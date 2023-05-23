package node

import (
	"github.com/stader-labs/stader-node/shared/services"
	"github.com/stader-labs/stader-node/shared/types/api"
	"github.com/stader-labs/stader-node/stader-lib/node"
	"github.com/urfave/cli"
)

func CanWithdrawRewards(c *cli.Context) (*api.CanWithdrawRewards, error) {
	if err := services.RequireNodeWallet(c); err != nil {
		return nil, err
	}
	if err := services.RequireNodeRegistered(c); err != nil {
		return nil, err
	}
	w, err := services.GetWallet(c)
	if err != nil {
		return nil, err
	}
	orc, err := services.GetOperatorRewardsCollectorContract(c)
	if err != nil {
		return nil, err
	}

	response := api.CanWithdrawRewards{}

	opts, err := w.GetNodeAccountTransactor()
	if err != nil {
		return nil, err
	}

	// estimate gas
	gasInfo, err := node.EstimateClaimOperatorRewards(orc, opts)
	if err != nil {
		return nil, err
	}

	response.GasInfo = gasInfo

	return &response, nil
}

func WithdrawRewards(c *cli.Context) (*api.WithdrawRewards, error) {
	if err := services.RequireNodeWallet(c); err != nil {
		return nil, err
	}
	if err := services.RequireNodeRegistered(c); err != nil {
		return nil, err
	}
	w, err := services.GetWallet(c)
	if err != nil {
		return nil, err
	}
	orc, err := services.GetOperatorRewardsCollectorContract(c)
	if err != nil {
		return nil, err
	}
	pnr, err := services.GetPermissionlessNodeRegistry(c)
	if err != nil {
		return nil, err
	}

	response := api.WithdrawRewards{}

	nodeAddress, err := w.GetNodeAccount()
	if err != nil {
		return nil, err
	}

	operatorId, err := node.GetOperatorId(pnr, nodeAddress.Address, nil)
	if err != nil {
		return nil, err
	}
	operatorInfo, err := node.GetOperatorInfo(pnr, operatorId, nil)
	if err != nil {
		return nil, err
	}

	operatorRewardsBalance, err := node.GetOperatorRewardsCollectorBalance(orc, nodeAddress.Address, nil)
	if err != nil {
		return nil, err
	}

	response.OperatorRewardsBalance = operatorRewardsBalance
	response.OperatorRewardAddress = operatorInfo.OperatorRewardAddress

	opts, err := w.GetNodeAccountTransactor()
	if err != nil {
		return nil, err
	}

	// estimate gas
	tx, err := node.ClaimOperatorRewards(orc, opts)
	if err != nil {
		return nil, err
	}

	response.TxHash = tx.Hash()

	return &response, nil
}
