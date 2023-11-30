package node

import (
	"math/big"

	"github.com/stader-labs/stader-node/stader-lib/node"
	"github.com/stader-labs/stader-node/stader/api/validator"

	"github.com/urfave/cli"

	"github.com/stader-labs/stader-node/shared/services"
	"github.com/stader-labs/stader-node/shared/types/api"
)

func getSDStatus(c *cli.Context) (*api.GetSdStatusResponse, error) {
	sdc, err := services.GetSdCollateralContract(c)
	if err != nil {
		return nil, err
	}

	sdt, err := services.GetSdTokenContract(c)
	if err != nil {
		return nil, err
	}

	sdu, err := services.GetSdUtilityContract(c)
	if err != nil {
		return nil, err
	}

	w, err := services.GetWallet(c)
	if err != nil {
		return nil, err
	}

	// Get node account
	nodeAccount, err := w.GetNodeAccount()
	if err != nil {
		return nil, err
	}

	prn, err := services.GetPermissionlessNodeRegistry(c)
	if err != nil {
		return nil, err
	}

	operatorID, err := node.GetOperatorId(prn, nodeAccount.Address, nil)
	if err != nil {
		return nil, err
	}

	totalValidatorKeys, err := node.GetTotalValidatorKeys(prn, operatorID, nil)
	if err != nil {
		return nil, err
	}

	totalValidatorNonTerminalKeys, err := node.GetTotalNonTerminalValidatorKeys(prn, nodeAccount.Address, totalValidatorKeys, nil)
	if err != nil {
		return nil, err
	}

	sdStatus, err := validator.GetSDStatus(sdc, sdu, sdt, nodeAccount.Address, big.NewInt(int64(totalValidatorNonTerminalKeys)))
	if err != nil {
		return nil, err
	}

	return &api.GetSdStatusResponse{
		SDStatus: sdStatus,
	}, nil
}
