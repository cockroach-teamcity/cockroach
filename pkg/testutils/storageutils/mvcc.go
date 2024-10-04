// Copyright 2022 The Cockroach Authors.
//
// Use of this software is governed by the CockroachDB Software License
// included in the /LICENSE file.

package storageutils

import (
	"testing"

	"github.com/cockroachdb/cockroach/pkg/storage"
	"github.com/stretchr/testify/require"
)

// MVCCGetRaw fetches a raw MVCC value, for use in tests.
func MVCCGetRaw(t *testing.T, r storage.Reader, key storage.MVCCKey) []byte {
	value, err := MVCCGetRawWithError(t, r, key)
	require.NoError(t, err)
	return value
}

// MVCCGetRawWithError is like MVCCGetRaw, but returns an error rather than
// failing the test.
func MVCCGetRawWithError(t *testing.T, r storage.Reader, key storage.MVCCKey) ([]byte, error) {
	iter := r.NewMVCCIterator(storage.MVCCKeyAndIntentsIterKind, storage.IterOptions{Prefix: true})
	defer iter.Close()
	iter.SeekGE(key)
	if ok, err := iter.Valid(); err != nil || !ok {
		return nil, err
	}
	return iter.Value()
}
