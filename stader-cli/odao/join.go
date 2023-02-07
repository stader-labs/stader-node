package odao

import (
	"fmt"
	"math/big"

	"github.com/rocket-pool/rocketpool-go/utils/eth"
	"github.com/urfave/cli"

	"github.com/stader-labs/stader-node/shared/services/gas"
	"github.com/stader-labs/stader-node/shared/services/stader"
	cliutils "github.com/stader-labs/stader-node/shared/utils/cli"
	"github.com/stader-labs/stader-node/shared/utils/math"
)

func join(c *cli.Context) error {

	// Get RP client
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

	// Get node status
	status, err := staderClient.NodeStatus()
	if err != nil {
		return err
	}

	// If a custom nonce is set, print the multi-transaction warning
	if c.GlobalUint64("nonce") != 0 {
		cliutils.PrintMultiTransactionNonceWarning()
	}

	// Check for fixed-supply RPL balance
	if status.AccountBalances.FixedSupplyRPL.Cmp(big.NewInt(0)) > 0 {

		// Confirm swapping RPL
		if c.Bool("swap") || cliutils.Confirm(fmt.Sprintf("The node has a balance of %.6f old RPL. Would you like to swap it for new RPL before transferring your bond?", math.RoundDown(eth.WeiToEth(status.AccountBalances.FixedSupplyRPL), 6))) {

			// Check allowance
			allowance, err := staderClient.GetNodeSwapRplAllowance()
			if err != nil {
				return err
			}

			if allowance.Allowance.Cmp(status.AccountBalances.FixedSupplyRPL) < 0 {
				fmt.Println("Before swapping legacy RPL for new RPL, you must first give the new RPL contract approval to interact with your legacy RPL.")
				fmt.Println("This only needs to be done once for your node.")

				// If a custom nonce is set, print the multi-transaction warning
				if c.GlobalUint64("nonce") != 0 {
					cliutils.PrintMultiTransactionNonceWarning()
				}

				// Calculate max uint256 value
				maxApproval := big.NewInt(2)
				maxApproval = maxApproval.Exp(maxApproval, big.NewInt(256), nil)
				maxApproval = maxApproval.Sub(maxApproval, big.NewInt(1))

				// Get approval gas
				approvalGas, err := staderClient.NodeSwapRplApprovalGas(maxApproval)
				if err != nil {
					return err
				}
				// Assign max fees
				err = gas.AssignMaxFeeAndLimit(approvalGas.GasInfo, staderClient, c.Bool("yes"))
				if err != nil {
					return err
				}

				// Prompt for confirmation
				if !(c.Bool("yes") || cliutils.Confirm(fmt.Sprintf("Do you want to let the new RPL contract interact with your legacy RPL?"))) {
					fmt.Println("Cancelled.")
					return nil
				}

				// Approve RPL for swapping
				response, err := staderClient.NodeSwapRplApprove(maxApproval)
				if err != nil {
					return err
				}
				hash := response.ApproveTxHash
				fmt.Printf("Approving legacy RPL for swapping...\n")
				cliutils.PrintTransactionHash(staderClient, hash)
				if _, err = staderClient.WaitForTransaction(hash); err != nil {
					return err
				}
				fmt.Println("Successfully approved access to legacy RPL.")

				// If a custom nonce is set, increment it for the next transaction
				if c.GlobalUint64("nonce") != 0 {
					staderClient.IncrementCustomNonce()
				}
			}

			// Check RPL can be swapped
			canSwap, err := staderClient.CanNodeSwapRpl(status.AccountBalances.FixedSupplyRPL)
			if err != nil {
				return err
			}
			if !canSwap.CanSwap {
				fmt.Println("Cannot swap RPL:")
				if canSwap.InsufficientBalance {
					fmt.Println("The node's old RPL balance is insufficient.")
				}
				return nil
			}
			fmt.Println("RPL Swap Gas Info:")
			// Assign max fees
			err = gas.AssignMaxFeeAndLimit(canSwap.GasInfo, staderClient, c.Bool("yes"))
			if err != nil {
				return err
			}

			// Prompt for confirmation
			if !(c.Bool("yes") || cliutils.Confirm(fmt.Sprintf("Are you sure you want to swap %.6f old RPL for new RPL?", math.RoundDown(eth.WeiToEth(status.AccountBalances.FixedSupplyRPL), 6)))) {
				fmt.Println("Cancelled.")
				return nil
			}

			// Swap RPL
			swapResponse, err := staderClient.NodeSwapRpl(status.AccountBalances.FixedSupplyRPL)
			if err != nil {
				return err
			}

			fmt.Printf("Swapping old RPL for new RPL...\n")
			cliutils.PrintTransactionHash(staderClient, swapResponse.SwapTxHash)
			if _, err = staderClient.WaitForTransaction(swapResponse.SwapTxHash); err != nil {
				return err
			}

			// Log
			fmt.Printf("Successfully swapped %.6f old RPL for new RPL.\n", math.RoundDown(eth.WeiToEth(status.AccountBalances.FixedSupplyRPL), 6))
			fmt.Println("")

			// If a custom nonce is set, increment it for the next transaction
			if c.GlobalUint64("nonce") != 0 {
				staderClient.IncrementCustomNonce()
			}
		}
	}

	// Check if node can join the oracle DAO
	canJoin, err := staderClient.CanJoinTNDAO()
	if err != nil {
		return err
	}
	if !canJoin.CanJoin {
		fmt.Println("Cannot join the oracle DAO:")
		if canJoin.ProposalExpired {
			fmt.Println("The proposal for you to join the oracle DAO does not exist or has expired.")
		}
		if canJoin.AlreadyMember {
			fmt.Println("The node is already a member of the oracle DAO.")
		}
		if canJoin.InsufficientRplBalance {
			fmt.Println("The node does not have enough RPL to pay the RPL bond.")
		}
		return nil
	}

	// Display gas estimate
	// Assign max fees
	err = gas.AssignMaxFeeAndLimit(canJoin.GasInfo, staderClient, c.Bool("yes"))
	if err != nil {
		return err
	}
	staderClient.PrintMultiTxWarning()

	// Prompt for confirmation
	if !(c.Bool("yes") || cliutils.Confirm("Are you sure you want to join the oracle DAO? Your RPL bond will be locked until you leave.")) {
		fmt.Println("Cancelled.")
		return nil
	}

	// Approve RPL for joining the ODAO
	response, err := staderClient.ApproveRPLToJoinTNDAO()
	if err != nil {
		return err
	}
	hash := response.ApproveTxHash
	fmt.Printf("Approving RPL for joining the Oracle DAO...\n")
	cliutils.PrintTransactionHashNoCancel(staderClient, hash)

	// If a custom nonce is set, increment it for the next transaction
	if c.GlobalUint64("nonce") != 0 {
		staderClient.IncrementCustomNonce()
	}

	// Join the ODAO
	joinResponse, err := staderClient.JoinTNDAO(hash)
	if err != nil {
		return err
	}
	fmt.Printf("Joining the ODAO...\n")
	cliutils.PrintTransactionHash(staderClient, joinResponse.JoinTxHash)
	if _, err = staderClient.WaitForTransaction(joinResponse.JoinTxHash); err != nil {
		return err
	}

	// Log & return
	fmt.Println("Successfully joined the oracle DAO.")
	return nil

}
