package types

import "math/big"

type RewardCycleDetails struct {
	CurrentIndex      *big.Int
	CurrentStartBlock *big.Int
	CurrentEndBlock   *big.Int
}

type CurrentRewardCycleDetails struct {
	StartBlock *big.Int
	EndBlock   *big.Int
}
