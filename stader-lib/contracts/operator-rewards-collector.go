// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package contracts

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

// OperatorRewardsCollectorMetaData contains all meta data concerning the OperatorRewardsCollector contract.
var OperatorRewardsCollectorMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"name\":\"TransferFailed\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ZeroAddress\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Claimed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"DepositedFor\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"version\",\"type\":\"uint8\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Paused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"previousAdminRole\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"newAdminRole\",\"type\":\"bytes32\"}],\"name\":\"RoleAdminChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleGranted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleRevoked\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Unpaused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"staderConfig\",\"type\":\"address\"}],\"name\":\"UpdatedStaderConfig\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"DEFAULT_ADMIN_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"balances\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"claim\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_receiver\",\"type\":\"address\"}],\"name\":\"depositFor\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"}],\"name\":\"getRoleAdmin\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"grantRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"hasRole\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_admin\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_staderConfig\",\"type\":\"address\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"paused\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"renounceRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"revokeRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"staderConfig\",\"outputs\":[{\"internalType\":\"contractIStaderConfig\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// OperatorRewardsCollectorABI is the input ABI used to generate the binding from.
// Deprecated: Use OperatorRewardsCollectorMetaData.ABI instead.
var OperatorRewardsCollectorABI = OperatorRewardsCollectorMetaData.ABI

// OperatorRewardsCollector is an auto generated Go binding around an Ethereum contract.
type OperatorRewardsCollector struct {
	OperatorRewardsCollectorCaller     // Read-only binding to the contract
	OperatorRewardsCollectorTransactor // Write-only binding to the contract
	OperatorRewardsCollectorFilterer   // Log filterer for contract events
}

// OperatorRewardsCollectorCaller is an auto generated read-only Go binding around an Ethereum contract.
type OperatorRewardsCollectorCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OperatorRewardsCollectorTransactor is an auto generated write-only Go binding around an Ethereum contract.
type OperatorRewardsCollectorTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OperatorRewardsCollectorFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type OperatorRewardsCollectorFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OperatorRewardsCollectorSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type OperatorRewardsCollectorSession struct {
	Contract     *OperatorRewardsCollector // Generic contract binding to set the session for
	CallOpts     bind.CallOpts             // Call options to use throughout this session
	TransactOpts bind.TransactOpts         // Transaction auth options to use throughout this session
}

// OperatorRewardsCollectorCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type OperatorRewardsCollectorCallerSession struct {
	Contract *OperatorRewardsCollectorCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                   // Call options to use throughout this session
}

// OperatorRewardsCollectorTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type OperatorRewardsCollectorTransactorSession struct {
	Contract     *OperatorRewardsCollectorTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                   // Transaction auth options to use throughout this session
}

// OperatorRewardsCollectorRaw is an auto generated low-level Go binding around an Ethereum contract.
type OperatorRewardsCollectorRaw struct {
	Contract *OperatorRewardsCollector // Generic contract binding to access the raw methods on
}

// OperatorRewardsCollectorCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type OperatorRewardsCollectorCallerRaw struct {
	Contract *OperatorRewardsCollectorCaller // Generic read-only contract binding to access the raw methods on
}

// OperatorRewardsCollectorTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type OperatorRewardsCollectorTransactorRaw struct {
	Contract *OperatorRewardsCollectorTransactor // Generic write-only contract binding to access the raw methods on
}

// NewOperatorRewardsCollector creates a new instance of OperatorRewardsCollector, bound to a specific deployed contract.
func NewOperatorRewardsCollector(address common.Address, backend bind.ContractBackend) (*OperatorRewardsCollector, error) {
	contract, err := bindOperatorRewardsCollector(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &OperatorRewardsCollector{OperatorRewardsCollectorCaller: OperatorRewardsCollectorCaller{contract: contract}, OperatorRewardsCollectorTransactor: OperatorRewardsCollectorTransactor{contract: contract}, OperatorRewardsCollectorFilterer: OperatorRewardsCollectorFilterer{contract: contract}}, nil
}

// NewOperatorRewardsCollectorCaller creates a new read-only instance of OperatorRewardsCollector, bound to a specific deployed contract.
func NewOperatorRewardsCollectorCaller(address common.Address, caller bind.ContractCaller) (*OperatorRewardsCollectorCaller, error) {
	contract, err := bindOperatorRewardsCollector(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &OperatorRewardsCollectorCaller{contract: contract}, nil
}

// NewOperatorRewardsCollectorTransactor creates a new write-only instance of OperatorRewardsCollector, bound to a specific deployed contract.
func NewOperatorRewardsCollectorTransactor(address common.Address, transactor bind.ContractTransactor) (*OperatorRewardsCollectorTransactor, error) {
	contract, err := bindOperatorRewardsCollector(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &OperatorRewardsCollectorTransactor{contract: contract}, nil
}

// NewOperatorRewardsCollectorFilterer creates a new log filterer instance of OperatorRewardsCollector, bound to a specific deployed contract.
func NewOperatorRewardsCollectorFilterer(address common.Address, filterer bind.ContractFilterer) (*OperatorRewardsCollectorFilterer, error) {
	contract, err := bindOperatorRewardsCollector(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &OperatorRewardsCollectorFilterer{contract: contract}, nil
}

// bindOperatorRewardsCollector binds a generic wrapper to an already deployed contract.
func bindOperatorRewardsCollector(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := OperatorRewardsCollectorMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_OperatorRewardsCollector *OperatorRewardsCollectorRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _OperatorRewardsCollector.Contract.OperatorRewardsCollectorCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_OperatorRewardsCollector *OperatorRewardsCollectorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _OperatorRewardsCollector.Contract.OperatorRewardsCollectorTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_OperatorRewardsCollector *OperatorRewardsCollectorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _OperatorRewardsCollector.Contract.OperatorRewardsCollectorTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_OperatorRewardsCollector *OperatorRewardsCollectorCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _OperatorRewardsCollector.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_OperatorRewardsCollector *OperatorRewardsCollectorTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _OperatorRewardsCollector.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_OperatorRewardsCollector *OperatorRewardsCollectorTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _OperatorRewardsCollector.Contract.contract.Transact(opts, method, params...)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_OperatorRewardsCollector *OperatorRewardsCollectorCaller) DEFAULTADMINROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _OperatorRewardsCollector.contract.Call(opts, &out, "DEFAULT_ADMIN_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_OperatorRewardsCollector *OperatorRewardsCollectorSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _OperatorRewardsCollector.Contract.DEFAULTADMINROLE(&_OperatorRewardsCollector.CallOpts)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_OperatorRewardsCollector *OperatorRewardsCollectorCallerSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _OperatorRewardsCollector.Contract.DEFAULTADMINROLE(&_OperatorRewardsCollector.CallOpts)
}

// Balances is a free data retrieval call binding the contract method 0x27e235e3.
//
// Solidity: function balances(address ) view returns(uint256)
func (_OperatorRewardsCollector *OperatorRewardsCollectorCaller) Balances(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _OperatorRewardsCollector.contract.Call(opts, &out, "balances", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Balances is a free data retrieval call binding the contract method 0x27e235e3.
//
// Solidity: function balances(address ) view returns(uint256)
func (_OperatorRewardsCollector *OperatorRewardsCollectorSession) Balances(arg0 common.Address) (*big.Int, error) {
	return _OperatorRewardsCollector.Contract.Balances(&_OperatorRewardsCollector.CallOpts, arg0)
}

// Balances is a free data retrieval call binding the contract method 0x27e235e3.
//
// Solidity: function balances(address ) view returns(uint256)
func (_OperatorRewardsCollector *OperatorRewardsCollectorCallerSession) Balances(arg0 common.Address) (*big.Int, error) {
	return _OperatorRewardsCollector.Contract.Balances(&_OperatorRewardsCollector.CallOpts, arg0)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_OperatorRewardsCollector *OperatorRewardsCollectorCaller) GetRoleAdmin(opts *bind.CallOpts, role [32]byte) ([32]byte, error) {
	var out []interface{}
	err := _OperatorRewardsCollector.contract.Call(opts, &out, "getRoleAdmin", role)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_OperatorRewardsCollector *OperatorRewardsCollectorSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _OperatorRewardsCollector.Contract.GetRoleAdmin(&_OperatorRewardsCollector.CallOpts, role)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_OperatorRewardsCollector *OperatorRewardsCollectorCallerSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _OperatorRewardsCollector.Contract.GetRoleAdmin(&_OperatorRewardsCollector.CallOpts, role)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_OperatorRewardsCollector *OperatorRewardsCollectorCaller) HasRole(opts *bind.CallOpts, role [32]byte, account common.Address) (bool, error) {
	var out []interface{}
	err := _OperatorRewardsCollector.contract.Call(opts, &out, "hasRole", role, account)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_OperatorRewardsCollector *OperatorRewardsCollectorSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _OperatorRewardsCollector.Contract.HasRole(&_OperatorRewardsCollector.CallOpts, role, account)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_OperatorRewardsCollector *OperatorRewardsCollectorCallerSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _OperatorRewardsCollector.Contract.HasRole(&_OperatorRewardsCollector.CallOpts, role, account)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_OperatorRewardsCollector *OperatorRewardsCollectorCaller) Paused(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _OperatorRewardsCollector.contract.Call(opts, &out, "paused")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_OperatorRewardsCollector *OperatorRewardsCollectorSession) Paused() (bool, error) {
	return _OperatorRewardsCollector.Contract.Paused(&_OperatorRewardsCollector.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_OperatorRewardsCollector *OperatorRewardsCollectorCallerSession) Paused() (bool, error) {
	return _OperatorRewardsCollector.Contract.Paused(&_OperatorRewardsCollector.CallOpts)
}

// StaderConfig is a free data retrieval call binding the contract method 0x490ffa35.
//
// Solidity: function staderConfig() view returns(address)
func (_OperatorRewardsCollector *OperatorRewardsCollectorCaller) StaderConfig(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _OperatorRewardsCollector.contract.Call(opts, &out, "staderConfig")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// StaderConfig is a free data retrieval call binding the contract method 0x490ffa35.
//
// Solidity: function staderConfig() view returns(address)
func (_OperatorRewardsCollector *OperatorRewardsCollectorSession) StaderConfig() (common.Address, error) {
	return _OperatorRewardsCollector.Contract.StaderConfig(&_OperatorRewardsCollector.CallOpts)
}

// StaderConfig is a free data retrieval call binding the contract method 0x490ffa35.
//
// Solidity: function staderConfig() view returns(address)
func (_OperatorRewardsCollector *OperatorRewardsCollectorCallerSession) StaderConfig() (common.Address, error) {
	return _OperatorRewardsCollector.Contract.StaderConfig(&_OperatorRewardsCollector.CallOpts)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_OperatorRewardsCollector *OperatorRewardsCollectorCaller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _OperatorRewardsCollector.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_OperatorRewardsCollector *OperatorRewardsCollectorSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _OperatorRewardsCollector.Contract.SupportsInterface(&_OperatorRewardsCollector.CallOpts, interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_OperatorRewardsCollector *OperatorRewardsCollectorCallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _OperatorRewardsCollector.Contract.SupportsInterface(&_OperatorRewardsCollector.CallOpts, interfaceId)
}

// Claim is a paid mutator transaction binding the contract method 0x4e71d92d.
//
// Solidity: function claim() returns()
func (_OperatorRewardsCollector *OperatorRewardsCollectorTransactor) Claim(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _OperatorRewardsCollector.contract.Transact(opts, "claim")
}

// Claim is a paid mutator transaction binding the contract method 0x4e71d92d.
//
// Solidity: function claim() returns()
func (_OperatorRewardsCollector *OperatorRewardsCollectorSession) Claim() (*types.Transaction, error) {
	return _OperatorRewardsCollector.Contract.Claim(&_OperatorRewardsCollector.TransactOpts)
}

// Claim is a paid mutator transaction binding the contract method 0x4e71d92d.
//
// Solidity: function claim() returns()
func (_OperatorRewardsCollector *OperatorRewardsCollectorTransactorSession) Claim() (*types.Transaction, error) {
	return _OperatorRewardsCollector.Contract.Claim(&_OperatorRewardsCollector.TransactOpts)
}

// DepositFor is a paid mutator transaction binding the contract method 0xaa67c919.
//
// Solidity: function depositFor(address _receiver) payable returns()
func (_OperatorRewardsCollector *OperatorRewardsCollectorTransactor) DepositFor(opts *bind.TransactOpts, _receiver common.Address) (*types.Transaction, error) {
	return _OperatorRewardsCollector.contract.Transact(opts, "depositFor", _receiver)
}

// DepositFor is a paid mutator transaction binding the contract method 0xaa67c919.
//
// Solidity: function depositFor(address _receiver) payable returns()
func (_OperatorRewardsCollector *OperatorRewardsCollectorSession) DepositFor(_receiver common.Address) (*types.Transaction, error) {
	return _OperatorRewardsCollector.Contract.DepositFor(&_OperatorRewardsCollector.TransactOpts, _receiver)
}

// DepositFor is a paid mutator transaction binding the contract method 0xaa67c919.
//
// Solidity: function depositFor(address _receiver) payable returns()
func (_OperatorRewardsCollector *OperatorRewardsCollectorTransactorSession) DepositFor(_receiver common.Address) (*types.Transaction, error) {
	return _OperatorRewardsCollector.Contract.DepositFor(&_OperatorRewardsCollector.TransactOpts, _receiver)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_OperatorRewardsCollector *OperatorRewardsCollectorTransactor) GrantRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _OperatorRewardsCollector.contract.Transact(opts, "grantRole", role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_OperatorRewardsCollector *OperatorRewardsCollectorSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _OperatorRewardsCollector.Contract.GrantRole(&_OperatorRewardsCollector.TransactOpts, role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_OperatorRewardsCollector *OperatorRewardsCollectorTransactorSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _OperatorRewardsCollector.Contract.GrantRole(&_OperatorRewardsCollector.TransactOpts, role, account)
}

// Initialize is a paid mutator transaction binding the contract method 0x485cc955.
//
// Solidity: function initialize(address _admin, address _staderConfig) returns()
func (_OperatorRewardsCollector *OperatorRewardsCollectorTransactor) Initialize(opts *bind.TransactOpts, _admin common.Address, _staderConfig common.Address) (*types.Transaction, error) {
	return _OperatorRewardsCollector.contract.Transact(opts, "initialize", _admin, _staderConfig)
}

// Initialize is a paid mutator transaction binding the contract method 0x485cc955.
//
// Solidity: function initialize(address _admin, address _staderConfig) returns()
func (_OperatorRewardsCollector *OperatorRewardsCollectorSession) Initialize(_admin common.Address, _staderConfig common.Address) (*types.Transaction, error) {
	return _OperatorRewardsCollector.Contract.Initialize(&_OperatorRewardsCollector.TransactOpts, _admin, _staderConfig)
}

// Initialize is a paid mutator transaction binding the contract method 0x485cc955.
//
// Solidity: function initialize(address _admin, address _staderConfig) returns()
func (_OperatorRewardsCollector *OperatorRewardsCollectorTransactorSession) Initialize(_admin common.Address, _staderConfig common.Address) (*types.Transaction, error) {
	return _OperatorRewardsCollector.Contract.Initialize(&_OperatorRewardsCollector.TransactOpts, _admin, _staderConfig)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address account) returns()
func (_OperatorRewardsCollector *OperatorRewardsCollectorTransactor) RenounceRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _OperatorRewardsCollector.contract.Transact(opts, "renounceRole", role, account)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address account) returns()
func (_OperatorRewardsCollector *OperatorRewardsCollectorSession) RenounceRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _OperatorRewardsCollector.Contract.RenounceRole(&_OperatorRewardsCollector.TransactOpts, role, account)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address account) returns()
func (_OperatorRewardsCollector *OperatorRewardsCollectorTransactorSession) RenounceRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _OperatorRewardsCollector.Contract.RenounceRole(&_OperatorRewardsCollector.TransactOpts, role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_OperatorRewardsCollector *OperatorRewardsCollectorTransactor) RevokeRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _OperatorRewardsCollector.contract.Transact(opts, "revokeRole", role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_OperatorRewardsCollector *OperatorRewardsCollectorSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _OperatorRewardsCollector.Contract.RevokeRole(&_OperatorRewardsCollector.TransactOpts, role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_OperatorRewardsCollector *OperatorRewardsCollectorTransactorSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _OperatorRewardsCollector.Contract.RevokeRole(&_OperatorRewardsCollector.TransactOpts, role, account)
}

// OperatorRewardsCollectorClaimedIterator is returned from FilterClaimed and is used to iterate over the raw logs and unpacked data for Claimed events raised by the OperatorRewardsCollector contract.
type OperatorRewardsCollectorClaimedIterator struct {
	Event *OperatorRewardsCollectorClaimed // Event containing the contract specifics and raw log

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
func (it *OperatorRewardsCollectorClaimedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OperatorRewardsCollectorClaimed)
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
		it.Event = new(OperatorRewardsCollectorClaimed)
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
func (it *OperatorRewardsCollectorClaimedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OperatorRewardsCollectorClaimedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OperatorRewardsCollectorClaimed represents a Claimed event raised by the OperatorRewardsCollector contract.
type OperatorRewardsCollectorClaimed struct {
	Receiver common.Address
	Amount   *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterClaimed is a free log retrieval operation binding the contract event 0xd8138f8a3f377c5259ca548e70e4c2de94f129f5a11036a15b69513cba2b426a.
//
// Solidity: event Claimed(address indexed receiver, uint256 amount)
func (_OperatorRewardsCollector *OperatorRewardsCollectorFilterer) FilterClaimed(opts *bind.FilterOpts, receiver []common.Address) (*OperatorRewardsCollectorClaimedIterator, error) {

	var receiverRule []interface{}
	for _, receiverItem := range receiver {
		receiverRule = append(receiverRule, receiverItem)
	}

	logs, sub, err := _OperatorRewardsCollector.contract.FilterLogs(opts, "Claimed", receiverRule)
	if err != nil {
		return nil, err
	}
	return &OperatorRewardsCollectorClaimedIterator{contract: _OperatorRewardsCollector.contract, event: "Claimed", logs: logs, sub: sub}, nil
}

// WatchClaimed is a free log subscription operation binding the contract event 0xd8138f8a3f377c5259ca548e70e4c2de94f129f5a11036a15b69513cba2b426a.
//
// Solidity: event Claimed(address indexed receiver, uint256 amount)
func (_OperatorRewardsCollector *OperatorRewardsCollectorFilterer) WatchClaimed(opts *bind.WatchOpts, sink chan<- *OperatorRewardsCollectorClaimed, receiver []common.Address) (event.Subscription, error) {

	var receiverRule []interface{}
	for _, receiverItem := range receiver {
		receiverRule = append(receiverRule, receiverItem)
	}

	logs, sub, err := _OperatorRewardsCollector.contract.WatchLogs(opts, "Claimed", receiverRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OperatorRewardsCollectorClaimed)
				if err := _OperatorRewardsCollector.contract.UnpackLog(event, "Claimed", log); err != nil {
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

// ParseClaimed is a log parse operation binding the contract event 0xd8138f8a3f377c5259ca548e70e4c2de94f129f5a11036a15b69513cba2b426a.
//
// Solidity: event Claimed(address indexed receiver, uint256 amount)
func (_OperatorRewardsCollector *OperatorRewardsCollectorFilterer) ParseClaimed(log types.Log) (*OperatorRewardsCollectorClaimed, error) {
	event := new(OperatorRewardsCollectorClaimed)
	if err := _OperatorRewardsCollector.contract.UnpackLog(event, "Claimed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// OperatorRewardsCollectorDepositedForIterator is returned from FilterDepositedFor and is used to iterate over the raw logs and unpacked data for DepositedFor events raised by the OperatorRewardsCollector contract.
type OperatorRewardsCollectorDepositedForIterator struct {
	Event *OperatorRewardsCollectorDepositedFor // Event containing the contract specifics and raw log

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
func (it *OperatorRewardsCollectorDepositedForIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OperatorRewardsCollectorDepositedFor)
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
		it.Event = new(OperatorRewardsCollectorDepositedFor)
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
func (it *OperatorRewardsCollectorDepositedForIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OperatorRewardsCollectorDepositedForIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OperatorRewardsCollectorDepositedFor represents a DepositedFor event raised by the OperatorRewardsCollector contract.
type OperatorRewardsCollectorDepositedFor struct {
	Sender   common.Address
	Receiver common.Address
	Amount   *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterDepositedFor is a free log retrieval operation binding the contract event 0x11fa725a956e1222d809b94ec211abe46e4218803a3c67d50f6fd9e46ba20a0e.
//
// Solidity: event DepositedFor(address indexed sender, address indexed receiver, uint256 amount)
func (_OperatorRewardsCollector *OperatorRewardsCollectorFilterer) FilterDepositedFor(opts *bind.FilterOpts, sender []common.Address, receiver []common.Address) (*OperatorRewardsCollectorDepositedForIterator, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}
	var receiverRule []interface{}
	for _, receiverItem := range receiver {
		receiverRule = append(receiverRule, receiverItem)
	}

	logs, sub, err := _OperatorRewardsCollector.contract.FilterLogs(opts, "DepositedFor", senderRule, receiverRule)
	if err != nil {
		return nil, err
	}
	return &OperatorRewardsCollectorDepositedForIterator{contract: _OperatorRewardsCollector.contract, event: "DepositedFor", logs: logs, sub: sub}, nil
}

// WatchDepositedFor is a free log subscription operation binding the contract event 0x11fa725a956e1222d809b94ec211abe46e4218803a3c67d50f6fd9e46ba20a0e.
//
// Solidity: event DepositedFor(address indexed sender, address indexed receiver, uint256 amount)
func (_OperatorRewardsCollector *OperatorRewardsCollectorFilterer) WatchDepositedFor(opts *bind.WatchOpts, sink chan<- *OperatorRewardsCollectorDepositedFor, sender []common.Address, receiver []common.Address) (event.Subscription, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}
	var receiverRule []interface{}
	for _, receiverItem := range receiver {
		receiverRule = append(receiverRule, receiverItem)
	}

	logs, sub, err := _OperatorRewardsCollector.contract.WatchLogs(opts, "DepositedFor", senderRule, receiverRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OperatorRewardsCollectorDepositedFor)
				if err := _OperatorRewardsCollector.contract.UnpackLog(event, "DepositedFor", log); err != nil {
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

// ParseDepositedFor is a log parse operation binding the contract event 0x11fa725a956e1222d809b94ec211abe46e4218803a3c67d50f6fd9e46ba20a0e.
//
// Solidity: event DepositedFor(address indexed sender, address indexed receiver, uint256 amount)
func (_OperatorRewardsCollector *OperatorRewardsCollectorFilterer) ParseDepositedFor(log types.Log) (*OperatorRewardsCollectorDepositedFor, error) {
	event := new(OperatorRewardsCollectorDepositedFor)
	if err := _OperatorRewardsCollector.contract.UnpackLog(event, "DepositedFor", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// OperatorRewardsCollectorInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the OperatorRewardsCollector contract.
type OperatorRewardsCollectorInitializedIterator struct {
	Event *OperatorRewardsCollectorInitialized // Event containing the contract specifics and raw log

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
func (it *OperatorRewardsCollectorInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OperatorRewardsCollectorInitialized)
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
		it.Event = new(OperatorRewardsCollectorInitialized)
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
func (it *OperatorRewardsCollectorInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OperatorRewardsCollectorInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OperatorRewardsCollectorInitialized represents a Initialized event raised by the OperatorRewardsCollector contract.
type OperatorRewardsCollectorInitialized struct {
	Version uint8
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_OperatorRewardsCollector *OperatorRewardsCollectorFilterer) FilterInitialized(opts *bind.FilterOpts) (*OperatorRewardsCollectorInitializedIterator, error) {

	logs, sub, err := _OperatorRewardsCollector.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &OperatorRewardsCollectorInitializedIterator{contract: _OperatorRewardsCollector.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_OperatorRewardsCollector *OperatorRewardsCollectorFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *OperatorRewardsCollectorInitialized) (event.Subscription, error) {

	logs, sub, err := _OperatorRewardsCollector.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OperatorRewardsCollectorInitialized)
				if err := _OperatorRewardsCollector.contract.UnpackLog(event, "Initialized", log); err != nil {
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

// ParseInitialized is a log parse operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_OperatorRewardsCollector *OperatorRewardsCollectorFilterer) ParseInitialized(log types.Log) (*OperatorRewardsCollectorInitialized, error) {
	event := new(OperatorRewardsCollectorInitialized)
	if err := _OperatorRewardsCollector.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// OperatorRewardsCollectorPausedIterator is returned from FilterPaused and is used to iterate over the raw logs and unpacked data for Paused events raised by the OperatorRewardsCollector contract.
type OperatorRewardsCollectorPausedIterator struct {
	Event *OperatorRewardsCollectorPaused // Event containing the contract specifics and raw log

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
func (it *OperatorRewardsCollectorPausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OperatorRewardsCollectorPaused)
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
		it.Event = new(OperatorRewardsCollectorPaused)
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
func (it *OperatorRewardsCollectorPausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OperatorRewardsCollectorPausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OperatorRewardsCollectorPaused represents a Paused event raised by the OperatorRewardsCollector contract.
type OperatorRewardsCollectorPaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterPaused is a free log retrieval operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_OperatorRewardsCollector *OperatorRewardsCollectorFilterer) FilterPaused(opts *bind.FilterOpts) (*OperatorRewardsCollectorPausedIterator, error) {

	logs, sub, err := _OperatorRewardsCollector.contract.FilterLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return &OperatorRewardsCollectorPausedIterator{contract: _OperatorRewardsCollector.contract, event: "Paused", logs: logs, sub: sub}, nil
}

// WatchPaused is a free log subscription operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_OperatorRewardsCollector *OperatorRewardsCollectorFilterer) WatchPaused(opts *bind.WatchOpts, sink chan<- *OperatorRewardsCollectorPaused) (event.Subscription, error) {

	logs, sub, err := _OperatorRewardsCollector.contract.WatchLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OperatorRewardsCollectorPaused)
				if err := _OperatorRewardsCollector.contract.UnpackLog(event, "Paused", log); err != nil {
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

// ParsePaused is a log parse operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_OperatorRewardsCollector *OperatorRewardsCollectorFilterer) ParsePaused(log types.Log) (*OperatorRewardsCollectorPaused, error) {
	event := new(OperatorRewardsCollectorPaused)
	if err := _OperatorRewardsCollector.contract.UnpackLog(event, "Paused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// OperatorRewardsCollectorRoleAdminChangedIterator is returned from FilterRoleAdminChanged and is used to iterate over the raw logs and unpacked data for RoleAdminChanged events raised by the OperatorRewardsCollector contract.
type OperatorRewardsCollectorRoleAdminChangedIterator struct {
	Event *OperatorRewardsCollectorRoleAdminChanged // Event containing the contract specifics and raw log

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
func (it *OperatorRewardsCollectorRoleAdminChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OperatorRewardsCollectorRoleAdminChanged)
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
		it.Event = new(OperatorRewardsCollectorRoleAdminChanged)
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
func (it *OperatorRewardsCollectorRoleAdminChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OperatorRewardsCollectorRoleAdminChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OperatorRewardsCollectorRoleAdminChanged represents a RoleAdminChanged event raised by the OperatorRewardsCollector contract.
type OperatorRewardsCollectorRoleAdminChanged struct {
	Role              [32]byte
	PreviousAdminRole [32]byte
	NewAdminRole      [32]byte
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterRoleAdminChanged is a free log retrieval operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_OperatorRewardsCollector *OperatorRewardsCollectorFilterer) FilterRoleAdminChanged(opts *bind.FilterOpts, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (*OperatorRewardsCollectorRoleAdminChangedIterator, error) {

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

	logs, sub, err := _OperatorRewardsCollector.contract.FilterLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return &OperatorRewardsCollectorRoleAdminChangedIterator{contract: _OperatorRewardsCollector.contract, event: "RoleAdminChanged", logs: logs, sub: sub}, nil
}

// WatchRoleAdminChanged is a free log subscription operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_OperatorRewardsCollector *OperatorRewardsCollectorFilterer) WatchRoleAdminChanged(opts *bind.WatchOpts, sink chan<- *OperatorRewardsCollectorRoleAdminChanged, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (event.Subscription, error) {

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

	logs, sub, err := _OperatorRewardsCollector.contract.WatchLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OperatorRewardsCollectorRoleAdminChanged)
				if err := _OperatorRewardsCollector.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
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
func (_OperatorRewardsCollector *OperatorRewardsCollectorFilterer) ParseRoleAdminChanged(log types.Log) (*OperatorRewardsCollectorRoleAdminChanged, error) {
	event := new(OperatorRewardsCollectorRoleAdminChanged)
	if err := _OperatorRewardsCollector.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// OperatorRewardsCollectorRoleGrantedIterator is returned from FilterRoleGranted and is used to iterate over the raw logs and unpacked data for RoleGranted events raised by the OperatorRewardsCollector contract.
type OperatorRewardsCollectorRoleGrantedIterator struct {
	Event *OperatorRewardsCollectorRoleGranted // Event containing the contract specifics and raw log

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
func (it *OperatorRewardsCollectorRoleGrantedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OperatorRewardsCollectorRoleGranted)
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
		it.Event = new(OperatorRewardsCollectorRoleGranted)
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
func (it *OperatorRewardsCollectorRoleGrantedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OperatorRewardsCollectorRoleGrantedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OperatorRewardsCollectorRoleGranted represents a RoleGranted event raised by the OperatorRewardsCollector contract.
type OperatorRewardsCollectorRoleGranted struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleGranted is a free log retrieval operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_OperatorRewardsCollector *OperatorRewardsCollectorFilterer) FilterRoleGranted(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*OperatorRewardsCollectorRoleGrantedIterator, error) {

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

	logs, sub, err := _OperatorRewardsCollector.contract.FilterLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &OperatorRewardsCollectorRoleGrantedIterator{contract: _OperatorRewardsCollector.contract, event: "RoleGranted", logs: logs, sub: sub}, nil
}

// WatchRoleGranted is a free log subscription operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_OperatorRewardsCollector *OperatorRewardsCollectorFilterer) WatchRoleGranted(opts *bind.WatchOpts, sink chan<- *OperatorRewardsCollectorRoleGranted, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _OperatorRewardsCollector.contract.WatchLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OperatorRewardsCollectorRoleGranted)
				if err := _OperatorRewardsCollector.contract.UnpackLog(event, "RoleGranted", log); err != nil {
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
func (_OperatorRewardsCollector *OperatorRewardsCollectorFilterer) ParseRoleGranted(log types.Log) (*OperatorRewardsCollectorRoleGranted, error) {
	event := new(OperatorRewardsCollectorRoleGranted)
	if err := _OperatorRewardsCollector.contract.UnpackLog(event, "RoleGranted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// OperatorRewardsCollectorRoleRevokedIterator is returned from FilterRoleRevoked and is used to iterate over the raw logs and unpacked data for RoleRevoked events raised by the OperatorRewardsCollector contract.
type OperatorRewardsCollectorRoleRevokedIterator struct {
	Event *OperatorRewardsCollectorRoleRevoked // Event containing the contract specifics and raw log

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
func (it *OperatorRewardsCollectorRoleRevokedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OperatorRewardsCollectorRoleRevoked)
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
		it.Event = new(OperatorRewardsCollectorRoleRevoked)
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
func (it *OperatorRewardsCollectorRoleRevokedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OperatorRewardsCollectorRoleRevokedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OperatorRewardsCollectorRoleRevoked represents a RoleRevoked event raised by the OperatorRewardsCollector contract.
type OperatorRewardsCollectorRoleRevoked struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleRevoked is a free log retrieval operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_OperatorRewardsCollector *OperatorRewardsCollectorFilterer) FilterRoleRevoked(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*OperatorRewardsCollectorRoleRevokedIterator, error) {

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

	logs, sub, err := _OperatorRewardsCollector.contract.FilterLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &OperatorRewardsCollectorRoleRevokedIterator{contract: _OperatorRewardsCollector.contract, event: "RoleRevoked", logs: logs, sub: sub}, nil
}

// WatchRoleRevoked is a free log subscription operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_OperatorRewardsCollector *OperatorRewardsCollectorFilterer) WatchRoleRevoked(opts *bind.WatchOpts, sink chan<- *OperatorRewardsCollectorRoleRevoked, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _OperatorRewardsCollector.contract.WatchLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OperatorRewardsCollectorRoleRevoked)
				if err := _OperatorRewardsCollector.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
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
func (_OperatorRewardsCollector *OperatorRewardsCollectorFilterer) ParseRoleRevoked(log types.Log) (*OperatorRewardsCollectorRoleRevoked, error) {
	event := new(OperatorRewardsCollectorRoleRevoked)
	if err := _OperatorRewardsCollector.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// OperatorRewardsCollectorUnpausedIterator is returned from FilterUnpaused and is used to iterate over the raw logs and unpacked data for Unpaused events raised by the OperatorRewardsCollector contract.
type OperatorRewardsCollectorUnpausedIterator struct {
	Event *OperatorRewardsCollectorUnpaused // Event containing the contract specifics and raw log

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
func (it *OperatorRewardsCollectorUnpausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OperatorRewardsCollectorUnpaused)
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
		it.Event = new(OperatorRewardsCollectorUnpaused)
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
func (it *OperatorRewardsCollectorUnpausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OperatorRewardsCollectorUnpausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OperatorRewardsCollectorUnpaused represents a Unpaused event raised by the OperatorRewardsCollector contract.
type OperatorRewardsCollectorUnpaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterUnpaused is a free log retrieval operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_OperatorRewardsCollector *OperatorRewardsCollectorFilterer) FilterUnpaused(opts *bind.FilterOpts) (*OperatorRewardsCollectorUnpausedIterator, error) {

	logs, sub, err := _OperatorRewardsCollector.contract.FilterLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return &OperatorRewardsCollectorUnpausedIterator{contract: _OperatorRewardsCollector.contract, event: "Unpaused", logs: logs, sub: sub}, nil
}

// WatchUnpaused is a free log subscription operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_OperatorRewardsCollector *OperatorRewardsCollectorFilterer) WatchUnpaused(opts *bind.WatchOpts, sink chan<- *OperatorRewardsCollectorUnpaused) (event.Subscription, error) {

	logs, sub, err := _OperatorRewardsCollector.contract.WatchLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OperatorRewardsCollectorUnpaused)
				if err := _OperatorRewardsCollector.contract.UnpackLog(event, "Unpaused", log); err != nil {
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

// ParseUnpaused is a log parse operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_OperatorRewardsCollector *OperatorRewardsCollectorFilterer) ParseUnpaused(log types.Log) (*OperatorRewardsCollectorUnpaused, error) {
	event := new(OperatorRewardsCollectorUnpaused)
	if err := _OperatorRewardsCollector.contract.UnpackLog(event, "Unpaused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// OperatorRewardsCollectorUpdatedStaderConfigIterator is returned from FilterUpdatedStaderConfig and is used to iterate over the raw logs and unpacked data for UpdatedStaderConfig events raised by the OperatorRewardsCollector contract.
type OperatorRewardsCollectorUpdatedStaderConfigIterator struct {
	Event *OperatorRewardsCollectorUpdatedStaderConfig // Event containing the contract specifics and raw log

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
func (it *OperatorRewardsCollectorUpdatedStaderConfigIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OperatorRewardsCollectorUpdatedStaderConfig)
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
		it.Event = new(OperatorRewardsCollectorUpdatedStaderConfig)
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
func (it *OperatorRewardsCollectorUpdatedStaderConfigIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OperatorRewardsCollectorUpdatedStaderConfigIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OperatorRewardsCollectorUpdatedStaderConfig represents a UpdatedStaderConfig event raised by the OperatorRewardsCollector contract.
type OperatorRewardsCollectorUpdatedStaderConfig struct {
	StaderConfig common.Address
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterUpdatedStaderConfig is a free log retrieval operation binding the contract event 0xdb2219043d7b197cb235f1af0cf6d782d77dee3de19e3f4fb6d39aae633b4485.
//
// Solidity: event UpdatedStaderConfig(address indexed staderConfig)
func (_OperatorRewardsCollector *OperatorRewardsCollectorFilterer) FilterUpdatedStaderConfig(opts *bind.FilterOpts, staderConfig []common.Address) (*OperatorRewardsCollectorUpdatedStaderConfigIterator, error) {

	var staderConfigRule []interface{}
	for _, staderConfigItem := range staderConfig {
		staderConfigRule = append(staderConfigRule, staderConfigItem)
	}

	logs, sub, err := _OperatorRewardsCollector.contract.FilterLogs(opts, "UpdatedStaderConfig", staderConfigRule)
	if err != nil {
		return nil, err
	}
	return &OperatorRewardsCollectorUpdatedStaderConfigIterator{contract: _OperatorRewardsCollector.contract, event: "UpdatedStaderConfig", logs: logs, sub: sub}, nil
}

// WatchUpdatedStaderConfig is a free log subscription operation binding the contract event 0xdb2219043d7b197cb235f1af0cf6d782d77dee3de19e3f4fb6d39aae633b4485.
//
// Solidity: event UpdatedStaderConfig(address indexed staderConfig)
func (_OperatorRewardsCollector *OperatorRewardsCollectorFilterer) WatchUpdatedStaderConfig(opts *bind.WatchOpts, sink chan<- *OperatorRewardsCollectorUpdatedStaderConfig, staderConfig []common.Address) (event.Subscription, error) {

	var staderConfigRule []interface{}
	for _, staderConfigItem := range staderConfig {
		staderConfigRule = append(staderConfigRule, staderConfigItem)
	}

	logs, sub, err := _OperatorRewardsCollector.contract.WatchLogs(opts, "UpdatedStaderConfig", staderConfigRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OperatorRewardsCollectorUpdatedStaderConfig)
				if err := _OperatorRewardsCollector.contract.UnpackLog(event, "UpdatedStaderConfig", log); err != nil {
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

// ParseUpdatedStaderConfig is a log parse operation binding the contract event 0xdb2219043d7b197cb235f1af0cf6d782d77dee3de19e3f4fb6d39aae633b4485.
//
// Solidity: event UpdatedStaderConfig(address indexed staderConfig)
func (_OperatorRewardsCollector *OperatorRewardsCollectorFilterer) ParseUpdatedStaderConfig(log types.Log) (*OperatorRewardsCollectorUpdatedStaderConfig, error) {
	event := new(OperatorRewardsCollectorUpdatedStaderConfig)
	if err := _OperatorRewardsCollector.contract.UnpackLog(event, "UpdatedStaderConfig", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
