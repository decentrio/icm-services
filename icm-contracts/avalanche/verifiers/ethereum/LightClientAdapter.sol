// (c) 2026, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

// SPDX-License-Identifier: LicenseRef-Ecosystem

pragma solidity ^0.8.30;

import {EthereumAdapter} from "./EthereumAdapter.sol";
import {EthereumLogVerifier} from "./EthereumLogVerifier.sol";
import {IBeaconRootSource} from "./IBeaconRootSource.sol";
import {Execution} from "./StateManagerLibrary.sol";

/**
 * THIS IS AN EXAMPLE CONTRACT THAT USES UN-AUDITED CODE.
 * DO NOT USE THIS CODE IN PRODUCTION.
 */

/**
 * @notice Adapts Ethereum log verification for the Teleporter messaging protocol, against the
 * finalized beacon roots of a separate source such as the ZK light client.
 * @dev verifyMessage accepts only logs emitted by this contract's own address on the source chain,
 * so it must be deployed at the same address on both chains (Nick's method), with the same
 * constructor arguments: on the source chain the beacon root source is not used.
 */
contract LightClientAdapter is EthereumAdapter {
    /// @notice Where the finalized beacon roots that messages are proven against come from
    IBeaconRootSource public immutable beaconRootSource;

    constructor(
        uint256 sourceChainId_,
        Execution.BeaconConfig memory beaconConfig_,
        IBeaconRootSource beaconRootSource_
    ) EthereumLogVerifier(sourceChainId_, beaconConfig_) {
        require(address(beaconRootSource_) != address(0), "Invalid beacon root source");
        beaconRootSource = beaconRootSource_;
    }

    function _beaconRoot(
        uint64 slot
    ) internal view override returns (bytes32) {
        return beaconRootSource.beaconRootAt(slot);
    }
}
