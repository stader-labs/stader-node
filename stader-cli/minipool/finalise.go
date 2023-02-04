package minipool

import (
	"fmt"

	"github.com/ethereum/go-ethereum/common"
	"github.com/urfave/cli"

	"github.com/stader-labs/stader-node/shared/services/gas"
	"github.com/stader-labs/stader-node/shared/services/stader"
	cliutils "github.com/stader-labs/stader-node/shared/utils/cli"
)

func finaliseMinipool(c *cli.Context, minipoolAddress common.Address) error {

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

	// Check if the minipool can be finalised
	canResponse, err := staderClient.CanFinaliseMinipool(minipoolAddress)
	if err != nil {
		return err
	}

	// Assign max fees
	err = gas.AssignMaxFeeAndLimit(canResponse.GasInfo, staderClient, c.Bool("yes"))
	if err != nil {
		return err
	}

	// Prompt for confirmation
	if !(c.Bool("yes") || cliutils.Confirm(fmt.Sprintf("Are you sure you want to finalize minipool %s?", minipoolAddress.Hex()))) {
		fmt.Println("Cancelled.")
		return nil
	}

	// Finalise the minipool
	response, err := staderClient.FinaliseMinipool(minipoolAddress)
	if err != nil {
		return err
	}

	fmt.Printf("Finalizing minipool %s...\n", minipoolAddress)
	cliutils.PrintTransactionHash(staderClient, response.TxHash)
	if _, err = staderClient.WaitForTransaction(response.TxHash); err != nil {
		return err
	}

	// Log & return
	fmt.Printf("Successfully finalized the minipool.\n")
	return nil

}
