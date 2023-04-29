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
	"time"
)

func canRequestSdWithdraw(c *cli.Context, amountWei *big.Int) (*api.CanRequestWithdrawSdResponse, error) {
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
	response := api.CanRequestWithdrawSdResponse{}

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

	withdrawReq, err := sd_collateral.GetOperatorWithdrawInfo(sdc, nodeAccount.Address, nil)
	if err != nil {
		return nil, err
	}

	effectiveOperatorSdCollateralBalance := operatorSdCollateral
	if operatorSdCollateral.Cmp(withdrawReq.TotalSDWithdrawReqAmount) > 0 {
		effectiveOperatorSdCollateralBalance = big.NewInt(0).Sub(operatorSdCollateral, withdrawReq.TotalSDWithdrawReqAmount)
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

	if effectiveOperatorSdCollateralBalance.Cmp(thresholdSdRequiredToWithdraw) < 0 {
		response.InsufficientSdCollateral = true
		return &response, nil
	}

	if operatorSdCollateral.Cmp(amountWei) < 0 {
		response.InsufficientSdCollateral = true
		return &response, nil
	}

	opts, err := w.GetNodeAccountTransactor()
	if err != nil {
		return nil, err
	}

	gasInfo, err := sd_collateral.EstimateRequestSdCollateralWithdraw(sdc, amountWei, opts)
	if err != nil {
		return nil, err
	}
	response.GasInfo = gasInfo

	return &response, nil
}

func requestSdWithdraw(c *cli.Context, amountWei *big.Int) (*api.RequestWithdrawSdResponse, error) {
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
	response := api.RequestWithdrawSdResponse{}

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

func canClaimSd(c *cli.Context) (*api.CanClaimSdResponse, error) {
	sdc, err := services.GetSdCollateralContract(c)
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

	response := api.CanClaimSdResponse{}

	operatorWithdrawInfo, err := sd_collateral.GetOperatorWithdrawInfo(sdc, nodeAccount.Address, nil)
	if err != nil {
		return nil, err
	}
	withdrawDelay, err := sd_collateral.GetWithdrawDelay(sdc, nil)
	if err != nil {
		return nil, err
	}

	currentTime := time.Now().Unix()
	// this is already in unix timestamp
	lastWithdrawReqTimestamp := operatorWithdrawInfo.LastWithdrawReqTimestamp.Int64()

	if operatorWithdrawInfo.TotalSDWithdrawReqAmount.Cmp(big.NewInt(0)) == 0 {
		response.NoExistingClaim = true
		return &response, nil
	}
	if lastWithdrawReqTimestamp+withdrawDelay.Int64()+20 > currentTime {
		response.ClaimIsInUnbondingPeriod = true
		return &response, nil
	}

	opts, err := w.GetNodeAccountTransactor()
	if err != nil {
		return nil, err
	}

	gasInfo, err := sd_collateral.EstimateClaimWithdrawnSd(sdc, opts)
	if err != nil {
		return nil, err
	}

	response.GasInfo = gasInfo

	return &response, nil
}

func claimSd(c *cli.Context) (*api.ClaimSdResponse, error) {
	sdc, err := services.GetSdCollateralContract(c)
	if err != nil {
		return nil, err
	}
	w, err := services.GetWallet(c)
	if err != nil {
		return nil, err
	}

	// Response
	response := api.ClaimSdResponse{}

	opts, err := w.GetNodeAccountTransactor()
	if err != nil {
		return nil, err
	}
	err = eth1.CheckForNonceOverride(c, opts)
	if err != nil {
		return nil, fmt.Errorf("Error checking for nonce override: %w", err)
	}
	tx, err := sd_collateral.ClaimWithdrawnSd(sdc, opts)
	if err != nil {
		return nil, err
	}

	response.TxHash = tx.Hash()

	return &response, nil
}
