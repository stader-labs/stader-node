package validator

import (
	"fmt"
	"math/big"

	"github.com/stader-labs/stader-node/shared/services/gas"
	"github.com/stader-labs/stader-node/shared/utils/log"
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

	baseAmount := eth.EthToWei(eth.BaseAmountInEth)

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

	// Get node SD status
	sdStatusResp, err := staderClient.GetSDStatus(big.NewInt(int64(numValidators)))
	if err != nil {
		return err
	}

	userBalance := status.AccountBalances.ETH
	amountToSend := new(big.Int).Mul(eth.EthToWei(eth.BaseAmountInEth), big.NewInt(int64(numValidators)))

	if userBalance.Cmp(amountToSend) < 0 {
		fmt.Printf("You don't have sufficient ETH in your Operator Address to add validators. Please deposit ETH into your Operator Address and try again to add validators to your node.")
		return nil
	}

	// Prompt for confirmation
	if !(c.Bool("yes") || cliutils.Confirm(fmt.Sprintf(
		"You are about to deposit %d ETH to create %d validators.\n"+
			"%sAre you sure you want to do this? Running a validator is a long-term commitment, and this action cannot be undone!%s",
		uint64(eth.BaseAmountInEth)*numValidators, numValidators,
		log.ColorYellow,
		log.ColorReset))) {
		fmt.Println("Cancelled.")
		return nil
	}

	sdStatus := sdStatusResp.SDStatus

	utilityAmount := big.NewInt(0)

	if sdStatus.NotEnoughSdCollateral {
		ops := []string{"Utilize SD from the SD Utility Pool", "Use your own SD (Self Bond)"}
		i, _ := cliutils.Select("You do not have sufficient SD collateral to add validators. Please choose one of the following options to add SD collateral. Enter 1 or 2:", ops)

		switch i {
		case 0:
			utilityAmount, err = node.PromptChooseUtilityAmount(sdStatus)
			if err != nil {
				return err
			}

			if !cliutils.Confirm(fmt.Sprintf("Are you sure you want to use %s from the Utility Pool? [Y/N] \nNote: A Utilization Fee of %s%s will be applied to the utilized SD from the Utility Pool. ", eth.DisplayAmountInUnits(utilityAmount, "sd"), sdStatus.UtilizationRate.String(), "%")) {
				fmt.Printf("Cancelled\n")
				return nil
			}
		case 1:
			selfBondAmount, err := node.PromptChooseSelfBondAmount(sdStatus)
			if err != nil {
				return err
			}

			if status.AccountBalances.Sd.Cmp(selfBondAmount) < 0 {
				fmt.Printf("You don't have sufficient SD in your Operator Address to add as collateral. Please deposit SD into your Operator Address and try again.\n")
				return nil
			}

			if !cliutils.Confirm(fmt.Sprintf("Are you sure you want to deposit %s as collateral?", eth.DisplayAmountInUnits(selfBondAmount, "sd"))) {
				fmt.Printf("Cancelled\n")
				return nil
			}

			autoConfirm := c.Bool("yes")
			nounce := c.GlobalUint64("nonce")

			err = node.DepositSdWithAmount(staderClient, selfBondAmount, autoConfirm, nounce)
			if err != nil {
				return err
			}

		default:
			return nil
		}

		fmt.Println("Continue with create validator...")
	}

	canNodeDepositResponse, err := staderClient.CanNodeDeposit(baseAmount, utilityAmount, big.NewInt(int64(numValidators)), true)
	if err != nil {
		return err
	}

	if canNodeDepositResponse.DepositPaused {
		fmt.Printf("Deposit is paused")
		return nil
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
	if utilityAmount.Cmp(big.NewInt(0)) == 0 {
		fmt.Printf("The node deposit of %d ETH was made successfully and total %d validators were created.\n", uint64(eth.BaseAmountInEth)*numValidators, numValidators)
	} else {
		fmt.Printf("The node deposit of %d ETH was made successfully and total %d validators were created by utilizing %s from the SD Utility Pool.\n", uint64(eth.BaseAmountInEth)*numValidators, numValidators, eth.DisplayAmountInUnits(utilityAmount, "sd"))
	}

	fmt.Println("Your validators are now in an Initialized state. Once the ETH deposits have been matched by the remaining 28 ETH, it will move to the Deposited state.")
	fmt.Println("You can check the status of your validator by executing the following command: `~/bin/stader-cli validator status`")

	return nil

}
