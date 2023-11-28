package node

import (
	"math/big"

	"github.com/urfave/cli"

	"github.com/stader-labs/stader-node/shared/services"
	"github.com/stader-labs/stader-node/shared/types/api"
	"github.com/stader-labs/stader-node/stader-lib/sd_utility"
)

func utilitySd(c *cli.Context, amountWei *big.Int) (*api.NodeUtilitySDResponse, error) {

	return &api.NodeUtilitySDResponse{}, nil
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
	response := api.NodeUtilitySDResponse{}

	tx, err := sd_utility.Utilize(sdu, amountWei, opts)
	if err != nil {
		return nil, err
	}

	response.TxHash = tx.Hash()

	return &response, nil

}

func canUtilitySd(c *cli.Context, amountWei *big.Int) (*api.CanUtilitySDResponse, error) {

	return &api.CanUtilitySDResponse{
		Error: "utility err",
	}, nil

	// Get services
	if err := services.RequireNodeWallet(c); err != nil {
		return nil, err
	}

	w, err := services.GetWallet(c)
	if err != nil {
		return nil, err
	}

	// Response
	response := api.CanUtilitySDResponse{}

	sdu, err := services.GetSdUtilityContract(c)
	if err != nil {
		return nil, err
	}

	// Get gas estimates
	opts, err := w.GetNodeAccountTransactor()
	if err != nil {
		return nil, err
	}

	gasInfo, err := sd_utility.EstimateUtilize(sdu, amountWei, opts)
	if err != nil {
		return nil, err
	}

	response.GasInfo = gasInfo

	return &response, nil

}
