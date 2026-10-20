// Copyright (C) 2026, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package proofs

import (
	"encoding/json"
	"testing"

	"github.com/ava-labs/libevm/common"
	"github.com/ava-labs/libevm/crypto"
	"github.com/stretchr/testify/require"
)

// mainnetBlockPath is a mainnet Prague/Osaka block (21 header fields), as
// returned by eth_getBlockByNumber.
const mainnetBlockPath = "testdata/mainnet/el_block_25575612.json"

// loadHeaderFields reads a saved block response as its JSON fields.
func loadHeaderFields(t *testing.T, path string) map[string]json.RawMessage {
	t.Helper()
	var fields map[string]json.RawMessage
	loadRPCResult(t, path, &fields)
	return fields
}

func marshalHeaderFields(t *testing.T, fields map[string]json.RawMessage) []byte {
	t.Helper()
	data, err := json.Marshal(fields)
	require.NoError(t, err)
	return data
}

// The encoding must hash to the block hash the node reports, which is the
// definition of the block hash, for headers of different forks: a mainnet
// Prague header (21 fields) and a Sepolia Amsterdam header (23 fields).
func TestEncodeExecutionHeaderMatchesBlockHash(t *testing.T) {
	for _, path := range []string{mainnetBlockPath, sepoliaBlockPath} {
		fields := loadHeaderFields(t, path)
		var hash common.Hash
		require.NoError(t, json.Unmarshal(fields["hash"], &hash))

		headerRLP, err := encodeExecutionHeader(marshalHeaderFields(t, fields))
		require.NoError(t, err, path)
		require.Equal(t, hash, crypto.Keccak256Hash(headerRLP), path)
	}

	// The Sepolia encoding must also equal the fixture's, built by an
	// independent encoder.
	fixture := loadSepoliaFixture(t)
	headerRLP, err := encodeExecutionHeader(marshalHeaderFields(t, loadHeaderFields(t, sepoliaBlockPath)))
	require.NoError(t, err)
	require.Equal(t, []byte(fixture.ExecutionProof.ExecutionHeaderRlp), headerRLP)
}

// Each fork appends header fields, so a header carries a prefix of the
// optional fields. A later field present after a missing one cannot be
// encoded at its position and must be rejected.
func TestEncodeExecutionHeaderRejectsGapInOptionalFields(t *testing.T) {
	fields := loadHeaderFields(t, sepoliaBlockPath)
	delete(fields, "withdrawalsRoot")

	_, err := encodeExecutionHeader(marshalHeaderFields(t, fields))
	require.ErrorContains(t, err, `header field "blobGasUsed" follows missing field "withdrawalsRoot"`)
}

// Every header has the 15 legacy fields.
func TestEncodeExecutionHeaderRejectsMissingLegacyField(t *testing.T) {
	fields := loadHeaderFields(t, sepoliaBlockPath)
	delete(fields, "mixHash")

	_, err := encodeExecutionHeader(marshalHeaderFields(t, fields))
	require.ErrorContains(t, err, `missing header field "mixHash"`)
}
