// (c) 2026, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

// SPDX-License-Identifier: LicenseRef-Ecosystem

pragma solidity ^0.8.30;

/**
 * @notice A source of finalized Ethereum beacon block roots, such as a light client.
 */
interface IBeaconRootSource {
    /**
     * @notice The finalized beacon block root at `slot`, or zero if the source has none.
     */
    function beaconRootAt(
        uint64 slot
    ) external view returns (bytes32);
}
