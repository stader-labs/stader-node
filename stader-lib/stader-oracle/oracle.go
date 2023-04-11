package stader_oracle

import (
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/stader-labs/stader-node/stader-lib/stader"
	"math/big"
)

func GetMerkleRootForCycle(so *stader.StaderOracleContractManager, cycle *big.Int, opts *bind.CallOpts) ([32]byte, error) {
	merkleRoot, err := so.StaderOracle.SocializingRewardsMerkleRoot(opts, cycle)
	if err != nil {
		return [32]byte{}, err
	}

	return merkleRoot, nil
}
