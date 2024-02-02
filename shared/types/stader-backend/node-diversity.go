package stader_backend

type NodeDiversityRequest struct {
	Signature string         `json:"signature"`
	Message   *NodeDiversity `json:"message"`
}

type NodeDiversity struct {
	ExecutionClient string `json:"executionClient"`
	ConsensusClient string `json:"consensusClient"`
	NodeAddress     string `json:"nodeAddress"`
	NodePublicKey   string `json:"nodePublicKey"`
	Relays          string `json:"relays"`
}

type NodeDiversityResponseType struct {
	Success bool   `json:"success"`
	Error   string `json:"error"`
}
