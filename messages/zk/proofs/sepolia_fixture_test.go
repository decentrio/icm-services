// Copyright (C) 2026, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package proofs

import (
	"encoding/json"
	"os"
	"testing"

	"github.com/ava-labs/libevm/common"
	"github.com/ava-labs/libevm/common/hexutil"
	"github.com/stretchr/testify/require"
)

// Real Sepolia data for a USDC Transfer log in EL block 11,955,865, a Gloas
// payload. The fixture was built and cross-checked by independent tooling
// (a Python SSZ/RLP/MPT implementation, dynamic-ssz and cast), and the EL
// files are raw JSON-RPC responses from a Sepolia node.
const (
	sepoliaFixturePath  = "testdata/sepolia_gloas_fixture.json"
	sepoliaBlockPath    = "testdata/sepolia/el_block_11955865.json"
	sepoliaReceiptsPath = "testdata/sepolia/el_receipts_11955865.json"
)

// sepoliaFixture mirrors the parts of the Sepolia Gloas fixture the tests
// read.
type sepoliaFixture struct {
	Metadata struct {
		TxHash      common.Hash `json:"txHash"`
		BlockNumber uint64      `json:"blockNumber"`
		BlockHash   common.Hash `json:"blockHash"`
	} `json:"metadata"`
	ExecutionProof struct {
		ExecutionHeaderRlp hexutil.Bytes `json:"executionHeaderRlp"`
	} `json:"executionProof"`
	ReceiptProof struct {
		Proof           []hexutil.Bytes `json:"proof"`
		Key             hexutil.Bytes   `json:"key"`
		Value           hexutil.Bytes   `json:"value"`
		LogIndex        uint            `json:"logIndex"`
		ExpectedEmitter common.Address  `json:"expectedEmitter"`
		ExpectedTopic0  common.Hash     `json:"expectedTopic0"`
	} `json:"receiptProof"`
}

func loadSepoliaFixture(t *testing.T) *sepoliaFixture {
	t.Helper()
	data, err := os.ReadFile(sepoliaFixturePath)
	require.NoError(t, err)
	var fixture sepoliaFixture
	require.NoError(t, json.Unmarshal(data, &fixture))
	return &fixture
}

// loadRPCResult reads a saved JSON-RPC response and decodes its result field
// into out.
func loadRPCResult(t *testing.T, path string, out any) {
	t.Helper()
	data, err := os.ReadFile(path)
	require.NoError(t, err)
	var response struct {
		Result json.RawMessage `json:"result"`
	}
	require.NoError(t, json.Unmarshal(data, &response))
	require.NoError(t, json.Unmarshal(response.Result, out))
}
