package node

import (
	"fmt"
	"github.com/stader-labs/stader-node/shared/services/gas"
	"github.com/stader-labs/stader-node/shared/services/stader"
	cliutils "github.com/stader-labs/stader-node/shared/utils/cli"
	"github.com/urfave/cli"
)

func updateOperatorName(c *cli.Context, operatorName string) error {

	staderClient, err := stader.NewClientFromCtx(c)
	if err != nil {
		return err
	}
	defer staderClient.Close()

	// check if we can update the el
	res, err := staderClient.CanUpdateOperatorName(operatorName)
	if err != nil {
		return err
	}
	if res.OperatorNotActive {
		fmt.Println("Operator not active")
		return nil
	}
	if res.OperatorNameTooLong {
		fmt.Println("Operator name too long")
		return nil
	}
	if res.NothingToUpdate {
		fmt.Println("Nothing to update")
		return nil
	}

	err = gas.AssignMaxFeeAndLimit(res.GasInfo, staderClient, c.Bool("yes"))
	if err != nil {
		return err
	}

	if !(c.Bool("yes") || cliutils.Confirm(fmt.Sprintf(
		"Are you sure you want to update your operator name?"))) {
		fmt.Println("Cancelled.")
		return nil
	}

	// update the socializing pool el
	response, err := staderClient.UpdateOperatorName(operatorName)
	if err != nil {
		return err
	}

	fmt.Println("Updating operator name...")

	cliutils.PrintTransactionHash(staderClient, response.TxHash)
	_, err = staderClient.WaitForTransaction(response.TxHash)
	if err != nil {
		return err
	}

	fmt.Println("Operator name updated!")

	return nil
}
