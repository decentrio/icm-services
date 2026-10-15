package ethereum_icm_verification

import (
	"context"
	"crypto/ecdsa"
	"math/big"

	"github.com/ava-labs/avalanchego/ids"
	teleportermessengerv2 "github.com/ava-labs/icm-services/abi-bindings/go/TeleporterMessengerV2"
	erc20tokenhomev2 "github.com/ava-labs/icm-services/abi-bindings/go/ictt-v2/TokenHome/ERC20TokenHome"
	erc20tokenremotev2 "github.com/ava-labs/icm-services/abi-bindings/go/ictt-v2/TokenRemote/ERC20TokenRemote"
	exampleerc20 "github.com/ava-labs/icm-services/abi-bindings/go/ictt/mocks/ExampleERC20Decimals"
	mockERC20SACR "github.com/ava-labs/icm-services/abi-bindings/go/ictt/mocks/MockERC20SendAndCallReceiver"
	localnetwork "github.com/ava-labs/icm-services/icm-contracts/tests/network"
	testinfo "github.com/ava-labs/icm-services/icm-contracts/tests/test-info"
	"github.com/ava-labs/icm-services/icm-contracts/tests/utils"
	"github.com/ava-labs/libevm/accounts/abi/bind"
	"github.com/ava-labs/libevm/common"
	. "github.com/onsi/gomega"
)

/**
* Test ICTT v2 moving an ERC20 token between an Avalanche L1 and Ethereum over TeleporterMessengerV2.
* - Avalanche L1 -> Ethereum: signed by the L1 validator set, delivered by the ICM relayer and verified by
*   the MerkleValidatorSetRegistry on Ethereum
* - Ethereum -> Avalanche L1: signed with ECDSA and verified by the mock ECDSAVerifier on the L1 until the
*   ZK verifier is relayed
* - C-Chain <-> Avalanche L1: Teleporter (V1) over Warp
*
* The test proceeds in the following steps.
* 1. Deploy the connection between the L1 and Ethereum, register the L1 validator set on Ethereum and
*    start the relayer.
* 2. Deploy an ERC20TokenHome on the L1 and route Ethereum through the connection's messenger.
* 3. Deploy an ERC20TokenRemote on Ethereum bound to the same messenger.
* 4. Register the remote with the home (Ethereum -> L1).
* 5. Send tokens from the L1 to Ethereum and check that the remote minted them against tokens locked in the home.
* 6. Send tokens and call a contract on Ethereum.
* 7. Multi-hop: send tokens from a remote on the C-Chain to Ethereum through the home, which receives
*    them over Warp and forwards them over the connection.
* 8. Send tokens back from Ethereum. A second messenger on the L1, whose verifier accepts the message,
*    cannot deliver it to the home; the connection's messenger can, and the home releases the tokens.
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
	fundedEthereumAddress, fundedEthereumKey := ethereumNetwork.GetFundedAccountInfo()

	// Step 1: Deploy the connection between the L1 and Ethereum
	bridge := newAvalancheEthereumBridge(
		ctx, avalancheNetwork, ethereumNetwork, ecdsaSigner, ecdsaVerifierContractAddress,
	)
	defer bridge.close()
	l1Info := bridge.l1Info
	ethInfo := bridge.ethInfo
	ethereumBlockchainID := ethInfo.ChainID()
	l1Opts, err := bind.NewKeyedTransactorWithChainID(fundedAvalancheKey, l1Info.EVMChainID)
	Expect(err).Should(BeNil())
	ethOpts, err := bind.NewKeyedTransactorWithChainID(fundedEthereumKey, ethInfo.EVMChainID)
	Expect(err).Should(BeNil())

	// Step 2: Deploy the token and its home on the L1, sending to Ethereum through the connection's messenger
	tokenAddress, token := utils.DeployExampleERC20Decimals(ctx, fundedAvalancheKey, l1Info, tokenDecimals)
	homeAddress, home := utils.DeployERC20TokenHomeV2(
		ctx, fundedAvalancheKey, &l1Info, fundedAvalancheAddress, tokenAddress, tokenDecimals,
	)
	utils.AddRouteV2(ctx, fundedAvalancheKey, &l1Info, home, ethereumBlockchainID, bridge.messengerAddress)

	// Step 3: Deploy the remote on Ethereum. Initialization routes the L1 through the same messenger.
	remoteAddress, remote := utils.DeployERC20TokenRemoteV2(
		ctx, fundedEthereumKey, ethInfo, bridge.messengerAddress, fundedEthereumAddress,
		l1Info.BlockchainID, homeAddress, tokenDecimals, "Wrapped Example Token", "WEXT", tokenDecimals,
	)
	ethRemote := backedRemote{blockchainID: ethereumBlockchainID, address: remoteAddress, remote: remote}

	// Step 4: Register the remote with the home
	tx, err := remote.RegisterWithHome(
		ethOpts,
		erc20tokenremotev2.TeleporterFeeInfo{FeeTokenAddress: common.Address{}, Amount: big.NewInt(0)},
	)
	Expect(err).Should(BeNil())
	receipt := utils.WaitForTransactionSuccess(ctx, ethInfo.EthClient, tx.Hash())
	receipt = bridge.toAvalanche.deliver(ctx, receipt)

	registerEvent, err := utils.GetEventFromLogsEmittedBy(receipt.Logs, homeAddress, home.ParseRemoteRegistered)
	Expect(err).Should(BeNil())
	Expect(registerEvent.RemoteBlockchainID[:]).Should(Equal(ethereumBlockchainID[:]))
	Expect(registerEvent.RemoteTokenTransferrerAddress).Should(Equal(remoteAddress))

	// Step 5: Send tokens from the L1 to Ethereum
	recipientAddress := common.HexToAddress("0xabcDEFabCdeFabcdEFAbCdefABCDeFAbCDEfaBcD")
	amount := new(big.Int).Mul(big.NewInt(1e18), big.NewInt(13))
	utils.ERC20DecimalsApprove(ctx, token, homeAddress, amount, l1Info, fundedAvalancheKey)
	tx, err = home.Send(
		l1Opts,
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
	sendEvent, err := utils.GetEventFromLogsEmittedBy(receipt.Logs, homeAddress, home.ParseTokensSent)
	Expect(err).Should(BeNil())
	utils.ExpectBigEqual(sendEvent.Amount, amount)

	receipt = bridge.toEthereum.deliver(ctx, receipt)
	withdrawEvent, err := utils.GetEventFromLogsEmittedBy(receipt.Logs, remoteAddress, remote.ParseTokensWithdrawn)
	Expect(err).Should(BeNil())
	Expect(withdrawEvent.Recipient).Should(Equal(recipientAddress))
	utils.ExpectBigEqual(withdrawEvent.Amount, amount)
	expectERC20Balance(remote, recipientAddress, amount)
	checkBacking(home, homeAddress, token, ethRemote)

	// Step 6: Send tokens and call a contract on Ethereum
	receiverAddress, tx, receiver, err := mockERC20SACR.DeployMockERC20SendAndCallReceiver(ethOpts, ethInfo.EthClient)
	Expect(err).Should(BeNil())
	utils.WaitForTransactionSuccess(ctx, ethInfo.EthClient, tx.Hash())

	callAmount := big.NewInt(1e18)
	payload := []byte("Hello Ethereum!")
	utils.ERC20DecimalsApprove(ctx, token, homeAddress, callAmount, l1Info, fundedAvalancheKey)
	tx, err = home.SendAndCall(
		l1Opts,
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

	receipt = bridge.toEthereum.deliver(ctx, receipt)
	receivedEvent, err := utils.GetEventFromLogsEmittedBy(receipt.Logs, receiverAddress, receiver.ParseTokensReceived)
	Expect(err).Should(BeNil())
	Expect(receivedEvent.SourceBlockchainID[:]).Should(Equal(l1Info.BlockchainID[:]))
	Expect(receivedEvent.OriginTokenTransferrerAddress).Should(Equal(homeAddress))
	Expect(receivedEvent.OriginSenderAddress).Should(Equal(fundedAvalancheAddress))
	Expect(receivedEvent.Token).Should(Equal(remoteAddress))
	utils.ExpectBigEqual(receivedEvent.Amount, callAmount)
	Expect(receivedEvent.Payload).Should(Equal(payload))
	expectERC20Balance(remote, receiverAddress, callAmount)
	checkBacking(home, homeAddress, token, ethRemote)

	// Step 7: Multi-hop from a remote on the C-Chain to Ethereum through the home
	cChainInfo := avalancheNetwork.GetPrimaryNetworkInfo()
	cChainOpts, err := bind.NewKeyedTransactorWithChainID(fundedAvalancheKey, cChainInfo.EVMChainID)
	Expect(err).Should(BeNil())
	warpMessengerAddress := deployWarpMessenger(ctx, fundedAvalancheKey, cChainInfo, l1Info)
	teleporterInfo := localnetwork.NewTeleporterTestInfo(avalancheNetwork)
	teleporterInfo.SetTeleporter(warpMessengerAddress, cChainInfo.BlockchainID)
	teleporterInfo.SetTeleporter(warpMessengerAddress, l1Info.BlockchainID)
	cChainToL1 := &warpRelay{
		teleporter:  teleporterInfo,
		source:      cChainInfo,
		destination: l1Info,
		relayerKey:  fundedAvalancheKey,
		aggregator:  bridge.signatureAggregator,
	}
	l1ToCChain := &warpRelay{
		teleporter:  teleporterInfo,
		source:      l1Info,
		destination: cChainInfo,
		relayerKey:  fundedAvalancheKey,
		aggregator:  bridge.signatureAggregator,
	}

	cChainRemoteAddress, cChainRemote := utils.DeployERC20TokenRemoteV2(
		ctx, fundedAvalancheKey, &cChainInfo, warpMessengerAddress, fundedAvalancheAddress,
		l1Info.BlockchainID, homeAddress, tokenDecimals, "C-Chain Example Token", "CEXT", tokenDecimals,
	)
	cChainBackedRemote := backedRemote{
		blockchainID: cChainInfo.BlockchainID, address: cChainRemoteAddress, remote: cChainRemote,
	}
	utils.AddRouteV2(ctx, fundedAvalancheKey, &l1Info, home, cChainInfo.BlockchainID, warpMessengerAddress)

	tx, err = cChainRemote.RegisterWithHome(
		cChainOpts,
		erc20tokenremotev2.TeleporterFeeInfo{FeeTokenAddress: common.Address{}, Amount: big.NewInt(0)},
	)
	Expect(err).Should(BeNil())
	receipt = utils.WaitForTransactionSuccess(ctx, cChainInfo.EthClient, tx.Hash())
	receipt = cChainToL1.deliver(ctx, receipt)
	_, err = utils.GetEventFromLogsEmittedBy(receipt.Logs, homeAddress, home.ParseRemoteRegistered)
	Expect(err).Should(BeNil())

	cChainAmount := new(big.Int).Mul(big.NewInt(1e18), big.NewInt(5))
	utils.ERC20DecimalsApprove(ctx, token, homeAddress, cChainAmount, l1Info, fundedAvalancheKey)
	tx, err = home.Send(
		l1Opts,
		erc20tokenhomev2.SendTokensInput{
			DestinationBlockchainID:            cChainInfo.BlockchainID,
			DestinationTokenTransferrerAddress: cChainRemoteAddress,
			Recipient:                          fundedAvalancheAddress,
			PrimaryFeeTokenAddress:             tokenAddress,
			PrimaryFee:                         big.NewInt(0),
			SecondaryFee:                       big.NewInt(0),
			RequiredGasLimit:                   utils.DefaultERC20RequiredGas,
			MultiHopFallback:                   common.Address{},
		},
		cChainAmount,
	)
	Expect(err).Should(BeNil())
	receipt = utils.WaitForTransactionSuccess(ctx, l1Info.EthClient, tx.Hash())
	l1ToCChain.deliver(ctx, receipt)
	expectERC20Balance(cChainRemote, fundedAvalancheAddress, cChainAmount)

	hopAmount := new(big.Int).Mul(big.NewInt(1e18), big.NewInt(2))
	tx, err = cChainRemote.Approve(cChainOpts, cChainRemoteAddress, hopAmount)
	Expect(err).Should(BeNil())
	utils.WaitForTransactionSuccess(ctx, cChainInfo.EthClient, tx.Hash())
	tx, err = cChainRemote.Send(
		cChainOpts,
		erc20tokenremotev2.SendTokensInput{
			DestinationBlockchainID:            ethereumBlockchainID,
			DestinationTokenTransferrerAddress: remoteAddress,
			Recipient:                          fundedEthereumAddress,
			PrimaryFeeTokenAddress:             cChainRemoteAddress,
			PrimaryFee:                         big.NewInt(0),
			SecondaryFee:                       big.NewInt(0),
			RequiredGasLimit:                   utils.DefaultERC20RequiredGas,
			MultiHopFallback:                   fundedAvalancheAddress,
		},
		hopAmount,
	)
	Expect(err).Should(BeNil())
	receipt = utils.WaitForTransactionSuccess(ctx, cChainInfo.EthClient, tx.Hash())

	// The home receives the first hop over Warp and forwards it over the connection's messenger
	receipt = cChainToL1.deliver(ctx, receipt)
	routedEvent, err := utils.GetEventFromLogsEmittedBy(receipt.Logs, homeAddress, home.ParseTokensRouted)
	Expect(err).Should(BeNil())
	utils.ExpectBigEqual(routedEvent.Amount, hopAmount)

	receipt = bridge.toEthereum.deliver(ctx, receipt)
	withdrawEvent, err = utils.GetEventFromLogsEmittedBy(receipt.Logs, remoteAddress, remote.ParseTokensWithdrawn)
	Expect(err).Should(BeNil())
	Expect(withdrawEvent.Recipient).Should(Equal(fundedEthereumAddress))
	utils.ExpectBigEqual(withdrawEvent.Amount, hopAmount)
	expectERC20Balance(remote, fundedEthereumAddress, hopAmount)
	expectERC20Balance(cChainRemote, fundedAvalancheAddress, new(big.Int).Sub(cChainAmount, hopAmount))
	checkBacking(home, homeAddress, token, ethRemote, cChainBackedRemote)

	// Step 8: Send tokens back from Ethereum
	backRecipientAddress := common.HexToAddress("0x5EAd0000000000000000000000000000000BaCcC")
	backAmount := big.NewInt(1e18)
	tx, err = remote.Approve(ethOpts, remoteAddress, backAmount)
	Expect(err).Should(BeNil())
	utils.WaitForTransactionSuccess(ctx, ethInfo.EthClient, tx.Hash())
	tx, err = remote.Send(
		ethOpts,
		erc20tokenremotev2.SendTokensInput{
			DestinationBlockchainID:            l1Info.BlockchainID,
			DestinationTokenTransferrerAddress: homeAddress,
			Recipient:                          backRecipientAddress,
			PrimaryFeeTokenAddress:             remoteAddress,
			PrimaryFee:                         big.NewInt(0),
			SecondaryFee:                       big.NewInt(0),
			RequiredGasLimit:                   utils.DefaultERC20RequiredGas,
			MultiHopFallback:                   common.Address{},
		},
		backAmount,
	)
	Expect(err).Should(BeNil())
	sendBackReceipt := utils.WaitForTransactionSuccess(ctx, ethInfo.EthClient, tx.Hash())
	lockedBefore, err := token.BalanceOf(&bind.CallOpts{}, homeAddress)
	Expect(err).Should(BeNil())

	// A second bridge messenger on the L1 verifies the message, but the home's route for Ethereum does
	// not include it, so the home rejects the delivery and the message fails to execute.
	secondMessengerAddress, secondMessenger := deploySecondMessenger(
		ctx, fundedAvalancheKey, l1Info, ethereumBlockchainID, ecdsaVerifierContractAddress,
	)
	sent := bridge.toAvalanche.sentMessage(sendBackReceipt)
	sent.OriginTeleporterAddress = secondMessengerAddress
	receipt = submitMessage(ctx, &l1Info, fundedAvalancheKey, secondMessenger, bridge.toAvalanche.attester.attest(sent))
	failedEvent, err := utils.GetEventFromLogsEmittedBy(
		receipt.Logs, secondMessengerAddress, secondMessenger.ParseMessageExecutionFailed,
	)
	Expect(err).Should(BeNil(), "the home accepted a message from a messenger outside its route")
	Expect(failedEvent.Message.DestinationAddress).Should(Equal(homeAddress))
	expectERC20Balance(token, homeAddress, lockedBefore)
	expectERC20Balance(token, backRecipientAddress, big.NewInt(0))

	// The connection's messenger delivers the same message, and the home releases the tokens
	receipt = bridge.toAvalanche.deliver(ctx, sendBackReceipt)
	homeWithdrawEvent, err := utils.GetEventFromLogsEmittedBy(receipt.Logs, homeAddress, home.ParseTokensWithdrawn)
	Expect(err).Should(BeNil())
	Expect(homeWithdrawEvent.Recipient).Should(Equal(backRecipientAddress))
	utils.ExpectBigEqual(homeWithdrawEvent.Amount, backAmount)
	expectERC20Balance(token, backRecipientAddress, backAmount)
	checkBacking(home, homeAddress, token, ethRemote, cChainBackedRemote)
}

// backedRemote is a remote of the home, with the chain it is on.
type backedRemote struct {
	blockchainID ids.ID
	address      common.Address
	remote       *erc20tokenremotev2.ERC20TokenRemote
}

// checkBacking checks the backing rule: each remote's total supply equals the tokens the home recorded
// as sent to it, and the home holds exactly their sum. The token uses the same decimals on every chain,
// so no scaling applies.
func checkBacking(
	home *erc20tokenhomev2.ERC20TokenHome,
	homeAddress common.Address,
	token *exampleerc20.ExampleERC20Decimals,
	remotes ...backedRemote,
) {
	locked := big.NewInt(0)
	for _, r := range remotes {
		totalSupply, err := r.remote.TotalSupply(&bind.CallOpts{})
		Expect(err).Should(BeNil())
		transferredBalance, err := home.GetTransferredBalance(&bind.CallOpts{}, r.blockchainID, r.address)
		Expect(err).Should(BeNil())
		utils.ExpectBigEqual(totalSupply, transferredBalance)
		locked.Add(locked, transferredBalance)
	}
	expectERC20Balance(token, homeAddress, locked)
}

type erc20 interface {
	BalanceOf(opts *bind.CallOpts, account common.Address) (*big.Int, error)
}

func expectERC20Balance(token erc20, account common.Address, expected *big.Int) {
	balance, err := token.BalanceOf(&bind.CallOpts{}, account)
	Expect(err).Should(BeNil())
	utils.ExpectBigEqual(balance, expected)
}

// deployWarpMessenger deploys the Teleporter (V1) messenger on [chains] with Nick's method, skipping any
// chain that already has it, and returns its address.
func deployWarpMessenger(
	ctx context.Context,
	fundedKey *ecdsa.PrivateKey,
	chains ...testinfo.L1TestInfo,
) common.Address {
	messengerAddress, deployerAddress, _ := utils.TeleporterDeploymentValues()
	deployerTransaction := utils.TeleporterDeployerTransaction()
	for _, chain := range chains {
		code, err := chain.EthClient.CodeAt(ctx, messengerAddress, nil)
		Expect(err).Should(BeNil())
		if len(code) > 0 {
			continue
		}
		utils.DeployWithNicksMethod(ctx, &chain, deployerTransaction, deployerAddress, messengerAddress, fundedKey)
	}
	return messengerAddress
}

// deploySecondMessenger deploys another TeleporterMessengerV2 on the L1, with its own adapter that
// accepts messages from Ethereum signed for the mock ECDSAVerifier, as the connection's adapter does.
func deploySecondMessenger(
	ctx context.Context,
	fundedKey *ecdsa.PrivateKey,
	l1Info testinfo.L1TestInfo,
	ethereumBlockchainID ids.ID,
	ecdsaVerifierAddress common.Address,
) (common.Address, *teleportermessengerv2.TeleporterMessengerV2) {
	// The adapter's address follows from its constructor arguments. Its Ethereum-bound side is never
	// used, and a placeholder there keeps it apart from the connection's adapter.
	adapterAddress := utils.DeployAdapter(
		ctx, &l1Info, l1Info.BlockchainID, ethereumBlockchainID, ecdsaVerifierAddress,
		common.HexToAddress("0x000000000000000000000000000000000000dEaD"), fundedKey,
	)
	messengerAddress := utils.DeployTeleporterV2(ctx, &l1Info, adapterAddress, fundedKey, fundedKey)
	messenger, err := teleportermessengerv2.NewTeleporterMessengerV2(messengerAddress, l1Info.EthClient)
	Expect(err).Should(BeNil())
	return messengerAddress, messenger
}
