package validator

import (
	"fmt"

	"github.com/stader-labs/stader-node/shared/services"
	"github.com/stader-labs/stader-node/shared/types/api"
	"github.com/stader-labs/stader-node/shared/types/config"
	"github.com/stader-labs/stader-node/shared/utils/eth2"
	"github.com/stader-labs/stader-node/shared/utils/validator"
	"github.com/stader-labs/stader-node/stader-lib/types"
	"github.com/urfave/cli"
	eth2types "github.com/wealdtech/go-eth2-types/v2"
)

func canExitValidator(c *cli.Context, validatorPubKey types.ValidatorPubkey) (*api.CanExitValidatorResponse, error) {

	// Get services
	if err := services.RequireNodeWallet(c); err != nil {
		return nil, err
	}
	if err := services.RequireNodeRegistered(c); err != nil {
		return nil, err
	}
	bc, err := services.GetBeaconClient(c)
	if err != nil {
		return nil, err
	}
	w, err := services.GetWallet(c)
	if err != nil {
		return nil, err
	}
	// Response
	response := api.CanExitValidatorResponse{}

	// check if the validator is key is available to sign the exit message
	_, err = w.GetValidatorKeyByPubkey(validatorPubKey)
	if err != nil {
		return nil, err
	}

	res, err := bc.GetValidatorStatus(validatorPubKey, nil)
	if err != nil {
		return nil, err
	}
	if !res.Exists {
		response.ValidatorNotRegistered = true
		return &response, nil
	}
	if !eth2.IsValidatorActive(res) {
		response.ValidatorNotActive = true
		return &response, nil
	}

	if eth2.IsValidatorExiting(res) {
		response.ValidatorExiting = true
		return &response, nil
	}

	beaconHead, err := bc.GetBeaconHead()
	if err != nil {
		return nil, err
	}
	currentEpoch := beaconHead.Epoch

	if res.ActivationEpoch+256 > currentEpoch {
		response.ValidatorTooYoung = true
		return &response, nil
	}

	return &response, nil
}

func exitValidator(c *cli.Context, validatorPubKey types.ValidatorPubkey) (*api.ExitValidatorResponse, error) {

	w, err := services.GetWallet(c)
	if err != nil {
		return nil, err
	}
	bc, err := services.GetBeaconClient(c)
	if err != nil {
		return nil, err
	}
	cfg, err := services.GetConfig(c)
	if err != nil {
		return nil, err
	}

	network, ok := cfg.StaderNode.Network.Value.(config.Network)
	if !ok {
		return nil, fmt.Errorf("invalid network configuration: %s", cfg.StaderNode.Network.Value)
	}

	// Response
	response := api.ExitValidatorResponse{}

	// Get beacon head
	head, err := bc.GetBeaconHead()
	if err != nil {
		return nil, err
	}

	// Get voluntary exit signature domain
	signatureDomain, err := bc.GetExitDomainData(eth2types.DomainVoluntaryExit[:], network)
	if err != nil {
		return nil, err
	}

	// Get validator index
	validatorIndex, err := bc.GetValidatorIndex(validatorPubKey)
	if err != nil {
		return nil, err
	}

	validatorKey, err := w.GetValidatorKeyByPubkey(validatorPubKey)
	if err != nil {
		return nil, err
	}

	// Get signed voluntary exit message
	signature, _, err := validator.GetSignedExitMessage(validatorKey, validatorIndex, head.Epoch, signatureDomain)
	if err != nil {
		return nil, err
	}

	// Broadcast voluntary exit message
	if err := bc.ExitValidator(validatorIndex, head.Epoch, signature); err != nil {
		return nil, err
	}

	response.BeaconChainUrl = cfg.StaderNode.GetBeaconChainUrl()

	// Return response
	return &response, nil

}
