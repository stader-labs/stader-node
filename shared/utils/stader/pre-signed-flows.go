package stader

import (
	"crypto/rsa"
	"encoding/json"
	"github.com/stader-labs/stader-node/shared/services/stader"
	stader_backend "github.com/stader-labs/stader-node/shared/types/stader-backend"
	"github.com/stader-labs/stader-node/shared/utils/crypto"
	"github.com/stader-labs/stader-node/shared/utils/net"
	"github.com/stader-labs/stader-node/stader-lib/types"
)

func SendPresignedMessageToStaderBackend(preSignedMessage stader_backend.PreSignSendApiRequestType) (*stader_backend.PreSignSendApiResponseType, error) {
	res, err := net.MakePostRequest(stader.PreSignSendApi, preSignedMessage)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

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

	res, err := net.MakePostRequest(stader.PreSignCheckApi, preSignCheckRequest)

	defer res.Body.Close()
	var preSignCheckResponse stader_backend.PreSignCheckApiResponseType
	err = json.NewDecoder(res.Body).Decode(&preSignCheckResponse)
	if err != nil {
		return false, err
	}
	return preSignCheckResponse.Value, nil
}

func GetPublicKey() (*rsa.PublicKey, error) {
	// get public key from api
	res, err := net.MakeGetRequest(stader.PublicKeyApi, struct{}{})
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

	publicKey, err := crypto.BytesToPublicKey(decodedPublicKey)
	if err != nil {
		return nil, err
	}

	return publicKey, nil
}
