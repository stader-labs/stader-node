package node

import (
	"fmt"
	"math/big"
	"strconv"

	"github.com/urfave/cli"

	"github.com/stader-labs/stader-node/shared/services/gas"
	"github.com/stader-labs/stader-node/shared/services/stader"
	cliutils "github.com/stader-labs/stader-node/shared/utils/cli"
	"github.com/stader-labs/stader-node/shared/utils/math"
	"github.com/stader-labs/stader-node/stader-lib/utils/eth"
)

func repaySD(c *cli.Context) error {
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

	amountInString := c.String("amount")

	amount, err := strconv.ParseFloat(amountInString, 64)
	if err != nil {
		return err
	}

	amountWei := eth.EthToWei(amount)

	contracts, err := staderClient.GetContractsInfo()
	if err != nil {
		return err
	}

	sdStatusResponse, err := staderClient.GetSDStatus(big.NewInt(0))
	if err != nil {
		return err
	}

	sdStatus := sdStatusResponse.SDStatus
	if sdStatus.SdUtilizerLatestBalance.Cmp(big.NewInt(0)) == 0 {
		fmt.Println("You do not have an existing Utilization Position.")
		return nil
	}

	// 1. Check if repay more than need
	if amountWei.Cmp(sdStatus.SdUtilizerLatestBalance) > 0 {
		fmt.Printf("Repayment amount greater than the Utilization position. Your current Utilization Position is %0.6f \n", eth.WeiToEth(sdStatus.SdUtilizerLatestBalance))
		return nil
	}

	// 2. If user had enough to repay
	if amountWei.Cmp(sdStatus.SdBalance) > 0 {
		fmt.Printf("You don't have sufficient SD in your Operator Address for the repayment. Please deposit SD into your Operator Address and try again.\n")
		return nil
	}

	// Check allowance
	allowance, err := staderClient.GetNodeSdAllowance(contracts.SdUtilityContract)
	if err != nil {
		return err
	}

	if allowance.Allowance.Cmp(amountWei) < 0 {
		fmt.Println("Before repaying the SD, you must first give the utility contract approval to interact with your SD. Amount to approve: ", eth.WeiToEth(amountWei))
		err = nodeApproveUtilitySd(c, amountInString)
		if err != nil {
			return err
		}
	}

	canRepaySdResponse, err := staderClient.CanRepaySd(amountWei)
	if err != nil {
		return err
	}

	err = gas.AssignMaxFeeAndLimit(canRepaySdResponse.GasInfo, staderClient, c.Bool("yes"))
	if err != nil {
		return err
	}

	// Prompt for confirmation
	if !(c.Bool("yes") || cliutils.Confirm(fmt.Sprintf(
		"Are you sure you want to repay %0.6f SD from your Operator Address and reduce or close your Utilization Position?", eth.WeiToEth(amountWei)))) {
		fmt.Println("Cancelled.")
		return nil
	}

	res, err := staderClient.RepaySd(amountWei)
	if err != nil {
		return err
	}

	cliutils.PrintTransactionHash(staderClient, res.TxHash)

	if _, err = staderClient.WaitForTransaction(res.TxHash); err != nil {
		return err
	}

	remainUtilize := new(big.Int).Sub(sdStatus.SdUtilizerLatestBalance, amountWei)
	fmt.Printf("Repayment of %.6f SD successful. Current Utilization Position: %.6f SD.\n", math.RoundDown(eth.WeiToEth(amountWei), 6), math.RoundDown(eth.WeiToEth(remainUtilize), 6))

	return nil
}
