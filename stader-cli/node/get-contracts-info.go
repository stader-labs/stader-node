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
	"github.com/stader-labs/stader-node/shared/services/stader"
	cliutils "github.com/stader-labs/stader-node/shared/utils/cli"
	"github.com/stader-labs/stader-node/shared/utils/log"
	"github.com/urfave/cli"
)

func getContractsInfo(c *cli.Context) error {
	staderClient, err := stader.NewClientFromCtx(c)
	if err != nil {
		return err
	}
	defer staderClient.Close()

	// Check and assign the EC status
	err = cliutils.CheckClientStatus(staderClient)
	if err != nil {
		return err
	}

	// Get node fee
	response, err := staderClient.GetContractsInfo()
	if err != nil {
		return err
	}

	fmt.Printf("%s=== Beacon Network Contract Details ===%s\n", log.ColorGreen, log.ColorReset)
	fmt.Printf("Beacon Network: %d\n\n", response.BeaconNetwork)
	fmt.Printf("Beacon Deposit Contract: %s\n\n", response.BeaconDepositContract)

	fmt.Printf("%s=== Stader Network Contract Details ===%s\n", log.ColorGreen, log.ColorReset)
	fmt.Printf("Network: %d\n\n", response.Network)
	fmt.Printf("Permissionless Node Registry: %s\n\n", response.PermissionlessNodeRegistry)
	fmt.Printf("Vault Factory: %s\n\n", response.VaultFactory)
	fmt.Printf("Sd Collateral Lock: %s\n\n", response.SdCollateralContract)
	fmt.Printf("EthX Token: %s\n\n", response.EthxToken)
	fmt.Printf("Sd Token: %s\n\n", response.SdToken)

	return nil
}
