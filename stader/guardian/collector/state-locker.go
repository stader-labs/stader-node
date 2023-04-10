package collector

import (
	"github.com/stader-labs/stader-node/shared/services/state"
	"sync"
)

type StateCache struct {
	state *state.NetworkStateCache

	lock *sync.Mutex
}

func NewStateCache() *StateCache {
	return &StateCache{
		lock: &sync.Mutex{},
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
