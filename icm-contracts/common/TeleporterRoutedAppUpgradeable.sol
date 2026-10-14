// (c) 2026, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

// SPDX-License-Identifier: LicenseRef-Ecosystem

pragma solidity 0.8.30;

import {ITeleporterRoutedApp, ITeleporterSender} from "./ITeleporterRoutedApp.sol";
import {ITeleporterReceiver} from "@teleporter/ITeleporterReceiver.sol";
import {TeleporterMessageInput} from "@teleporter/ITeleporterMessenger.sol";
import {ContextUpgradeable} from
    "@openzeppelin/contracts-upgradeable@5.1.0/utils/ContextUpgradeable.sol";
import {ReentrancyGuardUpgradeable} from
    "@openzeppelin/contracts-upgradeable@5.1.0/utils/ReentrancyGuardUpgradeable.sol";
import {SafeERC20} from "@openzeppelin/contracts@5.1.0/token/ERC20/utils/SafeERC20.sol";
import {IERC20} from "@openzeppelin/contracts@5.1.0/token/ERC20/IERC20.sol";
import {Initializable} from
    "@openzeppelin/contracts-upgradeable@5.1.0/proxy/utils/Initializable.sol";

/**
 * THIS IS AN EXAMPLE CONTRACT THAT USES UN-AUDITED CODE.
 * DO NOT USE THIS CODE IN PRODUCTION.
 */

/**
 * @dev Base contract for a Teleporter app that routes each counterparty chain through its own
 * messenger(s), in place of {TeleporterRegistryAppUpgradeable}.
 *
 * A message claiming to come from a chain is accepted only from a messenger bound to that chain, so a
 * messenger whose security scheme only covers one connection (e.g. Avalanche <-> Ethereum) cannot be
 * used to impersonate another. The blockchain ID of this chain is supplied at initialization rather than
 * read from the Warp precompile, so the same contract runs on chains without Warp.
 *
 * @custom:security-contact https://github.com/ava-labs/icm-contracts/blob/main/SECURITY.md
 */
abstract contract TeleporterRoutedAppUpgradeable is
    Initializable,
    ContextUpgradeable,
    ITeleporterReceiver,
    ITeleporterRoutedApp,
    ReentrancyGuardUpgradeable
{
    using SafeERC20 for IERC20;

    // solhint-disable private-vars-leading-underscore
    /**
     * @dev Namespace storage slots following the ERC-7201 standard to prevent
     * storage collisions between upgradeable contracts.
     *
     * @custom:storage-location erc7201:teleporter.storage.TeleporterRoutedApp
     */
    struct TeleporterRoutedAppStorage {
        /// @notice The blockchain ID of the chain this contract is deployed on.
        bytes32 _blockchainID;
        /// @notice The messenger used to send messages to each remote chain.
        mapping(bytes32 remoteBlockchainID => address messenger) _outboundMessengers;
        /// @notice The messengers allowed to deliver messages from each remote chain.
        mapping(bytes32 remoteBlockchainID => mapping(address messenger => bool allowed))
            _inboundMessengers;
        /// @notice Messengers paused for sending and receiving on every chain.
        mapping(address messenger => bool paused) _pausedMessengers;
    }
    // solhint-enable private-vars-leading-underscore

    /**
     * @dev Storage slot computed based off ERC-7201 formula
     * keccak256(abi.encode(uint256(keccak256("teleporter.storage.TeleporterRoutedApp")) - 1)) & ~bytes32(uint256(0xff));
     */
    bytes32 public constant TELEPORTER_ROUTED_APP_STORAGE_LOCATION =
        0x63108c0ddc3d3ab0f279f27d4b9ca8220d5efbdbeeb537d4327d5cf9e982b100;

    // solhint-disable ordering
    function _getTeleporterRoutedAppStorage()
        private
        pure
        returns (TeleporterRoutedAppStorage storage $)
    {
        // solhint-disable-next-line no-inline-assembly
        assembly {
            $.slot := TELEPORTER_ROUTED_APP_STORAGE_LOCATION
        }
    }

    // solhint-disable-next-line func-name-mixedcase
    function __TeleporterRoutedApp_init(
        bytes32 blockchainID
    ) internal onlyInitializing {
        __ReentrancyGuard_init();
        __Context_init();
        __TeleporterRoutedApp_init_unchained(blockchainID);
    }

    // solhint-disable-next-line func-name-mixedcase
    function __TeleporterRoutedApp_init_unchained(
        bytes32 blockchainID
    ) internal onlyInitializing {
        require(blockchainID != bytes32(0), "TeleporterRoutedApp: zero blockchain ID");
        _getTeleporterRoutedAppStorage()._blockchainID = blockchainID;
    }
    // solhint-enable ordering

    /**
     * @dev See {ITeleporterReceiver-receiveTeleporterMessage}
     *
     * Requirements:
     *
     * - The caller must be a messenger bound to `sourceBlockchainID`.
     * - The caller must not be paused.
     */
    function receiveTeleporterMessage(
        bytes32 sourceBlockchainID,
        address originSenderAddress,
        bytes calldata message
    ) external nonReentrant {
        TeleporterRoutedAppStorage storage $ = _getTeleporterRoutedAppStorage();
        require(
            $._inboundMessengers[sourceBlockchainID][_msgSender()],
            "TeleporterRoutedApp: unauthorized messenger"
        );
        require(!$._pausedMessengers[_msgSender()], "TeleporterRoutedApp: messenger paused");

        _receiveTeleporterMessage(sourceBlockchainID, originSenderAddress, message);
    }

    /**
     * @dev See {ITeleporterRoutedApp-addMessenger}
     */
    function addMessenger(bytes32 remoteBlockchainID, address messenger) public virtual {
        _checkTeleporterRoutedAppAccess();
        _addMessenger(remoteBlockchainID, messenger);
    }

    /**
     * @dev See {ITeleporterRoutedApp-removeMessenger}
     */
    function removeMessenger(bytes32 remoteBlockchainID, address messenger) public virtual {
        _checkTeleporterRoutedAppAccess();
        TeleporterRoutedAppStorage storage $ = _getTeleporterRoutedAppStorage();
        require(
            $._inboundMessengers[remoteBlockchainID][messenger],
            "TeleporterRoutedApp: messenger not added"
        );
        require(
            $._outboundMessengers[remoteBlockchainID] != messenger,
            "TeleporterRoutedApp: cannot remove outbound messenger"
        );
        $._inboundMessengers[remoteBlockchainID][messenger] = false;
        emit MessengerRemoved(remoteBlockchainID, messenger);
    }

    /**
     * @dev See {ITeleporterRoutedApp-setOutboundMessenger}
     */
    function setOutboundMessenger(bytes32 remoteBlockchainID, address messenger) public virtual {
        _checkTeleporterRoutedAppAccess();
        _setOutboundMessenger(remoteBlockchainID, messenger);
    }

    /**
     * @dev See {ITeleporterRoutedApp-pauseMessenger}
     */
    function pauseMessenger(
        address messenger
    ) public virtual {
        _checkTeleporterRoutedAppAccess();
        TeleporterRoutedAppStorage storage $ = _getTeleporterRoutedAppStorage();
        require(messenger != address(0), "TeleporterRoutedApp: zero messenger address");
        require(!$._pausedMessengers[messenger], "TeleporterRoutedApp: messenger already paused");
        $._pausedMessengers[messenger] = true;
        emit MessengerPaused(messenger);
    }

    /**
     * @dev See {ITeleporterRoutedApp-unpauseMessenger}
     */
    function unpauseMessenger(
        address messenger
    ) public virtual {
        _checkTeleporterRoutedAppAccess();
        TeleporterRoutedAppStorage storage $ = _getTeleporterRoutedAppStorage();
        require(messenger != address(0), "TeleporterRoutedApp: zero messenger address");
        require($._pausedMessengers[messenger], "TeleporterRoutedApp: messenger not paused");
        $._pausedMessengers[messenger] = false;
        emit MessengerUnpaused(messenger);
    }

    /**
     * @dev See {ITeleporterRoutedApp-getBlockchainID}
     */
    function getBlockchainID() public view returns (bytes32) {
        return _getTeleporterRoutedAppStorage()._blockchainID;
    }

    /**
     * @dev See {ITeleporterRoutedApp-getOutboundMessenger}
     */
    function getOutboundMessenger(
        bytes32 remoteBlockchainID
    ) public view returns (address) {
        return _getTeleporterRoutedAppStorage()._outboundMessengers[remoteBlockchainID];
    }

    /**
     * @dev See {ITeleporterRoutedApp-isInboundMessenger}
     */
    function isInboundMessenger(
        bytes32 remoteBlockchainID,
        address messenger
    ) public view returns (bool) {
        return _getTeleporterRoutedAppStorage()._inboundMessengers[remoteBlockchainID][messenger];
    }

    /**
     * @dev See {ITeleporterRoutedApp-isMessengerPaused}
     */
    function isMessengerPaused(
        address messenger
    ) public view returns (bool) {
        return _getTeleporterRoutedAppStorage()._pausedMessengers[messenger];
    }

    /**
     * @dev Allows `messenger` to deliver messages from `remoteBlockchainID`. Not access controlled, so
     * that inheriting contracts can bind messengers during initialization.
     *
     * Emits a {MessengerAdded} event.
     */
    function _addMessenger(bytes32 remoteBlockchainID, address messenger) internal {
        TeleporterRoutedAppStorage storage $ = _getTeleporterRoutedAppStorage();
        require(remoteBlockchainID != bytes32(0), "TeleporterRoutedApp: zero remote blockchain ID");
        require(
            remoteBlockchainID != $._blockchainID,
            "TeleporterRoutedApp: cannot route to own blockchain"
        );
        require(messenger != address(0), "TeleporterRoutedApp: zero messenger address");
        require(
            !$._inboundMessengers[remoteBlockchainID][messenger],
            "TeleporterRoutedApp: messenger already added"
        );

        // A V1 TeleporterMessenger learns its blockchain ID lazily and reports zero until then.
        // Any other value must match, otherwise this app and the messenger disagree on which chain
        // they are on and messages would be addressed to the wrong chain.
        bytes32 messengerBlockchainID = ITeleporterSender(messenger).blockchainID();
        require(
            messengerBlockchainID == bytes32(0) || messengerBlockchainID == $._blockchainID,
            "TeleporterRoutedApp: messenger blockchain ID mismatch"
        );

        $._inboundMessengers[remoteBlockchainID][messenger] = true;
        emit MessengerAdded(remoteBlockchainID, messenger);
    }

    /**
     * @dev Sends future messages to `remoteBlockchainID` through `messenger`. Not access controlled, so
     * that inheriting contracts can bind messengers during initialization.
     *
     * Emits an {OutboundMessengerSet} event.
     */
    function _setOutboundMessenger(bytes32 remoteBlockchainID, address messenger) internal {
        TeleporterRoutedAppStorage storage $ = _getTeleporterRoutedAppStorage();
        require(
            $._inboundMessengers[remoteBlockchainID][messenger],
            "TeleporterRoutedApp: messenger not added"
        );
        address oldMessenger = $._outboundMessengers[remoteBlockchainID];
        require(oldMessenger != messenger, "TeleporterRoutedApp: messenger already outbound");
        $._outboundMessengers[remoteBlockchainID] = messenger;
        emit OutboundMessengerSet(remoteBlockchainID, oldMessenger, messenger);
    }

    /**
     * @dev Sends a message through the outbound messenger of its destination chain, approving the
     * messenger to pull the fee if there is one.
     * @return The message ID assigned by the messenger.
     */
    function _sendTeleporterMessage(
        TeleporterMessageInput memory messageInput
    ) internal virtual returns (bytes32) {
        ITeleporterSender messenger = _getOutboundMessenger(messageInput.destinationBlockchainID);
        if (messageInput.feeInfo.amount > 0) {
            require(
                messageInput.feeInfo.feeTokenAddress != address(0),
                "TeleporterRoutedApp: zero fee token address"
            );
            IERC20(messageInput.feeInfo.feeTokenAddress).safeIncreaseAllowance(
                address(messenger), messageInput.feeInfo.amount
            );
        }
        return messenger.sendCrossChainMessage(messageInput);
    }

    /**
     * @dev Returns the outbound messenger of `remoteBlockchainID`, reverting if there is none or it
     * is paused.
     */
    function _getOutboundMessenger(
        bytes32 remoteBlockchainID
    ) internal view returns (ITeleporterSender) {
        TeleporterRoutedAppStorage storage $ = _getTeleporterRoutedAppStorage();
        address messenger = $._outboundMessengers[remoteBlockchainID];
        require(messenger != address(0), "TeleporterRoutedApp: no outbound messenger");
        require(!$._pausedMessengers[messenger], "TeleporterRoutedApp: messenger paused");
        return ITeleporterSender(messenger);
    }

    /**
     * @dev Receives a message whose messenger has already been checked against its source chain.
     */
    function _receiveTeleporterMessage(
        bytes32 sourceBlockchainID,
        address originSenderAddress,
        bytes memory message
    ) internal virtual;

    /**
     * @dev Checks that the caller may change routes and pause messengers. Must revert otherwise.
     */
    function _checkTeleporterRoutedAppAccess() internal virtual;
}
