// Copyright (C) 2026, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package proofs

import (
	"context"
	"fmt"
	"math/big"
	"testing"

	"github.com/ava-labs/avalanchego/graft/evm/rpc"
	"github.com/ava-labs/libevm/common"
	"github.com/ava-labs/libevm/core/rawdb"
	"github.com/ava-labs/libevm/core/types"
	"github.com/ava-labs/libevm/crypto"
	"github.com/ava-labs/libevm/rlp"
	"github.com/ava-labs/libevm/trie"
	"github.com/stretchr/testify/require"
)

// mockEthClient returns fixed test data instead of querying a real chain.
type mockEthClient struct {
	header   *types.Header
	receipts types.Receipts
}

func (f *mockEthClient) HeaderByNumber(_ context.Context, _ *big.Int) (*types.Header, error) {
	return f.header, nil
}

func (f *mockEthClient) BlockReceipts(_ context.Context, _ rpc.BlockNumberOrHash) ([]*types.Receipt, error) {
	return f.receipts, nil
}

// makeTestBlock builds n synthetic receipts and a header carrying their
// correct receipts root.
func makeTestBlock(t *testing.T, n int) (*mockEthClient, types.Receipts) {
	t.Helper()

	receipts := make(types.Receipts, n)
	for i := 0; i < n; i++ {
		receipts[i] = &types.Receipt{
			Type:              types.LegacyTxType,
			Status:            types.ReceiptStatusSuccessful,
			CumulativeGasUsed: uint64(21000 * (i + 1)),
			Logs: []*types.Log{
				{
					Address: common.BytesToAddress([]byte{0xEE, byte(i)}),
					Topics:  []common.Hash{common.BytesToHash([]byte{0xAA, byte(i)})},
					Data:    []byte{byte(i)},
				},
				{
					Address: common.BytesToAddress([]byte{0xFF, byte(i)}),
					Topics:  []common.Hash{common.BytesToHash([]byte{0xBB, byte(i)})},
					Data:    []byte{byte(i), byte(i)},
				},
			},
			TxHash:           common.BytesToHash([]byte{0xCC, byte(i)}),
			TransactionIndex: uint(i),
		}
	}

	header := &types.Header{
		Number:      big.NewInt(100),
		ReceiptHash: types.DeriveSha(receipts, trie.NewStackTrie(nil)),
	}

	return &mockEthClient{header: header, receipts: receipts}, receipts
}

// verifyProofAgainstRoot cryptographically verifies the returned proof nodes
// against the receipts root and returns the proven value. Plays the role of
// on-chain verifier.
func verifyProofAgainstRoot(t *testing.T, root common.Hash, key []byte, nodes [][]byte) []byte {
	t.Helper()
	proofDB := rawdb.NewMemoryDatabase()
	for _, node := range nodes {
		require.NoError(t, proofDB.Put(crypto.Keccak256(node), node))
	}
	value, err := trie.VerifyProof(root, key, proofDB)
	require.NoError(t, err)
	return value
}

// Happy path: the proof verifies against the header's receipts root and the
// proven value is the canonical encoding of the target receipt.
func TestBuildReceiptProof(t *testing.T) {
	client, receipts := makeTestBlock(t, 3)
	targetIdx := 1
	targetLogIdx := uint(1)
	targetHash := receipts[targetIdx].TxHash

	// Build proof.
	proof, err := BuildReceiptProof(context.Background(), client, 100, targetHash, targetLogIdx)
	require.NoError(t, err)

	expectedKey, err := rlp.EncodeToBytes(uint(targetIdx))
	require.NoError(t, err)
	require.Equal(t, expectedKey, proof.Key)

	// Verify.
	provenValue := verifyProofAgainstRoot(t, client.header.ReceiptHash, proof.Key, proof.Proof)
	require.Equal(t, proof.Value, provenValue)

	// Log metadata must match the requested log.
	targetLog := receipts[targetIdx].Logs[targetLogIdx]
	require.Equal(t, targetLogIdx, proof.LogIndex)
	require.Equal(t, targetLog.Address, proof.ExpectedEmitter)
	require.Equal(t, targetLog.Topics[0], proof.ExpectedTopic0)
}

// Verify every receipt in the block.
func TestBuildReceiptProofAllIndices(t *testing.T) {
	client, receipts := makeTestBlock(t, 5)
	for i, receipt := range receipts {
		proof, err := BuildReceiptProof(context.Background(), client, 100, receipt.TxHash, 0)
		require.NoError(t, err, "receipt %d", i)
		provenValue := verifyProofAgainstRoot(t, client.header.ReceiptHash, proof.Key, proof.Proof)
		require.Equal(t, proof.Value, provenValue, "receipt %d", i)
	}
}

// A log index beyond the receipt's logs must error.
func TestBuildReceiptProofLogIndexOutOfRange(t *testing.T) {
	client, receipts := makeTestBlock(t, 2)
	_, err := BuildReceiptProof(context.Background(), client, 100, receipts[0].TxHash, 99)
	require.ErrorContains(t, err, "log index")
}

// A transaction hash not present in the block must error.
func TestBuildReceiptProofTxNotInBlock(t *testing.T) {
	client, _ := makeTestBlock(t, 2)
	_, err := BuildReceiptProof(context.Background(), client, 100, common.HexToHash("0xdead"), 0)
	require.ErrorContains(t, err, "not found in block")
}

// sepoliaBlockClient serves a real Sepolia block's header and receipts. Like a
// node, it answers receipt requests only for the block's number or its real
// hash.
type sepoliaBlockClient struct {
	header   *types.Header
	receipts []*types.Receipt
	number   uint64
	hash     common.Hash
}

func (c *sepoliaBlockClient) HeaderByNumber(_ context.Context, _ *big.Int) (*types.Header, error) {
	return c.header, nil
}

func (c *sepoliaBlockClient) BlockReceipts(_ context.Context, id rpc.BlockNumberOrHash) ([]*types.Receipt, error) {
	if number, ok := id.Number(); ok && uint64(number) == c.number {
		return c.receipts, nil
	}
	if hash, ok := id.Hash(); ok && hash == c.hash {
		return c.receipts, nil
	}
	return nil, fmt.Errorf("no block %s", id.String())
}

// Receipts must not be fetched by libevm's Header.Hash(): it ignores the
// Prague and Amsterdam header fields, so for a Sepolia Gloas block it names a
// block that does not exist. The proof must match the fixture's, which was
// built by an independent MPT implementation.
func TestBuildReceiptProofSepoliaBlock(t *testing.T) {
	fixture := loadSepoliaFixture(t)
	client := &sepoliaBlockClient{
		number: fixture.Metadata.BlockNumber,
		hash:   fixture.Metadata.BlockHash,
	}
	loadRPCResult(t, sepoliaBlockPath, &client.header)
	loadRPCResult(t, sepoliaReceiptsPath, &client.receipts)

	proof, err := BuildReceiptProof(
		context.Background(), client, fixture.Metadata.BlockNumber, fixture.Metadata.TxHash, 0)
	require.NoError(t, err)

	expected := fixture.ReceiptProof
	require.Equal(t, []byte(expected.Key), proof.Key)
	require.Equal(t, []byte(expected.Value), proof.Value)
	require.Len(t, proof.Proof, len(expected.Proof))
	for i := range expected.Proof {
		require.Equal(t, []byte(expected.Proof[i]), proof.Proof[i], "proof node %d", i)
	}
	require.Equal(t, expected.LogIndex, proof.LogIndex)
	require.Equal(t, expected.ExpectedEmitter, proof.ExpectedEmitter)
	require.Equal(t, expected.ExpectedTopic0, proof.ExpectedTopic0)
}

// A header whose receipts root doesn't match the block's receipts must error
// rather than produce a proof against the wrong tree.
func TestBuildReceiptProofRootMismatch(t *testing.T) {
	client, receipts := makeTestBlock(t, 2)
	client.header.ReceiptHash = common.HexToHash("0xbadbadbad")
	_, err := BuildReceiptProof(context.Background(), client, 100, receipts[0].TxHash, 0)
	require.ErrorContains(t, err, "root mismatch")
}
