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
package etherchain

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math/big"
	"net/http"
)

const gasNowUrl string = "https://beaconcha.in/api/v1/execution/gasnow"

// Standard response
type gasNowResponse struct {
	Data struct {
		Rapid    uint64  `json:"rapid"`
		Fast     uint64  `json:"fast"`
		Standard uint64  `json:"standard"`
		Slow     uint64  `json:"slow"`
		PriceUSD float64 `json:"priceUSD"`
	} `json:"data"`
}

type GasFeeSuggestion struct {
	RapidWei  *big.Int
	RapidTime string

	FastWei  *big.Int
	FastTime string

	StandardWei  *big.Int
	StandardTime string

	SlowWei  *big.Int
	SlowTime string

	EthUsd float64
}

// Get gas prices
func GetGasPrices() (GasFeeSuggestion, error) {

	// Send request
	response, err := http.Get(gasNowUrl)
	if err != nil {
		return GasFeeSuggestion{}, err
	}
	defer func() {
		_ = response.Body.Close()
	}()

	// Check the response code
	if response.StatusCode != http.StatusOK {
		return GasFeeSuggestion{}, fmt.Errorf("request failed with code %d", response.StatusCode)
	}

	// Get response
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return GasFeeSuggestion{}, err
	}

	// Deserialize response
	var gnResponse gasNowResponse
	if err := json.Unmarshal(body, &gnResponse); err != nil {
		return GasFeeSuggestion{}, fmt.Errorf("Could not decode Gas Now response: %w", err)
	}

	suggestion := GasFeeSuggestion{
		RapidWei:  big.NewInt(0).SetUint64(gnResponse.Data.Rapid),
		RapidTime: "15 Seconds",

		FastWei:  big.NewInt(0).SetUint64(gnResponse.Data.Fast),
		FastTime: "1 Minute",

		StandardWei:  big.NewInt(0).SetUint64(gnResponse.Data.Standard),
		StandardTime: "3 Minutes",

		SlowWei:  big.NewInt(0).SetUint64(gnResponse.Data.Slow),
		SlowTime: ">10 Minutes",

		EthUsd: gnResponse.Data.PriceUSD,
	}

	// Return
	return suggestion, nil

}
