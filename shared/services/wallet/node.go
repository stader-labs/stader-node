/*
This work is licensed and released under GNU GPL v3 or any other later versions.
The full text of the license is below/ found at <http://www.gnu.org/licenses/>

(c) 2023 Rocket Pool Pty Ltd. Modified under GNU GPL v3. [1.0.0]

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
	"crypto/ecdsa"
	"errors"
	"fmt"
	"math/big"

	"github.com/btcsuite/btcd/btcutil/hdkeychain"
	"github.com/ethereum/go-ethereum/accounts"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/crypto"
)

// Get the node account
func (w *Wallet) GetNodeAccount() (accounts.Account, error) {

	// Check wallet is initialized
	if !w.IsInitialized() {
		return accounts.Account{}, errors.New("Wallet is not initialized")
	}

	// Get private key
	privateKey, path, err := w.getNodePrivateKey()
	if err != nil {
		return accounts.Account{}, err
	}

	// Get public key
	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		return accounts.Account{}, errors.New("Could not get node public key")
	}

	// Create & return account
	return accounts.Account{
		Address: crypto.PubkeyToAddress(*publicKeyECDSA),
		URL: accounts.URL{
			Scheme: "",
			Path:   path,
		},
	}, nil

}

// Get a transactor for the node account
func (w *Wallet) GetNodeAccountTransactor() (*bind.TransactOpts, error) {

	// Check wallet is initialized
	if !w.IsInitialized() {
		return nil, errors.New("Wallet is not initialized")
	}

	// Get private key
	privateKey, _, err := w.getNodePrivateKey()
	if err != nil {
		return nil, err
	}

	// Create & return transactor
	transactor, err := bind.NewKeyedTransactorWithChainID(privateKey, w.chainID)
	transactor.GasFeeCap = w.maxFee
	transactor.GasTipCap = w.maxPriorityFee
	transactor.GasLimit = w.gasLimit
	transactor.Context = context.Background()
	return transactor, err

}

func (w *Wallet) GetNodePrivateKey() (*ecdsa.PrivateKey, error) {
	// Check wallet is initialized
	if !w.IsInitialized() {
		return nil, errors.New("Wallet is not initialized")
	}

	// Get private key
	privateKey, _, err := w.getNodePrivateKey()
	if err != nil {
		return nil, err
	}

	return privateKey, err
}

// Get the node account private key bytes
func (w *Wallet) GetNodePrivateKeyBytes() ([]byte, error) {

	// Check wallet is initialized
	if !w.IsInitialized() {
		return nil, errors.New("Wallet is not initialized")
	}

	// Get private key
	privateKey, _, err := w.getNodePrivateKey()
	if err != nil {
		return nil, err
	}

	// Return private key bytes
	return crypto.FromECDSA(privateKey), nil

}

// Get the node private key
func (w *Wallet) getNodePrivateKey() (*ecdsa.PrivateKey, string, error) {

	// Check for cached node key
	if w.nodeKey != nil {
		return w.nodeKey, w.nodeKeyPath, nil
	}

	// Get derived key
	derivedKey, path, err := w.getNodeDerivedKey(w.ws.WalletIndex)
	if err != nil {
		return nil, "", err
	}

	// Get private key
	privateKey, err := derivedKey.ECPrivKey()
	if err != nil {
		return nil, "", fmt.Errorf("Could not get node private key: %w", err)
	}
	privateKeyECDSA := privateKey.ToECDSA()

	// Cache node key
	w.nodeKey = privateKeyECDSA
	w.nodeKeyPath = path
	// Return
	return privateKeyECDSA, path, nil

}

// Get the derived key & derivation path for the node account at the index
func (w *Wallet) getNodeDerivedKey(index uint) (*hdkeychain.ExtendedKey, string, error) {

	// Get derivation path
	if w.ws.DerivationPath == "" {
		w.ws.DerivationPath = DefaultNodeKeyPath
	}
	derivationPath := fmt.Sprintf(w.ws.DerivationPath, index)

	// Parse derivation path
	path, err := accounts.ParseDerivationPath(derivationPath)
	if err != nil {
		return nil, "", fmt.Errorf("Invalid node key derivation path '%s': %w", derivationPath, err)
	}

	// Follow derivation path
	key := w.mk
	for i, n := range path {
		// Use the legacy implementation for Goerli
		// TODO: remove this if Prater ever goes away!
		if w.chainID.Cmp(big.NewInt(5)) == 0 {
			key, err = key.DeriveNonStandard(n)
		} else {
			key, err = key.Derive(n)
		}
		if err == hdkeychain.ErrInvalidChild {
			return w.getNodeDerivedKey(index + 1)
		} else if err != nil {
			return nil, "", fmt.Errorf("Invalid child key at depth %d: %w", i, err)
		}
	}

	// Return
	return key, derivationPath, nil

}
