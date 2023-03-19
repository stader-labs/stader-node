package node

import (
	"fmt"
	"github.com/stader-labs/stader-node/stader-lib/sd-collateral"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/stader-labs/stader-node/stader-lib/tokens"
	"github.com/stader-labs/stader-node/stader-lib/utils"
	"github.com/urfave/cli"

	"github.com/stader-labs/stader-node/shared/services"
	"github.com/stader-labs/stader-node/shared/types/api"
	"github.com/stader-labs/stader-node/shared/utils/eth1"
)

func canNodeDepositSd(c *cli.Context, amountWei *big.Int) (*api.CanNodeDepositSdResponse, error) {

	// Get services
	w, err := services.GetWallet(c)
	if err != nil {
		return nil, err
	}
	sdt, err := services.GetSdTokenContract(c)
	if err != nil {
		return nil, err
	}
	sdc, err := services.GetSdCollateralContract(c)
	if err != nil {
		return nil, err
	}

	// Response
	response := api.CanNodeDepositSdResponse{}

	// Get node account
	nodeAccount, err := w.GetNodeAccount()
	if err != nil {
		return nil, err
	}

	// Check Sd balance
	// TODO - bchain we should abstract this out
	sdBalance, err := sdt.Erc20Token.BalanceOf(nil, nodeAccount.Address)
	if err != nil {
		return nil, err
	}
	response.InsufficientBalance = amountWei.Cmp(sdBalance) > 0

	// Get gas estimates
	opts, err := w.GetNodeAccountTransactor()
	if err != nil {
		return nil, err
	}
	gasInfo, err := sd_collateral.EstimateDepositSdAsCollateral(sdc, amountWei, opts)
	if err != nil {
		return nil, err
	}
	response.GasInfo = gasInfo

	// Update & return response
	response.CanStake = !response.InsufficientBalance
	return &response, nil

}

func getDepositSdApprovalGas(c *cli.Context, amountWei *big.Int) (*api.NodeDepositSdApproveGasResponse, error) {
	// Get services
	if err := services.RequireNodeWallet(c); err != nil {
		return nil, err
	}
	w, err := services.GetWallet(c)
	if err != nil {
		return nil, err
	}
	sdt, err := services.GetSdTokenContract(c)
	if err != nil {
		return nil, err
	}
	sdc, err := services.GetSdCollateralContract(c)
	if err != nil {
		return nil, err
	}

	// Response
	response := api.NodeDepositSdApproveGasResponse{}

	// Get gas estimates
	opts, err := w.GetNodeAccountTransactor()
	if err != nil {
		return nil, err
	}
	gasInfo, err := tokens.EstimateApproveGas(sdt, *sdc.SdCollateralContract.Address, amountWei, opts)
	if err != nil {
		return nil, err
	}
	response.GasInfo = gasInfo
	return &response, nil
}

func allowanceSd(c *cli.Context) (*api.NodeDepositSdAllowanceResponse, error) {

	// Get services
	if err := services.RequireNodeRegistered(c); err != nil {
		return nil, err
	}
	w, err := services.GetWallet(c)
	if err != nil {
		return nil, err
	}
	sdt, err := services.GetSdTokenContract(c)
	if err != nil {
		return nil, err
	}
	sdc, err := services.GetSdCollateralContract(c)
	if err != nil {
		return nil, err
	}

	// Response
	response := api.NodeDepositSdAllowanceResponse{}

	// Get node account
	account, err := w.GetNodeAccount()
	if err != nil {
		return nil, err
	}

	// Get node's RPL allowance
	allowance, err := tokens.Allowance(sdt, account.Address, *sdc.SdCollateralContract.Address, nil)
	if err != nil {
		return nil, err
	}

	response.Allowance = allowance

	return &response, nil
}

func approveSd(c *cli.Context, amountWei *big.Int) (*api.NodeDepositSdApproveResponse, error) {

	// Get services
	if err := services.RequireNodeRegistered(c); err != nil {
		return nil, err
	}
	w, err := services.GetWallet(c)
	if err != nil {
		return nil, err
	}
	sdt, err := services.GetSdTokenContract(c)
	if err != nil {
		return nil, err
	}
	sdc, err := services.GetSdCollateralContract(c)
	if err != nil {
		return nil, err
	}
	// Response
	response := api.NodeDepositSdApproveResponse{}

	// Approve RPL allowance
	opts, err := w.GetNodeAccountTransactor()
	if err != nil {
		return nil, err
	}
	err = eth1.CheckForNonceOverride(c, opts)
	if err != nil {
		return nil, fmt.Errorf("Error checking for nonce override: %w", err)
	}
	hash, err := tokens.Approve(sdt, *sdc.SdCollateralContract.Address, amountWei, opts)
	if err != nil {
		return nil, err
	}

	response.ApproveTxHash = hash

	// Return response
	return &response, nil

}

func waitForApprovalAndDepositSd(c *cli.Context, amountWei *big.Int, hash common.Hash) (*api.NodeDepositSdResponse, error) {

	// Get services
	if err := services.RequireNodeRegistered(c); err != nil {
		return nil, err
	}
	prn, err := services.GetPermissionlessNodeRegistry(c)
	if err != nil {
		return nil, err
	}

	// Wait for the RPL approval TX to successfully get included in a block
	_, err = utils.WaitForTransaction(prn.Client, hash)
	if err != nil {
		return nil, err
	}

	// Perform the stake
	return depositSdAsCollateral(c, amountWei)

}

func depositSdAsCollateral(c *cli.Context, amountWei *big.Int) (*api.NodeDepositSdResponse, error) {

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
	response := api.NodeDepositSdResponse{}

	// Stake RPL
	opts, err := w.GetNodeAccountTransactor()
	if err != nil {
		return nil, err
	}
	err = eth1.CheckForNonceOverride(c, opts)
	if err != nil {
		return nil, fmt.Errorf("Error checking for nonce override: %w", err)
	}
	tx, err := sd_collateral.DepositSdAsCollateral(sdc, amountWei, opts)
	if err != nil {
		return nil, err
	}

	response.StakeTxHash = tx.Hash()

	// Return response
	return &response, nil

}
