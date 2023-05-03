package types

import "math/big"

type PoolThresholdInfo struct {
	MinThreshold      *big.Int
	MaxThreshold      *big.Int
	WithdrawThreshold *big.Int
	Units             string
}

type OperatorWithdrawInfo struct {
	LastWithdrawReqTimestamp *big.Int
	TotalSDWithdrawReqAmount *big.Int
}
