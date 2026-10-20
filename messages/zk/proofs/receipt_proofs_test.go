// Copyright (C) 2026, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package proofs

import (
	"context"
	"encoding/json"
	"fmt"
	"math/big"
	"testing"

	"github.com/ava-labs/libevm/common"
	"github.com/ava-labs/libevm/common/hexutil"
	"github.com/ava-labs/libevm/core/rawdb"
	"github.com/ava-labs/libevm/core/types"
	"github.com/ava-labs/libevm/crypto"
	"github.com/ava-labs/libevm/rlp"
	"github.com/ava-labs/libevm/trie"
	"github.com/stretchr/testify/require"
)

// mockEthClient returns fixed test data instead of querying a real chain,
// serialized to JSON the way a node returns it.
type mockEthClient struct {
	header   *types.Header
	receipts types.Receipts
}

func (f *mockEthClient) CallContext(_ context.Context, result any, method string, _ ...any) error {
	switch method {
	case "eth_getBlockByNumber":
		return jsonRoundTrip(f.header, result)
	case "eth_getBlockReceipts":
		return jsonRoundTrip(f.receipts, result)
	default:
		return fmt.Errorf("unexpected method %s", method)
	}
}

// jsonRoundTrip decodes value's JSON encoding into result.
func jsonRoundTrip(value any, result any) error {
	data, err := json.Marshal(value)
	if err != nil {
		return err
	}
	return json.Unmarshal(data, result)
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

// A block the node does not know, which it reports as a null result, must
// error.
func TestBuildReceiptProofBlockNotFound(t *testing.T) {
	client, receipts := makeTestBlock(t, 2)
	client.header = nil
	_, err := BuildReceiptProof(context.Background(), client, 100, receipts[0].TxHash, 0)
	require.ErrorContains(t, err, "block 100 not found")
}

// rpcFileClient serves a block and its receipts from saved JSON-RPC
// responses. Like a node, it answers only for the block's number or its real
// hash.
type rpcFileClient struct {
	block        json.RawMessage
	receipts     json.RawMessage
	number       uint64
	hash         common.Hash
	receiptsRoot common.Hash
}

func newRPCFileClient(t *testing.T, blockPath, receiptsPath string) *rpcFileClient {
	t.Helper()
	var header struct {
		Number       hexutil.Uint64 `json:"number"`
		Hash         common.Hash    `json:"hash"`
		ReceiptsRoot common.Hash    `json:"receiptsRoot"`
	}
	loadRPCResult(t, blockPath, &header)
	client := &rpcFileClient{
		number:       uint64(header.Number),
		hash:         header.Hash,
		receiptsRoot: header.ReceiptsRoot,
	}
	loadRPCResult(t, blockPath, &client.block)
	loadRPCResult(t, receiptsPath, &client.receipts)
	return client
}

func (c *rpcFileClient) CallContext(_ context.Context, result any, method string, args ...any) error {
	if len(args) == 0 {
		return fmt.Errorf("%s: missing block argument", method)
	}
	id := args[0]
	byNumber := id == hexutil.EncodeUint64(c.number)
	byHash := id == c.hash || id == c.hash.Hex()
	switch {
	case method == "eth_getBlockByNumber" && byNumber:
		return json.Unmarshal(c.block, result)
	case method == "eth_getBlockReceipts" && (byNumber || byHash):
		return json.Unmarshal(c.receipts, result)
	default:
		return fmt.Errorf("%s: no block %v", method, id)
	}
}

// Receipts must not be fetched by libevm's Header.Hash(): it ignores the
// Prague and Amsterdam header fields, so for a Sepolia Gloas block it names a
// block that does not exist. The proof must match the fixture's, which was
// built by an independent MPT implementation.
func TestBuildReceiptProofSepoliaBlock(t *testing.T) {
	fixture := loadSepoliaFixture(t)
	client := newRPCFileClient(t, sepoliaBlockPath, sepoliaReceiptsPath)

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

// A block holding a type-4 (EIP-7702) receipt must reproduce the header's
// receipts root. libevm's receipt encoding writes only the type byte for
// type-4 receipts, so its trie misses the header's root.
func TestBuildReceiptProofSetCodeReceipt(t *testing.T) {
	const (
		blockPath    = "testdata/sepolia/el_block_11955870.json"
		receiptsPath = "testdata/sepolia/el_receipts_11955870.json"
	)
	// The block's only type-4 receipt, at transaction index 7.
	setCodeTx := common.HexToHash("0xe46a1e7377bafb5229b1b852d696f976b6945b57ff28a08e6c4c20bc7970a78c")
	client := newRPCFileClient(t, blockPath, receiptsPath)

	proof, err := BuildReceiptProof(context.Background(), client, client.number, setCodeTx, 0)
	require.NoError(t, err)

	provenValue := verifyProofAgainstRoot(t, client.receiptsRoot, proof.Key, proof.Proof)
	require.Equal(t, proof.Value, provenValue)
	require.Equal(t, byte(0x04), proof.Value[0], "EIP-7702 receipt type")
}

// A header whose receipts root doesn't match the block's receipts must error
// rather than produce a proof against the wrong tree.
func TestBuildReceiptProofRootMismatch(t *testing.T) {
	client, receipts := makeTestBlock(t, 2)
	client.header.ReceiptHash = common.HexToHash("0xbadbadbad")
	_, err := BuildReceiptProof(context.Background(), client, 100, receipts[0].TxHash, 0)
	require.ErrorContains(t, err, "root mismatch")
}
