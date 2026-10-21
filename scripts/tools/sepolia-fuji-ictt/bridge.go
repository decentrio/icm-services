// Copyright (C) 2026, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

// THIS IS AN EXAMPLE OF UNAUDITED CODE. DO NOT USE THIS IN PRODUCTION.

package main

import (
	"context"
	"errors"
	"flag"
	"fmt"
	"math/big"

	"github.com/ava-labs/avalanchego/utils/constants"
	adapterbinding "github.com/ava-labs/icm-services/abi-bindings/go/Adapter"
	merkleregistry "github.com/ava-labs/icm-services/abi-bindings/go/MerkleValidatorSetRegistry"
	teleportermessengerv2 "github.com/ava-labs/icm-services/abi-bindings/go/TeleporterMessengerV2"
	lightclientadapter "github.com/ava-labs/icm-services/abi-bindings/go/verifiers/ethereum/LightClientAdapter"
	"github.com/ava-labs/icm-services/config"
	"github.com/ava-labs/icm-services/messages/zk/proofs"
	"github.com/ava-labs/icm-services/peers/clients"
	"github.com/ava-labs/icm-services/relayer/validatorupdater"
	"github.com/ava-labs/libevm/accounts/abi"
	"github.com/ava-labs/libevm/accounts/abi/bind"
	"github.com/ava-labs/libevm/common"
	"github.com/ava-labs/libevm/crypto"
)

// create2Proxy is Arachnid's deterministic deployment proxy, on both Sepolia and Fuji: it deploys
// salt ‖ initcode with CREATE2, so the same initcode lands at the same address on every chain.
var create2Proxy = common.HexToAddress("0x4e59b44847b379578588920cA78FbF26c0B4956C")

// deployBridge deploys the contracts that carry messages between the chains, at the same address
// on both: the validator set registry, the LightClientAdapter, the Adapter that dispatches between
// them, and the TeleporterMessengerV2, which it then initializes on each chain.
func deployBridge(ctx context.Context, n networks, args []string) error {
	fs := flag.NewFlagSet("deploy-bridge", flag.ExitOnError)
	lightClient := fs.String("light-client", "", "ZK light client on Fuji (from avax-light-client)")
	saltText := fs.String("salt", "decentrio-sepolia-fuji-ictt-v1", "CREATE2 salt")
	pChainAPI := fs.String("p-chain-api", "https://api.avax-test.network", "Fuji P-Chain API")
	pChainHeight := fs.Uint64("p-chain-height", 0, "P-Chain height of the registry's validator set (0: latest)")
	d := deploymentFlag(fs)
	if err := fs.Parse(args); err != nil {
		return err
	}
	if !common.IsHexAddress(*lightClient) {
		return errors.New("-light-client is required")
	}
	acc, err := accountFromEnv()
	if err != nil {
		return err
	}
	sepolia, err := dial(ctx, "Sepolia", n.sepoliaRPC, acc)
	if err != nil {
		return err
	}
	fuji, err := dial(ctx, "Fuji", n.fujiRPC, acc)
	if err != nil {
		return err
	}

	set, err := fujiValidatorSet(ctx, *pChainAPI, *pChainHeight)
	if err != nil {
		return err
	}
	fmt.Printf("Fuji primary network at P-Chain height %d: %d validators, weight %d\n",
		set.height, set.count, set.totalWeight)

	salt := crypto.Keccak256Hash([]byte(*saltText))
	registryCode, err := initcode(merkleregistry.MerkleValidatorSetRegistryMetaData,
		uint32(fujiNetworkID), fujiCChainID, set.root, set.totalWeight, set.height, set.timestamp, false)
	if err != nil {
		return err
	}
	registry := create2Address(salt, registryCode)
	adapterCode, err := initcode(lightclientadapter.LightClientAdapterMetaData,
		big.NewInt(sepoliaChainID), sepoliaBeaconConfig(), common.HexToAddress(*lightClient))
	if err != nil {
		return err
	}
	lcAdapter := create2Address(salt, adapterCode)
	wrapperCode, err := initcode(adapterbinding.AdapterMetaData, fujiCChainID, sepoliaBlockchainID, lcAdapter, registry)
	if err != nil {
		return err
	}
	wrapper := create2Address(salt, wrapperCode)
	messengerCode, err := initcode(teleportermessengerv2.TeleporterMessengerV2MetaData, wrapper, acc.address)
	if err != nil {
		return err
	}
	messenger := create2Address(salt, messengerCode)

	for _, c := range []struct {
		chain        *chain
		blockchainID common.Hash
	}{{sepolia, sepoliaBlockchainID}, {fuji, fujiCChainID}} {
		fmt.Printf("%s:\n", c.chain.name)
		for _, contract := range []struct {
			name string
			code []byte
		}{
			{"MerkleValidatorSetRegistry", registryCode},
			{"LightClientAdapter", adapterCode},
			{"Adapter", wrapperCode},
			{"TeleporterMessengerV2", messengerCode},
		} {
			if _, err := c.chain.create2(salt, contract.name, contract.code); err != nil {
				return err
			}
		}
		if err := initializeMessenger(c.chain, messenger, c.blockchainID); err != nil {
			return err
		}
	}

	r, err := merkleregistry.NewMerkleValidatorSetRegistry(registry, sepolia.client)
	if err != nil {
		return err
	}
	registered, err := r.IsRegistered(&bind.CallOpts{Context: ctx}, fujiCChainID)
	if err != nil {
		return err
	}
	fmt.Printf("registry on Sepolia knows Fuji C-Chain's validators: %t\n", registered)

	return d.save(&deployment{
		LightClient:        common.HexToAddress(*lightClient),
		Registry:           registry,
		RegistryHeight:     set.height,
		LightClientAdapter: lcAdapter,
		Adapter:            wrapper,
		Messenger:          messenger,
	})
}

// sepoliaBeaconConfig is where Sepolia's beacon chain keeps execution data, before and from Gloas.
func sepoliaBeaconConfig() lightclientadapter.ExecutionBeaconConfig {
	return lightclientadapter.ExecutionBeaconConfig{
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
}

type validatorSet struct {
	height, timestamp, totalWeight uint64
	root                           [32]byte
	count                          int
}

// fujiValidatorSet is the canonical Fuji primary network validator set at a P-Chain height, as the
// registry commits to it: validators sorted by uncompressed BLS key, in a Merkle tree.
func fujiValidatorSet(ctx context.Context, api string, height uint64) (*validatorSet, error) {
	client := clients.NewCanonicalValidatorClient(&config.APIConfig{BaseURL: api})
	var err error
	if height == 0 {
		if height, err = client.GetLatestHeight(ctx); err != nil {
			return nil, err
		}
	}
	sets, err := client.GetAllValidatorSets(ctx, height)
	if err != nil {
		return nil, err
	}
	primary, ok := sets[constants.PrimaryNetworkID]
	if !ok {
		return nil, fmt.Errorf("no primary network validators at P-Chain height %d", height)
	}
	validators := make([]*validatorupdater.Validator, len(primary.Validators))
	for i, v := range primary.Validators {
		validators[i] = &validatorupdater.Validator{
			UncompressedPublicKeyBytes: [96]byte(v.PublicKey.Serialize()),
			Weight:                     v.Weight,
		}
	}
	validatorupdater.SortValidators(validators)
	set := &validatorSet{height: height, root: validatorupdater.BuildMerkleRoot(validators), count: len(validators)}
	for _, v := range validators {
		set.totalWeight += v.Weight
	}
	if set.timestamp, err = client.GetBlockTimestampAtHeight(ctx, height); err != nil {
		return nil, err
	}
	return set, nil
}

// initcode is a contract's creation bytecode followed by its constructor arguments.
func initcode(meta *bind.MetaData, args ...any) ([]byte, error) {
	parsed, err := meta.GetAbi()
	if err != nil {
		return nil, err
	}
	packed, err := parsed.Pack("", args...)
	if err != nil {
		return nil, err
	}
	return append(common.FromHex(meta.Bin), packed...), nil
}

func create2Address(salt common.Hash, initcode []byte) common.Address {
	return crypto.CreateAddress2(create2Proxy, salt, crypto.Keccak256(initcode))
}

// create2 deploys initcode through the deployment proxy, unless it is already deployed.
func (c *chain) create2(salt common.Hash, name string, initcode []byte) (common.Address, error) {
	address := create2Address(salt, initcode)
	code, err := c.client.CodeAt(c.opts.Context, address, nil)
	if err != nil {
		return common.Address{}, err
	}
	if len(code) > 0 {
		fmt.Printf("  %s already at %s\n", name, address)
		return address, nil
	}
	proxy := bind.NewBoundContract(create2Proxy, abi.ABI{}, c.client, c.client, c.client)
	tx, err := proxy.RawTransact(c.opts, append(salt.Bytes(), initcode...))
	if _, err := c.mined("deploy "+name, tx, err); err != nil {
		return common.Address{}, err
	}
	if code, err = c.client.CodeAt(c.opts.Context, address, nil); err != nil || len(code) == 0 {
		return common.Address{}, fmt.Errorf("%s not deployed at %s on %s: %v", name, address, c.name, err)
	}
	fmt.Printf("  %s at %s\n", name, address)
	return address, nil
}

// initializeMessenger sets the messenger's blockchain ID, unless it is already set.
func initializeMessenger(c *chain, address common.Address, blockchainID common.Hash) error {
	messenger, err := teleportermessengerv2.NewTeleporterMessengerV2(address, c.client)
	if err != nil {
		return err
	}
	current, err := messenger.BlockchainID(&bind.CallOpts{Context: c.opts.Context})
	if err != nil {
		return err
	}
	if current == blockchainID {
		return nil
	}
	if current != (common.Hash{}) {
		return fmt.Errorf("messenger on %s is initialized for %s", c.name, common.Hash(current))
	}
	tx, err := messenger.Initialize(c.opts, blockchainID)
	_, err = c.mined("initialize TeleporterMessengerV2", tx, err)
	return err
}
