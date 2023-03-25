/*
This work is licensed and released under GNU GPL v3 or any other later versions. 
The full text of the license is below/ found at <http://www.gnu.org/licenses/>

(c) 2023 Rocket Pool Pty Ltd. Modified under GNU GPL v3.

This program is free software: you can redistribute it and/or modify
it under the terms of the GNU General Public License as published by
the Free Software Foundation, either version 3 of the License, or
(at your option) any later version.

This program is distributed in the hope that it will be useful,
but WITHOUT ANY WARRANTY; without even the implied warranty of
MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
GNU General Public License for more details.

You should have received a copy of the GNU General Public License
along with this program.  If not, see <http://www.gnu.org/licenses/>.
*/
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
	response.SdCollateralContract = *sdc.SdCollateralContract.Address
	response.EthxToken = *ethx.Erc20TokenContract.Address
	response.SdToken = *sdt.Erc20TokenContract.Address
	response.PermissionlessNodeRegistry = *prn.PermissionlessNodeRegistryContract.Address
	response.VaultFactory = *vf.VaultFactoryContract.Address

	// Return response
	return &response, nil

}
