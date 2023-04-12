package node

import (
	"encoding/hex"
	"encoding/json"
	"fmt"
	"github.com/mitchellh/go-homedir"
	"github.com/stader-labs/stader-node/shared/services"
	"github.com/stader-labs/stader-node/shared/types/api"
	stader_backend "github.com/stader-labs/stader-node/shared/types/stader-backend"
	string_utils "github.com/stader-labs/stader-node/shared/utils/string-utils"
	"github.com/stader-labs/stader-node/stader-lib/node"
	socializing_pool "github.com/stader-labs/stader-node/stader-lib/socializing-pool"
	"github.com/urfave/cli"
	"math/big"
	"os"
)

func getClaimData(c *cli.Context, cycles []*big.Int) ([]*big.Int, []*big.Int, [][][32]byte, error) {
	cfg, err := services.GetConfig(c)
	if err != nil {
		return nil, nil, nil, err
	}

	// data to pass to socializing pool contract
	amountSd := []*big.Int{}
	amountEth := []*big.Int{}
	merkleProofs := [][][32]byte{}

	// get the proofs for each cycle and claim them. throw error if cycle proofs is not downloaded
	for _, cycle := range cycles {
		cycleMerkleRewardFile := cfg.StaderNode.GetSpRewardCyclePath(cycle.Int64(), true)
		expandedCycleMerkleRewardFile, err := homedir.Expand(cycleMerkleRewardFile)
		if err != nil {
			return nil, nil, nil, err
		}
		data, err := os.ReadFile(expandedCycleMerkleRewardFile)
		if err != nil {
			return nil, nil, nil, err
		}
		// need to make sure we have downloaded the data
		if os.IsNotExist(err) {
			return nil, nil, nil, err
		}
		merkleData := stader_backend.CycleMerkleProofs{}
		err = json.Unmarshal(data, &merkleData)
		if err != nil {
			return nil, nil, nil, err
		}

		amountSdBigInt := big.NewInt(0)
		amountSdBigInt, ok := amountSdBigInt.SetString(merkleData.Sd, 10)
		if !ok {
			return nil, nil, nil, fmt.Errorf("could not parse sd amount %s", merkleData.Sd)
		}

		// same thing above for eth
		amountEthBigInt := big.NewInt(0)
		amountEthBigInt, ok = amountEthBigInt.SetString(merkleData.Eth, 10)
		if !ok {
			return nil, nil, nil, fmt.Errorf("could not parse eth amount %s", merkleData.Eth)
		}

		amountSd = append(amountSd, amountSdBigInt)
		amountEth = append(amountEth, amountEthBigInt)

		// convert merkle proofs to [32]byte
		cycleMerkleProofs := [][32]byte{}
		for _, proof := range merkleData.Proof {
			merkleProofBytes, err := hex.DecodeString(proof[2:])
			if err != nil {
				return nil, nil, nil, err
			}
			var proofBytes [32]byte
			copy(proofBytes[:], merkleProofBytes[:32])
			cycleMerkleProofs = append(cycleMerkleProofs, proofBytes)
		}

		merkleProofs = append(merkleProofs, cycleMerkleProofs)
	}

	return amountSd, amountEth, merkleProofs, nil
}

func canClaimSpRewards(c *cli.Context) (*api.CanClaimSpRewardsResponse, error) {
	prn, err := services.GetPermissionlessNodeRegistry(c)
	if err != nil {
		return nil, err
	}
	sp, err := services.GetSocializingPoolContract(c)
	if err != nil {
		return nil, err
	}
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

	response := api.CanClaimSpRewardsResponse{}

	// check if operator is registered
	operatorId, err := node.GetOperatorId(prn, nodeAccount.Address, nil)
	if err != nil {
		return nil, err
	}
	if operatorId.Int64() == 0 {
		response.OperatorNotRegistered = true
		return &response, nil
	}

	claimedCycles := []*big.Int{}
	unclaimedCycles := []*big.Int{}
	// TODO - bchain - add check for this
	ineligibleCycles := []*big.Int{}
	cyclesToDownload := []*big.Int{}

	rewardDetails, err := socializing_pool.GetRewardDetails(sp, nil)
	if err != nil {
		return nil, err
	}

	for i := int64(1); i < rewardDetails.CurrentIndex.Int64(); i++ {
		cycle := big.NewInt(i)
		isClaimed, err := socializing_pool.HasClaimedRewards(sp, nodeAccount.Address, cycle, nil)
		if err != nil {
			return nil, err
		}
		if isClaimed {
			claimedCycles = append(claimedCycles, cycle)
			continue
		}

		unclaimedCycles = append(unclaimedCycles, cycle)
		// check if this cycle has been downloaded
		cycleMerkleRewardFile := cfg.StaderNode.GetSpRewardCyclePath(i, true)
		expandedCycleMerkleRewardFile, err := homedir.Expand(cycleMerkleRewardFile)
		if err != nil {
			return nil, err
		}
		_, err = os.Stat(expandedCycleMerkleRewardFile)
		if !os.IsNotExist(err) && err != nil {
			return nil, err
		}
		if os.IsNotExist(err) {
			cyclesToDownload = append(cyclesToDownload, cycle)
		}
	}

	response.ClaimedCycles = claimedCycles
	response.IneligibleCycles = ineligibleCycles
	response.UnclaimedCycles = unclaimedCycles
	response.CyclesToDownload = cyclesToDownload

	return &response, nil
}

func estimateSpRewardsGas(c *cli.Context, stringifiedCycles string) (*api.EstimateClaimSpRewardsGasResponse, error) {
	sp, err := services.GetSocializingPoolContract(c)
	if err != nil {
		return nil, err
	}
	w, err := services.GetWallet(c)
	if err != nil {
		return nil, err
	}

	response := api.EstimateClaimSpRewardsGasResponse{}

	cycles, err := string_utils.DestringifyArray(stringifiedCycles)
	if err != nil {
		return nil, err
	}

	amountSd, amountEth, merkleProofs, err := getClaimData(c, cycles)
	if err != nil {
		return nil, err
	}

	opts, err := w.GetNodeAccountTransactor()
	if err != nil {
		return nil, err
	}

	gasInfo, err := socializing_pool.EstimateClaimRewards(sp, cycles, amountSd, amountEth, merkleProofs, opts)
	if err != nil {
		return nil, err
	}

	response.GasInfo = gasInfo

	return &response, nil
}

func claimSpRewards(c *cli.Context, stringifiedCycles string) (*api.ClaimSpRewardsResponse, error) {
	sp, err := services.GetSocializingPoolContract(c)
	if err != nil {
		return nil, err
	}
	w, err := services.GetWallet(c)
	if err != nil {
		return nil, err
	}

	cycles, err := string_utils.DestringifyArray(stringifiedCycles)
	if err != nil {
		return nil, err
	}

	response := api.ClaimSpRewardsResponse{}

	amountSd, amountEth, merkleProofs, err := getClaimData(c, cycles)
	if err != nil {
		return nil, err
	}

	opts, err := w.GetNodeAccountTransactor()
	if err != nil {
		return nil, err
	}

	tx, err := socializing_pool.ClaimRewards(sp, cycles, amountSd, amountEth, merkleProofs, opts)
	if err != nil {
		return nil, err
	}

	response.TxHash = tx.Hash()

	return &response, nil
}
