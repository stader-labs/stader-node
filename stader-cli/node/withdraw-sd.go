package node

import (
	"fmt"
	"math/big"
	"strconv"

	"github.com/stader-labs/stader-node/shared/services/gas"
	"github.com/stader-labs/stader-node/shared/services/stader"
	cliutils "github.com/stader-labs/stader-node/shared/utils/cli"
	"github.com/stader-labs/stader-node/stader-lib/utils/eth"
	"github.com/urfave/cli"
)

func WithdrawSd(c *cli.Context) error {
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

	amountInString := c.String("amount")
	amount, err := strconv.ParseFloat(amountInString, 64)
	if err != nil {
		return err
	}
	amountWei := eth.EthToWei(amount)

	canWithdrawSdResponse, err := staderClient.CanWithdrawSd(amountWei)
	if err != nil {
		return err
	}

	if canWithdrawSdResponse.InsufficientSdCollateral {
		fmt.Println("Not enough excess SD collateral to withdraw")
		return nil
	}

	if canWithdrawSdResponse.InsufficientWithdrawableSd {
		fmt.Println("Insufficient withdrawable SD!")
		return nil
	}

	sdStatusResponse, err := staderClient.GetSDStatus(big.NewInt(0))
	if err != nil {
		return err
	}

	hasUtilizePosition := sdStatusResponse.SDStatus.SdUtilizerLatestBalance.Cmp(big.NewInt(0)) != 0

	// Check if the withdraw request amount if greater than the Utilization Position
	withdrawExcessUtilization := amountWei.Cmp(sdStatusResponse.SDStatus.SdUtilizerLatestBalance) >= 0
	if withdrawExcessUtilization {
		if !cliutils.Confirm(fmt.Sprintf("You have an existing Utilization Position of %s. The excess SD collateral you are trying to withdraw will be used to repay the utilized SD and the remaining SD will be sent to your Reward Address.\n Do you wish to proceed?", eth.DisplayAmountInUnits(sdStatusResponse.SDStatus.SdUtilizerLatestBalance, "sd"))) {
			fmt.Println("Cancelled.")
			return nil
		}
	} else {
		if !cliutils.Confirm(fmt.Sprintf("You have an existing Utilization Position of %s. The excess SD collateral you are trying to withdraw will be used to repay the utilized SD.\n Do you wish to proceed?", eth.DisplayAmountInUnits(sdStatusResponse.SDStatus.SdUtilizerLatestBalance, "sd"))) {
			fmt.Println("Cancelled.")
			return nil
		}
	}

	// Assign max fees
	err = gas.AssignMaxFeeAndLimit(canWithdrawSdResponse.GasInfo, staderClient, c.Bool("yes"))
	if err != nil {
		return err
	}

	res, err := staderClient.WithdrawSd(amountWei)
	if err != nil {
		return err
	}

	fmt.Printf("Withdrawing %s from the collateral contract.\n", eth.DisplayAmountInUnits(amountWei, "sd"))
	cliutils.PrintTransactionHash(staderClient, res.TxHash)
	if _, err = staderClient.WaitForTransaction(res.TxHash); err != nil {
		return err
	}

	remainUtilize := new(big.Int).Sub(sdStatusResponse.SDStatus.SdUtilizerLatestBalance, amountWei)

	// remain collateral is (SdUtilizerLatestBalance + SdCollateralCurrentAmount) - amountWei
	remainCollateral := new(big.Int).Add(sdStatusResponse.SDStatus.SdCollateralCurrentAmount, remainUtilize)

	// Log & return

	// withdraw request amount lesser than the Utilization Position

	if !hasUtilizePosition {
		fmt.Printf("Successfully withdrawn %s Collateral. \n", eth.DisplayAmountInUnits(amountWei, "sd"))

		return nil
	}

	if withdrawExcessUtilization {
		// withdraw request amount greater than the Utilization Position
		fmt.Printf("Repayment of %s successful using the excess SD Collateral.\n", eth.DisplayAmountInUnits(amountWei, "sd"))
		fmt.Printf("The remaining %s has been sent to your Operator Reward Address\n", eth.DisplayAmountInUnits(new(big.Int).Abs(remainUtilize), "sd"))
	} else {
		fmt.Printf("Repayment of %s successful using the excess SD Collateral.. \n", eth.DisplayAmountInUnits(amountWei, "sd"))
		fmt.Printf("Current Utilization Position: %s\nCurrent SD collateral:  %s\n", eth.DisplayAmountInUnits(remainUtilize, "sd"), eth.DisplayAmountInUnits(remainCollateral, "sd"))
	}

	return nil
}
