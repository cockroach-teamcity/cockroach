// Copyright 2016 The Cockroach Authors.
//
// Use of this software is governed by the CockroachDB Software License
// included in the /LICENSE file.

package main

import (
	"context"
	"os"
	"path/filepath"
	"reflect"
	"strconv"
	"strings"
	"testing"

	"github.com/cockroachdb/cockroach/pkg/testutils/datapathutils"
	"github.com/cockroachdb/cockroach/pkg/testutils/skip"
	"github.com/kr/pretty"
)

func TestPkgsFromDiff(t *testing.T) {
	for filename, expPkgs := range map[string]map[string]pkg{
		datapathutils.TestDataPath(t, "10305.diff"): {
			"pkg/roachpb": {tests: []string{"TestLeaseEquivalence"}},
			"pkg/storage": {tests: []string{"TestStoreRangeLease", "TestStoreRangeLeaseSwitcheroo"}},
		},
		datapathutils.TestDataPath(t, "skip.diff"): {
			"pkg/ccl/storageccl": {tests: []string{"TestPutS3"}},
		},
		// This PR had some churn and renamed packages. This was formerly problematic
		// because nonexistent packages would be emitted.
		datapathutils.TestDataPath(t, "27595.diff"): {
			"pkg/storage/closedts/transport": {tests: []string{"TestTransportConnectOnRequest", "TestTransportClientReceivesEntries"}},
			"pkg/storage/closedts/container": {tests: []string{"TestTwoNodes"}},
			"pkg/storage/closedts/storage":   {tests: []string{"TestConcurrent"}},
		},
		datapathutils.TestDataPath(t, "removed.diff"): {},
	} {
		t.Run(filename, func(t *testing.T) {
			f, err := os.Open(filename)
			if err != nil {
				t.Fatal(err)
			}
			defer f.Close()
			pkgs, err := pkgsFromDiff(f)
			if err != nil {
				t.Fatal(err)
			}
			if !reflect.DeepEqual(pkgs, expPkgs) {
				t.Errorf("expected %s, got %s", expPkgs, pkgs)
			}
		})
	}
}

func TestPkgsFromDiffHelper(t *testing.T) {
	// This helper can easily generate new test cases.
	skip.IgnoreLint(t, "only for manual use")

	ctx := context.Background()
	client := ghClient(ctx)

	const prNum = 10305

	diff, err := getDiff(ctx, client, "cockroachdb", "cockroach", prNum)
	if err != nil {
		t.Fatal(err)
	}
	wd, err := os.Getwd()
	if err != nil {
		t.Fatal(err)
	}

	name := filepath.Join(wd, "testdata", strconv.Itoa(prNum)+".diff")
	if err := os.WriteFile(name, []byte(diff), 0644); err != nil {
		t.Fatal(err)
	}

	pkgs, err := pkgsFromDiff(strings.NewReader(diff))
	if err != nil {
		t.Fatal(err)
	}
	t.Errorf("read the following information:\n%v\n\ndiff at %s", pretty.Sprint(pkgs), name)
}
