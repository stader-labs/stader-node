/*
This work is licensed and released under GNU GPL v3 or any other later versions.
The full text of the license is below/ found at <http://www.gnu.org/licenses/>

(c) 2023 Rocket Pool Pty Ltd. Modified under GNU GPL v3. [0.3.0-beta]

This program is free software: you can redistribute it and/or modify
it under the terms of the GNU General Public License as published by
the Free Software Foundation, either version 3 of the License, or
(at your option) any later version.

This program is distributed in the hope that it will be useful,
but WITHOUT ANY WARRANTY; without even the implied warranty of
MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
GNU General Public License for more details.

You should have received a copy of the GNU General Public License
along with this program.  If not, see <http://www.gnu.org/licenses/>.
*/
package eth2

import (
	"github.com/stader-labs/stader-node/shared/services/beacon"
)

// Get an eth2 epoch number by time
func EpochAt(config beacon.Eth2Config, time uint64) uint64 {
	return config.GenesisEpoch + (time-config.GenesisTime)/config.SecondsPerEpoch
}

func IsValidatorExiting(validatorStatus beacon.ValidatorStatus) bool {
	//fmt.Printf("Validator status: %v", validatorStatus)
	switch validatorStatus.Status {
	case beacon.ValidatorState_PendingInitialized:
		return false
	case beacon.ValidatorState_PendingQueued:
		return false
	case beacon.ValidatorState_ActiveOngoing:
		return false
	case beacon.ValidatorState_ActiveSlashed:
		return false
	}

	return true
}

func IsValidatorActive(validatorStatus beacon.ValidatorStatus) bool {
	switch validatorStatus.Status {
	case beacon.ValidatorState_ActiveOngoing:
		return true
	case beacon.ValidatorState_ActiveSlashed:
		return true
	}

	return false
}
