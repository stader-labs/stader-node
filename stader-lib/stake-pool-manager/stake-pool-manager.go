package stake_pool_manager

import (
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/stader-labs/stader-node/stader-lib/stader"
)

func GetTotalAssets(spm *stader.StakePoolManagerContractManager, opts *bind.CallOpts) (*big.Int, error) {
	return spm.StakePoolManager.TotalAssets(opts)
}
