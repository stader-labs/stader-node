package node

import (
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/stader-labs/stader-node/stader-lib/stader"
)

func IsPermissionlessNodeRegistryPaused(pnr *stader.PermissionlessNodeRegistryContractManager, opts *bind.CallOpts) (bool, error) {
	return pnr.PermissionlessNodeRegistry.Paused(opts)
}
