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
	"github.com/stader-labs/stader-node/stader-lib/utils/eth"
	"github.com/urfave/cli"
	"math/big"
	"strconv"

	"github.com/stader-labs/stader-node/shared/services/gas"
	"github.com/stader-labs/stader-node/shared/services/stader"
	cliutils "github.com/stader-labs/stader-node/shared/utils/cli"
	"github.com/stader-labs/stader-node/shared/utils/math"
)

// TODO - bchain make user experience such that we suggest users how much sd to stake
func nodeDepositSd(c *cli.Context) error {

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

	// If a custom nonce is set, print the multi-transaction warning
	if c.GlobalUint64("nonce") != 0 {
		cliutils.PrintMultiTransactionNonceWarning()
	}

	// Get stake mount
	amountInString := c.String("amount")
	amount, err := strconv.ParseFloat(amountInString, 64)
	if err != nil {
		return err
	}
	fmt.Printf("amount is %f\n", amount)
	amountWei := eth.EthToWei(amount)
	fmt.Printf("amountWei is %d\n", amountWei)

	// Check allowance
	allowance, err := staderClient.GetNodeDepositSdAllowance()
	if err != nil {
		return err
	}
	fmt.Printf("allowance is %d\n", allowance.Allowance.Int64())

	// Calculate max uint256 value
	maxApproval := big.NewInt(2)
	maxApproval = maxApproval.Exp(maxApproval, big.NewInt(256), nil)
	maxApproval = maxApproval.Sub(maxApproval, big.NewInt(1))

	if allowance.Allowance.Cmp(maxApproval) < 0 {
		fmt.Println("Before depositing SD, you must first give the collateral contract approval to interact with your SD.")
		fmt.Println("This only needs to be done once for your node.")

		// If a custom nonce is set, print the multi-transaction warning
		if c.GlobalUint64("nonce") != 0 {
			cliutils.PrintMultiTransactionNonceWarning()
		}

		// Get approval gas
		approvalGas, err := staderClient.NodeDepositSdApprovalGas(maxApproval)
		if err != nil {
			return err
		}
		// Assign max fees
		err = gas.AssignMaxFeeAndLimit(approvalGas.GasInfo, staderClient, c.Bool("yes"))
		if err != nil {
			return err
		}

		// Prompt for confirmation
		if !(c.Bool("yes") || cliutils.Confirm("Do you want to deposit your SD into the Collateral contract")) {
			fmt.Println("Cancelled.")
			return nil
		}

		response, err := staderClient.NodeDepositSdApprove(maxApproval)
		if err != nil {
			return err
		}
		hash := response.ApproveTxHash
		fmt.Printf("Approving SD for deposting...\n")
		cliutils.PrintTransactionHash(staderClient, hash)
		if _, err = staderClient.WaitForTransaction(hash); err != nil {
			return err
		}
		fmt.Println("Successfully approved SD to deposit.")

		// If a custom nonce is set, increment it for the next transaction
		if c.GlobalUint64("nonce") != 0 {
			staderClient.IncrementCustomNonce()
		}
	}

	canDeposit, err := staderClient.CanNodeDepositSd(amountWei)
	if err != nil {
		return err
	}
	if !canDeposit.CanDeposit {
		fmt.Println("Cannot deposit SD:")
		if canDeposit.InsufficientBalance {
			fmt.Println("The node's SD balance is insufficient.")
		}
		return nil
	}

	// Assign max fees
	err = gas.AssignMaxFeeAndLimit(canDeposit.GasInfo, staderClient, c.Bool("yes"))
	if err != nil {
		return err
	}

	// Prompt for confirmation
	if !(c.Bool("yes") || cliutils.Confirm(fmt.Sprintf("Are you sure you want to deposit %.6f SD? You will not be able to withdraw this SD until you exit your validators", math.RoundDown(eth.WeiToEth(amountWei), 6)))) {
		fmt.Println("Cancelled.")
		return nil
	}

	depositSdResponse, err := staderClient.NodeDepositSd(amountWei)
	if err != nil {
		return err
	}

	fmt.Printf("Deposting SD...\n")
	cliutils.PrintTransactionHash(staderClient, depositSdResponse.DepositTxHash)
	if _, err = staderClient.WaitForTransaction(depositSdResponse.DepositTxHash); err != nil {
		return err
	}

	// Log & return
	fmt.Printf("Successfully deposited %.6f SD.\n", math.RoundDown(eth.WeiToEth(amountWei), 6))
	return nil

}
