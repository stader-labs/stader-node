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

// ValidatorWithdrawVaultMetaData contains all meta data concerning the ValidatorWithdrawVault contract.
var ValidatorWithdrawVaultMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"CallerNotNodeRegistryContract\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"ETHTransferFailed\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InsufficientBalance\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidRewardAmount\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotEnoughRewardToDistribute\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ZeroAddress\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"rewardAmount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"rewardThreshold\",\"type\":\"uint256\"}],\"name\":\"DistributeRewardFailed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"userShare\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"operatorShare\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"protocolShare\",\"type\":\"uint256\"}],\"name\":\"DistributedRewards\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"ETHReceived\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"version\",\"type\":\"uint8\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"previousAdminRole\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"newAdminRole\",\"type\":\"bytes32\"}],\"name\":\"RoleAdminChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleGranted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleRevoked\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"userShare\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"operatorShare\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"protocolShare\",\"type\":\"uint256\"}],\"name\":\"SettledFunds\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_staderConfig\",\"type\":\"address\"}],\"name\":\"UpdatedStaderConfig\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"DEFAULT_ADMIN_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"calculateValidatorWithdrawalShare\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"_userShare\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_operatorShare\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_protocolShare\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"distributeRewards\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"}],\"name\":\"getRoleAdmin\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"grantRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"hasRole\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"_poolId\",\"type\":\"uint8\"},{\"internalType\":\"address\",\"name\":\"_staderConfig\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_validatorId\",\"type\":\"uint256\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"poolId\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"renounceRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"revokeRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"settleFunds\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"staderConfig\",\"outputs\":[{\"internalType\":\"contractIStaderConfig\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_staderConfig\",\"type\":\"address\"}],\"name\":\"updateStaderConfig\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"validatorId\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"vaultSettleStatus\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"stateMutability\":\"payable\",\"type\":\"receive\"}]",
}

// ValidatorWithdrawVaultABI is the input ABI used to generate the binding from.
// Deprecated: Use ValidatorWithdrawVaultMetaData.ABI instead.
var ValidatorWithdrawVaultABI = ValidatorWithdrawVaultMetaData.ABI

// ValidatorWithdrawVault is an auto generated Go binding around an Ethereum contract.
type ValidatorWithdrawVault struct {
	ValidatorWithdrawVaultCaller     // Read-only binding to the contract
	ValidatorWithdrawVaultTransactor // Write-only binding to the contract
	ValidatorWithdrawVaultFilterer   // Log filterer for contract events
}

// ValidatorWithdrawVaultCaller is an auto generated read-only Go binding around an Ethereum contract.
type ValidatorWithdrawVaultCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ValidatorWithdrawVaultTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ValidatorWithdrawVaultTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ValidatorWithdrawVaultFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ValidatorWithdrawVaultFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ValidatorWithdrawVaultSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ValidatorWithdrawVaultSession struct {
	Contract     *ValidatorWithdrawVault // Generic contract binding to set the session for
	CallOpts     bind.CallOpts           // Call options to use throughout this session
	TransactOpts bind.TransactOpts       // Transaction auth options to use throughout this session
}

// ValidatorWithdrawVaultCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ValidatorWithdrawVaultCallerSession struct {
	Contract *ValidatorWithdrawVaultCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                 // Call options to use throughout this session
}

// ValidatorWithdrawVaultTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ValidatorWithdrawVaultTransactorSession struct {
	Contract     *ValidatorWithdrawVaultTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                 // Transaction auth options to use throughout this session
}

// ValidatorWithdrawVaultRaw is an auto generated low-level Go binding around an Ethereum contract.
type ValidatorWithdrawVaultRaw struct {
	Contract *ValidatorWithdrawVault // Generic contract binding to access the raw methods on
}

// ValidatorWithdrawVaultCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ValidatorWithdrawVaultCallerRaw struct {
	Contract *ValidatorWithdrawVaultCaller // Generic read-only contract binding to access the raw methods on
}

// ValidatorWithdrawVaultTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ValidatorWithdrawVaultTransactorRaw struct {
	Contract *ValidatorWithdrawVaultTransactor // Generic write-only contract binding to access the raw methods on
}

// NewValidatorWithdrawVault creates a new instance of ValidatorWithdrawVault, bound to a specific deployed contract.
func NewValidatorWithdrawVault(address common.Address, backend bind.ContractBackend) (*ValidatorWithdrawVault, error) {
	contract, err := bindValidatorWithdrawVault(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ValidatorWithdrawVault{ValidatorWithdrawVaultCaller: ValidatorWithdrawVaultCaller{contract: contract}, ValidatorWithdrawVaultTransactor: ValidatorWithdrawVaultTransactor{contract: contract}, ValidatorWithdrawVaultFilterer: ValidatorWithdrawVaultFilterer{contract: contract}}, nil
}

// NewValidatorWithdrawVaultCaller creates a new read-only instance of ValidatorWithdrawVault, bound to a specific deployed contract.
func NewValidatorWithdrawVaultCaller(address common.Address, caller bind.ContractCaller) (*ValidatorWithdrawVaultCaller, error) {
	contract, err := bindValidatorWithdrawVault(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ValidatorWithdrawVaultCaller{contract: contract}, nil
}

// NewValidatorWithdrawVaultTransactor creates a new write-only instance of ValidatorWithdrawVault, bound to a specific deployed contract.
func NewValidatorWithdrawVaultTransactor(address common.Address, transactor bind.ContractTransactor) (*ValidatorWithdrawVaultTransactor, error) {
	contract, err := bindValidatorWithdrawVault(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ValidatorWithdrawVaultTransactor{contract: contract}, nil
}

// NewValidatorWithdrawVaultFilterer creates a new log filterer instance of ValidatorWithdrawVault, bound to a specific deployed contract.
func NewValidatorWithdrawVaultFilterer(address common.Address, filterer bind.ContractFilterer) (*ValidatorWithdrawVaultFilterer, error) {
	contract, err := bindValidatorWithdrawVault(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ValidatorWithdrawVaultFilterer{contract: contract}, nil
}

// bindValidatorWithdrawVault binds a generic wrapper to an already deployed contract.
func bindValidatorWithdrawVault(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := ValidatorWithdrawVaultMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ValidatorWithdrawVault *ValidatorWithdrawVaultRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ValidatorWithdrawVault.Contract.ValidatorWithdrawVaultCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ValidatorWithdrawVault *ValidatorWithdrawVaultRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ValidatorWithdrawVault.Contract.ValidatorWithdrawVaultTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ValidatorWithdrawVault *ValidatorWithdrawVaultRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ValidatorWithdrawVault.Contract.ValidatorWithdrawVaultTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ValidatorWithdrawVault *ValidatorWithdrawVaultCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ValidatorWithdrawVault.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ValidatorWithdrawVault *ValidatorWithdrawVaultTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ValidatorWithdrawVault.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ValidatorWithdrawVault *ValidatorWithdrawVaultTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ValidatorWithdrawVault.Contract.contract.Transact(opts, method, params...)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_ValidatorWithdrawVault *ValidatorWithdrawVaultCaller) DEFAULTADMINROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _ValidatorWithdrawVault.contract.Call(opts, &out, "DEFAULT_ADMIN_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_ValidatorWithdrawVault *ValidatorWithdrawVaultSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _ValidatorWithdrawVault.Contract.DEFAULTADMINROLE(&_ValidatorWithdrawVault.CallOpts)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_ValidatorWithdrawVault *ValidatorWithdrawVaultCallerSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _ValidatorWithdrawVault.Contract.DEFAULTADMINROLE(&_ValidatorWithdrawVault.CallOpts)
}

// CalculateValidatorWithdrawalShare is a free data retrieval call binding the contract method 0x99997b70.
//
// Solidity: function calculateValidatorWithdrawalShare() view returns(uint256 _userShare, uint256 _operatorShare, uint256 _protocolShare)
func (_ValidatorWithdrawVault *ValidatorWithdrawVaultCaller) CalculateValidatorWithdrawalShare(opts *bind.CallOpts) (struct {
	UserShare     *big.Int
	OperatorShare *big.Int
	ProtocolShare *big.Int
}, error) {
	var out []interface{}
	err := _ValidatorWithdrawVault.contract.Call(opts, &out, "calculateValidatorWithdrawalShare")

	outstruct := new(struct {
		UserShare     *big.Int
		OperatorShare *big.Int
		ProtocolShare *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.UserShare = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.OperatorShare = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.ProtocolShare = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// CalculateValidatorWithdrawalShare is a free data retrieval call binding the contract method 0x99997b70.
//
// Solidity: function calculateValidatorWithdrawalShare() view returns(uint256 _userShare, uint256 _operatorShare, uint256 _protocolShare)
func (_ValidatorWithdrawVault *ValidatorWithdrawVaultSession) CalculateValidatorWithdrawalShare() (struct {
	UserShare     *big.Int
	OperatorShare *big.Int
	ProtocolShare *big.Int
}, error) {
	return _ValidatorWithdrawVault.Contract.CalculateValidatorWithdrawalShare(&_ValidatorWithdrawVault.CallOpts)
}

// CalculateValidatorWithdrawalShare is a free data retrieval call binding the contract method 0x99997b70.
//
// Solidity: function calculateValidatorWithdrawalShare() view returns(uint256 _userShare, uint256 _operatorShare, uint256 _protocolShare)
func (_ValidatorWithdrawVault *ValidatorWithdrawVaultCallerSession) CalculateValidatorWithdrawalShare() (struct {
	UserShare     *big.Int
	OperatorShare *big.Int
	ProtocolShare *big.Int
}, error) {
	return _ValidatorWithdrawVault.Contract.CalculateValidatorWithdrawalShare(&_ValidatorWithdrawVault.CallOpts)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_ValidatorWithdrawVault *ValidatorWithdrawVaultCaller) GetRoleAdmin(opts *bind.CallOpts, role [32]byte) ([32]byte, error) {
	var out []interface{}
	err := _ValidatorWithdrawVault.contract.Call(opts, &out, "getRoleAdmin", role)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_ValidatorWithdrawVault *ValidatorWithdrawVaultSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _ValidatorWithdrawVault.Contract.GetRoleAdmin(&_ValidatorWithdrawVault.CallOpts, role)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_ValidatorWithdrawVault *ValidatorWithdrawVaultCallerSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _ValidatorWithdrawVault.Contract.GetRoleAdmin(&_ValidatorWithdrawVault.CallOpts, role)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_ValidatorWithdrawVault *ValidatorWithdrawVaultCaller) HasRole(opts *bind.CallOpts, role [32]byte, account common.Address) (bool, error) {
	var out []interface{}
	err := _ValidatorWithdrawVault.contract.Call(opts, &out, "hasRole", role, account)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_ValidatorWithdrawVault *ValidatorWithdrawVaultSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _ValidatorWithdrawVault.Contract.HasRole(&_ValidatorWithdrawVault.CallOpts, role, account)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_ValidatorWithdrawVault *ValidatorWithdrawVaultCallerSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _ValidatorWithdrawVault.Contract.HasRole(&_ValidatorWithdrawVault.CallOpts, role, account)
}

// PoolId is a free data retrieval call binding the contract method 0x3e0dc34e.
//
// Solidity: function poolId() view returns(uint8)
func (_ValidatorWithdrawVault *ValidatorWithdrawVaultCaller) PoolId(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _ValidatorWithdrawVault.contract.Call(opts, &out, "poolId")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// PoolId is a free data retrieval call binding the contract method 0x3e0dc34e.
//
// Solidity: function poolId() view returns(uint8)
func (_ValidatorWithdrawVault *ValidatorWithdrawVaultSession) PoolId() (uint8, error) {
	return _ValidatorWithdrawVault.Contract.PoolId(&_ValidatorWithdrawVault.CallOpts)
}

// PoolId is a free data retrieval call binding the contract method 0x3e0dc34e.
//
// Solidity: function poolId() view returns(uint8)
func (_ValidatorWithdrawVault *ValidatorWithdrawVaultCallerSession) PoolId() (uint8, error) {
	return _ValidatorWithdrawVault.Contract.PoolId(&_ValidatorWithdrawVault.CallOpts)
}

// StaderConfig is a free data retrieval call binding the contract method 0x490ffa35.
//
// Solidity: function staderConfig() view returns(address)
func (_ValidatorWithdrawVault *ValidatorWithdrawVaultCaller) StaderConfig(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ValidatorWithdrawVault.contract.Call(opts, &out, "staderConfig")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// StaderConfig is a free data retrieval call binding the contract method 0x490ffa35.
//
// Solidity: function staderConfig() view returns(address)
func (_ValidatorWithdrawVault *ValidatorWithdrawVaultSession) StaderConfig() (common.Address, error) {
	return _ValidatorWithdrawVault.Contract.StaderConfig(&_ValidatorWithdrawVault.CallOpts)
}

// StaderConfig is a free data retrieval call binding the contract method 0x490ffa35.
//
// Solidity: function staderConfig() view returns(address)
func (_ValidatorWithdrawVault *ValidatorWithdrawVaultCallerSession) StaderConfig() (common.Address, error) {
	return _ValidatorWithdrawVault.Contract.StaderConfig(&_ValidatorWithdrawVault.CallOpts)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_ValidatorWithdrawVault *ValidatorWithdrawVaultCaller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _ValidatorWithdrawVault.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_ValidatorWithdrawVault *ValidatorWithdrawVaultSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _ValidatorWithdrawVault.Contract.SupportsInterface(&_ValidatorWithdrawVault.CallOpts, interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_ValidatorWithdrawVault *ValidatorWithdrawVaultCallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _ValidatorWithdrawVault.Contract.SupportsInterface(&_ValidatorWithdrawVault.CallOpts, interfaceId)
}

// ValidatorId is a free data retrieval call binding the contract method 0x5c5f7dae.
//
// Solidity: function validatorId() view returns(uint256)
func (_ValidatorWithdrawVault *ValidatorWithdrawVaultCaller) ValidatorId(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _ValidatorWithdrawVault.contract.Call(opts, &out, "validatorId")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ValidatorId is a free data retrieval call binding the contract method 0x5c5f7dae.
//
// Solidity: function validatorId() view returns(uint256)
func (_ValidatorWithdrawVault *ValidatorWithdrawVaultSession) ValidatorId() (*big.Int, error) {
	return _ValidatorWithdrawVault.Contract.ValidatorId(&_ValidatorWithdrawVault.CallOpts)
}

// ValidatorId is a free data retrieval call binding the contract method 0x5c5f7dae.
//
// Solidity: function validatorId() view returns(uint256)
func (_ValidatorWithdrawVault *ValidatorWithdrawVaultCallerSession) ValidatorId() (*big.Int, error) {
	return _ValidatorWithdrawVault.Contract.ValidatorId(&_ValidatorWithdrawVault.CallOpts)
}

// VaultSettleStatus is a free data retrieval call binding the contract method 0x7ef4947d.
//
// Solidity: function vaultSettleStatus() view returns(bool)
func (_ValidatorWithdrawVault *ValidatorWithdrawVaultCaller) VaultSettleStatus(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _ValidatorWithdrawVault.contract.Call(opts, &out, "vaultSettleStatus")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// VaultSettleStatus is a free data retrieval call binding the contract method 0x7ef4947d.
//
// Solidity: function vaultSettleStatus() view returns(bool)
func (_ValidatorWithdrawVault *ValidatorWithdrawVaultSession) VaultSettleStatus() (bool, error) {
	return _ValidatorWithdrawVault.Contract.VaultSettleStatus(&_ValidatorWithdrawVault.CallOpts)
}

// VaultSettleStatus is a free data retrieval call binding the contract method 0x7ef4947d.
//
// Solidity: function vaultSettleStatus() view returns(bool)
func (_ValidatorWithdrawVault *ValidatorWithdrawVaultCallerSession) VaultSettleStatus() (bool, error) {
	return _ValidatorWithdrawVault.Contract.VaultSettleStatus(&_ValidatorWithdrawVault.CallOpts)
}

// DistributeRewards is a paid mutator transaction binding the contract method 0x6f4a2cd0.
//
// Solidity: function distributeRewards() returns()
func (_ValidatorWithdrawVault *ValidatorWithdrawVaultTransactor) DistributeRewards(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ValidatorWithdrawVault.contract.Transact(opts, "distributeRewards")
}

// DistributeRewards is a paid mutator transaction binding the contract method 0x6f4a2cd0.
//
// Solidity: function distributeRewards() returns()
func (_ValidatorWithdrawVault *ValidatorWithdrawVaultSession) DistributeRewards() (*types.Transaction, error) {
	return _ValidatorWithdrawVault.Contract.DistributeRewards(&_ValidatorWithdrawVault.TransactOpts)
}

// DistributeRewards is a paid mutator transaction binding the contract method 0x6f4a2cd0.
//
// Solidity: function distributeRewards() returns()
func (_ValidatorWithdrawVault *ValidatorWithdrawVaultTransactorSession) DistributeRewards() (*types.Transaction, error) {
	return _ValidatorWithdrawVault.Contract.DistributeRewards(&_ValidatorWithdrawVault.TransactOpts)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_ValidatorWithdrawVault *ValidatorWithdrawVaultTransactor) GrantRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _ValidatorWithdrawVault.contract.Transact(opts, "grantRole", role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_ValidatorWithdrawVault *ValidatorWithdrawVaultSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _ValidatorWithdrawVault.Contract.GrantRole(&_ValidatorWithdrawVault.TransactOpts, role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_ValidatorWithdrawVault *ValidatorWithdrawVaultTransactorSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _ValidatorWithdrawVault.Contract.GrantRole(&_ValidatorWithdrawVault.TransactOpts, role, account)
}

// Initialize is a paid mutator transaction binding the contract method 0x06d39109.
//
// Solidity: function initialize(uint8 _poolId, address _staderConfig, uint256 _validatorId) returns()
func (_ValidatorWithdrawVault *ValidatorWithdrawVaultTransactor) Initialize(opts *bind.TransactOpts, _poolId uint8, _staderConfig common.Address, _validatorId *big.Int) (*types.Transaction, error) {
	return _ValidatorWithdrawVault.contract.Transact(opts, "initialize", _poolId, _staderConfig, _validatorId)
}

// Initialize is a paid mutator transaction binding the contract method 0x06d39109.
//
// Solidity: function initialize(uint8 _poolId, address _staderConfig, uint256 _validatorId) returns()
func (_ValidatorWithdrawVault *ValidatorWithdrawVaultSession) Initialize(_poolId uint8, _staderConfig common.Address, _validatorId *big.Int) (*types.Transaction, error) {
	return _ValidatorWithdrawVault.Contract.Initialize(&_ValidatorWithdrawVault.TransactOpts, _poolId, _staderConfig, _validatorId)
}

// Initialize is a paid mutator transaction binding the contract method 0x06d39109.
//
// Solidity: function initialize(uint8 _poolId, address _staderConfig, uint256 _validatorId) returns()
func (_ValidatorWithdrawVault *ValidatorWithdrawVaultTransactorSession) Initialize(_poolId uint8, _staderConfig common.Address, _validatorId *big.Int) (*types.Transaction, error) {
	return _ValidatorWithdrawVault.Contract.Initialize(&_ValidatorWithdrawVault.TransactOpts, _poolId, _staderConfig, _validatorId)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address account) returns()
func (_ValidatorWithdrawVault *ValidatorWithdrawVaultTransactor) RenounceRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _ValidatorWithdrawVault.contract.Transact(opts, "renounceRole", role, account)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address account) returns()
func (_ValidatorWithdrawVault *ValidatorWithdrawVaultSession) RenounceRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _ValidatorWithdrawVault.Contract.RenounceRole(&_ValidatorWithdrawVault.TransactOpts, role, account)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address account) returns()
func (_ValidatorWithdrawVault *ValidatorWithdrawVaultTransactorSession) RenounceRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _ValidatorWithdrawVault.Contract.RenounceRole(&_ValidatorWithdrawVault.TransactOpts, role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_ValidatorWithdrawVault *ValidatorWithdrawVaultTransactor) RevokeRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _ValidatorWithdrawVault.contract.Transact(opts, "revokeRole", role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_ValidatorWithdrawVault *ValidatorWithdrawVaultSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _ValidatorWithdrawVault.Contract.RevokeRole(&_ValidatorWithdrawVault.TransactOpts, role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_ValidatorWithdrawVault *ValidatorWithdrawVaultTransactorSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _ValidatorWithdrawVault.Contract.RevokeRole(&_ValidatorWithdrawVault.TransactOpts, role, account)
}

// SettleFunds is a paid mutator transaction binding the contract method 0xbf8ac490.
//
// Solidity: function settleFunds() returns()
func (_ValidatorWithdrawVault *ValidatorWithdrawVaultTransactor) SettleFunds(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ValidatorWithdrawVault.contract.Transact(opts, "settleFunds")
}

// SettleFunds is a paid mutator transaction binding the contract method 0xbf8ac490.
//
// Solidity: function settleFunds() returns()
func (_ValidatorWithdrawVault *ValidatorWithdrawVaultSession) SettleFunds() (*types.Transaction, error) {
	return _ValidatorWithdrawVault.Contract.SettleFunds(&_ValidatorWithdrawVault.TransactOpts)
}

// SettleFunds is a paid mutator transaction binding the contract method 0xbf8ac490.
//
// Solidity: function settleFunds() returns()
func (_ValidatorWithdrawVault *ValidatorWithdrawVaultTransactorSession) SettleFunds() (*types.Transaction, error) {
	return _ValidatorWithdrawVault.Contract.SettleFunds(&_ValidatorWithdrawVault.TransactOpts)
}

// UpdateStaderConfig is a paid mutator transaction binding the contract method 0x9ee804cb.
//
// Solidity: function updateStaderConfig(address _staderConfig) returns()
func (_ValidatorWithdrawVault *ValidatorWithdrawVaultTransactor) UpdateStaderConfig(opts *bind.TransactOpts, _staderConfig common.Address) (*types.Transaction, error) {
	return _ValidatorWithdrawVault.contract.Transact(opts, "updateStaderConfig", _staderConfig)
}

// UpdateStaderConfig is a paid mutator transaction binding the contract method 0x9ee804cb.
//
// Solidity: function updateStaderConfig(address _staderConfig) returns()
func (_ValidatorWithdrawVault *ValidatorWithdrawVaultSession) UpdateStaderConfig(_staderConfig common.Address) (*types.Transaction, error) {
	return _ValidatorWithdrawVault.Contract.UpdateStaderConfig(&_ValidatorWithdrawVault.TransactOpts, _staderConfig)
}

// UpdateStaderConfig is a paid mutator transaction binding the contract method 0x9ee804cb.
//
// Solidity: function updateStaderConfig(address _staderConfig) returns()
func (_ValidatorWithdrawVault *ValidatorWithdrawVaultTransactorSession) UpdateStaderConfig(_staderConfig common.Address) (*types.Transaction, error) {
	return _ValidatorWithdrawVault.Contract.UpdateStaderConfig(&_ValidatorWithdrawVault.TransactOpts, _staderConfig)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_ValidatorWithdrawVault *ValidatorWithdrawVaultTransactor) Receive(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ValidatorWithdrawVault.contract.RawTransact(opts, nil) // calldata is disallowed for receive function
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_ValidatorWithdrawVault *ValidatorWithdrawVaultSession) Receive() (*types.Transaction, error) {
	return _ValidatorWithdrawVault.Contract.Receive(&_ValidatorWithdrawVault.TransactOpts)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_ValidatorWithdrawVault *ValidatorWithdrawVaultTransactorSession) Receive() (*types.Transaction, error) {
	return _ValidatorWithdrawVault.Contract.Receive(&_ValidatorWithdrawVault.TransactOpts)
}

// ValidatorWithdrawVaultDistributeRewardFailedIterator is returned from FilterDistributeRewardFailed and is used to iterate over the raw logs and unpacked data for DistributeRewardFailed events raised by the ValidatorWithdrawVault contract.
type ValidatorWithdrawVaultDistributeRewardFailedIterator struct {
	Event *ValidatorWithdrawVaultDistributeRewardFailed // Event containing the contract specifics and raw log

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
func (it *ValidatorWithdrawVaultDistributeRewardFailedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ValidatorWithdrawVaultDistributeRewardFailed)
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
		it.Event = new(ValidatorWithdrawVaultDistributeRewardFailed)
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
func (it *ValidatorWithdrawVaultDistributeRewardFailedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ValidatorWithdrawVaultDistributeRewardFailedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ValidatorWithdrawVaultDistributeRewardFailed represents a DistributeRewardFailed event raised by the ValidatorWithdrawVault contract.
type ValidatorWithdrawVaultDistributeRewardFailed struct {
	RewardAmount    *big.Int
	RewardThreshold *big.Int
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterDistributeRewardFailed is a free log retrieval operation binding the contract event 0x88248e03625d5954a12c62ce113d86a84d6166c264dcc1e3001139c8442a2767.
//
// Solidity: event DistributeRewardFailed(uint256 rewardAmount, uint256 rewardThreshold)
func (_ValidatorWithdrawVault *ValidatorWithdrawVaultFilterer) FilterDistributeRewardFailed(opts *bind.FilterOpts) (*ValidatorWithdrawVaultDistributeRewardFailedIterator, error) {

	logs, sub, err := _ValidatorWithdrawVault.contract.FilterLogs(opts, "DistributeRewardFailed")
	if err != nil {
		return nil, err
	}
	return &ValidatorWithdrawVaultDistributeRewardFailedIterator{contract: _ValidatorWithdrawVault.contract, event: "DistributeRewardFailed", logs: logs, sub: sub}, nil
}

// WatchDistributeRewardFailed is a free log subscription operation binding the contract event 0x88248e03625d5954a12c62ce113d86a84d6166c264dcc1e3001139c8442a2767.
//
// Solidity: event DistributeRewardFailed(uint256 rewardAmount, uint256 rewardThreshold)
func (_ValidatorWithdrawVault *ValidatorWithdrawVaultFilterer) WatchDistributeRewardFailed(opts *bind.WatchOpts, sink chan<- *ValidatorWithdrawVaultDistributeRewardFailed) (event.Subscription, error) {

	logs, sub, err := _ValidatorWithdrawVault.contract.WatchLogs(opts, "DistributeRewardFailed")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ValidatorWithdrawVaultDistributeRewardFailed)
				if err := _ValidatorWithdrawVault.contract.UnpackLog(event, "DistributeRewardFailed", log); err != nil {
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

// ParseDistributeRewardFailed is a log parse operation binding the contract event 0x88248e03625d5954a12c62ce113d86a84d6166c264dcc1e3001139c8442a2767.
//
// Solidity: event DistributeRewardFailed(uint256 rewardAmount, uint256 rewardThreshold)
func (_ValidatorWithdrawVault *ValidatorWithdrawVaultFilterer) ParseDistributeRewardFailed(log types.Log) (*ValidatorWithdrawVaultDistributeRewardFailed, error) {
	event := new(ValidatorWithdrawVaultDistributeRewardFailed)
	if err := _ValidatorWithdrawVault.contract.UnpackLog(event, "DistributeRewardFailed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ValidatorWithdrawVaultDistributedRewardsIterator is returned from FilterDistributedRewards and is used to iterate over the raw logs and unpacked data for DistributedRewards events raised by the ValidatorWithdrawVault contract.
type ValidatorWithdrawVaultDistributedRewardsIterator struct {
	Event *ValidatorWithdrawVaultDistributedRewards // Event containing the contract specifics and raw log

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
func (it *ValidatorWithdrawVaultDistributedRewardsIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ValidatorWithdrawVaultDistributedRewards)
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
		it.Event = new(ValidatorWithdrawVaultDistributedRewards)
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
func (it *ValidatorWithdrawVaultDistributedRewardsIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ValidatorWithdrawVaultDistributedRewardsIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ValidatorWithdrawVaultDistributedRewards represents a DistributedRewards event raised by the ValidatorWithdrawVault contract.
type ValidatorWithdrawVaultDistributedRewards struct {
	UserShare     *big.Int
	OperatorShare *big.Int
	ProtocolShare *big.Int
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterDistributedRewards is a free log retrieval operation binding the contract event 0x95a31bc3041897ca26d2debdc69c41333fbf5d1cb92040b8d0d35e62c5e01433.
//
// Solidity: event DistributedRewards(uint256 userShare, uint256 operatorShare, uint256 protocolShare)
func (_ValidatorWithdrawVault *ValidatorWithdrawVaultFilterer) FilterDistributedRewards(opts *bind.FilterOpts) (*ValidatorWithdrawVaultDistributedRewardsIterator, error) {

	logs, sub, err := _ValidatorWithdrawVault.contract.FilterLogs(opts, "DistributedRewards")
	if err != nil {
		return nil, err
	}
	return &ValidatorWithdrawVaultDistributedRewardsIterator{contract: _ValidatorWithdrawVault.contract, event: "DistributedRewards", logs: logs, sub: sub}, nil
}

// WatchDistributedRewards is a free log subscription operation binding the contract event 0x95a31bc3041897ca26d2debdc69c41333fbf5d1cb92040b8d0d35e62c5e01433.
//
// Solidity: event DistributedRewards(uint256 userShare, uint256 operatorShare, uint256 protocolShare)
func (_ValidatorWithdrawVault *ValidatorWithdrawVaultFilterer) WatchDistributedRewards(opts *bind.WatchOpts, sink chan<- *ValidatorWithdrawVaultDistributedRewards) (event.Subscription, error) {

	logs, sub, err := _ValidatorWithdrawVault.contract.WatchLogs(opts, "DistributedRewards")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ValidatorWithdrawVaultDistributedRewards)
				if err := _ValidatorWithdrawVault.contract.UnpackLog(event, "DistributedRewards", log); err != nil {
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

// ParseDistributedRewards is a log parse operation binding the contract event 0x95a31bc3041897ca26d2debdc69c41333fbf5d1cb92040b8d0d35e62c5e01433.
//
// Solidity: event DistributedRewards(uint256 userShare, uint256 operatorShare, uint256 protocolShare)
func (_ValidatorWithdrawVault *ValidatorWithdrawVaultFilterer) ParseDistributedRewards(log types.Log) (*ValidatorWithdrawVaultDistributedRewards, error) {
	event := new(ValidatorWithdrawVaultDistributedRewards)
	if err := _ValidatorWithdrawVault.contract.UnpackLog(event, "DistributedRewards", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ValidatorWithdrawVaultETHReceivedIterator is returned from FilterETHReceived and is used to iterate over the raw logs and unpacked data for ETHReceived events raised by the ValidatorWithdrawVault contract.
type ValidatorWithdrawVaultETHReceivedIterator struct {
	Event *ValidatorWithdrawVaultETHReceived // Event containing the contract specifics and raw log

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
func (it *ValidatorWithdrawVaultETHReceivedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ValidatorWithdrawVaultETHReceived)
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
		it.Event = new(ValidatorWithdrawVaultETHReceived)
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
func (it *ValidatorWithdrawVaultETHReceivedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ValidatorWithdrawVaultETHReceivedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ValidatorWithdrawVaultETHReceived represents a ETHReceived event raised by the ValidatorWithdrawVault contract.
type ValidatorWithdrawVaultETHReceived struct {
	Sender common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterETHReceived is a free log retrieval operation binding the contract event 0xbfe611b001dfcd411432f7bf0d79b82b4b2ee81511edac123a3403c357fb972a.
//
// Solidity: event ETHReceived(address indexed sender, uint256 amount)
func (_ValidatorWithdrawVault *ValidatorWithdrawVaultFilterer) FilterETHReceived(opts *bind.FilterOpts, sender []common.Address) (*ValidatorWithdrawVaultETHReceivedIterator, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _ValidatorWithdrawVault.contract.FilterLogs(opts, "ETHReceived", senderRule)
	if err != nil {
		return nil, err
	}
	return &ValidatorWithdrawVaultETHReceivedIterator{contract: _ValidatorWithdrawVault.contract, event: "ETHReceived", logs: logs, sub: sub}, nil
}

// WatchETHReceived is a free log subscription operation binding the contract event 0xbfe611b001dfcd411432f7bf0d79b82b4b2ee81511edac123a3403c357fb972a.
//
// Solidity: event ETHReceived(address indexed sender, uint256 amount)
func (_ValidatorWithdrawVault *ValidatorWithdrawVaultFilterer) WatchETHReceived(opts *bind.WatchOpts, sink chan<- *ValidatorWithdrawVaultETHReceived, sender []common.Address) (event.Subscription, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _ValidatorWithdrawVault.contract.WatchLogs(opts, "ETHReceived", senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ValidatorWithdrawVaultETHReceived)
				if err := _ValidatorWithdrawVault.contract.UnpackLog(event, "ETHReceived", log); err != nil {
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
func (_ValidatorWithdrawVault *ValidatorWithdrawVaultFilterer) ParseETHReceived(log types.Log) (*ValidatorWithdrawVaultETHReceived, error) {
	event := new(ValidatorWithdrawVaultETHReceived)
	if err := _ValidatorWithdrawVault.contract.UnpackLog(event, "ETHReceived", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ValidatorWithdrawVaultInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the ValidatorWithdrawVault contract.
type ValidatorWithdrawVaultInitializedIterator struct {
	Event *ValidatorWithdrawVaultInitialized // Event containing the contract specifics and raw log

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
func (it *ValidatorWithdrawVaultInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ValidatorWithdrawVaultInitialized)
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
		it.Event = new(ValidatorWithdrawVaultInitialized)
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
func (it *ValidatorWithdrawVaultInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ValidatorWithdrawVaultInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ValidatorWithdrawVaultInitialized represents a Initialized event raised by the ValidatorWithdrawVault contract.
type ValidatorWithdrawVaultInitialized struct {
	Version uint8
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_ValidatorWithdrawVault *ValidatorWithdrawVaultFilterer) FilterInitialized(opts *bind.FilterOpts) (*ValidatorWithdrawVaultInitializedIterator, error) {

	logs, sub, err := _ValidatorWithdrawVault.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &ValidatorWithdrawVaultInitializedIterator{contract: _ValidatorWithdrawVault.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_ValidatorWithdrawVault *ValidatorWithdrawVaultFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *ValidatorWithdrawVaultInitialized) (event.Subscription, error) {

	logs, sub, err := _ValidatorWithdrawVault.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ValidatorWithdrawVaultInitialized)
				if err := _ValidatorWithdrawVault.contract.UnpackLog(event, "Initialized", log); err != nil {
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
func (_ValidatorWithdrawVault *ValidatorWithdrawVaultFilterer) ParseInitialized(log types.Log) (*ValidatorWithdrawVaultInitialized, error) {
	event := new(ValidatorWithdrawVaultInitialized)
	if err := _ValidatorWithdrawVault.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ValidatorWithdrawVaultRoleAdminChangedIterator is returned from FilterRoleAdminChanged and is used to iterate over the raw logs and unpacked data for RoleAdminChanged events raised by the ValidatorWithdrawVault contract.
type ValidatorWithdrawVaultRoleAdminChangedIterator struct {
	Event *ValidatorWithdrawVaultRoleAdminChanged // Event containing the contract specifics and raw log

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
func (it *ValidatorWithdrawVaultRoleAdminChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ValidatorWithdrawVaultRoleAdminChanged)
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
		it.Event = new(ValidatorWithdrawVaultRoleAdminChanged)
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
func (it *ValidatorWithdrawVaultRoleAdminChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ValidatorWithdrawVaultRoleAdminChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ValidatorWithdrawVaultRoleAdminChanged represents a RoleAdminChanged event raised by the ValidatorWithdrawVault contract.
type ValidatorWithdrawVaultRoleAdminChanged struct {
	Role              [32]byte
	PreviousAdminRole [32]byte
	NewAdminRole      [32]byte
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterRoleAdminChanged is a free log retrieval operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_ValidatorWithdrawVault *ValidatorWithdrawVaultFilterer) FilterRoleAdminChanged(opts *bind.FilterOpts, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (*ValidatorWithdrawVaultRoleAdminChangedIterator, error) {

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

	logs, sub, err := _ValidatorWithdrawVault.contract.FilterLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return &ValidatorWithdrawVaultRoleAdminChangedIterator{contract: _ValidatorWithdrawVault.contract, event: "RoleAdminChanged", logs: logs, sub: sub}, nil
}

// WatchRoleAdminChanged is a free log subscription operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_ValidatorWithdrawVault *ValidatorWithdrawVaultFilterer) WatchRoleAdminChanged(opts *bind.WatchOpts, sink chan<- *ValidatorWithdrawVaultRoleAdminChanged, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (event.Subscription, error) {

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

	logs, sub, err := _ValidatorWithdrawVault.contract.WatchLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ValidatorWithdrawVaultRoleAdminChanged)
				if err := _ValidatorWithdrawVault.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
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
func (_ValidatorWithdrawVault *ValidatorWithdrawVaultFilterer) ParseRoleAdminChanged(log types.Log) (*ValidatorWithdrawVaultRoleAdminChanged, error) {
	event := new(ValidatorWithdrawVaultRoleAdminChanged)
	if err := _ValidatorWithdrawVault.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ValidatorWithdrawVaultRoleGrantedIterator is returned from FilterRoleGranted and is used to iterate over the raw logs and unpacked data for RoleGranted events raised by the ValidatorWithdrawVault contract.
type ValidatorWithdrawVaultRoleGrantedIterator struct {
	Event *ValidatorWithdrawVaultRoleGranted // Event containing the contract specifics and raw log

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
func (it *ValidatorWithdrawVaultRoleGrantedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ValidatorWithdrawVaultRoleGranted)
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
		it.Event = new(ValidatorWithdrawVaultRoleGranted)
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
func (it *ValidatorWithdrawVaultRoleGrantedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ValidatorWithdrawVaultRoleGrantedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ValidatorWithdrawVaultRoleGranted represents a RoleGranted event raised by the ValidatorWithdrawVault contract.
type ValidatorWithdrawVaultRoleGranted struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleGranted is a free log retrieval operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_ValidatorWithdrawVault *ValidatorWithdrawVaultFilterer) FilterRoleGranted(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*ValidatorWithdrawVaultRoleGrantedIterator, error) {

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

	logs, sub, err := _ValidatorWithdrawVault.contract.FilterLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &ValidatorWithdrawVaultRoleGrantedIterator{contract: _ValidatorWithdrawVault.contract, event: "RoleGranted", logs: logs, sub: sub}, nil
}

// WatchRoleGranted is a free log subscription operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_ValidatorWithdrawVault *ValidatorWithdrawVaultFilterer) WatchRoleGranted(opts *bind.WatchOpts, sink chan<- *ValidatorWithdrawVaultRoleGranted, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _ValidatorWithdrawVault.contract.WatchLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ValidatorWithdrawVaultRoleGranted)
				if err := _ValidatorWithdrawVault.contract.UnpackLog(event, "RoleGranted", log); err != nil {
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
func (_ValidatorWithdrawVault *ValidatorWithdrawVaultFilterer) ParseRoleGranted(log types.Log) (*ValidatorWithdrawVaultRoleGranted, error) {
	event := new(ValidatorWithdrawVaultRoleGranted)
	if err := _ValidatorWithdrawVault.contract.UnpackLog(event, "RoleGranted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ValidatorWithdrawVaultRoleRevokedIterator is returned from FilterRoleRevoked and is used to iterate over the raw logs and unpacked data for RoleRevoked events raised by the ValidatorWithdrawVault contract.
type ValidatorWithdrawVaultRoleRevokedIterator struct {
	Event *ValidatorWithdrawVaultRoleRevoked // Event containing the contract specifics and raw log

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
func (it *ValidatorWithdrawVaultRoleRevokedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ValidatorWithdrawVaultRoleRevoked)
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
		it.Event = new(ValidatorWithdrawVaultRoleRevoked)
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
func (it *ValidatorWithdrawVaultRoleRevokedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ValidatorWithdrawVaultRoleRevokedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ValidatorWithdrawVaultRoleRevoked represents a RoleRevoked event raised by the ValidatorWithdrawVault contract.
type ValidatorWithdrawVaultRoleRevoked struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleRevoked is a free log retrieval operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_ValidatorWithdrawVault *ValidatorWithdrawVaultFilterer) FilterRoleRevoked(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*ValidatorWithdrawVaultRoleRevokedIterator, error) {

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

	logs, sub, err := _ValidatorWithdrawVault.contract.FilterLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &ValidatorWithdrawVaultRoleRevokedIterator{contract: _ValidatorWithdrawVault.contract, event: "RoleRevoked", logs: logs, sub: sub}, nil
}

// WatchRoleRevoked is a free log subscription operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_ValidatorWithdrawVault *ValidatorWithdrawVaultFilterer) WatchRoleRevoked(opts *bind.WatchOpts, sink chan<- *ValidatorWithdrawVaultRoleRevoked, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _ValidatorWithdrawVault.contract.WatchLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ValidatorWithdrawVaultRoleRevoked)
				if err := _ValidatorWithdrawVault.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
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
func (_ValidatorWithdrawVault *ValidatorWithdrawVaultFilterer) ParseRoleRevoked(log types.Log) (*ValidatorWithdrawVaultRoleRevoked, error) {
	event := new(ValidatorWithdrawVaultRoleRevoked)
	if err := _ValidatorWithdrawVault.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ValidatorWithdrawVaultSettledFundsIterator is returned from FilterSettledFunds and is used to iterate over the raw logs and unpacked data for SettledFunds events raised by the ValidatorWithdrawVault contract.
type ValidatorWithdrawVaultSettledFundsIterator struct {
	Event *ValidatorWithdrawVaultSettledFunds // Event containing the contract specifics and raw log

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
func (it *ValidatorWithdrawVaultSettledFundsIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ValidatorWithdrawVaultSettledFunds)
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
		it.Event = new(ValidatorWithdrawVaultSettledFunds)
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
func (it *ValidatorWithdrawVaultSettledFundsIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ValidatorWithdrawVaultSettledFundsIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ValidatorWithdrawVaultSettledFunds represents a SettledFunds event raised by the ValidatorWithdrawVault contract.
type ValidatorWithdrawVaultSettledFunds struct {
	UserShare     *big.Int
	OperatorShare *big.Int
	ProtocolShare *big.Int
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterSettledFunds is a free log retrieval operation binding the contract event 0xddbef61b53bcb1654a5ef3e9d8a4e087b0db8a7812f4e7949203ecd09256260e.
//
// Solidity: event SettledFunds(uint256 userShare, uint256 operatorShare, uint256 protocolShare)
func (_ValidatorWithdrawVault *ValidatorWithdrawVaultFilterer) FilterSettledFunds(opts *bind.FilterOpts) (*ValidatorWithdrawVaultSettledFundsIterator, error) {

	logs, sub, err := _ValidatorWithdrawVault.contract.FilterLogs(opts, "SettledFunds")
	if err != nil {
		return nil, err
	}
	return &ValidatorWithdrawVaultSettledFundsIterator{contract: _ValidatorWithdrawVault.contract, event: "SettledFunds", logs: logs, sub: sub}, nil
}

// WatchSettledFunds is a free log subscription operation binding the contract event 0xddbef61b53bcb1654a5ef3e9d8a4e087b0db8a7812f4e7949203ecd09256260e.
//
// Solidity: event SettledFunds(uint256 userShare, uint256 operatorShare, uint256 protocolShare)
func (_ValidatorWithdrawVault *ValidatorWithdrawVaultFilterer) WatchSettledFunds(opts *bind.WatchOpts, sink chan<- *ValidatorWithdrawVaultSettledFunds) (event.Subscription, error) {

	logs, sub, err := _ValidatorWithdrawVault.contract.WatchLogs(opts, "SettledFunds")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ValidatorWithdrawVaultSettledFunds)
				if err := _ValidatorWithdrawVault.contract.UnpackLog(event, "SettledFunds", log); err != nil {
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

// ParseSettledFunds is a log parse operation binding the contract event 0xddbef61b53bcb1654a5ef3e9d8a4e087b0db8a7812f4e7949203ecd09256260e.
//
// Solidity: event SettledFunds(uint256 userShare, uint256 operatorShare, uint256 protocolShare)
func (_ValidatorWithdrawVault *ValidatorWithdrawVaultFilterer) ParseSettledFunds(log types.Log) (*ValidatorWithdrawVaultSettledFunds, error) {
	event := new(ValidatorWithdrawVaultSettledFunds)
	if err := _ValidatorWithdrawVault.contract.UnpackLog(event, "SettledFunds", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ValidatorWithdrawVaultUpdatedStaderConfigIterator is returned from FilterUpdatedStaderConfig and is used to iterate over the raw logs and unpacked data for UpdatedStaderConfig events raised by the ValidatorWithdrawVault contract.
type ValidatorWithdrawVaultUpdatedStaderConfigIterator struct {
	Event *ValidatorWithdrawVaultUpdatedStaderConfig // Event containing the contract specifics and raw log

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
func (it *ValidatorWithdrawVaultUpdatedStaderConfigIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ValidatorWithdrawVaultUpdatedStaderConfig)
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
		it.Event = new(ValidatorWithdrawVaultUpdatedStaderConfig)
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
func (it *ValidatorWithdrawVaultUpdatedStaderConfigIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ValidatorWithdrawVaultUpdatedStaderConfigIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ValidatorWithdrawVaultUpdatedStaderConfig represents a UpdatedStaderConfig event raised by the ValidatorWithdrawVault contract.
type ValidatorWithdrawVaultUpdatedStaderConfig struct {
	StaderConfig common.Address
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterUpdatedStaderConfig is a free log retrieval operation binding the contract event 0xdb2219043d7b197cb235f1af0cf6d782d77dee3de19e3f4fb6d39aae633b4485.
//
// Solidity: event UpdatedStaderConfig(address _staderConfig)
func (_ValidatorWithdrawVault *ValidatorWithdrawVaultFilterer) FilterUpdatedStaderConfig(opts *bind.FilterOpts) (*ValidatorWithdrawVaultUpdatedStaderConfigIterator, error) {

	logs, sub, err := _ValidatorWithdrawVault.contract.FilterLogs(opts, "UpdatedStaderConfig")
	if err != nil {
		return nil, err
	}
	return &ValidatorWithdrawVaultUpdatedStaderConfigIterator{contract: _ValidatorWithdrawVault.contract, event: "UpdatedStaderConfig", logs: logs, sub: sub}, nil
}

// WatchUpdatedStaderConfig is a free log subscription operation binding the contract event 0xdb2219043d7b197cb235f1af0cf6d782d77dee3de19e3f4fb6d39aae633b4485.
//
// Solidity: event UpdatedStaderConfig(address _staderConfig)
func (_ValidatorWithdrawVault *ValidatorWithdrawVaultFilterer) WatchUpdatedStaderConfig(opts *bind.WatchOpts, sink chan<- *ValidatorWithdrawVaultUpdatedStaderConfig) (event.Subscription, error) {

	logs, sub, err := _ValidatorWithdrawVault.contract.WatchLogs(opts, "UpdatedStaderConfig")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ValidatorWithdrawVaultUpdatedStaderConfig)
				if err := _ValidatorWithdrawVault.contract.UnpackLog(event, "UpdatedStaderConfig", log); err != nil {
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
// Solidity: event UpdatedStaderConfig(address _staderConfig)
func (_ValidatorWithdrawVault *ValidatorWithdrawVaultFilterer) ParseUpdatedStaderConfig(log types.Log) (*ValidatorWithdrawVaultUpdatedStaderConfig, error) {
	event := new(ValidatorWithdrawVaultUpdatedStaderConfig)
	if err := _ValidatorWithdrawVault.contract.UnpackLog(event, "UpdatedStaderConfig", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
