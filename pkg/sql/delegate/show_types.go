// Copyright 2020 The Cockroach Authors.
//
// Use of this software is governed by the Business Source License
// included in the file licenses/BSL.txt.
//
// As of the Change Date specified in that file, in accordance with
// the Business Source License, use of this software will be governed
// by the Apache License, Version 2.0, included in the file
// licenses/APL.txt.

package delegate

import (
	"fmt"

	"github.com/cockroachdb/cockroach/pkg/sql/lexbase"
	"github.com/cockroachdb/cockroach/pkg/sql/sem/catconstants"
	"github.com/cockroachdb/cockroach/pkg/sql/sem/tree"
	"github.com/cockroachdb/cockroach/pkg/sql/sqltelemetry"
)

func (d *delegator) delegateShowTypes(n *tree.ShowTypes) (tree.Statement, error) {

	dbName := lexbase.EscapeSQLIdent(d.evalCtx.SessionData().Database)
	commentColumn, commentJoin := ``, ``
	if n.WithComment {
		commentColumn = `, comment`
		commentJoin = fmt.Sprintf(`
			LEFT JOIN
				(
					SELECT 
						objoid, description as comment
					FROM
						%s.pg_catalog.pg_description
					WHERE
						classoid = %d
				) c
			ON
				 	(types.oid::int - 100000) = c.objoid`, dbName, catconstants.PgCatalogTypeTableID)
	}

	// Query all enum and composite types. Two explanations about the WHERE
	// clause we use:
	// - we filter out types defined in internally generated namespaces so that
	//   we only show user defined types.
	// - we filter based on typrelid to omit the type for the table record. It
	//   must be 0, for enums, or point back to its own oid for a standalone
	//   composite type.
	query := fmt.Sprintf(`
SELECT
	nsp.nspname AS schema,
	types.typname AS name,
	rl.rolname AS owner %[2]s
FROM
	%[1]s.pg_catalog.pg_type AS types
	LEFT JOIN %[1]s.pg_catalog.pg_roles AS rl on (types.typowner = rl.oid)
	JOIN %[1]s.pg_catalog.pg_namespace AS nsp ON (types.typnamespace = nsp.oid)
	%[3]s
WHERE types.typtype IN ('e','c') AND
      types.typrelid IN (0, types.oid) AND
      nsp.nspname NOT IN ('information_schema', 'pg_catalog', 'crdb_internal', 'pg_extension')
ORDER BY (nsp.nspname, types.typname)
`, dbName, commentColumn, commentJoin)
	return d.parse(query)
}

func (d *delegator) delegateShowCreateAllTypes() (tree.Statement, error) {
	sqltelemetry.IncrementShowCounter(sqltelemetry.Create)

	const showCreateAllTypesQuery = `
	SELECT crdb_internal.show_create_all_types(%[1]s) AS create_statement;
`
	databaseLiteral := d.evalCtx.SessionData().Database

	query := fmt.Sprintf(showCreateAllTypesQuery,
		lexbase.EscapeSQLString(databaseLiteral),
	)

	return d.parse(query)
}
