package node

import (
	"github.com/stader-labs/stader-minipool-go/node"
	"github.com/stader-labs/stader-minipool-go/tokens"
	"github.com/stader-labs/stader-node/shared/services"
	"github.com/stader-labs/stader-node/shared/types/api"
	"github.com/urfave/cli"
)

func getStatus(c *cli.Context) (*api.NodeStatusResponse, error) {

	// Get services
	if err := services.RequireNodeWallet(c); err != nil {
		return nil, err
	}
	w, err := services.GetWallet(c)
	if err != nil {
		return nil, err
	}
	sor, err := services.GetStaderOperatorRegistry(c)
	if err != nil {
		return nil, err
	}

	nodeAccount, err := w.GetNodeAccount()
	if err != nil {
		return nil, err
	}

	//rp, err := services.GetRocketPool(c)
	//if err != nil {
	//	return nil, err
	//}
	//bc, err := services.GetBeaconClient(c)
	//if err != nil {
	//	return nil, err
	//}
	//s, err := services.GetSnapshotDelegation(c)
	//if err != nil {
	//	return nil, err
	//}

	// Response
	response := api.NodeStatusResponse{}

	operatorRegistryInfo, err := node.GetOperatorRegistry(sor, nodeAccount.Address, nil)
	if err != nil {
		return nil, err
	}
	if operatorRegistryInfo.OperatorName == "" {
		response.Registered = false
		return &response, nil
	}

	response.AccountAddress = nodeAccount.Address
	response.AccountAddressFormatted = formatResolvedAddress(c, response.AccountAddress)
	response.OperatorName = operatorRegistryInfo.OperatorName
	response.OperatorId = operatorRegistryInfo.OperatorId
	response.OperatorRewardAddress = operatorRegistryInfo.OperatorRewardAddress
	response.EthBalance, err = tokens.GetAccountEthBalance(sor.StaderOperatorRegistryContract, nodeAccount.Address, nil)
	if err != nil {
		return nil, err
	}

	response.Registered = true

	// Return response
	return &response, nil
}
