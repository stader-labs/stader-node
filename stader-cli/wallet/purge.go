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
package wallet

// ROCKETPOOL-OWNED

import (
	"fmt"

	"github.com/stader-labs/stader-node/shared/utils/log"

	"github.com/urfave/cli"

	"github.com/stader-labs/stader-node/shared/services/stader"
	cliutils "github.com/stader-labs/stader-node/shared/utils/cli"
	"github.com/stader-labs/stader-node/stader-cli/service"
)

func purge(c *cli.Context) error {

	staderClient, err := stader.NewClientFromCtx(c)
	if err != nil {
		return err
	}
	defer staderClient.Close()

	// Get the config
	cfg, _, err := staderClient.LoadConfig()
	if err != nil {
		return fmt.Errorf("error loading user settings: %w", err)
	}

	if !cliutils.Confirm(fmt.Sprintf("%sWARNING: This will delete your node wallet, all of your validator keys (including externally-generated ones in the 'custom-keys' folder), and restart your Validator Client.\nYou will NO LONGER be able to attest with this machine anymore until you recover your wallet or initialize a new one.\n\nYou MUST have your node wallet's mnemonic recorded before running this, or you will lose access to your node wallet and your validators forever!\n\n%sDo you want to continue?", log.ColorRed, log.ColorReset)) {
		fmt.Println("Cancelled.")
		return nil
	}

	// Purge
	_, err = staderClient.Purge()
	if err != nil {
		return err
	}

	if !cfg.IsNativeMode {
		projectName := cfg.StaderNode.ProjectName.Value.(string)
		nodeName := projectName + service.NodeContainerSuffix
		guardianName := projectName + service.GuardianContainerSuffix

		// Restart node
		err := restartContainer(staderClient, nodeName)
		if err != nil {
			return err
		}

		// Restart guardian
		err = restartContainer(staderClient, guardianName)
		if err != nil {
			return err
		}
	} else {
		fmt.Printf("%sNOTE: As you are in Native mode, please restart your node and guardian services manually to remove the cached wallet information.%s\n\n", log.ColorYellow, log.ColorReset)
	}

	fmt.Printf("Deleted the node wallet and all validator keys.\n**Please verify that the keys have been removed by looking at your validator logs before continuing.**\n\n")
	fmt.Printf("%sWARNING: If you intend to use these keys for validating again on this or any other machine, you must wait **at least fifteen minutes** after running this command before you can safely begin validating with them again.\nFailure to wait **could cause you to be slashed!**%s\n", log.ColorYellow, log.ColorReset)
	return nil

}

func restartContainer(sd *stader.Client, containerName string) error {
	// Restart node
	result, err := sd.RestartContainer(containerName)
	if err != nil {
		return fmt.Errorf("Error stopping %s container: %w", containerName, err)
	}
	if result != containerName {
		return fmt.Errorf("Unexpected output while stopping %s container: %s", containerName, result)
	}
	return nil
}
