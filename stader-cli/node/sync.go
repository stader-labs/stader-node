/*
This work is licensed and released under GNU GPL v3 or any other later versions.
The full text of the license is below/ found at <http://www.gnu.org/licenses/>

(c) 2023 Rocket Pool Pty Ltd. Modified under GNU GPL v3.

This program is free software: you can redistribute it and/or modify
it under the terms of the GNU General Public License as published by
the Free Software Foundation, either version 3 of the License, or
(at your option) any later version.

This program is distributed in the hope that it will be useful,
but WITHOUT ANY WARRANTY; without even the implied warranty of
MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
GNU General Public License for more details.

You should have received a copy of the GNU General Public License
along with this program.  If not, see <http://www.gnu.org/licenses/>.
*/
package node

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
