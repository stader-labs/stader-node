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

// OperatorLiquidation is an auto generated low-level Go binding around an user-defined struct.
type OperatorLiquidation struct {
	TotalAmountInEth *big.Int
	TotalBonusInEth  *big.Int
	TotalFeeInEth    *big.Int
	IsRepaid         bool
	IsClaimed        bool
	Liquidator       common.Address
}

// UserData is an auto generated low-level Go binding around an user-defined struct.
type UserData struct {
	TotalInterestSD     *big.Int
	TotalCollateralInSD *big.Int
	HealthFactor        *big.Int
	LockedEth           *big.Int
}

// SDUtilityPoolMetaData contains all meta data concerning the SDUtilityPool contract.
var SDUtilityPoolMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"AccrualBlockNumberNotLatest\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"AlreadyClaimed\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"AlreadyLiquidated\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"CallerNotAuthorizedToRedeem\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"CallerNotManager\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"CallerNotStaderContract\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"CannotFindRequestId\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InsufficientPoolBalance\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidAmountOfWithdraw\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidInput\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidWithdrawAmount\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"MaxLimitOnWithdrawRequestCountReached\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotClaimable\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotLiquidatable\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotLiquidator\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"requestId\",\"type\":\"uint256\"}],\"name\":\"RequestIdNotFinalized\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"SDTransferFailed\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"SDUtilizeLimitReached\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"UndelegationPeriodNotPassed\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ZeroAddress\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"feeAccumulated\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"totalProtocolFee\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"totalUtilizedSD\",\"type\":\"uint256\"}],\"name\":\"AccruedFees\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"liquidator\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"liquidationBonusInEth\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"liquidationFeeInEth\",\"type\":\"uint256\"}],\"name\":\"ClaimedLiquidation\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"delegator\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"sdAmount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"sdXToMint\",\"type\":\"uint256\"}],\"name\":\"Delegated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"nextRequestIdToFinalize\",\"type\":\"uint256\"}],\"name\":\"FinalizedWithdrawRequest\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"version\",\"type\":\"uint8\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"totalLiquidationAmountInEth\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"liquidationBonusInEth\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"liquidationFeeInEth\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"liquidator\",\"type\":\"address\"}],\"name\":\"LiquidationCall\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Paused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"protocolFeeFactor\",\"type\":\"uint256\"}],\"name\":\"ProtocolFeeFactorUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"delegator\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"sdAmount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"sdXAmount\",\"type\":\"uint256\"}],\"name\":\"Redeemed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"utilizer\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"repayAmount\",\"type\":\"uint256\"}],\"name\":\"Repaid\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"caller\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"sdToTransfer\",\"type\":\"uint256\"}],\"name\":\"RequestRedeemed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"liquidationThreshold\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"liquidationBonusPercent\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"liquidationFeePercent\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"ltv\",\"type\":\"uint256\"}],\"name\":\"RiskConfigUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"previousAdminRole\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"newAdminRole\",\"type\":\"bytes32\"}],\"name\":\"RoleAdminChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleGranted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleRevoked\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"utilizer\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"utilizeAmount\",\"type\":\"uint256\"}],\"name\":\"SDUtilized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Unpaused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"finalizationBatchLimit\",\"type\":\"uint256\"}],\"name\":\"UpdatedFinalizationBatchLimit\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"maxETHWorthOfSDPerValidator\",\"type\":\"uint256\"}],\"name\":\"UpdatedMaxETHWorthOfSDPerValidator\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"count\",\"type\":\"uint256\"}],\"name\":\"UpdatedMaxNonRedeemedDelegatorRequestCount\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"minBlockDelayToFinalizeRequest\",\"type\":\"uint256\"}],\"name\":\"UpdatedMinBlockDelayToFinalizeRequest\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_staderConfig\",\"type\":\"address\"}],\"name\":\"UpdatedStaderConfig\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"undelegationPeriodInBlocks\",\"type\":\"uint256\"}],\"name\":\"UpdatedUndelegationPeriodInBlocks\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"utilizationRatePerBlock\",\"type\":\"uint256\"}],\"name\":\"UtilizationRatePerBlockUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"caller\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"nextRequestId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"sdAmountToWithdraw\",\"type\":\"uint256\"}],\"name\":\"WithdrawRequestReceived\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"WithdrawnProtocolFee\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"DECIMAL\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"DEFAULT_ADMIN_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"MAX_PROTOCOL_FEE\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"accrualBlockNumber\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"accrueFee\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"accumulatedProtocolFee\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"cTokenTotalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_requestId\",\"type\":\"uint256\"}],\"name\":\"claim\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"claimLiquidation\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"sdAmount\",\"type\":\"uint256\"}],\"name\":\"delegate\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"delegatorCTokenBalance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"delegatorWithdrawRequestedCTokenCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"delegatorWithdrawRequests\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amountOfCToken\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"sdExpected\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"sdFinalized\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"requestBlock\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"exchangeRateCurrent\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"exchangeRateStored\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"finalizationBatchLimit\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"finalizeDelegatorWithdrawalRequest\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getDelegationRate\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_delegator\",\"type\":\"address\"}],\"name\":\"getDelegatorLatestSDBalance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getLatestExchangeRate\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getLiquidationThreshold\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"getOperatorLiquidation\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"totalAmountInEth\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"totalBonusInEth\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"totalFeeInEth\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"isRepaid\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"isClaimed\",\"type\":\"bool\"},{\"internalType\":\"address\",\"name\":\"liquidator\",\"type\":\"address\"}],\"internalType\":\"structOperatorLiquidation\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"}],\"name\":\"getOperatorTotalEth\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getPoolAvailableSDBalance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_delegator\",\"type\":\"address\"}],\"name\":\"getRequestIdsByDelegator\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"\",\"type\":\"uint256[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"}],\"name\":\"getRoleAdmin\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"getUserData\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"totalInterestSD\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"totalCollateralInSD\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"healthFactor\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"lockedEth\",\"type\":\"uint256\"}],\"internalType\":\"structUserData\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_utilizer\",\"type\":\"address\"}],\"name\":\"getUtilizerLatestBalance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"grantRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"hasRole\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_admin\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_staderConfig\",\"type\":\"address\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"liquidationCall\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"liquidations\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"totalAmountInEth\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"totalBonusInEth\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"totalFeeInEth\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"isRepaid\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"isClaimed\",\"type\":\"bool\"},{\"internalType\":\"address\",\"name\":\"liquidator\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"maxApproveSD\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"maxETHWorthOfSDPerValidator\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"maxNonRedeemedDelegatorRequestCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"minBlockDelayToFinalizeRequest\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"nextRequestId\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"nextRequestIdToFinalize\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"paused\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"poolUtilization\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"protocolFeeFactor\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"renounceRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"repayAmount\",\"type\":\"uint256\"}],\"name\":\"repay\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"repaidAmount\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"repayLiquidation\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"utilizer\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"repayAmount\",\"type\":\"uint256\"}],\"name\":\"repayOnBehalf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"repaidAmount\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"requestIdsByDelegatorAddress\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_cTokenAmount\",\"type\":\"uint256\"}],\"name\":\"requestWithdraw\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"_requestId\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_sdAmount\",\"type\":\"uint256\"}],\"name\":\"requestWithdrawWithSDAmount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"_requestId\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"revokeRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"riskConfig\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"liquidationThreshold\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"liquidationBonusPercent\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"liquidationFeePercent\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"ltv\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"sdRequestedForWithdraw\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"sdReservedForClaim\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"staderConfig\",\"outputs\":[{\"internalType\":\"contractIStaderConfig\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalUtilizedSD\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"undelegationPeriodInBlocks\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"unpause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_finalizationBatchLimit\",\"type\":\"uint256\"}],\"name\":\"updateFinalizationBatchLimit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_maxETHWorthOfSDPerValidator\",\"type\":\"uint256\"}],\"name\":\"updateMaxETHWorthOfSDPerValidator\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_count\",\"type\":\"uint256\"}],\"name\":\"updateMaxNonRedeemedDelegatorRequestCount\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_minBlockDelayToFinalizeRequest\",\"type\":\"uint256\"}],\"name\":\"updateMinBlockDelayToFinalizeRequest\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_protocolFeeFactor\",\"type\":\"uint256\"}],\"name\":\"updateProtocolFeeFactor\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"liquidationThreshold\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"liquidationBonusPercent\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"liquidationFeePercent\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"ltv\",\"type\":\"uint256\"}],\"name\":\"updateRiskConfig\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_staderConfig\",\"type\":\"address\"}],\"name\":\"updateStaderConfig\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_undelegationPeriodInBlocks\",\"type\":\"uint256\"}],\"name\":\"updateUndelegationPeriodInBlocks\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_utilizationRatePerBlock\",\"type\":\"uint256\"}],\"name\":\"updateUtilizationRatePerBlock\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"utilizationRatePerBlock\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"utilizeAmount\",\"type\":\"uint256\"}],\"name\":\"utilize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"utilizeIndex\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"utilizeAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"nonTerminalKeyCount\",\"type\":\"uint256\"}],\"name\":\"utilizeWhileAddingKeys\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"utilizerBalanceCurrent\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"utilizerBalanceStored\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"utilizerData\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"principal\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"utilizeIndex\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"withdrawProtocolFee\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
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

// DECIMAL is a free data retrieval call binding the contract method 0x6ab55fd9.
//
// Solidity: function DECIMAL() view returns(uint256)
func (_SDUtilityPool *SDUtilityPoolCaller) DECIMAL(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _SDUtilityPool.contract.Call(opts, &out, "DECIMAL")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// DECIMAL is a free data retrieval call binding the contract method 0x6ab55fd9.
//
// Solidity: function DECIMAL() view returns(uint256)
func (_SDUtilityPool *SDUtilityPoolSession) DECIMAL() (*big.Int, error) {
	return _SDUtilityPool.Contract.DECIMAL(&_SDUtilityPool.CallOpts)
}

// DECIMAL is a free data retrieval call binding the contract method 0x6ab55fd9.
//
// Solidity: function DECIMAL() view returns(uint256)
func (_SDUtilityPool *SDUtilityPoolCallerSession) DECIMAL() (*big.Int, error) {
	return _SDUtilityPool.Contract.DECIMAL(&_SDUtilityPool.CallOpts)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_SDUtilityPool *SDUtilityPoolCaller) DEFAULTADMINROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _SDUtilityPool.contract.Call(opts, &out, "DEFAULT_ADMIN_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_SDUtilityPool *SDUtilityPoolSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _SDUtilityPool.Contract.DEFAULTADMINROLE(&_SDUtilityPool.CallOpts)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_SDUtilityPool *SDUtilityPoolCallerSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _SDUtilityPool.Contract.DEFAULTADMINROLE(&_SDUtilityPool.CallOpts)
}

// MAXPROTOCOLFEE is a free data retrieval call binding the contract method 0xb8ca3b83.
//
// Solidity: function MAX_PROTOCOL_FEE() view returns(uint256)
func (_SDUtilityPool *SDUtilityPoolCaller) MAXPROTOCOLFEE(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _SDUtilityPool.contract.Call(opts, &out, "MAX_PROTOCOL_FEE")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MAXPROTOCOLFEE is a free data retrieval call binding the contract method 0xb8ca3b83.
//
// Solidity: function MAX_PROTOCOL_FEE() view returns(uint256)
func (_SDUtilityPool *SDUtilityPoolSession) MAXPROTOCOLFEE() (*big.Int, error) {
	return _SDUtilityPool.Contract.MAXPROTOCOLFEE(&_SDUtilityPool.CallOpts)
}

// MAXPROTOCOLFEE is a free data retrieval call binding the contract method 0xb8ca3b83.
//
// Solidity: function MAX_PROTOCOL_FEE() view returns(uint256)
func (_SDUtilityPool *SDUtilityPoolCallerSession) MAXPROTOCOLFEE() (*big.Int, error) {
	return _SDUtilityPool.Contract.MAXPROTOCOLFEE(&_SDUtilityPool.CallOpts)
}

// AccrualBlockNumber is a free data retrieval call binding the contract method 0x6c540baf.
//
// Solidity: function accrualBlockNumber() view returns(uint256)
func (_SDUtilityPool *SDUtilityPoolCaller) AccrualBlockNumber(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _SDUtilityPool.contract.Call(opts, &out, "accrualBlockNumber")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// AccrualBlockNumber is a free data retrieval call binding the contract method 0x6c540baf.
//
// Solidity: function accrualBlockNumber() view returns(uint256)
func (_SDUtilityPool *SDUtilityPoolSession) AccrualBlockNumber() (*big.Int, error) {
	return _SDUtilityPool.Contract.AccrualBlockNumber(&_SDUtilityPool.CallOpts)
}

// AccrualBlockNumber is a free data retrieval call binding the contract method 0x6c540baf.
//
// Solidity: function accrualBlockNumber() view returns(uint256)
func (_SDUtilityPool *SDUtilityPoolCallerSession) AccrualBlockNumber() (*big.Int, error) {
	return _SDUtilityPool.Contract.AccrualBlockNumber(&_SDUtilityPool.CallOpts)
}

// AccumulatedProtocolFee is a free data retrieval call binding the contract method 0xa544a62c.
//
// Solidity: function accumulatedProtocolFee() view returns(uint256)
func (_SDUtilityPool *SDUtilityPoolCaller) AccumulatedProtocolFee(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _SDUtilityPool.contract.Call(opts, &out, "accumulatedProtocolFee")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// AccumulatedProtocolFee is a free data retrieval call binding the contract method 0xa544a62c.
//
// Solidity: function accumulatedProtocolFee() view returns(uint256)
func (_SDUtilityPool *SDUtilityPoolSession) AccumulatedProtocolFee() (*big.Int, error) {
	return _SDUtilityPool.Contract.AccumulatedProtocolFee(&_SDUtilityPool.CallOpts)
}

// AccumulatedProtocolFee is a free data retrieval call binding the contract method 0xa544a62c.
//
// Solidity: function accumulatedProtocolFee() view returns(uint256)
func (_SDUtilityPool *SDUtilityPoolCallerSession) AccumulatedProtocolFee() (*big.Int, error) {
	return _SDUtilityPool.Contract.AccumulatedProtocolFee(&_SDUtilityPool.CallOpts)
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

// DelegatorWithdrawRequestedCTokenCount is a free data retrieval call binding the contract method 0x4f8f7a37.
//
// Solidity: function delegatorWithdrawRequestedCTokenCount(address ) view returns(uint256)
func (_SDUtilityPool *SDUtilityPoolCaller) DelegatorWithdrawRequestedCTokenCount(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _SDUtilityPool.contract.Call(opts, &out, "delegatorWithdrawRequestedCTokenCount", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// DelegatorWithdrawRequestedCTokenCount is a free data retrieval call binding the contract method 0x4f8f7a37.
//
// Solidity: function delegatorWithdrawRequestedCTokenCount(address ) view returns(uint256)
func (_SDUtilityPool *SDUtilityPoolSession) DelegatorWithdrawRequestedCTokenCount(arg0 common.Address) (*big.Int, error) {
	return _SDUtilityPool.Contract.DelegatorWithdrawRequestedCTokenCount(&_SDUtilityPool.CallOpts, arg0)
}

// DelegatorWithdrawRequestedCTokenCount is a free data retrieval call binding the contract method 0x4f8f7a37.
//
// Solidity: function delegatorWithdrawRequestedCTokenCount(address ) view returns(uint256)
func (_SDUtilityPool *SDUtilityPoolCallerSession) DelegatorWithdrawRequestedCTokenCount(arg0 common.Address) (*big.Int, error) {
	return _SDUtilityPool.Contract.DelegatorWithdrawRequestedCTokenCount(&_SDUtilityPool.CallOpts, arg0)
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

// FinalizationBatchLimit is a free data retrieval call binding the contract method 0x6f2bd0b2.
//
// Solidity: function finalizationBatchLimit() view returns(uint256)
func (_SDUtilityPool *SDUtilityPoolCaller) FinalizationBatchLimit(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _SDUtilityPool.contract.Call(opts, &out, "finalizationBatchLimit")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// FinalizationBatchLimit is a free data retrieval call binding the contract method 0x6f2bd0b2.
//
// Solidity: function finalizationBatchLimit() view returns(uint256)
func (_SDUtilityPool *SDUtilityPoolSession) FinalizationBatchLimit() (*big.Int, error) {
	return _SDUtilityPool.Contract.FinalizationBatchLimit(&_SDUtilityPool.CallOpts)
}

// FinalizationBatchLimit is a free data retrieval call binding the contract method 0x6f2bd0b2.
//
// Solidity: function finalizationBatchLimit() view returns(uint256)
func (_SDUtilityPool *SDUtilityPoolCallerSession) FinalizationBatchLimit() (*big.Int, error) {
	return _SDUtilityPool.Contract.FinalizationBatchLimit(&_SDUtilityPool.CallOpts)
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

// GetLiquidationThreshold is a free data retrieval call binding the contract method 0x4ae9b8bc.
//
// Solidity: function getLiquidationThreshold() view returns(uint256)
func (_SDUtilityPool *SDUtilityPoolCaller) GetLiquidationThreshold(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _SDUtilityPool.contract.Call(opts, &out, "getLiquidationThreshold")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetLiquidationThreshold is a free data retrieval call binding the contract method 0x4ae9b8bc.
//
// Solidity: function getLiquidationThreshold() view returns(uint256)
func (_SDUtilityPool *SDUtilityPoolSession) GetLiquidationThreshold() (*big.Int, error) {
	return _SDUtilityPool.Contract.GetLiquidationThreshold(&_SDUtilityPool.CallOpts)
}

// GetLiquidationThreshold is a free data retrieval call binding the contract method 0x4ae9b8bc.
//
// Solidity: function getLiquidationThreshold() view returns(uint256)
func (_SDUtilityPool *SDUtilityPoolCallerSession) GetLiquidationThreshold() (*big.Int, error) {
	return _SDUtilityPool.Contract.GetLiquidationThreshold(&_SDUtilityPool.CallOpts)
}

// GetOperatorLiquidation is a free data retrieval call binding the contract method 0x12372ffe.
//
// Solidity: function getOperatorLiquidation(address account) view returns((uint256,uint256,uint256,bool,bool,address))
func (_SDUtilityPool *SDUtilityPoolCaller) GetOperatorLiquidation(opts *bind.CallOpts, account common.Address) (OperatorLiquidation, error) {
	var out []interface{}
	err := _SDUtilityPool.contract.Call(opts, &out, "getOperatorLiquidation", account)

	if err != nil {
		return *new(OperatorLiquidation), err
	}

	out0 := *abi.ConvertType(out[0], new(OperatorLiquidation)).(*OperatorLiquidation)

	return out0, err

}

// GetOperatorLiquidation is a free data retrieval call binding the contract method 0x12372ffe.
//
// Solidity: function getOperatorLiquidation(address account) view returns((uint256,uint256,uint256,bool,bool,address))
func (_SDUtilityPool *SDUtilityPoolSession) GetOperatorLiquidation(account common.Address) (OperatorLiquidation, error) {
	return _SDUtilityPool.Contract.GetOperatorLiquidation(&_SDUtilityPool.CallOpts, account)
}

// GetOperatorLiquidation is a free data retrieval call binding the contract method 0x12372ffe.
//
// Solidity: function getOperatorLiquidation(address account) view returns((uint256,uint256,uint256,bool,bool,address))
func (_SDUtilityPool *SDUtilityPoolCallerSession) GetOperatorLiquidation(account common.Address) (OperatorLiquidation, error) {
	return _SDUtilityPool.Contract.GetOperatorLiquidation(&_SDUtilityPool.CallOpts, account)
}

// GetOperatorTotalEth is a free data retrieval call binding the contract method 0x2170c14c.
//
// Solidity: function getOperatorTotalEth(address operator) view returns(uint256)
func (_SDUtilityPool *SDUtilityPoolCaller) GetOperatorTotalEth(opts *bind.CallOpts, operator common.Address) (*big.Int, error) {
	var out []interface{}
	err := _SDUtilityPool.contract.Call(opts, &out, "getOperatorTotalEth", operator)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetOperatorTotalEth is a free data retrieval call binding the contract method 0x2170c14c.
//
// Solidity: function getOperatorTotalEth(address operator) view returns(uint256)
func (_SDUtilityPool *SDUtilityPoolSession) GetOperatorTotalEth(operator common.Address) (*big.Int, error) {
	return _SDUtilityPool.Contract.GetOperatorTotalEth(&_SDUtilityPool.CallOpts, operator)
}

// GetOperatorTotalEth is a free data retrieval call binding the contract method 0x2170c14c.
//
// Solidity: function getOperatorTotalEth(address operator) view returns(uint256)
func (_SDUtilityPool *SDUtilityPoolCallerSession) GetOperatorTotalEth(operator common.Address) (*big.Int, error) {
	return _SDUtilityPool.Contract.GetOperatorTotalEth(&_SDUtilityPool.CallOpts, operator)
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
// Solidity: function getRequestIdsByDelegator(address _delegator) view returns(uint256[])
func (_SDUtilityPool *SDUtilityPoolCaller) GetRequestIdsByDelegator(opts *bind.CallOpts, _delegator common.Address) ([]*big.Int, error) {
	var out []interface{}
	err := _SDUtilityPool.contract.Call(opts, &out, "getRequestIdsByDelegator", _delegator)

	if err != nil {
		return *new([]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([]*big.Int)).(*[]*big.Int)

	return out0, err

}

// GetRequestIdsByDelegator is a free data retrieval call binding the contract method 0x99775f40.
//
// Solidity: function getRequestIdsByDelegator(address _delegator) view returns(uint256[])
func (_SDUtilityPool *SDUtilityPoolSession) GetRequestIdsByDelegator(_delegator common.Address) ([]*big.Int, error) {
	return _SDUtilityPool.Contract.GetRequestIdsByDelegator(&_SDUtilityPool.CallOpts, _delegator)
}

// GetRequestIdsByDelegator is a free data retrieval call binding the contract method 0x99775f40.
//
// Solidity: function getRequestIdsByDelegator(address _delegator) view returns(uint256[])
func (_SDUtilityPool *SDUtilityPoolCallerSession) GetRequestIdsByDelegator(_delegator common.Address) ([]*big.Int, error) {
	return _SDUtilityPool.Contract.GetRequestIdsByDelegator(&_SDUtilityPool.CallOpts, _delegator)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_SDUtilityPool *SDUtilityPoolCaller) GetRoleAdmin(opts *bind.CallOpts, role [32]byte) ([32]byte, error) {
	var out []interface{}
	err := _SDUtilityPool.contract.Call(opts, &out, "getRoleAdmin", role)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_SDUtilityPool *SDUtilityPoolSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _SDUtilityPool.Contract.GetRoleAdmin(&_SDUtilityPool.CallOpts, role)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_SDUtilityPool *SDUtilityPoolCallerSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _SDUtilityPool.Contract.GetRoleAdmin(&_SDUtilityPool.CallOpts, role)
}

// GetUserData is a free data retrieval call binding the contract method 0xffc9896b.
//
// Solidity: function getUserData(address account) view returns((uint256,uint256,uint256,uint256))
func (_SDUtilityPool *SDUtilityPoolCaller) GetUserData(opts *bind.CallOpts, account common.Address) (UserData, error) {
	var out []interface{}
	err := _SDUtilityPool.contract.Call(opts, &out, "getUserData", account)

	if err != nil {
		return *new(UserData), err
	}

	out0 := *abi.ConvertType(out[0], new(UserData)).(*UserData)

	return out0, err

}

// GetUserData is a free data retrieval call binding the contract method 0xffc9896b.
//
// Solidity: function getUserData(address account) view returns((uint256,uint256,uint256,uint256))
func (_SDUtilityPool *SDUtilityPoolSession) GetUserData(account common.Address) (UserData, error) {
	return _SDUtilityPool.Contract.GetUserData(&_SDUtilityPool.CallOpts, account)
}

// GetUserData is a free data retrieval call binding the contract method 0xffc9896b.
//
// Solidity: function getUserData(address account) view returns((uint256,uint256,uint256,uint256))
func (_SDUtilityPool *SDUtilityPoolCallerSession) GetUserData(account common.Address) (UserData, error) {
	return _SDUtilityPool.Contract.GetUserData(&_SDUtilityPool.CallOpts, account)
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

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_SDUtilityPool *SDUtilityPoolCaller) HasRole(opts *bind.CallOpts, role [32]byte, account common.Address) (bool, error) {
	var out []interface{}
	err := _SDUtilityPool.contract.Call(opts, &out, "hasRole", role, account)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_SDUtilityPool *SDUtilityPoolSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _SDUtilityPool.Contract.HasRole(&_SDUtilityPool.CallOpts, role, account)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_SDUtilityPool *SDUtilityPoolCallerSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _SDUtilityPool.Contract.HasRole(&_SDUtilityPool.CallOpts, role, account)
}

// Liquidations is a free data retrieval call binding the contract method 0xa66c84f4.
//
// Solidity: function liquidations(uint256 ) view returns(uint256 totalAmountInEth, uint256 totalBonusInEth, uint256 totalFeeInEth, bool isRepaid, bool isClaimed, address liquidator)
func (_SDUtilityPool *SDUtilityPoolCaller) Liquidations(opts *bind.CallOpts, arg0 *big.Int) (struct {
	TotalAmountInEth *big.Int
	TotalBonusInEth  *big.Int
	TotalFeeInEth    *big.Int
	IsRepaid         bool
	IsClaimed        bool
	Liquidator       common.Address
}, error) {
	var out []interface{}
	err := _SDUtilityPool.contract.Call(opts, &out, "liquidations", arg0)

	outstruct := new(struct {
		TotalAmountInEth *big.Int
		TotalBonusInEth  *big.Int
		TotalFeeInEth    *big.Int
		IsRepaid         bool
		IsClaimed        bool
		Liquidator       common.Address
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.TotalAmountInEth = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.TotalBonusInEth = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.TotalFeeInEth = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	outstruct.IsRepaid = *abi.ConvertType(out[3], new(bool)).(*bool)
	outstruct.IsClaimed = *abi.ConvertType(out[4], new(bool)).(*bool)
	outstruct.Liquidator = *abi.ConvertType(out[5], new(common.Address)).(*common.Address)

	return *outstruct, err

}

// Liquidations is a free data retrieval call binding the contract method 0xa66c84f4.
//
// Solidity: function liquidations(uint256 ) view returns(uint256 totalAmountInEth, uint256 totalBonusInEth, uint256 totalFeeInEth, bool isRepaid, bool isClaimed, address liquidator)
func (_SDUtilityPool *SDUtilityPoolSession) Liquidations(arg0 *big.Int) (struct {
	TotalAmountInEth *big.Int
	TotalBonusInEth  *big.Int
	TotalFeeInEth    *big.Int
	IsRepaid         bool
	IsClaimed        bool
	Liquidator       common.Address
}, error) {
	return _SDUtilityPool.Contract.Liquidations(&_SDUtilityPool.CallOpts, arg0)
}

// Liquidations is a free data retrieval call binding the contract method 0xa66c84f4.
//
// Solidity: function liquidations(uint256 ) view returns(uint256 totalAmountInEth, uint256 totalBonusInEth, uint256 totalFeeInEth, bool isRepaid, bool isClaimed, address liquidator)
func (_SDUtilityPool *SDUtilityPoolCallerSession) Liquidations(arg0 *big.Int) (struct {
	TotalAmountInEth *big.Int
	TotalBonusInEth  *big.Int
	TotalFeeInEth    *big.Int
	IsRepaid         bool
	IsClaimed        bool
	Liquidator       common.Address
}, error) {
	return _SDUtilityPool.Contract.Liquidations(&_SDUtilityPool.CallOpts, arg0)
}

// MaxETHWorthOfSDPerValidator is a free data retrieval call binding the contract method 0x2807c313.
//
// Solidity: function maxETHWorthOfSDPerValidator() view returns(uint256)
func (_SDUtilityPool *SDUtilityPoolCaller) MaxETHWorthOfSDPerValidator(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _SDUtilityPool.contract.Call(opts, &out, "maxETHWorthOfSDPerValidator")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MaxETHWorthOfSDPerValidator is a free data retrieval call binding the contract method 0x2807c313.
//
// Solidity: function maxETHWorthOfSDPerValidator() view returns(uint256)
func (_SDUtilityPool *SDUtilityPoolSession) MaxETHWorthOfSDPerValidator() (*big.Int, error) {
	return _SDUtilityPool.Contract.MaxETHWorthOfSDPerValidator(&_SDUtilityPool.CallOpts)
}

// MaxETHWorthOfSDPerValidator is a free data retrieval call binding the contract method 0x2807c313.
//
// Solidity: function maxETHWorthOfSDPerValidator() view returns(uint256)
func (_SDUtilityPool *SDUtilityPoolCallerSession) MaxETHWorthOfSDPerValidator() (*big.Int, error) {
	return _SDUtilityPool.Contract.MaxETHWorthOfSDPerValidator(&_SDUtilityPool.CallOpts)
}

// MaxNonRedeemedDelegatorRequestCount is a free data retrieval call binding the contract method 0xc083f3e6.
//
// Solidity: function maxNonRedeemedDelegatorRequestCount() view returns(uint256)
func (_SDUtilityPool *SDUtilityPoolCaller) MaxNonRedeemedDelegatorRequestCount(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _SDUtilityPool.contract.Call(opts, &out, "maxNonRedeemedDelegatorRequestCount")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MaxNonRedeemedDelegatorRequestCount is a free data retrieval call binding the contract method 0xc083f3e6.
//
// Solidity: function maxNonRedeemedDelegatorRequestCount() view returns(uint256)
func (_SDUtilityPool *SDUtilityPoolSession) MaxNonRedeemedDelegatorRequestCount() (*big.Int, error) {
	return _SDUtilityPool.Contract.MaxNonRedeemedDelegatorRequestCount(&_SDUtilityPool.CallOpts)
}

// MaxNonRedeemedDelegatorRequestCount is a free data retrieval call binding the contract method 0xc083f3e6.
//
// Solidity: function maxNonRedeemedDelegatorRequestCount() view returns(uint256)
func (_SDUtilityPool *SDUtilityPoolCallerSession) MaxNonRedeemedDelegatorRequestCount() (*big.Int, error) {
	return _SDUtilityPool.Contract.MaxNonRedeemedDelegatorRequestCount(&_SDUtilityPool.CallOpts)
}

// MinBlockDelayToFinalizeRequest is a free data retrieval call binding the contract method 0xbc1428f6.
//
// Solidity: function minBlockDelayToFinalizeRequest() view returns(uint256)
func (_SDUtilityPool *SDUtilityPoolCaller) MinBlockDelayToFinalizeRequest(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _SDUtilityPool.contract.Call(opts, &out, "minBlockDelayToFinalizeRequest")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MinBlockDelayToFinalizeRequest is a free data retrieval call binding the contract method 0xbc1428f6.
//
// Solidity: function minBlockDelayToFinalizeRequest() view returns(uint256)
func (_SDUtilityPool *SDUtilityPoolSession) MinBlockDelayToFinalizeRequest() (*big.Int, error) {
	return _SDUtilityPool.Contract.MinBlockDelayToFinalizeRequest(&_SDUtilityPool.CallOpts)
}

// MinBlockDelayToFinalizeRequest is a free data retrieval call binding the contract method 0xbc1428f6.
//
// Solidity: function minBlockDelayToFinalizeRequest() view returns(uint256)
func (_SDUtilityPool *SDUtilityPoolCallerSession) MinBlockDelayToFinalizeRequest() (*big.Int, error) {
	return _SDUtilityPool.Contract.MinBlockDelayToFinalizeRequest(&_SDUtilityPool.CallOpts)
}

// NextRequestId is a free data retrieval call binding the contract method 0x6a84a985.
//
// Solidity: function nextRequestId() view returns(uint256)
func (_SDUtilityPool *SDUtilityPoolCaller) NextRequestId(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _SDUtilityPool.contract.Call(opts, &out, "nextRequestId")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// NextRequestId is a free data retrieval call binding the contract method 0x6a84a985.
//
// Solidity: function nextRequestId() view returns(uint256)
func (_SDUtilityPool *SDUtilityPoolSession) NextRequestId() (*big.Int, error) {
	return _SDUtilityPool.Contract.NextRequestId(&_SDUtilityPool.CallOpts)
}

// NextRequestId is a free data retrieval call binding the contract method 0x6a84a985.
//
// Solidity: function nextRequestId() view returns(uint256)
func (_SDUtilityPool *SDUtilityPoolCallerSession) NextRequestId() (*big.Int, error) {
	return _SDUtilityPool.Contract.NextRequestId(&_SDUtilityPool.CallOpts)
}

// NextRequestIdToFinalize is a free data retrieval call binding the contract method 0xbbb84362.
//
// Solidity: function nextRequestIdToFinalize() view returns(uint256)
func (_SDUtilityPool *SDUtilityPoolCaller) NextRequestIdToFinalize(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _SDUtilityPool.contract.Call(opts, &out, "nextRequestIdToFinalize")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// NextRequestIdToFinalize is a free data retrieval call binding the contract method 0xbbb84362.
//
// Solidity: function nextRequestIdToFinalize() view returns(uint256)
func (_SDUtilityPool *SDUtilityPoolSession) NextRequestIdToFinalize() (*big.Int, error) {
	return _SDUtilityPool.Contract.NextRequestIdToFinalize(&_SDUtilityPool.CallOpts)
}

// NextRequestIdToFinalize is a free data retrieval call binding the contract method 0xbbb84362.
//
// Solidity: function nextRequestIdToFinalize() view returns(uint256)
func (_SDUtilityPool *SDUtilityPoolCallerSession) NextRequestIdToFinalize() (*big.Int, error) {
	return _SDUtilityPool.Contract.NextRequestIdToFinalize(&_SDUtilityPool.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_SDUtilityPool *SDUtilityPoolCaller) Paused(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _SDUtilityPool.contract.Call(opts, &out, "paused")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_SDUtilityPool *SDUtilityPoolSession) Paused() (bool, error) {
	return _SDUtilityPool.Contract.Paused(&_SDUtilityPool.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_SDUtilityPool *SDUtilityPoolCallerSession) Paused() (bool, error) {
	return _SDUtilityPool.Contract.Paused(&_SDUtilityPool.CallOpts)
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

// ProtocolFeeFactor is a free data retrieval call binding the contract method 0x11e3bdde.
//
// Solidity: function protocolFeeFactor() view returns(uint256)
func (_SDUtilityPool *SDUtilityPoolCaller) ProtocolFeeFactor(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _SDUtilityPool.contract.Call(opts, &out, "protocolFeeFactor")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ProtocolFeeFactor is a free data retrieval call binding the contract method 0x11e3bdde.
//
// Solidity: function protocolFeeFactor() view returns(uint256)
func (_SDUtilityPool *SDUtilityPoolSession) ProtocolFeeFactor() (*big.Int, error) {
	return _SDUtilityPool.Contract.ProtocolFeeFactor(&_SDUtilityPool.CallOpts)
}

// ProtocolFeeFactor is a free data retrieval call binding the contract method 0x11e3bdde.
//
// Solidity: function protocolFeeFactor() view returns(uint256)
func (_SDUtilityPool *SDUtilityPoolCallerSession) ProtocolFeeFactor() (*big.Int, error) {
	return _SDUtilityPool.Contract.ProtocolFeeFactor(&_SDUtilityPool.CallOpts)
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

// RiskConfig is a free data retrieval call binding the contract method 0x7c90a6ca.
//
// Solidity: function riskConfig() view returns(uint256 liquidationThreshold, uint256 liquidationBonusPercent, uint256 liquidationFeePercent, uint256 ltv)
func (_SDUtilityPool *SDUtilityPoolCaller) RiskConfig(opts *bind.CallOpts) (struct {
	LiquidationThreshold    *big.Int
	LiquidationBonusPercent *big.Int
	LiquidationFeePercent   *big.Int
	Ltv                     *big.Int
}, error) {
	var out []interface{}
	err := _SDUtilityPool.contract.Call(opts, &out, "riskConfig")

	outstruct := new(struct {
		LiquidationThreshold    *big.Int
		LiquidationBonusPercent *big.Int
		LiquidationFeePercent   *big.Int
		Ltv                     *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.LiquidationThreshold = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.LiquidationBonusPercent = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.LiquidationFeePercent = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	outstruct.Ltv = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// RiskConfig is a free data retrieval call binding the contract method 0x7c90a6ca.
//
// Solidity: function riskConfig() view returns(uint256 liquidationThreshold, uint256 liquidationBonusPercent, uint256 liquidationFeePercent, uint256 ltv)
func (_SDUtilityPool *SDUtilityPoolSession) RiskConfig() (struct {
	LiquidationThreshold    *big.Int
	LiquidationBonusPercent *big.Int
	LiquidationFeePercent   *big.Int
	Ltv                     *big.Int
}, error) {
	return _SDUtilityPool.Contract.RiskConfig(&_SDUtilityPool.CallOpts)
}

// RiskConfig is a free data retrieval call binding the contract method 0x7c90a6ca.
//
// Solidity: function riskConfig() view returns(uint256 liquidationThreshold, uint256 liquidationBonusPercent, uint256 liquidationFeePercent, uint256 ltv)
func (_SDUtilityPool *SDUtilityPoolCallerSession) RiskConfig() (struct {
	LiquidationThreshold    *big.Int
	LiquidationBonusPercent *big.Int
	LiquidationFeePercent   *big.Int
	Ltv                     *big.Int
}, error) {
	return _SDUtilityPool.Contract.RiskConfig(&_SDUtilityPool.CallOpts)
}

// SdRequestedForWithdraw is a free data retrieval call binding the contract method 0x3b92e3cf.
//
// Solidity: function sdRequestedForWithdraw() view returns(uint256)
func (_SDUtilityPool *SDUtilityPoolCaller) SdRequestedForWithdraw(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _SDUtilityPool.contract.Call(opts, &out, "sdRequestedForWithdraw")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// SdRequestedForWithdraw is a free data retrieval call binding the contract method 0x3b92e3cf.
//
// Solidity: function sdRequestedForWithdraw() view returns(uint256)
func (_SDUtilityPool *SDUtilityPoolSession) SdRequestedForWithdraw() (*big.Int, error) {
	return _SDUtilityPool.Contract.SdRequestedForWithdraw(&_SDUtilityPool.CallOpts)
}

// SdRequestedForWithdraw is a free data retrieval call binding the contract method 0x3b92e3cf.
//
// Solidity: function sdRequestedForWithdraw() view returns(uint256)
func (_SDUtilityPool *SDUtilityPoolCallerSession) SdRequestedForWithdraw() (*big.Int, error) {
	return _SDUtilityPool.Contract.SdRequestedForWithdraw(&_SDUtilityPool.CallOpts)
}

// SdReservedForClaim is a free data retrieval call binding the contract method 0x2b886941.
//
// Solidity: function sdReservedForClaim() view returns(uint256)
func (_SDUtilityPool *SDUtilityPoolCaller) SdReservedForClaim(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _SDUtilityPool.contract.Call(opts, &out, "sdReservedForClaim")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// SdReservedForClaim is a free data retrieval call binding the contract method 0x2b886941.
//
// Solidity: function sdReservedForClaim() view returns(uint256)
func (_SDUtilityPool *SDUtilityPoolSession) SdReservedForClaim() (*big.Int, error) {
	return _SDUtilityPool.Contract.SdReservedForClaim(&_SDUtilityPool.CallOpts)
}

// SdReservedForClaim is a free data retrieval call binding the contract method 0x2b886941.
//
// Solidity: function sdReservedForClaim() view returns(uint256)
func (_SDUtilityPool *SDUtilityPoolCallerSession) SdReservedForClaim() (*big.Int, error) {
	return _SDUtilityPool.Contract.SdReservedForClaim(&_SDUtilityPool.CallOpts)
}

// StaderConfig is a free data retrieval call binding the contract method 0x490ffa35.
//
// Solidity: function staderConfig() view returns(address)
func (_SDUtilityPool *SDUtilityPoolCaller) StaderConfig(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _SDUtilityPool.contract.Call(opts, &out, "staderConfig")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// StaderConfig is a free data retrieval call binding the contract method 0x490ffa35.
//
// Solidity: function staderConfig() view returns(address)
func (_SDUtilityPool *SDUtilityPoolSession) StaderConfig() (common.Address, error) {
	return _SDUtilityPool.Contract.StaderConfig(&_SDUtilityPool.CallOpts)
}

// StaderConfig is a free data retrieval call binding the contract method 0x490ffa35.
//
// Solidity: function staderConfig() view returns(address)
func (_SDUtilityPool *SDUtilityPoolCallerSession) StaderConfig() (common.Address, error) {
	return _SDUtilityPool.Contract.StaderConfig(&_SDUtilityPool.CallOpts)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_SDUtilityPool *SDUtilityPoolCaller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _SDUtilityPool.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_SDUtilityPool *SDUtilityPoolSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _SDUtilityPool.Contract.SupportsInterface(&_SDUtilityPool.CallOpts, interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_SDUtilityPool *SDUtilityPoolCallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _SDUtilityPool.Contract.SupportsInterface(&_SDUtilityPool.CallOpts, interfaceId)
}

// TotalUtilizedSD is a free data retrieval call binding the contract method 0x8763a328.
//
// Solidity: function totalUtilizedSD() view returns(uint256)
func (_SDUtilityPool *SDUtilityPoolCaller) TotalUtilizedSD(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _SDUtilityPool.contract.Call(opts, &out, "totalUtilizedSD")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalUtilizedSD is a free data retrieval call binding the contract method 0x8763a328.
//
// Solidity: function totalUtilizedSD() view returns(uint256)
func (_SDUtilityPool *SDUtilityPoolSession) TotalUtilizedSD() (*big.Int, error) {
	return _SDUtilityPool.Contract.TotalUtilizedSD(&_SDUtilityPool.CallOpts)
}

// TotalUtilizedSD is a free data retrieval call binding the contract method 0x8763a328.
//
// Solidity: function totalUtilizedSD() view returns(uint256)
func (_SDUtilityPool *SDUtilityPoolCallerSession) TotalUtilizedSD() (*big.Int, error) {
	return _SDUtilityPool.Contract.TotalUtilizedSD(&_SDUtilityPool.CallOpts)
}

// UndelegationPeriodInBlocks is a free data retrieval call binding the contract method 0xbbf1a4ef.
//
// Solidity: function undelegationPeriodInBlocks() view returns(uint256)
func (_SDUtilityPool *SDUtilityPoolCaller) UndelegationPeriodInBlocks(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _SDUtilityPool.contract.Call(opts, &out, "undelegationPeriodInBlocks")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// UndelegationPeriodInBlocks is a free data retrieval call binding the contract method 0xbbf1a4ef.
//
// Solidity: function undelegationPeriodInBlocks() view returns(uint256)
func (_SDUtilityPool *SDUtilityPoolSession) UndelegationPeriodInBlocks() (*big.Int, error) {
	return _SDUtilityPool.Contract.UndelegationPeriodInBlocks(&_SDUtilityPool.CallOpts)
}

// UndelegationPeriodInBlocks is a free data retrieval call binding the contract method 0xbbf1a4ef.
//
// Solidity: function undelegationPeriodInBlocks() view returns(uint256)
func (_SDUtilityPool *SDUtilityPoolCallerSession) UndelegationPeriodInBlocks() (*big.Int, error) {
	return _SDUtilityPool.Contract.UndelegationPeriodInBlocks(&_SDUtilityPool.CallOpts)
}

// UtilizationRatePerBlock is a free data retrieval call binding the contract method 0x962c7070.
//
// Solidity: function utilizationRatePerBlock() view returns(uint256)
func (_SDUtilityPool *SDUtilityPoolCaller) UtilizationRatePerBlock(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _SDUtilityPool.contract.Call(opts, &out, "utilizationRatePerBlock")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// UtilizationRatePerBlock is a free data retrieval call binding the contract method 0x962c7070.
//
// Solidity: function utilizationRatePerBlock() view returns(uint256)
func (_SDUtilityPool *SDUtilityPoolSession) UtilizationRatePerBlock() (*big.Int, error) {
	return _SDUtilityPool.Contract.UtilizationRatePerBlock(&_SDUtilityPool.CallOpts)
}

// UtilizationRatePerBlock is a free data retrieval call binding the contract method 0x962c7070.
//
// Solidity: function utilizationRatePerBlock() view returns(uint256)
func (_SDUtilityPool *SDUtilityPoolCallerSession) UtilizationRatePerBlock() (*big.Int, error) {
	return _SDUtilityPool.Contract.UtilizationRatePerBlock(&_SDUtilityPool.CallOpts)
}

// UtilizeIndex is a free data retrieval call binding the contract method 0x634c48c7.
//
// Solidity: function utilizeIndex() view returns(uint256)
func (_SDUtilityPool *SDUtilityPoolCaller) UtilizeIndex(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _SDUtilityPool.contract.Call(opts, &out, "utilizeIndex")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// UtilizeIndex is a free data retrieval call binding the contract method 0x634c48c7.
//
// Solidity: function utilizeIndex() view returns(uint256)
func (_SDUtilityPool *SDUtilityPoolSession) UtilizeIndex() (*big.Int, error) {
	return _SDUtilityPool.Contract.UtilizeIndex(&_SDUtilityPool.CallOpts)
}

// UtilizeIndex is a free data retrieval call binding the contract method 0x634c48c7.
//
// Solidity: function utilizeIndex() view returns(uint256)
func (_SDUtilityPool *SDUtilityPoolCallerSession) UtilizeIndex() (*big.Int, error) {
	return _SDUtilityPool.Contract.UtilizeIndex(&_SDUtilityPool.CallOpts)
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
// Solidity: function claim(uint256 _requestId) returns()
func (_SDUtilityPool *SDUtilityPoolTransactor) Claim(opts *bind.TransactOpts, _requestId *big.Int) (*types.Transaction, error) {
	return _SDUtilityPool.contract.Transact(opts, "claim", _requestId)
}

// Claim is a paid mutator transaction binding the contract method 0x379607f5.
//
// Solidity: function claim(uint256 _requestId) returns()
func (_SDUtilityPool *SDUtilityPoolSession) Claim(_requestId *big.Int) (*types.Transaction, error) {
	return _SDUtilityPool.Contract.Claim(&_SDUtilityPool.TransactOpts, _requestId)
}

// Claim is a paid mutator transaction binding the contract method 0x379607f5.
//
// Solidity: function claim(uint256 _requestId) returns()
func (_SDUtilityPool *SDUtilityPoolTransactorSession) Claim(_requestId *big.Int) (*types.Transaction, error) {
	return _SDUtilityPool.Contract.Claim(&_SDUtilityPool.TransactOpts, _requestId)
}

// ClaimLiquidation is a paid mutator transaction binding the contract method 0x732c8009.
//
// Solidity: function claimLiquidation(uint256 index) returns()
func (_SDUtilityPool *SDUtilityPoolTransactor) ClaimLiquidation(opts *bind.TransactOpts, index *big.Int) (*types.Transaction, error) {
	return _SDUtilityPool.contract.Transact(opts, "claimLiquidation", index)
}

// ClaimLiquidation is a paid mutator transaction binding the contract method 0x732c8009.
//
// Solidity: function claimLiquidation(uint256 index) returns()
func (_SDUtilityPool *SDUtilityPoolSession) ClaimLiquidation(index *big.Int) (*types.Transaction, error) {
	return _SDUtilityPool.Contract.ClaimLiquidation(&_SDUtilityPool.TransactOpts, index)
}

// ClaimLiquidation is a paid mutator transaction binding the contract method 0x732c8009.
//
// Solidity: function claimLiquidation(uint256 index) returns()
func (_SDUtilityPool *SDUtilityPoolTransactorSession) ClaimLiquidation(index *big.Int) (*types.Transaction, error) {
	return _SDUtilityPool.Contract.ClaimLiquidation(&_SDUtilityPool.TransactOpts, index)
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

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_SDUtilityPool *SDUtilityPoolTransactor) GrantRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _SDUtilityPool.contract.Transact(opts, "grantRole", role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_SDUtilityPool *SDUtilityPoolSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _SDUtilityPool.Contract.GrantRole(&_SDUtilityPool.TransactOpts, role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_SDUtilityPool *SDUtilityPoolTransactorSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _SDUtilityPool.Contract.GrantRole(&_SDUtilityPool.TransactOpts, role, account)
}

// Initialize is a paid mutator transaction binding the contract method 0x485cc955.
//
// Solidity: function initialize(address _admin, address _staderConfig) returns()
func (_SDUtilityPool *SDUtilityPoolTransactor) Initialize(opts *bind.TransactOpts, _admin common.Address, _staderConfig common.Address) (*types.Transaction, error) {
	return _SDUtilityPool.contract.Transact(opts, "initialize", _admin, _staderConfig)
}

// Initialize is a paid mutator transaction binding the contract method 0x485cc955.
//
// Solidity: function initialize(address _admin, address _staderConfig) returns()
func (_SDUtilityPool *SDUtilityPoolSession) Initialize(_admin common.Address, _staderConfig common.Address) (*types.Transaction, error) {
	return _SDUtilityPool.Contract.Initialize(&_SDUtilityPool.TransactOpts, _admin, _staderConfig)
}

// Initialize is a paid mutator transaction binding the contract method 0x485cc955.
//
// Solidity: function initialize(address _admin, address _staderConfig) returns()
func (_SDUtilityPool *SDUtilityPoolTransactorSession) Initialize(_admin common.Address, _staderConfig common.Address) (*types.Transaction, error) {
	return _SDUtilityPool.Contract.Initialize(&_SDUtilityPool.TransactOpts, _admin, _staderConfig)
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

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_SDUtilityPool *SDUtilityPoolTransactor) Pause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SDUtilityPool.contract.Transact(opts, "pause")
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_SDUtilityPool *SDUtilityPoolSession) Pause() (*types.Transaction, error) {
	return _SDUtilityPool.Contract.Pause(&_SDUtilityPool.TransactOpts)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_SDUtilityPool *SDUtilityPoolTransactorSession) Pause() (*types.Transaction, error) {
	return _SDUtilityPool.Contract.Pause(&_SDUtilityPool.TransactOpts)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address account) returns()
func (_SDUtilityPool *SDUtilityPoolTransactor) RenounceRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _SDUtilityPool.contract.Transact(opts, "renounceRole", role, account)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address account) returns()
func (_SDUtilityPool *SDUtilityPoolSession) RenounceRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _SDUtilityPool.Contract.RenounceRole(&_SDUtilityPool.TransactOpts, role, account)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address account) returns()
func (_SDUtilityPool *SDUtilityPoolTransactorSession) RenounceRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _SDUtilityPool.Contract.RenounceRole(&_SDUtilityPool.TransactOpts, role, account)
}

// Repay is a paid mutator transaction binding the contract method 0x371fd8e6.
//
// Solidity: function repay(uint256 repayAmount) returns(uint256 repaidAmount)
func (_SDUtilityPool *SDUtilityPoolTransactor) Repay(opts *bind.TransactOpts, repayAmount *big.Int) (*types.Transaction, error) {
	return _SDUtilityPool.contract.Transact(opts, "repay", repayAmount)
}

// Repay is a paid mutator transaction binding the contract method 0x371fd8e6.
//
// Solidity: function repay(uint256 repayAmount) returns(uint256 repaidAmount)
func (_SDUtilityPool *SDUtilityPoolSession) Repay(repayAmount *big.Int) (*types.Transaction, error) {
	return _SDUtilityPool.Contract.Repay(&_SDUtilityPool.TransactOpts, repayAmount)
}

// Repay is a paid mutator transaction binding the contract method 0x371fd8e6.
//
// Solidity: function repay(uint256 repayAmount) returns(uint256 repaidAmount)
func (_SDUtilityPool *SDUtilityPoolTransactorSession) Repay(repayAmount *big.Int) (*types.Transaction, error) {
	return _SDUtilityPool.Contract.Repay(&_SDUtilityPool.TransactOpts, repayAmount)
}

// RepayLiquidation is a paid mutator transaction binding the contract method 0x24faf5d9.
//
// Solidity: function repayLiquidation(address account) returns()
func (_SDUtilityPool *SDUtilityPoolTransactor) RepayLiquidation(opts *bind.TransactOpts, account common.Address) (*types.Transaction, error) {
	return _SDUtilityPool.contract.Transact(opts, "repayLiquidation", account)
}

// RepayLiquidation is a paid mutator transaction binding the contract method 0x24faf5d9.
//
// Solidity: function repayLiquidation(address account) returns()
func (_SDUtilityPool *SDUtilityPoolSession) RepayLiquidation(account common.Address) (*types.Transaction, error) {
	return _SDUtilityPool.Contract.RepayLiquidation(&_SDUtilityPool.TransactOpts, account)
}

// RepayLiquidation is a paid mutator transaction binding the contract method 0x24faf5d9.
//
// Solidity: function repayLiquidation(address account) returns()
func (_SDUtilityPool *SDUtilityPoolTransactorSession) RepayLiquidation(account common.Address) (*types.Transaction, error) {
	return _SDUtilityPool.Contract.RepayLiquidation(&_SDUtilityPool.TransactOpts, account)
}

// RepayOnBehalf is a paid mutator transaction binding the contract method 0x9f689e0b.
//
// Solidity: function repayOnBehalf(address utilizer, uint256 repayAmount) returns(uint256 repaidAmount)
func (_SDUtilityPool *SDUtilityPoolTransactor) RepayOnBehalf(opts *bind.TransactOpts, utilizer common.Address, repayAmount *big.Int) (*types.Transaction, error) {
	return _SDUtilityPool.contract.Transact(opts, "repayOnBehalf", utilizer, repayAmount)
}

// RepayOnBehalf is a paid mutator transaction binding the contract method 0x9f689e0b.
//
// Solidity: function repayOnBehalf(address utilizer, uint256 repayAmount) returns(uint256 repaidAmount)
func (_SDUtilityPool *SDUtilityPoolSession) RepayOnBehalf(utilizer common.Address, repayAmount *big.Int) (*types.Transaction, error) {
	return _SDUtilityPool.Contract.RepayOnBehalf(&_SDUtilityPool.TransactOpts, utilizer, repayAmount)
}

// RepayOnBehalf is a paid mutator transaction binding the contract method 0x9f689e0b.
//
// Solidity: function repayOnBehalf(address utilizer, uint256 repayAmount) returns(uint256 repaidAmount)
func (_SDUtilityPool *SDUtilityPoolTransactorSession) RepayOnBehalf(utilizer common.Address, repayAmount *big.Int) (*types.Transaction, error) {
	return _SDUtilityPool.Contract.RepayOnBehalf(&_SDUtilityPool.TransactOpts, utilizer, repayAmount)
}

// RequestWithdraw is a paid mutator transaction binding the contract method 0x745400c9.
//
// Solidity: function requestWithdraw(uint256 _cTokenAmount) returns(uint256 _requestId)
func (_SDUtilityPool *SDUtilityPoolTransactor) RequestWithdraw(opts *bind.TransactOpts, _cTokenAmount *big.Int) (*types.Transaction, error) {
	return _SDUtilityPool.contract.Transact(opts, "requestWithdraw", _cTokenAmount)
}

// RequestWithdraw is a paid mutator transaction binding the contract method 0x745400c9.
//
// Solidity: function requestWithdraw(uint256 _cTokenAmount) returns(uint256 _requestId)
func (_SDUtilityPool *SDUtilityPoolSession) RequestWithdraw(_cTokenAmount *big.Int) (*types.Transaction, error) {
	return _SDUtilityPool.Contract.RequestWithdraw(&_SDUtilityPool.TransactOpts, _cTokenAmount)
}

// RequestWithdraw is a paid mutator transaction binding the contract method 0x745400c9.
//
// Solidity: function requestWithdraw(uint256 _cTokenAmount) returns(uint256 _requestId)
func (_SDUtilityPool *SDUtilityPoolTransactorSession) RequestWithdraw(_cTokenAmount *big.Int) (*types.Transaction, error) {
	return _SDUtilityPool.Contract.RequestWithdraw(&_SDUtilityPool.TransactOpts, _cTokenAmount)
}

// RequestWithdrawWithSDAmount is a paid mutator transaction binding the contract method 0xc51cd1cc.
//
// Solidity: function requestWithdrawWithSDAmount(uint256 _sdAmount) returns(uint256 _requestId)
func (_SDUtilityPool *SDUtilityPoolTransactor) RequestWithdrawWithSDAmount(opts *bind.TransactOpts, _sdAmount *big.Int) (*types.Transaction, error) {
	return _SDUtilityPool.contract.Transact(opts, "requestWithdrawWithSDAmount", _sdAmount)
}

// RequestWithdrawWithSDAmount is a paid mutator transaction binding the contract method 0xc51cd1cc.
//
// Solidity: function requestWithdrawWithSDAmount(uint256 _sdAmount) returns(uint256 _requestId)
func (_SDUtilityPool *SDUtilityPoolSession) RequestWithdrawWithSDAmount(_sdAmount *big.Int) (*types.Transaction, error) {
	return _SDUtilityPool.Contract.RequestWithdrawWithSDAmount(&_SDUtilityPool.TransactOpts, _sdAmount)
}

// RequestWithdrawWithSDAmount is a paid mutator transaction binding the contract method 0xc51cd1cc.
//
// Solidity: function requestWithdrawWithSDAmount(uint256 _sdAmount) returns(uint256 _requestId)
func (_SDUtilityPool *SDUtilityPoolTransactorSession) RequestWithdrawWithSDAmount(_sdAmount *big.Int) (*types.Transaction, error) {
	return _SDUtilityPool.Contract.RequestWithdrawWithSDAmount(&_SDUtilityPool.TransactOpts, _sdAmount)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_SDUtilityPool *SDUtilityPoolTransactor) RevokeRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _SDUtilityPool.contract.Transact(opts, "revokeRole", role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_SDUtilityPool *SDUtilityPoolSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _SDUtilityPool.Contract.RevokeRole(&_SDUtilityPool.TransactOpts, role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_SDUtilityPool *SDUtilityPoolTransactorSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _SDUtilityPool.Contract.RevokeRole(&_SDUtilityPool.TransactOpts, role, account)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_SDUtilityPool *SDUtilityPoolTransactor) Unpause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SDUtilityPool.contract.Transact(opts, "unpause")
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_SDUtilityPool *SDUtilityPoolSession) Unpause() (*types.Transaction, error) {
	return _SDUtilityPool.Contract.Unpause(&_SDUtilityPool.TransactOpts)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_SDUtilityPool *SDUtilityPoolTransactorSession) Unpause() (*types.Transaction, error) {
	return _SDUtilityPool.Contract.Unpause(&_SDUtilityPool.TransactOpts)
}

// UpdateFinalizationBatchLimit is a paid mutator transaction binding the contract method 0x267fca73.
//
// Solidity: function updateFinalizationBatchLimit(uint256 _finalizationBatchLimit) returns()
func (_SDUtilityPool *SDUtilityPoolTransactor) UpdateFinalizationBatchLimit(opts *bind.TransactOpts, _finalizationBatchLimit *big.Int) (*types.Transaction, error) {
	return _SDUtilityPool.contract.Transact(opts, "updateFinalizationBatchLimit", _finalizationBatchLimit)
}

// UpdateFinalizationBatchLimit is a paid mutator transaction binding the contract method 0x267fca73.
//
// Solidity: function updateFinalizationBatchLimit(uint256 _finalizationBatchLimit) returns()
func (_SDUtilityPool *SDUtilityPoolSession) UpdateFinalizationBatchLimit(_finalizationBatchLimit *big.Int) (*types.Transaction, error) {
	return _SDUtilityPool.Contract.UpdateFinalizationBatchLimit(&_SDUtilityPool.TransactOpts, _finalizationBatchLimit)
}

// UpdateFinalizationBatchLimit is a paid mutator transaction binding the contract method 0x267fca73.
//
// Solidity: function updateFinalizationBatchLimit(uint256 _finalizationBatchLimit) returns()
func (_SDUtilityPool *SDUtilityPoolTransactorSession) UpdateFinalizationBatchLimit(_finalizationBatchLimit *big.Int) (*types.Transaction, error) {
	return _SDUtilityPool.Contract.UpdateFinalizationBatchLimit(&_SDUtilityPool.TransactOpts, _finalizationBatchLimit)
}

// UpdateMaxETHWorthOfSDPerValidator is a paid mutator transaction binding the contract method 0x5393618e.
//
// Solidity: function updateMaxETHWorthOfSDPerValidator(uint256 _maxETHWorthOfSDPerValidator) returns()
func (_SDUtilityPool *SDUtilityPoolTransactor) UpdateMaxETHWorthOfSDPerValidator(opts *bind.TransactOpts, _maxETHWorthOfSDPerValidator *big.Int) (*types.Transaction, error) {
	return _SDUtilityPool.contract.Transact(opts, "updateMaxETHWorthOfSDPerValidator", _maxETHWorthOfSDPerValidator)
}

// UpdateMaxETHWorthOfSDPerValidator is a paid mutator transaction binding the contract method 0x5393618e.
//
// Solidity: function updateMaxETHWorthOfSDPerValidator(uint256 _maxETHWorthOfSDPerValidator) returns()
func (_SDUtilityPool *SDUtilityPoolSession) UpdateMaxETHWorthOfSDPerValidator(_maxETHWorthOfSDPerValidator *big.Int) (*types.Transaction, error) {
	return _SDUtilityPool.Contract.UpdateMaxETHWorthOfSDPerValidator(&_SDUtilityPool.TransactOpts, _maxETHWorthOfSDPerValidator)
}

// UpdateMaxETHWorthOfSDPerValidator is a paid mutator transaction binding the contract method 0x5393618e.
//
// Solidity: function updateMaxETHWorthOfSDPerValidator(uint256 _maxETHWorthOfSDPerValidator) returns()
func (_SDUtilityPool *SDUtilityPoolTransactorSession) UpdateMaxETHWorthOfSDPerValidator(_maxETHWorthOfSDPerValidator *big.Int) (*types.Transaction, error) {
	return _SDUtilityPool.Contract.UpdateMaxETHWorthOfSDPerValidator(&_SDUtilityPool.TransactOpts, _maxETHWorthOfSDPerValidator)
}

// UpdateMaxNonRedeemedDelegatorRequestCount is a paid mutator transaction binding the contract method 0xee63e5f9.
//
// Solidity: function updateMaxNonRedeemedDelegatorRequestCount(uint256 _count) returns()
func (_SDUtilityPool *SDUtilityPoolTransactor) UpdateMaxNonRedeemedDelegatorRequestCount(opts *bind.TransactOpts, _count *big.Int) (*types.Transaction, error) {
	return _SDUtilityPool.contract.Transact(opts, "updateMaxNonRedeemedDelegatorRequestCount", _count)
}

// UpdateMaxNonRedeemedDelegatorRequestCount is a paid mutator transaction binding the contract method 0xee63e5f9.
//
// Solidity: function updateMaxNonRedeemedDelegatorRequestCount(uint256 _count) returns()
func (_SDUtilityPool *SDUtilityPoolSession) UpdateMaxNonRedeemedDelegatorRequestCount(_count *big.Int) (*types.Transaction, error) {
	return _SDUtilityPool.Contract.UpdateMaxNonRedeemedDelegatorRequestCount(&_SDUtilityPool.TransactOpts, _count)
}

// UpdateMaxNonRedeemedDelegatorRequestCount is a paid mutator transaction binding the contract method 0xee63e5f9.
//
// Solidity: function updateMaxNonRedeemedDelegatorRequestCount(uint256 _count) returns()
func (_SDUtilityPool *SDUtilityPoolTransactorSession) UpdateMaxNonRedeemedDelegatorRequestCount(_count *big.Int) (*types.Transaction, error) {
	return _SDUtilityPool.Contract.UpdateMaxNonRedeemedDelegatorRequestCount(&_SDUtilityPool.TransactOpts, _count)
}

// UpdateMinBlockDelayToFinalizeRequest is a paid mutator transaction binding the contract method 0x4a2965af.
//
// Solidity: function updateMinBlockDelayToFinalizeRequest(uint256 _minBlockDelayToFinalizeRequest) returns()
func (_SDUtilityPool *SDUtilityPoolTransactor) UpdateMinBlockDelayToFinalizeRequest(opts *bind.TransactOpts, _minBlockDelayToFinalizeRequest *big.Int) (*types.Transaction, error) {
	return _SDUtilityPool.contract.Transact(opts, "updateMinBlockDelayToFinalizeRequest", _minBlockDelayToFinalizeRequest)
}

// UpdateMinBlockDelayToFinalizeRequest is a paid mutator transaction binding the contract method 0x4a2965af.
//
// Solidity: function updateMinBlockDelayToFinalizeRequest(uint256 _minBlockDelayToFinalizeRequest) returns()
func (_SDUtilityPool *SDUtilityPoolSession) UpdateMinBlockDelayToFinalizeRequest(_minBlockDelayToFinalizeRequest *big.Int) (*types.Transaction, error) {
	return _SDUtilityPool.Contract.UpdateMinBlockDelayToFinalizeRequest(&_SDUtilityPool.TransactOpts, _minBlockDelayToFinalizeRequest)
}

// UpdateMinBlockDelayToFinalizeRequest is a paid mutator transaction binding the contract method 0x4a2965af.
//
// Solidity: function updateMinBlockDelayToFinalizeRequest(uint256 _minBlockDelayToFinalizeRequest) returns()
func (_SDUtilityPool *SDUtilityPoolTransactorSession) UpdateMinBlockDelayToFinalizeRequest(_minBlockDelayToFinalizeRequest *big.Int) (*types.Transaction, error) {
	return _SDUtilityPool.Contract.UpdateMinBlockDelayToFinalizeRequest(&_SDUtilityPool.TransactOpts, _minBlockDelayToFinalizeRequest)
}

// UpdateProtocolFeeFactor is a paid mutator transaction binding the contract method 0x865cbba3.
//
// Solidity: function updateProtocolFeeFactor(uint256 _protocolFeeFactor) returns()
func (_SDUtilityPool *SDUtilityPoolTransactor) UpdateProtocolFeeFactor(opts *bind.TransactOpts, _protocolFeeFactor *big.Int) (*types.Transaction, error) {
	return _SDUtilityPool.contract.Transact(opts, "updateProtocolFeeFactor", _protocolFeeFactor)
}

// UpdateProtocolFeeFactor is a paid mutator transaction binding the contract method 0x865cbba3.
//
// Solidity: function updateProtocolFeeFactor(uint256 _protocolFeeFactor) returns()
func (_SDUtilityPool *SDUtilityPoolSession) UpdateProtocolFeeFactor(_protocolFeeFactor *big.Int) (*types.Transaction, error) {
	return _SDUtilityPool.Contract.UpdateProtocolFeeFactor(&_SDUtilityPool.TransactOpts, _protocolFeeFactor)
}

// UpdateProtocolFeeFactor is a paid mutator transaction binding the contract method 0x865cbba3.
//
// Solidity: function updateProtocolFeeFactor(uint256 _protocolFeeFactor) returns()
func (_SDUtilityPool *SDUtilityPoolTransactorSession) UpdateProtocolFeeFactor(_protocolFeeFactor *big.Int) (*types.Transaction, error) {
	return _SDUtilityPool.Contract.UpdateProtocolFeeFactor(&_SDUtilityPool.TransactOpts, _protocolFeeFactor)
}

// UpdateRiskConfig is a paid mutator transaction binding the contract method 0x62569a51.
//
// Solidity: function updateRiskConfig(uint256 liquidationThreshold, uint256 liquidationBonusPercent, uint256 liquidationFeePercent, uint256 ltv) returns()
func (_SDUtilityPool *SDUtilityPoolTransactor) UpdateRiskConfig(opts *bind.TransactOpts, liquidationThreshold *big.Int, liquidationBonusPercent *big.Int, liquidationFeePercent *big.Int, ltv *big.Int) (*types.Transaction, error) {
	return _SDUtilityPool.contract.Transact(opts, "updateRiskConfig", liquidationThreshold, liquidationBonusPercent, liquidationFeePercent, ltv)
}

// UpdateRiskConfig is a paid mutator transaction binding the contract method 0x62569a51.
//
// Solidity: function updateRiskConfig(uint256 liquidationThreshold, uint256 liquidationBonusPercent, uint256 liquidationFeePercent, uint256 ltv) returns()
func (_SDUtilityPool *SDUtilityPoolSession) UpdateRiskConfig(liquidationThreshold *big.Int, liquidationBonusPercent *big.Int, liquidationFeePercent *big.Int, ltv *big.Int) (*types.Transaction, error) {
	return _SDUtilityPool.Contract.UpdateRiskConfig(&_SDUtilityPool.TransactOpts, liquidationThreshold, liquidationBonusPercent, liquidationFeePercent, ltv)
}

// UpdateRiskConfig is a paid mutator transaction binding the contract method 0x62569a51.
//
// Solidity: function updateRiskConfig(uint256 liquidationThreshold, uint256 liquidationBonusPercent, uint256 liquidationFeePercent, uint256 ltv) returns()
func (_SDUtilityPool *SDUtilityPoolTransactorSession) UpdateRiskConfig(liquidationThreshold *big.Int, liquidationBonusPercent *big.Int, liquidationFeePercent *big.Int, ltv *big.Int) (*types.Transaction, error) {
	return _SDUtilityPool.Contract.UpdateRiskConfig(&_SDUtilityPool.TransactOpts, liquidationThreshold, liquidationBonusPercent, liquidationFeePercent, ltv)
}

// UpdateStaderConfig is a paid mutator transaction binding the contract method 0x9ee804cb.
//
// Solidity: function updateStaderConfig(address _staderConfig) returns()
func (_SDUtilityPool *SDUtilityPoolTransactor) UpdateStaderConfig(opts *bind.TransactOpts, _staderConfig common.Address) (*types.Transaction, error) {
	return _SDUtilityPool.contract.Transact(opts, "updateStaderConfig", _staderConfig)
}

// UpdateStaderConfig is a paid mutator transaction binding the contract method 0x9ee804cb.
//
// Solidity: function updateStaderConfig(address _staderConfig) returns()
func (_SDUtilityPool *SDUtilityPoolSession) UpdateStaderConfig(_staderConfig common.Address) (*types.Transaction, error) {
	return _SDUtilityPool.Contract.UpdateStaderConfig(&_SDUtilityPool.TransactOpts, _staderConfig)
}

// UpdateStaderConfig is a paid mutator transaction binding the contract method 0x9ee804cb.
//
// Solidity: function updateStaderConfig(address _staderConfig) returns()
func (_SDUtilityPool *SDUtilityPoolTransactorSession) UpdateStaderConfig(_staderConfig common.Address) (*types.Transaction, error) {
	return _SDUtilityPool.Contract.UpdateStaderConfig(&_SDUtilityPool.TransactOpts, _staderConfig)
}

// UpdateUndelegationPeriodInBlocks is a paid mutator transaction binding the contract method 0xf9d1f10a.
//
// Solidity: function updateUndelegationPeriodInBlocks(uint256 _undelegationPeriodInBlocks) returns()
func (_SDUtilityPool *SDUtilityPoolTransactor) UpdateUndelegationPeriodInBlocks(opts *bind.TransactOpts, _undelegationPeriodInBlocks *big.Int) (*types.Transaction, error) {
	return _SDUtilityPool.contract.Transact(opts, "updateUndelegationPeriodInBlocks", _undelegationPeriodInBlocks)
}

// UpdateUndelegationPeriodInBlocks is a paid mutator transaction binding the contract method 0xf9d1f10a.
//
// Solidity: function updateUndelegationPeriodInBlocks(uint256 _undelegationPeriodInBlocks) returns()
func (_SDUtilityPool *SDUtilityPoolSession) UpdateUndelegationPeriodInBlocks(_undelegationPeriodInBlocks *big.Int) (*types.Transaction, error) {
	return _SDUtilityPool.Contract.UpdateUndelegationPeriodInBlocks(&_SDUtilityPool.TransactOpts, _undelegationPeriodInBlocks)
}

// UpdateUndelegationPeriodInBlocks is a paid mutator transaction binding the contract method 0xf9d1f10a.
//
// Solidity: function updateUndelegationPeriodInBlocks(uint256 _undelegationPeriodInBlocks) returns()
func (_SDUtilityPool *SDUtilityPoolTransactorSession) UpdateUndelegationPeriodInBlocks(_undelegationPeriodInBlocks *big.Int) (*types.Transaction, error) {
	return _SDUtilityPool.Contract.UpdateUndelegationPeriodInBlocks(&_SDUtilityPool.TransactOpts, _undelegationPeriodInBlocks)
}

// UpdateUtilizationRatePerBlock is a paid mutator transaction binding the contract method 0xcb2d89dd.
//
// Solidity: function updateUtilizationRatePerBlock(uint256 _utilizationRatePerBlock) returns()
func (_SDUtilityPool *SDUtilityPoolTransactor) UpdateUtilizationRatePerBlock(opts *bind.TransactOpts, _utilizationRatePerBlock *big.Int) (*types.Transaction, error) {
	return _SDUtilityPool.contract.Transact(opts, "updateUtilizationRatePerBlock", _utilizationRatePerBlock)
}

// UpdateUtilizationRatePerBlock is a paid mutator transaction binding the contract method 0xcb2d89dd.
//
// Solidity: function updateUtilizationRatePerBlock(uint256 _utilizationRatePerBlock) returns()
func (_SDUtilityPool *SDUtilityPoolSession) UpdateUtilizationRatePerBlock(_utilizationRatePerBlock *big.Int) (*types.Transaction, error) {
	return _SDUtilityPool.Contract.UpdateUtilizationRatePerBlock(&_SDUtilityPool.TransactOpts, _utilizationRatePerBlock)
}

// UpdateUtilizationRatePerBlock is a paid mutator transaction binding the contract method 0xcb2d89dd.
//
// Solidity: function updateUtilizationRatePerBlock(uint256 _utilizationRatePerBlock) returns()
func (_SDUtilityPool *SDUtilityPoolTransactorSession) UpdateUtilizationRatePerBlock(_utilizationRatePerBlock *big.Int) (*types.Transaction, error) {
	return _SDUtilityPool.Contract.UpdateUtilizationRatePerBlock(&_SDUtilityPool.TransactOpts, _utilizationRatePerBlock)
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

// WithdrawProtocolFee is a paid mutator transaction binding the contract method 0x424cd833.
//
// Solidity: function withdrawProtocolFee(uint256 _amount) returns()
func (_SDUtilityPool *SDUtilityPoolTransactor) WithdrawProtocolFee(opts *bind.TransactOpts, _amount *big.Int) (*types.Transaction, error) {
	return _SDUtilityPool.contract.Transact(opts, "withdrawProtocolFee", _amount)
}

// WithdrawProtocolFee is a paid mutator transaction binding the contract method 0x424cd833.
//
// Solidity: function withdrawProtocolFee(uint256 _amount) returns()
func (_SDUtilityPool *SDUtilityPoolSession) WithdrawProtocolFee(_amount *big.Int) (*types.Transaction, error) {
	return _SDUtilityPool.Contract.WithdrawProtocolFee(&_SDUtilityPool.TransactOpts, _amount)
}

// WithdrawProtocolFee is a paid mutator transaction binding the contract method 0x424cd833.
//
// Solidity: function withdrawProtocolFee(uint256 _amount) returns()
func (_SDUtilityPool *SDUtilityPoolTransactorSession) WithdrawProtocolFee(_amount *big.Int) (*types.Transaction, error) {
	return _SDUtilityPool.Contract.WithdrawProtocolFee(&_SDUtilityPool.TransactOpts, _amount)
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

// SDUtilityPoolClaimedLiquidationIterator is returned from FilterClaimedLiquidation and is used to iterate over the raw logs and unpacked data for ClaimedLiquidation events raised by the SDUtilityPool contract.
type SDUtilityPoolClaimedLiquidationIterator struct {
	Event *SDUtilityPoolClaimedLiquidation // Event containing the contract specifics and raw log

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
func (it *SDUtilityPoolClaimedLiquidationIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SDUtilityPoolClaimedLiquidation)
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
		it.Event = new(SDUtilityPoolClaimedLiquidation)
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
func (it *SDUtilityPoolClaimedLiquidationIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SDUtilityPoolClaimedLiquidationIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SDUtilityPoolClaimedLiquidation represents a ClaimedLiquidation event raised by the SDUtilityPool contract.
type SDUtilityPoolClaimedLiquidation struct {
	Liquidator            common.Address
	LiquidationBonusInEth *big.Int
	LiquidationFeeInEth   *big.Int
	Raw                   types.Log // Blockchain specific contextual infos
}

// FilterClaimedLiquidation is a free log retrieval operation binding the contract event 0x87e8e54247cd0a3f2aae4999ad6838c60057f6de710e7b4eec301fe80eb55199.
//
// Solidity: event ClaimedLiquidation(address indexed liquidator, uint256 liquidationBonusInEth, uint256 liquidationFeeInEth)
func (_SDUtilityPool *SDUtilityPoolFilterer) FilterClaimedLiquidation(opts *bind.FilterOpts, liquidator []common.Address) (*SDUtilityPoolClaimedLiquidationIterator, error) {

	var liquidatorRule []interface{}
	for _, liquidatorItem := range liquidator {
		liquidatorRule = append(liquidatorRule, liquidatorItem)
	}

	logs, sub, err := _SDUtilityPool.contract.FilterLogs(opts, "ClaimedLiquidation", liquidatorRule)
	if err != nil {
		return nil, err
	}
	return &SDUtilityPoolClaimedLiquidationIterator{contract: _SDUtilityPool.contract, event: "ClaimedLiquidation", logs: logs, sub: sub}, nil
}

// WatchClaimedLiquidation is a free log subscription operation binding the contract event 0x87e8e54247cd0a3f2aae4999ad6838c60057f6de710e7b4eec301fe80eb55199.
//
// Solidity: event ClaimedLiquidation(address indexed liquidator, uint256 liquidationBonusInEth, uint256 liquidationFeeInEth)
func (_SDUtilityPool *SDUtilityPoolFilterer) WatchClaimedLiquidation(opts *bind.WatchOpts, sink chan<- *SDUtilityPoolClaimedLiquidation, liquidator []common.Address) (event.Subscription, error) {

	var liquidatorRule []interface{}
	for _, liquidatorItem := range liquidator {
		liquidatorRule = append(liquidatorRule, liquidatorItem)
	}

	logs, sub, err := _SDUtilityPool.contract.WatchLogs(opts, "ClaimedLiquidation", liquidatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SDUtilityPoolClaimedLiquidation)
				if err := _SDUtilityPool.contract.UnpackLog(event, "ClaimedLiquidation", log); err != nil {
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

// ParseClaimedLiquidation is a log parse operation binding the contract event 0x87e8e54247cd0a3f2aae4999ad6838c60057f6de710e7b4eec301fe80eb55199.
//
// Solidity: event ClaimedLiquidation(address indexed liquidator, uint256 liquidationBonusInEth, uint256 liquidationFeeInEth)
func (_SDUtilityPool *SDUtilityPoolFilterer) ParseClaimedLiquidation(log types.Log) (*SDUtilityPoolClaimedLiquidation, error) {
	event := new(SDUtilityPoolClaimedLiquidation)
	if err := _SDUtilityPool.contract.UnpackLog(event, "ClaimedLiquidation", log); err != nil {
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

// SDUtilityPoolInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the SDUtilityPool contract.
type SDUtilityPoolInitializedIterator struct {
	Event *SDUtilityPoolInitialized // Event containing the contract specifics and raw log

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
func (it *SDUtilityPoolInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SDUtilityPoolInitialized)
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
		it.Event = new(SDUtilityPoolInitialized)
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
func (it *SDUtilityPoolInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SDUtilityPoolInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SDUtilityPoolInitialized represents a Initialized event raised by the SDUtilityPool contract.
type SDUtilityPoolInitialized struct {
	Version uint8
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_SDUtilityPool *SDUtilityPoolFilterer) FilterInitialized(opts *bind.FilterOpts) (*SDUtilityPoolInitializedIterator, error) {

	logs, sub, err := _SDUtilityPool.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &SDUtilityPoolInitializedIterator{contract: _SDUtilityPool.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_SDUtilityPool *SDUtilityPoolFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *SDUtilityPoolInitialized) (event.Subscription, error) {

	logs, sub, err := _SDUtilityPool.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SDUtilityPoolInitialized)
				if err := _SDUtilityPool.contract.UnpackLog(event, "Initialized", log); err != nil {
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
func (_SDUtilityPool *SDUtilityPoolFilterer) ParseInitialized(log types.Log) (*SDUtilityPoolInitialized, error) {
	event := new(SDUtilityPoolInitialized)
	if err := _SDUtilityPool.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SDUtilityPoolLiquidationCallIterator is returned from FilterLiquidationCall and is used to iterate over the raw logs and unpacked data for LiquidationCall events raised by the SDUtilityPool contract.
type SDUtilityPoolLiquidationCallIterator struct {
	Event *SDUtilityPoolLiquidationCall // Event containing the contract specifics and raw log

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
func (it *SDUtilityPoolLiquidationCallIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SDUtilityPoolLiquidationCall)
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
		it.Event = new(SDUtilityPoolLiquidationCall)
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
func (it *SDUtilityPoolLiquidationCallIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SDUtilityPoolLiquidationCallIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SDUtilityPoolLiquidationCall represents a LiquidationCall event raised by the SDUtilityPool contract.
type SDUtilityPoolLiquidationCall struct {
	Account                     common.Address
	TotalLiquidationAmountInEth *big.Int
	LiquidationBonusInEth       *big.Int
	LiquidationFeeInEth         *big.Int
	Liquidator                  common.Address
	Raw                         types.Log // Blockchain specific contextual infos
}

// FilterLiquidationCall is a free log retrieval operation binding the contract event 0xcc9de8be9ac1f02b70a8ca2612f451a769d6d160ad91de17dcc38e54c567a532.
//
// Solidity: event LiquidationCall(address indexed account, uint256 totalLiquidationAmountInEth, uint256 liquidationBonusInEth, uint256 liquidationFeeInEth, address indexed liquidator)
func (_SDUtilityPool *SDUtilityPoolFilterer) FilterLiquidationCall(opts *bind.FilterOpts, account []common.Address, liquidator []common.Address) (*SDUtilityPoolLiquidationCallIterator, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	var liquidatorRule []interface{}
	for _, liquidatorItem := range liquidator {
		liquidatorRule = append(liquidatorRule, liquidatorItem)
	}

	logs, sub, err := _SDUtilityPool.contract.FilterLogs(opts, "LiquidationCall", accountRule, liquidatorRule)
	if err != nil {
		return nil, err
	}
	return &SDUtilityPoolLiquidationCallIterator{contract: _SDUtilityPool.contract, event: "LiquidationCall", logs: logs, sub: sub}, nil
}

// WatchLiquidationCall is a free log subscription operation binding the contract event 0xcc9de8be9ac1f02b70a8ca2612f451a769d6d160ad91de17dcc38e54c567a532.
//
// Solidity: event LiquidationCall(address indexed account, uint256 totalLiquidationAmountInEth, uint256 liquidationBonusInEth, uint256 liquidationFeeInEth, address indexed liquidator)
func (_SDUtilityPool *SDUtilityPoolFilterer) WatchLiquidationCall(opts *bind.WatchOpts, sink chan<- *SDUtilityPoolLiquidationCall, account []common.Address, liquidator []common.Address) (event.Subscription, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	var liquidatorRule []interface{}
	for _, liquidatorItem := range liquidator {
		liquidatorRule = append(liquidatorRule, liquidatorItem)
	}

	logs, sub, err := _SDUtilityPool.contract.WatchLogs(opts, "LiquidationCall", accountRule, liquidatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SDUtilityPoolLiquidationCall)
				if err := _SDUtilityPool.contract.UnpackLog(event, "LiquidationCall", log); err != nil {
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

// ParseLiquidationCall is a log parse operation binding the contract event 0xcc9de8be9ac1f02b70a8ca2612f451a769d6d160ad91de17dcc38e54c567a532.
//
// Solidity: event LiquidationCall(address indexed account, uint256 totalLiquidationAmountInEth, uint256 liquidationBonusInEth, uint256 liquidationFeeInEth, address indexed liquidator)
func (_SDUtilityPool *SDUtilityPoolFilterer) ParseLiquidationCall(log types.Log) (*SDUtilityPoolLiquidationCall, error) {
	event := new(SDUtilityPoolLiquidationCall)
	if err := _SDUtilityPool.contract.UnpackLog(event, "LiquidationCall", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SDUtilityPoolPausedIterator is returned from FilterPaused and is used to iterate over the raw logs and unpacked data for Paused events raised by the SDUtilityPool contract.
type SDUtilityPoolPausedIterator struct {
	Event *SDUtilityPoolPaused // Event containing the contract specifics and raw log

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
func (it *SDUtilityPoolPausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SDUtilityPoolPaused)
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
		it.Event = new(SDUtilityPoolPaused)
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
func (it *SDUtilityPoolPausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SDUtilityPoolPausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SDUtilityPoolPaused represents a Paused event raised by the SDUtilityPool contract.
type SDUtilityPoolPaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterPaused is a free log retrieval operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_SDUtilityPool *SDUtilityPoolFilterer) FilterPaused(opts *bind.FilterOpts) (*SDUtilityPoolPausedIterator, error) {

	logs, sub, err := _SDUtilityPool.contract.FilterLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return &SDUtilityPoolPausedIterator{contract: _SDUtilityPool.contract, event: "Paused", logs: logs, sub: sub}, nil
}

// WatchPaused is a free log subscription operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_SDUtilityPool *SDUtilityPoolFilterer) WatchPaused(opts *bind.WatchOpts, sink chan<- *SDUtilityPoolPaused) (event.Subscription, error) {

	logs, sub, err := _SDUtilityPool.contract.WatchLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SDUtilityPoolPaused)
				if err := _SDUtilityPool.contract.UnpackLog(event, "Paused", log); err != nil {
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
func (_SDUtilityPool *SDUtilityPoolFilterer) ParsePaused(log types.Log) (*SDUtilityPoolPaused, error) {
	event := new(SDUtilityPoolPaused)
	if err := _SDUtilityPool.contract.UnpackLog(event, "Paused", log); err != nil {
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

// SDUtilityPoolRiskConfigUpdatedIterator is returned from FilterRiskConfigUpdated and is used to iterate over the raw logs and unpacked data for RiskConfigUpdated events raised by the SDUtilityPool contract.
type SDUtilityPoolRiskConfigUpdatedIterator struct {
	Event *SDUtilityPoolRiskConfigUpdated // Event containing the contract specifics and raw log

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
func (it *SDUtilityPoolRiskConfigUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SDUtilityPoolRiskConfigUpdated)
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
		it.Event = new(SDUtilityPoolRiskConfigUpdated)
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
func (it *SDUtilityPoolRiskConfigUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SDUtilityPoolRiskConfigUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SDUtilityPoolRiskConfigUpdated represents a RiskConfigUpdated event raised by the SDUtilityPool contract.
type SDUtilityPoolRiskConfigUpdated struct {
	LiquidationThreshold    *big.Int
	LiquidationBonusPercent *big.Int
	LiquidationFeePercent   *big.Int
	Ltv                     *big.Int
	Raw                     types.Log // Blockchain specific contextual infos
}

// FilterRiskConfigUpdated is a free log retrieval operation binding the contract event 0x62dd46f943681bd727e339f4baaa19fb66ea209fbe7d9b4a75aa74c33acc18f3.
//
// Solidity: event RiskConfigUpdated(uint256 liquidationThreshold, uint256 liquidationBonusPercent, uint256 liquidationFeePercent, uint256 ltv)
func (_SDUtilityPool *SDUtilityPoolFilterer) FilterRiskConfigUpdated(opts *bind.FilterOpts) (*SDUtilityPoolRiskConfigUpdatedIterator, error) {

	logs, sub, err := _SDUtilityPool.contract.FilterLogs(opts, "RiskConfigUpdated")
	if err != nil {
		return nil, err
	}
	return &SDUtilityPoolRiskConfigUpdatedIterator{contract: _SDUtilityPool.contract, event: "RiskConfigUpdated", logs: logs, sub: sub}, nil
}

// WatchRiskConfigUpdated is a free log subscription operation binding the contract event 0x62dd46f943681bd727e339f4baaa19fb66ea209fbe7d9b4a75aa74c33acc18f3.
//
// Solidity: event RiskConfigUpdated(uint256 liquidationThreshold, uint256 liquidationBonusPercent, uint256 liquidationFeePercent, uint256 ltv)
func (_SDUtilityPool *SDUtilityPoolFilterer) WatchRiskConfigUpdated(opts *bind.WatchOpts, sink chan<- *SDUtilityPoolRiskConfigUpdated) (event.Subscription, error) {

	logs, sub, err := _SDUtilityPool.contract.WatchLogs(opts, "RiskConfigUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SDUtilityPoolRiskConfigUpdated)
				if err := _SDUtilityPool.contract.UnpackLog(event, "RiskConfigUpdated", log); err != nil {
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

// ParseRiskConfigUpdated is a log parse operation binding the contract event 0x62dd46f943681bd727e339f4baaa19fb66ea209fbe7d9b4a75aa74c33acc18f3.
//
// Solidity: event RiskConfigUpdated(uint256 liquidationThreshold, uint256 liquidationBonusPercent, uint256 liquidationFeePercent, uint256 ltv)
func (_SDUtilityPool *SDUtilityPoolFilterer) ParseRiskConfigUpdated(log types.Log) (*SDUtilityPoolRiskConfigUpdated, error) {
	event := new(SDUtilityPoolRiskConfigUpdated)
	if err := _SDUtilityPool.contract.UnpackLog(event, "RiskConfigUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SDUtilityPoolRoleAdminChangedIterator is returned from FilterRoleAdminChanged and is used to iterate over the raw logs and unpacked data for RoleAdminChanged events raised by the SDUtilityPool contract.
type SDUtilityPoolRoleAdminChangedIterator struct {
	Event *SDUtilityPoolRoleAdminChanged // Event containing the contract specifics and raw log

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
func (it *SDUtilityPoolRoleAdminChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SDUtilityPoolRoleAdminChanged)
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
		it.Event = new(SDUtilityPoolRoleAdminChanged)
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
func (it *SDUtilityPoolRoleAdminChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SDUtilityPoolRoleAdminChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SDUtilityPoolRoleAdminChanged represents a RoleAdminChanged event raised by the SDUtilityPool contract.
type SDUtilityPoolRoleAdminChanged struct {
	Role              [32]byte
	PreviousAdminRole [32]byte
	NewAdminRole      [32]byte
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterRoleAdminChanged is a free log retrieval operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_SDUtilityPool *SDUtilityPoolFilterer) FilterRoleAdminChanged(opts *bind.FilterOpts, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (*SDUtilityPoolRoleAdminChangedIterator, error) {

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

	logs, sub, err := _SDUtilityPool.contract.FilterLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return &SDUtilityPoolRoleAdminChangedIterator{contract: _SDUtilityPool.contract, event: "RoleAdminChanged", logs: logs, sub: sub}, nil
}

// WatchRoleAdminChanged is a free log subscription operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_SDUtilityPool *SDUtilityPoolFilterer) WatchRoleAdminChanged(opts *bind.WatchOpts, sink chan<- *SDUtilityPoolRoleAdminChanged, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (event.Subscription, error) {

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

	logs, sub, err := _SDUtilityPool.contract.WatchLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SDUtilityPoolRoleAdminChanged)
				if err := _SDUtilityPool.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
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
func (_SDUtilityPool *SDUtilityPoolFilterer) ParseRoleAdminChanged(log types.Log) (*SDUtilityPoolRoleAdminChanged, error) {
	event := new(SDUtilityPoolRoleAdminChanged)
	if err := _SDUtilityPool.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SDUtilityPoolRoleGrantedIterator is returned from FilterRoleGranted and is used to iterate over the raw logs and unpacked data for RoleGranted events raised by the SDUtilityPool contract.
type SDUtilityPoolRoleGrantedIterator struct {
	Event *SDUtilityPoolRoleGranted // Event containing the contract specifics and raw log

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
func (it *SDUtilityPoolRoleGrantedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SDUtilityPoolRoleGranted)
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
		it.Event = new(SDUtilityPoolRoleGranted)
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
func (it *SDUtilityPoolRoleGrantedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SDUtilityPoolRoleGrantedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SDUtilityPoolRoleGranted represents a RoleGranted event raised by the SDUtilityPool contract.
type SDUtilityPoolRoleGranted struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleGranted is a free log retrieval operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_SDUtilityPool *SDUtilityPoolFilterer) FilterRoleGranted(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*SDUtilityPoolRoleGrantedIterator, error) {

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

	logs, sub, err := _SDUtilityPool.contract.FilterLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &SDUtilityPoolRoleGrantedIterator{contract: _SDUtilityPool.contract, event: "RoleGranted", logs: logs, sub: sub}, nil
}

// WatchRoleGranted is a free log subscription operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_SDUtilityPool *SDUtilityPoolFilterer) WatchRoleGranted(opts *bind.WatchOpts, sink chan<- *SDUtilityPoolRoleGranted, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _SDUtilityPool.contract.WatchLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SDUtilityPoolRoleGranted)
				if err := _SDUtilityPool.contract.UnpackLog(event, "RoleGranted", log); err != nil {
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
func (_SDUtilityPool *SDUtilityPoolFilterer) ParseRoleGranted(log types.Log) (*SDUtilityPoolRoleGranted, error) {
	event := new(SDUtilityPoolRoleGranted)
	if err := _SDUtilityPool.contract.UnpackLog(event, "RoleGranted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SDUtilityPoolRoleRevokedIterator is returned from FilterRoleRevoked and is used to iterate over the raw logs and unpacked data for RoleRevoked events raised by the SDUtilityPool contract.
type SDUtilityPoolRoleRevokedIterator struct {
	Event *SDUtilityPoolRoleRevoked // Event containing the contract specifics and raw log

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
func (it *SDUtilityPoolRoleRevokedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SDUtilityPoolRoleRevoked)
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
		it.Event = new(SDUtilityPoolRoleRevoked)
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
func (it *SDUtilityPoolRoleRevokedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SDUtilityPoolRoleRevokedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SDUtilityPoolRoleRevoked represents a RoleRevoked event raised by the SDUtilityPool contract.
type SDUtilityPoolRoleRevoked struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleRevoked is a free log retrieval operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_SDUtilityPool *SDUtilityPoolFilterer) FilterRoleRevoked(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*SDUtilityPoolRoleRevokedIterator, error) {

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

	logs, sub, err := _SDUtilityPool.contract.FilterLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &SDUtilityPoolRoleRevokedIterator{contract: _SDUtilityPool.contract, event: "RoleRevoked", logs: logs, sub: sub}, nil
}

// WatchRoleRevoked is a free log subscription operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_SDUtilityPool *SDUtilityPoolFilterer) WatchRoleRevoked(opts *bind.WatchOpts, sink chan<- *SDUtilityPoolRoleRevoked, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _SDUtilityPool.contract.WatchLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SDUtilityPoolRoleRevoked)
				if err := _SDUtilityPool.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
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
func (_SDUtilityPool *SDUtilityPoolFilterer) ParseRoleRevoked(log types.Log) (*SDUtilityPoolRoleRevoked, error) {
	event := new(SDUtilityPoolRoleRevoked)
	if err := _SDUtilityPool.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
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

// SDUtilityPoolUnpausedIterator is returned from FilterUnpaused and is used to iterate over the raw logs and unpacked data for Unpaused events raised by the SDUtilityPool contract.
type SDUtilityPoolUnpausedIterator struct {
	Event *SDUtilityPoolUnpaused // Event containing the contract specifics and raw log

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
func (it *SDUtilityPoolUnpausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SDUtilityPoolUnpaused)
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
		it.Event = new(SDUtilityPoolUnpaused)
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
func (it *SDUtilityPoolUnpausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SDUtilityPoolUnpausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SDUtilityPoolUnpaused represents a Unpaused event raised by the SDUtilityPool contract.
type SDUtilityPoolUnpaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterUnpaused is a free log retrieval operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_SDUtilityPool *SDUtilityPoolFilterer) FilterUnpaused(opts *bind.FilterOpts) (*SDUtilityPoolUnpausedIterator, error) {

	logs, sub, err := _SDUtilityPool.contract.FilterLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return &SDUtilityPoolUnpausedIterator{contract: _SDUtilityPool.contract, event: "Unpaused", logs: logs, sub: sub}, nil
}

// WatchUnpaused is a free log subscription operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_SDUtilityPool *SDUtilityPoolFilterer) WatchUnpaused(opts *bind.WatchOpts, sink chan<- *SDUtilityPoolUnpaused) (event.Subscription, error) {

	logs, sub, err := _SDUtilityPool.contract.WatchLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SDUtilityPoolUnpaused)
				if err := _SDUtilityPool.contract.UnpackLog(event, "Unpaused", log); err != nil {
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
func (_SDUtilityPool *SDUtilityPoolFilterer) ParseUnpaused(log types.Log) (*SDUtilityPoolUnpaused, error) {
	event := new(SDUtilityPoolUnpaused)
	if err := _SDUtilityPool.contract.UnpackLog(event, "Unpaused", log); err != nil {
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

// SDUtilityPoolWithdrawnProtocolFeeIterator is returned from FilterWithdrawnProtocolFee and is used to iterate over the raw logs and unpacked data for WithdrawnProtocolFee events raised by the SDUtilityPool contract.
type SDUtilityPoolWithdrawnProtocolFeeIterator struct {
	Event *SDUtilityPoolWithdrawnProtocolFee // Event containing the contract specifics and raw log

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
func (it *SDUtilityPoolWithdrawnProtocolFeeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SDUtilityPoolWithdrawnProtocolFee)
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
		it.Event = new(SDUtilityPoolWithdrawnProtocolFee)
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
func (it *SDUtilityPoolWithdrawnProtocolFeeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SDUtilityPoolWithdrawnProtocolFeeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SDUtilityPoolWithdrawnProtocolFee represents a WithdrawnProtocolFee event raised by the SDUtilityPool contract.
type SDUtilityPoolWithdrawnProtocolFee struct {
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterWithdrawnProtocolFee is a free log retrieval operation binding the contract event 0xa32a67c2a472d948ee4cdcce3fdf2df7a9a99b762c69858f0651b1b22067efdf.
//
// Solidity: event WithdrawnProtocolFee(uint256 amount)
func (_SDUtilityPool *SDUtilityPoolFilterer) FilterWithdrawnProtocolFee(opts *bind.FilterOpts) (*SDUtilityPoolWithdrawnProtocolFeeIterator, error) {

	logs, sub, err := _SDUtilityPool.contract.FilterLogs(opts, "WithdrawnProtocolFee")
	if err != nil {
		return nil, err
	}
	return &SDUtilityPoolWithdrawnProtocolFeeIterator{contract: _SDUtilityPool.contract, event: "WithdrawnProtocolFee", logs: logs, sub: sub}, nil
}

// WatchWithdrawnProtocolFee is a free log subscription operation binding the contract event 0xa32a67c2a472d948ee4cdcce3fdf2df7a9a99b762c69858f0651b1b22067efdf.
//
// Solidity: event WithdrawnProtocolFee(uint256 amount)
func (_SDUtilityPool *SDUtilityPoolFilterer) WatchWithdrawnProtocolFee(opts *bind.WatchOpts, sink chan<- *SDUtilityPoolWithdrawnProtocolFee) (event.Subscription, error) {

	logs, sub, err := _SDUtilityPool.contract.WatchLogs(opts, "WithdrawnProtocolFee")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SDUtilityPoolWithdrawnProtocolFee)
				if err := _SDUtilityPool.contract.UnpackLog(event, "WithdrawnProtocolFee", log); err != nil {
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

// ParseWithdrawnProtocolFee is a log parse operation binding the contract event 0xa32a67c2a472d948ee4cdcce3fdf2df7a9a99b762c69858f0651b1b22067efdf.
//
// Solidity: event WithdrawnProtocolFee(uint256 amount)
func (_SDUtilityPool *SDUtilityPoolFilterer) ParseWithdrawnProtocolFee(log types.Log) (*SDUtilityPoolWithdrawnProtocolFee, error) {
	event := new(SDUtilityPoolWithdrawnProtocolFee)
	if err := _SDUtilityPool.contract.UnpackLog(event, "WithdrawnProtocolFee", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
