// Copyright (C) 2026, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

// THIS IS AN EXAMPLE OF UNAUDITED CODE. DO NOT USE THIS IN PRODUCTION.

package lightclient

import (
	"context"
	"errors"
	"fmt"
	"math/big"
	"time"

	"github.com/ava-labs/avalanchego/ids"
	"github.com/ava-labs/avalanchego/utils/logging"
	"github.com/ava-labs/avalanchego/utils/set"
	teleportermessengerv2 "github.com/ava-labs/icm-services/abi-bindings/go/TeleporterMessengerV2"
	teleporterUtils "github.com/ava-labs/icm-services/icm-contracts/utils/teleporter-utils"
	"github.com/ava-labs/icm-services/messages"
	"github.com/ava-labs/icm-services/messages/zk/proofs"
	"github.com/ava-labs/icm-services/messages/zk/proofs/beacon"
	"github.com/ava-labs/icm-services/relayer/config"
	"github.com/ava-labs/icm-services/signature-aggregator/aggregator"
	"github.com/ava-labs/icm-services/vms"
	"github.com/ava-labs/icm-services/vms/evm"
	"github.com/ava-labs/libevm/accounts/abi/bind"
	"github.com/ava-labs/libevm/common"
	"github.com/ava-labs/libevm/common/hexutil"
	"github.com/ava-labs/libevm/core/types"
	"github.com/ava-labs/libevm/rpc"
	"go.uber.org/zap"
)

const (
	// defaultPollInterval is how often the handler checks whether the light client covers the
	// message: one Ethereum slot.
	defaultPollInterval = 12 * time.Second
	// defaultMaxWait bounds the wait for the light client. Ethereum finalizes in about 13 minutes,
	// and the light client follows within a few more.
	defaultMaxWait = time.Hour
	// deliveryGasOverhead is the gas the destination needs on top of the message's required gas
	// limit: verifying the execution and receipt proofs took 0.4-0.7M gas on Fuji, depending on the
	// receipt's size, plus the messenger's own work.
	deliveryGasOverhead = 1_500_000
)

// Config holds the settings of a zk-light-client message contract entry. The entry is keyed by
// the LightClientAdapter that emits the messages on the source chain.
type Config struct {
	// RewardAddress is credited as the relayer reward recipient on the destination.
	RewardAddress common.Address
	// LightClientAddress is the ZK light client of the source chain, on the destination chain.
	LightClientAddress common.Address
	// TeleporterAddress is TeleporterMessengerV2 on the destination chain.
	TeleporterAddress common.Address
	// BeaconAPIURL is a beacon node of the source chain, used to build execution proofs.
	BeaconAPIURL string
	// PollInterval is how often to check whether the light client covers a message.
	PollInterval time.Duration
	// MaxWait bounds the wait for the light client to cover a message.
	MaxWait time.Duration
}

// ConfigFromMap reads the settings of a zk-light-client message contract entry.
func ConfigFromMap(m map[string]any) (*Config, error) {
	address := func(key string) (common.Address, error) {
		s, ok := m[key].(string)
		if !ok || !common.IsHexAddress(s) {
			return common.Address{}, fmt.Errorf("%s: want a hex address, got %v", key, m[key])
		}
		return common.HexToAddress(s), nil
	}
	seconds := func(key string, fallback time.Duration) (time.Duration, error) {
		v, ok := m[key]
		if !ok {
			return fallback, nil
		}
		var s float64
		switch n := v.(type) {
		case float64:
			s = n
		case int:
			s = float64(n)
		case int64:
			s = float64(n)
		case uint64:
			s = float64(n)
		}
		if s <= 0 {
			return 0, fmt.Errorf("%s: want a positive number of seconds, got %v", key, v)
		}
		return time.Duration(s * float64(time.Second)), nil
	}

	var (
		c   Config
		err error
	)
	if c.RewardAddress, err = address("reward-address"); err != nil {
		return nil, err
	}
	if c.LightClientAddress, err = address("light-client-address"); err != nil {
		return nil, err
	}
	if c.TeleporterAddress, err = address("teleporter-address"); err != nil {
		return nil, err
	}
	if c.BeaconAPIURL, _ = m["beacon-api-url"].(string); c.BeaconAPIURL == "" {
		return nil, errors.New("beacon-api-url: want a beacon node URL")
	}
	if c.PollInterval, err = seconds("poll-interval-seconds", defaultPollInterval); err != nil {
		return nil, err
	}
	if c.MaxWait, err = seconds("max-wait-seconds", defaultMaxWait); err != nil {
		return nil, err
	}
	return &c, nil
}

type factory struct {
	config  *Config
	adapter common.Address
	source  *rpc.Client
	chainID uint64
	forks   proofs.ForkConfig
}

// NewMessageHandlerFactory returns the factory of handlers for the messages that the
// LightClientAdapter at adapter emits on the Ethereum chain served at sourceRPCURL. They are
// delivered with a proof against the beacon roots of the ZK light client on the destination.
func NewMessageHandlerFactory(
	adapter common.Address,
	cfg config.MessageProtocolConfig,
	sourceRPCURL string,
) (messages.MessageHandlerFactory, error) {
	c, err := ConfigFromMap(cfg.Settings)
	if err != nil {
		return nil, fmt.Errorf("invalid zk-light-client message protocol config: %w", err)
	}
	source, err := rpc.Dial(sourceRPCURL)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to the source chain: %w", err)
	}
	var chainID hexutil.Uint64
	if err := source.CallContext(context.Background(), &chainID, "eth_chainId"); err != nil {
		return nil, fmt.Errorf("failed to read the source chain ID: %w", err)
	}
	forks, err := forksOf(uint64(chainID))
	if err != nil {
		return nil, err
	}
	return &factory{config: c, adapter: adapter, source: source, chainID: uint64(chainID), forks: forks}, nil
}

// forksOf is the fork schedule proofs are built with on the Ethereum chain with EVM chain ID chainID.
func forksOf(chainID uint64) (proofs.ForkConfig, error) {
	switch chainID {
	case 1:
		return proofs.MainnetForkConfig, nil
	case 11155111:
		return proofs.SepoliaForkConfig, nil
	case 560048:
		return proofs.HoodiForkConfig, nil
	}
	return proofs.ForkConfig{}, fmt.Errorf("no fork schedule for source chain %d", chainID)
}

// EventFilter matches the adapter's TeleporterV2MessageSent logs.
func (f *factory) EventFilter() evm.EventFilter {
	return evm.EventFilter{
		Addresses: []common.Address{f.adapter},
		Topics:    [][]common.Hash{{messageSentTopic}},
	}
}

// GetMessageRoutingInfo routes the message by the adapter on both ends, as the zk message
// format does.
func (*factory) GetMessageRoutingInfo(message *messages.SourceMessage) (messages.MessageRoutingInfo, error) {
	m, err := decodeMessageSent(message.Payload)
	if err != nil {
		return messages.MessageRoutingInfo{}, err
	}
	return messages.MessageRoutingInfo{
		SourceChainID:      message.SourceBlockchainID,
		SenderAddress:      message.ProtocolAddress,
		DestinationChainID: ids.ID(m.DestinationBlockchainID),
		DestinationAddress: message.ProtocolAddress,
	}, nil
}

// NewMessageHandler locates the message's log in its source transaction. The signature
// aggregator, signing subnet and quorum are unused: the attestation is a proof against the light
// client's roots, not validator signatures.
func (f *factory) NewMessageHandler(
	logger logging.Logger,
	message *messages.SourceMessage,
	destinationClient vms.DestinationClient,
	_ *aggregator.SignatureAggregator,
	metrics messages.Metrics,
	_ ids.ID,
	_ uint64,
) (messages.MessageHandler, error) {
	if message.SourceTxID == (common.Hash{}) {
		return nil, errors.New("zk-light-client messages need their source transaction")
	}
	ctx := context.Background()
	sent, err := findSentMessage(ctx, f.source, message.SourceTxID, f.adapter, message.Payload)
	if err != nil {
		return nil, err
	}
	destination := destinationClient.Client()
	messenger, err := teleportermessengerv2.NewTeleporterMessengerV2(f.config.TeleporterAddress, destination)
	if err != nil {
		return nil, fmt.Errorf("failed to bind TeleporterMessengerV2: %w", err)
	}
	roots, err := NewLightClientRoots(f.config.LightClientAddress, destination)
	if err != nil {
		return nil, err
	}
	messageID, err := teleporterUtils.CalculateMessageID(
		f.config.TeleporterAddress,
		ids.ID(common.BigToHash(new(big.Int).SetUint64(f.chainID))),
		ids.ID(sent.Message.DestinationBlockchainID),
		sent.Message.MessageNonce,
	)
	if err != nil {
		return nil, fmt.Errorf("failed to compute the message ID: %w", err)
	}
	builder := proofs.NewProofBuilderWithForks(beacon.NewBeaconClient(f.config.BeaconAPIURL), f.source, f.forks)
	return &messageHandler{
		logger: logger.With(
			zap.Stringer("sourceTxID", message.SourceTxID),
			zap.Uint64("sourceBlock", sent.BlockNumber),
			zap.Stringer("messageID", messageID),
		),
		sent:        sent,
		config:      f.config,
		destination: destinationClient,
		metrics:     metrics,
		attest: func(ctx context.Context) (teleportermessengerv2.TeleporterICMMessage, error) {
			return Attest(ctx, builder, f.source, roots, sent, f.chainID)
		},
		delivered: func(ctx context.Context) (bool, error) {
			return messenger.MessageReceived(&bind.CallOpts{Context: ctx}, messageID)
		},
	}, nil
}

type messageHandler struct {
	logger      logging.Logger
	sent        *SentMessage
	config      *Config
	destination vms.DestinationClient
	metrics     messages.Metrics
	// attest builds the delivery against the light client's latest root, or returns
	// ErrNotCovered while that root does not cover the message.
	attest func(context.Context) (teleportermessengerv2.TeleporterICMMessage, error)
	// delivered reports whether the destination has received the message.
	delivered func(context.Context) (bool, error)
}

// ProcessMessage waits until the light client covers the message, then delivers it to the
// destination's TeleporterMessengerV2 with its proof. It returns the zero hash when the message
// is not this relayer's to deliver or has already been delivered.
func (h *messageHandler) ProcessMessage() (common.Hash, error) {
	ctx, cancel := context.WithTimeout(context.Background(), h.config.MaxWait)
	defer cancel()

	allowed := h.sent.Message.AllowedRelayerAddresses
	if !containsAllowedRelayer(allowed, h.destination.SenderAddresses()) {
		h.logger.Info("Relayer EOA not allowed to deliver this message")
		return common.Hash{}, nil
	}
	done, err := h.delivered(ctx)
	if err != nil {
		return common.Hash{}, fmt.Errorf("failed to check delivery: %w", err)
	}
	if done {
		h.logger.Info("Message already delivered to destination")
		return common.Hash{}, nil
	}

	h.logger.Info("Waiting for the light client to cover the message")
	message, err := h.waitForAttestation(ctx)
	if err != nil {
		h.metrics.IncFailedRelayMessageCount("failed to attest the message against the light client")
		return common.Hash{}, err
	}
	parsed, err := teleportermessengerv2.TeleporterMessengerV2MetaData.GetAbi()
	if err != nil {
		return common.Hash{}, err
	}
	callData, err := parsed.Pack("receiveCrossChainMessage", message, h.config.RewardAddress)
	if err != nil {
		return common.Hash{}, fmt.Errorf("failed to pack the delivery: %w", err)
	}
	receipt, err := h.destination.SendTx(
		h.logger,
		nil,
		set.Of(allowed...),
		h.config.TeleporterAddress,
		deliveryGasLimit(h.sent.Message.RequiredGasLimit, h.destination.BlockGasLimit()),
		callData,
	)
	if err != nil {
		h.metrics.IncFailedRelayMessageCount("failed to send the delivery")
		return common.Hash{}, fmt.Errorf("failed to send the delivery: %w", err)
	}
	if receipt.Status != types.ReceiptStatusSuccessful {
		// Another relayer may have delivered the message first.
		if done, err := h.delivered(ctx); err == nil && done {
			h.logger.Info("Delivery reverted, but the message was delivered", zap.Stringer("txID", receipt.TxHash))
			return receipt.TxHash, nil
		}
		h.metrics.IncFailedRelayMessageCount("delivery reverted")
		return common.Hash{}, fmt.Errorf("delivery %s reverted", receipt.TxHash)
	}
	h.logger.Info("Delivered message to destination chain", zap.Stringer("txID", receipt.TxHash))
	h.metrics.IncSuccessfulRelayMessageCount()
	return receipt.TxHash, nil
}

// waitForAttestation retries the attestation every PollInterval until it succeeds or ctx ends.
// Besides a light client that does not cover the message yet, a failure is usually a read of
// Ethereum or the beacon node that a later attempt gets through, so it is retried too.
func (h *messageHandler) waitForAttestation(ctx context.Context) (teleportermessengerv2.TeleporterICMMessage, error) {
	for {
		message, err := h.attest(ctx)
		if err == nil {
			return message, nil
		}
		if !errors.Is(err, ErrNotCovered) {
			h.logger.Warn("Failed to attest the message, retrying", zap.Error(err))
		}
		select {
		case <-ctx.Done():
			return message, fmt.Errorf("no attestation after %s: %w", h.config.MaxWait, err)
		case <-time.After(h.config.PollInterval):
		}
	}
}

// deliveryGasLimit is the message's required gas plus the cost of its proofs, within the block.
func deliveryGasLimit(required *big.Int, blockGasLimit uint64) uint64 {
	limit := deliveryGasOverhead + required.Uint64()
	if limit > blockGasLimit {
		return blockGasLimit
	}
	return limit
}

// containsAllowedRelayer reports whether one of eoas may deliver a message that names
// allowedRelayers; an empty list allows anyone.
func containsAllowedRelayer(allowedRelayers []common.Address, eoas []common.Address) bool {
	if len(allowedRelayers) == 0 {
		return true
	}
	for _, a := range allowedRelayers {
		for _, e := range eoas {
			if a == e {
				return true
			}
		}
	}
	return false
}
