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
	GIndexBlockStateRoot *big.Int
	GIndexExecRoot       *big.Int
	GIndexBaseStateRoots *big.Int
	StateRootsDepth      *big.Int
	GIndexReceiptsRoot   *big.Int
	StateRootsVectorSize *big.Int
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
	ABI: "[{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"sourceChainId_\",\"type\":\"uint256\"},{\"components\":[{\"components\":[{\"internalType\":\"uint64\",\"name\":\"epoch\",\"type\":\"uint64\"},{\"internalType\":\"bytes32\",\"name\":\"root\",\"type\":\"bytes32\"}],\"internalType\":\"structConsensus.Checkpoint\",\"name\":\"currentJustifiedCheckpoint\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint64\",\"name\":\"epoch\",\"type\":\"uint64\"},{\"internalType\":\"bytes32\",\"name\":\"root\",\"type\":\"bytes32\"}],\"internalType\":\"structConsensus.Checkpoint\",\"name\":\"finalizedCheckpoint\",\"type\":\"tuple\"}],\"internalType\":\"structConsensus.State\",\"name\":\"startingState\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"gIndexBlockStateRoot\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"gIndexExecRoot\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"gIndexBaseStateRoots\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"stateRootsDepth\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"gIndexReceiptsRoot\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"stateRootsVectorSize\",\"type\":\"uint256\"}],\"internalType\":\"structExecution.BeaconConfig\",\"name\":\"beaconConfig_\",\"type\":\"tuple\"},{\"internalType\":\"uint256\",\"name\":\"genesisTime_\",\"type\":\"uint256\"},{\"internalType\":\"uint24\",\"name\":\"permissibleTimespan_\",\"type\":\"uint24\"},{\"internalType\":\"address\",\"name\":\"verifier\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"imageID_\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"admin\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"superAdmin\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"AccessControlBadConfirmation\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"neededRole\",\"type\":\"bytes32\"}],\"name\":\"AccessControlUnauthorizedAccount\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidArgument\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidPreState\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"PermissibleTimespanLapsed\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"chainId\",\"type\":\"uint256\"}],\"name\":\"UnsupportedSourceChain\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint64\",\"name\":\"slot\",\"type\":\"uint64\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"root\",\"type\":\"bytes32\"}],\"name\":\"ConfirmedBeaconBlock\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"newImageID\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"oldImageID\",\"type\":\"bytes32\"}],\"name\":\"ImageIDUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint24\",\"name\":\"permissibleTimespan\",\"type\":\"uint24\"}],\"name\":\"PermissibleTimespanUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"previousAdminRole\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"newAdminRole\",\"type\":\"bytes32\"}],\"name\":\"RoleAdminChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleGranted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleRevoked\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"encodedMessage\",\"type\":\"bytes\"}],\"name\":\"TeleporterV2MessageSent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint64\",\"name\":\"preEpoch\",\"type\":\"uint64\"},{\"indexed\":true,\"internalType\":\"uint64\",\"name\":\"postEpoch\",\"type\":\"uint64\"},{\"components\":[{\"components\":[{\"internalType\":\"uint64\",\"name\":\"epoch\",\"type\":\"uint64\"},{\"internalType\":\"bytes32\",\"name\":\"root\",\"type\":\"bytes32\"}],\"internalType\":\"structConsensus.Checkpoint\",\"name\":\"currentJustifiedCheckpoint\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint64\",\"name\":\"epoch\",\"type\":\"uint64\"},{\"internalType\":\"bytes32\",\"name\":\"root\",\"type\":\"bytes32\"}],\"internalType\":\"structConsensus.Checkpoint\",\"name\":\"finalizedCheckpoint\",\"type\":\"tuple\"}],\"indexed\":false,\"internalType\":\"structConsensus.State\",\"name\":\"preState\",\"type\":\"tuple\"},{\"components\":[{\"components\":[{\"internalType\":\"uint64\",\"name\":\"epoch\",\"type\":\"uint64\"},{\"internalType\":\"bytes32\",\"name\":\"root\",\"type\":\"bytes32\"}],\"internalType\":\"structConsensus.Checkpoint\",\"name\":\"currentJustifiedCheckpoint\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint64\",\"name\":\"epoch\",\"type\":\"uint64\"},{\"internalType\":\"bytes32\",\"name\":\"root\",\"type\":\"bytes32\"}],\"internalType\":\"structConsensus.Checkpoint\",\"name\":\"finalizedCheckpoint\",\"type\":\"tuple\"}],\"indexed\":false,\"internalType\":\"structConsensus.State\",\"name\":\"postState\",\"type\":\"tuple\"}],\"name\":\"Transitioned\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"sourceChainId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"beaconSlot\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"emitter\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"executionRoot\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"logIndex\",\"type\":\"uint256\"}],\"name\":\"ZKEventImported\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"ADMIN_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"DEFAULT_ADMIN_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"SECONDS_PER_SLOT\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"SLOT_PER_EPOCH\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"UNDEFINED_ROOT\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"slot\",\"type\":\"uint64\"}],\"name\":\"beaconRootAt\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"epoch\",\"type\":\"uint64\"}],\"name\":\"epochTimestamp\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"genesisTime\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"slot\",\"type\":\"uint64\"}],\"name\":\"getBeaconBlockRoot\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"root\",\"type\":\"bytes32\"},{\"internalType\":\"bool\",\"name\":\"valid\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"}],\"name\":\"getRoleAdmin\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"grantRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"hasRole\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"imageID\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"journalData\",\"type\":\"bytes\"},{\"internalType\":\"uint64\",\"name\":\"finalizedSlot\",\"type\":\"uint64\"}],\"name\":\"manualTransition\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"permissibleTimespan\",\"outputs\":[{\"internalType\":\"uint24\",\"name\":\"\",\"type\":\"uint24\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"callerConfirmation\",\"type\":\"address\"}],\"name\":\"renounceRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"revokeRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"messageNonce\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"originSenderAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"originTeleporterAddress\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"destinationBlockchainID\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"destinationAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"requiredGasLimit\",\"type\":\"uint256\"},{\"internalType\":\"address[]\",\"name\":\"allowedRelayerAddresses\",\"type\":\"address[]\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"receivedMessageNonce\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"relayerRewardAddress\",\"type\":\"address\"}],\"internalType\":\"structTeleporterMessageReceipt[]\",\"name\":\"receipts\",\"type\":\"tuple[]\"},{\"internalType\":\"bytes\",\"name\":\"message\",\"type\":\"bytes\"}],\"internalType\":\"structTeleporterMessageV2\",\"name\":\"message\",\"type\":\"tuple\"}],\"name\":\"sendMessage\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"sourceChainId\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"sourceNetworkId\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"bytes\",\"name\":\"journalData\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"seal\",\"type\":\"bytes\"},{\"internalType\":\"uint64\",\"name\":\"finalizedSlot\",\"type\":\"uint64\"}],\"internalType\":\"structConsensusData\",\"name\":\"consensus\",\"type\":\"tuple\"}],\"name\":\"transition\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"newImageID\",\"type\":\"bytes32\"}],\"name\":\"updateImageID\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint24\",\"name\":\"newPermissibleTimespan\",\"type\":\"uint24\"}],\"name\":\"updatePermissibleTimespan\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newVerifier\",\"type\":\"address\"}],\"name\":\"updateVerifier\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"verifier\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint64\",\"name\":\"anchorSlot\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"targetSlot\",\"type\":\"uint64\"},{\"internalType\":\"bytes32\",\"name\":\"anchorBeaconStateRoot\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32[]\",\"name\":\"anchorBeaconStateProof\",\"type\":\"bytes32[]\"},{\"internalType\":\"bytes32\",\"name\":\"targetBeaconStateRoot\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32[]\",\"name\":\"targetBeaconStateProof\",\"type\":\"bytes32[]\"},{\"internalType\":\"bytes32\",\"name\":\"targetExecutionHeaderRoot\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32[]\",\"name\":\"targetExecutionHeaderProof\",\"type\":\"bytes32[]\"},{\"internalType\":\"bytes32\",\"name\":\"targetReceiptsRoot\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32[]\",\"name\":\"targetReceiptsProof\",\"type\":\"bytes32[]\"}],\"internalType\":\"structExecution.Proof\",\"name\":\"execProof\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"bytes[]\",\"name\":\"proof\",\"type\":\"bytes[]\"},{\"internalType\":\"bytes\",\"name\":\"key\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"value\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"logIndex\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"expectedEmitter\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"expectedTopic0\",\"type\":\"bytes32\"}],\"internalType\":\"structReceipt.Proof\",\"name\":\"logProof\",\"type\":\"tuple\"}],\"name\":\"verifyLogAndExtract\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"messageNonce\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"originSenderAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"originTeleporterAddress\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"destinationBlockchainID\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"destinationAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"requiredGasLimit\",\"type\":\"uint256\"},{\"internalType\":\"address[]\",\"name\":\"allowedRelayerAddresses\",\"type\":\"address[]\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"receivedMessageNonce\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"relayerRewardAddress\",\"type\":\"address\"}],\"internalType\":\"structTeleporterMessageReceipt[]\",\"name\":\"receipts\",\"type\":\"tuple[]\"},{\"internalType\":\"bytes\",\"name\":\"message\",\"type\":\"bytes\"}],\"internalType\":\"structTeleporterMessageV2\",\"name\":\"message\",\"type\":\"tuple\"},{\"internalType\":\"uint32\",\"name\":\"sourceNetworkID\",\"type\":\"uint32\"},{\"internalType\":\"bytes32\",\"name\":\"sourceBlockchainID\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"attestation\",\"type\":\"bytes\"}],\"internalType\":\"structTeleporterICMMessage\",\"name\":\"message\",\"type\":\"tuple\"}],\"name\":\"verifyMessage\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x60e060405234801561000f575f5ffd5b50604051614cf6380380614cf683398101604081905261002e91610488565b8888888888888888888887815f036100805760405162461bcd60e51b815260206004820152601060248201526f125b9d985b1a590818da185a5b88125160821b60448201526064015b60405180910390fd5b608091825280516001556020810151600255604081015160035560608101516004559081015160055560a001516006555f8690036101005760405162461bcd60e51b815260206004820152601460248201527f496e76616c69642067656e657369732074696d650000000000000000000000006044820152606401610077565b60a086905261012f7fa49807205ce4d355092ef5a8a18f56e8913cf4a201fbe287825b095693c217758361024c565b5061013a5f8261024c565b50602088015151429061014c906102f5565b111561019a5760405162461bcd60e51b815260206004820152601c60248201527f5374617274696e6720737461746520696e2074686520667574757265000000006044820152606401610077565b505085518051600780546001600160401b03199081166001600160401b039384161790915560209283015160085597820151805160098054909a169216919091179097559590950151600a55600d80546001600160b81b031916600160a01b62ffffff94909416939093026001600160a01b031916929092176001600160a01b03919091161790555050600c555060805161023490610323565b63ffffffff1660c0525061058b975050505050505050565b5f828152602081815260408083206001600160a01b038516845290915281205460ff166102ec575f838152602081815260408083206001600160a01b03861684529091529020805460ff191660011790556102a43390565b6001600160a01b0316826001600160a01b0316847f2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d60405160405180910390a45060016102ef565b505f5b92915050565b5f600c61030c60206001600160401b038516610561565b6103169190610561565b60a0516102ef9190610578565b5f8160010361033457506001919050565b8162aa36a70361034657506002919050565b6040516372c559db60e01b815260048101839052602401610077565b604080519081016001600160401b038111828210171561039057634e487b7160e01b5f52604160045260245ffd5b60405290565b5f604082840312156103a6575f5ffd5b6103ae610362565b82519091506001600160401b03811681146103c7575f5ffd5b815260209182015191810191909152919050565b5f60c082840312156103eb575f5ffd5b60405160c081016001600160401b038111828210171561041957634e487b7160e01b5f52604160045260245ffd5b604090815283518252602080850151908301528381015190820152606080840151908201526080808401519082015260a0928301519281019290925250919050565b805162ffffff8116811461046d575f5ffd5b919050565b80516001600160a01b038116811461046d575f5ffd5b5f5f5f5f5f5f5f5f5f898b036102208112156104a2575f5ffd5b8a5199506080601f19820112156104b7575f5ffd5b506104c0610362565b6104cd8c60208d01610396565b81526104dc8c60608d01610396565b602082015297506104f08b60a08c016103db565b6101608b015190975095506105086101808b0161045b565b94506105176101a08b01610472565b6101c08b0151909450925061052f6101e08b01610472565b915061053e6102008b01610472565b90509295985092959850929598565b634e487b7160e01b5f52601160045260245ffd5b80820281158282048414176102ef576102ef61054d565b808201808211156102ef576102ef61054d565b60805160a05160c0516147256105d15f395f81816103450152610ae201525f81816102e0015261055b01525f8181610216015281816109e60152610a7c01526147255ff3fe608060405234801561000f575f5ffd5b50600436106101bb575f3560e01c80636fda4c8b116100f3578063d085070811610093578063e97eb3621161006e578063e97eb36214610443578063eb97cd2c14610456578063ebeced0e14610469578063f1faff0014610489575f5ffd5b8063d0850708146103dc578063d547741f1461041d578063db2635b714610430575f5ffd5b806391d14854116100ce57806391d14854146103a357806397fc007c146103b6578063a217fddf1461031e578063c506e0be146103c9575f5ffd5b80636fda4c8b1461034057806375b238fc1461037c5780638e1e49ba14610390575f5ffd5b806336568abe1161015e5780634d3499b5116101395780634d3499b51461031557806354747fd41461031e578063547923a31461032557806365242a0e14610338575f5ffd5b806336568abe146102c857806342c6498a146102db5780634b0d0fa814610302575f5ffd5b8063248a9ca311610199578063248a9ca3146102465780632b7ac3f3146102685780632f2ff15d14610293578063304b9071146102a8575f5ffd5b806301ffc9a7146101bf57806304ae4327146101e75780631544298e14610211575b5f5ffd5b6101d26101cd3660046135fe565b61049c565b60405190151581526020015b60405180910390f35b600d546101fd90600160a01b900462ffffff1681565b60405162ffffff90911681526020016101de565b6102387f000000000000000000000000000000000000000000000000000000000000000081565b6040519081526020016101de565b610238610254366004613625565b5f9081526020819052604090206001015490565b600d5461027b906001600160a01b031681565b6040516001600160a01b0390911681526020016101de565b6102a66102a1366004613660565b6104d2565b005b6102b0600c81565b6040516001600160401b0390911681526020016101de565b6102a66102d6366004613660565b6104fc565b6102387f000000000000000000000000000000000000000000000000000000000000000081565b6102386103103660046136a4565b610534565b610238600c5481565b6102385f81565b6102a66103333660046136bd565b61057f565b6102b0602081565b6103677f000000000000000000000000000000000000000000000000000000000000000081565b60405163ffffffff90911681526020016101de565b6102385f5160206146d05f395f51905f5281565b61023861039e3660046136a4565b6105ed565b6101d26103b1366004613660565b61060a565b6102a66103c43660046136df565b610632565b6102a66103d73660046136fa565b6106f6565b6104086103ea3660046136a4565b6001600160401b03165f908152600b60205260409020549081151590565b604080519283529015156020830152016101de565b6102a661042b366004613660565b61072d565b6102a661043e366004613774565b610751565b6102a6610451366004613625565b610792565b6102a66104643660046137aa565b6107d1565b61047c6104773660046137e1565b610918565b6040516101de9190613891565b6101d26104973660046138a3565b610a57565b5f6001600160e01b03198216637965db0b60e01b14806104cc57506301ffc9a760e01b6001600160e01b03198316145b92915050565b5f828152602081905260409020600101546104ec81610d0e565b6104f68383610d1b565b50505050565b6001600160a01b03811633146105255760405163334bd91960e11b815260040160405180910390fd5b61052f8282610daa565b505050565b5f600c61054b60206001600160401b0385166138ed565b61055591906138ed565b6104cc907f0000000000000000000000000000000000000000000000000000000000000000613904565b5f5160206146d05f395f51905f5261059681610d0e565b600d5462ffffff600160a01b9091048116908316036105c85760405163a9cb9e0d60e01b815260040160405180910390fd5b50600d805462ffffff909216600160a01b0262ffffff60a01b19909216919091179055565b6001600160401b0381165f908152600b60205260408120546104cc565b5f918252602082815260408084206001600160a01b0393909316845291905290205460ff1690565b5f5160206146d05f395f51905f5261064981610d0e565b6001600160a01b0382166106a45760405162461bcd60e51b815260206004820152601860248201527f496e76616c69642076657269666965722061646472657373000000000000000060448201526064015b60405180910390fd5b600d546001600160a01b03908116908316036106d35760405163a9cb9e0d60e01b815260040160405180910390fd5b50600d80546001600160a01b0319166001600160a01b0392909216919091179055565b5f5160206146d05f395f51905f5261070d81610d0e565b5f61071a84860186613a3d565b90506107268184610e13565b5050505050565b5f8281526020819052604090206001015461074781610d0e565b6104f68383610daa565b5f61075c8280613aac565b8101906107699190613a3d565b90506107758183610eba565b61078e8161078960608501604086016136a4565b610e13565b5050565b5f5160206146d05f395f51905f526107a981610d0e565b600c5482036107cb5760405163a9cb9e0d60e01b815260040160405180910390fd5b50600c55565b6107e160608201604083016136df565b6001600160a01b0316336001600160a01b0316148061087f575061080b60608201604083016136df565b6001600160a01b031663d67bdd256040518163ffffffff1660e01b8152600401602060405180830381865afa158015610846573d5f5f3e3d5ffd5b505050506040513d601f19601f8201168201806040525081019061086a9190613af5565b6001600160a01b0316336001600160a01b0316145b6108c15760405162461bcd60e51b81526020600482015260136024820152723ab730baba3437b934bd32b21039b2b73232b960691b604482015260640161069b565b7f467f98200caa1453d527fd7de92faa829d8357fb650c10010e77188d93a3a3b6816040516020016108f39190613c91565b60408051601f198184030181529082905261090d91613891565b60405180910390a150565b60605f61094561092b60208601866136a4565b6001600160401b03165f908152600b602052604090205490565b9050806109945760405162461bcd60e51b815260206004820152601860248201527f416e63686f7220736c6f7420697320756e646566696e65640000000000000000604482015260640161069b565b6109a081856001611027565b5f6109b08561010001358561140c565b90506109c260a08501608086016136df565b6001600160a01b03166109db60408701602088016136a4565b6001600160401b03167f00000000000000000000000000000000000000000000000000000000000000007f17e2291123b98af9b416b45d04d46e4484bbf9d2c1061bbf567425a66bb1beb18860c001358860600135604051610a47929190918252602082015260400190565b60405180910390a4949350505050565b5f80610a666060840184613aac565b810190610a739190613fca565b905060408301357f000000000000000000000000000000000000000000000000000000000000000014610adb5760405162461bcd60e51b815260206004820152601060248201526f3130b21039b7bab931b29031b430b4b760811b604482015260640161069b565b63ffffffff7f000000000000000000000000000000000000000000000000000000000000000016610b12604085016020860161415b565b63ffffffff1614610b535760405162461bcd60e51b815260206004820152600b60248201526a626164206e6574776f726b60a81b604482015260640161069b565b6020810151608001516001600160a01b03163014610ba15760405162461bcd60e51b815260206004820152600b60248201526a3130b21032b6b4ba3a32b960a91b604482015260640161069b565b7f467f98200caa1453d527fd7de92faa829d8357fb650c10010e77188d93a3a3b6816020015160a0015114610c045760405162461bcd60e51b815260206004820152600960248201526862616420746f70696360b81b604482015260640161069b565b805160208201516040516375f6768760e11b81525f92309263ebeced0e92610c3092919060040161427a565b5f604051808303815f875af1158015610c4b573d5f5f3e3d5ffd5b505050506040513d5f823e601f3d908101601f19168201604052610c729190810190614367565b90505f81806020019051810190610c899190614367565b9050610c9585806143cf565b604051602001610ca59190613c91565b60405160208183030381529060405280519060200120818051906020012014610d035760405162461bcd60e51b815260206004820152601060248201526f0e0c2f2d8dec2c840dad2e6dac2e8c6d60831b604482015260640161069b565b506001949350505050565b610d188133611759565b50565b5f610d26838361060a565b610da3575f838152602081815260408083206001600160a01b03861684529091529020805460ff19166001179055610d5b3390565b6001600160a01b0316826001600160a01b0316847f2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d60405160405180910390a45060016104cc565b505f6104cc565b5f610db5838361060a565b15610da3575f838152602081815260408083206001600160a01b0386168085529252808320805460ff1916905551339286917ff6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b9190a45060016104cc565b60208083015180518051600780546001600160401b0392831667ffffffffffffffff19918216179091559184015160085582840151805160098054918416919094168117909355840151600a558551938401515160405192949116927f955222ca45c52776cd8084d242443738d71afc33822d853b8efeb94e070d94d092610e9c92919061442c565b60405180910390a361078e8183602001516020015160200151611792565b604080516080810182526007546001600160401b03908116828401908152600854606084015282528251808401909352600954168252600a546020838101919091528101919091528251610f0e9190611802565b610f2b57604051634247f10960e11b815260040160405180910390fd5b610f3c825f01518360200151611831565b610f5957604051632d776d6b60e11b815260040160405180910390fd5b5f6002610f668380613aac565b604051610f74929190614448565b602060405180830381855afa158015610f8f573d5f5f3e3d5ffd5b5050506040513d601f19601f82011682018060405250810190610fb29190614457565b600d549091506001600160a01b031663ab750e75610fd36020850185613aac565b600c54856040518563ffffffff1660e01b8152600401610ff6949392919061446e565b5f6040518083038186803b15801561100c575f5ffd5b505afa15801561101e573d5f5f3e3d5ffd5b50505050505050565b61103460208301836136a4565b6001600160401b031661104d60408401602085016136a4565b6001600160401b0316106110b55760405162461bcd60e51b815260206004820152602960248201527f54617267657420736c6f74206d757374206265206c657373207468616e20616e60448201526818da1bdc881cdb1bdd60ba1b606482015260840161069b565b60058101546110ca60408401602085016136a4565b6110d760208501856136a4565b6110e19190614494565b6001600160401b031611156111315760405162461bcd60e51b815260206004820152601660248201527515185c99d95d081cdb1bdd081a5cc81d1bdbc81bdb1960521b604482015260640161069b565b80545f906111849060408501359061114c60608701876144b3565b808060200260200160405190810160405280939291908181526020018383602002808284375f920191909152508a92506118ee915050565b9050806111d35760405162461bcd60e51b815260206004820152601960248201527f496e76616c696420616e63686f7220737461746520726f6f7400000000000000604482015260640161069b565b60058201545f906111ea60408601602087016136a4565b6001600160401b03166111fd919061450c565b90505f8184600301548560020154901b6112179190613904565b90505f61126a60808701358361123060a08a018a6144b3565b808060200260200160405190810160405280939291908181526020018383602002808284375f9201919091525050505060408a01356118ee565b9050806112b95760405162461bcd60e51b815260206004820181905260248201527f496e76616c69642074617267657420626561636f6e20737461746520726f6f74604482015260640161069b565b5f6113118760c001358760010154898060e001906112d791906144b3565b808060200260200160405190810160405280939291908181526020018383602002808284375f9201919091525050505060808b01356118ee565b9050806113605760405162461bcd60e51b815260206004820152601c60248201527f496e76616c696420657865637574696f6e20726f6f742070726f6f6600000000604482015260640161069b565b5f6113ba88610100013588600401548a80610120019061138091906144b3565b808060200260200160405190810160405280939291908181526020018383602002808284375f9201919091525050505060c08c01356118ee565b9050806114015760405162461bcd60e51b8152602060048201526015602482015274125b9d985b1a59081c9958d95a5c1d1cc81c9bdbdd605a1b604482015260640161069b565b505050505050505050565b6040805160018082528183019092526060915f9190816020015b606081526020019060019003908161142657905050905061144a6020840184613aac565b8080601f0160208091040260200160405190810160405280939291908181526020018383808284375f9201829052508551869450909250151590506114915761149161451f565b60209081029190910101525f6114ba856114ab86806144b3565b6114b491614533565b84611905565b905080516001146115205760405162461bcd60e51b815260206004820152602a60248201527f496e76616c6964206e756d626572206f6620726573756c747320696e2072656360448201526932b4b83a10383937b7b360b11b606482015260840161069b565b5f815f815181106115335761153361451f565b60200260200101516020015151116115855760405162461bcd60e51b815260206004820152601560248201527424b73b30b634b2103932b1b2b4b83a10383937b7b360591b604482015260640161069b565b5f6115b46115af835f8151811061159e5761159e61451f565b602002602001015160200151611d68565b611d94565b90508060800151518560600135106116025760405162461bcd60e51b8152602060048201526011602482015270092dcecc2d8d2c840d8dece40d2dcc8caf607b1b604482015260640161069b565b5f816080015186606001358151811061161d5761161d61451f565b6020026020010151905085608001602081019061163a91906136df565b6001600160a01b0316815f01516001600160a01b0316146116935760405162461bcd60e51b815260206004820152601360248201527224b73b30b634b2103637b39032b6b4ba3a32b960691b604482015260640161069b565b5f816020015151116116db5760405162461bcd60e51b81526020600482015260116024820152704c6f6720686173206e6f20746f7069637360781b604482015260640161069b565b8560a0013581602001515f815181106116f6576116f661451f565b60200260200101511461174b5760405162461bcd60e51b815260206004820152601760248201527f496e76616c6964206576656e74207369676e6174757265000000000000000000604482015260640161069b565b604001519695505050505050565b611763828261060a565b61078e5760405163e2517d3f60e01b81526001600160a01b03821660048201526024810183905260440161069b565b6001600160401b0382165f908152600b60205260409020546117c9576001600160401b0382165f908152600b602052604090208190555b60405181906001600160401b038416907f9f2417e1c378e9eae44d76184be142d77ed574882fbc98272fe0533053d8816d905f90a35050565b815181515f9161181191612047565b801561182a575061182a83602001518360200151612047565b9392505050565b6020808301515190820151515f91906001600160401b038083169082161161185d575f925050506104cc565b5f61186782610534565b90504281111561187c575f93505050506104cc565b600d54600160a01b900462ffffff16611895824261453f565b11156118a6575f93505050506104cc565b5f600c60206118b58686614494565b6001600160401b03166118c891906138ed565b6118d291906138ed565b600d54600160a01b900462ffffff161015979650505050505050565b5f816118fb868686612070565b1495945050505050565b60605f82516001600160401b0381111561192157611921613917565b60405190808252806020026020018201604052801561196657816020015b604080518082019091526060808252602082015281526020019060019003908161193f5790505b5090505f84516001600160401b0381111561198357611983613917565b6040519080825280602002602001820160405280156119c857816020015b604080518082019091525f8152606060208201528152602001906001900390816119a15790505b5090505f5b8551811015611a495760405180604001604052808783815181106119f3576119f361451f565b6020026020010151805190602001208152602001878381518110611a1957611a1961451f565b6020026020010151815250828281518110611a3657611a3661451f565b60209081029190910101526001016119cd565b505f5b8451811015611d5d57848181518110611a6757611a6761451f565b6020026020010151838281518110611a8157611a8161451f565b60200260200101515f01819052505f6040518060400160405280878481518110611aad57611aad61451f565b602002602001015181526020015f81525090505f611ad3611ace858b6121de565b612279565b905060015b8015611d5257604080516080810182525f8082526020820181905291810191909152606080820152611b09836123cd565b15611bd2575f611b18846123e1565b90505f60028660200151611b2c919061450c565b15611b525760028660200151611b429190614552565b611b4d906001613904565b611b63565b60028660200151611b639190614552565b90506040518060400160405280611b7d885f015184612486565b81526020015f8152509550611b95825f01518761254e565b15611bca57611ba8888360200151612573565b898881518110611bba57611bba61451f565b6020026020010151602001819052505b505050611d52565b60c083015115611c6b575f611be6846125bb565b9050611bf585825f01516126d7565b15611c5e575f611c07825f01516126f3565b8660200151611c169190613904565b90506040518060400160405280611c3c885f0151600285611c379190614552565b612486565b8152602001611c4c60028461450c565b81525095508160200151925050611c65565b5050611d52565b50611d2f565b60e083015115611d22575f611c7f84612714565b9050611c8a85612908565b15611cd4578051511515600103611c5e57611cac87825f015160200151612573565b888781518110611cbe57611cbe61451f565b6020026020010151602001819052505050611d52565b5f8160200151611ce4875f612919565b60108110611cf457611cf461451f565b602002015180519091501515600103611bca57611d1286600161299f565b9550806020015192505050611d2f565b825115611d2f5750611d52565b611d3c611ace8783612573565b9250508080611d4a90614565565b915050611ad8565b505050600101611a4c565b509095945050505050565b6040805180820182525f8082526020918201528151808301909252825182529182019181019190915290565b6040805160a0810182525f808252606060208301819052928201528181018290526080810191909152611dc6826129e0565b611e6157602082015180515f1a60018110801590611de8575060048160ff1611155b611e2a5760405162461bcd60e51b8152602060048201526013602482015272556e737570706f72746564207478207479706560681b604482015260640161069b565b60ff811683526040805180820190915284518190611e4a9060019061453f565b8152602001611e5a846001613904565b9052935050505b5f611e6b83612a17565b90508051600414611ec95760405162461bcd60e51b815260206004820152602260248201527f496e76616c696420726563656970743a2065787065637465642034206669656c604482015261647360f01b606482015260840161069b565b611eeb815f81518110611ede57611ede61451f565b6020026020010151612b1c565b8260200181905250611f1681600181518110611f0957611f0961451f565b6020026020010151612b8d565b6001600160401b031660408301528051611f3d9082906002908110611ede57611ede61451f565b82606001819052505f611f6982600381518110611f5c57611f5c61451f565b6020026020010151612a17565b8051909150806001600160401b03811115611f8657611f86613917565b604051908082528060200260200182016040528015611fe357816020015b611fd060405180606001604052805f6001600160a01b0316815260200160608152602001606081525090565b815260200190600190039081611fa45790505b5060808501525f5b8181101561203e576120158382815181106120085761200861451f565b6020026020010151612bd8565b8560800151828151811061202b5761202b61451f565b6020908102919091010152600101611feb565b50505050919050565b805182515f916001600160401b03918216911614801561182a5750506020908101519101511490565b5f83815b846001146121d55761208760028661450c565b6001036121245760028482815181106120a2576120a261451f565b6020026020010151836040516020016120c5929190918252602082015260400190565b60408051601f19818403018152908290526120df9161457d565b602060405180830381855afa1580156120fa573d5f5f3e3d5ffd5b5050506040513d601f19601f8201168201806040525081019061211d9190614457565b91506121b6565b6002828583815181106121395761213961451f565b602002602001015160405160200161215b929190918252602082015260400190565b60408051601f19818403018152908290526121759161457d565b602060405180830381855afa158015612190573d5f5f3e3d5ffd5b5050506040513d601f19601f820116820180604052508101906121b39190614457565b91505b6121c1600286614552565b9450806121cd81614565565b915050612074565b50949350505050565b60605f5b835181101561223c57828482815181106121fe576121fe61451f565b60200260200101515f015103612234578381815181106122205761222061451f565b6020026020010151602001519150506104cc565b6001016121e2565b5060405162461bcd60e51b8152602060048201526011602482015270496e636f6d706c6574652050726f6f662160781b604482015260640161069b565b6122816134c5565b6122896134c5565b5f60405180604001604052808581526020015f81525090506122e0846040518060400160405280602081526020017f56e81f171bcc55a6ff8345e692c0f86e5b48e01b996cadc001622fb5e363b421815250612dac565b156122f057506001815292915050565b5f6123026122fd86611d68565b612a17565b80519091505f81900361231c575050600182525092915050565b80600203612374575f61233a835f81518110611ede57611ede61451f565b602081015190915060fc1c60028114806123545750806003145b15612365576001602087015261236d565b600160c08701525b50506123bf565b8060110361238857600160e08501526123bf565b60405162461bcd60e51b815260206004820152600c60248201526b496e76616c6964206461746160a01b604482015260640161069b565b505061012082015292915050565b5f8160200151806104cc5750506040015190565b6123e961352a565b6123f161352a565b5f6124066122fd8561012001515f0151611d68565b90505f61241f82600181518110611ede57611ede61451f565b9050604051806040016040528061244b612444855f81518110611ede57611ede61451f565b6001612dda565b81525f602091820181905291855260408051608081018252838152808301939093526001908301526060820192909252908301525092915050565b81516060905f612496848361453f565b9050838210156124dc5760405162461bcd60e51b8152602060048201526011602482015270736c6963655f6f75744f66426f756e647360781b604482015260640161069b565b6060811580156124fa5760405191505f825260208201604052612544565b6040519150601f8316801560200281840101848101888315602002848c0101015b8183101561253357805183526020928301920161251b565b5050848452601f01601f1916604052505b5095945050505050565b5f612558826126f3565b612561846126f3565b14801561182a575061182a83836126d7565b606081604001511561258a575060608101516104cc565b8151156125a65761259f8383602001516121de565b90506104cc565b5060408051602081019091525f815292915050565b6125c361352a565b6125cb61352a565b5f6125e06122fd8561012001515f0151611d68565b90505f6125f982600181518110611ede57611ede61451f565b90505f6004612613845f81518110611ede57611ede61451f565b5f815181106126245761262461451f565b602001015160f81c60f81b6001600160f81b031916901c60f81c60011690506040518060400160405280612685612666865f81518110611ede57611ede61451f565b600261267386600161458e565b61267d91906145a7565b60ff16612dda565b815260ff83166020918201529085526040805160808101909152600181529081016126af84612e10565b81525f6020808301829052604080519283528282018152909201528501525091949350505050565b5f6126e1826126f3565b6126eb8484612e77565b149392505050565b5f81602001516002835f01515161270a91906138ed565b6104cc919061453f565b61271c613579565b612724613579565b5f6127396122fd8561012001515f0151611d68565b9050612743613595565b5f5b601081101561283c575f612764848381518110611ede57611ede61451f565b905080516020146127cb576040805180820182525f8082528251608081018452818152602080820183905281850183905284519283528281019094526060810191909152918101919091528383601081106127c1576127c161451f565b6020020152612833565b5f6127d582612e10565b6040805180820182526001808252825160808101845290815260208082018590525f828501819052845190815280820190945260608201939093529181019190915290915084846010811061282c5761282c61451f565b6020020152505b50600101612745565b5061285b61285683601081518110611ede57611ede61451f565b612fe9565b156128a7576040805180820182525f80825282516080810184528181526020808201839052818501839052845192835282810190945260608101919091529181019190915283526128fc565b604051806040016040528060011515815260200160405180608001604052805f151581526020015f5f1b81526020016001151581526020016128f586601081518110611ede57611ede61451f565b9052905283525b60208301525092915050565b5f612912826126f3565b1592915050565b5f5f600283856020015161292d9190613904565b6129379190614552565b90505f600284866020015161294c9190613904565b612956919061450c565b90505f855f0151838151811061296e5761296e61451f565b016020015160f81c90506001821461298d5760048160ff16901c612992565b80600f165b60ff169695505050505050565b60408051808201909152606081525f60208201526040518060400160405280845f015181526020018385602001516129d79190613904565b90529392505050565b80515f9081036129f157505f919050565b602082015180515f1a9060c0821015612a0d57505f9392505050565b5060019392505050565b6060612a22826129e0565b612a2a575f5ffd5b5f612a3483613051565b90505f816001600160401b03811115612a4f57612a4f613917565b604051908082528060200260200182016040528015612a9357816020015b604080518082019091525f8082526020820152815260200190600190039081612a6d5790505b5090505f612aa485602001516130d4565b8560200151612ab39190613904565b90505f805b84811015612b1157612ac98361314d565b9150604051806040016040528083815260200184815250848281518110612af257612af261451f565b6020908102919091010152612b078284613904565b9250600101612ab8565b509195945050505050565b8051606090612b29575f5ffd5b5f5f612b34846131f5565b915091505f816001600160401b03811115612b5157612b51613917565b6040519080825280601f01601f191660200182016040528015612b7b576020820181803683370190505b509050602081016121d5848285613237565b80515f9015801590612ba157508151602110155b612ba9575f5ffd5b5f5f612bb4846131f5565b815191935091506020821015612bd05760208290036101000a90045b949350505050565b612c0460405180606001604052805f6001600160a01b0316815260200160608152602001606081525090565b5f612c0e83612a17565b90508051600314612c615760405162461bcd60e51b815260206004820152601e60248201527f496e76616c6964206c6f673a2065787065637465642033206669656c64730000604482015260640161069b565b612c83815f81518110612c7657612c7661451f565b60200260200101516132b7565b6001600160a01b031682528051612ca79082906002908110611ede57611ede61451f565b82604001819052505f612cc682600181518110611f5c57611f5c61451f565b80519091506004811115612d1c5760405162461bcd60e51b815260206004820152601c60248201527f496e76616c6964206c6f673a20746f6f206d616e7920746f7069637300000000604482015260640161069b565b806001600160401b03811115612d3457612d34613917565b604051908082528060200260200182016040528015612d5d578160200160208202803683370190505b5060208501525f5b8181101561203e57612d82838281518110611f0957611f0961451f565b6020860151805183908110612d9957612d9961451f565b6020908102919091010152600101612d65565b5f8151835114612dbd57505f6104cc565b825160208381018281209186019283209091145b95945050505050565b60608251821115612de9575f5ffd5b5f828451612df7919061453f565b905060208401612dd1612e0a8583613904565b836132cf565b5f602082511015612e6f5760405162461bcd60e51b8152602060048201526024808201527f42797465733a3a20746f427974657333323a206461746120697320746f20736860448201526337b93a1760e11b606482015260840161069b565b506020015190565b5f5f60028460200151612e8a919061450c565b90505f60028460200151612e9e919061450c565b9050808203612f8e575f60028660200151612eb99190614552565b90505f60028660200151612ecd9190614552565b90505f8415612f4b578651805183908110612eea57612eea61451f565b016020015188518051600f60f81b9092169185908110612f0c57612f0c61451f565b0160200151600f60f81b1614612f29575f955050505050506104cc565b612f3283614565565b9250612f3d82614565565b9150612f4881614565565b90505b5f612f59895f015185612486565b90505f612f69895f015185612486565b905082612f768383613325565b612f809190613904565b9750505050505050506104cc565b5f612fa9612f9b876126f3565b612fa4876126f3565b61340b565b90505f5b81811015612fde57612fbf8682612919565b612fc98883612919565b03612fde57612fd781614565565b9050612fad565b93506104cc92505050565b5f5f82511180156104cc5750815f815181106130075761300761451f565b6020910101516001600160f81b031916600360fe1b14806104cc5750815f815181106130355761303561451f565b6020910101516001600160f81b031916600160ff1b1492915050565b80515f90810361306257505f919050565b5f5f90505f61307484602001516130d4565b84602001516130839190613904565b90505f845f015185602001516130999190613904565b90505b808210156130cb576130ad8261314d565b6130b79083613904565b9150826130c381614565565b93505061309c565b50909392505050565b80515f90811a60808110156130eb57505f92915050565b60b8811080613106575060c08110801590613106575060f881105b156131145750600192915050565b60c081101561314157613129600160b86145c8565b6131369060ff168261453f565b61182a906001613904565b613129600160f86145c8565b80515f908190811a608081101561316757600191506131ee565b60b881101561318d5761317b60808261453f565b613186906001613904565b91506131ee565b60c08110156131ba5760b78103600185019450806020036101000a855104600182018101935050506131ee565b60f88110156131ce5761317b60c08261453f565b60f78103600185019450806020036101000a855104600182018101935050505b5092915050565b5f5f5f61320584602001516130d4565b90505f8185602001516132189190613904565b90505f82865f015161322a919061453f565b9196919550909350505050565b805f0361324357505050565b6020811061327b578251825261325a602084613904565b9250613267602083613904565b915061327460208261453f565b9050613243565b801561052f575f600161328f83602061453f565b61329b906101006146c4565b6132a5919061453f565b84518451821691191617835250505050565b80515f906015146132c6575f5ffd5b6104cc82612b8d565b6060816001600160401b038111156132e9576132e9613917565b6040519080825280601f01601f191660200182016040528015613313576020820181803683370190505b509050602081016131ee848285613420565b5f5f6133338451845161340b565b90505f5b81811015613400578381815181106133515761335161451f565b602001015160f81c60f81b6001600160f81b0319168582815181106133785761337861451f565b01602001516001600160f81b031916146133f0576133d28582815181106133a1576133a161451f565b602001015160f81c60f81b8583815181106133be576133be61451f565b01602001516001600160f81b031916613487565b6133dd6002836138ed565b6133e79190613904565b925050506104cc565b6133f981614565565b9050613337565b612dd16002826138ed565b5f818310613419578161182a565b5090919050565b602081106134585782518252613437602083613904565b9150613444602084613904565b925061345160208261453f565b9050613420565b5f81156104f657600161346c83602061453f565b613478906101006146c4565b613482919061453f565b6132a5565b5f6001600160f81b0319808316908416036134a4575060026104cc565b8160f81c60f01660ff168360f81c60f01660ff1603610da3575060016104cc565b60408051610140810182525f808252602080830182905282840182905260608084018390526080840183905260a0840183905260c0840183905260e084018390526101008401839052845180860190955284528301529061012082015290565b905290565b60405180604001604052806135516040518060400160405280606081526020015f81525090565b8152604080516080810182525f80825260208281018290529282015260608082015291015290565b604051806040016040528061358c6135c3565b81526020016135255b6040518061020001604052806010905b6135ad6135c3565b8152602001906001900390816135a55790505090565b60405180604001604052805f15158152602001613525604080516080810182525f808252602082018190529181019190915260608082015290565b5f6020828403121561360e575f5ffd5b81356001600160e01b03198116811461182a575f5ffd5b5f60208284031215613635575f5ffd5b5035919050565b6001600160a01b0381168114610d18575f5ffd5b803561365b8161363c565b919050565b5f5f60408385031215613671575f5ffd5b8235915060208301356136838161363c565b809150509250929050565b80356001600160401b038116811461365b575f5ffd5b5f602082840312156136b4575f5ffd5b61182a8261368e565b5f602082840312156136cd575f5ffd5b813562ffffff8116811461182a575f5ffd5b5f602082840312156136ef575f5ffd5b813561182a8161363c565b5f5f5f6040848603121561370c575f5ffd5b83356001600160401b03811115613721575f5ffd5b8401601f81018613613731575f5ffd5b80356001600160401b03811115613746575f5ffd5b866020828401011115613757575f5ffd5b60209182019450925061376b90850161368e565b90509250925092565b5f60208284031215613784575f5ffd5b81356001600160401b03811115613799575f5ffd5b82016060818503121561182a575f5ffd5b5f602082840312156137ba575f5ffd5b81356001600160401b038111156137cf575f5ffd5b8201610120818503121561182a575f5ffd5b5f5f604083850312156137f2575f5ffd5b82356001600160401b03811115613807575f5ffd5b83016101408186031215613819575f5ffd5b915060208301356001600160401b03811115613833575f5ffd5b830160c08186031215613683575f5ffd5b5f5b8381101561385e578181015183820152602001613846565b50505f910152565b5f815180845261387d816020860160208601613844565b601f01601f19169290920160200192915050565b602081525f61182a6020830184613866565b5f602082840312156138b3575f5ffd5b81356001600160401b038111156138c8575f5ffd5b82016080818503121561182a575f5ffd5b634e487b7160e01b5f52601160045260245ffd5b80820281158282048414176104cc576104cc6138d9565b808201808211156104cc576104cc6138d9565b634e487b7160e01b5f52604160045260245ffd5b604080519081016001600160401b038111828210171561394d5761394d613917565b60405290565b60405160c081016001600160401b038111828210171561394d5761394d613917565b60405161014081016001600160401b038111828210171561394d5761394d613917565b604051601f8201601f191681016001600160401b03811182821017156139c0576139c0613917565b604052919050565b5f604082840312156139d8575f5ffd5b6139e061392b565b90506139eb8261368e565b815260209182013591810191909152919050565b5f60808284031215613a0f575f5ffd5b613a1761392b565b9050613a2383836139c8565b8152613a3283604084016139c8565b602082015292915050565b5f610120828403128015613a4f575f5ffd5b50604051606081016001600160401b0381118282101715613a7257613a72613917565b604052613a7f84846139ff565b8152613a8e84608085016139ff565b6020820152613aa0610100840161368e565b60408201529392505050565b5f5f8335601e19843603018112613ac1575f5ffd5b8301803591506001600160401b03821115613ada575f5ffd5b602001915036819003821315613aee575f5ffd5b9250929050565b5f60208284031215613b05575f5ffd5b815161182a8161363c565b5f5f8335601e19843603018112613b25575f5ffd5b83016020810192503590506001600160401b03811115613b43575f5ffd5b8060051b3603821315613aee575f5ffd5b8183526020830192505f815f5b84811015613b92578135613b748161363c565b6001600160a01b031686526020958601959190910190600101613b61565b5093949350505050565b5f5f8335601e19843603018112613bb1575f5ffd5b83016020810192503590506001600160401b03811115613bcf575f5ffd5b8060061b3603821315613aee575f5ffd5b8183526020830192505f815f5b84811015613b9257813586526020820135613c078161363c565b6001600160a01b031660208701526040958601959190910190600101613bed565b5f5f8335601e19843603018112613c3d575f5ffd5b83016020810192503590506001600160401b03811115613c5b575f5ffd5b803603821315613aee575f5ffd5b81835281816020850137505f828201602090810191909152601f909101601f19169091010190565b60208082528235828201525f90613ca9908401613650565b6001600160a01b038116604084015250613cc560408401613650565b6001600160a01b0381166060840152506060830135608083810191909152613cee908401613650565b6001600160a01b03811660a08401525060a083013560c083810191909152613d1890840184613b10565b61012060e0850152613d2f61014085018284613b54565b915050613d3f60e0850185613b9c565b848303601f1901610100860152613d57838284613be0565b92505050613d69610100850185613c28565b848303601f1901610120860152613d81838284613c69565b9695505050505050565b5f6001600160401b03821115613da357613da3613917565b5060051b60200190565b5f82601f830112613dbc575f5ffd5b8135613dcf613dca82613d8b565b613998565b8082825260208201915060208360051b860101925085831115613df0575f5ffd5b602085015b83811015612544578035835260209283019201613df5565b5f6001600160401b03821115613e2557613e25613917565b50601f01601f191660200190565b5f82601f830112613e42575f5ffd5b8135613e50613dca82613e0d565b818152846020838601011115613e64575f5ffd5b816020850160208301375f918101602001919091529392505050565b5f613e8d613dca84613d8b565b838152905060208101600584901b830185811115613ea9575f5ffd5b835b81811015613ee35780356001600160401b03811115613ec8575f5ffd5b613ed488828801613e33565b84525060209283019201613eab565b5050509392505050565b5f82601f830112613efc575f5ffd5b61182a83833560208501613e80565b5f60c08284031215613f1b575f5ffd5b613f23613953565b905081356001600160401b03811115613f3a575f5ffd5b613f4684828501613eed565b82525060208201356001600160401b03811115613f61575f5ffd5b613f6d84828501613e33565b60208301525060408201356001600160401b03811115613f8b575f5ffd5b613f9784828501613e33565b60408301525060608281013590820152613fb360808301613650565b608082015260a09182013591810191909152919050565b5f60208284031215613fda575f5ffd5b81356001600160401b03811115613fef575f5ffd5b820160408185031215614000575f5ffd5b61400861392b565b81356001600160401b0381111561401d575f5ffd5b8201610140818703121561402f575f5ffd5b614037613975565b6140408261368e565b815261404e6020830161368e565b60208201526040828101359082015260608201356001600160401b03811115614075575f5ffd5b61408188828501613dad565b6060830152506080828101359082015260a08201356001600160401b038111156140a9575f5ffd5b6140b588828501613dad565b60a08301525060c0828101359082015260e08201356001600160401b038111156140dd575f5ffd5b6140e988828501613dad565b60e08301525061010082810135908201526101208201356001600160401b03811115614113575f5ffd5b61411f88828501613dad565b6101208301525082525060208201356001600160401b03811115614141575f5ffd5b61414d86828501613f0b565b602083015250949350505050565b5f6020828403121561416b575f5ffd5b813563ffffffff8116811461182a575f5ffd5b5f8151808452602084019350602083015f5b82811015613b92578151865260209586019590910190600101614190565b5f60c08301825160c0855281815180845260e08701915060e08160051b88010193506020830192505f5b818110156142095760df198886030183526141f4858551613866565b945060209384019392909201916001016141d8565b50505050602083015184820360208601526142248282613866565b9150506040830151848203604086015261423e8282613866565b91505060608301516060850152608083015161426560808601826001600160a01b03169052565b5060a083015160a08501528091505092915050565b604081526142946040820184516001600160401b03169052565b5f60208401516142af60608401826001600160401b03169052565b5060408401516080830152606084015161014060a08401526142d561018084018261417e565b9050608085015160c084015260a0850151603f198483030160e08501526142fc828261417e565b91505060c085015161010084015260e0850151603f1984830301610120850152614326828261417e565b915050610100850151610140840152610120850151603f1984830301610160850152614352828261417e565b9150508281036020840152612dd181856141ae565b5f60208284031215614377575f5ffd5b81516001600160401b0381111561438c575f5ffd5b8201601f8101841361439c575f5ffd5b80516143aa613dca82613e0d565b8181528560208385010111156143be575f5ffd5b612dd1826020830160208601613844565b5f823561011e198336030181126143e4575f5ffd5b9190910192915050565b61440c82825180516001600160401b03168252602090810151910152565b60209081015180516001600160401b031660408401520151606090910152565b610100810161443b82856143ee565b61182a60808301846143ee565b818382375f9101908152919050565b5f60208284031215614467575f5ffd5b5051919050565b606081525f614481606083018688613c69565b6020830194909452506040015292915050565b6001600160401b0382811682821603908111156104cc576104cc6138d9565b5f5f8335601e198436030181126144c8575f5ffd5b8301803591506001600160401b038211156144e1575f5ffd5b6020019150600581901b3603821315613aee575f5ffd5b634e487b7160e01b5f52601260045260245ffd5b5f8261451a5761451a6144f8565b500690565b634e487b7160e01b5f52603260045260245ffd5b5f61182a368484613e80565b818103818111156104cc576104cc6138d9565b5f82614560576145606144f8565b500490565b5f60018201614576576145766138d9565b5060010190565b5f82516143e4818460208701613844565b60ff81811683821601908111156104cc576104cc6138d9565b5f60ff8316806145b9576145b96144f8565b8060ff84160691505092915050565b60ff82811682821603908111156104cc576104cc6138d9565b6001815b600184111561461c57808504811115614600576146006138d9565b600184161561460e57908102905b60019390931c9280026145e5565b935093915050565b5f82614632575060016104cc565b8161463e57505f6104cc565b8160018114614654576002811461465e5761467a565b60019150506104cc565b60ff84111561466f5761466f6138d9565b50506001821b6104cc565b5060208310610133831016604e8410600b841016171561469d575081810a6104cc565b6146a95f1984846145e1565b805f19048211156146bc576146bc6138d9565b029392505050565b5f61182a838361462456fea49807205ce4d355092ef5a8a18f56e8913cf4a201fbe287825b095693c21775a26469706673582212209fa137d02da2f66856cc85b03e8a93fc15e22cac711f73a16bd7fcde982f61b864736f6c634300081e0033",
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

// VerifyLogAndExtract is a paid mutator transaction binding the contract method 0xebeced0e.
//
// Solidity: function verifyLogAndExtract((uint64,uint64,bytes32,bytes32[],bytes32,bytes32[],bytes32,bytes32[],bytes32,bytes32[]) execProof, (bytes[],bytes,bytes,uint256,address,bytes32) logProof) returns(bytes)
func (_ZKAdapter *ZKAdapterTransactor) VerifyLogAndExtract(opts *bind.TransactOpts, execProof ExecutionProof, logProof ReceiptProof) (*types.Transaction, error) {
	return _ZKAdapter.contract.Transact(opts, "verifyLogAndExtract", execProof, logProof)
}

// VerifyLogAndExtract is a paid mutator transaction binding the contract method 0xebeced0e.
//
// Solidity: function verifyLogAndExtract((uint64,uint64,bytes32,bytes32[],bytes32,bytes32[],bytes32,bytes32[],bytes32,bytes32[]) execProof, (bytes[],bytes,bytes,uint256,address,bytes32) logProof) returns(bytes)
func (_ZKAdapter *ZKAdapterSession) VerifyLogAndExtract(execProof ExecutionProof, logProof ReceiptProof) (*types.Transaction, error) {
	return _ZKAdapter.Contract.VerifyLogAndExtract(&_ZKAdapter.TransactOpts, execProof, logProof)
}

// VerifyLogAndExtract is a paid mutator transaction binding the contract method 0xebeced0e.
//
// Solidity: function verifyLogAndExtract((uint64,uint64,bytes32,bytes32[],bytes32,bytes32[],bytes32,bytes32[],bytes32,bytes32[]) execProof, (bytes[],bytes,bytes,uint256,address,bytes32) logProof) returns(bytes)
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
