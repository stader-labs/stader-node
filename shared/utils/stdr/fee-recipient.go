package stdr

import (
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/stader-labs/stader-node/stader-lib/node"
	"github.com/stader-labs/stader-node/stader-lib/stader"
)

type FeeRecipientInfo struct {
	SocializingPoolAddress common.Address `json:"socializingPoolAddress"`
	FeeDistributorAddress  common.Address `json:"feeDistributorAddress"`
	IsInSocializingPool    bool           `json:"isInSocializingPool"`
}

func GetFeeRecipientInfo(prn *stader.PermissionlessNodeRegistryContractManager, vf *stader.VaultFactoryContractManager, pp *stader.PermissionlessPoolContractManager, nodeAddress common.Address, opts *bind.CallOpts) (*FeeRecipientInfo, error) {
	feeRecipientInfo := FeeRecipientInfo{
		SocializingPoolAddress: common.Address{},
		FeeDistributorAddress:  common.Address{},
		IsInSocializingPool:    false,
	}

	operatorId, err := node.GetOperatorId(prn, nodeAddress, opts)
	if err != nil {
		return nil, err
	}
	operatorInfo, err := node.GetOperatorInfo(prn, operatorId, opts)
	if err != nil {
		return nil, err
	}

	if operatorInfo.OptedForSocializingPool {
		feeRecipientInfo.IsInSocializingPool = true
		socializingPoolAddress, err := node.GetSocializingPoolContract(pp, opts)
		if err != nil {
			return nil, err
		}
		feeRecipientInfo.SocializingPoolAddress = socializingPoolAddress
	} else {
		nodeElRewardAddress, err := node.GetNodeElRewardAddress(vf, 1, operatorId, opts)
		if err != nil {
			return nil, err
		}
		feeRecipientInfo.FeeDistributorAddress = nodeElRewardAddress
	}

	return &feeRecipientInfo, nil
}
