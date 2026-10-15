// (c) 2026, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.
// SPDX-License-Identifier: LicenseRef-Ecosystem
pragma solidity ^0.8.30;

import {Test} from "@forge-std/Test.sol";
import {Initializable} from
    "@openzeppelin/contracts-upgradeable@5.1.0/proxy/utils/Initializable.sol";
import {
    IAdapter,
    TeleporterICMMessage,
    TeleporterMessageV2,
    TeleporterMessageInput,
    TeleporterMessageReceipt,
    TeleporterFeeInfo
} from "../ITeleporterMessengerV2.sol";
import {TeleporterMessengerV2} from "../TeleporterMessengerV2.sol";

/// @dev Adapter that accepts every message, so tests can reach the messenger's own checks.
contract AcceptAllAdapter is IAdapter {
    event MessageSent(TeleporterMessageV2 message);

    function sendMessage(
        TeleporterMessageV2 calldata message
    ) external {
        emit MessageSent(message);
    }

    function verifyMessage(
        TeleporterICMMessage calldata
    ) external pure returns (bool) {
        return true;
    }
}

/// @dev Receiver that rejects messages until it is opened, so a delivery can fail and be retried.
contract GatedReceiver {
    bool public open;
    bytes public lastMessage;

    function setOpen(
        bool open_
    ) external {
        open = open_;
    }

    function receiveTeleporterMessage(bytes32, address, bytes calldata message) external {
        require(open, "GatedReceiver: closed");
        lastMessage = message;
    }
}

contract TeleporterMessengerV2InitializeTest is Test {
    bytes32 private constant _BLOCKCHAIN_ID = hex"01";
    bytes32 private constant _OTHER_BLOCKCHAIN_ID = hex"02";
    address private constant _INITIALIZER = address(0x1234);
    address private constant _ATTACKER = address(0xBAD);

    AcceptAllAdapter private _adapter;
    TeleporterMessengerV2 private _teleporter;

    function setUp() public {
        _adapter = new AcceptAllAdapter();
        _teleporter = new TeleporterMessengerV2(address(_adapter), _INITIALIZER);
    }

    function testConstructorRevertsZeroInitializer() public {
        vm.expectRevert("TeleporterMessenger: zero initializer address");
        new TeleporterMessengerV2(address(_adapter), address(0));
    }

    function testConstructorRevertsZeroAdapter() public {
        vm.expectRevert("TeleporterMessenger: zero adapter address");
        new TeleporterMessengerV2(address(0), _INITIALIZER);
    }

    function testInitializeByInitializer() public {
        vm.prank(_INITIALIZER);
        _teleporter.initialize(_BLOCKCHAIN_ID);
        assertEq(_teleporter.blockchainID(), _BLOCKCHAIN_ID);
    }

    function testInitializeRevertsUnauthorized() public {
        vm.prank(_ATTACKER);
        vm.expectRevert("TeleporterMessenger: unauthorized initializer");
        _teleporter.initialize(_OTHER_BLOCKCHAIN_ID);
        assertEq(_teleporter.blockchainID(), bytes32(0));

        // A failed attempt must not consume the one-shot initializer.
        vm.prank(_INITIALIZER);
        _teleporter.initialize(_BLOCKCHAIN_ID);
        assertEq(_teleporter.blockchainID(), _BLOCKCHAIN_ID);
    }

    function testInitializeRevertsZeroBlockchainID() public {
        vm.prank(_INITIALIZER);
        vm.expectRevert("TeleporterMessenger: zero blockchain ID");
        _teleporter.initialize(bytes32(0));
    }

    function testInitializeRevertsWhenAlreadyInitialized() public {
        vm.startPrank(_INITIALIZER);
        _teleporter.initialize(_BLOCKCHAIN_ID);
        vm.expectRevert(Initializable.InvalidInitialization.selector);
        _teleporter.initialize(_OTHER_BLOCKCHAIN_ID);
        vm.stopPrank();
        assertEq(_teleporter.blockchainID(), _BLOCKCHAIN_ID);
    }

    /// @dev Before initialization the stored blockchain ID is zero. A message addressed to the zero
    /// blockchain ID must not be accepted as if it were addressed to this chain.
    function testReceiveRevertsWhenUninitialized() public {
        TeleporterICMMessage memory message = _buildMessage(bytes32(0));
        vm.expectRevert("TeleporterMessenger: zero blockchain ID");
        _teleporter.receiveCrossChainMessage(message, address(0));
    }

    /// @dev Sends must also wait for initialization, otherwise the message ID would be derived from
    /// the zero blockchain ID and could never be matched by later retries or receipts.
    function testSendRevertsWhenUninitialized() public {
        vm.expectRevert("TeleporterMessenger: zero blockchain ID");
        _teleporter.sendCrossChainMessage(_buildMessageInput());
    }

    function testSendSpecifiedReceiptsRevertsWhenUninitialized() public {
        vm.expectRevert("TeleporterMessenger: zero blockchain ID");
        _teleporter.sendSpecifiedReceipts(
            _OTHER_BLOCKCHAIN_ID,
            new bytes32[](0),
            TeleporterFeeInfo({feeTokenAddress: address(0), amount: 0}),
            new address[](0)
        );
    }

    function testGetNextMessageIDRevertsWhenUninitialized() public {
        vm.expectRevert("TeleporterMessenger: zero blockchain ID");
        _teleporter.getNextMessageID(_OTHER_BLOCKCHAIN_ID);
    }

    function testSendAfterInitialize() public {
        vm.prank(_INITIALIZER);
        _teleporter.initialize(_BLOCKCHAIN_ID);

        bytes32 expectedMessageID = _teleporter.getNextMessageID(_OTHER_BLOCKCHAIN_ID);
        bytes32 messageID = _teleporter.sendCrossChainMessage(_buildMessageInput());
        assertEq(messageID, expectedMessageID);
        assertEq(messageID, _teleporter.calculateMessageID(_BLOCKCHAIN_ID, _OTHER_BLOCKCHAIN_ID, 1));
    }

    function testReceiveRevertsWrongDestinationAfterInitialize() public {
        vm.prank(_INITIALIZER);
        _teleporter.initialize(_BLOCKCHAIN_ID);

        TeleporterICMMessage memory message = _buildMessage(_OTHER_BLOCKCHAIN_ID);
        vm.expectRevert("TeleporterMessenger: invalid destination chain ID");
        _teleporter.receiveCrossChainMessage(message, address(0));
    }

    function testConstructorSetsInitializer() public view {
        assertEq(_teleporter.initializerAddress(), _INITIALIZER);
        assertEq(_teleporter.blockchainID(), bytes32(0));
    }

    function _buildMessage(
        bytes32 destinationBlockchainID
    ) private view returns (TeleporterICMMessage memory) {
        TeleporterMessageV2 memory teleporterMessage = TeleporterMessageV2({
            messageNonce: 1,
            originSenderAddress: address(this),
            originTeleporterAddress: address(_teleporter),
            destinationBlockchainID: destinationBlockchainID,
            destinationAddress: address(this),
            requiredGasLimit: 100_000,
            allowedRelayerAddresses: new address[](0),
            receipts: new TeleporterMessageReceipt[](0),
            message: hex"deadbeef"
        });
        return TeleporterICMMessage({
            message: teleporterMessage,
            sourceNetworkID: 1,
            sourceBlockchainID: _OTHER_BLOCKCHAIN_ID,
            attestation: hex""
        });
    }

    function _buildMessageInput() private pure returns (TeleporterMessageInput memory) {
        return TeleporterMessageInput({
            destinationBlockchainID: _OTHER_BLOCKCHAIN_ID,
            destinationAddress: address(0xCAFE),
            feeInfo: TeleporterFeeInfo({feeTokenAddress: address(0), amount: 0}),
            requiredGasLimit: 100_000,
            allowedRelayerAddresses: new address[](0),
            message: hex"deadbeef"
        });
    }
}

contract TeleporterMessengerV2RetryTest is Test {
    bytes32 private constant _BLOCKCHAIN_ID = hex"01";
    bytes32 private constant _SOURCE_BLOCKCHAIN_ID = hex"02";

    TeleporterMessengerV2 private _teleporter;
    GatedReceiver private _receiver;

    function setUp() public {
        _teleporter = new TeleporterMessengerV2(address(new AcceptAllAdapter()), address(this));
        _teleporter.initialize(_BLOCKCHAIN_ID);
        _receiver = new GatedReceiver();
    }

    function testRetryFailedMessage() public {
        TeleporterMessageV2 memory message = _deliverFailingMessage();

        _receiver.setOpen(true);
        _teleporter.retryMessageExecution(_SOURCE_BLOCKCHAIN_ID, message);
        assertEq(_receiver.lastMessage(), hex"deadbeef");
        assertEq(_teleporter.receivedFailedMessageHashes(_messageID()), bytes32(0));

        vm.expectRevert("TeleporterMessenger: message not found");
        _teleporter.retryMessageExecution(_SOURCE_BLOCKCHAIN_ID, message);
    }

    function testRetryFailedMessageStillFailing() public {
        TeleporterMessageV2 memory message = _deliverFailingMessage();

        vm.expectRevert("TeleporterMessenger: retry execution failed");
        _teleporter.retryMessageExecution(_SOURCE_BLOCKCHAIN_ID, message);
        assertNotEq(_teleporter.receivedFailedMessageHashes(_messageID()), bytes32(0));
    }

    function testRetryRejectsAlteredMessage() public {
        TeleporterMessageV2 memory message = _deliverFailingMessage();
        _receiver.setOpen(true);

        message.message = hex"0bad";
        vm.expectRevert("TeleporterMessenger: invalid message hash");
        _teleporter.retryMessageExecution(_SOURCE_BLOCKCHAIN_ID, message);
    }

    function _deliverFailingMessage() private returns (TeleporterMessageV2 memory message) {
        message = TeleporterMessageV2({
            messageNonce: 1,
            originSenderAddress: address(this),
            originTeleporterAddress: address(_teleporter),
            destinationBlockchainID: _BLOCKCHAIN_ID,
            destinationAddress: address(_receiver),
            requiredGasLimit: 100_000,
            allowedRelayerAddresses: new address[](0),
            receipts: new TeleporterMessageReceipt[](0),
            message: hex"deadbeef"
        });
        _teleporter.receiveCrossChainMessage(
            TeleporterICMMessage({
                message: message,
                sourceNetworkID: 1,
                sourceBlockchainID: _SOURCE_BLOCKCHAIN_ID,
                attestation: hex""
            }),
            address(0)
        );
        assertNotEq(_teleporter.receivedFailedMessageHashes(_messageID()), bytes32(0));
    }

    function _messageID() private view returns (bytes32) {
        return _teleporter.calculateMessageID(_SOURCE_BLOCKCHAIN_ID, _BLOCKCHAIN_ID, 1);
    }
}
