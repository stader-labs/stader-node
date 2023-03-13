package node

import (
	"context"
	"fmt"
	"github.com/stader-labs/stader-node/stader-lib/stader"
	"github.com/stader-labs/stader-node/stader-lib/utils/eth"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/stader-labs/stader-node/stader-lib/tokens"
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
	sdt, err := services.GetSdTokenContract(c)
	if err != nil {
		return nil, err
	}
	ethxt, err := services.GetEthxTokenContract(c)
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
		response.InsufficientBalance = amountWei.Cmp(ethBalanceWei) > 0
		gasInfo, err := eth.EstimateSendTransactionGas(ec, nodeAccount.Address, opts)
		if err != nil {
			return nil, err
		}
		response.GasInfo = stader.GasInfo(gasInfo)

	case "sd":

		// Check node SD balance
		sdBalanceWei, err := tokens.BalanceOf(sdt, nodeAccount.Address, nil)
		if err != nil {
			return nil, err
		}
		response.InsufficientBalance = amountWei.Cmp(sdBalanceWei) > 0
		gasInfo, err := tokens.EstimateTransferGas(sdt, nodeAccount.Address, amountWei, opts)
		if err != nil {
			return nil, err
		}
		response.GasInfo = gasInfo

	case "ethx":

		// Check node EthX balance
		ethxBalanceWei, err := tokens.BalanceOf(ethxt, nodeAccount.Address, nil)
		if err != nil {
			return nil, err
		}
		response.InsufficientBalance = amountWei.Cmp(ethxBalanceWei) > 0
		gasInfo, err := tokens.EstimateTransferGas(ethxt, nodeAccount.Address, amountWei, opts)
		if err != nil {
			return nil, err
		}
		response.GasInfo = gasInfo

	}

	// Update & return response
	response.CanSend = !response.InsufficientBalance
	return &response, nil

}

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
	sdt, err := services.GetSdTokenContract(c)
	if err != nil {
		return nil, err
	}
	ethxt, err := services.GetEthxTokenContract(c)
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

	case "sd":

		// Transfer SD
		hash, err := tokens.Transfer(sdt, to, amountWei, opts)
		if err != nil {
			return nil, err
		}
		response.TxHash = hash

	case "ethx":

		// Transfer Ethx
		hash, err := tokens.Transfer(ethxt, to, amountWei, opts)
		if err != nil {
			return nil, err
		}
		response.TxHash = hash

	}

	// Return response
	return &response, nil

}
