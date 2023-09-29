package node

import (
	"fmt"

	"github.com/ethereum/go-ethereum/common"
	"github.com/stader-labs/stader-node/shared/services/gas"
	"github.com/stader-labs/stader-node/shared/services/stader"
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
		fmt.Printf("%sFor node security, only your existing Reward Address can \npropose a change. To propose and confirm a Reward address \nchange, please use the PermissionlessNodeRegistry Smart Contract: \n%s %s\n\n", colorLightBlue, infoResponse.PermissionlessNodeRegistry, colorReset)

		promptNote()
		return nil
	}

	err = gas.AssignMaxFeeAndLimit(res.GasInfo, staderClient, c.Bool("yes"))
	if err != nil {
		return err
	}

	if !(c.Bool("yes") || cliutils.Confirm(fmt.Sprintf(
		"\n%sThis action will change your Reward address. Once it's \nchanged, all future SD and ETH reward will be sent to the new address. \n\nYour new Reward address will initially be in a \"Confirmation pending\" state. \nTo confirm the change, please use \nthe PermissionlessNodeRegistry Smart Contract: %s \n\nStader will continue to send reward to your existing address until you \nconfirm the change using your new Reward Address. To complete \nthe confirmation process, make sure you have a web3-compatible wallet like \nMetaMask. Connect your wallet with the Smart Contract using \nyour new Reward address and confirm the Reward address change \nusing the ConfirmRewardAddressChange function. \n\n%sDo you wish to processed with the Reward address change?%s", colorLightBlue, infoResponse.PermissionlessNodeRegistry, colorGreen, colorReset))) {
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

	fmt.Printf("%sYour reward address changed in pending state. Please \nconfirm the request using this contract: %s and the reward address.%s", infoResponse.PermissionlessNodeRegistry, colorLightBlue, colorReset)

	return nil
}

func promptNote() {
	fmt.Printf("%sFollow these steps for your Reward address change: \nStep1: Propose the Reward address change by connecting your wallet \nto the Smart Contract using your EXISTING Reward address and the ProposeRewardAddress function. \nStep2: Confirm the Reward address change by connecting your wallet \nto the Smart Contract using your NEW Reward address and the ConfirmRewardAddressChange function. \n\n%sNote: Stader will continue to send reward to your existing Reward \naddress until you confirm the change using your new Reward Address%s", colorGreen, colorRed, colorReset)
}
