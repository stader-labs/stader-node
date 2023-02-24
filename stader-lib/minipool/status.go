package minipool

import (
	"fmt"
	"sync"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"

	"github.com/stader-labs/stader-node/stader-lib/stader"
)

// Estimate the gas of SubmitMinipoolWithdrawable
func EstimateSubmitMinipoolWithdrawableGas(rp *stader.PermissionlessNodeRegistryContractManager, minipoolAddress common.Address, opts *bind.TransactOpts) (stader.GasInfo, error) {
	rocketMinipoolStatus, err := getRocketMinipoolStatus(rp, nil)
	if err != nil {
		return stader.GasInfo{}, err
	}
	return rocketMinipoolStatus.GetTransactionGasInfo(opts, "submitMinipoolWithdrawable", minipoolAddress)
}

// Submit a minipool withdrawable event
func SubmitMinipoolWithdrawable(rp *stader.PermissionlessNodeRegistryContractManager, minipoolAddress common.Address, opts *bind.TransactOpts) (common.Hash, error) {
	rocketMinipoolStatus, err := getRocketMinipoolStatus(rp, nil)
	if err != nil {
		return common.Hash{}, err
	}
	tx, err := rocketMinipoolStatus.Transact(opts, "submitMinipoolWithdrawable", minipoolAddress)
	if err != nil {
		return common.Hash{}, fmt.Errorf("Could not submit minipool withdrawable event: %w", err)
	}
	return tx.Hash(), nil
}

// Get contracts
var rocketMinipoolStatusLock sync.Mutex

func getRocketMinipoolStatus(rp *stader.PermissionlessNodeRegistryContractManager, opts *bind.CallOpts) (*stader.Contract, error) {
	rocketMinipoolStatusLock.Lock()
	defer rocketMinipoolStatusLock.Unlock()
	return nil, nil
	//return rp.GetContract("rocketMinipoolStatus", opts)
}
