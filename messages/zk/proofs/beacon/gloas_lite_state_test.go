// Copyright (C) 2026, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

// THIS IS AN EXAMPLE OF UNAUDITED CODE. DO NOT USE THIS IN PRODUCTION.

package beacon

import (
	"encoding/json"
	"os"
	"path/filepath"
	"testing"

	"github.com/OffchainLabs/go-bitfield"
	"github.com/attestantio/go-eth2-client/spec/altair"
	"github.com/attestantio/go-eth2-client/spec/capella"
	"github.com/attestantio/go-eth2-client/spec/deneb"
	"github.com/attestantio/go-eth2-client/spec/electra"
	"github.com/attestantio/go-eth2-client/spec/gloas"
	"github.com/attestantio/go-eth2-client/spec/phase0"
	"github.com/ava-labs/libevm/common"
	ssz "github.com/ferranbt/fastssz"
	dynssz "github.com/pk910/dynamic-ssz"
	"github.com/stretchr/testify/require"
)

// Gindices into the Gloas BeaconState that the proofs walk to, duplicated
// here because the beacon package does not import proofs: state_roots[i] sits
// at (353 << 13) + i and latest_block_hash at 2947. The real-state test below
// pins them against Sepolia data.
const (
	testGIndexGloasStateRootsBase = 353 << 13
	testGIndexLatestBlockHash     = 2947
)

// sepoliaGloasStatesDirEnv names a directory holding the real Sepolia states
// the Gloas fixture was built from: anchor_state.ssz (slot 11,397,760) and
// target_state.ssz (slot 11,397,715). At ~3.5 MB each they are not committed,
// so the test using them skips unless this is set.
const sepoliaGloasStatesDirEnv = "SEPOLIA_GLOAS_STATES_DIR"

// The Gloas lite trees must reproduce the state root go-eth2-client computes
// (every field holds distinct data, so a wrong or misplaced field root changes
// it), and their proofs must equal the ones dynamic-ssz takes from its full
// state tree.
func TestGloasLiteStateTreesMatchOracles(t *testing.T) {
	state := minimalGloasBeaconState(t)
	expectedRoot, err := state.HashTreeRoot()
	require.NoError(t, err)
	stateSSZ, err := state.MarshalSSZ()
	require.NoError(t, err)
	oracle, err := dynssz.GetGlobalDynSsz().GetTree(state)
	require.NoError(t, err)

	lite, err := ParseGloasLiteBeaconState(stateSSZ)
	require.NoError(t, err)

	// Anchor role: state_roots expanded, so any of its entries can be proven.
	anchorTree, err := lite.AnchorStateTree(expectedRoot)
	require.NoError(t, err)
	for _, i := range []int{0, 5, 8191} {
		gIndex := testGIndexGloasStateRootsBase + i
		proof, err := anchorTree.Prove(gIndex)
		require.NoError(t, err)
		expectedProof, err := oracle.Prove(gIndex)
		require.NoError(t, err)
		require.Equal(t, state.StateRoots[i][:], proof.Leaf)
		require.Equal(t, expectedProof.Hashes, proof.Hashes, "state_roots[%d]", i)
		require.Equal(t, common.Hash(state.StateRoots[i]), lite.StateRootAt(uint64(i)))
	}

	// Target role: latest_block_hash is a field leaf.
	targetTree, err := lite.TargetStateTree(expectedRoot)
	require.NoError(t, err)
	proof, err := targetTree.Prove(testGIndexLatestBlockHash)
	require.NoError(t, err)
	expectedProof, err := oracle.Prove(testGIndexLatestBlockHash)
	require.NoError(t, err)
	require.Equal(t, state.LatestBlockHash[:], proof.Leaf)
	require.Equal(t, expectedProof.Hashes, proof.Hashes)
}

// A wrong expected root must be rejected by assembly's root check, in both
// roles.
func TestGloasLiteStateRootMismatch(t *testing.T) {
	stateSSZ, err := minimalGloasBeaconState(t).MarshalSSZ()
	require.NoError(t, err)
	lite, err := ParseGloasLiteBeaconState(stateSSZ)
	require.NoError(t, err)

	wrong := common.Hash{0xff}
	_, err = lite.AnchorStateTree(wrong)
	require.ErrorContains(t, err, "root mismatch")
	_, err = lite.TargetStateTree(wrong)
	require.ErrorContains(t, err, "root mismatch")
}

// Malformed bytes must be rejected at deserialization.
func TestGloasLiteStateRejectsGarbage(t *testing.T) {
	_, err := ParseGloasLiteBeaconState([]byte{0x01, 0x02, 0x03})
	require.ErrorContains(t, err, "failed to deserialize beacon state")
}

// On the real Sepolia states, the lite trees must reproduce the state roots
// and the branches of the fixture, which an independent merkleizer built.
func TestGloasLiteStateSepoliaStates(t *testing.T) {
	dir := os.Getenv(sepoliaGloasStatesDirEnv)
	if dir == "" {
		t.Skipf("%s must be set", sepoliaGloasStatesDirEnv)
	}

	var fixture struct {
		Metadata struct {
			TargetSlot uint64 `json:"targetSlot"`
		} `json:"metadata"`
		ExecutionProof struct {
			AnchorBeaconStateRoot    common.Hash   `json:"anchorBeaconStateRoot"`
			TargetBeaconStateGIndex  int           `json:"targetBeaconStateGIndex"`
			TargetBeaconStateRoot    common.Hash   `json:"targetBeaconStateRoot"`
			TargetBeaconStateProof   []common.Hash `json:"targetBeaconStateProof"`
			ExecutionBlockHashGIndex int           `json:"executionBlockHashGIndex"`
			ExecutionBlockHash       common.Hash   `json:"executionBlockHash"`
			ExecutionBlockHashProof  []common.Hash `json:"executionBlockHashProof"`
		} `json:"executionProof"`
	}
	data, err := os.ReadFile("../testdata/sepolia_gloas_fixture.json")
	require.NoError(t, err)
	require.NoError(t, json.Unmarshal(data, &fixture))
	expected := fixture.ExecutionProof
	targetSlot := fixture.Metadata.TargetSlot
	require.Equal(t, testGIndexGloasStateRootsBase+int(targetSlot%StateRootsVectorSize), expected.TargetBeaconStateGIndex)
	require.Equal(t, testGIndexLatestBlockHash, expected.ExecutionBlockHashGIndex)

	anchor := parseGloasStateFile(t, filepath.Join(dir, "anchor_state.ssz"))
	anchorTree, err := anchor.AnchorStateTree(expected.AnchorBeaconStateRoot)
	require.NoError(t, err)
	require.Equal(t, expected.TargetBeaconStateRoot, anchor.StateRootAt(targetSlot))
	proof, err := anchorTree.Prove(expected.TargetBeaconStateGIndex)
	require.NoError(t, err)
	require.Equal(t, expected.TargetBeaconStateRoot.Bytes(), proof.Leaf)
	require.Equal(t, expected.TargetBeaconStateProof, toHashes(proof))

	target := parseGloasStateFile(t, filepath.Join(dir, "target_state.ssz"))
	targetTree, err := target.TargetStateTree(expected.TargetBeaconStateRoot)
	require.NoError(t, err)
	proof, err = targetTree.Prove(expected.ExecutionBlockHashGIndex)
	require.NoError(t, err)
	require.Equal(t, expected.ExecutionBlockHash.Bytes(), proof.Leaf)
	require.Equal(t, expected.ExecutionBlockHashProof, toHashes(proof))
}

func parseGloasStateFile(t *testing.T, path string) *GloasLiteBeaconState {
	t.Helper()
	stateSSZ, err := os.ReadFile(path)
	require.NoError(t, err)
	lite, err := ParseGloasLiteBeaconState(stateSSZ)
	require.NoError(t, err)
	return lite
}

func toHashes(proof *ssz.Proof) []common.Hash {
	out := make([]common.Hash, len(proof.Hashes))
	for i, h := range proof.Hashes {
		out[i] = common.BytesToHash(h)
	}
	return out
}

// minimalGloasBeaconState builds a small gloas.BeaconState that
// go-eth2-client will serialize and hash. Fixed-size vectors are at spec
// length, and every field, list and vector holds distinct non-zero data, so
// that a missing, wrong or misplaced field root changes the state root. The
// progressive lists hold 5 elements to span two spine batches.
func minimalGloasBeaconState(t *testing.T) *gloas.BeaconState {
	t.Helper()

	syncCommittee := func(seed byte) *altair.SyncCommittee {
		committee := &altair.SyncCommittee{AggregatePubkey: phase0.BLSPubKey{seed}}
		for i := range committee.Pubkeys {
			committee.Pubkeys[i] = phase0.BLSPubKey{seed, byte(i)}
		}
		return committee
	}
	roots := func(n int, seed byte) []phase0.Root {
		out := make([]phase0.Root, n)
		for i := range out {
			out[i] = phase0.Root{seed, byte(i), byte(i >> 8)}
		}
		return out
	}
	gweis := func(n int, base uint64) []phase0.Gwei {
		out := make([]phase0.Gwei, n)
		for i := range out {
			out[i] = phase0.Gwei(base + uint64(i))
		}
		return out
	}
	indices := func(n int, base uint64) []phase0.ValidatorIndex {
		out := make([]phase0.ValidatorIndex, n)
		for i := range out {
			out[i] = phase0.ValidatorIndex(base + uint64(i))
		}
		return out
	}

	const listLen = 5
	state := &gloas.BeaconState{
		GenesisTime:           1,
		GenesisValidatorsRoot: phase0.Root{0x01},
		Slot:                  100,
		Fork: &phase0.Fork{
			PreviousVersion: phase0.Version{0x01},
			CurrentVersion:  phase0.Version{0x02},
			Epoch:           1,
		},
		LatestBlockHeader: &phase0.BeaconBlockHeader{
			Slot:       99,
			ParentRoot: phase0.Root{0x02},
			StateRoot:  phase0.Root{0x03},
			BodyRoot:   phase0.Root{0x04},
		},
		BlockRoots:      roots(8192, 0x10),
		StateRoots:      roots(8192, 0x11),
		HistoricalRoots: roots(2, 0x12),
		ETH1Data: &phase0.ETH1Data{
			DepositRoot:  phase0.Root{0x05},
			DepositCount: 1,
			BlockHash:    make([]byte, 32),
		},
		ETH1DataVotes: []*phase0.ETH1Data{{
			DepositRoot:  phase0.Root{0x06},
			DepositCount: 2,
			BlockHash:    make([]byte, 32),
		}},
		ETH1DepositIndex:             3,
		Balances:                     gweis(listLen, 32_000_000_000),
		RANDAOMixes:                  roots(65536, 0x13),
		Slashings:                    gweis(8192, 1),
		PreviousEpochParticipation:   []altair.ParticipationFlags{1, 2, 3, 4, 5},
		CurrentEpochParticipation:    []altair.ParticipationFlags{6, 7, 0, 1, 2},
		JustificationBits:            bitfield.Bitvector4{0x05},
		PreviousJustifiedCheckpoint:  &phase0.Checkpoint{Epoch: 1, Root: phase0.Root{0x07}},
		CurrentJustifiedCheckpoint:   &phase0.Checkpoint{Epoch: 2, Root: phase0.Root{0x08}},
		FinalizedCheckpoint:          &phase0.Checkpoint{Epoch: 3, Root: phase0.Root{0x09}},
		InactivityScores:             []uint64{11, 12, 13, 14, 15},
		CurrentSyncCommittee:         syncCommittee(0x20),
		NextSyncCommittee:            syncCommittee(0x21),
		LatestBlockHash:              phase0.Hash32{0x0a},
		NextWithdrawalIndex:          4,
		NextWithdrawalValidatorIndex: 5,
		HistoricalSummaries: []*capella.HistoricalSummary{{
			BlockSummaryRoot: phase0.Root{0x0b},
			StateSummaryRoot: phase0.Root{0x0c},
		}},
		DepositRequestsStartIndex:     6,
		DepositBalanceToConsume:       7,
		ExitBalanceToConsume:          8,
		EarliestExitEpoch:             9,
		ConsolidationBalanceToConsume: 10,
		EarliestConsolidationEpoch:    11,
		ProposerLookahead:             indices(64, 100),
		NextWithdrawalBuilderIndex:    12,
		ExecutionPayloadAvailability:  make([]uint8, 1024),
		BuilderPendingPayments:        make([]*gloas.BuilderPendingPayment, 64),
		LatestExecutionPayloadBid: &gloas.ExecutionPayloadBid{
			ParentBlockHash:       phase0.Hash32{0x0d},
			ParentBlockRoot:       phase0.Root{0x0e},
			BlockHash:             phase0.Hash32{0x0f},
			PrevRandao:            phase0.Root{0x10},
			GasLimit:              13,
			BuilderIndex:          14,
			Slot:                  100,
			Value:                 15,
			ExecutionPayment:      16,
			BlobKZGCommitments:    []deneb.KZGCommitment{{0x11}, {0x12}},
			ExecutionRequestsRoot: phase0.Root{0x13},
		},
		PTCWindow: make([][]phase0.ValidatorIndex, 96),
	}
	for i := range listLen {
		b := byte(i + 1)
		state.Validators = append(state.Validators, &phase0.Validator{
			PublicKey:                  phase0.BLSPubKey{0x30, b},
			WithdrawalCredentials:      append([]byte{0x31, b}, make([]byte, 30)...),
			EffectiveBalance:           phase0.Gwei(i + 1),
			Slashed:                    i%2 == 0,
			ActivationEligibilityEpoch: phase0.Epoch(i + 2),
			ActivationEpoch:            phase0.Epoch(i + 3),
			ExitEpoch:                  phase0.Epoch(i + 4),
			WithdrawableEpoch:          phase0.Epoch(i + 5),
		})
		state.PendingDeposits = append(state.PendingDeposits, &electra.PendingDeposit{
			Pubkey:                phase0.BLSPubKey{0x32, b},
			WithdrawalCredentials: append([]byte{0x33, b}, make([]byte, 30)...),
			Amount:                phase0.Gwei(i + 6),
			Signature:             phase0.BLSSignature{0x34, b},
			Slot:                  phase0.Slot(i + 7),
		})
		state.PendingPartialWithdrawals = append(state.PendingPartialWithdrawals, &electra.PendingPartialWithdrawal{
			ValidatorIndex:    phase0.ValidatorIndex(i + 8),
			Amount:            phase0.Gwei(i + 9),
			WithdrawableEpoch: phase0.Epoch(i + 10),
		})
		state.PendingConsolidations = append(state.PendingConsolidations, &electra.PendingConsolidation{
			SourceIndex: phase0.ValidatorIndex(i + 11),
			TargetIndex: phase0.ValidatorIndex(i + 12),
		})
		state.Builders = append(state.Builders, &gloas.Builder{
			PublicKey:         phase0.BLSPubKey{0x35, b},
			Version:           b,
			ExecutionAddress:  [20]byte{0x36, b},
			Balance:           phase0.Gwei(i + 13),
			DepositEpoch:      phase0.Epoch(i + 14),
			WithdrawableEpoch: phase0.Epoch(i + 15),
		})
		state.BuilderPendingWithdrawals = append(state.BuilderPendingWithdrawals, &gloas.BuilderPendingWithdrawal{
			FeeRecipient: [20]byte{0x37, b},
			Amount:       phase0.Gwei(i + 16),
			BuilderIndex: gloas.BuilderIndex(i + 17),
		})
		state.PayloadExpectedWithdrawals = append(state.PayloadExpectedWithdrawals, &capella.Withdrawal{
			Index:          capella.WithdrawalIndex(i + 18),
			ValidatorIndex: phase0.ValidatorIndex(i + 19),
			Address:        [20]byte{0x38, b},
			Amount:         phase0.Gwei(i + 20),
		})
	}
	for i := range state.ExecutionPayloadAvailability {
		state.ExecutionPayloadAvailability[i] = byte(i)
	}
	for i := range state.BuilderPendingPayments {
		state.BuilderPendingPayments[i] = &gloas.BuilderPendingPayment{
			Weight: phase0.Gwei(i + 1),
			Withdrawal: &gloas.BuilderPendingWithdrawal{
				FeeRecipient: [20]byte{0x39, byte(i)},
				Amount:       phase0.Gwei(i + 2),
				BuilderIndex: gloas.BuilderIndex(i + 3),
			},
			ProposerIndex: phase0.ValidatorIndex(i + 4),
		}
	}
	for i := range state.PTCWindow {
		state.PTCWindow[i] = indices(512, uint64(i)*1000)
	}
	return state
}
