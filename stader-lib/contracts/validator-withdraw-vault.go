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
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"CallerNotNodeRegistryContract\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidRewardAmount\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotEnoughRewardToDistribute\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"TransferFailed\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"rewardAmount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"rewardThreshold\",\"type\":\"uint256\"}],\"name\":\"DistributeRewardFailed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"userShare\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"operatorShare\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"protocolShare\",\"type\":\"uint256\"}],\"name\":\"DistributedRewards\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"ETHReceived\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"userShare\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"operatorShare\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"protocolShare\",\"type\":\"uint256\"}],\"name\":\"SettledFunds\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_staderConfig\",\"type\":\"address\"}],\"name\":\"UpdatedStaderConfig\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"calculateValidatorWithdrawalShare\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"_userShare\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_operatorShare\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_protocolShare\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"distributeRewards\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"settleFunds\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"stateMutability\":\"payable\",\"type\":\"receive\"}]",
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
