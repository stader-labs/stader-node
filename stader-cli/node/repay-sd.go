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

	// Check allowance
	allowance, err := staderClient.GetNodeSdAllowance(contracts.SdUtilityContract)
	if err != nil {
		return err
	}

	if allowance.Allowance.Cmp(amountWei) < 0 {
		fmt.Println("Before repay SD, you must first give the utility contract approval to interact with your SD. Amount to approve: ", eth.WeiToEth(amountWei))
		err = nodeApproveSd(c, contracts.SdUtilityContract.String(), amountInString)
		if err != nil {
			return err
		}
	}

	sdStatusResponse, err := staderClient.GetSDStatus(big.NewInt(0))
	if err != nil {
		return err
	}

	sdStatus := sdStatusResponse.SDStatus
	// 1. Check if repay more than need
	if amountWei.Cmp(sdStatus.SdUtilizerLatestBalance) > 0 {
		cliutils.PrintError(fmt.Sprintf("Repayment amount greater than the Utilization position: %s \n", sdStatus.SdUtilizerLatestBalance.String()))
		return nil
	}

	// 2. If user had enough to repay
	if amountWei.Cmp(sdStatus.SdBalance) > 0 {
		cliutils.PrintError(fmt.Sprintf("The node's SD balance is not enough SD, current SD available: %f \n", eth.WeiToEth(sdStatus.SdBalance)))
		return nil
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
		"Are you sure you want to repay %f SD?", amount))) {
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

	return nil
}
