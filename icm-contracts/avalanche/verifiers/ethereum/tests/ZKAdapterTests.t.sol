// (c) 2026, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

// SPDX-License-Identifier: LicenseRef-Ecosystem

pragma solidity ^0.8.30;

import {Test, Vm} from "forge-std/Test.sol";
import {IRiscZeroVerifier, Receipt as RiscZeroReceipt} from "@risc0/IRiscZeroVerifier.sol";
import {ZKAdapter} from "../ZKAdapter.sol";
import {Consensus, Execution} from "../StateManagerLibrary.sol";
import {
    IAdapter,
    TeleporterMessageV2,
    TeleporterICMMessage,
    TeleporterMessageInput,
    TeleporterFeeInfo
} from "@common/ITeleporterMessengerV2.sol";
import {Adapter} from "@common/Adapter.sol";
import {TeleporterMessengerV2} from "@common/TeleporterMessengerV2.sol";

// solhint-disable no-empty-blocks
contract ZKAdapterTestVerifier is IRiscZeroVerifier {
    function verify(bytes calldata, bytes32, bytes32) external pure {}

    function verifyIntegrity(
        RiscZeroReceipt calldata
    ) external pure {}
}

/// @dev Stands in for the adapter of the other direction behind an {Adapter} wrapper.
contract ZKAdapterTestOtherAdapter is IAdapter {
    function sendMessage(
        TeleporterMessageV2 calldata
    ) external {}

    function verifyMessage(
        TeleporterICMMessage calldata
    ) external pure returns (bool) {
        return false;
    }
}
// solhint-enable no-empty-blocks

/// @dev Minimal stand-in for a TeleporterMessengerV2: exposes the adapter it sends through.
contract ZKAdapterTestTeleporter {
    address public immutable messageSender;

    constructor(
        address messageSender_
    ) {
        messageSender = messageSender_;
    }
}

/**
 * @dev sendMessage emits a log that verifyMessage on the destination chain accepts as authentic, so
 * only the messenger named in the message, or the adapter that messenger sends through, may call it.
 * Accepting the adapter is what lets ZKAdapter sit behind an {Adapter} wrapper that dispatches by
 * destination chain, which a bidirectional Avalanche <-> Ethereum messenger needs.
 */
contract ZKAdapterSendMessageTest is Test {
    uint64 private constant _START_EPOCH = 1_000;
    uint256 private constant _GENESIS_TIME = 1_606_824_023;
    bytes32 private constant _AVALANCHE_BLOCKCHAIN_ID = keccak256("avalanche");
    bytes32 private constant _ETHEREUM_BLOCKCHAIN_ID = bytes32(uint256(1));
    address private constant _WRAPPER_ADAPTER = address(0xA11CE);
    address private constant _ATTACKER = address(0xBAD);

    ZKAdapter private _adapter;
    ZKAdapterTestTeleporter private _teleporter;

    event TeleporterV2MessageSent(bytes encodedMessage);

    function setUp() public {
        vm.warp(_GENESIS_TIME + uint256(_START_EPOCH + 2) * 32 * 12);
        Execution.BeaconConfig memory config;
        address admin = makeAddr("admin");
        _adapter = new ZKAdapter({
            sourceChainId_: 1,
            startingState: Consensus.State({
                currentJustifiedCheckpoint: Consensus.Checkpoint({
                    epoch: _START_EPOCH + 1,
                    root: keccak256("justified")
                }),
                finalizedCheckpoint: Consensus.Checkpoint({
                    epoch: _START_EPOCH,
                    root: keccak256("finalized")
                })
            }),
            beaconConfig_: config,
            genesisTime_: _GENESIS_TIME,
            permissibleTimespan_: 86_400,
            verifier: address(new ZKAdapterTestVerifier()),
            imageID_: bytes32(uint256(1)),
            admin: admin,
            superAdmin: admin
        });
        _teleporter = new ZKAdapterTestTeleporter(_WRAPPER_ADAPTER);
    }

    function testSendMessageFromOriginTeleporter() public {
        TeleporterMessageV2 memory message = _message();
        vm.expectEmit(true, true, true, true, address(_adapter));
        emit TeleporterV2MessageSent(abi.encode(message));
        vm.prank(address(_teleporter));
        _adapter.sendMessage(message);
    }

    function testSendMessageFromOriginTeleporterAdapter() public {
        TeleporterMessageV2 memory message = _message();
        vm.expectEmit(true, true, true, true, address(_adapter));
        emit TeleporterV2MessageSent(abi.encode(message));
        vm.prank(_WRAPPER_ADAPTER);
        _adapter.sendMessage(message);
    }

    function testSendMessageRevertsForUnauthorizedSender() public {
        TeleporterMessageV2 memory message = _message();
        vm.prank(_ATTACKER);
        vm.expectRevert("unauthorized sender");
        _adapter.sendMessage(message);
    }

    function testSendMessageRevertsWhenOriginTeleporterIsNotAContract() public {
        // Looking up the messenger's adapter fails on an address without code, which rejects the send.
        TeleporterMessageV2 memory message = _message();
        message.originTeleporterAddress = address(0xBEEF);
        vm.prank(_ATTACKER);
        vm.expectRevert();
        _adapter.sendMessage(message);
    }

    function testSendMessageRevertsForAdapterOfAnotherTeleporter() public {
        // A message naming the real messenger must not be accepted from the adapter of another one.
        ZKAdapterTestTeleporter attackerTeleporter = new ZKAdapterTestTeleporter(_ATTACKER);
        TeleporterMessageV2 memory message = _message();
        message.originTeleporterAddress = address(attackerTeleporter);
        vm.prank(_WRAPPER_ADAPTER);
        vm.expectRevert("unauthorized sender");
        _adapter.sendMessage(message);
    }

    /**
     * @dev On Ethereum, a bidirectional messenger sends to Avalanche through an {Adapter} that routes
     * messages destined for Avalanche to ZKAdapter, so ZKAdapter is called by the wrapper rather than
     * by the messenger.
     */
    function testSendThroughAdapterWrapper() public {
        Adapter wrapper = new Adapter(
            _AVALANCHE_BLOCKCHAIN_ID,
            _ETHEREUM_BLOCKCHAIN_ID,
            address(_adapter),
            address(new ZKAdapterTestOtherAdapter())
        );
        TeleporterMessengerV2 messenger = new TeleporterMessengerV2(address(wrapper), address(this));
        messenger.initialize(_ETHEREUM_BLOCKCHAIN_ID);

        vm.recordLogs();
        messenger.sendCrossChainMessage(
            TeleporterMessageInput({
                destinationBlockchainID: _AVALANCHE_BLOCKCHAIN_ID,
                destinationAddress: address(0xD5),
                feeInfo: TeleporterFeeInfo({feeTokenAddress: address(0), amount: 0}),
                requiredGasLimit: 100_000,
                allowedRelayerAddresses: new address[](0),
                message: hex"cafe"
            })
        );

        bool emitted;
        bytes32 topic = keccak256("TeleporterV2MessageSent(bytes)");
        Vm.Log[] memory logs = vm.getRecordedLogs();
        for (uint256 i; i < logs.length; ++i) {
            if (logs[i].emitter == address(_adapter) && logs[i].topics[0] == topic) {
                TeleporterMessageV2 memory sent =
                    abi.decode(abi.decode(logs[i].data, (bytes)), (TeleporterMessageV2));
                assertEq(sent.originTeleporterAddress, address(messenger));
                assertEq(sent.destinationBlockchainID, _AVALANCHE_BLOCKCHAIN_ID);
                assertEq(sent.message, hex"cafe");
                emitted = true;
            }
        }
        assertTrue(emitted);
    }

    function _message() internal view returns (TeleporterMessageV2 memory message) {
        message.messageNonce = 1;
        message.originSenderAddress = address(0x5E11DE5);
        message.originTeleporterAddress = address(_teleporter);
        message.destinationBlockchainID = _AVALANCHE_BLOCKCHAIN_ID;
        message.destinationAddress = address(0xD5);
        message.requiredGasLimit = 100_000;
        message.message = hex"deadbeef";
    }
}
