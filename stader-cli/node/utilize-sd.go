package node

import (
	"errors"
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

	sdStatusResponse, err := staderClient.GetSDStatus(big.NewInt(0), false)
	if err != nil {
		return err
	}

	amountWei, err := PromptChooseUtilityAmount(sdStatusResponse.SDStatus)
	if err != nil {
		return err
	}

	canNodeUtilizeSdResponse, err := staderClient.CanNodeUtilizeSd(amountWei)
	if err != nil {
		return err
	}

	err = gas.AssignMaxFeeAndLimit(canNodeUtilizeSdResponse.GasInfo, staderClient, c.Bool("yes"))
	if err != nil {
		return err
	}

	// Prompt for confirmation
	if !(c.Bool("yes") || cliutils.Confirm(fmt.Sprintf(
		"Are you sure you want to utilize %f SD? ", eth.WeiToEth(amountWei)))) {
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
		minUtility = big.NewInt(0)
	}

	return minUtility
}

func GetMaxUtility(sdStatus *api.SdStatusResponse) *big.Int {
	maxUtility := new(big.Int).Sub(sdStatus.SdMaxUtilizableAmount, sdStatus.SdUtilizerLatestBalance)

	if maxUtility.Cmp(sdStatus.PoolAvailableSDBalance) > 0 {
		maxUtility = sdStatus.PoolAvailableSDBalance
	}

	return maxUtility
}

func PromptChooseUtilityAmount(sdStatus *api.SdStatusResponse) (*big.Int, error) {
	minUtility := GetMinUtility(sdStatus)
	maxUtility := GetMaxUtility(sdStatus)

	// 1. If the pool had enough SD
	if minUtility.Cmp(sdStatus.PoolAvailableSDBalance) > 0 {
		msg := fmt.Sprintf("Pool available SD: %f not enough to min utility : %f \n", eth.WeiToEth(sdStatus.PoolAvailableSDBalance), eth.WeiToEth(minUtility))

		return nil, errors.New(msg)
	}

	// 2. If user had enough Eth
	if minUtility.Cmp(maxUtility) > 0 {
		msg := fmt.Sprintf("Do not had enough ETH bond to utility : %s \n", minUtility.String())

		return nil, errors.New(msg)
	}

	// Set max to pool available
	if sdStatus.PoolAvailableSDBalance.Cmp(maxUtility) <= 0 {
		maxUtility = sdStatus.PoolAvailableSDBalance
	}

	min := eth.WeiToEth(minUtility)
	max := eth.WeiToEth(maxUtility)

	var _utilityAmount int

	var err error

	msg := fmt.Sprintf("Please enter a number of SD to utilize in range %f and %f: ", min, max)

	for {
		s := cliutils.Prompt(
			msg,
			"^[1-9][0-9]*$",
			msg)

		_utilityAmount, err = strconv.Atoi(s)
		if err != nil {
			fmt.Println("Please enter a valid number.")
			continue
		}

		if _utilityAmount < int(min) || _utilityAmount > int(max) {
			fmt.Printf("Invalid input, please specify an amount within %f and %f range:\n", min, max)
			continue
		}

		break
	}

	utilityAmount := eth.EthToWei(float64(_utilityAmount))

	return utilityAmount, nil
}
