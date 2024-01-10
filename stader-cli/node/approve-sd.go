package node

import (
	"fmt"
	"math/big"
	"strconv"

	"github.com/ethereum/go-ethereum/common"
	"github.com/urfave/cli"

	"github.com/stader-labs/stader-node/shared/services/gas"
	"github.com/stader-labs/stader-node/shared/services/stader"
	cliutils "github.com/stader-labs/stader-node/shared/utils/cli"
	"github.com/stader-labs/stader-node/stader-lib/utils/eth"
)

func nodeApproveDepositSd(c *cli.Context, amountInString string) error {
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

	amount, err := strconv.ParseFloat(amountInString, 64)
	if err != nil {
		return err
	}

	amountWei := eth.EthToWei(amount)

	contracts, err := staderClient.GetContractsInfo()
	if err != nil {
		return err
	}

	autoConfirm := c.Bool("yes")

	nonce := c.GlobalUint64("nonce")

	err = nodeApproveSdWithAmountAndAddress(staderClient, amountWei, contracts.SdCollateralContract, autoConfirm, nonce)
	if err != nil {
		return err
	}

	return nil
}

func nodeApproveUtilitySd(c *cli.Context, amountInString string) error {
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

	amount, err := strconv.ParseFloat(amountInString, 64)
	if err != nil {
		return err
	}

	amountWei := eth.EthToWei(amount)

	contracts, err := staderClient.GetContractsInfo()
	if err != nil {
		return err
	}

	autoConfirm := c.Bool("yes")

	nonce := c.GlobalUint64("nonce")

	err = nodeApproveSdWithAmountAndAddress(staderClient, amountWei, contracts.SdUtilityContract, autoConfirm, nonce)
	if err != nil {
		return err
	}

	return nil
}

func nodeApproveSdWithAmountAndAddress(staderClient *stader.Client, amountWei *big.Int, address common.Address, autoConfirm bool, nonce uint64) error {
	// If a custom nonce is set, print the multi-transaction warning
	if nonce != 0 {
		cliutils.PrintMultiTransactionNonceWarning()
	}

	// Get approval gas
	approvalGas, err := staderClient.NodeSdApprovalGas(amountWei, address)
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

	response, err := staderClient.NodeSdApprove(amountWei, address)
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

// Calculate max uint256 value
func maxApprovalAmount() *big.Int {
	maxApproval := big.NewInt(2)
	maxApproval = maxApproval.Exp(maxApproval, big.NewInt(256), nil)
	maxApproval = maxApproval.Sub(maxApproval, big.NewInt(1))

	return maxApproval
}
