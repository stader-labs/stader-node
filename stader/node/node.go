package node

import (
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	stader_backend "github.com/stader-labs/stader-node/shared/types/stader-backend"
	"github.com/stader-labs/stader-node/shared/utils/crypto"
	"github.com/stader-labs/stader-node/shared/utils/stader"
	"github.com/stader-labs/stader-node/shared/utils/validator"
	"github.com/stader-labs/stader-node/stader-lib/types"
	eth2types "github.com/wealdtech/go-eth2-types/v2"
	"io/ioutil"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"sync"
	"time"

	"github.com/fatih/color"
	"github.com/urfave/cli"

	"github.com/stader-labs/stader-node/shared/services"
	"github.com/stader-labs/stader-node/shared/services/config"
	"github.com/stader-labs/stader-node/shared/services/wallet/keystore/lighthouse"
	"github.com/stader-labs/stader-node/shared/services/wallet/keystore/nimbus"
	"github.com/stader-labs/stader-node/shared/services/wallet/keystore/prysm"
	"github.com/stader-labs/stader-node/shared/services/wallet/keystore/teku"
	"github.com/stader-labs/stader-node/shared/utils/log"
)

// Config
var preSignedCooldown, _ = time.ParseDuration("12h")
var preSignedBatchCooldown, _ = time.ParseDuration("30s")
var preSignBatchSize = 10 // Go thru 10 keys in each pass
var tasksInterval, _ = time.ParseDuration("5m")
var taskCooldown, _ = time.ParseDuration("10s")

const (
	MaxConcurrentEth1Requests = 200

	MetricsColor            = color.FgHiYellow
	ManageFeeRecipientColor = color.FgHiCyan
	ErrorColor              = color.FgRed
	InfoColor               = color.FgHiCyan
)

// Register node command
func RegisterCommands(app *cli.App, name string, aliases []string) {
	app.Commands = append(app.Commands, cli.Command{
		Name:    name,
		Aliases: aliases,
		Usage:   "Run Rocket Pool node activity daemon",
		Action: func(c *cli.Context) error {
			return run(c)
		},
	})
}

// Run daemon
func run(c *cli.Context) error {

	// Handle the initial fee recipient file deployment
	err := deployDefaultFeeRecipientFile(c)
	if err != nil {
		return err
	}

	w, err := services.GetWallet(c)
	if err != nil {
		return err
	}
	bc, err := services.GetBeaconClient(c)
	if err != nil {
		return err
	}
	publicKey, err := stader.GetPublicKey()
	if err != nil {
		return err
	}

	// Configure
	configureHTTP()

	// Initialize tasks
	manageFeeRecipient, err := newManageFeeRecipient(c, log.NewColorLogger(ManageFeeRecipientColor))
	if err != nil {
		return err
	}

	// Initialize loggers
	errorLog := log.NewColorLogger(ErrorColor)
	infoLog := log.NewColorLogger(InfoColor)

	// Wait group to handle the various threads
	wg := new(sync.WaitGroup)
	wg.Add(3)

	// validator presigned loop
	go func() {
		for {
			infoLog.Println("Starting a pass of the presign daemon!")
			walletIndex := w.GetNextAccount()
			noOfBatches := walletIndex / uint(preSignBatchSize)
			batchIndex := 0

			currentHead, err := bc.GetBeaconHead()
			if err != nil {
				panic("not able to communicate with beacon chain!")
			}

			for i := uint(0); i < noOfBatches; i++ {
				for j := batchIndex; j < batchIndex+preSignBatchSize && j < int(walletIndex); j++ {
					infoLog.Printf("Checking validator index %d\n", j)
					// TODO - bchain - parallelize for each validator for each batch
					validatorPrivateKey, err := w.GetValidatorKeyAt(uint(j))
					// log the errors and continue. dont need to sleep post an error
					if err != nil {
						errorLog.Printf("Could not find validator private key for validator index %d\n", j)
						continue
					}

					validatorPubKey := types.BytesToValidatorPubkey(validatorPrivateKey.PublicKey().Marshal())

					// check if validator has not yet been registered
					validatorStatus, err := bc.GetValidatorStatus(validatorPubKey, nil)
					if validatorStatus.Index == 0 || err != nil {
						errorLog.Printf("Could not find validator status for validator pub key: %s\n", validatorPubKey)
						continue
					}

					// check if the presigned message has been registered. if it has been registered, then continue
					isRegistered, err := stader.IsPresignedKeyRegistered(validatorPubKey)
					if isRegistered {
						infoLog.Printf("Validator pub key: %s already registered\n", validatorPubKey)
						continue
					} else if err != nil {
						errorLog.Printf("Could not query presign api to check if validator: %s is registered\n", validatorPubKey)
					}

					// exit epoch should be > activation_epoch + 256
					// exit epoch should be > current epoch
					exitEpoch := currentHead.Epoch + 1
					epochsSinceActivation := currentHead.Epoch - validatorStatus.ActivationEpoch
					if epochsSinceActivation < 256 {
						exitEpoch = exitEpoch + (256 - epochsSinceActivation) + 1
					}

					signatureDomain, err := bc.GetDomainData(eth2types.DomainVoluntaryExit[:], exitEpoch)
					if err != nil {
						errorLog.Printf("Failed to get the signature domain from beacon chain\n")
						continue
					}

					// get the presigned msg
					exitSignature, srHash, err := validator.GetSignedExitMessage(validatorPrivateKey, validatorStatus.Index, exitEpoch, signatureDomain)
					if err != nil {
						errorLog.Printf("Failed to generate the SignedExitMessage for validator with beacon chain index: %d\n", validatorStatus.Index)
						continue
					}
					srHashHex := common.Bytes2Hex(srHash[:])

					// encrypt the signature and srHash
					exitSignatureEncrypted, err := crypto.EncryptUsingPublicKey([]byte(exitSignature.String()), publicKey)
					if err != nil {
						errorLog.Printf("Failed to encrypt exit signature for validator: %s\n", validatorPubKey)
						continue
					}
					// TODO - bchain - revise the naming
					exitSignatureEncryptedString := crypto.EncodeBase64(exitSignatureEncrypted)

					messageHashEncrypted, err := crypto.EncryptUsingPublicKey([]byte(srHashHex), publicKey)
					if err != nil {
						errorLog.Printf("Failed to encrypt message hash for validator: %s\n", validatorPubKey)
						continue
					}
					messageHashEncryptedString := crypto.EncodeBase64(messageHashEncrypted)

					// send it to the presigned api
					backendRes, err := stader.SendPresignedMessageToStaderBackend(stader_backend.PreSignSendApiRequestType{
						Message: struct {
							Epoch          string `json:"epoch"`
							ValidatorIndex string `json:"validator_index"`
						}{
							Epoch:          strconv.FormatUint(exitEpoch, 10),
							ValidatorIndex: strconv.FormatUint(validatorStatus.Index, 10),
						},
						MessageHash:        messageHashEncryptedString,
						Signature:          exitSignatureEncryptedString,
						ValidatorPublicKey: validatorPubKey.String(),
					})
					if !backendRes.Success {
						errorLog.Printf("Failed to send the presigned api: %s\n", backendRes.Message)
						continue
					} else if backendRes.Success {
						errorLog.Printf("Successfully sent the presigned message for validator: %s\n", validatorPubKey)
						continue
					} else if err != nil {
						errorLog.Printf("Sending presigned message failed with %v\n", err)
						continue
					}

					time.Sleep(preSignedBatchCooldown)
				}

				batchIndex = batchIndex + preSignBatchSize
			}

			errorLog.Printf("Done with the pass of presign daemon")
			// run loop every 12 hours
			time.Sleep(preSignedCooldown)
		}

		wg.Done()
	}()

	// Run task loop
	go func() {
		for {
			// Check the EC status
			err := services.WaitEthClientSynced(c, false) // Force refresh the primary / fallback EC status
			if err != nil {
				errorLog.Println(err)
			} else {
				// Check the BC status
				err := services.WaitBeaconClientSynced(c, false) // Force refresh the primary / fallback BC status
				if err != nil {
					errorLog.Println(err)
				} else {
					// Manage the fee recipient for the node
					if err := manageFeeRecipient.run(); err != nil {
						errorLog.Println(err)
					}
					time.Sleep(taskCooldown)
				}
			}
			time.Sleep(tasksInterval)
		}
		wg.Done()
	}()

	// Run metrics loop
	go func() {
		err := runMetricsServer(c, log.NewColorLogger(MetricsColor))
		if err != nil {
			errorLog.Println(err)
		}
		wg.Done()
	}()

	// Wait for both threads to stop
	wg.Wait()
	return nil

}

// Configure HTTP transport settings
func configureHTTP() {

	// The watchtower daemon makes a large number of concurrent RPC requests to the Eth1 client
	// The HTTP transport is set to cache connections for future re-use equal to the maximum expected number of concurrent requests
	// This prevents issues related to memory consumption and address allowance from repeatedly opening and closing connections
	http.DefaultTransport.(*http.Transport).MaxIdleConnsPerHost = MaxConcurrentEth1Requests

}

// Copy the default fee recipient file into the proper location
func deployDefaultFeeRecipientFile(c *cli.Context) error {

	cfg, err := services.GetConfig(c)
	if err != nil {
		return err
	}

	feeRecipientPath := cfg.Smartnode.GetFeeRecipientFilePath()
	_, err = os.Stat(feeRecipientPath)
	if os.IsNotExist(err) {
		// Make sure the validators dir is created
		validatorsFolder := filepath.Dir(feeRecipientPath)
		err = os.MkdirAll(validatorsFolder, 0755)
		if err != nil {
			return fmt.Errorf("could not create validators directory: %w", err)
		}

		// Create the file
		var defaultFeeRecipientFileContents string
		if cfg.IsNativeMode {
			// Native mode needs an environment variable definition
			defaultFeeRecipientFileContents = fmt.Sprintf("%s=%s", config.FeeRecipientEnvVar, cfg.Smartnode.GetRethAddress().Hex())
		} else {
			// Docker and Hybrid just need the address itself
			defaultFeeRecipientFileContents = cfg.Smartnode.GetRethAddress().Hex()
		}
		err := ioutil.WriteFile(feeRecipientPath, []byte(defaultFeeRecipientFileContents), 0664)
		if err != nil {
			return fmt.Errorf("could not write default fee recipient file to %s: %w", feeRecipientPath, err)
		}
	} else if err != nil {
		return fmt.Errorf("error checking fee recipient file status: %w", err)
	}

	return nil

}

// Remove the old fee recipient files that were created in v1.5.0
func removeLegacyFeeRecipientFiles(c *cli.Context) error {

	legacyFeeRecipientFile := "rp-fee-recipient.txt"

	cfg, err := services.GetConfig(c)
	if err != nil {
		return err
	}

	validatorsFolder := cfg.Smartnode.GetValidatorKeychainPath()

	// Remove the legacy files
	keystoreDirs := []string{lighthouse.KeystoreDir, nimbus.KeystoreDir, prysm.KeystoreDir, teku.KeystoreDir}
	for _, keystoreDir := range keystoreDirs {
		oldFile := filepath.Join(validatorsFolder, keystoreDir, legacyFeeRecipientFile)
		_, err = os.Stat(oldFile)
		if !os.IsNotExist(err) {
			err = os.Remove(oldFile)
			if err != nil {
				fmt.Printf("NOTE: Couldn't remove old fee recipient file (%s): %s\nThis file is no longer used, you may remove it manually if you wish.\n", oldFile, err.Error())
			}
		}
	}

	return nil

}
