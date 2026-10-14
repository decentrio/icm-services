// (c) 2026, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.
// SPDX-License-Identifier: LicenseRef-Ecosystem
pragma solidity 0.8.30;

import {Test} from "@forge-std/Test.sol";
import {TeleporterRoutedOwnableAppUpgradeable} from "../TeleporterRoutedOwnableAppUpgradeable.sol";
import {ITeleporterSender} from "../ITeleporterRoutedApp.sol";
import {
    IAdapter,
    TeleporterMessageInput,
    TeleporterFeeInfo,
    TeleporterMessageV2,
    TeleporterMessageReceipt,
    TeleporterICMMessage
} from "../ITeleporterMessengerV2.sol";
import {TeleporterMessengerV2} from "../TeleporterMessengerV2.sol";
import {ExampleERC20} from "@mocks/ExampleERC20.sol";
import {IERC20} from "@openzeppelin/contracts@5.1.0/token/ERC20/IERC20.sol";
import {OwnableUpgradeable} from
    "@openzeppelin/contracts-upgradeable@5.1.0/access/OwnableUpgradeable.sol";

/**
 * @dev Minimal routed app that exposes sending and records what it receives.
 */
contract ExampleRoutedApp is TeleporterRoutedOwnableAppUpgradeable {
    bytes32 public lastSourceBlockchainID;
    address public lastOriginSenderAddress;
    bytes public lastMessage;

    function initialize(bytes32 blockchainID, address owner) external initializer {
        __TeleporterRoutedOwnableApp_init(blockchainID, owner);
    }

    function sendTeleporterMessage(
        TeleporterMessageInput calldata messageInput
    ) external returns (bytes32) {
        return _sendTeleporterMessage(messageInput);
    }

    function _receiveTeleporterMessage(
        bytes32 sourceBlockchainID,
        address originSenderAddress,
        bytes memory message
    ) internal override {
        lastSourceBlockchainID = sourceBlockchainID;
        lastOriginSenderAddress = originSenderAddress;
        lastMessage = message;
    }
}

/**
 * @dev Messenger stand-in that reports a configurable blockchain ID and records sends.
 */
contract MockMessenger is ITeleporterSender {
    bytes32 public blockchainID;
    uint256 public sendCount;
    bytes32 public lastDestinationBlockchainID;

    constructor(
        bytes32 blockchainID_
    ) {
        blockchainID = blockchainID_;
    }

    function sendCrossChainMessage(
        TeleporterMessageInput calldata messageInput
    ) external returns (bytes32) {
        ++sendCount;
        lastDestinationBlockchainID = messageInput.destinationBlockchainID;
        return bytes32(sendCount);
    }
}

/**
 * @dev Adapter that emits nothing and accepts every message, so that a real
 * {TeleporterMessengerV2} can be driven directly in a single-chain test.
 */
contract AcceptAllAdapter is IAdapter {
    // solhint-disable no-empty-blocks
    function sendMessage(
        TeleporterMessageV2 calldata
    ) external {}
    // solhint-enable no-empty-blocks

    function verifyMessage(
        TeleporterICMMessage calldata
    ) external pure returns (bool) {
        return true;
    }
}

contract TeleporterRoutedAppTest is Test {
    bytes32 private constant _LOCAL_BLOCKCHAIN_ID = bytes32(uint256(43114));
    bytes32 private constant _ETHEREUM_BLOCKCHAIN_ID = bytes32(uint256(1));
    bytes32 private constant _L1_BLOCKCHAIN_ID = keccak256("L1");
    address private constant _OWNER = address(0x0123);
    address private constant _ORIGIN_SENDER = address(0x0456);

    ExampleRoutedApp private _app;
    MockMessenger private _ethMessenger;
    MockMessenger private _l1Messenger;

    event MessengerAdded(bytes32 indexed remoteBlockchainID, address indexed messenger);
    event MessengerRemoved(bytes32 indexed remoteBlockchainID, address indexed messenger);
    event OutboundMessengerSet(
        bytes32 indexed remoteBlockchainID,
        address indexed oldMessenger,
        address indexed newMessenger
    );
    event MessengerPaused(address indexed messenger);
    event MessengerUnpaused(address indexed messenger);

    function setUp() public {
        _app = new ExampleRoutedApp();
        _app.initialize(_LOCAL_BLOCKCHAIN_ID, _OWNER);
        _ethMessenger = new MockMessenger(_LOCAL_BLOCKCHAIN_ID);
        _l1Messenger = new MockMessenger(_LOCAL_BLOCKCHAIN_ID);
    }

    /**
     * Initialization
     */
    function testInitializeZeroBlockchainID() public {
        ExampleRoutedApp app = new ExampleRoutedApp();
        vm.expectRevert("TeleporterRoutedApp: zero blockchain ID");
        app.initialize(bytes32(0), _OWNER);
    }

    /**
     * addMessenger
     */
    function testAddMessenger() public {
        vm.expectEmit(true, true, true, true, address(_app));
        emit MessengerAdded(_ETHEREUM_BLOCKCHAIN_ID, address(_ethMessenger));
        _addMessenger(_ETHEREUM_BLOCKCHAIN_ID, address(_ethMessenger));

        assertTrue(_app.isInboundMessenger(_ETHEREUM_BLOCKCHAIN_ID, address(_ethMessenger)));
        assertFalse(_app.isInboundMessenger(_L1_BLOCKCHAIN_ID, address(_ethMessenger)));
        // Adding a messenger does not make it the outbound messenger.
        assertEq(_app.getOutboundMessenger(_ETHEREUM_BLOCKCHAIN_ID), address(0));
    }

    function testAddMessengerNotOwner() public {
        vm.expectRevert(
            abi.encodeWithSelector(
                OwnableUpgradeable.OwnableUnauthorizedAccount.selector, address(this)
            )
        );
        _app.addMessenger(_ETHEREUM_BLOCKCHAIN_ID, address(_ethMessenger));
    }

    function testAddMessengerZeroRemoteBlockchainID() public {
        vm.expectRevert("TeleporterRoutedApp: zero remote blockchain ID");
        _addMessenger(bytes32(0), address(_ethMessenger));
    }

    function testAddMessengerOwnBlockchain() public {
        vm.expectRevert("TeleporterRoutedApp: cannot route to own blockchain");
        _addMessenger(_LOCAL_BLOCKCHAIN_ID, address(_ethMessenger));
    }

    function testAddMessengerZeroAddress() public {
        vm.expectRevert("TeleporterRoutedApp: zero messenger address");
        _addMessenger(_ETHEREUM_BLOCKCHAIN_ID, address(0));
    }

    function testAddMessengerAlreadyAdded() public {
        _addMessenger(_ETHEREUM_BLOCKCHAIN_ID, address(_ethMessenger));
        vm.expectRevert("TeleporterRoutedApp: messenger already added");
        _addMessenger(_ETHEREUM_BLOCKCHAIN_ID, address(_ethMessenger));
    }

    function testAddMessengerBlockchainIDMismatch() public {
        MockMessenger foreignMessenger = new MockMessenger(_ETHEREUM_BLOCKCHAIN_ID);
        vm.expectRevert("TeleporterRoutedApp: messenger blockchain ID mismatch");
        _addMessenger(_ETHEREUM_BLOCKCHAIN_ID, address(foreignMessenger));
    }

    function testAddMessengerWithUninitializedBlockchainID() public {
        // A V1 TeleporterMessenger reports zero until it learns its blockchain ID.
        MockMessenger lazyMessenger = new MockMessenger(bytes32(0));
        _addMessenger(_L1_BLOCKCHAIN_ID, address(lazyMessenger));
        assertTrue(_app.isInboundMessenger(_L1_BLOCKCHAIN_ID, address(lazyMessenger)));
    }

    function testAddMessengerWithoutCode() public {
        vm.expectRevert();
        _addMessenger(_ETHEREUM_BLOCKCHAIN_ID, address(0xdead));
    }

    /**
     * setOutboundMessenger
     */
    function testSetOutboundMessenger() public {
        _addMessenger(_ETHEREUM_BLOCKCHAIN_ID, address(_ethMessenger));
        vm.expectEmit(true, true, true, true, address(_app));
        emit OutboundMessengerSet(_ETHEREUM_BLOCKCHAIN_ID, address(0), address(_ethMessenger));
        _setOutboundMessenger(_ETHEREUM_BLOCKCHAIN_ID, address(_ethMessenger));
        assertEq(_app.getOutboundMessenger(_ETHEREUM_BLOCKCHAIN_ID), address(_ethMessenger));
    }

    function testSetOutboundMessengerNotAdded() public {
        vm.expectRevert("TeleporterRoutedApp: messenger not added");
        _setOutboundMessenger(_ETHEREUM_BLOCKCHAIN_ID, address(_ethMessenger));
    }

    function testSetOutboundMessengerAddedForOtherChain() public {
        _addMessenger(_L1_BLOCKCHAIN_ID, address(_ethMessenger));
        vm.expectRevert("TeleporterRoutedApp: messenger not added");
        _setOutboundMessenger(_ETHEREUM_BLOCKCHAIN_ID, address(_ethMessenger));
    }

    function testSetOutboundMessengerAlreadyOutbound() public {
        _route(_ETHEREUM_BLOCKCHAIN_ID, address(_ethMessenger));
        vm.expectRevert("TeleporterRoutedApp: messenger already outbound");
        _setOutboundMessenger(_ETHEREUM_BLOCKCHAIN_ID, address(_ethMessenger));
    }

    function testSetOutboundMessengerNotOwner() public {
        _addMessenger(_ETHEREUM_BLOCKCHAIN_ID, address(_ethMessenger));
        vm.expectRevert(
            abi.encodeWithSelector(
                OwnableUpgradeable.OwnableUnauthorizedAccount.selector, address(this)
            )
        );
        _app.setOutboundMessenger(_ETHEREUM_BLOCKCHAIN_ID, address(_ethMessenger));
    }

    /**
     * removeMessenger
     */
    function testRemoveMessenger() public {
        _addMessenger(_ETHEREUM_BLOCKCHAIN_ID, address(_ethMessenger));
        vm.expectEmit(true, true, true, true, address(_app));
        emit MessengerRemoved(_ETHEREUM_BLOCKCHAIN_ID, address(_ethMessenger));
        vm.prank(_OWNER);
        _app.removeMessenger(_ETHEREUM_BLOCKCHAIN_ID, address(_ethMessenger));
        assertFalse(_app.isInboundMessenger(_ETHEREUM_BLOCKCHAIN_ID, address(_ethMessenger)));
    }

    function testRemoveMessengerNotAdded() public {
        vm.prank(_OWNER);
        vm.expectRevert("TeleporterRoutedApp: messenger not added");
        _app.removeMessenger(_ETHEREUM_BLOCKCHAIN_ID, address(_ethMessenger));
    }

    function testRemoveOutboundMessenger() public {
        _route(_ETHEREUM_BLOCKCHAIN_ID, address(_ethMessenger));
        vm.prank(_OWNER);
        vm.expectRevert("TeleporterRoutedApp: cannot remove outbound messenger");
        _app.removeMessenger(_ETHEREUM_BLOCKCHAIN_ID, address(_ethMessenger));
    }

    function testRotateMessenger() public {
        _route(_ETHEREUM_BLOCKCHAIN_ID, address(_ethMessenger));
        MockMessenger newMessenger = new MockMessenger(_LOCAL_BLOCKCHAIN_ID);

        // 1. Accept both messengers inbound.
        _addMessenger(_ETHEREUM_BLOCKCHAIN_ID, address(newMessenger));
        _receive(address(_ethMessenger), _ETHEREUM_BLOCKCHAIN_ID);
        _receive(address(newMessenger), _ETHEREUM_BLOCKCHAIN_ID);

        // 2. Send through the new messenger.
        vm.expectEmit(true, true, true, true, address(_app));
        emit OutboundMessengerSet(
            _ETHEREUM_BLOCKCHAIN_ID, address(_ethMessenger), address(newMessenger)
        );
        _setOutboundMessenger(_ETHEREUM_BLOCKCHAIN_ID, address(newMessenger));
        _app.sendTeleporterMessage(_messageInput(_ETHEREUM_BLOCKCHAIN_ID));
        assertEq(newMessenger.sendCount(), 1);
        assertEq(_ethMessenger.sendCount(), 0);

        // 3. Drop the old messenger once drained.
        vm.prank(_OWNER);
        _app.removeMessenger(_ETHEREUM_BLOCKCHAIN_ID, address(_ethMessenger));
        vm.prank(address(_ethMessenger));
        vm.expectRevert("TeleporterRoutedApp: unauthorized messenger");
        _app.receiveTeleporterMessage(_ETHEREUM_BLOCKCHAIN_ID, _ORIGIN_SENDER, hex"01");
    }

    /**
     * pauseMessenger / unpauseMessenger
     */
    function testPauseMessenger() public {
        _route(_ETHEREUM_BLOCKCHAIN_ID, address(_ethMessenger));
        vm.expectEmit(true, true, true, true, address(_app));
        emit MessengerPaused(address(_ethMessenger));
        vm.prank(_OWNER);
        _app.pauseMessenger(address(_ethMessenger));
        assertTrue(_app.isMessengerPaused(address(_ethMessenger)));

        vm.expectRevert("TeleporterRoutedApp: messenger paused");
        _app.sendTeleporterMessage(_messageInput(_ETHEREUM_BLOCKCHAIN_ID));

        vm.prank(address(_ethMessenger));
        vm.expectRevert("TeleporterRoutedApp: messenger paused");
        _app.receiveTeleporterMessage(_ETHEREUM_BLOCKCHAIN_ID, _ORIGIN_SENDER, hex"01");
    }

    function testUnpauseMessenger() public {
        _route(_ETHEREUM_BLOCKCHAIN_ID, address(_ethMessenger));
        vm.startPrank(_OWNER);
        _app.pauseMessenger(address(_ethMessenger));
        vm.expectEmit(true, true, true, true, address(_app));
        emit MessengerUnpaused(address(_ethMessenger));
        _app.unpauseMessenger(address(_ethMessenger));
        vm.stopPrank();

        assertFalse(_app.isMessengerPaused(address(_ethMessenger)));
        _app.sendTeleporterMessage(_messageInput(_ETHEREUM_BLOCKCHAIN_ID));
        _receive(address(_ethMessenger), _ETHEREUM_BLOCKCHAIN_ID);
    }

    function testPauseMessengerZeroAddress() public {
        vm.prank(_OWNER);
        vm.expectRevert("TeleporterRoutedApp: zero messenger address");
        _app.pauseMessenger(address(0));
    }

    function testPauseMessengerAlreadyPaused() public {
        vm.startPrank(_OWNER);
        _app.pauseMessenger(address(_ethMessenger));
        vm.expectRevert("TeleporterRoutedApp: messenger already paused");
        _app.pauseMessenger(address(_ethMessenger));
        vm.stopPrank();
    }

    function testUnpauseMessengerNotPaused() public {
        vm.prank(_OWNER);
        vm.expectRevert("TeleporterRoutedApp: messenger not paused");
        _app.unpauseMessenger(address(_ethMessenger));
    }

    function testPauseMessengerNotOwner() public {
        vm.expectRevert(
            abi.encodeWithSelector(
                OwnableUpgradeable.OwnableUnauthorizedAccount.selector, address(this)
            )
        );
        _app.pauseMessenger(address(_ethMessenger));
    }

    /**
     * Receiving
     */
    function testReceive() public {
        _route(_ETHEREUM_BLOCKCHAIN_ID, address(_ethMessenger));
        _receive(address(_ethMessenger), _ETHEREUM_BLOCKCHAIN_ID);
        assertEq(_app.lastSourceBlockchainID(), _ETHEREUM_BLOCKCHAIN_ID);
        assertEq(_app.lastOriginSenderAddress(), _ORIGIN_SENDER);
        assertEq(_app.lastMessage(), hex"01");
    }

    function testReceiveFromUnroutedChain() public {
        vm.prank(address(_ethMessenger));
        vm.expectRevert("TeleporterRoutedApp: unauthorized messenger");
        _app.receiveTeleporterMessage(_ETHEREUM_BLOCKCHAIN_ID, _ORIGIN_SENDER, hex"01");
    }

    function testReceiveFromMessengerOfOtherChain() public {
        // The L1 (Warp) messenger cannot deliver messages claiming to come from Ethereum,
        // and the Ethereum messenger cannot deliver messages claiming to come from the L1.
        _route(_ETHEREUM_BLOCKCHAIN_ID, address(_ethMessenger));
        _route(_L1_BLOCKCHAIN_ID, address(_l1Messenger));

        vm.prank(address(_l1Messenger));
        vm.expectRevert("TeleporterRoutedApp: unauthorized messenger");
        _app.receiveTeleporterMessage(_ETHEREUM_BLOCKCHAIN_ID, _ORIGIN_SENDER, hex"01");

        vm.prank(address(_ethMessenger));
        vm.expectRevert("TeleporterRoutedApp: unauthorized messenger");
        _app.receiveTeleporterMessage(_L1_BLOCKCHAIN_ID, _ORIGIN_SENDER, hex"01");
    }

    /**
     * Sending
     */
    function testSendUsesRouteOfDestination() public {
        _route(_ETHEREUM_BLOCKCHAIN_ID, address(_ethMessenger));
        _route(_L1_BLOCKCHAIN_ID, address(_l1Messenger));

        _app.sendTeleporterMessage(_messageInput(_ETHEREUM_BLOCKCHAIN_ID));
        _app.sendTeleporterMessage(_messageInput(_L1_BLOCKCHAIN_ID));
        _app.sendTeleporterMessage(_messageInput(_L1_BLOCKCHAIN_ID));

        assertEq(_ethMessenger.sendCount(), 1);
        assertEq(_ethMessenger.lastDestinationBlockchainID(), _ETHEREUM_BLOCKCHAIN_ID);
        assertEq(_l1Messenger.sendCount(), 2);
        assertEq(_l1Messenger.lastDestinationBlockchainID(), _L1_BLOCKCHAIN_ID);
    }

    function testSendNoOutboundMessenger() public {
        // Inbound alone is not enough to send.
        _addMessenger(_ETHEREUM_BLOCKCHAIN_ID, address(_ethMessenger));
        vm.expectRevert("TeleporterRoutedApp: no outbound messenger");
        _app.sendTeleporterMessage(_messageInput(_ETHEREUM_BLOCKCHAIN_ID));
    }

    function testSendApprovesFee() public {
        _route(_ETHEREUM_BLOCKCHAIN_ID, address(_ethMessenger));
        IERC20 feeToken = new ExampleERC20();
        TeleporterMessageInput memory input = _messageInput(_ETHEREUM_BLOCKCHAIN_ID);
        input.feeInfo = TeleporterFeeInfo({feeTokenAddress: address(feeToken), amount: 100});

        _app.sendTeleporterMessage(input);
        assertEq(feeToken.allowance(address(_app), address(_ethMessenger)), 100);
    }

    function testSendZeroFeeTokenAddress() public {
        _route(_ETHEREUM_BLOCKCHAIN_ID, address(_ethMessenger));
        TeleporterMessageInput memory input = _messageInput(_ETHEREUM_BLOCKCHAIN_ID);
        input.feeInfo = TeleporterFeeInfo({feeTokenAddress: address(0), amount: 100});

        vm.expectRevert("TeleporterRoutedApp: zero fee token address");
        _app.sendTeleporterMessage(input);
    }

    /**
     * Integration with TeleporterMessengerV2
     */
    function testWithTeleporterMessengerV2() public {
        TeleporterMessengerV2 messenger =
            new TeleporterMessengerV2(address(new AcceptAllAdapter()), address(this));
        messenger.initialize(_LOCAL_BLOCKCHAIN_ID);
        _route(_ETHEREUM_BLOCKCHAIN_ID, address(messenger));

        // Sending goes through the V2 messenger.
        bytes32 messageID = _app.sendTeleporterMessage(_messageInput(_ETHEREUM_BLOCKCHAIN_ID));
        assertEq(
            messageID,
            messenger.calculateMessageID(_LOCAL_BLOCKCHAIN_ID, _ETHEREUM_BLOCKCHAIN_ID, 1)
        );

        // Receiving through the V2 messenger reaches the app.
        TeleporterMessageV2 memory message = TeleporterMessageV2({
            messageNonce: 1,
            originSenderAddress: _ORIGIN_SENDER,
            originTeleporterAddress: address(messenger),
            destinationBlockchainID: _LOCAL_BLOCKCHAIN_ID,
            destinationAddress: address(_app),
            requiredGasLimit: 200_000,
            allowedRelayerAddresses: new address[](0),
            receipts: new TeleporterMessageReceipt[](0),
            message: hex"cafe"
        });
        messenger.receiveCrossChainMessage(
            TeleporterICMMessage({
                message: message,
                sourceNetworkID: 1,
                sourceBlockchainID: _ETHEREUM_BLOCKCHAIN_ID,
                attestation: new bytes(0)
            }),
            address(this)
        );
        assertEq(_app.lastSourceBlockchainID(), _ETHEREUM_BLOCKCHAIN_ID);
        assertEq(_app.lastOriginSenderAddress(), _ORIGIN_SENDER);
        assertEq(_app.lastMessage(), hex"cafe");
    }

    function testWithTeleporterMessengerV2OnOtherBlockchain() public {
        TeleporterMessengerV2 messenger =
            new TeleporterMessengerV2(address(new AcceptAllAdapter()), address(this));
        messenger.initialize(_ETHEREUM_BLOCKCHAIN_ID);
        vm.expectRevert("TeleporterRoutedApp: messenger blockchain ID mismatch");
        _addMessenger(_L1_BLOCKCHAIN_ID, address(messenger));
    }

    function testInitialize() public view {
        assertEq(_app.getBlockchainID(), _LOCAL_BLOCKCHAIN_ID);
        assertEq(_app.owner(), _OWNER);
        assertEq(_app.getOutboundMessenger(_ETHEREUM_BLOCKCHAIN_ID), address(0));
    }

    function _addMessenger(bytes32 remoteBlockchainID, address messenger) private {
        vm.prank(_OWNER);
        _app.addMessenger(remoteBlockchainID, messenger);
    }

    function _setOutboundMessenger(bytes32 remoteBlockchainID, address messenger) private {
        vm.prank(_OWNER);
        _app.setOutboundMessenger(remoteBlockchainID, messenger);
    }

    function _route(bytes32 remoteBlockchainID, address messenger) private {
        _addMessenger(remoteBlockchainID, messenger);
        _setOutboundMessenger(remoteBlockchainID, messenger);
    }

    function _receive(address messenger, bytes32 sourceBlockchainID) private {
        vm.prank(messenger);
        _app.receiveTeleporterMessage(sourceBlockchainID, _ORIGIN_SENDER, hex"01");
    }

    function _messageInput(
        bytes32 destinationBlockchainID
    ) private pure returns (TeleporterMessageInput memory) {
        return TeleporterMessageInput({
            destinationBlockchainID: destinationBlockchainID,
            destinationAddress: address(0xbeef),
            feeInfo: TeleporterFeeInfo({feeTokenAddress: address(0), amount: 0}),
            requiredGasLimit: 100_000,
            allowedRelayerAddresses: new address[](0),
            message: hex"01"
        });
    }
}
