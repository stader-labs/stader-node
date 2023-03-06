package node

import (
	"github.com/urfave/cli"
)

func nodeWithdrawRpl(c *cli.Context) error {
	//
	//// Get RP client
	//staderClient, err := stader.NewClientFromCtx(c)
	//if err != nil {
	//	return err
	//}
	//defer staderClient.Close()
	//
	//// Check and assign the EC status
	//err = cliutils.CheckClientStatus(staderClient)
	//if err != nil {
	//	return err
	//}
	//
	//// Get withdrawal mount
	//var amountWei *big.Int
	//if c.String("amount") == "max" {
	//
	//	// Get node status
	//	status, err := staderClient.NodeStatus()
	//	if err != nil {
	//		return err
	//	}
	//
	//	// Set amount to maximum withdrawable amount
	//	var maxAmount big.Int
	//	if status.RplStake.Cmp(status.MinimumRplStake) > 0 {
	//		maxAmount.Sub(status.RplStake, status.MinimumRplStake)
	//	}
	//	amountWei = &maxAmount
	//
	//} else if c.String("amount") != "" {
	//
	//	// Parse amount
	//	withdrawalAmount, err := strconv.ParseFloat(c.String("amount"), 64)
	//	if err != nil {
	//		return fmt.Errorf("Invalid withdrawal amount '%s': %w", c.String("amount"), err)
	//	}
	//	amountWei = eth.EthToWei(withdrawalAmount)
	//
	//} else {
	//
	//	// Get node status
	//	status, err := staderClient.NodeStatus()
	//	if err != nil {
	//		return err
	//	}
	//
	//	// Get maximum withdrawable amount
	//	var maxAmount big.Int
	//	maxAmount.Sub(status.RplStake, status.MaximumRplStake)
	//	if maxAmount.Sign() == 1 {
	//		// Prompt for maximum amount
	//		if cliutils.Confirm(fmt.Sprintf("Would you like to withdraw the maximum amount of staked RPL (%.6f RPL)?", math.RoundDown(eth.WeiToEth(&maxAmount), 6))) {
	//			amountWei = &maxAmount
	//		} else {
	//
	//			// Prompt for custom amount
	//			inputAmount := cliutils.Prompt("Please enter an amount of staked RPL to withdraw:", "^\\d+(\\.\\d+)?$", "Invalid amount")
	//			withdrawalAmount, err := strconv.ParseFloat(inputAmount, 64)
	//			if err != nil {
	//				return fmt.Errorf("Invalid withdrawal amount '%s': %w", inputAmount, err)
	//			}
	//			amountWei = eth.EthToWei(withdrawalAmount)
	//
	//		}
	//	} else {
	//		fmt.Printf("Cannot withdraw staked RPL - you have %.6f RPL staked, but are not allowed to withdraw below %.6f RPL (150%% collateral).\n",
	//			math.RoundDown(eth.WeiToEth(status.RplStake), 6),
	//			math.RoundDown(eth.WeiToEth(status.MaximumRplStake), 6))
	//		return nil
	//	}
	//
	//}
	//
	//// Check RPL can be withdrawn
	//canWithdraw, err := staderClient.CanNodeWithdrawRpl(amountWei)
	//if err != nil {
	//	return err
	//}
	//if !canWithdraw.CanWithdraw {
	//	fmt.Println("Cannot withdraw staked RPL:")
	//	if canWithdraw.InsufficientBalance {
	//		fmt.Println("The node's staked RPL balance is insufficient.")
	//	}
	//	if canWithdraw.MinipoolsUndercollateralized {
	//		fmt.Println("Remaining staked RPL is not enough to collateralize the node's minipools.")
	//	}
	//	if canWithdraw.WithdrawalDelayActive {
	//		fmt.Println("The withdrawal delay period has not passed.")
	//	}
	//	if !canWithdraw.InConsensus {
	//		fmt.Println("The RPL price and total effective staked RPL of the network are still being voted on by the Oracle DAO.\nPlease try again in a few minutes.")
	//	}
	//	return nil
	//}
	//
	//// Assign max fees
	//err = gas.AssignMaxFeeAndLimit(canWithdraw.GasInfo, staderClient, c.Bool("yes"))
	//if err != nil {
	//	return err
	//}
	//
	//// Prompt for confirmation
	//if !(c.Bool("yes") || cliutils.Confirm(fmt.Sprintf("Are you sure you want to withdraw %.6f staked RPL? This may decrease your node's RPL rewards.", math.RoundDown(eth.WeiToEth(amountWei), 6)))) {
	//	fmt.Println("Cancelled.")
	//	return nil
	//}
	//
	//// Withdraw RPL
	//response, err := staderClient.NodeWithdrawRpl(amountWei)
	//if err != nil {
	//	return err
	//}
	//
	//fmt.Printf("Withdrawing RPL...\n")
	//cliutils.PrintTransactionHash(staderClient, response.TxHash)
	//if _, err = staderClient.WaitForTransaction(response.TxHash); err != nil {
	//	return err
	//}
	//
	//// Log & return
	//fmt.Printf("Successfully withdrew %.6f staked RPL.\n", math.RoundDown(eth.WeiToEth(amountWei), 6))
	//return nil
	return nil
}
