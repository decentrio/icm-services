// (c) 2026, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

// SPDX-License-Identifier: LicenseRef-Ecosystem

pragma solidity ^0.8.30;

import {Test} from "forge-std/Test.sol";
import {LightClientAdapter} from "../LightClientAdapter.sol";
import {EthereumAdapter} from "../EthereumAdapter.sol";
import {EthereumLogVerifier} from "../EthereumLogVerifier.sol";
import {IBeaconRootSource} from "../IBeaconRootSource.sol";
import {ZKStateManager, Journal} from "../ZKStateManager.sol";
// Aliased: forge-std's Test declares its own Receipt struct.
import {Consensus, Execution, Receipt as LogReceipt} from "../StateManagerLibrary.sol";
import {EthereumAdapterTest, ZKAdapterTestVerifier} from "./ZKAdapterTests.t.sol";

/// @dev A beacon root source whose roots the test sets, standing in for a light client.
contract TestBeaconRootSource is IBeaconRootSource {
    mapping(uint64 slot => bytes32 root) private _roots;

    function setRoot(
        uint64 slot,
        bytes32 root
    ) external {
        _roots[slot] = root;
    }

    function beaconRootAt(
        uint64 slot
    ) external view returns (bytes32) {
        return _roots[slot];
    }
}

/// @dev The behaviour LightClientAdapter shares with every Ethereum adapter.
contract LightClientAdapterSharedTest is EthereumAdapterTest {
    function _newAdapter() internal override returns (EthereumAdapter) {
        Execution.BeaconConfig memory config;
        return new LightClientAdapter(1, config, new TestBeaconRootSource());
    }
}

/**
 * @dev LightClientAdapter proves Ethereum logs against the beacon roots of a source, such as the
 * ZK light client, using a real Ethereum mainnet log (a USDC transfer) and its proofs.
 */
contract LightClientAdapterTest is Test {
    uint64 private constant _START_EPOCH = 1_000;
    uint256 private constant _GENESIS_TIME = 1_606_824_023;

    struct Fixture {
        bytes32 anchorRoot;
        Execution.Proof execProof;
        LogReceipt.Proof logProof;
    }

    event ZKEventImported(
        uint256 indexed sourceChainId,
        uint256 indexed beaconSlot,
        address indexed emitter,
        bytes32 executionRoot,
        uint256 logIndex
    );

    function testVerifiesARealMainnetLogAgainstTheSourcesRoot() public {
        Fixture memory f = _fixture();
        TestBeaconRootSource source = new TestBeaconRootSource();
        source.setRoot(f.execProof.anchorSlot, f.anchorRoot);
        LightClientAdapter adapter = new LightClientAdapter(1, _fuluConfig(), source);

        vm.expectEmit(true, true, true, true, address(adapter));
        emit ZKEventImported(
            1,
            f.execProof.targetSlot,
            f.logProof.expectedEmitter,
            f.execProof.targetExecutionHeaderRoot,
            f.logProof.logIndex
        );
        bytes memory logData = adapter.verifyLogAndExtract(f.execProof, f.logProof);
        // A Transfer log carries the amount as its only non-indexed field.
        assertEq(logData.length, 32);
    }

    function testRejectsAnAnchorSlotTheSourceHasNoRootFor() public {
        Fixture memory f = _fixture();
        LightClientAdapter adapter =
            new LightClientAdapter(1, _fuluConfig(), new TestBeaconRootSource());

        vm.expectRevert("Anchor slot is undefined");
        adapter.verifyLogAndExtract(f.execProof, f.logProof);
    }

    function testRejectsAProofOfAnotherBeaconBlock() public {
        Fixture memory f = _fixture();
        TestBeaconRootSource source = new TestBeaconRootSource();
        source.setRoot(f.execProof.anchorSlot, keccak256("another block"));
        LightClientAdapter adapter = new LightClientAdapter(1, _fuluConfig(), source);

        vm.expectRevert();
        adapter.verifyLogAndExtract(f.execProof, f.logProof);
    }

    function testReadsRootsFromItsSourceAtVerificationTime() public {
        Fixture memory f = _fixture();
        TestBeaconRootSource source = new TestBeaconRootSource();
        LightClientAdapter adapter = new LightClientAdapter(1, _fuluConfig(), source);
        assertEq(address(adapter.beaconRootSource()), address(source));

        source.setRoot(f.execProof.anchorSlot, f.anchorRoot);
        adapter.verifyLogAndExtract(f.execProof, f.logProof);
    }

    /// @dev The Boundless state manager is a beacon root source too, so ICM can run either source.
    function testVerifiesAgainstRootsConfirmedByTheBoundlessStateManager() public {
        Fixture memory f = _fixture();
        ZKStateManager manager = _stateManager();
        Journal memory journal;
        journal.postState.finalizedCheckpoint.root = f.anchorRoot;
        manager.manualTransition(abi.encode(journal), f.execProof.anchorSlot);
        assertEq(manager.beaconRootAt(f.execProof.anchorSlot), f.anchorRoot);

        LightClientAdapter adapter = new LightClientAdapter(1, _fuluConfig(), manager);
        adapter.verifyLogAndExtract(f.execProof, f.logProof);
    }

    function testRejectsAZeroSourceChainId() public {
        TestBeaconRootSource source = new TestBeaconRootSource();
        vm.expectRevert("Invalid chain ID");
        new LightClientAdapter(0, _fuluConfig(), source);
    }

    function testRejectsAMissingSource() public {
        vm.expectRevert("Invalid beacon root source");
        new LightClientAdapter(1, _fuluConfig(), IBeaconRootSource(address(0)));
    }

    /// @dev Generalized indices of the Fulu beacon state and block layouts.
    function _fuluConfig() private pure returns (Execution.BeaconConfig memory) {
        return Execution.BeaconConfig({
            gIndexBlockStateRoot: 11,
            gIndexExecRoot: 88,
            gIndexBaseStateRoots: 70,
            stateRootsDepth: 13,
            gIndexReceiptsRoot: 35,
            stateRootsVectorSize: 8192
        });
    }

    function _stateManager() private returns (ZKStateManager) {
        vm.warp(_GENESIS_TIME + uint256(_START_EPOCH + 2) * 32 * 12);
        Consensus.State memory start;
        start.finalizedCheckpoint.epoch = _START_EPOCH;
        return new ZKStateManager({
            newSourceChainId: 1,
            startingState: start,
            beaconConfig: _fuluConfig(),
            genesisTime_: _GENESIS_TIME,
            permissibleTimespan_: 86_400,
            verifier_: address(new ZKAdapterTestVerifier()),
            imageID_: bytes32(uint256(1)),
            admin: address(this),
            superAdmin: address(this)
        });
    }

    /// @dev The mainnet fixture of the Go e2e test: a USDC transfer log and its proofs.
    function _fixture() private view returns (Fixture memory f) {
        string memory json = vm.readFile("tests/testdata/ethereum_fixture.json");
        f.anchorRoot = vm.parseJsonBytes32(json, ".anchorBeaconBlockRoot");
        f.execProof = Execution.Proof({
            anchorSlot: uint64(vm.parseJsonUint(json, ".executionProof.anchorSlot")),
            targetSlot: uint64(vm.parseJsonUint(json, ".executionProof.targetSlot")),
            anchorBeaconStateRoot: vm.parseJsonBytes32(
                json, ".executionProof.anchorBeaconStateRoot"
            ),
            anchorBeaconStateProof: vm.parseJsonBytes32Array(
                json, ".executionProof.anchorBeaconStateProof"
            ),
            targetBeaconStateRoot: vm.parseJsonBytes32(
                json, ".executionProof.targetBeaconStateRoot"
            ),
            targetBeaconStateProof: vm.parseJsonBytes32Array(
                json, ".executionProof.targetBeaconStateProof"
            ),
            targetExecutionHeaderRoot: vm.parseJsonBytes32(
                json, ".executionProof.targetExecutionHeaderRoot"
            ),
            targetExecutionHeaderProof: vm.parseJsonBytes32Array(
                json, ".executionProof.targetExecutionHeaderProof"
            ),
            targetReceiptsRoot: vm.parseJsonBytes32(json, ".executionProof.targetReceiptsRoot"),
            targetReceiptsProof: vm.parseJsonBytes32Array(
                json, ".executionProof.targetReceiptsProof"
            )
        });
        f.logProof = LogReceipt.Proof({
            proof: vm.parseJsonBytesArray(json, ".receiptProof.proof"),
            key: vm.parseJsonBytes(json, ".receiptProof.key"),
            value: vm.parseJsonBytes(json, ".receiptProof.value"),
            logIndex: vm.parseJsonUint(json, ".receiptProof.logIndex"),
            expectedEmitter: vm.parseJsonAddress(json, ".receiptProof.expectedEmitter"),
            expectedTopic0: vm.parseJsonBytes32(json, ".receiptProof.expectedTopic0")
        });
    }
}
