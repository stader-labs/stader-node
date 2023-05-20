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

	fmt.Printf("eth2DepositContract: %+v\n", eth2DepositContract)
	response.BeaconNetwork = eth2DepositContract.ChainID
	response.BeaconDepositContract = eth2DepositContract.Address
	fmt.Printf("Getting sd collateral address\n")
	response.SdCollateralContract, err = services.GetSdCollateralAddress(c)
	if err != nil {
		return nil, err
	}
	fmt.Printf("Getting ethx token address\n")
	response.EthxToken, err = services.GetEthxTokenAddress(c)
	if err != nil {
		return nil, err
	}
	fmt.Printf("Getting sd token address\n")
	response.SdToken, err = services.GetSdTokenAddress(c)
	if err != nil {
		return nil, err
	}
	fmt.Printf("Getting permissionless node registry address\n")
	response.PermissionlessNodeRegistry, err = services.GetPermissionlessNodeRegistryAddress(c)
	if err != nil {
		return nil, err
	}
	fmt.Printf("Getting vault factory address\n")
	response.VaultFactory, err = services.GetVaultFactoryAddress(c)
	if err != nil {
		return nil, err
	}
	fmt.Printf("Getting socializing pool address\n")
	response.SocializingPoolContract, err = services.GetSocializingPoolAddress(c)
	if err != nil {
		return nil, err
	}
	fmt.Printf("Getting permissionless pool address\n")
	response.PermisionlessPool, err = services.GetPermissionlessPoolAddress(c)
	if err != nil {
		return nil, err
	}
	fmt.Printf("Getting stader oracle address\n")
	response.StaderOracle, err = services.GetStaderOracleAddress(c)
	if err != nil {
		return nil, err
	}
	fmt.Printf("Getting stake pool manager addres\n")
	response.StakePoolManager, err = services.GetStakePoolManagerAddress(c)
	if err != nil {
		return nil, err
	}

	// Return response
	return &response, nil

}
