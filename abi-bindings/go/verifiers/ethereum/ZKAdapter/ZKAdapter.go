// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package zkadapter

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

// ConsensusCheckpoint is an auto generated low-level Go binding around an user-defined struct.
type ConsensusCheckpoint struct {
	Epoch uint64
	Root  [32]byte
}

// ConsensusData is an auto generated low-level Go binding around an user-defined struct.
type ConsensusData struct {
	JournalData   []byte
	Seal          []byte
	FinalizedSlot uint64
}

// ConsensusState is an auto generated low-level Go binding around an user-defined struct.
type ConsensusState struct {
	CurrentJustifiedCheckpoint ConsensusCheckpoint
	FinalizedCheckpoint        ConsensusCheckpoint
}

// ExecutionBeaconConfig is an auto generated low-level Go binding around an user-defined struct.
type ExecutionBeaconConfig struct {
	GIndexBlockStateRoot      *big.Int
	GIndexExecRoot            *big.Int
	GIndexBaseStateRoots      *big.Int
	StateRootsDepth           *big.Int
	GIndexReceiptsRoot        *big.Int
	StateRootsVectorSize      *big.Int
	GloasForkSlot             *big.Int
	GIndexBaseStateRootsGloas *big.Int
	GIndexLatestBlockHash     *big.Int
}

// ExecutionProof is an auto generated low-level Go binding around an user-defined struct.
type ExecutionProof struct {
	AnchorSlot                 uint64
	TargetSlot                 uint64
	AnchorBeaconStateRoot      [32]byte
	AnchorBeaconStateProof     [][32]byte
	TargetBeaconStateRoot      [32]byte
	TargetBeaconStateProof     [][32]byte
	TargetExecutionHeaderRoot  [32]byte
	TargetExecutionHeaderProof [][32]byte
	TargetReceiptsRoot         [32]byte
	TargetReceiptsProof        [][32]byte
	ExecutionHeaderRlp         []byte
}

// ReceiptProof is an auto generated low-level Go binding around an user-defined struct.
type ReceiptProof struct {
	Proof           [][]byte
	Key             []byte
	Value           []byte
	LogIndex        *big.Int
	ExpectedEmitter common.Address
	ExpectedTopic0  [32]byte
}

// TeleporterICMMessage is an auto generated low-level Go binding around an user-defined struct.
type TeleporterICMMessage struct {
	Message            TeleporterMessageV2
	SourceNetworkID    uint32
	SourceBlockchainID [32]byte
	Attestation        []byte
}

// TeleporterMessageReceipt is an auto generated low-level Go binding around an user-defined struct.
type TeleporterMessageReceipt struct {
	ReceivedMessageNonce *big.Int
	RelayerRewardAddress common.Address
}

// TeleporterMessageV2 is an auto generated low-level Go binding around an user-defined struct.
type TeleporterMessageV2 struct {
	MessageNonce            *big.Int
	OriginSenderAddress     common.Address
	OriginTeleporterAddress common.Address
	DestinationBlockchainID [32]byte
	DestinationAddress      common.Address
	RequiredGasLimit        *big.Int
	AllowedRelayerAddresses []common.Address
	Receipts                []TeleporterMessageReceipt
	Message                 []byte
}

// ZKAdapterMetaData contains all meta data concerning the ZKAdapter contract.
var ZKAdapterMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"sourceChainId_\",\"type\":\"uint256\"},{\"components\":[{\"components\":[{\"internalType\":\"uint64\",\"name\":\"epoch\",\"type\":\"uint64\"},{\"internalType\":\"bytes32\",\"name\":\"root\",\"type\":\"bytes32\"}],\"internalType\":\"structConsensus.Checkpoint\",\"name\":\"currentJustifiedCheckpoint\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint64\",\"name\":\"epoch\",\"type\":\"uint64\"},{\"internalType\":\"bytes32\",\"name\":\"root\",\"type\":\"bytes32\"}],\"internalType\":\"structConsensus.Checkpoint\",\"name\":\"finalizedCheckpoint\",\"type\":\"tuple\"}],\"internalType\":\"structConsensus.State\",\"name\":\"startingState\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"gIndexBlockStateRoot\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"gIndexExecRoot\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"gIndexBaseStateRoots\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"stateRootsDepth\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"gIndexReceiptsRoot\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"stateRootsVectorSize\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"gloasForkSlot\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"gIndexBaseStateRootsGloas\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"gIndexLatestBlockHash\",\"type\":\"uint256\"}],\"internalType\":\"structExecution.BeaconConfig\",\"name\":\"beaconConfig_\",\"type\":\"tuple\"},{\"internalType\":\"uint256\",\"name\":\"genesisTime_\",\"type\":\"uint256\"},{\"internalType\":\"uint24\",\"name\":\"permissibleTimespan_\",\"type\":\"uint24\"},{\"internalType\":\"address\",\"name\":\"verifier\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"imageID_\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"admin\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"superAdmin\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"AccessControlBadConfirmation\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"neededRole\",\"type\":\"bytes32\"}],\"name\":\"AccessControlUnauthorizedAccount\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidArgument\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidPreState\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"PermissibleTimespanLapsed\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"chainId\",\"type\":\"uint256\"}],\"name\":\"UnsupportedSourceChain\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint64\",\"name\":\"slot\",\"type\":\"uint64\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"root\",\"type\":\"bytes32\"}],\"name\":\"ConfirmedBeaconBlock\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"newImageID\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"oldImageID\",\"type\":\"bytes32\"}],\"name\":\"ImageIDUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint24\",\"name\":\"permissibleTimespan\",\"type\":\"uint24\"}],\"name\":\"PermissibleTimespanUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"previousAdminRole\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"newAdminRole\",\"type\":\"bytes32\"}],\"name\":\"RoleAdminChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleGranted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleRevoked\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"encodedMessage\",\"type\":\"bytes\"}],\"name\":\"TeleporterV2MessageSent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint64\",\"name\":\"preEpoch\",\"type\":\"uint64\"},{\"indexed\":true,\"internalType\":\"uint64\",\"name\":\"postEpoch\",\"type\":\"uint64\"},{\"components\":[{\"components\":[{\"internalType\":\"uint64\",\"name\":\"epoch\",\"type\":\"uint64\"},{\"internalType\":\"bytes32\",\"name\":\"root\",\"type\":\"bytes32\"}],\"internalType\":\"structConsensus.Checkpoint\",\"name\":\"currentJustifiedCheckpoint\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint64\",\"name\":\"epoch\",\"type\":\"uint64\"},{\"internalType\":\"bytes32\",\"name\":\"root\",\"type\":\"bytes32\"}],\"internalType\":\"structConsensus.Checkpoint\",\"name\":\"finalizedCheckpoint\",\"type\":\"tuple\"}],\"indexed\":false,\"internalType\":\"structConsensus.State\",\"name\":\"preState\",\"type\":\"tuple\"},{\"components\":[{\"components\":[{\"internalType\":\"uint64\",\"name\":\"epoch\",\"type\":\"uint64\"},{\"internalType\":\"bytes32\",\"name\":\"root\",\"type\":\"bytes32\"}],\"internalType\":\"structConsensus.Checkpoint\",\"name\":\"currentJustifiedCheckpoint\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint64\",\"name\":\"epoch\",\"type\":\"uint64\"},{\"internalType\":\"bytes32\",\"name\":\"root\",\"type\":\"bytes32\"}],\"internalType\":\"structConsensus.Checkpoint\",\"name\":\"finalizedCheckpoint\",\"type\":\"tuple\"}],\"indexed\":false,\"internalType\":\"structConsensus.State\",\"name\":\"postState\",\"type\":\"tuple\"}],\"name\":\"Transitioned\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"sourceChainId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"beaconSlot\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"emitter\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"executionRoot\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"logIndex\",\"type\":\"uint256\"}],\"name\":\"ZKEventImported\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"ADMIN_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"DEFAULT_ADMIN_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"SECONDS_PER_SLOT\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"SLOT_PER_EPOCH\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"UNDEFINED_ROOT\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"slot\",\"type\":\"uint64\"}],\"name\":\"beaconRootAt\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"epoch\",\"type\":\"uint64\"}],\"name\":\"epochTimestamp\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"genesisTime\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"slot\",\"type\":\"uint64\"}],\"name\":\"getBeaconBlockRoot\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"root\",\"type\":\"bytes32\"},{\"internalType\":\"bool\",\"name\":\"valid\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"}],\"name\":\"getRoleAdmin\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"grantRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"hasRole\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"imageID\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"journalData\",\"type\":\"bytes\"},{\"internalType\":\"uint64\",\"name\":\"finalizedSlot\",\"type\":\"uint64\"}],\"name\":\"manualTransition\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"permissibleTimespan\",\"outputs\":[{\"internalType\":\"uint24\",\"name\":\"\",\"type\":\"uint24\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"callerConfirmation\",\"type\":\"address\"}],\"name\":\"renounceRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"revokeRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"messageNonce\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"originSenderAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"originTeleporterAddress\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"destinationBlockchainID\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"destinationAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"requiredGasLimit\",\"type\":\"uint256\"},{\"internalType\":\"address[]\",\"name\":\"allowedRelayerAddresses\",\"type\":\"address[]\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"receivedMessageNonce\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"relayerRewardAddress\",\"type\":\"address\"}],\"internalType\":\"structTeleporterMessageReceipt[]\",\"name\":\"receipts\",\"type\":\"tuple[]\"},{\"internalType\":\"bytes\",\"name\":\"message\",\"type\":\"bytes\"}],\"internalType\":\"structTeleporterMessageV2\",\"name\":\"message\",\"type\":\"tuple\"}],\"name\":\"sendMessage\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"sourceChainId\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"sourceNetworkId\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"bytes\",\"name\":\"journalData\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"seal\",\"type\":\"bytes\"},{\"internalType\":\"uint64\",\"name\":\"finalizedSlot\",\"type\":\"uint64\"}],\"internalType\":\"structConsensusData\",\"name\":\"consensus\",\"type\":\"tuple\"}],\"name\":\"transition\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"newImageID\",\"type\":\"bytes32\"}],\"name\":\"updateImageID\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint24\",\"name\":\"newPermissibleTimespan\",\"type\":\"uint24\"}],\"name\":\"updatePermissibleTimespan\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newVerifier\",\"type\":\"address\"}],\"name\":\"updateVerifier\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"verifier\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint64\",\"name\":\"anchorSlot\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"targetSlot\",\"type\":\"uint64\"},{\"internalType\":\"bytes32\",\"name\":\"anchorBeaconStateRoot\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32[]\",\"name\":\"anchorBeaconStateProof\",\"type\":\"bytes32[]\"},{\"internalType\":\"bytes32\",\"name\":\"targetBeaconStateRoot\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32[]\",\"name\":\"targetBeaconStateProof\",\"type\":\"bytes32[]\"},{\"internalType\":\"bytes32\",\"name\":\"targetExecutionHeaderRoot\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32[]\",\"name\":\"targetExecutionHeaderProof\",\"type\":\"bytes32[]\"},{\"internalType\":\"bytes32\",\"name\":\"targetReceiptsRoot\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32[]\",\"name\":\"targetReceiptsProof\",\"type\":\"bytes32[]\"},{\"internalType\":\"bytes\",\"name\":\"executionHeaderRlp\",\"type\":\"bytes\"}],\"internalType\":\"structExecution.Proof\",\"name\":\"execProof\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"bytes[]\",\"name\":\"proof\",\"type\":\"bytes[]\"},{\"internalType\":\"bytes\",\"name\":\"key\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"value\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"logIndex\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"expectedEmitter\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"expectedTopic0\",\"type\":\"bytes32\"}],\"internalType\":\"structReceipt.Proof\",\"name\":\"logProof\",\"type\":\"tuple\"}],\"name\":\"verifyLogAndExtract\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"messageNonce\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"originSenderAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"originTeleporterAddress\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"destinationBlockchainID\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"destinationAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"requiredGasLimit\",\"type\":\"uint256\"},{\"internalType\":\"address[]\",\"name\":\"allowedRelayerAddresses\",\"type\":\"address[]\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"receivedMessageNonce\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"relayerRewardAddress\",\"type\":\"address\"}],\"internalType\":\"structTeleporterMessageReceipt[]\",\"name\":\"receipts\",\"type\":\"tuple[]\"},{\"internalType\":\"bytes\",\"name\":\"message\",\"type\":\"bytes\"}],\"internalType\":\"structTeleporterMessageV2\",\"name\":\"message\",\"type\":\"tuple\"},{\"internalType\":\"uint32\",\"name\":\"sourceNetworkID\",\"type\":\"uint32\"},{\"internalType\":\"bytes32\",\"name\":\"sourceBlockchainID\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"attestation\",\"type\":\"bytes\"}],\"internalType\":\"structTeleporterICMMessage\",\"name\":\"message\",\"type\":\"tuple\"}],\"name\":\"verifyMessage\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x60e060405234801561000f575f5ffd5b5060405161539638038061539683398101604081905261002e916104c7565b8888888888888888888887815f036100805760405162461bcd60e51b815260206004820152601060248201526f125b9d985b1a590818da185a5b88125160821b60448201526064015b60405180910390fd5b608091825280516001556020810151600255604081015160035560608101516004559081015160055560a081015160065560c081015160075560e081015160085561010001516009555f8690036101195760405162461bcd60e51b815260206004820152601460248201527f496e76616c69642067656e657369732074696d650000000000000000000000006044820152606401610077565b60a08690526101487fa49807205ce4d355092ef5a8a18f56e8913cf4a201fbe287825b095693c2177583610265565b506101535f82610265565b5060208801515142906101659061030e565b11156101b35760405162461bcd60e51b815260206004820152601c60248201527f5374617274696e6720737461746520696e2074686520667574757265000000006044820152606401610077565b505085518051600a80546001600160401b03199081166001600160401b0393841617909155602092830151600b55978201518051600c8054909a169216919091179097559590950151600d55601080546001600160b81b031916600160a01b62ffffff94909416939093026001600160a01b031916929092176001600160a01b03919091161790555050600f555060805161024d9061033c565b63ffffffff1660c052506105ca975050505050505050565b5f828152602081815260408083206001600160a01b038516845290915281205460ff16610305575f838152602081815260408083206001600160a01b03861684529091529020805460ff191660011790556102bd3390565b6001600160a01b0316826001600160a01b0316847f2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d60405160405180910390a4506001610308565b505f5b92915050565b5f600c61032560206001600160401b0385166105a0565b61032f91906105a0565b60a05161030891906105b7565b5f8160010361034d57506001919050565b8162aa36a70361035f57506002919050565b6040516372c559db60e01b815260048101839052602401610077565b604080519081016001600160401b03811182821017156103a957634e487b7160e01b5f52604160045260245ffd5b60405290565b60405161012081016001600160401b03811182821017156103a957634e487b7160e01b5f52604160045260245ffd5b5f604082840312156103ee575f5ffd5b6103f661037b565b82519091506001600160401b038116811461040f575f5ffd5b815260209182015191810191909152919050565b5f6101208284031215610434575f5ffd5b61043c6103af565b825181526020808401519082015260408084015190820152606080840151908201526080808401519082015260a0808401519082015260c0808401519082015260e08084015190820152610100928301519281019290925250919050565b805162ffffff811681146104ac575f5ffd5b919050565b80516001600160a01b03811681146104ac575f5ffd5b5f5f5f5f5f5f5f5f5f898b036102808112156104e1575f5ffd5b8a5199506080601f19820112156104f6575f5ffd5b506104ff61037b565b61050c8c60208d016103de565b815261051b8c60608d016103de565b6020820152975061052f8b60a08c01610423565b6101c08b015190975095506105476101e08b0161049a565b94506105566102008b016104b1565b6102208b0151909450925061056e6102408b016104b1565b915061057d6102608b016104b1565b90509295985092959850929598565b634e487b7160e01b5f52601160045260245ffd5b80820281158282048414176103085761030861058c565b808201808211156103085761030861058c565b60805160a05160c051614d866106105f395f81816103450152610ae101525f81816102e0015261055b01525f8181610216015281816107c20152610a7b0152614d865ff3fe608060405234801561000f575f5ffd5b50600436106101bb575f3560e01c80636fda4c8b116100f3578063c506e0be11610093578063db2635b71161006e578063db2635b714610450578063e97eb36214610463578063eb97cd2c14610476578063f1faff0014610489575f5ffd5b8063c506e0be146103e9578063d0850708146103fc578063d547741f1461043d575f5ffd5b806391d14854116100ce57806391d14854146103a357806397fc007c146103b657806398ad485d146103c9578063a217fddf1461031e575f5ffd5b80636fda4c8b1461034057806375b238fc1461037c5780638e1e49ba14610390575f5ffd5b806336568abe1161015e5780634d3499b5116101395780634d3499b51461031557806354747fd41461031e578063547923a31461032557806365242a0e14610338575f5ffd5b806336568abe146102c857806342c6498a146102db5780634b0d0fa814610302575f5ffd5b8063248a9ca311610199578063248a9ca3146102465780632b7ac3f3146102685780632f2ff15d14610293578063304b9071146102a8575f5ffd5b806301ffc9a7146101bf57806304ae4327146101e75780631544298e14610211575b5f5ffd5b6101d26101cd366004613bf0565b61049c565b60405190151581526020015b60405180910390f35b6010546101fd90600160a01b900462ffffff1681565b60405162ffffff90911681526020016101de565b6102387f000000000000000000000000000000000000000000000000000000000000000081565b6040519081526020016101de565b610238610254366004613c17565b5f9081526020819052604090206001015490565b60105461027b906001600160a01b031681565b6040516001600160a01b0390911681526020016101de565b6102a66102a1366004613c52565b6104d2565b005b6102b0600c81565b6040516001600160401b0390911681526020016101de565b6102a66102d6366004613c52565b6104fc565b6102387f000000000000000000000000000000000000000000000000000000000000000081565b610238610310366004613c96565b610534565b610238600f5481565b6102385f81565b6102a6610333366004613caf565b61057f565b6102b0602081565b6103677f000000000000000000000000000000000000000000000000000000000000000081565b60405163ffffffff90911681526020016101de565b6102385f516020614d315f395f51905f5281565b61023861039e366004613c96565b6105ed565b6101d26103b1366004613c52565b61060a565b6102a66103c4366004613cd1565b610632565b6103dc6103d7366004613cec565b6106f6565b6040516101de9190613d9c565b6102a66103f7366004613dae565b610834565b61042861040a366004613c96565b6001600160401b03165f908152600e60205260409020549081151590565b604080519283529015156020830152016101de565b6102a661044b366004613c52565b61086b565b6102a661045e366004613e28565b61088f565b6102a6610471366004613c17565b6108d0565b6102a6610484366004613e5e565b61090f565b6101d2610497366004613e95565b610a56565b5f6001600160e01b03198216637965db0b60e01b14806104cc57506301ffc9a760e01b6001600160e01b03198316145b92915050565b5f828152602081905260409020600101546104ec81610d0d565b6104f68383610d1a565b50505050565b6001600160a01b03811633146105255760405163334bd91960e11b815260040160405180910390fd5b61052f8282610da9565b505050565b5f600c61054b60206001600160401b038516613edf565b6105559190613edf565b6104cc907f0000000000000000000000000000000000000000000000000000000000000000613ef6565b5f516020614d315f395f51905f5261059681610d0d565b60105462ffffff600160a01b9091048116908316036105c85760405163a9cb9e0d60e01b815260040160405180910390fd5b506010805462ffffff909216600160a01b0262ffffff60a01b19909216919091179055565b6001600160401b0381165f908152600e60205260408120546104cc565b5f918252602082815260408084206001600160a01b0393909316845291905290205460ff1690565b5f516020614d315f395f51905f5261064981610d0d565b6001600160a01b0382166106a45760405162461bcd60e51b815260206004820152601860248201527f496e76616c69642076657269666965722061646472657373000000000000000060448201526064015b60405180910390fd5b6010546001600160a01b03908116908316036106d35760405163a9cb9e0d60e01b815260040160405180910390fd5b50601080546001600160a01b0319166001600160a01b0392909216919091179055565b60605f6107236107096020860186613c96565b6001600160401b03165f908152600e602052604090205490565b9050806107725760405162461bcd60e51b815260206004820152601860248201527f416e63686f7220736c6f7420697320756e646566696e65640000000000000000604482015260640161069b565b5f61077f82866001610e12565b90505f61078c82866112c4565b905061079e60a0860160808701613cd1565b6001600160a01b03166107b76040880160208901613c96565b6001600160401b03167f00000000000000000000000000000000000000000000000000000000000000007f17e2291123b98af9b416b45d04d46e4484bbf9d2c1061bbf567425a66bb1beb18960c001358960600135604051610823929190918252602082015260400190565b60405180910390a495945050505050565b5f516020614d315f395f51905f5261084b81610d0d565b5f6108588486018661402f565b90506108648184611611565b5050505050565b5f8281526020819052604090206001015461088581610d0d565b6104f68383610da9565b5f61089a828061409e565b8101906108a7919061402f565b90506108b381836116b8565b6108cc816108c76060850160408601613c96565b611611565b5050565b5f516020614d315f395f51905f526108e781610d0d565b600f5482036109095760405163a9cb9e0d60e01b815260040160405180910390fd5b50600f55565b61091f6060820160408301613cd1565b6001600160a01b0316336001600160a01b031614806109bd57506109496060820160408301613cd1565b6001600160a01b031663d67bdd256040518163ffffffff1660e01b8152600401602060405180830381865afa158015610984573d5f5f3e3d5ffd5b505050506040513d601f19601f820116820180604052508101906109a891906140e7565b6001600160a01b0316336001600160a01b0316145b6109ff5760405162461bcd60e51b81526020600482015260136024820152723ab730baba3437b934bd32b21039b2b73232b960691b604482015260640161069b565b7f467f98200caa1453d527fd7de92faa829d8357fb650c10010e77188d93a3a3b681604051602001610a319190614283565b60408051601f1981840301815290829052610a4b91613d9c565b60405180910390a150565b5f80610a65606084018461409e565b810190610a7291906145bc565b905060408301357f000000000000000000000000000000000000000000000000000000000000000014610ada5760405162461bcd60e51b815260206004820152601060248201526f3130b21039b7bab931b29031b430b4b760811b604482015260640161069b565b63ffffffff7f000000000000000000000000000000000000000000000000000000000000000016610b116040850160208601614779565b63ffffffff1614610b525760405162461bcd60e51b815260206004820152600b60248201526a626164206e6574776f726b60a81b604482015260640161069b565b6020810151608001516001600160a01b03163014610ba05760405162461bcd60e51b815260206004820152600b60248201526a3130b21032b6b4ba3a32b960a91b604482015260640161069b565b7f467f98200caa1453d527fd7de92faa829d8357fb650c10010e77188d93a3a3b6816020015160a0015114610c035760405162461bcd60e51b815260206004820152600960248201526862616420746f70696360b81b604482015260640161069b565b805160208201516040516398ad485d60e01b81525f9230926398ad485d92610c2f929190600401614898565b5f604051808303815f875af1158015610c4a573d5f5f3e3d5ffd5b505050506040513d5f823e601f3d908101601f19168201604052610c7191908101906149a5565b90505f81806020019051810190610c8891906149a5565b9050610c948580614a0d565b604051602001610ca49190614283565b60405160208183030381529060405280519060200120818051906020012014610d025760405162461bcd60e51b815260206004820152601060248201526f0e0c2f2d8dec2c840dad2e6dac2e8c6d60831b604482015260640161069b565b506001949350505050565b610d178133611825565b50565b5f610d25838361060a565b610da2575f838152602081815260408083206001600160a01b03861684529091529020805460ff19166001179055610d5a3390565b6001600160a01b0316826001600160a01b0316847f2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d60405160405180910390a45060016104cc565b505f6104cc565b5f610db4838361060a565b15610da2575f838152602081815260408083206001600160a01b0386168085529252808320805460ff1916905551339286917ff6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b9190a45060016104cc565b5f610e206020840184613c96565b6001600160401b0316610e396040850160208601613c96565b6001600160401b031610610ea15760405162461bcd60e51b815260206004820152602960248201527f54617267657420736c6f74206d757374206265206c657373207468616e20616e60448201526818da1bdc881cdb1bdd60ba1b606482015260840161069b565b6005820154610eb66040850160208601613c96565b610ec36020860186613c96565b610ecd9190614a2c565b6001600160401b03161115610f1d5760405162461bcd60e51b815260206004820152601660248201527515185c99d95d081cdb1bdd081a5cc81d1bdbc81bdb1960521b604482015260640161069b565b81545f90610f7090604086013590610f386060880188614a4b565b808060200260200160405190810160405280939291908181526020018383602002808284375f920191909152508b925061185e915050565b905080610fbf5760405162461bcd60e51b815260206004820152601960248201527f496e76616c696420616e63686f7220737461746520726f6f7400000000000000604482015260640161069b565b60058301545f90610fd66040870160208801613c96565b6001600160401b0316610fe99190614aa4565b60068501549091505f906110006020880188613c96565b6001600160401b0316101561101957846002015461101f565b84600701545b90505f82866003015483901b6110359190613ef6565b90505f61108860808901358361104e60a08c018c614a4b565b808060200260200160405190810160405280939291908181526020018383602002808284375f9201919091525050505060408c013561185e565b9050806110d75760405162461bcd60e51b815260206004820181905260248201527f496e76616c69642074617267657420626561636f6e20737461746520726f6f74604482015260640161069b565b60068701546110ec60408a0160208b01613c96565b6001600160401b031610611110576111048888611901565b955050505050506112bd565b61111e61014089018961409e565b1590506111655760405162461bcd60e51b81526020600482015260156024820152740556e65787065637465642068656164657220524c5605c1b604482015260640161069b565b5f6111bd8960c0013589600101548b8060e001906111839190614a4b565b808060200260200160405190810160405280939291908181526020018383602002808284375f9201919091525050505060808d013561185e565b90508061120c5760405162461bcd60e51b815260206004820152601c60248201527f496e76616c696420657865637574696f6e20726f6f742070726f6f6600000000604482015260640161069b565b5f6112668a61010001358a600401548c80610120019061122c9190614a4b565b808060200260200160405190810160405280939291908181526020018383602002808284375f9201919091525050505060c08e013561185e565b9050806112ad5760405162461bcd60e51b8152602060048201526015602482015274125b9d985b1a59081c9958d95a5c1d1cc81c9bdbdd605a1b604482015260640161069b565b8961010001359750505050505050505b9392505050565b6040805160018082528183019092526060915f9190816020015b60608152602001906001900390816112de579050509050611302602084018461409e565b8080601f0160208091040260200160405190810160405280939291908181526020018383808284375f92018290525085518694509092501515905061134957611349614ab7565b60209081029190910101525f611372856113638680614a4b565b61136c91614acb565b84611b21565b905080516001146113d85760405162461bcd60e51b815260206004820152602a60248201527f496e76616c6964206e756d626572206f6620726573756c747320696e2072656360448201526932b4b83a10383937b7b360b11b606482015260840161069b565b5f815f815181106113eb576113eb614ab7565b602002602001015160200151511161143d5760405162461bcd60e51b815260206004820152601560248201527424b73b30b634b2103932b1b2b4b83a10383937b7b360591b604482015260640161069b565b5f61146c611467835f8151811061145657611456614ab7565b602002602001015160200151611f84565b611fb0565b90508060800151518560600135106114ba5760405162461bcd60e51b8152602060048201526011602482015270092dcecc2d8d2c840d8dece40d2dcc8caf607b1b604482015260640161069b565b5f81608001518660600135815181106114d5576114d5614ab7565b602002602001015190508560800160208101906114f29190613cd1565b6001600160a01b0316815f01516001600160a01b03161461154b5760405162461bcd60e51b815260206004820152601360248201527224b73b30b634b2103637b39032b6b4ba3a32b960691b604482015260640161069b565b5f816020015151116115935760405162461bcd60e51b81526020600482015260116024820152704c6f6720686173206e6f20746f7069637360781b604482015260640161069b565b8560a0013581602001515f815181106115ae576115ae614ab7565b6020026020010151146116035760405162461bcd60e51b815260206004820152601760248201527f496e76616c6964206576656e74207369676e6174757265000000000000000000604482015260640161069b565b604001519695505050505050565b60208083015180518051600a80546001600160401b0392831667ffffffffffffffff199182161790915591840151600b55828401518051600c8054918416919094168117909355840151600d558551938401515160405192949116927f955222ca45c52776cd8084d242443738d71afc33822d853b8efeb94e070d94d09261169a929190614b15565b60405180910390a36108cc8183602001516020015160200151612263565b60408051608081018252600a546001600160401b03908116828401908152600b54606084015282528251808401909352600c54168252600d54602083810191909152810191909152825161170c91906122d3565b61172957604051634247f10960e11b815260040160405180910390fd5b61173a825f015183602001516122fb565b61175757604051632d776d6b60e11b815260040160405180910390fd5b5f6002611764838061409e565b604051611772929190614b31565b602060405180830381855afa15801561178d573d5f5f3e3d5ffd5b5050506040513d601f19601f820116820180604052508101906117b09190614b40565b6010549091506001600160a01b031663ab750e756117d1602085018561409e565b600f54856040518563ffffffff1660e01b81526004016117f49493929190614b57565b5f6040518083038186803b15801561180a575f5ffd5b505afa15801561181c573d5f5f3e3d5ffd5b50505050505050565b61182f828261060a565b6108cc5760405163e2517d3f60e01b81526001600160a01b03821660048201526024810183905260440161069b565b5f6118db84600160806fffffffffffffffffffffffffffffffff83110291821c60406001600160401b0382110290811c602063ffffffff82110290811c601061ffff82110290811c600860ff82110290811c6004600f82110290811c6002600382110290811c969096119490960192909201010192909201010190565b8351146118e957505f6118f9565b816118f58686866123b8565b1490505b949350505050565b5f611910610120840184614a4b565b15905061195f5760405162461bcd60e51b815260206004820152601960248201527f556e65787065637465642072656365697074732070726f6f6600000000000000604482015260640161069b565b5f6119b78460c001358460080154868060e0019061197d9190614a4b565b808060200260200160405190810160405280939291908181526020018383602002808284375f92019190915250505050608088013561185e565b905080611a115760405162461bcd60e51b815260206004820152602260248201527f496e76616c696420657865637574696f6e20626c6f636b20686173682070726f60448201526137b360f11b606482015260840161069b565b60c0840135611a2461014086018661409e565b604051611a32929190614b31565b604051809103902014611a875760405162461bcd60e51b815260206004820181905260248201527f48656164657220646f6573206e6f74206d6174636820626c6f636b2068617368604482015260640161069b565b5f611ad2611a9961014087018761409e565b8080601f0160208091040260200160405190810160405280939291908181526020018383808284375f9201919091525061252692505050565b905084610100013581146118f95760405162461bcd60e51b81526020600482015260166024820152750a4cac6cad2e0e8e640e4dedee840dad2e6dac2e8c6d60531b604482015260640161069b565b60605f82516001600160401b03811115611b3d57611b3d613f09565b604051908082528060200260200182016040528015611b8257816020015b6040805180820190915260608082526020820152815260200190600190039081611b5b5790505b5090505f84516001600160401b03811115611b9f57611b9f613f09565b604051908082528060200260200182016040528015611be457816020015b604080518082019091525f815260606020820152815260200190600190039081611bbd5790505b5090505f5b8551811015611c65576040518060400160405280878381518110611c0f57611c0f614ab7565b6020026020010151805190602001208152602001878381518110611c3557611c35614ab7565b6020026020010151815250828281518110611c5257611c52614ab7565b6020908102919091010152600101611be9565b505f5b8451811015611f7957848181518110611c8357611c83614ab7565b6020026020010151838281518110611c9d57611c9d614ab7565b60200260200101515f01819052505f6040518060400160405280878481518110611cc957611cc9614ab7565b602002602001015181526020015f81525090505f611cef611cea858b6126c2565b61275d565b905060015b8015611f6e57604080516080810182525f8082526020820181905291810191909152606080820152611d25836128b1565b15611dee575f611d34846128c5565b90505f60028660200151611d489190614aa4565b15611d6e5760028660200151611d5e9190614b7d565b611d69906001613ef6565b611d7f565b60028660200151611d7f9190614b7d565b90506040518060400160405280611d99885f01518461296a565b81526020015f8152509550611db1825f015187612a32565b15611de657611dc4888360200151612a57565b898881518110611dd657611dd6614ab7565b6020026020010151602001819052505b505050611f6e565b60c083015115611e87575f611e0284612a9f565b9050611e1185825f0151612bbb565b15611e7a575f611e23825f0151612bd7565b8660200151611e329190613ef6565b90506040518060400160405280611e58885f0151600285611e539190614b7d565b61296a565b8152602001611e68600284614aa4565b81525095508160200151925050611e81565b5050611f6e565b50611f4b565b60e083015115611f3e575f611e9b84612bf8565b9050611ea685612dec565b15611ef0578051511515600103611e7a57611ec887825f015160200151612a57565b888781518110611eda57611eda614ab7565b6020026020010151602001819052505050611f6e565b5f8160200151611f00875f612dfd565b60108110611f1057611f10614ab7565b602002015180519091501515600103611de657611f2e866001612e83565b9550806020015192505050611f4b565b825115611f4b5750611f6e565b611f58611cea8783612a57565b9250508080611f6690614b90565b915050611cf4565b505050600101611c68565b509095945050505050565b6040805180820182525f8082526020918201528151808301909252825182529182019181019190915290565b6040805160a0810182525f808252606060208301819052928201528181018290526080810191909152611fe282612ec4565b61207d57602082015180515f1a60018110801590612004575060048160ff1611155b6120465760405162461bcd60e51b8152602060048201526013602482015272556e737570706f72746564207478207479706560681b604482015260640161069b565b60ff81168352604080518082019091528451819061206690600190614ba8565b8152602001612076846001613ef6565b9052935050505b5f61208783612efb565b905080516004146120e55760405162461bcd60e51b815260206004820152602260248201527f496e76616c696420726563656970743a2065787065637465642034206669656c604482015261647360f01b606482015260840161069b565b612107815f815181106120fa576120fa614ab7565b6020026020010151613000565b82602001819052506121328160018151811061212557612125614ab7565b6020026020010151613071565b6001600160401b03166040830152805161215990829060029081106120fa576120fa614ab7565b82606001819052505f6121858260038151811061217857612178614ab7565b6020026020010151612efb565b8051909150806001600160401b038111156121a2576121a2613f09565b6040519080825280602002602001820160405280156121ff57816020015b6121ec60405180606001604052805f6001600160a01b0316815260200160608152602001606081525090565b8152602001906001900390816121c05790505b5060808501525f5b8181101561225a5761223183828151811061222457612224614ab7565b60200260200101516130bb565b8560800151828151811061224757612247614ab7565b6020908102919091010152600101612207565b50505050919050565b6001600160401b0382165f908152600e602052604090205461229a576001600160401b0382165f908152600e602052604090208190555b60405181906001600160401b038416907f9f2417e1c378e9eae44d76184be142d77ed574882fbc98272fe0533053d8816d905f90a35050565b815181515f916122e29161328f565b80156112bd57506112bd8360200151836020015161328f565b6020808301515190820151515f91906001600160401b0380831690821611612327575f925050506104cc565b5f61233182610534565b905042811115612346575f93505050506104cc565b601054600160a01b900462ffffff1661235f8242614ba8565b1115612370575f93505050506104cc565b5f600c602061237f8686614a2c565b6001600160401b03166123929190613edf565b61239c9190613edf565b601054600160a01b900462ffffff161015979650505050505050565b5f83815b8460011461251d576123cf600286614aa4565b60010361246c5760028482815181106123ea576123ea614ab7565b60200260200101518360405160200161240d929190918252602082015260400190565b60408051601f198184030181529082905261242791614bbb565b602060405180830381855afa158015612442573d5f5f3e3d5ffd5b5050506040513d601f19601f820116820180604052508101906124659190614b40565b91506124fe565b60028285838151811061248157612481614ab7565b60200260200101516040516020016124a3929190918252602082015260400190565b60408051601f19818403018152908290526124bd91614bbb565b602060405180830381855afa1580156124d8573d5f5f3e3d5ffd5b5050506040513d601f19601f820116820180604052508101906124fb9190614b40565b91505b612509600286614b7d565b94508061251581614b90565b9150506123bc565b50949350505050565b5f5f61253183611f84565b905061253c81612ec4565b61257f5760405162461bcd60e51b815260206004820152601460248201527312195859195c881a5cc81b9bdd0818481b1a5cdd60621b604482015260640161069b565b5f612589826132b8565b90505f5b60058110156125f95761259f82613321565b6125e75760405162461bcd60e51b81526020600482015260196024820152784865616465722068617320746f6f20666577206669656c647360381b604482015260640161069b565b6125f082613344565b5060010161258d565b5061260381613321565b61264b5760405162461bcd60e51b81526020600482015260196024820152784865616465722068617320746f6f20666577206669656c647360381b604482015260640161069b565b5f61265d61265883613344565b613000565b905080516020146126b05760405162461bcd60e51b815260206004820152601f60248201527f496e76616c696420726563656970747320726f6f7420696e2068656164657200604482015260640161069b565b6126b981614bcc565b95945050505050565b60605f5b835181101561272057828482815181106126e2576126e2614ab7565b60200260200101515f0151036127185783818151811061270457612704614ab7565b6020026020010151602001519150506104cc565b6001016126c6565b5060405162461bcd60e51b8152602060048201526011602482015270496e636f6d706c6574652050726f6f662160781b604482015260640161069b565b612765613ab7565b61276d613ab7565b5f60405180604001604052808581526020015f81525090506127c4846040518060400160405280602081526020017f56e81f171bcc55a6ff8345e692c0f86e5b48e01b996cadc001622fb5e363b42181525061339e565b156127d457506001815292915050565b5f6127e66127e186611f84565b612efb565b80519091505f819003612800575050600182525092915050565b80600203612858575f61281e835f815181106120fa576120fa614ab7565b602081015190915060fc1c60028114806128385750806003145b156128495760016020870152612851565b600160c08701525b50506128a3565b8060110361286c57600160e08501526128a3565b60405162461bcd60e51b815260206004820152600c60248201526b496e76616c6964206461746160a01b604482015260640161069b565b505061012082015292915050565b5f8160200151806104cc5750506040015190565b6128cd613b1c565b6128d5613b1c565b5f6128ea6127e18561012001515f0151611f84565b90505f612903826001815181106120fa576120fa614ab7565b9050604051806040016040528061292f612928855f815181106120fa576120fa614ab7565b60016133c6565b81525f602091820181905291855260408051608081018252838152808301939093526001908301526060820192909252908301525092915050565b81516060905f61297a8483614ba8565b9050838210156129c05760405162461bcd60e51b8152602060048201526011602482015270736c6963655f6f75744f66426f756e647360781b604482015260640161069b565b6060811580156129de5760405191505f825260208201604052612a28565b6040519150601f8316801560200281840101848101888315602002848c0101015b81831015612a175780518352602092830192016129ff565b5050848452601f01601f1916604052505b5095945050505050565b5f612a3c82612bd7565b612a4584612bd7565b1480156112bd57506112bd8383612bbb565b6060816040015115612a6e575060608101516104cc565b815115612a8a57612a838383602001516126c2565b90506104cc565b5060408051602081019091525f815292915050565b612aa7613b1c565b612aaf613b1c565b5f612ac46127e18561012001515f0151611f84565b90505f612add826001815181106120fa576120fa614ab7565b90505f6004612af7845f815181106120fa576120fa614ab7565b5f81518110612b0857612b08614ab7565b602001015160f81c60f81b6001600160f81b031916901c60f81c60011690506040518060400160405280612b69612b4a865f815181106120fa576120fa614ab7565b6002612b57866001614bef565b612b619190614c08565b60ff166133c6565b815260ff8316602091820152908552604080516080810190915260018152908101612b93846133fc565b81525f6020808301829052604080519283528282018152909201528501525091949350505050565b5f612bc582612bd7565b612bcf8484613463565b149392505050565b5f81602001516002835f015151612bee9190613edf565b6104cc9190614ba8565b612c00613b6b565b612c08613b6b565b5f612c1d6127e18561012001515f0151611f84565b9050612c27613b87565b5f5b6010811015612d20575f612c488483815181106120fa576120fa614ab7565b90508051602014612caf576040805180820182525f808252825160808101845281815260208082018390528185018390528451928352828101909452606081019190915291810191909152838360108110612ca557612ca5614ab7565b6020020152612d17565b5f612cb9826133fc565b6040805180820182526001808252825160808101845290815260208082018590525f8285018190528451908152808201909452606082019390935291810191909152909150848460108110612d1057612d10614ab7565b6020020152505b50600101612c29565b50612d3f612d3a836010815181106120fa576120fa614ab7565b6135d5565b15612d8b576040805180820182525f8082528251608081018452818152602080820183905281850183905284519283528281019094526060810191909152918101919091528352612de0565b604051806040016040528060011515815260200160405180608001604052805f151581526020015f5f1b8152602001600115158152602001612dd9866010815181106120fa576120fa614ab7565b9052905283525b60208301525092915050565b5f612df682612bd7565b1592915050565b5f5f6002838560200151612e119190613ef6565b612e1b9190614b7d565b90505f6002848660200151612e309190613ef6565b612e3a9190614aa4565b90505f855f01518381518110612e5257612e52614ab7565b016020015160f81c905060018214612e715760048160ff16901c612e76565b80600f165b60ff169695505050505050565b60408051808201909152606081525f60208201526040518060400160405280845f01518152602001838560200151612ebb9190613ef6565b90529392505050565b80515f908103612ed557505f919050565b602082015180515f1a9060c0821015612ef157505f9392505050565b5060019392505050565b6060612f0682612ec4565b612f0e575f5ffd5b5f612f188361363d565b90505f816001600160401b03811115612f3357612f33613f09565b604051908082528060200260200182016040528015612f7757816020015b604080518082019091525f8082526020820152815260200190600190039081612f515790505b5090505f612f8885602001516136c0565b8560200151612f979190613ef6565b90505f805b84811015612ff557612fad8361373f565b9150604051806040016040528083815260200184815250848281518110612fd657612fd6614ab7565b6020908102919091010152612feb8284613ef6565b9250600101612f9c565b509195945050505050565b805160609061300d575f5ffd5b5f5f613018846137e7565b915091505f816001600160401b0381111561303557613035613f09565b6040519080825280601f01601f19166020018201604052801561305f576020820181803683370190505b5090506020810161251d848285613829565b80515f901580159061308557508151602110155b61308d575f5ffd5b5f5f613098846137e7565b8151919350915060208210156118f95760208290036101000a9004949350505050565b6130e760405180606001604052805f6001600160a01b0316815260200160608152602001606081525090565b5f6130f183612efb565b905080516003146131445760405162461bcd60e51b815260206004820152601e60248201527f496e76616c6964206c6f673a2065787065637465642033206669656c64730000604482015260640161069b565b613166815f8151811061315957613159614ab7565b60200260200101516138a9565b6001600160a01b03168252805161318a90829060029081106120fa576120fa614ab7565b82604001819052505f6131a98260018151811061217857612178614ab7565b805190915060048111156131ff5760405162461bcd60e51b815260206004820152601c60248201527f496e76616c6964206c6f673a20746f6f206d616e7920746f7069637300000000604482015260640161069b565b806001600160401b0381111561321757613217613f09565b604051908082528060200260200182016040528015613240578160200160208202803683370190505b5060208501525f5b8181101561225a5761326583828151811061212557612125614ab7565b602086015180518390811061327c5761327c614ab7565b6020908102919091010152600101613248565b805182515f916001600160401b0391821691161480156112bd5750506020908101519101511490565b604080516080810182525f91810182815260608201839052815260208101919091526132e382612ec4565b6132eb575f5ffd5b5f6132f983602001516136c0565b83602001516133089190613ef6565b6040805180820190915293845260208401525090919050565b8051805160208201515f929161333691613ef6565b836020015110915050919050565b604080518082019091525f808252602082015261336082613321565b613368575f5ffd5b60208201515f6133778261373f565b90506133838183613ef6565b60209485015260408051808201909152908152928301525090565b5f81518351146133af57505f6104cc565b508151602091820181902091909201919091201490565b606082518211156133d5575f5ffd5b5f8284516133e39190614ba8565b9050602084016126b96133f68583613ef6565b836138c1565b5f60208251101561345b5760405162461bcd60e51b8152602060048201526024808201527f42797465733a3a20746f427974657333323a206461746120697320746f20736860448201526337b93a1760e11b606482015260840161069b565b506020015190565b5f5f600284602001516134769190614aa4565b90505f6002846020015161348a9190614aa4565b905080820361357a575f600286602001516134a59190614b7d565b90505f600286602001516134b99190614b7d565b90505f84156135375786518051839081106134d6576134d6614ab7565b016020015188518051600f60f81b90921691859081106134f8576134f8614ab7565b0160200151600f60f81b1614613515575f955050505050506104cc565b61351e83614b90565b925061352982614b90565b915061353481614b90565b90505b5f613545895f01518561296a565b90505f613555895f01518561296a565b9050826135628383613917565b61356c9190613ef6565b9750505050505050506104cc565b5f61359561358787612bd7565b61359087612bd7565b6139fd565b90505f5b818110156135ca576135ab8682612dfd565b6135b58883612dfd565b036135ca576135c381614b90565b9050613599565b93506104cc92505050565b5f5f82511180156104cc5750815f815181106135f3576135f3614ab7565b6020910101516001600160f81b031916600360fe1b14806104cc5750815f8151811061362157613621614ab7565b6020910101516001600160f81b031916600160ff1b1492915050565b80515f90810361364e57505f919050565b5f5f90505f61366084602001516136c0565b846020015161366f9190613ef6565b90505f845f015185602001516136859190613ef6565b90505b808210156136b7576136998261373f565b6136a39083613ef6565b9150826136af81614b90565b935050613688565b50909392505050565b80515f90811a60808110156136d757505f92915050565b60b88110806136f2575060c081108015906136f2575060f881105b156137005750600192915050565b60c081101561372d57613715600160b8614c29565b6137229060ff1682614ba8565b6112bd906001613ef6565b613715600160f8614c29565b50919050565b80515f908190811a608081101561375957600191506137e0565b60b881101561377f5761376d608082614ba8565b613778906001613ef6565b91506137e0565b60c08110156137ac5760b78103600185019450806020036101000a855104600182018101935050506137e0565b60f88110156137c05761376d60c082614ba8565b60f78103600185019450806020036101000a855104600182018101935050505b5092915050565b5f5f5f6137f784602001516136c0565b90505f81856020015161380a9190613ef6565b90505f82865f015161381c9190614ba8565b9196919550909350505050565b805f0361383557505050565b6020811061386d578251825261384c602084613ef6565b9250613859602083613ef6565b9150613866602082614ba8565b9050613835565b801561052f575f6001613881836020614ba8565b61388d90610100614d25565b6138979190614ba8565b84518451821691191617835250505050565b80515f906015146138b8575f5ffd5b6104cc82613071565b6060816001600160401b038111156138db576138db613f09565b6040519080825280601f01601f191660200182016040528015613905576020820181803683370190505b509050602081016137e0848285613a12565b5f5f613925845184516139fd565b90505f5b818110156139f25783818151811061394357613943614ab7565b602001015160f81c60f81b6001600160f81b03191685828151811061396a5761396a614ab7565b01602001516001600160f81b031916146139e2576139c485828151811061399357613993614ab7565b602001015160f81c60f81b8583815181106139b0576139b0614ab7565b01602001516001600160f81b031916613a79565b6139cf600283613edf565b6139d99190613ef6565b925050506104cc565b6139eb81614b90565b9050613929565b6126b9600282613edf565b5f818310613a0b57816112bd565b5090919050565b60208110613a4a5782518252613a29602083613ef6565b9150613a36602084613ef6565b9250613a43602082614ba8565b9050613a12565b5f81156104f6576001613a5e836020614ba8565b613a6a90610100614d25565b613a749190614ba8565b613897565b5f6001600160f81b031980831690841603613a96575060026104cc565b8160f81c60f01660ff168360f81c60f01660ff1603610da2575060016104cc565b60408051610140810182525f808252602080830182905282840182905260608084018390526080840183905260a0840183905260c0840183905260e084018390526101008401839052845180860190955284528301529061012082015290565b905290565b6040518060400160405280613b436040518060400160405280606081526020015f81525090565b8152604080516080810182525f80825260208281018290529282015260608082015291015290565b6040518060400160405280613b7e613bb5565b8152602001613b175b6040518061020001604052806010905b613b9f613bb5565b815260200190600190039081613b975790505090565b60405180604001604052805f15158152602001613b17604080516080810182525f808252602082018190529181019190915260608082015290565b5f60208284031215613c00575f5ffd5b81356001600160e01b0319811681146112bd575f5ffd5b5f60208284031215613c27575f5ffd5b5035919050565b6001600160a01b0381168114610d17575f5ffd5b8035613c4d81613c2e565b919050565b5f5f60408385031215613c63575f5ffd5b823591506020830135613c7581613c2e565b809150509250929050565b80356001600160401b0381168114613c4d575f5ffd5b5f60208284031215613ca6575f5ffd5b6112bd82613c80565b5f60208284031215613cbf575f5ffd5b813562ffffff811681146112bd575f5ffd5b5f60208284031215613ce1575f5ffd5b81356112bd81613c2e565b5f5f60408385031215613cfd575f5ffd5b82356001600160401b03811115613d12575f5ffd5b83016101608186031215613d24575f5ffd5b915060208301356001600160401b03811115613d3e575f5ffd5b830160c08186031215613c75575f5ffd5b5f5b83811015613d69578181015183820152602001613d51565b50505f910152565b5f8151808452613d88816020860160208601613d4f565b601f01601f19169290920160200192915050565b602081525f6112bd6020830184613d71565b5f5f5f60408486031215613dc0575f5ffd5b83356001600160401b03811115613dd5575f5ffd5b8401601f81018613613de5575f5ffd5b80356001600160401b03811115613dfa575f5ffd5b866020828401011115613e0b575f5ffd5b602091820194509250613e1f908501613c80565b90509250925092565b5f60208284031215613e38575f5ffd5b81356001600160401b03811115613e4d575f5ffd5b8201606081850312156112bd575f5ffd5b5f60208284031215613e6e575f5ffd5b81356001600160401b03811115613e83575f5ffd5b820161012081850312156112bd575f5ffd5b5f60208284031215613ea5575f5ffd5b81356001600160401b03811115613eba575f5ffd5b8201608081850312156112bd575f5ffd5b634e487b7160e01b5f52601160045260245ffd5b80820281158282048414176104cc576104cc613ecb565b808201808211156104cc576104cc613ecb565b634e487b7160e01b5f52604160045260245ffd5b604080519081016001600160401b0381118282101715613f3f57613f3f613f09565b60405290565b60405160c081016001600160401b0381118282101715613f3f57613f3f613f09565b60405161016081016001600160401b0381118282101715613f3f57613f3f613f09565b604051601f8201601f191681016001600160401b0381118282101715613fb257613fb2613f09565b604052919050565b5f60408284031215613fca575f5ffd5b613fd2613f1d565b9050613fdd82613c80565b815260209182013591810191909152919050565b5f60808284031215614001575f5ffd5b614009613f1d565b90506140158383613fba565b81526140248360408401613fba565b602082015292915050565b5f610120828403128015614041575f5ffd5b50604051606081016001600160401b038111828210171561406457614064613f09565b6040526140718484613ff1565b81526140808460808501613ff1565b60208201526140926101008401613c80565b60408201529392505050565b5f5f8335601e198436030181126140b3575f5ffd5b8301803591506001600160401b038211156140cc575f5ffd5b6020019150368190038213156140e0575f5ffd5b9250929050565b5f602082840312156140f7575f5ffd5b81516112bd81613c2e565b5f5f8335601e19843603018112614117575f5ffd5b83016020810192503590506001600160401b03811115614135575f5ffd5b8060051b36038213156140e0575f5ffd5b8183526020830192505f815f5b8481101561418457813561416681613c2e565b6001600160a01b031686526020958601959190910190600101614153565b5093949350505050565b5f5f8335601e198436030181126141a3575f5ffd5b83016020810192503590506001600160401b038111156141c1575f5ffd5b8060061b36038213156140e0575f5ffd5b8183526020830192505f815f5b84811015614184578135865260208201356141f981613c2e565b6001600160a01b0316602087015260409586019591909101906001016141df565b5f5f8335601e1984360301811261422f575f5ffd5b83016020810192503590506001600160401b0381111561424d575f5ffd5b8036038213156140e0575f5ffd5b81835281816020850137505f828201602090810191909152601f909101601f19169091010190565b60208082528235828201525f9061429b908401613c42565b6001600160a01b0381166040840152506142b760408401613c42565b6001600160a01b03811660608401525060608301356080838101919091526142e0908401613c42565b6001600160a01b03811660a08401525060a083013560c08381019190915261430a90840184614102565b61012060e085015261432161014085018284614146565b91505061433160e085018561418e565b848303601f19016101008601526143498382846141d2565b9250505061435b61010085018561421a565b848303601f190161012086015261437383828461425b565b9695505050505050565b5f6001600160401b0382111561439557614395613f09565b5060051b60200190565b5f82601f8301126143ae575f5ffd5b81356143c16143bc8261437d565b613f8a565b8082825260208201915060208360051b8601019250858311156143e2575f5ffd5b602085015b83811015612a285780358352602092830192016143e7565b5f6001600160401b0382111561441757614417613f09565b50601f01601f191660200190565b5f82601f830112614434575f5ffd5b81356144426143bc826143ff565b818152846020838601011115614456575f5ffd5b816020850160208301375f918101602001919091529392505050565b5f61447f6143bc8461437d565b838152905060208101600584901b83018581111561449b575f5ffd5b835b818110156144d55780356001600160401b038111156144ba575f5ffd5b6144c688828801614425565b8452506020928301920161449d565b5050509392505050565b5f82601f8301126144ee575f5ffd5b6112bd83833560208501614472565b5f60c0828403121561450d575f5ffd5b614515613f45565b905081356001600160401b0381111561452c575f5ffd5b614538848285016144df565b82525060208201356001600160401b03811115614553575f5ffd5b61455f84828501614425565b60208301525060408201356001600160401b0381111561457d575f5ffd5b61458984828501614425565b604083015250606082810135908201526145a560808301613c42565b608082015260a09182013591810191909152919050565b5f602082840312156145cc575f5ffd5b81356001600160401b038111156145e1575f5ffd5b8201604081850312156145f2575f5ffd5b6145fa613f1d565b81356001600160401b0381111561460f575f5ffd5b82016101608187031215614621575f5ffd5b614629613f67565b61463282613c80565b815261464060208301613c80565b60208201526040828101359082015260608201356001600160401b03811115614667575f5ffd5b6146738882850161439f565b6060830152506080828101359082015260a08201356001600160401b0381111561469b575f5ffd5b6146a78882850161439f565b60a08301525060c0828101359082015260e08201356001600160401b038111156146cf575f5ffd5b6146db8882850161439f565b60e08301525061010082810135908201526101208201356001600160401b03811115614705575f5ffd5b6147118882850161439f565b610120830152506101408201356001600160401b03811115614731575f5ffd5b61473d88828501614425565b6101408301525082525060208201356001600160401b0381111561475f575f5ffd5b61476b868285016144fd565b602083015250949350505050565b5f60208284031215614789575f5ffd5b813563ffffffff811681146112bd575f5ffd5b5f8151808452602084019350602083015f5b828110156141845781518652602095860195909101906001016147ae565b5f60c08301825160c0855281815180845260e08701915060e08160051b88010193506020830192505f5b818110156148275760df19888603018352614812858551613d71565b945060209384019392909201916001016147f6565b50505050602083015184820360208601526148428282613d71565b9150506040830151848203604086015261485c8282613d71565b91505060608301516060850152608083015161488360808601826001600160a01b03169052565b5060a083015160a08501528091505092915050565b604081526148b26040820184516001600160401b03169052565b5f60208401516148cd60608401826001600160401b03169052565b5060408401516080830152606084015161016060a08401526148f36101a084018261479c565b9050608085015160c084015260a0850151603f198483030160e085015261491a828261479c565b91505060c085015161010084015260e0850151603f1984830301610120850152614944828261479c565b915050610100850151610140840152610120850151603f1984830301610160850152614970828261479c565b915050610140850151603f19848303016101808501526149908282613d71565b91505082810360208401526126b981856147cc565b5f602082840312156149b5575f5ffd5b81516001600160401b038111156149ca575f5ffd5b8201601f810184136149da575f5ffd5b80516149e86143bc826143ff565b8181528560208385010111156149fc575f5ffd5b6126b9826020830160208601613d4f565b5f823561011e19833603018112614a22575f5ffd5b9190910192915050565b6001600160401b0382811682821603908111156104cc576104cc613ecb565b5f5f8335601e19843603018112614a60575f5ffd5b8301803591506001600160401b03821115614a79575f5ffd5b6020019150600581901b36038213156140e0575f5ffd5b634e487b7160e01b5f52601260045260245ffd5b5f82614ab257614ab2614a90565b500690565b634e487b7160e01b5f52603260045260245ffd5b5f6112bd368484614472565b614af582825180516001600160401b03168252602090810151910152565b60209081015180516001600160401b031660408401520151606090910152565b6101008101614b248285614ad7565b6112bd6080830184614ad7565b818382375f9101908152919050565b5f60208284031215614b50575f5ffd5b5051919050565b606081525f614b6a60608301868861425b565b6020830194909452506040015292915050565b5f82614b8b57614b8b614a90565b500490565b5f60018201614ba157614ba1613ecb565b5060010190565b818103818111156104cc576104cc613ecb565b5f8251614a22818460208701613d4f565b80516020808301519190811015613739575f1960209190910360031b1b16919050565b60ff81811683821601908111156104cc576104cc613ecb565b5f60ff831680614c1a57614c1a614a90565b8060ff84160691505092915050565b60ff82811682821603908111156104cc576104cc613ecb565b6001815b6001841115614c7d57808504811115614c6157614c61613ecb565b6001841615614c6f57908102905b60019390931c928002614c46565b935093915050565b5f82614c93575060016104cc565b81614c9f57505f6104cc565b8160018114614cb55760028114614cbf57614cdb565b60019150506104cc565b60ff841115614cd057614cd0613ecb565b50506001821b6104cc565b5060208310610133831016604e8410600b8410161715614cfe575081810a6104cc565b614d0a5f198484614c42565b805f1904821115614d1d57614d1d613ecb565b029392505050565b5f6112bd8383614c8556fea49807205ce4d355092ef5a8a18f56e8913cf4a201fbe287825b095693c21775a26469706673582212204003e53abfee4ab5f8b248d57a3dc2849d4ebeadf718e792d4493a79fd370e4d64736f6c634300081e0033",
}

// ZKAdapterABI is the input ABI used to generate the binding from.
// Deprecated: Use ZKAdapterMetaData.ABI instead.
var ZKAdapterABI = ZKAdapterMetaData.ABI

// ZKAdapterBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use ZKAdapterMetaData.Bin instead.
var ZKAdapterBin = ZKAdapterMetaData.Bin

// DeployZKAdapter deploys a new Ethereum contract, binding an instance of ZKAdapter to it.
func DeployZKAdapter(auth *bind.TransactOpts, backend bind.ContractBackend, sourceChainId_ *big.Int, startingState ConsensusState, beaconConfig_ ExecutionBeaconConfig, genesisTime_ *big.Int, permissibleTimespan_ *big.Int, verifier common.Address, imageID_ [32]byte, admin common.Address, superAdmin common.Address) (common.Address, *types.Transaction, *ZKAdapter, error) {
	parsed, err := ZKAdapterMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(ZKAdapterBin), backend, sourceChainId_, startingState, beaconConfig_, genesisTime_, permissibleTimespan_, verifier, imageID_, admin, superAdmin)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &ZKAdapter{ZKAdapterCaller: ZKAdapterCaller{contract: contract}, ZKAdapterTransactor: ZKAdapterTransactor{contract: contract}, ZKAdapterFilterer: ZKAdapterFilterer{contract: contract}}, nil
}

// ZKAdapter is an auto generated Go binding around an Ethereum contract.
type ZKAdapter struct {
	ZKAdapterCaller     // Read-only binding to the contract
	ZKAdapterTransactor // Write-only binding to the contract
	ZKAdapterFilterer   // Log filterer for contract events
}

// ZKAdapterCaller is an auto generated read-only Go binding around an Ethereum contract.
type ZKAdapterCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ZKAdapterTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ZKAdapterTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ZKAdapterFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ZKAdapterFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ZKAdapterSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ZKAdapterSession struct {
	Contract     *ZKAdapter        // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ZKAdapterCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ZKAdapterCallerSession struct {
	Contract *ZKAdapterCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts    // Call options to use throughout this session
}

// ZKAdapterTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ZKAdapterTransactorSession struct {
	Contract     *ZKAdapterTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// ZKAdapterRaw is an auto generated low-level Go binding around an Ethereum contract.
type ZKAdapterRaw struct {
	Contract *ZKAdapter // Generic contract binding to access the raw methods on
}

// ZKAdapterCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ZKAdapterCallerRaw struct {
	Contract *ZKAdapterCaller // Generic read-only contract binding to access the raw methods on
}

// ZKAdapterTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ZKAdapterTransactorRaw struct {
	Contract *ZKAdapterTransactor // Generic write-only contract binding to access the raw methods on
}

// NewZKAdapter creates a new instance of ZKAdapter, bound to a specific deployed contract.
func NewZKAdapter(address common.Address, backend bind.ContractBackend) (*ZKAdapter, error) {
	contract, err := bindZKAdapter(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ZKAdapter{ZKAdapterCaller: ZKAdapterCaller{contract: contract}, ZKAdapterTransactor: ZKAdapterTransactor{contract: contract}, ZKAdapterFilterer: ZKAdapterFilterer{contract: contract}}, nil
}

// NewZKAdapterCaller creates a new read-only instance of ZKAdapter, bound to a specific deployed contract.
func NewZKAdapterCaller(address common.Address, caller bind.ContractCaller) (*ZKAdapterCaller, error) {
	contract, err := bindZKAdapter(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ZKAdapterCaller{contract: contract}, nil
}

// NewZKAdapterTransactor creates a new write-only instance of ZKAdapter, bound to a specific deployed contract.
func NewZKAdapterTransactor(address common.Address, transactor bind.ContractTransactor) (*ZKAdapterTransactor, error) {
	contract, err := bindZKAdapter(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ZKAdapterTransactor{contract: contract}, nil
}

// NewZKAdapterFilterer creates a new log filterer instance of ZKAdapter, bound to a specific deployed contract.
func NewZKAdapterFilterer(address common.Address, filterer bind.ContractFilterer) (*ZKAdapterFilterer, error) {
	contract, err := bindZKAdapter(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ZKAdapterFilterer{contract: contract}, nil
}

// bindZKAdapter binds a generic wrapper to an already deployed contract.
func bindZKAdapter(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := ZKAdapterMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ZKAdapter *ZKAdapterRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ZKAdapter.Contract.ZKAdapterCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ZKAdapter *ZKAdapterRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ZKAdapter.Contract.ZKAdapterTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ZKAdapter *ZKAdapterRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ZKAdapter.Contract.ZKAdapterTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ZKAdapter *ZKAdapterCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ZKAdapter.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ZKAdapter *ZKAdapterTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ZKAdapter.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ZKAdapter *ZKAdapterTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ZKAdapter.Contract.contract.Transact(opts, method, params...)
}

// ADMINROLE is a free data retrieval call binding the contract method 0x75b238fc.
//
// Solidity: function ADMIN_ROLE() view returns(bytes32)
func (_ZKAdapter *ZKAdapterCaller) ADMINROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _ZKAdapter.contract.Call(opts, &out, "ADMIN_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// ADMINROLE is a free data retrieval call binding the contract method 0x75b238fc.
//
// Solidity: function ADMIN_ROLE() view returns(bytes32)
func (_ZKAdapter *ZKAdapterSession) ADMINROLE() ([32]byte, error) {
	return _ZKAdapter.Contract.ADMINROLE(&_ZKAdapter.CallOpts)
}

// ADMINROLE is a free data retrieval call binding the contract method 0x75b238fc.
//
// Solidity: function ADMIN_ROLE() view returns(bytes32)
func (_ZKAdapter *ZKAdapterCallerSession) ADMINROLE() ([32]byte, error) {
	return _ZKAdapter.Contract.ADMINROLE(&_ZKAdapter.CallOpts)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_ZKAdapter *ZKAdapterCaller) DEFAULTADMINROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _ZKAdapter.contract.Call(opts, &out, "DEFAULT_ADMIN_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_ZKAdapter *ZKAdapterSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _ZKAdapter.Contract.DEFAULTADMINROLE(&_ZKAdapter.CallOpts)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_ZKAdapter *ZKAdapterCallerSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _ZKAdapter.Contract.DEFAULTADMINROLE(&_ZKAdapter.CallOpts)
}

// SECONDSPERSLOT is a free data retrieval call binding the contract method 0x304b9071.
//
// Solidity: function SECONDS_PER_SLOT() view returns(uint64)
func (_ZKAdapter *ZKAdapterCaller) SECONDSPERSLOT(opts *bind.CallOpts) (uint64, error) {
	var out []interface{}
	err := _ZKAdapter.contract.Call(opts, &out, "SECONDS_PER_SLOT")

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// SECONDSPERSLOT is a free data retrieval call binding the contract method 0x304b9071.
//
// Solidity: function SECONDS_PER_SLOT() view returns(uint64)
func (_ZKAdapter *ZKAdapterSession) SECONDSPERSLOT() (uint64, error) {
	return _ZKAdapter.Contract.SECONDSPERSLOT(&_ZKAdapter.CallOpts)
}

// SECONDSPERSLOT is a free data retrieval call binding the contract method 0x304b9071.
//
// Solidity: function SECONDS_PER_SLOT() view returns(uint64)
func (_ZKAdapter *ZKAdapterCallerSession) SECONDSPERSLOT() (uint64, error) {
	return _ZKAdapter.Contract.SECONDSPERSLOT(&_ZKAdapter.CallOpts)
}

// SLOTPEREPOCH is a free data retrieval call binding the contract method 0x65242a0e.
//
// Solidity: function SLOT_PER_EPOCH() view returns(uint64)
func (_ZKAdapter *ZKAdapterCaller) SLOTPEREPOCH(opts *bind.CallOpts) (uint64, error) {
	var out []interface{}
	err := _ZKAdapter.contract.Call(opts, &out, "SLOT_PER_EPOCH")

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// SLOTPEREPOCH is a free data retrieval call binding the contract method 0x65242a0e.
//
// Solidity: function SLOT_PER_EPOCH() view returns(uint64)
func (_ZKAdapter *ZKAdapterSession) SLOTPEREPOCH() (uint64, error) {
	return _ZKAdapter.Contract.SLOTPEREPOCH(&_ZKAdapter.CallOpts)
}

// SLOTPEREPOCH is a free data retrieval call binding the contract method 0x65242a0e.
//
// Solidity: function SLOT_PER_EPOCH() view returns(uint64)
func (_ZKAdapter *ZKAdapterCallerSession) SLOTPEREPOCH() (uint64, error) {
	return _ZKAdapter.Contract.SLOTPEREPOCH(&_ZKAdapter.CallOpts)
}

// UNDEFINEDROOT is a free data retrieval call binding the contract method 0x54747fd4.
//
// Solidity: function UNDEFINED_ROOT() view returns(bytes32)
func (_ZKAdapter *ZKAdapterCaller) UNDEFINEDROOT(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _ZKAdapter.contract.Call(opts, &out, "UNDEFINED_ROOT")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// UNDEFINEDROOT is a free data retrieval call binding the contract method 0x54747fd4.
//
// Solidity: function UNDEFINED_ROOT() view returns(bytes32)
func (_ZKAdapter *ZKAdapterSession) UNDEFINEDROOT() ([32]byte, error) {
	return _ZKAdapter.Contract.UNDEFINEDROOT(&_ZKAdapter.CallOpts)
}

// UNDEFINEDROOT is a free data retrieval call binding the contract method 0x54747fd4.
//
// Solidity: function UNDEFINED_ROOT() view returns(bytes32)
func (_ZKAdapter *ZKAdapterCallerSession) UNDEFINEDROOT() ([32]byte, error) {
	return _ZKAdapter.Contract.UNDEFINEDROOT(&_ZKAdapter.CallOpts)
}

// BeaconRootAt is a free data retrieval call binding the contract method 0x8e1e49ba.
//
// Solidity: function beaconRootAt(uint64 slot) view returns(bytes32)
func (_ZKAdapter *ZKAdapterCaller) BeaconRootAt(opts *bind.CallOpts, slot uint64) ([32]byte, error) {
	var out []interface{}
	err := _ZKAdapter.contract.Call(opts, &out, "beaconRootAt", slot)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// BeaconRootAt is a free data retrieval call binding the contract method 0x8e1e49ba.
//
// Solidity: function beaconRootAt(uint64 slot) view returns(bytes32)
func (_ZKAdapter *ZKAdapterSession) BeaconRootAt(slot uint64) ([32]byte, error) {
	return _ZKAdapter.Contract.BeaconRootAt(&_ZKAdapter.CallOpts, slot)
}

// BeaconRootAt is a free data retrieval call binding the contract method 0x8e1e49ba.
//
// Solidity: function beaconRootAt(uint64 slot) view returns(bytes32)
func (_ZKAdapter *ZKAdapterCallerSession) BeaconRootAt(slot uint64) ([32]byte, error) {
	return _ZKAdapter.Contract.BeaconRootAt(&_ZKAdapter.CallOpts, slot)
}

// EpochTimestamp is a free data retrieval call binding the contract method 0x4b0d0fa8.
//
// Solidity: function epochTimestamp(uint64 epoch) view returns(uint256)
func (_ZKAdapter *ZKAdapterCaller) EpochTimestamp(opts *bind.CallOpts, epoch uint64) (*big.Int, error) {
	var out []interface{}
	err := _ZKAdapter.contract.Call(opts, &out, "epochTimestamp", epoch)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// EpochTimestamp is a free data retrieval call binding the contract method 0x4b0d0fa8.
//
// Solidity: function epochTimestamp(uint64 epoch) view returns(uint256)
func (_ZKAdapter *ZKAdapterSession) EpochTimestamp(epoch uint64) (*big.Int, error) {
	return _ZKAdapter.Contract.EpochTimestamp(&_ZKAdapter.CallOpts, epoch)
}

// EpochTimestamp is a free data retrieval call binding the contract method 0x4b0d0fa8.
//
// Solidity: function epochTimestamp(uint64 epoch) view returns(uint256)
func (_ZKAdapter *ZKAdapterCallerSession) EpochTimestamp(epoch uint64) (*big.Int, error) {
	return _ZKAdapter.Contract.EpochTimestamp(&_ZKAdapter.CallOpts, epoch)
}

// GenesisTime is a free data retrieval call binding the contract method 0x42c6498a.
//
// Solidity: function genesisTime() view returns(uint256)
func (_ZKAdapter *ZKAdapterCaller) GenesisTime(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _ZKAdapter.contract.Call(opts, &out, "genesisTime")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GenesisTime is a free data retrieval call binding the contract method 0x42c6498a.
//
// Solidity: function genesisTime() view returns(uint256)
func (_ZKAdapter *ZKAdapterSession) GenesisTime() (*big.Int, error) {
	return _ZKAdapter.Contract.GenesisTime(&_ZKAdapter.CallOpts)
}

// GenesisTime is a free data retrieval call binding the contract method 0x42c6498a.
//
// Solidity: function genesisTime() view returns(uint256)
func (_ZKAdapter *ZKAdapterCallerSession) GenesisTime() (*big.Int, error) {
	return _ZKAdapter.Contract.GenesisTime(&_ZKAdapter.CallOpts)
}

// GetBeaconBlockRoot is a free data retrieval call binding the contract method 0xd0850708.
//
// Solidity: function getBeaconBlockRoot(uint64 slot) view returns(bytes32 root, bool valid)
func (_ZKAdapter *ZKAdapterCaller) GetBeaconBlockRoot(opts *bind.CallOpts, slot uint64) (struct {
	Root  [32]byte
	Valid bool
}, error) {
	var out []interface{}
	err := _ZKAdapter.contract.Call(opts, &out, "getBeaconBlockRoot", slot)

	outstruct := new(struct {
		Root  [32]byte
		Valid bool
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Root = *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)
	outstruct.Valid = *abi.ConvertType(out[1], new(bool)).(*bool)

	return *outstruct, err

}

// GetBeaconBlockRoot is a free data retrieval call binding the contract method 0xd0850708.
//
// Solidity: function getBeaconBlockRoot(uint64 slot) view returns(bytes32 root, bool valid)
func (_ZKAdapter *ZKAdapterSession) GetBeaconBlockRoot(slot uint64) (struct {
	Root  [32]byte
	Valid bool
}, error) {
	return _ZKAdapter.Contract.GetBeaconBlockRoot(&_ZKAdapter.CallOpts, slot)
}

// GetBeaconBlockRoot is a free data retrieval call binding the contract method 0xd0850708.
//
// Solidity: function getBeaconBlockRoot(uint64 slot) view returns(bytes32 root, bool valid)
func (_ZKAdapter *ZKAdapterCallerSession) GetBeaconBlockRoot(slot uint64) (struct {
	Root  [32]byte
	Valid bool
}, error) {
	return _ZKAdapter.Contract.GetBeaconBlockRoot(&_ZKAdapter.CallOpts, slot)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_ZKAdapter *ZKAdapterCaller) GetRoleAdmin(opts *bind.CallOpts, role [32]byte) ([32]byte, error) {
	var out []interface{}
	err := _ZKAdapter.contract.Call(opts, &out, "getRoleAdmin", role)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_ZKAdapter *ZKAdapterSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _ZKAdapter.Contract.GetRoleAdmin(&_ZKAdapter.CallOpts, role)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_ZKAdapter *ZKAdapterCallerSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _ZKAdapter.Contract.GetRoleAdmin(&_ZKAdapter.CallOpts, role)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_ZKAdapter *ZKAdapterCaller) HasRole(opts *bind.CallOpts, role [32]byte, account common.Address) (bool, error) {
	var out []interface{}
	err := _ZKAdapter.contract.Call(opts, &out, "hasRole", role, account)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_ZKAdapter *ZKAdapterSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _ZKAdapter.Contract.HasRole(&_ZKAdapter.CallOpts, role, account)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_ZKAdapter *ZKAdapterCallerSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _ZKAdapter.Contract.HasRole(&_ZKAdapter.CallOpts, role, account)
}

// ImageID is a free data retrieval call binding the contract method 0x4d3499b5.
//
// Solidity: function imageID() view returns(bytes32)
func (_ZKAdapter *ZKAdapterCaller) ImageID(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _ZKAdapter.contract.Call(opts, &out, "imageID")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// ImageID is a free data retrieval call binding the contract method 0x4d3499b5.
//
// Solidity: function imageID() view returns(bytes32)
func (_ZKAdapter *ZKAdapterSession) ImageID() ([32]byte, error) {
	return _ZKAdapter.Contract.ImageID(&_ZKAdapter.CallOpts)
}

// ImageID is a free data retrieval call binding the contract method 0x4d3499b5.
//
// Solidity: function imageID() view returns(bytes32)
func (_ZKAdapter *ZKAdapterCallerSession) ImageID() ([32]byte, error) {
	return _ZKAdapter.Contract.ImageID(&_ZKAdapter.CallOpts)
}

// PermissibleTimespan is a free data retrieval call binding the contract method 0x04ae4327.
//
// Solidity: function permissibleTimespan() view returns(uint24)
func (_ZKAdapter *ZKAdapterCaller) PermissibleTimespan(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _ZKAdapter.contract.Call(opts, &out, "permissibleTimespan")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PermissibleTimespan is a free data retrieval call binding the contract method 0x04ae4327.
//
// Solidity: function permissibleTimespan() view returns(uint24)
func (_ZKAdapter *ZKAdapterSession) PermissibleTimespan() (*big.Int, error) {
	return _ZKAdapter.Contract.PermissibleTimespan(&_ZKAdapter.CallOpts)
}

// PermissibleTimespan is a free data retrieval call binding the contract method 0x04ae4327.
//
// Solidity: function permissibleTimespan() view returns(uint24)
func (_ZKAdapter *ZKAdapterCallerSession) PermissibleTimespan() (*big.Int, error) {
	return _ZKAdapter.Contract.PermissibleTimespan(&_ZKAdapter.CallOpts)
}

// SourceChainId is a free data retrieval call binding the contract method 0x1544298e.
//
// Solidity: function sourceChainId() view returns(uint256)
func (_ZKAdapter *ZKAdapterCaller) SourceChainId(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _ZKAdapter.contract.Call(opts, &out, "sourceChainId")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// SourceChainId is a free data retrieval call binding the contract method 0x1544298e.
//
// Solidity: function sourceChainId() view returns(uint256)
func (_ZKAdapter *ZKAdapterSession) SourceChainId() (*big.Int, error) {
	return _ZKAdapter.Contract.SourceChainId(&_ZKAdapter.CallOpts)
}

// SourceChainId is a free data retrieval call binding the contract method 0x1544298e.
//
// Solidity: function sourceChainId() view returns(uint256)
func (_ZKAdapter *ZKAdapterCallerSession) SourceChainId() (*big.Int, error) {
	return _ZKAdapter.Contract.SourceChainId(&_ZKAdapter.CallOpts)
}

// SourceNetworkId is a free data retrieval call binding the contract method 0x6fda4c8b.
//
// Solidity: function sourceNetworkId() view returns(uint32)
func (_ZKAdapter *ZKAdapterCaller) SourceNetworkId(opts *bind.CallOpts) (uint32, error) {
	var out []interface{}
	err := _ZKAdapter.contract.Call(opts, &out, "sourceNetworkId")

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

// SourceNetworkId is a free data retrieval call binding the contract method 0x6fda4c8b.
//
// Solidity: function sourceNetworkId() view returns(uint32)
func (_ZKAdapter *ZKAdapterSession) SourceNetworkId() (uint32, error) {
	return _ZKAdapter.Contract.SourceNetworkId(&_ZKAdapter.CallOpts)
}

// SourceNetworkId is a free data retrieval call binding the contract method 0x6fda4c8b.
//
// Solidity: function sourceNetworkId() view returns(uint32)
func (_ZKAdapter *ZKAdapterCallerSession) SourceNetworkId() (uint32, error) {
	return _ZKAdapter.Contract.SourceNetworkId(&_ZKAdapter.CallOpts)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_ZKAdapter *ZKAdapterCaller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _ZKAdapter.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_ZKAdapter *ZKAdapterSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _ZKAdapter.Contract.SupportsInterface(&_ZKAdapter.CallOpts, interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_ZKAdapter *ZKAdapterCallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _ZKAdapter.Contract.SupportsInterface(&_ZKAdapter.CallOpts, interfaceId)
}

// Verifier is a free data retrieval call binding the contract method 0x2b7ac3f3.
//
// Solidity: function verifier() view returns(address)
func (_ZKAdapter *ZKAdapterCaller) Verifier(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ZKAdapter.contract.Call(opts, &out, "verifier")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Verifier is a free data retrieval call binding the contract method 0x2b7ac3f3.
//
// Solidity: function verifier() view returns(address)
func (_ZKAdapter *ZKAdapterSession) Verifier() (common.Address, error) {
	return _ZKAdapter.Contract.Verifier(&_ZKAdapter.CallOpts)
}

// Verifier is a free data retrieval call binding the contract method 0x2b7ac3f3.
//
// Solidity: function verifier() view returns(address)
func (_ZKAdapter *ZKAdapterCallerSession) Verifier() (common.Address, error) {
	return _ZKAdapter.Contract.Verifier(&_ZKAdapter.CallOpts)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_ZKAdapter *ZKAdapterTransactor) GrantRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _ZKAdapter.contract.Transact(opts, "grantRole", role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_ZKAdapter *ZKAdapterSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _ZKAdapter.Contract.GrantRole(&_ZKAdapter.TransactOpts, role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_ZKAdapter *ZKAdapterTransactorSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _ZKAdapter.Contract.GrantRole(&_ZKAdapter.TransactOpts, role, account)
}

// ManualTransition is a paid mutator transaction binding the contract method 0xc506e0be.
//
// Solidity: function manualTransition(bytes journalData, uint64 finalizedSlot) returns()
func (_ZKAdapter *ZKAdapterTransactor) ManualTransition(opts *bind.TransactOpts, journalData []byte, finalizedSlot uint64) (*types.Transaction, error) {
	return _ZKAdapter.contract.Transact(opts, "manualTransition", journalData, finalizedSlot)
}

// ManualTransition is a paid mutator transaction binding the contract method 0xc506e0be.
//
// Solidity: function manualTransition(bytes journalData, uint64 finalizedSlot) returns()
func (_ZKAdapter *ZKAdapterSession) ManualTransition(journalData []byte, finalizedSlot uint64) (*types.Transaction, error) {
	return _ZKAdapter.Contract.ManualTransition(&_ZKAdapter.TransactOpts, journalData, finalizedSlot)
}

// ManualTransition is a paid mutator transaction binding the contract method 0xc506e0be.
//
// Solidity: function manualTransition(bytes journalData, uint64 finalizedSlot) returns()
func (_ZKAdapter *ZKAdapterTransactorSession) ManualTransition(journalData []byte, finalizedSlot uint64) (*types.Transaction, error) {
	return _ZKAdapter.Contract.ManualTransition(&_ZKAdapter.TransactOpts, journalData, finalizedSlot)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address callerConfirmation) returns()
func (_ZKAdapter *ZKAdapterTransactor) RenounceRole(opts *bind.TransactOpts, role [32]byte, callerConfirmation common.Address) (*types.Transaction, error) {
	return _ZKAdapter.contract.Transact(opts, "renounceRole", role, callerConfirmation)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address callerConfirmation) returns()
func (_ZKAdapter *ZKAdapterSession) RenounceRole(role [32]byte, callerConfirmation common.Address) (*types.Transaction, error) {
	return _ZKAdapter.Contract.RenounceRole(&_ZKAdapter.TransactOpts, role, callerConfirmation)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address callerConfirmation) returns()
func (_ZKAdapter *ZKAdapterTransactorSession) RenounceRole(role [32]byte, callerConfirmation common.Address) (*types.Transaction, error) {
	return _ZKAdapter.Contract.RenounceRole(&_ZKAdapter.TransactOpts, role, callerConfirmation)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_ZKAdapter *ZKAdapterTransactor) RevokeRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _ZKAdapter.contract.Transact(opts, "revokeRole", role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_ZKAdapter *ZKAdapterSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _ZKAdapter.Contract.RevokeRole(&_ZKAdapter.TransactOpts, role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_ZKAdapter *ZKAdapterTransactorSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _ZKAdapter.Contract.RevokeRole(&_ZKAdapter.TransactOpts, role, account)
}

// SendMessage is a paid mutator transaction binding the contract method 0xeb97cd2c.
//
// Solidity: function sendMessage((uint256,address,address,bytes32,address,uint256,address[],(uint256,address)[],bytes) message) returns()
func (_ZKAdapter *ZKAdapterTransactor) SendMessage(opts *bind.TransactOpts, message TeleporterMessageV2) (*types.Transaction, error) {
	return _ZKAdapter.contract.Transact(opts, "sendMessage", message)
}

// SendMessage is a paid mutator transaction binding the contract method 0xeb97cd2c.
//
// Solidity: function sendMessage((uint256,address,address,bytes32,address,uint256,address[],(uint256,address)[],bytes) message) returns()
func (_ZKAdapter *ZKAdapterSession) SendMessage(message TeleporterMessageV2) (*types.Transaction, error) {
	return _ZKAdapter.Contract.SendMessage(&_ZKAdapter.TransactOpts, message)
}

// SendMessage is a paid mutator transaction binding the contract method 0xeb97cd2c.
//
// Solidity: function sendMessage((uint256,address,address,bytes32,address,uint256,address[],(uint256,address)[],bytes) message) returns()
func (_ZKAdapter *ZKAdapterTransactorSession) SendMessage(message TeleporterMessageV2) (*types.Transaction, error) {
	return _ZKAdapter.Contract.SendMessage(&_ZKAdapter.TransactOpts, message)
}

// Transition is a paid mutator transaction binding the contract method 0xdb2635b7.
//
// Solidity: function transition((bytes,bytes,uint64) consensus) returns()
func (_ZKAdapter *ZKAdapterTransactor) Transition(opts *bind.TransactOpts, consensus ConsensusData) (*types.Transaction, error) {
	return _ZKAdapter.contract.Transact(opts, "transition", consensus)
}

// Transition is a paid mutator transaction binding the contract method 0xdb2635b7.
//
// Solidity: function transition((bytes,bytes,uint64) consensus) returns()
func (_ZKAdapter *ZKAdapterSession) Transition(consensus ConsensusData) (*types.Transaction, error) {
	return _ZKAdapter.Contract.Transition(&_ZKAdapter.TransactOpts, consensus)
}

// Transition is a paid mutator transaction binding the contract method 0xdb2635b7.
//
// Solidity: function transition((bytes,bytes,uint64) consensus) returns()
func (_ZKAdapter *ZKAdapterTransactorSession) Transition(consensus ConsensusData) (*types.Transaction, error) {
	return _ZKAdapter.Contract.Transition(&_ZKAdapter.TransactOpts, consensus)
}

// UpdateImageID is a paid mutator transaction binding the contract method 0xe97eb362.
//
// Solidity: function updateImageID(bytes32 newImageID) returns()
func (_ZKAdapter *ZKAdapterTransactor) UpdateImageID(opts *bind.TransactOpts, newImageID [32]byte) (*types.Transaction, error) {
	return _ZKAdapter.contract.Transact(opts, "updateImageID", newImageID)
}

// UpdateImageID is a paid mutator transaction binding the contract method 0xe97eb362.
//
// Solidity: function updateImageID(bytes32 newImageID) returns()
func (_ZKAdapter *ZKAdapterSession) UpdateImageID(newImageID [32]byte) (*types.Transaction, error) {
	return _ZKAdapter.Contract.UpdateImageID(&_ZKAdapter.TransactOpts, newImageID)
}

// UpdateImageID is a paid mutator transaction binding the contract method 0xe97eb362.
//
// Solidity: function updateImageID(bytes32 newImageID) returns()
func (_ZKAdapter *ZKAdapterTransactorSession) UpdateImageID(newImageID [32]byte) (*types.Transaction, error) {
	return _ZKAdapter.Contract.UpdateImageID(&_ZKAdapter.TransactOpts, newImageID)
}

// UpdatePermissibleTimespan is a paid mutator transaction binding the contract method 0x547923a3.
//
// Solidity: function updatePermissibleTimespan(uint24 newPermissibleTimespan) returns()
func (_ZKAdapter *ZKAdapterTransactor) UpdatePermissibleTimespan(opts *bind.TransactOpts, newPermissibleTimespan *big.Int) (*types.Transaction, error) {
	return _ZKAdapter.contract.Transact(opts, "updatePermissibleTimespan", newPermissibleTimespan)
}

// UpdatePermissibleTimespan is a paid mutator transaction binding the contract method 0x547923a3.
//
// Solidity: function updatePermissibleTimespan(uint24 newPermissibleTimespan) returns()
func (_ZKAdapter *ZKAdapterSession) UpdatePermissibleTimespan(newPermissibleTimespan *big.Int) (*types.Transaction, error) {
	return _ZKAdapter.Contract.UpdatePermissibleTimespan(&_ZKAdapter.TransactOpts, newPermissibleTimespan)
}

// UpdatePermissibleTimespan is a paid mutator transaction binding the contract method 0x547923a3.
//
// Solidity: function updatePermissibleTimespan(uint24 newPermissibleTimespan) returns()
func (_ZKAdapter *ZKAdapterTransactorSession) UpdatePermissibleTimespan(newPermissibleTimespan *big.Int) (*types.Transaction, error) {
	return _ZKAdapter.Contract.UpdatePermissibleTimespan(&_ZKAdapter.TransactOpts, newPermissibleTimespan)
}

// UpdateVerifier is a paid mutator transaction binding the contract method 0x97fc007c.
//
// Solidity: function updateVerifier(address newVerifier) returns()
func (_ZKAdapter *ZKAdapterTransactor) UpdateVerifier(opts *bind.TransactOpts, newVerifier common.Address) (*types.Transaction, error) {
	return _ZKAdapter.contract.Transact(opts, "updateVerifier", newVerifier)
}

// UpdateVerifier is a paid mutator transaction binding the contract method 0x97fc007c.
//
// Solidity: function updateVerifier(address newVerifier) returns()
func (_ZKAdapter *ZKAdapterSession) UpdateVerifier(newVerifier common.Address) (*types.Transaction, error) {
	return _ZKAdapter.Contract.UpdateVerifier(&_ZKAdapter.TransactOpts, newVerifier)
}

// UpdateVerifier is a paid mutator transaction binding the contract method 0x97fc007c.
//
// Solidity: function updateVerifier(address newVerifier) returns()
func (_ZKAdapter *ZKAdapterTransactorSession) UpdateVerifier(newVerifier common.Address) (*types.Transaction, error) {
	return _ZKAdapter.Contract.UpdateVerifier(&_ZKAdapter.TransactOpts, newVerifier)
}

// VerifyLogAndExtract is a paid mutator transaction binding the contract method 0x98ad485d.
//
// Solidity: function verifyLogAndExtract((uint64,uint64,bytes32,bytes32[],bytes32,bytes32[],bytes32,bytes32[],bytes32,bytes32[],bytes) execProof, (bytes[],bytes,bytes,uint256,address,bytes32) logProof) returns(bytes)
func (_ZKAdapter *ZKAdapterTransactor) VerifyLogAndExtract(opts *bind.TransactOpts, execProof ExecutionProof, logProof ReceiptProof) (*types.Transaction, error) {
	return _ZKAdapter.contract.Transact(opts, "verifyLogAndExtract", execProof, logProof)
}

// VerifyLogAndExtract is a paid mutator transaction binding the contract method 0x98ad485d.
//
// Solidity: function verifyLogAndExtract((uint64,uint64,bytes32,bytes32[],bytes32,bytes32[],bytes32,bytes32[],bytes32,bytes32[],bytes) execProof, (bytes[],bytes,bytes,uint256,address,bytes32) logProof) returns(bytes)
func (_ZKAdapter *ZKAdapterSession) VerifyLogAndExtract(execProof ExecutionProof, logProof ReceiptProof) (*types.Transaction, error) {
	return _ZKAdapter.Contract.VerifyLogAndExtract(&_ZKAdapter.TransactOpts, execProof, logProof)
}

// VerifyLogAndExtract is a paid mutator transaction binding the contract method 0x98ad485d.
//
// Solidity: function verifyLogAndExtract((uint64,uint64,bytes32,bytes32[],bytes32,bytes32[],bytes32,bytes32[],bytes32,bytes32[],bytes) execProof, (bytes[],bytes,bytes,uint256,address,bytes32) logProof) returns(bytes)
func (_ZKAdapter *ZKAdapterTransactorSession) VerifyLogAndExtract(execProof ExecutionProof, logProof ReceiptProof) (*types.Transaction, error) {
	return _ZKAdapter.Contract.VerifyLogAndExtract(&_ZKAdapter.TransactOpts, execProof, logProof)
}

// VerifyMessage is a paid mutator transaction binding the contract method 0xf1faff00.
//
// Solidity: function verifyMessage(((uint256,address,address,bytes32,address,uint256,address[],(uint256,address)[],bytes),uint32,bytes32,bytes) message) returns(bool)
func (_ZKAdapter *ZKAdapterTransactor) VerifyMessage(opts *bind.TransactOpts, message TeleporterICMMessage) (*types.Transaction, error) {
	return _ZKAdapter.contract.Transact(opts, "verifyMessage", message)
}

// VerifyMessage is a paid mutator transaction binding the contract method 0xf1faff00.
//
// Solidity: function verifyMessage(((uint256,address,address,bytes32,address,uint256,address[],(uint256,address)[],bytes),uint32,bytes32,bytes) message) returns(bool)
func (_ZKAdapter *ZKAdapterSession) VerifyMessage(message TeleporterICMMessage) (*types.Transaction, error) {
	return _ZKAdapter.Contract.VerifyMessage(&_ZKAdapter.TransactOpts, message)
}

// VerifyMessage is a paid mutator transaction binding the contract method 0xf1faff00.
//
// Solidity: function verifyMessage(((uint256,address,address,bytes32,address,uint256,address[],(uint256,address)[],bytes),uint32,bytes32,bytes) message) returns(bool)
func (_ZKAdapter *ZKAdapterTransactorSession) VerifyMessage(message TeleporterICMMessage) (*types.Transaction, error) {
	return _ZKAdapter.Contract.VerifyMessage(&_ZKAdapter.TransactOpts, message)
}

// ZKAdapterConfirmedBeaconBlockIterator is returned from FilterConfirmedBeaconBlock and is used to iterate over the raw logs and unpacked data for ConfirmedBeaconBlock events raised by the ZKAdapter contract.
type ZKAdapterConfirmedBeaconBlockIterator struct {
	Event *ZKAdapterConfirmedBeaconBlock // Event containing the contract specifics and raw log

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
func (it *ZKAdapterConfirmedBeaconBlockIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ZKAdapterConfirmedBeaconBlock)
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
		it.Event = new(ZKAdapterConfirmedBeaconBlock)
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
func (it *ZKAdapterConfirmedBeaconBlockIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ZKAdapterConfirmedBeaconBlockIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ZKAdapterConfirmedBeaconBlock represents a ConfirmedBeaconBlock event raised by the ZKAdapter contract.
type ZKAdapterConfirmedBeaconBlock struct {
	Slot uint64
	Root [32]byte
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterConfirmedBeaconBlock is a free log retrieval operation binding the contract event 0x9f2417e1c378e9eae44d76184be142d77ed574882fbc98272fe0533053d8816d.
//
// Solidity: event ConfirmedBeaconBlock(uint64 indexed slot, bytes32 indexed root)
func (_ZKAdapter *ZKAdapterFilterer) FilterConfirmedBeaconBlock(opts *bind.FilterOpts, slot []uint64, root [][32]byte) (*ZKAdapterConfirmedBeaconBlockIterator, error) {

	var slotRule []interface{}
	for _, slotItem := range slot {
		slotRule = append(slotRule, slotItem)
	}
	var rootRule []interface{}
	for _, rootItem := range root {
		rootRule = append(rootRule, rootItem)
	}

	logs, sub, err := _ZKAdapter.contract.FilterLogs(opts, "ConfirmedBeaconBlock", slotRule, rootRule)
	if err != nil {
		return nil, err
	}
	return &ZKAdapterConfirmedBeaconBlockIterator{contract: _ZKAdapter.contract, event: "ConfirmedBeaconBlock", logs: logs, sub: sub}, nil
}

// WatchConfirmedBeaconBlock is a free log subscription operation binding the contract event 0x9f2417e1c378e9eae44d76184be142d77ed574882fbc98272fe0533053d8816d.
//
// Solidity: event ConfirmedBeaconBlock(uint64 indexed slot, bytes32 indexed root)
func (_ZKAdapter *ZKAdapterFilterer) WatchConfirmedBeaconBlock(opts *bind.WatchOpts, sink chan<- *ZKAdapterConfirmedBeaconBlock, slot []uint64, root [][32]byte) (event.Subscription, error) {

	var slotRule []interface{}
	for _, slotItem := range slot {
		slotRule = append(slotRule, slotItem)
	}
	var rootRule []interface{}
	for _, rootItem := range root {
		rootRule = append(rootRule, rootItem)
	}

	logs, sub, err := _ZKAdapter.contract.WatchLogs(opts, "ConfirmedBeaconBlock", slotRule, rootRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ZKAdapterConfirmedBeaconBlock)
				if err := _ZKAdapter.contract.UnpackLog(event, "ConfirmedBeaconBlock", log); err != nil {
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

// ParseConfirmedBeaconBlock is a log parse operation binding the contract event 0x9f2417e1c378e9eae44d76184be142d77ed574882fbc98272fe0533053d8816d.
//
// Solidity: event ConfirmedBeaconBlock(uint64 indexed slot, bytes32 indexed root)
func (_ZKAdapter *ZKAdapterFilterer) ParseConfirmedBeaconBlock(log types.Log) (*ZKAdapterConfirmedBeaconBlock, error) {
	event := new(ZKAdapterConfirmedBeaconBlock)
	if err := _ZKAdapter.contract.UnpackLog(event, "ConfirmedBeaconBlock", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ZKAdapterImageIDUpdatedIterator is returned from FilterImageIDUpdated and is used to iterate over the raw logs and unpacked data for ImageIDUpdated events raised by the ZKAdapter contract.
type ZKAdapterImageIDUpdatedIterator struct {
	Event *ZKAdapterImageIDUpdated // Event containing the contract specifics and raw log

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
func (it *ZKAdapterImageIDUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ZKAdapterImageIDUpdated)
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
		it.Event = new(ZKAdapterImageIDUpdated)
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
func (it *ZKAdapterImageIDUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ZKAdapterImageIDUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ZKAdapterImageIDUpdated represents a ImageIDUpdated event raised by the ZKAdapter contract.
type ZKAdapterImageIDUpdated struct {
	NewImageID [32]byte
	OldImageID [32]byte
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterImageIDUpdated is a free log retrieval operation binding the contract event 0xe0282e75f3c31dcf030de54b09dd5050b91ba56b46be09d5d2a069cb37e10730.
//
// Solidity: event ImageIDUpdated(bytes32 indexed newImageID, bytes32 indexed oldImageID)
func (_ZKAdapter *ZKAdapterFilterer) FilterImageIDUpdated(opts *bind.FilterOpts, newImageID [][32]byte, oldImageID [][32]byte) (*ZKAdapterImageIDUpdatedIterator, error) {

	var newImageIDRule []interface{}
	for _, newImageIDItem := range newImageID {
		newImageIDRule = append(newImageIDRule, newImageIDItem)
	}
	var oldImageIDRule []interface{}
	for _, oldImageIDItem := range oldImageID {
		oldImageIDRule = append(oldImageIDRule, oldImageIDItem)
	}

	logs, sub, err := _ZKAdapter.contract.FilterLogs(opts, "ImageIDUpdated", newImageIDRule, oldImageIDRule)
	if err != nil {
		return nil, err
	}
	return &ZKAdapterImageIDUpdatedIterator{contract: _ZKAdapter.contract, event: "ImageIDUpdated", logs: logs, sub: sub}, nil
}

// WatchImageIDUpdated is a free log subscription operation binding the contract event 0xe0282e75f3c31dcf030de54b09dd5050b91ba56b46be09d5d2a069cb37e10730.
//
// Solidity: event ImageIDUpdated(bytes32 indexed newImageID, bytes32 indexed oldImageID)
func (_ZKAdapter *ZKAdapterFilterer) WatchImageIDUpdated(opts *bind.WatchOpts, sink chan<- *ZKAdapterImageIDUpdated, newImageID [][32]byte, oldImageID [][32]byte) (event.Subscription, error) {

	var newImageIDRule []interface{}
	for _, newImageIDItem := range newImageID {
		newImageIDRule = append(newImageIDRule, newImageIDItem)
	}
	var oldImageIDRule []interface{}
	for _, oldImageIDItem := range oldImageID {
		oldImageIDRule = append(oldImageIDRule, oldImageIDItem)
	}

	logs, sub, err := _ZKAdapter.contract.WatchLogs(opts, "ImageIDUpdated", newImageIDRule, oldImageIDRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ZKAdapterImageIDUpdated)
				if err := _ZKAdapter.contract.UnpackLog(event, "ImageIDUpdated", log); err != nil {
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

// ParseImageIDUpdated is a log parse operation binding the contract event 0xe0282e75f3c31dcf030de54b09dd5050b91ba56b46be09d5d2a069cb37e10730.
//
// Solidity: event ImageIDUpdated(bytes32 indexed newImageID, bytes32 indexed oldImageID)
func (_ZKAdapter *ZKAdapterFilterer) ParseImageIDUpdated(log types.Log) (*ZKAdapterImageIDUpdated, error) {
	event := new(ZKAdapterImageIDUpdated)
	if err := _ZKAdapter.contract.UnpackLog(event, "ImageIDUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ZKAdapterPermissibleTimespanUpdatedIterator is returned from FilterPermissibleTimespanUpdated and is used to iterate over the raw logs and unpacked data for PermissibleTimespanUpdated events raised by the ZKAdapter contract.
type ZKAdapterPermissibleTimespanUpdatedIterator struct {
	Event *ZKAdapterPermissibleTimespanUpdated // Event containing the contract specifics and raw log

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
func (it *ZKAdapterPermissibleTimespanUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ZKAdapterPermissibleTimespanUpdated)
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
		it.Event = new(ZKAdapterPermissibleTimespanUpdated)
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
func (it *ZKAdapterPermissibleTimespanUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ZKAdapterPermissibleTimespanUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ZKAdapterPermissibleTimespanUpdated represents a PermissibleTimespanUpdated event raised by the ZKAdapter contract.
type ZKAdapterPermissibleTimespanUpdated struct {
	PermissibleTimespan *big.Int
	Raw                 types.Log // Blockchain specific contextual infos
}

// FilterPermissibleTimespanUpdated is a free log retrieval operation binding the contract event 0x838e79d53629a47c909a6df3129cfd79736efcace61e839a00556cc3252cb174.
//
// Solidity: event PermissibleTimespanUpdated(uint24 indexed permissibleTimespan)
func (_ZKAdapter *ZKAdapterFilterer) FilterPermissibleTimespanUpdated(opts *bind.FilterOpts, permissibleTimespan []*big.Int) (*ZKAdapterPermissibleTimespanUpdatedIterator, error) {

	var permissibleTimespanRule []interface{}
	for _, permissibleTimespanItem := range permissibleTimespan {
		permissibleTimespanRule = append(permissibleTimespanRule, permissibleTimespanItem)
	}

	logs, sub, err := _ZKAdapter.contract.FilterLogs(opts, "PermissibleTimespanUpdated", permissibleTimespanRule)
	if err != nil {
		return nil, err
	}
	return &ZKAdapterPermissibleTimespanUpdatedIterator{contract: _ZKAdapter.contract, event: "PermissibleTimespanUpdated", logs: logs, sub: sub}, nil
}

// WatchPermissibleTimespanUpdated is a free log subscription operation binding the contract event 0x838e79d53629a47c909a6df3129cfd79736efcace61e839a00556cc3252cb174.
//
// Solidity: event PermissibleTimespanUpdated(uint24 indexed permissibleTimespan)
func (_ZKAdapter *ZKAdapterFilterer) WatchPermissibleTimespanUpdated(opts *bind.WatchOpts, sink chan<- *ZKAdapterPermissibleTimespanUpdated, permissibleTimespan []*big.Int) (event.Subscription, error) {

	var permissibleTimespanRule []interface{}
	for _, permissibleTimespanItem := range permissibleTimespan {
		permissibleTimespanRule = append(permissibleTimespanRule, permissibleTimespanItem)
	}

	logs, sub, err := _ZKAdapter.contract.WatchLogs(opts, "PermissibleTimespanUpdated", permissibleTimespanRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ZKAdapterPermissibleTimespanUpdated)
				if err := _ZKAdapter.contract.UnpackLog(event, "PermissibleTimespanUpdated", log); err != nil {
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

// ParsePermissibleTimespanUpdated is a log parse operation binding the contract event 0x838e79d53629a47c909a6df3129cfd79736efcace61e839a00556cc3252cb174.
//
// Solidity: event PermissibleTimespanUpdated(uint24 indexed permissibleTimespan)
func (_ZKAdapter *ZKAdapterFilterer) ParsePermissibleTimespanUpdated(log types.Log) (*ZKAdapterPermissibleTimespanUpdated, error) {
	event := new(ZKAdapterPermissibleTimespanUpdated)
	if err := _ZKAdapter.contract.UnpackLog(event, "PermissibleTimespanUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ZKAdapterRoleAdminChangedIterator is returned from FilterRoleAdminChanged and is used to iterate over the raw logs and unpacked data for RoleAdminChanged events raised by the ZKAdapter contract.
type ZKAdapterRoleAdminChangedIterator struct {
	Event *ZKAdapterRoleAdminChanged // Event containing the contract specifics and raw log

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
func (it *ZKAdapterRoleAdminChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ZKAdapterRoleAdminChanged)
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
		it.Event = new(ZKAdapterRoleAdminChanged)
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
func (it *ZKAdapterRoleAdminChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ZKAdapterRoleAdminChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ZKAdapterRoleAdminChanged represents a RoleAdminChanged event raised by the ZKAdapter contract.
type ZKAdapterRoleAdminChanged struct {
	Role              [32]byte
	PreviousAdminRole [32]byte
	NewAdminRole      [32]byte
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterRoleAdminChanged is a free log retrieval operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_ZKAdapter *ZKAdapterFilterer) FilterRoleAdminChanged(opts *bind.FilterOpts, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (*ZKAdapterRoleAdminChangedIterator, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var previousAdminRoleRule []interface{}
	for _, previousAdminRoleItem := range previousAdminRole {
		previousAdminRoleRule = append(previousAdminRoleRule, previousAdminRoleItem)
	}
	var newAdminRoleRule []interface{}
	for _, newAdminRoleItem := range newAdminRole {
		newAdminRoleRule = append(newAdminRoleRule, newAdminRoleItem)
	}

	logs, sub, err := _ZKAdapter.contract.FilterLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return &ZKAdapterRoleAdminChangedIterator{contract: _ZKAdapter.contract, event: "RoleAdminChanged", logs: logs, sub: sub}, nil
}

// WatchRoleAdminChanged is a free log subscription operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_ZKAdapter *ZKAdapterFilterer) WatchRoleAdminChanged(opts *bind.WatchOpts, sink chan<- *ZKAdapterRoleAdminChanged, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (event.Subscription, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var previousAdminRoleRule []interface{}
	for _, previousAdminRoleItem := range previousAdminRole {
		previousAdminRoleRule = append(previousAdminRoleRule, previousAdminRoleItem)
	}
	var newAdminRoleRule []interface{}
	for _, newAdminRoleItem := range newAdminRole {
		newAdminRoleRule = append(newAdminRoleRule, newAdminRoleItem)
	}

	logs, sub, err := _ZKAdapter.contract.WatchLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ZKAdapterRoleAdminChanged)
				if err := _ZKAdapter.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
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

// ParseRoleAdminChanged is a log parse operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_ZKAdapter *ZKAdapterFilterer) ParseRoleAdminChanged(log types.Log) (*ZKAdapterRoleAdminChanged, error) {
	event := new(ZKAdapterRoleAdminChanged)
	if err := _ZKAdapter.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ZKAdapterRoleGrantedIterator is returned from FilterRoleGranted and is used to iterate over the raw logs and unpacked data for RoleGranted events raised by the ZKAdapter contract.
type ZKAdapterRoleGrantedIterator struct {
	Event *ZKAdapterRoleGranted // Event containing the contract specifics and raw log

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
func (it *ZKAdapterRoleGrantedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ZKAdapterRoleGranted)
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
		it.Event = new(ZKAdapterRoleGranted)
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
func (it *ZKAdapterRoleGrantedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ZKAdapterRoleGrantedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ZKAdapterRoleGranted represents a RoleGranted event raised by the ZKAdapter contract.
type ZKAdapterRoleGranted struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleGranted is a free log retrieval operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_ZKAdapter *ZKAdapterFilterer) FilterRoleGranted(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*ZKAdapterRoleGrantedIterator, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _ZKAdapter.contract.FilterLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &ZKAdapterRoleGrantedIterator{contract: _ZKAdapter.contract, event: "RoleGranted", logs: logs, sub: sub}, nil
}

// WatchRoleGranted is a free log subscription operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_ZKAdapter *ZKAdapterFilterer) WatchRoleGranted(opts *bind.WatchOpts, sink chan<- *ZKAdapterRoleGranted, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _ZKAdapter.contract.WatchLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ZKAdapterRoleGranted)
				if err := _ZKAdapter.contract.UnpackLog(event, "RoleGranted", log); err != nil {
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

// ParseRoleGranted is a log parse operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_ZKAdapter *ZKAdapterFilterer) ParseRoleGranted(log types.Log) (*ZKAdapterRoleGranted, error) {
	event := new(ZKAdapterRoleGranted)
	if err := _ZKAdapter.contract.UnpackLog(event, "RoleGranted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ZKAdapterRoleRevokedIterator is returned from FilterRoleRevoked and is used to iterate over the raw logs and unpacked data for RoleRevoked events raised by the ZKAdapter contract.
type ZKAdapterRoleRevokedIterator struct {
	Event *ZKAdapterRoleRevoked // Event containing the contract specifics and raw log

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
func (it *ZKAdapterRoleRevokedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ZKAdapterRoleRevoked)
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
		it.Event = new(ZKAdapterRoleRevoked)
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
func (it *ZKAdapterRoleRevokedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ZKAdapterRoleRevokedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ZKAdapterRoleRevoked represents a RoleRevoked event raised by the ZKAdapter contract.
type ZKAdapterRoleRevoked struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleRevoked is a free log retrieval operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_ZKAdapter *ZKAdapterFilterer) FilterRoleRevoked(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*ZKAdapterRoleRevokedIterator, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _ZKAdapter.contract.FilterLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &ZKAdapterRoleRevokedIterator{contract: _ZKAdapter.contract, event: "RoleRevoked", logs: logs, sub: sub}, nil
}

// WatchRoleRevoked is a free log subscription operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_ZKAdapter *ZKAdapterFilterer) WatchRoleRevoked(opts *bind.WatchOpts, sink chan<- *ZKAdapterRoleRevoked, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _ZKAdapter.contract.WatchLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ZKAdapterRoleRevoked)
				if err := _ZKAdapter.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
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

// ParseRoleRevoked is a log parse operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_ZKAdapter *ZKAdapterFilterer) ParseRoleRevoked(log types.Log) (*ZKAdapterRoleRevoked, error) {
	event := new(ZKAdapterRoleRevoked)
	if err := _ZKAdapter.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ZKAdapterTeleporterV2MessageSentIterator is returned from FilterTeleporterV2MessageSent and is used to iterate over the raw logs and unpacked data for TeleporterV2MessageSent events raised by the ZKAdapter contract.
type ZKAdapterTeleporterV2MessageSentIterator struct {
	Event *ZKAdapterTeleporterV2MessageSent // Event containing the contract specifics and raw log

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
func (it *ZKAdapterTeleporterV2MessageSentIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ZKAdapterTeleporterV2MessageSent)
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
		it.Event = new(ZKAdapterTeleporterV2MessageSent)
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
func (it *ZKAdapterTeleporterV2MessageSentIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ZKAdapterTeleporterV2MessageSentIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ZKAdapterTeleporterV2MessageSent represents a TeleporterV2MessageSent event raised by the ZKAdapter contract.
type ZKAdapterTeleporterV2MessageSent struct {
	EncodedMessage []byte
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterTeleporterV2MessageSent is a free log retrieval operation binding the contract event 0x467f98200caa1453d527fd7de92faa829d8357fb650c10010e77188d93a3a3b6.
//
// Solidity: event TeleporterV2MessageSent(bytes encodedMessage)
func (_ZKAdapter *ZKAdapterFilterer) FilterTeleporterV2MessageSent(opts *bind.FilterOpts) (*ZKAdapterTeleporterV2MessageSentIterator, error) {

	logs, sub, err := _ZKAdapter.contract.FilterLogs(opts, "TeleporterV2MessageSent")
	if err != nil {
		return nil, err
	}
	return &ZKAdapterTeleporterV2MessageSentIterator{contract: _ZKAdapter.contract, event: "TeleporterV2MessageSent", logs: logs, sub: sub}, nil
}

// WatchTeleporterV2MessageSent is a free log subscription operation binding the contract event 0x467f98200caa1453d527fd7de92faa829d8357fb650c10010e77188d93a3a3b6.
//
// Solidity: event TeleporterV2MessageSent(bytes encodedMessage)
func (_ZKAdapter *ZKAdapterFilterer) WatchTeleporterV2MessageSent(opts *bind.WatchOpts, sink chan<- *ZKAdapterTeleporterV2MessageSent) (event.Subscription, error) {

	logs, sub, err := _ZKAdapter.contract.WatchLogs(opts, "TeleporterV2MessageSent")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ZKAdapterTeleporterV2MessageSent)
				if err := _ZKAdapter.contract.UnpackLog(event, "TeleporterV2MessageSent", log); err != nil {
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

// ParseTeleporterV2MessageSent is a log parse operation binding the contract event 0x467f98200caa1453d527fd7de92faa829d8357fb650c10010e77188d93a3a3b6.
//
// Solidity: event TeleporterV2MessageSent(bytes encodedMessage)
func (_ZKAdapter *ZKAdapterFilterer) ParseTeleporterV2MessageSent(log types.Log) (*ZKAdapterTeleporterV2MessageSent, error) {
	event := new(ZKAdapterTeleporterV2MessageSent)
	if err := _ZKAdapter.contract.UnpackLog(event, "TeleporterV2MessageSent", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ZKAdapterTransitionedIterator is returned from FilterTransitioned and is used to iterate over the raw logs and unpacked data for Transitioned events raised by the ZKAdapter contract.
type ZKAdapterTransitionedIterator struct {
	Event *ZKAdapterTransitioned // Event containing the contract specifics and raw log

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
func (it *ZKAdapterTransitionedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ZKAdapterTransitioned)
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
		it.Event = new(ZKAdapterTransitioned)
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
func (it *ZKAdapterTransitionedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ZKAdapterTransitionedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ZKAdapterTransitioned represents a Transitioned event raised by the ZKAdapter contract.
type ZKAdapterTransitioned struct {
	PreEpoch  uint64
	PostEpoch uint64
	PreState  ConsensusState
	PostState ConsensusState
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterTransitioned is a free log retrieval operation binding the contract event 0x955222ca45c52776cd8084d242443738d71afc33822d853b8efeb94e070d94d0.
//
// Solidity: event Transitioned(uint64 indexed preEpoch, uint64 indexed postEpoch, ((uint64,bytes32),(uint64,bytes32)) preState, ((uint64,bytes32),(uint64,bytes32)) postState)
func (_ZKAdapter *ZKAdapterFilterer) FilterTransitioned(opts *bind.FilterOpts, preEpoch []uint64, postEpoch []uint64) (*ZKAdapterTransitionedIterator, error) {

	var preEpochRule []interface{}
	for _, preEpochItem := range preEpoch {
		preEpochRule = append(preEpochRule, preEpochItem)
	}
	var postEpochRule []interface{}
	for _, postEpochItem := range postEpoch {
		postEpochRule = append(postEpochRule, postEpochItem)
	}

	logs, sub, err := _ZKAdapter.contract.FilterLogs(opts, "Transitioned", preEpochRule, postEpochRule)
	if err != nil {
		return nil, err
	}
	return &ZKAdapterTransitionedIterator{contract: _ZKAdapter.contract, event: "Transitioned", logs: logs, sub: sub}, nil
}

// WatchTransitioned is a free log subscription operation binding the contract event 0x955222ca45c52776cd8084d242443738d71afc33822d853b8efeb94e070d94d0.
//
// Solidity: event Transitioned(uint64 indexed preEpoch, uint64 indexed postEpoch, ((uint64,bytes32),(uint64,bytes32)) preState, ((uint64,bytes32),(uint64,bytes32)) postState)
func (_ZKAdapter *ZKAdapterFilterer) WatchTransitioned(opts *bind.WatchOpts, sink chan<- *ZKAdapterTransitioned, preEpoch []uint64, postEpoch []uint64) (event.Subscription, error) {

	var preEpochRule []interface{}
	for _, preEpochItem := range preEpoch {
		preEpochRule = append(preEpochRule, preEpochItem)
	}
	var postEpochRule []interface{}
	for _, postEpochItem := range postEpoch {
		postEpochRule = append(postEpochRule, postEpochItem)
	}

	logs, sub, err := _ZKAdapter.contract.WatchLogs(opts, "Transitioned", preEpochRule, postEpochRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ZKAdapterTransitioned)
				if err := _ZKAdapter.contract.UnpackLog(event, "Transitioned", log); err != nil {
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

// ParseTransitioned is a log parse operation binding the contract event 0x955222ca45c52776cd8084d242443738d71afc33822d853b8efeb94e070d94d0.
//
// Solidity: event Transitioned(uint64 indexed preEpoch, uint64 indexed postEpoch, ((uint64,bytes32),(uint64,bytes32)) preState, ((uint64,bytes32),(uint64,bytes32)) postState)
func (_ZKAdapter *ZKAdapterFilterer) ParseTransitioned(log types.Log) (*ZKAdapterTransitioned, error) {
	event := new(ZKAdapterTransitioned)
	if err := _ZKAdapter.contract.UnpackLog(event, "Transitioned", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ZKAdapterZKEventImportedIterator is returned from FilterZKEventImported and is used to iterate over the raw logs and unpacked data for ZKEventImported events raised by the ZKAdapter contract.
type ZKAdapterZKEventImportedIterator struct {
	Event *ZKAdapterZKEventImported // Event containing the contract specifics and raw log

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
func (it *ZKAdapterZKEventImportedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ZKAdapterZKEventImported)
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
		it.Event = new(ZKAdapterZKEventImported)
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
func (it *ZKAdapterZKEventImportedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ZKAdapterZKEventImportedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ZKAdapterZKEventImported represents a ZKEventImported event raised by the ZKAdapter contract.
type ZKAdapterZKEventImported struct {
	SourceChainId *big.Int
	BeaconSlot    *big.Int
	Emitter       common.Address
	ExecutionRoot [32]byte
	LogIndex      *big.Int
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterZKEventImported is a free log retrieval operation binding the contract event 0x17e2291123b98af9b416b45d04d46e4484bbf9d2c1061bbf567425a66bb1beb1.
//
// Solidity: event ZKEventImported(uint256 indexed sourceChainId, uint256 indexed beaconSlot, address indexed emitter, bytes32 executionRoot, uint256 logIndex)
func (_ZKAdapter *ZKAdapterFilterer) FilterZKEventImported(opts *bind.FilterOpts, sourceChainId []*big.Int, beaconSlot []*big.Int, emitter []common.Address) (*ZKAdapterZKEventImportedIterator, error) {

	var sourceChainIdRule []interface{}
	for _, sourceChainIdItem := range sourceChainId {
		sourceChainIdRule = append(sourceChainIdRule, sourceChainIdItem)
	}
	var beaconSlotRule []interface{}
	for _, beaconSlotItem := range beaconSlot {
		beaconSlotRule = append(beaconSlotRule, beaconSlotItem)
	}
	var emitterRule []interface{}
	for _, emitterItem := range emitter {
		emitterRule = append(emitterRule, emitterItem)
	}

	logs, sub, err := _ZKAdapter.contract.FilterLogs(opts, "ZKEventImported", sourceChainIdRule, beaconSlotRule, emitterRule)
	if err != nil {
		return nil, err
	}
	return &ZKAdapterZKEventImportedIterator{contract: _ZKAdapter.contract, event: "ZKEventImported", logs: logs, sub: sub}, nil
}

// WatchZKEventImported is a free log subscription operation binding the contract event 0x17e2291123b98af9b416b45d04d46e4484bbf9d2c1061bbf567425a66bb1beb1.
//
// Solidity: event ZKEventImported(uint256 indexed sourceChainId, uint256 indexed beaconSlot, address indexed emitter, bytes32 executionRoot, uint256 logIndex)
func (_ZKAdapter *ZKAdapterFilterer) WatchZKEventImported(opts *bind.WatchOpts, sink chan<- *ZKAdapterZKEventImported, sourceChainId []*big.Int, beaconSlot []*big.Int, emitter []common.Address) (event.Subscription, error) {

	var sourceChainIdRule []interface{}
	for _, sourceChainIdItem := range sourceChainId {
		sourceChainIdRule = append(sourceChainIdRule, sourceChainIdItem)
	}
	var beaconSlotRule []interface{}
	for _, beaconSlotItem := range beaconSlot {
		beaconSlotRule = append(beaconSlotRule, beaconSlotItem)
	}
	var emitterRule []interface{}
	for _, emitterItem := range emitter {
		emitterRule = append(emitterRule, emitterItem)
	}

	logs, sub, err := _ZKAdapter.contract.WatchLogs(opts, "ZKEventImported", sourceChainIdRule, beaconSlotRule, emitterRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ZKAdapterZKEventImported)
				if err := _ZKAdapter.contract.UnpackLog(event, "ZKEventImported", log); err != nil {
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

// ParseZKEventImported is a log parse operation binding the contract event 0x17e2291123b98af9b416b45d04d46e4484bbf9d2c1061bbf567425a66bb1beb1.
//
// Solidity: event ZKEventImported(uint256 indexed sourceChainId, uint256 indexed beaconSlot, address indexed emitter, bytes32 executionRoot, uint256 logIndex)
func (_ZKAdapter *ZKAdapterFilterer) ParseZKEventImported(log types.Log) (*ZKAdapterZKEventImported, error) {
	event := new(ZKAdapterZKEventImported)
	if err := _ZKAdapter.contract.UnpackLog(event, "ZKEventImported", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
