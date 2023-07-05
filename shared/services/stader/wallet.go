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
package stader

import (
	"encoding/json"
	"fmt"

	"github.com/ethereum/go-ethereum/common"
	"github.com/stader-labs/stader-node/shared/types/api"
)

// Get wallet status
func (c *Client) WalletStatus() (api.WalletStatusResponse, error) {
	responseBytes, err := c.callAPI("wallet status")
	if err != nil {
		return api.WalletStatusResponse{}, fmt.Errorf("Could not get wallet status: %w", err)
	}
	var response api.WalletStatusResponse
	if err := json.Unmarshal(responseBytes, &response); err != nil {
		return api.WalletStatusResponse{}, fmt.Errorf("Could not decode wallet status response: %w", err)
	}
	if response.Error != "" {
		return api.WalletStatusResponse{}, fmt.Errorf("Could not get wallet status: %s", response.Error)
	}
	return response, nil
}

// Set wallet password
func (c *Client) SetPassword(password string) (api.SetPasswordResponse, error) {
	responseBytes, err := c.callAPI("wallet set-password", password)
	if err != nil {
		return api.SetPasswordResponse{}, fmt.Errorf("Could not set wallet password: %w", err)
	}
	var response api.SetPasswordResponse
	if err := json.Unmarshal(responseBytes, &response); err != nil {
		return api.SetPasswordResponse{}, fmt.Errorf("Could not decode set wallet password response: %w", err)
	}
	if response.Error != "" {
		return api.SetPasswordResponse{}, fmt.Errorf("Could not set wallet password: %s", response.Error)
	}
	return response, nil
}

// Initialize wallet
func (c *Client) InitWallet(derivationPath string) (api.InitWalletResponse, error) {
	responseBytes, err := c.callAPI("wallet init --derivation-path", derivationPath)
	if err != nil {
		return api.InitWalletResponse{}, fmt.Errorf("Could not initialize wallet: %w", err)
	}
	var response api.InitWalletResponse
	if err := json.Unmarshal(responseBytes, &response); err != nil {
		return api.InitWalletResponse{}, fmt.Errorf("Could not decode initialize wallet response: %w", err)
	}
	if response.Error != "" {
		return api.InitWalletResponse{}, fmt.Errorf("Could not initialize wallet: %s", response.Error)
	}
	return response, nil
}

// Recover wallet
func (c *Client) RecoverWallet(mnemonic string, skipValidatorKeyRecovery bool, derivationPath string, walletIndex uint) (api.RecoverWalletResponse, error) {
	command := "wallet recover "
	if skipValidatorKeyRecovery {
		command += "--skip-validator-key-recovery "
	}
	if walletIndex != 0 {
		command += fmt.Sprintf("--wallet-index %d ", walletIndex)
	}
	command += "--derivation-path"

	responseBytes, err := c.callAPI(command, derivationPath, mnemonic)
	if err != nil {
		return api.RecoverWalletResponse{}, fmt.Errorf("Could not recover wallet: %w", err)
	}
	var response api.RecoverWalletResponse
	if err := json.Unmarshal(responseBytes, &response); err != nil {
		return api.RecoverWalletResponse{}, fmt.Errorf("Could not decode recover wallet response: %w", err)
	}
	if response.Error != "" {
		return api.RecoverWalletResponse{}, fmt.Errorf("Could not recover wallet: %s", response.Error)
	}
	return response, nil
}

// Search and recover wallet
func (c *Client) SearchAndRecoverWallet(mnemonic string, address common.Address, skipValidatorKeyRecovery bool) (api.SearchAndRecoverWalletResponse, error) {
	command := "wallet search-and-recover "
	if skipValidatorKeyRecovery {
		command += "--skip-validator-key-recovery "
	}

	responseBytes, err := c.callAPI(command, mnemonic, address.Hex())
	if err != nil {
		return api.SearchAndRecoverWalletResponse{}, fmt.Errorf("Could not search and recover wallet: %w", err)
	}
	var response api.SearchAndRecoverWalletResponse
	if err := json.Unmarshal(responseBytes, &response); err != nil {
		return api.SearchAndRecoverWalletResponse{}, fmt.Errorf("Could not decode search-and-recover wallet response: %w", err)
	}
	if response.Error != "" {
		return api.SearchAndRecoverWalletResponse{}, fmt.Errorf("Could not search and recover wallet: %s", response.Error)
	}
	return response, nil
}

// Recover wallet
func (c *Client) TestRecoverWallet(mnemonic string, skipValidatorKeyRecovery bool, derivationPath string, walletIndex uint) (api.RecoverWalletResponse, error) {
	command := "wallet test-recovery "
	if skipValidatorKeyRecovery {
		command += "--skip-validator-key-recovery "
	}
	if walletIndex != 0 {
		command += fmt.Sprintf("--wallet-index %d ", walletIndex)
	}
	command += "--derivation-path"

	responseBytes, err := c.callAPI(command, derivationPath, mnemonic)
	if err != nil {
		return api.RecoverWalletResponse{}, fmt.Errorf("Could not test recover wallet: %w", err)
	}
	var response api.RecoverWalletResponse
	if err := json.Unmarshal(responseBytes, &response); err != nil {
		return api.RecoverWalletResponse{}, fmt.Errorf("Could not decode test recover wallet response: %w", err)
	}
	if response.Error != "" {
		return api.RecoverWalletResponse{}, fmt.Errorf("Could not test recover wallet: %s", response.Error)
	}
	return response, nil
}

// Search and recover wallet
func (c *Client) TestSearchAndRecoverWallet(mnemonic string, address common.Address, skipValidatorKeyRecovery bool) (api.SearchAndRecoverWalletResponse, error) {
	command := "wallet test-search-and-recover "
	if skipValidatorKeyRecovery {
		command += "--skip-validator-key-recovery "
	}

	responseBytes, err := c.callAPI(command, mnemonic, address.Hex())
	if err != nil {
		return api.SearchAndRecoverWalletResponse{}, fmt.Errorf("Could not test search and recover wallet: %w", err)
	}
	var response api.SearchAndRecoverWalletResponse
	if err := json.Unmarshal(responseBytes, &response); err != nil {
		return api.SearchAndRecoverWalletResponse{}, fmt.Errorf("Could not decode test-search-and-recover wallet response: %w", err)
	}
	if response.Error != "" {
		return api.SearchAndRecoverWalletResponse{}, fmt.Errorf("Could not test search and recover wallet: %s", response.Error)
	}
	return response, nil
}

// Rebuild wallet
func (c *Client) RebuildWallet() (api.RebuildWalletResponse, error) {
	responseBytes, err := c.callAPI("wallet rebuild")
	if err != nil {
		return api.RebuildWalletResponse{}, fmt.Errorf("Could not rebuild wallet: %w", err)
	}
	var response api.RebuildWalletResponse
	if err := json.Unmarshal(responseBytes, &response); err != nil {
		return api.RebuildWalletResponse{}, fmt.Errorf("Could not decode rebuild wallet response: %w", err)
	}
	if response.Error != "" {
		return api.RebuildWalletResponse{}, fmt.Errorf("Could not rebuild wallet: %s", response.Error)
	}
	return response, nil
}

// Purge the node wallet and validator keys
func (c *Client) Purge() (api.PurgeResponse, error) {
	responseBytes, err := c.callAPI("wallet purge")
	if err != nil {
		return api.PurgeResponse{}, fmt.Errorf("Could not purge wallet and keys: %w", err)
	}
	var response api.PurgeResponse
	if err := json.Unmarshal(responseBytes, &response); err != nil {
		return api.PurgeResponse{}, fmt.Errorf("Could not decode purge response: %w", err)
	}
	if response.Error != "" {
		return api.PurgeResponse{}, fmt.Errorf("Could not purge wallet and keys: %s", response.Error)
	}
	return response, nil
}

// Estimate the gas required to set an ENS reverse record to a name
func (c *Client) EstimateGasSetEnsName(name string) (api.SetEnsNameResponse, error) {
	responseBytes, err := c.callAPI(fmt.Sprintf("wallet estimate-gas-set-ens-name %s", name))
	if err != nil {
		return api.SetEnsNameResponse{}, fmt.Errorf("Could not get estimate-gas-set-ens-name response: %w", err)
	}
	var response api.SetEnsNameResponse
	if err := json.Unmarshal(responseBytes, &response); err != nil {
		return api.SetEnsNameResponse{}, fmt.Errorf("Could not decode estimate-gas-set-ens-name response: %w", err)
	}
	if response.Error != "" {
		return api.SetEnsNameResponse{}, fmt.Errorf("Could not get estimate-gas-set-ens-name response: %s", response.Error)
	}
	return response, nil
}

// Set an ENS reverse record to a name
func (c *Client) SetEnsName(name string) (api.SetEnsNameResponse, error) {
	responseBytes, err := c.callAPI(fmt.Sprintf("wallet set-ens-name %s", name))
	if err != nil {
		return api.SetEnsNameResponse{}, fmt.Errorf("Could not update ENS record: %w", err)
	}
	var response api.SetEnsNameResponse
	if err := json.Unmarshal(responseBytes, &response); err != nil {
		return api.SetEnsNameResponse{}, fmt.Errorf("Could not decode set-ens-name response: %w", err)
	}
	if response.Error != "" {
		return api.SetEnsNameResponse{}, fmt.Errorf("Could not update ENS record: %s", response.Error)
	}
	return response, nil
}

// Export wallet
func (c *Client) ExportWallet() (api.ExportWalletResponse, error) {
	responseBytes, err := c.callAPI("wallet export")
	if err != nil {
		return api.ExportWalletResponse{}, fmt.Errorf("Could not export wallet: %w", err)
	}
	var response api.ExportWalletResponse
	if err := json.Unmarshal(responseBytes, &response); err != nil {
		return api.ExportWalletResponse{}, fmt.Errorf("Could not decode export wallet response: %w", err)
	}
	if response.Error != "" {
		return api.ExportWalletResponse{}, fmt.Errorf("Could not export wallet: %s", response.Error)
	}
	return response, nil
}
