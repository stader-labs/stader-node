package node

import (
	"fmt"

	"github.com/rocket-pool/rocketpool-go/utils/eth"
	"github.com/urfave/cli"

	"github.com/stader-labs/stader-node/shared/services/gas"
	"github.com/stader-labs/stader-node/shared/services/stader"
	cliutils "github.com/stader-labs/stader-node/shared/utils/cli"
	"github.com/stader-labs/stader-node/shared/utils/math"
)

func nodeBurn(c *cli.Context, amount float64, token string) error {

	// Get RP client
	staderNode, err := stader.NewClientFromCtx(c)
	if err != nil {
		return err
	}
	defer staderNode.Close()

	// Check and assign the EC status
	err = cliutils.CheckClientStatus(staderNode)
	if err != nil {
		return err
	}

	// Get amount in wei
	amountWei := eth.EthToWei(amount)

	// Check tokens can be burned
	canBurn, err := staderNode.CanNodeBurn(amountWei, token)
	if err != nil {
		return err
	}
	if !canBurn.CanBurn {
		fmt.Println("Cannot burn tokens:")
		if canBurn.InsufficientBalance {
			fmt.Printf("The node's %s balance is insufficient.\n", token)
		}
		if canBurn.InsufficientCollateral {
			fmt.Printf("There is insufficient ETH collateral to trade %s for.\n", token)
		}
		return nil
	}

	// Assign max fees
	err = gas.AssignMaxFeeAndLimit(canBurn.GasInfo, staderNode, c.Bool("yes"))
	if err != nil {
		return err
	}

	// Prompt for confirmation
	if !(c.Bool("yes") || cliutils.Confirm(fmt.Sprintf("Are you sure you want to burn %.6f %s for ETH?", math.RoundDown(eth.WeiToEth(amountWei), 6), token))) {
		fmt.Println("Cancelled.")
		return nil
	}

	// Burn tokens
	response, err := staderNode.NodeBurn(amountWei, token)
	if err != nil {
		return err
	}

	fmt.Printf("Burning tokens...\n")
	cliutils.PrintTransactionHash(staderNode, response.TxHash)
	if _, err = staderNode.WaitForTransaction(response.TxHash); err != nil {
		return err
	}

	// Log & return
	fmt.Printf("Successfully burned %.6f %s for ETH.\n", math.RoundDown(eth.WeiToEth(amountWei), 6), token)
	return nil

}
