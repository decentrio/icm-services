// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package erc20tokenremoteupgradeable

import (
	"errors"
	"math/big"
	"strings"

	ethereum "github.com/ava-labs/libevm"
	"github.com/ava-labs/libevm/accounts/abi"
	"github.com/ava-labs/libevm/accounts/abi/bind"
	"github.com/ava-labs/libevm/common"
	"github.com/ava-labs/libevm/core/types"
	"github.com/ava-labs/libevm/event"
)

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = errors.New
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
	_ = abi.ConvertType
)

// SendAndCallInput is an auto generated low-level Go binding around an user-defined struct.
type SendAndCallInput struct {
	DestinationBlockchainID            [32]byte
	DestinationTokenTransferrerAddress common.Address
	RecipientContract                  common.Address
	RecipientPayload                   []byte
	RequiredGasLimit                   *big.Int
	RecipientGasLimit                  *big.Int
	MultiHopFallback                   common.Address
	FallbackRecipient                  common.Address
	PrimaryFeeTokenAddress             common.Address
	PrimaryFee                         *big.Int
	SecondaryFee                       *big.Int
}

// SendTokensInput is an auto generated low-level Go binding around an user-defined struct.
type SendTokensInput struct {
	DestinationBlockchainID            [32]byte
	DestinationTokenTransferrerAddress common.Address
	Recipient                          common.Address
	PrimaryFeeTokenAddress             common.Address
	PrimaryFee                         *big.Int
	SecondaryFee                       *big.Int
	RequiredGasLimit                   *big.Int
	MultiHopFallback                   common.Address
}

// TeleporterFeeInfo is an auto generated low-level Go binding around an user-defined struct.
type TeleporterFeeInfo struct {
	FeeTokenAddress common.Address
	Amount          *big.Int
}

// TokenRemoteSettings is an auto generated low-level Go binding around an user-defined struct.
type TokenRemoteSettings struct {
	BlockchainID          [32]byte
	TeleporterMessenger   common.Address
	TeleporterManager     common.Address
	TokenHomeBlockchainID [32]byte
	TokenHomeAddress      common.Address
	TokenHomeDecimals     uint8
}

// ERC20TokenRemoteUpgradeableMetaData contains all meta data concerning the ERC20TokenRemoteUpgradeable contract.
var ERC20TokenRemoteUpgradeableMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"enumICMInitializable\",\"name\":\"init\",\"type\":\"uint8\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"allowance\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"needed\",\"type\":\"uint256\"}],\"name\":\"ERC20InsufficientAllowance\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"balance\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"needed\",\"type\":\"uint256\"}],\"name\":\"ERC20InsufficientBalance\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"approver\",\"type\":\"address\"}],\"name\":\"ERC20InvalidApprover\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"}],\"name\":\"ERC20InvalidReceiver\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"ERC20InvalidSender\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"}],\"name\":\"ERC20InvalidSpender\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidInitialization\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotInitializing\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"OwnableInvalidOwner\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"OwnableUnauthorizedAccount\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ReentrancyGuardReentrantCall\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"SafeERC20FailedOperation\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"recipientContract\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"CallFailed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"recipientContract\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"CallSucceeded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"version\",\"type\":\"uint64\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"remoteBlockchainID\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"messenger\",\"type\":\"address\"}],\"name\":\"MessengerAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"messenger\",\"type\":\"address\"}],\"name\":\"MessengerPaused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"remoteBlockchainID\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"messenger\",\"type\":\"address\"}],\"name\":\"MessengerRemoved\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"messenger\",\"type\":\"address\"}],\"name\":\"MessengerUnpaused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"remoteBlockchainID\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"oldMessenger\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newMessenger\",\"type\":\"address\"}],\"name\":\"OutboundMessengerSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"teleporterMessageID\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"destinationBlockchainID\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"destinationTokenTransferrerAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"recipientContract\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"recipientPayload\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"requiredGasLimit\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"recipientGasLimit\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"multiHopFallback\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"fallbackRecipient\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"primaryFeeTokenAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"primaryFee\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"secondaryFee\",\"type\":\"uint256\"}],\"indexed\":false,\"internalType\":\"structSendAndCallInput\",\"name\":\"input\",\"type\":\"tuple\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"TokensAndCallSent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"teleporterMessageID\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"destinationBlockchainID\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"destinationTokenTransferrerAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"primaryFeeTokenAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"primaryFee\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"secondaryFee\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"requiredGasLimit\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"multiHopFallback\",\"type\":\"address\"}],\"indexed\":false,\"internalType\":\"structSendTokensInput\",\"name\":\"input\",\"type\":\"tuple\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"TokensSent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"TokensWithdrawn\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"ERC20_TOKEN_REMOTE_STORAGE_LOCATION\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"MULTI_HOP_CALL_GAS_PER_WORD\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"MULTI_HOP_CALL_REQUIRED_GAS\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"MULTI_HOP_SEND_REQUIRED_GAS\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"REGISTER_REMOTE_REQUIRED_GAS\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"TELEPORTER_ROUTED_APP_STORAGE_LOCATION\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"TOKEN_REMOTE_STORAGE_LOCATION\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"remoteBlockchainID\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"messenger\",\"type\":\"address\"}],\"name\":\"addMessenger\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"payloadSize\",\"type\":\"uint256\"}],\"name\":\"calculateNumWords\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"decimals\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getBlockchainID\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getInitialReserveImbalance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getIsCollateralized\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getMultiplyOnRemote\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"remoteBlockchainID\",\"type\":\"bytes32\"}],\"name\":\"getOutboundMessenger\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getTokenHomeAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getTokenHomeBlockchainID\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getTokenMultiplier\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"blockchainID\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"teleporterMessenger\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"teleporterManager\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"tokenHomeBlockchainID\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"tokenHomeAddress\",\"type\":\"address\"},{\"internalType\":\"uint8\",\"name\":\"tokenHomeDecimals\",\"type\":\"uint8\"}],\"internalType\":\"structTokenRemoteSettings\",\"name\":\"settings\",\"type\":\"tuple\"},{\"internalType\":\"string\",\"name\":\"tokenName\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"tokenSymbol\",\"type\":\"string\"},{\"internalType\":\"uint8\",\"name\":\"tokenDecimals\",\"type\":\"uint8\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"remoteBlockchainID\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"messenger\",\"type\":\"address\"}],\"name\":\"isInboundMessenger\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"messenger\",\"type\":\"address\"}],\"name\":\"isMessengerPaused\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"messenger\",\"type\":\"address\"}],\"name\":\"pauseMessenger\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"sourceBlockchainID\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"originSenderAddress\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"message\",\"type\":\"bytes\"}],\"name\":\"receiveTeleporterMessage\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"feeTokenAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"internalType\":\"structTeleporterFeeInfo\",\"name\":\"feeInfo\",\"type\":\"tuple\"}],\"name\":\"registerWithHome\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"remoteBlockchainID\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"messenger\",\"type\":\"address\"}],\"name\":\"removeMessenger\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"destinationBlockchainID\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"destinationTokenTransferrerAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"primaryFeeTokenAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"primaryFee\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"secondaryFee\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"requiredGasLimit\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"multiHopFallback\",\"type\":\"address\"}],\"internalType\":\"structSendTokensInput\",\"name\":\"input\",\"type\":\"tuple\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"send\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"destinationBlockchainID\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"destinationTokenTransferrerAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"recipientContract\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"recipientPayload\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"requiredGasLimit\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"recipientGasLimit\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"multiHopFallback\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"fallbackRecipient\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"primaryFeeTokenAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"primaryFee\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"secondaryFee\",\"type\":\"uint256\"}],\"internalType\":\"structSendAndCallInput\",\"name\":\"input\",\"type\":\"tuple\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"sendAndCall\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"remoteBlockchainID\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"messenger\",\"type\":\"address\"}],\"name\":\"setOutboundMessenger\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"messenger\",\"type\":\"address\"}],\"name\":\"unpauseMessenger\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x608060405234801561000f575f5ffd5b50604051614ba7380380614ba783398101604081905261002e91610107565b60018160018111156100425761004261012c565b0361004f5761004f610055565b50610140565b7ff0c57e16840df040f15088dc2f81fe391c3923bec73e23a9662efc9c229c6a00805468010000000000000000900460ff16156100a55760405163f92ee8a960e01b815260040160405180910390fd5b80546001600160401b03908116146101045780546001600160401b0319166001600160401b0390811782556040519081527fc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d29060200160405180910390a15b50565b5f60208284031215610117575f5ffd5b815160028110610125575f5ffd5b9392505050565b634e487b7160e01b5f52602160045260245ffd5b614a5a8061014d5f395ff3fe608060405234801561000f575f5ffd5b506004361061023f575f3560e01c806383c002d011610135578063d73a0641116100b4578063ef793e2a11610079578063ef793e2a14610639578063f2fde38b14610660578063f3f981d814610673578063f4be94a714610686578063fc56082114610699575f5ffd5b8063d73a0641146105c6578063d90f1b10146105d9578063dd5a20fd146105ec578063dd62ed3e146105ff578063e0fd9cb814610612575f5ffd5b8063b2432529116100fa578063b243252914610549578063b8a46d021461055c578063c1c6353e1461056f578063c3cd692714610583578063c868efaa146105b3575f5ffd5b806383c002d01461045c5780638da5cb5b146104a657806395a19443146104ea57806395d89b411461052e578063a9059cbb14610536575f5ffd5b806335cac159116101c1578063656900381161018657806365690038146103d957806370a08231146103ec578063715018a61461042057806371717c18146104285780637ee3779a14610432575f5ffd5b806335cac1591461036d5780634213cf78146103815780635507f3d1146103955780635d16225d1461039f57806362431a65146103b2575f5ffd5b806315beb59f1161020757806315beb59f146102ec57806318160ddd146102f557806323b872dd1461031c578063254ac1601461032f578063313ce56714610339575f5ffd5b806302a30c7d1461024357806303f789511461027e57806306fdde03146102935780630733c8c8146102a8578063095ea7b3146102d9575b5f5ffd5b7f600d6a9b283d1eda563de594ce4843869b6f128a4baa222422ed94a60b0cef065460ff165b60405190151581526020015b60405180910390f35b61029161028c3660046138ed565b6106e0565b005b61029b6107f3565b6040516102759190613955565b7f600d6a9b283d1eda563de594ce4843869b6f128a4baa222422ed94a60b0cef03545b604051908152602001610275565b6102696102e7366004613967565b6108b3565b6102cb6105dc81565b7f52c63247e1f47db19d5ce0460030c497f067ca4cebf71ba98eeadabe20bace02546102cb565b61026961032a366004613991565b6108cc565b6102cb6201fbd081565b7f9b9029a3537fcf0e984763da4ac33bbf592a3462819171bf424e91cf626223005460405160ff9091168152602001610275565b6102cb5f5160206149c55f395f51905f5281565b5f5160206149e55f395f51905f52546102cb565b6102cb6205302081565b6102916103ad3660046139cf565b6108ef565b6102cb7f9b9029a3537fcf0e984763da4ac33bbf592a3462819171bf424e91cf6262230081565b6102916103e7366004613a00565b6108fd565b6102cb6103fa3660046138ed565b6001600160a01b03165f9081525f5160206149a55f395f51905f52602052604090205490565b610291610907565b6102cb6205573081565b7f600d6a9b283d1eda563de594ce4843869b6f128a4baa222422ed94a60b0cef045460ff16610269565b61026961046a3660046138ed565b6001600160a01b03165f9081527f63108c0ddc3d3ab0f279f27d4b9ca8220d5efbdbeeb537d4327d5cf9e982b103602052604090205460ff1690565b7f9016d09d72d40fdae2fd8ceac6b6234c7706214fd39c1cd1e609a0528c199300546001600160a01b03165b6040516001600160a01b039091168152602001610275565b6102696104f8366004613a38565b5f9182525f516020614a055f395f51905f52602090815260408084206001600160a01b0393909316845291905290205460ff1690565b61029b61091a565b610269610544366004613967565b610958565b610291610557366004613a38565b610965565b61029161056a366004613a66565b610977565b6102cb5f5160206149e55f395f51905f5281565b7f600d6a9b283d1eda563de594ce4843869b6f128a4baa222422ed94a60b0cef02546001600160a01b03166104d2565b6102916105c1366004613a7f565b610b5e565b6102916105d4366004613c3b565b610c7d565b6102916105e7366004613a38565b610d8f565b6102916105fa3660046138ed565b610da1565b6102cb61060d366004613d26565b610ea3565b7f600d6a9b283d1eda563de594ce4843869b6f128a4baa222422ed94a60b0cef01546102cb565b7f600d6a9b283d1eda563de594ce4843869b6f128a4baa222422ed94a60b0cef05546102cb565b61029161066e3660046138ed565b610eec565b6102cb610681366004613d52565b610f29565b610291610694366004613a38565b610f3f565b6104d26106a7366004613d52565b5f9081527f63108c0ddc3d3ab0f279f27d4b9ca8220d5efbdbeeb537d4327d5cf9e982b10160205260409020546001600160a01b031690565b6106e8611085565b5f5160206149e55f395f51905f526001600160a01b0382166107255760405162461bcd60e51b815260040161071c90613d69565b60405180910390fd5b6001600160a01b0382165f90815260038201602052604090205460ff16156107a55760405162461bcd60e51b815260206004820152602d60248201527f54656c65706f72746572526f757465644170703a206d657373656e676572206160448201526c1b1c9958591e481c185d5cd959609a1b606482015260840161071c565b6001600160a01b0382165f818152600383016020526040808220805460ff19166001179055517f46def370ef3a1ec72618480f66f1063a80fe68b27d09575bc3562ad489079b619190a25050565b7f52c63247e1f47db19d5ce0460030c497f067ca4cebf71ba98eeadabe20bace0380546060915f5160206149a55f395f51905f529161083190613db4565b80601f016020809104026020016040519081016040528092919081815260200182805461085d90613db4565b80156108a85780601f1061087f576101008083540402835291602001916108a8565b820191905f5260205f20905b81548152906001019060200180831161088b57829003601f168201915b505050505091505090565b5f336108c081858561108d565b60019150505b92915050565b5f336108d985828561109f565b6108e48585856110fc565b506001949350505050565b6108f98282611159565b5050565b6108f982826111e4565b61090f61126f565b6109185f6112ca565b565b7f52c63247e1f47db19d5ce0460030c497f067ca4cebf71ba98eeadabe20bace0480546060915f5160206149a55f395f51905f529161083190613db4565b5f336108c08185856110fc565b61096d611085565b6108f9828261133a565b7f600d6a9b283d1eda563de594ce4843869b6f128a4baa222422ed94a60b0cef06545f5160206149c55f395f51905f5290610100900460ff16156109fd5760405162461bcd60e51b815260206004820152601f60248201527f546f6b656e52656d6f74653a20616c7265616479207265676973746572656400604482015260640161071c565b604080516060808201835260058401548252600284015460ff600160a01b820481166020808601918252600160a81b9093048216858701908152865180880188525f808252885188518188015293518516848a01529151909316828601528651808303909501855260809091019095528082019290925291929091610a9290610a88908701876138ed565b86602001356115ff565b6040805160c0810182526001870154815260028701546001600160a01b031660208083019190915282518084018452939450610b56939192830191908190610adc908b018b6138ed565b6001600160a01b0316815260209081018690529082526201fbd0908201526040015f5b604051908082528060200260200182016040528015610b28578160200160208202803683370190505b50815260200184604051602001610b3f9190613e00565b60405160208183030381529060405281525061164e565b505050505050565b610b6661176b565b5f8481525f516020614a055f395f51905f52602090815260408083203384529091529020545f5160206149e55f395f51905f529060ff16610bfd5760405162461bcd60e51b815260206004820152602b60248201527f54656c65706f72746572526f757465644170703a20756e617574686f72697a6560448201526a321036b2b9b9b2b733b2b960a91b606482015260840161071c565b335f90815260038201602052604090205460ff1615610c2e5760405162461bcd60e51b815260040161071c90613e42565b610c6e858585858080601f0160208091040260200160405190810160405280939291908181526020018383808284375f920191909152506117b592505050565b50610c776119ed565b50505050565b7ff0c57e16840df040f15088dc2f81fe391c3923bec73e23a9662efc9c229c6a008054600160401b810460ff1615906001600160401b03165f81158015610cc15750825b90505f826001600160401b03166001148015610cdc5750303b155b905081158015610cea575080155b15610d085760405163f92ee8a960e01b815260040160405180910390fd5b845467ffffffffffffffff191660011785558315610d3257845460ff60401b1916600160401b1785555b610d3e89898989611a17565b8315610d8457845460ff60401b19168555604051600181527fc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d29060200160405180910390a15b505050505050505050565b610d97611085565b6108f98282611a67565b610da9611085565b5f5160206149e55f395f51905f526001600160a01b038216610ddd5760405162461bcd60e51b815260040161071c90613d69565b6001600160a01b0382165f90815260038201602052604090205460ff16610e585760405162461bcd60e51b815260206004820152602960248201527f54656c65706f72746572526f757465644170703a206d657373656e676572206e6044820152681bdd081c185d5cd95960ba1b606482015260840161071c565b6001600160a01b0382165f818152600383016020526040808220805460ff19169055517ff5574433a9d82924af98e239704e51afb34ae8c15f513d78a22ae0c58f5f99659190a25050565b6001600160a01b039182165f9081527f52c63247e1f47db19d5ce0460030c497f067ca4cebf71ba98eeadabe20bace016020908152604080832093909416825291909152205490565b610ef461126f565b6001600160a01b038116610f1d57604051631e4fbdf760e01b81525f600482015260240161071c565b610f26816112ca565b50565b5f6005610f3783601f613e9b565b901c92915050565b610f47611085565b5f8281525f516020614a055f395f51905f52602090815260408083206001600160a01b03851684529091529020545f5160206149e55f395f51905f529060ff16610fa35760405162461bcd60e51b815260040161071c90613eae565b5f8381526001820160205260409020546001600160a01b0380841691160361102b5760405162461bcd60e51b815260206004820152603560248201527f54656c65706f72746572526f757465644170703a2063616e6e6f742072656d6f6044820152743b329037baba3137bab7321036b2b9b9b2b733b2b960591b606482015260840161071c565b5f83815260028201602090815260408083206001600160a01b0386168085529252808320805460ff1916905551909185917f8350b16da8b8ec7d7e6ce37ca2ccb4f63d628b5ad53215d072763b6afce995249190a3505050565b61091861126f565b61109a8383836001611ba7565b505050565b5f6110aa8484610ea3565b90505f198114610c7757818110156110ee57604051637dc7a0d960e11b81526001600160a01b0384166004820152602481018290526044810183905260640161071c565b610c7784848484035f611ba7565b6001600160a01b03831661112557604051634b637e8f60e11b81525f600482015260240161071c565b6001600160a01b03821661114e5760405163ec442f0560e01b81525f600482015260240161071c565b61109a838383611c8a565b7fd2f1ed38b7d242bfb8b41862afb813a15193219a4bc717f2056607593e6c7500805460011461119b5760405162461bcd60e51b815260040161071c90613ef6565b600281555f5160206149c55f395f51905f526111b684611dc3565b60018101548435036111d1576111cc8484611eae565b6111db565b6111db8484612035565b50600190555050565b7fd2f1ed38b7d242bfb8b41862afb813a15193219a4bc717f2056607593e6c750080546001146112265760405162461bcd60e51b815260040161071c90613ef6565b600281555f5160206149c55f395f51905f5261124184612203565b600181015484350361125d57611257848461243d565b50611267565b6112578484612651565b600190555050565b336112a17f9016d09d72d40fdae2fd8ceac6b6234c7706214fd39c1cd1e609a0528c199300546001600160a01b031690565b6001600160a01b0316146109185760405163118cdaa760e01b815233600482015260240161071c565b7f9016d09d72d40fdae2fd8ceac6b6234c7706214fd39c1cd1e609a0528c19930080546001600160a01b031981166001600160a01b03848116918217845560405192169182907f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0905f90a3505050565b5f5160206149e55f395f51905f52826113ac5760405162461bcd60e51b815260206004820152602e60248201527f54656c65706f72746572526f757465644170703a207a65726f2072656d6f746560448201526d08189b1bd8dad8da185a5b88125160921b606482015260840161071c565b805483036114185760405162461bcd60e51b815260206004820152603360248201527f54656c65706f72746572526f757465644170703a2063616e6e6f7420726f757460448201527232903a379037bbb710313637b1b5b1b430b4b760691b606482015260840161071c565b6001600160a01b03821661143e5760405162461bcd60e51b815260040161071c90613d69565b5f83815260028201602090815260408083206001600160a01b038616845290915290205460ff16156114c75760405162461bcd60e51b815260206004820152602c60248201527f54656c65706f72746572526f757465644170703a206d657373656e676572206160448201526b1b1c9958591e48185919195960a21b606482015260840161071c565b5f826001600160a01b031663d127dc9b6040518163ffffffff1660e01b8152600401602060405180830381865afa158015611504573d5f5f3e3d5ffd5b505050506040513d601f19601f820116820180604052508101906115289190613f3a565b90508015806115375750815481145b6115a15760405162461bcd60e51b815260206004820152603560248201527f54656c65706f72746572526f757465644170703a206d657373656e67657220626044820152740d8dec6d6c6d0c2d2dc40928840dad2e6dac2e8c6d605b1b606482015260840161071c565b5f84815260028301602090815260408083206001600160a01b0387168085529252808320805460ff1916600117905551909186917fcc2f89e16e8979d86d3c3b49c5646f3a013f76626a479400e1cb3596f3dc2b229190a350505050565b5f815f0361160e57505f6108c6565b306001600160a01b0384160361163c5761162a335b308461109f565b6116353330846110fc565b50806108c6565b6116478333846128f0565b9392505050565b5f5f61165c835f0151612a53565b604084015160200151909150156116ff576040830151516001600160a01b03166116dc5760405162461bcd60e51b815260206004820152602b60248201527f54656c65706f72746572526f757465644170703a207a65726f2066656520746f60448201526a6b656e206164647265737360a81b606482015260840161071c565b6040830151602081015190516116ff916001600160a01b03909116908390612b32565b604051630624488560e41b81526001600160a01b0382169063624488509061172b908690600401613f51565b6020604051808303815f875af1158015611747573d5f5f3e3d5ffd5b505050506040513d601f19601f820116820180604052508101906116479190613f3a565b7f9b779b17422d0df92223018b32b4d1fa46e071723d6817e2486d003becc55f008054600119016117af57604051633ee5aeb560e01b815260040160405180910390fd5b60029055565b7f600d6a9b283d1eda563de594ce4843869b6f128a4baa222422ed94a60b0cef01545f5160206149c55f395f51905f529084146118465760405162461bcd60e51b815260206004820152602960248201527f546f6b656e52656d6f74653a20696e76616c696420736f7572636520626c6f636044820152681ad8da185a5b88125160ba1b606482015260840161071c565b60028101546001600160a01b038481169116146118b85760405162461bcd60e51b815260206004820152602a60248201527f546f6b656e52656d6f74653a20696e76616c6964206f726967696e2073656e646044820152696572206164647265737360b01b606482015260840161071c565b5f828060200190518101906118cd919061404e565b6006830154909150610100900460ff1615806118ee5750600682015460ff16155b156119055760068201805461ffff19166101011790555b60018151600481111561191a5761191a613dec565b03611951575f816020015180602001905181019061193891906140d9565b905061194b815f01518260200151612bb9565b506119e6565b60028151600481111561196657611966613dec565b03611994575f81602001518060200190518101906119849190614113565b905061194b818260800151612c06565b60405162461bcd60e51b815260206004820152602160248201527f546f6b656e52656d6f74653a20696e76616c6964206d657373616765207479706044820152606560f81b606482015260840161071c565b5050505050565b5f7f9b779b17422d0df92223018b32b4d1fa46e071723d6817e2486d003becc55f005b6001905550565b611a1f612d5f565b611a298383612da8565b611a34845f83612dba565b7f9b9029a3537fcf0e984763da4ac33bbf592a3462819171bf424e91cf62622300805460ff191660ff8316179055610c77565b5f8281525f516020614a055f395f51905f52602090815260408083206001600160a01b03851684529091529020545f5160206149e55f395f51905f529060ff16611ac35760405162461bcd60e51b815260040161071c90613eae565b5f8381526001820160205260409020546001600160a01b039081169083168103611b475760405162461bcd60e51b815260206004820152602f60248201527f54656c65706f72746572526f757465644170703a206d657373656e676572206160448201526e1b1c9958591e481bdd5d189bdd5b99608a1b606482015260840161071c565b5f84815260018301602052604080822080546001600160a01b0319166001600160a01b03878116918217909255915191929084169187917fd00498d0475cde985344ea08d024c9315ca654d9d157211295aac4349637411591a450505050565b5f5160206149a55f395f51905f526001600160a01b038516611bde5760405163e602df0560e01b81525f600482015260240161071c565b6001600160a01b038416611c0757604051634a1406b160e11b81525f600482015260240161071c565b6001600160a01b038086165f908152600183016020908152604080832093881683529290522083905581156119e657836001600160a01b0316856001600160a01b03167f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b92585604051611c7b91815260200190565b60405180910390a35050505050565b5f5160206149a55f395f51905f526001600160a01b038416611cc45781816002015f828254611cb99190613e9b565b90915550611d349050565b6001600160a01b0384165f9081526020829052604090205482811015611d165760405163391434e360e21b81526001600160a01b0386166004820152602481018290526044810184905260640161071c565b6001600160a01b0385165f9081526020839052604090209083900390555b6001600160a01b038316611d52576002810180548390039055611d70565b6001600160a01b0383165f9081526020829052604090208054830190555b826001600160a01b0316846001600160a01b03167fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef84604051611db591815260200190565b60405180910390a350505050565b5f611dd460608301604084016138ed565b6001600160a01b031603611e365760405162461bcd60e51b815260206004820152602360248201527f546f6b656e52656d6f74653a207a65726f20726563697069656e74206164647260448201526265737360e81b606482015260840161071c565b5f8160c0013511611e595760405162461bcd60e51b815260040161071c906141e0565b8035611e775760405162461bcd60e51b815260040161071c90614224565b5f611e8860408301602084016138ed565b6001600160a01b031603610f265760405162461bcd60e51b815260040161071c9061426f565b5f5160206149c55f395f51905f52611eea611ecf60408501602086016138ed565b60a0850135611ee5610100870160e088016138ed565b612de6565b5f611f0e83611eff60808701606088016138ed565b86608001358760a00135612f03565b6040805180820190915291945091505f9080600181526020016040518060400160405280886040016020810190611f4591906138ed565b6001600160a01b0316815260200187815250604051602001611f6791906142cc565b60408051601f198184030181529181529152805160c0810182526001860154815260028601546001600160a01b03166020820152815180830183529293505f92611fe79282019080611fbf60808c0160608d016138ed565b6001600160a01b03168152602090810188905290825260c08a0135908201526040015f610aff565b9050336001600160a01b0316817f93f19bf1ec58a15dc643b37e7e18a1c13e85e06cd11929e283154691ace9fb5288886040516120259291906142ec565b60405180910390a3505050505050565b5f5160206149c55f395f51905f5261206e833561205860408601602087016138ed565b612069610100870160e088016138ed565b612fc9565b5f61208383611eff60808701606088016138ed565b60408051808201825260038152815160e081018352883581529396509193505f92602080840192828201916120bc918b01908b016138ed565b6001600160a01b031681526020016120da60608a0160408b016138ed565b6001600160a01b031681526020810188905260a0890135604082015260c089013560608201526080016121146101008a0160e08b016138ed565b6001600160a01b031690526040516121849190602001815181526020808301516001600160a01b0390811691830191909152604080840151821690830152606080840151908301526080808401519083015260a0808401519083015260c092830151169181019190915260e00190565b60408051601f198184030181529181529152805160c0810182526001860154815260028601546001600160a01b03166020820152815180830183529293505f92611fe792820190806121dc60808c0160608d016138ed565b6001600160a01b03168152602090810188905290825262053020908201526040015f610aff565b80356122215760405162461bcd60e51b815260040161071c90614224565b5f61223260408301602084016138ed565b6001600160a01b0316036122585760405162461bcd60e51b815260040161071c9061426f565b5f61226960608301604084016138ed565b6001600160a01b0316036122d45760405162461bcd60e51b815260206004820152602c60248201527f546f6b656e52656d6f74653a207a65726f20726563697069656e7420636f6e7460448201526b72616374206164647265737360a01b606482015260840161071c565b5f8160800135116122f75760405162461bcd60e51b815260040161071c906141e0565b5f8160a00135116123585760405162461bcd60e51b815260206004820152602560248201527f546f6b656e52656d6f74653a207a65726f20726563697069656e7420676173206044820152641b1a5b5a5d60da1b606482015260840161071c565b80608001358160a00135106123c05760405162461bcd60e51b815260206004820152602860248201527f546f6b656e52656d6f74653a20696e76616c696420726563697069656e742067604482015267185cc81b1a5b5a5d60c21b606482015260840161071c565b5f6123d2610100830160e084016138ed565b6001600160a01b031603610f265760405162461bcd60e51b815260206004820152602c60248201527f546f6b656e52656d6f74653a207a65726f2066616c6c6261636b20726563697060448201526b69656e74206164647265737360a01b606482015260840161071c565b5f5160206149c55f395f51905f5261247461245e60408501602086016138ed565b610140850135611ee560e0870160c088016138ed565b5f61249c8361248b610120870161010088016138ed565b866101200135876101400135612f03565b6040805180820190915291945091505f9080600281526020016040518061010001604052806124d65f5160206149e55f395f51905f525490565b8152306020820152604001336001600160a01b0316815260200161250060608a0160408b016138ed565b6001600160a01b031681526020810188905260400161252260608a018a61438e565b8080601f0160208091040260200160405190810160405280939291908181526020018383808284375f9201919091525050509082525060a089013560208201526040016125766101008a0160e08b016138ed565b6001600160a01b0316905260405161259191906020016143d0565b60408051601f198184030181529181529152805160c0810182526001860154815260028601546001600160a01b03166020820152815180830183529293505f9261261392820190806125eb6101208c016101008d016138ed565b6001600160a01b03168152602090810188905290825260808a0135908201526040015f610aff565b9050336001600160a01b0316817f5d76dff81bf773b908b050fa113d39f7d8135bb4175398f313ea19cd3a1a0b1688886040516120259291906144d8565b5f5160206149c55f395f51905f52612684833561267460408601602087016138ed565b61206960e0870160c088016138ed565b5f61269b8361248b610120870161010088016138ed565b6040805180820190915291945091505f9080600481526020016040518061016001604052806126c73390565b6001600160a01b03168152602001885f013581526020018860200160208101906126f191906138ed565b6001600160a01b0316815260200161270f60608a0160408b016138ed565b6001600160a01b031681526020810188905260400161273160608a018a61438e565b8080601f0160208091040260200160405190810160405280939291908181526020018383808284375f9201919091525050509082525060a089013560208201526040016127856101008a0160e08b016138ed565b6001600160a01b03168152608089013560208201526040016127ad60e08a0160c08b016138ed565b6001600160a01b031681526101408901356020918201526040516127d29291016145cc565b60408051601f19818403018152919052905290505f6105dc6128016127fa606089018961438e565b9050610f29565b61280b91906146ad565b6128189062055730613e9b565b6040805160c0810182526001870154815260028701546001600160a01b03166020820152815180830183529293505f926128a192820190806128626101208d016101008e016138ed565b6001600160a01b031681526020908101899052908252818101869052604080515f815280830182528184015251606090920191610b3f91889101613e00565b9050336001600160a01b0316817f5d76dff81bf773b908b050fa113d39f7d8135bb4175398f313ea19cd3a1a0b1689896040516128df9291906144d8565b60405180910390a350505050505050565b6040516370a0823160e01b81523060048201525f9081906001600160a01b038616906370a0823190602401602060405180830381865afa158015612936573d5f5f3e3d5ffd5b505050506040513d601f19601f8201168201806040525081019061295a9190613f3a565b90506129716001600160a01b038616853086613068565b6040516370a0823160e01b81523060048201525f906001600160a01b038716906370a0823190602401602060405180830381865afa1580156129b5573d5f5f3e3d5ffd5b505050506040513d601f19601f820116820180604052508101906129d99190613f3a565b9050818111612a3f5760405162461bcd60e51b815260206004820152602c60248201527f5361666545524332305472616e7366657246726f6d3a2062616c616e6365206e60448201526b1bdd081a5b98dc99585cd95960a21b606482015260840161071c565b612a4982826146c4565b9695505050505050565b5f8181527f63108c0ddc3d3ab0f279f27d4b9ca8220d5efbdbeeb537d4327d5cf9e982b10160205260408120545f5160206149e55f395f51905f52906001600160a01b031680612af85760405162461bcd60e51b815260206004820152602a60248201527f54656c65706f72746572526f757465644170703a206e6f206f7574626f756e646044820152691036b2b9b9b2b733b2b960b11b606482015260840161071c565b6001600160a01b0381165f90815260038301602052604090205460ff16156116475760405162461bcd60e51b815260040161071c90613e42565b604051636eb1769f60e11b81523060048201526001600160a01b0383811660248301525f919085169063dd62ed3e90604401602060405180830381865afa158015612b7f573d5f5f3e3d5ffd5b505050506040513d601f19601f82011682018060405250810190612ba39190613f3a565b9050610c778484612bb48585613e9b565b6130cf565b816001600160a01b03167f6352c5382c4a4578e712449ca65e83cdb392d045dfcf1cad9615189db2da244b82604051612bf491815260200190565b60405180910390a26108f9828261315e565b612c10308261315e565b612c1f3083606001518361108d565b5f825f01518360200151846040015130858760a00151604051602401612c4a969594939291906146d7565b60408051601f198184030181529190526020810180516001600160e01b03166394395edd60e01b17905260c084015160608501519192505f91612c8e919084613192565b90505f612c9f308660600151610ea3565b9050612cb03086606001515f61108d565b8115612d025784606001516001600160a01b03167f104deb555f67e63782bb817bc26c39050894645f9b9f29c4be8ae68d0e8b7ff485604051612cf591815260200190565b60405180910390a2612d4a565b84606001516001600160a01b03167fb9eaeae386d339f8115782f297a9e5f0e13fb587cd6b0d502f113cb8dd4d6cb085604051612d4191815260200190565b60405180910390a25b80156119e6576119e6308660e00151836110fc565b7ff0c57e16840df040f15088dc2f81fe391c3923bec73e23a9662efc9c229c6a0054600160401b900460ff1661091857604051631afcd79f60e31b815260040160405180910390fd5b612db0612d5f565b6108f982826131a7565b612dc2612d5f565b612dd3835f015184604001516131f7565b612ddb613211565b61109a838383613221565b7f600d6a9b283d1eda563de594ce4843869b6f128a4baa222422ed94a60b0cef02545f5160206149c55f395f51905f52906001600160a01b03858116911614612e415760405162461bcd60e51b815260040161071c90614723565b8215612e9b5760405162461bcd60e51b815260206004820152602360248201527f546f6b656e52656d6f74653a206e6f6e2d7a65726f207365636f6e646172792060448201526266656560e81b606482015260840161071c565b6001600160a01b03821615610c775760405162461bcd60e51b815260206004820152602860248201527f546f6b656e52656d6f74653a206e6f6e2d7a65726f206d756c74692d686f702060448201526766616c6c6261636b60c01b606482015260840161071c565b5f805f5160206149c55f395f51905f52612f1c8761350d565b9650612f2886866115ff565b60038201546004830154919650612f429160ff1686613525565b60038201546004830154612f5a919060ff168a613525565b11612fbc5760405162461bcd60e51b815260206004820152602c60248201527f546f6b656e52656d6f74653a20696e73756666696369656e7420746f6b656e7360448201526b103a37903a3930b739b332b960a11b606482015260840161071c565b5094959294509192505050565b5f5160206149e55f395f51905f5254830361300657306001600160a01b038316036130065760405162461bcd60e51b815260040161071c90614723565b6001600160a01b03811661109a5760405162461bcd60e51b8152602060048201526024808201527f546f6b656e52656d6f74653a207a65726f206d756c74692d686f702066616c6c6044820152636261636b60e01b606482015260840161071c565b6040516001600160a01b038481166024830152838116604483015260648201839052610c779186918216906323b872dd906084015b604051602081830303815290604052915060e01b6020820180516001600160e01b038381831617835250505050613532565b604080516001600160a01b038416602482015260448082018490528251808303909101815260649091019091526020810180516001600160e01b031663095ea7b360e01b179052613120848261359e565b610c77576040516001600160a01b0384811660248301525f604483015261315491869182169063095ea7b39060640161309d565b610c778482613532565b6001600160a01b0382166131875760405163ec442f0560e01b81525f600482015260240161071c565b6108f95f8383611c8a565b5f61319f845f85856135e3565b949350505050565b6131af612d5f565b5f5160206149a55f395f51905f527f52c63247e1f47db19d5ce0460030c497f067ca4cebf71ba98eeadabe20bace036131e884826147c4565b5060048101610c7783826147c4565b6131ff612d5f565b61320882613724565b6108f981613745565b613219612d5f565b610918613756565b613229612d5f565b60608301515f5160206149c55f395f51905f529061329c5760405162461bcd60e51b815260206004820152602a60248201527f546f6b656e52656d6f74653a207a65726f20746f6b656e20686f6d6520626c6f60448201526918dad8da185a5b88125160b21b606482015260840161071c565b5f5160206149e55f395f51905f52548460600151036133235760405162461bcd60e51b815260206004820152603b60248201527f546f6b656e52656d6f74653a2063616e6e6f74206465706c6f7920746f20736160448201527f6d6520626c6f636b636861696e20617320746f6b656e20686f6d650000000000606482015260840161071c565b60808401516001600160a01b03166133895760405162461bcd60e51b8152602060048201526024808201527f546f6b656e52656d6f74653a207a65726f20746f6b656e20686f6d65206164646044820152637265737360e01b606482015260840161071c565b60128460a0015160ff1611156133f35760405162461bcd60e51b815260206004820152602960248201527f546f6b656e52656d6f74653a20746f6b656e20686f6d6520646563696d616c73604482015268040e8dede40d0d2ced60bb1b606482015260840161071c565b60128260ff1611156134535760405162461bcd60e51b8152602060048201526024808201527f546f6b656e52656d6f74653a20746f6b656e20646563696d616c7320746f6f206044820152630d0d2ced60e31b606482015260840161071c565b60608401516001820155608084015160028201805460058401869055600684018054871560ff1990911617905560a08701516001600160a01b039093166001600160a81b031990911617600160a01b60ff808516919091029190911760ff60a81b1916600160a81b918616919091021790556134cf908361377d565b60048301805460ff19169115159190911790556003820155606084015160208501516134fb919061133a565b610c7784606001518560200151611a67565b5f61351733611623565b61352133836137c7565b5090565b5f61319f8484845f6137fb565b5f5f60205f8451602086015f885af180613551576040513d5f823e3d81fd5b50505f513d91508115613568578060011415613575565b6001600160a01b0384163b155b15610c7757604051635274afe760e01b81526001600160a01b038516600482015260240161071c565b5f5f5f5f60205f8651602088015f8a5af192503d91505f519050828015612a49575081156135cf5780600114612a49565b50505050506001600160a01b03163b151590565b5f845a10156136345760405162461bcd60e51b815260206004820152601b60248201527f43616c6c5574696c733a20696e73756666696369656e74206761730000000000604482015260640161071c565b834710156136845760405162461bcd60e51b815260206004820152601d60248201527f43616c6c5574696c733a20696e73756666696369656e742076616c7565000000604482015260640161071c565b826001600160a01b03163b5f0361369c57505f61319f565b6127106136aa603f8761487e565b6136b49087613e9b565b6136be9190613e9b565b5a101561370d5760405162461bcd60e51b815260206004820152601b60248201527f43616c6c5574696c733a20696e73756666696369656e74206761730000000000604482015260640161071c565b5f5f5f84516020860188888bf19695505050505050565b61372c612d5f565b61373461382b565b61373c61383b565b610f2681613843565b61374d612d5f565b610f26816138b9565b5f7fd2f1ed38b7d242bfb8b41862afb813a15193219a4bc717f2056607593e6c7500611a10565b5f8060ff8085169084161181816137a057613798858761489d565b60ff166137ae565b6137aa868661489d565b60ff165b6137b990600a614999565b9350909150505b9250929050565b6001600160a01b0382166137f057604051634b637e8f60e11b81525f600482015260240161071c565b6108f9825f83611c8a565b5f811515841515036138185761381185846146ad565b905061319f565b613822858461487e565b95945050505050565b613833612d5f565b6109186138c1565b610918612d5f565b61384b612d5f565b806138a85760405162461bcd60e51b815260206004820152602760248201527f54656c65706f72746572526f757465644170703a207a65726f20626c6f636b636044820152661a185a5b88125160ca1b606482015260840161071c565b5f5160206149e55f395f51905f5255565b610ef4612d5f565b6119ed612d5f565b6001600160a01b0381168114610f26575f5ffd5b80356138e8816138c9565b919050565b5f602082840312156138fd575f5ffd5b8135611647816138c9565b5f5b8381101561392257818101518382015260200161390a565b50505f910152565b5f8151808452613941816020860160208601613908565b601f01601f19169290920160200192915050565b602081525f611647602083018461392a565b5f5f60408385031215613978575f5ffd5b8235613983816138c9565b946020939093013593505050565b5f5f5f606084860312156139a3575f5ffd5b83356139ae816138c9565b925060208401356139be816138c9565b929592945050506040919091013590565b5f5f8284036101208112156139e2575f5ffd5b6101008112156139f0575f5ffd5b5091936101008501359350915050565b5f5f60408385031215613a11575f5ffd5b82356001600160401b03811115613a26575f5ffd5b83016101608186031215613983575f5ffd5b5f5f60408385031215613a49575f5ffd5b823591506020830135613a5b816138c9565b809150509250929050565b5f6040828403128015613a77575f5ffd5b509092915050565b5f5f5f5f60608587031215613a92575f5ffd5b843593506020850135613aa4816138c9565b925060408501356001600160401b03811115613abe575f5ffd5b8501601f81018713613ace575f5ffd5b80356001600160401b03811115613ae3575f5ffd5b876020828401011115613af4575f5ffd5b949793965060200194505050565b634e487b7160e01b5f52604160045260245ffd5b60405160c081016001600160401b0381118282101715613b3857613b38613b02565b60405290565b604080519081016001600160401b0381118282101715613b3857613b38613b02565b60405161010081016001600160401b0381118282101715613b3857613b38613b02565b604051601f8201601f191681016001600160401b0381118282101715613bab57613bab613b02565b604052919050565b803560ff811681146138e8575f5ffd5b5f6001600160401b03821115613bdb57613bdb613b02565b50601f01601f191660200190565b5f82601f830112613bf8575f5ffd5b8135613c0b613c0682613bc3565b613b83565b818152846020838601011115613c1f575f5ffd5b816020850160208301375f918101602001919091529392505050565b5f5f5f5f848603610120811215613c50575f5ffd5b60c0811215613c5d575f5ffd5b50613c66613b16565b853581526020860135613c78816138c9565b60208201526040860135613c8b816138c9565b6040820152606086810135908201526080860135613ca8816138c9565b6080820152613cb960a08701613bb3565b60a0820152935060c08501356001600160401b03811115613cd8575f5ffd5b613ce487828801613be9565b93505060e08501356001600160401b03811115613cff575f5ffd5b613d0b87828801613be9565b925050613d1b6101008601613bb3565b905092959194509250565b5f5f60408385031215613d37575f5ffd5b8235613d42816138c9565b91506020830135613a5b816138c9565b5f60208284031215613d62575f5ffd5b5035919050565b6020808252602b908201527f54656c65706f72746572526f757465644170703a207a65726f206d657373656e60408201526a676572206164647265737360a81b606082015260800190565b600181811c90821680613dc857607f821691505b602082108103613de657634e487b7160e01b5f52602260045260245ffd5b50919050565b634e487b7160e01b5f52602160045260245ffd5b602081525f825160058110613e2357634e487b7160e01b5f52602160045260245ffd5b80602084015250602083015160408084015261319f606084018261392a565b60208082526025908201527f54656c65706f72746572526f757465644170703a206d657373656e6765722070604082015264185d5cd95960da1b606082015260800190565b634e487b7160e01b5f52601160045260245ffd5b808201808211156108c6576108c6613e87565b60208082526028908201527f54656c65706f72746572526f757465644170703a206d657373656e676572206e6040820152671bdd08185919195960c21b606082015260800190565b60208082526024908201527f53656e645265656e7472616e637947756172643a2073656e64207265656e7472604082015263616e637960e01b606082015260800190565b5f60208284031215613f4a575f5ffd5b5051919050565b602081525f61010082018351602084015260018060a01b0360208501511660408401526040840151613f99606085018280516001600160a01b03168252602090810151910152565b50606084015160a0840152608084015160e060c0850152805191829052602001905f906101208501905b80831015613fee5783516001600160a01b031682526020938401936001939093019290910190613fc3565b5060a0860151858203601f190160e08701529250612a49818461392a565b5f82601f83011261401b575f5ffd5b8151614029613c0682613bc3565b81815284602083860101111561403d575f5ffd5b61319f826020830160208701613908565b5f6020828403121561405e575f5ffd5b81516001600160401b03811115614073575f5ffd5b820160408185031215614084575f5ffd5b61408c613b3e565b81516005811061409a575f5ffd5b815260208201516001600160401b038111156140b4575f5ffd5b6140c08682850161400c565b602083015250949350505050565b80516138e8816138c9565b5f60408284031280156140ea575f5ffd5b506140f3613b3e565b82516140fe816138c9565b81526020928301519281019290925250919050565b5f60208284031215614123575f5ffd5b81516001600160401b03811115614138575f5ffd5b8201610100818503121561414a575f5ffd5b614152613b60565b81518152614162602083016140ce565b6020820152614173604083016140ce565b6040820152614184606083016140ce565b60608201526080828101519082015260a08201516001600160401b038111156141ab575f5ffd5b6141b78682850161400c565b60a08301525060c082810151908201526141d360e083016140ce565b60e0820152949350505050565b60208082526024908201527f546f6b656e52656d6f74653a207a65726f20726571756972656420676173206c6040820152631a5b5a5d60e21b606082015260800190565b6020808252602b908201527f546f6b656e52656d6f74653a207a65726f2064657374696e6174696f6e20626c60408201526a1bd8dad8da185a5b88125160aa1b606082015260800190565b60208082526037908201527f546f6b656e52656d6f74653a207a65726f2064657374696e6174696f6e20746f60408201527f6b656e207472616e736665727265722061646472657373000000000000000000606082015260800190565b81516001600160a01b0316815260208083015190820152604081016108c6565b8235815261012081016020840135614303816138c9565b6001600160a01b03166020830152604084013561431f816138c9565b6001600160a01b03166040830152614339606085016138dd565b6001600160a01b031660608301526080848101359083015260a0808501359083015260c0808501359083015261437160e085016138dd565b6001600160a01b031660e083015261010090910191909152919050565b5f5f8335601e198436030181126143a3575f5ffd5b8301803591506001600160401b038211156143bc575f5ffd5b6020019150368190038213156137c0575f5ffd5b602081528151602082015260018060a01b03602083015116604082015260018060a01b0360408301511660608201525f606083015161441a60808401826001600160a01b03169052565b50608083015160a083015260a083015161010060c084015261444061012084018261392a565b905060c084015160e084015260e08401516144676101008501826001600160a01b03169052565b509392505050565b5f5f8335601e19843603018112614484575f5ffd5b83016020810192503590506001600160401b038111156144a2575f5ffd5b8036038213156137c0575f5ffd5b81835281816020850137505f828201602090810191909152601f909101601f19169091010190565b60408082528335908201525f6144f0602085016138dd565b6001600160a01b0316606083015261450a604085016138dd565b6001600160a01b03166080830152614525606085018561446f565b61016060a085015261453c6101a0850182846144b0565b608087013560c08681019190915260a088013560e0870152909250614563915086016138dd565b6001600160a01b031661010084015261457e60e086016138dd565b6001600160a01b031661012084015261459a61010086016138dd565b6001600160a01b0381166101408501525061012085013561016084015261014090940135610180830152506020015290565b602081526145e66020820183516001600160a01b03169052565b602082015160408201525f604083015161460b60608401826001600160a01b03169052565b5060608301516001600160a01b038116608084015250608083015160a083015260a083015161016060c084015261464661018084018261392a565b905060c084015160e084015260e084015161466d6101008501826001600160a01b03169052565b506101008401516101208401526101208401516146966101408501826001600160a01b03169052565b506101408401516101608401528091505092915050565b80820281158282048414176108c6576108c6613e87565b818103818111156108c6576108c6613e87565b8681526001600160a01b0386811660208301528581166040830152841660608201526080810183905260c060a082018190525f906147179083018461392a565b98975050505050505050565b6020808252603a908201527f546f6b656e52656d6f74653a20696e76616c69642064657374696e6174696f6e60408201527f20746f6b656e207472616e736665727265722061646472657373000000000000606082015260800190565b601f82111561109a57805f5260205f20601f840160051c810160208510156147a55750805b601f840160051c820191505b818110156119e6575f81556001016147b1565b81516001600160401b038111156147dd576147dd613b02565b6147f1816147eb8454613db4565b84614780565b6020601f821160018114614823575f831561480c5750848201515b5f19600385901b1c1916600184901b1784556119e6565b5f84815260208120601f198516915b828110156148525787850151825560209485019460019092019101614832565b508482101561486f57868401515f19600387901b60f8161c191681555b50505050600190811b01905550565b5f8261489857634e487b7160e01b5f52601260045260245ffd5b500490565b60ff82811682821603908111156108c6576108c6613e87565b6001815b60018411156148f1578085048111156148d5576148d5613e87565b60018416156148e357908102905b60019390931c9280026148ba565b935093915050565b5f82614907575060016108c6565b8161491357505f6108c6565b816001811461492957600281146149335761494f565b60019150506108c6565b60ff84111561494457614944613e87565b50506001821b6108c6565b5060208310610133831016604e8410600b8410161715614972575081810a6108c6565b61497e5f1984846148b6565b805f190482111561499157614991613e87565b029392505050565b5f61164783836148f956fe52c63247e1f47db19d5ce0460030c497f067ca4cebf71ba98eeadabe20bace00600d6a9b283d1eda563de594ce4843869b6f128a4baa222422ed94a60b0cef0063108c0ddc3d3ab0f279f27d4b9ca8220d5efbdbeeb537d4327d5cf9e982b10063108c0ddc3d3ab0f279f27d4b9ca8220d5efbdbeeb537d4327d5cf9e982b102a2646970667358221220929b651ae379ef527c73f7d9b86f08a7fc65dac895269c9edeb8bb7474a41d4f64736f6c634300081e0033",
}

// ERC20TokenRemoteUpgradeableABI is the input ABI used to generate the binding from.
// Deprecated: Use ERC20TokenRemoteUpgradeableMetaData.ABI instead.
var ERC20TokenRemoteUpgradeableABI = ERC20TokenRemoteUpgradeableMetaData.ABI

// ERC20TokenRemoteUpgradeableBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use ERC20TokenRemoteUpgradeableMetaData.Bin instead.
var ERC20TokenRemoteUpgradeableBin = ERC20TokenRemoteUpgradeableMetaData.Bin

// DeployERC20TokenRemoteUpgradeable deploys a new Ethereum contract, binding an instance of ERC20TokenRemoteUpgradeable to it.
func DeployERC20TokenRemoteUpgradeable(auth *bind.TransactOpts, backend bind.ContractBackend, init uint8) (common.Address, *types.Transaction, *ERC20TokenRemoteUpgradeable, error) {
	parsed, err := ERC20TokenRemoteUpgradeableMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(ERC20TokenRemoteUpgradeableBin), backend, init)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &ERC20TokenRemoteUpgradeable{ERC20TokenRemoteUpgradeableCaller: ERC20TokenRemoteUpgradeableCaller{contract: contract}, ERC20TokenRemoteUpgradeableTransactor: ERC20TokenRemoteUpgradeableTransactor{contract: contract}, ERC20TokenRemoteUpgradeableFilterer: ERC20TokenRemoteUpgradeableFilterer{contract: contract}}, nil
}

// ERC20TokenRemoteUpgradeable is an auto generated Go binding around an Ethereum contract.
type ERC20TokenRemoteUpgradeable struct {
	ERC20TokenRemoteUpgradeableCaller     // Read-only binding to the contract
	ERC20TokenRemoteUpgradeableTransactor // Write-only binding to the contract
	ERC20TokenRemoteUpgradeableFilterer   // Log filterer for contract events
}

// ERC20TokenRemoteUpgradeableCaller is an auto generated read-only Go binding around an Ethereum contract.
type ERC20TokenRemoteUpgradeableCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ERC20TokenRemoteUpgradeableTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ERC20TokenRemoteUpgradeableTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ERC20TokenRemoteUpgradeableFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ERC20TokenRemoteUpgradeableFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ERC20TokenRemoteUpgradeableSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ERC20TokenRemoteUpgradeableSession struct {
	Contract     *ERC20TokenRemoteUpgradeable // Generic contract binding to set the session for
	CallOpts     bind.CallOpts                // Call options to use throughout this session
	TransactOpts bind.TransactOpts            // Transaction auth options to use throughout this session
}

// ERC20TokenRemoteUpgradeableCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ERC20TokenRemoteUpgradeableCallerSession struct {
	Contract *ERC20TokenRemoteUpgradeableCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                      // Call options to use throughout this session
}

// ERC20TokenRemoteUpgradeableTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ERC20TokenRemoteUpgradeableTransactorSession struct {
	Contract     *ERC20TokenRemoteUpgradeableTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                      // Transaction auth options to use throughout this session
}

// ERC20TokenRemoteUpgradeableRaw is an auto generated low-level Go binding around an Ethereum contract.
type ERC20TokenRemoteUpgradeableRaw struct {
	Contract *ERC20TokenRemoteUpgradeable // Generic contract binding to access the raw methods on
}

// ERC20TokenRemoteUpgradeableCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ERC20TokenRemoteUpgradeableCallerRaw struct {
	Contract *ERC20TokenRemoteUpgradeableCaller // Generic read-only contract binding to access the raw methods on
}

// ERC20TokenRemoteUpgradeableTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ERC20TokenRemoteUpgradeableTransactorRaw struct {
	Contract *ERC20TokenRemoteUpgradeableTransactor // Generic write-only contract binding to access the raw methods on
}

// NewERC20TokenRemoteUpgradeable creates a new instance of ERC20TokenRemoteUpgradeable, bound to a specific deployed contract.
func NewERC20TokenRemoteUpgradeable(address common.Address, backend bind.ContractBackend) (*ERC20TokenRemoteUpgradeable, error) {
	contract, err := bindERC20TokenRemoteUpgradeable(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ERC20TokenRemoteUpgradeable{ERC20TokenRemoteUpgradeableCaller: ERC20TokenRemoteUpgradeableCaller{contract: contract}, ERC20TokenRemoteUpgradeableTransactor: ERC20TokenRemoteUpgradeableTransactor{contract: contract}, ERC20TokenRemoteUpgradeableFilterer: ERC20TokenRemoteUpgradeableFilterer{contract: contract}}, nil
}

// NewERC20TokenRemoteUpgradeableCaller creates a new read-only instance of ERC20TokenRemoteUpgradeable, bound to a specific deployed contract.
func NewERC20TokenRemoteUpgradeableCaller(address common.Address, caller bind.ContractCaller) (*ERC20TokenRemoteUpgradeableCaller, error) {
	contract, err := bindERC20TokenRemoteUpgradeable(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ERC20TokenRemoteUpgradeableCaller{contract: contract}, nil
}

// NewERC20TokenRemoteUpgradeableTransactor creates a new write-only instance of ERC20TokenRemoteUpgradeable, bound to a specific deployed contract.
func NewERC20TokenRemoteUpgradeableTransactor(address common.Address, transactor bind.ContractTransactor) (*ERC20TokenRemoteUpgradeableTransactor, error) {
	contract, err := bindERC20TokenRemoteUpgradeable(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ERC20TokenRemoteUpgradeableTransactor{contract: contract}, nil
}

// NewERC20TokenRemoteUpgradeableFilterer creates a new log filterer instance of ERC20TokenRemoteUpgradeable, bound to a specific deployed contract.
func NewERC20TokenRemoteUpgradeableFilterer(address common.Address, filterer bind.ContractFilterer) (*ERC20TokenRemoteUpgradeableFilterer, error) {
	contract, err := bindERC20TokenRemoteUpgradeable(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ERC20TokenRemoteUpgradeableFilterer{contract: contract}, nil
}

// bindERC20TokenRemoteUpgradeable binds a generic wrapper to an already deployed contract.
func bindERC20TokenRemoteUpgradeable(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := ERC20TokenRemoteUpgradeableMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ERC20TokenRemoteUpgradeable *ERC20TokenRemoteUpgradeableRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ERC20TokenRemoteUpgradeable.Contract.ERC20TokenRemoteUpgradeableCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ERC20TokenRemoteUpgradeable *ERC20TokenRemoteUpgradeableRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ERC20TokenRemoteUpgradeable.Contract.ERC20TokenRemoteUpgradeableTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ERC20TokenRemoteUpgradeable *ERC20TokenRemoteUpgradeableRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ERC20TokenRemoteUpgradeable.Contract.ERC20TokenRemoteUpgradeableTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ERC20TokenRemoteUpgradeable *ERC20TokenRemoteUpgradeableCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ERC20TokenRemoteUpgradeable.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ERC20TokenRemoteUpgradeable *ERC20TokenRemoteUpgradeableTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ERC20TokenRemoteUpgradeable.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ERC20TokenRemoteUpgradeable *ERC20TokenRemoteUpgradeableTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ERC20TokenRemoteUpgradeable.Contract.contract.Transact(opts, method, params...)
}

// ERC20TOKENREMOTESTORAGELOCATION is a free data retrieval call binding the contract method 0x62431a65.
//
// Solidity: function ERC20_TOKEN_REMOTE_STORAGE_LOCATION() view returns(bytes32)
func (_ERC20TokenRemoteUpgradeable *ERC20TokenRemoteUpgradeableCaller) ERC20TOKENREMOTESTORAGELOCATION(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _ERC20TokenRemoteUpgradeable.contract.Call(opts, &out, "ERC20_TOKEN_REMOTE_STORAGE_LOCATION")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// ERC20TOKENREMOTESTORAGELOCATION is a free data retrieval call binding the contract method 0x62431a65.
//
// Solidity: function ERC20_TOKEN_REMOTE_STORAGE_LOCATION() view returns(bytes32)
func (_ERC20TokenRemoteUpgradeable *ERC20TokenRemoteUpgradeableSession) ERC20TOKENREMOTESTORAGELOCATION() ([32]byte, error) {
	return _ERC20TokenRemoteUpgradeable.Contract.ERC20TOKENREMOTESTORAGELOCATION(&_ERC20TokenRemoteUpgradeable.CallOpts)
}

// ERC20TOKENREMOTESTORAGELOCATION is a free data retrieval call binding the contract method 0x62431a65.
//
// Solidity: function ERC20_TOKEN_REMOTE_STORAGE_LOCATION() view returns(bytes32)
func (_ERC20TokenRemoteUpgradeable *ERC20TokenRemoteUpgradeableCallerSession) ERC20TOKENREMOTESTORAGELOCATION() ([32]byte, error) {
	return _ERC20TokenRemoteUpgradeable.Contract.ERC20TOKENREMOTESTORAGELOCATION(&_ERC20TokenRemoteUpgradeable.CallOpts)
}

// MULTIHOPCALLGASPERWORD is a free data retrieval call binding the contract method 0x15beb59f.
//
// Solidity: function MULTI_HOP_CALL_GAS_PER_WORD() view returns(uint256)
func (_ERC20TokenRemoteUpgradeable *ERC20TokenRemoteUpgradeableCaller) MULTIHOPCALLGASPERWORD(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _ERC20TokenRemoteUpgradeable.contract.Call(opts, &out, "MULTI_HOP_CALL_GAS_PER_WORD")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MULTIHOPCALLGASPERWORD is a free data retrieval call binding the contract method 0x15beb59f.
//
// Solidity: function MULTI_HOP_CALL_GAS_PER_WORD() view returns(uint256)
func (_ERC20TokenRemoteUpgradeable *ERC20TokenRemoteUpgradeableSession) MULTIHOPCALLGASPERWORD() (*big.Int, error) {
	return _ERC20TokenRemoteUpgradeable.Contract.MULTIHOPCALLGASPERWORD(&_ERC20TokenRemoteUpgradeable.CallOpts)
}

// MULTIHOPCALLGASPERWORD is a free data retrieval call binding the contract method 0x15beb59f.
//
// Solidity: function MULTI_HOP_CALL_GAS_PER_WORD() view returns(uint256)
func (_ERC20TokenRemoteUpgradeable *ERC20TokenRemoteUpgradeableCallerSession) MULTIHOPCALLGASPERWORD() (*big.Int, error) {
	return _ERC20TokenRemoteUpgradeable.Contract.MULTIHOPCALLGASPERWORD(&_ERC20TokenRemoteUpgradeable.CallOpts)
}

// MULTIHOPCALLREQUIREDGAS is a free data retrieval call binding the contract method 0x71717c18.
//
// Solidity: function MULTI_HOP_CALL_REQUIRED_GAS() view returns(uint256)
func (_ERC20TokenRemoteUpgradeable *ERC20TokenRemoteUpgradeableCaller) MULTIHOPCALLREQUIREDGAS(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _ERC20TokenRemoteUpgradeable.contract.Call(opts, &out, "MULTI_HOP_CALL_REQUIRED_GAS")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MULTIHOPCALLREQUIREDGAS is a free data retrieval call binding the contract method 0x71717c18.
//
// Solidity: function MULTI_HOP_CALL_REQUIRED_GAS() view returns(uint256)
func (_ERC20TokenRemoteUpgradeable *ERC20TokenRemoteUpgradeableSession) MULTIHOPCALLREQUIREDGAS() (*big.Int, error) {
	return _ERC20TokenRemoteUpgradeable.Contract.MULTIHOPCALLREQUIREDGAS(&_ERC20TokenRemoteUpgradeable.CallOpts)
}

// MULTIHOPCALLREQUIREDGAS is a free data retrieval call binding the contract method 0x71717c18.
//
// Solidity: function MULTI_HOP_CALL_REQUIRED_GAS() view returns(uint256)
func (_ERC20TokenRemoteUpgradeable *ERC20TokenRemoteUpgradeableCallerSession) MULTIHOPCALLREQUIREDGAS() (*big.Int, error) {
	return _ERC20TokenRemoteUpgradeable.Contract.MULTIHOPCALLREQUIREDGAS(&_ERC20TokenRemoteUpgradeable.CallOpts)
}

// MULTIHOPSENDREQUIREDGAS is a free data retrieval call binding the contract method 0x5507f3d1.
//
// Solidity: function MULTI_HOP_SEND_REQUIRED_GAS() view returns(uint256)
func (_ERC20TokenRemoteUpgradeable *ERC20TokenRemoteUpgradeableCaller) MULTIHOPSENDREQUIREDGAS(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _ERC20TokenRemoteUpgradeable.contract.Call(opts, &out, "MULTI_HOP_SEND_REQUIRED_GAS")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MULTIHOPSENDREQUIREDGAS is a free data retrieval call binding the contract method 0x5507f3d1.
//
// Solidity: function MULTI_HOP_SEND_REQUIRED_GAS() view returns(uint256)
func (_ERC20TokenRemoteUpgradeable *ERC20TokenRemoteUpgradeableSession) MULTIHOPSENDREQUIREDGAS() (*big.Int, error) {
	return _ERC20TokenRemoteUpgradeable.Contract.MULTIHOPSENDREQUIREDGAS(&_ERC20TokenRemoteUpgradeable.CallOpts)
}

// MULTIHOPSENDREQUIREDGAS is a free data retrieval call binding the contract method 0x5507f3d1.
//
// Solidity: function MULTI_HOP_SEND_REQUIRED_GAS() view returns(uint256)
func (_ERC20TokenRemoteUpgradeable *ERC20TokenRemoteUpgradeableCallerSession) MULTIHOPSENDREQUIREDGAS() (*big.Int, error) {
	return _ERC20TokenRemoteUpgradeable.Contract.MULTIHOPSENDREQUIREDGAS(&_ERC20TokenRemoteUpgradeable.CallOpts)
}

// REGISTERREMOTEREQUIREDGAS is a free data retrieval call binding the contract method 0x254ac160.
//
// Solidity: function REGISTER_REMOTE_REQUIRED_GAS() view returns(uint256)
func (_ERC20TokenRemoteUpgradeable *ERC20TokenRemoteUpgradeableCaller) REGISTERREMOTEREQUIREDGAS(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _ERC20TokenRemoteUpgradeable.contract.Call(opts, &out, "REGISTER_REMOTE_REQUIRED_GAS")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// REGISTERREMOTEREQUIREDGAS is a free data retrieval call binding the contract method 0x254ac160.
//
// Solidity: function REGISTER_REMOTE_REQUIRED_GAS() view returns(uint256)
func (_ERC20TokenRemoteUpgradeable *ERC20TokenRemoteUpgradeableSession) REGISTERREMOTEREQUIREDGAS() (*big.Int, error) {
	return _ERC20TokenRemoteUpgradeable.Contract.REGISTERREMOTEREQUIREDGAS(&_ERC20TokenRemoteUpgradeable.CallOpts)
}

// REGISTERREMOTEREQUIREDGAS is a free data retrieval call binding the contract method 0x254ac160.
//
// Solidity: function REGISTER_REMOTE_REQUIRED_GAS() view returns(uint256)
func (_ERC20TokenRemoteUpgradeable *ERC20TokenRemoteUpgradeableCallerSession) REGISTERREMOTEREQUIREDGAS() (*big.Int, error) {
	return _ERC20TokenRemoteUpgradeable.Contract.REGISTERREMOTEREQUIREDGAS(&_ERC20TokenRemoteUpgradeable.CallOpts)
}

// TELEPORTERROUTEDAPPSTORAGELOCATION is a free data retrieval call binding the contract method 0xc1c6353e.
//
// Solidity: function TELEPORTER_ROUTED_APP_STORAGE_LOCATION() view returns(bytes32)
func (_ERC20TokenRemoteUpgradeable *ERC20TokenRemoteUpgradeableCaller) TELEPORTERROUTEDAPPSTORAGELOCATION(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _ERC20TokenRemoteUpgradeable.contract.Call(opts, &out, "TELEPORTER_ROUTED_APP_STORAGE_LOCATION")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// TELEPORTERROUTEDAPPSTORAGELOCATION is a free data retrieval call binding the contract method 0xc1c6353e.
//
// Solidity: function TELEPORTER_ROUTED_APP_STORAGE_LOCATION() view returns(bytes32)
func (_ERC20TokenRemoteUpgradeable *ERC20TokenRemoteUpgradeableSession) TELEPORTERROUTEDAPPSTORAGELOCATION() ([32]byte, error) {
	return _ERC20TokenRemoteUpgradeable.Contract.TELEPORTERROUTEDAPPSTORAGELOCATION(&_ERC20TokenRemoteUpgradeable.CallOpts)
}

// TELEPORTERROUTEDAPPSTORAGELOCATION is a free data retrieval call binding the contract method 0xc1c6353e.
//
// Solidity: function TELEPORTER_ROUTED_APP_STORAGE_LOCATION() view returns(bytes32)
func (_ERC20TokenRemoteUpgradeable *ERC20TokenRemoteUpgradeableCallerSession) TELEPORTERROUTEDAPPSTORAGELOCATION() ([32]byte, error) {
	return _ERC20TokenRemoteUpgradeable.Contract.TELEPORTERROUTEDAPPSTORAGELOCATION(&_ERC20TokenRemoteUpgradeable.CallOpts)
}

// TOKENREMOTESTORAGELOCATION is a free data retrieval call binding the contract method 0x35cac159.
//
// Solidity: function TOKEN_REMOTE_STORAGE_LOCATION() view returns(bytes32)
func (_ERC20TokenRemoteUpgradeable *ERC20TokenRemoteUpgradeableCaller) TOKENREMOTESTORAGELOCATION(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _ERC20TokenRemoteUpgradeable.contract.Call(opts, &out, "TOKEN_REMOTE_STORAGE_LOCATION")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// TOKENREMOTESTORAGELOCATION is a free data retrieval call binding the contract method 0x35cac159.
//
// Solidity: function TOKEN_REMOTE_STORAGE_LOCATION() view returns(bytes32)
func (_ERC20TokenRemoteUpgradeable *ERC20TokenRemoteUpgradeableSession) TOKENREMOTESTORAGELOCATION() ([32]byte, error) {
	return _ERC20TokenRemoteUpgradeable.Contract.TOKENREMOTESTORAGELOCATION(&_ERC20TokenRemoteUpgradeable.CallOpts)
}

// TOKENREMOTESTORAGELOCATION is a free data retrieval call binding the contract method 0x35cac159.
//
// Solidity: function TOKEN_REMOTE_STORAGE_LOCATION() view returns(bytes32)
func (_ERC20TokenRemoteUpgradeable *ERC20TokenRemoteUpgradeableCallerSession) TOKENREMOTESTORAGELOCATION() ([32]byte, error) {
	return _ERC20TokenRemoteUpgradeable.Contract.TOKENREMOTESTORAGELOCATION(&_ERC20TokenRemoteUpgradeable.CallOpts)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_ERC20TokenRemoteUpgradeable *ERC20TokenRemoteUpgradeableCaller) Allowance(opts *bind.CallOpts, owner common.Address, spender common.Address) (*big.Int, error) {
	var out []interface{}
	err := _ERC20TokenRemoteUpgradeable.contract.Call(opts, &out, "allowance", owner, spender)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_ERC20TokenRemoteUpgradeable *ERC20TokenRemoteUpgradeableSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _ERC20TokenRemoteUpgradeable.Contract.Allowance(&_ERC20TokenRemoteUpgradeable.CallOpts, owner, spender)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_ERC20TokenRemoteUpgradeable *ERC20TokenRemoteUpgradeableCallerSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _ERC20TokenRemoteUpgradeable.Contract.Allowance(&_ERC20TokenRemoteUpgradeable.CallOpts, owner, spender)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_ERC20TokenRemoteUpgradeable *ERC20TokenRemoteUpgradeableCaller) BalanceOf(opts *bind.CallOpts, account common.Address) (*big.Int, error) {
	var out []interface{}
	err := _ERC20TokenRemoteUpgradeable.contract.Call(opts, &out, "balanceOf", account)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_ERC20TokenRemoteUpgradeable *ERC20TokenRemoteUpgradeableSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _ERC20TokenRemoteUpgradeable.Contract.BalanceOf(&_ERC20TokenRemoteUpgradeable.CallOpts, account)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_ERC20TokenRemoteUpgradeable *ERC20TokenRemoteUpgradeableCallerSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _ERC20TokenRemoteUpgradeable.Contract.BalanceOf(&_ERC20TokenRemoteUpgradeable.CallOpts, account)
}

// CalculateNumWords is a free data retrieval call binding the contract method 0xf3f981d8.
//
// Solidity: function calculateNumWords(uint256 payloadSize) pure returns(uint256)
func (_ERC20TokenRemoteUpgradeable *ERC20TokenRemoteUpgradeableCaller) CalculateNumWords(opts *bind.CallOpts, payloadSize *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _ERC20TokenRemoteUpgradeable.contract.Call(opts, &out, "calculateNumWords", payloadSize)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CalculateNumWords is a free data retrieval call binding the contract method 0xf3f981d8.
//
// Solidity: function calculateNumWords(uint256 payloadSize) pure returns(uint256)
func (_ERC20TokenRemoteUpgradeable *ERC20TokenRemoteUpgradeableSession) CalculateNumWords(payloadSize *big.Int) (*big.Int, error) {
	return _ERC20TokenRemoteUpgradeable.Contract.CalculateNumWords(&_ERC20TokenRemoteUpgradeable.CallOpts, payloadSize)
}

// CalculateNumWords is a free data retrieval call binding the contract method 0xf3f981d8.
//
// Solidity: function calculateNumWords(uint256 payloadSize) pure returns(uint256)
func (_ERC20TokenRemoteUpgradeable *ERC20TokenRemoteUpgradeableCallerSession) CalculateNumWords(payloadSize *big.Int) (*big.Int, error) {
	return _ERC20TokenRemoteUpgradeable.Contract.CalculateNumWords(&_ERC20TokenRemoteUpgradeable.CallOpts, payloadSize)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_ERC20TokenRemoteUpgradeable *ERC20TokenRemoteUpgradeableCaller) Decimals(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _ERC20TokenRemoteUpgradeable.contract.Call(opts, &out, "decimals")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_ERC20TokenRemoteUpgradeable *ERC20TokenRemoteUpgradeableSession) Decimals() (uint8, error) {
	return _ERC20TokenRemoteUpgradeable.Contract.Decimals(&_ERC20TokenRemoteUpgradeable.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_ERC20TokenRemoteUpgradeable *ERC20TokenRemoteUpgradeableCallerSession) Decimals() (uint8, error) {
	return _ERC20TokenRemoteUpgradeable.Contract.Decimals(&_ERC20TokenRemoteUpgradeable.CallOpts)
}

// GetBlockchainID is a free data retrieval call binding the contract method 0x4213cf78.
//
// Solidity: function getBlockchainID() view returns(bytes32)
func (_ERC20TokenRemoteUpgradeable *ERC20TokenRemoteUpgradeableCaller) GetBlockchainID(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _ERC20TokenRemoteUpgradeable.contract.Call(opts, &out, "getBlockchainID")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetBlockchainID is a free data retrieval call binding the contract method 0x4213cf78.
//
// Solidity: function getBlockchainID() view returns(bytes32)
func (_ERC20TokenRemoteUpgradeable *ERC20TokenRemoteUpgradeableSession) GetBlockchainID() ([32]byte, error) {
	return _ERC20TokenRemoteUpgradeable.Contract.GetBlockchainID(&_ERC20TokenRemoteUpgradeable.CallOpts)
}

// GetBlockchainID is a free data retrieval call binding the contract method 0x4213cf78.
//
// Solidity: function getBlockchainID() view returns(bytes32)
func (_ERC20TokenRemoteUpgradeable *ERC20TokenRemoteUpgradeableCallerSession) GetBlockchainID() ([32]byte, error) {
	return _ERC20TokenRemoteUpgradeable.Contract.GetBlockchainID(&_ERC20TokenRemoteUpgradeable.CallOpts)
}

// GetInitialReserveImbalance is a free data retrieval call binding the contract method 0xef793e2a.
//
// Solidity: function getInitialReserveImbalance() view returns(uint256)
func (_ERC20TokenRemoteUpgradeable *ERC20TokenRemoteUpgradeableCaller) GetInitialReserveImbalance(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _ERC20TokenRemoteUpgradeable.contract.Call(opts, &out, "getInitialReserveImbalance")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetInitialReserveImbalance is a free data retrieval call binding the contract method 0xef793e2a.
//
// Solidity: function getInitialReserveImbalance() view returns(uint256)
func (_ERC20TokenRemoteUpgradeable *ERC20TokenRemoteUpgradeableSession) GetInitialReserveImbalance() (*big.Int, error) {
	return _ERC20TokenRemoteUpgradeable.Contract.GetInitialReserveImbalance(&_ERC20TokenRemoteUpgradeable.CallOpts)
}

// GetInitialReserveImbalance is a free data retrieval call binding the contract method 0xef793e2a.
//
// Solidity: function getInitialReserveImbalance() view returns(uint256)
func (_ERC20TokenRemoteUpgradeable *ERC20TokenRemoteUpgradeableCallerSession) GetInitialReserveImbalance() (*big.Int, error) {
	return _ERC20TokenRemoteUpgradeable.Contract.GetInitialReserveImbalance(&_ERC20TokenRemoteUpgradeable.CallOpts)
}

// GetIsCollateralized is a free data retrieval call binding the contract method 0x02a30c7d.
//
// Solidity: function getIsCollateralized() view returns(bool)
func (_ERC20TokenRemoteUpgradeable *ERC20TokenRemoteUpgradeableCaller) GetIsCollateralized(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _ERC20TokenRemoteUpgradeable.contract.Call(opts, &out, "getIsCollateralized")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// GetIsCollateralized is a free data retrieval call binding the contract method 0x02a30c7d.
//
// Solidity: function getIsCollateralized() view returns(bool)
func (_ERC20TokenRemoteUpgradeable *ERC20TokenRemoteUpgradeableSession) GetIsCollateralized() (bool, error) {
	return _ERC20TokenRemoteUpgradeable.Contract.GetIsCollateralized(&_ERC20TokenRemoteUpgradeable.CallOpts)
}

// GetIsCollateralized is a free data retrieval call binding the contract method 0x02a30c7d.
//
// Solidity: function getIsCollateralized() view returns(bool)
func (_ERC20TokenRemoteUpgradeable *ERC20TokenRemoteUpgradeableCallerSession) GetIsCollateralized() (bool, error) {
	return _ERC20TokenRemoteUpgradeable.Contract.GetIsCollateralized(&_ERC20TokenRemoteUpgradeable.CallOpts)
}

// GetMultiplyOnRemote is a free data retrieval call binding the contract method 0x7ee3779a.
//
// Solidity: function getMultiplyOnRemote() view returns(bool)
func (_ERC20TokenRemoteUpgradeable *ERC20TokenRemoteUpgradeableCaller) GetMultiplyOnRemote(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _ERC20TokenRemoteUpgradeable.contract.Call(opts, &out, "getMultiplyOnRemote")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// GetMultiplyOnRemote is a free data retrieval call binding the contract method 0x7ee3779a.
//
// Solidity: function getMultiplyOnRemote() view returns(bool)
func (_ERC20TokenRemoteUpgradeable *ERC20TokenRemoteUpgradeableSession) GetMultiplyOnRemote() (bool, error) {
	return _ERC20TokenRemoteUpgradeable.Contract.GetMultiplyOnRemote(&_ERC20TokenRemoteUpgradeable.CallOpts)
}

// GetMultiplyOnRemote is a free data retrieval call binding the contract method 0x7ee3779a.
//
// Solidity: function getMultiplyOnRemote() view returns(bool)
func (_ERC20TokenRemoteUpgradeable *ERC20TokenRemoteUpgradeableCallerSession) GetMultiplyOnRemote() (bool, error) {
	return _ERC20TokenRemoteUpgradeable.Contract.GetMultiplyOnRemote(&_ERC20TokenRemoteUpgradeable.CallOpts)
}

// GetOutboundMessenger is a free data retrieval call binding the contract method 0xfc560821.
//
// Solidity: function getOutboundMessenger(bytes32 remoteBlockchainID) view returns(address)
func (_ERC20TokenRemoteUpgradeable *ERC20TokenRemoteUpgradeableCaller) GetOutboundMessenger(opts *bind.CallOpts, remoteBlockchainID [32]byte) (common.Address, error) {
	var out []interface{}
	err := _ERC20TokenRemoteUpgradeable.contract.Call(opts, &out, "getOutboundMessenger", remoteBlockchainID)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetOutboundMessenger is a free data retrieval call binding the contract method 0xfc560821.
//
// Solidity: function getOutboundMessenger(bytes32 remoteBlockchainID) view returns(address)
func (_ERC20TokenRemoteUpgradeable *ERC20TokenRemoteUpgradeableSession) GetOutboundMessenger(remoteBlockchainID [32]byte) (common.Address, error) {
	return _ERC20TokenRemoteUpgradeable.Contract.GetOutboundMessenger(&_ERC20TokenRemoteUpgradeable.CallOpts, remoteBlockchainID)
}

// GetOutboundMessenger is a free data retrieval call binding the contract method 0xfc560821.
//
// Solidity: function getOutboundMessenger(bytes32 remoteBlockchainID) view returns(address)
func (_ERC20TokenRemoteUpgradeable *ERC20TokenRemoteUpgradeableCallerSession) GetOutboundMessenger(remoteBlockchainID [32]byte) (common.Address, error) {
	return _ERC20TokenRemoteUpgradeable.Contract.GetOutboundMessenger(&_ERC20TokenRemoteUpgradeable.CallOpts, remoteBlockchainID)
}

// GetTokenHomeAddress is a free data retrieval call binding the contract method 0xc3cd6927.
//
// Solidity: function getTokenHomeAddress() view returns(address)
func (_ERC20TokenRemoteUpgradeable *ERC20TokenRemoteUpgradeableCaller) GetTokenHomeAddress(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ERC20TokenRemoteUpgradeable.contract.Call(opts, &out, "getTokenHomeAddress")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetTokenHomeAddress is a free data retrieval call binding the contract method 0xc3cd6927.
//
// Solidity: function getTokenHomeAddress() view returns(address)
func (_ERC20TokenRemoteUpgradeable *ERC20TokenRemoteUpgradeableSession) GetTokenHomeAddress() (common.Address, error) {
	return _ERC20TokenRemoteUpgradeable.Contract.GetTokenHomeAddress(&_ERC20TokenRemoteUpgradeable.CallOpts)
}

// GetTokenHomeAddress is a free data retrieval call binding the contract method 0xc3cd6927.
//
// Solidity: function getTokenHomeAddress() view returns(address)
func (_ERC20TokenRemoteUpgradeable *ERC20TokenRemoteUpgradeableCallerSession) GetTokenHomeAddress() (common.Address, error) {
	return _ERC20TokenRemoteUpgradeable.Contract.GetTokenHomeAddress(&_ERC20TokenRemoteUpgradeable.CallOpts)
}

// GetTokenHomeBlockchainID is a free data retrieval call binding the contract method 0xe0fd9cb8.
//
// Solidity: function getTokenHomeBlockchainID() view returns(bytes32)
func (_ERC20TokenRemoteUpgradeable *ERC20TokenRemoteUpgradeableCaller) GetTokenHomeBlockchainID(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _ERC20TokenRemoteUpgradeable.contract.Call(opts, &out, "getTokenHomeBlockchainID")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetTokenHomeBlockchainID is a free data retrieval call binding the contract method 0xe0fd9cb8.
//
// Solidity: function getTokenHomeBlockchainID() view returns(bytes32)
func (_ERC20TokenRemoteUpgradeable *ERC20TokenRemoteUpgradeableSession) GetTokenHomeBlockchainID() ([32]byte, error) {
	return _ERC20TokenRemoteUpgradeable.Contract.GetTokenHomeBlockchainID(&_ERC20TokenRemoteUpgradeable.CallOpts)
}

// GetTokenHomeBlockchainID is a free data retrieval call binding the contract method 0xe0fd9cb8.
//
// Solidity: function getTokenHomeBlockchainID() view returns(bytes32)
func (_ERC20TokenRemoteUpgradeable *ERC20TokenRemoteUpgradeableCallerSession) GetTokenHomeBlockchainID() ([32]byte, error) {
	return _ERC20TokenRemoteUpgradeable.Contract.GetTokenHomeBlockchainID(&_ERC20TokenRemoteUpgradeable.CallOpts)
}

// GetTokenMultiplier is a free data retrieval call binding the contract method 0x0733c8c8.
//
// Solidity: function getTokenMultiplier() view returns(uint256)
func (_ERC20TokenRemoteUpgradeable *ERC20TokenRemoteUpgradeableCaller) GetTokenMultiplier(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _ERC20TokenRemoteUpgradeable.contract.Call(opts, &out, "getTokenMultiplier")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetTokenMultiplier is a free data retrieval call binding the contract method 0x0733c8c8.
//
// Solidity: function getTokenMultiplier() view returns(uint256)
func (_ERC20TokenRemoteUpgradeable *ERC20TokenRemoteUpgradeableSession) GetTokenMultiplier() (*big.Int, error) {
	return _ERC20TokenRemoteUpgradeable.Contract.GetTokenMultiplier(&_ERC20TokenRemoteUpgradeable.CallOpts)
}

// GetTokenMultiplier is a free data retrieval call binding the contract method 0x0733c8c8.
//
// Solidity: function getTokenMultiplier() view returns(uint256)
func (_ERC20TokenRemoteUpgradeable *ERC20TokenRemoteUpgradeableCallerSession) GetTokenMultiplier() (*big.Int, error) {
	return _ERC20TokenRemoteUpgradeable.Contract.GetTokenMultiplier(&_ERC20TokenRemoteUpgradeable.CallOpts)
}

// IsInboundMessenger is a free data retrieval call binding the contract method 0x95a19443.
//
// Solidity: function isInboundMessenger(bytes32 remoteBlockchainID, address messenger) view returns(bool)
func (_ERC20TokenRemoteUpgradeable *ERC20TokenRemoteUpgradeableCaller) IsInboundMessenger(opts *bind.CallOpts, remoteBlockchainID [32]byte, messenger common.Address) (bool, error) {
	var out []interface{}
	err := _ERC20TokenRemoteUpgradeable.contract.Call(opts, &out, "isInboundMessenger", remoteBlockchainID, messenger)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsInboundMessenger is a free data retrieval call binding the contract method 0x95a19443.
//
// Solidity: function isInboundMessenger(bytes32 remoteBlockchainID, address messenger) view returns(bool)
func (_ERC20TokenRemoteUpgradeable *ERC20TokenRemoteUpgradeableSession) IsInboundMessenger(remoteBlockchainID [32]byte, messenger common.Address) (bool, error) {
	return _ERC20TokenRemoteUpgradeable.Contract.IsInboundMessenger(&_ERC20TokenRemoteUpgradeable.CallOpts, remoteBlockchainID, messenger)
}

// IsInboundMessenger is a free data retrieval call binding the contract method 0x95a19443.
//
// Solidity: function isInboundMessenger(bytes32 remoteBlockchainID, address messenger) view returns(bool)
func (_ERC20TokenRemoteUpgradeable *ERC20TokenRemoteUpgradeableCallerSession) IsInboundMessenger(remoteBlockchainID [32]byte, messenger common.Address) (bool, error) {
	return _ERC20TokenRemoteUpgradeable.Contract.IsInboundMessenger(&_ERC20TokenRemoteUpgradeable.CallOpts, remoteBlockchainID, messenger)
}

// IsMessengerPaused is a free data retrieval call binding the contract method 0x83c002d0.
//
// Solidity: function isMessengerPaused(address messenger) view returns(bool)
func (_ERC20TokenRemoteUpgradeable *ERC20TokenRemoteUpgradeableCaller) IsMessengerPaused(opts *bind.CallOpts, messenger common.Address) (bool, error) {
	var out []interface{}
	err := _ERC20TokenRemoteUpgradeable.contract.Call(opts, &out, "isMessengerPaused", messenger)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsMessengerPaused is a free data retrieval call binding the contract method 0x83c002d0.
//
// Solidity: function isMessengerPaused(address messenger) view returns(bool)
func (_ERC20TokenRemoteUpgradeable *ERC20TokenRemoteUpgradeableSession) IsMessengerPaused(messenger common.Address) (bool, error) {
	return _ERC20TokenRemoteUpgradeable.Contract.IsMessengerPaused(&_ERC20TokenRemoteUpgradeable.CallOpts, messenger)
}

// IsMessengerPaused is a free data retrieval call binding the contract method 0x83c002d0.
//
// Solidity: function isMessengerPaused(address messenger) view returns(bool)
func (_ERC20TokenRemoteUpgradeable *ERC20TokenRemoteUpgradeableCallerSession) IsMessengerPaused(messenger common.Address) (bool, error) {
	return _ERC20TokenRemoteUpgradeable.Contract.IsMessengerPaused(&_ERC20TokenRemoteUpgradeable.CallOpts, messenger)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_ERC20TokenRemoteUpgradeable *ERC20TokenRemoteUpgradeableCaller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _ERC20TokenRemoteUpgradeable.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_ERC20TokenRemoteUpgradeable *ERC20TokenRemoteUpgradeableSession) Name() (string, error) {
	return _ERC20TokenRemoteUpgradeable.Contract.Name(&_ERC20TokenRemoteUpgradeable.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_ERC20TokenRemoteUpgradeable *ERC20TokenRemoteUpgradeableCallerSession) Name() (string, error) {
	return _ERC20TokenRemoteUpgradeable.Contract.Name(&_ERC20TokenRemoteUpgradeable.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_ERC20TokenRemoteUpgradeable *ERC20TokenRemoteUpgradeableCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ERC20TokenRemoteUpgradeable.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_ERC20TokenRemoteUpgradeable *ERC20TokenRemoteUpgradeableSession) Owner() (common.Address, error) {
	return _ERC20TokenRemoteUpgradeable.Contract.Owner(&_ERC20TokenRemoteUpgradeable.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_ERC20TokenRemoteUpgradeable *ERC20TokenRemoteUpgradeableCallerSession) Owner() (common.Address, error) {
	return _ERC20TokenRemoteUpgradeable.Contract.Owner(&_ERC20TokenRemoteUpgradeable.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_ERC20TokenRemoteUpgradeable *ERC20TokenRemoteUpgradeableCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _ERC20TokenRemoteUpgradeable.contract.Call(opts, &out, "symbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_ERC20TokenRemoteUpgradeable *ERC20TokenRemoteUpgradeableSession) Symbol() (string, error) {
	return _ERC20TokenRemoteUpgradeable.Contract.Symbol(&_ERC20TokenRemoteUpgradeable.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_ERC20TokenRemoteUpgradeable *ERC20TokenRemoteUpgradeableCallerSession) Symbol() (string, error) {
	return _ERC20TokenRemoteUpgradeable.Contract.Symbol(&_ERC20TokenRemoteUpgradeable.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_ERC20TokenRemoteUpgradeable *ERC20TokenRemoteUpgradeableCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _ERC20TokenRemoteUpgradeable.contract.Call(opts, &out, "totalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_ERC20TokenRemoteUpgradeable *ERC20TokenRemoteUpgradeableSession) TotalSupply() (*big.Int, error) {
	return _ERC20TokenRemoteUpgradeable.Contract.TotalSupply(&_ERC20TokenRemoteUpgradeable.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_ERC20TokenRemoteUpgradeable *ERC20TokenRemoteUpgradeableCallerSession) TotalSupply() (*big.Int, error) {
	return _ERC20TokenRemoteUpgradeable.Contract.TotalSupply(&_ERC20TokenRemoteUpgradeable.CallOpts)
}

// AddMessenger is a paid mutator transaction binding the contract method 0xb2432529.
//
// Solidity: function addMessenger(bytes32 remoteBlockchainID, address messenger) returns()
func (_ERC20TokenRemoteUpgradeable *ERC20TokenRemoteUpgradeableTransactor) AddMessenger(opts *bind.TransactOpts, remoteBlockchainID [32]byte, messenger common.Address) (*types.Transaction, error) {
	return _ERC20TokenRemoteUpgradeable.contract.Transact(opts, "addMessenger", remoteBlockchainID, messenger)
}

// AddMessenger is a paid mutator transaction binding the contract method 0xb2432529.
//
// Solidity: function addMessenger(bytes32 remoteBlockchainID, address messenger) returns()
func (_ERC20TokenRemoteUpgradeable *ERC20TokenRemoteUpgradeableSession) AddMessenger(remoteBlockchainID [32]byte, messenger common.Address) (*types.Transaction, error) {
	return _ERC20TokenRemoteUpgradeable.Contract.AddMessenger(&_ERC20TokenRemoteUpgradeable.TransactOpts, remoteBlockchainID, messenger)
}

// AddMessenger is a paid mutator transaction binding the contract method 0xb2432529.
//
// Solidity: function addMessenger(bytes32 remoteBlockchainID, address messenger) returns()
func (_ERC20TokenRemoteUpgradeable *ERC20TokenRemoteUpgradeableTransactorSession) AddMessenger(remoteBlockchainID [32]byte, messenger common.Address) (*types.Transaction, error) {
	return _ERC20TokenRemoteUpgradeable.Contract.AddMessenger(&_ERC20TokenRemoteUpgradeable.TransactOpts, remoteBlockchainID, messenger)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 value) returns(bool)
func (_ERC20TokenRemoteUpgradeable *ERC20TokenRemoteUpgradeableTransactor) Approve(opts *bind.TransactOpts, spender common.Address, value *big.Int) (*types.Transaction, error) {
	return _ERC20TokenRemoteUpgradeable.contract.Transact(opts, "approve", spender, value)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 value) returns(bool)
func (_ERC20TokenRemoteUpgradeable *ERC20TokenRemoteUpgradeableSession) Approve(spender common.Address, value *big.Int) (*types.Transaction, error) {
	return _ERC20TokenRemoteUpgradeable.Contract.Approve(&_ERC20TokenRemoteUpgradeable.TransactOpts, spender, value)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 value) returns(bool)
func (_ERC20TokenRemoteUpgradeable *ERC20TokenRemoteUpgradeableTransactorSession) Approve(spender common.Address, value *big.Int) (*types.Transaction, error) {
	return _ERC20TokenRemoteUpgradeable.Contract.Approve(&_ERC20TokenRemoteUpgradeable.TransactOpts, spender, value)
}

// Initialize is a paid mutator transaction binding the contract method 0xd73a0641.
//
// Solidity: function initialize((bytes32,address,address,bytes32,address,uint8) settings, string tokenName, string tokenSymbol, uint8 tokenDecimals) returns()
func (_ERC20TokenRemoteUpgradeable *ERC20TokenRemoteUpgradeableTransactor) Initialize(opts *bind.TransactOpts, settings TokenRemoteSettings, tokenName string, tokenSymbol string, tokenDecimals uint8) (*types.Transaction, error) {
	return _ERC20TokenRemoteUpgradeable.contract.Transact(opts, "initialize", settings, tokenName, tokenSymbol, tokenDecimals)
}

// Initialize is a paid mutator transaction binding the contract method 0xd73a0641.
//
// Solidity: function initialize((bytes32,address,address,bytes32,address,uint8) settings, string tokenName, string tokenSymbol, uint8 tokenDecimals) returns()
func (_ERC20TokenRemoteUpgradeable *ERC20TokenRemoteUpgradeableSession) Initialize(settings TokenRemoteSettings, tokenName string, tokenSymbol string, tokenDecimals uint8) (*types.Transaction, error) {
	return _ERC20TokenRemoteUpgradeable.Contract.Initialize(&_ERC20TokenRemoteUpgradeable.TransactOpts, settings, tokenName, tokenSymbol, tokenDecimals)
}

// Initialize is a paid mutator transaction binding the contract method 0xd73a0641.
//
// Solidity: function initialize((bytes32,address,address,bytes32,address,uint8) settings, string tokenName, string tokenSymbol, uint8 tokenDecimals) returns()
func (_ERC20TokenRemoteUpgradeable *ERC20TokenRemoteUpgradeableTransactorSession) Initialize(settings TokenRemoteSettings, tokenName string, tokenSymbol string, tokenDecimals uint8) (*types.Transaction, error) {
	return _ERC20TokenRemoteUpgradeable.Contract.Initialize(&_ERC20TokenRemoteUpgradeable.TransactOpts, settings, tokenName, tokenSymbol, tokenDecimals)
}

// PauseMessenger is a paid mutator transaction binding the contract method 0x03f78951.
//
// Solidity: function pauseMessenger(address messenger) returns()
func (_ERC20TokenRemoteUpgradeable *ERC20TokenRemoteUpgradeableTransactor) PauseMessenger(opts *bind.TransactOpts, messenger common.Address) (*types.Transaction, error) {
	return _ERC20TokenRemoteUpgradeable.contract.Transact(opts, "pauseMessenger", messenger)
}

// PauseMessenger is a paid mutator transaction binding the contract method 0x03f78951.
//
// Solidity: function pauseMessenger(address messenger) returns()
func (_ERC20TokenRemoteUpgradeable *ERC20TokenRemoteUpgradeableSession) PauseMessenger(messenger common.Address) (*types.Transaction, error) {
	return _ERC20TokenRemoteUpgradeable.Contract.PauseMessenger(&_ERC20TokenRemoteUpgradeable.TransactOpts, messenger)
}

// PauseMessenger is a paid mutator transaction binding the contract method 0x03f78951.
//
// Solidity: function pauseMessenger(address messenger) returns()
func (_ERC20TokenRemoteUpgradeable *ERC20TokenRemoteUpgradeableTransactorSession) PauseMessenger(messenger common.Address) (*types.Transaction, error) {
	return _ERC20TokenRemoteUpgradeable.Contract.PauseMessenger(&_ERC20TokenRemoteUpgradeable.TransactOpts, messenger)
}

// ReceiveTeleporterMessage is a paid mutator transaction binding the contract method 0xc868efaa.
//
// Solidity: function receiveTeleporterMessage(bytes32 sourceBlockchainID, address originSenderAddress, bytes message) returns()
func (_ERC20TokenRemoteUpgradeable *ERC20TokenRemoteUpgradeableTransactor) ReceiveTeleporterMessage(opts *bind.TransactOpts, sourceBlockchainID [32]byte, originSenderAddress common.Address, message []byte) (*types.Transaction, error) {
	return _ERC20TokenRemoteUpgradeable.contract.Transact(opts, "receiveTeleporterMessage", sourceBlockchainID, originSenderAddress, message)
}

// ReceiveTeleporterMessage is a paid mutator transaction binding the contract method 0xc868efaa.
//
// Solidity: function receiveTeleporterMessage(bytes32 sourceBlockchainID, address originSenderAddress, bytes message) returns()
func (_ERC20TokenRemoteUpgradeable *ERC20TokenRemoteUpgradeableSession) ReceiveTeleporterMessage(sourceBlockchainID [32]byte, originSenderAddress common.Address, message []byte) (*types.Transaction, error) {
	return _ERC20TokenRemoteUpgradeable.Contract.ReceiveTeleporterMessage(&_ERC20TokenRemoteUpgradeable.TransactOpts, sourceBlockchainID, originSenderAddress, message)
}

// ReceiveTeleporterMessage is a paid mutator transaction binding the contract method 0xc868efaa.
//
// Solidity: function receiveTeleporterMessage(bytes32 sourceBlockchainID, address originSenderAddress, bytes message) returns()
func (_ERC20TokenRemoteUpgradeable *ERC20TokenRemoteUpgradeableTransactorSession) ReceiveTeleporterMessage(sourceBlockchainID [32]byte, originSenderAddress common.Address, message []byte) (*types.Transaction, error) {
	return _ERC20TokenRemoteUpgradeable.Contract.ReceiveTeleporterMessage(&_ERC20TokenRemoteUpgradeable.TransactOpts, sourceBlockchainID, originSenderAddress, message)
}

// RegisterWithHome is a paid mutator transaction binding the contract method 0xb8a46d02.
//
// Solidity: function registerWithHome((address,uint256) feeInfo) returns()
func (_ERC20TokenRemoteUpgradeable *ERC20TokenRemoteUpgradeableTransactor) RegisterWithHome(opts *bind.TransactOpts, feeInfo TeleporterFeeInfo) (*types.Transaction, error) {
	return _ERC20TokenRemoteUpgradeable.contract.Transact(opts, "registerWithHome", feeInfo)
}

// RegisterWithHome is a paid mutator transaction binding the contract method 0xb8a46d02.
//
// Solidity: function registerWithHome((address,uint256) feeInfo) returns()
func (_ERC20TokenRemoteUpgradeable *ERC20TokenRemoteUpgradeableSession) RegisterWithHome(feeInfo TeleporterFeeInfo) (*types.Transaction, error) {
	return _ERC20TokenRemoteUpgradeable.Contract.RegisterWithHome(&_ERC20TokenRemoteUpgradeable.TransactOpts, feeInfo)
}

// RegisterWithHome is a paid mutator transaction binding the contract method 0xb8a46d02.
//
// Solidity: function registerWithHome((address,uint256) feeInfo) returns()
func (_ERC20TokenRemoteUpgradeable *ERC20TokenRemoteUpgradeableTransactorSession) RegisterWithHome(feeInfo TeleporterFeeInfo) (*types.Transaction, error) {
	return _ERC20TokenRemoteUpgradeable.Contract.RegisterWithHome(&_ERC20TokenRemoteUpgradeable.TransactOpts, feeInfo)
}

// RemoveMessenger is a paid mutator transaction binding the contract method 0xf4be94a7.
//
// Solidity: function removeMessenger(bytes32 remoteBlockchainID, address messenger) returns()
func (_ERC20TokenRemoteUpgradeable *ERC20TokenRemoteUpgradeableTransactor) RemoveMessenger(opts *bind.TransactOpts, remoteBlockchainID [32]byte, messenger common.Address) (*types.Transaction, error) {
	return _ERC20TokenRemoteUpgradeable.contract.Transact(opts, "removeMessenger", remoteBlockchainID, messenger)
}

// RemoveMessenger is a paid mutator transaction binding the contract method 0xf4be94a7.
//
// Solidity: function removeMessenger(bytes32 remoteBlockchainID, address messenger) returns()
func (_ERC20TokenRemoteUpgradeable *ERC20TokenRemoteUpgradeableSession) RemoveMessenger(remoteBlockchainID [32]byte, messenger common.Address) (*types.Transaction, error) {
	return _ERC20TokenRemoteUpgradeable.Contract.RemoveMessenger(&_ERC20TokenRemoteUpgradeable.TransactOpts, remoteBlockchainID, messenger)
}

// RemoveMessenger is a paid mutator transaction binding the contract method 0xf4be94a7.
//
// Solidity: function removeMessenger(bytes32 remoteBlockchainID, address messenger) returns()
func (_ERC20TokenRemoteUpgradeable *ERC20TokenRemoteUpgradeableTransactorSession) RemoveMessenger(remoteBlockchainID [32]byte, messenger common.Address) (*types.Transaction, error) {
	return _ERC20TokenRemoteUpgradeable.Contract.RemoveMessenger(&_ERC20TokenRemoteUpgradeable.TransactOpts, remoteBlockchainID, messenger)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_ERC20TokenRemoteUpgradeable *ERC20TokenRemoteUpgradeableTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ERC20TokenRemoteUpgradeable.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_ERC20TokenRemoteUpgradeable *ERC20TokenRemoteUpgradeableSession) RenounceOwnership() (*types.Transaction, error) {
	return _ERC20TokenRemoteUpgradeable.Contract.RenounceOwnership(&_ERC20TokenRemoteUpgradeable.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_ERC20TokenRemoteUpgradeable *ERC20TokenRemoteUpgradeableTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _ERC20TokenRemoteUpgradeable.Contract.RenounceOwnership(&_ERC20TokenRemoteUpgradeable.TransactOpts)
}

// Send is a paid mutator transaction binding the contract method 0x5d16225d.
//
// Solidity: function send((bytes32,address,address,address,uint256,uint256,uint256,address) input, uint256 amount) returns()
func (_ERC20TokenRemoteUpgradeable *ERC20TokenRemoteUpgradeableTransactor) Send(opts *bind.TransactOpts, input SendTokensInput, amount *big.Int) (*types.Transaction, error) {
	return _ERC20TokenRemoteUpgradeable.contract.Transact(opts, "send", input, amount)
}

// Send is a paid mutator transaction binding the contract method 0x5d16225d.
//
// Solidity: function send((bytes32,address,address,address,uint256,uint256,uint256,address) input, uint256 amount) returns()
func (_ERC20TokenRemoteUpgradeable *ERC20TokenRemoteUpgradeableSession) Send(input SendTokensInput, amount *big.Int) (*types.Transaction, error) {
	return _ERC20TokenRemoteUpgradeable.Contract.Send(&_ERC20TokenRemoteUpgradeable.TransactOpts, input, amount)
}

// Send is a paid mutator transaction binding the contract method 0x5d16225d.
//
// Solidity: function send((bytes32,address,address,address,uint256,uint256,uint256,address) input, uint256 amount) returns()
func (_ERC20TokenRemoteUpgradeable *ERC20TokenRemoteUpgradeableTransactorSession) Send(input SendTokensInput, amount *big.Int) (*types.Transaction, error) {
	return _ERC20TokenRemoteUpgradeable.Contract.Send(&_ERC20TokenRemoteUpgradeable.TransactOpts, input, amount)
}

// SendAndCall is a paid mutator transaction binding the contract method 0x65690038.
//
// Solidity: function sendAndCall((bytes32,address,address,bytes,uint256,uint256,address,address,address,uint256,uint256) input, uint256 amount) returns()
func (_ERC20TokenRemoteUpgradeable *ERC20TokenRemoteUpgradeableTransactor) SendAndCall(opts *bind.TransactOpts, input SendAndCallInput, amount *big.Int) (*types.Transaction, error) {
	return _ERC20TokenRemoteUpgradeable.contract.Transact(opts, "sendAndCall", input, amount)
}

// SendAndCall is a paid mutator transaction binding the contract method 0x65690038.
//
// Solidity: function sendAndCall((bytes32,address,address,bytes,uint256,uint256,address,address,address,uint256,uint256) input, uint256 amount) returns()
func (_ERC20TokenRemoteUpgradeable *ERC20TokenRemoteUpgradeableSession) SendAndCall(input SendAndCallInput, amount *big.Int) (*types.Transaction, error) {
	return _ERC20TokenRemoteUpgradeable.Contract.SendAndCall(&_ERC20TokenRemoteUpgradeable.TransactOpts, input, amount)
}

// SendAndCall is a paid mutator transaction binding the contract method 0x65690038.
//
// Solidity: function sendAndCall((bytes32,address,address,bytes,uint256,uint256,address,address,address,uint256,uint256) input, uint256 amount) returns()
func (_ERC20TokenRemoteUpgradeable *ERC20TokenRemoteUpgradeableTransactorSession) SendAndCall(input SendAndCallInput, amount *big.Int) (*types.Transaction, error) {
	return _ERC20TokenRemoteUpgradeable.Contract.SendAndCall(&_ERC20TokenRemoteUpgradeable.TransactOpts, input, amount)
}

// SetOutboundMessenger is a paid mutator transaction binding the contract method 0xd90f1b10.
//
// Solidity: function setOutboundMessenger(bytes32 remoteBlockchainID, address messenger) returns()
func (_ERC20TokenRemoteUpgradeable *ERC20TokenRemoteUpgradeableTransactor) SetOutboundMessenger(opts *bind.TransactOpts, remoteBlockchainID [32]byte, messenger common.Address) (*types.Transaction, error) {
	return _ERC20TokenRemoteUpgradeable.contract.Transact(opts, "setOutboundMessenger", remoteBlockchainID, messenger)
}

// SetOutboundMessenger is a paid mutator transaction binding the contract method 0xd90f1b10.
//
// Solidity: function setOutboundMessenger(bytes32 remoteBlockchainID, address messenger) returns()
func (_ERC20TokenRemoteUpgradeable *ERC20TokenRemoteUpgradeableSession) SetOutboundMessenger(remoteBlockchainID [32]byte, messenger common.Address) (*types.Transaction, error) {
	return _ERC20TokenRemoteUpgradeable.Contract.SetOutboundMessenger(&_ERC20TokenRemoteUpgradeable.TransactOpts, remoteBlockchainID, messenger)
}

// SetOutboundMessenger is a paid mutator transaction binding the contract method 0xd90f1b10.
//
// Solidity: function setOutboundMessenger(bytes32 remoteBlockchainID, address messenger) returns()
func (_ERC20TokenRemoteUpgradeable *ERC20TokenRemoteUpgradeableTransactorSession) SetOutboundMessenger(remoteBlockchainID [32]byte, messenger common.Address) (*types.Transaction, error) {
	return _ERC20TokenRemoteUpgradeable.Contract.SetOutboundMessenger(&_ERC20TokenRemoteUpgradeable.TransactOpts, remoteBlockchainID, messenger)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 value) returns(bool)
func (_ERC20TokenRemoteUpgradeable *ERC20TokenRemoteUpgradeableTransactor) Transfer(opts *bind.TransactOpts, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _ERC20TokenRemoteUpgradeable.contract.Transact(opts, "transfer", to, value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 value) returns(bool)
func (_ERC20TokenRemoteUpgradeable *ERC20TokenRemoteUpgradeableSession) Transfer(to common.Address, value *big.Int) (*types.Transaction, error) {
	return _ERC20TokenRemoteUpgradeable.Contract.Transfer(&_ERC20TokenRemoteUpgradeable.TransactOpts, to, value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 value) returns(bool)
func (_ERC20TokenRemoteUpgradeable *ERC20TokenRemoteUpgradeableTransactorSession) Transfer(to common.Address, value *big.Int) (*types.Transaction, error) {
	return _ERC20TokenRemoteUpgradeable.Contract.Transfer(&_ERC20TokenRemoteUpgradeable.TransactOpts, to, value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 value) returns(bool)
func (_ERC20TokenRemoteUpgradeable *ERC20TokenRemoteUpgradeableTransactor) TransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _ERC20TokenRemoteUpgradeable.contract.Transact(opts, "transferFrom", from, to, value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 value) returns(bool)
func (_ERC20TokenRemoteUpgradeable *ERC20TokenRemoteUpgradeableSession) TransferFrom(from common.Address, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _ERC20TokenRemoteUpgradeable.Contract.TransferFrom(&_ERC20TokenRemoteUpgradeable.TransactOpts, from, to, value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 value) returns(bool)
func (_ERC20TokenRemoteUpgradeable *ERC20TokenRemoteUpgradeableTransactorSession) TransferFrom(from common.Address, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _ERC20TokenRemoteUpgradeable.Contract.TransferFrom(&_ERC20TokenRemoteUpgradeable.TransactOpts, from, to, value)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_ERC20TokenRemoteUpgradeable *ERC20TokenRemoteUpgradeableTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _ERC20TokenRemoteUpgradeable.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_ERC20TokenRemoteUpgradeable *ERC20TokenRemoteUpgradeableSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _ERC20TokenRemoteUpgradeable.Contract.TransferOwnership(&_ERC20TokenRemoteUpgradeable.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_ERC20TokenRemoteUpgradeable *ERC20TokenRemoteUpgradeableTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _ERC20TokenRemoteUpgradeable.Contract.TransferOwnership(&_ERC20TokenRemoteUpgradeable.TransactOpts, newOwner)
}

// UnpauseMessenger is a paid mutator transaction binding the contract method 0xdd5a20fd.
//
// Solidity: function unpauseMessenger(address messenger) returns()
func (_ERC20TokenRemoteUpgradeable *ERC20TokenRemoteUpgradeableTransactor) UnpauseMessenger(opts *bind.TransactOpts, messenger common.Address) (*types.Transaction, error) {
	return _ERC20TokenRemoteUpgradeable.contract.Transact(opts, "unpauseMessenger", messenger)
}

// UnpauseMessenger is a paid mutator transaction binding the contract method 0xdd5a20fd.
//
// Solidity: function unpauseMessenger(address messenger) returns()
func (_ERC20TokenRemoteUpgradeable *ERC20TokenRemoteUpgradeableSession) UnpauseMessenger(messenger common.Address) (*types.Transaction, error) {
	return _ERC20TokenRemoteUpgradeable.Contract.UnpauseMessenger(&_ERC20TokenRemoteUpgradeable.TransactOpts, messenger)
}

// UnpauseMessenger is a paid mutator transaction binding the contract method 0xdd5a20fd.
//
// Solidity: function unpauseMessenger(address messenger) returns()
func (_ERC20TokenRemoteUpgradeable *ERC20TokenRemoteUpgradeableTransactorSession) UnpauseMessenger(messenger common.Address) (*types.Transaction, error) {
	return _ERC20TokenRemoteUpgradeable.Contract.UnpauseMessenger(&_ERC20TokenRemoteUpgradeable.TransactOpts, messenger)
}

// ERC20TokenRemoteUpgradeableApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the ERC20TokenRemoteUpgradeable contract.
type ERC20TokenRemoteUpgradeableApprovalIterator struct {
	Event *ERC20TokenRemoteUpgradeableApproval // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ERC20TokenRemoteUpgradeableApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ERC20TokenRemoteUpgradeableApproval)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ERC20TokenRemoteUpgradeableApproval)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ERC20TokenRemoteUpgradeableApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ERC20TokenRemoteUpgradeableApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ERC20TokenRemoteUpgradeableApproval represents a Approval event raised by the ERC20TokenRemoteUpgradeable contract.
type ERC20TokenRemoteUpgradeableApproval struct {
	Owner   common.Address
	Spender common.Address
	Value   *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_ERC20TokenRemoteUpgradeable *ERC20TokenRemoteUpgradeableFilterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, spender []common.Address) (*ERC20TokenRemoteUpgradeableApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _ERC20TokenRemoteUpgradeable.contract.FilterLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return &ERC20TokenRemoteUpgradeableApprovalIterator{contract: _ERC20TokenRemoteUpgradeable.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_ERC20TokenRemoteUpgradeable *ERC20TokenRemoteUpgradeableFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *ERC20TokenRemoteUpgradeableApproval, owner []common.Address, spender []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _ERC20TokenRemoteUpgradeable.contract.WatchLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ERC20TokenRemoteUpgradeableApproval)
				if err := _ERC20TokenRemoteUpgradeable.contract.UnpackLog(event, "Approval", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseApproval is a log parse operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_ERC20TokenRemoteUpgradeable *ERC20TokenRemoteUpgradeableFilterer) ParseApproval(log types.Log) (*ERC20TokenRemoteUpgradeableApproval, error) {
	event := new(ERC20TokenRemoteUpgradeableApproval)
	if err := _ERC20TokenRemoteUpgradeable.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ERC20TokenRemoteUpgradeableCallFailedIterator is returned from FilterCallFailed and is used to iterate over the raw logs and unpacked data for CallFailed events raised by the ERC20TokenRemoteUpgradeable contract.
type ERC20TokenRemoteUpgradeableCallFailedIterator struct {
	Event *ERC20TokenRemoteUpgradeableCallFailed // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ERC20TokenRemoteUpgradeableCallFailedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ERC20TokenRemoteUpgradeableCallFailed)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ERC20TokenRemoteUpgradeableCallFailed)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ERC20TokenRemoteUpgradeableCallFailedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ERC20TokenRemoteUpgradeableCallFailedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ERC20TokenRemoteUpgradeableCallFailed represents a CallFailed event raised by the ERC20TokenRemoteUpgradeable contract.
type ERC20TokenRemoteUpgradeableCallFailed struct {
	RecipientContract common.Address
	Amount            *big.Int
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterCallFailed is a free log retrieval operation binding the contract event 0xb9eaeae386d339f8115782f297a9e5f0e13fb587cd6b0d502f113cb8dd4d6cb0.
//
// Solidity: event CallFailed(address indexed recipientContract, uint256 amount)
func (_ERC20TokenRemoteUpgradeable *ERC20TokenRemoteUpgradeableFilterer) FilterCallFailed(opts *bind.FilterOpts, recipientContract []common.Address) (*ERC20TokenRemoteUpgradeableCallFailedIterator, error) {

	var recipientContractRule []interface{}
	for _, recipientContractItem := range recipientContract {
		recipientContractRule = append(recipientContractRule, recipientContractItem)
	}

	logs, sub, err := _ERC20TokenRemoteUpgradeable.contract.FilterLogs(opts, "CallFailed", recipientContractRule)
	if err != nil {
		return nil, err
	}
	return &ERC20TokenRemoteUpgradeableCallFailedIterator{contract: _ERC20TokenRemoteUpgradeable.contract, event: "CallFailed", logs: logs, sub: sub}, nil
}

// WatchCallFailed is a free log subscription operation binding the contract event 0xb9eaeae386d339f8115782f297a9e5f0e13fb587cd6b0d502f113cb8dd4d6cb0.
//
// Solidity: event CallFailed(address indexed recipientContract, uint256 amount)
func (_ERC20TokenRemoteUpgradeable *ERC20TokenRemoteUpgradeableFilterer) WatchCallFailed(opts *bind.WatchOpts, sink chan<- *ERC20TokenRemoteUpgradeableCallFailed, recipientContract []common.Address) (event.Subscription, error) {

	var recipientContractRule []interface{}
	for _, recipientContractItem := range recipientContract {
		recipientContractRule = append(recipientContractRule, recipientContractItem)
	}

	logs, sub, err := _ERC20TokenRemoteUpgradeable.contract.WatchLogs(opts, "CallFailed", recipientContractRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ERC20TokenRemoteUpgradeableCallFailed)
				if err := _ERC20TokenRemoteUpgradeable.contract.UnpackLog(event, "CallFailed", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseCallFailed is a log parse operation binding the contract event 0xb9eaeae386d339f8115782f297a9e5f0e13fb587cd6b0d502f113cb8dd4d6cb0.
//
// Solidity: event CallFailed(address indexed recipientContract, uint256 amount)
func (_ERC20TokenRemoteUpgradeable *ERC20TokenRemoteUpgradeableFilterer) ParseCallFailed(log types.Log) (*ERC20TokenRemoteUpgradeableCallFailed, error) {
	event := new(ERC20TokenRemoteUpgradeableCallFailed)
	if err := _ERC20TokenRemoteUpgradeable.contract.UnpackLog(event, "CallFailed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ERC20TokenRemoteUpgradeableCallSucceededIterator is returned from FilterCallSucceeded and is used to iterate over the raw logs and unpacked data for CallSucceeded events raised by the ERC20TokenRemoteUpgradeable contract.
type ERC20TokenRemoteUpgradeableCallSucceededIterator struct {
	Event *ERC20TokenRemoteUpgradeableCallSucceeded // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ERC20TokenRemoteUpgradeableCallSucceededIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ERC20TokenRemoteUpgradeableCallSucceeded)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ERC20TokenRemoteUpgradeableCallSucceeded)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ERC20TokenRemoteUpgradeableCallSucceededIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ERC20TokenRemoteUpgradeableCallSucceededIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ERC20TokenRemoteUpgradeableCallSucceeded represents a CallSucceeded event raised by the ERC20TokenRemoteUpgradeable contract.
type ERC20TokenRemoteUpgradeableCallSucceeded struct {
	RecipientContract common.Address
	Amount            *big.Int
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterCallSucceeded is a free log retrieval operation binding the contract event 0x104deb555f67e63782bb817bc26c39050894645f9b9f29c4be8ae68d0e8b7ff4.
//
// Solidity: event CallSucceeded(address indexed recipientContract, uint256 amount)
func (_ERC20TokenRemoteUpgradeable *ERC20TokenRemoteUpgradeableFilterer) FilterCallSucceeded(opts *bind.FilterOpts, recipientContract []common.Address) (*ERC20TokenRemoteUpgradeableCallSucceededIterator, error) {

	var recipientContractRule []interface{}
	for _, recipientContractItem := range recipientContract {
		recipientContractRule = append(recipientContractRule, recipientContractItem)
	}

	logs, sub, err := _ERC20TokenRemoteUpgradeable.contract.FilterLogs(opts, "CallSucceeded", recipientContractRule)
	if err != nil {
		return nil, err
	}
	return &ERC20TokenRemoteUpgradeableCallSucceededIterator{contract: _ERC20TokenRemoteUpgradeable.contract, event: "CallSucceeded", logs: logs, sub: sub}, nil
}

// WatchCallSucceeded is a free log subscription operation binding the contract event 0x104deb555f67e63782bb817bc26c39050894645f9b9f29c4be8ae68d0e8b7ff4.
//
// Solidity: event CallSucceeded(address indexed recipientContract, uint256 amount)
func (_ERC20TokenRemoteUpgradeable *ERC20TokenRemoteUpgradeableFilterer) WatchCallSucceeded(opts *bind.WatchOpts, sink chan<- *ERC20TokenRemoteUpgradeableCallSucceeded, recipientContract []common.Address) (event.Subscription, error) {

	var recipientContractRule []interface{}
	for _, recipientContractItem := range recipientContract {
		recipientContractRule = append(recipientContractRule, recipientContractItem)
	}

	logs, sub, err := _ERC20TokenRemoteUpgradeable.contract.WatchLogs(opts, "CallSucceeded", recipientContractRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ERC20TokenRemoteUpgradeableCallSucceeded)
				if err := _ERC20TokenRemoteUpgradeable.contract.UnpackLog(event, "CallSucceeded", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseCallSucceeded is a log parse operation binding the contract event 0x104deb555f67e63782bb817bc26c39050894645f9b9f29c4be8ae68d0e8b7ff4.
//
// Solidity: event CallSucceeded(address indexed recipientContract, uint256 amount)
func (_ERC20TokenRemoteUpgradeable *ERC20TokenRemoteUpgradeableFilterer) ParseCallSucceeded(log types.Log) (*ERC20TokenRemoteUpgradeableCallSucceeded, error) {
	event := new(ERC20TokenRemoteUpgradeableCallSucceeded)
	if err := _ERC20TokenRemoteUpgradeable.contract.UnpackLog(event, "CallSucceeded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ERC20TokenRemoteUpgradeableInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the ERC20TokenRemoteUpgradeable contract.
type ERC20TokenRemoteUpgradeableInitializedIterator struct {
	Event *ERC20TokenRemoteUpgradeableInitialized // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ERC20TokenRemoteUpgradeableInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ERC20TokenRemoteUpgradeableInitialized)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ERC20TokenRemoteUpgradeableInitialized)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ERC20TokenRemoteUpgradeableInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ERC20TokenRemoteUpgradeableInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ERC20TokenRemoteUpgradeableInitialized represents a Initialized event raised by the ERC20TokenRemoteUpgradeable contract.
type ERC20TokenRemoteUpgradeableInitialized struct {
	Version uint64
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_ERC20TokenRemoteUpgradeable *ERC20TokenRemoteUpgradeableFilterer) FilterInitialized(opts *bind.FilterOpts) (*ERC20TokenRemoteUpgradeableInitializedIterator, error) {

	logs, sub, err := _ERC20TokenRemoteUpgradeable.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &ERC20TokenRemoteUpgradeableInitializedIterator{contract: _ERC20TokenRemoteUpgradeable.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_ERC20TokenRemoteUpgradeable *ERC20TokenRemoteUpgradeableFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *ERC20TokenRemoteUpgradeableInitialized) (event.Subscription, error) {

	logs, sub, err := _ERC20TokenRemoteUpgradeable.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ERC20TokenRemoteUpgradeableInitialized)
				if err := _ERC20TokenRemoteUpgradeable.contract.UnpackLog(event, "Initialized", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseInitialized is a log parse operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_ERC20TokenRemoteUpgradeable *ERC20TokenRemoteUpgradeableFilterer) ParseInitialized(log types.Log) (*ERC20TokenRemoteUpgradeableInitialized, error) {
	event := new(ERC20TokenRemoteUpgradeableInitialized)
	if err := _ERC20TokenRemoteUpgradeable.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ERC20TokenRemoteUpgradeableMessengerAddedIterator is returned from FilterMessengerAdded and is used to iterate over the raw logs and unpacked data for MessengerAdded events raised by the ERC20TokenRemoteUpgradeable contract.
type ERC20TokenRemoteUpgradeableMessengerAddedIterator struct {
	Event *ERC20TokenRemoteUpgradeableMessengerAdded // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ERC20TokenRemoteUpgradeableMessengerAddedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ERC20TokenRemoteUpgradeableMessengerAdded)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ERC20TokenRemoteUpgradeableMessengerAdded)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ERC20TokenRemoteUpgradeableMessengerAddedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ERC20TokenRemoteUpgradeableMessengerAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ERC20TokenRemoteUpgradeableMessengerAdded represents a MessengerAdded event raised by the ERC20TokenRemoteUpgradeable contract.
type ERC20TokenRemoteUpgradeableMessengerAdded struct {
	RemoteBlockchainID [32]byte
	Messenger          common.Address
	Raw                types.Log // Blockchain specific contextual infos
}

// FilterMessengerAdded is a free log retrieval operation binding the contract event 0xcc2f89e16e8979d86d3c3b49c5646f3a013f76626a479400e1cb3596f3dc2b22.
//
// Solidity: event MessengerAdded(bytes32 indexed remoteBlockchainID, address indexed messenger)
func (_ERC20TokenRemoteUpgradeable *ERC20TokenRemoteUpgradeableFilterer) FilterMessengerAdded(opts *bind.FilterOpts, remoteBlockchainID [][32]byte, messenger []common.Address) (*ERC20TokenRemoteUpgradeableMessengerAddedIterator, error) {

	var remoteBlockchainIDRule []interface{}
	for _, remoteBlockchainIDItem := range remoteBlockchainID {
		remoteBlockchainIDRule = append(remoteBlockchainIDRule, remoteBlockchainIDItem)
	}
	var messengerRule []interface{}
	for _, messengerItem := range messenger {
		messengerRule = append(messengerRule, messengerItem)
	}

	logs, sub, err := _ERC20TokenRemoteUpgradeable.contract.FilterLogs(opts, "MessengerAdded", remoteBlockchainIDRule, messengerRule)
	if err != nil {
		return nil, err
	}
	return &ERC20TokenRemoteUpgradeableMessengerAddedIterator{contract: _ERC20TokenRemoteUpgradeable.contract, event: "MessengerAdded", logs: logs, sub: sub}, nil
}

// WatchMessengerAdded is a free log subscription operation binding the contract event 0xcc2f89e16e8979d86d3c3b49c5646f3a013f76626a479400e1cb3596f3dc2b22.
//
// Solidity: event MessengerAdded(bytes32 indexed remoteBlockchainID, address indexed messenger)
func (_ERC20TokenRemoteUpgradeable *ERC20TokenRemoteUpgradeableFilterer) WatchMessengerAdded(opts *bind.WatchOpts, sink chan<- *ERC20TokenRemoteUpgradeableMessengerAdded, remoteBlockchainID [][32]byte, messenger []common.Address) (event.Subscription, error) {

	var remoteBlockchainIDRule []interface{}
	for _, remoteBlockchainIDItem := range remoteBlockchainID {
		remoteBlockchainIDRule = append(remoteBlockchainIDRule, remoteBlockchainIDItem)
	}
	var messengerRule []interface{}
	for _, messengerItem := range messenger {
		messengerRule = append(messengerRule, messengerItem)
	}

	logs, sub, err := _ERC20TokenRemoteUpgradeable.contract.WatchLogs(opts, "MessengerAdded", remoteBlockchainIDRule, messengerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ERC20TokenRemoteUpgradeableMessengerAdded)
				if err := _ERC20TokenRemoteUpgradeable.contract.UnpackLog(event, "MessengerAdded", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseMessengerAdded is a log parse operation binding the contract event 0xcc2f89e16e8979d86d3c3b49c5646f3a013f76626a479400e1cb3596f3dc2b22.
//
// Solidity: event MessengerAdded(bytes32 indexed remoteBlockchainID, address indexed messenger)
func (_ERC20TokenRemoteUpgradeable *ERC20TokenRemoteUpgradeableFilterer) ParseMessengerAdded(log types.Log) (*ERC20TokenRemoteUpgradeableMessengerAdded, error) {
	event := new(ERC20TokenRemoteUpgradeableMessengerAdded)
	if err := _ERC20TokenRemoteUpgradeable.contract.UnpackLog(event, "MessengerAdded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ERC20TokenRemoteUpgradeableMessengerPausedIterator is returned from FilterMessengerPaused and is used to iterate over the raw logs and unpacked data for MessengerPaused events raised by the ERC20TokenRemoteUpgradeable contract.
type ERC20TokenRemoteUpgradeableMessengerPausedIterator struct {
	Event *ERC20TokenRemoteUpgradeableMessengerPaused // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ERC20TokenRemoteUpgradeableMessengerPausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ERC20TokenRemoteUpgradeableMessengerPaused)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ERC20TokenRemoteUpgradeableMessengerPaused)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ERC20TokenRemoteUpgradeableMessengerPausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ERC20TokenRemoteUpgradeableMessengerPausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ERC20TokenRemoteUpgradeableMessengerPaused represents a MessengerPaused event raised by the ERC20TokenRemoteUpgradeable contract.
type ERC20TokenRemoteUpgradeableMessengerPaused struct {
	Messenger common.Address
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterMessengerPaused is a free log retrieval operation binding the contract event 0x46def370ef3a1ec72618480f66f1063a80fe68b27d09575bc3562ad489079b61.
//
// Solidity: event MessengerPaused(address indexed messenger)
func (_ERC20TokenRemoteUpgradeable *ERC20TokenRemoteUpgradeableFilterer) FilterMessengerPaused(opts *bind.FilterOpts, messenger []common.Address) (*ERC20TokenRemoteUpgradeableMessengerPausedIterator, error) {

	var messengerRule []interface{}
	for _, messengerItem := range messenger {
		messengerRule = append(messengerRule, messengerItem)
	}

	logs, sub, err := _ERC20TokenRemoteUpgradeable.contract.FilterLogs(opts, "MessengerPaused", messengerRule)
	if err != nil {
		return nil, err
	}
	return &ERC20TokenRemoteUpgradeableMessengerPausedIterator{contract: _ERC20TokenRemoteUpgradeable.contract, event: "MessengerPaused", logs: logs, sub: sub}, nil
}

// WatchMessengerPaused is a free log subscription operation binding the contract event 0x46def370ef3a1ec72618480f66f1063a80fe68b27d09575bc3562ad489079b61.
//
// Solidity: event MessengerPaused(address indexed messenger)
func (_ERC20TokenRemoteUpgradeable *ERC20TokenRemoteUpgradeableFilterer) WatchMessengerPaused(opts *bind.WatchOpts, sink chan<- *ERC20TokenRemoteUpgradeableMessengerPaused, messenger []common.Address) (event.Subscription, error) {

	var messengerRule []interface{}
	for _, messengerItem := range messenger {
		messengerRule = append(messengerRule, messengerItem)
	}

	logs, sub, err := _ERC20TokenRemoteUpgradeable.contract.WatchLogs(opts, "MessengerPaused", messengerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ERC20TokenRemoteUpgradeableMessengerPaused)
				if err := _ERC20TokenRemoteUpgradeable.contract.UnpackLog(event, "MessengerPaused", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseMessengerPaused is a log parse operation binding the contract event 0x46def370ef3a1ec72618480f66f1063a80fe68b27d09575bc3562ad489079b61.
//
// Solidity: event MessengerPaused(address indexed messenger)
func (_ERC20TokenRemoteUpgradeable *ERC20TokenRemoteUpgradeableFilterer) ParseMessengerPaused(log types.Log) (*ERC20TokenRemoteUpgradeableMessengerPaused, error) {
	event := new(ERC20TokenRemoteUpgradeableMessengerPaused)
	if err := _ERC20TokenRemoteUpgradeable.contract.UnpackLog(event, "MessengerPaused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ERC20TokenRemoteUpgradeableMessengerRemovedIterator is returned from FilterMessengerRemoved and is used to iterate over the raw logs and unpacked data for MessengerRemoved events raised by the ERC20TokenRemoteUpgradeable contract.
type ERC20TokenRemoteUpgradeableMessengerRemovedIterator struct {
	Event *ERC20TokenRemoteUpgradeableMessengerRemoved // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ERC20TokenRemoteUpgradeableMessengerRemovedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ERC20TokenRemoteUpgradeableMessengerRemoved)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ERC20TokenRemoteUpgradeableMessengerRemoved)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ERC20TokenRemoteUpgradeableMessengerRemovedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ERC20TokenRemoteUpgradeableMessengerRemovedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ERC20TokenRemoteUpgradeableMessengerRemoved represents a MessengerRemoved event raised by the ERC20TokenRemoteUpgradeable contract.
type ERC20TokenRemoteUpgradeableMessengerRemoved struct {
	RemoteBlockchainID [32]byte
	Messenger          common.Address
	Raw                types.Log // Blockchain specific contextual infos
}

// FilterMessengerRemoved is a free log retrieval operation binding the contract event 0x8350b16da8b8ec7d7e6ce37ca2ccb4f63d628b5ad53215d072763b6afce99524.
//
// Solidity: event MessengerRemoved(bytes32 indexed remoteBlockchainID, address indexed messenger)
func (_ERC20TokenRemoteUpgradeable *ERC20TokenRemoteUpgradeableFilterer) FilterMessengerRemoved(opts *bind.FilterOpts, remoteBlockchainID [][32]byte, messenger []common.Address) (*ERC20TokenRemoteUpgradeableMessengerRemovedIterator, error) {

	var remoteBlockchainIDRule []interface{}
	for _, remoteBlockchainIDItem := range remoteBlockchainID {
		remoteBlockchainIDRule = append(remoteBlockchainIDRule, remoteBlockchainIDItem)
	}
	var messengerRule []interface{}
	for _, messengerItem := range messenger {
		messengerRule = append(messengerRule, messengerItem)
	}

	logs, sub, err := _ERC20TokenRemoteUpgradeable.contract.FilterLogs(opts, "MessengerRemoved", remoteBlockchainIDRule, messengerRule)
	if err != nil {
		return nil, err
	}
	return &ERC20TokenRemoteUpgradeableMessengerRemovedIterator{contract: _ERC20TokenRemoteUpgradeable.contract, event: "MessengerRemoved", logs: logs, sub: sub}, nil
}

// WatchMessengerRemoved is a free log subscription operation binding the contract event 0x8350b16da8b8ec7d7e6ce37ca2ccb4f63d628b5ad53215d072763b6afce99524.
//
// Solidity: event MessengerRemoved(bytes32 indexed remoteBlockchainID, address indexed messenger)
func (_ERC20TokenRemoteUpgradeable *ERC20TokenRemoteUpgradeableFilterer) WatchMessengerRemoved(opts *bind.WatchOpts, sink chan<- *ERC20TokenRemoteUpgradeableMessengerRemoved, remoteBlockchainID [][32]byte, messenger []common.Address) (event.Subscription, error) {

	var remoteBlockchainIDRule []interface{}
	for _, remoteBlockchainIDItem := range remoteBlockchainID {
		remoteBlockchainIDRule = append(remoteBlockchainIDRule, remoteBlockchainIDItem)
	}
	var messengerRule []interface{}
	for _, messengerItem := range messenger {
		messengerRule = append(messengerRule, messengerItem)
	}

	logs, sub, err := _ERC20TokenRemoteUpgradeable.contract.WatchLogs(opts, "MessengerRemoved", remoteBlockchainIDRule, messengerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ERC20TokenRemoteUpgradeableMessengerRemoved)
				if err := _ERC20TokenRemoteUpgradeable.contract.UnpackLog(event, "MessengerRemoved", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseMessengerRemoved is a log parse operation binding the contract event 0x8350b16da8b8ec7d7e6ce37ca2ccb4f63d628b5ad53215d072763b6afce99524.
//
// Solidity: event MessengerRemoved(bytes32 indexed remoteBlockchainID, address indexed messenger)
func (_ERC20TokenRemoteUpgradeable *ERC20TokenRemoteUpgradeableFilterer) ParseMessengerRemoved(log types.Log) (*ERC20TokenRemoteUpgradeableMessengerRemoved, error) {
	event := new(ERC20TokenRemoteUpgradeableMessengerRemoved)
	if err := _ERC20TokenRemoteUpgradeable.contract.UnpackLog(event, "MessengerRemoved", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ERC20TokenRemoteUpgradeableMessengerUnpausedIterator is returned from FilterMessengerUnpaused and is used to iterate over the raw logs and unpacked data for MessengerUnpaused events raised by the ERC20TokenRemoteUpgradeable contract.
type ERC20TokenRemoteUpgradeableMessengerUnpausedIterator struct {
	Event *ERC20TokenRemoteUpgradeableMessengerUnpaused // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ERC20TokenRemoteUpgradeableMessengerUnpausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ERC20TokenRemoteUpgradeableMessengerUnpaused)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ERC20TokenRemoteUpgradeableMessengerUnpaused)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ERC20TokenRemoteUpgradeableMessengerUnpausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ERC20TokenRemoteUpgradeableMessengerUnpausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ERC20TokenRemoteUpgradeableMessengerUnpaused represents a MessengerUnpaused event raised by the ERC20TokenRemoteUpgradeable contract.
type ERC20TokenRemoteUpgradeableMessengerUnpaused struct {
	Messenger common.Address
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterMessengerUnpaused is a free log retrieval operation binding the contract event 0xf5574433a9d82924af98e239704e51afb34ae8c15f513d78a22ae0c58f5f9965.
//
// Solidity: event MessengerUnpaused(address indexed messenger)
func (_ERC20TokenRemoteUpgradeable *ERC20TokenRemoteUpgradeableFilterer) FilterMessengerUnpaused(opts *bind.FilterOpts, messenger []common.Address) (*ERC20TokenRemoteUpgradeableMessengerUnpausedIterator, error) {

	var messengerRule []interface{}
	for _, messengerItem := range messenger {
		messengerRule = append(messengerRule, messengerItem)
	}

	logs, sub, err := _ERC20TokenRemoteUpgradeable.contract.FilterLogs(opts, "MessengerUnpaused", messengerRule)
	if err != nil {
		return nil, err
	}
	return &ERC20TokenRemoteUpgradeableMessengerUnpausedIterator{contract: _ERC20TokenRemoteUpgradeable.contract, event: "MessengerUnpaused", logs: logs, sub: sub}, nil
}

// WatchMessengerUnpaused is a free log subscription operation binding the contract event 0xf5574433a9d82924af98e239704e51afb34ae8c15f513d78a22ae0c58f5f9965.
//
// Solidity: event MessengerUnpaused(address indexed messenger)
func (_ERC20TokenRemoteUpgradeable *ERC20TokenRemoteUpgradeableFilterer) WatchMessengerUnpaused(opts *bind.WatchOpts, sink chan<- *ERC20TokenRemoteUpgradeableMessengerUnpaused, messenger []common.Address) (event.Subscription, error) {

	var messengerRule []interface{}
	for _, messengerItem := range messenger {
		messengerRule = append(messengerRule, messengerItem)
	}

	logs, sub, err := _ERC20TokenRemoteUpgradeable.contract.WatchLogs(opts, "MessengerUnpaused", messengerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ERC20TokenRemoteUpgradeableMessengerUnpaused)
				if err := _ERC20TokenRemoteUpgradeable.contract.UnpackLog(event, "MessengerUnpaused", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseMessengerUnpaused is a log parse operation binding the contract event 0xf5574433a9d82924af98e239704e51afb34ae8c15f513d78a22ae0c58f5f9965.
//
// Solidity: event MessengerUnpaused(address indexed messenger)
func (_ERC20TokenRemoteUpgradeable *ERC20TokenRemoteUpgradeableFilterer) ParseMessengerUnpaused(log types.Log) (*ERC20TokenRemoteUpgradeableMessengerUnpaused, error) {
	event := new(ERC20TokenRemoteUpgradeableMessengerUnpaused)
	if err := _ERC20TokenRemoteUpgradeable.contract.UnpackLog(event, "MessengerUnpaused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ERC20TokenRemoteUpgradeableOutboundMessengerSetIterator is returned from FilterOutboundMessengerSet and is used to iterate over the raw logs and unpacked data for OutboundMessengerSet events raised by the ERC20TokenRemoteUpgradeable contract.
type ERC20TokenRemoteUpgradeableOutboundMessengerSetIterator struct {
	Event *ERC20TokenRemoteUpgradeableOutboundMessengerSet // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ERC20TokenRemoteUpgradeableOutboundMessengerSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ERC20TokenRemoteUpgradeableOutboundMessengerSet)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ERC20TokenRemoteUpgradeableOutboundMessengerSet)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ERC20TokenRemoteUpgradeableOutboundMessengerSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ERC20TokenRemoteUpgradeableOutboundMessengerSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ERC20TokenRemoteUpgradeableOutboundMessengerSet represents a OutboundMessengerSet event raised by the ERC20TokenRemoteUpgradeable contract.
type ERC20TokenRemoteUpgradeableOutboundMessengerSet struct {
	RemoteBlockchainID [32]byte
	OldMessenger       common.Address
	NewMessenger       common.Address
	Raw                types.Log // Blockchain specific contextual infos
}

// FilterOutboundMessengerSet is a free log retrieval operation binding the contract event 0xd00498d0475cde985344ea08d024c9315ca654d9d157211295aac43496374115.
//
// Solidity: event OutboundMessengerSet(bytes32 indexed remoteBlockchainID, address indexed oldMessenger, address indexed newMessenger)
func (_ERC20TokenRemoteUpgradeable *ERC20TokenRemoteUpgradeableFilterer) FilterOutboundMessengerSet(opts *bind.FilterOpts, remoteBlockchainID [][32]byte, oldMessenger []common.Address, newMessenger []common.Address) (*ERC20TokenRemoteUpgradeableOutboundMessengerSetIterator, error) {

	var remoteBlockchainIDRule []interface{}
	for _, remoteBlockchainIDItem := range remoteBlockchainID {
		remoteBlockchainIDRule = append(remoteBlockchainIDRule, remoteBlockchainIDItem)
	}
	var oldMessengerRule []interface{}
	for _, oldMessengerItem := range oldMessenger {
		oldMessengerRule = append(oldMessengerRule, oldMessengerItem)
	}
	var newMessengerRule []interface{}
	for _, newMessengerItem := range newMessenger {
		newMessengerRule = append(newMessengerRule, newMessengerItem)
	}

	logs, sub, err := _ERC20TokenRemoteUpgradeable.contract.FilterLogs(opts, "OutboundMessengerSet", remoteBlockchainIDRule, oldMessengerRule, newMessengerRule)
	if err != nil {
		return nil, err
	}
	return &ERC20TokenRemoteUpgradeableOutboundMessengerSetIterator{contract: _ERC20TokenRemoteUpgradeable.contract, event: "OutboundMessengerSet", logs: logs, sub: sub}, nil
}

// WatchOutboundMessengerSet is a free log subscription operation binding the contract event 0xd00498d0475cde985344ea08d024c9315ca654d9d157211295aac43496374115.
//
// Solidity: event OutboundMessengerSet(bytes32 indexed remoteBlockchainID, address indexed oldMessenger, address indexed newMessenger)
func (_ERC20TokenRemoteUpgradeable *ERC20TokenRemoteUpgradeableFilterer) WatchOutboundMessengerSet(opts *bind.WatchOpts, sink chan<- *ERC20TokenRemoteUpgradeableOutboundMessengerSet, remoteBlockchainID [][32]byte, oldMessenger []common.Address, newMessenger []common.Address) (event.Subscription, error) {

	var remoteBlockchainIDRule []interface{}
	for _, remoteBlockchainIDItem := range remoteBlockchainID {
		remoteBlockchainIDRule = append(remoteBlockchainIDRule, remoteBlockchainIDItem)
	}
	var oldMessengerRule []interface{}
	for _, oldMessengerItem := range oldMessenger {
		oldMessengerRule = append(oldMessengerRule, oldMessengerItem)
	}
	var newMessengerRule []interface{}
	for _, newMessengerItem := range newMessenger {
		newMessengerRule = append(newMessengerRule, newMessengerItem)
	}

	logs, sub, err := _ERC20TokenRemoteUpgradeable.contract.WatchLogs(opts, "OutboundMessengerSet", remoteBlockchainIDRule, oldMessengerRule, newMessengerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ERC20TokenRemoteUpgradeableOutboundMessengerSet)
				if err := _ERC20TokenRemoteUpgradeable.contract.UnpackLog(event, "OutboundMessengerSet", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseOutboundMessengerSet is a log parse operation binding the contract event 0xd00498d0475cde985344ea08d024c9315ca654d9d157211295aac43496374115.
//
// Solidity: event OutboundMessengerSet(bytes32 indexed remoteBlockchainID, address indexed oldMessenger, address indexed newMessenger)
func (_ERC20TokenRemoteUpgradeable *ERC20TokenRemoteUpgradeableFilterer) ParseOutboundMessengerSet(log types.Log) (*ERC20TokenRemoteUpgradeableOutboundMessengerSet, error) {
	event := new(ERC20TokenRemoteUpgradeableOutboundMessengerSet)
	if err := _ERC20TokenRemoteUpgradeable.contract.UnpackLog(event, "OutboundMessengerSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ERC20TokenRemoteUpgradeableOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the ERC20TokenRemoteUpgradeable contract.
type ERC20TokenRemoteUpgradeableOwnershipTransferredIterator struct {
	Event *ERC20TokenRemoteUpgradeableOwnershipTransferred // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ERC20TokenRemoteUpgradeableOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ERC20TokenRemoteUpgradeableOwnershipTransferred)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ERC20TokenRemoteUpgradeableOwnershipTransferred)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ERC20TokenRemoteUpgradeableOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ERC20TokenRemoteUpgradeableOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ERC20TokenRemoteUpgradeableOwnershipTransferred represents a OwnershipTransferred event raised by the ERC20TokenRemoteUpgradeable contract.
type ERC20TokenRemoteUpgradeableOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_ERC20TokenRemoteUpgradeable *ERC20TokenRemoteUpgradeableFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*ERC20TokenRemoteUpgradeableOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _ERC20TokenRemoteUpgradeable.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &ERC20TokenRemoteUpgradeableOwnershipTransferredIterator{contract: _ERC20TokenRemoteUpgradeable.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_ERC20TokenRemoteUpgradeable *ERC20TokenRemoteUpgradeableFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *ERC20TokenRemoteUpgradeableOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _ERC20TokenRemoteUpgradeable.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ERC20TokenRemoteUpgradeableOwnershipTransferred)
				if err := _ERC20TokenRemoteUpgradeable.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseOwnershipTransferred is a log parse operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_ERC20TokenRemoteUpgradeable *ERC20TokenRemoteUpgradeableFilterer) ParseOwnershipTransferred(log types.Log) (*ERC20TokenRemoteUpgradeableOwnershipTransferred, error) {
	event := new(ERC20TokenRemoteUpgradeableOwnershipTransferred)
	if err := _ERC20TokenRemoteUpgradeable.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ERC20TokenRemoteUpgradeableTokensAndCallSentIterator is returned from FilterTokensAndCallSent and is used to iterate over the raw logs and unpacked data for TokensAndCallSent events raised by the ERC20TokenRemoteUpgradeable contract.
type ERC20TokenRemoteUpgradeableTokensAndCallSentIterator struct {
	Event *ERC20TokenRemoteUpgradeableTokensAndCallSent // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ERC20TokenRemoteUpgradeableTokensAndCallSentIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ERC20TokenRemoteUpgradeableTokensAndCallSent)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ERC20TokenRemoteUpgradeableTokensAndCallSent)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ERC20TokenRemoteUpgradeableTokensAndCallSentIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ERC20TokenRemoteUpgradeableTokensAndCallSentIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ERC20TokenRemoteUpgradeableTokensAndCallSent represents a TokensAndCallSent event raised by the ERC20TokenRemoteUpgradeable contract.
type ERC20TokenRemoteUpgradeableTokensAndCallSent struct {
	TeleporterMessageID [32]byte
	Sender              common.Address
	Input               SendAndCallInput
	Amount              *big.Int
	Raw                 types.Log // Blockchain specific contextual infos
}

// FilterTokensAndCallSent is a free log retrieval operation binding the contract event 0x5d76dff81bf773b908b050fa113d39f7d8135bb4175398f313ea19cd3a1a0b16.
//
// Solidity: event TokensAndCallSent(bytes32 indexed teleporterMessageID, address indexed sender, (bytes32,address,address,bytes,uint256,uint256,address,address,address,uint256,uint256) input, uint256 amount)
func (_ERC20TokenRemoteUpgradeable *ERC20TokenRemoteUpgradeableFilterer) FilterTokensAndCallSent(opts *bind.FilterOpts, teleporterMessageID [][32]byte, sender []common.Address) (*ERC20TokenRemoteUpgradeableTokensAndCallSentIterator, error) {

	var teleporterMessageIDRule []interface{}
	for _, teleporterMessageIDItem := range teleporterMessageID {
		teleporterMessageIDRule = append(teleporterMessageIDRule, teleporterMessageIDItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _ERC20TokenRemoteUpgradeable.contract.FilterLogs(opts, "TokensAndCallSent", teleporterMessageIDRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &ERC20TokenRemoteUpgradeableTokensAndCallSentIterator{contract: _ERC20TokenRemoteUpgradeable.contract, event: "TokensAndCallSent", logs: logs, sub: sub}, nil
}

// WatchTokensAndCallSent is a free log subscription operation binding the contract event 0x5d76dff81bf773b908b050fa113d39f7d8135bb4175398f313ea19cd3a1a0b16.
//
// Solidity: event TokensAndCallSent(bytes32 indexed teleporterMessageID, address indexed sender, (bytes32,address,address,bytes,uint256,uint256,address,address,address,uint256,uint256) input, uint256 amount)
func (_ERC20TokenRemoteUpgradeable *ERC20TokenRemoteUpgradeableFilterer) WatchTokensAndCallSent(opts *bind.WatchOpts, sink chan<- *ERC20TokenRemoteUpgradeableTokensAndCallSent, teleporterMessageID [][32]byte, sender []common.Address) (event.Subscription, error) {

	var teleporterMessageIDRule []interface{}
	for _, teleporterMessageIDItem := range teleporterMessageID {
		teleporterMessageIDRule = append(teleporterMessageIDRule, teleporterMessageIDItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _ERC20TokenRemoteUpgradeable.contract.WatchLogs(opts, "TokensAndCallSent", teleporterMessageIDRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ERC20TokenRemoteUpgradeableTokensAndCallSent)
				if err := _ERC20TokenRemoteUpgradeable.contract.UnpackLog(event, "TokensAndCallSent", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseTokensAndCallSent is a log parse operation binding the contract event 0x5d76dff81bf773b908b050fa113d39f7d8135bb4175398f313ea19cd3a1a0b16.
//
// Solidity: event TokensAndCallSent(bytes32 indexed teleporterMessageID, address indexed sender, (bytes32,address,address,bytes,uint256,uint256,address,address,address,uint256,uint256) input, uint256 amount)
func (_ERC20TokenRemoteUpgradeable *ERC20TokenRemoteUpgradeableFilterer) ParseTokensAndCallSent(log types.Log) (*ERC20TokenRemoteUpgradeableTokensAndCallSent, error) {
	event := new(ERC20TokenRemoteUpgradeableTokensAndCallSent)
	if err := _ERC20TokenRemoteUpgradeable.contract.UnpackLog(event, "TokensAndCallSent", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ERC20TokenRemoteUpgradeableTokensSentIterator is returned from FilterTokensSent and is used to iterate over the raw logs and unpacked data for TokensSent events raised by the ERC20TokenRemoteUpgradeable contract.
type ERC20TokenRemoteUpgradeableTokensSentIterator struct {
	Event *ERC20TokenRemoteUpgradeableTokensSent // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ERC20TokenRemoteUpgradeableTokensSentIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ERC20TokenRemoteUpgradeableTokensSent)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ERC20TokenRemoteUpgradeableTokensSent)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ERC20TokenRemoteUpgradeableTokensSentIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ERC20TokenRemoteUpgradeableTokensSentIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ERC20TokenRemoteUpgradeableTokensSent represents a TokensSent event raised by the ERC20TokenRemoteUpgradeable contract.
type ERC20TokenRemoteUpgradeableTokensSent struct {
	TeleporterMessageID [32]byte
	Sender              common.Address
	Input               SendTokensInput
	Amount              *big.Int
	Raw                 types.Log // Blockchain specific contextual infos
}

// FilterTokensSent is a free log retrieval operation binding the contract event 0x93f19bf1ec58a15dc643b37e7e18a1c13e85e06cd11929e283154691ace9fb52.
//
// Solidity: event TokensSent(bytes32 indexed teleporterMessageID, address indexed sender, (bytes32,address,address,address,uint256,uint256,uint256,address) input, uint256 amount)
func (_ERC20TokenRemoteUpgradeable *ERC20TokenRemoteUpgradeableFilterer) FilterTokensSent(opts *bind.FilterOpts, teleporterMessageID [][32]byte, sender []common.Address) (*ERC20TokenRemoteUpgradeableTokensSentIterator, error) {

	var teleporterMessageIDRule []interface{}
	for _, teleporterMessageIDItem := range teleporterMessageID {
		teleporterMessageIDRule = append(teleporterMessageIDRule, teleporterMessageIDItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _ERC20TokenRemoteUpgradeable.contract.FilterLogs(opts, "TokensSent", teleporterMessageIDRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &ERC20TokenRemoteUpgradeableTokensSentIterator{contract: _ERC20TokenRemoteUpgradeable.contract, event: "TokensSent", logs: logs, sub: sub}, nil
}

// WatchTokensSent is a free log subscription operation binding the contract event 0x93f19bf1ec58a15dc643b37e7e18a1c13e85e06cd11929e283154691ace9fb52.
//
// Solidity: event TokensSent(bytes32 indexed teleporterMessageID, address indexed sender, (bytes32,address,address,address,uint256,uint256,uint256,address) input, uint256 amount)
func (_ERC20TokenRemoteUpgradeable *ERC20TokenRemoteUpgradeableFilterer) WatchTokensSent(opts *bind.WatchOpts, sink chan<- *ERC20TokenRemoteUpgradeableTokensSent, teleporterMessageID [][32]byte, sender []common.Address) (event.Subscription, error) {

	var teleporterMessageIDRule []interface{}
	for _, teleporterMessageIDItem := range teleporterMessageID {
		teleporterMessageIDRule = append(teleporterMessageIDRule, teleporterMessageIDItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _ERC20TokenRemoteUpgradeable.contract.WatchLogs(opts, "TokensSent", teleporterMessageIDRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ERC20TokenRemoteUpgradeableTokensSent)
				if err := _ERC20TokenRemoteUpgradeable.contract.UnpackLog(event, "TokensSent", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseTokensSent is a log parse operation binding the contract event 0x93f19bf1ec58a15dc643b37e7e18a1c13e85e06cd11929e283154691ace9fb52.
//
// Solidity: event TokensSent(bytes32 indexed teleporterMessageID, address indexed sender, (bytes32,address,address,address,uint256,uint256,uint256,address) input, uint256 amount)
func (_ERC20TokenRemoteUpgradeable *ERC20TokenRemoteUpgradeableFilterer) ParseTokensSent(log types.Log) (*ERC20TokenRemoteUpgradeableTokensSent, error) {
	event := new(ERC20TokenRemoteUpgradeableTokensSent)
	if err := _ERC20TokenRemoteUpgradeable.contract.UnpackLog(event, "TokensSent", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ERC20TokenRemoteUpgradeableTokensWithdrawnIterator is returned from FilterTokensWithdrawn and is used to iterate over the raw logs and unpacked data for TokensWithdrawn events raised by the ERC20TokenRemoteUpgradeable contract.
type ERC20TokenRemoteUpgradeableTokensWithdrawnIterator struct {
	Event *ERC20TokenRemoteUpgradeableTokensWithdrawn // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ERC20TokenRemoteUpgradeableTokensWithdrawnIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ERC20TokenRemoteUpgradeableTokensWithdrawn)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ERC20TokenRemoteUpgradeableTokensWithdrawn)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ERC20TokenRemoteUpgradeableTokensWithdrawnIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ERC20TokenRemoteUpgradeableTokensWithdrawnIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ERC20TokenRemoteUpgradeableTokensWithdrawn represents a TokensWithdrawn event raised by the ERC20TokenRemoteUpgradeable contract.
type ERC20TokenRemoteUpgradeableTokensWithdrawn struct {
	Recipient common.Address
	Amount    *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterTokensWithdrawn is a free log retrieval operation binding the contract event 0x6352c5382c4a4578e712449ca65e83cdb392d045dfcf1cad9615189db2da244b.
//
// Solidity: event TokensWithdrawn(address indexed recipient, uint256 amount)
func (_ERC20TokenRemoteUpgradeable *ERC20TokenRemoteUpgradeableFilterer) FilterTokensWithdrawn(opts *bind.FilterOpts, recipient []common.Address) (*ERC20TokenRemoteUpgradeableTokensWithdrawnIterator, error) {

	var recipientRule []interface{}
	for _, recipientItem := range recipient {
		recipientRule = append(recipientRule, recipientItem)
	}

	logs, sub, err := _ERC20TokenRemoteUpgradeable.contract.FilterLogs(opts, "TokensWithdrawn", recipientRule)
	if err != nil {
		return nil, err
	}
	return &ERC20TokenRemoteUpgradeableTokensWithdrawnIterator{contract: _ERC20TokenRemoteUpgradeable.contract, event: "TokensWithdrawn", logs: logs, sub: sub}, nil
}

// WatchTokensWithdrawn is a free log subscription operation binding the contract event 0x6352c5382c4a4578e712449ca65e83cdb392d045dfcf1cad9615189db2da244b.
//
// Solidity: event TokensWithdrawn(address indexed recipient, uint256 amount)
func (_ERC20TokenRemoteUpgradeable *ERC20TokenRemoteUpgradeableFilterer) WatchTokensWithdrawn(opts *bind.WatchOpts, sink chan<- *ERC20TokenRemoteUpgradeableTokensWithdrawn, recipient []common.Address) (event.Subscription, error) {

	var recipientRule []interface{}
	for _, recipientItem := range recipient {
		recipientRule = append(recipientRule, recipientItem)
	}

	logs, sub, err := _ERC20TokenRemoteUpgradeable.contract.WatchLogs(opts, "TokensWithdrawn", recipientRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ERC20TokenRemoteUpgradeableTokensWithdrawn)
				if err := _ERC20TokenRemoteUpgradeable.contract.UnpackLog(event, "TokensWithdrawn", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseTokensWithdrawn is a log parse operation binding the contract event 0x6352c5382c4a4578e712449ca65e83cdb392d045dfcf1cad9615189db2da244b.
//
// Solidity: event TokensWithdrawn(address indexed recipient, uint256 amount)
func (_ERC20TokenRemoteUpgradeable *ERC20TokenRemoteUpgradeableFilterer) ParseTokensWithdrawn(log types.Log) (*ERC20TokenRemoteUpgradeableTokensWithdrawn, error) {
	event := new(ERC20TokenRemoteUpgradeableTokensWithdrawn)
	if err := _ERC20TokenRemoteUpgradeable.contract.UnpackLog(event, "TokensWithdrawn", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ERC20TokenRemoteUpgradeableTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the ERC20TokenRemoteUpgradeable contract.
type ERC20TokenRemoteUpgradeableTransferIterator struct {
	Event *ERC20TokenRemoteUpgradeableTransfer // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ERC20TokenRemoteUpgradeableTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ERC20TokenRemoteUpgradeableTransfer)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ERC20TokenRemoteUpgradeableTransfer)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ERC20TokenRemoteUpgradeableTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ERC20TokenRemoteUpgradeableTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ERC20TokenRemoteUpgradeableTransfer represents a Transfer event raised by the ERC20TokenRemoteUpgradeable contract.
type ERC20TokenRemoteUpgradeableTransfer struct {
	From  common.Address
	To    common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_ERC20TokenRemoteUpgradeable *ERC20TokenRemoteUpgradeableFilterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*ERC20TokenRemoteUpgradeableTransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _ERC20TokenRemoteUpgradeable.contract.FilterLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &ERC20TokenRemoteUpgradeableTransferIterator{contract: _ERC20TokenRemoteUpgradeable.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_ERC20TokenRemoteUpgradeable *ERC20TokenRemoteUpgradeableFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *ERC20TokenRemoteUpgradeableTransfer, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _ERC20TokenRemoteUpgradeable.contract.WatchLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ERC20TokenRemoteUpgradeableTransfer)
				if err := _ERC20TokenRemoteUpgradeable.contract.UnpackLog(event, "Transfer", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseTransfer is a log parse operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_ERC20TokenRemoteUpgradeable *ERC20TokenRemoteUpgradeableFilterer) ParseTransfer(log types.Log) (*ERC20TokenRemoteUpgradeableTransfer, error) {
	event := new(ERC20TokenRemoteUpgradeableTransfer)
	if err := _ERC20TokenRemoteUpgradeable.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
