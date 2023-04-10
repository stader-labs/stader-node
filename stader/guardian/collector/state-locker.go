package collector

import (
	"github.com/stader-labs/stader-node/shared/services/state"
	"sync"
)

type StateLocker struct {
	state *state.NetworkState

	lock *sync.Mutex
}

func NewStateLocker() *StateLocker {
	return &StateLocker{
		lock: &sync.Mutex{},
	}
}

func (l *StateLocker) UpdateState(state *state.NetworkState) {
	l.lock.Lock()
	defer l.lock.Unlock()
	l.state = state
}

func (l *StateLocker) GetState() *state.NetworkState {
	l.lock.Lock()
	defer l.lock.Unlock()
	return l.state
}
