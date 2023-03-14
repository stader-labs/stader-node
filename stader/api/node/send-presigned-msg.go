package node

import (
	"crypto/rsa"
	"crypto/x509"
	"encoding/base64"
	"encoding/json"
	"encoding/pem"
	"fmt"
	"github.com/stader-labs/stader-node/shared/services"
	"github.com/stader-labs/stader-node/shared/types/api"
	"github.com/stader-labs/stader-node/shared/utils/validator"
	"github.com/stader-labs/stader-node/stader-lib/types"
	"github.com/urfave/cli"
	eth2types "github.com/wealdtech/go-eth2-types/v2"
	"net/http"
)

// TODO - refactor

const preSignSendApi = "https://v6s3vqe7va.execute-api.us-east-1.amazonaws.com/prod/presign"
const preSignCheckApi = "https://v6s3vqe7va.execute-api.us-east-1.amazonaws.com/prod/msgSubmitted"
const publicKeyApi = "https://v6s3vqe7va.execute-api.us-east-1.amazonaws.com/prod/publicKey"

type PreSignCheckApiRequestType struct {
	ValidatorPublicKey string `json:"validatorPublicKey"`
}

type PreSignCheckApiResponseType struct {
	Value bool `json:"value"`
}

type PreSignSendApiResponseType struct {
}

type PreSignedSendApiRequest struct {
	Data []byte `json:"data"`
}

type PreSignSendUnEncryptedType struct {
	Message struct {
		Epoch          string `json:"epoch"`
		ValidatorIndex string `json:"validator_index"`
	} `json:"message"`
	MessageHash        string `json:"messageHash"`
	Signature          string `json:"signature"`
	ValidatorPublicKey string `json:"validatorPublicKey"`
}

type PublicKeyApiResponse struct {
	Value string `json:"value"`
}

func BytesToPublicKey(pub []byte) (*rsa.PublicKey, error) {
	block, _ := pem.Decode(pub)
	b := block.Bytes
	var err error

	ifc, err := x509.ParsePKIXPublicKey(b)
	if err != nil {
		fmt.Printf("Error using x509.ParsePKIXPublicKey %v\n", err)
		return nil, err
	}
	key, ok := ifc.(*rsa.PublicKey)
	if !ok {
		fmt.Printf("Error validating public key %v\n", err)
		return nil, err
	}
	return key, nil
}

func canSendPresignedMsg(c *cli.Context, validatorPubKey types.ValidatorPubkey) error {

	return nil
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
		exitEpoch = exitEpoch + (256 - epochsSinceActivation) + 1
	}

	signatureDomain, err := bc.GetDomainData(eth2types.DomainVoluntaryExit[:], exitEpoch)
	if err != nil {
		return nil, err
	}

	_, _, err = validator.GetSignedExitMessage(validatorPrivateKey, validatorStatus.Index, exitEpoch, signatureDomain)
	if err != nil {
		return nil, err
	}

	// get the public key
	res, err := http.Get(publicKeyApi)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()
	var publicKeyResponse PublicKeyApiResponse
	err = json.NewDecoder(res.Body).Decode(&publicKeyResponse)
	if err != nil {
		return nil, err
	}
	fmt.Printf("public key is %s\n", publicKeyResponse.Value)

	// encrypt using the public key
	encodedPublicKey, err := base64.StdEncoding.DecodeString(publicKeyResponse.Value)
	if err != nil {
		panic(err)
	}

	publicKey, err := BytesToPublicKey(encodedPublicKey)
	if err != nil {
		return nil, err
	}
	fmt.Printf("public key is %v\n", publicKey)

	//// check if it is already there
	//preSignCheckRequest := PreSignCheckApiRequestType{
	//	ValidatorPublicKey: validatorPubKey.String(),
	//}
	//requestData, err := json.Marshal(preSignCheckRequest)
	//if err != nil {
	//	return nil, err
	//}
	//
	//preSignCheckRes, err := http.Post(preSignCheckApi, "application/json", bytes.NewBuffer(requestData))
	//if err != nil {
	//	return nil, err
	//}
	//defer preSignCheckRes.Body.Close()
	//var preSignCheckResponse PreSignCheckApiResponseType
	//err = json.NewDecoder(preSignCheckRes.Body).Decode(&preSignCheckResponse)
	//if err != nil {
	//	return nil, err
	//}
	//fmt.Printf("PreSigned check output %t\n", preSignCheckResponse.Value)

	// encrypt the presigned exit message object

	// send the encrypted presigned exit message object to the api

	return &response, nil
}
