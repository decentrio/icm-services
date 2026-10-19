// (c) 2026, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

// SPDX-License-Identifier: LicenseRef-Ecosystem

pragma solidity ^0.8.30;

import {ZKStateManager} from "./ZKStateManager.sol";
import {EthereumAdapter} from "./EthereumAdapter.sol";
import {Consensus, Execution} from "./StateManagerLibrary.sol";

/**
 * THIS IS AN EXAMPLE CONTRACT THAT USES UN-AUDITED CODE.
 * DO NOT USE THIS CODE IN PRODUCTION.
 */

/**
 * @notice Adapts ZKStateManager for the Teleporter messaging protocol: messages are verified
 * against the beacon roots this contract's own Boundless transitions confirm.
 */
contract ZKAdapter is ZKStateManager, EthereumAdapter {
    constructor(
        uint256 sourceChainId_,
        Consensus.State memory startingState,
        Execution.BeaconConfig memory beaconConfig_,
        uint256 genesisTime_,
        uint24 permissibleTimespan_,
        address verifier,
        bytes32 imageID_,
        address admin,
        address superAdmin
    )
        ZKStateManager(
            sourceChainId_,
            startingState,
            beaconConfig_,
            genesisTime_,
            permissibleTimespan_,
            verifier,
            imageID_,
            admin,
            superAdmin
        )
    {}
}
