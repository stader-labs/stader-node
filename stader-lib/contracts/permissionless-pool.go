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

// PermissionlessPoolMetaData contains all meta data concerning the PermissionlessPool contract.
var PermissionlessPoolMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"CallerNotManager\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"CallerNotStaderContract\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidCommission\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OperatorFeeUnchanged\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ProtocolFeeUnchanged\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"UnsupportedOperation\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ZeroAddress\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"version\",\"type\":\"uint8\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"ReceivedCollateralETH\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"ReceivedInsuranceFund\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"previousAdminRole\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"newAdminRole\",\"type\":\"bytes32\"}],\"name\":\"RoleAdminChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleGranted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleRevoked\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"TransferredETHToSSPMForDefectiveKeys\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"protocolFee\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"operatorFee\",\"type\":\"uint256\"}],\"name\":\"UpdatedCommissionFees\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"staderConfig\",\"type\":\"address\"}],\"name\":\"UpdatedStaderConfig\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"validatorId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"pubKey\",\"type\":\"bytes\"}],\"name\":\"ValidatorDepositedOnBeaconChain\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes\",\"name\":\"pubKey\",\"type\":\"bytes\"}],\"name\":\"ValidatorPreDepositedOnBeaconChain\",\"type\":\"event\"},{\"stateMutability\":\"payable\",\"type\":\"fallback\"},{\"inputs\":[],\"name\":\"DEFAULT_ADMIN_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"DEPOSIT_NODE_BOND\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"MAX_COMMISSION_LIMIT_BIPS\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"POOL_ID\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"_pubkey\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"_signature\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"_withdrawCredential\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"_depositAmount\",\"type\":\"uint256\"}],\"name\":\"computeDepositDataRoot\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_pageNumber\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_pageSize\",\"type\":\"uint256\"}],\"name\":\"getAllSocializingPoolOptOutOperators\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getCollateralETH\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getNodeRegistry\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_nodeOperator\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_startIndex\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_endIndex\",\"type\":\"uint256\"}],\"name\":\"getOperatorTotalNonTerminalKeys\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"}],\"name\":\"getRoleAdmin\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getSocializingPoolAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getTotalActiveValidatorCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getTotalQueuedValidatorCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"grantRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"hasRole\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_admin\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_staderConfig\",\"type\":\"address\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_operAddr\",\"type\":\"address\"}],\"name\":\"isExistingOperator\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"_pubkey\",\"type\":\"bytes\"}],\"name\":\"isExistingPubkey\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"operatorFee\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes[]\",\"name\":\"_pubkey\",\"type\":\"bytes[]\"},{\"internalType\":\"bytes[]\",\"name\":\"_preDepositSignature\",\"type\":\"bytes[]\"},{\"internalType\":\"uint256\",\"name\":\"_operatorId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_operatorTotalKeys\",\"type\":\"uint256\"}],\"name\":\"preDepositOnBeaconChain\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"protocolFee\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"receiveRemainingCollateralETH\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"renounceRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"revokeRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_protocolFee\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_operatorFee\",\"type\":\"uint256\"}],\"name\":\"setCommissionFees\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"staderConfig\",\"outputs\":[{\"internalType\":\"contractIStaderConfig\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"stakeUserETHToBeaconChain\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_staderConfig\",\"type\":\"address\"}],\"name\":\"updateStaderConfig\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"stateMutability\":\"payable\",\"type\":\"receive\"}]",
}

// PermissionlessPoolABI is the input ABI used to generate the binding from.
// Deprecated: Use PermissionlessPoolMetaData.ABI instead.
var PermissionlessPoolABI = PermissionlessPoolMetaData.ABI

// PermissionlessPool is an auto generated Go binding around an Ethereum contract.
type PermissionlessPool struct {
	PermissionlessPoolCaller     // Read-only binding to the contract
	PermissionlessPoolTransactor // Write-only binding to the contract
	PermissionlessPoolFilterer   // Log filterer for contract events
}

// PermissionlessPoolCaller is an auto generated read-only Go binding around an Ethereum contract.
type PermissionlessPoolCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PermissionlessPoolTransactor is an auto generated write-only Go binding around an Ethereum contract.
type PermissionlessPoolTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PermissionlessPoolFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type PermissionlessPoolFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PermissionlessPoolSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type PermissionlessPoolSession struct {
	Contract     *PermissionlessPool // Generic contract binding to set the session for
	CallOpts     bind.CallOpts       // Call options to use throughout this session
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// PermissionlessPoolCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type PermissionlessPoolCallerSession struct {
	Contract *PermissionlessPoolCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts             // Call options to use throughout this session
}

// PermissionlessPoolTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type PermissionlessPoolTransactorSession struct {
	Contract     *PermissionlessPoolTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts             // Transaction auth options to use throughout this session
}

// PermissionlessPoolRaw is an auto generated low-level Go binding around an Ethereum contract.
type PermissionlessPoolRaw struct {
	Contract *PermissionlessPool // Generic contract binding to access the raw methods on
}

// PermissionlessPoolCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type PermissionlessPoolCallerRaw struct {
	Contract *PermissionlessPoolCaller // Generic read-only contract binding to access the raw methods on
}

// PermissionlessPoolTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type PermissionlessPoolTransactorRaw struct {
	Contract *PermissionlessPoolTransactor // Generic write-only contract binding to access the raw methods on
}

// NewPermissionlessPool creates a new instance of PermissionlessPool, bound to a specific deployed contract.
func NewPermissionlessPool(address common.Address, backend bind.ContractBackend) (*PermissionlessPool, error) {
	contract, err := bindPermissionlessPool(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &PermissionlessPool{PermissionlessPoolCaller: PermissionlessPoolCaller{contract: contract}, PermissionlessPoolTransactor: PermissionlessPoolTransactor{contract: contract}, PermissionlessPoolFilterer: PermissionlessPoolFilterer{contract: contract}}, nil
}

// NewPermissionlessPoolCaller creates a new read-only instance of PermissionlessPool, bound to a specific deployed contract.
func NewPermissionlessPoolCaller(address common.Address, caller bind.ContractCaller) (*PermissionlessPoolCaller, error) {
	contract, err := bindPermissionlessPool(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &PermissionlessPoolCaller{contract: contract}, nil
}

// NewPermissionlessPoolTransactor creates a new write-only instance of PermissionlessPool, bound to a specific deployed contract.
func NewPermissionlessPoolTransactor(address common.Address, transactor bind.ContractTransactor) (*PermissionlessPoolTransactor, error) {
	contract, err := bindPermissionlessPool(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &PermissionlessPoolTransactor{contract: contract}, nil
}

// NewPermissionlessPoolFilterer creates a new log filterer instance of PermissionlessPool, bound to a specific deployed contract.
func NewPermissionlessPoolFilterer(address common.Address, filterer bind.ContractFilterer) (*PermissionlessPoolFilterer, error) {
	contract, err := bindPermissionlessPool(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &PermissionlessPoolFilterer{contract: contract}, nil
}

// bindPermissionlessPool binds a generic wrapper to an already deployed contract.
func bindPermissionlessPool(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := PermissionlessPoolMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_PermissionlessPool *PermissionlessPoolRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _PermissionlessPool.Contract.PermissionlessPoolCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_PermissionlessPool *PermissionlessPoolRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PermissionlessPool.Contract.PermissionlessPoolTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_PermissionlessPool *PermissionlessPoolRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _PermissionlessPool.Contract.PermissionlessPoolTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_PermissionlessPool *PermissionlessPoolCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _PermissionlessPool.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_PermissionlessPool *PermissionlessPoolTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PermissionlessPool.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_PermissionlessPool *PermissionlessPoolTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _PermissionlessPool.Contract.contract.Transact(opts, method, params...)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_PermissionlessPool *PermissionlessPoolCaller) DEFAULTADMINROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _PermissionlessPool.contract.Call(opts, &out, "DEFAULT_ADMIN_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_PermissionlessPool *PermissionlessPoolSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _PermissionlessPool.Contract.DEFAULTADMINROLE(&_PermissionlessPool.CallOpts)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_PermissionlessPool *PermissionlessPoolCallerSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _PermissionlessPool.Contract.DEFAULTADMINROLE(&_PermissionlessPool.CallOpts)
}

// DEPOSITNODEBOND is a free data retrieval call binding the contract method 0x24f69706.
//
// Solidity: function DEPOSIT_NODE_BOND() view returns(uint256)
func (_PermissionlessPool *PermissionlessPoolCaller) DEPOSITNODEBOND(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _PermissionlessPool.contract.Call(opts, &out, "DEPOSIT_NODE_BOND")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// DEPOSITNODEBOND is a free data retrieval call binding the contract method 0x24f69706.
//
// Solidity: function DEPOSIT_NODE_BOND() view returns(uint256)
func (_PermissionlessPool *PermissionlessPoolSession) DEPOSITNODEBOND() (*big.Int, error) {
	return _PermissionlessPool.Contract.DEPOSITNODEBOND(&_PermissionlessPool.CallOpts)
}

// DEPOSITNODEBOND is a free data retrieval call binding the contract method 0x24f69706.
//
// Solidity: function DEPOSIT_NODE_BOND() view returns(uint256)
func (_PermissionlessPool *PermissionlessPoolCallerSession) DEPOSITNODEBOND() (*big.Int, error) {
	return _PermissionlessPool.Contract.DEPOSITNODEBOND(&_PermissionlessPool.CallOpts)
}

// MAXCOMMISSIONLIMITBIPS is a free data retrieval call binding the contract method 0x9cd6dd56.
//
// Solidity: function MAX_COMMISSION_LIMIT_BIPS() view returns(uint256)
func (_PermissionlessPool *PermissionlessPoolCaller) MAXCOMMISSIONLIMITBIPS(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _PermissionlessPool.contract.Call(opts, &out, "MAX_COMMISSION_LIMIT_BIPS")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MAXCOMMISSIONLIMITBIPS is a free data retrieval call binding the contract method 0x9cd6dd56.
//
// Solidity: function MAX_COMMISSION_LIMIT_BIPS() view returns(uint256)
func (_PermissionlessPool *PermissionlessPoolSession) MAXCOMMISSIONLIMITBIPS() (*big.Int, error) {
	return _PermissionlessPool.Contract.MAXCOMMISSIONLIMITBIPS(&_PermissionlessPool.CallOpts)
}

// MAXCOMMISSIONLIMITBIPS is a free data retrieval call binding the contract method 0x9cd6dd56.
//
// Solidity: function MAX_COMMISSION_LIMIT_BIPS() view returns(uint256)
func (_PermissionlessPool *PermissionlessPoolCallerSession) MAXCOMMISSIONLIMITBIPS() (*big.Int, error) {
	return _PermissionlessPool.Contract.MAXCOMMISSIONLIMITBIPS(&_PermissionlessPool.CallOpts)
}

// POOLID is a free data retrieval call binding the contract method 0xe0d7d0e9.
//
// Solidity: function POOL_ID() view returns(uint8)
func (_PermissionlessPool *PermissionlessPoolCaller) POOLID(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _PermissionlessPool.contract.Call(opts, &out, "POOL_ID")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// POOLID is a free data retrieval call binding the contract method 0xe0d7d0e9.
//
// Solidity: function POOL_ID() view returns(uint8)
func (_PermissionlessPool *PermissionlessPoolSession) POOLID() (uint8, error) {
	return _PermissionlessPool.Contract.POOLID(&_PermissionlessPool.CallOpts)
}

// POOLID is a free data retrieval call binding the contract method 0xe0d7d0e9.
//
// Solidity: function POOL_ID() view returns(uint8)
func (_PermissionlessPool *PermissionlessPoolCallerSession) POOLID() (uint8, error) {
	return _PermissionlessPool.Contract.POOLID(&_PermissionlessPool.CallOpts)
}

// ComputeDepositDataRoot is a free data retrieval call binding the contract method 0x5c164e53.
//
// Solidity: function computeDepositDataRoot(bytes _pubkey, bytes _signature, bytes _withdrawCredential, uint256 _depositAmount) pure returns(bytes32)
func (_PermissionlessPool *PermissionlessPoolCaller) ComputeDepositDataRoot(opts *bind.CallOpts, _pubkey []byte, _signature []byte, _withdrawCredential []byte, _depositAmount *big.Int) ([32]byte, error) {
	var out []interface{}
	err := _PermissionlessPool.contract.Call(opts, &out, "computeDepositDataRoot", _pubkey, _signature, _withdrawCredential, _depositAmount)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// ComputeDepositDataRoot is a free data retrieval call binding the contract method 0x5c164e53.
//
// Solidity: function computeDepositDataRoot(bytes _pubkey, bytes _signature, bytes _withdrawCredential, uint256 _depositAmount) pure returns(bytes32)
func (_PermissionlessPool *PermissionlessPoolSession) ComputeDepositDataRoot(_pubkey []byte, _signature []byte, _withdrawCredential []byte, _depositAmount *big.Int) ([32]byte, error) {
	return _PermissionlessPool.Contract.ComputeDepositDataRoot(&_PermissionlessPool.CallOpts, _pubkey, _signature, _withdrawCredential, _depositAmount)
}

// ComputeDepositDataRoot is a free data retrieval call binding the contract method 0x5c164e53.
//
// Solidity: function computeDepositDataRoot(bytes _pubkey, bytes _signature, bytes _withdrawCredential, uint256 _depositAmount) pure returns(bytes32)
func (_PermissionlessPool *PermissionlessPoolCallerSession) ComputeDepositDataRoot(_pubkey []byte, _signature []byte, _withdrawCredential []byte, _depositAmount *big.Int) ([32]byte, error) {
	return _PermissionlessPool.Contract.ComputeDepositDataRoot(&_PermissionlessPool.CallOpts, _pubkey, _signature, _withdrawCredential, _depositAmount)
}

// GetAllSocializingPoolOptOutOperators is a free data retrieval call binding the contract method 0x61e2f809.
//
// Solidity: function getAllSocializingPoolOptOutOperators(uint256 _pageNumber, uint256 _pageSize) view returns(address[])
func (_PermissionlessPool *PermissionlessPoolCaller) GetAllSocializingPoolOptOutOperators(opts *bind.CallOpts, _pageNumber *big.Int, _pageSize *big.Int) ([]common.Address, error) {
	var out []interface{}
	err := _PermissionlessPool.contract.Call(opts, &out, "getAllSocializingPoolOptOutOperators", _pageNumber, _pageSize)

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// GetAllSocializingPoolOptOutOperators is a free data retrieval call binding the contract method 0x61e2f809.
//
// Solidity: function getAllSocializingPoolOptOutOperators(uint256 _pageNumber, uint256 _pageSize) view returns(address[])
func (_PermissionlessPool *PermissionlessPoolSession) GetAllSocializingPoolOptOutOperators(_pageNumber *big.Int, _pageSize *big.Int) ([]common.Address, error) {
	return _PermissionlessPool.Contract.GetAllSocializingPoolOptOutOperators(&_PermissionlessPool.CallOpts, _pageNumber, _pageSize)
}

// GetAllSocializingPoolOptOutOperators is a free data retrieval call binding the contract method 0x61e2f809.
//
// Solidity: function getAllSocializingPoolOptOutOperators(uint256 _pageNumber, uint256 _pageSize) view returns(address[])
func (_PermissionlessPool *PermissionlessPoolCallerSession) GetAllSocializingPoolOptOutOperators(_pageNumber *big.Int, _pageSize *big.Int) ([]common.Address, error) {
	return _PermissionlessPool.Contract.GetAllSocializingPoolOptOutOperators(&_PermissionlessPool.CallOpts, _pageNumber, _pageSize)
}

// GetCollateralETH is a free data retrieval call binding the contract method 0xb01db078.
//
// Solidity: function getCollateralETH() view returns(uint256)
func (_PermissionlessPool *PermissionlessPoolCaller) GetCollateralETH(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _PermissionlessPool.contract.Call(opts, &out, "getCollateralETH")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetCollateralETH is a free data retrieval call binding the contract method 0xb01db078.
//
// Solidity: function getCollateralETH() view returns(uint256)
func (_PermissionlessPool *PermissionlessPoolSession) GetCollateralETH() (*big.Int, error) {
	return _PermissionlessPool.Contract.GetCollateralETH(&_PermissionlessPool.CallOpts)
}

// GetCollateralETH is a free data retrieval call binding the contract method 0xb01db078.
//
// Solidity: function getCollateralETH() view returns(uint256)
func (_PermissionlessPool *PermissionlessPoolCallerSession) GetCollateralETH() (*big.Int, error) {
	return _PermissionlessPool.Contract.GetCollateralETH(&_PermissionlessPool.CallOpts)
}

// GetNodeRegistry is a free data retrieval call binding the contract method 0xb6fb3fac.
//
// Solidity: function getNodeRegistry() view returns(address)
func (_PermissionlessPool *PermissionlessPoolCaller) GetNodeRegistry(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _PermissionlessPool.contract.Call(opts, &out, "getNodeRegistry")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetNodeRegistry is a free data retrieval call binding the contract method 0xb6fb3fac.
//
// Solidity: function getNodeRegistry() view returns(address)
func (_PermissionlessPool *PermissionlessPoolSession) GetNodeRegistry() (common.Address, error) {
	return _PermissionlessPool.Contract.GetNodeRegistry(&_PermissionlessPool.CallOpts)
}

// GetNodeRegistry is a free data retrieval call binding the contract method 0xb6fb3fac.
//
// Solidity: function getNodeRegistry() view returns(address)
func (_PermissionlessPool *PermissionlessPoolCallerSession) GetNodeRegistry() (common.Address, error) {
	return _PermissionlessPool.Contract.GetNodeRegistry(&_PermissionlessPool.CallOpts)
}

// GetOperatorTotalNonTerminalKeys is a free data retrieval call binding the contract method 0x8a25bcec.
//
// Solidity: function getOperatorTotalNonTerminalKeys(address _nodeOperator, uint256 _startIndex, uint256 _endIndex) view returns(uint256)
func (_PermissionlessPool *PermissionlessPoolCaller) GetOperatorTotalNonTerminalKeys(opts *bind.CallOpts, _nodeOperator common.Address, _startIndex *big.Int, _endIndex *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _PermissionlessPool.contract.Call(opts, &out, "getOperatorTotalNonTerminalKeys", _nodeOperator, _startIndex, _endIndex)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetOperatorTotalNonTerminalKeys is a free data retrieval call binding the contract method 0x8a25bcec.
//
// Solidity: function getOperatorTotalNonTerminalKeys(address _nodeOperator, uint256 _startIndex, uint256 _endIndex) view returns(uint256)
func (_PermissionlessPool *PermissionlessPoolSession) GetOperatorTotalNonTerminalKeys(_nodeOperator common.Address, _startIndex *big.Int, _endIndex *big.Int) (*big.Int, error) {
	return _PermissionlessPool.Contract.GetOperatorTotalNonTerminalKeys(&_PermissionlessPool.CallOpts, _nodeOperator, _startIndex, _endIndex)
}

// GetOperatorTotalNonTerminalKeys is a free data retrieval call binding the contract method 0x8a25bcec.
//
// Solidity: function getOperatorTotalNonTerminalKeys(address _nodeOperator, uint256 _startIndex, uint256 _endIndex) view returns(uint256)
func (_PermissionlessPool *PermissionlessPoolCallerSession) GetOperatorTotalNonTerminalKeys(_nodeOperator common.Address, _startIndex *big.Int, _endIndex *big.Int) (*big.Int, error) {
	return _PermissionlessPool.Contract.GetOperatorTotalNonTerminalKeys(&_PermissionlessPool.CallOpts, _nodeOperator, _startIndex, _endIndex)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_PermissionlessPool *PermissionlessPoolCaller) GetRoleAdmin(opts *bind.CallOpts, role [32]byte) ([32]byte, error) {
	var out []interface{}
	err := _PermissionlessPool.contract.Call(opts, &out, "getRoleAdmin", role)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_PermissionlessPool *PermissionlessPoolSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _PermissionlessPool.Contract.GetRoleAdmin(&_PermissionlessPool.CallOpts, role)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_PermissionlessPool *PermissionlessPoolCallerSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _PermissionlessPool.Contract.GetRoleAdmin(&_PermissionlessPool.CallOpts, role)
}

// GetSocializingPoolAddress is a free data retrieval call binding the contract method 0xf74b4cd1.
//
// Solidity: function getSocializingPoolAddress() view returns(address)
func (_PermissionlessPool *PermissionlessPoolCaller) GetSocializingPoolAddress(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _PermissionlessPool.contract.Call(opts, &out, "getSocializingPoolAddress")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetSocializingPoolAddress is a free data retrieval call binding the contract method 0xf74b4cd1.
//
// Solidity: function getSocializingPoolAddress() view returns(address)
func (_PermissionlessPool *PermissionlessPoolSession) GetSocializingPoolAddress() (common.Address, error) {
	return _PermissionlessPool.Contract.GetSocializingPoolAddress(&_PermissionlessPool.CallOpts)
}

// GetSocializingPoolAddress is a free data retrieval call binding the contract method 0xf74b4cd1.
//
// Solidity: function getSocializingPoolAddress() view returns(address)
func (_PermissionlessPool *PermissionlessPoolCallerSession) GetSocializingPoolAddress() (common.Address, error) {
	return _PermissionlessPool.Contract.GetSocializingPoolAddress(&_PermissionlessPool.CallOpts)
}

// GetTotalActiveValidatorCount is a free data retrieval call binding the contract method 0x77c359e1.
//
// Solidity: function getTotalActiveValidatorCount() view returns(uint256)
func (_PermissionlessPool *PermissionlessPoolCaller) GetTotalActiveValidatorCount(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _PermissionlessPool.contract.Call(opts, &out, "getTotalActiveValidatorCount")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetTotalActiveValidatorCount is a free data retrieval call binding the contract method 0x77c359e1.
//
// Solidity: function getTotalActiveValidatorCount() view returns(uint256)
func (_PermissionlessPool *PermissionlessPoolSession) GetTotalActiveValidatorCount() (*big.Int, error) {
	return _PermissionlessPool.Contract.GetTotalActiveValidatorCount(&_PermissionlessPool.CallOpts)
}

// GetTotalActiveValidatorCount is a free data retrieval call binding the contract method 0x77c359e1.
//
// Solidity: function getTotalActiveValidatorCount() view returns(uint256)
func (_PermissionlessPool *PermissionlessPoolCallerSession) GetTotalActiveValidatorCount() (*big.Int, error) {
	return _PermissionlessPool.Contract.GetTotalActiveValidatorCount(&_PermissionlessPool.CallOpts)
}

// GetTotalQueuedValidatorCount is a free data retrieval call binding the contract method 0x7bd977d9.
//
// Solidity: function getTotalQueuedValidatorCount() view returns(uint256)
func (_PermissionlessPool *PermissionlessPoolCaller) GetTotalQueuedValidatorCount(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _PermissionlessPool.contract.Call(opts, &out, "getTotalQueuedValidatorCount")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetTotalQueuedValidatorCount is a free data retrieval call binding the contract method 0x7bd977d9.
//
// Solidity: function getTotalQueuedValidatorCount() view returns(uint256)
func (_PermissionlessPool *PermissionlessPoolSession) GetTotalQueuedValidatorCount() (*big.Int, error) {
	return _PermissionlessPool.Contract.GetTotalQueuedValidatorCount(&_PermissionlessPool.CallOpts)
}

// GetTotalQueuedValidatorCount is a free data retrieval call binding the contract method 0x7bd977d9.
//
// Solidity: function getTotalQueuedValidatorCount() view returns(uint256)
func (_PermissionlessPool *PermissionlessPoolCallerSession) GetTotalQueuedValidatorCount() (*big.Int, error) {
	return _PermissionlessPool.Contract.GetTotalQueuedValidatorCount(&_PermissionlessPool.CallOpts)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_PermissionlessPool *PermissionlessPoolCaller) HasRole(opts *bind.CallOpts, role [32]byte, account common.Address) (bool, error) {
	var out []interface{}
	err := _PermissionlessPool.contract.Call(opts, &out, "hasRole", role, account)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_PermissionlessPool *PermissionlessPoolSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _PermissionlessPool.Contract.HasRole(&_PermissionlessPool.CallOpts, role, account)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_PermissionlessPool *PermissionlessPoolCallerSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _PermissionlessPool.Contract.HasRole(&_PermissionlessPool.CallOpts, role, account)
}

// IsExistingOperator is a free data retrieval call binding the contract method 0xf9c4dda4.
//
// Solidity: function isExistingOperator(address _operAddr) view returns(bool)
func (_PermissionlessPool *PermissionlessPoolCaller) IsExistingOperator(opts *bind.CallOpts, _operAddr common.Address) (bool, error) {
	var out []interface{}
	err := _PermissionlessPool.contract.Call(opts, &out, "isExistingOperator", _operAddr)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsExistingOperator is a free data retrieval call binding the contract method 0xf9c4dda4.
//
// Solidity: function isExistingOperator(address _operAddr) view returns(bool)
func (_PermissionlessPool *PermissionlessPoolSession) IsExistingOperator(_operAddr common.Address) (bool, error) {
	return _PermissionlessPool.Contract.IsExistingOperator(&_PermissionlessPool.CallOpts, _operAddr)
}

// IsExistingOperator is a free data retrieval call binding the contract method 0xf9c4dda4.
//
// Solidity: function isExistingOperator(address _operAddr) view returns(bool)
func (_PermissionlessPool *PermissionlessPoolCallerSession) IsExistingOperator(_operAddr common.Address) (bool, error) {
	return _PermissionlessPool.Contract.IsExistingOperator(&_PermissionlessPool.CallOpts, _operAddr)
}

// IsExistingPubkey is a free data retrieval call binding the contract method 0x36514d9f.
//
// Solidity: function isExistingPubkey(bytes _pubkey) view returns(bool)
func (_PermissionlessPool *PermissionlessPoolCaller) IsExistingPubkey(opts *bind.CallOpts, _pubkey []byte) (bool, error) {
	var out []interface{}
	err := _PermissionlessPool.contract.Call(opts, &out, "isExistingPubkey", _pubkey)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsExistingPubkey is a free data retrieval call binding the contract method 0x36514d9f.
//
// Solidity: function isExistingPubkey(bytes _pubkey) view returns(bool)
func (_PermissionlessPool *PermissionlessPoolSession) IsExistingPubkey(_pubkey []byte) (bool, error) {
	return _PermissionlessPool.Contract.IsExistingPubkey(&_PermissionlessPool.CallOpts, _pubkey)
}

// IsExistingPubkey is a free data retrieval call binding the contract method 0x36514d9f.
//
// Solidity: function isExistingPubkey(bytes _pubkey) view returns(bool)
func (_PermissionlessPool *PermissionlessPoolCallerSession) IsExistingPubkey(_pubkey []byte) (bool, error) {
	return _PermissionlessPool.Contract.IsExistingPubkey(&_PermissionlessPool.CallOpts, _pubkey)
}

// OperatorFee is a free data retrieval call binding the contract method 0x89afc0f1.
//
// Solidity: function operatorFee() view returns(uint256)
func (_PermissionlessPool *PermissionlessPoolCaller) OperatorFee(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _PermissionlessPool.contract.Call(opts, &out, "operatorFee")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// OperatorFee is a free data retrieval call binding the contract method 0x89afc0f1.
//
// Solidity: function operatorFee() view returns(uint256)
func (_PermissionlessPool *PermissionlessPoolSession) OperatorFee() (*big.Int, error) {
	return _PermissionlessPool.Contract.OperatorFee(&_PermissionlessPool.CallOpts)
}

// OperatorFee is a free data retrieval call binding the contract method 0x89afc0f1.
//
// Solidity: function operatorFee() view returns(uint256)
func (_PermissionlessPool *PermissionlessPoolCallerSession) OperatorFee() (*big.Int, error) {
	return _PermissionlessPool.Contract.OperatorFee(&_PermissionlessPool.CallOpts)
}

// ProtocolFee is a free data retrieval call binding the contract method 0xb0e21e8a.
//
// Solidity: function protocolFee() view returns(uint256)
func (_PermissionlessPool *PermissionlessPoolCaller) ProtocolFee(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _PermissionlessPool.contract.Call(opts, &out, "protocolFee")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ProtocolFee is a free data retrieval call binding the contract method 0xb0e21e8a.
//
// Solidity: function protocolFee() view returns(uint256)
func (_PermissionlessPool *PermissionlessPoolSession) ProtocolFee() (*big.Int, error) {
	return _PermissionlessPool.Contract.ProtocolFee(&_PermissionlessPool.CallOpts)
}

// ProtocolFee is a free data retrieval call binding the contract method 0xb0e21e8a.
//
// Solidity: function protocolFee() view returns(uint256)
func (_PermissionlessPool *PermissionlessPoolCallerSession) ProtocolFee() (*big.Int, error) {
	return _PermissionlessPool.Contract.ProtocolFee(&_PermissionlessPool.CallOpts)
}

// StaderConfig is a free data retrieval call binding the contract method 0x490ffa35.
//
// Solidity: function staderConfig() view returns(address)
func (_PermissionlessPool *PermissionlessPoolCaller) StaderConfig(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _PermissionlessPool.contract.Call(opts, &out, "staderConfig")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// StaderConfig is a free data retrieval call binding the contract method 0x490ffa35.
//
// Solidity: function staderConfig() view returns(address)
func (_PermissionlessPool *PermissionlessPoolSession) StaderConfig() (common.Address, error) {
	return _PermissionlessPool.Contract.StaderConfig(&_PermissionlessPool.CallOpts)
}

// StaderConfig is a free data retrieval call binding the contract method 0x490ffa35.
//
// Solidity: function staderConfig() view returns(address)
func (_PermissionlessPool *PermissionlessPoolCallerSession) StaderConfig() (common.Address, error) {
	return _PermissionlessPool.Contract.StaderConfig(&_PermissionlessPool.CallOpts)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_PermissionlessPool *PermissionlessPoolCaller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _PermissionlessPool.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_PermissionlessPool *PermissionlessPoolSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _PermissionlessPool.Contract.SupportsInterface(&_PermissionlessPool.CallOpts, interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_PermissionlessPool *PermissionlessPoolCallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _PermissionlessPool.Contract.SupportsInterface(&_PermissionlessPool.CallOpts, interfaceId)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_PermissionlessPool *PermissionlessPoolTransactor) GrantRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _PermissionlessPool.contract.Transact(opts, "grantRole", role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_PermissionlessPool *PermissionlessPoolSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _PermissionlessPool.Contract.GrantRole(&_PermissionlessPool.TransactOpts, role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_PermissionlessPool *PermissionlessPoolTransactorSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _PermissionlessPool.Contract.GrantRole(&_PermissionlessPool.TransactOpts, role, account)
}

// Initialize is a paid mutator transaction binding the contract method 0x485cc955.
//
// Solidity: function initialize(address _admin, address _staderConfig) returns()
func (_PermissionlessPool *PermissionlessPoolTransactor) Initialize(opts *bind.TransactOpts, _admin common.Address, _staderConfig common.Address) (*types.Transaction, error) {
	return _PermissionlessPool.contract.Transact(opts, "initialize", _admin, _staderConfig)
}

// Initialize is a paid mutator transaction binding the contract method 0x485cc955.
//
// Solidity: function initialize(address _admin, address _staderConfig) returns()
func (_PermissionlessPool *PermissionlessPoolSession) Initialize(_admin common.Address, _staderConfig common.Address) (*types.Transaction, error) {
	return _PermissionlessPool.Contract.Initialize(&_PermissionlessPool.TransactOpts, _admin, _staderConfig)
}

// Initialize is a paid mutator transaction binding the contract method 0x485cc955.
//
// Solidity: function initialize(address _admin, address _staderConfig) returns()
func (_PermissionlessPool *PermissionlessPoolTransactorSession) Initialize(_admin common.Address, _staderConfig common.Address) (*types.Transaction, error) {
	return _PermissionlessPool.Contract.Initialize(&_PermissionlessPool.TransactOpts, _admin, _staderConfig)
}

// PreDepositOnBeaconChain is a paid mutator transaction binding the contract method 0xeda0ae12.
//
// Solidity: function preDepositOnBeaconChain(bytes[] _pubkey, bytes[] _preDepositSignature, uint256 _operatorId, uint256 _operatorTotalKeys) payable returns()
func (_PermissionlessPool *PermissionlessPoolTransactor) PreDepositOnBeaconChain(opts *bind.TransactOpts, _pubkey [][]byte, _preDepositSignature [][]byte, _operatorId *big.Int, _operatorTotalKeys *big.Int) (*types.Transaction, error) {
	return _PermissionlessPool.contract.Transact(opts, "preDepositOnBeaconChain", _pubkey, _preDepositSignature, _operatorId, _operatorTotalKeys)
}

// PreDepositOnBeaconChain is a paid mutator transaction binding the contract method 0xeda0ae12.
//
// Solidity: function preDepositOnBeaconChain(bytes[] _pubkey, bytes[] _preDepositSignature, uint256 _operatorId, uint256 _operatorTotalKeys) payable returns()
func (_PermissionlessPool *PermissionlessPoolSession) PreDepositOnBeaconChain(_pubkey [][]byte, _preDepositSignature [][]byte, _operatorId *big.Int, _operatorTotalKeys *big.Int) (*types.Transaction, error) {
	return _PermissionlessPool.Contract.PreDepositOnBeaconChain(&_PermissionlessPool.TransactOpts, _pubkey, _preDepositSignature, _operatorId, _operatorTotalKeys)
}

// PreDepositOnBeaconChain is a paid mutator transaction binding the contract method 0xeda0ae12.
//
// Solidity: function preDepositOnBeaconChain(bytes[] _pubkey, bytes[] _preDepositSignature, uint256 _operatorId, uint256 _operatorTotalKeys) payable returns()
func (_PermissionlessPool *PermissionlessPoolTransactorSession) PreDepositOnBeaconChain(_pubkey [][]byte, _preDepositSignature [][]byte, _operatorId *big.Int, _operatorTotalKeys *big.Int) (*types.Transaction, error) {
	return _PermissionlessPool.Contract.PreDepositOnBeaconChain(&_PermissionlessPool.TransactOpts, _pubkey, _preDepositSignature, _operatorId, _operatorTotalKeys)
}

// ReceiveRemainingCollateralETH is a paid mutator transaction binding the contract method 0x1f033ef0.
//
// Solidity: function receiveRemainingCollateralETH() payable returns()
func (_PermissionlessPool *PermissionlessPoolTransactor) ReceiveRemainingCollateralETH(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PermissionlessPool.contract.Transact(opts, "receiveRemainingCollateralETH")
}

// ReceiveRemainingCollateralETH is a paid mutator transaction binding the contract method 0x1f033ef0.
//
// Solidity: function receiveRemainingCollateralETH() payable returns()
func (_PermissionlessPool *PermissionlessPoolSession) ReceiveRemainingCollateralETH() (*types.Transaction, error) {
	return _PermissionlessPool.Contract.ReceiveRemainingCollateralETH(&_PermissionlessPool.TransactOpts)
}

// ReceiveRemainingCollateralETH is a paid mutator transaction binding the contract method 0x1f033ef0.
//
// Solidity: function receiveRemainingCollateralETH() payable returns()
func (_PermissionlessPool *PermissionlessPoolTransactorSession) ReceiveRemainingCollateralETH() (*types.Transaction, error) {
	return _PermissionlessPool.Contract.ReceiveRemainingCollateralETH(&_PermissionlessPool.TransactOpts)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address account) returns()
func (_PermissionlessPool *PermissionlessPoolTransactor) RenounceRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _PermissionlessPool.contract.Transact(opts, "renounceRole", role, account)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address account) returns()
func (_PermissionlessPool *PermissionlessPoolSession) RenounceRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _PermissionlessPool.Contract.RenounceRole(&_PermissionlessPool.TransactOpts, role, account)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address account) returns()
func (_PermissionlessPool *PermissionlessPoolTransactorSession) RenounceRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _PermissionlessPool.Contract.RenounceRole(&_PermissionlessPool.TransactOpts, role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_PermissionlessPool *PermissionlessPoolTransactor) RevokeRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _PermissionlessPool.contract.Transact(opts, "revokeRole", role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_PermissionlessPool *PermissionlessPoolSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _PermissionlessPool.Contract.RevokeRole(&_PermissionlessPool.TransactOpts, role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_PermissionlessPool *PermissionlessPoolTransactorSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _PermissionlessPool.Contract.RevokeRole(&_PermissionlessPool.TransactOpts, role, account)
}

// SetCommissionFees is a paid mutator transaction binding the contract method 0x21066d18.
//
// Solidity: function setCommissionFees(uint256 _protocolFee, uint256 _operatorFee) returns()
func (_PermissionlessPool *PermissionlessPoolTransactor) SetCommissionFees(opts *bind.TransactOpts, _protocolFee *big.Int, _operatorFee *big.Int) (*types.Transaction, error) {
	return _PermissionlessPool.contract.Transact(opts, "setCommissionFees", _protocolFee, _operatorFee)
}

// SetCommissionFees is a paid mutator transaction binding the contract method 0x21066d18.
//
// Solidity: function setCommissionFees(uint256 _protocolFee, uint256 _operatorFee) returns()
func (_PermissionlessPool *PermissionlessPoolSession) SetCommissionFees(_protocolFee *big.Int, _operatorFee *big.Int) (*types.Transaction, error) {
	return _PermissionlessPool.Contract.SetCommissionFees(&_PermissionlessPool.TransactOpts, _protocolFee, _operatorFee)
}

// SetCommissionFees is a paid mutator transaction binding the contract method 0x21066d18.
//
// Solidity: function setCommissionFees(uint256 _protocolFee, uint256 _operatorFee) returns()
func (_PermissionlessPool *PermissionlessPoolTransactorSession) SetCommissionFees(_protocolFee *big.Int, _operatorFee *big.Int) (*types.Transaction, error) {
	return _PermissionlessPool.Contract.SetCommissionFees(&_PermissionlessPool.TransactOpts, _protocolFee, _operatorFee)
}

// StakeUserETHToBeaconChain is a paid mutator transaction binding the contract method 0x9b26728e.
//
// Solidity: function stakeUserETHToBeaconChain() payable returns()
func (_PermissionlessPool *PermissionlessPoolTransactor) StakeUserETHToBeaconChain(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PermissionlessPool.contract.Transact(opts, "stakeUserETHToBeaconChain")
}

// StakeUserETHToBeaconChain is a paid mutator transaction binding the contract method 0x9b26728e.
//
// Solidity: function stakeUserETHToBeaconChain() payable returns()
func (_PermissionlessPool *PermissionlessPoolSession) StakeUserETHToBeaconChain() (*types.Transaction, error) {
	return _PermissionlessPool.Contract.StakeUserETHToBeaconChain(&_PermissionlessPool.TransactOpts)
}

// StakeUserETHToBeaconChain is a paid mutator transaction binding the contract method 0x9b26728e.
//
// Solidity: function stakeUserETHToBeaconChain() payable returns()
func (_PermissionlessPool *PermissionlessPoolTransactorSession) StakeUserETHToBeaconChain() (*types.Transaction, error) {
	return _PermissionlessPool.Contract.StakeUserETHToBeaconChain(&_PermissionlessPool.TransactOpts)
}

// UpdateStaderConfig is a paid mutator transaction binding the contract method 0x9ee804cb.
//
// Solidity: function updateStaderConfig(address _staderConfig) returns()
func (_PermissionlessPool *PermissionlessPoolTransactor) UpdateStaderConfig(opts *bind.TransactOpts, _staderConfig common.Address) (*types.Transaction, error) {
	return _PermissionlessPool.contract.Transact(opts, "updateStaderConfig", _staderConfig)
}

// UpdateStaderConfig is a paid mutator transaction binding the contract method 0x9ee804cb.
//
// Solidity: function updateStaderConfig(address _staderConfig) returns()
func (_PermissionlessPool *PermissionlessPoolSession) UpdateStaderConfig(_staderConfig common.Address) (*types.Transaction, error) {
	return _PermissionlessPool.Contract.UpdateStaderConfig(&_PermissionlessPool.TransactOpts, _staderConfig)
}

// UpdateStaderConfig is a paid mutator transaction binding the contract method 0x9ee804cb.
//
// Solidity: function updateStaderConfig(address _staderConfig) returns()
func (_PermissionlessPool *PermissionlessPoolTransactorSession) UpdateStaderConfig(_staderConfig common.Address) (*types.Transaction, error) {
	return _PermissionlessPool.Contract.UpdateStaderConfig(&_PermissionlessPool.TransactOpts, _staderConfig)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() payable returns()
func (_PermissionlessPool *PermissionlessPoolTransactor) Fallback(opts *bind.TransactOpts, calldata []byte) (*types.Transaction, error) {
	return _PermissionlessPool.contract.RawTransact(opts, calldata)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() payable returns()
func (_PermissionlessPool *PermissionlessPoolSession) Fallback(calldata []byte) (*types.Transaction, error) {
	return _PermissionlessPool.Contract.Fallback(&_PermissionlessPool.TransactOpts, calldata)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() payable returns()
func (_PermissionlessPool *PermissionlessPoolTransactorSession) Fallback(calldata []byte) (*types.Transaction, error) {
	return _PermissionlessPool.Contract.Fallback(&_PermissionlessPool.TransactOpts, calldata)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_PermissionlessPool *PermissionlessPoolTransactor) Receive(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PermissionlessPool.contract.RawTransact(opts, nil) // calldata is disallowed for receive function
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_PermissionlessPool *PermissionlessPoolSession) Receive() (*types.Transaction, error) {
	return _PermissionlessPool.Contract.Receive(&_PermissionlessPool.TransactOpts)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_PermissionlessPool *PermissionlessPoolTransactorSession) Receive() (*types.Transaction, error) {
	return _PermissionlessPool.Contract.Receive(&_PermissionlessPool.TransactOpts)
}

// PermissionlessPoolInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the PermissionlessPool contract.
type PermissionlessPoolInitializedIterator struct {
	Event *PermissionlessPoolInitialized // Event containing the contract specifics and raw log

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
func (it *PermissionlessPoolInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PermissionlessPoolInitialized)
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
		it.Event = new(PermissionlessPoolInitialized)
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
func (it *PermissionlessPoolInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PermissionlessPoolInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PermissionlessPoolInitialized represents a Initialized event raised by the PermissionlessPool contract.
type PermissionlessPoolInitialized struct {
	Version uint8
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_PermissionlessPool *PermissionlessPoolFilterer) FilterInitialized(opts *bind.FilterOpts) (*PermissionlessPoolInitializedIterator, error) {

	logs, sub, err := _PermissionlessPool.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &PermissionlessPoolInitializedIterator{contract: _PermissionlessPool.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_PermissionlessPool *PermissionlessPoolFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *PermissionlessPoolInitialized) (event.Subscription, error) {

	logs, sub, err := _PermissionlessPool.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PermissionlessPoolInitialized)
				if err := _PermissionlessPool.contract.UnpackLog(event, "Initialized", log); err != nil {
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
func (_PermissionlessPool *PermissionlessPoolFilterer) ParseInitialized(log types.Log) (*PermissionlessPoolInitialized, error) {
	event := new(PermissionlessPoolInitialized)
	if err := _PermissionlessPool.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PermissionlessPoolReceivedCollateralETHIterator is returned from FilterReceivedCollateralETH and is used to iterate over the raw logs and unpacked data for ReceivedCollateralETH events raised by the PermissionlessPool contract.
type PermissionlessPoolReceivedCollateralETHIterator struct {
	Event *PermissionlessPoolReceivedCollateralETH // Event containing the contract specifics and raw log

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
func (it *PermissionlessPoolReceivedCollateralETHIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PermissionlessPoolReceivedCollateralETH)
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
		it.Event = new(PermissionlessPoolReceivedCollateralETH)
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
func (it *PermissionlessPoolReceivedCollateralETHIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PermissionlessPoolReceivedCollateralETHIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PermissionlessPoolReceivedCollateralETH represents a ReceivedCollateralETH event raised by the PermissionlessPool contract.
type PermissionlessPoolReceivedCollateralETH struct {
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterReceivedCollateralETH is a free log retrieval operation binding the contract event 0x3726a70cbf9252aacb06774b2f9f7108d837fc60afd7143f86eec77d7e3da94b.
//
// Solidity: event ReceivedCollateralETH(uint256 amount)
func (_PermissionlessPool *PermissionlessPoolFilterer) FilterReceivedCollateralETH(opts *bind.FilterOpts) (*PermissionlessPoolReceivedCollateralETHIterator, error) {

	logs, sub, err := _PermissionlessPool.contract.FilterLogs(opts, "ReceivedCollateralETH")
	if err != nil {
		return nil, err
	}
	return &PermissionlessPoolReceivedCollateralETHIterator{contract: _PermissionlessPool.contract, event: "ReceivedCollateralETH", logs: logs, sub: sub}, nil
}

// WatchReceivedCollateralETH is a free log subscription operation binding the contract event 0x3726a70cbf9252aacb06774b2f9f7108d837fc60afd7143f86eec77d7e3da94b.
//
// Solidity: event ReceivedCollateralETH(uint256 amount)
func (_PermissionlessPool *PermissionlessPoolFilterer) WatchReceivedCollateralETH(opts *bind.WatchOpts, sink chan<- *PermissionlessPoolReceivedCollateralETH) (event.Subscription, error) {

	logs, sub, err := _PermissionlessPool.contract.WatchLogs(opts, "ReceivedCollateralETH")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PermissionlessPoolReceivedCollateralETH)
				if err := _PermissionlessPool.contract.UnpackLog(event, "ReceivedCollateralETH", log); err != nil {
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

// ParseReceivedCollateralETH is a log parse operation binding the contract event 0x3726a70cbf9252aacb06774b2f9f7108d837fc60afd7143f86eec77d7e3da94b.
//
// Solidity: event ReceivedCollateralETH(uint256 amount)
func (_PermissionlessPool *PermissionlessPoolFilterer) ParseReceivedCollateralETH(log types.Log) (*PermissionlessPoolReceivedCollateralETH, error) {
	event := new(PermissionlessPoolReceivedCollateralETH)
	if err := _PermissionlessPool.contract.UnpackLog(event, "ReceivedCollateralETH", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PermissionlessPoolReceivedInsuranceFundIterator is returned from FilterReceivedInsuranceFund and is used to iterate over the raw logs and unpacked data for ReceivedInsuranceFund events raised by the PermissionlessPool contract.
type PermissionlessPoolReceivedInsuranceFundIterator struct {
	Event *PermissionlessPoolReceivedInsuranceFund // Event containing the contract specifics and raw log

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
func (it *PermissionlessPoolReceivedInsuranceFundIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PermissionlessPoolReceivedInsuranceFund)
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
		it.Event = new(PermissionlessPoolReceivedInsuranceFund)
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
func (it *PermissionlessPoolReceivedInsuranceFundIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PermissionlessPoolReceivedInsuranceFundIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PermissionlessPoolReceivedInsuranceFund represents a ReceivedInsuranceFund event raised by the PermissionlessPool contract.
type PermissionlessPoolReceivedInsuranceFund struct {
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterReceivedInsuranceFund is a free log retrieval operation binding the contract event 0xf43316e35306192a57665c134c0731f8e965fd6b625c620983cc68c2dd183726.
//
// Solidity: event ReceivedInsuranceFund(uint256 amount)
func (_PermissionlessPool *PermissionlessPoolFilterer) FilterReceivedInsuranceFund(opts *bind.FilterOpts) (*PermissionlessPoolReceivedInsuranceFundIterator, error) {

	logs, sub, err := _PermissionlessPool.contract.FilterLogs(opts, "ReceivedInsuranceFund")
	if err != nil {
		return nil, err
	}
	return &PermissionlessPoolReceivedInsuranceFundIterator{contract: _PermissionlessPool.contract, event: "ReceivedInsuranceFund", logs: logs, sub: sub}, nil
}

// WatchReceivedInsuranceFund is a free log subscription operation binding the contract event 0xf43316e35306192a57665c134c0731f8e965fd6b625c620983cc68c2dd183726.
//
// Solidity: event ReceivedInsuranceFund(uint256 amount)
func (_PermissionlessPool *PermissionlessPoolFilterer) WatchReceivedInsuranceFund(opts *bind.WatchOpts, sink chan<- *PermissionlessPoolReceivedInsuranceFund) (event.Subscription, error) {

	logs, sub, err := _PermissionlessPool.contract.WatchLogs(opts, "ReceivedInsuranceFund")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PermissionlessPoolReceivedInsuranceFund)
				if err := _PermissionlessPool.contract.UnpackLog(event, "ReceivedInsuranceFund", log); err != nil {
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

// ParseReceivedInsuranceFund is a log parse operation binding the contract event 0xf43316e35306192a57665c134c0731f8e965fd6b625c620983cc68c2dd183726.
//
// Solidity: event ReceivedInsuranceFund(uint256 amount)
func (_PermissionlessPool *PermissionlessPoolFilterer) ParseReceivedInsuranceFund(log types.Log) (*PermissionlessPoolReceivedInsuranceFund, error) {
	event := new(PermissionlessPoolReceivedInsuranceFund)
	if err := _PermissionlessPool.contract.UnpackLog(event, "ReceivedInsuranceFund", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PermissionlessPoolRoleAdminChangedIterator is returned from FilterRoleAdminChanged and is used to iterate over the raw logs and unpacked data for RoleAdminChanged events raised by the PermissionlessPool contract.
type PermissionlessPoolRoleAdminChangedIterator struct {
	Event *PermissionlessPoolRoleAdminChanged // Event containing the contract specifics and raw log

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
func (it *PermissionlessPoolRoleAdminChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PermissionlessPoolRoleAdminChanged)
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
		it.Event = new(PermissionlessPoolRoleAdminChanged)
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
func (it *PermissionlessPoolRoleAdminChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PermissionlessPoolRoleAdminChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PermissionlessPoolRoleAdminChanged represents a RoleAdminChanged event raised by the PermissionlessPool contract.
type PermissionlessPoolRoleAdminChanged struct {
	Role              [32]byte
	PreviousAdminRole [32]byte
	NewAdminRole      [32]byte
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterRoleAdminChanged is a free log retrieval operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_PermissionlessPool *PermissionlessPoolFilterer) FilterRoleAdminChanged(opts *bind.FilterOpts, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (*PermissionlessPoolRoleAdminChangedIterator, error) {

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

	logs, sub, err := _PermissionlessPool.contract.FilterLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return &PermissionlessPoolRoleAdminChangedIterator{contract: _PermissionlessPool.contract, event: "RoleAdminChanged", logs: logs, sub: sub}, nil
}

// WatchRoleAdminChanged is a free log subscription operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_PermissionlessPool *PermissionlessPoolFilterer) WatchRoleAdminChanged(opts *bind.WatchOpts, sink chan<- *PermissionlessPoolRoleAdminChanged, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (event.Subscription, error) {

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

	logs, sub, err := _PermissionlessPool.contract.WatchLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PermissionlessPoolRoleAdminChanged)
				if err := _PermissionlessPool.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
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
func (_PermissionlessPool *PermissionlessPoolFilterer) ParseRoleAdminChanged(log types.Log) (*PermissionlessPoolRoleAdminChanged, error) {
	event := new(PermissionlessPoolRoleAdminChanged)
	if err := _PermissionlessPool.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PermissionlessPoolRoleGrantedIterator is returned from FilterRoleGranted and is used to iterate over the raw logs and unpacked data for RoleGranted events raised by the PermissionlessPool contract.
type PermissionlessPoolRoleGrantedIterator struct {
	Event *PermissionlessPoolRoleGranted // Event containing the contract specifics and raw log

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
func (it *PermissionlessPoolRoleGrantedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PermissionlessPoolRoleGranted)
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
		it.Event = new(PermissionlessPoolRoleGranted)
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
func (it *PermissionlessPoolRoleGrantedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PermissionlessPoolRoleGrantedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PermissionlessPoolRoleGranted represents a RoleGranted event raised by the PermissionlessPool contract.
type PermissionlessPoolRoleGranted struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleGranted is a free log retrieval operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_PermissionlessPool *PermissionlessPoolFilterer) FilterRoleGranted(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*PermissionlessPoolRoleGrantedIterator, error) {

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

	logs, sub, err := _PermissionlessPool.contract.FilterLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &PermissionlessPoolRoleGrantedIterator{contract: _PermissionlessPool.contract, event: "RoleGranted", logs: logs, sub: sub}, nil
}

// WatchRoleGranted is a free log subscription operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_PermissionlessPool *PermissionlessPoolFilterer) WatchRoleGranted(opts *bind.WatchOpts, sink chan<- *PermissionlessPoolRoleGranted, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _PermissionlessPool.contract.WatchLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PermissionlessPoolRoleGranted)
				if err := _PermissionlessPool.contract.UnpackLog(event, "RoleGranted", log); err != nil {
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
func (_PermissionlessPool *PermissionlessPoolFilterer) ParseRoleGranted(log types.Log) (*PermissionlessPoolRoleGranted, error) {
	event := new(PermissionlessPoolRoleGranted)
	if err := _PermissionlessPool.contract.UnpackLog(event, "RoleGranted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PermissionlessPoolRoleRevokedIterator is returned from FilterRoleRevoked and is used to iterate over the raw logs and unpacked data for RoleRevoked events raised by the PermissionlessPool contract.
type PermissionlessPoolRoleRevokedIterator struct {
	Event *PermissionlessPoolRoleRevoked // Event containing the contract specifics and raw log

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
func (it *PermissionlessPoolRoleRevokedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PermissionlessPoolRoleRevoked)
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
		it.Event = new(PermissionlessPoolRoleRevoked)
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
func (it *PermissionlessPoolRoleRevokedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PermissionlessPoolRoleRevokedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PermissionlessPoolRoleRevoked represents a RoleRevoked event raised by the PermissionlessPool contract.
type PermissionlessPoolRoleRevoked struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleRevoked is a free log retrieval operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_PermissionlessPool *PermissionlessPoolFilterer) FilterRoleRevoked(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*PermissionlessPoolRoleRevokedIterator, error) {

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

	logs, sub, err := _PermissionlessPool.contract.FilterLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &PermissionlessPoolRoleRevokedIterator{contract: _PermissionlessPool.contract, event: "RoleRevoked", logs: logs, sub: sub}, nil
}

// WatchRoleRevoked is a free log subscription operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_PermissionlessPool *PermissionlessPoolFilterer) WatchRoleRevoked(opts *bind.WatchOpts, sink chan<- *PermissionlessPoolRoleRevoked, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _PermissionlessPool.contract.WatchLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PermissionlessPoolRoleRevoked)
				if err := _PermissionlessPool.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
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
func (_PermissionlessPool *PermissionlessPoolFilterer) ParseRoleRevoked(log types.Log) (*PermissionlessPoolRoleRevoked, error) {
	event := new(PermissionlessPoolRoleRevoked)
	if err := _PermissionlessPool.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PermissionlessPoolTransferredETHToSSPMForDefectiveKeysIterator is returned from FilterTransferredETHToSSPMForDefectiveKeys and is used to iterate over the raw logs and unpacked data for TransferredETHToSSPMForDefectiveKeys events raised by the PermissionlessPool contract.
type PermissionlessPoolTransferredETHToSSPMForDefectiveKeysIterator struct {
	Event *PermissionlessPoolTransferredETHToSSPMForDefectiveKeys // Event containing the contract specifics and raw log

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
func (it *PermissionlessPoolTransferredETHToSSPMForDefectiveKeysIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PermissionlessPoolTransferredETHToSSPMForDefectiveKeys)
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
		it.Event = new(PermissionlessPoolTransferredETHToSSPMForDefectiveKeys)
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
func (it *PermissionlessPoolTransferredETHToSSPMForDefectiveKeysIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PermissionlessPoolTransferredETHToSSPMForDefectiveKeysIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PermissionlessPoolTransferredETHToSSPMForDefectiveKeys represents a TransferredETHToSSPMForDefectiveKeys event raised by the PermissionlessPool contract.
type PermissionlessPoolTransferredETHToSSPMForDefectiveKeys struct {
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterTransferredETHToSSPMForDefectiveKeys is a free log retrieval operation binding the contract event 0x1149ac7366ffa1a2dee383a534b6742d013df82e7c15fd5720fd1ce82592a36f.
//
// Solidity: event TransferredETHToSSPMForDefectiveKeys(uint256 amount)
func (_PermissionlessPool *PermissionlessPoolFilterer) FilterTransferredETHToSSPMForDefectiveKeys(opts *bind.FilterOpts) (*PermissionlessPoolTransferredETHToSSPMForDefectiveKeysIterator, error) {

	logs, sub, err := _PermissionlessPool.contract.FilterLogs(opts, "TransferredETHToSSPMForDefectiveKeys")
	if err != nil {
		return nil, err
	}
	return &PermissionlessPoolTransferredETHToSSPMForDefectiveKeysIterator{contract: _PermissionlessPool.contract, event: "TransferredETHToSSPMForDefectiveKeys", logs: logs, sub: sub}, nil
}

// WatchTransferredETHToSSPMForDefectiveKeys is a free log subscription operation binding the contract event 0x1149ac7366ffa1a2dee383a534b6742d013df82e7c15fd5720fd1ce82592a36f.
//
// Solidity: event TransferredETHToSSPMForDefectiveKeys(uint256 amount)
func (_PermissionlessPool *PermissionlessPoolFilterer) WatchTransferredETHToSSPMForDefectiveKeys(opts *bind.WatchOpts, sink chan<- *PermissionlessPoolTransferredETHToSSPMForDefectiveKeys) (event.Subscription, error) {

	logs, sub, err := _PermissionlessPool.contract.WatchLogs(opts, "TransferredETHToSSPMForDefectiveKeys")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PermissionlessPoolTransferredETHToSSPMForDefectiveKeys)
				if err := _PermissionlessPool.contract.UnpackLog(event, "TransferredETHToSSPMForDefectiveKeys", log); err != nil {
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

// ParseTransferredETHToSSPMForDefectiveKeys is a log parse operation binding the contract event 0x1149ac7366ffa1a2dee383a534b6742d013df82e7c15fd5720fd1ce82592a36f.
//
// Solidity: event TransferredETHToSSPMForDefectiveKeys(uint256 amount)
func (_PermissionlessPool *PermissionlessPoolFilterer) ParseTransferredETHToSSPMForDefectiveKeys(log types.Log) (*PermissionlessPoolTransferredETHToSSPMForDefectiveKeys, error) {
	event := new(PermissionlessPoolTransferredETHToSSPMForDefectiveKeys)
	if err := _PermissionlessPool.contract.UnpackLog(event, "TransferredETHToSSPMForDefectiveKeys", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PermissionlessPoolUpdatedCommissionFeesIterator is returned from FilterUpdatedCommissionFees and is used to iterate over the raw logs and unpacked data for UpdatedCommissionFees events raised by the PermissionlessPool contract.
type PermissionlessPoolUpdatedCommissionFeesIterator struct {
	Event *PermissionlessPoolUpdatedCommissionFees // Event containing the contract specifics and raw log

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
func (it *PermissionlessPoolUpdatedCommissionFeesIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PermissionlessPoolUpdatedCommissionFees)
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
		it.Event = new(PermissionlessPoolUpdatedCommissionFees)
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
func (it *PermissionlessPoolUpdatedCommissionFeesIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PermissionlessPoolUpdatedCommissionFeesIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PermissionlessPoolUpdatedCommissionFees represents a UpdatedCommissionFees event raised by the PermissionlessPool contract.
type PermissionlessPoolUpdatedCommissionFees struct {
	ProtocolFee *big.Int
	OperatorFee *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterUpdatedCommissionFees is a free log retrieval operation binding the contract event 0xe8525a7862504bbe091b0440fe96979769664cae1625591ff0a9816512a5287b.
//
// Solidity: event UpdatedCommissionFees(uint256 protocolFee, uint256 operatorFee)
func (_PermissionlessPool *PermissionlessPoolFilterer) FilterUpdatedCommissionFees(opts *bind.FilterOpts) (*PermissionlessPoolUpdatedCommissionFeesIterator, error) {

	logs, sub, err := _PermissionlessPool.contract.FilterLogs(opts, "UpdatedCommissionFees")
	if err != nil {
		return nil, err
	}
	return &PermissionlessPoolUpdatedCommissionFeesIterator{contract: _PermissionlessPool.contract, event: "UpdatedCommissionFees", logs: logs, sub: sub}, nil
}

// WatchUpdatedCommissionFees is a free log subscription operation binding the contract event 0xe8525a7862504bbe091b0440fe96979769664cae1625591ff0a9816512a5287b.
//
// Solidity: event UpdatedCommissionFees(uint256 protocolFee, uint256 operatorFee)
func (_PermissionlessPool *PermissionlessPoolFilterer) WatchUpdatedCommissionFees(opts *bind.WatchOpts, sink chan<- *PermissionlessPoolUpdatedCommissionFees) (event.Subscription, error) {

	logs, sub, err := _PermissionlessPool.contract.WatchLogs(opts, "UpdatedCommissionFees")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PermissionlessPoolUpdatedCommissionFees)
				if err := _PermissionlessPool.contract.UnpackLog(event, "UpdatedCommissionFees", log); err != nil {
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

// ParseUpdatedCommissionFees is a log parse operation binding the contract event 0xe8525a7862504bbe091b0440fe96979769664cae1625591ff0a9816512a5287b.
//
// Solidity: event UpdatedCommissionFees(uint256 protocolFee, uint256 operatorFee)
func (_PermissionlessPool *PermissionlessPoolFilterer) ParseUpdatedCommissionFees(log types.Log) (*PermissionlessPoolUpdatedCommissionFees, error) {
	event := new(PermissionlessPoolUpdatedCommissionFees)
	if err := _PermissionlessPool.contract.UnpackLog(event, "UpdatedCommissionFees", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PermissionlessPoolUpdatedStaderConfigIterator is returned from FilterUpdatedStaderConfig and is used to iterate over the raw logs and unpacked data for UpdatedStaderConfig events raised by the PermissionlessPool contract.
type PermissionlessPoolUpdatedStaderConfigIterator struct {
	Event *PermissionlessPoolUpdatedStaderConfig // Event containing the contract specifics and raw log

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
func (it *PermissionlessPoolUpdatedStaderConfigIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PermissionlessPoolUpdatedStaderConfig)
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
		it.Event = new(PermissionlessPoolUpdatedStaderConfig)
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
func (it *PermissionlessPoolUpdatedStaderConfigIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PermissionlessPoolUpdatedStaderConfigIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PermissionlessPoolUpdatedStaderConfig represents a UpdatedStaderConfig event raised by the PermissionlessPool contract.
type PermissionlessPoolUpdatedStaderConfig struct {
	StaderConfig common.Address
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterUpdatedStaderConfig is a free log retrieval operation binding the contract event 0xdb2219043d7b197cb235f1af0cf6d782d77dee3de19e3f4fb6d39aae633b4485.
//
// Solidity: event UpdatedStaderConfig(address staderConfig)
func (_PermissionlessPool *PermissionlessPoolFilterer) FilterUpdatedStaderConfig(opts *bind.FilterOpts) (*PermissionlessPoolUpdatedStaderConfigIterator, error) {

	logs, sub, err := _PermissionlessPool.contract.FilterLogs(opts, "UpdatedStaderConfig")
	if err != nil {
		return nil, err
	}
	return &PermissionlessPoolUpdatedStaderConfigIterator{contract: _PermissionlessPool.contract, event: "UpdatedStaderConfig", logs: logs, sub: sub}, nil
}

// WatchUpdatedStaderConfig is a free log subscription operation binding the contract event 0xdb2219043d7b197cb235f1af0cf6d782d77dee3de19e3f4fb6d39aae633b4485.
//
// Solidity: event UpdatedStaderConfig(address staderConfig)
func (_PermissionlessPool *PermissionlessPoolFilterer) WatchUpdatedStaderConfig(opts *bind.WatchOpts, sink chan<- *PermissionlessPoolUpdatedStaderConfig) (event.Subscription, error) {

	logs, sub, err := _PermissionlessPool.contract.WatchLogs(opts, "UpdatedStaderConfig")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PermissionlessPoolUpdatedStaderConfig)
				if err := _PermissionlessPool.contract.UnpackLog(event, "UpdatedStaderConfig", log); err != nil {
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
func (_PermissionlessPool *PermissionlessPoolFilterer) ParseUpdatedStaderConfig(log types.Log) (*PermissionlessPoolUpdatedStaderConfig, error) {
	event := new(PermissionlessPoolUpdatedStaderConfig)
	if err := _PermissionlessPool.contract.UnpackLog(event, "UpdatedStaderConfig", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PermissionlessPoolValidatorDepositedOnBeaconChainIterator is returned from FilterValidatorDepositedOnBeaconChain and is used to iterate over the raw logs and unpacked data for ValidatorDepositedOnBeaconChain events raised by the PermissionlessPool contract.
type PermissionlessPoolValidatorDepositedOnBeaconChainIterator struct {
	Event *PermissionlessPoolValidatorDepositedOnBeaconChain // Event containing the contract specifics and raw log

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
func (it *PermissionlessPoolValidatorDepositedOnBeaconChainIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PermissionlessPoolValidatorDepositedOnBeaconChain)
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
		it.Event = new(PermissionlessPoolValidatorDepositedOnBeaconChain)
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
func (it *PermissionlessPoolValidatorDepositedOnBeaconChainIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PermissionlessPoolValidatorDepositedOnBeaconChainIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PermissionlessPoolValidatorDepositedOnBeaconChain represents a ValidatorDepositedOnBeaconChain event raised by the PermissionlessPool contract.
type PermissionlessPoolValidatorDepositedOnBeaconChain struct {
	ValidatorId *big.Int
	PubKey      []byte
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterValidatorDepositedOnBeaconChain is a free log retrieval operation binding the contract event 0xbef89de94658b7ef396ba7f9316542858c893c9011602906b1a2ad18d0a99c35.
//
// Solidity: event ValidatorDepositedOnBeaconChain(uint256 indexed validatorId, bytes pubKey)
func (_PermissionlessPool *PermissionlessPoolFilterer) FilterValidatorDepositedOnBeaconChain(opts *bind.FilterOpts, validatorId []*big.Int) (*PermissionlessPoolValidatorDepositedOnBeaconChainIterator, error) {

	var validatorIdRule []interface{}
	for _, validatorIdItem := range validatorId {
		validatorIdRule = append(validatorIdRule, validatorIdItem)
	}

	logs, sub, err := _PermissionlessPool.contract.FilterLogs(opts, "ValidatorDepositedOnBeaconChain", validatorIdRule)
	if err != nil {
		return nil, err
	}
	return &PermissionlessPoolValidatorDepositedOnBeaconChainIterator{contract: _PermissionlessPool.contract, event: "ValidatorDepositedOnBeaconChain", logs: logs, sub: sub}, nil
}

// WatchValidatorDepositedOnBeaconChain is a free log subscription operation binding the contract event 0xbef89de94658b7ef396ba7f9316542858c893c9011602906b1a2ad18d0a99c35.
//
// Solidity: event ValidatorDepositedOnBeaconChain(uint256 indexed validatorId, bytes pubKey)
func (_PermissionlessPool *PermissionlessPoolFilterer) WatchValidatorDepositedOnBeaconChain(opts *bind.WatchOpts, sink chan<- *PermissionlessPoolValidatorDepositedOnBeaconChain, validatorId []*big.Int) (event.Subscription, error) {

	var validatorIdRule []interface{}
	for _, validatorIdItem := range validatorId {
		validatorIdRule = append(validatorIdRule, validatorIdItem)
	}

	logs, sub, err := _PermissionlessPool.contract.WatchLogs(opts, "ValidatorDepositedOnBeaconChain", validatorIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PermissionlessPoolValidatorDepositedOnBeaconChain)
				if err := _PermissionlessPool.contract.UnpackLog(event, "ValidatorDepositedOnBeaconChain", log); err != nil {
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

// ParseValidatorDepositedOnBeaconChain is a log parse operation binding the contract event 0xbef89de94658b7ef396ba7f9316542858c893c9011602906b1a2ad18d0a99c35.
//
// Solidity: event ValidatorDepositedOnBeaconChain(uint256 indexed validatorId, bytes pubKey)
func (_PermissionlessPool *PermissionlessPoolFilterer) ParseValidatorDepositedOnBeaconChain(log types.Log) (*PermissionlessPoolValidatorDepositedOnBeaconChain, error) {
	event := new(PermissionlessPoolValidatorDepositedOnBeaconChain)
	if err := _PermissionlessPool.contract.UnpackLog(event, "ValidatorDepositedOnBeaconChain", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PermissionlessPoolValidatorPreDepositedOnBeaconChainIterator is returned from FilterValidatorPreDepositedOnBeaconChain and is used to iterate over the raw logs and unpacked data for ValidatorPreDepositedOnBeaconChain events raised by the PermissionlessPool contract.
type PermissionlessPoolValidatorPreDepositedOnBeaconChainIterator struct {
	Event *PermissionlessPoolValidatorPreDepositedOnBeaconChain // Event containing the contract specifics and raw log

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
func (it *PermissionlessPoolValidatorPreDepositedOnBeaconChainIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PermissionlessPoolValidatorPreDepositedOnBeaconChain)
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
		it.Event = new(PermissionlessPoolValidatorPreDepositedOnBeaconChain)
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
func (it *PermissionlessPoolValidatorPreDepositedOnBeaconChainIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PermissionlessPoolValidatorPreDepositedOnBeaconChainIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PermissionlessPoolValidatorPreDepositedOnBeaconChain represents a ValidatorPreDepositedOnBeaconChain event raised by the PermissionlessPool contract.
type PermissionlessPoolValidatorPreDepositedOnBeaconChain struct {
	PubKey common.Hash
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterValidatorPreDepositedOnBeaconChain is a free log retrieval operation binding the contract event 0xa35366ad083efbaee7949ff15c68508b95e2c0441248e80d73da97ac82bc1f10.
//
// Solidity: event ValidatorPreDepositedOnBeaconChain(bytes indexed pubKey)
func (_PermissionlessPool *PermissionlessPoolFilterer) FilterValidatorPreDepositedOnBeaconChain(opts *bind.FilterOpts, pubKey [][]byte) (*PermissionlessPoolValidatorPreDepositedOnBeaconChainIterator, error) {

	var pubKeyRule []interface{}
	for _, pubKeyItem := range pubKey {
		pubKeyRule = append(pubKeyRule, pubKeyItem)
	}

	logs, sub, err := _PermissionlessPool.contract.FilterLogs(opts, "ValidatorPreDepositedOnBeaconChain", pubKeyRule)
	if err != nil {
		return nil, err
	}
	return &PermissionlessPoolValidatorPreDepositedOnBeaconChainIterator{contract: _PermissionlessPool.contract, event: "ValidatorPreDepositedOnBeaconChain", logs: logs, sub: sub}, nil
}

// WatchValidatorPreDepositedOnBeaconChain is a free log subscription operation binding the contract event 0xa35366ad083efbaee7949ff15c68508b95e2c0441248e80d73da97ac82bc1f10.
//
// Solidity: event ValidatorPreDepositedOnBeaconChain(bytes indexed pubKey)
func (_PermissionlessPool *PermissionlessPoolFilterer) WatchValidatorPreDepositedOnBeaconChain(opts *bind.WatchOpts, sink chan<- *PermissionlessPoolValidatorPreDepositedOnBeaconChain, pubKey [][]byte) (event.Subscription, error) {

	var pubKeyRule []interface{}
	for _, pubKeyItem := range pubKey {
		pubKeyRule = append(pubKeyRule, pubKeyItem)
	}

	logs, sub, err := _PermissionlessPool.contract.WatchLogs(opts, "ValidatorPreDepositedOnBeaconChain", pubKeyRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PermissionlessPoolValidatorPreDepositedOnBeaconChain)
				if err := _PermissionlessPool.contract.UnpackLog(event, "ValidatorPreDepositedOnBeaconChain", log); err != nil {
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

// ParseValidatorPreDepositedOnBeaconChain is a log parse operation binding the contract event 0xa35366ad083efbaee7949ff15c68508b95e2c0441248e80d73da97ac82bc1f10.
//
// Solidity: event ValidatorPreDepositedOnBeaconChain(bytes indexed pubKey)
func (_PermissionlessPool *PermissionlessPoolFilterer) ParseValidatorPreDepositedOnBeaconChain(log types.Log) (*PermissionlessPoolValidatorPreDepositedOnBeaconChain, error) {
	event := new(PermissionlessPoolValidatorPreDepositedOnBeaconChain)
	if err := _PermissionlessPool.contract.UnpackLog(event, "ValidatorPreDepositedOnBeaconChain", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
