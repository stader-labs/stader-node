package node

import (
	"fmt"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/stader-labs/stader-node/stader-lib/stader"
	types2 "github.com/stader-labs/stader-node/stader-lib/types"
	"math/big"
)

func EstimateOnboardNodeOperator(pnr *stader.PermissionlessNodeRegistryContractManager, mevSocialize bool, operatorName string, operatorRewarderAddress common.Address, opts *bind.TransactOpts) (stader.GasInfo, error) {
	return pnr.PermissionlessNodeRegistryContract.GetTransactionGasInfo(opts, "onboardNodeOperator", mevSocialize, operatorName, operatorRewarderAddress)
}

func OnboardNodeOperator(pnr *stader.PermissionlessNodeRegistryContractManager, mevSocialize bool, operatorName string, operatorRewarderAddress common.Address, opts *bind.TransactOpts) (*types.Transaction, error) {
	tx, err := pnr.PermissionlessNodeRegistry.OnboardNodeOperator(opts, mevSocialize, operatorName, operatorRewarderAddress)
	if err != nil {
		return nil, fmt.Errorf("Could not onboard node operator: %w", err)
	}

	return tx, nil
}

func EstimateChangeSocializingPoolState(pnr *stader.PermissionlessNodeRegistryContractManager, socializeEl bool, opts *bind.TransactOpts) (stader.GasInfo, error) {
	return pnr.PermissionlessNodeRegistryContract.GetTransactionGasInfo(opts, "changeSocializingPoolState", socializeEl)
}

func ChangeSocializingPoolState(pnr *stader.PermissionlessNodeRegistryContractManager, socializeEl bool, opts *bind.TransactOpts) (*types.Transaction, error) {
	tx, err := pnr.PermissionlessNodeRegistry.ChangeSocializingPoolState(opts, socializeEl)
	if err != nil {
		return nil, err
	}

	return tx, nil
}

func EstimateUpdateOperatorDetails(pnr *stader.PermissionlessNodeRegistryContractManager, operatorName string, operatorRewarderAddress common.Address, opts *bind.TransactOpts) (stader.GasInfo, error) {
	return pnr.PermissionlessNodeRegistryContract.GetTransactionGasInfo(opts, "updateOperatorDetails", operatorName, operatorRewarderAddress)
}

func UpdateOperatorDetails(pnr *stader.PermissionlessNodeRegistryContractManager, operatorName string, operatorRewarderAddress common.Address, opts *bind.TransactOpts) (*types.Transaction, error) {
	tx, err := pnr.PermissionlessNodeRegistry.UpdateOperatorDetails(opts, operatorName, operatorRewarderAddress)
	if err != nil {
		return nil, err
	}

	return tx, nil
}

func EstimateWithdrawFromNodeElVault(client stader.ExecutionClient, nevAddress common.Address, opts *bind.TransactOpts) (stader.GasInfo, error) {
	nev, err := stader.NewNodeElRewardVaultFactory(client, nevAddress)
	if err != nil {
		return stader.GasInfo{}, err
	}
	return nev.NodeElRewardVaultContract.GetTransactionGasInfo(opts, "withdraw")
}

func WithdrawFromNodeElVault(client stader.ExecutionClient, nevAddress common.Address, opts *bind.TransactOpts) (*types.Transaction, error) {
	nev, err := stader.NewNodeElRewardVaultFactory(client, nevAddress)
	if err != nil {
		return nil, err
	}
	return nev.NodeElRewardVault.Withdraw(opts)
}

func GetOperatorId(pnr *stader.PermissionlessNodeRegistryContractManager, nodeAddress common.Address, opts *bind.CallOpts) (*big.Int, error) {
	operatorId, err := pnr.PermissionlessNodeRegistry.OperatorIDByAddress(opts, nodeAddress)
	if err != nil {
		return nil, err
	}

	return operatorId, nil
}

func GetOperatorInfo(pnr *stader.PermissionlessNodeRegistryContractManager, operatorId *big.Int, opts *bind.CallOpts) (types2.OperatorInfo, error) {
	operatorInfo, err := pnr.PermissionlessNodeRegistry.OperatorStructById(opts, operatorId)
	if err != nil {
		return types2.OperatorInfo{}, err
	}

	return operatorInfo, nil
}

func GetNodeElRewardAddress(vf *stader.VaultFactoryContractManager, poolId uint8, operatorId *big.Int, opts *bind.CallOpts) (common.Address, error) {
	return vf.VaultFactory.ComputeNodeELRewardVaultAddress(opts, poolId, operatorId)
}

func GetSocializingPoolContract(pp *stader.PermissionlessPoolContractManager, opts *bind.CallOpts) (common.Address, error) {
	return pp.PermissionlessPool.GetSocializingPoolAddress(opts)
}

func GetSocializingPoolStateChangeBlock(pnr *stader.PermissionlessNodeRegistryContractManager, operatorId *big.Int, opts *bind.CallOpts) (*big.Int, error) {
	return pnr.PermissionlessNodeRegistry.GetSocializingPoolStateChangeBlock(opts, operatorId)
}

func GetNextOperatorId(pnr *stader.PermissionlessNodeRegistryContractManager, opts *bind.CallOpts) (*big.Int, error) {
	return pnr.PermissionlessNodeRegistry.NextOperatorId(opts)
}
