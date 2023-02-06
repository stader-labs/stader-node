package odao

import (
	"fmt"

	"github.com/urfave/cli"

	"github.com/stader-labs/stader-node/shared/services/gas"
	"github.com/stader-labs/stader-node/shared/services/stader"
	cliutils "github.com/stader-labs/stader-node/shared/utils/cli"
)

func proposeLeave(c *cli.Context) error {

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

	// Check if proposal can be made
	canPropose, err := staderClient.CanProposeLeaveTNDAO()
	if err != nil {
		return err
	}
	if !canPropose.CanPropose {
		fmt.Println("Cannot propose leaving:")
		if canPropose.ProposalCooldownActive {
			fmt.Println("The node must wait for the proposal cooldown period to pass before making another proposal.")
		}
		if canPropose.InsufficientMembers {
			fmt.Println("There are not enough members in the oracle DAO to allow a member to leave.")
		}
		return nil
	}
	// Assign max fees
	err = gas.AssignMaxFeeAndLimit(canPropose.GasInfo, staderClient, c.Bool("yes"))
	if err != nil {
		return err
	}

	// Prompt for confirmation
	if !(c.Bool("yes") || cliutils.Confirm("Are you sure you want to submit this proposal?")) {
		fmt.Println("Cancelled.")
		return nil
	}

	// Submit proposal
	response, err := staderClient.ProposeLeaveTNDAO()
	if err != nil {
		return err
	}

	fmt.Printf("Proposing a leave from the oracle DAO...\n")
	cliutils.PrintTransactionHash(staderClient, response.TxHash)
	if _, err = staderClient.WaitForTransaction(response.TxHash); err != nil {
		return err
	}

	// Log & return
	fmt.Printf("Successfully submitted a leave proposal with ID %d.\n", response.ProposalId)
	return nil

}
