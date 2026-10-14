// (c) 2026, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

// SPDX-License-Identifier: LicenseRef-Ecosystem

pragma solidity 0.8.30;

import {TeleporterRoutedAppUpgradeable} from "./TeleporterRoutedAppUpgradeable.sol";
import {OwnableUpgradeable} from
    "@openzeppelin/contracts-upgradeable@5.1.0/access/OwnableUpgradeable.sol";

/**
 * THIS IS AN EXAMPLE CONTRACT THAT USES UN-AUDITED CODE.
 * DO NOT USE THIS CODE IN PRODUCTION.
 */

/**
 * @dev Contract that inherits {TeleporterRoutedAppUpgradeable} and allows only the owner of the
 * contract to change messenger routes or pause and unpause messengers.
 *
 * @custom:security-contact https://github.com/ava-labs/icm-contracts/blob/main/SECURITY.md
 */
abstract contract TeleporterRoutedOwnableAppUpgradeable is
    TeleporterRoutedAppUpgradeable,
    OwnableUpgradeable
{
    // solhint-disable-next-line func-name-mixedcase
    function __TeleporterRoutedOwnableApp_init(
        bytes32 blockchainID,
        address initialOwner
    ) internal onlyInitializing {
        __TeleporterRoutedApp_init(blockchainID);
        __Ownable_init(initialOwner);
    }

    // solhint-disable-next-line func-name-mixedcase, no-empty-blocks
    function __TeleporterRoutedOwnableApp_init_unchained() internal onlyInitializing {}

    /**
     * @dev See {TeleporterRoutedAppUpgradeable-_checkTeleporterRoutedAppAccess}
     *
     * Checks that the caller is the owner of the contract.
     */
    function _checkTeleporterRoutedAppAccess() internal view virtual override {
        _checkOwner();
    }
}
