// Copyright (C) 2025  The GoHBase Authors.  All rights reserved.
// This file is part of GoHBase.
// Use of this source code is governed by the Apache License 2.0
// that can be found in the COPYING file.

package gohbase

import (
	"context"
	"errors"
	"fmt"
	"log/slog"

	"github.com/tsuna/gohbase/hrpc"
	"github.com/tsuna/gohbase/pb"
	"github.com/tsuna/gohbase/region"
	"github.com/tsuna/gohbase/zk"
)

// AdminServerClient is a HBase client that performs Admin Service functions
// TODO - naming
type AdminServerClient interface {
	GetRegionInfo(gr *hrpc.GetRegionInfo) (*pb.RegionInfo,
		pb.GetRegionInfoResponse_CompactionState, error)
	CompactRegion(cr *hrpc.CompactRegion) error
	GetLastMajorCompactionTimestamp(lr *hrpc.LastMajorCompaction) (int64, error)
	Compact(ctx context.Context, table []byte) error
}

func NewAdminServerClient(zkquorum string, options ...Option) AdminServerClient {
	return newAdminServerClient(zkquorum, options...)
}

func newAdminServerClient(zkquorum string, options ...Option) AdminServerClient {
	c := &client{
		clientType:    region.AdminServerClient,
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
	c.logger.Debug("Creating new admin server client.", "Host",
		slog.StringValue(zkquorum))

	c.zkClient = zk.NewClient(zkquorum, c.zkTimeout, c.zkDialer, c.logger)
	return c
}

// GetRegionInfo gets the region info for the given region.
// Current compaction state can be requested as well via the WithCompactionState option for the
// GetRegionInfo request.
func (c *client) GetRegionInfo(gr *hrpc.GetRegionInfo) (*pb.RegionInfo,
	pb.GetRegionInfoResponse_CompactionState, error) {
	// TODO don't use SendRPC
	pbmsg, err := c.SendRPC(gr)
	if err != nil {
		return nil, 0, err
	}
	res, ok := pbmsg.(*pb.GetRegionInfoResponse)
	if !ok {
		return nil, 0, errors.New("sendPRC returned not a GetRegionInfoResponse")
	}

	return res.GetRegionInfo(), res.GetCompactionState(), nil
}

// CompactRegion executes a compaction on the given region.
func (c *client) CompactRegion(cr *hrpc.CompactRegion) error {
	// TODO provide base.key to make sure it goes to the right regionserver.
	// Or lookup to find which regionserver has the ID
	pbmsg, err := c.SendRPC(cr) // TODO replace SendRPC
	if err != nil {
		return err
	}
	_, ok := pbmsg.(*pb.CompactRegionResponse)
	if !ok {
		return errors.New("SendPRC did not return a CompactRegionResponse")
	}
	return nil
}

// GetLastMajorCompactionTimestamp gets the timestamp of the last major compaction for the
// given table. The timestamp is for the oldest HFile resulting from a major compaction of that
// table, or 0 if there are no such HFiles for that table.
func (c *client) GetLastMajorCompactionTimestamp(lr *hrpc.LastMajorCompaction) (int64, error) {
	// TODO java.lang.UnsupportedOperationExceptionF
	// TODO probably don't need this one, can get with ClusterStatus
	pbmsg, err := c.SendRPC(lr)
	if err != nil {
		return 0, err
	}
	res, ok := pbmsg.(*pb.MajorCompactionTimestampResponse)
	if !ok {
		return 0, errors.New("sendPRC returned not a MajorCompactionTimestampResponse")
	}
	return res.GetCompactionTimestamp(), nil
}

// Compact performs a major compaction on the given table
func (c *client) Compact(ctx context.Context, table []byte) error {
	// Internal client to access meta:
	// TODO refactor this
	metaClient := newClient("TODO", RpcQueueSize(1))
	defer metaClient.Close()

	risAndAddrs, err := metaClient.findAllRegions(ctx, table)
	if err != nil {
		return fmt.Errorf("error finding all regions for table %v : %w", table, err)
	}
	c.logger.Debug(fmt.Sprintf("Found %d regions to compact for table %s", len(risAndAddrs),
		string(table)))

	// The HBase protobuf API only exposes CompactRegion currently, so need to send a CompactRegion
	// request for each region in the table.
	// TODO use SendBatch for all the Compacts?
	var cr *hrpc.CompactRegion
	for _, ria := range risAndAddrs {
		cr, err = hrpc.NewCompactRegion(ctx, ria.regionInfo.Name(), table,
			hrpc.WithMajor(true))
		if err != nil {
			return err
		}
		// TODO retryable?
		err = c.CompactRegion(cr) // TODO java.lang.UnsupportedOperationException
		if err != nil {
			return fmt.Errorf("error compacting region %v in table %v, error: %w",
				ria.regionInfo.Name(), table, err)
		}
		c.logger.Debug("SUCCESS Compacted a region")
	}

	return nil
}
