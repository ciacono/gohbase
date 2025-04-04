// Copyright (C) 2015  The GoHBase Authors.  All rights reserved.
// This file is part of GoHBase.
// Use of this source code is governed by the Apache License 2.0
// that can be found in the COPYING file.

package gohbase

import (
	"context"
	"encoding/binary"
	"encoding/json"
	"fmt"
	"log/slog"
	"net"
	"sync"
	"time"

	"github.com/tsuna/gohbase/compression"
	"github.com/tsuna/gohbase/hrpc"
	"github.com/tsuna/gohbase/pb"
	"github.com/tsuna/gohbase/region"
	"github.com/tsuna/gohbase/zk"
	"google.golang.org/protobuf/proto"
	"modernc.org/b/v2"
)

const (
	defaultRPCQueueSize  = 100
	defaultFlushInterval = 20 * time.Millisecond
	defaultZkRoot        = "/hbase"
	defaultZkTimeout     = 30 * time.Second
	defaultEffectiveUser = "root"
)

// Client a regular HBase client
type Client interface {
	Scan(s *hrpc.Scan) hrpc.Scanner
	Get(g *hrpc.Get) (*hrpc.Result, error)
	Put(p *hrpc.Mutate) (*hrpc.Result, error)
	Delete(d *hrpc.Mutate) (*hrpc.Result, error)
	Append(a *hrpc.Mutate) (*hrpc.Result, error)
	Increment(i *hrpc.Mutate) (int64, error)
	CheckAndPut(p *hrpc.Mutate, family string, qualifier string,
		expectedValue []byte) (bool, error)
	SendBatch(ctx context.Context, batch []hrpc.Call) (res []hrpc.RPCResult, allOK bool)
	CacheRegions(table []byte) error
	Close()
}

// RPCClient is core client of gohbase. It's exposed for testing.
type RPCClient interface {
	SendRPC(rpc hrpc.Call) (proto.Message, error)
}

// Option is a function used to configure optional config items for a Client.
type Option func(*client)

// A Client provides access to an HBase cluster.
type client struct {
	clientType region.ClientType

	regions keyRegionCache

	// Maps a hrpc.RegionInfo to the *region.Client that we think currently
	// serves it.
	clients clientRegionCache

	metaRegionInfo hrpc.RegionInfo

	adminRegionInfo hrpc.RegionInfo

	// The maximum size of the RPC queue in the region client
	rpcQueueSize int

	// zkClient is zookeeper for retrieving meta and admin information
	zkClient zk.Client

	// The root zookeeper path for Hbase. By default, this is usually "/hbase".
	zkRoot string

	// The zookeeper session timeout
	zkTimeout time.Duration

	// The timeout before flushing the RPC queue in the region client
	flushInterval time.Duration

	// The user used when accessing regions.
	effectiveUser string

	// How long to wait for a region lookup (either meta lookup or finding
	// meta in ZooKeeper).  Should be greater than or equal to the ZooKeeper
	// session timeout.
	regionLookupTimeout time.Duration

	// regionReadTimeout is the maximum amount of time to wait for regionserver reply
	regionReadTimeout time.Duration

	done      chan struct{}
	closeOnce sync.Once

	newRegionClientFn func(string, region.ClientType, int, time.Duration,
		string, time.Duration, compression.Codec,
		func(ctx context.Context, network, addr string) (net.Conn, error),
		*slog.Logger) hrpc.RegionClient

	compressionCodec compression.Codec

	// zkDialer is used in the zkClient to connect to the quorum
	zkDialer func(ctx context.Context, network, addr string) (net.Conn, error)
	// regionDialer is passed into the region client to connect to hbase in a custom way,
	// such as SOCKS proxy.
	regionDialer func(ctx context.Context, network, addr string) (net.Conn, error)
	// logger that could be defined by user
	logger *slog.Logger
}

// NewClient creates a new HBase client.
func NewClient(zkquorum string, options ...Option) Client {
	return newClient(zkquorum, options...)
}

func newClient(zkquorum string, options ...Option) *client {
	c := &client{
		clientType:    region.RegionClient,
		rpcQueueSize:  defaultRPCQueueSize,
		flushInterval: defaultFlushInterval,
		metaRegionInfo: region.NewInfo(
			0,
			[]byte("hbase"),
			[]byte("meta"),
			[]byte("hbase:meta,,1"),
			nil,
			nil),
		zkRoot:              defaultZkRoot,
		zkTimeout:           defaultZkTimeout,
		effectiveUser:       defaultEffectiveUser,
		regionLookupTimeout: region.DefaultLookupTimeout,
		regionReadTimeout:   region.DefaultReadTimeout,
		done:                make(chan struct{}),
		newRegionClientFn:   region.NewClient,
		logger:              slog.Default(),
	}
	for _, option := range options {
		option(c)
	}
	c.logger.Debug("Creating new client.", "Host", slog.StringValue(zkquorum))

	//Have to create the zkClient after the Options have been set
	//since the zkTimeout could be changed as an option
	c.zkClient = zk.NewClient(zkquorum, c.zkTimeout, c.zkDialer, c.logger)
	c.regions = keyRegionCache{
		logger:  c.logger,
		regions: b.TreeNew[[]byte, hrpc.RegionInfo](region.Compare),
	}
	c.clients = clientRegionCache{
		logger:  c.logger,
		regions: make(map[hrpc.RegionClient]map[hrpc.RegionInfo]struct{}),
	}

	return c
}

// DebugState information about the clients keyRegionCache, and clientRegionCache
func DebugState(c Client) ([]byte, error) {

	debugInfoJson, err := json.Marshal(c)
	if err != nil {
		if cclient, ok := c.(*client); ok {
			cclient.logger.Error("Cannot turn client into JSON bytes array", "error", err)
		} else {
			slog.Error("Cannot turn client into JSON bytes array", "error", err)
		}
	}
	return debugInfoJson, err
}

func (c *client) MarshalJSON() ([]byte, error) {

	var done string
	if c.done != nil {
		select {
		case <-c.done:
			done = "Closed"
		default:
			done = "Not Closed"
		}
	} else {
		done = "nil"
	}

	rcc := &c.clients
	krc := &c.regions

	// create map for all ClientRegions (clntRegion Ptr -> JSONified Client Region)
	clientRegionsMap := map[string]hrpc.RegionClient{}
	// create map for all RegionInfos (Region Ptr -> JSONified RegionInfo)
	keyRegionInfosMap := map[string]hrpc.RegionInfo{}

	clientRegionCacheValues := rcc.debugInfo(keyRegionInfosMap, clientRegionsMap)
	keyRegionCacheValues := krc.debugInfo(keyRegionInfosMap)

	state := struct {
		ClientType          region.ClientType
		ClientRegionMap     map[string]hrpc.RegionClient
		RegionInfoMap       map[string]hrpc.RegionInfo
		KeyRegionCache      map[string]string
		ClientRegionCache   map[string][]string
		MetaRegionInfo      hrpc.RegionInfo
		AdminRegionInfo     hrpc.RegionInfo
		Done_Status         string
		RegionLookupTimeout time.Duration
		RegionReadTimeout   time.Duration
	}{
		ClientType:          c.clientType,
		ClientRegionMap:     clientRegionsMap,
		RegionInfoMap:       keyRegionInfosMap,
		KeyRegionCache:      keyRegionCacheValues,
		ClientRegionCache:   clientRegionCacheValues,
		MetaRegionInfo:      c.metaRegionInfo,
		AdminRegionInfo:     c.adminRegionInfo,
		Done_Status:         done,
		RegionLookupTimeout: c.regionLookupTimeout,
		RegionReadTimeout:   c.regionReadTimeout,
	}

	jsonVal, err := json.Marshal(state)
	return jsonVal, err
}

// RpcQueueSize will return an option that will set the size of the RPC queues
// used in a given client
func RpcQueueSize(size int) Option {
	return func(c *client) {
		c.rpcQueueSize = size
	}
}

// ZookeeperRoot will return an option that will set the zookeeper root path used in a given client.
func ZookeeperRoot(root string) Option {
	return func(c *client) {
		c.zkRoot = root
	}
}

// ZookeeperTimeout will return an option that will set the zookeeper session timeout.
func ZookeeperTimeout(to time.Duration) Option {
	return func(c *client) {
		c.zkTimeout = to
	}
}

// RegionLookupTimeout will return an option that sets the region lookup timeout
func RegionLookupTimeout(to time.Duration) Option {
	return func(c *client) {
		c.regionLookupTimeout = to
	}
}

// RegionReadTimeout will return an option that sets the region read timeout
func RegionReadTimeout(to time.Duration) Option {
	return func(c *client) {
		c.regionReadTimeout = to
	}
}

// EffectiveUser will return an option that will set the user used when accessing regions.
func EffectiveUser(user string) Option {
	return func(c *client) {
		c.effectiveUser = user
	}
}

// FlushInterval will return an option that will set the timeout for flushing
// the RPC queues used in a given client
func FlushInterval(interval time.Duration) Option {
	return func(c *client) {
		c.flushInterval = interval
	}
}

// CompressionCodec will return an option to set compression codec between
// client and server. The only currently supported codec is "snappy".
// Default is no compression.
func CompressionCodec(codec string) Option {
	return func(c *client) {
		c.compressionCodec = compression.New(codec)
	}
}

// ZooKeeperDialer will return an option to pass the given dialer function
// into the ZooKeeper client Connect() call, which allows for customizing
// network connections.
func ZooKeeperDialer(dialer func(
	ctx context.Context, network, addr string) (net.Conn, error)) Option {
	return func(c *client) {
		c.zkDialer = dialer
	}
}

// RegionDialer will return an option that uses the specified Dialer for
// connecting to region servers. This allows for connecting through proxies.
func RegionDialer(dialer func(
	ctx context.Context, network, addr string) (net.Conn, error)) Option {
	return func(c *client) {
		c.regionDialer = dialer
	}
}

// Logger will return an option to set *slog.Logger instance
func Logger(logger *slog.Logger) Option {
	return func(c *client) {
		c.logger = logger
	}
}

// Close closes connections to hbase master and regionservers
func (c *client) Close() {
	c.closeOnce.Do(func() {
		close(c.done)
		if c.clientType == region.MasterClient {
			if ac := c.adminRegionInfo.Client(); ac != nil {
				ac.Close()
			}
		}
		c.clients.closeAll()
	})
}

func (c *client) Scan(s *hrpc.Scan) hrpc.Scanner {
	return newScanner(c, s, c.logger)
}

func (c *client) Get(g *hrpc.Get) (*hrpc.Result, error) {
	pbmsg, err := c.SendRPC(g)
	if err != nil {
		return nil, err
	}

	r, ok := pbmsg.(*pb.GetResponse)
	if !ok {
		return nil, fmt.Errorf("sendRPC returned not a GetResponse")
	}

	return hrpc.ToLocalResult(r.Result), nil
}

func (c *client) Put(p *hrpc.Mutate) (*hrpc.Result, error) {
	return c.mutate(p)
}

func (c *client) Delete(d *hrpc.Mutate) (*hrpc.Result, error) {
	return c.mutate(d)
}

func (c *client) Append(a *hrpc.Mutate) (*hrpc.Result, error) {
	return c.mutate(a)
}

func (c *client) Increment(i *hrpc.Mutate) (int64, error) {
	r, err := c.mutate(i)
	if err != nil {
		return 0, err
	}

	if len(r.Cells) != 1 {
		return 0, fmt.Errorf("increment returned %d cells, but we expected exactly one",
			len(r.Cells))
	}

	val := binary.BigEndian.Uint64(r.Cells[0].Value)
	return int64(val), nil
}

func (c *client) mutate(m *hrpc.Mutate) (*hrpc.Result, error) {
	pbmsg, err := c.SendRPC(m)
	if err != nil {
		return nil, err
	}

	r, ok := pbmsg.(*pb.MutateResponse)
	if !ok {
		return nil, fmt.Errorf("sendRPC returned not a MutateResponse")
	}

	return hrpc.ToLocalResult(r.Result), nil
}

func (c *client) CheckAndPut(p *hrpc.Mutate, family string,
	qualifier string, expectedValue []byte) (bool, error) {
	cas, err := hrpc.NewCheckAndPut(p, family, qualifier, expectedValue)
	if err != nil {
		return false, err
	}

	pbmsg, err := c.SendRPC(cas)
	if err != nil {
		return false, err
	}

	r, ok := pbmsg.(*pb.MutateResponse)
	if !ok {
		return false, fmt.Errorf("sendRPC returned a %T instead of MutateResponse", pbmsg)
	}

	if r.Processed == nil {
		return false, fmt.Errorf("protobuf in the response didn't contain the field "+
			"indicating whether the CheckAndPut was successful or not: %s", r)
	}

	return r.GetProcessed(), nil
}

// CacheRegions scan the meta region to get all the regions and populate to cache.
// This can be used to warm up region cache
func (c *client) CacheRegions(table []byte) error {
	_, err := c.findAllRegions(context.Background(), table)
	return err
}
