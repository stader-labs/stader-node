package validator

import (
	"github.com/urfave/cli"

	"github.com/stader-labs/stader-node/shared/services"
	"github.com/stader-labs/stader-node/shared/types/api"
)

func getStatus(c *cli.Context) (*api.MinipoolStatusResponse, error) {

	// Get services
	if err := services.RequireNodeRegistered(c); err != nil {
		return nil, err
	}
	if err := services.RequireBeaconClientSynced(c); err != nil {
		return nil, err
	}
	w, err := services.GetWallet(c)
	if err != nil {
		return nil, err
	}
	rp, err := services.GetRocketPool(c)
	if err != nil {
		return nil, err
	}
	bc, err := services.GetBeaconClient(c)
	if err != nil {
		return nil, err
	}

	// Response
	response := api.MinipoolStatusResponse{}

	// Get minipool details
	nodeAccount, err := w.GetNodeAccount()
	if err != nil {
		return nil, err
	}
	details, err := getNodeMinipoolDetails(rp, bc, nodeAccount.Address)
	if err != nil {
		return nil, err
	}
	response.Minipools = details

	// Return response
	return &response, nil

}
