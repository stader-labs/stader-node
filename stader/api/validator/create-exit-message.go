package validator

import (
	"fmt"

	"github.com/rocket-pool/rocketpool-go/types"
	"github.com/urfave/cli"
	eth2types "github.com/wealdtech/go-eth2-types/v2"

	"github.com/stader-labs/stader-node/shared/services"
	"github.com/stader-labs/stader-node/shared/utils/validator"
)

type VoluntaryExitMessage struct {
	Epoch          uint64 `json:"epoch"`
	ValidatorIndex uint64 `json:"validator_index"`
}

type VoluntaryExitRequest struct {
	Message   VoluntaryExitMessage `json:"message"`
	Signature []byte               `json:"signature"`
}

func CreateExitMessage(c *cli.Context, validatorPubkey types.ValidatorPubkey) (*VoluntaryExitRequest, error) {

	// Get services
	// if err := services.RequireNodeRegistered(c); err != nil {
	// 	return nil, err
	// }
	// if err := services.RequireBeaconClientSynced(c); err != nil {
	// 	return nil, err
	// }
	w, err := services.GetWallet(c)
	if err != nil {
		return nil, err
	}

	bc, err := services.GetBeaconClient(c)
	if err != nil {
		return nil, err
	}

	// Get validator private key
	validatorKey, err := w.GetValidatorKeyByPubkey(validatorPubkey)
	if err != nil {
		return nil, err
	}

	// Get beacon head
	head, err := bc.GetBeaconHead()
	if err != nil {
		return nil, err
	}

	// Get voluntary exit signature domain
	signatureDomain, err := bc.GetDomainData(eth2types.DomainVoluntaryExit[:], head.Epoch)
	if err != nil {
		return nil, err
	}

	// Get validator index
	validatorIndex, err := bc.GetValidatorIndex(validatorPubkey)
	if err != nil {
		return nil, err
	}

	// Get signed voluntary exit message
	signature, err := validator.GetSignedExitMessage(validatorKey, validatorIndex, head.Epoch, signatureDomain)
	if err != nil {
		return nil, err
	}

	fmt.Println("validatorIndex: ", validatorIndex)
	fmt.Println("head.Epoch: ", head.Epoch)
	fmt.Println("signature: ", signature)

	response := VoluntaryExitRequest{
		Message: VoluntaryExitMessage{
			Epoch:          uint64(head.Epoch),
			ValidatorIndex: uint64(validatorIndex),
		},
		Signature: signature.Bytes(),
	}

	// Return response
	return &response, nil

}
