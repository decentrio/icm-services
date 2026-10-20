// (c) 2026, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.
// SPDX-License-Identifier: LicenseRef-Ecosystem
pragma solidity ^0.8.30;

import {Test} from "forge-std/Test.sol";
import {Execution} from "../StateManagerLibrary.sol";
import {SSZ} from "../../../utilities/SimpleSerialize.sol";
import {Math} from "@openzeppelin/contracts@5.1.0/utils/math/Math.sol";

contract ExecutionHarness {
    Execution.BeaconConfig internal _config;

    function setConfig(
        Execution.BeaconConfig memory newConfig
    ) external {
        _config = newConfig;
    }

    function getConfig() external view returns (Execution.BeaconConfig memory) {
        return _config;
    }

    function verify(
        bytes32 trustedBeaconBlockRoot,
        Execution.Proof calldata proof
    ) external view returns (bytes32) {
        return Execution.verify(trustedBeaconBlockRoot, proof, _config);
    }

    function receiptsRootOf(
        bytes memory headerRlp
    ) external pure returns (bytes32) {
        return Execution.receiptsRootOf(headerRlp);
    }
}

contract ExecutionTest is Test {
    ExecutionHarness internal _harness;

    // Generalized indices for the Deneb/Electra Ethereum beacon state.
    Execution.BeaconConfig internal _defaultConfig = Execution.BeaconConfig({
        gIndexBlockStateRoot: 11,
        gIndexExecRoot: 1794,
        gIndexBaseStateRoots: 38,
        stateRootsDepth: 13,
        gIndexReceiptsRoot: 1795,
        stateRootsVectorSize: 8192,
        gloasForkSlot: type(uint256).max,
        gIndexBaseStateRootsGloas: 353,
        gIndexLatestBlockHash: 2947
    });

    function setUp() public {
        _harness = new ExecutionHarness();
        _harness.setConfig(_defaultConfig);
    }

    /**
     * @notice Corrupting any single step's proof should revert at exactly that step.
     */
    function testVerifyRevertOnCorruptedProofAtEachStep() public {
        (Execution.Proof memory proof, bytes32 beaconBlockRoot) =
            _buildFullProof(bytes32(uint256(0xCAFE)), 0, 1);
        bytes32 original;

        // Corrupt anchor
        original = proof.anchorBeaconStateProof[0];
        proof.anchorBeaconStateProof[0] = bytes32(uint256(0xBEEF));
        vm.expectRevert("Invalid anchor state root");
        _harness.verify(beaconBlockRoot, proof);
        proof.anchorBeaconStateProof[0] = original;

        // Corrupt history
        original = proof.targetBeaconStateProof[0];
        proof.targetBeaconStateProof[0] = bytes32(uint256(0xBEEF));
        vm.expectRevert("Invalid target beacon state root");
        _harness.verify(beaconBlockRoot, proof);
        proof.targetBeaconStateProof[0] = original;

        // Corrupt execution
        original = proof.targetExecutionHeaderProof[0];
        proof.targetExecutionHeaderProof[0] = bytes32(uint256(0xBEEF));
        vm.expectRevert("Invalid execution root proof");
        _harness.verify(beaconBlockRoot, proof);
        proof.targetExecutionHeaderProof[0] = original;

        // Corrupt receipts
        original = proof.targetReceiptsProof[0];
        proof.targetReceiptsProof[0] = bytes32(uint256(0xBEEF));
        vm.expectRevert("Invalid receipts root");
        _harness.verify(beaconBlockRoot, proof);
        proof.targetReceiptsProof[0] = original;
    }

    /**
     * @notice A branch must have exactly the depth of its generalized index: extra siblings are not
     * ignored.
     */
    function testVerifyRevertsOnABranchLongerThanItsIndex() public {
        (Execution.Proof memory proof, bytes32 beaconBlockRoot) =
            _buildFullProof(bytes32(uint256(0xCAFE)), 0, 1);
        bytes32[] memory longer = new bytes32[](proof.targetReceiptsProof.length + 1);
        for (uint256 i = 0; i < proof.targetReceiptsProof.length; i++) {
            longer[i] = proof.targetReceiptsProof[i];
        }
        proof.targetReceiptsProof = longer;

        vm.expectRevert("Invalid receipts root");
        _harness.verify(beaconBlockRoot, proof);
    }

    /**
     * @notice After Gloas the anchor state has the Gloas layout, while a target before the fork
     * keeps the Fulu one: each leg uses the layout of its own slot.
     */
    function testVerifyAGloasAnchorAndAFuluTarget() public {
        Execution.BeaconConfig memory config = _defaultConfig;
        config.gloasForkSlot = 1000;
        _harness.setConfig(config);
        bytes32 receiptsRoot = bytes32(uint256(0xCAFE));

        (Execution.Proof memory proof, bytes32 beaconBlockRoot) =
            _buildFullProof(receiptsRoot, 999, 1000);
        // Built at the Fulu index of state_roots, it does not verify in a Gloas anchor state.
        vm.expectRevert("Invalid target beacon state root");
        _harness.verify(beaconBlockRoot, proof);

        (proof, beaconBlockRoot) =
            _buildProofWithHistoryBase(receiptsRoot, 999, 1000, config.gIndexBaseStateRootsGloas);
        assertEq(_harness.verify(beaconBlockRoot, proof), receiptsRoot);
    }

    function testReceiptsRootOfRejectsMalformedHeaders() public {
        bytes[] memory five = new bytes[](5);
        for (uint256 i = 0; i < five.length; i++) {
            five[i] = _rlpBytes32(bytes32(i));
        }
        vm.expectRevert("Header has too few fields");
        _harness.receiptsRootOf(_rlpList(five));

        vm.expectRevert("Header is not a list");
        _harness.receiptsRootOf(_rlpBytes32(bytes32(uint256(1))));

        bytes[] memory shortRoot = new bytes[](6);
        for (uint256 i = 0; i < shortRoot.length; i++) {
            shortRoot[i] = _rlpBytes32(bytes32(i));
        }
        shortRoot[5] = hex"8401020304";
        vm.expectRevert("Invalid receipts root in header");
        _harness.receiptsRootOf(_rlpList(shortRoot));
    }

    function testVerifyRevertSlotBoundaries() public {
        Execution.Proof memory proof = _makeBaseProof();
        // target == anchor
        proof.targetSlot = 1000;
        proof.anchorSlot = 1000;
        vm.expectRevert("Target slot must be less than anchor slot");
        _harness.verify(bytes32(uint256(0xBEEF)), proof);
        // target > anchor
        proof.targetSlot = 2000;
        proof.anchorSlot = 1000;
        vm.expectRevert("Target slot must be less than anchor slot");
        _harness.verify(bytes32(uint256(0xBEEF)), proof);
        // exceeds vector size by one
        proof.targetSlot = 1000;
        proof.anchorSlot = 1000 + 8193;
        vm.expectRevert("Target slot is too old");
        _harness.verify(bytes32(uint256(0xBEEF)), proof);
    }

    /**
     * @notice Changing config values should break verification.
     */
    function testVerifyConfigMutation() public {
        (Execution.Proof memory proof, bytes32 beaconBlockRoot) =
            _buildFullProof(bytes32(uint256(0xcafe)), 0, 1);

        // Changing gIndexBlockStateRoot breaks anchor check
        Execution.BeaconConfig memory c1 = _defaultConfig;
        c1.gIndexBlockStateRoot = 12;
        _harness.setConfig(c1);
        vm.expectRevert("Invalid anchor state root");
        _harness.verify(beaconBlockRoot, proof);

        // Changing gIndexReceiptsRoot breaks receipts check
        Execution.BeaconConfig memory c2 = _defaultConfig;
        c2.gIndexReceiptsRoot = 1796;
        _harness.setConfig(c2);
        vm.expectRevert("Invalid receipts root");
        _harness.verify(beaconBlockRoot, proof);

        // Shrink the valid historical slots vector
        _harness.setConfig(_defaultConfig);
        (Execution.Proof memory farProof, bytes32 farRoot) =
            _buildFullProof(bytes32(uint256(0xcafe)), 1000, 1000 + 5000);
        // Verify it passes with default config
        _harness.verify(farRoot, farProof);
        // Now shrink the vector and the verification should fail
        Execution.BeaconConfig memory c3 = _defaultConfig;
        c3.stateRootsVectorSize = 4096;
        _harness.setConfig(c3);
        vm.expectRevert("Target slot is too old");
        _harness.verify(farRoot, farProof);
    }

    /**
     * @notice Invalid slot relationships should always revert.
     */
    function testFuzzVerifyRevertInvalidSlots(
        uint64 target,
        uint64 anchor
    ) public {
        vm.assume(target >= anchor);
        Execution.Proof memory proof = _makeBaseProof();
        proof.targetSlot = target;
        proof.anchorSlot = anchor;
        vm.expectRevert();
        _harness.verify(bytes32(uint256(0xBEEF)), proof);
    }

    /**
     * @notice The receipts root is item 5 of the execution header's RLP list.
     */
    function testReceiptsRootOfReadsTheSixthHeaderField() public view {
        bytes32 receiptsRoot = keccak256("receipts");
        bytes[] memory fields = new bytes[](7);
        for (uint256 i = 0; i < fields.length; i++) {
            fields[i] = _rlpBytes32(keccak256(abi.encode(i)));
        }
        fields[5] = _rlpBytes32(receiptsRoot);
        assertEq(_harness.receiptsRootOf(_rlpList(fields)), receiptsRoot);
    }

    /**
     * @notice Full happy path. Builds valid proofs, sanity checks each step, then verifies the chain.
     */
    function testVerifySuccessFullChainOfTrust() public view {
        bytes32 receiptsRoot = bytes32(uint256(0xcafe));
        uint64 targetSlot = 0;
        uint64 anchorSlot = 1;

        // Build proof
        (Execution.Proof memory proof, bytes32 beaconBlockRoot) =
            _buildFullProof(receiptsRoot, targetSlot, anchorSlot);

        // Verify
        uint256 historyGIndex = _historyGIndex(targetSlot);
        assertEq(
            SSZ.restoreMerkleRoot(
                receiptsRoot, _harness.getConfig().gIndexReceiptsRoot, proof.targetReceiptsProof
            ),
            proof.targetExecutionHeaderRoot
        );
        assertEq(
            SSZ.restoreMerkleRoot(
                proof.targetExecutionHeaderRoot,
                _harness.getConfig().gIndexExecRoot,
                proof.targetExecutionHeaderProof
            ),
            proof.targetBeaconStateRoot
        );
        assertEq(
            SSZ.restoreMerkleRoot(
                proof.targetBeaconStateRoot, historyGIndex, proof.targetBeaconStateProof
            ),
            proof.anchorBeaconStateRoot
        );
        assertEq(
            SSZ.restoreMerkleRoot(
                proof.anchorBeaconStateRoot,
                _harness.getConfig().gIndexBlockStateRoot,
                proof.anchorBeaconStateProof
            ),
            beaconBlockRoot
        );

        _harness.verify(beaconBlockRoot, proof);
    }

    /**
     * @notice Any non-zero receipts root with valid slots should pass.
     */
    function testFuzzVerifySuccessArbitraryInputs(
        bytes32 receiptsRoot,
        uint64 targetSlot,
        uint64 gap
    ) public view {
        vm.assume(receiptsRoot != bytes32(0));
        vm.assume(gap >= 1 && gap <= 8192);
        vm.assume(uint256(targetSlot) + uint256(gap) <= type(uint64).max);
        uint64 anchorSlot = targetSlot + gap;
        (Execution.Proof memory proof, bytes32 beaconBlockRoot) =
            _buildFullProof(receiptsRoot, targetSlot, anchorSlot);
        _harness.verify(beaconBlockRoot, proof);
    }

    function _buildFullProof(
        bytes32 receiptsRoot,
        uint64 targetSlot,
        uint64 anchorSlot
    ) internal view returns (Execution.Proof memory proof, bytes32 beaconBlockRoot) {
        (bytes32 execHeaderRoot, bytes32[] memory receiptsProof) =
            _buildProof(receiptsRoot, _harness.getConfig().gIndexReceiptsRoot);
        (bytes32 targetStateRoot, bytes32[] memory execProof) =
            _buildProof(execHeaderRoot, _harness.getConfig().gIndexExecRoot);

        uint256 historyGIndex = _historyGIndex(targetSlot);
        (bytes32 anchorStateRoot, bytes32[] memory historyProof) =
            _buildProof(targetStateRoot, historyGIndex);
        (bytes32 blockRoot, bytes32[] memory anchorProof) =
            _buildProof(anchorStateRoot, _harness.getConfig().gIndexBlockStateRoot);

        proof = Execution.Proof({
            targetSlot: targetSlot,
            anchorSlot: anchorSlot,
            anchorBeaconStateRoot: anchorStateRoot,
            anchorBeaconStateProof: anchorProof,
            targetBeaconStateRoot: targetStateRoot,
            targetBeaconStateProof: historyProof,
            targetExecutionHeaderRoot: execHeaderRoot,
            targetExecutionHeaderProof: execProof,
            targetReceiptsRoot: receiptsRoot,
            targetReceiptsProof: receiptsProof,
            executionHeaderRlp: ""
        });

        beaconBlockRoot = blockRoot;
    }

    function _buildProofWithHistoryBase(
        bytes32 receiptsRoot,
        uint64 targetSlot,
        uint64 anchorSlot,
        uint256 historyBase
    ) internal view returns (Execution.Proof memory proof, bytes32 beaconBlockRoot) {
        Execution.BeaconConfig memory config = _harness.getConfig();
        proof.targetSlot = targetSlot;
        proof.anchorSlot = anchorSlot;
        proof.targetReceiptsRoot = receiptsRoot;
        (proof.targetExecutionHeaderRoot, proof.targetReceiptsProof) =
            _buildProof(receiptsRoot, config.gIndexReceiptsRoot);
        (proof.targetBeaconStateRoot, proof.targetExecutionHeaderProof) =
            _buildProof(proof.targetExecutionHeaderRoot, config.gIndexExecRoot);
        (proof.anchorBeaconStateRoot, proof.targetBeaconStateProof) = _buildProof(
            proof.targetBeaconStateRoot,
            (historyBase << config.stateRootsDepth) + targetSlot % config.stateRootsVectorSize
        );
        (beaconBlockRoot, proof.anchorBeaconStateProof) =
            _buildProof(proof.anchorBeaconStateRoot, config.gIndexBlockStateRoot);
    }

    function _historyGIndex(
        uint64 targetSlot
    ) internal view returns (uint256) {
        uint256 vectorIndex = uint256(targetSlot) % _harness.getConfig().stateRootsVectorSize;
        return (_harness.getConfig().gIndexBaseStateRoots << _harness.getConfig().stateRootsDepth)
            + vectorIndex;
    }

    function _buildProof(
        bytes32 leaf,
        uint256 gIndex
    ) internal pure returns (bytes32 root, bytes32[] memory proof) {
        uint256 depth = Math.log2(gIndex);
        proof = new bytes32[](depth);

        bytes32 computed = leaf;
        uint256 index = gIndex;

        for (uint256 i = 0; i < depth; i++) {
            proof[i] = keccak256(abi.encodePacked("sibling", gIndex, i));

            if (index % 2 == 1) {
                computed = sha256(abi.encodePacked(proof[i], computed));
            } else {
                computed = sha256(abi.encodePacked(computed, proof[i]));
            }
            index /= 2;
        }

        root = computed;
    }

    function _makeBaseProof() internal pure returns (Execution.Proof memory) {
        bytes32[] memory dummyBranch = new bytes32[](1);
        dummyBranch[0] = bytes32(uint256(0xFF));

        return Execution.Proof({
            targetSlot: 500,
            anchorSlot: 600,
            anchorBeaconStateRoot: bytes32(uint256(0xA1)),
            anchorBeaconStateProof: dummyBranch,
            targetBeaconStateRoot: bytes32(uint256(0xB2)),
            targetBeaconStateProof: dummyBranch,
            targetExecutionHeaderRoot: bytes32(uint256(0xC3)),
            targetExecutionHeaderProof: dummyBranch,
            targetReceiptsRoot: bytes32(uint256(0xD4)),
            targetReceiptsProof: dummyBranch,
            executionHeaderRlp: ""
        });
    }

    /// @dev RLP of a 32-byte string.
    function _rlpBytes32(
        bytes32 value
    ) internal pure returns (bytes memory) {
        return abi.encodePacked(uint8(0x80 + 32), value);
    }

    /// @dev RLP of a list of RLP items, whose payload is under 64 KiB.
    function _rlpList(
        bytes[] memory items
    ) internal pure returns (bytes memory payload) {
        for (uint256 i = 0; i < items.length; i++) {
            payload = bytes.concat(payload, items[i]);
        }
        if (payload.length < 56) {
            return bytes.concat(bytes1(uint8(0xc0 + payload.length)), payload);
        }
        if (payload.length < 256) {
            return bytes.concat(hex"f8", bytes1(uint8(payload.length)), payload);
        }
        return bytes.concat(hex"f9", bytes2(uint16(payload.length)), payload);
    }
}
