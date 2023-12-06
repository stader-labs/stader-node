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
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"CallerNotManager\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"CallerNotStaderContract\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"CallerNotWithdrawVault\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"operatorSDCollateral\",\"type\":\"uint256\"}],\"name\":\"InsufficientSDToWithdraw\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidPoolId\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidPoolLimit\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NoStateChange\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"SDTransferFailed\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ZeroAddress\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"version\",\"type\":\"uint8\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"sdAmount\",\"type\":\"uint256\"}],\"name\":\"ReducedUtilizedPosition\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"previousAdminRole\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"newAdminRole\",\"type\":\"bytes32\"}],\"name\":\"RoleAdminChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleGranted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleRevoked\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"sdAmount\",\"type\":\"uint256\"}],\"name\":\"SDDeposited\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"repayAmount\",\"type\":\"uint256\"}],\"name\":\"SDRepaid\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"auction\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"sdSlashed\",\"type\":\"uint256\"}],\"name\":\"SDSlashed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"sdSlashFromUtilized\",\"type\":\"uint256\"}],\"name\":\"SDSlashedFromUtilize\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"sdAmount\",\"type\":\"uint256\"}],\"name\":\"SDWithdrawn\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"poolId\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"}],\"name\":\"UpdatedPoolIdForOperator\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"poolId\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"minThreshold\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"withdrawThreshold\",\"type\":\"uint256\"}],\"name\":\"UpdatedPoolThreshold\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"staderConfig\",\"type\":\"address\"}],\"name\":\"UpdatedStaderConfig\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"sdAmount\",\"type\":\"uint256\"}],\"name\":\"UtilizedSDDeposited\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"DEFAULT_ADMIN_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_ethAmount\",\"type\":\"uint256\"}],\"name\":\"convertETHToSD\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_sdAmount\",\"type\":\"uint256\"}],\"name\":\"convertSDToETH\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_sdAmount\",\"type\":\"uint256\"}],\"name\":\"depositSDAsCollateral\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_operator\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_sdAmount\",\"type\":\"uint256\"}],\"name\":\"depositSDAsCollateralOnBehalf\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_operator\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_sdAmount\",\"type\":\"uint256\"}],\"name\":\"depositUtilizedSD\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"_poolId\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"_numValidator\",\"type\":\"uint256\"}],\"name\":\"getMinimumSDToBond\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"_minSDToBond\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_operator\",\"type\":\"address\"}],\"name\":\"getOperatorInfo\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"_poolId\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"_operatorId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_validatorCount\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_operator\",\"type\":\"address\"}],\"name\":\"getOperatorWithdrawThreshold\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"operatorWithdrawThreshold\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_operator\",\"type\":\"address\"},{\"internalType\":\"uint8\",\"name\":\"_poolId\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"_numValidator\",\"type\":\"uint256\"}],\"name\":\"getRemainingSDToBond\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_operator\",\"type\":\"address\"}],\"name\":\"getRewardEligibleSD\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"_rewardEligibleSD\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"}],\"name\":\"getRoleAdmin\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"grantRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_operator\",\"type\":\"address\"},{\"internalType\":\"uint8\",\"name\":\"_poolId\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"_numValidator\",\"type\":\"uint256\"}],\"name\":\"hasEnoughSDCollateral\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"hasRole\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_admin\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_staderConfig\",\"type\":\"address\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"maxApproveSD\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"operatorSDBalance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"operatorUtilizedSDBalance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"name\":\"poolThresholdbyPoolId\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"minThreshold\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxThreshold\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"withdrawThreshold\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"units\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_operator\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_sdAmount\",\"type\":\"uint256\"}],\"name\":\"reduceUtilizedSDPosition\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"renounceRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"revokeRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_validatorId\",\"type\":\"uint256\"},{\"internalType\":\"uint8\",\"name\":\"_poolId\",\"type\":\"uint8\"}],\"name\":\"slashValidatorSD\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"staderConfig\",\"outputs\":[{\"internalType\":\"contractIStaderConfig\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"_poolId\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"_minThreshold\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_maxThreshold\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_withdrawThreshold\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"_units\",\"type\":\"string\"}],\"name\":\"updatePoolThreshold\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_staderConfig\",\"type\":\"address\"}],\"name\":\"updateStaderConfig\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_requestedSD\",\"type\":\"uint256\"}],\"name\":\"withdraw\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
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

// GetMinimumSDToBond is a free data retrieval call binding the contract method 0x379b727e.
//
// Solidity: function getMinimumSDToBond(uint8 _poolId, uint256 _numValidator) view returns(uint256 _minSDToBond)
func (_SdCollateral *SdCollateralCaller) GetMinimumSDToBond(opts *bind.CallOpts, _poolId uint8, _numValidator *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _SdCollateral.contract.Call(opts, &out, "getMinimumSDToBond", _poolId, _numValidator)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetMinimumSDToBond is a free data retrieval call binding the contract method 0x379b727e.
//
// Solidity: function getMinimumSDToBond(uint8 _poolId, uint256 _numValidator) view returns(uint256 _minSDToBond)
func (_SdCollateral *SdCollateralSession) GetMinimumSDToBond(_poolId uint8, _numValidator *big.Int) (*big.Int, error) {
	return _SdCollateral.Contract.GetMinimumSDToBond(&_SdCollateral.CallOpts, _poolId, _numValidator)
}

// GetMinimumSDToBond is a free data retrieval call binding the contract method 0x379b727e.
//
// Solidity: function getMinimumSDToBond(uint8 _poolId, uint256 _numValidator) view returns(uint256 _minSDToBond)
func (_SdCollateral *SdCollateralCallerSession) GetMinimumSDToBond(_poolId uint8, _numValidator *big.Int) (*big.Int, error) {
	return _SdCollateral.Contract.GetMinimumSDToBond(&_SdCollateral.CallOpts, _poolId, _numValidator)
}

// GetOperatorInfo is a free data retrieval call binding the contract method 0x27d9ab5d.
//
// Solidity: function getOperatorInfo(address _operator) view returns(uint8 _poolId, uint256 _operatorId, uint256 _validatorCount)
func (_SdCollateral *SdCollateralCaller) GetOperatorInfo(opts *bind.CallOpts, _operator common.Address) (struct {
	PoolId         uint8
	OperatorId     *big.Int
	ValidatorCount *big.Int
}, error) {
	var out []interface{}
	err := _SdCollateral.contract.Call(opts, &out, "getOperatorInfo", _operator)

	outstruct := new(struct {
		PoolId         uint8
		OperatorId     *big.Int
		ValidatorCount *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.PoolId = *abi.ConvertType(out[0], new(uint8)).(*uint8)
	outstruct.OperatorId = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.ValidatorCount = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// GetOperatorInfo is a free data retrieval call binding the contract method 0x27d9ab5d.
//
// Solidity: function getOperatorInfo(address _operator) view returns(uint8 _poolId, uint256 _operatorId, uint256 _validatorCount)
func (_SdCollateral *SdCollateralSession) GetOperatorInfo(_operator common.Address) (struct {
	PoolId         uint8
	OperatorId     *big.Int
	ValidatorCount *big.Int
}, error) {
	return _SdCollateral.Contract.GetOperatorInfo(&_SdCollateral.CallOpts, _operator)
}

// GetOperatorInfo is a free data retrieval call binding the contract method 0x27d9ab5d.
//
// Solidity: function getOperatorInfo(address _operator) view returns(uint8 _poolId, uint256 _operatorId, uint256 _validatorCount)
func (_SdCollateral *SdCollateralCallerSession) GetOperatorInfo(_operator common.Address) (struct {
	PoolId         uint8
	OperatorId     *big.Int
	ValidatorCount *big.Int
}, error) {
	return _SdCollateral.Contract.GetOperatorInfo(&_SdCollateral.CallOpts, _operator)
}

// GetOperatorWithdrawThreshold is a free data retrieval call binding the contract method 0x9871a30a.
//
// Solidity: function getOperatorWithdrawThreshold(address _operator) view returns(uint256 operatorWithdrawThreshold)
func (_SdCollateral *SdCollateralCaller) GetOperatorWithdrawThreshold(opts *bind.CallOpts, _operator common.Address) (*big.Int, error) {
	var out []interface{}
	err := _SdCollateral.contract.Call(opts, &out, "getOperatorWithdrawThreshold", _operator)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetOperatorWithdrawThreshold is a free data retrieval call binding the contract method 0x9871a30a.
//
// Solidity: function getOperatorWithdrawThreshold(address _operator) view returns(uint256 operatorWithdrawThreshold)
func (_SdCollateral *SdCollateralSession) GetOperatorWithdrawThreshold(_operator common.Address) (*big.Int, error) {
	return _SdCollateral.Contract.GetOperatorWithdrawThreshold(&_SdCollateral.CallOpts, _operator)
}

// GetOperatorWithdrawThreshold is a free data retrieval call binding the contract method 0x9871a30a.
//
// Solidity: function getOperatorWithdrawThreshold(address _operator) view returns(uint256 operatorWithdrawThreshold)
func (_SdCollateral *SdCollateralCallerSession) GetOperatorWithdrawThreshold(_operator common.Address) (*big.Int, error) {
	return _SdCollateral.Contract.GetOperatorWithdrawThreshold(&_SdCollateral.CallOpts, _operator)
}

// GetRemainingSDToBond is a free data retrieval call binding the contract method 0x351691ab.
//
// Solidity: function getRemainingSDToBond(address _operator, uint8 _poolId, uint256 _numValidator) view returns(uint256)
func (_SdCollateral *SdCollateralCaller) GetRemainingSDToBond(opts *bind.CallOpts, _operator common.Address, _poolId uint8, _numValidator *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _SdCollateral.contract.Call(opts, &out, "getRemainingSDToBond", _operator, _poolId, _numValidator)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetRemainingSDToBond is a free data retrieval call binding the contract method 0x351691ab.
//
// Solidity: function getRemainingSDToBond(address _operator, uint8 _poolId, uint256 _numValidator) view returns(uint256)
func (_SdCollateral *SdCollateralSession) GetRemainingSDToBond(_operator common.Address, _poolId uint8, _numValidator *big.Int) (*big.Int, error) {
	return _SdCollateral.Contract.GetRemainingSDToBond(&_SdCollateral.CallOpts, _operator, _poolId, _numValidator)
}

// GetRemainingSDToBond is a free data retrieval call binding the contract method 0x351691ab.
//
// Solidity: function getRemainingSDToBond(address _operator, uint8 _poolId, uint256 _numValidator) view returns(uint256)
func (_SdCollateral *SdCollateralCallerSession) GetRemainingSDToBond(_operator common.Address, _poolId uint8, _numValidator *big.Int) (*big.Int, error) {
	return _SdCollateral.Contract.GetRemainingSDToBond(&_SdCollateral.CallOpts, _operator, _poolId, _numValidator)
}

// GetRewardEligibleSD is a free data retrieval call binding the contract method 0x3909afd3.
//
// Solidity: function getRewardEligibleSD(address _operator) view returns(uint256 _rewardEligibleSD)
func (_SdCollateral *SdCollateralCaller) GetRewardEligibleSD(opts *bind.CallOpts, _operator common.Address) (*big.Int, error) {
	var out []interface{}
	err := _SdCollateral.contract.Call(opts, &out, "getRewardEligibleSD", _operator)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetRewardEligibleSD is a free data retrieval call binding the contract method 0x3909afd3.
//
// Solidity: function getRewardEligibleSD(address _operator) view returns(uint256 _rewardEligibleSD)
func (_SdCollateral *SdCollateralSession) GetRewardEligibleSD(_operator common.Address) (*big.Int, error) {
	return _SdCollateral.Contract.GetRewardEligibleSD(&_SdCollateral.CallOpts, _operator)
}

// GetRewardEligibleSD is a free data retrieval call binding the contract method 0x3909afd3.
//
// Solidity: function getRewardEligibleSD(address _operator) view returns(uint256 _rewardEligibleSD)
func (_SdCollateral *SdCollateralCallerSession) GetRewardEligibleSD(_operator common.Address) (*big.Int, error) {
	return _SdCollateral.Contract.GetRewardEligibleSD(&_SdCollateral.CallOpts, _operator)
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

// HasEnoughSDCollateral is a free data retrieval call binding the contract method 0xb178e38e.
//
// Solidity: function hasEnoughSDCollateral(address _operator, uint8 _poolId, uint256 _numValidator) view returns(bool)
func (_SdCollateral *SdCollateralCaller) HasEnoughSDCollateral(opts *bind.CallOpts, _operator common.Address, _poolId uint8, _numValidator *big.Int) (bool, error) {
	var out []interface{}
	err := _SdCollateral.contract.Call(opts, &out, "hasEnoughSDCollateral", _operator, _poolId, _numValidator)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// HasEnoughSDCollateral is a free data retrieval call binding the contract method 0xb178e38e.
//
// Solidity: function hasEnoughSDCollateral(address _operator, uint8 _poolId, uint256 _numValidator) view returns(bool)
func (_SdCollateral *SdCollateralSession) HasEnoughSDCollateral(_operator common.Address, _poolId uint8, _numValidator *big.Int) (bool, error) {
	return _SdCollateral.Contract.HasEnoughSDCollateral(&_SdCollateral.CallOpts, _operator, _poolId, _numValidator)
}

// HasEnoughSDCollateral is a free data retrieval call binding the contract method 0xb178e38e.
//
// Solidity: function hasEnoughSDCollateral(address _operator, uint8 _poolId, uint256 _numValidator) view returns(bool)
func (_SdCollateral *SdCollateralCallerSession) HasEnoughSDCollateral(_operator common.Address, _poolId uint8, _numValidator *big.Int) (bool, error) {
	return _SdCollateral.Contract.HasEnoughSDCollateral(&_SdCollateral.CallOpts, _operator, _poolId, _numValidator)
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

// OperatorSDBalance is a free data retrieval call binding the contract method 0xf9af40b8.
//
// Solidity: function operatorSDBalance(address ) view returns(uint256)
func (_SdCollateral *SdCollateralCaller) OperatorSDBalance(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _SdCollateral.contract.Call(opts, &out, "operatorSDBalance", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// OperatorSDBalance is a free data retrieval call binding the contract method 0xf9af40b8.
//
// Solidity: function operatorSDBalance(address ) view returns(uint256)
func (_SdCollateral *SdCollateralSession) OperatorSDBalance(arg0 common.Address) (*big.Int, error) {
	return _SdCollateral.Contract.OperatorSDBalance(&_SdCollateral.CallOpts, arg0)
}

// OperatorSDBalance is a free data retrieval call binding the contract method 0xf9af40b8.
//
// Solidity: function operatorSDBalance(address ) view returns(uint256)
func (_SdCollateral *SdCollateralCallerSession) OperatorSDBalance(arg0 common.Address) (*big.Int, error) {
	return _SdCollateral.Contract.OperatorSDBalance(&_SdCollateral.CallOpts, arg0)
}

// OperatorUtilizedSDBalance is a free data retrieval call binding the contract method 0xb11a3a9b.
//
// Solidity: function operatorUtilizedSDBalance(address ) view returns(uint256)
func (_SdCollateral *SdCollateralCaller) OperatorUtilizedSDBalance(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _SdCollateral.contract.Call(opts, &out, "operatorUtilizedSDBalance", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// OperatorUtilizedSDBalance is a free data retrieval call binding the contract method 0xb11a3a9b.
//
// Solidity: function operatorUtilizedSDBalance(address ) view returns(uint256)
func (_SdCollateral *SdCollateralSession) OperatorUtilizedSDBalance(arg0 common.Address) (*big.Int, error) {
	return _SdCollateral.Contract.OperatorUtilizedSDBalance(&_SdCollateral.CallOpts, arg0)
}

// OperatorUtilizedSDBalance is a free data retrieval call binding the contract method 0xb11a3a9b.
//
// Solidity: function operatorUtilizedSDBalance(address ) view returns(uint256)
func (_SdCollateral *SdCollateralCallerSession) OperatorUtilizedSDBalance(arg0 common.Address) (*big.Int, error) {
	return _SdCollateral.Contract.OperatorUtilizedSDBalance(&_SdCollateral.CallOpts, arg0)
}

// PoolThresholdbyPoolId is a free data retrieval call binding the contract method 0x8a9b3738.
//
// Solidity: function poolThresholdbyPoolId(uint8 ) view returns(uint256 minThreshold, uint256 maxThreshold, uint256 withdrawThreshold, string units)
func (_SdCollateral *SdCollateralCaller) PoolThresholdbyPoolId(opts *bind.CallOpts, arg0 uint8) (struct {
	MinThreshold      *big.Int
	MaxThreshold      *big.Int
	WithdrawThreshold *big.Int
	Units             string
}, error) {
	var out []interface{}
	err := _SdCollateral.contract.Call(opts, &out, "poolThresholdbyPoolId", arg0)

	outstruct := new(struct {
		MinThreshold      *big.Int
		MaxThreshold      *big.Int
		WithdrawThreshold *big.Int
		Units             string
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.MinThreshold = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.MaxThreshold = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.WithdrawThreshold = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	outstruct.Units = *abi.ConvertType(out[3], new(string)).(*string)

	return *outstruct, err

}

// PoolThresholdbyPoolId is a free data retrieval call binding the contract method 0x8a9b3738.
//
// Solidity: function poolThresholdbyPoolId(uint8 ) view returns(uint256 minThreshold, uint256 maxThreshold, uint256 withdrawThreshold, string units)
func (_SdCollateral *SdCollateralSession) PoolThresholdbyPoolId(arg0 uint8) (struct {
	MinThreshold      *big.Int
	MaxThreshold      *big.Int
	WithdrawThreshold *big.Int
	Units             string
}, error) {
	return _SdCollateral.Contract.PoolThresholdbyPoolId(&_SdCollateral.CallOpts, arg0)
}

// PoolThresholdbyPoolId is a free data retrieval call binding the contract method 0x8a9b3738.
//
// Solidity: function poolThresholdbyPoolId(uint8 ) view returns(uint256 minThreshold, uint256 maxThreshold, uint256 withdrawThreshold, string units)
func (_SdCollateral *SdCollateralCallerSession) PoolThresholdbyPoolId(arg0 uint8) (struct {
	MinThreshold      *big.Int
	MaxThreshold      *big.Int
	WithdrawThreshold *big.Int
	Units             string
}, error) {
	return _SdCollateral.Contract.PoolThresholdbyPoolId(&_SdCollateral.CallOpts, arg0)
}

// StaderConfig is a free data retrieval call binding the contract method 0x490ffa35.
//
// Solidity: function staderConfig() view returns(address)
func (_SdCollateral *SdCollateralCaller) StaderConfig(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _SdCollateral.contract.Call(opts, &out, "staderConfig")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// StaderConfig is a free data retrieval call binding the contract method 0x490ffa35.
//
// Solidity: function staderConfig() view returns(address)
func (_SdCollateral *SdCollateralSession) StaderConfig() (common.Address, error) {
	return _SdCollateral.Contract.StaderConfig(&_SdCollateral.CallOpts)
}

// StaderConfig is a free data retrieval call binding the contract method 0x490ffa35.
//
// Solidity: function staderConfig() view returns(address)
func (_SdCollateral *SdCollateralCallerSession) StaderConfig() (common.Address, error) {
	return _SdCollateral.Contract.StaderConfig(&_SdCollateral.CallOpts)
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

// DepositSDAsCollateralOnBehalf is a paid mutator transaction binding the contract method 0xc35d0016.
//
// Solidity: function depositSDAsCollateralOnBehalf(address _operator, uint256 _sdAmount) returns()
func (_SdCollateral *SdCollateralTransactor) DepositSDAsCollateralOnBehalf(opts *bind.TransactOpts, _operator common.Address, _sdAmount *big.Int) (*types.Transaction, error) {
	return _SdCollateral.contract.Transact(opts, "depositSDAsCollateralOnBehalf", _operator, _sdAmount)
}

// DepositSDAsCollateralOnBehalf is a paid mutator transaction binding the contract method 0xc35d0016.
//
// Solidity: function depositSDAsCollateralOnBehalf(address _operator, uint256 _sdAmount) returns()
func (_SdCollateral *SdCollateralSession) DepositSDAsCollateralOnBehalf(_operator common.Address, _sdAmount *big.Int) (*types.Transaction, error) {
	return _SdCollateral.Contract.DepositSDAsCollateralOnBehalf(&_SdCollateral.TransactOpts, _operator, _sdAmount)
}

// DepositSDAsCollateralOnBehalf is a paid mutator transaction binding the contract method 0xc35d0016.
//
// Solidity: function depositSDAsCollateralOnBehalf(address _operator, uint256 _sdAmount) returns()
func (_SdCollateral *SdCollateralTransactorSession) DepositSDAsCollateralOnBehalf(_operator common.Address, _sdAmount *big.Int) (*types.Transaction, error) {
	return _SdCollateral.Contract.DepositSDAsCollateralOnBehalf(&_SdCollateral.TransactOpts, _operator, _sdAmount)
}

// DepositUtilizedSD is a paid mutator transaction binding the contract method 0xd52eb99e.
//
// Solidity: function depositUtilizedSD(address _operator, uint256 _sdAmount) returns()
func (_SdCollateral *SdCollateralTransactor) DepositUtilizedSD(opts *bind.TransactOpts, _operator common.Address, _sdAmount *big.Int) (*types.Transaction, error) {
	return _SdCollateral.contract.Transact(opts, "depositUtilizedSD", _operator, _sdAmount)
}

// DepositUtilizedSD is a paid mutator transaction binding the contract method 0xd52eb99e.
//
// Solidity: function depositUtilizedSD(address _operator, uint256 _sdAmount) returns()
func (_SdCollateral *SdCollateralSession) DepositUtilizedSD(_operator common.Address, _sdAmount *big.Int) (*types.Transaction, error) {
	return _SdCollateral.Contract.DepositUtilizedSD(&_SdCollateral.TransactOpts, _operator, _sdAmount)
}

// DepositUtilizedSD is a paid mutator transaction binding the contract method 0xd52eb99e.
//
// Solidity: function depositUtilizedSD(address _operator, uint256 _sdAmount) returns()
func (_SdCollateral *SdCollateralTransactorSession) DepositUtilizedSD(_operator common.Address, _sdAmount *big.Int) (*types.Transaction, error) {
	return _SdCollateral.Contract.DepositUtilizedSD(&_SdCollateral.TransactOpts, _operator, _sdAmount)
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

// Initialize is a paid mutator transaction binding the contract method 0x485cc955.
//
// Solidity: function initialize(address _admin, address _staderConfig) returns()
func (_SdCollateral *SdCollateralTransactor) Initialize(opts *bind.TransactOpts, _admin common.Address, _staderConfig common.Address) (*types.Transaction, error) {
	return _SdCollateral.contract.Transact(opts, "initialize", _admin, _staderConfig)
}

// Initialize is a paid mutator transaction binding the contract method 0x485cc955.
//
// Solidity: function initialize(address _admin, address _staderConfig) returns()
func (_SdCollateral *SdCollateralSession) Initialize(_admin common.Address, _staderConfig common.Address) (*types.Transaction, error) {
	return _SdCollateral.Contract.Initialize(&_SdCollateral.TransactOpts, _admin, _staderConfig)
}

// Initialize is a paid mutator transaction binding the contract method 0x485cc955.
//
// Solidity: function initialize(address _admin, address _staderConfig) returns()
func (_SdCollateral *SdCollateralTransactorSession) Initialize(_admin common.Address, _staderConfig common.Address) (*types.Transaction, error) {
	return _SdCollateral.Contract.Initialize(&_SdCollateral.TransactOpts, _admin, _staderConfig)
}

// MaxApproveSD is a paid mutator transaction binding the contract method 0x3e04cd35.
//
// Solidity: function maxApproveSD() returns()
func (_SdCollateral *SdCollateralTransactor) MaxApproveSD(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SdCollateral.contract.Transact(opts, "maxApproveSD")
}

// MaxApproveSD is a paid mutator transaction binding the contract method 0x3e04cd35.
//
// Solidity: function maxApproveSD() returns()
func (_SdCollateral *SdCollateralSession) MaxApproveSD() (*types.Transaction, error) {
	return _SdCollateral.Contract.MaxApproveSD(&_SdCollateral.TransactOpts)
}

// MaxApproveSD is a paid mutator transaction binding the contract method 0x3e04cd35.
//
// Solidity: function maxApproveSD() returns()
func (_SdCollateral *SdCollateralTransactorSession) MaxApproveSD() (*types.Transaction, error) {
	return _SdCollateral.Contract.MaxApproveSD(&_SdCollateral.TransactOpts)
}

// ReduceUtilizedSDPosition is a paid mutator transaction binding the contract method 0x956b95e7.
//
// Solidity: function reduceUtilizedSDPosition(address _operator, uint256 _sdAmount) returns()
func (_SdCollateral *SdCollateralTransactor) ReduceUtilizedSDPosition(opts *bind.TransactOpts, _operator common.Address, _sdAmount *big.Int) (*types.Transaction, error) {
	return _SdCollateral.contract.Transact(opts, "reduceUtilizedSDPosition", _operator, _sdAmount)
}

// ReduceUtilizedSDPosition is a paid mutator transaction binding the contract method 0x956b95e7.
//
// Solidity: function reduceUtilizedSDPosition(address _operator, uint256 _sdAmount) returns()
func (_SdCollateral *SdCollateralSession) ReduceUtilizedSDPosition(_operator common.Address, _sdAmount *big.Int) (*types.Transaction, error) {
	return _SdCollateral.Contract.ReduceUtilizedSDPosition(&_SdCollateral.TransactOpts, _operator, _sdAmount)
}

// ReduceUtilizedSDPosition is a paid mutator transaction binding the contract method 0x956b95e7.
//
// Solidity: function reduceUtilizedSDPosition(address _operator, uint256 _sdAmount) returns()
func (_SdCollateral *SdCollateralTransactorSession) ReduceUtilizedSDPosition(_operator common.Address, _sdAmount *big.Int) (*types.Transaction, error) {
	return _SdCollateral.Contract.ReduceUtilizedSDPosition(&_SdCollateral.TransactOpts, _operator, _sdAmount)
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

// SlashValidatorSD is a paid mutator transaction binding the contract method 0x4c538f58.
//
// Solidity: function slashValidatorSD(uint256 _validatorId, uint8 _poolId) returns()
func (_SdCollateral *SdCollateralTransactor) SlashValidatorSD(opts *bind.TransactOpts, _validatorId *big.Int, _poolId uint8) (*types.Transaction, error) {
	return _SdCollateral.contract.Transact(opts, "slashValidatorSD", _validatorId, _poolId)
}

// SlashValidatorSD is a paid mutator transaction binding the contract method 0x4c538f58.
//
// Solidity: function slashValidatorSD(uint256 _validatorId, uint8 _poolId) returns()
func (_SdCollateral *SdCollateralSession) SlashValidatorSD(_validatorId *big.Int, _poolId uint8) (*types.Transaction, error) {
	return _SdCollateral.Contract.SlashValidatorSD(&_SdCollateral.TransactOpts, _validatorId, _poolId)
}

// SlashValidatorSD is a paid mutator transaction binding the contract method 0x4c538f58.
//
// Solidity: function slashValidatorSD(uint256 _validatorId, uint8 _poolId) returns()
func (_SdCollateral *SdCollateralTransactorSession) SlashValidatorSD(_validatorId *big.Int, _poolId uint8) (*types.Transaction, error) {
	return _SdCollateral.Contract.SlashValidatorSD(&_SdCollateral.TransactOpts, _validatorId, _poolId)
}

// UpdatePoolThreshold is a paid mutator transaction binding the contract method 0xe0412f0e.
//
// Solidity: function updatePoolThreshold(uint8 _poolId, uint256 _minThreshold, uint256 _maxThreshold, uint256 _withdrawThreshold, string _units) returns()
func (_SdCollateral *SdCollateralTransactor) UpdatePoolThreshold(opts *bind.TransactOpts, _poolId uint8, _minThreshold *big.Int, _maxThreshold *big.Int, _withdrawThreshold *big.Int, _units string) (*types.Transaction, error) {
	return _SdCollateral.contract.Transact(opts, "updatePoolThreshold", _poolId, _minThreshold, _maxThreshold, _withdrawThreshold, _units)
}

// UpdatePoolThreshold is a paid mutator transaction binding the contract method 0xe0412f0e.
//
// Solidity: function updatePoolThreshold(uint8 _poolId, uint256 _minThreshold, uint256 _maxThreshold, uint256 _withdrawThreshold, string _units) returns()
func (_SdCollateral *SdCollateralSession) UpdatePoolThreshold(_poolId uint8, _minThreshold *big.Int, _maxThreshold *big.Int, _withdrawThreshold *big.Int, _units string) (*types.Transaction, error) {
	return _SdCollateral.Contract.UpdatePoolThreshold(&_SdCollateral.TransactOpts, _poolId, _minThreshold, _maxThreshold, _withdrawThreshold, _units)
}

// UpdatePoolThreshold is a paid mutator transaction binding the contract method 0xe0412f0e.
//
// Solidity: function updatePoolThreshold(uint8 _poolId, uint256 _minThreshold, uint256 _maxThreshold, uint256 _withdrawThreshold, string _units) returns()
func (_SdCollateral *SdCollateralTransactorSession) UpdatePoolThreshold(_poolId uint8, _minThreshold *big.Int, _maxThreshold *big.Int, _withdrawThreshold *big.Int, _units string) (*types.Transaction, error) {
	return _SdCollateral.Contract.UpdatePoolThreshold(&_SdCollateral.TransactOpts, _poolId, _minThreshold, _maxThreshold, _withdrawThreshold, _units)
}

// UpdateStaderConfig is a paid mutator transaction binding the contract method 0x9ee804cb.
//
// Solidity: function updateStaderConfig(address _staderConfig) returns()
func (_SdCollateral *SdCollateralTransactor) UpdateStaderConfig(opts *bind.TransactOpts, _staderConfig common.Address) (*types.Transaction, error) {
	return _SdCollateral.contract.Transact(opts, "updateStaderConfig", _staderConfig)
}

// UpdateStaderConfig is a paid mutator transaction binding the contract method 0x9ee804cb.
//
// Solidity: function updateStaderConfig(address _staderConfig) returns()
func (_SdCollateral *SdCollateralSession) UpdateStaderConfig(_staderConfig common.Address) (*types.Transaction, error) {
	return _SdCollateral.Contract.UpdateStaderConfig(&_SdCollateral.TransactOpts, _staderConfig)
}

// UpdateStaderConfig is a paid mutator transaction binding the contract method 0x9ee804cb.
//
// Solidity: function updateStaderConfig(address _staderConfig) returns()
func (_SdCollateral *SdCollateralTransactorSession) UpdateStaderConfig(_staderConfig common.Address) (*types.Transaction, error) {
	return _SdCollateral.Contract.UpdateStaderConfig(&_SdCollateral.TransactOpts, _staderConfig)
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

// SdCollateralReducedUtilizedPositionIterator is returned from FilterReducedUtilizedPosition and is used to iterate over the raw logs and unpacked data for ReducedUtilizedPosition events raised by the SdCollateral contract.
type SdCollateralReducedUtilizedPositionIterator struct {
	Event *SdCollateralReducedUtilizedPosition // Event containing the contract specifics and raw log

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
func (it *SdCollateralReducedUtilizedPositionIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SdCollateralReducedUtilizedPosition)
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
		it.Event = new(SdCollateralReducedUtilizedPosition)
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
func (it *SdCollateralReducedUtilizedPositionIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SdCollateralReducedUtilizedPositionIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SdCollateralReducedUtilizedPosition represents a ReducedUtilizedPosition event raised by the SdCollateral contract.
type SdCollateralReducedUtilizedPosition struct {
	Operator common.Address
	SdAmount *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterReducedUtilizedPosition is a free log retrieval operation binding the contract event 0x3819b1ebb313cb71368635e9c11bf524c47e692fbb7e68d46885341aa0d9fdeb.
//
// Solidity: event ReducedUtilizedPosition(address indexed operator, uint256 sdAmount)
func (_SdCollateral *SdCollateralFilterer) FilterReducedUtilizedPosition(opts *bind.FilterOpts, operator []common.Address) (*SdCollateralReducedUtilizedPositionIterator, error) {

	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _SdCollateral.contract.FilterLogs(opts, "ReducedUtilizedPosition", operatorRule)
	if err != nil {
		return nil, err
	}
	return &SdCollateralReducedUtilizedPositionIterator{contract: _SdCollateral.contract, event: "ReducedUtilizedPosition", logs: logs, sub: sub}, nil
}

// WatchReducedUtilizedPosition is a free log subscription operation binding the contract event 0x3819b1ebb313cb71368635e9c11bf524c47e692fbb7e68d46885341aa0d9fdeb.
//
// Solidity: event ReducedUtilizedPosition(address indexed operator, uint256 sdAmount)
func (_SdCollateral *SdCollateralFilterer) WatchReducedUtilizedPosition(opts *bind.WatchOpts, sink chan<- *SdCollateralReducedUtilizedPosition, operator []common.Address) (event.Subscription, error) {

	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _SdCollateral.contract.WatchLogs(opts, "ReducedUtilizedPosition", operatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SdCollateralReducedUtilizedPosition)
				if err := _SdCollateral.contract.UnpackLog(event, "ReducedUtilizedPosition", log); err != nil {
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

// ParseReducedUtilizedPosition is a log parse operation binding the contract event 0x3819b1ebb313cb71368635e9c11bf524c47e692fbb7e68d46885341aa0d9fdeb.
//
// Solidity: event ReducedUtilizedPosition(address indexed operator, uint256 sdAmount)
func (_SdCollateral *SdCollateralFilterer) ParseReducedUtilizedPosition(log types.Log) (*SdCollateralReducedUtilizedPosition, error) {
	event := new(SdCollateralReducedUtilizedPosition)
	if err := _SdCollateral.contract.UnpackLog(event, "ReducedUtilizedPosition", log); err != nil {
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

// SdCollateralSDDepositedIterator is returned from FilterSDDeposited and is used to iterate over the raw logs and unpacked data for SDDeposited events raised by the SdCollateral contract.
type SdCollateralSDDepositedIterator struct {
	Event *SdCollateralSDDeposited // Event containing the contract specifics and raw log

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
func (it *SdCollateralSDDepositedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SdCollateralSDDeposited)
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
		it.Event = new(SdCollateralSDDeposited)
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
func (it *SdCollateralSDDepositedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SdCollateralSDDepositedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SdCollateralSDDeposited represents a SDDeposited event raised by the SdCollateral contract.
type SdCollateralSDDeposited struct {
	Operator common.Address
	SdAmount *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterSDDeposited is a free log retrieval operation binding the contract event 0x112973aa2e3b182a34447572d830c1e97afc414558a08f33554be8e545224259.
//
// Solidity: event SDDeposited(address indexed operator, uint256 sdAmount)
func (_SdCollateral *SdCollateralFilterer) FilterSDDeposited(opts *bind.FilterOpts, operator []common.Address) (*SdCollateralSDDepositedIterator, error) {

	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _SdCollateral.contract.FilterLogs(opts, "SDDeposited", operatorRule)
	if err != nil {
		return nil, err
	}
	return &SdCollateralSDDepositedIterator{contract: _SdCollateral.contract, event: "SDDeposited", logs: logs, sub: sub}, nil
}

// WatchSDDeposited is a free log subscription operation binding the contract event 0x112973aa2e3b182a34447572d830c1e97afc414558a08f33554be8e545224259.
//
// Solidity: event SDDeposited(address indexed operator, uint256 sdAmount)
func (_SdCollateral *SdCollateralFilterer) WatchSDDeposited(opts *bind.WatchOpts, sink chan<- *SdCollateralSDDeposited, operator []common.Address) (event.Subscription, error) {

	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _SdCollateral.contract.WatchLogs(opts, "SDDeposited", operatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SdCollateralSDDeposited)
				if err := _SdCollateral.contract.UnpackLog(event, "SDDeposited", log); err != nil {
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

// ParseSDDeposited is a log parse operation binding the contract event 0x112973aa2e3b182a34447572d830c1e97afc414558a08f33554be8e545224259.
//
// Solidity: event SDDeposited(address indexed operator, uint256 sdAmount)
func (_SdCollateral *SdCollateralFilterer) ParseSDDeposited(log types.Log) (*SdCollateralSDDeposited, error) {
	event := new(SdCollateralSDDeposited)
	if err := _SdCollateral.contract.UnpackLog(event, "SDDeposited", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SdCollateralSDRepaidIterator is returned from FilterSDRepaid and is used to iterate over the raw logs and unpacked data for SDRepaid events raised by the SdCollateral contract.
type SdCollateralSDRepaidIterator struct {
	Event *SdCollateralSDRepaid // Event containing the contract specifics and raw log

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
func (it *SdCollateralSDRepaidIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SdCollateralSDRepaid)
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
		it.Event = new(SdCollateralSDRepaid)
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
func (it *SdCollateralSDRepaidIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SdCollateralSDRepaidIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SdCollateralSDRepaid represents a SDRepaid event raised by the SdCollateral contract.
type SdCollateralSDRepaid struct {
	Operator    common.Address
	RepayAmount *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterSDRepaid is a free log retrieval operation binding the contract event 0x78403a1d3d348f0cc8bb6e4a35d395cb9da1458aecf5bf1f147c6df1e20749b0.
//
// Solidity: event SDRepaid(address operator, uint256 repayAmount)
func (_SdCollateral *SdCollateralFilterer) FilterSDRepaid(opts *bind.FilterOpts) (*SdCollateralSDRepaidIterator, error) {

	logs, sub, err := _SdCollateral.contract.FilterLogs(opts, "SDRepaid")
	if err != nil {
		return nil, err
	}
	return &SdCollateralSDRepaidIterator{contract: _SdCollateral.contract, event: "SDRepaid", logs: logs, sub: sub}, nil
}

// WatchSDRepaid is a free log subscription operation binding the contract event 0x78403a1d3d348f0cc8bb6e4a35d395cb9da1458aecf5bf1f147c6df1e20749b0.
//
// Solidity: event SDRepaid(address operator, uint256 repayAmount)
func (_SdCollateral *SdCollateralFilterer) WatchSDRepaid(opts *bind.WatchOpts, sink chan<- *SdCollateralSDRepaid) (event.Subscription, error) {

	logs, sub, err := _SdCollateral.contract.WatchLogs(opts, "SDRepaid")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SdCollateralSDRepaid)
				if err := _SdCollateral.contract.UnpackLog(event, "SDRepaid", log); err != nil {
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

// ParseSDRepaid is a log parse operation binding the contract event 0x78403a1d3d348f0cc8bb6e4a35d395cb9da1458aecf5bf1f147c6df1e20749b0.
//
// Solidity: event SDRepaid(address operator, uint256 repayAmount)
func (_SdCollateral *SdCollateralFilterer) ParseSDRepaid(log types.Log) (*SdCollateralSDRepaid, error) {
	event := new(SdCollateralSDRepaid)
	if err := _SdCollateral.contract.UnpackLog(event, "SDRepaid", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SdCollateralSDSlashedIterator is returned from FilterSDSlashed and is used to iterate over the raw logs and unpacked data for SDSlashed events raised by the SdCollateral contract.
type SdCollateralSDSlashedIterator struct {
	Event *SdCollateralSDSlashed // Event containing the contract specifics and raw log

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
func (it *SdCollateralSDSlashedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SdCollateralSDSlashed)
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
		it.Event = new(SdCollateralSDSlashed)
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
func (it *SdCollateralSDSlashedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SdCollateralSDSlashedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SdCollateralSDSlashed represents a SDSlashed event raised by the SdCollateral contract.
type SdCollateralSDSlashed struct {
	Operator  common.Address
	Auction   common.Address
	SdSlashed *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterSDSlashed is a free log retrieval operation binding the contract event 0xe4a6f5b1a676a94b2af7a506093c172e23d46a5bea6f2d783d4d32c9047800f4.
//
// Solidity: event SDSlashed(address indexed operator, address indexed auction, uint256 sdSlashed)
func (_SdCollateral *SdCollateralFilterer) FilterSDSlashed(opts *bind.FilterOpts, operator []common.Address, auction []common.Address) (*SdCollateralSDSlashedIterator, error) {

	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}
	var auctionRule []interface{}
	for _, auctionItem := range auction {
		auctionRule = append(auctionRule, auctionItem)
	}

	logs, sub, err := _SdCollateral.contract.FilterLogs(opts, "SDSlashed", operatorRule, auctionRule)
	if err != nil {
		return nil, err
	}
	return &SdCollateralSDSlashedIterator{contract: _SdCollateral.contract, event: "SDSlashed", logs: logs, sub: sub}, nil
}

// WatchSDSlashed is a free log subscription operation binding the contract event 0xe4a6f5b1a676a94b2af7a506093c172e23d46a5bea6f2d783d4d32c9047800f4.
//
// Solidity: event SDSlashed(address indexed operator, address indexed auction, uint256 sdSlashed)
func (_SdCollateral *SdCollateralFilterer) WatchSDSlashed(opts *bind.WatchOpts, sink chan<- *SdCollateralSDSlashed, operator []common.Address, auction []common.Address) (event.Subscription, error) {

	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}
	var auctionRule []interface{}
	for _, auctionItem := range auction {
		auctionRule = append(auctionRule, auctionItem)
	}

	logs, sub, err := _SdCollateral.contract.WatchLogs(opts, "SDSlashed", operatorRule, auctionRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SdCollateralSDSlashed)
				if err := _SdCollateral.contract.UnpackLog(event, "SDSlashed", log); err != nil {
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

// ParseSDSlashed is a log parse operation binding the contract event 0xe4a6f5b1a676a94b2af7a506093c172e23d46a5bea6f2d783d4d32c9047800f4.
//
// Solidity: event SDSlashed(address indexed operator, address indexed auction, uint256 sdSlashed)
func (_SdCollateral *SdCollateralFilterer) ParseSDSlashed(log types.Log) (*SdCollateralSDSlashed, error) {
	event := new(SdCollateralSDSlashed)
	if err := _SdCollateral.contract.UnpackLog(event, "SDSlashed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SdCollateralSDSlashedFromUtilizeIterator is returned from FilterSDSlashedFromUtilize and is used to iterate over the raw logs and unpacked data for SDSlashedFromUtilize events raised by the SdCollateral contract.
type SdCollateralSDSlashedFromUtilizeIterator struct {
	Event *SdCollateralSDSlashedFromUtilize // Event containing the contract specifics and raw log

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
func (it *SdCollateralSDSlashedFromUtilizeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SdCollateralSDSlashedFromUtilize)
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
		it.Event = new(SdCollateralSDSlashedFromUtilize)
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
func (it *SdCollateralSDSlashedFromUtilizeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SdCollateralSDSlashedFromUtilizeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SdCollateralSDSlashedFromUtilize represents a SDSlashedFromUtilize event raised by the SdCollateral contract.
type SdCollateralSDSlashedFromUtilize struct {
	Operator            common.Address
	SdSlashFromUtilized *big.Int
	Raw                 types.Log // Blockchain specific contextual infos
}

// FilterSDSlashedFromUtilize is a free log retrieval operation binding the contract event 0xe0674d9d9473524ba6cf67d596e4357df10898dc3feaf4edf2a7fb9ffdb41802.
//
// Solidity: event SDSlashedFromUtilize(address operator, uint256 sdSlashFromUtilized)
func (_SdCollateral *SdCollateralFilterer) FilterSDSlashedFromUtilize(opts *bind.FilterOpts) (*SdCollateralSDSlashedFromUtilizeIterator, error) {

	logs, sub, err := _SdCollateral.contract.FilterLogs(opts, "SDSlashedFromUtilize")
	if err != nil {
		return nil, err
	}
	return &SdCollateralSDSlashedFromUtilizeIterator{contract: _SdCollateral.contract, event: "SDSlashedFromUtilize", logs: logs, sub: sub}, nil
}

// WatchSDSlashedFromUtilize is a free log subscription operation binding the contract event 0xe0674d9d9473524ba6cf67d596e4357df10898dc3feaf4edf2a7fb9ffdb41802.
//
// Solidity: event SDSlashedFromUtilize(address operator, uint256 sdSlashFromUtilized)
func (_SdCollateral *SdCollateralFilterer) WatchSDSlashedFromUtilize(opts *bind.WatchOpts, sink chan<- *SdCollateralSDSlashedFromUtilize) (event.Subscription, error) {

	logs, sub, err := _SdCollateral.contract.WatchLogs(opts, "SDSlashedFromUtilize")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SdCollateralSDSlashedFromUtilize)
				if err := _SdCollateral.contract.UnpackLog(event, "SDSlashedFromUtilize", log); err != nil {
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

// ParseSDSlashedFromUtilize is a log parse operation binding the contract event 0xe0674d9d9473524ba6cf67d596e4357df10898dc3feaf4edf2a7fb9ffdb41802.
//
// Solidity: event SDSlashedFromUtilize(address operator, uint256 sdSlashFromUtilized)
func (_SdCollateral *SdCollateralFilterer) ParseSDSlashedFromUtilize(log types.Log) (*SdCollateralSDSlashedFromUtilize, error) {
	event := new(SdCollateralSDSlashedFromUtilize)
	if err := _SdCollateral.contract.UnpackLog(event, "SDSlashedFromUtilize", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SdCollateralSDWithdrawnIterator is returned from FilterSDWithdrawn and is used to iterate over the raw logs and unpacked data for SDWithdrawn events raised by the SdCollateral contract.
type SdCollateralSDWithdrawnIterator struct {
	Event *SdCollateralSDWithdrawn // Event containing the contract specifics and raw log

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
func (it *SdCollateralSDWithdrawnIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SdCollateralSDWithdrawn)
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
		it.Event = new(SdCollateralSDWithdrawn)
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
func (it *SdCollateralSDWithdrawnIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SdCollateralSDWithdrawnIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SdCollateralSDWithdrawn represents a SDWithdrawn event raised by the SdCollateral contract.
type SdCollateralSDWithdrawn struct {
	Operator common.Address
	SdAmount *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterSDWithdrawn is a free log retrieval operation binding the contract event 0x48c1f846fa4bc05385324ee60316f9c6778ed2b5f205a6319678a609b8767607.
//
// Solidity: event SDWithdrawn(address indexed operator, uint256 sdAmount)
func (_SdCollateral *SdCollateralFilterer) FilterSDWithdrawn(opts *bind.FilterOpts, operator []common.Address) (*SdCollateralSDWithdrawnIterator, error) {

	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _SdCollateral.contract.FilterLogs(opts, "SDWithdrawn", operatorRule)
	if err != nil {
		return nil, err
	}
	return &SdCollateralSDWithdrawnIterator{contract: _SdCollateral.contract, event: "SDWithdrawn", logs: logs, sub: sub}, nil
}

// WatchSDWithdrawn is a free log subscription operation binding the contract event 0x48c1f846fa4bc05385324ee60316f9c6778ed2b5f205a6319678a609b8767607.
//
// Solidity: event SDWithdrawn(address indexed operator, uint256 sdAmount)
func (_SdCollateral *SdCollateralFilterer) WatchSDWithdrawn(opts *bind.WatchOpts, sink chan<- *SdCollateralSDWithdrawn, operator []common.Address) (event.Subscription, error) {

	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _SdCollateral.contract.WatchLogs(opts, "SDWithdrawn", operatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SdCollateralSDWithdrawn)
				if err := _SdCollateral.contract.UnpackLog(event, "SDWithdrawn", log); err != nil {
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

// ParseSDWithdrawn is a log parse operation binding the contract event 0x48c1f846fa4bc05385324ee60316f9c6778ed2b5f205a6319678a609b8767607.
//
// Solidity: event SDWithdrawn(address indexed operator, uint256 sdAmount)
func (_SdCollateral *SdCollateralFilterer) ParseSDWithdrawn(log types.Log) (*SdCollateralSDWithdrawn, error) {
	event := new(SdCollateralSDWithdrawn)
	if err := _SdCollateral.contract.UnpackLog(event, "SDWithdrawn", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SdCollateralUpdatedPoolIdForOperatorIterator is returned from FilterUpdatedPoolIdForOperator and is used to iterate over the raw logs and unpacked data for UpdatedPoolIdForOperator events raised by the SdCollateral contract.
type SdCollateralUpdatedPoolIdForOperatorIterator struct {
	Event *SdCollateralUpdatedPoolIdForOperator // Event containing the contract specifics and raw log

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
func (it *SdCollateralUpdatedPoolIdForOperatorIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SdCollateralUpdatedPoolIdForOperator)
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
		it.Event = new(SdCollateralUpdatedPoolIdForOperator)
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
func (it *SdCollateralUpdatedPoolIdForOperatorIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SdCollateralUpdatedPoolIdForOperatorIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SdCollateralUpdatedPoolIdForOperator represents a UpdatedPoolIdForOperator event raised by the SdCollateral contract.
type SdCollateralUpdatedPoolIdForOperator struct {
	PoolId   uint8
	Operator common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterUpdatedPoolIdForOperator is a free log retrieval operation binding the contract event 0x834f00ba6adeb9f7123fa03b8252cdda3f81509cc96c3c2239420138fa1b895e.
//
// Solidity: event UpdatedPoolIdForOperator(uint8 poolId, address operator)
func (_SdCollateral *SdCollateralFilterer) FilterUpdatedPoolIdForOperator(opts *bind.FilterOpts) (*SdCollateralUpdatedPoolIdForOperatorIterator, error) {

	logs, sub, err := _SdCollateral.contract.FilterLogs(opts, "UpdatedPoolIdForOperator")
	if err != nil {
		return nil, err
	}
	return &SdCollateralUpdatedPoolIdForOperatorIterator{contract: _SdCollateral.contract, event: "UpdatedPoolIdForOperator", logs: logs, sub: sub}, nil
}

// WatchUpdatedPoolIdForOperator is a free log subscription operation binding the contract event 0x834f00ba6adeb9f7123fa03b8252cdda3f81509cc96c3c2239420138fa1b895e.
//
// Solidity: event UpdatedPoolIdForOperator(uint8 poolId, address operator)
func (_SdCollateral *SdCollateralFilterer) WatchUpdatedPoolIdForOperator(opts *bind.WatchOpts, sink chan<- *SdCollateralUpdatedPoolIdForOperator) (event.Subscription, error) {

	logs, sub, err := _SdCollateral.contract.WatchLogs(opts, "UpdatedPoolIdForOperator")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SdCollateralUpdatedPoolIdForOperator)
				if err := _SdCollateral.contract.UnpackLog(event, "UpdatedPoolIdForOperator", log); err != nil {
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

// ParseUpdatedPoolIdForOperator is a log parse operation binding the contract event 0x834f00ba6adeb9f7123fa03b8252cdda3f81509cc96c3c2239420138fa1b895e.
//
// Solidity: event UpdatedPoolIdForOperator(uint8 poolId, address operator)
func (_SdCollateral *SdCollateralFilterer) ParseUpdatedPoolIdForOperator(log types.Log) (*SdCollateralUpdatedPoolIdForOperator, error) {
	event := new(SdCollateralUpdatedPoolIdForOperator)
	if err := _SdCollateral.contract.UnpackLog(event, "UpdatedPoolIdForOperator", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SdCollateralUpdatedPoolThresholdIterator is returned from FilterUpdatedPoolThreshold and is used to iterate over the raw logs and unpacked data for UpdatedPoolThreshold events raised by the SdCollateral contract.
type SdCollateralUpdatedPoolThresholdIterator struct {
	Event *SdCollateralUpdatedPoolThreshold // Event containing the contract specifics and raw log

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
func (it *SdCollateralUpdatedPoolThresholdIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SdCollateralUpdatedPoolThreshold)
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
		it.Event = new(SdCollateralUpdatedPoolThreshold)
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
func (it *SdCollateralUpdatedPoolThresholdIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SdCollateralUpdatedPoolThresholdIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SdCollateralUpdatedPoolThreshold represents a UpdatedPoolThreshold event raised by the SdCollateral contract.
type SdCollateralUpdatedPoolThreshold struct {
	PoolId            uint8
	MinThreshold      *big.Int
	WithdrawThreshold *big.Int
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterUpdatedPoolThreshold is a free log retrieval operation binding the contract event 0x18757dd1fbfe2ad823e1bd4de3f8a2ee76b49f92f6aa34cc7cbf717cdf4d1758.
//
// Solidity: event UpdatedPoolThreshold(uint8 poolId, uint256 minThreshold, uint256 withdrawThreshold)
func (_SdCollateral *SdCollateralFilterer) FilterUpdatedPoolThreshold(opts *bind.FilterOpts) (*SdCollateralUpdatedPoolThresholdIterator, error) {

	logs, sub, err := _SdCollateral.contract.FilterLogs(opts, "UpdatedPoolThreshold")
	if err != nil {
		return nil, err
	}
	return &SdCollateralUpdatedPoolThresholdIterator{contract: _SdCollateral.contract, event: "UpdatedPoolThreshold", logs: logs, sub: sub}, nil
}

// WatchUpdatedPoolThreshold is a free log subscription operation binding the contract event 0x18757dd1fbfe2ad823e1bd4de3f8a2ee76b49f92f6aa34cc7cbf717cdf4d1758.
//
// Solidity: event UpdatedPoolThreshold(uint8 poolId, uint256 minThreshold, uint256 withdrawThreshold)
func (_SdCollateral *SdCollateralFilterer) WatchUpdatedPoolThreshold(opts *bind.WatchOpts, sink chan<- *SdCollateralUpdatedPoolThreshold) (event.Subscription, error) {

	logs, sub, err := _SdCollateral.contract.WatchLogs(opts, "UpdatedPoolThreshold")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SdCollateralUpdatedPoolThreshold)
				if err := _SdCollateral.contract.UnpackLog(event, "UpdatedPoolThreshold", log); err != nil {
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

// ParseUpdatedPoolThreshold is a log parse operation binding the contract event 0x18757dd1fbfe2ad823e1bd4de3f8a2ee76b49f92f6aa34cc7cbf717cdf4d1758.
//
// Solidity: event UpdatedPoolThreshold(uint8 poolId, uint256 minThreshold, uint256 withdrawThreshold)
func (_SdCollateral *SdCollateralFilterer) ParseUpdatedPoolThreshold(log types.Log) (*SdCollateralUpdatedPoolThreshold, error) {
	event := new(SdCollateralUpdatedPoolThreshold)
	if err := _SdCollateral.contract.UnpackLog(event, "UpdatedPoolThreshold", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SdCollateralUpdatedStaderConfigIterator is returned from FilterUpdatedStaderConfig and is used to iterate over the raw logs and unpacked data for UpdatedStaderConfig events raised by the SdCollateral contract.
type SdCollateralUpdatedStaderConfigIterator struct {
	Event *SdCollateralUpdatedStaderConfig // Event containing the contract specifics and raw log

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
func (it *SdCollateralUpdatedStaderConfigIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SdCollateralUpdatedStaderConfig)
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
		it.Event = new(SdCollateralUpdatedStaderConfig)
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
func (it *SdCollateralUpdatedStaderConfigIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SdCollateralUpdatedStaderConfigIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SdCollateralUpdatedStaderConfig represents a UpdatedStaderConfig event raised by the SdCollateral contract.
type SdCollateralUpdatedStaderConfig struct {
	StaderConfig common.Address
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterUpdatedStaderConfig is a free log retrieval operation binding the contract event 0xdb2219043d7b197cb235f1af0cf6d782d77dee3de19e3f4fb6d39aae633b4485.
//
// Solidity: event UpdatedStaderConfig(address indexed staderConfig)
func (_SdCollateral *SdCollateralFilterer) FilterUpdatedStaderConfig(opts *bind.FilterOpts, staderConfig []common.Address) (*SdCollateralUpdatedStaderConfigIterator, error) {

	var staderConfigRule []interface{}
	for _, staderConfigItem := range staderConfig {
		staderConfigRule = append(staderConfigRule, staderConfigItem)
	}

	logs, sub, err := _SdCollateral.contract.FilterLogs(opts, "UpdatedStaderConfig", staderConfigRule)
	if err != nil {
		return nil, err
	}
	return &SdCollateralUpdatedStaderConfigIterator{contract: _SdCollateral.contract, event: "UpdatedStaderConfig", logs: logs, sub: sub}, nil
}

// WatchUpdatedStaderConfig is a free log subscription operation binding the contract event 0xdb2219043d7b197cb235f1af0cf6d782d77dee3de19e3f4fb6d39aae633b4485.
//
// Solidity: event UpdatedStaderConfig(address indexed staderConfig)
func (_SdCollateral *SdCollateralFilterer) WatchUpdatedStaderConfig(opts *bind.WatchOpts, sink chan<- *SdCollateralUpdatedStaderConfig, staderConfig []common.Address) (event.Subscription, error) {

	var staderConfigRule []interface{}
	for _, staderConfigItem := range staderConfig {
		staderConfigRule = append(staderConfigRule, staderConfigItem)
	}

	logs, sub, err := _SdCollateral.contract.WatchLogs(opts, "UpdatedStaderConfig", staderConfigRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SdCollateralUpdatedStaderConfig)
				if err := _SdCollateral.contract.UnpackLog(event, "UpdatedStaderConfig", log); err != nil {
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
func (_SdCollateral *SdCollateralFilterer) ParseUpdatedStaderConfig(log types.Log) (*SdCollateralUpdatedStaderConfig, error) {
	event := new(SdCollateralUpdatedStaderConfig)
	if err := _SdCollateral.contract.UnpackLog(event, "UpdatedStaderConfig", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SdCollateralUtilizedSDDepositedIterator is returned from FilterUtilizedSDDeposited and is used to iterate over the raw logs and unpacked data for UtilizedSDDeposited events raised by the SdCollateral contract.
type SdCollateralUtilizedSDDepositedIterator struct {
	Event *SdCollateralUtilizedSDDeposited // Event containing the contract specifics and raw log

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
func (it *SdCollateralUtilizedSDDepositedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SdCollateralUtilizedSDDeposited)
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
		it.Event = new(SdCollateralUtilizedSDDeposited)
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
func (it *SdCollateralUtilizedSDDepositedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SdCollateralUtilizedSDDepositedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SdCollateralUtilizedSDDeposited represents a UtilizedSDDeposited event raised by the SdCollateral contract.
type SdCollateralUtilizedSDDeposited struct {
	Operator common.Address
	SdAmount *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterUtilizedSDDeposited is a free log retrieval operation binding the contract event 0x48ad8145b270d3b23a3a4e232b1e7900f334a73d223ef71986cefe090631b652.
//
// Solidity: event UtilizedSDDeposited(address indexed operator, uint256 sdAmount)
func (_SdCollateral *SdCollateralFilterer) FilterUtilizedSDDeposited(opts *bind.FilterOpts, operator []common.Address) (*SdCollateralUtilizedSDDepositedIterator, error) {

	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _SdCollateral.contract.FilterLogs(opts, "UtilizedSDDeposited", operatorRule)
	if err != nil {
		return nil, err
	}
	return &SdCollateralUtilizedSDDepositedIterator{contract: _SdCollateral.contract, event: "UtilizedSDDeposited", logs: logs, sub: sub}, nil
}

// WatchUtilizedSDDeposited is a free log subscription operation binding the contract event 0x48ad8145b270d3b23a3a4e232b1e7900f334a73d223ef71986cefe090631b652.
//
// Solidity: event UtilizedSDDeposited(address indexed operator, uint256 sdAmount)
func (_SdCollateral *SdCollateralFilterer) WatchUtilizedSDDeposited(opts *bind.WatchOpts, sink chan<- *SdCollateralUtilizedSDDeposited, operator []common.Address) (event.Subscription, error) {

	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _SdCollateral.contract.WatchLogs(opts, "UtilizedSDDeposited", operatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SdCollateralUtilizedSDDeposited)
				if err := _SdCollateral.contract.UnpackLog(event, "UtilizedSDDeposited", log); err != nil {
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

// ParseUtilizedSDDeposited is a log parse operation binding the contract event 0x48ad8145b270d3b23a3a4e232b1e7900f334a73d223ef71986cefe090631b652.
//
// Solidity: event UtilizedSDDeposited(address indexed operator, uint256 sdAmount)
func (_SdCollateral *SdCollateralFilterer) ParseUtilizedSDDeposited(log types.Log) (*SdCollateralUtilizedSDDeposited, error) {
	event := new(SdCollateralUtilizedSDDeposited)
	if err := _SdCollateral.contract.UnpackLog(event, "UtilizedSDDeposited", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
