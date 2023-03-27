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
package nimbus

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/google/uuid"
	stadertypes "github.com/stader-labs/stader-node/stader-lib/types"
	eth2types "github.com/wealdtech/go-eth2-types/v2"
	eth2ks "github.com/wealdtech/go-eth2-wallet-encryptor-keystorev4"

	"github.com/stader-labs/stader-node/shared/services/passwords"
	keystore "github.com/stader-labs/stader-node/shared/services/wallet/keystore"
	hexutil "github.com/stader-labs/stader-node/shared/utils/hex"
)

// Config
const (
	KeystoreDir   = "nimbus"
	SecretsDir    = "secrets"
	ValidatorsDir = "validators"
	KeyFileName   = "keystore.json"
	DirMode       = 0770
	FileMode      = 0640
)

// Nimbus keystore
type Keystore struct {
	keystorePath string
	pm           *passwords.PasswordManager
	encryptor    *eth2ks.Encryptor
}

// Encrypted validator key store
type validatorKey struct {
	Crypto  map[string]interface{}      `json:"crypto"`
	Version uint                        `json:"version"`
	UUID    uuid.UUID                   `json:"uuid"`
	Path    string                      `json:"path"`
	Pubkey  stadertypes.ValidatorPubkey `json:"pubkey"`
}

// Create new nimbus keystore
func NewKeystore(keystorePath string, passwordManager *passwords.PasswordManager) *Keystore {
	return &Keystore{
		keystorePath: keystorePath,
		pm:           passwordManager,
		encryptor:    eth2ks.New(),
	}
}

// Get the keystore directory
func (ks *Keystore) GetKeystoreDir() string {
	return filepath.Join(ks.keystorePath, KeystoreDir)
}

// Store a validator key
func (ks *Keystore) StoreValidatorKey(key *eth2types.BLSPrivateKey, derivationPath string) error {

	// Get validator pubkey
	pubkey := stadertypes.BytesToValidatorPubkey(key.PublicKey().Marshal())

	// Create a new password
	password, err := keystore.GenerateRandomPassword()
	if err != nil {
		return fmt.Errorf("Could not generate random password: %w", err)
	}

	// Encrypt key
	encryptedKey, err := ks.encryptor.Encrypt(key.Marshal(), password)
	if err != nil {
		return fmt.Errorf("Could not encrypt validator key: %w", err)
	}

	// Create key store
	keyStore := validatorKey{
		Crypto:  encryptedKey,
		Version: ks.encryptor.Version(),
		UUID:    uuid.New(),
		Path:    derivationPath,
		Pubkey:  pubkey,
	}

	// Encode key store
	keyStoreBytes, err := json.Marshal(keyStore)
	if err != nil {
		return fmt.Errorf("Could not encode validator key: %w", err)
	}

	// Get secret file path
	secretFilePath := filepath.Join(ks.keystorePath, KeystoreDir, SecretsDir, hexutil.AddPrefix(pubkey.Hex()))

	// Create secrets dir
	if err := os.MkdirAll(filepath.Dir(secretFilePath), DirMode); err != nil {
		return fmt.Errorf("Could not create validator secrets folder: %w", err)
	}

	// Write secret to disk
	if err := ioutil.WriteFile(secretFilePath, []byte(password), FileMode); err != nil {
		return fmt.Errorf("Could not write validator secret to disk: %w", err)
	}

	// Get key file path
	keyFilePath := filepath.Join(ks.keystorePath, KeystoreDir, ValidatorsDir, hexutil.AddPrefix(pubkey.Hex()), KeyFileName)

	// Create key dir
	if err := os.MkdirAll(filepath.Dir(keyFilePath), DirMode); err != nil {
		return fmt.Errorf("Could not create validator key folder: %w", err)
	}

	// Write key store to disk
	if err := ioutil.WriteFile(keyFilePath, keyStoreBytes, FileMode); err != nil {
		return fmt.Errorf("Could not write validator key to disk: %w", err)
	}

	// Return
	return nil

}
