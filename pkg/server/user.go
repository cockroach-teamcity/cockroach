// Copyright 2022 The Cockroach Authors.
//
// Use of this software is governed by the CockroachDB Software License
// included in the /LICENSE file.

package server

import (
	"context"

	"github.com/cockroachdb/cockroach/pkg/server/serverpb"
	"github.com/cockroachdb/cockroach/pkg/sql/privilege"
	"github.com/cockroachdb/cockroach/pkg/sql/roleoption"
)

// UserSQLRoles return a list of the logged in SQL user roles.
func (s *baseStatusServer) UserSQLRoles(
	ctx context.Context, req *serverpb.UserSQLRolesRequest,
) (_ *serverpb.UserSQLRolesResponse, retErr error) {
	ctx = forwardSQLIdentityThroughRPCCalls(ctx)
	ctx = s.AnnotateCtx(ctx)

	username, isAdmin, err := s.privilegeChecker.getUserAndRole(ctx)
	if err != nil {
		return nil, serverError(ctx, err)
	}

	var resp serverpb.UserSQLRolesResponse
	if !isAdmin {
		for _, privKind := range privilege.GlobalPrivileges {
			privName := privKind.String()
			hasPriv, err := s.privilegeChecker.hasGlobalPrivilege(ctx, username, privKind)
			if err != nil {
				return nil, serverError(ctx, err)
			}
			if hasPriv {
				resp.Roles = append(resp.Roles, privName)
				continue
			}
			roleOpt, ok := roleoption.ByName[privName]
			if !ok {
				continue
			}
			hasRole, err := s.privilegeChecker.hasRoleOption(ctx, username, roleOpt)
			if err != nil {
				return nil, serverError(ctx, err)
			}
			if hasRole {
				resp.Roles = append(resp.Roles, privName)
			}
		}
	} else {
		resp.Roles = append(resp.Roles, "ADMIN")
	}
	return &resp, nil
}
