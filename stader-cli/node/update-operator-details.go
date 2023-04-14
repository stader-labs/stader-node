package node

import (
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/stader-labs/stader-node/shared/services/gas"
	"github.com/stader-labs/stader-node/shared/services/stader"
	cliutils "github.com/stader-labs/stader-node/shared/utils/cli"
	"github.com/urfave/cli"
)

func updateOperatorDetails(c *cli.Context, operatorName string, operatorRewardAddress common.Address) error {

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

	// check if we can update the el
	res, err := staderClient.CanUpdateOperatorDetails(operatorName, operatorRewardAddress)
	if err != nil {
		return err
	}
	if res.OperatorNotRegistered {
		fmt.Println("Operator not registered")
		return nil
	}
	if res.OperatorNotActive {
		fmt.Println("Operator not active")
		return nil
	}
	if res.OperatorNameTooLong {
		fmt.Println("Operator name too long")
		return nil
	}
	if res.OperatorRewardAddressZero {
		fmt.Println("Operator reward address cannot be zero")
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
		"Are you sure you want to update your operator details?"))) {
		fmt.Println("Cancelled.")
		return nil
	}

	// update the socializing pool el
	response, err := staderClient.UpdateOperatorDetails(operatorName, operatorRewardAddress)
	if err != nil {
		return err
	}

	fmt.Println("Updating operator details...")

	cliutils.PrintTransactionHash(staderClient, response.TxHash)
	_, err = staderClient.WaitForTransaction(response.TxHash)
	if err != nil {
		return err
	}

	fmt.Println("Operator details updated!")

	return nil
}
