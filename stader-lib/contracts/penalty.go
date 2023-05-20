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

// PenaltyTrackerMetaData contains all meta data concerning the PenaltyTracker contract.
var PenaltyTrackerMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"CallerNotManager\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"CallerNotWithdrawVault\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidPubkeyLength\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ValidatorSettled\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ZeroAddress\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"pubkey\",\"type\":\"bytes\"}],\"name\":\"ForceExitValidator\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"version\",\"type\":\"uint8\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"previousAdminRole\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"newAdminRole\",\"type\":\"bytes32\"}],\"name\":\"RoleAdminChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleGranted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleRevoked\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"pubkey\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"UpdatedAdditionalPenaltyAmount\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"mevTheftPenalty\",\"type\":\"uint256\"}],\"name\":\"UpdatedMEVTheftPenaltyPerStrike\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"missedAttestationPenalty\",\"type\":\"uint256\"}],\"name\":\"UpdatedMissedAttestationPenaltyPerStrike\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"penaltyOracleAddress\",\"type\":\"address\"}],\"name\":\"UpdatedPenaltyOracleAddress\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"staderConfig\",\"type\":\"address\"}],\"name\":\"UpdatedStaderConfig\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"totalPenaltyThreshold\",\"type\":\"uint256\"}],\"name\":\"UpdatedValidatorExitPenaltyThreshold\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"pubkey\",\"type\":\"bytes\"}],\"name\":\"ValidatorMarkedAsSettled\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"DEFAULT_ADMIN_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"additionalPenaltyAmount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_pubkeyRoot\",\"type\":\"bytes32\"}],\"name\":\"calculateMEVTheftPenalty\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_pubkeyRoot\",\"type\":\"bytes32\"}],\"name\":\"calculateMissedAttestationPenalty\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"_pubkey\",\"type\":\"bytes\"}],\"name\":\"getAdditionalPenaltyAmount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"}],\"name\":\"getRoleAdmin\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"grantRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"hasRole\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_admin\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_staderConfig\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_ratedOracleAddress\",\"type\":\"address\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"_poolId\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"_validatorId\",\"type\":\"uint256\"}],\"name\":\"markValidatorSettled\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"mevTheftPenaltyPerStrike\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"missedAttestationPenaltyPerStrike\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"ratedOracleAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"renounceRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"revokeRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"_pubkey\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"setAdditionalPenaltyAmount\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"staderConfig\",\"outputs\":[{\"internalType\":\"contractIStaderConfig\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"name\":\"totalPenaltyAmount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_mevTheftPenaltyPerStrike\",\"type\":\"uint256\"}],\"name\":\"updateMEVTheftPenaltyPerStrike\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_missedAttestationPenaltyPerStrike\",\"type\":\"uint256\"}],\"name\":\"updateMissedAttestationPenaltyPerStrike\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_ratedOracleAddress\",\"type\":\"address\"}],\"name\":\"updateRatedOracleAddress\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_staderConfig\",\"type\":\"address\"}],\"name\":\"updateStaderConfig\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes[]\",\"name\":\"_pubkey\",\"type\":\"bytes[]\"}],\"name\":\"updateTotalPenaltyAmount\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_validatorExitPenaltyThreshold\",\"type\":\"uint256\"}],\"name\":\"updateValidatorExitPenaltyThreshold\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"validatorExitPenaltyThreshold\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// PenaltyTrackerABI is the input ABI used to generate the binding from.
// Deprecated: Use PenaltyTrackerMetaData.ABI instead.
var PenaltyTrackerABI = PenaltyTrackerMetaData.ABI

// PenaltyTracker is an auto generated Go binding around an Ethereum contract.
type PenaltyTracker struct {
	PenaltyTrackerCaller     // Read-only binding to the contract
	PenaltyTrackerTransactor // Write-only binding to the contract
	PenaltyTrackerFilterer   // Log filterer for contract events
}

// PenaltyTrackerCaller is an auto generated read-only Go binding around an Ethereum contract.
type PenaltyTrackerCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PenaltyTrackerTransactor is an auto generated write-only Go binding around an Ethereum contract.
type PenaltyTrackerTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PenaltyTrackerFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type PenaltyTrackerFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PenaltyTrackerSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type PenaltyTrackerSession struct {
	Contract     *PenaltyTracker   // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// PenaltyTrackerCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type PenaltyTrackerCallerSession struct {
	Contract *PenaltyTrackerCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts         // Call options to use throughout this session
}

// PenaltyTrackerTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type PenaltyTrackerTransactorSession struct {
	Contract     *PenaltyTrackerTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts         // Transaction auth options to use throughout this session
}

// PenaltyTrackerRaw is an auto generated low-level Go binding around an Ethereum contract.
type PenaltyTrackerRaw struct {
	Contract *PenaltyTracker // Generic contract binding to access the raw methods on
}

// PenaltyTrackerCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type PenaltyTrackerCallerRaw struct {
	Contract *PenaltyTrackerCaller // Generic read-only contract binding to access the raw methods on
}

// PenaltyTrackerTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type PenaltyTrackerTransactorRaw struct {
	Contract *PenaltyTrackerTransactor // Generic write-only contract binding to access the raw methods on
}

// NewPenaltyTracker creates a new instance of PenaltyTracker, bound to a specific deployed contract.
func NewPenaltyTracker(address common.Address, backend bind.ContractBackend) (*PenaltyTracker, error) {
	contract, err := bindPenaltyTracker(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &PenaltyTracker{PenaltyTrackerCaller: PenaltyTrackerCaller{contract: contract}, PenaltyTrackerTransactor: PenaltyTrackerTransactor{contract: contract}, PenaltyTrackerFilterer: PenaltyTrackerFilterer{contract: contract}}, nil
}

// NewPenaltyTrackerCaller creates a new read-only instance of PenaltyTracker, bound to a specific deployed contract.
func NewPenaltyTrackerCaller(address common.Address, caller bind.ContractCaller) (*PenaltyTrackerCaller, error) {
	contract, err := bindPenaltyTracker(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &PenaltyTrackerCaller{contract: contract}, nil
}

// NewPenaltyTrackerTransactor creates a new write-only instance of PenaltyTracker, bound to a specific deployed contract.
func NewPenaltyTrackerTransactor(address common.Address, transactor bind.ContractTransactor) (*PenaltyTrackerTransactor, error) {
	contract, err := bindPenaltyTracker(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &PenaltyTrackerTransactor{contract: contract}, nil
}

// NewPenaltyTrackerFilterer creates a new log filterer instance of PenaltyTracker, bound to a specific deployed contract.
func NewPenaltyTrackerFilterer(address common.Address, filterer bind.ContractFilterer) (*PenaltyTrackerFilterer, error) {
	contract, err := bindPenaltyTracker(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &PenaltyTrackerFilterer{contract: contract}, nil
}

// bindPenaltyTracker binds a generic wrapper to an already deployed contract.
func bindPenaltyTracker(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := PenaltyTrackerMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_PenaltyTracker *PenaltyTrackerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _PenaltyTracker.Contract.PenaltyTrackerCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_PenaltyTracker *PenaltyTrackerRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PenaltyTracker.Contract.PenaltyTrackerTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_PenaltyTracker *PenaltyTrackerRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _PenaltyTracker.Contract.PenaltyTrackerTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_PenaltyTracker *PenaltyTrackerCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _PenaltyTracker.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_PenaltyTracker *PenaltyTrackerTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PenaltyTracker.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_PenaltyTracker *PenaltyTrackerTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _PenaltyTracker.Contract.contract.Transact(opts, method, params...)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_PenaltyTracker *PenaltyTrackerCaller) DEFAULTADMINROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _PenaltyTracker.contract.Call(opts, &out, "DEFAULT_ADMIN_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_PenaltyTracker *PenaltyTrackerSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _PenaltyTracker.Contract.DEFAULTADMINROLE(&_PenaltyTracker.CallOpts)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_PenaltyTracker *PenaltyTrackerCallerSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _PenaltyTracker.Contract.DEFAULTADMINROLE(&_PenaltyTracker.CallOpts)
}

// AdditionalPenaltyAmount is a free data retrieval call binding the contract method 0xde8b4b77.
//
// Solidity: function additionalPenaltyAmount(bytes32 ) view returns(uint256)
func (_PenaltyTracker *PenaltyTrackerCaller) AdditionalPenaltyAmount(opts *bind.CallOpts, arg0 [32]byte) (*big.Int, error) {
	var out []interface{}
	err := _PenaltyTracker.contract.Call(opts, &out, "additionalPenaltyAmount", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// AdditionalPenaltyAmount is a free data retrieval call binding the contract method 0xde8b4b77.
//
// Solidity: function additionalPenaltyAmount(bytes32 ) view returns(uint256)
func (_PenaltyTracker *PenaltyTrackerSession) AdditionalPenaltyAmount(arg0 [32]byte) (*big.Int, error) {
	return _PenaltyTracker.Contract.AdditionalPenaltyAmount(&_PenaltyTracker.CallOpts, arg0)
}

// AdditionalPenaltyAmount is a free data retrieval call binding the contract method 0xde8b4b77.
//
// Solidity: function additionalPenaltyAmount(bytes32 ) view returns(uint256)
func (_PenaltyTracker *PenaltyTrackerCallerSession) AdditionalPenaltyAmount(arg0 [32]byte) (*big.Int, error) {
	return _PenaltyTracker.Contract.AdditionalPenaltyAmount(&_PenaltyTracker.CallOpts, arg0)
}

// CalculateMissedAttestationPenalty is a free data retrieval call binding the contract method 0x40b75c0a.
//
// Solidity: function calculateMissedAttestationPenalty(bytes32 _pubkeyRoot) view returns(uint256)
func (_PenaltyTracker *PenaltyTrackerCaller) CalculateMissedAttestationPenalty(opts *bind.CallOpts, _pubkeyRoot [32]byte) (*big.Int, error) {
	var out []interface{}
	err := _PenaltyTracker.contract.Call(opts, &out, "calculateMissedAttestationPenalty", _pubkeyRoot)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CalculateMissedAttestationPenalty is a free data retrieval call binding the contract method 0x40b75c0a.
//
// Solidity: function calculateMissedAttestationPenalty(bytes32 _pubkeyRoot) view returns(uint256)
func (_PenaltyTracker *PenaltyTrackerSession) CalculateMissedAttestationPenalty(_pubkeyRoot [32]byte) (*big.Int, error) {
	return _PenaltyTracker.Contract.CalculateMissedAttestationPenalty(&_PenaltyTracker.CallOpts, _pubkeyRoot)
}

// CalculateMissedAttestationPenalty is a free data retrieval call binding the contract method 0x40b75c0a.
//
// Solidity: function calculateMissedAttestationPenalty(bytes32 _pubkeyRoot) view returns(uint256)
func (_PenaltyTracker *PenaltyTrackerCallerSession) CalculateMissedAttestationPenalty(_pubkeyRoot [32]byte) (*big.Int, error) {
	return _PenaltyTracker.Contract.CalculateMissedAttestationPenalty(&_PenaltyTracker.CallOpts, _pubkeyRoot)
}

// GetAdditionalPenaltyAmount is a free data retrieval call binding the contract method 0x683e7cb9.
//
// Solidity: function getAdditionalPenaltyAmount(bytes _pubkey) view returns(uint256)
func (_PenaltyTracker *PenaltyTrackerCaller) GetAdditionalPenaltyAmount(opts *bind.CallOpts, _pubkey []byte) (*big.Int, error) {
	var out []interface{}
	err := _PenaltyTracker.contract.Call(opts, &out, "getAdditionalPenaltyAmount", _pubkey)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetAdditionalPenaltyAmount is a free data retrieval call binding the contract method 0x683e7cb9.
//
// Solidity: function getAdditionalPenaltyAmount(bytes _pubkey) view returns(uint256)
func (_PenaltyTracker *PenaltyTrackerSession) GetAdditionalPenaltyAmount(_pubkey []byte) (*big.Int, error) {
	return _PenaltyTracker.Contract.GetAdditionalPenaltyAmount(&_PenaltyTracker.CallOpts, _pubkey)
}

// GetAdditionalPenaltyAmount is a free data retrieval call binding the contract method 0x683e7cb9.
//
// Solidity: function getAdditionalPenaltyAmount(bytes _pubkey) view returns(uint256)
func (_PenaltyTracker *PenaltyTrackerCallerSession) GetAdditionalPenaltyAmount(_pubkey []byte) (*big.Int, error) {
	return _PenaltyTracker.Contract.GetAdditionalPenaltyAmount(&_PenaltyTracker.CallOpts, _pubkey)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_PenaltyTracker *PenaltyTrackerCaller) GetRoleAdmin(opts *bind.CallOpts, role [32]byte) ([32]byte, error) {
	var out []interface{}
	err := _PenaltyTracker.contract.Call(opts, &out, "getRoleAdmin", role)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_PenaltyTracker *PenaltyTrackerSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _PenaltyTracker.Contract.GetRoleAdmin(&_PenaltyTracker.CallOpts, role)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_PenaltyTracker *PenaltyTrackerCallerSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _PenaltyTracker.Contract.GetRoleAdmin(&_PenaltyTracker.CallOpts, role)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_PenaltyTracker *PenaltyTrackerCaller) HasRole(opts *bind.CallOpts, role [32]byte, account common.Address) (bool, error) {
	var out []interface{}
	err := _PenaltyTracker.contract.Call(opts, &out, "hasRole", role, account)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_PenaltyTracker *PenaltyTrackerSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _PenaltyTracker.Contract.HasRole(&_PenaltyTracker.CallOpts, role, account)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_PenaltyTracker *PenaltyTrackerCallerSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _PenaltyTracker.Contract.HasRole(&_PenaltyTracker.CallOpts, role, account)
}

// MevTheftPenaltyPerStrike is a free data retrieval call binding the contract method 0x57628707.
//
// Solidity: function mevTheftPenaltyPerStrike() view returns(uint256)
func (_PenaltyTracker *PenaltyTrackerCaller) MevTheftPenaltyPerStrike(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _PenaltyTracker.contract.Call(opts, &out, "mevTheftPenaltyPerStrike")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MevTheftPenaltyPerStrike is a free data retrieval call binding the contract method 0x57628707.
//
// Solidity: function mevTheftPenaltyPerStrike() view returns(uint256)
func (_PenaltyTracker *PenaltyTrackerSession) MevTheftPenaltyPerStrike() (*big.Int, error) {
	return _PenaltyTracker.Contract.MevTheftPenaltyPerStrike(&_PenaltyTracker.CallOpts)
}

// MevTheftPenaltyPerStrike is a free data retrieval call binding the contract method 0x57628707.
//
// Solidity: function mevTheftPenaltyPerStrike() view returns(uint256)
func (_PenaltyTracker *PenaltyTrackerCallerSession) MevTheftPenaltyPerStrike() (*big.Int, error) {
	return _PenaltyTracker.Contract.MevTheftPenaltyPerStrike(&_PenaltyTracker.CallOpts)
}

// MissedAttestationPenaltyPerStrike is a free data retrieval call binding the contract method 0xb8a16865.
//
// Solidity: function missedAttestationPenaltyPerStrike() view returns(uint256)
func (_PenaltyTracker *PenaltyTrackerCaller) MissedAttestationPenaltyPerStrike(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _PenaltyTracker.contract.Call(opts, &out, "missedAttestationPenaltyPerStrike")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MissedAttestationPenaltyPerStrike is a free data retrieval call binding the contract method 0xb8a16865.
//
// Solidity: function missedAttestationPenaltyPerStrike() view returns(uint256)
func (_PenaltyTracker *PenaltyTrackerSession) MissedAttestationPenaltyPerStrike() (*big.Int, error) {
	return _PenaltyTracker.Contract.MissedAttestationPenaltyPerStrike(&_PenaltyTracker.CallOpts)
}

// MissedAttestationPenaltyPerStrike is a free data retrieval call binding the contract method 0xb8a16865.
//
// Solidity: function missedAttestationPenaltyPerStrike() view returns(uint256)
func (_PenaltyTracker *PenaltyTrackerCallerSession) MissedAttestationPenaltyPerStrike() (*big.Int, error) {
	return _PenaltyTracker.Contract.MissedAttestationPenaltyPerStrike(&_PenaltyTracker.CallOpts)
}

// RatedOracleAddress is a free data retrieval call binding the contract method 0x1d135104.
//
// Solidity: function ratedOracleAddress() view returns(address)
func (_PenaltyTracker *PenaltyTrackerCaller) RatedOracleAddress(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _PenaltyTracker.contract.Call(opts, &out, "ratedOracleAddress")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// RatedOracleAddress is a free data retrieval call binding the contract method 0x1d135104.
//
// Solidity: function ratedOracleAddress() view returns(address)
func (_PenaltyTracker *PenaltyTrackerSession) RatedOracleAddress() (common.Address, error) {
	return _PenaltyTracker.Contract.RatedOracleAddress(&_PenaltyTracker.CallOpts)
}

// RatedOracleAddress is a free data retrieval call binding the contract method 0x1d135104.
//
// Solidity: function ratedOracleAddress() view returns(address)
func (_PenaltyTracker *PenaltyTrackerCallerSession) RatedOracleAddress() (common.Address, error) {
	return _PenaltyTracker.Contract.RatedOracleAddress(&_PenaltyTracker.CallOpts)
}

// StaderConfig is a free data retrieval call binding the contract method 0x490ffa35.
//
// Solidity: function staderConfig() view returns(address)
func (_PenaltyTracker *PenaltyTrackerCaller) StaderConfig(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _PenaltyTracker.contract.Call(opts, &out, "staderConfig")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// StaderConfig is a free data retrieval call binding the contract method 0x490ffa35.
//
// Solidity: function staderConfig() view returns(address)
func (_PenaltyTracker *PenaltyTrackerSession) StaderConfig() (common.Address, error) {
	return _PenaltyTracker.Contract.StaderConfig(&_PenaltyTracker.CallOpts)
}

// StaderConfig is a free data retrieval call binding the contract method 0x490ffa35.
//
// Solidity: function staderConfig() view returns(address)
func (_PenaltyTracker *PenaltyTrackerCallerSession) StaderConfig() (common.Address, error) {
	return _PenaltyTracker.Contract.StaderConfig(&_PenaltyTracker.CallOpts)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_PenaltyTracker *PenaltyTrackerCaller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _PenaltyTracker.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_PenaltyTracker *PenaltyTrackerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _PenaltyTracker.Contract.SupportsInterface(&_PenaltyTracker.CallOpts, interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_PenaltyTracker *PenaltyTrackerCallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _PenaltyTracker.Contract.SupportsInterface(&_PenaltyTracker.CallOpts, interfaceId)
}

// TotalPenaltyAmount is a free data retrieval call binding the contract method 0xe3b9f45c.
//
// Solidity: function totalPenaltyAmount(bytes ) view returns(uint256)
func (_PenaltyTracker *PenaltyTrackerCaller) TotalPenaltyAmount(opts *bind.CallOpts, arg0 []byte) (*big.Int, error) {
	var out []interface{}
	err := _PenaltyTracker.contract.Call(opts, &out, "totalPenaltyAmount", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalPenaltyAmount is a free data retrieval call binding the contract method 0xe3b9f45c.
//
// Solidity: function totalPenaltyAmount(bytes ) view returns(uint256)
func (_PenaltyTracker *PenaltyTrackerSession) TotalPenaltyAmount(arg0 []byte) (*big.Int, error) {
	return _PenaltyTracker.Contract.TotalPenaltyAmount(&_PenaltyTracker.CallOpts, arg0)
}

// TotalPenaltyAmount is a free data retrieval call binding the contract method 0xe3b9f45c.
//
// Solidity: function totalPenaltyAmount(bytes ) view returns(uint256)
func (_PenaltyTracker *PenaltyTrackerCallerSession) TotalPenaltyAmount(arg0 []byte) (*big.Int, error) {
	return _PenaltyTracker.Contract.TotalPenaltyAmount(&_PenaltyTracker.CallOpts, arg0)
}

// ValidatorExitPenaltyThreshold is a free data retrieval call binding the contract method 0xad67dfbb.
//
// Solidity: function validatorExitPenaltyThreshold() view returns(uint256)
func (_PenaltyTracker *PenaltyTrackerCaller) ValidatorExitPenaltyThreshold(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _PenaltyTracker.contract.Call(opts, &out, "validatorExitPenaltyThreshold")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ValidatorExitPenaltyThreshold is a free data retrieval call binding the contract method 0xad67dfbb.
//
// Solidity: function validatorExitPenaltyThreshold() view returns(uint256)
func (_PenaltyTracker *PenaltyTrackerSession) ValidatorExitPenaltyThreshold() (*big.Int, error) {
	return _PenaltyTracker.Contract.ValidatorExitPenaltyThreshold(&_PenaltyTracker.CallOpts)
}

// ValidatorExitPenaltyThreshold is a free data retrieval call binding the contract method 0xad67dfbb.
//
// Solidity: function validatorExitPenaltyThreshold() view returns(uint256)
func (_PenaltyTracker *PenaltyTrackerCallerSession) ValidatorExitPenaltyThreshold() (*big.Int, error) {
	return _PenaltyTracker.Contract.ValidatorExitPenaltyThreshold(&_PenaltyTracker.CallOpts)
}

// CalculateMEVTheftPenalty is a paid mutator transaction binding the contract method 0xd27a7cd1.
//
// Solidity: function calculateMEVTheftPenalty(bytes32 _pubkeyRoot) returns(uint256)
func (_PenaltyTracker *PenaltyTrackerTransactor) CalculateMEVTheftPenalty(opts *bind.TransactOpts, _pubkeyRoot [32]byte) (*types.Transaction, error) {
	return _PenaltyTracker.contract.Transact(opts, "calculateMEVTheftPenalty", _pubkeyRoot)
}

// CalculateMEVTheftPenalty is a paid mutator transaction binding the contract method 0xd27a7cd1.
//
// Solidity: function calculateMEVTheftPenalty(bytes32 _pubkeyRoot) returns(uint256)
func (_PenaltyTracker *PenaltyTrackerSession) CalculateMEVTheftPenalty(_pubkeyRoot [32]byte) (*types.Transaction, error) {
	return _PenaltyTracker.Contract.CalculateMEVTheftPenalty(&_PenaltyTracker.TransactOpts, _pubkeyRoot)
}

// CalculateMEVTheftPenalty is a paid mutator transaction binding the contract method 0xd27a7cd1.
//
// Solidity: function calculateMEVTheftPenalty(bytes32 _pubkeyRoot) returns(uint256)
func (_PenaltyTracker *PenaltyTrackerTransactorSession) CalculateMEVTheftPenalty(_pubkeyRoot [32]byte) (*types.Transaction, error) {
	return _PenaltyTracker.Contract.CalculateMEVTheftPenalty(&_PenaltyTracker.TransactOpts, _pubkeyRoot)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_PenaltyTracker *PenaltyTrackerTransactor) GrantRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _PenaltyTracker.contract.Transact(opts, "grantRole", role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_PenaltyTracker *PenaltyTrackerSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _PenaltyTracker.Contract.GrantRole(&_PenaltyTracker.TransactOpts, role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_PenaltyTracker *PenaltyTrackerTransactorSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _PenaltyTracker.Contract.GrantRole(&_PenaltyTracker.TransactOpts, role, account)
}

// Initialize is a paid mutator transaction binding the contract method 0xc0c53b8b.
//
// Solidity: function initialize(address _admin, address _staderConfig, address _ratedOracleAddress) returns()
func (_PenaltyTracker *PenaltyTrackerTransactor) Initialize(opts *bind.TransactOpts, _admin common.Address, _staderConfig common.Address, _ratedOracleAddress common.Address) (*types.Transaction, error) {
	return _PenaltyTracker.contract.Transact(opts, "initialize", _admin, _staderConfig, _ratedOracleAddress)
}

// Initialize is a paid mutator transaction binding the contract method 0xc0c53b8b.
//
// Solidity: function initialize(address _admin, address _staderConfig, address _ratedOracleAddress) returns()
func (_PenaltyTracker *PenaltyTrackerSession) Initialize(_admin common.Address, _staderConfig common.Address, _ratedOracleAddress common.Address) (*types.Transaction, error) {
	return _PenaltyTracker.Contract.Initialize(&_PenaltyTracker.TransactOpts, _admin, _staderConfig, _ratedOracleAddress)
}

// Initialize is a paid mutator transaction binding the contract method 0xc0c53b8b.
//
// Solidity: function initialize(address _admin, address _staderConfig, address _ratedOracleAddress) returns()
func (_PenaltyTracker *PenaltyTrackerTransactorSession) Initialize(_admin common.Address, _staderConfig common.Address, _ratedOracleAddress common.Address) (*types.Transaction, error) {
	return _PenaltyTracker.Contract.Initialize(&_PenaltyTracker.TransactOpts, _admin, _staderConfig, _ratedOracleAddress)
}

// MarkValidatorSettled is a paid mutator transaction binding the contract method 0x27f000a6.
//
// Solidity: function markValidatorSettled(uint8 _poolId, uint256 _validatorId) returns()
func (_PenaltyTracker *PenaltyTrackerTransactor) MarkValidatorSettled(opts *bind.TransactOpts, _poolId uint8, _validatorId *big.Int) (*types.Transaction, error) {
	return _PenaltyTracker.contract.Transact(opts, "markValidatorSettled", _poolId, _validatorId)
}

// MarkValidatorSettled is a paid mutator transaction binding the contract method 0x27f000a6.
//
// Solidity: function markValidatorSettled(uint8 _poolId, uint256 _validatorId) returns()
func (_PenaltyTracker *PenaltyTrackerSession) MarkValidatorSettled(_poolId uint8, _validatorId *big.Int) (*types.Transaction, error) {
	return _PenaltyTracker.Contract.MarkValidatorSettled(&_PenaltyTracker.TransactOpts, _poolId, _validatorId)
}

// MarkValidatorSettled is a paid mutator transaction binding the contract method 0x27f000a6.
//
// Solidity: function markValidatorSettled(uint8 _poolId, uint256 _validatorId) returns()
func (_PenaltyTracker *PenaltyTrackerTransactorSession) MarkValidatorSettled(_poolId uint8, _validatorId *big.Int) (*types.Transaction, error) {
	return _PenaltyTracker.Contract.MarkValidatorSettled(&_PenaltyTracker.TransactOpts, _poolId, _validatorId)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address account) returns()
func (_PenaltyTracker *PenaltyTrackerTransactor) RenounceRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _PenaltyTracker.contract.Transact(opts, "renounceRole", role, account)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address account) returns()
func (_PenaltyTracker *PenaltyTrackerSession) RenounceRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _PenaltyTracker.Contract.RenounceRole(&_PenaltyTracker.TransactOpts, role, account)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address account) returns()
func (_PenaltyTracker *PenaltyTrackerTransactorSession) RenounceRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _PenaltyTracker.Contract.RenounceRole(&_PenaltyTracker.TransactOpts, role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_PenaltyTracker *PenaltyTrackerTransactor) RevokeRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _PenaltyTracker.contract.Transact(opts, "revokeRole", role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_PenaltyTracker *PenaltyTrackerSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _PenaltyTracker.Contract.RevokeRole(&_PenaltyTracker.TransactOpts, role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_PenaltyTracker *PenaltyTrackerTransactorSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _PenaltyTracker.Contract.RevokeRole(&_PenaltyTracker.TransactOpts, role, account)
}

// SetAdditionalPenaltyAmount is a paid mutator transaction binding the contract method 0x21c0e552.
//
// Solidity: function setAdditionalPenaltyAmount(bytes _pubkey, uint256 _amount) returns()
func (_PenaltyTracker *PenaltyTrackerTransactor) SetAdditionalPenaltyAmount(opts *bind.TransactOpts, _pubkey []byte, _amount *big.Int) (*types.Transaction, error) {
	return _PenaltyTracker.contract.Transact(opts, "setAdditionalPenaltyAmount", _pubkey, _amount)
}

// SetAdditionalPenaltyAmount is a paid mutator transaction binding the contract method 0x21c0e552.
//
// Solidity: function setAdditionalPenaltyAmount(bytes _pubkey, uint256 _amount) returns()
func (_PenaltyTracker *PenaltyTrackerSession) SetAdditionalPenaltyAmount(_pubkey []byte, _amount *big.Int) (*types.Transaction, error) {
	return _PenaltyTracker.Contract.SetAdditionalPenaltyAmount(&_PenaltyTracker.TransactOpts, _pubkey, _amount)
}

// SetAdditionalPenaltyAmount is a paid mutator transaction binding the contract method 0x21c0e552.
//
// Solidity: function setAdditionalPenaltyAmount(bytes _pubkey, uint256 _amount) returns()
func (_PenaltyTracker *PenaltyTrackerTransactorSession) SetAdditionalPenaltyAmount(_pubkey []byte, _amount *big.Int) (*types.Transaction, error) {
	return _PenaltyTracker.Contract.SetAdditionalPenaltyAmount(&_PenaltyTracker.TransactOpts, _pubkey, _amount)
}

// UpdateMEVTheftPenaltyPerStrike is a paid mutator transaction binding the contract method 0x4b45a3f5.
//
// Solidity: function updateMEVTheftPenaltyPerStrike(uint256 _mevTheftPenaltyPerStrike) returns()
func (_PenaltyTracker *PenaltyTrackerTransactor) UpdateMEVTheftPenaltyPerStrike(opts *bind.TransactOpts, _mevTheftPenaltyPerStrike *big.Int) (*types.Transaction, error) {
	return _PenaltyTracker.contract.Transact(opts, "updateMEVTheftPenaltyPerStrike", _mevTheftPenaltyPerStrike)
}

// UpdateMEVTheftPenaltyPerStrike is a paid mutator transaction binding the contract method 0x4b45a3f5.
//
// Solidity: function updateMEVTheftPenaltyPerStrike(uint256 _mevTheftPenaltyPerStrike) returns()
func (_PenaltyTracker *PenaltyTrackerSession) UpdateMEVTheftPenaltyPerStrike(_mevTheftPenaltyPerStrike *big.Int) (*types.Transaction, error) {
	return _PenaltyTracker.Contract.UpdateMEVTheftPenaltyPerStrike(&_PenaltyTracker.TransactOpts, _mevTheftPenaltyPerStrike)
}

// UpdateMEVTheftPenaltyPerStrike is a paid mutator transaction binding the contract method 0x4b45a3f5.
//
// Solidity: function updateMEVTheftPenaltyPerStrike(uint256 _mevTheftPenaltyPerStrike) returns()
func (_PenaltyTracker *PenaltyTrackerTransactorSession) UpdateMEVTheftPenaltyPerStrike(_mevTheftPenaltyPerStrike *big.Int) (*types.Transaction, error) {
	return _PenaltyTracker.Contract.UpdateMEVTheftPenaltyPerStrike(&_PenaltyTracker.TransactOpts, _mevTheftPenaltyPerStrike)
}

// UpdateMissedAttestationPenaltyPerStrike is a paid mutator transaction binding the contract method 0xb679c856.
//
// Solidity: function updateMissedAttestationPenaltyPerStrike(uint256 _missedAttestationPenaltyPerStrike) returns()
func (_PenaltyTracker *PenaltyTrackerTransactor) UpdateMissedAttestationPenaltyPerStrike(opts *bind.TransactOpts, _missedAttestationPenaltyPerStrike *big.Int) (*types.Transaction, error) {
	return _PenaltyTracker.contract.Transact(opts, "updateMissedAttestationPenaltyPerStrike", _missedAttestationPenaltyPerStrike)
}

// UpdateMissedAttestationPenaltyPerStrike is a paid mutator transaction binding the contract method 0xb679c856.
//
// Solidity: function updateMissedAttestationPenaltyPerStrike(uint256 _missedAttestationPenaltyPerStrike) returns()
func (_PenaltyTracker *PenaltyTrackerSession) UpdateMissedAttestationPenaltyPerStrike(_missedAttestationPenaltyPerStrike *big.Int) (*types.Transaction, error) {
	return _PenaltyTracker.Contract.UpdateMissedAttestationPenaltyPerStrike(&_PenaltyTracker.TransactOpts, _missedAttestationPenaltyPerStrike)
}

// UpdateMissedAttestationPenaltyPerStrike is a paid mutator transaction binding the contract method 0xb679c856.
//
// Solidity: function updateMissedAttestationPenaltyPerStrike(uint256 _missedAttestationPenaltyPerStrike) returns()
func (_PenaltyTracker *PenaltyTrackerTransactorSession) UpdateMissedAttestationPenaltyPerStrike(_missedAttestationPenaltyPerStrike *big.Int) (*types.Transaction, error) {
	return _PenaltyTracker.Contract.UpdateMissedAttestationPenaltyPerStrike(&_PenaltyTracker.TransactOpts, _missedAttestationPenaltyPerStrike)
}

// UpdateRatedOracleAddress is a paid mutator transaction binding the contract method 0x36e9df0f.
//
// Solidity: function updateRatedOracleAddress(address _ratedOracleAddress) returns()
func (_PenaltyTracker *PenaltyTrackerTransactor) UpdateRatedOracleAddress(opts *bind.TransactOpts, _ratedOracleAddress common.Address) (*types.Transaction, error) {
	return _PenaltyTracker.contract.Transact(opts, "updateRatedOracleAddress", _ratedOracleAddress)
}

// UpdateRatedOracleAddress is a paid mutator transaction binding the contract method 0x36e9df0f.
//
// Solidity: function updateRatedOracleAddress(address _ratedOracleAddress) returns()
func (_PenaltyTracker *PenaltyTrackerSession) UpdateRatedOracleAddress(_ratedOracleAddress common.Address) (*types.Transaction, error) {
	return _PenaltyTracker.Contract.UpdateRatedOracleAddress(&_PenaltyTracker.TransactOpts, _ratedOracleAddress)
}

// UpdateRatedOracleAddress is a paid mutator transaction binding the contract method 0x36e9df0f.
//
// Solidity: function updateRatedOracleAddress(address _ratedOracleAddress) returns()
func (_PenaltyTracker *PenaltyTrackerTransactorSession) UpdateRatedOracleAddress(_ratedOracleAddress common.Address) (*types.Transaction, error) {
	return _PenaltyTracker.Contract.UpdateRatedOracleAddress(&_PenaltyTracker.TransactOpts, _ratedOracleAddress)
}

// UpdateStaderConfig is a paid mutator transaction binding the contract method 0x9ee804cb.
//
// Solidity: function updateStaderConfig(address _staderConfig) returns()
func (_PenaltyTracker *PenaltyTrackerTransactor) UpdateStaderConfig(opts *bind.TransactOpts, _staderConfig common.Address) (*types.Transaction, error) {
	return _PenaltyTracker.contract.Transact(opts, "updateStaderConfig", _staderConfig)
}

// UpdateStaderConfig is a paid mutator transaction binding the contract method 0x9ee804cb.
//
// Solidity: function updateStaderConfig(address _staderConfig) returns()
func (_PenaltyTracker *PenaltyTrackerSession) UpdateStaderConfig(_staderConfig common.Address) (*types.Transaction, error) {
	return _PenaltyTracker.Contract.UpdateStaderConfig(&_PenaltyTracker.TransactOpts, _staderConfig)
}

// UpdateStaderConfig is a paid mutator transaction binding the contract method 0x9ee804cb.
//
// Solidity: function updateStaderConfig(address _staderConfig) returns()
func (_PenaltyTracker *PenaltyTrackerTransactorSession) UpdateStaderConfig(_staderConfig common.Address) (*types.Transaction, error) {
	return _PenaltyTracker.Contract.UpdateStaderConfig(&_PenaltyTracker.TransactOpts, _staderConfig)
}

// UpdateTotalPenaltyAmount is a paid mutator transaction binding the contract method 0xd68f725d.
//
// Solidity: function updateTotalPenaltyAmount(bytes[] _pubkey) returns()
func (_PenaltyTracker *PenaltyTrackerTransactor) UpdateTotalPenaltyAmount(opts *bind.TransactOpts, _pubkey [][]byte) (*types.Transaction, error) {
	return _PenaltyTracker.contract.Transact(opts, "updateTotalPenaltyAmount", _pubkey)
}

// UpdateTotalPenaltyAmount is a paid mutator transaction binding the contract method 0xd68f725d.
//
// Solidity: function updateTotalPenaltyAmount(bytes[] _pubkey) returns()
func (_PenaltyTracker *PenaltyTrackerSession) UpdateTotalPenaltyAmount(_pubkey [][]byte) (*types.Transaction, error) {
	return _PenaltyTracker.Contract.UpdateTotalPenaltyAmount(&_PenaltyTracker.TransactOpts, _pubkey)
}

// UpdateTotalPenaltyAmount is a paid mutator transaction binding the contract method 0xd68f725d.
//
// Solidity: function updateTotalPenaltyAmount(bytes[] _pubkey) returns()
func (_PenaltyTracker *PenaltyTrackerTransactorSession) UpdateTotalPenaltyAmount(_pubkey [][]byte) (*types.Transaction, error) {
	return _PenaltyTracker.Contract.UpdateTotalPenaltyAmount(&_PenaltyTracker.TransactOpts, _pubkey)
}

// UpdateValidatorExitPenaltyThreshold is a paid mutator transaction binding the contract method 0xb42cd621.
//
// Solidity: function updateValidatorExitPenaltyThreshold(uint256 _validatorExitPenaltyThreshold) returns()
func (_PenaltyTracker *PenaltyTrackerTransactor) UpdateValidatorExitPenaltyThreshold(opts *bind.TransactOpts, _validatorExitPenaltyThreshold *big.Int) (*types.Transaction, error) {
	return _PenaltyTracker.contract.Transact(opts, "updateValidatorExitPenaltyThreshold", _validatorExitPenaltyThreshold)
}

// UpdateValidatorExitPenaltyThreshold is a paid mutator transaction binding the contract method 0xb42cd621.
//
// Solidity: function updateValidatorExitPenaltyThreshold(uint256 _validatorExitPenaltyThreshold) returns()
func (_PenaltyTracker *PenaltyTrackerSession) UpdateValidatorExitPenaltyThreshold(_validatorExitPenaltyThreshold *big.Int) (*types.Transaction, error) {
	return _PenaltyTracker.Contract.UpdateValidatorExitPenaltyThreshold(&_PenaltyTracker.TransactOpts, _validatorExitPenaltyThreshold)
}

// UpdateValidatorExitPenaltyThreshold is a paid mutator transaction binding the contract method 0xb42cd621.
//
// Solidity: function updateValidatorExitPenaltyThreshold(uint256 _validatorExitPenaltyThreshold) returns()
func (_PenaltyTracker *PenaltyTrackerTransactorSession) UpdateValidatorExitPenaltyThreshold(_validatorExitPenaltyThreshold *big.Int) (*types.Transaction, error) {
	return _PenaltyTracker.Contract.UpdateValidatorExitPenaltyThreshold(&_PenaltyTracker.TransactOpts, _validatorExitPenaltyThreshold)
}

// PenaltyTrackerForceExitValidatorIterator is returned from FilterForceExitValidator and is used to iterate over the raw logs and unpacked data for ForceExitValidator events raised by the PenaltyTracker contract.
type PenaltyTrackerForceExitValidatorIterator struct {
	Event *PenaltyTrackerForceExitValidator // Event containing the contract specifics and raw log

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
func (it *PenaltyTrackerForceExitValidatorIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PenaltyTrackerForceExitValidator)
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
		it.Event = new(PenaltyTrackerForceExitValidator)
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
func (it *PenaltyTrackerForceExitValidatorIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PenaltyTrackerForceExitValidatorIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PenaltyTrackerForceExitValidator represents a ForceExitValidator event raised by the PenaltyTracker contract.
type PenaltyTrackerForceExitValidator struct {
	Pubkey []byte
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterForceExitValidator is a free log retrieval operation binding the contract event 0xb3e2417cfe8e24e37e70b305ab38cac40f3543e07257edc3bf5f95e492ce97cd.
//
// Solidity: event ForceExitValidator(bytes pubkey)
func (_PenaltyTracker *PenaltyTrackerFilterer) FilterForceExitValidator(opts *bind.FilterOpts) (*PenaltyTrackerForceExitValidatorIterator, error) {

	logs, sub, err := _PenaltyTracker.contract.FilterLogs(opts, "ForceExitValidator")
	if err != nil {
		return nil, err
	}
	return &PenaltyTrackerForceExitValidatorIterator{contract: _PenaltyTracker.contract, event: "ForceExitValidator", logs: logs, sub: sub}, nil
}

// WatchForceExitValidator is a free log subscription operation binding the contract event 0xb3e2417cfe8e24e37e70b305ab38cac40f3543e07257edc3bf5f95e492ce97cd.
//
// Solidity: event ForceExitValidator(bytes pubkey)
func (_PenaltyTracker *PenaltyTrackerFilterer) WatchForceExitValidator(opts *bind.WatchOpts, sink chan<- *PenaltyTrackerForceExitValidator) (event.Subscription, error) {

	logs, sub, err := _PenaltyTracker.contract.WatchLogs(opts, "ForceExitValidator")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PenaltyTrackerForceExitValidator)
				if err := _PenaltyTracker.contract.UnpackLog(event, "ForceExitValidator", log); err != nil {
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

// ParseForceExitValidator is a log parse operation binding the contract event 0xb3e2417cfe8e24e37e70b305ab38cac40f3543e07257edc3bf5f95e492ce97cd.
//
// Solidity: event ForceExitValidator(bytes pubkey)
func (_PenaltyTracker *PenaltyTrackerFilterer) ParseForceExitValidator(log types.Log) (*PenaltyTrackerForceExitValidator, error) {
	event := new(PenaltyTrackerForceExitValidator)
	if err := _PenaltyTracker.contract.UnpackLog(event, "ForceExitValidator", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PenaltyTrackerInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the PenaltyTracker contract.
type PenaltyTrackerInitializedIterator struct {
	Event *PenaltyTrackerInitialized // Event containing the contract specifics and raw log

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
func (it *PenaltyTrackerInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PenaltyTrackerInitialized)
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
		it.Event = new(PenaltyTrackerInitialized)
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
func (it *PenaltyTrackerInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PenaltyTrackerInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PenaltyTrackerInitialized represents a Initialized event raised by the PenaltyTracker contract.
type PenaltyTrackerInitialized struct {
	Version uint8
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_PenaltyTracker *PenaltyTrackerFilterer) FilterInitialized(opts *bind.FilterOpts) (*PenaltyTrackerInitializedIterator, error) {

	logs, sub, err := _PenaltyTracker.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &PenaltyTrackerInitializedIterator{contract: _PenaltyTracker.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_PenaltyTracker *PenaltyTrackerFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *PenaltyTrackerInitialized) (event.Subscription, error) {

	logs, sub, err := _PenaltyTracker.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PenaltyTrackerInitialized)
				if err := _PenaltyTracker.contract.UnpackLog(event, "Initialized", log); err != nil {
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
func (_PenaltyTracker *PenaltyTrackerFilterer) ParseInitialized(log types.Log) (*PenaltyTrackerInitialized, error) {
	event := new(PenaltyTrackerInitialized)
	if err := _PenaltyTracker.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PenaltyTrackerRoleAdminChangedIterator is returned from FilterRoleAdminChanged and is used to iterate over the raw logs and unpacked data for RoleAdminChanged events raised by the PenaltyTracker contract.
type PenaltyTrackerRoleAdminChangedIterator struct {
	Event *PenaltyTrackerRoleAdminChanged // Event containing the contract specifics and raw log

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
func (it *PenaltyTrackerRoleAdminChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PenaltyTrackerRoleAdminChanged)
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
		it.Event = new(PenaltyTrackerRoleAdminChanged)
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
func (it *PenaltyTrackerRoleAdminChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PenaltyTrackerRoleAdminChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PenaltyTrackerRoleAdminChanged represents a RoleAdminChanged event raised by the PenaltyTracker contract.
type PenaltyTrackerRoleAdminChanged struct {
	Role              [32]byte
	PreviousAdminRole [32]byte
	NewAdminRole      [32]byte
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterRoleAdminChanged is a free log retrieval operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_PenaltyTracker *PenaltyTrackerFilterer) FilterRoleAdminChanged(opts *bind.FilterOpts, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (*PenaltyTrackerRoleAdminChangedIterator, error) {

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

	logs, sub, err := _PenaltyTracker.contract.FilterLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return &PenaltyTrackerRoleAdminChangedIterator{contract: _PenaltyTracker.contract, event: "RoleAdminChanged", logs: logs, sub: sub}, nil
}

// WatchRoleAdminChanged is a free log subscription operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_PenaltyTracker *PenaltyTrackerFilterer) WatchRoleAdminChanged(opts *bind.WatchOpts, sink chan<- *PenaltyTrackerRoleAdminChanged, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (event.Subscription, error) {

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

	logs, sub, err := _PenaltyTracker.contract.WatchLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PenaltyTrackerRoleAdminChanged)
				if err := _PenaltyTracker.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
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
func (_PenaltyTracker *PenaltyTrackerFilterer) ParseRoleAdminChanged(log types.Log) (*PenaltyTrackerRoleAdminChanged, error) {
	event := new(PenaltyTrackerRoleAdminChanged)
	if err := _PenaltyTracker.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PenaltyTrackerRoleGrantedIterator is returned from FilterRoleGranted and is used to iterate over the raw logs and unpacked data for RoleGranted events raised by the PenaltyTracker contract.
type PenaltyTrackerRoleGrantedIterator struct {
	Event *PenaltyTrackerRoleGranted // Event containing the contract specifics and raw log

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
func (it *PenaltyTrackerRoleGrantedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PenaltyTrackerRoleGranted)
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
		it.Event = new(PenaltyTrackerRoleGranted)
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
func (it *PenaltyTrackerRoleGrantedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PenaltyTrackerRoleGrantedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PenaltyTrackerRoleGranted represents a RoleGranted event raised by the PenaltyTracker contract.
type PenaltyTrackerRoleGranted struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleGranted is a free log retrieval operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_PenaltyTracker *PenaltyTrackerFilterer) FilterRoleGranted(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*PenaltyTrackerRoleGrantedIterator, error) {

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

	logs, sub, err := _PenaltyTracker.contract.FilterLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &PenaltyTrackerRoleGrantedIterator{contract: _PenaltyTracker.contract, event: "RoleGranted", logs: logs, sub: sub}, nil
}

// WatchRoleGranted is a free log subscription operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_PenaltyTracker *PenaltyTrackerFilterer) WatchRoleGranted(opts *bind.WatchOpts, sink chan<- *PenaltyTrackerRoleGranted, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _PenaltyTracker.contract.WatchLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PenaltyTrackerRoleGranted)
				if err := _PenaltyTracker.contract.UnpackLog(event, "RoleGranted", log); err != nil {
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
func (_PenaltyTracker *PenaltyTrackerFilterer) ParseRoleGranted(log types.Log) (*PenaltyTrackerRoleGranted, error) {
	event := new(PenaltyTrackerRoleGranted)
	if err := _PenaltyTracker.contract.UnpackLog(event, "RoleGranted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PenaltyTrackerRoleRevokedIterator is returned from FilterRoleRevoked and is used to iterate over the raw logs and unpacked data for RoleRevoked events raised by the PenaltyTracker contract.
type PenaltyTrackerRoleRevokedIterator struct {
	Event *PenaltyTrackerRoleRevoked // Event containing the contract specifics and raw log

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
func (it *PenaltyTrackerRoleRevokedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PenaltyTrackerRoleRevoked)
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
		it.Event = new(PenaltyTrackerRoleRevoked)
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
func (it *PenaltyTrackerRoleRevokedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PenaltyTrackerRoleRevokedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PenaltyTrackerRoleRevoked represents a RoleRevoked event raised by the PenaltyTracker contract.
type PenaltyTrackerRoleRevoked struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleRevoked is a free log retrieval operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_PenaltyTracker *PenaltyTrackerFilterer) FilterRoleRevoked(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*PenaltyTrackerRoleRevokedIterator, error) {

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

	logs, sub, err := _PenaltyTracker.contract.FilterLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &PenaltyTrackerRoleRevokedIterator{contract: _PenaltyTracker.contract, event: "RoleRevoked", logs: logs, sub: sub}, nil
}

// WatchRoleRevoked is a free log subscription operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_PenaltyTracker *PenaltyTrackerFilterer) WatchRoleRevoked(opts *bind.WatchOpts, sink chan<- *PenaltyTrackerRoleRevoked, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _PenaltyTracker.contract.WatchLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PenaltyTrackerRoleRevoked)
				if err := _PenaltyTracker.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
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
func (_PenaltyTracker *PenaltyTrackerFilterer) ParseRoleRevoked(log types.Log) (*PenaltyTrackerRoleRevoked, error) {
	event := new(PenaltyTrackerRoleRevoked)
	if err := _PenaltyTracker.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PenaltyTrackerUpdatedAdditionalPenaltyAmountIterator is returned from FilterUpdatedAdditionalPenaltyAmount and is used to iterate over the raw logs and unpacked data for UpdatedAdditionalPenaltyAmount events raised by the PenaltyTracker contract.
type PenaltyTrackerUpdatedAdditionalPenaltyAmountIterator struct {
	Event *PenaltyTrackerUpdatedAdditionalPenaltyAmount // Event containing the contract specifics and raw log

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
func (it *PenaltyTrackerUpdatedAdditionalPenaltyAmountIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PenaltyTrackerUpdatedAdditionalPenaltyAmount)
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
		it.Event = new(PenaltyTrackerUpdatedAdditionalPenaltyAmount)
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
func (it *PenaltyTrackerUpdatedAdditionalPenaltyAmountIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PenaltyTrackerUpdatedAdditionalPenaltyAmountIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PenaltyTrackerUpdatedAdditionalPenaltyAmount represents a UpdatedAdditionalPenaltyAmount event raised by the PenaltyTracker contract.
type PenaltyTrackerUpdatedAdditionalPenaltyAmount struct {
	Pubkey []byte
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterUpdatedAdditionalPenaltyAmount is a free log retrieval operation binding the contract event 0xe2adfb4363f117cf16e1b1d555a780e217e9ebc1d1474da2a065c5fe1a3cf768.
//
// Solidity: event UpdatedAdditionalPenaltyAmount(bytes pubkey, uint256 amount)
func (_PenaltyTracker *PenaltyTrackerFilterer) FilterUpdatedAdditionalPenaltyAmount(opts *bind.FilterOpts) (*PenaltyTrackerUpdatedAdditionalPenaltyAmountIterator, error) {

	logs, sub, err := _PenaltyTracker.contract.FilterLogs(opts, "UpdatedAdditionalPenaltyAmount")
	if err != nil {
		return nil, err
	}
	return &PenaltyTrackerUpdatedAdditionalPenaltyAmountIterator{contract: _PenaltyTracker.contract, event: "UpdatedAdditionalPenaltyAmount", logs: logs, sub: sub}, nil
}

// WatchUpdatedAdditionalPenaltyAmount is a free log subscription operation binding the contract event 0xe2adfb4363f117cf16e1b1d555a780e217e9ebc1d1474da2a065c5fe1a3cf768.
//
// Solidity: event UpdatedAdditionalPenaltyAmount(bytes pubkey, uint256 amount)
func (_PenaltyTracker *PenaltyTrackerFilterer) WatchUpdatedAdditionalPenaltyAmount(opts *bind.WatchOpts, sink chan<- *PenaltyTrackerUpdatedAdditionalPenaltyAmount) (event.Subscription, error) {

	logs, sub, err := _PenaltyTracker.contract.WatchLogs(opts, "UpdatedAdditionalPenaltyAmount")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PenaltyTrackerUpdatedAdditionalPenaltyAmount)
				if err := _PenaltyTracker.contract.UnpackLog(event, "UpdatedAdditionalPenaltyAmount", log); err != nil {
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

// ParseUpdatedAdditionalPenaltyAmount is a log parse operation binding the contract event 0xe2adfb4363f117cf16e1b1d555a780e217e9ebc1d1474da2a065c5fe1a3cf768.
//
// Solidity: event UpdatedAdditionalPenaltyAmount(bytes pubkey, uint256 amount)
func (_PenaltyTracker *PenaltyTrackerFilterer) ParseUpdatedAdditionalPenaltyAmount(log types.Log) (*PenaltyTrackerUpdatedAdditionalPenaltyAmount, error) {
	event := new(PenaltyTrackerUpdatedAdditionalPenaltyAmount)
	if err := _PenaltyTracker.contract.UnpackLog(event, "UpdatedAdditionalPenaltyAmount", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PenaltyTrackerUpdatedMEVTheftPenaltyPerStrikeIterator is returned from FilterUpdatedMEVTheftPenaltyPerStrike and is used to iterate over the raw logs and unpacked data for UpdatedMEVTheftPenaltyPerStrike events raised by the PenaltyTracker contract.
type PenaltyTrackerUpdatedMEVTheftPenaltyPerStrikeIterator struct {
	Event *PenaltyTrackerUpdatedMEVTheftPenaltyPerStrike // Event containing the contract specifics and raw log

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
func (it *PenaltyTrackerUpdatedMEVTheftPenaltyPerStrikeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PenaltyTrackerUpdatedMEVTheftPenaltyPerStrike)
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
		it.Event = new(PenaltyTrackerUpdatedMEVTheftPenaltyPerStrike)
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
func (it *PenaltyTrackerUpdatedMEVTheftPenaltyPerStrikeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PenaltyTrackerUpdatedMEVTheftPenaltyPerStrikeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PenaltyTrackerUpdatedMEVTheftPenaltyPerStrike represents a UpdatedMEVTheftPenaltyPerStrike event raised by the PenaltyTracker contract.
type PenaltyTrackerUpdatedMEVTheftPenaltyPerStrike struct {
	MevTheftPenalty *big.Int
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterUpdatedMEVTheftPenaltyPerStrike is a free log retrieval operation binding the contract event 0x7452b3d7dc2c5652111f3ad07bbe1e71403b3466bf29a5583e57d09ddb668a69.
//
// Solidity: event UpdatedMEVTheftPenaltyPerStrike(uint256 mevTheftPenalty)
func (_PenaltyTracker *PenaltyTrackerFilterer) FilterUpdatedMEVTheftPenaltyPerStrike(opts *bind.FilterOpts) (*PenaltyTrackerUpdatedMEVTheftPenaltyPerStrikeIterator, error) {

	logs, sub, err := _PenaltyTracker.contract.FilterLogs(opts, "UpdatedMEVTheftPenaltyPerStrike")
	if err != nil {
		return nil, err
	}
	return &PenaltyTrackerUpdatedMEVTheftPenaltyPerStrikeIterator{contract: _PenaltyTracker.contract, event: "UpdatedMEVTheftPenaltyPerStrike", logs: logs, sub: sub}, nil
}

// WatchUpdatedMEVTheftPenaltyPerStrike is a free log subscription operation binding the contract event 0x7452b3d7dc2c5652111f3ad07bbe1e71403b3466bf29a5583e57d09ddb668a69.
//
// Solidity: event UpdatedMEVTheftPenaltyPerStrike(uint256 mevTheftPenalty)
func (_PenaltyTracker *PenaltyTrackerFilterer) WatchUpdatedMEVTheftPenaltyPerStrike(opts *bind.WatchOpts, sink chan<- *PenaltyTrackerUpdatedMEVTheftPenaltyPerStrike) (event.Subscription, error) {

	logs, sub, err := _PenaltyTracker.contract.WatchLogs(opts, "UpdatedMEVTheftPenaltyPerStrike")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PenaltyTrackerUpdatedMEVTheftPenaltyPerStrike)
				if err := _PenaltyTracker.contract.UnpackLog(event, "UpdatedMEVTheftPenaltyPerStrike", log); err != nil {
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

// ParseUpdatedMEVTheftPenaltyPerStrike is a log parse operation binding the contract event 0x7452b3d7dc2c5652111f3ad07bbe1e71403b3466bf29a5583e57d09ddb668a69.
//
// Solidity: event UpdatedMEVTheftPenaltyPerStrike(uint256 mevTheftPenalty)
func (_PenaltyTracker *PenaltyTrackerFilterer) ParseUpdatedMEVTheftPenaltyPerStrike(log types.Log) (*PenaltyTrackerUpdatedMEVTheftPenaltyPerStrike, error) {
	event := new(PenaltyTrackerUpdatedMEVTheftPenaltyPerStrike)
	if err := _PenaltyTracker.contract.UnpackLog(event, "UpdatedMEVTheftPenaltyPerStrike", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PenaltyTrackerUpdatedMissedAttestationPenaltyPerStrikeIterator is returned from FilterUpdatedMissedAttestationPenaltyPerStrike and is used to iterate over the raw logs and unpacked data for UpdatedMissedAttestationPenaltyPerStrike events raised by the PenaltyTracker contract.
type PenaltyTrackerUpdatedMissedAttestationPenaltyPerStrikeIterator struct {
	Event *PenaltyTrackerUpdatedMissedAttestationPenaltyPerStrike // Event containing the contract specifics and raw log

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
func (it *PenaltyTrackerUpdatedMissedAttestationPenaltyPerStrikeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PenaltyTrackerUpdatedMissedAttestationPenaltyPerStrike)
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
		it.Event = new(PenaltyTrackerUpdatedMissedAttestationPenaltyPerStrike)
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
func (it *PenaltyTrackerUpdatedMissedAttestationPenaltyPerStrikeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PenaltyTrackerUpdatedMissedAttestationPenaltyPerStrikeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PenaltyTrackerUpdatedMissedAttestationPenaltyPerStrike represents a UpdatedMissedAttestationPenaltyPerStrike event raised by the PenaltyTracker contract.
type PenaltyTrackerUpdatedMissedAttestationPenaltyPerStrike struct {
	MissedAttestationPenalty *big.Int
	Raw                      types.Log // Blockchain specific contextual infos
}

// FilterUpdatedMissedAttestationPenaltyPerStrike is a free log retrieval operation binding the contract event 0xd8b2f639322ed6a8b319f2db0a2d0711eb3f658a862135c51d2b44407426fd93.
//
// Solidity: event UpdatedMissedAttestationPenaltyPerStrike(uint256 missedAttestationPenalty)
func (_PenaltyTracker *PenaltyTrackerFilterer) FilterUpdatedMissedAttestationPenaltyPerStrike(opts *bind.FilterOpts) (*PenaltyTrackerUpdatedMissedAttestationPenaltyPerStrikeIterator, error) {

	logs, sub, err := _PenaltyTracker.contract.FilterLogs(opts, "UpdatedMissedAttestationPenaltyPerStrike")
	if err != nil {
		return nil, err
	}
	return &PenaltyTrackerUpdatedMissedAttestationPenaltyPerStrikeIterator{contract: _PenaltyTracker.contract, event: "UpdatedMissedAttestationPenaltyPerStrike", logs: logs, sub: sub}, nil
}

// WatchUpdatedMissedAttestationPenaltyPerStrike is a free log subscription operation binding the contract event 0xd8b2f639322ed6a8b319f2db0a2d0711eb3f658a862135c51d2b44407426fd93.
//
// Solidity: event UpdatedMissedAttestationPenaltyPerStrike(uint256 missedAttestationPenalty)
func (_PenaltyTracker *PenaltyTrackerFilterer) WatchUpdatedMissedAttestationPenaltyPerStrike(opts *bind.WatchOpts, sink chan<- *PenaltyTrackerUpdatedMissedAttestationPenaltyPerStrike) (event.Subscription, error) {

	logs, sub, err := _PenaltyTracker.contract.WatchLogs(opts, "UpdatedMissedAttestationPenaltyPerStrike")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PenaltyTrackerUpdatedMissedAttestationPenaltyPerStrike)
				if err := _PenaltyTracker.contract.UnpackLog(event, "UpdatedMissedAttestationPenaltyPerStrike", log); err != nil {
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

// ParseUpdatedMissedAttestationPenaltyPerStrike is a log parse operation binding the contract event 0xd8b2f639322ed6a8b319f2db0a2d0711eb3f658a862135c51d2b44407426fd93.
//
// Solidity: event UpdatedMissedAttestationPenaltyPerStrike(uint256 missedAttestationPenalty)
func (_PenaltyTracker *PenaltyTrackerFilterer) ParseUpdatedMissedAttestationPenaltyPerStrike(log types.Log) (*PenaltyTrackerUpdatedMissedAttestationPenaltyPerStrike, error) {
	event := new(PenaltyTrackerUpdatedMissedAttestationPenaltyPerStrike)
	if err := _PenaltyTracker.contract.UnpackLog(event, "UpdatedMissedAttestationPenaltyPerStrike", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PenaltyTrackerUpdatedPenaltyOracleAddressIterator is returned from FilterUpdatedPenaltyOracleAddress and is used to iterate over the raw logs and unpacked data for UpdatedPenaltyOracleAddress events raised by the PenaltyTracker contract.
type PenaltyTrackerUpdatedPenaltyOracleAddressIterator struct {
	Event *PenaltyTrackerUpdatedPenaltyOracleAddress // Event containing the contract specifics and raw log

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
func (it *PenaltyTrackerUpdatedPenaltyOracleAddressIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PenaltyTrackerUpdatedPenaltyOracleAddress)
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
		it.Event = new(PenaltyTrackerUpdatedPenaltyOracleAddress)
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
func (it *PenaltyTrackerUpdatedPenaltyOracleAddressIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PenaltyTrackerUpdatedPenaltyOracleAddressIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PenaltyTrackerUpdatedPenaltyOracleAddress represents a UpdatedPenaltyOracleAddress event raised by the PenaltyTracker contract.
type PenaltyTrackerUpdatedPenaltyOracleAddress struct {
	PenaltyOracleAddress common.Address
	Raw                  types.Log // Blockchain specific contextual infos
}

// FilterUpdatedPenaltyOracleAddress is a free log retrieval operation binding the contract event 0xd6608b6b84376c8d4b6b072c8c53f25921b1943e56c413ffca340c4e88b610ba.
//
// Solidity: event UpdatedPenaltyOracleAddress(address penaltyOracleAddress)
func (_PenaltyTracker *PenaltyTrackerFilterer) FilterUpdatedPenaltyOracleAddress(opts *bind.FilterOpts) (*PenaltyTrackerUpdatedPenaltyOracleAddressIterator, error) {

	logs, sub, err := _PenaltyTracker.contract.FilterLogs(opts, "UpdatedPenaltyOracleAddress")
	if err != nil {
		return nil, err
	}
	return &PenaltyTrackerUpdatedPenaltyOracleAddressIterator{contract: _PenaltyTracker.contract, event: "UpdatedPenaltyOracleAddress", logs: logs, sub: sub}, nil
}

// WatchUpdatedPenaltyOracleAddress is a free log subscription operation binding the contract event 0xd6608b6b84376c8d4b6b072c8c53f25921b1943e56c413ffca340c4e88b610ba.
//
// Solidity: event UpdatedPenaltyOracleAddress(address penaltyOracleAddress)
func (_PenaltyTracker *PenaltyTrackerFilterer) WatchUpdatedPenaltyOracleAddress(opts *bind.WatchOpts, sink chan<- *PenaltyTrackerUpdatedPenaltyOracleAddress) (event.Subscription, error) {

	logs, sub, err := _PenaltyTracker.contract.WatchLogs(opts, "UpdatedPenaltyOracleAddress")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PenaltyTrackerUpdatedPenaltyOracleAddress)
				if err := _PenaltyTracker.contract.UnpackLog(event, "UpdatedPenaltyOracleAddress", log); err != nil {
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

// ParseUpdatedPenaltyOracleAddress is a log parse operation binding the contract event 0xd6608b6b84376c8d4b6b072c8c53f25921b1943e56c413ffca340c4e88b610ba.
//
// Solidity: event UpdatedPenaltyOracleAddress(address penaltyOracleAddress)
func (_PenaltyTracker *PenaltyTrackerFilterer) ParseUpdatedPenaltyOracleAddress(log types.Log) (*PenaltyTrackerUpdatedPenaltyOracleAddress, error) {
	event := new(PenaltyTrackerUpdatedPenaltyOracleAddress)
	if err := _PenaltyTracker.contract.UnpackLog(event, "UpdatedPenaltyOracleAddress", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PenaltyTrackerUpdatedStaderConfigIterator is returned from FilterUpdatedStaderConfig and is used to iterate over the raw logs and unpacked data for UpdatedStaderConfig events raised by the PenaltyTracker contract.
type PenaltyTrackerUpdatedStaderConfigIterator struct {
	Event *PenaltyTrackerUpdatedStaderConfig // Event containing the contract specifics and raw log

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
func (it *PenaltyTrackerUpdatedStaderConfigIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PenaltyTrackerUpdatedStaderConfig)
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
		it.Event = new(PenaltyTrackerUpdatedStaderConfig)
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
func (it *PenaltyTrackerUpdatedStaderConfigIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PenaltyTrackerUpdatedStaderConfigIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PenaltyTrackerUpdatedStaderConfig represents a UpdatedStaderConfig event raised by the PenaltyTracker contract.
type PenaltyTrackerUpdatedStaderConfig struct {
	StaderConfig common.Address
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterUpdatedStaderConfig is a free log retrieval operation binding the contract event 0xdb2219043d7b197cb235f1af0cf6d782d77dee3de19e3f4fb6d39aae633b4485.
//
// Solidity: event UpdatedStaderConfig(address staderConfig)
func (_PenaltyTracker *PenaltyTrackerFilterer) FilterUpdatedStaderConfig(opts *bind.FilterOpts) (*PenaltyTrackerUpdatedStaderConfigIterator, error) {

	logs, sub, err := _PenaltyTracker.contract.FilterLogs(opts, "UpdatedStaderConfig")
	if err != nil {
		return nil, err
	}
	return &PenaltyTrackerUpdatedStaderConfigIterator{contract: _PenaltyTracker.contract, event: "UpdatedStaderConfig", logs: logs, sub: sub}, nil
}

// WatchUpdatedStaderConfig is a free log subscription operation binding the contract event 0xdb2219043d7b197cb235f1af0cf6d782d77dee3de19e3f4fb6d39aae633b4485.
//
// Solidity: event UpdatedStaderConfig(address staderConfig)
func (_PenaltyTracker *PenaltyTrackerFilterer) WatchUpdatedStaderConfig(opts *bind.WatchOpts, sink chan<- *PenaltyTrackerUpdatedStaderConfig) (event.Subscription, error) {

	logs, sub, err := _PenaltyTracker.contract.WatchLogs(opts, "UpdatedStaderConfig")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PenaltyTrackerUpdatedStaderConfig)
				if err := _PenaltyTracker.contract.UnpackLog(event, "UpdatedStaderConfig", log); err != nil {
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
func (_PenaltyTracker *PenaltyTrackerFilterer) ParseUpdatedStaderConfig(log types.Log) (*PenaltyTrackerUpdatedStaderConfig, error) {
	event := new(PenaltyTrackerUpdatedStaderConfig)
	if err := _PenaltyTracker.contract.UnpackLog(event, "UpdatedStaderConfig", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PenaltyTrackerUpdatedValidatorExitPenaltyThresholdIterator is returned from FilterUpdatedValidatorExitPenaltyThreshold and is used to iterate over the raw logs and unpacked data for UpdatedValidatorExitPenaltyThreshold events raised by the PenaltyTracker contract.
type PenaltyTrackerUpdatedValidatorExitPenaltyThresholdIterator struct {
	Event *PenaltyTrackerUpdatedValidatorExitPenaltyThreshold // Event containing the contract specifics and raw log

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
func (it *PenaltyTrackerUpdatedValidatorExitPenaltyThresholdIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PenaltyTrackerUpdatedValidatorExitPenaltyThreshold)
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
		it.Event = new(PenaltyTrackerUpdatedValidatorExitPenaltyThreshold)
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
func (it *PenaltyTrackerUpdatedValidatorExitPenaltyThresholdIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PenaltyTrackerUpdatedValidatorExitPenaltyThresholdIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PenaltyTrackerUpdatedValidatorExitPenaltyThreshold represents a UpdatedValidatorExitPenaltyThreshold event raised by the PenaltyTracker contract.
type PenaltyTrackerUpdatedValidatorExitPenaltyThreshold struct {
	TotalPenaltyThreshold *big.Int
	Raw                   types.Log // Blockchain specific contextual infos
}

// FilterUpdatedValidatorExitPenaltyThreshold is a free log retrieval operation binding the contract event 0x1560afd566079daff010f57d64aed4272dcd0f975cdd8c22e6bb4eee0f543eda.
//
// Solidity: event UpdatedValidatorExitPenaltyThreshold(uint256 totalPenaltyThreshold)
func (_PenaltyTracker *PenaltyTrackerFilterer) FilterUpdatedValidatorExitPenaltyThreshold(opts *bind.FilterOpts) (*PenaltyTrackerUpdatedValidatorExitPenaltyThresholdIterator, error) {

	logs, sub, err := _PenaltyTracker.contract.FilterLogs(opts, "UpdatedValidatorExitPenaltyThreshold")
	if err != nil {
		return nil, err
	}
	return &PenaltyTrackerUpdatedValidatorExitPenaltyThresholdIterator{contract: _PenaltyTracker.contract, event: "UpdatedValidatorExitPenaltyThreshold", logs: logs, sub: sub}, nil
}

// WatchUpdatedValidatorExitPenaltyThreshold is a free log subscription operation binding the contract event 0x1560afd566079daff010f57d64aed4272dcd0f975cdd8c22e6bb4eee0f543eda.
//
// Solidity: event UpdatedValidatorExitPenaltyThreshold(uint256 totalPenaltyThreshold)
func (_PenaltyTracker *PenaltyTrackerFilterer) WatchUpdatedValidatorExitPenaltyThreshold(opts *bind.WatchOpts, sink chan<- *PenaltyTrackerUpdatedValidatorExitPenaltyThreshold) (event.Subscription, error) {

	logs, sub, err := _PenaltyTracker.contract.WatchLogs(opts, "UpdatedValidatorExitPenaltyThreshold")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PenaltyTrackerUpdatedValidatorExitPenaltyThreshold)
				if err := _PenaltyTracker.contract.UnpackLog(event, "UpdatedValidatorExitPenaltyThreshold", log); err != nil {
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

// ParseUpdatedValidatorExitPenaltyThreshold is a log parse operation binding the contract event 0x1560afd566079daff010f57d64aed4272dcd0f975cdd8c22e6bb4eee0f543eda.
//
// Solidity: event UpdatedValidatorExitPenaltyThreshold(uint256 totalPenaltyThreshold)
func (_PenaltyTracker *PenaltyTrackerFilterer) ParseUpdatedValidatorExitPenaltyThreshold(log types.Log) (*PenaltyTrackerUpdatedValidatorExitPenaltyThreshold, error) {
	event := new(PenaltyTrackerUpdatedValidatorExitPenaltyThreshold)
	if err := _PenaltyTracker.contract.UnpackLog(event, "UpdatedValidatorExitPenaltyThreshold", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PenaltyTrackerValidatorMarkedAsSettledIterator is returned from FilterValidatorMarkedAsSettled and is used to iterate over the raw logs and unpacked data for ValidatorMarkedAsSettled events raised by the PenaltyTracker contract.
type PenaltyTrackerValidatorMarkedAsSettledIterator struct {
	Event *PenaltyTrackerValidatorMarkedAsSettled // Event containing the contract specifics and raw log

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
func (it *PenaltyTrackerValidatorMarkedAsSettledIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PenaltyTrackerValidatorMarkedAsSettled)
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
		it.Event = new(PenaltyTrackerValidatorMarkedAsSettled)
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
func (it *PenaltyTrackerValidatorMarkedAsSettledIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PenaltyTrackerValidatorMarkedAsSettledIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PenaltyTrackerValidatorMarkedAsSettled represents a ValidatorMarkedAsSettled event raised by the PenaltyTracker contract.
type PenaltyTrackerValidatorMarkedAsSettled struct {
	Pubkey []byte
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterValidatorMarkedAsSettled is a free log retrieval operation binding the contract event 0xcb48ff45614a20a9afe7c91270e1efcd412e49852d2c104e639feb398ea84d07.
//
// Solidity: event ValidatorMarkedAsSettled(bytes pubkey)
func (_PenaltyTracker *PenaltyTrackerFilterer) FilterValidatorMarkedAsSettled(opts *bind.FilterOpts) (*PenaltyTrackerValidatorMarkedAsSettledIterator, error) {

	logs, sub, err := _PenaltyTracker.contract.FilterLogs(opts, "ValidatorMarkedAsSettled")
	if err != nil {
		return nil, err
	}
	return &PenaltyTrackerValidatorMarkedAsSettledIterator{contract: _PenaltyTracker.contract, event: "ValidatorMarkedAsSettled", logs: logs, sub: sub}, nil
}

// WatchValidatorMarkedAsSettled is a free log subscription operation binding the contract event 0xcb48ff45614a20a9afe7c91270e1efcd412e49852d2c104e639feb398ea84d07.
//
// Solidity: event ValidatorMarkedAsSettled(bytes pubkey)
func (_PenaltyTracker *PenaltyTrackerFilterer) WatchValidatorMarkedAsSettled(opts *bind.WatchOpts, sink chan<- *PenaltyTrackerValidatorMarkedAsSettled) (event.Subscription, error) {

	logs, sub, err := _PenaltyTracker.contract.WatchLogs(opts, "ValidatorMarkedAsSettled")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PenaltyTrackerValidatorMarkedAsSettled)
				if err := _PenaltyTracker.contract.UnpackLog(event, "ValidatorMarkedAsSettled", log); err != nil {
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

// ParseValidatorMarkedAsSettled is a log parse operation binding the contract event 0xcb48ff45614a20a9afe7c91270e1efcd412e49852d2c104e639feb398ea84d07.
//
// Solidity: event ValidatorMarkedAsSettled(bytes pubkey)
func (_PenaltyTracker *PenaltyTrackerFilterer) ParseValidatorMarkedAsSettled(log types.Log) (*PenaltyTrackerValidatorMarkedAsSettled, error) {
	event := new(PenaltyTrackerValidatorMarkedAsSettled)
	if err := _PenaltyTracker.contract.UnpackLog(event, "ValidatorMarkedAsSettled", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
