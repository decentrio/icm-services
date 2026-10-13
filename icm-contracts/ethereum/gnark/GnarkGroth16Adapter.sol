// SPDX-License-Identifier: MIT
pragma solidity ^0.8.28;

import {ISP1Verifier} from "@sp1-contracts/ISP1Verifier.sol";

/// @notice The verifier gnark exports for the Avalanche quorum circuit.
interface IAvaxQuorumGroth16Verifier {
    function verifyProof(
        bytes calldata proof,
        uint256[2] calldata input
    ) external view;
}

/**
 * @title GnarkGroth16Adapter
 * @notice Lets an unmodified ZKValidatorSetRegistry verify gnark Groth16 proofs in place of
 * SP1 proofs. The registry is deployed with this adapter as its `sp1Verifier_` and with
 * VKEY as its `attestationProgramVKey_`.
 * @dev The gnark circuit proves the same statement as the SP1 guest program and exposes
 * sha256(publicValues) as two public inputs: the high and low 128 bits. The registry checks
 * the decoded public values against its own state; this adapter binds those exact bytes to
 * the proof.
 */
contract GnarkGroth16Adapter is ISP1Verifier {
    /// @notice The gnark-exported Groth16 verifier for the circuit.
    IAvaxQuorumGroth16Verifier public immutable VERIFIER;
    /// @notice Identifier of the circuit's verifying key (sha256 of the serialized key).
    bytes32 public immutable VKEY;

    error InvalidConfig();
    error WrongVKey(bytes32 programVKey);

    constructor(
        address verifier,
        bytes32 vkey
    ) {
        if (verifier == address(0) || vkey == bytes32(0)) {
            revert InvalidConfig();
        }
        VERIFIER = IAvaxQuorumGroth16Verifier(verifier);
        VKEY = vkey;
    }

    /// @inheritdoc ISP1Verifier
    function verifyProof(
        bytes32 programVKey,
        bytes calldata publicValues,
        bytes calldata proofBytes
    ) external view {
        if (programVKey != VKEY) {
            revert WrongVKey(programVKey);
        }
        uint256 digest = uint256(sha256(publicValues));
        uint256[2] memory input = [digest >> 128, digest & type(uint128).max];
        VERIFIER.verifyProof(proofBytes, input);
    }
}
