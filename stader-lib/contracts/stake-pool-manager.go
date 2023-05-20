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

// StakePoolManagerMetaData contains all meta data concerning the StakePoolManager contract.
var StakePoolManagerMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"CallerNotManager\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"CallerNotStaderContract\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"CooldownNotComplete\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InsufficientBalance\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidDepositAmount\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"PoolIdDoesNotExit\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"TransferFailed\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"UnsupportedOperation\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"UnsupportedOperationInSafeMode\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ZeroAddress\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"AuctionedEthReceived\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"caller\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"assets\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"shares\",\"type\":\"uint256\"}],\"name\":\"Deposited\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"poolId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"poolAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"validatorCount\",\"type\":\"uint256\"}],\"name\":\"ETHTransferredToPool\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"ExecutionLayerRewardsReceived\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"version\",\"type\":\"uint8\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Paused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint8\",\"name\":\"poolId\",\"type\":\"uint8\"}],\"name\":\"ReceivedExcessEthFromPool\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"previousAdminRole\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"newAdminRole\",\"type\":\"bytes32\"}],\"name\":\"RoleAdminChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleGranted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleRevoked\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"TransferredETHToUserWithdrawManager\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Unpaused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"excessETHDepositCoolDown\",\"type\":\"uint256\"}],\"name\":\"UpdatedExcessETHDepositCoolDown\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"staderConfig\",\"type\":\"address\"}],\"name\":\"UpdatedStaderConfig\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"WithdrawVaultUserShareReceived\",\"type\":\"event\"},{\"stateMutability\":\"payable\",\"type\":\"fallback\"},{\"inputs\":[],\"name\":\"DEFAULT_ADMIN_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_shares\",\"type\":\"uint256\"}],\"name\":\"convertToAssets\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_assets\",\"type\":\"uint256\"}],\"name\":\"convertToShares\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_receiver\",\"type\":\"address\"}],\"name\":\"deposit\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"depositETHOverTargetWeight\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"excessETHDepositCoolDown\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getExchangeRate\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"}],\"name\":\"getRoleAdmin\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"grantRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"hasRole\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_admin\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_staderConfig\",\"type\":\"address\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"isVaultHealthy\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"lastExcessETHDepositBlock\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"maxDeposit\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"minDeposit\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"paused\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_assets\",\"type\":\"uint256\"}],\"name\":\"previewDeposit\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_shares\",\"type\":\"uint256\"}],\"name\":\"previewWithdraw\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"receiveEthFromAuction\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"_poolId\",\"type\":\"uint8\"}],\"name\":\"receiveExcessEthFromPool\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"receiveExecutionLayerRewards\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"receiveWithdrawVaultUserShare\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"renounceRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"revokeRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"staderConfig\",\"outputs\":[{\"internalType\":\"contractIStaderConfig\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalAssets\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"transferETHToUserWithdrawManager\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"unpause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_excessETHDepositCoolDown\",\"type\":\"uint256\"}],\"name\":\"updateExcessETHDepositCoolDown\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_staderConfig\",\"type\":\"address\"}],\"name\":\"updateStaderConfig\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"_poolId\",\"type\":\"uint8\"}],\"name\":\"validatorBatchDeposit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"stateMutability\":\"payable\",\"type\":\"receive\"}]",
}

// StakePoolManagerABI is the input ABI used to generate the binding from.
// Deprecated: Use StakePoolManagerMetaData.ABI instead.
var StakePoolManagerABI = StakePoolManagerMetaData.ABI

// StakePoolManager is an auto generated Go binding around an Ethereum contract.
type StakePoolManager struct {
	StakePoolManagerCaller     // Read-only binding to the contract
	StakePoolManagerTransactor // Write-only binding to the contract
	StakePoolManagerFilterer   // Log filterer for contract events
}

// StakePoolManagerCaller is an auto generated read-only Go binding around an Ethereum contract.
type StakePoolManagerCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StakePoolManagerTransactor is an auto generated write-only Go binding around an Ethereum contract.
type StakePoolManagerTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StakePoolManagerFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type StakePoolManagerFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StakePoolManagerSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type StakePoolManagerSession struct {
	Contract     *StakePoolManager // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// StakePoolManagerCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type StakePoolManagerCallerSession struct {
	Contract *StakePoolManagerCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts           // Call options to use throughout this session
}

// StakePoolManagerTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type StakePoolManagerTransactorSession struct {
	Contract     *StakePoolManagerTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts           // Transaction auth options to use throughout this session
}

// StakePoolManagerRaw is an auto generated low-level Go binding around an Ethereum contract.
type StakePoolManagerRaw struct {
	Contract *StakePoolManager // Generic contract binding to access the raw methods on
}

// StakePoolManagerCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type StakePoolManagerCallerRaw struct {
	Contract *StakePoolManagerCaller // Generic read-only contract binding to access the raw methods on
}

// StakePoolManagerTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type StakePoolManagerTransactorRaw struct {
	Contract *StakePoolManagerTransactor // Generic write-only contract binding to access the raw methods on
}

// NewStakePoolManager creates a new instance of StakePoolManager, bound to a specific deployed contract.
func NewStakePoolManager(address common.Address, backend bind.ContractBackend) (*StakePoolManager, error) {
	contract, err := bindStakePoolManager(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &StakePoolManager{StakePoolManagerCaller: StakePoolManagerCaller{contract: contract}, StakePoolManagerTransactor: StakePoolManagerTransactor{contract: contract}, StakePoolManagerFilterer: StakePoolManagerFilterer{contract: contract}}, nil
}

// NewStakePoolManagerCaller creates a new read-only instance of StakePoolManager, bound to a specific deployed contract.
func NewStakePoolManagerCaller(address common.Address, caller bind.ContractCaller) (*StakePoolManagerCaller, error) {
	contract, err := bindStakePoolManager(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &StakePoolManagerCaller{contract: contract}, nil
}

// NewStakePoolManagerTransactor creates a new write-only instance of StakePoolManager, bound to a specific deployed contract.
func NewStakePoolManagerTransactor(address common.Address, transactor bind.ContractTransactor) (*StakePoolManagerTransactor, error) {
	contract, err := bindStakePoolManager(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &StakePoolManagerTransactor{contract: contract}, nil
}

// NewStakePoolManagerFilterer creates a new log filterer instance of StakePoolManager, bound to a specific deployed contract.
func NewStakePoolManagerFilterer(address common.Address, filterer bind.ContractFilterer) (*StakePoolManagerFilterer, error) {
	contract, err := bindStakePoolManager(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &StakePoolManagerFilterer{contract: contract}, nil
}

// bindStakePoolManager binds a generic wrapper to an already deployed contract.
func bindStakePoolManager(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := StakePoolManagerMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_StakePoolManager *StakePoolManagerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _StakePoolManager.Contract.StakePoolManagerCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_StakePoolManager *StakePoolManagerRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _StakePoolManager.Contract.StakePoolManagerTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_StakePoolManager *StakePoolManagerRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _StakePoolManager.Contract.StakePoolManagerTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_StakePoolManager *StakePoolManagerCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _StakePoolManager.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_StakePoolManager *StakePoolManagerTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _StakePoolManager.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_StakePoolManager *StakePoolManagerTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _StakePoolManager.Contract.contract.Transact(opts, method, params...)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_StakePoolManager *StakePoolManagerCaller) DEFAULTADMINROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _StakePoolManager.contract.Call(opts, &out, "DEFAULT_ADMIN_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_StakePoolManager *StakePoolManagerSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _StakePoolManager.Contract.DEFAULTADMINROLE(&_StakePoolManager.CallOpts)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_StakePoolManager *StakePoolManagerCallerSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _StakePoolManager.Contract.DEFAULTADMINROLE(&_StakePoolManager.CallOpts)
}

// ConvertToAssets is a free data retrieval call binding the contract method 0x07a2d13a.
//
// Solidity: function convertToAssets(uint256 _shares) view returns(uint256)
func (_StakePoolManager *StakePoolManagerCaller) ConvertToAssets(opts *bind.CallOpts, _shares *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _StakePoolManager.contract.Call(opts, &out, "convertToAssets", _shares)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ConvertToAssets is a free data retrieval call binding the contract method 0x07a2d13a.
//
// Solidity: function convertToAssets(uint256 _shares) view returns(uint256)
func (_StakePoolManager *StakePoolManagerSession) ConvertToAssets(_shares *big.Int) (*big.Int, error) {
	return _StakePoolManager.Contract.ConvertToAssets(&_StakePoolManager.CallOpts, _shares)
}

// ConvertToAssets is a free data retrieval call binding the contract method 0x07a2d13a.
//
// Solidity: function convertToAssets(uint256 _shares) view returns(uint256)
func (_StakePoolManager *StakePoolManagerCallerSession) ConvertToAssets(_shares *big.Int) (*big.Int, error) {
	return _StakePoolManager.Contract.ConvertToAssets(&_StakePoolManager.CallOpts, _shares)
}

// ConvertToShares is a free data retrieval call binding the contract method 0xc6e6f592.
//
// Solidity: function convertToShares(uint256 _assets) view returns(uint256)
func (_StakePoolManager *StakePoolManagerCaller) ConvertToShares(opts *bind.CallOpts, _assets *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _StakePoolManager.contract.Call(opts, &out, "convertToShares", _assets)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ConvertToShares is a free data retrieval call binding the contract method 0xc6e6f592.
//
// Solidity: function convertToShares(uint256 _assets) view returns(uint256)
func (_StakePoolManager *StakePoolManagerSession) ConvertToShares(_assets *big.Int) (*big.Int, error) {
	return _StakePoolManager.Contract.ConvertToShares(&_StakePoolManager.CallOpts, _assets)
}

// ConvertToShares is a free data retrieval call binding the contract method 0xc6e6f592.
//
// Solidity: function convertToShares(uint256 _assets) view returns(uint256)
func (_StakePoolManager *StakePoolManagerCallerSession) ConvertToShares(_assets *big.Int) (*big.Int, error) {
	return _StakePoolManager.Contract.ConvertToShares(&_StakePoolManager.CallOpts, _assets)
}

// ExcessETHDepositCoolDown is a free data retrieval call binding the contract method 0xfa43245f.
//
// Solidity: function excessETHDepositCoolDown() view returns(uint256)
func (_StakePoolManager *StakePoolManagerCaller) ExcessETHDepositCoolDown(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _StakePoolManager.contract.Call(opts, &out, "excessETHDepositCoolDown")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ExcessETHDepositCoolDown is a free data retrieval call binding the contract method 0xfa43245f.
//
// Solidity: function excessETHDepositCoolDown() view returns(uint256)
func (_StakePoolManager *StakePoolManagerSession) ExcessETHDepositCoolDown() (*big.Int, error) {
	return _StakePoolManager.Contract.ExcessETHDepositCoolDown(&_StakePoolManager.CallOpts)
}

// ExcessETHDepositCoolDown is a free data retrieval call binding the contract method 0xfa43245f.
//
// Solidity: function excessETHDepositCoolDown() view returns(uint256)
func (_StakePoolManager *StakePoolManagerCallerSession) ExcessETHDepositCoolDown() (*big.Int, error) {
	return _StakePoolManager.Contract.ExcessETHDepositCoolDown(&_StakePoolManager.CallOpts)
}

// GetExchangeRate is a free data retrieval call binding the contract method 0xe6aa216c.
//
// Solidity: function getExchangeRate() view returns(uint256)
func (_StakePoolManager *StakePoolManagerCaller) GetExchangeRate(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _StakePoolManager.contract.Call(opts, &out, "getExchangeRate")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetExchangeRate is a free data retrieval call binding the contract method 0xe6aa216c.
//
// Solidity: function getExchangeRate() view returns(uint256)
func (_StakePoolManager *StakePoolManagerSession) GetExchangeRate() (*big.Int, error) {
	return _StakePoolManager.Contract.GetExchangeRate(&_StakePoolManager.CallOpts)
}

// GetExchangeRate is a free data retrieval call binding the contract method 0xe6aa216c.
//
// Solidity: function getExchangeRate() view returns(uint256)
func (_StakePoolManager *StakePoolManagerCallerSession) GetExchangeRate() (*big.Int, error) {
	return _StakePoolManager.Contract.GetExchangeRate(&_StakePoolManager.CallOpts)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_StakePoolManager *StakePoolManagerCaller) GetRoleAdmin(opts *bind.CallOpts, role [32]byte) ([32]byte, error) {
	var out []interface{}
	err := _StakePoolManager.contract.Call(opts, &out, "getRoleAdmin", role)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_StakePoolManager *StakePoolManagerSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _StakePoolManager.Contract.GetRoleAdmin(&_StakePoolManager.CallOpts, role)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_StakePoolManager *StakePoolManagerCallerSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _StakePoolManager.Contract.GetRoleAdmin(&_StakePoolManager.CallOpts, role)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_StakePoolManager *StakePoolManagerCaller) HasRole(opts *bind.CallOpts, role [32]byte, account common.Address) (bool, error) {
	var out []interface{}
	err := _StakePoolManager.contract.Call(opts, &out, "hasRole", role, account)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_StakePoolManager *StakePoolManagerSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _StakePoolManager.Contract.HasRole(&_StakePoolManager.CallOpts, role, account)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_StakePoolManager *StakePoolManagerCallerSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _StakePoolManager.Contract.HasRole(&_StakePoolManager.CallOpts, role, account)
}

// IsVaultHealthy is a free data retrieval call binding the contract method 0xd5c9cfb0.
//
// Solidity: function isVaultHealthy() view returns(bool)
func (_StakePoolManager *StakePoolManagerCaller) IsVaultHealthy(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _StakePoolManager.contract.Call(opts, &out, "isVaultHealthy")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsVaultHealthy is a free data retrieval call binding the contract method 0xd5c9cfb0.
//
// Solidity: function isVaultHealthy() view returns(bool)
func (_StakePoolManager *StakePoolManagerSession) IsVaultHealthy() (bool, error) {
	return _StakePoolManager.Contract.IsVaultHealthy(&_StakePoolManager.CallOpts)
}

// IsVaultHealthy is a free data retrieval call binding the contract method 0xd5c9cfb0.
//
// Solidity: function isVaultHealthy() view returns(bool)
func (_StakePoolManager *StakePoolManagerCallerSession) IsVaultHealthy() (bool, error) {
	return _StakePoolManager.Contract.IsVaultHealthy(&_StakePoolManager.CallOpts)
}

// LastExcessETHDepositBlock is a free data retrieval call binding the contract method 0x83770c74.
//
// Solidity: function lastExcessETHDepositBlock() view returns(uint256)
func (_StakePoolManager *StakePoolManagerCaller) LastExcessETHDepositBlock(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _StakePoolManager.contract.Call(opts, &out, "lastExcessETHDepositBlock")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// LastExcessETHDepositBlock is a free data retrieval call binding the contract method 0x83770c74.
//
// Solidity: function lastExcessETHDepositBlock() view returns(uint256)
func (_StakePoolManager *StakePoolManagerSession) LastExcessETHDepositBlock() (*big.Int, error) {
	return _StakePoolManager.Contract.LastExcessETHDepositBlock(&_StakePoolManager.CallOpts)
}

// LastExcessETHDepositBlock is a free data retrieval call binding the contract method 0x83770c74.
//
// Solidity: function lastExcessETHDepositBlock() view returns(uint256)
func (_StakePoolManager *StakePoolManagerCallerSession) LastExcessETHDepositBlock() (*big.Int, error) {
	return _StakePoolManager.Contract.LastExcessETHDepositBlock(&_StakePoolManager.CallOpts)
}

// MaxDeposit is a free data retrieval call binding the contract method 0x6083e59a.
//
// Solidity: function maxDeposit() view returns(uint256)
func (_StakePoolManager *StakePoolManagerCaller) MaxDeposit(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _StakePoolManager.contract.Call(opts, &out, "maxDeposit")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MaxDeposit is a free data retrieval call binding the contract method 0x6083e59a.
//
// Solidity: function maxDeposit() view returns(uint256)
func (_StakePoolManager *StakePoolManagerSession) MaxDeposit() (*big.Int, error) {
	return _StakePoolManager.Contract.MaxDeposit(&_StakePoolManager.CallOpts)
}

// MaxDeposit is a free data retrieval call binding the contract method 0x6083e59a.
//
// Solidity: function maxDeposit() view returns(uint256)
func (_StakePoolManager *StakePoolManagerCallerSession) MaxDeposit() (*big.Int, error) {
	return _StakePoolManager.Contract.MaxDeposit(&_StakePoolManager.CallOpts)
}

// MinDeposit is a free data retrieval call binding the contract method 0x41b3d185.
//
// Solidity: function minDeposit() view returns(uint256)
func (_StakePoolManager *StakePoolManagerCaller) MinDeposit(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _StakePoolManager.contract.Call(opts, &out, "minDeposit")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MinDeposit is a free data retrieval call binding the contract method 0x41b3d185.
//
// Solidity: function minDeposit() view returns(uint256)
func (_StakePoolManager *StakePoolManagerSession) MinDeposit() (*big.Int, error) {
	return _StakePoolManager.Contract.MinDeposit(&_StakePoolManager.CallOpts)
}

// MinDeposit is a free data retrieval call binding the contract method 0x41b3d185.
//
// Solidity: function minDeposit() view returns(uint256)
func (_StakePoolManager *StakePoolManagerCallerSession) MinDeposit() (*big.Int, error) {
	return _StakePoolManager.Contract.MinDeposit(&_StakePoolManager.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_StakePoolManager *StakePoolManagerCaller) Paused(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _StakePoolManager.contract.Call(opts, &out, "paused")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_StakePoolManager *StakePoolManagerSession) Paused() (bool, error) {
	return _StakePoolManager.Contract.Paused(&_StakePoolManager.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_StakePoolManager *StakePoolManagerCallerSession) Paused() (bool, error) {
	return _StakePoolManager.Contract.Paused(&_StakePoolManager.CallOpts)
}

// PreviewDeposit is a free data retrieval call binding the contract method 0xef8b30f7.
//
// Solidity: function previewDeposit(uint256 _assets) view returns(uint256)
func (_StakePoolManager *StakePoolManagerCaller) PreviewDeposit(opts *bind.CallOpts, _assets *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _StakePoolManager.contract.Call(opts, &out, "previewDeposit", _assets)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PreviewDeposit is a free data retrieval call binding the contract method 0xef8b30f7.
//
// Solidity: function previewDeposit(uint256 _assets) view returns(uint256)
func (_StakePoolManager *StakePoolManagerSession) PreviewDeposit(_assets *big.Int) (*big.Int, error) {
	return _StakePoolManager.Contract.PreviewDeposit(&_StakePoolManager.CallOpts, _assets)
}

// PreviewDeposit is a free data retrieval call binding the contract method 0xef8b30f7.
//
// Solidity: function previewDeposit(uint256 _assets) view returns(uint256)
func (_StakePoolManager *StakePoolManagerCallerSession) PreviewDeposit(_assets *big.Int) (*big.Int, error) {
	return _StakePoolManager.Contract.PreviewDeposit(&_StakePoolManager.CallOpts, _assets)
}

// PreviewWithdraw is a free data retrieval call binding the contract method 0x0a28a477.
//
// Solidity: function previewWithdraw(uint256 _shares) view returns(uint256)
func (_StakePoolManager *StakePoolManagerCaller) PreviewWithdraw(opts *bind.CallOpts, _shares *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _StakePoolManager.contract.Call(opts, &out, "previewWithdraw", _shares)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PreviewWithdraw is a free data retrieval call binding the contract method 0x0a28a477.
//
// Solidity: function previewWithdraw(uint256 _shares) view returns(uint256)
func (_StakePoolManager *StakePoolManagerSession) PreviewWithdraw(_shares *big.Int) (*big.Int, error) {
	return _StakePoolManager.Contract.PreviewWithdraw(&_StakePoolManager.CallOpts, _shares)
}

// PreviewWithdraw is a free data retrieval call binding the contract method 0x0a28a477.
//
// Solidity: function previewWithdraw(uint256 _shares) view returns(uint256)
func (_StakePoolManager *StakePoolManagerCallerSession) PreviewWithdraw(_shares *big.Int) (*big.Int, error) {
	return _StakePoolManager.Contract.PreviewWithdraw(&_StakePoolManager.CallOpts, _shares)
}

// StaderConfig is a free data retrieval call binding the contract method 0x490ffa35.
//
// Solidity: function staderConfig() view returns(address)
func (_StakePoolManager *StakePoolManagerCaller) StaderConfig(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _StakePoolManager.contract.Call(opts, &out, "staderConfig")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// StaderConfig is a free data retrieval call binding the contract method 0x490ffa35.
//
// Solidity: function staderConfig() view returns(address)
func (_StakePoolManager *StakePoolManagerSession) StaderConfig() (common.Address, error) {
	return _StakePoolManager.Contract.StaderConfig(&_StakePoolManager.CallOpts)
}

// StaderConfig is a free data retrieval call binding the contract method 0x490ffa35.
//
// Solidity: function staderConfig() view returns(address)
func (_StakePoolManager *StakePoolManagerCallerSession) StaderConfig() (common.Address, error) {
	return _StakePoolManager.Contract.StaderConfig(&_StakePoolManager.CallOpts)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_StakePoolManager *StakePoolManagerCaller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _StakePoolManager.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_StakePoolManager *StakePoolManagerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _StakePoolManager.Contract.SupportsInterface(&_StakePoolManager.CallOpts, interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_StakePoolManager *StakePoolManagerCallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _StakePoolManager.Contract.SupportsInterface(&_StakePoolManager.CallOpts, interfaceId)
}

// TotalAssets is a free data retrieval call binding the contract method 0x01e1d114.
//
// Solidity: function totalAssets() view returns(uint256)
func (_StakePoolManager *StakePoolManagerCaller) TotalAssets(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _StakePoolManager.contract.Call(opts, &out, "totalAssets")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalAssets is a free data retrieval call binding the contract method 0x01e1d114.
//
// Solidity: function totalAssets() view returns(uint256)
func (_StakePoolManager *StakePoolManagerSession) TotalAssets() (*big.Int, error) {
	return _StakePoolManager.Contract.TotalAssets(&_StakePoolManager.CallOpts)
}

// TotalAssets is a free data retrieval call binding the contract method 0x01e1d114.
//
// Solidity: function totalAssets() view returns(uint256)
func (_StakePoolManager *StakePoolManagerCallerSession) TotalAssets() (*big.Int, error) {
	return _StakePoolManager.Contract.TotalAssets(&_StakePoolManager.CallOpts)
}

// Deposit is a paid mutator transaction binding the contract method 0xf340fa01.
//
// Solidity: function deposit(address _receiver) payable returns(uint256)
func (_StakePoolManager *StakePoolManagerTransactor) Deposit(opts *bind.TransactOpts, _receiver common.Address) (*types.Transaction, error) {
	return _StakePoolManager.contract.Transact(opts, "deposit", _receiver)
}

// Deposit is a paid mutator transaction binding the contract method 0xf340fa01.
//
// Solidity: function deposit(address _receiver) payable returns(uint256)
func (_StakePoolManager *StakePoolManagerSession) Deposit(_receiver common.Address) (*types.Transaction, error) {
	return _StakePoolManager.Contract.Deposit(&_StakePoolManager.TransactOpts, _receiver)
}

// Deposit is a paid mutator transaction binding the contract method 0xf340fa01.
//
// Solidity: function deposit(address _receiver) payable returns(uint256)
func (_StakePoolManager *StakePoolManagerTransactorSession) Deposit(_receiver common.Address) (*types.Transaction, error) {
	return _StakePoolManager.Contract.Deposit(&_StakePoolManager.TransactOpts, _receiver)
}

// DepositETHOverTargetWeight is a paid mutator transaction binding the contract method 0xbf040ea1.
//
// Solidity: function depositETHOverTargetWeight() returns()
func (_StakePoolManager *StakePoolManagerTransactor) DepositETHOverTargetWeight(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _StakePoolManager.contract.Transact(opts, "depositETHOverTargetWeight")
}

// DepositETHOverTargetWeight is a paid mutator transaction binding the contract method 0xbf040ea1.
//
// Solidity: function depositETHOverTargetWeight() returns()
func (_StakePoolManager *StakePoolManagerSession) DepositETHOverTargetWeight() (*types.Transaction, error) {
	return _StakePoolManager.Contract.DepositETHOverTargetWeight(&_StakePoolManager.TransactOpts)
}

// DepositETHOverTargetWeight is a paid mutator transaction binding the contract method 0xbf040ea1.
//
// Solidity: function depositETHOverTargetWeight() returns()
func (_StakePoolManager *StakePoolManagerTransactorSession) DepositETHOverTargetWeight() (*types.Transaction, error) {
	return _StakePoolManager.Contract.DepositETHOverTargetWeight(&_StakePoolManager.TransactOpts)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_StakePoolManager *StakePoolManagerTransactor) GrantRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _StakePoolManager.contract.Transact(opts, "grantRole", role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_StakePoolManager *StakePoolManagerSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _StakePoolManager.Contract.GrantRole(&_StakePoolManager.TransactOpts, role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_StakePoolManager *StakePoolManagerTransactorSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _StakePoolManager.Contract.GrantRole(&_StakePoolManager.TransactOpts, role, account)
}

// Initialize is a paid mutator transaction binding the contract method 0x485cc955.
//
// Solidity: function initialize(address _admin, address _staderConfig) returns()
func (_StakePoolManager *StakePoolManagerTransactor) Initialize(opts *bind.TransactOpts, _admin common.Address, _staderConfig common.Address) (*types.Transaction, error) {
	return _StakePoolManager.contract.Transact(opts, "initialize", _admin, _staderConfig)
}

// Initialize is a paid mutator transaction binding the contract method 0x485cc955.
//
// Solidity: function initialize(address _admin, address _staderConfig) returns()
func (_StakePoolManager *StakePoolManagerSession) Initialize(_admin common.Address, _staderConfig common.Address) (*types.Transaction, error) {
	return _StakePoolManager.Contract.Initialize(&_StakePoolManager.TransactOpts, _admin, _staderConfig)
}

// Initialize is a paid mutator transaction binding the contract method 0x485cc955.
//
// Solidity: function initialize(address _admin, address _staderConfig) returns()
func (_StakePoolManager *StakePoolManagerTransactorSession) Initialize(_admin common.Address, _staderConfig common.Address) (*types.Transaction, error) {
	return _StakePoolManager.Contract.Initialize(&_StakePoolManager.TransactOpts, _admin, _staderConfig)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_StakePoolManager *StakePoolManagerTransactor) Pause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _StakePoolManager.contract.Transact(opts, "pause")
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_StakePoolManager *StakePoolManagerSession) Pause() (*types.Transaction, error) {
	return _StakePoolManager.Contract.Pause(&_StakePoolManager.TransactOpts)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_StakePoolManager *StakePoolManagerTransactorSession) Pause() (*types.Transaction, error) {
	return _StakePoolManager.Contract.Pause(&_StakePoolManager.TransactOpts)
}

// ReceiveEthFromAuction is a paid mutator transaction binding the contract method 0x6f3ca778.
//
// Solidity: function receiveEthFromAuction() payable returns()
func (_StakePoolManager *StakePoolManagerTransactor) ReceiveEthFromAuction(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _StakePoolManager.contract.Transact(opts, "receiveEthFromAuction")
}

// ReceiveEthFromAuction is a paid mutator transaction binding the contract method 0x6f3ca778.
//
// Solidity: function receiveEthFromAuction() payable returns()
func (_StakePoolManager *StakePoolManagerSession) ReceiveEthFromAuction() (*types.Transaction, error) {
	return _StakePoolManager.Contract.ReceiveEthFromAuction(&_StakePoolManager.TransactOpts)
}

// ReceiveEthFromAuction is a paid mutator transaction binding the contract method 0x6f3ca778.
//
// Solidity: function receiveEthFromAuction() payable returns()
func (_StakePoolManager *StakePoolManagerTransactorSession) ReceiveEthFromAuction() (*types.Transaction, error) {
	return _StakePoolManager.Contract.ReceiveEthFromAuction(&_StakePoolManager.TransactOpts)
}

// ReceiveExcessEthFromPool is a paid mutator transaction binding the contract method 0xa55d3088.
//
// Solidity: function receiveExcessEthFromPool(uint8 _poolId) payable returns()
func (_StakePoolManager *StakePoolManagerTransactor) ReceiveExcessEthFromPool(opts *bind.TransactOpts, _poolId uint8) (*types.Transaction, error) {
	return _StakePoolManager.contract.Transact(opts, "receiveExcessEthFromPool", _poolId)
}

// ReceiveExcessEthFromPool is a paid mutator transaction binding the contract method 0xa55d3088.
//
// Solidity: function receiveExcessEthFromPool(uint8 _poolId) payable returns()
func (_StakePoolManager *StakePoolManagerSession) ReceiveExcessEthFromPool(_poolId uint8) (*types.Transaction, error) {
	return _StakePoolManager.Contract.ReceiveExcessEthFromPool(&_StakePoolManager.TransactOpts, _poolId)
}

// ReceiveExcessEthFromPool is a paid mutator transaction binding the contract method 0xa55d3088.
//
// Solidity: function receiveExcessEthFromPool(uint8 _poolId) payable returns()
func (_StakePoolManager *StakePoolManagerTransactorSession) ReceiveExcessEthFromPool(_poolId uint8) (*types.Transaction, error) {
	return _StakePoolManager.Contract.ReceiveExcessEthFromPool(&_StakePoolManager.TransactOpts, _poolId)
}

// ReceiveExecutionLayerRewards is a paid mutator transaction binding the contract method 0x33cf0ef0.
//
// Solidity: function receiveExecutionLayerRewards() payable returns()
func (_StakePoolManager *StakePoolManagerTransactor) ReceiveExecutionLayerRewards(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _StakePoolManager.contract.Transact(opts, "receiveExecutionLayerRewards")
}

// ReceiveExecutionLayerRewards is a paid mutator transaction binding the contract method 0x33cf0ef0.
//
// Solidity: function receiveExecutionLayerRewards() payable returns()
func (_StakePoolManager *StakePoolManagerSession) ReceiveExecutionLayerRewards() (*types.Transaction, error) {
	return _StakePoolManager.Contract.ReceiveExecutionLayerRewards(&_StakePoolManager.TransactOpts)
}

// ReceiveExecutionLayerRewards is a paid mutator transaction binding the contract method 0x33cf0ef0.
//
// Solidity: function receiveExecutionLayerRewards() payable returns()
func (_StakePoolManager *StakePoolManagerTransactorSession) ReceiveExecutionLayerRewards() (*types.Transaction, error) {
	return _StakePoolManager.Contract.ReceiveExecutionLayerRewards(&_StakePoolManager.TransactOpts)
}

// ReceiveWithdrawVaultUserShare is a paid mutator transaction binding the contract method 0x24477f11.
//
// Solidity: function receiveWithdrawVaultUserShare() payable returns()
func (_StakePoolManager *StakePoolManagerTransactor) ReceiveWithdrawVaultUserShare(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _StakePoolManager.contract.Transact(opts, "receiveWithdrawVaultUserShare")
}

// ReceiveWithdrawVaultUserShare is a paid mutator transaction binding the contract method 0x24477f11.
//
// Solidity: function receiveWithdrawVaultUserShare() payable returns()
func (_StakePoolManager *StakePoolManagerSession) ReceiveWithdrawVaultUserShare() (*types.Transaction, error) {
	return _StakePoolManager.Contract.ReceiveWithdrawVaultUserShare(&_StakePoolManager.TransactOpts)
}

// ReceiveWithdrawVaultUserShare is a paid mutator transaction binding the contract method 0x24477f11.
//
// Solidity: function receiveWithdrawVaultUserShare() payable returns()
func (_StakePoolManager *StakePoolManagerTransactorSession) ReceiveWithdrawVaultUserShare() (*types.Transaction, error) {
	return _StakePoolManager.Contract.ReceiveWithdrawVaultUserShare(&_StakePoolManager.TransactOpts)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address account) returns()
func (_StakePoolManager *StakePoolManagerTransactor) RenounceRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _StakePoolManager.contract.Transact(opts, "renounceRole", role, account)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address account) returns()
func (_StakePoolManager *StakePoolManagerSession) RenounceRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _StakePoolManager.Contract.RenounceRole(&_StakePoolManager.TransactOpts, role, account)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address account) returns()
func (_StakePoolManager *StakePoolManagerTransactorSession) RenounceRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _StakePoolManager.Contract.RenounceRole(&_StakePoolManager.TransactOpts, role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_StakePoolManager *StakePoolManagerTransactor) RevokeRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _StakePoolManager.contract.Transact(opts, "revokeRole", role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_StakePoolManager *StakePoolManagerSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _StakePoolManager.Contract.RevokeRole(&_StakePoolManager.TransactOpts, role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_StakePoolManager *StakePoolManagerTransactorSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _StakePoolManager.Contract.RevokeRole(&_StakePoolManager.TransactOpts, role, account)
}

// TransferETHToUserWithdrawManager is a paid mutator transaction binding the contract method 0x1cdfeb8f.
//
// Solidity: function transferETHToUserWithdrawManager(uint256 _amount) returns()
func (_StakePoolManager *StakePoolManagerTransactor) TransferETHToUserWithdrawManager(opts *bind.TransactOpts, _amount *big.Int) (*types.Transaction, error) {
	return _StakePoolManager.contract.Transact(opts, "transferETHToUserWithdrawManager", _amount)
}

// TransferETHToUserWithdrawManager is a paid mutator transaction binding the contract method 0x1cdfeb8f.
//
// Solidity: function transferETHToUserWithdrawManager(uint256 _amount) returns()
func (_StakePoolManager *StakePoolManagerSession) TransferETHToUserWithdrawManager(_amount *big.Int) (*types.Transaction, error) {
	return _StakePoolManager.Contract.TransferETHToUserWithdrawManager(&_StakePoolManager.TransactOpts, _amount)
}

// TransferETHToUserWithdrawManager is a paid mutator transaction binding the contract method 0x1cdfeb8f.
//
// Solidity: function transferETHToUserWithdrawManager(uint256 _amount) returns()
func (_StakePoolManager *StakePoolManagerTransactorSession) TransferETHToUserWithdrawManager(_amount *big.Int) (*types.Transaction, error) {
	return _StakePoolManager.Contract.TransferETHToUserWithdrawManager(&_StakePoolManager.TransactOpts, _amount)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_StakePoolManager *StakePoolManagerTransactor) Unpause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _StakePoolManager.contract.Transact(opts, "unpause")
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_StakePoolManager *StakePoolManagerSession) Unpause() (*types.Transaction, error) {
	return _StakePoolManager.Contract.Unpause(&_StakePoolManager.TransactOpts)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_StakePoolManager *StakePoolManagerTransactorSession) Unpause() (*types.Transaction, error) {
	return _StakePoolManager.Contract.Unpause(&_StakePoolManager.TransactOpts)
}

// UpdateExcessETHDepositCoolDown is a paid mutator transaction binding the contract method 0x7fafeb8e.
//
// Solidity: function updateExcessETHDepositCoolDown(uint256 _excessETHDepositCoolDown) returns()
func (_StakePoolManager *StakePoolManagerTransactor) UpdateExcessETHDepositCoolDown(opts *bind.TransactOpts, _excessETHDepositCoolDown *big.Int) (*types.Transaction, error) {
	return _StakePoolManager.contract.Transact(opts, "updateExcessETHDepositCoolDown", _excessETHDepositCoolDown)
}

// UpdateExcessETHDepositCoolDown is a paid mutator transaction binding the contract method 0x7fafeb8e.
//
// Solidity: function updateExcessETHDepositCoolDown(uint256 _excessETHDepositCoolDown) returns()
func (_StakePoolManager *StakePoolManagerSession) UpdateExcessETHDepositCoolDown(_excessETHDepositCoolDown *big.Int) (*types.Transaction, error) {
	return _StakePoolManager.Contract.UpdateExcessETHDepositCoolDown(&_StakePoolManager.TransactOpts, _excessETHDepositCoolDown)
}

// UpdateExcessETHDepositCoolDown is a paid mutator transaction binding the contract method 0x7fafeb8e.
//
// Solidity: function updateExcessETHDepositCoolDown(uint256 _excessETHDepositCoolDown) returns()
func (_StakePoolManager *StakePoolManagerTransactorSession) UpdateExcessETHDepositCoolDown(_excessETHDepositCoolDown *big.Int) (*types.Transaction, error) {
	return _StakePoolManager.Contract.UpdateExcessETHDepositCoolDown(&_StakePoolManager.TransactOpts, _excessETHDepositCoolDown)
}

// UpdateStaderConfig is a paid mutator transaction binding the contract method 0x9ee804cb.
//
// Solidity: function updateStaderConfig(address _staderConfig) returns()
func (_StakePoolManager *StakePoolManagerTransactor) UpdateStaderConfig(opts *bind.TransactOpts, _staderConfig common.Address) (*types.Transaction, error) {
	return _StakePoolManager.contract.Transact(opts, "updateStaderConfig", _staderConfig)
}

// UpdateStaderConfig is a paid mutator transaction binding the contract method 0x9ee804cb.
//
// Solidity: function updateStaderConfig(address _staderConfig) returns()
func (_StakePoolManager *StakePoolManagerSession) UpdateStaderConfig(_staderConfig common.Address) (*types.Transaction, error) {
	return _StakePoolManager.Contract.UpdateStaderConfig(&_StakePoolManager.TransactOpts, _staderConfig)
}

// UpdateStaderConfig is a paid mutator transaction binding the contract method 0x9ee804cb.
//
// Solidity: function updateStaderConfig(address _staderConfig) returns()
func (_StakePoolManager *StakePoolManagerTransactorSession) UpdateStaderConfig(_staderConfig common.Address) (*types.Transaction, error) {
	return _StakePoolManager.Contract.UpdateStaderConfig(&_StakePoolManager.TransactOpts, _staderConfig)
}

// ValidatorBatchDeposit is a paid mutator transaction binding the contract method 0x3e05b7dd.
//
// Solidity: function validatorBatchDeposit(uint8 _poolId) returns()
func (_StakePoolManager *StakePoolManagerTransactor) ValidatorBatchDeposit(opts *bind.TransactOpts, _poolId uint8) (*types.Transaction, error) {
	return _StakePoolManager.contract.Transact(opts, "validatorBatchDeposit", _poolId)
}

// ValidatorBatchDeposit is a paid mutator transaction binding the contract method 0x3e05b7dd.
//
// Solidity: function validatorBatchDeposit(uint8 _poolId) returns()
func (_StakePoolManager *StakePoolManagerSession) ValidatorBatchDeposit(_poolId uint8) (*types.Transaction, error) {
	return _StakePoolManager.Contract.ValidatorBatchDeposit(&_StakePoolManager.TransactOpts, _poolId)
}

// ValidatorBatchDeposit is a paid mutator transaction binding the contract method 0x3e05b7dd.
//
// Solidity: function validatorBatchDeposit(uint8 _poolId) returns()
func (_StakePoolManager *StakePoolManagerTransactorSession) ValidatorBatchDeposit(_poolId uint8) (*types.Transaction, error) {
	return _StakePoolManager.Contract.ValidatorBatchDeposit(&_StakePoolManager.TransactOpts, _poolId)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() payable returns()
func (_StakePoolManager *StakePoolManagerTransactor) Fallback(opts *bind.TransactOpts, calldata []byte) (*types.Transaction, error) {
	return _StakePoolManager.contract.RawTransact(opts, calldata)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() payable returns()
func (_StakePoolManager *StakePoolManagerSession) Fallback(calldata []byte) (*types.Transaction, error) {
	return _StakePoolManager.Contract.Fallback(&_StakePoolManager.TransactOpts, calldata)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() payable returns()
func (_StakePoolManager *StakePoolManagerTransactorSession) Fallback(calldata []byte) (*types.Transaction, error) {
	return _StakePoolManager.Contract.Fallback(&_StakePoolManager.TransactOpts, calldata)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_StakePoolManager *StakePoolManagerTransactor) Receive(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _StakePoolManager.contract.RawTransact(opts, nil) // calldata is disallowed for receive function
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_StakePoolManager *StakePoolManagerSession) Receive() (*types.Transaction, error) {
	return _StakePoolManager.Contract.Receive(&_StakePoolManager.TransactOpts)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_StakePoolManager *StakePoolManagerTransactorSession) Receive() (*types.Transaction, error) {
	return _StakePoolManager.Contract.Receive(&_StakePoolManager.TransactOpts)
}

// StakePoolManagerAuctionedEthReceivedIterator is returned from FilterAuctionedEthReceived and is used to iterate over the raw logs and unpacked data for AuctionedEthReceived events raised by the StakePoolManager contract.
type StakePoolManagerAuctionedEthReceivedIterator struct {
	Event *StakePoolManagerAuctionedEthReceived // Event containing the contract specifics and raw log

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
func (it *StakePoolManagerAuctionedEthReceivedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StakePoolManagerAuctionedEthReceived)
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
		it.Event = new(StakePoolManagerAuctionedEthReceived)
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
func (it *StakePoolManagerAuctionedEthReceivedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StakePoolManagerAuctionedEthReceivedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StakePoolManagerAuctionedEthReceived represents a AuctionedEthReceived event raised by the StakePoolManager contract.
type StakePoolManagerAuctionedEthReceived struct {
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterAuctionedEthReceived is a free log retrieval operation binding the contract event 0xffb49c9f940c060c51ce2a0b874b4fd4f5c0bc9cb4d60f0e9a333760dcb236ed.
//
// Solidity: event AuctionedEthReceived(uint256 amount)
func (_StakePoolManager *StakePoolManagerFilterer) FilterAuctionedEthReceived(opts *bind.FilterOpts) (*StakePoolManagerAuctionedEthReceivedIterator, error) {

	logs, sub, err := _StakePoolManager.contract.FilterLogs(opts, "AuctionedEthReceived")
	if err != nil {
		return nil, err
	}
	return &StakePoolManagerAuctionedEthReceivedIterator{contract: _StakePoolManager.contract, event: "AuctionedEthReceived", logs: logs, sub: sub}, nil
}

// WatchAuctionedEthReceived is a free log subscription operation binding the contract event 0xffb49c9f940c060c51ce2a0b874b4fd4f5c0bc9cb4d60f0e9a333760dcb236ed.
//
// Solidity: event AuctionedEthReceived(uint256 amount)
func (_StakePoolManager *StakePoolManagerFilterer) WatchAuctionedEthReceived(opts *bind.WatchOpts, sink chan<- *StakePoolManagerAuctionedEthReceived) (event.Subscription, error) {

	logs, sub, err := _StakePoolManager.contract.WatchLogs(opts, "AuctionedEthReceived")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StakePoolManagerAuctionedEthReceived)
				if err := _StakePoolManager.contract.UnpackLog(event, "AuctionedEthReceived", log); err != nil {
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

// ParseAuctionedEthReceived is a log parse operation binding the contract event 0xffb49c9f940c060c51ce2a0b874b4fd4f5c0bc9cb4d60f0e9a333760dcb236ed.
//
// Solidity: event AuctionedEthReceived(uint256 amount)
func (_StakePoolManager *StakePoolManagerFilterer) ParseAuctionedEthReceived(log types.Log) (*StakePoolManagerAuctionedEthReceived, error) {
	event := new(StakePoolManagerAuctionedEthReceived)
	if err := _StakePoolManager.contract.UnpackLog(event, "AuctionedEthReceived", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StakePoolManagerDepositedIterator is returned from FilterDeposited and is used to iterate over the raw logs and unpacked data for Deposited events raised by the StakePoolManager contract.
type StakePoolManagerDepositedIterator struct {
	Event *StakePoolManagerDeposited // Event containing the contract specifics and raw log

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
func (it *StakePoolManagerDepositedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StakePoolManagerDeposited)
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
		it.Event = new(StakePoolManagerDeposited)
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
func (it *StakePoolManagerDepositedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StakePoolManagerDepositedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StakePoolManagerDeposited represents a Deposited event raised by the StakePoolManager contract.
type StakePoolManagerDeposited struct {
	Caller common.Address
	Owner  common.Address
	Assets *big.Int
	Shares *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterDeposited is a free log retrieval operation binding the contract event 0xf5681f9d0db1b911ac18ee83d515a1cf1051853a9eae418316a2fdf7dea427c5.
//
// Solidity: event Deposited(address indexed caller, address indexed owner, uint256 assets, uint256 shares)
func (_StakePoolManager *StakePoolManagerFilterer) FilterDeposited(opts *bind.FilterOpts, caller []common.Address, owner []common.Address) (*StakePoolManagerDepositedIterator, error) {

	var callerRule []interface{}
	for _, callerItem := range caller {
		callerRule = append(callerRule, callerItem)
	}
	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}

	logs, sub, err := _StakePoolManager.contract.FilterLogs(opts, "Deposited", callerRule, ownerRule)
	if err != nil {
		return nil, err
	}
	return &StakePoolManagerDepositedIterator{contract: _StakePoolManager.contract, event: "Deposited", logs: logs, sub: sub}, nil
}

// WatchDeposited is a free log subscription operation binding the contract event 0xf5681f9d0db1b911ac18ee83d515a1cf1051853a9eae418316a2fdf7dea427c5.
//
// Solidity: event Deposited(address indexed caller, address indexed owner, uint256 assets, uint256 shares)
func (_StakePoolManager *StakePoolManagerFilterer) WatchDeposited(opts *bind.WatchOpts, sink chan<- *StakePoolManagerDeposited, caller []common.Address, owner []common.Address) (event.Subscription, error) {

	var callerRule []interface{}
	for _, callerItem := range caller {
		callerRule = append(callerRule, callerItem)
	}
	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}

	logs, sub, err := _StakePoolManager.contract.WatchLogs(opts, "Deposited", callerRule, ownerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StakePoolManagerDeposited)
				if err := _StakePoolManager.contract.UnpackLog(event, "Deposited", log); err != nil {
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

// ParseDeposited is a log parse operation binding the contract event 0xf5681f9d0db1b911ac18ee83d515a1cf1051853a9eae418316a2fdf7dea427c5.
//
// Solidity: event Deposited(address indexed caller, address indexed owner, uint256 assets, uint256 shares)
func (_StakePoolManager *StakePoolManagerFilterer) ParseDeposited(log types.Log) (*StakePoolManagerDeposited, error) {
	event := new(StakePoolManagerDeposited)
	if err := _StakePoolManager.contract.UnpackLog(event, "Deposited", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StakePoolManagerETHTransferredToPoolIterator is returned from FilterETHTransferredToPool and is used to iterate over the raw logs and unpacked data for ETHTransferredToPool events raised by the StakePoolManager contract.
type StakePoolManagerETHTransferredToPoolIterator struct {
	Event *StakePoolManagerETHTransferredToPool // Event containing the contract specifics and raw log

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
func (it *StakePoolManagerETHTransferredToPoolIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StakePoolManagerETHTransferredToPool)
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
		it.Event = new(StakePoolManagerETHTransferredToPool)
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
func (it *StakePoolManagerETHTransferredToPoolIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StakePoolManagerETHTransferredToPoolIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StakePoolManagerETHTransferredToPool represents a ETHTransferredToPool event raised by the StakePoolManager contract.
type StakePoolManagerETHTransferredToPool struct {
	PoolId         *big.Int
	PoolAddress    common.Address
	ValidatorCount *big.Int
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterETHTransferredToPool is a free log retrieval operation binding the contract event 0x0f4ee8a1358b01e75e0b5291e986fa643035327081fc296a9bb60449257e988a.
//
// Solidity: event ETHTransferredToPool(uint256 indexed poolId, address poolAddress, uint256 validatorCount)
func (_StakePoolManager *StakePoolManagerFilterer) FilterETHTransferredToPool(opts *bind.FilterOpts, poolId []*big.Int) (*StakePoolManagerETHTransferredToPoolIterator, error) {

	var poolIdRule []interface{}
	for _, poolIdItem := range poolId {
		poolIdRule = append(poolIdRule, poolIdItem)
	}

	logs, sub, err := _StakePoolManager.contract.FilterLogs(opts, "ETHTransferredToPool", poolIdRule)
	if err != nil {
		return nil, err
	}
	return &StakePoolManagerETHTransferredToPoolIterator{contract: _StakePoolManager.contract, event: "ETHTransferredToPool", logs: logs, sub: sub}, nil
}

// WatchETHTransferredToPool is a free log subscription operation binding the contract event 0x0f4ee8a1358b01e75e0b5291e986fa643035327081fc296a9bb60449257e988a.
//
// Solidity: event ETHTransferredToPool(uint256 indexed poolId, address poolAddress, uint256 validatorCount)
func (_StakePoolManager *StakePoolManagerFilterer) WatchETHTransferredToPool(opts *bind.WatchOpts, sink chan<- *StakePoolManagerETHTransferredToPool, poolId []*big.Int) (event.Subscription, error) {

	var poolIdRule []interface{}
	for _, poolIdItem := range poolId {
		poolIdRule = append(poolIdRule, poolIdItem)
	}

	logs, sub, err := _StakePoolManager.contract.WatchLogs(opts, "ETHTransferredToPool", poolIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StakePoolManagerETHTransferredToPool)
				if err := _StakePoolManager.contract.UnpackLog(event, "ETHTransferredToPool", log); err != nil {
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

// ParseETHTransferredToPool is a log parse operation binding the contract event 0x0f4ee8a1358b01e75e0b5291e986fa643035327081fc296a9bb60449257e988a.
//
// Solidity: event ETHTransferredToPool(uint256 indexed poolId, address poolAddress, uint256 validatorCount)
func (_StakePoolManager *StakePoolManagerFilterer) ParseETHTransferredToPool(log types.Log) (*StakePoolManagerETHTransferredToPool, error) {
	event := new(StakePoolManagerETHTransferredToPool)
	if err := _StakePoolManager.contract.UnpackLog(event, "ETHTransferredToPool", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StakePoolManagerExecutionLayerRewardsReceivedIterator is returned from FilterExecutionLayerRewardsReceived and is used to iterate over the raw logs and unpacked data for ExecutionLayerRewardsReceived events raised by the StakePoolManager contract.
type StakePoolManagerExecutionLayerRewardsReceivedIterator struct {
	Event *StakePoolManagerExecutionLayerRewardsReceived // Event containing the contract specifics and raw log

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
func (it *StakePoolManagerExecutionLayerRewardsReceivedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StakePoolManagerExecutionLayerRewardsReceived)
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
		it.Event = new(StakePoolManagerExecutionLayerRewardsReceived)
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
func (it *StakePoolManagerExecutionLayerRewardsReceivedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StakePoolManagerExecutionLayerRewardsReceivedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StakePoolManagerExecutionLayerRewardsReceived represents a ExecutionLayerRewardsReceived event raised by the StakePoolManager contract.
type StakePoolManagerExecutionLayerRewardsReceived struct {
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterExecutionLayerRewardsReceived is a free log retrieval operation binding the contract event 0xe2abdf52ca718fb2feac4ce491e7a4b908e49e318734391dfee4e03fc5acf4d6.
//
// Solidity: event ExecutionLayerRewardsReceived(uint256 amount)
func (_StakePoolManager *StakePoolManagerFilterer) FilterExecutionLayerRewardsReceived(opts *bind.FilterOpts) (*StakePoolManagerExecutionLayerRewardsReceivedIterator, error) {

	logs, sub, err := _StakePoolManager.contract.FilterLogs(opts, "ExecutionLayerRewardsReceived")
	if err != nil {
		return nil, err
	}
	return &StakePoolManagerExecutionLayerRewardsReceivedIterator{contract: _StakePoolManager.contract, event: "ExecutionLayerRewardsReceived", logs: logs, sub: sub}, nil
}

// WatchExecutionLayerRewardsReceived is a free log subscription operation binding the contract event 0xe2abdf52ca718fb2feac4ce491e7a4b908e49e318734391dfee4e03fc5acf4d6.
//
// Solidity: event ExecutionLayerRewardsReceived(uint256 amount)
func (_StakePoolManager *StakePoolManagerFilterer) WatchExecutionLayerRewardsReceived(opts *bind.WatchOpts, sink chan<- *StakePoolManagerExecutionLayerRewardsReceived) (event.Subscription, error) {

	logs, sub, err := _StakePoolManager.contract.WatchLogs(opts, "ExecutionLayerRewardsReceived")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StakePoolManagerExecutionLayerRewardsReceived)
				if err := _StakePoolManager.contract.UnpackLog(event, "ExecutionLayerRewardsReceived", log); err != nil {
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

// ParseExecutionLayerRewardsReceived is a log parse operation binding the contract event 0xe2abdf52ca718fb2feac4ce491e7a4b908e49e318734391dfee4e03fc5acf4d6.
//
// Solidity: event ExecutionLayerRewardsReceived(uint256 amount)
func (_StakePoolManager *StakePoolManagerFilterer) ParseExecutionLayerRewardsReceived(log types.Log) (*StakePoolManagerExecutionLayerRewardsReceived, error) {
	event := new(StakePoolManagerExecutionLayerRewardsReceived)
	if err := _StakePoolManager.contract.UnpackLog(event, "ExecutionLayerRewardsReceived", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StakePoolManagerInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the StakePoolManager contract.
type StakePoolManagerInitializedIterator struct {
	Event *StakePoolManagerInitialized // Event containing the contract specifics and raw log

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
func (it *StakePoolManagerInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StakePoolManagerInitialized)
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
		it.Event = new(StakePoolManagerInitialized)
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
func (it *StakePoolManagerInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StakePoolManagerInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StakePoolManagerInitialized represents a Initialized event raised by the StakePoolManager contract.
type StakePoolManagerInitialized struct {
	Version uint8
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_StakePoolManager *StakePoolManagerFilterer) FilterInitialized(opts *bind.FilterOpts) (*StakePoolManagerInitializedIterator, error) {

	logs, sub, err := _StakePoolManager.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &StakePoolManagerInitializedIterator{contract: _StakePoolManager.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_StakePoolManager *StakePoolManagerFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *StakePoolManagerInitialized) (event.Subscription, error) {

	logs, sub, err := _StakePoolManager.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StakePoolManagerInitialized)
				if err := _StakePoolManager.contract.UnpackLog(event, "Initialized", log); err != nil {
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
func (_StakePoolManager *StakePoolManagerFilterer) ParseInitialized(log types.Log) (*StakePoolManagerInitialized, error) {
	event := new(StakePoolManagerInitialized)
	if err := _StakePoolManager.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StakePoolManagerPausedIterator is returned from FilterPaused and is used to iterate over the raw logs and unpacked data for Paused events raised by the StakePoolManager contract.
type StakePoolManagerPausedIterator struct {
	Event *StakePoolManagerPaused // Event containing the contract specifics and raw log

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
func (it *StakePoolManagerPausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StakePoolManagerPaused)
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
		it.Event = new(StakePoolManagerPaused)
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
func (it *StakePoolManagerPausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StakePoolManagerPausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StakePoolManagerPaused represents a Paused event raised by the StakePoolManager contract.
type StakePoolManagerPaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterPaused is a free log retrieval operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_StakePoolManager *StakePoolManagerFilterer) FilterPaused(opts *bind.FilterOpts) (*StakePoolManagerPausedIterator, error) {

	logs, sub, err := _StakePoolManager.contract.FilterLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return &StakePoolManagerPausedIterator{contract: _StakePoolManager.contract, event: "Paused", logs: logs, sub: sub}, nil
}

// WatchPaused is a free log subscription operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_StakePoolManager *StakePoolManagerFilterer) WatchPaused(opts *bind.WatchOpts, sink chan<- *StakePoolManagerPaused) (event.Subscription, error) {

	logs, sub, err := _StakePoolManager.contract.WatchLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StakePoolManagerPaused)
				if err := _StakePoolManager.contract.UnpackLog(event, "Paused", log); err != nil {
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
func (_StakePoolManager *StakePoolManagerFilterer) ParsePaused(log types.Log) (*StakePoolManagerPaused, error) {
	event := new(StakePoolManagerPaused)
	if err := _StakePoolManager.contract.UnpackLog(event, "Paused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StakePoolManagerReceivedExcessEthFromPoolIterator is returned from FilterReceivedExcessEthFromPool and is used to iterate over the raw logs and unpacked data for ReceivedExcessEthFromPool events raised by the StakePoolManager contract.
type StakePoolManagerReceivedExcessEthFromPoolIterator struct {
	Event *StakePoolManagerReceivedExcessEthFromPool // Event containing the contract specifics and raw log

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
func (it *StakePoolManagerReceivedExcessEthFromPoolIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StakePoolManagerReceivedExcessEthFromPool)
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
		it.Event = new(StakePoolManagerReceivedExcessEthFromPool)
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
func (it *StakePoolManagerReceivedExcessEthFromPoolIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StakePoolManagerReceivedExcessEthFromPoolIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StakePoolManagerReceivedExcessEthFromPool represents a ReceivedExcessEthFromPool event raised by the StakePoolManager contract.
type StakePoolManagerReceivedExcessEthFromPool struct {
	PoolId uint8
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterReceivedExcessEthFromPool is a free log retrieval operation binding the contract event 0x9d702bbfa67b9ff48c9c450576e1f53296afd8f7bf30211d771128708586f5b1.
//
// Solidity: event ReceivedExcessEthFromPool(uint8 indexed poolId)
func (_StakePoolManager *StakePoolManagerFilterer) FilterReceivedExcessEthFromPool(opts *bind.FilterOpts, poolId []uint8) (*StakePoolManagerReceivedExcessEthFromPoolIterator, error) {

	var poolIdRule []interface{}
	for _, poolIdItem := range poolId {
		poolIdRule = append(poolIdRule, poolIdItem)
	}

	logs, sub, err := _StakePoolManager.contract.FilterLogs(opts, "ReceivedExcessEthFromPool", poolIdRule)
	if err != nil {
		return nil, err
	}
	return &StakePoolManagerReceivedExcessEthFromPoolIterator{contract: _StakePoolManager.contract, event: "ReceivedExcessEthFromPool", logs: logs, sub: sub}, nil
}

// WatchReceivedExcessEthFromPool is a free log subscription operation binding the contract event 0x9d702bbfa67b9ff48c9c450576e1f53296afd8f7bf30211d771128708586f5b1.
//
// Solidity: event ReceivedExcessEthFromPool(uint8 indexed poolId)
func (_StakePoolManager *StakePoolManagerFilterer) WatchReceivedExcessEthFromPool(opts *bind.WatchOpts, sink chan<- *StakePoolManagerReceivedExcessEthFromPool, poolId []uint8) (event.Subscription, error) {

	var poolIdRule []interface{}
	for _, poolIdItem := range poolId {
		poolIdRule = append(poolIdRule, poolIdItem)
	}

	logs, sub, err := _StakePoolManager.contract.WatchLogs(opts, "ReceivedExcessEthFromPool", poolIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StakePoolManagerReceivedExcessEthFromPool)
				if err := _StakePoolManager.contract.UnpackLog(event, "ReceivedExcessEthFromPool", log); err != nil {
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

// ParseReceivedExcessEthFromPool is a log parse operation binding the contract event 0x9d702bbfa67b9ff48c9c450576e1f53296afd8f7bf30211d771128708586f5b1.
//
// Solidity: event ReceivedExcessEthFromPool(uint8 indexed poolId)
func (_StakePoolManager *StakePoolManagerFilterer) ParseReceivedExcessEthFromPool(log types.Log) (*StakePoolManagerReceivedExcessEthFromPool, error) {
	event := new(StakePoolManagerReceivedExcessEthFromPool)
	if err := _StakePoolManager.contract.UnpackLog(event, "ReceivedExcessEthFromPool", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StakePoolManagerRoleAdminChangedIterator is returned from FilterRoleAdminChanged and is used to iterate over the raw logs and unpacked data for RoleAdminChanged events raised by the StakePoolManager contract.
type StakePoolManagerRoleAdminChangedIterator struct {
	Event *StakePoolManagerRoleAdminChanged // Event containing the contract specifics and raw log

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
func (it *StakePoolManagerRoleAdminChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StakePoolManagerRoleAdminChanged)
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
		it.Event = new(StakePoolManagerRoleAdminChanged)
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
func (it *StakePoolManagerRoleAdminChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StakePoolManagerRoleAdminChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StakePoolManagerRoleAdminChanged represents a RoleAdminChanged event raised by the StakePoolManager contract.
type StakePoolManagerRoleAdminChanged struct {
	Role              [32]byte
	PreviousAdminRole [32]byte
	NewAdminRole      [32]byte
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterRoleAdminChanged is a free log retrieval operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_StakePoolManager *StakePoolManagerFilterer) FilterRoleAdminChanged(opts *bind.FilterOpts, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (*StakePoolManagerRoleAdminChangedIterator, error) {

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

	logs, sub, err := _StakePoolManager.contract.FilterLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return &StakePoolManagerRoleAdminChangedIterator{contract: _StakePoolManager.contract, event: "RoleAdminChanged", logs: logs, sub: sub}, nil
}

// WatchRoleAdminChanged is a free log subscription operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_StakePoolManager *StakePoolManagerFilterer) WatchRoleAdminChanged(opts *bind.WatchOpts, sink chan<- *StakePoolManagerRoleAdminChanged, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (event.Subscription, error) {

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

	logs, sub, err := _StakePoolManager.contract.WatchLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StakePoolManagerRoleAdminChanged)
				if err := _StakePoolManager.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
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
func (_StakePoolManager *StakePoolManagerFilterer) ParseRoleAdminChanged(log types.Log) (*StakePoolManagerRoleAdminChanged, error) {
	event := new(StakePoolManagerRoleAdminChanged)
	if err := _StakePoolManager.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StakePoolManagerRoleGrantedIterator is returned from FilterRoleGranted and is used to iterate over the raw logs and unpacked data for RoleGranted events raised by the StakePoolManager contract.
type StakePoolManagerRoleGrantedIterator struct {
	Event *StakePoolManagerRoleGranted // Event containing the contract specifics and raw log

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
func (it *StakePoolManagerRoleGrantedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StakePoolManagerRoleGranted)
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
		it.Event = new(StakePoolManagerRoleGranted)
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
func (it *StakePoolManagerRoleGrantedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StakePoolManagerRoleGrantedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StakePoolManagerRoleGranted represents a RoleGranted event raised by the StakePoolManager contract.
type StakePoolManagerRoleGranted struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleGranted is a free log retrieval operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_StakePoolManager *StakePoolManagerFilterer) FilterRoleGranted(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*StakePoolManagerRoleGrantedIterator, error) {

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

	logs, sub, err := _StakePoolManager.contract.FilterLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &StakePoolManagerRoleGrantedIterator{contract: _StakePoolManager.contract, event: "RoleGranted", logs: logs, sub: sub}, nil
}

// WatchRoleGranted is a free log subscription operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_StakePoolManager *StakePoolManagerFilterer) WatchRoleGranted(opts *bind.WatchOpts, sink chan<- *StakePoolManagerRoleGranted, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _StakePoolManager.contract.WatchLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StakePoolManagerRoleGranted)
				if err := _StakePoolManager.contract.UnpackLog(event, "RoleGranted", log); err != nil {
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
func (_StakePoolManager *StakePoolManagerFilterer) ParseRoleGranted(log types.Log) (*StakePoolManagerRoleGranted, error) {
	event := new(StakePoolManagerRoleGranted)
	if err := _StakePoolManager.contract.UnpackLog(event, "RoleGranted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StakePoolManagerRoleRevokedIterator is returned from FilterRoleRevoked and is used to iterate over the raw logs and unpacked data for RoleRevoked events raised by the StakePoolManager contract.
type StakePoolManagerRoleRevokedIterator struct {
	Event *StakePoolManagerRoleRevoked // Event containing the contract specifics and raw log

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
func (it *StakePoolManagerRoleRevokedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StakePoolManagerRoleRevoked)
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
		it.Event = new(StakePoolManagerRoleRevoked)
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
func (it *StakePoolManagerRoleRevokedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StakePoolManagerRoleRevokedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StakePoolManagerRoleRevoked represents a RoleRevoked event raised by the StakePoolManager contract.
type StakePoolManagerRoleRevoked struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleRevoked is a free log retrieval operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_StakePoolManager *StakePoolManagerFilterer) FilterRoleRevoked(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*StakePoolManagerRoleRevokedIterator, error) {

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

	logs, sub, err := _StakePoolManager.contract.FilterLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &StakePoolManagerRoleRevokedIterator{contract: _StakePoolManager.contract, event: "RoleRevoked", logs: logs, sub: sub}, nil
}

// WatchRoleRevoked is a free log subscription operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_StakePoolManager *StakePoolManagerFilterer) WatchRoleRevoked(opts *bind.WatchOpts, sink chan<- *StakePoolManagerRoleRevoked, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _StakePoolManager.contract.WatchLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StakePoolManagerRoleRevoked)
				if err := _StakePoolManager.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
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
func (_StakePoolManager *StakePoolManagerFilterer) ParseRoleRevoked(log types.Log) (*StakePoolManagerRoleRevoked, error) {
	event := new(StakePoolManagerRoleRevoked)
	if err := _StakePoolManager.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StakePoolManagerTransferredETHToUserWithdrawManagerIterator is returned from FilterTransferredETHToUserWithdrawManager and is used to iterate over the raw logs and unpacked data for TransferredETHToUserWithdrawManager events raised by the StakePoolManager contract.
type StakePoolManagerTransferredETHToUserWithdrawManagerIterator struct {
	Event *StakePoolManagerTransferredETHToUserWithdrawManager // Event containing the contract specifics and raw log

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
func (it *StakePoolManagerTransferredETHToUserWithdrawManagerIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StakePoolManagerTransferredETHToUserWithdrawManager)
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
		it.Event = new(StakePoolManagerTransferredETHToUserWithdrawManager)
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
func (it *StakePoolManagerTransferredETHToUserWithdrawManagerIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StakePoolManagerTransferredETHToUserWithdrawManagerIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StakePoolManagerTransferredETHToUserWithdrawManager represents a TransferredETHToUserWithdrawManager event raised by the StakePoolManager contract.
type StakePoolManagerTransferredETHToUserWithdrawManager struct {
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterTransferredETHToUserWithdrawManager is a free log retrieval operation binding the contract event 0xfcf1373cbfb78832a864dcce3862324e51116876bd08423a61b5ed6d5c03f421.
//
// Solidity: event TransferredETHToUserWithdrawManager(uint256 amount)
func (_StakePoolManager *StakePoolManagerFilterer) FilterTransferredETHToUserWithdrawManager(opts *bind.FilterOpts) (*StakePoolManagerTransferredETHToUserWithdrawManagerIterator, error) {

	logs, sub, err := _StakePoolManager.contract.FilterLogs(opts, "TransferredETHToUserWithdrawManager")
	if err != nil {
		return nil, err
	}
	return &StakePoolManagerTransferredETHToUserWithdrawManagerIterator{contract: _StakePoolManager.contract, event: "TransferredETHToUserWithdrawManager", logs: logs, sub: sub}, nil
}

// WatchTransferredETHToUserWithdrawManager is a free log subscription operation binding the contract event 0xfcf1373cbfb78832a864dcce3862324e51116876bd08423a61b5ed6d5c03f421.
//
// Solidity: event TransferredETHToUserWithdrawManager(uint256 amount)
func (_StakePoolManager *StakePoolManagerFilterer) WatchTransferredETHToUserWithdrawManager(opts *bind.WatchOpts, sink chan<- *StakePoolManagerTransferredETHToUserWithdrawManager) (event.Subscription, error) {

	logs, sub, err := _StakePoolManager.contract.WatchLogs(opts, "TransferredETHToUserWithdrawManager")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StakePoolManagerTransferredETHToUserWithdrawManager)
				if err := _StakePoolManager.contract.UnpackLog(event, "TransferredETHToUserWithdrawManager", log); err != nil {
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

// ParseTransferredETHToUserWithdrawManager is a log parse operation binding the contract event 0xfcf1373cbfb78832a864dcce3862324e51116876bd08423a61b5ed6d5c03f421.
//
// Solidity: event TransferredETHToUserWithdrawManager(uint256 amount)
func (_StakePoolManager *StakePoolManagerFilterer) ParseTransferredETHToUserWithdrawManager(log types.Log) (*StakePoolManagerTransferredETHToUserWithdrawManager, error) {
	event := new(StakePoolManagerTransferredETHToUserWithdrawManager)
	if err := _StakePoolManager.contract.UnpackLog(event, "TransferredETHToUserWithdrawManager", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StakePoolManagerUnpausedIterator is returned from FilterUnpaused and is used to iterate over the raw logs and unpacked data for Unpaused events raised by the StakePoolManager contract.
type StakePoolManagerUnpausedIterator struct {
	Event *StakePoolManagerUnpaused // Event containing the contract specifics and raw log

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
func (it *StakePoolManagerUnpausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StakePoolManagerUnpaused)
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
		it.Event = new(StakePoolManagerUnpaused)
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
func (it *StakePoolManagerUnpausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StakePoolManagerUnpausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StakePoolManagerUnpaused represents a Unpaused event raised by the StakePoolManager contract.
type StakePoolManagerUnpaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterUnpaused is a free log retrieval operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_StakePoolManager *StakePoolManagerFilterer) FilterUnpaused(opts *bind.FilterOpts) (*StakePoolManagerUnpausedIterator, error) {

	logs, sub, err := _StakePoolManager.contract.FilterLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return &StakePoolManagerUnpausedIterator{contract: _StakePoolManager.contract, event: "Unpaused", logs: logs, sub: sub}, nil
}

// WatchUnpaused is a free log subscription operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_StakePoolManager *StakePoolManagerFilterer) WatchUnpaused(opts *bind.WatchOpts, sink chan<- *StakePoolManagerUnpaused) (event.Subscription, error) {

	logs, sub, err := _StakePoolManager.contract.WatchLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StakePoolManagerUnpaused)
				if err := _StakePoolManager.contract.UnpackLog(event, "Unpaused", log); err != nil {
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
func (_StakePoolManager *StakePoolManagerFilterer) ParseUnpaused(log types.Log) (*StakePoolManagerUnpaused, error) {
	event := new(StakePoolManagerUnpaused)
	if err := _StakePoolManager.contract.UnpackLog(event, "Unpaused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StakePoolManagerUpdatedExcessETHDepositCoolDownIterator is returned from FilterUpdatedExcessETHDepositCoolDown and is used to iterate over the raw logs and unpacked data for UpdatedExcessETHDepositCoolDown events raised by the StakePoolManager contract.
type StakePoolManagerUpdatedExcessETHDepositCoolDownIterator struct {
	Event *StakePoolManagerUpdatedExcessETHDepositCoolDown // Event containing the contract specifics and raw log

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
func (it *StakePoolManagerUpdatedExcessETHDepositCoolDownIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StakePoolManagerUpdatedExcessETHDepositCoolDown)
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
		it.Event = new(StakePoolManagerUpdatedExcessETHDepositCoolDown)
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
func (it *StakePoolManagerUpdatedExcessETHDepositCoolDownIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StakePoolManagerUpdatedExcessETHDepositCoolDownIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StakePoolManagerUpdatedExcessETHDepositCoolDown represents a UpdatedExcessETHDepositCoolDown event raised by the StakePoolManager contract.
type StakePoolManagerUpdatedExcessETHDepositCoolDown struct {
	ExcessETHDepositCoolDown *big.Int
	Raw                      types.Log // Blockchain specific contextual infos
}

// FilterUpdatedExcessETHDepositCoolDown is a free log retrieval operation binding the contract event 0xdf91929e13446a67ab8b4f37b193a2650935a1ac5cd8f39586f386afd552b724.
//
// Solidity: event UpdatedExcessETHDepositCoolDown(uint256 excessETHDepositCoolDown)
func (_StakePoolManager *StakePoolManagerFilterer) FilterUpdatedExcessETHDepositCoolDown(opts *bind.FilterOpts) (*StakePoolManagerUpdatedExcessETHDepositCoolDownIterator, error) {

	logs, sub, err := _StakePoolManager.contract.FilterLogs(opts, "UpdatedExcessETHDepositCoolDown")
	if err != nil {
		return nil, err
	}
	return &StakePoolManagerUpdatedExcessETHDepositCoolDownIterator{contract: _StakePoolManager.contract, event: "UpdatedExcessETHDepositCoolDown", logs: logs, sub: sub}, nil
}

// WatchUpdatedExcessETHDepositCoolDown is a free log subscription operation binding the contract event 0xdf91929e13446a67ab8b4f37b193a2650935a1ac5cd8f39586f386afd552b724.
//
// Solidity: event UpdatedExcessETHDepositCoolDown(uint256 excessETHDepositCoolDown)
func (_StakePoolManager *StakePoolManagerFilterer) WatchUpdatedExcessETHDepositCoolDown(opts *bind.WatchOpts, sink chan<- *StakePoolManagerUpdatedExcessETHDepositCoolDown) (event.Subscription, error) {

	logs, sub, err := _StakePoolManager.contract.WatchLogs(opts, "UpdatedExcessETHDepositCoolDown")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StakePoolManagerUpdatedExcessETHDepositCoolDown)
				if err := _StakePoolManager.contract.UnpackLog(event, "UpdatedExcessETHDepositCoolDown", log); err != nil {
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

// ParseUpdatedExcessETHDepositCoolDown is a log parse operation binding the contract event 0xdf91929e13446a67ab8b4f37b193a2650935a1ac5cd8f39586f386afd552b724.
//
// Solidity: event UpdatedExcessETHDepositCoolDown(uint256 excessETHDepositCoolDown)
func (_StakePoolManager *StakePoolManagerFilterer) ParseUpdatedExcessETHDepositCoolDown(log types.Log) (*StakePoolManagerUpdatedExcessETHDepositCoolDown, error) {
	event := new(StakePoolManagerUpdatedExcessETHDepositCoolDown)
	if err := _StakePoolManager.contract.UnpackLog(event, "UpdatedExcessETHDepositCoolDown", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StakePoolManagerUpdatedStaderConfigIterator is returned from FilterUpdatedStaderConfig and is used to iterate over the raw logs and unpacked data for UpdatedStaderConfig events raised by the StakePoolManager contract.
type StakePoolManagerUpdatedStaderConfigIterator struct {
	Event *StakePoolManagerUpdatedStaderConfig // Event containing the contract specifics and raw log

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
func (it *StakePoolManagerUpdatedStaderConfigIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StakePoolManagerUpdatedStaderConfig)
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
		it.Event = new(StakePoolManagerUpdatedStaderConfig)
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
func (it *StakePoolManagerUpdatedStaderConfigIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StakePoolManagerUpdatedStaderConfigIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StakePoolManagerUpdatedStaderConfig represents a UpdatedStaderConfig event raised by the StakePoolManager contract.
type StakePoolManagerUpdatedStaderConfig struct {
	StaderConfig common.Address
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterUpdatedStaderConfig is a free log retrieval operation binding the contract event 0xdb2219043d7b197cb235f1af0cf6d782d77dee3de19e3f4fb6d39aae633b4485.
//
// Solidity: event UpdatedStaderConfig(address staderConfig)
func (_StakePoolManager *StakePoolManagerFilterer) FilterUpdatedStaderConfig(opts *bind.FilterOpts) (*StakePoolManagerUpdatedStaderConfigIterator, error) {

	logs, sub, err := _StakePoolManager.contract.FilterLogs(opts, "UpdatedStaderConfig")
	if err != nil {
		return nil, err
	}
	return &StakePoolManagerUpdatedStaderConfigIterator{contract: _StakePoolManager.contract, event: "UpdatedStaderConfig", logs: logs, sub: sub}, nil
}

// WatchUpdatedStaderConfig is a free log subscription operation binding the contract event 0xdb2219043d7b197cb235f1af0cf6d782d77dee3de19e3f4fb6d39aae633b4485.
//
// Solidity: event UpdatedStaderConfig(address staderConfig)
func (_StakePoolManager *StakePoolManagerFilterer) WatchUpdatedStaderConfig(opts *bind.WatchOpts, sink chan<- *StakePoolManagerUpdatedStaderConfig) (event.Subscription, error) {

	logs, sub, err := _StakePoolManager.contract.WatchLogs(opts, "UpdatedStaderConfig")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StakePoolManagerUpdatedStaderConfig)
				if err := _StakePoolManager.contract.UnpackLog(event, "UpdatedStaderConfig", log); err != nil {
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
func (_StakePoolManager *StakePoolManagerFilterer) ParseUpdatedStaderConfig(log types.Log) (*StakePoolManagerUpdatedStaderConfig, error) {
	event := new(StakePoolManagerUpdatedStaderConfig)
	if err := _StakePoolManager.contract.UnpackLog(event, "UpdatedStaderConfig", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StakePoolManagerWithdrawVaultUserShareReceivedIterator is returned from FilterWithdrawVaultUserShareReceived and is used to iterate over the raw logs and unpacked data for WithdrawVaultUserShareReceived events raised by the StakePoolManager contract.
type StakePoolManagerWithdrawVaultUserShareReceivedIterator struct {
	Event *StakePoolManagerWithdrawVaultUserShareReceived // Event containing the contract specifics and raw log

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
func (it *StakePoolManagerWithdrawVaultUserShareReceivedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StakePoolManagerWithdrawVaultUserShareReceived)
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
		it.Event = new(StakePoolManagerWithdrawVaultUserShareReceived)
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
func (it *StakePoolManagerWithdrawVaultUserShareReceivedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StakePoolManagerWithdrawVaultUserShareReceivedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StakePoolManagerWithdrawVaultUserShareReceived represents a WithdrawVaultUserShareReceived event raised by the StakePoolManager contract.
type StakePoolManagerWithdrawVaultUserShareReceived struct {
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterWithdrawVaultUserShareReceived is a free log retrieval operation binding the contract event 0x5569069b23fc6ce6fbe88bdd95e974a82fb3d433cc2ebcbe4dd70af6ac63f83b.
//
// Solidity: event WithdrawVaultUserShareReceived(uint256 amount)
func (_StakePoolManager *StakePoolManagerFilterer) FilterWithdrawVaultUserShareReceived(opts *bind.FilterOpts) (*StakePoolManagerWithdrawVaultUserShareReceivedIterator, error) {

	logs, sub, err := _StakePoolManager.contract.FilterLogs(opts, "WithdrawVaultUserShareReceived")
	if err != nil {
		return nil, err
	}
	return &StakePoolManagerWithdrawVaultUserShareReceivedIterator{contract: _StakePoolManager.contract, event: "WithdrawVaultUserShareReceived", logs: logs, sub: sub}, nil
}

// WatchWithdrawVaultUserShareReceived is a free log subscription operation binding the contract event 0x5569069b23fc6ce6fbe88bdd95e974a82fb3d433cc2ebcbe4dd70af6ac63f83b.
//
// Solidity: event WithdrawVaultUserShareReceived(uint256 amount)
func (_StakePoolManager *StakePoolManagerFilterer) WatchWithdrawVaultUserShareReceived(opts *bind.WatchOpts, sink chan<- *StakePoolManagerWithdrawVaultUserShareReceived) (event.Subscription, error) {

	logs, sub, err := _StakePoolManager.contract.WatchLogs(opts, "WithdrawVaultUserShareReceived")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StakePoolManagerWithdrawVaultUserShareReceived)
				if err := _StakePoolManager.contract.UnpackLog(event, "WithdrawVaultUserShareReceived", log); err != nil {
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

// ParseWithdrawVaultUserShareReceived is a log parse operation binding the contract event 0x5569069b23fc6ce6fbe88bdd95e974a82fb3d433cc2ebcbe4dd70af6ac63f83b.
//
// Solidity: event WithdrawVaultUserShareReceived(uint256 amount)
func (_StakePoolManager *StakePoolManagerFilterer) ParseWithdrawVaultUserShareReceived(log types.Log) (*StakePoolManagerWithdrawVaultUserShareReceived, error) {
	event := new(StakePoolManagerWithdrawVaultUserShareReceived)
	if err := _StakePoolManager.contract.UnpackLog(event, "WithdrawVaultUserShareReceived", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
