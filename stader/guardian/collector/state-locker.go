package collector

import (
	"math/big"
	"sync"

	"github.com/stader-labs/stader-node/stader-lib/contracts"

	"github.com/stader-labs/stader-node/shared/services/beacon"
	"github.com/stader-labs/stader-node/shared/services/state"
	"github.com/stader-labs/stader-node/stader-lib/types"
)

type MetricsCacheContainer struct {
	state *state.MetricsCache

	lock *sync.Mutex
}

func NewMetricsCacheContainer() *MetricsCacheContainer {
	return &MetricsCacheContainer{
		lock: &sync.Mutex{},
		state: &state.MetricsCache{
			StaderNetworkDetails: state.MetricDetails{
				SdPrice:                              0,
				TotalValidators:                      big.NewInt(0),
				TotalOperators:                       big.NewInt(0),
				TotalStakedSd:                        0,
				TotalStakedEthByNos:                  big.NewInt(0),
				TotalEthxSupply:                      0,
				TotalStakedEthByUsers:                big.NewInt(0),
				ActiveValidators:                     big.NewInt(0),
				BeaconChainQueuedValidators:          big.NewInt(0),
				SlashedValidators:                    big.NewInt(0),
				ExitingValidators:                    big.NewInt(0),
				WithdrawnValidators:                  big.NewInt(0),
				InitializedValidators:                big.NewInt(0),
				FrontRunValidators:                   big.NewInt(0),
				InvalidSignatureValidators:           big.NewInt(0),
				FundsSettledValidators:               big.NewInt(0),
				CumulativePenalty:                    0,
				UnclaimedClRewards:                   0,
				UnclaimedNonSocializingPoolElRewards: 0,
				ClaimedSocializingPoolElRewards:      0,
				ClaimedSocializingPoolSdRewards:      0,
				UnclaimedSocializingPoolElRewards:    0,
				TotalActiveValidators:                big.NewInt(0),
				TotalQueuedValidators:                0,
				UnclaimedSocializingPoolSDRewards:    0,
				OperatorStakedSd:                     0,
				OperatorEthCollateral:                0,
				NextSocializingPoolRewardCycle: types.RewardCycleDetails{
					CurrentIndex:      big.NewInt(0),
					CurrentStartBlock: big.NewInt(0),
					CurrentEndBlock:   big.NewInt(0),
				},
				ValidatorStatusMap:            make(map[types.ValidatorPubkey]beacon.ValidatorStatus),
				ValidatorInfoMap:              make(map[types.ValidatorPubkey]contracts.Validator),
				CollateralRatio:               0,
				CollateralRatioInSd:           0,
				OperatorSDUtilized:            0,
				OperatorSDInterest:            0,
				OperatorSDUtilizationPosition: 0,
				SdUtilityPoolBalance:          0,
				OperatorSDSelfBond:            0,
				SdCollateralPct:               0,
				LockedEth:                     0,
				HealthFactor:                  0,
				LiquidationStatus:             0,
				StaderQueuedValidators:        big.NewInt(0),
			},
		},
	}
}

func (l *MetricsCacheContainer) UpdateMetricsContainer(state *state.MetricsCache) {
	l.lock.Lock()
	defer l.lock.Unlock()
	l.state = state
}

func (l *MetricsCacheContainer) GetMetricsContainer() *state.MetricsCache {
	l.lock.Lock()
	defer l.lock.Unlock()
	return l.state
}
