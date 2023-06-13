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

// ELContractMetaData contains all meta data concerning the ELContract contract.
var ELContractMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_storageContract\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"stateMutability\":\"payable\",\"type\":\"fallback\"},{\"inputs\":[],\"name\":\"elContractDelegateKey\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"storageContract\",\"outputs\":[{\"internalType\":\"contractIStorage\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"stateMutability\":\"payable\",\"type\":\"receive\"}]",
	Bin: "0x60c060405234801561000f575f80fd5b506040516102df3803806102df83398101604081905261002e91610096565b6001600160a01b03811660a0526040516f636f6e74726163742e6164647265737360801b602082015271454c436f6e747261637444656c656761746560701b603082015260420160408051601f198184030181529190528051602090910120608052506100c3565b5f602082840312156100a6575f80fd5b81516001600160a01b03811681146100bc575f80fd5b9392505050565b60805160a0516101ef6100f05f395f8181606a015261010c01525f81816042015261015c01526101ef5ff3fe60806040526004361061002c575f3560e01c806311ce0267146100fb578063638927951461014b57610033565b3661003357005b6040516321f8a72160e01b81527f000000000000000000000000000000000000000000000000000000000000000060048201525f907f00000000000000000000000000000000000000000000000000000000000000006001600160a01b0316906321f8a72190602401602060405180830381865afa1580156100b7573d5f803e3d5ffd5b505050506040513d601f19601f820116820180604052508101906100db919061018c565b9050365f80375f80365f845af43d5f803e8080156100f7573d5ff35b3d5ffd5b348015610106575f80fd5b5061012e7f000000000000000000000000000000000000000000000000000000000000000081565b6040516001600160a01b0390911681526020015b60405180910390f35b348015610156575f80fd5b5061017e7f000000000000000000000000000000000000000000000000000000000000000081565b604051908152602001610142565b5f6020828403121561019c575f80fd5b81516001600160a01b03811681146101b2575f80fd5b939250505056fea2646970667358221220d756c64b850292594ad4eb5deb84983dac4e9ece2dff895b5ef38ce2efe23f7664736f6c63430008140033",
}

// ELContractABI is the input ABI used to generate the binding from.
// Deprecated: Use ELContractMetaData.ABI instead.
var ELContractABI = ELContractMetaData.ABI

// ELContractBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use ELContractMetaData.Bin instead.
var ELContractBin = ELContractMetaData.Bin

// DeployELContract deploys a new Ethereum contract, binding an instance of ELContract to it.
func DeployELContract(auth *bind.TransactOpts, backend bind.ContractBackend, _storageContract common.Address) (common.Address, *types.Transaction, *ELContract, error) {
	parsed, err := ELContractMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(ELContractBin), backend, _storageContract)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &ELContract{ELContractCaller: ELContractCaller{contract: contract}, ELContractTransactor: ELContractTransactor{contract: contract}, ELContractFilterer: ELContractFilterer{contract: contract}}, nil
}

// ELContract is an auto generated Go binding around an Ethereum contract.
type ELContract struct {
	ELContractCaller     // Read-only binding to the contract
	ELContractTransactor // Write-only binding to the contract
	ELContractFilterer   // Log filterer for contract events
}

// ELContractCaller is an auto generated read-only Go binding around an Ethereum contract.
type ELContractCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ELContractTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ELContractTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ELContractFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ELContractFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ELContractSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ELContractSession struct {
	Contract     *ELContract       // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ELContractCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ELContractCallerSession struct {
	Contract *ELContractCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts     // Call options to use throughout this session
}

// ELContractTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ELContractTransactorSession struct {
	Contract     *ELContractTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// ELContractRaw is an auto generated low-level Go binding around an Ethereum contract.
type ELContractRaw struct {
	Contract *ELContract // Generic contract binding to access the raw methods on
}

// ELContractCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ELContractCallerRaw struct {
	Contract *ELContractCaller // Generic read-only contract binding to access the raw methods on
}

// ELContractTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ELContractTransactorRaw struct {
	Contract *ELContractTransactor // Generic write-only contract binding to access the raw methods on
}

// NewELContract creates a new instance of ELContract, bound to a specific deployed contract.
func NewELContract(address common.Address, backend bind.ContractBackend) (*ELContract, error) {
	contract, err := bindELContract(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ELContract{ELContractCaller: ELContractCaller{contract: contract}, ELContractTransactor: ELContractTransactor{contract: contract}, ELContractFilterer: ELContractFilterer{contract: contract}}, nil
}

// NewELContractCaller creates a new read-only instance of ELContract, bound to a specific deployed contract.
func NewELContractCaller(address common.Address, caller bind.ContractCaller) (*ELContractCaller, error) {
	contract, err := bindELContract(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ELContractCaller{contract: contract}, nil
}

// NewELContractTransactor creates a new write-only instance of ELContract, bound to a specific deployed contract.
func NewELContractTransactor(address common.Address, transactor bind.ContractTransactor) (*ELContractTransactor, error) {
	contract, err := bindELContract(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ELContractTransactor{contract: contract}, nil
}

// NewELContractFilterer creates a new log filterer instance of ELContract, bound to a specific deployed contract.
func NewELContractFilterer(address common.Address, filterer bind.ContractFilterer) (*ELContractFilterer, error) {
	contract, err := bindELContract(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ELContractFilterer{contract: contract}, nil
}

// bindELContract binds a generic wrapper to an already deployed contract.
func bindELContract(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := ELContractMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ELContract *ELContractRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ELContract.Contract.ELContractCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ELContract *ELContractRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ELContract.Contract.ELContractTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ELContract *ELContractRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ELContract.Contract.ELContractTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ELContract *ELContractCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ELContract.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ELContract *ELContractTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ELContract.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ELContract *ELContractTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ELContract.Contract.contract.Transact(opts, method, params...)
}

// ElContractDelegateKey is a free data retrieval call binding the contract method 0x63892795.
//
// Solidity: function elContractDelegateKey() view returns(bytes32)
func (_ELContract *ELContractCaller) ElContractDelegateKey(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _ELContract.contract.Call(opts, &out, "elContractDelegateKey")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// ElContractDelegateKey is a free data retrieval call binding the contract method 0x63892795.
//
// Solidity: function elContractDelegateKey() view returns(bytes32)
func (_ELContract *ELContractSession) ElContractDelegateKey() ([32]byte, error) {
	return _ELContract.Contract.ElContractDelegateKey(&_ELContract.CallOpts)
}

// ElContractDelegateKey is a free data retrieval call binding the contract method 0x63892795.
//
// Solidity: function elContractDelegateKey() view returns(bytes32)
func (_ELContract *ELContractCallerSession) ElContractDelegateKey() ([32]byte, error) {
	return _ELContract.Contract.ElContractDelegateKey(&_ELContract.CallOpts)
}

// StorageContract is a free data retrieval call binding the contract method 0x11ce0267.
//
// Solidity: function storageContract() view returns(address)
func (_ELContract *ELContractCaller) StorageContract(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ELContract.contract.Call(opts, &out, "storageContract")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// StorageContract is a free data retrieval call binding the contract method 0x11ce0267.
//
// Solidity: function storageContract() view returns(address)
func (_ELContract *ELContractSession) StorageContract() (common.Address, error) {
	return _ELContract.Contract.StorageContract(&_ELContract.CallOpts)
}

// StorageContract is a free data retrieval call binding the contract method 0x11ce0267.
//
// Solidity: function storageContract() view returns(address)
func (_ELContract *ELContractCallerSession) StorageContract() (common.Address, error) {
	return _ELContract.Contract.StorageContract(&_ELContract.CallOpts)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() payable returns()
func (_ELContract *ELContractTransactor) Fallback(opts *bind.TransactOpts, calldata []byte) (*types.Transaction, error) {
	return _ELContract.contract.RawTransact(opts, calldata)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() payable returns()
func (_ELContract *ELContractSession) Fallback(calldata []byte) (*types.Transaction, error) {
	return _ELContract.Contract.Fallback(&_ELContract.TransactOpts, calldata)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() payable returns()
func (_ELContract *ELContractTransactorSession) Fallback(calldata []byte) (*types.Transaction, error) {
	return _ELContract.Contract.Fallback(&_ELContract.TransactOpts, calldata)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_ELContract *ELContractTransactor) Receive(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ELContract.contract.RawTransact(opts, nil) // calldata is disallowed for receive function
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_ELContract *ELContractSession) Receive() (*types.Transaction, error) {
	return _ELContract.Contract.Receive(&_ELContract.TransactOpts)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_ELContract *ELContractTransactorSession) Receive() (*types.Transaction, error) {
	return _ELContract.Contract.Receive(&_ELContract.TransactOpts)
}
