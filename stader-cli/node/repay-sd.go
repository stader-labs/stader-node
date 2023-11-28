package node

import (
	"fmt"
	"strconv"

	"github.com/urfave/cli"

	"github.com/stader-labs/stader-node/shared/services/gas"
	"github.com/stader-labs/stader-node/shared/services/stader"
	cliutils "github.com/stader-labs/stader-node/shared/utils/cli"
	"github.com/stader-labs/stader-node/stader-lib/utils/eth"
)

func repaySD(c *cli.Context) error {
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
	if err != nil {
		return err
	}

	// Get stake mount
	amountInString := c.String("amount")

	amount, err := strconv.ParseFloat(amountInString, 64)
	if err != nil {
		return err
	}

	amountWei := eth.EthToWei(amount)
	// Check if we can Withdraw El Rewards
	canClaimElRewardsResponse, err := staderClient.CanNodeRepaySd(amountWei)
	if err != nil {
		return err
	}

	// Check if repay more than need
	if amountWei.Cmp(canClaimElRewardsResponse.SdStatusResponse.SdUtilityBalance) >= 1 {
		fmt.Printf("The amount is more than amount need to repay: %s \n", canClaimElRewardsResponse.SdStatusResponse.SdUtilityBalance.String())
		return nil
	}

	if amountWei.Cmp(canClaimElRewardsResponse.SdStatusResponse.SdBalance) >= 1 {
		fmt.Printf("The node's SD balance is not enough: %s \n", canClaimElRewardsResponse.SdStatusResponse.SdBalance.String())
		return nil
	}

	err = gas.AssignMaxFeeAndLimit(canClaimElRewardsResponse.GasInfo, staderClient, c.Bool("yes"))
	if err != nil {
		return err
	}

	// Prompt for confirmation
	if !(c.Bool("yes") || cliutils.Confirm(fmt.Sprintln(
		"Are you sure you want to repay SD?"))) {
		fmt.Println("Cancelled.")
		return nil
	}

	res, err := staderClient.NodeRepaySd(amountWei)
	if err != nil {
		return err
	}

	cliutils.PrintTransactionHash(staderClient, res.TxHash)

	if _, err = staderClient.WaitForTransaction(res.TxHash); err != nil {
		return err
	}

	return nil
}
