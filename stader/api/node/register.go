/*
This work is licensed and released under GNU GPL v3 or any other later versions. 
The full text of the license is below/ found at <http://www.gnu.org/licenses/>

(c) 2023 Rocket Pool Pty Ltd. Modified under GNU GPL v3.

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
	"fmt"
	"github.com/ethereum/go-ethereum/common"

	"github.com/stader-labs/stader-node/shared/services"
	"github.com/stader-labs/stader-node/shared/types/api"
	"github.com/stader-labs/stader-node/shared/utils/eth1"
	"github.com/stader-labs/stader-node/stader-lib/node"
	"github.com/urfave/cli"
)

func canRegisterNode(c *cli.Context, operatorName string, operatorRewardAddress common.Address, socializeMev bool) (*api.CanRegisterNodeResponse, error) {

	// Get services
	if err := services.RequireNodeWallet(c); err != nil {
		return nil, err
	}
	_, err := services.GetWallet(c)
	if err != nil {
		return nil, err
	}
	pnr, err := services.GetPermissionlessNodeRegistry(c)
	if err != nil {
		return nil, err
	}
	w, err := services.GetWallet(c)
	if err != nil {
		return nil, err
	}

	// Response
	response := api.CanRegisterNodeResponse{}

	nodeAccount, err := w.GetNodeAccount()

	operatorId, err := node.GetOperatorId(pnr, nodeAccount.Address, nil)
	if err != nil {
		return nil, err
	}

	if operatorId.Int64() != 0 {
		response.AlreadyRegistered = true
		return &response, nil
	}

	opts, err := w.GetNodeAccountTransactor()
	if err != nil {
		return nil, err
	}

	isPermissionlessRegistryPaused, err := node.IsPermissionlessNodeRegistryPaused(pnr, nil)
	if err != nil {
		return nil, err
	}
	if isPermissionlessRegistryPaused {
		response.RegistrationPaused = true
		return &response, nil
	}

	gasInfo, err := node.EstimateOnboardNodeOperator(pnr, socializeMev, operatorName, operatorRewardAddress, opts)
	if err != nil {
		return nil, err
	}

	response.CanRegister = true
	response.GasInfo = gasInfo

	return &response, nil

}

func registerNode(c *cli.Context, operatorName string, operatorRewardAddress common.Address, mevSocialize bool) (*api.RegisterNodeResponse, error) {

	// Get services
	w, err := services.GetWallet(c)
	if err != nil {
		return nil, err
	}
	prn, err := services.GetPermissionlessNodeRegistry(c)

	// Response
	response := api.RegisterNodeResponse{}

	// Get transactor
	opts, err := w.GetNodeAccountTransactor()
	if err != nil {
		return nil, err
	}

	// Override the provided pending TX if requested
	err = eth1.CheckForNonceOverride(c, opts)
	if err != nil {
		return nil, fmt.Errorf("Error checking for nonce override: %w", err)
	}

	// Register node
	tx, err := node.OnboardNodeOperator(prn, mevSocialize, operatorName, operatorRewardAddress, opts)
	if err != nil {
		return nil, err
	}
	response.TxHash = tx.Hash()

	// Return response
	return &response, nil

}
