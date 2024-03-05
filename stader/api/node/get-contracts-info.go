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
	response.EncryptionKey = config.StaderNode.GetPresignEncryptionKey()

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

	response.SdCollateralContract, err = services.GetSdCollateralAddress(c)
	if err != nil {
		return nil, err
	}
	response.StaderConfig = config.StaderNode.GetStaderConfigAddress()

	response.EthxToken, err = services.GetEthxTokenAddress(c)
	if err != nil {
		return nil, err
	}

	response.SdToken, err = services.GetSdTokenAddress(c)
	if err != nil {
		return nil, err
	}

	response.PermissionlessNodeRegistry, err = services.GetPermissionlessNodeRegistryAddress(c)
	if err != nil {
		return nil, err
	}

	response.VaultFactory, err = services.GetVaultFactoryAddress(c)
	if err != nil {
		return nil, err
	}

	response.SocializingPoolContract, err = services.GetSocializingPoolAddress(c)
	if err != nil {
		return nil, err
	}

	response.PermisionlessPool, err = services.GetPermissionlessPoolAddress(c)
	if err != nil {
		return nil, err
	}

	response.StaderOracle, err = services.GetStaderOracleAddress(c)
	if err != nil {
		return nil, err
	}

	response.StakePoolManager, err = services.GetStakePoolManagerAddress(c)
	if err != nil {
		return nil, err
	}

	response.SdUtilityContract, err = services.GetSdUtilityAddress(c)
	if err != nil {
		return nil, err
	}

	// Return response
	return &response, nil

}
