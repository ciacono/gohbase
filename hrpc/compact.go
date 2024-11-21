// Copyright (C) 2024  The GoHBase Authors.  All rights reserved.
// This file is part of GoHBase.
// Use of this source code is governed by the Apache License 2.0
// that can be found in the COPYING file.

package hrpc

import (
	"context"
	"fmt"

	"github.com/tsuna/gohbase/pb"
	"google.golang.org/protobuf/proto"
)

type CompactRegion struct {
	base
	req    *pb.CompactRegionRequest
	region *pb.RegionSpecifier
	major  bool
	family string
}

// Name returns the name of this RPC call.
func (c *CompactRegion) Name() string {
	return "CompactRegion"
}

// ToProto converts the RPC into a protobuf message.
func (c *CompactRegion) ToProto() proto.Message {
	return &pb.CompactRegionRequest{
		Region: nil, // TODO
		Major:  proto.Bool(c.major),
		Family: []byte(c.family),
	}
}

// NewResponse creates an empty protobuf message to read the response of this RPC.
func (c *CompactRegion) NewResponse() proto.Message {
	return &pb.CompactRegionResponse{}
}

// Description returns the description of this RPC call.
func (c *CompactRegion) Description() string {
	return c.Name()
}

// NewCompactRegion is TODO.
func NewCompactRegion(ctx context.Context, opts ...func(Call) error) (*CompactRegion, error) {
	cr := &CompactRegion{
		base: base{},
		req:  &pb.CompactRegionRequest{},
	}
	if err := applyOptions(cr, opts...); err != nil {
		return nil, err
	}
	return nil, fmt.Errorf("not yet implemented")
}

// WithMajor determines if the compaction requested will be major or minor.
// To do a major compaction, set major to true, else it will be minor.
func WithMajor(major bool) func(Call) error {
	return func(c Call) error {
		cr, ok := c.(*CompactRegion)
		if !ok {
			return fmt.Errorf("WithMajor can only be used with CompactRegion")
		}
		cr.major = major
		return nil
	}
}

// WithFamily sets the ColumnFamily for the compaction request.
func WithFamily(family string) func(Call) error {
	return func(c Call) error {
		cr, ok := c.(*CompactRegion)
		if !ok {
			return fmt.Errorf("WithFamily can only be used with CompactRegion")
		}
		cr.family = family
		return nil
	}
}

// WithRegion is TODO.
func WithRegion(region *pb.RegionSpecifier) func(Call) error {
	return func(c Call) error {
		cr, ok := c.(*CompactRegion)
		if !ok {
			return fmt.Errorf("WithRegion can only be used with CompactRegion")
		}
		cr.region = region
		return nil
	}
}
