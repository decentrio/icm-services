// (c) 2026, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

// SPDX-License-Identifier: LicenseRef-Ecosystem

pragma solidity ^0.8.30;

import {EthereumLogVerifier} from "./EthereumLogVerifier.sol";
import {
    TeleporterICMMessage,
    TeleporterMessageV2,
    IAdapter,
    ITeleporterMessengerV2
} from "../../../common/ITeleporterMessengerV2.sol";
import {Execution, Receipt} from "./StateManagerLibrary.sol";

/**
 * THIS IS AN EXAMPLE CONTRACT THAT USES UN-AUDITED CODE.
 * DO NOT USE THIS CODE IN PRODUCTION.
 */
struct Attestation {
    Execution.Proof execProof;
    Receipt.Proof logProof;
}

/**
 * @notice Adapts Ethereum log verification for the Teleporter messaging protocol. Where the
 * trusted beacon roots come from is left to the contract that inherits it.
 */
abstract contract EthereumAdapter is EthereumLogVerifier, IAdapter {
    /// @notice topic0 of the TeleporterV2MessageSent event which is the only log topic verifyMessage accepts
    bytes32 private constant _MESSAGE_SENT_TOPIC = keccak256("TeleporterV2MessageSent(bytes)");

    /// @notice The network ID (mainnet/testnet/local) messages must originate from
    uint32 public immutable sourceNetworkId;

    event TeleporterV2MessageSent(bytes encodedMessage);

    error UnsupportedSourceChain(uint256 chainId);

    constructor() {
        sourceNetworkId = _networkIdForChain(sourceChainId);
    }

    /**
     * @notice Originates a message by emitting it as a log for the relayer to prove.
     * @dev Only the message's originTeleporterAddress, or the adapter that messenger sends through (e.g. an
     * Adapter wrapper dispatching by destination chain), may call this. verifyMessage on the destination
     * chain accepts any log emitted by this same address, so without this check anyone could originate a
     * message that is delivered as if the messenger had sent it. See {ITeleporterMessengerV2-messageSender}.
     */
    function sendMessage(
        TeleporterMessageV2 calldata message
    ) external {
        require(
            msg.sender == message.originTeleporterAddress
                || msg.sender
                    == address(
                        ITeleporterMessengerV2(message.originTeleporterAddress).messageSender()
                    ),
            "unauthorized sender"
        );
        emit TeleporterV2MessageSent(abi.encode(message));
    }

    /**
     * @notice Verifies a Teleporter message was emitted on the source chain by validating the attestation.
     * @dev Reverts unless the attestation proves that a `TeleporterV2MessageSent` log was emitted by the
     * trusted source-side adapter, whose address via Nick's method is `address(this)`, whose payload
     * equals the message being verified.
     */
    function verifyMessage(
        TeleporterICMMessage calldata message
    ) external returns (bool) {
        Attestation memory att = abi.decode(message.attestation, (Attestation));
        require(message.sourceBlockchainID == bytes32(sourceChainId), "bad source chain");
        require(message.sourceNetworkID == sourceNetworkId, "bad network");

        require(att.logProof.expectedEmitter == address(this), "bad emitter");
        require(att.logProof.expectedTopic0 == _MESSAGE_SENT_TOPIC, "bad topic");

        bytes memory logData = this.verifyLogAndExtract(att.execProof, att.logProof);
        bytes memory emitted = abi.decode(logData, (bytes));
        require(keccak256(emitted) == keccak256(abi.encode(message.message)), "payload mismatch");

        return true;
    }

    /// @notice Maps a source EVM chain ID to its ICM network ID.
    function _networkIdForChain(
        uint256 chainId_
    ) internal pure returns (uint32) {
        if (chainId_ == 1) return 1;
        if (chainId_ == 11155111) return 2;
        revert UnsupportedSourceChain(chainId_);
    }
}
