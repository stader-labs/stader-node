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
	amountWei := eth.EthToWei(amount)

	// Check allowance
	allowance, err := staderClient.GetNodeDepositSdAllowance()
	if err != nil {
		return err
	}

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
		if !(c.Bool("yes") || cliutils.Confirm("Do you want to approve SD to be spent by the Collateral Contract?")) {
			fmt.Println("Cancelled.")
			return nil
		}

		response, err := staderClient.NodeDepositSdApprove(maxApproval)
		if err != nil {
			return err
		}
		hash := response.ApproveTxHash
		fmt.Printf("Approving SD for depositing...\n")
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
	if canDeposit.InsufficientBalance {
		fmt.Println("The node's SD balance is insufficient.")
		return nil
	}
	if canDeposit.CollateralContractPaused {
		fmt.Println("The collateral contract is paused.")
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

	fmt.Printf("Depositing SD...\n")
	cliutils.PrintTransactionHash(staderClient, depositSdResponse.DepositTxHash)
	if _, err = staderClient.WaitForTransaction(depositSdResponse.DepositTxHash); err != nil {
		return err
	}

	// Log & return
	fmt.Printf("Successfully deposited %.6f SD.\n", math.RoundDown(eth.WeiToEth(amountWei), 6))
	return nil

}
