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

// BPSuckerDeployerConfig is an auto generated low-level Go binding around an user-defined struct.
type BPSuckerDeployerConfig struct {
	Deployer common.Address
	Mappings []BPTokenMapping
}

// BPTokenMapping is an auto generated low-level Go binding around an user-defined struct.
type BPTokenMapping struct {
	LocalToken      common.Address
	MinGas          uint32
	RemoteToken     common.Address
	MinBridgeAmount *big.Int
}

// BPSuckerRegistryMetaData contains all meta data concerning the BPSuckerRegistry contract.
var BPSuckerRegistryMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"constructor\",\"inputs\":[{\"name\":\"projects\",\"type\":\"address\",\"internalType\":\"contractIJBProjects\"},{\"name\":\"permissions\",\"type\":\"address\",\"internalType\":\"contractIJBPermissions\"},{\"name\":\"_initialOwner\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"PERMISSIONS\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIJBPermissions\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"PROJECTS\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIJBProjects\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"allowSuckerDeployer\",\"inputs\":[{\"name\":\"deployer\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"allowSuckerDeployers\",\"inputs\":[{\"name\":\"deployers\",\"type\":\"address[]\",\"internalType\":\"address[]\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"deploySuckersFor\",\"inputs\":[{\"name\":\"projectId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"salt\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"configurations\",\"type\":\"tuple[]\",\"internalType\":\"structBPSuckerDeployerConfig[]\",\"components\":[{\"name\":\"deployer\",\"type\":\"address\",\"internalType\":\"contractIBPSuckerDeployer\"},{\"name\":\"mappings\",\"type\":\"tuple[]\",\"internalType\":\"structBPTokenMapping[]\",\"components\":[{\"name\":\"localToken\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"minGas\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"remoteToken\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"minBridgeAmount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}]}],\"outputs\":[{\"name\":\"suckers\",\"type\":\"address[]\",\"internalType\":\"address[]\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"isSuckerOf\",\"inputs\":[{\"name\":\"projectId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"suckerAddress\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"jbOwner\",\"inputs\":[],\"outputs\":[{\"name\":\"owner\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"projectId\",\"type\":\"uint88\",\"internalType\":\"uint88\"},{\"name\":\"permissionId\",\"type\":\"uint8\",\"internalType\":\"uint8\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"owner\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"renounceOwnership\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setPermissionId\",\"inputs\":[{\"name\":\"permissionId\",\"type\":\"uint8\",\"internalType\":\"uint8\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"suckerDeployerIsAllowed\",\"inputs\":[{\"name\":\"deployer\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"suckersOf\",\"inputs\":[{\"name\":\"projectId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"address[]\",\"internalType\":\"address[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"transferOwnership\",\"inputs\":[{\"name\":\"newOwner\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"transferOwnershipToProject\",\"inputs\":[{\"name\":\"projectId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"event\",\"name\":\"OwnershipTransferred\",\"inputs\":[{\"name\":\"previousOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"newOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"PermissionIdChanged\",\"inputs\":[{\"name\":\"newIndex\",\"type\":\"uint8\",\"indexed\":false,\"internalType\":\"uint8\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"SuckerDeployerAllowed\",\"inputs\":[{\"name\":\"deployer\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"EnumerableMapNonexistentKey\",\"inputs\":[{\"name\":\"key\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]},{\"type\":\"error\",\"name\":\"INVALID_DEPLOYER\",\"inputs\":[{\"name\":\"deployer\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"INVALID_NEW_OWNER\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"UNAUTHORIZED\",\"inputs\":[]}]",
}

// BPSuckerRegistryABI is the input ABI used to generate the binding from.
// Deprecated: Use BPSuckerRegistryMetaData.ABI instead.
var BPSuckerRegistryABI = BPSuckerRegistryMetaData.ABI

// BPSuckerRegistry is an auto generated Go binding around an Ethereum contract.
type BPSuckerRegistry struct {
	BPSuckerRegistryCaller     // Read-only binding to the contract
	BPSuckerRegistryTransactor // Write-only binding to the contract
	BPSuckerRegistryFilterer   // Log filterer for contract events
}

// BPSuckerRegistryCaller is an auto generated read-only Go binding around an Ethereum contract.
type BPSuckerRegistryCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BPSuckerRegistryTransactor is an auto generated write-only Go binding around an Ethereum contract.
type BPSuckerRegistryTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BPSuckerRegistryFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type BPSuckerRegistryFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BPSuckerRegistrySession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type BPSuckerRegistrySession struct {
	Contract     *BPSuckerRegistry // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// BPSuckerRegistryCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type BPSuckerRegistryCallerSession struct {
	Contract *BPSuckerRegistryCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts           // Call options to use throughout this session
}

// BPSuckerRegistryTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type BPSuckerRegistryTransactorSession struct {
	Contract     *BPSuckerRegistryTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts           // Transaction auth options to use throughout this session
}

// BPSuckerRegistryRaw is an auto generated low-level Go binding around an Ethereum contract.
type BPSuckerRegistryRaw struct {
	Contract *BPSuckerRegistry // Generic contract binding to access the raw methods on
}

// BPSuckerRegistryCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type BPSuckerRegistryCallerRaw struct {
	Contract *BPSuckerRegistryCaller // Generic read-only contract binding to access the raw methods on
}

// BPSuckerRegistryTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type BPSuckerRegistryTransactorRaw struct {
	Contract *BPSuckerRegistryTransactor // Generic write-only contract binding to access the raw methods on
}

// NewBPSuckerRegistry creates a new instance of BPSuckerRegistry, bound to a specific deployed contract.
func NewBPSuckerRegistry(address common.Address, backend bind.ContractBackend) (*BPSuckerRegistry, error) {
	contract, err := bindBPSuckerRegistry(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &BPSuckerRegistry{BPSuckerRegistryCaller: BPSuckerRegistryCaller{contract: contract}, BPSuckerRegistryTransactor: BPSuckerRegistryTransactor{contract: contract}, BPSuckerRegistryFilterer: BPSuckerRegistryFilterer{contract: contract}}, nil
}

// NewBPSuckerRegistryCaller creates a new read-only instance of BPSuckerRegistry, bound to a specific deployed contract.
func NewBPSuckerRegistryCaller(address common.Address, caller bind.ContractCaller) (*BPSuckerRegistryCaller, error) {
	contract, err := bindBPSuckerRegistry(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &BPSuckerRegistryCaller{contract: contract}, nil
}

// NewBPSuckerRegistryTransactor creates a new write-only instance of BPSuckerRegistry, bound to a specific deployed contract.
func NewBPSuckerRegistryTransactor(address common.Address, transactor bind.ContractTransactor) (*BPSuckerRegistryTransactor, error) {
	contract, err := bindBPSuckerRegistry(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &BPSuckerRegistryTransactor{contract: contract}, nil
}

// NewBPSuckerRegistryFilterer creates a new log filterer instance of BPSuckerRegistry, bound to a specific deployed contract.
func NewBPSuckerRegistryFilterer(address common.Address, filterer bind.ContractFilterer) (*BPSuckerRegistryFilterer, error) {
	contract, err := bindBPSuckerRegistry(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &BPSuckerRegistryFilterer{contract: contract}, nil
}

// bindBPSuckerRegistry binds a generic wrapper to an already deployed contract.
func bindBPSuckerRegistry(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := BPSuckerRegistryMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_BPSuckerRegistry *BPSuckerRegistryRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _BPSuckerRegistry.Contract.BPSuckerRegistryCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_BPSuckerRegistry *BPSuckerRegistryRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BPSuckerRegistry.Contract.BPSuckerRegistryTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_BPSuckerRegistry *BPSuckerRegistryRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _BPSuckerRegistry.Contract.BPSuckerRegistryTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_BPSuckerRegistry *BPSuckerRegistryCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _BPSuckerRegistry.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_BPSuckerRegistry *BPSuckerRegistryTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BPSuckerRegistry.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_BPSuckerRegistry *BPSuckerRegistryTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _BPSuckerRegistry.Contract.contract.Transact(opts, method, params...)
}

// PERMISSIONS is a free data retrieval call binding the contract method 0xf434c914.
//
// Solidity: function PERMISSIONS() view returns(address)
func (_BPSuckerRegistry *BPSuckerRegistryCaller) PERMISSIONS(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _BPSuckerRegistry.contract.Call(opts, &out, "PERMISSIONS")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// PERMISSIONS is a free data retrieval call binding the contract method 0xf434c914.
//
// Solidity: function PERMISSIONS() view returns(address)
func (_BPSuckerRegistry *BPSuckerRegistrySession) PERMISSIONS() (common.Address, error) {
	return _BPSuckerRegistry.Contract.PERMISSIONS(&_BPSuckerRegistry.CallOpts)
}

// PERMISSIONS is a free data retrieval call binding the contract method 0xf434c914.
//
// Solidity: function PERMISSIONS() view returns(address)
func (_BPSuckerRegistry *BPSuckerRegistryCallerSession) PERMISSIONS() (common.Address, error) {
	return _BPSuckerRegistry.Contract.PERMISSIONS(&_BPSuckerRegistry.CallOpts)
}

// PROJECTS is a free data retrieval call binding the contract method 0x293c4999.
//
// Solidity: function PROJECTS() view returns(address)
func (_BPSuckerRegistry *BPSuckerRegistryCaller) PROJECTS(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _BPSuckerRegistry.contract.Call(opts, &out, "PROJECTS")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// PROJECTS is a free data retrieval call binding the contract method 0x293c4999.
//
// Solidity: function PROJECTS() view returns(address)
func (_BPSuckerRegistry *BPSuckerRegistrySession) PROJECTS() (common.Address, error) {
	return _BPSuckerRegistry.Contract.PROJECTS(&_BPSuckerRegistry.CallOpts)
}

// PROJECTS is a free data retrieval call binding the contract method 0x293c4999.
//
// Solidity: function PROJECTS() view returns(address)
func (_BPSuckerRegistry *BPSuckerRegistryCallerSession) PROJECTS() (common.Address, error) {
	return _BPSuckerRegistry.Contract.PROJECTS(&_BPSuckerRegistry.CallOpts)
}

// IsSuckerOf is a free data retrieval call binding the contract method 0x83db9d01.
//
// Solidity: function isSuckerOf(uint256 projectId, address suckerAddress) view returns(bool)
func (_BPSuckerRegistry *BPSuckerRegistryCaller) IsSuckerOf(opts *bind.CallOpts, projectId *big.Int, suckerAddress common.Address) (bool, error) {
	var out []interface{}
	err := _BPSuckerRegistry.contract.Call(opts, &out, "isSuckerOf", projectId, suckerAddress)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsSuckerOf is a free data retrieval call binding the contract method 0x83db9d01.
//
// Solidity: function isSuckerOf(uint256 projectId, address suckerAddress) view returns(bool)
func (_BPSuckerRegistry *BPSuckerRegistrySession) IsSuckerOf(projectId *big.Int, suckerAddress common.Address) (bool, error) {
	return _BPSuckerRegistry.Contract.IsSuckerOf(&_BPSuckerRegistry.CallOpts, projectId, suckerAddress)
}

// IsSuckerOf is a free data retrieval call binding the contract method 0x83db9d01.
//
// Solidity: function isSuckerOf(uint256 projectId, address suckerAddress) view returns(bool)
func (_BPSuckerRegistry *BPSuckerRegistryCallerSession) IsSuckerOf(projectId *big.Int, suckerAddress common.Address) (bool, error) {
	return _BPSuckerRegistry.Contract.IsSuckerOf(&_BPSuckerRegistry.CallOpts, projectId, suckerAddress)
}

// JbOwner is a free data retrieval call binding the contract method 0xba23c36e.
//
// Solidity: function jbOwner() view returns(address owner, uint88 projectId, uint8 permissionId)
func (_BPSuckerRegistry *BPSuckerRegistryCaller) JbOwner(opts *bind.CallOpts) (struct {
	Owner        common.Address
	ProjectId    *big.Int
	PermissionId uint8
}, error) {
	var out []interface{}
	err := _BPSuckerRegistry.contract.Call(opts, &out, "jbOwner")

	outstruct := new(struct {
		Owner        common.Address
		ProjectId    *big.Int
		PermissionId uint8
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Owner = *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	outstruct.ProjectId = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.PermissionId = *abi.ConvertType(out[2], new(uint8)).(*uint8)

	return *outstruct, err

}

// JbOwner is a free data retrieval call binding the contract method 0xba23c36e.
//
// Solidity: function jbOwner() view returns(address owner, uint88 projectId, uint8 permissionId)
func (_BPSuckerRegistry *BPSuckerRegistrySession) JbOwner() (struct {
	Owner        common.Address
	ProjectId    *big.Int
	PermissionId uint8
}, error) {
	return _BPSuckerRegistry.Contract.JbOwner(&_BPSuckerRegistry.CallOpts)
}

// JbOwner is a free data retrieval call binding the contract method 0xba23c36e.
//
// Solidity: function jbOwner() view returns(address owner, uint88 projectId, uint8 permissionId)
func (_BPSuckerRegistry *BPSuckerRegistryCallerSession) JbOwner() (struct {
	Owner        common.Address
	ProjectId    *big.Int
	PermissionId uint8
}, error) {
	return _BPSuckerRegistry.Contract.JbOwner(&_BPSuckerRegistry.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_BPSuckerRegistry *BPSuckerRegistryCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _BPSuckerRegistry.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_BPSuckerRegistry *BPSuckerRegistrySession) Owner() (common.Address, error) {
	return _BPSuckerRegistry.Contract.Owner(&_BPSuckerRegistry.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_BPSuckerRegistry *BPSuckerRegistryCallerSession) Owner() (common.Address, error) {
	return _BPSuckerRegistry.Contract.Owner(&_BPSuckerRegistry.CallOpts)
}

// SuckerDeployerIsAllowed is a free data retrieval call binding the contract method 0xa1031a4f.
//
// Solidity: function suckerDeployerIsAllowed(address deployer) view returns(bool)
func (_BPSuckerRegistry *BPSuckerRegistryCaller) SuckerDeployerIsAllowed(opts *bind.CallOpts, deployer common.Address) (bool, error) {
	var out []interface{}
	err := _BPSuckerRegistry.contract.Call(opts, &out, "suckerDeployerIsAllowed", deployer)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SuckerDeployerIsAllowed is a free data retrieval call binding the contract method 0xa1031a4f.
//
// Solidity: function suckerDeployerIsAllowed(address deployer) view returns(bool)
func (_BPSuckerRegistry *BPSuckerRegistrySession) SuckerDeployerIsAllowed(deployer common.Address) (bool, error) {
	return _BPSuckerRegistry.Contract.SuckerDeployerIsAllowed(&_BPSuckerRegistry.CallOpts, deployer)
}

// SuckerDeployerIsAllowed is a free data retrieval call binding the contract method 0xa1031a4f.
//
// Solidity: function suckerDeployerIsAllowed(address deployer) view returns(bool)
func (_BPSuckerRegistry *BPSuckerRegistryCallerSession) SuckerDeployerIsAllowed(deployer common.Address) (bool, error) {
	return _BPSuckerRegistry.Contract.SuckerDeployerIsAllowed(&_BPSuckerRegistry.CallOpts, deployer)
}

// SuckersOf is a free data retrieval call binding the contract method 0x3d12cb87.
//
// Solidity: function suckersOf(uint256 projectId) view returns(address[])
func (_BPSuckerRegistry *BPSuckerRegistryCaller) SuckersOf(opts *bind.CallOpts, projectId *big.Int) ([]common.Address, error) {
	var out []interface{}
	err := _BPSuckerRegistry.contract.Call(opts, &out, "suckersOf", projectId)

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// SuckersOf is a free data retrieval call binding the contract method 0x3d12cb87.
//
// Solidity: function suckersOf(uint256 projectId) view returns(address[])
func (_BPSuckerRegistry *BPSuckerRegistrySession) SuckersOf(projectId *big.Int) ([]common.Address, error) {
	return _BPSuckerRegistry.Contract.SuckersOf(&_BPSuckerRegistry.CallOpts, projectId)
}

// SuckersOf is a free data retrieval call binding the contract method 0x3d12cb87.
//
// Solidity: function suckersOf(uint256 projectId) view returns(address[])
func (_BPSuckerRegistry *BPSuckerRegistryCallerSession) SuckersOf(projectId *big.Int) ([]common.Address, error) {
	return _BPSuckerRegistry.Contract.SuckersOf(&_BPSuckerRegistry.CallOpts, projectId)
}

// AllowSuckerDeployer is a paid mutator transaction binding the contract method 0xf73ac20f.
//
// Solidity: function allowSuckerDeployer(address deployer) returns()
func (_BPSuckerRegistry *BPSuckerRegistryTransactor) AllowSuckerDeployer(opts *bind.TransactOpts, deployer common.Address) (*types.Transaction, error) {
	return _BPSuckerRegistry.contract.Transact(opts, "allowSuckerDeployer", deployer)
}

// AllowSuckerDeployer is a paid mutator transaction binding the contract method 0xf73ac20f.
//
// Solidity: function allowSuckerDeployer(address deployer) returns()
func (_BPSuckerRegistry *BPSuckerRegistrySession) AllowSuckerDeployer(deployer common.Address) (*types.Transaction, error) {
	return _BPSuckerRegistry.Contract.AllowSuckerDeployer(&_BPSuckerRegistry.TransactOpts, deployer)
}

// AllowSuckerDeployer is a paid mutator transaction binding the contract method 0xf73ac20f.
//
// Solidity: function allowSuckerDeployer(address deployer) returns()
func (_BPSuckerRegistry *BPSuckerRegistryTransactorSession) AllowSuckerDeployer(deployer common.Address) (*types.Transaction, error) {
	return _BPSuckerRegistry.Contract.AllowSuckerDeployer(&_BPSuckerRegistry.TransactOpts, deployer)
}

// AllowSuckerDeployers is a paid mutator transaction binding the contract method 0x669745a5.
//
// Solidity: function allowSuckerDeployers(address[] deployers) returns()
func (_BPSuckerRegistry *BPSuckerRegistryTransactor) AllowSuckerDeployers(opts *bind.TransactOpts, deployers []common.Address) (*types.Transaction, error) {
	return _BPSuckerRegistry.contract.Transact(opts, "allowSuckerDeployers", deployers)
}

// AllowSuckerDeployers is a paid mutator transaction binding the contract method 0x669745a5.
//
// Solidity: function allowSuckerDeployers(address[] deployers) returns()
func (_BPSuckerRegistry *BPSuckerRegistrySession) AllowSuckerDeployers(deployers []common.Address) (*types.Transaction, error) {
	return _BPSuckerRegistry.Contract.AllowSuckerDeployers(&_BPSuckerRegistry.TransactOpts, deployers)
}

// AllowSuckerDeployers is a paid mutator transaction binding the contract method 0x669745a5.
//
// Solidity: function allowSuckerDeployers(address[] deployers) returns()
func (_BPSuckerRegistry *BPSuckerRegistryTransactorSession) AllowSuckerDeployers(deployers []common.Address) (*types.Transaction, error) {
	return _BPSuckerRegistry.Contract.AllowSuckerDeployers(&_BPSuckerRegistry.TransactOpts, deployers)
}

// DeploySuckersFor is a paid mutator transaction binding the contract method 0x2f3f50e0.
//
// Solidity: function deploySuckersFor(uint256 projectId, bytes32 salt, (address,(address,uint32,address,uint256)[])[] configurations) returns(address[] suckers)
func (_BPSuckerRegistry *BPSuckerRegistryTransactor) DeploySuckersFor(opts *bind.TransactOpts, projectId *big.Int, salt [32]byte, configurations []BPSuckerDeployerConfig) (*types.Transaction, error) {
	return _BPSuckerRegistry.contract.Transact(opts, "deploySuckersFor", projectId, salt, configurations)
}

// DeploySuckersFor is a paid mutator transaction binding the contract method 0x2f3f50e0.
//
// Solidity: function deploySuckersFor(uint256 projectId, bytes32 salt, (address,(address,uint32,address,uint256)[])[] configurations) returns(address[] suckers)
func (_BPSuckerRegistry *BPSuckerRegistrySession) DeploySuckersFor(projectId *big.Int, salt [32]byte, configurations []BPSuckerDeployerConfig) (*types.Transaction, error) {
	return _BPSuckerRegistry.Contract.DeploySuckersFor(&_BPSuckerRegistry.TransactOpts, projectId, salt, configurations)
}

// DeploySuckersFor is a paid mutator transaction binding the contract method 0x2f3f50e0.
//
// Solidity: function deploySuckersFor(uint256 projectId, bytes32 salt, (address,(address,uint32,address,uint256)[])[] configurations) returns(address[] suckers)
func (_BPSuckerRegistry *BPSuckerRegistryTransactorSession) DeploySuckersFor(projectId *big.Int, salt [32]byte, configurations []BPSuckerDeployerConfig) (*types.Transaction, error) {
	return _BPSuckerRegistry.Contract.DeploySuckersFor(&_BPSuckerRegistry.TransactOpts, projectId, salt, configurations)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_BPSuckerRegistry *BPSuckerRegistryTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BPSuckerRegistry.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_BPSuckerRegistry *BPSuckerRegistrySession) RenounceOwnership() (*types.Transaction, error) {
	return _BPSuckerRegistry.Contract.RenounceOwnership(&_BPSuckerRegistry.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_BPSuckerRegistry *BPSuckerRegistryTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _BPSuckerRegistry.Contract.RenounceOwnership(&_BPSuckerRegistry.TransactOpts)
}

// SetPermissionId is a paid mutator transaction binding the contract method 0xc0fb5e85.
//
// Solidity: function setPermissionId(uint8 permissionId) returns()
func (_BPSuckerRegistry *BPSuckerRegistryTransactor) SetPermissionId(opts *bind.TransactOpts, permissionId uint8) (*types.Transaction, error) {
	return _BPSuckerRegistry.contract.Transact(opts, "setPermissionId", permissionId)
}

// SetPermissionId is a paid mutator transaction binding the contract method 0xc0fb5e85.
//
// Solidity: function setPermissionId(uint8 permissionId) returns()
func (_BPSuckerRegistry *BPSuckerRegistrySession) SetPermissionId(permissionId uint8) (*types.Transaction, error) {
	return _BPSuckerRegistry.Contract.SetPermissionId(&_BPSuckerRegistry.TransactOpts, permissionId)
}

// SetPermissionId is a paid mutator transaction binding the contract method 0xc0fb5e85.
//
// Solidity: function setPermissionId(uint8 permissionId) returns()
func (_BPSuckerRegistry *BPSuckerRegistryTransactorSession) SetPermissionId(permissionId uint8) (*types.Transaction, error) {
	return _BPSuckerRegistry.Contract.SetPermissionId(&_BPSuckerRegistry.TransactOpts, permissionId)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_BPSuckerRegistry *BPSuckerRegistryTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _BPSuckerRegistry.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_BPSuckerRegistry *BPSuckerRegistrySession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _BPSuckerRegistry.Contract.TransferOwnership(&_BPSuckerRegistry.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_BPSuckerRegistry *BPSuckerRegistryTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _BPSuckerRegistry.Contract.TransferOwnership(&_BPSuckerRegistry.TransactOpts, newOwner)
}

// TransferOwnershipToProject is a paid mutator transaction binding the contract method 0xa220d696.
//
// Solidity: function transferOwnershipToProject(uint256 projectId) returns()
func (_BPSuckerRegistry *BPSuckerRegistryTransactor) TransferOwnershipToProject(opts *bind.TransactOpts, projectId *big.Int) (*types.Transaction, error) {
	return _BPSuckerRegistry.contract.Transact(opts, "transferOwnershipToProject", projectId)
}

// TransferOwnershipToProject is a paid mutator transaction binding the contract method 0xa220d696.
//
// Solidity: function transferOwnershipToProject(uint256 projectId) returns()
func (_BPSuckerRegistry *BPSuckerRegistrySession) TransferOwnershipToProject(projectId *big.Int) (*types.Transaction, error) {
	return _BPSuckerRegistry.Contract.TransferOwnershipToProject(&_BPSuckerRegistry.TransactOpts, projectId)
}

// TransferOwnershipToProject is a paid mutator transaction binding the contract method 0xa220d696.
//
// Solidity: function transferOwnershipToProject(uint256 projectId) returns()
func (_BPSuckerRegistry *BPSuckerRegistryTransactorSession) TransferOwnershipToProject(projectId *big.Int) (*types.Transaction, error) {
	return _BPSuckerRegistry.Contract.TransferOwnershipToProject(&_BPSuckerRegistry.TransactOpts, projectId)
}

// BPSuckerRegistryOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the BPSuckerRegistry contract.
type BPSuckerRegistryOwnershipTransferredIterator struct {
	Event *BPSuckerRegistryOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *BPSuckerRegistryOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BPSuckerRegistryOwnershipTransferred)
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
		it.Event = new(BPSuckerRegistryOwnershipTransferred)
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
func (it *BPSuckerRegistryOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BPSuckerRegistryOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BPSuckerRegistryOwnershipTransferred represents a OwnershipTransferred event raised by the BPSuckerRegistry contract.
type BPSuckerRegistryOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_BPSuckerRegistry *BPSuckerRegistryFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*BPSuckerRegistryOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _BPSuckerRegistry.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &BPSuckerRegistryOwnershipTransferredIterator{contract: _BPSuckerRegistry.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_BPSuckerRegistry *BPSuckerRegistryFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *BPSuckerRegistryOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _BPSuckerRegistry.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BPSuckerRegistryOwnershipTransferred)
				if err := _BPSuckerRegistry.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_BPSuckerRegistry *BPSuckerRegistryFilterer) ParseOwnershipTransferred(log types.Log) (*BPSuckerRegistryOwnershipTransferred, error) {
	event := new(BPSuckerRegistryOwnershipTransferred)
	if err := _BPSuckerRegistry.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BPSuckerRegistryPermissionIdChangedIterator is returned from FilterPermissionIdChanged and is used to iterate over the raw logs and unpacked data for PermissionIdChanged events raised by the BPSuckerRegistry contract.
type BPSuckerRegistryPermissionIdChangedIterator struct {
	Event *BPSuckerRegistryPermissionIdChanged // Event containing the contract specifics and raw log

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
func (it *BPSuckerRegistryPermissionIdChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BPSuckerRegistryPermissionIdChanged)
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
		it.Event = new(BPSuckerRegistryPermissionIdChanged)
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
func (it *BPSuckerRegistryPermissionIdChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BPSuckerRegistryPermissionIdChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BPSuckerRegistryPermissionIdChanged represents a PermissionIdChanged event raised by the BPSuckerRegistry contract.
type BPSuckerRegistryPermissionIdChanged struct {
	NewIndex uint8
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterPermissionIdChanged is a free log retrieval operation binding the contract event 0x6f4725627084d54c46c628b050805457a1cd38c5cd20c1f00132960f8b546fed.
//
// Solidity: event PermissionIdChanged(uint8 newIndex)
func (_BPSuckerRegistry *BPSuckerRegistryFilterer) FilterPermissionIdChanged(opts *bind.FilterOpts) (*BPSuckerRegistryPermissionIdChangedIterator, error) {

	logs, sub, err := _BPSuckerRegistry.contract.FilterLogs(opts, "PermissionIdChanged")
	if err != nil {
		return nil, err
	}
	return &BPSuckerRegistryPermissionIdChangedIterator{contract: _BPSuckerRegistry.contract, event: "PermissionIdChanged", logs: logs, sub: sub}, nil
}

// WatchPermissionIdChanged is a free log subscription operation binding the contract event 0x6f4725627084d54c46c628b050805457a1cd38c5cd20c1f00132960f8b546fed.
//
// Solidity: event PermissionIdChanged(uint8 newIndex)
func (_BPSuckerRegistry *BPSuckerRegistryFilterer) WatchPermissionIdChanged(opts *bind.WatchOpts, sink chan<- *BPSuckerRegistryPermissionIdChanged) (event.Subscription, error) {

	logs, sub, err := _BPSuckerRegistry.contract.WatchLogs(opts, "PermissionIdChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BPSuckerRegistryPermissionIdChanged)
				if err := _BPSuckerRegistry.contract.UnpackLog(event, "PermissionIdChanged", log); err != nil {
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

// ParsePermissionIdChanged is a log parse operation binding the contract event 0x6f4725627084d54c46c628b050805457a1cd38c5cd20c1f00132960f8b546fed.
//
// Solidity: event PermissionIdChanged(uint8 newIndex)
func (_BPSuckerRegistry *BPSuckerRegistryFilterer) ParsePermissionIdChanged(log types.Log) (*BPSuckerRegistryPermissionIdChanged, error) {
	event := new(BPSuckerRegistryPermissionIdChanged)
	if err := _BPSuckerRegistry.contract.UnpackLog(event, "PermissionIdChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BPSuckerRegistrySuckerDeployerAllowedIterator is returned from FilterSuckerDeployerAllowed and is used to iterate over the raw logs and unpacked data for SuckerDeployerAllowed events raised by the BPSuckerRegistry contract.
type BPSuckerRegistrySuckerDeployerAllowedIterator struct {
	Event *BPSuckerRegistrySuckerDeployerAllowed // Event containing the contract specifics and raw log

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
func (it *BPSuckerRegistrySuckerDeployerAllowedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BPSuckerRegistrySuckerDeployerAllowed)
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
		it.Event = new(BPSuckerRegistrySuckerDeployerAllowed)
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
func (it *BPSuckerRegistrySuckerDeployerAllowedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BPSuckerRegistrySuckerDeployerAllowedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BPSuckerRegistrySuckerDeployerAllowed represents a SuckerDeployerAllowed event raised by the BPSuckerRegistry contract.
type BPSuckerRegistrySuckerDeployerAllowed struct {
	Deployer common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterSuckerDeployerAllowed is a free log retrieval operation binding the contract event 0xf6e6facda29906a01d337146a9b928ac09b423901f28e0e001d43550448b370b.
//
// Solidity: event SuckerDeployerAllowed(address deployer)
func (_BPSuckerRegistry *BPSuckerRegistryFilterer) FilterSuckerDeployerAllowed(opts *bind.FilterOpts) (*BPSuckerRegistrySuckerDeployerAllowedIterator, error) {

	logs, sub, err := _BPSuckerRegistry.contract.FilterLogs(opts, "SuckerDeployerAllowed")
	if err != nil {
		return nil, err
	}
	return &BPSuckerRegistrySuckerDeployerAllowedIterator{contract: _BPSuckerRegistry.contract, event: "SuckerDeployerAllowed", logs: logs, sub: sub}, nil
}

// WatchSuckerDeployerAllowed is a free log subscription operation binding the contract event 0xf6e6facda29906a01d337146a9b928ac09b423901f28e0e001d43550448b370b.
//
// Solidity: event SuckerDeployerAllowed(address deployer)
func (_BPSuckerRegistry *BPSuckerRegistryFilterer) WatchSuckerDeployerAllowed(opts *bind.WatchOpts, sink chan<- *BPSuckerRegistrySuckerDeployerAllowed) (event.Subscription, error) {

	logs, sub, err := _BPSuckerRegistry.contract.WatchLogs(opts, "SuckerDeployerAllowed")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BPSuckerRegistrySuckerDeployerAllowed)
				if err := _BPSuckerRegistry.contract.UnpackLog(event, "SuckerDeployerAllowed", log); err != nil {
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

// ParseSuckerDeployerAllowed is a log parse operation binding the contract event 0xf6e6facda29906a01d337146a9b928ac09b423901f28e0e001d43550448b370b.
//
// Solidity: event SuckerDeployerAllowed(address deployer)
func (_BPSuckerRegistry *BPSuckerRegistryFilterer) ParseSuckerDeployerAllowed(log types.Log) (*BPSuckerRegistrySuckerDeployerAllowed, error) {
	event := new(BPSuckerRegistrySuckerDeployerAllowed)
	if err := _BPSuckerRegistry.contract.UnpackLog(event, "SuckerDeployerAllowed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
