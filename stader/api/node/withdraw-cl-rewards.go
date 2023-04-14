package node

import (
	"github.com/stader-labs/stader-node/shared/services"
	"github.com/stader-labs/stader-node/shared/types/api"
	"github.com/stader-labs/stader-node/stader-lib/node"
	pool_utils "github.com/stader-labs/stader-node/stader-lib/pool-utils"
	stader_config "github.com/stader-labs/stader-node/stader-lib/stader-config"
	"github.com/stader-labs/stader-node/stader-lib/tokens"
	"github.com/stader-labs/stader-node/stader-lib/types"
	"github.com/urfave/cli"
)

func CanWithdrawClRewards(c *cli.Context, validatorPubKey types.ValidatorPubkey) (*api.CanWithdrawClRewardsResponse, error) {
	// Get services
	pnr, err := services.GetPermissionlessNodeRegistry(c)
	if err != nil {
		return nil, err
	}
	sdcfg, err := services.GetStaderConfigContract(c)
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

	// Response
	response := api.CanWithdrawClRewardsResponse{}

	operatorId, err := node.GetOperatorId(pnr, nodeAccount.Address, nil)
	if err != nil {
		return nil, err
	}
	if operatorId.Int64() == 0 {
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

	// TODO - bchain - check if withdraw vault is settled

	withdrawVaultRewardShares, err := node.CalculateValidatorWithdrawVaultWithdrawShare(pnr.Client, validatorContractInfo.WithdrawVaultAddress, nil)
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

	opts, err := w.GetNodeAccountTransactor()
	if err != nil {
		return nil, err
	}

	gasInfo, err := node.EstimateDistributeRewards(pnr.Client, validatorContractInfo.WithdrawVaultAddress, opts)
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
	putils, err := services.GetPoolUtilsContract(c)
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
	withdrawVaultRewardShares, err := pool_utils.CalculateNodeElRewardShare(putils, 1, withdrawVaultBalance, nil)
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
