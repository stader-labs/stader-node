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
package stader

import (
	"context"
	"math/big"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
)

// This is the common interface for execution clients.
type ExecutionClient interface {

	/// ========================
	/// ContractCaller Functions
	/// ========================

	// CodeAt returns the code of the given account. This is needed to differentiate
	// between contract internal errors and the local chain being out of sync.
	CodeAt(ctx context.Context, contract common.Address, blockNumber *big.Int) ([]byte, error)

	// CallContract executes an Ethereum contract call with the specified data as the
	// input.
	CallContract(ctx context.Context, call ethereum.CallMsg, blockNumber *big.Int) ([]byte, error)

	/// ============================
	/// ContractTransactor Functions
	/// ============================

	// HeaderByHash returns the block header with the given hash.
	HeaderByHash(ctx context.Context, hash common.Hash) (*types.Header, error)

	// HeaderByNumber returns a block header from the current canonical chain. If number is
	// nil, the latest known header is returned.
	HeaderByNumber(ctx context.Context, number *big.Int) (*types.Header, error)

	// PendingCodeAt returns the code of the given account in the pending state.
	PendingCodeAt(ctx context.Context, account common.Address) ([]byte, error)

	// PendingNonceAt retrieves the current pending nonce associated with an account.
	PendingNonceAt(ctx context.Context, account common.Address) (uint64, error)

	// SuggestGasPrice retrieves the currently suggested gas price to allow a timely
	// execution of a transaction.
	SuggestGasPrice(ctx context.Context) (*big.Int, error)

	// SuggestGasTipCap retrieves the currently suggested 1559 priority fee to allow
	// a timely execution of a transaction.
	SuggestGasTipCap(ctx context.Context) (*big.Int, error)

	// EstimateGas tries to estimate the gas needed to execute a specific
	// transaction based on the current pending state of the backend blockchain.
	// There is no guarantee that this is the true gas limit requirement as other
	// transactions may be added or removed by miners, but it should provide a basis
	// for setting a reasonable default.
	EstimateGas(ctx context.Context, call ethereum.CallMsg) (gas uint64, err error)

	// SendTransaction injects the transaction into the pending pool for execution.
	SendTransaction(ctx context.Context, tx *types.Transaction) error

	/// ==========================
	/// ContractFilterer Functions
	/// ==========================

	// FilterLogs executes a log filter operation, blocking during execution and
	// returning all the results in one batch.
	//
	FilterLogs(ctx context.Context, query ethereum.FilterQuery) ([]types.Log, error)

	// SubscribeFilterLogs creates a background log filtering operation, returning
	// a subscription immediately, which can be used to stream the found events.
	SubscribeFilterLogs(ctx context.Context, query ethereum.FilterQuery, ch chan<- types.Log) (ethereum.Subscription, error)

	/// =======================
	/// DeployBackend Functions
	/// =======================

	// TransactionReceipt returns the receipt of a transaction by transaction hash.
	// Note that the receipt is not available for pending transactions.
	TransactionReceipt(ctx context.Context, txHash common.Hash) (*types.Receipt, error)

	/// ================
	/// Client functions
	/// ================

	// BlockNumber returns the most recent block number
	BlockNumber(ctx context.Context) (uint64, error)

	// BalanceAt returns the wei balance of the given account.
	// The block number can be nil, in which case the balance is taken from the latest known block.
	BalanceAt(ctx context.Context, account common.Address, blockNumber *big.Int) (*big.Int, error)

	// TransactionByHash returns the transaction with the given hash.
	TransactionByHash(ctx context.Context, hash common.Hash) (tx *types.Transaction, isPending bool, err error)

	// NonceAt returns the account nonce of the given account.
	// The block number can be nil, in which case the nonce is taken from the latest known block.
	NonceAt(ctx context.Context, account common.Address, blockNumber *big.Int) (uint64, error)

	// SyncProgress retrieves the current progress of the sync algorithm. If there's
	// no sync currently running, it returns nil.
	SyncProgress(ctx context.Context) (*ethereum.SyncProgress, error)
}
