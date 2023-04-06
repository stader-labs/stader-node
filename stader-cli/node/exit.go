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
		fmt.Printf("Validator not registered!")
		return nil
	}

	// now exit
	_, err = staderClient.ExitValidator(validatorPubKey)
	if err != nil {
		return err
	}

	fmt.Printf("Exiting validator %s, you check check the validator status on the explorer", validatorPubKey)

	return nil
}
