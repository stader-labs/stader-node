package node

import (
	"fmt"

	"github.com/urfave/cli"

	"github.com/stader-labs/stader-node/shared/services"
	"github.com/stader-labs/stader-node/shared/types/api"
)

func getContractsInfo(c *cli.Context) (*api.ContractsInfoResponse, error) {
	// Response
	response := api.ContractsInfoResponse{}

	// Get the ETH1 network ID
	config, err := services.GetConfig(c)
	if err != nil {
		return nil, fmt.Errorf("Error getting configuration: %w", err)
	}
	response.Network = uint64(config.StaderNode.GetChainID())

	// Get the Beacon Client info
	bc, err := services.GetBeaconClient(c)
	if err != nil {
		return nil, fmt.Errorf("Error getting beacon client: %w", err)
	}

	eth2DepositContract, err := bc.GetEth2DepositContract()
	if err != nil {
		return nil, fmt.Errorf("Error getting beacon client deposit contract: %w", err)
	}

	response.BeaconNetwork = eth2DepositContract.ChainID
	response.BeaconDepositContract = eth2DepositContract.Address
	response.SdCollateralContract = config.StaderNode.GetSdCollateralContractAddress()
	response.EthxToken = config.StaderNode.GetEthxTokenAddress()
	response.SdToken = config.StaderNode.GetSdTokenAddress()
	response.PermissionlessNodeRegistry = config.StaderNode.GetPermissionlessNodeRegistryAddress()
	response.VaultFactory = config.StaderNode.GetVaultFactoryAddress()
	response.SocializingPoolContract = config.StaderNode.GetSocializingPoolAddress()
	response.PermisionlessPool = config.StaderNode.GetPermissionlessPoolAddress()
	response.StaderOracle = config.StaderNode.GetStaderOracleAddress()

	// Return response
	return &response, nil

}
