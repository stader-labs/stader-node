package node

import (
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	pool_utils "github.com/stader-labs/stader-node/stader-lib/pool-utils"
	stader_config "github.com/stader-labs/stader-node/stader-lib/stader-config"

	"github.com/stader-labs/stader-node/shared/services"
	"github.com/stader-labs/stader-node/shared/types/api"
	"github.com/stader-labs/stader-node/shared/utils/eth1"
	"github.com/stader-labs/stader-node/stader-lib/node"
	"github.com/urfave/cli"
)

func canRegisterNode(c *cli.Context, operatorName string, operatorRewardAddress common.Address, socializeMev bool) (*api.CanRegisterNodeResponse, error) {

	// Get services
	if err := services.RequireNodeWallet(c); err != nil {
		return nil, err
	}
	_, err := services.GetWallet(c)
	if err != nil {
		return nil, err
	}
	pnr, err := services.GetPermissionlessNodeRegistry(c)
	sdcfg, err := services.GetStaderConfigContract(c)
	if err != nil {
		return nil, err
	}
	putils, err := services.GetPoolUtilsContract(c)
	if err != nil {
		return nil, err
	}
	w, err := services.GetWallet(c)
	if err != nil {
		return nil, err
	}

	// Response
	response := api.CanRegisterNodeResponse{}

	nodeAccount, err := w.GetNodeAccount()

	isPermissionlessRegistryPaused, err := node.IsPermissionlessNodeRegistryPaused(pnr, nil)
	if err != nil {
		return nil, err
	}
	if isPermissionlessRegistryPaused {
		response.RegistrationPaused = true
		return &response, nil
	}

	// this checks across all pools
	isExistingOperator, err := pool_utils.IsExistingOperator(putils, nodeAccount.Address, nil)
	if err != nil {
		return nil, err
	}
	if isExistingOperator {
		response.AlreadyRegistered = true
		return &response, nil
	}

	operatorNameMaxLength, err := stader_config.GetOperatorNameMaxLength(sdcfg, nil)
	if err != nil {
		return nil, err
	}
	if len(operatorName) > int(operatorNameMaxLength.Int64()) {
		response.OperatorNameTooLong = true
		return &response, nil
	}
	if eth1.IsZeroAddress(operatorRewardAddress) {
		response.OperatorRewardAddressZero = true
		return &response, nil
	}

	opts, err := w.GetNodeAccountTransactor()
	if err != nil {
		return nil, err
	}

	gasInfo, err := node.EstimateOnboardNodeOperator(pnr, socializeMev, operatorName, operatorRewardAddress, opts)
	if err != nil {
		return nil, err
	}

	response.GasInfo = gasInfo

	return &response, nil

}

func registerNode(c *cli.Context, operatorName string, operatorRewardAddress common.Address, mevSocialize bool) (*api.RegisterNodeResponse, error) {

	// Get services
	w, err := services.GetWallet(c)
	if err != nil {
		return nil, err
	}
	prn, err := services.GetPermissionlessNodeRegistry(c)

	// Response
	response := api.RegisterNodeResponse{}

	// Get transactor
	opts, err := w.GetNodeAccountTransactor()
	if err != nil {
		return nil, err
	}

	// Override the provided pending TX if requested
	err = eth1.CheckForNonceOverride(c, opts)
	if err != nil {
		return nil, fmt.Errorf("Error checking for nonce override: %w", err)
	}

	// Register node
	tx, err := node.OnboardNodeOperator(prn, mevSocialize, operatorName, operatorRewardAddress, opts)
	if err != nil {
		return nil, err
	}
	response.TxHash = tx.Hash()

	// Return response
	return &response, nil

}
