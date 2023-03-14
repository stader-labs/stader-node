package node

import (
	"crypto/rsa"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"github.com/stader-labs/stader-node/shared/services"
	"github.com/stader-labs/stader-node/shared/types/api"
	"github.com/stader-labs/stader-node/shared/types/stader-backend"
	"github.com/stader-labs/stader-node/shared/utils/crypto"
	"github.com/stader-labs/stader-node/shared/utils/net"
	"github.com/stader-labs/stader-node/shared/utils/validator"
	"github.com/stader-labs/stader-node/stader-lib/types"
	"github.com/urfave/cli"
	eth2types "github.com/wealdtech/go-eth2-types/v2"
)

// TODO - refactor these urls somehow
const preSignSendApi = "https://v6s3vqe7va.execute-api.us-east-1.amazonaws.com/prod/presign"
const preSignCheckApi = "https://v6s3vqe7va.execute-api.us-east-1.amazonaws.com/prod/msgSubmitted"
const publicKeyApi = "https://v6s3vqe7va.execute-api.us-east-1.amazonaws.com/prod/publicKey"

func sendPresignedMessageToStaderBackend(preSignedMessage stader_backend.PreSignSendApiRequestType) (*stader_backend.PreSignSendApiResponseType, error) {
	res, err := net.MakePostRequest(preSignSendApi, preSignedMessage)
	if err != nil {
		return nil, err
	}

	var preSignSendResponse stader_backend.PreSignSendApiResponseType
	err = json.NewDecoder(res.Body).Decode(&preSignSendResponse)
	if err != nil {
		return nil, err
	}

	return &preSignSendResponse, nil
}

func isPresignedKeyRegistered(validatorPubKey types.ValidatorPubkey) (bool, error) {
	//// check if it is already there
	preSignCheckRequest := stader_backend.PreSignCheckApiRequestType{
		ValidatorPublicKey: validatorPubKey.String(),
	}

	preSignCheckRes, err := net.MakePostRequest(preSignCheckApi, preSignCheckRequest)

	defer preSignCheckRes.Body.Close()
	var preSignCheckResponse stader_backend.PreSignCheckApiResponseType
	err = json.NewDecoder(preSignCheckRes.Body).Decode(&preSignCheckResponse)
	if err != nil {
		return false, err
	}
	return preSignCheckResponse.Value, nil
}

func getPublicKey() (*rsa.PublicKey, error) {
	// get public key from api
	res, err := net.MakeGetRequest(publicKeyApi, struct{}{})
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	var publicKeyResponse stader_backend.PublicKeyApiResponse
	err = json.NewDecoder(res.Body).Decode(&publicKeyResponse)
	if err != nil {
		return nil, err
	}

	decodedPublicKey, err := base64.StdEncoding.DecodeString(publicKeyResponse.Value)
	if err != nil {
		return nil, err
	}

	publicKey, err := crypto.BytesToPublicKey(decodedPublicKey)
	if err != nil {
		return nil, err
	}

	return publicKey, nil
}

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
	isRegistered, err := isPresignedKeyRegistered(validatorPubKey)
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
	fmt.Printf("validator private key is %s\n", validatorPrivateKey)

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
		exitEpoch = exitEpoch + (256 - epochsSinceActivation) + 1
	}

	signatureDomain, err := bc.GetDomainData(eth2types.DomainVoluntaryExit[:], exitEpoch)
	if err != nil {
		return nil, err
	}

	exitMsg, srHash, err := validator.GetSignedExitMessage(validatorPrivateKey, validatorStatus.Index, exitEpoch, signatureDomain)
	if err != nil {
		return nil, err
	}

	// get the public key
	publicKey, err := getPublicKey()
	if err != nil {
		return nil, err
	}

	exitSignatureEncrypted, err := crypto.EncryptUsingPublicKey(exitMsg.Bytes(), publicKey)
	if err != nil {
		return nil, err
	}

	messageHashEncrypted, err := crypto.EncryptUsingPublicKey(srHash[:], publicKey)
	if err != nil {
		return nil, err
	}

	fmt.Printf("encrypted oaep message hash is %s\n", messageHashEncrypted)

	// encrypt the presigned exit message object
	preSignedMessageRequest := stader_backend.PreSignSendApiRequestType{
		Message: struct {
			Epoch          uint64 `json:"epoch"`
			ValidatorIndex uint64 `json:"validator_index"`
		}{
			Epoch:          exitEpoch,
			ValidatorIndex: validatorStatus.Index,
		},
		MessageHash:        messageHashEncrypted,
		Signature:          exitSignatureEncrypted,
		ValidatorPublicKey: validatorPubKey.Hex(),
	}

	_, err = sendPresignedMessageToStaderBackend(preSignedMessageRequest)
	if err != nil {
		return nil, err
	}

	response.ValidatorIndex = validatorStatus.Index
	response.ValidatorPubKey = validatorPubKey
	response.ExitEpoch = exitEpoch
	response.SignedMsg = exitMsg

	return &response, nil
}
