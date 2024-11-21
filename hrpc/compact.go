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
	req             *pb.CompactRegionRequest
	regionSpecifier *pb.RegionSpecifier
	major           bool
	family          string
}

// Name returns the name of this RPC call.
func (c *CompactRegion) Name() string {
	return "CompactRegion"
}

// ToProto converts the RPC into a protobuf message.
func (c *CompactRegion) ToProto() proto.Message {
	return &pb.CompactRegionRequest{
		Region: c.regionSpecifier,
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

// NewCompactRegion creates a CompactRegionRequest which will compact the given region in the
// given HBase table.
func NewCompactRegion(ctx context.Context, regionName []byte, table []byte,
	opts ...func(Call) error) (*CompactRegion, error) {
	b := base{
		table:    table,
		ctx:      ctx,
		resultch: make(chan RPCResult, 1),
	}
	rs := &pb.RegionSpecifier{
		Type:  RegionSpecifierRegionName,
		Value: regionName,
	}
	cr := &CompactRegion{
		base: b,
		req: &pb.CompactRegionRequest{
			Region: rs,
		},
		regionSpecifier: rs,
	}
	if err := applyOptions(cr, opts...); err != nil {
		return nil, err
	}
	return cr, nil
}

// NewCompactRegionEncoded creates a CompactRegionRequest which will compact the given region in the
// given HBase table. It uses the encoded region name to form the request
func NewCompactRegionEncoded(ctx context.Context, encodedRegion []byte, table []byte,
	opts ...func(Call) error) (*CompactRegion, error) {
	b := base{
		table:    table,
		ctx:      ctx,
		resultch: make(chan RPCResult, 1),
	}
	rs := &pb.RegionSpecifier{
		Type:  pb.RegionSpecifier_ENCODED_REGION_NAME.Enum(),
		Value: encodedRegion,
	}
	cr := &CompactRegion{
		base: b,
		req: &pb.CompactRegionRequest{
			Region: rs,
		},
		regionSpecifier: rs,
	}
	if err := applyOptions(cr, opts...); err != nil {
		return nil, err
	}
	return cr, nil
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
