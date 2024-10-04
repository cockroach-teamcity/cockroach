// Copyright 2023 The Cockroach Authors.
//
// Use of this software is governed by the CockroachDB Software License
// included in the /LICENSE file.

package upgrades

import (
	"context"
	"fmt"

	"github.com/cockroachdb/cockroach/pkg/clusterversion"
	"github.com/cockroachdb/cockroach/pkg/sql/sem/tree"
	"github.com/cockroachdb/cockroach/pkg/upgrade"
	"github.com/cockroachdb/cockroach/pkg/util/envutil"
	"github.com/cockroachdb/cockroach/pkg/util/log"
	"github.com/cockroachdb/errors"
)

// RunFirstUpgradePrecondition short-circuits FirstUpgradeFromReleasePrecondition if set to false.
var RunFirstUpgradePrecondition = envutil.EnvOrDefaultBool("COCKROACH_RUN_FIRST_UPGRADE_PRECONDITION", false)

// FirstUpgradeFromReleasePrecondition is the precondition check for upgrading
// from any supported major release.
//
// This precondition function performs health checks on the catalog metadata.
// This prevents cluster version upgrades from proceeding when it detects
// known corruptions. This forces cluster operations to repair the metadata,
// which may be somewhat annoying, but it's considered a lesser evil to being
// stuck in a mixed-version state due to a later upgrade step throwing
// errors.
//
// The implementation of this function may evolve in the future to include
// other checks beyond querying the invalid_objects virtual table, which
// mainly checks descriptor validity, as that has historically been the main
// source of problems.
func FirstUpgradeFromReleasePrecondition(
	ctx context.Context, _ clusterversion.ClusterVersion, d upgrade.TenantDeps,
) error {
	if !RunFirstUpgradePrecondition {
		return nil
	}
	// For performance reasons, we look back in time when performing
	// a diagnostic query. If no corruptions were found back then, we assume that
	// there are no corruptions now. Otherwise, we retry and do everything
	// without an AOST clause henceforth.
	withAOST := true
	diagnose := func(tbl string) (hasRows bool, err error) {
		q := fmt.Sprintf("SELECT count(*) FROM \"\".crdb_internal.%s", tbl)
		if withAOST {
			q = q + " AS OF SYSTEM TIME '-10s'"
		}
		row, err := d.InternalExecutor.QueryRow(ctx, "query-"+tbl, nil /* txn */, q)
		if err == nil && row[0].String() != "0" {
			hasRows = true
		}
		return hasRows, err
	}
	// Check for possibility of time travel.
	if hasRows, err := diagnose("databases"); err != nil {
		return err
	} else if !hasRows {
		// We're looking back in time to before the cluster was bootstrapped
		// and no databases exist at that point. Disable time-travel henceforth.
		withAOST = false
	}
	// Check for repairable catalog corruptions.
	if hasRows, err := diagnose("kv_repairable_catalog_corruptions"); err != nil {
		return err
	} else if hasRows {
		// Attempt to repair catalog corruptions in batches.
		log.Info(ctx, "auto-repairing catalog corruptions detected during upgrade attempt")
		var n int
		const repairQuery = `
SELECT
	count(*)
FROM
	(
		SELECT
			crdb_internal.repair_catalog_corruption(id, corruption) AS was_repaired
		FROM
			"".crdb_internal.kv_repairable_catalog_corruptions
		LIMIT
			1000
	)
WHERE
	was_repaired`
		for {
			row, err := d.InternalExecutor.QueryRow(
				ctx, "repair-catalog-corruptions", nil /* txn */, repairQuery,
			)
			if err != nil {
				return err
			}
			c := tree.MustBeDInt(row[0])
			if c == 0 {
				break
			}
			n += int(c)
			log.Infof(ctx, "repaired %d catalog corruptions", c)
		}
		if n == 0 {
			log.Info(ctx, "no catalog corruptions found to repair during upgrade attempt")
		} else {
			// Repairs have actually been performed: stop all time travel henceforth.
			withAOST = false
			log.Infof(ctx, "%d catalog corruptions have been repaired in total", n)
		}
	}
	// Check for all known catalog corruptions.
	if hasRows, err := diagnose("invalid_objects"); err != nil {
		return err
	} else if !hasRows {
		return nil
	}
	if !withAOST {
		return errors.AssertionFailedf("\"\".crdb_internal.invalid_objects is not empty")
	}
	// At this point, corruptions were found using the AS OF SYSTEM TIME clause.
	// Re-run the diagnosis without the clause, because we might not be seeing
	// repairs which might have taken place recently.
	withAOST = false
	if hasRows, err := diagnose("invalid_objects"); err != nil {
		return err
	} else if !hasRows {
		return nil
	}
	return errors.AssertionFailedf("\"\".crdb_internal.invalid_objects is not empty")
}
