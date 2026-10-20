// Copyright (C) 2026, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package proofs

import (
	"context"
	"errors"
	"fmt"
	"os"
	"testing"

	"github.com/OffchainLabs/go-bitfield"
	"github.com/attestantio/go-eth2-client/spec/altair"
	"github.com/attestantio/go-eth2-client/spec/gloas"
	"github.com/attestantio/go-eth2-client/spec/phase0"
	"github.com/ava-labs/icm-services/messages/zk/proofs/beacon"
	"github.com/ava-labs/libevm/common"
	"github.com/stretchr/testify/require"
)

// Slots of the Sepolia fixture: the payload's own slot s and the slot c of
// the first block after it, which built on it.
const (
	sepoliaPayloadSlot = uint64(11397714)
	sepoliaTargetSlot  = uint64(11397715)
	sepoliaAnchorSlot  = uint64(11397760)
)

// slotBlocks serves beacon blocks by slot. Slots with no block are missed,
// unless errs holds a failure for them.
type slotBlocks struct {
	blocks  map[uint64]*beacon.SSZObject
	errs    map[uint64]error
	fetched []uint64
}

func (s *slotBlocks) Block(_ context.Context, slot uint64) (*beacon.SSZObject, error) {
	s.fetched = append(s.fetched, slot)
	if err, ok := s.errs[slot]; ok {
		return nil, err
	}
	block, ok := s.blocks[slot]
	if !ok {
		return nil, fmt.Errorf("%w %d", beacon.ErrNoBlockAtSlot, slot)
	}
	return block, nil
}

// readSepoliaBlock reads a real Sepolia Gloas block from testdata.
func readSepoliaBlock(t *testing.T, slot uint64) *beacon.SSZObject {
	t.Helper()
	data, err := os.ReadFile(fmt.Sprintf("testdata/sepolia/beacon_block_%d.ssz", slot))
	require.NoError(t, err)
	return &beacon.SSZObject{Fork: beacon.ForkGloas, Data: data}
}

// gloasBlock builds a minimal Gloas block at slot whose bid builds on the
// execution block parentBlockHash.
func gloasBlock(t *testing.T, slot uint64, parentBlockHash common.Hash) *beacon.SSZObject {
	t.Helper()
	block := &gloas.SignedBeaconBlock{
		Message: &gloas.BeaconBlock{
			Slot: phase0.Slot(slot),
			Body: &gloas.BeaconBlockBody{
				ETH1Data: &phase0.ETH1Data{BlockHash: make([]byte, 32)},
				SyncAggregate: &altair.SyncAggregate{
					SyncCommitteeBits: bitfield.NewBitvector512(),
				},
				SignedExecutionPayloadBid: &gloas.SignedExecutionPayloadBid{
					Message: &gloas.ExecutionPayloadBid{
						ParentBlockHash: phase0.Hash32(parentBlockHash),
						BlockHash:       phase0.Hash32{0xbb},
						Slot:            phase0.Slot(slot),
					},
				},
			},
		},
	}
	data, err := block.MarshalSSZ()
	require.NoError(t, err)
	return &beacon.SSZObject{Fork: beacon.ForkGloas, Data: data}
}

// On the fixture's data, the block right after the payload's slot builds on
// it, so that block's slot is c.
func TestLocateCommittingSlotNextSlot(t *testing.T) {
	fixture := loadSepoliaFixture(t)
	blocks := &slotBlocks{blocks: map[uint64]*beacon.SSZObject{
		sepoliaTargetSlot: readSepoliaBlock(t, sepoliaTargetSlot),
	}}

	slot, err := locateCommittingSlot(
		context.Background(), blocks, sepoliaPayloadSlot, sepoliaAnchorSlot, fixture.Metadata.BlockHash)
	require.NoError(t, err)
	require.Equal(t, sepoliaTargetSlot, slot)
	require.Equal(t, []uint64{sepoliaTargetSlot}, blocks.fetched)
}

// Missed slots after the payload's slot are skipped.
func TestLocateCommittingSlotSkipsMissedSlots(t *testing.T) {
	blockHash := common.Hash{0xaa}
	blocks := &slotBlocks{blocks: map[uint64]*beacon.SSZObject{
		103: gloasBlock(t, 103, blockHash),
	}}

	slot, err := locateCommittingSlot(context.Background(), blocks, 100, 200, blockHash)
	require.NoError(t, err)
	require.Equal(t, uint64(103), slot)
	require.Equal(t, []uint64{101, 102, 103}, blocks.fetched)
}

// If the first block after the payload's slot does not build on the payload,
// the payload was withheld (EMPTY) and its execution block is not canonical.
// On the fixture's data, block s builds on the parent payload, not on a
// payload at s-1 that the test pretends to look for.
func TestLocateCommittingSlotPayloadNotIncluded(t *testing.T) {
	blocks := &slotBlocks{blocks: map[uint64]*beacon.SSZObject{
		sepoliaPayloadSlot: readSepoliaBlock(t, sepoliaPayloadSlot),
	}}

	_, err := locateCommittingSlot(
		context.Background(), blocks, sepoliaPayloadSlot-1, sepoliaAnchorSlot, common.Hash{0xaa})
	require.ErrorIs(t, err, ErrPayloadNotIncluded)
}

// If no block between the payload's slot and the anchor exists, the payload
// is not committed before the anchor yet; the anchor itself is not searched.
func TestLocateCommittingSlotNotCommittedBeforeAnchor(t *testing.T) {
	blockHash := common.Hash{0xaa}
	blocks := &slotBlocks{blocks: map[uint64]*beacon.SSZObject{
		103: gloasBlock(t, 103, blockHash),
	}}

	_, err := locateCommittingSlot(context.Background(), blocks, 100, 103, blockHash)
	require.ErrorIs(t, err, ErrNotCommittedBeforeAnchor)
	require.Equal(t, []uint64{101, 102}, blocks.fetched)
}

// A pre-Gloas block has no bid to check, so it cannot commit a Gloas payload.
func TestLocateCommittingSlotRejectsPreGloasBlock(t *testing.T) {
	blocks := &slotBlocks{blocks: map[uint64]*beacon.SSZObject{
		101: {Fork: beacon.ForkFulu, Data: []byte{0x01}},
	}}

	_, err := locateCommittingSlot(context.Background(), blocks, 100, 200, common.Hash{0xaa})
	require.ErrorContains(t, err, `beacon block of fork "fulu" has no execution payload bid`)
}

// A failed fetch is not a missed slot and must surface.
func TestLocateCommittingSlotFetchFailure(t *testing.T) {
	blocks := &slotBlocks{errs: map[uint64]error{101: errors.New("node unavailable")}}

	_, err := locateCommittingSlot(context.Background(), blocks, 100, 200, common.Hash{0xaa})
	require.ErrorContains(t, err, "node unavailable")
	require.NotErrorIs(t, err, ErrNotCommittedBeforeAnchor)
}
