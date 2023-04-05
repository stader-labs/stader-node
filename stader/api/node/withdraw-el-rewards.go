package node

import (
	"github.com/stader-labs/stader-node/shared/services"
	"github.com/stader-labs/stader-node/shared/types/api"
	"github.com/stader-labs/stader-node/stader-lib/node"
	"github.com/stader-labs/stader-node/stader-lib/tokens"
	"github.com/urfave/cli"
)

func CanWithdrawElRewards(c *cli.Context) (*api.CanWithdrawElRewardsResponse, error) {
	// Get services
	if err := services.RequireNodeRegistered(c); err != nil {
		return nil, err
	}
	w, err := services.GetWallet(c)
	if err != nil {
		return nil, err
	}
	pnr, err := services.GetPermissionlessNodeRegistry(c)
	if err != nil {
		return nil, err
	}
	vf, err := services.GetVaultFactory(c)
	if err != nil {
		return nil, err
	}

	nodeAccount, err := w.GetNodeAccount()
	if err != nil {
		return nil, err
	}
	// Response
	response := api.CanWithdrawElRewardsResponse{}

	operatorId, err := node.GetOperatorId(pnr, nodeAccount.Address, nil)
	if err != nil {
		return nil, err
	}

	operatorElRewardAddress, err := node.GetNodeElRewardAddress(vf, 1, operatorId, nil)
	if err != nil {
		return nil, err
	}
	elRewardAddressBalance, err := tokens.GetEthBalance(pnr.Client, operatorElRewardAddress, nil)
	if err != nil {
		return nil, err
	}
	operatorElRewards, err := node.CalculateElRewardShare(pnr.Client, operatorElRewardAddress, elRewardAddressBalance, nil)
	if err != nil {
		return nil, err
	}
	if operatorElRewards.OperatorShare.Int64() < 0 {
		response.NoElRewards = true
		return &response, nil
	}

	gasInfo, err := node.EstimateWithdrawFromNodeElVault(pnr.Client, nodeAccount.Address, nil)
	if err != nil {
		return nil, err
	}
	response.GasInfo = gasInfo

	return &response, nil
}

func WithdrawElRewards(c *cli.Context) (*api.WithdrawElRewardsResponse, error) {
	w, err := services.GetWallet(c)
	if err != nil {
		return nil, err
	}
	pnr, err := services.GetPermissionlessNodeRegistry(c)
	if err != nil {
		return nil, err
	}
	vf, err := services.GetVaultFactory(c)
	if err != nil {
		return nil, err
	}
	nodeAccount, err := w.GetNodeAccount()
	if err != nil {
		return nil, err
	}

	response := api.WithdrawElRewardsResponse{}

	operatorId, err := node.GetOperatorId(pnr, nodeAccount.Address, nil)
	if err != nil {
		return nil, err
	}
	operatorInfo, err := node.GetOperatorInfo(pnr, operatorId, nil)
	if err != nil {
		return nil, err
	}
	operatorElRewardAddress, err := node.GetNodeElRewardAddress(vf, 1, operatorId, nil)
	if err != nil {
		return nil, err
	}
	elRewardAddressBalance, err := tokens.GetEthBalance(pnr.Client, operatorElRewardAddress, nil)
	if err != nil {
		return nil, err
	}
	operatorElRewards, err := node.CalculateElRewardShare(pnr.Client, operatorElRewardAddress, elRewardAddressBalance, nil)
	if err != nil {
		return nil, err
	}

	response.ElRewardsAmount = operatorElRewards.OperatorShare
	response.OperatorRewardAddress = operatorInfo.OperatorRewardAddress

	tx, err := node.WithdrawFromNodeElVault(pnr.Client, nodeAccount.Address, nil)
	if err != nil {
		return nil, err
	}

	response.TxHash = tx.Hash()

	return &response, nil
}
