package node

import (
	"fmt"
	"github.com/stader-labs/stader-node/shared/services/stader"
	cliutils "github.com/stader-labs/stader-node/shared/utils/cli"
	"github.com/urfave/cli"
)

func ClaimSpRewards(c *cli.Context) error {
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

	// If a custom nonce is set, print the multi-transaction warning
	if c.GlobalUint64("nonce") != 0 {
		cliutils.PrintMultiTransactionNonceWarning()
	}

	// prompt user to select the cycles to claim from
	canClaimSpRewards, err := staderClient.CanClaimSpRewards()
	if err != nil {
		return err
	}
	if canClaimSpRewards.OperatorNotRegistered {
		fmt.Println("Operator not registered!")
		return nil
	}
	if len(canClaimSpRewards.UnclaimedCycles) == 0 {
		fmt.Println("No cycles to claim!")
		return nil
	}

	if len(canClaimSpRewards.CyclesToDownload) > 0 && !c.Bool("download-merkles-proofs") {
		fmt.Println("You have unclaimed rewards from cycles that you have not downloaded the merkle proofs for. Please download the merkle proofs for these cycles before claiming your rewards.")
		return nil
	}

	fmt.Println("Downloading the merkle proofs for the cycles you have unclaimed rewards for...")
	if len(canClaimSpRewards.CyclesToDownload) > 0 {
		_, err := staderClient.DownloadSpMerkleProofs()
		if err != nil {
			return err
		}
	}
	fmt.Println("Merkle proofs downloaded!")

	//fmt.Println("Following are the unclaimed cycles, Please enter in a comma seperated string the cycles you want to claim rewards for:")
	//for i, cycle := range canClaimSpRewards.UnclaimedCycles {
	//	fmt.Printf("%d) %d\n", i, cycle.Int64())
	//}
	//
	//cyclesToClaim := []int64{}
	//cycleSelection := cliutils.Prompt("Which cycles would you like to claim? Use a comma separated list (such as '1,2,3') or leave it blank to claim all cycles at once.", "^$|^\\d+(,\\d+)*$", "Invalid cycle selection")
	//if cycleSelection == "" {
	//	for _, cycle := range canClaimSpRewards.UnclaimedCycles {
	//		cyclesToClaim = append(cyclesToClaim, cycle.Int64())
	//	}
	//} else {
	//	elements := strings.Split(cycleSelection, ",")
	//	for _, element := range elements {
	//
	//	}
	//
	//}

	fmt.Printf("Claiming rewards for cycles %v\n", canClaimSpRewards.UnclaimedCycles)
	res, err := staderClient.ClaimSpRewards(canClaimSpRewards.UnclaimedCycles)
	if err != nil {
		return err
	}

	cliutils.PrintTransactionHash(staderClient, res.TxHash)
	_, err = staderClient.WaitForTransaction(res.TxHash)
	if err != nil {
		return err
	}
	fmt.Println("Rewards claimed! You can check your balance with the node status command.")

	return nil
}
