package node

import (
	"encoding/json"
	"fmt"
	"github.com/mitchellh/go-homedir"
	"github.com/stader-labs/stader-node/shared/services"
	"github.com/stader-labs/stader-node/shared/types/api"
	"github.com/stader-labs/stader-node/shared/utils/stader"
	socializing_pool "github.com/stader-labs/stader-node/stader-lib/socializing-pool"
	"github.com/urfave/cli"
	"os"
)

func canDownloadSpMerkleProofs(c *cli.Context) (*api.CanDownloadSpMerkleProofsResponse, error) {
	if err := services.RequireNodeWallet(c); err != nil {
		return nil, err
	}
	if err := services.RequireNodeRegistered(c); err != nil {
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

	response := api.CanDownloadSpMerkleProofsResponse{}

	// check if all cycles are present
	rewardDetails, err := socializing_pool.GetRewardDetails(sp, nil)
	if err != nil {
		return nil, err
	}
	currentIndex := rewardDetails.CurrentIndex.Int64()
	missingCycles := []int64{}
	// iterate thru all cycles starting from 1
	for i := int64(1); i < currentIndex; i++ {
		// download all cycles irrespective if the NO claim or not claimed.
		cycleMerkleRewardFile := cfg.StaderNode.GetSpRewardCyclePath(i, true)
		// check if file exists or not
		_, err = os.Stat(cycleMerkleRewardFile)
		if !os.IsNotExist(err) && err != nil {
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
	nodeAccount, err := w.GetNodeAccount()
	if err != nil {
		return nil, err
	}

	response := api.DownloadSpMerkleProofsResponse{}

	allMerkleProofs, err := stader.GetAllMerkleProofsForOperator(c, nodeAccount.Address)
	if err != nil {
		return nil, err
	}

	downloadedCycles := []int64{}

	for _, cycleMerkleProof := range allMerkleProofs {

		cycleMerkleProofFile := cfg.StaderNode.GetSpRewardCyclePath(cycleMerkleProof.Cycle, true)
		absolutePathOfProofFile, err := homedir.Expand(cycleMerkleProofFile)
		if err != nil {
			return nil, err
		}

		// proof has already been downloaded
		_, err = os.Stat(cycleMerkleProofFile)
		if !os.IsNotExist(err) && err != nil {
			return nil, err
		}
		if !os.IsNotExist(err) {
			continue
		}

		file, err := os.Create(absolutePathOfProofFile)
		if err != nil {
			return nil, err
		}

		encoder := json.NewEncoder(file)
		err = encoder.Encode(cycleMerkleProof)
		if err != nil {
			return nil, fmt.Errorf("error encoding JSON: %v", err)
		}

		downloadedCycles = append(downloadedCycles, cycleMerkleProof.Cycle)
	}

	response.DownloadedCycles = downloadedCycles

	return &response, nil
}
