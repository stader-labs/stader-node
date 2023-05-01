package node

import (
	"fmt"
	"github.com/stader-labs/stader-node/shared/services/stader"
	cliutils "github.com/stader-labs/stader-node/shared/utils/cli"
	"github.com/stader-labs/stader-node/stader-lib/types"
	"github.com/urfave/cli"
)

func ExitValidator(c *cli.Context, validatorPubKey types.ValidatorPubkey) error {
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

	// check canExit
	response, err := staderClient.CanExitValidator(validatorPubKey)
	if err != nil {
		return err
	}
	if response.ValidatorNotRegistered {
		fmt.Println("Validator not registered!")
		return nil
	}
	if response.ValidatorTooYoung {
		fmt.Println("Validator too young!")
		return nil
	}
	if response.ValidatorExiting {
		fmt.Println("Validator already exiting!")
		return nil
	}

	// Prompt for confirmation
	if !(c.Bool("yes") || cliutils.Confirm(fmt.Sprintf(
		"Are you sure you want to exit validator %s?", validatorPubKey))) {
		fmt.Println("Cancelled.")
		return nil
	}

	// now exit
	_, err = staderClient.ExitValidator(validatorPubKey)
	if err != nil {
		return err
	}

	fmt.Printf("Exiting validator %s, you check check the validator status at %s\n", validatorPubKey, fmt.Sprintf("https://prater.beaconcha.in/validator/%s#withdrawals", validatorPubKey))

	return nil
}
