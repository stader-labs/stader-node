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

func EstimateDistributeRewards(executionClient stader.ExecutionClient, validatorWithdrawVaultAddress common.Address, opts *bind.TransactOpts) (stader.GasInfo, error) {
	vwv, err := stader.NewValidatorWithdrawVaultFactory(executionClient, validatorWithdrawVaultAddress)
	if err != nil {
		return stader.GasInfo{}, err
	}

	return vwv.ValidatorWithdrawVaultContract.GetTransactionGasInfo(opts, "distributeRewards")
}

func DistributeRewards(executionClient stader.ExecutionClient, validatorWithdrawVaultAddress common.Address, opts *bind.TransactOpts) (*types.Transaction, error) {
	vwv, err := stader.NewValidatorWithdrawVaultFactory(executionClient, validatorWithdrawVaultAddress)
	if err != nil {
		return nil, err
	}

	return vwv.ValidatorWithdrawVault.DistributeRewards(opts)
}

func GetTotalValidatorKeys(pnr *stader.PermissionlessNodeRegistryContractManager, operatorId *big.Int, opts *bind.CallOpts) (*big.Int, error) {
	return pnr.PermissionlessNodeRegistry.GetOperatorTotalKeys(opts, operatorId)
}

func GetTotalNonTerminalValidatorKeys(pnr *stader.PermissionlessNodeRegistryContractManager, operatorAddress common.Address, maxPaginationIndex *big.Int, opts *bind.CallOpts) (uint64, error) {
	return pnr.PermissionlessNodeRegistry.GetOperatorTotalNonTerminalKeys(opts, operatorAddress, big.NewInt(0), maxPaginationIndex)
}

func GetMaxValidatorKeysPerOperator(pnr *stader.PermissionlessNodeRegistryContractManager, opts *bind.CallOpts) (uint64, error) {
	return pnr.PermissionlessNodeRegistry.MaxNonTerminalKeyPerOperator(opts)
}

func GetValidatorIdByOperatorId(pnr *stader.PermissionlessNodeRegistryContractManager, operatorId *big.Int, validatorIndex *big.Int, opts *bind.CallOpts) (*big.Int, error) {
	return pnr.PermissionlessNodeRegistry.ValidatorIdsByOperatorId(opts, operatorId, validatorIndex)
}

// TODO - bchain - make these inline structs into types
func GetValidatorInfo(pnr *stader.PermissionlessNodeRegistryContractManager, validatorId *big.Int, opts *bind.CallOpts) (struct {
	Status               uint8
	Pubkey               []byte
	PreDepositSignature  []byte
	DepositSignature     []byte
	WithdrawVaultAddress common.Address
	OperatorId           *big.Int
	DepositBlock         *big.Int
	WithdrawnBlock       *big.Int
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

func GetValidatorWithdrawVaultSettleStatus(executionClient stader.ExecutionClient, validatorWithdrawVaultAddress common.Address, opts *bind.CallOpts) (bool, error) {
	vwv, err := stader.NewValidatorWithdrawVaultFactory(executionClient, validatorWithdrawVaultAddress)
	if err != nil {
		return false, err
	}

	return vwv.ValidatorWithdrawVault.VaultSettleStatus(opts)
}

func GetValidatorIdByPubKey(pnr *stader.PermissionlessNodeRegistryContractManager, validatorPubKey []byte, opts *bind.CallOpts) (*big.Int, error) {
	return pnr.PermissionlessNodeRegistry.ValidatorIdByPubkey(opts, validatorPubKey)
}

func GetNextValidatorId(pnr *stader.PermissionlessNodeRegistryContractManager, opts *bind.CallOpts) (*big.Int, error) {
	return pnr.PermissionlessNodeRegistry.NextValidatorId(opts)
}
