package node

import (
	"fmt"
	"github.com/stader-labs/stader-node/shared/services"
	"github.com/stader-labs/stader-node/shared/types/api"
	"github.com/stader-labs/stader-node/shared/utils/eth1"
	"github.com/stader-labs/stader-node/stader-lib/node"
	sd_collateral "github.com/stader-labs/stader-node/stader-lib/sd-collateral"
	"github.com/urfave/cli"
	"math/big"
)

func canWithdrawSd(c *cli.Context, amountWei *big.Int) (*api.CanWithdrawSdResponse, error) {
	// Get services
	w, err := services.GetWallet(c)
	if err != nil {
		return nil, err
	}
	pnr, err := services.GetPermissionlessNodeRegistry(c)
	if err != nil {
		return nil, err
	}
	sdc, err := services.GetSdCollateralContract(c)
	if err != nil {
		return nil, err
	}

	// Response
	response := api.CanWithdrawSdResponse{}

	// Get node account
	nodeAccount, err := w.GetNodeAccount()
	if err != nil {
		return nil, err
	}
	operatorId, err := node.GetOperatorId(pnr, nodeAccount.Address, nil)
	if err != nil {
		return nil, err
	}
	if operatorId.Uint64() == 0 {
		response.OperatorNotRegistered = true
		return &response, nil
	}

	// get current sd collateral
	operatorSdCollateral, err := sd_collateral.GetOperatorSdBalance(sdc, nodeAccount.Address, nil)
	if err != nil {
		return nil, err
	}
	if operatorSdCollateral.Cmp(amountWei) < 0 {
		response.InsufficientSdCollateral = true
		return &response, nil
	}

	sdCollateralAmountPostWithdrawal := operatorSdCollateral.Sub(operatorSdCollateral, amountWei)

	// get number of registered validators
	totalValidatorKeys, err := node.GetTotalValidatorKeys(pnr, operatorId, nil)
	if err != nil {
		return nil, err
	}
	totalSdWorthValidators, err := sd_collateral.GetMaxValidatorSpawnable(sdc, sdCollateralAmountPostWithdrawal, 1, nil)
	if err != nil {
		return nil, err
	}
	if totalValidatorKeys.Cmp(totalSdWorthValidators) > 0 {
		response.InsufficientSdCollateral = true
		return &response, nil
	}

	gasInfo, err := sd_collateral.EstimateRequestSdCollateralWithdraw(sdc, amountWei, nil)
	if err != nil {
		return nil, err
	}
	response.GasInfo = gasInfo

	return &response, nil
}

func withdrawSd(c *cli.Context, amountWei *big.Int) (*api.WithdrawSdResponse, error) {
	// Get services
	w, err := services.GetWallet(c)
	if err != nil {
		return nil, err
	}
	sdc, err := services.GetSdCollateralContract(c)
	if err != nil {
		return nil, err
	}

	// Response
	response := api.WithdrawSdResponse{}

	opts, err := w.GetNodeAccountTransactor()
	if err != nil {
		return nil, err
	}
	err = eth1.CheckForNonceOverride(c, opts)
	if err != nil {
		return nil, fmt.Errorf("Error checking for nonce override: %w", err)
	}
	tx, err := sd_collateral.RequestSdCollateralWithdraw(sdc, amountWei, opts)
	if err != nil {
		return nil, err
	}

	response.TxHash = tx.Hash()

	// Return response
	return &response, nil
}
