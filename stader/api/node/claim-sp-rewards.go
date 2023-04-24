package node

import (
	"fmt"
	"math/big"
	"os"

	"github.com/mitchellh/go-homedir"
	"github.com/stader-labs/stader-node/shared/services"
	"github.com/stader-labs/stader-node/shared/types/api"
	string_utils "github.com/stader-labs/stader-node/shared/utils/string-utils"
	"github.com/stader-labs/stader-node/stader-lib/node"
	socializing_pool "github.com/stader-labs/stader-node/stader-lib/socializing-pool"
	"github.com/urfave/cli"
)

func IsEligibleForCycle(c *cli.Context, cycle *big.Int) (bool, error) {
	sp, err := services.GetSocializingPoolContract(c)
	if err != nil {
		return false, err
	}
	pnr, err := services.GetPermissionlessNodeRegistry(c)
	if err != nil {
		return false, err
	}
	w, err := services.GetWallet(c)
	if err != nil {
		return false, err
	}
	nodeAccount, err := w.GetNodeAccount()
	if err != nil {
		return false, err
	}

	cycleDetails, err := socializing_pool.GetRewardCycleDetails(sp, cycle, nil)
	if err != nil {
		return false, err
	}

	operatorId, err := node.GetOperatorId(pnr, nodeAccount.Address, nil)
	if err != nil {
		return false, err
	}
	operatorInfo, err := node.GetOperatorInfo(pnr, operatorId, nil)
	if err != nil {
		return false, err
	}

	lastSocializingPoolChangeBlock, err := node.GetSocializingPoolStateChangeBlock(pnr, operatorId, nil)
	if err != nil {
		return false, err
	}
	lastSocializingPoolChangeBlockUint := lastSocializingPoolChangeBlock.Uint64()
	cycleStartBlock := cycleDetails.StartBlock.Uint64()
	cycleEndBlock := cycleDetails.EndBlock.Uint64()

	// TODO - bchain - simplify this logic using boolean algebra and make it succint
	switch operatorInfo.OptedForSocializingPool {
	case true:
		return lastSocializingPoolChangeBlockUint < cycleStartBlock || (lastSocializingPoolChangeBlockUint > cycleStartBlock && lastSocializingPoolChangeBlockUint < cycleEndBlock) || !(lastSocializingPoolChangeBlockUint > cycleEndBlock), nil
	case false:
		return !(lastSocializingPoolChangeBlockUint < cycleStartBlock) || (lastSocializingPoolChangeBlockUint > cycleStartBlock && lastSocializingPoolChangeBlockUint < cycleEndBlock) || (lastSocializingPoolChangeBlockUint > cycleEndBlock), nil
	default:
		// will never happen
		return false, fmt.Errorf("operator info opted for socializing pool is not a boolean")
	}
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
	cfg, err := services.GetConfig(c)
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
	ineligibleCycles := []*big.Int{}
	cyclesToDownload := []*big.Int{}

	rewardDetails, err := socializing_pool.GetRewardDetails(sp, nil)
	if err != nil {
		return nil, err
	}

	for i := int64(1); i < rewardDetails.CurrentIndex.Int64(); i++ {
		cycle := big.NewInt(i)
		isEligible, err := IsEligibleForCycle(c, big.NewInt(i))
		if err != nil {
			return nil, err
		}
		if !isEligible {
			ineligibleCycles = append(ineligibleCycles, cycle)
			continue
		}

		isClaimed, err := socializing_pool.HasClaimedRewards(sp, nodeAccount.Address, cycle, nil)
		if err != nil {
			return nil, err
		}
		if isClaimed {
			claimedCycles = append(claimedCycles, cycle)
		} else {
			unclaimedCycles = append(unclaimedCycles, cycle)
		}

		// download merkle proofs if they don't exist even for claimed. it is useful for status displaying
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
