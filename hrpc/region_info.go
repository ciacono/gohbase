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

type GetRegionInfo struct {
	base
	regionSpecifier *pb.RegionSpecifier
	compactionState bool
}

func (gr *GetRegionInfo) Name() string {
	return "GetRegionInfo"
}

func (gr *GetRegionInfo) ToProto() proto.Message {
	return &pb.GetRegionInfoRequest{
		Region:          gr.regionSpecifier,
		CompactionState: proto.Bool(gr.compactionState),
	}
}

func (gr *GetRegionInfo) NewResponse() proto.Message {
	return &pb.GetRegionInfoResponse{}
}

func (gr *GetRegionInfo) Description() string {
	return gr.Name()
}

func NewGetRegionInfo(ctx context.Context, rs *pb.RegionSpecifier,
	opts ...func(Call) error) (*GetRegionInfo, error) {
	gr := &GetRegionInfo{
		base: base{
			ctx:      ctx,
			resultch: make(chan RPCResult, 1),
		},
		regionSpecifier: rs,
		compactionState: false, // By default, do not retrieve the compaction state.
	}
	if err := applyOptions(gr, opts...); err != nil {
		return nil, err
	}
	return gr, nil
}

// WithCompactionState is used to request the compaction state of the region in the GetRegionInfo
// request when include is set to true.
func WithCompactionState(include bool) func(Call) error {
	return func(call Call) error {
		gr, ok := call.(*GetRegionInfo)
		if !ok {
			return fmt.Errorf("WithCompactionState can only be used with GetRegionInfo")
		}
		gr.compactionState = include
		return nil
	}
}
