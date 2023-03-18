package stdr

// ROCKETPOOL-OWNED

import (
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/stader-labs/stader-node/shared/services/beacon"
	"github.com/stader-labs/stader-node/stader-lib/stader"
)

type FeeRecipientInfo struct {
	SmoothingPoolAddress  common.Address `json:"smoothingPoolAddress"`
	FeeDistributorAddress common.Address `json:"feeDistributorAddress"`
	IsInSmoothingPool     bool           `json:"isInSmoothingPool"`
	IsInOptOutCooldown    bool           `json:"isInOptOutCooldown"`
	OptOutEpoch           uint64         `json:"optOutEpoch"`
}

func GetFeeRecipientInfo(prn *stader.PermissionlessNodeRegistryContractManager, bc beacon.Client, nodeAddress common.Address, opts *bind.CallOpts) (*FeeRecipientInfo, error) {
	//
	//info := &FeeRecipientInfo{
	//	IsInOptOutCooldown: false,
	//	OptOutEpoch:        0,
	//}
	//
	//// Sync
	//var wg errgroup.Group
	//
	//// Get the smoothing pool address
	//wg.Go(func() error {
	//	smoothingPoolContract, err := stdr.GetContract("rocketSmoothingPool", opts)
	//	if err != nil {
	//		return fmt.Errorf("Error getting smoothing pool contract: %w", err)
	//	}
	//	info.SmoothingPoolAddress = *smoothingPoolContract.Address
	//	return nil
	//})
	//
	//// Get the node's fee distributor
	//wg.Go(func() error {
	//	distributorAddress, err := node.GetDistributorAddress(stdr, nodeAddress, nil)
	//	if err != nil {
	//		return fmt.Errorf("Error getting the fee distributor for %s: %w", nodeAddress.Hex(), err)
	//	}
	//	info.FeeDistributorAddress = distributorAddress
	//	return nil
	//})
	//
	//// Check if the user's opted into the smoothing pool
	//wg.Go(func() error {
	//	isOptedIn, err := node.GetSmoothingPoolRegistrationState(stdr, nodeAddress, nil)
	//	if err != nil {
	//		return err
	//	}
	//	info.IsInSmoothingPool = isOptedIn
	//	return nil
	//})
	//
	//// Wait for data
	//if err := wg.Wait(); err != nil {
	//	return nil, err
	//}
	//
	//// Calculate the safe opt-out epoch if applicable
	//if !info.IsInSmoothingPool {
	//	// Get the opt out time
	//	optOutTime, err := node.GetSmoothingPoolRegistrationChanged(stdr, nodeAddress, nil)
	//	if err != nil {
	//		return nil, fmt.Errorf("Error getting smoothing pool opt-out time: %w", err)
	//	}
	//
	//	// Get the Beacon info
	//	beaconConfig, err := bc.GetEth2Config()
	//	if err != nil {
	//		return nil, fmt.Errorf("Error getting Beacon config: %w", err)
	//	}
	//	beaconHead, err := bc.GetBeaconHead()
	//	if err != nil {
	//		return nil, fmt.Errorf("Error getting Beacon head: %w", err)
	//	}
	//
	//	// Check if the user just opted out
	//	if optOutTime != time.Unix(0, 0) {
	//		// Get the epoch for that time
	//		genesisTime := time.Unix(int64(beaconConfig.GenesisTime), 0)
	//		secondsSinceGenesis := optOutTime.Sub(genesisTime)
	//		epoch := uint64(secondsSinceGenesis.Seconds()) / beaconConfig.SecondsPerEpoch
	//
	//		// Make sure epoch + 1 is finalized - if not, they're still on cooldown
	//		targetEpoch := epoch + 1
	//		if beaconHead.FinalizedEpoch < targetEpoch {
	//			info.IsInOptOutCooldown = true
	//			info.OptOutEpoch = targetEpoch
	//		}
	//	}
	//}
	//
	//return info, nil

	return nil, nil
}
