package stader_config

import (
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/stader-labs/stader-node/stader-lib/stader"
	"math/big"
)

func GetRewardsThreshold(sdConfig *stader.StaderConfigContractManager, opts *bind.CallOpts) (*big.Int, error) {
	return sdConfig.StaderConfig.GetRewardsThreshold(opts)
}

func GetSocializingPoolChangeThreshold(sdConfig *stader.StaderConfigContractManager, opts *bind.CallOpts) (*big.Int, error) {
	return sdConfig.StaderConfig.GetSocializingPoolOptInCoolingPeriod(opts)
}
