package node

import (
	"fmt"

	"github.com/urfave/cli"

	"github.com/stader-labs/stader-node/shared/services/gas"
	"github.com/stader-labs/stader-node/shared/services/stader"
	cliutils "github.com/stader-labs/stader-node/shared/utils/cli"
)

func setTimezoneLocation(c *cli.Context) error {

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

	// Get the gas estimate
	canResponse, err := staderClient.CanSetNodeTimezone(timezoneLocation)
	if err != nil {
		return err
	}

	// Assign max fees
	err = gas.AssignMaxFeeAndLimit(canResponse.GasInfo, staderClient, c.Bool("yes"))
	if err != nil {
		return err
	}

	// Prompt for confirmation
	if !(c.Bool("yes") || cliutils.Confirm("Are you sure you want to set your timezone?")) {
		fmt.Println("Cancelled.")
		return nil
	}

	// Set node's timezone location
	response, err := staderClient.SetNodeTimezone(timezoneLocation)
	if err != nil {
		return err
	}

	fmt.Printf("Setting timezone...\n")
	cliutils.PrintTransactionHash(staderClient, response.TxHash)
	if _, err = staderClient.WaitForTransaction(response.TxHash); err != nil {
		return err
	}

	// Log & return
	fmt.Printf("The node's timezone location was successfully updated to '%s'.\n", timezoneLocation)
	return nil

}
