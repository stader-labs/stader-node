package network

import (
	"fmt"
	"math/big"
	"sync"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"

	"github.com/stader-labs/stader-node/stader-lib/stader"
	"github.com/stader-labs/stader-node/stader-lib/utils/eth"
)

// Get the current network node demand in ETH
func GetNodeDemand(rp *stader.PermissionlessNodeRegistryContractManager, opts *bind.CallOpts) (*big.Int, error) {
	rocketNetworkFees, err := getRocketNetworkFees(rp, opts)
	if err != nil {
		return nil, err
	}
	nodeDemand := new(*big.Int)
	if err := rocketNetworkFees.Call(opts, nodeDemand, "getNodeDemand"); err != nil {
		return nil, fmt.Errorf("Could not get network node demand: %w", err)
	}
	return *nodeDemand, nil
}

// Get the current network node commission rate
func GetNodeFee(rp *stader.PermissionlessNodeRegistryContractManager, opts *bind.CallOpts) (float64, error) {
	rocketNetworkFees, err := getRocketNetworkFees(rp, opts)
	if err != nil {
		return 0, err
	}
	nodeFee := new(*big.Int)
	if err := rocketNetworkFees.Call(opts, nodeFee, "getNodeFee"); err != nil {
		return 0, fmt.Errorf("Could not get network node fee: %w", err)
	}
	return eth.WeiToEth(*nodeFee), nil
}

// Get the network node fee for a node demand value
func GetNodeFeeByDemand(rp *stader.PermissionlessNodeRegistryContractManager, nodeDemand *big.Int, opts *bind.CallOpts) (float64, error) {
	rocketNetworkFees, err := getRocketNetworkFees(rp, opts)
	if err != nil {
		return 0, err
	}
	nodeFee := new(*big.Int)
	if err := rocketNetworkFees.Call(opts, nodeFee, "getNodeFeeByDemand", nodeDemand); err != nil {
		return 0, fmt.Errorf("Could not get node fee by node demand: %w", err)
	}
	return eth.WeiToEth(*nodeFee), nil
}

// Get contracts
var rocketNetworkFeesLock sync.Mutex

func getRocketNetworkFees(rp *stader.PermissionlessNodeRegistryContractManager, opts *bind.CallOpts) (*stader.Contract, error) {
	rocketNetworkFeesLock.Lock()
	defer rocketNetworkFeesLock.Unlock()
	return nil, nil
	//return rp.GetContract("rocketNetworkFees", opts)
}
