package ethereum_icm_verification

import (
	"context"
	"crypto/ecdsa"

	"github.com/ava-labs/avalanchego/ids"
	teleportermessengerv2 "github.com/ava-labs/icm-services/abi-bindings/go/TeleporterMessengerV2"
	ecdsaverifier "github.com/ava-labs/icm-services/abi-bindings/go/mocks/ECDSAVerifier"
	testinfo "github.com/ava-labs/icm-services/icm-contracts/tests/test-info"
	"github.com/ava-labs/icm-services/icm-contracts/tests/utils"
	teleporterUtils "github.com/ava-labs/icm-services/icm-contracts/utils/teleporter-utils"
	"github.com/ava-labs/libevm/accounts/abi/bind"
	"github.com/ava-labs/libevm/common"
	"github.com/ava-labs/libevm/core/types"
	. "github.com/onsi/gomega"
)

// A relay delivers the ICM message that a transaction sent on its source chain. It returns the receipt
// of the delivery on the destination chain, after checking that the message executed there. Each
// direction of a connection has its own relay.
type relay interface {
	deliver(ctx context.Context, sourceReceipt *types.Receipt) *types.Receipt
}

// relayerRelay waits for the ICM relayer to deliver a TeleporterMessengerV2 message from an Avalanche
// L1 to Ethereum, where the MerkleValidatorSetRegistry verifies it.
type relayerRelay struct {
	source *testinfo.L1TestInfo
	// sourceKey funds the transactions that advance the source chain.
	sourceKey            *ecdsa.PrivateKey
	destination          *testinfo.EthereumTestInfo
	messengerAddress     common.Address
	sourceMessenger      *teleportermessengerv2.TeleporterMessengerV2
	destinationMessenger *teleportermessengerv2.TeleporterMessengerV2
}

func (r *relayerRelay) deliver(ctx context.Context, sourceReceipt *types.Receipt) *types.Receipt {
	sendEvent, err := utils.GetEventFromLogsEmittedBy(
		sourceReceipt.Logs, r.messengerAddress, r.sourceMessenger.ParseSendCrossChainMessage,
	)
	Expect(err).Should(BeNil())
	messageID, err := teleporterUtils.CalculateMessageID(
		r.messengerAddress, r.source.BlockchainID, r.destination.ChainID(), sendEvent.Message.MessageNonce,
	)
	Expect(err).Should(BeNil())

	// subnet-evm only builds blocks when it has transactions to include, so advance the source chain
	// for the relayer's subscription to surface the block that holds the message.
	Expect(utils.IssueTxsToAdvanceChain(
		ctx, r.source.EVMChainID, r.sourceKey, r.source.EthClient, 5,
	)).Should(BeNil())
	Expect(utils.WaitTeleporterMessageDeliveredV2(ctx, r.destinationMessenger, messageID)).Should(
		BeNil(), "the relayer did not deliver the message to Ethereum",
	)

	deliveries, err := r.destinationMessenger.FilterReceiveCrossChainMessage(
		&bind.FilterOpts{Context: ctx}, [][32]byte{messageID}, nil, nil,
	)
	Expect(err).Should(BeNil())
	defer deliveries.Close()
	Expect(deliveries.Next()).Should(BeTrue(), "no ReceiveCrossChainMessage log for the delivered message")
	receipt, err := r.destination.EthClient.TransactionReceipt(ctx, deliveries.Event.Raw.TxHash)
	Expect(err).Should(BeNil())

	_, err = utils.GetEventFromLogsEmittedBy(
		receipt.Logs, r.messengerAddress, r.destinationMessenger.ParseMessageExecuted,
	)
	Expect(err).Should(BeNil(), "message was delivered to Ethereum but did not execute")
	return receipt
}

// ecdsaAttester attests messages from Ethereum for the mock ECDSAVerifier, which accepts any message
// signed by the key it trusts.
type ecdsaAttester struct {
	ethereumBlockchainID ids.ID
	verifierAddress      common.Address
	signer               *ecdsa.PrivateKey
}

func (a *ecdsaAttester) attest(message ecdsaverifier.TeleporterMessageV2) teleportermessengerv2.TeleporterICMMessage {
	receipts := make([]teleportermessengerv2.TeleporterMessageReceipt, len(message.Receipts))
	for i, r := range message.Receipts {
		receipts[i] = teleportermessengerv2.TeleporterMessageReceipt{
			ReceivedMessageNonce: r.ReceivedMessageNonce,
			RelayerRewardAddress: r.RelayerRewardAddress,
		}
	}
	return teleportermessengerv2.TeleporterICMMessage{
		Message: teleportermessengerv2.TeleporterMessageV2{
			MessageNonce:            message.MessageNonce,
			OriginSenderAddress:     message.OriginSenderAddress,
			OriginTeleporterAddress: message.OriginTeleporterAddress,
			DestinationBlockchainID: message.DestinationBlockchainID,
			DestinationAddress:      message.DestinationAddress,
			RequiredGasLimit:        message.RequiredGasLimit,
			AllowedRelayerAddresses: message.AllowedRelayerAddresses,
			Receipts:                receipts,
			Message:                 message.Message,
		},
		SourceNetworkID:    0,
		SourceBlockchainID: a.ethereumBlockchainID,
		Attestation:        signMessageEcdsa(message, a.ethereumBlockchainID, a.verifierAddress, a.signer),
	}
}

// ecdsaRelay delivers a TeleporterMessengerV2 message from Ethereum to an Avalanche L1, attested for the
// mock ECDSAVerifier. It stands in for the ZK verifier until the relayer supports it.
type ecdsaRelay struct {
	attester *ecdsaAttester
	// sourceVerifier is the ECDSAVerifier on Ethereum, which emits the messages it sends.
	sourceVerifier              *ecdsaverifier.ECDSAVerifier
	destination                 *testinfo.L1TestInfo
	destinationKey              *ecdsa.PrivateKey
	destinationMessengerAddress common.Address
	destinationMessenger        *teleportermessengerv2.TeleporterMessengerV2
}

func (r *ecdsaRelay) deliver(ctx context.Context, sourceReceipt *types.Receipt) *types.Receipt {
	receipt := submitMessage(
		ctx, r.destination, r.destinationKey, r.destinationMessenger,
		r.attester.attest(r.sentMessage(sourceReceipt)),
	)
	_, err := utils.GetEventFromLogsEmittedBy(
		receipt.Logs, r.destinationMessengerAddress, r.destinationMessenger.ParseMessageExecuted,
	)
	Expect(err).Should(BeNil(), "message was delivered to the L1 but did not execute")
	return receipt
}

// sentMessage returns the message that the ECDSAVerifier on Ethereum sent in [sourceReceipt].
func (r *ecdsaRelay) sentMessage(sourceReceipt *types.Receipt) ecdsaverifier.TeleporterMessageV2 {
	sendEvent, err := utils.GetEventFromLogsEmittedBy(
		sourceReceipt.Logs, r.attester.verifierAddress, r.sourceVerifier.ParseECDSAVerifierSendMessage,
	)
	Expect(err).Should(BeNil())
	return sendEvent.Message
}

// warpRelay delivers a TeleporterMessenger (V1) message between two Avalanche chains, attested by the
// source chain's validators through Warp.
type warpRelay struct {
	teleporter  utils.TeleporterTestInfo
	source      testinfo.L1TestInfo
	destination testinfo.L1TestInfo
	relayerKey  *ecdsa.PrivateKey
	aggregator  *utils.SignatureAggregator
}

func (r *warpRelay) deliver(ctx context.Context, sourceReceipt *types.Receipt) *types.Receipt {
	receipt := r.teleporter.RelayTeleporterMessage(
		ctx, sourceReceipt, r.source, r.destination, true, r.relayerKey, nil, r.aggregator,
	)
	_, err := utils.GetEventFromLogsEmittedBy(
		receipt.Logs,
		r.teleporter.TeleporterMessengerAddress(r.destination.BlockchainID),
		r.teleporter.TeleporterMessenger(&r.destination).ParseMessageExecuted,
	)
	Expect(err).Should(BeNil(), "message was delivered but did not execute")
	return receipt
}

// submitMessage delivers an attested message to [messenger] on [chain] and returns the receipt, whether
// or not the message executed.
func submitMessage(
	ctx context.Context,
	chain *testinfo.L1TestInfo,
	key *ecdsa.PrivateKey,
	messenger *teleportermessengerv2.TeleporterMessengerV2,
	message teleportermessengerv2.TeleporterICMMessage,
) *types.Receipt {
	opts, err := bind.NewKeyedTransactorWithChainID(key, chain.EVMChainID)
	Expect(err).Should(BeNil())
	tx, err := messenger.ReceiveCrossChainMessage(opts, message, common.Address{})
	Expect(err).Should(BeNil())
	return utils.WaitForTransactionSuccess(ctx, chain.EthClient, tx.Hash())
}
