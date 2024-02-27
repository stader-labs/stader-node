package node

import (
	"fmt"
	"math/big"

	"github.com/urfave/cli"

	"github.com/stader-labs/stader-node/shared/services/gas"
	"github.com/stader-labs/stader-node/shared/services/stader"
	"github.com/stader-labs/stader-node/shared/types/api"
	cliutils "github.com/stader-labs/stader-node/shared/utils/cli"
	"github.com/stader-labs/stader-node/shared/utils/log"
	"github.com/stader-labs/stader-node/stader-lib/utils/eth"
	"github.com/stader-labs/stader-node/stader-lib/utils/sd"
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

	// Check allowance
	allowance, err := staderClient.GetNodeSdAllowance(contracts.SdUtilityContract)
	if err != nil {
		return err
	}

	ops := []string{"Repay full amount", "Repay a custom amount"}
	i, _ := cliutils.Select("Please choose one of the following options for repayment. Enter 1 or 2:", ops)

	fullRepay := false

	var amountWei *big.Int

	switch i {
	case 0:
		// 1. If user had enough to repay
		fullRepay = true
		amountWei = sdStatus.SdUtilizerLatestBalance

	case 1:
		amountWei, err = PromptChooseRepayAmount(sdStatus)
		if err != nil {
			return err
		}

	default:
		return nil
	}

	if amountWei.Cmp(sdStatus.SdBalance) > 0 {
		fmt.Printf("You don't have sufficient SD in your Operator Address for the repayment. Please deposit SD into your Operator Address and try again.\n")
		return nil
	}

	if allowance.Allowance.Cmp(sdStatus.SdUtilizerLatestBalance) < 0 {
		fmt.Println("Before repaying the SD, you must first give the utility contract approval to interact with your SD.")

		maxApproval := maxUint256()

		err = nodeApproveUtilitySd(c, maxApproval.String())
		if err != nil {
			return err
		}
	}

	// Prompt for confirmation
	if !cliutils.Confirm(fmt.Sprintf("Are you sure you want to repay %s from your Operator Address and close your Utilization Position? ", eth.DisplayAmountInUnits(amountWei, "sd"))) {
		fmt.Printf("Cancelled\n")
		return nil
	}

	if fullRepay {
		amountWei = maxUint256()
	}

	canRepaySdResponse, err := staderClient.CanRepaySd(amountWei)
	if err != nil {
		return err
	}

	err = gas.AssignMaxFeeAndLimit(canRepaySdResponse.GasInfo, staderClient, c.Bool("yes"))
	if err != nil {
		return err
	}

	res, err := staderClient.RepaySd(amountWei)
	if err != nil {
		return err
	}

	cliutils.PrintTransactionHash(staderClient, res.TxHash)

	if _, err = staderClient.WaitForTransaction(res.TxHash); err != nil {
		return err
	}

	if amountWei.Cmp(maxUint256()) == 0 {
		fmt.Printf("Repayment of %s successful. Current Utilization Position: 0 SD.\n", eth.DisplayAmountInUnits(sdStatus.SdUtilizerLatestBalance, "sd"))
	} else {
		remainUtilize := new(big.Int).Sub(sdStatus.SdUtilizerLatestBalance, amountWei)
		fmt.Printf("Repayment of %s successful. Current Utilization Position: %s.\n", eth.DisplayAmountInUnits(amountWei, "sd"), eth.DisplayAmountInUnits(remainUtilize, "sd"))
	}

	return nil
}

func PromptChooseRepayAmount(sdStatus *api.SdStatusResponse) (*big.Int, error) {
	msg := fmt.Sprintf(`%sPlease enter the amount of SD you wish to repay. Your current Utilization Position is %s%s`, log.ColorYellow, eth.DisplayAmountInUnits(sdStatus.SdUtilizerLatestBalance, "sd"), log.ColorReset)

	errMsg := fmt.Sprintf("%sInvalid input, please specify a valid amount of SD you wish to repay. Your current Utilization Position is %s SD%s", log.ColorRed, eth.DisplayAmountInUnits(sdStatus.SdUtilizerLatestBalance, "sd"), log.ColorReset)

	utilityAmount, err := sd.PromptChooseSDWithMaxMin(msg, errMsg, big.NewInt(0), sdStatus.SdUtilizerLatestBalance)
	if err != nil {
		return nil, err
	}

	if sd.WeiAlmostEqual(utilityAmount, sdStatus.SdUtilizerLatestBalance) {
		utilityAmount = sdStatus.SdUtilizerLatestBalance
	}

	return utilityAmount, nil
}
