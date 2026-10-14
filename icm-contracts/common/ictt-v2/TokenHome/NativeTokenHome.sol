// (c) 2026, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

// SPDX-License-Identifier: LicenseRef-Ecosystem

pragma solidity 0.8.30;

import {NativeTokenHomeUpgradeable} from "./NativeTokenHomeUpgradeable.sol";
import {ICMInitializable} from "@utilities/ICMInitializable.sol";

/**
 * @title NativeTokenHome
 * @notice A non-upgradeable version of {NativeTokenHomeUpgradeable} that calls the parent upgradeable contract's initialize function.
 * @custom:security-contact https://github.com/ava-labs/icm-contracts/blob/main/SECURITY.md
 */
contract NativeTokenHome is NativeTokenHomeUpgradeable {
    constructor(
        bytes32 blockchainID,
        address teleporterManager,
        address wrappedTokenAddress
    ) NativeTokenHomeUpgradeable(ICMInitializable.Allowed) {
        initialize(blockchainID, teleporterManager, wrappedTokenAddress);
    }
}
