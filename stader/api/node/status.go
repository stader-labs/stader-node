package node

import (
	"github.com/stader-labs/stader-node/shared/services"
	"github.com/stader-labs/stader-node/shared/types/api"
	"github.com/stader-labs/stader-node/shared/utils/stdr"
	"github.com/stader-labs/stader-node/stader-lib/node"
	sd_collateral "github.com/stader-labs/stader-node/stader-lib/sd-collateral"
	"github.com/stader-labs/stader-node/stader-lib/tokens"
	"github.com/urfave/cli"
	"math/big"
)

func getStatus(c *cli.Context) (*api.NodeStatusResponse, error) {

	// Get services
	if err := services.RequireNodeWallet(c); err != nil {
		return nil, err
	}
	w, err := services.GetWallet(c)
	if err != nil {
		return nil, err
	}
	pnr, err := services.GetPermissionlessNodeRegistry(c)
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
	response := api.NodeStatusResponse{}

	nodeAccount, err := w.GetNodeAccount()
	if err != nil {
		return nil, err
	}

	response.AccountAddress = nodeAccount.Address

	accountEthBalance, err := tokens.GetEthBalance(pnr.Client, nodeAccount.Address, nil)
	if err != nil {
		return nil, err
	}
	accountSdBalance, err := tokens.BalanceOf(sdt, nodeAccount.Address, nil)
	if err != nil {
		return nil, err
	}

	response.AccountBalances.ETH = accountEthBalance
	response.AccountBalances.Sd = accountSdBalance

	operatorId, err := node.GetOperatorId(pnr, nodeAccount.Address, nil)
	if err != nil {
		return nil, err
	}
	operatorRegistry, err := node.GetOperatorInfo(pnr, operatorId, nil)
	if err != nil {
		return nil, err
	}

	if operatorRegistry.OperatorName != "" {
		response.Registered = true
		response.OperatorId = operatorId
		response.OperatorName = operatorRegistry.OperatorName
		response.OperatorRewardAddress = operatorRegistry.OperatorRewardAddress

		// get operator deposited sd collateral
		operatorSdCollateral, err := sd_collateral.GetOperatorSdBalance(sdc, nodeAccount.Address, nil)
		if err != nil {
			return nil, err
		}
		response.DepositedSdCollateral = operatorSdCollateral

		// total registerable validators
		totalSdWorthValidators, err := sd_collateral.GetMaxValidatorSpawnable(sdc, operatorSdCollateral, 1, nil)
		if err != nil {
			return nil, err
		}
		response.SdCollateralWorthValidators = totalSdWorthValidators

		// TODO - bchain - work on getting validator statuses
		totalValidatorKeys, err := node.GetTotalValidatorKeys(pnr, operatorId, nil)
		if err != nil {
			return nil, err
		}
		validatorInfoArray := make([]stdr.ValidatorInfo, totalValidatorKeys.Int64())

		for i := int64(0); i < totalValidatorKeys.Int64(); i++ {
			validatorIndex, err := node.GetValidatorIdByOperatorId(pnr, operatorId, big.NewInt(i), nil)
			if err != nil {
				return nil, err
			}
			validatorInfo, err := node.GetValidatorInfo(pnr, validatorIndex, nil)
			if err != nil {
				return nil, err
			}
			validatorInfoArray[i] = validatorInfo
		}

		response.ValidatorInfos = validatorInfoArray

	} else {
		response.Registered = false
	}

	// Return response
	return &response, nil
}
