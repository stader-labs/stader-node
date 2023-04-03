package node

import (
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/stader-labs/stader-node/shared/services"
	"github.com/stader-labs/stader-node/shared/types/api"
	"github.com/stader-labs/stader-node/shared/types/stader-backend"
	"github.com/stader-labs/stader-node/shared/utils/crypto"
	"github.com/stader-labs/stader-node/shared/utils/eth2"
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
	validatorStatus, err := bc.GetValidatorStatus(validatorPubKey, nil)
	if validatorStatus.Index == 0 || err != nil {
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

	if eth2.IsValidatorExiting(validatorStatus) {
		canSendPresignedMsgResponse.ValidatorIsExiting = true
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

	signatureDomain, err := bc.GetDomainData(eth2types.DomainVoluntaryExit[:], exitEpoch, false)
	if err != nil {
		return nil, err
	}

	exitMsg, srHash, err := validator.GetSignedExitMessage(validatorPrivateKey, validatorStatus.Index, exitEpoch, signatureDomain)
	if err != nil {
		return nil, err
	}
	srHashHex := common.Bytes2Hex(srHash[:])

	// get the public key
	publicKey, err := stader.GetPublicKey()
	if err != nil {
		return nil, err
	}

	exitSignatureEncrypted, err := crypto.EncryptUsingPublicKey([]byte(exitMsg.String()), publicKey)
	if err != nil {
		return nil, err
	}
	exitSignatureEncryptedString := crypto.EncodeBase64(exitSignatureEncrypted)

	messageHashEncrypted, err := crypto.EncryptUsingPublicKey([]byte(srHashHex), publicKey)
	if err != nil {
		return nil, err
	}

	messageHashEncryptedString := crypto.EncodeBase64(messageHashEncrypted)

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

	res, err := stader.SendPresignedMessageToStaderBackend(preSignedMessageRequest)
	if err != nil {
		return nil, err
	}
	if !res.Success {
		return nil, fmt.Errorf("send-presigned-message failed: %s\n", res.Message)
	}

	response.ValidatorIndex = validatorStatus.Index
	response.ValidatorPubKey = validatorPubKey
	response.ExitEpoch = exitEpoch
	response.SignedMsg = exitMsg

	return &response, nil
}
