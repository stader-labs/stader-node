package node

import (
	"fmt"

	"github.com/urfave/cli"

	"github.com/stader-labs/stader-node/shared/services"
	"github.com/stader-labs/stader-node/shared/types/api"
)

func getContractsInfo(c *cli.Context) (*api.ContractsInfoResponse, error) {
	prn, err := services.GetPermissionlessNodeRegistry(c)
	if err != nil {
		return nil, err
	}
	vf, err := services.GetVaultFactory(c)
	if err != nil {
		return nil, err
	}
	sdc, err := services.GetSdCollateralContract(c)
	if err != nil {
		return nil, err
	}
	ethx, err := services.GetEthxTokenContract(c)
	if err != nil {
		return nil, err
	}
	sdt, err := services.GetSdTokenContract(c)
	if err != nil {
		return nil, err
	}

	// Response
	response := api.ContractsInfoResponse{}

	// Get the ETH1 network ID
	fmt.Printf("Getting ETH1 Network ID!")
	config, err := services.GetConfig(c)
	if err != nil {
		return nil, fmt.Errorf("Error getting configuration: %w", err)
	}
	response.Network = uint64(config.Smartnode.GetChainID())
	fmt.Printf("Got Eth1 Network Id: %s\n", response.Network)

	// Get the Beacon Client info
	fmt.Printf("Get Beacon client info!")
	bc, err := services.GetBeaconClient(c)
	if err != nil {
		return nil, fmt.Errorf("Error getting beacon client: %w", err)
	}
	fmt.Printf("Got Beacon client info!")
	fmt.Printf("Getting eth2 deposit contract!")
	eth2DepositContract, err := bc.GetEth2DepositContract()
	if err != nil {
		return nil, fmt.Errorf("Error getting beacon client deposit contract: %w", err)
	}
	fmt.Printf("Got eth2DepositContract: %s\n", eth2DepositContract)

	response.BeaconNetwork = eth2DepositContract.ChainID
	response.BeaconDepositContract = eth2DepositContract.Address
	response.SdCollateralContract = *sdc.SdCollateralContract.Address
	response.EthxToken = *ethx.Erc20TokenContract.Address
	response.SdToken = *sdt.Erc20TokenContract.Address
	response.PermissionlessNodeRegistry = *prn.PermissionlessNodeRegistryContract.Address
	response.VaultFactory = *vf.VaultFactoryContract.Address

	// Return response
	return &response, nil

}
