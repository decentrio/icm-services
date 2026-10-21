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

	remoteregistrar "github.com/ava-labs/icm-services/abi-bindings/go/ictt-v2/RemoteRegistrar"
	erc20tokenhome "github.com/ava-labs/icm-services/abi-bindings/go/ictt-v2/TokenHome/ERC20TokenHome"
	erc20tokenremote "github.com/ava-labs/icm-services/abi-bindings/go/ictt-v2/TokenRemote/ERC20TokenRemote"
	exampleerc20 "github.com/ava-labs/icm-services/abi-bindings/go/ictt/mocks/ExampleERC20Decimals"
	"github.com/ava-labs/libevm/accounts/abi/bind"
	"github.com/ava-labs/libevm/common"
)

// sendGasLimit is the gas the remote may use to mint on Fuji.
const sendGasLimit = 250_000

// deployICTT deploys the token's home on Sepolia, routed to Fuji through the messenger, and its
// remote on Fuji. Without -token it first deploys a demo ERC-20, minted to the account.
func deployICTT(ctx context.Context, n networks, args []string) error {
	fs := flag.NewFlagSet("deploy-ictt", flag.ExitOnError)
	token := fs.String("token", "", "ERC-20 on Sepolia; a demo token is deployed if empty")
	decimals := fs.Uint("decimals", 18, "decimals of the demo token")
	name := fs.String("name", "Sepolia Demo Token", "remote token's name")
	symbol := fs.String("symbol", "SDT.e", "remote token's symbol")
	d := deploymentFlag(fs)
	if err := fs.Parse(args); err != nil {
		return err
	}
	dep, err := d.load()
	if err != nil {
		return err
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

	fmt.Println("Sepolia:")
	if *token == "" {
		address, tx, _, err := exampleerc20.DeployExampleERC20Decimals(sepolia.opts, sepolia.client, uint8(*decimals))
		if _, err := sepolia.mined("deploy demo ERC-20", tx, err); err != nil {
			return err
		}
		dep.Token, dep.TokenDecimals = address, uint8(*decimals)
	} else {
		dep.Token = common.HexToAddress(*token)
		erc20, err := exampleerc20.NewExampleERC20Decimals(dep.Token, sepolia.client)
		if err != nil {
			return err
		}
		if dep.TokenDecimals, err = erc20.Decimals(&bind.CallOpts{Context: ctx}); err != nil {
			return fmt.Errorf("token decimals: %w", err)
		}
	}
	fmt.Printf("  token %s (%d decimals)\n", dep.Token, dep.TokenDecimals)

	homeAddress, tx, home, err := erc20tokenhome.DeployERC20TokenHome(
		sepolia.opts, sepolia.client, sepoliaBlockchainID, acc.address, dep.Token, dep.TokenDecimals)
	if _, err := sepolia.mined("deploy ERC20TokenHome", tx, err); err != nil {
		return err
	}
	dep.TokenHome = homeAddress
	tx, err = home.AddMessenger(sepolia.opts, fujiCChainID, dep.Messenger)
	if _, err := sepolia.mined("route Fuji through the messenger (inbound)", tx, err); err != nil {
		return err
	}
	tx, err = home.SetOutboundMessenger(sepolia.opts, fujiCChainID, dep.Messenger)
	if _, err := sepolia.mined("route Fuji through the messenger (outbound)", tx, err); err != nil {
		return err
	}
	fmt.Printf("  home %s\n", homeAddress)

	fmt.Println("Fuji:")
	remoteAddress, tx, _, err := erc20tokenremote.DeployERC20TokenRemote(fuji.opts, fuji.client,
		erc20tokenremote.TokenRemoteSettingsV2{
			BlockchainID:          fujiCChainID,
			TeleporterMessenger:   dep.Messenger,
			TeleporterManager:     acc.address,
			TokenHomeBlockchainID: sepoliaBlockchainID,
			TokenHomeAddress:      homeAddress,
			TokenHomeDecimals:     dep.TokenDecimals,
		}, *name, *symbol, dep.TokenDecimals)
	if _, err := fuji.mined("deploy ERC20TokenRemote", tx, err); err != nil {
		return err
	}
	dep.TokenRemote = remoteAddress
	fmt.Printf("  remote %s\n", remoteAddress)
	return d.save(dep)
}

// register registers the remote with the home on the home owner's word: it deploys a
// RemoteRegistrar, adds it as the home's messenger for Fuji, registers the remote, and removes it.
// Fuji's messages cannot reach Sepolia without a Fuji node of our own; see README.md.
func register(ctx context.Context, n networks, args []string) error {
	fs := flag.NewFlagSet("register", flag.ExitOnError)
	d := deploymentFlag(fs)
	if err := fs.Parse(args); err != nil {
		return err
	}
	dep, err := d.load()
	if err != nil {
		return err
	}
	acc, err := accountFromEnv()
	if err != nil {
		return err
	}
	sepolia, err := dial(ctx, "Sepolia", n.sepoliaRPC, acc)
	if err != nil {
		return err
	}
	home, err := erc20tokenhome.NewERC20TokenHome(dep.TokenHome, sepolia.client)
	if err != nil {
		return err
	}
	settings, err := home.GetRemoteTokenTransferrerSettings(&bind.CallOpts{Context: ctx}, fujiCChainID, dep.TokenRemote)
	if err != nil {
		return err
	}
	if settings.Registered {
		fmt.Println("remote already registered")
		return nil
	}

	registrarAddress, tx, registrar, err := remoteregistrar.DeployRemoteRegistrar(
		sepolia.opts, sepolia.client, dep.TokenHome, sepoliaBlockchainID)
	if _, err := sepolia.mined("deploy RemoteRegistrar", tx, err); err != nil {
		return err
	}
	tx, err = home.AddMessenger(sepolia.opts, fujiCChainID, registrarAddress)
	if _, err := sepolia.mined("add the registrar as a messenger", tx, err); err != nil {
		return err
	}
	tx, err = registrar.Register(sepolia.opts, fujiCChainID, dep.TokenRemote, dep.TokenDecimals, dep.TokenDecimals)
	if _, err := sepolia.mined("register the remote", tx, err); err != nil {
		return err
	}
	tx, err = home.RemoveMessenger(sepolia.opts, fujiCChainID, registrarAddress)
	if _, err := sepolia.mined("remove the registrar", tx, err); err != nil {
		return err
	}
	settings, err = home.GetRemoteTokenTransferrerSettings(
		&bind.CallOpts{Context: ctx}, fujiCChainID, dep.TokenRemote)
	if err != nil {
		return err
	}
	if !settings.Registered {
		return errors.New("the home does not list the remote as registered")
	}
	fmt.Printf("remote %s registered with home %s\n", dep.TokenRemote, dep.TokenHome)
	return nil
}

// send sends tokens from the account on Sepolia to a recipient on Fuji, through the home.
func send(ctx context.Context, n networks, args []string) error {
	fs := flag.NewFlagSet("send", flag.ExitOnError)
	amount := fs.String("amount", "1", "amount in whole tokens")
	recipient := fs.String("recipient", "", "recipient on Fuji (default: the account)")
	d := deploymentFlag(fs)
	if err := fs.Parse(args); err != nil {
		return err
	}
	dep, err := d.load()
	if err != nil {
		return err
	}
	acc, err := accountFromEnv()
	if err != nil {
		return err
	}
	to := acc.address
	if *recipient != "" {
		to = common.HexToAddress(*recipient)
	}
	units, err := tokenUnits(*amount, dep.TokenDecimals)
	if err != nil {
		return err
	}
	sepolia, err := dial(ctx, "Sepolia", n.sepoliaRPC, acc)
	if err != nil {
		return err
	}
	erc20, err := exampleerc20.NewExampleERC20Decimals(dep.Token, sepolia.client)
	if err != nil {
		return err
	}
	tx, err := erc20.Approve(sepolia.opts, dep.TokenHome, units)
	if _, err := sepolia.mined("approve the home", tx, err); err != nil {
		return err
	}
	home, err := erc20tokenhome.NewERC20TokenHome(dep.TokenHome, sepolia.client)
	if err != nil {
		return err
	}
	tx, err = home.Send(sepolia.opts, erc20tokenhome.SendTokensInput{
		DestinationBlockchainID:            fujiCChainID,
		DestinationTokenTransferrerAddress: dep.TokenRemote,
		Recipient:                          to,
		PrimaryFeeTokenAddress:             dep.Token,
		PrimaryFee:                         new(big.Int),
		SecondaryFee:                       new(big.Int),
		RequiredGasLimit:                   big.NewInt(sendGasLimit),
	}, units)
	if _, err := sepolia.mined("send to Fuji", tx, err); err != nil {
		return err
	}
	fmt.Printf("\nsent %s tokens to %s on Fuji. Deliver them with:\n  sepolia-fuji-ictt relay -tx %s\n",
		*amount, to, tx.Hash())
	return nil
}

// balance prints a Fuji account's balance of the remote token, and what the home holds for Fuji.
func balance(ctx context.Context, n networks, args []string) error {
	fs := flag.NewFlagSet("balance", flag.ExitOnError)
	holder := fs.String("account", "", "account on Fuji (default: the KEY account)")
	d := deploymentFlag(fs)
	if err := fs.Parse(args); err != nil {
		return err
	}
	dep, err := d.load()
	if err != nil {
		return err
	}
	var who common.Address
	if *holder != "" {
		who = common.HexToAddress(*holder)
	} else {
		acc, err := accountFromEnv()
		if err != nil {
			return err
		}
		who = acc.address
	}
	fuji, err := dial(ctx, "Fuji", n.fujiRPC, nil)
	if err != nil {
		return err
	}
	remote, err := erc20tokenremote.NewERC20TokenRemote(dep.TokenRemote, fuji.client)
	if err != nil {
		return err
	}
	b, err := remote.BalanceOf(&bind.CallOpts{Context: ctx}, who)
	if err != nil {
		return err
	}
	fmt.Printf("%s holds %s of the remote token on Fuji (raw units, %d decimals)\n", who, b, dep.TokenDecimals)
	sepolia, err := dial(ctx, "Sepolia", n.sepoliaRPC, nil)
	if err != nil {
		return err
	}
	home, err := erc20tokenhome.NewERC20TokenHome(dep.TokenHome, sepolia.client)
	if err != nil {
		return err
	}
	sent, err := home.GetTransferredBalance(&bind.CallOpts{Context: ctx}, fujiCChainID, dep.TokenRemote)
	if err != nil {
		return err
	}
	fmt.Printf("the home on Sepolia holds %s for the remote on Fuji\n", sent)
	return nil
}

// tokenUnits converts a decimal amount of whole tokens to the token's smallest units.
func tokenUnits(amount string, decimals uint8) (*big.Int, error) {
	value, ok := new(big.Float).SetPrec(256).SetString(amount)
	if !ok || value.Sign() <= 0 {
		return nil, fmt.Errorf("invalid amount %q", amount)
	}
	scale := new(big.Float).SetInt(new(big.Int).Exp(big.NewInt(10), big.NewInt(int64(decimals)), nil))
	units, _ := value.Mul(value, scale).Int(nil)
	return units, nil
}
