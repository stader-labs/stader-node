package node

import (
	"fmt"
	"github.com/stader-labs/stader-node/shared/services"
	"github.com/stader-labs/stader-node/shared/types/api"
	"github.com/stader-labs/stader-node/shared/types/stader-backend"
	"github.com/stader-labs/stader-node/shared/utils/crypto"
	"github.com/stader-labs/stader-node/shared/utils/eth2"
	"github.com/stader-labs/stader-node/shared/utils/stader"
	"github.com/stader-labs/stader-node/shared/utils/validator"
	"github.com/stader-labs/stader-node/stader-lib/node"
	"github.com/stader-labs/stader-node/stader-lib/types"
	"github.com/urfave/cli"
	eth2types "github.com/wealdtech/go-eth2-types/v2"
	"strconv"
)

func canSendPresignedMsg(c *cli.Context, validatorPubKey types.ValidatorPubkey) (*api.CanSendPresignedMsgResponse, error) {
	if err := services.RequireNodeWallet(c); err != nil {
		return nil, err
	}
	pnr, err := services.GetPermissionlessNodeRegistry(c)
	if err != nil {
		return nil, err
	}
	bc, err := services.GetBeaconClient(c)
	if err != nil {
		return nil, err
	}
	w, err := services.GetWallet(c)
	if err != nil {
		return nil, err
	}
	nodeAccount, err := w.GetNodeAccount()
	if err != nil {
		return nil, err
	}

	canSendPresignedMsgResponse := api.CanSendPresignedMsgResponse{}

	// TODO - check if the validator is registered with this operator
	// TODO - check if the validator key is present in the system

	operatorId, err := node.GetOperatorId(pnr, nodeAccount.Address, nil)
	if err != nil {
		return nil, err
	}
	if operatorId.Int64() == 0 {
		canSendPresignedMsgResponse.OperatorNotRegistered = true
		return &canSendPresignedMsgResponse, nil
	}

	// check if validator is present by querying validator index
	validatorStatus, err := bc.GetValidatorStatus(validatorPubKey, nil)
	if validatorStatus.Index == 0 || err != nil {
		canSendPresignedMsgResponse.ValidatorNotRegistered = true
		return &canSendPresignedMsgResponse, nil
	}

	if eth2.IsValidatorExiting(validatorStatus) {
		canSendPresignedMsgResponse.ValidatorIsNotActive = true
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

	// exit epoch should be > current epoch
	exitEpoch := currentHead.Epoch

	signatureDomain, err := bc.GetDomainData(eth2types.DomainVoluntaryExit[:], exitEpoch, false)
	if err != nil {
		return nil, err
	}

	exitMsg, _, err := validator.GetSignedExitMessage(validatorPrivateKey, validatorStatus.Index, exitEpoch, signatureDomain)
	if err != nil {
		return nil, err
	}

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

	// encrypt the presigned exit message object
	preSignedMessageRequest := stader_backend.PreSignSendApiRequestType{
		Message: struct {
			Epoch          string `json:"epoch"`
			ValidatorIndex string `json:"validator_index"`
		}{
			Epoch:          strconv.FormatUint(exitEpoch, 10),
			ValidatorIndex: strconv.FormatUint(validatorStatus.Index, 10),
		},
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
