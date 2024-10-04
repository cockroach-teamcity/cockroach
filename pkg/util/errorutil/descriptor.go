// Copyright 2019 The Cockroach Authors.
//
// Use of this software is governed by the CockroachDB Software License
// included in the /LICENSE file.

package errorutil

import (
	"fmt"

	"github.com/cockroachdb/cockroach/pkg/roachpb"
	"github.com/cockroachdb/errors"
)

type descriptorNotFound struct {
	msg string
}

func (e *descriptorNotFound) Error() string {
	return e.msg
}

func IsDescriptorNotFoundError(err error) bool {
	return errors.HasType(err, (*descriptorNotFound)(nil))
}

func NewNodeNotFoundError(nodeID roachpb.NodeID) error {
	return &descriptorNotFound{fmt.Sprintf("unable to look up descriptor for n%d", nodeID)}
}

func NewStoreNotFoundError(storeID roachpb.StoreID) error {
	return &descriptorNotFound{fmt.Sprintf("unable to look up descriptor for store ID %d", storeID)}
}
