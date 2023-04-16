package node

import (
	"github.com/stader-labs/stader-node/shared/services"
	"github.com/stader-labs/stader-node/shared/types/api"
	"github.com/stader-labs/stader-node/shared/utils/eth2"
	"github.com/stader-labs/stader-node/shared/utils/validator"
	"github.com/stader-labs/stader-node/stader-lib/node"
	"github.com/stader-labs/stader-node/stader-lib/types"
	"github.com/urfave/cli"
	eth2types "github.com/wealdtech/go-eth2-types/v2"
)

func canExitValidator(c *cli.Context, validatorPubKey types.ValidatorPubkey) (*api.CanExitValidatorResponse, error) {

	// Get services
	if err := services.RequireNodeWallet(c); err != nil {
		return nil, err
	}
	pnr, err := services.GetPermissionlessNodeRegistry(c)
	if err != nil {
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
	nodeAccount, err := w.GetNodeAccount()
	if err != nil {
		return nil, err
	}

	// Response
	response := api.CanExitValidatorResponse{}

	operatorId, err := node.GetOperatorId(pnr, nodeAccount.Address, nil)
	if err != nil {
		return nil, err
	}
	if operatorId.Int64() != 0 {
		response.OperatorNotRegistered = true
		return &response, nil
	}

	operatorInfo, err := node.GetOperatorInfo(pnr, operatorId, nil)
	if err != nil {
		return nil, err
	}
	if !operatorInfo.Active {
		response.OperatorNotActive = true
		return &response, nil
	}

	res, err := bc.GetValidatorStatus(validatorPubKey, nil)
	if err != nil || !res.Exists {
		response.ValidatorNotRegistered = true
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
	signatureDomain, err := bc.GetDomainData(eth2types.DomainVoluntaryExit[:], head.Epoch, false)
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
