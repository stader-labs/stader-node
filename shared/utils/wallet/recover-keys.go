package wallet

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/stader-labs/stader-node/shared/services/wallet"
	"github.com/stader-labs/stader-node/stader-lib/stader"
	"github.com/stader-labs/stader-node/stader-lib/types"
	"github.com/urfave/cli"
)

const (
	bucketSize  uint = 20
	bucketLimit uint = 2000
)

func RecoverMinipoolKeys(c *cli.Context, pnr stader.PermissionlessNodeRegistryContractManager, address common.Address, w *wallet.Wallet, testOnly bool) ([]types.ValidatorPubkey, error) {
	//
	//cfg, err := services.GetConfig(c)
	//if err != nil {
	//	return nil, err
	//}
	//
	//// Get node's validating pubkeys
	//pubkeys, err := minipool.GetNodeValidatingMinipoolPubkeys(rp, address, nil)
	//if err != nil {
	//	return nil, err
	//}
	//
	//// Remove zero pubkeys
	//zeroPubkey := types.ValidatorPubkey{}
	//filteredPubkeys := []types.ValidatorPubkey{}
	//for _, pubkey := range pubkeys {
	//	if !bytes.Equal(pubkey[:], zeroPubkey[:]) {
	//		filteredPubkeys = append(filteredPubkeys, pubkey)
	//	}
	//}
	//pubkeys = filteredPubkeys
	//
	//pubkeyMap := map[types.ValidatorPubkey]bool{}
	//for _, pubkey := range pubkeys {
	//	pubkeyMap[pubkey] = true
	//}
	//
	//// Load custom validator keys
	//customKeyDir := cfg.Smartnode.GetCustomKeyPath()
	//info, err := os.Stat(customKeyDir)
	//if !os.IsNotExist(err) && info.IsDir() {
	//
	//	// Get the custom keystore files
	//	files, err := ioutil.ReadDir(customKeyDir)
	//	if err != nil {
	//		return nil, fmt.Errorf("error enumerating custom keystores: %w", err)
	//	}
	//
	//	// Initialize the BLS library
	//	err = eth2types.InitBLS()
	//	if err != nil {
	//		return nil, fmt.Errorf("error initializing BLS: %w", err)
	//	}
	//
	//	if len(files) > 0 {
	//
	//		// Deserialize the password file
	//		passwordFile := cfg.Smartnode.GetCustomKeyPasswordFilePath()
	//		fileBytes, err := ioutil.ReadFile(passwordFile)
	//		if err != nil {
	//			return nil, fmt.Errorf("%d custom keystores were found but the password file could not be loaded: %w", len(files), err)
	//		}
	//		passwords := map[string]string{}
	//		err = yaml.Unmarshal(fileBytes, &passwords)
	//		if err != nil {
	//			return nil, fmt.Errorf("error unmarshalling custom keystore password file: %w", err)
	//		}
	//
	//		// Process every custom key
	//		for _, file := range files {
	//			// Read the file
	//			bytes, err := ioutil.ReadFile(filepath.Join(customKeyDir, file.Name()))
	//			if err != nil {
	//				return nil, fmt.Errorf("error reading custom keystore %s: %w", file.Name(), err)
	//			}
	//
	//			// Deserialize it
	//			keystore := api.ValidatorKeystore{}
	//			err = json.Unmarshal(bytes, &keystore)
	//			if err != nil {
	//				return nil, fmt.Errorf("error deserializing custom keystore %s: %w", file.Name(), err)
	//			}
	//
	//			// Check if it's one of the pubkeys for the minipool
	//			_, exists := pubkeyMap[keystore.Pubkey]
	//			if !exists {
	//				// This pubkey isn't for any of this node's minipools so ignore it
	//				continue
	//			}
	//
	//			// Get the password for it
	//			formattedPubkey := string-utils.ToUpper(hexutils.RemovePrefix(keystore.Pubkey.Hex()))
	//			password, exists := passwords[formattedPubkey]
	//			if !exists {
	//				return nil, fmt.Errorf("custom keystore for pubkey %s needs a password, but none was provided", keystore.Pubkey.Hex())
	//			}
	//
	//			// Get the encryption function it uses
	//			kdf, exists := keystore.Crypto["kdf"]
	//			if !exists {
	//				return nil, fmt.Errorf("error processing custom keystore %s: \"crypto\" didn't contain a subkey named \"kdf\"", file.Name())
	//			}
	//			kdfMap := kdf.(map[string]interface{})
	//			function, exists := kdfMap["function"]
	//			if !exists {
	//				return nil, fmt.Errorf("error processing custom keystore %s: \"crypto.kdf\" didn't contain a subkey named \"function\"", file.Name())
	//			}
	//			functionString := function.(string)
	//
	//			// Decrypt the private key
	//			encryptor := eth2ks.New(eth2ks.WithCipher(functionString))
	//			decryptedKey, err := encryptor.Decrypt(keystore.Crypto, password)
	//			if err != nil {
	//				return nil, fmt.Errorf("error decrypting keystore for validator %s: %w", keystore.Pubkey.Hex(), err)
	//			}
	//			privateKey, err := eth2types.BLSPrivateKeyFromBytes(decryptedKey)
	//			if err != nil {
	//				return nil, fmt.Errorf("error recreating private key for validator %s: %w", keystore.Pubkey.Hex(), err)
	//			}
	//
	//			// Verify the private key matches the public key
	//			reconstructedPubkey := types.BytesToValidatorPubkey(privateKey.PublicKey().Marshal())
	//			if reconstructedPubkey != keystore.Pubkey {
	//				return nil, fmt.Errorf("private keystore file %s claims to be for validator %s but it's for validator %s", file.Name(), keystore.Pubkey.Hex(), reconstructedPubkey.Hex())
	//			}
	//
	//			// Store the key
	//			if !testOnly {
	//				err = w.StoreValidatorKey(privateKey, keystore.Path)
	//				if err != nil {
	//					return nil, fmt.Errorf("error storing private keystore for %s: %w", reconstructedPubkey.Hex(), err)
	//				}
	//			}
	//
	//			// Remove the pubkey from pending minipools to handle
	//			delete(pubkeyMap, reconstructedPubkey)
	//		}
	//	}
	//}
	//
	//// Recover conventionally generated keys
	//bucketStart := uint(0)
	//for {
	//	if bucketStart >= bucketLimit {
	//		return nil, fmt.Errorf("attempt limit exceeded (%d keys)", bucketLimit)
	//	}
	//	bucketEnd := bucketStart + bucketSize
	//	if bucketEnd > bucketLimit {
	//		bucketEnd = bucketLimit
	//	}
	//
	//	// Get the keys for this bucket
	//	keys, err := w.GetValidatorKeys(bucketStart, bucketEnd-bucketStart)
	//	if err != nil {
	//		return nil, err
	//	}
	//	for _, validatorKey := range keys {
	//		_, exists := pubkeyMap[validatorKey.PublicKey]
	//		if exists {
	//			// Found one!
	//			delete(pubkeyMap, validatorKey.PublicKey)
	//			if !testOnly {
	//				err := w.SaveValidatorKey(validatorKey)
	//				if err != nil {
	//					return nil, fmt.Errorf("error recovering validator keys: %w", err)
	//				}
	//			}
	//		}
	//	}
	//
	//	if len(pubkeyMap) == 0 {
	//		// All keys recovered!
	//		break
	//	}
	//
	//	// Run another iteration with the next bucket
	//	bucketStart = bucketEnd
	//}

	return []types.ValidatorPubkey{}, nil

}
