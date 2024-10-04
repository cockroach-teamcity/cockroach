// Copyright 2021 The Cockroach Authors.
//
// Use of this software is governed by the CockroachDB Software License
// included in the /LICENSE file.

package upgrades_test

import (
	"context"
	"testing"

	"github.com/cockroachdb/cockroach/pkg/base"
	"github.com/cockroachdb/cockroach/pkg/clusterversion"
	"github.com/cockroachdb/cockroach/pkg/server"
	"github.com/cockroachdb/cockroach/pkg/testutils/sqlutils"
	"github.com/cockroachdb/cockroach/pkg/testutils/testcluster"
	"github.com/cockroachdb/cockroach/pkg/util/leaktest"
	"github.com/cockroachdb/cockroach/pkg/util/log"
)

func TestIsAtLeastVersionBuiltin(t *testing.T) {
	defer leaktest.AfterTest(t)()
	defer log.Scope(t).Close(t)

	clusterArgs := base.TestClusterArgs{
		ServerArgs: base.TestServerArgs{
			Knobs: base.TestingKnobs{
				Server: &server.TestingKnobs{
					DisableAutomaticVersionUpgrade: make(chan struct{}),
					BinaryVersionOverride:          clusterversion.ByKey(clusterversion.V22_2),
				},
			},
		},
	}

	var (
		ctx   = context.Background()
		tc    = testcluster.StartTestCluster(t, 1, clusterArgs)
		conn  = tc.ServerConn(0)
		sqlDB = sqlutils.MakeSQLRunner(conn)
	)
	defer tc.Stopper().Stop(ctx)

	v := clusterversion.ByKey(clusterversion.V23_1Start).String()
	// Check that the builtin returns false when comparing against the new version
	// version because we are still on the bootstrap version.
	sqlDB.CheckQueryResults(t, "SELECT crdb_internal.is_at_least_version('"+v+"')", [][]string{{"false"}})

	// Run the upgrade.
	sqlDB.Exec(t, "SET CLUSTER SETTING version = $1", v)

	// It should now return true.
	sqlDB.CheckQueryResultsRetry(t, "SELECT crdb_internal.is_at_least_version('"+v+"')", [][]string{{"true"}})
}
