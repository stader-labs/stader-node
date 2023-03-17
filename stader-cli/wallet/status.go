package wallet

import (
	"fmt"

	"github.com/stader-labs/stader-node/shared/services/stader"
	"github.com/urfave/cli"

	cliutils "github.com/stader-labs/stader-node/shared/utils/cli"
)

func getStatus(c *cli.Context) error {

	staderClient, err := stader.NewClientFromCtx(c)
	if err != nil {
		return err
	}
	defer staderClient.Close()

	// Print what network we're on
	err = cliutils.PrintNetwork(staderClient)
	if err != nil {
		return err
	}

	// Get wallet status
	status, err := staderClient.WalletStatus()
	if err != nil {
		return err
	}

	// Print status & return
	if status.WalletInitialized {
		fmt.Println("The node wallet is initialized.")
		fmt.Printf("Node account: %s\n", status.AccountAddress.Hex())
	} else {
		fmt.Println("The node wallet has not been initialized.")
	}
	return nil

}
