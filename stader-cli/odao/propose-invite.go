package odao

import (
	"fmt"

	"github.com/ethereum/go-ethereum/common"
	"github.com/urfave/cli"

	"github.com/stader-labs/stader-node/shared/services/gas"
	"github.com/stader-labs/stader-node/shared/services/stader"
	cliutils "github.com/stader-labs/stader-node/shared/utils/cli"
)

func proposeInvite(c *cli.Context, memberAddress common.Address, memberId, memberUrl string) error {

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
	canPropose, err := staderClient.CanProposeInviteToTNDAO(memberAddress, memberId, memberUrl)
	if err != nil {
		return err
	}
	if !canPropose.CanPropose {
		fmt.Println("Cannot propose inviting member:")
		if canPropose.ProposalCooldownActive {
			fmt.Println("The node must wait for the proposal cooldown period to pass before making another proposal.")
		}
		if canPropose.MemberAlreadyExists {
			fmt.Printf("The node %s is already a member of the oracle DAO.\n", memberAddress.Hex())
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
	response, err := staderClient.ProposeInviteToTNDAO(memberAddress, memberId, memberUrl)
	if err != nil {
		return err
	}

	fmt.Printf("Inviting %s to the oracle DAO...\n", memberAddress.Hex())
	cliutils.PrintTransactionHash(staderClient, response.TxHash)
	if _, err = staderClient.WaitForTransaction(response.TxHash); err != nil {
		return err
	}

	// Log & return
	fmt.Printf("Successfully submitted an invite proposal with ID %d for node %s.\n", response.ProposalId, memberAddress.Hex())
	return nil

}
