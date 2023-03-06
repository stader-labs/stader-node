package protocol

import (
	"fmt"
	"math/big"
	"sync"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"

	"github.com/stader-labs/stader-node/stader-lib/stader"
	"github.com/stader-labs/stader-node/stader-lib/utils/eth"
)

// Estimate the gas of BootstrapBool
func EstimateBootstrapBoolGas(rp *stader.PermissionlessNodeRegistryContractManager, contractName, settingPath string, value bool, opts *bind.TransactOpts) (stader.GasInfo, error) {
	rocketDAOProtocol, err := getRocketDAOProtocol(rp, nil)
	if err != nil {
		return stader.GasInfo{}, err
	}
	return rocketDAOProtocol.GetTransactionGasInfo(opts, "bootstrapSettingBool", contractName, settingPath, value)
}

// Bootstrap a bool setting
func BootstrapBool(rp *stader.PermissionlessNodeRegistryContractManager, contractName, settingPath string, value bool, opts *bind.TransactOpts) (common.Hash, error) {
	rocketDAOProtocol, err := getRocketDAOProtocol(rp, nil)
	if err != nil {
		return common.Hash{}, err
	}
	tx, err := rocketDAOProtocol.Transact(opts, "bootstrapSettingBool", contractName, settingPath, value)
	if err != nil {
		return common.Hash{}, fmt.Errorf("Could not bootstrap protocol setting %s.%s: %w", contractName, settingPath, err)
	}
	return tx.Hash(), nil
}

// Estimate the gas of BootstrapUint
func EstimateBootstrapUintGas(rp *stader.PermissionlessNodeRegistryContractManager, contractName, settingPath string, value *big.Int, opts *bind.TransactOpts) (stader.GasInfo, error) {
	rocketDAOProtocol, err := getRocketDAOProtocol(rp, nil)
	if err != nil {
		return stader.GasInfo{}, err
	}
	return rocketDAOProtocol.GetTransactionGasInfo(opts, "bootstrapSettingUint", contractName, settingPath, value)
}

// Bootstrap a uint256 setting
func BootstrapUint(rp *stader.PermissionlessNodeRegistryContractManager, contractName, settingPath string, value *big.Int, opts *bind.TransactOpts) (common.Hash, error) {
	rocketDAOProtocol, err := getRocketDAOProtocol(rp, nil)
	if err != nil {
		return common.Hash{}, err
	}
	tx, err := rocketDAOProtocol.Transact(opts, "bootstrapSettingUint", contractName, settingPath, value)
	if err != nil {
		return common.Hash{}, fmt.Errorf("Could not bootstrap protocol setting %s.%s: %w", contractName, settingPath, err)
	}
	return tx.Hash(), nil
}

// Estimate the gas of BootstrapAddress
func EstimateBootstrapAddressGas(rp *stader.PermissionlessNodeRegistryContractManager, contractName, settingPath string, value common.Address, opts *bind.TransactOpts) (stader.GasInfo, error) {
	rocketDAOProtocol, err := getRocketDAOProtocol(rp, nil)
	if err != nil {
		return stader.GasInfo{}, err
	}
	return rocketDAOProtocol.GetTransactionGasInfo(opts, "bootstrapSettingAddress", contractName, settingPath, value)
}

// Bootstrap an address setting
func BootstrapAddress(rp *stader.PermissionlessNodeRegistryContractManager, contractName, settingPath string, value common.Address, opts *bind.TransactOpts) (common.Hash, error) {
	rocketDAOProtocol, err := getRocketDAOProtocol(rp, nil)
	if err != nil {
		return common.Hash{}, err
	}
	tx, err := rocketDAOProtocol.Transact(opts, "bootstrapSettingAddress", contractName, settingPath, value)
	if err != nil {
		return common.Hash{}, fmt.Errorf("Could not bootstrap protocol setting %s.%s: %w", contractName, settingPath, err)
	}
	return tx.Hash(), nil
}

// Estimate the gas of BootstrapClaimer
func EstimateBootstrapClaimerGas(rp *stader.PermissionlessNodeRegistryContractManager, contractName string, amount float64, opts *bind.TransactOpts) (stader.GasInfo, error) {
	rocketDAOProtocol, err := getRocketDAOProtocol(rp, nil)
	if err != nil {
		return stader.GasInfo{}, err
	}
	return rocketDAOProtocol.GetTransactionGasInfo(opts, "bootstrapSettingClaimer", contractName, eth.EthToWei(amount))
}

// Bootstrap a rewards claimer
func BootstrapClaimer(rp *stader.PermissionlessNodeRegistryContractManager, contractName string, amount float64, opts *bind.TransactOpts) (common.Hash, error) {
	rocketDAOProtocol, err := getRocketDAOProtocol(rp, nil)
	if err != nil {
		return common.Hash{}, err
	}
	tx, err := rocketDAOProtocol.Transact(opts, "bootstrapSettingClaimer", contractName, eth.EthToWei(amount))
	if err != nil {
		return common.Hash{}, fmt.Errorf("Could not bootstrap protocol rewards claimer %s: %w", contractName, err)
	}
	return tx.Hash(), nil
}

// Get contracts
var rocketDAOProtocolLock sync.Mutex

func getRocketDAOProtocol(rp *stader.PermissionlessNodeRegistryContractManager, opts *bind.CallOpts) (*stader.Contract, error) {
	rocketDAOProtocolLock.Lock()
	defer rocketDAOProtocolLock.Unlock()
	return nil, nil
	//return rp.GetContract("rocketDAOProtocol", opts)
}
