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

// BPClaim is an auto generated low-level Go binding around an user-defined struct.
type BPClaim struct {
	Token common.Address
	Leaf  BPLeaf
	Proof [32][32]byte
}

// BPInboxTreeRoot is an auto generated low-level Go binding around an user-defined struct.
type BPInboxTreeRoot struct {
	Nonce uint64
	Root  [32]byte
}

// BPLeaf is an auto generated low-level Go binding around an user-defined struct.
type BPLeaf struct {
	Index               *big.Int
	Beneficiary         common.Address
	ProjectTokenAmount  *big.Int
	TerminalTokenAmount *big.Int
}

// BPMessageRoot is an auto generated low-level Go binding around an user-defined struct.
type BPMessageRoot struct {
	Token      common.Address
	Amount     *big.Int
	RemoteRoot BPInboxTreeRoot
}

// BPTokenMapping is an auto generated low-level Go binding around an user-defined struct.
type BPTokenMapping struct {
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

// BPSuckerMetaData contains all meta data concerning the BPSucker contract.
var BPSuckerMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"receive\",\"stateMutability\":\"payable\"},{\"type\":\"function\",\"name\":\"ADD_TO_BALANCE_MODE\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint8\",\"internalType\":\"enumBPAddToBalanceMode\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"DEPLOYER\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"DIRECTORY\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIJBDirectory\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"PEER\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"PERMISSIONS\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIJBPermissions\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"PROJECT_ID\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"TOKENS\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIJBTokens\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"addOutstandingAmountToBalance\",\"inputs\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"amountToAddToBalance\",\"inputs\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"claim\",\"inputs\":[{\"name\":\"claimData\",\"type\":\"tuple\",\"internalType\":\"structBPClaim\",\"components\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"leaf\",\"type\":\"tuple\",\"internalType\":\"structBPLeaf\",\"components\":[{\"name\":\"index\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"beneficiary\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"projectTokenAmount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"terminalTokenAmount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"proof\",\"type\":\"bytes32[32]\",\"internalType\":\"bytes32[32]\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"claim\",\"inputs\":[{\"name\":\"claims\",\"type\":\"tuple[]\",\"internalType\":\"structBPClaim[]\",\"components\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"leaf\",\"type\":\"tuple\",\"internalType\":\"structBPLeaf\",\"components\":[{\"name\":\"index\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"beneficiary\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"projectTokenAmount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"terminalTokenAmount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"proof\",\"type\":\"bytes32[32]\",\"internalType\":\"bytes32[32]\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"fromRemote\",\"inputs\":[{\"name\":\"root\",\"type\":\"tuple\",\"internalType\":\"structBPMessageRoot\",\"components\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"remoteRoot\",\"type\":\"tuple\",\"internalType\":\"structBPInboxTreeRoot\",\"components\":[{\"name\":\"nonce\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"root\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]}]}],\"outputs\":[],\"stateMutability\":\"payable\"},{\"type\":\"function\",\"name\":\"inbox\",\"inputs\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"nonce\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"root\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"isMapped\",\"inputs\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"mapToken\",\"inputs\":[{\"name\":\"map\",\"type\":\"tuple\",\"internalType\":\"structBPTokenMapping\",\"components\":[{\"name\":\"localToken\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"minGas\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"remoteToken\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"minBridgeAmount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"mapTokens\",\"inputs\":[{\"name\":\"maps\",\"type\":\"tuple[]\",\"internalType\":\"structBPTokenMapping[]\",\"components\":[{\"name\":\"localToken\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"minGas\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"remoteToken\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"minBridgeAmount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"outbox\",\"inputs\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"nonce\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"balance\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"tree\",\"type\":\"tuple\",\"internalType\":\"structMerkleLib.Tree\",\"components\":[{\"name\":\"branch\",\"type\":\"bytes32[32]\",\"internalType\":\"bytes32[32]\"},{\"name\":\"count\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"prepare\",\"inputs\":[{\"name\":\"projectTokenAmount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"beneficiary\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"minTokensReclaimed\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"remoteTokenFor\",\"inputs\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"minGas\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"addr\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"minBridgeAmount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"toRemote\",\"inputs\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"payable\"},{\"type\":\"event\",\"name\":\"InsertToOutboxTree\",\"inputs\":[{\"name\":\"beneficiary\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"terminalToken\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"hashed\",\"type\":\"bytes32\",\"indexed\":false,\"internalType\":\"bytes32\"},{\"name\":\"index\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"root\",\"type\":\"bytes32\",\"indexed\":false,\"internalType\":\"bytes32\"},{\"name\":\"projectTokenAmount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"terminalTokenAmount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"NewInboxTreeRoot\",\"inputs\":[{\"name\":\"token\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"nonce\",\"type\":\"uint64\",\"indexed\":false,\"internalType\":\"uint64\"},{\"name\":\"root\",\"type\":\"bytes32\",\"indexed\":false,\"internalType\":\"bytes32\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"RootToRemote\",\"inputs\":[{\"name\":\"root\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"terminalToken\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"index\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"nonce\",\"type\":\"uint64\",\"indexed\":false,\"internalType\":\"uint64\"}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"AddressEmptyCode\",\"inputs\":[{\"name\":\"target\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"AddressInsufficientBalance\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"BELOW_MIN_GAS\",\"inputs\":[{\"name\":\"minGas\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"suppliedGas\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"BENEFICIARY_NOT_ALLOWED\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ERC20_TOKEN_REQUIRED\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"FailedInnerCall\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"INSUFFICIENT_BALANCE\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"INVALID_NATIVE_REMOTE_ADDRESS\",\"inputs\":[{\"name\":\"addr\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"INVALID_PROOF\",\"inputs\":[{\"name\":\"expectedRoot\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"proofRoot\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]},{\"type\":\"error\",\"name\":\"LEAF_ALREADY_EXECUTED\",\"inputs\":[{\"name\":\"index\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"MANUAL_NOT_ALLOWED\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"MerkleLib__insert_treeIsFull\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"NOT_PEER\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"NO_TERMINAL_FOR\",\"inputs\":[{\"name\":\"projectId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"QUEUE_INSUFFECIENT_SIZE\",\"inputs\":[{\"name\":\"minSize\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"currentSize\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"SafeERC20FailedOperation\",\"inputs\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"TOKEN_NOT_MAPPED\",\"inputs\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"UNAUTHORIZED\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"UNEXPECTED_MSG_VALUE\",\"inputs\":[]}]",
}

// BPSuckerABI is the input ABI used to generate the binding from.
// Deprecated: Use BPSuckerMetaData.ABI instead.
var BPSuckerABI = BPSuckerMetaData.ABI

// BPSucker is an auto generated Go binding around an Ethereum contract.
type BPSucker struct {
	BPSuckerCaller     // Read-only binding to the contract
	BPSuckerTransactor // Write-only binding to the contract
	BPSuckerFilterer   // Log filterer for contract events
}

// BPSuckerCaller is an auto generated read-only Go binding around an Ethereum contract.
type BPSuckerCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BPSuckerTransactor is an auto generated write-only Go binding around an Ethereum contract.
type BPSuckerTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BPSuckerFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type BPSuckerFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BPSuckerSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type BPSuckerSession struct {
	Contract     *BPSucker         // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// BPSuckerCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type BPSuckerCallerSession struct {
	Contract *BPSuckerCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts   // Call options to use throughout this session
}

// BPSuckerTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type BPSuckerTransactorSession struct {
	Contract     *BPSuckerTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// BPSuckerRaw is an auto generated low-level Go binding around an Ethereum contract.
type BPSuckerRaw struct {
	Contract *BPSucker // Generic contract binding to access the raw methods on
}

// BPSuckerCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type BPSuckerCallerRaw struct {
	Contract *BPSuckerCaller // Generic read-only contract binding to access the raw methods on
}

// BPSuckerTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type BPSuckerTransactorRaw struct {
	Contract *BPSuckerTransactor // Generic write-only contract binding to access the raw methods on
}

// NewBPSucker creates a new instance of BPSucker, bound to a specific deployed contract.
func NewBPSucker(address common.Address, backend bind.ContractBackend) (*BPSucker, error) {
	contract, err := bindBPSucker(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &BPSucker{BPSuckerCaller: BPSuckerCaller{contract: contract}, BPSuckerTransactor: BPSuckerTransactor{contract: contract}, BPSuckerFilterer: BPSuckerFilterer{contract: contract}}, nil
}

// NewBPSuckerCaller creates a new read-only instance of BPSucker, bound to a specific deployed contract.
func NewBPSuckerCaller(address common.Address, caller bind.ContractCaller) (*BPSuckerCaller, error) {
	contract, err := bindBPSucker(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &BPSuckerCaller{contract: contract}, nil
}

// NewBPSuckerTransactor creates a new write-only instance of BPSucker, bound to a specific deployed contract.
func NewBPSuckerTransactor(address common.Address, transactor bind.ContractTransactor) (*BPSuckerTransactor, error) {
	contract, err := bindBPSucker(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &BPSuckerTransactor{contract: contract}, nil
}

// NewBPSuckerFilterer creates a new log filterer instance of BPSucker, bound to a specific deployed contract.
func NewBPSuckerFilterer(address common.Address, filterer bind.ContractFilterer) (*BPSuckerFilterer, error) {
	contract, err := bindBPSucker(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &BPSuckerFilterer{contract: contract}, nil
}

// bindBPSucker binds a generic wrapper to an already deployed contract.
func bindBPSucker(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := BPSuckerMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_BPSucker *BPSuckerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _BPSucker.Contract.BPSuckerCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_BPSucker *BPSuckerRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BPSucker.Contract.BPSuckerTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_BPSucker *BPSuckerRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _BPSucker.Contract.BPSuckerTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_BPSucker *BPSuckerCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _BPSucker.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_BPSucker *BPSuckerTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BPSucker.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_BPSucker *BPSuckerTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _BPSucker.Contract.contract.Transact(opts, method, params...)
}

// ADDTOBALANCEMODE is a free data retrieval call binding the contract method 0xc6525f2c.
//
// Solidity: function ADD_TO_BALANCE_MODE() view returns(uint8)
func (_BPSucker *BPSuckerCaller) ADDTOBALANCEMODE(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _BPSucker.contract.Call(opts, &out, "ADD_TO_BALANCE_MODE")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// ADDTOBALANCEMODE is a free data retrieval call binding the contract method 0xc6525f2c.
//
// Solidity: function ADD_TO_BALANCE_MODE() view returns(uint8)
func (_BPSucker *BPSuckerSession) ADDTOBALANCEMODE() (uint8, error) {
	return _BPSucker.Contract.ADDTOBALANCEMODE(&_BPSucker.CallOpts)
}

// ADDTOBALANCEMODE is a free data retrieval call binding the contract method 0xc6525f2c.
//
// Solidity: function ADD_TO_BALANCE_MODE() view returns(uint8)
func (_BPSucker *BPSuckerCallerSession) ADDTOBALANCEMODE() (uint8, error) {
	return _BPSucker.Contract.ADDTOBALANCEMODE(&_BPSucker.CallOpts)
}

// DEPLOYER is a free data retrieval call binding the contract method 0xc1b8411a.
//
// Solidity: function DEPLOYER() view returns(address)
func (_BPSucker *BPSuckerCaller) DEPLOYER(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _BPSucker.contract.Call(opts, &out, "DEPLOYER")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// DEPLOYER is a free data retrieval call binding the contract method 0xc1b8411a.
//
// Solidity: function DEPLOYER() view returns(address)
func (_BPSucker *BPSuckerSession) DEPLOYER() (common.Address, error) {
	return _BPSucker.Contract.DEPLOYER(&_BPSucker.CallOpts)
}

// DEPLOYER is a free data retrieval call binding the contract method 0xc1b8411a.
//
// Solidity: function DEPLOYER() view returns(address)
func (_BPSucker *BPSuckerCallerSession) DEPLOYER() (common.Address, error) {
	return _BPSucker.Contract.DEPLOYER(&_BPSucker.CallOpts)
}

// DIRECTORY is a free data retrieval call binding the contract method 0x88bc2ef3.
//
// Solidity: function DIRECTORY() view returns(address)
func (_BPSucker *BPSuckerCaller) DIRECTORY(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _BPSucker.contract.Call(opts, &out, "DIRECTORY")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// DIRECTORY is a free data retrieval call binding the contract method 0x88bc2ef3.
//
// Solidity: function DIRECTORY() view returns(address)
func (_BPSucker *BPSuckerSession) DIRECTORY() (common.Address, error) {
	return _BPSucker.Contract.DIRECTORY(&_BPSucker.CallOpts)
}

// DIRECTORY is a free data retrieval call binding the contract method 0x88bc2ef3.
//
// Solidity: function DIRECTORY() view returns(address)
func (_BPSucker *BPSuckerCallerSession) DIRECTORY() (common.Address, error) {
	return _BPSucker.Contract.DIRECTORY(&_BPSucker.CallOpts)
}

// PEER is a free data retrieval call binding the contract method 0xad42938b.
//
// Solidity: function PEER() view returns(address)
func (_BPSucker *BPSuckerCaller) PEER(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _BPSucker.contract.Call(opts, &out, "PEER")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// PEER is a free data retrieval call binding the contract method 0xad42938b.
//
// Solidity: function PEER() view returns(address)
func (_BPSucker *BPSuckerSession) PEER() (common.Address, error) {
	return _BPSucker.Contract.PEER(&_BPSucker.CallOpts)
}

// PEER is a free data retrieval call binding the contract method 0xad42938b.
//
// Solidity: function PEER() view returns(address)
func (_BPSucker *BPSuckerCallerSession) PEER() (common.Address, error) {
	return _BPSucker.Contract.PEER(&_BPSucker.CallOpts)
}

// PERMISSIONS is a free data retrieval call binding the contract method 0xf434c914.
//
// Solidity: function PERMISSIONS() view returns(address)
func (_BPSucker *BPSuckerCaller) PERMISSIONS(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _BPSucker.contract.Call(opts, &out, "PERMISSIONS")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// PERMISSIONS is a free data retrieval call binding the contract method 0xf434c914.
//
// Solidity: function PERMISSIONS() view returns(address)
func (_BPSucker *BPSuckerSession) PERMISSIONS() (common.Address, error) {
	return _BPSucker.Contract.PERMISSIONS(&_BPSucker.CallOpts)
}

// PERMISSIONS is a free data retrieval call binding the contract method 0xf434c914.
//
// Solidity: function PERMISSIONS() view returns(address)
func (_BPSucker *BPSuckerCallerSession) PERMISSIONS() (common.Address, error) {
	return _BPSucker.Contract.PERMISSIONS(&_BPSucker.CallOpts)
}

// PROJECTID is a free data retrieval call binding the contract method 0x56539f39.
//
// Solidity: function PROJECT_ID() view returns(uint256)
func (_BPSucker *BPSuckerCaller) PROJECTID(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _BPSucker.contract.Call(opts, &out, "PROJECT_ID")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PROJECTID is a free data retrieval call binding the contract method 0x56539f39.
//
// Solidity: function PROJECT_ID() view returns(uint256)
func (_BPSucker *BPSuckerSession) PROJECTID() (*big.Int, error) {
	return _BPSucker.Contract.PROJECTID(&_BPSucker.CallOpts)
}

// PROJECTID is a free data retrieval call binding the contract method 0x56539f39.
//
// Solidity: function PROJECT_ID() view returns(uint256)
func (_BPSucker *BPSuckerCallerSession) PROJECTID() (*big.Int, error) {
	return _BPSucker.Contract.PROJECTID(&_BPSucker.CallOpts)
}

// TOKENS is a free data retrieval call binding the contract method 0x1d831d5c.
//
// Solidity: function TOKENS() view returns(address)
func (_BPSucker *BPSuckerCaller) TOKENS(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _BPSucker.contract.Call(opts, &out, "TOKENS")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// TOKENS is a free data retrieval call binding the contract method 0x1d831d5c.
//
// Solidity: function TOKENS() view returns(address)
func (_BPSucker *BPSuckerSession) TOKENS() (common.Address, error) {
	return _BPSucker.Contract.TOKENS(&_BPSucker.CallOpts)
}

// TOKENS is a free data retrieval call binding the contract method 0x1d831d5c.
//
// Solidity: function TOKENS() view returns(address)
func (_BPSucker *BPSuckerCallerSession) TOKENS() (common.Address, error) {
	return _BPSucker.Contract.TOKENS(&_BPSucker.CallOpts)
}

// AmountToAddToBalance is a free data retrieval call binding the contract method 0xe7f82c75.
//
// Solidity: function amountToAddToBalance(address token) view returns(uint256 amount)
func (_BPSucker *BPSuckerCaller) AmountToAddToBalance(opts *bind.CallOpts, token common.Address) (*big.Int, error) {
	var out []interface{}
	err := _BPSucker.contract.Call(opts, &out, "amountToAddToBalance", token)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// AmountToAddToBalance is a free data retrieval call binding the contract method 0xe7f82c75.
//
// Solidity: function amountToAddToBalance(address token) view returns(uint256 amount)
func (_BPSucker *BPSuckerSession) AmountToAddToBalance(token common.Address) (*big.Int, error) {
	return _BPSucker.Contract.AmountToAddToBalance(&_BPSucker.CallOpts, token)
}

// AmountToAddToBalance is a free data retrieval call binding the contract method 0xe7f82c75.
//
// Solidity: function amountToAddToBalance(address token) view returns(uint256 amount)
func (_BPSucker *BPSuckerCallerSession) AmountToAddToBalance(token common.Address) (*big.Int, error) {
	return _BPSucker.Contract.AmountToAddToBalance(&_BPSucker.CallOpts, token)
}

// Inbox is a free data retrieval call binding the contract method 0xed0e26b9.
//
// Solidity: function inbox(address token) view returns(uint64 nonce, bytes32 root)
func (_BPSucker *BPSuckerCaller) Inbox(opts *bind.CallOpts, token common.Address) (struct {
	Nonce uint64
	Root  [32]byte
}, error) {
	var out []interface{}
	err := _BPSucker.contract.Call(opts, &out, "inbox", token)

	outstruct := new(struct {
		Nonce uint64
		Root  [32]byte
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Nonce = *abi.ConvertType(out[0], new(uint64)).(*uint64)
	outstruct.Root = *abi.ConvertType(out[1], new([32]byte)).(*[32]byte)

	return *outstruct, err

}

// Inbox is a free data retrieval call binding the contract method 0xed0e26b9.
//
// Solidity: function inbox(address token) view returns(uint64 nonce, bytes32 root)
func (_BPSucker *BPSuckerSession) Inbox(token common.Address) (struct {
	Nonce uint64
	Root  [32]byte
}, error) {
	return _BPSucker.Contract.Inbox(&_BPSucker.CallOpts, token)
}

// Inbox is a free data retrieval call binding the contract method 0xed0e26b9.
//
// Solidity: function inbox(address token) view returns(uint64 nonce, bytes32 root)
func (_BPSucker *BPSuckerCallerSession) Inbox(token common.Address) (struct {
	Nonce uint64
	Root  [32]byte
}, error) {
	return _BPSucker.Contract.Inbox(&_BPSucker.CallOpts, token)
}

// IsMapped is a free data retrieval call binding the contract method 0xbf557aee.
//
// Solidity: function isMapped(address token) view returns(bool)
func (_BPSucker *BPSuckerCaller) IsMapped(opts *bind.CallOpts, token common.Address) (bool, error) {
	var out []interface{}
	err := _BPSucker.contract.Call(opts, &out, "isMapped", token)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsMapped is a free data retrieval call binding the contract method 0xbf557aee.
//
// Solidity: function isMapped(address token) view returns(bool)
func (_BPSucker *BPSuckerSession) IsMapped(token common.Address) (bool, error) {
	return _BPSucker.Contract.IsMapped(&_BPSucker.CallOpts, token)
}

// IsMapped is a free data retrieval call binding the contract method 0xbf557aee.
//
// Solidity: function isMapped(address token) view returns(bool)
func (_BPSucker *BPSuckerCallerSession) IsMapped(token common.Address) (bool, error) {
	return _BPSucker.Contract.IsMapped(&_BPSucker.CallOpts, token)
}

// Outbox is a free data retrieval call binding the contract method 0x6dfa7268.
//
// Solidity: function outbox(address token) view returns(uint64 nonce, uint256 balance, (bytes32[32],uint256) tree)
func (_BPSucker *BPSuckerCaller) Outbox(opts *bind.CallOpts, token common.Address) (struct {
	Nonce   uint64
	Balance *big.Int
	Tree    MerkleLibTree
}, error) {
	var out []interface{}
	err := _BPSucker.contract.Call(opts, &out, "outbox", token)

	outstruct := new(struct {
		Nonce   uint64
		Balance *big.Int
		Tree    MerkleLibTree
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Nonce = *abi.ConvertType(out[0], new(uint64)).(*uint64)
	outstruct.Balance = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.Tree = *abi.ConvertType(out[2], new(MerkleLibTree)).(*MerkleLibTree)

	return *outstruct, err

}

// Outbox is a free data retrieval call binding the contract method 0x6dfa7268.
//
// Solidity: function outbox(address token) view returns(uint64 nonce, uint256 balance, (bytes32[32],uint256) tree)
func (_BPSucker *BPSuckerSession) Outbox(token common.Address) (struct {
	Nonce   uint64
	Balance *big.Int
	Tree    MerkleLibTree
}, error) {
	return _BPSucker.Contract.Outbox(&_BPSucker.CallOpts, token)
}

// Outbox is a free data retrieval call binding the contract method 0x6dfa7268.
//
// Solidity: function outbox(address token) view returns(uint64 nonce, uint256 balance, (bytes32[32],uint256) tree)
func (_BPSucker *BPSuckerCallerSession) Outbox(token common.Address) (struct {
	Nonce   uint64
	Balance *big.Int
	Tree    MerkleLibTree
}, error) {
	return _BPSucker.Contract.Outbox(&_BPSucker.CallOpts, token)
}

// RemoteTokenFor is a free data retrieval call binding the contract method 0x7de6fba2.
//
// Solidity: function remoteTokenFor(address token) view returns(uint32 minGas, address addr, uint256 minBridgeAmount)
func (_BPSucker *BPSuckerCaller) RemoteTokenFor(opts *bind.CallOpts, token common.Address) (struct {
	MinGas          uint32
	Addr            common.Address
	MinBridgeAmount *big.Int
}, error) {
	var out []interface{}
	err := _BPSucker.contract.Call(opts, &out, "remoteTokenFor", token)

	outstruct := new(struct {
		MinGas          uint32
		Addr            common.Address
		MinBridgeAmount *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.MinGas = *abi.ConvertType(out[0], new(uint32)).(*uint32)
	outstruct.Addr = *abi.ConvertType(out[1], new(common.Address)).(*common.Address)
	outstruct.MinBridgeAmount = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// RemoteTokenFor is a free data retrieval call binding the contract method 0x7de6fba2.
//
// Solidity: function remoteTokenFor(address token) view returns(uint32 minGas, address addr, uint256 minBridgeAmount)
func (_BPSucker *BPSuckerSession) RemoteTokenFor(token common.Address) (struct {
	MinGas          uint32
	Addr            common.Address
	MinBridgeAmount *big.Int
}, error) {
	return _BPSucker.Contract.RemoteTokenFor(&_BPSucker.CallOpts, token)
}

// RemoteTokenFor is a free data retrieval call binding the contract method 0x7de6fba2.
//
// Solidity: function remoteTokenFor(address token) view returns(uint32 minGas, address addr, uint256 minBridgeAmount)
func (_BPSucker *BPSuckerCallerSession) RemoteTokenFor(token common.Address) (struct {
	MinGas          uint32
	Addr            common.Address
	MinBridgeAmount *big.Int
}, error) {
	return _BPSucker.Contract.RemoteTokenFor(&_BPSucker.CallOpts, token)
}

// AddOutstandingAmountToBalance is a paid mutator transaction binding the contract method 0x69ce2233.
//
// Solidity: function addOutstandingAmountToBalance(address token) returns()
func (_BPSucker *BPSuckerTransactor) AddOutstandingAmountToBalance(opts *bind.TransactOpts, token common.Address) (*types.Transaction, error) {
	return _BPSucker.contract.Transact(opts, "addOutstandingAmountToBalance", token)
}

// AddOutstandingAmountToBalance is a paid mutator transaction binding the contract method 0x69ce2233.
//
// Solidity: function addOutstandingAmountToBalance(address token) returns()
func (_BPSucker *BPSuckerSession) AddOutstandingAmountToBalance(token common.Address) (*types.Transaction, error) {
	return _BPSucker.Contract.AddOutstandingAmountToBalance(&_BPSucker.TransactOpts, token)
}

// AddOutstandingAmountToBalance is a paid mutator transaction binding the contract method 0x69ce2233.
//
// Solidity: function addOutstandingAmountToBalance(address token) returns()
func (_BPSucker *BPSuckerTransactorSession) AddOutstandingAmountToBalance(token common.Address) (*types.Transaction, error) {
	return _BPSucker.Contract.AddOutstandingAmountToBalance(&_BPSucker.TransactOpts, token)
}

// Claim is a paid mutator transaction binding the contract method 0xb59ef50f.
//
// Solidity: function claim((address,(uint256,address,uint256,uint256),bytes32[32]) claimData) returns()
func (_BPSucker *BPSuckerTransactor) Claim(opts *bind.TransactOpts, claimData BPClaim) (*types.Transaction, error) {
	return _BPSucker.contract.Transact(opts, "claim", claimData)
}

// Claim is a paid mutator transaction binding the contract method 0xb59ef50f.
//
// Solidity: function claim((address,(uint256,address,uint256,uint256),bytes32[32]) claimData) returns()
func (_BPSucker *BPSuckerSession) Claim(claimData BPClaim) (*types.Transaction, error) {
	return _BPSucker.Contract.Claim(&_BPSucker.TransactOpts, claimData)
}

// Claim is a paid mutator transaction binding the contract method 0xb59ef50f.
//
// Solidity: function claim((address,(uint256,address,uint256,uint256),bytes32[32]) claimData) returns()
func (_BPSucker *BPSuckerTransactorSession) Claim(claimData BPClaim) (*types.Transaction, error) {
	return _BPSucker.Contract.Claim(&_BPSucker.TransactOpts, claimData)
}

// Claim0 is a paid mutator transaction binding the contract method 0xc1854129.
//
// Solidity: function claim((address,(uint256,address,uint256,uint256),bytes32[32])[] claims) returns()
func (_BPSucker *BPSuckerTransactor) Claim0(opts *bind.TransactOpts, claims []BPClaim) (*types.Transaction, error) {
	return _BPSucker.contract.Transact(opts, "claim0", claims)
}

// Claim0 is a paid mutator transaction binding the contract method 0xc1854129.
//
// Solidity: function claim((address,(uint256,address,uint256,uint256),bytes32[32])[] claims) returns()
func (_BPSucker *BPSuckerSession) Claim0(claims []BPClaim) (*types.Transaction, error) {
	return _BPSucker.Contract.Claim0(&_BPSucker.TransactOpts, claims)
}

// Claim0 is a paid mutator transaction binding the contract method 0xc1854129.
//
// Solidity: function claim((address,(uint256,address,uint256,uint256),bytes32[32])[] claims) returns()
func (_BPSucker *BPSuckerTransactorSession) Claim0(claims []BPClaim) (*types.Transaction, error) {
	return _BPSucker.Contract.Claim0(&_BPSucker.TransactOpts, claims)
}

// FromRemote is a paid mutator transaction binding the contract method 0x69d31fdd.
//
// Solidity: function fromRemote((address,uint256,(uint64,bytes32)) root) payable returns()
func (_BPSucker *BPSuckerTransactor) FromRemote(opts *bind.TransactOpts, root BPMessageRoot) (*types.Transaction, error) {
	return _BPSucker.contract.Transact(opts, "fromRemote", root)
}

// FromRemote is a paid mutator transaction binding the contract method 0x69d31fdd.
//
// Solidity: function fromRemote((address,uint256,(uint64,bytes32)) root) payable returns()
func (_BPSucker *BPSuckerSession) FromRemote(root BPMessageRoot) (*types.Transaction, error) {
	return _BPSucker.Contract.FromRemote(&_BPSucker.TransactOpts, root)
}

// FromRemote is a paid mutator transaction binding the contract method 0x69d31fdd.
//
// Solidity: function fromRemote((address,uint256,(uint64,bytes32)) root) payable returns()
func (_BPSucker *BPSuckerTransactorSession) FromRemote(root BPMessageRoot) (*types.Transaction, error) {
	return _BPSucker.Contract.FromRemote(&_BPSucker.TransactOpts, root)
}

// MapToken is a paid mutator transaction binding the contract method 0x38eb2044.
//
// Solidity: function mapToken((address,uint32,address,uint256) map) returns()
func (_BPSucker *BPSuckerTransactor) MapToken(opts *bind.TransactOpts, arg0 BPTokenMapping) (*types.Transaction, error) {
	return _BPSucker.contract.Transact(opts, "mapToken", arg0)
}

// MapToken is a paid mutator transaction binding the contract method 0x38eb2044.
//
// Solidity: function mapToken((address,uint32,address,uint256) map) returns()
func (_BPSucker *BPSuckerSession) MapToken(arg0 BPTokenMapping) (*types.Transaction, error) {
	return _BPSucker.Contract.MapToken(&_BPSucker.TransactOpts, arg0)
}

// MapToken is a paid mutator transaction binding the contract method 0x38eb2044.
//
// Solidity: function mapToken((address,uint32,address,uint256) map) returns()
func (_BPSucker *BPSuckerTransactorSession) MapToken(arg0 BPTokenMapping) (*types.Transaction, error) {
	return _BPSucker.Contract.MapToken(&_BPSucker.TransactOpts, arg0)
}

// MapTokens is a paid mutator transaction binding the contract method 0x2241dc52.
//
// Solidity: function mapTokens((address,uint32,address,uint256)[] maps) returns()
func (_BPSucker *BPSuckerTransactor) MapTokens(opts *bind.TransactOpts, maps []BPTokenMapping) (*types.Transaction, error) {
	return _BPSucker.contract.Transact(opts, "mapTokens", maps)
}

// MapTokens is a paid mutator transaction binding the contract method 0x2241dc52.
//
// Solidity: function mapTokens((address,uint32,address,uint256)[] maps) returns()
func (_BPSucker *BPSuckerSession) MapTokens(maps []BPTokenMapping) (*types.Transaction, error) {
	return _BPSucker.Contract.MapTokens(&_BPSucker.TransactOpts, maps)
}

// MapTokens is a paid mutator transaction binding the contract method 0x2241dc52.
//
// Solidity: function mapTokens((address,uint32,address,uint256)[] maps) returns()
func (_BPSucker *BPSuckerTransactorSession) MapTokens(maps []BPTokenMapping) (*types.Transaction, error) {
	return _BPSucker.Contract.MapTokens(&_BPSucker.TransactOpts, maps)
}

// Prepare is a paid mutator transaction binding the contract method 0x551ebbf0.
//
// Solidity: function prepare(uint256 projectTokenAmount, address beneficiary, uint256 minTokensReclaimed, address token) returns()
func (_BPSucker *BPSuckerTransactor) Prepare(opts *bind.TransactOpts, projectTokenAmount *big.Int, beneficiary common.Address, minTokensReclaimed *big.Int, token common.Address) (*types.Transaction, error) {
	return _BPSucker.contract.Transact(opts, "prepare", projectTokenAmount, beneficiary, minTokensReclaimed, token)
}

// Prepare is a paid mutator transaction binding the contract method 0x551ebbf0.
//
// Solidity: function prepare(uint256 projectTokenAmount, address beneficiary, uint256 minTokensReclaimed, address token) returns()
func (_BPSucker *BPSuckerSession) Prepare(projectTokenAmount *big.Int, beneficiary common.Address, minTokensReclaimed *big.Int, token common.Address) (*types.Transaction, error) {
	return _BPSucker.Contract.Prepare(&_BPSucker.TransactOpts, projectTokenAmount, beneficiary, minTokensReclaimed, token)
}

// Prepare is a paid mutator transaction binding the contract method 0x551ebbf0.
//
// Solidity: function prepare(uint256 projectTokenAmount, address beneficiary, uint256 minTokensReclaimed, address token) returns()
func (_BPSucker *BPSuckerTransactorSession) Prepare(projectTokenAmount *big.Int, beneficiary common.Address, minTokensReclaimed *big.Int, token common.Address) (*types.Transaction, error) {
	return _BPSucker.Contract.Prepare(&_BPSucker.TransactOpts, projectTokenAmount, beneficiary, minTokensReclaimed, token)
}

// ToRemote is a paid mutator transaction binding the contract method 0xb71c1179.
//
// Solidity: function toRemote(address token) payable returns()
func (_BPSucker *BPSuckerTransactor) ToRemote(opts *bind.TransactOpts, token common.Address) (*types.Transaction, error) {
	return _BPSucker.contract.Transact(opts, "toRemote", token)
}

// ToRemote is a paid mutator transaction binding the contract method 0xb71c1179.
//
// Solidity: function toRemote(address token) payable returns()
func (_BPSucker *BPSuckerSession) ToRemote(token common.Address) (*types.Transaction, error) {
	return _BPSucker.Contract.ToRemote(&_BPSucker.TransactOpts, token)
}

// ToRemote is a paid mutator transaction binding the contract method 0xb71c1179.
//
// Solidity: function toRemote(address token) payable returns()
func (_BPSucker *BPSuckerTransactorSession) ToRemote(token common.Address) (*types.Transaction, error) {
	return _BPSucker.Contract.ToRemote(&_BPSucker.TransactOpts, token)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_BPSucker *BPSuckerTransactor) Receive(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BPSucker.contract.RawTransact(opts, nil) // calldata is disallowed for receive function
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_BPSucker *BPSuckerSession) Receive() (*types.Transaction, error) {
	return _BPSucker.Contract.Receive(&_BPSucker.TransactOpts)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_BPSucker *BPSuckerTransactorSession) Receive() (*types.Transaction, error) {
	return _BPSucker.Contract.Receive(&_BPSucker.TransactOpts)
}

// BPSuckerInsertToOutboxTreeIterator is returned from FilterInsertToOutboxTree and is used to iterate over the raw logs and unpacked data for InsertToOutboxTree events raised by the BPSucker contract.
type BPSuckerInsertToOutboxTreeIterator struct {
	Event *BPSuckerInsertToOutboxTree // Event containing the contract specifics and raw log

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
func (it *BPSuckerInsertToOutboxTreeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BPSuckerInsertToOutboxTree)
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
		it.Event = new(BPSuckerInsertToOutboxTree)
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
func (it *BPSuckerInsertToOutboxTreeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BPSuckerInsertToOutboxTreeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BPSuckerInsertToOutboxTree represents a InsertToOutboxTree event raised by the BPSucker contract.
type BPSuckerInsertToOutboxTree struct {
	Beneficiary         common.Address
	TerminalToken       common.Address
	Hashed              [32]byte
	Index               *big.Int
	Root                [32]byte
	ProjectTokenAmount  *big.Int
	TerminalTokenAmount *big.Int
	Raw                 types.Log // Blockchain specific contextual infos
}

// FilterInsertToOutboxTree is a free log retrieval operation binding the contract event 0x7b58dbb78e7edc690b0688a53dea0c443d630c4a5db66c3c47d4ff2044c97728.
//
// Solidity: event InsertToOutboxTree(address indexed beneficiary, address indexed terminalToken, bytes32 hashed, uint256 index, bytes32 root, uint256 projectTokenAmount, uint256 terminalTokenAmount)
func (_BPSucker *BPSuckerFilterer) FilterInsertToOutboxTree(opts *bind.FilterOpts, beneficiary []common.Address, terminalToken []common.Address) (*BPSuckerInsertToOutboxTreeIterator, error) {

	var beneficiaryRule []interface{}
	for _, beneficiaryItem := range beneficiary {
		beneficiaryRule = append(beneficiaryRule, beneficiaryItem)
	}
	var terminalTokenRule []interface{}
	for _, terminalTokenItem := range terminalToken {
		terminalTokenRule = append(terminalTokenRule, terminalTokenItem)
	}

	logs, sub, err := _BPSucker.contract.FilterLogs(opts, "InsertToOutboxTree", beneficiaryRule, terminalTokenRule)
	if err != nil {
		return nil, err
	}
	return &BPSuckerInsertToOutboxTreeIterator{contract: _BPSucker.contract, event: "InsertToOutboxTree", logs: logs, sub: sub}, nil
}

// WatchInsertToOutboxTree is a free log subscription operation binding the contract event 0x7b58dbb78e7edc690b0688a53dea0c443d630c4a5db66c3c47d4ff2044c97728.
//
// Solidity: event InsertToOutboxTree(address indexed beneficiary, address indexed terminalToken, bytes32 hashed, uint256 index, bytes32 root, uint256 projectTokenAmount, uint256 terminalTokenAmount)
func (_BPSucker *BPSuckerFilterer) WatchInsertToOutboxTree(opts *bind.WatchOpts, sink chan<- *BPSuckerInsertToOutboxTree, beneficiary []common.Address, terminalToken []common.Address) (event.Subscription, error) {

	var beneficiaryRule []interface{}
	for _, beneficiaryItem := range beneficiary {
		beneficiaryRule = append(beneficiaryRule, beneficiaryItem)
	}
	var terminalTokenRule []interface{}
	for _, terminalTokenItem := range terminalToken {
		terminalTokenRule = append(terminalTokenRule, terminalTokenItem)
	}

	logs, sub, err := _BPSucker.contract.WatchLogs(opts, "InsertToOutboxTree", beneficiaryRule, terminalTokenRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BPSuckerInsertToOutboxTree)
				if err := _BPSucker.contract.UnpackLog(event, "InsertToOutboxTree", log); err != nil {
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

// ParseInsertToOutboxTree is a log parse operation binding the contract event 0x7b58dbb78e7edc690b0688a53dea0c443d630c4a5db66c3c47d4ff2044c97728.
//
// Solidity: event InsertToOutboxTree(address indexed beneficiary, address indexed terminalToken, bytes32 hashed, uint256 index, bytes32 root, uint256 projectTokenAmount, uint256 terminalTokenAmount)
func (_BPSucker *BPSuckerFilterer) ParseInsertToOutboxTree(log types.Log) (*BPSuckerInsertToOutboxTree, error) {
	event := new(BPSuckerInsertToOutboxTree)
	if err := _BPSucker.contract.UnpackLog(event, "InsertToOutboxTree", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BPSuckerNewInboxTreeRootIterator is returned from FilterNewInboxTreeRoot and is used to iterate over the raw logs and unpacked data for NewInboxTreeRoot events raised by the BPSucker contract.
type BPSuckerNewInboxTreeRootIterator struct {
	Event *BPSuckerNewInboxTreeRoot // Event containing the contract specifics and raw log

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
func (it *BPSuckerNewInboxTreeRootIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BPSuckerNewInboxTreeRoot)
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
		it.Event = new(BPSuckerNewInboxTreeRoot)
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
func (it *BPSuckerNewInboxTreeRootIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BPSuckerNewInboxTreeRootIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BPSuckerNewInboxTreeRoot represents a NewInboxTreeRoot event raised by the BPSucker contract.
type BPSuckerNewInboxTreeRoot struct {
	Token common.Address
	Nonce uint64
	Root  [32]byte
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterNewInboxTreeRoot is a free log retrieval operation binding the contract event 0x551b6b99b4c48e3265d2885ab0d93779be2402293261e074f7bf74b2c8f45b6f.
//
// Solidity: event NewInboxTreeRoot(address indexed token, uint64 nonce, bytes32 root)
func (_BPSucker *BPSuckerFilterer) FilterNewInboxTreeRoot(opts *bind.FilterOpts, token []common.Address) (*BPSuckerNewInboxTreeRootIterator, error) {

	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}

	logs, sub, err := _BPSucker.contract.FilterLogs(opts, "NewInboxTreeRoot", tokenRule)
	if err != nil {
		return nil, err
	}
	return &BPSuckerNewInboxTreeRootIterator{contract: _BPSucker.contract, event: "NewInboxTreeRoot", logs: logs, sub: sub}, nil
}

// WatchNewInboxTreeRoot is a free log subscription operation binding the contract event 0x551b6b99b4c48e3265d2885ab0d93779be2402293261e074f7bf74b2c8f45b6f.
//
// Solidity: event NewInboxTreeRoot(address indexed token, uint64 nonce, bytes32 root)
func (_BPSucker *BPSuckerFilterer) WatchNewInboxTreeRoot(opts *bind.WatchOpts, sink chan<- *BPSuckerNewInboxTreeRoot, token []common.Address) (event.Subscription, error) {

	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}

	logs, sub, err := _BPSucker.contract.WatchLogs(opts, "NewInboxTreeRoot", tokenRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BPSuckerNewInboxTreeRoot)
				if err := _BPSucker.contract.UnpackLog(event, "NewInboxTreeRoot", log); err != nil {
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

// ParseNewInboxTreeRoot is a log parse operation binding the contract event 0x551b6b99b4c48e3265d2885ab0d93779be2402293261e074f7bf74b2c8f45b6f.
//
// Solidity: event NewInboxTreeRoot(address indexed token, uint64 nonce, bytes32 root)
func (_BPSucker *BPSuckerFilterer) ParseNewInboxTreeRoot(log types.Log) (*BPSuckerNewInboxTreeRoot, error) {
	event := new(BPSuckerNewInboxTreeRoot)
	if err := _BPSucker.contract.UnpackLog(event, "NewInboxTreeRoot", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BPSuckerRootToRemoteIterator is returned from FilterRootToRemote and is used to iterate over the raw logs and unpacked data for RootToRemote events raised by the BPSucker contract.
type BPSuckerRootToRemoteIterator struct {
	Event *BPSuckerRootToRemote // Event containing the contract specifics and raw log

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
func (it *BPSuckerRootToRemoteIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BPSuckerRootToRemote)
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
		it.Event = new(BPSuckerRootToRemote)
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
func (it *BPSuckerRootToRemoteIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BPSuckerRootToRemoteIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BPSuckerRootToRemote represents a RootToRemote event raised by the BPSucker contract.
type BPSuckerRootToRemote struct {
	Root          [32]byte
	TerminalToken common.Address
	Index         *big.Int
	Nonce         uint64
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterRootToRemote is a free log retrieval operation binding the contract event 0x3e069c1f498cba23d7d2b74b84311469e8da6e527cade62061d642c61b64e0de.
//
// Solidity: event RootToRemote(bytes32 indexed root, address indexed terminalToken, uint256 index, uint64 nonce)
func (_BPSucker *BPSuckerFilterer) FilterRootToRemote(opts *bind.FilterOpts, root [][32]byte, terminalToken []common.Address) (*BPSuckerRootToRemoteIterator, error) {

	var rootRule []interface{}
	for _, rootItem := range root {
		rootRule = append(rootRule, rootItem)
	}
	var terminalTokenRule []interface{}
	for _, terminalTokenItem := range terminalToken {
		terminalTokenRule = append(terminalTokenRule, terminalTokenItem)
	}

	logs, sub, err := _BPSucker.contract.FilterLogs(opts, "RootToRemote", rootRule, terminalTokenRule)
	if err != nil {
		return nil, err
	}
	return &BPSuckerRootToRemoteIterator{contract: _BPSucker.contract, event: "RootToRemote", logs: logs, sub: sub}, nil
}

// WatchRootToRemote is a free log subscription operation binding the contract event 0x3e069c1f498cba23d7d2b74b84311469e8da6e527cade62061d642c61b64e0de.
//
// Solidity: event RootToRemote(bytes32 indexed root, address indexed terminalToken, uint256 index, uint64 nonce)
func (_BPSucker *BPSuckerFilterer) WatchRootToRemote(opts *bind.WatchOpts, sink chan<- *BPSuckerRootToRemote, root [][32]byte, terminalToken []common.Address) (event.Subscription, error) {

	var rootRule []interface{}
	for _, rootItem := range root {
		rootRule = append(rootRule, rootItem)
	}
	var terminalTokenRule []interface{}
	for _, terminalTokenItem := range terminalToken {
		terminalTokenRule = append(terminalTokenRule, terminalTokenItem)
	}

	logs, sub, err := _BPSucker.contract.WatchLogs(opts, "RootToRemote", rootRule, terminalTokenRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BPSuckerRootToRemote)
				if err := _BPSucker.contract.UnpackLog(event, "RootToRemote", log); err != nil {
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

// ParseRootToRemote is a log parse operation binding the contract event 0x3e069c1f498cba23d7d2b74b84311469e8da6e527cade62061d642c61b64e0de.
//
// Solidity: event RootToRemote(bytes32 indexed root, address indexed terminalToken, uint256 index, uint64 nonce)
func (_BPSucker *BPSuckerFilterer) ParseRootToRemote(log types.Log) (*BPSuckerRootToRemote, error) {
	event := new(BPSuckerRootToRemote)
	if err := _BPSucker.contract.UnpackLog(event, "RootToRemote", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
