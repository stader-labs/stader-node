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

// VaultProxyMetaData contains all meta data concerning the VaultProxy contract.
var VaultProxyMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"CallerNotOwner\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ZeroAddress\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"UpdatedOwner\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"staderConfig\",\"type\":\"address\"}],\"name\":\"UpdatedStaderConfig\",\"type\":\"event\"},{\"stateMutability\":\"payable\",\"type\":\"fallback\"},{\"inputs\":[],\"name\":\"id\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bool\",\"name\":\"_isValidatorWithdrawalVault\",\"type\":\"bool\"},{\"internalType\":\"uint8\",\"name\":\"_poolId\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"_id\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_staderConfig\",\"type\":\"address\"}],\"name\":\"initialise\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"isValidatorWithdrawalVault\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"poolId\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"staderConfig\",\"outputs\":[{\"internalType\":\"contractIStaderConfig\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_owner\",\"type\":\"address\"}],\"name\":\"updateOwner\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_staderConfig\",\"type\":\"address\"}],\"name\":\"updateStaderConfig\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"vaultSettleStatus\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// VaultProxyABI is the input ABI used to generate the binding from.
// Deprecated: Use VaultProxyMetaData.ABI instead.
var VaultProxyABI = VaultProxyMetaData.ABI

// VaultProxy is an auto generated Go binding around an Ethereum contract.
type VaultProxy struct {
	VaultProxyCaller     // Read-only binding to the contract
	VaultProxyTransactor // Write-only binding to the contract
	VaultProxyFilterer   // Log filterer for contract events
}

// VaultProxyCaller is an auto generated read-only Go binding around an Ethereum contract.
type VaultProxyCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// VaultProxyTransactor is an auto generated write-only Go binding around an Ethereum contract.
type VaultProxyTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// VaultProxyFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type VaultProxyFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// VaultProxySession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type VaultProxySession struct {
	Contract     *VaultProxy       // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// VaultProxyCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type VaultProxyCallerSession struct {
	Contract *VaultProxyCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts     // Call options to use throughout this session
}

// VaultProxyTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type VaultProxyTransactorSession struct {
	Contract     *VaultProxyTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// VaultProxyRaw is an auto generated low-level Go binding around an Ethereum contract.
type VaultProxyRaw struct {
	Contract *VaultProxy // Generic contract binding to access the raw methods on
}

// VaultProxyCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type VaultProxyCallerRaw struct {
	Contract *VaultProxyCaller // Generic read-only contract binding to access the raw methods on
}

// VaultProxyTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type VaultProxyTransactorRaw struct {
	Contract *VaultProxyTransactor // Generic write-only contract binding to access the raw methods on
}

// NewVaultProxy creates a new instance of VaultProxy, bound to a specific deployed contract.
func NewVaultProxy(address common.Address, backend bind.ContractBackend) (*VaultProxy, error) {
	contract, err := bindVaultProxy(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &VaultProxy{VaultProxyCaller: VaultProxyCaller{contract: contract}, VaultProxyTransactor: VaultProxyTransactor{contract: contract}, VaultProxyFilterer: VaultProxyFilterer{contract: contract}}, nil
}

// NewVaultProxyCaller creates a new read-only instance of VaultProxy, bound to a specific deployed contract.
func NewVaultProxyCaller(address common.Address, caller bind.ContractCaller) (*VaultProxyCaller, error) {
	contract, err := bindVaultProxy(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &VaultProxyCaller{contract: contract}, nil
}

// NewVaultProxyTransactor creates a new write-only instance of VaultProxy, bound to a specific deployed contract.
func NewVaultProxyTransactor(address common.Address, transactor bind.ContractTransactor) (*VaultProxyTransactor, error) {
	contract, err := bindVaultProxy(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &VaultProxyTransactor{contract: contract}, nil
}

// NewVaultProxyFilterer creates a new log filterer instance of VaultProxy, bound to a specific deployed contract.
func NewVaultProxyFilterer(address common.Address, filterer bind.ContractFilterer) (*VaultProxyFilterer, error) {
	contract, err := bindVaultProxy(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &VaultProxyFilterer{contract: contract}, nil
}

// bindVaultProxy binds a generic wrapper to an already deployed contract.
func bindVaultProxy(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := VaultProxyMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_VaultProxy *VaultProxyRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _VaultProxy.Contract.VaultProxyCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_VaultProxy *VaultProxyRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _VaultProxy.Contract.VaultProxyTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_VaultProxy *VaultProxyRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _VaultProxy.Contract.VaultProxyTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_VaultProxy *VaultProxyCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _VaultProxy.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_VaultProxy *VaultProxyTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _VaultProxy.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_VaultProxy *VaultProxyTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _VaultProxy.Contract.contract.Transact(opts, method, params...)
}

// Id is a free data retrieval call binding the contract method 0xaf640d0f.
//
// Solidity: function id() view returns(uint256)
func (_VaultProxy *VaultProxyCaller) Id(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _VaultProxy.contract.Call(opts, &out, "id")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Id is a free data retrieval call binding the contract method 0xaf640d0f.
//
// Solidity: function id() view returns(uint256)
func (_VaultProxy *VaultProxySession) Id() (*big.Int, error) {
	return _VaultProxy.Contract.Id(&_VaultProxy.CallOpts)
}

// Id is a free data retrieval call binding the contract method 0xaf640d0f.
//
// Solidity: function id() view returns(uint256)
func (_VaultProxy *VaultProxyCallerSession) Id() (*big.Int, error) {
	return _VaultProxy.Contract.Id(&_VaultProxy.CallOpts)
}

// IsValidatorWithdrawalVault is a free data retrieval call binding the contract method 0x7cc0bb90.
//
// Solidity: function isValidatorWithdrawalVault() view returns(bool)
func (_VaultProxy *VaultProxyCaller) IsValidatorWithdrawalVault(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _VaultProxy.contract.Call(opts, &out, "isValidatorWithdrawalVault")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsValidatorWithdrawalVault is a free data retrieval call binding the contract method 0x7cc0bb90.
//
// Solidity: function isValidatorWithdrawalVault() view returns(bool)
func (_VaultProxy *VaultProxySession) IsValidatorWithdrawalVault() (bool, error) {
	return _VaultProxy.Contract.IsValidatorWithdrawalVault(&_VaultProxy.CallOpts)
}

// IsValidatorWithdrawalVault is a free data retrieval call binding the contract method 0x7cc0bb90.
//
// Solidity: function isValidatorWithdrawalVault() view returns(bool)
func (_VaultProxy *VaultProxyCallerSession) IsValidatorWithdrawalVault() (bool, error) {
	return _VaultProxy.Contract.IsValidatorWithdrawalVault(&_VaultProxy.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_VaultProxy *VaultProxyCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _VaultProxy.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_VaultProxy *VaultProxySession) Owner() (common.Address, error) {
	return _VaultProxy.Contract.Owner(&_VaultProxy.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_VaultProxy *VaultProxyCallerSession) Owner() (common.Address, error) {
	return _VaultProxy.Contract.Owner(&_VaultProxy.CallOpts)
}

// PoolId is a free data retrieval call binding the contract method 0x3e0dc34e.
//
// Solidity: function poolId() view returns(uint8)
func (_VaultProxy *VaultProxyCaller) PoolId(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _VaultProxy.contract.Call(opts, &out, "poolId")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// PoolId is a free data retrieval call binding the contract method 0x3e0dc34e.
//
// Solidity: function poolId() view returns(uint8)
func (_VaultProxy *VaultProxySession) PoolId() (uint8, error) {
	return _VaultProxy.Contract.PoolId(&_VaultProxy.CallOpts)
}

// PoolId is a free data retrieval call binding the contract method 0x3e0dc34e.
//
// Solidity: function poolId() view returns(uint8)
func (_VaultProxy *VaultProxyCallerSession) PoolId() (uint8, error) {
	return _VaultProxy.Contract.PoolId(&_VaultProxy.CallOpts)
}

// StaderConfig is a free data retrieval call binding the contract method 0x490ffa35.
//
// Solidity: function staderConfig() view returns(address)
func (_VaultProxy *VaultProxyCaller) StaderConfig(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _VaultProxy.contract.Call(opts, &out, "staderConfig")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// StaderConfig is a free data retrieval call binding the contract method 0x490ffa35.
//
// Solidity: function staderConfig() view returns(address)
func (_VaultProxy *VaultProxySession) StaderConfig() (common.Address, error) {
	return _VaultProxy.Contract.StaderConfig(&_VaultProxy.CallOpts)
}

// StaderConfig is a free data retrieval call binding the contract method 0x490ffa35.
//
// Solidity: function staderConfig() view returns(address)
func (_VaultProxy *VaultProxyCallerSession) StaderConfig() (common.Address, error) {
	return _VaultProxy.Contract.StaderConfig(&_VaultProxy.CallOpts)
}

// VaultSettleStatus is a free data retrieval call binding the contract method 0x7ef4947d.
//
// Solidity: function vaultSettleStatus() view returns(bool)
func (_VaultProxy *VaultProxyCaller) VaultSettleStatus(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _VaultProxy.contract.Call(opts, &out, "vaultSettleStatus")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// VaultSettleStatus is a free data retrieval call binding the contract method 0x7ef4947d.
//
// Solidity: function vaultSettleStatus() view returns(bool)
func (_VaultProxy *VaultProxySession) VaultSettleStatus() (bool, error) {
	return _VaultProxy.Contract.VaultSettleStatus(&_VaultProxy.CallOpts)
}

// VaultSettleStatus is a free data retrieval call binding the contract method 0x7ef4947d.
//
// Solidity: function vaultSettleStatus() view returns(bool)
func (_VaultProxy *VaultProxyCallerSession) VaultSettleStatus() (bool, error) {
	return _VaultProxy.Contract.VaultSettleStatus(&_VaultProxy.CallOpts)
}

// Initialise is a paid mutator transaction binding the contract method 0x120c59e4.
//
// Solidity: function initialise(bool _isValidatorWithdrawalVault, uint8 _poolId, uint256 _id, address _staderConfig) returns()
func (_VaultProxy *VaultProxyTransactor) Initialise(opts *bind.TransactOpts, _isValidatorWithdrawalVault bool, _poolId uint8, _id *big.Int, _staderConfig common.Address) (*types.Transaction, error) {
	return _VaultProxy.contract.Transact(opts, "initialise", _isValidatorWithdrawalVault, _poolId, _id, _staderConfig)
}

// Initialise is a paid mutator transaction binding the contract method 0x120c59e4.
//
// Solidity: function initialise(bool _isValidatorWithdrawalVault, uint8 _poolId, uint256 _id, address _staderConfig) returns()
func (_VaultProxy *VaultProxySession) Initialise(_isValidatorWithdrawalVault bool, _poolId uint8, _id *big.Int, _staderConfig common.Address) (*types.Transaction, error) {
	return _VaultProxy.Contract.Initialise(&_VaultProxy.TransactOpts, _isValidatorWithdrawalVault, _poolId, _id, _staderConfig)
}

// Initialise is a paid mutator transaction binding the contract method 0x120c59e4.
//
// Solidity: function initialise(bool _isValidatorWithdrawalVault, uint8 _poolId, uint256 _id, address _staderConfig) returns()
func (_VaultProxy *VaultProxyTransactorSession) Initialise(_isValidatorWithdrawalVault bool, _poolId uint8, _id *big.Int, _staderConfig common.Address) (*types.Transaction, error) {
	return _VaultProxy.Contract.Initialise(&_VaultProxy.TransactOpts, _isValidatorWithdrawalVault, _poolId, _id, _staderConfig)
}

// UpdateOwner is a paid mutator transaction binding the contract method 0x880cdc31.
//
// Solidity: function updateOwner(address _owner) returns()
func (_VaultProxy *VaultProxyTransactor) UpdateOwner(opts *bind.TransactOpts, _owner common.Address) (*types.Transaction, error) {
	return _VaultProxy.contract.Transact(opts, "updateOwner", _owner)
}

// UpdateOwner is a paid mutator transaction binding the contract method 0x880cdc31.
//
// Solidity: function updateOwner(address _owner) returns()
func (_VaultProxy *VaultProxySession) UpdateOwner(_owner common.Address) (*types.Transaction, error) {
	return _VaultProxy.Contract.UpdateOwner(&_VaultProxy.TransactOpts, _owner)
}

// UpdateOwner is a paid mutator transaction binding the contract method 0x880cdc31.
//
// Solidity: function updateOwner(address _owner) returns()
func (_VaultProxy *VaultProxyTransactorSession) UpdateOwner(_owner common.Address) (*types.Transaction, error) {
	return _VaultProxy.Contract.UpdateOwner(&_VaultProxy.TransactOpts, _owner)
}

// UpdateStaderConfig is a paid mutator transaction binding the contract method 0x9ee804cb.
//
// Solidity: function updateStaderConfig(address _staderConfig) returns()
func (_VaultProxy *VaultProxyTransactor) UpdateStaderConfig(opts *bind.TransactOpts, _staderConfig common.Address) (*types.Transaction, error) {
	return _VaultProxy.contract.Transact(opts, "updateStaderConfig", _staderConfig)
}

// UpdateStaderConfig is a paid mutator transaction binding the contract method 0x9ee804cb.
//
// Solidity: function updateStaderConfig(address _staderConfig) returns()
func (_VaultProxy *VaultProxySession) UpdateStaderConfig(_staderConfig common.Address) (*types.Transaction, error) {
	return _VaultProxy.Contract.UpdateStaderConfig(&_VaultProxy.TransactOpts, _staderConfig)
}

// UpdateStaderConfig is a paid mutator transaction binding the contract method 0x9ee804cb.
//
// Solidity: function updateStaderConfig(address _staderConfig) returns()
func (_VaultProxy *VaultProxyTransactorSession) UpdateStaderConfig(_staderConfig common.Address) (*types.Transaction, error) {
	return _VaultProxy.Contract.UpdateStaderConfig(&_VaultProxy.TransactOpts, _staderConfig)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() payable returns()
func (_VaultProxy *VaultProxyTransactor) Fallback(opts *bind.TransactOpts, calldata []byte) (*types.Transaction, error) {
	return _VaultProxy.contract.RawTransact(opts, calldata)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() payable returns()
func (_VaultProxy *VaultProxySession) Fallback(calldata []byte) (*types.Transaction, error) {
	return _VaultProxy.Contract.Fallback(&_VaultProxy.TransactOpts, calldata)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() payable returns()
func (_VaultProxy *VaultProxyTransactorSession) Fallback(calldata []byte) (*types.Transaction, error) {
	return _VaultProxy.Contract.Fallback(&_VaultProxy.TransactOpts, calldata)
}

// VaultProxyUpdatedOwnerIterator is returned from FilterUpdatedOwner and is used to iterate over the raw logs and unpacked data for UpdatedOwner events raised by the VaultProxy contract.
type VaultProxyUpdatedOwnerIterator struct {
	Event *VaultProxyUpdatedOwner // Event containing the contract specifics and raw log

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
func (it *VaultProxyUpdatedOwnerIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VaultProxyUpdatedOwner)
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
		it.Event = new(VaultProxyUpdatedOwner)
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
func (it *VaultProxyUpdatedOwnerIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VaultProxyUpdatedOwnerIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VaultProxyUpdatedOwner represents a UpdatedOwner event raised by the VaultProxy contract.
type VaultProxyUpdatedOwner struct {
	Owner common.Address
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterUpdatedOwner is a free log retrieval operation binding the contract event 0x957090e72c0a1b3ebf83c682eb8c1f88c2a18cd0578b91a819efb28859f0f3a3.
//
// Solidity: event UpdatedOwner(address owner)
func (_VaultProxy *VaultProxyFilterer) FilterUpdatedOwner(opts *bind.FilterOpts) (*VaultProxyUpdatedOwnerIterator, error) {

	logs, sub, err := _VaultProxy.contract.FilterLogs(opts, "UpdatedOwner")
	if err != nil {
		return nil, err
	}
	return &VaultProxyUpdatedOwnerIterator{contract: _VaultProxy.contract, event: "UpdatedOwner", logs: logs, sub: sub}, nil
}

// WatchUpdatedOwner is a free log subscription operation binding the contract event 0x957090e72c0a1b3ebf83c682eb8c1f88c2a18cd0578b91a819efb28859f0f3a3.
//
// Solidity: event UpdatedOwner(address owner)
func (_VaultProxy *VaultProxyFilterer) WatchUpdatedOwner(opts *bind.WatchOpts, sink chan<- *VaultProxyUpdatedOwner) (event.Subscription, error) {

	logs, sub, err := _VaultProxy.contract.WatchLogs(opts, "UpdatedOwner")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VaultProxyUpdatedOwner)
				if err := _VaultProxy.contract.UnpackLog(event, "UpdatedOwner", log); err != nil {
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

// ParseUpdatedOwner is a log parse operation binding the contract event 0x957090e72c0a1b3ebf83c682eb8c1f88c2a18cd0578b91a819efb28859f0f3a3.
//
// Solidity: event UpdatedOwner(address owner)
func (_VaultProxy *VaultProxyFilterer) ParseUpdatedOwner(log types.Log) (*VaultProxyUpdatedOwner, error) {
	event := new(VaultProxyUpdatedOwner)
	if err := _VaultProxy.contract.UnpackLog(event, "UpdatedOwner", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// VaultProxyUpdatedStaderConfigIterator is returned from FilterUpdatedStaderConfig and is used to iterate over the raw logs and unpacked data for UpdatedStaderConfig events raised by the VaultProxy contract.
type VaultProxyUpdatedStaderConfigIterator struct {
	Event *VaultProxyUpdatedStaderConfig // Event containing the contract specifics and raw log

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
func (it *VaultProxyUpdatedStaderConfigIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VaultProxyUpdatedStaderConfig)
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
		it.Event = new(VaultProxyUpdatedStaderConfig)
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
func (it *VaultProxyUpdatedStaderConfigIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VaultProxyUpdatedStaderConfigIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VaultProxyUpdatedStaderConfig represents a UpdatedStaderConfig event raised by the VaultProxy contract.
type VaultProxyUpdatedStaderConfig struct {
	StaderConfig common.Address
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterUpdatedStaderConfig is a free log retrieval operation binding the contract event 0xdb2219043d7b197cb235f1af0cf6d782d77dee3de19e3f4fb6d39aae633b4485.
//
// Solidity: event UpdatedStaderConfig(address staderConfig)
func (_VaultProxy *VaultProxyFilterer) FilterUpdatedStaderConfig(opts *bind.FilterOpts) (*VaultProxyUpdatedStaderConfigIterator, error) {

	logs, sub, err := _VaultProxy.contract.FilterLogs(opts, "UpdatedStaderConfig")
	if err != nil {
		return nil, err
	}
	return &VaultProxyUpdatedStaderConfigIterator{contract: _VaultProxy.contract, event: "UpdatedStaderConfig", logs: logs, sub: sub}, nil
}

// WatchUpdatedStaderConfig is a free log subscription operation binding the contract event 0xdb2219043d7b197cb235f1af0cf6d782d77dee3de19e3f4fb6d39aae633b4485.
//
// Solidity: event UpdatedStaderConfig(address staderConfig)
func (_VaultProxy *VaultProxyFilterer) WatchUpdatedStaderConfig(opts *bind.WatchOpts, sink chan<- *VaultProxyUpdatedStaderConfig) (event.Subscription, error) {

	logs, sub, err := _VaultProxy.contract.WatchLogs(opts, "UpdatedStaderConfig")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VaultProxyUpdatedStaderConfig)
				if err := _VaultProxy.contract.UnpackLog(event, "UpdatedStaderConfig", log); err != nil {
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
func (_VaultProxy *VaultProxyFilterer) ParseUpdatedStaderConfig(log types.Log) (*VaultProxyUpdatedStaderConfig, error) {
	event := new(VaultProxyUpdatedStaderConfig)
	if err := _VaultProxy.contract.UnpackLog(event, "UpdatedStaderConfig", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
