package trustednode

import (
	"fmt"
	"math/big"
	"sync"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"

	trustednodedao "github.com/stader-labs/stader-node/stader-lib/dao/trustednode"
	"github.com/stader-labs/stader-node/stader-lib/stader"
	"github.com/stader-labs/stader-node/stader-lib/utils/eth"
)

// Config
const (
	MembersSettingsContractName       = "rocketDAONodeTrustedSettingsMembers"
	QuorumSettingPath                 = "members.quorum"
	RPLBondSettingPath                = "members.rplbond"
	MinipoolUnbondedMaxSettingPath    = "members.minipool.unbonded.max"
	MinipoolUnbondedMinFeeSettingPath = "members.minipool.unbonded.min.fee"
	ChallengeCooldownSettingPath      = "members.challenge.cooldown"
	ChallengeWindowSettingPath        = "members.challenge.window"
	ChallengeCostSettingPath          = "members.challenge.cost"
)

// Member proposal quorum threshold
func GetQuorum(rp *stader.PermissionlessNodeRegistryContractManager, opts *bind.CallOpts) (float64, error) {
	membersSettingsContract, err := getMembersSettingsContract(rp, opts)
	if err != nil {
		return 0, err
	}
	value := new(*big.Int)
	if err := membersSettingsContract.Call(opts, value, "getQuorum"); err != nil {
		return 0, fmt.Errorf("Could not get member quorum threshold: %w", err)
	}
	return eth.WeiToEth(*value), nil
}
func BootstrapQuorum(rp *stader.PermissionlessNodeRegistryContractManager, value float64, opts *bind.TransactOpts) (common.Hash, error) {
	return trustednodedao.BootstrapUint(rp, MembersSettingsContractName, QuorumSettingPath, eth.EthToWei(value), opts)
}
func ProposeQuorum(rp *stader.PermissionlessNodeRegistryContractManager, value float64, opts *bind.TransactOpts) (uint64, common.Hash, error) {
	return trustednodedao.ProposeSetUint(rp, fmt.Sprintf("set %s", QuorumSettingPath), MembersSettingsContractName, QuorumSettingPath, eth.EthToWei(value), opts)
}
func EstimateProposeQuorumGas(rp *stader.PermissionlessNodeRegistryContractManager, value float64, opts *bind.TransactOpts) (stader.GasInfo, error) {
	return trustednodedao.EstimateProposeSetUintGas(rp, fmt.Sprintf("set %s", QuorumSettingPath), MembersSettingsContractName, QuorumSettingPath, eth.EthToWei(value), opts)
}

// RPL bond required for a member
func GetRPLBond(rp *stader.PermissionlessNodeRegistryContractManager, opts *bind.CallOpts) (*big.Int, error) {
	membersSettingsContract, err := getMembersSettingsContract(rp, opts)
	if err != nil {
		return nil, err
	}
	value := new(*big.Int)
	if err := membersSettingsContract.Call(opts, value, "getRPLBond"); err != nil {
		return nil, fmt.Errorf("Could not get member RPL bond amount: %w", err)
	}
	return *value, nil
}
func BootstrapRPLBond(rp *stader.PermissionlessNodeRegistryContractManager, value *big.Int, opts *bind.TransactOpts) (common.Hash, error) {
	return trustednodedao.BootstrapUint(rp, MembersSettingsContractName, RPLBondSettingPath, value, opts)
}
func ProposeRPLBond(rp *stader.PermissionlessNodeRegistryContractManager, value *big.Int, opts *bind.TransactOpts) (uint64, common.Hash, error) {
	return trustednodedao.ProposeSetUint(rp, fmt.Sprintf("set %s", RPLBondSettingPath), MembersSettingsContractName, RPLBondSettingPath, value, opts)
}
func EstimateProposeRPLBondGas(rp *stader.PermissionlessNodeRegistryContractManager, value *big.Int, opts *bind.TransactOpts) (stader.GasInfo, error) {
	return trustednodedao.EstimateProposeSetUintGas(rp, fmt.Sprintf("set %s", RPLBondSettingPath), MembersSettingsContractName, RPLBondSettingPath, value, opts)
}

// The maximum number of unbonded minipools a member can run
func GetMinipoolUnbondedMax(rp *stader.PermissionlessNodeRegistryContractManager, opts *bind.CallOpts) (uint64, error) {
	membersSettingsContract, err := getMembersSettingsContract(rp, opts)
	if err != nil {
		return 0, err
	}
	value := new(*big.Int)
	if err := membersSettingsContract.Call(opts, value, "getMinipoolUnbondedMax"); err != nil {
		return 0, fmt.Errorf("Could not get member unbonded minipool limit: %w", err)
	}
	return (*value).Uint64(), nil
}
func BootstrapMinipoolUnbondedMax(rp *stader.PermissionlessNodeRegistryContractManager, value uint64, opts *bind.TransactOpts) (common.Hash, error) {
	return trustednodedao.BootstrapUint(rp, MembersSettingsContractName, MinipoolUnbondedMaxSettingPath, big.NewInt(int64(value)), opts)
}
func ProposeMinipoolUnbondedMax(rp *stader.PermissionlessNodeRegistryContractManager, value uint64, opts *bind.TransactOpts) (uint64, common.Hash, error) {
	return trustednodedao.ProposeSetUint(rp, fmt.Sprintf("set %s", MinipoolUnbondedMaxSettingPath), MembersSettingsContractName, MinipoolUnbondedMaxSettingPath, big.NewInt(int64(value)), opts)
}
func EstimateProposeMinipoolUnbondedMaxGas(rp *stader.PermissionlessNodeRegistryContractManager, value uint64, opts *bind.TransactOpts) (stader.GasInfo, error) {
	return trustednodedao.EstimateProposeSetUintGas(rp, fmt.Sprintf("set %s", MinipoolUnbondedMaxSettingPath), MembersSettingsContractName, MinipoolUnbondedMaxSettingPath, big.NewInt(int64(value)), opts)
}

// The minimum commission rate before unbonded minipools are allowed
func GetMinipoolUnbondedMinFee(rp *stader.PermissionlessNodeRegistryContractManager, opts *bind.CallOpts) (uint64, error) {
	membersSettingsContract, err := getMembersSettingsContract(rp, opts)
	if err != nil {
		return 0, err
	}
	value := new(*big.Int)
	if err := membersSettingsContract.Call(opts, value, "getMinipoolUnbondedMinFee"); err != nil {
		return 0, fmt.Errorf("Could not get member unbonded minipool minimum fee: %w", err)
	}
	return (*value).Uint64(), nil
}
func BootstrapMinipoolUnbondedMinFee(rp *stader.PermissionlessNodeRegistryContractManager, value uint64, opts *bind.TransactOpts) (common.Hash, error) {
	return trustednodedao.BootstrapUint(rp, MembersSettingsContractName, MinipoolUnbondedMinFeeSettingPath, big.NewInt(int64(value)), opts)
}
func ProposeMinipoolUnbondedMinFee(rp *stader.PermissionlessNodeRegistryContractManager, value uint64, opts *bind.TransactOpts) (uint64, common.Hash, error) {
	return trustednodedao.ProposeSetUint(rp, fmt.Sprintf("set %s", MinipoolUnbondedMinFeeSettingPath), MembersSettingsContractName, MinipoolUnbondedMinFeeSettingPath, big.NewInt(int64(value)), opts)
}
func EstimateProposeMinipoolUnbondedMinFeeGas(rp *stader.PermissionlessNodeRegistryContractManager, value uint64, opts *bind.TransactOpts) (stader.GasInfo, error) {
	return trustednodedao.EstimateProposeSetUintGas(rp, fmt.Sprintf("set %s", MinipoolUnbondedMinFeeSettingPath), MembersSettingsContractName, MinipoolUnbondedMinFeeSettingPath, big.NewInt(int64(value)), opts)
}

// The period a member must wait for before submitting another challenge, in blocks
func GetChallengeCooldown(rp *stader.PermissionlessNodeRegistryContractManager, opts *bind.CallOpts) (uint64, error) {
	membersSettingsContract, err := getMembersSettingsContract(rp, opts)
	if err != nil {
		return 0, err
	}
	value := new(*big.Int)
	if err := membersSettingsContract.Call(opts, value, "getChallengeCooldown"); err != nil {
		return 0, fmt.Errorf("Could not get member challenge cooldown period: %w", err)
	}
	return (*value).Uint64(), nil
}
func BootstrapChallengeCooldown(rp *stader.PermissionlessNodeRegistryContractManager, value uint64, opts *bind.TransactOpts) (common.Hash, error) {
	return trustednodedao.BootstrapUint(rp, MembersSettingsContractName, ChallengeCooldownSettingPath, big.NewInt(int64(value)), opts)
}
func ProposeChallengeCooldown(rp *stader.PermissionlessNodeRegistryContractManager, value uint64, opts *bind.TransactOpts) (uint64, common.Hash, error) {
	return trustednodedao.ProposeSetUint(rp, fmt.Sprintf("set %s", ChallengeCooldownSettingPath), MembersSettingsContractName, ChallengeCooldownSettingPath, big.NewInt(int64(value)), opts)
}
func EstimateProposeChallengeCooldownGas(rp *stader.PermissionlessNodeRegistryContractManager, value uint64, opts *bind.TransactOpts) (stader.GasInfo, error) {
	return trustednodedao.EstimateProposeSetUintGas(rp, fmt.Sprintf("set %s", ChallengeCooldownSettingPath), MembersSettingsContractName, ChallengeCooldownSettingPath, big.NewInt(int64(value)), opts)
}

// The period during which a member can respond to a challenge, in blocks
func GetChallengeWindow(rp *stader.PermissionlessNodeRegistryContractManager, opts *bind.CallOpts) (uint64, error) {
	membersSettingsContract, err := getMembersSettingsContract(rp, opts)
	if err != nil {
		return 0, err
	}
	value := new(*big.Int)
	if err := membersSettingsContract.Call(opts, value, "getChallengeWindow"); err != nil {
		return 0, fmt.Errorf("Could not get member challenge window period: %w", err)
	}
	return (*value).Uint64(), nil
}
func BootstrapChallengeWindow(rp *stader.PermissionlessNodeRegistryContractManager, value uint64, opts *bind.TransactOpts) (common.Hash, error) {
	return trustednodedao.BootstrapUint(rp, MembersSettingsContractName, ChallengeWindowSettingPath, big.NewInt(int64(value)), opts)
}
func ProposeChallengeWindow(rp *stader.PermissionlessNodeRegistryContractManager, value uint64, opts *bind.TransactOpts) (uint64, common.Hash, error) {
	return trustednodedao.ProposeSetUint(rp, fmt.Sprintf("set %s", ChallengeWindowSettingPath), MembersSettingsContractName, ChallengeWindowSettingPath, big.NewInt(int64(value)), opts)
}
func EstimateProposeChallengeWindowGas(rp *stader.PermissionlessNodeRegistryContractManager, value uint64, opts *bind.TransactOpts) (stader.GasInfo, error) {
	return trustednodedao.EstimateProposeSetUintGas(rp, fmt.Sprintf("set %s", ChallengeWindowSettingPath), MembersSettingsContractName, ChallengeWindowSettingPath, big.NewInt(int64(value)), opts)
}

// The fee for a non-member to challenge a member, in wei
func GetChallengeCost(rp *stader.PermissionlessNodeRegistryContractManager, opts *bind.CallOpts) (*big.Int, error) {
	membersSettingsContract, err := getMembersSettingsContract(rp, opts)
	if err != nil {
		return nil, err
	}
	value := new(*big.Int)
	if err := membersSettingsContract.Call(opts, value, "getChallengeCost"); err != nil {
		return nil, fmt.Errorf("Could not get member challenge cost: %w", err)
	}
	return *value, nil
}
func BootstrapChallengeCost(rp *stader.PermissionlessNodeRegistryContractManager, value *big.Int, opts *bind.TransactOpts) (common.Hash, error) {
	return trustednodedao.BootstrapUint(rp, MembersSettingsContractName, ChallengeCostSettingPath, value, opts)
}
func ProposeChallengeCost(rp *stader.PermissionlessNodeRegistryContractManager, value *big.Int, opts *bind.TransactOpts) (uint64, common.Hash, error) {
	return trustednodedao.ProposeSetUint(rp, fmt.Sprintf("set %s", ChallengeCostSettingPath), MembersSettingsContractName, ChallengeCostSettingPath, value, opts)
}
func EstimateProposeChallengeCostGas(rp *stader.PermissionlessNodeRegistryContractManager, value *big.Int, opts *bind.TransactOpts) (stader.GasInfo, error) {
	return trustednodedao.EstimateProposeSetUintGas(rp, fmt.Sprintf("set %s", ChallengeCostSettingPath), MembersSettingsContractName, ChallengeCostSettingPath, value, opts)
}

// Get contracts
var membersSettingsContractLock sync.Mutex

func getMembersSettingsContract(rp *stader.PermissionlessNodeRegistryContractManager, opts *bind.CallOpts) (*stader.Contract, error) {
	membersSettingsContractLock.Lock()
	defer membersSettingsContractLock.Unlock()
	return nil, nil
	//return rp.GetContract(MembersSettingsContractName, opts)
}
