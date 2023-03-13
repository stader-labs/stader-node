package node

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/stader-labs/stader-node/stader-lib/types"
	eth2types "github.com/wealdtech/go-eth2-types/v2"
	"math/big"

	"github.com/urfave/cli"

	"github.com/stader-labs/stader-node/shared/services"
	"github.com/stader-labs/stader-node/shared/types/api"
	"github.com/stader-labs/stader-node/shared/utils/validator"
)

func DebugExit(c *cli.Context, valIndex *big.Int, epochDelta *big.Int) (*api.DebugExitResponse, error) {
	w, err := services.GetWallet(c)
	if err != nil {
		return nil, err
	}
	bc, err := services.GetBeaconClient(c)
	if err != nil {
		return nil, err
	}

	validatorPrivateKey, err := w.GetValidatorKeyAt(uint(valIndex.Uint64()))
	if err != nil {
		return nil, err
	}

	validatorPubKey := validatorPrivateKey.PublicKey()

	currentHead, err := bc.GetBeaconHead()
	if err != nil {
		return nil, err
	}

	exitEpoch := currentHead.Epoch + epochDelta.Uint64()

	signatureDomain, err := bc.GetDomainData(eth2types.DomainVoluntaryExit[:], exitEpoch)
	if err != nil {
		return nil, err
	}

	validatorIndex, err := bc.GetValidatorIndex(types.BytesToValidatorPubkey(validatorPubKey.Marshal()))
	if err != nil {
		return nil, err
	}

	exitMsg, srHash, err := validator.GetSignedExitMessage(validatorPrivateKey, validatorIndex, exitEpoch, signatureDomain)
	if err != nil {
		return nil, err
	}

	// Response
	response := api.DebugExitResponse{}
	response.ExitEpoch = exitEpoch
	response.CurrentEpoch = currentHead.Epoch
	response.ValidatorIndex = validatorIndex
	response.ValidatorPubKey = types.BytesToValidatorPubkey(validatorPubKey.Marshal())
	response.SignatureDomain = common.BytesToHash(signatureDomain)
	response.SrHash = common.BytesToHash(srHash[:])
	response.Status = "success"
	response.SignedMsg = exitMsg

	// Return response
	return &response, nil

}
