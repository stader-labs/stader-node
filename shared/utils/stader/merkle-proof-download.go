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
	if res.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("error while getting all merkle proofs for operator %s", operator.Hex())
	}

	var allMerkleProofs []*stader_backend.CycleMerkleProofs
	err = json.NewDecoder(res.Body).Decode(&allMerkleProofs)
	if err != nil {
		return nil, err
	}
	return allMerkleProofs, nil
	//return []*stader_backend.CycleMerkleProofs{
	//	{
	//		Cycle: 1,
	//		Proof: []string{
	//			"0x7899e58593b56ccecf51b92a0480a83450306b3eae",
	//		},
	//		Eth:  "1000000000000000000",
	//		Sd:   "1000000000000000000",
	//		Root: "0x813f98cbd1ddea7097db5c5eca5da343a4625ecdbab384ab681ef6df307dfaa2",
	//	},
	//	{
	//		Cycle: 2,
	//		Proof: []string{
	//			"0xf95a29cc9b510fc71767250ad824809eb15bcc79b1ba9a89816671a1579635d4",
	//		},
	//		Eth:  "2000000000000000000",
	//		Sd:   "2000000000000000000",
	//		Root: "0xeb6a5b823dbe72a46a224e32fd396e634ce4439fab72791627d02a67fb04cb1a",
	//	},
	//	{
	//		Cycle: 3,
	//		Proof: []string{
	//			"0x90f7004ee52d2a946a61280aa31fae56c1138e76f69d8cfebda590c1369590fe",
	//		},
	//		Eth:  "1000000000000000000",
	//		Sd:   "2000000000000000000",
	//		Root: "0x640675395b978ca7cac723ca7e5115dc187d0b6b8d1980afb057ee47cedf384f",
	//	},
	//	{
	//		Cycle: 4,
	//		Proof: []string{
	//			"0x4d48b42e0f17246aa92a5d6e4e18fb6fa613379fe104e9ca54f3c166d64f5575",
	//		},
	//		Eth:  "4000000000000000000",
	//		Sd:   "3000000000000000000",
	//		Root: "0x87c75479c7888db177c43ab658b532f13a07c4d8e8f88da2d39f2148b370948e",
	//	},
	//}, nil
}
