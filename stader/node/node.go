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
	"github.com/stader-labs/stader-node/shared/utils/log"
)

// Config
var preSignedCooldown, _ = time.ParseDuration("10s")
var preSignedBatchCooldown, _ = time.ParseDuration("5s")
var preSignBatchSize = 10 // Go thru 100 keys in each pass
var tasksInterval, _ = time.ParseDuration("5m")
var taskCooldown, _ = time.ParseDuration("10s")

const (
	MaxConcurrentEth1Requests = 200
	ManageFeeRecipientColor   = color.FgHiCyan
	ErrorColor                = color.FgRed
	InfoColor                 = color.FgHiGreen
)

// Register node command
func RegisterCommands(app *cli.App, name string, aliases []string) {
	app.Commands = append(app.Commands, cli.Command{
		Name:    name,
		Aliases: aliases,
		Usage:   "Run Stader node activity daemon",
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
	wg.Add(2)

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

			for i := uint(0); i <= noOfBatches; i++ {
				for j := batchIndex; j < batchIndex+preSignBatchSize && j < int(walletIndex); j++ {
					infoLog.Printf("Checking validator index %d\n", j)
					validatorPrivateKey, err := w.GetValidatorKeyAt(uint(j))
					// log the errors and continue. dont need to sleep post an error
					if err != nil {
						errorLog.Printf("Could not find validator private key for validator index %d\n", j)
						continue
					}

					validatorPubKey := types.BytesToValidatorPubkey(validatorPrivateKey.PublicKey().Marshal())
					infoLog.Printf("Checking validator Pub key: %s\n", validatorPubKey.String())

					// check if validator has not yet been registered
					validatorStatus, err := bc.GetValidatorStatus(validatorPubKey, nil)
					if err != nil {
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
					// TODO - bchain - verify with sigma prime
					exitEpoch := currentHead.Epoch
					epochsSinceActivation := currentHead.Epoch - validatorStatus.ActivationEpoch
					if epochsSinceActivation < 256 {
						exitEpoch = exitEpoch + (256 - epochsSinceActivation)
					}

					signatureDomain, err := bc.GetDomainData(eth2types.DomainVoluntaryExit[:], exitEpoch, false)
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

	// Wait for both threads to stop
	wg.Wait()
	return nil

}

// Configure HTTP transport settings
func configureHTTP() {

	// The guardian daemon makes a large number of concurrent RPC requests to the Eth1 client
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

	feeRecipientPath := cfg.StaderNode.GetFeeRecipientFilePath()
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
			defaultFeeRecipientFileContents = fmt.Sprintf("%s=%s", config.FeeRecipientEnvVar, cfg.StaderNode.GetEthxTokenAddress().Hex())
		} else {
			// Docker and Hybrid just need the address itself
			defaultFeeRecipientFileContents = cfg.StaderNode.GetEthxTokenAddress().Hex()
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
