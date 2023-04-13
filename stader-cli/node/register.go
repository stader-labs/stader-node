package node

import (
	"fmt"

	"github.com/ethereum/go-ethereum/common"
	"github.com/urfave/cli"

	"github.com/stader-labs/stader-node/shared/services/gas"
	"github.com/stader-labs/stader-node/shared/services/stader"
	cliutils "github.com/stader-labs/stader-node/shared/utils/cli"
)

func registerNode(c *cli.Context) error {

	staderClient, err := stader.NewClientFromCtx(c)
	if err != nil {
		return err
	}
	defer staderClient.Close()

	walletStatus, err := staderClient.WalletStatus()
	if err != nil {
		return err
	}

	// Check and assign the EC status
	err = cliutils.CheckClientStatus(staderClient)
	if err != nil {
		return err
	}

	operatorName := c.String("operator-name")
	operatorRewardAddressString := c.String("operator-reward-address")
	if operatorRewardAddressString == "" {
		operatorRewardAddressString = walletStatus.AccountAddress.String()
	}
	socializeEl := c.BoolT("socialize-el")

	// Check node can be registered
	canRegister, err := staderClient.CanRegisterNode(operatorName, common.HexToAddress(operatorRewardAddressString), socializeEl)
	if err != nil {
		return err
	}
	if canRegister.AlreadyRegistered {
		fmt.Println("The node is already registered with Stader.")
		return nil
	}
	if canRegister.RegistrationPaused {
		fmt.Println("Node registrations are currently disabled.")
		return nil
	}
	if canRegister.OperatorNameTooLong {
		fmt.Println("The operator name is too long.")
		return nil
	}
	if canRegister.OperatorRewardAddressZero {
		fmt.Println("The operator reward address cannot be zero.")
		return nil
	}

	// Assign max fees
	err = gas.AssignMaxFeeAndLimit(canRegister.GasInfo, staderClient, c.Bool("yes"))
	if err != nil {
		return err
	}

	// Prompt for confirmation
	if !(c.Bool("yes") || cliutils.Confirm("Are you sure you want to register this node?")) {
		fmt.Println("Cancelled.")
		return nil
	}

	// Register node
	response, err := staderClient.RegisterNode(operatorName, common.HexToAddress(operatorRewardAddressString), socializeEl)
	if err != nil {
		return err
	}

	fmt.Printf("Registering node...\n")
	cliutils.PrintTransactionHash(staderClient, response.TxHash)
	if _, err = staderClient.WaitForTransaction(response.TxHash); err != nil {
		return err
	}

	// Log & return
	fmt.Println("The node was successfully registered with Stader.")
	return nil

}
