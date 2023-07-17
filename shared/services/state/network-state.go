package state

import (
	"fmt"
	"github.com/stader-labs/stader-node/shared/services"
	"github.com/stader-labs/stader-node/shared/utils/stdr"
	"github.com/stader-labs/stader-node/stader-lib/contracts"
	"github.com/urfave/cli"
	"math/big"
	"time"

	"github.com/stader-labs/stader-node/shared/utils/math"
	"github.com/stader-labs/stader-node/stader-lib/utils/eth"

	"github.com/stader-labs/stader-node/shared/utils/eth2"
	penalty_tracker "github.com/stader-labs/stader-node/stader-lib/penalty-tracker"
	pool_utils "github.com/stader-labs/stader-node/stader-lib/pool-utils"
	socializing_pool "github.com/stader-labs/stader-node/stader-lib/socializing-pool"
	stader_config "github.com/stader-labs/stader-node/stader-lib/stader-config"
	stake_pool_manager "github.com/stader-labs/stader-node/stader-lib/stake-pool-manager"

	"github.com/stader-labs/stader-node/shared/services/beacon"
	"github.com/stader-labs/stader-node/shared/services/config"
	"github.com/stader-labs/stader-node/shared/utils/log"
	"github.com/stader-labs/stader-node/stader-lib/node"
	sd_collateral "github.com/stader-labs/stader-node/stader-lib/sd-collateral"
	"github.com/stader-labs/stader-node/stader-lib/stader"
	"github.com/stader-labs/stader-node/stader-lib/tokens"
	"github.com/stader-labs/stader-node/stader-lib/types"

	"github.com/ethereum/go-ethereum/common"
)

type MetricDetails struct {
	// Network details

	// done
	SdPrice float64

	EthPrice float64
	// done
	TotalValidators *big.Int
	// done
	TotalActiveValidators *big.Int
	// done
	TotalQueuedValidators *big.Int
	// done
	TotalOperators *big.Int
	// done
	TotalStakedSd float64
	// done
	TotalStakedEthByNos *big.Int
	// done
	TotalEthxSupply float64
	// done
	TotalStakedEthByUsers *big.Int
	// done
	MinEthThreshold float64
	// done
	MaxEthThreshold float64

	// Validator specific info

	// done
	ActiveValidators *big.Int
	// done
	StaderQueuedValidators *big.Int
	// done
	BeaconChainQueuedValidators *big.Int
	// done
	SlashedValidators *big.Int
	// done
	ExitingValidators *big.Int
	// done
	WithdrawnValidators *big.Int
	// done
	InitializedValidators *big.Int
	// done
	InvalidSignatureValidators *big.Int
	// done
	FrontRunValidators *big.Int
	// done
	FundsSettledValidators *big.Int
	// done
	ValidatorStatusMap map[types.ValidatorPubkey]beacon.ValidatorStatus
	ValidatorInfoMap   map[types.ValidatorPubkey]contracts.Validator

	// done
	CumulativePenalty float64
	// done
	UnclaimedClRewards float64
	// done
	UnclaimedNonSocializingPoolElRewards float64
	// done
	CollateralRatio float64
	//
	CollateralRatioInSd float64

	// done
	ClaimedSocializingPoolElRewards float64
	// done
	ClaimedSocializingPoolSdRewards float64
	// done
	UnclaimedSocializingPoolElRewards float64
	// done
	UnclaimedSocializingPoolSDRewards float64
	// done
	NextSocializingPoolRewardCycle types.RewardCycleDetails
	// done
	OperatorStakedSd      float64
	OperatorStakedSdInEth float64
	// done
	OperatorEthCollateral float64
}

type MetricsCache struct {
	// Block / slot for this state
	ElBlockNumber    uint64
	BeaconSlotNumber uint64
	BeaconConfig     beacon.Eth2Config

	StaderNetworkDetails MetricDetails

	// Validator details
	ValidatorDetails map[types.ValidatorPubkey]beacon.ValidatorStatus

	// Internal fields
	log *log.ColorLogger
}

func CreateMetricsCache(
	c *cli.Context,
	cfg *config.StaderNodeConfig,
	ec stader.ExecutionClient,
	bc beacon.Client,
	log *log.ColorLogger,
	slotNumber uint64,
	beaconConfig beacon.Eth2Config,
	nodeAddress common.Address,
) (*MetricsCache, error) {
	prnAddress, err := services.GetPermissionlessNodeRegistryAddress(c)
	if err != nil {
		return nil, err
	}
	ptAddress, err := services.GetPenaltyTrackerAddress(c)
	if err != nil {
		return nil, err
	}
	sdcAddress, err := services.GetSdCollateralAddress(c)
	if err != nil {
		return nil, err
	}
	ethxAddress, err := services.GetEthxTokenAddress(c)
	if err != nil {
		return nil, err
	}
	sdTokenAddress, err := services.GetSdTokenAddress(c)
	if err != nil {
		return nil, err
	}
	stakePoolManagerAddress, err := services.GetStakePoolManagerAddress(c)
	if err != nil {
		return nil, err
	}
	poolUtilsAddress, err := services.GetPoolUtilsAddress(c)
	if err != nil {
		return nil, err
	}
	staderConfigAddress := cfg.GetStaderConfigAddress()
	socializingPoolAddress, err := services.GetSocializingPoolAddress(c)
	if err != nil {
		return nil, err
	}

	prn, err := stader.NewPermissionlessNodeRegistry(ec, prnAddress)
	if err != nil {
		return nil, err
	}
	sdc, err := stader.NewSdCollateralContract(ec, sdcAddress)
	if err != nil {
		return nil, err
	}
	ethx, err := stader.NewErc20TokenContract(ec, ethxAddress)
	if err != nil {
		return nil, err
	}
	sdt, err := stader.NewErc20TokenContract(ec, sdTokenAddress)
	if err != nil {
		return nil, err
	}
	pt, err := stader.NewPenaltyTracker(ec, ptAddress)
	if err != nil {
		return nil, err
	}
	spm, err := stader.NewStakePoolManager(ec, stakePoolManagerAddress)
	if err != nil {
		return nil, err
	}
	putils, err := stader.NewPoolUtils(ec, poolUtilsAddress)
	if err != nil {
		return nil, err
	}
	sdcfg, err := stader.NewStaderConfig(ec, staderConfigAddress)
	if err != nil {
		return nil, err
	}
	sp, err := stader.NewSocializingPool(ec, socializingPoolAddress)
	if err != nil {
		return nil, err
	}

	// Get the execution block for the given slot
	beaconBlock, exists, err := bc.GetBeaconBlock(fmt.Sprintf("%d", slotNumber))
	if err != nil {
		return nil, fmt.Errorf("error getting Beacon block for slot %d: %w", slotNumber, err)
	}
	if !exists {
		return nil, fmt.Errorf("slot %d did not have a Beacon block", slotNumber)
	}

	// Get the corresponding block on the EL
	elBlockNumber := beaconBlock.ExecutionBlockNumber

	// Create the state wrapper
	state := &MetricsCache{
		BeaconSlotNumber: slotNumber,
		ElBlockNumber:    elBlockNumber,
		BeaconConfig:     beaconConfig,
		log:              log,
	}

	state.logLine("Getting network state for EL block %d, Beacon slot %d", elBlockNumber, slotNumber)

	start := time.Now()

	// fetch all validator pub keys
	operatorId, err := node.GetOperatorId(prn, nodeAddress, nil)
	if err != nil {
		return nil, err
	}
	operatorElRewardAddress, err := node.GetNodeElRewardAddress(prn, 1, operatorId, nil)
	if err != nil {
		return nil, err
	}
	elRewardAddressBalance, err := tokens.GetEthBalance(prn.Client, operatorElRewardAddress, nil)
	if err != nil {
		return nil, err
	}
	operatorElRewards, err := pool_utils.CalculateRewardShare(putils, 1, elRewardAddressBalance, nil)
	if err != nil {
		return nil, err
	}
	operatorSdColletaral, err := sd_collateral.GetOperatorSdBalance(sdc, nodeAddress, nil)
	if err != nil {
		return nil, err
	}
	totalValidatorKeys, err := node.GetTotalValidatorKeys(prn, operatorId, nil)
	if err != nil {
		return nil, err
	}
	poolThreshold, err := sd_collateral.GetPoolThreshold(sdc, 1, nil)
	if err != nil {
		return nil, err
	}
	operatorSdCollateralInEth, err := sd_collateral.ConvertSdToEth(sdc, operatorSdColletaral, nil)
	if err != nil {
		return nil, err
	}

	operatorNonTerminalKeys, err := node.GetTotalNonTerminalValidatorKeys(prn, nodeAddress, totalValidatorKeys, nil)
	if err != nil {
		return nil, err
	}
	operatorEthCollateral := float64(4 * operatorNonTerminalKeys)

	nextRewardCycleDetails, err := socializing_pool.GetRewardDetails(sp, nil)
	if err != nil {
		return nil, err
	}

	validatorInfoMap, pubkeys, err := stdr.GetAllValidatorsRegisteredWithOperator(prn, operatorId, nodeAddress, nil)
	if err != nil {
		return nil, err
	}

	activeValidators := big.NewInt(0)
	slashedValidators := big.NewInt(0)
	staderQueuedValidators := big.NewInt(0)
	beaconChainQueuedValidators := big.NewInt(0)
	exitingValidators := big.NewInt(0)
	withdrawnValidators := big.NewInt(0)
	initializedValidators := big.NewInt(0)
	fundsSettledValidators := big.NewInt(0)
	invalidSignatureValidators := big.NewInt(0)
	frontRunValidators := big.NewInt(0)
	totalClRewards := big.NewInt(0)
	cumulativePenalty := big.NewInt(0)

	// Get the validator stats from Beacon
	statusMap, err := bc.GetValidatorStatuses(pubkeys, &beacon.ValidatorStatusOptions{
		Slot: &slotNumber,
	})
	if err != nil {
		return nil, err
	}
	for _, pubKey := range pubkeys {
		totalValidatorPenalty, err := penalty_tracker.GetCumulativeValidatorPenalty(pt, pubKey, nil)
		if err != nil {
			return nil, err
		}
		cumulativePenalty.Add(cumulativePenalty, totalValidatorPenalty)

		validatorContractInfo, ok := validatorInfoMap[pubKey]
		if !ok {
			state.logLine("pub key is not found in validatorInfoMap: %s\n", pubKey)
		}

		status, inBeaconChain := statusMap[pubKey]
		if !inBeaconChain {
			state.logLine("pub key is not found in statusMap: %s\n", pubKey)
		}

		if validatorContractInfo.Status == 0 {
			initializedValidators.Add(initializedValidators, big.NewInt(1))
			continue
		}
		if validatorContractInfo.Status == 1 {
			invalidSignatureValidators.Add(invalidSignatureValidators, big.NewInt(1))
			continue
		}
		if validatorContractInfo.Status == 2 {
			frontRunValidators.Add(frontRunValidators, big.NewInt(1))
			continue
		}
		if !inBeaconChain && validatorContractInfo.Status == 3 {
			staderQueuedValidators.Add(staderQueuedValidators, big.NewInt(1))
			continue
		}
		if !inBeaconChain && validatorContractInfo.Status == 4 {
			staderQueuedValidators.Add(staderQueuedValidators, big.NewInt(1))
			continue
		}
		if validatorContractInfo.Status == 5 {
			fundsSettledValidators.Add(fundsSettledValidators, big.NewInt(1))
			continue
		}
		if inBeaconChain && eth2.IsValidatorExitingButNotWithdrawn(status) {
			exitingValidators.Add(exitingValidators, big.NewInt(1))
			continue
		}
		if inBeaconChain && eth2.IsValidatorWithdrawn(status) {
			withdrawnValidators.Add(withdrawnValidators, big.NewInt(1))
			continue
		}
		if inBeaconChain && eth2.IsValidatorQueued(status) {
			beaconChainQueuedValidators.Add(beaconChainQueuedValidators, big.NewInt(1))
		}
		if inBeaconChain && eth2.IsValidatorSlashed(status) {
			slashedValidators.Add(slashedValidators, big.NewInt(1))
		}
		if inBeaconChain && eth2.IsValidatorActive(status) {
			activeValidators.Add(activeValidators, big.NewInt(1))
		}

		validatorWithdrawVault := validatorInfoMap[pubKey].WithdrawVaultAddress
		withdrawVaultBalance, err := tokens.GetEthBalance(prn.Client, validatorWithdrawVault, nil)
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
		if withdrawVaultRewardShares.OperatorShare.Cmp(rewardsThreshold) > 0 {
			continue
		} else {
			totalClRewards.Add(totalClRewards, withdrawVaultRewardShares.OperatorShare)
		}
	}

	state.ValidatorDetails = statusMap

	state.logLine("Retrieved validator details (total time: %s)", time.Since(start))

	state.logLine("Retrieved Socializing Pool Reward Details")

	start = time.Now()

	rewardClaimData, err := getClaimedAndUnclaimedSocializingSdAndEth(cfg, sp, nodeAddress)
	if err != nil {
		return nil, err
	}

	state.logLine("Retrieved Socializing Pool Reward Details (total time: %s)", time.Since(start))

	state.logLine("Getting Stader Network Details")

	start = time.Now()

	metricsDetails := MetricDetails{}

	sdPrice, err := sd_collateral.ConvertEthToSd(sdc, big.NewInt(1000000000000000000), nil)
	if err != nil {
		return nil, err
	}
	ethPrice, err := sd_collateral.ConvertSdToEth(sdc, big.NewInt(1000000000000000000), nil)
	if err != nil {
		return nil, err
	}
	totalOperators, err := node.GetNextOperatorId(prn, nil)
	if err != nil {
		return nil, err
	}
	totalValidators, err := node.GetNextValidatorId(prn, nil)
	if err != nil {
		return nil, err
	}
	totalActiveValidators, err := node.GetTotalActiveValidators(prn, nil)
	if err != nil {
		return nil, err
	}
	totalQueuedValidators, err := node.GetTotalQueuedValidators(prn, nil)
	if err != nil {
		return nil, err
	}
	totalSdCollateral, err := tokens.BalanceOf(sdt, sdcAddress, nil)
	if err != nil {
		return nil, err
	}
	permissionlessPoolThreshold, err := sd_collateral.GetPoolThreshold(sdc, 1, nil)
	if err != nil {
		return nil, err
	}
	ethxSupply, err := tokens.TotalSupply(ethx, nil)
	if err != nil {
		return nil, err
	}
	totalStakedAssets, err := stake_pool_manager.GetTotalAssets(spm, nil)
	if err != nil {
		return nil, err
	}

	minThreshold := math.RoundDown(eth.WeiToEth(permissionlessPoolThreshold.MinThreshold), 2)
	sdPriceFormatted := math.RoundDown(eth.WeiToEth(sdPrice), 2)
	collateralRatioInSd := minThreshold * sdPriceFormatted

	metricsDetails.SdPrice = sdPriceFormatted
	metricsDetails.EthPrice = math.RoundDown(eth.WeiToEth(ethPrice), 10)
	metricsDetails.OperatorStakedSd = math.RoundDown(eth.WeiToEth(operatorSdColletaral), 10)
	metricsDetails.OperatorStakedSdInEth = math.RoundDown(eth.WeiToEth(operatorSdCollateralInEth), 10)
	metricsDetails.OperatorEthCollateral = operatorEthCollateral
	metricsDetails.TotalOperators = totalOperators.Sub(totalOperators, big.NewInt(1))
	metricsDetails.TotalValidators = totalValidators.Sub(totalValidators, big.NewInt(1))
	metricsDetails.TotalActiveValidators = totalActiveValidators
	metricsDetails.TotalQueuedValidators = totalQueuedValidators
	metricsDetails.TotalStakedSd = math.RoundDown(eth.WeiToEth(totalSdCollateral), 10)
	metricsDetails.TotalEthxSupply = math.RoundDown(eth.WeiToEth(ethxSupply), 10)
	metricsDetails.TotalStakedEthByUsers = totalStakedAssets
	metricsDetails.TotalStakedEthByNos = big.NewInt(0).Mul(totalValidators, big.NewInt(4))
	metricsDetails.CollateralRatio = math.RoundDown(eth.WeiToEth(permissionlessPoolThreshold.MinThreshold), 2)
	metricsDetails.CollateralRatioInSd = collateralRatioInSd
	metricsDetails.MinEthThreshold = math.RoundDown(eth.WeiToEth(poolThreshold.MinThreshold), 4)
	metricsDetails.MaxEthThreshold = math.RoundDown(eth.WeiToEth(poolThreshold.MaxThreshold), 4)

	metricsDetails.ValidatorStatusMap = statusMap
	metricsDetails.ValidatorInfoMap = validatorInfoMap
	metricsDetails.ActiveValidators = activeValidators
	metricsDetails.BeaconChainQueuedValidators = beaconChainQueuedValidators
	metricsDetails.StaderQueuedValidators = staderQueuedValidators
	metricsDetails.ExitingValidators = exitingValidators
	metricsDetails.SlashedValidators = slashedValidators
	metricsDetails.WithdrawnValidators = withdrawnValidators
	metricsDetails.InitializedValidators = initializedValidators
	metricsDetails.FrontRunValidators = frontRunValidators
	metricsDetails.InvalidSignatureValidators = invalidSignatureValidators
	metricsDetails.FundsSettledValidators = fundsSettledValidators
	metricsDetails.CumulativePenalty = math.RoundDown(eth.WeiToEth(cumulativePenalty), 2)
	metricsDetails.UnclaimedClRewards = math.RoundDown(eth.WeiToEth(totalClRewards), 18)
	metricsDetails.NextSocializingPoolRewardCycle = nextRewardCycleDetails
	metricsDetails.UnclaimedNonSocializingPoolElRewards = math.RoundDown(eth.WeiToEth(operatorElRewards.OperatorShare), 2)
	metricsDetails.ClaimedSocializingPoolElRewards = math.RoundDown(eth.WeiToEth(rewardClaimData.claimedEth), 2)
	metricsDetails.ClaimedSocializingPoolSdRewards = math.RoundDown(eth.WeiToEth(rewardClaimData.claimedSd), 2)
	metricsDetails.UnclaimedSocializingPoolElRewards = math.RoundDown(eth.WeiToEth(rewardClaimData.unclaimedEth), 2)
	metricsDetails.UnclaimedSocializingPoolSDRewards = math.RoundDown(eth.WeiToEth(rewardClaimData.unclaimedSd), 2)

	state.StaderNetworkDetails = metricsDetails

	state.logLine("Retrieved Stader Network Details (total time: %s)", time.Since(start))

	return state, nil
}

// Logs a line if the logger is specified
func (s *MetricsCache) logLine(format string, v ...interface{}) {
	if s.log != nil {
		s.log.Printlnf(format, v...)
	}
}

func getClaimedAndUnclaimedSocializingSdAndEth(
	cfg *config.StaderNodeConfig,
	sp *stader.SocializingPoolContractManager,
	nodeAccount common.Address,
) (struct {
	unclaimedEth *big.Int
	unclaimedSd  *big.Int
	claimedEth   *big.Int
	claimedSd    *big.Int
}, error) {
	outstruct := struct {
		unclaimedEth *big.Int
		unclaimedSd  *big.Int
		claimedEth   *big.Int
		claimedSd    *big.Int
	}{
		unclaimedEth: big.NewInt(0),
		unclaimedSd:  big.NewInt(0),
		claimedEth:   big.NewInt(0),
		claimedSd:    big.NewInt(0),
	}

	outstruct.unclaimedEth = big.NewInt(0)
	outstruct.unclaimedSd = big.NewInt(0)
	outstruct.claimedEth = big.NewInt(0)
	outstruct.claimedSd = big.NewInt(0)

	rewardDetails, err := socializing_pool.GetRewardDetails(sp, nil)
	if err != nil {
		return outstruct, err
	}

	unclaimedEth := big.NewInt(0)
	unclaimedSd := big.NewInt(0)
	claimedEth := big.NewInt(0)
	claimedSd := big.NewInt(0)
	for i := int64(1); i < rewardDetails.CurrentIndex.Int64(); i++ {
		cycleMerkleProof, exists, err := cfg.ReadCycleCache(i)
		if err != nil {
			return outstruct, err
		}
		if !exists {
			continue
		}
		claimed, err := socializing_pool.HasClaimedRewards(sp, nodeAccount, big.NewInt(i), nil)
		if err != nil {
			return outstruct, err
		}

		if claimed {
			ethClaimed, ok := big.NewInt(0).SetString(cycleMerkleProof.Eth, 10)
			if !ok {
				return outstruct, fmt.Errorf("failed to parse eth claimed: %s", cycleMerkleProof.Eth)
			}
			sdClaimed, ok := big.NewInt(0).SetString(cycleMerkleProof.Sd, 10)
			if !ok {
				return outstruct, fmt.Errorf("failed to parse sd claimed: %s", cycleMerkleProof.Sd)
			}
			claimedEth.Add(claimedEth, ethClaimed)
			claimedSd.Add(claimedSd, sdClaimed)
		} else {
			ethUnclaimed, ok := big.NewInt(0).SetString(cycleMerkleProof.Eth, 10)
			if !ok {
				return outstruct, fmt.Errorf("failed to parse eth unclaimed: %s", cycleMerkleProof.Eth)
			}
			sdUnclaimed, ok := big.NewInt(0).SetString(cycleMerkleProof.Sd, 10)
			if !ok {
				return outstruct, fmt.Errorf("failed to parse sd unclaimed: %s", cycleMerkleProof.Sd)
			}
			unclaimedEth.Add(unclaimedEth, ethUnclaimed)
			unclaimedSd.Add(unclaimedSd, sdUnclaimed)
		}
	}

	outstruct.unclaimedEth = unclaimedEth
	outstruct.unclaimedSd = unclaimedSd
	outstruct.claimedEth = claimedEth
	outstruct.claimedSd = claimedSd

	return outstruct, nil
}
