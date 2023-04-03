package stader

import (
	"crypto/rsa"
	"encoding/json"
	"fmt"
	stader_backend "github.com/stader-labs/stader-node/shared/types/stader-backend"
	"github.com/stader-labs/stader-node/shared/utils/crypto"
	"github.com/stader-labs/stader-node/shared/utils/net"
	"github.com/stader-labs/stader-node/stader-lib/types"
)

// TODO - refactor these urls somehow
const preSignSendApi = "https://v6s3vqe7va.execute-api.us-east-1.amazonaws.com/prod/presign"
const preSignCheckApi = "https://v6s3vqe7va.execute-api.us-east-1.amazonaws.com/prod/msgSubmitted"
const publicKeyApi = "https://v6s3vqe7va.execute-api.us-east-1.amazonaws.com/prod/publicKey"

func SendPresignedMessageToStaderBackend(preSignedMessage stader_backend.PreSignSendApiRequestType) (*stader_backend.PreSignSendApiResponseType, error) {
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

func IsPresignedKeyRegistered(validatorPubKey types.ValidatorPubkey) (bool, error) {
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

func GetPublicKey() (*rsa.PublicKey, error) {
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

	decodedPublicKey, err := crypto.DecodeBase64(publicKeyResponse.Value)
	if err != nil {
		return nil, err
	}
	fmt.Printf("Decoded public key is %s\n", string(decodedPublicKey))

	publicKey, err := crypto.BytesToPublicKey(decodedPublicKey)
	if err != nil {
		return nil, err
	}

	return publicKey, nil
}
