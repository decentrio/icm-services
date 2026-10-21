// Copyright (C) 2026, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

// THIS IS AN EXAMPLE OF UNAUDITED CODE. DO NOT USE THIS IN PRODUCTION.

// Package lightclient delivers TeleporterMessengerV2 messages from Ethereum to an Avalanche chain,
// proven against the finalized beacon roots of a ZK light client. It finds the message a
// LightClientAdapter emitted on Ethereum, builds the execution and receipt proofs of its log, and
// packs them as the attestation LightClientAdapter.verifyMessage takes.
package lightclient

import (
	"context"
	"errors"
	"fmt"
	"math/big"
	"strings"

	teleportermessengerv2 "github.com/ava-labs/icm-services/abi-bindings/go/TeleporterMessengerV2"
	lightclientadapter "github.com/ava-labs/icm-services/abi-bindings/go/verifiers/ethereum/LightClientAdapter"
	zkadapter "github.com/ava-labs/icm-services/abi-bindings/go/verifiers/ethereum/ZKAdapter"
	"github.com/ava-labs/icm-services/messages/zk/proofs"
	"github.com/ava-labs/libevm/accounts/abi"
	"github.com/ava-labs/libevm/accounts/abi/bind"
	"github.com/ava-labs/libevm/common"
	"github.com/ava-labs/libevm/common/hexutil"
	"github.com/ava-labs/libevm/crypto"
)

var (
	// ErrNoMessage means the transaction emitted no message from the adapter.
	ErrNoMessage = errors.New("transaction sent no message through the adapter")
	// ErrNotCovered means the light client's latest finalized root does not cover the message's
	// block yet: try again after it finalizes a later checkpoint.
	ErrNotCovered = errors.New("light client does not cover the message's block yet")

	// messageSentTopic is the topic of EthereumAdapter's TeleporterV2MessageSent(bytes).
	messageSentTopic = crypto.Keccak256Hash([]byte("TeleporterV2MessageSent(bytes)"))
)

// ethClient is the subset of rpc.Client used to read the source chain.
type ethClient interface {
	CallContext(ctx context.Context, result any, method string, args ...any) error
}

// SentMessage is a message an adapter emitted on the source chain, and where its log is.
type SentMessage struct {
	Message     teleportermessengerv2.TeleporterMessageV2
	TxHash      common.Hash
	BlockNumber uint64
	BlockHash   common.Hash
	// ReceiptLogIndex is the log's position in its receipt, as Receipt.Proof takes it.
	ReceiptLogIndex uint
}

// FindSentMessage returns the first message that adapter emitted in transaction txHash.
func FindSentMessage(
	ctx context.Context,
	el ethClient,
	txHash common.Hash,
	adapter common.Address,
) (*SentMessage, error) {
	var receipt *struct {
		BlockNumber hexutil.Uint64 `json:"blockNumber"`
		BlockHash   common.Hash    `json:"blockHash"`
		Logs        []struct {
			Address common.Address `json:"address"`
			Topics  []common.Hash  `json:"topics"`
			Data    hexutil.Bytes  `json:"data"`
		} `json:"logs"`
	}
	if err := el.CallContext(ctx, &receipt, "eth_getTransactionReceipt", txHash); err != nil {
		return nil, fmt.Errorf("failed to fetch receipt of %s: %w", txHash, err)
	}
	if receipt == nil {
		return nil, fmt.Errorf("transaction %s not found", txHash)
	}
	for i, l := range receipt.Logs {
		if l.Address != adapter || len(l.Topics) == 0 || l.Topics[0] != messageSentTopic {
			continue
		}
		message, err := decodeMessageSent(l.Data)
		if err != nil {
			return nil, err
		}
		return &SentMessage{
			Message:         message,
			TxHash:          txHash,
			BlockNumber:     uint64(receipt.BlockNumber),
			BlockHash:       receipt.BlockHash,
			ReceiptLogIndex: uint(i),
		}, nil
	}
	return nil, fmt.Errorf("%w: %s from %s", ErrNoMessage, txHash, adapter)
}

// decodeMessageSent decodes TeleporterV2MessageSent's data: abi.encode(bytes), where the bytes are
// abi.encode(message).
func decodeMessageSent(data []byte) (teleportermessengerv2.TeleporterMessageV2, error) {
	bytesType, err := abi.NewType("bytes", "", nil)
	if err != nil {
		return teleportermessengerv2.TeleporterMessageV2{}, err
	}
	values, err := abi.Arguments{{Type: bytesType}}.Unpack(data)
	if err != nil {
		return teleportermessengerv2.TeleporterMessageV2{}, fmt.Errorf("failed to decode message log: %w", err)
	}
	var message teleportermessengerv2.TeleporterMessageV2
	if err := message.Unpack(values[0].([]byte)); err != nil {
		return teleportermessengerv2.TeleporterMessageV2{}, err
	}
	return message, nil
}

// EncodeAttestation returns abi.encode(Attestation{execProof, logProof}), the attestation
// EthereumAdapter.verifyMessage decodes.
func EncodeAttestation(execProof zkadapter.ExecutionProof, logProof proofs.ReceiptProof) ([]byte, error) {
	parsed, err := lightclientadapter.LightClientAdapterMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	// verifyLogAndExtract takes the attestation's two fields as its arguments, so its arguments'
	// encoding is the encoding of the Attestation tuple.
	tuple, err := parsed.Methods["verifyLogAndExtract"].Inputs.Pack(
		lightclientadapter.ExecutionProof(execProof),
		lightclientadapter.ReceiptProof{
			Proof:           logProof.Proof,
			Key:             logProof.Key,
			Value:           logProof.Value,
			LogIndex:        new(big.Int).SetUint64(uint64(logProof.LogIndex)),
			ExpectedEmitter: logProof.ExpectedEmitter,
			ExpectedTopic0:  logProof.ExpectedTopic0,
		},
	)
	if err != nil {
		return nil, fmt.Errorf("failed to encode attestation: %w", err)
	}
	// abi.encode of one dynamic tuple is the offset of its encoding, then the encoding.
	return append(common.LeftPadBytes([]byte{0x20}, 32), tuple...), nil
}

// ICMMessage is the message as TeleporterMessengerV2.receiveCrossChainMessage takes it from the
// source chain with EVM chain ID sourceChainID, with its attestation.
func ICMMessage(
	sent *SentMessage,
	sourceChainID uint64,
	attestation []byte,
) (teleportermessengerv2.TeleporterICMMessage, error) {
	networkID, err := networkID(sourceChainID)
	if err != nil {
		return teleportermessengerv2.TeleporterICMMessage{}, err
	}
	return teleportermessengerv2.TeleporterICMMessage{
		Message:            sent.Message,
		SourceNetworkID:    networkID,
		SourceBlockchainID: common.BigToHash(new(big.Int).SetUint64(sourceChainID)),
		Attestation:        attestation,
	}, nil
}

// networkID is the ICM network ID of an Ethereum chain, as EthereumAdapter maps it.
func networkID(chainID uint64) (uint32, error) {
	switch chainID {
	case 1:
		return 1, nil
	case 11155111:
		return 2, nil
	}
	return 0, fmt.Errorf("unsupported source chain %d", chainID)
}

// RootSource is the light client as seen by the relay: its latest finalized slot and the beacon
// root it stored at a slot.
type RootSource interface {
	LatestFinalizedSlot(ctx context.Context) (uint64, error)
	BeaconRootAt(ctx context.Context, slot uint64) (common.Hash, error)
}

// Attest builds the delivery of sent, from the source chain with EVM chain ID sourceChainID,
// against the light client's latest finalized root. It returns ErrNotCovered while that root does
// not cover the message's block.
func Attest(
	ctx context.Context,
	builder *proofs.ProofBuilder,
	el ethClient,
	roots RootSource,
	sent *SentMessage,
	sourceChainID uint64,
) (teleportermessengerv2.TeleporterICMMessage, error) {
	var block *struct {
		SlotNumber *hexutil.Uint64 `json:"slotNumber"`
	}
	if err := el.CallContext(ctx, &block, "eth_getBlockByHash", sent.BlockHash, false); err != nil {
		return teleportermessengerv2.TeleporterICMMessage{}, err
	}
	if block == nil || block.SlotNumber == nil {
		return teleportermessengerv2.TeleporterICMMessage{}, fmt.Errorf("block %s has no slotNumber", sent.BlockHash)
	}
	payloadSlot := uint64(*block.SlotNumber)

	anchor, err := roots.LatestFinalizedSlot(ctx)
	if err != nil {
		return teleportermessengerv2.TeleporterICMMessage{}, err
	}
	// The payload is committed by a later block, which must be before the anchor.
	if payloadSlot+1 >= anchor {
		return teleportermessengerv2.TeleporterICMMessage{}, ErrNotCovered
	}
	anchorRoot, err := roots.BeaconRootAt(ctx, anchor)
	if err != nil {
		return teleportermessengerv2.TeleporterICMMessage{}, err
	}
	execProof, err := builder.BuildExecutionProofForBlock(ctx, anchor, anchorRoot,
		proofs.ExecutionTarget{BlockHash: sent.BlockHash, PayloadSlot: payloadSlot})
	if errors.Is(err, proofs.ErrNotCommittedBeforeAnchor) {
		return teleportermessengerv2.TeleporterICMMessage{}, ErrNotCovered
	}
	if err != nil {
		return teleportermessengerv2.TeleporterICMMessage{}, err
	}
	logProof, err := proofs.BuildReceiptProof(ctx, el, sent.BlockNumber, sent.TxHash, sent.ReceiptLogIndex)
	if err != nil {
		return teleportermessengerv2.TeleporterICMMessage{}, err
	}
	attestation, err := EncodeAttestation(*execProof, *logProof)
	if err != nil {
		return teleportermessengerv2.TeleporterICMMessage{}, err
	}
	return ICMMessage(sent, sourceChainID, attestation)
}

// lightClientABI is the part of the ZK light client the relay reads.
const lightClientABI = `[
	{"type":"function","name":"latestFinalizedSlot","stateMutability":"view","inputs":[],
	 "outputs":[{"type":"uint64"}]},
	{"type":"function","name":"beaconRootAt","stateMutability":"view",
	 "inputs":[{"type":"uint64"}],"outputs":[{"type":"bytes32"}]}
]`

type lightClientRoots struct{ contract *bind.BoundContract }

// NewLightClientRoots reads the roots of the ZK light client at address.
func NewLightClientRoots(address common.Address, backend bind.ContractCaller) (RootSource, error) {
	parsed, err := abi.JSON(strings.NewReader(lightClientABI))
	if err != nil {
		return nil, err
	}
	return &lightClientRoots{bind.NewBoundContract(address, parsed, backend, nil, nil)}, nil
}

func (r *lightClientRoots) LatestFinalizedSlot(ctx context.Context) (uint64, error) {
	var out []any
	if err := r.contract.Call(&bind.CallOpts{Context: ctx}, &out, "latestFinalizedSlot"); err != nil {
		return 0, err
	}
	return out[0].(uint64), nil
}

func (r *lightClientRoots) BeaconRootAt(ctx context.Context, slot uint64) (common.Hash, error) {
	var out []any
	if err := r.contract.Call(&bind.CallOpts{Context: ctx}, &out, "beaconRootAt", slot); err != nil {
		return common.Hash{}, err
	}
	return common.Hash(out[0].([32]byte)), nil
}
