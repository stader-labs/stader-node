package node

import (
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/rocket-pool/rocketpool-go/rocketpool"

	"github.com/urfave/cli"

	"github.com/stader-labs/stader-node/shared/services/gas"
	"github.com/stader-labs/stader-node/shared/services/stader"
	cliutils "github.com/stader-labs/stader-node/shared/utils/cli"
)

func registerNode(c *cli.Context) error {

	// Get RP client
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

	// Check node can be registered
	canRegister, err := staderClient.CanRegisterNode()
	if err != nil {
		return err
	}
	if !canRegister.CanRegister {
		fmt.Println("The node cannot be registered:")
		if canRegister.AlreadyRegistered {
			fmt.Println("The node is already registered with Stader.")
		}
		if canRegister.RegistrationDisabled {
			fmt.Println("Node registrations are currently disabled.")
		}
		return nil
	}

	// Assign max fees
	err = gas.AssignMaxFeeAndLimit(rocketpool.GasInfo{
		EstGasLimit:  10000000,
		SafeGasLimit: 25000000,
	}, staderClient, c.Bool("yes"))
	if err != nil {
		return err
	}

	// Prompt for confirmation
	if !(c.Bool("yes") || cliutils.Confirm("Are you sure you want to register this node?")) {
		fmt.Println("Cancelled.")
		return nil
	}

	operatorName := c.String("operator-name")
	operatorRewardAddress := c.String("operator-reward-address")
	socializeMev := c.Bool("socialize-mev")

	fmt.Printf("cli: Register-Node: operator reward address is %s\n\n", common.HexToAddress(operatorRewardAddress))

	// Register node
	response, err := staderClient.RegisterNode(operatorName, common.HexToAddress(operatorRewardAddress), socializeMev)
	if err != nil {
		return err
	}

	fmt.Printf("Registering node...\n")
	cliutils.PrintTransactionHash(staderClient, response.TxHash)
	if _, err = staderClient.WaitForTransaction(response.TxHash); err != nil {
		return err
	}

	// Log & return
	fmt.Println("The node was successfully registered with Rocket Pool.")
	return nil

}
