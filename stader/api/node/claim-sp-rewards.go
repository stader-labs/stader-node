package node

import (
	"github.com/stader-labs/stader-node/shared/services"
	"github.com/stader-labs/stader-node/shared/types/api"
	"github.com/stader-labs/stader-node/shared/utils/eth1"
	string_utils "github.com/stader-labs/stader-node/shared/utils/string-utils"
	socializing_pool "github.com/stader-labs/stader-node/stader-lib/socializing-pool"
	"github.com/urfave/cli"
	"math/big"
)

func GetCyclesDetailedInfo(c *cli.Context, stringifiedCycles string) (*api.CyclesDetailedInfo, error) {
	cfg, err := services.GetConfig(c)
	if err != nil {
		return nil, err
	}
	sp, err := services.GetSocializingPoolContract(c)
	if err != nil {
		return nil, err
	}

	cycles, err := string_utils.DestringifyArray(stringifiedCycles)
	if err != nil {
		return nil, err
	}

	response := api.CyclesDetailedInfo{}
	merkleProofs := []api.DetailedMerkleProofInfo{}
	for _, cycle := range cycles {
		merkleCycleProof, exists, err := cfg.StaderNode.ReadCycleCache(cycle.Int64())
		if err != nil {
			return nil, err
		}
		if !exists {
			continue
		}

		currentBlock, err := eth1.GetCurrentBlockNumber(c)
		if err != nil {
			return nil, err
		}
		cycleDetails, err := socializing_pool.GetRewardCycleDetails(sp, cycle, nil)
		if err != nil {
			return nil, err
		}

		cycleStartBlock := currentBlock
		if cycleDetails.StartBlock.Cmp(big.NewInt(int64(currentBlock))) < 0 {
			cycleStartBlock = cycleDetails.StartBlock.Uint64()
		}
		cycleStartTime, err := eth1.ConvertBlockToTimestamp(c, int64(cycleStartBlock))
		if err != nil {
			return nil, err
		}

		merkleProofs = append(merkleProofs, api.DetailedMerkleProofInfo{
			MerkleProofInfo: merkleCycleProof,
			CycleTime:       cycleStartTime,
		})
	}

	response.DetailedCyclesInfo = merkleProofs

	return &response, nil
}

func canClaimSpRewards(c *cli.Context) (*api.CanClaimSpRewardsResponse, error) {
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
	w, err := services.GetWallet(c)
	if err != nil {
		return nil, err
	}
	nodeAccount, err := w.GetNodeAccount()
	if err != nil {
		return nil, err
	}

	response := api.CanClaimSpRewardsResponse{}

	isPaused, err := socializing_pool.IsSocializingPoolPaused(sp, nil)
	if err != nil {
		return nil, err
	}
	if isPaused {
		response.SocializingPoolContractPaused = true
		return &response, nil
	}

	claimedCycles := []*big.Int{}
	unclaimedCycles := []*big.Int{}
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
		} else {
			unclaimedCycles = append(unclaimedCycles, cycle)
		}
	}

	response.ClaimedCycles = claimedCycles
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
	cfg, err := services.GetConfig(c)
	if err != nil {
		return nil, err
	}
	response := api.EstimateClaimSpRewardsGasResponse{}

	cycles, err := string_utils.DestringifyArray(stringifiedCycles)
	if err != nil {
		return nil, err
	}

	amountSd, amountEth, merkleProofs, err := cfg.StaderNode.GetClaimData(cycles)
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

	cfg, err := services.GetConfig(c)
	if err != nil {
		return nil, err
	}

	response := api.ClaimSpRewardsResponse{}
	amountSd, amountEth, merkleProofs, err := cfg.StaderNode.GetClaimData(cycles)
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
