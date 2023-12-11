package node

import (
	"math/big"

	"github.com/stader-labs/stader-node/stader-lib/node"
	"github.com/urfave/cli"

	"github.com/stader-labs/stader-node/shared/services"
	"github.com/stader-labs/stader-node/shared/types/api"
	"github.com/stader-labs/stader-node/stader-lib/sdutility"
)

func canUtilitySd(c *cli.Context, amountWei *big.Int) (*api.CanUtilitySDResponse, error) {
	if err := services.RequireNodeWallet(c); err != nil {
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

	prn, err := services.GetPermissionlessNodeRegistry(c)
	if err != nil {
		return nil, err
	}

	// Response
	response := api.CanUtilitySDResponse{}

	sdu, err := services.GetSdUtilityContract(c)
	if err != nil {
		return nil, err
	}

	operatorId, err := node.GetOperatorId(prn, nodeAccount.Address, nil)
	if err != nil {
		return nil, err
	}

	totalValidatorKeys, err := node.GetTotalValidatorKeys(prn, operatorId, nil)
	if err != nil {
		return nil, err
	}
	totalValidatorNonTerminalKeys, err := node.GetTotalNonTerminalValidatorKeys(prn, nodeAccount.Address, totalValidatorKeys, nil)
	if err != nil {
		return nil, err
	}

	response.NonTerminalValidators = totalValidatorNonTerminalKeys

	// Get gas estimates
	opts, err := w.GetNodeAccountTransactor()
	if err != nil {
		return nil, err
	}

	gasInfo, err := sdutility.EstimateUtilize(sdu, amountWei, opts)
	if err != nil {
		return nil, err
	}

	response.GasInfo = gasInfo

	return &response, nil
}

func utilitySd(c *cli.Context, amountWei *big.Int) (*api.NodeUtilitySDResponse, error) {
	// Get services
	if err := services.RequireNodeWallet(c); err != nil {
		return nil, err
	}

	w, err := services.GetWallet(c)
	if err != nil {
		return nil, err
	}

	sdu, err := services.GetSdUtilityContract(c)
	if err != nil {
		return nil, err
	}

	// Get gas estimates
	opts, err := w.GetNodeAccountTransactor()
	if err != nil {
		return nil, err
	}

	// Response
	response := api.NodeUtilitySDResponse{}

	tx, err := sdutility.Utilize(sdu, amountWei, opts)
	if err != nil {
		return nil, err
	}

	response.TxHash = tx.Hash()

	return &response, nil
}
