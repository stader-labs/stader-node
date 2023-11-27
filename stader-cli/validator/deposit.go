package validator

import (
	"fmt"
	"math/big"
	"strconv"

	"github.com/stader-labs/stader-node/shared/services/gas"
	"github.com/stader-labs/stader-node/shared/utils/log"
	"github.com/stader-labs/stader-node/shared/utils/math"
	"github.com/stader-labs/stader-node/stader-cli/node"

	"github.com/stader-labs/stader-node/shared/services/stader"
	cliutils "github.com/stader-labs/stader-node/shared/utils/cli"
	"github.com/stader-labs/stader-node/stader-lib/utils/eth"
	"github.com/urfave/cli"
)

func nodeDeposit(c *cli.Context) error {

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

	numValidators := c.Uint64("num-validators")
	autoConfirm := c.Bool("yes")

	baseAmountInEth := 4
	baseAmount := eth.EthToWei(4.0)

	// Check to see if eth2 is synced
	syncResponse, err := staderClient.NodeSync()
	if err != nil {
		fmt.Printf("%s**WARNING**: Can't verify the sync status of your consensus client.\nYOU WILL LOSE ETH if your validator is activated before it is fully synced.\n"+
			"Reason: %s\n%s", log.ColorRed, err, log.ColorReset)
	} else {
		if syncResponse.BcStatus.PrimaryClientStatus.IsSynced {
			fmt.Printf("Your consensus client is synced, you may safely create a validator.\n")
		} else if syncResponse.BcStatus.FallbackEnabled {
			if syncResponse.BcStatus.FallbackClientStatus.IsSynced {
				fmt.Printf("Your fallback consensus client is synced, you may safely create a validator.\n")
			} else {
				fmt.Printf("%s**WARNING**: neither your primary nor fallback consensus clients are fully synced.\nYOU WILL LOSE ETH if your validator is activated before they are fully synced.\n%s", log.ColorRed, log.ColorReset)
			}
		} else {
			fmt.Printf("%s**WARNING**: your primary consensus client is either not fully synced or offline and you do not have a fallback client configured.\nYOU WILL LOSE ETH if your validator is activated before it is fully synced.\n%s", log.ColorRed, log.ColorReset)
		}
	}

	// Get node status
	status, err := staderClient.NodeStatus()
	if err != nil {
		return err
	}

	fmt.Printf(
		"The node %s%s%s has a balance of %.6f SD.\n\n",
		log.ColorBlue,
		status.AccountAddress,
		log.ColorReset,
		math.RoundDown(eth.WeiToEth(status.AccountBalances.Sd), 18))

	canNodeDepositResponse, err := staderClient.CanNodeDeposit(baseAmount, big.NewInt(int64(numValidators)), true)
	if err != nil {
		return err
	}

	if canNodeDepositResponse.InsufficientBalance {
		fmt.Printf("Account does not have enough balance!")
		return nil
	}
	if canNodeDepositResponse.DepositPaused {
		fmt.Printf("Deposit is paused")
		return nil
	}

	sdStatus := canNodeDepositResponse.SdStatus
	amountToCollateral := new(big.Int).Sub(sdStatus.SdCollateralRequireAmount, sdStatus.SdCollateralCurrentAmount)

	if amountToCollateral.Cmp(big.NewInt(0)) >= 1 {
		fmt.Printf(
			"The node %s%s%s need %.6f SD to meet collateral require.\n\n",
			log.ColorBlue,
			status.AccountAddress,
			log.ColorReset,
			math.RoundDown(eth.WeiToEth(amountToCollateral), 18))
	}

	utilityAmount := big.NewInt(0)

	if sdStatus.NotEnoughSdCollateral {
		// had SD but not in contract

		// User had enough to deposit-sd
		// 1. deposit-sd
		// 2. call deposit again
		if amountToCollateral.Cmp(sdStatus.SdBalance) <= 0 {
			err = node.NodeDepositSdWithAmount(staderClient, amountToCollateral, autoConfirm, 0)
			if err != nil {
				return err
			}

			return nodeDeposit(c)
		} else {
			// User had not enough to deposit-sd
			// Check if can borrow here
			// ask utilize or deposit-sd

			minUtility := new(big.Int).Sub(sdStatus.SdCollateralRequireAmount, sdStatus.SdCollateralCurrentAmount)

			// Max
			maxUtility := new(big.Int).Sub(sdStatus.SdMaxCollateralAmount, sdStatus.SdUtilityBalance)

			if minUtility.Cmp(sdStatus.PoolAvailableSDBalance) >= 0 {
				fmt.Printf("Pool available SD: %s not enough to min utility : %s \n", sdStatus.PoolAvailableSDBalance.String(), minUtility.String())
				return nil
			}

			// Set max to pool available
			if sdStatus.PoolAvailableSDBalance.Cmp(maxUtility) <= 0 {
				fmt.Printf("Pool available SD: %f, max utility : %f \n", eth.WeiToEth(sdStatus.PoolAvailableSDBalance), eth.WeiToEth(maxUtility))
				maxUtility = sdStatus.PoolAvailableSDBalance
			}

			min := eth.WeiToEth(minUtility)
			max := eth.WeiToEth(maxUtility)

			fmt.Printf("Min utility %+v max %+v \n", min, max)

			var _utilityAmount int
			msg := fmt.Sprintf("Please enter a valid number in range %f <> %f.", min, max)
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
					continue
				}

				break
			}

			utilityAmount = eth.EthToWei(float64(_utilityAmount))

			if !cliutils.Confirm(fmt.Sprintf("You're about to utility %d SD: ", _utilityAmount)) {
				fmt.Printf("Cancel \n")
				return nil
			}
		}
	}

	if canNodeDepositResponse.MaxValidatorLimitReached {
		fmt.Printf("Max validator limit reached")
		return nil
	}
	if canNodeDepositResponse.InputKeyLimitReached {
		fmt.Printf("You can only add %d keys at a time\n", canNodeDepositResponse.InputKeyLimit)
		return nil
	}

	//Assign max fees
	err = gas.AssignMaxFeeAndLimit(canNodeDepositResponse.GasInfo, staderClient, c.Bool("yes"))
	if err != nil {
		return err
	}

	// Prompt for confirmation
	if !(c.Bool("yes") || cliutils.Confirm(fmt.Sprintf(
		"You are about to utility %f SD to create %d validators.\n"+
			"%sARE YOU SURE YOU WANT TO DO THIS?!%s",
		eth.WeiToEth(utilityAmount), numValidators,
		log.ColorYellow,
		log.ColorReset))) {
		fmt.Println("Cancelled.")
		return nil
	}
	// 1 eth max per val0.8

	// Prompt for confirmation
	if !(c.Bool("yes") || cliutils.Confirm(fmt.Sprintf(
		"You are about to deposit %d ETH to create %d validators.\n"+
			"%sARE YOU SURE YOU WANT TO DO THIS? Running a validator is a long-term commitment, and this action cannot be undone!%s",
		uint64(baseAmountInEth)*numValidators, numValidators,
		log.ColorYellow,
		log.ColorReset))) {
		fmt.Println("Cancelled.")
		return nil
	}

	// Make deposit
	response, err := staderClient.NodeDeposit(baseAmount, big.NewInt(int64(numValidators)), utilityAmount, true)
	if err != nil {
		return err
	}

	fmt.Printf("Creating %d validators...\n", numValidators)
	cliutils.PrintTransactionHash(staderClient, response.TxHash)
	_, err = staderClient.WaitForTransaction(response.TxHash)
	if err != nil {
		return err
	}

	// Log & return
	fmt.Printf("The node deposit of %d ETH was made successfully!\n", uint64(baseAmountInEth)*numValidators)
	fmt.Printf("Total %d validators were created\n", numValidators)

	fmt.Println("Your validators are now in Initialized status.")
	fmt.Println("Once the ETH deposits have been matched by the remaining 28ETH, it will move to Deposited status.")
	fmt.Println("You can check the status of your validator with `stader-cli validator status`.")

	return nil

}
