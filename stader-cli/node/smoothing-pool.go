package node

import (
	"fmt"

	"github.com/urfave/cli"

	"github.com/stader-labs/stader-node/shared/services/gas"
	"github.com/stader-labs/stader-node/shared/services/stader"
	cliutils "github.com/stader-labs/stader-node/shared/utils/cli"
)

func joinSmoothingPool(c *cli.Context) error {

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

	// Get the node's registration status
	status, err := staderClient.NodeGetSmoothingPoolRegistrationStatus()
	if err != nil {
		return err
	}

	if status.NodeRegistered {
		fmt.Println("The node is already joined to the Smoothing Pool.")
		return nil
	}

	if status.TimeLeftUntilChangeable > 0 {
		fmt.Printf("You have recently left the Smoothing Pool. You must wait %s until you can join it again.\n", status.TimeLeftUntilChangeable)
		return nil
	}

	// Print some info
	fmt.Println("You are about to opt into the Smoothing Pool.\nYour fee recipient will be changed to the Smoothing Pool contract.\nAll priority fees and MEV you earn via proposals will be shared equally with other members of the Smoothing Pool.\n\nIf you desire, you can opt back out after one full rewards interval has passed.\n")

	// Get the gas estimate
	canResponse, err := staderClient.CanNodeSetSmoothingPoolStatus(true)
	if err != nil {
		return err
	}

	// Assign max fees
	err = gas.AssignMaxFeeAndLimit(canResponse.GasInfo, staderClient, c.Bool("yes"))
	if err != nil {
		return err
	}

	fmt.Printf("%sNOTE: This process will restart your node's validator client.\nYou may miss an attestation if you are currently scheduled to produce one.%s\n\n", colorYellow, colorReset)

	// Prompt for confirmation
	if !(c.Bool("yes") || cliutils.Confirm("Are you sure you want to join the Smoothing Pool?")) {
		fmt.Println("Cancelled.")
		return nil
	}

	// Set the fee recipient to the Smoothing Pool
	response, err := staderClient.NodeSetSmoothingPoolStatus(true)
	if err != nil {
		return err
	}

	fmt.Printf("Joining the Smoothing Pool...\n")
	cliutils.PrintTransactionHash(staderClient, response.TxHash)
	if _, err = staderClient.WaitForTransaction(response.TxHash); err != nil {
		return fmt.Errorf("%w\nYour fee recipient will be automatically reset to your node's distributor in a few minutes, and your validator client will restart.", err)
	}

	// Log & return
	fmt.Println("Successfully joined the Smoothing Pool.")
	return nil

}

func leaveSmoothingPool(c *cli.Context) error {

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

	// Get the node's registration status
	status, err := staderClient.NodeGetSmoothingPoolRegistrationStatus()
	if err != nil {
		return err
	}

	if !status.NodeRegistered {
		fmt.Println("The node is not currently joined to the Smoothing Pool.")
		return nil
	}

	if status.TimeLeftUntilChangeable > 0 {
		fmt.Printf("You have recently joined the Smoothing Pool. You must wait %s until you can leave it.\n", status.TimeLeftUntilChangeable)
		return nil
	}

	// Print some info
	fmt.Println("You are about to opt out of the Smoothing Pool.\nYour fee recipient will be changed back to your node's distributor contract once the next Epoch has been finalized.\nAll priority fees and MEV you earn via proposals will go directly to your distributor and will not be shared by the Smoothing Pool members.\n\nIf you desire, you can opt back in after one full rewards interval has passed.\n")

	// Get the gas estimate
	canResponse, err := staderClient.CanNodeSetSmoothingPoolStatus(false)
	if err != nil {
		return err
	}

	// Assign max fees
	err = gas.AssignMaxFeeAndLimit(canResponse.GasInfo, staderClient, c.Bool("yes"))
	if err != nil {
		return err
	}

	// Prompt for confirmation
	if !(c.Bool("yes") || cliutils.Confirm("Are you sure you want to leave the Smoothing Pool?")) {
		fmt.Println("Cancelled.")
		return nil
	}

	// Set the fee recipient to the Fee Distributor
	response, err := staderClient.NodeSetSmoothingPoolStatus(false)
	if err != nil {
		return err
	}

	fmt.Printf("Leaving the Smoothing Pool...\n")
	cliutils.PrintTransactionHash(staderClient, response.TxHash)
	if _, err = staderClient.WaitForTransaction(response.TxHash); err != nil {
		return err
	}

	// Log & return
	fmt.Println("Successfully left the Smoothing Pool.")
	fmt.Printf("%sNOTE: Your validator client will restart to change its fee recipient back to your node's distributor once the next Epoch has been finalized.\nYou may miss an attestation when this happens (or multiple if you have Doppelganger Protection enabled); this is normal.%s\n", colorYellow, colorReset)
	return nil

}
