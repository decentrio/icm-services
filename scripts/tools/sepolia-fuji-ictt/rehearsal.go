// Copyright (C) 2026, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

// THIS IS AN EXAMPLE OF UNAUDITED CODE. DO NOT USE THIS IN PRODUCTION.

//go:build rehearsal

package main

import (
	"context"
	"errors"
	"flag"
	"fmt"

	teleportermessengerv2 "github.com/ava-labs/icm-services/abi-bindings/go/TeleporterMessengerV2"
	"github.com/ava-labs/icm-services/messages/zk/lightclient"
	"github.com/ava-labs/libevm/common"
	"github.com/ava-labs/libevm/common/hexutil"
	"github.com/ava-labs/libevm/rpc"
)

// acceptAll is runtime code that returns true to any call: it stands in for the
// LightClientAdapter while a message from a Sepolia fork, which no beacon root covers, is
// delivered on a Fuji fork.
var acceptAll = common.FromHex("0x600160005260206000f3")

func init() {
	commands["rehearse-delivery"] = rehearseDelivery
}

// rehearseDelivery delivers a message sent on a Sepolia fork to a Fuji fork without a proof, to
// rehearse everything after verification: it swaps the LightClientAdapter's code on the Fuji fork
// for acceptAll (anvil_setCode), delivers, and puts the code back. Forks only: real chains do not
// answer anvil_setCode.
func rehearseDelivery(ctx context.Context, n networks, args []string) error {
	fs := flag.NewFlagSet("rehearse-delivery", flag.ExitOnError)
	tx := fs.String("tx", "", "transaction on the Sepolia fork that sent the message")
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
	sent, err := lightclient.FindSentMessage(ctx, el, common.HexToHash(*tx), dep.LightClientAdapter)
	if err != nil {
		return err
	}
	message, err := lightclient.ICMMessage(sent, sepoliaChainID, nil)
	if err != nil {
		return err
	}
	fmt.Printf("message %s from %s to %s\n",
		sent.Message.MessageNonce, sent.Message.OriginSenderAddress, sent.Message.DestinationAddress)

	fujiRPC, err := rpc.DialContext(ctx, n.fujiRPC)
	if err != nil {
		return err
	}
	var original hexutil.Bytes
	if err := fujiRPC.CallContext(ctx, &original, "eth_getCode", dep.LightClientAdapter, "latest"); err != nil {
		return err
	}
	err = fujiRPC.CallContext(ctx, nil, "anvil_setCode", dep.LightClientAdapter, hexutil.Bytes(acceptAll))
	if err != nil {
		return fmt.Errorf("not a fork (anvil_setCode): %w", err)
	}
	defer func() {
		if err := fujiRPC.CallContext(ctx, nil, "anvil_setCode", dep.LightClientAdapter, original); err != nil {
			fmt.Println("failed to restore the LightClientAdapter's code:", err)
		}
	}()

	fuji, err := dial(ctx, "Fuji", n.fujiRPC, acc)
	if err != nil {
		return err
	}
	messenger, err := teleportermessengerv2.NewTeleporterMessengerV2(dep.Messenger, fuji.client)
	if err != nil {
		return err
	}
	delivery, err := messenger.ReceiveCrossChainMessage(fuji.opts, message, common.Address{})
	receipt, err := fuji.mined("receiveCrossChainMessage (unproven)", delivery, err)
	if err != nil {
		return err
	}
	for _, l := range receipt.Logs {
		if l.Address != dep.Messenger {
			continue
		}
		if _, err := messenger.ParseMessageExecuted(*l); err == nil {
			fmt.Println("message executed on Fuji")
			return nil
		}
		if _, err := messenger.ParseMessageExecutionFailed(*l); err == nil {
			return errors.New("message delivered but its execution failed")
		}
	}
	return errors.New("message delivered but the messenger reported no execution")
}
