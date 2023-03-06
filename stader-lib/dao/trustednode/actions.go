package trustednode

import (
	"fmt"
	"sync"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"

	"github.com/stader-labs/stader-node/stader-lib/stader"
)

// Estimate the gas of Join
func EstimateJoinGas(rp *stader.PermissionlessNodeRegistryContractManager, opts *bind.TransactOpts) (stader.GasInfo, error) {
	rocketDAONodeTrustedActions, err := getRocketDAONodeTrustedActions(rp, nil)
	if err != nil {
		return stader.GasInfo{}, err
	}
	return rocketDAONodeTrustedActions.GetTransactionGasInfo(opts, "actionJoin")
}

// Join the trusted node DAO
// Requires an executed invite proposal
func Join(rp *stader.PermissionlessNodeRegistryContractManager, opts *bind.TransactOpts) (common.Hash, error) {
	rocketDAONodeTrustedActions, err := getRocketDAONodeTrustedActions(rp, nil)
	if err != nil {
		return common.Hash{}, err
	}
	tx, err := rocketDAONodeTrustedActions.Transact(opts, "actionJoin")
	if err != nil {
		return common.Hash{}, fmt.Errorf("Could not join the trusted node DAO: %w", err)
	}
	return tx.Hash(), nil
}

// Estimate the gas of Leave
func EstimateLeaveGas(rp *stader.PermissionlessNodeRegistryContractManager, rplBondRefundAddress common.Address, opts *bind.TransactOpts) (stader.GasInfo, error) {
	rocketDAONodeTrustedActions, err := getRocketDAONodeTrustedActions(rp, nil)
	if err != nil {
		return stader.GasInfo{}, err
	}
	return rocketDAONodeTrustedActions.GetTransactionGasInfo(opts, "actionLeave", rplBondRefundAddress)
}

// Leave the trusted node DAO
// Requires an executed leave proposal
func Leave(rp *stader.PermissionlessNodeRegistryContractManager, rplBondRefundAddress common.Address, opts *bind.TransactOpts) (common.Hash, error) {
	rocketDAONodeTrustedActions, err := getRocketDAONodeTrustedActions(rp, nil)
	if err != nil {
		return common.Hash{}, err
	}
	tx, err := rocketDAONodeTrustedActions.Transact(opts, "actionLeave", rplBondRefundAddress)
	if err != nil {
		return common.Hash{}, fmt.Errorf("Could not leave the trusted node DAO: %w", err)
	}
	return tx.Hash(), nil
}

// Estimate the gas of MakeChallenge
func EstimateMakeChallengeGas(rp *stader.PermissionlessNodeRegistryContractManager, memberAddress common.Address, opts *bind.TransactOpts) (stader.GasInfo, error) {
	rocketDAONodeTrustedActions, err := getRocketDAONodeTrustedActions(rp, nil)
	if err != nil {
		return stader.GasInfo{}, err
	}
	return rocketDAONodeTrustedActions.GetTransactionGasInfo(opts, "actionChallengeMake", memberAddress)
}

// Make a challenge against a node
func MakeChallenge(rp *stader.PermissionlessNodeRegistryContractManager, memberAddress common.Address, opts *bind.TransactOpts) (common.Hash, error) {
	rocketDAONodeTrustedActions, err := getRocketDAONodeTrustedActions(rp, nil)
	if err != nil {
		return common.Hash{}, err
	}
	tx, err := rocketDAONodeTrustedActions.Transact(opts, "actionChallengeMake", memberAddress)
	if err != nil {
		return common.Hash{}, fmt.Errorf("Could not challenge trusted node DAO member %s: %w", memberAddress.Hex(), err)
	}
	return tx.Hash(), nil
}

// Estimate the gas of DecideChallenge
func EstimateDecideChallengeGas(rp *stader.PermissionlessNodeRegistryContractManager, memberAddress common.Address, opts *bind.TransactOpts) (stader.GasInfo, error) {
	rocketDAONodeTrustedActions, err := getRocketDAONodeTrustedActions(rp, nil)
	if err != nil {
		return stader.GasInfo{}, err
	}
	return rocketDAONodeTrustedActions.GetTransactionGasInfo(opts, "actionChallengeDecide", memberAddress)
}

// Decide a challenge against a node
func DecideChallenge(rp *stader.PermissionlessNodeRegistryContractManager, memberAddress common.Address, opts *bind.TransactOpts) (common.Hash, error) {
	rocketDAONodeTrustedActions, err := getRocketDAONodeTrustedActions(rp, nil)
	if err != nil {
		return common.Hash{}, err
	}
	tx, err := rocketDAONodeTrustedActions.Transact(opts, "actionChallengeDecide", memberAddress)
	if err != nil {
		return common.Hash{}, fmt.Errorf("Could not decide the challenge against trusted node DAO member %s: %w", memberAddress.Hex(), err)
	}
	return tx.Hash(), nil
}

// Get contracts
var rocketDAONodeTrustedActionsLock sync.Mutex

func getRocketDAONodeTrustedActions(rp *stader.PermissionlessNodeRegistryContractManager, opts *bind.CallOpts) (*stader.Contract, error) {
	rocketDAONodeTrustedActionsLock.Lock()
	defer rocketDAONodeTrustedActionsLock.Unlock()
	return nil, nil
	//return rp.GetContract("rocketDAONodeTrustedActions", opts)
}
