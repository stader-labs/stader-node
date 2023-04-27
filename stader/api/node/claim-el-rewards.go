package node

import (
	"github.com/stader-labs/stader-node/shared/services"
	"github.com/stader-labs/stader-node/shared/types/api"
	"github.com/stader-labs/stader-node/stader-lib/node"
	pool_utils "github.com/stader-labs/stader-node/stader-lib/pool-utils"
	"github.com/stader-labs/stader-node/stader-lib/tokens"
	"github.com/urfave/cli"
	"math/big"
)

func CanClaimElRewards(c *cli.Context) (*api.CanClaimElRewardsResponse, error) {
	if err := services.RequireNodeWallet(c); err != nil {
		return nil, err
	}
	if err := services.RequireNodeRegistered(c); err != nil {
		return nil, err
	}
	// Get services
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
	putils, err := services.GetPoolUtilsContract(c)
	if err != nil {
		return nil, err
	}

	nodeAccount, err := w.GetNodeAccount()
	if err != nil {
		return nil, err
	}
	// Response
	response := api.CanClaimElRewardsResponse{}

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
	operatorElRewards, err := pool_utils.CalculateRewardShare(putils, 1, elRewardAddressBalance, nil)
	if err != nil {
		return nil, err
	}

	if operatorElRewards.OperatorShare.Cmp(new(big.Int).SetUint64(1000000000)) < 0 {
		response.NoElRewards = true
		return &response, nil
	}

	opts, err := w.GetNodeAccountTransactor()
	if err != nil {
		return nil, err
	}

	gasInfo, err := node.EstimateWithdrawFromNodeElVault(pnr.Client, operatorElRewardAddress, opts)
	if err != nil {
		return nil, err
	}
	response.GasInfo = gasInfo

	return &response, nil
}

func ClaimElRewards(c *cli.Context) (*api.ClaimElRewardsResponse, error) {
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
	putils, err := services.GetPoolUtilsContract(c)
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

	response := api.ClaimElRewardsResponse{}

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
	operatorElRewards, err := pool_utils.CalculateRewardShare(putils, 1, elRewardAddressBalance, nil)
	if err != nil {
		return nil, err
	}

	response.ElRewardsAmount = operatorElRewards.OperatorShare
	response.OperatorRewardAddress = operatorInfo.OperatorRewardAddress

	tx, err := node.WithdrawFromNodeElVault(pnr.Client, operatorElRewardAddress, opts)
	if err != nil {
		return nil, err
	}

	response.TxHash = tx.Hash()

	return &response, nil
}
