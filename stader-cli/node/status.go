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
		totalEthCollateral)

	fmt.Printf(
		"The node %s%s%s has a deposited %.6f SD as collateral.\n\n",
		log.ColorBlue,
		status.AccountAddress,
		log.ColorReset,
		math.RoundDown(eth.WeiToEth(status.DepositedSdCollateral), 18))

	fmt.Printf(
		"The node %s%s%s can register %d more validators based on the ETH balance and the SD collateral provided.\n\n",
		log.ColorBlue,
		status.AccountAddress,
		log.ColorReset,
		noOfValidatorsWeCanRegister)

	fmt.Printf("%s=== Operator Registration Details ===%s\n", log.ColorGreen, log.ColorReset)

	if !status.Registered {
		fmt.Printf("The node is not registered with Stader. Please use the %sstader-cli node register%s to register with Stader", log.ColorGreen, log.ColorReset)
		return nil
	}

	fmt.Printf("The operator is registered with Stader. Below are operator details:\n")
	fmt.Printf("Operator Id: %d\n\n", status.OperatorId)
	fmt.Printf("Operator Name: %s\n\n", status.OperatorName)
	fmt.Printf("Operator Address: %s\n\n", status.OperatorAddress.String())
	fmt.Printf("Operator Reward Address: %s\n\n", status.OperatorRewardAddress.String())
	if status.OperatorActive {
		fmt.Printf("Operator Status: Active\n\n")
	} else {
		fmt.Printf("Operator Status: Not Active\n\n")
	}
	fmt.Printf("The Operator has registered a total of %d validators\n\n", len(status.ValidatorInfos))

	if totalUnclaimedSocializingPoolSd.Cmp(big.NewInt(0)) > 0 {
		fmt.Printf("The Operator reward address %s has %.6f SD as unclaimed SD rewards\n\n", status.OperatorAddress.String(), math.RoundDown(eth.WeiToEth(totalUnclaimedSocializingPoolSd), 18))
		fmt.Printf("To claim SD rewards using the %sstader-cli node claim-sp-rewards%s command\n\n", log.ColorGreen, log.ColorReset)
	}

	if !status.OptedInForSocializingPool {
		fmt.Printf("Operator has Opted Out for Socializing Pool\n\n")
		fmt.Printf("Operator Fee Recepient: %s\n\n", status.OperatorELRewardsAddress.String())
		if status.OperatorELRewardsAddressBalance.Cmp(big.NewInt(0)) > 0 {
			fmt.Printf(
				"The Operator %s%s%s fee recepient %s%s%s has unclaimed rewards of %.6f ETH.\n\n",
				log.ColorBlue,
				status.AccountAddress,
				log.ColorReset,
				log.ColorBlue,
				status.OperatorELRewardsAddress,
				log.ColorReset,
				math.RoundDown(eth.WeiToEth(status.OperatorELRewardsAddressBalance), 6))
			fmt.Printf("To claim fee recepient EL rewards use the %sstader-cli node claim-el-rewards%s command\n\n", log.ColorGreen, log.ColorReset)
		}
	} else {
		fmt.Printf("Operator has Opted In for Socializing Pool\n\n")

		if totalUnclaimedSocializingPoolSd.Cmp(big.NewInt(0)) > 0 {
			fmt.Printf("The Operator reward address %s has %.6f ETH as unclaimed EL rewards through socializing pool till\n\n", status.OperatorAddress.String(), math.RoundDown(eth.WeiToEth(totalUnclaimedSocializingPoolEth), 18))
			fmt.Printf("To claim Socializing pool EL rewards using the %sstader-cli node claim-sp-rewards%s command\n\n", log.ColorGreen, log.ColorReset)
		}
	}

	fmt.Printf("%s=== Registered Validator Details ===%s\n", log.ColorGreen, log.ColorReset)

	if totalRegisteredValidators.Int64() <= 0 {
		fmt.Printf("The node has no registered validators. Please use the %sstader-cli node deposit%s command to register a validator with Stader\n\n", log.ColorGreen, log.ColorReset)
		return nil
	}

	for i := 0; i < len(status.ValidatorInfos); i++ {
		fmt.Printf("%d)\n", i+1)
		validatorInfo := status.ValidatorInfos[i]
		validatorPubKey := types.BytesToValidatorPubkey(validatorInfo.Pubkey)
		fmt.Printf("-Validator Pub Key: %s\n\n", validatorPubKey)
		fmt.Printf("-Validator Status: %s\n\n", validatorInfo.StatusToDisplay)
		fmt.Printf("-Validator Withdraw Vault: %s\n\n", validatorInfo.WithdrawVaultAddress)
		if validatorInfo.WithdrawVaultRewardBalance.Int64() > 0 && !validatorInfo.CrossedRewardsThreshold {
			fmt.Printf("-Validator Skimmed Rewards: %.6f\n", math.RoundDown(eth.WeiToEth(validatorInfo.WithdrawVaultRewardBalance), 18))
			fmt.Printf("To claim skimmed rewards use the %sstader-cli node claim-cl-rewards %s command\n\n", log.ColorGreen, log.ColorReset)
		} else if validatorInfo.CrossedRewardsThreshold {
			fmt.Printf("Too many CL rewards to withdraw for validator %s.\n", validatorPubKey.String())
			fmt.Printf("If you have exited the validator, Please wait for Stader Oracles to settle your funds!\n")
			fmt.Printf("If you have not exited the validator, Please reach out to the Stader Team on discord!\n")
		}

		if validatorInfo.Status > 3 {
			fmt.Printf("-Deposit time: %s\n\n", validatorInfo.DepositTime.Format("2006-01-02 15:04:05"))
		}

		if validatorInfo.WithdrawnBlock.Int64() > 0 {
			fmt.Printf("-Withdraw Time: %s\n\n", validatorInfo.WithdrawnTime.Format("2006-01-02 15:04:05"))
		}

		fmt.Printf("\n\n")
	}

	return nil
}
