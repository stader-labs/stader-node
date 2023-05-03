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
	if res.SocializingPoolContractPaused {
		fmt.Println("The socializing pool contract is paused!")
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
		fmt.Println("You are in cooldown period!")
		return nil
	}

	err = gas.AssignMaxFeeAndLimit(res.GasInfo, staderClient, c.Bool("yes"))
	if err != nil {
		return err
	}

	if !(c.Bool("yes") || cliutils.Confirm(fmt.Sprintf(
		"Are you sure you want to update socializing pool participation?"))) {
		fmt.Println("Cancelled.")
		return nil
	}

	// update the socializing pool el
	response, err := staderClient.UpdateSocializeEl(socializeEl)
	if err != nil {
		return err
	}

	if socializeEl {
		fmt.Printf("Opting in for socializing pool...\n")
	} else {
		fmt.Printf("Opting out for socializing pool...\n")
	}
	cliutils.PrintTransactionHash(staderClient, response.TxHash)
	_, err = staderClient.WaitForTransaction(response.TxHash)
	if err != nil {
		return err
	}

	if socializeEl {
		fmt.Printf("Opted in of socializing pool...\n")
	} else {
		fmt.Printf("Opted out of socializing pool...\n")
	}

	return nil

}
