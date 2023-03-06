package minipool

import (
	"fmt"
	"sync"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"

	"github.com/stader-labs/stader-node/stader-lib/stader"
)

// Get the CreationCode binary for the RocketMinipool contract that will be created by node deposits
func GetMinipoolBytecode(rp *stader.PermissionlessNodeRegistryContractManager, opts *bind.CallOpts) ([]byte, error) {
	rocketMinipoolFactory, err := getRocketMinipoolFactory(rp, opts)
	if err != nil {
		return []byte{}, err
	}
	bytecode := new([]byte)
	if err := rocketMinipoolFactory.Call(opts, bytecode, "getMinipoolBytecode"); err != nil {
		return []byte{}, fmt.Errorf("Could not get minipool contract bytecode: %w", err)
	}
	return *bytecode, nil
}

// Get contracts
var rocketMinipoolFactoryLock sync.Mutex

func getRocketMinipoolFactory(rp *stader.PermissionlessNodeRegistryContractManager, opts *bind.CallOpts) (*stader.Contract, error) {
	rocketMinipoolFactoryLock.Lock()
	defer rocketMinipoolFactoryLock.Unlock()
	return nil, nil
	//return rp.GetContract("rocketMinipoolFactory", opts)
}
