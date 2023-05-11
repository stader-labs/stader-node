package node

import (
	"crypto/rand"
	"fmt"
	"math/big"

	"github.com/stader-labs/stader-node/shared/services/gas"
	"github.com/stader-labs/stader-node/shared/utils/log"

	"github.com/stader-labs/stader-node/stader-lib/utils/eth"
	"github.com/urfave/cli"

	"github.com/stader-labs/stader-node/shared/services/stader"
	cliutils "github.com/stader-labs/stader-node/shared/utils/cli"
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

	baseAmountInEth := 4
	baseAmount := eth.EthToWei(4.0)

	var salt *big.Int
	if c.String("salt") != "" {
		var success bool
		salt, success = big.NewInt(0).SetString(c.String("salt"), 0)
		if !success {
			return fmt.Errorf("Invalid validator salt: %s", c.String("salt"))
		}
	} else {
		buffer := make([]byte, 32)
		_, err = rand.Read(buffer)
		if err != nil {
			return fmt.Errorf("Error generating random salt: %w", err)
		}
		salt = big.NewInt(0).SetBytes(buffer)
	}

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

	canNodeDepositResponse, err := staderClient.CanNodeDeposit(baseAmount, salt, big.NewInt(int64(numValidators)), true)
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
	if canNodeDepositResponse.NotEnoughSdCollateral {
		fmt.Printf("Not enough SD as collateral")
		return nil
	}
	if canNodeDepositResponse.MaxValidatorLimitReached {
		fmt.Printf("Max validator limit reached")
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
	response, err := staderClient.NodeDeposit(baseAmount, salt, big.NewInt(int64(numValidators)), false)
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
	fmt.Println("You can check the status of your validator with `stader-cli node status`.")

	return nil

}
