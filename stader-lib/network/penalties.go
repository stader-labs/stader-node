package network

import (
	"fmt"
	"math/big"
	"sync"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"

	"github.com/stader-labs/stader-node/stader-lib/stader"
)

// Estimate the gas of SubmitPenalty
func EstimateSubmitPenaltyGas(rp *stader.PermissionlessNodeRegistryContractManager, minipoolAddress common.Address, block *big.Int, opts *bind.TransactOpts) (stader.GasInfo, error) {
	rocketNetworkPenalties, err := getRocketNetworkPenalties(rp, nil)
	if err != nil {
		return stader.GasInfo{}, err
	}
	return rocketNetworkPenalties.GetTransactionGasInfo(opts, "submitPenalty", minipoolAddress, block)
}

// Submit penalty for given minipool
func SubmitPenalty(rp *stader.PermissionlessNodeRegistryContractManager, minipoolAddress common.Address, block *big.Int, opts *bind.TransactOpts) (common.Hash, error) {
	rocketNetworkPrices, err := getRocketNetworkPenalties(rp, nil)
	if err != nil {
		return common.Hash{}, err
	}
	tx, err := rocketNetworkPrices.Transact(opts, "submitPenalty", minipoolAddress, block)
	if err != nil {
		return common.Hash{}, fmt.Errorf("Could not submit penalty: %w", err)
	}
	return tx.Hash(), nil
}

// Get contracts
var rocketNetworkPenaltiesLock sync.Mutex

func getRocketNetworkPenalties(rp *stader.PermissionlessNodeRegistryContractManager, opts *bind.CallOpts) (*stader.Contract, error) {
	rocketNetworkPenaltiesLock.Lock()
	defer rocketNetworkPenaltiesLock.Unlock()
	return nil, nil
	//return rp.GetContract("rocketNetworkPenalties", opts)
}
