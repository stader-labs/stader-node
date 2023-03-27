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
	"bytes"
	"errors"
	"fmt"
	"os"
	"sync"

	"github.com/stader-labs/stader-node/stader-lib/types"
	stadertypes "github.com/stader-labs/stader-node/stader-lib/types"
	eth2types "github.com/wealdtech/go-eth2-types/v2"
	eth2util "github.com/wealdtech/go-eth2-util"
)

// Config
const (
	ValidatorKeyPath               string = "m/12381/3600/%d/0/0"
	MaxValidatorKeyRecoverAttempts uint   = 1000
)

// A validator private/public key pair
type ValidatorKey struct {
	PublicKey      types.ValidatorPubkey
	PrivateKey     *eth2types.BLSPrivateKey
	DerivationPath string
	WalletIndex    uint
}

// Get the number of validator keys recorded in the wallet
func (w *Wallet) GetValidatorKeyCount() (uint, error) {

	// Check wallet is initialized
	if !w.IsInitialized() {
		return 0, errors.New("Wallet is not initialized")
	}

	// Return validator key count
	return w.ws.NextAccount, nil

}

// Get a validator key by index
func (w *Wallet) GetValidatorKeyAt(index uint) (*eth2types.BLSPrivateKey, error) {

	// Check wallet is initialized
	if !w.IsInitialized() {
		return nil, errors.New("Wallet is not initialized")
	}

	// Return validator key
	key, _, err := w.getValidatorPrivateKey(index)
	return key, err

}

// Get a validator key by public key
func (w *Wallet) GetValidatorKeyByPubkey(pubkey stadertypes.ValidatorPubkey) (*eth2types.BLSPrivateKey, error) {

	// Check wallet is initialized
	if !w.IsInitialized() {
		return nil, errors.New("Wallet is not initialized")
	}

	// Get pubkey hex string
	pubkeyHex := pubkey.Hex()

	// Check for cached validator key index
	if index, ok := w.validatorKeyIndices[pubkeyHex]; ok {
		if key, _, err := w.getValidatorPrivateKey(index); err != nil {
			return nil, err
		} else if bytes.Equal(pubkey.Bytes(), key.PublicKey().Marshal()) {
			return key, nil
		}
	}

	// Find matching validator key
	var index uint
	var validatorKey *eth2types.BLSPrivateKey
	for index = 0; index < w.ws.NextAccount; index++ {
		if key, _, err := w.getValidatorPrivateKey(index); err != nil {
			return nil, err
		} else if bytes.Equal(pubkey.Bytes(), key.PublicKey().Marshal()) {
			validatorKey = key
			break
		}
	}

	// Check validator key
	if validatorKey == nil {
		return nil, fmt.Errorf("Validator %s key not found", pubkeyHex)
	}

	// Cache validator key index
	w.validatorKeyIndices[pubkeyHex] = index

	// Return
	return validatorKey, nil

}

func (w *Wallet) CreateValidatorKeyFromIndex(index uint) (*eth2types.BLSPrivateKey, error) {
	// Check wallet is initialized
	if !w.IsInitialized() {
		return nil, errors.New("Wallet is not initialized")
	}

	// Get validator key
	key, _, err := w.getValidatorPrivateKey(index)
	if err != nil {
		return nil, err
	}

	return key, err

}

func (w *Wallet) GetNextAccount() uint {
	return w.ws.NextAccount
}

// Create a new validator key
func (w *Wallet) CreateValidatorKey() (*eth2types.BLSPrivateKey, error) {

	// Check wallet is initialized
	if !w.IsInitialized() {
		return nil, errors.New("Wallet is not initialized")
	}

	// Get & increment account index
	index := w.ws.NextAccount
	w.ws.NextAccount++

	// Get validator key
	key, path, err := w.getValidatorPrivateKey(index)
	if err != nil {
		return nil, err
	}

	// Update keystores
	err = w.StoreValidatorKey(key, path)
	if err != nil {
		return nil, err
	}

	// Return validator key
	return key, nil

}

func (w *Wallet) StoreValidatorKey(key *eth2types.BLSPrivateKey, path string) error {

	for name := range w.keystores {
		// Update the keystore in the wallet - using an iterator variable only runs it on the local copy
		if err := w.keystores[name].StoreValidatorKey(key, path); err != nil {
			return fmt.Errorf("Could not store %s validator key: %w", name, err)
		}
	}

	// Return validator key
	return nil

}

// Deletes all of the keystore directories and persistent VC storage
func (w *Wallet) DeleteValidatorStores() error {

	for name := range w.keystores {
		keystorePath := w.keystores[name].GetKeystoreDir()
		err := os.RemoveAll(keystorePath)
		if err != nil {
			return fmt.Errorf("error deleting validator directory for %s: %w", name, err)
		}
	}

	return nil

}

// Returns the next validator key that will be generated without saving it
func (w *Wallet) GetNextValidatorKey() (*eth2types.BLSPrivateKey, error) {

	// Check wallet is initialized
	if !w.IsInitialized() {
		return nil, errors.New("Wallet is not initialized")
	}

	// Get account index
	index := w.ws.NextAccount

	// Get validator key
	key, _, err := w.getValidatorPrivateKey(index)
	if err != nil {
		return nil, err
	}

	// Return validator key
	return key, nil

}

// Recover a set of validator keys by their public key
func (w *Wallet) GetValidatorKeys(startIndex uint, length uint) ([]ValidatorKey, error) {

	// Check wallet is initialized
	if !w.IsInitialized() {
		return nil, errors.New("Wallet is not initialized")
	}

	validatorKeys := make([]ValidatorKey, 0, length)
	for index := startIndex; index < startIndex+length; index++ {
		key, path, err := w.getValidatorPrivateKey(index)
		if err != nil {
			return nil, fmt.Errorf("error getting validator key for index %d: %w", index, err)
		}
		validatorKey := ValidatorKey{
			PublicKey:      types.BytesToValidatorPubkey(key.PublicKey().Marshal()),
			PrivateKey:     key,
			DerivationPath: path,
			WalletIndex:    index,
		}
		validatorKeys = append(validatorKeys, validatorKey)
	}

	return validatorKeys, nil

}

// Save a validator key
func (w *Wallet) SaveValidatorKey(key ValidatorKey) error {

	// Update account index
	if key.WalletIndex > w.ws.NextAccount {
		w.ws.NextAccount = key.WalletIndex
	}

	// Update keystores
	for name := range w.keystores {
		// Update the keystore in the wallet - using an iterator variable only runs it on the local copy
		if err := w.keystores[name].StoreValidatorKey(key.PrivateKey, key.DerivationPath); err != nil {
			return fmt.Errorf("could not store validator key %s in %s keystore: %w", key.PublicKey.Hex(), name, err)
		}
	}

	// Return
	return nil

}

// Recover a validator key by public key
func (w *Wallet) RecoverValidatorKey(pubkey stadertypes.ValidatorPubkey, startIndex uint) (uint, error) {

	// Check wallet is initialized
	if !w.IsInitialized() {
		return 0, errors.New("Wallet is not initialized")
	}

	// Find matching validator key
	var index uint
	var validatorKey *eth2types.BLSPrivateKey
	var derivationPath string
	for index = 0; index < MaxValidatorKeyRecoverAttempts; index++ {
		if key, path, err := w.getValidatorPrivateKey(index + startIndex); err != nil {
			return 0, err
		} else if bytes.Equal(pubkey.Bytes(), key.PublicKey().Marshal()) {
			validatorKey = key
			derivationPath = path
			break
		}
	}

	// Check validator key
	if validatorKey == nil {
		return 0, fmt.Errorf("Validator %s key not found", pubkey.Hex())
	}

	// Update account index
	nextIndex := index + startIndex + 1
	if nextIndex > w.ws.NextAccount {
		w.ws.NextAccount = nextIndex
	}

	// Update keystores
	for name := range w.keystores {
		// Update the keystore in the wallet - using an iterator variable only runs it on the local copy
		if err := w.keystores[name].StoreValidatorKey(validatorKey, derivationPath); err != nil {
			return 0, fmt.Errorf("Could not store %s validator key: %w", name, err)
		}
	}

	// Return
	return index + startIndex, nil

}

// Test recovery of a validator key by public key
func (w *Wallet) TestRecoverValidatorKey(pubkey stadertypes.ValidatorPubkey, startIndex uint) (uint, error) {

	// Check wallet is initialized
	if !w.IsInitialized() {
		return 0, errors.New("Wallet is not initialized")
	}

	// Find matching validator key
	var index uint
	var validatorKey *eth2types.BLSPrivateKey
	for index = 0; index < MaxValidatorKeyRecoverAttempts; index++ {
		if key, _, err := w.getValidatorPrivateKey(index + startIndex); err != nil {
			return 0, err
		} else if bytes.Equal(pubkey.Bytes(), key.PublicKey().Marshal()) {
			validatorKey = key
			break
		}
	}

	// Check validator key
	if validatorKey == nil {
		return 0, fmt.Errorf("Validator %s key not found", pubkey.Hex())
	}

	// Return
	return index + startIndex, nil

}

// Get a validator private key by index
func (w *Wallet) getValidatorPrivateKey(index uint) (*eth2types.BLSPrivateKey, string, error) {

	// Get derivation path
	derivationPath := fmt.Sprintf(ValidatorKeyPath, index)

	// Check for cached validator key
	if validatorKey, ok := w.validatorKeys[index]; ok {
		return validatorKey, derivationPath, nil
	}

	// Initialize BLS support
	if err := initializeBLS(); err != nil {
		return nil, "", fmt.Errorf("Could not initialize BLS library: %w", err)
	}

	// Get private key
	privateKey, err := eth2util.PrivateKeyFromSeedAndPath(w.seed, derivationPath)
	if err != nil {
		return nil, "", fmt.Errorf("Could not get validator %d private key: %w", index, err)
	}

	// Cache validator key
	w.validatorKeys[index] = privateKey

	// Return
	return privateKey, derivationPath, nil

}

// Initialize BLS support
var initBLS sync.Once

func initializeBLS() error {
	var err error
	initBLS.Do(func() {
		err = eth2types.InitBLS()
	})
	return err
}
