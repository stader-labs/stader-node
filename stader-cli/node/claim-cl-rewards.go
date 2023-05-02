package node

import (
	"fmt"
	"github.com/stader-labs/stader-node/shared/services/gas"

	"github.com/stader-labs/stader-node/shared/services/stader"
	cliutils "github.com/stader-labs/stader-node/shared/utils/cli"
	"github.com/stader-labs/stader-node/shared/utils/math"
	"github.com/stader-labs/stader-node/stader-lib/types"
	"github.com/stader-labs/stader-node/stader-lib/utils/eth"
	"github.com/urfave/cli"
)

func ClaimClRewards(c *cli.Context, validatorPubKey types.ValidatorPubkey) error {
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

	canClaimClRewardsResponse, err := staderClient.CanClaimClRewards(validatorPubKey)
	if err != nil {
		return err
	}
	if canClaimClRewardsResponse.NoClRewards {
		fmt.Printf("No CL rewards to withdraw for validator %s\n", validatorPubKey.String())
		return nil
	}
	if canClaimClRewardsResponse.TooManyClRewards {
		fmt.Printf("Too many CL rewards to withdraw for validator %s.\n", validatorPubKey.String())
		fmt.Printf("If you have exited the validator, Please wait for Stader Oracles to settle your funds!\n")
		fmt.Printf("If you have not exited the validator, Please reach out to the Stader Team on discord!\n")
		return nil
	}
	if canClaimClRewardsResponse.ValidatorNotFound {
		fmt.Printf("Validator %s not found\n", validatorPubKey.String())
		return nil
	}
	if canClaimClRewardsResponse.ValidatorWithdrawn {
		fmt.Printf("Validator %s has withdrawn all the staked funds\n", validatorPubKey.String())
		return nil
	}
	if canClaimClRewardsResponse.VaultAlreadySettled {
		fmt.Printf("Vault for validator %s has already been settled\n", validatorPubKey.String())
		return nil
	}

	err = gas.AssignMaxFeeAndLimit(canClaimClRewardsResponse.GasInfo, staderClient, c.Bool("yes"))
	if err != nil {
		return err
	}

	// Prompt for confirmation
	if !(c.Bool("yes") || cliutils.Confirm(fmt.Sprintf(
		"Are you sure you want to claim CL rewards for validator %s?", validatorPubKey))) {
		fmt.Println("Cancelled.")
		return nil
	}

	res, err := staderClient.ClaimClRewards(validatorPubKey)
	if err != nil {
		return err
	}

	fmt.Printf("Claiming %.6f CL Rewards to Operator Reward Address: %s\n\n", math.RoundDown(eth.WeiToEth(res.ClRewardsAmount), 6), res.OperatorRewardAddress)
	cliutils.PrintTransactionHash(staderClient, res.TxHash)
	if _, err = staderClient.WaitForTransaction(res.TxHash); err != nil {
		return err
	}

	// Log & return
	fmt.Printf("Successfully Claimed %.6f CL Rewards to Operator Reward Address: %s\n\n", math.RoundDown(eth.WeiToEth(res.ClRewardsAmount), 6), res.OperatorRewardAddress)

	return nil
}
