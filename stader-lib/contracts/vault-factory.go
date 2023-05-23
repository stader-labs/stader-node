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

// VaultFactoryMetaData contains all meta data concerning the VaultFactory contract.
var VaultFactoryMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"ZeroAddress\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"version\",\"type\":\"uint8\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"nodeDistributor\",\"type\":\"address\"}],\"name\":\"NodeELRewardVaultCreated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"previousAdminRole\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"newAdminRole\",\"type\":\"bytes32\"}],\"name\":\"RoleAdminChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleGranted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleRevoked\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"staderConfig\",\"type\":\"address\"}],\"name\":\"UpdatedStaderConfig\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"withdrawVault\",\"type\":\"address\"}],\"name\":\"WithdrawVaultCreated\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"DEFAULT_ADMIN_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"NODE_REGISTRY_CONTRACT\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"_poolId\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"_operatorId\",\"type\":\"uint256\"}],\"name\":\"computeNodeELRewardVaultAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"_poolId\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"_operatorId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_validatorCount\",\"type\":\"uint256\"}],\"name\":\"computeWithdrawVaultAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"_poolId\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"_operatorId\",\"type\":\"uint256\"}],\"name\":\"deployNodeELRewardVault\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"_poolId\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"_operatorId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_validatorCount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_validatorId\",\"type\":\"uint256\"}],\"name\":\"deployWithdrawVault\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"}],\"name\":\"getRoleAdmin\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_withdrawVault\",\"type\":\"address\"}],\"name\":\"getValidatorWithdrawCredential\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"grantRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"hasRole\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_admin\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_staderConfig\",\"type\":\"address\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"nodeELRewardVaultImplementation\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"renounceRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"revokeRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"staderConfig\",\"outputs\":[{\"internalType\":\"contractIStaderConfig\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_staderConfig\",\"type\":\"address\"}],\"name\":\"updateStaderConfig\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"validatorWithdrawalVaultImplementation\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// VaultFactoryABI is the input ABI used to generate the binding from.
// Deprecated: Use VaultFactoryMetaData.ABI instead.
var VaultFactoryABI = VaultFactoryMetaData.ABI

// VaultFactory is an auto generated Go binding around an Ethereum contract.
type VaultFactory struct {
	VaultFactoryCaller     // Read-only binding to the contract
	VaultFactoryTransactor // Write-only binding to the contract
	VaultFactoryFilterer   // Log filterer for contract events
}

// VaultFactoryCaller is an auto generated read-only Go binding around an Ethereum contract.
type VaultFactoryCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// VaultFactoryTransactor is an auto generated write-only Go binding around an Ethereum contract.
type VaultFactoryTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// VaultFactoryFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type VaultFactoryFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// VaultFactorySession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type VaultFactorySession struct {
	Contract     *VaultFactory     // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// VaultFactoryCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type VaultFactoryCallerSession struct {
	Contract *VaultFactoryCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts       // Call options to use throughout this session
}

// VaultFactoryTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type VaultFactoryTransactorSession struct {
	Contract     *VaultFactoryTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts       // Transaction auth options to use throughout this session
}

// VaultFactoryRaw is an auto generated low-level Go binding around an Ethereum contract.
type VaultFactoryRaw struct {
	Contract *VaultFactory // Generic contract binding to access the raw methods on
}

// VaultFactoryCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type VaultFactoryCallerRaw struct {
	Contract *VaultFactoryCaller // Generic read-only contract binding to access the raw methods on
}

// VaultFactoryTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type VaultFactoryTransactorRaw struct {
	Contract *VaultFactoryTransactor // Generic write-only contract binding to access the raw methods on
}

// NewVaultFactory creates a new instance of VaultFactory, bound to a specific deployed contract.
func NewVaultFactory(address common.Address, backend bind.ContractBackend) (*VaultFactory, error) {
	contract, err := bindVaultFactory(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &VaultFactory{VaultFactoryCaller: VaultFactoryCaller{contract: contract}, VaultFactoryTransactor: VaultFactoryTransactor{contract: contract}, VaultFactoryFilterer: VaultFactoryFilterer{contract: contract}}, nil
}

// NewVaultFactoryCaller creates a new read-only instance of VaultFactory, bound to a specific deployed contract.
func NewVaultFactoryCaller(address common.Address, caller bind.ContractCaller) (*VaultFactoryCaller, error) {
	contract, err := bindVaultFactory(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &VaultFactoryCaller{contract: contract}, nil
}

// NewVaultFactoryTransactor creates a new write-only instance of VaultFactory, bound to a specific deployed contract.
func NewVaultFactoryTransactor(address common.Address, transactor bind.ContractTransactor) (*VaultFactoryTransactor, error) {
	contract, err := bindVaultFactory(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &VaultFactoryTransactor{contract: contract}, nil
}

// NewVaultFactoryFilterer creates a new log filterer instance of VaultFactory, bound to a specific deployed contract.
func NewVaultFactoryFilterer(address common.Address, filterer bind.ContractFilterer) (*VaultFactoryFilterer, error) {
	contract, err := bindVaultFactory(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &VaultFactoryFilterer{contract: contract}, nil
}

// bindVaultFactory binds a generic wrapper to an already deployed contract.
func bindVaultFactory(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := VaultFactoryMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_VaultFactory *VaultFactoryRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _VaultFactory.Contract.VaultFactoryCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_VaultFactory *VaultFactoryRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _VaultFactory.Contract.VaultFactoryTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_VaultFactory *VaultFactoryRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _VaultFactory.Contract.VaultFactoryTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_VaultFactory *VaultFactoryCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _VaultFactory.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_VaultFactory *VaultFactoryTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _VaultFactory.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_VaultFactory *VaultFactoryTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _VaultFactory.Contract.contract.Transact(opts, method, params...)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_VaultFactory *VaultFactoryCaller) DEFAULTADMINROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _VaultFactory.contract.Call(opts, &out, "DEFAULT_ADMIN_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_VaultFactory *VaultFactorySession) DEFAULTADMINROLE() ([32]byte, error) {
	return _VaultFactory.Contract.DEFAULTADMINROLE(&_VaultFactory.CallOpts)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_VaultFactory *VaultFactoryCallerSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _VaultFactory.Contract.DEFAULTADMINROLE(&_VaultFactory.CallOpts)
}

// NODEREGISTRYCONTRACT is a free data retrieval call binding the contract method 0xf8bc93a5.
//
// Solidity: function NODE_REGISTRY_CONTRACT() view returns(bytes32)
func (_VaultFactory *VaultFactoryCaller) NODEREGISTRYCONTRACT(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _VaultFactory.contract.Call(opts, &out, "NODE_REGISTRY_CONTRACT")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// NODEREGISTRYCONTRACT is a free data retrieval call binding the contract method 0xf8bc93a5.
//
// Solidity: function NODE_REGISTRY_CONTRACT() view returns(bytes32)
func (_VaultFactory *VaultFactorySession) NODEREGISTRYCONTRACT() ([32]byte, error) {
	return _VaultFactory.Contract.NODEREGISTRYCONTRACT(&_VaultFactory.CallOpts)
}

// NODEREGISTRYCONTRACT is a free data retrieval call binding the contract method 0xf8bc93a5.
//
// Solidity: function NODE_REGISTRY_CONTRACT() view returns(bytes32)
func (_VaultFactory *VaultFactoryCallerSession) NODEREGISTRYCONTRACT() ([32]byte, error) {
	return _VaultFactory.Contract.NODEREGISTRYCONTRACT(&_VaultFactory.CallOpts)
}

// ComputeNodeELRewardVaultAddress is a free data retrieval call binding the contract method 0xca934f60.
//
// Solidity: function computeNodeELRewardVaultAddress(uint8 _poolId, uint256 _operatorId) view returns(address)
func (_VaultFactory *VaultFactoryCaller) ComputeNodeELRewardVaultAddress(opts *bind.CallOpts, _poolId uint8, _operatorId *big.Int) (common.Address, error) {
	var out []interface{}
	err := _VaultFactory.contract.Call(opts, &out, "computeNodeELRewardVaultAddress", _poolId, _operatorId)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// ComputeNodeELRewardVaultAddress is a free data retrieval call binding the contract method 0xca934f60.
//
// Solidity: function computeNodeELRewardVaultAddress(uint8 _poolId, uint256 _operatorId) view returns(address)
func (_VaultFactory *VaultFactorySession) ComputeNodeELRewardVaultAddress(_poolId uint8, _operatorId *big.Int) (common.Address, error) {
	return _VaultFactory.Contract.ComputeNodeELRewardVaultAddress(&_VaultFactory.CallOpts, _poolId, _operatorId)
}

// ComputeNodeELRewardVaultAddress is a free data retrieval call binding the contract method 0xca934f60.
//
// Solidity: function computeNodeELRewardVaultAddress(uint8 _poolId, uint256 _operatorId) view returns(address)
func (_VaultFactory *VaultFactoryCallerSession) ComputeNodeELRewardVaultAddress(_poolId uint8, _operatorId *big.Int) (common.Address, error) {
	return _VaultFactory.Contract.ComputeNodeELRewardVaultAddress(&_VaultFactory.CallOpts, _poolId, _operatorId)
}

// ComputeWithdrawVaultAddress is a free data retrieval call binding the contract method 0x74903b02.
//
// Solidity: function computeWithdrawVaultAddress(uint8 _poolId, uint256 _operatorId, uint256 _validatorCount) view returns(address)
func (_VaultFactory *VaultFactoryCaller) ComputeWithdrawVaultAddress(opts *bind.CallOpts, _poolId uint8, _operatorId *big.Int, _validatorCount *big.Int) (common.Address, error) {
	var out []interface{}
	err := _VaultFactory.contract.Call(opts, &out, "computeWithdrawVaultAddress", _poolId, _operatorId, _validatorCount)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// ComputeWithdrawVaultAddress is a free data retrieval call binding the contract method 0x74903b02.
//
// Solidity: function computeWithdrawVaultAddress(uint8 _poolId, uint256 _operatorId, uint256 _validatorCount) view returns(address)
func (_VaultFactory *VaultFactorySession) ComputeWithdrawVaultAddress(_poolId uint8, _operatorId *big.Int, _validatorCount *big.Int) (common.Address, error) {
	return _VaultFactory.Contract.ComputeWithdrawVaultAddress(&_VaultFactory.CallOpts, _poolId, _operatorId, _validatorCount)
}

// ComputeWithdrawVaultAddress is a free data retrieval call binding the contract method 0x74903b02.
//
// Solidity: function computeWithdrawVaultAddress(uint8 _poolId, uint256 _operatorId, uint256 _validatorCount) view returns(address)
func (_VaultFactory *VaultFactoryCallerSession) ComputeWithdrawVaultAddress(_poolId uint8, _operatorId *big.Int, _validatorCount *big.Int) (common.Address, error) {
	return _VaultFactory.Contract.ComputeWithdrawVaultAddress(&_VaultFactory.CallOpts, _poolId, _operatorId, _validatorCount)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_VaultFactory *VaultFactoryCaller) GetRoleAdmin(opts *bind.CallOpts, role [32]byte) ([32]byte, error) {
	var out []interface{}
	err := _VaultFactory.contract.Call(opts, &out, "getRoleAdmin", role)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_VaultFactory *VaultFactorySession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _VaultFactory.Contract.GetRoleAdmin(&_VaultFactory.CallOpts, role)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_VaultFactory *VaultFactoryCallerSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _VaultFactory.Contract.GetRoleAdmin(&_VaultFactory.CallOpts, role)
}

// GetValidatorWithdrawCredential is a free data retrieval call binding the contract method 0xae4e4e45.
//
// Solidity: function getValidatorWithdrawCredential(address _withdrawVault) pure returns(bytes)
func (_VaultFactory *VaultFactoryCaller) GetValidatorWithdrawCredential(opts *bind.CallOpts, _withdrawVault common.Address) ([]byte, error) {
	var out []interface{}
	err := _VaultFactory.contract.Call(opts, &out, "getValidatorWithdrawCredential", _withdrawVault)

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// GetValidatorWithdrawCredential is a free data retrieval call binding the contract method 0xae4e4e45.
//
// Solidity: function getValidatorWithdrawCredential(address _withdrawVault) pure returns(bytes)
func (_VaultFactory *VaultFactorySession) GetValidatorWithdrawCredential(_withdrawVault common.Address) ([]byte, error) {
	return _VaultFactory.Contract.GetValidatorWithdrawCredential(&_VaultFactory.CallOpts, _withdrawVault)
}

// GetValidatorWithdrawCredential is a free data retrieval call binding the contract method 0xae4e4e45.
//
// Solidity: function getValidatorWithdrawCredential(address _withdrawVault) pure returns(bytes)
func (_VaultFactory *VaultFactoryCallerSession) GetValidatorWithdrawCredential(_withdrawVault common.Address) ([]byte, error) {
	return _VaultFactory.Contract.GetValidatorWithdrawCredential(&_VaultFactory.CallOpts, _withdrawVault)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_VaultFactory *VaultFactoryCaller) HasRole(opts *bind.CallOpts, role [32]byte, account common.Address) (bool, error) {
	var out []interface{}
	err := _VaultFactory.contract.Call(opts, &out, "hasRole", role, account)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_VaultFactory *VaultFactorySession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _VaultFactory.Contract.HasRole(&_VaultFactory.CallOpts, role, account)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_VaultFactory *VaultFactoryCallerSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _VaultFactory.Contract.HasRole(&_VaultFactory.CallOpts, role, account)
}

// NodeELRewardVaultImplementation is a free data retrieval call binding the contract method 0x6a9f22ba.
//
// Solidity: function nodeELRewardVaultImplementation() view returns(address)
func (_VaultFactory *VaultFactoryCaller) NodeELRewardVaultImplementation(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _VaultFactory.contract.Call(opts, &out, "nodeELRewardVaultImplementation")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// NodeELRewardVaultImplementation is a free data retrieval call binding the contract method 0x6a9f22ba.
//
// Solidity: function nodeELRewardVaultImplementation() view returns(address)
func (_VaultFactory *VaultFactorySession) NodeELRewardVaultImplementation() (common.Address, error) {
	return _VaultFactory.Contract.NodeELRewardVaultImplementation(&_VaultFactory.CallOpts)
}

// NodeELRewardVaultImplementation is a free data retrieval call binding the contract method 0x6a9f22ba.
//
// Solidity: function nodeELRewardVaultImplementation() view returns(address)
func (_VaultFactory *VaultFactoryCallerSession) NodeELRewardVaultImplementation() (common.Address, error) {
	return _VaultFactory.Contract.NodeELRewardVaultImplementation(&_VaultFactory.CallOpts)
}

// StaderConfig is a free data retrieval call binding the contract method 0x490ffa35.
//
// Solidity: function staderConfig() view returns(address)
func (_VaultFactory *VaultFactoryCaller) StaderConfig(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _VaultFactory.contract.Call(opts, &out, "staderConfig")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// StaderConfig is a free data retrieval call binding the contract method 0x490ffa35.
//
// Solidity: function staderConfig() view returns(address)
func (_VaultFactory *VaultFactorySession) StaderConfig() (common.Address, error) {
	return _VaultFactory.Contract.StaderConfig(&_VaultFactory.CallOpts)
}

// StaderConfig is a free data retrieval call binding the contract method 0x490ffa35.
//
// Solidity: function staderConfig() view returns(address)
func (_VaultFactory *VaultFactoryCallerSession) StaderConfig() (common.Address, error) {
	return _VaultFactory.Contract.StaderConfig(&_VaultFactory.CallOpts)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_VaultFactory *VaultFactoryCaller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _VaultFactory.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_VaultFactory *VaultFactorySession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _VaultFactory.Contract.SupportsInterface(&_VaultFactory.CallOpts, interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_VaultFactory *VaultFactoryCallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _VaultFactory.Contract.SupportsInterface(&_VaultFactory.CallOpts, interfaceId)
}

// ValidatorWithdrawalVaultImplementation is a free data retrieval call binding the contract method 0x436c55e6.
//
// Solidity: function validatorWithdrawalVaultImplementation() view returns(address)
func (_VaultFactory *VaultFactoryCaller) ValidatorWithdrawalVaultImplementation(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _VaultFactory.contract.Call(opts, &out, "validatorWithdrawalVaultImplementation")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// ValidatorWithdrawalVaultImplementation is a free data retrieval call binding the contract method 0x436c55e6.
//
// Solidity: function validatorWithdrawalVaultImplementation() view returns(address)
func (_VaultFactory *VaultFactorySession) ValidatorWithdrawalVaultImplementation() (common.Address, error) {
	return _VaultFactory.Contract.ValidatorWithdrawalVaultImplementation(&_VaultFactory.CallOpts)
}

// ValidatorWithdrawalVaultImplementation is a free data retrieval call binding the contract method 0x436c55e6.
//
// Solidity: function validatorWithdrawalVaultImplementation() view returns(address)
func (_VaultFactory *VaultFactoryCallerSession) ValidatorWithdrawalVaultImplementation() (common.Address, error) {
	return _VaultFactory.Contract.ValidatorWithdrawalVaultImplementation(&_VaultFactory.CallOpts)
}

// DeployNodeELRewardVault is a paid mutator transaction binding the contract method 0x6a0b6881.
//
// Solidity: function deployNodeELRewardVault(uint8 _poolId, uint256 _operatorId) returns(address)
func (_VaultFactory *VaultFactoryTransactor) DeployNodeELRewardVault(opts *bind.TransactOpts, _poolId uint8, _operatorId *big.Int) (*types.Transaction, error) {
	return _VaultFactory.contract.Transact(opts, "deployNodeELRewardVault", _poolId, _operatorId)
}

// DeployNodeELRewardVault is a paid mutator transaction binding the contract method 0x6a0b6881.
//
// Solidity: function deployNodeELRewardVault(uint8 _poolId, uint256 _operatorId) returns(address)
func (_VaultFactory *VaultFactorySession) DeployNodeELRewardVault(_poolId uint8, _operatorId *big.Int) (*types.Transaction, error) {
	return _VaultFactory.Contract.DeployNodeELRewardVault(&_VaultFactory.TransactOpts, _poolId, _operatorId)
}

// DeployNodeELRewardVault is a paid mutator transaction binding the contract method 0x6a0b6881.
//
// Solidity: function deployNodeELRewardVault(uint8 _poolId, uint256 _operatorId) returns(address)
func (_VaultFactory *VaultFactoryTransactorSession) DeployNodeELRewardVault(_poolId uint8, _operatorId *big.Int) (*types.Transaction, error) {
	return _VaultFactory.Contract.DeployNodeELRewardVault(&_VaultFactory.TransactOpts, _poolId, _operatorId)
}

// DeployWithdrawVault is a paid mutator transaction binding the contract method 0x7f70ce0d.
//
// Solidity: function deployWithdrawVault(uint8 _poolId, uint256 _operatorId, uint256 _validatorCount, uint256 _validatorId) returns(address)
func (_VaultFactory *VaultFactoryTransactor) DeployWithdrawVault(opts *bind.TransactOpts, _poolId uint8, _operatorId *big.Int, _validatorCount *big.Int, _validatorId *big.Int) (*types.Transaction, error) {
	return _VaultFactory.contract.Transact(opts, "deployWithdrawVault", _poolId, _operatorId, _validatorCount, _validatorId)
}

// DeployWithdrawVault is a paid mutator transaction binding the contract method 0x7f70ce0d.
//
// Solidity: function deployWithdrawVault(uint8 _poolId, uint256 _operatorId, uint256 _validatorCount, uint256 _validatorId) returns(address)
func (_VaultFactory *VaultFactorySession) DeployWithdrawVault(_poolId uint8, _operatorId *big.Int, _validatorCount *big.Int, _validatorId *big.Int) (*types.Transaction, error) {
	return _VaultFactory.Contract.DeployWithdrawVault(&_VaultFactory.TransactOpts, _poolId, _operatorId, _validatorCount, _validatorId)
}

// DeployWithdrawVault is a paid mutator transaction binding the contract method 0x7f70ce0d.
//
// Solidity: function deployWithdrawVault(uint8 _poolId, uint256 _operatorId, uint256 _validatorCount, uint256 _validatorId) returns(address)
func (_VaultFactory *VaultFactoryTransactorSession) DeployWithdrawVault(_poolId uint8, _operatorId *big.Int, _validatorCount *big.Int, _validatorId *big.Int) (*types.Transaction, error) {
	return _VaultFactory.Contract.DeployWithdrawVault(&_VaultFactory.TransactOpts, _poolId, _operatorId, _validatorCount, _validatorId)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_VaultFactory *VaultFactoryTransactor) GrantRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _VaultFactory.contract.Transact(opts, "grantRole", role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_VaultFactory *VaultFactorySession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _VaultFactory.Contract.GrantRole(&_VaultFactory.TransactOpts, role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_VaultFactory *VaultFactoryTransactorSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _VaultFactory.Contract.GrantRole(&_VaultFactory.TransactOpts, role, account)
}

// Initialize is a paid mutator transaction binding the contract method 0x485cc955.
//
// Solidity: function initialize(address _admin, address _staderConfig) returns()
func (_VaultFactory *VaultFactoryTransactor) Initialize(opts *bind.TransactOpts, _admin common.Address, _staderConfig common.Address) (*types.Transaction, error) {
	return _VaultFactory.contract.Transact(opts, "initialize", _admin, _staderConfig)
}

// Initialize is a paid mutator transaction binding the contract method 0x485cc955.
//
// Solidity: function initialize(address _admin, address _staderConfig) returns()
func (_VaultFactory *VaultFactorySession) Initialize(_admin common.Address, _staderConfig common.Address) (*types.Transaction, error) {
	return _VaultFactory.Contract.Initialize(&_VaultFactory.TransactOpts, _admin, _staderConfig)
}

// Initialize is a paid mutator transaction binding the contract method 0x485cc955.
//
// Solidity: function initialize(address _admin, address _staderConfig) returns()
func (_VaultFactory *VaultFactoryTransactorSession) Initialize(_admin common.Address, _staderConfig common.Address) (*types.Transaction, error) {
	return _VaultFactory.Contract.Initialize(&_VaultFactory.TransactOpts, _admin, _staderConfig)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address account) returns()
func (_VaultFactory *VaultFactoryTransactor) RenounceRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _VaultFactory.contract.Transact(opts, "renounceRole", role, account)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address account) returns()
func (_VaultFactory *VaultFactorySession) RenounceRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _VaultFactory.Contract.RenounceRole(&_VaultFactory.TransactOpts, role, account)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address account) returns()
func (_VaultFactory *VaultFactoryTransactorSession) RenounceRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _VaultFactory.Contract.RenounceRole(&_VaultFactory.TransactOpts, role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_VaultFactory *VaultFactoryTransactor) RevokeRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _VaultFactory.contract.Transact(opts, "revokeRole", role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_VaultFactory *VaultFactorySession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _VaultFactory.Contract.RevokeRole(&_VaultFactory.TransactOpts, role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_VaultFactory *VaultFactoryTransactorSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _VaultFactory.Contract.RevokeRole(&_VaultFactory.TransactOpts, role, account)
}

// UpdateStaderConfig is a paid mutator transaction binding the contract method 0x9ee804cb.
//
// Solidity: function updateStaderConfig(address _staderConfig) returns()
func (_VaultFactory *VaultFactoryTransactor) UpdateStaderConfig(opts *bind.TransactOpts, _staderConfig common.Address) (*types.Transaction, error) {
	return _VaultFactory.contract.Transact(opts, "updateStaderConfig", _staderConfig)
}

// UpdateStaderConfig is a paid mutator transaction binding the contract method 0x9ee804cb.
//
// Solidity: function updateStaderConfig(address _staderConfig) returns()
func (_VaultFactory *VaultFactorySession) UpdateStaderConfig(_staderConfig common.Address) (*types.Transaction, error) {
	return _VaultFactory.Contract.UpdateStaderConfig(&_VaultFactory.TransactOpts, _staderConfig)
}

// UpdateStaderConfig is a paid mutator transaction binding the contract method 0x9ee804cb.
//
// Solidity: function updateStaderConfig(address _staderConfig) returns()
func (_VaultFactory *VaultFactoryTransactorSession) UpdateStaderConfig(_staderConfig common.Address) (*types.Transaction, error) {
	return _VaultFactory.Contract.UpdateStaderConfig(&_VaultFactory.TransactOpts, _staderConfig)
}

// VaultFactoryInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the VaultFactory contract.
type VaultFactoryInitializedIterator struct {
	Event *VaultFactoryInitialized // Event containing the contract specifics and raw log

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
func (it *VaultFactoryInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VaultFactoryInitialized)
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
		it.Event = new(VaultFactoryInitialized)
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
func (it *VaultFactoryInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VaultFactoryInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VaultFactoryInitialized represents a Initialized event raised by the VaultFactory contract.
type VaultFactoryInitialized struct {
	Version uint8
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_VaultFactory *VaultFactoryFilterer) FilterInitialized(opts *bind.FilterOpts) (*VaultFactoryInitializedIterator, error) {

	logs, sub, err := _VaultFactory.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &VaultFactoryInitializedIterator{contract: _VaultFactory.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_VaultFactory *VaultFactoryFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *VaultFactoryInitialized) (event.Subscription, error) {

	logs, sub, err := _VaultFactory.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VaultFactoryInitialized)
				if err := _VaultFactory.contract.UnpackLog(event, "Initialized", log); err != nil {
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
func (_VaultFactory *VaultFactoryFilterer) ParseInitialized(log types.Log) (*VaultFactoryInitialized, error) {
	event := new(VaultFactoryInitialized)
	if err := _VaultFactory.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// VaultFactoryNodeELRewardVaultCreatedIterator is returned from FilterNodeELRewardVaultCreated and is used to iterate over the raw logs and unpacked data for NodeELRewardVaultCreated events raised by the VaultFactory contract.
type VaultFactoryNodeELRewardVaultCreatedIterator struct {
	Event *VaultFactoryNodeELRewardVaultCreated // Event containing the contract specifics and raw log

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
func (it *VaultFactoryNodeELRewardVaultCreatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VaultFactoryNodeELRewardVaultCreated)
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
		it.Event = new(VaultFactoryNodeELRewardVaultCreated)
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
func (it *VaultFactoryNodeELRewardVaultCreatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VaultFactoryNodeELRewardVaultCreatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VaultFactoryNodeELRewardVaultCreated represents a NodeELRewardVaultCreated event raised by the VaultFactory contract.
type VaultFactoryNodeELRewardVaultCreated struct {
	NodeDistributor common.Address
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterNodeELRewardVaultCreated is a free log retrieval operation binding the contract event 0x534e63617d48745c45efd754f53f9292453257107e03d0140db96fc8e3bf7ec9.
//
// Solidity: event NodeELRewardVaultCreated(address nodeDistributor)
func (_VaultFactory *VaultFactoryFilterer) FilterNodeELRewardVaultCreated(opts *bind.FilterOpts) (*VaultFactoryNodeELRewardVaultCreatedIterator, error) {

	logs, sub, err := _VaultFactory.contract.FilterLogs(opts, "NodeELRewardVaultCreated")
	if err != nil {
		return nil, err
	}
	return &VaultFactoryNodeELRewardVaultCreatedIterator{contract: _VaultFactory.contract, event: "NodeELRewardVaultCreated", logs: logs, sub: sub}, nil
}

// WatchNodeELRewardVaultCreated is a free log subscription operation binding the contract event 0x534e63617d48745c45efd754f53f9292453257107e03d0140db96fc8e3bf7ec9.
//
// Solidity: event NodeELRewardVaultCreated(address nodeDistributor)
func (_VaultFactory *VaultFactoryFilterer) WatchNodeELRewardVaultCreated(opts *bind.WatchOpts, sink chan<- *VaultFactoryNodeELRewardVaultCreated) (event.Subscription, error) {

	logs, sub, err := _VaultFactory.contract.WatchLogs(opts, "NodeELRewardVaultCreated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VaultFactoryNodeELRewardVaultCreated)
				if err := _VaultFactory.contract.UnpackLog(event, "NodeELRewardVaultCreated", log); err != nil {
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

// ParseNodeELRewardVaultCreated is a log parse operation binding the contract event 0x534e63617d48745c45efd754f53f9292453257107e03d0140db96fc8e3bf7ec9.
//
// Solidity: event NodeELRewardVaultCreated(address nodeDistributor)
func (_VaultFactory *VaultFactoryFilterer) ParseNodeELRewardVaultCreated(log types.Log) (*VaultFactoryNodeELRewardVaultCreated, error) {
	event := new(VaultFactoryNodeELRewardVaultCreated)
	if err := _VaultFactory.contract.UnpackLog(event, "NodeELRewardVaultCreated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// VaultFactoryRoleAdminChangedIterator is returned from FilterRoleAdminChanged and is used to iterate over the raw logs and unpacked data for RoleAdminChanged events raised by the VaultFactory contract.
type VaultFactoryRoleAdminChangedIterator struct {
	Event *VaultFactoryRoleAdminChanged // Event containing the contract specifics and raw log

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
func (it *VaultFactoryRoleAdminChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VaultFactoryRoleAdminChanged)
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
		it.Event = new(VaultFactoryRoleAdminChanged)
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
func (it *VaultFactoryRoleAdminChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VaultFactoryRoleAdminChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VaultFactoryRoleAdminChanged represents a RoleAdminChanged event raised by the VaultFactory contract.
type VaultFactoryRoleAdminChanged struct {
	Role              [32]byte
	PreviousAdminRole [32]byte
	NewAdminRole      [32]byte
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterRoleAdminChanged is a free log retrieval operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_VaultFactory *VaultFactoryFilterer) FilterRoleAdminChanged(opts *bind.FilterOpts, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (*VaultFactoryRoleAdminChangedIterator, error) {

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

	logs, sub, err := _VaultFactory.contract.FilterLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return &VaultFactoryRoleAdminChangedIterator{contract: _VaultFactory.contract, event: "RoleAdminChanged", logs: logs, sub: sub}, nil
}

// WatchRoleAdminChanged is a free log subscription operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_VaultFactory *VaultFactoryFilterer) WatchRoleAdminChanged(opts *bind.WatchOpts, sink chan<- *VaultFactoryRoleAdminChanged, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (event.Subscription, error) {

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

	logs, sub, err := _VaultFactory.contract.WatchLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VaultFactoryRoleAdminChanged)
				if err := _VaultFactory.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
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
func (_VaultFactory *VaultFactoryFilterer) ParseRoleAdminChanged(log types.Log) (*VaultFactoryRoleAdminChanged, error) {
	event := new(VaultFactoryRoleAdminChanged)
	if err := _VaultFactory.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// VaultFactoryRoleGrantedIterator is returned from FilterRoleGranted and is used to iterate over the raw logs and unpacked data for RoleGranted events raised by the VaultFactory contract.
type VaultFactoryRoleGrantedIterator struct {
	Event *VaultFactoryRoleGranted // Event containing the contract specifics and raw log

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
func (it *VaultFactoryRoleGrantedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VaultFactoryRoleGranted)
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
		it.Event = new(VaultFactoryRoleGranted)
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
func (it *VaultFactoryRoleGrantedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VaultFactoryRoleGrantedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VaultFactoryRoleGranted represents a RoleGranted event raised by the VaultFactory contract.
type VaultFactoryRoleGranted struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleGranted is a free log retrieval operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_VaultFactory *VaultFactoryFilterer) FilterRoleGranted(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*VaultFactoryRoleGrantedIterator, error) {

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

	logs, sub, err := _VaultFactory.contract.FilterLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &VaultFactoryRoleGrantedIterator{contract: _VaultFactory.contract, event: "RoleGranted", logs: logs, sub: sub}, nil
}

// WatchRoleGranted is a free log subscription operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_VaultFactory *VaultFactoryFilterer) WatchRoleGranted(opts *bind.WatchOpts, sink chan<- *VaultFactoryRoleGranted, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _VaultFactory.contract.WatchLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VaultFactoryRoleGranted)
				if err := _VaultFactory.contract.UnpackLog(event, "RoleGranted", log); err != nil {
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
func (_VaultFactory *VaultFactoryFilterer) ParseRoleGranted(log types.Log) (*VaultFactoryRoleGranted, error) {
	event := new(VaultFactoryRoleGranted)
	if err := _VaultFactory.contract.UnpackLog(event, "RoleGranted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// VaultFactoryRoleRevokedIterator is returned from FilterRoleRevoked and is used to iterate over the raw logs and unpacked data for RoleRevoked events raised by the VaultFactory contract.
type VaultFactoryRoleRevokedIterator struct {
	Event *VaultFactoryRoleRevoked // Event containing the contract specifics and raw log

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
func (it *VaultFactoryRoleRevokedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VaultFactoryRoleRevoked)
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
		it.Event = new(VaultFactoryRoleRevoked)
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
func (it *VaultFactoryRoleRevokedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VaultFactoryRoleRevokedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VaultFactoryRoleRevoked represents a RoleRevoked event raised by the VaultFactory contract.
type VaultFactoryRoleRevoked struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleRevoked is a free log retrieval operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_VaultFactory *VaultFactoryFilterer) FilterRoleRevoked(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*VaultFactoryRoleRevokedIterator, error) {

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

	logs, sub, err := _VaultFactory.contract.FilterLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &VaultFactoryRoleRevokedIterator{contract: _VaultFactory.contract, event: "RoleRevoked", logs: logs, sub: sub}, nil
}

// WatchRoleRevoked is a free log subscription operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_VaultFactory *VaultFactoryFilterer) WatchRoleRevoked(opts *bind.WatchOpts, sink chan<- *VaultFactoryRoleRevoked, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _VaultFactory.contract.WatchLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VaultFactoryRoleRevoked)
				if err := _VaultFactory.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
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
func (_VaultFactory *VaultFactoryFilterer) ParseRoleRevoked(log types.Log) (*VaultFactoryRoleRevoked, error) {
	event := new(VaultFactoryRoleRevoked)
	if err := _VaultFactory.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// VaultFactoryUpdatedStaderConfigIterator is returned from FilterUpdatedStaderConfig and is used to iterate over the raw logs and unpacked data for UpdatedStaderConfig events raised by the VaultFactory contract.
type VaultFactoryUpdatedStaderConfigIterator struct {
	Event *VaultFactoryUpdatedStaderConfig // Event containing the contract specifics and raw log

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
func (it *VaultFactoryUpdatedStaderConfigIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VaultFactoryUpdatedStaderConfig)
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
		it.Event = new(VaultFactoryUpdatedStaderConfig)
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
func (it *VaultFactoryUpdatedStaderConfigIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VaultFactoryUpdatedStaderConfigIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VaultFactoryUpdatedStaderConfig represents a UpdatedStaderConfig event raised by the VaultFactory contract.
type VaultFactoryUpdatedStaderConfig struct {
	StaderConfig common.Address
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterUpdatedStaderConfig is a free log retrieval operation binding the contract event 0xdb2219043d7b197cb235f1af0cf6d782d77dee3de19e3f4fb6d39aae633b4485.
//
// Solidity: event UpdatedStaderConfig(address staderConfig)
func (_VaultFactory *VaultFactoryFilterer) FilterUpdatedStaderConfig(opts *bind.FilterOpts) (*VaultFactoryUpdatedStaderConfigIterator, error) {

	logs, sub, err := _VaultFactory.contract.FilterLogs(opts, "UpdatedStaderConfig")
	if err != nil {
		return nil, err
	}
	return &VaultFactoryUpdatedStaderConfigIterator{contract: _VaultFactory.contract, event: "UpdatedStaderConfig", logs: logs, sub: sub}, nil
}

// WatchUpdatedStaderConfig is a free log subscription operation binding the contract event 0xdb2219043d7b197cb235f1af0cf6d782d77dee3de19e3f4fb6d39aae633b4485.
//
// Solidity: event UpdatedStaderConfig(address staderConfig)
func (_VaultFactory *VaultFactoryFilterer) WatchUpdatedStaderConfig(opts *bind.WatchOpts, sink chan<- *VaultFactoryUpdatedStaderConfig) (event.Subscription, error) {

	logs, sub, err := _VaultFactory.contract.WatchLogs(opts, "UpdatedStaderConfig")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VaultFactoryUpdatedStaderConfig)
				if err := _VaultFactory.contract.UnpackLog(event, "UpdatedStaderConfig", log); err != nil {
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
func (_VaultFactory *VaultFactoryFilterer) ParseUpdatedStaderConfig(log types.Log) (*VaultFactoryUpdatedStaderConfig, error) {
	event := new(VaultFactoryUpdatedStaderConfig)
	if err := _VaultFactory.contract.UnpackLog(event, "UpdatedStaderConfig", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// VaultFactoryWithdrawVaultCreatedIterator is returned from FilterWithdrawVaultCreated and is used to iterate over the raw logs and unpacked data for WithdrawVaultCreated events raised by the VaultFactory contract.
type VaultFactoryWithdrawVaultCreatedIterator struct {
	Event *VaultFactoryWithdrawVaultCreated // Event containing the contract specifics and raw log

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
func (it *VaultFactoryWithdrawVaultCreatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VaultFactoryWithdrawVaultCreated)
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
		it.Event = new(VaultFactoryWithdrawVaultCreated)
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
func (it *VaultFactoryWithdrawVaultCreatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VaultFactoryWithdrawVaultCreatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VaultFactoryWithdrawVaultCreated represents a WithdrawVaultCreated event raised by the VaultFactory contract.
type VaultFactoryWithdrawVaultCreated struct {
	WithdrawVault common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterWithdrawVaultCreated is a free log retrieval operation binding the contract event 0x3961159e01aea37eec106905b883c13f18c7c19abfde04fef6835c94d2af51b9.
//
// Solidity: event WithdrawVaultCreated(address withdrawVault)
func (_VaultFactory *VaultFactoryFilterer) FilterWithdrawVaultCreated(opts *bind.FilterOpts) (*VaultFactoryWithdrawVaultCreatedIterator, error) {

	logs, sub, err := _VaultFactory.contract.FilterLogs(opts, "WithdrawVaultCreated")
	if err != nil {
		return nil, err
	}
	return &VaultFactoryWithdrawVaultCreatedIterator{contract: _VaultFactory.contract, event: "WithdrawVaultCreated", logs: logs, sub: sub}, nil
}

// WatchWithdrawVaultCreated is a free log subscription operation binding the contract event 0x3961159e01aea37eec106905b883c13f18c7c19abfde04fef6835c94d2af51b9.
//
// Solidity: event WithdrawVaultCreated(address withdrawVault)
func (_VaultFactory *VaultFactoryFilterer) WatchWithdrawVaultCreated(opts *bind.WatchOpts, sink chan<- *VaultFactoryWithdrawVaultCreated) (event.Subscription, error) {

	logs, sub, err := _VaultFactory.contract.WatchLogs(opts, "WithdrawVaultCreated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VaultFactoryWithdrawVaultCreated)
				if err := _VaultFactory.contract.UnpackLog(event, "WithdrawVaultCreated", log); err != nil {
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

// ParseWithdrawVaultCreated is a log parse operation binding the contract event 0x3961159e01aea37eec106905b883c13f18c7c19abfde04fef6835c94d2af51b9.
//
// Solidity: event WithdrawVaultCreated(address withdrawVault)
func (_VaultFactory *VaultFactoryFilterer) ParseWithdrawVaultCreated(log types.Log) (*VaultFactoryWithdrawVaultCreated, error) {
	event := new(VaultFactoryWithdrawVaultCreated)
	if err := _VaultFactory.contract.UnpackLog(event, "WithdrawVaultCreated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
