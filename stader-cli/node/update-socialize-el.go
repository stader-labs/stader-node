package node

import (
	"fmt"

	"github.com/stader-labs/stader-node/shared/services/gas"
	"github.com/urfave/cli"

	"github.com/stader-labs/stader-node/shared/services/stader"
	cliutils "github.com/stader-labs/stader-node/shared/utils/cli"
)

func UpdateSocializeEl(c *cli.Context, socializeEl bool) error {

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

	// check if we can update the el
	res, err := staderClient.CanUpdateSocializeEl(socializeEl)
	if err != nil {
		return err
	}
	if res.IsPermissionlessNodeRegistryPaused {
		fmt.Println("Permissionless node registry is paused!")
		return nil
	}
	if res.AlreadyOptedIn {
		fmt.Println("You have already opted in to the socializing pool!")
		return nil
	}
	if res.AlreadyOptedOut {
		fmt.Println("You have already opted out of the socializing pool!")
		return nil
	}
	if res.InCooldown {
		fmt.Printf("You are in cooldown period! You can opt in/out after %d block\n", res.NextUpdatableBlock)
		return nil
	}

	err = gas.AssignMaxFeeAndLimit(res.GasInfo, staderClient, c.Bool("yes"))
	if err != nil {
		return err
	}

	var confirmText string

	if socializeEl {
		confirmText = `
Note: There is a 56-day cool-off period for both Opt-In and Opt-Out of the socializing pool. For example, once you Opt-In, you will have to wait for 56 days to Opt-Out of the Socializing pool, and vice versa.
Are you sure you want to Opt-Into the socializing pool? `
	} else {
		confirmText = `
Note: There is a 56-day cool-off period for both Opt-In and Opt-Out of the socializing pool. For example, once you Opt-Out, you will have to wait for 56 days to Opt Into the Socializing pool, and vice versa.
Are you sure you want to Opt-Out of the socializing pool?`
	}

	if !(c.Bool("yes") || cliutils.Confirm(confirmText)) {
		fmt.Println("Cancelled.")
		return nil
	}

	// update the socializing pool el
	response, err := staderClient.UpdateSocializeEl(socializeEl)
	if err != nil {
		return err
	}

	if socializeEl {
		fmt.Printf("Opting into the socializing pool...\n")
	} else {
		fmt.Printf("Opting out of the socializing pool...\n")
	}
	cliutils.PrintTransactionHash(staderClient, response.TxHash)
	_, err = staderClient.WaitForTransaction(response.TxHash)
	if err != nil {
		return err
	}

	if socializeEl {
		fmt.Printf("Your request to Opt-In to the socializing pool is successful and will take effect after 3 epochs, approximately 20 minutes. \n")
	} else {
		fmt.Printf("Your request to Opt-Out of the socializing pool is successful and will take effect after 3 epochs, approximately 20 minutes. \n")
	}

	return nil

}
