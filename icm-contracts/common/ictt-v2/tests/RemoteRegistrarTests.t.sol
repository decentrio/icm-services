// (c) 2026, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

// SPDX-License-Identifier: LicenseRef-Ecosystem

pragma solidity 0.8.30;

import {Test} from "forge-std/Test.sol";
import {ERC20TokenHome} from "../TokenHome/ERC20TokenHome.sol";
import {RemoteRegistrar} from "../RemoteRegistrar.sol";
import {
    ITokenHome,
    RemoteTokenTransferrerSettings
} from "../../../avalanche/ictt/TokenHome/interfaces/ITokenHome.sol";
import {ExampleERC20} from "@mocks/ExampleERC20.sol";

contract RemoteRegistrarTest is Test {
    bytes32 internal constant _HOME_CHAIN = bytes32(uint256(11_155_111));
    bytes32 internal constant _REMOTE_CHAIN =
        0x7fc93d85c6d62c5b2ac0b519c87010ea5294012d1e407030d6acd0021cac10d5;
    address internal constant _REMOTE = address(0xFEED);
    uint8 internal constant _DECIMALS = 18;

    ERC20TokenHome internal _home;
    RemoteRegistrar internal _registrar;

    function setUp() public {
        _home =
            new ERC20TokenHome(_HOME_CHAIN, address(this), address(new ExampleERC20()), _DECIMALS);
        _registrar = new RemoteRegistrar(_home, _HOME_CHAIN);
    }

    /// The home's owner adds the registrar as a messenger for the remote's chain, has it register
    /// the remote, and removes it.
    function testRegistersARemoteForTheHomesOwner() public {
        _home.addMessenger(_REMOTE_CHAIN, address(_registrar));

        vm.expectEmit(true, true, true, true, address(_home));
        emit ITokenHome.RemoteRegistered(_REMOTE_CHAIN, _REMOTE, 0, _DECIMALS);
        _registrar.register(_REMOTE_CHAIN, _REMOTE, _DECIMALS, _DECIMALS);
        _home.removeMessenger(_REMOTE_CHAIN, address(_registrar));

        RemoteTokenTransferrerSettings memory settings =
            _home.getRemoteTokenTransferrerSettings(_REMOTE_CHAIN, _REMOTE);
        assertTrue(settings.registered);
        assertEq(settings.collateralNeeded, 0);
        assertEq(settings.tokenMultiplier, 1);
    }

    /// Only whoever deployed the registrar may use it while it is a messenger.
    function testOnlyItsDeployerRegisters() public {
        _home.addMessenger(_REMOTE_CHAIN, address(_registrar));

        vm.prank(address(0xBAD));
        vm.expectRevert("RemoteRegistrar: not owner");
        _registrar.register(_REMOTE_CHAIN, _REMOTE, _DECIMALS, _DECIMALS);
    }

    /// The home accepts the registration only while its owner has added the registrar.
    function testRegistersNothingUnlessTheHomeAddedIt() public {
        vm.expectRevert("TeleporterRoutedApp: unauthorized messenger");
        _registrar.register(_REMOTE_CHAIN, _REMOTE, _DECIMALS, _DECIMALS);

        _home.addMessenger(_REMOTE_CHAIN, address(_registrar));
        _home.removeMessenger(_REMOTE_CHAIN, address(_registrar));
        vm.expectRevert("TeleporterRoutedApp: unauthorized messenger");
        _registrar.register(_REMOTE_CHAIN, _REMOTE, _DECIMALS, _DECIMALS);
    }

    /// The home's rules still apply: a remote with other decimals than the token's is refused.
    function testKeepsTheHomesRegistrationRules() public {
        _home.addMessenger(_REMOTE_CHAIN, address(_registrar));

        vm.expectRevert("TokenHome: invalid home token decimals");
        _registrar.register(_REMOTE_CHAIN, _REMOTE, _DECIMALS - 1, _DECIMALS);
    }

    /// The home checks that a messenger it adds is on its own chain.
    function testReportsTheHomesChain() public view {
        assertEq(_registrar.blockchainID(), _HOME_CHAIN);
    }
}
