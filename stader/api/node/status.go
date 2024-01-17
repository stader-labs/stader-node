package node

import (
	"math/big"
	"time"

	stader_backend "github.com/stader-labs/stader-node/shared/types/stader-backend"
	"github.com/stader-labs/stader-node/shared/utils/eth1"
	pool_utils "github.com/stader-labs/stader-node/stader-lib/pool-utils"
	socializing_pool "github.com/stader-labs/stader-node/stader-lib/socializing-pool"
	stader_config "github.com/stader-labs/stader-node/stader-lib/stader-config"
	"github.com/stader-labs/stader-node/stader-lib/types"

	"github.com/stader-labs/stader-node/shared/services"
	"github.com/stader-labs/stader-node/shared/types/api"
	"github.com/stader-labs/stader-node/shared/utils/stdr"
	"github.com/stader-labs/stader-node/stader-lib/node"
	sd_collateral "github.com/stader-labs/stader-node/stader-lib/sd-collateral"
	"github.com/stader-labs/stader-node/stader-lib/tokens"
	"github.com/urfave/cli"
)

func GetClaimedAndUnclaimedSocializingPoolMerkles(c *cli.Context) ([]stader_backend.CycleMerkleProofs, []stader_backend.CycleMerkleProofs, error) {
	cfg, err := services.GetConfig(c)
	if err != nil {
		return nil, nil, err
	}
	sp, err := services.GetSocializingPoolContract(c)
	if err != nil {
		return nil, nil, err
	}
	w, err := services.GetWallet(c)
	if err != nil {
		return nil, nil, err
	}
	nodeAccount, err := w.GetNodeAccount()
	if err != nil {
		return nil, nil, err
	}

	rewardDetails, err := socializing_pool.GetRewardDetails(sp, nil)
	if err != nil {
		return nil, nil, err
	}

	unclaimedMerkles := []stader_backend.CycleMerkleProofs{}
	claimedMerkles := []stader_backend.CycleMerkleProofs{}
	for i := int64(1); i < rewardDetails.CurrentIndex.Int64(); i++ {
		cycleMerkleProof, exists, err := cfg.StaderNode.ReadCycleCache(i)
		if err != nil {
			return nil, nil, err
		}
		if !exists {
			continue
		}
		claimed, err := socializing_pool.HasClaimedRewards(sp, nodeAccount.Address, big.NewInt(i), nil)
		if err != nil {
			return nil, nil, err
		}

		if claimed {
			claimedMerkles = append(claimedMerkles, cycleMerkleProof)
		} else {
			unclaimedMerkles = append(unclaimedMerkles, cycleMerkleProof)
		}
	}

	return claimedMerkles, unclaimedMerkles, nil
}

func getStatus(c *cli.Context) (*api.NodeStatusResponse, error) {

	// Get services
	if err := services.RequireNodeWallet(c); err != nil {
		return nil, err
	}
	w, err := services.GetWallet(c)
	if err != nil {
		return nil, err
	}
	pnr, err := services.GetPermissionlessNodeRegistry(c)
	if err != nil {
		return nil, err
	}
	sdt, err := services.GetSdTokenContract(c)
	if err != nil {
		return nil, err
	}
	sdc, err := services.GetSdCollateralContract(c)
	if err != nil {
		return nil, err
	}
	sdcfg, err := services.GetStaderConfigContract(c)
	if err != nil {
		return nil, err
	}
	putils, err := services.GetPoolUtilsContract(c)
	if err != nil {
		return nil, err
	}
	bc, err := services.GetBeaconClient(c)
	if err != nil {
		return nil, err
	}
	sp, err := services.GetSocializingPoolContract(c)
	if err != nil {
		return nil, err
	}
	orc, err := services.GetOperatorRewardsCollectorContract(c)
	if err != nil {
		return nil, err
	}

	// Response
	response := api.NodeStatusResponse{}

	//fmt.Printf("Getting node account...\n")
	nodeAccount, err := w.GetNodeAccount()
	if err != nil {
		return nil, err
	}

	response.AccountAddress = nodeAccount.Address

	//fmt.Printf("Getting node account balances...\n")
	accountEthBalance, err := tokens.GetEthBalance(pnr.Client, nodeAccount.Address, nil)
	if err != nil {
		return nil, err
	}
	//fmt.Printf("Getting node account SD balance...\n")
	accountSdBalance, err := tokens.BalanceOf(sdt, nodeAccount.Address, nil)
	if err != nil {
		return nil, err
	}

	response.AccountBalances.ETH = accountEthBalance
	response.AccountBalances.Sd = accountSdBalance

	//fmt.Printf("Getting socializing pool address...\n")
	socializingPoolAddress, err := stader_config.GetSocializingPoolContractAddress(sdcfg, nil)
	if err != nil {
		return nil, err
	}
	response.SocializingPoolAddress = socializingPoolAddress

	//fmt.Printf("Getting operator id...\n")
	operatorId, err := node.GetOperatorId(pnr, nodeAccount.Address, nil)
	if err != nil {
		return nil, err
	}
	//fmt.Printf("Getting operator info...\n")
	operatorRegistry, err := node.GetOperatorInfo(pnr, operatorId, nil)
	if err != nil {
		return nil, err
	}

	if operatorRegistry.OperatorName != "" {
		response.Registered = true
		response.OperatorId = operatorId
		response.OperatorName = operatorRegistry.OperatorName
		response.OperatorActive = operatorRegistry.Active
		response.OperatorAddress = operatorRegistry.OperatorAddress
		response.OperatorRewardAddress = operatorRegistry.OperatorRewardAddress
		response.OptedInForSocializingPool = operatorRegistry.OptedForSocializingPool

		//fmt.Printf("Getting operator node el reward balance\n")
		// non socializing pool fee recepient
		operatorElRewardAddress, err := node.GetNodeElRewardAddress(pnr, 1, operatorId, nil)
		if err != nil {
			return nil, err
		}
		//fmt.Printf("Getting operator node el reward balance\n")
		elRewardAddressBalance, err := tokens.GetEthBalance(pnr.Client, operatorElRewardAddress, nil)
		if err != nil {
			return nil, err
		}
		//fmt.Printf("Getting operator node el reward share\n")
		operatorElRewards, err := pool_utils.CalculateRewardShare(putils, 1, elRewardAddressBalance, nil)
		if err != nil {
			return nil, err
		}
		response.OperatorELRewardsAddress = operatorElRewardAddress
		response.OperatorELRewardsAddressBalance = operatorElRewards.OperatorShare

		operatorRewardCollectorBalance, err := node.GetOperatorRewardsCollectorBalance(orc, nodeAccount.Address, nil)
		if err != nil {
			return nil, err
		}
		response.OperatorRewardCollectorBalance = operatorRewardCollectorBalance

		operatorWithdrawableEth, err := node.WithdrawableInEth(orc, nodeAccount.Address, nil)
		if err != nil {
			return nil, err
		}
		response.OperatorWithdrawableEth = operatorWithdrawableEth

		//fmt.Printf("Getting operator reward address balance\n")
		operatorReward, err := tokens.GetEthBalance(pnr.Client, operatorRegistry.OperatorRewardAddress, nil)
		if err != nil {
			return nil, err
		}
		response.OperatorRewardInETH = operatorReward

		//fmt.Printf("getting operator sd collateral balance\n")
		// get operator deposited sd collateral
		operatorSdCollateral, err := sd_collateral.GetOperatorSdBalance(sdc, nodeAccount.Address, nil)
		if err != nil {
			return nil, err
		}
		response.DepositedSdCollateral = operatorSdCollateral

		//fmt.Printf("getting operator sd collateral worth validators\n")
		// total registerable validators
		totalSdWorthValidators, err := sd_collateral.GetMaxValidatorSpawnable(sdc, operatorSdCollateral, 1, nil)
		if err != nil {
			return nil, err
		}
		response.SdCollateralWorthValidators = totalSdWorthValidators

		//fmt.Printf("Getting reward details\n")
		rewardCycleDetails, err := socializing_pool.GetRewardDetails(sp, nil)
		if err != nil {
			return nil, err
		}
		response.SocializingPoolRewardCycleDetails = rewardCycleDetails
		socializingPoolStartTimestamp := time.Now()
		response.SocializingPoolStartTime = socializingPoolStartTimestamp

		//fmt.Printf("Get total validator keys\n")
		totalValidatorKeys, err := node.GetTotalValidatorKeys(pnr, operatorId, nil)
		if err != nil {
			return nil, err
		}

		//fmt.Printf("Get total non terminal validator keys\n")
		totalNonTerminalValidatorKeys, err := node.GetTotalNonTerminalValidatorKeys(pnr, nodeAccount.Address, totalValidatorKeys, nil)
		if err != nil {
			return nil, err
		}
		//fmt.Printf("Total non terminal validators %d\n", totalNonTerminalValidatorKeys)

		response.TotalNonTerminalValidators = big.NewInt(int64(totalNonTerminalValidatorKeys))

		totalValidatorClRewards := big.NewInt(0)
		validatorInfoArray := make([]stdr.ValidatorInfo, totalValidatorKeys.Int64())

		validatorInfoMap, _, err := stdr.GetAllValidatorsRegisteredWithOperator(pnr, operatorId, nodeAccount.Address, nil)
		if err != nil {
			return nil, err
		}

		i := 0
		for _, validatorContractInfo := range validatorInfoMap {
			withdrawVaultBalance, err := tokens.GetEthBalance(pnr.Client, validatorContractInfo.WithdrawVaultAddress, nil)
			if err != nil {
				return nil, err
			}
			withdrawVaultRewardShares, err := pool_utils.CalculateRewardShare(putils, 1, withdrawVaultBalance, nil)
			if err != nil {
				return nil, err
			}
			rewardsThreshold, err := stader_config.GetRewardsThreshold(sdcfg, nil)
			if err != nil {
				return nil, err
			}
			crossedRewardThreshold := false
			if withdrawVaultBalance.Cmp(rewardsThreshold) > 0 {
				crossedRewardThreshold = true
			} else {
				totalValidatorClRewards.Add(totalValidatorClRewards, withdrawVaultRewardShares.OperatorShare)
			}

			withdrawVaultWithdrawShares, err := node.CalculateValidatorWithdrawVaultWithdrawShare(pnr.Client, validatorContractInfo.WithdrawVaultAddress, nil)
			if err != nil {
				return nil, err
			}
			validatorWithdrawVaultWithdrawShares := withdrawVaultWithdrawShares.OperatorShare

			validatorBeaconStatus, err := bc.GetValidatorStatus(types.BytesToValidatorPubkey(validatorContractInfo.Pubkey), nil)
			if err != nil {
				return nil, err
			}

			validatorDisplayStatus, err := stdr.GetValidatorRunningStatus(validatorBeaconStatus, validatorContractInfo)
			if err != nil {
				return nil, err
			}

			depositTime, err := eth1.ConvertBlockToTimestamp(c, validatorContractInfo.DepositBlock.Int64())
			if err != nil {
				return nil, err
			}
			withdrawTime, err := eth1.ConvertBlockToTimestamp(c, validatorContractInfo.WithdrawnBlock.Int64())
			if err != nil {
				return nil, err
			}

			validatorInfo := stdr.ValidatorInfo{
				Status:                           validatorContractInfo.Status,
				StatusToDisplay:                  validatorDisplayStatus,
				Pubkey:                           validatorContractInfo.Pubkey,
				PreDepositSignature:              validatorContractInfo.PreDepositSignature,
				DepositSignature:                 validatorContractInfo.DepositSignature,
				WithdrawVaultAddress:             validatorContractInfo.WithdrawVaultAddress,
				WithdrawVaultRewardBalance:       withdrawVaultRewardShares.OperatorShare,
				CrossedRewardsThreshold:          crossedRewardThreshold,
				WithdrawVaultWithdrawableBalance: validatorWithdrawVaultWithdrawShares,
				OperatorId:                       validatorContractInfo.OperatorId,
				DepositBlock:                     validatorContractInfo.DepositBlock,
				DepositTime:                      depositTime,
				WithdrawnBlock:                   validatorContractInfo.WithdrawnBlock,
				WithdrawnTime:                    withdrawTime,
			}

			validatorInfoArray[i] = validatorInfo

			i += 1
		}

		response.ValidatorInfos = validatorInfoArray
		response.TotalValidatorClRewards = totalValidatorClRewards

		//fmt.Printf("Getting operator claimed and unclaimed socializing pool merkles\n")
		claimedMerkles, unclaimedMerkles, err := GetClaimedAndUnclaimedSocializingPoolMerkles(c)
		if err != nil {
			return nil, err
		}

		response.ClaimedSocializingPoolMerkles = claimedMerkles
		response.UnclaimedSocializingPoolMerkles = unclaimedMerkles

	} else {
		response.DepositedSdCollateral = big.NewInt(0)
		response.SdCollateralWorthValidators = big.NewInt(0)
		response.ValidatorInfos = []stdr.ValidatorInfo{}
		response.Registered = false
	}

	// Return response
	return &response, nil
}
