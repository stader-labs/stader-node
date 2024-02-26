package node

import (
	"fmt"

	"github.com/stader-labs/stader-node/shared/services/gas"
	"github.com/stader-labs/stader-node/shared/services/stader"
	cliutils "github.com/stader-labs/stader-node/shared/utils/cli"
	"github.com/stader-labs/stader-node/stader-lib/utils/eth"
	"github.com/urfave/cli"
)

func SendElRewards(c *cli.Context) error {
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
	canClaimElRewardsResponse, err := staderClient.CanSendElRewards()
	if err != nil {
		return err
	}
	if canClaimElRewardsResponse.NoElRewards {
		fmt.Printf("No El Rewards to withdraw\n")
		return nil
	}

	err = gas.AssignMaxFeeAndLimit(canClaimElRewardsResponse.GasInfo, staderClient, c.Bool("yes"))
	if err != nil {
		return err
	}

	// Prompt for confirmation
	if !(c.Bool("yes") || cliutils.Confirm(fmt.Sprintf(
		"Are you sure you want to send El Rewards to claim vault?"))) {
		fmt.Println("Cancelled.")
		return nil
	}

	res, err := staderClient.SendElRewards()
	if err != nil {
		return err
	}

	fmt.Printf("Sending %s EL Rewards to Claim Vault\n\n", eth.DisplayAmountInUnits(res.ElRewardsAmount, "eth"))
	cliutils.PrintTransactionHash(staderClient, res.TxHash)

	if _, err = staderClient.WaitForTransaction(res.TxHash); err != nil {
		return err
	}

	// Log & return
	fmt.Printf("Sent %s EL Rewards to Claim Vault\n\n", eth.DisplayAmountInUnits(res.ElRewardsAmount, "eth"))
	return nil
}
