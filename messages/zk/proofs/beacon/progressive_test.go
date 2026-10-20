// Copyright (C) 2026, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

// THIS IS AN EXAMPLE OF UNAUDITED CODE. DO NOT USE THIS IN PRODUCTION.

package beacon

import (
	"testing"

	"github.com/attestantio/go-eth2-client/spec/phase0"
	ssz "github.com/ferranbt/fastssz"
	dynssz "github.com/pk910/dynamic-ssz"
	"github.com/stretchr/testify/require"
)

// progressiveRoots is a container holding one progressive list of roots. A
// one-field container's root is its field's root, so hashing it with
// dynamic-ssz, the SSZ engine behind go-eth2-client's Gloas types, gives the
// list's root.
type progressiveRoots struct {
	Roots []phase0.Root `ssz-type:"progressive-list"`
}

// Both progressive merkleization helpers must agree with dynamic-ssz at every
// spine-batch boundary. Batches hold 1, 4, 16 and 64 chunks, so lists of 0,
// 1, 2, 5, 21 and 85 chunks each end in a different batch.
func TestProgressiveMerkleizationMatchesDynamicSSZ(t *testing.T) {
	for _, n := range []int{0, 1, 2, 5, 21, 85} {
		roots := make([]phase0.Root, n)
		chunks := make([][32]byte, n)
		leaves := make([]*ssz.Node, n)
		for i := range roots {
			roots[i][0] = byte(i + 1)
			roots[i][31] = 0xaa
			chunks[i] = roots[i]
			leaves[i] = ssz.LeafFromBytes(roots[i][:])
		}
		expected, err := dynssz.GetGlobalDynSsz().HashTreeRoot(&progressiveRoots{Roots: roots})
		require.NoError(t, err)

		require.Equal(t, expected, progressiveListRoot(chunks), "%d chunks", n)

		tree, err := progressiveTree(leaves)
		require.NoError(t, err)
		require.Equal(t, expected, mixinLength([32]byte(tree.Hash()), uint64(n)), "%d chunks", n)
	}
}
