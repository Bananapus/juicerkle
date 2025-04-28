// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package main

import (
	"errors"
	"math/big"
	"strings"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/event"
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

// JBClaim is an auto generated low-level Go binding around an user-defined struct.
type JBClaim struct {
	Token common.Address
	Leaf  JBLeaf
	Proof [32][32]byte
}

// JBInboxTreeRoot is an auto generated low-level Go binding around an user-defined struct.
type JBInboxTreeRoot struct {
	Nonce uint64
	Root  [32]byte
}

// JBLeaf is an auto generated low-level Go binding around an user-defined struct.
type JBLeaf struct {
	Index               *big.Int
	Beneficiary         common.Address
	ProjectTokenCount   *big.Int
	TerminalTokenAmount *big.Int
}

// JBMessageRoot is an auto generated low-level Go binding around an user-defined struct.
type JBMessageRoot struct {
	Token      common.Address
	Amount     *big.Int
	RemoteRoot JBInboxTreeRoot
}

// JBOutboxTree is an auto generated low-level Go binding around an user-defined struct.
type JBOutboxTree struct {
	Nonce              uint64
	Balance            *big.Int
	Tree               MerkleLibTree
	NumberOfClaimsSent *big.Int
}

// JBRemoteToken is an auto generated low-level Go binding around an user-defined struct.
type JBRemoteToken struct {
	Enabled         bool
	EmergencyHatch  bool
	MinGas          uint32
	Addr            common.Address
	MinBridgeAmount *big.Int
}

// JBTokenMapping is an auto generated low-level Go binding around an user-defined struct.
type JBTokenMapping struct {
	LocalToken      common.Address
	MinGas          uint32
	RemoteToken     common.Address
	MinBridgeAmount *big.Int
}

// MerkleLibTree is an auto generated low-level Go binding around an user-defined struct.
type MerkleLibTree struct {
	Branch [32][32]byte
	Count  *big.Int
}

// JBSuckerMetaData contains all meta data concerning the JBSucker contract.
var JBSuckerMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"receive\",\"stateMutability\":\"payable\"},{\"type\":\"function\",\"name\":\"ADD_TO_BALANCE_MODE\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint8\",\"internalType\":\"enumJBAddToBalanceMode\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"DEPLOYER\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"DIRECTORY\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIJBDirectory\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"MESSENGER_BASE_GAS_LIMIT\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"MESSENGER_ERC20_MIN_GAS_LIMIT\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"PERMISSIONS\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIJBPermissions\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"TOKENS\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIJBTokens\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"addOutstandingAmountToBalance\",\"inputs\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"amountToAddToBalanceOf\",\"inputs\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"claim\",\"inputs\":[{\"name\":\"claimData\",\"type\":\"tuple\",\"internalType\":\"structJBClaim\",\"components\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"leaf\",\"type\":\"tuple\",\"internalType\":\"structJBLeaf\",\"components\":[{\"name\":\"index\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"beneficiary\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"projectTokenCount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"terminalTokenAmount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"proof\",\"type\":\"bytes32[32]\",\"internalType\":\"bytes32[32]\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"claim\",\"inputs\":[{\"name\":\"claims\",\"type\":\"tuple[]\",\"internalType\":\"structJBClaim[]\",\"components\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"leaf\",\"type\":\"tuple\",\"internalType\":\"structJBLeaf\",\"components\":[{\"name\":\"index\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"beneficiary\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"projectTokenCount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"terminalTokenAmount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"proof\",\"type\":\"bytes32[32]\",\"internalType\":\"bytes32[32]\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"enableEmergencyHatchFor\",\"inputs\":[{\"name\":\"tokens\",\"type\":\"address[]\",\"internalType\":\"address[]\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"exitThroughEmergencyHatch\",\"inputs\":[{\"name\":\"claimData\",\"type\":\"tuple\",\"internalType\":\"structJBClaim\",\"components\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"leaf\",\"type\":\"tuple\",\"internalType\":\"structJBLeaf\",\"components\":[{\"name\":\"index\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"beneficiary\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"projectTokenCount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"terminalTokenAmount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"proof\",\"type\":\"bytes32[32]\",\"internalType\":\"bytes32[32]\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"fromRemote\",\"inputs\":[{\"name\":\"root\",\"type\":\"tuple\",\"internalType\":\"structJBMessageRoot\",\"components\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"remoteRoot\",\"type\":\"tuple\",\"internalType\":\"structJBInboxTreeRoot\",\"components\":[{\"name\":\"nonce\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"root\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]}]}],\"outputs\":[],\"stateMutability\":\"payable\"},{\"type\":\"function\",\"name\":\"inboxOf\",\"inputs\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structJBInboxTreeRoot\",\"components\":[{\"name\":\"nonce\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"root\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"initialize\",\"inputs\":[{\"name\":\"__projectId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"isMapped\",\"inputs\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"isTrustedForwarder\",\"inputs\":[{\"name\":\"forwarder\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"mapToken\",\"inputs\":[{\"name\":\"map\",\"type\":\"tuple\",\"internalType\":\"structJBTokenMapping\",\"components\":[{\"name\":\"localToken\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"minGas\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"remoteToken\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"minBridgeAmount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}],\"outputs\":[],\"stateMutability\":\"payable\"},{\"type\":\"function\",\"name\":\"mapTokens\",\"inputs\":[{\"name\":\"maps\",\"type\":\"tuple[]\",\"internalType\":\"structJBTokenMapping[]\",\"components\":[{\"name\":\"localToken\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"minGas\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"remoteToken\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"minBridgeAmount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}],\"outputs\":[],\"stateMutability\":\"payable\"},{\"type\":\"function\",\"name\":\"outboxOf\",\"inputs\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structJBOutboxTree\",\"components\":[{\"name\":\"nonce\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"balance\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"tree\",\"type\":\"tuple\",\"internalType\":\"structMerkleLib.Tree\",\"components\":[{\"name\":\"branch\",\"type\":\"bytes32[32]\",\"internalType\":\"bytes32[32]\"},{\"name\":\"count\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"numberOfClaimsSent\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"peer\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"peerChainId\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"prepare\",\"inputs\":[{\"name\":\"projectTokenCount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"beneficiary\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"minTokensReclaimed\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"projectId\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"remoteTokenFor\",\"inputs\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structJBRemoteToken\",\"components\":[{\"name\":\"enabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"emergencyHatch\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"minGas\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"addr\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"minBridgeAmount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"setDeprecation\",\"inputs\":[{\"name\":\"timestamp\",\"type\":\"uint40\",\"internalType\":\"uint40\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"state\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint8\",\"internalType\":\"enumJBSuckerState\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"supportsInterface\",\"inputs\":[{\"name\":\"interfaceId\",\"type\":\"bytes4\",\"internalType\":\"bytes4\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"toRemote\",\"inputs\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"payable\"},{\"type\":\"function\",\"name\":\"trustedForwarder\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"event\",\"name\":\"Claimed\",\"inputs\":[{\"name\":\"beneficiary\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"token\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"projectTokenCount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"terminalTokenAmount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"index\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"autoAddedToBalance\",\"type\":\"bool\",\"indexed\":false,\"internalType\":\"bool\"},{\"name\":\"caller\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"DeprecationTimeUpdated\",\"inputs\":[{\"name\":\"timestamp\",\"type\":\"uint40\",\"indexed\":false,\"internalType\":\"uint40\"},{\"name\":\"caller\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"EmergencyHatchOpened\",\"inputs\":[{\"name\":\"tokens\",\"type\":\"address[]\",\"indexed\":false,\"internalType\":\"address[]\"},{\"name\":\"caller\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Initialized\",\"inputs\":[{\"name\":\"version\",\"type\":\"uint64\",\"indexed\":false,\"internalType\":\"uint64\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"InsertToOutboxTree\",\"inputs\":[{\"name\":\"beneficiary\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"token\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"hashed\",\"type\":\"bytes32\",\"indexed\":false,\"internalType\":\"bytes32\"},{\"name\":\"index\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"root\",\"type\":\"bytes32\",\"indexed\":false,\"internalType\":\"bytes32\"},{\"name\":\"projectTokenCount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"terminalTokenAmount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"caller\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"NewInboxTreeRoot\",\"inputs\":[{\"name\":\"token\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"nonce\",\"type\":\"uint64\",\"indexed\":false,\"internalType\":\"uint64\"},{\"name\":\"root\",\"type\":\"bytes32\",\"indexed\":false,\"internalType\":\"bytes32\"},{\"name\":\"caller\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"RootToRemote\",\"inputs\":[{\"name\":\"root\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"token\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"index\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"nonce\",\"type\":\"uint64\",\"indexed\":false,\"internalType\":\"uint64\"},{\"name\":\"caller\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"InvalidInitialization\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"JBPermissioned_Unauthorized\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"sender\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"projectId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"permissionId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"JBSucker_BelowMinGas\",\"inputs\":[{\"name\":\"minGas\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"minGasLimit\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"JBSucker_Deprecated\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"JBSucker_DeprecationTimestampTooSoon\",\"inputs\":[{\"name\":\"givenTime\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"minimumTime\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"JBSucker_ExpectedMsgValue\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"JBSucker_InsufficientBalance\",\"inputs\":[{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"balance\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"JBSucker_InsufficientMsgValue\",\"inputs\":[{\"name\":\"received\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"expected\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"JBSucker_InvalidNativeRemoteAddress\",\"inputs\":[{\"name\":\"remoteToken\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"JBSucker_InvalidProof\",\"inputs\":[{\"name\":\"root\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"inboxRoot\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]},{\"type\":\"error\",\"name\":\"JBSucker_LeafAlreadyExecuted\",\"inputs\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"index\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"JBSucker_ManualNotAllowed\",\"inputs\":[{\"name\":\"mode\",\"type\":\"uint8\",\"internalType\":\"enumJBAddToBalanceMode\"}]},{\"type\":\"error\",\"name\":\"JBSucker_NoTerminalForToken\",\"inputs\":[{\"name\":\"projectId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"JBSucker_NotPeer\",\"inputs\":[{\"name\":\"caller\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"JBSucker_QueueInsufficientSize\",\"inputs\":[{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"minimumAmount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"JBSucker_TokenAlreadyMapped\",\"inputs\":[{\"name\":\"localToken\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"mappedTo\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"JBSucker_TokenHasInvalidEmergencyHatchState\",\"inputs\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"JBSucker_TokenNotMapped\",\"inputs\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"JBSucker_UnexpectedMsgValue\",\"inputs\":[{\"name\":\"value\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"JBSucker_ZeroBeneficiary\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"JBSucker_ZeroERC20Token\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"MerkleLib__insert_treeIsFull\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"NotInitializing\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"SafeERC20FailedOperation\",\"inputs\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"}]}]",
}

// JBSuckerABI is the input ABI used to generate the binding from.
// Deprecated: Use JBSuckerMetaData.ABI instead.
var JBSuckerABI = JBSuckerMetaData.ABI

// JBSucker is an auto generated Go binding around an Ethereum contract.
type JBSucker struct {
	JBSuckerCaller     // Read-only binding to the contract
	JBSuckerTransactor // Write-only binding to the contract
	JBSuckerFilterer   // Log filterer for contract events
}

// JBSuckerCaller is an auto generated read-only Go binding around an Ethereum contract.
type JBSuckerCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// JBSuckerTransactor is an auto generated write-only Go binding around an Ethereum contract.
type JBSuckerTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// JBSuckerFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type JBSuckerFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// JBSuckerSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type JBSuckerSession struct {
	Contract     *JBSucker         // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// JBSuckerCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type JBSuckerCallerSession struct {
	Contract *JBSuckerCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts   // Call options to use throughout this session
}

// JBSuckerTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type JBSuckerTransactorSession struct {
	Contract     *JBSuckerTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// JBSuckerRaw is an auto generated low-level Go binding around an Ethereum contract.
type JBSuckerRaw struct {
	Contract *JBSucker // Generic contract binding to access the raw methods on
}

// JBSuckerCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type JBSuckerCallerRaw struct {
	Contract *JBSuckerCaller // Generic read-only contract binding to access the raw methods on
}

// JBSuckerTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type JBSuckerTransactorRaw struct {
	Contract *JBSuckerTransactor // Generic write-only contract binding to access the raw methods on
}

// NewJBSucker creates a new instance of JBSucker, bound to a specific deployed contract.
func NewJBSucker(address common.Address, backend bind.ContractBackend) (*JBSucker, error) {
	contract, err := bindJBSucker(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &JBSucker{JBSuckerCaller: JBSuckerCaller{contract: contract}, JBSuckerTransactor: JBSuckerTransactor{contract: contract}, JBSuckerFilterer: JBSuckerFilterer{contract: contract}}, nil
}

// NewJBSuckerCaller creates a new read-only instance of JBSucker, bound to a specific deployed contract.
func NewJBSuckerCaller(address common.Address, caller bind.ContractCaller) (*JBSuckerCaller, error) {
	contract, err := bindJBSucker(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &JBSuckerCaller{contract: contract}, nil
}

// NewJBSuckerTransactor creates a new write-only instance of JBSucker, bound to a specific deployed contract.
func NewJBSuckerTransactor(address common.Address, transactor bind.ContractTransactor) (*JBSuckerTransactor, error) {
	contract, err := bindJBSucker(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &JBSuckerTransactor{contract: contract}, nil
}

// NewJBSuckerFilterer creates a new log filterer instance of JBSucker, bound to a specific deployed contract.
func NewJBSuckerFilterer(address common.Address, filterer bind.ContractFilterer) (*JBSuckerFilterer, error) {
	contract, err := bindJBSucker(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &JBSuckerFilterer{contract: contract}, nil
}

// bindJBSucker binds a generic wrapper to an already deployed contract.
func bindJBSucker(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := JBSuckerMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_JBSucker *JBSuckerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _JBSucker.Contract.JBSuckerCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_JBSucker *JBSuckerRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _JBSucker.Contract.JBSuckerTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_JBSucker *JBSuckerRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _JBSucker.Contract.JBSuckerTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_JBSucker *JBSuckerCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _JBSucker.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_JBSucker *JBSuckerTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _JBSucker.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_JBSucker *JBSuckerTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _JBSucker.Contract.contract.Transact(opts, method, params...)
}

// ADDTOBALANCEMODE is a free data retrieval call binding the contract method 0xc6525f2c.
//
// Solidity: function ADD_TO_BALANCE_MODE() view returns(uint8)
func (_JBSucker *JBSuckerCaller) ADDTOBALANCEMODE(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _JBSucker.contract.Call(opts, &out, "ADD_TO_BALANCE_MODE")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// ADDTOBALANCEMODE is a free data retrieval call binding the contract method 0xc6525f2c.
//
// Solidity: function ADD_TO_BALANCE_MODE() view returns(uint8)
func (_JBSucker *JBSuckerSession) ADDTOBALANCEMODE() (uint8, error) {
	return _JBSucker.Contract.ADDTOBALANCEMODE(&_JBSucker.CallOpts)
}

// ADDTOBALANCEMODE is a free data retrieval call binding the contract method 0xc6525f2c.
//
// Solidity: function ADD_TO_BALANCE_MODE() view returns(uint8)
func (_JBSucker *JBSuckerCallerSession) ADDTOBALANCEMODE() (uint8, error) {
	return _JBSucker.Contract.ADDTOBALANCEMODE(&_JBSucker.CallOpts)
}

// DEPLOYER is a free data retrieval call binding the contract method 0xc1b8411a.
//
// Solidity: function DEPLOYER() view returns(address)
func (_JBSucker *JBSuckerCaller) DEPLOYER(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _JBSucker.contract.Call(opts, &out, "DEPLOYER")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// DEPLOYER is a free data retrieval call binding the contract method 0xc1b8411a.
//
// Solidity: function DEPLOYER() view returns(address)
func (_JBSucker *JBSuckerSession) DEPLOYER() (common.Address, error) {
	return _JBSucker.Contract.DEPLOYER(&_JBSucker.CallOpts)
}

// DEPLOYER is a free data retrieval call binding the contract method 0xc1b8411a.
//
// Solidity: function DEPLOYER() view returns(address)
func (_JBSucker *JBSuckerCallerSession) DEPLOYER() (common.Address, error) {
	return _JBSucker.Contract.DEPLOYER(&_JBSucker.CallOpts)
}

// DIRECTORY is a free data retrieval call binding the contract method 0x88bc2ef3.
//
// Solidity: function DIRECTORY() view returns(address)
func (_JBSucker *JBSuckerCaller) DIRECTORY(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _JBSucker.contract.Call(opts, &out, "DIRECTORY")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// DIRECTORY is a free data retrieval call binding the contract method 0x88bc2ef3.
//
// Solidity: function DIRECTORY() view returns(address)
func (_JBSucker *JBSuckerSession) DIRECTORY() (common.Address, error) {
	return _JBSucker.Contract.DIRECTORY(&_JBSucker.CallOpts)
}

// DIRECTORY is a free data retrieval call binding the contract method 0x88bc2ef3.
//
// Solidity: function DIRECTORY() view returns(address)
func (_JBSucker *JBSuckerCallerSession) DIRECTORY() (common.Address, error) {
	return _JBSucker.Contract.DIRECTORY(&_JBSucker.CallOpts)
}

// MESSENGERBASEGASLIMIT is a free data retrieval call binding the contract method 0x92ee1ea7.
//
// Solidity: function MESSENGER_BASE_GAS_LIMIT() view returns(uint32)
func (_JBSucker *JBSuckerCaller) MESSENGERBASEGASLIMIT(opts *bind.CallOpts) (uint32, error) {
	var out []interface{}
	err := _JBSucker.contract.Call(opts, &out, "MESSENGER_BASE_GAS_LIMIT")

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

// MESSENGERBASEGASLIMIT is a free data retrieval call binding the contract method 0x92ee1ea7.
//
// Solidity: function MESSENGER_BASE_GAS_LIMIT() view returns(uint32)
func (_JBSucker *JBSuckerSession) MESSENGERBASEGASLIMIT() (uint32, error) {
	return _JBSucker.Contract.MESSENGERBASEGASLIMIT(&_JBSucker.CallOpts)
}

// MESSENGERBASEGASLIMIT is a free data retrieval call binding the contract method 0x92ee1ea7.
//
// Solidity: function MESSENGER_BASE_GAS_LIMIT() view returns(uint32)
func (_JBSucker *JBSuckerCallerSession) MESSENGERBASEGASLIMIT() (uint32, error) {
	return _JBSucker.Contract.MESSENGERBASEGASLIMIT(&_JBSucker.CallOpts)
}

// MESSENGERERC20MINGASLIMIT is a free data retrieval call binding the contract method 0x3a5de07d.
//
// Solidity: function MESSENGER_ERC20_MIN_GAS_LIMIT() view returns(uint32)
func (_JBSucker *JBSuckerCaller) MESSENGERERC20MINGASLIMIT(opts *bind.CallOpts) (uint32, error) {
	var out []interface{}
	err := _JBSucker.contract.Call(opts, &out, "MESSENGER_ERC20_MIN_GAS_LIMIT")

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

// MESSENGERERC20MINGASLIMIT is a free data retrieval call binding the contract method 0x3a5de07d.
//
// Solidity: function MESSENGER_ERC20_MIN_GAS_LIMIT() view returns(uint32)
func (_JBSucker *JBSuckerSession) MESSENGERERC20MINGASLIMIT() (uint32, error) {
	return _JBSucker.Contract.MESSENGERERC20MINGASLIMIT(&_JBSucker.CallOpts)
}

// MESSENGERERC20MINGASLIMIT is a free data retrieval call binding the contract method 0x3a5de07d.
//
// Solidity: function MESSENGER_ERC20_MIN_GAS_LIMIT() view returns(uint32)
func (_JBSucker *JBSuckerCallerSession) MESSENGERERC20MINGASLIMIT() (uint32, error) {
	return _JBSucker.Contract.MESSENGERERC20MINGASLIMIT(&_JBSucker.CallOpts)
}

// PERMISSIONS is a free data retrieval call binding the contract method 0xf434c914.
//
// Solidity: function PERMISSIONS() view returns(address)
func (_JBSucker *JBSuckerCaller) PERMISSIONS(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _JBSucker.contract.Call(opts, &out, "PERMISSIONS")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// PERMISSIONS is a free data retrieval call binding the contract method 0xf434c914.
//
// Solidity: function PERMISSIONS() view returns(address)
func (_JBSucker *JBSuckerSession) PERMISSIONS() (common.Address, error) {
	return _JBSucker.Contract.PERMISSIONS(&_JBSucker.CallOpts)
}

// PERMISSIONS is a free data retrieval call binding the contract method 0xf434c914.
//
// Solidity: function PERMISSIONS() view returns(address)
func (_JBSucker *JBSuckerCallerSession) PERMISSIONS() (common.Address, error) {
	return _JBSucker.Contract.PERMISSIONS(&_JBSucker.CallOpts)
}

// TOKENS is a free data retrieval call binding the contract method 0x1d831d5c.
//
// Solidity: function TOKENS() view returns(address)
func (_JBSucker *JBSuckerCaller) TOKENS(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _JBSucker.contract.Call(opts, &out, "TOKENS")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// TOKENS is a free data retrieval call binding the contract method 0x1d831d5c.
//
// Solidity: function TOKENS() view returns(address)
func (_JBSucker *JBSuckerSession) TOKENS() (common.Address, error) {
	return _JBSucker.Contract.TOKENS(&_JBSucker.CallOpts)
}

// TOKENS is a free data retrieval call binding the contract method 0x1d831d5c.
//
// Solidity: function TOKENS() view returns(address)
func (_JBSucker *JBSuckerCallerSession) TOKENS() (common.Address, error) {
	return _JBSucker.Contract.TOKENS(&_JBSucker.CallOpts)
}

// AmountToAddToBalanceOf is a free data retrieval call binding the contract method 0x9e6d7d43.
//
// Solidity: function amountToAddToBalanceOf(address token) view returns(uint256)
func (_JBSucker *JBSuckerCaller) AmountToAddToBalanceOf(opts *bind.CallOpts, token common.Address) (*big.Int, error) {
	var out []interface{}
	err := _JBSucker.contract.Call(opts, &out, "amountToAddToBalanceOf", token)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// AmountToAddToBalanceOf is a free data retrieval call binding the contract method 0x9e6d7d43.
//
// Solidity: function amountToAddToBalanceOf(address token) view returns(uint256)
func (_JBSucker *JBSuckerSession) AmountToAddToBalanceOf(token common.Address) (*big.Int, error) {
	return _JBSucker.Contract.AmountToAddToBalanceOf(&_JBSucker.CallOpts, token)
}

// AmountToAddToBalanceOf is a free data retrieval call binding the contract method 0x9e6d7d43.
//
// Solidity: function amountToAddToBalanceOf(address token) view returns(uint256)
func (_JBSucker *JBSuckerCallerSession) AmountToAddToBalanceOf(token common.Address) (*big.Int, error) {
	return _JBSucker.Contract.AmountToAddToBalanceOf(&_JBSucker.CallOpts, token)
}

// InboxOf is a free data retrieval call binding the contract method 0x6d9e384b.
//
// Solidity: function inboxOf(address token) view returns((uint64,bytes32))
func (_JBSucker *JBSuckerCaller) InboxOf(opts *bind.CallOpts, token common.Address) (JBInboxTreeRoot, error) {
	var out []interface{}
	err := _JBSucker.contract.Call(opts, &out, "inboxOf", token)

	if err != nil {
		return *new(JBInboxTreeRoot), err
	}

	out0 := *abi.ConvertType(out[0], new(JBInboxTreeRoot)).(*JBInboxTreeRoot)

	return out0, err

}

// InboxOf is a free data retrieval call binding the contract method 0x6d9e384b.
//
// Solidity: function inboxOf(address token) view returns((uint64,bytes32))
func (_JBSucker *JBSuckerSession) InboxOf(token common.Address) (JBInboxTreeRoot, error) {
	return _JBSucker.Contract.InboxOf(&_JBSucker.CallOpts, token)
}

// InboxOf is a free data retrieval call binding the contract method 0x6d9e384b.
//
// Solidity: function inboxOf(address token) view returns((uint64,bytes32))
func (_JBSucker *JBSuckerCallerSession) InboxOf(token common.Address) (JBInboxTreeRoot, error) {
	return _JBSucker.Contract.InboxOf(&_JBSucker.CallOpts, token)
}

// IsMapped is a free data retrieval call binding the contract method 0xbf557aee.
//
// Solidity: function isMapped(address token) view returns(bool)
func (_JBSucker *JBSuckerCaller) IsMapped(opts *bind.CallOpts, token common.Address) (bool, error) {
	var out []interface{}
	err := _JBSucker.contract.Call(opts, &out, "isMapped", token)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsMapped is a free data retrieval call binding the contract method 0xbf557aee.
//
// Solidity: function isMapped(address token) view returns(bool)
func (_JBSucker *JBSuckerSession) IsMapped(token common.Address) (bool, error) {
	return _JBSucker.Contract.IsMapped(&_JBSucker.CallOpts, token)
}

// IsMapped is a free data retrieval call binding the contract method 0xbf557aee.
//
// Solidity: function isMapped(address token) view returns(bool)
func (_JBSucker *JBSuckerCallerSession) IsMapped(token common.Address) (bool, error) {
	return _JBSucker.Contract.IsMapped(&_JBSucker.CallOpts, token)
}

// IsTrustedForwarder is a free data retrieval call binding the contract method 0x572b6c05.
//
// Solidity: function isTrustedForwarder(address forwarder) view returns(bool)
func (_JBSucker *JBSuckerCaller) IsTrustedForwarder(opts *bind.CallOpts, forwarder common.Address) (bool, error) {
	var out []interface{}
	err := _JBSucker.contract.Call(opts, &out, "isTrustedForwarder", forwarder)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsTrustedForwarder is a free data retrieval call binding the contract method 0x572b6c05.
//
// Solidity: function isTrustedForwarder(address forwarder) view returns(bool)
func (_JBSucker *JBSuckerSession) IsTrustedForwarder(forwarder common.Address) (bool, error) {
	return _JBSucker.Contract.IsTrustedForwarder(&_JBSucker.CallOpts, forwarder)
}

// IsTrustedForwarder is a free data retrieval call binding the contract method 0x572b6c05.
//
// Solidity: function isTrustedForwarder(address forwarder) view returns(bool)
func (_JBSucker *JBSuckerCallerSession) IsTrustedForwarder(forwarder common.Address) (bool, error) {
	return _JBSucker.Contract.IsTrustedForwarder(&_JBSucker.CallOpts, forwarder)
}

// OutboxOf is a free data retrieval call binding the contract method 0x802c8fa0.
//
// Solidity: function outboxOf(address token) view returns((uint64,uint256,(bytes32[32],uint256),uint256))
func (_JBSucker *JBSuckerCaller) OutboxOf(opts *bind.CallOpts, token common.Address) (JBOutboxTree, error) {
	var out []interface{}
	err := _JBSucker.contract.Call(opts, &out, "outboxOf", token)

	if err != nil {
		return *new(JBOutboxTree), err
	}

	out0 := *abi.ConvertType(out[0], new(JBOutboxTree)).(*JBOutboxTree)

	return out0, err

}

// OutboxOf is a free data retrieval call binding the contract method 0x802c8fa0.
//
// Solidity: function outboxOf(address token) view returns((uint64,uint256,(bytes32[32],uint256),uint256))
func (_JBSucker *JBSuckerSession) OutboxOf(token common.Address) (JBOutboxTree, error) {
	return _JBSucker.Contract.OutboxOf(&_JBSucker.CallOpts, token)
}

// OutboxOf is a free data retrieval call binding the contract method 0x802c8fa0.
//
// Solidity: function outboxOf(address token) view returns((uint64,uint256,(bytes32[32],uint256),uint256))
func (_JBSucker *JBSuckerCallerSession) OutboxOf(token common.Address) (JBOutboxTree, error) {
	return _JBSucker.Contract.OutboxOf(&_JBSucker.CallOpts, token)
}

// Peer is a free data retrieval call binding the contract method 0x11cda415.
//
// Solidity: function peer() view returns(address)
func (_JBSucker *JBSuckerCaller) Peer(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _JBSucker.contract.Call(opts, &out, "peer")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Peer is a free data retrieval call binding the contract method 0x11cda415.
//
// Solidity: function peer() view returns(address)
func (_JBSucker *JBSuckerSession) Peer() (common.Address, error) {
	return _JBSucker.Contract.Peer(&_JBSucker.CallOpts)
}

// Peer is a free data retrieval call binding the contract method 0x11cda415.
//
// Solidity: function peer() view returns(address)
func (_JBSucker *JBSuckerCallerSession) Peer() (common.Address, error) {
	return _JBSucker.Contract.Peer(&_JBSucker.CallOpts)
}

// PeerChainId is a free data retrieval call binding the contract method 0xcdfceeba.
//
// Solidity: function peerChainId() view returns(uint256)
func (_JBSucker *JBSuckerCaller) PeerChainId(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _JBSucker.contract.Call(opts, &out, "peerChainId")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PeerChainId is a free data retrieval call binding the contract method 0xcdfceeba.
//
// Solidity: function peerChainId() view returns(uint256)
func (_JBSucker *JBSuckerSession) PeerChainId() (*big.Int, error) {
	return _JBSucker.Contract.PeerChainId(&_JBSucker.CallOpts)
}

// PeerChainId is a free data retrieval call binding the contract method 0xcdfceeba.
//
// Solidity: function peerChainId() view returns(uint256)
func (_JBSucker *JBSuckerCallerSession) PeerChainId() (*big.Int, error) {
	return _JBSucker.Contract.PeerChainId(&_JBSucker.CallOpts)
}

// ProjectId is a free data retrieval call binding the contract method 0x3fafa127.
//
// Solidity: function projectId() view returns(uint256)
func (_JBSucker *JBSuckerCaller) ProjectId(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _JBSucker.contract.Call(opts, &out, "projectId")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ProjectId is a free data retrieval call binding the contract method 0x3fafa127.
//
// Solidity: function projectId() view returns(uint256)
func (_JBSucker *JBSuckerSession) ProjectId() (*big.Int, error) {
	return _JBSucker.Contract.ProjectId(&_JBSucker.CallOpts)
}

// ProjectId is a free data retrieval call binding the contract method 0x3fafa127.
//
// Solidity: function projectId() view returns(uint256)
func (_JBSucker *JBSuckerCallerSession) ProjectId() (*big.Int, error) {
	return _JBSucker.Contract.ProjectId(&_JBSucker.CallOpts)
}

// RemoteTokenFor is a free data retrieval call binding the contract method 0x7de6fba2.
//
// Solidity: function remoteTokenFor(address token) view returns((bool,bool,uint32,address,uint256))
func (_JBSucker *JBSuckerCaller) RemoteTokenFor(opts *bind.CallOpts, token common.Address) (JBRemoteToken, error) {
	var out []interface{}
	err := _JBSucker.contract.Call(opts, &out, "remoteTokenFor", token)

	if err != nil {
		return *new(JBRemoteToken), err
	}

	out0 := *abi.ConvertType(out[0], new(JBRemoteToken)).(*JBRemoteToken)

	return out0, err

}

// RemoteTokenFor is a free data retrieval call binding the contract method 0x7de6fba2.
//
// Solidity: function remoteTokenFor(address token) view returns((bool,bool,uint32,address,uint256))
func (_JBSucker *JBSuckerSession) RemoteTokenFor(token common.Address) (JBRemoteToken, error) {
	return _JBSucker.Contract.RemoteTokenFor(&_JBSucker.CallOpts, token)
}

// RemoteTokenFor is a free data retrieval call binding the contract method 0x7de6fba2.
//
// Solidity: function remoteTokenFor(address token) view returns((bool,bool,uint32,address,uint256))
func (_JBSucker *JBSuckerCallerSession) RemoteTokenFor(token common.Address) (JBRemoteToken, error) {
	return _JBSucker.Contract.RemoteTokenFor(&_JBSucker.CallOpts, token)
}

// State is a free data retrieval call binding the contract method 0xc19d93fb.
//
// Solidity: function state() view returns(uint8)
func (_JBSucker *JBSuckerCaller) State(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _JBSucker.contract.Call(opts, &out, "state")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// State is a free data retrieval call binding the contract method 0xc19d93fb.
//
// Solidity: function state() view returns(uint8)
func (_JBSucker *JBSuckerSession) State() (uint8, error) {
	return _JBSucker.Contract.State(&_JBSucker.CallOpts)
}

// State is a free data retrieval call binding the contract method 0xc19d93fb.
//
// Solidity: function state() view returns(uint8)
func (_JBSucker *JBSuckerCallerSession) State() (uint8, error) {
	return _JBSucker.Contract.State(&_JBSucker.CallOpts)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_JBSucker *JBSuckerCaller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _JBSucker.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_JBSucker *JBSuckerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _JBSucker.Contract.SupportsInterface(&_JBSucker.CallOpts, interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_JBSucker *JBSuckerCallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _JBSucker.Contract.SupportsInterface(&_JBSucker.CallOpts, interfaceId)
}

// TrustedForwarder is a free data retrieval call binding the contract method 0x7da0a877.
//
// Solidity: function trustedForwarder() view returns(address)
func (_JBSucker *JBSuckerCaller) TrustedForwarder(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _JBSucker.contract.Call(opts, &out, "trustedForwarder")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// TrustedForwarder is a free data retrieval call binding the contract method 0x7da0a877.
//
// Solidity: function trustedForwarder() view returns(address)
func (_JBSucker *JBSuckerSession) TrustedForwarder() (common.Address, error) {
	return _JBSucker.Contract.TrustedForwarder(&_JBSucker.CallOpts)
}

// TrustedForwarder is a free data retrieval call binding the contract method 0x7da0a877.
//
// Solidity: function trustedForwarder() view returns(address)
func (_JBSucker *JBSuckerCallerSession) TrustedForwarder() (common.Address, error) {
	return _JBSucker.Contract.TrustedForwarder(&_JBSucker.CallOpts)
}

// AddOutstandingAmountToBalance is a paid mutator transaction binding the contract method 0x69ce2233.
//
// Solidity: function addOutstandingAmountToBalance(address token) returns()
func (_JBSucker *JBSuckerTransactor) AddOutstandingAmountToBalance(opts *bind.TransactOpts, token common.Address) (*types.Transaction, error) {
	return _JBSucker.contract.Transact(opts, "addOutstandingAmountToBalance", token)
}

// AddOutstandingAmountToBalance is a paid mutator transaction binding the contract method 0x69ce2233.
//
// Solidity: function addOutstandingAmountToBalance(address token) returns()
func (_JBSucker *JBSuckerSession) AddOutstandingAmountToBalance(token common.Address) (*types.Transaction, error) {
	return _JBSucker.Contract.AddOutstandingAmountToBalance(&_JBSucker.TransactOpts, token)
}

// AddOutstandingAmountToBalance is a paid mutator transaction binding the contract method 0x69ce2233.
//
// Solidity: function addOutstandingAmountToBalance(address token) returns()
func (_JBSucker *JBSuckerTransactorSession) AddOutstandingAmountToBalance(token common.Address) (*types.Transaction, error) {
	return _JBSucker.Contract.AddOutstandingAmountToBalance(&_JBSucker.TransactOpts, token)
}

// Claim is a paid mutator transaction binding the contract method 0xb59ef50f.
//
// Solidity: function claim((address,(uint256,address,uint256,uint256),bytes32[32]) claimData) returns()
func (_JBSucker *JBSuckerTransactor) Claim(opts *bind.TransactOpts, claimData JBClaim) (*types.Transaction, error) {
	return _JBSucker.contract.Transact(opts, "claim", claimData)
}

// Claim is a paid mutator transaction binding the contract method 0xb59ef50f.
//
// Solidity: function claim((address,(uint256,address,uint256,uint256),bytes32[32]) claimData) returns()
func (_JBSucker *JBSuckerSession) Claim(claimData JBClaim) (*types.Transaction, error) {
	return _JBSucker.Contract.Claim(&_JBSucker.TransactOpts, claimData)
}

// Claim is a paid mutator transaction binding the contract method 0xb59ef50f.
//
// Solidity: function claim((address,(uint256,address,uint256,uint256),bytes32[32]) claimData) returns()
func (_JBSucker *JBSuckerTransactorSession) Claim(claimData JBClaim) (*types.Transaction, error) {
	return _JBSucker.Contract.Claim(&_JBSucker.TransactOpts, claimData)
}

// Claim0 is a paid mutator transaction binding the contract method 0xc1854129.
//
// Solidity: function claim((address,(uint256,address,uint256,uint256),bytes32[32])[] claims) returns()
func (_JBSucker *JBSuckerTransactor) Claim0(opts *bind.TransactOpts, claims []JBClaim) (*types.Transaction, error) {
	return _JBSucker.contract.Transact(opts, "claim0", claims)
}

// Claim0 is a paid mutator transaction binding the contract method 0xc1854129.
//
// Solidity: function claim((address,(uint256,address,uint256,uint256),bytes32[32])[] claims) returns()
func (_JBSucker *JBSuckerSession) Claim0(claims []JBClaim) (*types.Transaction, error) {
	return _JBSucker.Contract.Claim0(&_JBSucker.TransactOpts, claims)
}

// Claim0 is a paid mutator transaction binding the contract method 0xc1854129.
//
// Solidity: function claim((address,(uint256,address,uint256,uint256),bytes32[32])[] claims) returns()
func (_JBSucker *JBSuckerTransactorSession) Claim0(claims []JBClaim) (*types.Transaction, error) {
	return _JBSucker.Contract.Claim0(&_JBSucker.TransactOpts, claims)
}

// EnableEmergencyHatchFor is a paid mutator transaction binding the contract method 0x75451559.
//
// Solidity: function enableEmergencyHatchFor(address[] tokens) returns()
func (_JBSucker *JBSuckerTransactor) EnableEmergencyHatchFor(opts *bind.TransactOpts, tokens []common.Address) (*types.Transaction, error) {
	return _JBSucker.contract.Transact(opts, "enableEmergencyHatchFor", tokens)
}

// EnableEmergencyHatchFor is a paid mutator transaction binding the contract method 0x75451559.
//
// Solidity: function enableEmergencyHatchFor(address[] tokens) returns()
func (_JBSucker *JBSuckerSession) EnableEmergencyHatchFor(tokens []common.Address) (*types.Transaction, error) {
	return _JBSucker.Contract.EnableEmergencyHatchFor(&_JBSucker.TransactOpts, tokens)
}

// EnableEmergencyHatchFor is a paid mutator transaction binding the contract method 0x75451559.
//
// Solidity: function enableEmergencyHatchFor(address[] tokens) returns()
func (_JBSucker *JBSuckerTransactorSession) EnableEmergencyHatchFor(tokens []common.Address) (*types.Transaction, error) {
	return _JBSucker.Contract.EnableEmergencyHatchFor(&_JBSucker.TransactOpts, tokens)
}

// ExitThroughEmergencyHatch is a paid mutator transaction binding the contract method 0x0f648214.
//
// Solidity: function exitThroughEmergencyHatch((address,(uint256,address,uint256,uint256),bytes32[32]) claimData) returns()
func (_JBSucker *JBSuckerTransactor) ExitThroughEmergencyHatch(opts *bind.TransactOpts, claimData JBClaim) (*types.Transaction, error) {
	return _JBSucker.contract.Transact(opts, "exitThroughEmergencyHatch", claimData)
}

// ExitThroughEmergencyHatch is a paid mutator transaction binding the contract method 0x0f648214.
//
// Solidity: function exitThroughEmergencyHatch((address,(uint256,address,uint256,uint256),bytes32[32]) claimData) returns()
func (_JBSucker *JBSuckerSession) ExitThroughEmergencyHatch(claimData JBClaim) (*types.Transaction, error) {
	return _JBSucker.Contract.ExitThroughEmergencyHatch(&_JBSucker.TransactOpts, claimData)
}

// ExitThroughEmergencyHatch is a paid mutator transaction binding the contract method 0x0f648214.
//
// Solidity: function exitThroughEmergencyHatch((address,(uint256,address,uint256,uint256),bytes32[32]) claimData) returns()
func (_JBSucker *JBSuckerTransactorSession) ExitThroughEmergencyHatch(claimData JBClaim) (*types.Transaction, error) {
	return _JBSucker.Contract.ExitThroughEmergencyHatch(&_JBSucker.TransactOpts, claimData)
}

// FromRemote is a paid mutator transaction binding the contract method 0x69d31fdd.
//
// Solidity: function fromRemote((address,uint256,(uint64,bytes32)) root) payable returns()
func (_JBSucker *JBSuckerTransactor) FromRemote(opts *bind.TransactOpts, root JBMessageRoot) (*types.Transaction, error) {
	return _JBSucker.contract.Transact(opts, "fromRemote", root)
}

// FromRemote is a paid mutator transaction binding the contract method 0x69d31fdd.
//
// Solidity: function fromRemote((address,uint256,(uint64,bytes32)) root) payable returns()
func (_JBSucker *JBSuckerSession) FromRemote(root JBMessageRoot) (*types.Transaction, error) {
	return _JBSucker.Contract.FromRemote(&_JBSucker.TransactOpts, root)
}

// FromRemote is a paid mutator transaction binding the contract method 0x69d31fdd.
//
// Solidity: function fromRemote((address,uint256,(uint64,bytes32)) root) payable returns()
func (_JBSucker *JBSuckerTransactorSession) FromRemote(root JBMessageRoot) (*types.Transaction, error) {
	return _JBSucker.Contract.FromRemote(&_JBSucker.TransactOpts, root)
}

// Initialize is a paid mutator transaction binding the contract method 0xfe4b84df.
//
// Solidity: function initialize(uint256 __projectId) returns()
func (_JBSucker *JBSuckerTransactor) Initialize(opts *bind.TransactOpts, __projectId *big.Int) (*types.Transaction, error) {
	return _JBSucker.contract.Transact(opts, "initialize", __projectId)
}

// Initialize is a paid mutator transaction binding the contract method 0xfe4b84df.
//
// Solidity: function initialize(uint256 __projectId) returns()
func (_JBSucker *JBSuckerSession) Initialize(__projectId *big.Int) (*types.Transaction, error) {
	return _JBSucker.Contract.Initialize(&_JBSucker.TransactOpts, __projectId)
}

// Initialize is a paid mutator transaction binding the contract method 0xfe4b84df.
//
// Solidity: function initialize(uint256 __projectId) returns()
func (_JBSucker *JBSuckerTransactorSession) Initialize(__projectId *big.Int) (*types.Transaction, error) {
	return _JBSucker.Contract.Initialize(&_JBSucker.TransactOpts, __projectId)
}

// MapToken is a paid mutator transaction binding the contract method 0x38eb2044.
//
// Solidity: function mapToken((address,uint32,address,uint256) map) payable returns()
func (_JBSucker *JBSuckerTransactor) MapToken(opts *bind.TransactOpts, arg0 JBTokenMapping) (*types.Transaction, error) {
	return _JBSucker.contract.Transact(opts, "mapToken", arg0)
}

// MapToken is a paid mutator transaction binding the contract method 0x38eb2044.
//
// Solidity: function mapToken((address,uint32,address,uint256) map) payable returns()
func (_JBSucker *JBSuckerSession) MapToken(arg0 JBTokenMapping) (*types.Transaction, error) {
	return _JBSucker.Contract.MapToken(&_JBSucker.TransactOpts, arg0)
}

// MapToken is a paid mutator transaction binding the contract method 0x38eb2044.
//
// Solidity: function mapToken((address,uint32,address,uint256) map) payable returns()
func (_JBSucker *JBSuckerTransactorSession) MapToken(arg0 JBTokenMapping) (*types.Transaction, error) {
	return _JBSucker.Contract.MapToken(&_JBSucker.TransactOpts, arg0)
}

// MapTokens is a paid mutator transaction binding the contract method 0x2241dc52.
//
// Solidity: function mapTokens((address,uint32,address,uint256)[] maps) payable returns()
func (_JBSucker *JBSuckerTransactor) MapTokens(opts *bind.TransactOpts, maps []JBTokenMapping) (*types.Transaction, error) {
	return _JBSucker.contract.Transact(opts, "mapTokens", maps)
}

// MapTokens is a paid mutator transaction binding the contract method 0x2241dc52.
//
// Solidity: function mapTokens((address,uint32,address,uint256)[] maps) payable returns()
func (_JBSucker *JBSuckerSession) MapTokens(maps []JBTokenMapping) (*types.Transaction, error) {
	return _JBSucker.Contract.MapTokens(&_JBSucker.TransactOpts, maps)
}

// MapTokens is a paid mutator transaction binding the contract method 0x2241dc52.
//
// Solidity: function mapTokens((address,uint32,address,uint256)[] maps) payable returns()
func (_JBSucker *JBSuckerTransactorSession) MapTokens(maps []JBTokenMapping) (*types.Transaction, error) {
	return _JBSucker.Contract.MapTokens(&_JBSucker.TransactOpts, maps)
}

// Prepare is a paid mutator transaction binding the contract method 0x551ebbf0.
//
// Solidity: function prepare(uint256 projectTokenCount, address beneficiary, uint256 minTokensReclaimed, address token) returns()
func (_JBSucker *JBSuckerTransactor) Prepare(opts *bind.TransactOpts, projectTokenCount *big.Int, beneficiary common.Address, minTokensReclaimed *big.Int, token common.Address) (*types.Transaction, error) {
	return _JBSucker.contract.Transact(opts, "prepare", projectTokenCount, beneficiary, minTokensReclaimed, token)
}

// Prepare is a paid mutator transaction binding the contract method 0x551ebbf0.
//
// Solidity: function prepare(uint256 projectTokenCount, address beneficiary, uint256 minTokensReclaimed, address token) returns()
func (_JBSucker *JBSuckerSession) Prepare(projectTokenCount *big.Int, beneficiary common.Address, minTokensReclaimed *big.Int, token common.Address) (*types.Transaction, error) {
	return _JBSucker.Contract.Prepare(&_JBSucker.TransactOpts, projectTokenCount, beneficiary, minTokensReclaimed, token)
}

// Prepare is a paid mutator transaction binding the contract method 0x551ebbf0.
//
// Solidity: function prepare(uint256 projectTokenCount, address beneficiary, uint256 minTokensReclaimed, address token) returns()
func (_JBSucker *JBSuckerTransactorSession) Prepare(projectTokenCount *big.Int, beneficiary common.Address, minTokensReclaimed *big.Int, token common.Address) (*types.Transaction, error) {
	return _JBSucker.Contract.Prepare(&_JBSucker.TransactOpts, projectTokenCount, beneficiary, minTokensReclaimed, token)
}

// SetDeprecation is a paid mutator transaction binding the contract method 0x6284386a.
//
// Solidity: function setDeprecation(uint40 timestamp) returns()
func (_JBSucker *JBSuckerTransactor) SetDeprecation(opts *bind.TransactOpts, timestamp *big.Int) (*types.Transaction, error) {
	return _JBSucker.contract.Transact(opts, "setDeprecation", timestamp)
}

// SetDeprecation is a paid mutator transaction binding the contract method 0x6284386a.
//
// Solidity: function setDeprecation(uint40 timestamp) returns()
func (_JBSucker *JBSuckerSession) SetDeprecation(timestamp *big.Int) (*types.Transaction, error) {
	return _JBSucker.Contract.SetDeprecation(&_JBSucker.TransactOpts, timestamp)
}

// SetDeprecation is a paid mutator transaction binding the contract method 0x6284386a.
//
// Solidity: function setDeprecation(uint40 timestamp) returns()
func (_JBSucker *JBSuckerTransactorSession) SetDeprecation(timestamp *big.Int) (*types.Transaction, error) {
	return _JBSucker.Contract.SetDeprecation(&_JBSucker.TransactOpts, timestamp)
}

// ToRemote is a paid mutator transaction binding the contract method 0xb71c1179.
//
// Solidity: function toRemote(address token) payable returns()
func (_JBSucker *JBSuckerTransactor) ToRemote(opts *bind.TransactOpts, token common.Address) (*types.Transaction, error) {
	return _JBSucker.contract.Transact(opts, "toRemote", token)
}

// ToRemote is a paid mutator transaction binding the contract method 0xb71c1179.
//
// Solidity: function toRemote(address token) payable returns()
func (_JBSucker *JBSuckerSession) ToRemote(token common.Address) (*types.Transaction, error) {
	return _JBSucker.Contract.ToRemote(&_JBSucker.TransactOpts, token)
}

// ToRemote is a paid mutator transaction binding the contract method 0xb71c1179.
//
// Solidity: function toRemote(address token) payable returns()
func (_JBSucker *JBSuckerTransactorSession) ToRemote(token common.Address) (*types.Transaction, error) {
	return _JBSucker.Contract.ToRemote(&_JBSucker.TransactOpts, token)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_JBSucker *JBSuckerTransactor) Receive(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _JBSucker.contract.RawTransact(opts, nil) // calldata is disallowed for receive function
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_JBSucker *JBSuckerSession) Receive() (*types.Transaction, error) {
	return _JBSucker.Contract.Receive(&_JBSucker.TransactOpts)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_JBSucker *JBSuckerTransactorSession) Receive() (*types.Transaction, error) {
	return _JBSucker.Contract.Receive(&_JBSucker.TransactOpts)
}

// JBSuckerClaimedIterator is returned from FilterClaimed and is used to iterate over the raw logs and unpacked data for Claimed events raised by the JBSucker contract.
type JBSuckerClaimedIterator struct {
	Event *JBSuckerClaimed // Event containing the contract specifics and raw log

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
func (it *JBSuckerClaimedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(JBSuckerClaimed)
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
		it.Event = new(JBSuckerClaimed)
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
func (it *JBSuckerClaimedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *JBSuckerClaimedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// JBSuckerClaimed represents a Claimed event raised by the JBSucker contract.
type JBSuckerClaimed struct {
	Beneficiary         common.Address
	Token               common.Address
	ProjectTokenCount   *big.Int
	TerminalTokenAmount *big.Int
	Index               *big.Int
	AutoAddedToBalance  bool
	Caller              common.Address
	Raw                 types.Log // Blockchain specific contextual infos
}

// FilterClaimed is a free log retrieval operation binding the contract event 0xaaca6f6829d340365b5e6599e0d513cbb4602b94f58ccf16ac77d8d3ae7139ef.
//
// Solidity: event Claimed(address beneficiary, address token, uint256 projectTokenCount, uint256 terminalTokenAmount, uint256 index, bool autoAddedToBalance, address caller)
func (_JBSucker *JBSuckerFilterer) FilterClaimed(opts *bind.FilterOpts) (*JBSuckerClaimedIterator, error) {

	logs, sub, err := _JBSucker.contract.FilterLogs(opts, "Claimed")
	if err != nil {
		return nil, err
	}
	return &JBSuckerClaimedIterator{contract: _JBSucker.contract, event: "Claimed", logs: logs, sub: sub}, nil
}

// WatchClaimed is a free log subscription operation binding the contract event 0xaaca6f6829d340365b5e6599e0d513cbb4602b94f58ccf16ac77d8d3ae7139ef.
//
// Solidity: event Claimed(address beneficiary, address token, uint256 projectTokenCount, uint256 terminalTokenAmount, uint256 index, bool autoAddedToBalance, address caller)
func (_JBSucker *JBSuckerFilterer) WatchClaimed(opts *bind.WatchOpts, sink chan<- *JBSuckerClaimed) (event.Subscription, error) {

	logs, sub, err := _JBSucker.contract.WatchLogs(opts, "Claimed")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(JBSuckerClaimed)
				if err := _JBSucker.contract.UnpackLog(event, "Claimed", log); err != nil {
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

// ParseClaimed is a log parse operation binding the contract event 0xaaca6f6829d340365b5e6599e0d513cbb4602b94f58ccf16ac77d8d3ae7139ef.
//
// Solidity: event Claimed(address beneficiary, address token, uint256 projectTokenCount, uint256 terminalTokenAmount, uint256 index, bool autoAddedToBalance, address caller)
func (_JBSucker *JBSuckerFilterer) ParseClaimed(log types.Log) (*JBSuckerClaimed, error) {
	event := new(JBSuckerClaimed)
	if err := _JBSucker.contract.UnpackLog(event, "Claimed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// JBSuckerDeprecationTimeUpdatedIterator is returned from FilterDeprecationTimeUpdated and is used to iterate over the raw logs and unpacked data for DeprecationTimeUpdated events raised by the JBSucker contract.
type JBSuckerDeprecationTimeUpdatedIterator struct {
	Event *JBSuckerDeprecationTimeUpdated // Event containing the contract specifics and raw log

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
func (it *JBSuckerDeprecationTimeUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(JBSuckerDeprecationTimeUpdated)
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
		it.Event = new(JBSuckerDeprecationTimeUpdated)
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
func (it *JBSuckerDeprecationTimeUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *JBSuckerDeprecationTimeUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// JBSuckerDeprecationTimeUpdated represents a DeprecationTimeUpdated event raised by the JBSucker contract.
type JBSuckerDeprecationTimeUpdated struct {
	Timestamp *big.Int
	Caller    common.Address
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterDeprecationTimeUpdated is a free log retrieval operation binding the contract event 0x455efa76af2a8686dc0dd3041bda89ece5ba4691277d525e2161003bbca8b245.
//
// Solidity: event DeprecationTimeUpdated(uint40 timestamp, address caller)
func (_JBSucker *JBSuckerFilterer) FilterDeprecationTimeUpdated(opts *bind.FilterOpts) (*JBSuckerDeprecationTimeUpdatedIterator, error) {

	logs, sub, err := _JBSucker.contract.FilterLogs(opts, "DeprecationTimeUpdated")
	if err != nil {
		return nil, err
	}
	return &JBSuckerDeprecationTimeUpdatedIterator{contract: _JBSucker.contract, event: "DeprecationTimeUpdated", logs: logs, sub: sub}, nil
}

// WatchDeprecationTimeUpdated is a free log subscription operation binding the contract event 0x455efa76af2a8686dc0dd3041bda89ece5ba4691277d525e2161003bbca8b245.
//
// Solidity: event DeprecationTimeUpdated(uint40 timestamp, address caller)
func (_JBSucker *JBSuckerFilterer) WatchDeprecationTimeUpdated(opts *bind.WatchOpts, sink chan<- *JBSuckerDeprecationTimeUpdated) (event.Subscription, error) {

	logs, sub, err := _JBSucker.contract.WatchLogs(opts, "DeprecationTimeUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(JBSuckerDeprecationTimeUpdated)
				if err := _JBSucker.contract.UnpackLog(event, "DeprecationTimeUpdated", log); err != nil {
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

// ParseDeprecationTimeUpdated is a log parse operation binding the contract event 0x455efa76af2a8686dc0dd3041bda89ece5ba4691277d525e2161003bbca8b245.
//
// Solidity: event DeprecationTimeUpdated(uint40 timestamp, address caller)
func (_JBSucker *JBSuckerFilterer) ParseDeprecationTimeUpdated(log types.Log) (*JBSuckerDeprecationTimeUpdated, error) {
	event := new(JBSuckerDeprecationTimeUpdated)
	if err := _JBSucker.contract.UnpackLog(event, "DeprecationTimeUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// JBSuckerEmergencyHatchOpenedIterator is returned from FilterEmergencyHatchOpened and is used to iterate over the raw logs and unpacked data for EmergencyHatchOpened events raised by the JBSucker contract.
type JBSuckerEmergencyHatchOpenedIterator struct {
	Event *JBSuckerEmergencyHatchOpened // Event containing the contract specifics and raw log

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
func (it *JBSuckerEmergencyHatchOpenedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(JBSuckerEmergencyHatchOpened)
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
		it.Event = new(JBSuckerEmergencyHatchOpened)
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
func (it *JBSuckerEmergencyHatchOpenedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *JBSuckerEmergencyHatchOpenedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// JBSuckerEmergencyHatchOpened represents a EmergencyHatchOpened event raised by the JBSucker contract.
type JBSuckerEmergencyHatchOpened struct {
	Tokens []common.Address
	Caller common.Address
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterEmergencyHatchOpened is a free log retrieval operation binding the contract event 0x8021bce60b20a9afb57085869a3552e0ee5b272b868df894fb30c5e90e74c072.
//
// Solidity: event EmergencyHatchOpened(address[] tokens, address caller)
func (_JBSucker *JBSuckerFilterer) FilterEmergencyHatchOpened(opts *bind.FilterOpts) (*JBSuckerEmergencyHatchOpenedIterator, error) {

	logs, sub, err := _JBSucker.contract.FilterLogs(opts, "EmergencyHatchOpened")
	if err != nil {
		return nil, err
	}
	return &JBSuckerEmergencyHatchOpenedIterator{contract: _JBSucker.contract, event: "EmergencyHatchOpened", logs: logs, sub: sub}, nil
}

// WatchEmergencyHatchOpened is a free log subscription operation binding the contract event 0x8021bce60b20a9afb57085869a3552e0ee5b272b868df894fb30c5e90e74c072.
//
// Solidity: event EmergencyHatchOpened(address[] tokens, address caller)
func (_JBSucker *JBSuckerFilterer) WatchEmergencyHatchOpened(opts *bind.WatchOpts, sink chan<- *JBSuckerEmergencyHatchOpened) (event.Subscription, error) {

	logs, sub, err := _JBSucker.contract.WatchLogs(opts, "EmergencyHatchOpened")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(JBSuckerEmergencyHatchOpened)
				if err := _JBSucker.contract.UnpackLog(event, "EmergencyHatchOpened", log); err != nil {
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

// ParseEmergencyHatchOpened is a log parse operation binding the contract event 0x8021bce60b20a9afb57085869a3552e0ee5b272b868df894fb30c5e90e74c072.
//
// Solidity: event EmergencyHatchOpened(address[] tokens, address caller)
func (_JBSucker *JBSuckerFilterer) ParseEmergencyHatchOpened(log types.Log) (*JBSuckerEmergencyHatchOpened, error) {
	event := new(JBSuckerEmergencyHatchOpened)
	if err := _JBSucker.contract.UnpackLog(event, "EmergencyHatchOpened", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// JBSuckerInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the JBSucker contract.
type JBSuckerInitializedIterator struct {
	Event *JBSuckerInitialized // Event containing the contract specifics and raw log

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
func (it *JBSuckerInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(JBSuckerInitialized)
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
		it.Event = new(JBSuckerInitialized)
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
func (it *JBSuckerInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *JBSuckerInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// JBSuckerInitialized represents a Initialized event raised by the JBSucker contract.
type JBSuckerInitialized struct {
	Version uint64
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_JBSucker *JBSuckerFilterer) FilterInitialized(opts *bind.FilterOpts) (*JBSuckerInitializedIterator, error) {

	logs, sub, err := _JBSucker.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &JBSuckerInitializedIterator{contract: _JBSucker.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_JBSucker *JBSuckerFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *JBSuckerInitialized) (event.Subscription, error) {

	logs, sub, err := _JBSucker.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(JBSuckerInitialized)
				if err := _JBSucker.contract.UnpackLog(event, "Initialized", log); err != nil {
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
func (_JBSucker *JBSuckerFilterer) ParseInitialized(log types.Log) (*JBSuckerInitialized, error) {
	event := new(JBSuckerInitialized)
	if err := _JBSucker.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// JBSuckerInsertToOutboxTreeIterator is returned from FilterInsertToOutboxTree and is used to iterate over the raw logs and unpacked data for InsertToOutboxTree events raised by the JBSucker contract.
type JBSuckerInsertToOutboxTreeIterator struct {
	Event *JBSuckerInsertToOutboxTree // Event containing the contract specifics and raw log

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
func (it *JBSuckerInsertToOutboxTreeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(JBSuckerInsertToOutboxTree)
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
		it.Event = new(JBSuckerInsertToOutboxTree)
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
func (it *JBSuckerInsertToOutboxTreeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *JBSuckerInsertToOutboxTreeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// JBSuckerInsertToOutboxTree represents a InsertToOutboxTree event raised by the JBSucker contract.
type JBSuckerInsertToOutboxTree struct {
	Beneficiary         common.Address
	Token               common.Address
	Hashed              [32]byte
	Index               *big.Int
	Root                [32]byte
	ProjectTokenCount   *big.Int
	TerminalTokenAmount *big.Int
	Caller              common.Address
	Raw                 types.Log // Blockchain specific contextual infos
}

// FilterInsertToOutboxTree is a free log retrieval operation binding the contract event 0x584a41ff082fad99bb6e50c5767c347c762057fe5726539f0a8de49d84549c23.
//
// Solidity: event InsertToOutboxTree(address indexed beneficiary, address indexed token, bytes32 hashed, uint256 index, bytes32 root, uint256 projectTokenCount, uint256 terminalTokenAmount, address caller)
func (_JBSucker *JBSuckerFilterer) FilterInsertToOutboxTree(opts *bind.FilterOpts, beneficiary []common.Address, token []common.Address) (*JBSuckerInsertToOutboxTreeIterator, error) {

	var beneficiaryRule []interface{}
	for _, beneficiaryItem := range beneficiary {
		beneficiaryRule = append(beneficiaryRule, beneficiaryItem)
	}
	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}

	logs, sub, err := _JBSucker.contract.FilterLogs(opts, "InsertToOutboxTree", beneficiaryRule, tokenRule)
	if err != nil {
		return nil, err
	}
	return &JBSuckerInsertToOutboxTreeIterator{contract: _JBSucker.contract, event: "InsertToOutboxTree", logs: logs, sub: sub}, nil
}

// WatchInsertToOutboxTree is a free log subscription operation binding the contract event 0x584a41ff082fad99bb6e50c5767c347c762057fe5726539f0a8de49d84549c23.
//
// Solidity: event InsertToOutboxTree(address indexed beneficiary, address indexed token, bytes32 hashed, uint256 index, bytes32 root, uint256 projectTokenCount, uint256 terminalTokenAmount, address caller)
func (_JBSucker *JBSuckerFilterer) WatchInsertToOutboxTree(opts *bind.WatchOpts, sink chan<- *JBSuckerInsertToOutboxTree, beneficiary []common.Address, token []common.Address) (event.Subscription, error) {

	var beneficiaryRule []interface{}
	for _, beneficiaryItem := range beneficiary {
		beneficiaryRule = append(beneficiaryRule, beneficiaryItem)
	}
	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}

	logs, sub, err := _JBSucker.contract.WatchLogs(opts, "InsertToOutboxTree", beneficiaryRule, tokenRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(JBSuckerInsertToOutboxTree)
				if err := _JBSucker.contract.UnpackLog(event, "InsertToOutboxTree", log); err != nil {
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

// ParseInsertToOutboxTree is a log parse operation binding the contract event 0x584a41ff082fad99bb6e50c5767c347c762057fe5726539f0a8de49d84549c23.
//
// Solidity: event InsertToOutboxTree(address indexed beneficiary, address indexed token, bytes32 hashed, uint256 index, bytes32 root, uint256 projectTokenCount, uint256 terminalTokenAmount, address caller)
func (_JBSucker *JBSuckerFilterer) ParseInsertToOutboxTree(log types.Log) (*JBSuckerInsertToOutboxTree, error) {
	event := new(JBSuckerInsertToOutboxTree)
	if err := _JBSucker.contract.UnpackLog(event, "InsertToOutboxTree", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// JBSuckerNewInboxTreeRootIterator is returned from FilterNewInboxTreeRoot and is used to iterate over the raw logs and unpacked data for NewInboxTreeRoot events raised by the JBSucker contract.
type JBSuckerNewInboxTreeRootIterator struct {
	Event *JBSuckerNewInboxTreeRoot // Event containing the contract specifics and raw log

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
func (it *JBSuckerNewInboxTreeRootIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(JBSuckerNewInboxTreeRoot)
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
		it.Event = new(JBSuckerNewInboxTreeRoot)
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
func (it *JBSuckerNewInboxTreeRootIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *JBSuckerNewInboxTreeRootIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// JBSuckerNewInboxTreeRoot represents a NewInboxTreeRoot event raised by the JBSucker contract.
type JBSuckerNewInboxTreeRoot struct {
	Token  common.Address
	Nonce  uint64
	Root   [32]byte
	Caller common.Address
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterNewInboxTreeRoot is a free log retrieval operation binding the contract event 0xc69d3412489554fb7820c4c2784cb644b6c79bfc117444b93977d3d33a9b8eac.
//
// Solidity: event NewInboxTreeRoot(address indexed token, uint64 nonce, bytes32 root, address caller)
func (_JBSucker *JBSuckerFilterer) FilterNewInboxTreeRoot(opts *bind.FilterOpts, token []common.Address) (*JBSuckerNewInboxTreeRootIterator, error) {

	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}

	logs, sub, err := _JBSucker.contract.FilterLogs(opts, "NewInboxTreeRoot", tokenRule)
	if err != nil {
		return nil, err
	}
	return &JBSuckerNewInboxTreeRootIterator{contract: _JBSucker.contract, event: "NewInboxTreeRoot", logs: logs, sub: sub}, nil
}

// WatchNewInboxTreeRoot is a free log subscription operation binding the contract event 0xc69d3412489554fb7820c4c2784cb644b6c79bfc117444b93977d3d33a9b8eac.
//
// Solidity: event NewInboxTreeRoot(address indexed token, uint64 nonce, bytes32 root, address caller)
func (_JBSucker *JBSuckerFilterer) WatchNewInboxTreeRoot(opts *bind.WatchOpts, sink chan<- *JBSuckerNewInboxTreeRoot, token []common.Address) (event.Subscription, error) {

	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}

	logs, sub, err := _JBSucker.contract.WatchLogs(opts, "NewInboxTreeRoot", tokenRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(JBSuckerNewInboxTreeRoot)
				if err := _JBSucker.contract.UnpackLog(event, "NewInboxTreeRoot", log); err != nil {
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

// ParseNewInboxTreeRoot is a log parse operation binding the contract event 0xc69d3412489554fb7820c4c2784cb644b6c79bfc117444b93977d3d33a9b8eac.
//
// Solidity: event NewInboxTreeRoot(address indexed token, uint64 nonce, bytes32 root, address caller)
func (_JBSucker *JBSuckerFilterer) ParseNewInboxTreeRoot(log types.Log) (*JBSuckerNewInboxTreeRoot, error) {
	event := new(JBSuckerNewInboxTreeRoot)
	if err := _JBSucker.contract.UnpackLog(event, "NewInboxTreeRoot", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// JBSuckerRootToRemoteIterator is returned from FilterRootToRemote and is used to iterate over the raw logs and unpacked data for RootToRemote events raised by the JBSucker contract.
type JBSuckerRootToRemoteIterator struct {
	Event *JBSuckerRootToRemote // Event containing the contract specifics and raw log

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
func (it *JBSuckerRootToRemoteIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(JBSuckerRootToRemote)
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
		it.Event = new(JBSuckerRootToRemote)
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
func (it *JBSuckerRootToRemoteIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *JBSuckerRootToRemoteIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// JBSuckerRootToRemote represents a RootToRemote event raised by the JBSucker contract.
type JBSuckerRootToRemote struct {
	Root   [32]byte
	Token  common.Address
	Index  *big.Int
	Nonce  uint64
	Caller common.Address
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterRootToRemote is a free log retrieval operation binding the contract event 0x47c5615513b0f669da2dd7e3362a34610ca2695d0a0e2da2160cee952eab203c.
//
// Solidity: event RootToRemote(bytes32 indexed root, address indexed token, uint256 index, uint64 nonce, address caller)
func (_JBSucker *JBSuckerFilterer) FilterRootToRemote(opts *bind.FilterOpts, root [][32]byte, token []common.Address) (*JBSuckerRootToRemoteIterator, error) {

	var rootRule []interface{}
	for _, rootItem := range root {
		rootRule = append(rootRule, rootItem)
	}
	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}

	logs, sub, err := _JBSucker.contract.FilterLogs(opts, "RootToRemote", rootRule, tokenRule)
	if err != nil {
		return nil, err
	}
	return &JBSuckerRootToRemoteIterator{contract: _JBSucker.contract, event: "RootToRemote", logs: logs, sub: sub}, nil
}

// WatchRootToRemote is a free log subscription operation binding the contract event 0x47c5615513b0f669da2dd7e3362a34610ca2695d0a0e2da2160cee952eab203c.
//
// Solidity: event RootToRemote(bytes32 indexed root, address indexed token, uint256 index, uint64 nonce, address caller)
func (_JBSucker *JBSuckerFilterer) WatchRootToRemote(opts *bind.WatchOpts, sink chan<- *JBSuckerRootToRemote, root [][32]byte, token []common.Address) (event.Subscription, error) {

	var rootRule []interface{}
	for _, rootItem := range root {
		rootRule = append(rootRule, rootItem)
	}
	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}

	logs, sub, err := _JBSucker.contract.WatchLogs(opts, "RootToRemote", rootRule, tokenRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(JBSuckerRootToRemote)
				if err := _JBSucker.contract.UnpackLog(event, "RootToRemote", log); err != nil {
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

// ParseRootToRemote is a log parse operation binding the contract event 0x47c5615513b0f669da2dd7e3362a34610ca2695d0a0e2da2160cee952eab203c.
//
// Solidity: event RootToRemote(bytes32 indexed root, address indexed token, uint256 index, uint64 nonce, address caller)
func (_JBSucker *JBSuckerFilterer) ParseRootToRemote(log types.Log) (*JBSuckerRootToRemote, error) {
	event := new(JBSuckerRootToRemote)
	if err := _JBSucker.contract.UnpackLog(event, "RootToRemote", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
