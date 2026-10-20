// Copyright (C) 2026, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

// THIS IS AN EXAMPLE OF UNAUDITED CODE. DO NOT USE THIS IN PRODUCTION.

package proofs

import (
	"context"
	"errors"
	"fmt"

	"github.com/ava-labs/icm-services/messages/zk/proofs/beacon"
	"github.com/ava-labs/libevm/common"
)

var (
	// ErrPayloadNotIncluded means the first beacon block after a Gloas payload's
	// slot did not build on the payload: the payload was withheld (EMPTY), so
	// its execution block is not canonical. This is final for the payload.
	ErrPayloadNotIncluded = errors.New("execution payload was not included")

	// ErrNotCommittedBeforeAnchor means no beacon block between a Gloas
	// payload's slot and the anchor has built on the payload yet. A later
	// anchor will have one.
	ErrNotCommittedBeforeAnchor = errors.New("execution payload is not committed before the anchor slot")
)

// beaconBlockSource fetches SSZ encoded beacon blocks by slot.
type beaconBlockSource interface {
	Block(ctx context.Context, slot uint64) (*beacon.SSZObject, error)
}

// locateCommittingSlot returns c, the slot of the beacon block that marked
// the Gloas payload at payloadSlot FULL, which must precede anchorSlot.
//
// Under Gloas, a beacon block carries only a bid for its slot's payload, and
// it is the next block that decides whether that payload was included, by
// building on it or not. So the beacon state holding the payload's block hash
// as latest_block_hash is the state at c, the first block after payloadSlot,
// not the state at payloadSlot, which still holds the parent's hash.
func locateCommittingSlot(
	ctx context.Context,
	client beaconBlockSource,
	payloadSlot uint64,
	anchorSlot uint64,
	blockHash common.Hash,
) (uint64, error) {
	for slot := payloadSlot + 1; slot < anchorSlot; slot++ {
		block, err := client.Block(ctx, slot)
		if errors.Is(err, beacon.ErrNoBlockAtSlot) {
			continue
		}
		if err != nil {
			return 0, fmt.Errorf("failed to fetch beacon block at slot %d: %w", slot, err)
		}
		parentBlockHash, err := beacon.BidParentBlockHash(block)
		if err != nil {
			return 0, fmt.Errorf("beacon block at slot %d: %w", slot, err)
		}
		if parentBlockHash != blockHash {
			return 0, fmt.Errorf("%w: block %s at slot %d, but the block at slot %d builds on %s",
				ErrPayloadNotIncluded, blockHash, payloadSlot, slot, parentBlockHash)
		}
		return slot, nil
	}
	return 0, fmt.Errorf("%w: block %s at slot %d, anchor slot %d",
		ErrNotCommittedBeforeAnchor, blockHash, payloadSlot, anchorSlot)
}
