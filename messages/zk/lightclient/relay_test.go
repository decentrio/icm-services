// Copyright (C) 2026, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package lightclient

import (
	"context"
	"encoding/json"
	"math/big"
	"os"
	"path/filepath"
	"testing"

	zkadapter "github.com/ava-labs/icm-services/abi-bindings/go/verifiers/ethereum/ZKAdapter"
	"github.com/ava-labs/icm-services/messages/zk/proofs"
	"github.com/ava-labs/libevm/common"
	"github.com/ava-labs/libevm/common/hexutil"
	"github.com/stretchr/testify/require"
)

// fakeEL serves eth_getTransactionReceipt from a captured receipt.
type fakeEL struct{ receipt json.RawMessage }

func (f fakeEL) CallContext(_ context.Context, result any, method string, _ ...any) error {
	if method != "eth_getTransactionReceipt" {
		return os.ErrNotExist
	}
	return json.Unmarshal(f.receipt, result)
}

func readJSON(t *testing.T, path string, v any) {
	t.Helper()
	b, err := os.ReadFile(path)
	require.NoError(t, err)
	require.NoError(t, json.Unmarshal(b, v))
}

// The receipt of a TeleporterV2MessageSent log that a LightClientAdapter emitted on a Sepolia fork,
// for this message.
var (
	forkAdapter = common.HexToAddress("0xcc870FFaCAed7CF9Ab4030E268dcEeb4E182eE9C")
	forkTx      = common.HexToHash("0xae1f9f5bd8487110ca6cccb6e2a0df1380e7f69ecdec4503d10b9d5d1aa684ce")
	forkSender  = common.HexToAddress("0x3C44CdDdB6a900fa2b585dd299e03d12FA4293BC")
)

func loadForkReceipt(t *testing.T) fakeEL {
	t.Helper()
	b, err := os.ReadFile(filepath.Join("testdata", "sent_message_receipt.json"))
	require.NoError(t, err)
	return fakeEL{receipt: b}
}

func TestFindSentMessageDecodesTheAdaptersLog(t *testing.T) {
	sent, err := FindSentMessage(context.Background(), loadForkReceipt(t), forkTx, forkAdapter)
	require.NoError(t, err)

	m := sent.Message
	require.Equal(t, big.NewInt(7), m.MessageNonce)
	require.Equal(t, common.HexToAddress("0xaa"), m.OriginSenderAddress)
	require.Equal(t, forkSender, m.OriginTeleporterAddress)
	require.Equal(t, common.HexToHash("0x7fc93d85c6d62c5b2ac0b519c87010ea5294012d1e407030d6acd0021cac10d5"),
		common.Hash(m.DestinationBlockchainID))
	require.Equal(t, common.HexToAddress("0xbb"), m.DestinationAddress)
	require.Equal(t, big.NewInt(250000), m.RequiredGasLimit)
	require.Equal(t, []byte{0xde, 0xad, 0xbe, 0xef}, m.Message)

	var receipt struct {
		BlockNumber hexutil.Uint64 `json:"blockNumber"`
		BlockHash   common.Hash    `json:"blockHash"`
	}
	readJSON(t, filepath.Join("testdata", "sent_message_receipt.json"), &receipt)
	require.Equal(t, uint64(receipt.BlockNumber), sent.BlockNumber)
	require.Equal(t, receipt.BlockHash, sent.BlockHash)
	require.Equal(t, uint(0), sent.ReceiptLogIndex)
}

// Only the adapter's own logs are messages: verifyMessage on the destination accepts no other emitter.
func TestFindSentMessageIgnoresLogsOfOtherContracts(t *testing.T) {
	_, err := FindSentMessage(context.Background(), loadForkReceipt(t), forkTx, common.HexToAddress("0x1"))
	require.ErrorIs(t, err, ErrNoMessage)
}

// The attestation is abi.encode(Attestation{execProof, logProof}) as EthereumAdapter.verifyMessage
// decodes it; the expected bytes were encoded by Foundry's cast from the same proofs.
func TestEncodeAttestationMatchesSolidityEncoding(t *testing.T) {
	var f struct {
		ExecutionProof struct {
			AnchorSlot              uint64        `json:"anchorSlot"`
			TargetSlot              uint64        `json:"targetSlot"`
			AnchorBeaconStateRoot   common.Hash   `json:"anchorBeaconStateRoot"`
			AnchorBeaconStateProof  []common.Hash `json:"anchorBeaconStateProof"`
			TargetBeaconStateRoot   common.Hash   `json:"targetBeaconStateRoot"`
			TargetBeaconStateProof  []common.Hash `json:"targetBeaconStateProof"`
			ExecutionBlockHash      common.Hash   `json:"executionBlockHash"`
			ExecutionBlockHashProof []common.Hash `json:"executionBlockHashProof"`
			ReceiptsRoot            common.Hash   `json:"receiptsRoot"`
			ExecutionHeaderRlp      hexutil.Bytes `json:"executionHeaderRlp"`
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
	readJSON(t, filepath.Join("..", "proofs", "testdata", "sepolia_gloas_fixture.json"), &f)
	var expected struct {
		Attestation hexutil.Bytes `json:"attestation"`
	}
	readJSON(t, filepath.Join("testdata", "sepolia_attestation.json"), &expected)

	e := f.ExecutionProof
	exec := zkadapter.ExecutionProof{
		AnchorSlot:                 e.AnchorSlot,
		TargetSlot:                 e.TargetSlot,
		AnchorBeaconStateRoot:      e.AnchorBeaconStateRoot,
		AnchorBeaconStateProof:     hashes(e.AnchorBeaconStateProof),
		TargetBeaconStateRoot:      e.TargetBeaconStateRoot,
		TargetBeaconStateProof:     hashes(e.TargetBeaconStateProof),
		TargetExecutionHeaderRoot:  e.ExecutionBlockHash,
		TargetExecutionHeaderProof: hashes(e.ExecutionBlockHashProof),
		TargetReceiptsRoot:         e.ReceiptsRoot,
		TargetReceiptsProof:        [][32]byte{},
		ExecutionHeaderRlp:         e.ExecutionHeaderRlp,
	}
	r := f.ReceiptProof
	var nodes [][]byte
	for _, n := range r.Proof {
		nodes = append(nodes, n)
	}
	log := proofs.ReceiptProof{
		Proof: nodes, Key: r.Key, Value: r.Value, LogIndex: r.LogIndex,
		ExpectedEmitter: r.ExpectedEmitter, ExpectedTopic0: r.ExpectedTopic0,
	}

	got, err := EncodeAttestation(exec, log)
	require.NoError(t, err)
	require.Equal(t, []byte(expected.Attestation), got)
}

// EthereumAdapter.verifyMessage requires the source chain's ID as the blockchain ID and its ICM
// network ID: 1 for mainnet, 2 for Sepolia.
func TestICMMessageNamesTheSourceChainAsEthereumAdapterExpects(t *testing.T) {
	sent, err := FindSentMessage(context.Background(), loadForkReceipt(t), forkTx, forkAdapter)
	require.NoError(t, err)

	msg, err := ICMMessage(sent, 11155111, []byte{1})
	require.NoError(t, err)
	require.Equal(t, uint32(2), msg.SourceNetworkID)
	require.Equal(t, common.BigToHash(big.NewInt(11155111)), common.Hash(msg.SourceBlockchainID))
	require.Equal(t, sent.Message, msg.Message)
	require.Equal(t, []byte{1}, msg.Attestation)

	msg, err = ICMMessage(sent, 1, nil)
	require.NoError(t, err)
	require.Equal(t, uint32(1), msg.SourceNetworkID)

	_, err = ICMMessage(sent, 17000, nil)
	require.Error(t, err)
}

func hashes(hs []common.Hash) [][32]byte {
	out := make([][32]byte, len(hs))
	for i, h := range hs {
		out[i] = h
	}
	return out
}
