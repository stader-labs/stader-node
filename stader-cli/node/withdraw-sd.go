package node

import (
	"fmt"
	"math/big"
	"strconv"

	"github.com/stader-labs/stader-node/shared/services/gas"
	"github.com/stader-labs/stader-node/shared/services/stader"
	cliutils "github.com/stader-labs/stader-node/shared/utils/cli"
	"github.com/stader-labs/stader-node/shared/utils/math"
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
	if canWithdrawSdResponse.InsufficientWithdrawableSd {
		fmt.Println("Insufficient withdrawable SD!")
		return nil
	}
	if canWithdrawSdResponse.InsufficientSdCollateral {
		fmt.Println("SD collateral less than 200%")
		return nil
	}

	sdStatusResponse, err := staderClient.GetSDStatus(big.NewInt(0))
	if err != nil {
		return err
	}

	if sdStatusResponse.SDStatus.SdUtilizerLatestBalance.Cmp(amountWei) > 0 {
		confirm := cliutils.Confirm(fmt.Sprintf("You have an existing Utilization Position of %.6f SD. The excess SD collateral you are trying to withdraw will be used to repay the utilized SD.\n Do you wish to proceed?", math.RoundDown(eth.WeiToEth(sdStatusResponse.SDStatus.SdUtilizerLatestBalance), 6)))
		if !confirm {
			fmt.Println("Cancelled.")
			return nil
		}
	} else {
		confirm := cliutils.Confirm(fmt.Sprintf("You have an existing Utilization Position of %.6f SD. The excess SD collateral you are trying to withdraw will be used to repay the utilized SD and the remaining SD will be sent to your Reward Address.\n Do you wish to proceed?", math.RoundDown(eth.WeiToEth(sdStatusResponse.SDStatus.SdUtilizerLatestBalance), 6)))
		if !confirm {
			fmt.Println("Cancelled.")
			return nil
		}
	}

	// Assign max fees
	err = gas.AssignMaxFeeAndLimit(canWithdrawSdResponse.GasInfo, staderClient, c.Bool("yes"))
	if err != nil {
		return err
	}

	// Prompt for confirmation
	if !(c.Bool("yes") || cliutils.Confirm(fmt.Sprintf(
		"Are you sure you want to withdraw %.6f SD from the collateral contract?", math.RoundDown(eth.WeiToEth(amountWei), 6)))) {
		fmt.Println("Cancelled.")
		return nil
	}

	res, err := staderClient.WithdrawSd(amountWei)
	if err != nil {
		return err
	}

	fmt.Printf("Withdrawing %s SD from the collateral contract.\n", amountInString)
	cliutils.PrintTransactionHash(staderClient, res.TxHash)
	if _, err = staderClient.WaitForTransaction(res.TxHash); err != nil {
		return err
	}

	remainUtilize := new(big.Int).Sub(sdStatusResponse.SDStatus.SdUtilizerLatestBalance, amountWei)

	// remain collateral is (SdUtilizerLatestBalance + SdCollateralCurrentAmount) - amountWei
	remainCollateral := new(big.Int).Add(sdStatusResponse.SDStatus.SdCollateralCurrentAmount, remainUtilize)

	// Log & return

	// withdraw request amount lesser than the Utilization Position
	if remainUtilize.Cmp(big.NewInt(0)) > 0 {
		fmt.Printf("Successfully withdrawn %.6f SD Collateral. \n", math.RoundDown(eth.WeiToEth(amountWei), 6))
		fmt.Printf("Current Utilization Position: %.6f SD \nCurrent SD collateral:  %.6f SD\n", math.RoundDown(eth.WeiToEth(remainUtilize), 6), math.RoundDown(eth.WeiToEth(remainCollateral), 6))

		return nil
	}

	// withdraw request amount greater than the Utilization Position
	fmt.Printf("Repayment of %.6f SD successful using the excess SD Collateral.\n", math.RoundDown(eth.WeiToEth(amountWei), 6))
	fmt.Printf("The remaining %.6f SD has been sent to your Operator Reward Address\n", math.RoundDown(eth.WeiToEth(new(big.Int).Abs(remainUtilize)), 6))

	return nil
}
