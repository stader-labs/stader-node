/*
This work is licensed and released under GNU GPL v3 or any other later versions.
The full text of the license is below/ found at <http://www.gnu.org/licenses/>

(c) 2023 Rocket Pool Pty Ltd. Modified under GNU GPL v3. [1.2.1]

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
package etherscan

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
)

const gasOracleUrl string = "https://api.etherscan.io/api?module=gastracker&action=gasoracle"

// Standard response
type gasOracleResponse struct {
	Status  uinteger `json:"status"`
	Message string   `json:"message"`
	Result  struct {
		SafeGasPrice    uinteger `json:"SafeGasPrice"`
		ProposeGasPrice uinteger `json:"ProposeGasPrice"`
		FastGasPrice    uinteger `json:"FastGasPrice"`
	} `json:"result"`
}

type GasFeeSuggestion struct {
	SlowGwei     float64
	StandardGwei float64
	FastGwei     float64
}

// Get gas prices
func GetGasPrices() (GasFeeSuggestion, error) {

	// Send request
	response, err := http.Get(gasOracleUrl)
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
	var gOResponse gasOracleResponse
	if err := json.Unmarshal(body, &gOResponse); err != nil {
		return GasFeeSuggestion{}, fmt.Errorf("Could not decode Etherscan gas oracle response: %w", err)
	}
	if gOResponse.Status != 1 {
		return GasFeeSuggestion{}, fmt.Errorf("Error retrieving Etherscan gas oracle response: %s", gOResponse.Message)
	}

	suggestion := GasFeeSuggestion{
		SlowGwei:     float64(gOResponse.Result.SafeGasPrice),
		StandardGwei: float64(gOResponse.Result.ProposeGasPrice),
		FastGwei:     float64(gOResponse.Result.FastGasPrice),
	}

	// Return
	return suggestion, nil

}

// Unsigned integer type
type uinteger uint64

func (i uinteger) MarshalJSON() ([]byte, error) {
	return json.Marshal(strconv.Itoa(int(i)))
}

func (i *uinteger) UnmarshalJSON(data []byte) error {

	// Unmarshal string
	var dataStr string
	if err := json.Unmarshal(data, &dataStr); err != nil {
		return err
	}

	// Parse integer value
	value, err := strconv.ParseUint(dataStr, 10, 64)
	if err != nil {
		return err
	}

	// Set value and return
	*i = uinteger(value)
	return nil
}
