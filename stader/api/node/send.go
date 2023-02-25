package node

import (
	"context"
	"fmt"
	"github.com/stader-labs/stader-minipool-go/stader"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/rocket-pool/rocketpool-go/tokens"
	"github.com/rocket-pool/rocketpool-go/utils/eth"
	"github.com/urfave/cli"

	"github.com/stader-labs/stader-node/shared/services"
	"github.com/stader-labs/stader-node/shared/types/api"
	"github.com/stader-labs/stader-node/shared/utils/eth1"
)

func canNodeSend(c *cli.Context, amountWei *big.Int, token string) (*api.CanNodeSendResponse, error) {

	// Get services
	if err := services.RequireNodeWallet(c); err != nil {
		return nil, err
	}
	w, err := services.GetWallet(c)
	if err != nil {
		return nil, err
	}
	ec, err := services.GetEthClient(c)
	if err != nil {
		return nil, err
	}
	rp, err := services.GetRocketPool(c)
	if err != nil {
		return nil, err
	}

	// Response
	response := api.CanNodeSendResponse{}

	// Get node account
	nodeAccount, err := w.GetNodeAccount()
	if err != nil {
		return nil, err
	}

	// Get gas estimate
	opts, err := w.GetNodeAccountTransactor()
	if err != nil {
		return nil, err
	}

	// Handle token type
	switch token {
	case "eth":

		// Check node ETH balance
		ethBalanceWei, err := ec.BalanceAt(context.Background(), nodeAccount.Address, nil)
		if err != nil {
			return nil, err
		}
		response.InsufficientBalance = (amountWei.Cmp(ethBalanceWei) > 0)
		gasInfo, err := eth.EstimateSendTransactionGas(ec, nodeAccount.Address, opts)
		if err != nil {
			return nil, err
		}
		response.GasInfo = stader.GasInfo(gasInfo)

	case "rpl":

		// Get RocketStorage
		if err := services.RequireRocketStorage(c); err != nil {
			return nil, err
		}
		// Check node RPL balance
		rplBalanceWei, err := tokens.GetRPLBalance(rp, nodeAccount.Address, nil)
		if err != nil {
			return nil, err
		}
		response.InsufficientBalance = (amountWei.Cmp(rplBalanceWei) > 0)
		gasInfo, err := tokens.EstimateTransferRPLGas(rp, nodeAccount.Address, amountWei, opts)
		if err != nil {
			return nil, err
		}
		response.GasInfo = stader.GasInfo(gasInfo)

	case "fsrpl":

		// Get RocketStorage
		if err := services.RequireRocketStorage(c); err != nil {
			return nil, err
		}
		// Check node fixed-supply RPL balance
		fixedSupplyRplBalanceWei, err := tokens.GetFixedSupplyRPLBalance(rp, nodeAccount.Address, nil)
		if err != nil {
			return nil, err
		}
		response.InsufficientBalance = (amountWei.Cmp(fixedSupplyRplBalanceWei) > 0)
		gasInfo, err := tokens.EstimateTransferFixedSupplyRPLGas(rp, nodeAccount.Address, amountWei, opts)
		if err != nil {
			return nil, err
		}
		response.GasInfo = stader.GasInfo(gasInfo)

	case "reth":

		// Get RocketStorage
		if err := services.RequireRocketStorage(c); err != nil {
			return nil, err
		}
		// Check node rETH balance
		rethBalanceWei, err := tokens.GetRETHBalance(rp, nodeAccount.Address, nil)
		if err != nil {
			return nil, err
		}
		response.InsufficientBalance = (amountWei.Cmp(rethBalanceWei) > 0)
		gasInfo, err := tokens.EstimateTransferRETHGas(rp, nodeAccount.Address, amountWei, opts)
		if err != nil {
			return nil, err
		}
		response.GasInfo = stader.GasInfo(gasInfo)

	}

	// Update & return response
	response.CanSend = !response.InsufficientBalance
	return &response, nil

}

// TODO - bchain fix this method
func nodeSend(c *cli.Context, amountWei *big.Int, token string, to common.Address) (*api.NodeSendResponse, error) {

	// Get services
	if err := services.RequireNodeWallet(c); err != nil {
		return nil, err
	}
	w, err := services.GetWallet(c)
	if err != nil {
		return nil, err
	}
	ec, err := services.GetEthClient(c)
	if err != nil {
		return nil, err
	}
	rp, err := services.GetRocketPool(c)
	if err != nil {
		return nil, err
	}

	// Response
	response := api.NodeSendResponse{}

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

	// Handle token type
	switch token {
	case "eth":

		// Transfer ETH
		opts.Value = amountWei
		hash, err := eth.SendTransaction(ec, to, w.GetChainID(), opts)
		if err != nil {
			return nil, err
		}
		response.TxHash = hash

	case "ethx":

		// Get RocketStorage
		if err := services.RequireRocketStorage(c); err != nil {
			return nil, err
		}
		// Transfer RPL
		hash, err := tokens.TransferRPL(rp, to, amountWei, opts)
		if err != nil {
			return nil, err
		}
		response.TxHash = hash

	case "sd":

		// Get RocketStorage
		if err := services.RequireRocketStorage(c); err != nil {
			return nil, err
		}
		// Transfer fixed-supply RPL
		hash, err := tokens.TransferFixedSupplyRPL(rp, to, amountWei, opts)
		if err != nil {
			return nil, err
		}
		response.TxHash = hash

	}

	// Return response
	return &response, nil

}
