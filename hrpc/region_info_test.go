// Copyright (C) 2024  The GoHBase Authors.  All rights reserved.
// This file is part of GoHBase.
// Use of this source code is governed by the Apache License 2.0
// that can be found in the COPYING file.

package hrpc

import (
	"bytes"
	"context"
	"testing"
)

func TestGetRegionInfo(t *testing.T) {
	rn := []byte("name of a region")
	gri, err := NewGetRegionInfo(context.Background(), rn)
	if err != nil {
		t.Fatal(err)
	}
	validateGetRegionInfo(t, gri, rn, false)

	gri, err = NewGetRegionInfo(context.Background(), rn, WithCompactionState(true))
	if err != nil {
		t.Fatal(err)
	}
	validateGetRegionInfo(t, gri, rn, true)

	gri, err = NewGetRegionInfo(context.Background(), rn, WithCompactionState(false))
	if err != nil {
		t.Fatal(err)
	}
	validateGetRegionInfo(t, gri, rn, false)
}

func validateGetRegionInfo(t *testing.T, actual *GetRegionInfo,
	expectedName []byte, expectedCompactionState bool) {
	if !bytes.Equal(actual.regionName, expectedName) {
		t.Fatalf("region name is %s, want %s", actual.regionName, expectedName)
	}
	if actual.compactionState != expectedCompactionState {
		t.Fatalf("CompactionState should be set to %v.", expectedCompactionState)
	}
}
