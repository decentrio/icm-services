// (c) 2026, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

// SPDX-License-Identifier: LicenseRef-Ecosystem

pragma solidity 0.8.30;

import {ERC20TokenHomeUpgradeable} from "./ERC20TokenHomeUpgradeable.sol";
import {ICMInitializable} from "@utilities/ICMInitializable.sol";

/**
 * THIS IS AN EXAMPLE CONTRACT THAT USES UN-AUDITED CODE.
 * DO NOT USE THIS CODE IN PRODUCTION.
 */

/**
 * @title ERC20TokenHome
 * @notice A non-upgradeable version of {ERC20TokenHomeUpgradeable} that calls the parent upgradeable contract's initialize function.
 * @custom:security-contact https://github.com/ava-labs/icm-contracts/blob/main/SECURITY.md
 */
contract ERC20TokenHome is ERC20TokenHomeUpgradeable {
    constructor(
        bytes32 blockchainID,
        address teleporterManager,
        address tokenAddress,
        uint8 tokenDecimals
    ) ERC20TokenHomeUpgradeable(ICMInitializable.Allowed) {
        initialize(blockchainID, teleporterManager, tokenAddress, tokenDecimals);
    }
}
