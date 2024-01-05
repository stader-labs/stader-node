package node

import (
	"fmt"
	"math/big"
	"strconv"

	"github.com/stader-labs/stader-node/stader-lib/utils/eth"
	"github.com/urfave/cli"

	"github.com/stader-labs/stader-node/shared/services/gas"
	"github.com/stader-labs/stader-node/shared/services/stader"
	cliutils "github.com/stader-labs/stader-node/shared/utils/cli"
)

func nodeDepositSd(c *cli.Context) error {

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

	nounce := c.GlobalUint64("nonce")
	// If a custom nonce is set, print the multi-transaction warning
	if nounce != 0 {
		cliutils.PrintMultiTransactionNonceWarning()
	}

	amountInString := c.String("amount")

	amount, err := strconv.ParseFloat(amountInString, 64)
	if err != nil {
		return err
	}

	autoConfirm := c.Bool("yes")

	amountWei := eth.EthToWei(amount)

	return DepositSdWithAmount(staderClient, amountWei, autoConfirm, nounce)
}

func DepositSdWithAmount(staderClient *stader.Client, amountWei *big.Int, autoConfirm bool, nonce uint64) error {
	contracts, err := staderClient.GetContractsInfo()
	if err != nil {
		return err
	}

	// Check allowance
	allowance, err := staderClient.GetNodeSdAllowance(contracts.SdCollateralContract)
	if err != nil {
		return err
	}

	if allowance.Allowance.Cmp(amountWei) < 0 {
		fmt.Println("Before depositing SD, you must first give the collateral contract approval to interact with your SD. Amount to approve: ", eth.DisplayAmountInUnits(amountWei, "sd"))
		err = nodeApproveSdWithAmountAndAddress(staderClient, amountWei, contracts.SdCollateralContract, autoConfirm, nonce)
		if err != nil {
			return err
		}
	}

	canDeposit, err := staderClient.CanNodeDepositSd(amountWei)
	if err != nil {
		return err
	}

	if canDeposit.InsufficientBalance {
		fmt.Println("You don't have sufficient SD in your Operator Address to add as collateral. Please deposit SD into your Operator Address and try again.")
		return nil
	}

	if canDeposit.CollateralContractPaused {
		fmt.Println("The collateral contract is paused.")
		return nil
	}

	// Assign max fees
	err = gas.AssignMaxFeeAndLimit(canDeposit.GasInfo, staderClient, autoConfirm)
	if err != nil {
		return err
	}

	// Prompt for confirmation
	if !(autoConfirm || cliutils.Confirm(fmt.Sprintf("Are you sure you want to deposit %s as collateral?", eth.DisplayAmountInUnits(amountWei, "sd")))) {
		fmt.Println("Cancelled.")
		return nil
	}

	depositSdResponse, err := staderClient.NodeDepositSd(amountWei)
	if err != nil {
		return err
	}

	fmt.Printf("Depositing SD...\n")
	cliutils.PrintTransactionHash(staderClient, depositSdResponse.DepositTxHash)

	if _, err = staderClient.WaitForTransaction(depositSdResponse.DepositTxHash); err != nil {
		return err
	}

	// Log & return
	fmt.Printf("Successfully deposited %s.\n", eth.DisplayAmountInUnits(amountWei, "sd"))

	return nil
}
