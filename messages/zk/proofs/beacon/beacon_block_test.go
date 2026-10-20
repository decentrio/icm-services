// Copyright (C) 2026, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

// THIS IS AN EXAMPLE OF UNAUDITED CODE. DO NOT USE THIS IN PRODUCTION.

package beacon

import (
	"os"
	"testing"

	"github.com/OffchainLabs/go-bitfield"
	"github.com/attestantio/go-eth2-client/spec/altair"
	"github.com/attestantio/go-eth2-client/spec/bellatrix"
	"github.com/attestantio/go-eth2-client/spec/capella"
	"github.com/attestantio/go-eth2-client/spec/deneb"
	"github.com/attestantio/go-eth2-client/spec/electra"
	"github.com/attestantio/go-eth2-client/spec/phase0"
	"github.com/ava-labs/libevm/common"
	"github.com/holiman/uint256"
	"github.com/stretchr/testify/require"
)

const testGIndexBlockStateRoot = 11

// The block tree built from SSZ bytes must hash to the same root attestantio
// computes for the block.
func TestParseBlockTreeMatchesTypedRoot(t *testing.T) {
	signedBlock := minimalSignedBeaconBlock(t, 100, phase0.Root{0xaa})
	expectedRoot, err := signedBlock.Message.HashTreeRoot()
	require.NoError(t, err)

	blockSSZ, err := signedBlock.MarshalSSZ()
	require.NoError(t, err)

	tree, stateRoot, err := ParseBlockTree(&SSZObject{Fork: ForkFulu, Data: blockSSZ})
	require.NoError(t, err)
	require.Equal(t, expectedRoot[:], tree.Hash())
	require.Equal(t, [32]byte(signedBlock.Message.StateRoot), stateRoot)

	leaf, err := tree.Get(testGIndexBlockStateRoot)
	require.NoError(t, err)
	require.Equal(t, signedBlock.Message.StateRoot[:], leaf.Hash())
}

// Real Sepolia Gloas blocks must give the block roots and state roots that the
// Sepolia fixture records (anchorBeaconBlockRoot and its header's state_root,
// and beaconBlocks.c).
func TestParseBlockTreeSepoliaGloasBlocks(t *testing.T) {
	for _, tc := range []struct {
		path      string
		root      string
		stateRoot string
	}{
		{
			path:      "../testdata/sepolia/beacon_block_11397760.ssz",
			root:      "0xb5101888b495fa63efc0a0dc9ca161b2dd02ac18d516d3877e856b9270c1cd67",
			stateRoot: "0xf11a57a39171758f0b38af98fed5d231e1f5239189961b7b9a46d92c157ca157",
		},
		{
			path:      "../testdata/sepolia/beacon_block_11397715.ssz",
			root:      "0x9b11a47e73a8605a73bd41e18ee5001d5aee4fa4c3b770b476b0414fff52bf28",
			stateRoot: "0x09626e51cae8c681dc7a8c116df21147bab677d9e9d0ffbd07ac633d8085753b",
		},
	} {
		blockSSZ, err := os.ReadFile(tc.path)
		require.NoError(t, err)

		tree, stateRoot, err := ParseBlockTree(&SSZObject{Fork: ForkGloas, Data: blockSSZ})
		require.NoError(t, err, tc.path)
		require.Equal(t, common.HexToHash(tc.root).Bytes(), tree.Hash(), tc.path)
		require.Equal(t, common.HexToHash(tc.stateRoot), common.Hash(stateRoot), tc.path)

		leaf, err := tree.Get(testGIndexBlockStateRoot)
		require.NoError(t, err)
		require.Equal(t, stateRoot[:], leaf.Hash(), tc.path)
	}
}

// Malformed bytes must be rejected at deserialization, for every fork.
func TestParseBlockTreeRejectsGarbage(t *testing.T) {
	for _, fork := range []Fork{ForkFulu, ForkGloas} {
		_, _, err := ParseBlockTree(&SSZObject{Fork: fork, Data: []byte{0x01, 0x02, 0x03}})
		require.ErrorContains(t, err, "failed to deserialize beacon block", "fork %s", fork)
	}
}

// A block of a fork this package cannot decode must be rejected.
func TestParseBlockTreeRejectsUnknownFork(t *testing.T) {
	_, _, err := ParseBlockTree(&SSZObject{Fork: "electra", Data: []byte{0x01}})
	require.ErrorContains(t, err, `unsupported consensus fork "electra"`)
}

// minimalSignedBeaconBlock builds the smallest signed block attestantio will
// serialize and hash, at the given slot and with the given state_root.
// TODO: This is duplicated in the proofs package's tests pending package consolidation.
// See Issue: https://github.com/ava-labs/icm-services/issues/1541
func minimalSignedBeaconBlock(t *testing.T, slot uint64, stateRoot phase0.Root) *electra.SignedBeaconBlock {
	t.Helper()
	return &electra.SignedBeaconBlock{
		Message: &electra.BeaconBlock{
			Slot:          phase0.Slot(slot),
			ProposerIndex: 7,
			ParentRoot:    phase0.Root{0x01},
			StateRoot:     stateRoot,
			Body: &electra.BeaconBlockBody{
				ETH1Data: &phase0.ETH1Data{
					DepositRoot: phase0.Root{0x02},
					BlockHash:   make([]byte, 32),
				},
				ProposerSlashings: []*phase0.ProposerSlashing{},
				AttesterSlashings: []*electra.AttesterSlashing{},
				Attestations:      []*electra.Attestation{},
				Deposits:          []*phase0.Deposit{},
				VoluntaryExits:    []*phase0.SignedVoluntaryExit{},
				SyncAggregate: &altair.SyncAggregate{
					SyncCommitteeBits: bitfield.NewBitvector512(),
				},
				ExecutionPayload: &deneb.ExecutionPayload{
					BaseFeePerGas: uint256.NewInt(1),
					Transactions:  []bellatrix.Transaction{},
					Withdrawals:   []*capella.Withdrawal{},
				},
				BLSToExecutionChanges: []*capella.SignedBLSToExecutionChange{},
				BlobKZGCommitments:    []deneb.KZGCommitment{},
				ExecutionRequests: &electra.ExecutionRequests{
					Deposits:       []*electra.DepositRequest{},
					Withdrawals:    []*electra.WithdrawalRequest{},
					Consolidations: []*electra.ConsolidationRequest{},
				},
			},
		},
	}
}
