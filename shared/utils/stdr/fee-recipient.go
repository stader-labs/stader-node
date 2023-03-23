package stdr

// ROCKETPOOL-OWNED

import (
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/stader-labs/stader-node/shared/services/beacon"
	"github.com/stader-labs/stader-node/stader-lib/stader"
)

type FeeRecipientInfo struct {
	SocializingPoolAddress common.Address `json:"socializingPoolAddress"`
	FeeDistributorAddress  common.Address `json:"feeDistributorAddress"`
	IsInSocializingPool    bool           `json:"isInSocializingPool"`
	IsInOptOutCooldown     bool           `json:"isInOptOutCooldown"`
	OptOutEpoch            uint64         `json:"optOutEpoch"`
}

// TODO - add fee receipient info for socializing pools
func GetFeeRecipientInfo(prn *stader.PermissionlessNodeRegistryContractManager, bc beacon.Client, nodeAddress common.Address, opts *bind.CallOpts) (*FeeRecipientInfo, error) {
	// TODO - Get fee recipient info from sanjay
	return &FeeRecipientInfo{
		SocializingPoolAddress: common.BytesToAddress([]byte("0x1Dd8f12e17519732b56bdcFaea9D10FB55E2cb6C")),
		FeeDistributorAddress:  common.Address{},
		IsInSocializingPool:    true,
		IsInOptOutCooldown:     false,
		OptOutEpoch:            0,
	}, nil
}
