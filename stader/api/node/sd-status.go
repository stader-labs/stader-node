package node

import (
	"math/big"

	"github.com/stader-labs/stader-node/stader-lib/node"
	sd_collateral "github.com/stader-labs/stader-node/stader-lib/sd-collateral"
	"github.com/stader-labs/stader-node/stader/api/validator"

	"github.com/urfave/cli"

	"github.com/stader-labs/stader-node/shared/services"
	"github.com/stader-labs/stader-node/shared/types/api"
)

func getSDStatus(c *cli.Context, numValidators *big.Int) (*api.GetSdStatusResponse, error) {
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

	numValidatorsPostAdd := new(big.Int).Add(numValidators, big.NewInt(int64(totalValidatorNonTerminalKeys)))

	sdStatus, err := validator.GetSDStatus(sdc, sdu, sdt, nodeAccount.Address, numValidatorsPostAdd)
	if err != nil {
		return nil, err
	}

	hasEnoughSdCollateral, err := sd_collateral.HasEnoughSdCollateral(sdc, nodeAccount.Address, 1, numValidatorsPostAdd, nil)

	if err != nil {
		return nil, err
	}

	sdStatus.NotEnoughSdCollateral = !hasEnoughSdCollateral
	return &api.GetSdStatusResponse{
		SDStatus: sdStatus,
	}, nil
}
