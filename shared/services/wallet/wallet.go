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
	"crypto/ecdsa"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"math/big"
	"os"

	"github.com/btcsuite/btcd/btcutil/hdkeychain"
	"github.com/btcsuite/btcd/chaincfg"
	"github.com/ethereum/go-ethereum/accounts"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/google/uuid"
	"github.com/tyler-smith/go-bip39"
	eth2types "github.com/wealdtech/go-eth2-types/v2"
	eth2ks "github.com/wealdtech/go-eth2-wallet-encryptor-keystorev4"

	"github.com/stader-labs/stader-node/shared/services/passwords"
	"github.com/stader-labs/stader-node/shared/services/wallet/keystore"
)

// Config
const (
	EntropyBits              = 256
	FileMode                 = 0600
	DefaultNodeKeyPath       = "m/44'/60'/0'/0/%d"
	LedgerLiveNodeKeyPath    = "m/44'/60'/%d/0/0"
	MyEtherWalletNodeKeyPath = "m/44'/60'/0'/%d"
)

// Wallet
type Wallet struct {

	// Core
	walletPath string
	pm         *passwords.PasswordManager
	encryptor  *eth2ks.Encryptor
	chainID    *big.Int

	// Encrypted store
	ws *walletStore

	// Seed & master key
	seed []byte
	mk   *hdkeychain.ExtendedKey

	// Node key cache
	nodeKey     *ecdsa.PrivateKey
	nodeKeyPath string

	// Validator key caches
	validatorKeys       map[uint]*eth2types.BLSPrivateKey
	validatorKeyIndices map[string]uint

	// Keystores
	keystores map[string]keystore.Keystore

	// Desired gas price & limit from config
	maxFee         *big.Int
	maxPriorityFee *big.Int
	gasLimit       uint64
}

// Encrypted wallet store
type walletStore struct {
	Crypto         map[string]interface{} `json:"crypto"`
	Name           string                 `json:"name"`
	Version        uint                   `json:"version"`
	UUID           uuid.UUID              `json:"uuid"`
	DerivationPath string                 `json:"derivationPath,omitempty"`
	WalletIndex    uint                   `json:"walletIndex,omitempty"`
	NextAccount    uint                   `json:"next_account"`
}

// Create new wallet
func NewWallet(walletPath string, chainId uint, maxFee *big.Int, maxPriorityFee *big.Int, gasLimit uint64, passwordManager *passwords.PasswordManager) (*Wallet, error) {

	// Initialize wallet
	w := &Wallet{
		walletPath:          walletPath,
		pm:                  passwordManager,
		encryptor:           eth2ks.New(),
		chainID:             big.NewInt(int64(chainId)),
		validatorKeys:       map[uint]*eth2types.BLSPrivateKey{},
		validatorKeyIndices: map[string]uint{},
		keystores:           map[string]keystore.Keystore{},
		maxFee:              maxFee,
		maxPriorityFee:      maxPriorityFee,
		gasLimit:            gasLimit,
	}

	// Load & decrypt wallet store
	if _, err := w.loadStore(); err != nil {
		return nil, err
	}

	// Return
	return w, nil

}

// Gets the wallet's chain ID
func (w *Wallet) GetChainID() *big.Int {
	copy := big.NewInt(0).Set(w.chainID)
	return copy
}

// Add a keystore to the wallet
func (w *Wallet) AddKeystore(name string, ks keystore.Keystore) {
	w.keystores[name] = ks
}

// Check if the wallet has been initialized
func (w *Wallet) IsInitialized() bool {
	return (w.ws != nil && w.seed != nil && w.mk != nil)
}

// Attempt to initialize the wallet if not initialized and return status
func (w *Wallet) GetInitialized() (bool, error) {
	if w.IsInitialized() {
		return true, nil
	}
	return w.loadStore()
}

// Serialize the wallet to a JSON string
func (w *Wallet) String() (string, error) {

	// Check wallet is initialized
	if !w.IsInitialized() {
		return "", errors.New("Wallet is not initialized")
	}

	// Encode wallet store
	wsBytes, err := json.Marshal(w.ws)
	if err != nil {
		return "", fmt.Errorf("Could not encode wallet: %w", err)
	}

	// Return
	return string(wsBytes), nil

}

// Initialize the wallet from a random seed
func (w *Wallet) Initialize(derivationPath string, walletIndex uint) (string, error) {

	// Check wallet is not initialized
	if w.IsInitialized() {
		return "", errors.New("Wallet is already initialized")
	}

	// Generate mnemonic entropy
	entropy, err := bip39.NewEntropy(EntropyBits)
	if err != nil {
		return "", fmt.Errorf("Could not generate wallet mnemonic entropy bytes: %w", err)
	}

	// Generate mnemonic
	mnemonic, err := bip39.NewMnemonic(entropy)
	if err != nil {
		return "", fmt.Errorf("Could not generate wallet mnemonic: %w", err)
	}

	// Initialize wallet store
	if err := w.initializeStore(derivationPath, walletIndex, mnemonic); err != nil {
		return "", err
	}

	// Return
	return mnemonic, nil

}

// Recover a wallet from a mnemonic
func (w *Wallet) Recover(derivationPath string, walletIndex uint, mnemonic string) error {

	// Check wallet is not initialized
	if w.IsInitialized() {
		return errors.New("Wallet is already initialized")
	}

	// Check mnemonic
	if !bip39.IsMnemonicValid(mnemonic) {
		return fmt.Errorf("Invalid mnemonic '%s'", mnemonic)
	}

	// Initialize wallet store
	if err := w.initializeStore(derivationPath, walletIndex, mnemonic); err != nil {
		return err
	}

	// Return
	return nil

}

// Recover a wallet from a mnemonic - only used for testing mnemonics
func (w *Wallet) TestRecovery(derivationPath string, walletIndex uint, mnemonic string) error {

	// Check mnemonic
	if !bip39.IsMnemonicValid(mnemonic) {
		return fmt.Errorf("Invalid mnemonic '%s'", mnemonic)
	}

	// Generate seed
	w.seed = bip39.NewSeed(mnemonic, "")

	// Create master key
	var err error
	w.mk, err = hdkeychain.NewMaster(w.seed, &chaincfg.MainNetParams)
	if err != nil {
		return fmt.Errorf("Could not create wallet master key: %w", err)
	}

	// Create wallet store
	w.ws = &walletStore{
		Name:           w.encryptor.Name(),
		Version:        w.encryptor.Version(),
		UUID:           uuid.New(),
		DerivationPath: derivationPath,
		WalletIndex:    walletIndex,
		NextAccount:    0,
	}

	// Return
	return nil

}

// Save the wallet store to disk
func (w *Wallet) Save() error {

	// Check wallet is initialized
	if !w.IsInitialized() {
		return errors.New("Wallet is not initialized")
	}

	// Encode wallet store
	wsBytes, err := json.Marshal(w.ws)
	if err != nil {
		return fmt.Errorf("Could not encode wallet: %w", err)
	}

	// Write wallet store to disk
	if err := ioutil.WriteFile(w.walletPath, wsBytes, FileMode); err != nil {
		return fmt.Errorf("Could not write wallet to disk: %w", err)
	}

	// Return
	return nil

}

// Delete the wallet store from disk
func (w *Wallet) Delete() error {

	// Check if it exists
	_, err := os.Stat(w.walletPath)
	if os.IsNotExist(err) {
		return nil
	} else if err != nil {
		return fmt.Errorf("error checking wallet file path: %w", err)
	}

	// Write wallet store to disk
	err = os.Remove(w.walletPath)
	return err

}

// Signs a serialized TX using the wallet's private key
func (w *Wallet) Sign(serializedTx []byte) ([]byte, error) {
	// Get private key
	privateKey, _, err := w.getNodePrivateKey()
	if err != nil {
		return nil, err
	}

	tx := types.Transaction{}
	err = tx.UnmarshalBinary(serializedTx)
	if err != nil {
		return nil, fmt.Errorf("Error unmarshalling TX: %w", err)
	}

	signer := types.NewLondonSigner(w.chainID)
	signedTx, err := types.SignTx(&tx, signer, privateKey)
	if err != nil {
		return nil, fmt.Errorf("Error signing TX: %w", err)
	}

	signedData, err := signedTx.MarshalBinary()
	if err != nil {
		return nil, fmt.Errorf("Error marshalling signed TX to binary: %w", err)
	}

	return signedData, nil
}

// Signs an arbitrary message using the wallet's private key
func (w *Wallet) SignMessage(message string) ([]byte, error) {
	// Get the wallet's private key
	privateKey, _, err := w.getNodePrivateKey()
	if err != nil {
		return nil, err
	}

	messageHash := accounts.TextHash([]byte(message))
	signedMessage, err := crypto.Sign(messageHash, privateKey)
	if err != nil {
		return nil, fmt.Errorf("Error signing message: %w", err)
	}

	// fix the ECDSA 'v' (see https://medium.com/mycrypto/the-magic-of-digital-signatures-on-ethereum-98fe184dc9c7#:~:text=The%20version%20number,2%E2%80%9D%20was%20introduced)
	signedMessage[crypto.RecoveryIDOffset] += 27
	return signedMessage, nil
}

// Reloads wallet from disk
func (w *Wallet) Reload() error {
	_, err := w.loadStore()
	return err
}

// Load the wallet store from disk and decrypt it
func (w *Wallet) loadStore() (bool, error) {

	// Read wallet store from disk; cancel if not found
	wsBytes, err := ioutil.ReadFile(w.walletPath)
	if err != nil {
		return false, nil
	}

	// Decode wallet store
	w.ws = new(walletStore)
	if err = json.Unmarshal(wsBytes, w.ws); err != nil {
		return false, fmt.Errorf("Could not decode wallet: %w", err)
	}

	// Upgrade legacy wallets to include derivation paths
	if w.ws.DerivationPath == "" {
		w.ws.DerivationPath = DefaultNodeKeyPath
	}

	// Get wallet password
	password, err := w.pm.GetPassword()
	if err != nil {
		return false, fmt.Errorf("Could not get wallet password: %w", err)
	}

	// Decrypt seed
	w.seed, err = w.encryptor.Decrypt(w.ws.Crypto, password)
	if err != nil {
		return false, fmt.Errorf("Could not decrypt wallet seed: %w", err)
	}

	// Create master key
	w.mk, err = hdkeychain.NewMaster(w.seed, &chaincfg.MainNetParams)
	if err != nil {
		return false, fmt.Errorf("Could not create wallet master key: %w", err)
	}

	// Return
	return true, nil

}

// Initialize the encrypted wallet store from a mnemonic
func (w *Wallet) initializeStore(derivationPath string, walletIndex uint, mnemonic string) error {

	// Generate seed
	w.seed = bip39.NewSeed(mnemonic, "")

	// Create master key
	var err error
	w.mk, err = hdkeychain.NewMaster(w.seed, &chaincfg.MainNetParams)
	if err != nil {
		return fmt.Errorf("Could not create wallet master key: %w", err)
	}

	// Get wallet password
	password, err := w.pm.GetPassword()
	if err != nil {
		return fmt.Errorf("Could not get wallet password: %w", err)
	}

	// Encrypt seed
	encryptedSeed, err := w.encryptor.Encrypt(w.seed, password)
	if err != nil {
		return fmt.Errorf("Could not encrypt wallet seed: %w", err)
	}

	// Create wallet store
	w.ws = &walletStore{
		Crypto:         encryptedSeed,
		Name:           w.encryptor.Name(),
		Version:        w.encryptor.Version(),
		UUID:           uuid.New(),
		DerivationPath: derivationPath,
		WalletIndex:    walletIndex,
		NextAccount:    0,
	}

	// Return
	return nil

}
