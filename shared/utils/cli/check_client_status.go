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
package cli

import (
	"fmt"

	"github.com/stader-labs/stader-node/shared/services/stader"
	"github.com/stader-labs/stader-node/shared/types/api"
)

// Check the status of the Execution and Consensus client(s) and provision the API with them
func CheckClientStatus(staderClient *stader.Client) error {

	// Check if the primary clients are up, synced, and able to respond to requests - if not, forces the use of the fallbacks for this command
	response, err := staderClient.GetClientStatus()
	if err != nil {
		return err
	}

	ecMgrStatus := response.EcManagerStatus
	bcMgrStatus := response.BcManagerStatus

	// Primary EC and CC are good
	if ecMgrStatus.PrimaryClientStatus.IsSynced && bcMgrStatus.PrimaryClientStatus.IsSynced {
		staderClient.SetClientStatusFlags(true, false)
		return nil
	}

	// Get the status messages
	primaryEcStatus := getClientStatusString(ecMgrStatus.PrimaryClientStatus)
	primaryBcStatus := getClientStatusString(bcMgrStatus.PrimaryClientStatus)
	fallbackEcStatus := getClientStatusString(ecMgrStatus.FallbackClientStatus)
	fallbackBcStatus := getClientStatusString(bcMgrStatus.FallbackClientStatus)

	// Check the fallbacks if enabled
	if ecMgrStatus.FallbackEnabled && bcMgrStatus.FallbackEnabled {

		// Fallback EC and CC are good
		if ecMgrStatus.FallbackClientStatus.IsSynced && bcMgrStatus.FallbackClientStatus.IsSynced {
			fmt.Printf("%sNOTE: primary clients are not ready, using fallback clients...\n\tPrimary EC status: %s\n\tPrimary CC status: %s%s\n\n", colorYellow, primaryEcStatus, primaryBcStatus, colorReset)
			staderClient.SetClientStatusFlags(true, true)
			return nil
		}

		// Both pairs aren't ready
		return fmt.Errorf("Error: neither primary nor fallback client pairs are ready.\n\tPrimary EC status: %s\n\tFallback EC status: %s\n\tPrimary CC status: %s\n\tFallback CC status: %s", primaryEcStatus, fallbackEcStatus, primaryBcStatus, fallbackBcStatus)

	}

	// Primary isn't ready and fallback isn't enabled
	return fmt.Errorf("Error: primary client pair isn't ready and fallback clients aren't enabled.\n\tPrimary EC status: %s\n\tPrimary CC status: %s", primaryEcStatus, primaryBcStatus)

}

func getClientStatusString(clientStatus api.ClientStatus) string {
	if clientStatus.IsSynced {
		return "synced and ready"
	} else if clientStatus.IsWorking {
		return fmt.Sprintf("syncing (%.2f%%)", clientStatus.SyncProgress*100)
	} else {
		return fmt.Sprintf("unavailable (%s)", clientStatus.Error)
	}
}
