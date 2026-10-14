// (c) 2026, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

// SPDX-License-Identifier: LicenseRef-Ecosystem

pragma solidity 0.8.30;

import {TeleporterMessageInput} from "@teleporter/ITeleporterMessenger.sol";

/**
 * THIS IS AN EXAMPLE CONTRACT THAT USES UN-AUDITED CODE.
 * DO NOT USE THIS CODE IN PRODUCTION.
 */

/**
 * @dev The part of a Teleporter messenger that a routed app depends on. Both {TeleporterMessenger}
 * and {TeleporterMessengerV2} implement it, so an app can route some chains through Warp and
 * others through an external bridge.
 */
interface ITeleporterSender {
    function sendCrossChainMessage(
        TeleporterMessageInput calldata messageInput
    ) external returns (bytes32);

    function blockchainID() external view returns (bytes32);
}

/**
 * @dev Interface for a Teleporter app that binds each counterparty chain to the messengers allowed to
 * carry its messages, instead of trusting every messenger in a TeleporterRegistry for every chain.
 *
 * @custom:security-contact https://github.com/ava-labs/icm-contracts/blob/main/SECURITY.md
 */
interface ITeleporterRoutedApp {
    /**
     * @notice Emitted when a messenger is allowed to deliver messages from a remote chain.
     */
    event MessengerAdded(bytes32 indexed remoteBlockchainID, address indexed messenger);

    /**
     * @notice Emitted when a messenger is no longer allowed to deliver messages from a remote chain.
     */
    event MessengerRemoved(bytes32 indexed remoteBlockchainID, address indexed messenger);

    /**
     * @notice Emitted when the messenger used to send messages to a remote chain changes.
     */
    event OutboundMessengerSet(
        bytes32 indexed remoteBlockchainID,
        address indexed oldMessenger,
        address indexed newMessenger
    );

    /**
     * @notice Emitted when a messenger is paused for sending and receiving on every chain.
     */
    event MessengerPaused(address indexed messenger);

    /**
     * @notice Emitted when a paused messenger is unpaused.
     */
    event MessengerUnpaused(address indexed messenger);

    /**
     * @notice Allows `messenger` to deliver messages from `remoteBlockchainID`.
     * @dev The messenger must report this chain's blockchain ID, or zero if it has yet to learn it.
     */
    function addMessenger(bytes32 remoteBlockchainID, address messenger) external;

    /**
     * @notice Stops `messenger` from delivering messages from `remoteBlockchainID`.
     * @dev The outbound messenger of a chain cannot be removed; switch it first.
     */
    function removeMessenger(bytes32 remoteBlockchainID, address messenger) external;

    /**
     * @notice Sends future messages to `remoteBlockchainID` through `messenger`.
     * @dev `messenger` must already be allowed to deliver from that chain, since replies come back
     * through the same messenger.
     */
    function setOutboundMessenger(bytes32 remoteBlockchainID, address messenger) external;

    /**
     * @notice Pauses sending and receiving through `messenger` on every chain.
     */
    function pauseMessenger(
        address messenger
    ) external;

    /**
     * @notice Unpauses a paused messenger.
     */
    function unpauseMessenger(
        address messenger
    ) external;

    /**
     * @notice The blockchain ID of the chain this contract is deployed on.
     */
    function getBlockchainID() external view returns (bytes32);

    /**
     * @notice The messenger used to send messages to `remoteBlockchainID`, or zero if there is none.
     */
    function getOutboundMessenger(
        bytes32 remoteBlockchainID
    ) external view returns (address);

    /**
     * @notice Whether `messenger` may deliver messages from `remoteBlockchainID`.
     */
    function isInboundMessenger(
        bytes32 remoteBlockchainID,
        address messenger
    ) external view returns (bool);

    /**
     * @notice Whether `messenger` is paused.
     */
    function isMessengerPaused(
        address messenger
    ) external view returns (bool);
}
