// (c) 2026, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

// SPDX-License-Identifier: LicenseRef-Ecosystem

pragma solidity ^0.8.30;

import {Test, Vm} from "forge-std/Test.sol";
import {IRiscZeroVerifier, Receipt as RiscZeroReceipt} from "@risc0/IRiscZeroVerifier.sol";
import {ZKAdapter} from "../ZKAdapter.sol";
import {EthereumAdapter, Attestation} from "../EthereumAdapter.sol";
import {EthereumLogVerifier} from "../EthereumLogVerifier.sol";
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
 * @dev Behaviour every Ethereum adapter shares, whatever its source of beacon roots.
 *
 * sendMessage emits a log that verifyMessage on the destination chain accepts as authentic, so
 * only the messenger named in the message, or the adapter that messenger sends through, may call it.
 * Accepting the adapter is what lets the adapter sit behind an {Adapter} wrapper that dispatches by
 * destination chain, which a bidirectional Avalanche <-> Ethereum messenger needs.
 *
 * verifyMessage accepts only a TeleporterV2MessageSent log emitted by the adapter itself, from the
 * source chain and network it serves, whose payload is the message being delivered.
 */
abstract contract EthereumAdapterTest is Test {
    bytes32 private constant _AVALANCHE_BLOCKCHAIN_ID = keccak256("avalanche");
    bytes32 private constant _ETHEREUM_BLOCKCHAIN_ID = bytes32(uint256(1));
    address private constant _WRAPPER_ADAPTER = address(0xA11CE);
    address private constant _ATTACKER = address(0xBAD);

    EthereumAdapter private _adapter;
    ZKAdapterTestTeleporter private _teleporter;

    event TeleporterV2MessageSent(bytes encodedMessage);

    function setUp() public {
        _adapter = _newAdapter();
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

    function testVerifyMessageRejectsAnotherSourceChain() public {
        TeleporterICMMessage memory icm = _icmMessage();
        icm.sourceBlockchainID = bytes32(uint256(11_155_111));
        vm.expectRevert("bad source chain");
        _adapter.verifyMessage(icm);
    }

    function testVerifyMessageRejectsAnotherNetwork() public {
        TeleporterICMMessage memory icm = _icmMessage();
        icm.sourceNetworkID = 2;
        vm.expectRevert("bad network");
        _adapter.verifyMessage(icm);
    }

    function testVerifyMessageRejectsALogFromAnotherEmitter() public {
        Attestation memory att = _attestation();
        att.logProof.expectedEmitter = _ATTACKER;
        TeleporterICMMessage memory icm = _icmMessage();
        icm.attestation = abi.encode(att);
        vm.expectRevert("bad emitter");
        _adapter.verifyMessage(icm);
    }

    function testVerifyMessageRejectsAnotherEvent() public {
        Attestation memory att = _attestation();
        att.logProof.expectedTopic0 = keccak256("Transfer(address,address,uint256)");
        TeleporterICMMessage memory icm = _icmMessage();
        icm.attestation = abi.encode(att);
        vm.expectRevert("bad topic");
        _adapter.verifyMessage(icm);
    }

    function testVerifyMessageAcceptsTheLogOfThisMessage() public {
        TeleporterICMMessage memory icm = _icmMessage();
        _mockProvenLog(abi.encode(icm.message));
        assertTrue(_adapter.verifyMessage(icm));
    }

    function testVerifyMessageRejectsTheLogOfAnotherMessage() public {
        TeleporterICMMessage memory icm = _icmMessage();
        TeleporterMessageV2 memory other = _message();
        other.message = hex"0bad";
        _mockProvenLog(abi.encode(other));
        vm.expectRevert("payload mismatch");
        _adapter.verifyMessage(icm);
    }

    /// @dev The adapter under test, serving Ethereum mainnet (chain ID 1).
    function _newAdapter() internal virtual returns (EthereumAdapter);

    /**
     * @dev Stands in for the proofs: the adapter's verifyLogAndExtract returns a
     * TeleporterV2MessageSent log carrying `encodedMessage`.
     */
    function _mockProvenLog(
        bytes memory encodedMessage
    ) internal {
        vm.mockCall(
            address(_adapter),
            abi.encodeWithSelector(EthereumLogVerifier.verifyLogAndExtract.selector),
            abi.encode(abi.encode(encodedMessage))
        );
    }

    function _attestation() internal view returns (Attestation memory att) {
        att.logProof.expectedEmitter = address(_adapter);
        att.logProof.expectedTopic0 = keccak256("TeleporterV2MessageSent(bytes)");
    }

    function _icmMessage() internal view returns (TeleporterICMMessage memory icm) {
        icm.message = _message();
        icm.sourceNetworkID = 1;
        icm.sourceBlockchainID = bytes32(uint256(1));
        icm.attestation = abi.encode(_attestation());
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

contract ZKAdapterTest is EthereumAdapterTest {
    uint64 private constant _START_EPOCH = 1_000;
    uint256 private constant _GENESIS_TIME = 1_606_824_023;

    function _newAdapter() internal override returns (EthereumAdapter) {
        // The starting state must not be in the future.
        vm.warp(_GENESIS_TIME + uint256(_START_EPOCH + 2) * 32 * 12);
        Execution.BeaconConfig memory config;
        address admin = makeAddr("admin");
        return new ZKAdapter({
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
    }
}
