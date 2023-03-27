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
package wallet

import (
	"errors"
	"fmt"

	"github.com/ethereum/go-ethereum/common"
	"github.com/urfave/cli"

	"github.com/stader-labs/stader-node/shared/services"
	"github.com/stader-labs/stader-node/shared/services/wallet"
	"github.com/stader-labs/stader-node/shared/types/api"
)

const (
	findIterations uint = 100000
)

func recoverWallet(c *cli.Context, mnemonic string) (*api.RecoverWalletResponse, error) {

	// Get services
	if err := services.RequireNodePassword(c); err != nil {
		return nil, err
	}
	w, err := services.GetWallet(c)
	if err != nil {
		return nil, err
	}

	// Response
	response := api.RecoverWalletResponse{}

	// Check if wallet is already initialized
	if w.IsInitialized() {
		return nil, errors.New("the wallet is already initialized")
	}

	// Get the derivation path
	path := c.String("derivation-path")
	switch path {
	case "":
		path = wallet.DefaultNodeKeyPath
	case "ledgerLive":
		path = wallet.LedgerLiveNodeKeyPath
	case "mew":
		path = wallet.MyEtherWalletNodeKeyPath
	}

	// Get the wallet index
	walletIndex := c.Uint("wallet-index")

	// Recover wallet
	if err := w.Recover(path, walletIndex, mnemonic); err != nil {
		return nil, err
	}

	// Get node account
	nodeAccount, err := w.GetNodeAccount()
	if err != nil {
		return nil, err
	}
	response.AccountAddress = nodeAccount.Address

	// Save wallet
	if err := w.Save(); err != nil {
		return nil, err
	}

	// Return response
	return &response, nil

}

func searchAndRecoverWallet(c *cli.Context, mnemonic string, address common.Address) (*api.SearchAndRecoverWalletResponse, error) {

	// Get services
	if err := services.RequireNodePassword(c); err != nil {
		return nil, err
	}
	w, err := services.GetWallet(c)
	if err != nil {
		return nil, err
	}

	// Response
	response := api.SearchAndRecoverWalletResponse{}

	// Check if wallet is already initialized
	if w.IsInitialized() {
		return nil, errors.New("the wallet is already initialized")
	}

	// Try each derivation path across all of the iterations
	paths := []string{
		wallet.DefaultNodeKeyPath,
		wallet.LedgerLiveNodeKeyPath,
		wallet.MyEtherWalletNodeKeyPath,
	}
	for i := uint(0); i < findIterations; i++ {
		for j := 0; j < len(paths); j++ {
			derivationPath := paths[j]
			recoveredWallet, err := wallet.NewWallet("", uint(w.GetChainID().Uint64()), nil, nil, 0, nil)
			if err != nil {
				return nil, fmt.Errorf("error generating new wallet: %w", err)
			}
			err = recoveredWallet.TestRecovery(derivationPath, i, mnemonic)
			if err != nil {
				return nil, fmt.Errorf("error recovering wallet with path [%s], index [%d]: %w", derivationPath, i, err)
			}

			// Get recovered account
			recoveredAccount, err := recoveredWallet.GetNodeAccount()
			if err != nil {
				return nil, fmt.Errorf("error getting recovered account: %w", err)
			}
			if recoveredAccount.Address == address {
				// We found the correct derivation path and index
				response.FoundWallet = true
				response.DerivationPath = derivationPath
				response.Index = i
				break
			}
		}
		if response.FoundWallet {
			break
		}
	}

	if !response.FoundWallet {
		return nil, fmt.Errorf("exhausted all derivation paths and indices from 0 to %d, wallet not found", findIterations)
	}

	// Recover wallet
	if err := w.Recover(response.DerivationPath, response.Index, mnemonic); err != nil {
		return nil, err
	}

	// Get node account
	nodeAccount, err := w.GetNodeAccount()
	if err != nil {
		return nil, err
	}
	response.AccountAddress = nodeAccount.Address

	// Save wallet
	if err := w.Save(); err != nil {
		return nil, err
	}

	// Return response
	return &response, nil

}
