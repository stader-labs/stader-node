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
package tokens

import (
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/stader-labs/stader-node/stader-lib/stader"
)

// Token balances
type Balances struct {
	ETH  *big.Int `json:"eth"`
	Sd   *big.Int `json:"sd"`
	Ethx *big.Int `json:"ethx"`
}

// Get a token's total supply
func TotalSupply(tokenContract *stader.Erc20TokenContractManager, opts *bind.CallOpts) (*big.Int, error) {
	return tokenContract.Erc20Token.TotalSupply(opts)
}

// Get a token balance
func BalanceOf(tokenContract *stader.Erc20TokenContractManager, address common.Address, opts *bind.CallOpts) (*big.Int, error) {
	return tokenContract.Erc20Token.BalanceOf(opts, address)
}

// Get a spender's Allowance for an address
func Allowance(tokenContract *stader.Erc20TokenContractManager, owner common.Address, spender common.Address, opts *bind.CallOpts) (*big.Int, error) {
	return tokenContract.Erc20Token.Allowance(opts, owner, spender)
}

// Estimate the gas of Transfer
func EstimateTransferGas(tokenContract *stader.Erc20TokenContractManager, to common.Address, amount *big.Int, opts *bind.TransactOpts) (stader.GasInfo, error) {
	return tokenContract.Erc20TokenContract.GetTransactionGasInfo(opts, "transfer", to, amount)
}

// Transfer tokens to an address
func Transfer(tokenContract *stader.Erc20TokenContractManager, to common.Address, amount *big.Int, opts *bind.TransactOpts) (common.Hash, error) {
	tx, err := tokenContract.Erc20Token.Transfer(opts, to, amount)
	if err != nil {
		return common.Hash{}, err
	}

	return tx.Hash(), nil
}

// Estimate the gas of Approve
func EstimateApproveGas(tokenContract *stader.Erc20TokenContractManager, spender common.Address, amount *big.Int, opts *bind.TransactOpts) (stader.GasInfo, error) {
	return tokenContract.Erc20TokenContract.GetTransactionGasInfo(opts, "approve", spender, amount)
}

// Approve a token Allowance for a spender
func Approve(tokenContract *stader.Erc20TokenContractManager, spender common.Address, amount *big.Int, opts *bind.TransactOpts) (common.Hash, error) {
	tx, err := tokenContract.Erc20Token.Approve(opts, spender, amount)
	if err != nil {
		return common.Hash{}, err
	}

	return tx.Hash(), nil
}

// Estimate the gas of TransferFrom
func EstimateTransferFromGas(tokenContract *stader.Erc20TokenContractManager, from, to common.Address, amount *big.Int, opts *bind.TransactOpts) (stader.GasInfo, error) {
	return tokenContract.Erc20TokenContract.GetTransactionGasInfo(opts, "transferFrom", from, to, amount)
}

// Transfer tokens from a sender to an address
func TransferFrom(tokenContract *stader.Erc20TokenContractManager, from, to common.Address, amount *big.Int, opts *bind.TransactOpts) (common.Hash, error) {
	tx, err := tokenContract.Erc20Token.TransferFrom(opts, from, to, amount)
	if err != nil {
		return common.Hash{}, err
	}

	return tx.Hash(), nil
}
