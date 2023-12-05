package node

import (
	"fmt"
	"math/big"
	"strconv"

	"github.com/urfave/cli"

	"github.com/stader-labs/stader-node/shared/services/gas"
	"github.com/stader-labs/stader-node/shared/services/stader"
	cliutils "github.com/stader-labs/stader-node/shared/utils/cli"
	"github.com/stader-labs/stader-node/stader-lib/utils/eth"
)

func nodeApproveSd(c *cli.Context) error {
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

	// If a custom nonce is set, print the multi-transaction warning
	if c.GlobalUint64("nonce") != 0 {
		cliutils.PrintMultiTransactionNonceWarning()
	}

	amountInString := c.String("amount")

	amount, err := strconv.ParseFloat(amountInString, 64)
	if err != nil {
		return err
	}

	amountWei := eth.EthToWei(amount)

	autoConfirm := c.Bool("yes")

	nonce := c.GlobalUint64("nonce")

	err = nodeApproveSdWithAmount(staderClient, amountWei, autoConfirm, nonce)
	if err != nil {
		return err
	}

	return nil
}

func nodeApproveSdWithAmount(staderClient *stader.Client, amountWei *big.Int, autoConfirm bool, nonce uint64) error {
	// If a custom nonce is set, print the multi-transaction warning
	if nonce != 0 {
		cliutils.PrintMultiTransactionNonceWarning()
	}

	// Get approval gas
	approvalGas, err := staderClient.NodeDepositSdApprovalGas(amountWei)
	if err != nil {
		return err
	}
	// Assign max fees
	err = gas.AssignMaxFeeAndLimit(approvalGas.GasInfo, staderClient, autoConfirm)
	if err != nil {
		return err
	}

	// Prompt for confirmation
	if !(autoConfirm || cliutils.Confirm("Do you want to approve SD to be spent by the Collateral Contract?")) {
		fmt.Println("Cancelled.")
		return nil
	}

	response, err := staderClient.NodeDepositSdApprove(amountWei)
	if err != nil {
		return err
	}

	hash := response.ApproveTxHash

	fmt.Println("Approving SD..")
	cliutils.PrintTransactionHash(staderClient, hash)

	if _, err = staderClient.WaitForTransaction(hash); err != nil {
		return err
	}

	fmt.Println("Successfully approved SD.")

	// If a custom nonce is set, increment it for the next transaction
	if nonce != 0 {
		staderClient.IncrementCustomNonce()
	}

	return nil
}

func nodeApproveUtilitySd(c *cli.Context) error {
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

	// If a custom nonce is set, print the multi-transaction warning
	if c.GlobalUint64("nonce") != 0 {
		cliutils.PrintMultiTransactionNonceWarning()
	}

	amountInString := c.String("amount")

	amount, err := strconv.ParseFloat(amountInString, 64)
	if err != nil {
		return err
	}

	amountWei := eth.EthToWei(amount)

	autoConfirm := c.Bool("yes")

	nonce := c.GlobalUint64("nonce")

	if nonce != 0 {
		cliutils.PrintMultiTransactionNonceWarning()
	}

	// Get approval gas
	approvalGas, err := staderClient.SDPoolUtilitySdApproveGas(amountWei)
	if err != nil {
		return err
	}

	// Assign max fees
	err = gas.AssignMaxFeeAndLimit(approvalGas.GasInfo, staderClient, autoConfirm)
	if err != nil {
		return err
	}

	// Prompt for confirmation
	if !(autoConfirm || cliutils.Confirm("Do you want to approve SD to be spent by the Utility Contract?")) {
		fmt.Println("Cancelled.")
		return nil
	}

	response, err := staderClient.SDPoolUtilitySdApprove(amountWei)
	if err != nil {
		return err
	}

	hash := response.ApproveTxHash

	fmt.Println("Approving SD..")
	cliutils.PrintTransactionHash(staderClient, hash)

	if _, err = staderClient.WaitForTransaction(hash); err != nil {
		return err
	}

	fmt.Println("Successfully approved SD.")

	// If a custom nonce is set, increment it for the next transaction
	if nonce != 0 {
		staderClient.IncrementCustomNonce()
	}
	return nil
}
