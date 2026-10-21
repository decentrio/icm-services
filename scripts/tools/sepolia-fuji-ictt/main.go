// Copyright (C) 2026, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

// THIS IS AN EXAMPLE OF UNAUDITED CODE. DO NOT USE THIS IN PRODUCTION.

// Command sepolia-fuji-ictt moves an ERC-20 from Ethereum Sepolia (the token's home) to Avalanche
// Fuji C-Chain (a remote) with ICTT v2. Messages from Sepolia are verified on Fuji by a
// LightClientAdapter against the ZK light client's beacon roots, and relayed by this tool.
//
// The private key of one account funded on both chains is read from the KEY environment variable.
// Addresses of a deployment are kept in a JSON file (-deployment). See README.md for the flow.
package main

import (
	"context"
	"crypto/ecdsa"
	"encoding/json"
	"flag"
	"fmt"
	"math/big"
	"os"
	"strings"

	"github.com/ava-labs/libevm/accounts/abi/bind"
	"github.com/ava-labs/libevm/common"
	"github.com/ava-labs/libevm/core/types"
	"github.com/ava-labs/libevm/crypto"
	"github.com/ava-labs/libevm/ethclient"
)

const (
	sepoliaChainID = 11155111
	// fujiNetworkID is the Avalanche network ID of Fuji.
	fujiNetworkID = 5
)

var (
	// sepoliaBlockchainID is the ICM blockchain ID of Sepolia: its EVM chain ID as bytes32, which
	// EthereumAdapter requires as the source of messages from it.
	sepoliaBlockchainID = common.BigToHash(big.NewInt(sepoliaChainID))
	// fujiCChainID is the blockchain ID of Fuji C-Chain (yH8D7ThNJkxmtkuv2jgBa4P1Rn3Qpr4pPr7QYNfcdoS6k6HWp).
	fujiCChainID = common.HexToHash("0x7fc93d85c6d62c5b2ac0b519c87010ea5294012d1e407030d6acd0021cac10d5")
)

// networks are the endpoints of the two chains.
type networks struct {
	sepoliaRPC, fujiRPC, beaconURL string
}

// deployment holds the addresses of a deployment, written by the deploy commands and read by the
// others.
type deployment struct {
	// LightClient is the ZK light client on Fuji (deployed from avax-light-client).
	LightClient common.Address `json:"lightClient"`
	// The bridge contracts sit at the same address on both chains (CREATE2).
	Registry           common.Address `json:"registry"`
	RegistryHeight     uint64         `json:"registryPChainHeight"`
	LightClientAdapter common.Address `json:"lightClientAdapter"`
	Adapter            common.Address `json:"adapter"`
	Messenger          common.Address `json:"messenger"`
	// The token, its home on Sepolia and its remote on Fuji.
	Token         common.Address `json:"token,omitempty"`
	TokenDecimals uint8          `json:"tokenDecimals,omitempty"`
	TokenHome     common.Address `json:"tokenHome,omitempty"`
	TokenRemote   common.Address `json:"tokenRemote,omitempty"`
}

type deploymentFile struct{ path *string }

func deploymentFlag(fs *flag.FlagSet) deploymentFile {
	return deploymentFile{fs.String("deployment", "deployment.json", "deployment addresses file")}
}

func (d deploymentFile) load() (*deployment, error) {
	b, err := os.ReadFile(*d.path)
	if err != nil {
		return nil, err
	}
	var dep deployment
	if err := json.Unmarshal(b, &dep); err != nil {
		return nil, fmt.Errorf("%s: %w", *d.path, err)
	}
	return &dep, nil
}

func (d deploymentFile) save(dep *deployment) error {
	b, err := json.MarshalIndent(dep, "", "  ")
	if err != nil {
		return err
	}
	fmt.Println("wrote", *d.path)
	return os.WriteFile(*d.path, b, 0o644)
}

// account is the funded account the commands send from, on either chain.
type account struct {
	key     *ecdsa.PrivateKey
	address common.Address
}

func accountFromEnv() (*account, error) {
	key, err := crypto.HexToECDSA(strings.TrimPrefix(os.Getenv("KEY"), "0x"))
	if err != nil {
		return nil, fmt.Errorf("KEY: %w", err)
	}
	return &account{key: key, address: crypto.PubkeyToAddress(key.PublicKey)}, nil
}

// chain is one of the two chains, with the account's transactor on it.
type chain struct {
	name   string
	client *ethclient.Client
	opts   *bind.TransactOpts
}

// dial connects to a chain, sending from acc if it is not nil.
func dial(ctx context.Context, name, url string, acc *account) (*chain, error) {
	client, err := ethclient.DialContext(ctx, url)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", name, err)
	}
	chainID, err := client.ChainID(ctx)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", name, err)
	}
	// Without an account the chain is only read.
	opts := &bind.TransactOpts{Context: ctx}
	if acc != nil {
		if opts, err = bind.NewKeyedTransactorWithChainID(acc.key, chainID); err != nil {
			return nil, err
		}
		opts.Context = ctx
	}
	return &chain{name: name, client: client, opts: opts}, nil
}

// mined waits for tx and fails if it reverted.
func (c *chain) mined(what string, tx *types.Transaction, err error) (*types.Receipt, error) {
	if err != nil {
		return nil, fmt.Errorf("%s on %s: %w", what, c.name, err)
	}
	receipt, err := bind.WaitMined(c.opts.Context, c.client, tx)
	if err != nil {
		return nil, err
	}
	if receipt.Status != types.ReceiptStatusSuccessful {
		return nil, fmt.Errorf("%s on %s reverted in %s", what, c.name, tx.Hash())
	}
	fmt.Printf("  %s on %s: %s (%d gas)\n", what, c.name, tx.Hash(), receipt.GasUsed)
	return receipt, nil
}

func main() {
	n := networks{}
	global := flag.NewFlagSet("sepolia-fuji-ictt", flag.ExitOnError)
	global.StringVar(&n.sepoliaRPC, "sepolia-rpc", "https://ethereum-sepolia-rpc.publicnode.com", "Sepolia execution RPC")
	global.StringVar(&n.fujiRPC, "fuji-rpc", "https://api.avax-test.network/ext/bc/C/rpc", "Fuji C-Chain RPC")
	global.StringVar(&n.beaconURL, "beacon", "https://lodestar-sepolia.chainsafe.io", "Sepolia beacon API")
	global.Usage = func() {
		fmt.Fprintln(os.Stderr, "usage: sepolia-fuji-ictt [network flags] <command> [flags]")
		fmt.Fprintln(os.Stderr, "commands, in order: deploy-bridge, deploy-ictt, register, send, relay, balance")
		global.PrintDefaults()
	}
	_ = global.Parse(os.Args[1:])
	if global.NArg() == 0 {
		global.Usage()
		os.Exit(2)
	}
	cmd, ok := commands[global.Arg(0)]
	if !ok {
		global.Usage()
		os.Exit(2)
	}
	if err := cmd(context.Background(), n, global.Args()[1:]); err != nil {
		fmt.Fprintln(os.Stderr, "error:", err)
		os.Exit(1)
	}
}

var commands = map[string]func(context.Context, networks, []string) error{
	"deploy-bridge": deployBridge,
	"deploy-ictt":   deployICTT,
	"register":      register,
	"send":          send,
	"relay":         relay,
	"balance":       balance,
}
