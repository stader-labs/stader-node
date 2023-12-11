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

	sdStatusResponse, err := staderClient.GetSDStatus(big.NewInt(0))
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

	if canNodeUtilizeSdResponse.NonTerminalValidators == 0 {
		fmt.Printf("Please add a validator to your node first before utilizing SD from a Utility Pool. Execute the following command to add a validator to your node: stader-cli validator deposit --num-validators <number of validators you wish to add> \n")
		return nil
	}

	err = gas.AssignMaxFeeAndLimit(canNodeUtilizeSdResponse.GasInfo, staderClient, c.Bool("yes"))
	if err != nil {
		return err
	}

	// Prompt for confirmation
	if !(c.Bool("yes") || cliutils.Confirm(fmt.Sprintf(
		"Are you sure you want to use %f SD from the utility pool? (y/n). Note: A Utilization fee of %.6f APR will be applied to the utilized SD from the utility pool.\n", eth.WeiToEth(amountWei), 0.5))) {
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
		return nil, errors.New("There is not sufficient free SD in the Utility Pool for utilization at the moment. Please try again later when there is enough free SD in the Utility Pool")
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

	msg := fmt.Sprintf(`Please enter the amount of SD you wish to utilize from the SD Utility Pool:
SD Utility Pool balance: %f SD
Minimum utilization amount: %f SD 
Maximum utilization amount: %f SD`, eth.WeiToEth(sdStatus.PoolAvailableSDBalance), min, max)

	errMsg := fmt.Sprintf("Invalid input, please specify an amount within %f and %f SD range\n", min, max)

	for {
		s := cliutils.Prompt(
			msg,
			"^[1-9][0-9]*$",
			errMsg)

		_utilityAmount, err = strconv.Atoi(s)
		if err != nil {
			fmt.Println(errMsg)
			continue
		}

		if _utilityAmount < int(min) || _utilityAmount > int(max) {
			fmt.Println(errMsg)
			continue
		}

		break
	}

	utilityAmount := eth.EthToWei(float64(_utilityAmount))

	return utilityAmount, nil
}

func PromptChooseSelfBondAmount(sdStatus *api.SdStatusResponse) (*big.Int, error) {

	amountToCollateralRemain := new(big.Int).Sub(sdStatus.SdCollateralRequireAmount, sdStatus.SdCollateralCurrentAmount)

	sdRewardEligibleRemain := new(big.Int).Sub(sdStatus.SdRewardEligible, sdStatus.SdCollateralCurrentAmount)

	min := eth.WeiToEth(amountToCollateralRemain)
	max := eth.WeiToEth(sdRewardEligibleRemain)

	var _utilityAmount int

	var err error

	msg := fmt.Sprintf(`Please enter the amount of SD you wish to deposit as collateral.
Minimum bond: %f SD 
Maximum bond: %f SD`, min, max)

	errMsg := fmt.Sprintf("Invalid input, please specify an amount within %f and %f SD range\n", min, max)

	for {
		s := cliutils.Prompt(
			msg,
			"^[1-9][0-9]*$",
			errMsg)

		_utilityAmount, err = strconv.Atoi(s)
		if err != nil {
			fmt.Println(errMsg)
			continue
		}

		if _utilityAmount < int(min) || _utilityAmount > int(max) {
			fmt.Println(errMsg)
			continue
		}

		break
	}

	utilityAmount := eth.EthToWei(float64(_utilityAmount))

	return utilityAmount, nil
}
