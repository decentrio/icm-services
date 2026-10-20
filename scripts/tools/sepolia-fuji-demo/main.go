// Copyright (C) 2026, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

// THIS IS AN EXAMPLE OF UNAUDITED CODE. DO NOT USE THIS IN PRODUCTION.

// Command sepolia-fuji-demo proves a real Sepolia ERC-20 transfer on an Avalanche chain against
// the ZK light client. It finds a recent transfer of the token that the light client's latest
// finalized beacon root covers, builds the Gloas execution proof and the receipt proof from
// Sepolia beacon and execution nodes, and has a LightClientAdapter verify the transfer's log.
//
// The light client must already follow Sepolia (see avax-light-client's cmd/updater). The
// sender's private key is read from the DEMO_KEY environment variable.
package main

import (
	"context"
	"errors"
	"flag"
	"fmt"
	"math/big"
	"os"
	"sort"
	"strings"
	"time"

	lightclientadapter "github.com/ava-labs/icm-services/abi-bindings/go/verifiers/ethereum/LightClientAdapter"
	zkadapter "github.com/ava-labs/icm-services/abi-bindings/go/verifiers/ethereum/ZKAdapter"
	"github.com/ava-labs/icm-services/messages/zk/proofs"
	"github.com/ava-labs/icm-services/messages/zk/proofs/beacon"
	"github.com/ava-labs/libevm/accounts/abi"
	"github.com/ava-labs/libevm/accounts/abi/bind"
	"github.com/ava-labs/libevm/common"
	"github.com/ava-labs/libevm/common/hexutil"
	"github.com/ava-labs/libevm/crypto"
	"github.com/ava-labs/libevm/ethclient"
	"github.com/ava-labs/libevm/rpc"
)

const (
	sepoliaChainID = 11155111
	// Blocks searched for transfers below the execution chain's finalized block.
	searchDepth = 100
)

// transferTopic is keccak256("Transfer(address,address,uint256)").
var transferTopic = crypto.Keccak256Hash([]byte("Transfer(address,address,uint256)"))

// beaconRootSourceABI is the part of the light client the demo reads.
const beaconRootSourceABI = `[
	{"type":"function","name":"latestFinalizedSlot","stateMutability":"view","inputs":[],"outputs":[{"type":"uint64"}]},
	{"type":"function","name":"beaconRootAt","stateMutability":"view",
	 "inputs":[{"type":"uint64"}],"outputs":[{"type":"bytes32"}]}
]`

type config struct {
	rpc, sepoliaRPC, beaconURL string
	lightClient, adapter       common.Address
	token                      common.Address
	tx                         common.Hash
	wait                       time.Duration
}

func main() {
	var c config
	var lightClient, adapter, token, tx string
	flag.StringVar(&c.rpc, "rpc", "http://127.0.0.1:8547", "RPC of the Avalanche chain the light client is on")
	flag.StringVar(&c.sepoliaRPC, "sepolia-rpc", "https://ethereum-sepolia-rpc.publicnode.com", "Sepolia execution RPC")
	flag.StringVar(&c.beaconURL, "beacon", "https://lodestar-sepolia.chainsafe.io", "Sepolia beacon API")
	flag.StringVar(&lightClient, "light-client", "", "ZK light client address")
	flag.StringVar(&adapter, "adapter", "", "LightClientAdapter address; deployed if empty")
	flag.StringVar(&token, "token", "0x1c7d4b196cb0c7b01d743fbc6116a902379c7238", "ERC-20 on Sepolia (default: USDC)")
	flag.StringVar(&tx, "tx", "", "Sepolia transaction to prove; the latest covered transfer if empty")
	flag.DurationVar(&c.wait, "wait", 30*time.Minute, "how long to wait for the light client to cover a transfer")
	flag.Parse()
	c.lightClient = common.HexToAddress(lightClient)
	c.adapter = common.HexToAddress(adapter)
	c.token = common.HexToAddress(token)
	if tx != "" {
		c.tx = common.HexToHash(tx)
	}
	if err := run(context.Background(), c); err != nil {
		fmt.Fprintln(os.Stderr, "error:", err)
		os.Exit(1)
	}
}

func run(ctx context.Context, c config) error {
	key, err := crypto.HexToECDSA(strings.TrimPrefix(os.Getenv("DEMO_KEY"), "0x"))
	if err != nil {
		return fmt.Errorf("DEMO_KEY: %w", err)
	}
	avax, err := ethclient.DialContext(ctx, c.rpc)
	if err != nil {
		return err
	}
	chainID, err := avax.ChainID(ctx)
	if err != nil {
		return err
	}
	auth, err := bind.NewKeyedTransactorWithChainID(key, chainID)
	if err != nil {
		return err
	}
	el, err := rpc.DialContext(ctx, c.sepoliaRPC)
	if err != nil {
		return err
	}
	parsed, err := abi.JSON(strings.NewReader(beaconRootSourceABI))
	if err != nil {
		return err
	}
	source := bind.NewBoundContract(c.lightClient, parsed, avax, avax, avax)

	adapterAddress := c.adapter
	if adapterAddress == (common.Address{}) {
		if adapterAddress, err = deployAdapter(ctx, auth, avax, c.lightClient); err != nil {
			return err
		}
	}
	adapter, err := lightclientadapter.NewLightClientAdapter(adapterAddress, avax)
	if err != nil {
		return err
	}
	fmt.Printf("chain %s: light client %s, adapter %s\n", chainID, c.lightClient, adapterAddress)

	builder := proofs.NewProofBuilderWithForks(beacon.NewBeaconClient(c.beaconURL), el, proofs.SepoliaForkConfig)
	t, execProof, err := proveTransfer(ctx, c, source, el, builder)
	if err != nil {
		return err
	}
	receiptProof, err := proofs.BuildReceiptProof(ctx, el, t.blockNumber, t.txHash, t.receiptLogIndex)
	if err != nil {
		return err
	}
	fmt.Printf("receipt proof: %d trie nodes\n", len(receiptProof.Proof))

	exec := lightclientadapter.ExecutionProof(*execProof)
	logProof := lightclientadapter.ReceiptProof{
		Proof:           receiptProof.Proof,
		Key:             receiptProof.Key,
		Value:           receiptProof.Value,
		LogIndex:        new(big.Int).SetUint64(uint64(receiptProof.LogIndex)),
		ExpectedEmitter: receiptProof.ExpectedEmitter,
		ExpectedTopic0:  receiptProof.ExpectedTopic0,
	}
	auth.Context = ctx
	tx, err := adapter.VerifyLogAndExtract(auth, exec, logProof)
	if err != nil {
		return fmt.Errorf("verifyLogAndExtract: %w", err)
	}
	receipt, err := bind.WaitMined(ctx, avax, tx)
	if err != nil {
		return err
	}
	if receipt.Status != 1 {
		return fmt.Errorf("verifyLogAndExtract %s reverted", tx.Hash())
	}
	for _, l := range receipt.Logs {
		ev, err := adapter.ParseZKEventImported(*l)
		if err != nil {
			continue
		}
		fmt.Printf("\nverified on chain %s in tx %s (%d gas)\n", chainID, tx.Hash(), receipt.GasUsed)
		fmt.Printf("  ZKEventImported: source chain %s, beacon slot %s, emitter %s, execution block %s, log %s\n",
			ev.SourceChainId, ev.BeaconSlot, ev.Emitter, common.Hash(ev.ExecutionRoot), ev.LogIndex)
	}
	fmt.Printf("  Sepolia transfer %s: %s -> %s, amount %s (raw units)\n", t.txHash, t.from, t.to, t.amount)
	return nil
}

// deployAdapter deploys a LightClientAdapter for Sepolia, whose Gloas fork starts at epoch 353024.
func deployAdapter(
	ctx context.Context,
	auth *bind.TransactOpts,
	avax *ethclient.Client,
	lightClient common.Address,
) (common.Address, error) {
	cfg := lightclientadapter.ExecutionBeaconConfig{
		GIndexBlockStateRoot:      big.NewInt(11),
		GIndexExecRoot:            big.NewInt(88),
		GIndexBaseStateRoots:      big.NewInt(70),
		StateRootsDepth:           big.NewInt(13),
		GIndexReceiptsRoot:        big.NewInt(35),
		StateRootsVectorSize:      big.NewInt(8192),
		GloasForkSlot:             new(big.Int).SetUint64(proofs.SepoliaForkConfig.GloasForkSlot),
		GIndexBaseStateRootsGloas: big.NewInt(353),
		GIndexLatestBlockHash:     big.NewInt(2947),
	}
	auth.Context = ctx
	address, tx, _, err := lightclientadapter.DeployLightClientAdapter(
		auth, avax, big.NewInt(sepoliaChainID), cfg, lightClient)
	if err != nil {
		return common.Address{}, fmt.Errorf("deploy LightClientAdapter: %w", err)
	}
	if _, err := bind.WaitDeployed(ctx, avax, tx); err != nil {
		return common.Address{}, err
	}
	return address, nil
}

type transfer struct {
	txHash          common.Hash
	blockNumber     uint64
	blockHash       common.Hash
	payloadSlot     uint64
	receiptLogIndex uint
	from, to        common.Address
	amount          *big.Int
}

type rpcLog struct {
	Address     common.Address `json:"address"`
	Topics      []common.Hash  `json:"topics"`
	Data        hexutil.Bytes  `json:"data"`
	BlockNumber hexutil.Uint64 `json:"blockNumber"`
	BlockHash   common.Hash    `json:"blockHash"`
	TxHash      common.Hash    `json:"transactionHash"`
	LogIndex    hexutil.Uint64 `json:"logIndex"`
}

// proveTransfer waits until the light client's latest finalized root covers a transfer of the
// token, and builds the transfer's execution proof against that root.
func proveTransfer(
	ctx context.Context,
	c config,
	source *bind.BoundContract,
	el *rpc.Client,
	builder *proofs.ProofBuilder,
) (*transfer, *zkadapter.ExecutionProof, error) {
	deadline := time.Now().Add(c.wait)
	for {
		anchor, anchorRoot, err := latestRoot(ctx, source)
		if err != nil {
			return nil, nil, err
		}
		candidates, err := transfers(ctx, c, el)
		if err != nil {
			return nil, nil, err
		}
		for _, t := range candidates {
			if t.payloadSlot+1 >= anchor {
				continue // the payload is committed by a block after it, which must be before the anchor
			}
			fmt.Printf("light client anchor: slot %d, root %s\n", anchor, anchorRoot)
			fmt.Printf("transfer %s in execution block %d (payload slot %d)\n", t.txHash, t.blockNumber, t.payloadSlot)
			start := time.Now()
			proof, err := builder.BuildExecutionProofForBlock(ctx, anchor, anchorRoot,
				proofs.ExecutionTarget{BlockHash: t.blockHash, PayloadSlot: t.payloadSlot})
			if errors.Is(err, proofs.ErrNotCommittedBeforeAnchor) || errors.Is(err, proofs.ErrPayloadNotIncluded) {
				fmt.Println("  skipped:", err)
				continue
			}
			if err != nil {
				return nil, nil, err
			}
			fmt.Printf("execution proof in %s: target slot %d, %d + %d + %d SSZ siblings, %d-byte header\n",
				time.Since(start).Round(time.Second), proof.TargetSlot, len(proof.AnchorBeaconStateProof),
				len(proof.TargetBeaconStateProof), len(proof.TargetExecutionHeaderProof), len(proof.ExecutionHeaderRlp))
			return t, proof, nil
		}
		if time.Now().After(deadline) {
			return nil, nil, fmt.Errorf("no transfer covered by the light client's root at slot %d", anchor)
		}
		fmt.Printf("waiting for the light client to finalize past a transfer (anchor slot %d)\n", anchor)
		select {
		case <-ctx.Done():
			return nil, nil, ctx.Err()
		case <-time.After(time.Minute):
		}
	}
}

func latestRoot(ctx context.Context, source *bind.BoundContract) (uint64, common.Hash, error) {
	var out []any
	if err := source.Call(&bind.CallOpts{Context: ctx}, &out, "latestFinalizedSlot"); err != nil {
		return 0, common.Hash{}, err
	}
	slot := out[0].(uint64)
	out = nil
	if err := source.Call(&bind.CallOpts{Context: ctx}, &out, "beaconRootAt", slot); err != nil {
		return 0, common.Hash{}, err
	}
	return slot, common.Hash(out[0].([32]byte)), nil
}

// transfers returns the token's transfers, newest first: the given transaction's, or those in
// the blocks below the execution chain's finalized block.
func transfers(ctx context.Context, c config, el *rpc.Client) ([]*transfer, error) {
	var logs []rpcLog
	if c.tx != (common.Hash{}) {
		var receipt *struct {
			Logs []rpcLog `json:"logs"`
		}
		if err := el.CallContext(ctx, &receipt, "eth_getTransactionReceipt", c.tx); err != nil {
			return nil, err
		}
		if receipt == nil {
			return nil, fmt.Errorf("transaction %s not found", c.tx)
		}
		logs = receipt.Logs
	} else {
		var finalized *struct {
			Number hexutil.Uint64 `json:"number"`
		}
		if err := el.CallContext(ctx, &finalized, "eth_getBlockByNumber", "finalized", false); err != nil {
			return nil, err
		}
		filter := map[string]any{
			"address":   c.token,
			"topics":    []any{transferTopic},
			"fromBlock": hexutil.Uint64(uint64(finalized.Number) - searchDepth),
			"toBlock":   finalized.Number,
		}
		if err := el.CallContext(ctx, &logs, "eth_getLogs", filter); err != nil {
			return nil, err
		}
	}
	sort.Slice(logs, func(i, j int) bool { return logs[i].BlockNumber > logs[j].BlockNumber })

	var out []*transfer
	for _, l := range logs {
		if l.Address != c.token || len(l.Topics) != 3 || l.Topics[0] != transferTopic {
			continue
		}
		t, err := describe(ctx, el, l)
		if err != nil {
			return nil, err
		}
		out = append(out, t)
	}
	return out, nil
}

// describe completes a transfer log with its block's payload slot and its index in its receipt.
func describe(ctx context.Context, el *rpc.Client, l rpcLog) (*transfer, error) {
	var block *struct {
		SlotNumber *hexutil.Uint64 `json:"slotNumber"`
	}
	if err := el.CallContext(ctx, &block, "eth_getBlockByHash", l.BlockHash, false); err != nil {
		return nil, err
	}
	if block == nil || block.SlotNumber == nil {
		return nil, fmt.Errorf("block %s has no slotNumber", l.BlockHash)
	}
	var receipt *struct {
		Logs []rpcLog `json:"logs"`
	}
	if err := el.CallContext(ctx, &receipt, "eth_getTransactionReceipt", l.TxHash); err != nil {
		return nil, err
	}
	index := -1
	for i, rl := range receipt.Logs {
		if rl.LogIndex == l.LogIndex {
			index = i
		}
	}
	if index < 0 {
		return nil, fmt.Errorf("log %d not in the receipt of %s", l.LogIndex, l.TxHash)
	}
	return &transfer{
		txHash:          l.TxHash,
		blockNumber:     uint64(l.BlockNumber),
		blockHash:       l.BlockHash,
		payloadSlot:     uint64(*block.SlotNumber),
		receiptLogIndex: uint(index),
		from:            common.BytesToAddress(l.Topics[1].Bytes()),
		to:              common.BytesToAddress(l.Topics[2].Bytes()),
		amount:          new(big.Int).SetBytes(l.Data),
	}, nil
}
