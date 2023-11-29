package node

import (
	"fmt"
	"math/big"
	"strconv"

	"github.com/urfave/cli"

	"github.com/stader-labs/stader-node/shared/services/gas"
	"github.com/stader-labs/stader-node/shared/services/stader"
	"github.com/stader-labs/stader-node/shared/types/api"
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

	// min
	minUtility := GetMinUtility(sdStatus)

	// Max
	maxUtility := GetMaxUtility(sdStatus)

	if maxUtility.Cmp(minUtility) < 0 {
		cliutils.PrintError("Insufficient ETH bonding")
		return nil
	}
	// 1. If there's enough SD in pool
	if sdStatus.PoolAvailableSDBalance.Cmp(amountWei) < 0 {
		fmt.Printf("The amount in pool: %f is not enough\n", eth.WeiToEth(sdStatus.PoolAvailableSDBalance))
		return nil
	}

	// 2. If not utilize to max amount
	if maxUtility.Cmp(amountWei) < 0 || minUtility.Cmp(amountWei) > 0 {
		cliutils.PrintError(fmt.Sprintf("Invalid input, please specify an amount within %f and %f range \n", eth.WeiToEth(minUtility), eth.WeiToEth(maxUtility)))
		return nil
	}

	err = gas.AssignMaxFeeAndLimit(canNodeUtilizeSdResponse.GasInfo, staderClient, c.Bool("yes"))
	if err != nil {
		return err
	}

	// Prompt for confirmation
	if !(c.Bool("yes") || cliutils.Confirm(fmt.Sprintf(
		"Are you sure you want to utilize %+v SD? ", amount))) {
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

func GetMinUtility(sdStatus *api.SdStatusResponse) *big.Int {
	minUtility := new(big.Int).Sub(sdStatus.SdCollateralRequireAmount, sdStatus.SdCollateralCurrentAmount)

	if minUtility.Cmp(big.NewInt(0)) < 0 {
		cliutils.PrintWarning("Over collateral")

		minUtility = big.NewInt(0)
	}

	return minUtility
}

func GetMaxUtility(sdStatus *api.SdStatusResponse) *big.Int {
	maxUtility := new(big.Int).Sub(sdStatus.SdMaxCollateralAmount, sdStatus.SdUtilizerLatestBalance)

	if maxUtility.Cmp(sdStatus.PoolAvailableSDBalance) > 0 {
		cliutils.PrintWarning("Pool not had enough SD")

		maxUtility = sdStatus.PoolAvailableSDBalance
	}

	return maxUtility
}
