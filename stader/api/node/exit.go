package node

import (
	"github.com/stader-labs/stader-node/shared/services"
	"github.com/stader-labs/stader-node/shared/types/api"
	"github.com/stader-labs/stader-node/shared/utils/validator"
	"github.com/stader-labs/stader-node/stader-lib/types"
	"github.com/urfave/cli"
	eth2types "github.com/wealdtech/go-eth2-types/v2"
)

func canExitValidator(c *cli.Context, validatorPubKey types.ValidatorPubkey) (*api.CanExitValidatorResponse, error) {

	// Get services
	if err := services.RequireNodeRegistered(c); err != nil {
		return nil, err
	}
	bc, err := services.GetBeaconClient(c)
	if err != nil {
		return nil, err
	}

	// Response
	response := api.CanExitValidatorResponse{}

	_, err = bc.GetValidatorStatus(validatorPubKey, nil)
	if err != nil {
		response.ValidatorNotRegistered = true
		return &response, nil
	}

	response.CanExit = true
	return &response, nil
}

func exitValidator(c *cli.Context, validatorPubKey types.ValidatorPubkey) (*api.ExitValidatorResponse, error) {

	// Get services
	if err := services.RequireNodeRegistered(c); err != nil {
		return nil, err
	}
	if err := services.RequireBeaconClientSynced(c); err != nil {
		return nil, err
	}
	w, err := services.GetWallet(c)
	if err != nil {
		return nil, err
	}
	bc, err := services.GetBeaconClient(c)
	if err != nil {
		return nil, err
	}

	// Response
	response := api.ExitValidatorResponse{}

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
	validatorIndex, err := bc.GetValidatorIndex(validatorPubKey)
	if err != nil {
		return nil, err
	}

	validatorKey, err := w.GetValidatorKeyByPubkey(validatorPubKey)

	// Get signed voluntary exit message
	signature, _, err := validator.GetSignedExitMessage(validatorKey, validatorIndex, head.Epoch, signatureDomain)
	if err != nil {
		return nil, err
	}

	// Broadcast voluntary exit message
	if err := bc.ExitValidator(validatorIndex, head.Epoch, signature); err != nil {
		return nil, err
	}

	// Return response
	return &response, nil

}
