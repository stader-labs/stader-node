package node

import (
	"fmt"
	"github.com/stader-labs/stader-node/shared/utils/log"
	"github.com/stader-labs/stader-node/stader-lib/utils/eth"
	"github.com/urfave/cli"

	"github.com/stader-labs/stader-node/shared/services/stader"
	cliutils "github.com/stader-labs/stader-node/shared/utils/cli"
	"github.com/stader-labs/stader-node/shared/utils/math"
)

const ()

func getStatus(c *cli.Context) error {

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

	// Print what network we're on
	err = cliutils.PrintNetwork(staderClient)
	if err != nil {
		return err
	}

	// Get node status
	status, err := staderClient.NodeStatus()
	if err != nil {
		return err
	}

	// Account address & balances
	fmt.Printf("%s=== Account and Balances ===%s\n", log.ColorGreen, log.ColorReset)
	fmt.Printf(
		"The node %s%s%s has a balance of %.6f ETH.\n\n",
		log.ColorBlue,
		status.AccountAddressFormatted,
		log.ColorReset,
		math.RoundDown(eth.WeiToEth(status.AccountBalances.ETH), 6))
	fmt.Printf(
		"The node %s%s%s has a balance of %.6f SD.\n\n",
		log.ColorBlue,
		status.AccountAddressFormatted,
		log.ColorReset,
		math.RoundDown(eth.WeiToEth(status.AccountBalances.Sd), 18))
	fmt.Printf(
		"The node %s%s%s has a deposited %.6f SD as collateral.\n\n",
		log.ColorBlue,
		status.AccountAddressFormatted,
		log.ColorReset,
		math.RoundDown(eth.WeiToEth(status.DepositedSdCollateral), 18))

	fmt.Printf("%s=== Operator Registration Details ===%s\n", log.ColorGreen, log.ColorReset)

	if status.Registered {
		fmt.Printf("The node is registered with Stader. Below are node details:\n")
		fmt.Printf("Operator Id: %d\n", status.OperatorId)
		fmt.Printf("Operator Name: %s\n", status.OperatorName)
		fmt.Printf("Operator Reward Address: %s\n", status.OperatorRewardAddress.String())
	} else {
		fmt.Printf("The node is not registered with Stader. Please use the %sstader-cli node register%s to register with stdr", log.ColorGreen, log.ColorReset)
	}

	// Return
	return nil

}
