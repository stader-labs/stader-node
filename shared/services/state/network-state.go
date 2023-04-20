package state

import (
	"fmt"
	"math/big"
	"time"

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

type NetworkDetails struct {
	SdPrice               *big.Int
	TotalValidators       *big.Int
	TotalOperators        *big.Int
	TotalStakedSd         *big.Int
	TotalStakedEthByNos   *big.Int
	TotalEthxSupply       *big.Int
	TotalStakedEthByUsers *big.Int

	ActiveValidators                  *big.Int
	QueuedValidators                  *big.Int
	SlashedValidators                 *big.Int
	TotalETHBonded                    *big.Int
	TotalSDBonded                     *big.Int
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

	// bcm, err := getBea
	if err != nil {
		return nil, err
	}

	vf, err := stader.NewVaultFactory(ec, cfg.GetVaultFactoryAddress())
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

	activeValidators, err := node.GetActiveValidators(prn, nil)
	if err != nil {
		return nil, err
	}

	queuedValidators, err := node.GetQueuedValidators(prn, nil)
	if err != nil {
		return nil, err
	}

	totalETHBonded, err := node.GetCollateralETH(prn, nil)
	if err != nil {
		return nil, err
	}

	totalSDBonded, err := node.GetCollateralETH(prn, nil)
	if err != nil {
		return nil, err
	}
	operatorElRewardAddress, err := node.GetNodeElRewardAddress(vf, 1, operatorId, nil)
	if err != nil {
		return nil, err
	}
	elRewardBalance, err := tokens.GetEthBalance(ec, operatorElRewardAddress, nil)
	if err != nil {
		return nil, err
	}

	networkDetails.ActiveValidators = activeValidators // OK
	networkDetails.QueuedValidators = queuedValidators // OK

	networkDetails.SlashedValidators = totalETHBonded // CHECK
	networkDetails.TotalETHBonded = totalETHBonded    // CHECK
	networkDetails.TotalSDBonded = totalSDBonded      // CHECK

	networkDetails.SdCollateral = totalSdCollateral                      // OK
	networkDetails.BeaconchainReward = totalSdCollateral                 // CHECK
	networkDetails.ElReward = elRewardBalance                            // OK
	networkDetails.SDReward = totalSdCollateral                          // CHECK
	networkDetails.ETHAPR = totalSdCollateral                            // CHECK
	networkDetails.SDAPR = totalSdCollateral                             // CHECK
	networkDetails.CumulativePenalty = totalSdCollateral                 // CHECK
	networkDetails.ClaimedBeaconchainRewards = totalSdCollateral         // CHECK
	networkDetails.ClaimedELRewards = totalSdCollateral                  // CHECK
	networkDetails.ClaimedSDrewards = totalSdCollateral                  // CHECK
	networkDetails.UnclaimedELRewards = totalSdCollateral                // CHECK
	networkDetails.UnclaimedSDRewards = totalSdCollateral                // CHECK
	networkDetails.NextSDOrELAndSDRewardsCheckpoint = totalSdCollateral  // CHECK
	networkDetails.TotalAttestations = totalSdCollateral                 // CHECK
	networkDetails.AttestationPercent = totalSdCollateral                // CHECK
	networkDetails.BlocksProduced = totalSdCollateral                    // CHECK
	networkDetails.BlocksProducedPercent = totalSdCollateral             // CHECK
	networkDetails.AttestationInclusionEffectiveness = totalSdCollateral // CHECK
	networkDetails.UptimePercent = totalSdCollateral                     // CHECK

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
