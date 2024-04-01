package node

import (
	"fmt"

	"github.com/ethereum/go-ethereum/common"
	"github.com/stader-labs/stader-node/shared/services/config"
	"github.com/stader-labs/stader-node/shared/services/gas"
	"github.com/stader-labs/stader-node/shared/services/stader"
	cfTypes "github.com/stader-labs/stader-node/shared/types/config"
	cliutils "github.com/stader-labs/stader-node/shared/utils/cli"
	"github.com/urfave/cli"
)

const (
	colorReset     string = "\033[0m"
	colorBold      string = "\033[1m"
	colorRed       string = "\033[31m"
	colorYellow    string = "\033[33m"
	colorGreen     string = "\033[32m"
	colorLightBlue string = "\033[36m"
	clearLine      string = "\033[2K"
)

func SetRewardAddress(c *cli.Context, operatorRewardAddress common.Address) error {
	staderClient, err := stader.NewClientFromCtx(c)
	if err != nil {
		return err
	}
	defer staderClient.Close()

	cfg, _, err := staderClient.LoadConfig()
	if err != nil {
		return err
	}
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
		promptHowToChangeReward(cfg, infoResponse.PermissionlessNodeRegistry)
		return nil
	}

	err = gas.AssignMaxFeeAndLimit(res.GasInfo, staderClient, c.Bool("yes"))
	if err != nil {
		return err
	}

	confirmMessage := `
This action will change your Reward Address. Once it's changed, all future SD and ETH rewards will be sent to the New Reward Address.

After you propose the change, your New Reward Address will initially be in a
'Confirmation pending' state until you confirm the change using your New RewardAddress on the
PermissionlessNodeRegistry Smart Contract. Please make sure that your New Reward Address
is linked to a web3-compatible wallet, such as MetaMask, to connect with the Smart Contract

Do you wish to proceed with the Reward Address change?`

	if !(c.Bool("yes") || cliutils.Confirm(fmt.Sprintf(
		"\n%s %s %s", colorLightBlue, confirmMessage, colorReset))) {
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

	promptSuccessChangedRewardAndNextStep(cfg, infoResponse.PermissionlessNodeRegistry)
	return nil
}

func promptSuccessChangedRewardAndNextStep(cfg *config.StaderConfig, contractAddr common.Address) {
	switch cfg.StaderNode.Network.Value.(cfTypes.Network) {
	case cfTypes.Network_Mainnet:
		msg := `
You have successfully raised a request to change your Reward Address.

To confirm the Reward Address change please follow these steps:
Step 1: Visit the PermissionlessNodeRegistry Smart Contract: https://etherscan.io/address/%s#writeProxyContract#F3
Step 2: Confirm the Reward Address change by connecting your New Reward Address wallet with the Smart Contract and execute the "ConfirmRewardAddressChange" function.

Please refer to the Reward Address change guide here - https://www.staderlabs.com/docs-v1/Ethereum/Node-operator/permissionless-node-operator/Update-reward-address/

Note: Stader will continue to send rewards to your existing Reward Address until you confirm the change using your new Reward Address.
`
		msg = fmt.Sprintf(msg, contractAddr.String())
		fmt.Printf("%s %s %s\n", colorLightBlue, msg, colorReset)
	case cfTypes.Network_Holesky:
		msg := `
You have successfully raised a request to change your Reward Address.

To confirm the Reward Address change please follow these steps:
Step 1: Visit the PermissionlessNodeRegistry Smart Contract: https://goerli.etherscan.io/address/%s#writeProxyContract#F3
Step 2: Confirm the Reward Address change by connecting your New Reward Address wallet with the Smart Contract and execute the "ConfirmRewardAddressChange" function.

Please refer to the Reward Address change guide here - https://www.staderlabs.com/docs-v1/Ethereum/Node-operator/permissionless-node-operator/Update-reward-address/

Note: Stader will continue to send rewards to your existing Reward Address until you confirm the change using your new Reward Address.`
		msg = fmt.Sprintf(msg, contractAddr.String())
		fmt.Printf("%s %s %s\n", colorLightBlue, msg, colorReset)
	default:
		fmt.Println("Unsupported network")
	}
}

func promptHowToChangeReward(cfg *config.StaderConfig, contractAddr common.Address) {
	network := cfg.StaderNode.Network.Value.(cfTypes.Network)
	switch network {
	case cfTypes.Network_Mainnet:
		msg := `
For node security, only your existing Reward Address can propose a change. To propose and confirm a Reward Address change, please use the PermissionlessNodeRegistry Smart Contract: https://etherscan.io/address/%s#writeProxyContract#F10

Follow these steps for your Reward address change:
Step 1: Propose the Reward Address change by connecting your Existing Reward Address wallet with the Smart Contract and execute the "ProposeRewardAddress" function.
Step 2: Confirm the Reward Address change by connecting your New Reward Address wallet with the Smart Contract and execute the "ConfirmRewardAddressChange" function

Please refer to the Reward Address change guide here - https://www.staderlabs.com/docs-v1/Ethereum/Node-operator/permissionless-node-operator/Update-reward-address/

Note: Stader will continue to send rewards to your existing Reward Address until you confirm the change using your New Reward Address.
`
		msg = fmt.Sprintf(msg, contractAddr.String())
		fmt.Printf("%s %s %s\n\n", colorLightBlue, msg, colorReset)
	case cfTypes.Network_Holesky:
		msg := `
For node security, only your existing Reward Address can propose a change. To propose and confirm a Reward Address change, please use the PermissionlessNodeRegistry Smart Contract: https://goerli.etherscan.io/address/%s#writeProxyContract#F10

Follow these steps for your Reward Address change:
Step 1: Propose the Reward Address change by connecting your Existing Reward Address wallet with the Smart Contract and execute the "ProposeRewardAddress" function.
Step 2: Confirm the Reward Address change by connecting your New Reward Address wallet with the Smart Contract and execute the "ConfirmRewardAddressChange" function

Please refer to the Reward Address change guide here - https://www.staderlabs.com/docs-v1/Ethereum/Node-operator/permissionless-node-operator/Update-reward-address/

Note: Stader will continue to send rewards to your existing Reward Address until you confirm the change using your New Reward Address.
`
		msg = fmt.Sprintf(msg, contractAddr.String())
		fmt.Printf("%s %s %s\n\n", colorLightBlue, msg, colorReset)
	default:
		fmt.Println("Unsupported network")
	}
}
