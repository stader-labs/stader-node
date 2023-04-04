package node

import (
	"fmt"
	"github.com/stader-labs/stader-node/shared/services/gas"
	"github.com/stader-labs/stader-node/shared/utils/log"
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

	syncResponse, err := staderClient.NodeSync()
	if err != nil {
		fmt.Printf("%s**WARNING**: Can't verify the sync status of your consensus client.\nYOU WILL LOSE ETH if your validator is activated before it is fully synced.\n"+
			"Reason: %s\n%s", log.ColorRed, err, log.ColorReset)
	} else {
		if syncResponse.BcStatus.PrimaryClientStatus.IsSynced {
			fmt.Printf("Your consensus client is synced, you may safely create a validator.\n")
		} else if syncResponse.BcStatus.FallbackEnabled {
			if syncResponse.BcStatus.FallbackClientStatus.IsSynced {
				fmt.Printf("Your fallback consensus client is synced, you may safely create a validator.\n")
			} else {
				fmt.Printf("%s**WARNING**: neither your primary nor fallback consensus clients are fully synced.\nYOU WILL LOSE ETH if your validator is activated before they are fully synced.\n%s", log.ColorRed, log.ColorReset)
			}
		} else {
			fmt.Printf("%s**WARNING**: your primary consensus client is either not fully synced or offline and you do not have a fallback client configured.\nYOU WILL LOSE ETH if your validator is activated before it is fully synced.\n%s", log.ColorRed, log.ColorReset)
		}
	}

	// check if we can update the el
	res, err := staderClient.CanUpdateSocializeEl(socializeEl)
	if err != nil {
		return err
	}
	if res.AlreadyOptedIn {
		fmt.Println("You have already opted in to the socializing pool!")
		return nil
	}
	if res.AlreadyOptedOut {
		fmt.Println("You have already opted out of the socializing pool!")
		return nil
	}

	err = gas.AssignMaxFeeAndLimit(res.GasInfo, staderClient, c.Bool("yes"))
	if err != nil {
		return err
	}

	fmt.Printf("socializeEl is %t\n", socializeEl)
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
