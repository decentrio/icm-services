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
    struct Fixture {
        bytes32 anchorRoot;
        Execution.Proof execProof;
        LogReceipt.Proof logProof;
    }

    uint64 private constant _START_EPOCH = 1_000;
    uint256 private constant _GENESIS_TIME = 1_606_824_023;
    uint256 private constant _SEPOLIA_CHAIN_ID = 11_155_111;
    // GLOAS_FORK_EPOCH * SLOTS_PER_EPOCH on Sepolia.
    uint256 private constant _SEPOLIA_GLOAS_FORK_SLOT = 353_024 * 32;

    function testVerifiesARealMainnetLogAgainstTheSourcesRoot() public {
        Fixture memory f = _fixture();
        TestBeaconRootSource source = new TestBeaconRootSource();
        source.setRoot(f.execProof.anchorSlot, f.anchorRoot);
        LightClientAdapter adapter = new LightClientAdapter(1, _fuluConfig(), source);

        vm.expectEmit(true, true, true, true, address(adapter));
        emit EthereumLogVerifier.ZKEventImported(
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

    /**
     * @dev A real Sepolia log after Gloas, which took the execution payload header out of the
     * beacon state: the execution block is reached through the state of the first block that built
     * on its payload, whose latest_block_hash is the hash of the execution header given in full.
     */
    function testVerifiesARealSepoliaLogAfterGloas() public {
        Fixture memory f = _gloasFixture();
        TestBeaconRootSource source = new TestBeaconRootSource();
        source.setRoot(f.execProof.anchorSlot, f.anchorRoot);
        LightClientAdapter adapter =
            new LightClientAdapter(_SEPOLIA_CHAIN_ID, _sepoliaConfig(), source);

        string memory json = vm.readFile("tests/testdata/sepolia_gloas_fixture.json");
        vm.expectEmit(true, true, true, true, address(adapter));
        emit EthereumLogVerifier.ZKEventImported(
            _SEPOLIA_CHAIN_ID,
            vm.parseJsonUint(json, ".metadata.targetSlot"),
            vm.parseJsonAddress(json, ".receiptProof.expectedEmitter"),
            vm.parseJsonBytes32(json, ".metadata.blockHash"),
            0
        );
        bytes memory logData = adapter.verifyLogAndExtract(f.execProof, f.logProof);
        // The transfer moved 0.5 USDC (6 decimals).
        assertEq(abi.decode(logData, (uint256)), 500_000);
    }

    /// @dev The layout is the one of each slot's fork: a Gloas proof does not verify on a chain
    /// where Gloas is not scheduled.
    function testRejectsAGloasProofWhereGloasIsNotScheduled() public {
        Fixture memory f = _gloasFixture();
        LightClientAdapter adapter = _sepoliaAdapter(f, _fuluConfig());

        vm.expectRevert("Invalid target beacon state root");
        adapter.verifyLogAndExtract(f.execProof, f.logProof);
    }

    /// @dev ...nor a Fulu proof where its slots are in Gloas. Its branch is shorter than the Gloas
    /// index's depth, which must revert like any other invalid branch.
    function testRejectsAFuluProofWhereItsSlotsAreInGloas() public {
        Fixture memory f = _fixture();
        Execution.BeaconConfig memory config = _fuluConfig();
        config.gloasForkSlot = 0;
        TestBeaconRootSource source = new TestBeaconRootSource();
        source.setRoot(f.execProof.anchorSlot, f.anchorRoot);
        LightClientAdapter adapter = new LightClientAdapter(1, config, source);

        vm.expectRevert("Invalid target beacon state root");
        adapter.verifyLogAndExtract(f.execProof, f.logProof);
    }

    function testRejectsAGloasBlockHashNotInTheTargetState() public {
        Fixture memory f = _gloasFixture();
        LightClientAdapter adapter = _sepoliaAdapter(f, _sepoliaConfig());
        f.execProof.targetExecutionHeaderProof[0] ^= bytes32(uint256(1));

        vm.expectRevert("Invalid execution block hash proof");
        adapter.verifyLogAndExtract(f.execProof, f.logProof);
    }

    function testRejectsAHeaderThatIsNotTheBlockHashs() public {
        Fixture memory f = _gloasFixture();
        LightClientAdapter adapter = _sepoliaAdapter(f, _sepoliaConfig());
        f.execProof.executionHeaderRlp[200] ^= 0x01;

        vm.expectRevert("Header does not match block hash");
        adapter.verifyLogAndExtract(f.execProof, f.logProof);
    }

    /**
     * @dev In Gloas the state at the payload's own slot holds the parent's block hash: the payload
     * is only known to be included once a later block builds on it. A proof through that state
     * proves the parent block, whose receipts root does not hold the log.
     */
    function testRejectsTheStateOfThePayloadsOwnSlot() public {
        Fixture memory f = _gloasFixture();
        LightClientAdapter adapter = _sepoliaAdapter(f, _sepoliaConfig());
        string memory json = vm.readFile("tests/testdata/sepolia_gloas_fixture.json");
        string memory s = ".negative.wrongTargetSlotIsPayloadSlot";
        f.execProof.targetSlot = uint64(vm.parseJsonUint(json, string.concat(s, ".targetSlot")));
        f.execProof.targetBeaconStateRoot =
            vm.parseJsonBytes32(json, string.concat(s, ".targetBeaconStateRoot"));
        f.execProof.targetBeaconStateProof =
            vm.parseJsonBytes32Array(json, string.concat(s, ".targetBeaconStateProof"));
        f.execProof.targetExecutionHeaderRoot =
            vm.parseJsonBytes32(json, string.concat(s, ".executionBlockHash"));
        f.execProof.targetExecutionHeaderProof =
            vm.parseJsonBytes32Array(json, string.concat(s, ".executionBlockHashProof"));

        // The target's own header does not hash to the parent's block hash.
        vm.expectRevert("Header does not match block hash");
        adapter.verifyLogAndExtract(f.execProof, f.logProof);

        // The parent's header does, but its receipts root is not the claimed one...
        f.execProof.executionHeaderRlp =
            vm.parseJsonBytes(json, string.concat(s, ".parentHeaderRlp"));
        vm.expectRevert("Receipts root mismatch");
        adapter.verifyLogAndExtract(f.execProof, f.logProof);

        // ...and the log's receipt is not under the parent's receipts root.
        f.execProof.targetReceiptsRoot = Execution.receiptsRootOf(f.execProof.executionHeaderRlp);
        vm.expectRevert();
        adapter.verifyLogAndExtract(f.execProof, f.logProof);
    }

    /// @dev A Gloas proof has no receipts branch: the receipts root comes from the header.
    function testRejectsAGloasProofWithAReceiptsBranch() public {
        Fixture memory f = _gloasFixture();
        LightClientAdapter adapter = _sepoliaAdapter(f, _sepoliaConfig());
        f.execProof.targetReceiptsProof = new bytes32[](1);

        vm.expectRevert("Unexpected receipts proof");
        adapter.verifyLogAndExtract(f.execProof, f.logProof);
    }

    /// @dev A Fulu proof has no execution header: the receipts root comes from its branch.
    function testRejectsAFuluProofWithAnExecutionHeader() public {
        Fixture memory f = _fixture();
        TestBeaconRootSource source = new TestBeaconRootSource();
        source.setRoot(f.execProof.anchorSlot, f.anchorRoot);
        LightClientAdapter adapter = new LightClientAdapter(1, _fuluConfig(), source);
        f.execProof.executionHeaderRlp = hex"c0";

        vm.expectRevert("Unexpected header RLP");
        adapter.verifyLogAndExtract(f.execProof, f.logProof);
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

    /// @dev The Boundless state manager proves logs against the roots its own transitions confirm.
    function testBoundlessStateManagerVerifiesAgainstItsOwnConfirmedRoots() public {
        Fixture memory f = _fixture();
        ZKStateManager manager = _stateManager();
        vm.expectRevert("Anchor slot is undefined");
        manager.verifyLogAndExtract(f.execProof, f.logProof);

        Journal memory journal;
        journal.postState.finalizedCheckpoint.root = f.anchorRoot;
        manager.manualTransition(abi.encode(journal), f.execProof.anchorSlot);
        manager.verifyLogAndExtract(f.execProof, f.logProof);
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

    function _sepoliaAdapter(
        Fixture memory f,
        Execution.BeaconConfig memory config
    ) private returns (LightClientAdapter) {
        TestBeaconRootSource source = new TestBeaconRootSource();
        source.setRoot(f.execProof.anchorSlot, f.anchorRoot);
        return new LightClientAdapter(_SEPOLIA_CHAIN_ID, config, source);
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
            ),
            executionHeaderRlp: ""
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

    /// @dev The Sepolia fixture after Gloas: a USDC transfer log and its proofs.
    function _gloasFixture() private view returns (Fixture memory f) {
        string memory json = vm.readFile("tests/testdata/sepolia_gloas_fixture.json");
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
                json, ".executionProof.executionBlockHash"
            ),
            targetExecutionHeaderProof: vm.parseJsonBytes32Array(
                json, ".executionProof.executionBlockHashProof"
            ),
            targetReceiptsRoot: vm.parseJsonBytes32(json, ".executionProof.receiptsRoot"),
            targetReceiptsProof: new bytes32[](0),
            executionHeaderRlp: vm.parseJsonBytes(json, ".executionProof.executionHeaderRlp")
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

    /// @dev Generalized indices of the Fulu beacon state and block layouts; Gloas not scheduled.
    function _fuluConfig() private pure returns (Execution.BeaconConfig memory config) {
        config = _layouts();
        config.gloasForkSlot = type(uint256).max;
    }

    /// @dev Sepolia: the Fulu layouts, and Gloas's from its fork.
    function _sepoliaConfig() private pure returns (Execution.BeaconConfig memory config) {
        config = _layouts();
        config.gloasForkSlot = _SEPOLIA_GLOAS_FORK_SLOT;
    }

    /// @dev Generalized indices of the Fulu and Gloas beacon state and block layouts.
    function _layouts() private pure returns (Execution.BeaconConfig memory) {
        return Execution.BeaconConfig({
            gIndexBlockStateRoot: 11,
            gIndexExecRoot: 88,
            gIndexBaseStateRoots: 70,
            stateRootsDepth: 13,
            gIndexReceiptsRoot: 35,
            stateRootsVectorSize: 8192,
            gloasForkSlot: 0,
            gIndexBaseStateRootsGloas: 353,
            gIndexLatestBlockHash: 2947
        });
    }
}
