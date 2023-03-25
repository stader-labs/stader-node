/*
This work is licensed and released under GNU GPL v3 or any other later versions. 
The full text of the license is below/ found at <http://www.gnu.org/licenses/>

(c) 2023 Rocket Pool Pty Ltd. Modified under GNU GPL v3.

This program is free software: you can redistribute it and/or modify
it under the terms of the GNU General Public License as published by
the Free Software Foundation, either version 3 of the License, or
(at your option) any later version.

This program is distributed in the hope that it will be useful,
but WITHOUT ANY WARRANTY; without even the implied warranty of
MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
GNU General Public License for more details.

You should have received a copy of the GNU General Public License
along with this program.  If not, see <http://www.gnu.org/licenses/>.
*/
package sd_collateral

import (
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/stader-labs/stader-node/stader-lib/stader"
	"math/big"
)

func EstimateDepositSdAsCollateral(sdc *stader.SdCollateralContractManager, amount *big.Int, opts *bind.TransactOpts) (stader.GasInfo, error) {
	return sdc.SdCollateralContract.GetTransactionGasInfo(opts, "depositSDAsCollateral", amount)
}

func DepositSdAsCollateral(sdc *stader.SdCollateralContractManager, amount *big.Int, opts *bind.TransactOpts) (*types.Transaction, error) {
	tx, err := sdc.SdCollateral.DepositSDAsCollateral(opts, amount)
	if err != nil {
		return nil, err
	}

	return tx, nil
}

func GetOperatorSdBalance(sdc *stader.SdCollateralContractManager, operatorAddress common.Address, opts *bind.CallOpts) (*big.Int, error) {
	balance, err := sdc.SdCollateral.GetOperatorSDBalance(opts, operatorAddress)
	if err != nil {
		return nil, err
	}

	return balance, err
}

func HasEnoughSdCollateral(sdc *stader.SdCollateralContractManager, operatorAddress common.Address, poolType uint8, numValidators *big.Int, opts *bind.CallOpts) (bool, error) {
	hasEnoughSdCollateral, err := sdc.SdCollateral.HasEnoughSDCollateral(opts, operatorAddress, poolType, numValidators)
	if err != nil {
		return false, err
	}

	return hasEnoughSdCollateral, nil
}

func GetMaxValidatorSpawnable(sdc *stader.SdCollateralContractManager, sdAmount *big.Int, poolType uint8, opts *bind.CallOpts) (*big.Int, error) {
	maxValidatorSpawanable, err := sdc.SdCollateral.GetMaxValidatorSpawnable(opts, sdAmount, poolType)
	if err != nil {
		return nil, err
	}

	return maxValidatorSpawanable, nil
}
