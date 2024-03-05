package node

import (
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/stader-labs/stader-node/stader-lib/sdutility"
	"github.com/stader-labs/stader-node/stader-lib/stader"

	"github.com/urfave/cli"

	"github.com/stader-labs/stader-node/shared/services"
	"github.com/stader-labs/stader-node/shared/types/api"
)

func canRepaySD(c *cli.Context, amountWei *big.Int) (*api.CanRepaySDResponse, error) {
	// Get services
	if err := services.RequireNodeWallet(c); err != nil {
		return nil, err
	}

	w, err := services.GetWallet(c)
	if err != nil {
		return nil, err
	}

	// Response
	response := api.CanRepaySDResponse{}

	sdu, err := services.GetSdUtilityContract(c)
	if err != nil {
		return nil, err
	}

	isAmountMaxUint256 := amountWei.Cmp(abi.MaxUint256) == 0

	// Get gas estimates
	opts, err := w.GetNodeAccountTransactor()
	if err != nil {
		return nil, err
	}

	var gasInfo stader.GasInfo

	if isAmountMaxUint256 {
		gasInfo, err = sdutility.EstimateRepayFullAmount(sdu, opts)
		if err != nil {
			return nil, err
		}
	} else {
		gasInfo, err = sdutility.EstimateRepay(sdu, amountWei, opts)
		if err != nil {
			return nil, err
		}
	}

	response.GasInfo = gasInfo

	return &response, nil
}

func repaySD(c *cli.Context, amountWei *big.Int) (*api.NodeRepaySDResponse, error) {
	// Get services
	if err := services.RequireNodeWallet(c); err != nil {
		return nil, err
	}

	w, err := services.GetWallet(c)
	if err != nil {
		return nil, err
	}

	sdu, err := services.GetSdUtilityContract(c)
	if err != nil {
		return nil, err
	}

	// Get gas estimates
	opts, err := w.GetNodeAccountTransactor()
	if err != nil {
		return nil, err
	}

	isAmountMaxUint256 := amountWei.Cmp(abi.MaxUint256) == 0

	// Response
	response := api.NodeRepaySDResponse{}

	var tx *types.Transaction
	if isAmountMaxUint256 {
		tx, err = sdutility.RepayFullAmount(sdu, opts)
		if err != nil {
			return nil, err
		}
	} else {
		tx, err = sdutility.Repay(sdu, amountWei, opts)
		if err != nil {
			return nil, err
		}
	}

	response.TxHash = tx.Hash()

	return &response, nil
}
