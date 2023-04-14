package pool_utils

import (
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/stader-labs/stader-node/stader-lib/stader"
	"math/big"
)

func CalculateNodeElRewardShare(pool_utils *stader.PoolUtilsContractManager, poolId uint8, totalRewards *big.Int, opts *bind.CallOpts) (struct {
	UserShare     *big.Int
	OperatorShare *big.Int
	ProtocolShare *big.Int
}, error) {
	return pool_utils.PoolUtils.CalculateRewardShare(opts, poolId, totalRewards)
}
