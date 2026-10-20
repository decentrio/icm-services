// Copyright (C) 2026, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

// THIS IS AN EXAMPLE OF UNAUDITED CODE. DO NOT USE THIS IN PRODUCTION.

package proofs

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/ava-labs/avalanchego/cache/lru"
	zkadapter "github.com/ava-labs/icm-services/abi-bindings/go/verifiers/ethereum/ZKAdapter"
	"github.com/ava-labs/icm-services/messages/zk/proofs/beacon"
	"github.com/ava-labs/libevm/common"
)

// Overview: orchestrator.go connects the three main individual components of the proofs package.
// 1) the beacon client, which fetches SSZ-encoded beacon blocks and states by slot
// 2) the lite beacon tree builder, which parses the SSZ into lite (memory-efficient) Merkle trees
// 3) the execution proof builder, which produces a proof from the anchor block root to the target receipts root
//
// BuildExecutionProofForSlots is the main entry point for this orchestrator, which takes in a beacon client,
// an anchor slot, target slot, and the anchor block root, and returns an execution proof ready to be sent to the
// contract. BuildExecutionProofForBlock wraps it for a target execution block, and adds the Gloas path, where
// the target slot must be located and the proof also needs the execution header from an EL.
// TODO: Export in future PR.
type beaconClient interface {
	Block(ctx context.Context, slot uint64) (*beacon.SSZObject, error)
	State(ctx context.Context, slot uint64) (*beacon.SSZObject, error)
}

// liteStateCacheSize is the number of parsed lite beacon states kept in memory. A single anchor
// beacon state may be used to verify multiple target slots within the 8192-slot window of the
// anchor state's state_roots vector, so caching avoids refetching ~200MB state.
const liteStateCacheSize = 16

// ProofBuilder builds execution proofs from beacon data fetched via the beaconClient and
// cached parsed lite beacon states by slot.
type ProofBuilder struct {
	client          beaconClient
	el              ethClient
	forks           ForkConfig
	liteStates      *lru.Cache[uint64, *beacon.LiteBeaconState]
	gloasLiteStates *lru.Cache[uint64, *beacon.GloasLiteBeaconState]
}

// NewProofBuilder returns a ProofBuilder for Fulu payloads only.
func NewProofBuilder(client beaconClient) *ProofBuilder {
	return NewProofBuilderWithForks(client, nil, MainnetForkConfig)
}

// NewProofBuilderWithForks returns a ProofBuilder that picks each payload's proof path by its
// slot's fork in forks. Gloas proofs carry the execution header, which it fetches from el.
func NewProofBuilderWithForks(client beaconClient, el ethClient, forks ForkConfig) *ProofBuilder {
	return &ProofBuilder{
		client:          client,
		el:              el,
		forks:           forks,
		liteStates:      lru.NewCache[uint64, *beacon.LiteBeaconState](liteStateCacheSize),
		gloasLiteStates: lru.NewCache[uint64, *beacon.GloasLiteBeaconState](liteStateCacheSize),
	}
}

// ExecutionTarget identifies the execution block whose receipts root a proof reaches: its
// hash, and the slot of the beacon block that carried it as its execution payload. Under
// Gloas, that slot is the header's slotNumber (EIP-7843).
type ExecutionTarget struct {
	BlockHash   common.Hash
	PayloadSlot uint64
}

// BuildExecutionProofForBlock builds the complete execution proof linking the target
// execution block's receipts root to the confirmed anchor beacon block root at anchorSlot,
// taking the proof path of the payload slot's fork.
//
// A Fulu payload's proof is BuildExecutionProofForSlots's, with the payload slot as the target
// slot. A Gloas payload's proof targets c, the slot of the first beacon block after the
// payload's, which marks the payload as included: its post-state holds the execution block
// hash, which the execution header fetched from the EL must hash to.
func (b *ProofBuilder) BuildExecutionProofForBlock(
	ctx context.Context,
	anchorSlot uint64,
	anchorBlockRoot common.Hash,
	target ExecutionTarget,
) (*zkadapter.ExecutionProof, error) {
	if !b.forks.isGloas(target.PayloadSlot) {
		return b.BuildExecutionProofForSlots(ctx, anchorSlot, target.PayloadSlot, anchorBlockRoot)
	}
	return b.buildGloasExecutionProof(ctx, anchorSlot, anchorBlockRoot, target)
}

// buildGloasExecutionProof builds a Gloas payload's execution proof. Like
// BuildExecutionProofForSlots, it fetches and parses the anchor block and the anchor and
// target states, verifying each against the chain of trust, then also fetches the execution
// header for BuildGloasExecutionProof to check against the proven block hash.
func (b *ProofBuilder) buildGloasExecutionProof(
	ctx context.Context,
	anchorSlot uint64,
	anchorBlockRoot common.Hash,
	target ExecutionTarget,
) (*zkadapter.ExecutionProof, error) {
	targetSlot, err := locateCommittingSlot(ctx, b.client, target.PayloadSlot, anchorSlot, target.BlockHash)
	if err != nil {
		return nil, err
	}
	if err := validateSlotWindow(anchorSlot, targetSlot); err != nil {
		return nil, err
	}

	block, err := b.client.Block(ctx, anchorSlot)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch anchor block at slot %d: %w", anchorSlot, err)
	}
	anchorBlockTree, anchorStateRoot, err := beacon.ParseBlockTree(block)
	if err != nil {
		return nil, err
	}
	if got := common.BytesToHash(anchorBlockTree.Hash()); got != anchorBlockRoot {
		return nil, fmt.Errorf("anchor block at slot %d has root %s, expected confirmed root %s",
			anchorSlot, got, anchorBlockRoot)
	}

	anchorLite, err := b.getGloasLiteState(ctx, anchorSlot)
	if err != nil {
		return nil, fmt.Errorf("anchor state: %w", err)
	}
	anchorStateTree, err := anchorLite.AnchorStateTree(anchorStateRoot)
	if err != nil {
		return nil, fmt.Errorf("anchor state at slot %d: %w", anchorSlot, err)
	}

	targetLite, err := b.getGloasLiteState(ctx, targetSlot)
	if err != nil {
		return nil, fmt.Errorf("target state: %w", err)
	}
	targetStateTree, err := targetLite.TargetStateTree(anchorLite.StateRootAt(targetSlot))
	if err != nil {
		return nil, fmt.Errorf("target state at slot %d: %w", targetSlot, err)
	}

	headerRLP, err := fetchExecutionHeader(ctx, b.el, target.BlockHash)
	if err != nil {
		return nil, err
	}

	return BuildGloasExecutionProof(
		anchorBlockTree,
		anchorStateTree,
		targetStateTree,
		headerRLP,
		anchorBlockRoot,
		anchorSlot,
		targetSlot,
	)
}

// BuildExecutionProofForSlots builds the complete execution proof linking
// the targetSlot's receipts root to the confirmed anchor beacon block root at
// anchorSlot.
//
// In more detail, this function performs the following tasks:
// 1. Fetches the anchor beacon block and both beacon states (anchor and target)
// from the beaconClient. Note the anchorBlockRoot is the root ZKAdapter has confirmed for the
// anchorSlot. Beacon states are served from the lite state cache when present.
// 2. Parse the anchor beacon block into a regular Merkle tree, and the two beacon states into
// lite Merkle trees (to save memory).
// 3. Use the execution proof builder to create a proof from the parsed trees and provided slots.
// The execution proof builder will verify each tree against the expected chain of trust.
//
// TODO: We can cache the anchor's parsed lite state. The reason is that a single anchor beacon state
// may be used to verify multiple target slots within the 8192-slot window of the anchor state's
// state_roots vector. Issue: https://github.com/ava-labs/icm-services/issues/1542
func (b *ProofBuilder) BuildExecutionProofForSlots(
	ctx context.Context,
	anchorSlot uint64,
	targetSlot uint64,
	anchorBlockRoot common.Hash,
) (*zkadapter.ExecutionProof, error) {
	// Safety check
	if err := validateSlotWindow(anchorSlot, targetSlot); err != nil {
		return nil, err
	}

	// Get the SSZ-encoded anchor beacon block from the beaconClient and parse it into a tree.
	block, err := b.client.Block(ctx, anchorSlot)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch anchor block at slot %d: %w", anchorSlot, err)
	}
	// This path decodes Fulu objects only. The anchor's fork bounds the target's.
	if block.Fork != beacon.ForkFulu {
		return nil, fmt.Errorf("anchor block at slot %d is a %q object, expected %q",
			anchorSlot, block.Fork, beacon.ForkFulu)
	}
	anchorBlockTree, anchorStateRoot, err := beacon.ParseBlockTree(block)
	if err != nil {
		return nil, err
	}

	// Verify the fetched block is the confirmed anchor before fetching ~200MB
	// of state against it.
	if got := common.BytesToHash(anchorBlockTree.Hash()); got != anchorBlockRoot {
		return nil, fmt.Errorf("anchor block at slot %d has root %s, expected confirmed root %s",
			anchorSlot, got, anchorBlockRoot)
	}

	// Get the SSZ-encoded anchor beacon state from the beaconClient or the cache
	// and parse it into a lite tree.
	anchorLite, err := b.getLiteState(ctx, anchorSlot)
	if err != nil {
		return nil, fmt.Errorf("anchor state: %w", err)
	}

	anchorStateTree, err := anchorLite.AnchorStateTree(anchorStateRoot)
	if err != nil {
		return nil, fmt.Errorf("anchor state at slot %d: %w", anchorSlot, err)
	}

	// Get the SSZ-encoded target beacon state from the beaconClient or the cache
	// and parse it into a lite tree.
	targetLite, err := b.getLiteState(ctx, targetSlot)
	if err != nil {
		return nil, fmt.Errorf("target state: %w", err)
	}

	targetStateRoot := anchorLite.StateRootAt(targetSlot)
	targetStateTree, execHeaderTree, err := targetLite.TargetStateTree(targetStateRoot)
	if err != nil {
		return nil, fmt.Errorf("target state at slot %d: %w", targetSlot, err)
	}

	// Use the assembled trees to produce an execution proof from the beacon block root
	// down to the receipts root via the chain of trust.
	return BuildExecutionProof(
		anchorBlockTree,
		anchorStateTree,
		targetStateTree,
		execHeaderTree,
		anchorBlockRoot,
		anchorSlot,
		targetSlot,
	)
}

// liteState returns the parsed lite beacon state at the given slot, fetching the SSZ-encoded
// beacon state from the beaconClient and parsing it on a cache miss.
func (b *ProofBuilder) getLiteState(ctx context.Context, slot uint64) (*beacon.LiteBeaconState, error) {
	if lite, ok := b.liteStates.Get(slot); ok {
		return lite, nil
	}
	state, err := b.client.State(ctx, slot)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch beacon state at slot %d: %w", slot, err)
	}
	lite, err := beacon.ParseLiteBeaconState(state.Data)
	if err != nil {
		return nil, fmt.Errorf("failed to parse beacon state at slot %d: %w", slot, err)
	}
	b.liteStates.Put(slot, lite)
	return lite, nil
}

// getGloasLiteState returns the parsed Gloas lite beacon state at the given slot, fetching the
// SSZ-encoded beacon state from the beaconClient and parsing it on a cache miss.
func (b *ProofBuilder) getGloasLiteState(ctx context.Context, slot uint64) (*beacon.GloasLiteBeaconState, error) {
	if lite, ok := b.gloasLiteStates.Get(slot); ok {
		return lite, nil
	}
	state, err := b.client.State(ctx, slot)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch beacon state at slot %d: %w", slot, err)
	}
	lite, err := beacon.ParseGloasLiteBeaconState(state.Data)
	if err != nil {
		return nil, fmt.Errorf("failed to parse beacon state at slot %d: %w", slot, err)
	}
	b.gloasLiteStates.Put(slot, lite)
	return lite, nil
}

// fetchExecutionHeader returns the RLP encoding of the execution block header with the given
// hash, as reported by the EL. Callers must check that it hashes to a trusted block hash.
func fetchExecutionHeader(ctx context.Context, client ethClient, blockHash common.Hash) ([]byte, error) {
	var header json.RawMessage
	if err := client.CallContext(ctx, &header, "eth_getBlockByHash", blockHash, false); err != nil {
		return nil, fmt.Errorf("failed to fetch execution block %s: %w", blockHash, err)
	}
	headerRLP, err := encodeExecutionHeader(header)
	if err != nil {
		return nil, fmt.Errorf("execution block %s: %w", blockHash, err)
	}
	return headerRLP, nil
}
