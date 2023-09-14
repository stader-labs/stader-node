package node

import (
	"fmt"

	"github.com/ethereum/go-ethereum/common"
	"github.com/stader-labs/stader-node/shared/services/gas"
	"github.com/stader-labs/stader-node/shared/services/stader"
	cliutils "github.com/stader-labs/stader-node/shared/utils/cli"
	"github.com/urfave/cli"
)

func SetRewardAddress(c *cli.Context, operatorRewardAddress common.Address) error {
	staderClient, err := stader.NewClientFromCtx(c)
	if err != nil {
		return err
	}
	defer staderClient.Close()

	// check if we can update the el
	res, err := staderClient.CanUpdateOperatorRewardAddress(operatorRewardAddress)
	if err != nil {
		return err
	}
	if res.OperatorNotActive {
		fmt.Println("Operator not active")
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
	if res.IsPermissionlessNodeRegistryPaused {
		fmt.Println("Permissionless Node Registry is paused.")
		return nil
	}

	infoResponse, err := staderClient.GetContractsInfo()
	if err != nil {
		return err
	}

	if res.OperatorAddressAndRewardNotTheSame {
		fmt.Printf("Please follow the given steps using this contract: %s to \nchange your reward address: \nStep1: Initiate the change using your existing reward address. \nStep2: Confirm the change using new reward address.\n", infoResponse.PermissionlessNodeRegistry)

		return nil
	}

	err = gas.AssignMaxFeeAndLimit(res.GasInfo, staderClient, c.Bool("yes"))
	if err != nil {
		return err
	}

	if !(c.Bool("yes") || cliutils.Confirm(fmt.Sprintf(
		"Are you sure you want to update your operator reward address?"))) {
		fmt.Println("Cancelled.")
		return nil
	}

	// update the socializing pool el
	response, err := staderClient.SetRewardAddress(operatorRewardAddress)
	if err != nil {
		return err
	}

	fmt.Println("Updating operator reward address...")

	cliutils.PrintTransactionHash(staderClient, response.TxHash)
	_, err = staderClient.WaitForTransaction(response.TxHash)
	if err != nil {
		return err
	}

	fmt.Printf("Your reward address changed in pending state. Please \nconfirm the request using this contract: %s and the reward address.", infoResponse.PermissionlessNodeRegistry)

	return nil
}
