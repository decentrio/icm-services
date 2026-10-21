// (c) 2026, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

// SPDX-License-Identifier: LicenseRef-Ecosystem

pragma solidity ^0.8.30;

import {ITeleporterReceiver} from "@teleporter/ITeleporterReceiver.sol";
import {
    TransferrerMessage,
    TransferrerMessageType,
    RegisterRemoteMessage
} from "../../avalanche/ictt/interfaces/ITokenTransferrer.sol";

/**
 * THIS IS AN EXAMPLE CONTRACT THAT USES UN-AUDITED CODE.
 * DO NOT USE THIS CODE IN PRODUCTION.
 */

/**
 * @notice Registers a remote with a token home on its owner's word, for a remote chain whose
 * messages cannot reach the home yet. The home's owner adds the registrar as an inbound messenger
 * for the remote's chain, calls register, and removes it again. The home trusts it exactly as it
 * trusts any messenger its owner adds, and applies its usual registration rules.
 */
contract RemoteRegistrar {
    /// @notice Who may register through this registrar: whoever deployed it.
    address public immutable owner;
    /// @notice The token home this registrar registers remotes with.
    ITeleporterReceiver public immutable home;
    /// @notice The home's blockchain ID, which the home checks when it adds a messenger.
    bytes32 public immutable blockchainID;

    constructor(
        ITeleporterReceiver home_,
        bytes32 homeBlockchainID
    ) {
        owner = msg.sender;
        home = home_;
        blockchainID = homeBlockchainID;
    }

    /**
     * @notice Delivers to the home the message a remote's registerWithHome would send, with no
     * initial reserve imbalance.
     */
    function register(
        bytes32 remoteBlockchainID,
        address remote,
        uint8 homeTokenDecimals,
        uint8 remoteTokenDecimals
    ) external {
        require(msg.sender == owner, "RemoteRegistrar: not owner");
        bytes memory payload = abi.encode(
            RegisterRemoteMessage({
                initialReserveImbalance: 0,
                homeTokenDecimals: homeTokenDecimals,
                remoteTokenDecimals: remoteTokenDecimals
            })
        );
        home.receiveTeleporterMessage(
            remoteBlockchainID,
            remote,
            abi.encode(
                TransferrerMessage({
                    messageType: TransferrerMessageType.REGISTER_REMOTE, payload: payload
                })
            )
        );
    }
}
