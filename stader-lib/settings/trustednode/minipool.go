package trustednode

import (
	"fmt"
	"math/big"
	"sync"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"

	trustednodedao "github.com/stader-labs/stader-node/stader-lib/dao/trustednode"
	"github.com/stader-labs/stader-node/stader-lib/stader"
)

// Config
const (
	MinipoolSettingsContractName = "rocketDAONodeTrustedSettingsMinipool"
	ScrubPeriodPath              = "minipool.scrub.period"
	ScrubPenaltyEnabledPath      = "minipool.scrub.penalty.enabled"
)

// The cooldown period a member must wait after making a proposal before making another in seconds
func GetScrubPeriod(rp *stader.PermissionlessNodeRegistryContractManager, opts *bind.CallOpts) (uint64, error) {
	minipoolSettingsContract, err := getMinipoolSettingsContract(rp, opts)
	if err != nil {
		return 0, err
	}
	value := new(*big.Int)
	if err := minipoolSettingsContract.Call(opts, value, "getScrubPeriod"); err != nil {
		return 0, fmt.Errorf("Could not get scrub period: %w", err)
	}
	return (*value).Uint64(), nil
}
func BootstrapScrubPeriod(rp *stader.PermissionlessNodeRegistryContractManager, value uint64, opts *bind.TransactOpts) (common.Hash, error) {
	return trustednodedao.BootstrapUint(rp, MinipoolSettingsContractName, ScrubPeriodPath, big.NewInt(int64(value)), opts)
}
func ProposeScrubPeriod(rp *stader.PermissionlessNodeRegistryContractManager, value uint64, opts *bind.TransactOpts) (uint64, common.Hash, error) {
	return trustednodedao.ProposeSetUint(rp, fmt.Sprintf("set %s", ScrubPeriodPath), MinipoolSettingsContractName, ScrubPeriodPath, big.NewInt(int64(value)), opts)
}
func EstimateProposeScrubPeriodGas(rp *stader.PermissionlessNodeRegistryContractManager, value uint64, opts *bind.TransactOpts) (stader.GasInfo, error) {
	return trustednodedao.EstimateProposeSetUintGas(rp, fmt.Sprintf("set %s", ScrubPeriodPath), MinipoolSettingsContractName, ScrubPeriodPath, big.NewInt(int64(value)), opts)
}

// Whether or not the RPL slashing penalty is applied to scrubbed minipools
func GetScrubPenaltyEnabled(rp *stader.PermissionlessNodeRegistryContractManager, opts *bind.CallOpts) (bool, error) {
	minipoolSettingsContract, err := getMinipoolSettingsContract(rp, opts)
	if err != nil {
		return false, err
	}
	value := new(bool)
	if err := minipoolSettingsContract.Call(opts, value, "getScrubPenaltyEnabled"); err != nil {
		return false, fmt.Errorf("Could not get scrub penalty setting: %w", err)
	}
	return (*value), nil
}
func BootstrapScrubPenaltyEnabled(rp *stader.PermissionlessNodeRegistryContractManager, value bool, opts *bind.TransactOpts) (common.Hash, error) {
	return trustednodedao.BootstrapBool(rp, MinipoolSettingsContractName, ScrubPenaltyEnabledPath, value, opts)
}
func ProposeScrubPenaltyEnabled(rp *stader.PermissionlessNodeRegistryContractManager, value bool, opts *bind.TransactOpts) (uint64, common.Hash, error) {
	return trustednodedao.ProposeSetBool(rp, fmt.Sprintf("set %s", ScrubPenaltyEnabledPath), MinipoolSettingsContractName, ScrubPenaltyEnabledPath, value, opts)
}
func EstimateProposeScrubPenaltyEnabledGas(rp *stader.PermissionlessNodeRegistryContractManager, value bool, opts *bind.TransactOpts) (stader.GasInfo, error) {
	return trustednodedao.EstimateProposeSetBoolGas(rp, fmt.Sprintf("set %s", ScrubPenaltyEnabledPath), MinipoolSettingsContractName, ScrubPenaltyEnabledPath, value, opts)
}

// Get contracts
var minipoolSettingsContractLock sync.Mutex

func getMinipoolSettingsContract(rp *stader.PermissionlessNodeRegistryContractManager, opts *bind.CallOpts) (*stader.Contract, error) {
	minipoolSettingsContractLock.Lock()
	defer minipoolSettingsContractLock.Unlock()
	return nil, nil
	//return rp.GetContract(MinipoolSettingsContractName, opts)
}
