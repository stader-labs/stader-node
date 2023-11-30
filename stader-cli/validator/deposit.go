package validator

import (
	"fmt"
	"math/big"

	"github.com/stader-labs/stader-node/shared/services/gas"
	"github.com/stader-labs/stader-node/shared/utils/log"
	"github.com/stader-labs/stader-node/shared/utils/math"
	"github.com/stader-labs/stader-node/stader-cli/node"

	"github.com/stader-labs/stader-node/shared/services/stader"
	cliutils "github.com/stader-labs/stader-node/shared/utils/cli"
	"github.com/stader-labs/stader-node/stader-lib/utils/eth"
	"github.com/urfave/cli"
)

const Decimal = 18

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
		fmt.Printf("Account does not have enough ETH balance!")
		return nil
	}
	if canNodeDepositResponse.DepositPaused {
		fmt.Printf("Deposit is paused")
		return nil
	}

	sdStatus := canNodeDepositResponse.SdStatusResponse
	amountToCollateral := new(big.Int).Sub(sdStatus.SdCollateralRequireAmount, sdStatus.SdCollateralCurrentAmount)

	fmt.Printf(
		"The node %s%s%s current had %.6f SD in collateral.\n\n",
		log.ColorBlue,
		status.AccountAddress,
		log.ColorReset,
		math.RoundDown(eth.WeiToEth(sdStatus.SdCollateralCurrentAmount), Decimal))

	fmt.Printf(
		"The node %s%s%s need %.6f SD in collateral after deposit.\n\n",
		log.ColorBlue,
		status.AccountAddress,
		log.ColorReset,
		math.RoundDown(eth.WeiToEth(sdStatus.SdCollateralRequireAmount), Decimal))

	fmt.Printf(
		"The node %s%s%s can had max %.6f SD in collateral after deposit.\n\n",
		log.ColorBlue,
		status.AccountAddress,
		log.ColorReset,
		math.RoundDown(eth.WeiToEth(sdStatus.SdMaxCollateralAmount), Decimal))

	fmt.Printf(
		"The node %s%s%s current utility %.6f SD.\n\n",
		log.ColorBlue,
		status.AccountAddress,
		log.ColorReset,
		math.RoundDown(eth.WeiToEth(sdStatus.SdUtilizerLatestBalance), Decimal))

	if amountToCollateral.Cmp(big.NewInt(0)) >= 1 {
		fmt.Printf(
			"The node %s%s%s need %.6f SD to meet collateral require.\n\n",
			log.ColorBlue,
			status.AccountAddress,
			log.ColorReset,
			math.RoundDown(eth.WeiToEth(amountToCollateral), Decimal))
	}

	utilityAmount := big.NewInt(0)

	if sdStatus.NotEnoughSdCollateral {
		fmt.Printf(
			"The node %s%s%s had not enough SD in collateral please deposit SD.\n\n",
			log.ColorBlue,
			status.AccountAddress,
			log.ColorReset,
		)

		ops := []string{"Deposit from node wallet", "Utility SD"}
		i, _ := cliutils.Select("Choose option", ops)

		switch i {
		case 0:
			if status.AccountBalances.Sd.Cmp(amountToCollateral) < 0 {
				cliutils.PrintError("You do not had enough SD in your wallet. Please deposit and try again")
				return nil
			}

			err = node.NodeDepositSdWithAmount(staderClient, amountToCollateral, autoConfirm, 0)
			if err != nil {
				return err
			}

			return nodeDeposit(c)

		case 1:
			utilityAmount = node.PromChooseUtilityAmount(sdStatus)

			if !cliutils.Confirm(fmt.Sprintf("You're about to utility %f SD: ", eth.WeiToEth(utilityAmount))) {
				fmt.Printf("Cancel \n")
				return nil
			}
		default:
			return nil
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
		"You are about to deposit %d ETH to create %d validators.\n"+
			"%sARE YOU SURE YOU WANT TO DO THIS? Running a validator is a long-term commitment, and this action cannot be undone!%s",
		uint64(baseAmountInEth)*numValidators, numValidators,
		log.ColorYellow,
		log.ColorReset))) {
		fmt.Println("Cancelled.")
		return nil
	}

	// Make deposit
	response, err := staderClient.NodeDeposit(baseAmount, big.NewInt(int64(numValidators)), utilityAmount, false)
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
