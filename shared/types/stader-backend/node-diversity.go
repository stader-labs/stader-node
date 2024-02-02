package stader_backend

type NodeDiversityRequest struct {
	Signature string `json:"signature"`
	Message   string `json:"message"`
}

type NodeDiversity struct {
	ExecutionClient string `json:"executionClient"`
	ConsensusClient string `json:"consensusClient"`
	NodeAddress     string `json:"nodeAddress"`
	NodePublicKey   string `json:"nodePublicKey"`
	Relays          string `json:"relays"`
}
