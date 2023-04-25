package socializing_pool

import (
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/stader-labs/stader-node/stader-lib/stader"
	types2 "github.com/stader-labs/stader-node/stader-lib/types"
	"math/big"
)

func EstimateClaimRewards(sp *stader.SocializingPoolContractManager, index []*big.Int, amountSd []*big.Int, amountEth []*big.Int, merkleProof [][][32]byte, opts *bind.TransactOpts) (stader.GasInfo, error) {
	return sp.SocializingPoolContract.GetTransactionGasInfo(opts, "claim", index, amountSd, amountEth, merkleProof)
}

func ClaimRewards(sp *stader.SocializingPoolContractManager, index []*big.Int, amountSd []*big.Int, amountEth []*big.Int, merkleProof [][][32]byte, opts *bind.TransactOpts) (*types.Transaction, error) {
	tx, err := sp.SocializingPool.Claim(opts, index, amountSd, amountEth, merkleProof)
	if err != nil {
		return nil, err
	}

	return tx, nil
}

func GetRewardDetails(sp *stader.SocializingPoolContractManager, opts *bind.CallOpts) (types2.RewardCycleDetails, error) {
	return sp.SocializingPool.GetRewardDetails(opts)
}

func GetRewardCycleDetails(sp *stader.SocializingPoolContractManager, cycle *big.Int, opts *bind.CallOpts) (types2.CurrentRewardCycleDetails, error) {
	return sp.SocializingPool.GetRewardCycleDetails(opts, cycle)
}

func HasClaimedRewards(sp *stader.SocializingPoolContractManager, address common.Address, index *big.Int, opts *bind.CallOpts) (bool, error) {
	return sp.SocializingPool.ClaimedRewards(opts, address, index)
}

func IsSocializingPoolPaused(sp *stader.SocializingPoolContractManager, opts *bind.CallOpts) (bool, error) {
	return sp.SocializingPool.Paused(opts)
}
