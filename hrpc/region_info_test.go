// Copyright (C) 2024  The GoHBase Authors.  All rights reserved.
// This file is part of GoHBase.
// Use of this source code is governed by the Apache License 2.0
// that can be found in the COPYING file.

package hrpc

import (
	"bytes"
	"context"
	"testing"

	"github.com/tsuna/gohbase/pb"
)

func TestGetRegionInfo(t *testing.T) {
	rn := []byte("name of a region")
	rs := &pb.RegionSpecifier{
		Type:  RegionSpecifierRegionName,
		Value: rn,
	}
	gri, err := NewGetRegionInfo(context.Background(), rs)
	if err != nil {
		t.Fatal(err)
	}
	validateGetRegionInfo(t, gri, rs, false)

	gri, err = NewGetRegionInfo(context.Background(), rs, WithCompactionState(true))
	if err != nil {
		t.Fatal(err)
	}
	validateGetRegionInfo(t, gri, rs, true)

	gri, err = NewGetRegionInfo(context.Background(), rs, WithCompactionState(false))
	if err != nil {
		t.Fatal(err)
	}
	validateGetRegionInfo(t, gri, rs, false)
}

func validateGetRegionInfo(t *testing.T, actual *GetRegionInfo,
	expectedRegionSpecifier *pb.RegionSpecifier, expectedCompactionState bool) {
	if !bytes.Equal(actual.regionSpecifier.GetValue(), expectedRegionSpecifier.GetValue()) {
		t.Fatalf("region name is %s, want %s", actual.regionSpecifier.GetValue(),
			expectedRegionSpecifier.GetValue())
	}
	if expectedRegionSpecifier.GetType() != actual.regionSpecifier.GetType() {
		t.Fatalf("region type is %d, want %d", actual.regionSpecifier.GetType(),
			expectedRegionSpecifier.GetType())
	}
	if actual.compactionState != expectedCompactionState {
		t.Fatalf("CompactionState should be set to %v.", expectedCompactionState)
	}
}
