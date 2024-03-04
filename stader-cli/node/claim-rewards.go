package node

import (
	"fmt"
	"math/big"

	"github.com/stader-labs/stader-node/shared/services/gas"
	"github.com/stader-labs/stader-node/shared/services/stader"
	cliutils "github.com/stader-labs/stader-node/shared/utils/cli"
	"github.com/stader-labs/stader-node/stader-lib/utils/eth"
	"github.com/urfave/cli"
)

func ClaimRewards(c *cli.Context) error {
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

	// Print what network we're on
	err = cliutils.PrintNetwork(staderClient)

	// Check if we can Withdraw El Rewards
	canClaimRewardsResponse, err := staderClient.CanClaimRewards()
	if err != nil {
		return err
	}
	if canClaimRewardsResponse.NoRewards {
		fmt.Println("No rewards to claim.")
		return nil
	}

	sdStatusResponse, err := staderClient.GetSDStatus(big.NewInt(0))
	if err != nil {
		return err
	}

	sdStatus := sdStatusResponse.SDStatus

	// totalFee := sdStatus.AccumulatedInterest
	// if withdrawableInEth < claimsBalance, then there is an existing utilization position
	if canClaimRewardsResponse.ClaimsBalance.Cmp(canClaimRewardsResponse.WithdrawableInEth) != 0 {
		if sdStatus.SdUtilizerLatestBalance.Cmp(big.NewInt(0)) > 0 {
			fmt.Printf("You currently have an existing SD Utilization Position of %s. Based on the current Health Factor, you can claim upto %s", eth.DisplayAmountInUnits(sdStatusResponse.SDStatus.SdUtilizerLatestBalance, "sd"), eth.DisplayAmountInUnits(canClaimRewardsResponse.WithdrawableInEth, "eth"))

			fmt.Printf("Note: Please repay your utilized SD by using the following command to claim the remaining ETH: stader-cli sd repay --amount <amount of SD to be repaid>.\n")

			if !cliutils.Confirm("Do you wish to proceed?\n\n") {
				fmt.Println("Cancelled.")
				return nil
			}
		}
	}

	err = gas.AssignMaxFeeAndLimit(canClaimRewardsResponse.GasInfo, staderClient, c.Bool("yes"))
	if err != nil {
		return err
	}

	// Prompt for confirmation
	if !(c.Bool("yes") || cliutils.Confirm(fmt.Sprintf(
		"Are you sure you want to send rewards to your operator reward address?"))) {
		fmt.Println("Cancelled.")
		return nil
	}

	// Withdraw El Rewards
	res, err := staderClient.ClaimRewards()
	if err != nil {
		return err
	}

	fmt.Printf("Withdrawing %s Rewards to Operator Reward Address: %s\n\n", eth.DisplayAmountInUnits(res.RewardsClaimed, "eth"), res.OperatorRewardAddress)
	cliutils.PrintTransactionHash(staderClient, res.TxHash)

	if _, err = staderClient.WaitForTransaction(res.TxHash); err != nil {
		return err
	}

	// Log & return
	fmt.Printf("Successful withdrawal of %s to Operator Reward Address: %s\n\n", eth.DisplayAmountInUnits(res.RewardsClaimed, "eth"), res.OperatorRewardAddress)
	return nil
}
