package node

import (
	"github.com/stader-labs/stader-node/shared/services"
	"github.com/stader-labs/stader-node/shared/types/api"
	"github.com/stader-labs/stader-node/stader-lib/node"
	"github.com/stader-labs/stader-node/stader-lib/types"
	"github.com/urfave/cli"
	"math/big"
)

func CanSettleExitFunds(c *cli.Context, validatorPubKey types.ValidatorPubkey) (*api.CanSettleExitFunds, error) {
	pnr, err := services.GetPermissionlessNodeRegistry(c)
	if err != nil {
		return nil, err
	}
	w, err := services.GetWallet(c)
	if err != nil {
		return nil, err
	}
	nodeAccount, err := w.GetNodeAccount()
	if err != nil {
		return nil, err
	}
	opts, err := w.GetNodeAccountTransactor()
	if err != nil {
		return nil, err
	}

	response := api.CanSettleExitFunds{}

	operatorId, err := node.GetOperatorId(pnr, nodeAccount.Address, nil)
	if err != nil {
		return nil, err
	}
	if operatorId.Cmp(big.NewInt(0)) == 0 {
		response.OperatorNotRegistered = true
		return &response, nil
	}

	operatorInfo, err := node.GetOperatorInfo(pnr, operatorId, nil)
	if err != nil {
		return nil, err
	}
	if !operatorInfo.Active {
		response.OperatorNotActive = true
		return &response, nil
	}

	// make sure validator state is in withdrawn
	validatorId, err := node.GetValidatorIdByPubKey(pnr, validatorPubKey.Bytes(), nil)
	if err != nil {
		return nil, err
	}
	validatorInfo, err := node.GetValidatorInfo(pnr, validatorId, nil)
	if err != nil {
		return nil, err
	}
	if validatorInfo.Status < 8 {
		response.ValidatorNotWithdrawn = true
		return &response, nil
	}

	vaultSettleStatus, err := node.GetValidatorWithdrawVaultSettleStatus(pnr.Client, validatorInfo.WithdrawVaultAddress, nil)
	if err != nil {
		return nil, err
	}
	if vaultSettleStatus {
		response.VaultAlreadySettled = true
		return &response, nil
	}

	withdrawVaultWithdrawShares, err := node.CalculateValidatorWithdrawVaultWithdrawShare(pnr.Client, validatorInfo.WithdrawVaultAddress, nil)
	if err != nil {
		return nil, err
	}
	validatorWithdrawVaultWithdrawShares := withdrawVaultWithdrawShares.OperatorShare
	// make sure validator has eth to withdraw
	if validatorWithdrawVaultWithdrawShares.Int64() == 0 {
		response.NotEthToWithdraw = true
		return &response, nil
	}

	gasInfo, err := node.EstimateSettleFunds(pnr.Client, validatorInfo.WithdrawVaultAddress, opts)
	if err != nil {
		return nil, err
	}
	response.GasInfo = gasInfo

	return &response, nil
}

func SettleExitFunds(c *cli.Context, validatorPubKey types.ValidatorPubkey) (*api.SettleExitFunds, error) {
	pnr, err := services.GetPermissionlessNodeRegistry(c)
	if err != nil {
		return nil, err
	}
	w, err := services.GetWallet(c)
	if err != nil {
		return nil, err
	}
	nodeAccount, err := w.GetNodeAccount()
	if err != nil {
		return nil, err
	}
	opts, err := w.GetNodeAccountTransactor()
	if err != nil {
		return nil, err
	}

	response := api.SettleExitFunds{}

	operatorId, err := node.GetOperatorId(pnr, nodeAccount.Address, nil)
	if err != nil {
		return nil, err
	}
	operatorInfo, err := node.GetOperatorInfo(pnr, operatorId, nil)
	if err != nil {
		return nil, err
	}

	validatorId, err := node.GetValidatorIdByPubKey(pnr, validatorPubKey.Bytes(), nil)
	if err != nil {
		return nil, err
	}
	validatorInfo, err := node.GetValidatorInfo(pnr, validatorId, nil)
	if err != nil {
		return nil, err
	}
	withdrawVaultWithdrawShares, err := node.CalculateValidatorWithdrawVaultWithdrawShare(pnr.Client, validatorInfo.WithdrawVaultAddress, nil)
	if err != nil {
		return nil, err
	}
	validatorWithdrawVaultWithdrawShares := withdrawVaultWithdrawShares.OperatorShare

	response.OperatorRewardAddress = operatorInfo.OperatorRewardAddress
	response.ExitAmount = validatorWithdrawVaultWithdrawShares

	tx, err := node.SettleFunds(pnr.Client, validatorInfo.WithdrawVaultAddress, opts)
	if err != nil {
		return nil, err
	}

	response.TxHash = tx.Hash()

	return &response, nil
}
