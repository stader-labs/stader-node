package sdutility

import (
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/stader-labs/stader-node/stader-lib/stader"
)

func GetUtilizerLatestBalance(sp *stader.SDUtilityPoolContractManager, address common.Address, opts *bind.CallOpts) (*big.Int, error) {
	return sp.SDUtilityPool.GetUtilizerLatestBalance(opts, address)
}

func GetPoolAvailableSDBalance(sp *stader.SDUtilityPoolContractManager, opts *bind.CallOpts) (*big.Int, error) {
	return sp.SDUtilityPool.GetPoolAvailableSDBalance(opts)
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

func GetUtilizationRate(sp *stader.SDUtilityPoolContractManager, opts *bind.CallOpts) (*big.Float, error) {
	utilizationRatePerBlockInWei, err := sp.SDUtilityPool.UtilizationRatePerBlock(opts)
	if err != nil {
		return nil, err
	}

	utilizationRatePerYear := new(big.Int).Mul(utilizationRatePerBlockInWei, big.NewInt(2628000)) // 2628000 block per year

	utilizationRatePerYearF := new(big.Float).SetInt(utilizationRatePerYear)

	utilizationRateInPercent := new(big.Float).Quo(utilizationRatePerYearF, big.NewFloat(1e16))

	return utilizationRateInPercent, nil
}
