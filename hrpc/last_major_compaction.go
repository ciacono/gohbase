// Copyright (C) 2025  The GoHBase Authors.  All rights reserved.
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

type LastMajorCompaction struct {
	base
	// A note on namespaces:
	// Two special namespaces:
	// 1. "hbase" - system namespace, used to contain hbase internal tables
	// 2. "default" - tables with no explicit specified namespace will automatically fall into
	// this namespace.
	namespace []byte
}

// NewLastMajorCompaction creates a hrpc to get the timestamp of the last major compaction for
// a table. By default, uses the "default" namespace unless option specified.
func NewLastMajorCompaction(ctx context.Context, table []byte,
	opts ...func(Call) error) (*LastMajorCompaction, error) {
	lr := &LastMajorCompaction{
		namespace: []byte("default"),
		base: base{
			ctx:      ctx,
			table:    table,
			resultch: make(chan RPCResult, 1),
		},
	}
	if err := applyOptions(lr, opts...); err != nil {
		return nil, err
	}

	return lr, nil
}

func WithNamespace(namespace []byte) func(Call) error {
	return func(call Call) error {
		lr, ok := call.(*LastMajorCompaction)
		if !ok {
			return fmt.Errorf("WithNameSpace can only be used with LastMajorCompaction")
		}
		lr.namespace = namespace
		return nil
	}
}

func (l *LastMajorCompaction) Name() string {
	return "LastMajorCompaction"
}

func (l *LastMajorCompaction) ToProto() proto.Message {
	return &pb.MajorCompactionTimestampRequest{
		TableName: &pb.TableName{
			Namespace: l.namespace,
			Qualifier: l.Table(),
		},
	}
}

func (l *LastMajorCompaction) NewResponse() proto.Message {
	return &pb.MajorCompactionTimestampResponse{}
}

func (l *LastMajorCompaction) Description() string {
	return l.Name()
}
