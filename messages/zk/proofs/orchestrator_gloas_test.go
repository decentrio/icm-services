// Copyright (C) 2026, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

// THIS IS AN EXAMPLE OF UNAUDITED CODE. DO NOT USE THIS IN PRODUCTION.

package proofs

import (
	"context"
	"encoding/json"
	"os"
	"path/filepath"
	"testing"

	"github.com/attestantio/go-eth2-client/spec/altair"
	"github.com/attestantio/go-eth2-client/spec/gloas"
	"github.com/attestantio/go-eth2-client/spec/phase0"
	"github.com/ava-labs/icm-services/messages/zk/proofs/beacon"
	"github.com/ava-labs/libevm/common"
	"github.com/ava-labs/libevm/rlp"
	ssz "github.com/ferranbt/fastssz"
	"github.com/stretchr/testify/require"
)

// sepoliaGloasStatesDirEnv names a directory holding the real Sepolia states
// the Gloas fixture was built from: anchor_state.ssz (slot 11,397,760) and
// target_state.ssz (slot 11,397,715). At ~3.5 MB each they are not committed,
// so the test using them skips unless this is set.
const sepoliaGloasStatesDirEnv = "SEPOLIA_GLOAS_STATES_DIR"

// gloasMockFixture chains synthetic Gloas beacon data to the real Sepolia
// execution block, at the fixture's slots: block c's bid builds on the
// block, the state at c holds its hash as latest_block_hash, and the anchor
// state's state_roots holds that state's root. The EL serves the real block.
type gloasMockFixture struct {
	beacon          *mockBeaconClient
	el              *rpcFileClient
	target          ExecutionTarget
	anchorBlockRoot common.Hash
	anchorStateRoot common.Hash
	targetStateRoot common.Hash
}

func newGloasMockFixture(t *testing.T) *gloasMockFixture {
	t.Helper()
	fixture := loadSepoliaFixture(t)
	blockHash := fixture.Metadata.BlockHash

	targetState := minimalGloasBeaconState(sepoliaTargetSlot)
	targetState.LatestBlockHash = phase0.Hash32(blockHash)
	targetStateRoot, err := targetState.HashTreeRoot()
	require.NoError(t, err)

	anchorState := minimalGloasBeaconState(sepoliaAnchorSlot)
	anchorState.StateRoots[sepoliaTargetSlot%StateRootsVectorSize] = targetStateRoot
	anchorStateRoot, err := anchorState.HashTreeRoot()
	require.NoError(t, err)

	anchorBlock := minimalGloasBlock(sepoliaAnchorSlot, anchorStateRoot, common.Hash{0xcc})
	anchorBlockRoot, err := anchorBlock.Message.HashTreeRoot()
	require.NoError(t, err)

	return &gloasMockFixture{
		beacon: &mockBeaconClient{
			fork: beacon.ForkGloas,
			blocks: map[uint64][]byte{
				sepoliaAnchorSlot: mustMarshalSSZ(t, anchorBlock),
				sepoliaTargetSlot: mustMarshalSSZ(t, minimalGloasBlock(sepoliaTargetSlot, phase0.Root{}, blockHash)),
			},
			states: map[uint64][]byte{
				sepoliaAnchorSlot: mustMarshalSSZ(t, anchorState),
				sepoliaTargetSlot: mustMarshalSSZ(t, targetState),
			},
			stateFetches: map[uint64]int{},
		},
		el:              newRPCFileClient(t, sepoliaBlockPath, sepoliaReceiptsPath),
		target:          ExecutionTarget{BlockHash: blockHash, PayloadSlot: sepoliaPayloadSlot},
		anchorBlockRoot: common.Hash(anchorBlockRoot),
		anchorStateRoot: common.Hash(anchorStateRoot),
		targetStateRoot: common.Hash(targetStateRoot),
	}
}

func (f *gloasMockFixture) builder() *ProofBuilder {
	return NewProofBuilderWithForks(f.beacon, f.el, SepoliaForkConfig)
}

// The fork configs must hold each network's GLOAS_FORK_EPOCH, in slots.
func TestForkConfigs(t *testing.T) {
	fixture := loadSepoliaFixture(t)
	require.Equal(t, fixture.Metadata.GloasForkEpoch*32, SepoliaForkConfig.GloasForkSlot)
	require.Equal(t, uint64(132352*32), HoodiForkConfig.GloasForkSlot)
	require.False(t, MainnetForkConfig.isGloas(1<<62), "Gloas is not scheduled on mainnet")

	require.False(t, SepoliaForkConfig.isGloas(SepoliaForkConfig.GloasForkSlot-1))
	require.True(t, SepoliaForkConfig.isGloas(SepoliaForkConfig.GloasForkSlot))
}

// A Gloas payload's proof must run from the anchor to the state at c, then
// to its latest_block_hash, and carry the header that hashes to it. Every
// branch must verify against the roots the synthetic data was built with.
func TestBuildExecutionProofForBlockGloas(t *testing.T) {
	f := newGloasMockFixture(t)
	fixture := loadSepoliaFixture(t)

	proof, err := f.builder().BuildExecutionProofForBlock(
		context.Background(), sepoliaAnchorSlot, f.anchorBlockRoot, f.target)
	require.NoError(t, err)

	require.Equal(t, sepoliaAnchorSlot, proof.AnchorSlot)
	require.Equal(t, sepoliaTargetSlot, proof.TargetSlot)
	require.Equal(t, f.anchorStateRoot, common.Hash(proof.AnchorBeaconStateRoot))
	verifyBranch(t, f.anchorBlockRoot, gIndexBlockStateRoot, f.anchorStateRoot, proof.AnchorBeaconStateProof)

	require.Equal(t, f.targetStateRoot, common.Hash(proof.TargetBeaconStateRoot))
	historyGIndex := (gIndexBaseStateRootsGloas << stateRootsDepth) + int(sepoliaTargetSlot%StateRootsVectorSize)
	verifyBranch(t, f.anchorStateRoot, historyGIndex, f.targetStateRoot, proof.TargetBeaconStateProof)

	require.Equal(t, f.target.BlockHash, common.Hash(proof.TargetExecutionHeaderRoot))
	verifyBranch(t, f.targetStateRoot, gIndexLatestBlockHash, f.target.BlockHash, proof.TargetExecutionHeaderProof)

	require.Equal(t, []byte(fixture.ExecutionProof.ExecutionHeaderRlp), proof.ExecutionHeaderRlp)
	require.Equal(t, fixture.ExecutionProof.ReceiptsRoot, common.Hash(proof.TargetReceiptsRoot))
	require.Empty(t, proof.TargetReceiptsProof)
}

// A header that does not hash to the proven latest_block_hash must be
// rejected: its receipts root would not be the block's.
func TestBuildExecutionProofForBlockGloasHeaderMismatch(t *testing.T) {
	f := newGloasMockFixture(t)
	var header map[string]json.RawMessage
	require.NoError(t, json.Unmarshal(f.el.block, &header))
	header["gasUsed"] = json.RawMessage(`"0x1"`)
	var err error
	f.el.block, err = json.Marshal(header)
	require.NoError(t, err)

	_, err = f.builder().BuildExecutionProofForBlock(
		context.Background(), sepoliaAnchorSlot, f.anchorBlockRoot, f.target)
	require.ErrorContains(t, err, "does not hash to the proven execution block hash")
}

// A payload the next block did not build on cannot be proven.
func TestBuildExecutionProofForBlockGloasPayloadNotIncluded(t *testing.T) {
	f := newGloasMockFixture(t)
	f.target.BlockHash = common.Hash{0xaa}

	_, err := f.builder().BuildExecutionProofForBlock(
		context.Background(), sepoliaAnchorSlot, f.anchorBlockRoot, f.target)
	require.ErrorIs(t, err, ErrPayloadNotIncluded)
}

// The beacon block committing the payload must be within the anchor's
// state_roots window.
func TestBuildExecutionProofForBlockGloasOutsideWindow(t *testing.T) {
	f := newGloasMockFixture(t)

	_, err := f.builder().BuildExecutionProofForBlock(
		context.Background(), sepoliaTargetSlot+StateRootsVectorSize+1, f.anchorBlockRoot, f.target)
	require.ErrorContains(t, err, "state_roots window")
}

// A non-confirmed anchor beacon block root must be rejected.
func TestBuildExecutionProofForBlockGloasWrongAnchorRoot(t *testing.T) {
	f := newGloasMockFixture(t)

	_, err := f.builder().BuildExecutionProofForBlock(
		context.Background(), sepoliaAnchorSlot, common.Hash{0xba, 0xd}, f.target)
	require.ErrorContains(t, err, "expected confirmed root")
}

// A pre-Gloas payload must take the Fulu path, unchanged: the proof equals
// BuildExecutionProofForSlots's, with no header RLP.
func TestBuildExecutionProofForBlockFulu(t *testing.T) {
	f := newMockFixture(t)
	forks := ForkConfig{GloasForkSlot: f.anchorSlot + 1}
	builder := NewProofBuilderWithForks(f.client, nil, forks)

	proof, err := builder.BuildExecutionProofForBlock(
		context.Background(), f.anchorSlot, f.anchorBlockRoot, ExecutionTarget{PayloadSlot: f.targetSlot})
	require.NoError(t, err)

	expected, err := NewProofBuilder(f.client).BuildExecutionProofForSlots(
		context.Background(), f.anchorSlot, f.targetSlot, f.anchorBlockRoot)
	require.NoError(t, err)
	require.Equal(t, expected, proof)
	require.Empty(t, proof.ExecutionHeaderRlp)
}

// On the real Sepolia data, the pipeline must reproduce the fixture's
// execution proof byte for byte. The fixture was built by an independent
// implementation, so this is the cross-implementation check for Gloas.
func TestBuildExecutionProofForBlockSepolia(t *testing.T) {
	dir := os.Getenv(sepoliaGloasStatesDirEnv)
	if dir == "" {
		t.Skipf("%s must be set", sepoliaGloasStatesDirEnv)
	}
	fixture := loadSepoliaFixture(t)
	expected := fixture.ExecutionProof
	readState := func(name string) []byte {
		data, err := os.ReadFile(filepath.Join(dir, name))
		require.NoError(t, err)
		return data
	}
	beaconClient := &mockBeaconClient{
		fork: beacon.ForkGloas,
		blocks: map[uint64][]byte{
			sepoliaAnchorSlot: readSepoliaBlock(t, sepoliaAnchorSlot).Data,
			sepoliaTargetSlot: readSepoliaBlock(t, sepoliaTargetSlot).Data,
		},
		states: map[uint64][]byte{
			sepoliaAnchorSlot: readState("anchor_state.ssz"),
			sepoliaTargetSlot: readState("target_state.ssz"),
		},
		stateFetches: map[uint64]int{},
	}
	el := newRPCFileClient(t, sepoliaBlockPath, sepoliaReceiptsPath)
	builder := NewProofBuilderWithForks(beaconClient, el, SepoliaForkConfig)

	proof, err := builder.BuildExecutionProofForBlock(
		context.Background(),
		expected.AnchorSlot,
		fixture.AnchorBeaconBlockRoot,
		ExecutionTarget{BlockHash: fixture.Metadata.BlockHash, PayloadSlot: fixture.Metadata.PayloadSlot},
	)
	require.NoError(t, err)

	require.Equal(t, expected.AnchorSlot, proof.AnchorSlot)
	require.Equal(t, expected.TargetSlot, proof.TargetSlot)
	require.Equal(t, expected.AnchorBeaconStateRoot, common.Hash(proof.AnchorBeaconStateRoot))
	require.Equal(t, expected.AnchorBeaconStateProof, toCommonHashes(proof.AnchorBeaconStateProof))
	require.Equal(t, expected.TargetBeaconStateRoot, common.Hash(proof.TargetBeaconStateRoot))
	require.Equal(t, expected.TargetBeaconStateProof, toCommonHashes(proof.TargetBeaconStateProof))
	require.Equal(t, expected.ExecutionBlockHash, common.Hash(proof.TargetExecutionHeaderRoot))
	require.Equal(t, expected.ExecutionBlockHashProof, toCommonHashes(proof.TargetExecutionHeaderProof))
	require.Equal(t, []byte(expected.ExecutionHeaderRlp), proof.ExecutionHeaderRlp)
	require.Equal(t, expected.ReceiptsRoot, common.Hash(proof.TargetReceiptsRoot))
	require.Empty(t, proof.TargetReceiptsProof)
}

// The receipts root is the header's RLP item 5; anything that is not a list
// with a 32-byte item there is rejected.
func TestHeaderReceiptsRoot(t *testing.T) {
	receiptsRoot := common.Hash{0x55}
	encode := func(items ...any) []byte {
		data, err := rlp.EncodeToBytes(items)
		require.NoError(t, err)
		return data
	}
	zero := []byte{}

	got, err := headerReceiptsRoot(encode(zero, zero, zero, zero, zero, receiptsRoot, zero))
	require.NoError(t, err)
	require.Equal(t, receiptsRoot, got)

	_, err = headerReceiptsRoot(encode(zero, zero, zero, zero, zero))
	require.ErrorContains(t, err, "execution header has only 5 fields")

	_, err = headerReceiptsRoot(encode(zero, zero, zero, zero, zero, []byte{0x55}))
	require.ErrorContains(t, err, "receipts root")

	notAList, err := rlp.EncodeToBytes(receiptsRoot)
	require.NoError(t, err)
	_, err = headerReceiptsRoot(notAList)
	require.ErrorContains(t, err, "failed to decode execution header")
}

// verifyBranch checks that leaf sits at gIndex under root, given the
// branch's siblings.
func verifyBranch(t *testing.T, root common.Hash, gIndex int, leaf common.Hash, branch [][32]byte) {
	t.Helper()
	hashes := make([][]byte, len(branch))
	for i := range branch {
		hashes[i] = branch[i][:]
	}
	ok, err := ssz.VerifyProof(root.Bytes(), &ssz.Proof{Index: gIndex, Leaf: leaf.Bytes(), Hashes: hashes})
	require.NoError(t, err)
	require.True(t, ok, "branch to gindex %d", gIndex)
}

func toCommonHashes(hashes [][32]byte) []common.Hash {
	out := make([]common.Hash, len(hashes))
	for i := range hashes {
		out[i] = hashes[i]
	}
	return out
}

func mustMarshalSSZ(t *testing.T, obj interface{ MarshalSSZ() ([]byte, error) }) []byte {
	t.Helper()
	data, err := obj.MarshalSSZ()
	require.NoError(t, err)
	return data
}

// minimalGloasBeaconState builds the smallest gloas.BeaconState that
// go-eth2-client will serialize and hash: fixed-size vectors at spec length,
// lists empty, containers populated.
// TODO: This is duplicated in the beacon package's tests pending package
// consolidation. See Issue: https://github.com/ava-labs/icm-services/issues/1541
func minimalGloasBeaconState(slot uint64) *gloas.BeaconState {
	syncCommittee := &altair.SyncCommittee{}
	return &gloas.BeaconState{
		Slot:                         phase0.Slot(slot),
		Fork:                         &phase0.Fork{},
		LatestBlockHeader:            &phase0.BeaconBlockHeader{Slot: phase0.Slot(slot) - 1},
		BlockRoots:                   make([]phase0.Root, 8192),
		StateRoots:                   make([]phase0.Root, 8192),
		ETH1Data:                     &phase0.ETH1Data{BlockHash: make([]byte, 32)},
		RANDAOMixes:                  make([]phase0.Root, 65536),
		Slashings:                    make([]phase0.Gwei, 8192),
		JustificationBits:            make([]byte, 1),
		PreviousJustifiedCheckpoint:  &phase0.Checkpoint{},
		CurrentJustifiedCheckpoint:   &phase0.Checkpoint{},
		FinalizedCheckpoint:          &phase0.Checkpoint{},
		CurrentSyncCommittee:         syncCommittee,
		NextSyncCommittee:            syncCommittee,
		ProposerLookahead:            make([]phase0.ValidatorIndex, 64),
		ExecutionPayloadAvailability: make([]uint8, 1024),
		BuilderPendingPayments:       make([]*gloas.BuilderPendingPayment, 64),
		LatestExecutionPayloadBid:    &gloas.ExecutionPayloadBid{},
		PTCWindow:                    make([][]phase0.ValidatorIndex, 96),
	}
}
