package stader_backend

type PreSignCheckApiRequestType struct {
	ValidatorPublicKey string `json:"validatorPublicKey"`
}

type PreSignCheckApiResponseType struct {
	Value bool `json:"value"`
}

type PreSignSendApiResponseType struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
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
type BulkPreSignSendApiResponseType = []PreSignSendApiResponseType

type BulkPreSignCheckApiRequestType = []string
type BulkPreSignCheckApiResponseType = map[string]bool

type PublicKeyApiResponse struct {
	Value string `json:"value"`
}
