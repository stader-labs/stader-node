package tokens

import (
	"context"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/stader-labs/stader-node/stader-lib/stader"
	"math/big"
)

func GetEthBalance(client stader.ExecutionClient, address common.Address, opts *bind.CallOpts) (*big.Int, error) {
	var blockNumber *big.Int
	if opts != nil {
		blockNumber = opts.BlockNumber
	}

	ethBalance, err := client.BalanceAt(context.Background(), address, blockNumber)
	if err != nil {
		return nil, err
	}

	return ethBalance, nil
}
