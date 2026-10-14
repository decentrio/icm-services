// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package nativetokenhome

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

// RemoteTokenTransferrerSettings is an auto generated low-level Go binding around an user-defined struct.
type RemoteTokenTransferrerSettings struct {
	Registered       bool
	CollateralNeeded *big.Int
	TokenMultiplier  *big.Int
	MultiplyOnRemote bool
}

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

// NativeTokenHomeMetaData contains all meta data concerning the NativeTokenHome contract.
var NativeTokenHomeMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"blockchainID\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"teleporterManager\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"wrappedTokenAddress\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"FailedCall\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"balance\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"needed\",\"type\":\"uint256\"}],\"name\":\"InsufficientBalance\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidInitialization\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotInitializing\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"OwnableInvalidOwner\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"OwnableUnauthorizedAccount\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ReentrancyGuardReentrantCall\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"SafeERC20FailedOperation\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"recipientContract\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"CallFailed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"recipientContract\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"CallSucceeded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"remoteBlockchainID\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"remoteTokenTransferrerAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"remaining\",\"type\":\"uint256\"}],\"name\":\"CollateralAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"version\",\"type\":\"uint64\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"remoteBlockchainID\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"messenger\",\"type\":\"address\"}],\"name\":\"MessengerAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"messenger\",\"type\":\"address\"}],\"name\":\"MessengerPaused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"remoteBlockchainID\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"messenger\",\"type\":\"address\"}],\"name\":\"MessengerRemoved\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"messenger\",\"type\":\"address\"}],\"name\":\"MessengerUnpaused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"remoteBlockchainID\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"oldMessenger\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newMessenger\",\"type\":\"address\"}],\"name\":\"OutboundMessengerSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"remoteBlockchainID\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"remoteTokenTransferrerAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"initialCollateralNeeded\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"tokenDecimals\",\"type\":\"uint8\"}],\"name\":\"RemoteRegistered\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"teleporterMessageID\",\"type\":\"bytes32\"},{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"destinationBlockchainID\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"destinationTokenTransferrerAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"recipientContract\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"recipientPayload\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"requiredGasLimit\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"recipientGasLimit\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"multiHopFallback\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"fallbackRecipient\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"primaryFeeTokenAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"primaryFee\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"secondaryFee\",\"type\":\"uint256\"}],\"indexed\":false,\"internalType\":\"structSendAndCallInput\",\"name\":\"input\",\"type\":\"tuple\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"TokensAndCallRouted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"teleporterMessageID\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"destinationBlockchainID\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"destinationTokenTransferrerAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"recipientContract\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"recipientPayload\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"requiredGasLimit\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"recipientGasLimit\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"multiHopFallback\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"fallbackRecipient\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"primaryFeeTokenAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"primaryFee\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"secondaryFee\",\"type\":\"uint256\"}],\"indexed\":false,\"internalType\":\"structSendAndCallInput\",\"name\":\"input\",\"type\":\"tuple\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"TokensAndCallSent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"teleporterMessageID\",\"type\":\"bytes32\"},{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"destinationBlockchainID\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"destinationTokenTransferrerAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"primaryFeeTokenAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"primaryFee\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"secondaryFee\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"requiredGasLimit\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"multiHopFallback\",\"type\":\"address\"}],\"indexed\":false,\"internalType\":\"structSendTokensInput\",\"name\":\"input\",\"type\":\"tuple\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"TokensRouted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"teleporterMessageID\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"destinationBlockchainID\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"destinationTokenTransferrerAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"primaryFeeTokenAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"primaryFee\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"secondaryFee\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"requiredGasLimit\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"multiHopFallback\",\"type\":\"address\"}],\"indexed\":false,\"internalType\":\"structSendTokensInput\",\"name\":\"input\",\"type\":\"tuple\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"TokensSent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"TokensWithdrawn\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"NATIVE_TOKEN_HOME_STORAGE_LOCATION\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"TELEPORTER_ROUTED_APP_STORAGE_LOCATION\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"TOKEN_HOME_STORAGE_LOCATION\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"remoteBlockchainID\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"remoteTokenTransferrerAddress\",\"type\":\"address\"}],\"name\":\"addCollateral\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"remoteBlockchainID\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"messenger\",\"type\":\"address\"}],\"name\":\"addMessenger\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getBlockchainID\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"remoteBlockchainID\",\"type\":\"bytes32\"}],\"name\":\"getOutboundMessenger\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"remoteBlockchainID\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"remoteTokenTransferrerAddress\",\"type\":\"address\"}],\"name\":\"getRemoteTokenTransferrerSettings\",\"outputs\":[{\"components\":[{\"internalType\":\"bool\",\"name\":\"registered\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"collateralNeeded\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"tokenMultiplier\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"multiplyOnRemote\",\"type\":\"bool\"}],\"internalType\":\"structRemoteTokenTransferrerSettings\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getTokenAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"remoteBlockchainID\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"remoteTokenTransferrerAddress\",\"type\":\"address\"}],\"name\":\"getTransferredBalance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"blockchainID\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"teleporterManager\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"wrappedTokenAddress\",\"type\":\"address\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"remoteBlockchainID\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"messenger\",\"type\":\"address\"}],\"name\":\"isInboundMessenger\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"messenger\",\"type\":\"address\"}],\"name\":\"isMessengerPaused\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"messenger\",\"type\":\"address\"}],\"name\":\"pauseMessenger\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"sourceBlockchainID\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"originSenderAddress\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"message\",\"type\":\"bytes\"}],\"name\":\"receiveTeleporterMessage\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"remoteBlockchainID\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"messenger\",\"type\":\"address\"}],\"name\":\"removeMessenger\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"destinationBlockchainID\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"destinationTokenTransferrerAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"primaryFeeTokenAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"primaryFee\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"secondaryFee\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"requiredGasLimit\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"multiHopFallback\",\"type\":\"address\"}],\"internalType\":\"structSendTokensInput\",\"name\":\"input\",\"type\":\"tuple\"}],\"name\":\"send\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"destinationBlockchainID\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"destinationTokenTransferrerAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"recipientContract\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"recipientPayload\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"requiredGasLimit\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"recipientGasLimit\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"multiHopFallback\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"fallbackRecipient\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"primaryFeeTokenAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"primaryFee\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"secondaryFee\",\"type\":\"uint256\"}],\"internalType\":\"structSendAndCallInput\",\"name\":\"input\",\"type\":\"tuple\"}],\"name\":\"sendAndCall\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"remoteBlockchainID\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"messenger\",\"type\":\"address\"}],\"name\":\"setOutboundMessenger\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"messenger\",\"type\":\"address\"}],\"name\":\"unpauseMessenger\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"stateMutability\":\"payable\",\"type\":\"receive\"}]",
	Bin: "0x608060405234801561000f575f5ffd5b5060405161513138038061513183398101604081905261002e9161054e565b610039838383610044565b505050610587565b50565b5f5160206151115f395f51905f52805468010000000000000000810460ff1615906001600160401b03165f8115801561007a5750825b90505f826001600160401b031660011480156100955750303b155b9050811580156100a3575080155b156100c15760405163f92ee8a960e01b815260040160405180910390fd5b84546001600160401b031916600117855583156100ef57845460ff60401b1916680100000000000000001785555b6100fa88888861014a565b831561014057845460ff60401b19168555604051600181527fc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d29060200160405180910390a15b5050505050505050565b61015261016d565b61015f83838360126101aa565b610168816101d4565b505050565b5f5160206151115f395f51905f525468010000000000000000900460ff166101a857604051631afcd79f60e31b815260040160405180910390fd5b565b6101b261016d565b6101bc848461021d565b6101c461023b565b6101ce828261024b565b50505050565b6101dc61016d565b7f3b5030f10c94fcbdaa3022348ff0b82dbd4c0c71339e41ff59d0bdc92179d60080546001600160a01b0319166001600160a01b0392909216919091179055565b61022561016d565b61022e8261035d565b6102378161037e565b5050565b61024361016d565b6101a861038f565b61025361016d565b6001600160a01b0382166102ae5760405162461bcd60e51b815260206004820152601d60248201527f546f6b656e486f6d653a207a65726f20746f6b656e206164647265737300000060448201526064015b60405180910390fd5b60128160ff16111561030d5760405162461bcd60e51b815260206004820152602260248201527f546f6b656e486f6d653a20746f6b656e20646563696d616c7320746f6f2068696044820152610ced60f31b60648201526084016102a5565b7f9316912b5a9db88acbe872c934fdd0a46c436c6dcba332d649c4d57c7bc9e601805460ff92909216600160a01b026001600160a81b03199092166001600160a01b039390931692909217179055565b61036561016d565b61036d6103b9565b6103756103c9565b610041816103d1565b61038661016d565b6100418161045a565b5f7fd2f1ed38b7d242bfb8b41862afb813a15193219a4bc717f2056607593e6c75005b6001905550565b6103c161016d565b6101a8610494565b6101a861016d565b6103d961016d565b806104365760405162461bcd60e51b815260206004820152602760248201527f54656c65706f72746572526f757465644170703a207a65726f20626c6f636b636044820152661a185a5b88125160ca1b60648201526084016102a5565b7f63108c0ddc3d3ab0f279f27d4b9ca8220d5efbdbeeb537d4327d5cf9e982b10055565b61046261016d565b6001600160a01b03811661048b57604051631e4fbdf760e01b81525f60048201526024016102a5565b610041816104c3565b61049c61016d565b5f7f9b779b17422d0df92223018b32b4d1fa46e071723d6817e2486d003becc55f006103b2565b7f9016d09d72d40fdae2fd8ceac6b6234c7706214fd39c1cd1e609a0528c19930080546001600160a01b031981166001600160a01b03848116918217845560405192169182907f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0905f90a3505050565b80516001600160a01b0381168114610549575f5ffd5b919050565b5f5f5f60608486031215610560575f5ffd5b8351925061057060208501610533565b915061057e60408501610533565b90509250925092565b614b7d806105945f395ff3fe60806040526004361061013f575f3560e01c806395a19443116100b3578063d90f1b101161006d578063d90f1b101461054a578063dd5a20fd14610569578063efb5b95e14610588578063f2fde38b146105a8578063f4be94a7146105c7578063fc560821146105e6575f5ffd5b806395a19443146103a9578063b0b78b26146103f9578063b24325291461040c578063c1c6353e1461042b578063c8511ada1461044b578063c868efaa1461052b575f5ffd5b80636e6eef8d116101045780636e6eef8d146102ae578063715018a6146102c157806383c002d0146102d557806388b117b31461033b5780638bf2fa941461035a5780638da5cb5b1461036d575f5ffd5b806303f78951146101cd57806310fe9ae8146101ec578063154d625a146102415780634213cf781461026e57806362e3901b1461028e575f5ffd5b366101c9575f516020614aa85f395f51905f52546001600160a01b031633146101c75760405162461bcd60e51b815260206004820152602f60248201527f4e6174697665546f6b656e486f6d653a20696e76616c6964207265636569766560448201526e103830bcb0b136329039b2b73232b960891b60648201526084015b60405180910390fd5b005b5f5ffd5b3480156101d8575f5ffd5b506101c76101e7366004613b60565b610639565b3480156101f7575f5ffd5b507f9316912b5a9db88acbe872c934fdd0a46c436c6dcba332d649c4d57c7bc9e601546001600160a01b03165b6040516001600160a01b0390911681526020015b60405180910390f35b34801561024c575f5ffd5b5061026061025b366004613b7b565b610743565b604051908152602001610238565b348015610279575f5ffd5b505f516020614b085f395f51905f5254610260565b348015610299575f5ffd5b506102605f516020614a885f395f51905f5281565b6101c76102bc366004613ba9565b61078b565b3480156102cc575f5ffd5b506101c76107b7565b3480156102e0575f5ffd5b5061032b6102ef366004613b60565b6001600160a01b03165f9081527f63108c0ddc3d3ab0f279f27d4b9ca8220d5efbdbeeb537d4327d5cf9e982b103602052604090205460ff1690565b6040519015158152602001610238565b348015610346575f5ffd5b506101c7610355366004613be0565b6107ca565b6101c7610368366004613c1f565b6108da565b348015610378575f5ffd5b507f9016d09d72d40fdae2fd8ceac6b6234c7706214fd39c1cd1e609a0528c199300546001600160a01b0316610224565b3480156103b4575f5ffd5b5061032b6103c3366004613b7b565b5f9182525f516020614b285f395f51905f52602090815260408084206001600160a01b0393909316845291905290205460ff1690565b6101c7610407366004613b7b565b6108f2565b348015610417575f5ffd5b506101c7610426366004613b7b565b610901565b348015610436575f5ffd5b506102605f516020614b085f395f51905f5281565b348015610456575f5ffd5b506104f4610465366004613b7b565b60408051608080820183525f808352602080840182905283850182905260609384018290529581525f516020614ac85f395f51905f5286528381206001600160a01b039590951681529385529282902082519384018352805460ff9081161515855260018201549585019590955260028101549284019290925260039091015490921615159181019190915290565b6040516102389190815115158152602080830151908201526040808301519082015260609182015115159181019190915260800190565b348015610536575f5ffd5b506101c7610545366004613c39565b610913565b348015610555575f5ffd5b506101c7610564366004613b7b565b610a32565b348015610574575f5ffd5b506101c7610583366004613b60565b610a44565b348015610593575f5ffd5b506102605f516020614aa85f395f51905f5281565b3480156105b3575f5ffd5b506101c76105c2366004613b60565b610b46565b3480156105d2575f5ffd5b506101c76105e1366004613b7b565b610b80565b3480156105f1575f5ffd5b50610224610600366004613cbc565b5f9081527f63108c0ddc3d3ab0f279f27d4b9ca8220d5efbdbeeb537d4327d5cf9e982b10160205260409020546001600160a01b031690565b610641610cc6565b5f516020614b085f395f51905f526001600160a01b0382166106755760405162461bcd60e51b81526004016101be90613cd3565b6001600160a01b0382165f90815260038201602052604090205460ff16156106f55760405162461bcd60e51b815260206004820152602d60248201527f54656c65706f72746572526f757465644170703a206d657373656e676572206160448201526c1b1c9958591e481c185d5cd959609a1b60648201526084016101be565b6001600160a01b0382165f818152600383016020526040808220805460ff19166001179055517f46def370ef3a1ec72618480f66f1063a80fe68b27d09575bc3562ad489079b619190a25050565b5f8281527f9316912b5a9db88acbe872c934fdd0a46c436c6dcba332d649c4d57c7bc9e603602090815260408083206001600160a01b03851684529091529020545b92915050565b6107b46107a35f516020614b085f395f51905f525490565b30336107ae85613e6a565b34610cce565b50565b6107bf610f31565b6107c85f610f8c565b565b7ff0c57e16840df040f15088dc2f81fe391c3923bec73e23a9662efc9c229c6a008054600160401b810460ff1615906001600160401b03165f8115801561080e5750825b90505f826001600160401b031660011480156108295750303b155b905081158015610837575080155b156108555760405163f92ee8a960e01b815260040160405180910390fd5b845467ffffffffffffffff19166001178555831561087f57845460ff60401b1916600160401b1785555b61088a888888610ffc565b83156108d057845460ff60401b19168555604051600181527fc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d29060200160405180910390a15b5050505050505050565b6107b46108ec36839003830183613f39565b3461101f565b6108fd8282346111d9565b5050565b610909610cc6565b6108fd82826113c9565b61091b61168e565b5f8481525f516020614b285f395f51905f52602090815260408083203384529091529020545f516020614b085f395f51905f529060ff166109b25760405162461bcd60e51b815260206004820152602b60248201527f54656c65706f72746572526f757465644170703a20756e617574686f72697a6560448201526a321036b2b9b9b2b733b2b960a91b60648201526084016101be565b335f90815260038201602052604090205460ff16156109e35760405162461bcd60e51b81526004016101be90613fc3565b610a23858585858080601f0160208091040260200160405190810160405280939291908181526020018383808284375f920191909152506116d892505050565b50610a2c611ac2565b50505050565b610a3a610cc6565b6108fd8282611aec565b610a4c610cc6565b5f516020614b085f395f51905f526001600160a01b038216610a805760405162461bcd60e51b81526004016101be90613cd3565b6001600160a01b0382165f90815260038201602052604090205460ff16610afb5760405162461bcd60e51b815260206004820152602960248201527f54656c65706f72746572526f757465644170703a206d657373656e676572206e6044820152681bdd081c185d5cd95960ba1b60648201526084016101be565b6001600160a01b0382165f818152600383016020526040808220805460ff19169055517ff5574433a9d82924af98e239704e51afb34ae8c15f513d78a22ae0c58f5f99659190a25050565b610b4e610f31565b6001600160a01b038116610b7757604051631e4fbdf760e01b81525f60048201526024016101be565b6107b481610f8c565b610b88610cc6565b5f8281525f516020614b285f395f51905f52602090815260408083206001600160a01b03851684529091529020545f516020614b085f395f51905f529060ff16610be45760405162461bcd60e51b81526004016101be90614008565b5f8381526001820160205260409020546001600160a01b03808416911603610c6c5760405162461bcd60e51b815260206004820152603560248201527f54656c65706f72746572526f757465644170703a2063616e6e6f742072656d6f6044820152743b329037baba3137bab7321036b2b9b9b2b733b2b960591b60648201526084016101be565b5f83815260028201602090815260408083206001600160a01b0386168085529252808320805460ff1916905551909185917f8350b16da8b8ec7d7e6ce37ca2ccb4f63d628b5ad53215d072763b6afce995249190a3505050565b6107c8610f31565b5f516020614ae85f395f51905f528054600114610cfd5760405162461bcd60e51b81526004016101be90614050565b60028155610d0a83611c2c565b60c08301516001600160a01b031615610d355760405162461bcd60e51b81526004016101be90614094565b5f5f610d55855f0151866020015186886101000151896101200151611e11565b915091505f604051806040016040528060026004811115610d7857610d786140da565b81526020016040518061010001604052808c81526020018b6001600160a01b031681526020018a6001600160a01b0316815260200189604001516001600160a01b03168152602001868152602001896060015181526020018960a0015181526020018960e001516001600160a01b0316815250604051602001610dfb919061413b565b60405160208183030381529060405281525090505f610edb6040518060c00160405280895f0151815260200189602001516001600160a01b0316815260200160405180604001604052808b61010001516001600160a01b03168152602001878152508152602001896080015181526020015f6001600160401b03811115610e8457610e84613d1e565b604051908082528060200260200182016040528015610ead578160200160208202803683370190505b50815260200184604051602001610ec491906141da565b604051602081830303815290604052815250611fd6565b9050876001600160a01b0316817f5d76dff81bf773b908b050fa113d39f7d8135bb4175398f313ea19cd3a1a0b168987604051610f1992919061421c565b60405180910390a35050600190925550505050505050565b33610f637f9016d09d72d40fdae2fd8ceac6b6234c7706214fd39c1cd1e609a0528c199300546001600160a01b031690565b6001600160a01b0316146107c85760405163118cdaa760e01b81523360048201526024016101be565b7f9016d09d72d40fdae2fd8ceac6b6234c7706214fd39c1cd1e609a0528c19930080546001600160a01b031981166001600160a01b03848116918217845560405192169182907f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0905f90a3505050565b6110046120fa565b6110118383836012612143565b61101a81612167565b505050565b5f516020614ae85f395f51905f52805460011461104e5760405162461bcd60e51b81526004016101be90614050565b6002815561105b8361219d565b60e08301516001600160a01b0316156110865760405162461bcd60e51b81526004016101be90614094565b5f5f6110a4855f015186602001518688606001518960800151611e11565b915091505f6040518060400160405280600160048111156110c7576110c76140da565b8152602001604051806040016040528089604001516001600160a01b03168152602001868152506040516020016110fe91906142f1565b60405160208183030381529060405281525090505f6111866040518060c00160405280895f0151815260200189602001516001600160a01b0316815260200160405180604001604052808b606001516001600160a01b031681526020018781525081526020018960c0015181526020015f6001600160401b03811115610e8457610e84613d1e565b9050336001600160a01b0316817f93f19bf1ec58a15dc643b37e7e18a1c13e85e06cd11929e283154691ace9fb5289876040516111c4929190614311565b60405180910390a35050600190925550505050565b5f516020614ae85f395f51905f5280546001146112085760405162461bcd60e51b81526004016101be90614050565b60028082555f8581525f516020614ac85f395f51905f52602090815260408083206001600160a01b03881684528252918290208251608081018452815460ff9081161515808352600184015494830194909452948201549381019390935260030154909216151560608201525f516020614a885f395f51905f529161129f5760405162461bcd60e51b81526004016101be9061437e565b5f8160200151116112fc5760405162461bcd60e51b815260206004820152602160248201527f546f6b656e486f6d653a207a65726f20636f6c6c61746572616c206e656564656044820152601960fa1b60648201526084016101be565b61130584612246565b93505f5f826020015186106113345760208301515f925061132690876143c7565b905082602001519550611347565b85836020015161134491906143c7565b91505b5f88815260028501602090815260408083206001600160a01b038b168085529083529281902060010185905580518981529182018590528a917f6769a5f9bfc8b6e0db839ab981cbf9239274ae72d2d035081a9157d43bd33cb6910160405180910390a380156113bb576113bb338261226c565b505060019092555050505050565b5f516020614b085f395f51905f528261143b5760405162461bcd60e51b815260206004820152602e60248201527f54656c65706f72746572526f757465644170703a207a65726f2072656d6f746560448201526d08189b1bd8dad8da185a5b88125160921b60648201526084016101be565b805483036114a75760405162461bcd60e51b815260206004820152603360248201527f54656c65706f72746572526f757465644170703a2063616e6e6f7420726f757460448201527232903a379037bbb710313637b1b5b1b430b4b760691b60648201526084016101be565b6001600160a01b0382166114cd5760405162461bcd60e51b81526004016101be90613cd3565b5f83815260028201602090815260408083206001600160a01b038616845290915290205460ff16156115565760405162461bcd60e51b815260206004820152602c60248201527f54656c65706f72746572526f757465644170703a206d657373656e676572206160448201526b1b1c9958591e48185919195960a21b60648201526084016101be565b5f826001600160a01b031663d127dc9b6040518163ffffffff1660e01b8152600401602060405180830381865afa158015611593573d5f5f3e3d5ffd5b505050506040513d601f19601f820116820180604052508101906115b791906143da565b90508015806115c65750815481145b6116305760405162461bcd60e51b815260206004820152603560248201527f54656c65706f72746572526f757465644170703a206d657373656e67657220626044820152740d8dec6d6c6d0c2d2dc40928840dad2e6dac2e8c6d605b1b60648201526084016101be565b5f84815260028301602090815260408083206001600160a01b0387168085529252808320805460ff1916600117905551909186917fcc2f89e16e8979d86d3c3b49c5646f3a013f76626a479400e1cb3596f3dc2b229190a350505050565b7f9b779b17422d0df92223018b32b4d1fa46e071723d6817e2486d003becc55f008054600119016116d257604051633ee5aeb560e01b815260040160405180910390fd5b60029055565b5f5f516020614a885f395f51905f5290505f828060200190518101906116fe9190614433565b9050600181516004811115611715576117156140da565b0361175d575f816020015180602001905181019061173391906144be565b90505f61174587878460200151612325565b9050611754825f01518261226c565b50505050505050565b600281516004811115611772576117726140da565b0361188b575f816020015180602001905181019061179091906144f8565b90505f6117a287878460800151612325565b825190915087146118085760405162461bcd60e51b815260206004820152602a60248201527f546f6b656e486f6d653a206d69736d61746368656420736f7572636520626c6f60448201526918dad8da185a5b88125160b21b60648201526084016101be565b856001600160a01b031682602001516001600160a01b0316146118815760405162461bcd60e51b815260206004820152602b60248201527f546f6b656e486f6d653a206d69736d617463686564206f726967696e2073656e60448201526a646572206164647265737360a81b60648201526084016101be565b61175482826123b0565b6003815160048111156118a0576118a06140da565b0361196a575f81602001518060200190518101906118be91906145c5565b90505f5f6118d6888885606001518660800151612536565b915091506108d0604051806101000160405280855f0151815260200185602001516001600160a01b0316815260200185604001516001600160a01b03168152602001876001015f9054906101000a90046001600160a01b03166001600160a01b031681526020018381526020015f81526020018560a0015181526020018560c001516001600160a01b0316815250836125e1565b60048151600481111561197f5761197f6140da565b03611a7a575f816020015180602001905181019061199d9190614641565b90505f5f6119b688888560800151866101400151612536565b915091506108d08888855f01516040518061016001604052808860200151815260200188604001516001600160a01b0316815260200188606001516001600160a01b031681526020018860a00151815260200188610100015181526020018860c0015181526020018861012001516001600160a01b031681526020018860e001516001600160a01b031681526020018a6001015f9054906101000a90046001600160a01b03166001600160a01b031681526020018681526020015f8152508661276c565b5f81516004811115611a8e57611a8e6140da565b03611abb575f8160200151806020019051810190611aac9190614749565b9050611ab9868683612949565b505b5050505050565b5f7f9b779b17422d0df92223018b32b4d1fa46e071723d6817e2486d003becc55f005b6001905550565b5f8281525f516020614b285f395f51905f52602090815260408083206001600160a01b03851684529091529020545f516020614b085f395f51905f529060ff16611b485760405162461bcd60e51b81526004016101be90614008565b5f8381526001820160205260409020546001600160a01b039081169083168103611bcc5760405162461bcd60e51b815260206004820152602f60248201527f54656c65706f72746572526f757465644170703a206d657373656e676572206160448201526e1b1c9958591e481bdd5d189bdd5b99608a1b60648201526084016101be565b5f84815260018301602052604080822080546001600160a01b0319166001600160a01b03878116918217909255915191929084169187917fd00498d0475cde985344ea08d024c9315ca654d9d157211295aac4349637411591a450505050565b60408101516001600160a01b0316611c995760405162461bcd60e51b815260206004820152602a60248201527f546f6b656e486f6d653a207a65726f20726563697069656e7420636f6e7472616044820152696374206164647265737360b01b60648201526084016101be565b5f816080015111611cbc5760405162461bcd60e51b81526004016101be906147ad565b5f8160a0015111611d1b5760405162461bcd60e51b815260206004820152602360248201527f546f6b656e486f6d653a207a65726f20726563697069656e7420676173206c696044820152621b5a5d60ea1b60648201526084016101be565b80608001518160a0015110611d815760405162461bcd60e51b815260206004820152602660248201527f546f6b656e486f6d653a20696e76616c696420726563697069656e7420676173604482015265081b1a5b5a5d60d21b60648201526084016101be565b60e08101516001600160a01b0316611dee5760405162461bcd60e51b815260206004820152602a60248201527f546f6b656e486f6d653a207a65726f2066616c6c6261636b20726563697069656044820152696e74206164647265737360b01b60648201526084016101be565b610140810151156107b45760405162461bcd60e51b81526004016101be906147ef565b5f8581525f516020614ac85f395f51905f52602090815260408083206001600160a01b038816845282528083208151608081018352815460ff90811615158083526001840154958301959095526002830154938201939093526003909101549091161515606082015282915f516020614a885f395f51905f529190611ea85760405162461bcd60e51b81526004016101be9061437e565b602081015115611f0a5760405162461bcd60e51b815260206004820152602760248201527f546f6b656e486f6d653a20636f6c6c61746572616c206e656564656420666f726044820152662072656d6f746560c81b60648201526084016101be565b611f1387612246565b96508415611f2957611f26863387612d5d565b94505b5f611f3d826040015183606001518a612eb6565b90505f8111611f8e5760405162461bcd60e51b815260206004820152601d60248201527f546f6b656e486f6d653a207a65726f207363616c656420616d6f756e7400000060448201526064016101be565b5f8a815260038401602090815260408083206001600160a01b038d16845290915281208054839290611fc1908490614830565b90915550909a95995094975050505050505050565b5f5f611fe4835f0151612ecc565b60408401516020015190915015612087576040830151516001600160a01b03166120645760405162461bcd60e51b815260206004820152602b60248201527f54656c65706f72746572526f757465644170703a207a65726f2066656520746f60448201526a6b656e206164647265737360a81b60648201526084016101be565b604083015160208101519051612087916001600160a01b03909116908390612fab565b604051630624488560e41b81526001600160a01b038216906362448850906120b3908690600401614843565b6020604051808303815f875af11580156120cf573d5f5f3e3d5ffd5b505050506040513d601f19601f820116820180604052508101906120f391906143da565b9392505050565b7ff0c57e16840df040f15088dc2f81fe391c3923bec73e23a9662efc9c229c6a0054600160401b900460ff166107c857604051631afcd79f60e31b815260040160405180910390fd5b61214b6120fa565b6121558484613032565b61215d61304c565b610a2c828261305c565b61216f6120fa565b5f516020614aa85f395f51905f5280546001600160a01b0319166001600160a01b0392909216919091179055565b60408101516001600160a01b03166122015760405162461bcd60e51b815260206004820152602160248201527f546f6b656e486f6d653a207a65726f20726563697069656e74206164647265736044820152607360f81b60648201526084016101be565b5f8160c00151116122245760405162461bcd60e51b81526004016101be906147ad565b60a0810151156107b45760405162461bcd60e51b81526004016101be906147ef565b5f516020614aa85f395f51905f5280545f91906120f3906001600160a01b031684613169565b6040518181525f516020614aa85f395f51905f52906001600160a01b038416907f6352c5382c4a4578e712449ca65e83cdb392d045dfcf1cad9615189db2da244b9060200160405180910390a28054604051632e1a7d4d60e01b8152600481018490526001600160a01b0390911690632e1a7d4d906024015f604051808303815f87803b1580156122fb575f5ffd5b505af115801561230d573d5f5f3e3d5ffd5b5061101a925050506001600160a01b0384168361330e565b5f8381525f516020614ac85f395f51905f52602090815260408083206001600160a01b038616845282528083208151608081018352815460ff9081161515825260018301549482019490945260028201549281019290925260030154909116151560608201525f516020614a885f395f51905f52906123a6818787876133a8565b9695505050505050565b5f5f516020614aa85f395f51905f528054604051632e1a7d4d60e01b8152600481018590529192506001600160a01b031690632e1a7d4d906024015f604051808303815f87803b158015612402575f5ffd5b505af1158015612414573d5f5f3e3d5ffd5b50508451602086015160408088015160a089015191515f965061243d95509091906024016148fe565b60408051601f198184030181529190526020810180516001600160e01b031663161b12ff60e11b17905260c085015160608601519192505f916124839190869085613496565b905080156124d75784606001516001600160a01b03167f104deb555f67e63782bb817bc26c39050894645f9b9f29c4be8ae68d0e8b7ff4856040516124ca91815260200190565b60405180910390a2611abb565b84606001516001600160a01b03167fb9eaeae386d339f8115782f297a9e5f0e13fb587cd6b0d502f113cb8dd4d6cb08560405161251691815260200190565b60405180910390a260e0850151611abb906001600160a01b03168561330e565b5f8481525f516020614ac85f395f51905f52602090815260408083206001600160a01b038716845282528083208151608081018352815460ff90811615158252600183015494820194909452600282015492810192909252600301549091161515606082015281905f516020614a885f395f51905f5290826125ba828a8a8a6133a8565b90505f6125d083604001518460600151896135d7565b919a91995090975050505050505050565b5f516020614ae85f395f51905f5280546001146126105760405162461bcd60e51b81526004016101be90614050565b6002815561261d8361219d565b5f612635845f015185602001518587608001516135e4565b9050805f036126525761264c8460e001518461226c565b50612764565b604080518082019091525f908060018152602001604051806040016040528088604001516001600160a01b031681526020018581525060405160200161269891906142f1565b60405160208183030381529060405281525090505f6127246040518060c00160405280885f0151815260200188602001516001600160a01b0316815260200160405180604001604052808a606001516001600160a01b031681526020018a6080015181525081526020018860c0015181526020015f6001600160401b03811115610e8457610e84613d1e565b9050807f825080857c76cef4a1629c0705a7f8b4ef0282ddcafde0b6715c4fb34b68aaf08785604051612758929190614311565b60405180910390a25050505b600190555050565b5f516020614ae85f395f51905f52805460011461279b5760405162461bcd60e51b81526004016101be90614050565b600281556127a883611c2c565b5f6127c1845f01518560200151858761012001516135e4565b9050805f036127de576127d88460c001518461226c565b5061293e565b604080518082019091525f9080600281526020016040518061010001604052808b81526020018a6001600160a01b03168152602001896001600160a01b0316815260200188604001516001600160a01b03168152602001858152602001886060015181526020018860a0015181526020018860e001516001600160a01b0316815250604051602001612870919061413b565b60405160208183030381529060405281525090505f6128fe6040518060c00160405280885f0151815260200188602001516001600160a01b0316815260200160405180604001604052808a61010001516001600160a01b031681526020018a61012001518152508152602001886080015181526020015f6001600160401b03811115610e8457610e84613d1e565b9050807f42eff9005856e3c586b096d67211a566dc926052119fd7cc08023c70937ecb30878560405161293292919061421c565b60405180910390a25050505b600190555050505050565b5f516020614a885f395f51905f52836129b05760405162461bcd60e51b8152602060048201526024808201527f546f6b656e486f6d653a207a65726f2072656d6f746520626c6f636b636861696044820152631b88125160e21b60648201526084016101be565b5f516020614b085f395f51905f52548403612a255760405162461bcd60e51b815260206004820152602f60248201527f546f6b656e486f6d653a2063616e6e6f742072656769737465722072656d6f7460448201526e329037b71039b0b6b29031b430b4b760891b60648201526084016101be565b6001600160a01b038316612a945760405162461bcd60e51b815260206004820152603060248201527f546f6b656e486f6d653a207a65726f2072656d6f746520746f6b656e2074726160448201526f6e73666572726572206164647265737360801b60648201526084016101be565b5f84815260028201602090815260408083206001600160a01b038716845290915290205460ff1615612b145760405162461bcd60e51b8152602060048201526024808201527f546f6b656e486f6d653a2072656d6f746520616c726561647920726567697374604482015263195c995960e21b60648201526084016101be565b6012826040015160ff161115612b7e5760405162461bcd60e51b815260206004820152602960248201527f546f6b656e486f6d653a2072656d6f746520746f6b656e20646563696d616c73604482015268040e8dede40d0d2ced60bb1b60648201526084016101be565b6001810154602083015160ff908116600160a01b9092041614612bf25760405162461bcd60e51b815260206004820152602660248201527f546f6b656e486f6d653a20696e76616c696420686f6d6520746f6b656e20646560448201526563696d616c7360d01b60648201526084016101be565b5f5f612c138360010160149054906101000a900460ff168560400151613755565b915091505f612c268383875f01516135d7565b9050818015612c4057508451612c3d908490614943565b15155b15612c5357612c50600182614830565b90505b6040518060800160405280600115158152602001828152602001848152602001831515815250846002015f8981526020019081526020015f205f886001600160a01b03166001600160a01b031681526020019081526020015f205f820151815f015f6101000a81548160ff02191690831515021790555060208201518160010155604082015181600201556060820151816003015f6101000a81548160ff021916908315150217905550905050856001600160a01b0316877ff229b02a51a4c8d5ef03a096ae0dd727d7b48b710d21b50ebebb560eef739b90838860400151604051612d4c92919091825260ff16602082015260400190565b60405180910390a350505050505050565b6040516370a0823160e01b81523060048201525f9081906001600160a01b038616906370a0823190602401602060405180830381865afa158015612da3573d5f5f3e3d5ffd5b505050506040513d601f19601f82011682018060405250810190612dc791906143da565b9050612dde6001600160a01b03861685308661379d565b6040516370a0823160e01b81523060048201525f906001600160a01b038716906370a0823190602401602060405180830381865afa158015612e22573d5f5f3e3d5ffd5b505050506040513d601f19601f82011682018060405250810190612e4691906143da565b9050818111612eac5760405162461bcd60e51b815260206004820152602c60248201527f5361666545524332305472616e7366657246726f6d3a2062616c616e6365206e60448201526b1bdd081a5b98dc99585cd95960a21b60648201526084016101be565b6123a682826143c7565b5f612ec48484846001613804565b949350505050565b5f8181527f63108c0ddc3d3ab0f279f27d4b9ca8220d5efbdbeeb537d4327d5cf9e982b10160205260408120545f516020614b085f395f51905f52906001600160a01b031680612f715760405162461bcd60e51b815260206004820152602a60248201527f54656c65706f72746572526f757465644170703a206e6f206f7574626f756e646044820152691036b2b9b9b2b733b2b960b11b60648201526084016101be565b6001600160a01b0381165f90815260038301602052604090205460ff16156120f35760405162461bcd60e51b81526004016101be90613fc3565b604051636eb1769f60e11b81523060048201526001600160a01b0383811660248301525f919085169063dd62ed3e90604401602060405180830381865afa158015612ff8573d5f5f3e3d5ffd5b505050506040513d601f19601f8201168201806040525081019061301c91906143da565b9050610a2c848461302d8585614830565b61382b565b61303a6120fa565b613043826138ba565b6108fd816138db565b6130546120fa565b6107c86138ec565b6130646120fa565b6001600160a01b0382166130ba5760405162461bcd60e51b815260206004820152601d60248201527f546f6b656e486f6d653a207a65726f20746f6b656e206164647265737300000060448201526064016101be565b60128160ff1611156131195760405162461bcd60e51b815260206004820152602260248201527f546f6b656e486f6d653a20746f6b656e20646563696d616c7320746f6f2068696044820152610ced60f31b60648201526084016101be565b7f9316912b5a9db88acbe872c934fdd0a46c436c6dcba332d649c4d57c7bc9e601805460ff92909216600160a01b026001600160a81b03199092166001600160a01b039390931692909217179055565b6040516370a0823160e01b81523060048201525f9081906001600160a01b038516906370a0823190602401602060405180830381865afa1580156131af573d5f5f3e3d5ffd5b505050506040513d601f19601f820116820180604052508101906131d391906143da565b9050836001600160a01b031663d0e30db0846040518263ffffffff1660e01b81526004015f604051808303818588803b15801561320e575f5ffd5b505af1158015613220573d5f5f3e3d5ffd5b50506040516370a0823160e01b81523060048201525f93506001600160a01b03881692506370a082319150602401602060405180830381865afa158015613269573d5f5f3e3d5ffd5b505050506040513d601f19601f8201168201806040525081019061328d91906143da565b90508181116132fb5760405162461bcd60e51b815260206004820152603460248201527f53616665577261707065644e6174697665546f6b656e4465706f7369743a2062604482015273185b185b98d9481b9bdd081a5b98dc99585cd95960621b60648201526084016101be565b61330582826143c7565b95945050505050565b804710156133385760405163cf47918160e01b8152476004820152602481018290526044016101be565b5f826001600160a01b0316826040515f6040518083038185875af1925050503d805f8114613381576040519150601f19603f3d011682016040523d82523d5f602084013e613386565b606091505b505090508061101a5760405163d6bda27560e01b815260040160405180910390fd5b83515f906133c85760405162461bcd60e51b81526004016101be9061437e565b6020850151156134265760405162461bcd60e51b8152602060048201526024808201527f546f6b656e486f6d653a2072656d6f7465206e6f7420636f6c6c61746572616c6044820152631a5e995960e21b60648201526084016101be565b613431848484613900565b5f61344586604001518760600151856135d7565b90505f81116133055760405162461bcd60e51b815260206004820152601c60248201527f546f6b656e486f6d653a207a65726f20746f6b656e20616d6f756e740000000060448201526064016101be565b5f845a10156134e75760405162461bcd60e51b815260206004820152601b60248201527f43616c6c5574696c733a20696e73756666696369656e7420676173000000000060448201526064016101be565b834710156135375760405162461bcd60e51b815260206004820152601d60248201527f43616c6c5574696c733a20696e73756666696369656e742076616c756500000060448201526064016101be565b826001600160a01b03163b5f0361354f57505f612ec4565b61271061355d603f87614956565b6135679087614830565b6135719190614830565b5a10156135c05760405162461bcd60e51b815260206004820152601b60248201527f43616c6c5574696c733a20696e73756666696369656e7420676173000000000060448201526064016101be565b5f5f5f84516020860188888bf19695505050505050565b5f612ec48484845f613804565b5f8481525f516020614ac85f395f51905f52602090815260408083206001600160a01b038716845282528083208151608081018352815460ff9081161580158352600184015495830195909552600283015493820193909352600390910154909116151560608201525f516020614a885f395f51905f52918061366a57505f8160200151115b15613679575f92505050612ec4565b8385116136dd5760405162461bcd60e51b815260206004820152602c60248201527f546f6b656e486f6d653a20696e73756666696369656e7420616d6f756e74207460448201526b6f20636f766572206665657360a01b60648201526084016101be565b6136e784866143c7565b94505f6136fd8260400151836060015188612eb6565b9050805f03613711575f9350505050612ec4565b5f88815260038401602090815260408083206001600160a01b038b16845290915281208054839290613744908490614830565b909155509098975050505050505050565b5f8060ff808516908416118181613778576137708587614969565b60ff16613786565b6137828686614969565b60ff165b61379190600a614a65565b96919550909350505050565b6040516001600160a01b038481166024830152838116604483015260648201839052610a2c9186918216906323b872dd906084015b604051602081830303815290604052915060e01b6020820180516001600160e01b0383818316178352505050506139ed565b5f811515841515036138215761381a8584614a70565b9050612ec4565b6133058584614956565b604080516001600160a01b038416602482015260448082018490528251808303909101815260649091019091526020810180516001600160e01b031663095ea7b360e01b17905261387c8482613a59565b610a2c576040516001600160a01b0384811660248301525f60448301526138b091869182169063095ea7b3906064016137d2565b610a2c84826139ed565b6138c26120fa565b6138ca613a9e565b6138d2613aae565b6107b481613ab6565b6138e36120fa565b6107b481613b2c565b5f5f516020614ae85f395f51905f52611ae5565b5f8381527f9316912b5a9db88acbe872c934fdd0a46c436c6dcba332d649c4d57c7bc9e603602090815260408083206001600160a01b03861684529091529020545f516020614a885f395f51905f5290828110156139b75760405162461bcd60e51b815260206004820152602e60248201527f546f6b656e486f6d653a20696e73756666696369656e7420746f6b656e20747260448201526d616e736665722062616c616e636560901b60648201526084016101be565b6139c183826143c7565b5f9586526003909201602090815260408087206001600160a01b03909616875294905250919092205550565b5f5f60205f8451602086015f885af180613a0c576040513d5f823e3d81fd5b50505f513d91508115613a23578060011415613a30565b6001600160a01b0384163b155b15610a2c57604051635274afe760e01b81526001600160a01b03851660048201526024016101be565b5f5f5f5f60205f8651602088015f8a5af192503d91505f5190508280156123a657508115613a8a57806001146123a6565b50505050506001600160a01b03163b151590565b613aa66120fa565b6107c8613b34565b6107c86120fa565b613abe6120fa565b80613b1b5760405162461bcd60e51b815260206004820152602760248201527f54656c65706f72746572526f757465644170703a207a65726f20626c6f636b636044820152661a185a5b88125160ca1b60648201526084016101be565b5f516020614b085f395f51905f5255565b610b4e6120fa565b611ac26120fa565b6001600160a01b03811681146107b4575f5ffd5b8035613b5b81613b3c565b919050565b5f60208284031215613b70575f5ffd5b81356120f381613b3c565b5f5f60408385031215613b8c575f5ffd5b823591506020830135613b9e81613b3c565b809150509250929050565b5f60208284031215613bb9575f5ffd5b81356001600160401b03811115613bce575f5ffd5b820161016081850312156120f3575f5ffd5b5f5f5f60608486031215613bf2575f5ffd5b833592506020840135613c0481613b3c565b91506040840135613c1481613b3c565b809150509250925092565b5f610100828403128015613c31575f5ffd5b509092915050565b5f5f5f5f60608587031215613c4c575f5ffd5b843593506020850135613c5e81613b3c565b925060408501356001600160401b03811115613c78575f5ffd5b8501601f81018713613c88575f5ffd5b80356001600160401b03811115613c9d575f5ffd5b876020828401011115613cae575f5ffd5b949793965060200194505050565b5f60208284031215613ccc575f5ffd5b5035919050565b6020808252602b908201527f54656c65706f72746572526f757465644170703a207a65726f206d657373656e60408201526a676572206164647265737360a81b606082015260800190565b634e487b7160e01b5f52604160045260245ffd5b60405161016081016001600160401b0381118282101715613d5557613d55613d1e565b60405290565b60405161010081016001600160401b0381118282101715613d5557613d55613d1e565b604080519081016001600160401b0381118282101715613d5557613d55613d1e565b60405160e081016001600160401b0381118282101715613d5557613d55613d1e565b604051601f8201601f191681016001600160401b0381118282101715613dea57613dea613d1e565b604052919050565b5f6001600160401b03821115613e0a57613e0a613d1e565b50601f01601f191660200190565b5f82601f830112613e27575f5ffd5b8135613e3a613e3582613df2565b613dc2565b818152846020838601011115613e4e575f5ffd5b816020850160208301375f918101602001919091529392505050565b5f6101608236031215613e7b575f5ffd5b613e83613d32565b82358152613e9360208401613b50565b6020820152613ea460408401613b50565b604082015260608301356001600160401b03811115613ec1575f5ffd5b613ecd36828601613e18565b6060830152506080838101359082015260a08084013590820152613ef360c08401613b50565b60c0820152613f0460e08401613b50565b60e0820152613f166101008401613b50565b610100820152610120838101359082015261014092830135928101929092525090565b5f610100828403128015613f4b575f5ffd5b50613f54613d5b565b823581526020830135613f6681613b3c565b6020820152613f7760408401613b50565b6040820152613f8860608401613b50565b60608201526080838101359082015260a0808401359082015260c08084013590820152613fb760e08401613b50565b60e08201529392505050565b60208082526025908201527f54656c65706f72746572526f757465644170703a206d657373656e6765722070604082015264185d5cd95960da1b606082015260800190565b60208082526028908201527f54656c65706f72746572526f757465644170703a206d657373656e676572206e6040820152671bdd08185919195960c21b606082015260800190565b60208082526024908201527f53656e645265656e7472616e637947756172643a2073656e64207265656e7472604082015263616e637960e01b606082015260800190565b60208082526026908201527f546f6b656e486f6d653a206e6f6e2d7a65726f206d756c74692d686f702066616040820152656c6c6261636b60d01b606082015260800190565b634e487b7160e01b5f52602160045260245ffd5b5f5b838110156141085781810151838201526020016140f0565b50505f910152565b5f81518084526141278160208601602086016140ee565b601f01601f19169290920160200192915050565b602081528151602082015260018060a01b03602083015116604082015260018060a01b0360408301511660608201525f606083015161418560808401826001600160a01b03169052565b50608083015160a083015260a083015161010060c08401526141ab610120840182614110565b905060c084015160e084015260e08401516141d26101008501826001600160a01b03169052565b509392505050565b602081525f8251600581106141fd57634e487b7160e01b5f52602160045260245ffd5b806020840152506020830151604080840152612ec46060840182614110565b60408152825160408201525f602084015161424260608401826001600160a01b03169052565b5060408401516001600160a01b03166080830152606084015161016060a08401526142716101a0840182614110565b9050608085015160c084015260a085015160e084015260c08501516142a26101008501826001600160a01b03169052565b5060e08501516001600160a01b038116610120850152506101008501516001600160a01b0381166101408501525061012085015161016084015261014090940151610180830152506020015290565b81516001600160a01b031681526020808301519082015260408101610785565b825181526020808401516001600160a01b03908116918301919091526040808501518216908301526060808501518216908301526080808501519083015260a0808501519083015260c0808501519083015260e09384015116928101929092526101008201526101200190565b6020808252818101527f546f6b656e486f6d653a2072656d6f7465206e6f742072656769737465726564604082015260600190565b634e487b7160e01b5f52601160045260245ffd5b81810381811115610785576107856143b3565b5f602082840312156143ea575f5ffd5b5051919050565b5f82601f830112614400575f5ffd5b815161440e613e3582613df2565b818152846020838601011115614422575f5ffd5b612ec48260208301602087016140ee565b5f60208284031215614443575f5ffd5b81516001600160401b03811115614458575f5ffd5b820160408185031215614469575f5ffd5b614471613d7e565b81516005811061447f575f5ffd5b815260208201516001600160401b03811115614499575f5ffd5b6144a5868285016143f1565b602083015250949350505050565b8051613b5b81613b3c565b5f60408284031280156144cf575f5ffd5b506144d8613d7e565b82516144e381613b3c565b81526020928301519281019290925250919050565b5f60208284031215614508575f5ffd5b81516001600160401b0381111561451d575f5ffd5b8201610100818503121561452f575f5ffd5b614537613d5b565b81518152614547602083016144b3565b6020820152614558604083016144b3565b6040820152614569606083016144b3565b60608201526080828101519082015260a08201516001600160401b03811115614590575f5ffd5b61459c868285016143f1565b60a08301525060c082810151908201526145b860e083016144b3565b60e0820152949350505050565b5f60e08284031280156145d6575f5ffd5b506145df613da0565b8251815260208301516145f181613b3c565b6020820152604083015161460481613b3c565b6040820152606083810151908201526080808401519082015260a0808401519082015260c083015161463581613b3c565b60c08201529392505050565b5f60208284031215614651575f5ffd5b81516001600160401b03811115614666575f5ffd5b82016101608185031215614678575f5ffd5b614680613d32565b614689826144b3565b8152602082810151908201526146a1604083016144b3565b60408201526146b2606083016144b3565b60608201526080828101519082015260a08201516001600160401b038111156146d9575f5ffd5b6146e5868285016143f1565b60a08301525060c0828101519082015261470160e083016144b3565b60e0820152610100828101519082015261471e61012083016144b3565b61012082015261014091820151918101919091529392505050565b805160ff81168114613b5b575f5ffd5b5f606082840312801561475a575f5ffd5b50604051606081016001600160401b038111828210171561477d5761477d613d1e565b6040528251815261479060208401614739565b60208201526147a160408401614739565b60408201529392505050565b60208082526022908201527f546f6b656e486f6d653a207a65726f20726571756972656420676173206c696d6040820152611a5d60f21b606082015260800190565b60208082526021908201527f546f6b656e486f6d653a206e6f6e2d7a65726f207365636f6e646172792066656040820152606560f81b606082015260800190565b80820180821115610785576107856143b3565b602081525f61010082018351602084015260018060a01b036020850151166040840152604084015161488b606085018280516001600160a01b03168252602090810151910152565b50606084015160a0840152608084015160e060c0850152805191829052602001905f906101208501905b808310156148e05783516001600160a01b0316825260209384019360019390930192909101906148b5565b5060a0860151858203601f190160e087015292506123a68184614110565b8481526001600160a01b038481166020830152831660408201526080606082018190525f906123a690830184614110565b634e487b7160e01b5f52601260045260245ffd5b5f826149515761495161492f565b500690565b5f826149645761496461492f565b500490565b60ff8281168282160390811115610785576107856143b3565b6001815b60018411156149bd578085048111156149a1576149a16143b3565b60018416156149af57908102905b60019390931c928002614986565b935093915050565b5f826149d357506001610785565b816149df57505f610785565b81600181146149f557600281146149ff57614a1b565b6001915050610785565b60ff841115614a1057614a106143b3565b50506001821b610785565b5060208310610133831016604e8410600b8410161715614a3e575081810a610785565b614a4a5f198484614982565b805f1904821115614a5d57614a5d6143b3565b029392505050565b5f6120f383836149c5565b8082028115828204841417610785576107856143b356fe9316912b5a9db88acbe872c934fdd0a46c436c6dcba332d649c4d57c7bc9e6003b5030f10c94fcbdaa3022348ff0b82dbd4c0c71339e41ff59d0bdc92179d6009316912b5a9db88acbe872c934fdd0a46c436c6dcba332d649c4d57c7bc9e602d2f1ed38b7d242bfb8b41862afb813a15193219a4bc717f2056607593e6c750063108c0ddc3d3ab0f279f27d4b9ca8220d5efbdbeeb537d4327d5cf9e982b10063108c0ddc3d3ab0f279f27d4b9ca8220d5efbdbeeb537d4327d5cf9e982b102a2646970667358221220c343a3be67bbeeafc0b2abb563209778362c52c3e7e11db6fd1a9242177fb12864736f6c634300081e0033f0c57e16840df040f15088dc2f81fe391c3923bec73e23a9662efc9c229c6a00",
}

// NativeTokenHomeABI is the input ABI used to generate the binding from.
// Deprecated: Use NativeTokenHomeMetaData.ABI instead.
var NativeTokenHomeABI = NativeTokenHomeMetaData.ABI

// NativeTokenHomeBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use NativeTokenHomeMetaData.Bin instead.
var NativeTokenHomeBin = NativeTokenHomeMetaData.Bin

// DeployNativeTokenHome deploys a new Ethereum contract, binding an instance of NativeTokenHome to it.
func DeployNativeTokenHome(auth *bind.TransactOpts, backend bind.ContractBackend, blockchainID [32]byte, teleporterManager common.Address, wrappedTokenAddress common.Address) (common.Address, *types.Transaction, *NativeTokenHome, error) {
	parsed, err := NativeTokenHomeMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(NativeTokenHomeBin), backend, blockchainID, teleporterManager, wrappedTokenAddress)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &NativeTokenHome{NativeTokenHomeCaller: NativeTokenHomeCaller{contract: contract}, NativeTokenHomeTransactor: NativeTokenHomeTransactor{contract: contract}, NativeTokenHomeFilterer: NativeTokenHomeFilterer{contract: contract}}, nil
}

// NativeTokenHome is an auto generated Go binding around an Ethereum contract.
type NativeTokenHome struct {
	NativeTokenHomeCaller     // Read-only binding to the contract
	NativeTokenHomeTransactor // Write-only binding to the contract
	NativeTokenHomeFilterer   // Log filterer for contract events
}

// NativeTokenHomeCaller is an auto generated read-only Go binding around an Ethereum contract.
type NativeTokenHomeCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// NativeTokenHomeTransactor is an auto generated write-only Go binding around an Ethereum contract.
type NativeTokenHomeTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// NativeTokenHomeFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type NativeTokenHomeFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// NativeTokenHomeSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type NativeTokenHomeSession struct {
	Contract     *NativeTokenHome  // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// NativeTokenHomeCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type NativeTokenHomeCallerSession struct {
	Contract *NativeTokenHomeCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts          // Call options to use throughout this session
}

// NativeTokenHomeTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type NativeTokenHomeTransactorSession struct {
	Contract     *NativeTokenHomeTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts          // Transaction auth options to use throughout this session
}

// NativeTokenHomeRaw is an auto generated low-level Go binding around an Ethereum contract.
type NativeTokenHomeRaw struct {
	Contract *NativeTokenHome // Generic contract binding to access the raw methods on
}

// NativeTokenHomeCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type NativeTokenHomeCallerRaw struct {
	Contract *NativeTokenHomeCaller // Generic read-only contract binding to access the raw methods on
}

// NativeTokenHomeTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type NativeTokenHomeTransactorRaw struct {
	Contract *NativeTokenHomeTransactor // Generic write-only contract binding to access the raw methods on
}

// NewNativeTokenHome creates a new instance of NativeTokenHome, bound to a specific deployed contract.
func NewNativeTokenHome(address common.Address, backend bind.ContractBackend) (*NativeTokenHome, error) {
	contract, err := bindNativeTokenHome(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &NativeTokenHome{NativeTokenHomeCaller: NativeTokenHomeCaller{contract: contract}, NativeTokenHomeTransactor: NativeTokenHomeTransactor{contract: contract}, NativeTokenHomeFilterer: NativeTokenHomeFilterer{contract: contract}}, nil
}

// NewNativeTokenHomeCaller creates a new read-only instance of NativeTokenHome, bound to a specific deployed contract.
func NewNativeTokenHomeCaller(address common.Address, caller bind.ContractCaller) (*NativeTokenHomeCaller, error) {
	contract, err := bindNativeTokenHome(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &NativeTokenHomeCaller{contract: contract}, nil
}

// NewNativeTokenHomeTransactor creates a new write-only instance of NativeTokenHome, bound to a specific deployed contract.
func NewNativeTokenHomeTransactor(address common.Address, transactor bind.ContractTransactor) (*NativeTokenHomeTransactor, error) {
	contract, err := bindNativeTokenHome(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &NativeTokenHomeTransactor{contract: contract}, nil
}

// NewNativeTokenHomeFilterer creates a new log filterer instance of NativeTokenHome, bound to a specific deployed contract.
func NewNativeTokenHomeFilterer(address common.Address, filterer bind.ContractFilterer) (*NativeTokenHomeFilterer, error) {
	contract, err := bindNativeTokenHome(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &NativeTokenHomeFilterer{contract: contract}, nil
}

// bindNativeTokenHome binds a generic wrapper to an already deployed contract.
func bindNativeTokenHome(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := NativeTokenHomeMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_NativeTokenHome *NativeTokenHomeRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _NativeTokenHome.Contract.NativeTokenHomeCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_NativeTokenHome *NativeTokenHomeRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _NativeTokenHome.Contract.NativeTokenHomeTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_NativeTokenHome *NativeTokenHomeRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _NativeTokenHome.Contract.NativeTokenHomeTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_NativeTokenHome *NativeTokenHomeCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _NativeTokenHome.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_NativeTokenHome *NativeTokenHomeTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _NativeTokenHome.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_NativeTokenHome *NativeTokenHomeTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _NativeTokenHome.Contract.contract.Transact(opts, method, params...)
}

// NATIVETOKENHOMESTORAGELOCATION is a free data retrieval call binding the contract method 0xefb5b95e.
//
// Solidity: function NATIVE_TOKEN_HOME_STORAGE_LOCATION() view returns(bytes32)
func (_NativeTokenHome *NativeTokenHomeCaller) NATIVETOKENHOMESTORAGELOCATION(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _NativeTokenHome.contract.Call(opts, &out, "NATIVE_TOKEN_HOME_STORAGE_LOCATION")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// NATIVETOKENHOMESTORAGELOCATION is a free data retrieval call binding the contract method 0xefb5b95e.
//
// Solidity: function NATIVE_TOKEN_HOME_STORAGE_LOCATION() view returns(bytes32)
func (_NativeTokenHome *NativeTokenHomeSession) NATIVETOKENHOMESTORAGELOCATION() ([32]byte, error) {
	return _NativeTokenHome.Contract.NATIVETOKENHOMESTORAGELOCATION(&_NativeTokenHome.CallOpts)
}

// NATIVETOKENHOMESTORAGELOCATION is a free data retrieval call binding the contract method 0xefb5b95e.
//
// Solidity: function NATIVE_TOKEN_HOME_STORAGE_LOCATION() view returns(bytes32)
func (_NativeTokenHome *NativeTokenHomeCallerSession) NATIVETOKENHOMESTORAGELOCATION() ([32]byte, error) {
	return _NativeTokenHome.Contract.NATIVETOKENHOMESTORAGELOCATION(&_NativeTokenHome.CallOpts)
}

// TELEPORTERROUTEDAPPSTORAGELOCATION is a free data retrieval call binding the contract method 0xc1c6353e.
//
// Solidity: function TELEPORTER_ROUTED_APP_STORAGE_LOCATION() view returns(bytes32)
func (_NativeTokenHome *NativeTokenHomeCaller) TELEPORTERROUTEDAPPSTORAGELOCATION(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _NativeTokenHome.contract.Call(opts, &out, "TELEPORTER_ROUTED_APP_STORAGE_LOCATION")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// TELEPORTERROUTEDAPPSTORAGELOCATION is a free data retrieval call binding the contract method 0xc1c6353e.
//
// Solidity: function TELEPORTER_ROUTED_APP_STORAGE_LOCATION() view returns(bytes32)
func (_NativeTokenHome *NativeTokenHomeSession) TELEPORTERROUTEDAPPSTORAGELOCATION() ([32]byte, error) {
	return _NativeTokenHome.Contract.TELEPORTERROUTEDAPPSTORAGELOCATION(&_NativeTokenHome.CallOpts)
}

// TELEPORTERROUTEDAPPSTORAGELOCATION is a free data retrieval call binding the contract method 0xc1c6353e.
//
// Solidity: function TELEPORTER_ROUTED_APP_STORAGE_LOCATION() view returns(bytes32)
func (_NativeTokenHome *NativeTokenHomeCallerSession) TELEPORTERROUTEDAPPSTORAGELOCATION() ([32]byte, error) {
	return _NativeTokenHome.Contract.TELEPORTERROUTEDAPPSTORAGELOCATION(&_NativeTokenHome.CallOpts)
}

// TOKENHOMESTORAGELOCATION is a free data retrieval call binding the contract method 0x62e3901b.
//
// Solidity: function TOKEN_HOME_STORAGE_LOCATION() view returns(bytes32)
func (_NativeTokenHome *NativeTokenHomeCaller) TOKENHOMESTORAGELOCATION(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _NativeTokenHome.contract.Call(opts, &out, "TOKEN_HOME_STORAGE_LOCATION")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// TOKENHOMESTORAGELOCATION is a free data retrieval call binding the contract method 0x62e3901b.
//
// Solidity: function TOKEN_HOME_STORAGE_LOCATION() view returns(bytes32)
func (_NativeTokenHome *NativeTokenHomeSession) TOKENHOMESTORAGELOCATION() ([32]byte, error) {
	return _NativeTokenHome.Contract.TOKENHOMESTORAGELOCATION(&_NativeTokenHome.CallOpts)
}

// TOKENHOMESTORAGELOCATION is a free data retrieval call binding the contract method 0x62e3901b.
//
// Solidity: function TOKEN_HOME_STORAGE_LOCATION() view returns(bytes32)
func (_NativeTokenHome *NativeTokenHomeCallerSession) TOKENHOMESTORAGELOCATION() ([32]byte, error) {
	return _NativeTokenHome.Contract.TOKENHOMESTORAGELOCATION(&_NativeTokenHome.CallOpts)
}

// GetBlockchainID is a free data retrieval call binding the contract method 0x4213cf78.
//
// Solidity: function getBlockchainID() view returns(bytes32)
func (_NativeTokenHome *NativeTokenHomeCaller) GetBlockchainID(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _NativeTokenHome.contract.Call(opts, &out, "getBlockchainID")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetBlockchainID is a free data retrieval call binding the contract method 0x4213cf78.
//
// Solidity: function getBlockchainID() view returns(bytes32)
func (_NativeTokenHome *NativeTokenHomeSession) GetBlockchainID() ([32]byte, error) {
	return _NativeTokenHome.Contract.GetBlockchainID(&_NativeTokenHome.CallOpts)
}

// GetBlockchainID is a free data retrieval call binding the contract method 0x4213cf78.
//
// Solidity: function getBlockchainID() view returns(bytes32)
func (_NativeTokenHome *NativeTokenHomeCallerSession) GetBlockchainID() ([32]byte, error) {
	return _NativeTokenHome.Contract.GetBlockchainID(&_NativeTokenHome.CallOpts)
}

// GetOutboundMessenger is a free data retrieval call binding the contract method 0xfc560821.
//
// Solidity: function getOutboundMessenger(bytes32 remoteBlockchainID) view returns(address)
func (_NativeTokenHome *NativeTokenHomeCaller) GetOutboundMessenger(opts *bind.CallOpts, remoteBlockchainID [32]byte) (common.Address, error) {
	var out []interface{}
	err := _NativeTokenHome.contract.Call(opts, &out, "getOutboundMessenger", remoteBlockchainID)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetOutboundMessenger is a free data retrieval call binding the contract method 0xfc560821.
//
// Solidity: function getOutboundMessenger(bytes32 remoteBlockchainID) view returns(address)
func (_NativeTokenHome *NativeTokenHomeSession) GetOutboundMessenger(remoteBlockchainID [32]byte) (common.Address, error) {
	return _NativeTokenHome.Contract.GetOutboundMessenger(&_NativeTokenHome.CallOpts, remoteBlockchainID)
}

// GetOutboundMessenger is a free data retrieval call binding the contract method 0xfc560821.
//
// Solidity: function getOutboundMessenger(bytes32 remoteBlockchainID) view returns(address)
func (_NativeTokenHome *NativeTokenHomeCallerSession) GetOutboundMessenger(remoteBlockchainID [32]byte) (common.Address, error) {
	return _NativeTokenHome.Contract.GetOutboundMessenger(&_NativeTokenHome.CallOpts, remoteBlockchainID)
}

// GetRemoteTokenTransferrerSettings is a free data retrieval call binding the contract method 0xc8511ada.
//
// Solidity: function getRemoteTokenTransferrerSettings(bytes32 remoteBlockchainID, address remoteTokenTransferrerAddress) view returns((bool,uint256,uint256,bool))
func (_NativeTokenHome *NativeTokenHomeCaller) GetRemoteTokenTransferrerSettings(opts *bind.CallOpts, remoteBlockchainID [32]byte, remoteTokenTransferrerAddress common.Address) (RemoteTokenTransferrerSettings, error) {
	var out []interface{}
	err := _NativeTokenHome.contract.Call(opts, &out, "getRemoteTokenTransferrerSettings", remoteBlockchainID, remoteTokenTransferrerAddress)

	if err != nil {
		return *new(RemoteTokenTransferrerSettings), err
	}

	out0 := *abi.ConvertType(out[0], new(RemoteTokenTransferrerSettings)).(*RemoteTokenTransferrerSettings)

	return out0, err

}

// GetRemoteTokenTransferrerSettings is a free data retrieval call binding the contract method 0xc8511ada.
//
// Solidity: function getRemoteTokenTransferrerSettings(bytes32 remoteBlockchainID, address remoteTokenTransferrerAddress) view returns((bool,uint256,uint256,bool))
func (_NativeTokenHome *NativeTokenHomeSession) GetRemoteTokenTransferrerSettings(remoteBlockchainID [32]byte, remoteTokenTransferrerAddress common.Address) (RemoteTokenTransferrerSettings, error) {
	return _NativeTokenHome.Contract.GetRemoteTokenTransferrerSettings(&_NativeTokenHome.CallOpts, remoteBlockchainID, remoteTokenTransferrerAddress)
}

// GetRemoteTokenTransferrerSettings is a free data retrieval call binding the contract method 0xc8511ada.
//
// Solidity: function getRemoteTokenTransferrerSettings(bytes32 remoteBlockchainID, address remoteTokenTransferrerAddress) view returns((bool,uint256,uint256,bool))
func (_NativeTokenHome *NativeTokenHomeCallerSession) GetRemoteTokenTransferrerSettings(remoteBlockchainID [32]byte, remoteTokenTransferrerAddress common.Address) (RemoteTokenTransferrerSettings, error) {
	return _NativeTokenHome.Contract.GetRemoteTokenTransferrerSettings(&_NativeTokenHome.CallOpts, remoteBlockchainID, remoteTokenTransferrerAddress)
}

// GetTokenAddress is a free data retrieval call binding the contract method 0x10fe9ae8.
//
// Solidity: function getTokenAddress() view returns(address)
func (_NativeTokenHome *NativeTokenHomeCaller) GetTokenAddress(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _NativeTokenHome.contract.Call(opts, &out, "getTokenAddress")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetTokenAddress is a free data retrieval call binding the contract method 0x10fe9ae8.
//
// Solidity: function getTokenAddress() view returns(address)
func (_NativeTokenHome *NativeTokenHomeSession) GetTokenAddress() (common.Address, error) {
	return _NativeTokenHome.Contract.GetTokenAddress(&_NativeTokenHome.CallOpts)
}

// GetTokenAddress is a free data retrieval call binding the contract method 0x10fe9ae8.
//
// Solidity: function getTokenAddress() view returns(address)
func (_NativeTokenHome *NativeTokenHomeCallerSession) GetTokenAddress() (common.Address, error) {
	return _NativeTokenHome.Contract.GetTokenAddress(&_NativeTokenHome.CallOpts)
}

// GetTransferredBalance is a free data retrieval call binding the contract method 0x154d625a.
//
// Solidity: function getTransferredBalance(bytes32 remoteBlockchainID, address remoteTokenTransferrerAddress) view returns(uint256)
func (_NativeTokenHome *NativeTokenHomeCaller) GetTransferredBalance(opts *bind.CallOpts, remoteBlockchainID [32]byte, remoteTokenTransferrerAddress common.Address) (*big.Int, error) {
	var out []interface{}
	err := _NativeTokenHome.contract.Call(opts, &out, "getTransferredBalance", remoteBlockchainID, remoteTokenTransferrerAddress)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetTransferredBalance is a free data retrieval call binding the contract method 0x154d625a.
//
// Solidity: function getTransferredBalance(bytes32 remoteBlockchainID, address remoteTokenTransferrerAddress) view returns(uint256)
func (_NativeTokenHome *NativeTokenHomeSession) GetTransferredBalance(remoteBlockchainID [32]byte, remoteTokenTransferrerAddress common.Address) (*big.Int, error) {
	return _NativeTokenHome.Contract.GetTransferredBalance(&_NativeTokenHome.CallOpts, remoteBlockchainID, remoteTokenTransferrerAddress)
}

// GetTransferredBalance is a free data retrieval call binding the contract method 0x154d625a.
//
// Solidity: function getTransferredBalance(bytes32 remoteBlockchainID, address remoteTokenTransferrerAddress) view returns(uint256)
func (_NativeTokenHome *NativeTokenHomeCallerSession) GetTransferredBalance(remoteBlockchainID [32]byte, remoteTokenTransferrerAddress common.Address) (*big.Int, error) {
	return _NativeTokenHome.Contract.GetTransferredBalance(&_NativeTokenHome.CallOpts, remoteBlockchainID, remoteTokenTransferrerAddress)
}

// IsInboundMessenger is a free data retrieval call binding the contract method 0x95a19443.
//
// Solidity: function isInboundMessenger(bytes32 remoteBlockchainID, address messenger) view returns(bool)
func (_NativeTokenHome *NativeTokenHomeCaller) IsInboundMessenger(opts *bind.CallOpts, remoteBlockchainID [32]byte, messenger common.Address) (bool, error) {
	var out []interface{}
	err := _NativeTokenHome.contract.Call(opts, &out, "isInboundMessenger", remoteBlockchainID, messenger)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsInboundMessenger is a free data retrieval call binding the contract method 0x95a19443.
//
// Solidity: function isInboundMessenger(bytes32 remoteBlockchainID, address messenger) view returns(bool)
func (_NativeTokenHome *NativeTokenHomeSession) IsInboundMessenger(remoteBlockchainID [32]byte, messenger common.Address) (bool, error) {
	return _NativeTokenHome.Contract.IsInboundMessenger(&_NativeTokenHome.CallOpts, remoteBlockchainID, messenger)
}

// IsInboundMessenger is a free data retrieval call binding the contract method 0x95a19443.
//
// Solidity: function isInboundMessenger(bytes32 remoteBlockchainID, address messenger) view returns(bool)
func (_NativeTokenHome *NativeTokenHomeCallerSession) IsInboundMessenger(remoteBlockchainID [32]byte, messenger common.Address) (bool, error) {
	return _NativeTokenHome.Contract.IsInboundMessenger(&_NativeTokenHome.CallOpts, remoteBlockchainID, messenger)
}

// IsMessengerPaused is a free data retrieval call binding the contract method 0x83c002d0.
//
// Solidity: function isMessengerPaused(address messenger) view returns(bool)
func (_NativeTokenHome *NativeTokenHomeCaller) IsMessengerPaused(opts *bind.CallOpts, messenger common.Address) (bool, error) {
	var out []interface{}
	err := _NativeTokenHome.contract.Call(opts, &out, "isMessengerPaused", messenger)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsMessengerPaused is a free data retrieval call binding the contract method 0x83c002d0.
//
// Solidity: function isMessengerPaused(address messenger) view returns(bool)
func (_NativeTokenHome *NativeTokenHomeSession) IsMessengerPaused(messenger common.Address) (bool, error) {
	return _NativeTokenHome.Contract.IsMessengerPaused(&_NativeTokenHome.CallOpts, messenger)
}

// IsMessengerPaused is a free data retrieval call binding the contract method 0x83c002d0.
//
// Solidity: function isMessengerPaused(address messenger) view returns(bool)
func (_NativeTokenHome *NativeTokenHomeCallerSession) IsMessengerPaused(messenger common.Address) (bool, error) {
	return _NativeTokenHome.Contract.IsMessengerPaused(&_NativeTokenHome.CallOpts, messenger)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_NativeTokenHome *NativeTokenHomeCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _NativeTokenHome.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_NativeTokenHome *NativeTokenHomeSession) Owner() (common.Address, error) {
	return _NativeTokenHome.Contract.Owner(&_NativeTokenHome.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_NativeTokenHome *NativeTokenHomeCallerSession) Owner() (common.Address, error) {
	return _NativeTokenHome.Contract.Owner(&_NativeTokenHome.CallOpts)
}

// AddCollateral is a paid mutator transaction binding the contract method 0xb0b78b26.
//
// Solidity: function addCollateral(bytes32 remoteBlockchainID, address remoteTokenTransferrerAddress) payable returns()
func (_NativeTokenHome *NativeTokenHomeTransactor) AddCollateral(opts *bind.TransactOpts, remoteBlockchainID [32]byte, remoteTokenTransferrerAddress common.Address) (*types.Transaction, error) {
	return _NativeTokenHome.contract.Transact(opts, "addCollateral", remoteBlockchainID, remoteTokenTransferrerAddress)
}

// AddCollateral is a paid mutator transaction binding the contract method 0xb0b78b26.
//
// Solidity: function addCollateral(bytes32 remoteBlockchainID, address remoteTokenTransferrerAddress) payable returns()
func (_NativeTokenHome *NativeTokenHomeSession) AddCollateral(remoteBlockchainID [32]byte, remoteTokenTransferrerAddress common.Address) (*types.Transaction, error) {
	return _NativeTokenHome.Contract.AddCollateral(&_NativeTokenHome.TransactOpts, remoteBlockchainID, remoteTokenTransferrerAddress)
}

// AddCollateral is a paid mutator transaction binding the contract method 0xb0b78b26.
//
// Solidity: function addCollateral(bytes32 remoteBlockchainID, address remoteTokenTransferrerAddress) payable returns()
func (_NativeTokenHome *NativeTokenHomeTransactorSession) AddCollateral(remoteBlockchainID [32]byte, remoteTokenTransferrerAddress common.Address) (*types.Transaction, error) {
	return _NativeTokenHome.Contract.AddCollateral(&_NativeTokenHome.TransactOpts, remoteBlockchainID, remoteTokenTransferrerAddress)
}

// AddMessenger is a paid mutator transaction binding the contract method 0xb2432529.
//
// Solidity: function addMessenger(bytes32 remoteBlockchainID, address messenger) returns()
func (_NativeTokenHome *NativeTokenHomeTransactor) AddMessenger(opts *bind.TransactOpts, remoteBlockchainID [32]byte, messenger common.Address) (*types.Transaction, error) {
	return _NativeTokenHome.contract.Transact(opts, "addMessenger", remoteBlockchainID, messenger)
}

// AddMessenger is a paid mutator transaction binding the contract method 0xb2432529.
//
// Solidity: function addMessenger(bytes32 remoteBlockchainID, address messenger) returns()
func (_NativeTokenHome *NativeTokenHomeSession) AddMessenger(remoteBlockchainID [32]byte, messenger common.Address) (*types.Transaction, error) {
	return _NativeTokenHome.Contract.AddMessenger(&_NativeTokenHome.TransactOpts, remoteBlockchainID, messenger)
}

// AddMessenger is a paid mutator transaction binding the contract method 0xb2432529.
//
// Solidity: function addMessenger(bytes32 remoteBlockchainID, address messenger) returns()
func (_NativeTokenHome *NativeTokenHomeTransactorSession) AddMessenger(remoteBlockchainID [32]byte, messenger common.Address) (*types.Transaction, error) {
	return _NativeTokenHome.Contract.AddMessenger(&_NativeTokenHome.TransactOpts, remoteBlockchainID, messenger)
}

// Initialize is a paid mutator transaction binding the contract method 0x88b117b3.
//
// Solidity: function initialize(bytes32 blockchainID, address teleporterManager, address wrappedTokenAddress) returns()
func (_NativeTokenHome *NativeTokenHomeTransactor) Initialize(opts *bind.TransactOpts, blockchainID [32]byte, teleporterManager common.Address, wrappedTokenAddress common.Address) (*types.Transaction, error) {
	return _NativeTokenHome.contract.Transact(opts, "initialize", blockchainID, teleporterManager, wrappedTokenAddress)
}

// Initialize is a paid mutator transaction binding the contract method 0x88b117b3.
//
// Solidity: function initialize(bytes32 blockchainID, address teleporterManager, address wrappedTokenAddress) returns()
func (_NativeTokenHome *NativeTokenHomeSession) Initialize(blockchainID [32]byte, teleporterManager common.Address, wrappedTokenAddress common.Address) (*types.Transaction, error) {
	return _NativeTokenHome.Contract.Initialize(&_NativeTokenHome.TransactOpts, blockchainID, teleporterManager, wrappedTokenAddress)
}

// Initialize is a paid mutator transaction binding the contract method 0x88b117b3.
//
// Solidity: function initialize(bytes32 blockchainID, address teleporterManager, address wrappedTokenAddress) returns()
func (_NativeTokenHome *NativeTokenHomeTransactorSession) Initialize(blockchainID [32]byte, teleporterManager common.Address, wrappedTokenAddress common.Address) (*types.Transaction, error) {
	return _NativeTokenHome.Contract.Initialize(&_NativeTokenHome.TransactOpts, blockchainID, teleporterManager, wrappedTokenAddress)
}

// PauseMessenger is a paid mutator transaction binding the contract method 0x03f78951.
//
// Solidity: function pauseMessenger(address messenger) returns()
func (_NativeTokenHome *NativeTokenHomeTransactor) PauseMessenger(opts *bind.TransactOpts, messenger common.Address) (*types.Transaction, error) {
	return _NativeTokenHome.contract.Transact(opts, "pauseMessenger", messenger)
}

// PauseMessenger is a paid mutator transaction binding the contract method 0x03f78951.
//
// Solidity: function pauseMessenger(address messenger) returns()
func (_NativeTokenHome *NativeTokenHomeSession) PauseMessenger(messenger common.Address) (*types.Transaction, error) {
	return _NativeTokenHome.Contract.PauseMessenger(&_NativeTokenHome.TransactOpts, messenger)
}

// PauseMessenger is a paid mutator transaction binding the contract method 0x03f78951.
//
// Solidity: function pauseMessenger(address messenger) returns()
func (_NativeTokenHome *NativeTokenHomeTransactorSession) PauseMessenger(messenger common.Address) (*types.Transaction, error) {
	return _NativeTokenHome.Contract.PauseMessenger(&_NativeTokenHome.TransactOpts, messenger)
}

// ReceiveTeleporterMessage is a paid mutator transaction binding the contract method 0xc868efaa.
//
// Solidity: function receiveTeleporterMessage(bytes32 sourceBlockchainID, address originSenderAddress, bytes message) returns()
func (_NativeTokenHome *NativeTokenHomeTransactor) ReceiveTeleporterMessage(opts *bind.TransactOpts, sourceBlockchainID [32]byte, originSenderAddress common.Address, message []byte) (*types.Transaction, error) {
	return _NativeTokenHome.contract.Transact(opts, "receiveTeleporterMessage", sourceBlockchainID, originSenderAddress, message)
}

// ReceiveTeleporterMessage is a paid mutator transaction binding the contract method 0xc868efaa.
//
// Solidity: function receiveTeleporterMessage(bytes32 sourceBlockchainID, address originSenderAddress, bytes message) returns()
func (_NativeTokenHome *NativeTokenHomeSession) ReceiveTeleporterMessage(sourceBlockchainID [32]byte, originSenderAddress common.Address, message []byte) (*types.Transaction, error) {
	return _NativeTokenHome.Contract.ReceiveTeleporterMessage(&_NativeTokenHome.TransactOpts, sourceBlockchainID, originSenderAddress, message)
}

// ReceiveTeleporterMessage is a paid mutator transaction binding the contract method 0xc868efaa.
//
// Solidity: function receiveTeleporterMessage(bytes32 sourceBlockchainID, address originSenderAddress, bytes message) returns()
func (_NativeTokenHome *NativeTokenHomeTransactorSession) ReceiveTeleporterMessage(sourceBlockchainID [32]byte, originSenderAddress common.Address, message []byte) (*types.Transaction, error) {
	return _NativeTokenHome.Contract.ReceiveTeleporterMessage(&_NativeTokenHome.TransactOpts, sourceBlockchainID, originSenderAddress, message)
}

// RemoveMessenger is a paid mutator transaction binding the contract method 0xf4be94a7.
//
// Solidity: function removeMessenger(bytes32 remoteBlockchainID, address messenger) returns()
func (_NativeTokenHome *NativeTokenHomeTransactor) RemoveMessenger(opts *bind.TransactOpts, remoteBlockchainID [32]byte, messenger common.Address) (*types.Transaction, error) {
	return _NativeTokenHome.contract.Transact(opts, "removeMessenger", remoteBlockchainID, messenger)
}

// RemoveMessenger is a paid mutator transaction binding the contract method 0xf4be94a7.
//
// Solidity: function removeMessenger(bytes32 remoteBlockchainID, address messenger) returns()
func (_NativeTokenHome *NativeTokenHomeSession) RemoveMessenger(remoteBlockchainID [32]byte, messenger common.Address) (*types.Transaction, error) {
	return _NativeTokenHome.Contract.RemoveMessenger(&_NativeTokenHome.TransactOpts, remoteBlockchainID, messenger)
}

// RemoveMessenger is a paid mutator transaction binding the contract method 0xf4be94a7.
//
// Solidity: function removeMessenger(bytes32 remoteBlockchainID, address messenger) returns()
func (_NativeTokenHome *NativeTokenHomeTransactorSession) RemoveMessenger(remoteBlockchainID [32]byte, messenger common.Address) (*types.Transaction, error) {
	return _NativeTokenHome.Contract.RemoveMessenger(&_NativeTokenHome.TransactOpts, remoteBlockchainID, messenger)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_NativeTokenHome *NativeTokenHomeTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _NativeTokenHome.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_NativeTokenHome *NativeTokenHomeSession) RenounceOwnership() (*types.Transaction, error) {
	return _NativeTokenHome.Contract.RenounceOwnership(&_NativeTokenHome.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_NativeTokenHome *NativeTokenHomeTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _NativeTokenHome.Contract.RenounceOwnership(&_NativeTokenHome.TransactOpts)
}

// Send is a paid mutator transaction binding the contract method 0x8bf2fa94.
//
// Solidity: function send((bytes32,address,address,address,uint256,uint256,uint256,address) input) payable returns()
func (_NativeTokenHome *NativeTokenHomeTransactor) Send(opts *bind.TransactOpts, input SendTokensInput) (*types.Transaction, error) {
	return _NativeTokenHome.contract.Transact(opts, "send", input)
}

// Send is a paid mutator transaction binding the contract method 0x8bf2fa94.
//
// Solidity: function send((bytes32,address,address,address,uint256,uint256,uint256,address) input) payable returns()
func (_NativeTokenHome *NativeTokenHomeSession) Send(input SendTokensInput) (*types.Transaction, error) {
	return _NativeTokenHome.Contract.Send(&_NativeTokenHome.TransactOpts, input)
}

// Send is a paid mutator transaction binding the contract method 0x8bf2fa94.
//
// Solidity: function send((bytes32,address,address,address,uint256,uint256,uint256,address) input) payable returns()
func (_NativeTokenHome *NativeTokenHomeTransactorSession) Send(input SendTokensInput) (*types.Transaction, error) {
	return _NativeTokenHome.Contract.Send(&_NativeTokenHome.TransactOpts, input)
}

// SendAndCall is a paid mutator transaction binding the contract method 0x6e6eef8d.
//
// Solidity: function sendAndCall((bytes32,address,address,bytes,uint256,uint256,address,address,address,uint256,uint256) input) payable returns()
func (_NativeTokenHome *NativeTokenHomeTransactor) SendAndCall(opts *bind.TransactOpts, input SendAndCallInput) (*types.Transaction, error) {
	return _NativeTokenHome.contract.Transact(opts, "sendAndCall", input)
}

// SendAndCall is a paid mutator transaction binding the contract method 0x6e6eef8d.
//
// Solidity: function sendAndCall((bytes32,address,address,bytes,uint256,uint256,address,address,address,uint256,uint256) input) payable returns()
func (_NativeTokenHome *NativeTokenHomeSession) SendAndCall(input SendAndCallInput) (*types.Transaction, error) {
	return _NativeTokenHome.Contract.SendAndCall(&_NativeTokenHome.TransactOpts, input)
}

// SendAndCall is a paid mutator transaction binding the contract method 0x6e6eef8d.
//
// Solidity: function sendAndCall((bytes32,address,address,bytes,uint256,uint256,address,address,address,uint256,uint256) input) payable returns()
func (_NativeTokenHome *NativeTokenHomeTransactorSession) SendAndCall(input SendAndCallInput) (*types.Transaction, error) {
	return _NativeTokenHome.Contract.SendAndCall(&_NativeTokenHome.TransactOpts, input)
}

// SetOutboundMessenger is a paid mutator transaction binding the contract method 0xd90f1b10.
//
// Solidity: function setOutboundMessenger(bytes32 remoteBlockchainID, address messenger) returns()
func (_NativeTokenHome *NativeTokenHomeTransactor) SetOutboundMessenger(opts *bind.TransactOpts, remoteBlockchainID [32]byte, messenger common.Address) (*types.Transaction, error) {
	return _NativeTokenHome.contract.Transact(opts, "setOutboundMessenger", remoteBlockchainID, messenger)
}

// SetOutboundMessenger is a paid mutator transaction binding the contract method 0xd90f1b10.
//
// Solidity: function setOutboundMessenger(bytes32 remoteBlockchainID, address messenger) returns()
func (_NativeTokenHome *NativeTokenHomeSession) SetOutboundMessenger(remoteBlockchainID [32]byte, messenger common.Address) (*types.Transaction, error) {
	return _NativeTokenHome.Contract.SetOutboundMessenger(&_NativeTokenHome.TransactOpts, remoteBlockchainID, messenger)
}

// SetOutboundMessenger is a paid mutator transaction binding the contract method 0xd90f1b10.
//
// Solidity: function setOutboundMessenger(bytes32 remoteBlockchainID, address messenger) returns()
func (_NativeTokenHome *NativeTokenHomeTransactorSession) SetOutboundMessenger(remoteBlockchainID [32]byte, messenger common.Address) (*types.Transaction, error) {
	return _NativeTokenHome.Contract.SetOutboundMessenger(&_NativeTokenHome.TransactOpts, remoteBlockchainID, messenger)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_NativeTokenHome *NativeTokenHomeTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _NativeTokenHome.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_NativeTokenHome *NativeTokenHomeSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _NativeTokenHome.Contract.TransferOwnership(&_NativeTokenHome.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_NativeTokenHome *NativeTokenHomeTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _NativeTokenHome.Contract.TransferOwnership(&_NativeTokenHome.TransactOpts, newOwner)
}

// UnpauseMessenger is a paid mutator transaction binding the contract method 0xdd5a20fd.
//
// Solidity: function unpauseMessenger(address messenger) returns()
func (_NativeTokenHome *NativeTokenHomeTransactor) UnpauseMessenger(opts *bind.TransactOpts, messenger common.Address) (*types.Transaction, error) {
	return _NativeTokenHome.contract.Transact(opts, "unpauseMessenger", messenger)
}

// UnpauseMessenger is a paid mutator transaction binding the contract method 0xdd5a20fd.
//
// Solidity: function unpauseMessenger(address messenger) returns()
func (_NativeTokenHome *NativeTokenHomeSession) UnpauseMessenger(messenger common.Address) (*types.Transaction, error) {
	return _NativeTokenHome.Contract.UnpauseMessenger(&_NativeTokenHome.TransactOpts, messenger)
}

// UnpauseMessenger is a paid mutator transaction binding the contract method 0xdd5a20fd.
//
// Solidity: function unpauseMessenger(address messenger) returns()
func (_NativeTokenHome *NativeTokenHomeTransactorSession) UnpauseMessenger(messenger common.Address) (*types.Transaction, error) {
	return _NativeTokenHome.Contract.UnpauseMessenger(&_NativeTokenHome.TransactOpts, messenger)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_NativeTokenHome *NativeTokenHomeTransactor) Receive(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _NativeTokenHome.contract.RawTransact(opts, nil) // calldata is disallowed for receive function
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_NativeTokenHome *NativeTokenHomeSession) Receive() (*types.Transaction, error) {
	return _NativeTokenHome.Contract.Receive(&_NativeTokenHome.TransactOpts)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_NativeTokenHome *NativeTokenHomeTransactorSession) Receive() (*types.Transaction, error) {
	return _NativeTokenHome.Contract.Receive(&_NativeTokenHome.TransactOpts)
}

// NativeTokenHomeCallFailedIterator is returned from FilterCallFailed and is used to iterate over the raw logs and unpacked data for CallFailed events raised by the NativeTokenHome contract.
type NativeTokenHomeCallFailedIterator struct {
	Event *NativeTokenHomeCallFailed // Event containing the contract specifics and raw log

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
func (it *NativeTokenHomeCallFailedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NativeTokenHomeCallFailed)
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
		it.Event = new(NativeTokenHomeCallFailed)
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
func (it *NativeTokenHomeCallFailedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NativeTokenHomeCallFailedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NativeTokenHomeCallFailed represents a CallFailed event raised by the NativeTokenHome contract.
type NativeTokenHomeCallFailed struct {
	RecipientContract common.Address
	Amount            *big.Int
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterCallFailed is a free log retrieval operation binding the contract event 0xb9eaeae386d339f8115782f297a9e5f0e13fb587cd6b0d502f113cb8dd4d6cb0.
//
// Solidity: event CallFailed(address indexed recipientContract, uint256 amount)
func (_NativeTokenHome *NativeTokenHomeFilterer) FilterCallFailed(opts *bind.FilterOpts, recipientContract []common.Address) (*NativeTokenHomeCallFailedIterator, error) {

	var recipientContractRule []interface{}
	for _, recipientContractItem := range recipientContract {
		recipientContractRule = append(recipientContractRule, recipientContractItem)
	}

	logs, sub, err := _NativeTokenHome.contract.FilterLogs(opts, "CallFailed", recipientContractRule)
	if err != nil {
		return nil, err
	}
	return &NativeTokenHomeCallFailedIterator{contract: _NativeTokenHome.contract, event: "CallFailed", logs: logs, sub: sub}, nil
}

// WatchCallFailed is a free log subscription operation binding the contract event 0xb9eaeae386d339f8115782f297a9e5f0e13fb587cd6b0d502f113cb8dd4d6cb0.
//
// Solidity: event CallFailed(address indexed recipientContract, uint256 amount)
func (_NativeTokenHome *NativeTokenHomeFilterer) WatchCallFailed(opts *bind.WatchOpts, sink chan<- *NativeTokenHomeCallFailed, recipientContract []common.Address) (event.Subscription, error) {

	var recipientContractRule []interface{}
	for _, recipientContractItem := range recipientContract {
		recipientContractRule = append(recipientContractRule, recipientContractItem)
	}

	logs, sub, err := _NativeTokenHome.contract.WatchLogs(opts, "CallFailed", recipientContractRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NativeTokenHomeCallFailed)
				if err := _NativeTokenHome.contract.UnpackLog(event, "CallFailed", log); err != nil {
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
func (_NativeTokenHome *NativeTokenHomeFilterer) ParseCallFailed(log types.Log) (*NativeTokenHomeCallFailed, error) {
	event := new(NativeTokenHomeCallFailed)
	if err := _NativeTokenHome.contract.UnpackLog(event, "CallFailed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// NativeTokenHomeCallSucceededIterator is returned from FilterCallSucceeded and is used to iterate over the raw logs and unpacked data for CallSucceeded events raised by the NativeTokenHome contract.
type NativeTokenHomeCallSucceededIterator struct {
	Event *NativeTokenHomeCallSucceeded // Event containing the contract specifics and raw log

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
func (it *NativeTokenHomeCallSucceededIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NativeTokenHomeCallSucceeded)
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
		it.Event = new(NativeTokenHomeCallSucceeded)
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
func (it *NativeTokenHomeCallSucceededIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NativeTokenHomeCallSucceededIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NativeTokenHomeCallSucceeded represents a CallSucceeded event raised by the NativeTokenHome contract.
type NativeTokenHomeCallSucceeded struct {
	RecipientContract common.Address
	Amount            *big.Int
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterCallSucceeded is a free log retrieval operation binding the contract event 0x104deb555f67e63782bb817bc26c39050894645f9b9f29c4be8ae68d0e8b7ff4.
//
// Solidity: event CallSucceeded(address indexed recipientContract, uint256 amount)
func (_NativeTokenHome *NativeTokenHomeFilterer) FilterCallSucceeded(opts *bind.FilterOpts, recipientContract []common.Address) (*NativeTokenHomeCallSucceededIterator, error) {

	var recipientContractRule []interface{}
	for _, recipientContractItem := range recipientContract {
		recipientContractRule = append(recipientContractRule, recipientContractItem)
	}

	logs, sub, err := _NativeTokenHome.contract.FilterLogs(opts, "CallSucceeded", recipientContractRule)
	if err != nil {
		return nil, err
	}
	return &NativeTokenHomeCallSucceededIterator{contract: _NativeTokenHome.contract, event: "CallSucceeded", logs: logs, sub: sub}, nil
}

// WatchCallSucceeded is a free log subscription operation binding the contract event 0x104deb555f67e63782bb817bc26c39050894645f9b9f29c4be8ae68d0e8b7ff4.
//
// Solidity: event CallSucceeded(address indexed recipientContract, uint256 amount)
func (_NativeTokenHome *NativeTokenHomeFilterer) WatchCallSucceeded(opts *bind.WatchOpts, sink chan<- *NativeTokenHomeCallSucceeded, recipientContract []common.Address) (event.Subscription, error) {

	var recipientContractRule []interface{}
	for _, recipientContractItem := range recipientContract {
		recipientContractRule = append(recipientContractRule, recipientContractItem)
	}

	logs, sub, err := _NativeTokenHome.contract.WatchLogs(opts, "CallSucceeded", recipientContractRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NativeTokenHomeCallSucceeded)
				if err := _NativeTokenHome.contract.UnpackLog(event, "CallSucceeded", log); err != nil {
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
func (_NativeTokenHome *NativeTokenHomeFilterer) ParseCallSucceeded(log types.Log) (*NativeTokenHomeCallSucceeded, error) {
	event := new(NativeTokenHomeCallSucceeded)
	if err := _NativeTokenHome.contract.UnpackLog(event, "CallSucceeded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// NativeTokenHomeCollateralAddedIterator is returned from FilterCollateralAdded and is used to iterate over the raw logs and unpacked data for CollateralAdded events raised by the NativeTokenHome contract.
type NativeTokenHomeCollateralAddedIterator struct {
	Event *NativeTokenHomeCollateralAdded // Event containing the contract specifics and raw log

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
func (it *NativeTokenHomeCollateralAddedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NativeTokenHomeCollateralAdded)
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
		it.Event = new(NativeTokenHomeCollateralAdded)
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
func (it *NativeTokenHomeCollateralAddedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NativeTokenHomeCollateralAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NativeTokenHomeCollateralAdded represents a CollateralAdded event raised by the NativeTokenHome contract.
type NativeTokenHomeCollateralAdded struct {
	RemoteBlockchainID            [32]byte
	RemoteTokenTransferrerAddress common.Address
	Amount                        *big.Int
	Remaining                     *big.Int
	Raw                           types.Log // Blockchain specific contextual infos
}

// FilterCollateralAdded is a free log retrieval operation binding the contract event 0x6769a5f9bfc8b6e0db839ab981cbf9239274ae72d2d035081a9157d43bd33cb6.
//
// Solidity: event CollateralAdded(bytes32 indexed remoteBlockchainID, address indexed remoteTokenTransferrerAddress, uint256 amount, uint256 remaining)
func (_NativeTokenHome *NativeTokenHomeFilterer) FilterCollateralAdded(opts *bind.FilterOpts, remoteBlockchainID [][32]byte, remoteTokenTransferrerAddress []common.Address) (*NativeTokenHomeCollateralAddedIterator, error) {

	var remoteBlockchainIDRule []interface{}
	for _, remoteBlockchainIDItem := range remoteBlockchainID {
		remoteBlockchainIDRule = append(remoteBlockchainIDRule, remoteBlockchainIDItem)
	}
	var remoteTokenTransferrerAddressRule []interface{}
	for _, remoteTokenTransferrerAddressItem := range remoteTokenTransferrerAddress {
		remoteTokenTransferrerAddressRule = append(remoteTokenTransferrerAddressRule, remoteTokenTransferrerAddressItem)
	}

	logs, sub, err := _NativeTokenHome.contract.FilterLogs(opts, "CollateralAdded", remoteBlockchainIDRule, remoteTokenTransferrerAddressRule)
	if err != nil {
		return nil, err
	}
	return &NativeTokenHomeCollateralAddedIterator{contract: _NativeTokenHome.contract, event: "CollateralAdded", logs: logs, sub: sub}, nil
}

// WatchCollateralAdded is a free log subscription operation binding the contract event 0x6769a5f9bfc8b6e0db839ab981cbf9239274ae72d2d035081a9157d43bd33cb6.
//
// Solidity: event CollateralAdded(bytes32 indexed remoteBlockchainID, address indexed remoteTokenTransferrerAddress, uint256 amount, uint256 remaining)
func (_NativeTokenHome *NativeTokenHomeFilterer) WatchCollateralAdded(opts *bind.WatchOpts, sink chan<- *NativeTokenHomeCollateralAdded, remoteBlockchainID [][32]byte, remoteTokenTransferrerAddress []common.Address) (event.Subscription, error) {

	var remoteBlockchainIDRule []interface{}
	for _, remoteBlockchainIDItem := range remoteBlockchainID {
		remoteBlockchainIDRule = append(remoteBlockchainIDRule, remoteBlockchainIDItem)
	}
	var remoteTokenTransferrerAddressRule []interface{}
	for _, remoteTokenTransferrerAddressItem := range remoteTokenTransferrerAddress {
		remoteTokenTransferrerAddressRule = append(remoteTokenTransferrerAddressRule, remoteTokenTransferrerAddressItem)
	}

	logs, sub, err := _NativeTokenHome.contract.WatchLogs(opts, "CollateralAdded", remoteBlockchainIDRule, remoteTokenTransferrerAddressRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NativeTokenHomeCollateralAdded)
				if err := _NativeTokenHome.contract.UnpackLog(event, "CollateralAdded", log); err != nil {
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

// ParseCollateralAdded is a log parse operation binding the contract event 0x6769a5f9bfc8b6e0db839ab981cbf9239274ae72d2d035081a9157d43bd33cb6.
//
// Solidity: event CollateralAdded(bytes32 indexed remoteBlockchainID, address indexed remoteTokenTransferrerAddress, uint256 amount, uint256 remaining)
func (_NativeTokenHome *NativeTokenHomeFilterer) ParseCollateralAdded(log types.Log) (*NativeTokenHomeCollateralAdded, error) {
	event := new(NativeTokenHomeCollateralAdded)
	if err := _NativeTokenHome.contract.UnpackLog(event, "CollateralAdded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// NativeTokenHomeInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the NativeTokenHome contract.
type NativeTokenHomeInitializedIterator struct {
	Event *NativeTokenHomeInitialized // Event containing the contract specifics and raw log

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
func (it *NativeTokenHomeInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NativeTokenHomeInitialized)
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
		it.Event = new(NativeTokenHomeInitialized)
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
func (it *NativeTokenHomeInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NativeTokenHomeInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NativeTokenHomeInitialized represents a Initialized event raised by the NativeTokenHome contract.
type NativeTokenHomeInitialized struct {
	Version uint64
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_NativeTokenHome *NativeTokenHomeFilterer) FilterInitialized(opts *bind.FilterOpts) (*NativeTokenHomeInitializedIterator, error) {

	logs, sub, err := _NativeTokenHome.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &NativeTokenHomeInitializedIterator{contract: _NativeTokenHome.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_NativeTokenHome *NativeTokenHomeFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *NativeTokenHomeInitialized) (event.Subscription, error) {

	logs, sub, err := _NativeTokenHome.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NativeTokenHomeInitialized)
				if err := _NativeTokenHome.contract.UnpackLog(event, "Initialized", log); err != nil {
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
func (_NativeTokenHome *NativeTokenHomeFilterer) ParseInitialized(log types.Log) (*NativeTokenHomeInitialized, error) {
	event := new(NativeTokenHomeInitialized)
	if err := _NativeTokenHome.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// NativeTokenHomeMessengerAddedIterator is returned from FilterMessengerAdded and is used to iterate over the raw logs and unpacked data for MessengerAdded events raised by the NativeTokenHome contract.
type NativeTokenHomeMessengerAddedIterator struct {
	Event *NativeTokenHomeMessengerAdded // Event containing the contract specifics and raw log

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
func (it *NativeTokenHomeMessengerAddedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NativeTokenHomeMessengerAdded)
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
		it.Event = new(NativeTokenHomeMessengerAdded)
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
func (it *NativeTokenHomeMessengerAddedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NativeTokenHomeMessengerAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NativeTokenHomeMessengerAdded represents a MessengerAdded event raised by the NativeTokenHome contract.
type NativeTokenHomeMessengerAdded struct {
	RemoteBlockchainID [32]byte
	Messenger          common.Address
	Raw                types.Log // Blockchain specific contextual infos
}

// FilterMessengerAdded is a free log retrieval operation binding the contract event 0xcc2f89e16e8979d86d3c3b49c5646f3a013f76626a479400e1cb3596f3dc2b22.
//
// Solidity: event MessengerAdded(bytes32 indexed remoteBlockchainID, address indexed messenger)
func (_NativeTokenHome *NativeTokenHomeFilterer) FilterMessengerAdded(opts *bind.FilterOpts, remoteBlockchainID [][32]byte, messenger []common.Address) (*NativeTokenHomeMessengerAddedIterator, error) {

	var remoteBlockchainIDRule []interface{}
	for _, remoteBlockchainIDItem := range remoteBlockchainID {
		remoteBlockchainIDRule = append(remoteBlockchainIDRule, remoteBlockchainIDItem)
	}
	var messengerRule []interface{}
	for _, messengerItem := range messenger {
		messengerRule = append(messengerRule, messengerItem)
	}

	logs, sub, err := _NativeTokenHome.contract.FilterLogs(opts, "MessengerAdded", remoteBlockchainIDRule, messengerRule)
	if err != nil {
		return nil, err
	}
	return &NativeTokenHomeMessengerAddedIterator{contract: _NativeTokenHome.contract, event: "MessengerAdded", logs: logs, sub: sub}, nil
}

// WatchMessengerAdded is a free log subscription operation binding the contract event 0xcc2f89e16e8979d86d3c3b49c5646f3a013f76626a479400e1cb3596f3dc2b22.
//
// Solidity: event MessengerAdded(bytes32 indexed remoteBlockchainID, address indexed messenger)
func (_NativeTokenHome *NativeTokenHomeFilterer) WatchMessengerAdded(opts *bind.WatchOpts, sink chan<- *NativeTokenHomeMessengerAdded, remoteBlockchainID [][32]byte, messenger []common.Address) (event.Subscription, error) {

	var remoteBlockchainIDRule []interface{}
	for _, remoteBlockchainIDItem := range remoteBlockchainID {
		remoteBlockchainIDRule = append(remoteBlockchainIDRule, remoteBlockchainIDItem)
	}
	var messengerRule []interface{}
	for _, messengerItem := range messenger {
		messengerRule = append(messengerRule, messengerItem)
	}

	logs, sub, err := _NativeTokenHome.contract.WatchLogs(opts, "MessengerAdded", remoteBlockchainIDRule, messengerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NativeTokenHomeMessengerAdded)
				if err := _NativeTokenHome.contract.UnpackLog(event, "MessengerAdded", log); err != nil {
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
func (_NativeTokenHome *NativeTokenHomeFilterer) ParseMessengerAdded(log types.Log) (*NativeTokenHomeMessengerAdded, error) {
	event := new(NativeTokenHomeMessengerAdded)
	if err := _NativeTokenHome.contract.UnpackLog(event, "MessengerAdded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// NativeTokenHomeMessengerPausedIterator is returned from FilterMessengerPaused and is used to iterate over the raw logs and unpacked data for MessengerPaused events raised by the NativeTokenHome contract.
type NativeTokenHomeMessengerPausedIterator struct {
	Event *NativeTokenHomeMessengerPaused // Event containing the contract specifics and raw log

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
func (it *NativeTokenHomeMessengerPausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NativeTokenHomeMessengerPaused)
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
		it.Event = new(NativeTokenHomeMessengerPaused)
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
func (it *NativeTokenHomeMessengerPausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NativeTokenHomeMessengerPausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NativeTokenHomeMessengerPaused represents a MessengerPaused event raised by the NativeTokenHome contract.
type NativeTokenHomeMessengerPaused struct {
	Messenger common.Address
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterMessengerPaused is a free log retrieval operation binding the contract event 0x46def370ef3a1ec72618480f66f1063a80fe68b27d09575bc3562ad489079b61.
//
// Solidity: event MessengerPaused(address indexed messenger)
func (_NativeTokenHome *NativeTokenHomeFilterer) FilterMessengerPaused(opts *bind.FilterOpts, messenger []common.Address) (*NativeTokenHomeMessengerPausedIterator, error) {

	var messengerRule []interface{}
	for _, messengerItem := range messenger {
		messengerRule = append(messengerRule, messengerItem)
	}

	logs, sub, err := _NativeTokenHome.contract.FilterLogs(opts, "MessengerPaused", messengerRule)
	if err != nil {
		return nil, err
	}
	return &NativeTokenHomeMessengerPausedIterator{contract: _NativeTokenHome.contract, event: "MessengerPaused", logs: logs, sub: sub}, nil
}

// WatchMessengerPaused is a free log subscription operation binding the contract event 0x46def370ef3a1ec72618480f66f1063a80fe68b27d09575bc3562ad489079b61.
//
// Solidity: event MessengerPaused(address indexed messenger)
func (_NativeTokenHome *NativeTokenHomeFilterer) WatchMessengerPaused(opts *bind.WatchOpts, sink chan<- *NativeTokenHomeMessengerPaused, messenger []common.Address) (event.Subscription, error) {

	var messengerRule []interface{}
	for _, messengerItem := range messenger {
		messengerRule = append(messengerRule, messengerItem)
	}

	logs, sub, err := _NativeTokenHome.contract.WatchLogs(opts, "MessengerPaused", messengerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NativeTokenHomeMessengerPaused)
				if err := _NativeTokenHome.contract.UnpackLog(event, "MessengerPaused", log); err != nil {
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
func (_NativeTokenHome *NativeTokenHomeFilterer) ParseMessengerPaused(log types.Log) (*NativeTokenHomeMessengerPaused, error) {
	event := new(NativeTokenHomeMessengerPaused)
	if err := _NativeTokenHome.contract.UnpackLog(event, "MessengerPaused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// NativeTokenHomeMessengerRemovedIterator is returned from FilterMessengerRemoved and is used to iterate over the raw logs and unpacked data for MessengerRemoved events raised by the NativeTokenHome contract.
type NativeTokenHomeMessengerRemovedIterator struct {
	Event *NativeTokenHomeMessengerRemoved // Event containing the contract specifics and raw log

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
func (it *NativeTokenHomeMessengerRemovedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NativeTokenHomeMessengerRemoved)
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
		it.Event = new(NativeTokenHomeMessengerRemoved)
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
func (it *NativeTokenHomeMessengerRemovedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NativeTokenHomeMessengerRemovedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NativeTokenHomeMessengerRemoved represents a MessengerRemoved event raised by the NativeTokenHome contract.
type NativeTokenHomeMessengerRemoved struct {
	RemoteBlockchainID [32]byte
	Messenger          common.Address
	Raw                types.Log // Blockchain specific contextual infos
}

// FilterMessengerRemoved is a free log retrieval operation binding the contract event 0x8350b16da8b8ec7d7e6ce37ca2ccb4f63d628b5ad53215d072763b6afce99524.
//
// Solidity: event MessengerRemoved(bytes32 indexed remoteBlockchainID, address indexed messenger)
func (_NativeTokenHome *NativeTokenHomeFilterer) FilterMessengerRemoved(opts *bind.FilterOpts, remoteBlockchainID [][32]byte, messenger []common.Address) (*NativeTokenHomeMessengerRemovedIterator, error) {

	var remoteBlockchainIDRule []interface{}
	for _, remoteBlockchainIDItem := range remoteBlockchainID {
		remoteBlockchainIDRule = append(remoteBlockchainIDRule, remoteBlockchainIDItem)
	}
	var messengerRule []interface{}
	for _, messengerItem := range messenger {
		messengerRule = append(messengerRule, messengerItem)
	}

	logs, sub, err := _NativeTokenHome.contract.FilterLogs(opts, "MessengerRemoved", remoteBlockchainIDRule, messengerRule)
	if err != nil {
		return nil, err
	}
	return &NativeTokenHomeMessengerRemovedIterator{contract: _NativeTokenHome.contract, event: "MessengerRemoved", logs: logs, sub: sub}, nil
}

// WatchMessengerRemoved is a free log subscription operation binding the contract event 0x8350b16da8b8ec7d7e6ce37ca2ccb4f63d628b5ad53215d072763b6afce99524.
//
// Solidity: event MessengerRemoved(bytes32 indexed remoteBlockchainID, address indexed messenger)
func (_NativeTokenHome *NativeTokenHomeFilterer) WatchMessengerRemoved(opts *bind.WatchOpts, sink chan<- *NativeTokenHomeMessengerRemoved, remoteBlockchainID [][32]byte, messenger []common.Address) (event.Subscription, error) {

	var remoteBlockchainIDRule []interface{}
	for _, remoteBlockchainIDItem := range remoteBlockchainID {
		remoteBlockchainIDRule = append(remoteBlockchainIDRule, remoteBlockchainIDItem)
	}
	var messengerRule []interface{}
	for _, messengerItem := range messenger {
		messengerRule = append(messengerRule, messengerItem)
	}

	logs, sub, err := _NativeTokenHome.contract.WatchLogs(opts, "MessengerRemoved", remoteBlockchainIDRule, messengerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NativeTokenHomeMessengerRemoved)
				if err := _NativeTokenHome.contract.UnpackLog(event, "MessengerRemoved", log); err != nil {
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
func (_NativeTokenHome *NativeTokenHomeFilterer) ParseMessengerRemoved(log types.Log) (*NativeTokenHomeMessengerRemoved, error) {
	event := new(NativeTokenHomeMessengerRemoved)
	if err := _NativeTokenHome.contract.UnpackLog(event, "MessengerRemoved", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// NativeTokenHomeMessengerUnpausedIterator is returned from FilterMessengerUnpaused and is used to iterate over the raw logs and unpacked data for MessengerUnpaused events raised by the NativeTokenHome contract.
type NativeTokenHomeMessengerUnpausedIterator struct {
	Event *NativeTokenHomeMessengerUnpaused // Event containing the contract specifics and raw log

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
func (it *NativeTokenHomeMessengerUnpausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NativeTokenHomeMessengerUnpaused)
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
		it.Event = new(NativeTokenHomeMessengerUnpaused)
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
func (it *NativeTokenHomeMessengerUnpausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NativeTokenHomeMessengerUnpausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NativeTokenHomeMessengerUnpaused represents a MessengerUnpaused event raised by the NativeTokenHome contract.
type NativeTokenHomeMessengerUnpaused struct {
	Messenger common.Address
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterMessengerUnpaused is a free log retrieval operation binding the contract event 0xf5574433a9d82924af98e239704e51afb34ae8c15f513d78a22ae0c58f5f9965.
//
// Solidity: event MessengerUnpaused(address indexed messenger)
func (_NativeTokenHome *NativeTokenHomeFilterer) FilterMessengerUnpaused(opts *bind.FilterOpts, messenger []common.Address) (*NativeTokenHomeMessengerUnpausedIterator, error) {

	var messengerRule []interface{}
	for _, messengerItem := range messenger {
		messengerRule = append(messengerRule, messengerItem)
	}

	logs, sub, err := _NativeTokenHome.contract.FilterLogs(opts, "MessengerUnpaused", messengerRule)
	if err != nil {
		return nil, err
	}
	return &NativeTokenHomeMessengerUnpausedIterator{contract: _NativeTokenHome.contract, event: "MessengerUnpaused", logs: logs, sub: sub}, nil
}

// WatchMessengerUnpaused is a free log subscription operation binding the contract event 0xf5574433a9d82924af98e239704e51afb34ae8c15f513d78a22ae0c58f5f9965.
//
// Solidity: event MessengerUnpaused(address indexed messenger)
func (_NativeTokenHome *NativeTokenHomeFilterer) WatchMessengerUnpaused(opts *bind.WatchOpts, sink chan<- *NativeTokenHomeMessengerUnpaused, messenger []common.Address) (event.Subscription, error) {

	var messengerRule []interface{}
	for _, messengerItem := range messenger {
		messengerRule = append(messengerRule, messengerItem)
	}

	logs, sub, err := _NativeTokenHome.contract.WatchLogs(opts, "MessengerUnpaused", messengerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NativeTokenHomeMessengerUnpaused)
				if err := _NativeTokenHome.contract.UnpackLog(event, "MessengerUnpaused", log); err != nil {
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
func (_NativeTokenHome *NativeTokenHomeFilterer) ParseMessengerUnpaused(log types.Log) (*NativeTokenHomeMessengerUnpaused, error) {
	event := new(NativeTokenHomeMessengerUnpaused)
	if err := _NativeTokenHome.contract.UnpackLog(event, "MessengerUnpaused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// NativeTokenHomeOutboundMessengerSetIterator is returned from FilterOutboundMessengerSet and is used to iterate over the raw logs and unpacked data for OutboundMessengerSet events raised by the NativeTokenHome contract.
type NativeTokenHomeOutboundMessengerSetIterator struct {
	Event *NativeTokenHomeOutboundMessengerSet // Event containing the contract specifics and raw log

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
func (it *NativeTokenHomeOutboundMessengerSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NativeTokenHomeOutboundMessengerSet)
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
		it.Event = new(NativeTokenHomeOutboundMessengerSet)
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
func (it *NativeTokenHomeOutboundMessengerSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NativeTokenHomeOutboundMessengerSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NativeTokenHomeOutboundMessengerSet represents a OutboundMessengerSet event raised by the NativeTokenHome contract.
type NativeTokenHomeOutboundMessengerSet struct {
	RemoteBlockchainID [32]byte
	OldMessenger       common.Address
	NewMessenger       common.Address
	Raw                types.Log // Blockchain specific contextual infos
}

// FilterOutboundMessengerSet is a free log retrieval operation binding the contract event 0xd00498d0475cde985344ea08d024c9315ca654d9d157211295aac43496374115.
//
// Solidity: event OutboundMessengerSet(bytes32 indexed remoteBlockchainID, address indexed oldMessenger, address indexed newMessenger)
func (_NativeTokenHome *NativeTokenHomeFilterer) FilterOutboundMessengerSet(opts *bind.FilterOpts, remoteBlockchainID [][32]byte, oldMessenger []common.Address, newMessenger []common.Address) (*NativeTokenHomeOutboundMessengerSetIterator, error) {

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

	logs, sub, err := _NativeTokenHome.contract.FilterLogs(opts, "OutboundMessengerSet", remoteBlockchainIDRule, oldMessengerRule, newMessengerRule)
	if err != nil {
		return nil, err
	}
	return &NativeTokenHomeOutboundMessengerSetIterator{contract: _NativeTokenHome.contract, event: "OutboundMessengerSet", logs: logs, sub: sub}, nil
}

// WatchOutboundMessengerSet is a free log subscription operation binding the contract event 0xd00498d0475cde985344ea08d024c9315ca654d9d157211295aac43496374115.
//
// Solidity: event OutboundMessengerSet(bytes32 indexed remoteBlockchainID, address indexed oldMessenger, address indexed newMessenger)
func (_NativeTokenHome *NativeTokenHomeFilterer) WatchOutboundMessengerSet(opts *bind.WatchOpts, sink chan<- *NativeTokenHomeOutboundMessengerSet, remoteBlockchainID [][32]byte, oldMessenger []common.Address, newMessenger []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _NativeTokenHome.contract.WatchLogs(opts, "OutboundMessengerSet", remoteBlockchainIDRule, oldMessengerRule, newMessengerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NativeTokenHomeOutboundMessengerSet)
				if err := _NativeTokenHome.contract.UnpackLog(event, "OutboundMessengerSet", log); err != nil {
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
func (_NativeTokenHome *NativeTokenHomeFilterer) ParseOutboundMessengerSet(log types.Log) (*NativeTokenHomeOutboundMessengerSet, error) {
	event := new(NativeTokenHomeOutboundMessengerSet)
	if err := _NativeTokenHome.contract.UnpackLog(event, "OutboundMessengerSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// NativeTokenHomeOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the NativeTokenHome contract.
type NativeTokenHomeOwnershipTransferredIterator struct {
	Event *NativeTokenHomeOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *NativeTokenHomeOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NativeTokenHomeOwnershipTransferred)
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
		it.Event = new(NativeTokenHomeOwnershipTransferred)
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
func (it *NativeTokenHomeOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NativeTokenHomeOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NativeTokenHomeOwnershipTransferred represents a OwnershipTransferred event raised by the NativeTokenHome contract.
type NativeTokenHomeOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_NativeTokenHome *NativeTokenHomeFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*NativeTokenHomeOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _NativeTokenHome.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &NativeTokenHomeOwnershipTransferredIterator{contract: _NativeTokenHome.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_NativeTokenHome *NativeTokenHomeFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *NativeTokenHomeOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _NativeTokenHome.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NativeTokenHomeOwnershipTransferred)
				if err := _NativeTokenHome.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_NativeTokenHome *NativeTokenHomeFilterer) ParseOwnershipTransferred(log types.Log) (*NativeTokenHomeOwnershipTransferred, error) {
	event := new(NativeTokenHomeOwnershipTransferred)
	if err := _NativeTokenHome.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// NativeTokenHomeRemoteRegisteredIterator is returned from FilterRemoteRegistered and is used to iterate over the raw logs and unpacked data for RemoteRegistered events raised by the NativeTokenHome contract.
type NativeTokenHomeRemoteRegisteredIterator struct {
	Event *NativeTokenHomeRemoteRegistered // Event containing the contract specifics and raw log

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
func (it *NativeTokenHomeRemoteRegisteredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NativeTokenHomeRemoteRegistered)
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
		it.Event = new(NativeTokenHomeRemoteRegistered)
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
func (it *NativeTokenHomeRemoteRegisteredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NativeTokenHomeRemoteRegisteredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NativeTokenHomeRemoteRegistered represents a RemoteRegistered event raised by the NativeTokenHome contract.
type NativeTokenHomeRemoteRegistered struct {
	RemoteBlockchainID            [32]byte
	RemoteTokenTransferrerAddress common.Address
	InitialCollateralNeeded       *big.Int
	TokenDecimals                 uint8
	Raw                           types.Log // Blockchain specific contextual infos
}

// FilterRemoteRegistered is a free log retrieval operation binding the contract event 0xf229b02a51a4c8d5ef03a096ae0dd727d7b48b710d21b50ebebb560eef739b90.
//
// Solidity: event RemoteRegistered(bytes32 indexed remoteBlockchainID, address indexed remoteTokenTransferrerAddress, uint256 initialCollateralNeeded, uint8 tokenDecimals)
func (_NativeTokenHome *NativeTokenHomeFilterer) FilterRemoteRegistered(opts *bind.FilterOpts, remoteBlockchainID [][32]byte, remoteTokenTransferrerAddress []common.Address) (*NativeTokenHomeRemoteRegisteredIterator, error) {

	var remoteBlockchainIDRule []interface{}
	for _, remoteBlockchainIDItem := range remoteBlockchainID {
		remoteBlockchainIDRule = append(remoteBlockchainIDRule, remoteBlockchainIDItem)
	}
	var remoteTokenTransferrerAddressRule []interface{}
	for _, remoteTokenTransferrerAddressItem := range remoteTokenTransferrerAddress {
		remoteTokenTransferrerAddressRule = append(remoteTokenTransferrerAddressRule, remoteTokenTransferrerAddressItem)
	}

	logs, sub, err := _NativeTokenHome.contract.FilterLogs(opts, "RemoteRegistered", remoteBlockchainIDRule, remoteTokenTransferrerAddressRule)
	if err != nil {
		return nil, err
	}
	return &NativeTokenHomeRemoteRegisteredIterator{contract: _NativeTokenHome.contract, event: "RemoteRegistered", logs: logs, sub: sub}, nil
}

// WatchRemoteRegistered is a free log subscription operation binding the contract event 0xf229b02a51a4c8d5ef03a096ae0dd727d7b48b710d21b50ebebb560eef739b90.
//
// Solidity: event RemoteRegistered(bytes32 indexed remoteBlockchainID, address indexed remoteTokenTransferrerAddress, uint256 initialCollateralNeeded, uint8 tokenDecimals)
func (_NativeTokenHome *NativeTokenHomeFilterer) WatchRemoteRegistered(opts *bind.WatchOpts, sink chan<- *NativeTokenHomeRemoteRegistered, remoteBlockchainID [][32]byte, remoteTokenTransferrerAddress []common.Address) (event.Subscription, error) {

	var remoteBlockchainIDRule []interface{}
	for _, remoteBlockchainIDItem := range remoteBlockchainID {
		remoteBlockchainIDRule = append(remoteBlockchainIDRule, remoteBlockchainIDItem)
	}
	var remoteTokenTransferrerAddressRule []interface{}
	for _, remoteTokenTransferrerAddressItem := range remoteTokenTransferrerAddress {
		remoteTokenTransferrerAddressRule = append(remoteTokenTransferrerAddressRule, remoteTokenTransferrerAddressItem)
	}

	logs, sub, err := _NativeTokenHome.contract.WatchLogs(opts, "RemoteRegistered", remoteBlockchainIDRule, remoteTokenTransferrerAddressRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NativeTokenHomeRemoteRegistered)
				if err := _NativeTokenHome.contract.UnpackLog(event, "RemoteRegistered", log); err != nil {
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

// ParseRemoteRegistered is a log parse operation binding the contract event 0xf229b02a51a4c8d5ef03a096ae0dd727d7b48b710d21b50ebebb560eef739b90.
//
// Solidity: event RemoteRegistered(bytes32 indexed remoteBlockchainID, address indexed remoteTokenTransferrerAddress, uint256 initialCollateralNeeded, uint8 tokenDecimals)
func (_NativeTokenHome *NativeTokenHomeFilterer) ParseRemoteRegistered(log types.Log) (*NativeTokenHomeRemoteRegistered, error) {
	event := new(NativeTokenHomeRemoteRegistered)
	if err := _NativeTokenHome.contract.UnpackLog(event, "RemoteRegistered", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// NativeTokenHomeTokensAndCallRoutedIterator is returned from FilterTokensAndCallRouted and is used to iterate over the raw logs and unpacked data for TokensAndCallRouted events raised by the NativeTokenHome contract.
type NativeTokenHomeTokensAndCallRoutedIterator struct {
	Event *NativeTokenHomeTokensAndCallRouted // Event containing the contract specifics and raw log

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
func (it *NativeTokenHomeTokensAndCallRoutedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NativeTokenHomeTokensAndCallRouted)
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
		it.Event = new(NativeTokenHomeTokensAndCallRouted)
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
func (it *NativeTokenHomeTokensAndCallRoutedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NativeTokenHomeTokensAndCallRoutedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NativeTokenHomeTokensAndCallRouted represents a TokensAndCallRouted event raised by the NativeTokenHome contract.
type NativeTokenHomeTokensAndCallRouted struct {
	TeleporterMessageID [32]byte
	Input               SendAndCallInput
	Amount              *big.Int
	Raw                 types.Log // Blockchain specific contextual infos
}

// FilterTokensAndCallRouted is a free log retrieval operation binding the contract event 0x42eff9005856e3c586b096d67211a566dc926052119fd7cc08023c70937ecb30.
//
// Solidity: event TokensAndCallRouted(bytes32 indexed teleporterMessageID, (bytes32,address,address,bytes,uint256,uint256,address,address,address,uint256,uint256) input, uint256 amount)
func (_NativeTokenHome *NativeTokenHomeFilterer) FilterTokensAndCallRouted(opts *bind.FilterOpts, teleporterMessageID [][32]byte) (*NativeTokenHomeTokensAndCallRoutedIterator, error) {

	var teleporterMessageIDRule []interface{}
	for _, teleporterMessageIDItem := range teleporterMessageID {
		teleporterMessageIDRule = append(teleporterMessageIDRule, teleporterMessageIDItem)
	}

	logs, sub, err := _NativeTokenHome.contract.FilterLogs(opts, "TokensAndCallRouted", teleporterMessageIDRule)
	if err != nil {
		return nil, err
	}
	return &NativeTokenHomeTokensAndCallRoutedIterator{contract: _NativeTokenHome.contract, event: "TokensAndCallRouted", logs: logs, sub: sub}, nil
}

// WatchTokensAndCallRouted is a free log subscription operation binding the contract event 0x42eff9005856e3c586b096d67211a566dc926052119fd7cc08023c70937ecb30.
//
// Solidity: event TokensAndCallRouted(bytes32 indexed teleporterMessageID, (bytes32,address,address,bytes,uint256,uint256,address,address,address,uint256,uint256) input, uint256 amount)
func (_NativeTokenHome *NativeTokenHomeFilterer) WatchTokensAndCallRouted(opts *bind.WatchOpts, sink chan<- *NativeTokenHomeTokensAndCallRouted, teleporterMessageID [][32]byte) (event.Subscription, error) {

	var teleporterMessageIDRule []interface{}
	for _, teleporterMessageIDItem := range teleporterMessageID {
		teleporterMessageIDRule = append(teleporterMessageIDRule, teleporterMessageIDItem)
	}

	logs, sub, err := _NativeTokenHome.contract.WatchLogs(opts, "TokensAndCallRouted", teleporterMessageIDRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NativeTokenHomeTokensAndCallRouted)
				if err := _NativeTokenHome.contract.UnpackLog(event, "TokensAndCallRouted", log); err != nil {
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

// ParseTokensAndCallRouted is a log parse operation binding the contract event 0x42eff9005856e3c586b096d67211a566dc926052119fd7cc08023c70937ecb30.
//
// Solidity: event TokensAndCallRouted(bytes32 indexed teleporterMessageID, (bytes32,address,address,bytes,uint256,uint256,address,address,address,uint256,uint256) input, uint256 amount)
func (_NativeTokenHome *NativeTokenHomeFilterer) ParseTokensAndCallRouted(log types.Log) (*NativeTokenHomeTokensAndCallRouted, error) {
	event := new(NativeTokenHomeTokensAndCallRouted)
	if err := _NativeTokenHome.contract.UnpackLog(event, "TokensAndCallRouted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// NativeTokenHomeTokensAndCallSentIterator is returned from FilterTokensAndCallSent and is used to iterate over the raw logs and unpacked data for TokensAndCallSent events raised by the NativeTokenHome contract.
type NativeTokenHomeTokensAndCallSentIterator struct {
	Event *NativeTokenHomeTokensAndCallSent // Event containing the contract specifics and raw log

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
func (it *NativeTokenHomeTokensAndCallSentIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NativeTokenHomeTokensAndCallSent)
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
		it.Event = new(NativeTokenHomeTokensAndCallSent)
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
func (it *NativeTokenHomeTokensAndCallSentIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NativeTokenHomeTokensAndCallSentIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NativeTokenHomeTokensAndCallSent represents a TokensAndCallSent event raised by the NativeTokenHome contract.
type NativeTokenHomeTokensAndCallSent struct {
	TeleporterMessageID [32]byte
	Sender              common.Address
	Input               SendAndCallInput
	Amount              *big.Int
	Raw                 types.Log // Blockchain specific contextual infos
}

// FilterTokensAndCallSent is a free log retrieval operation binding the contract event 0x5d76dff81bf773b908b050fa113d39f7d8135bb4175398f313ea19cd3a1a0b16.
//
// Solidity: event TokensAndCallSent(bytes32 indexed teleporterMessageID, address indexed sender, (bytes32,address,address,bytes,uint256,uint256,address,address,address,uint256,uint256) input, uint256 amount)
func (_NativeTokenHome *NativeTokenHomeFilterer) FilterTokensAndCallSent(opts *bind.FilterOpts, teleporterMessageID [][32]byte, sender []common.Address) (*NativeTokenHomeTokensAndCallSentIterator, error) {

	var teleporterMessageIDRule []interface{}
	for _, teleporterMessageIDItem := range teleporterMessageID {
		teleporterMessageIDRule = append(teleporterMessageIDRule, teleporterMessageIDItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _NativeTokenHome.contract.FilterLogs(opts, "TokensAndCallSent", teleporterMessageIDRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &NativeTokenHomeTokensAndCallSentIterator{contract: _NativeTokenHome.contract, event: "TokensAndCallSent", logs: logs, sub: sub}, nil
}

// WatchTokensAndCallSent is a free log subscription operation binding the contract event 0x5d76dff81bf773b908b050fa113d39f7d8135bb4175398f313ea19cd3a1a0b16.
//
// Solidity: event TokensAndCallSent(bytes32 indexed teleporterMessageID, address indexed sender, (bytes32,address,address,bytes,uint256,uint256,address,address,address,uint256,uint256) input, uint256 amount)
func (_NativeTokenHome *NativeTokenHomeFilterer) WatchTokensAndCallSent(opts *bind.WatchOpts, sink chan<- *NativeTokenHomeTokensAndCallSent, teleporterMessageID [][32]byte, sender []common.Address) (event.Subscription, error) {

	var teleporterMessageIDRule []interface{}
	for _, teleporterMessageIDItem := range teleporterMessageID {
		teleporterMessageIDRule = append(teleporterMessageIDRule, teleporterMessageIDItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _NativeTokenHome.contract.WatchLogs(opts, "TokensAndCallSent", teleporterMessageIDRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NativeTokenHomeTokensAndCallSent)
				if err := _NativeTokenHome.contract.UnpackLog(event, "TokensAndCallSent", log); err != nil {
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
func (_NativeTokenHome *NativeTokenHomeFilterer) ParseTokensAndCallSent(log types.Log) (*NativeTokenHomeTokensAndCallSent, error) {
	event := new(NativeTokenHomeTokensAndCallSent)
	if err := _NativeTokenHome.contract.UnpackLog(event, "TokensAndCallSent", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// NativeTokenHomeTokensRoutedIterator is returned from FilterTokensRouted and is used to iterate over the raw logs and unpacked data for TokensRouted events raised by the NativeTokenHome contract.
type NativeTokenHomeTokensRoutedIterator struct {
	Event *NativeTokenHomeTokensRouted // Event containing the contract specifics and raw log

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
func (it *NativeTokenHomeTokensRoutedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NativeTokenHomeTokensRouted)
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
		it.Event = new(NativeTokenHomeTokensRouted)
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
func (it *NativeTokenHomeTokensRoutedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NativeTokenHomeTokensRoutedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NativeTokenHomeTokensRouted represents a TokensRouted event raised by the NativeTokenHome contract.
type NativeTokenHomeTokensRouted struct {
	TeleporterMessageID [32]byte
	Input               SendTokensInput
	Amount              *big.Int
	Raw                 types.Log // Blockchain specific contextual infos
}

// FilterTokensRouted is a free log retrieval operation binding the contract event 0x825080857c76cef4a1629c0705a7f8b4ef0282ddcafde0b6715c4fb34b68aaf0.
//
// Solidity: event TokensRouted(bytes32 indexed teleporterMessageID, (bytes32,address,address,address,uint256,uint256,uint256,address) input, uint256 amount)
func (_NativeTokenHome *NativeTokenHomeFilterer) FilterTokensRouted(opts *bind.FilterOpts, teleporterMessageID [][32]byte) (*NativeTokenHomeTokensRoutedIterator, error) {

	var teleporterMessageIDRule []interface{}
	for _, teleporterMessageIDItem := range teleporterMessageID {
		teleporterMessageIDRule = append(teleporterMessageIDRule, teleporterMessageIDItem)
	}

	logs, sub, err := _NativeTokenHome.contract.FilterLogs(opts, "TokensRouted", teleporterMessageIDRule)
	if err != nil {
		return nil, err
	}
	return &NativeTokenHomeTokensRoutedIterator{contract: _NativeTokenHome.contract, event: "TokensRouted", logs: logs, sub: sub}, nil
}

// WatchTokensRouted is a free log subscription operation binding the contract event 0x825080857c76cef4a1629c0705a7f8b4ef0282ddcafde0b6715c4fb34b68aaf0.
//
// Solidity: event TokensRouted(bytes32 indexed teleporterMessageID, (bytes32,address,address,address,uint256,uint256,uint256,address) input, uint256 amount)
func (_NativeTokenHome *NativeTokenHomeFilterer) WatchTokensRouted(opts *bind.WatchOpts, sink chan<- *NativeTokenHomeTokensRouted, teleporterMessageID [][32]byte) (event.Subscription, error) {

	var teleporterMessageIDRule []interface{}
	for _, teleporterMessageIDItem := range teleporterMessageID {
		teleporterMessageIDRule = append(teleporterMessageIDRule, teleporterMessageIDItem)
	}

	logs, sub, err := _NativeTokenHome.contract.WatchLogs(opts, "TokensRouted", teleporterMessageIDRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NativeTokenHomeTokensRouted)
				if err := _NativeTokenHome.contract.UnpackLog(event, "TokensRouted", log); err != nil {
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

// ParseTokensRouted is a log parse operation binding the contract event 0x825080857c76cef4a1629c0705a7f8b4ef0282ddcafde0b6715c4fb34b68aaf0.
//
// Solidity: event TokensRouted(bytes32 indexed teleporterMessageID, (bytes32,address,address,address,uint256,uint256,uint256,address) input, uint256 amount)
func (_NativeTokenHome *NativeTokenHomeFilterer) ParseTokensRouted(log types.Log) (*NativeTokenHomeTokensRouted, error) {
	event := new(NativeTokenHomeTokensRouted)
	if err := _NativeTokenHome.contract.UnpackLog(event, "TokensRouted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// NativeTokenHomeTokensSentIterator is returned from FilterTokensSent and is used to iterate over the raw logs and unpacked data for TokensSent events raised by the NativeTokenHome contract.
type NativeTokenHomeTokensSentIterator struct {
	Event *NativeTokenHomeTokensSent // Event containing the contract specifics and raw log

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
func (it *NativeTokenHomeTokensSentIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NativeTokenHomeTokensSent)
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
		it.Event = new(NativeTokenHomeTokensSent)
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
func (it *NativeTokenHomeTokensSentIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NativeTokenHomeTokensSentIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NativeTokenHomeTokensSent represents a TokensSent event raised by the NativeTokenHome contract.
type NativeTokenHomeTokensSent struct {
	TeleporterMessageID [32]byte
	Sender              common.Address
	Input               SendTokensInput
	Amount              *big.Int
	Raw                 types.Log // Blockchain specific contextual infos
}

// FilterTokensSent is a free log retrieval operation binding the contract event 0x93f19bf1ec58a15dc643b37e7e18a1c13e85e06cd11929e283154691ace9fb52.
//
// Solidity: event TokensSent(bytes32 indexed teleporterMessageID, address indexed sender, (bytes32,address,address,address,uint256,uint256,uint256,address) input, uint256 amount)
func (_NativeTokenHome *NativeTokenHomeFilterer) FilterTokensSent(opts *bind.FilterOpts, teleporterMessageID [][32]byte, sender []common.Address) (*NativeTokenHomeTokensSentIterator, error) {

	var teleporterMessageIDRule []interface{}
	for _, teleporterMessageIDItem := range teleporterMessageID {
		teleporterMessageIDRule = append(teleporterMessageIDRule, teleporterMessageIDItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _NativeTokenHome.contract.FilterLogs(opts, "TokensSent", teleporterMessageIDRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &NativeTokenHomeTokensSentIterator{contract: _NativeTokenHome.contract, event: "TokensSent", logs: logs, sub: sub}, nil
}

// WatchTokensSent is a free log subscription operation binding the contract event 0x93f19bf1ec58a15dc643b37e7e18a1c13e85e06cd11929e283154691ace9fb52.
//
// Solidity: event TokensSent(bytes32 indexed teleporterMessageID, address indexed sender, (bytes32,address,address,address,uint256,uint256,uint256,address) input, uint256 amount)
func (_NativeTokenHome *NativeTokenHomeFilterer) WatchTokensSent(opts *bind.WatchOpts, sink chan<- *NativeTokenHomeTokensSent, teleporterMessageID [][32]byte, sender []common.Address) (event.Subscription, error) {

	var teleporterMessageIDRule []interface{}
	for _, teleporterMessageIDItem := range teleporterMessageID {
		teleporterMessageIDRule = append(teleporterMessageIDRule, teleporterMessageIDItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _NativeTokenHome.contract.WatchLogs(opts, "TokensSent", teleporterMessageIDRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NativeTokenHomeTokensSent)
				if err := _NativeTokenHome.contract.UnpackLog(event, "TokensSent", log); err != nil {
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
func (_NativeTokenHome *NativeTokenHomeFilterer) ParseTokensSent(log types.Log) (*NativeTokenHomeTokensSent, error) {
	event := new(NativeTokenHomeTokensSent)
	if err := _NativeTokenHome.contract.UnpackLog(event, "TokensSent", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// NativeTokenHomeTokensWithdrawnIterator is returned from FilterTokensWithdrawn and is used to iterate over the raw logs and unpacked data for TokensWithdrawn events raised by the NativeTokenHome contract.
type NativeTokenHomeTokensWithdrawnIterator struct {
	Event *NativeTokenHomeTokensWithdrawn // Event containing the contract specifics and raw log

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
func (it *NativeTokenHomeTokensWithdrawnIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NativeTokenHomeTokensWithdrawn)
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
		it.Event = new(NativeTokenHomeTokensWithdrawn)
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
func (it *NativeTokenHomeTokensWithdrawnIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NativeTokenHomeTokensWithdrawnIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NativeTokenHomeTokensWithdrawn represents a TokensWithdrawn event raised by the NativeTokenHome contract.
type NativeTokenHomeTokensWithdrawn struct {
	Recipient common.Address
	Amount    *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterTokensWithdrawn is a free log retrieval operation binding the contract event 0x6352c5382c4a4578e712449ca65e83cdb392d045dfcf1cad9615189db2da244b.
//
// Solidity: event TokensWithdrawn(address indexed recipient, uint256 amount)
func (_NativeTokenHome *NativeTokenHomeFilterer) FilterTokensWithdrawn(opts *bind.FilterOpts, recipient []common.Address) (*NativeTokenHomeTokensWithdrawnIterator, error) {

	var recipientRule []interface{}
	for _, recipientItem := range recipient {
		recipientRule = append(recipientRule, recipientItem)
	}

	logs, sub, err := _NativeTokenHome.contract.FilterLogs(opts, "TokensWithdrawn", recipientRule)
	if err != nil {
		return nil, err
	}
	return &NativeTokenHomeTokensWithdrawnIterator{contract: _NativeTokenHome.contract, event: "TokensWithdrawn", logs: logs, sub: sub}, nil
}

// WatchTokensWithdrawn is a free log subscription operation binding the contract event 0x6352c5382c4a4578e712449ca65e83cdb392d045dfcf1cad9615189db2da244b.
//
// Solidity: event TokensWithdrawn(address indexed recipient, uint256 amount)
func (_NativeTokenHome *NativeTokenHomeFilterer) WatchTokensWithdrawn(opts *bind.WatchOpts, sink chan<- *NativeTokenHomeTokensWithdrawn, recipient []common.Address) (event.Subscription, error) {

	var recipientRule []interface{}
	for _, recipientItem := range recipient {
		recipientRule = append(recipientRule, recipientItem)
	}

	logs, sub, err := _NativeTokenHome.contract.WatchLogs(opts, "TokensWithdrawn", recipientRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NativeTokenHomeTokensWithdrawn)
				if err := _NativeTokenHome.contract.UnpackLog(event, "TokensWithdrawn", log); err != nil {
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
func (_NativeTokenHome *NativeTokenHomeFilterer) ParseTokensWithdrawn(log types.Log) (*NativeTokenHomeTokensWithdrawn, error) {
	event := new(NativeTokenHomeTokensWithdrawn)
	if err := _NativeTokenHome.contract.UnpackLog(event, "TokensWithdrawn", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
