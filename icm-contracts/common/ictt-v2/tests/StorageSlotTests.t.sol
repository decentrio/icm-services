// (c) 2026, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

// SPDX-License-Identifier: LicenseRef-Ecosystem

pragma solidity 0.8.30;

import {Test} from "@forge-std/Test.sol";
import {ERC20TokenHomeUpgradeable} from "../TokenHome/ERC20TokenHomeUpgradeable.sol";
import {ERC20TokenRemoteUpgradeable} from "../TokenRemote/ERC20TokenRemoteUpgradeable.sol";
import {NativeTokenHomeUpgradeable} from "../TokenHome/NativeTokenHomeUpgradeable.sol";
import {ICMInitializable} from "@utilities/ICMInitializable.sol";

/**
 * @dev ICTT v2 keeps the v1 storage namespaces, so these slots must stay equal to v1's.
 */
contract StorageSlotTest is Test {
    function testERC20TokenHomeStorageSlot() public {
        assertEq(
            new ERC20TokenHomeUpgradeable(ICMInitializable.Disallowed)
                .ERC20_TOKEN_HOME_STORAGE_LOCATION(),
            _erc7201StorageSlot("ERC20TokenHome")
        );
    }

    function testERC20TokenRemoteStorageSlot() public {
        assertEq(
            new ERC20TokenRemoteUpgradeable(ICMInitializable.Disallowed)
                .ERC20_TOKEN_REMOTE_STORAGE_LOCATION(),
            _erc7201StorageSlot("ERC20TokenRemote")
        );
    }

    function testNativeTokenHomeStorageSlot() public {
        assertEq(
            new NativeTokenHomeUpgradeable(ICMInitializable.Disallowed)
                .NATIVE_TOKEN_HOME_STORAGE_LOCATION(),
            _erc7201StorageSlot("NativeTokenHome")
        );
    }

    function testTokenHomeStorageSlot() public {
        assertEq(
            new ERC20TokenHomeUpgradeable(ICMInitializable.Disallowed).TOKEN_HOME_STORAGE_LOCATION(),
            _erc7201StorageSlot("TokenHome")
        );
    }

    function testTokenRemoteStorageSlot() public {
        assertEq(
            new ERC20TokenRemoteUpgradeable(ICMInitializable.Disallowed)
                .TOKEN_REMOTE_STORAGE_LOCATION(),
            _erc7201StorageSlot("TokenRemote")
        );
    }

    function _erc7201StorageSlot(
        bytes memory storageName
    ) private pure returns (bytes32) {
        return keccak256(
            abi.encode(
                uint256(keccak256(abi.encodePacked("avalanche-ictt.storage.", storageName))) - 1
            )
        ) & ~bytes32(uint256(0xff));
    }
}
