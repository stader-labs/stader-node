package pool_utils

import (
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/stader-labs/stader-node/stader-lib/stader"
	"github.com/stader-labs/stader-node/stader-lib/types"
	"math/big"
)

func CalculateRewardShare(pool_utils *stader.PoolUtilsContractManager, poolId uint8, totalRewards *big.Int, opts *bind.CallOpts) (types.RewardShare, error) {
	return pool_utils.PoolUtils.CalculateRewardShare(opts, poolId, totalRewards)
}

func IsExistingOperator(pool_utils *stader.PoolUtilsContractManager, operatorAddress common.Address, opts *bind.CallOpts) (bool, error) {
	return pool_utils.PoolUtils.IsExistingOperator(opts, operatorAddress)
}
