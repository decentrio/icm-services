// Copyright (C) 2026, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

// THIS IS AN EXAMPLE OF UNAUDITED CODE. DO NOT USE THIS IN PRODUCTION.

package main

import (
	"context"
	"errors"
	"flag"
	"fmt"
	"time"

	teleportermessengerv2 "github.com/ava-labs/icm-services/abi-bindings/go/TeleporterMessengerV2"
	"github.com/ava-labs/icm-services/messages/zk/lightclient"
	"github.com/ava-labs/icm-services/messages/zk/proofs"
	"github.com/ava-labs/icm-services/messages/zk/proofs/beacon"
	"github.com/ava-labs/libevm/common"
	"github.com/ava-labs/libevm/rpc"
)

// relay delivers a message sent on Sepolia to Fuji: it waits until the ZK light client covers the
// message's block, proves the adapter's log against the light client's root, and calls
// receiveCrossChainMessage on the Fuji messenger.
func relay(ctx context.Context, n networks, args []string) error {
	fs := flag.NewFlagSet("relay", flag.ExitOnError)
	tx := fs.String("tx", "", "Sepolia transaction that sent the message")
	wait := fs.Duration("wait", 45*time.Minute, "how long to wait for the light client to cover the message")
	d := deploymentFlag(fs)
	if err := fs.Parse(args); err != nil {
		return err
	}
	dep, err := d.load()
	if err != nil {
		return err
	}
	if *tx == "" {
		return errors.New("-tx is required")
	}
	acc, err := accountFromEnv()
	if err != nil {
		return err
	}

	el, err := rpc.DialContext(ctx, n.sepoliaRPC)
	if err != nil {
		return err
	}
	fuji, err := dial(ctx, "Fuji", n.fujiRPC, acc)
	if err != nil {
		return err
	}
	sent, err := lightclient.FindSentMessage(ctx, el, common.HexToHash(*tx), dep.LightClientAdapter)
	if err != nil {
		return err
	}
	fmt.Printf("message %s from %s to %s on %s, in Sepolia block %d\n",
		sent.Message.MessageNonce, sent.Message.OriginSenderAddress, sent.Message.DestinationAddress,
		common.Hash(sent.Message.DestinationBlockchainID), sent.BlockNumber)

	roots, err := lightclient.NewLightClientRoots(dep.LightClient, fuji.client)
	if err != nil {
		return err
	}
	builder := proofs.NewProofBuilderWithForks(beacon.NewBeaconClient(n.beaconURL), el, proofs.SepoliaForkConfig)
	deadline := time.Now().Add(*wait)
	var message teleportermessengerv2.TeleporterICMMessage
	for {
		message, err = lightclient.Attest(ctx, builder, el, roots, sent, sepoliaChainID)
		if err == nil {
			break
		}
		if !errors.Is(err, lightclient.ErrNotCovered) || time.Now().After(deadline) {
			return err
		}
		fmt.Println("waiting for the light client to finalize past the message's block")
		select {
		case <-ctx.Done():
			return ctx.Err()
		case <-time.After(time.Minute):
		}
	}
	fmt.Printf("attestation: %d bytes\n", len(message.Attestation))

	messenger, err := teleportermessengerv2.NewTeleporterMessengerV2(dep.Messenger, fuji.client)
	if err != nil {
		return err
	}
	delivery, err := messenger.ReceiveCrossChainMessage(fuji.opts, message, common.Address{})
	receipt, err := fuji.mined("receiveCrossChainMessage", delivery, err)
	if err != nil {
		return err
	}
	for _, l := range receipt.Logs {
		if l.Address != dep.Messenger {
			continue
		}
		if ev, err := messenger.ParseMessageExecuted(*l); err == nil {
			fmt.Printf("delivered and executed on Fuji in %s (%d gas), message ID %s\n",
				delivery.Hash(), receipt.GasUsed, common.Hash(ev.MessageID))
			return nil
		}
		if ev, err := messenger.ParseMessageExecutionFailed(*l); err == nil {
			return fmt.Errorf("delivered in %s but execution failed, message ID %s", delivery.Hash(),
				common.Hash(ev.MessageID))
		}
	}
	return fmt.Errorf("delivered in %s, but the messenger reported no execution", delivery.Hash())
}
