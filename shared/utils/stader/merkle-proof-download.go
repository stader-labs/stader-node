package stader

import (
	"encoding/json"
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	stader_backend "github.com/stader-labs/stader-node/shared/types/stader-backend"
	"github.com/stader-labs/stader-node/shared/utils/net"
)

const merkleProofGetterApi = "https://v6s3vqe7va.execute-api.us-east-1.amazonaws.com/prod/merklesForElRewards/%d/%s"

// TODO - akhilesh - get an api which will return the all merkle proofs for all cycles for a given operator
func GetCycleMerkleProofsForOperator(cycle int64, operator common.Address) (*stader_backend.CycleMerkleProofs, error) {
	res, err := net.MakeGetRequest(fmt.Sprintf(merkleProofGetterApi, cycle, operator.Hex()), struct{}{})
	fmt.Printf("Making api call to %s\n", fmt.Sprintf(merkleProofGetterApi, cycle, operator.Hex()))
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()
	var cycleMerkleProofs stader_backend.CycleMerkleProofs
	err = json.NewDecoder(res.Body).Decode(&cycleMerkleProofs)
	if err != nil {
		return nil, err
	}
	return &cycleMerkleProofs, nil
}
