package collector

import (
	"math/big"
	"sync"

	"github.com/stader-labs/stader-node/shared/services/state"
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
				TotalStakedSd:                        big.NewInt(0),
				TotalStakedEthByNos:                  big.NewInt(0),
				TotalEthxSupply:                      big.NewInt(0),
				TotalStakedEthByUsers:                big.NewInt(0),
				ActiveValidators:                     big.NewInt(0),
				QueuedValidators:                     big.NewInt(0),
				SlashedValidators:                    big.NewInt(0),
				ExitingValidators:                    big.NewInt(0),
				WithdrawnValidators:                  big.NewInt(0),
				CumulativePenalty:                    big.NewInt(0),
				UnclaimedClRewards:                   big.NewInt(0),
				UnclaimedNonSocializingPoolElRewards: big.NewInt(0),
				ClaimedSocializingPoolElRewards:      big.NewInt(0),
				ClaimedSocializingPoolSdRewards:      big.NewInt(0),
				UnclaimedSocializingPoolElRewards:    big.NewInt(0),
				UnclaimedSocializingPoolSDRewards:    big.NewInt(0),
				EthApr:                               big.NewInt(0),
				SdApr:                                big.NewInt(0),
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
