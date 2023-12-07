package node

import (
	"fmt"
	"math/big"

	sd_collateral "github.com/stader-labs/stader-node/stader-lib/sd-collateral"

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
	response := api.CanNodeDepositSdResponse{}

	// Get node account
	nodeAccount, err := w.GetNodeAccount()
	if err != nil {
		return nil, err
	}

	// Check Sd balance
	sdBalance, err := sdt.Erc20Token.BalanceOf(nil, nodeAccount.Address)
	if err != nil {
		return nil, err
	}
	if amountWei.Cmp(sdBalance) > 0 {
		response.InsufficientBalance = true
		return &response, nil
	}

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

	return &response, nil

}

func getDepositSdApprovalGas(c *cli.Context, amountWei *big.Int, contractAddress common.Address) (*api.SdApproveGasResponse, error) {
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

	// Response
	response := api.SdApproveGasResponse{}

	// Get gas estimates
	opts, err := w.GetNodeAccountTransactor()
	if err != nil {
		return nil, err
	}
	gasInfo, err := tokens.EstimateApproveGas(sdt, contractAddress, amountWei, opts)
	if err != nil {
		return nil, err
	}
	response.GasInfo = gasInfo
	return &response, nil
}

func allowanceSd(c *cli.Context, contractAddress common.Address) (*api.SdAllowanceResponse, error) {
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

	// Response
	response := api.SdAllowanceResponse{}

	// Get node account
	account, err := w.GetNodeAccount()
	if err != nil {
		return nil, err
	}

	allowance, err := tokens.Allowance(sdt, account.Address, contractAddress, nil)
	if err != nil {
		return nil, err
	}

	response.Allowance = allowance

	return &response, nil
}

func approveSd(c *cli.Context, amountWei *big.Int, contractAddress common.Address) (*api.SdApproveResponse, error) {
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

	// Response
	response := api.SdApproveResponse{}

	opts, err := w.GetNodeAccountTransactor()
	if err != nil {
		return nil, err
	}
	err = eth1.CheckForNonceOverride(c, opts)
	if err != nil {
		return nil, fmt.Errorf("Error checking for nonce override: %w", err)
	}
	hash, err := tokens.Approve(sdt, contractAddress, amountWei, opts)
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

	response.DepositTxHash = tx.Hash()

	// Return response
	return &response, nil

}
