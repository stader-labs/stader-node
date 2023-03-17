package api

import (
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/stader-labs/stader-node/shared/services/config"
	"github.com/stader-labs/stader-node/shared/utils/log"
	"github.com/stader-labs/stader-node/shared/utils/math"
	"github.com/stader-labs/stader-node/stader-lib/stader"
	"github.com/stader-labs/stader-node/stader-lib/utils"
	"github.com/stader-labs/stader-node/stader-lib/utils/eth"
	"math/big"
)

// TODO - ROCKETPOOL-OWNER

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

	txWatchUrl := cfg.Smartnode.GetTxWatchUrl()
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
