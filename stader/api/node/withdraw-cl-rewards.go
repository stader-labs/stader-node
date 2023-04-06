package node

import (
	"github.com/stader-labs/stader-node/shared/services"
	"github.com/stader-labs/stader-node/shared/types/api"
	"github.com/stader-labs/stader-node/stader-lib/node"
	stader_config "github.com/stader-labs/stader-node/stader-lib/stader-config"
	"github.com/stader-labs/stader-node/stader-lib/tokens"
	"github.com/stader-labs/stader-node/stader-lib/types"
	"github.com/urfave/cli"
)

func CanWithdrawClRewards(c *cli.Context, validatorPubKey types.ValidatorPubkey) (*api.CanWithdrawClRewardsResponse, error) {
	// Get services
	if err := services.RequireNodeRegistered(c); err != nil {
		return nil, err
	}
	pnr, err := services.GetPermissionlessNodeRegistry(c)
	if err != nil {
		return nil, err
	}
	sdcfg, err := services.GetStaderConfigContract(c)
	if err != nil {
		return nil, err
	}

	// Response
	response := api.CanWithdrawClRewardsResponse{}

	validatorId, err := node.GetValidatorIdByPubKey(pnr, validatorPubKey.Bytes(), nil)
	if err != nil {
		return nil, err
	}
	if validatorId.Int64() == 0 {
		response.ValidatorNotFound = true
		return &response, nil
	}

	validatorContractInfo, err := node.GetValidatorInfo(pnr, validatorId, nil)
	if err != nil {
		return nil, err
	}
	if validatorContractInfo.Status > 8 {
		response.ValidatorWithdrawn = true
		return &response, nil
	}

	withdrawVaultBalance, err := tokens.GetEthBalance(pnr.Client, validatorContractInfo.WithdrawVaultAddress, nil)
	if err != nil {
		return nil, err
	}
	withdrawVaultRewardShares, err := node.CalculateValidatorWithdrawVaultRewardShare(pnr.Client, validatorContractInfo.WithdrawVaultAddress, withdrawVaultBalance, nil)
	if err != nil {
		return nil, err
	}
	if withdrawVaultRewardShares.OperatorShare.Int64() == 0 {
		response.NoClRewards = true
		return &response, nil
	}

	rewardsThreshold, err := stader_config.GetRewardsThreshold(sdcfg, nil)
	if err != nil {
		return nil, err
	}

	if withdrawVaultRewardShares.OperatorShare.Cmp(rewardsThreshold) > 0 {
		response.TooManyClRewards = true
		return &response, nil
	}

	gasInfo, err := node.EstimateDistributeRewards(pnr.Client, validatorContractInfo.WithdrawVaultAddress, nil)
	if err != nil {
		return nil, err
	}
	response.GasInfo = gasInfo

	return &response, nil
}

func WithdrawClRewards(c *cli.Context, validatorPubKey types.ValidatorPubkey) (*api.WithdrawClRewardsResponse, error) {
	w, err := services.GetWallet(c)
	if err != nil {
		return nil, err
	}
	pnr, err := services.GetPermissionlessNodeRegistry(c)
	if err != nil {
		return nil, err
	}

	nodeAccount, err := w.GetNodeAccount()
	if err != nil {
		return nil, err
	}

	response := api.WithdrawClRewardsResponse{}

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

	validatorContractInfo, err := node.GetValidatorInfo(pnr, validatorId, nil)
	if err != nil {
		return nil, err
	}
	withdrawVaultBalance, err := tokens.GetEthBalance(pnr.Client, validatorContractInfo.WithdrawVaultAddress, nil)
	if err != nil {
		return nil, err
	}
	withdrawVaultRewardShares, err := node.CalculateValidatorWithdrawVaultRewardShare(pnr.Client, validatorContractInfo.WithdrawVaultAddress, withdrawVaultBalance, nil)
	if err != nil {
		return nil, err
	}

	response.ClRewardsAmount = withdrawVaultRewardShares.OperatorShare
	response.OperatorRewardAddress = operatorInfo.OperatorRewardAddress

	tx, err := node.DistributeRewards(pnr.Client, validatorContractInfo.WithdrawVaultAddress, nil)
	if err != nil {
		return nil, err
	}

	response.TxHash = tx.Hash()

	return &response, nil
}
