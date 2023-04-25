package stake_pool_manager

import (
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/stader-labs/stader-node/stader-lib/stader"
	"math/big"
)

func GetTotalAssets(spm *stader.StakePoolManagerContractManager, opts *bind.CallOpts) (*big.Int, error) {
	return spm.StakePoolManager.TotalAssets(opts)
}
