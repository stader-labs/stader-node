package types

import (
	"github.com/ethereum/go-ethereum/common"
	"math/big"
)

type ValidatorContractInfo struct {
	Status               uint8
	Pubkey               []byte
	PreDepositSignature  []byte
	DepositSignature     []byte
	WithdrawVaultAddress common.Address
	OperatorId           *big.Int
	DepositBlock         *big.Int
	WithdrawnBlock       *big.Int
}
