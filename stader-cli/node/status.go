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
package node

import (
	"fmt"
	"github.com/stader-labs/stader-node/shared/utils/log"
	"github.com/stader-labs/stader-node/shared/utils/stdr"
	"github.com/stader-labs/stader-node/stader-lib/types"
	"github.com/stader-labs/stader-node/stader-lib/utils/eth"
	"github.com/urfave/cli"
	"time"

	"github.com/stader-labs/stader-node/shared/services/stader"
	cliutils "github.com/stader-labs/stader-node/shared/utils/cli"
	"github.com/stader-labs/stader-node/shared/utils/math"
)

func getStatus(c *cli.Context) error {

	staderClient, err := stader.NewClientFromCtx(c)
	if err != nil {
		return err
	}
	defer staderClient.Close()

	// Check and assign the EC status
	err = cliutils.CheckClientStatus(staderClient)
	if err != nil {
		return err
	}

	// Print what network we're on
	err = cliutils.PrintNetwork(staderClient)
	if err != nil {
		return err
	}

	// Get node status
	status, err := staderClient.NodeStatus()
	if err != nil {
		return err
	}

	totalRegisteredValidators := len(status.ValidatorInfos)
	totalRegisterableValidators := status.SdCollateralWorthValidators

	noOfValidatorsWhichWeCanRegister := int64(0)
	if totalRegisterableValidators.Int64() > int64(totalRegisteredValidators) {
		noOfValidatorsWhichWeCanRegister = totalRegisterableValidators.Int64() - int64(totalRegisteredValidators)
	}

	// Account address & balances
	fmt.Printf("%s=== Account and Balances ===%s\n", log.ColorGreen, log.ColorReset)
	fmt.Printf(
		"The node %s%s%s has a balance of %.6f ETH.\n\n",
		log.ColorBlue,
		status.AccountAddress,
		log.ColorReset,
		math.RoundDown(eth.WeiToEth(status.AccountBalances.ETH), 6))
	fmt.Printf(
		"The node %s%s%s has a balance of %.6f SD.\n\n",
		log.ColorBlue,
		status.AccountAddress,
		log.ColorReset,
		math.RoundDown(eth.WeiToEth(status.AccountBalances.Sd), 18))
	fmt.Printf(
		"The node %s%s%s has a deposited %.6f SD as collateral.\n\n",
		log.ColorBlue,
		status.AccountAddress,
		log.ColorReset,
		math.RoundDown(eth.WeiToEth(status.DepositedSdCollateral), 18))
	fmt.Printf(
		"The node %s%s%s has registered %d validators.\n\n",
		log.ColorBlue,
		status.AccountAddress,
		log.ColorReset,
		totalRegisteredValidators)
	fmt.Printf(
		"The node %s%s%s can register %d more validators based on the amount of SD collateral it has provided.\n\n",
		log.ColorBlue,
		status.AccountAddress,
		log.ColorReset,
		noOfValidatorsWhichWeCanRegister)

	fmt.Printf("%s=== Operator Registration Details ===%s\n", log.ColorGreen, log.ColorReset)

	if status.Registered {
		fmt.Printf("The node is registered with Stader. Below are node details:\n")
		fmt.Printf("Operator Id: %d\n\n", status.OperatorId)
		fmt.Printf("Operator Name: %s\n\n", status.OperatorName)
		fmt.Printf("Operator Reward Address: %s\n\n", status.OperatorRewardAddress.String())

		fmt.Printf("%s=== Registered Validator Details ===%s\n", log.ColorGreen, log.ColorReset)
		// display validators
		if totalRegisteredValidators > 0 {
			for i := 0; i < totalRegisteredValidators; i++ {
				fmt.Printf("%d)\n", i+1)
				validatorInfo := status.ValidatorInfos[i]
				fmt.Printf("-Validator Pub Key: %s\n\n", types.BytesToValidatorPubkey(validatorInfo.Pubkey))
				fmt.Printf("-Validator Status: %s\n\n", stdr.ValidatorState[validatorInfo.Status])
				fmt.Printf("-Validator Withdraw Vault: %s\n\n", validatorInfo.WithdrawVaultAddress)

				if validatorInfo.Status > 3 {
					fmt.Printf("-Deposit time: %s\n\n", time.Unix(validatorInfo.DepositTime.Int64(), 0))
				}

				// Validator has withdrawn
				if validatorInfo.Status == 8 {
					fmt.Printf("-Withdrawn time: %d\n\n", time.Unix(validatorInfo.WithdrawnTime.Int64(), 0))
				}
				fmt.Printf("\n\n")
			}
		} else {
			fmt.Printf("The node has no registered validators. Please use the %sstader-cli node deposit%s command to register a validator with Stader", log.ColorGreen, log.ColorReset)
		}
	} else {
		fmt.Printf("The node is not registered with Stader. Please use the %sstader-cli node register%s to register with Stader", log.ColorGreen, log.ColorReset)
	}

	// Return
	return nil

}
