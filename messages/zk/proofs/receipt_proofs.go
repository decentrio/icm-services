// Copyright (C) 2026, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

// THIS IS AN EXAMPLE OF UNAUDITED CODE. DO NOT USE THIS IN PRODUCTION.

package proofs

import (
	"context"
	"fmt"

	"github.com/ava-labs/libevm/common"
	"github.com/ava-labs/libevm/common/hexutil"
	"github.com/ava-labs/libevm/core/rawdb"
	"github.com/ava-labs/libevm/rlp"
	"github.com/ava-labs/libevm/trie"
	"github.com/ava-labs/libevm/triedb"
)

// ethClient is the subset of rpc.Client used for proof building. Results are
// decoded from raw JSON here rather than into libevm's types, which stop at
// the Cancun fork: libevm's header hash is wrong for Prague and later
// headers, and it encodes type-4 (EIP-7702) receipts as just their type byte.
type ethClient interface {
	CallContext(ctx context.Context, result any, method string, args ...any) error
}

// ReceiptProof is a Merkle Patricia Tree (MPT) inclusion proof for
// a specific log inside a receipt and is against a block's receipts root.
// The field layout below mirrors the ZKAdapter's Receipt.Proof struct
// See: https://github.com/ava-labs/icm-services/blob/39eb811943f8017843e30dfeb059e829fe65099e/icm-contracts/avalanche/verifiers/ethereum/StateManagerLibrary.sol#L160
//
//nolint:lll
type ReceiptProof struct {
	// Proof is the list of trie nodes on the path from the receipts root to
	// the receipt, in root-to-leaf order.
	Proof [][]byte
	// Key is the RLP-encoded transaction index of the receipt in the block.
	Key []byte
	// Value is the RLP-encoded (EIP-2718 typed) receipt data.
	Value []byte
	// LogIndex of the message event within the receipt.
	LogIndex uint
	// ExpectedEmitter is the contract that emitted the log.
	ExpectedEmitter common.Address
	// ExpectedTopic0 is the log's topic0 (event signature hash).
	ExpectedTopic0 common.Hash
}

// rpcReceipt holds the fields of an eth_getBlockReceipts entry that the
// receipts trie commits to, plus the transaction it belongs to.
type rpcReceipt struct {
	Type              hexutil.Uint64 `json:"type"`
	Status            hexutil.Uint64 `json:"status"`
	CumulativeGasUsed hexutil.Uint64 `json:"cumulativeGasUsed"`
	LogsBloom         hexutil.Bytes  `json:"logsBloom"`
	Logs              []rpcLog       `json:"logs"`
	TxHash            common.Hash    `json:"transactionHash"`
	TxIndex           hexutil.Uint64 `json:"transactionIndex"`
}

// rpcLog holds the fields of a log that the receipts trie commits to. Its
// field order is the log's RLP encoding.
type rpcLog struct {
	Address common.Address `json:"address"`
	Topics  []common.Hash  `json:"topics"`
	Data    hexutil.Bytes  `json:"data"`
}

// BuildReceiptProof reconstructs the receipts trie for the block containing
// txHash, verifies its root against the block header, and extracts the
// inclusion proof for that transaction's receipt and the log at logIndex.
func BuildReceiptProof(
	ctx context.Context,
	client ethClient,
	blockNumber uint64,
	txHash common.Hash,
	logIndex uint,
) (*ReceiptProof, error) {
	// Get the target receipt corresponding to txHash and its log. Both are
	// fetched by number; a reorg between the two fetches fails the receipts
	// root check below.
	number := hexutil.EncodeUint64(blockNumber)
	var header *struct {
		ReceiptsRoot common.Hash `json:"receiptsRoot"`
	}
	if err := client.CallContext(ctx, &header, "eth_getBlockByNumber", number, false); err != nil {
		return nil, fmt.Errorf("failed to fetch header %d: %w", blockNumber, err)
	}
	if header == nil {
		return nil, fmt.Errorf("block %d not found", blockNumber)
	}
	var receipts []*rpcReceipt
	if err := client.CallContext(ctx, &receipts, "eth_getBlockReceipts", number); err != nil {
		return nil, fmt.Errorf("failed to fetch receipts for block %d: %w", blockNumber, err)
	}
	targetReceipt, err := receiptByTxHash(receipts, txHash)
	if err != nil {
		return nil, err
	}
	if int(logIndex) >= len(targetReceipt.Logs) {
		return nil, fmt.Errorf(
			"log index %d out of range (total %d logs) in receipt %s",
			logIndex,
			len(targetReceipt.Logs),
			txHash.Hex(),
		)
	}
	targetLog := targetReceipt.Logs[logIndex]

	// Build the receipts trie and verify its root against the header receipts root.
	receiptsTrie, targetValue, err := buildReceiptsTrie(
		receipts,
		uint64(targetReceipt.TxIndex),
		header.ReceiptsRoot,
	)
	if err != nil {
		return nil, err
	}

	// Compute the inclusion proof for the target receipt.
	key, err := rlp.EncodeToBytes(uint64(targetReceipt.TxIndex))
	if err != nil {
		return nil, fmt.Errorf("failed to encode target trie key: %w", err)
	}
	var proof proofList
	if err := receiptsTrie.Prove(key, &proof); err != nil {
		return nil, fmt.Errorf("failed to generate receipt proof: %w", err)
	}

	return &ReceiptProof{
		Proof:           proof,
		Key:             key,
		Value:           targetValue,
		LogIndex:        logIndex,
		ExpectedEmitter: targetLog.Address,
		ExpectedTopic0:  targetLog.Topics[0],
	}, nil
}

// buildReceiptsTrie reconstructs the block's receipts trie and returns it
// along with the encoded target receipt. The trie root is verified against
// the header's receipts root, so any encoding drift fails here rather than
// producing proofs against the wrong tree.
func buildReceiptsTrie(
	receipts []*rpcReceipt,
	targetIndex uint64,
	receiptsRoot common.Hash,
) (*trie.Trie, []byte, error) {
	receiptsTrie := trie.NewEmpty(triedb.NewDatabase(rawdb.NewMemoryDatabase(), nil))
	var targetValue []byte
	for i, receipt := range receipts {
		key, err := rlp.EncodeToBytes(uint64(i))
		if err != nil {
			return nil, nil, fmt.Errorf("failed to encode trie key %d: %w", i, err)
		}
		value, err := encodeReceipt(receipt)
		if err != nil {
			return nil, nil, fmt.Errorf("failed to encode receipt %d: %w", i, err)
		}
		if err := receiptsTrie.Update(key, value); err != nil {
			return nil, nil, fmt.Errorf("failed to insert receipt %d into trie: %w", i, err)
		}
		if uint64(i) == targetIndex {
			targetValue = value
		}
	}

	if receiptsTrie.Hash() != receiptsRoot {
		return nil, nil, fmt.Errorf(
			"receipts trie root mismatch: computed %s, header %s",
			receiptsTrie.Hash().Hex(),
			receiptsRoot.Hex(),
		)
	}
	return receiptsTrie, targetValue, nil
}

// encodeReceipt returns the receipt's consensus encoding, the value the
// receipts trie stores: rlp([status, cumulativeGasUsed, logsBloom, logs]),
// prefixed by the type byte for typed (EIP-2718) receipts. The body is the
// same for every receipt type since Byzantium.
func encodeReceipt(r *rpcReceipt) ([]byte, error) {
	body, err := rlp.EncodeToBytes([]any{
		uint64(r.Status),
		uint64(r.CumulativeGasUsed),
		[]byte(r.LogsBloom),
		r.Logs,
	})
	if err != nil {
		return nil, err
	}
	if r.Type == 0 {
		return body, nil
	}
	return append([]byte{byte(r.Type)}, body...), nil
}

func receiptByTxHash(receipts []*rpcReceipt, txHash common.Hash) (*rpcReceipt, error) {
	for _, r := range receipts {
		if r.TxHash == txHash {
			return r, nil
		}
	}
	return nil, fmt.Errorf("receipt for tx %s not found in block", txHash.Hex())
}

// proofList collects trie proof nodes in insertion (root-to-leaf) order.
type proofList [][]byte

func (p *proofList) Put(_ []byte, value []byte) error {
	*p = append(*p, value)
	return nil
}

func (p *proofList) Delete([]byte) error { return nil }
