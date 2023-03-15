package node

import (
	"fmt"
	"github.com/stader-labs/stader-node/shared/services"
	"github.com/stader-labs/stader-node/shared/types/api"
	"github.com/stader-labs/stader-node/shared/types/stader-backend"
	"github.com/stader-labs/stader-node/shared/utils/crypto"
	"github.com/stader-labs/stader-node/shared/utils/stader"
	"github.com/stader-labs/stader-node/shared/utils/validator"
	"github.com/stader-labs/stader-node/stader-lib/types"
	"github.com/urfave/cli"
	eth2types "github.com/wealdtech/go-eth2-types/v2"
	"strconv"
)

func canSendPresignedMsg(c *cli.Context, validatorPubKey types.ValidatorPubkey) (*api.CanSendPresignedMsgResponse, error) {
	canSendPresignedMsgResponse := api.CanSendPresignedMsgResponse{}

	bc, err := services.GetBeaconClient(c)
	if err != nil {
		return nil, err
	}

	// check if validator is present by querying validator index
	validatorIndex, err := bc.GetValidatorIndex(validatorPubKey)
	if err != nil {
		return nil, err
	}
	if validatorIndex == 0 {
		canSendPresignedMsgResponse.ValidatorNotRegistered = true
		return &canSendPresignedMsgResponse, nil
	}

	// check if already registered
	isRegistered, err := stader.IsPresignedKeyRegistered(validatorPubKey)
	if err != nil {
		return nil, err
	}
	if isRegistered {
		canSendPresignedMsgResponse.ValidatorPreSignKeyAlreadyRegistered = true
		return &canSendPresignedMsgResponse, nil
	}

	return &canSendPresignedMsgResponse, nil
}

func sendPresignedMsg(c *cli.Context, validatorPubKey types.ValidatorPubkey) (*api.SendPresignedMsgResponse, error) {
	// generate presigned message
	w, err := services.GetWallet(c)
	if err != nil {
		return nil, err
	}
	bc, err := services.GetBeaconClient(c)
	if err != nil {
		return nil, err
	}

	validatorPrivateKey, err := w.GetValidatorKeyByPubkey(validatorPubKey)
	if err != nil {
		return nil, err
	}

	currentHead, err := bc.GetBeaconHead()
	if err != nil {
		return nil, err
	}

	response := api.SendPresignedMsgResponse{}

	validatorStatus, err := bc.GetValidatorStatus(validatorPubKey, nil)
	if err != nil {
		return nil, err
	}

	// exit epoch should be > activation_epoch + 256
	// exit epoch should be > current epoch
	exitEpoch := currentHead.Epoch + 1
	epochsSinceActivation := currentHead.Epoch - validatorStatus.ActivationEpoch
	if epochsSinceActivation < 256 {
		exitEpoch = exitEpoch + (256 - epochsSinceActivation)
	}
	fmt.Printf("exitEpoch is %d\n", exitEpoch)

	signatureDomain, err := bc.GetDomainData(eth2types.DomainVoluntaryExit[:], exitEpoch, false)
	if err != nil {
		return nil, err
	}

	exitMsg, srHash, err := validator.GetSignedExitMessage(validatorPrivateKey, validatorStatus.Index, exitEpoch, signatureDomain)
	if err != nil {
		return nil, err
	}
	fmt.Printf("exitMsg is %s\n", exitMsg.String())
	fmt.Printf("srHash is %s\n", srHash[:])

	// get the public key
	fmt.Printf("Getting public key!")
	publicKey, err := stader.GetPublicKey()
	if err != nil {
		return nil, err
	}
	fmt.Printf("Got the public key! %v\n", publicKey)

	fmt.Println("Encrypting exitSignature!")
	exitSignatureEncrypted, err := crypto.EncryptUsingPublicKey([]byte(exitMsg.String()), publicKey)
	if err != nil {
		return nil, err
	}
	fmt.Printf("exitSigntaure encrypted is %s\n", exitSignatureEncrypted)
	exitSignatureEncryptedString := crypto.EncodeBase64(exitSignatureEncrypted)
	fmt.Printf("base64 encoded exit signature is %s\n", exitSignatureEncryptedString)

	fmt.Println("Encrypting message hash")
	messageHashEncrypted, err := crypto.EncryptUsingPublicKey(srHash[:], publicKey)
	if err != nil {
		return nil, err
	}
	fmt.Println("Encrypted message hash")
	fmt.Printf("message hash encrypted is %s\n", messageHashEncrypted)
	messageHashEncryptedString := crypto.EncodeBase64(messageHashEncrypted)
	fmt.Printf("base64 encoded message hash is %s\n", messageHashEncryptedString)

	fmt.Printf("Sending the presigned message\n")
	// encrypt the presigned exit message object
	preSignedMessageRequest := stader_backend.PreSignSendApiRequestType{
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
	}

	fmt.Printf("preSignedMessageRequest is %s\n", preSignedMessageRequest)

	res, err := stader.SendPresignedMessageToStaderBackend(preSignedMessageRequest)
	if err != nil {
		return nil, err
	}
	fmt.Printf("res of send presigned message to stader backend is %v\n", res)

	response.ValidatorIndex = validatorStatus.Index
	response.ValidatorPubKey = validatorPubKey
	response.ExitEpoch = exitEpoch
	response.SignedMsg = exitMsg

	return &response, nil
}
