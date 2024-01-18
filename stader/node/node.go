/*
This work is licensed and released under GNU GPL v3 or any other later versions.
The full text of the license is below/ found at <http://www.gnu.org/licenses/>

(c) 2023 Rocket Pool Pty Ltd. Modified under GNU GPL v3. [1.4.4]

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
	_ "embed"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"sync"
	"time"

	cfgtypes "github.com/stader-labs/stader-node/shared/types/config"
	stader_backend "github.com/stader-labs/stader-node/shared/types/stader-backend"
	"github.com/stader-labs/stader-node/shared/utils/crypto"
	"github.com/stader-labs/stader-node/shared/utils/eth2"
	"github.com/stader-labs/stader-node/shared/utils/stader"
	"github.com/stader-labs/stader-node/shared/utils/stdr"
	"github.com/stader-labs/stader-node/shared/utils/validator"
	"github.com/stader-labs/stader-node/stader-lib/node"
	eth2types "github.com/wealdtech/go-eth2-types/v2"

	"github.com/fatih/color"
	"github.com/urfave/cli"

	"github.com/stader-labs/stader-node/shared/services"
	"github.com/stader-labs/stader-node/shared/services/config"
	"github.com/stader-labs/stader-node/shared/utils/log"
)

// Config
var preSignedCooldown, _ = time.ParseDuration("1h")
var feeRecepientPollingInterval, _ = time.ParseDuration("5m")
var taskCooldown, _ = time.ParseDuration("10s")
var merkleProofsDownloadInterval, _ = time.ParseDuration("3h")

const (
	MaxConcurrentEth1Requests   = 200
	ManageFeeRecipientColor     = color.FgHiCyan
	MerkleProofsDownloaderColor = color.FgHiBlue
	ErrorColor                  = color.FgRed
	InfoColor                   = color.FgHiGreen
	blocksPerThreeEpoch         = 96
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

// run daemon
func run(c *cli.Context) error {

	// Handle the initial fee recipient file deployment
	err := deployDefaultFeeRecipientFile(c)
	if err != nil {
		return err
	}

	// Configure
	configureHTTP()

	w, err := services.GetWallet(c)
	if err != nil {
		return err
	}
	nodeAccount, err := w.GetNodeAccount()
	if err != nil {
		return err
	}

	err = services.WaitEthClientSynced(c, true)
	if err != nil {
		return err
	}
	err = services.WaitNodeRegistered(c, nodeAccount.Address, true)
	if err != nil {
		return err
	}

	bc, err := services.GetBeaconClient(c)
	if err != nil {
		return err
	}
	pnr, err := services.GetPermissionlessNodeRegistry(c)
	if err != nil {
		return err
	}

	// Initialize tasks
	manageFeeRecipient, err := newManageFeeRecipient(c, log.NewColorLogger(ManageFeeRecipientColor))
	if err != nil {
		return err
	}
	merkleProofsDownloader, err := NewMerkleProofsDownloader(c, log.NewColorLogger(MerkleProofsDownloaderColor))
	if err != nil {
		return err
	}

	// Initialize loggers
	errorLog := log.NewColorLogger(ErrorColor)
	infoLog := log.NewColorLogger(InfoColor)

	publicKey, err := stader.GetPublicKey(c)
	if err != nil {
		return err
	}

	// Wait group to handle the various threads
	wg := new(sync.WaitGroup)
	wg.Add(3)

	// validator presigned loop
	go func() {
		for {
			// Check the EC status
			err := services.WaitEthClientSynced(c, false) // Force refresh the primary / fallback EC status
			if err != nil {
				errorLog.Println(err)
				continue
			} else {
				// Check the BC status
				err := services.WaitBeaconClientSynced(c, false) // Force refresh the primary / fallback BC status
				if err != nil {
					errorLog.Println(err)
					continue
				}
			}

			cfg, err := services.GetConfig(c)
			if err != nil {
				errorLog.Printf("Failed to get config with error %s\n", err.Error())
				continue
			}

			network, ok := cfg.StaderNode.Network.Value.(cfgtypes.Network)
			if !ok {
				errorLog.Printf("Failed to get network from config: %s\n", cfg.StaderNode.Network.Value)
				continue
			}

			operatorId, err := node.GetOperatorId(pnr, nodeAccount.Address, nil)
			if err != nil {
				errorLog.Printf("Failed to get operator id: %s\n", err.Error())
				continue
			}

			// make a map of all validators actually registered with stader
			// user might just move the validator keys to the directory. we don't wanna send the presigned msg of them
			infoLog.Println("Building a map of user validators registered with stader")
			registeredValidators, validatorPubKeys, err := stdr.GetAllValidatorsRegisteredWithOperator(pnr, operatorId, nodeAccount.Address, nil)
			if err != nil {
				errorLog.Printf("Could not get all validators registered with operator %s with error %s\n", operatorId, err.Error())
				continue
			}

			infoLog.Printlnf("Found %d validators registered with operator %s", len(registeredValidators), operatorId)
			infoLog.Println("Starting a pass of the presign daemon!")

			currentHead, err := bc.GetBeaconHead()
			if err != nil {
				errorLog.Printf("Could not get beacon head with error %s\n", err.Error())
				continue
			}

			err = w.Reload()
			if err != nil {
				errorLog.Printf("Could not reload wallet: %s\n", err.Error())
				continue
			}

			preSignRegisteredMap, err := stader.BulkIsPresignedKeyRegistered(c, validatorPubKeys)
			if err != nil {
				errorLog.Printf("Could not bulk check presigned keys with error %s\n", err.Error())
				continue
			}

			pageNumber := 0
			pageSize := 5
			for {
				startIndex := pageNumber * pageSize
				if startIndex > len(validatorPubKeys) {
					break
				}
				endIndex := (pageNumber + 1) * pageSize
				if endIndex > len(validatorPubKeys) {
					endIndex = len(validatorPubKeys)
				}
				infoLog.Printf("Starting index: %d, End index: %d\n", startIndex, endIndex)

				validatorKeyBatch := validatorPubKeys[startIndex:endIndex]
				infoLog.Printf("Checking %d validator keys\n", len(validatorKeyBatch))

				preSignSendMessages := []stader_backend.PreSignSendApiRequestType{}

				for _, validatorPubKey := range validatorKeyBatch {
					infoLog.Printf("Checking validator pubkey %s\n", validatorPubKey.String())
					validatorKeyPair, err := w.GetValidatorKeyByPubkey(validatorPubKey)
					// log the errors and continue. dont need to sleep post an error
					if err != nil {
						errorLog.Printf("Could not find validator private key for %s with err: %s\n", validatorPubKey, err.Error())
						continue
					}

					validatorInfo, ok := registeredValidators[validatorPubKey]
					if !ok {
						errorLog.Printf("Validator pub key: %s not found in stader contracts\n", validatorPubKey)
						continue
					}
					if stdr.IsValidatorTerminal(validatorInfo) {
						errorLog.Printf("Validator pub key: %s is in terminal state in the stader contracts\n", validatorPubKey)
						continue
					}

					registeredPresign, ok := preSignRegisteredMap[validatorPubKey.String()]
					if !ok {
						errorLog.Printf("Could not query presign api to check if validator: %s is registered\n", validatorPubKey)
						continue
					}
					if registeredPresign {
						infoLog.Printf("Validator pub key: %s pre signed key already registered\n", validatorPubKey)
						continue
					} else {
						infoLog.Printf("Validator pub key: %s pre signed key not registered. Creating presigned message\n", validatorPubKey)
					}

					// check if validator has not yet been registered on beacon chain
					validatorStatus, err := bc.GetValidatorStatus(validatorPubKey, nil)
					if err != nil {
						errorLog.Printf("Error finding validator status for validator: %s with err: %s\n", validatorPubKey, err.Error())
						continue
					}
					if !validatorStatus.Exists {
						errorLog.Printf("Validator pub key: %s not found on beacon chain\n", validatorPubKey)
						continue
					}

					// check if validator is already in an exiting phase, then no point sending a pre-signed message
					if eth2.IsValidatorExiting(validatorStatus) {
						errorLog.Printf("Validator pub key: %s already exiting or exited with status %s", validatorPubKey, validatorStatus.Status)
						continue
					}

					exitEpoch := currentHead.Epoch

					signatureDomain, err := bc.GetExitDomainData(eth2types.DomainVoluntaryExit[:], network)
					if err != nil {
						errorLog.Printf("Failed to get the signature domain from beacon chain with err: %s\n", err.Error())
						continue
					}

					// get the presigned msg
					exitSignature, _, err := validator.GetSignedExitMessage(validatorKeyPair, validatorStatus.Index, exitEpoch, signatureDomain)
					if err != nil {
						errorLog.Printf("Failed to generate the SignedExitMessage for validator with beacon chain index: %d with err: %s\n", validatorStatus.Index, err.Error())
						continue
					}

					// encrypt the signature and srHash
					exitSignatureEncrypted, err := crypto.EncryptUsingPublicKey([]byte(exitSignature.String()), publicKey)
					if err != nil {
						errorLog.Printf("Failed to encrypt exit signature for validator: %s with err: %s\n", validatorPubKey, err.Error())
						continue
					}
					exitSignatureEncryptedString := crypto.EncodeBase64(exitSignatureEncrypted)

					// send it to the presigned api
					preSignSendMessages = append(preSignSendMessages, stader_backend.PreSignSendApiRequestType{
						Message: struct {
							Epoch          string `json:"epoch"`
							ValidatorIndex string `json:"validator_index"`
						}{
							Epoch:          strconv.FormatUint(exitEpoch, 10),
							ValidatorIndex: strconv.FormatUint(validatorStatus.Index, 10),
						},
						Signature:          exitSignatureEncryptedString,
						ValidatorPublicKey: validatorPubKey.String(),
					})
				}

				//fmt.Printf("Sending %d presigned messages to stader backend\n", len(preSignSendMessages))
				fmt.Printf("Sending %d presigned messages to stader backend\n", len(preSignSendMessages))
				//fmt.Printf("Pre-sign messages being sent are %v\n", preSignSendMessages)
				if len(preSignSendMessages) > 0 {
					res, err := stader.SendBulkPresignedMessageToStaderBackend(c, preSignSendMessages)
					if err != nil {
						errorLog.Printf("Sending bulk presigned message failed with %v\n", err.Error())
					} else {
						for pubKey, response := range *res {
							if response.Success {
								infoLog.Printf("Successfully sent the presigned message for validator: %s\n", pubKey)
							} else {
								errorLog.Printf("Failed to send the presigned api for validator: %s with err: %s\n", pubKey, response.Error)
							}
						}
					}
				}

				pageNumber += 1
			}

			infoLog.Printf("Done with the pass of presign daemon")
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
			time.Sleep(feeRecepientPollingInterval)
		}
		wg.Done()
	}()

	go func() {
		for {
			infoLog.Printlnf("Checking if there are any available merkle proofs to download")
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
					if err := merkleProofsDownloader.run(); err != nil {
						errorLog.Println(err)
					}
					time.Sleep(taskCooldown)
				}
			}
			infoLog.Printlnf("Done checking for merkle proofs to download")
			time.Sleep(merkleProofsDownloadInterval)
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
