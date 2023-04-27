package collector

import (
	"math/big"
	"sync"

	"github.com/stader-labs/stader-node/shared/services/beacon"
	"github.com/stader-labs/stader-node/shared/services/state"
	"github.com/stader-labs/stader-node/stader-lib/types"
)

type StateCache struct {
	state *state.NetworkStateCache

	lock *sync.Mutex
}

func NewStateCache() *StateCache {
	return &StateCache{
		lock: &sync.Mutex{},
		state: &state.NetworkStateCache{
			StaderNetworkDetails: state.NetworkDetails{
				SdPrice:                              big.NewInt(0),
				TotalValidators:                      big.NewInt(0),
				TotalOperators:                       big.NewInt(0),
				TotalStakedSd:                        0,
				TotalStakedEthByNos:                  big.NewInt(0),
				TotalEthxSupply:                      0,
				TotalStakedEthByUsers:                big.NewInt(0),
				ActiveValidators:                     big.NewInt(0),
				QueuedValidators:                     big.NewInt(0),
				SlashedValidators:                    big.NewInt(0),
				ExitingValidators:                    big.NewInt(0),
				WithdrawnValidators:                  big.NewInt(0),
				CumulativePenalty:                    0,
				UnclaimedClRewards:                   0,
				UnclaimedNonSocializingPoolElRewards: 0,
				ClaimedSocializingPoolElRewards:      0,
				ClaimedSocializingPoolSdRewards:      0,
				UnclaimedSocializingPoolElRewards:    0,
				TotalActiveValidators:                big.NewInt(0),
				TotalQueuedValidators:                big.NewInt(0),
				UnclaimedSocializingPoolSDRewards:    0,
				OperatorStakedSd:                     0,
				OperatorEthCollateral:                0,
				NextSocializingPoolRewardCycle: types.RewardCycleDetails{
					CurrentIndex:      big.NewInt(0),
					CurrentStartBlock: big.NewInt(0),
					CurrentEndBlock:   big.NewInt(0),
					NextIndex:         big.NewInt(0),
					NextStartBlock:    big.NewInt(0),
					NextEndBlock:      big.NewInt(0),
				},
				ValidatorStatusMap: make(map[types.ValidatorPubkey]beacon.ValidatorStatus),
				ValidatorInfoMap:   make(map[types.ValidatorPubkey]types.ValidatorContractInfo),
				CollateralRatio:    0,
			},
		},
	}
}

func (l *StateCache) UpdateState(state *state.NetworkStateCache) {
	l.lock.Lock()
	defer l.lock.Unlock()
	l.state = state
}

func (l *StateCache) GetState() *state.NetworkStateCache {
	l.lock.Lock()
	defer l.lock.Unlock()
	return l.state
}
