/*
This work is licensed and released under GNU GPL v3 or any other later versions.
The full text of the license is below/ found at <http://www.gnu.org/licenses/>

(c) 2023 Rocket Pool Pty Ltd. Modified under GNU GPL v3. [1.0.0]

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
package api

import (
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/stader-labs/stader-node/shared/services/config"
	"github.com/stader-labs/stader-node/shared/utils/log"
	"github.com/stader-labs/stader-node/shared/utils/math"
	"github.com/stader-labs/stader-node/stader-lib/stader"
	"github.com/stader-labs/stader-node/stader-lib/utils"
	"github.com/stader-labs/stader-node/stader-lib/utils/eth"
)

// The fraction of the timeout period to trigger overdue transactions
const TimeoutSafetyFactor int = 2

// Print the gas price and cost of a TX
func PrintAndCheckGasInfo(gasInfo stader.GasInfo, checkThreshold bool, gasThresholdGwei float64, logger log.ColorLogger, maxFeeWei *big.Int, gasLimit uint64) bool {

	// Check the gas threshold if requested
	if checkThreshold {
		gasThresholdWei := math.RoundUp(gasThresholdGwei*eth.WeiPerGwei, 0)
		gasThreshold := new(big.Int).SetUint64(uint64(gasThresholdWei))
		if maxFeeWei.Cmp(gasThreshold) != -1 {
			logger.Printlnf("Current network gas price is %.2f Gwei, which is not lower than the set threshold of %.2f Gwei. "+
				"Aborting the transaction.", eth.WeiToGwei(maxFeeWei), gasThresholdGwei)
			return false
		}
	} else {
		logger.Println("This transaction does not check the gas threshold limit, continuing...")
	}

	// Print the total TX cost
	var gas *big.Int
	var safeGas *big.Int
	if gasLimit != 0 {
		gas = new(big.Int).SetUint64(gasLimit)
		safeGas = gas
	} else {
		gas = new(big.Int).SetUint64(gasInfo.EstGasLimit)
		safeGas = new(big.Int).SetUint64(gasInfo.SafeGasLimit)
	}
	totalGasWei := new(big.Int).Mul(maxFeeWei, gas)
	totalSafeGasWei := new(big.Int).Mul(maxFeeWei, safeGas)
	logger.Printlnf("This transaction will use a max fee of %.6f Gwei, for a total of up to %.6f - %.6f ETH.",
		eth.WeiToGwei(maxFeeWei),
		math.RoundDown(eth.WeiToEth(totalGasWei), 6),
		math.RoundDown(eth.WeiToEth(totalSafeGasWei), 6))

	return true
}

// Print a TX's details to the logger and waits for it to validated.
func PrintAndWaitForTransaction(cfg *config.StaderConfig, hash common.Hash, ec stader.ExecutionClient, logger log.ColorLogger) error {

	txWatchUrl := cfg.StaderNode.GetTxWatchUrl()
	hashString := hash.String()

	logger.Printlnf("Transaction has been submitted with hash %s.", hashString)
	if txWatchUrl != "" {
		logger.Printlnf("You may follow its progress by visiting:")
		logger.Printlnf("%s/%s\n", txWatchUrl, hashString)
	}
	logger.Println("Waiting for the transaction to be validated...")

	// Wait for the TX to be included in a block
	if _, err := utils.WaitForTransaction(ec, hash); err != nil {
		return fmt.Errorf("Error waiting for transaction: %w", err)
	}

	return nil

}
