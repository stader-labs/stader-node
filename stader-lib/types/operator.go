package types

import "github.com/ethereum/go-ethereum/common"

type OperatorInfo struct {
	Active                  bool
	OptedForSocializingPool bool
	OperatorName            string
	OperatorRewardAddress   common.Address
	OperatorAddress         common.Address
}
