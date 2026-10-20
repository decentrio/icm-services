// (c) 2026, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

// SPDX-License-Identifier: LicenseRef-Ecosystem

pragma solidity ^0.8.30;

import {Execution, Receipt} from "./StateManagerLibrary.sol";

/**
 * THIS IS AN EXAMPLE CONTRACT THAT USES UN-AUDITED CODE.
 * DO NOT USE THIS CODE IN PRODUCTION.
 */

/**
 * @notice Verifies Ethereum logs against trusted beacon block roots. Where the roots come from is
 * left to the contract that inherits it.
 */
abstract contract EthereumLogVerifier {
    bytes32 public constant UNDEFINED_ROOT = bytes32(0);

    uint256 public immutable sourceChainId;

    /**
     * @notice The beacon chain config used to verify execution layer data.
     */
    Execution.BeaconConfig internal _beaconConfig;

    /// @dev Named for the ZK sources it was first emitted for; the name is kept for the ABI.
    event ZKEventImported(
        uint256 indexed sourceChainId,
        uint256 indexed beaconSlot,
        address indexed emitter,
        bytes32 executionRoot,
        uint256 logIndex
    );

    constructor(
        uint256 sourceChainId_,
        Execution.BeaconConfig memory beaconConfig_
    ) {
        require(sourceChainId_ != 0, "Invalid chain ID");
        sourceChainId = sourceChainId_;
        _beaconConfig = beaconConfig_;
    }

    /**
     * @notice Verifies that a specific log was emitted on the beacon chain and returns the log.
     * @dev This is the main entry point for bridging events. It performs two key verifications:
     * 1. Execution Verification: Validates that the provided `targetReceiptsRoot` is part of the canonical beacon chain history using the `execProof`.
     * 2. Log Verification: Uses the validated `targetReceiptsRoot` to verify the inclusion  of a specific Receipt (and Log) via the `logProof`.
     * Once verified, it emits a `ZKEventImported` event.
     * @param execProof The execution proof linking a receipt root to a trusted beacon block root.
     * @param logProof The log proof establishing the receipt and log inclusion against a trusted receipts root.
     */
    function verifyLogAndExtract(
        Execution.Proof calldata execProof,
        Receipt.Proof calldata logProof
    ) external returns (bytes memory) {
        bytes32 anchorRoot = _beaconRoot(execProof.anchorSlot);
        require(anchorRoot != UNDEFINED_ROOT, "Anchor slot is undefined");
        bytes32 receiptsRoot = Execution.verify(anchorRoot, execProof, _beaconConfig);
        bytes memory logData = Receipt.verifyAndExtractLog(receiptsRoot, logProof);
        emit ZKEventImported(
            sourceChainId,
            execProof.targetSlot,
            logProof.expectedEmitter,
            execProof.targetExecutionHeaderRoot,
            logProof.logIndex
        );
        return logData;
    }

    /**
     * @notice The trusted beacon block root at `slot`, or `UNDEFINED_ROOT` if there is none.
     */
    function _beaconRoot(
        uint64 slot
    ) internal view virtual returns (bytes32);
}
