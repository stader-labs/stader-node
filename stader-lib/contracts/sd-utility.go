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

// SDUtilityPoolMetaData contains all meta data concerning the SDUtilityPool contract.
var SDUtilityPoolMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"name\":\"AccrualBlockNumberNotLatest\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"CallerNotAuthorizedToRedeem\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"CannotFindRequestId\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InsufficientPoolBalance\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidAmountOfWithdraw\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"MaxLimitOnWithdrawRequestCountReached\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"requestId\",\"type\":\"uint256\"}],\"name\":\"RequestIdNotFinalized\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"SDTransferFailed\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"SDUtilizeLimitReached\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"UndelegationPeriodNotPassed\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"feeAccumulated\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"totalProtocolFee\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"totalUtilizedSD\",\"type\":\"uint256\"}],\"name\":\"AccruedFees\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"delegator\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"sdAmount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"sdXToMint\",\"type\":\"uint256\"}],\"name\":\"Delegated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"nextRequestIdToFinalize\",\"type\":\"uint256\"}],\"name\":\"FinalizedWithdrawRequest\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"protocolFeeFactor\",\"type\":\"uint256\"}],\"name\":\"ProtocolFeeFactorUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"delegator\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"sdAmount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"sdXAmount\",\"type\":\"uint256\"}],\"name\":\"Redeemed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"utilizer\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"repayAmount\",\"type\":\"uint256\"}],\"name\":\"Repaid\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"caller\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"sdToTransfer\",\"type\":\"uint256\"}],\"name\":\"RequestRedeemed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"utilizer\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"utilizeAmount\",\"type\":\"uint256\"}],\"name\":\"SDUtilized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"finalizationBatchLimit\",\"type\":\"uint256\"}],\"name\":\"UpdatedFinalizationBatchLimit\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"maxETHWorthOfSDPerValidator\",\"type\":\"uint256\"}],\"name\":\"UpdatedMaxETHWorthOfSDPerValidator\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"count\",\"type\":\"uint256\"}],\"name\":\"UpdatedMaxNonRedeemedDelegatorRequestCount\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"minBlockDelayToFinalizeRequest\",\"type\":\"uint256\"}],\"name\":\"UpdatedMinBlockDelayToFinalizeRequest\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_staderConfig\",\"type\":\"address\"}],\"name\":\"UpdatedStaderConfig\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"undelegationPeriodInBlocks\",\"type\":\"uint256\"}],\"name\":\"UpdatedUndelegationPeriodInBlocks\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"utilizationRatePerBlock\",\"type\":\"uint256\"}],\"name\":\"UtilizationRatePerBlockUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"utilizer\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"slashSDAmount\",\"type\":\"uint256\"}],\"name\":\"UtilizerSDSlashingHandled\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"caller\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"nextRequestId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"sdAmountToWithdraw\",\"type\":\"uint256\"}],\"name\":\"WithdrawRequestReceived\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"accrueFee\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"cTokenTotalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"requestId\",\"type\":\"uint256\"}],\"name\":\"claim\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"sdAmount\",\"type\":\"uint256\"}],\"name\":\"delegate\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"delegatorCTokenBalance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"delegatorWithdrawRequests\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amountOfCToken\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"sdExpected\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"sdFinalized\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"requestBlock\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"exchangeRateCurrent\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"exchangeRateStored\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"finalizeDelegatorWithdrawalRequest\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getDelegationRate\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_delegator\",\"type\":\"address\"}],\"name\":\"getDelegatorLatestSDBalance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getLatestExchangeRate\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getPoolAvailableSDBalance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_owner\",\"type\":\"address\"}],\"name\":\"getRequestIdsByDelegator\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"\",\"type\":\"uint256[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_utilizer\",\"type\":\"address\"}],\"name\":\"getUtilizerLatestBalance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_utilizer\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_slashSDAmount\",\"type\":\"uint256\"}],\"name\":\"handleUtilizerSDSlashing\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"liquidationCall\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"maxApproveSD\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"poolUtilization\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"repayAmount\",\"type\":\"uint256\"}],\"name\":\"repay\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"utilizer\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"repayAmount\",\"type\":\"uint256\"}],\"name\":\"repayOnBehalf\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"requestIdsByDelegatorAddress\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"cTokenAmount\",\"type\":\"uint256\"}],\"name\":\"requestWithdraw\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"sdAmount\",\"type\":\"uint256\"}],\"name\":\"requestWithdrawWithSDAmount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"utilizeAmount\",\"type\":\"uint256\"}],\"name\":\"utilize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"utilizeAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"nonTerminalKeyCount\",\"type\":\"uint256\"}],\"name\":\"utilizeWhileAddingKeys\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"utilizerBalanceCurrent\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"utilizerBalanceStored\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"utilizerData\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"principal\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"utilizeIndex\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// SDUtilityPoolABI is the input ABI used to generate the binding from.
// Deprecated: Use SDUtilityPoolMetaData.ABI instead.
var SDUtilityPoolABI = SDUtilityPoolMetaData.ABI

// SDUtilityPool is an auto generated Go binding around an Ethereum contract.
type SDUtilityPool struct {
	SDUtilityPoolCaller     // Read-only binding to the contract
	SDUtilityPoolTransactor // Write-only binding to the contract
	SDUtilityPoolFilterer   // Log filterer for contract events
}

// SDUtilityPoolCaller is an auto generated read-only Go binding around an Ethereum contract.
type SDUtilityPoolCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SDUtilityPoolTransactor is an auto generated write-only Go binding around an Ethereum contract.
type SDUtilityPoolTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SDUtilityPoolFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type SDUtilityPoolFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SDUtilityPoolSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type SDUtilityPoolSession struct {
	Contract     *SDUtilityPool    // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// SDUtilityPoolCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type SDUtilityPoolCallerSession struct {
	Contract *SDUtilityPoolCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts        // Call options to use throughout this session
}

// SDUtilityPoolTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type SDUtilityPoolTransactorSession struct {
	Contract     *SDUtilityPoolTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts        // Transaction auth options to use throughout this session
}

// SDUtilityPoolRaw is an auto generated low-level Go binding around an Ethereum contract.
type SDUtilityPoolRaw struct {
	Contract *SDUtilityPool // Generic contract binding to access the raw methods on
}

// SDUtilityPoolCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type SDUtilityPoolCallerRaw struct {
	Contract *SDUtilityPoolCaller // Generic read-only contract binding to access the raw methods on
}

// SDUtilityPoolTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type SDUtilityPoolTransactorRaw struct {
	Contract *SDUtilityPoolTransactor // Generic write-only contract binding to access the raw methods on
}

// NewSDUtilityPool creates a new instance of SDUtilityPool, bound to a specific deployed contract.
func NewSDUtilityPool(address common.Address, backend bind.ContractBackend) (*SDUtilityPool, error) {
	contract, err := bindSDUtilityPool(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &SDUtilityPool{SDUtilityPoolCaller: SDUtilityPoolCaller{contract: contract}, SDUtilityPoolTransactor: SDUtilityPoolTransactor{contract: contract}, SDUtilityPoolFilterer: SDUtilityPoolFilterer{contract: contract}}, nil
}

// NewSDUtilityPoolCaller creates a new read-only instance of SDUtilityPool, bound to a specific deployed contract.
func NewSDUtilityPoolCaller(address common.Address, caller bind.ContractCaller) (*SDUtilityPoolCaller, error) {
	contract, err := bindSDUtilityPool(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &SDUtilityPoolCaller{contract: contract}, nil
}

// NewSDUtilityPoolTransactor creates a new write-only instance of SDUtilityPool, bound to a specific deployed contract.
func NewSDUtilityPoolTransactor(address common.Address, transactor bind.ContractTransactor) (*SDUtilityPoolTransactor, error) {
	contract, err := bindSDUtilityPool(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &SDUtilityPoolTransactor{contract: contract}, nil
}

// NewSDUtilityPoolFilterer creates a new log filterer instance of SDUtilityPool, bound to a specific deployed contract.
func NewSDUtilityPoolFilterer(address common.Address, filterer bind.ContractFilterer) (*SDUtilityPoolFilterer, error) {
	contract, err := bindSDUtilityPool(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &SDUtilityPoolFilterer{contract: contract}, nil
}

// bindSDUtilityPool binds a generic wrapper to an already deployed contract.
func bindSDUtilityPool(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := SDUtilityPoolMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SDUtilityPool *SDUtilityPoolRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _SDUtilityPool.Contract.SDUtilityPoolCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SDUtilityPool *SDUtilityPoolRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SDUtilityPool.Contract.SDUtilityPoolTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SDUtilityPool *SDUtilityPoolRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SDUtilityPool.Contract.SDUtilityPoolTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SDUtilityPool *SDUtilityPoolCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _SDUtilityPool.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SDUtilityPool *SDUtilityPoolTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SDUtilityPool.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SDUtilityPool *SDUtilityPoolTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SDUtilityPool.Contract.contract.Transact(opts, method, params...)
}

// CTokenTotalSupply is a free data retrieval call binding the contract method 0x37a4adf7.
//
// Solidity: function cTokenTotalSupply() view returns(uint256)
func (_SDUtilityPool *SDUtilityPoolCaller) CTokenTotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _SDUtilityPool.contract.Call(opts, &out, "cTokenTotalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CTokenTotalSupply is a free data retrieval call binding the contract method 0x37a4adf7.
//
// Solidity: function cTokenTotalSupply() view returns(uint256)
func (_SDUtilityPool *SDUtilityPoolSession) CTokenTotalSupply() (*big.Int, error) {
	return _SDUtilityPool.Contract.CTokenTotalSupply(&_SDUtilityPool.CallOpts)
}

// CTokenTotalSupply is a free data retrieval call binding the contract method 0x37a4adf7.
//
// Solidity: function cTokenTotalSupply() view returns(uint256)
func (_SDUtilityPool *SDUtilityPoolCallerSession) CTokenTotalSupply() (*big.Int, error) {
	return _SDUtilityPool.Contract.CTokenTotalSupply(&_SDUtilityPool.CallOpts)
}

// DelegatorCTokenBalance is a free data retrieval call binding the contract method 0xabf9db02.
//
// Solidity: function delegatorCTokenBalance(address ) view returns(uint256)
func (_SDUtilityPool *SDUtilityPoolCaller) DelegatorCTokenBalance(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _SDUtilityPool.contract.Call(opts, &out, "delegatorCTokenBalance", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// DelegatorCTokenBalance is a free data retrieval call binding the contract method 0xabf9db02.
//
// Solidity: function delegatorCTokenBalance(address ) view returns(uint256)
func (_SDUtilityPool *SDUtilityPoolSession) DelegatorCTokenBalance(arg0 common.Address) (*big.Int, error) {
	return _SDUtilityPool.Contract.DelegatorCTokenBalance(&_SDUtilityPool.CallOpts, arg0)
}

// DelegatorCTokenBalance is a free data retrieval call binding the contract method 0xabf9db02.
//
// Solidity: function delegatorCTokenBalance(address ) view returns(uint256)
func (_SDUtilityPool *SDUtilityPoolCallerSession) DelegatorCTokenBalance(arg0 common.Address) (*big.Int, error) {
	return _SDUtilityPool.Contract.DelegatorCTokenBalance(&_SDUtilityPool.CallOpts, arg0)
}

// DelegatorWithdrawRequests is a free data retrieval call binding the contract method 0xe41b55d9.
//
// Solidity: function delegatorWithdrawRequests(uint256 ) view returns(address owner, uint256 amountOfCToken, uint256 sdExpected, uint256 sdFinalized, uint256 requestBlock)
func (_SDUtilityPool *SDUtilityPoolCaller) DelegatorWithdrawRequests(opts *bind.CallOpts, arg0 *big.Int) (struct {
	Owner          common.Address
	AmountOfCToken *big.Int
	SdExpected     *big.Int
	SdFinalized    *big.Int
	RequestBlock   *big.Int
}, error) {
	var out []interface{}
	err := _SDUtilityPool.contract.Call(opts, &out, "delegatorWithdrawRequests", arg0)

	outstruct := new(struct {
		Owner          common.Address
		AmountOfCToken *big.Int
		SdExpected     *big.Int
		SdFinalized    *big.Int
		RequestBlock   *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Owner = *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	outstruct.AmountOfCToken = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.SdExpected = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	outstruct.SdFinalized = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)
	outstruct.RequestBlock = *abi.ConvertType(out[4], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// DelegatorWithdrawRequests is a free data retrieval call binding the contract method 0xe41b55d9.
//
// Solidity: function delegatorWithdrawRequests(uint256 ) view returns(address owner, uint256 amountOfCToken, uint256 sdExpected, uint256 sdFinalized, uint256 requestBlock)
func (_SDUtilityPool *SDUtilityPoolSession) DelegatorWithdrawRequests(arg0 *big.Int) (struct {
	Owner          common.Address
	AmountOfCToken *big.Int
	SdExpected     *big.Int
	SdFinalized    *big.Int
	RequestBlock   *big.Int
}, error) {
	return _SDUtilityPool.Contract.DelegatorWithdrawRequests(&_SDUtilityPool.CallOpts, arg0)
}

// DelegatorWithdrawRequests is a free data retrieval call binding the contract method 0xe41b55d9.
//
// Solidity: function delegatorWithdrawRequests(uint256 ) view returns(address owner, uint256 amountOfCToken, uint256 sdExpected, uint256 sdFinalized, uint256 requestBlock)
func (_SDUtilityPool *SDUtilityPoolCallerSession) DelegatorWithdrawRequests(arg0 *big.Int) (struct {
	Owner          common.Address
	AmountOfCToken *big.Int
	SdExpected     *big.Int
	SdFinalized    *big.Int
	RequestBlock   *big.Int
}, error) {
	return _SDUtilityPool.Contract.DelegatorWithdrawRequests(&_SDUtilityPool.CallOpts, arg0)
}

// ExchangeRateStored is a free data retrieval call binding the contract method 0x182df0f5.
//
// Solidity: function exchangeRateStored() view returns(uint256)
func (_SDUtilityPool *SDUtilityPoolCaller) ExchangeRateStored(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _SDUtilityPool.contract.Call(opts, &out, "exchangeRateStored")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ExchangeRateStored is a free data retrieval call binding the contract method 0x182df0f5.
//
// Solidity: function exchangeRateStored() view returns(uint256)
func (_SDUtilityPool *SDUtilityPoolSession) ExchangeRateStored() (*big.Int, error) {
	return _SDUtilityPool.Contract.ExchangeRateStored(&_SDUtilityPool.CallOpts)
}

// ExchangeRateStored is a free data retrieval call binding the contract method 0x182df0f5.
//
// Solidity: function exchangeRateStored() view returns(uint256)
func (_SDUtilityPool *SDUtilityPoolCallerSession) ExchangeRateStored() (*big.Int, error) {
	return _SDUtilityPool.Contract.ExchangeRateStored(&_SDUtilityPool.CallOpts)
}

// GetDelegationRate is a free data retrieval call binding the contract method 0x0c7e5c23.
//
// Solidity: function getDelegationRate() view returns(uint256)
func (_SDUtilityPool *SDUtilityPoolCaller) GetDelegationRate(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _SDUtilityPool.contract.Call(opts, &out, "getDelegationRate")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetDelegationRate is a free data retrieval call binding the contract method 0x0c7e5c23.
//
// Solidity: function getDelegationRate() view returns(uint256)
func (_SDUtilityPool *SDUtilityPoolSession) GetDelegationRate() (*big.Int, error) {
	return _SDUtilityPool.Contract.GetDelegationRate(&_SDUtilityPool.CallOpts)
}

// GetDelegationRate is a free data retrieval call binding the contract method 0x0c7e5c23.
//
// Solidity: function getDelegationRate() view returns(uint256)
func (_SDUtilityPool *SDUtilityPoolCallerSession) GetDelegationRate() (*big.Int, error) {
	return _SDUtilityPool.Contract.GetDelegationRate(&_SDUtilityPool.CallOpts)
}

// GetDelegatorLatestSDBalance is a free data retrieval call binding the contract method 0x22291528.
//
// Solidity: function getDelegatorLatestSDBalance(address _delegator) view returns(uint256)
func (_SDUtilityPool *SDUtilityPoolCaller) GetDelegatorLatestSDBalance(opts *bind.CallOpts, _delegator common.Address) (*big.Int, error) {
	var out []interface{}
	err := _SDUtilityPool.contract.Call(opts, &out, "getDelegatorLatestSDBalance", _delegator)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetDelegatorLatestSDBalance is a free data retrieval call binding the contract method 0x22291528.
//
// Solidity: function getDelegatorLatestSDBalance(address _delegator) view returns(uint256)
func (_SDUtilityPool *SDUtilityPoolSession) GetDelegatorLatestSDBalance(_delegator common.Address) (*big.Int, error) {
	return _SDUtilityPool.Contract.GetDelegatorLatestSDBalance(&_SDUtilityPool.CallOpts, _delegator)
}

// GetDelegatorLatestSDBalance is a free data retrieval call binding the contract method 0x22291528.
//
// Solidity: function getDelegatorLatestSDBalance(address _delegator) view returns(uint256)
func (_SDUtilityPool *SDUtilityPoolCallerSession) GetDelegatorLatestSDBalance(_delegator common.Address) (*big.Int, error) {
	return _SDUtilityPool.Contract.GetDelegatorLatestSDBalance(&_SDUtilityPool.CallOpts, _delegator)
}

// GetLatestExchangeRate is a free data retrieval call binding the contract method 0x34d093f6.
//
// Solidity: function getLatestExchangeRate() view returns(uint256)
func (_SDUtilityPool *SDUtilityPoolCaller) GetLatestExchangeRate(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _SDUtilityPool.contract.Call(opts, &out, "getLatestExchangeRate")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetLatestExchangeRate is a free data retrieval call binding the contract method 0x34d093f6.
//
// Solidity: function getLatestExchangeRate() view returns(uint256)
func (_SDUtilityPool *SDUtilityPoolSession) GetLatestExchangeRate() (*big.Int, error) {
	return _SDUtilityPool.Contract.GetLatestExchangeRate(&_SDUtilityPool.CallOpts)
}

// GetLatestExchangeRate is a free data retrieval call binding the contract method 0x34d093f6.
//
// Solidity: function getLatestExchangeRate() view returns(uint256)
func (_SDUtilityPool *SDUtilityPoolCallerSession) GetLatestExchangeRate() (*big.Int, error) {
	return _SDUtilityPool.Contract.GetLatestExchangeRate(&_SDUtilityPool.CallOpts)
}

// GetPoolAvailableSDBalance is a free data retrieval call binding the contract method 0xda695857.
//
// Solidity: function getPoolAvailableSDBalance() view returns(uint256)
func (_SDUtilityPool *SDUtilityPoolCaller) GetPoolAvailableSDBalance(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _SDUtilityPool.contract.Call(opts, &out, "getPoolAvailableSDBalance")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetPoolAvailableSDBalance is a free data retrieval call binding the contract method 0xda695857.
//
// Solidity: function getPoolAvailableSDBalance() view returns(uint256)
func (_SDUtilityPool *SDUtilityPoolSession) GetPoolAvailableSDBalance() (*big.Int, error) {
	return _SDUtilityPool.Contract.GetPoolAvailableSDBalance(&_SDUtilityPool.CallOpts)
}

// GetPoolAvailableSDBalance is a free data retrieval call binding the contract method 0xda695857.
//
// Solidity: function getPoolAvailableSDBalance() view returns(uint256)
func (_SDUtilityPool *SDUtilityPoolCallerSession) GetPoolAvailableSDBalance() (*big.Int, error) {
	return _SDUtilityPool.Contract.GetPoolAvailableSDBalance(&_SDUtilityPool.CallOpts)
}

// GetRequestIdsByDelegator is a free data retrieval call binding the contract method 0x99775f40.
//
// Solidity: function getRequestIdsByDelegator(address _owner) view returns(uint256[])
func (_SDUtilityPool *SDUtilityPoolCaller) GetRequestIdsByDelegator(opts *bind.CallOpts, _owner common.Address) ([]*big.Int, error) {
	var out []interface{}
	err := _SDUtilityPool.contract.Call(opts, &out, "getRequestIdsByDelegator", _owner)

	if err != nil {
		return *new([]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([]*big.Int)).(*[]*big.Int)

	return out0, err

}

// GetRequestIdsByDelegator is a free data retrieval call binding the contract method 0x99775f40.
//
// Solidity: function getRequestIdsByDelegator(address _owner) view returns(uint256[])
func (_SDUtilityPool *SDUtilityPoolSession) GetRequestIdsByDelegator(_owner common.Address) ([]*big.Int, error) {
	return _SDUtilityPool.Contract.GetRequestIdsByDelegator(&_SDUtilityPool.CallOpts, _owner)
}

// GetRequestIdsByDelegator is a free data retrieval call binding the contract method 0x99775f40.
//
// Solidity: function getRequestIdsByDelegator(address _owner) view returns(uint256[])
func (_SDUtilityPool *SDUtilityPoolCallerSession) GetRequestIdsByDelegator(_owner common.Address) ([]*big.Int, error) {
	return _SDUtilityPool.Contract.GetRequestIdsByDelegator(&_SDUtilityPool.CallOpts, _owner)
}

// GetUtilizerLatestBalance is a free data retrieval call binding the contract method 0x36978412.
//
// Solidity: function getUtilizerLatestBalance(address _utilizer) view returns(uint256)
func (_SDUtilityPool *SDUtilityPoolCaller) GetUtilizerLatestBalance(opts *bind.CallOpts, _utilizer common.Address) (*big.Int, error) {
	var out []interface{}
	err := _SDUtilityPool.contract.Call(opts, &out, "getUtilizerLatestBalance", _utilizer)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetUtilizerLatestBalance is a free data retrieval call binding the contract method 0x36978412.
//
// Solidity: function getUtilizerLatestBalance(address _utilizer) view returns(uint256)
func (_SDUtilityPool *SDUtilityPoolSession) GetUtilizerLatestBalance(_utilizer common.Address) (*big.Int, error) {
	return _SDUtilityPool.Contract.GetUtilizerLatestBalance(&_SDUtilityPool.CallOpts, _utilizer)
}

// GetUtilizerLatestBalance is a free data retrieval call binding the contract method 0x36978412.
//
// Solidity: function getUtilizerLatestBalance(address _utilizer) view returns(uint256)
func (_SDUtilityPool *SDUtilityPoolCallerSession) GetUtilizerLatestBalance(_utilizer common.Address) (*big.Int, error) {
	return _SDUtilityPool.Contract.GetUtilizerLatestBalance(&_SDUtilityPool.CallOpts, _utilizer)
}

// PoolUtilization is a free data retrieval call binding the contract method 0x9a3263ee.
//
// Solidity: function poolUtilization() view returns(uint256)
func (_SDUtilityPool *SDUtilityPoolCaller) PoolUtilization(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _SDUtilityPool.contract.Call(opts, &out, "poolUtilization")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PoolUtilization is a free data retrieval call binding the contract method 0x9a3263ee.
//
// Solidity: function poolUtilization() view returns(uint256)
func (_SDUtilityPool *SDUtilityPoolSession) PoolUtilization() (*big.Int, error) {
	return _SDUtilityPool.Contract.PoolUtilization(&_SDUtilityPool.CallOpts)
}

// PoolUtilization is a free data retrieval call binding the contract method 0x9a3263ee.
//
// Solidity: function poolUtilization() view returns(uint256)
func (_SDUtilityPool *SDUtilityPoolCallerSession) PoolUtilization() (*big.Int, error) {
	return _SDUtilityPool.Contract.PoolUtilization(&_SDUtilityPool.CallOpts)
}

// RequestIdsByDelegatorAddress is a free data retrieval call binding the contract method 0x7844e3af.
//
// Solidity: function requestIdsByDelegatorAddress(address , uint256 ) view returns(uint256)
func (_SDUtilityPool *SDUtilityPoolCaller) RequestIdsByDelegatorAddress(opts *bind.CallOpts, arg0 common.Address, arg1 *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _SDUtilityPool.contract.Call(opts, &out, "requestIdsByDelegatorAddress", arg0, arg1)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// RequestIdsByDelegatorAddress is a free data retrieval call binding the contract method 0x7844e3af.
//
// Solidity: function requestIdsByDelegatorAddress(address , uint256 ) view returns(uint256)
func (_SDUtilityPool *SDUtilityPoolSession) RequestIdsByDelegatorAddress(arg0 common.Address, arg1 *big.Int) (*big.Int, error) {
	return _SDUtilityPool.Contract.RequestIdsByDelegatorAddress(&_SDUtilityPool.CallOpts, arg0, arg1)
}

// RequestIdsByDelegatorAddress is a free data retrieval call binding the contract method 0x7844e3af.
//
// Solidity: function requestIdsByDelegatorAddress(address , uint256 ) view returns(uint256)
func (_SDUtilityPool *SDUtilityPoolCallerSession) RequestIdsByDelegatorAddress(arg0 common.Address, arg1 *big.Int) (*big.Int, error) {
	return _SDUtilityPool.Contract.RequestIdsByDelegatorAddress(&_SDUtilityPool.CallOpts, arg0, arg1)
}

// UtilizerBalanceStored is a free data retrieval call binding the contract method 0x6e236aee.
//
// Solidity: function utilizerBalanceStored(address account) view returns(uint256)
func (_SDUtilityPool *SDUtilityPoolCaller) UtilizerBalanceStored(opts *bind.CallOpts, account common.Address) (*big.Int, error) {
	var out []interface{}
	err := _SDUtilityPool.contract.Call(opts, &out, "utilizerBalanceStored", account)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// UtilizerBalanceStored is a free data retrieval call binding the contract method 0x6e236aee.
//
// Solidity: function utilizerBalanceStored(address account) view returns(uint256)
func (_SDUtilityPool *SDUtilityPoolSession) UtilizerBalanceStored(account common.Address) (*big.Int, error) {
	return _SDUtilityPool.Contract.UtilizerBalanceStored(&_SDUtilityPool.CallOpts, account)
}

// UtilizerBalanceStored is a free data retrieval call binding the contract method 0x6e236aee.
//
// Solidity: function utilizerBalanceStored(address account) view returns(uint256)
func (_SDUtilityPool *SDUtilityPoolCallerSession) UtilizerBalanceStored(account common.Address) (*big.Int, error) {
	return _SDUtilityPool.Contract.UtilizerBalanceStored(&_SDUtilityPool.CallOpts, account)
}

// UtilizerData is a free data retrieval call binding the contract method 0x9e070088.
//
// Solidity: function utilizerData(address ) view returns(uint256 principal, uint256 utilizeIndex)
func (_SDUtilityPool *SDUtilityPoolCaller) UtilizerData(opts *bind.CallOpts, arg0 common.Address) (struct {
	Principal    *big.Int
	UtilizeIndex *big.Int
}, error) {
	var out []interface{}
	err := _SDUtilityPool.contract.Call(opts, &out, "utilizerData", arg0)

	outstruct := new(struct {
		Principal    *big.Int
		UtilizeIndex *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Principal = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.UtilizeIndex = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// UtilizerData is a free data retrieval call binding the contract method 0x9e070088.
//
// Solidity: function utilizerData(address ) view returns(uint256 principal, uint256 utilizeIndex)
func (_SDUtilityPool *SDUtilityPoolSession) UtilizerData(arg0 common.Address) (struct {
	Principal    *big.Int
	UtilizeIndex *big.Int
}, error) {
	return _SDUtilityPool.Contract.UtilizerData(&_SDUtilityPool.CallOpts, arg0)
}

// UtilizerData is a free data retrieval call binding the contract method 0x9e070088.
//
// Solidity: function utilizerData(address ) view returns(uint256 principal, uint256 utilizeIndex)
func (_SDUtilityPool *SDUtilityPoolCallerSession) UtilizerData(arg0 common.Address) (struct {
	Principal    *big.Int
	UtilizeIndex *big.Int
}, error) {
	return _SDUtilityPool.Contract.UtilizerData(&_SDUtilityPool.CallOpts, arg0)
}

// AccrueFee is a paid mutator transaction binding the contract method 0xb26cc394.
//
// Solidity: function accrueFee() returns()
func (_SDUtilityPool *SDUtilityPoolTransactor) AccrueFee(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SDUtilityPool.contract.Transact(opts, "accrueFee")
}

// AccrueFee is a paid mutator transaction binding the contract method 0xb26cc394.
//
// Solidity: function accrueFee() returns()
func (_SDUtilityPool *SDUtilityPoolSession) AccrueFee() (*types.Transaction, error) {
	return _SDUtilityPool.Contract.AccrueFee(&_SDUtilityPool.TransactOpts)
}

// AccrueFee is a paid mutator transaction binding the contract method 0xb26cc394.
//
// Solidity: function accrueFee() returns()
func (_SDUtilityPool *SDUtilityPoolTransactorSession) AccrueFee() (*types.Transaction, error) {
	return _SDUtilityPool.Contract.AccrueFee(&_SDUtilityPool.TransactOpts)
}

// Claim is a paid mutator transaction binding the contract method 0x379607f5.
//
// Solidity: function claim(uint256 requestId) returns()
func (_SDUtilityPool *SDUtilityPoolTransactor) Claim(opts *bind.TransactOpts, requestId *big.Int) (*types.Transaction, error) {
	return _SDUtilityPool.contract.Transact(opts, "claim", requestId)
}

// Claim is a paid mutator transaction binding the contract method 0x379607f5.
//
// Solidity: function claim(uint256 requestId) returns()
func (_SDUtilityPool *SDUtilityPoolSession) Claim(requestId *big.Int) (*types.Transaction, error) {
	return _SDUtilityPool.Contract.Claim(&_SDUtilityPool.TransactOpts, requestId)
}

// Claim is a paid mutator transaction binding the contract method 0x379607f5.
//
// Solidity: function claim(uint256 requestId) returns()
func (_SDUtilityPool *SDUtilityPoolTransactorSession) Claim(requestId *big.Int) (*types.Transaction, error) {
	return _SDUtilityPool.Contract.Claim(&_SDUtilityPool.TransactOpts, requestId)
}

// Delegate is a paid mutator transaction binding the contract method 0x9fa6dd35.
//
// Solidity: function delegate(uint256 sdAmount) returns()
func (_SDUtilityPool *SDUtilityPoolTransactor) Delegate(opts *bind.TransactOpts, sdAmount *big.Int) (*types.Transaction, error) {
	return _SDUtilityPool.contract.Transact(opts, "delegate", sdAmount)
}

// Delegate is a paid mutator transaction binding the contract method 0x9fa6dd35.
//
// Solidity: function delegate(uint256 sdAmount) returns()
func (_SDUtilityPool *SDUtilityPoolSession) Delegate(sdAmount *big.Int) (*types.Transaction, error) {
	return _SDUtilityPool.Contract.Delegate(&_SDUtilityPool.TransactOpts, sdAmount)
}

// Delegate is a paid mutator transaction binding the contract method 0x9fa6dd35.
//
// Solidity: function delegate(uint256 sdAmount) returns()
func (_SDUtilityPool *SDUtilityPoolTransactorSession) Delegate(sdAmount *big.Int) (*types.Transaction, error) {
	return _SDUtilityPool.Contract.Delegate(&_SDUtilityPool.TransactOpts, sdAmount)
}

// ExchangeRateCurrent is a paid mutator transaction binding the contract method 0xbd6d894d.
//
// Solidity: function exchangeRateCurrent() returns(uint256)
func (_SDUtilityPool *SDUtilityPoolTransactor) ExchangeRateCurrent(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SDUtilityPool.contract.Transact(opts, "exchangeRateCurrent")
}

// ExchangeRateCurrent is a paid mutator transaction binding the contract method 0xbd6d894d.
//
// Solidity: function exchangeRateCurrent() returns(uint256)
func (_SDUtilityPool *SDUtilityPoolSession) ExchangeRateCurrent() (*types.Transaction, error) {
	return _SDUtilityPool.Contract.ExchangeRateCurrent(&_SDUtilityPool.TransactOpts)
}

// ExchangeRateCurrent is a paid mutator transaction binding the contract method 0xbd6d894d.
//
// Solidity: function exchangeRateCurrent() returns(uint256)
func (_SDUtilityPool *SDUtilityPoolTransactorSession) ExchangeRateCurrent() (*types.Transaction, error) {
	return _SDUtilityPool.Contract.ExchangeRateCurrent(&_SDUtilityPool.TransactOpts)
}

// FinalizeDelegatorWithdrawalRequest is a paid mutator transaction binding the contract method 0x71898b4e.
//
// Solidity: function finalizeDelegatorWithdrawalRequest() returns()
func (_SDUtilityPool *SDUtilityPoolTransactor) FinalizeDelegatorWithdrawalRequest(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SDUtilityPool.contract.Transact(opts, "finalizeDelegatorWithdrawalRequest")
}

// FinalizeDelegatorWithdrawalRequest is a paid mutator transaction binding the contract method 0x71898b4e.
//
// Solidity: function finalizeDelegatorWithdrawalRequest() returns()
func (_SDUtilityPool *SDUtilityPoolSession) FinalizeDelegatorWithdrawalRequest() (*types.Transaction, error) {
	return _SDUtilityPool.Contract.FinalizeDelegatorWithdrawalRequest(&_SDUtilityPool.TransactOpts)
}

// FinalizeDelegatorWithdrawalRequest is a paid mutator transaction binding the contract method 0x71898b4e.
//
// Solidity: function finalizeDelegatorWithdrawalRequest() returns()
func (_SDUtilityPool *SDUtilityPoolTransactorSession) FinalizeDelegatorWithdrawalRequest() (*types.Transaction, error) {
	return _SDUtilityPool.Contract.FinalizeDelegatorWithdrawalRequest(&_SDUtilityPool.TransactOpts)
}

// HandleUtilizerSDSlashing is a paid mutator transaction binding the contract method 0xfa72bf63.
//
// Solidity: function handleUtilizerSDSlashing(address _utilizer, uint256 _slashSDAmount) returns()
func (_SDUtilityPool *SDUtilityPoolTransactor) HandleUtilizerSDSlashing(opts *bind.TransactOpts, _utilizer common.Address, _slashSDAmount *big.Int) (*types.Transaction, error) {
	return _SDUtilityPool.contract.Transact(opts, "handleUtilizerSDSlashing", _utilizer, _slashSDAmount)
}

// HandleUtilizerSDSlashing is a paid mutator transaction binding the contract method 0xfa72bf63.
//
// Solidity: function handleUtilizerSDSlashing(address _utilizer, uint256 _slashSDAmount) returns()
func (_SDUtilityPool *SDUtilityPoolSession) HandleUtilizerSDSlashing(_utilizer common.Address, _slashSDAmount *big.Int) (*types.Transaction, error) {
	return _SDUtilityPool.Contract.HandleUtilizerSDSlashing(&_SDUtilityPool.TransactOpts, _utilizer, _slashSDAmount)
}

// HandleUtilizerSDSlashing is a paid mutator transaction binding the contract method 0xfa72bf63.
//
// Solidity: function handleUtilizerSDSlashing(address _utilizer, uint256 _slashSDAmount) returns()
func (_SDUtilityPool *SDUtilityPoolTransactorSession) HandleUtilizerSDSlashing(_utilizer common.Address, _slashSDAmount *big.Int) (*types.Transaction, error) {
	return _SDUtilityPool.Contract.HandleUtilizerSDSlashing(&_SDUtilityPool.TransactOpts, _utilizer, _slashSDAmount)
}

// LiquidationCall is a paid mutator transaction binding the contract method 0xd2ac4a3d.
//
// Solidity: function liquidationCall(address account) returns()
func (_SDUtilityPool *SDUtilityPoolTransactor) LiquidationCall(opts *bind.TransactOpts, account common.Address) (*types.Transaction, error) {
	return _SDUtilityPool.contract.Transact(opts, "liquidationCall", account)
}

// LiquidationCall is a paid mutator transaction binding the contract method 0xd2ac4a3d.
//
// Solidity: function liquidationCall(address account) returns()
func (_SDUtilityPool *SDUtilityPoolSession) LiquidationCall(account common.Address) (*types.Transaction, error) {
	return _SDUtilityPool.Contract.LiquidationCall(&_SDUtilityPool.TransactOpts, account)
}

// LiquidationCall is a paid mutator transaction binding the contract method 0xd2ac4a3d.
//
// Solidity: function liquidationCall(address account) returns()
func (_SDUtilityPool *SDUtilityPoolTransactorSession) LiquidationCall(account common.Address) (*types.Transaction, error) {
	return _SDUtilityPool.Contract.LiquidationCall(&_SDUtilityPool.TransactOpts, account)
}

// MaxApproveSD is a paid mutator transaction binding the contract method 0x3e04cd35.
//
// Solidity: function maxApproveSD() returns()
func (_SDUtilityPool *SDUtilityPoolTransactor) MaxApproveSD(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SDUtilityPool.contract.Transact(opts, "maxApproveSD")
}

// MaxApproveSD is a paid mutator transaction binding the contract method 0x3e04cd35.
//
// Solidity: function maxApproveSD() returns()
func (_SDUtilityPool *SDUtilityPoolSession) MaxApproveSD() (*types.Transaction, error) {
	return _SDUtilityPool.Contract.MaxApproveSD(&_SDUtilityPool.TransactOpts)
}

// MaxApproveSD is a paid mutator transaction binding the contract method 0x3e04cd35.
//
// Solidity: function maxApproveSD() returns()
func (_SDUtilityPool *SDUtilityPoolTransactorSession) MaxApproveSD() (*types.Transaction, error) {
	return _SDUtilityPool.Contract.MaxApproveSD(&_SDUtilityPool.TransactOpts)
}

// Repay is a paid mutator transaction binding the contract method 0x371fd8e6.
//
// Solidity: function repay(uint256 repayAmount) returns()
func (_SDUtilityPool *SDUtilityPoolTransactor) Repay(opts *bind.TransactOpts, repayAmount *big.Int) (*types.Transaction, error) {
	return _SDUtilityPool.contract.Transact(opts, "repay", repayAmount)
}

// Repay is a paid mutator transaction binding the contract method 0x371fd8e6.
//
// Solidity: function repay(uint256 repayAmount) returns()
func (_SDUtilityPool *SDUtilityPoolSession) Repay(repayAmount *big.Int) (*types.Transaction, error) {
	return _SDUtilityPool.Contract.Repay(&_SDUtilityPool.TransactOpts, repayAmount)
}

// Repay is a paid mutator transaction binding the contract method 0x371fd8e6.
//
// Solidity: function repay(uint256 repayAmount) returns()
func (_SDUtilityPool *SDUtilityPoolTransactorSession) Repay(repayAmount *big.Int) (*types.Transaction, error) {
	return _SDUtilityPool.Contract.Repay(&_SDUtilityPool.TransactOpts, repayAmount)
}

// RepayOnBehalf is a paid mutator transaction binding the contract method 0x9f689e0b.
//
// Solidity: function repayOnBehalf(address utilizer, uint256 repayAmount) returns()
func (_SDUtilityPool *SDUtilityPoolTransactor) RepayOnBehalf(opts *bind.TransactOpts, utilizer common.Address, repayAmount *big.Int) (*types.Transaction, error) {
	return _SDUtilityPool.contract.Transact(opts, "repayOnBehalf", utilizer, repayAmount)
}

// RepayOnBehalf is a paid mutator transaction binding the contract method 0x9f689e0b.
//
// Solidity: function repayOnBehalf(address utilizer, uint256 repayAmount) returns()
func (_SDUtilityPool *SDUtilityPoolSession) RepayOnBehalf(utilizer common.Address, repayAmount *big.Int) (*types.Transaction, error) {
	return _SDUtilityPool.Contract.RepayOnBehalf(&_SDUtilityPool.TransactOpts, utilizer, repayAmount)
}

// RepayOnBehalf is a paid mutator transaction binding the contract method 0x9f689e0b.
//
// Solidity: function repayOnBehalf(address utilizer, uint256 repayAmount) returns()
func (_SDUtilityPool *SDUtilityPoolTransactorSession) RepayOnBehalf(utilizer common.Address, repayAmount *big.Int) (*types.Transaction, error) {
	return _SDUtilityPool.Contract.RepayOnBehalf(&_SDUtilityPool.TransactOpts, utilizer, repayAmount)
}

// RequestWithdraw is a paid mutator transaction binding the contract method 0x745400c9.
//
// Solidity: function requestWithdraw(uint256 cTokenAmount) returns(uint256)
func (_SDUtilityPool *SDUtilityPoolTransactor) RequestWithdraw(opts *bind.TransactOpts, cTokenAmount *big.Int) (*types.Transaction, error) {
	return _SDUtilityPool.contract.Transact(opts, "requestWithdraw", cTokenAmount)
}

// RequestWithdraw is a paid mutator transaction binding the contract method 0x745400c9.
//
// Solidity: function requestWithdraw(uint256 cTokenAmount) returns(uint256)
func (_SDUtilityPool *SDUtilityPoolSession) RequestWithdraw(cTokenAmount *big.Int) (*types.Transaction, error) {
	return _SDUtilityPool.Contract.RequestWithdraw(&_SDUtilityPool.TransactOpts, cTokenAmount)
}

// RequestWithdraw is a paid mutator transaction binding the contract method 0x745400c9.
//
// Solidity: function requestWithdraw(uint256 cTokenAmount) returns(uint256)
func (_SDUtilityPool *SDUtilityPoolTransactorSession) RequestWithdraw(cTokenAmount *big.Int) (*types.Transaction, error) {
	return _SDUtilityPool.Contract.RequestWithdraw(&_SDUtilityPool.TransactOpts, cTokenAmount)
}

// RequestWithdrawWithSDAmount is a paid mutator transaction binding the contract method 0xc51cd1cc.
//
// Solidity: function requestWithdrawWithSDAmount(uint256 sdAmount) returns(uint256)
func (_SDUtilityPool *SDUtilityPoolTransactor) RequestWithdrawWithSDAmount(opts *bind.TransactOpts, sdAmount *big.Int) (*types.Transaction, error) {
	return _SDUtilityPool.contract.Transact(opts, "requestWithdrawWithSDAmount", sdAmount)
}

// RequestWithdrawWithSDAmount is a paid mutator transaction binding the contract method 0xc51cd1cc.
//
// Solidity: function requestWithdrawWithSDAmount(uint256 sdAmount) returns(uint256)
func (_SDUtilityPool *SDUtilityPoolSession) RequestWithdrawWithSDAmount(sdAmount *big.Int) (*types.Transaction, error) {
	return _SDUtilityPool.Contract.RequestWithdrawWithSDAmount(&_SDUtilityPool.TransactOpts, sdAmount)
}

// RequestWithdrawWithSDAmount is a paid mutator transaction binding the contract method 0xc51cd1cc.
//
// Solidity: function requestWithdrawWithSDAmount(uint256 sdAmount) returns(uint256)
func (_SDUtilityPool *SDUtilityPoolTransactorSession) RequestWithdrawWithSDAmount(sdAmount *big.Int) (*types.Transaction, error) {
	return _SDUtilityPool.Contract.RequestWithdrawWithSDAmount(&_SDUtilityPool.TransactOpts, sdAmount)
}

// Utilize is a paid mutator transaction binding the contract method 0xec29c551.
//
// Solidity: function utilize(uint256 utilizeAmount) returns()
func (_SDUtilityPool *SDUtilityPoolTransactor) Utilize(opts *bind.TransactOpts, utilizeAmount *big.Int) (*types.Transaction, error) {
	return _SDUtilityPool.contract.Transact(opts, "utilize", utilizeAmount)
}

// Utilize is a paid mutator transaction binding the contract method 0xec29c551.
//
// Solidity: function utilize(uint256 utilizeAmount) returns()
func (_SDUtilityPool *SDUtilityPoolSession) Utilize(utilizeAmount *big.Int) (*types.Transaction, error) {
	return _SDUtilityPool.Contract.Utilize(&_SDUtilityPool.TransactOpts, utilizeAmount)
}

// Utilize is a paid mutator transaction binding the contract method 0xec29c551.
//
// Solidity: function utilize(uint256 utilizeAmount) returns()
func (_SDUtilityPool *SDUtilityPoolTransactorSession) Utilize(utilizeAmount *big.Int) (*types.Transaction, error) {
	return _SDUtilityPool.Contract.Utilize(&_SDUtilityPool.TransactOpts, utilizeAmount)
}

// UtilizeWhileAddingKeys is a paid mutator transaction binding the contract method 0x23c4ac1a.
//
// Solidity: function utilizeWhileAddingKeys(address operator, uint256 utilizeAmount, uint256 nonTerminalKeyCount) returns()
func (_SDUtilityPool *SDUtilityPoolTransactor) UtilizeWhileAddingKeys(opts *bind.TransactOpts, operator common.Address, utilizeAmount *big.Int, nonTerminalKeyCount *big.Int) (*types.Transaction, error) {
	return _SDUtilityPool.contract.Transact(opts, "utilizeWhileAddingKeys", operator, utilizeAmount, nonTerminalKeyCount)
}

// UtilizeWhileAddingKeys is a paid mutator transaction binding the contract method 0x23c4ac1a.
//
// Solidity: function utilizeWhileAddingKeys(address operator, uint256 utilizeAmount, uint256 nonTerminalKeyCount) returns()
func (_SDUtilityPool *SDUtilityPoolSession) UtilizeWhileAddingKeys(operator common.Address, utilizeAmount *big.Int, nonTerminalKeyCount *big.Int) (*types.Transaction, error) {
	return _SDUtilityPool.Contract.UtilizeWhileAddingKeys(&_SDUtilityPool.TransactOpts, operator, utilizeAmount, nonTerminalKeyCount)
}

// UtilizeWhileAddingKeys is a paid mutator transaction binding the contract method 0x23c4ac1a.
//
// Solidity: function utilizeWhileAddingKeys(address operator, uint256 utilizeAmount, uint256 nonTerminalKeyCount) returns()
func (_SDUtilityPool *SDUtilityPoolTransactorSession) UtilizeWhileAddingKeys(operator common.Address, utilizeAmount *big.Int, nonTerminalKeyCount *big.Int) (*types.Transaction, error) {
	return _SDUtilityPool.Contract.UtilizeWhileAddingKeys(&_SDUtilityPool.TransactOpts, operator, utilizeAmount, nonTerminalKeyCount)
}

// UtilizerBalanceCurrent is a paid mutator transaction binding the contract method 0xe65efbe4.
//
// Solidity: function utilizerBalanceCurrent(address account) returns(uint256)
func (_SDUtilityPool *SDUtilityPoolTransactor) UtilizerBalanceCurrent(opts *bind.TransactOpts, account common.Address) (*types.Transaction, error) {
	return _SDUtilityPool.contract.Transact(opts, "utilizerBalanceCurrent", account)
}

// UtilizerBalanceCurrent is a paid mutator transaction binding the contract method 0xe65efbe4.
//
// Solidity: function utilizerBalanceCurrent(address account) returns(uint256)
func (_SDUtilityPool *SDUtilityPoolSession) UtilizerBalanceCurrent(account common.Address) (*types.Transaction, error) {
	return _SDUtilityPool.Contract.UtilizerBalanceCurrent(&_SDUtilityPool.TransactOpts, account)
}

// UtilizerBalanceCurrent is a paid mutator transaction binding the contract method 0xe65efbe4.
//
// Solidity: function utilizerBalanceCurrent(address account) returns(uint256)
func (_SDUtilityPool *SDUtilityPoolTransactorSession) UtilizerBalanceCurrent(account common.Address) (*types.Transaction, error) {
	return _SDUtilityPool.Contract.UtilizerBalanceCurrent(&_SDUtilityPool.TransactOpts, account)
}

// SDUtilityPoolAccruedFeesIterator is returned from FilterAccruedFees and is used to iterate over the raw logs and unpacked data for AccruedFees events raised by the SDUtilityPool contract.
type SDUtilityPoolAccruedFeesIterator struct {
	Event *SDUtilityPoolAccruedFees // Event containing the contract specifics and raw log

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
func (it *SDUtilityPoolAccruedFeesIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SDUtilityPoolAccruedFees)
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
		it.Event = new(SDUtilityPoolAccruedFees)
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
func (it *SDUtilityPoolAccruedFeesIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SDUtilityPoolAccruedFeesIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SDUtilityPoolAccruedFees represents a AccruedFees event raised by the SDUtilityPool contract.
type SDUtilityPoolAccruedFees struct {
	FeeAccumulated   *big.Int
	TotalProtocolFee *big.Int
	TotalUtilizedSD  *big.Int
	Raw              types.Log // Blockchain specific contextual infos
}

// FilterAccruedFees is a free log retrieval operation binding the contract event 0x7119249986febcaf2eaa8565a4a5f37df51951d3933512847e77ad489aff89a5.
//
// Solidity: event AccruedFees(uint256 feeAccumulated, uint256 totalProtocolFee, uint256 totalUtilizedSD)
func (_SDUtilityPool *SDUtilityPoolFilterer) FilterAccruedFees(opts *bind.FilterOpts) (*SDUtilityPoolAccruedFeesIterator, error) {

	logs, sub, err := _SDUtilityPool.contract.FilterLogs(opts, "AccruedFees")
	if err != nil {
		return nil, err
	}
	return &SDUtilityPoolAccruedFeesIterator{contract: _SDUtilityPool.contract, event: "AccruedFees", logs: logs, sub: sub}, nil
}

// WatchAccruedFees is a free log subscription operation binding the contract event 0x7119249986febcaf2eaa8565a4a5f37df51951d3933512847e77ad489aff89a5.
//
// Solidity: event AccruedFees(uint256 feeAccumulated, uint256 totalProtocolFee, uint256 totalUtilizedSD)
func (_SDUtilityPool *SDUtilityPoolFilterer) WatchAccruedFees(opts *bind.WatchOpts, sink chan<- *SDUtilityPoolAccruedFees) (event.Subscription, error) {

	logs, sub, err := _SDUtilityPool.contract.WatchLogs(opts, "AccruedFees")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SDUtilityPoolAccruedFees)
				if err := _SDUtilityPool.contract.UnpackLog(event, "AccruedFees", log); err != nil {
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

// ParseAccruedFees is a log parse operation binding the contract event 0x7119249986febcaf2eaa8565a4a5f37df51951d3933512847e77ad489aff89a5.
//
// Solidity: event AccruedFees(uint256 feeAccumulated, uint256 totalProtocolFee, uint256 totalUtilizedSD)
func (_SDUtilityPool *SDUtilityPoolFilterer) ParseAccruedFees(log types.Log) (*SDUtilityPoolAccruedFees, error) {
	event := new(SDUtilityPoolAccruedFees)
	if err := _SDUtilityPool.contract.UnpackLog(event, "AccruedFees", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SDUtilityPoolDelegatedIterator is returned from FilterDelegated and is used to iterate over the raw logs and unpacked data for Delegated events raised by the SDUtilityPool contract.
type SDUtilityPoolDelegatedIterator struct {
	Event *SDUtilityPoolDelegated // Event containing the contract specifics and raw log

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
func (it *SDUtilityPoolDelegatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SDUtilityPoolDelegated)
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
		it.Event = new(SDUtilityPoolDelegated)
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
func (it *SDUtilityPoolDelegatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SDUtilityPoolDelegatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SDUtilityPoolDelegated represents a Delegated event raised by the SDUtilityPool contract.
type SDUtilityPoolDelegated struct {
	Delegator common.Address
	SdAmount  *big.Int
	SdXToMint *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterDelegated is a free log retrieval operation binding the contract event 0x9a8f44850296624dadfd9c246d17e47171d35727a181bd090aa14bbbe00238bb.
//
// Solidity: event Delegated(address indexed delegator, uint256 sdAmount, uint256 sdXToMint)
func (_SDUtilityPool *SDUtilityPoolFilterer) FilterDelegated(opts *bind.FilterOpts, delegator []common.Address) (*SDUtilityPoolDelegatedIterator, error) {

	var delegatorRule []interface{}
	for _, delegatorItem := range delegator {
		delegatorRule = append(delegatorRule, delegatorItem)
	}

	logs, sub, err := _SDUtilityPool.contract.FilterLogs(opts, "Delegated", delegatorRule)
	if err != nil {
		return nil, err
	}
	return &SDUtilityPoolDelegatedIterator{contract: _SDUtilityPool.contract, event: "Delegated", logs: logs, sub: sub}, nil
}

// WatchDelegated is a free log subscription operation binding the contract event 0x9a8f44850296624dadfd9c246d17e47171d35727a181bd090aa14bbbe00238bb.
//
// Solidity: event Delegated(address indexed delegator, uint256 sdAmount, uint256 sdXToMint)
func (_SDUtilityPool *SDUtilityPoolFilterer) WatchDelegated(opts *bind.WatchOpts, sink chan<- *SDUtilityPoolDelegated, delegator []common.Address) (event.Subscription, error) {

	var delegatorRule []interface{}
	for _, delegatorItem := range delegator {
		delegatorRule = append(delegatorRule, delegatorItem)
	}

	logs, sub, err := _SDUtilityPool.contract.WatchLogs(opts, "Delegated", delegatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SDUtilityPoolDelegated)
				if err := _SDUtilityPool.contract.UnpackLog(event, "Delegated", log); err != nil {
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

// ParseDelegated is a log parse operation binding the contract event 0x9a8f44850296624dadfd9c246d17e47171d35727a181bd090aa14bbbe00238bb.
//
// Solidity: event Delegated(address indexed delegator, uint256 sdAmount, uint256 sdXToMint)
func (_SDUtilityPool *SDUtilityPoolFilterer) ParseDelegated(log types.Log) (*SDUtilityPoolDelegated, error) {
	event := new(SDUtilityPoolDelegated)
	if err := _SDUtilityPool.contract.UnpackLog(event, "Delegated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SDUtilityPoolFinalizedWithdrawRequestIterator is returned from FilterFinalizedWithdrawRequest and is used to iterate over the raw logs and unpacked data for FinalizedWithdrawRequest events raised by the SDUtilityPool contract.
type SDUtilityPoolFinalizedWithdrawRequestIterator struct {
	Event *SDUtilityPoolFinalizedWithdrawRequest // Event containing the contract specifics and raw log

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
func (it *SDUtilityPoolFinalizedWithdrawRequestIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SDUtilityPoolFinalizedWithdrawRequest)
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
		it.Event = new(SDUtilityPoolFinalizedWithdrawRequest)
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
func (it *SDUtilityPoolFinalizedWithdrawRequestIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SDUtilityPoolFinalizedWithdrawRequestIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SDUtilityPoolFinalizedWithdrawRequest represents a FinalizedWithdrawRequest event raised by the SDUtilityPool contract.
type SDUtilityPoolFinalizedWithdrawRequest struct {
	NextRequestIdToFinalize *big.Int
	Raw                     types.Log // Blockchain specific contextual infos
}

// FilterFinalizedWithdrawRequest is a free log retrieval operation binding the contract event 0x12a00f5e4c3614409f2dd90dc5be91b9b64ef89bac58a5b034ec0094376dbd37.
//
// Solidity: event FinalizedWithdrawRequest(uint256 nextRequestIdToFinalize)
func (_SDUtilityPool *SDUtilityPoolFilterer) FilterFinalizedWithdrawRequest(opts *bind.FilterOpts) (*SDUtilityPoolFinalizedWithdrawRequestIterator, error) {

	logs, sub, err := _SDUtilityPool.contract.FilterLogs(opts, "FinalizedWithdrawRequest")
	if err != nil {
		return nil, err
	}
	return &SDUtilityPoolFinalizedWithdrawRequestIterator{contract: _SDUtilityPool.contract, event: "FinalizedWithdrawRequest", logs: logs, sub: sub}, nil
}

// WatchFinalizedWithdrawRequest is a free log subscription operation binding the contract event 0x12a00f5e4c3614409f2dd90dc5be91b9b64ef89bac58a5b034ec0094376dbd37.
//
// Solidity: event FinalizedWithdrawRequest(uint256 nextRequestIdToFinalize)
func (_SDUtilityPool *SDUtilityPoolFilterer) WatchFinalizedWithdrawRequest(opts *bind.WatchOpts, sink chan<- *SDUtilityPoolFinalizedWithdrawRequest) (event.Subscription, error) {

	logs, sub, err := _SDUtilityPool.contract.WatchLogs(opts, "FinalizedWithdrawRequest")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SDUtilityPoolFinalizedWithdrawRequest)
				if err := _SDUtilityPool.contract.UnpackLog(event, "FinalizedWithdrawRequest", log); err != nil {
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

// ParseFinalizedWithdrawRequest is a log parse operation binding the contract event 0x12a00f5e4c3614409f2dd90dc5be91b9b64ef89bac58a5b034ec0094376dbd37.
//
// Solidity: event FinalizedWithdrawRequest(uint256 nextRequestIdToFinalize)
func (_SDUtilityPool *SDUtilityPoolFilterer) ParseFinalizedWithdrawRequest(log types.Log) (*SDUtilityPoolFinalizedWithdrawRequest, error) {
	event := new(SDUtilityPoolFinalizedWithdrawRequest)
	if err := _SDUtilityPool.contract.UnpackLog(event, "FinalizedWithdrawRequest", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SDUtilityPoolProtocolFeeFactorUpdatedIterator is returned from FilterProtocolFeeFactorUpdated and is used to iterate over the raw logs and unpacked data for ProtocolFeeFactorUpdated events raised by the SDUtilityPool contract.
type SDUtilityPoolProtocolFeeFactorUpdatedIterator struct {
	Event *SDUtilityPoolProtocolFeeFactorUpdated // Event containing the contract specifics and raw log

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
func (it *SDUtilityPoolProtocolFeeFactorUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SDUtilityPoolProtocolFeeFactorUpdated)
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
		it.Event = new(SDUtilityPoolProtocolFeeFactorUpdated)
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
func (it *SDUtilityPoolProtocolFeeFactorUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SDUtilityPoolProtocolFeeFactorUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SDUtilityPoolProtocolFeeFactorUpdated represents a ProtocolFeeFactorUpdated event raised by the SDUtilityPool contract.
type SDUtilityPoolProtocolFeeFactorUpdated struct {
	ProtocolFeeFactor *big.Int
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterProtocolFeeFactorUpdated is a free log retrieval operation binding the contract event 0xba813ee7ea736ec5148f515c7fe651c522fa84413c6c5ce693bd74abade775d3.
//
// Solidity: event ProtocolFeeFactorUpdated(uint256 protocolFeeFactor)
func (_SDUtilityPool *SDUtilityPoolFilterer) FilterProtocolFeeFactorUpdated(opts *bind.FilterOpts) (*SDUtilityPoolProtocolFeeFactorUpdatedIterator, error) {

	logs, sub, err := _SDUtilityPool.contract.FilterLogs(opts, "ProtocolFeeFactorUpdated")
	if err != nil {
		return nil, err
	}
	return &SDUtilityPoolProtocolFeeFactorUpdatedIterator{contract: _SDUtilityPool.contract, event: "ProtocolFeeFactorUpdated", logs: logs, sub: sub}, nil
}

// WatchProtocolFeeFactorUpdated is a free log subscription operation binding the contract event 0xba813ee7ea736ec5148f515c7fe651c522fa84413c6c5ce693bd74abade775d3.
//
// Solidity: event ProtocolFeeFactorUpdated(uint256 protocolFeeFactor)
func (_SDUtilityPool *SDUtilityPoolFilterer) WatchProtocolFeeFactorUpdated(opts *bind.WatchOpts, sink chan<- *SDUtilityPoolProtocolFeeFactorUpdated) (event.Subscription, error) {

	logs, sub, err := _SDUtilityPool.contract.WatchLogs(opts, "ProtocolFeeFactorUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SDUtilityPoolProtocolFeeFactorUpdated)
				if err := _SDUtilityPool.contract.UnpackLog(event, "ProtocolFeeFactorUpdated", log); err != nil {
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

// ParseProtocolFeeFactorUpdated is a log parse operation binding the contract event 0xba813ee7ea736ec5148f515c7fe651c522fa84413c6c5ce693bd74abade775d3.
//
// Solidity: event ProtocolFeeFactorUpdated(uint256 protocolFeeFactor)
func (_SDUtilityPool *SDUtilityPoolFilterer) ParseProtocolFeeFactorUpdated(log types.Log) (*SDUtilityPoolProtocolFeeFactorUpdated, error) {
	event := new(SDUtilityPoolProtocolFeeFactorUpdated)
	if err := _SDUtilityPool.contract.UnpackLog(event, "ProtocolFeeFactorUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SDUtilityPoolRedeemedIterator is returned from FilterRedeemed and is used to iterate over the raw logs and unpacked data for Redeemed events raised by the SDUtilityPool contract.
type SDUtilityPoolRedeemedIterator struct {
	Event *SDUtilityPoolRedeemed // Event containing the contract specifics and raw log

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
func (it *SDUtilityPoolRedeemedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SDUtilityPoolRedeemed)
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
		it.Event = new(SDUtilityPoolRedeemed)
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
func (it *SDUtilityPoolRedeemedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SDUtilityPoolRedeemedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SDUtilityPoolRedeemed represents a Redeemed event raised by the SDUtilityPool contract.
type SDUtilityPoolRedeemed struct {
	Delegator common.Address
	SdAmount  *big.Int
	SdXAmount *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterRedeemed is a free log retrieval operation binding the contract event 0xf3a670cd3af7d64b488926880889d08a8585a138ff455227af6737339a1ec262.
//
// Solidity: event Redeemed(address indexed delegator, uint256 sdAmount, uint256 sdXAmount)
func (_SDUtilityPool *SDUtilityPoolFilterer) FilterRedeemed(opts *bind.FilterOpts, delegator []common.Address) (*SDUtilityPoolRedeemedIterator, error) {

	var delegatorRule []interface{}
	for _, delegatorItem := range delegator {
		delegatorRule = append(delegatorRule, delegatorItem)
	}

	logs, sub, err := _SDUtilityPool.contract.FilterLogs(opts, "Redeemed", delegatorRule)
	if err != nil {
		return nil, err
	}
	return &SDUtilityPoolRedeemedIterator{contract: _SDUtilityPool.contract, event: "Redeemed", logs: logs, sub: sub}, nil
}

// WatchRedeemed is a free log subscription operation binding the contract event 0xf3a670cd3af7d64b488926880889d08a8585a138ff455227af6737339a1ec262.
//
// Solidity: event Redeemed(address indexed delegator, uint256 sdAmount, uint256 sdXAmount)
func (_SDUtilityPool *SDUtilityPoolFilterer) WatchRedeemed(opts *bind.WatchOpts, sink chan<- *SDUtilityPoolRedeemed, delegator []common.Address) (event.Subscription, error) {

	var delegatorRule []interface{}
	for _, delegatorItem := range delegator {
		delegatorRule = append(delegatorRule, delegatorItem)
	}

	logs, sub, err := _SDUtilityPool.contract.WatchLogs(opts, "Redeemed", delegatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SDUtilityPoolRedeemed)
				if err := _SDUtilityPool.contract.UnpackLog(event, "Redeemed", log); err != nil {
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

// ParseRedeemed is a log parse operation binding the contract event 0xf3a670cd3af7d64b488926880889d08a8585a138ff455227af6737339a1ec262.
//
// Solidity: event Redeemed(address indexed delegator, uint256 sdAmount, uint256 sdXAmount)
func (_SDUtilityPool *SDUtilityPoolFilterer) ParseRedeemed(log types.Log) (*SDUtilityPoolRedeemed, error) {
	event := new(SDUtilityPoolRedeemed)
	if err := _SDUtilityPool.contract.UnpackLog(event, "Redeemed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SDUtilityPoolRepaidIterator is returned from FilterRepaid and is used to iterate over the raw logs and unpacked data for Repaid events raised by the SDUtilityPool contract.
type SDUtilityPoolRepaidIterator struct {
	Event *SDUtilityPoolRepaid // Event containing the contract specifics and raw log

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
func (it *SDUtilityPoolRepaidIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SDUtilityPoolRepaid)
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
		it.Event = new(SDUtilityPoolRepaid)
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
func (it *SDUtilityPoolRepaidIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SDUtilityPoolRepaidIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SDUtilityPoolRepaid represents a Repaid event raised by the SDUtilityPool contract.
type SDUtilityPoolRepaid struct {
	Utilizer    common.Address
	RepayAmount *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterRepaid is a free log retrieval operation binding the contract event 0x0516911bcc3a0a7412a44601057c0a0a1ec628bde049a84284bc428866534488.
//
// Solidity: event Repaid(address indexed utilizer, uint256 repayAmount)
func (_SDUtilityPool *SDUtilityPoolFilterer) FilterRepaid(opts *bind.FilterOpts, utilizer []common.Address) (*SDUtilityPoolRepaidIterator, error) {

	var utilizerRule []interface{}
	for _, utilizerItem := range utilizer {
		utilizerRule = append(utilizerRule, utilizerItem)
	}

	logs, sub, err := _SDUtilityPool.contract.FilterLogs(opts, "Repaid", utilizerRule)
	if err != nil {
		return nil, err
	}
	return &SDUtilityPoolRepaidIterator{contract: _SDUtilityPool.contract, event: "Repaid", logs: logs, sub: sub}, nil
}

// WatchRepaid is a free log subscription operation binding the contract event 0x0516911bcc3a0a7412a44601057c0a0a1ec628bde049a84284bc428866534488.
//
// Solidity: event Repaid(address indexed utilizer, uint256 repayAmount)
func (_SDUtilityPool *SDUtilityPoolFilterer) WatchRepaid(opts *bind.WatchOpts, sink chan<- *SDUtilityPoolRepaid, utilizer []common.Address) (event.Subscription, error) {

	var utilizerRule []interface{}
	for _, utilizerItem := range utilizer {
		utilizerRule = append(utilizerRule, utilizerItem)
	}

	logs, sub, err := _SDUtilityPool.contract.WatchLogs(opts, "Repaid", utilizerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SDUtilityPoolRepaid)
				if err := _SDUtilityPool.contract.UnpackLog(event, "Repaid", log); err != nil {
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

// ParseRepaid is a log parse operation binding the contract event 0x0516911bcc3a0a7412a44601057c0a0a1ec628bde049a84284bc428866534488.
//
// Solidity: event Repaid(address indexed utilizer, uint256 repayAmount)
func (_SDUtilityPool *SDUtilityPoolFilterer) ParseRepaid(log types.Log) (*SDUtilityPoolRepaid, error) {
	event := new(SDUtilityPoolRepaid)
	if err := _SDUtilityPool.contract.UnpackLog(event, "Repaid", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SDUtilityPoolRequestRedeemedIterator is returned from FilterRequestRedeemed and is used to iterate over the raw logs and unpacked data for RequestRedeemed events raised by the SDUtilityPool contract.
type SDUtilityPoolRequestRedeemedIterator struct {
	Event *SDUtilityPoolRequestRedeemed // Event containing the contract specifics and raw log

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
func (it *SDUtilityPoolRequestRedeemedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SDUtilityPoolRequestRedeemed)
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
		it.Event = new(SDUtilityPoolRequestRedeemed)
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
func (it *SDUtilityPoolRequestRedeemedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SDUtilityPoolRequestRedeemedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SDUtilityPoolRequestRedeemed represents a RequestRedeemed event raised by the SDUtilityPool contract.
type SDUtilityPoolRequestRedeemed struct {
	Caller       common.Address
	SdToTransfer *big.Int
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterRequestRedeemed is a free log retrieval operation binding the contract event 0xf7aca15eb385953329daa8e235ce29ead201e26c1e36e1083d929821401bb7d2.
//
// Solidity: event RequestRedeemed(address caller, uint256 sdToTransfer)
func (_SDUtilityPool *SDUtilityPoolFilterer) FilterRequestRedeemed(opts *bind.FilterOpts) (*SDUtilityPoolRequestRedeemedIterator, error) {

	logs, sub, err := _SDUtilityPool.contract.FilterLogs(opts, "RequestRedeemed")
	if err != nil {
		return nil, err
	}
	return &SDUtilityPoolRequestRedeemedIterator{contract: _SDUtilityPool.contract, event: "RequestRedeemed", logs: logs, sub: sub}, nil
}

// WatchRequestRedeemed is a free log subscription operation binding the contract event 0xf7aca15eb385953329daa8e235ce29ead201e26c1e36e1083d929821401bb7d2.
//
// Solidity: event RequestRedeemed(address caller, uint256 sdToTransfer)
func (_SDUtilityPool *SDUtilityPoolFilterer) WatchRequestRedeemed(opts *bind.WatchOpts, sink chan<- *SDUtilityPoolRequestRedeemed) (event.Subscription, error) {

	logs, sub, err := _SDUtilityPool.contract.WatchLogs(opts, "RequestRedeemed")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SDUtilityPoolRequestRedeemed)
				if err := _SDUtilityPool.contract.UnpackLog(event, "RequestRedeemed", log); err != nil {
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

// ParseRequestRedeemed is a log parse operation binding the contract event 0xf7aca15eb385953329daa8e235ce29ead201e26c1e36e1083d929821401bb7d2.
//
// Solidity: event RequestRedeemed(address caller, uint256 sdToTransfer)
func (_SDUtilityPool *SDUtilityPoolFilterer) ParseRequestRedeemed(log types.Log) (*SDUtilityPoolRequestRedeemed, error) {
	event := new(SDUtilityPoolRequestRedeemed)
	if err := _SDUtilityPool.contract.UnpackLog(event, "RequestRedeemed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SDUtilityPoolSDUtilizedIterator is returned from FilterSDUtilized and is used to iterate over the raw logs and unpacked data for SDUtilized events raised by the SDUtilityPool contract.
type SDUtilityPoolSDUtilizedIterator struct {
	Event *SDUtilityPoolSDUtilized // Event containing the contract specifics and raw log

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
func (it *SDUtilityPoolSDUtilizedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SDUtilityPoolSDUtilized)
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
		it.Event = new(SDUtilityPoolSDUtilized)
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
func (it *SDUtilityPoolSDUtilizedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SDUtilityPoolSDUtilizedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SDUtilityPoolSDUtilized represents a SDUtilized event raised by the SDUtilityPool contract.
type SDUtilityPoolSDUtilized struct {
	Utilizer      common.Address
	UtilizeAmount *big.Int
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterSDUtilized is a free log retrieval operation binding the contract event 0x29f5a3c94649e9fd569197488f431f12d8b1fe681596964128403511823a6386.
//
// Solidity: event SDUtilized(address utilizer, uint256 utilizeAmount)
func (_SDUtilityPool *SDUtilityPoolFilterer) FilterSDUtilized(opts *bind.FilterOpts) (*SDUtilityPoolSDUtilizedIterator, error) {

	logs, sub, err := _SDUtilityPool.contract.FilterLogs(opts, "SDUtilized")
	if err != nil {
		return nil, err
	}
	return &SDUtilityPoolSDUtilizedIterator{contract: _SDUtilityPool.contract, event: "SDUtilized", logs: logs, sub: sub}, nil
}

// WatchSDUtilized is a free log subscription operation binding the contract event 0x29f5a3c94649e9fd569197488f431f12d8b1fe681596964128403511823a6386.
//
// Solidity: event SDUtilized(address utilizer, uint256 utilizeAmount)
func (_SDUtilityPool *SDUtilityPoolFilterer) WatchSDUtilized(opts *bind.WatchOpts, sink chan<- *SDUtilityPoolSDUtilized) (event.Subscription, error) {

	logs, sub, err := _SDUtilityPool.contract.WatchLogs(opts, "SDUtilized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SDUtilityPoolSDUtilized)
				if err := _SDUtilityPool.contract.UnpackLog(event, "SDUtilized", log); err != nil {
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

// ParseSDUtilized is a log parse operation binding the contract event 0x29f5a3c94649e9fd569197488f431f12d8b1fe681596964128403511823a6386.
//
// Solidity: event SDUtilized(address utilizer, uint256 utilizeAmount)
func (_SDUtilityPool *SDUtilityPoolFilterer) ParseSDUtilized(log types.Log) (*SDUtilityPoolSDUtilized, error) {
	event := new(SDUtilityPoolSDUtilized)
	if err := _SDUtilityPool.contract.UnpackLog(event, "SDUtilized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SDUtilityPoolUpdatedFinalizationBatchLimitIterator is returned from FilterUpdatedFinalizationBatchLimit and is used to iterate over the raw logs and unpacked data for UpdatedFinalizationBatchLimit events raised by the SDUtilityPool contract.
type SDUtilityPoolUpdatedFinalizationBatchLimitIterator struct {
	Event *SDUtilityPoolUpdatedFinalizationBatchLimit // Event containing the contract specifics and raw log

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
func (it *SDUtilityPoolUpdatedFinalizationBatchLimitIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SDUtilityPoolUpdatedFinalizationBatchLimit)
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
		it.Event = new(SDUtilityPoolUpdatedFinalizationBatchLimit)
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
func (it *SDUtilityPoolUpdatedFinalizationBatchLimitIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SDUtilityPoolUpdatedFinalizationBatchLimitIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SDUtilityPoolUpdatedFinalizationBatchLimit represents a UpdatedFinalizationBatchLimit event raised by the SDUtilityPool contract.
type SDUtilityPoolUpdatedFinalizationBatchLimit struct {
	FinalizationBatchLimit *big.Int
	Raw                    types.Log // Blockchain specific contextual infos
}

// FilterUpdatedFinalizationBatchLimit is a free log retrieval operation binding the contract event 0x7ffbe87ac4b7820fd4ca4ac8c7c7820cc79c7cdd9631ac083c06bb833be63587.
//
// Solidity: event UpdatedFinalizationBatchLimit(uint256 finalizationBatchLimit)
func (_SDUtilityPool *SDUtilityPoolFilterer) FilterUpdatedFinalizationBatchLimit(opts *bind.FilterOpts) (*SDUtilityPoolUpdatedFinalizationBatchLimitIterator, error) {

	logs, sub, err := _SDUtilityPool.contract.FilterLogs(opts, "UpdatedFinalizationBatchLimit")
	if err != nil {
		return nil, err
	}
	return &SDUtilityPoolUpdatedFinalizationBatchLimitIterator{contract: _SDUtilityPool.contract, event: "UpdatedFinalizationBatchLimit", logs: logs, sub: sub}, nil
}

// WatchUpdatedFinalizationBatchLimit is a free log subscription operation binding the contract event 0x7ffbe87ac4b7820fd4ca4ac8c7c7820cc79c7cdd9631ac083c06bb833be63587.
//
// Solidity: event UpdatedFinalizationBatchLimit(uint256 finalizationBatchLimit)
func (_SDUtilityPool *SDUtilityPoolFilterer) WatchUpdatedFinalizationBatchLimit(opts *bind.WatchOpts, sink chan<- *SDUtilityPoolUpdatedFinalizationBatchLimit) (event.Subscription, error) {

	logs, sub, err := _SDUtilityPool.contract.WatchLogs(opts, "UpdatedFinalizationBatchLimit")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SDUtilityPoolUpdatedFinalizationBatchLimit)
				if err := _SDUtilityPool.contract.UnpackLog(event, "UpdatedFinalizationBatchLimit", log); err != nil {
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

// ParseUpdatedFinalizationBatchLimit is a log parse operation binding the contract event 0x7ffbe87ac4b7820fd4ca4ac8c7c7820cc79c7cdd9631ac083c06bb833be63587.
//
// Solidity: event UpdatedFinalizationBatchLimit(uint256 finalizationBatchLimit)
func (_SDUtilityPool *SDUtilityPoolFilterer) ParseUpdatedFinalizationBatchLimit(log types.Log) (*SDUtilityPoolUpdatedFinalizationBatchLimit, error) {
	event := new(SDUtilityPoolUpdatedFinalizationBatchLimit)
	if err := _SDUtilityPool.contract.UnpackLog(event, "UpdatedFinalizationBatchLimit", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SDUtilityPoolUpdatedMaxETHWorthOfSDPerValidatorIterator is returned from FilterUpdatedMaxETHWorthOfSDPerValidator and is used to iterate over the raw logs and unpacked data for UpdatedMaxETHWorthOfSDPerValidator events raised by the SDUtilityPool contract.
type SDUtilityPoolUpdatedMaxETHWorthOfSDPerValidatorIterator struct {
	Event *SDUtilityPoolUpdatedMaxETHWorthOfSDPerValidator // Event containing the contract specifics and raw log

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
func (it *SDUtilityPoolUpdatedMaxETHWorthOfSDPerValidatorIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SDUtilityPoolUpdatedMaxETHWorthOfSDPerValidator)
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
		it.Event = new(SDUtilityPoolUpdatedMaxETHWorthOfSDPerValidator)
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
func (it *SDUtilityPoolUpdatedMaxETHWorthOfSDPerValidatorIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SDUtilityPoolUpdatedMaxETHWorthOfSDPerValidatorIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SDUtilityPoolUpdatedMaxETHWorthOfSDPerValidator represents a UpdatedMaxETHWorthOfSDPerValidator event raised by the SDUtilityPool contract.
type SDUtilityPoolUpdatedMaxETHWorthOfSDPerValidator struct {
	MaxETHWorthOfSDPerValidator *big.Int
	Raw                         types.Log // Blockchain specific contextual infos
}

// FilterUpdatedMaxETHWorthOfSDPerValidator is a free log retrieval operation binding the contract event 0x9491876f555a700a2938143a5a27937708ab194a6f8d94cd7eb8eefae92f34f5.
//
// Solidity: event UpdatedMaxETHWorthOfSDPerValidator(uint256 maxETHWorthOfSDPerValidator)
func (_SDUtilityPool *SDUtilityPoolFilterer) FilterUpdatedMaxETHWorthOfSDPerValidator(opts *bind.FilterOpts) (*SDUtilityPoolUpdatedMaxETHWorthOfSDPerValidatorIterator, error) {

	logs, sub, err := _SDUtilityPool.contract.FilterLogs(opts, "UpdatedMaxETHWorthOfSDPerValidator")
	if err != nil {
		return nil, err
	}
	return &SDUtilityPoolUpdatedMaxETHWorthOfSDPerValidatorIterator{contract: _SDUtilityPool.contract, event: "UpdatedMaxETHWorthOfSDPerValidator", logs: logs, sub: sub}, nil
}

// WatchUpdatedMaxETHWorthOfSDPerValidator is a free log subscription operation binding the contract event 0x9491876f555a700a2938143a5a27937708ab194a6f8d94cd7eb8eefae92f34f5.
//
// Solidity: event UpdatedMaxETHWorthOfSDPerValidator(uint256 maxETHWorthOfSDPerValidator)
func (_SDUtilityPool *SDUtilityPoolFilterer) WatchUpdatedMaxETHWorthOfSDPerValidator(opts *bind.WatchOpts, sink chan<- *SDUtilityPoolUpdatedMaxETHWorthOfSDPerValidator) (event.Subscription, error) {

	logs, sub, err := _SDUtilityPool.contract.WatchLogs(opts, "UpdatedMaxETHWorthOfSDPerValidator")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SDUtilityPoolUpdatedMaxETHWorthOfSDPerValidator)
				if err := _SDUtilityPool.contract.UnpackLog(event, "UpdatedMaxETHWorthOfSDPerValidator", log); err != nil {
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

// ParseUpdatedMaxETHWorthOfSDPerValidator is a log parse operation binding the contract event 0x9491876f555a700a2938143a5a27937708ab194a6f8d94cd7eb8eefae92f34f5.
//
// Solidity: event UpdatedMaxETHWorthOfSDPerValidator(uint256 maxETHWorthOfSDPerValidator)
func (_SDUtilityPool *SDUtilityPoolFilterer) ParseUpdatedMaxETHWorthOfSDPerValidator(log types.Log) (*SDUtilityPoolUpdatedMaxETHWorthOfSDPerValidator, error) {
	event := new(SDUtilityPoolUpdatedMaxETHWorthOfSDPerValidator)
	if err := _SDUtilityPool.contract.UnpackLog(event, "UpdatedMaxETHWorthOfSDPerValidator", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SDUtilityPoolUpdatedMaxNonRedeemedDelegatorRequestCountIterator is returned from FilterUpdatedMaxNonRedeemedDelegatorRequestCount and is used to iterate over the raw logs and unpacked data for UpdatedMaxNonRedeemedDelegatorRequestCount events raised by the SDUtilityPool contract.
type SDUtilityPoolUpdatedMaxNonRedeemedDelegatorRequestCountIterator struct {
	Event *SDUtilityPoolUpdatedMaxNonRedeemedDelegatorRequestCount // Event containing the contract specifics and raw log

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
func (it *SDUtilityPoolUpdatedMaxNonRedeemedDelegatorRequestCountIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SDUtilityPoolUpdatedMaxNonRedeemedDelegatorRequestCount)
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
		it.Event = new(SDUtilityPoolUpdatedMaxNonRedeemedDelegatorRequestCount)
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
func (it *SDUtilityPoolUpdatedMaxNonRedeemedDelegatorRequestCountIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SDUtilityPoolUpdatedMaxNonRedeemedDelegatorRequestCountIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SDUtilityPoolUpdatedMaxNonRedeemedDelegatorRequestCount represents a UpdatedMaxNonRedeemedDelegatorRequestCount event raised by the SDUtilityPool contract.
type SDUtilityPoolUpdatedMaxNonRedeemedDelegatorRequestCount struct {
	Count *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterUpdatedMaxNonRedeemedDelegatorRequestCount is a free log retrieval operation binding the contract event 0x620864df809b46dfaffcfb35289f5efd1abd7c309562a0bdb89ee19903b2ed10.
//
// Solidity: event UpdatedMaxNonRedeemedDelegatorRequestCount(uint256 count)
func (_SDUtilityPool *SDUtilityPoolFilterer) FilterUpdatedMaxNonRedeemedDelegatorRequestCount(opts *bind.FilterOpts) (*SDUtilityPoolUpdatedMaxNonRedeemedDelegatorRequestCountIterator, error) {

	logs, sub, err := _SDUtilityPool.contract.FilterLogs(opts, "UpdatedMaxNonRedeemedDelegatorRequestCount")
	if err != nil {
		return nil, err
	}
	return &SDUtilityPoolUpdatedMaxNonRedeemedDelegatorRequestCountIterator{contract: _SDUtilityPool.contract, event: "UpdatedMaxNonRedeemedDelegatorRequestCount", logs: logs, sub: sub}, nil
}

// WatchUpdatedMaxNonRedeemedDelegatorRequestCount is a free log subscription operation binding the contract event 0x620864df809b46dfaffcfb35289f5efd1abd7c309562a0bdb89ee19903b2ed10.
//
// Solidity: event UpdatedMaxNonRedeemedDelegatorRequestCount(uint256 count)
func (_SDUtilityPool *SDUtilityPoolFilterer) WatchUpdatedMaxNonRedeemedDelegatorRequestCount(opts *bind.WatchOpts, sink chan<- *SDUtilityPoolUpdatedMaxNonRedeemedDelegatorRequestCount) (event.Subscription, error) {

	logs, sub, err := _SDUtilityPool.contract.WatchLogs(opts, "UpdatedMaxNonRedeemedDelegatorRequestCount")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SDUtilityPoolUpdatedMaxNonRedeemedDelegatorRequestCount)
				if err := _SDUtilityPool.contract.UnpackLog(event, "UpdatedMaxNonRedeemedDelegatorRequestCount", log); err != nil {
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

// ParseUpdatedMaxNonRedeemedDelegatorRequestCount is a log parse operation binding the contract event 0x620864df809b46dfaffcfb35289f5efd1abd7c309562a0bdb89ee19903b2ed10.
//
// Solidity: event UpdatedMaxNonRedeemedDelegatorRequestCount(uint256 count)
func (_SDUtilityPool *SDUtilityPoolFilterer) ParseUpdatedMaxNonRedeemedDelegatorRequestCount(log types.Log) (*SDUtilityPoolUpdatedMaxNonRedeemedDelegatorRequestCount, error) {
	event := new(SDUtilityPoolUpdatedMaxNonRedeemedDelegatorRequestCount)
	if err := _SDUtilityPool.contract.UnpackLog(event, "UpdatedMaxNonRedeemedDelegatorRequestCount", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SDUtilityPoolUpdatedMinBlockDelayToFinalizeRequestIterator is returned from FilterUpdatedMinBlockDelayToFinalizeRequest and is used to iterate over the raw logs and unpacked data for UpdatedMinBlockDelayToFinalizeRequest events raised by the SDUtilityPool contract.
type SDUtilityPoolUpdatedMinBlockDelayToFinalizeRequestIterator struct {
	Event *SDUtilityPoolUpdatedMinBlockDelayToFinalizeRequest // Event containing the contract specifics and raw log

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
func (it *SDUtilityPoolUpdatedMinBlockDelayToFinalizeRequestIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SDUtilityPoolUpdatedMinBlockDelayToFinalizeRequest)
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
		it.Event = new(SDUtilityPoolUpdatedMinBlockDelayToFinalizeRequest)
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
func (it *SDUtilityPoolUpdatedMinBlockDelayToFinalizeRequestIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SDUtilityPoolUpdatedMinBlockDelayToFinalizeRequestIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SDUtilityPoolUpdatedMinBlockDelayToFinalizeRequest represents a UpdatedMinBlockDelayToFinalizeRequest event raised by the SDUtilityPool contract.
type SDUtilityPoolUpdatedMinBlockDelayToFinalizeRequest struct {
	MinBlockDelayToFinalizeRequest *big.Int
	Raw                            types.Log // Blockchain specific contextual infos
}

// FilterUpdatedMinBlockDelayToFinalizeRequest is a free log retrieval operation binding the contract event 0xaa3bf534da453100a74c0c499340ebed87ce7f16483706a8b1e5ca11b9982789.
//
// Solidity: event UpdatedMinBlockDelayToFinalizeRequest(uint256 minBlockDelayToFinalizeRequest)
func (_SDUtilityPool *SDUtilityPoolFilterer) FilterUpdatedMinBlockDelayToFinalizeRequest(opts *bind.FilterOpts) (*SDUtilityPoolUpdatedMinBlockDelayToFinalizeRequestIterator, error) {

	logs, sub, err := _SDUtilityPool.contract.FilterLogs(opts, "UpdatedMinBlockDelayToFinalizeRequest")
	if err != nil {
		return nil, err
	}
	return &SDUtilityPoolUpdatedMinBlockDelayToFinalizeRequestIterator{contract: _SDUtilityPool.contract, event: "UpdatedMinBlockDelayToFinalizeRequest", logs: logs, sub: sub}, nil
}

// WatchUpdatedMinBlockDelayToFinalizeRequest is a free log subscription operation binding the contract event 0xaa3bf534da453100a74c0c499340ebed87ce7f16483706a8b1e5ca11b9982789.
//
// Solidity: event UpdatedMinBlockDelayToFinalizeRequest(uint256 minBlockDelayToFinalizeRequest)
func (_SDUtilityPool *SDUtilityPoolFilterer) WatchUpdatedMinBlockDelayToFinalizeRequest(opts *bind.WatchOpts, sink chan<- *SDUtilityPoolUpdatedMinBlockDelayToFinalizeRequest) (event.Subscription, error) {

	logs, sub, err := _SDUtilityPool.contract.WatchLogs(opts, "UpdatedMinBlockDelayToFinalizeRequest")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SDUtilityPoolUpdatedMinBlockDelayToFinalizeRequest)
				if err := _SDUtilityPool.contract.UnpackLog(event, "UpdatedMinBlockDelayToFinalizeRequest", log); err != nil {
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

// ParseUpdatedMinBlockDelayToFinalizeRequest is a log parse operation binding the contract event 0xaa3bf534da453100a74c0c499340ebed87ce7f16483706a8b1e5ca11b9982789.
//
// Solidity: event UpdatedMinBlockDelayToFinalizeRequest(uint256 minBlockDelayToFinalizeRequest)
func (_SDUtilityPool *SDUtilityPoolFilterer) ParseUpdatedMinBlockDelayToFinalizeRequest(log types.Log) (*SDUtilityPoolUpdatedMinBlockDelayToFinalizeRequest, error) {
	event := new(SDUtilityPoolUpdatedMinBlockDelayToFinalizeRequest)
	if err := _SDUtilityPool.contract.UnpackLog(event, "UpdatedMinBlockDelayToFinalizeRequest", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SDUtilityPoolUpdatedStaderConfigIterator is returned from FilterUpdatedStaderConfig and is used to iterate over the raw logs and unpacked data for UpdatedStaderConfig events raised by the SDUtilityPool contract.
type SDUtilityPoolUpdatedStaderConfigIterator struct {
	Event *SDUtilityPoolUpdatedStaderConfig // Event containing the contract specifics and raw log

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
func (it *SDUtilityPoolUpdatedStaderConfigIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SDUtilityPoolUpdatedStaderConfig)
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
		it.Event = new(SDUtilityPoolUpdatedStaderConfig)
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
func (it *SDUtilityPoolUpdatedStaderConfigIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SDUtilityPoolUpdatedStaderConfigIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SDUtilityPoolUpdatedStaderConfig represents a UpdatedStaderConfig event raised by the SDUtilityPool contract.
type SDUtilityPoolUpdatedStaderConfig struct {
	StaderConfig common.Address
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterUpdatedStaderConfig is a free log retrieval operation binding the contract event 0xdb2219043d7b197cb235f1af0cf6d782d77dee3de19e3f4fb6d39aae633b4485.
//
// Solidity: event UpdatedStaderConfig(address indexed _staderConfig)
func (_SDUtilityPool *SDUtilityPoolFilterer) FilterUpdatedStaderConfig(opts *bind.FilterOpts, _staderConfig []common.Address) (*SDUtilityPoolUpdatedStaderConfigIterator, error) {

	var _staderConfigRule []interface{}
	for _, _staderConfigItem := range _staderConfig {
		_staderConfigRule = append(_staderConfigRule, _staderConfigItem)
	}

	logs, sub, err := _SDUtilityPool.contract.FilterLogs(opts, "UpdatedStaderConfig", _staderConfigRule)
	if err != nil {
		return nil, err
	}
	return &SDUtilityPoolUpdatedStaderConfigIterator{contract: _SDUtilityPool.contract, event: "UpdatedStaderConfig", logs: logs, sub: sub}, nil
}

// WatchUpdatedStaderConfig is a free log subscription operation binding the contract event 0xdb2219043d7b197cb235f1af0cf6d782d77dee3de19e3f4fb6d39aae633b4485.
//
// Solidity: event UpdatedStaderConfig(address indexed _staderConfig)
func (_SDUtilityPool *SDUtilityPoolFilterer) WatchUpdatedStaderConfig(opts *bind.WatchOpts, sink chan<- *SDUtilityPoolUpdatedStaderConfig, _staderConfig []common.Address) (event.Subscription, error) {

	var _staderConfigRule []interface{}
	for _, _staderConfigItem := range _staderConfig {
		_staderConfigRule = append(_staderConfigRule, _staderConfigItem)
	}

	logs, sub, err := _SDUtilityPool.contract.WatchLogs(opts, "UpdatedStaderConfig", _staderConfigRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SDUtilityPoolUpdatedStaderConfig)
				if err := _SDUtilityPool.contract.UnpackLog(event, "UpdatedStaderConfig", log); err != nil {
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
// Solidity: event UpdatedStaderConfig(address indexed _staderConfig)
func (_SDUtilityPool *SDUtilityPoolFilterer) ParseUpdatedStaderConfig(log types.Log) (*SDUtilityPoolUpdatedStaderConfig, error) {
	event := new(SDUtilityPoolUpdatedStaderConfig)
	if err := _SDUtilityPool.contract.UnpackLog(event, "UpdatedStaderConfig", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SDUtilityPoolUpdatedUndelegationPeriodInBlocksIterator is returned from FilterUpdatedUndelegationPeriodInBlocks and is used to iterate over the raw logs and unpacked data for UpdatedUndelegationPeriodInBlocks events raised by the SDUtilityPool contract.
type SDUtilityPoolUpdatedUndelegationPeriodInBlocksIterator struct {
	Event *SDUtilityPoolUpdatedUndelegationPeriodInBlocks // Event containing the contract specifics and raw log

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
func (it *SDUtilityPoolUpdatedUndelegationPeriodInBlocksIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SDUtilityPoolUpdatedUndelegationPeriodInBlocks)
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
		it.Event = new(SDUtilityPoolUpdatedUndelegationPeriodInBlocks)
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
func (it *SDUtilityPoolUpdatedUndelegationPeriodInBlocksIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SDUtilityPoolUpdatedUndelegationPeriodInBlocksIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SDUtilityPoolUpdatedUndelegationPeriodInBlocks represents a UpdatedUndelegationPeriodInBlocks event raised by the SDUtilityPool contract.
type SDUtilityPoolUpdatedUndelegationPeriodInBlocks struct {
	UndelegationPeriodInBlocks *big.Int
	Raw                        types.Log // Blockchain specific contextual infos
}

// FilterUpdatedUndelegationPeriodInBlocks is a free log retrieval operation binding the contract event 0x727294fbc7a5976b29da226c439fc4cf0ec4f6a81a2d36e157af8667819aa7ac.
//
// Solidity: event UpdatedUndelegationPeriodInBlocks(uint256 undelegationPeriodInBlocks)
func (_SDUtilityPool *SDUtilityPoolFilterer) FilterUpdatedUndelegationPeriodInBlocks(opts *bind.FilterOpts) (*SDUtilityPoolUpdatedUndelegationPeriodInBlocksIterator, error) {

	logs, sub, err := _SDUtilityPool.contract.FilterLogs(opts, "UpdatedUndelegationPeriodInBlocks")
	if err != nil {
		return nil, err
	}
	return &SDUtilityPoolUpdatedUndelegationPeriodInBlocksIterator{contract: _SDUtilityPool.contract, event: "UpdatedUndelegationPeriodInBlocks", logs: logs, sub: sub}, nil
}

// WatchUpdatedUndelegationPeriodInBlocks is a free log subscription operation binding the contract event 0x727294fbc7a5976b29da226c439fc4cf0ec4f6a81a2d36e157af8667819aa7ac.
//
// Solidity: event UpdatedUndelegationPeriodInBlocks(uint256 undelegationPeriodInBlocks)
func (_SDUtilityPool *SDUtilityPoolFilterer) WatchUpdatedUndelegationPeriodInBlocks(opts *bind.WatchOpts, sink chan<- *SDUtilityPoolUpdatedUndelegationPeriodInBlocks) (event.Subscription, error) {

	logs, sub, err := _SDUtilityPool.contract.WatchLogs(opts, "UpdatedUndelegationPeriodInBlocks")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SDUtilityPoolUpdatedUndelegationPeriodInBlocks)
				if err := _SDUtilityPool.contract.UnpackLog(event, "UpdatedUndelegationPeriodInBlocks", log); err != nil {
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

// ParseUpdatedUndelegationPeriodInBlocks is a log parse operation binding the contract event 0x727294fbc7a5976b29da226c439fc4cf0ec4f6a81a2d36e157af8667819aa7ac.
//
// Solidity: event UpdatedUndelegationPeriodInBlocks(uint256 undelegationPeriodInBlocks)
func (_SDUtilityPool *SDUtilityPoolFilterer) ParseUpdatedUndelegationPeriodInBlocks(log types.Log) (*SDUtilityPoolUpdatedUndelegationPeriodInBlocks, error) {
	event := new(SDUtilityPoolUpdatedUndelegationPeriodInBlocks)
	if err := _SDUtilityPool.contract.UnpackLog(event, "UpdatedUndelegationPeriodInBlocks", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SDUtilityPoolUtilizationRatePerBlockUpdatedIterator is returned from FilterUtilizationRatePerBlockUpdated and is used to iterate over the raw logs and unpacked data for UtilizationRatePerBlockUpdated events raised by the SDUtilityPool contract.
type SDUtilityPoolUtilizationRatePerBlockUpdatedIterator struct {
	Event *SDUtilityPoolUtilizationRatePerBlockUpdated // Event containing the contract specifics and raw log

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
func (it *SDUtilityPoolUtilizationRatePerBlockUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SDUtilityPoolUtilizationRatePerBlockUpdated)
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
		it.Event = new(SDUtilityPoolUtilizationRatePerBlockUpdated)
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
func (it *SDUtilityPoolUtilizationRatePerBlockUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SDUtilityPoolUtilizationRatePerBlockUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SDUtilityPoolUtilizationRatePerBlockUpdated represents a UtilizationRatePerBlockUpdated event raised by the SDUtilityPool contract.
type SDUtilityPoolUtilizationRatePerBlockUpdated struct {
	UtilizationRatePerBlock *big.Int
	Raw                     types.Log // Blockchain specific contextual infos
}

// FilterUtilizationRatePerBlockUpdated is a free log retrieval operation binding the contract event 0x3d9659ac5decde6b265b661cde27cd4a357992c35cfa014eb789ad7cbe89ff8b.
//
// Solidity: event UtilizationRatePerBlockUpdated(uint256 utilizationRatePerBlock)
func (_SDUtilityPool *SDUtilityPoolFilterer) FilterUtilizationRatePerBlockUpdated(opts *bind.FilterOpts) (*SDUtilityPoolUtilizationRatePerBlockUpdatedIterator, error) {

	logs, sub, err := _SDUtilityPool.contract.FilterLogs(opts, "UtilizationRatePerBlockUpdated")
	if err != nil {
		return nil, err
	}
	return &SDUtilityPoolUtilizationRatePerBlockUpdatedIterator{contract: _SDUtilityPool.contract, event: "UtilizationRatePerBlockUpdated", logs: logs, sub: sub}, nil
}

// WatchUtilizationRatePerBlockUpdated is a free log subscription operation binding the contract event 0x3d9659ac5decde6b265b661cde27cd4a357992c35cfa014eb789ad7cbe89ff8b.
//
// Solidity: event UtilizationRatePerBlockUpdated(uint256 utilizationRatePerBlock)
func (_SDUtilityPool *SDUtilityPoolFilterer) WatchUtilizationRatePerBlockUpdated(opts *bind.WatchOpts, sink chan<- *SDUtilityPoolUtilizationRatePerBlockUpdated) (event.Subscription, error) {

	logs, sub, err := _SDUtilityPool.contract.WatchLogs(opts, "UtilizationRatePerBlockUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SDUtilityPoolUtilizationRatePerBlockUpdated)
				if err := _SDUtilityPool.contract.UnpackLog(event, "UtilizationRatePerBlockUpdated", log); err != nil {
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

// ParseUtilizationRatePerBlockUpdated is a log parse operation binding the contract event 0x3d9659ac5decde6b265b661cde27cd4a357992c35cfa014eb789ad7cbe89ff8b.
//
// Solidity: event UtilizationRatePerBlockUpdated(uint256 utilizationRatePerBlock)
func (_SDUtilityPool *SDUtilityPoolFilterer) ParseUtilizationRatePerBlockUpdated(log types.Log) (*SDUtilityPoolUtilizationRatePerBlockUpdated, error) {
	event := new(SDUtilityPoolUtilizationRatePerBlockUpdated)
	if err := _SDUtilityPool.contract.UnpackLog(event, "UtilizationRatePerBlockUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SDUtilityPoolUtilizerSDSlashingHandledIterator is returned from FilterUtilizerSDSlashingHandled and is used to iterate over the raw logs and unpacked data for UtilizerSDSlashingHandled events raised by the SDUtilityPool contract.
type SDUtilityPoolUtilizerSDSlashingHandledIterator struct {
	Event *SDUtilityPoolUtilizerSDSlashingHandled // Event containing the contract specifics and raw log

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
func (it *SDUtilityPoolUtilizerSDSlashingHandledIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SDUtilityPoolUtilizerSDSlashingHandled)
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
		it.Event = new(SDUtilityPoolUtilizerSDSlashingHandled)
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
func (it *SDUtilityPoolUtilizerSDSlashingHandledIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SDUtilityPoolUtilizerSDSlashingHandledIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SDUtilityPoolUtilizerSDSlashingHandled represents a UtilizerSDSlashingHandled event raised by the SDUtilityPool contract.
type SDUtilityPoolUtilizerSDSlashingHandled struct {
	Utilizer      common.Address
	SlashSDAmount *big.Int
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterUtilizerSDSlashingHandled is a free log retrieval operation binding the contract event 0xf583a13dc98ab9935d418df5efc1b6e0b3b4f642ccc3a8601aec4646e965dc0b.
//
// Solidity: event UtilizerSDSlashingHandled(address utilizer, uint256 slashSDAmount)
func (_SDUtilityPool *SDUtilityPoolFilterer) FilterUtilizerSDSlashingHandled(opts *bind.FilterOpts) (*SDUtilityPoolUtilizerSDSlashingHandledIterator, error) {

	logs, sub, err := _SDUtilityPool.contract.FilterLogs(opts, "UtilizerSDSlashingHandled")
	if err != nil {
		return nil, err
	}
	return &SDUtilityPoolUtilizerSDSlashingHandledIterator{contract: _SDUtilityPool.contract, event: "UtilizerSDSlashingHandled", logs: logs, sub: sub}, nil
}

// WatchUtilizerSDSlashingHandled is a free log subscription operation binding the contract event 0xf583a13dc98ab9935d418df5efc1b6e0b3b4f642ccc3a8601aec4646e965dc0b.
//
// Solidity: event UtilizerSDSlashingHandled(address utilizer, uint256 slashSDAmount)
func (_SDUtilityPool *SDUtilityPoolFilterer) WatchUtilizerSDSlashingHandled(opts *bind.WatchOpts, sink chan<- *SDUtilityPoolUtilizerSDSlashingHandled) (event.Subscription, error) {

	logs, sub, err := _SDUtilityPool.contract.WatchLogs(opts, "UtilizerSDSlashingHandled")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SDUtilityPoolUtilizerSDSlashingHandled)
				if err := _SDUtilityPool.contract.UnpackLog(event, "UtilizerSDSlashingHandled", log); err != nil {
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

// ParseUtilizerSDSlashingHandled is a log parse operation binding the contract event 0xf583a13dc98ab9935d418df5efc1b6e0b3b4f642ccc3a8601aec4646e965dc0b.
//
// Solidity: event UtilizerSDSlashingHandled(address utilizer, uint256 slashSDAmount)
func (_SDUtilityPool *SDUtilityPoolFilterer) ParseUtilizerSDSlashingHandled(log types.Log) (*SDUtilityPoolUtilizerSDSlashingHandled, error) {
	event := new(SDUtilityPoolUtilizerSDSlashingHandled)
	if err := _SDUtilityPool.contract.UnpackLog(event, "UtilizerSDSlashingHandled", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SDUtilityPoolWithdrawRequestReceivedIterator is returned from FilterWithdrawRequestReceived and is used to iterate over the raw logs and unpacked data for WithdrawRequestReceived events raised by the SDUtilityPool contract.
type SDUtilityPoolWithdrawRequestReceivedIterator struct {
	Event *SDUtilityPoolWithdrawRequestReceived // Event containing the contract specifics and raw log

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
func (it *SDUtilityPoolWithdrawRequestReceivedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SDUtilityPoolWithdrawRequestReceived)
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
		it.Event = new(SDUtilityPoolWithdrawRequestReceived)
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
func (it *SDUtilityPoolWithdrawRequestReceivedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SDUtilityPoolWithdrawRequestReceivedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SDUtilityPoolWithdrawRequestReceived represents a WithdrawRequestReceived event raised by the SDUtilityPool contract.
type SDUtilityPoolWithdrawRequestReceived struct {
	Caller             common.Address
	NextRequestId      *big.Int
	SdAmountToWithdraw *big.Int
	Raw                types.Log // Blockchain specific contextual infos
}

// FilterWithdrawRequestReceived is a free log retrieval operation binding the contract event 0x0edfc24f4f80277416f78f699d4733f7bb58fd6fb8838e2b1033162cee5fd7aa.
//
// Solidity: event WithdrawRequestReceived(address caller, uint256 nextRequestId, uint256 sdAmountToWithdraw)
func (_SDUtilityPool *SDUtilityPoolFilterer) FilterWithdrawRequestReceived(opts *bind.FilterOpts) (*SDUtilityPoolWithdrawRequestReceivedIterator, error) {

	logs, sub, err := _SDUtilityPool.contract.FilterLogs(opts, "WithdrawRequestReceived")
	if err != nil {
		return nil, err
	}
	return &SDUtilityPoolWithdrawRequestReceivedIterator{contract: _SDUtilityPool.contract, event: "WithdrawRequestReceived", logs: logs, sub: sub}, nil
}

// WatchWithdrawRequestReceived is a free log subscription operation binding the contract event 0x0edfc24f4f80277416f78f699d4733f7bb58fd6fb8838e2b1033162cee5fd7aa.
//
// Solidity: event WithdrawRequestReceived(address caller, uint256 nextRequestId, uint256 sdAmountToWithdraw)
func (_SDUtilityPool *SDUtilityPoolFilterer) WatchWithdrawRequestReceived(opts *bind.WatchOpts, sink chan<- *SDUtilityPoolWithdrawRequestReceived) (event.Subscription, error) {

	logs, sub, err := _SDUtilityPool.contract.WatchLogs(opts, "WithdrawRequestReceived")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SDUtilityPoolWithdrawRequestReceived)
				if err := _SDUtilityPool.contract.UnpackLog(event, "WithdrawRequestReceived", log); err != nil {
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

// ParseWithdrawRequestReceived is a log parse operation binding the contract event 0x0edfc24f4f80277416f78f699d4733f7bb58fd6fb8838e2b1033162cee5fd7aa.
//
// Solidity: event WithdrawRequestReceived(address caller, uint256 nextRequestId, uint256 sdAmountToWithdraw)
func (_SDUtilityPool *SDUtilityPoolFilterer) ParseWithdrawRequestReceived(log types.Log) (*SDUtilityPoolWithdrawRequestReceived, error) {
	event := new(SDUtilityPoolWithdrawRequestReceived)
	if err := _SDUtilityPool.contract.UnpackLog(event, "WithdrawRequestReceived", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
