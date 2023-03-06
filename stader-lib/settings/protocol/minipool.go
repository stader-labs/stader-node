package protocol

import (
	"fmt"
	"math/big"
	"sync"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"

	protocoldao "github.com/stader-labs/stader-node/stader-lib/dao/protocol"
	"github.com/stader-labs/stader-node/stader-lib/stader"
)

// Config
const MinipoolSettingsContractName = "rocketDAOProtocolSettingsMinipool"

// Get the minipool launch balance
func GetMinipoolLaunchBalance(rp *stader.PermissionlessNodeRegistryContractManager, opts *bind.CallOpts) (*big.Int, error) {
	minipoolSettingsContract, err := getMinipoolSettingsContract(rp, opts)
	if err != nil {
		return nil, err
	}
	value := new(*big.Int)
	if err := minipoolSettingsContract.Call(opts, value, "getLaunchBalance"); err != nil {
		return nil, fmt.Errorf("Could not get minipool launch balance: %w", err)
	}
	return *value, nil
}

// Required node deposit amounts
func GetMinipoolFullDepositNodeAmount(rp *stader.PermissionlessNodeRegistryContractManager, opts *bind.CallOpts) (*big.Int, error) {
	minipoolSettingsContract, err := getMinipoolSettingsContract(rp, opts)
	if err != nil {
		return nil, err
	}
	value := new(*big.Int)
	if err := minipoolSettingsContract.Call(opts, value, "getFullDepositNodeAmount"); err != nil {
		return nil, fmt.Errorf("Could not get full minipool deposit node amount: %w", err)
	}
	return *value, nil
}
func GetMinipoolHalfDepositNodeAmount(rp *stader.PermissionlessNodeRegistryContractManager, opts *bind.CallOpts) (*big.Int, error) {
	minipoolSettingsContract, err := getMinipoolSettingsContract(rp, opts)
	if err != nil {
		return nil, err
	}
	value := new(*big.Int)
	if err := minipoolSettingsContract.Call(opts, value, "getHalfDepositNodeAmount"); err != nil {
		return nil, fmt.Errorf("Could not get half minipool deposit node amount: %w", err)
	}
	return *value, nil
}
func GetMinipoolEmptyDepositNodeAmount(rp *stader.PermissionlessNodeRegistryContractManager, opts *bind.CallOpts) (*big.Int, error) {
	minipoolSettingsContract, err := getMinipoolSettingsContract(rp, opts)
	if err != nil {
		return nil, err
	}
	value := new(*big.Int)
	if err := minipoolSettingsContract.Call(opts, value, "getEmptyDepositNodeAmount"); err != nil {
		return nil, fmt.Errorf("Could not get empty minipool deposit node amount: %w", err)
	}
	return *value, nil
}

// Required user deposit amounts
func GetMinipoolFullDepositUserAmount(rp *stader.PermissionlessNodeRegistryContractManager, opts *bind.CallOpts) (*big.Int, error) {
	minipoolSettingsContract, err := getMinipoolSettingsContract(rp, opts)
	if err != nil {
		return nil, err
	}
	value := new(*big.Int)
	if err := minipoolSettingsContract.Call(opts, value, "getFullDepositUserAmount"); err != nil {
		return nil, fmt.Errorf("Could not get full minipool deposit user amount: %w", err)
	}
	return *value, nil
}
func GetMinipoolHalfDepositUserAmount(rp *stader.PermissionlessNodeRegistryContractManager, opts *bind.CallOpts) (*big.Int, error) {
	minipoolSettingsContract, err := getMinipoolSettingsContract(rp, opts)
	if err != nil {
		return nil, err
	}
	value := new(*big.Int)
	if err := minipoolSettingsContract.Call(opts, value, "getHalfDepositUserAmount"); err != nil {
		return nil, fmt.Errorf("Could not get half minipool deposit user amount: %w", err)
	}
	return *value, nil
}
func GetMinipoolEmptyDepositUserAmount(rp *stader.PermissionlessNodeRegistryContractManager, opts *bind.CallOpts) (*big.Int, error) {
	minipoolSettingsContract, err := getMinipoolSettingsContract(rp, opts)
	if err != nil {
		return nil, err
	}
	value := new(*big.Int)
	if err := minipoolSettingsContract.Call(opts, value, "getEmptyDepositUserAmount"); err != nil {
		return nil, fmt.Errorf("Could not get empty minipool deposit user amount: %w", err)
	}
	return *value, nil
}

// Minipool withdrawable event submissions currently enabled
func GetMinipoolSubmitWithdrawableEnabled(rp *stader.PermissionlessNodeRegistryContractManager, opts *bind.CallOpts) (bool, error) {
	minipoolSettingsContract, err := getMinipoolSettingsContract(rp, opts)
	if err != nil {
		return false, err
	}
	value := new(bool)
	if err := minipoolSettingsContract.Call(opts, value, "getSubmitWithdrawableEnabled"); err != nil {
		return false, fmt.Errorf("Could not get minipool withdrawable submissions enabled status: %w", err)
	}
	return *value, nil
}
func BootstrapMinipoolSubmitWithdrawableEnabled(rp *stader.PermissionlessNodeRegistryContractManager, value bool, opts *bind.TransactOpts) (common.Hash, error) {
	return protocoldao.BootstrapBool(rp, MinipoolSettingsContractName, "minipool.submit.withdrawable.enabled", value, opts)
}

// Timeout period in seconds for prelaunch minipools to launch
func GetMinipoolLaunchTimeout(rp *stader.PermissionlessNodeRegistryContractManager, opts *bind.CallOpts) (time.Duration, error) {
	minipoolSettingsContract, err := getMinipoolSettingsContract(rp, opts)
	if err != nil {
		return 0, err
	}
	value := new(*big.Int)
	if err := minipoolSettingsContract.Call(opts, value, "getLaunchTimeout"); err != nil {
		return 0, fmt.Errorf("Could not get minipool launch timeout: %w", err)
	}
	seconds := time.Duration((*value).Int64()) * time.Second
	return seconds, nil
}
func BootstrapMinipoolLaunchTimeout(rp *stader.PermissionlessNodeRegistryContractManager, value time.Duration, opts *bind.TransactOpts) (common.Hash, error) {
	return protocoldao.BootstrapUint(rp, MinipoolSettingsContractName, "minipool.launch.timeout", big.NewInt(int64(value.Seconds())), opts)
}

// Get contracts
var minipoolSettingsContractLock sync.Mutex

func getMinipoolSettingsContract(rp *stader.PermissionlessNodeRegistryContractManager, opts *bind.CallOpts) (*stader.Contract, error) {
	minipoolSettingsContractLock.Lock()
	defer minipoolSettingsContractLock.Unlock()
	return nil, nil
	//return rp.GetContract(MinipoolSettingsContractName, opts)
}
