package node

// ROCKETPOOL-OWNED

import (
	"fmt"
	"github.com/urfave/cli"

	"github.com/stader-labs/stader-node/shared/services/stader"
	cliutils "github.com/stader-labs/stader-node/shared/utils/cli"
)

func getSyncProgress(c *cli.Context) error {

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

	// Get node status
	status, err := staderClient.NodeSync()
	if err != nil {
		return err
	}

	// Print EC status
	if status.EcStatus.PrimaryClientStatus.Error != "" {
		fmt.Printf("Your primary execution client is unavailable (%s).\n", status.EcStatus.PrimaryClientStatus.Error)
	} else if status.EcStatus.PrimaryClientStatus.IsSynced {
		fmt.Print("Your primary execution client is fully synced.\n")
	} else {
		fmt.Printf("Your primary execution client is still syncing (%0.2f%%).\n", status.EcStatus.PrimaryClientStatus.SyncProgress*100)
		if status.EcStatus.PrimaryClientStatus.SyncProgress == 0 {
			fmt.Println("\tNOTE: your execution client may not report sync progress.\n\tYou should check its logs to review it.")
		}
	}

	// Print fallback EC status
	if status.EcStatus.FallbackEnabled {
		if status.EcStatus.FallbackClientStatus.Error != "" {
			fmt.Printf("Your fallback execution client is unavailable (%s).\n", status.EcStatus.FallbackClientStatus.Error)
		} else if status.EcStatus.FallbackClientStatus.IsSynced {
			fmt.Print("Your fallback execution client is fully synced.\n")
		} else {
			fmt.Printf("Your fallback execution client is still syncing (%0.2f%%).\n", status.EcStatus.FallbackClientStatus.SyncProgress*100)
			if status.EcStatus.FallbackClientStatus.SyncProgress == 0 {
				fmt.Println("\tNOTE: your execution client may not report sync progress.\n\tYou should check your its logs to review it.")
			}
		}
	} else {
		fmt.Printf("You do not have a fallback execution client enabled.\n")
	}

	// Print CC status
	if status.BcStatus.PrimaryClientStatus.Error != "" {
		fmt.Printf("Your primary consensus client is unavailable (%s).\n", status.BcStatus.PrimaryClientStatus.Error)
	} else if status.BcStatus.PrimaryClientStatus.IsSynced {
		fmt.Print("Your primary consensus client is fully synced.\n")
	} else {
		fmt.Printf("Your primary consensus client is still syncing (%0.2f%%).\n", status.BcStatus.PrimaryClientStatus.SyncProgress*100)
	}

	// Print fallback CC status
	if status.BcStatus.FallbackEnabled {
		if status.BcStatus.FallbackClientStatus.Error != "" {
			fmt.Printf("Your fallback consensus client is unavailable (%s).\n", status.BcStatus.FallbackClientStatus.Error)
		} else if status.BcStatus.FallbackClientStatus.IsSynced {
			fmt.Print("Your fallback consensus client is fully synced.\n")
		} else {
			fmt.Printf("Your fallback consensus client is still syncing (%0.2f%%).\n", status.BcStatus.FallbackClientStatus.SyncProgress*100)
		}
	} else {
		fmt.Printf("You do not have a fallback consensus client enabled.\n")
	}

	// Return
	return nil

}
