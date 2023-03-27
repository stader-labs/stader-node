/*
This work is licensed and released under GNU GPL v3 or any other later versions.
The full text of the license is below/ found at <http://www.gnu.org/licenses/>

(c) 2023 Rocket Pool Pty Ltd. Modified under GNU GPL v3. [0.3.0-beta]

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
package gas

import (
	"fmt"
	"math/big"
	"strconv"

	"github.com/stader-labs/stader-node/shared/utils/log"

	"github.com/stader-labs/stader-node/shared/services/gas/etherchain"
	"github.com/stader-labs/stader-node/shared/services/gas/etherscan"
	"github.com/stader-labs/stader-node/shared/services/stader"
	cliutils "github.com/stader-labs/stader-node/shared/utils/cli"
	"github.com/stader-labs/stader-node/shared/utils/math"
	staderCore "github.com/stader-labs/stader-node/stader-lib/stader"
	"github.com/stader-labs/stader-node/stader-lib/utils/eth"
)

func AssignMaxFeeAndLimit(gasInfo staderCore.GasInfo, staderClient *stader.Client, headless bool) error {

	cfg, isNew, err := staderClient.LoadConfig()
	if err != nil {
		return fmt.Errorf("Error getting Stader configuration: %w", err)
	}
	if isNew {
		return fmt.Errorf("couldn't locate settings file. Execute the stader-cli service config command to set up your Stader Node")
	}

	// Get the current settings from the CLI arguments
	maxFeeGwei, maxPriorityFeeGwei, gasLimit := staderClient.GetGasSettings()

	// Get the max fee - prioritize the CLI arguments, default to the config file setting
	if maxFeeGwei == 0 {
		maxFee := eth.GweiToWei(cfg.StaderNode.ManualMaxFee.Value.(float64))
		if maxFee != nil && maxFee.Uint64() != 0 {
			maxFeeGwei = eth.WeiToGwei(maxFee)
		}
	}

	// Get the priority fee - prioritize the CLI arguments, default to the config file setting
	if maxPriorityFeeGwei == 0 {
		maxPriorityFee := eth.GweiToWei(cfg.StaderNode.PriorityFee.Value.(float64))
		if maxPriorityFee == nil || maxPriorityFee.Uint64() == 0 {
			fmt.Printf("%sNOTE: max priority fee not set or set to 0, defaulting to 2 gwei%s\n", log.ColorYellow, log.ColorReset)
			maxPriorityFeeGwei = 2
		} else {
			maxPriorityFeeGwei = eth.WeiToGwei(maxPriorityFee)
		}
	}

	// Use the requested max fee and priority fee if provided
	if maxFeeGwei != 0 {
		fmt.Printf("%sUsing the requested max fee of %.2f gwei (including a max priority fee of %.2f gwei).\n", log.ColorYellow, maxFeeGwei, maxPriorityFeeGwei)

		var lowLimit float64
		var highLimit float64
		if gasLimit == 0 {
			lowLimit = maxFeeGwei / eth.WeiPerGwei * float64(gasInfo.EstGasLimit)
			highLimit = maxFeeGwei / eth.WeiPerGwei * float64(gasInfo.SafeGasLimit)
		} else {
			lowLimit = maxFeeGwei / eth.WeiPerGwei * float64(gasLimit)
			highLimit = lowLimit
		}
		fmt.Printf("Total cost: %.4f to %.4f ETH%s\n", lowLimit, highLimit, log.ColorReset)

	} else {
		if headless {
			maxFeeWei, err := GetHeadlessMaxFeeWei()
			if err != nil {
				return err
			}
			maxFeeGwei = eth.WeiToGwei(maxFeeWei)
		} else {
			// Try to get the latest gas prices from Etherchain
			etherchainData, err := etherchain.GetGasPrices()
			if err == nil {
				// Print the Etherchain data and ask for an amount
				maxFeeGwei = handleEtherchainGasPrices(etherchainData, gasInfo, maxPriorityFeeGwei, gasLimit)

			} else {
				// Fallback to Etherscan
				fmt.Printf("%sWarning: couldn't get gas estimates from Etherchain - %s\nFalling back to Etherscan%s\n", log.ColorYellow, err.Error(), log.ColorReset)
				etherscanData, err := etherscan.GetGasPrices()
				if err == nil {
					// Print the Etherscan data and ask for an amount
					maxFeeGwei = handleEtherscanGasPrices(etherscanData, gasInfo, maxPriorityFeeGwei, gasLimit)
				} else {
					return fmt.Errorf("Error getting gas price suggestions: %w", err)
				}
			}
		}
		fmt.Printf("%sUsing a max fee of %.2f gwei and a priority fee of %.2f gwei.\n%s", log.ColorBlue, maxFeeGwei, maxPriorityFeeGwei, log.ColorReset)
	}

	// Use the requested gas limit if provided
	if gasLimit != 0 {
		fmt.Printf("Using the requested gas limit of %d units.\n%sNOTE: if you set this too low, your transaction may fail but you will still have to pay the gas fee!%s\n", gasLimit, log.ColorYellow, log.ColorReset)
	}

	if maxPriorityFeeGwei > maxFeeGwei {
		return fmt.Errorf("Priority fee cannot be greater than max fee.")
	}
	staderClient.AssignGasSettings(maxFeeGwei, maxPriorityFeeGwei, gasLimit)
	return nil

}

// Get the suggested max fee for service operations
func GetHeadlessMaxFeeWei() (*big.Int, error) {
	etherchainData, err := etherchain.GetGasPrices()
	if err == nil {
		return etherchainData.RapidWei, nil
	}

	fmt.Printf("%sWarning: couldn't get gas estimates from Etherchain - %s\nFalling back to Etherscan%s\n", log.ColorYellow, err.Error(), log.ColorReset)
	etherscanData, err := etherscan.GetGasPrices()
	if err == nil {
		return eth.GweiToWei(etherscanData.FastGwei), nil
	}

	return nil, fmt.Errorf("Error getting gas price suggestions: %w", err)
}

func handleEtherchainGasPrices(gasSuggestion etherchain.GasFeeSuggestion, gasInfo staderCore.GasInfo, priorityFee float64, gasLimit uint64) float64 {

	rapidGwei := math.RoundUp(eth.WeiToGwei(gasSuggestion.RapidWei)+priorityFee, 0)
	rapidEth := eth.WeiToEth(gasSuggestion.RapidWei)

	var rapidLowLimit float64
	var rapidHighLimit float64
	if gasLimit == 0 {
		rapidLowLimit = rapidEth * float64(gasInfo.EstGasLimit)
		rapidHighLimit = rapidEth * float64(gasInfo.SafeGasLimit)
	} else {
		rapidLowLimit = rapidEth * float64(gasLimit)
		rapidHighLimit = rapidLowLimit
	}

	fastGwei := math.RoundUp(eth.WeiToGwei(gasSuggestion.FastWei)+priorityFee, 0)
	fastEth := eth.WeiToEth(gasSuggestion.FastWei)

	var fastLowLimit float64
	var fastHighLimit float64
	if gasLimit == 0 {
		fastLowLimit = fastEth * float64(gasInfo.EstGasLimit)
		fastHighLimit = fastEth * float64(gasInfo.SafeGasLimit)
	} else {
		fastLowLimit = fastEth * float64(gasLimit)
		fastHighLimit = fastLowLimit
	}

	standardGwei := math.RoundUp(eth.WeiToGwei(gasSuggestion.StandardWei)+priorityFee, 0)
	standardEth := eth.WeiToEth(gasSuggestion.StandardWei)

	var standardLowLimit float64
	var standardHighLimit float64
	if gasLimit == 0 {
		standardLowLimit = standardEth * float64(gasInfo.EstGasLimit)
		standardHighLimit = standardEth * float64(gasInfo.SafeGasLimit)
	} else {
		standardLowLimit = standardEth * float64(gasLimit)
		standardHighLimit = standardLowLimit
	}

	slowGwei := math.RoundUp(eth.WeiToGwei(gasSuggestion.SlowWei)+priorityFee, 0)
	slowEth := eth.WeiToEth(gasSuggestion.SlowWei)

	var slowLowLimit float64
	var slowHighLimit float64
	if gasLimit == 0 {
		slowLowLimit = slowEth * float64(gasInfo.EstGasLimit)
		slowHighLimit = slowEth * float64(gasInfo.SafeGasLimit)
	} else {
		slowLowLimit = slowEth * float64(gasLimit)
		slowHighLimit = slowLowLimit
	}

	fmt.Printf("%s+============== Suggested Gas Prices ==============+\n", log.ColorBlue)
	fmt.Println("| Avg Wait Time |  Max Fee  |    Total Gas Cost    |")
	fmt.Printf("| %-13s | %-9s | %.4f to %.4f ETH |\n",
		gasSuggestion.RapidTime, fmt.Sprintf("%d gwei", int(rapidGwei)), rapidLowLimit, rapidHighLimit)
	fmt.Printf("| %-13s | %-9s | %.4f to %.4f ETH |\n",
		gasSuggestion.FastTime, fmt.Sprintf("%d gwei", int(fastGwei)), fastLowLimit, fastHighLimit)
	fmt.Printf("| %-13s | %-9s | %.4f to %.4f ETH |\n",
		gasSuggestion.StandardTime, fmt.Sprintf("%d gwei", int(standardGwei)), standardLowLimit, standardHighLimit)
	fmt.Printf("| %-13s | %-9s | %.4f to %.4f ETH |\n",
		gasSuggestion.SlowTime, fmt.Sprintf("%d gwei", int(slowGwei)), slowLowLimit, slowHighLimit)
	fmt.Printf("+==================================================+\n\n%s", log.ColorReset)

	fmt.Printf("These prices include a maximum priority fee of %.2f gwei.\n", priorityFee)

	for {
		desiredPrice := cliutils.Prompt(
			fmt.Sprintf("Please enter your max fee (including the priority fee) or leave blank for the default of %d gwei:", int(fastGwei)),
			"^(?:[1-9]\\d*|0)?(?:\\.\\d+)?$",
			"Not a valid gas price, try again:")

		if desiredPrice == "" {
			return fastGwei
		}

		desiredPriceFloat, err := strconv.ParseFloat(desiredPrice, 64)
		if err != nil {
			fmt.Println("Not a valid gas price (%s), try again.", err.Error())
			continue
		}
		if desiredPriceFloat <= 0 {
			fmt.Println("Max fee must be greater than zero.")
			continue
		}

		return desiredPriceFloat
	}

}

func handleEtherscanGasPrices(gasSuggestion etherscan.GasFeeSuggestion, gasInfo staderCore.GasInfo, priorityFee float64, gasLimit uint64) float64 {

	fastGwei := math.RoundUp(gasSuggestion.FastGwei+priorityFee, 0)
	fastEth := gasSuggestion.FastGwei / eth.WeiPerGwei

	var fastLowLimit float64
	var fastHighLimit float64
	if gasLimit == 0 {
		fastLowLimit = fastEth * float64(gasInfo.EstGasLimit)
		fastHighLimit = fastEth * float64(gasInfo.SafeGasLimit)
	} else {
		fastLowLimit = fastEth * float64(gasLimit)
		fastHighLimit = fastLowLimit
	}

	standardGwei := math.RoundUp(gasSuggestion.StandardGwei+priorityFee, 0)
	standardEth := gasSuggestion.StandardGwei / eth.WeiPerGwei

	var standardLowLimit float64
	var standardHighLimit float64
	if gasLimit == 0 {
		standardLowLimit = standardEth * float64(gasInfo.EstGasLimit)
		standardHighLimit = standardEth * float64(gasInfo.SafeGasLimit)
	} else {
		standardLowLimit = standardEth * float64(gasLimit)
		standardHighLimit = standardLowLimit
	}

	slowGwei := math.RoundUp(gasSuggestion.SlowGwei+priorityFee, 0)
	slowEth := gasSuggestion.SlowGwei / eth.WeiPerGwei

	var slowLowLimit float64
	var slowHighLimit float64
	if gasLimit == 0 {
		slowLowLimit = slowEth * float64(gasInfo.EstGasLimit)
		slowHighLimit = slowEth * float64(gasInfo.SafeGasLimit)
	} else {
		slowLowLimit = slowEth * float64(gasLimit)
		slowHighLimit = slowLowLimit
	}

	fmt.Printf("%s+============ Suggested Gas Prices ============+\n", log.ColorBlue)
	fmt.Println("|   Speed   |  Max Fee  |    Total Gas Cost    |")
	fmt.Printf("| Fast      | %-9s | %.4f to %.4f ETH |\n",
		fmt.Sprintf("%d gwei", int(fastGwei)), fastLowLimit, fastHighLimit)
	fmt.Printf("| Standard  | %-9s | %.4f to %.4f ETH |\n",
		fmt.Sprintf("%d gwei", int(standardGwei)), standardLowLimit, standardHighLimit)
	fmt.Printf("| Slow      | %-9s | %.4f to %.4f ETH |\n",
		fmt.Sprintf("%d gwei", int(slowGwei)), slowLowLimit, slowHighLimit)
	fmt.Printf("+==============================================+\n\n%s", log.ColorReset)

	fmt.Printf("These prices include a maximum priority fee of %.2f gwei.\n", priorityFee)

	for {
		desiredPrice := cliutils.Prompt(
			fmt.Sprintf("Please enter your max fee (including the priority fee) or leave blank for the default of %d gwei:", int(fastGwei)),
			"^(?:[1-9]\\d*|0)?(?:\\.\\d+)?$",
			"Not a valid gas price, try again:")

		if desiredPrice == "" {
			return fastGwei
		}

		desiredPriceFloat, err := strconv.ParseFloat(desiredPrice, 64)
		if err != nil {
			fmt.Println("Not a valid gas price (%s), try again.", err.Error())
			continue
		}
		if desiredPriceFloat <= 0 {
			fmt.Println("Max fee must be greater than zero.")
			continue
		}

		return desiredPriceFloat
	}

}
