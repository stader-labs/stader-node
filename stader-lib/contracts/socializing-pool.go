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

// RewardsData is an auto generated low-level Go binding around an user-defined struct.
type RewardsData struct {
	ReportingBlockNumber *big.Int
	Index                *big.Int
	MerkleRoot           [32]byte
	PoolId               uint8
	OperatorETHRewards   *big.Int
	UserETHRewards       *big.Int
	ProtocolETHRewards   *big.Int
	OperatorSDRewards    *big.Int
}

// SocializingPoolMetaData contains all meta data concerning the SocializingPool contract.
var SocializingPoolMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"CallerNotManager\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"CallerNotStaderContract\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"ETHTransferFailed\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"FutureCycleIndex\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InsufficientETHRewards\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InsufficientSDRewards\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidAmount\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidCycleIndex\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"cycle\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"}],\"name\":\"InvalidProof\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"cycle\",\"type\":\"uint256\"}],\"name\":\"RewardAlreadyClaimed\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"RewardAlreadyHandled\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"SDTransferFailed\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ZeroAddress\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"ETHReceived\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"version\",\"type\":\"uint8\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"ethRewards\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"sdRewards\",\"type\":\"uint256\"}],\"name\":\"OperatorRewardsClaimed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"ethRewards\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"totalETHRewards\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"sdRewards\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"totalSDRewards\",\"type\":\"uint256\"}],\"name\":\"OperatorRewardsUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Paused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"ethRewards\",\"type\":\"uint256\"}],\"name\":\"ProtocolETHRewardsTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"previousAdminRole\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"newAdminRole\",\"type\":\"bytes32\"}],\"name\":\"RoleAdminChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleGranted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleRevoked\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Unpaused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"staderConfig\",\"type\":\"address\"}],\"name\":\"UpdatedStaderConfig\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"staderOperatorRegistry\",\"type\":\"address\"}],\"name\":\"UpdatedStaderOperatorRegistry\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"staderValidatorRegistry\",\"type\":\"address\"}],\"name\":\"UpdatedStaderValidatorRegistry\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"ethRewards\",\"type\":\"uint256\"}],\"name\":\"UserETHRewardsTransferred\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"DEFAULT_ADMIN_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256[]\",\"name\":\"_index\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256[]\",\"name\":\"_amountSD\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256[]\",\"name\":\"_amountETH\",\"type\":\"uint256[]\"},{\"internalType\":\"bytes32[][]\",\"name\":\"_merkleProof\",\"type\":\"bytes32[][]\"}],\"name\":\"claim\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"claimedRewards\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getCurrentRewardsIndex\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_index\",\"type\":\"uint256\"}],\"name\":\"getRewardCycleDetails\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"_startBlock\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_endBlock\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getRewardDetails\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"currentIndex\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"currentStartBlock\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"currentEndBlock\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"}],\"name\":\"getRoleAdmin\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"grantRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"reportingBlockNumber\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"merkleRoot\",\"type\":\"bytes32\"},{\"internalType\":\"uint8\",\"name\":\"poolId\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"operatorETHRewards\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"userETHRewards\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"protocolETHRewards\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"operatorSDRewards\",\"type\":\"uint256\"}],\"internalType\":\"structRewardsData\",\"name\":\"_rewardsData\",\"type\":\"tuple\"}],\"name\":\"handleRewards\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"handledRewards\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"hasRole\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"initialBlock\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_admin\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_staderConfig\",\"type\":\"address\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"lastReportedRewardsData\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"reportingBlockNumber\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"merkleRoot\",\"type\":\"bytes32\"},{\"internalType\":\"uint8\",\"name\":\"poolId\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"operatorETHRewards\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"userETHRewards\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"protocolETHRewards\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"operatorSDRewards\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"paused\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"renounceRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"revokeRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"rewardsDataMap\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"reportingBlockNumber\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"merkleRoot\",\"type\":\"bytes32\"},{\"internalType\":\"uint8\",\"name\":\"poolId\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"operatorETHRewards\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"userETHRewards\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"protocolETHRewards\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"operatorSDRewards\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"staderConfig\",\"outputs\":[{\"internalType\":\"contractIStaderConfig\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalOperatorETHRewardsRemaining\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalOperatorSDRewardsRemaining\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"unpause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_staderConfig\",\"type\":\"address\"}],\"name\":\"updateStaderConfig\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_index\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_operator\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amountSD\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_amountETH\",\"type\":\"uint256\"},{\"internalType\":\"bytes32[]\",\"name\":\"_merkleProof\",\"type\":\"bytes32[]\"}],\"name\":\"verifyProof\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"stateMutability\":\"payable\",\"type\":\"receive\"}]",
}

// SocializingPoolABI is the input ABI used to generate the binding from.
// Deprecated: Use SocializingPoolMetaData.ABI instead.
var SocializingPoolABI = SocializingPoolMetaData.ABI

// SocializingPool is an auto generated Go binding around an Ethereum contract.
type SocializingPool struct {
	SocializingPoolCaller     // Read-only binding to the contract
	SocializingPoolTransactor // Write-only binding to the contract
	SocializingPoolFilterer   // Log filterer for contract events
}

// SocializingPoolCaller is an auto generated read-only Go binding around an Ethereum contract.
type SocializingPoolCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SocializingPoolTransactor is an auto generated write-only Go binding around an Ethereum contract.
type SocializingPoolTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SocializingPoolFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type SocializingPoolFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SocializingPoolSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type SocializingPoolSession struct {
	Contract     *SocializingPool  // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// SocializingPoolCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type SocializingPoolCallerSession struct {
	Contract *SocializingPoolCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts          // Call options to use throughout this session
}

// SocializingPoolTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type SocializingPoolTransactorSession struct {
	Contract     *SocializingPoolTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts          // Transaction auth options to use throughout this session
}

// SocializingPoolRaw is an auto generated low-level Go binding around an Ethereum contract.
type SocializingPoolRaw struct {
	Contract *SocializingPool // Generic contract binding to access the raw methods on
}

// SocializingPoolCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type SocializingPoolCallerRaw struct {
	Contract *SocializingPoolCaller // Generic read-only contract binding to access the raw methods on
}

// SocializingPoolTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type SocializingPoolTransactorRaw struct {
	Contract *SocializingPoolTransactor // Generic write-only contract binding to access the raw methods on
}

// NewSocializingPool creates a new instance of SocializingPool, bound to a specific deployed contract.
func NewSocializingPool(address common.Address, backend bind.ContractBackend) (*SocializingPool, error) {
	contract, err := bindSocializingPool(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &SocializingPool{SocializingPoolCaller: SocializingPoolCaller{contract: contract}, SocializingPoolTransactor: SocializingPoolTransactor{contract: contract}, SocializingPoolFilterer: SocializingPoolFilterer{contract: contract}}, nil
}

// NewSocializingPoolCaller creates a new read-only instance of SocializingPool, bound to a specific deployed contract.
func NewSocializingPoolCaller(address common.Address, caller bind.ContractCaller) (*SocializingPoolCaller, error) {
	contract, err := bindSocializingPool(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &SocializingPoolCaller{contract: contract}, nil
}

// NewSocializingPoolTransactor creates a new write-only instance of SocializingPool, bound to a specific deployed contract.
func NewSocializingPoolTransactor(address common.Address, transactor bind.ContractTransactor) (*SocializingPoolTransactor, error) {
	contract, err := bindSocializingPool(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &SocializingPoolTransactor{contract: contract}, nil
}

// NewSocializingPoolFilterer creates a new log filterer instance of SocializingPool, bound to a specific deployed contract.
func NewSocializingPoolFilterer(address common.Address, filterer bind.ContractFilterer) (*SocializingPoolFilterer, error) {
	contract, err := bindSocializingPool(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &SocializingPoolFilterer{contract: contract}, nil
}

// bindSocializingPool binds a generic wrapper to an already deployed contract.
func bindSocializingPool(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := SocializingPoolMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SocializingPool *SocializingPoolRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _SocializingPool.Contract.SocializingPoolCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SocializingPool *SocializingPoolRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SocializingPool.Contract.SocializingPoolTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SocializingPool *SocializingPoolRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SocializingPool.Contract.SocializingPoolTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SocializingPool *SocializingPoolCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _SocializingPool.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SocializingPool *SocializingPoolTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SocializingPool.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SocializingPool *SocializingPoolTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SocializingPool.Contract.contract.Transact(opts, method, params...)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_SocializingPool *SocializingPoolCaller) DEFAULTADMINROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _SocializingPool.contract.Call(opts, &out, "DEFAULT_ADMIN_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_SocializingPool *SocializingPoolSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _SocializingPool.Contract.DEFAULTADMINROLE(&_SocializingPool.CallOpts)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_SocializingPool *SocializingPoolCallerSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _SocializingPool.Contract.DEFAULTADMINROLE(&_SocializingPool.CallOpts)
}

// ClaimedRewards is a free data retrieval call binding the contract method 0xfb831b9a.
//
// Solidity: function claimedRewards(address , uint256 ) view returns(bool)
func (_SocializingPool *SocializingPoolCaller) ClaimedRewards(opts *bind.CallOpts, arg0 common.Address, arg1 *big.Int) (bool, error) {
	var out []interface{}
	err := _SocializingPool.contract.Call(opts, &out, "claimedRewards", arg0, arg1)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// ClaimedRewards is a free data retrieval call binding the contract method 0xfb831b9a.
//
// Solidity: function claimedRewards(address , uint256 ) view returns(bool)
func (_SocializingPool *SocializingPoolSession) ClaimedRewards(arg0 common.Address, arg1 *big.Int) (bool, error) {
	return _SocializingPool.Contract.ClaimedRewards(&_SocializingPool.CallOpts, arg0, arg1)
}

// ClaimedRewards is a free data retrieval call binding the contract method 0xfb831b9a.
//
// Solidity: function claimedRewards(address , uint256 ) view returns(bool)
func (_SocializingPool *SocializingPoolCallerSession) ClaimedRewards(arg0 common.Address, arg1 *big.Int) (bool, error) {
	return _SocializingPool.Contract.ClaimedRewards(&_SocializingPool.CallOpts, arg0, arg1)
}

// GetCurrentRewardsIndex is a free data retrieval call binding the contract method 0x189956a2.
//
// Solidity: function getCurrentRewardsIndex() view returns(uint256 index)
func (_SocializingPool *SocializingPoolCaller) GetCurrentRewardsIndex(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _SocializingPool.contract.Call(opts, &out, "getCurrentRewardsIndex")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetCurrentRewardsIndex is a free data retrieval call binding the contract method 0x189956a2.
//
// Solidity: function getCurrentRewardsIndex() view returns(uint256 index)
func (_SocializingPool *SocializingPoolSession) GetCurrentRewardsIndex() (*big.Int, error) {
	return _SocializingPool.Contract.GetCurrentRewardsIndex(&_SocializingPool.CallOpts)
}

// GetCurrentRewardsIndex is a free data retrieval call binding the contract method 0x189956a2.
//
// Solidity: function getCurrentRewardsIndex() view returns(uint256 index)
func (_SocializingPool *SocializingPoolCallerSession) GetCurrentRewardsIndex() (*big.Int, error) {
	return _SocializingPool.Contract.GetCurrentRewardsIndex(&_SocializingPool.CallOpts)
}

// GetRewardCycleDetails is a free data retrieval call binding the contract method 0xd2bff5ed.
//
// Solidity: function getRewardCycleDetails(uint256 _index) view returns(uint256 _startBlock, uint256 _endBlock)
func (_SocializingPool *SocializingPoolCaller) GetRewardCycleDetails(opts *bind.CallOpts, _index *big.Int) (struct {
	StartBlock *big.Int
	EndBlock   *big.Int
}, error) {
	var out []interface{}
	err := _SocializingPool.contract.Call(opts, &out, "getRewardCycleDetails", _index)

	outstruct := new(struct {
		StartBlock *big.Int
		EndBlock   *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.StartBlock = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.EndBlock = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// GetRewardCycleDetails is a free data retrieval call binding the contract method 0xd2bff5ed.
//
// Solidity: function getRewardCycleDetails(uint256 _index) view returns(uint256 _startBlock, uint256 _endBlock)
func (_SocializingPool *SocializingPoolSession) GetRewardCycleDetails(_index *big.Int) (struct {
	StartBlock *big.Int
	EndBlock   *big.Int
}, error) {
	return _SocializingPool.Contract.GetRewardCycleDetails(&_SocializingPool.CallOpts, _index)
}

// GetRewardCycleDetails is a free data retrieval call binding the contract method 0xd2bff5ed.
//
// Solidity: function getRewardCycleDetails(uint256 _index) view returns(uint256 _startBlock, uint256 _endBlock)
func (_SocializingPool *SocializingPoolCallerSession) GetRewardCycleDetails(_index *big.Int) (struct {
	StartBlock *big.Int
	EndBlock   *big.Int
}, error) {
	return _SocializingPool.Contract.GetRewardCycleDetails(&_SocializingPool.CallOpts, _index)
}

// GetRewardDetails is a free data retrieval call binding the contract method 0xd0c40276.
//
// Solidity: function getRewardDetails() view returns(uint256 currentIndex, uint256 currentStartBlock, uint256 currentEndBlock)
func (_SocializingPool *SocializingPoolCaller) GetRewardDetails(opts *bind.CallOpts) (struct {
	CurrentIndex      *big.Int
	CurrentStartBlock *big.Int
	CurrentEndBlock   *big.Int
}, error) {
	var out []interface{}
	err := _SocializingPool.contract.Call(opts, &out, "getRewardDetails")

	outstruct := new(struct {
		CurrentIndex      *big.Int
		CurrentStartBlock *big.Int
		CurrentEndBlock   *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.CurrentIndex = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.CurrentStartBlock = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.CurrentEndBlock = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// GetRewardDetails is a free data retrieval call binding the contract method 0xd0c40276.
//
// Solidity: function getRewardDetails() view returns(uint256 currentIndex, uint256 currentStartBlock, uint256 currentEndBlock)
func (_SocializingPool *SocializingPoolSession) GetRewardDetails() (struct {
	CurrentIndex      *big.Int
	CurrentStartBlock *big.Int
	CurrentEndBlock   *big.Int
}, error) {
	return _SocializingPool.Contract.GetRewardDetails(&_SocializingPool.CallOpts)
}

// GetRewardDetails is a free data retrieval call binding the contract method 0xd0c40276.
//
// Solidity: function getRewardDetails() view returns(uint256 currentIndex, uint256 currentStartBlock, uint256 currentEndBlock)
func (_SocializingPool *SocializingPoolCallerSession) GetRewardDetails() (struct {
	CurrentIndex      *big.Int
	CurrentStartBlock *big.Int
	CurrentEndBlock   *big.Int
}, error) {
	return _SocializingPool.Contract.GetRewardDetails(&_SocializingPool.CallOpts)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_SocializingPool *SocializingPoolCaller) GetRoleAdmin(opts *bind.CallOpts, role [32]byte) ([32]byte, error) {
	var out []interface{}
	err := _SocializingPool.contract.Call(opts, &out, "getRoleAdmin", role)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_SocializingPool *SocializingPoolSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _SocializingPool.Contract.GetRoleAdmin(&_SocializingPool.CallOpts, role)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_SocializingPool *SocializingPoolCallerSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _SocializingPool.Contract.GetRoleAdmin(&_SocializingPool.CallOpts, role)
}

// HandledRewards is a free data retrieval call binding the contract method 0xebc0f5f7.
//
// Solidity: function handledRewards(uint256 ) view returns(bool)
func (_SocializingPool *SocializingPoolCaller) HandledRewards(opts *bind.CallOpts, arg0 *big.Int) (bool, error) {
	var out []interface{}
	err := _SocializingPool.contract.Call(opts, &out, "handledRewards", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// HandledRewards is a free data retrieval call binding the contract method 0xebc0f5f7.
//
// Solidity: function handledRewards(uint256 ) view returns(bool)
func (_SocializingPool *SocializingPoolSession) HandledRewards(arg0 *big.Int) (bool, error) {
	return _SocializingPool.Contract.HandledRewards(&_SocializingPool.CallOpts, arg0)
}

// HandledRewards is a free data retrieval call binding the contract method 0xebc0f5f7.
//
// Solidity: function handledRewards(uint256 ) view returns(bool)
func (_SocializingPool *SocializingPoolCallerSession) HandledRewards(arg0 *big.Int) (bool, error) {
	return _SocializingPool.Contract.HandledRewards(&_SocializingPool.CallOpts, arg0)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_SocializingPool *SocializingPoolCaller) HasRole(opts *bind.CallOpts, role [32]byte, account common.Address) (bool, error) {
	var out []interface{}
	err := _SocializingPool.contract.Call(opts, &out, "hasRole", role, account)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_SocializingPool *SocializingPoolSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _SocializingPool.Contract.HasRole(&_SocializingPool.CallOpts, role, account)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_SocializingPool *SocializingPoolCallerSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _SocializingPool.Contract.HasRole(&_SocializingPool.CallOpts, role, account)
}

// InitialBlock is a free data retrieval call binding the contract method 0x2cb15864.
//
// Solidity: function initialBlock() view returns(uint256)
func (_SocializingPool *SocializingPoolCaller) InitialBlock(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _SocializingPool.contract.Call(opts, &out, "initialBlock")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// InitialBlock is a free data retrieval call binding the contract method 0x2cb15864.
//
// Solidity: function initialBlock() view returns(uint256)
func (_SocializingPool *SocializingPoolSession) InitialBlock() (*big.Int, error) {
	return _SocializingPool.Contract.InitialBlock(&_SocializingPool.CallOpts)
}

// InitialBlock is a free data retrieval call binding the contract method 0x2cb15864.
//
// Solidity: function initialBlock() view returns(uint256)
func (_SocializingPool *SocializingPoolCallerSession) InitialBlock() (*big.Int, error) {
	return _SocializingPool.Contract.InitialBlock(&_SocializingPool.CallOpts)
}

// LastReportedRewardsData is a free data retrieval call binding the contract method 0x251272e0.
//
// Solidity: function lastReportedRewardsData() view returns(uint256 reportingBlockNumber, uint256 index, bytes32 merkleRoot, uint8 poolId, uint256 operatorETHRewards, uint256 userETHRewards, uint256 protocolETHRewards, uint256 operatorSDRewards)
func (_SocializingPool *SocializingPoolCaller) LastReportedRewardsData(opts *bind.CallOpts) (struct {
	ReportingBlockNumber *big.Int
	Index                *big.Int
	MerkleRoot           [32]byte
	PoolId               uint8
	OperatorETHRewards   *big.Int
	UserETHRewards       *big.Int
	ProtocolETHRewards   *big.Int
	OperatorSDRewards    *big.Int
}, error) {
	var out []interface{}
	err := _SocializingPool.contract.Call(opts, &out, "lastReportedRewardsData")

	outstruct := new(struct {
		ReportingBlockNumber *big.Int
		Index                *big.Int
		MerkleRoot           [32]byte
		PoolId               uint8
		OperatorETHRewards   *big.Int
		UserETHRewards       *big.Int
		ProtocolETHRewards   *big.Int
		OperatorSDRewards    *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.ReportingBlockNumber = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.Index = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.MerkleRoot = *abi.ConvertType(out[2], new([32]byte)).(*[32]byte)
	outstruct.PoolId = *abi.ConvertType(out[3], new(uint8)).(*uint8)
	outstruct.OperatorETHRewards = *abi.ConvertType(out[4], new(*big.Int)).(**big.Int)
	outstruct.UserETHRewards = *abi.ConvertType(out[5], new(*big.Int)).(**big.Int)
	outstruct.ProtocolETHRewards = *abi.ConvertType(out[6], new(*big.Int)).(**big.Int)
	outstruct.OperatorSDRewards = *abi.ConvertType(out[7], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// LastReportedRewardsData is a free data retrieval call binding the contract method 0x251272e0.
//
// Solidity: function lastReportedRewardsData() view returns(uint256 reportingBlockNumber, uint256 index, bytes32 merkleRoot, uint8 poolId, uint256 operatorETHRewards, uint256 userETHRewards, uint256 protocolETHRewards, uint256 operatorSDRewards)
func (_SocializingPool *SocializingPoolSession) LastReportedRewardsData() (struct {
	ReportingBlockNumber *big.Int
	Index                *big.Int
	MerkleRoot           [32]byte
	PoolId               uint8
	OperatorETHRewards   *big.Int
	UserETHRewards       *big.Int
	ProtocolETHRewards   *big.Int
	OperatorSDRewards    *big.Int
}, error) {
	return _SocializingPool.Contract.LastReportedRewardsData(&_SocializingPool.CallOpts)
}

// LastReportedRewardsData is a free data retrieval call binding the contract method 0x251272e0.
//
// Solidity: function lastReportedRewardsData() view returns(uint256 reportingBlockNumber, uint256 index, bytes32 merkleRoot, uint8 poolId, uint256 operatorETHRewards, uint256 userETHRewards, uint256 protocolETHRewards, uint256 operatorSDRewards)
func (_SocializingPool *SocializingPoolCallerSession) LastReportedRewardsData() (struct {
	ReportingBlockNumber *big.Int
	Index                *big.Int
	MerkleRoot           [32]byte
	PoolId               uint8
	OperatorETHRewards   *big.Int
	UserETHRewards       *big.Int
	ProtocolETHRewards   *big.Int
	OperatorSDRewards    *big.Int
}, error) {
	return _SocializingPool.Contract.LastReportedRewardsData(&_SocializingPool.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_SocializingPool *SocializingPoolCaller) Paused(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _SocializingPool.contract.Call(opts, &out, "paused")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_SocializingPool *SocializingPoolSession) Paused() (bool, error) {
	return _SocializingPool.Contract.Paused(&_SocializingPool.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_SocializingPool *SocializingPoolCallerSession) Paused() (bool, error) {
	return _SocializingPool.Contract.Paused(&_SocializingPool.CallOpts)
}

// RewardsDataMap is a free data retrieval call binding the contract method 0x4a321b79.
//
// Solidity: function rewardsDataMap(uint256 ) view returns(uint256 reportingBlockNumber, uint256 index, bytes32 merkleRoot, uint8 poolId, uint256 operatorETHRewards, uint256 userETHRewards, uint256 protocolETHRewards, uint256 operatorSDRewards)
func (_SocializingPool *SocializingPoolCaller) RewardsDataMap(opts *bind.CallOpts, arg0 *big.Int) (struct {
	ReportingBlockNumber *big.Int
	Index                *big.Int
	MerkleRoot           [32]byte
	PoolId               uint8
	OperatorETHRewards   *big.Int
	UserETHRewards       *big.Int
	ProtocolETHRewards   *big.Int
	OperatorSDRewards    *big.Int
}, error) {
	var out []interface{}
	err := _SocializingPool.contract.Call(opts, &out, "rewardsDataMap", arg0)

	outstruct := new(struct {
		ReportingBlockNumber *big.Int
		Index                *big.Int
		MerkleRoot           [32]byte
		PoolId               uint8
		OperatorETHRewards   *big.Int
		UserETHRewards       *big.Int
		ProtocolETHRewards   *big.Int
		OperatorSDRewards    *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.ReportingBlockNumber = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.Index = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.MerkleRoot = *abi.ConvertType(out[2], new([32]byte)).(*[32]byte)
	outstruct.PoolId = *abi.ConvertType(out[3], new(uint8)).(*uint8)
	outstruct.OperatorETHRewards = *abi.ConvertType(out[4], new(*big.Int)).(**big.Int)
	outstruct.UserETHRewards = *abi.ConvertType(out[5], new(*big.Int)).(**big.Int)
	outstruct.ProtocolETHRewards = *abi.ConvertType(out[6], new(*big.Int)).(**big.Int)
	outstruct.OperatorSDRewards = *abi.ConvertType(out[7], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// RewardsDataMap is a free data retrieval call binding the contract method 0x4a321b79.
//
// Solidity: function rewardsDataMap(uint256 ) view returns(uint256 reportingBlockNumber, uint256 index, bytes32 merkleRoot, uint8 poolId, uint256 operatorETHRewards, uint256 userETHRewards, uint256 protocolETHRewards, uint256 operatorSDRewards)
func (_SocializingPool *SocializingPoolSession) RewardsDataMap(arg0 *big.Int) (struct {
	ReportingBlockNumber *big.Int
	Index                *big.Int
	MerkleRoot           [32]byte
	PoolId               uint8
	OperatorETHRewards   *big.Int
	UserETHRewards       *big.Int
	ProtocolETHRewards   *big.Int
	OperatorSDRewards    *big.Int
}, error) {
	return _SocializingPool.Contract.RewardsDataMap(&_SocializingPool.CallOpts, arg0)
}

// RewardsDataMap is a free data retrieval call binding the contract method 0x4a321b79.
//
// Solidity: function rewardsDataMap(uint256 ) view returns(uint256 reportingBlockNumber, uint256 index, bytes32 merkleRoot, uint8 poolId, uint256 operatorETHRewards, uint256 userETHRewards, uint256 protocolETHRewards, uint256 operatorSDRewards)
func (_SocializingPool *SocializingPoolCallerSession) RewardsDataMap(arg0 *big.Int) (struct {
	ReportingBlockNumber *big.Int
	Index                *big.Int
	MerkleRoot           [32]byte
	PoolId               uint8
	OperatorETHRewards   *big.Int
	UserETHRewards       *big.Int
	ProtocolETHRewards   *big.Int
	OperatorSDRewards    *big.Int
}, error) {
	return _SocializingPool.Contract.RewardsDataMap(&_SocializingPool.CallOpts, arg0)
}

// StaderConfig is a free data retrieval call binding the contract method 0x490ffa35.
//
// Solidity: function staderConfig() view returns(address)
func (_SocializingPool *SocializingPoolCaller) StaderConfig(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _SocializingPool.contract.Call(opts, &out, "staderConfig")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// StaderConfig is a free data retrieval call binding the contract method 0x490ffa35.
//
// Solidity: function staderConfig() view returns(address)
func (_SocializingPool *SocializingPoolSession) StaderConfig() (common.Address, error) {
	return _SocializingPool.Contract.StaderConfig(&_SocializingPool.CallOpts)
}

// StaderConfig is a free data retrieval call binding the contract method 0x490ffa35.
//
// Solidity: function staderConfig() view returns(address)
func (_SocializingPool *SocializingPoolCallerSession) StaderConfig() (common.Address, error) {
	return _SocializingPool.Contract.StaderConfig(&_SocializingPool.CallOpts)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_SocializingPool *SocializingPoolCaller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _SocializingPool.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_SocializingPool *SocializingPoolSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _SocializingPool.Contract.SupportsInterface(&_SocializingPool.CallOpts, interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_SocializingPool *SocializingPoolCallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _SocializingPool.Contract.SupportsInterface(&_SocializingPool.CallOpts, interfaceId)
}

// TotalOperatorETHRewardsRemaining is a free data retrieval call binding the contract method 0xc8725d82.
//
// Solidity: function totalOperatorETHRewardsRemaining() view returns(uint256)
func (_SocializingPool *SocializingPoolCaller) TotalOperatorETHRewardsRemaining(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _SocializingPool.contract.Call(opts, &out, "totalOperatorETHRewardsRemaining")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalOperatorETHRewardsRemaining is a free data retrieval call binding the contract method 0xc8725d82.
//
// Solidity: function totalOperatorETHRewardsRemaining() view returns(uint256)
func (_SocializingPool *SocializingPoolSession) TotalOperatorETHRewardsRemaining() (*big.Int, error) {
	return _SocializingPool.Contract.TotalOperatorETHRewardsRemaining(&_SocializingPool.CallOpts)
}

// TotalOperatorETHRewardsRemaining is a free data retrieval call binding the contract method 0xc8725d82.
//
// Solidity: function totalOperatorETHRewardsRemaining() view returns(uint256)
func (_SocializingPool *SocializingPoolCallerSession) TotalOperatorETHRewardsRemaining() (*big.Int, error) {
	return _SocializingPool.Contract.TotalOperatorETHRewardsRemaining(&_SocializingPool.CallOpts)
}

// TotalOperatorSDRewardsRemaining is a free data retrieval call binding the contract method 0x47675c9e.
//
// Solidity: function totalOperatorSDRewardsRemaining() view returns(uint256)
func (_SocializingPool *SocializingPoolCaller) TotalOperatorSDRewardsRemaining(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _SocializingPool.contract.Call(opts, &out, "totalOperatorSDRewardsRemaining")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalOperatorSDRewardsRemaining is a free data retrieval call binding the contract method 0x47675c9e.
//
// Solidity: function totalOperatorSDRewardsRemaining() view returns(uint256)
func (_SocializingPool *SocializingPoolSession) TotalOperatorSDRewardsRemaining() (*big.Int, error) {
	return _SocializingPool.Contract.TotalOperatorSDRewardsRemaining(&_SocializingPool.CallOpts)
}

// TotalOperatorSDRewardsRemaining is a free data retrieval call binding the contract method 0x47675c9e.
//
// Solidity: function totalOperatorSDRewardsRemaining() view returns(uint256)
func (_SocializingPool *SocializingPoolCallerSession) TotalOperatorSDRewardsRemaining() (*big.Int, error) {
	return _SocializingPool.Contract.TotalOperatorSDRewardsRemaining(&_SocializingPool.CallOpts)
}

// VerifyProof is a free data retrieval call binding the contract method 0xfffbe459.
//
// Solidity: function verifyProof(uint256 _index, address _operator, uint256 _amountSD, uint256 _amountETH, bytes32[] _merkleProof) view returns(bool)
func (_SocializingPool *SocializingPoolCaller) VerifyProof(opts *bind.CallOpts, _index *big.Int, _operator common.Address, _amountSD *big.Int, _amountETH *big.Int, _merkleProof [][32]byte) (bool, error) {
	var out []interface{}
	err := _SocializingPool.contract.Call(opts, &out, "verifyProof", _index, _operator, _amountSD, _amountETH, _merkleProof)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// VerifyProof is a free data retrieval call binding the contract method 0xfffbe459.
//
// Solidity: function verifyProof(uint256 _index, address _operator, uint256 _amountSD, uint256 _amountETH, bytes32[] _merkleProof) view returns(bool)
func (_SocializingPool *SocializingPoolSession) VerifyProof(_index *big.Int, _operator common.Address, _amountSD *big.Int, _amountETH *big.Int, _merkleProof [][32]byte) (bool, error) {
	return _SocializingPool.Contract.VerifyProof(&_SocializingPool.CallOpts, _index, _operator, _amountSD, _amountETH, _merkleProof)
}

// VerifyProof is a free data retrieval call binding the contract method 0xfffbe459.
//
// Solidity: function verifyProof(uint256 _index, address _operator, uint256 _amountSD, uint256 _amountETH, bytes32[] _merkleProof) view returns(bool)
func (_SocializingPool *SocializingPoolCallerSession) VerifyProof(_index *big.Int, _operator common.Address, _amountSD *big.Int, _amountETH *big.Int, _merkleProof [][32]byte) (bool, error) {
	return _SocializingPool.Contract.VerifyProof(&_SocializingPool.CallOpts, _index, _operator, _amountSD, _amountETH, _merkleProof)
}

// Claim is a paid mutator transaction binding the contract method 0xd009b3d0.
//
// Solidity: function claim(uint256[] _index, uint256[] _amountSD, uint256[] _amountETH, bytes32[][] _merkleProof) returns()
func (_SocializingPool *SocializingPoolTransactor) Claim(opts *bind.TransactOpts, _index []*big.Int, _amountSD []*big.Int, _amountETH []*big.Int, _merkleProof [][][32]byte) (*types.Transaction, error) {
	return _SocializingPool.contract.Transact(opts, "claim", _index, _amountSD, _amountETH, _merkleProof)
}

// Claim is a paid mutator transaction binding the contract method 0xd009b3d0.
//
// Solidity: function claim(uint256[] _index, uint256[] _amountSD, uint256[] _amountETH, bytes32[][] _merkleProof) returns()
func (_SocializingPool *SocializingPoolSession) Claim(_index []*big.Int, _amountSD []*big.Int, _amountETH []*big.Int, _merkleProof [][][32]byte) (*types.Transaction, error) {
	return _SocializingPool.Contract.Claim(&_SocializingPool.TransactOpts, _index, _amountSD, _amountETH, _merkleProof)
}

// Claim is a paid mutator transaction binding the contract method 0xd009b3d0.
//
// Solidity: function claim(uint256[] _index, uint256[] _amountSD, uint256[] _amountETH, bytes32[][] _merkleProof) returns()
func (_SocializingPool *SocializingPoolTransactorSession) Claim(_index []*big.Int, _amountSD []*big.Int, _amountETH []*big.Int, _merkleProof [][][32]byte) (*types.Transaction, error) {
	return _SocializingPool.Contract.Claim(&_SocializingPool.TransactOpts, _index, _amountSD, _amountETH, _merkleProof)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_SocializingPool *SocializingPoolTransactor) GrantRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _SocializingPool.contract.Transact(opts, "grantRole", role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_SocializingPool *SocializingPoolSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _SocializingPool.Contract.GrantRole(&_SocializingPool.TransactOpts, role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_SocializingPool *SocializingPoolTransactorSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _SocializingPool.Contract.GrantRole(&_SocializingPool.TransactOpts, role, account)
}

// HandleRewards is a paid mutator transaction binding the contract method 0x0d83e4ed.
//
// Solidity: function handleRewards((uint256,uint256,bytes32,uint8,uint256,uint256,uint256,uint256) _rewardsData) returns()
func (_SocializingPool *SocializingPoolTransactor) HandleRewards(opts *bind.TransactOpts, _rewardsData RewardsData) (*types.Transaction, error) {
	return _SocializingPool.contract.Transact(opts, "handleRewards", _rewardsData)
}

// HandleRewards is a paid mutator transaction binding the contract method 0x0d83e4ed.
//
// Solidity: function handleRewards((uint256,uint256,bytes32,uint8,uint256,uint256,uint256,uint256) _rewardsData) returns()
func (_SocializingPool *SocializingPoolSession) HandleRewards(_rewardsData RewardsData) (*types.Transaction, error) {
	return _SocializingPool.Contract.HandleRewards(&_SocializingPool.TransactOpts, _rewardsData)
}

// HandleRewards is a paid mutator transaction binding the contract method 0x0d83e4ed.
//
// Solidity: function handleRewards((uint256,uint256,bytes32,uint8,uint256,uint256,uint256,uint256) _rewardsData) returns()
func (_SocializingPool *SocializingPoolTransactorSession) HandleRewards(_rewardsData RewardsData) (*types.Transaction, error) {
	return _SocializingPool.Contract.HandleRewards(&_SocializingPool.TransactOpts, _rewardsData)
}

// Initialize is a paid mutator transaction binding the contract method 0x485cc955.
//
// Solidity: function initialize(address _admin, address _staderConfig) returns()
func (_SocializingPool *SocializingPoolTransactor) Initialize(opts *bind.TransactOpts, _admin common.Address, _staderConfig common.Address) (*types.Transaction, error) {
	return _SocializingPool.contract.Transact(opts, "initialize", _admin, _staderConfig)
}

// Initialize is a paid mutator transaction binding the contract method 0x485cc955.
//
// Solidity: function initialize(address _admin, address _staderConfig) returns()
func (_SocializingPool *SocializingPoolSession) Initialize(_admin common.Address, _staderConfig common.Address) (*types.Transaction, error) {
	return _SocializingPool.Contract.Initialize(&_SocializingPool.TransactOpts, _admin, _staderConfig)
}

// Initialize is a paid mutator transaction binding the contract method 0x485cc955.
//
// Solidity: function initialize(address _admin, address _staderConfig) returns()
func (_SocializingPool *SocializingPoolTransactorSession) Initialize(_admin common.Address, _staderConfig common.Address) (*types.Transaction, error) {
	return _SocializingPool.Contract.Initialize(&_SocializingPool.TransactOpts, _admin, _staderConfig)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_SocializingPool *SocializingPoolTransactor) Pause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SocializingPool.contract.Transact(opts, "pause")
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_SocializingPool *SocializingPoolSession) Pause() (*types.Transaction, error) {
	return _SocializingPool.Contract.Pause(&_SocializingPool.TransactOpts)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_SocializingPool *SocializingPoolTransactorSession) Pause() (*types.Transaction, error) {
	return _SocializingPool.Contract.Pause(&_SocializingPool.TransactOpts)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address account) returns()
func (_SocializingPool *SocializingPoolTransactor) RenounceRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _SocializingPool.contract.Transact(opts, "renounceRole", role, account)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address account) returns()
func (_SocializingPool *SocializingPoolSession) RenounceRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _SocializingPool.Contract.RenounceRole(&_SocializingPool.TransactOpts, role, account)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address account) returns()
func (_SocializingPool *SocializingPoolTransactorSession) RenounceRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _SocializingPool.Contract.RenounceRole(&_SocializingPool.TransactOpts, role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_SocializingPool *SocializingPoolTransactor) RevokeRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _SocializingPool.contract.Transact(opts, "revokeRole", role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_SocializingPool *SocializingPoolSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _SocializingPool.Contract.RevokeRole(&_SocializingPool.TransactOpts, role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_SocializingPool *SocializingPoolTransactorSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _SocializingPool.Contract.RevokeRole(&_SocializingPool.TransactOpts, role, account)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_SocializingPool *SocializingPoolTransactor) Unpause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SocializingPool.contract.Transact(opts, "unpause")
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_SocializingPool *SocializingPoolSession) Unpause() (*types.Transaction, error) {
	return _SocializingPool.Contract.Unpause(&_SocializingPool.TransactOpts)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_SocializingPool *SocializingPoolTransactorSession) Unpause() (*types.Transaction, error) {
	return _SocializingPool.Contract.Unpause(&_SocializingPool.TransactOpts)
}

// UpdateStaderConfig is a paid mutator transaction binding the contract method 0x9ee804cb.
//
// Solidity: function updateStaderConfig(address _staderConfig) returns()
func (_SocializingPool *SocializingPoolTransactor) UpdateStaderConfig(opts *bind.TransactOpts, _staderConfig common.Address) (*types.Transaction, error) {
	return _SocializingPool.contract.Transact(opts, "updateStaderConfig", _staderConfig)
}

// UpdateStaderConfig is a paid mutator transaction binding the contract method 0x9ee804cb.
//
// Solidity: function updateStaderConfig(address _staderConfig) returns()
func (_SocializingPool *SocializingPoolSession) UpdateStaderConfig(_staderConfig common.Address) (*types.Transaction, error) {
	return _SocializingPool.Contract.UpdateStaderConfig(&_SocializingPool.TransactOpts, _staderConfig)
}

// UpdateStaderConfig is a paid mutator transaction binding the contract method 0x9ee804cb.
//
// Solidity: function updateStaderConfig(address _staderConfig) returns()
func (_SocializingPool *SocializingPoolTransactorSession) UpdateStaderConfig(_staderConfig common.Address) (*types.Transaction, error) {
	return _SocializingPool.Contract.UpdateStaderConfig(&_SocializingPool.TransactOpts, _staderConfig)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_SocializingPool *SocializingPoolTransactor) Receive(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SocializingPool.contract.RawTransact(opts, nil) // calldata is disallowed for receive function
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_SocializingPool *SocializingPoolSession) Receive() (*types.Transaction, error) {
	return _SocializingPool.Contract.Receive(&_SocializingPool.TransactOpts)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_SocializingPool *SocializingPoolTransactorSession) Receive() (*types.Transaction, error) {
	return _SocializingPool.Contract.Receive(&_SocializingPool.TransactOpts)
}

// SocializingPoolETHReceivedIterator is returned from FilterETHReceived and is used to iterate over the raw logs and unpacked data for ETHReceived events raised by the SocializingPool contract.
type SocializingPoolETHReceivedIterator struct {
	Event *SocializingPoolETHReceived // Event containing the contract specifics and raw log

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
func (it *SocializingPoolETHReceivedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SocializingPoolETHReceived)
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
		it.Event = new(SocializingPoolETHReceived)
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
func (it *SocializingPoolETHReceivedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SocializingPoolETHReceivedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SocializingPoolETHReceived represents a ETHReceived event raised by the SocializingPool contract.
type SocializingPoolETHReceived struct {
	Sender common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterETHReceived is a free log retrieval operation binding the contract event 0xbfe611b001dfcd411432f7bf0d79b82b4b2ee81511edac123a3403c357fb972a.
//
// Solidity: event ETHReceived(address indexed sender, uint256 amount)
func (_SocializingPool *SocializingPoolFilterer) FilterETHReceived(opts *bind.FilterOpts, sender []common.Address) (*SocializingPoolETHReceivedIterator, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _SocializingPool.contract.FilterLogs(opts, "ETHReceived", senderRule)
	if err != nil {
		return nil, err
	}
	return &SocializingPoolETHReceivedIterator{contract: _SocializingPool.contract, event: "ETHReceived", logs: logs, sub: sub}, nil
}

// WatchETHReceived is a free log subscription operation binding the contract event 0xbfe611b001dfcd411432f7bf0d79b82b4b2ee81511edac123a3403c357fb972a.
//
// Solidity: event ETHReceived(address indexed sender, uint256 amount)
func (_SocializingPool *SocializingPoolFilterer) WatchETHReceived(opts *bind.WatchOpts, sink chan<- *SocializingPoolETHReceived, sender []common.Address) (event.Subscription, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _SocializingPool.contract.WatchLogs(opts, "ETHReceived", senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SocializingPoolETHReceived)
				if err := _SocializingPool.contract.UnpackLog(event, "ETHReceived", log); err != nil {
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
func (_SocializingPool *SocializingPoolFilterer) ParseETHReceived(log types.Log) (*SocializingPoolETHReceived, error) {
	event := new(SocializingPoolETHReceived)
	if err := _SocializingPool.contract.UnpackLog(event, "ETHReceived", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SocializingPoolInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the SocializingPool contract.
type SocializingPoolInitializedIterator struct {
	Event *SocializingPoolInitialized // Event containing the contract specifics and raw log

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
func (it *SocializingPoolInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SocializingPoolInitialized)
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
		it.Event = new(SocializingPoolInitialized)
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
func (it *SocializingPoolInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SocializingPoolInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SocializingPoolInitialized represents a Initialized event raised by the SocializingPool contract.
type SocializingPoolInitialized struct {
	Version uint8
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_SocializingPool *SocializingPoolFilterer) FilterInitialized(opts *bind.FilterOpts) (*SocializingPoolInitializedIterator, error) {

	logs, sub, err := _SocializingPool.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &SocializingPoolInitializedIterator{contract: _SocializingPool.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_SocializingPool *SocializingPoolFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *SocializingPoolInitialized) (event.Subscription, error) {

	logs, sub, err := _SocializingPool.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SocializingPoolInitialized)
				if err := _SocializingPool.contract.UnpackLog(event, "Initialized", log); err != nil {
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
func (_SocializingPool *SocializingPoolFilterer) ParseInitialized(log types.Log) (*SocializingPoolInitialized, error) {
	event := new(SocializingPoolInitialized)
	if err := _SocializingPool.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SocializingPoolOperatorRewardsClaimedIterator is returned from FilterOperatorRewardsClaimed and is used to iterate over the raw logs and unpacked data for OperatorRewardsClaimed events raised by the SocializingPool contract.
type SocializingPoolOperatorRewardsClaimedIterator struct {
	Event *SocializingPoolOperatorRewardsClaimed // Event containing the contract specifics and raw log

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
func (it *SocializingPoolOperatorRewardsClaimedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SocializingPoolOperatorRewardsClaimed)
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
		it.Event = new(SocializingPoolOperatorRewardsClaimed)
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
func (it *SocializingPoolOperatorRewardsClaimedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SocializingPoolOperatorRewardsClaimedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SocializingPoolOperatorRewardsClaimed represents a OperatorRewardsClaimed event raised by the SocializingPool contract.
type SocializingPoolOperatorRewardsClaimed struct {
	Recipient  common.Address
	EthRewards *big.Int
	SdRewards  *big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterOperatorRewardsClaimed is a free log retrieval operation binding the contract event 0x62bc6d6d870f047ea4dd686d08bfda93e24cac6b8dae0d740f6fa33071f3f0af.
//
// Solidity: event OperatorRewardsClaimed(address indexed recipient, uint256 ethRewards, uint256 sdRewards)
func (_SocializingPool *SocializingPoolFilterer) FilterOperatorRewardsClaimed(opts *bind.FilterOpts, recipient []common.Address) (*SocializingPoolOperatorRewardsClaimedIterator, error) {

	var recipientRule []interface{}
	for _, recipientItem := range recipient {
		recipientRule = append(recipientRule, recipientItem)
	}

	logs, sub, err := _SocializingPool.contract.FilterLogs(opts, "OperatorRewardsClaimed", recipientRule)
	if err != nil {
		return nil, err
	}
	return &SocializingPoolOperatorRewardsClaimedIterator{contract: _SocializingPool.contract, event: "OperatorRewardsClaimed", logs: logs, sub: sub}, nil
}

// WatchOperatorRewardsClaimed is a free log subscription operation binding the contract event 0x62bc6d6d870f047ea4dd686d08bfda93e24cac6b8dae0d740f6fa33071f3f0af.
//
// Solidity: event OperatorRewardsClaimed(address indexed recipient, uint256 ethRewards, uint256 sdRewards)
func (_SocializingPool *SocializingPoolFilterer) WatchOperatorRewardsClaimed(opts *bind.WatchOpts, sink chan<- *SocializingPoolOperatorRewardsClaimed, recipient []common.Address) (event.Subscription, error) {

	var recipientRule []interface{}
	for _, recipientItem := range recipient {
		recipientRule = append(recipientRule, recipientItem)
	}

	logs, sub, err := _SocializingPool.contract.WatchLogs(opts, "OperatorRewardsClaimed", recipientRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SocializingPoolOperatorRewardsClaimed)
				if err := _SocializingPool.contract.UnpackLog(event, "OperatorRewardsClaimed", log); err != nil {
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

// ParseOperatorRewardsClaimed is a log parse operation binding the contract event 0x62bc6d6d870f047ea4dd686d08bfda93e24cac6b8dae0d740f6fa33071f3f0af.
//
// Solidity: event OperatorRewardsClaimed(address indexed recipient, uint256 ethRewards, uint256 sdRewards)
func (_SocializingPool *SocializingPoolFilterer) ParseOperatorRewardsClaimed(log types.Log) (*SocializingPoolOperatorRewardsClaimed, error) {
	event := new(SocializingPoolOperatorRewardsClaimed)
	if err := _SocializingPool.contract.UnpackLog(event, "OperatorRewardsClaimed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SocializingPoolOperatorRewardsUpdatedIterator is returned from FilterOperatorRewardsUpdated and is used to iterate over the raw logs and unpacked data for OperatorRewardsUpdated events raised by the SocializingPool contract.
type SocializingPoolOperatorRewardsUpdatedIterator struct {
	Event *SocializingPoolOperatorRewardsUpdated // Event containing the contract specifics and raw log

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
func (it *SocializingPoolOperatorRewardsUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SocializingPoolOperatorRewardsUpdated)
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
		it.Event = new(SocializingPoolOperatorRewardsUpdated)
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
func (it *SocializingPoolOperatorRewardsUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SocializingPoolOperatorRewardsUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SocializingPoolOperatorRewardsUpdated represents a OperatorRewardsUpdated event raised by the SocializingPool contract.
type SocializingPoolOperatorRewardsUpdated struct {
	EthRewards      *big.Int
	TotalETHRewards *big.Int
	SdRewards       *big.Int
	TotalSDRewards  *big.Int
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterOperatorRewardsUpdated is a free log retrieval operation binding the contract event 0x0e357ad2594fa2d9d8c6dc7c280141cc1b89ba4c9714a96fd3f409f4fded31d0.
//
// Solidity: event OperatorRewardsUpdated(uint256 ethRewards, uint256 totalETHRewards, uint256 sdRewards, uint256 totalSDRewards)
func (_SocializingPool *SocializingPoolFilterer) FilterOperatorRewardsUpdated(opts *bind.FilterOpts) (*SocializingPoolOperatorRewardsUpdatedIterator, error) {

	logs, sub, err := _SocializingPool.contract.FilterLogs(opts, "OperatorRewardsUpdated")
	if err != nil {
		return nil, err
	}
	return &SocializingPoolOperatorRewardsUpdatedIterator{contract: _SocializingPool.contract, event: "OperatorRewardsUpdated", logs: logs, sub: sub}, nil
}

// WatchOperatorRewardsUpdated is a free log subscription operation binding the contract event 0x0e357ad2594fa2d9d8c6dc7c280141cc1b89ba4c9714a96fd3f409f4fded31d0.
//
// Solidity: event OperatorRewardsUpdated(uint256 ethRewards, uint256 totalETHRewards, uint256 sdRewards, uint256 totalSDRewards)
func (_SocializingPool *SocializingPoolFilterer) WatchOperatorRewardsUpdated(opts *bind.WatchOpts, sink chan<- *SocializingPoolOperatorRewardsUpdated) (event.Subscription, error) {

	logs, sub, err := _SocializingPool.contract.WatchLogs(opts, "OperatorRewardsUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SocializingPoolOperatorRewardsUpdated)
				if err := _SocializingPool.contract.UnpackLog(event, "OperatorRewardsUpdated", log); err != nil {
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

// ParseOperatorRewardsUpdated is a log parse operation binding the contract event 0x0e357ad2594fa2d9d8c6dc7c280141cc1b89ba4c9714a96fd3f409f4fded31d0.
//
// Solidity: event OperatorRewardsUpdated(uint256 ethRewards, uint256 totalETHRewards, uint256 sdRewards, uint256 totalSDRewards)
func (_SocializingPool *SocializingPoolFilterer) ParseOperatorRewardsUpdated(log types.Log) (*SocializingPoolOperatorRewardsUpdated, error) {
	event := new(SocializingPoolOperatorRewardsUpdated)
	if err := _SocializingPool.contract.UnpackLog(event, "OperatorRewardsUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SocializingPoolPausedIterator is returned from FilterPaused and is used to iterate over the raw logs and unpacked data for Paused events raised by the SocializingPool contract.
type SocializingPoolPausedIterator struct {
	Event *SocializingPoolPaused // Event containing the contract specifics and raw log

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
func (it *SocializingPoolPausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SocializingPoolPaused)
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
		it.Event = new(SocializingPoolPaused)
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
func (it *SocializingPoolPausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SocializingPoolPausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SocializingPoolPaused represents a Paused event raised by the SocializingPool contract.
type SocializingPoolPaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterPaused is a free log retrieval operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_SocializingPool *SocializingPoolFilterer) FilterPaused(opts *bind.FilterOpts) (*SocializingPoolPausedIterator, error) {

	logs, sub, err := _SocializingPool.contract.FilterLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return &SocializingPoolPausedIterator{contract: _SocializingPool.contract, event: "Paused", logs: logs, sub: sub}, nil
}

// WatchPaused is a free log subscription operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_SocializingPool *SocializingPoolFilterer) WatchPaused(opts *bind.WatchOpts, sink chan<- *SocializingPoolPaused) (event.Subscription, error) {

	logs, sub, err := _SocializingPool.contract.WatchLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SocializingPoolPaused)
				if err := _SocializingPool.contract.UnpackLog(event, "Paused", log); err != nil {
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

// ParsePaused is a log parse operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_SocializingPool *SocializingPoolFilterer) ParsePaused(log types.Log) (*SocializingPoolPaused, error) {
	event := new(SocializingPoolPaused)
	if err := _SocializingPool.contract.UnpackLog(event, "Paused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SocializingPoolProtocolETHRewardsTransferredIterator is returned from FilterProtocolETHRewardsTransferred and is used to iterate over the raw logs and unpacked data for ProtocolETHRewardsTransferred events raised by the SocializingPool contract.
type SocializingPoolProtocolETHRewardsTransferredIterator struct {
	Event *SocializingPoolProtocolETHRewardsTransferred // Event containing the contract specifics and raw log

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
func (it *SocializingPoolProtocolETHRewardsTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SocializingPoolProtocolETHRewardsTransferred)
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
		it.Event = new(SocializingPoolProtocolETHRewardsTransferred)
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
func (it *SocializingPoolProtocolETHRewardsTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SocializingPoolProtocolETHRewardsTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SocializingPoolProtocolETHRewardsTransferred represents a ProtocolETHRewardsTransferred event raised by the SocializingPool contract.
type SocializingPoolProtocolETHRewardsTransferred struct {
	EthRewards *big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterProtocolETHRewardsTransferred is a free log retrieval operation binding the contract event 0x292921846cb4d7dd20ae1c60a15192efacaaa30028f2043998f925c6c10a8150.
//
// Solidity: event ProtocolETHRewardsTransferred(uint256 ethRewards)
func (_SocializingPool *SocializingPoolFilterer) FilterProtocolETHRewardsTransferred(opts *bind.FilterOpts) (*SocializingPoolProtocolETHRewardsTransferredIterator, error) {

	logs, sub, err := _SocializingPool.contract.FilterLogs(opts, "ProtocolETHRewardsTransferred")
	if err != nil {
		return nil, err
	}
	return &SocializingPoolProtocolETHRewardsTransferredIterator{contract: _SocializingPool.contract, event: "ProtocolETHRewardsTransferred", logs: logs, sub: sub}, nil
}

// WatchProtocolETHRewardsTransferred is a free log subscription operation binding the contract event 0x292921846cb4d7dd20ae1c60a15192efacaaa30028f2043998f925c6c10a8150.
//
// Solidity: event ProtocolETHRewardsTransferred(uint256 ethRewards)
func (_SocializingPool *SocializingPoolFilterer) WatchProtocolETHRewardsTransferred(opts *bind.WatchOpts, sink chan<- *SocializingPoolProtocolETHRewardsTransferred) (event.Subscription, error) {

	logs, sub, err := _SocializingPool.contract.WatchLogs(opts, "ProtocolETHRewardsTransferred")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SocializingPoolProtocolETHRewardsTransferred)
				if err := _SocializingPool.contract.UnpackLog(event, "ProtocolETHRewardsTransferred", log); err != nil {
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

// ParseProtocolETHRewardsTransferred is a log parse operation binding the contract event 0x292921846cb4d7dd20ae1c60a15192efacaaa30028f2043998f925c6c10a8150.
//
// Solidity: event ProtocolETHRewardsTransferred(uint256 ethRewards)
func (_SocializingPool *SocializingPoolFilterer) ParseProtocolETHRewardsTransferred(log types.Log) (*SocializingPoolProtocolETHRewardsTransferred, error) {
	event := new(SocializingPoolProtocolETHRewardsTransferred)
	if err := _SocializingPool.contract.UnpackLog(event, "ProtocolETHRewardsTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SocializingPoolRoleAdminChangedIterator is returned from FilterRoleAdminChanged and is used to iterate over the raw logs and unpacked data for RoleAdminChanged events raised by the SocializingPool contract.
type SocializingPoolRoleAdminChangedIterator struct {
	Event *SocializingPoolRoleAdminChanged // Event containing the contract specifics and raw log

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
func (it *SocializingPoolRoleAdminChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SocializingPoolRoleAdminChanged)
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
		it.Event = new(SocializingPoolRoleAdminChanged)
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
func (it *SocializingPoolRoleAdminChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SocializingPoolRoleAdminChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SocializingPoolRoleAdminChanged represents a RoleAdminChanged event raised by the SocializingPool contract.
type SocializingPoolRoleAdminChanged struct {
	Role              [32]byte
	PreviousAdminRole [32]byte
	NewAdminRole      [32]byte
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterRoleAdminChanged is a free log retrieval operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_SocializingPool *SocializingPoolFilterer) FilterRoleAdminChanged(opts *bind.FilterOpts, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (*SocializingPoolRoleAdminChangedIterator, error) {

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

	logs, sub, err := _SocializingPool.contract.FilterLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return &SocializingPoolRoleAdminChangedIterator{contract: _SocializingPool.contract, event: "RoleAdminChanged", logs: logs, sub: sub}, nil
}

// WatchRoleAdminChanged is a free log subscription operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_SocializingPool *SocializingPoolFilterer) WatchRoleAdminChanged(opts *bind.WatchOpts, sink chan<- *SocializingPoolRoleAdminChanged, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (event.Subscription, error) {

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

	logs, sub, err := _SocializingPool.contract.WatchLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SocializingPoolRoleAdminChanged)
				if err := _SocializingPool.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
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
func (_SocializingPool *SocializingPoolFilterer) ParseRoleAdminChanged(log types.Log) (*SocializingPoolRoleAdminChanged, error) {
	event := new(SocializingPoolRoleAdminChanged)
	if err := _SocializingPool.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SocializingPoolRoleGrantedIterator is returned from FilterRoleGranted and is used to iterate over the raw logs and unpacked data for RoleGranted events raised by the SocializingPool contract.
type SocializingPoolRoleGrantedIterator struct {
	Event *SocializingPoolRoleGranted // Event containing the contract specifics and raw log

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
func (it *SocializingPoolRoleGrantedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SocializingPoolRoleGranted)
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
		it.Event = new(SocializingPoolRoleGranted)
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
func (it *SocializingPoolRoleGrantedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SocializingPoolRoleGrantedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SocializingPoolRoleGranted represents a RoleGranted event raised by the SocializingPool contract.
type SocializingPoolRoleGranted struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleGranted is a free log retrieval operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_SocializingPool *SocializingPoolFilterer) FilterRoleGranted(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*SocializingPoolRoleGrantedIterator, error) {

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

	logs, sub, err := _SocializingPool.contract.FilterLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &SocializingPoolRoleGrantedIterator{contract: _SocializingPool.contract, event: "RoleGranted", logs: logs, sub: sub}, nil
}

// WatchRoleGranted is a free log subscription operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_SocializingPool *SocializingPoolFilterer) WatchRoleGranted(opts *bind.WatchOpts, sink chan<- *SocializingPoolRoleGranted, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _SocializingPool.contract.WatchLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SocializingPoolRoleGranted)
				if err := _SocializingPool.contract.UnpackLog(event, "RoleGranted", log); err != nil {
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
func (_SocializingPool *SocializingPoolFilterer) ParseRoleGranted(log types.Log) (*SocializingPoolRoleGranted, error) {
	event := new(SocializingPoolRoleGranted)
	if err := _SocializingPool.contract.UnpackLog(event, "RoleGranted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SocializingPoolRoleRevokedIterator is returned from FilterRoleRevoked and is used to iterate over the raw logs and unpacked data for RoleRevoked events raised by the SocializingPool contract.
type SocializingPoolRoleRevokedIterator struct {
	Event *SocializingPoolRoleRevoked // Event containing the contract specifics and raw log

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
func (it *SocializingPoolRoleRevokedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SocializingPoolRoleRevoked)
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
		it.Event = new(SocializingPoolRoleRevoked)
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
func (it *SocializingPoolRoleRevokedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SocializingPoolRoleRevokedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SocializingPoolRoleRevoked represents a RoleRevoked event raised by the SocializingPool contract.
type SocializingPoolRoleRevoked struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleRevoked is a free log retrieval operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_SocializingPool *SocializingPoolFilterer) FilterRoleRevoked(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*SocializingPoolRoleRevokedIterator, error) {

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

	logs, sub, err := _SocializingPool.contract.FilterLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &SocializingPoolRoleRevokedIterator{contract: _SocializingPool.contract, event: "RoleRevoked", logs: logs, sub: sub}, nil
}

// WatchRoleRevoked is a free log subscription operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_SocializingPool *SocializingPoolFilterer) WatchRoleRevoked(opts *bind.WatchOpts, sink chan<- *SocializingPoolRoleRevoked, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _SocializingPool.contract.WatchLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SocializingPoolRoleRevoked)
				if err := _SocializingPool.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
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
func (_SocializingPool *SocializingPoolFilterer) ParseRoleRevoked(log types.Log) (*SocializingPoolRoleRevoked, error) {
	event := new(SocializingPoolRoleRevoked)
	if err := _SocializingPool.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SocializingPoolUnpausedIterator is returned from FilterUnpaused and is used to iterate over the raw logs and unpacked data for Unpaused events raised by the SocializingPool contract.
type SocializingPoolUnpausedIterator struct {
	Event *SocializingPoolUnpaused // Event containing the contract specifics and raw log

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
func (it *SocializingPoolUnpausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SocializingPoolUnpaused)
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
		it.Event = new(SocializingPoolUnpaused)
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
func (it *SocializingPoolUnpausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SocializingPoolUnpausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SocializingPoolUnpaused represents a Unpaused event raised by the SocializingPool contract.
type SocializingPoolUnpaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterUnpaused is a free log retrieval operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_SocializingPool *SocializingPoolFilterer) FilterUnpaused(opts *bind.FilterOpts) (*SocializingPoolUnpausedIterator, error) {

	logs, sub, err := _SocializingPool.contract.FilterLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return &SocializingPoolUnpausedIterator{contract: _SocializingPool.contract, event: "Unpaused", logs: logs, sub: sub}, nil
}

// WatchUnpaused is a free log subscription operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_SocializingPool *SocializingPoolFilterer) WatchUnpaused(opts *bind.WatchOpts, sink chan<- *SocializingPoolUnpaused) (event.Subscription, error) {

	logs, sub, err := _SocializingPool.contract.WatchLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SocializingPoolUnpaused)
				if err := _SocializingPool.contract.UnpackLog(event, "Unpaused", log); err != nil {
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

// ParseUnpaused is a log parse operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_SocializingPool *SocializingPoolFilterer) ParseUnpaused(log types.Log) (*SocializingPoolUnpaused, error) {
	event := new(SocializingPoolUnpaused)
	if err := _SocializingPool.contract.UnpackLog(event, "Unpaused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SocializingPoolUpdatedStaderConfigIterator is returned from FilterUpdatedStaderConfig and is used to iterate over the raw logs and unpacked data for UpdatedStaderConfig events raised by the SocializingPool contract.
type SocializingPoolUpdatedStaderConfigIterator struct {
	Event *SocializingPoolUpdatedStaderConfig // Event containing the contract specifics and raw log

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
func (it *SocializingPoolUpdatedStaderConfigIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SocializingPoolUpdatedStaderConfig)
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
		it.Event = new(SocializingPoolUpdatedStaderConfig)
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
func (it *SocializingPoolUpdatedStaderConfigIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SocializingPoolUpdatedStaderConfigIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SocializingPoolUpdatedStaderConfig represents a UpdatedStaderConfig event raised by the SocializingPool contract.
type SocializingPoolUpdatedStaderConfig struct {
	StaderConfig common.Address
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterUpdatedStaderConfig is a free log retrieval operation binding the contract event 0xdb2219043d7b197cb235f1af0cf6d782d77dee3de19e3f4fb6d39aae633b4485.
//
// Solidity: event UpdatedStaderConfig(address indexed staderConfig)
func (_SocializingPool *SocializingPoolFilterer) FilterUpdatedStaderConfig(opts *bind.FilterOpts, staderConfig []common.Address) (*SocializingPoolUpdatedStaderConfigIterator, error) {

	var staderConfigRule []interface{}
	for _, staderConfigItem := range staderConfig {
		staderConfigRule = append(staderConfigRule, staderConfigItem)
	}

	logs, sub, err := _SocializingPool.contract.FilterLogs(opts, "UpdatedStaderConfig", staderConfigRule)
	if err != nil {
		return nil, err
	}
	return &SocializingPoolUpdatedStaderConfigIterator{contract: _SocializingPool.contract, event: "UpdatedStaderConfig", logs: logs, sub: sub}, nil
}

// WatchUpdatedStaderConfig is a free log subscription operation binding the contract event 0xdb2219043d7b197cb235f1af0cf6d782d77dee3de19e3f4fb6d39aae633b4485.
//
// Solidity: event UpdatedStaderConfig(address indexed staderConfig)
func (_SocializingPool *SocializingPoolFilterer) WatchUpdatedStaderConfig(opts *bind.WatchOpts, sink chan<- *SocializingPoolUpdatedStaderConfig, staderConfig []common.Address) (event.Subscription, error) {

	var staderConfigRule []interface{}
	for _, staderConfigItem := range staderConfig {
		staderConfigRule = append(staderConfigRule, staderConfigItem)
	}

	logs, sub, err := _SocializingPool.contract.WatchLogs(opts, "UpdatedStaderConfig", staderConfigRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SocializingPoolUpdatedStaderConfig)
				if err := _SocializingPool.contract.UnpackLog(event, "UpdatedStaderConfig", log); err != nil {
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
func (_SocializingPool *SocializingPoolFilterer) ParseUpdatedStaderConfig(log types.Log) (*SocializingPoolUpdatedStaderConfig, error) {
	event := new(SocializingPoolUpdatedStaderConfig)
	if err := _SocializingPool.contract.UnpackLog(event, "UpdatedStaderConfig", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SocializingPoolUpdatedStaderOperatorRegistryIterator is returned from FilterUpdatedStaderOperatorRegistry and is used to iterate over the raw logs and unpacked data for UpdatedStaderOperatorRegistry events raised by the SocializingPool contract.
type SocializingPoolUpdatedStaderOperatorRegistryIterator struct {
	Event *SocializingPoolUpdatedStaderOperatorRegistry // Event containing the contract specifics and raw log

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
func (it *SocializingPoolUpdatedStaderOperatorRegistryIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SocializingPoolUpdatedStaderOperatorRegistry)
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
		it.Event = new(SocializingPoolUpdatedStaderOperatorRegistry)
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
func (it *SocializingPoolUpdatedStaderOperatorRegistryIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SocializingPoolUpdatedStaderOperatorRegistryIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SocializingPoolUpdatedStaderOperatorRegistry represents a UpdatedStaderOperatorRegistry event raised by the SocializingPool contract.
type SocializingPoolUpdatedStaderOperatorRegistry struct {
	StaderOperatorRegistry common.Address
	Raw                    types.Log // Blockchain specific contextual infos
}

// FilterUpdatedStaderOperatorRegistry is a free log retrieval operation binding the contract event 0x9ea3d4ab5ce0102a316617bb8bbf02dbb10d19c7f7fd9903efd2e136658ebefb.
//
// Solidity: event UpdatedStaderOperatorRegistry(address indexed staderOperatorRegistry)
func (_SocializingPool *SocializingPoolFilterer) FilterUpdatedStaderOperatorRegistry(opts *bind.FilterOpts, staderOperatorRegistry []common.Address) (*SocializingPoolUpdatedStaderOperatorRegistryIterator, error) {

	var staderOperatorRegistryRule []interface{}
	for _, staderOperatorRegistryItem := range staderOperatorRegistry {
		staderOperatorRegistryRule = append(staderOperatorRegistryRule, staderOperatorRegistryItem)
	}

	logs, sub, err := _SocializingPool.contract.FilterLogs(opts, "UpdatedStaderOperatorRegistry", staderOperatorRegistryRule)
	if err != nil {
		return nil, err
	}
	return &SocializingPoolUpdatedStaderOperatorRegistryIterator{contract: _SocializingPool.contract, event: "UpdatedStaderOperatorRegistry", logs: logs, sub: sub}, nil
}

// WatchUpdatedStaderOperatorRegistry is a free log subscription operation binding the contract event 0x9ea3d4ab5ce0102a316617bb8bbf02dbb10d19c7f7fd9903efd2e136658ebefb.
//
// Solidity: event UpdatedStaderOperatorRegistry(address indexed staderOperatorRegistry)
func (_SocializingPool *SocializingPoolFilterer) WatchUpdatedStaderOperatorRegistry(opts *bind.WatchOpts, sink chan<- *SocializingPoolUpdatedStaderOperatorRegistry, staderOperatorRegistry []common.Address) (event.Subscription, error) {

	var staderOperatorRegistryRule []interface{}
	for _, staderOperatorRegistryItem := range staderOperatorRegistry {
		staderOperatorRegistryRule = append(staderOperatorRegistryRule, staderOperatorRegistryItem)
	}

	logs, sub, err := _SocializingPool.contract.WatchLogs(opts, "UpdatedStaderOperatorRegistry", staderOperatorRegistryRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SocializingPoolUpdatedStaderOperatorRegistry)
				if err := _SocializingPool.contract.UnpackLog(event, "UpdatedStaderOperatorRegistry", log); err != nil {
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

// ParseUpdatedStaderOperatorRegistry is a log parse operation binding the contract event 0x9ea3d4ab5ce0102a316617bb8bbf02dbb10d19c7f7fd9903efd2e136658ebefb.
//
// Solidity: event UpdatedStaderOperatorRegistry(address indexed staderOperatorRegistry)
func (_SocializingPool *SocializingPoolFilterer) ParseUpdatedStaderOperatorRegistry(log types.Log) (*SocializingPoolUpdatedStaderOperatorRegistry, error) {
	event := new(SocializingPoolUpdatedStaderOperatorRegistry)
	if err := _SocializingPool.contract.UnpackLog(event, "UpdatedStaderOperatorRegistry", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SocializingPoolUpdatedStaderValidatorRegistryIterator is returned from FilterUpdatedStaderValidatorRegistry and is used to iterate over the raw logs and unpacked data for UpdatedStaderValidatorRegistry events raised by the SocializingPool contract.
type SocializingPoolUpdatedStaderValidatorRegistryIterator struct {
	Event *SocializingPoolUpdatedStaderValidatorRegistry // Event containing the contract specifics and raw log

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
func (it *SocializingPoolUpdatedStaderValidatorRegistryIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SocializingPoolUpdatedStaderValidatorRegistry)
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
		it.Event = new(SocializingPoolUpdatedStaderValidatorRegistry)
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
func (it *SocializingPoolUpdatedStaderValidatorRegistryIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SocializingPoolUpdatedStaderValidatorRegistryIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SocializingPoolUpdatedStaderValidatorRegistry represents a UpdatedStaderValidatorRegistry event raised by the SocializingPool contract.
type SocializingPoolUpdatedStaderValidatorRegistry struct {
	StaderValidatorRegistry common.Address
	Raw                     types.Log // Blockchain specific contextual infos
}

// FilterUpdatedStaderValidatorRegistry is a free log retrieval operation binding the contract event 0xf583e0ea5b9579df6531ea89b81be75889dde9f34d35e2402ca38e93c0b5db0a.
//
// Solidity: event UpdatedStaderValidatorRegistry(address indexed staderValidatorRegistry)
func (_SocializingPool *SocializingPoolFilterer) FilterUpdatedStaderValidatorRegistry(opts *bind.FilterOpts, staderValidatorRegistry []common.Address) (*SocializingPoolUpdatedStaderValidatorRegistryIterator, error) {

	var staderValidatorRegistryRule []interface{}
	for _, staderValidatorRegistryItem := range staderValidatorRegistry {
		staderValidatorRegistryRule = append(staderValidatorRegistryRule, staderValidatorRegistryItem)
	}

	logs, sub, err := _SocializingPool.contract.FilterLogs(opts, "UpdatedStaderValidatorRegistry", staderValidatorRegistryRule)
	if err != nil {
		return nil, err
	}
	return &SocializingPoolUpdatedStaderValidatorRegistryIterator{contract: _SocializingPool.contract, event: "UpdatedStaderValidatorRegistry", logs: logs, sub: sub}, nil
}

// WatchUpdatedStaderValidatorRegistry is a free log subscription operation binding the contract event 0xf583e0ea5b9579df6531ea89b81be75889dde9f34d35e2402ca38e93c0b5db0a.
//
// Solidity: event UpdatedStaderValidatorRegistry(address indexed staderValidatorRegistry)
func (_SocializingPool *SocializingPoolFilterer) WatchUpdatedStaderValidatorRegistry(opts *bind.WatchOpts, sink chan<- *SocializingPoolUpdatedStaderValidatorRegistry, staderValidatorRegistry []common.Address) (event.Subscription, error) {

	var staderValidatorRegistryRule []interface{}
	for _, staderValidatorRegistryItem := range staderValidatorRegistry {
		staderValidatorRegistryRule = append(staderValidatorRegistryRule, staderValidatorRegistryItem)
	}

	logs, sub, err := _SocializingPool.contract.WatchLogs(opts, "UpdatedStaderValidatorRegistry", staderValidatorRegistryRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SocializingPoolUpdatedStaderValidatorRegistry)
				if err := _SocializingPool.contract.UnpackLog(event, "UpdatedStaderValidatorRegistry", log); err != nil {
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

// ParseUpdatedStaderValidatorRegistry is a log parse operation binding the contract event 0xf583e0ea5b9579df6531ea89b81be75889dde9f34d35e2402ca38e93c0b5db0a.
//
// Solidity: event UpdatedStaderValidatorRegistry(address indexed staderValidatorRegistry)
func (_SocializingPool *SocializingPoolFilterer) ParseUpdatedStaderValidatorRegistry(log types.Log) (*SocializingPoolUpdatedStaderValidatorRegistry, error) {
	event := new(SocializingPoolUpdatedStaderValidatorRegistry)
	if err := _SocializingPool.contract.UnpackLog(event, "UpdatedStaderValidatorRegistry", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SocializingPoolUserETHRewardsTransferredIterator is returned from FilterUserETHRewardsTransferred and is used to iterate over the raw logs and unpacked data for UserETHRewardsTransferred events raised by the SocializingPool contract.
type SocializingPoolUserETHRewardsTransferredIterator struct {
	Event *SocializingPoolUserETHRewardsTransferred // Event containing the contract specifics and raw log

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
func (it *SocializingPoolUserETHRewardsTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SocializingPoolUserETHRewardsTransferred)
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
		it.Event = new(SocializingPoolUserETHRewardsTransferred)
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
func (it *SocializingPoolUserETHRewardsTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SocializingPoolUserETHRewardsTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SocializingPoolUserETHRewardsTransferred represents a UserETHRewardsTransferred event raised by the SocializingPool contract.
type SocializingPoolUserETHRewardsTransferred struct {
	EthRewards *big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterUserETHRewardsTransferred is a free log retrieval operation binding the contract event 0x7083eaccdb1f2834d37a767b05f3b72d54217404ee1cb70aa1b774e4e8a02dda.
//
// Solidity: event UserETHRewardsTransferred(uint256 ethRewards)
func (_SocializingPool *SocializingPoolFilterer) FilterUserETHRewardsTransferred(opts *bind.FilterOpts) (*SocializingPoolUserETHRewardsTransferredIterator, error) {

	logs, sub, err := _SocializingPool.contract.FilterLogs(opts, "UserETHRewardsTransferred")
	if err != nil {
		return nil, err
	}
	return &SocializingPoolUserETHRewardsTransferredIterator{contract: _SocializingPool.contract, event: "UserETHRewardsTransferred", logs: logs, sub: sub}, nil
}

// WatchUserETHRewardsTransferred is a free log subscription operation binding the contract event 0x7083eaccdb1f2834d37a767b05f3b72d54217404ee1cb70aa1b774e4e8a02dda.
//
// Solidity: event UserETHRewardsTransferred(uint256 ethRewards)
func (_SocializingPool *SocializingPoolFilterer) WatchUserETHRewardsTransferred(opts *bind.WatchOpts, sink chan<- *SocializingPoolUserETHRewardsTransferred) (event.Subscription, error) {

	logs, sub, err := _SocializingPool.contract.WatchLogs(opts, "UserETHRewardsTransferred")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SocializingPoolUserETHRewardsTransferred)
				if err := _SocializingPool.contract.UnpackLog(event, "UserETHRewardsTransferred", log); err != nil {
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

// ParseUserETHRewardsTransferred is a log parse operation binding the contract event 0x7083eaccdb1f2834d37a767b05f3b72d54217404ee1cb70aa1b774e4e8a02dda.
//
// Solidity: event UserETHRewardsTransferred(uint256 ethRewards)
func (_SocializingPool *SocializingPoolFilterer) ParseUserETHRewardsTransferred(log types.Log) (*SocializingPoolUserETHRewardsTransferred, error) {
	event := new(SocializingPoolUserETHRewardsTransferred)
	if err := _SocializingPool.contract.UnpackLog(event, "UserETHRewardsTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
