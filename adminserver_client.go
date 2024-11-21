// Copyright (C) 2025  The GoHBase Authors.  All rights reserved.
// This file is part of GoHBase.
// Use of this source code is governed by the Apache License 2.0
// that can be found in the COPYING file.

package gohbase

import (
	"errors"
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
