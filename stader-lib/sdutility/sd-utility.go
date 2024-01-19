package sdutility

import (
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/stader-labs/stader-node/stader-lib/contracts"
	"github.com/stader-labs/stader-node/stader-lib/stader"
)

func GetUtilizerLatestBalance(sp *stader.SDUtilityPoolContractManager, address common.Address, opts *bind.CallOpts) (*big.Int, error) {
	return sp.SDUtilityPool.GetUtilizerLatestBalance(opts, address)
}

func GetPoolAvailableSDBalance(sp *stader.SDUtilityPoolContractManager, opts *bind.CallOpts) (*big.Int, error) {
	accumulatedProtocolFee, err := GetAccumulatedProtocolFee(sp, opts)
	if err != nil {
		return nil, err
	}
	sdRequestedForWithdraw, err := GetSdRequestedForWithdraw(sp, opts)
	if err != nil {
		return nil, err
	}
	utilityPoolBalance, err := GetUtilityPoolBalance(sp, opts)
	if err != nil {
		return nil, err
	}

	utilityPoolBalanceMinusSdForWithdraw := big.NewInt(0).Sub(utilityPoolBalance, sdRequestedForWithdraw)
	if utilityPoolBalanceMinusSdForWithdraw.Cmp(big.NewInt(0)) < 0 {
		return big.NewInt(0), nil
	}

	availableSdBalance := big.NewInt(0).Sub(utilityPoolBalanceMinusSdForWithdraw, accumulatedProtocolFee)
	if availableSdBalance.Cmp(big.NewInt(0)) < 0 {
		return big.NewInt(0), nil
	}

	return availableSdBalance, nil
}

func GetUtilityPoolBalance(sp *stader.SDUtilityPoolContractManager, opts *bind.CallOpts) (*big.Int, error) {
	return sp.SDUtilityPool.GetPoolAvailableSDBalance(opts)
}

func GetAccumulatedProtocolFee(sp *stader.SDUtilityPoolContractManager, opts *bind.CallOpts) (*big.Int, error) {
	return sp.SDUtilityPool.AccumulatedProtocolFee(opts)
}

func GetSdRequestedForWithdraw(sp *stader.SDUtilityPoolContractManager, opts *bind.CallOpts) (*big.Int, error) {
	return sp.SDUtilityPool.SdRequestedForWithdraw(opts)
}

// Estimate the gas of Utilize
func EstimateUtilize(sp *stader.SDUtilityPoolContractManager, utilityAmount *big.Int, opts *bind.TransactOpts) (stader.GasInfo, error) {
	return sp.SDUtilityPoolContract.GetTransactionGasInfo(opts, "utilize", utilityAmount)
}

func Utilize(sp *stader.SDUtilityPoolContractManager, utilityAmount *big.Int, opts *bind.TransactOpts) (*types.Transaction, error) {
	return sp.SDUtilityPool.Utilize(opts, utilityAmount)
}

func EstimateRepay(sp *stader.SDUtilityPoolContractManager, utilityAmount *big.Int, opts *bind.TransactOpts) (stader.GasInfo, error) {
	return sp.SDUtilityPoolContract.GetTransactionGasInfo(opts, "repay", utilityAmount)
}

func Repay(sp *stader.SDUtilityPoolContractManager, utilityAmount *big.Int, opts *bind.TransactOpts) (*types.Transaction, error) {
	return sp.SDUtilityPool.Repay(opts, utilityAmount)
}

func SDMaxUtilizableAmount(sp *stader.SDUtilityPoolContractManager,
	sdc *stader.SdCollateralContractManager, numValidators *big.Int, opts *bind.CallOpts) (*big.Int, error) {
	maxThreshold, err := sp.SDUtilityPool.MaxETHWorthOfSDPerValidator(opts)
	ethAmount := new(big.Int).Mul(maxThreshold, numValidators)

	sdAmount, err := sdc.SdCollateral.ConvertETHToSD(opts, ethAmount)
	if err != nil {
		return nil, err
	}

	return sdAmount, nil
}

func GetUtilizationRatePercent(sp *stader.SDUtilityPoolContractManager, opts *bind.CallOpts) (*big.Float, error) {
	utilizationRatePerBlockInWei, err := sp.SDUtilityPool.UtilizationRatePerBlock(opts)
	if err != nil {
		return nil, err
	}

	utilizationRatePerYear := new(big.Int).Mul(utilizationRatePerBlockInWei, big.NewInt(2628000)) // 2628000 block per year

	utilizationRatePerYearF := new(big.Float).SetInt(utilizationRatePerYear)

	utilizationRateInPercent := new(big.Float).Quo(utilizationRatePerYearF, big.NewFloat(1e16))

	return utilizationRateInPercent, nil
}

func GetUserData(sp *stader.SDUtilityPoolContractManager, address common.Address, opts *bind.CallOpts) (*contracts.UserData, error) {
	userData, err := sp.SDUtilityPool.GetUserData(opts, address)
	if err != nil {
		return nil, err
	}

	return &userData, nil
}

func AlreadyLiquidated(sp *stader.SDUtilityPoolContractManager, address common.Address, opts *bind.CallOpts) (bool, error) {
	liquidationIndex, err := LiquidationIndexByOperator(sp, address, nil)
	if err != nil {
		return false, err
	}

	isAlreadyLiquidated := liquidationIndex.Cmp(big.NewInt(0)) != 0

	return isAlreadyLiquidated, nil
}

func LiquidationIndexByOperator(sp *stader.SDUtilityPoolContractManager, address common.Address, opts *bind.CallOpts) (*big.Int, error) {
	return sp.SDUtilityPool.LiquidationIndexByOperator(opts, address)
}
