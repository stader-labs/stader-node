package node

import (
	"fmt"
	"math/big"
	"strconv"
	"strings"

	"github.com/stader-labs/stader-node/shared/services/gas"
	"github.com/stader-labs/stader-node/shared/services/stader"
	stader_backend "github.com/stader-labs/stader-node/shared/types/stader-backend"
	cliutils "github.com/stader-labs/stader-node/shared/utils/cli"
	"github.com/stader-labs/stader-node/shared/utils/math"
	"github.com/stader-labs/stader-node/stader-lib/utils/eth"
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

	downloadRes, err := staderClient.DownloadSpMerkleProofs()
	if err != nil {
		return err
	}
	if len(downloadRes.DownloadedCycles) != 0 {
		fmt.Printf("Merkle proofs downloaded for cycles %v!\n", downloadRes.DownloadedCycles)
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

	indexedDetailedCyclesInfo := map[int64]stader_backend.CycleMerkleProofs{}
	for _, cycleInfo := range detailedCyclesInfo.DetailedCyclesInfo {
		indexedDetailedCyclesInfo[cycleInfo.MerkleProofInfo.Cycle] = cycleInfo.MerkleProofInfo
	}

	cycleIndexes := []*big.Int{}
	for _, cycleInfo := range detailedCyclesInfo.DetailedCyclesInfo {
		cycleIndexes = append(cycleIndexes, big.NewInt(cycleInfo.MerkleProofInfo.Cycle))
	}

	fmt.Println("Following are the unclaimed cycles, Please enter in a comma separated string the cycles you want to claim rewards for:")

	fmt.Printf("\n%-18s%-14.30s%-14.10s%-10s\n", "Cycle Number", "Cycle Date", "ETH Rewards", "SD Rewards")
	cyclesToClaim := map[int64]bool{}
	for {
		for _, cycleInfo := range detailedCyclesInfo.DetailedCyclesInfo {
			ethRewards, ok := big.NewInt(0).SetString(cycleInfo.MerkleProofInfo.Eth, 10)
			if !ok {
				return fmt.Errorf("Unable to parse eth rewards: %s", cycleInfo.MerkleProofInfo.Eth)
			}
			ethRewardsConverted := math.RoundDown(eth.WeiToEth(ethRewards), 5)
			sdRewards, ok := big.NewInt(0).SetString(cycleInfo.MerkleProofInfo.Sd, 10)
			if !ok {
				return fmt.Errorf("Unable to parse sd rewards: %s", cycleInfo.MerkleProofInfo.Sd)
			}
			sdRewardsConverted := math.RoundDown(eth.WeiToEth(sdRewards), 5)

			if ethRewards.Cmp(big.NewInt(0)) == 0 && sdRewards.Cmp(big.NewInt(0)) == 0 {
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

	totalClaimableEth := big.NewInt(0)
	totalClaimableSd := big.NewInt(0)
	for _, cycle := range cyclesToClaimArray {
		cycleInfo := indexedDetailedCyclesInfo[cycle.Int64()]
		ethRewards, ok := big.NewInt(0).SetString(cycleInfo.Eth, 10)
		if !ok {
			return fmt.Errorf("Unable to parse eth rewards: %s", cycleInfo.Eth)
		}
		totalClaimableEth = totalClaimableEth.Add(totalClaimableEth, ethRewards)
		sdRewards, ok := big.NewInt(0).SetString(cycleInfo.Sd, 10)
		if !ok {
			return fmt.Errorf("Unable to parse sd rewards: %s", cycleInfo.Sd)
		}
		totalClaimableSd = totalClaimableSd.Add(totalClaimableSd, sdRewards)
	}

	depositSd := false
	if totalClaimableSd.Cmp(big.NewInt(0)) > 0 {
		fmt.Printf("You will claim %s and %s with the following selection - cycles %v\n\n", eth.DisplayAmountInUnits(totalClaimableSd, "sd"), eth.DisplayAmountInUnits(totalClaimableEth, "eth"), cyclesToClaimArray)
		fmt.Printf("Your ETH rewards will be sent to your Reward Address\n")
		fmt.Printf("For SD rewards, you can claim all the rewards to your Reward Address or redeposit them as SD collateral to earn more rewards\n")

		fmt.Printf("Please select one of the following options:\n")
		fmt.Printf("1. Claim all SD rewards to your Reward Address\n")
		fmt.Printf("2. Redeposit all the SD rewards as SD collateral for additional earnings\n")

		option := cliutils.Prompt("", "^(1|2)$", "Please enter a valid option")
		if option == "1" {
			if !cliutils.Confirm(fmt.Sprintf(
				"Are you sure you want to claim %s and %s for cycles %v to your reward address?", eth.DisplayAmountInUnits(totalClaimableEth, "eth"), eth.DisplayAmountInUnits(totalClaimableSd, "sd"), cyclesToClaimArray)) {
				fmt.Println("Claim Cancelled.")
				return nil
			}
		} else if option == "2" {
			if !cliutils.Confirm(fmt.Sprintf(
				"Your %s rewards will be sent to your Reward Address.\nFor your %s rewards, are you sure you want to re-deposit it as SD collateral for additional earnings?", eth.DisplayAmountInUnits(totalClaimableEth, "eth"), eth.DisplayAmountInUnits(totalClaimableSd, "sd"))) {
				fmt.Println("Claim Cancelled.")
				return nil
			}
			depositSd = true
		}
	} else {
		if !cliutils.Confirm(fmt.Sprintf(
			"Are you sure you want to claim %s ETH for cycles %v to your reward address?", totalClaimableEth.String(), cyclesToClaimArray)) {
			fmt.Println("Cancelled.")
			return nil
		}
	}

	// estimate gas
	fmt.Println("Estimating gas...")
	estimateGasResponse, err := staderClient.EstimateClaimSpRewardsGas(cyclesToClaimArray, depositSd)
	if err != nil {
		return err
	}

	//Assign max fees
	err = gas.AssignMaxFeeAndLimit(estimateGasResponse.GasInfo, staderClient, c.Bool("yes"))
	if err != nil {
		return err
	}

	fmt.Printf("Claiming rewards for cycles %v\n", cyclesToClaimArray)
	res, err := staderClient.ClaimSpRewards(cyclesToClaimArray, depositSd)
	if err != nil {
		return err
	}

	cliutils.PrintTransactionHash(staderClient, res.TxHash)
	_, err = staderClient.WaitForTransaction(res.TxHash)
	if err != nil {
		return err
	}

	fmt.Printf("Transaction Successful\n")
	if depositSd {
		fmt.Printf("%s rewards have been sent to your Reward Address and %s rewards have been re-deposited as SD collateral\n", eth.DisplayAmountInUnits(totalClaimableEth, "eth"), eth.DisplayAmountInUnits(totalClaimableSd, "sd"))
	} else {
		if totalClaimableSd.Cmp(big.NewInt(0)) <= 0 {
			fmt.Printf("%s rewards have been sent to your Reward Address\n", eth.DisplayAmountInUnits(totalClaimableEth, "eth"))
		} else {
			fmt.Printf("%s rewards and %s rewards have been sent to your Reward Address\n", eth.DisplayAmountInUnits(totalClaimableSd, "sd"), eth.DisplayAmountInUnits(totalClaimableEth, "eth"))
		}
	}

	return nil
}
