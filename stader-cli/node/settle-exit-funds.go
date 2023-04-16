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

func SettleExitFunds(c *cli.Context, validatorPubKey types.ValidatorPubkey) error {
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

	canSettleExitFunds, err := staderClient.CanSettleExitFunds(validatorPubKey)
	if err != nil {
		return err
	}
	if canSettleExitFunds.OperatorNotRegistered {
		fmt.Printf("Operator %s not registered\n", validatorPubKey.String())
		return nil
	}
	if canSettleExitFunds.ValidatorNotRegistered {
		fmt.Printf("Validator %s not registered\n", validatorPubKey.String())
		return nil
	}
	if canSettleExitFunds.ValidatorNotWithdrawn {
		fmt.Printf("Validator %s funds have not been withdrawn yet\n", validatorPubKey.String())
		return nil
	}
	if canSettleExitFunds.NoEthToWithdraw {
		fmt.Printf("No eth to withdraw for validator %s\n", validatorPubKey.String())
		return nil
	}
	if canSettleExitFunds.VaultAlreadySettled {
		fmt.Printf("Vault for validator %s already settled\n", validatorPubKey.String())
		return nil
	}

	// Assign max fees
	err = gas.AssignMaxFeeAndLimit(canSettleExitFunds.GasInfo, staderClient, c.Bool("yes"))
	if err != nil {
		return err
	}

	// Prompt for confirmation
	if !(c.Bool("yes") || cliutils.Confirm(fmt.Sprintf(
		"Are you sure you want to withdraw exited funds for validator %s?", validatorPubKey))) {
		fmt.Println("Cancelled.")
		return nil
	}

	res, err := staderClient.SettleExitFunds(validatorPubKey)
	if err != nil {
		return err
	}

	fmt.Printf("Withdrawing %.6f Exited Funds to Operator Reward Address: %s\n\n", math.RoundDown(eth.WeiToEth(res.ExitAmount), 6), res.OperatorRewardAddress)
	cliutils.PrintTransactionHash(staderClient, res.TxHash)
	if _, err = staderClient.WaitForTransaction(res.TxHash); err != nil {
		return err
	}

	// Log & return
	fmt.Printf("Successfully Withdrawn %.6f Exited Funds to Operator Reward Address: %s\n\n", math.RoundDown(eth.WeiToEth(res.ExitAmount), 6), res.OperatorRewardAddress)

	return nil
}
