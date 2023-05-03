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

// NodeElRewardVaultMetaData contains all meta data concerning the NodeElRewardVault contract.
var NodeElRewardVaultMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"ETHTransferFailed\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotEnoughRewardToWithdraw\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ZeroAddress\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"ETHReceived\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"version\",\"type\":\"uint8\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"previousAdminRole\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"newAdminRole\",\"type\":\"bytes32\"}],\"name\":\"RoleAdminChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleGranted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleRevoked\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"staderConfig\",\"type\":\"address\"}],\"name\":\"UpdatedStaderConfig\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"protocolAmount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"operatorAmount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"userAmount\",\"type\":\"uint256\"}],\"name\":\"Withdrawal\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"DEFAULT_ADMIN_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"}],\"name\":\"getRoleAdmin\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"grantRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"hasRole\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"_poolId\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"_operatorId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_staderConfig\",\"type\":\"address\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"operatorId\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"poolId\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"renounceRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"revokeRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"staderConfig\",\"outputs\":[{\"internalType\":\"contractIStaderConfig\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_staderConfig\",\"type\":\"address\"}],\"name\":\"updateStaderConfig\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"withdraw\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"stateMutability\":\"payable\",\"type\":\"receive\"}]",
}

// NodeElRewardVaultABI is the input ABI used to generate the binding from.
// Deprecated: Use NodeElRewardVaultMetaData.ABI instead.
var NodeElRewardVaultABI = NodeElRewardVaultMetaData.ABI

// NodeElRewardVault is an auto generated Go binding around an Ethereum contract.
type NodeElRewardVault struct {
	NodeElRewardVaultCaller     // Read-only binding to the contract
	NodeElRewardVaultTransactor // Write-only binding to the contract
	NodeElRewardVaultFilterer   // Log filterer for contract events
}

// NodeElRewardVaultCaller is an auto generated read-only Go binding around an Ethereum contract.
type NodeElRewardVaultCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// NodeElRewardVaultTransactor is an auto generated write-only Go binding around an Ethereum contract.
type NodeElRewardVaultTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// NodeElRewardVaultFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type NodeElRewardVaultFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// NodeElRewardVaultSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type NodeElRewardVaultSession struct {
	Contract     *NodeElRewardVault // Generic contract binding to set the session for
	CallOpts     bind.CallOpts      // Call options to use throughout this session
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// NodeElRewardVaultCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type NodeElRewardVaultCallerSession struct {
	Contract *NodeElRewardVaultCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts            // Call options to use throughout this session
}

// NodeElRewardVaultTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type NodeElRewardVaultTransactorSession struct {
	Contract     *NodeElRewardVaultTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts            // Transaction auth options to use throughout this session
}

// NodeElRewardVaultRaw is an auto generated low-level Go binding around an Ethereum contract.
type NodeElRewardVaultRaw struct {
	Contract *NodeElRewardVault // Generic contract binding to access the raw methods on
}

// NodeElRewardVaultCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type NodeElRewardVaultCallerRaw struct {
	Contract *NodeElRewardVaultCaller // Generic read-only contract binding to access the raw methods on
}

// NodeElRewardVaultTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type NodeElRewardVaultTransactorRaw struct {
	Contract *NodeElRewardVaultTransactor // Generic write-only contract binding to access the raw methods on
}

// NewNodeElRewardVault creates a new instance of NodeElRewardVault, bound to a specific deployed contract.
func NewNodeElRewardVault(address common.Address, backend bind.ContractBackend) (*NodeElRewardVault, error) {
	contract, err := bindNodeElRewardVault(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &NodeElRewardVault{NodeElRewardVaultCaller: NodeElRewardVaultCaller{contract: contract}, NodeElRewardVaultTransactor: NodeElRewardVaultTransactor{contract: contract}, NodeElRewardVaultFilterer: NodeElRewardVaultFilterer{contract: contract}}, nil
}

// NewNodeElRewardVaultCaller creates a new read-only instance of NodeElRewardVault, bound to a specific deployed contract.
func NewNodeElRewardVaultCaller(address common.Address, caller bind.ContractCaller) (*NodeElRewardVaultCaller, error) {
	contract, err := bindNodeElRewardVault(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &NodeElRewardVaultCaller{contract: contract}, nil
}

// NewNodeElRewardVaultTransactor creates a new write-only instance of NodeElRewardVault, bound to a specific deployed contract.
func NewNodeElRewardVaultTransactor(address common.Address, transactor bind.ContractTransactor) (*NodeElRewardVaultTransactor, error) {
	contract, err := bindNodeElRewardVault(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &NodeElRewardVaultTransactor{contract: contract}, nil
}

// NewNodeElRewardVaultFilterer creates a new log filterer instance of NodeElRewardVault, bound to a specific deployed contract.
func NewNodeElRewardVaultFilterer(address common.Address, filterer bind.ContractFilterer) (*NodeElRewardVaultFilterer, error) {
	contract, err := bindNodeElRewardVault(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &NodeElRewardVaultFilterer{contract: contract}, nil
}

// bindNodeElRewardVault binds a generic wrapper to an already deployed contract.
func bindNodeElRewardVault(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := NodeElRewardVaultMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_NodeElRewardVault *NodeElRewardVaultRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _NodeElRewardVault.Contract.NodeElRewardVaultCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_NodeElRewardVault *NodeElRewardVaultRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _NodeElRewardVault.Contract.NodeElRewardVaultTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_NodeElRewardVault *NodeElRewardVaultRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _NodeElRewardVault.Contract.NodeElRewardVaultTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_NodeElRewardVault *NodeElRewardVaultCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _NodeElRewardVault.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_NodeElRewardVault *NodeElRewardVaultTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _NodeElRewardVault.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_NodeElRewardVault *NodeElRewardVaultTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _NodeElRewardVault.Contract.contract.Transact(opts, method, params...)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_NodeElRewardVault *NodeElRewardVaultCaller) DEFAULTADMINROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _NodeElRewardVault.contract.Call(opts, &out, "DEFAULT_ADMIN_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_NodeElRewardVault *NodeElRewardVaultSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _NodeElRewardVault.Contract.DEFAULTADMINROLE(&_NodeElRewardVault.CallOpts)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_NodeElRewardVault *NodeElRewardVaultCallerSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _NodeElRewardVault.Contract.DEFAULTADMINROLE(&_NodeElRewardVault.CallOpts)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_NodeElRewardVault *NodeElRewardVaultCaller) GetRoleAdmin(opts *bind.CallOpts, role [32]byte) ([32]byte, error) {
	var out []interface{}
	err := _NodeElRewardVault.contract.Call(opts, &out, "getRoleAdmin", role)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_NodeElRewardVault *NodeElRewardVaultSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _NodeElRewardVault.Contract.GetRoleAdmin(&_NodeElRewardVault.CallOpts, role)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_NodeElRewardVault *NodeElRewardVaultCallerSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _NodeElRewardVault.Contract.GetRoleAdmin(&_NodeElRewardVault.CallOpts, role)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_NodeElRewardVault *NodeElRewardVaultCaller) HasRole(opts *bind.CallOpts, role [32]byte, account common.Address) (bool, error) {
	var out []interface{}
	err := _NodeElRewardVault.contract.Call(opts, &out, "hasRole", role, account)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_NodeElRewardVault *NodeElRewardVaultSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _NodeElRewardVault.Contract.HasRole(&_NodeElRewardVault.CallOpts, role, account)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_NodeElRewardVault *NodeElRewardVaultCallerSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _NodeElRewardVault.Contract.HasRole(&_NodeElRewardVault.CallOpts, role, account)
}

// OperatorId is a free data retrieval call binding the contract method 0xbf68b816.
//
// Solidity: function operatorId() view returns(uint256)
func (_NodeElRewardVault *NodeElRewardVaultCaller) OperatorId(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _NodeElRewardVault.contract.Call(opts, &out, "operatorId")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// OperatorId is a free data retrieval call binding the contract method 0xbf68b816.
//
// Solidity: function operatorId() view returns(uint256)
func (_NodeElRewardVault *NodeElRewardVaultSession) OperatorId() (*big.Int, error) {
	return _NodeElRewardVault.Contract.OperatorId(&_NodeElRewardVault.CallOpts)
}

// OperatorId is a free data retrieval call binding the contract method 0xbf68b816.
//
// Solidity: function operatorId() view returns(uint256)
func (_NodeElRewardVault *NodeElRewardVaultCallerSession) OperatorId() (*big.Int, error) {
	return _NodeElRewardVault.Contract.OperatorId(&_NodeElRewardVault.CallOpts)
}

// PoolId is a free data retrieval call binding the contract method 0x3e0dc34e.
//
// Solidity: function poolId() view returns(uint8)
func (_NodeElRewardVault *NodeElRewardVaultCaller) PoolId(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _NodeElRewardVault.contract.Call(opts, &out, "poolId")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// PoolId is a free data retrieval call binding the contract method 0x3e0dc34e.
//
// Solidity: function poolId() view returns(uint8)
func (_NodeElRewardVault *NodeElRewardVaultSession) PoolId() (uint8, error) {
	return _NodeElRewardVault.Contract.PoolId(&_NodeElRewardVault.CallOpts)
}

// PoolId is a free data retrieval call binding the contract method 0x3e0dc34e.
//
// Solidity: function poolId() view returns(uint8)
func (_NodeElRewardVault *NodeElRewardVaultCallerSession) PoolId() (uint8, error) {
	return _NodeElRewardVault.Contract.PoolId(&_NodeElRewardVault.CallOpts)
}

// StaderConfig is a free data retrieval call binding the contract method 0x490ffa35.
//
// Solidity: function staderConfig() view returns(address)
func (_NodeElRewardVault *NodeElRewardVaultCaller) StaderConfig(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _NodeElRewardVault.contract.Call(opts, &out, "staderConfig")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// StaderConfig is a free data retrieval call binding the contract method 0x490ffa35.
//
// Solidity: function staderConfig() view returns(address)
func (_NodeElRewardVault *NodeElRewardVaultSession) StaderConfig() (common.Address, error) {
	return _NodeElRewardVault.Contract.StaderConfig(&_NodeElRewardVault.CallOpts)
}

// StaderConfig is a free data retrieval call binding the contract method 0x490ffa35.
//
// Solidity: function staderConfig() view returns(address)
func (_NodeElRewardVault *NodeElRewardVaultCallerSession) StaderConfig() (common.Address, error) {
	return _NodeElRewardVault.Contract.StaderConfig(&_NodeElRewardVault.CallOpts)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_NodeElRewardVault *NodeElRewardVaultCaller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _NodeElRewardVault.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_NodeElRewardVault *NodeElRewardVaultSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _NodeElRewardVault.Contract.SupportsInterface(&_NodeElRewardVault.CallOpts, interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_NodeElRewardVault *NodeElRewardVaultCallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _NodeElRewardVault.Contract.SupportsInterface(&_NodeElRewardVault.CallOpts, interfaceId)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_NodeElRewardVault *NodeElRewardVaultTransactor) GrantRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _NodeElRewardVault.contract.Transact(opts, "grantRole", role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_NodeElRewardVault *NodeElRewardVaultSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _NodeElRewardVault.Contract.GrantRole(&_NodeElRewardVault.TransactOpts, role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_NodeElRewardVault *NodeElRewardVaultTransactorSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _NodeElRewardVault.Contract.GrantRole(&_NodeElRewardVault.TransactOpts, role, account)
}

// Initialize is a paid mutator transaction binding the contract method 0x98f7f55e.
//
// Solidity: function initialize(uint8 _poolId, uint256 _operatorId, address _staderConfig) returns()
func (_NodeElRewardVault *NodeElRewardVaultTransactor) Initialize(opts *bind.TransactOpts, _poolId uint8, _operatorId *big.Int, _staderConfig common.Address) (*types.Transaction, error) {
	return _NodeElRewardVault.contract.Transact(opts, "initialize", _poolId, _operatorId, _staderConfig)
}

// Initialize is a paid mutator transaction binding the contract method 0x98f7f55e.
//
// Solidity: function initialize(uint8 _poolId, uint256 _operatorId, address _staderConfig) returns()
func (_NodeElRewardVault *NodeElRewardVaultSession) Initialize(_poolId uint8, _operatorId *big.Int, _staderConfig common.Address) (*types.Transaction, error) {
	return _NodeElRewardVault.Contract.Initialize(&_NodeElRewardVault.TransactOpts, _poolId, _operatorId, _staderConfig)
}

// Initialize is a paid mutator transaction binding the contract method 0x98f7f55e.
//
// Solidity: function initialize(uint8 _poolId, uint256 _operatorId, address _staderConfig) returns()
func (_NodeElRewardVault *NodeElRewardVaultTransactorSession) Initialize(_poolId uint8, _operatorId *big.Int, _staderConfig common.Address) (*types.Transaction, error) {
	return _NodeElRewardVault.Contract.Initialize(&_NodeElRewardVault.TransactOpts, _poolId, _operatorId, _staderConfig)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address account) returns()
func (_NodeElRewardVault *NodeElRewardVaultTransactor) RenounceRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _NodeElRewardVault.contract.Transact(opts, "renounceRole", role, account)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address account) returns()
func (_NodeElRewardVault *NodeElRewardVaultSession) RenounceRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _NodeElRewardVault.Contract.RenounceRole(&_NodeElRewardVault.TransactOpts, role, account)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address account) returns()
func (_NodeElRewardVault *NodeElRewardVaultTransactorSession) RenounceRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _NodeElRewardVault.Contract.RenounceRole(&_NodeElRewardVault.TransactOpts, role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_NodeElRewardVault *NodeElRewardVaultTransactor) RevokeRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _NodeElRewardVault.contract.Transact(opts, "revokeRole", role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_NodeElRewardVault *NodeElRewardVaultSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _NodeElRewardVault.Contract.RevokeRole(&_NodeElRewardVault.TransactOpts, role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_NodeElRewardVault *NodeElRewardVaultTransactorSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _NodeElRewardVault.Contract.RevokeRole(&_NodeElRewardVault.TransactOpts, role, account)
}

// UpdateStaderConfig is a paid mutator transaction binding the contract method 0x9ee804cb.
//
// Solidity: function updateStaderConfig(address _staderConfig) returns()
func (_NodeElRewardVault *NodeElRewardVaultTransactor) UpdateStaderConfig(opts *bind.TransactOpts, _staderConfig common.Address) (*types.Transaction, error) {
	return _NodeElRewardVault.contract.Transact(opts, "updateStaderConfig", _staderConfig)
}

// UpdateStaderConfig is a paid mutator transaction binding the contract method 0x9ee804cb.
//
// Solidity: function updateStaderConfig(address _staderConfig) returns()
func (_NodeElRewardVault *NodeElRewardVaultSession) UpdateStaderConfig(_staderConfig common.Address) (*types.Transaction, error) {
	return _NodeElRewardVault.Contract.UpdateStaderConfig(&_NodeElRewardVault.TransactOpts, _staderConfig)
}

// UpdateStaderConfig is a paid mutator transaction binding the contract method 0x9ee804cb.
//
// Solidity: function updateStaderConfig(address _staderConfig) returns()
func (_NodeElRewardVault *NodeElRewardVaultTransactorSession) UpdateStaderConfig(_staderConfig common.Address) (*types.Transaction, error) {
	return _NodeElRewardVault.Contract.UpdateStaderConfig(&_NodeElRewardVault.TransactOpts, _staderConfig)
}

// Withdraw is a paid mutator transaction binding the contract method 0x3ccfd60b.
//
// Solidity: function withdraw() returns()
func (_NodeElRewardVault *NodeElRewardVaultTransactor) Withdraw(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _NodeElRewardVault.contract.Transact(opts, "withdraw")
}

// Withdraw is a paid mutator transaction binding the contract method 0x3ccfd60b.
//
// Solidity: function withdraw() returns()
func (_NodeElRewardVault *NodeElRewardVaultSession) Withdraw() (*types.Transaction, error) {
	return _NodeElRewardVault.Contract.Withdraw(&_NodeElRewardVault.TransactOpts)
}

// Withdraw is a paid mutator transaction binding the contract method 0x3ccfd60b.
//
// Solidity: function withdraw() returns()
func (_NodeElRewardVault *NodeElRewardVaultTransactorSession) Withdraw() (*types.Transaction, error) {
	return _NodeElRewardVault.Contract.Withdraw(&_NodeElRewardVault.TransactOpts)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_NodeElRewardVault *NodeElRewardVaultTransactor) Receive(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _NodeElRewardVault.contract.RawTransact(opts, nil) // calldata is disallowed for receive function
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_NodeElRewardVault *NodeElRewardVaultSession) Receive() (*types.Transaction, error) {
	return _NodeElRewardVault.Contract.Receive(&_NodeElRewardVault.TransactOpts)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_NodeElRewardVault *NodeElRewardVaultTransactorSession) Receive() (*types.Transaction, error) {
	return _NodeElRewardVault.Contract.Receive(&_NodeElRewardVault.TransactOpts)
}

// NodeElRewardVaultETHReceivedIterator is returned from FilterETHReceived and is used to iterate over the raw logs and unpacked data for ETHReceived events raised by the NodeElRewardVault contract.
type NodeElRewardVaultETHReceivedIterator struct {
	Event *NodeElRewardVaultETHReceived // Event containing the contract specifics and raw log

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
func (it *NodeElRewardVaultETHReceivedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NodeElRewardVaultETHReceived)
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
		it.Event = new(NodeElRewardVaultETHReceived)
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
func (it *NodeElRewardVaultETHReceivedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NodeElRewardVaultETHReceivedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NodeElRewardVaultETHReceived represents a ETHReceived event raised by the NodeElRewardVault contract.
type NodeElRewardVaultETHReceived struct {
	Sender common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterETHReceived is a free log retrieval operation binding the contract event 0xbfe611b001dfcd411432f7bf0d79b82b4b2ee81511edac123a3403c357fb972a.
//
// Solidity: event ETHReceived(address indexed sender, uint256 amount)
func (_NodeElRewardVault *NodeElRewardVaultFilterer) FilterETHReceived(opts *bind.FilterOpts, sender []common.Address) (*NodeElRewardVaultETHReceivedIterator, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _NodeElRewardVault.contract.FilterLogs(opts, "ETHReceived", senderRule)
	if err != nil {
		return nil, err
	}
	return &NodeElRewardVaultETHReceivedIterator{contract: _NodeElRewardVault.contract, event: "ETHReceived", logs: logs, sub: sub}, nil
}

// WatchETHReceived is a free log subscription operation binding the contract event 0xbfe611b001dfcd411432f7bf0d79b82b4b2ee81511edac123a3403c357fb972a.
//
// Solidity: event ETHReceived(address indexed sender, uint256 amount)
func (_NodeElRewardVault *NodeElRewardVaultFilterer) WatchETHReceived(opts *bind.WatchOpts, sink chan<- *NodeElRewardVaultETHReceived, sender []common.Address) (event.Subscription, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _NodeElRewardVault.contract.WatchLogs(opts, "ETHReceived", senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NodeElRewardVaultETHReceived)
				if err := _NodeElRewardVault.contract.UnpackLog(event, "ETHReceived", log); err != nil {
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

// ParseETHReceived is a log parse operation binding the contract event 0xbfe611b001dfcd411432f7bf0d79b82b4b2ee81511edac123a3403c357fb972a.
//
// Solidity: event ETHReceived(address indexed sender, uint256 amount)
func (_NodeElRewardVault *NodeElRewardVaultFilterer) ParseETHReceived(log types.Log) (*NodeElRewardVaultETHReceived, error) {
	event := new(NodeElRewardVaultETHReceived)
	if err := _NodeElRewardVault.contract.UnpackLog(event, "ETHReceived", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// NodeElRewardVaultInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the NodeElRewardVault contract.
type NodeElRewardVaultInitializedIterator struct {
	Event *NodeElRewardVaultInitialized // Event containing the contract specifics and raw log

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
func (it *NodeElRewardVaultInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NodeElRewardVaultInitialized)
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
		it.Event = new(NodeElRewardVaultInitialized)
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
func (it *NodeElRewardVaultInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NodeElRewardVaultInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NodeElRewardVaultInitialized represents a Initialized event raised by the NodeElRewardVault contract.
type NodeElRewardVaultInitialized struct {
	Version uint8
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_NodeElRewardVault *NodeElRewardVaultFilterer) FilterInitialized(opts *bind.FilterOpts) (*NodeElRewardVaultInitializedIterator, error) {

	logs, sub, err := _NodeElRewardVault.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &NodeElRewardVaultInitializedIterator{contract: _NodeElRewardVault.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_NodeElRewardVault *NodeElRewardVaultFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *NodeElRewardVaultInitialized) (event.Subscription, error) {

	logs, sub, err := _NodeElRewardVault.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NodeElRewardVaultInitialized)
				if err := _NodeElRewardVault.contract.UnpackLog(event, "Initialized", log); err != nil {
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
func (_NodeElRewardVault *NodeElRewardVaultFilterer) ParseInitialized(log types.Log) (*NodeElRewardVaultInitialized, error) {
	event := new(NodeElRewardVaultInitialized)
	if err := _NodeElRewardVault.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// NodeElRewardVaultRoleAdminChangedIterator is returned from FilterRoleAdminChanged and is used to iterate over the raw logs and unpacked data for RoleAdminChanged events raised by the NodeElRewardVault contract.
type NodeElRewardVaultRoleAdminChangedIterator struct {
	Event *NodeElRewardVaultRoleAdminChanged // Event containing the contract specifics and raw log

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
func (it *NodeElRewardVaultRoleAdminChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NodeElRewardVaultRoleAdminChanged)
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
		it.Event = new(NodeElRewardVaultRoleAdminChanged)
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
func (it *NodeElRewardVaultRoleAdminChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NodeElRewardVaultRoleAdminChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NodeElRewardVaultRoleAdminChanged represents a RoleAdminChanged event raised by the NodeElRewardVault contract.
type NodeElRewardVaultRoleAdminChanged struct {
	Role              [32]byte
	PreviousAdminRole [32]byte
	NewAdminRole      [32]byte
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterRoleAdminChanged is a free log retrieval operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_NodeElRewardVault *NodeElRewardVaultFilterer) FilterRoleAdminChanged(opts *bind.FilterOpts, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (*NodeElRewardVaultRoleAdminChangedIterator, error) {

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

	logs, sub, err := _NodeElRewardVault.contract.FilterLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return &NodeElRewardVaultRoleAdminChangedIterator{contract: _NodeElRewardVault.contract, event: "RoleAdminChanged", logs: logs, sub: sub}, nil
}

// WatchRoleAdminChanged is a free log subscription operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_NodeElRewardVault *NodeElRewardVaultFilterer) WatchRoleAdminChanged(opts *bind.WatchOpts, sink chan<- *NodeElRewardVaultRoleAdminChanged, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (event.Subscription, error) {

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

	logs, sub, err := _NodeElRewardVault.contract.WatchLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NodeElRewardVaultRoleAdminChanged)
				if err := _NodeElRewardVault.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
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
func (_NodeElRewardVault *NodeElRewardVaultFilterer) ParseRoleAdminChanged(log types.Log) (*NodeElRewardVaultRoleAdminChanged, error) {
	event := new(NodeElRewardVaultRoleAdminChanged)
	if err := _NodeElRewardVault.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// NodeElRewardVaultRoleGrantedIterator is returned from FilterRoleGranted and is used to iterate over the raw logs and unpacked data for RoleGranted events raised by the NodeElRewardVault contract.
type NodeElRewardVaultRoleGrantedIterator struct {
	Event *NodeElRewardVaultRoleGranted // Event containing the contract specifics and raw log

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
func (it *NodeElRewardVaultRoleGrantedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NodeElRewardVaultRoleGranted)
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
		it.Event = new(NodeElRewardVaultRoleGranted)
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
func (it *NodeElRewardVaultRoleGrantedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NodeElRewardVaultRoleGrantedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NodeElRewardVaultRoleGranted represents a RoleGranted event raised by the NodeElRewardVault contract.
type NodeElRewardVaultRoleGranted struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleGranted is a free log retrieval operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_NodeElRewardVault *NodeElRewardVaultFilterer) FilterRoleGranted(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*NodeElRewardVaultRoleGrantedIterator, error) {

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

	logs, sub, err := _NodeElRewardVault.contract.FilterLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &NodeElRewardVaultRoleGrantedIterator{contract: _NodeElRewardVault.contract, event: "RoleGranted", logs: logs, sub: sub}, nil
}

// WatchRoleGranted is a free log subscription operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_NodeElRewardVault *NodeElRewardVaultFilterer) WatchRoleGranted(opts *bind.WatchOpts, sink chan<- *NodeElRewardVaultRoleGranted, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _NodeElRewardVault.contract.WatchLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NodeElRewardVaultRoleGranted)
				if err := _NodeElRewardVault.contract.UnpackLog(event, "RoleGranted", log); err != nil {
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
func (_NodeElRewardVault *NodeElRewardVaultFilterer) ParseRoleGranted(log types.Log) (*NodeElRewardVaultRoleGranted, error) {
	event := new(NodeElRewardVaultRoleGranted)
	if err := _NodeElRewardVault.contract.UnpackLog(event, "RoleGranted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// NodeElRewardVaultRoleRevokedIterator is returned from FilterRoleRevoked and is used to iterate over the raw logs and unpacked data for RoleRevoked events raised by the NodeElRewardVault contract.
type NodeElRewardVaultRoleRevokedIterator struct {
	Event *NodeElRewardVaultRoleRevoked // Event containing the contract specifics and raw log

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
func (it *NodeElRewardVaultRoleRevokedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NodeElRewardVaultRoleRevoked)
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
		it.Event = new(NodeElRewardVaultRoleRevoked)
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
func (it *NodeElRewardVaultRoleRevokedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NodeElRewardVaultRoleRevokedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NodeElRewardVaultRoleRevoked represents a RoleRevoked event raised by the NodeElRewardVault contract.
type NodeElRewardVaultRoleRevoked struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleRevoked is a free log retrieval operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_NodeElRewardVault *NodeElRewardVaultFilterer) FilterRoleRevoked(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*NodeElRewardVaultRoleRevokedIterator, error) {

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

	logs, sub, err := _NodeElRewardVault.contract.FilterLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &NodeElRewardVaultRoleRevokedIterator{contract: _NodeElRewardVault.contract, event: "RoleRevoked", logs: logs, sub: sub}, nil
}

// WatchRoleRevoked is a free log subscription operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_NodeElRewardVault *NodeElRewardVaultFilterer) WatchRoleRevoked(opts *bind.WatchOpts, sink chan<- *NodeElRewardVaultRoleRevoked, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _NodeElRewardVault.contract.WatchLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NodeElRewardVaultRoleRevoked)
				if err := _NodeElRewardVault.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
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
func (_NodeElRewardVault *NodeElRewardVaultFilterer) ParseRoleRevoked(log types.Log) (*NodeElRewardVaultRoleRevoked, error) {
	event := new(NodeElRewardVaultRoleRevoked)
	if err := _NodeElRewardVault.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// NodeElRewardVaultUpdatedStaderConfigIterator is returned from FilterUpdatedStaderConfig and is used to iterate over the raw logs and unpacked data for UpdatedStaderConfig events raised by the NodeElRewardVault contract.
type NodeElRewardVaultUpdatedStaderConfigIterator struct {
	Event *NodeElRewardVaultUpdatedStaderConfig // Event containing the contract specifics and raw log

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
func (it *NodeElRewardVaultUpdatedStaderConfigIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NodeElRewardVaultUpdatedStaderConfig)
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
		it.Event = new(NodeElRewardVaultUpdatedStaderConfig)
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
func (it *NodeElRewardVaultUpdatedStaderConfigIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NodeElRewardVaultUpdatedStaderConfigIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NodeElRewardVaultUpdatedStaderConfig represents a UpdatedStaderConfig event raised by the NodeElRewardVault contract.
type NodeElRewardVaultUpdatedStaderConfig struct {
	StaderConfig common.Address
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterUpdatedStaderConfig is a free log retrieval operation binding the contract event 0xdb2219043d7b197cb235f1af0cf6d782d77dee3de19e3f4fb6d39aae633b4485.
//
// Solidity: event UpdatedStaderConfig(address staderConfig)
func (_NodeElRewardVault *NodeElRewardVaultFilterer) FilterUpdatedStaderConfig(opts *bind.FilterOpts) (*NodeElRewardVaultUpdatedStaderConfigIterator, error) {

	logs, sub, err := _NodeElRewardVault.contract.FilterLogs(opts, "UpdatedStaderConfig")
	if err != nil {
		return nil, err
	}
	return &NodeElRewardVaultUpdatedStaderConfigIterator{contract: _NodeElRewardVault.contract, event: "UpdatedStaderConfig", logs: logs, sub: sub}, nil
}

// WatchUpdatedStaderConfig is a free log subscription operation binding the contract event 0xdb2219043d7b197cb235f1af0cf6d782d77dee3de19e3f4fb6d39aae633b4485.
//
// Solidity: event UpdatedStaderConfig(address staderConfig)
func (_NodeElRewardVault *NodeElRewardVaultFilterer) WatchUpdatedStaderConfig(opts *bind.WatchOpts, sink chan<- *NodeElRewardVaultUpdatedStaderConfig) (event.Subscription, error) {

	logs, sub, err := _NodeElRewardVault.contract.WatchLogs(opts, "UpdatedStaderConfig")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NodeElRewardVaultUpdatedStaderConfig)
				if err := _NodeElRewardVault.contract.UnpackLog(event, "UpdatedStaderConfig", log); err != nil {
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
// Solidity: event UpdatedStaderConfig(address staderConfig)
func (_NodeElRewardVault *NodeElRewardVaultFilterer) ParseUpdatedStaderConfig(log types.Log) (*NodeElRewardVaultUpdatedStaderConfig, error) {
	event := new(NodeElRewardVaultUpdatedStaderConfig)
	if err := _NodeElRewardVault.contract.UnpackLog(event, "UpdatedStaderConfig", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// NodeElRewardVaultWithdrawalIterator is returned from FilterWithdrawal and is used to iterate over the raw logs and unpacked data for Withdrawal events raised by the NodeElRewardVault contract.
type NodeElRewardVaultWithdrawalIterator struct {
	Event *NodeElRewardVaultWithdrawal // Event containing the contract specifics and raw log

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
func (it *NodeElRewardVaultWithdrawalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NodeElRewardVaultWithdrawal)
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
		it.Event = new(NodeElRewardVaultWithdrawal)
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
func (it *NodeElRewardVaultWithdrawalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NodeElRewardVaultWithdrawalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NodeElRewardVaultWithdrawal represents a Withdrawal event raised by the NodeElRewardVault contract.
type NodeElRewardVaultWithdrawal struct {
	ProtocolAmount *big.Int
	OperatorAmount *big.Int
	UserAmount     *big.Int
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterWithdrawal is a free log retrieval operation binding the contract event 0x4de79efceb38f14026173a22fb2113555409a7a88343c2a780064d2ce0a00a87.
//
// Solidity: event Withdrawal(uint256 protocolAmount, uint256 operatorAmount, uint256 userAmount)
func (_NodeElRewardVault *NodeElRewardVaultFilterer) FilterWithdrawal(opts *bind.FilterOpts) (*NodeElRewardVaultWithdrawalIterator, error) {

	logs, sub, err := _NodeElRewardVault.contract.FilterLogs(opts, "Withdrawal")
	if err != nil {
		return nil, err
	}
	return &NodeElRewardVaultWithdrawalIterator{contract: _NodeElRewardVault.contract, event: "Withdrawal", logs: logs, sub: sub}, nil
}

// WatchWithdrawal is a free log subscription operation binding the contract event 0x4de79efceb38f14026173a22fb2113555409a7a88343c2a780064d2ce0a00a87.
//
// Solidity: event Withdrawal(uint256 protocolAmount, uint256 operatorAmount, uint256 userAmount)
func (_NodeElRewardVault *NodeElRewardVaultFilterer) WatchWithdrawal(opts *bind.WatchOpts, sink chan<- *NodeElRewardVaultWithdrawal) (event.Subscription, error) {

	logs, sub, err := _NodeElRewardVault.contract.WatchLogs(opts, "Withdrawal")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NodeElRewardVaultWithdrawal)
				if err := _NodeElRewardVault.contract.UnpackLog(event, "Withdrawal", log); err != nil {
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

// ParseWithdrawal is a log parse operation binding the contract event 0x4de79efceb38f14026173a22fb2113555409a7a88343c2a780064d2ce0a00a87.
//
// Solidity: event Withdrawal(uint256 protocolAmount, uint256 operatorAmount, uint256 userAmount)
func (_NodeElRewardVault *NodeElRewardVaultFilterer) ParseWithdrawal(log types.Log) (*NodeElRewardVaultWithdrawal, error) {
	event := new(NodeElRewardVaultWithdrawal)
	if err := _NodeElRewardVault.contract.UnpackLog(event, "Withdrawal", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
