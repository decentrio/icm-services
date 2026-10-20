// Copyright (C) 2026, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

// THIS IS AN EXAMPLE OF UNAUDITED CODE. DO NOT USE THIS IN PRODUCTION.

package beacon

import (
	"fmt"

	electra "github.com/attestantio/go-eth2-client/spec/electra"
	"github.com/attestantio/go-eth2-client/spec/gloas"
	"github.com/attestantio/go-eth2-client/spec/phase0"
	"github.com/ava-labs/libevm/common"
	ssz "github.com/ferranbt/fastssz"
)

// numBlockFieldLeaves pads the BeaconBlock's 5 fields to 2^3 = 8 leaves.
// BeaconBlock is still this phase0 container in Gloas.
// https://github.com/ethereum/consensus-specs/blob/master/specs/phase0/beacon-chain.md#beaconblock
const numBlockFieldLeaves = 8

// ParseBlockTree deserializes an SSZ-encoded SignedBeaconBlock of the
// object's fork and builds the Merkle tree of its BeaconBlock. Note that a
// "lite tree" (a memory-efficient Merkle tree) is not needed here because we
// are not expanding any of the anchor beacon block's subtrees.
//
// Note the block's state_root field, which is the Merkle root of the anchor beacon state,
// is also returned.
func ParseBlockTree(block *SSZObject) (*ssz.Node, [32]byte, error) {
	header, err := parseBlockHeader(block)
	if err != nil {
		return nil, [32]byte{}, err
	}

	fieldRoots := make([][32]byte, numBlockFieldLeaves)
	fieldRoots[0] = uint64Root(uint64(header.Slot))
	fieldRoots[1] = uint64Root(uint64(header.ProposerIndex))
	fieldRoots[2] = header.ParentRoot
	fieldRoots[3] = header.StateRoot
	fieldRoots[4] = header.BodyRoot

	tree, err := ssz.TreeFromChunks(toByteSlices(fieldRoots))
	if err != nil {
		return nil, [32]byte{}, fmt.Errorf("failed to build block tree: %w", err)
	}
	return tree, header.StateRoot, nil
}

// parseBlockHeader deserializes a SignedBeaconBlock of the object's fork and
// reduces its BeaconBlock to a header: the same fields with the body replaced
// by its root, which leaves the root unchanged.
func parseBlockHeader(block *SSZObject) (*phase0.BeaconBlockHeader, error) {
	switch block.Fork {
	case ForkFulu:
		var signedBlock electra.SignedBeaconBlock
		if err := signedBlock.UnmarshalSSZ(block.Data); err != nil {
			return nil, fmt.Errorf("failed to deserialize beacon block: %w", err)
		}
		b := signedBlock.Message
		return blockHeader(b.Slot, b.ProposerIndex, b.ParentRoot, b.StateRoot, b.Body)
	case ForkGloas:
		var signedBlock gloas.SignedBeaconBlock
		if err := signedBlock.UnmarshalSSZ(block.Data); err != nil {
			return nil, fmt.Errorf("failed to deserialize beacon block: %w", err)
		}
		b := signedBlock.Message
		return blockHeader(b.Slot, b.ProposerIndex, b.ParentRoot, b.StateRoot, b.Body)
	default:
		return nil, fmt.Errorf("unsupported consensus fork %q for beacon block", block.Fork)
	}
}

// BidParentBlockHash deserializes a Gloas SignedBeaconBlock and returns the
// parent_block_hash of its execution payload bid: the execution block the
// block builds on. A block whose bid builds on a payload marks that payload
// FULL, and its post-state's latest_block_hash is that payload's block hash.
func BidParentBlockHash(block *SSZObject) (common.Hash, error) {
	if block.Fork != ForkGloas {
		return common.Hash{}, fmt.Errorf("beacon block of fork %q has no execution payload bid", block.Fork)
	}
	var signedBlock gloas.SignedBeaconBlock
	if err := signedBlock.UnmarshalSSZ(block.Data); err != nil {
		return common.Hash{}, fmt.Errorf("failed to deserialize beacon block: %w", err)
	}
	return common.Hash(signedBlock.Message.Body.SignedExecutionPayloadBid.Message.ParentBlockHash), nil
}

func blockHeader(
	slot phase0.Slot,
	proposerIndex phase0.ValidatorIndex,
	parentRoot phase0.Root,
	stateRoot phase0.Root,
	body interface{ HashTreeRoot() ([32]byte, error) },
) (*phase0.BeaconBlockHeader, error) {
	bodyRoot, err := body.HashTreeRoot()
	if err != nil {
		return nil, fmt.Errorf("failed to hash block body: %w", err)
	}
	return &phase0.BeaconBlockHeader{
		Slot:          slot,
		ProposerIndex: proposerIndex,
		ParentRoot:    parentRoot,
		StateRoot:     stateRoot,
		BodyRoot:      bodyRoot,
	}, nil
}
