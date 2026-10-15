package ethereum_icm_verification

import (
	"context"
	"crypto/ecdsa"
	"math/big"

	"github.com/ava-labs/avalanchego/ids"
	"github.com/ava-labs/avalanchego/utils/constants"
	merkleregistry "github.com/ava-labs/icm-services/abi-bindings/go/MerkleValidatorSetRegistry"
	teleportermessengerv2 "github.com/ava-labs/icm-services/abi-bindings/go/TeleporterMessengerV2"
	erc20tokenhomev2 "github.com/ava-labs/icm-services/abi-bindings/go/ictt-v2/TokenHome/ERC20TokenHome"
	erc20tokenremotev2 "github.com/ava-labs/icm-services/abi-bindings/go/ictt-v2/TokenRemote/ERC20TokenRemote"
	exampleerc20 "github.com/ava-labs/icm-services/abi-bindings/go/ictt/mocks/ExampleERC20Decimals"
	mockERC20SACR "github.com/ava-labs/icm-services/abi-bindings/go/ictt/mocks/MockERC20SendAndCallReceiver"
	ecdsaverifier "github.com/ava-labs/icm-services/abi-bindings/go/mocks/ECDSAVerifier"
	"github.com/ava-labs/icm-services/config"
	localnetwork "github.com/ava-labs/icm-services/icm-contracts/tests/network"
	testinfo "github.com/ava-labs/icm-services/icm-contracts/tests/test-info"
	"github.com/ava-labs/icm-services/icm-contracts/tests/utils"
	"github.com/ava-labs/icm-services/peers/clients"
	"github.com/ava-labs/icm-services/relayer/validatorupdater"
	"github.com/ava-labs/libevm/accounts/abi/bind"
	"github.com/ava-labs/libevm/common"
	"github.com/ava-labs/libevm/core/types"
	. "github.com/onsi/gomega"
)

/**
* Test ICTT v2 transferring an ERC20 token from an Avalanche L1 to Ethereum over TeleporterMessengerV2.
* - Avalanche L1 -> Ethereum: signed by the L1 validator set and verified by the MerkleValidatorSetRegistry on Ethereum
* - Ethereum -> Avalanche L1: only used to register the remote. Signed with ECDSA and verified by the mock
*   ECDSAVerifier on the L1 until the ZK path is relayed.
*
* The test proceeds in the following steps.
* 1. Deploy a two-way TeleporterMessengerV2 connection between the L1 and Ethereum, and register the
*    L1 validator set on Ethereum.
* 2. Deploy an ERC20TokenHome on the L1 and route Ethereum through the connection's messenger.
* 3. Deploy an ERC20TokenRemote on Ethereum bound to the same messenger.
* 4. Register the remote with the home (Ethereum -> L1).
* 5. Send tokens from the L1 to Ethereum and check that the remote minted them against tokens locked in the home.
* 6. Send tokens and call a contract on Ethereum.
 */
func ICTTV2AvalancheToEthereum(
	ctx context.Context,
	avalancheNetwork *localnetwork.LocalAvalancheNetwork,
	ethereumNetwork *localnetwork.LocalEthereumNetwork,
	ecdsaSigner *ecdsa.PrivateKey,
	ecdsaVerifierContractAddress common.Address,
) {
	const tokenDecimals = uint8(18)

	fundedAvalancheAddress, fundedAvalancheKey := avalancheNetwork.GetFundedAccountInfo()
	fundedEthereumAddress, _ := ethereumNetwork.GetFundedAccountInfo()

	// Step 1: Deploy the two-way connection between the L1 and Ethereum
	bridge := newAvalancheEthereumBridge(
		ctx, avalancheNetwork, ethereumNetwork, ecdsaSigner, ecdsaVerifierContractAddress,
	)
	defer bridge.signatureAggregator.Shutdown()
	l1Info := bridge.l1Info
	ethInfo := bridge.ethInfo
	ethereumBlockchainID := ethInfo.ChainID()

	// Step 2: Deploy the token and its home on the L1, sending to Ethereum through the connection's messenger
	tokenAddress, token := utils.DeployExampleERC20Decimals(ctx, fundedAvalancheKey, l1Info, tokenDecimals)
	homeAddress, tx, home, err := erc20tokenhomev2.DeployERC20TokenHome(
		bridge.l1Opts,
		l1Info.EthClient,
		l1Info.BlockchainID,
		fundedAvalancheAddress,
		tokenAddress,
		tokenDecimals,
	)
	Expect(err).Should(BeNil())
	utils.WaitForTransactionSuccess(ctx, l1Info.EthClient, tx.Hash())

	tx, err = home.AddMessenger(bridge.l1Opts, ethereumBlockchainID, bridge.messengerAddress)
	Expect(err).Should(BeNil())
	utils.WaitForTransactionSuccess(ctx, l1Info.EthClient, tx.Hash())
	tx, err = home.SetOutboundMessenger(bridge.l1Opts, ethereumBlockchainID, bridge.messengerAddress)
	Expect(err).Should(BeNil())
	utils.WaitForTransactionSuccess(ctx, l1Info.EthClient, tx.Hash())

	// Step 3: Deploy the remote on Ethereum. Initialization binds the messenger to the L1.
	remoteAddress, tx, remote, err := erc20tokenremotev2.DeployERC20TokenRemote(
		bridge.ethOpts,
		ethInfo.EthClient,
		erc20tokenremotev2.TokenRemoteSettingsV2{
			BlockchainID:          ethereumBlockchainID,
			TeleporterMessenger:   bridge.messengerAddress,
			TeleporterManager:     fundedEthereumAddress,
			TokenHomeBlockchainID: l1Info.BlockchainID,
			TokenHomeAddress:      homeAddress,
			TokenHomeDecimals:     tokenDecimals,
		},
		"Wrapped Example Token",
		"WEXT",
		tokenDecimals,
	)
	Expect(err).Should(BeNil())
	utils.WaitForTransactionSuccess(ctx, ethInfo.EthClient, tx.Hash())

	// Step 4: Register the remote with the home
	tx, err = remote.RegisterWithHome(
		bridge.ethOpts,
		erc20tokenremotev2.TeleporterFeeInfo{FeeTokenAddress: common.Address{}, Amount: big.NewInt(0)},
	)
	Expect(err).Should(BeNil())
	receipt := utils.WaitForTransactionSuccess(ctx, ethInfo.EthClient, tx.Hash())
	receipt = bridge.relayToAvalanche(ctx, receipt)

	registerEvent, err := utils.GetEventFromLogs(receipt.Logs, home.ParseRemoteRegistered)
	Expect(err).Should(BeNil())
	Expect(registerEvent.RemoteBlockchainID[:]).Should(Equal(ethereumBlockchainID[:]))
	Expect(registerEvent.RemoteTokenTransferrerAddress).Should(Equal(remoteAddress))

	// Step 5: Send tokens from the L1 to Ethereum
	recipientAddress := common.HexToAddress("0xabcDEFabCdeFabcdEFAbCdefABCDeFAbCDEfaBcD")
	amount := new(big.Int).Mul(big.NewInt(1e18), big.NewInt(13))
	utils.ERC20DecimalsApprove(ctx, token, homeAddress, amount, l1Info, fundedAvalancheKey)
	tx, err = home.Send(
		bridge.l1Opts,
		erc20tokenhomev2.SendTokensInput{
			DestinationBlockchainID:            ethereumBlockchainID,
			DestinationTokenTransferrerAddress: remoteAddress,
			Recipient:                          recipientAddress,
			PrimaryFeeTokenAddress:             tokenAddress,
			PrimaryFee:                         big.NewInt(0),
			SecondaryFee:                       big.NewInt(0),
			RequiredGasLimit:                   utils.DefaultERC20RequiredGas,
			MultiHopFallback:                   common.Address{},
		},
		amount,
	)
	Expect(err).Should(BeNil())
	receipt = utils.WaitForTransactionSuccess(ctx, l1Info.EthClient, tx.Hash())
	sendEvent, err := utils.GetEventFromLogs(receipt.Logs, home.ParseTokensSent)
	Expect(err).Should(BeNil())
	utils.ExpectBigEqual(sendEvent.Amount, amount)

	receipt = bridge.relayToEthereum(ctx, receipt)
	withdrawEvent, err := utils.GetEventFromLogs(receipt.Logs, remote.ParseTokensWithdrawn)
	Expect(err).Should(BeNil())
	Expect(withdrawEvent.Recipient).Should(Equal(recipientAddress))
	utils.ExpectBigEqual(withdrawEvent.Amount, amount)

	recipientBalance, err := remote.BalanceOf(&bind.CallOpts{}, recipientAddress)
	Expect(err).Should(BeNil())
	utils.ExpectBigEqual(recipientBalance, amount)

	// Every token minted on Ethereum is backed by a token locked in the home
	checkBacking(home, homeAddress, token, remote, remoteAddress, ethereumBlockchainID, amount)

	// Step 6: Send tokens and call a contract on Ethereum
	receiverAddress, tx, receiver, err := mockERC20SACR.DeployMockERC20SendAndCallReceiver(
		bridge.ethOpts, ethInfo.EthClient,
	)
	Expect(err).Should(BeNil())
	utils.WaitForTransactionSuccess(ctx, ethInfo.EthClient, tx.Hash())

	callAmount := big.NewInt(1e18)
	payload := []byte("Hello Ethereum!")
	utils.ERC20DecimalsApprove(ctx, token, homeAddress, callAmount, l1Info, fundedAvalancheKey)
	tx, err = home.SendAndCall(
		bridge.l1Opts,
		erc20tokenhomev2.SendAndCallInput{
			DestinationBlockchainID:            ethereumBlockchainID,
			DestinationTokenTransferrerAddress: remoteAddress,
			RecipientContract:                  receiverAddress,
			RecipientPayload:                   payload,
			RequiredGasLimit:                   utils.BigIntMul(big.NewInt(10), utils.DefaultERC20RequiredGas),
			RecipientGasLimit:                  utils.BigIntMul(big.NewInt(5), utils.DefaultERC20RequiredGas),
			MultiHopFallback:                   common.Address{},
			FallbackRecipient:                  fundedEthereumAddress,
			PrimaryFeeTokenAddress:             tokenAddress,
			PrimaryFee:                         big.NewInt(0),
			SecondaryFee:                       big.NewInt(0),
		},
		callAmount,
	)
	Expect(err).Should(BeNil())
	receipt = utils.WaitForTransactionSuccess(ctx, l1Info.EthClient, tx.Hash())

	receipt = bridge.relayToEthereum(ctx, receipt)
	receivedEvent, err := utils.GetEventFromLogs(receipt.Logs, receiver.ParseTokensReceived)
	Expect(err).Should(BeNil())
	Expect(receivedEvent.SourceBlockchainID[:]).Should(Equal(l1Info.BlockchainID[:]))
	Expect(receivedEvent.OriginTokenTransferrerAddress).Should(Equal(homeAddress))
	Expect(receivedEvent.OriginSenderAddress).Should(Equal(fundedAvalancheAddress))
	Expect(receivedEvent.Token).Should(Equal(remoteAddress))
	utils.ExpectBigEqual(receivedEvent.Amount, callAmount)
	Expect(receivedEvent.Payload).Should(Equal(payload))

	receiverBalance, err := remote.BalanceOf(&bind.CallOpts{}, receiverAddress)
	Expect(err).Should(BeNil())
	utils.ExpectBigEqual(receiverBalance, callAmount)

	checkBacking(
		home, homeAddress, token, remote, remoteAddress, ethereumBlockchainID,
		new(big.Int).Add(amount, callAmount),
	)
}

// checkBacking checks that the remote's total supply on Ethereum equals both the tokens the home
// recorded as sent to it and the tokens the home holds. The token uses the same decimals on both
// chains, so no scaling applies.
func checkBacking(
	home *erc20tokenhomev2.ERC20TokenHome,
	homeAddress common.Address,
	token *exampleerc20.ExampleERC20Decimals,
	remote *erc20tokenremotev2.ERC20TokenRemote,
	remoteAddress common.Address,
	ethereumBlockchainID ids.ID,
	expected *big.Int,
) {
	totalSupply, err := remote.TotalSupply(&bind.CallOpts{})
	Expect(err).Should(BeNil())
	utils.ExpectBigEqual(totalSupply, expected)

	transferredBalance, err := home.GetTransferredBalance(&bind.CallOpts{}, ethereumBlockchainID, remoteAddress)
	Expect(err).Should(BeNil())
	utils.ExpectBigEqual(transferredBalance, expected)

	lockedBalance, err := token.BalanceOf(&bind.CallOpts{}, homeAddress)
	Expect(err).Should(BeNil())
	utils.ExpectBigEqual(lockedBalance, expected)
}

// avalancheEthereumBridge is a two-way TeleporterMessengerV2 connection between an Avalanche L1 and
// Ethereum, with the messenger at the same address on both chains.
// - Avalanche L1 -> Ethereum: Warp, verified by the MerkleValidatorSetRegistry on Ethereum
// - Ethereum -> Avalanche L1: verified by the mock ECDSAVerifier on the L1
type avalancheEthereumBridge struct {
	l1Info               testinfo.L1TestInfo
	ethInfo              *testinfo.EthereumTestInfo
	networkID            uint32
	messengerAddress     common.Address
	l1Messenger          *teleportermessengerv2.TeleporterMessengerV2
	ethMessenger         *teleportermessengerv2.TeleporterMessengerV2
	l1Opts               *bind.TransactOpts
	ethOpts              *bind.TransactOpts
	signatureAggregator  *utils.SignatureAggregator
	pChainClient         *clients.CanonicalValidatorClient
	registeredHeight     uint64
	ecdsaSigner          *ecdsa.PrivateKey
	ecdsaVerifierAddress common.Address
}

// newAvalancheEthereumBridge deploys the registry, adapter and messenger on both chains and registers
// the L1 validator set on Ethereum. The caller must shut down the returned signature aggregator.
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
	avalancheOpts, err := bind.NewKeyedTransactorWithChainID(fundedAvalancheKey, l1Info.EVMChainID)
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

	signatureAggregator := avalancheNetwork.GetSignatureAggregator()
	registeredHeight := registerL1ValidatorSet(
		ctx, pChainClient, pChainValidators, signatureAggregator,
		networkID, l1Info, merkleRegistry, ethereumOpts, ethereumNetwork.EthClient,
	)

	return &avalancheEthereumBridge{
		l1Info:               l1Info,
		ethInfo:              ethInfo,
		networkID:            networkID,
		messengerAddress:     l1MessengerAddr,
		l1Messenger:          l1Messenger,
		ethMessenger:         ethMessenger,
		l1Opts:               avalancheOpts,
		ethOpts:              ethereumOpts,
		signatureAggregator:  signatureAggregator,
		pChainClient:         pChainClient,
		registeredHeight:     registeredHeight,
		ecdsaSigner:          ecdsaSigner,
		ecdsaVerifierAddress: ecdsaVerifierContractAddress,
	}
}

// relayToEthereum delivers the message sent by the L1 messenger in [sourceReceipt] to Ethereum. The message
// is attested by the L1 validators' aggregate BLS signature and a Merkle proof of the signers against the
// commitment registered on Ethereum. Returns the delivery receipt after checking that the message executed.
func (b *avalancheEthereumBridge) relayToEthereum(ctx context.Context, sourceReceipt *types.Receipt) *types.Receipt {
	sendEvent, err := utils.GetEventFromLogs(sourceReceipt.Logs, b.l1Messenger.ParseSendCrossChainMessage)
	Expect(err).Should(BeNil())

	attestation := signMessageAvalancheMerkle(
		ctx, sourceReceipt, b.l1Info, b.signatureAggregator, b.pChainClient, b.registeredHeight,
	)
	tx, err := b.ethMessenger.ReceiveCrossChainMessage(
		b.ethOpts,
		teleportermessengerv2.TeleporterICMMessage{
			Message:            sendEvent.Message,
			SourceNetworkID:    b.networkID,
			SourceBlockchainID: b.l1Info.BlockchainID,
			Attestation:        attestation,
		},
		common.Address{},
	)
	Expect(err).Should(BeNil())
	receipt := utils.WaitForTransactionSuccess(ctx, b.ethInfo.EthClient, tx.Hash())

	_, err = utils.GetEventFromLogs(receipt.Logs, b.ethMessenger.ParseMessageExecuted)
	Expect(err).Should(BeNil(), "message was delivered to Ethereum but did not execute")
	return receipt
}

// relayToAvalanche delivers the message sent through the ECDSAVerifier on Ethereum in [sourceReceipt] to
// the L1, signing it with the trusted ECDSA key. Returns the delivery receipt after checking that the
// message executed.
func (b *avalancheEthereumBridge) relayToAvalanche(ctx context.Context, sourceReceipt *types.Receipt) *types.Receipt {
	ecdsaVerifier, err := ecdsaverifier.NewECDSAVerifier(b.ecdsaVerifierAddress, b.ethInfo.EthClient)
	Expect(err).Should(BeNil())
	sendEvent, err := utils.GetEventFromLogs(sourceReceipt.Logs, ecdsaVerifier.ParseECDSAVerifierSendMessage)
	Expect(err).Should(BeNil())

	sent := sendEvent.Message
	receipts := make([]teleportermessengerv2.TeleporterMessageReceipt, len(sent.Receipts))
	for i, r := range sent.Receipts {
		receipts[i] = teleportermessengerv2.TeleporterMessageReceipt{
			ReceivedMessageNonce: r.ReceivedMessageNonce,
			RelayerRewardAddress: r.RelayerRewardAddress,
		}
	}
	tx, err := b.l1Messenger.ReceiveCrossChainMessage(
		b.l1Opts,
		teleportermessengerv2.TeleporterICMMessage{
			Message: teleportermessengerv2.TeleporterMessageV2{
				MessageNonce:            sent.MessageNonce,
				OriginSenderAddress:     sent.OriginSenderAddress,
				OriginTeleporterAddress: sent.OriginTeleporterAddress,
				DestinationBlockchainID: sent.DestinationBlockchainID,
				DestinationAddress:      sent.DestinationAddress,
				RequiredGasLimit:        sent.RequiredGasLimit,
				AllowedRelayerAddresses: sent.AllowedRelayerAddresses,
				Receipts:                receipts,
				Message:                 sent.Message,
			},
			SourceNetworkID:    0,
			SourceBlockchainID: b.ethInfo.ChainID(),
			Attestation:        signMessageEcdsa(sent, b.ethInfo.ChainID(), b.ecdsaVerifierAddress, b.ecdsaSigner),
		},
		common.Address{},
	)
	Expect(err).Should(BeNil())
	receipt := utils.WaitForTransactionSuccess(ctx, b.l1Info.EthClient, tx.Hash())

	_, err = utils.GetEventFromLogs(receipt.Logs, b.l1Messenger.ParseMessageExecuted)
	Expect(err).Should(BeNil(), "message was delivered to the L1 but did not execute")
	return receipt
}
