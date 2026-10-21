// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package remoteregistrar

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

// RemoteRegistrarMetaData contains all meta data concerning the RemoteRegistrar contract.
var RemoteRegistrarMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"contractITeleporterReceiver\",\"name\":\"home_\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"homeBlockchainID\",\"type\":\"bytes32\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"blockchainID\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"home\",\"outputs\":[{\"internalType\":\"contractITeleporterReceiver\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"remoteBlockchainID\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"remote\",\"type\":\"address\"},{\"internalType\":\"uint8\",\"name\":\"homeTokenDecimals\",\"type\":\"uint8\"},{\"internalType\":\"uint8\",\"name\":\"remoteTokenDecimals\",\"type\":\"uint8\"}],\"name\":\"register\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x60e060405234801561000f575f5ffd5b5060405161049438038061049483398101604081905261002e91610049565b336080526001600160a01b039190911660a05260c052610080565b5f5f6040838503121561005a575f5ffd5b82516001600160a01b0381168114610070575f5ffd5b6020939093015192949293505050565b60805160a05160c0516103df6100b55f395f60d301525f8181609701526101e001525f81816053015261010e01526103df5ff3fe608060405234801561000f575f5ffd5b506004361061004a575f3560e01c80638da5cb5b1461004e5780639fa92f9d14610092578063bdd740f1146100b9578063d127dc9b146100ce575b5f5ffd5b6100757f000000000000000000000000000000000000000000000000000000000000000081565b6040516001600160a01b0390911681526020015b60405180910390f35b6100757f000000000000000000000000000000000000000000000000000000000000000081565b6100cc6100c7366004610293565b610103565b005b6100f57f000000000000000000000000000000000000000000000000000000000000000081565b604051908152602001610089565b336001600160a01b037f0000000000000000000000000000000000000000000000000000000000000000161461017f5760405162461bcd60e51b815260206004820152601a60248201527f52656d6f74655265676973747261723a206e6f74206f776e6572000000000000604482015260640160405180910390fd5b60408051606080820183525f80835260ff8681166020808601918252878316958701958652865180820185905291518316828801529451909116818401528451808203909301835260c0810185526080810191825260a001829052925190927f00000000000000000000000000000000000000000000000000000000000000006001600160a01b03169263c868efaa928992899261021d920161032d565b6040516020818303038152906040526040518463ffffffff1660e01b815260040161024a93929190610377565b5f604051808303815f87803b158015610261575f5ffd5b505af1158015610273573d5f5f3e3d5ffd5b505050505050505050565b803560ff8116811461028e575f5ffd5b919050565b5f5f5f5f608085870312156102a6575f5ffd5b8435935060208501356001600160a01b03811681146102c3575f5ffd5b92506102d16040860161027e565b91506102df6060860161027e565b905092959194509250565b5f81518084525f5b8181101561030e576020818501810151868301820152016102f2565b505f602082860101526020601f19601f83011685010191505092915050565b602081525f82516005811061035057634e487b7160e01b5f52602160045260245ffd5b80602084015250602083015160408084015261036f60608401826102ea565b949350505050565b8381526001600160a01b03831660208201526060604082018190525f906103a0908301846102ea565b9594505050505056fea264697066735822122060d5609de65ff7903f7adff8b799827b1296cfb30e1adeb745b727778727254d64736f6c634300081e0033",
}

// RemoteRegistrarABI is the input ABI used to generate the binding from.
// Deprecated: Use RemoteRegistrarMetaData.ABI instead.
var RemoteRegistrarABI = RemoteRegistrarMetaData.ABI

// RemoteRegistrarBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use RemoteRegistrarMetaData.Bin instead.
var RemoteRegistrarBin = RemoteRegistrarMetaData.Bin

// DeployRemoteRegistrar deploys a new Ethereum contract, binding an instance of RemoteRegistrar to it.
func DeployRemoteRegistrar(auth *bind.TransactOpts, backend bind.ContractBackend, home_ common.Address, homeBlockchainID [32]byte) (common.Address, *types.Transaction, *RemoteRegistrar, error) {
	parsed, err := RemoteRegistrarMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(RemoteRegistrarBin), backend, home_, homeBlockchainID)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &RemoteRegistrar{RemoteRegistrarCaller: RemoteRegistrarCaller{contract: contract}, RemoteRegistrarTransactor: RemoteRegistrarTransactor{contract: contract}, RemoteRegistrarFilterer: RemoteRegistrarFilterer{contract: contract}}, nil
}

// RemoteRegistrar is an auto generated Go binding around an Ethereum contract.
type RemoteRegistrar struct {
	RemoteRegistrarCaller     // Read-only binding to the contract
	RemoteRegistrarTransactor // Write-only binding to the contract
	RemoteRegistrarFilterer   // Log filterer for contract events
}

// RemoteRegistrarCaller is an auto generated read-only Go binding around an Ethereum contract.
type RemoteRegistrarCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RemoteRegistrarTransactor is an auto generated write-only Go binding around an Ethereum contract.
type RemoteRegistrarTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RemoteRegistrarFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type RemoteRegistrarFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RemoteRegistrarSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type RemoteRegistrarSession struct {
	Contract     *RemoteRegistrar  // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// RemoteRegistrarCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type RemoteRegistrarCallerSession struct {
	Contract *RemoteRegistrarCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts          // Call options to use throughout this session
}

// RemoteRegistrarTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type RemoteRegistrarTransactorSession struct {
	Contract     *RemoteRegistrarTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts          // Transaction auth options to use throughout this session
}

// RemoteRegistrarRaw is an auto generated low-level Go binding around an Ethereum contract.
type RemoteRegistrarRaw struct {
	Contract *RemoteRegistrar // Generic contract binding to access the raw methods on
}

// RemoteRegistrarCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type RemoteRegistrarCallerRaw struct {
	Contract *RemoteRegistrarCaller // Generic read-only contract binding to access the raw methods on
}

// RemoteRegistrarTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type RemoteRegistrarTransactorRaw struct {
	Contract *RemoteRegistrarTransactor // Generic write-only contract binding to access the raw methods on
}

// NewRemoteRegistrar creates a new instance of RemoteRegistrar, bound to a specific deployed contract.
func NewRemoteRegistrar(address common.Address, backend bind.ContractBackend) (*RemoteRegistrar, error) {
	contract, err := bindRemoteRegistrar(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &RemoteRegistrar{RemoteRegistrarCaller: RemoteRegistrarCaller{contract: contract}, RemoteRegistrarTransactor: RemoteRegistrarTransactor{contract: contract}, RemoteRegistrarFilterer: RemoteRegistrarFilterer{contract: contract}}, nil
}

// NewRemoteRegistrarCaller creates a new read-only instance of RemoteRegistrar, bound to a specific deployed contract.
func NewRemoteRegistrarCaller(address common.Address, caller bind.ContractCaller) (*RemoteRegistrarCaller, error) {
	contract, err := bindRemoteRegistrar(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &RemoteRegistrarCaller{contract: contract}, nil
}

// NewRemoteRegistrarTransactor creates a new write-only instance of RemoteRegistrar, bound to a specific deployed contract.
func NewRemoteRegistrarTransactor(address common.Address, transactor bind.ContractTransactor) (*RemoteRegistrarTransactor, error) {
	contract, err := bindRemoteRegistrar(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &RemoteRegistrarTransactor{contract: contract}, nil
}

// NewRemoteRegistrarFilterer creates a new log filterer instance of RemoteRegistrar, bound to a specific deployed contract.
func NewRemoteRegistrarFilterer(address common.Address, filterer bind.ContractFilterer) (*RemoteRegistrarFilterer, error) {
	contract, err := bindRemoteRegistrar(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &RemoteRegistrarFilterer{contract: contract}, nil
}

// bindRemoteRegistrar binds a generic wrapper to an already deployed contract.
func bindRemoteRegistrar(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := RemoteRegistrarMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_RemoteRegistrar *RemoteRegistrarRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _RemoteRegistrar.Contract.RemoteRegistrarCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_RemoteRegistrar *RemoteRegistrarRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _RemoteRegistrar.Contract.RemoteRegistrarTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_RemoteRegistrar *RemoteRegistrarRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _RemoteRegistrar.Contract.RemoteRegistrarTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_RemoteRegistrar *RemoteRegistrarCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _RemoteRegistrar.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_RemoteRegistrar *RemoteRegistrarTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _RemoteRegistrar.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_RemoteRegistrar *RemoteRegistrarTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _RemoteRegistrar.Contract.contract.Transact(opts, method, params...)
}

// BlockchainID is a free data retrieval call binding the contract method 0xd127dc9b.
//
// Solidity: function blockchainID() view returns(bytes32)
func (_RemoteRegistrar *RemoteRegistrarCaller) BlockchainID(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _RemoteRegistrar.contract.Call(opts, &out, "blockchainID")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// BlockchainID is a free data retrieval call binding the contract method 0xd127dc9b.
//
// Solidity: function blockchainID() view returns(bytes32)
func (_RemoteRegistrar *RemoteRegistrarSession) BlockchainID() ([32]byte, error) {
	return _RemoteRegistrar.Contract.BlockchainID(&_RemoteRegistrar.CallOpts)
}

// BlockchainID is a free data retrieval call binding the contract method 0xd127dc9b.
//
// Solidity: function blockchainID() view returns(bytes32)
func (_RemoteRegistrar *RemoteRegistrarCallerSession) BlockchainID() ([32]byte, error) {
	return _RemoteRegistrar.Contract.BlockchainID(&_RemoteRegistrar.CallOpts)
}

// Home is a free data retrieval call binding the contract method 0x9fa92f9d.
//
// Solidity: function home() view returns(address)
func (_RemoteRegistrar *RemoteRegistrarCaller) Home(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _RemoteRegistrar.contract.Call(opts, &out, "home")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Home is a free data retrieval call binding the contract method 0x9fa92f9d.
//
// Solidity: function home() view returns(address)
func (_RemoteRegistrar *RemoteRegistrarSession) Home() (common.Address, error) {
	return _RemoteRegistrar.Contract.Home(&_RemoteRegistrar.CallOpts)
}

// Home is a free data retrieval call binding the contract method 0x9fa92f9d.
//
// Solidity: function home() view returns(address)
func (_RemoteRegistrar *RemoteRegistrarCallerSession) Home() (common.Address, error) {
	return _RemoteRegistrar.Contract.Home(&_RemoteRegistrar.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_RemoteRegistrar *RemoteRegistrarCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _RemoteRegistrar.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_RemoteRegistrar *RemoteRegistrarSession) Owner() (common.Address, error) {
	return _RemoteRegistrar.Contract.Owner(&_RemoteRegistrar.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_RemoteRegistrar *RemoteRegistrarCallerSession) Owner() (common.Address, error) {
	return _RemoteRegistrar.Contract.Owner(&_RemoteRegistrar.CallOpts)
}

// Register is a paid mutator transaction binding the contract method 0xbdd740f1.
//
// Solidity: function register(bytes32 remoteBlockchainID, address remote, uint8 homeTokenDecimals, uint8 remoteTokenDecimals) returns()
func (_RemoteRegistrar *RemoteRegistrarTransactor) Register(opts *bind.TransactOpts, remoteBlockchainID [32]byte, remote common.Address, homeTokenDecimals uint8, remoteTokenDecimals uint8) (*types.Transaction, error) {
	return _RemoteRegistrar.contract.Transact(opts, "register", remoteBlockchainID, remote, homeTokenDecimals, remoteTokenDecimals)
}

// Register is a paid mutator transaction binding the contract method 0xbdd740f1.
//
// Solidity: function register(bytes32 remoteBlockchainID, address remote, uint8 homeTokenDecimals, uint8 remoteTokenDecimals) returns()
func (_RemoteRegistrar *RemoteRegistrarSession) Register(remoteBlockchainID [32]byte, remote common.Address, homeTokenDecimals uint8, remoteTokenDecimals uint8) (*types.Transaction, error) {
	return _RemoteRegistrar.Contract.Register(&_RemoteRegistrar.TransactOpts, remoteBlockchainID, remote, homeTokenDecimals, remoteTokenDecimals)
}

// Register is a paid mutator transaction binding the contract method 0xbdd740f1.
//
// Solidity: function register(bytes32 remoteBlockchainID, address remote, uint8 homeTokenDecimals, uint8 remoteTokenDecimals) returns()
func (_RemoteRegistrar *RemoteRegistrarTransactorSession) Register(remoteBlockchainID [32]byte, remote common.Address, homeTokenDecimals uint8, remoteTokenDecimals uint8) (*types.Transaction, error) {
	return _RemoteRegistrar.Contract.Register(&_RemoteRegistrar.TransactOpts, remoteBlockchainID, remote, homeTokenDecimals, remoteTokenDecimals)
}
