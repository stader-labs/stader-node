package state

import (
	"fmt"
	"github.com/stader-labs/stader-node/shared/utils/eth2"
	"math/big"
	"time"

	"github.com/stader-labs/stader-node/shared/services/beacon"
	"github.com/stader-labs/stader-node/shared/services/config"
	"github.com/stader-labs/stader-node/shared/utils/log"
	"github.com/stader-labs/stader-node/stader-lib/node"
	sd_collateral "github.com/stader-labs/stader-node/stader-lib/sd-collateral"

	"github.com/stader-labs/stader-node/stader-lib/contracts"
	"github.com/stader-labs/stader-node/stader-lib/stader"
	"github.com/stader-labs/stader-node/stader-lib/tokens"
	"github.com/stader-labs/stader-node/stader-lib/types"

	"github.com/ethereum/go-ethereum/common"
)

type NetworkDetails struct {
	// Network details
	SdPrice               *big.Int
	TotalValidators       *big.Int
	TotalOperators        *big.Int
	TotalStakedSd         *big.Int
	TotalStakedEthByNos   *big.Int
	TotalEthxSupply       *big.Int
	TotalStakedEthByUsers *big.Int

	// Validator specific info
	ActiveValidators    *big.Int
	QueuedValidators    *big.Int
	SlashedValidators   *big.Int
	ExitingValidators   *big.Int
	WithdrawnValidators *big.Int
	ValidatorStatusMap  map[types.ValidatorPubkey]beacon.ValidatorStatus

	TotalETHBonded                    *big.Int
	SdCollateral                      *big.Int
	BeaconchainReward                 *big.Int
	ElReward                          *big.Int
	SDReward                          *big.Int
	ETHAPR                            *big.Int
	SDAPR                             *big.Int
	CumulativePenalty                 *big.Int
	ClaimedBeaconchainRewards         *big.Int
	ClaimedELRewards                  *big.Int
	ClaimedSDrewards                  *big.Int
	UnclaimedELRewards                *big.Int
	UnclaimedSDRewards                *big.Int
	NextSDOrELAndSDRewardsCheckpoint  *big.Int
	TotalAttestations                 *big.Int
	AttestationPercent                *big.Int
	BlocksProduced                    *big.Int
	BlocksProducedPercent             *big.Int
	AttestationInclusionEffectiveness *big.Int
	UptimePercent                     *big.Int
}

type NetworkStateCache struct {
	// Block / slot for this state
	ElBlockNumber    uint64
	BeaconSlotNumber uint64
	BeaconConfig     beacon.Eth2Config

	StaderNetworkDetails NetworkDetails

	// Validator details
	ValidatorDetails map[types.ValidatorPubkey]beacon.ValidatorStatus

	// Internal fields
	log *log.ColorLogger
}

func CreateNetworkStateCache(
	cfg *config.StaderNodeConfig,
	ec stader.ExecutionClient,
	bc beacon.Client,
	log *log.ColorLogger,
	slotNumber uint64,
	beaconConfig beacon.Eth2Config,
	nodeAddress common.Address,
) (*NetworkStateCache, error) {
	prnAddress := cfg.GetPermissionlessNodeRegistryAddress()
	sdcAddress := cfg.GetSdCollateralContractAddress()
	ethxAddress := cfg.GetEthxTokenAddress()

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
	state := &NetworkStateCache{
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

	totalValidatorKeys, err := node.GetTotalValidatorKeys(prn, operatorId, nil)
	if err != nil {
		return nil, err
	}

	pubkeys := make([]types.ValidatorPubkey, 0, totalValidatorKeys.Int64())
	for i := 0; i < int(totalValidatorKeys.Int64()); i++ {
		validatorId, err := node.GetValidatorIdByOperatorId(prn, operatorId, big.NewInt(int64(i)), nil)
		if err != nil {
			return nil, err
		}
		validatorInfo, err := node.GetValidatorInfo(prn, validatorId, nil)
		if err != nil {
			return nil, err
		}

		pubkeys = append(pubkeys, types.BytesToValidatorPubkey(validatorInfo.Pubkey))
	}

	// Get the validator stats from Beacon
	statusMap, err := bc.GetValidatorStatuses(pubkeys, &beacon.ValidatorStatusOptions{
		Slot: &slotNumber,
	})
	if err != nil {
		return nil, err
	}
	state.ValidatorDetails = statusMap
	activeValidators := big.NewInt(0)
	slashedValidators := big.NewInt(0)
	queuedValidators := big.NewInt(0)
	exitingValidators := big.NewInt(0)
	withdrawnValidators := big.NewInt(0)
	for _, status := range statusMap {
		if eth2.IsValidatorQueued(status) {
			queuedValidators.Add(queuedValidators, big.NewInt(1))
		}
		if eth2.IsValidatorSlashed(status) {
			slashedValidators.Add(slashedValidators, big.NewInt(1))
		}
		if eth2.IsValidatorExiting(status) {
			exitingValidators.Add(exitingValidators, big.NewInt(1))
		}
		if eth2.IsValidatorWithdrawn(status) {
			withdrawnValidators.Add(withdrawnValidators, big.NewInt(1))
		}
		if eth2.IsValidatorActive(status) {
			activeValidators.Add(activeValidators, big.NewInt(1))
		}
	}

	state.logLine("Retrieved validator details (total time: %s)", time.Since(start))

	state.logLine("Getting Stader Network Details")

	//// TODO - bchain - parallelize these calls
	start = time.Now()

	networkDetails := NetworkDetails{}

	sdPrice, err := sd_collateral.ConvertEthToSd(sdc, big.NewInt(1), nil)
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
	totalSdCollateral, err := sd_collateral.GetTotalSdCollateral(sdc, nil)
	if err != nil {
		return nil, err
	}
	ethxSupply, err := tokens.TotalSupply(ethx, nil)
	if err != nil {
		return nil, err
	}

	networkDetails.SdPrice = sdPrice
	networkDetails.TotalOperators = totalOperators
	networkDetails.TotalValidators = totalValidators
	networkDetails.TotalStakedSd = totalSdCollateral
	networkDetails.TotalEthxSupply = ethxSupply
	networkDetails.TotalStakedEthByUsers = big.NewInt(0)
	networkDetails.TotalStakedEthByNos = big.NewInt(0).Mul(totalValidators, big.NewInt(4))

	networkDetails.ValidatorStatusMap = statusMap
	networkDetails.ActiveValidators = activeValidators
	networkDetails.QueuedValidators = queuedValidators
	networkDetails.ExitingValidators = exitingValidators
	networkDetails.SlashedValidators = slashedValidators
	networkDetails.WithdrawnValidators = withdrawnValidators

	state.StaderNetworkDetails = networkDetails

	state.logLine("Retrieved Stader Network Details (total time: %s)", time.Since(start))

	return state, nil
}

// Logs a line if the logger is specified
func (s *NetworkStateCache) logLine(format string, v ...interface{}) {
	if s.log != nil {
		s.log.Printlnf(format, v...)
	}
}

func spInfo(sp *stader.SocializingPoolContractManager) (struct {
	ethRewardRemaining *big.Int
	sdRewardRemaining  *big.Int
	rewardChan         <-chan *contracts.SocializingPoolOperatorRewardsClaimed
}, error) {

	outstruct := new(struct {
		ethRewardRemaining *big.Int
		sdRewardRemaining  *big.Int
		rewardChan         <-chan *contracts.SocializingPoolOperatorRewardsClaimed
	})
	ethRewardRemaining, err := sp.SocializingPool.TotalOperatorETHRewardsRemaining(nil)
	if err != nil {
		return *outstruct, err
	}

	sdRewardRemaining, err := sp.SocializingPool.TotalOperatorSDRewardsRemaining(nil)
	if err != nil {
		return *outstruct, err
	}
	sink := make(chan *contracts.SocializingPoolOperatorRewardsClaimed)
	// var recipient []common.Address
	// _, err = sp.SocializingPool.WatchOperatorRewardsClaimed(nil, sink, recipient)
	// if err != nil {
	// 	return *outstruct, err
	// }

	outstruct.ethRewardRemaining = ethRewardRemaining
	outstruct.sdRewardRemaining = sdRewardRemaining
	outstruct.rewardChan = sink

	return *outstruct, nil
}
