package node

import (
	"fmt"
	"github.com/ethereum/go-ethereum/common"

	"github.com/stader-labs/stader-minipool-go/node"
	"github.com/stader-labs/stader-node/shared/services"
	"github.com/stader-labs/stader-node/shared/types/api"
	"github.com/stader-labs/stader-node/shared/utils/eth1"
	"github.com/urfave/cli"
)

func canRegisterNode(c *cli.Context) (*api.CanRegisterNodeResponse, error) {

	// Get services
	if err := services.RequireNodeWallet(c); err != nil {
		return nil, err
	}
	_, err := services.GetWallet(c)
	if err != nil {
		return nil, err
	}
	sor, err := services.GetStaderOperatorRegistry(c)
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

	operatorRegistry, err := node.GetOperatorRegistry(sor, nodeAccount.Address, nil)
	if err != nil {
		return nil, err
	}

	if operatorRegistry.OperatorName != "" {
		response.AlreadyRegistered = true
		response.CanRegister = false
	} else {
		response.CanRegister = true
	}

	return &response, nil

}

func registerNode(c *cli.Context, operatorName string, operatorRewardAddress common.Address, mevSocialize bool) (*api.RegisterNodeResponse, error) {

	// Get services
	w, err := services.GetWallet(c)
	if err != nil {
		return nil, err
	}
	sor, err := services.GetStaderOperatorRegistry(c)
	if err != nil {
		return nil, err
	}

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

	//fmt.Printf("mev socialize is %d\n", mevSocialize)
	// Register node
	tx, err := node.OnboardNodeOperator(sor, false, 0, operatorName, operatorRewardAddress, opts)
	if err != nil {
		return nil, err
	}
	response.TxHash = tx.Hash()

	// Return response
	return &response, nil

}
