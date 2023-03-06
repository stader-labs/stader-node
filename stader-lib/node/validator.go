package node

import (
	"fmt"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/stader-labs/stader-node/stader-lib/stader"
	"math/big"
)

func EstimateAddValidatorKeys(pnr *stader.PermissionlessNodeRegistryContractManager, pubKeys [][]byte, preDepositSignatures [][]byte, depositSignatures [][]byte, opts *bind.TransactOpts) (stader.GasInfo, error) {
	return pnr.PermissionlessNodeRegistryContract.GetTransactionGasInfo(opts, "addValidatorKeys", pubKeys, preDepositSignatures, depositSignatures)
}

func AddValidatorKeys(pnr *stader.PermissionlessNodeRegistryContractManager, pubKeys [][]byte, preDepositSignatures [][]byte, depositSignatures [][]byte, opts *bind.TransactOpts) (*types.Transaction, error) {
	tx, err := pnr.PermissionlessNodeRegistry.AddValidatorKeys(opts, pubKeys, preDepositSignatures, depositSignatures)
	if err != nil {
		return nil, fmt.Errorf("could not add validator keys: %w", err)
	}

	return tx, nil
}

func GetTotalValidatorKeys(pnr *stader.PermissionlessNodeRegistryContractManager, operatorId *big.Int, opts *bind.CallOpts) (*big.Int, error) {
	return pnr.PermissionlessNodeRegistry.GetOperatorTotalKeys(opts, operatorId)
}

func ComputeWithdrawVaultAddress(vfcm *stader.VaultFactoryContractManager, poolType uint8, operatorId *big.Int, validatorCount *big.Int, opts *bind.CallOpts) (common.Address, error) {
	return vfcm.VaultFactory.ComputeWithdrawVaultAddress(opts, poolType, operatorId, validatorCount)
}

func GetValidatorWithdrawalCredential(vfcm *stader.VaultFactoryContractManager, withdrawVaultAddress common.Address, opts *bind.CallOpts) (common.Hash, error) {
	withdrawalCredentials := new(common.Hash)
	if err := vfcm.VaultFactoryContract.Call(opts, withdrawalCredentials, "getValidatorWithdrawCredential", withdrawVaultAddress); err != nil {
		return common.Hash{}, fmt.Errorf("Could not get vault withdrawal credentials: %w", err)
	}

	return *withdrawalCredentials, nil
}
