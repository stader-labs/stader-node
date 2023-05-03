package stader_config

import (
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/stader-labs/stader-node/stader-lib/stader"
	"math/big"
)

func GetRewardsThreshold(sdConfig *stader.StaderConfigContractManager, opts *bind.CallOpts) (*big.Int, error) {
	return sdConfig.StaderConfig.GetRewardsThreshold(opts)
}

func GetSocializingPoolChangeThreshold(sdConfig *stader.StaderConfigContractManager, opts *bind.CallOpts) (*big.Int, error) {
	return sdConfig.StaderConfig.GetSocializingPoolOptInCoolingPeriod(opts)
}

func GetOperatorNameMaxLength(sdConfig *stader.StaderConfigContractManager, opts *bind.CallOpts) (*big.Int, error) {
	return sdConfig.StaderConfig.GetOperatorMaxNameLength(opts)
}

func GetSocializingPoolContractAddress(sdConfig *stader.StaderConfigContractManager, opts *bind.CallOpts) (common.Address, error) {
	return sdConfig.StaderConfig.GetSocializingPool(opts)
}
