package stader

import (
	"encoding/json"
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/stader-labs/stader-node/shared/services/stader"
	stader_backend "github.com/stader-labs/stader-node/shared/types/stader-backend"
	"github.com/stader-labs/stader-node/shared/utils/net"
	"net/http"
)

func GetAllMerkleProofsForOperator(operator common.Address) ([]*stader_backend.CycleMerkleProofs, error) {
	res, err := net.MakeGetRequest(fmt.Sprintf(stader.MerkleProofAggregateGetterApi, operator.Hex()), struct{}{})
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()
	if res.StatusCode == http.StatusBadRequest {
		return []*stader_backend.CycleMerkleProofs{}, nil
	}
	if res.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("error while getting all merkle proofs for operator %s", operator.Hex())
	}

	var allMerkleProofs []*stader_backend.CycleMerkleProofs
	err = json.NewDecoder(res.Body).Decode(&allMerkleProofs)
	if err != nil {
		return nil, err
	}
	return allMerkleProofs, nil
}
