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
package utils

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/stader-labs/stader-node/stader-lib/stader"
)

// Wait for a transaction to get mined
func WaitForTransaction(client stader.ExecutionClient, hash common.Hash) (*types.Receipt, error) {

	var tx *types.Transaction
	var err error

	// Get the transaction from its hash, retrying for 30 sec if it wasn't found
	for i := 0; i < 30; i++ {
		if i == 29 {
			return nil, fmt.Errorf("Transaction not found after 30 seconds.")
		}

		tx, _, err = client.TransactionByHash(context.Background(), hash)
		if err != nil {
			if err.Error() == "not found" {
				time.Sleep(1 * time.Second)
				continue
			}
			return nil, err
		} else {
			break
		}
	}

	// Wait for transaction to be mined
	txReceipt, err := bind.WaitMined(context.Background(), client, tx)
	if err != nil {
		return nil, err
	}

	// Check transaction status
	if txReceipt.Status == 0 {
		return txReceipt, errors.New("Transaction failed with status 0")
	}

	// Return
	return txReceipt, nil
}
