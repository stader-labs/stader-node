package storage

import (
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/stader-labs/stader-node/stader-lib/stader"
)

// Get a node's withdrawal address
func GetNodeWithdrawalAddress(rp *stader.PermissionlessNodeRegistryContractManager, nodeAddress common.Address, opts *bind.CallOpts) (common.Address, error) {
	//withdrawalAddress := new(common.Address)
	//if err := rp.RocketStorageContract.Call(opts, withdrawalAddress, "getNodeWithdrawalAddress", nodeAddress); err != nil {
	//	return common.Address{}, fmt.Errorf("Could not get node %s withdrawal address: %w", nodeAddress.Hex(), err)
	//}
	//return *withdrawalAddress, nil
	return common.BytesToAddress([]byte("abc")), nil
}

// Get a node's pending withdrawal address
func GetNodePendingWithdrawalAddress(rp *stader.PermissionlessNodeRegistryContractManager, nodeAddress common.Address, opts *bind.CallOpts) (common.Address, error) {
	//withdrawalAddress := new(common.Address)
	//if err := rp.RocketStorageContract.Call(opts, withdrawalAddress, "getNodePendingWithdrawalAddress", nodeAddress); err != nil {
	//	return common.Address{}, fmt.Errorf("Could not get node %s pending withdrawal address: %w", nodeAddress.Hex(), err)
	//}
	return common.BytesToAddress([]byte("abc")), nil
}

// Estimate the gas of SetWithdrawalAddress
func EstimateSetWithdrawalAddressGas(rp *stader.PermissionlessNodeRegistryContractManager, nodeAddress common.Address, withdrawalAddress common.Address, confirm bool, opts *bind.TransactOpts) (stader.GasInfo, error) {
	//return rp.RocketStorageContract.GetTransactionGasInfo(opts, "setWithdrawalAddress", nodeAddress, withdrawalAddress, confirm)
	return stader.GasInfo{
		EstGasLimit:  0,
		SafeGasLimit: 0,
	}, nil
}

// Set a node's withdrawal address
func SetWithdrawalAddress(rp *stader.PermissionlessNodeRegistryContractManager, nodeAddress common.Address, withdrawalAddress common.Address, confirm bool, opts *bind.TransactOpts) (common.Hash, error) {
	//tx, err := rp.RocketStorageContract.Transact(opts, "setWithdrawalAddress", nodeAddress, withdrawalAddress, confirm)
	//if err != nil {
	//	return common.Hash{}, fmt.Errorf("Could not set node withdrawal address: %w", err)
	//}
	//return tx.Hash(), nil
	return common.BytesToHash([]byte("abc")), nil
}

// Estimate the gas of ConfirmWithdrawalAddress
func EstimateConfirmWithdrawalAddressGas(rp *stader.PermissionlessNodeRegistryContractManager, nodeAddress common.Address, opts *bind.TransactOpts) (stader.GasInfo, error) {
	//return rp.RocketStorageContract.GetTransactionGasInfo(opts, "confirmWithdrawalAddress", nodeAddress)
	return stader.GasInfo{
		EstGasLimit:  0,
		SafeGasLimit: 0,
	}, nil
}

// Set a node's withdrawal address
func ConfirmWithdrawalAddress(rp *stader.PermissionlessNodeRegistryContractManager, nodeAddress common.Address, opts *bind.TransactOpts) (common.Hash, error) {
	//tx, err := rp.RocketStorageContract.Transact(opts, "confirmWithdrawalAddress", nodeAddress)
	//if err != nil {
	//	return common.Hash{}, fmt.Errorf("Could not confirm node withdrawal address: %w", err)
	//}
	//return tx.Hash(), nil
	return common.BytesToHash([]byte("abc")), nil
}
