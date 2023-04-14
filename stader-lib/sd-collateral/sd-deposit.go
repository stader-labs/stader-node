package sd_collateral

import (
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/stader-labs/stader-node/stader-lib/stader"
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
	maxValidatorSpawanable, err := sdc.SdCollateral.GetMaxValidatorSpawnable(opts, sdAmount, poolType)
	if err != nil {
		return nil, err
	}

	return maxValidatorSpawanable, nil
}

func ConvertEthToSd(sdc *stader.SdCollateralContractManager, ethAmount *big.Int, opts *bind.CallOpts) (*big.Int, error) {
	sdAmount, err := sdc.SdCollateral.ConvertETHToSD(opts, ethAmount)
	if err != nil {
		return nil, err
	}

	return sdAmount, nil
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
