// Copyright (C) 2026, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

// THIS IS AN EXAMPLE OF UNAUDITED CODE. DO NOT USE THIS IN PRODUCTION.

package beacon

import (
	ssz "github.com/ferranbt/fastssz"
)

// Progressive merkleization (EIP-7916), which Gloas uses for the BeaconState
// container (EIP-7688) and for its unbounded lists:
//
//	merkleize_progressive(chunks, n=1) =
//	    hash(merkleize(chunks[:n], n), merkleize_progressive(chunks[n:], 4n))
//
// Each spine level holds a batch of 1, 4, 16, 64, ... chunks as its left
// subtree and the rest of the spine as its right; an empty rest is a zero
// chunk. Unlike a bounded list, the tree's shape depends only on how many
// chunks it holds, so it needs no limit.
//
// https://github.com/ethereum/EIPs/blob/master/EIPS/eip-7916.md

// merkleizeProgressive returns the progressive tree root of chunks.
func merkleizeProgressive(chunks [][32]byte) [32]byte {
	root := [32]byte{}
	// Fold the spine from its deepest batch up.
	batches := progressiveBatches(len(chunks))
	for i := len(batches) - 1; i >= 0; i-- {
		start, size := batches[i].start, batches[i].size
		end := min(start+size, len(chunks))
		root = hashPair(merkleizeToLimit(chunks[start:end], uint64(size)), root)
	}
	return root
}

// progressiveListRoot merkleizes a progressive list of chunks, mixing in the
// element count. Every progressive list in the Gloas BeaconState holds one
// element per chunk except the packed basic-type lists, which mix in their
// own element count.
func progressiveListRoot(chunks [][32]byte) [32]byte {
	return mixinLength(merkleizeProgressive(chunks), uint64(len(chunks)))
}

// progressiveTree builds the progressive tree over leaves as fastssz nodes,
// so proofs can be taken from it. Leaves may be expanded subtrees.
func progressiveTree(leaves []*ssz.Node) (*ssz.Node, error) {
	spine := ssz.EmptyLeaf()
	batches := progressiveBatches(len(leaves))
	for i := len(batches) - 1; i >= 0; i-- {
		start, size := batches[i].start, batches[i].size
		end := min(start+size, len(leaves))
		batch, err := ssz.TreeFromNodes(leaves[start:end], size)
		if err != nil {
			return nil, err
		}
		spine = ssz.NewNodeWithLR(batch, spine)
	}
	return spine, nil
}

// progressiveBatch locates one spine level's left subtree: it holds up to
// size chunks, starting at chunk start.
type progressiveBatch struct {
	start, size int
}

// progressiveBatches lists the spine levels needed for n chunks, top first.
func progressiveBatches(n int) []progressiveBatch {
	var batches []progressiveBatch
	for start, size := 0, 1; start < n; start, size = start+size, size*4 {
		batches = append(batches, progressiveBatch{start: start, size: size})
	}
	return batches
}
