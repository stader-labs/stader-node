package node

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/stader-labs/stader-node/shared/types/eth2"
	"github.com/stader-labs/stader-node/stader-lib/types"
	eth2types "github.com/wealdtech/go-eth2-types/v2"
	"math/big"

	"github.com/urfave/cli"

	"github.com/stader-labs/stader-node/shared/services"
	"github.com/stader-labs/stader-node/shared/types/api"
)

// Get a voluntary exit message signature for a given validator key and index
func GetSignedExitMessage(validatorKey *eth2types.BLSPrivateKey, validatorIndex uint64, epoch uint64, signatureDomain []byte) (types.ValidatorSignature, [32]byte, error) {

	// Build voluntary exit message
	exitMessage := eth2.VoluntaryExit{
		Epoch:          epoch,
		ValidatorIndex: validatorIndex,
	}

	// Get object root
	or, err := exitMessage.HashTreeRoot()
	if err != nil {
		return types.ValidatorSignature{}, [32]byte{}, err
	}

	// Get signing root
	sr := eth2.SigningRoot{
		ObjectRoot: or[:],
		Domain:     signatureDomain,
	}

	srHash, err := sr.HashTreeRoot()
	if err != nil {
		return types.ValidatorSignature{}, [32]byte{}, nil
	}

	// Sign message
	signature := validatorKey.Sign(srHash[:]).Marshal()

	// Return
	return types.BytesToValidatorSignature(signature), srHash, nil

}

func DebugExit(c *cli.Context, valIndex *big.Int) (*api.DebugExitResponse, error) {
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

	signatureDomain, err := bc.GetDomainData(eth2types.DomainVoluntaryExit[:], currentHead.Epoch)
	if err != nil {
		return nil, err
	}

	validatorIndex, err := bc.GetValidatorIndex(types.BytesToValidatorPubkey(validatorPubKey.Marshal()))
	if err != nil {
		return nil, err
	}

	exitMsg, srHash, err := GetSignedExitMessage(validatorPrivateKey, validatorIndex, currentHead.Epoch, signatureDomain)
	if err != nil {
		return nil, err
	}

	// Response
	response := api.DebugExitResponse{}
	response.ExitEpoch = currentHead.Epoch
	response.ValidatorIndex = validatorIndex
	response.ValidatorPubKey = types.BytesToValidatorPubkey(validatorPubKey.Marshal())
	response.SignatureDomain = common.BytesToHash(signatureDomain)
	response.SrHash = common.BytesToHash(srHash[:])
	response.Status = "success"
	response.SignedMsg = exitMsg

	// Return response
	return &response, nil

}
