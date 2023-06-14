package node

import (
	"fmt"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/stader-labs/stader-node/stader-lib/contracts"
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

func EstimateClaimOperatorRewards(orc *stader.OperatorRewardsCollectorContractManager, opts *bind.TransactOpts) (stader.GasInfo, error) {
	return orc.OperatorRewardsCollectorContract.GetTransactionGasInfo(opts, "claim")
}

func ClaimOperatorRewards(orc *stader.OperatorRewardsCollectorContractManager, opts *bind.TransactOpts) (*types.Transaction, error) {
	tx, err := orc.OperatorRewardsCollector.Claim(opts)
	if err != nil {
		return nil, fmt.Errorf("Could not claim operator rewards: %w", err)
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

func GetNodeElRewardAddress(pnr *stader.PermissionlessNodeRegistryContractManager, poolId uint8, operatorId *big.Int, opts *bind.CallOpts) (common.Address, error) {
	//return vf.VaultFactory.ComputeNodeELRewardVaultAddress(opts, poolId, operatorId)
	return pnr.PermissionlessNodeRegistry.NodeELRewardVaultByOperatorId(opts, operatorId)
}

func GetSocializingPoolStateChangeBlock(pnr *stader.PermissionlessNodeRegistryContractManager, operatorId *big.Int, opts *bind.CallOpts) (*big.Int, error) {
	return pnr.PermissionlessNodeRegistry.GetSocializingPoolStateChangeBlock(opts, operatorId)
}

func GetNextOperatorId(pnr *stader.PermissionlessNodeRegistryContractManager, opts *bind.CallOpts) (*big.Int, error) {
	return pnr.PermissionlessNodeRegistry.NextOperatorId(opts)
}

func GetOperatorRewardsCollectorBalance(orc *stader.OperatorRewardsCollectorContractManager, operatorRewardAddress common.Address, opts *bind.CallOpts) (*big.Int, error) {
	return orc.OperatorRewardsCollector.Balances(opts, operatorRewardAddress)
}

func GetValidatorInfosByOperator(pnr *stader.PermissionlessNodeRegistryContractManager, operatorAddress common.Address, pageNumber *big.Int, pageSize *big.Int, opts *bind.CallOpts) ([]contracts.Validator, error) {
	return pnr.PermissionlessNodeRegistry.GetValidatorsByOperator(opts, operatorAddress, pageNumber, pageSize)
}

func GetAllValidatorsInfoByOperator(pnr *stader.PermissionlessNodeRegistryContractManager, operatorAddress common.Address, opts *bind.CallOpts) ([]contracts.Validator, error) {
	finalValidators := []contracts.Validator{}
	pageNumber := big.NewInt(1)
	pageSize := big.NewInt(100)

	for {
		validators, err := pnr.PermissionlessNodeRegistry.GetValidatorsByOperator(opts, operatorAddress, pageNumber, pageSize)
		if err != nil {
			return nil, err
		}
		if len(validators) == 0 {
			break
		}

		finalValidators = append(finalValidators, validators...)
		pageNumber.Add(pageNumber, big.NewInt(1))
	}

	return finalValidators, nil
}
