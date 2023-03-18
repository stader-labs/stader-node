package eth2

// ROCKETPOOL-OWNED

import (
	"github.com/stader-labs/stader-node/shared/services/beacon"
)

// Get an eth2 epoch number by time
func EpochAt(config beacon.Eth2Config, time uint64) uint64 {
	return config.GenesisEpoch + (time-config.GenesisTime)/config.SecondsPerEpoch
}
