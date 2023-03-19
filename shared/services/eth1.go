package services

// ROCKETPOOL-OWNED

import (
	"context"
	"github.com/stader-labs/stader-node/stader-lib/stader"
)

// ROCKETPOOL-OWNED

func GetEthClientLatestBlockTimestamp(ec stader.ExecutionClient) (uint64, error) {
	// Get latest block
	header, err := ec.HeaderByNumber(context.Background(), nil)
	if err != nil {
		return 0, err
	}

	// Return block timestamp
	return header.Time, nil
}
