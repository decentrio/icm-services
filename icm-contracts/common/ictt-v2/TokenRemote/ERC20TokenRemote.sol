// (c) 2026, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

// SPDX-License-Identifier: LicenseRef-Ecosystem

pragma solidity 0.8.30;

import {ERC20TokenRemoteUpgradeable} from "./ERC20TokenRemoteUpgradeable.sol";
import {TokenRemoteSettingsV2} from "./interfaces/TokenRemoteSettingsV2.sol";
import {ICMInitializable} from "@utilities/ICMInitializable.sol";

/**
 * THIS IS AN EXAMPLE CONTRACT THAT USES UN-AUDITED CODE.
 * DO NOT USE THIS CODE IN PRODUCTION.
 */

/**
 * @title ERC20TokenRemote
 * @notice A non-upgradeable version of {ERC20TokenRemoteUpgradeable} that calls the parent upgradeable contract's initialize function.
 * @custom:security-contact https://github.com/ava-labs/icm-contracts/blob/main/SECURITY.md
 */
contract ERC20TokenRemote is ERC20TokenRemoteUpgradeable {
    constructor(
        TokenRemoteSettingsV2 memory settings,
        string memory tokenName,
        string memory tokenSymbol,
        uint8 tokenDecimals
    ) ERC20TokenRemoteUpgradeable(ICMInitializable.Allowed) {
        initialize(settings, tokenName, tokenSymbol, tokenDecimals);
    }
}
