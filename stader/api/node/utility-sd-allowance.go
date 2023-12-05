package node

import (
	"fmt"
	"math/big"

	"github.com/stader-labs/stader-node/stader-lib/tokens"
	"github.com/urfave/cli"

	"github.com/stader-labs/stader-node/shared/services"
	"github.com/stader-labs/stader-node/shared/types/api"
	"github.com/stader-labs/stader-node/shared/utils/eth1"
)

func utilityAllowanceSd(c *cli.Context) (*api.SDPoolUtilitySdAllowanceResponse, error) {
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

	sduAddress, err := services.GetSdUtilityAddress(c)
	if err != nil {
		return nil, err
	}

	// Response
	response := api.SDPoolUtilitySdAllowanceResponse{}

	// Get node account
	account, err := w.GetNodeAccount()
	if err != nil {
		return nil, err
	}

	allowance, err := tokens.Allowance(sdt, account.Address, sduAddress, nil)
	if err != nil {
		return nil, err
	}

	response.Allowance = allowance

	return &response, nil
}

func utilityApproveSd(c *cli.Context, amountWei *big.Int) (*api.NodeDepositSdApproveResponse, error) {
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

	sduAddress, err := services.GetSdUtilityAddress(c)
	if err != nil {
		return nil, err
	}
	// Response
	response := api.NodeDepositSdApproveResponse{}

	opts, err := w.GetNodeAccountTransactor()
	if err != nil {
		return nil, err
	}

	err = eth1.CheckForNonceOverride(c, opts)
	if err != nil {
		return nil, fmt.Errorf("Error checking for nonce override: %w", err)
	}

	hash, err := tokens.Approve(sdt, sduAddress, amountWei, opts)
	if err != nil {
		return nil, err
	}

	response.ApproveTxHash = hash

	// Return response
	return &response, nil
}

func getUtilitySdApprovalGas(c *cli.Context, amountWei *big.Int) (*api.NodeDepositSdApproveGasResponse, error) {
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

	utilityAddress, err := services.GetPoolUtilsAddress(c)
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

	gasInfo, err := tokens.EstimateApproveGas(sdt, utilityAddress, amountWei, opts)
	if err != nil {
		return nil, err
	}

	response.GasInfo = gasInfo
	return &response, nil
}
