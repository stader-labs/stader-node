package node

import (
	"encoding/json"
	"fmt"
	"github.com/mitchellh/go-homedir"
	"github.com/stader-labs/stader-node/shared/services"
	"github.com/stader-labs/stader-node/shared/types/api"
	"github.com/stader-labs/stader-node/shared/utils/stader"
	"github.com/stader-labs/stader-node/stader-lib/node"
	socializing_pool "github.com/stader-labs/stader-node/stader-lib/socializing-pool"
	"github.com/urfave/cli"
	"os"
)

func canDownloadSpMerkleProofs(c *cli.Context) (*api.CanDownloadSpMerkleProofsResponse, error) {
	// check if operator is registered or not
	w, err := services.GetWallet(c)
	if err != nil {
		return nil, err
	}
	prn, err := services.GetPermissionlessNodeRegistry(c)
	if err != nil {
		return nil, err
	}
	sp, err := services.GetSocializingPoolContract(c)
	if err != nil {
		return nil, err
	}
	cfg, err := services.GetConfig(c)
	if err != nil {
		return nil, err
	}

	nodeAccount, err := w.GetNodeAccount()
	if err != nil {
		return nil, err
	}

	response := api.CanDownloadSpMerkleProofsResponse{}
	operatorId, err := node.GetOperatorId(prn, nodeAccount.Address, nil)
	if err != nil {
		return nil, err
	}
	if operatorId.Int64() == 0 {
		response.OperatorNotRegistered = true
		return &response, nil
	}

	fmt.Printf("Getting reward details for socializing pool...\n")
	// check if all cycles are present
	rewardDetails, err := socializing_pool.GetRewardDetails(sp, nil)
	if err != nil {
		return nil, err
	}
	fmt.Printf("Reward details: %v\n", rewardDetails)
	currentIndex := rewardDetails.CurrentIndex.Int64()
	missingCycles := []int64{}
	// iterate thru all cycles starting from 1
	for i := int64(1); i <= currentIndex; i++ {
		// TODO - check if operator is eligible for cycle before downloading it
		cycleMerkleRewardFile := cfg.StaderNode.GetSpRewardCyclePath(i, true)
		// check if file exists or not
		_, err := os.Stat(cycleMerkleRewardFile)
		if err != nil {
			return nil, err
		}
		if os.IsNotExist(err) {
			missingCycles = append(missingCycles, i)
		}
	}

	// no missing cycles
	if len(missingCycles) == 0 {
		response.NoMissingCycles = true
		return &response, nil
	}

	response.MissingCycles = missingCycles
	response.CurrentCycle = currentIndex

	return &response, nil
}

func downloadSpMerkleProofs(c *cli.Context) (*api.DownloadSpMerkleProofsResponse, error) {
	w, err := services.GetWallet(c)
	if err != nil {
		return nil, err
	}
	cfg, err := services.GetConfig(c)
	if err != nil {
		return nil, err
	}
	sp, err := services.GetSocializingPoolContract(c)
	if err != nil {
		return nil, err
	}
	rewardDetails, err := socializing_pool.GetRewardDetails(sp, nil)
	if err != nil {
		return nil, err
	}
	nodeAccount, err := w.GetNodeAccount()
	if err != nil {
		return nil, err
	}

	currentIndex := rewardDetails.CurrentIndex.Int64()
	missingCycles := []int64{}
	// iterate thru all cycles starting from 1
	for i := int64(1); i <= currentIndex; i++ {
		// TODO - bchain - add check for eligibility
		cycleRewardFile := cfg.StaderNode.GetSpRewardCyclePath(i, true)
		// check if file exists or not
		_, err := os.Stat(cycleRewardFile)
		if err != nil {
			return nil, err
		}
		if os.IsNotExist(err) {
			missingCycles = append(missingCycles, i)
		}
	}

	for _, cycle := range missingCycles {
		cycleMerkleProof, err := stader.GetCycleMerkleProofsForOperator(cycle, nodeAccount.Address)
		if err != nil {
			return nil, err
		}
		fmt.Printf("downloadSpMerkleProof: cycleMerkleProof: %+v", cycleMerkleProof)

		cycleMerkleProofFile := cfg.StaderNode.GetSpRewardCyclePath(cycle, true)
		absolutePathOfProofFile, err := homedir.Expand(cycleMerkleProofFile)
		fmt.Printf("downloadSpMerkleProof: absolutePathOfProofFile: %+v", absolutePathOfProofFile)
		if err != nil {
			return nil, err
		}

		file, err := os.Create(absolutePathOfProofFile)
		if err != nil {
			return nil, err
		}

		encoder := json.NewEncoder(file)
		err = encoder.Encode(cycleMerkleProof)
		if err != nil {
			return nil, fmt.Errorf("Error encoding JSON: %v", err)
		}
	}

	return nil, nil
}
