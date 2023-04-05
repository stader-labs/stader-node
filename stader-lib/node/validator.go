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

func EstimateSettleFunds(executionClient stader.ExecutionClient, validatorWithdrawVaultAddress common.Address, opts *bind.TransactOpts) (stader.GasInfo, error) {
	vwv, err := stader.NewValidatorWithdrawVaultFactory(executionClient, validatorWithdrawVaultAddress)
	if err != nil {
		return stader.GasInfo{}, err
	}

	return vwv.ValidatorWithdrawVaultContract.GetTransactionGasInfo(opts, "settleFunds")

}

func SettleFunds(executionClient stader.ExecutionClient, validatorWithdrawVaultAddress common.Address, opts *bind.TransactOpts) (*types.Transaction, error) {
	vwv, err := stader.NewValidatorWithdrawVaultFactory(executionClient, validatorWithdrawVaultAddress)
	if err != nil {
		return nil, err
	}

	return vwv.ValidatorWithdrawVault.SettleFunds(opts)
}

func GetTotalValidatorKeys(pnr *stader.PermissionlessNodeRegistryContractManager, operatorId *big.Int, opts *bind.CallOpts) (*big.Int, error) {
	return pnr.PermissionlessNodeRegistry.GetOperatorTotalKeys(opts, operatorId)
}

func GetValidatorIdByOperatorId(pnr *stader.PermissionlessNodeRegistryContractManager, operatorId *big.Int, validatorIndex *big.Int, opts *bind.CallOpts) (*big.Int, error) {
	return pnr.PermissionlessNodeRegistry.ValidatorIdsByOperatorId(opts, operatorId, validatorIndex)
}

func GetValidatorInfo(pnr *stader.PermissionlessNodeRegistryContractManager, validatorId *big.Int, opts *bind.CallOpts) (struct {
	Status               uint8
	Pubkey               []byte
	PreDepositSignature  []byte
	DepositSignature     []byte
	WithdrawVaultAddress common.Address
	OperatorId           *big.Int
	InitialBondEth       *big.Int
	DepositTime          *big.Int
	WithdrawnTime        *big.Int
}, error) {
	return pnr.PermissionlessNodeRegistry.ValidatorRegistry(opts, validatorId)
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

func CalculateValidatorWithdrawVaultRewardShare(executionClient stader.ExecutionClient, validatorWithdrawVaultAddress common.Address, totalRewards *big.Int, opts *bind.CallOpts) (struct {
	UserShare     *big.Int
	OperatorShare *big.Int
	ProtocolShare *big.Int
}, error) {
	vwv, err := stader.NewValidatorWithdrawVaultFactory(executionClient, validatorWithdrawVaultAddress)
	if err != nil {
		return struct {
			UserShare     *big.Int
			OperatorShare *big.Int
			ProtocolShare *big.Int
		}{}, err
	}

	return vwv.ValidatorWithdrawVault.CalculateRewardShare(opts, totalRewards)
}

func CalculateValidatorWithdrawVaultWithdrawShare(executionClient stader.ExecutionClient, validatorWithdrawVaultAddress common.Address, opts *bind.CallOpts) (struct {
	UserShare     *big.Int
	OperatorShare *big.Int
	ProtocolShare *big.Int
}, error) {
	vwv, err := stader.NewValidatorWithdrawVaultFactory(executionClient, validatorWithdrawVaultAddress)
	if err != nil {
		return struct {
			UserShare     *big.Int
			OperatorShare *big.Int
			ProtocolShare *big.Int
		}{}, err
	}

	return vwv.ValidatorWithdrawVault.CalculateValidatorWithdrawalShare(opts)
}
