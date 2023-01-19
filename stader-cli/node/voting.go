package node

import (
	"fmt"
	"strings"

	"github.com/ethereum/go-ethereum/common"
	"github.com/urfave/cli"

	"github.com/stader-labs/stader-node/shared/services/gas"
	"github.com/stader-labs/stader-node/shared/services/stader"
	cliutils "github.com/stader-labs/stader-node/shared/utils/cli"
)

func nodeSetVotingDelegate(c *cli.Context, nameOrAddress string) error {
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
	var address common.Address
	var addressString string
	if strings.Contains(nameOrAddress, ".") {
		response, err := staderClient.ResolveEnsName(nameOrAddress)
		if err != nil {
			return err
		}
		address = response.Address
		addressString = fmt.Sprintf("%s (%s)", nameOrAddress, address.Hex())
	} else {
		address, err = cliutils.ValidateAddress("delegate", nameOrAddress)
		if err != nil {
			return err
		}
		addressString = address.Hex()
	}

	// Get the gas estimation
	gasEstimate, err := staderClient.EstimateSetSnapshotDelegateGas(address)
	if err != nil {
		return err
	}

	// Assign max fees
	err = gas.AssignMaxFeeAndLimit(gasEstimate.GasInfo, staderClient, c.Bool("yes"))
	if err != nil {
		return err
	}

	// Prompt for confirmation
	if !(c.Bool("yes") || cliutils.Confirm(fmt.Sprintf("Are you sure you want %s to represent your node in Rocket Pool governance proposals?", addressString))) {
		fmt.Println("Cancelled.")
		return nil
	}

	// Set delegate
	response, err := staderClient.SetSnapshotDelegate(address)
	if err != nil {
		return err
	}

	fmt.Printf("Setting delegate...\n")
	cliutils.PrintTransactionHash(staderClient, response.TxHash)
	if _, err = staderClient.WaitForTransaction(response.TxHash); err != nil {
		return err
	}

	// Log & return
	fmt.Printf("The node's voting delegate was successfuly set to %s.\n", addressString)
	return nil

}

func nodeClearVotingDelegate(c *cli.Context) error {

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

	// Get the gas estimation
	gasEstimate, err := staderClient.EstimateClearSnapshotDelegateGas()
	if err != nil {
		return err
	}

	// Assign max fees
	err = gas.AssignMaxFeeAndLimit(gasEstimate.GasInfo, staderClient, c.Bool("yes"))
	if err != nil {
		return err
	}

	// Prompt for confirmation
	if !(c.Bool("yes") || cliutils.Confirm("Are you sure you remove your node's current delegate address for voting on governance proposals?")) {
		fmt.Println("Cancelled.")
		return nil
	}

	// Set delegate
	response, err := staderClient.ClearSnapshotDelegate()
	if err != nil {
		return err
	}

	fmt.Printf("Removing delegate...\n")
	cliutils.PrintTransactionHash(staderClient, response.TxHash)
	if _, err = staderClient.WaitForTransaction(response.TxHash); err != nil {
		return err
	}

	// Log & return
	fmt.Println("The node's voting delegate has been removed.")
	return nil

}
