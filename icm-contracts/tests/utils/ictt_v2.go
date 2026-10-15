// Copyright (C) 2026, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package utils

import (
	"context"
	"crypto/ecdsa"

	"github.com/ava-labs/avalanchego/ids"
	erc20tokenhomev2 "github.com/ava-labs/icm-services/abi-bindings/go/ictt-v2/TokenHome/ERC20TokenHome"
	erc20tokenremotev2 "github.com/ava-labs/icm-services/abi-bindings/go/ictt-v2/TokenRemote/ERC20TokenRemote"
	testinfo "github.com/ava-labs/icm-services/icm-contracts/tests/test-info"
	"github.com/ava-labs/libevm/accounts/abi/bind"
	"github.com/ava-labs/libevm/common"
	"github.com/ava-labs/libevm/core/types"
	. "github.com/onsi/gomega"
)

// RoutedApp is the part of an ICTT v2 contract that configures its routes. The Home and Remote
// bindings both implement it.
type RoutedApp interface {
	AddMessenger(
		opts *bind.TransactOpts,
		remoteBlockchainID [32]byte,
		messenger common.Address,
	) (*types.Transaction, error)
	SetOutboundMessenger(
		opts *bind.TransactOpts,
		remoteBlockchainID [32]byte,
		messenger common.Address,
	) (*types.Transaction, error)
}

// DeployERC20TokenHomeV2 deploys an ICTT v2 ERC20TokenHome on [chain], owned by [teleporterManager].
// The home has no routes until [AddRouteV2] is called for each counterparty chain.
func DeployERC20TokenHomeV2(
	ctx context.Context,
	senderKey *ecdsa.PrivateKey,
	chain testinfo.NetworkTestInfo,
	teleporterManager common.Address,
	tokenAddress common.Address,
	tokenDecimals uint8,
) (common.Address, *erc20tokenhomev2.ERC20TokenHome) {
	evmInfo := chain.GetEVMTestInfo()
	opts, err := bind.NewKeyedTransactorWithChainID(senderKey, evmInfo.EVMChainID)
	Expect(err).Should(BeNil())
	address, tx, home, err := erc20tokenhomev2.DeployERC20TokenHome(
		opts,
		evmInfo.EthClient,
		chain.ChainID(),
		teleporterManager,
		tokenAddress,
		tokenDecimals,
	)
	Expect(err).Should(BeNil())
	WaitForTransactionSuccess(ctx, evmInfo.EthClient, tx.Hash())
	return address, home
}

// DeployERC20TokenRemoteV2 deploys an ICTT v2 ERC20TokenRemote on [chain]. Its initialization routes
// the home's chain through [messenger].
func DeployERC20TokenRemoteV2(
	ctx context.Context,
	senderKey *ecdsa.PrivateKey,
	chain testinfo.NetworkTestInfo,
	messenger common.Address,
	teleporterManager common.Address,
	tokenHomeBlockchainID ids.ID,
	tokenHomeAddress common.Address,
	tokenHomeDecimals uint8,
	tokenName string,
	tokenSymbol string,
	tokenDecimals uint8,
) (common.Address, *erc20tokenremotev2.ERC20TokenRemote) {
	evmInfo := chain.GetEVMTestInfo()
	opts, err := bind.NewKeyedTransactorWithChainID(senderKey, evmInfo.EVMChainID)
	Expect(err).Should(BeNil())
	address, tx, remote, err := erc20tokenremotev2.DeployERC20TokenRemote(
		opts,
		evmInfo.EthClient,
		erc20tokenremotev2.TokenRemoteSettingsV2{
			BlockchainID:          chain.ChainID(),
			TeleporterMessenger:   messenger,
			TeleporterManager:     teleporterManager,
			TokenHomeBlockchainID: tokenHomeBlockchainID,
			TokenHomeAddress:      tokenHomeAddress,
			TokenHomeDecimals:     tokenHomeDecimals,
		},
		tokenName,
		tokenSymbol,
		tokenDecimals,
	)
	Expect(err).Should(BeNil())
	WaitForTransactionSuccess(ctx, evmInfo.EthClient, tx.Hash())
	return address, remote
}

// AddRouteV2 routes [remoteBlockchainID] through [messenger] on [app], for both sending and receiving.
// [ownerKey] must own [app].
func AddRouteV2(
	ctx context.Context,
	ownerKey *ecdsa.PrivateKey,
	chain testinfo.NetworkTestInfo,
	app RoutedApp,
	remoteBlockchainID ids.ID,
	messenger common.Address,
) {
	evmInfo := chain.GetEVMTestInfo()
	opts, err := bind.NewKeyedTransactorWithChainID(ownerKey, evmInfo.EVMChainID)
	Expect(err).Should(BeNil())
	tx, err := app.AddMessenger(opts, remoteBlockchainID, messenger)
	Expect(err).Should(BeNil())
	WaitForTransactionSuccess(ctx, evmInfo.EthClient, tx.Hash())
	tx, err = app.SetOutboundMessenger(opts, remoteBlockchainID, messenger)
	Expect(err).Should(BeNil())
	WaitForTransactionSuccess(ctx, evmInfo.EthClient, tx.Hash())
}
