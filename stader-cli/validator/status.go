package validator

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
)

func getValidatorStatus(c *cli.Context) error {

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

	if !status.Registered {
		fmt.Printf("The node is not registered with Stader. Please use the %sstader-cli node register%s to register with Stader", log.ColorGreen, log.ColorReset)
		return nil
	}

	fmt.Printf("The Operator has registered a total of %d validators\n\n", len(status.ValidatorInfos))

	if status.TotalValidatorClRewards.Cmp(big.NewInt(0)) > 0 {
		fmt.Printf(
			"The Operator has a total of %.6f ETH as CL rewards for all validators. These rewards are sent to the claim vault periodically by Stader. Once it is sent to the claim vault, the operator can use the %sstader-cli node claim-rewards%s command to claim for all validators in one transaction\n", math.RoundDown(eth.WeiToEth(status.TotalValidatorClRewards), 6), log.ColorGreen, log.ColorReset)
		fmt.Println("If the operator wishes to claim CL rewards by themselves, follow these steps:")
		fmt.Printf("1. Use the %sstader-cli node send-cl-rewards --validator-pub-key%s command to claim the CL rewards\n", log.ColorGreen, log.ColorReset)
		fmt.Printf("2. Use the %sstader-cli node claim-rewards%s command to claim the CL rewards from the claim vault to your operator reward address\n\n", log.ColorGreen, log.ColorReset)
	}

	fmt.Printf("%s=== Registered Validator Details ===%s\n", log.ColorGreen, log.ColorReset)

	if len(status.ValidatorInfos) == 0 {
		fmt.Printf("The node has no registered validators. Please use the %sstader-cli validator deposit%s command to register a validator with Stader\n\n", log.ColorGreen, log.ColorReset)
		return nil
	}

	fmt.Printf("Download all your validator details to a csv file using %sstader-cli validator export%scommand\n\n", log.ColorGreen, log.ColorReset)

	for i := 0; i < len(status.ValidatorInfos); i++ {
		fmt.Printf("%d)\n", i+1)
		validatorInfo := status.ValidatorInfos[i]
		validatorPubKey := types.BytesToValidatorPubkey(validatorInfo.Pubkey)
		fmt.Printf("-Validator Pub Key: %s\n\n", validatorPubKey)
		fmt.Printf("-Validator Status: %s\n\n", validatorInfo.StatusToDisplay)
		fmt.Printf("-Validator Withdraw Vault: %s\n\n", validatorInfo.WithdrawVaultAddress)
		if validatorInfo.WithdrawVaultRewardBalance.Int64() > 0 && !validatorInfo.CrossedRewardsThreshold {
			fmt.Printf("-Validator Consensus Layer Rewards: %.6f\n\n", math.RoundDown(eth.WeiToEth(validatorInfo.WithdrawVaultRewardBalance), 18))
		} else if validatorInfo.CrossedRewardsThreshold {
			fmt.Printf("If you have exited the validator, Please wait for Stader Oracles to settle your funds!\n\n")
		} else if validatorInfo.Status == 5 {
			fmt.Printf("Your validator has been successfully settled by the oracles. Please withdraw your exited settled funds using %sstader-cli node claim-rewards%s", log.ColorGreen, log.ColorReset)
		}

		if validatorInfo.DepositBlock.Int64() > 0 {
			fmt.Printf("-Deposit time: %s\n\n", validatorInfo.DepositTime.Format("2006-01-02 15:04:05"))
		}

		if validatorInfo.WithdrawnBlock.Int64() > 0 {
			fmt.Printf("-Withdraw Time: %s\n\n", validatorInfo.WithdrawnTime.Format("2006-01-02 15:04:05"))
		}

		fmt.Printf("\n\n")
	}

	return nil
}
