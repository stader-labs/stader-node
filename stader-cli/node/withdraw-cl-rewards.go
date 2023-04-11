package node

import (
	"fmt"

	"github.com/stader-labs/stader-node/shared/services/stader"
	cliutils "github.com/stader-labs/stader-node/shared/utils/cli"
	"github.com/stader-labs/stader-node/shared/utils/math"
	"github.com/stader-labs/stader-node/stader-lib/types"
	"github.com/stader-labs/stader-node/stader-lib/utils/eth"
	"github.com/urfave/cli"
)

func WithdrawClRewards(c *cli.Context, validatorPubKey types.ValidatorPubkey) error {
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

	canWithdrawClRewardsResponse, err := staderClient.CanWithdrawClRewards(validatorPubKey)
	if err != nil {
		return err
	}
	if canWithdrawClRewardsResponse.NoClRewards {
		fmt.Printf("No CL rewards to withdraw for validator %s\n", validatorPubKey.String())
		return nil
	}
	if canWithdrawClRewardsResponse.TooManyClRewards {
		fmt.Printf("Too many CL rewards to withdraw for validator %s\n. Please use stader-cli node settle-funds command to withdraw remaining amount", validatorPubKey.String())
		return nil
	}
	if canWithdrawClRewardsResponse.ValidatorNotFound {
		fmt.Printf("Validator %s not found\n", validatorPubKey.String())
		return nil
	}
	if canWithdrawClRewardsResponse.ValidatorWithdrawn {
		fmt.Printf("Validator %s has withdrawn all the staked funds\n", validatorPubKey.String())
		return nil
	}

	// Prompt for confirmation
	if !(c.Bool("yes") || cliutils.Confirm(fmt.Sprintf(
		"Are you sure you want to withdraw CL rewards for validator %s? (y/n)", validatorPubKey))) {
		fmt.Println("Cancelled.")
		return nil
	}

	res, err := staderClient.WithdrawClRewards(validatorPubKey)
	if err != nil {
		return err
	}

	fmt.Printf("Withdrawing %.6f CL Rewards to Operator Reward Address: %s\n\n", math.RoundDown(eth.WeiToEth(res.ClRewardsAmount), 6), res.OperatorRewardAddress)
	cliutils.PrintTransactionHash(staderClient, res.TxHash)
	if _, err = staderClient.WaitForTransaction(res.TxHash); err != nil {
		return err
	}

	// Log & return
	fmt.Printf("Successfully Withdrawn %.6f CL Rewards to Operator Reward Address: %s\n\n", math.RoundDown(eth.WeiToEth(res.ClRewardsAmount), 6), res.OperatorRewardAddress)

	return nil
}
