package node

import (
	"fmt"
	"math/big"

	"github.com/stader-labs/stader-node/shared/services/stader"
	cliutils "github.com/stader-labs/stader-node/shared/utils/cli"
	"github.com/stader-labs/stader-node/shared/utils/log"
	"github.com/stader-labs/stader-node/stader-lib/utils/eth"
	"github.com/urfave/cli"
)

func getNodeStatus(c *cli.Context) error {

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

	totalRegisteredValidators := status.TotalNonTerminalValidators
	totalRegisterableValidators := status.SdCollateralWorthValidators
	totalEthCollateral := totalRegisteredValidators.Int64() * 4

	totalUnclaimedSocializingPoolEth := big.NewInt(0)
	totalUnclaimedSocializingPoolSd := big.NewInt(0)
	for _, merkle := range status.UnclaimedSocializingPoolMerkles {
		ethRewards, ok := big.NewInt(0).SetString(merkle.Eth, 10)
		if !ok {
			return fmt.Errorf("error while converting eth rewards: %s to big int", merkle.Eth)
		}

		sdRewards, ok := big.NewInt(0).SetString(merkle.Sd, 10)
		if !ok {
			return fmt.Errorf("error while converting sd rewards: %s to big int", merkle.Sd)
		}

		totalUnclaimedSocializingPoolEth.Add(totalUnclaimedSocializingPoolEth, ethRewards)
		totalUnclaimedSocializingPoolSd.Add(totalUnclaimedSocializingPoolSd, sdRewards)
	}

	noOfValidatorsWhichWeCanRegisterBasedOnSdCollateral := totalRegisterableValidators.Int64() - totalRegisteredValidators.Int64()
	if noOfValidatorsWhichWeCanRegisterBasedOnSdCollateral < 0 {
		noOfValidatorsWhichWeCanRegisterBasedOnSdCollateral = 0
	}

	noOfValidatorsWeCanRegisterBasedOnEthBalance := int64(eth.WeiToEth(status.AccountBalances.ETH) / 4)

	noOfValidatorsWeCanRegister := noOfValidatorsWhichWeCanRegisterBasedOnSdCollateral
	if noOfValidatorsWhichWeCanRegisterBasedOnSdCollateral > noOfValidatorsWeCanRegisterBasedOnEthBalance {
		noOfValidatorsWeCanRegister = noOfValidatorsWeCanRegisterBasedOnEthBalance
	}

	fmt.Printf("%s=== Operator Registration Details ===%s\n", log.ColorGreen, log.ColorReset)

	if !status.Registered {
		fmt.Printf("The node is not registered with Stader. Please use the %sstader-cli node register%s to register with Stader", log.ColorGreen, log.ColorReset)
		return nil
	}

	fmt.Printf("Operator Id: %d\n\n", status.OperatorId)
	fmt.Printf("Operator Name: %s\n\n", status.OperatorName)
	fmt.Printf("Operator Address: %s\n\n", status.OperatorAddress.String())
	fmt.Printf("Operator Reward Address: %s\n\n", status.OperatorRewardAddress.String())
	if status.OperatorActive {
		fmt.Printf("Operator Status: Active\n\n")
	} else {
		fmt.Printf("Operator Status: Not Active\n\n")
	}

	fmt.Printf("The Operator has registered a total of %d validators (To view details of each validator, please use the `stader-cli validator status` command)\n\n", totalRegisteredValidators)

	if !status.OptedInForSocializingPool {
		fmt.Printf("The Operator has Opted Out for Socializing Pool\n\n")
	} else {
		fmt.Printf("The Operator has Opted In for Socializing Pool\n\n")
	}

	if totalUnclaimedSocializingPoolSd.Cmp(big.NewInt(0)) > 0 {
		fmt.Printf("The Operator has %s as unclaimed SD rewards. To claim SD rewards using the %sstader-cli node claim-sp-rewards%s command\n\n", eth.DisplayAmountInUnits(totalUnclaimedSocializingPoolSd, "sd"), log.ColorGreen, log.ColorReset)
	}

	if totalUnclaimedSocializingPoolEth.Cmp(big.NewInt(0)) > 0 {
		fmt.Printf("The Operator has %s as unclaimed Socializing Pool EL rewards. To claim Socialized EL rewards using the %sstader-cli node claim-sp-rewards%s command\n\n", eth.DisplayAmountInUnits(totalUnclaimedSocializingPoolEth, "eth"), log.ColorGreen, log.ColorReset)
	}

	if status.OperatorELRewardsAddressBalance.Cmp(big.NewInt(0)) > 0 {
		fmt.Printf(
			"The Operator has a total of %s as EL rewards for all validators.\n"+
				"These rewards are sent to the claim vault periodically by Stader.\n"+
				"Once it is sent to the claim vault, the operator can use the %sstader-cli node claim-rewards%s command to claim EL rewards to their operator reward address\n", eth.DisplayAmountInUnits(status.OperatorELRewardsAddressBalance, "eth"), log.ColorGreen, log.ColorReset)
		fmt.Println("If the operator wishes to claim EL rewards by themselves, follow these steps:")
		fmt.Printf("1. Use the %sstader-cli node send-el-rewards%s command to claim the EL rewards\n", log.ColorGreen, log.ColorReset)
		fmt.Printf("2. Use the %sstader-cli node claim-rewards%s command to claim the EL rewards from the claim vault to your operator reward address\n\n", log.ColorGreen, log.ColorReset)
	}

	operatorWithdrawableAmount := status.OperatorRewardCollectorBalance
	if status.OperatorWithdrawableEth.Cmp(operatorWithdrawableAmount) < 0 {
		operatorWithdrawableAmount = status.OperatorWithdrawableEth
	}

	if operatorWithdrawableAmount.Cmp(big.NewInt(0)) > 0 {
		fmt.Printf(
			"The Operator has aggregated total claims of %s in the claim vault\n",
			eth.DisplayAmountInUnits(operatorWithdrawableAmount, "eth"))
		fmt.Printf("To transfer the claims to your operator reward address use the %sstader-cli node claim-rewards%s command\n\n", log.ColorGreen, log.ColorReset)
	}

	fmt.Printf(
		"The Operator can register %d more validators considering the ETH balance and SD availability in the Operator Address and Utility Pool.\n\n",
		noOfValidatorsWeCanRegister)

	// Account address & balances

	fmt.Printf("%s=== Account and Balances ===%s\n", log.ColorGreen, log.ColorReset)
	fmt.Printf(
		"The Operator %s%s%s has a balance of %s.\n\n",
		log.ColorBlue,
		status.AccountAddress,
		log.ColorReset,
		eth.DisplayAmountInUnits(status.AccountBalances.ETH, "eth"))
	fmt.Printf(
		"The Operator %s%s%s has a balance of %s.\n\n",
		log.ColorBlue,
		status.AccountAddress,
		log.ColorReset,
		eth.DisplayAmountInUnits(status.AccountBalances.Sd, "sd"))

	fmt.Printf(
		"The Operator has a deposited %d ETH as collateral.\n\n",
		totalEthCollateral)

	// Get node SD status
	sdStatusResp, err := staderClient.GetSDStatus(big.NewInt(0))
	if err != nil {
		return err
	}

	collateralPct := 0.0
	sdStatus := sdStatusResp.SDStatus
	totalCollateral := new(big.Int).Add(sdStatus.SdCollateralCurrentAmount, sdStatus.SdUtilizedBalance)

	current := eth.WeiToEth(totalCollateral)
	require := eth.WeiToEth(sdStatus.SdCollateralRequireAmount)

	if require > 0 {
		collateralPct = current / require * 10
	}

	fmt.Printf(
		"The Operator has a deposited %s (%.6f%s) as collateral. Below is the break-up: \n",
		eth.DisplayAmountInUnits(totalCollateral, "sd"), collateralPct, "%")

	fmt.Printf(
		"Self-bonded: %s \n",
		eth.DisplayAmountInUnits(sdStatus.SdCollateralCurrentAmount, "sd"))

	fmt.Printf(
		"Utilized from the Utility Pool: %s.\n",
		eth.DisplayAmountInUnits(sdStatus.SdUtilizedBalance, "sd"))

	if totalRegisteredValidators.Cmp(big.NewInt(0)) != 0 {
		fmt.Printf(
			"Note: For the %d validator, the minimum SD collateral should be %s (%s) to be eligible for the SD rewards. Please ensure that the SD collateral percentage is greater than %s. The SD collateral snapshots are taken daily at a random block, and if the SD collateral value falls below the %s limit, the node operator will not earn SD rewards for that day.\n\n",
			totalRegisteredValidators,
			eth.DisplayAmountInUnits(sdStatus.SdCollateralRequireAmount, "sd"),
			"10%", "10%", "10%")
	} else {
		fmt.Println("")
	}

	fmt.Printf("%s=== SD utilization Details ===%s\n", log.ColorGreen, log.ColorReset)

	fmt.Printf("The Operator has utilized %s from the Utility Pool.\n\n", eth.DisplayAmountInUnits(sdStatus.SdUtilizedBalance, "sd"))

	fmt.Printf("The Operator has a current Utilization Position of %s. (including the utilization fee)\n",
		eth.DisplayAmountInUnits(sdStatus.SdUtilizerLatestBalance, "sd"))

	if sdStatus.SdUtilizerLatestBalance.Cmp(big.NewInt(0)) == 0 {
		fmt.Println("")
	} else {
		fmt.Printf("Note: For repayment of your utilized SD, please use the `stader-cli node repay-sd <amount to repay>` command.\n\n")
	}

	maxUtilizable := new(big.Int).Sub(sdStatus.SdMaxUtilizableAmount, sdStatus.SdUtilizedBalance)
	if maxUtilizable.Cmp(sdStatus.PoolAvailableSDBalance) > 0 {
		maxUtilizable = sdStatus.PoolAvailableSDBalance
	}

	if maxUtilizable.Sign() < 0 {
		maxUtilizable = big.NewInt(0)
	}

	fmt.Printf(
		"The Operator can utilize up to %s more.\nNote: Each Operator can utilize a maximum of 1 ETH worth SD per validator.\n\n",
		eth.DisplayAmountInUnits(maxUtilizable, "sd"))

	if sdStatus.SdUtilizedBalance.Cmp(big.NewInt(0)) != 0 {
		fmt.Printf(
			"The Operator has a Health Factor of %.6f. \nNote: Please ensure your Health Factor is greater than 1 to avoid liquidations.\n\n",
			eth.WeiToEth(sdStatus.HealthFactor))
	}

	fmt.Printf(
		"The Utility Pool currently has a balance of %s.\n\n",
		eth.DisplayAmountInUnits(sdStatus.PoolAvailableSDBalance, "sd"))

	fmt.Printf("Current Utilization Rate %s%s.\n", sdStatus.UtilizationRate.String(), "%")

	return nil
}
