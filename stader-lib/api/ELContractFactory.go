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

// ELContractFactoryMetaData contains all meta data concerning the ELContractFactory contract.
var ELContractFactoryMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_implementation\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"_pubkey\",\"type\":\"bytes\"}],\"name\":\"createProxy\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"_pubkey\",\"type\":\"bytes\"}],\"name\":\"getProxyAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"_pubkey\",\"type\":\"bytes\"}],\"name\":\"getPubkeyRoot\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"implementation\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Bin: "0x608060405234801561000f575f80fd5b506040516104c63803806104c683398101604081905261002e9161009b565b6001600160a01b0381166100775760405162461bcd60e51b815260206004820152600c60248201526b5a65726f206164647265737360a01b604482015260640160405180910390fd5b5f80546001600160a01b0319166001600160a01b03929092169190911790556100c8565b5f602082840312156100ab575f80fd5b81516001600160a01b03811681146100c1575f80fd5b9392505050565b6103f1806100d55f395ff3fe608060405234801561000f575f80fd5b506004361061004a575f3560e01c80635c60da1b1461004e5780638f295d591461007d578063c016c13714610092578063f0c418c8146100a5575b5f80fd5b5f54610060906001600160a01b031681565b6040516001600160a01b0390911681526020015b60405180910390f35b61009061008b3660046102e5565b6100c6565b005b6100606100a03660046102e5565b6100ef565b6100b86100b33660046102e5565b61011d565b604051908152602001610074565b5f6100d1838361011d565b5f549091506100e9906001600160a01b0316826101e2565b50505050565b5f806100fb848461011d565b5f54909150610113906001600160a01b031682610280565b9150505b92915050565b5f6030821461016b5760405162461bcd60e51b8152602060048201526015602482015274092dcecc2d8d2c840e0eac4d6caf240d8cadccee8d605b1b60448201526064015b60405180910390fd5b60405160029061018390859085905f90602001610351565b60408051601f198184030181529082905261019d91610378565b602060405180830381855afa1580156101b8573d5f803e3d5ffd5b5050506040513d601f19601f820116820180604052508101906101db91906103a4565b9392505050565b5f604051733d602d80600a3d3981f3363d3d373d3d3d363d7360601b81528360601b60148201526e5af43d82803e903d91602b57fd5bf360881b6028820152826037825ff59150506001600160a01b0381166101175760405162461bcd60e51b815260206004820152601760248201527f455243313136373a2063726561746532206661696c65640000000000000000006044820152606401610162565b5f6101db838330604051733d602d80600a3d3981f3363d3d373d3d3d363d7360601b8152606093841b60148201526f5af43d82803e903d91602b57fd5bf3ff60801b6028820152921b6038830152604c8201526037808220606c830152605591012090565b5f80602083850312156102f6575f80fd5b823567ffffffffffffffff8082111561030d575f80fd5b818501915085601f830112610320575f80fd5b81358181111561032e575f80fd5b86602082850101111561033f575f80fd5b60209290920196919550909350505050565b828482376fffffffffffffffffffffffffffffffff19919091169101908152601001919050565b5f82515f5b81811015610397576020818601810151858301520161037d565b505f920191825250919050565b5f602082840312156103b4575f80fd5b505191905056fea26469706673582212206d63d55da23c642bc1b07e2a4533f3e313fd9945346808e0424993d9b0a2ab4c64736f6c63430008140033",
}

// ELContractFactoryABI is the input ABI used to generate the binding from.
// Deprecated: Use ELContractFactoryMetaData.ABI instead.
var ELContractFactoryABI = ELContractFactoryMetaData.ABI

// ELContractFactoryBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use ELContractFactoryMetaData.Bin instead.
var ELContractFactoryBin = ELContractFactoryMetaData.Bin

// DeployELContractFactory deploys a new Ethereum contract, binding an instance of ELContractFactory to it.
func DeployELContractFactory(auth *bind.TransactOpts, backend bind.ContractBackend, _implementation common.Address) (common.Address, *types.Transaction, *ELContractFactory, error) {
	parsed, err := ELContractFactoryMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(ELContractFactoryBin), backend, _implementation)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &ELContractFactory{ELContractFactoryCaller: ELContractFactoryCaller{contract: contract}, ELContractFactoryTransactor: ELContractFactoryTransactor{contract: contract}, ELContractFactoryFilterer: ELContractFactoryFilterer{contract: contract}}, nil
}

// ELContractFactory is an auto generated Go binding around an Ethereum contract.
type ELContractFactory struct {
	ELContractFactoryCaller     // Read-only binding to the contract
	ELContractFactoryTransactor // Write-only binding to the contract
	ELContractFactoryFilterer   // Log filterer for contract events
}

// ELContractFactoryCaller is an auto generated read-only Go binding around an Ethereum contract.
type ELContractFactoryCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ELContractFactoryTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ELContractFactoryTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ELContractFactoryFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ELContractFactoryFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ELContractFactorySession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ELContractFactorySession struct {
	Contract     *ELContractFactory // Generic contract binding to set the session for
	CallOpts     bind.CallOpts      // Call options to use throughout this session
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// ELContractFactoryCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ELContractFactoryCallerSession struct {
	Contract *ELContractFactoryCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts            // Call options to use throughout this session
}

// ELContractFactoryTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ELContractFactoryTransactorSession struct {
	Contract     *ELContractFactoryTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts            // Transaction auth options to use throughout this session
}

// ELContractFactoryRaw is an auto generated low-level Go binding around an Ethereum contract.
type ELContractFactoryRaw struct {
	Contract *ELContractFactory // Generic contract binding to access the raw methods on
}

// ELContractFactoryCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ELContractFactoryCallerRaw struct {
	Contract *ELContractFactoryCaller // Generic read-only contract binding to access the raw methods on
}

// ELContractFactoryTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ELContractFactoryTransactorRaw struct {
	Contract *ELContractFactoryTransactor // Generic write-only contract binding to access the raw methods on
}

// NewELContractFactory creates a new instance of ELContractFactory, bound to a specific deployed contract.
func NewELContractFactory(address common.Address, backend bind.ContractBackend) (*ELContractFactory, error) {
	contract, err := bindELContractFactory(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ELContractFactory{ELContractFactoryCaller: ELContractFactoryCaller{contract: contract}, ELContractFactoryTransactor: ELContractFactoryTransactor{contract: contract}, ELContractFactoryFilterer: ELContractFactoryFilterer{contract: contract}}, nil
}

// NewELContractFactoryCaller creates a new read-only instance of ELContractFactory, bound to a specific deployed contract.
func NewELContractFactoryCaller(address common.Address, caller bind.ContractCaller) (*ELContractFactoryCaller, error) {
	contract, err := bindELContractFactory(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ELContractFactoryCaller{contract: contract}, nil
}

// NewELContractFactoryTransactor creates a new write-only instance of ELContractFactory, bound to a specific deployed contract.
func NewELContractFactoryTransactor(address common.Address, transactor bind.ContractTransactor) (*ELContractFactoryTransactor, error) {
	contract, err := bindELContractFactory(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ELContractFactoryTransactor{contract: contract}, nil
}

// NewELContractFactoryFilterer creates a new log filterer instance of ELContractFactory, bound to a specific deployed contract.
func NewELContractFactoryFilterer(address common.Address, filterer bind.ContractFilterer) (*ELContractFactoryFilterer, error) {
	contract, err := bindELContractFactory(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ELContractFactoryFilterer{contract: contract}, nil
}

// bindELContractFactory binds a generic wrapper to an already deployed contract.
func bindELContractFactory(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := ELContractFactoryMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ELContractFactory *ELContractFactoryRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ELContractFactory.Contract.ELContractFactoryCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ELContractFactory *ELContractFactoryRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ELContractFactory.Contract.ELContractFactoryTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ELContractFactory *ELContractFactoryRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ELContractFactory.Contract.ELContractFactoryTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ELContractFactory *ELContractFactoryCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ELContractFactory.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ELContractFactory *ELContractFactoryTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ELContractFactory.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ELContractFactory *ELContractFactoryTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ELContractFactory.Contract.contract.Transact(opts, method, params...)
}

// GetProxyAddress is a free data retrieval call binding the contract method 0xc016c137.
//
// Solidity: function getProxyAddress(bytes _pubkey) view returns(address)
func (_ELContractFactory *ELContractFactoryCaller) GetProxyAddress(opts *bind.CallOpts, _pubkey []byte) (common.Address, error) {
	var out []interface{}
	err := _ELContractFactory.contract.Call(opts, &out, "getProxyAddress", _pubkey)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetProxyAddress is a free data retrieval call binding the contract method 0xc016c137.
//
// Solidity: function getProxyAddress(bytes _pubkey) view returns(address)
func (_ELContractFactory *ELContractFactorySession) GetProxyAddress(_pubkey []byte) (common.Address, error) {
	return _ELContractFactory.Contract.GetProxyAddress(&_ELContractFactory.CallOpts, _pubkey)
}

// GetProxyAddress is a free data retrieval call binding the contract method 0xc016c137.
//
// Solidity: function getProxyAddress(bytes _pubkey) view returns(address)
func (_ELContractFactory *ELContractFactoryCallerSession) GetProxyAddress(_pubkey []byte) (common.Address, error) {
	return _ELContractFactory.Contract.GetProxyAddress(&_ELContractFactory.CallOpts, _pubkey)
}

// GetPubkeyRoot is a free data retrieval call binding the contract method 0xf0c418c8.
//
// Solidity: function getPubkeyRoot(bytes _pubkey) pure returns(bytes32)
func (_ELContractFactory *ELContractFactoryCaller) GetPubkeyRoot(opts *bind.CallOpts, _pubkey []byte) ([32]byte, error) {
	var out []interface{}
	err := _ELContractFactory.contract.Call(opts, &out, "getPubkeyRoot", _pubkey)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetPubkeyRoot is a free data retrieval call binding the contract method 0xf0c418c8.
//
// Solidity: function getPubkeyRoot(bytes _pubkey) pure returns(bytes32)
func (_ELContractFactory *ELContractFactorySession) GetPubkeyRoot(_pubkey []byte) ([32]byte, error) {
	return _ELContractFactory.Contract.GetPubkeyRoot(&_ELContractFactory.CallOpts, _pubkey)
}

// GetPubkeyRoot is a free data retrieval call binding the contract method 0xf0c418c8.
//
// Solidity: function getPubkeyRoot(bytes _pubkey) pure returns(bytes32)
func (_ELContractFactory *ELContractFactoryCallerSession) GetPubkeyRoot(_pubkey []byte) ([32]byte, error) {
	return _ELContractFactory.Contract.GetPubkeyRoot(&_ELContractFactory.CallOpts, _pubkey)
}

// Implementation is a free data retrieval call binding the contract method 0x5c60da1b.
//
// Solidity: function implementation() view returns(address)
func (_ELContractFactory *ELContractFactoryCaller) Implementation(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ELContractFactory.contract.Call(opts, &out, "implementation")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Implementation is a free data retrieval call binding the contract method 0x5c60da1b.
//
// Solidity: function implementation() view returns(address)
func (_ELContractFactory *ELContractFactorySession) Implementation() (common.Address, error) {
	return _ELContractFactory.Contract.Implementation(&_ELContractFactory.CallOpts)
}

// Implementation is a free data retrieval call binding the contract method 0x5c60da1b.
//
// Solidity: function implementation() view returns(address)
func (_ELContractFactory *ELContractFactoryCallerSession) Implementation() (common.Address, error) {
	return _ELContractFactory.Contract.Implementation(&_ELContractFactory.CallOpts)
}

// CreateProxy is a paid mutator transaction binding the contract method 0x8f295d59.
//
// Solidity: function createProxy(bytes _pubkey) returns()
func (_ELContractFactory *ELContractFactoryTransactor) CreateProxy(opts *bind.TransactOpts, _pubkey []byte) (*types.Transaction, error) {
	return _ELContractFactory.contract.Transact(opts, "createProxy", _pubkey)
}

// CreateProxy is a paid mutator transaction binding the contract method 0x8f295d59.
//
// Solidity: function createProxy(bytes _pubkey) returns()
func (_ELContractFactory *ELContractFactorySession) CreateProxy(_pubkey []byte) (*types.Transaction, error) {
	return _ELContractFactory.Contract.CreateProxy(&_ELContractFactory.TransactOpts, _pubkey)
}

// CreateProxy is a paid mutator transaction binding the contract method 0x8f295d59.
//
// Solidity: function createProxy(bytes _pubkey) returns()
func (_ELContractFactory *ELContractFactoryTransactorSession) CreateProxy(_pubkey []byte) (*types.Transaction, error) {
	return _ELContractFactory.Contract.CreateProxy(&_ELContractFactory.TransactOpts, _pubkey)
}
