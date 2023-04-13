package node

//
//// TODO - bchain - complete this abstraction
//
//import (
//	"github.com/stader-labs/stader-node/shared/services"
//	"github.com/stader-labs/stader-node/shared/services/wallet"
//	stader_backend "github.com/stader-labs/stader-node/shared/types/stader-backend"
//	"github.com/stader-labs/stader-node/shared/utils/crypto"
//	"github.com/stader-labs/stader-node/shared/utils/eth2"
//	"github.com/stader-labs/stader-node/shared/utils/log"
//	"github.com/stader-labs/stader-node/shared/utils/validator"
//	"github.com/stader-labs/stader-node/shared/utils/stader"
//	stader_contracts "github.com/stader-labs/stader-node/stader-lib/stader"
//
//	"github.com/stader-labs/stader-node/stader-lib/types"
//	"github.com/urfave/cli"
//	eth2types "github.com/wealdtech/go-eth2-types/v2"
//	"strconv"
//	"time"
//)
//
//type preSignDaemon struct {
//	c   *cli.Context
//	infoLog log.ColorLogger
//	errorLog log.ColorLogger
//	prn *stader_contracts.PermissionlessNodeRegistryContractManager
//	w *wallet.Wallet
//	bc *services.BeaconClientManager
//}
//
//func NewPreSignDaemon(c *cli.Context, infoLog log.ColorLogger, errorLog log.ColorLogger) (*preSignDaemon, error) {
//	prn, err := services.GetPermissionlessNodeRegistry(c)
//	if err != nil {
//		return nil, err
//	}
//	w, err := services.GetWallet(c)
//	if err != nil {
//		return nil, err
//	}
//	bc, err := services.GetBeaconClient(c)
//	if err != nil {
//		return nil, err
//	}
//
//	return &preSignDaemon{
//		c: c,
//		infoLog: infoLog,
//		errorLog: errorLog,
//		prn: prn,
//		w: w,
//		bc: bc,
//	}, nil
//}
//
//func (psd *preSignDaemon) run() error {
//	psd.infoLog.Println("Starting a pass of the presign daemon!")
//	walletIndex := psd.w.GetNextAccount()
//	noOfBatches := walletIndex / uint(preSignBatchSize)
//	batchIndex := 0
//
//	currentHead, err := psd.bc.GetBeaconHead()
//	if err != nil {
//		panic("not able to communicate with beacon chain!")
//	}
//
//	for i := uint(0); i <= noOfBatches; i++ {
//		for j := batchIndex; j < batchIndex+preSignBatchSize && j < int(walletIndex); j++ {
//			psd.infoLog.Printf("Checking validator index %d\n", j)
//			validatorPrivateKey, err := psd.w.GetValidatorKeyAt(uint(j))
//			// log the errors and continue. dont need to sleep post an error
//			if err != nil {
//				psd.errorLog.Printf("Could not find validator private key for validator index %d\n", j)
//				continue
//			}
//
//			validatorPubKey := types.BytesToValidatorPubkey(validatorPrivateKey.PublicKey().Marshal())
//			psd.infoLog.Printf("Checking validator Pub key: %s\n", validatorPubKey.String())
//
//			// check if validator has not yet been registered
//			validatorStatus, err := psd.bc.GetValidatorStatus(validatorPubKey, nil)
//			if err != nil {
//				psd.errorLog.Printf("Could not find validator status for validator pub key: %s\n", validatorPubKey)
//				continue
//			}
//			if !eth2.IsValidatorActive(validatorStatus) {
//				psd.errorLog.Printf("Validator is not active yet. Validator pub key: %s\n", validatorPubKey)
//				continue
//			}
//
//			// check if the presigned message has been registered. if it has been registered, then continue
//			isRegistered, err := stader.IsPresignedKeyRegistered(validatorPubKey)
//			if isRegistered {
//				psd.infoLog.Printf("Validator pub key: %s already registered\n", validatorPubKey)
//				continue
//			} else if err != nil {
//				psd.errorLog.Printf("Could not query presign api to check if validator: %s is registered\n", validatorPubKey)
//			}
//
//			exitEpoch := currentHead.Epoch
//
//			signatureDomain, err := psd.bc.GetDomainData(eth2types.DomainVoluntaryExit[:], exitEpoch, false)
//			if err != nil {
//				psd.errorLog.Printf("Failed to get the signature domain from beacon chain\n")
//				continue
//			}
//
//			// get the presigned msg
//			exitSignature, _, err := validator.GetSignedExitMessage(validatorPrivateKey, validatorStatus.Index, exitEpoch, signatureDomain)
//			if err != nil {
//				psd.errorLog.Printf("Failed to generate the SignedExitMessage for validator with beacon chain index: %d\n", validatorStatus.Index)
//				continue
//			}
//
//			// encrypt the signature and srHash
//			exitSignatureEncrypted, err := crypto.EncryptUsingPublicKey([]byte(exitSignature.String()), publicKey)
//			if err != nil {
//				psd.errorLog.Printf("Failed to encrypt exit signature for validator: %s\n", validatorPubKey)
//				continue
//			}
//			exitSignatureEncryptedString := crypto.EncodeBase64(exitSignatureEncrypted)
//
//			// send it to the presigned api
//			backendRes, err := stader.SendPresignedMessageToStaderBackend(stader_backend.PreSignSendApiRequestType{
//				Message: struct {
//					Epoch          string `json:"epoch"`
//					ValidatorIndex string `json:"validator_index"`
//				}{
//					Epoch:          strconv.FormatUint(exitEpoch, 10),
//					ValidatorIndex: strconv.FormatUint(validatorStatus.Index, 10),
//				},
//				Signature:          exitSignatureEncryptedString,
//				ValidatorPublicKey: validatorPubKey.String(),
//			})
//			if !backendRes.Success {
//				psd.errorLog.Printf("Failed to send the presigned api: %s\n", backendRes.Message)
//				continue
//			} else if backendRes.Success {
//				psd.errorLog.Printf("Successfully sent the presigned message for validator: %s\n", validatorPubKey)
//				continue
//			} else if err != nil {
//				psd.errorLog.Printf("Sending presigned message failed with %v\n", err)
//				continue
//			}
//
//			time.Sleep(preSignedBatchCooldown)
//		}
//
//		batchIndex = batchIndex + preSignBatchSize
//	}
//
//	psd.errorLog.Printf("Done with the pass of presign daemon")
//
//	return nil
//}
