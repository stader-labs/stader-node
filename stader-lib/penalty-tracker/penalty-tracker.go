package penalty_tracker

import (
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/stader-labs/stader-node/stader-lib/stader"
	"github.com/stader-labs/stader-node/stader-lib/types"
	"math/big"
)

func GetCumulativeValidatorPenalty(pt *stader.PenaltyTrackerContractManager, validatorPubKey types.ValidatorPubkey, opts *bind.CallOpts) (*big.Int, error) {
	return pt.Penalty.TotalPenaltyAmount(opts, validatorPubKey.Bytes())
}
