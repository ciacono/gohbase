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

func TestNewCompactRegion(t *testing.T) {
	var (
		tb = []byte("tablename")
		// Using pb.RegionSpecifier_REGION_NAME which looks like:
		// <tablename>,<startkey>,<regionId>.<encodedName>
		n = []byte("tablename,aa,12356,asdfghjkl")
		f = "familyyy"
	)

	// No opts:
	cr, err := NewCompactRegion(context.Background(), n, tb)
	if err != nil {
		t.Error(err)
	}
	if cr.regionSpecifier.GetType() != pb.RegionSpecifier_REGION_NAME &&
		!bytes.Equal(cr.regionSpecifier.GetValue(), n) {
		t.Errorf("RegionSpecifier not set correctly in CompactRegion, got: %s",
			cr.regionSpecifier.String())
	}
	if cr.major {
		t.Errorf("Expected WithMajor to not be set")
	}
	if cr.family != "" {
		t.Errorf("Expected WithFamily to not be set")
	}

	// With opts:
	cr, err = NewCompactRegion(context.Background(), n, tb, WithMajor(true), WithFamily(f))
	if err != nil {
		t.Error(err)
	}
	if cr.regionSpecifier.GetType() != pb.RegionSpecifier_REGION_NAME &&
		!bytes.Equal(cr.regionSpecifier.GetValue(), n) {
		t.Errorf("RegionSpecifier not set correctly in CompactRegion, got: %s",
			cr.regionSpecifier.String())
	}
	if !cr.major {
		t.Errorf("Expected WithMajor to be true")
	}
	if cr.family != f {
		t.Errorf("Expected WithFamily to be set, got %s", cr.family)
	}

	// Encoded:
	en := []byte("encodedname")
	cr, err = NewCompactRegionEncoded(context.Background(), en, tb)
	if err != nil {
		t.Error(err)
	}
	if cr.regionSpecifier.GetType() != pb.RegionSpecifier_ENCODED_REGION_NAME &&
		!bytes.Equal(cr.regionSpecifier.GetValue(), en) {
		t.Errorf("RegionSpecifier not set correctly in CompactRegion, got: %s",
			cr.regionSpecifier.String())
	}
	if cr.major {
		t.Errorf("Expected WithMajor to not be set")
	}
	if cr.family != "" {
		t.Errorf("Expected WithFamily to not be set")
	}
}
