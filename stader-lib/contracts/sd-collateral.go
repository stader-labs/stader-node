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

// SdCollateralMetaData contains all meta data concerning the SdCollateral contract.
var SdCollateralMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"version\",\"type\":\"uint8\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Paused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"previousAdminRole\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"newAdminRole\",\"type\":\"bytes32\"}],\"name\":\"RoleAdminChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleGranted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleRevoked\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Unpaused\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"DEFAULT_ADMIN_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_ethAmount\",\"type\":\"uint256\"}],\"name\":\"convertETHToSD\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_sdAmount\",\"type\":\"uint256\"}],\"name\":\"convertSDToETH\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_sdAmount\",\"type\":\"uint256\"}],\"name\":\"convertSDToShares\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_numShares\",\"type\":\"uint256\"}],\"name\":\"convertSharesToSD\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_sdAmount\",\"type\":\"uint256\"}],\"name\":\"depositSDAsCollateral\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_operator\",\"type\":\"address\"}],\"name\":\"getOperatorSDBalance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"}],\"name\":\"getRoleAdmin\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"grantRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_operator\",\"type\":\"address\"},{\"internalType\":\"uint8\",\"name\":\"_poolId\",\"type\":\"uint8\"},{\"internalType\":\"uint32\",\"name\":\"_numValidators\",\"type\":\"uint32\"}],\"name\":\"hasEnoughSDCollateral\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"hasRole\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_admin\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_sdERC20Addr\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_priceFetcherAddr\",\"type\":\"address\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"operatorShares\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"paused\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"name\":\"poolThresholdbyPoolId\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"minThreshold\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"withdrawThreshold\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"units\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"priceFetcher\",\"outputs\":[{\"internalType\":\"contractIPriceFetcher\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"renounceRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"revokeRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"sdERC20\",\"outputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalSDCollateral\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalShares\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"_poolId\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"_minThreshold\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_withdrawThreshold\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"_units\",\"type\":\"string\"}],\"name\":\"updatePoolThreshold\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_requestedSD\",\"type\":\"uint256\"}],\"name\":\"withdraw\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// SdCollateralABI is the input ABI used to generate the binding from.
// Deprecated: Use SdCollateralMetaData.ABI instead.
var SdCollateralABI = SdCollateralMetaData.ABI

// SdCollateral is an auto generated Go binding around an Ethereum contract.
type SdCollateral struct {
	SdCollateralCaller     // Read-only binding to the contract
	SdCollateralTransactor // Write-only binding to the contract
	SdCollateralFilterer   // Log filterer for contract events
}

// SdCollateralCaller is an auto generated read-only Go binding around an Ethereum contract.
type SdCollateralCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SdCollateralTransactor is an auto generated write-only Go binding around an Ethereum contract.
type SdCollateralTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SdCollateralFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type SdCollateralFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SdCollateralSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type SdCollateralSession struct {
	Contract     *SdCollateral     // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// SdCollateralCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type SdCollateralCallerSession struct {
	Contract *SdCollateralCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts       // Call options to use throughout this session
}

// SdCollateralTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type SdCollateralTransactorSession struct {
	Contract     *SdCollateralTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts       // Transaction auth options to use throughout this session
}

// SdCollateralRaw is an auto generated low-level Go binding around an Ethereum contract.
type SdCollateralRaw struct {
	Contract *SdCollateral // Generic contract binding to access the raw methods on
}

// SdCollateralCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type SdCollateralCallerRaw struct {
	Contract *SdCollateralCaller // Generic read-only contract binding to access the raw methods on
}

// SdCollateralTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type SdCollateralTransactorRaw struct {
	Contract *SdCollateralTransactor // Generic write-only contract binding to access the raw methods on
}

// NewSdCollateral creates a new instance of SdCollateral, bound to a specific deployed contract.
func NewSdCollateral(address common.Address, backend bind.ContractBackend) (*SdCollateral, error) {
	contract, err := bindSdCollateral(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &SdCollateral{SdCollateralCaller: SdCollateralCaller{contract: contract}, SdCollateralTransactor: SdCollateralTransactor{contract: contract}, SdCollateralFilterer: SdCollateralFilterer{contract: contract}}, nil
}

// NewSdCollateralCaller creates a new read-only instance of SdCollateral, bound to a specific deployed contract.
func NewSdCollateralCaller(address common.Address, caller bind.ContractCaller) (*SdCollateralCaller, error) {
	contract, err := bindSdCollateral(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &SdCollateralCaller{contract: contract}, nil
}

// NewSdCollateralTransactor creates a new write-only instance of SdCollateral, bound to a specific deployed contract.
func NewSdCollateralTransactor(address common.Address, transactor bind.ContractTransactor) (*SdCollateralTransactor, error) {
	contract, err := bindSdCollateral(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &SdCollateralTransactor{contract: contract}, nil
}

// NewSdCollateralFilterer creates a new log filterer instance of SdCollateral, bound to a specific deployed contract.
func NewSdCollateralFilterer(address common.Address, filterer bind.ContractFilterer) (*SdCollateralFilterer, error) {
	contract, err := bindSdCollateral(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &SdCollateralFilterer{contract: contract}, nil
}

// bindSdCollateral binds a generic wrapper to an already deployed contract.
func bindSdCollateral(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := SdCollateralMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SdCollateral *SdCollateralRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _SdCollateral.Contract.SdCollateralCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SdCollateral *SdCollateralRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SdCollateral.Contract.SdCollateralTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SdCollateral *SdCollateralRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SdCollateral.Contract.SdCollateralTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SdCollateral *SdCollateralCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _SdCollateral.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SdCollateral *SdCollateralTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SdCollateral.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SdCollateral *SdCollateralTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SdCollateral.Contract.contract.Transact(opts, method, params...)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_SdCollateral *SdCollateralCaller) DEFAULTADMINROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _SdCollateral.contract.Call(opts, &out, "DEFAULT_ADMIN_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_SdCollateral *SdCollateralSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _SdCollateral.Contract.DEFAULTADMINROLE(&_SdCollateral.CallOpts)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_SdCollateral *SdCollateralCallerSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _SdCollateral.Contract.DEFAULTADMINROLE(&_SdCollateral.CallOpts)
}

// ConvertETHToSD is a free data retrieval call binding the contract method 0xe614e17c.
//
// Solidity: function convertETHToSD(uint256 _ethAmount) view returns(uint256)
func (_SdCollateral *SdCollateralCaller) ConvertETHToSD(opts *bind.CallOpts, _ethAmount *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _SdCollateral.contract.Call(opts, &out, "convertETHToSD", _ethAmount)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ConvertETHToSD is a free data retrieval call binding the contract method 0xe614e17c.
//
// Solidity: function convertETHToSD(uint256 _ethAmount) view returns(uint256)
func (_SdCollateral *SdCollateralSession) ConvertETHToSD(_ethAmount *big.Int) (*big.Int, error) {
	return _SdCollateral.Contract.ConvertETHToSD(&_SdCollateral.CallOpts, _ethAmount)
}

// ConvertETHToSD is a free data retrieval call binding the contract method 0xe614e17c.
//
// Solidity: function convertETHToSD(uint256 _ethAmount) view returns(uint256)
func (_SdCollateral *SdCollateralCallerSession) ConvertETHToSD(_ethAmount *big.Int) (*big.Int, error) {
	return _SdCollateral.Contract.ConvertETHToSD(&_SdCollateral.CallOpts, _ethAmount)
}

// ConvertSDToETH is a free data retrieval call binding the contract method 0xdfdafccb.
//
// Solidity: function convertSDToETH(uint256 _sdAmount) view returns(uint256)
func (_SdCollateral *SdCollateralCaller) ConvertSDToETH(opts *bind.CallOpts, _sdAmount *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _SdCollateral.contract.Call(opts, &out, "convertSDToETH", _sdAmount)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ConvertSDToETH is a free data retrieval call binding the contract method 0xdfdafccb.
//
// Solidity: function convertSDToETH(uint256 _sdAmount) view returns(uint256)
func (_SdCollateral *SdCollateralSession) ConvertSDToETH(_sdAmount *big.Int) (*big.Int, error) {
	return _SdCollateral.Contract.ConvertSDToETH(&_SdCollateral.CallOpts, _sdAmount)
}

// ConvertSDToETH is a free data retrieval call binding the contract method 0xdfdafccb.
//
// Solidity: function convertSDToETH(uint256 _sdAmount) view returns(uint256)
func (_SdCollateral *SdCollateralCallerSession) ConvertSDToETH(_sdAmount *big.Int) (*big.Int, error) {
	return _SdCollateral.Contract.ConvertSDToETH(&_SdCollateral.CallOpts, _sdAmount)
}

// ConvertSDToShares is a free data retrieval call binding the contract method 0x2f0356e6.
//
// Solidity: function convertSDToShares(uint256 _sdAmount) view returns(uint256)
func (_SdCollateral *SdCollateralCaller) ConvertSDToShares(opts *bind.CallOpts, _sdAmount *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _SdCollateral.contract.Call(opts, &out, "convertSDToShares", _sdAmount)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ConvertSDToShares is a free data retrieval call binding the contract method 0x2f0356e6.
//
// Solidity: function convertSDToShares(uint256 _sdAmount) view returns(uint256)
func (_SdCollateral *SdCollateralSession) ConvertSDToShares(_sdAmount *big.Int) (*big.Int, error) {
	return _SdCollateral.Contract.ConvertSDToShares(&_SdCollateral.CallOpts, _sdAmount)
}

// ConvertSDToShares is a free data retrieval call binding the contract method 0x2f0356e6.
//
// Solidity: function convertSDToShares(uint256 _sdAmount) view returns(uint256)
func (_SdCollateral *SdCollateralCallerSession) ConvertSDToShares(_sdAmount *big.Int) (*big.Int, error) {
	return _SdCollateral.Contract.ConvertSDToShares(&_SdCollateral.CallOpts, _sdAmount)
}

// ConvertSharesToSD is a free data retrieval call binding the contract method 0xd4513719.
//
// Solidity: function convertSharesToSD(uint256 _numShares) view returns(uint256)
func (_SdCollateral *SdCollateralCaller) ConvertSharesToSD(opts *bind.CallOpts, _numShares *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _SdCollateral.contract.Call(opts, &out, "convertSharesToSD", _numShares)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ConvertSharesToSD is a free data retrieval call binding the contract method 0xd4513719.
//
// Solidity: function convertSharesToSD(uint256 _numShares) view returns(uint256)
func (_SdCollateral *SdCollateralSession) ConvertSharesToSD(_numShares *big.Int) (*big.Int, error) {
	return _SdCollateral.Contract.ConvertSharesToSD(&_SdCollateral.CallOpts, _numShares)
}

// ConvertSharesToSD is a free data retrieval call binding the contract method 0xd4513719.
//
// Solidity: function convertSharesToSD(uint256 _numShares) view returns(uint256)
func (_SdCollateral *SdCollateralCallerSession) ConvertSharesToSD(_numShares *big.Int) (*big.Int, error) {
	return _SdCollateral.Contract.ConvertSharesToSD(&_SdCollateral.CallOpts, _numShares)
}

// GetOperatorSDBalance is a free data retrieval call binding the contract method 0xbbd65e26.
//
// Solidity: function getOperatorSDBalance(address _operator) view returns(uint256)
func (_SdCollateral *SdCollateralCaller) GetOperatorSDBalance(opts *bind.CallOpts, _operator common.Address) (*big.Int, error) {
	var out []interface{}
	err := _SdCollateral.contract.Call(opts, &out, "getOperatorSDBalance", _operator)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetOperatorSDBalance is a free data retrieval call binding the contract method 0xbbd65e26.
//
// Solidity: function getOperatorSDBalance(address _operator) view returns(uint256)
func (_SdCollateral *SdCollateralSession) GetOperatorSDBalance(_operator common.Address) (*big.Int, error) {
	return _SdCollateral.Contract.GetOperatorSDBalance(&_SdCollateral.CallOpts, _operator)
}

// GetOperatorSDBalance is a free data retrieval call binding the contract method 0xbbd65e26.
//
// Solidity: function getOperatorSDBalance(address _operator) view returns(uint256)
func (_SdCollateral *SdCollateralCallerSession) GetOperatorSDBalance(_operator common.Address) (*big.Int, error) {
	return _SdCollateral.Contract.GetOperatorSDBalance(&_SdCollateral.CallOpts, _operator)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_SdCollateral *SdCollateralCaller) GetRoleAdmin(opts *bind.CallOpts, role [32]byte) ([32]byte, error) {
	var out []interface{}
	err := _SdCollateral.contract.Call(opts, &out, "getRoleAdmin", role)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_SdCollateral *SdCollateralSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _SdCollateral.Contract.GetRoleAdmin(&_SdCollateral.CallOpts, role)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_SdCollateral *SdCollateralCallerSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _SdCollateral.Contract.GetRoleAdmin(&_SdCollateral.CallOpts, role)
}

// HasEnoughSDCollateral is a free data retrieval call binding the contract method 0xbf137ba0.
//
// Solidity: function hasEnoughSDCollateral(address _operator, uint8 _poolId, uint32 _numValidators) view returns(bool)
func (_SdCollateral *SdCollateralCaller) HasEnoughSDCollateral(opts *bind.CallOpts, _operator common.Address, _poolId uint8, _numValidators uint32) (bool, error) {
	var out []interface{}
	err := _SdCollateral.contract.Call(opts, &out, "hasEnoughSDCollateral", _operator, _poolId, _numValidators)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// HasEnoughSDCollateral is a free data retrieval call binding the contract method 0xbf137ba0.
//
// Solidity: function hasEnoughSDCollateral(address _operator, uint8 _poolId, uint32 _numValidators) view returns(bool)
func (_SdCollateral *SdCollateralSession) HasEnoughSDCollateral(_operator common.Address, _poolId uint8, _numValidators uint32) (bool, error) {
	return _SdCollateral.Contract.HasEnoughSDCollateral(&_SdCollateral.CallOpts, _operator, _poolId, _numValidators)
}

// HasEnoughSDCollateral is a free data retrieval call binding the contract method 0xbf137ba0.
//
// Solidity: function hasEnoughSDCollateral(address _operator, uint8 _poolId, uint32 _numValidators) view returns(bool)
func (_SdCollateral *SdCollateralCallerSession) HasEnoughSDCollateral(_operator common.Address, _poolId uint8, _numValidators uint32) (bool, error) {
	return _SdCollateral.Contract.HasEnoughSDCollateral(&_SdCollateral.CallOpts, _operator, _poolId, _numValidators)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_SdCollateral *SdCollateralCaller) HasRole(opts *bind.CallOpts, role [32]byte, account common.Address) (bool, error) {
	var out []interface{}
	err := _SdCollateral.contract.Call(opts, &out, "hasRole", role, account)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_SdCollateral *SdCollateralSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _SdCollateral.Contract.HasRole(&_SdCollateral.CallOpts, role, account)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_SdCollateral *SdCollateralCallerSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _SdCollateral.Contract.HasRole(&_SdCollateral.CallOpts, role, account)
}

// OperatorShares is a free data retrieval call binding the contract method 0x371e21e9.
//
// Solidity: function operatorShares(address ) view returns(uint256)
func (_SdCollateral *SdCollateralCaller) OperatorShares(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _SdCollateral.contract.Call(opts, &out, "operatorShares", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// OperatorShares is a free data retrieval call binding the contract method 0x371e21e9.
//
// Solidity: function operatorShares(address ) view returns(uint256)
func (_SdCollateral *SdCollateralSession) OperatorShares(arg0 common.Address) (*big.Int, error) {
	return _SdCollateral.Contract.OperatorShares(&_SdCollateral.CallOpts, arg0)
}

// OperatorShares is a free data retrieval call binding the contract method 0x371e21e9.
//
// Solidity: function operatorShares(address ) view returns(uint256)
func (_SdCollateral *SdCollateralCallerSession) OperatorShares(arg0 common.Address) (*big.Int, error) {
	return _SdCollateral.Contract.OperatorShares(&_SdCollateral.CallOpts, arg0)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_SdCollateral *SdCollateralCaller) Paused(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _SdCollateral.contract.Call(opts, &out, "paused")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_SdCollateral *SdCollateralSession) Paused() (bool, error) {
	return _SdCollateral.Contract.Paused(&_SdCollateral.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_SdCollateral *SdCollateralCallerSession) Paused() (bool, error) {
	return _SdCollateral.Contract.Paused(&_SdCollateral.CallOpts)
}

// PoolThresholdbyPoolId is a free data retrieval call binding the contract method 0x8a9b3738.
//
// Solidity: function poolThresholdbyPoolId(uint8 ) view returns(uint256 minThreshold, uint256 withdrawThreshold, string units)
func (_SdCollateral *SdCollateralCaller) PoolThresholdbyPoolId(opts *bind.CallOpts, arg0 uint8) (struct {
	MinThreshold      *big.Int
	WithdrawThreshold *big.Int
	Units             string
}, error) {
	var out []interface{}
	err := _SdCollateral.contract.Call(opts, &out, "poolThresholdbyPoolId", arg0)

	outstruct := new(struct {
		MinThreshold      *big.Int
		WithdrawThreshold *big.Int
		Units             string
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.MinThreshold = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.WithdrawThreshold = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.Units = *abi.ConvertType(out[2], new(string)).(*string)

	return *outstruct, err

}

// PoolThresholdbyPoolId is a free data retrieval call binding the contract method 0x8a9b3738.
//
// Solidity: function poolThresholdbyPoolId(uint8 ) view returns(uint256 minThreshold, uint256 withdrawThreshold, string units)
func (_SdCollateral *SdCollateralSession) PoolThresholdbyPoolId(arg0 uint8) (struct {
	MinThreshold      *big.Int
	WithdrawThreshold *big.Int
	Units             string
}, error) {
	return _SdCollateral.Contract.PoolThresholdbyPoolId(&_SdCollateral.CallOpts, arg0)
}

// PoolThresholdbyPoolId is a free data retrieval call binding the contract method 0x8a9b3738.
//
// Solidity: function poolThresholdbyPoolId(uint8 ) view returns(uint256 minThreshold, uint256 withdrawThreshold, string units)
func (_SdCollateral *SdCollateralCallerSession) PoolThresholdbyPoolId(arg0 uint8) (struct {
	MinThreshold      *big.Int
	WithdrawThreshold *big.Int
	Units             string
}, error) {
	return _SdCollateral.Contract.PoolThresholdbyPoolId(&_SdCollateral.CallOpts, arg0)
}

// PriceFetcher is a free data retrieval call binding the contract method 0x2421d726.
//
// Solidity: function priceFetcher() view returns(address)
func (_SdCollateral *SdCollateralCaller) PriceFetcher(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _SdCollateral.contract.Call(opts, &out, "priceFetcher")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// PriceFetcher is a free data retrieval call binding the contract method 0x2421d726.
//
// Solidity: function priceFetcher() view returns(address)
func (_SdCollateral *SdCollateralSession) PriceFetcher() (common.Address, error) {
	return _SdCollateral.Contract.PriceFetcher(&_SdCollateral.CallOpts)
}

// PriceFetcher is a free data retrieval call binding the contract method 0x2421d726.
//
// Solidity: function priceFetcher() view returns(address)
func (_SdCollateral *SdCollateralCallerSession) PriceFetcher() (common.Address, error) {
	return _SdCollateral.Contract.PriceFetcher(&_SdCollateral.CallOpts)
}

// SdERC20 is a free data retrieval call binding the contract method 0xf493540a.
//
// Solidity: function sdERC20() view returns(address)
func (_SdCollateral *SdCollateralCaller) SdERC20(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _SdCollateral.contract.Call(opts, &out, "sdERC20")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// SdERC20 is a free data retrieval call binding the contract method 0xf493540a.
//
// Solidity: function sdERC20() view returns(address)
func (_SdCollateral *SdCollateralSession) SdERC20() (common.Address, error) {
	return _SdCollateral.Contract.SdERC20(&_SdCollateral.CallOpts)
}

// SdERC20 is a free data retrieval call binding the contract method 0xf493540a.
//
// Solidity: function sdERC20() view returns(address)
func (_SdCollateral *SdCollateralCallerSession) SdERC20() (common.Address, error) {
	return _SdCollateral.Contract.SdERC20(&_SdCollateral.CallOpts)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_SdCollateral *SdCollateralCaller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _SdCollateral.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_SdCollateral *SdCollateralSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _SdCollateral.Contract.SupportsInterface(&_SdCollateral.CallOpts, interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_SdCollateral *SdCollateralCallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _SdCollateral.Contract.SupportsInterface(&_SdCollateral.CallOpts, interfaceId)
}

// TotalSDCollateral is a free data retrieval call binding the contract method 0x33869a2c.
//
// Solidity: function totalSDCollateral() view returns(uint256)
func (_SdCollateral *SdCollateralCaller) TotalSDCollateral(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _SdCollateral.contract.Call(opts, &out, "totalSDCollateral")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSDCollateral is a free data retrieval call binding the contract method 0x33869a2c.
//
// Solidity: function totalSDCollateral() view returns(uint256)
func (_SdCollateral *SdCollateralSession) TotalSDCollateral() (*big.Int, error) {
	return _SdCollateral.Contract.TotalSDCollateral(&_SdCollateral.CallOpts)
}

// TotalSDCollateral is a free data retrieval call binding the contract method 0x33869a2c.
//
// Solidity: function totalSDCollateral() view returns(uint256)
func (_SdCollateral *SdCollateralCallerSession) TotalSDCollateral() (*big.Int, error) {
	return _SdCollateral.Contract.TotalSDCollateral(&_SdCollateral.CallOpts)
}

// TotalShares is a free data retrieval call binding the contract method 0x3a98ef39.
//
// Solidity: function totalShares() view returns(uint256)
func (_SdCollateral *SdCollateralCaller) TotalShares(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _SdCollateral.contract.Call(opts, &out, "totalShares")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalShares is a free data retrieval call binding the contract method 0x3a98ef39.
//
// Solidity: function totalShares() view returns(uint256)
func (_SdCollateral *SdCollateralSession) TotalShares() (*big.Int, error) {
	return _SdCollateral.Contract.TotalShares(&_SdCollateral.CallOpts)
}

// TotalShares is a free data retrieval call binding the contract method 0x3a98ef39.
//
// Solidity: function totalShares() view returns(uint256)
func (_SdCollateral *SdCollateralCallerSession) TotalShares() (*big.Int, error) {
	return _SdCollateral.Contract.TotalShares(&_SdCollateral.CallOpts)
}

// DepositSDAsCollateral is a paid mutator transaction binding the contract method 0xfcb7e032.
//
// Solidity: function depositSDAsCollateral(uint256 _sdAmount) returns()
func (_SdCollateral *SdCollateralTransactor) DepositSDAsCollateral(opts *bind.TransactOpts, _sdAmount *big.Int) (*types.Transaction, error) {
	return _SdCollateral.contract.Transact(opts, "depositSDAsCollateral", _sdAmount)
}

// DepositSDAsCollateral is a paid mutator transaction binding the contract method 0xfcb7e032.
//
// Solidity: function depositSDAsCollateral(uint256 _sdAmount) returns()
func (_SdCollateral *SdCollateralSession) DepositSDAsCollateral(_sdAmount *big.Int) (*types.Transaction, error) {
	return _SdCollateral.Contract.DepositSDAsCollateral(&_SdCollateral.TransactOpts, _sdAmount)
}

// DepositSDAsCollateral is a paid mutator transaction binding the contract method 0xfcb7e032.
//
// Solidity: function depositSDAsCollateral(uint256 _sdAmount) returns()
func (_SdCollateral *SdCollateralTransactorSession) DepositSDAsCollateral(_sdAmount *big.Int) (*types.Transaction, error) {
	return _SdCollateral.Contract.DepositSDAsCollateral(&_SdCollateral.TransactOpts, _sdAmount)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_SdCollateral *SdCollateralTransactor) GrantRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _SdCollateral.contract.Transact(opts, "grantRole", role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_SdCollateral *SdCollateralSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _SdCollateral.Contract.GrantRole(&_SdCollateral.TransactOpts, role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_SdCollateral *SdCollateralTransactorSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _SdCollateral.Contract.GrantRole(&_SdCollateral.TransactOpts, role, account)
}

// Initialize is a paid mutator transaction binding the contract method 0xc0c53b8b.
//
// Solidity: function initialize(address _admin, address _sdERC20Addr, address _priceFetcherAddr) returns()
func (_SdCollateral *SdCollateralTransactor) Initialize(opts *bind.TransactOpts, _admin common.Address, _sdERC20Addr common.Address, _priceFetcherAddr common.Address) (*types.Transaction, error) {
	return _SdCollateral.contract.Transact(opts, "initialize", _admin, _sdERC20Addr, _priceFetcherAddr)
}

// Initialize is a paid mutator transaction binding the contract method 0xc0c53b8b.
//
// Solidity: function initialize(address _admin, address _sdERC20Addr, address _priceFetcherAddr) returns()
func (_SdCollateral *SdCollateralSession) Initialize(_admin common.Address, _sdERC20Addr common.Address, _priceFetcherAddr common.Address) (*types.Transaction, error) {
	return _SdCollateral.Contract.Initialize(&_SdCollateral.TransactOpts, _admin, _sdERC20Addr, _priceFetcherAddr)
}

// Initialize is a paid mutator transaction binding the contract method 0xc0c53b8b.
//
// Solidity: function initialize(address _admin, address _sdERC20Addr, address _priceFetcherAddr) returns()
func (_SdCollateral *SdCollateralTransactorSession) Initialize(_admin common.Address, _sdERC20Addr common.Address, _priceFetcherAddr common.Address) (*types.Transaction, error) {
	return _SdCollateral.Contract.Initialize(&_SdCollateral.TransactOpts, _admin, _sdERC20Addr, _priceFetcherAddr)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address account) returns()
func (_SdCollateral *SdCollateralTransactor) RenounceRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _SdCollateral.contract.Transact(opts, "renounceRole", role, account)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address account) returns()
func (_SdCollateral *SdCollateralSession) RenounceRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _SdCollateral.Contract.RenounceRole(&_SdCollateral.TransactOpts, role, account)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address account) returns()
func (_SdCollateral *SdCollateralTransactorSession) RenounceRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _SdCollateral.Contract.RenounceRole(&_SdCollateral.TransactOpts, role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_SdCollateral *SdCollateralTransactor) RevokeRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _SdCollateral.contract.Transact(opts, "revokeRole", role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_SdCollateral *SdCollateralSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _SdCollateral.Contract.RevokeRole(&_SdCollateral.TransactOpts, role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_SdCollateral *SdCollateralTransactorSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _SdCollateral.Contract.RevokeRole(&_SdCollateral.TransactOpts, role, account)
}

// UpdatePoolThreshold is a paid mutator transaction binding the contract method 0x401bd27a.
//
// Solidity: function updatePoolThreshold(uint8 _poolId, uint256 _minThreshold, uint256 _withdrawThreshold, string _units) returns()
func (_SdCollateral *SdCollateralTransactor) UpdatePoolThreshold(opts *bind.TransactOpts, _poolId uint8, _minThreshold *big.Int, _withdrawThreshold *big.Int, _units string) (*types.Transaction, error) {
	return _SdCollateral.contract.Transact(opts, "updatePoolThreshold", _poolId, _minThreshold, _withdrawThreshold, _units)
}

// UpdatePoolThreshold is a paid mutator transaction binding the contract method 0x401bd27a.
//
// Solidity: function updatePoolThreshold(uint8 _poolId, uint256 _minThreshold, uint256 _withdrawThreshold, string _units) returns()
func (_SdCollateral *SdCollateralSession) UpdatePoolThreshold(_poolId uint8, _minThreshold *big.Int, _withdrawThreshold *big.Int, _units string) (*types.Transaction, error) {
	return _SdCollateral.Contract.UpdatePoolThreshold(&_SdCollateral.TransactOpts, _poolId, _minThreshold, _withdrawThreshold, _units)
}

// UpdatePoolThreshold is a paid mutator transaction binding the contract method 0x401bd27a.
//
// Solidity: function updatePoolThreshold(uint8 _poolId, uint256 _minThreshold, uint256 _withdrawThreshold, string _units) returns()
func (_SdCollateral *SdCollateralTransactorSession) UpdatePoolThreshold(_poolId uint8, _minThreshold *big.Int, _withdrawThreshold *big.Int, _units string) (*types.Transaction, error) {
	return _SdCollateral.Contract.UpdatePoolThreshold(&_SdCollateral.TransactOpts, _poolId, _minThreshold, _withdrawThreshold, _units)
}

// Withdraw is a paid mutator transaction binding the contract method 0x2e1a7d4d.
//
// Solidity: function withdraw(uint256 _requestedSD) returns()
func (_SdCollateral *SdCollateralTransactor) Withdraw(opts *bind.TransactOpts, _requestedSD *big.Int) (*types.Transaction, error) {
	return _SdCollateral.contract.Transact(opts, "withdraw", _requestedSD)
}

// Withdraw is a paid mutator transaction binding the contract method 0x2e1a7d4d.
//
// Solidity: function withdraw(uint256 _requestedSD) returns()
func (_SdCollateral *SdCollateralSession) Withdraw(_requestedSD *big.Int) (*types.Transaction, error) {
	return _SdCollateral.Contract.Withdraw(&_SdCollateral.TransactOpts, _requestedSD)
}

// Withdraw is a paid mutator transaction binding the contract method 0x2e1a7d4d.
//
// Solidity: function withdraw(uint256 _requestedSD) returns()
func (_SdCollateral *SdCollateralTransactorSession) Withdraw(_requestedSD *big.Int) (*types.Transaction, error) {
	return _SdCollateral.Contract.Withdraw(&_SdCollateral.TransactOpts, _requestedSD)
}

// SdCollateralInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the SdCollateral contract.
type SdCollateralInitializedIterator struct {
	Event *SdCollateralInitialized // Event containing the contract specifics and raw log

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
func (it *SdCollateralInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SdCollateralInitialized)
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
		it.Event = new(SdCollateralInitialized)
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
func (it *SdCollateralInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SdCollateralInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SdCollateralInitialized represents a Initialized event raised by the SdCollateral contract.
type SdCollateralInitialized struct {
	Version uint8
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_SdCollateral *SdCollateralFilterer) FilterInitialized(opts *bind.FilterOpts) (*SdCollateralInitializedIterator, error) {

	logs, sub, err := _SdCollateral.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &SdCollateralInitializedIterator{contract: _SdCollateral.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_SdCollateral *SdCollateralFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *SdCollateralInitialized) (event.Subscription, error) {

	logs, sub, err := _SdCollateral.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SdCollateralInitialized)
				if err := _SdCollateral.contract.UnpackLog(event, "Initialized", log); err != nil {
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
func (_SdCollateral *SdCollateralFilterer) ParseInitialized(log types.Log) (*SdCollateralInitialized, error) {
	event := new(SdCollateralInitialized)
	if err := _SdCollateral.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SdCollateralPausedIterator is returned from FilterPaused and is used to iterate over the raw logs and unpacked data for Paused events raised by the SdCollateral contract.
type SdCollateralPausedIterator struct {
	Event *SdCollateralPaused // Event containing the contract specifics and raw log

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
func (it *SdCollateralPausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SdCollateralPaused)
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
		it.Event = new(SdCollateralPaused)
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
func (it *SdCollateralPausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SdCollateralPausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SdCollateralPaused represents a Paused event raised by the SdCollateral contract.
type SdCollateralPaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterPaused is a free log retrieval operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_SdCollateral *SdCollateralFilterer) FilterPaused(opts *bind.FilterOpts) (*SdCollateralPausedIterator, error) {

	logs, sub, err := _SdCollateral.contract.FilterLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return &SdCollateralPausedIterator{contract: _SdCollateral.contract, event: "Paused", logs: logs, sub: sub}, nil
}

// WatchPaused is a free log subscription operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_SdCollateral *SdCollateralFilterer) WatchPaused(opts *bind.WatchOpts, sink chan<- *SdCollateralPaused) (event.Subscription, error) {

	logs, sub, err := _SdCollateral.contract.WatchLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SdCollateralPaused)
				if err := _SdCollateral.contract.UnpackLog(event, "Paused", log); err != nil {
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
func (_SdCollateral *SdCollateralFilterer) ParsePaused(log types.Log) (*SdCollateralPaused, error) {
	event := new(SdCollateralPaused)
	if err := _SdCollateral.contract.UnpackLog(event, "Paused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SdCollateralRoleAdminChangedIterator is returned from FilterRoleAdminChanged and is used to iterate over the raw logs and unpacked data for RoleAdminChanged events raised by the SdCollateral contract.
type SdCollateralRoleAdminChangedIterator struct {
	Event *SdCollateralRoleAdminChanged // Event containing the contract specifics and raw log

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
func (it *SdCollateralRoleAdminChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SdCollateralRoleAdminChanged)
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
		it.Event = new(SdCollateralRoleAdminChanged)
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
func (it *SdCollateralRoleAdminChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SdCollateralRoleAdminChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SdCollateralRoleAdminChanged represents a RoleAdminChanged event raised by the SdCollateral contract.
type SdCollateralRoleAdminChanged struct {
	Role              [32]byte
	PreviousAdminRole [32]byte
	NewAdminRole      [32]byte
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterRoleAdminChanged is a free log retrieval operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_SdCollateral *SdCollateralFilterer) FilterRoleAdminChanged(opts *bind.FilterOpts, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (*SdCollateralRoleAdminChangedIterator, error) {

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

	logs, sub, err := _SdCollateral.contract.FilterLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return &SdCollateralRoleAdminChangedIterator{contract: _SdCollateral.contract, event: "RoleAdminChanged", logs: logs, sub: sub}, nil
}

// WatchRoleAdminChanged is a free log subscription operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_SdCollateral *SdCollateralFilterer) WatchRoleAdminChanged(opts *bind.WatchOpts, sink chan<- *SdCollateralRoleAdminChanged, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (event.Subscription, error) {

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

	logs, sub, err := _SdCollateral.contract.WatchLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SdCollateralRoleAdminChanged)
				if err := _SdCollateral.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
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
func (_SdCollateral *SdCollateralFilterer) ParseRoleAdminChanged(log types.Log) (*SdCollateralRoleAdminChanged, error) {
	event := new(SdCollateralRoleAdminChanged)
	if err := _SdCollateral.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SdCollateralRoleGrantedIterator is returned from FilterRoleGranted and is used to iterate over the raw logs and unpacked data for RoleGranted events raised by the SdCollateral contract.
type SdCollateralRoleGrantedIterator struct {
	Event *SdCollateralRoleGranted // Event containing the contract specifics and raw log

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
func (it *SdCollateralRoleGrantedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SdCollateralRoleGranted)
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
		it.Event = new(SdCollateralRoleGranted)
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
func (it *SdCollateralRoleGrantedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SdCollateralRoleGrantedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SdCollateralRoleGranted represents a RoleGranted event raised by the SdCollateral contract.
type SdCollateralRoleGranted struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleGranted is a free log retrieval operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_SdCollateral *SdCollateralFilterer) FilterRoleGranted(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*SdCollateralRoleGrantedIterator, error) {

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

	logs, sub, err := _SdCollateral.contract.FilterLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &SdCollateralRoleGrantedIterator{contract: _SdCollateral.contract, event: "RoleGranted", logs: logs, sub: sub}, nil
}

// WatchRoleGranted is a free log subscription operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_SdCollateral *SdCollateralFilterer) WatchRoleGranted(opts *bind.WatchOpts, sink chan<- *SdCollateralRoleGranted, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _SdCollateral.contract.WatchLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SdCollateralRoleGranted)
				if err := _SdCollateral.contract.UnpackLog(event, "RoleGranted", log); err != nil {
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
func (_SdCollateral *SdCollateralFilterer) ParseRoleGranted(log types.Log) (*SdCollateralRoleGranted, error) {
	event := new(SdCollateralRoleGranted)
	if err := _SdCollateral.contract.UnpackLog(event, "RoleGranted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SdCollateralRoleRevokedIterator is returned from FilterRoleRevoked and is used to iterate over the raw logs and unpacked data for RoleRevoked events raised by the SdCollateral contract.
type SdCollateralRoleRevokedIterator struct {
	Event *SdCollateralRoleRevoked // Event containing the contract specifics and raw log

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
func (it *SdCollateralRoleRevokedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SdCollateralRoleRevoked)
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
		it.Event = new(SdCollateralRoleRevoked)
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
func (it *SdCollateralRoleRevokedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SdCollateralRoleRevokedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SdCollateralRoleRevoked represents a RoleRevoked event raised by the SdCollateral contract.
type SdCollateralRoleRevoked struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleRevoked is a free log retrieval operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_SdCollateral *SdCollateralFilterer) FilterRoleRevoked(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*SdCollateralRoleRevokedIterator, error) {

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

	logs, sub, err := _SdCollateral.contract.FilterLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &SdCollateralRoleRevokedIterator{contract: _SdCollateral.contract, event: "RoleRevoked", logs: logs, sub: sub}, nil
}

// WatchRoleRevoked is a free log subscription operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_SdCollateral *SdCollateralFilterer) WatchRoleRevoked(opts *bind.WatchOpts, sink chan<- *SdCollateralRoleRevoked, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _SdCollateral.contract.WatchLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SdCollateralRoleRevoked)
				if err := _SdCollateral.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
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
func (_SdCollateral *SdCollateralFilterer) ParseRoleRevoked(log types.Log) (*SdCollateralRoleRevoked, error) {
	event := new(SdCollateralRoleRevoked)
	if err := _SdCollateral.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SdCollateralUnpausedIterator is returned from FilterUnpaused and is used to iterate over the raw logs and unpacked data for Unpaused events raised by the SdCollateral contract.
type SdCollateralUnpausedIterator struct {
	Event *SdCollateralUnpaused // Event containing the contract specifics and raw log

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
func (it *SdCollateralUnpausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SdCollateralUnpaused)
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
		it.Event = new(SdCollateralUnpaused)
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
func (it *SdCollateralUnpausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SdCollateralUnpausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SdCollateralUnpaused represents a Unpaused event raised by the SdCollateral contract.
type SdCollateralUnpaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterUnpaused is a free log retrieval operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_SdCollateral *SdCollateralFilterer) FilterUnpaused(opts *bind.FilterOpts) (*SdCollateralUnpausedIterator, error) {

	logs, sub, err := _SdCollateral.contract.FilterLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return &SdCollateralUnpausedIterator{contract: _SdCollateral.contract, event: "Unpaused", logs: logs, sub: sub}, nil
}

// WatchUnpaused is a free log subscription operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_SdCollateral *SdCollateralFilterer) WatchUnpaused(opts *bind.WatchOpts, sink chan<- *SdCollateralUnpaused) (event.Subscription, error) {

	logs, sub, err := _SdCollateral.contract.WatchLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SdCollateralUnpaused)
				if err := _SdCollateral.contract.UnpackLog(event, "Unpaused", log); err != nil {
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
func (_SdCollateral *SdCollateralFilterer) ParseUnpaused(log types.Log) (*SdCollateralUnpaused, error) {
	event := new(SdCollateralUnpaused)
	if err := _SdCollateral.contract.UnpackLog(event, "Unpaused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
