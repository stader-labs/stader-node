/*
This work is licensed and released under GNU GPL v3 or any other later versions.
The full text of the license is below/ found at <http://www.gnu.org/licenses/>

(c) 2023 Rocket Pool Pty Ltd. Modified under GNU GPL v3. [1.2.0]

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
package eth

import (
	"math/big"
	"strconv"
)

// Conversion factors
const (
	WeiPerEth  float64 = 1e18
	WeiPerGwei float64 = 1e9
)

// Convert wei to eth
func WeiToEth(wei *big.Int) float64 {
	var weiFloat big.Float
	var eth big.Float
	weiFloat.SetInt(wei)
	eth.Quo(&weiFloat, big.NewFloat(WeiPerEth))
	eth64, _ := eth.Float64()
	return eth64
}

// Convert eth to wei
func EthToWei(eth float64) *big.Int {
	var ethFloat big.Float
	var weiFloat big.Float
	var wei big.Int
	ethFloat.SetString(strconv.FormatFloat(eth, 'f', -1, 64))
	weiFloat.Mul(&ethFloat, big.NewFloat(WeiPerEth))
	weiFloat.Int(&wei)
	return &wei
}

// Convert wei to gigawei
func WeiToGwei(wei *big.Int) float64 {
	var weiFloat big.Float
	var gwei big.Float
	weiFloat.SetInt(wei)
	gwei.Quo(&weiFloat, big.NewFloat(WeiPerGwei))
	gwei64, _ := gwei.Float64()
	return gwei64
}

// Convert gigawei to wei
func GweiToWei(gwei float64) *big.Int {
	var gweiFloat big.Float
	var weiFloat big.Float
	var wei big.Int
	gweiFloat.SetString(strconv.FormatFloat(gwei, 'f', -1, 64))
	weiFloat.Mul(&gweiFloat, big.NewFloat(WeiPerGwei))
	weiFloat.Int(&wei)
	return &wei
}
