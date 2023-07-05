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
	"context"
	"math/big"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/stader-labs/stader-node/stader-lib/stader"
)

// Estimate the gas of SendTransaction
func EstimateSendTransactionGas(client stader.ExecutionClient, toAddress common.Address, opts *bind.TransactOpts) (stader.GasInfo, error) {

	// User-defined settings
	response := stader.GasInfo{}

	// Set default value
	value := opts.Value
	if value == nil {
		value = big.NewInt(0)
	}

	// Estimate gas limit
	gasLimit, err := client.EstimateGas(context.Background(), ethereum.CallMsg{
		From:     opts.From,
		To:       &toAddress,
		GasPrice: big.NewInt(0), // set to 0 for simulation
		Value:    value,
	})
	if err != nil {
		return stader.GasInfo{}, err
	}
	response.EstGasLimit = gasLimit
	response.SafeGasLimit = gasLimit

	return response, err
}

// Send a transaction to an address
func SendTransaction(client stader.ExecutionClient, toAddress common.Address, chainID *big.Int, opts *bind.TransactOpts) (common.Hash, error) {
	var err error

	// Get from address nonce
	var nonce uint64
	if opts.Nonce == nil {
		nonce, err = client.PendingNonceAt(context.Background(), opts.From)
		if err != nil {
			return common.Hash{}, err
		}
	} else {
		nonce = opts.Nonce.Uint64()
	}

	// Set default value
	value := opts.Value
	if value == nil {
		value = big.NewInt(0)
	}

	// Estimate gas limit
	gasLimit := opts.GasLimit
	if gasLimit == 0 {
		gasLimit, err = client.EstimateGas(context.Background(), ethereum.CallMsg{
			From:     opts.From,
			To:       &toAddress,
			GasPrice: big.NewInt(0), // use 0 gwei for simulation
			Value:    value,
		})
		if err != nil {
			return common.Hash{}, err
		}
	}

	// Initialize transaction
	tx := types.NewTx(&types.DynamicFeeTx{
		ChainID:    chainID,
		Nonce:      nonce,
		GasTipCap:  opts.GasTipCap,
		GasFeeCap:  opts.GasFeeCap,
		Gas:        gasLimit,
		To:         &toAddress,
		Value:      value,
		Data:       []byte{},
		AccessList: []types.AccessTuple{},
	})

	// Sign transaction
	signedTx, err := opts.Signer(opts.From, tx)
	if err != nil {
		return common.Hash{}, err
	}

	// Send transaction
	if err = client.SendTransaction(context.Background(), signedTx); err != nil {
		return common.Hash{}, err
	}

	return signedTx.Hash(), nil

}
