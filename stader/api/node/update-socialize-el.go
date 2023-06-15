package node

import (
	"fmt"
	"github.com/stader-labs/stader-node/shared/services"
	"github.com/stader-labs/stader-node/shared/types/api"
	"github.com/stader-labs/stader-node/shared/utils/eth1"
	node "github.com/stader-labs/stader-node/stader-lib/node"
	stader_config "github.com/stader-labs/stader-node/stader-lib/stader-config"
	"github.com/urfave/cli"
)

func canUpdateSocializeEl(c *cli.Context, socializeEl bool) (*api.CanUpdateSocializeElResponse, error) {
	if err := services.RequireNodeWallet(c); err != nil {
		return nil, err
	}
	if err := services.RequireNodeRegistered(c); err != nil {
		return nil, err
	}
	if err := services.RequireNodeActive(c); err != nil {
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
	sdcfg, err := services.GetStaderConfigContract(c)
	if err != nil {
		return nil, err
	}
	nodeAccount, err := w.GetNodeAccount()
	if err != nil {
		return nil, err
	}
	// Response
	response := api.CanUpdateSocializeElResponse{}

	isPermissionlessRegistryPaused, err := node.IsPermissionlessNodeRegistryPaused(pnr, nil)
	if err != nil {
		return nil, err
	}
	if isPermissionlessRegistryPaused {
		response.IsPermissionlessNodeRegistryPaused = true
		return &response, nil
	}

	operatorId, err := node.GetOperatorId(pnr, nodeAccount.Address, nil)
	if err != nil {
		return nil, err
	}

	operatorInfo, err := node.GetOperatorInfo(pnr, operatorId, nil)
	if err != nil {
		return nil, err
	}

	if operatorInfo.OptedForSocializingPool && socializeEl {
		response.AlreadyOptedIn = true
		return &response, nil
	}
	if !operatorInfo.OptedForSocializingPool && !socializeEl {
		response.AlreadyOptedOut = true
		return &response, nil
	}

	currentBlock, err := eth1.GetCurrentBlockNumber(c)
	if err != nil {
		return nil, err
	}
	lastChangeBlock, err := node.GetSocializingPoolStateChangeBlock(pnr, operatorId, nil)
	if err != nil {
		return nil, err
	}
	coolDownPeriod, err := stader_config.GetSocializingPoolChangeThreshold(sdcfg, nil)
	if err != nil {
		return nil, err
	}
	if currentBlock < lastChangeBlock.Add(lastChangeBlock, coolDownPeriod).Uint64() {
		response.InCooldown = true
		response.NextUpdatableBlock = lastChangeBlock.Uint64()
		return &response, nil
	}

	opts, err := w.GetNodeAccountTransactor()
	if err != nil {
		return nil, err
	}
	gasInfo, err := node.EstimateChangeSocializingPoolState(pnr, socializeEl, opts)
	if err != nil {
		return nil, err
	}
	response.GasInfo = gasInfo

	return &response, nil
}

func updateSocializeEl(c *cli.Context, socializeEl bool) (*api.UpdateSocializeElResponse, error) {

	w, err := services.GetWallet(c)
	if err != nil {
		return nil, err
	}
	opts, err := w.GetNodeAccountTransactor()
	if err != nil {
		return nil, err
	}
	err = eth1.CheckForNonceOverride(c, opts)
	if err != nil {
		return nil, fmt.Errorf("error checking for nonce override: %w", err)
	}
	pnr, err := services.GetPermissionlessNodeRegistry(c)
	if err != nil {
		return nil, err
	}

	// Response
	response := api.UpdateSocializeElResponse{}

	tx, err := node.ChangeSocializingPoolState(pnr, socializeEl, opts)
	if err != nil {
		return nil, err
	}

	response.TxHash = tx.Hash()

	// Return response
	return &response, nil

}
