package node

import (
	"math/big"

	"github.com/stader-labs/stader-node/stader-lib/sdutility"

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

	// Get gas estimates
	opts, err := w.GetNodeAccountTransactor()
	if err != nil {
		return nil, err
	}

	gasInfo, err := sdutility.EstimateRepay(sdu, amountWei, opts)
	if err != nil {
		return nil, err
	}

	sdStatus, err := getSDStatus(c, big.NewInt(0))
	if err != nil {
		return nil, err
	}

	response.SdStatusResponse = sdStatus.SDStatus

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

	// Response
	response := api.NodeRepaySDResponse{}

	tx, err := sdutility.Repay(sdu, amountWei, opts)
	if err != nil {
		return nil, err
	}

	response.TxHash = tx.Hash()

	return &response, nil
}
