package stader_backend

import "github.com/stader-labs/stader-node/stader-lib/types"

type PreSignCheckApiRequestType struct {
	ValidatorPublicKey string `json:"validatorPublicKey"`
}

type PreSignCheckApiResponseType struct {
	Value bool `json:"value"`
}

type PreSignSendApiResponseType struct {
	Success bool   `json:"success"`
	Error   string `json:"error"`
}

type PreSignSendApiRequestType struct {
	Message struct {
		Epoch          string `json:"epoch"`
		ValidatorIndex string `json:"validator_index"`
	} `json:"message"`
	Signature          string `json:"signature"`
	ValidatorPublicKey string `json:"validatorPublicKey"`
}

type BulkPreSignSendApiRequestType = []PreSignSendApiRequestType
type BulkPreSignSendApiResponseType = map[string]PreSignSendApiResponseType

type BulkPreSignCheckApiRequestType struct {
	ValidatorPubKeys []types.ValidatorPubkey `json:"pubkeys"`
}

type BulkPreSignCheckApiResponseType = map[string]bool

type PublicKeyApiResponse struct {
	Value string `json:"value"`
}
