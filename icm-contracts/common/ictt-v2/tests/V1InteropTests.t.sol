// (c) 2026, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

// SPDX-License-Identifier: LicenseRef-Ecosystem

pragma solidity 0.8.30;

import {Test} from "@forge-std/Test.sol";
import {ITeleporterSender} from "../../ITeleporterRoutedApp.sol";
import {ERC20TokenHome} from "../TokenHome/ERC20TokenHome.sol";
import {NativeTokenRemoteUpgradeable} from
    "../../../avalanche/ictt/TokenRemote/NativeTokenRemoteUpgradeable.sol";
import {TokenRemoteSettings} from "../../../avalanche/ictt/TokenRemote/interfaces/ITokenRemote.sol";
import {SendTokensInput} from "../../../avalanche/ictt/interfaces/ITokenTransferrer.sol";
import {RemoteTokenTransferrerSettings} from
    "../../../avalanche/ictt/TokenHome/interfaces/ITokenHome.sol";
import {TeleporterRegistry} from "@teleporter/registry/TeleporterRegistry.sol";
import {TeleporterMessageInput, TeleporterFeeInfo} from "@teleporter/ITeleporterMessenger.sol";
import {IWarpMessenger} from "@subnet-evm/IWarpMessenger.sol";
import {INativeMinter} from "@subnet-evm/INativeMinter.sol";
import {ICMInitializable} from "@utilities/ICMInitializable.sol";
import {ExampleERC20} from "@mocks/ExampleERC20.sol";

/**
 * @dev Messenger stand-in that keeps the last message sent through it, so a test can deliver it.
 */
contract RecordingMessenger is ITeleporterSender {
    bytes32 public blockchainID;
    TeleporterMessageInput private _lastMessageInput;

    constructor(
        bytes32 blockchainID_
    ) {
        blockchainID = blockchainID_;
    }

    function sendCrossChainMessage(
        TeleporterMessageInput calldata messageInput
    ) external returns (bytes32) {
        _lastMessageInput = messageInput;
        return keccak256(abi.encode(messageInput));
    }

    function lastMessageInput() external view returns (TeleporterMessageInput memory) {
        return _lastMessageInput;
    }
}

/**
 * @dev An ICTT v2 home and an ICTT v1 NativeTokenRemote exchange the same messages, so a v1 remote
 * on an Avalanche L1 can keep working with a v2 home. Each side runs on its own messenger: the
 * v2 home through its route to the remote's chain, and the v1 remote through its TeleporterRegistry.
 */
contract V1NativeTokenRemoteInteropTest is Test {
    bytes32 private constant _HOME_BLOCKCHAIN_ID = keccak256("home");
    bytes32 private constant _REMOTE_BLOCKCHAIN_ID = keccak256("remote");
    address private constant _WARP_PRECOMPILE_ADDRESS = 0x0200000000000000000000000000000000000005;
    address private constant _NATIVE_MINTER_PRECOMPILE_ADDRESS =
        0x0200000000000000000000000000000000000001;
    address private constant _REGISTRY_ADDRESS = address(0x5E9);
    address private constant _ALICE = address(0xA11CE);
    address private constant _BOB = address(0xB0B);
    uint256 private constant _INITIAL_RESERVE_IMBALANCE = 100e18;

    ExampleERC20 private _token;
    ERC20TokenHome private _home;
    NativeTokenRemoteUpgradeable private _remote;
    RecordingMessenger private _homeMessenger;
    RecordingMessenger private _remoteMessenger;

    function setUp() public {
        _token = new ExampleERC20();
        _home = new ERC20TokenHome(_HOME_BLOCKCHAIN_ID, address(this), address(_token), 18);
        _homeMessenger = new RecordingMessenger(_HOME_BLOCKCHAIN_ID);
        _home.addMessenger(_REMOTE_BLOCKCHAIN_ID, address(_homeMessenger));
        _home.setOutboundMessenger(_REMOTE_BLOCKCHAIN_ID, address(_homeMessenger));

        // The v1 remote reads its blockchain ID from Warp and its messenger from a TeleporterRegistry.
        _remoteMessenger = new RecordingMessenger(_REMOTE_BLOCKCHAIN_ID);
        vm.mockCall(
            _WARP_PRECOMPILE_ADDRESS,
            abi.encodeCall(IWarpMessenger.getBlockchainID, ()),
            abi.encode(_REMOTE_BLOCKCHAIN_ID)
        );
        vm.mockCall(
            _REGISTRY_ADDRESS,
            abi.encodeWithSelector(TeleporterRegistry(_REGISTRY_ADDRESS).latestVersion.selector),
            abi.encode(1)
        );
        vm.mockCall(
            _REGISTRY_ADDRESS,
            abi.encodeWithSelector(TeleporterRegistry.getLatestTeleporter.selector),
            abi.encode(address(_remoteMessenger))
        );
        vm.mockCall(
            _REGISTRY_ADDRESS,
            abi.encodeWithSelector(TeleporterRegistry.getVersionFromAddress.selector),
            abi.encode(1)
        );

        _remote = new NativeTokenRemoteUpgradeable(ICMInitializable.Allowed);
        _remote.initialize({
            settings: TokenRemoteSettings({
                teleporterRegistryAddress: _REGISTRY_ADDRESS,
                teleporterManager: address(this),
                minTeleporterVersion: 1,
                tokenHomeBlockchainID: _HOME_BLOCKCHAIN_ID,
                tokenHomeAddress: address(_home),
                tokenHomeDecimals: 18
            }),
            nativeAssetSymbol: "NAT",
            initialReserveImbalance: _INITIAL_RESERVE_IMBALANCE,
            burnedFeesReportingRewardPercentage: 1
        });
    }

    function testRoundTrip() public {
        // The v1 remote registers with the v2 home.
        _remote.registerWithHome(TeleporterFeeInfo({feeTokenAddress: address(0), amount: 0}));
        _deliverToHome();
        RemoteTokenTransferrerSettings memory settings =
            _home.getRemoteTokenTransferrerSettings(_REMOTE_BLOCKCHAIN_ID, address(_remote));
        assertTrue(settings.registered);
        assertEq(settings.collateralNeeded, _INITIAL_RESERVE_IMBALANCE);

        // Collateralize the remote, then send tokens from the v2 home to the v1 remote.
        _token.approve(address(_home), _INITIAL_RESERVE_IMBALANCE + 10e18);
        _home.addCollateral(_REMOTE_BLOCKCHAIN_ID, address(_remote), _INITIAL_RESERVE_IMBALANCE);
        // solhint-disable-next-line check-send-result
        _home.send(_sendTokensInput(_REMOTE_BLOCKCHAIN_ID, address(_remote), _ALICE), 10e18);
        vm.expectCall(
            _NATIVE_MINTER_PRECOMPILE_ADDRESS,
            abi.encodeCall(INativeMinter.mintNativeCoin, (_ALICE, 10e18))
        );
        vm.mockCall(
            _NATIVE_MINTER_PRECOMPILE_ADDRESS,
            abi.encodeCall(INativeMinter.mintNativeCoin, (_ALICE, 10e18)),
            new bytes(0)
        );
        _deliverToRemote();
        assertTrue(_remote.getIsCollateralized());

        // Send part of it back from the v1 remote to the v2 home.
        vm.deal(_ALICE, 10e18);
        vm.prank(_ALICE);
        // solhint-disable-next-line check-send-result, multiple-sends
        _remote.send{value: 4e18}(_sendTokensInput(_HOME_BLOCKCHAIN_ID, address(_home), _BOB));
        _deliverToHome();
        assertEq(_token.balanceOf(_BOB), 4e18);
    }

    function _deliverToHome() private {
        TeleporterMessageInput memory input = _remoteMessenger.lastMessageInput();
        assertEq(input.destinationBlockchainID, _HOME_BLOCKCHAIN_ID);
        assertEq(input.destinationAddress, address(_home));
        vm.prank(address(_homeMessenger));
        _home.receiveTeleporterMessage(_REMOTE_BLOCKCHAIN_ID, address(_remote), input.message);
    }

    function _deliverToRemote() private {
        TeleporterMessageInput memory input = _homeMessenger.lastMessageInput();
        assertEq(input.destinationBlockchainID, _REMOTE_BLOCKCHAIN_ID);
        assertEq(input.destinationAddress, address(_remote));
        vm.prank(address(_remoteMessenger));
        _remote.receiveTeleporterMessage(_HOME_BLOCKCHAIN_ID, address(_home), input.message);
    }

    function _sendTokensInput(
        bytes32 destinationBlockchainID,
        address destinationTokenTransferrerAddress,
        address recipient
    ) private pure returns (SendTokensInput memory) {
        return SendTokensInput({
            destinationBlockchainID: destinationBlockchainID,
            destinationTokenTransferrerAddress: destinationTokenTransferrerAddress,
            recipient: recipient,
            primaryFeeTokenAddress: address(0),
            primaryFee: 0,
            secondaryFee: 0,
            requiredGasLimit: 250_000,
            multiHopFallback: address(0)
        });
    }
}
