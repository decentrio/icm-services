// Copyright (C) 2026, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package lightclient

import (
	"context"
	"errors"
	"math/big"
	"testing"
	"time"

	"github.com/ava-labs/avalanchego/ids"
	"github.com/ava-labs/avalanchego/utils/logging"
	teleportermessengerv2 "github.com/ava-labs/icm-services/abi-bindings/go/TeleporterMessengerV2"
	"github.com/ava-labs/icm-services/messages"
	"github.com/ava-labs/icm-services/vms/evm"
	"github.com/ava-labs/icm-services/vms/mocks"
	"github.com/ava-labs/libevm/common"
	"github.com/ava-labs/libevm/core/types"
	"github.com/ava-labs/libevm/crypto"
	"github.com/stretchr/testify/require"
	"go.uber.org/mock/gomock"
)

var (
	testTeleporter = common.HexToAddress("0xa2ae68a3F1E5C9fE71dEDfA681900d178c2cD79a")
	testReward     = common.HexToAddress("0x000000000000000000000000000000000000bEEF")
	testRelayer    = common.HexToAddress("0x97B1863e137D8B78C7308Dd82503450de2B2142c")
	testDelivery   = common.HexToHash("0x2de038b32865de8856218c3a6a2593def3a7b0b67107c6fcd89f2d586b1595f2")
)

type countingMetrics struct{ succeeded, failed int }

func (m *countingMetrics) IncSuccessfulRelayMessageCount()       { m.succeeded++ }
func (m *countingMetrics) IncFailedRelayMessageCount(string)     { m.failed++ }
func (*countingMetrics) IncAbandonedRelayMessageCount(string)    {}
func (*countingMetrics) IncFetchSignatureAppRequestCount()       {}
func (*countingMetrics) SetCreateSignedMessageLatencyMS(float64) {}

// testHandler relays the fork receipt's message to dest, attesting and checking delivery with the
// given functions.
func testHandler(
	t *testing.T,
	dest *mocks.MockDestinationClient,
	attest func(context.Context) (teleportermessengerv2.TeleporterICMMessage, error),
	delivered func(context.Context) (bool, error),
) (*messageHandler, *countingMetrics) {
	t.Helper()
	sent, err := FindSentMessage(context.Background(), loadForkReceipt(t), forkTx, forkAdapter)
	require.NoError(t, err)
	metrics := &countingMetrics{}
	return &messageHandler{
		logger: logging.NoLog{},
		sent:   sent,
		config: &Config{
			RewardAddress:     testReward,
			TeleporterAddress: testTeleporter,
			PollInterval:      time.Millisecond,
			MaxWait:           time.Second,
		},
		destination: dest,
		metrics:     metrics,
		attest:      attest,
		delivered:   delivered,
	}, metrics
}

func testICMMessage(t *testing.T, h *messageHandler) teleportermessengerv2.TeleporterICMMessage {
	t.Helper()
	m, err := ICMMessage(h.sent, 11155111, []byte{0x01, 0x02})
	require.NoError(t, err)
	return m
}

func notDelivered(context.Context) (bool, error) { return false, nil }

// receiveCallData is the call TeleporterMessengerV2's ABI defines for delivering m.
func receiveCallData(t *testing.T, m teleportermessengerv2.TeleporterICMMessage) []byte {
	t.Helper()
	parsed, err := teleportermessengerv2.TeleporterMessengerV2MetaData.GetAbi()
	require.NoError(t, err)
	data, err := parsed.Pack("receiveCrossChainMessage", m, testReward)
	require.NoError(t, err)
	return data
}

// Ethereum messages can only be proven once the light client finalizes past their block, about
// 13 minutes after they are sent, so the handler waits for that instead of failing.
func TestProcessMessageWaitsForTheLightClientThenDelivers(t *testing.T) {
	ctrl := gomock.NewController(t)
	dest := mocks.NewMockDestinationClient(ctrl)
	attempts := 0
	var message teleportermessengerv2.TeleporterICMMessage
	h, metrics := testHandler(t, dest, func(context.Context) (teleportermessengerv2.TeleporterICMMessage, error) {
		attempts++
		if attempts < 3 {
			return teleportermessengerv2.TeleporterICMMessage{}, ErrNotCovered
		}
		return message, nil
	}, notDelivered)
	message = testICMMessage(t, h)

	dest.EXPECT().SenderAddresses().Return([]common.Address{testRelayer}).AnyTimes()
	dest.EXPECT().BlockGasLimit().Return(uint64(15_000_000)).AnyTimes()
	dest.EXPECT().SendTx(gomock.Any(), gomock.Nil(), gomock.Any(), testTeleporter,
		uint64(250_000+deliveryGasOverhead), receiveCallData(t, message)).
		Return(&types.Receipt{Status: types.ReceiptStatusSuccessful, TxHash: testDelivery}, nil)

	txHash, err := h.ProcessMessage()
	require.NoError(t, err)
	require.Equal(t, testDelivery, txHash)
	require.Equal(t, 3, attempts)
	require.Equal(t, 1, metrics.succeeded)
}

// A message already delivered, by this relayer before a restart or by another one, is not sent
// again.
func TestProcessMessageSkipsADeliveredMessage(t *testing.T) {
	ctrl := gomock.NewController(t)
	dest := mocks.NewMockDestinationClient(ctrl)
	h, _ := testHandler(t, dest, func(context.Context) (teleportermessengerv2.TeleporterICMMessage, error) {
		t.Fatal("attested a delivered message")
		return teleportermessengerv2.TeleporterICMMessage{}, nil
	}, func(context.Context) (bool, error) { return true, nil })
	dest.EXPECT().SenderAddresses().Return([]common.Address{testRelayer}).AnyTimes()

	txHash, err := h.ProcessMessage()
	require.NoError(t, err)
	require.Equal(t, common.Hash{}, txHash)
}

// A message that names its relayers is left to them.
func TestProcessMessageLeavesAMessageToItsAllowedRelayers(t *testing.T) {
	ctrl := gomock.NewController(t)
	dest := mocks.NewMockDestinationClient(ctrl)
	h, _ := testHandler(t, dest, func(context.Context) (teleportermessengerv2.TeleporterICMMessage, error) {
		t.Fatal("attested a message this relayer may not deliver")
		return teleportermessengerv2.TeleporterICMMessage{}, nil
	}, notDelivered)
	h.sent.Message.AllowedRelayerAddresses = []common.Address{common.HexToAddress("0x1")}
	dest.EXPECT().SenderAddresses().Return([]common.Address{testRelayer}).AnyTimes()

	txHash, err := h.ProcessMessage()
	require.NoError(t, err)
	require.Equal(t, common.Hash{}, txHash)
}

// A delivery that reverts because another relayer delivered first is not a failure.
func TestProcessMessageAcceptsARevertWhenTheMessageWasDelivered(t *testing.T) {
	ctrl := gomock.NewController(t)
	dest := mocks.NewMockDestinationClient(ctrl)
	checks := 0
	var message teleportermessengerv2.TeleporterICMMessage
	h, _ := testHandler(t, dest, func(context.Context) (teleportermessengerv2.TeleporterICMMessage, error) {
		return message, nil
	}, func(context.Context) (bool, error) {
		checks++
		return checks > 1, nil
	})
	message = testICMMessage(t, h)
	dest.EXPECT().SenderAddresses().Return([]common.Address{testRelayer}).AnyTimes()
	dest.EXPECT().BlockGasLimit().Return(uint64(15_000_000)).AnyTimes()
	dest.EXPECT().SendTx(gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any()).
		Return(&types.Receipt{Status: types.ReceiptStatusFailed, TxHash: testDelivery}, nil)

	txHash, err := h.ProcessMessage()
	require.NoError(t, err)
	require.Equal(t, testDelivery, txHash)
}

func TestProcessMessageFailsWhenTheDeliveryReverts(t *testing.T) {
	ctrl := gomock.NewController(t)
	dest := mocks.NewMockDestinationClient(ctrl)
	var message teleportermessengerv2.TeleporterICMMessage
	h, metrics := testHandler(t, dest, func(context.Context) (teleportermessengerv2.TeleporterICMMessage, error) {
		return message, nil
	}, notDelivered)
	message = testICMMessage(t, h)
	dest.EXPECT().SenderAddresses().Return([]common.Address{testRelayer}).AnyTimes()
	dest.EXPECT().BlockGasLimit().Return(uint64(15_000_000)).AnyTimes()
	dest.EXPECT().SendTx(gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any()).
		Return(&types.Receipt{Status: types.ReceiptStatusFailed, TxHash: testDelivery}, nil)

	_, err := h.ProcessMessage()
	require.Error(t, err)
	require.Equal(t, 1, metrics.failed)
}

// A failed read of Ethereum or the beacon node, e.g. while the network drops, is retried like a
// light client that is still behind: giving up would make the relayer abandon the message.
func TestProcessMessageRetriesAttestationFailures(t *testing.T) {
	ctrl := gomock.NewController(t)
	dest := mocks.NewMockDestinationClient(ctrl)
	attempts := 0
	var message teleportermessengerv2.TeleporterICMMessage
	h, _ := testHandler(t, dest, func(context.Context) (teleportermessengerv2.TeleporterICMMessage, error) {
		attempts++
		if attempts == 1 {
			return teleportermessengerv2.TeleporterICMMessage{}, errors.New("read: can't assign requested address")
		}
		return message, nil
	}, notDelivered)
	message = testICMMessage(t, h)
	dest.EXPECT().SenderAddresses().Return([]common.Address{testRelayer}).AnyTimes()
	dest.EXPECT().BlockGasLimit().Return(uint64(15_000_000)).AnyTimes()
	dest.EXPECT().SendTx(gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any()).
		Return(&types.Receipt{Status: types.ReceiptStatusSuccessful, TxHash: testDelivery}, nil)

	txHash, err := h.ProcessMessage()
	require.NoError(t, err)
	require.Equal(t, testDelivery, txHash)
	require.Equal(t, 2, attempts)
}

// A failure that persists until MaxWait is reported.
func TestProcessMessageReportsAPersistentAttestationFailure(t *testing.T) {
	ctrl := gomock.NewController(t)
	dest := mocks.NewMockDestinationClient(ctrl)
	boom := errors.New("beacon node unavailable")
	h, _ := testHandler(t, dest, func(context.Context) (teleportermessengerv2.TeleporterICMMessage, error) {
		return teleportermessengerv2.TeleporterICMMessage{}, boom
	}, notDelivered)
	h.config.MaxWait = 20 * time.Millisecond
	dest.EXPECT().SenderAddresses().Return([]common.Address{testRelayer}).AnyTimes()

	_, err := h.ProcessMessage()
	require.ErrorIs(t, err, boom)
}

func TestProcessMessageGivesUpAfterMaxWait(t *testing.T) {
	ctrl := gomock.NewController(t)
	dest := mocks.NewMockDestinationClient(ctrl)
	h, _ := testHandler(t, dest, func(context.Context) (teleportermessengerv2.TeleporterICMMessage, error) {
		return teleportermessengerv2.TeleporterICMMessage{}, ErrNotCovered
	}, notDelivered)
	h.config.MaxWait = 20 * time.Millisecond
	dest.EXPECT().SenderAddresses().Return([]common.Address{testRelayer}).AnyTimes()

	_, err := h.ProcessMessage()
	require.ErrorIs(t, err, ErrNotCovered)
}

// The relayer reads the adapter's TeleporterV2MessageSent logs, and no other contract's.
func TestEventFilterMatchesTheAdaptersMessageLogs(t *testing.T) {
	f := &factory{adapter: forkAdapter}
	require.Equal(t, evm.EventFilter{
		Addresses: []common.Address{forkAdapter},
		Topics:    [][]common.Hash{{crypto.Keccak256Hash([]byte("TeleporterV2MessageSent(bytes)"))}},
	}, f.EventFilter())
}

func TestRoutingInfoNamesTheMessagesDestinationChain(t *testing.T) {
	var receipt struct {
		Logs []struct {
			Data string `json:"data"`
		} `json:"logs"`
	}
	readJSON(t, "testdata/sent_message_receipt.json", &receipt)
	source := ids.GenerateTestID()
	info, err := (&factory{adapter: forkAdapter}).GetMessageRoutingInfo(&messages.SourceMessage{
		SourceBlockchainID: source,
		ProtocolAddress:    forkAdapter,
		Payload:            common.FromHex(receipt.Logs[0].Data),
		SourceTxID:         forkTx,
	})
	require.NoError(t, err)
	require.Equal(t, source, info.SourceChainID)
	require.Equal(t, ids.ID(common.HexToHash("0x7fc93d85c6d62c5b2ac0b519c87010ea5294012d1e407030d6acd0021cac10d5")),
		info.DestinationChainID)
}

// A transaction may send several messages: the handler relays the one whose log the relayer read.
func TestFindSentMessageMatchesThePayload(t *testing.T) {
	_, err := findSentMessage(context.Background(), loadForkReceipt(t), forkTx, forkAdapter, []byte{0x01})
	require.ErrorIs(t, err, ErrNoMessage)
}

func TestConfigFromMap(t *testing.T) {
	valid := func() map[string]any {
		return map[string]any{
			"reward-address":       testReward.Hex(),
			"light-client-address": "0x4Fc01ae40bf56e5B517Af7B1f6424D1074313E5c",
			"teleporter-address":   testTeleporter.Hex(),
			"beacon-api-url":       "https://lodestar-sepolia.chainsafe.io",
		}
	}
	c, err := ConfigFromMap(valid())
	require.NoError(t, err)
	require.Equal(t, testTeleporter, c.TeleporterAddress)
	require.Equal(t, defaultPollInterval, c.PollInterval)
	require.Equal(t, defaultMaxWait, c.MaxWait)

	m := valid()
	m["poll-interval-seconds"] = float64(3)
	m["max-wait-seconds"] = 600
	c, err = ConfigFromMap(m)
	require.NoError(t, err)
	require.Equal(t, 3*time.Second, c.PollInterval)
	require.Equal(t, 10*time.Minute, c.MaxWait)

	for _, key := range []string{"reward-address", "light-client-address", "teleporter-address", "beacon-api-url"} {
		m := valid()
		delete(m, key)
		_, err := ConfigFromMap(m)
		require.ErrorContains(t, err, key)
	}
	m = valid()
	m["light-client-address"] = "not an address"
	_, err = ConfigFromMap(m)
	require.ErrorContains(t, err, "light-client-address")
	m = valid()
	m["poll-interval-seconds"] = float64(0)
	_, err = ConfigFromMap(m)
	require.ErrorContains(t, err, "poll-interval-seconds")
}

func TestGasLimitCoversTheRequiredGasPlusTheProofsWithinTheBlock(t *testing.T) {
	require.Equal(t, uint64(250_000+deliveryGasOverhead), deliveryGasLimit(big.NewInt(250_000), 15_000_000))
	require.Equal(t, uint64(8_000_000), deliveryGasLimit(big.NewInt(9_000_000), 8_000_000))
}
