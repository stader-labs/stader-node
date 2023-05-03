package types

import "math/big"

type RewardShare struct {
	UserShare     *big.Int
	OperatorShare *big.Int
	ProtocolShare *big.Int
}
