// Copyright 2020 The Cockroach Authors.
//
// Use of this software is governed by the CockroachDB Software License
// included in the /LICENSE file.

// Package kvtenant provides utilities required by SQL-only tenant processes in
// order to interact with the key-value layer.
package kvtenant

import (
	"github.com/cockroachdb/cockroach/pkg/config/zonepb"
	"github.com/cockroachdb/cockroach/pkg/roachpb"
	"github.com/cockroachdb/cockroach/pkg/rpc"
	"github.com/cockroachdb/cockroach/pkg/util/log"
	"github.com/cockroachdb/cockroach/pkg/util/retry"
)

// Factory is a hook for binaries that include CCL code to inject a
// ConnectorFactory.
var Factory ConnectorFactory = connectorFactory{}

// ConnectorConfig encompasses the configuration required to create a Connector.
type ConnectorConfig struct {
	TenantID          roachpb.TenantID
	AmbientCtx        log.AmbientContext
	RPCContext        *rpc.Context
	RPCRetryOptions   retry.Options
	DefaultZoneConfig *zonepb.ZoneConfig
}

// KVAddressConfig encompasses the network addresses, pointing to KV nodes,
// required to create a Connector.
type KVAddressConfig struct {
	RemoteAddresses []string
	LoopbackAddress string
}

// ConnectorFactory constructs a new tenant Connector from the provided network
// addresses pointing to KV nodes.
type ConnectorFactory interface {
	NewConnector(cfg ConnectorConfig, addressConfig KVAddressConfig) (Connector, error)
}
