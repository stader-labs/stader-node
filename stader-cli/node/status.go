package node

import (
	"fmt"
	"github.com/stader-labs/stader-node/shared/services/stader"
	cliutils "github.com/stader-labs/stader-node/shared/utils/cli"
	"github.com/stader-labs/stader-node/shared/utils/log"
	"github.com/stader-labs/stader-node/shared/utils/math"
	"github.com/stader-labs/stader-node/stader-lib/types"
	"github.com/stader-labs/stader-node/stader-lib/utils/eth"
	"github.com/urfave/cli"
	"math/big"
	"time"
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

	totalRegisteredValidators := status.TotalNonTerminalValidators
	totalRegisterableValidators := status.SdCollateralWorthValidators
	totalEthCollateral := status.TotalNonTerminalValidators.Mul(status.TotalNonTerminalValidators, big.NewInt(4))

	fmt.Printf("totalRegisteredValidators: %d\n", totalRegisteredValidators.Int64())
	fmt.Printf("totalRegisterableValidators: %d\n", totalRegisterableValidators.Int64())
	noOfValidatorsWhichWeCanRegisterBasedOnSdCollateral := totalRegisterableValidators.Int64() - totalRegisteredValidators.Int64()

	noOfValidatorsWeCanRegisterBasedOnEthBalance := int64(eth.WeiToEth(status.AccountBalances.ETH) / 4)

	fmt.Printf("noOfValidatorsWhichWeCanRegisterBasedOnSdCollateral: %d\n", noOfValidatorsWhichWeCanRegisterBasedOnSdCollateral)
	fmt.Printf("noOfValidatorsWeCanRegisterBasedOnEthBalance: %d\n", noOfValidatorsWeCanRegisterBasedOnEthBalance)
	noOfValidatorsWeCanRegister := noOfValidatorsWhichWeCanRegisterBasedOnSdCollateral
	if noOfValidatorsWhichWeCanRegisterBasedOnSdCollateral > noOfValidatorsWeCanRegisterBasedOnEthBalance {
		noOfValidatorsWeCanRegister = noOfValidatorsWeCanRegisterBasedOnEthBalance
	}
	fmt.Printf("noOfValidatorsWeCanRegister: %d\n", noOfValidatorsWeCanRegister)

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

	if status.SdCollateralRequestedToWithdraw.Cmp(big.NewInt(0)) > 0 {
		currentTime := time.Unix(time.Now().Unix(), 0)
		if status.SdCollateralWithdrawTime.Int64() > currentTime.Unix() {
			withdrawTime := time.Unix(status.SdCollateralWithdrawTime.Int64(), 0)
			timeRemaining := withdrawTime.Sub(currentTime)

			fmt.Printf(
				"The node %s%s%s %.6f SD in unbonding phase. You can withdraw it in %v \n\n",
				log.ColorBlue,
				status.AccountAddress,
				log.ColorReset,
				math.RoundDown(eth.WeiToEth(status.SdCollateralRequestedToWithdraw), 18), timeRemaining)
		} else {
			fmt.Printf(
				"The node %s%s%s can claim %.6f SD. You can use the %sstader-cli node claim-sd%s command \n\n",
				log.ColorBlue,
				status.AccountAddress,
				log.ColorReset,
				math.RoundDown(eth.WeiToEth(status.SdCollateralRequestedToWithdraw), 18), log.ColorGreen, log.ColorReset)
		}
	}

	fmt.Printf(
		"The node %s%s%s has a deposited %d Eth as collateral.\n\n",
		log.ColorBlue,
		status.AccountAddress,
		log.ColorReset,
		totalEthCollateral.Int64())

	fmt.Printf(
		"The node %s%s%s has a deposited %.6f SD as collateral.\n\n",
		log.ColorBlue,
		status.AccountAddress,
		log.ColorReset,
		math.RoundDown(eth.WeiToEth(status.DepositedSdCollateral), 18))

	fmt.Printf(
		"The node %s%s%s can register %d more validators based on the amount of SD collateral it has provided.\n\n",
		log.ColorBlue,
		status.AccountAddress,
		log.ColorReset,
		noOfValidatorsWeCanRegister)

	fmt.Printf("%s=== Operator Registration Details ===%s\n", log.ColorGreen, log.ColorReset)

	if !status.Registered {
		fmt.Printf("The node is not registered with Stader. Please use the %sstader-cli node register%s to register with Stader", log.ColorGreen, log.ColorReset)
		return nil
	}

	fmt.Printf("The node is registered with Stader. Below are node details:\n")
	fmt.Printf("Operator Id: %d\n\n", status.OperatorId)
	fmt.Printf("Operator Name: %s\n\n", status.OperatorName)
	fmt.Printf("Operator Address: %s\n\n", status.OperatorAddress.String())
	fmt.Printf("Operator Reward Address: %s\n\n", status.OperatorRewardAddress.String())
	if status.OperatorActive {
		fmt.Printf("Operator Status: Active\n\n")
	} else {
		fmt.Printf("Operator Status: Not Active\n\n")
	}
	if !status.OptedInForSocializingPool {
		fmt.Printf("Operator has Opted Out for Socializing Pool\n\n")
		fmt.Printf("Operator Execution layer reward vault: %s\n\n", status.OperatorELRewardsAddress.String())
		fmt.Printf(
			"The node %s%s%s execution layer reward vault has a balance %.6f ETH.\n\n",
			log.ColorBlue,
			status.AccountAddress,
			log.ColorReset,
			math.RoundDown(eth.WeiToEth(status.OperatorELRewardsAddressBalance), 6))
	} else {
		fmt.Printf("Operator has Opted In for Socializing Pool\n\n")
		fmt.Printf("Operator Socializing Pool Fee Recepient: %s\n\n", status.OperatorELRewardsAddress.String())
	}

	fmt.Printf("%s=== Registered Validator Details ===%s\n", log.ColorGreen, log.ColorReset)

	if totalRegisteredValidators.Int64() <= 0 {
		fmt.Printf("The node has no registered validators. Please use the %sstader-cli node deposit%s command to register a validator with Stader\n\n", log.ColorGreen, log.ColorReset)
		return nil
	}

	for i := 0; i < len(status.ValidatorInfos); i++ {
		fmt.Printf("%d)\n", i+1)
		validatorInfo := status.ValidatorInfos[i]
		fmt.Printf("-Validator Pub Key: %s\n\n", types.BytesToValidatorPubkey(validatorInfo.Pubkey))
		fmt.Printf("-Validator Status: %s\n\n", validatorInfo.StatusToDisplay)
		fmt.Printf("-Validator Withdraw Vault: %s\n\n", validatorInfo.WithdrawVaultAddress)
		if validatorInfo.WithdrawVaultRewardBalance.Int64() > 0 && !validatorInfo.CrossedRewardsThreshold {
			fmt.Printf("-Validator Skimmed Rewards: %.6f\n", math.RoundDown(eth.WeiToEth(validatorInfo.WithdrawVaultRewardBalance), 18))
			fmt.Printf("To withdraw skimmed rewards use the %sstader-cli node withdraw-cl-rewards%s command\n\n", log.ColorGreen, log.ColorReset)
		} else if validatorInfo.CrossedRewardsThreshold {
			fmt.Printf("-Validator Skimmed Rewards: Crossed threshold\n")
		}

		if validatorInfo.Status > 3 {
			fmt.Printf("-Deposit block: %s\n\n", validatorInfo.DepositBlock)
		}

		if validatorInfo.WithdrawnBlock.Int64() > 0 {
			// Validator has withdrawn
			if validatorInfo.WithdrawVaultWithdrawableBalance.Int64() > 0 {
				fmt.Printf("-Withdrawable Amount: %.6f\n", math.RoundDown(eth.WeiToEth(validatorInfo.WithdrawVaultWithdrawableBalance), 18))
				fmt.Printf("To withdraw exit amount use the %sstader-cli node settle-exit-funds%s command\n\n", log.ColorGreen, log.ColorReset)
			}
			fmt.Printf("-Withdraw block: %s\n\n", validatorInfo.WithdrawnBlock)
		}

		fmt.Printf("\n\n")
	}

	return nil
}
