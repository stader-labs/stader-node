package node

import (
	"fmt"

	"github.com/urfave/cli"

	"github.com/stader-labs/stader-node/shared/services/gas"
	"github.com/stader-labs/stader-node/shared/services/stader"
	cliutils "github.com/stader-labs/stader-node/shared/utils/cli"
)

func registerNode(c *cli.Context) error {

	// Get RP client
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

	// Prompt for timezone location
	var timezoneLocation string
	if c.String("timezone") != "" {
		timezoneLocation = c.String("timezone")
	} else {
		timezoneLocation = promptTimezone()
	}

	// Check node can be registered
	canRegister, err := staderClient.CanRegisterNode(timezoneLocation)
	if err != nil {
		return err
	}
	if !canRegister.CanRegister {
		fmt.Println("The node cannot be registered:")
		if canRegister.AlreadyRegistered {
			fmt.Println("The node is already registered with Rocket Pool.")
		}
		if canRegister.RegistrationDisabled {
			fmt.Println("Node registrations are currently disabled.")
		}
		return nil
	}

	// Assign max fees
	err = gas.AssignMaxFeeAndLimit(canRegister.GasInfo, staderClient, c.Bool("yes"))
	if err != nil {
		return err
	}

	// Prompt for confirmation
	if !(c.Bool("yes") || cliutils.Confirm("Are you sure you want to register this node?")) {
		fmt.Println("Cancelled.")
		return nil
	}

	// Register node
	response, err := staderClient.RegisterNode(timezoneLocation)
	if err != nil {
		return err
	}

	fmt.Printf("Registering node...\n")
	cliutils.PrintTransactionHash(staderClient, response.TxHash)
	if _, err = staderClient.WaitForTransaction(response.TxHash); err != nil {
		return err
	}

	// Log & return
	fmt.Println("The node was successfully registered with Rocket Pool.")
	return nil

}
