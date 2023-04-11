package socializing_pool

import (
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/stader-labs/stader-node/stader-lib/stader"
	"math/big"
)

func ClaimRewards(sp *stader.SocializingPoolContractManager, index []*big.Int, amountSd []*big.Int, amountEth []*big.Int, merkleProof [][][32]byte, opts *bind.TransactOpts) (*types.Transaction, error) {
	tx, err := sp.SocializingPool.Claim(opts, index, amountSd, amountEth, merkleProof)
	if err != nil {
		return nil, err
	}

	return tx, nil
}

func GetRewardDetails(sp *stader.SocializingPoolContractManager, opts *bind.CallOpts) (struct {
	CurrentIndex      *big.Int
	CurrentStartBlock *big.Int
	CurrentEndBlock   *big.Int
	NextIndex         *big.Int
	NextStartBlock    *big.Int
	NextEndBlock      *big.Int
}, error) {
	return sp.SocializingPool.GetRewardDetails(nil)
}

func HasClaimedRewards(sp *stader.SocializingPoolContractManager, address common.Address, index *big.Int, opts *bind.CallOpts) (bool, error) {
	return sp.SocializingPool.ClaimedRewards(opts, address, index)
}
