package sd_collateral

import (
	"fmt"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/stader-labs/stader-node/stader-lib/stader"
	types2 "github.com/stader-labs/stader-node/stader-lib/types"
	"math/big"
)

func EstimateDepositSdAsCollateral(sdc *stader.SdCollateralContractManager, amount *big.Int, opts *bind.TransactOpts) (stader.GasInfo, error) {
	return sdc.SdCollateralContract.GetTransactionGasInfo(opts, "depositSDAsCollateral", amount)
}

func DepositSdAsCollateral(sdc *stader.SdCollateralContractManager, amount *big.Int, opts *bind.TransactOpts) (*types.Transaction, error) {
	tx, err := sdc.SdCollateral.DepositSDAsCollateral(opts, amount)
	if err != nil {
		return nil, err
	}

	return tx, nil
}

func EstimateRequestSdCollateralWithdraw(sdc *stader.SdCollateralContractManager, amount *big.Int, opts *bind.TransactOpts) (stader.GasInfo, error) {
	return sdc.SdCollateralContract.GetTransactionGasInfo(opts, "requestWithdraw", amount)
}

func RequestSdCollateralWithdraw(sdc *stader.SdCollateralContractManager, amount *big.Int, opts *bind.TransactOpts) (*types.Transaction, error) {
	tx, err := sdc.SdCollateral.RequestWithdraw(opts, amount)
	if err != nil {
		return nil, err
	}

	return tx, nil
}

func EstimateClaimWithdrawnSd(sdc *stader.SdCollateralContractManager, opts *bind.TransactOpts) (stader.GasInfo, error) {
	return sdc.SdCollateralContract.GetTransactionGasInfo(opts, "claimWithdraw")
}

func ClaimWithdrawnSd(sdc *stader.SdCollateralContractManager, opts *bind.TransactOpts) (*types.Transaction, error) {
	tx, err := sdc.SdCollateral.ClaimWithdraw(opts)
	if err != nil {
		return nil, err
	}

	return tx, nil
}

func GetOperatorSdBalance(sdc *stader.SdCollateralContractManager, operatorAddress common.Address, opts *bind.CallOpts) (*big.Int, error) {
	balance, err := sdc.SdCollateral.OperatorSDBalance(opts, operatorAddress)
	if err != nil {
		return nil, err
	}

	return balance, err
}

func HasEnoughSdCollateral(sdc *stader.SdCollateralContractManager, operatorAddress common.Address, poolType uint8, numValidators *big.Int, opts *bind.CallOpts) (bool, error) {
	hasEnoughSdCollateral, err := sdc.SdCollateral.HasEnoughSDCollateral(opts, operatorAddress, poolType, numValidators)
	if err != nil {
		return false, err
	}

	return hasEnoughSdCollateral, nil
}

func GetMaxValidatorSpawnable(sdc *stader.SdCollateralContractManager, sdAmount *big.Int, poolType uint8, opts *bind.CallOpts) (*big.Int, error) {
	pThreshold, err := sdc.SdCollateral.PoolThresholdbyPoolId(opts, poolType)
	if err != nil {
		return nil, err
	}

	ethAmount, err := sdc.SdCollateral.ConvertSDToETH(opts, sdAmount)
	if err != nil {
		return nil, err
	}

	if pThreshold.MinThreshold.Cmp(big.NewInt(0)) == 0 {
		return nil, fmt.Errorf("pool min threshold is 0")
	}

	return ethAmount.Div(ethAmount, pThreshold.MinThreshold), nil
}

func ConvertEthToSd(sdc *stader.SdCollateralContractManager, ethAmount *big.Int, opts *bind.CallOpts) (*big.Int, error) {
	sdAmount, err := sdc.SdCollateral.ConvertETHToSD(opts, ethAmount)
	if err != nil {
		return nil, err
	}

	return sdAmount, nil
}

func ConvertSdToEth(sdc *stader.SdCollateralContractManager, sdAmount *big.Int, opts *bind.CallOpts) (*big.Int, error) {
	ethAmount, err := sdc.SdCollateral.ConvertSDToETH(opts, sdAmount)
	if err != nil {
		return nil, err
	}

	return ethAmount, nil
}

func GetTotalSdCollateral(sdc *stader.SdCollateralContractManager, opts *bind.CallOpts) (*big.Int, error) {
	totalSdCollateral, err := sdc.SdCollateral.TotalSDCollateral(opts)
	if err != nil {
		return nil, err
	}

	return totalSdCollateral, nil
}

func GetWithdrawDelay(sdc *stader.SdCollateralContractManager, opts *bind.CallOpts) (*big.Int, error) {
	withdrawDelay, err := sdc.SdCollateral.WithdrawDelay(opts)
	if err != nil {
		return nil, err
	}

	return withdrawDelay, nil
}

func GetOperatorWithdrawInfo(sdc *stader.SdCollateralContractManager, operatorAddress common.Address, opts *bind.CallOpts) (types2.OperatorWithdrawInfo, error) {
	withdrawInfo, err := sdc.SdCollateral.WithdrawReq(opts, operatorAddress)
	if err != nil {
		return types2.OperatorWithdrawInfo{}, err
	}

	return withdrawInfo, nil
}

func GetPoolThreshold(sdc *stader.SdCollateralContractManager, poolType uint8, opts *bind.CallOpts) (types2.PoolThresholdInfo, error) {
	poolThreshold, err := sdc.SdCollateral.PoolThresholdbyPoolId(opts, poolType)
	if err != nil {
		return types2.PoolThresholdInfo{}, err
	}

	return poolThreshold, nil
}
