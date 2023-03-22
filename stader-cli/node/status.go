package node

import (
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/stader-labs/stader-node/shared/utils/log"
	"github.com/stader-labs/stader-node/shared/utils/stdr"
	"github.com/stader-labs/stader-node/stader-lib/utils/eth"
	"github.com/urfave/cli"

	"github.com/stader-labs/stader-node/shared/services/stader"
	cliutils "github.com/stader-labs/stader-node/shared/utils/cli"
	"github.com/stader-labs/stader-node/shared/utils/math"
)

func getStatus(c *cli.Context) error {

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

	totalRegisteredValidators := len(status.ValidatorInfos)
	totalRegisterableValidators := status.SdCollateralWorthValidators
	noOfValidatorsWhichWeCanRegister := totalRegisterableValidators.Int64() - int64(totalRegisteredValidators)

	// Account address & balances
	fmt.Printf("%s=== Account and Balances ===%s\n", log.ColorGreen, log.ColorReset)
	fmt.Printf(
		"The node %s%s%s has a balance of %.6f ETH.\n\n",
		log.ColorBlue,
		status.AccountAddress,
		log.ColorReset,
		math.RoundDown(eth.WeiToEth(status.AccountBalances.ETH), 6))
	fmt.Printf(
		"The node %s%s%s has a balance of %.6f SD.\n\n",
		log.ColorBlue,
		status.AccountAddress,
		log.ColorReset,
		math.RoundDown(eth.WeiToEth(status.AccountBalances.Sd), 18))
	fmt.Printf(
		"The node %s%s%s has a deposited %.6f SD as collateral.\n\n",
		log.ColorBlue,
		status.AccountAddress,
		log.ColorReset,
		math.RoundDown(eth.WeiToEth(status.DepositedSdCollateral), 18))
	fmt.Printf(
		"The node %s%s%s has registered %d validators.\n\n",
		log.ColorBlue,
		status.AccountAddress,
		log.ColorReset,
		totalRegisteredValidators)
	fmt.Printf(
		"The node %s%s%s can register %d more validators based on the amount of SD collateral it has provided.\n\n",
		log.ColorBlue,
		status.AccountAddress,
		log.ColorReset,
		noOfValidatorsWhichWeCanRegister)

	fmt.Printf("%s=== Operator Registration Details ===%s\n", log.ColorGreen, log.ColorReset)

	if status.Registered {
		fmt.Printf("The node is registered with Stader. Below are node details:\n")
		fmt.Printf("Operator Id: %d\n\n", status.OperatorId)
		fmt.Printf("Operator Name: %s\n\n", status.OperatorName)
		fmt.Printf("Operator Reward Address: %s\n\n", status.OperatorRewardAddress.String())

		fmt.Printf("%s=== Registered Validator Details ===%s\n", log.ColorGreen, log.ColorReset)
		// display validators
		if totalRegisteredValidators > 0 {
			for i := 0; i < totalRegisteredValidators; i++ {
				fmt.Printf("%d)\n", i+1)
				validatorInfo := status.ValidatorInfos[i]
				fmt.Printf("-Validator Pub Key: %s\n\n", common.BytesToAddress(validatorInfo.Pubkey))
				fmt.Printf("-Validator Status: %s\n\n", stdr.ValidatorState[validatorInfo.Status])
				fmt.Printf("-Validator Withdraw Vault: %s\n\n", validatorInfo.WithdrawVaultAddress)

				if validatorInfo.Status > 3 {
					fmt.Printf("-Deposit time: %d\n\n", validatorInfo.DepositTime)
				}

				// Validator has withdrawn
				if validatorInfo.Status == 8 {
					fmt.Printf("-Withdrawn time: %d\n\n", validatorInfo.WithdrawnTime)
				}
				fmt.Printf("\n\n")
			}
		} else {
			fmt.Printf("The node has no registered validators. Please use the %sstader-cli node deposit%s command to register a validator with Stader", log.ColorGreen, log.ColorReset)
		}
	} else {
		fmt.Printf("The node is not registered with Stader. Please use the %sstader-cli node register%s to register with Stader", log.ColorGreen, log.ColorReset)
	}

	// Return
	return nil

}
