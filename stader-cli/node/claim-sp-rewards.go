package node

import (
	"fmt"
	"github.com/stader-labs/stader-node/shared/services/gas"
	"github.com/stader-labs/stader-node/shared/services/stader"
	cliutils "github.com/stader-labs/stader-node/shared/utils/cli"
	"github.com/stader-labs/stader-node/shared/utils/math"
	"github.com/stader-labs/stader-node/stader-lib/utils/eth"
	"github.com/urfave/cli"
	"math/big"
	"strconv"
	"strings"
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

	fmt.Println("Downloading the merkle proofs for the cycles you may have not downloaded yet...")
	downloadRes, err := staderClient.DownloadSpMerkleProofs()
	if err != nil {
		return err
	}
	if len(downloadRes.DownloadedCycles) != 0 {
		fmt.Printf("Merkle proofs downloaded for cycles %v!\n", downloadRes.DownloadedCycles)
	} else {
		fmt.Println("No new merkle proofs downloaded!")
	}

	// prompt user to select the cycles to claim from
	canClaimSpRewards, err := staderClient.CanClaimSpRewards()
	if err != nil {
		return err
	}
	if canClaimSpRewards.SocializingPoolContractPaused {
		fmt.Println("The socializing pool contract is paused!")
		return nil
	}
	if len(canClaimSpRewards.UnclaimedCycles) == 0 {
		fmt.Println("You have no unclaimed cycles!")
		return nil
	}

	fmt.Printf("Getting the detailed cycles info...\n")
	detailedCyclesInfo, err := staderClient.GetDetailedCyclesInfo(canClaimSpRewards.UnclaimedCycles)
	if err != nil {
		return err
	}

	// this is post checking the cache.
	if len(detailedCyclesInfo.DetailedCyclesInfo) == 0 {
		fmt.Println("You have no unclaimed cycles!")
		return nil
	}

	cycleIndexes := []*big.Int{}
	for _, cycleInfo := range detailedCyclesInfo.DetailedCyclesInfo {
		cycleIndexes = append(cycleIndexes, big.NewInt(cycleInfo.MerkleProofInfo.Cycle))
	}

	fmt.Println("Following are the unclaimed cycles, Please enter in a comma seperated string the cycles you want to claim rewards for:\n")

	fmt.Printf("%-18s%-14.30s%-14.10s%-10s\n", "Cycle Number", "Cycle Date", "ETH Rewards", "SD Rewards")
	cyclesToClaim := map[int64]bool{}
	for {
		for _, cycleInfo := range detailedCyclesInfo.DetailedCyclesInfo {
			ethRewards, ok := big.NewInt(0).SetString(cycleInfo.MerkleProofInfo.Eth, 10)
			if !ok {
				return fmt.Errorf("Unable to parse eth rewards: %s", cycleInfo.MerkleProofInfo.Eth)
			}
			ethRewardsConverted := math.RoundDown(eth.WeiToEth(ethRewards), 2)
			sdRewards, ok := big.NewInt(0).SetString(cycleInfo.MerkleProofInfo.Sd, 10)
			if !ok {
				return fmt.Errorf("Unable to parse sd rewards: %s", cycleInfo.MerkleProofInfo.Sd)
			}
			sdRewardsConverted := math.RoundDown(eth.WeiToEth(sdRewards), 2)

			if ethRewardsConverted == 0 && sdRewardsConverted == 0 {
				continue
			}

			fmt.Printf("%-18d%-14.30s%-14.4f%-.4f\n", cycleInfo.MerkleProofInfo.Cycle, cycleInfo.CycleTime.Format("2006-01-02"), ethRewardsConverted, sdRewardsConverted)
		}

		cycleSelection := cliutils.Prompt("Select the cycles for which you wish to claim the rewards. Enter the cycles numbers in a comma separate format without any space (e.g. 1,2,3,4) or leave it blank to claim all cycles at once.", "^$|^\\d+(,\\d+)*$", "Unexpected input. Please enter a comma separated list of cycle numbers or leave it blank to claim all cycles at once.")
		if cycleSelection == "" {
			for _, cycle := range cycleIndexes {
				cyclesToClaim[cycle.Int64()] = true
			}
			break
		} else {
			elements := strings.Split(cycleSelection, ",")
			allValid := true
			for _, element := range elements {
				cycle, err := strconv.ParseUint(element, 0, 64)
				if err != nil {
					fmt.Printf("Unable to parse element: %s", element)
					allValid = false
				}

				// check if unclaimedCycles contains the cycle
				found := false
				for _, unclaimedCycle := range cycleIndexes {
					if unclaimedCycle.Int64() == int64(cycle) {
						found = true
						break
					}
				}
				if !found {
					fmt.Printf("Cycle %d is not in the list of unclaimed cycles. Please enter a valid cycle number\n", cycle)
					allValid = false
				} else {
					cyclesToClaim[int64(cycle)] = true
				}
			}

			if allValid {
				break
			}
		}
	}

	// convert the map to a cycle big int array
	cyclesToClaimArray := []*big.Int{}
	for cycle := range cyclesToClaim {
		cyclesToClaimArray = append(cyclesToClaimArray, big.NewInt(cycle))
	}

	// estimate gas
	fmt.Println("Estimating gas...")
	estimateGasResponse, err := staderClient.EstimateClaimSpRewardsGas(cyclesToClaimArray)
	if err != nil {
		return err
	}

	//Assign max fees
	err = gas.AssignMaxFeeAndLimit(estimateGasResponse.GasInfo, staderClient, c.Bool("yes"))
	if err != nil {
		return err
	}

	if !(c.Bool("yes") || cliutils.Confirm(fmt.Sprintf(
		"Are you sure you want to claim the rewards for cycles %v?", cyclesToClaimArray))) {
		fmt.Println("Cancelled.")
		return nil
	}

	fmt.Printf("Claiming rewards for cycles %v\n", cyclesToClaimArray)
	res, err := staderClient.ClaimSpRewards(cyclesToClaimArray)
	if err != nil {
		return err
	}

	cliutils.PrintTransactionHash(staderClient, res.TxHash)
	_, err = staderClient.WaitForTransaction(res.TxHash)
	if err != nil {
		return err
	}

	fmt.Printf("Rewards Claim successful for cycles: %v\n", cyclesToClaimArray)
	fmt.Printf("Please check your Operator Reward Address for the claimed rewards\n")

	return nil
}
