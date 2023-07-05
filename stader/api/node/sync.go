/*
This work is licensed and released under GNU GPL v3 or any other later versions.
The full text of the license is below/ found at <http://www.gnu.org/licenses/>

(c) 2023 Rocket Pool Pty Ltd. Modified under GNU GPL v3. [1.2.0]

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
package node

import (
	"github.com/urfave/cli"

	"github.com/stader-labs/stader-node/shared/services"
	"github.com/stader-labs/stader-node/shared/types/api"
)

func getSyncProgress(c *cli.Context) (*api.NodeSyncProgressResponse, error) {

	// Response
	response := api.NodeSyncProgressResponse{}

	cfg, err := services.GetConfig(c)
	if err != nil {
		return nil, err
	}

	// Get the EC manager
	ecMgr, err := services.GetEthClient(c)
	if err != nil {
		return nil, err
	}

	// Get the status of the EC and fallback EC
	ecStatus := ecMgr.CheckStatus(cfg)
	response.EcStatus = *ecStatus

	// Get the BC manager
	bcMgr, err := services.GetBeaconClient(c)
	if err != nil {
		return nil, err
	}

	// Get the status of the BC and fallback BC
	bcStatus := bcMgr.CheckStatus()
	response.BcStatus = *bcStatus

	// Return response
	return &response, nil

}
