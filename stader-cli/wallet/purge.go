package wallet

import (
	"fmt"
	"github.com/stader-labs/stader-node/shared/utils/log"

	"github.com/urfave/cli"

	"github.com/stader-labs/stader-node/shared/services/stader"
	cliutils "github.com/stader-labs/stader-node/shared/utils/cli"
	"github.com/stader-labs/stader-node/stader-cli/service"
)

func purge(c *cli.Context) error {

	// Get RP client
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

	// Restart RP node and watchtower now that the wallet's gone
	if !cfg.IsNativeMode {
		projectName := cfg.Smartnode.ProjectName.Value.(string)
		nodeName := projectName + service.NodeContainerSuffix
		watchtowerName := projectName + service.WatchtowerContainerSuffix

		// Restart node
		err := restartContainer(staderClient, nodeName)
		if err != nil {
			return err
		}

		// Restart watchtower
		err = restartContainer(staderClient, watchtowerName)
		if err != nil {
			return err
		}
	} else {
		fmt.Printf("%sNOTE: As you are in Native mode, please restart your node and watchtower services manually to remove the cached wallet information.%s\n\n", log.ColorYellow, log.ColorReset)
	}

	fmt.Printf("Deleted the node wallet and all validator keys.\n**Please verify that the keys have been removed by looking at your validator logs before continuing.**\n\n")
	fmt.Printf("%sWARNING: If you intend to use these keys for validating again on this or any other machine, you must wait **at least fifteen minutes** after running this command before you can safely begin validating with them again.\nFailure to wait **could cause you to be slashed!**%s\n", log.ColorYellow, log.ColorReset)
	return nil

}

func restartContainer(rp *stader.Client, containerName string) error {
	// Restart node
	result, err := rp.RestartContainer(containerName)
	if err != nil {
		return fmt.Errorf("Error stopping %s container: %w", containerName, err)
	}
	if result != containerName {
		return fmt.Errorf("Unexpected output while stopping %s container: %s", containerName, result)
	}
	return nil
}
