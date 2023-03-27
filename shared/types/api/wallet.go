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
package api

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/google/uuid"
	"github.com/stader-labs/stader-node/stader-lib/stader"
	"github.com/stader-labs/stader-node/stader-lib/types"
)

// Encrypted validator keystore following the EIP-2335 standard
// (https://eips.ethereum.org/EIPS/eip-2335)
type ValidatorKeystore struct {
	Crypto  map[string]interface{} `json:"crypto"`
	Version uint                   `json:"version"`
	UUID    uuid.UUID              `json:"uuid"`
	Path    string                 `json:"path"`
	Pubkey  types.ValidatorPubkey  `json:"pubkey"`
}

type WalletStatusResponse struct {
	Status            string         `json:"status"`
	Error             string         `json:"error"`
	PasswordSet       bool           `json:"passwordSet"`
	WalletInitialized bool           `json:"walletInitialized"`
	AccountAddress    common.Address `json:"accountAddress"`
}

type SetPasswordResponse struct {
	Status string `json:"status"`
	Error  string `json:"error"`
}

type InitWalletResponse struct {
	Status         string         `json:"status"`
	Error          string         `json:"error"`
	Mnemonic       string         `json:"mnemonic"`
	AccountAddress common.Address `json:"accountAddress"`
}

type RecoverWalletResponse struct {
	Status         string                  `json:"status"`
	Error          string                  `json:"error"`
	AccountAddress common.Address          `json:"accountAddress"`
	ValidatorKeys  []types.ValidatorPubkey `json:"validatorKeys"`
}

type SearchAndRecoverWalletResponse struct {
	Status         string                  `json:"status"`
	Error          string                  `json:"error"`
	FoundWallet    bool                    `json:"foundWallet"`
	AccountAddress common.Address          `json:"accountAddress"`
	DerivationPath string                  `json:"derivationPath"`
	Index          uint                    `json:"index"`
	ValidatorKeys  []types.ValidatorPubkey `json:"validatorKeys"`
}

type RebuildWalletResponse struct {
	Status        string                  `json:"status"`
	Error         string                  `json:"error"`
	ValidatorKeys []types.ValidatorPubkey `json:"validatorKeys"`
}

type ExportWalletResponse struct {
	Status            string `json:"status"`
	Error             string `json:"error"`
	Password          string `json:"password"`
	Wallet            string `json:"wallet"`
	AccountPrivateKey string `json:"accountPrivateKey"`
}

type SetEnsNameResponse struct {
	Status  string         `json:"status"`
	Error   string         `json:"error"`
	Address common.Address `json:"address"`
	EnsName string         `json:"ensName"`
	TxHash  common.Hash    `json:"txHash"`
	GasInfo stader.GasInfo `json:"gasInfo"`
}

type TestMnemonicResponse struct {
	Status           string         `json:"status"`
	Error            string         `json:"error"`
	CurrentAddress   common.Address `json:"currentAddress"`
	RecoveredAddress common.Address `json:"recoveredAddress"`
}

type PurgeResponse struct {
	Status string `json:"status"`
	Error  string `json:"error"`
}
