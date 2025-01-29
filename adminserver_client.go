// Copyright (C) 2025  The GoHBase Authors.  All rights reserved.
// This file is part of GoHBase.
// Use of this source code is governed by the Apache License 2.0
// that can be found in the COPYING file.

package gohbase

import (
	"log/slog"

	"github.com/tsuna/gohbase/region"
	"github.com/tsuna/gohbase/zk"
)

// AdminServerClient is a HBase client that performs Admin Service functions
// TODO - naming
type AdminServerClient interface {
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
