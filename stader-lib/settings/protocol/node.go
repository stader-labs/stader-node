package protocol

import (
	"fmt"
	"math/big"
	"sync"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"

	protocoldao "github.com/stader-labs/stader-node/stader-lib/dao/protocol"
	"github.com/stader-labs/stader-node/stader-lib/stader"
	"github.com/stader-labs/stader-node/stader-lib/utils/eth"
)

// Config
const NodeSettingsContractName = "rocketDAOProtocolSettingsNode"

// Node registrations currently enabled
func GetNodeRegistrationEnabled(rp *stader.PermissionlessNodeRegistryContractManager, opts *bind.CallOpts) (bool, error) {
	nodeSettingsContract, err := getNodeSettingsContract(rp, opts)
	if err != nil {
		return false, err
	}
	value := new(bool)
	if err := nodeSettingsContract.Call(opts, value, "getRegistrationEnabled"); err != nil {
		return false, fmt.Errorf("Could not get node registrations enabled status: %w", err)
	}
	return *value, nil
}
func BootstrapNodeRegistrationEnabled(rp *stader.PermissionlessNodeRegistryContractManager, value bool, opts *bind.TransactOpts) (common.Hash, error) {
	return protocoldao.BootstrapBool(rp, NodeSettingsContractName, "node.registration.enabled", value, opts)
}

// Node deposits currently enabled
func GetNodeDepositEnabled(rp *stader.PermissionlessNodeRegistryContractManager, opts *bind.CallOpts) (bool, error) {
	nodeSettingsContract, err := getNodeSettingsContract(rp, opts)
	if err != nil {
		return false, err
	}
	value := new(bool)
	if err := nodeSettingsContract.Call(opts, value, "getDepositEnabled"); err != nil {
		return false, fmt.Errorf("Could not get node deposits enabled status: %w", err)
	}
	return *value, nil
}
func BootstrapNodeDepositEnabled(rp *stader.PermissionlessNodeRegistryContractManager, value bool, opts *bind.TransactOpts) (common.Hash, error) {
	return protocoldao.BootstrapBool(rp, NodeSettingsContractName, "node.deposit.enabled", value, opts)
}

// The minimum RPL stake per minipool as a fraction of assigned user ETH
func GetMinimumPerMinipoolStake(rp *stader.PermissionlessNodeRegistryContractManager, opts *bind.CallOpts) (float64, error) {
	nodeSettingsContract, err := getNodeSettingsContract(rp, opts)
	if err != nil {
		return 0, err
	}
	value := new(*big.Int)
	if err := nodeSettingsContract.Call(opts, value, "getMinimumPerMinipoolStake"); err != nil {
		return 0, fmt.Errorf("Could not get minimum RPL stake per minipool: %w", err)
	}
	return eth.WeiToEth(*value), nil
}

// The minimum RPL stake per minipool as a fraction of assigned user ETH
func GetMinimumPerMinipoolStakeRaw(rp *stader.PermissionlessNodeRegistryContractManager, opts *bind.CallOpts) (*big.Int, error) {
	nodeSettingsContract, err := getNodeSettingsContract(rp, opts)
	if err != nil {
		return nil, err
	}
	value := new(*big.Int)
	if err := nodeSettingsContract.Call(opts, value, "getMinimumPerMinipoolStake"); err != nil {
		return nil, fmt.Errorf("Could not get minimum RPL stake per minipool: %w", err)
	}
	return *value, nil
}
func BootstrapMinimumPerMinipoolStake(rp *stader.PermissionlessNodeRegistryContractManager, value float64, opts *bind.TransactOpts) (common.Hash, error) {
	return protocoldao.BootstrapUint(rp, NodeSettingsContractName, "node.per.minipool.stake.minimum", eth.EthToWei(value), opts)
}

// The maximum RPL stake per minipool as a fraction of assigned user ETH
func GetMaximumPerMinipoolStake(rp *stader.PermissionlessNodeRegistryContractManager, opts *bind.CallOpts) (float64, error) {
	nodeSettingsContract, err := getNodeSettingsContract(rp, opts)
	if err != nil {
		return 0, err
	}
	value := new(*big.Int)
	if err := nodeSettingsContract.Call(opts, value, "getMaximumPerMinipoolStake"); err != nil {
		return 0, fmt.Errorf("Could not get maximum RPL stake per minipool: %w", err)
	}
	return eth.WeiToEth(*value), nil
}

// The maximum RPL stake per minipool as a fraction of assigned user ETH
func GetMaximumPerMinipoolStakeRaw(rp *stader.PermissionlessNodeRegistryContractManager, opts *bind.CallOpts) (*big.Int, error) {
	nodeSettingsContract, err := getNodeSettingsContract(rp, opts)
	if err != nil {
		return nil, err
	}
	value := new(*big.Int)
	if err := nodeSettingsContract.Call(opts, value, "getMaximumPerMinipoolStake"); err != nil {
		return nil, fmt.Errorf("Could not get maximum RPL stake per minipool: %w", err)
	}
	return *value, nil
}
func BootstrapMaximumPerMinipoolStake(rp *stader.PermissionlessNodeRegistryContractManager, value float64, opts *bind.TransactOpts) (common.Hash, error) {
	return protocoldao.BootstrapUint(rp, NodeSettingsContractName, "node.per.minipool.stake.maximum", eth.EthToWei(value), opts)
}

// Get contracts
var nodeSettingsContractLock sync.Mutex

func getNodeSettingsContract(rp *stader.PermissionlessNodeRegistryContractManager, opts *bind.CallOpts) (*stader.Contract, error) {
	nodeSettingsContractLock.Lock()
	defer nodeSettingsContractLock.Unlock()
	return nil, nil
	//return rp.GetContract(NodeSettingsContractName, opts)
}
