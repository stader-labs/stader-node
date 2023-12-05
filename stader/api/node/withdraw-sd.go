package node

import (
	"fmt"
	"math/big"

	"github.com/stader-labs/stader-node/shared/services"
	"github.com/stader-labs/stader-node/shared/types/api"
	"github.com/stader-labs/stader-node/shared/utils/eth1"
	"github.com/stader-labs/stader-node/stader-lib/node"
	sd_collateral "github.com/stader-labs/stader-node/stader-lib/sd-collateral"
	"github.com/urfave/cli"
)

func canWithdrawSd(c *cli.Context, amountWei *big.Int) (*api.CanWithdrawSdResponse, error) {
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

	// get current sd collateral
	operatorSdCollateral, err := sd_collateral.GetOperatorSdBalance(sdc, nodeAccount.Address, nil)
	if err != nil {
		return nil, err
	}

	totalKeys, err := node.GetTotalValidatorKeys(pnr, operatorId, nil)
	if err != nil {
		return nil, err
	}
	nonTerminalKeys, err := node.GetTotalNonTerminalValidatorKeys(pnr, nodeAccount.Address, totalKeys, nil)
	if err != nil {
		return nil, err
	}
	poolThreshold, err := sd_collateral.GetPoolThreshold(sdc, 1, nil)
	if err != nil {
		return nil, err
	}
	withdrawThreshold := big.NewInt(0).Mul(poolThreshold.WithdrawThreshold, big.NewInt(int64(nonTerminalKeys)))
	withdrawThresholdInSd, err := sd_collateral.ConvertEthToSd(sdc, withdrawThreshold, nil)
	if err != nil {
		return nil, err
	}

	thresholdSdRequiredToWithdraw := withdrawThresholdInSd.Add(withdrawThresholdInSd, amountWei)

	if operatorSdCollateral.Cmp(thresholdSdRequiredToWithdraw) < 0 {
		response.InsufficientSdCollateral = true
		return &response, nil
	}

	if operatorSdCollateral.Cmp(amountWei) < 0 {
		response.InsufficientWithdrawableSd = true
		return &response, nil
	}

	opts, err := w.GetNodeAccountTransactor()
	if err != nil {
		return nil, err
	}

	gasInfo, err := sd_collateral.EstimateWithdrawSd(sdc, amountWei, opts)
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
	tx, err := sd_collateral.WithdrawSd(sdc, amountWei, opts)
	if err != nil {
		return nil, err
	}

	response.TxHash = tx.Hash()

	// Return response
	return &response, nil
}
