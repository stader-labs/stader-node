package node

import (
	"fmt"
	"math/big"
	"strconv"

	"github.com/urfave/cli"

	"github.com/stader-labs/stader-node/shared/services/gas"
	"github.com/stader-labs/stader-node/shared/services/stader"
	cliutils "github.com/stader-labs/stader-node/shared/utils/cli"
	"github.com/stader-labs/stader-node/stader-lib/utils/eth"
)

func utilizeSD(c *cli.Context) error {
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

	// Get stake mount
	amountInString := c.String("amount")

	amount, err := strconv.ParseFloat(amountInString, 64)
	if err != nil {
		return err
	}

	amountWei := eth.EthToWei(amount)

	canNodeUtilizeSdResponse, err := staderClient.CanNodeUtilizeSd(amountWei)
	if err != nil {
		return err
	}

	sdStatus := canNodeUtilizeSdResponse.SdStatusResponse

	// Max
	maxUtility := new(big.Int).Sub(sdStatus.SdMaxCollateralAmount, sdStatus.SdUtilizerLatestBalance)
	// 1.
	if sdStatus.PoolAvailableSDBalance.Cmp(amountWei) < 0 {
		fmt.Printf("The amount in pool is not enough: %s \n", sdStatus.PoolAvailableSDBalance.String())
		return nil
	}

	// 1.
	if maxUtility.Cmp(amountWei) < 0 {
		fmt.Printf("The max amount is : %s \n", maxUtility.String())
		return nil
	}

	err = gas.AssignMaxFeeAndLimit(canNodeUtilizeSdResponse.GasInfo, staderClient, c.Bool("yes"))
	if err != nil {
		return err
	}

	// Prompt for confirmation
	if !(c.Bool("yes") || cliutils.Confirm(fmt.Sprintln(
		"Are you sure you want to utilize SD?"))) {
		fmt.Println("Cancelled.")
		return nil
	}

	res, err := staderClient.NodeUtilizeSd(amountWei)
	if err != nil {
		return err
	}

	cliutils.PrintTransactionHash(staderClient, res.TxHash)

	if _, err = staderClient.WaitForTransaction(res.TxHash); err != nil {
		return err
	}

	return nil
}
