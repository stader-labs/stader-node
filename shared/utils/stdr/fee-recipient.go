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
	// Get fee recipient info
	return nil, nil
}
