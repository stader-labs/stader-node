package node

import (
	"errors"
	"fmt"
	"math/big"

	"github.com/stader-labs/stader-node/shared/services/gas"
	"github.com/stader-labs/stader-node/shared/services/stader"
	"github.com/stader-labs/stader-node/shared/types/api"
	cliutils "github.com/stader-labs/stader-node/shared/utils/cli"
	"github.com/stader-labs/stader-node/shared/utils/log"
	"github.com/stader-labs/stader-node/stader-lib/utils/eth"
	"github.com/stader-labs/stader-node/stader-lib/utils/sd"
	"github.com/urfave/cli"
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

	if sdStatusResponse.SDStatus.AlreadyLiquidated {
		fmt.Printf("Your node is under the liquidation process and you can not utilize SD at the moment.")
		return nil
	}

	healthFactor := eth.WeiToEth(sdStatusResponse.SDStatus.HealthFactor)
	if healthFactor < 1 {
		fmt.Printf(
			"Your node has a Health Factor of %.6f which is less than 1 and you can not utilize SD at the moment. Please repay your utilize SD and ensure your Health Factor is greater than 1 to avoid liquidation.\n\n",
			healthFactor)

		fmt.Printf("Note: Please repay your utilized SD by using the following command: stader-cli sd repay --amount <amount of SD to be repaid>.\n")

		return nil
	}

	if sdStatusResponse.SDStatus.SdCollateralRequireAmount.Cmp(big.NewInt(0)) == 0 {
		fmt.Printf("Please add a validator to your node first before utilizing SD from the SD Utility Pool. Execute the following command to add a validator to your node: ~/bin/stader-cli validator deposit --num-validators <number of validators you wish to add> \n")
		return nil
	}

	minUtility := eth.EthToWei(1)
	maxUtility := GetMaxUtility(sdStatusResponse.SDStatus)

	amountWei, err := PromptChooseUtilityAmount(sdStatusResponse.SDStatus, minUtility, maxUtility)
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

	rate := sdStatusResponse.SDStatus.UtilizationRate
	// Prompt for confirmation
	if !(c.Bool("yes") || cliutils.Confirm(fmt.Sprintf(
		"Are you sure you want to use %s from the SD Utility Pool?.\nNote: A Utilization fee of %s%s will be applied to the utilized SD from the SD Utility Pool.", eth.DisplayAmountInUnits(amountWei, "sd"), rate.String(), "%"))) {
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

	// Log & return
	fmt.Printf("Successfully deposited %s to the collateral contract by utilizing SD from the SD Utility Pool.\n", eth.DisplayAmountInUnits(amountWei, "sd"))

	return nil
}

func GetMinUtility(sdStatus *api.SdStatusResponse) *big.Int {
	totalCollateral := new(big.Int).Add(sdStatus.SdUtilizedBalance, sdStatus.SdCollateralCurrentAmount)

	minUtility := new(big.Int).Sub(sdStatus.SdCollateralRequireAmount, totalCollateral)

	if minUtility.Cmp(big.NewInt(0)) < 0 {
		minUtility = big.NewInt(0)
	}

	return minUtility
}

func GetMaxUtility(sdStatus *api.SdStatusResponse) *big.Int {
	maxUtility := new(big.Int).Sub(sdStatus.SdMaxUtilizableAmount, sdStatus.SdUtilizedBalance)

	if maxUtility.Cmp(sdStatus.PoolAvailableSDBalance) > 0 {
		maxUtility = sdStatus.PoolAvailableSDBalance
	}

	return maxUtility
}

func PromptChooseUtilityAmount(sdStatus *api.SdStatusResponse, minUtility, maxUtility *big.Int) (*big.Int, error) {
	// 1. If the pool had enough SD
	if minUtility.Cmp(sdStatus.PoolAvailableSDBalance) > 0 {
		return nil, errors.New("There is not sufficient free SD in the SD Utility Pool for utilization at the moment. Please try again later when there is enough free SD in the SD Utility Pool")
	}

	// 2. If user had enough Eth
	if minUtility.Cmp(maxUtility) >= 0 {
		return nil, errors.New("Do not had enough ETH bond to utility")
	}

	// Set maxSd to pool available
	if sdStatus.PoolAvailableSDBalance.Cmp(maxUtility) <= 0 {
		maxUtility = sdStatus.PoolAvailableSDBalance
	}

	minSd := eth.WeiToEth(minUtility)
	maxSd := eth.WeiToEth(maxUtility)

	msg := fmt.Sprintf(`
%sPlease enter the amount of SD you wish to utilize from the SD Utility Pool:%s
SD Utility Pool balance: %s
Minimum utilization amount: %s 
Maximum utilization amount: %s
`, log.ColorYellow, log.ColorReset, eth.DisplayAmountInUnits(sdStatus.PoolAvailableSDBalance, "sd"), eth.DisplayAmountInUnits(minUtility, "sd"), eth.DisplayAmountInUnits(maxUtility, "sd"))

	errMsg := fmt.Sprintf("%sInvalid input, please specify an amount within %f and %f SD range%s", log.ColorRed, minSd, maxSd, log.ColorReset)

	utilityAmount, err := sd.PromptChooseSDWithMaxMin(msg, errMsg, minUtility, maxUtility)
	if err != nil {
		return nil, err
	}

	return utilityAmount, nil
}

func PromptChooseSelfBondAmount(sdStatus *api.SdStatusResponse) (*big.Int, error) {
	totalCollateral := new(big.Int).Add(sdStatus.SdCollateralCurrentAmount, sdStatus.SdUtilizedBalance)

	amountToCollateralRemain := new(big.Int).Sub(sdStatus.SdCollateralRequireAmount, totalCollateral)

	sdRewardEligibleRemain := new(big.Int).Sub(sdStatus.SdRewardEligible, totalCollateral)

	if amountToCollateralRemain.Cmp(big.NewInt(0)) < 0 {
		amountToCollateralRemain = big.NewInt(0)
	}

	minSd := eth.WeiToEth(amountToCollateralRemain)
	maxSd := eth.WeiToEth(sdRewardEligibleRemain)

	msg := fmt.Sprintf(`Minimum bond: %s 
Maximum bond: %s

%sPlease enter the amount of SD you wish to deposit as collateral:%s`, eth.DisplayAmountInUnits(amountToCollateralRemain, "sd"), eth.DisplayAmountInUnits(sdRewardEligibleRemain, "sd"), log.ColorYellow, log.ColorReset)

	errMsg := fmt.Sprintf("Invalid input, please specify an amount within %f and %f SD range", minSd, maxSd)

	utilityAmount, err := sd.PromptChooseSDWithMaxMin(msg, errMsg, amountToCollateralRemain, sdRewardEligibleRemain)
	if err != nil {
		return nil, err
	}

	return utilityAmount, nil
}
