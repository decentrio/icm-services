// Copyright (C) 2026, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

// THIS IS AN EXAMPLE OF UNAUDITED CODE. DO NOT USE THIS IN PRODUCTION.

package beacon

import (
	"fmt"

	"github.com/attestantio/go-eth2-client/spec/gloas"
	"github.com/attestantio/go-eth2-client/spec/phase0"
	"github.com/ava-labs/libevm/common"
	ssz "github.com/ferranbt/fastssz"
)

// Constants for building the Gloas lite beacon state, derived from the Gloas
// BeaconState definition:
// https://github.com/ethereum/consensus-specs/blob/master/specs/gloas/beacon-chain.md#beaconstate
const (
	// numGloasStateFields is the BeaconState's field count. Gloas makes it a
	// progressive container (EIP-7688), so the fields are not padded to a
	// power of two: they sit in the progressive tree's spine batches.
	numGloasStateFields = 46

	// gloasPTCWindowRows is the number of rows of ptc_window,
	// (2 + MIN_SEED_LOOKAHEAD) * SLOTS_PER_EPOCH.
	gloasPTCWindowRows = 96
)

// gloasActiveFields is the BeaconState's active_fields bitvector, packed
// into a chunk: one bit per field, all 46 set. The progressive container's
// root mixes it in: hash(merkleize_progressive(field roots), active_fields).
var gloasActiveFields = func() [32]byte {
	var bits [32]byte
	for i := range numGloasStateFields {
		bits[i/8] |= 1 << (i % 8)
	}
	return bits
}()

// GloasLiteBeaconState is a minimal Gloas beacon state reduced to what proof
// building needs, like LiteBeaconState for Fulu: the 46 field roots, plus the
// state_roots vector so the anchor role can expand it. The target role needs
// no expanded subtree, since latest_block_hash is itself a field leaf.
type GloasLiteBeaconState struct {
	fieldRoots [numGloasStateFields][32]byte
	stateRoots [][32]byte
}

// ParseGloasLiteBeaconState deserializes an SSZ encoded Gloas BeaconState and
// reduces it to a GloasLiteBeaconState. Field roots are computed per field, in
// consensus spec order; the full state tree is never built. A wrong or
// missing field changes the assembled root and is rejected by the tree
// constructors' root check.
func ParseGloasLiteBeaconState(stateSSZ []byte) (*GloasLiteBeaconState, error) {
	var state gloas.BeaconState
	if err := state.UnmarshalSSZ(stateSSZ); err != nil {
		return nil, fmt.Errorf("failed to deserialize beacon state: %w", err)
	}

	lite := &GloasLiteBeaconState{stateRoots: rootsToFixed(state.StateRoots)}
	r := lite.fieldRoots[:]
	r[0] = uint64Root(state.GenesisTime)
	r[1] = state.GenesisValidatorsRoot
	r[2] = uint64Root(uint64(state.Slot))
	r[3] = mustRoot(state.Fork.HashTreeRoot())
	r[4] = mustRoot(state.LatestBlockHeader.HashTreeRoot())
	r[5] = rootsVectorRoot(state.BlockRoots)
	r[6] = rootsVectorRoot(state.StateRoots)
	r[7] = rootsListRoot(state.HistoricalRoots, historicalRootsLimit)
	r[8] = mustRoot(state.ETH1Data.HashTreeRoot())
	r[9] = containersListRoot(state.ETH1DataVotes, eth1DataVotesLimit)
	r[10] = uint64Root(state.ETH1DepositIndex)
	r[11] = containersProgressiveListRoot(state.Validators)
	r[12] = uint64sProgressiveListRoot(gweiToUint64s(state.Balances))
	r[13] = rootsVectorRoot(state.RANDAOMixes)
	r[14] = uint64sVectorRoot(gweiToUint64s(state.Slashings))
	r[15] = bytesProgressiveListRoot(participationToBytes(state.PreviousEpochParticipation))
	r[16] = bytesProgressiveListRoot(participationToBytes(state.CurrentEpochParticipation))
	r[17] = bytesChunkRoot(state.JustificationBits)
	r[18] = mustRoot(state.PreviousJustifiedCheckpoint.HashTreeRoot())
	r[19] = mustRoot(state.CurrentJustifiedCheckpoint.HashTreeRoot())
	r[20] = mustRoot(state.FinalizedCheckpoint.HashTreeRoot())
	r[21] = uint64sProgressiveListRoot(state.InactivityScores)
	r[22] = mustRoot(state.CurrentSyncCommittee.HashTreeRoot())
	r[23] = mustRoot(state.NextSyncCommittee.HashTreeRoot())
	r[24] = state.LatestBlockHash
	r[25] = uint64Root(uint64(state.NextWithdrawalIndex))
	r[26] = uint64Root(uint64(state.NextWithdrawalValidatorIndex))
	r[27] = containersListRoot(state.HistoricalSummaries, historicalSummariesLimit)
	r[28] = uint64Root(state.DepositRequestsStartIndex)
	r[29] = uint64Root(uint64(state.DepositBalanceToConsume))
	r[30] = uint64Root(uint64(state.ExitBalanceToConsume))
	r[31] = uint64Root(uint64(state.EarliestExitEpoch))
	r[32] = uint64Root(uint64(state.ConsolidationBalanceToConsume))
	r[33] = uint64Root(uint64(state.EarliestConsolidationEpoch))
	r[34] = containersProgressiveListRoot(state.PendingDeposits)
	r[35] = containersProgressiveListRoot(state.PendingPartialWithdrawals)
	r[36] = containersProgressiveListRoot(state.PendingConsolidations)
	r[37] = uint64sVectorRoot(validatorIndicesToUint64s(state.ProposerLookahead))
	r[38] = containersProgressiveListRoot(state.Builders)
	r[39] = uint64Root(uint64(state.NextWithdrawalBuilderIndex))
	r[40] = merkleizeChunks(packBytes(state.ExecutionPayloadAvailability))
	r[41] = containersVectorRoot(state.BuilderPendingPayments)
	r[42] = containersProgressiveListRoot(state.BuilderPendingWithdrawals)
	r[43] = mustRoot(state.LatestExecutionPayloadBid.HashTreeRoot())
	r[44] = containersProgressiveListRoot(state.PayloadExpectedWithdrawals)
	r[45] = ptcWindowRoot(state.PTCWindow)

	return lite, nil
}

// AnchorStateTree builds the anchor lite beacon state tree, verified against
// expectedRoot. state_roots is expanded as a real subtree and every other
// field is kept as a root leaf.
func (s *GloasLiteBeaconState) AnchorStateTree(expectedRoot common.Hash) (*ssz.Node, error) {
	stateRootsSubtree, err := ssz.TreeFromChunks(toByteSlices(s.stateRoots))
	if err != nil {
		return nil, fmt.Errorf("failed to build state_roots subtree: %w", err)
	}
	return s.assemble(map[int]*ssz.Node{stateRootsLeafIndex: stateRootsSubtree}, expectedRoot)
}

// TargetStateTree builds the target lite beacon state tree, verified against
// expectedRoot, the target state root found in the anchor state's
// state_roots. Every field is kept as a root leaf: the execution block hash
// is proven as the latest_block_hash leaf itself.
func (s *GloasLiteBeaconState) TargetStateTree(expectedRoot common.Hash) (*ssz.Node, error) {
	return s.assemble(nil, expectedRoot)
}

// StateRootAt returns the state root the state_roots history vector holds
// for the given slot.
func (s *GloasLiteBeaconState) StateRootAt(slot uint64) common.Hash {
	return common.Hash(s.stateRoots[slot%StateRootsVectorSize])
}

// assemble builds the progressive container tree of the beacon state: the
// progressive tree of the field leaves on the left and the active_fields
// chunk on the right. Expanded subtrees are provided in the map, keyed by
// their field index.
func (s *GloasLiteBeaconState) assemble(
	expanded map[int]*ssz.Node,
	expectedRoot common.Hash,
) (*ssz.Node, error) {
	leaves := make([]*ssz.Node, numGloasStateFields)
	for i := range leaves {
		if subtree, ok := expanded[i]; ok {
			leaves[i] = subtree
			continue
		}
		root := s.fieldRoots[i]
		leaves[i] = ssz.LeafFromBytes(root[:])
	}
	fields, err := progressiveTree(leaves)
	if err != nil {
		return nil, fmt.Errorf("failed to assemble lite state tree: %w", err)
	}
	tree := ssz.NewNodeWithLR(fields, ssz.LeafFromBytes(gloasActiveFields[:]))

	if got := common.BytesToHash(tree.Hash()); got != expectedRoot {
		return nil, fmt.Errorf("lite state root mismatch: computed %s, expected %s", got, expectedRoot)
	}
	return tree, nil
}

// --- Gloas field root helpers ----------------------------------------------

// containersProgressiveListRoot merkleizes a progressive list of SSZ
// containers, mixing in the element count. Each element's root comes from
// its own HashTreeRoot.
func containersProgressiveListRoot[T interface{ HashTreeRoot() ([32]byte, error) }](items []T) [32]byte {
	return progressiveListRoot(containerRoots(items))
}

// uint64sProgressiveListRoot merkleizes a progressive list of uint64s,
// packed four per chunk, mixing in the element count.
func uint64sProgressiveListRoot(values []uint64) [32]byte {
	return mixinLength(merkleizeProgressive(packUint64s(values)), uint64(len(values)))
}

// bytesProgressiveListRoot merkleizes a progressive byte list, packed 32 per
// chunk, mixing in the byte count.
func bytesProgressiveListRoot(data []byte) [32]byte {
	return mixinLength(merkleizeProgressive(packBytes(data)), uint64(len(data)))
}

// containersVectorRoot merkleizes a fixed-size vector of SSZ containers.
func containersVectorRoot[T interface{ HashTreeRoot() ([32]byte, error) }](items []T) [32]byte {
	return merkleizeChunks(containerRoots(items))
}

// ptcWindowRoot merkleizes ptc_window, a vector of rows, each a vector of
// validator indices.
func ptcWindowRoot(rows [][]phase0.ValidatorIndex) [32]byte {
	rowRoots := make([][32]byte, gloasPTCWindowRows)
	for i, row := range rows {
		rowRoots[i] = uint64sVectorRoot(validatorIndicesToUint64s(row))
	}
	return merkleizeChunks(rowRoots)
}

func containerRoots[T interface{ HashTreeRoot() ([32]byte, error) }](items []T) [][32]byte {
	roots := make([][32]byte, len(items))
	for i, item := range items {
		roots[i] = mustRoot(item.HashTreeRoot())
	}
	return roots
}

// packBytes packs bytes 32 per chunk, zero-padding the last chunk.
func packBytes(data []byte) [][32]byte {
	chunks := make([][32]byte, (len(data)+31)/32)
	for i := range chunks {
		copy(chunks[i][:], data[i*32:min((i+1)*32, len(data))])
	}
	return chunks
}
