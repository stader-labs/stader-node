package node

import (
	"fmt"
	"math/big"

	"github.com/urfave/cli"

	"github.com/stader-labs/stader-node/shared/services/gas"
	"github.com/stader-labs/stader-node/shared/services/stader"
	cliutils "github.com/stader-labs/stader-node/shared/utils/cli"
)

func nodeApproveSd(c *cli.Context) error {

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

	err = nodeApproveSdWithClient(c, staderClient)
	if err != nil {
		return err
	}

	return nil
}

func nodeApproveSdWithClient(c *cli.Context, staderClient *stader.Client) error {
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

		fmt.Println("Approving SD for depositing..")
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

	return nil
}
