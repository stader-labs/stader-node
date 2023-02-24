package node

import (
	"fmt"
	"math/big"
	"sync"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"

	"github.com/stader-labs/stader-node/stader-lib/stader"
	rptypes "github.com/stader-labs/stader-node/stader-lib/types"
	"github.com/stader-labs/stader-node/stader-lib/utils/eth"
)

// Estimate the gas of Deposit
func EstimateDepositGas(rp *stader.PermissionlessNodeRegistryContractManager, minimumNodeFee float64, validatorPubkey rptypes.ValidatorPubkey, validatorSignature rptypes.ValidatorSignature, depositDataRoot common.Hash, salt *big.Int, expectedMinipoolAddress common.Address, opts *bind.TransactOpts) (stader.GasInfo, error) {
	rocketNodeDeposit, err := getRocketNodeDeposit(rp, nil)
	if err != nil {
		return stader.GasInfo{}, err
	}
	return rocketNodeDeposit.GetTransactionGasInfo(opts, "deposit", eth.EthToWei(minimumNodeFee), validatorPubkey[:], validatorSignature[:], depositDataRoot, salt, expectedMinipoolAddress)
}

// Make a node deposit
func Deposit(excm *stader.PermissionlessNodeRegistryContractManager, minimumNodeFee float64, validatorPubkey rptypes.ValidatorPubkey, validatorSignature rptypes.ValidatorSignature, depositDataRoot common.Hash, salt *big.Int, expectedMinipoolAddress common.Address, opts *bind.TransactOpts) (*types.Transaction, error) {
	//rocketNodeDeposit, err := getRocketNodeDeposit(rp, nil)
	//if err != nil {
	//	return nil, err
	//}
	//tx, err := excm.PermissionlessNodeRegistry.Transact(opts, "nodeDeposit", validatorPubkey[:], validatorSignature[:], depositDataRoot, salt, expectedMinipoolAddress)
	//if err != nil {
	//	return nil, fmt.Errorf("Could not make node deposit: %w", err)
	//}
	return nil, nil
}

func StaderNodeDeposit(excm *stader.PermissionlessNodeRegistryContractManager, minimumNodeFee float64, validatorPubkey rptypes.ValidatorPubkey, validatorSignature rptypes.ValidatorSignature, depositDataRoot common.Hash, operatorRewardAddress common.Address, operatorName string, opts *bind.TransactOpts) (*types.Transaction, error) {
	//rocketNodeDeposit, err := getRocketNodeDeposit(rp, nil)
	//if err != nil {
	//	return nil, err
	//}
	//operatorId := big.NewInt(10)
	//tx, err := excm.EthxPermisionlessPoolContract.Transact(opts, "nodeDeposit", validatorPubkey[:], validatorSignature[:], depositDataRoot, operatorRewardAddress, operatorName, operatorId)
	//if err != nil {
	//	return nil, fmt.Errorf("Could not make node deposit: %w", err)
	//}
	//return tx, nil
	return nil, nil
}

// Get the type of a deposit based on the amount
func GetDepositType(rp *stader.PermissionlessNodeRegistryContractManager, amount *big.Int, opts *bind.CallOpts) (rptypes.MinipoolDeposit, error) {
	rocketNodeDeposit, err := getRocketNodeDeposit(rp, opts)
	if err != nil {
		return rptypes.Empty, err
	}

	depositType := new(uint8)
	if err := rocketNodeDeposit.Call(opts, depositType, "getDepositType", amount); err != nil {
		return rptypes.Empty, fmt.Errorf("Could not get deposit type: %w", err)
	}
	return rptypes.MinipoolDeposit(*depositType), nil
}

// Get contracts
var rocketNodeDepositLock sync.Mutex

func getRocketNodeDeposit(rp *stader.PermissionlessNodeRegistryContractManager, opts *bind.CallOpts) (*stader.Contract, error) {
	rocketNodeDepositLock.Lock()
	defer rocketNodeDepositLock.Unlock()
	return &stader.Contract{
		Contract: nil,
		Address:  nil,
		ABI:      nil,
		Client:   nil,
	}, nil
}
