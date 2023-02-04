package odao

import (
	"fmt"

	"github.com/ethereum/go-ethereum/common"
	"github.com/urfave/cli"

	"github.com/stader-labs/stader-node/shared/services/gas"
	"github.com/stader-labs/stader-node/shared/services/stader"
	cliutils "github.com/stader-labs/stader-node/shared/utils/cli"
)

func leave(c *cli.Context) error {

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

	// Get the RPL bond refund address
	var bondRefundAddress common.Address
	if c.String("refund-address") == "node" {

		// Set bond refund address to node address
		wallet, err := staderClient.WalletStatus()
		if err != nil {
			return err
		}
		bondRefundAddress = wallet.AccountAddress

	} else if c.String("refund-address") != "" {

		// Parse bond refund address
		bondRefundAddress = common.HexToAddress(c.String("refund-address"))

	} else {

		// Get wallet status
		wallet, err := staderClient.WalletStatus()
		if err != nil {
			return err
		}

		// Prompt for node address
		if cliutils.Confirm(fmt.Sprintf("Would you like to refund your RPL bond to your node account (%s)?", wallet.AccountAddress.Hex())) {
			bondRefundAddress = wallet.AccountAddress
		} else {

			// Prompt for custom address
			inputAddress := cliutils.Prompt("Please enter the address to refund your RPL bond to:", "^0x[0-9a-fA-F]{40}$", "Invalid address")
			bondRefundAddress = common.HexToAddress(inputAddress)

		}

	}

	// Check if node can leave the oracle DAO
	canLeave, err := staderClient.CanLeaveTNDAO()
	if err != nil {
		return err
	}
	if !canLeave.CanLeave {
		fmt.Println("Cannot leave the oracle DAO:")
		if canLeave.ProposalExpired {
			fmt.Println("The proposal for you to leave the oracle DAO does not exist or has expired.")
		}
		if canLeave.InsufficientMembers {
			fmt.Println("There are not enough members in the oracle DAO to allow a member to leave.")
		}
		return nil
	}

	// Assign max fees
	err = gas.AssignMaxFeeAndLimit(canLeave.GasInfo, staderClient, c.Bool("yes"))
	if err != nil {
		return err
	}

	// Prompt for confirmation
	if !(c.Bool("yes") || cliutils.Confirm(fmt.Sprintf("Are you sure you want to leave the oracle DAO and refund your RPL bond to %s? This action cannot be undone!", bondRefundAddress.Hex()))) {
		fmt.Println("Cancelled.")
		return nil
	}

	// Leave the oracle DAO
	response, err := staderClient.LeaveTNDAO(bondRefundAddress)
	if err != nil {
		return err
	}

	fmt.Printf("Leaving oracle DAO...\n")
	cliutils.PrintTransactionHash(staderClient, response.TxHash)
	if _, err = staderClient.WaitForTransaction(response.TxHash); err != nil {
		return err
	}

	// Log & return
	fmt.Println("Successfully left the oracle DAO.")
	return nil

}
