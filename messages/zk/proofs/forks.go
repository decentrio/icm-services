// Copyright (C) 2026, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

// THIS IS AN EXAMPLE OF UNAUDITED CODE. DO NOT USE THIS IN PRODUCTION.

package proofs

import "math"

// ForkConfig is the part of a network's fork schedule that proof building
// depends on. Gloas moves the receipts root out of the beacon state, so its
// payloads take a different proof path. The contract's config must hold the
// same fork slot, otherwise proofs built here revert on-chain.
type ForkConfig struct {
	// GloasForkSlot is the first slot of the Gloas fork: GLOAS_FORK_EPOCH *
	// SLOTS_PER_EPOCH, or math.MaxUint64 while Gloas is not scheduled.
	GloasForkSlot uint64
}

// Fork configs of the public networks, from the eth-clients network configs.
var (
	MainnetForkConfig = ForkConfig{GloasForkSlot: math.MaxUint64}
	SepoliaForkConfig = ForkConfig{GloasForkSlot: 353024 * 32}
	HoodiForkConfig   = ForkConfig{GloasForkSlot: 132352 * 32}
)

// isGloas reports whether slot is in the Gloas fork.
func (c ForkConfig) isGloas(slot uint64) bool {
	return slot >= c.GloasForkSlot
}
