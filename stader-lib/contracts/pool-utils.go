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

// PoolUtilsMetaData contains all meta data concerning the PoolUtils contract.
var PoolUtilsMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"CallerNotManager\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"CallerNotOperator\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"EmptyNameString\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ExistingOrMismatchingPoolId\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidLengthOfPubkey\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidLengthOfSignature\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NameCrossedMaxLength\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OperatorIsNotOnboarded\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"PoolIdNotPresent\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"PubkeyAlreadyExist\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"PubkeyDoesNotExit\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ZeroAddress\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint8\",\"name\":\"poolId\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"poolAddress\",\"type\":\"address\"}],\"name\":\"DeactivatedPool\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"pubkey\",\"type\":\"bytes\"}],\"name\":\"ExitValidator\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"version\",\"type\":\"uint8\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint8\",\"name\":\"poolId\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"poolAddress\",\"type\":\"address\"}],\"name\":\"PoolAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint8\",\"name\":\"poolId\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"poolAddress\",\"type\":\"address\"}],\"name\":\"PoolAddressUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"previousAdminRole\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"newAdminRole\",\"type\":\"bytes32\"}],\"name\":\"RoleAdminChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleGranted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleRevoked\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"staderConfig\",\"type\":\"address\"}],\"name\":\"UpdatedStaderConfig\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"DEFAULT_ADMIN_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"_poolId\",\"type\":\"uint8\"},{\"internalType\":\"address\",\"name\":\"_poolAddress\",\"type\":\"address\"}],\"name\":\"addNewPool\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"_poolId\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"_totalRewards\",\"type\":\"uint256\"}],\"name\":\"calculateRewardShare\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"userShare\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"operatorShare\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"protocolShare\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"_poolId\",\"type\":\"uint8\"}],\"name\":\"getActiveValidatorCountByPool\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"_poolId\",\"type\":\"uint8\"}],\"name\":\"getCollateralETH\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"_poolId\",\"type\":\"uint8\"}],\"name\":\"getNodeRegistry\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"_poolId\",\"type\":\"uint8\"}],\"name\":\"getOperatorFee\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_operAddr\",\"type\":\"address\"}],\"name\":\"getOperatorPoolId\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"_poolId\",\"type\":\"uint8\"},{\"internalType\":\"address\",\"name\":\"_nodeOperator\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_startIndex\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_endIndex\",\"type\":\"uint256\"}],\"name\":\"getOperatorTotalNonTerminalKeys\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getPoolIdArray\",\"outputs\":[{\"internalType\":\"uint8[]\",\"name\":\"\",\"type\":\"uint8[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"_poolId\",\"type\":\"uint8\"}],\"name\":\"getProtocolFee\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"_poolId\",\"type\":\"uint8\"}],\"name\":\"getQueuedValidatorCountByPool\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"}],\"name\":\"getRoleAdmin\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"_poolId\",\"type\":\"uint8\"}],\"name\":\"getSocializingPoolAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getTotalActiveValidatorCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"_pubkey\",\"type\":\"bytes\"}],\"name\":\"getValidatorPoolId\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"grantRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"hasRole\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_admin\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_staderConfig\",\"type\":\"address\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_operAddr\",\"type\":\"address\"}],\"name\":\"isExistingOperator\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"_poolId\",\"type\":\"uint8\"}],\"name\":\"isExistingPoolId\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"_pubkey\",\"type\":\"bytes\"}],\"name\":\"isExistingPubkey\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"_pubkey\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"_preDepositSignature\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"_depositSignature\",\"type\":\"bytes\"}],\"name\":\"onlyValidKeys\",\"outputs\":[],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_name\",\"type\":\"string\"}],\"name\":\"onlyValidName\",\"outputs\":[],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"name\":\"poolAddressById\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"poolIdArray\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes[]\",\"name\":\"_pubkeys\",\"type\":\"bytes[]\"}],\"name\":\"processValidatorExitList\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"renounceRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"revokeRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"staderConfig\",\"outputs\":[{\"internalType\":\"contractIStaderConfig\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"_poolId\",\"type\":\"uint8\"},{\"internalType\":\"address\",\"name\":\"_newPoolAddress\",\"type\":\"address\"}],\"name\":\"updatePoolAddress\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_staderConfig\",\"type\":\"address\"}],\"name\":\"updateStaderConfig\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// PoolUtilsABI is the input ABI used to generate the binding from.
// Deprecated: Use PoolUtilsMetaData.ABI instead.
var PoolUtilsABI = PoolUtilsMetaData.ABI

// PoolUtils is an auto generated Go binding around an Ethereum contract.
type PoolUtils struct {
	PoolUtilsCaller     // Read-only binding to the contract
	PoolUtilsTransactor // Write-only binding to the contract
	PoolUtilsFilterer   // Log filterer for contract events
}

// PoolUtilsCaller is an auto generated read-only Go binding around an Ethereum contract.
type PoolUtilsCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PoolUtilsTransactor is an auto generated write-only Go binding around an Ethereum contract.
type PoolUtilsTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PoolUtilsFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type PoolUtilsFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PoolUtilsSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type PoolUtilsSession struct {
	Contract     *PoolUtils        // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// PoolUtilsCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type PoolUtilsCallerSession struct {
	Contract *PoolUtilsCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts    // Call options to use throughout this session
}

// PoolUtilsTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type PoolUtilsTransactorSession struct {
	Contract     *PoolUtilsTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// PoolUtilsRaw is an auto generated low-level Go binding around an Ethereum contract.
type PoolUtilsRaw struct {
	Contract *PoolUtils // Generic contract binding to access the raw methods on
}

// PoolUtilsCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type PoolUtilsCallerRaw struct {
	Contract *PoolUtilsCaller // Generic read-only contract binding to access the raw methods on
}

// PoolUtilsTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type PoolUtilsTransactorRaw struct {
	Contract *PoolUtilsTransactor // Generic write-only contract binding to access the raw methods on
}

// NewPoolUtils creates a new instance of PoolUtils, bound to a specific deployed contract.
func NewPoolUtils(address common.Address, backend bind.ContractBackend) (*PoolUtils, error) {
	contract, err := bindPoolUtils(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &PoolUtils{PoolUtilsCaller: PoolUtilsCaller{contract: contract}, PoolUtilsTransactor: PoolUtilsTransactor{contract: contract}, PoolUtilsFilterer: PoolUtilsFilterer{contract: contract}}, nil
}

// NewPoolUtilsCaller creates a new read-only instance of PoolUtils, bound to a specific deployed contract.
func NewPoolUtilsCaller(address common.Address, caller bind.ContractCaller) (*PoolUtilsCaller, error) {
	contract, err := bindPoolUtils(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &PoolUtilsCaller{contract: contract}, nil
}

// NewPoolUtilsTransactor creates a new write-only instance of PoolUtils, bound to a specific deployed contract.
func NewPoolUtilsTransactor(address common.Address, transactor bind.ContractTransactor) (*PoolUtilsTransactor, error) {
	contract, err := bindPoolUtils(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &PoolUtilsTransactor{contract: contract}, nil
}

// NewPoolUtilsFilterer creates a new log filterer instance of PoolUtils, bound to a specific deployed contract.
func NewPoolUtilsFilterer(address common.Address, filterer bind.ContractFilterer) (*PoolUtilsFilterer, error) {
	contract, err := bindPoolUtils(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &PoolUtilsFilterer{contract: contract}, nil
}

// bindPoolUtils binds a generic wrapper to an already deployed contract.
func bindPoolUtils(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := PoolUtilsMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_PoolUtils *PoolUtilsRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _PoolUtils.Contract.PoolUtilsCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_PoolUtils *PoolUtilsRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PoolUtils.Contract.PoolUtilsTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_PoolUtils *PoolUtilsRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _PoolUtils.Contract.PoolUtilsTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_PoolUtils *PoolUtilsCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _PoolUtils.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_PoolUtils *PoolUtilsTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PoolUtils.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_PoolUtils *PoolUtilsTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _PoolUtils.Contract.contract.Transact(opts, method, params...)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_PoolUtils *PoolUtilsCaller) DEFAULTADMINROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _PoolUtils.contract.Call(opts, &out, "DEFAULT_ADMIN_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_PoolUtils *PoolUtilsSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _PoolUtils.Contract.DEFAULTADMINROLE(&_PoolUtils.CallOpts)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_PoolUtils *PoolUtilsCallerSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _PoolUtils.Contract.DEFAULTADMINROLE(&_PoolUtils.CallOpts)
}

// CalculateRewardShare is a free data retrieval call binding the contract method 0xafc2afba.
//
// Solidity: function calculateRewardShare(uint8 _poolId, uint256 _totalRewards) view returns(uint256 userShare, uint256 operatorShare, uint256 protocolShare)
func (_PoolUtils *PoolUtilsCaller) CalculateRewardShare(opts *bind.CallOpts, _poolId uint8, _totalRewards *big.Int) (struct {
	UserShare     *big.Int
	OperatorShare *big.Int
	ProtocolShare *big.Int
}, error) {
	var out []interface{}
	err := _PoolUtils.contract.Call(opts, &out, "calculateRewardShare", _poolId, _totalRewards)

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

// CalculateRewardShare is a free data retrieval call binding the contract method 0xafc2afba.
//
// Solidity: function calculateRewardShare(uint8 _poolId, uint256 _totalRewards) view returns(uint256 userShare, uint256 operatorShare, uint256 protocolShare)
func (_PoolUtils *PoolUtilsSession) CalculateRewardShare(_poolId uint8, _totalRewards *big.Int) (struct {
	UserShare     *big.Int
	OperatorShare *big.Int
	ProtocolShare *big.Int
}, error) {
	return _PoolUtils.Contract.CalculateRewardShare(&_PoolUtils.CallOpts, _poolId, _totalRewards)
}

// CalculateRewardShare is a free data retrieval call binding the contract method 0xafc2afba.
//
// Solidity: function calculateRewardShare(uint8 _poolId, uint256 _totalRewards) view returns(uint256 userShare, uint256 operatorShare, uint256 protocolShare)
func (_PoolUtils *PoolUtilsCallerSession) CalculateRewardShare(_poolId uint8, _totalRewards *big.Int) (struct {
	UserShare     *big.Int
	OperatorShare *big.Int
	ProtocolShare *big.Int
}, error) {
	return _PoolUtils.Contract.CalculateRewardShare(&_PoolUtils.CallOpts, _poolId, _totalRewards)
}

// GetActiveValidatorCountByPool is a free data retrieval call binding the contract method 0x1ec2db3c.
//
// Solidity: function getActiveValidatorCountByPool(uint8 _poolId) view returns(uint256)
func (_PoolUtils *PoolUtilsCaller) GetActiveValidatorCountByPool(opts *bind.CallOpts, _poolId uint8) (*big.Int, error) {
	var out []interface{}
	err := _PoolUtils.contract.Call(opts, &out, "getActiveValidatorCountByPool", _poolId)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetActiveValidatorCountByPool is a free data retrieval call binding the contract method 0x1ec2db3c.
//
// Solidity: function getActiveValidatorCountByPool(uint8 _poolId) view returns(uint256)
func (_PoolUtils *PoolUtilsSession) GetActiveValidatorCountByPool(_poolId uint8) (*big.Int, error) {
	return _PoolUtils.Contract.GetActiveValidatorCountByPool(&_PoolUtils.CallOpts, _poolId)
}

// GetActiveValidatorCountByPool is a free data retrieval call binding the contract method 0x1ec2db3c.
//
// Solidity: function getActiveValidatorCountByPool(uint8 _poolId) view returns(uint256)
func (_PoolUtils *PoolUtilsCallerSession) GetActiveValidatorCountByPool(_poolId uint8) (*big.Int, error) {
	return _PoolUtils.Contract.GetActiveValidatorCountByPool(&_PoolUtils.CallOpts, _poolId)
}

// GetCollateralETH is a free data retrieval call binding the contract method 0xef7bba86.
//
// Solidity: function getCollateralETH(uint8 _poolId) view returns(uint256)
func (_PoolUtils *PoolUtilsCaller) GetCollateralETH(opts *bind.CallOpts, _poolId uint8) (*big.Int, error) {
	var out []interface{}
	err := _PoolUtils.contract.Call(opts, &out, "getCollateralETH", _poolId)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetCollateralETH is a free data retrieval call binding the contract method 0xef7bba86.
//
// Solidity: function getCollateralETH(uint8 _poolId) view returns(uint256)
func (_PoolUtils *PoolUtilsSession) GetCollateralETH(_poolId uint8) (*big.Int, error) {
	return _PoolUtils.Contract.GetCollateralETH(&_PoolUtils.CallOpts, _poolId)
}

// GetCollateralETH is a free data retrieval call binding the contract method 0xef7bba86.
//
// Solidity: function getCollateralETH(uint8 _poolId) view returns(uint256)
func (_PoolUtils *PoolUtilsCallerSession) GetCollateralETH(_poolId uint8) (*big.Int, error) {
	return _PoolUtils.Contract.GetCollateralETH(&_PoolUtils.CallOpts, _poolId)
}

// GetNodeRegistry is a free data retrieval call binding the contract method 0x99d055c8.
//
// Solidity: function getNodeRegistry(uint8 _poolId) view returns(address)
func (_PoolUtils *PoolUtilsCaller) GetNodeRegistry(opts *bind.CallOpts, _poolId uint8) (common.Address, error) {
	var out []interface{}
	err := _PoolUtils.contract.Call(opts, &out, "getNodeRegistry", _poolId)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetNodeRegistry is a free data retrieval call binding the contract method 0x99d055c8.
//
// Solidity: function getNodeRegistry(uint8 _poolId) view returns(address)
func (_PoolUtils *PoolUtilsSession) GetNodeRegistry(_poolId uint8) (common.Address, error) {
	return _PoolUtils.Contract.GetNodeRegistry(&_PoolUtils.CallOpts, _poolId)
}

// GetNodeRegistry is a free data retrieval call binding the contract method 0x99d055c8.
//
// Solidity: function getNodeRegistry(uint8 _poolId) view returns(address)
func (_PoolUtils *PoolUtilsCallerSession) GetNodeRegistry(_poolId uint8) (common.Address, error) {
	return _PoolUtils.Contract.GetNodeRegistry(&_PoolUtils.CallOpts, _poolId)
}

// GetOperatorFee is a free data retrieval call binding the contract method 0xb0ef1e18.
//
// Solidity: function getOperatorFee(uint8 _poolId) view returns(uint256)
func (_PoolUtils *PoolUtilsCaller) GetOperatorFee(opts *bind.CallOpts, _poolId uint8) (*big.Int, error) {
	var out []interface{}
	err := _PoolUtils.contract.Call(opts, &out, "getOperatorFee", _poolId)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetOperatorFee is a free data retrieval call binding the contract method 0xb0ef1e18.
//
// Solidity: function getOperatorFee(uint8 _poolId) view returns(uint256)
func (_PoolUtils *PoolUtilsSession) GetOperatorFee(_poolId uint8) (*big.Int, error) {
	return _PoolUtils.Contract.GetOperatorFee(&_PoolUtils.CallOpts, _poolId)
}

// GetOperatorFee is a free data retrieval call binding the contract method 0xb0ef1e18.
//
// Solidity: function getOperatorFee(uint8 _poolId) view returns(uint256)
func (_PoolUtils *PoolUtilsCallerSession) GetOperatorFee(_poolId uint8) (*big.Int, error) {
	return _PoolUtils.Contract.GetOperatorFee(&_PoolUtils.CallOpts, _poolId)
}

// GetOperatorPoolId is a free data retrieval call binding the contract method 0x8e43c53a.
//
// Solidity: function getOperatorPoolId(address _operAddr) view returns(uint8)
func (_PoolUtils *PoolUtilsCaller) GetOperatorPoolId(opts *bind.CallOpts, _operAddr common.Address) (uint8, error) {
	var out []interface{}
	err := _PoolUtils.contract.Call(opts, &out, "getOperatorPoolId", _operAddr)

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// GetOperatorPoolId is a free data retrieval call binding the contract method 0x8e43c53a.
//
// Solidity: function getOperatorPoolId(address _operAddr) view returns(uint8)
func (_PoolUtils *PoolUtilsSession) GetOperatorPoolId(_operAddr common.Address) (uint8, error) {
	return _PoolUtils.Contract.GetOperatorPoolId(&_PoolUtils.CallOpts, _operAddr)
}

// GetOperatorPoolId is a free data retrieval call binding the contract method 0x8e43c53a.
//
// Solidity: function getOperatorPoolId(address _operAddr) view returns(uint8)
func (_PoolUtils *PoolUtilsCallerSession) GetOperatorPoolId(_operAddr common.Address) (uint8, error) {
	return _PoolUtils.Contract.GetOperatorPoolId(&_PoolUtils.CallOpts, _operAddr)
}

// GetOperatorTotalNonTerminalKeys is a free data retrieval call binding the contract method 0x5d713ec3.
//
// Solidity: function getOperatorTotalNonTerminalKeys(uint8 _poolId, address _nodeOperator, uint256 _startIndex, uint256 _endIndex) view returns(uint256)
func (_PoolUtils *PoolUtilsCaller) GetOperatorTotalNonTerminalKeys(opts *bind.CallOpts, _poolId uint8, _nodeOperator common.Address, _startIndex *big.Int, _endIndex *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _PoolUtils.contract.Call(opts, &out, "getOperatorTotalNonTerminalKeys", _poolId, _nodeOperator, _startIndex, _endIndex)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetOperatorTotalNonTerminalKeys is a free data retrieval call binding the contract method 0x5d713ec3.
//
// Solidity: function getOperatorTotalNonTerminalKeys(uint8 _poolId, address _nodeOperator, uint256 _startIndex, uint256 _endIndex) view returns(uint256)
func (_PoolUtils *PoolUtilsSession) GetOperatorTotalNonTerminalKeys(_poolId uint8, _nodeOperator common.Address, _startIndex *big.Int, _endIndex *big.Int) (*big.Int, error) {
	return _PoolUtils.Contract.GetOperatorTotalNonTerminalKeys(&_PoolUtils.CallOpts, _poolId, _nodeOperator, _startIndex, _endIndex)
}

// GetOperatorTotalNonTerminalKeys is a free data retrieval call binding the contract method 0x5d713ec3.
//
// Solidity: function getOperatorTotalNonTerminalKeys(uint8 _poolId, address _nodeOperator, uint256 _startIndex, uint256 _endIndex) view returns(uint256)
func (_PoolUtils *PoolUtilsCallerSession) GetOperatorTotalNonTerminalKeys(_poolId uint8, _nodeOperator common.Address, _startIndex *big.Int, _endIndex *big.Int) (*big.Int, error) {
	return _PoolUtils.Contract.GetOperatorTotalNonTerminalKeys(&_PoolUtils.CallOpts, _poolId, _nodeOperator, _startIndex, _endIndex)
}

// GetPoolIdArray is a free data retrieval call binding the contract method 0xa92e1faf.
//
// Solidity: function getPoolIdArray() view returns(uint8[])
func (_PoolUtils *PoolUtilsCaller) GetPoolIdArray(opts *bind.CallOpts) ([]uint8, error) {
	var out []interface{}
	err := _PoolUtils.contract.Call(opts, &out, "getPoolIdArray")

	if err != nil {
		return *new([]uint8), err
	}

	out0 := *abi.ConvertType(out[0], new([]uint8)).(*[]uint8)

	return out0, err

}

// GetPoolIdArray is a free data retrieval call binding the contract method 0xa92e1faf.
//
// Solidity: function getPoolIdArray() view returns(uint8[])
func (_PoolUtils *PoolUtilsSession) GetPoolIdArray() ([]uint8, error) {
	return _PoolUtils.Contract.GetPoolIdArray(&_PoolUtils.CallOpts)
}

// GetPoolIdArray is a free data retrieval call binding the contract method 0xa92e1faf.
//
// Solidity: function getPoolIdArray() view returns(uint8[])
func (_PoolUtils *PoolUtilsCallerSession) GetPoolIdArray() ([]uint8, error) {
	return _PoolUtils.Contract.GetPoolIdArray(&_PoolUtils.CallOpts)
}

// GetProtocolFee is a free data retrieval call binding the contract method 0x261d41f5.
//
// Solidity: function getProtocolFee(uint8 _poolId) view returns(uint256)
func (_PoolUtils *PoolUtilsCaller) GetProtocolFee(opts *bind.CallOpts, _poolId uint8) (*big.Int, error) {
	var out []interface{}
	err := _PoolUtils.contract.Call(opts, &out, "getProtocolFee", _poolId)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetProtocolFee is a free data retrieval call binding the contract method 0x261d41f5.
//
// Solidity: function getProtocolFee(uint8 _poolId) view returns(uint256)
func (_PoolUtils *PoolUtilsSession) GetProtocolFee(_poolId uint8) (*big.Int, error) {
	return _PoolUtils.Contract.GetProtocolFee(&_PoolUtils.CallOpts, _poolId)
}

// GetProtocolFee is a free data retrieval call binding the contract method 0x261d41f5.
//
// Solidity: function getProtocolFee(uint8 _poolId) view returns(uint256)
func (_PoolUtils *PoolUtilsCallerSession) GetProtocolFee(_poolId uint8) (*big.Int, error) {
	return _PoolUtils.Contract.GetProtocolFee(&_PoolUtils.CallOpts, _poolId)
}

// GetQueuedValidatorCountByPool is a free data retrieval call binding the contract method 0xb7b32d4b.
//
// Solidity: function getQueuedValidatorCountByPool(uint8 _poolId) view returns(uint256)
func (_PoolUtils *PoolUtilsCaller) GetQueuedValidatorCountByPool(opts *bind.CallOpts, _poolId uint8) (*big.Int, error) {
	var out []interface{}
	err := _PoolUtils.contract.Call(opts, &out, "getQueuedValidatorCountByPool", _poolId)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetQueuedValidatorCountByPool is a free data retrieval call binding the contract method 0xb7b32d4b.
//
// Solidity: function getQueuedValidatorCountByPool(uint8 _poolId) view returns(uint256)
func (_PoolUtils *PoolUtilsSession) GetQueuedValidatorCountByPool(_poolId uint8) (*big.Int, error) {
	return _PoolUtils.Contract.GetQueuedValidatorCountByPool(&_PoolUtils.CallOpts, _poolId)
}

// GetQueuedValidatorCountByPool is a free data retrieval call binding the contract method 0xb7b32d4b.
//
// Solidity: function getQueuedValidatorCountByPool(uint8 _poolId) view returns(uint256)
func (_PoolUtils *PoolUtilsCallerSession) GetQueuedValidatorCountByPool(_poolId uint8) (*big.Int, error) {
	return _PoolUtils.Contract.GetQueuedValidatorCountByPool(&_PoolUtils.CallOpts, _poolId)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_PoolUtils *PoolUtilsCaller) GetRoleAdmin(opts *bind.CallOpts, role [32]byte) ([32]byte, error) {
	var out []interface{}
	err := _PoolUtils.contract.Call(opts, &out, "getRoleAdmin", role)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_PoolUtils *PoolUtilsSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _PoolUtils.Contract.GetRoleAdmin(&_PoolUtils.CallOpts, role)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_PoolUtils *PoolUtilsCallerSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _PoolUtils.Contract.GetRoleAdmin(&_PoolUtils.CallOpts, role)
}

// GetSocializingPoolAddress is a free data retrieval call binding the contract method 0x7526d429.
//
// Solidity: function getSocializingPoolAddress(uint8 _poolId) view returns(address)
func (_PoolUtils *PoolUtilsCaller) GetSocializingPoolAddress(opts *bind.CallOpts, _poolId uint8) (common.Address, error) {
	var out []interface{}
	err := _PoolUtils.contract.Call(opts, &out, "getSocializingPoolAddress", _poolId)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetSocializingPoolAddress is a free data retrieval call binding the contract method 0x7526d429.
//
// Solidity: function getSocializingPoolAddress(uint8 _poolId) view returns(address)
func (_PoolUtils *PoolUtilsSession) GetSocializingPoolAddress(_poolId uint8) (common.Address, error) {
	return _PoolUtils.Contract.GetSocializingPoolAddress(&_PoolUtils.CallOpts, _poolId)
}

// GetSocializingPoolAddress is a free data retrieval call binding the contract method 0x7526d429.
//
// Solidity: function getSocializingPoolAddress(uint8 _poolId) view returns(address)
func (_PoolUtils *PoolUtilsCallerSession) GetSocializingPoolAddress(_poolId uint8) (common.Address, error) {
	return _PoolUtils.Contract.GetSocializingPoolAddress(&_PoolUtils.CallOpts, _poolId)
}

// GetTotalActiveValidatorCount is a free data retrieval call binding the contract method 0x77c359e1.
//
// Solidity: function getTotalActiveValidatorCount() view returns(uint256)
func (_PoolUtils *PoolUtilsCaller) GetTotalActiveValidatorCount(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _PoolUtils.contract.Call(opts, &out, "getTotalActiveValidatorCount")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetTotalActiveValidatorCount is a free data retrieval call binding the contract method 0x77c359e1.
//
// Solidity: function getTotalActiveValidatorCount() view returns(uint256)
func (_PoolUtils *PoolUtilsSession) GetTotalActiveValidatorCount() (*big.Int, error) {
	return _PoolUtils.Contract.GetTotalActiveValidatorCount(&_PoolUtils.CallOpts)
}

// GetTotalActiveValidatorCount is a free data retrieval call binding the contract method 0x77c359e1.
//
// Solidity: function getTotalActiveValidatorCount() view returns(uint256)
func (_PoolUtils *PoolUtilsCallerSession) GetTotalActiveValidatorCount() (*big.Int, error) {
	return _PoolUtils.Contract.GetTotalActiveValidatorCount(&_PoolUtils.CallOpts)
}

// GetValidatorPoolId is a free data retrieval call binding the contract method 0xbda0bc89.
//
// Solidity: function getValidatorPoolId(bytes _pubkey) view returns(uint8)
func (_PoolUtils *PoolUtilsCaller) GetValidatorPoolId(opts *bind.CallOpts, _pubkey []byte) (uint8, error) {
	var out []interface{}
	err := _PoolUtils.contract.Call(opts, &out, "getValidatorPoolId", _pubkey)

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// GetValidatorPoolId is a free data retrieval call binding the contract method 0xbda0bc89.
//
// Solidity: function getValidatorPoolId(bytes _pubkey) view returns(uint8)
func (_PoolUtils *PoolUtilsSession) GetValidatorPoolId(_pubkey []byte) (uint8, error) {
	return _PoolUtils.Contract.GetValidatorPoolId(&_PoolUtils.CallOpts, _pubkey)
}

// GetValidatorPoolId is a free data retrieval call binding the contract method 0xbda0bc89.
//
// Solidity: function getValidatorPoolId(bytes _pubkey) view returns(uint8)
func (_PoolUtils *PoolUtilsCallerSession) GetValidatorPoolId(_pubkey []byte) (uint8, error) {
	return _PoolUtils.Contract.GetValidatorPoolId(&_PoolUtils.CallOpts, _pubkey)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_PoolUtils *PoolUtilsCaller) HasRole(opts *bind.CallOpts, role [32]byte, account common.Address) (bool, error) {
	var out []interface{}
	err := _PoolUtils.contract.Call(opts, &out, "hasRole", role, account)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_PoolUtils *PoolUtilsSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _PoolUtils.Contract.HasRole(&_PoolUtils.CallOpts, role, account)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_PoolUtils *PoolUtilsCallerSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _PoolUtils.Contract.HasRole(&_PoolUtils.CallOpts, role, account)
}

// IsExistingOperator is a free data retrieval call binding the contract method 0xf9c4dda4.
//
// Solidity: function isExistingOperator(address _operAddr) view returns(bool)
func (_PoolUtils *PoolUtilsCaller) IsExistingOperator(opts *bind.CallOpts, _operAddr common.Address) (bool, error) {
	var out []interface{}
	err := _PoolUtils.contract.Call(opts, &out, "isExistingOperator", _operAddr)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsExistingOperator is a free data retrieval call binding the contract method 0xf9c4dda4.
//
// Solidity: function isExistingOperator(address _operAddr) view returns(bool)
func (_PoolUtils *PoolUtilsSession) IsExistingOperator(_operAddr common.Address) (bool, error) {
	return _PoolUtils.Contract.IsExistingOperator(&_PoolUtils.CallOpts, _operAddr)
}

// IsExistingOperator is a free data retrieval call binding the contract method 0xf9c4dda4.
//
// Solidity: function isExistingOperator(address _operAddr) view returns(bool)
func (_PoolUtils *PoolUtilsCallerSession) IsExistingOperator(_operAddr common.Address) (bool, error) {
	return _PoolUtils.Contract.IsExistingOperator(&_PoolUtils.CallOpts, _operAddr)
}

// IsExistingPoolId is a free data retrieval call binding the contract method 0x6cdf1252.
//
// Solidity: function isExistingPoolId(uint8 _poolId) view returns(bool)
func (_PoolUtils *PoolUtilsCaller) IsExistingPoolId(opts *bind.CallOpts, _poolId uint8) (bool, error) {
	var out []interface{}
	err := _PoolUtils.contract.Call(opts, &out, "isExistingPoolId", _poolId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsExistingPoolId is a free data retrieval call binding the contract method 0x6cdf1252.
//
// Solidity: function isExistingPoolId(uint8 _poolId) view returns(bool)
func (_PoolUtils *PoolUtilsSession) IsExistingPoolId(_poolId uint8) (bool, error) {
	return _PoolUtils.Contract.IsExistingPoolId(&_PoolUtils.CallOpts, _poolId)
}

// IsExistingPoolId is a free data retrieval call binding the contract method 0x6cdf1252.
//
// Solidity: function isExistingPoolId(uint8 _poolId) view returns(bool)
func (_PoolUtils *PoolUtilsCallerSession) IsExistingPoolId(_poolId uint8) (bool, error) {
	return _PoolUtils.Contract.IsExistingPoolId(&_PoolUtils.CallOpts, _poolId)
}

// IsExistingPubkey is a free data retrieval call binding the contract method 0x36514d9f.
//
// Solidity: function isExistingPubkey(bytes _pubkey) view returns(bool)
func (_PoolUtils *PoolUtilsCaller) IsExistingPubkey(opts *bind.CallOpts, _pubkey []byte) (bool, error) {
	var out []interface{}
	err := _PoolUtils.contract.Call(opts, &out, "isExistingPubkey", _pubkey)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsExistingPubkey is a free data retrieval call binding the contract method 0x36514d9f.
//
// Solidity: function isExistingPubkey(bytes _pubkey) view returns(bool)
func (_PoolUtils *PoolUtilsSession) IsExistingPubkey(_pubkey []byte) (bool, error) {
	return _PoolUtils.Contract.IsExistingPubkey(&_PoolUtils.CallOpts, _pubkey)
}

// IsExistingPubkey is a free data retrieval call binding the contract method 0x36514d9f.
//
// Solidity: function isExistingPubkey(bytes _pubkey) view returns(bool)
func (_PoolUtils *PoolUtilsCallerSession) IsExistingPubkey(_pubkey []byte) (bool, error) {
	return _PoolUtils.Contract.IsExistingPubkey(&_PoolUtils.CallOpts, _pubkey)
}

// OnlyValidKeys is a free data retrieval call binding the contract method 0x9f55941b.
//
// Solidity: function onlyValidKeys(bytes _pubkey, bytes _preDepositSignature, bytes _depositSignature) view returns()
func (_PoolUtils *PoolUtilsCaller) OnlyValidKeys(opts *bind.CallOpts, _pubkey []byte, _preDepositSignature []byte, _depositSignature []byte) error {
	var out []interface{}
	err := _PoolUtils.contract.Call(opts, &out, "onlyValidKeys", _pubkey, _preDepositSignature, _depositSignature)

	if err != nil {
		return err
	}

	return err

}

// OnlyValidKeys is a free data retrieval call binding the contract method 0x9f55941b.
//
// Solidity: function onlyValidKeys(bytes _pubkey, bytes _preDepositSignature, bytes _depositSignature) view returns()
func (_PoolUtils *PoolUtilsSession) OnlyValidKeys(_pubkey []byte, _preDepositSignature []byte, _depositSignature []byte) error {
	return _PoolUtils.Contract.OnlyValidKeys(&_PoolUtils.CallOpts, _pubkey, _preDepositSignature, _depositSignature)
}

// OnlyValidKeys is a free data retrieval call binding the contract method 0x9f55941b.
//
// Solidity: function onlyValidKeys(bytes _pubkey, bytes _preDepositSignature, bytes _depositSignature) view returns()
func (_PoolUtils *PoolUtilsCallerSession) OnlyValidKeys(_pubkey []byte, _preDepositSignature []byte, _depositSignature []byte) error {
	return _PoolUtils.Contract.OnlyValidKeys(&_PoolUtils.CallOpts, _pubkey, _preDepositSignature, _depositSignature)
}

// OnlyValidName is a free data retrieval call binding the contract method 0x9f7053f5.
//
// Solidity: function onlyValidName(string _name) view returns()
func (_PoolUtils *PoolUtilsCaller) OnlyValidName(opts *bind.CallOpts, _name string) error {
	var out []interface{}
	err := _PoolUtils.contract.Call(opts, &out, "onlyValidName", _name)

	if err != nil {
		return err
	}

	return err

}

// OnlyValidName is a free data retrieval call binding the contract method 0x9f7053f5.
//
// Solidity: function onlyValidName(string _name) view returns()
func (_PoolUtils *PoolUtilsSession) OnlyValidName(_name string) error {
	return _PoolUtils.Contract.OnlyValidName(&_PoolUtils.CallOpts, _name)
}

// OnlyValidName is a free data retrieval call binding the contract method 0x9f7053f5.
//
// Solidity: function onlyValidName(string _name) view returns()
func (_PoolUtils *PoolUtilsCallerSession) OnlyValidName(_name string) error {
	return _PoolUtils.Contract.OnlyValidName(&_PoolUtils.CallOpts, _name)
}

// PoolAddressById is a free data retrieval call binding the contract method 0xdf8984fe.
//
// Solidity: function poolAddressById(uint8 ) view returns(address)
func (_PoolUtils *PoolUtilsCaller) PoolAddressById(opts *bind.CallOpts, arg0 uint8) (common.Address, error) {
	var out []interface{}
	err := _PoolUtils.contract.Call(opts, &out, "poolAddressById", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// PoolAddressById is a free data retrieval call binding the contract method 0xdf8984fe.
//
// Solidity: function poolAddressById(uint8 ) view returns(address)
func (_PoolUtils *PoolUtilsSession) PoolAddressById(arg0 uint8) (common.Address, error) {
	return _PoolUtils.Contract.PoolAddressById(&_PoolUtils.CallOpts, arg0)
}

// PoolAddressById is a free data retrieval call binding the contract method 0xdf8984fe.
//
// Solidity: function poolAddressById(uint8 ) view returns(address)
func (_PoolUtils *PoolUtilsCallerSession) PoolAddressById(arg0 uint8) (common.Address, error) {
	return _PoolUtils.Contract.PoolAddressById(&_PoolUtils.CallOpts, arg0)
}

// PoolIdArray is a free data retrieval call binding the contract method 0x8465bef5.
//
// Solidity: function poolIdArray(uint256 ) view returns(uint8)
func (_PoolUtils *PoolUtilsCaller) PoolIdArray(opts *bind.CallOpts, arg0 *big.Int) (uint8, error) {
	var out []interface{}
	err := _PoolUtils.contract.Call(opts, &out, "poolIdArray", arg0)

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// PoolIdArray is a free data retrieval call binding the contract method 0x8465bef5.
//
// Solidity: function poolIdArray(uint256 ) view returns(uint8)
func (_PoolUtils *PoolUtilsSession) PoolIdArray(arg0 *big.Int) (uint8, error) {
	return _PoolUtils.Contract.PoolIdArray(&_PoolUtils.CallOpts, arg0)
}

// PoolIdArray is a free data retrieval call binding the contract method 0x8465bef5.
//
// Solidity: function poolIdArray(uint256 ) view returns(uint8)
func (_PoolUtils *PoolUtilsCallerSession) PoolIdArray(arg0 *big.Int) (uint8, error) {
	return _PoolUtils.Contract.PoolIdArray(&_PoolUtils.CallOpts, arg0)
}

// StaderConfig is a free data retrieval call binding the contract method 0x490ffa35.
//
// Solidity: function staderConfig() view returns(address)
func (_PoolUtils *PoolUtilsCaller) StaderConfig(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _PoolUtils.contract.Call(opts, &out, "staderConfig")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// StaderConfig is a free data retrieval call binding the contract method 0x490ffa35.
//
// Solidity: function staderConfig() view returns(address)
func (_PoolUtils *PoolUtilsSession) StaderConfig() (common.Address, error) {
	return _PoolUtils.Contract.StaderConfig(&_PoolUtils.CallOpts)
}

// StaderConfig is a free data retrieval call binding the contract method 0x490ffa35.
//
// Solidity: function staderConfig() view returns(address)
func (_PoolUtils *PoolUtilsCallerSession) StaderConfig() (common.Address, error) {
	return _PoolUtils.Contract.StaderConfig(&_PoolUtils.CallOpts)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_PoolUtils *PoolUtilsCaller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _PoolUtils.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_PoolUtils *PoolUtilsSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _PoolUtils.Contract.SupportsInterface(&_PoolUtils.CallOpts, interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_PoolUtils *PoolUtilsCallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _PoolUtils.Contract.SupportsInterface(&_PoolUtils.CallOpts, interfaceId)
}

// AddNewPool is a paid mutator transaction binding the contract method 0x63d0d5c0.
//
// Solidity: function addNewPool(uint8 _poolId, address _poolAddress) returns()
func (_PoolUtils *PoolUtilsTransactor) AddNewPool(opts *bind.TransactOpts, _poolId uint8, _poolAddress common.Address) (*types.Transaction, error) {
	return _PoolUtils.contract.Transact(opts, "addNewPool", _poolId, _poolAddress)
}

// AddNewPool is a paid mutator transaction binding the contract method 0x63d0d5c0.
//
// Solidity: function addNewPool(uint8 _poolId, address _poolAddress) returns()
func (_PoolUtils *PoolUtilsSession) AddNewPool(_poolId uint8, _poolAddress common.Address) (*types.Transaction, error) {
	return _PoolUtils.Contract.AddNewPool(&_PoolUtils.TransactOpts, _poolId, _poolAddress)
}

// AddNewPool is a paid mutator transaction binding the contract method 0x63d0d5c0.
//
// Solidity: function addNewPool(uint8 _poolId, address _poolAddress) returns()
func (_PoolUtils *PoolUtilsTransactorSession) AddNewPool(_poolId uint8, _poolAddress common.Address) (*types.Transaction, error) {
	return _PoolUtils.Contract.AddNewPool(&_PoolUtils.TransactOpts, _poolId, _poolAddress)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_PoolUtils *PoolUtilsTransactor) GrantRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _PoolUtils.contract.Transact(opts, "grantRole", role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_PoolUtils *PoolUtilsSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _PoolUtils.Contract.GrantRole(&_PoolUtils.TransactOpts, role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_PoolUtils *PoolUtilsTransactorSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _PoolUtils.Contract.GrantRole(&_PoolUtils.TransactOpts, role, account)
}

// Initialize is a paid mutator transaction binding the contract method 0x485cc955.
//
// Solidity: function initialize(address _admin, address _staderConfig) returns()
func (_PoolUtils *PoolUtilsTransactor) Initialize(opts *bind.TransactOpts, _admin common.Address, _staderConfig common.Address) (*types.Transaction, error) {
	return _PoolUtils.contract.Transact(opts, "initialize", _admin, _staderConfig)
}

// Initialize is a paid mutator transaction binding the contract method 0x485cc955.
//
// Solidity: function initialize(address _admin, address _staderConfig) returns()
func (_PoolUtils *PoolUtilsSession) Initialize(_admin common.Address, _staderConfig common.Address) (*types.Transaction, error) {
	return _PoolUtils.Contract.Initialize(&_PoolUtils.TransactOpts, _admin, _staderConfig)
}

// Initialize is a paid mutator transaction binding the contract method 0x485cc955.
//
// Solidity: function initialize(address _admin, address _staderConfig) returns()
func (_PoolUtils *PoolUtilsTransactorSession) Initialize(_admin common.Address, _staderConfig common.Address) (*types.Transaction, error) {
	return _PoolUtils.Contract.Initialize(&_PoolUtils.TransactOpts, _admin, _staderConfig)
}

// ProcessValidatorExitList is a paid mutator transaction binding the contract method 0xff6bceec.
//
// Solidity: function processValidatorExitList(bytes[] _pubkeys) returns()
func (_PoolUtils *PoolUtilsTransactor) ProcessValidatorExitList(opts *bind.TransactOpts, _pubkeys [][]byte) (*types.Transaction, error) {
	return _PoolUtils.contract.Transact(opts, "processValidatorExitList", _pubkeys)
}

// ProcessValidatorExitList is a paid mutator transaction binding the contract method 0xff6bceec.
//
// Solidity: function processValidatorExitList(bytes[] _pubkeys) returns()
func (_PoolUtils *PoolUtilsSession) ProcessValidatorExitList(_pubkeys [][]byte) (*types.Transaction, error) {
	return _PoolUtils.Contract.ProcessValidatorExitList(&_PoolUtils.TransactOpts, _pubkeys)
}

// ProcessValidatorExitList is a paid mutator transaction binding the contract method 0xff6bceec.
//
// Solidity: function processValidatorExitList(bytes[] _pubkeys) returns()
func (_PoolUtils *PoolUtilsTransactorSession) ProcessValidatorExitList(_pubkeys [][]byte) (*types.Transaction, error) {
	return _PoolUtils.Contract.ProcessValidatorExitList(&_PoolUtils.TransactOpts, _pubkeys)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address account) returns()
func (_PoolUtils *PoolUtilsTransactor) RenounceRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _PoolUtils.contract.Transact(opts, "renounceRole", role, account)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address account) returns()
func (_PoolUtils *PoolUtilsSession) RenounceRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _PoolUtils.Contract.RenounceRole(&_PoolUtils.TransactOpts, role, account)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address account) returns()
func (_PoolUtils *PoolUtilsTransactorSession) RenounceRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _PoolUtils.Contract.RenounceRole(&_PoolUtils.TransactOpts, role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_PoolUtils *PoolUtilsTransactor) RevokeRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _PoolUtils.contract.Transact(opts, "revokeRole", role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_PoolUtils *PoolUtilsSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _PoolUtils.Contract.RevokeRole(&_PoolUtils.TransactOpts, role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_PoolUtils *PoolUtilsTransactorSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _PoolUtils.Contract.RevokeRole(&_PoolUtils.TransactOpts, role, account)
}

// UpdatePoolAddress is a paid mutator transaction binding the contract method 0xd97ac847.
//
// Solidity: function updatePoolAddress(uint8 _poolId, address _newPoolAddress) returns()
func (_PoolUtils *PoolUtilsTransactor) UpdatePoolAddress(opts *bind.TransactOpts, _poolId uint8, _newPoolAddress common.Address) (*types.Transaction, error) {
	return _PoolUtils.contract.Transact(opts, "updatePoolAddress", _poolId, _newPoolAddress)
}

// UpdatePoolAddress is a paid mutator transaction binding the contract method 0xd97ac847.
//
// Solidity: function updatePoolAddress(uint8 _poolId, address _newPoolAddress) returns()
func (_PoolUtils *PoolUtilsSession) UpdatePoolAddress(_poolId uint8, _newPoolAddress common.Address) (*types.Transaction, error) {
	return _PoolUtils.Contract.UpdatePoolAddress(&_PoolUtils.TransactOpts, _poolId, _newPoolAddress)
}

// UpdatePoolAddress is a paid mutator transaction binding the contract method 0xd97ac847.
//
// Solidity: function updatePoolAddress(uint8 _poolId, address _newPoolAddress) returns()
func (_PoolUtils *PoolUtilsTransactorSession) UpdatePoolAddress(_poolId uint8, _newPoolAddress common.Address) (*types.Transaction, error) {
	return _PoolUtils.Contract.UpdatePoolAddress(&_PoolUtils.TransactOpts, _poolId, _newPoolAddress)
}

// UpdateStaderConfig is a paid mutator transaction binding the contract method 0x9ee804cb.
//
// Solidity: function updateStaderConfig(address _staderConfig) returns()
func (_PoolUtils *PoolUtilsTransactor) UpdateStaderConfig(opts *bind.TransactOpts, _staderConfig common.Address) (*types.Transaction, error) {
	return _PoolUtils.contract.Transact(opts, "updateStaderConfig", _staderConfig)
}

// UpdateStaderConfig is a paid mutator transaction binding the contract method 0x9ee804cb.
//
// Solidity: function updateStaderConfig(address _staderConfig) returns()
func (_PoolUtils *PoolUtilsSession) UpdateStaderConfig(_staderConfig common.Address) (*types.Transaction, error) {
	return _PoolUtils.Contract.UpdateStaderConfig(&_PoolUtils.TransactOpts, _staderConfig)
}

// UpdateStaderConfig is a paid mutator transaction binding the contract method 0x9ee804cb.
//
// Solidity: function updateStaderConfig(address _staderConfig) returns()
func (_PoolUtils *PoolUtilsTransactorSession) UpdateStaderConfig(_staderConfig common.Address) (*types.Transaction, error) {
	return _PoolUtils.Contract.UpdateStaderConfig(&_PoolUtils.TransactOpts, _staderConfig)
}

// PoolUtilsDeactivatedPoolIterator is returned from FilterDeactivatedPool and is used to iterate over the raw logs and unpacked data for DeactivatedPool events raised by the PoolUtils contract.
type PoolUtilsDeactivatedPoolIterator struct {
	Event *PoolUtilsDeactivatedPool // Event containing the contract specifics and raw log

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
func (it *PoolUtilsDeactivatedPoolIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PoolUtilsDeactivatedPool)
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
		it.Event = new(PoolUtilsDeactivatedPool)
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
func (it *PoolUtilsDeactivatedPoolIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PoolUtilsDeactivatedPoolIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PoolUtilsDeactivatedPool represents a DeactivatedPool event raised by the PoolUtils contract.
type PoolUtilsDeactivatedPool struct {
	PoolId      uint8
	PoolAddress common.Address
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterDeactivatedPool is a free log retrieval operation binding the contract event 0xf711845001a9a7fade2a40e12b1fb02c31952a41f7c999ae1f84a283d32671f6.
//
// Solidity: event DeactivatedPool(uint8 indexed poolId, address poolAddress)
func (_PoolUtils *PoolUtilsFilterer) FilterDeactivatedPool(opts *bind.FilterOpts, poolId []uint8) (*PoolUtilsDeactivatedPoolIterator, error) {

	var poolIdRule []interface{}
	for _, poolIdItem := range poolId {
		poolIdRule = append(poolIdRule, poolIdItem)
	}

	logs, sub, err := _PoolUtils.contract.FilterLogs(opts, "DeactivatedPool", poolIdRule)
	if err != nil {
		return nil, err
	}
	return &PoolUtilsDeactivatedPoolIterator{contract: _PoolUtils.contract, event: "DeactivatedPool", logs: logs, sub: sub}, nil
}

// WatchDeactivatedPool is a free log subscription operation binding the contract event 0xf711845001a9a7fade2a40e12b1fb02c31952a41f7c999ae1f84a283d32671f6.
//
// Solidity: event DeactivatedPool(uint8 indexed poolId, address poolAddress)
func (_PoolUtils *PoolUtilsFilterer) WatchDeactivatedPool(opts *bind.WatchOpts, sink chan<- *PoolUtilsDeactivatedPool, poolId []uint8) (event.Subscription, error) {

	var poolIdRule []interface{}
	for _, poolIdItem := range poolId {
		poolIdRule = append(poolIdRule, poolIdItem)
	}

	logs, sub, err := _PoolUtils.contract.WatchLogs(opts, "DeactivatedPool", poolIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PoolUtilsDeactivatedPool)
				if err := _PoolUtils.contract.UnpackLog(event, "DeactivatedPool", log); err != nil {
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

// ParseDeactivatedPool is a log parse operation binding the contract event 0xf711845001a9a7fade2a40e12b1fb02c31952a41f7c999ae1f84a283d32671f6.
//
// Solidity: event DeactivatedPool(uint8 indexed poolId, address poolAddress)
func (_PoolUtils *PoolUtilsFilterer) ParseDeactivatedPool(log types.Log) (*PoolUtilsDeactivatedPool, error) {
	event := new(PoolUtilsDeactivatedPool)
	if err := _PoolUtils.contract.UnpackLog(event, "DeactivatedPool", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PoolUtilsExitValidatorIterator is returned from FilterExitValidator and is used to iterate over the raw logs and unpacked data for ExitValidator events raised by the PoolUtils contract.
type PoolUtilsExitValidatorIterator struct {
	Event *PoolUtilsExitValidator // Event containing the contract specifics and raw log

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
func (it *PoolUtilsExitValidatorIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PoolUtilsExitValidator)
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
		it.Event = new(PoolUtilsExitValidator)
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
func (it *PoolUtilsExitValidatorIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PoolUtilsExitValidatorIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PoolUtilsExitValidator represents a ExitValidator event raised by the PoolUtils contract.
type PoolUtilsExitValidator struct {
	Pubkey []byte
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterExitValidator is a free log retrieval operation binding the contract event 0xce4931e23262c5aa14c3b95b5e67c07bb38447fda706a4e5e4019e4f70142812.
//
// Solidity: event ExitValidator(bytes pubkey)
func (_PoolUtils *PoolUtilsFilterer) FilterExitValidator(opts *bind.FilterOpts) (*PoolUtilsExitValidatorIterator, error) {

	logs, sub, err := _PoolUtils.contract.FilterLogs(opts, "ExitValidator")
	if err != nil {
		return nil, err
	}
	return &PoolUtilsExitValidatorIterator{contract: _PoolUtils.contract, event: "ExitValidator", logs: logs, sub: sub}, nil
}

// WatchExitValidator is a free log subscription operation binding the contract event 0xce4931e23262c5aa14c3b95b5e67c07bb38447fda706a4e5e4019e4f70142812.
//
// Solidity: event ExitValidator(bytes pubkey)
func (_PoolUtils *PoolUtilsFilterer) WatchExitValidator(opts *bind.WatchOpts, sink chan<- *PoolUtilsExitValidator) (event.Subscription, error) {

	logs, sub, err := _PoolUtils.contract.WatchLogs(opts, "ExitValidator")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PoolUtilsExitValidator)
				if err := _PoolUtils.contract.UnpackLog(event, "ExitValidator", log); err != nil {
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

// ParseExitValidator is a log parse operation binding the contract event 0xce4931e23262c5aa14c3b95b5e67c07bb38447fda706a4e5e4019e4f70142812.
//
// Solidity: event ExitValidator(bytes pubkey)
func (_PoolUtils *PoolUtilsFilterer) ParseExitValidator(log types.Log) (*PoolUtilsExitValidator, error) {
	event := new(PoolUtilsExitValidator)
	if err := _PoolUtils.contract.UnpackLog(event, "ExitValidator", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PoolUtilsInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the PoolUtils contract.
type PoolUtilsInitializedIterator struct {
	Event *PoolUtilsInitialized // Event containing the contract specifics and raw log

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
func (it *PoolUtilsInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PoolUtilsInitialized)
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
		it.Event = new(PoolUtilsInitialized)
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
func (it *PoolUtilsInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PoolUtilsInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PoolUtilsInitialized represents a Initialized event raised by the PoolUtils contract.
type PoolUtilsInitialized struct {
	Version uint8
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_PoolUtils *PoolUtilsFilterer) FilterInitialized(opts *bind.FilterOpts) (*PoolUtilsInitializedIterator, error) {

	logs, sub, err := _PoolUtils.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &PoolUtilsInitializedIterator{contract: _PoolUtils.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_PoolUtils *PoolUtilsFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *PoolUtilsInitialized) (event.Subscription, error) {

	logs, sub, err := _PoolUtils.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PoolUtilsInitialized)
				if err := _PoolUtils.contract.UnpackLog(event, "Initialized", log); err != nil {
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
func (_PoolUtils *PoolUtilsFilterer) ParseInitialized(log types.Log) (*PoolUtilsInitialized, error) {
	event := new(PoolUtilsInitialized)
	if err := _PoolUtils.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PoolUtilsPoolAddedIterator is returned from FilterPoolAdded and is used to iterate over the raw logs and unpacked data for PoolAdded events raised by the PoolUtils contract.
type PoolUtilsPoolAddedIterator struct {
	Event *PoolUtilsPoolAdded // Event containing the contract specifics and raw log

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
func (it *PoolUtilsPoolAddedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PoolUtilsPoolAdded)
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
		it.Event = new(PoolUtilsPoolAdded)
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
func (it *PoolUtilsPoolAddedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PoolUtilsPoolAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PoolUtilsPoolAdded represents a PoolAdded event raised by the PoolUtils contract.
type PoolUtilsPoolAdded struct {
	PoolId      uint8
	PoolAddress common.Address
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterPoolAdded is a free log retrieval operation binding the contract event 0x697362d5a2939aff718fb2db4145cb1b4ffc68872c82b2e64d805d8e94845af1.
//
// Solidity: event PoolAdded(uint8 indexed poolId, address poolAddress)
func (_PoolUtils *PoolUtilsFilterer) FilterPoolAdded(opts *bind.FilterOpts, poolId []uint8) (*PoolUtilsPoolAddedIterator, error) {

	var poolIdRule []interface{}
	for _, poolIdItem := range poolId {
		poolIdRule = append(poolIdRule, poolIdItem)
	}

	logs, sub, err := _PoolUtils.contract.FilterLogs(opts, "PoolAdded", poolIdRule)
	if err != nil {
		return nil, err
	}
	return &PoolUtilsPoolAddedIterator{contract: _PoolUtils.contract, event: "PoolAdded", logs: logs, sub: sub}, nil
}

// WatchPoolAdded is a free log subscription operation binding the contract event 0x697362d5a2939aff718fb2db4145cb1b4ffc68872c82b2e64d805d8e94845af1.
//
// Solidity: event PoolAdded(uint8 indexed poolId, address poolAddress)
func (_PoolUtils *PoolUtilsFilterer) WatchPoolAdded(opts *bind.WatchOpts, sink chan<- *PoolUtilsPoolAdded, poolId []uint8) (event.Subscription, error) {

	var poolIdRule []interface{}
	for _, poolIdItem := range poolId {
		poolIdRule = append(poolIdRule, poolIdItem)
	}

	logs, sub, err := _PoolUtils.contract.WatchLogs(opts, "PoolAdded", poolIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PoolUtilsPoolAdded)
				if err := _PoolUtils.contract.UnpackLog(event, "PoolAdded", log); err != nil {
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

// ParsePoolAdded is a log parse operation binding the contract event 0x697362d5a2939aff718fb2db4145cb1b4ffc68872c82b2e64d805d8e94845af1.
//
// Solidity: event PoolAdded(uint8 indexed poolId, address poolAddress)
func (_PoolUtils *PoolUtilsFilterer) ParsePoolAdded(log types.Log) (*PoolUtilsPoolAdded, error) {
	event := new(PoolUtilsPoolAdded)
	if err := _PoolUtils.contract.UnpackLog(event, "PoolAdded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PoolUtilsPoolAddressUpdatedIterator is returned from FilterPoolAddressUpdated and is used to iterate over the raw logs and unpacked data for PoolAddressUpdated events raised by the PoolUtils contract.
type PoolUtilsPoolAddressUpdatedIterator struct {
	Event *PoolUtilsPoolAddressUpdated // Event containing the contract specifics and raw log

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
func (it *PoolUtilsPoolAddressUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PoolUtilsPoolAddressUpdated)
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
		it.Event = new(PoolUtilsPoolAddressUpdated)
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
func (it *PoolUtilsPoolAddressUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PoolUtilsPoolAddressUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PoolUtilsPoolAddressUpdated represents a PoolAddressUpdated event raised by the PoolUtils contract.
type PoolUtilsPoolAddressUpdated struct {
	PoolId      uint8
	PoolAddress common.Address
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterPoolAddressUpdated is a free log retrieval operation binding the contract event 0xf732deab68331ad20834cfc15d686fed4bac945cf3af5d7f729205c9bf846199.
//
// Solidity: event PoolAddressUpdated(uint8 indexed poolId, address poolAddress)
func (_PoolUtils *PoolUtilsFilterer) FilterPoolAddressUpdated(opts *bind.FilterOpts, poolId []uint8) (*PoolUtilsPoolAddressUpdatedIterator, error) {

	var poolIdRule []interface{}
	for _, poolIdItem := range poolId {
		poolIdRule = append(poolIdRule, poolIdItem)
	}

	logs, sub, err := _PoolUtils.contract.FilterLogs(opts, "PoolAddressUpdated", poolIdRule)
	if err != nil {
		return nil, err
	}
	return &PoolUtilsPoolAddressUpdatedIterator{contract: _PoolUtils.contract, event: "PoolAddressUpdated", logs: logs, sub: sub}, nil
}

// WatchPoolAddressUpdated is a free log subscription operation binding the contract event 0xf732deab68331ad20834cfc15d686fed4bac945cf3af5d7f729205c9bf846199.
//
// Solidity: event PoolAddressUpdated(uint8 indexed poolId, address poolAddress)
func (_PoolUtils *PoolUtilsFilterer) WatchPoolAddressUpdated(opts *bind.WatchOpts, sink chan<- *PoolUtilsPoolAddressUpdated, poolId []uint8) (event.Subscription, error) {

	var poolIdRule []interface{}
	for _, poolIdItem := range poolId {
		poolIdRule = append(poolIdRule, poolIdItem)
	}

	logs, sub, err := _PoolUtils.contract.WatchLogs(opts, "PoolAddressUpdated", poolIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PoolUtilsPoolAddressUpdated)
				if err := _PoolUtils.contract.UnpackLog(event, "PoolAddressUpdated", log); err != nil {
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

// ParsePoolAddressUpdated is a log parse operation binding the contract event 0xf732deab68331ad20834cfc15d686fed4bac945cf3af5d7f729205c9bf846199.
//
// Solidity: event PoolAddressUpdated(uint8 indexed poolId, address poolAddress)
func (_PoolUtils *PoolUtilsFilterer) ParsePoolAddressUpdated(log types.Log) (*PoolUtilsPoolAddressUpdated, error) {
	event := new(PoolUtilsPoolAddressUpdated)
	if err := _PoolUtils.contract.UnpackLog(event, "PoolAddressUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PoolUtilsRoleAdminChangedIterator is returned from FilterRoleAdminChanged and is used to iterate over the raw logs and unpacked data for RoleAdminChanged events raised by the PoolUtils contract.
type PoolUtilsRoleAdminChangedIterator struct {
	Event *PoolUtilsRoleAdminChanged // Event containing the contract specifics and raw log

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
func (it *PoolUtilsRoleAdminChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PoolUtilsRoleAdminChanged)
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
		it.Event = new(PoolUtilsRoleAdminChanged)
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
func (it *PoolUtilsRoleAdminChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PoolUtilsRoleAdminChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PoolUtilsRoleAdminChanged represents a RoleAdminChanged event raised by the PoolUtils contract.
type PoolUtilsRoleAdminChanged struct {
	Role              [32]byte
	PreviousAdminRole [32]byte
	NewAdminRole      [32]byte
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterRoleAdminChanged is a free log retrieval operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_PoolUtils *PoolUtilsFilterer) FilterRoleAdminChanged(opts *bind.FilterOpts, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (*PoolUtilsRoleAdminChangedIterator, error) {

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

	logs, sub, err := _PoolUtils.contract.FilterLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return &PoolUtilsRoleAdminChangedIterator{contract: _PoolUtils.contract, event: "RoleAdminChanged", logs: logs, sub: sub}, nil
}

// WatchRoleAdminChanged is a free log subscription operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_PoolUtils *PoolUtilsFilterer) WatchRoleAdminChanged(opts *bind.WatchOpts, sink chan<- *PoolUtilsRoleAdminChanged, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (event.Subscription, error) {

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

	logs, sub, err := _PoolUtils.contract.WatchLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PoolUtilsRoleAdminChanged)
				if err := _PoolUtils.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
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
func (_PoolUtils *PoolUtilsFilterer) ParseRoleAdminChanged(log types.Log) (*PoolUtilsRoleAdminChanged, error) {
	event := new(PoolUtilsRoleAdminChanged)
	if err := _PoolUtils.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PoolUtilsRoleGrantedIterator is returned from FilterRoleGranted and is used to iterate over the raw logs and unpacked data for RoleGranted events raised by the PoolUtils contract.
type PoolUtilsRoleGrantedIterator struct {
	Event *PoolUtilsRoleGranted // Event containing the contract specifics and raw log

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
func (it *PoolUtilsRoleGrantedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PoolUtilsRoleGranted)
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
		it.Event = new(PoolUtilsRoleGranted)
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
func (it *PoolUtilsRoleGrantedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PoolUtilsRoleGrantedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PoolUtilsRoleGranted represents a RoleGranted event raised by the PoolUtils contract.
type PoolUtilsRoleGranted struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleGranted is a free log retrieval operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_PoolUtils *PoolUtilsFilterer) FilterRoleGranted(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*PoolUtilsRoleGrantedIterator, error) {

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

	logs, sub, err := _PoolUtils.contract.FilterLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &PoolUtilsRoleGrantedIterator{contract: _PoolUtils.contract, event: "RoleGranted", logs: logs, sub: sub}, nil
}

// WatchRoleGranted is a free log subscription operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_PoolUtils *PoolUtilsFilterer) WatchRoleGranted(opts *bind.WatchOpts, sink chan<- *PoolUtilsRoleGranted, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _PoolUtils.contract.WatchLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PoolUtilsRoleGranted)
				if err := _PoolUtils.contract.UnpackLog(event, "RoleGranted", log); err != nil {
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
func (_PoolUtils *PoolUtilsFilterer) ParseRoleGranted(log types.Log) (*PoolUtilsRoleGranted, error) {
	event := new(PoolUtilsRoleGranted)
	if err := _PoolUtils.contract.UnpackLog(event, "RoleGranted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PoolUtilsRoleRevokedIterator is returned from FilterRoleRevoked and is used to iterate over the raw logs and unpacked data for RoleRevoked events raised by the PoolUtils contract.
type PoolUtilsRoleRevokedIterator struct {
	Event *PoolUtilsRoleRevoked // Event containing the contract specifics and raw log

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
func (it *PoolUtilsRoleRevokedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PoolUtilsRoleRevoked)
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
		it.Event = new(PoolUtilsRoleRevoked)
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
func (it *PoolUtilsRoleRevokedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PoolUtilsRoleRevokedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PoolUtilsRoleRevoked represents a RoleRevoked event raised by the PoolUtils contract.
type PoolUtilsRoleRevoked struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleRevoked is a free log retrieval operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_PoolUtils *PoolUtilsFilterer) FilterRoleRevoked(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*PoolUtilsRoleRevokedIterator, error) {

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

	logs, sub, err := _PoolUtils.contract.FilterLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &PoolUtilsRoleRevokedIterator{contract: _PoolUtils.contract, event: "RoleRevoked", logs: logs, sub: sub}, nil
}

// WatchRoleRevoked is a free log subscription operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_PoolUtils *PoolUtilsFilterer) WatchRoleRevoked(opts *bind.WatchOpts, sink chan<- *PoolUtilsRoleRevoked, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _PoolUtils.contract.WatchLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PoolUtilsRoleRevoked)
				if err := _PoolUtils.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
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
func (_PoolUtils *PoolUtilsFilterer) ParseRoleRevoked(log types.Log) (*PoolUtilsRoleRevoked, error) {
	event := new(PoolUtilsRoleRevoked)
	if err := _PoolUtils.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PoolUtilsUpdatedStaderConfigIterator is returned from FilterUpdatedStaderConfig and is used to iterate over the raw logs and unpacked data for UpdatedStaderConfig events raised by the PoolUtils contract.
type PoolUtilsUpdatedStaderConfigIterator struct {
	Event *PoolUtilsUpdatedStaderConfig // Event containing the contract specifics and raw log

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
func (it *PoolUtilsUpdatedStaderConfigIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PoolUtilsUpdatedStaderConfig)
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
		it.Event = new(PoolUtilsUpdatedStaderConfig)
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
func (it *PoolUtilsUpdatedStaderConfigIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PoolUtilsUpdatedStaderConfigIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PoolUtilsUpdatedStaderConfig represents a UpdatedStaderConfig event raised by the PoolUtils contract.
type PoolUtilsUpdatedStaderConfig struct {
	StaderConfig common.Address
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterUpdatedStaderConfig is a free log retrieval operation binding the contract event 0xdb2219043d7b197cb235f1af0cf6d782d77dee3de19e3f4fb6d39aae633b4485.
//
// Solidity: event UpdatedStaderConfig(address staderConfig)
func (_PoolUtils *PoolUtilsFilterer) FilterUpdatedStaderConfig(opts *bind.FilterOpts) (*PoolUtilsUpdatedStaderConfigIterator, error) {

	logs, sub, err := _PoolUtils.contract.FilterLogs(opts, "UpdatedStaderConfig")
	if err != nil {
		return nil, err
	}
	return &PoolUtilsUpdatedStaderConfigIterator{contract: _PoolUtils.contract, event: "UpdatedStaderConfig", logs: logs, sub: sub}, nil
}

// WatchUpdatedStaderConfig is a free log subscription operation binding the contract event 0xdb2219043d7b197cb235f1af0cf6d782d77dee3de19e3f4fb6d39aae633b4485.
//
// Solidity: event UpdatedStaderConfig(address staderConfig)
func (_PoolUtils *PoolUtilsFilterer) WatchUpdatedStaderConfig(opts *bind.WatchOpts, sink chan<- *PoolUtilsUpdatedStaderConfig) (event.Subscription, error) {

	logs, sub, err := _PoolUtils.contract.WatchLogs(opts, "UpdatedStaderConfig")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PoolUtilsUpdatedStaderConfig)
				if err := _PoolUtils.contract.UnpackLog(event, "UpdatedStaderConfig", log); err != nil {
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
func (_PoolUtils *PoolUtilsFilterer) ParseUpdatedStaderConfig(log types.Log) (*PoolUtilsUpdatedStaderConfig, error) {
	event := new(PoolUtilsUpdatedStaderConfig)
	if err := _PoolUtils.contract.UnpackLog(event, "UpdatedStaderConfig", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
