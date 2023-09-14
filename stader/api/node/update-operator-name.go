package node

import (
	"fmt"

	"github.com/stader-labs/stader-node/shared/services"
	"github.com/stader-labs/stader-node/shared/types/api"
	"github.com/stader-labs/stader-node/shared/utils/eth1"
	"github.com/stader-labs/stader-node/stader-lib/node"
	stader_config "github.com/stader-labs/stader-node/stader-lib/stader-config"
	"github.com/urfave/cli"
)

func CanUpdateOperatorName(c *cli.Context, operatorName string) (*api.CanUpdateOperatorName, error) {
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

	response := api.CanUpdateOperatorName{}

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
	if !operatorInfo.Active {
		response.OperatorNotActive = true
		return &response, nil
	}

	if operatorInfo.OperatorName == operatorName {
		response.NothingToUpdate = true
		return &response, nil
	}

	maxNameLength, err := stader_config.GetOperatorNameMaxLength(sdcfg, nil)
	if err != nil {
		return nil, err
	}
	if len(operatorName) > int(maxNameLength.Int64()) {
		response.OperatorNameTooLong = true
		return &response, nil
	}

	opts, err := w.GetNodeAccountTransactor()
	if err != nil {
		return nil, err
	}

	// estimate gas
	gasInfo, err := node.EstimateUpdateOperatorDetails(pnr, operatorName, operatorInfo.OperatorRewardAddress, opts)
	if err != nil {
		return nil, err
	}

	response.GasInfo = gasInfo

	return &response, nil
}

func UpdateOperatorName(c *cli.Context, operatorName string) (*api.UpdateOperatorName, error) {
	w, err := services.GetWallet(c)
	if err != nil {
		return nil, err
	}
	pnr, err := services.GetPermissionlessNodeRegistry(c)
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

	response := api.UpdateOperatorName{}

	tx, err := node.UpdateOperatorName(pnr, operatorName, opts)
	if err != nil {
		return nil, err
	}

	response.TxHash = tx.Hash()

	return &response, nil
}
