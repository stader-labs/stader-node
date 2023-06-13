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
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"ETHTransferFailed\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotEnoughRewardToWithdraw\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"TransferFailed\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"ETHReceived\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"staderConfig\",\"type\":\"address\"}],\"name\":\"UpdatedStaderConfig\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"protocolAmount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"operatorAmount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"userAmount\",\"type\":\"uint256\"}],\"name\":\"Withdrawal\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"withdraw\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"stateMutability\":\"payable\",\"type\":\"receive\"}]",
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
