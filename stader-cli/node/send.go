/*
This work is licensed and released under GNU GPL v3 or any other later versions.
The full text of the license is below/ found at <http://www.gnu.org/licenses/>

(c) 2023 Rocket Pool Pty Ltd. Modified under GNU GPL v3. [0.3.0-beta]

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
	"github.com/stader-labs/stader-node/stader-lib/utils/eth"
	"github.com/urfave/cli"

	"github.com/stader-labs/stader-node/shared/services/gas"
	"github.com/stader-labs/stader-node/shared/services/stader"
	cliutils "github.com/stader-labs/stader-node/shared/utils/cli"
	"github.com/stader-labs/stader-node/shared/utils/math"
)

func nodeSend(c *cli.Context, amount float64, token string, toAddressOrENS string) error {

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

	// Get amount in wei
	amountWei := eth.EthToWei(amount)

	// Check tokens can be sent
	canSend, err := staderClient.CanNodeSend(amountWei, token)
	if err != nil {
		return err
	}
	if !canSend.CanSend {
		fmt.Println("Cannot send tokens:")
		if canSend.InsufficientBalance {
			fmt.Printf("The node's %s balance is insufficient.\n", token)
		}
		return nil
	}
	var toAddress common.Address
	var toAddressString string

	toAddress, err = cliutils.ValidateAddress("to address", toAddressOrENS)
	if err != nil {
		return err
	}
	toAddressString = toAddress.Hex()

	// Prompt for confirmation
	if !(c.Bool("yes") || cliutils.Confirm(fmt.Sprintf("Are you sure you want to send %.6f %s to %s? This action cannot be undone!", math.RoundDown(eth.WeiToEth(amountWei), 6), token, toAddressString))) {
		fmt.Println("Cancelled.")
		return nil
	}

	// Assign max fees
	err = gas.AssignMaxFeeAndLimit(canSend.GasInfo, staderClient, c.Bool("yes"))
	if err != nil {
		return err
	}

	// Send tokens
	response, err := staderClient.NodeSend(amountWei, token, toAddress)
	if err != nil {
		return err
	}

	fmt.Printf("Sending %s to %s...\n", token, toAddressString)
	cliutils.PrintTransactionHash(staderClient, response.TxHash)
	if _, err = staderClient.WaitForTransaction(response.TxHash); err != nil {
		return err
	}

	// Log & return
	fmt.Printf("Successfully sent %.6f %s to %s.\n", math.RoundDown(eth.WeiToEth(amountWei), 6), token, toAddressString)
	return nil

}
