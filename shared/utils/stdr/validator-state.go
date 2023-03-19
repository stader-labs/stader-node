package stdr

import (
	"github.com/ethereum/go-ethereum/common"
	"math/big"
)

var ValidatorState = map[uint8]string{
	0: "Initialized",
	1: "Invalid Signature",
	2: "Front Run",
	3: "Pre Deposit",
	4: "Deposited",
	5: "In Activation Queue",
	6: "Active",
	7: "In Exit Queue",
	8: "Exited",
	9: "Withdrawn",
}

type ValidatorInfo struct {
	Status               uint8
	Pubkey               []byte
	Signature            []byte
	WithdrawVaultAddress common.Address
	OperatorId           *big.Int
	InitialBondEth       *big.Int
	DepositTime          *big.Int
	WithdrawnTime        *big.Int
}