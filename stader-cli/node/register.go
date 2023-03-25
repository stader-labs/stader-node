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

	"github.com/ethereum/go-ethereum/common"
	"github.com/urfave/cli"

	"github.com/stader-labs/stader-node/shared/services/gas"
	"github.com/stader-labs/stader-node/shared/services/stader"
	cliutils "github.com/stader-labs/stader-node/shared/utils/cli"
)

func registerNode(c *cli.Context) error {

	staderClient, err := stader.NewClientFromCtx(c)
	if err != nil {
		return err
	}
	defer staderClient.Close()

	walletStatus, err := staderClient.WalletStatus()
	if err != nil {
		return err
	}

	// Check and assign the EC status
	err = cliutils.CheckClientStatus(staderClient)
	if err != nil {
		return err
	}

	operatorName := c.String("operator-name")
	operatorRewardAddressString := c.String("operator-reward-address")
	if operatorRewardAddressString == "" {
		operatorRewardAddressString = walletStatus.AccountAddress.String()
	}
	socializeMev := c.String("socialize-mev")
	// default socialize mev to true
	if socializeMev == "" {
		socializeMev = "true"
	}
	socializeMevBool := parseToBool(socializeMev)

	fmt.Printf("cli: Register-Node: operator reward address is %s\n\n", common.HexToAddress(operatorRewardAddressString))
	fmt.Printf("cli: Register-Node: socializeMevBool is %d\n", socializeMevBool)

	// Check node can be registered
	canRegister, err := staderClient.CanRegisterNode(operatorName, common.HexToAddress(operatorRewardAddressString), socializeMevBool)
	if err != nil {
		return err
	}
	if !canRegister.CanRegister {
		fmt.Println("The node cannot be registered:")
		if canRegister.AlreadyRegistered {
			fmt.Println("The node is already registered with Stader.")
		}
		if canRegister.RegistrationPaused {
			fmt.Println("Node registrations are currently disabled.")
		}
		return nil
	}

	// Assign max fees
	err = gas.AssignMaxFeeAndLimit(canRegister.GasInfo, staderClient, c.Bool("yes"))
	if err != nil {
		return err
	}

	// Prompt for confirmation
	if !(c.Bool("yes") || cliutils.Confirm("Are you sure you want to register this node?")) {
		fmt.Println("Cancelled.")
		return nil
	}

	// Register node
	response, err := staderClient.RegisterNode(operatorName, common.HexToAddress(operatorRewardAddressString), socializeMevBool)
	if err != nil {
		return err
	}

	fmt.Printf("Registering node...\n")
	cliutils.PrintTransactionHash(staderClient, response.TxHash)
	if _, err = staderClient.WaitForTransaction(response.TxHash); err != nil {
		return err
	}

	// Log & return
	fmt.Println("The node was successfully registered with Stader.")
	return nil

}

func parseToBool(c string) bool {
	if c == "true" {
		return true
	}

	return false
}
