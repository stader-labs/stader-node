package node

import (
	"fmt"
	"github.com/stader-labs/stader-node/shared/services/stader"
	cliutils "github.com/stader-labs/stader-node/shared/utils/cli"
	"github.com/urfave/cli"
)

func claimSd(c *cli.Context) error {
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

	canClaimSdResponse, err := staderClient.CanClaimSd()
	if err != nil {
		return err
	}
	if canClaimSdResponse.NoExistingClaim {
		fmt.Println("No existing claim!")
		return nil
	}
	if canClaimSdResponse.ClaimIsInUnbondingPeriod {
		fmt.Println("Claim is in unbonding period!")
		return nil
	}

	// Prompt for confirmation
	if !(c.Bool("yes") || cliutils.Confirm(fmt.Sprintf(
		"Are you sure you want to claim SD? (y/n) "))) {
		fmt.Println("Cancelled.")
		return nil
	}

	res, err := staderClient.ClaimSd()
	if err != nil {
		return err
	}

	fmt.Printf("Claiming SD... \n")
	cliutils.PrintTransactionHash(staderClient, res.TxHash)
	if _, err = staderClient.WaitForTransaction(res.TxHash); err != nil {
		return err
	}

	// Log & return
	fmt.Printf("Claimed SD successfully!\n")

	return nil
}
