package wallet

import (
	"fmt"
	staderCore "github.com/stader-labs/stader-minipool-go/stader"

	"github.com/stader-labs/stader-node/shared/services/gas"
	"github.com/stader-labs/stader-node/shared/services/stader"
	cliutils "github.com/stader-labs/stader-node/shared/utils/cli"
	"github.com/urfave/cli"
)

func setEnsName(c *cli.Context, name string) error {

	// Get RP client
	staderClient, err := stader.NewClientFromCtx(c)
	if err != nil {
		return err
	}
	defer staderClient.Close()

	fmt.Printf("This will confirm the node's ENS name as '%s'.\n\n%sNOTE: to confirm your name, you must first register it with the ENS application at https://app.ens.domains.\nWe recommend using a hardware wallet as the base domain, and registering your node as a subdomain of it.%s\n\n", name, colorYellow, colorReset)

	// Get gas estimate
	estimateGasSetName, err := staderClient.EstimateGasSetEnsName(name)
	if err != nil {
		return err
	}

	// Assign max fees
	err = gas.AssignMaxFeeAndLimit(staderCore.GasInfo(estimateGasSetName.GasInfo), staderClient, c.Bool("yes"))
	if err != nil {
		return err
	}

	if !cliutils.Confirm("Are you sure you want to confirm your node's ENS name?") {
		fmt.Println("Cancelled.")
		return nil
	}

	// Set the name
	response, err := staderClient.SetEnsName(name)
	if err != nil {
		return err
	}

	fmt.Printf("Setting ENS name...\n")
	cliutils.PrintTransactionHash(staderClient, response.TxHash)
	if _, err = staderClient.WaitForTransaction(response.TxHash); err != nil {
		return err
	}

	fmt.Printf("The ENS name associated with your node account is now '%s'.\n\n", name)
	return nil

}
