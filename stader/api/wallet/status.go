/*
This work is licensed and released under GNU GPL v3 or any other later versions.
The full text of the license is below/ found at <http://www.gnu.org/licenses/>

(c) 2023 Rocket Pool Pty Ltd. Modified under GNU GPL v3. [0.4.0-beta]

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
package wallet

import (
	"context"
	"github.com/urfave/cli"
	"math/big"

	"github.com/stader-labs/stader-node/shared/services"
	"github.com/stader-labs/stader-node/shared/types/api"
)

func getStatus(c *cli.Context) (*api.WalletStatusResponse, error) {

	// Get services
	pm, err := services.GetPasswordManager(c)
	if err != nil {
		return nil, err
	}
	w, err := services.GetWallet(c)
	if err != nil {
		return nil, err
	}
	ec, err := services.GetEthClient(c)
	if err != nil {
		return nil, err
	}

	// Response
	response := api.WalletStatusResponse{}

	// Get wallet status
	response.PasswordSet = pm.IsPasswordSet()
	response.WalletInitialized = w.IsInitialized()

	// Get accounts if initialized
	if response.WalletInitialized {

		// Get node account
		nodeAccount, err := w.GetNodeAccount()
		if err != nil {
			return nil, err
		}
		response.AccountAddress = nodeAccount.Address

		currentBlockNumber, err := ec.BlockNumber(context.Background())
		if err != nil {
			return nil, err
		}

		currentNonce, err := ec.NonceAt(context.Background(), nodeAccount.Address, big.NewInt(int64(currentBlockNumber)))
		if err != nil {
			return nil, err
		}
		pendingNonce, err := ec.PendingNonceAt(context.Background(), nodeAccount.Address)
		if err != nil {
			return nil, err
		}

		response.PendingNonce = big.NewInt(int64(pendingNonce))
		response.CurrentNonce = big.NewInt(int64(currentNonce))
	}

	// Return response
	return &response, nil

}
