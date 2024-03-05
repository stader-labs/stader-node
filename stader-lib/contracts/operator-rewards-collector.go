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
	ABI: "[{\"inputs\":[],\"name\":\"InsufficientBalance\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"WethTransferFailed\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Claimed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"DepositedFor\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"staderConfig\",\"type\":\"address\"}],\"name\":\"UpdatedStaderConfig\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"weth\",\"type\":\"address\"}],\"name\":\"UpdatedWethAddress\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"claim\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"}],\"name\":\"claimLiquidation\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"claimWithAmount\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_receiver\",\"type\":\"address\"}],\"name\":\"depositFor\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"}],\"name\":\"getBalance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"}],\"name\":\"withdrawableInEth\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
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

// GetBalance is a free data retrieval call binding the contract method 0xf8b2cb4f.
//
// Solidity: function getBalance(address operator) view returns(uint256)
func (_OperatorRewardsCollector *OperatorRewardsCollectorCaller) GetBalance(opts *bind.CallOpts, operator common.Address) (*big.Int, error) {
	var out []interface{}
	err := _OperatorRewardsCollector.contract.Call(opts, &out, "getBalance", operator)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetBalance is a free data retrieval call binding the contract method 0xf8b2cb4f.
//
// Solidity: function getBalance(address operator) view returns(uint256)
func (_OperatorRewardsCollector *OperatorRewardsCollectorSession) GetBalance(operator common.Address) (*big.Int, error) {
	return _OperatorRewardsCollector.Contract.GetBalance(&_OperatorRewardsCollector.CallOpts, operator)
}

// GetBalance is a free data retrieval call binding the contract method 0xf8b2cb4f.
//
// Solidity: function getBalance(address operator) view returns(uint256)
func (_OperatorRewardsCollector *OperatorRewardsCollectorCallerSession) GetBalance(operator common.Address) (*big.Int, error) {
	return _OperatorRewardsCollector.Contract.GetBalance(&_OperatorRewardsCollector.CallOpts, operator)
}

// WithdrawableInEth is a free data retrieval call binding the contract method 0x96198d0f.
//
// Solidity: function withdrawableInEth(address operator) view returns(uint256)
func (_OperatorRewardsCollector *OperatorRewardsCollectorCaller) WithdrawableInEth(opts *bind.CallOpts, operator common.Address) (*big.Int, error) {
	var out []interface{}
	err := _OperatorRewardsCollector.contract.Call(opts, &out, "withdrawableInEth", operator)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// WithdrawableInEth is a free data retrieval call binding the contract method 0x96198d0f.
//
// Solidity: function withdrawableInEth(address operator) view returns(uint256)
func (_OperatorRewardsCollector *OperatorRewardsCollectorSession) WithdrawableInEth(operator common.Address) (*big.Int, error) {
	return _OperatorRewardsCollector.Contract.WithdrawableInEth(&_OperatorRewardsCollector.CallOpts, operator)
}

// WithdrawableInEth is a free data retrieval call binding the contract method 0x96198d0f.
//
// Solidity: function withdrawableInEth(address operator) view returns(uint256)
func (_OperatorRewardsCollector *OperatorRewardsCollectorCallerSession) WithdrawableInEth(operator common.Address) (*big.Int, error) {
	return _OperatorRewardsCollector.Contract.WithdrawableInEth(&_OperatorRewardsCollector.CallOpts, operator)
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

// ClaimLiquidation is a paid mutator transaction binding the contract method 0x1cc1c626.
//
// Solidity: function claimLiquidation(address operator) returns()
func (_OperatorRewardsCollector *OperatorRewardsCollectorTransactor) ClaimLiquidation(opts *bind.TransactOpts, operator common.Address) (*types.Transaction, error) {
	return _OperatorRewardsCollector.contract.Transact(opts, "claimLiquidation", operator)
}

// ClaimLiquidation is a paid mutator transaction binding the contract method 0x1cc1c626.
//
// Solidity: function claimLiquidation(address operator) returns()
func (_OperatorRewardsCollector *OperatorRewardsCollectorSession) ClaimLiquidation(operator common.Address) (*types.Transaction, error) {
	return _OperatorRewardsCollector.Contract.ClaimLiquidation(&_OperatorRewardsCollector.TransactOpts, operator)
}

// ClaimLiquidation is a paid mutator transaction binding the contract method 0x1cc1c626.
//
// Solidity: function claimLiquidation(address operator) returns()
func (_OperatorRewardsCollector *OperatorRewardsCollectorTransactorSession) ClaimLiquidation(operator common.Address) (*types.Transaction, error) {
	return _OperatorRewardsCollector.Contract.ClaimLiquidation(&_OperatorRewardsCollector.TransactOpts, operator)
}

// ClaimWithAmount is a paid mutator transaction binding the contract method 0xa4557a11.
//
// Solidity: function claimWithAmount(uint256 _amount) returns()
func (_OperatorRewardsCollector *OperatorRewardsCollectorTransactor) ClaimWithAmount(opts *bind.TransactOpts, _amount *big.Int) (*types.Transaction, error) {
	return _OperatorRewardsCollector.contract.Transact(opts, "claimWithAmount", _amount)
}

// ClaimWithAmount is a paid mutator transaction binding the contract method 0xa4557a11.
//
// Solidity: function claimWithAmount(uint256 _amount) returns()
func (_OperatorRewardsCollector *OperatorRewardsCollectorSession) ClaimWithAmount(_amount *big.Int) (*types.Transaction, error) {
	return _OperatorRewardsCollector.Contract.ClaimWithAmount(&_OperatorRewardsCollector.TransactOpts, _amount)
}

// ClaimWithAmount is a paid mutator transaction binding the contract method 0xa4557a11.
//
// Solidity: function claimWithAmount(uint256 _amount) returns()
func (_OperatorRewardsCollector *OperatorRewardsCollectorTransactorSession) ClaimWithAmount(_amount *big.Int) (*types.Transaction, error) {
	return _OperatorRewardsCollector.Contract.ClaimWithAmount(&_OperatorRewardsCollector.TransactOpts, _amount)
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

// OperatorRewardsCollectorUpdatedWethAddressIterator is returned from FilterUpdatedWethAddress and is used to iterate over the raw logs and unpacked data for UpdatedWethAddress events raised by the OperatorRewardsCollector contract.
type OperatorRewardsCollectorUpdatedWethAddressIterator struct {
	Event *OperatorRewardsCollectorUpdatedWethAddress // Event containing the contract specifics and raw log

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
func (it *OperatorRewardsCollectorUpdatedWethAddressIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OperatorRewardsCollectorUpdatedWethAddress)
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
		it.Event = new(OperatorRewardsCollectorUpdatedWethAddress)
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
func (it *OperatorRewardsCollectorUpdatedWethAddressIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OperatorRewardsCollectorUpdatedWethAddressIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OperatorRewardsCollectorUpdatedWethAddress represents a UpdatedWethAddress event raised by the OperatorRewardsCollector contract.
type OperatorRewardsCollectorUpdatedWethAddress struct {
	Weth common.Address
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterUpdatedWethAddress is a free log retrieval operation binding the contract event 0xa1de212dff029f064e368842c50f600a79a6958de54219159b77f9fa9b84e023.
//
// Solidity: event UpdatedWethAddress(address indexed weth)
func (_OperatorRewardsCollector *OperatorRewardsCollectorFilterer) FilterUpdatedWethAddress(opts *bind.FilterOpts, weth []common.Address) (*OperatorRewardsCollectorUpdatedWethAddressIterator, error) {

	var wethRule []interface{}
	for _, wethItem := range weth {
		wethRule = append(wethRule, wethItem)
	}

	logs, sub, err := _OperatorRewardsCollector.contract.FilterLogs(opts, "UpdatedWethAddress", wethRule)
	if err != nil {
		return nil, err
	}
	return &OperatorRewardsCollectorUpdatedWethAddressIterator{contract: _OperatorRewardsCollector.contract, event: "UpdatedWethAddress", logs: logs, sub: sub}, nil
}

// WatchUpdatedWethAddress is a free log subscription operation binding the contract event 0xa1de212dff029f064e368842c50f600a79a6958de54219159b77f9fa9b84e023.
//
// Solidity: event UpdatedWethAddress(address indexed weth)
func (_OperatorRewardsCollector *OperatorRewardsCollectorFilterer) WatchUpdatedWethAddress(opts *bind.WatchOpts, sink chan<- *OperatorRewardsCollectorUpdatedWethAddress, weth []common.Address) (event.Subscription, error) {

	var wethRule []interface{}
	for _, wethItem := range weth {
		wethRule = append(wethRule, wethItem)
	}

	logs, sub, err := _OperatorRewardsCollector.contract.WatchLogs(opts, "UpdatedWethAddress", wethRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OperatorRewardsCollectorUpdatedWethAddress)
				if err := _OperatorRewardsCollector.contract.UnpackLog(event, "UpdatedWethAddress", log); err != nil {
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

// ParseUpdatedWethAddress is a log parse operation binding the contract event 0xa1de212dff029f064e368842c50f600a79a6958de54219159b77f9fa9b84e023.
//
// Solidity: event UpdatedWethAddress(address indexed weth)
func (_OperatorRewardsCollector *OperatorRewardsCollectorFilterer) ParseUpdatedWethAddress(log types.Log) (*OperatorRewardsCollectorUpdatedWethAddress, error) {
	event := new(OperatorRewardsCollectorUpdatedWethAddress)
	if err := _OperatorRewardsCollector.contract.UnpackLog(event, "UpdatedWethAddress", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
