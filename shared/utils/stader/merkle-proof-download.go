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
	//fmt.Printf("Making api call to %s\n", fmt.Sprintf(stader.MerkleProofAggregateGetterApi, operator.Hex()))
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()
	if res.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("error while getting ALL merkle proofs for operator %s", operator.Hex())
	}

	var allMerkleProofs []*stader_backend.CycleMerkleProofs
	err = json.NewDecoder(res.Body).Decode(&allMerkleProofs)
	if err != nil {
		return nil, err
	}
	return []*stader_backend.CycleMerkleProofs{
		{
			Cycle: 1,
			Proof: []string{},
			Eth:   "1000000000000000000",
			Sd:    "1000000000000000000",
			Root:  "0x7c13648108d8172a507408b98160a0f09d22eb409c4911c7cab4a24b59c90042",
		},
		{
			Cycle: 2,
			Proof: []string{},
			Eth:   "2000000000000000000",
			Sd:    "2000000000000000000",
			Root:  "0xf2c90d506b4a46d0fa5b7517ec6df08614cc507c52212d0f7da0df38526ccf8b",
		},
		{
			Cycle: 3,
			Proof: []string{},
			Eth:   "3000000000000000000",
			Sd:    "3000000000000000000",
			Root:  "0xaa6e55e3ab5e25bd672c48f695de4321cdae4b591dd8607dd5f8e3983266c1b3",
		},
	}, nil
}
