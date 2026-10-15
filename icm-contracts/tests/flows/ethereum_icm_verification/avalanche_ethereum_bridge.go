package ethereum_icm_verification

import (
	"context"
	"crypto/ecdsa"
	"encoding/hex"
	"math/big"
	"net"
	"time"

	"github.com/ava-labs/avalanchego/ids"
	"github.com/ava-labs/avalanchego/utils/constants"
	merkleregistry "github.com/ava-labs/icm-services/abi-bindings/go/MerkleValidatorSetRegistry"
	teleportermessengerv2 "github.com/ava-labs/icm-services/abi-bindings/go/TeleporterMessengerV2"
	ecdsaverifier "github.com/ava-labs/icm-services/abi-bindings/go/mocks/ECDSAVerifier"
	"github.com/ava-labs/icm-services/config"
	localnetwork "github.com/ava-labs/icm-services/icm-contracts/tests/network"
	testinfo "github.com/ava-labs/icm-services/icm-contracts/tests/test-info"
	"github.com/ava-labs/icm-services/icm-contracts/tests/utils"
	"github.com/ava-labs/icm-services/log"
	"github.com/ava-labs/icm-services/peers/clients"
	relayercfg "github.com/ava-labs/icm-services/relayer/config"
	"github.com/ava-labs/icm-services/relayer/validatorupdater"
	"github.com/ava-labs/libevm/accounts/abi/bind"
	"github.com/ava-labs/libevm/common"
	"github.com/ava-labs/libevm/crypto"
	. "github.com/onsi/gomega"
)

const (
	// How often the relayer's validator-set updater polls, and the longest it waits between updates.
	relayerPollIntervalSeconds      uint64 = 5
	relayerMaxUpdateIntervalSeconds uint64 = 60
)

// avalancheEthereumBridge is a connection between an Avalanche L1 and Ethereum: a TeleporterMessengerV2
// at the same address on both chains, with one relay for each direction.
//   - Avalanche L1 -> Ethereum: delivered by the ICM relayer, verified by the MerkleValidatorSetRegistry
//   - Ethereum -> Avalanche L1: relayed by the test, verified by the mock ECDSAVerifier until the ZK
//     verifier is relayed
type avalancheEthereumBridge struct {
	l1Info              testinfo.L1TestInfo
	ethInfo             *testinfo.EthereumTestInfo
	messengerAddress    common.Address
	l1Messenger         *teleportermessengerv2.TeleporterMessengerV2
	ethMessenger        *teleportermessengerv2.TeleporterMessengerV2
	toEthereum          relay
	toAvalanche         *ecdsaRelay
	signatureAggregator *utils.SignatureAggregator
	stopRelayer         context.CancelFunc
}

// newAvalancheEthereumBridge deploys the registry, adapter and messenger on both chains, registers the
// L1 validator set on Ethereum and starts the relayer. The caller must close the returned bridge.
func newAvalancheEthereumBridge(
	ctx context.Context,
	avalancheNetwork *localnetwork.LocalAvalancheNetwork,
	ethereumNetwork *localnetwork.LocalEthereumNetwork,
	ecdsaSigner *ecdsa.PrivateKey,
	ecdsaVerifierContractAddress common.Address,
) *avalancheEthereumBridge {
	l1Info := avalancheNetwork.GetL1Infos()[0]
	primaryNetworkInfo := avalancheNetwork.GetPrimaryNetworkInfo()
	ethInfo := ethereumNetwork.EthereumTestInfo()
	networkID := avalancheNetwork.GetNetworkID()
	_, ethFundedKey := ethereumNetwork.GetFundedAccountInfo()
	fundedAvalancheAddress, fundedAvalancheKey := avalancheNetwork.GetFundedAccountInfo()
	ethereumOpts, err := bind.NewKeyedTransactorWithChainID(ethFundedKey, ethereumNetwork.ChainID)
	Expect(err).Should(BeNil())

	pChainClient := clients.NewCanonicalValidatorClient(&config.APIConfig{
		BaseURL: primaryNetworkInfo.NodeURIs[0],
	})
	pChainHeight, err := pChainClient.GetLatestHeight(ctx)
	Expect(err).Should(BeNil())
	pChainTimestamp, err := pChainClient.GetBlockTimestampAtHeight(ctx, pChainHeight)
	Expect(err).Should(BeNil())
	pChainWarpSet, err := pChainClient.GetProposedValidators(ctx, ids.Empty)
	Expect(err).Should(BeNil())

	pChainValidators := make([]*validatorupdater.Validator, len(pChainWarpSet.Validators))
	for i, val := range pChainWarpSet.Validators {
		pChainValidators[i] = &validatorupdater.Validator{
			UncompressedPublicKeyBytes: [96]byte(val.PublicKey.Serialize()),
			Weight:                     val.Weight,
		}
	}
	utils.SortValidators(pChainValidators)
	pChainRoot := validatorupdater.BuildMerkleRoot(pChainValidators)
	var pChainTotalWeight uint64
	for _, v := range pChainValidators {
		pChainTotalWeight += v.Weight
	}

	// Registries are deployed with Nick's method, so a registry with the same constructor arguments as
	// one deployed by another test in this suite would collide with it. Allowing the P-Chain fallback
	// gives this connection its own registry, adapter and messenger addresses.
	merkleRegistryAddr := utils.DeployMerkleValidatorSetRegistry(
		ctx, ethInfo, ethFundedKey, networkID, constants.PlatformChainID,
		pChainRoot, pChainTotalWeight, pChainHeight, pChainTimestamp, true,
	)
	merkleRegistryAddrL1 := utils.DeployMerkleValidatorSetRegistry(
		ctx, &l1Info, fundedAvalancheKey, networkID, constants.PlatformChainID,
		pChainRoot, pChainTotalWeight, pChainHeight, pChainTimestamp, true,
	)
	Expect(merkleRegistryAddrL1).Should(Equal(merkleRegistryAddr))
	merkleRegistry, err := merkleregistry.NewMerkleValidatorSetRegistry(merkleRegistryAddr, ethereumNetwork.EthClient)
	Expect(err).Should(BeNil())

	adapterAddr := utils.DeployAdapter(
		ctx, ethInfo, l1Info.BlockchainID, ethInfo.ChainID(),
		ecdsaVerifierContractAddress, merkleRegistryAddr, ethFundedKey,
	)
	adapterAddrL1 := utils.DeployAdapter(
		ctx, &l1Info, l1Info.BlockchainID, ethInfo.ChainID(),
		ecdsaVerifierContractAddress, merkleRegistryAddr, fundedAvalancheKey,
	)
	Expect(adapterAddrL1).Should(Equal(adapterAddr))

	// The same initializer key must be used on both chains for the messengers to share an address.
	// Fund it on Ethereum so it can send the initialize transaction there.
	ethereumNetwork.FundAccount(ctx, fundedAvalancheAddress, new(big.Int).Mul(big.NewInt(1e18), big.NewInt(10)))
	l1MessengerAddr := utils.DeployTeleporterV2(ctx, &l1Info, adapterAddr, fundedAvalancheKey, fundedAvalancheKey)
	ethMessengerAddr := utils.DeployTeleporterV2(ctx, ethInfo, adapterAddr, ethFundedKey, fundedAvalancheKey)
	Expect(l1MessengerAddr).Should(Equal(ethMessengerAddr))

	l1Messenger, err := teleportermessengerv2.NewTeleporterMessengerV2(l1MessengerAddr, l1Info.EthClient)
	Expect(err).Should(BeNil())
	ethMessenger, err := teleportermessengerv2.NewTeleporterMessengerV2(ethMessengerAddr, ethInfo.EthClient)
	Expect(err).Should(BeNil())
	ethVerifier, err := ecdsaverifier.NewECDSAVerifier(ecdsaVerifierContractAddress, ethInfo.EthClient)
	Expect(err).Should(BeNil())

	// Register the L1 validator set before the relayer starts, so it can attest the first message.
	signatureAggregator := avalancheNetwork.GetSignatureAggregator()
	registerL1ValidatorSet(
		ctx, pChainClient, pChainValidators, signatureAggregator,
		networkID, l1Info, merkleRegistry, ethereumOpts, ethereumNetwork.EthClient,
	)
	stopRelayer := startRelayer(
		ctx, avalancheNetwork, ethereumNetwork, l1Info, merkleRegistryAddr, l1MessengerAddr,
	)

	return &avalancheEthereumBridge{
		l1Info:           l1Info,
		ethInfo:          ethInfo,
		messengerAddress: l1MessengerAddr,
		l1Messenger:      l1Messenger,
		ethMessenger:     ethMessenger,
		toEthereum: &relayerRelay{
			source:               &l1Info,
			sourceKey:            fundedAvalancheKey,
			destination:          ethInfo,
			messengerAddress:     l1MessengerAddr,
			sourceMessenger:      l1Messenger,
			destinationMessenger: ethMessenger,
		},
		toAvalanche: &ecdsaRelay{
			attester: &ecdsaAttester{
				ethereumBlockchainID: ethInfo.ChainID(),
				verifierAddress:      ecdsaVerifierContractAddress,
				signer:               ecdsaSigner,
			},
			sourceVerifier:              ethVerifier,
			destination:                 &l1Info,
			destinationKey:              fundedAvalancheKey,
			destinationMessengerAddress: l1MessengerAddr,
			destinationMessenger:        l1Messenger,
		},
		signatureAggregator: signatureAggregator,
		stopRelayer:         stopRelayer,
	}
}

// close stops the relayer and the signature aggregator.
func (b *avalancheEthereumBridge) close() {
	b.stopRelayer()
	b.signatureAggregator.Shutdown()
}

// startRelayer runs the ICM relayer for the bridge's Avalanche L1 -> Ethereum direction and waits until
// it is ready. Messages sent before it starts are not relayed. Returns the function that stops it.
func startRelayer(
	ctx context.Context,
	avalancheNetwork *localnetwork.LocalAvalancheNetwork,
	ethereumNetwork *localnetwork.LocalEthereumNetwork,
	l1Info testinfo.L1TestInfo,
	registryAddress common.Address,
	messengerAddress common.Address,
) context.CancelFunc {
	fundedAvalancheAddress, fundedAvalancheKey := avalancheNetwork.GetFundedAccountInfo()
	ethereumBlockchainID := ethereumNetwork.EthereumTestInfo().ChainID()
	Expect(utils.ClearRelayerStorage()).Should(BeNil())

	relayerKey, err := crypto.GenerateKey()
	Expect(err).Should(BeNil())
	utils.FundRelayers(ctx, []testinfo.L1TestInfo{l1Info}, fundedAvalancheKey, relayerKey)

	// The validator-set updater and the delivery client both send Ethereum transactions, so each gets
	// its own account, apart from each other and from the test's.
	updaterKey := newFundedEthereumKey(ctx, ethereumNetwork)
	deliveryKey := newFundedEthereumKey(ctx, ethereumNetwork)

	relayerConfig := utils.CreateDefaultRelayerConfig(
		log.Root(),
		localnetwork.NewTeleporterTestInfo(avalancheNetwork),
		[]testinfo.L1TestInfo{l1Info},
		[]testinfo.L1TestInfo{l1Info},
		fundedAvalancheAddress,
		relayerKey,
	)
	relayerConfig.APIPort = freePort()
	relayerConfig.MetricsPort = freePort()

	// Relay only the bridge's messages to Ethereum. They reach Warp through the Merkle registry, which
	// makes the registry their message contract.
	source := relayerConfig.SourceBlockchains[0]
	source.MessageContracts = map[string]relayercfg.MessageProtocolConfig{
		registryAddress.Hex(): {
			MessageFormat: relayercfg.TELEPORTER_V2.String(),
			Settings: map[string]interface{}{
				"reward-address":     fundedAvalancheAddress.Hex(),
				"registry-address":   registryAddress.Hex(),
				"teleporter-address": messengerAddress.Hex(),
			},
		},
	}
	source.SupportedDestinations = []*relayercfg.SupportedDestination{
		{BlockchainID: ethereumBlockchainID.String()},
	}
	relayerConfig.ExternalEVMDestinations = []*relayercfg.ExternalEVMDestination{
		{
			RPCEndpoint:              ethereumNetwork.BaseURL,
			PrivateKey:               hex.EncodeToString(crypto.FromECDSA(updaterKey)),
			ContractAddress:          registryAddress.Hex(),
			BlockchainID:             l1Info.BlockchainID.String(),
			SubnetID:                 l1Info.SubnetID.String(),
			PollIntervalSeconds:      relayerPollIntervalSeconds,
			MaxUpdateIntervalSeconds: relayerMaxUpdateIntervalSeconds,
			DestinationBlockchainID:  ethereumBlockchainID.String(),
			TeleporterAddress:        messengerAddress.Hex(),
			DeliveryPrivateKey:       hex.EncodeToString(crypto.FromECDSA(deliveryKey)),
		},
	}
	Expect(relayerConfig.Validate()).Should(BeNil())

	relayerConfigPath := utils.WriteRelayerConfig(log.Root(), relayerConfig, utils.DefaultRelayerCfgFname)
	stopRelayer, ready := utils.RunRelayerExecutable(ctx, log.Root(), relayerConfigPath, relayerConfig)
	startupCtx, cancel := context.WithTimeout(ctx, 60*time.Second)
	defer cancel()
	utils.WaitForChannelClose(startupCtx, ready)
	return stopRelayer
}

func newFundedEthereumKey(ctx context.Context, ethereumNetwork *localnetwork.LocalEthereumNetwork) *ecdsa.PrivateKey {
	key, err := crypto.GenerateKey()
	Expect(err).Should(BeNil())
	ethereumNetwork.FundAccount(
		ctx, crypto.PubkeyToAddress(key.PublicKey), new(big.Int).Mul(big.NewInt(1e18), big.NewInt(10)),
	)
	return key
}

// freePort returns a TCP port that is free on this machine, so that the relayer does not collide with
// other services listening locally.
func freePort() uint16 {
	listener, err := net.Listen("tcp", "127.0.0.1:0")
	Expect(err).Should(BeNil())
	defer listener.Close()
	return uint16(listener.Addr().(*net.TCPAddr).Port)
}
