package node

import (
	"math/big"

	"github.com/stader-labs/stader-node/shared/services"
	"github.com/stader-labs/stader-node/shared/types/api"
	"github.com/stader-labs/stader-node/stader-lib/node"
	"github.com/urfave/cli"
)

func CanClaimRewards(c *cli.Context) (*api.CanClaimRewards, error) {
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
	nodeAccount, err := w.GetNodeAccount()
	if err != nil {
		return nil, err
	}

	response := api.CanClaimRewards{}

	operatorClaimVaultBalance, err := node.GetOperatorRewardsCollectorBalance(orc, nodeAccount.Address, nil)
	if err != nil {
		return nil, err
	}
	if operatorClaimVaultBalance.Cmp(big.NewInt(0)) == 0 {
		response.NoRewards = true
		return &response, nil
	}

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

func ClaimRewards(c *cli.Context) (*api.ClaimRewards, error) {
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

	response := api.ClaimRewards{}

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
