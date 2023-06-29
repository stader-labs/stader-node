// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package contracts

import (
	"errors"
	"math/big"
	"strings"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/event"
)

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = errors.New
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
	_ = abi.ConvertType
)

// StaderConfigMetaData contains all meta data concerning the StaderConfig contract.
var StaderConfigMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_admin\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_ethDepositContract\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"InvalidLimits\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidMaxDepositValue\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidMaxWithdrawValue\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidMinDepositValue\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidMinWithdrawValue\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ZeroAddress\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"version\",\"type\":\"uint8\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"previousAdminRole\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"newAdminRole\",\"type\":\"bytes32\"}],\"name\":\"RoleAdminChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleGranted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleRevoked\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"key\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newAddress\",\"type\":\"address\"}],\"name\":\"SetAccount\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"key\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"SetConstant\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"key\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newAddress\",\"type\":\"address\"}],\"name\":\"SetContract\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"key\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newAddress\",\"type\":\"address\"}],\"name\":\"SetToken\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"key\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"SetVariable\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"ADMIN\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"AUCTION_CONTRACT\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"DECIMALS\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"DEFAULT_ADMIN_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"ETHX_SUPPLY_POR_FEED\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"ETH_BALANCE_POR_FEED\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"ETH_DEPOSIT_CONTRACT\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"ETH_PER_NODE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"ETHx\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"FULL_DEPOSIT_SIZE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"MANAGER\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"MAX_DEPOSIT_AMOUNT\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"MAX_WITHDRAW_AMOUNT\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"MIN_BLOCK_DELAY_TO_FINALIZE_WITHDRAW_REQUEST\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"MIN_DEPOSIT_AMOUNT\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"MIN_WITHDRAW_AMOUNT\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"NODE_EL_REWARD_VAULT_IMPLEMENTATION\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"OPERATOR\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"OPERATOR_MAX_NAME_LENGTH\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"OPERATOR_REWARD_COLLECTOR\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"PENALTY_CONTRACT\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"PERMISSIONED_NODE_REGISTRY\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"PERMISSIONED_POOL\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"PERMISSIONED_SOCIALIZING_POOL\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"PERMISSIONLESS_NODE_REGISTRY\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"PERMISSIONLESS_POOL\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"PERMISSIONLESS_SOCIALIZING_POOL\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"POOL_SELECTOR\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"POOL_UTILS\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"PRE_DEPOSIT_SIZE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"REWARD_THRESHOLD\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"SD\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"SD_COLLATERAL\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"SOCIALIZING_POOL_CYCLE_DURATION\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"SOCIALIZING_POOL_OPT_IN_COOLING_PERIOD\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"STADER_INSURANCE_FUND\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"STADER_ORACLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"STADER_TREASURY\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"STAKE_POOL_MANAGER\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"TOTAL_FEE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"USER_WITHDRAW_MANAGER\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"VALIDATOR_WITHDRAWAL_VAULT_IMPLEMENTATION\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"VAULT_FACTORY\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"WITHDRAWN_KEYS_BATCH_SIZE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getAdmin\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getAuctionContract\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getDecimals\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getETHBalancePORFeedProxy\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getETHDepositContract\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getETHXSupplyPORFeedProxy\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getETHxToken\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getFullDepositSize\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getMaxDepositAmount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getMaxWithdrawAmount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getMinBlockDelayToFinalizeWithdrawRequest\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getMinDepositAmount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getMinWithdrawAmount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getNodeELRewardVaultImplementation\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getOperatorMaxNameLength\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getOperatorRewardsCollector\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getPenaltyContract\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getPermissionedNodeRegistry\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getPermissionedPool\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getPermissionedSocializingPool\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getPermissionlessNodeRegistry\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getPermissionlessPool\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getPermissionlessSocializingPool\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getPoolSelector\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getPoolUtils\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getPreDepositSize\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getRewardsThreshold\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"}],\"name\":\"getRoleAdmin\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getSDCollateral\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getSocializingPoolCycleDuration\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getSocializingPoolOptInCoolingPeriod\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getStaderInsuranceFund\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getStaderOracle\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getStaderToken\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getStaderTreasury\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getStakePoolManager\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getStakedEthPerNode\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getTotalFee\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getUserWithdrawManager\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getValidatorWithdrawalVaultImplementation\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getVaultFactory\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getWithdrawnKeyBatchSize\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"grantRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"hasRole\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"onlyManagerRole\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"onlyOperatorRole\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_addr\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"_contractName\",\"type\":\"bytes32\"}],\"name\":\"onlyStaderContract\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"renounceRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"revokeRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_admin\",\"type\":\"address\"}],\"name\":\"updateAdmin\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_auctionContract\",\"type\":\"address\"}],\"name\":\"updateAuctionContract\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_ethBalanceProxy\",\"type\":\"address\"}],\"name\":\"updateETHBalancePORFeedProxy\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_ethXSupplyProxy\",\"type\":\"address\"}],\"name\":\"updateETHXSupplyPORFeedProxy\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_ethX\",\"type\":\"address\"}],\"name\":\"updateETHxToken\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_maxDepositAmount\",\"type\":\"uint256\"}],\"name\":\"updateMaxDepositAmount\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_maxWithdrawAmount\",\"type\":\"uint256\"}],\"name\":\"updateMaxWithdrawAmount\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_minBlockDelay\",\"type\":\"uint256\"}],\"name\":\"updateMinBlockDelayToFinalizeWithdrawRequest\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_minDepositAmount\",\"type\":\"uint256\"}],\"name\":\"updateMinDepositAmount\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_minWithdrawAmount\",\"type\":\"uint256\"}],\"name\":\"updateMinWithdrawAmount\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_nodeELRewardVaultImpl\",\"type\":\"address\"}],\"name\":\"updateNodeELRewardImplementation\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_operatorRewardsCollector\",\"type\":\"address\"}],\"name\":\"updateOperatorRewardsCollector\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_penaltyContract\",\"type\":\"address\"}],\"name\":\"updatePenaltyContract\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_permissionedNodeRegistry\",\"type\":\"address\"}],\"name\":\"updatePermissionedNodeRegistry\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_permissionedPool\",\"type\":\"address\"}],\"name\":\"updatePermissionedPool\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_permissionedSocializePool\",\"type\":\"address\"}],\"name\":\"updatePermissionedSocializingPool\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_permissionlessNodeRegistry\",\"type\":\"address\"}],\"name\":\"updatePermissionlessNodeRegistry\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_permissionlessPool\",\"type\":\"address\"}],\"name\":\"updatePermissionlessPool\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_permissionlessSocializePool\",\"type\":\"address\"}],\"name\":\"updatePermissionlessSocializingPool\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_poolSelector\",\"type\":\"address\"}],\"name\":\"updatePoolSelector\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_poolUtils\",\"type\":\"address\"}],\"name\":\"updatePoolUtils\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_rewardsThreshold\",\"type\":\"uint256\"}],\"name\":\"updateRewardsThreshold\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_sdCollateral\",\"type\":\"address\"}],\"name\":\"updateSDCollateral\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_socializingPoolCycleDuration\",\"type\":\"uint256\"}],\"name\":\"updateSocializingPoolCycleDuration\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_SocializePoolOptInCoolingPeriod\",\"type\":\"uint256\"}],\"name\":\"updateSocializingPoolOptInCoolingPeriod\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_staderInsuranceFund\",\"type\":\"address\"}],\"name\":\"updateStaderInsuranceFund\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_staderOracle\",\"type\":\"address\"}],\"name\":\"updateStaderOracle\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_staderToken\",\"type\":\"address\"}],\"name\":\"updateStaderToken\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_staderTreasury\",\"type\":\"address\"}],\"name\":\"updateStaderTreasury\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_stakePoolManager\",\"type\":\"address\"}],\"name\":\"updateStakePoolManager\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_userWithdrawManager\",\"type\":\"address\"}],\"name\":\"updateUserWithdrawManager\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_validatorWithdrawalVaultImpl\",\"type\":\"address\"}],\"name\":\"updateValidatorWithdrawalVaultImplementation\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_vaultFactory\",\"type\":\"address\"}],\"name\":\"updateVaultFactory\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_withdrawnKeysBatchSize\",\"type\":\"uint256\"}],\"name\":\"updateWithdrawnKeysBatchSize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x608060405234801562000010575f80fd5b506040516200373e3803806200373e83398101604081905262000033916200062c565b5f54610100900460ff16158080156200005257505f54600160ff909116105b806200006d5750303b1580156200006d57505f5460ff166001145b620000d65760405162461bcd60e51b815260206004820152602e60248201527f496e697469616c697a61626c653a20636f6e747261637420697320616c72656160448201526d191e481a5b9a5d1a585b1a5e995960921b60648201526084015b60405180910390fd5b5f805460ff191660011790558015620000f8575f805461ff0019166101001790555b6200010383620003f7565b6200010e82620003f7565b6200011862000422565b6200014d7ff822b1f0c3b886ce1cdf1c2a5317844145470db33b02c63cae4813f8c9b2dc176801bc16d674ec80000062000490565b620001817f9b1ae66636378b5626322a52e22518dd40bb04881cf0440ed16a20c0f902b242670de0b6b3a764000062000490565b620001b67f876943525608da6d95be5925fe6c4fe80e8622c8a76e7414f80e8ba210e0711c6801ae361fc1451c000062000490565b620001e47f33271b56873d8abb908de4853f90a8a0ef8829548ec0bf6c298feed3917c50a261271062000490565b620002187f08593985ae1bebfb02f6c30105edffb176a6d87c9fad54c434bf9b58f67e81b6670de0b6b3a764000062000490565b620002457f59b5f464ec5829246a81f005456c8cb714ee224aea800742e2dae497263e466960ff62000490565b620002777ffa5a84fed05ba4c93fcc5ba1f4ad010e3bef3e6394b367aa10b3ec01997375cc655af3107a4000620004cd565b620002ad7f712c13b90acf399d7bc7625370ce37c64b5eba41011b0961a88c2ef1648870cf69021e19e0c9bab2400000620004cd565b620002df7fb18278bb399a7088b8b0b26f4896d5ebaba4497c611bbe9d43abe92d9a1fe83d655af3107a4000620004cd565b620003157f1c2fe98ddbbbffbcf7735c7446ffcddb5ccd2a4ec2ace0f7d90f73e9ff13fcc769021e19e0c9bab2400000620004cd565b620003427f6f8d0b773ad4970d3e7d47623dc9ce06a1b4fe833bf451d06a47e774f9acaa636032620004cd565b620003707f2cf2377da51daa9c0d7e3f98c7532a67ee5e9398afad7b7db6e578b978a27094610258620004cd565b6200039c7f84b42b3d5e6851893d4418c6ebc9a4727e78afdf84e73674c8b9c1c2b1904e2d8362000503565b620003a85f846200056d565b8015620003ee575f805461ff0019169055604051600181527f7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb38474024989060200160405180910390a15b50505062000662565b6001600160a01b0381166200041f5760405163d92e233d60e01b815260040160405180910390fd5b50565b5f54610100900460ff166200048e5760405162461bcd60e51b815260206004820152602b60248201527f496e697469616c697a61626c653a20636f6e7472616374206973206e6f74206960448201526a6e697469616c697a696e6760a81b6064820152608401620000cd565b565b5f8281526097602090815260409182902083905581518481529081018390525f805160206200371e83398151915291015b60405180910390a15050565b5f8281526098602090815260409182902083905581518481529081018390525f805160206200371e8339815191529101620004c1565b6200050e81620003f7565b5f828152609a602090815260409182902080546001600160a01b0319166001600160a01b0385169081179091558251858152918201527f5de40a806536a2029221dac2c8887ac9f11952fcc1ed3d7cfb4476dd5259b7409101620004c1565b5f8281526065602090815260408083206001600160a01b038516845290915290205460ff166200060c575f8281526065602090815260408083206001600160a01b03851684529091529020805460ff19166001179055620005cb3390565b6001600160a01b0316816001600160a01b0316837f2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d60405160405180910390a45b5050565b80516001600160a01b038116811462000627575f80fd5b919050565b5f80604083850312156200063e575f80fd5b620006498362000610565b9150620006596020840162000610565b90509250929050565b6130ae80620006705f395ff3fe608060405234801561000f575f80fd5b506004361061076f575f3560e01c80636870bb2b116103cd578063b11c699d11610200578063defd024d1161011f578063f122961f116100b4578063f83c778711610084578063f83c778714611d94578063fa71fcbb14611da7578063ff387f3a14611df6578063ff4f354614611e45575f80fd5b8063f122961f14611d0a578063f4914d3314611d31578063f63718e714611d5a578063f6c278c114611d6d575f80fd5b8063e7bdba32116100ef578063e7bdba3214611be4578063e8fe187314611c0b578063ecf170a814611c63578063f0141d8414611cbb575f80fd5b8063defd024d14611b0e578063e069f71414611b66578063e2f273bd14611bbe578063e4f59b6c14611bd1575f80fd5b8063bedcb34c11610195578063cc45dabe11610165578063cc45dabe14611a2d578063d2cee8ba14611a85578063d547741f14611ad4578063dde63e8f14611ae7575f80fd5b8063bedcb34c146119b9578063c20573c1146119e0578063c60470d3146119f3578063ca78360c14611a1a575f80fd5b8063b5cfee6c116101d0578063b5cfee6c14611927578063b68578441461197f578063b9894a1114611993578063bbb99bb5146119a6575f80fd5b8063b11c699d14611871578063b312392214611898578063b479a517146118c5578063b549dbff14611914575f80fd5b806384780205116102ec57806398c3592711610281578063a469e24711610251578063a469e24714611787578063a53bddd6146117df578063aa2f56c714611806578063aa95379514611819575f80fd5b806398c35927146116bd5780639ca76b73146116d0578063a0b4079f14611728578063a217fddf14611780575f80fd5b80638a4cfb58116102bc5780638a4cfb58146116185780638f8b38671461162b57806391d1485414611683578063983d273714611696575f80fd5b8063847802051461155f57806385e2fcd31461157257806388993d8b146115995780638910115c146115c0575f80fd5b806377e8a0c3116103625780637ae316d0116103325780637ae316d0146114c35780637b4cd7ec146115125780638314859314611525578063841b83b314611538575f80fd5b806377e8a0c31461142757806379175a741461144e578063792c8cc3146114755780637a87fa0b1461149c575f80fd5b80636e9960c31161039d5780636e9960c31461136457806372195b3e146113a9578063723b732c146113bc57806372ce78b0146113cf575f80fd5b80636870bb2b146112525780636ccb9d70146112655780636d28ad1c146112bd5780636e0fddfc14611315575f80fd5b80632f2ff15d116105a55780634c34a982116104c45780635b9cc8b1116104595780636240fb9c116104295780636240fb9c146111ca57806363db7eae146111dd57806367dcf13414611204578063686a8b671461122b575f80fd5b80635b9cc8b1146111255780635be6ce69146111385780635edc686e1461114b5780636176bbde146111a3575f80fd5b80635458a106116104945780635458a106146110585780635726a356146110b0578063572c686a146110ff5780635b5961fc14611112575f80fd5b80634c34a98214610fd057806352112bd314610ff757806353f5713b1461101e5780635455e47214611031575f80fd5b8063384002a21161053a578063403efe7f1161050a578063403efe7f14610f2a5780634191e0fe14610f3d57806344ba0ea214610f64578063489ed65114610f78575f80fd5b8063384002a214610ea25780633871d0f114610ec95780633b6bcca014610ef05780633c128dad14610f17575f80fd5b806336568abe1161057557806336568abe14610dfd57806336854d6314610e10578063368f9d1714610e3757806336c157f414610e4a575f80fd5b80632f2ff15d14610d44578063326a16a314610d5757806334d17d7414610d93578063360374a414610da6575f80fd5b806318bcb28411610691578063248a9ca3116106265780632a9cc2c4116105f65780632a9cc2c414610c465780632ca03f6614610c6d5780632e0f262514610cc55780632ec5e01814610cec575f80fd5b8063248a9ca314610ba55780632651644c14610bc7578063278671bb14610bda5780632a0acc6a14610c32575f80fd5b80631c55cccd116106615780631c55cccd14610b085780631ca197a514610b2f5780631de03db814610b7e5780631ea30fef14610b91575f80fd5b806318bcb28414610a4e5780631af0fff314610aa65780631b2df85014610acd5780631bf6a41c14610ae1575f80fd5b8063103f290711610707578063121669f1116106d7578063121669f11461099d57806314e1b8fd146109b0578063152a91da146109d957806318829fc3146109ff575f80fd5b8063103f2907146108ed5780631049e32e1461091457806310deba2b146109275780631202007514610976575f80fd5b8063088ee72d11610742578063088ee72d146108345780630945d42c146108475780630a3fbd9a1461085a5780630bdf3166146108c6575f80fd5b806301ffc9a7146107735780630430246e1461079b578063047cb439146107d057806308297645146107e5575b5f80fd5b610786610781366004612d91565b611e58565b60405190151581526020015b60405180910390f35b6107c27f9b1ae66636378b5626322a52e22518dd40bb04881cf0440ed16a20c0f902b24281565b604051908152602001610792565b6107e36107de366004612dd3565b611e8e565b005b7f9b1ae66636378b5626322a52e22518dd40bb04881cf0440ed16a20c0f902b2425f5260976020527f2b5f44404b80fc874d00ce3803444dc1d8415bef002ea5e3d4c6a1fc229b361b546107c2565b6107e3610842366004612dd3565b611ec6565b6107e3610855366004612dec565b611efa565b7fc5b1a6a0b843563e6a17ca90bc59d2315c523be427d0c9c2ba08d77ced4f46b15f52609a6020527f93bda0178f178a956e1154aad6f6d04aca130dc29bb626bd6774e853c8c9f354546001600160a01b03165b6040516001600160a01b039091168152602001610792565b6107c27f3e4ded42f360c2e6b1251d584085ae1d9aa9cbed18687fac6b6aef8eed1c5ad381565b6107c27f8d4341681b282735dd0d55670ff8e0ad68a80cbfc2cee847065e9f771470f88f81565b6107e3610922366004612dd3565b611f43565b7f59b5f464ec5829246a81f005456c8cb714ee224aea800742e2dae497263e46695f5260976020527ff1d631be95f382e871541957d68e9595b265874c488308836f37d0f22a9fbae9546107c2565b6107c27f4c9466ca1bf288a7334a7494f09a0acc38ee31628eaf8c68b574b9f0ec22a9c181565b6107e36109ab366004612dd3565b611f77565b5f80516020612fd98339815191525f5260986020525f80516020613019833981519152546107c2565b6107c27e665c1b06e0667c56a1ca1706b7573435d1b9162c6327b5d0ea1daeb491ad0d81565b7f46b41285bb7b8513ce3a9d95cdf6916699fb00b47326e8d3850be1b6186e03495f5260986020527f4d508419d31c3547aff85909df3c1fcaa249c360d3c9fa4e4f9e9c899cebbedc546107c2565b7f8d4341681b282735dd0d55670ff8e0ad68a80cbfc2cee847065e9f771470f88f5f52609a6020527f510a692d092451633b86b6d5ebd49dd58b5ea01b6d0783a379a8169a08baac9f546001600160a01b03166108ae565b6107c27fe5240448c78dfcff5bda4e4eed69ba9635df15d79da0e8a4cf889217106fa45b81565b6107c25f80516020612f9983398151915281565b6107c27f29384ec8473b541e7a7850226a4d1906a700f14cc394266ee08800ba62dc3af981565b6107c27fbd34382cd421c5250595893a4ed6cdb2125e6be7d5e0a9dbc469de5d583adfcf81565b7f3c6dcff840f36f9818a73b67d9d00197362f63687bd52e3c277bd0ffb30dde335f5260986020527f9e4fbca7af476428837bb1c0659b29a978bd5be1038b9848cfd6837f97c0c036546107c2565b6107e3610b8c366004612dd3565b611fab565b6107c25f8051602061305983398151915281565b6107c2610bb3366004612dec565b5f9081526065602052604090206001015490565b6107e3610bd5366004612dd3565b611fdf565b7f5be667ef1f4c6c279e2aa7e62595a1045043db6a43145cb438c6d36e7a3c3ed85f52609a6020527f3f1c1b82007b7a87a83473281505b32822fde2464206a16635328330125264a8546001600160a01b03166108ae565b6107c25f80516020612ff983398151915281565b6107c27fb134afa3abad633a84ab2d33dd5171f2b371e38b0f7bca001383aaf08ed6d2d181565b7fb134afa3abad633a84ab2d33dd5171f2b371e38b0f7bca001383aaf08ed6d2d15f52609a6020527fb5c61d48a513a298b438559aede2612ccf11b8fe4c725b0f159efab727297353546001600160a01b03166108ae565b6107c27f08593985ae1bebfb02f6c30105edffb176a6d87c9fad54c434bf9b58f67e81b681565b7f3d88d1233771c5c30791fb6805b7f91424dae1e5a68a57da846ca7ff83c640295f52609a6020527f018f2aef664aeeb1561d5a44d318b67f16f75b697bf95eeabc62c48d36323e72546001600160a01b03166108ae565b6107e3610d52366004612e03565b612013565b5f80516020612fb98339815191525f5260986020527fd179a4a9329ee39fba707fd91c699ec0f088afc56731eb89ff424b873ac70844546107c2565b6107e3610da1366004612dd3565b61203c565b7e665c1b06e0667c56a1ca1706b7573435d1b9162c6327b5d0ea1daeb491ad0d5f52609a6020527fe107fed811895732bef768006b62e8ce98d10a188d78cab697a91a201b5e2404546001600160a01b03166108ae565b6107e3610e0b366004612e03565b612070565b6107c27ff935b8bf66b325637ad32ca875b588849cf4026791b79b4dc20623cd3dd36e2081565b6107e3610e45366004612dd3565b6120ef565b7f8e96355022bb9b9f4d9d4e01fe2b58f45e78549c982c401c96f75f33c5de457e5f52609a6020527fe74d6d5cda9d4a34ee9d4950f99c58c26803c1cf17dbd9d3e9f82fcea7feb01e546001600160a01b03166108ae565b6107c27f95bf18d68834a11aaae7b73ff6037326f163a81a7b5ea80cba96856ce2284fbd81565b6107c27f9f919a2294d86593fbcec81ea71aa683cec51c78771c642f8894ba8f3949705281565b6107c27fc5b1a6a0b843563e6a17ca90bc59d2315c523be427d0c9c2ba08d77ced4f46b181565b6107e3610f25366004612dd3565b612123565b6107e3610f38366004612dd3565b612157565b6107c27fa4083e7a78dd898def03c51ce199cb4286b8828be4f6f46e04aec6176196747181565b6107c25f80516020612fb983398151915281565b7f690795c57e13eaf2526f76202b6799e9afdb069afca1e572f693953d013569d85f52609a6020527f38e84315fdfc8f1b16767d9fd043998a9ff60cfbcb629d8f48542b4e3ee87096546001600160a01b03166108ae565b6107c27f712c13b90acf399d7bc7625370ce37c64b5eba41011b0961a88c2ef1648870cf81565b6107c27f602490b12960e59ddb584affd1da6cd5692f4455c1ba0cc4e865af81e111ebe281565b61078661102c366004612dd3565b61218b565b6107c27f59b5f464ec5829246a81f005456c8cb714ee224aea800742e2dae497263e466981565b7fdb5d1c2a9350ca010dcdf3953da11a9e8f7c5e2918cdfa65500e84e7fd4fde7d5f52609a6020527f642611b82cedca4c0a5510e3234bea9632cc7eb6e135d12e2ef4f8c68dc23add546001600160a01b03166108ae565b7f712c13b90acf399d7bc7625370ce37c64b5eba41011b0961a88c2ef1648870cf5f5260986020527ffcacc1044a5a1b4eb9c058396306426a857813d37a4fb6ccf5a3adde30e0c914546107c2565b6107e361110d366004612dec565b6121b6565b6107e3611120366004612dd3565b6121f7565b6107e3611133366004612dec565b61222b565b6107e3611146366004612dd3565b61224c565b7fa4083e7a78dd898def03c51ce199cb4286b8828be4f6f46e04aec617619674715f52609a6020527f863e03b3878962463f3668c14c10a4aeeabb7baa9c7a9b990796f179109d8692546001600160a01b03166108ae565b6107c27f2cf2377da51daa9c0d7e3f98c7532a67ee5e9398afad7b7db6e578b978a2709481565b6107866111d8366004612dd3565b612280565b6107c27f33271b56873d8abb908de4853f90a8a0ef8829548ec0bf6c298feed3917c50a281565b6107c27ff822b1f0c3b886ce1cdf1c2a5317844145470db33b02c63cae4813f8c9b2dc1781565b6107c27fc54a7590fe6738d7a81f393c1cf5ab3e577c91781037d93a5a9f5ce44f19eb5181565b6107e3611260366004612dec565b612298565b7fd7e49a298cb2719de62e5df1024257eed316db6337361b3a30d56a75324046075f52609a6020527f294ce448c5d68d362948bb2b78c5571986464589b6911cc804ca52d7abbad2e3546001600160a01b03166108ae565b7fbd34382cd421c5250595893a4ed6cdb2125e6be7d5e0a9dbc469de5d583adfcf5f52609a6020527f3195564ffd56571794a8c7ffc14e3d393758b399f23318e874273db13addfdfe546001600160a01b03166108ae565b7fc54a7590fe6738d7a81f393c1cf5ab3e577c91781037d93a5a9f5ce44f19eb515f5260986020527f4d985796191711ecc0d75f056488220f1f755856cdfe3ebd45de3537c37b9b50546107c2565b5f80516020612ff98339815191525f5260996020527f15be86566e203c1f41b9ae149d9fbb01b2c14f503704423d739a6e3d2db5a9ee546001600160a01b03166108ae565b6107e36113b7366004612dd3565b6122d9565b6107e36113ca366004612dec565b61230d565b7f8567f5af844d68168987760a7ce1762804b9de703165fc50ce4fa85246016c915f5260996020527f2c1f6cfa08e101d854b66353df53d6eb32e981bfc1a8351f458fd54b64cfc181546001600160a01b03166108ae565b6107c27f84b42b3d5e6851893d4418c6ebc9a4727e78afdf84e73674c8b9c1c2b1904e2d81565b6107c27f5be667ef1f4c6c279e2aa7e62595a1045043db6a43145cb438c6d36e7a3c3ed881565b6107c27f876943525608da6d95be5925fe6c4fe80e8622c8a76e7414f80e8ba210e0711c81565b6107c27f76d62e541b8d573110ca3eb9003e96426f530422a76712d1356f6c6ce50541ca81565b7f33271b56873d8abb908de4853f90a8a0ef8829548ec0bf6c298feed3917c50a25f5260976020527f799f922a2554690a852ce3427a174a9d0f64f94f53730bd0c6e1e1fdc54799ae546107c2565b6107e3611520366004612dd3565b612341565b6107e3611533366004612dd3565b612382565b6107c27f8567f5af844d68168987760a7ce1762804b9de703165fc50ce4fa85246016c9181565b6107e361156d366004612dec565b6123b6565b6107c27fd7e49a298cb2719de62e5df1024257eed316db6337361b3a30d56a753240460781565b6107c27f6f8d0b773ad4970d3e7d47623dc9ce06a1b4fe833bf451d06a47e774f9acaa6381565b7f29384ec8473b541e7a7850226a4d1906a700f14cc394266ee08800ba62dc3af95f52609a6020527f249a87d52af73222d4a479ebe40b904ebabf543d4706240658e6092ca9388c26546001600160a01b03166108ae565b6107e3611626366004612dec565b6123e4565b7f84b42b3d5e6851893d4418c6ebc9a4727e78afdf84e73674c8b9c1c2b1904e2d5f52609a6020527f492656d26f3accf1cea0a783c131178deb1c8733d9c679e5cecde8df27a9ad95546001600160a01b03166108ae565b610786611691366004612e03565b612425565b6107c27f523a704056dcd17bcf83bed8b68c59416dac1119be77755efe3bde0a64e46e0c81565b6107e36116cb366004612dd3565b61244f565b7f76d62e541b8d573110ca3eb9003e96426f530422a76712d1356f6c6ce50541ca5f52609a6020527f86012a00795dbb89a313ebfe1e3a458a84ce87cdb7c6a7971caf999119513627546001600160a01b03166108ae565b7f602490b12960e59ddb584affd1da6cd5692f4455c1ba0cc4e865af81e111ebe25f52609a6020527fd54531c6bba5beed207277daa8e0e65bdfb6aece3f974fb0394154eb989d1d42546001600160a01b03166108ae565b6107c25f81565b7f4c9466ca1bf288a7334a7494f09a0acc38ee31628eaf8c68b574b9f0ec22a9c15f52609a6020527f2df8b6a0a0cdef82de21edc971a252888647231024af6c12c533010687315b1f546001600160a01b03166108ae565b6107c27f3d88d1233771c5c30791fb6805b7f91424dae1e5a68a57da846ca7ff83c6402981565b6107e3611814366004612dd3565b612483565b7f5c00ec259bace293b50174e499c413ca897b4bcb54ed468b7e6bade51c6a9f965f52609a6020527fcda3409ebc466b6ac691341dcf169fdb28e448f6cf860239292340843aa52984546001600160a01b03166108ae565b6107c27f8e96355022bb9b9f4d9d4e01fe2b58f45e78549c982c401c96f75f33c5de457e81565b6107866118a6366004612e2d565b5f908152609a60205260409020546001600160a01b0390811691161490565b7f6f8d0b773ad4970d3e7d47623dc9ce06a1b4fe833bf451d06a47e774f9acaa635f5260986020527f72873426992e590ffa79a15175a7f2c8cf191cf402b7484af189cd125376fcdc546107c2565b6107e3611922366004612dd3565b6124b7565b7fe5240448c78dfcff5bda4e4eed69ba9635df15d79da0e8a4cf889217106fa45b5f52609a6020527fcce26741946f801b25ce3c49451d2dd729b689d4d0d23ea57849f6c666bb5ee3546001600160a01b03166108ae565b6107c25f80516020612fd983398151915281565b6107e36119a1366004612dd3565b6124eb565b6107e36119b4366004612dec565b61251f565b6107c27f3c6dcff840f36f9818a73b67d9d00197362f63687bd52e3c277bd0ffb30dde3381565b6107e36119ee366004612dd3565b612573565b6107c27f690795c57e13eaf2526f76202b6799e9afdb069afca1e572f693953d013569d881565b6107e3611a28366004612dd3565b6125a7565b7f09dfa94a9be22222b511ecf509f49718fc08fbe3ada37a44d2022489eca3b44c5f52609b6020527fe98ed444639fcf7afa9e33a4ea67ac4155aa97d88f546111c8d1357c98dbca00546001600160a01b03166108ae565b7f2cf2377da51daa9c0d7e3f98c7532a67ee5e9398afad7b7db6e578b978a270945f5260986020527f0ccceacf55cd457ff25dca300775a2cb43db2c0b890d3ee063f4abba210c504f546107c2565b6107e3611ae2366004612e03565b6125db565b6107c27fdb5d1c2a9350ca010dcdf3953da11a9e8f7c5e2918cdfa65500e84e7fd4fde7d81565b7f9f919a2294d86593fbcec81ea71aa683cec51c78771c642f8894ba8f394970525f52609a6020527f99c8bd240e5bd2ee897b6a14ca3ca43a06f489dad5e38985ad188e67459dc6d7546001600160a01b03166108ae565b7f95bf18d68834a11aaae7b73ff6037326f163a81a7b5ea80cba96856ce2284fbd5f52609b6020527f20c8b2f4826823ac4cd62278270e8be9c7f63b9fe22e1f148f5369ec26bc69f4546001600160a01b03166108ae565b6107e3611bcc366004612dd3565b6125ff565b6107e3611bdf366004612dd3565b612677565b6107c27f46b41285bb7b8513ce3a9d95cdf6916699fb00b47326e8d3850be1b6186e034981565b7f3e4ded42f360c2e6b1251d584085ae1d9aa9cbed18687fac6b6aef8eed1c5ad35f52609a6020527fe298efc0f606c3be77912795055e173991a2c395633d4b0a06597a13b46e0c0b546001600160a01b03166108ae565b7ff935b8bf66b325637ad32ca875b588849cf4026791b79b4dc20623cd3dd36e205f52609a6020527f18d210dd586fe31598c73b0131261a1f7a576051e2667bbf5a4f8a01cf2f1392546001600160a01b03166108ae565b7f08593985ae1bebfb02f6c30105edffb176a6d87c9fad54c434bf9b58f67e81b65f5260976020527fb645ae2edae7c0716931b638cd9631a05f9a39fec3f15294f7f3af49f2f51ca8546107c2565b6107c27f5c00ec259bace293b50174e499c413ca897b4bcb54ed468b7e6bade51c6a9f9681565b5f805160206130598339815191525f5260986020525f80516020613039833981519152546107c2565b6107e3611d68366004612dd3565b6126ab565b6107c27f09dfa94a9be22222b511ecf509f49718fc08fbe3ada37a44d2022489eca3b44c81565b6107e3611da2366004612dd3565b6126de565b7f876943525608da6d95be5925fe6c4fe80e8622c8a76e7414f80e8ba210e0711c5f5260976020527fea561c0677f20715a0e74899b0381a0fa1265a58e9e02fb4a5a398d87555d1fe546107c2565b7ff822b1f0c3b886ce1cdf1c2a5317844145470db33b02c63cae4813f8c9b2dc175f5260976020527f9863915096f3522486953e53c4b97560d72679216b36fd98b4bdd4eca3a01eaa546107c2565b6107e3611e53366004612dec565b612712565b5f6001600160e01b03198216637965db0b60e01b1480611e8857506301ffc9a760e01b6001600160e01b03198316145b92915050565b5f611e9881612733565b611ec27fdb5d1c2a9350ca010dcdf3953da11a9e8f7c5e2918cdfa65500e84e7fd4fde7d83612740565b5050565b5f611ed081612733565b611ec27ff935b8bf66b325637ad32ca875b588849cf4026791b79b4dc20623cd3dd36e2083612740565b5f80516020612f99833981519152611f1181612733565b611f3b7f712c13b90acf399d7bc7625370ce37c64b5eba41011b0961a88c2ef1648870cf836127af565b611ec26127f6565b5f611f4d81612733565b611ec27fc5b1a6a0b843563e6a17ca90bc59d2315c523be427d0c9c2ba08d77ced4f46b183612740565b5f611f8181612733565b611ec27f8e96355022bb9b9f4d9d4e01fe2b58f45e78549c982c401c96f75f33c5de457e83612740565b5f611fb581612733565b611ec27f5be667ef1f4c6c279e2aa7e62595a1045043db6a43145cb438c6d36e7a3c3ed883612740565b5f611fe981612733565b611ec27fd7e49a298cb2719de62e5df1024257eed316db6337361b3a30d56a753240460783612740565b5f8281526065602052604090206001015461202d81612733565b61203783836129a9565b505050565b5f61204681612733565b611ec27f5c00ec259bace293b50174e499c413ca897b4bcb54ed468b7e6bade51c6a9f9683612740565b6001600160a01b03811633146120e55760405162461bcd60e51b815260206004820152602f60248201527f416363657373436f6e74726f6c3a2063616e206f6e6c792072656e6f756e636560448201526e103937b632b9903337b91039b2b63360891b60648201526084015b60405180910390fd5b611ec28282612a2e565b5f6120f981612733565b611ec27f3d88d1233771c5c30791fb6805b7f91424dae1e5a68a57da846ca7ff83c6402983612740565b5f61212d81612733565b611ec27fa4083e7a78dd898def03c51ce199cb4286b8828be4f6f46e04aec6176196747183612740565b5f61216181612733565b611ec27f690795c57e13eaf2526f76202b6799e9afdb069afca1e572f693953d013569d883612740565b5f611e887f523a704056dcd17bcf83bed8b68c59416dac1119be77755efe3bde0a64e46e0c83612425565b5f80516020612f998339815191526121cd81612733565b611ec27f46b41285bb7b8513ce3a9d95cdf6916699fb00b47326e8d3850be1b6186e0349836127af565b5f61220181612733565b611ec27fbd34382cd421c5250595893a4ed6cdb2125e6be7d5e0a9dbc469de5d583adfcf83612740565b5f61223581612733565b611f3b5f80516020612fb9833981519152836127af565b5f61225681612733565b611ec27f3e4ded42f360c2e6b1251d584085ae1d9aa9cbed18687fac6b6aef8eed1c5ad383612740565b5f611e885f80516020612f9983398151915283612425565b5f80516020612f998339815191526122af81612733565b611ec27f3c6dcff840f36f9818a73b67d9d00197362f63687bd52e3c277bd0ffb30dde33836127af565b5f6122e381612733565b611ec27fb134afa3abad633a84ab2d33dd5171f2b371e38b0f7bca001383aaf08ed6d2d183612740565b5f61231781612733565b611ec27f2cf2377da51daa9c0d7e3f98c7532a67ee5e9398afad7b7db6e578b978a27094836127af565b5f80516020612f9983398151915261235881612733565b611ec27f8567f5af844d68168987760a7ce1762804b9de703165fc50ce4fa85246016c9183612a94565b5f61238c81612733565b611ec27f95bf18d68834a11aaae7b73ff6037326f163a81a7b5ea80cba96856ce2284fbd83612afb565b5f80516020612f998339815191526123cd81612733565b611f3b5f80516020613059833981519152836127af565b5f80516020612f998339815191526123fb81612733565b611ec27fc54a7590fe6738d7a81f393c1cf5ab3e577c91781037d93a5a9f5ce44f19eb51836127af565b5f9182526065602090815260408084206001600160a01b0393909316845291905290205460ff1690565b5f61245981612733565b611ec27f8d4341681b282735dd0d55670ff8e0ad68a80cbfc2cee847065e9f771470f88f83612740565b5f61248d81612733565b611ec27fe5240448c78dfcff5bda4e4eed69ba9635df15d79da0e8a4cf889217106fa45b83612740565b5f6124c181612733565b611ec27f602490b12960e59ddb584affd1da6cd5692f4455c1ba0cc4e865af81e111ebe283612740565b5f6124f581612733565b611ec27f09dfa94a9be22222b511ecf509f49718fc08fbe3ada37a44d2022489eca3b44c83612afb565b7f523a704056dcd17bcf83bed8b68c59416dac1119be77755efe3bde0a64e46e0c61254981612733565b611ec27f6f8d0b773ad4970d3e7d47623dc9ce06a1b4fe833bf451d06a47e774f9acaa63836127af565b5f61257d81612733565b611ec27f76d62e541b8d573110ca3eb9003e96426f530422a76712d1356f6c6ce50541ca83612740565b5f6125b181612733565b611ec27f9f919a2294d86593fbcec81ea71aa683cec51c78771c642f8894ba8f3949705283612740565b5f828152606560205260409020600101546125f581612733565b6120378383612a2e565b5f61260981612733565b5f80516020612ff98339815191525f90815260996020527f15be86566e203c1f41b9ae149d9fbb01b2c14f503704423d739a6e3d2db5a9ee546001600160a01b03169061265690846129a9565b61266d5f80516020612ff983398151915284612a94565b6120375f82612a2e565b5f61268181612733565b611ec27f4c9466ca1bf288a7334a7494f09a0acc38ee31628eaf8c68b574b9f0ec22a9c183612740565b5f6126b581612733565b611ec27e665c1b06e0667c56a1ca1706b7573435d1b9162c6327b5d0ea1daeb491ad0d83612740565b5f6126e881612733565b611ec27f29384ec8473b541e7a7850226a4d1906a700f14cc394266ee08800ba62dc3af983612740565b5f61271c81612733565b611f3b5f80516020612fd9833981519152836127af565b61273d8133612b62565b50565b61274981612bbb565b5f828152609a602090815260409182902080546001600160a01b0319166001600160a01b0385169081179091558251858152918201527f5de40a806536a2029221dac2c8887ac9f11952fcc1ed3d7cfb4476dd5259b74091015b60405180910390a15050565b5f8281526098602090815260409182902083905581518481529081018390527f9094260c4234c0cb4c44e4a035abb5816b84e5505f9dc571c3ff397c4658163091016127a3565b5f805160206130598339815191525f5260986020525f80516020613039833981519152541580159061284a57505f80516020612fd98339815191525f5260986020525f805160206130198339815191525415155b801561289a575060986020527ffcacc1044a5a1b4eb9c058396306426a857813d37a4fb6ccf5a3adde30e0c914545f805160206130598339815191525f525f805160206130398339815191525411155b80156128ea575060986020527fd179a4a9329ee39fba707fd91c699ec0f088afc56731eb89ff424b873ac70844545f80516020612fd98339815191525f525f805160206130198339815191525411155b8015612927575060986020525f80516020613039833981519152545f80516020612fd98339815191525f525f805160206130198339815191525411155b801561298a575060986020527ffcacc1044a5a1b4eb9c058396306426a857813d37a4fb6ccf5a3adde30e0c914545f80516020612fb98339815191525f527fd179a4a9329ee39fba707fd91c699ec0f088afc56731eb89ff424b873ac708445410155b6129a75760405163e773e0a960e01b815260040160405180910390fd5b565b6129b38282612425565b611ec2575f8281526065602090815260408083206001600160a01b03851684529091529020805460ff191660011790556129ea3390565b6001600160a01b0316816001600160a01b0316837f2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d60405160405180910390a45050565b612a388282612425565b15611ec2575f8281526065602090815260408083206001600160a01b0385168085529252808320805460ff1916905551339285917ff6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b9190a45050565b612a9d81612bbb565b5f8281526099602090815260409182902080546001600160a01b0319166001600160a01b0385169081179091558251858152918201527fcbdd341876786c7241ad12a5ce5ea46739a4ce7b1587d0c216dfa655a98e50a691016127a3565b612b0481612bbb565b5f828152609b602090815260409182902080546001600160a01b0319166001600160a01b0385169081179091558251858152918201527f19aab10c6a9f5d648eaa15e2d515f8dfda570ee221e7c8cb9dc07694e68005bc91016127a3565b612b6c8282612425565b611ec257612b7981612be2565b612b84836020612bf4565b604051602001612b95929190612e77565b60408051601f198184030181529082905262461bcd60e51b82526120dc91600401612eeb565b6001600160a01b03811661273d5760405163d92e233d60e01b815260040160405180910390fd5b6060611e886001600160a01b03831660145b60605f612c02836002612f31565b612c0d906002612f48565b67ffffffffffffffff811115612c2557612c25612f5b565b6040519080825280601f01601f191660200182016040528015612c4f576020820181803683370190505b509050600360fc1b815f81518110612c6957612c69612f6f565b60200101906001600160f81b03191690815f1a905350600f60fb1b81600181518110612c9757612c97612f6f565b60200101906001600160f81b03191690815f1a9053505f612cb9846002612f31565b612cc4906001612f48565b90505b6001811115612d3b576f181899199a1a9b1b9c1cb0b131b232b360811b85600f1660108110612cf857612cf8612f6f565b1a60f81b828281518110612d0e57612d0e612f6f565b60200101906001600160f81b03191690815f1a90535060049490941c93612d3481612f83565b9050612cc7565b508315612d8a5760405162461bcd60e51b815260206004820181905260248201527f537472696e67733a20686578206c656e67746820696e73756666696369656e7460448201526064016120dc565b9392505050565b5f60208284031215612da1575f80fd5b81356001600160e01b031981168114612d8a575f80fd5b80356001600160a01b0381168114612dce575f80fd5b919050565b5f60208284031215612de3575f80fd5b612d8a82612db8565b5f60208284031215612dfc575f80fd5b5035919050565b5f8060408385031215612e14575f80fd5b82359150612e2460208401612db8565b90509250929050565b5f8060408385031215612e3e575f80fd5b612e4783612db8565b946020939093013593505050565b5f5b83811015612e6f578181015183820152602001612e57565b50505f910152565b7f416363657373436f6e74726f6c3a206163636f756e742000000000000000000081525f8351612eae816017850160208801612e55565b7001034b99036b4b9b9b4b733903937b6329607d1b6017918401918201528351612edf816028840160208801612e55565b01602801949350505050565b602081525f8251806020840152612f09816040850160208701612e55565b601f01601f19169190910160400192915050565b634e487b7160e01b5f52601160045260245ffd5b8082028115828204841417611e8857611e88612f1d565b80820180821115611e8857611e88612f1d565b634e487b7160e01b5f52604160045260245ffd5b634e487b7160e01b5f52603260045260245ffd5b5f81612f9157612f91612f1d565b505f19019056feaf290d8680820aad922855f39b306097b20e28774d6c1ad35a20325630c3a02c1c2fe98ddbbbffbcf7735c7446ffcddb5ccd2a4ec2ace0f7d90f73e9ff13fcc7b18278bb399a7088b8b0b26f4896d5ebaba4497c611bbe9d43abe92d9a1fe83ddf8b4c520ffe197c5343c6f5aec59570151ef9a492f2c624fd45ddde6135ec428f1b9b075a455aa4e85ab4edea73c8fe6d4e2e5e4c6675d6135fefdca5e95a258489bc07817c82dd59579d43388f707a6a0a4a614b58e7df61bb06baec0de2c1fa5a84fed05ba4c93fcc5ba1f4ad010e3bef3e6394b367aa10b3ec01997375cca2646970667358221220d0bbc42566fb90656cf788bbeb4cb6ef98851a82ee8814bf66f60651001b822764736f6c634300081400339094260c4234c0cb4c44e4a035abb5816b84e5505f9dc571c3ff397c46581630",
}

// StaderConfigABI is the input ABI used to generate the binding from.
// Deprecated: Use StaderConfigMetaData.ABI instead.
var StaderConfigABI = StaderConfigMetaData.ABI

// StaderConfigBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use StaderConfigMetaData.Bin instead.
var StaderConfigBin = StaderConfigMetaData.Bin

// DeployStaderConfig deploys a new Ethereum contract, binding an instance of StaderConfig to it.
func DeployStaderConfig(auth *bind.TransactOpts, backend bind.ContractBackend, _admin common.Address, _ethDepositContract common.Address) (common.Address, *types.Transaction, *StaderConfig, error) {
	parsed, err := StaderConfigMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(StaderConfigBin), backend, _admin, _ethDepositContract)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &StaderConfig{StaderConfigCaller: StaderConfigCaller{contract: contract}, StaderConfigTransactor: StaderConfigTransactor{contract: contract}, StaderConfigFilterer: StaderConfigFilterer{contract: contract}}, nil
}

// StaderConfig is an auto generated Go binding around an Ethereum contract.
type StaderConfig struct {
	StaderConfigCaller     // Read-only binding to the contract
	StaderConfigTransactor // Write-only binding to the contract
	StaderConfigFilterer   // Log filterer for contract events
}

// StaderConfigCaller is an auto generated read-only Go binding around an Ethereum contract.
type StaderConfigCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StaderConfigTransactor is an auto generated write-only Go binding around an Ethereum contract.
type StaderConfigTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StaderConfigFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type StaderConfigFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StaderConfigSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type StaderConfigSession struct {
	Contract     *StaderConfig     // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// StaderConfigCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type StaderConfigCallerSession struct {
	Contract *StaderConfigCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts       // Call options to use throughout this session
}

// StaderConfigTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type StaderConfigTransactorSession struct {
	Contract     *StaderConfigTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts       // Transaction auth options to use throughout this session
}

// StaderConfigRaw is an auto generated low-level Go binding around an Ethereum contract.
type StaderConfigRaw struct {
	Contract *StaderConfig // Generic contract binding to access the raw methods on
}

// StaderConfigCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type StaderConfigCallerRaw struct {
	Contract *StaderConfigCaller // Generic read-only contract binding to access the raw methods on
}

// StaderConfigTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type StaderConfigTransactorRaw struct {
	Contract *StaderConfigTransactor // Generic write-only contract binding to access the raw methods on
}

// NewStaderConfig creates a new instance of StaderConfig, bound to a specific deployed contract.
func NewStaderConfig(address common.Address, backend bind.ContractBackend) (*StaderConfig, error) {
	contract, err := bindStaderConfig(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &StaderConfig{StaderConfigCaller: StaderConfigCaller{contract: contract}, StaderConfigTransactor: StaderConfigTransactor{contract: contract}, StaderConfigFilterer: StaderConfigFilterer{contract: contract}}, nil
}

// NewStaderConfigCaller creates a new read-only instance of StaderConfig, bound to a specific deployed contract.
func NewStaderConfigCaller(address common.Address, caller bind.ContractCaller) (*StaderConfigCaller, error) {
	contract, err := bindStaderConfig(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &StaderConfigCaller{contract: contract}, nil
}

// NewStaderConfigTransactor creates a new write-only instance of StaderConfig, bound to a specific deployed contract.
func NewStaderConfigTransactor(address common.Address, transactor bind.ContractTransactor) (*StaderConfigTransactor, error) {
	contract, err := bindStaderConfig(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &StaderConfigTransactor{contract: contract}, nil
}

// NewStaderConfigFilterer creates a new log filterer instance of StaderConfig, bound to a specific deployed contract.
func NewStaderConfigFilterer(address common.Address, filterer bind.ContractFilterer) (*StaderConfigFilterer, error) {
	contract, err := bindStaderConfig(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &StaderConfigFilterer{contract: contract}, nil
}

// bindStaderConfig binds a generic wrapper to an already deployed contract.
func bindStaderConfig(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := StaderConfigMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_StaderConfig *StaderConfigRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _StaderConfig.Contract.StaderConfigCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_StaderConfig *StaderConfigRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _StaderConfig.Contract.StaderConfigTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_StaderConfig *StaderConfigRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _StaderConfig.Contract.StaderConfigTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_StaderConfig *StaderConfigCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _StaderConfig.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_StaderConfig *StaderConfigTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _StaderConfig.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_StaderConfig *StaderConfigTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _StaderConfig.Contract.contract.Transact(opts, method, params...)
}

// ADMIN is a free data retrieval call binding the contract method 0x2a0acc6a.
//
// Solidity: function ADMIN() view returns(bytes32)
func (_StaderConfig *StaderConfigCaller) ADMIN(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _StaderConfig.contract.Call(opts, &out, "ADMIN")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// ADMIN is a free data retrieval call binding the contract method 0x2a0acc6a.
//
// Solidity: function ADMIN() view returns(bytes32)
func (_StaderConfig *StaderConfigSession) ADMIN() ([32]byte, error) {
	return _StaderConfig.Contract.ADMIN(&_StaderConfig.CallOpts)
}

// ADMIN is a free data retrieval call binding the contract method 0x2a0acc6a.
//
// Solidity: function ADMIN() view returns(bytes32)
func (_StaderConfig *StaderConfigCallerSession) ADMIN() ([32]byte, error) {
	return _StaderConfig.Contract.ADMIN(&_StaderConfig.CallOpts)
}

// AUCTIONCONTRACT is a free data retrieval call binding the contract method 0xb11c699d.
//
// Solidity: function AUCTION_CONTRACT() view returns(bytes32)
func (_StaderConfig *StaderConfigCaller) AUCTIONCONTRACT(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _StaderConfig.contract.Call(opts, &out, "AUCTION_CONTRACT")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// AUCTIONCONTRACT is a free data retrieval call binding the contract method 0xb11c699d.
//
// Solidity: function AUCTION_CONTRACT() view returns(bytes32)
func (_StaderConfig *StaderConfigSession) AUCTIONCONTRACT() ([32]byte, error) {
	return _StaderConfig.Contract.AUCTIONCONTRACT(&_StaderConfig.CallOpts)
}

// AUCTIONCONTRACT is a free data retrieval call binding the contract method 0xb11c699d.
//
// Solidity: function AUCTION_CONTRACT() view returns(bytes32)
func (_StaderConfig *StaderConfigCallerSession) AUCTIONCONTRACT() ([32]byte, error) {
	return _StaderConfig.Contract.AUCTIONCONTRACT(&_StaderConfig.CallOpts)
}

// DECIMALS is a free data retrieval call binding the contract method 0x2e0f2625.
//
// Solidity: function DECIMALS() view returns(bytes32)
func (_StaderConfig *StaderConfigCaller) DECIMALS(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _StaderConfig.contract.Call(opts, &out, "DECIMALS")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// DECIMALS is a free data retrieval call binding the contract method 0x2e0f2625.
//
// Solidity: function DECIMALS() view returns(bytes32)
func (_StaderConfig *StaderConfigSession) DECIMALS() ([32]byte, error) {
	return _StaderConfig.Contract.DECIMALS(&_StaderConfig.CallOpts)
}

// DECIMALS is a free data retrieval call binding the contract method 0x2e0f2625.
//
// Solidity: function DECIMALS() view returns(bytes32)
func (_StaderConfig *StaderConfigCallerSession) DECIMALS() ([32]byte, error) {
	return _StaderConfig.Contract.DECIMALS(&_StaderConfig.CallOpts)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_StaderConfig *StaderConfigCaller) DEFAULTADMINROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _StaderConfig.contract.Call(opts, &out, "DEFAULT_ADMIN_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_StaderConfig *StaderConfigSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _StaderConfig.Contract.DEFAULTADMINROLE(&_StaderConfig.CallOpts)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_StaderConfig *StaderConfigCallerSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _StaderConfig.Contract.DEFAULTADMINROLE(&_StaderConfig.CallOpts)
}

// ETHXSUPPLYPORFEED is a free data retrieval call binding the contract method 0x2a9cc2c4.
//
// Solidity: function ETHX_SUPPLY_POR_FEED() view returns(bytes32)
func (_StaderConfig *StaderConfigCaller) ETHXSUPPLYPORFEED(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _StaderConfig.contract.Call(opts, &out, "ETHX_SUPPLY_POR_FEED")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// ETHXSUPPLYPORFEED is a free data retrieval call binding the contract method 0x2a9cc2c4.
//
// Solidity: function ETHX_SUPPLY_POR_FEED() view returns(bytes32)
func (_StaderConfig *StaderConfigSession) ETHXSUPPLYPORFEED() ([32]byte, error) {
	return _StaderConfig.Contract.ETHXSUPPLYPORFEED(&_StaderConfig.CallOpts)
}

// ETHXSUPPLYPORFEED is a free data retrieval call binding the contract method 0x2a9cc2c4.
//
// Solidity: function ETHX_SUPPLY_POR_FEED() view returns(bytes32)
func (_StaderConfig *StaderConfigCallerSession) ETHXSUPPLYPORFEED() ([32]byte, error) {
	return _StaderConfig.Contract.ETHXSUPPLYPORFEED(&_StaderConfig.CallOpts)
}

// ETHBALANCEPORFEED is a free data retrieval call binding the contract method 0xc60470d3.
//
// Solidity: function ETH_BALANCE_POR_FEED() view returns(bytes32)
func (_StaderConfig *StaderConfigCaller) ETHBALANCEPORFEED(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _StaderConfig.contract.Call(opts, &out, "ETH_BALANCE_POR_FEED")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// ETHBALANCEPORFEED is a free data retrieval call binding the contract method 0xc60470d3.
//
// Solidity: function ETH_BALANCE_POR_FEED() view returns(bytes32)
func (_StaderConfig *StaderConfigSession) ETHBALANCEPORFEED() ([32]byte, error) {
	return _StaderConfig.Contract.ETHBALANCEPORFEED(&_StaderConfig.CallOpts)
}

// ETHBALANCEPORFEED is a free data retrieval call binding the contract method 0xc60470d3.
//
// Solidity: function ETH_BALANCE_POR_FEED() view returns(bytes32)
func (_StaderConfig *StaderConfigCallerSession) ETHBALANCEPORFEED() ([32]byte, error) {
	return _StaderConfig.Contract.ETHBALANCEPORFEED(&_StaderConfig.CallOpts)
}

// ETHDEPOSITCONTRACT is a free data retrieval call binding the contract method 0x77e8a0c3.
//
// Solidity: function ETH_DEPOSIT_CONTRACT() view returns(bytes32)
func (_StaderConfig *StaderConfigCaller) ETHDEPOSITCONTRACT(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _StaderConfig.contract.Call(opts, &out, "ETH_DEPOSIT_CONTRACT")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// ETHDEPOSITCONTRACT is a free data retrieval call binding the contract method 0x77e8a0c3.
//
// Solidity: function ETH_DEPOSIT_CONTRACT() view returns(bytes32)
func (_StaderConfig *StaderConfigSession) ETHDEPOSITCONTRACT() ([32]byte, error) {
	return _StaderConfig.Contract.ETHDEPOSITCONTRACT(&_StaderConfig.CallOpts)
}

// ETHDEPOSITCONTRACT is a free data retrieval call binding the contract method 0x77e8a0c3.
//
// Solidity: function ETH_DEPOSIT_CONTRACT() view returns(bytes32)
func (_StaderConfig *StaderConfigCallerSession) ETHDEPOSITCONTRACT() ([32]byte, error) {
	return _StaderConfig.Contract.ETHDEPOSITCONTRACT(&_StaderConfig.CallOpts)
}

// ETHPERNODE is a free data retrieval call binding the contract method 0x67dcf134.
//
// Solidity: function ETH_PER_NODE() view returns(bytes32)
func (_StaderConfig *StaderConfigCaller) ETHPERNODE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _StaderConfig.contract.Call(opts, &out, "ETH_PER_NODE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// ETHPERNODE is a free data retrieval call binding the contract method 0x67dcf134.
//
// Solidity: function ETH_PER_NODE() view returns(bytes32)
func (_StaderConfig *StaderConfigSession) ETHPERNODE() ([32]byte, error) {
	return _StaderConfig.Contract.ETHPERNODE(&_StaderConfig.CallOpts)
}

// ETHPERNODE is a free data retrieval call binding the contract method 0x67dcf134.
//
// Solidity: function ETH_PER_NODE() view returns(bytes32)
func (_StaderConfig *StaderConfigCallerSession) ETHPERNODE() ([32]byte, error) {
	return _StaderConfig.Contract.ETHPERNODE(&_StaderConfig.CallOpts)
}

// ETHx is a free data retrieval call binding the contract method 0xf6c278c1.
//
// Solidity: function ETHx() view returns(bytes32)
func (_StaderConfig *StaderConfigCaller) ETHx(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _StaderConfig.contract.Call(opts, &out, "ETHx")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// ETHx is a free data retrieval call binding the contract method 0xf6c278c1.
//
// Solidity: function ETHx() view returns(bytes32)
func (_StaderConfig *StaderConfigSession) ETHx() ([32]byte, error) {
	return _StaderConfig.Contract.ETHx(&_StaderConfig.CallOpts)
}

// ETHx is a free data retrieval call binding the contract method 0xf6c278c1.
//
// Solidity: function ETHx() view returns(bytes32)
func (_StaderConfig *StaderConfigCallerSession) ETHx() ([32]byte, error) {
	return _StaderConfig.Contract.ETHx(&_StaderConfig.CallOpts)
}

// FULLDEPOSITSIZE is a free data retrieval call binding the contract method 0x792c8cc3.
//
// Solidity: function FULL_DEPOSIT_SIZE() view returns(bytes32)
func (_StaderConfig *StaderConfigCaller) FULLDEPOSITSIZE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _StaderConfig.contract.Call(opts, &out, "FULL_DEPOSIT_SIZE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// FULLDEPOSITSIZE is a free data retrieval call binding the contract method 0x792c8cc3.
//
// Solidity: function FULL_DEPOSIT_SIZE() view returns(bytes32)
func (_StaderConfig *StaderConfigSession) FULLDEPOSITSIZE() ([32]byte, error) {
	return _StaderConfig.Contract.FULLDEPOSITSIZE(&_StaderConfig.CallOpts)
}

// FULLDEPOSITSIZE is a free data retrieval call binding the contract method 0x792c8cc3.
//
// Solidity: function FULL_DEPOSIT_SIZE() view returns(bytes32)
func (_StaderConfig *StaderConfigCallerSession) FULLDEPOSITSIZE() ([32]byte, error) {
	return _StaderConfig.Contract.FULLDEPOSITSIZE(&_StaderConfig.CallOpts)
}

// MANAGER is a free data retrieval call binding the contract method 0x1b2df850.
//
// Solidity: function MANAGER() view returns(bytes32)
func (_StaderConfig *StaderConfigCaller) MANAGER(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _StaderConfig.contract.Call(opts, &out, "MANAGER")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// MANAGER is a free data retrieval call binding the contract method 0x1b2df850.
//
// Solidity: function MANAGER() view returns(bytes32)
func (_StaderConfig *StaderConfigSession) MANAGER() ([32]byte, error) {
	return _StaderConfig.Contract.MANAGER(&_StaderConfig.CallOpts)
}

// MANAGER is a free data retrieval call binding the contract method 0x1b2df850.
//
// Solidity: function MANAGER() view returns(bytes32)
func (_StaderConfig *StaderConfigCallerSession) MANAGER() ([32]byte, error) {
	return _StaderConfig.Contract.MANAGER(&_StaderConfig.CallOpts)
}

// MAXDEPOSITAMOUNT is a free data retrieval call binding the contract method 0x4c34a982.
//
// Solidity: function MAX_DEPOSIT_AMOUNT() view returns(bytes32)
func (_StaderConfig *StaderConfigCaller) MAXDEPOSITAMOUNT(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _StaderConfig.contract.Call(opts, &out, "MAX_DEPOSIT_AMOUNT")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// MAXDEPOSITAMOUNT is a free data retrieval call binding the contract method 0x4c34a982.
//
// Solidity: function MAX_DEPOSIT_AMOUNT() view returns(bytes32)
func (_StaderConfig *StaderConfigSession) MAXDEPOSITAMOUNT() ([32]byte, error) {
	return _StaderConfig.Contract.MAXDEPOSITAMOUNT(&_StaderConfig.CallOpts)
}

// MAXDEPOSITAMOUNT is a free data retrieval call binding the contract method 0x4c34a982.
//
// Solidity: function MAX_DEPOSIT_AMOUNT() view returns(bytes32)
func (_StaderConfig *StaderConfigCallerSession) MAXDEPOSITAMOUNT() ([32]byte, error) {
	return _StaderConfig.Contract.MAXDEPOSITAMOUNT(&_StaderConfig.CallOpts)
}

// MAXWITHDRAWAMOUNT is a free data retrieval call binding the contract method 0x44ba0ea2.
//
// Solidity: function MAX_WITHDRAW_AMOUNT() view returns(bytes32)
func (_StaderConfig *StaderConfigCaller) MAXWITHDRAWAMOUNT(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _StaderConfig.contract.Call(opts, &out, "MAX_WITHDRAW_AMOUNT")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// MAXWITHDRAWAMOUNT is a free data retrieval call binding the contract method 0x44ba0ea2.
//
// Solidity: function MAX_WITHDRAW_AMOUNT() view returns(bytes32)
func (_StaderConfig *StaderConfigSession) MAXWITHDRAWAMOUNT() ([32]byte, error) {
	return _StaderConfig.Contract.MAXWITHDRAWAMOUNT(&_StaderConfig.CallOpts)
}

// MAXWITHDRAWAMOUNT is a free data retrieval call binding the contract method 0x44ba0ea2.
//
// Solidity: function MAX_WITHDRAW_AMOUNT() view returns(bytes32)
func (_StaderConfig *StaderConfigCallerSession) MAXWITHDRAWAMOUNT() ([32]byte, error) {
	return _StaderConfig.Contract.MAXWITHDRAWAMOUNT(&_StaderConfig.CallOpts)
}

// MINBLOCKDELAYTOFINALIZEWITHDRAWREQUEST is a free data retrieval call binding the contract method 0x6176bbde.
//
// Solidity: function MIN_BLOCK_DELAY_TO_FINALIZE_WITHDRAW_REQUEST() view returns(bytes32)
func (_StaderConfig *StaderConfigCaller) MINBLOCKDELAYTOFINALIZEWITHDRAWREQUEST(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _StaderConfig.contract.Call(opts, &out, "MIN_BLOCK_DELAY_TO_FINALIZE_WITHDRAW_REQUEST")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// MINBLOCKDELAYTOFINALIZEWITHDRAWREQUEST is a free data retrieval call binding the contract method 0x6176bbde.
//
// Solidity: function MIN_BLOCK_DELAY_TO_FINALIZE_WITHDRAW_REQUEST() view returns(bytes32)
func (_StaderConfig *StaderConfigSession) MINBLOCKDELAYTOFINALIZEWITHDRAWREQUEST() ([32]byte, error) {
	return _StaderConfig.Contract.MINBLOCKDELAYTOFINALIZEWITHDRAWREQUEST(&_StaderConfig.CallOpts)
}

// MINBLOCKDELAYTOFINALIZEWITHDRAWREQUEST is a free data retrieval call binding the contract method 0x6176bbde.
//
// Solidity: function MIN_BLOCK_DELAY_TO_FINALIZE_WITHDRAW_REQUEST() view returns(bytes32)
func (_StaderConfig *StaderConfigCallerSession) MINBLOCKDELAYTOFINALIZEWITHDRAWREQUEST() ([32]byte, error) {
	return _StaderConfig.Contract.MINBLOCKDELAYTOFINALIZEWITHDRAWREQUEST(&_StaderConfig.CallOpts)
}

// MINDEPOSITAMOUNT is a free data retrieval call binding the contract method 0x1ea30fef.
//
// Solidity: function MIN_DEPOSIT_AMOUNT() view returns(bytes32)
func (_StaderConfig *StaderConfigCaller) MINDEPOSITAMOUNT(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _StaderConfig.contract.Call(opts, &out, "MIN_DEPOSIT_AMOUNT")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// MINDEPOSITAMOUNT is a free data retrieval call binding the contract method 0x1ea30fef.
//
// Solidity: function MIN_DEPOSIT_AMOUNT() view returns(bytes32)
func (_StaderConfig *StaderConfigSession) MINDEPOSITAMOUNT() ([32]byte, error) {
	return _StaderConfig.Contract.MINDEPOSITAMOUNT(&_StaderConfig.CallOpts)
}

// MINDEPOSITAMOUNT is a free data retrieval call binding the contract method 0x1ea30fef.
//
// Solidity: function MIN_DEPOSIT_AMOUNT() view returns(bytes32)
func (_StaderConfig *StaderConfigCallerSession) MINDEPOSITAMOUNT() ([32]byte, error) {
	return _StaderConfig.Contract.MINDEPOSITAMOUNT(&_StaderConfig.CallOpts)
}

// MINWITHDRAWAMOUNT is a free data retrieval call binding the contract method 0xb6857844.
//
// Solidity: function MIN_WITHDRAW_AMOUNT() view returns(bytes32)
func (_StaderConfig *StaderConfigCaller) MINWITHDRAWAMOUNT(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _StaderConfig.contract.Call(opts, &out, "MIN_WITHDRAW_AMOUNT")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// MINWITHDRAWAMOUNT is a free data retrieval call binding the contract method 0xb6857844.
//
// Solidity: function MIN_WITHDRAW_AMOUNT() view returns(bytes32)
func (_StaderConfig *StaderConfigSession) MINWITHDRAWAMOUNT() ([32]byte, error) {
	return _StaderConfig.Contract.MINWITHDRAWAMOUNT(&_StaderConfig.CallOpts)
}

// MINWITHDRAWAMOUNT is a free data retrieval call binding the contract method 0xb6857844.
//
// Solidity: function MIN_WITHDRAW_AMOUNT() view returns(bytes32)
func (_StaderConfig *StaderConfigCallerSession) MINWITHDRAWAMOUNT() ([32]byte, error) {
	return _StaderConfig.Contract.MINWITHDRAWAMOUNT(&_StaderConfig.CallOpts)
}

// NODEELREWARDVAULTIMPLEMENTATION is a free data retrieval call binding the contract method 0x0bdf3166.
//
// Solidity: function NODE_EL_REWARD_VAULT_IMPLEMENTATION() view returns(bytes32)
func (_StaderConfig *StaderConfigCaller) NODEELREWARDVAULTIMPLEMENTATION(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _StaderConfig.contract.Call(opts, &out, "NODE_EL_REWARD_VAULT_IMPLEMENTATION")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// NODEELREWARDVAULTIMPLEMENTATION is a free data retrieval call binding the contract method 0x0bdf3166.
//
// Solidity: function NODE_EL_REWARD_VAULT_IMPLEMENTATION() view returns(bytes32)
func (_StaderConfig *StaderConfigSession) NODEELREWARDVAULTIMPLEMENTATION() ([32]byte, error) {
	return _StaderConfig.Contract.NODEELREWARDVAULTIMPLEMENTATION(&_StaderConfig.CallOpts)
}

// NODEELREWARDVAULTIMPLEMENTATION is a free data retrieval call binding the contract method 0x0bdf3166.
//
// Solidity: function NODE_EL_REWARD_VAULT_IMPLEMENTATION() view returns(bytes32)
func (_StaderConfig *StaderConfigCallerSession) NODEELREWARDVAULTIMPLEMENTATION() ([32]byte, error) {
	return _StaderConfig.Contract.NODEELREWARDVAULTIMPLEMENTATION(&_StaderConfig.CallOpts)
}

// OPERATOR is a free data retrieval call binding the contract method 0x983d2737.
//
// Solidity: function OPERATOR() view returns(bytes32)
func (_StaderConfig *StaderConfigCaller) OPERATOR(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _StaderConfig.contract.Call(opts, &out, "OPERATOR")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// OPERATOR is a free data retrieval call binding the contract method 0x983d2737.
//
// Solidity: function OPERATOR() view returns(bytes32)
func (_StaderConfig *StaderConfigSession) OPERATOR() ([32]byte, error) {
	return _StaderConfig.Contract.OPERATOR(&_StaderConfig.CallOpts)
}

// OPERATOR is a free data retrieval call binding the contract method 0x983d2737.
//
// Solidity: function OPERATOR() view returns(bytes32)
func (_StaderConfig *StaderConfigCallerSession) OPERATOR() ([32]byte, error) {
	return _StaderConfig.Contract.OPERATOR(&_StaderConfig.CallOpts)
}

// OPERATORMAXNAMELENGTH is a free data retrieval call binding the contract method 0x5455e472.
//
// Solidity: function OPERATOR_MAX_NAME_LENGTH() view returns(bytes32)
func (_StaderConfig *StaderConfigCaller) OPERATORMAXNAMELENGTH(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _StaderConfig.contract.Call(opts, &out, "OPERATOR_MAX_NAME_LENGTH")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// OPERATORMAXNAMELENGTH is a free data retrieval call binding the contract method 0x5455e472.
//
// Solidity: function OPERATOR_MAX_NAME_LENGTH() view returns(bytes32)
func (_StaderConfig *StaderConfigSession) OPERATORMAXNAMELENGTH() ([32]byte, error) {
	return _StaderConfig.Contract.OPERATORMAXNAMELENGTH(&_StaderConfig.CallOpts)
}

// OPERATORMAXNAMELENGTH is a free data retrieval call binding the contract method 0x5455e472.
//
// Solidity: function OPERATOR_MAX_NAME_LENGTH() view returns(bytes32)
func (_StaderConfig *StaderConfigCallerSession) OPERATORMAXNAMELENGTH() ([32]byte, error) {
	return _StaderConfig.Contract.OPERATORMAXNAMELENGTH(&_StaderConfig.CallOpts)
}

// OPERATORREWARDCOLLECTOR is a free data retrieval call binding the contract method 0x79175a74.
//
// Solidity: function OPERATOR_REWARD_COLLECTOR() view returns(bytes32)
func (_StaderConfig *StaderConfigCaller) OPERATORREWARDCOLLECTOR(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _StaderConfig.contract.Call(opts, &out, "OPERATOR_REWARD_COLLECTOR")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// OPERATORREWARDCOLLECTOR is a free data retrieval call binding the contract method 0x79175a74.
//
// Solidity: function OPERATOR_REWARD_COLLECTOR() view returns(bytes32)
func (_StaderConfig *StaderConfigSession) OPERATORREWARDCOLLECTOR() ([32]byte, error) {
	return _StaderConfig.Contract.OPERATORREWARDCOLLECTOR(&_StaderConfig.CallOpts)
}

// OPERATORREWARDCOLLECTOR is a free data retrieval call binding the contract method 0x79175a74.
//
// Solidity: function OPERATOR_REWARD_COLLECTOR() view returns(bytes32)
func (_StaderConfig *StaderConfigCallerSession) OPERATORREWARDCOLLECTOR() ([32]byte, error) {
	return _StaderConfig.Contract.OPERATORREWARDCOLLECTOR(&_StaderConfig.CallOpts)
}

// PENALTYCONTRACT is a free data retrieval call binding the contract method 0x1bf6a41c.
//
// Solidity: function PENALTY_CONTRACT() view returns(bytes32)
func (_StaderConfig *StaderConfigCaller) PENALTYCONTRACT(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _StaderConfig.contract.Call(opts, &out, "PENALTY_CONTRACT")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// PENALTYCONTRACT is a free data retrieval call binding the contract method 0x1bf6a41c.
//
// Solidity: function PENALTY_CONTRACT() view returns(bytes32)
func (_StaderConfig *StaderConfigSession) PENALTYCONTRACT() ([32]byte, error) {
	return _StaderConfig.Contract.PENALTYCONTRACT(&_StaderConfig.CallOpts)
}

// PENALTYCONTRACT is a free data retrieval call binding the contract method 0x1bf6a41c.
//
// Solidity: function PENALTY_CONTRACT() view returns(bytes32)
func (_StaderConfig *StaderConfigCallerSession) PENALTYCONTRACT() ([32]byte, error) {
	return _StaderConfig.Contract.PENALTYCONTRACT(&_StaderConfig.CallOpts)
}

// PERMISSIONEDNODEREGISTRY is a free data retrieval call binding the contract method 0x4191e0fe.
//
// Solidity: function PERMISSIONED_NODE_REGISTRY() view returns(bytes32)
func (_StaderConfig *StaderConfigCaller) PERMISSIONEDNODEREGISTRY(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _StaderConfig.contract.Call(opts, &out, "PERMISSIONED_NODE_REGISTRY")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// PERMISSIONEDNODEREGISTRY is a free data retrieval call binding the contract method 0x4191e0fe.
//
// Solidity: function PERMISSIONED_NODE_REGISTRY() view returns(bytes32)
func (_StaderConfig *StaderConfigSession) PERMISSIONEDNODEREGISTRY() ([32]byte, error) {
	return _StaderConfig.Contract.PERMISSIONEDNODEREGISTRY(&_StaderConfig.CallOpts)
}

// PERMISSIONEDNODEREGISTRY is a free data retrieval call binding the contract method 0x4191e0fe.
//
// Solidity: function PERMISSIONED_NODE_REGISTRY() view returns(bytes32)
func (_StaderConfig *StaderConfigCallerSession) PERMISSIONEDNODEREGISTRY() ([32]byte, error) {
	return _StaderConfig.Contract.PERMISSIONEDNODEREGISTRY(&_StaderConfig.CallOpts)
}

// PERMISSIONEDPOOL is a free data retrieval call binding the contract method 0x52112bd3.
//
// Solidity: function PERMISSIONED_POOL() view returns(bytes32)
func (_StaderConfig *StaderConfigCaller) PERMISSIONEDPOOL(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _StaderConfig.contract.Call(opts, &out, "PERMISSIONED_POOL")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// PERMISSIONEDPOOL is a free data retrieval call binding the contract method 0x52112bd3.
//
// Solidity: function PERMISSIONED_POOL() view returns(bytes32)
func (_StaderConfig *StaderConfigSession) PERMISSIONEDPOOL() ([32]byte, error) {
	return _StaderConfig.Contract.PERMISSIONEDPOOL(&_StaderConfig.CallOpts)
}

// PERMISSIONEDPOOL is a free data retrieval call binding the contract method 0x52112bd3.
//
// Solidity: function PERMISSIONED_POOL() view returns(bytes32)
func (_StaderConfig *StaderConfigCallerSession) PERMISSIONEDPOOL() ([32]byte, error) {
	return _StaderConfig.Contract.PERMISSIONEDPOOL(&_StaderConfig.CallOpts)
}

// PERMISSIONEDSOCIALIZINGPOOL is a free data retrieval call binding the contract method 0x12020075.
//
// Solidity: function PERMISSIONED_SOCIALIZING_POOL() view returns(bytes32)
func (_StaderConfig *StaderConfigCaller) PERMISSIONEDSOCIALIZINGPOOL(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _StaderConfig.contract.Call(opts, &out, "PERMISSIONED_SOCIALIZING_POOL")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// PERMISSIONEDSOCIALIZINGPOOL is a free data retrieval call binding the contract method 0x12020075.
//
// Solidity: function PERMISSIONED_SOCIALIZING_POOL() view returns(bytes32)
func (_StaderConfig *StaderConfigSession) PERMISSIONEDSOCIALIZINGPOOL() ([32]byte, error) {
	return _StaderConfig.Contract.PERMISSIONEDSOCIALIZINGPOOL(&_StaderConfig.CallOpts)
}

// PERMISSIONEDSOCIALIZINGPOOL is a free data retrieval call binding the contract method 0x12020075.
//
// Solidity: function PERMISSIONED_SOCIALIZING_POOL() view returns(bytes32)
func (_StaderConfig *StaderConfigCallerSession) PERMISSIONEDSOCIALIZINGPOOL() ([32]byte, error) {
	return _StaderConfig.Contract.PERMISSIONEDSOCIALIZINGPOOL(&_StaderConfig.CallOpts)
}

// PERMISSIONLESSNODEREGISTRY is a free data retrieval call binding the contract method 0x152a91da.
//
// Solidity: function PERMISSIONLESS_NODE_REGISTRY() view returns(bytes32)
func (_StaderConfig *StaderConfigCaller) PERMISSIONLESSNODEREGISTRY(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _StaderConfig.contract.Call(opts, &out, "PERMISSIONLESS_NODE_REGISTRY")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// PERMISSIONLESSNODEREGISTRY is a free data retrieval call binding the contract method 0x152a91da.
//
// Solidity: function PERMISSIONLESS_NODE_REGISTRY() view returns(bytes32)
func (_StaderConfig *StaderConfigSession) PERMISSIONLESSNODEREGISTRY() ([32]byte, error) {
	return _StaderConfig.Contract.PERMISSIONLESSNODEREGISTRY(&_StaderConfig.CallOpts)
}

// PERMISSIONLESSNODEREGISTRY is a free data retrieval call binding the contract method 0x152a91da.
//
// Solidity: function PERMISSIONLESS_NODE_REGISTRY() view returns(bytes32)
func (_StaderConfig *StaderConfigCallerSession) PERMISSIONLESSNODEREGISTRY() ([32]byte, error) {
	return _StaderConfig.Contract.PERMISSIONLESSNODEREGISTRY(&_StaderConfig.CallOpts)
}

// PERMISSIONLESSPOOL is a free data retrieval call binding the contract method 0x7a87fa0b.
//
// Solidity: function PERMISSIONLESS_POOL() view returns(bytes32)
func (_StaderConfig *StaderConfigCaller) PERMISSIONLESSPOOL(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _StaderConfig.contract.Call(opts, &out, "PERMISSIONLESS_POOL")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// PERMISSIONLESSPOOL is a free data retrieval call binding the contract method 0x7a87fa0b.
//
// Solidity: function PERMISSIONLESS_POOL() view returns(bytes32)
func (_StaderConfig *StaderConfigSession) PERMISSIONLESSPOOL() ([32]byte, error) {
	return _StaderConfig.Contract.PERMISSIONLESSPOOL(&_StaderConfig.CallOpts)
}

// PERMISSIONLESSPOOL is a free data retrieval call binding the contract method 0x7a87fa0b.
//
// Solidity: function PERMISSIONLESS_POOL() view returns(bytes32)
func (_StaderConfig *StaderConfigCallerSession) PERMISSIONLESSPOOL() ([32]byte, error) {
	return _StaderConfig.Contract.PERMISSIONLESSPOOL(&_StaderConfig.CallOpts)
}

// PERMISSIONLESSSOCIALIZINGPOOL is a free data retrieval call binding the contract method 0x3b6bcca0.
//
// Solidity: function PERMISSIONLESS_SOCIALIZING_POOL() view returns(bytes32)
func (_StaderConfig *StaderConfigCaller) PERMISSIONLESSSOCIALIZINGPOOL(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _StaderConfig.contract.Call(opts, &out, "PERMISSIONLESS_SOCIALIZING_POOL")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// PERMISSIONLESSSOCIALIZINGPOOL is a free data retrieval call binding the contract method 0x3b6bcca0.
//
// Solidity: function PERMISSIONLESS_SOCIALIZING_POOL() view returns(bytes32)
func (_StaderConfig *StaderConfigSession) PERMISSIONLESSSOCIALIZINGPOOL() ([32]byte, error) {
	return _StaderConfig.Contract.PERMISSIONLESSSOCIALIZINGPOOL(&_StaderConfig.CallOpts)
}

// PERMISSIONLESSSOCIALIZINGPOOL is a free data retrieval call binding the contract method 0x3b6bcca0.
//
// Solidity: function PERMISSIONLESS_SOCIALIZING_POOL() view returns(bytes32)
func (_StaderConfig *StaderConfigCallerSession) PERMISSIONLESSSOCIALIZINGPOOL() ([32]byte, error) {
	return _StaderConfig.Contract.PERMISSIONLESSSOCIALIZINGPOOL(&_StaderConfig.CallOpts)
}

// POOLSELECTOR is a free data retrieval call binding the contract method 0xdde63e8f.
//
// Solidity: function POOL_SELECTOR() view returns(bytes32)
func (_StaderConfig *StaderConfigCaller) POOLSELECTOR(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _StaderConfig.contract.Call(opts, &out, "POOL_SELECTOR")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// POOLSELECTOR is a free data retrieval call binding the contract method 0xdde63e8f.
//
// Solidity: function POOL_SELECTOR() view returns(bytes32)
func (_StaderConfig *StaderConfigSession) POOLSELECTOR() ([32]byte, error) {
	return _StaderConfig.Contract.POOLSELECTOR(&_StaderConfig.CallOpts)
}

// POOLSELECTOR is a free data retrieval call binding the contract method 0xdde63e8f.
//
// Solidity: function POOL_SELECTOR() view returns(bytes32)
func (_StaderConfig *StaderConfigCallerSession) POOLSELECTOR() ([32]byte, error) {
	return _StaderConfig.Contract.POOLSELECTOR(&_StaderConfig.CallOpts)
}

// POOLUTILS is a free data retrieval call binding the contract method 0x85e2fcd3.
//
// Solidity: function POOL_UTILS() view returns(bytes32)
func (_StaderConfig *StaderConfigCaller) POOLUTILS(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _StaderConfig.contract.Call(opts, &out, "POOL_UTILS")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// POOLUTILS is a free data retrieval call binding the contract method 0x85e2fcd3.
//
// Solidity: function POOL_UTILS() view returns(bytes32)
func (_StaderConfig *StaderConfigSession) POOLUTILS() ([32]byte, error) {
	return _StaderConfig.Contract.POOLUTILS(&_StaderConfig.CallOpts)
}

// POOLUTILS is a free data retrieval call binding the contract method 0x85e2fcd3.
//
// Solidity: function POOL_UTILS() view returns(bytes32)
func (_StaderConfig *StaderConfigCallerSession) POOLUTILS() ([32]byte, error) {
	return _StaderConfig.Contract.POOLUTILS(&_StaderConfig.CallOpts)
}

// PREDEPOSITSIZE is a free data retrieval call binding the contract method 0x0430246e.
//
// Solidity: function PRE_DEPOSIT_SIZE() view returns(bytes32)
func (_StaderConfig *StaderConfigCaller) PREDEPOSITSIZE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _StaderConfig.contract.Call(opts, &out, "PRE_DEPOSIT_SIZE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// PREDEPOSITSIZE is a free data retrieval call binding the contract method 0x0430246e.
//
// Solidity: function PRE_DEPOSIT_SIZE() view returns(bytes32)
func (_StaderConfig *StaderConfigSession) PREDEPOSITSIZE() ([32]byte, error) {
	return _StaderConfig.Contract.PREDEPOSITSIZE(&_StaderConfig.CallOpts)
}

// PREDEPOSITSIZE is a free data retrieval call binding the contract method 0x0430246e.
//
// Solidity: function PRE_DEPOSIT_SIZE() view returns(bytes32)
func (_StaderConfig *StaderConfigCallerSession) PREDEPOSITSIZE() ([32]byte, error) {
	return _StaderConfig.Contract.PREDEPOSITSIZE(&_StaderConfig.CallOpts)
}

// REWARDTHRESHOLD is a free data retrieval call binding the contract method 0xe7bdba32.
//
// Solidity: function REWARD_THRESHOLD() view returns(bytes32)
func (_StaderConfig *StaderConfigCaller) REWARDTHRESHOLD(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _StaderConfig.contract.Call(opts, &out, "REWARD_THRESHOLD")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// REWARDTHRESHOLD is a free data retrieval call binding the contract method 0xe7bdba32.
//
// Solidity: function REWARD_THRESHOLD() view returns(bytes32)
func (_StaderConfig *StaderConfigSession) REWARDTHRESHOLD() ([32]byte, error) {
	return _StaderConfig.Contract.REWARDTHRESHOLD(&_StaderConfig.CallOpts)
}

// REWARDTHRESHOLD is a free data retrieval call binding the contract method 0xe7bdba32.
//
// Solidity: function REWARD_THRESHOLD() view returns(bytes32)
func (_StaderConfig *StaderConfigCallerSession) REWARDTHRESHOLD() ([32]byte, error) {
	return _StaderConfig.Contract.REWARDTHRESHOLD(&_StaderConfig.CallOpts)
}

// SD is a free data retrieval call binding the contract method 0x384002a2.
//
// Solidity: function SD() view returns(bytes32)
func (_StaderConfig *StaderConfigCaller) SD(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _StaderConfig.contract.Call(opts, &out, "SD")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// SD is a free data retrieval call binding the contract method 0x384002a2.
//
// Solidity: function SD() view returns(bytes32)
func (_StaderConfig *StaderConfigSession) SD() ([32]byte, error) {
	return _StaderConfig.Contract.SD(&_StaderConfig.CallOpts)
}

// SD is a free data retrieval call binding the contract method 0x384002a2.
//
// Solidity: function SD() view returns(bytes32)
func (_StaderConfig *StaderConfigCallerSession) SD() ([32]byte, error) {
	return _StaderConfig.Contract.SD(&_StaderConfig.CallOpts)
}

// SDCOLLATERAL is a free data retrieval call binding the contract method 0xf122961f.
//
// Solidity: function SD_COLLATERAL() view returns(bytes32)
func (_StaderConfig *StaderConfigCaller) SDCOLLATERAL(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _StaderConfig.contract.Call(opts, &out, "SD_COLLATERAL")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// SDCOLLATERAL is a free data retrieval call binding the contract method 0xf122961f.
//
// Solidity: function SD_COLLATERAL() view returns(bytes32)
func (_StaderConfig *StaderConfigSession) SDCOLLATERAL() ([32]byte, error) {
	return _StaderConfig.Contract.SDCOLLATERAL(&_StaderConfig.CallOpts)
}

// SDCOLLATERAL is a free data retrieval call binding the contract method 0xf122961f.
//
// Solidity: function SD_COLLATERAL() view returns(bytes32)
func (_StaderConfig *StaderConfigCallerSession) SDCOLLATERAL() ([32]byte, error) {
	return _StaderConfig.Contract.SDCOLLATERAL(&_StaderConfig.CallOpts)
}

// SOCIALIZINGPOOLCYCLEDURATION is a free data retrieval call binding the contract method 0xbedcb34c.
//
// Solidity: function SOCIALIZING_POOL_CYCLE_DURATION() view returns(bytes32)
func (_StaderConfig *StaderConfigCaller) SOCIALIZINGPOOLCYCLEDURATION(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _StaderConfig.contract.Call(opts, &out, "SOCIALIZING_POOL_CYCLE_DURATION")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// SOCIALIZINGPOOLCYCLEDURATION is a free data retrieval call binding the contract method 0xbedcb34c.
//
// Solidity: function SOCIALIZING_POOL_CYCLE_DURATION() view returns(bytes32)
func (_StaderConfig *StaderConfigSession) SOCIALIZINGPOOLCYCLEDURATION() ([32]byte, error) {
	return _StaderConfig.Contract.SOCIALIZINGPOOLCYCLEDURATION(&_StaderConfig.CallOpts)
}

// SOCIALIZINGPOOLCYCLEDURATION is a free data retrieval call binding the contract method 0xbedcb34c.
//
// Solidity: function SOCIALIZING_POOL_CYCLE_DURATION() view returns(bytes32)
func (_StaderConfig *StaderConfigCallerSession) SOCIALIZINGPOOLCYCLEDURATION() ([32]byte, error) {
	return _StaderConfig.Contract.SOCIALIZINGPOOLCYCLEDURATION(&_StaderConfig.CallOpts)
}

// SOCIALIZINGPOOLOPTINCOOLINGPERIOD is a free data retrieval call binding the contract method 0x686a8b67.
//
// Solidity: function SOCIALIZING_POOL_OPT_IN_COOLING_PERIOD() view returns(bytes32)
func (_StaderConfig *StaderConfigCaller) SOCIALIZINGPOOLOPTINCOOLINGPERIOD(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _StaderConfig.contract.Call(opts, &out, "SOCIALIZING_POOL_OPT_IN_COOLING_PERIOD")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// SOCIALIZINGPOOLOPTINCOOLINGPERIOD is a free data retrieval call binding the contract method 0x686a8b67.
//
// Solidity: function SOCIALIZING_POOL_OPT_IN_COOLING_PERIOD() view returns(bytes32)
func (_StaderConfig *StaderConfigSession) SOCIALIZINGPOOLOPTINCOOLINGPERIOD() ([32]byte, error) {
	return _StaderConfig.Contract.SOCIALIZINGPOOLOPTINCOOLINGPERIOD(&_StaderConfig.CallOpts)
}

// SOCIALIZINGPOOLOPTINCOOLINGPERIOD is a free data retrieval call binding the contract method 0x686a8b67.
//
// Solidity: function SOCIALIZING_POOL_OPT_IN_COOLING_PERIOD() view returns(bytes32)
func (_StaderConfig *StaderConfigCallerSession) SOCIALIZINGPOOLOPTINCOOLINGPERIOD() ([32]byte, error) {
	return _StaderConfig.Contract.SOCIALIZINGPOOLOPTINCOOLINGPERIOD(&_StaderConfig.CallOpts)
}

// STADERINSURANCEFUND is a free data retrieval call binding the contract method 0x1af0fff3.
//
// Solidity: function STADER_INSURANCE_FUND() view returns(bytes32)
func (_StaderConfig *StaderConfigCaller) STADERINSURANCEFUND(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _StaderConfig.contract.Call(opts, &out, "STADER_INSURANCE_FUND")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// STADERINSURANCEFUND is a free data retrieval call binding the contract method 0x1af0fff3.
//
// Solidity: function STADER_INSURANCE_FUND() view returns(bytes32)
func (_StaderConfig *StaderConfigSession) STADERINSURANCEFUND() ([32]byte, error) {
	return _StaderConfig.Contract.STADERINSURANCEFUND(&_StaderConfig.CallOpts)
}

// STADERINSURANCEFUND is a free data retrieval call binding the contract method 0x1af0fff3.
//
// Solidity: function STADER_INSURANCE_FUND() view returns(bytes32)
func (_StaderConfig *StaderConfigCallerSession) STADERINSURANCEFUND() ([32]byte, error) {
	return _StaderConfig.Contract.STADERINSURANCEFUND(&_StaderConfig.CallOpts)
}

// STADERORACLE is a free data retrieval call binding the contract method 0x3871d0f1.
//
// Solidity: function STADER_ORACLE() view returns(bytes32)
func (_StaderConfig *StaderConfigCaller) STADERORACLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _StaderConfig.contract.Call(opts, &out, "STADER_ORACLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// STADERORACLE is a free data retrieval call binding the contract method 0x3871d0f1.
//
// Solidity: function STADER_ORACLE() view returns(bytes32)
func (_StaderConfig *StaderConfigSession) STADERORACLE() ([32]byte, error) {
	return _StaderConfig.Contract.STADERORACLE(&_StaderConfig.CallOpts)
}

// STADERORACLE is a free data retrieval call binding the contract method 0x3871d0f1.
//
// Solidity: function STADER_ORACLE() view returns(bytes32)
func (_StaderConfig *StaderConfigCallerSession) STADERORACLE() ([32]byte, error) {
	return _StaderConfig.Contract.STADERORACLE(&_StaderConfig.CallOpts)
}

// STADERTREASURY is a free data retrieval call binding the contract method 0x841b83b3.
//
// Solidity: function STADER_TREASURY() view returns(bytes32)
func (_StaderConfig *StaderConfigCaller) STADERTREASURY(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _StaderConfig.contract.Call(opts, &out, "STADER_TREASURY")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// STADERTREASURY is a free data retrieval call binding the contract method 0x841b83b3.
//
// Solidity: function STADER_TREASURY() view returns(bytes32)
func (_StaderConfig *StaderConfigSession) STADERTREASURY() ([32]byte, error) {
	return _StaderConfig.Contract.STADERTREASURY(&_StaderConfig.CallOpts)
}

// STADERTREASURY is a free data retrieval call binding the contract method 0x841b83b3.
//
// Solidity: function STADER_TREASURY() view returns(bytes32)
func (_StaderConfig *StaderConfigCallerSession) STADERTREASURY() ([32]byte, error) {
	return _StaderConfig.Contract.STADERTREASURY(&_StaderConfig.CallOpts)
}

// STAKEPOOLMANAGER is a free data retrieval call binding the contract method 0xa53bddd6.
//
// Solidity: function STAKE_POOL_MANAGER() view returns(bytes32)
func (_StaderConfig *StaderConfigCaller) STAKEPOOLMANAGER(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _StaderConfig.contract.Call(opts, &out, "STAKE_POOL_MANAGER")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// STAKEPOOLMANAGER is a free data retrieval call binding the contract method 0xa53bddd6.
//
// Solidity: function STAKE_POOL_MANAGER() view returns(bytes32)
func (_StaderConfig *StaderConfigSession) STAKEPOOLMANAGER() ([32]byte, error) {
	return _StaderConfig.Contract.STAKEPOOLMANAGER(&_StaderConfig.CallOpts)
}

// STAKEPOOLMANAGER is a free data retrieval call binding the contract method 0xa53bddd6.
//
// Solidity: function STAKE_POOL_MANAGER() view returns(bytes32)
func (_StaderConfig *StaderConfigCallerSession) STAKEPOOLMANAGER() ([32]byte, error) {
	return _StaderConfig.Contract.STAKEPOOLMANAGER(&_StaderConfig.CallOpts)
}

// TOTALFEE is a free data retrieval call binding the contract method 0x63db7eae.
//
// Solidity: function TOTAL_FEE() view returns(bytes32)
func (_StaderConfig *StaderConfigCaller) TOTALFEE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _StaderConfig.contract.Call(opts, &out, "TOTAL_FEE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// TOTALFEE is a free data retrieval call binding the contract method 0x63db7eae.
//
// Solidity: function TOTAL_FEE() view returns(bytes32)
func (_StaderConfig *StaderConfigSession) TOTALFEE() ([32]byte, error) {
	return _StaderConfig.Contract.TOTALFEE(&_StaderConfig.CallOpts)
}

// TOTALFEE is a free data retrieval call binding the contract method 0x63db7eae.
//
// Solidity: function TOTAL_FEE() view returns(bytes32)
func (_StaderConfig *StaderConfigCallerSession) TOTALFEE() ([32]byte, error) {
	return _StaderConfig.Contract.TOTALFEE(&_StaderConfig.CallOpts)
}

// USERWITHDRAWMANAGER is a free data retrieval call binding the contract method 0x36854d63.
//
// Solidity: function USER_WITHDRAW_MANAGER() view returns(bytes32)
func (_StaderConfig *StaderConfigCaller) USERWITHDRAWMANAGER(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _StaderConfig.contract.Call(opts, &out, "USER_WITHDRAW_MANAGER")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// USERWITHDRAWMANAGER is a free data retrieval call binding the contract method 0x36854d63.
//
// Solidity: function USER_WITHDRAW_MANAGER() view returns(bytes32)
func (_StaderConfig *StaderConfigSession) USERWITHDRAWMANAGER() ([32]byte, error) {
	return _StaderConfig.Contract.USERWITHDRAWMANAGER(&_StaderConfig.CallOpts)
}

// USERWITHDRAWMANAGER is a free data retrieval call binding the contract method 0x36854d63.
//
// Solidity: function USER_WITHDRAW_MANAGER() view returns(bytes32)
func (_StaderConfig *StaderConfigCallerSession) USERWITHDRAWMANAGER() ([32]byte, error) {
	return _StaderConfig.Contract.USERWITHDRAWMANAGER(&_StaderConfig.CallOpts)
}

// VALIDATORWITHDRAWALVAULTIMPLEMENTATION is a free data retrieval call binding the contract method 0x1c55cccd.
//
// Solidity: function VALIDATOR_WITHDRAWAL_VAULT_IMPLEMENTATION() view returns(bytes32)
func (_StaderConfig *StaderConfigCaller) VALIDATORWITHDRAWALVAULTIMPLEMENTATION(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _StaderConfig.contract.Call(opts, &out, "VALIDATOR_WITHDRAWAL_VAULT_IMPLEMENTATION")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// VALIDATORWITHDRAWALVAULTIMPLEMENTATION is a free data retrieval call binding the contract method 0x1c55cccd.
//
// Solidity: function VALIDATOR_WITHDRAWAL_VAULT_IMPLEMENTATION() view returns(bytes32)
func (_StaderConfig *StaderConfigSession) VALIDATORWITHDRAWALVAULTIMPLEMENTATION() ([32]byte, error) {
	return _StaderConfig.Contract.VALIDATORWITHDRAWALVAULTIMPLEMENTATION(&_StaderConfig.CallOpts)
}

// VALIDATORWITHDRAWALVAULTIMPLEMENTATION is a free data retrieval call binding the contract method 0x1c55cccd.
//
// Solidity: function VALIDATOR_WITHDRAWAL_VAULT_IMPLEMENTATION() view returns(bytes32)
func (_StaderConfig *StaderConfigCallerSession) VALIDATORWITHDRAWALVAULTIMPLEMENTATION() ([32]byte, error) {
	return _StaderConfig.Contract.VALIDATORWITHDRAWALVAULTIMPLEMENTATION(&_StaderConfig.CallOpts)
}

// VAULTFACTORY is a free data retrieval call binding the contract method 0x103f2907.
//
// Solidity: function VAULT_FACTORY() view returns(bytes32)
func (_StaderConfig *StaderConfigCaller) VAULTFACTORY(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _StaderConfig.contract.Call(opts, &out, "VAULT_FACTORY")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// VAULTFACTORY is a free data retrieval call binding the contract method 0x103f2907.
//
// Solidity: function VAULT_FACTORY() view returns(bytes32)
func (_StaderConfig *StaderConfigSession) VAULTFACTORY() ([32]byte, error) {
	return _StaderConfig.Contract.VAULTFACTORY(&_StaderConfig.CallOpts)
}

// VAULTFACTORY is a free data retrieval call binding the contract method 0x103f2907.
//
// Solidity: function VAULT_FACTORY() view returns(bytes32)
func (_StaderConfig *StaderConfigCallerSession) VAULTFACTORY() ([32]byte, error) {
	return _StaderConfig.Contract.VAULTFACTORY(&_StaderConfig.CallOpts)
}

// WITHDRAWNKEYSBATCHSIZE is a free data retrieval call binding the contract method 0x88993d8b.
//
// Solidity: function WITHDRAWN_KEYS_BATCH_SIZE() view returns(bytes32)
func (_StaderConfig *StaderConfigCaller) WITHDRAWNKEYSBATCHSIZE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _StaderConfig.contract.Call(opts, &out, "WITHDRAWN_KEYS_BATCH_SIZE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// WITHDRAWNKEYSBATCHSIZE is a free data retrieval call binding the contract method 0x88993d8b.
//
// Solidity: function WITHDRAWN_KEYS_BATCH_SIZE() view returns(bytes32)
func (_StaderConfig *StaderConfigSession) WITHDRAWNKEYSBATCHSIZE() ([32]byte, error) {
	return _StaderConfig.Contract.WITHDRAWNKEYSBATCHSIZE(&_StaderConfig.CallOpts)
}

// WITHDRAWNKEYSBATCHSIZE is a free data retrieval call binding the contract method 0x88993d8b.
//
// Solidity: function WITHDRAWN_KEYS_BATCH_SIZE() view returns(bytes32)
func (_StaderConfig *StaderConfigCallerSession) WITHDRAWNKEYSBATCHSIZE() ([32]byte, error) {
	return _StaderConfig.Contract.WITHDRAWNKEYSBATCHSIZE(&_StaderConfig.CallOpts)
}

// GetAdmin is a free data retrieval call binding the contract method 0x6e9960c3.
//
// Solidity: function getAdmin() view returns(address)
func (_StaderConfig *StaderConfigCaller) GetAdmin(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _StaderConfig.contract.Call(opts, &out, "getAdmin")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetAdmin is a free data retrieval call binding the contract method 0x6e9960c3.
//
// Solidity: function getAdmin() view returns(address)
func (_StaderConfig *StaderConfigSession) GetAdmin() (common.Address, error) {
	return _StaderConfig.Contract.GetAdmin(&_StaderConfig.CallOpts)
}

// GetAdmin is a free data retrieval call binding the contract method 0x6e9960c3.
//
// Solidity: function getAdmin() view returns(address)
func (_StaderConfig *StaderConfigCallerSession) GetAdmin() (common.Address, error) {
	return _StaderConfig.Contract.GetAdmin(&_StaderConfig.CallOpts)
}

// GetAuctionContract is a free data retrieval call binding the contract method 0x36c157f4.
//
// Solidity: function getAuctionContract() view returns(address)
func (_StaderConfig *StaderConfigCaller) GetAuctionContract(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _StaderConfig.contract.Call(opts, &out, "getAuctionContract")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetAuctionContract is a free data retrieval call binding the contract method 0x36c157f4.
//
// Solidity: function getAuctionContract() view returns(address)
func (_StaderConfig *StaderConfigSession) GetAuctionContract() (common.Address, error) {
	return _StaderConfig.Contract.GetAuctionContract(&_StaderConfig.CallOpts)
}

// GetAuctionContract is a free data retrieval call binding the contract method 0x36c157f4.
//
// Solidity: function getAuctionContract() view returns(address)
func (_StaderConfig *StaderConfigCallerSession) GetAuctionContract() (common.Address, error) {
	return _StaderConfig.Contract.GetAuctionContract(&_StaderConfig.CallOpts)
}

// GetDecimals is a free data retrieval call binding the contract method 0xf0141d84.
//
// Solidity: function getDecimals() view returns(uint256)
func (_StaderConfig *StaderConfigCaller) GetDecimals(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _StaderConfig.contract.Call(opts, &out, "getDecimals")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetDecimals is a free data retrieval call binding the contract method 0xf0141d84.
//
// Solidity: function getDecimals() view returns(uint256)
func (_StaderConfig *StaderConfigSession) GetDecimals() (*big.Int, error) {
	return _StaderConfig.Contract.GetDecimals(&_StaderConfig.CallOpts)
}

// GetDecimals is a free data retrieval call binding the contract method 0xf0141d84.
//
// Solidity: function getDecimals() view returns(uint256)
func (_StaderConfig *StaderConfigCallerSession) GetDecimals() (*big.Int, error) {
	return _StaderConfig.Contract.GetDecimals(&_StaderConfig.CallOpts)
}

// GetETHBalancePORFeedProxy is a free data retrieval call binding the contract method 0x489ed651.
//
// Solidity: function getETHBalancePORFeedProxy() view returns(address)
func (_StaderConfig *StaderConfigCaller) GetETHBalancePORFeedProxy(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _StaderConfig.contract.Call(opts, &out, "getETHBalancePORFeedProxy")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetETHBalancePORFeedProxy is a free data retrieval call binding the contract method 0x489ed651.
//
// Solidity: function getETHBalancePORFeedProxy() view returns(address)
func (_StaderConfig *StaderConfigSession) GetETHBalancePORFeedProxy() (common.Address, error) {
	return _StaderConfig.Contract.GetETHBalancePORFeedProxy(&_StaderConfig.CallOpts)
}

// GetETHBalancePORFeedProxy is a free data retrieval call binding the contract method 0x489ed651.
//
// Solidity: function getETHBalancePORFeedProxy() view returns(address)
func (_StaderConfig *StaderConfigCallerSession) GetETHBalancePORFeedProxy() (common.Address, error) {
	return _StaderConfig.Contract.GetETHBalancePORFeedProxy(&_StaderConfig.CallOpts)
}

// GetETHDepositContract is a free data retrieval call binding the contract method 0x8f8b3867.
//
// Solidity: function getETHDepositContract() view returns(address)
func (_StaderConfig *StaderConfigCaller) GetETHDepositContract(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _StaderConfig.contract.Call(opts, &out, "getETHDepositContract")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetETHDepositContract is a free data retrieval call binding the contract method 0x8f8b3867.
//
// Solidity: function getETHDepositContract() view returns(address)
func (_StaderConfig *StaderConfigSession) GetETHDepositContract() (common.Address, error) {
	return _StaderConfig.Contract.GetETHDepositContract(&_StaderConfig.CallOpts)
}

// GetETHDepositContract is a free data retrieval call binding the contract method 0x8f8b3867.
//
// Solidity: function getETHDepositContract() view returns(address)
func (_StaderConfig *StaderConfigCallerSession) GetETHDepositContract() (common.Address, error) {
	return _StaderConfig.Contract.GetETHDepositContract(&_StaderConfig.CallOpts)
}

// GetETHXSupplyPORFeedProxy is a free data retrieval call binding the contract method 0x2ca03f66.
//
// Solidity: function getETHXSupplyPORFeedProxy() view returns(address)
func (_StaderConfig *StaderConfigCaller) GetETHXSupplyPORFeedProxy(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _StaderConfig.contract.Call(opts, &out, "getETHXSupplyPORFeedProxy")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetETHXSupplyPORFeedProxy is a free data retrieval call binding the contract method 0x2ca03f66.
//
// Solidity: function getETHXSupplyPORFeedProxy() view returns(address)
func (_StaderConfig *StaderConfigSession) GetETHXSupplyPORFeedProxy() (common.Address, error) {
	return _StaderConfig.Contract.GetETHXSupplyPORFeedProxy(&_StaderConfig.CallOpts)
}

// GetETHXSupplyPORFeedProxy is a free data retrieval call binding the contract method 0x2ca03f66.
//
// Solidity: function getETHXSupplyPORFeedProxy() view returns(address)
func (_StaderConfig *StaderConfigCallerSession) GetETHXSupplyPORFeedProxy() (common.Address, error) {
	return _StaderConfig.Contract.GetETHXSupplyPORFeedProxy(&_StaderConfig.CallOpts)
}

// GetETHxToken is a free data retrieval call binding the contract method 0xcc45dabe.
//
// Solidity: function getETHxToken() view returns(address)
func (_StaderConfig *StaderConfigCaller) GetETHxToken(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _StaderConfig.contract.Call(opts, &out, "getETHxToken")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetETHxToken is a free data retrieval call binding the contract method 0xcc45dabe.
//
// Solidity: function getETHxToken() view returns(address)
func (_StaderConfig *StaderConfigSession) GetETHxToken() (common.Address, error) {
	return _StaderConfig.Contract.GetETHxToken(&_StaderConfig.CallOpts)
}

// GetETHxToken is a free data retrieval call binding the contract method 0xcc45dabe.
//
// Solidity: function getETHxToken() view returns(address)
func (_StaderConfig *StaderConfigCallerSession) GetETHxToken() (common.Address, error) {
	return _StaderConfig.Contract.GetETHxToken(&_StaderConfig.CallOpts)
}

// GetFullDepositSize is a free data retrieval call binding the contract method 0xfa71fcbb.
//
// Solidity: function getFullDepositSize() view returns(uint256)
func (_StaderConfig *StaderConfigCaller) GetFullDepositSize(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _StaderConfig.contract.Call(opts, &out, "getFullDepositSize")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetFullDepositSize is a free data retrieval call binding the contract method 0xfa71fcbb.
//
// Solidity: function getFullDepositSize() view returns(uint256)
func (_StaderConfig *StaderConfigSession) GetFullDepositSize() (*big.Int, error) {
	return _StaderConfig.Contract.GetFullDepositSize(&_StaderConfig.CallOpts)
}

// GetFullDepositSize is a free data retrieval call binding the contract method 0xfa71fcbb.
//
// Solidity: function getFullDepositSize() view returns(uint256)
func (_StaderConfig *StaderConfigCallerSession) GetFullDepositSize() (*big.Int, error) {
	return _StaderConfig.Contract.GetFullDepositSize(&_StaderConfig.CallOpts)
}

// GetMaxDepositAmount is a free data retrieval call binding the contract method 0x5726a356.
//
// Solidity: function getMaxDepositAmount() view returns(uint256)
func (_StaderConfig *StaderConfigCaller) GetMaxDepositAmount(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _StaderConfig.contract.Call(opts, &out, "getMaxDepositAmount")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetMaxDepositAmount is a free data retrieval call binding the contract method 0x5726a356.
//
// Solidity: function getMaxDepositAmount() view returns(uint256)
func (_StaderConfig *StaderConfigSession) GetMaxDepositAmount() (*big.Int, error) {
	return _StaderConfig.Contract.GetMaxDepositAmount(&_StaderConfig.CallOpts)
}

// GetMaxDepositAmount is a free data retrieval call binding the contract method 0x5726a356.
//
// Solidity: function getMaxDepositAmount() view returns(uint256)
func (_StaderConfig *StaderConfigCallerSession) GetMaxDepositAmount() (*big.Int, error) {
	return _StaderConfig.Contract.GetMaxDepositAmount(&_StaderConfig.CallOpts)
}

// GetMaxWithdrawAmount is a free data retrieval call binding the contract method 0x326a16a3.
//
// Solidity: function getMaxWithdrawAmount() view returns(uint256)
func (_StaderConfig *StaderConfigCaller) GetMaxWithdrawAmount(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _StaderConfig.contract.Call(opts, &out, "getMaxWithdrawAmount")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetMaxWithdrawAmount is a free data retrieval call binding the contract method 0x326a16a3.
//
// Solidity: function getMaxWithdrawAmount() view returns(uint256)
func (_StaderConfig *StaderConfigSession) GetMaxWithdrawAmount() (*big.Int, error) {
	return _StaderConfig.Contract.GetMaxWithdrawAmount(&_StaderConfig.CallOpts)
}

// GetMaxWithdrawAmount is a free data retrieval call binding the contract method 0x326a16a3.
//
// Solidity: function getMaxWithdrawAmount() view returns(uint256)
func (_StaderConfig *StaderConfigCallerSession) GetMaxWithdrawAmount() (*big.Int, error) {
	return _StaderConfig.Contract.GetMaxWithdrawAmount(&_StaderConfig.CallOpts)
}

// GetMinBlockDelayToFinalizeWithdrawRequest is a free data retrieval call binding the contract method 0xd2cee8ba.
//
// Solidity: function getMinBlockDelayToFinalizeWithdrawRequest() view returns(uint256)
func (_StaderConfig *StaderConfigCaller) GetMinBlockDelayToFinalizeWithdrawRequest(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _StaderConfig.contract.Call(opts, &out, "getMinBlockDelayToFinalizeWithdrawRequest")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetMinBlockDelayToFinalizeWithdrawRequest is a free data retrieval call binding the contract method 0xd2cee8ba.
//
// Solidity: function getMinBlockDelayToFinalizeWithdrawRequest() view returns(uint256)
func (_StaderConfig *StaderConfigSession) GetMinBlockDelayToFinalizeWithdrawRequest() (*big.Int, error) {
	return _StaderConfig.Contract.GetMinBlockDelayToFinalizeWithdrawRequest(&_StaderConfig.CallOpts)
}

// GetMinBlockDelayToFinalizeWithdrawRequest is a free data retrieval call binding the contract method 0xd2cee8ba.
//
// Solidity: function getMinBlockDelayToFinalizeWithdrawRequest() view returns(uint256)
func (_StaderConfig *StaderConfigCallerSession) GetMinBlockDelayToFinalizeWithdrawRequest() (*big.Int, error) {
	return _StaderConfig.Contract.GetMinBlockDelayToFinalizeWithdrawRequest(&_StaderConfig.CallOpts)
}

// GetMinDepositAmount is a free data retrieval call binding the contract method 0xf4914d33.
//
// Solidity: function getMinDepositAmount() view returns(uint256)
func (_StaderConfig *StaderConfigCaller) GetMinDepositAmount(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _StaderConfig.contract.Call(opts, &out, "getMinDepositAmount")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetMinDepositAmount is a free data retrieval call binding the contract method 0xf4914d33.
//
// Solidity: function getMinDepositAmount() view returns(uint256)
func (_StaderConfig *StaderConfigSession) GetMinDepositAmount() (*big.Int, error) {
	return _StaderConfig.Contract.GetMinDepositAmount(&_StaderConfig.CallOpts)
}

// GetMinDepositAmount is a free data retrieval call binding the contract method 0xf4914d33.
//
// Solidity: function getMinDepositAmount() view returns(uint256)
func (_StaderConfig *StaderConfigCallerSession) GetMinDepositAmount() (*big.Int, error) {
	return _StaderConfig.Contract.GetMinDepositAmount(&_StaderConfig.CallOpts)
}

// GetMinWithdrawAmount is a free data retrieval call binding the contract method 0x14e1b8fd.
//
// Solidity: function getMinWithdrawAmount() view returns(uint256)
func (_StaderConfig *StaderConfigCaller) GetMinWithdrawAmount(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _StaderConfig.contract.Call(opts, &out, "getMinWithdrawAmount")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetMinWithdrawAmount is a free data retrieval call binding the contract method 0x14e1b8fd.
//
// Solidity: function getMinWithdrawAmount() view returns(uint256)
func (_StaderConfig *StaderConfigSession) GetMinWithdrawAmount() (*big.Int, error) {
	return _StaderConfig.Contract.GetMinWithdrawAmount(&_StaderConfig.CallOpts)
}

// GetMinWithdrawAmount is a free data retrieval call binding the contract method 0x14e1b8fd.
//
// Solidity: function getMinWithdrawAmount() view returns(uint256)
func (_StaderConfig *StaderConfigCallerSession) GetMinWithdrawAmount() (*big.Int, error) {
	return _StaderConfig.Contract.GetMinWithdrawAmount(&_StaderConfig.CallOpts)
}

// GetNodeELRewardVaultImplementation is a free data retrieval call binding the contract method 0xe8fe1873.
//
// Solidity: function getNodeELRewardVaultImplementation() view returns(address)
func (_StaderConfig *StaderConfigCaller) GetNodeELRewardVaultImplementation(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _StaderConfig.contract.Call(opts, &out, "getNodeELRewardVaultImplementation")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetNodeELRewardVaultImplementation is a free data retrieval call binding the contract method 0xe8fe1873.
//
// Solidity: function getNodeELRewardVaultImplementation() view returns(address)
func (_StaderConfig *StaderConfigSession) GetNodeELRewardVaultImplementation() (common.Address, error) {
	return _StaderConfig.Contract.GetNodeELRewardVaultImplementation(&_StaderConfig.CallOpts)
}

// GetNodeELRewardVaultImplementation is a free data retrieval call binding the contract method 0xe8fe1873.
//
// Solidity: function getNodeELRewardVaultImplementation() view returns(address)
func (_StaderConfig *StaderConfigCallerSession) GetNodeELRewardVaultImplementation() (common.Address, error) {
	return _StaderConfig.Contract.GetNodeELRewardVaultImplementation(&_StaderConfig.CallOpts)
}

// GetOperatorMaxNameLength is a free data retrieval call binding the contract method 0x10deba2b.
//
// Solidity: function getOperatorMaxNameLength() view returns(uint256)
func (_StaderConfig *StaderConfigCaller) GetOperatorMaxNameLength(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _StaderConfig.contract.Call(opts, &out, "getOperatorMaxNameLength")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetOperatorMaxNameLength is a free data retrieval call binding the contract method 0x10deba2b.
//
// Solidity: function getOperatorMaxNameLength() view returns(uint256)
func (_StaderConfig *StaderConfigSession) GetOperatorMaxNameLength() (*big.Int, error) {
	return _StaderConfig.Contract.GetOperatorMaxNameLength(&_StaderConfig.CallOpts)
}

// GetOperatorMaxNameLength is a free data retrieval call binding the contract method 0x10deba2b.
//
// Solidity: function getOperatorMaxNameLength() view returns(uint256)
func (_StaderConfig *StaderConfigCallerSession) GetOperatorMaxNameLength() (*big.Int, error) {
	return _StaderConfig.Contract.GetOperatorMaxNameLength(&_StaderConfig.CallOpts)
}

// GetOperatorRewardsCollector is a free data retrieval call binding the contract method 0x278671bb.
//
// Solidity: function getOperatorRewardsCollector() view returns(address)
func (_StaderConfig *StaderConfigCaller) GetOperatorRewardsCollector(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _StaderConfig.contract.Call(opts, &out, "getOperatorRewardsCollector")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetOperatorRewardsCollector is a free data retrieval call binding the contract method 0x278671bb.
//
// Solidity: function getOperatorRewardsCollector() view returns(address)
func (_StaderConfig *StaderConfigSession) GetOperatorRewardsCollector() (common.Address, error) {
	return _StaderConfig.Contract.GetOperatorRewardsCollector(&_StaderConfig.CallOpts)
}

// GetOperatorRewardsCollector is a free data retrieval call binding the contract method 0x278671bb.
//
// Solidity: function getOperatorRewardsCollector() view returns(address)
func (_StaderConfig *StaderConfigCallerSession) GetOperatorRewardsCollector() (common.Address, error) {
	return _StaderConfig.Contract.GetOperatorRewardsCollector(&_StaderConfig.CallOpts)
}

// GetPenaltyContract is a free data retrieval call binding the contract method 0x8910115c.
//
// Solidity: function getPenaltyContract() view returns(address)
func (_StaderConfig *StaderConfigCaller) GetPenaltyContract(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _StaderConfig.contract.Call(opts, &out, "getPenaltyContract")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetPenaltyContract is a free data retrieval call binding the contract method 0x8910115c.
//
// Solidity: function getPenaltyContract() view returns(address)
func (_StaderConfig *StaderConfigSession) GetPenaltyContract() (common.Address, error) {
	return _StaderConfig.Contract.GetPenaltyContract(&_StaderConfig.CallOpts)
}

// GetPenaltyContract is a free data retrieval call binding the contract method 0x8910115c.
//
// Solidity: function getPenaltyContract() view returns(address)
func (_StaderConfig *StaderConfigCallerSession) GetPenaltyContract() (common.Address, error) {
	return _StaderConfig.Contract.GetPenaltyContract(&_StaderConfig.CallOpts)
}

// GetPermissionedNodeRegistry is a free data retrieval call binding the contract method 0x5edc686e.
//
// Solidity: function getPermissionedNodeRegistry() view returns(address)
func (_StaderConfig *StaderConfigCaller) GetPermissionedNodeRegistry(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _StaderConfig.contract.Call(opts, &out, "getPermissionedNodeRegistry")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetPermissionedNodeRegistry is a free data retrieval call binding the contract method 0x5edc686e.
//
// Solidity: function getPermissionedNodeRegistry() view returns(address)
func (_StaderConfig *StaderConfigSession) GetPermissionedNodeRegistry() (common.Address, error) {
	return _StaderConfig.Contract.GetPermissionedNodeRegistry(&_StaderConfig.CallOpts)
}

// GetPermissionedNodeRegistry is a free data retrieval call binding the contract method 0x5edc686e.
//
// Solidity: function getPermissionedNodeRegistry() view returns(address)
func (_StaderConfig *StaderConfigCallerSession) GetPermissionedNodeRegistry() (common.Address, error) {
	return _StaderConfig.Contract.GetPermissionedNodeRegistry(&_StaderConfig.CallOpts)
}

// GetPermissionedPool is a free data retrieval call binding the contract method 0xa0b4079f.
//
// Solidity: function getPermissionedPool() view returns(address)
func (_StaderConfig *StaderConfigCaller) GetPermissionedPool(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _StaderConfig.contract.Call(opts, &out, "getPermissionedPool")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetPermissionedPool is a free data retrieval call binding the contract method 0xa0b4079f.
//
// Solidity: function getPermissionedPool() view returns(address)
func (_StaderConfig *StaderConfigSession) GetPermissionedPool() (common.Address, error) {
	return _StaderConfig.Contract.GetPermissionedPool(&_StaderConfig.CallOpts)
}

// GetPermissionedPool is a free data retrieval call binding the contract method 0xa0b4079f.
//
// Solidity: function getPermissionedPool() view returns(address)
func (_StaderConfig *StaderConfigCallerSession) GetPermissionedPool() (common.Address, error) {
	return _StaderConfig.Contract.GetPermissionedPool(&_StaderConfig.CallOpts)
}

// GetPermissionedSocializingPool is a free data retrieval call binding the contract method 0xa469e247.
//
// Solidity: function getPermissionedSocializingPool() view returns(address)
func (_StaderConfig *StaderConfigCaller) GetPermissionedSocializingPool(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _StaderConfig.contract.Call(opts, &out, "getPermissionedSocializingPool")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetPermissionedSocializingPool is a free data retrieval call binding the contract method 0xa469e247.
//
// Solidity: function getPermissionedSocializingPool() view returns(address)
func (_StaderConfig *StaderConfigSession) GetPermissionedSocializingPool() (common.Address, error) {
	return _StaderConfig.Contract.GetPermissionedSocializingPool(&_StaderConfig.CallOpts)
}

// GetPermissionedSocializingPool is a free data retrieval call binding the contract method 0xa469e247.
//
// Solidity: function getPermissionedSocializingPool() view returns(address)
func (_StaderConfig *StaderConfigCallerSession) GetPermissionedSocializingPool() (common.Address, error) {
	return _StaderConfig.Contract.GetPermissionedSocializingPool(&_StaderConfig.CallOpts)
}

// GetPermissionlessNodeRegistry is a free data retrieval call binding the contract method 0x360374a4.
//
// Solidity: function getPermissionlessNodeRegistry() view returns(address)
func (_StaderConfig *StaderConfigCaller) GetPermissionlessNodeRegistry(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _StaderConfig.contract.Call(opts, &out, "getPermissionlessNodeRegistry")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetPermissionlessNodeRegistry is a free data retrieval call binding the contract method 0x360374a4.
//
// Solidity: function getPermissionlessNodeRegistry() view returns(address)
func (_StaderConfig *StaderConfigSession) GetPermissionlessNodeRegistry() (common.Address, error) {
	return _StaderConfig.Contract.GetPermissionlessNodeRegistry(&_StaderConfig.CallOpts)
}

// GetPermissionlessNodeRegistry is a free data retrieval call binding the contract method 0x360374a4.
//
// Solidity: function getPermissionlessNodeRegistry() view returns(address)
func (_StaderConfig *StaderConfigCallerSession) GetPermissionlessNodeRegistry() (common.Address, error) {
	return _StaderConfig.Contract.GetPermissionlessNodeRegistry(&_StaderConfig.CallOpts)
}

// GetPermissionlessPool is a free data retrieval call binding the contract method 0x9ca76b73.
//
// Solidity: function getPermissionlessPool() view returns(address)
func (_StaderConfig *StaderConfigCaller) GetPermissionlessPool(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _StaderConfig.contract.Call(opts, &out, "getPermissionlessPool")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetPermissionlessPool is a free data retrieval call binding the contract method 0x9ca76b73.
//
// Solidity: function getPermissionlessPool() view returns(address)
func (_StaderConfig *StaderConfigSession) GetPermissionlessPool() (common.Address, error) {
	return _StaderConfig.Contract.GetPermissionlessPool(&_StaderConfig.CallOpts)
}

// GetPermissionlessPool is a free data retrieval call binding the contract method 0x9ca76b73.
//
// Solidity: function getPermissionlessPool() view returns(address)
func (_StaderConfig *StaderConfigCallerSession) GetPermissionlessPool() (common.Address, error) {
	return _StaderConfig.Contract.GetPermissionlessPool(&_StaderConfig.CallOpts)
}

// GetPermissionlessSocializingPool is a free data retrieval call binding the contract method 0x0a3fbd9a.
//
// Solidity: function getPermissionlessSocializingPool() view returns(address)
func (_StaderConfig *StaderConfigCaller) GetPermissionlessSocializingPool(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _StaderConfig.contract.Call(opts, &out, "getPermissionlessSocializingPool")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetPermissionlessSocializingPool is a free data retrieval call binding the contract method 0x0a3fbd9a.
//
// Solidity: function getPermissionlessSocializingPool() view returns(address)
func (_StaderConfig *StaderConfigSession) GetPermissionlessSocializingPool() (common.Address, error) {
	return _StaderConfig.Contract.GetPermissionlessSocializingPool(&_StaderConfig.CallOpts)
}

// GetPermissionlessSocializingPool is a free data retrieval call binding the contract method 0x0a3fbd9a.
//
// Solidity: function getPermissionlessSocializingPool() view returns(address)
func (_StaderConfig *StaderConfigCallerSession) GetPermissionlessSocializingPool() (common.Address, error) {
	return _StaderConfig.Contract.GetPermissionlessSocializingPool(&_StaderConfig.CallOpts)
}

// GetPoolSelector is a free data retrieval call binding the contract method 0x5458a106.
//
// Solidity: function getPoolSelector() view returns(address)
func (_StaderConfig *StaderConfigCaller) GetPoolSelector(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _StaderConfig.contract.Call(opts, &out, "getPoolSelector")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetPoolSelector is a free data retrieval call binding the contract method 0x5458a106.
//
// Solidity: function getPoolSelector() view returns(address)
func (_StaderConfig *StaderConfigSession) GetPoolSelector() (common.Address, error) {
	return _StaderConfig.Contract.GetPoolSelector(&_StaderConfig.CallOpts)
}

// GetPoolSelector is a free data retrieval call binding the contract method 0x5458a106.
//
// Solidity: function getPoolSelector() view returns(address)
func (_StaderConfig *StaderConfigCallerSession) GetPoolSelector() (common.Address, error) {
	return _StaderConfig.Contract.GetPoolSelector(&_StaderConfig.CallOpts)
}

// GetPoolUtils is a free data retrieval call binding the contract method 0x6ccb9d70.
//
// Solidity: function getPoolUtils() view returns(address)
func (_StaderConfig *StaderConfigCaller) GetPoolUtils(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _StaderConfig.contract.Call(opts, &out, "getPoolUtils")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetPoolUtils is a free data retrieval call binding the contract method 0x6ccb9d70.
//
// Solidity: function getPoolUtils() view returns(address)
func (_StaderConfig *StaderConfigSession) GetPoolUtils() (common.Address, error) {
	return _StaderConfig.Contract.GetPoolUtils(&_StaderConfig.CallOpts)
}

// GetPoolUtils is a free data retrieval call binding the contract method 0x6ccb9d70.
//
// Solidity: function getPoolUtils() view returns(address)
func (_StaderConfig *StaderConfigCallerSession) GetPoolUtils() (common.Address, error) {
	return _StaderConfig.Contract.GetPoolUtils(&_StaderConfig.CallOpts)
}

// GetPreDepositSize is a free data retrieval call binding the contract method 0x08297645.
//
// Solidity: function getPreDepositSize() view returns(uint256)
func (_StaderConfig *StaderConfigCaller) GetPreDepositSize(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _StaderConfig.contract.Call(opts, &out, "getPreDepositSize")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetPreDepositSize is a free data retrieval call binding the contract method 0x08297645.
//
// Solidity: function getPreDepositSize() view returns(uint256)
func (_StaderConfig *StaderConfigSession) GetPreDepositSize() (*big.Int, error) {
	return _StaderConfig.Contract.GetPreDepositSize(&_StaderConfig.CallOpts)
}

// GetPreDepositSize is a free data retrieval call binding the contract method 0x08297645.
//
// Solidity: function getPreDepositSize() view returns(uint256)
func (_StaderConfig *StaderConfigCallerSession) GetPreDepositSize() (*big.Int, error) {
	return _StaderConfig.Contract.GetPreDepositSize(&_StaderConfig.CallOpts)
}

// GetRewardsThreshold is a free data retrieval call binding the contract method 0x18829fc3.
//
// Solidity: function getRewardsThreshold() view returns(uint256)
func (_StaderConfig *StaderConfigCaller) GetRewardsThreshold(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _StaderConfig.contract.Call(opts, &out, "getRewardsThreshold")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetRewardsThreshold is a free data retrieval call binding the contract method 0x18829fc3.
//
// Solidity: function getRewardsThreshold() view returns(uint256)
func (_StaderConfig *StaderConfigSession) GetRewardsThreshold() (*big.Int, error) {
	return _StaderConfig.Contract.GetRewardsThreshold(&_StaderConfig.CallOpts)
}

// GetRewardsThreshold is a free data retrieval call binding the contract method 0x18829fc3.
//
// Solidity: function getRewardsThreshold() view returns(uint256)
func (_StaderConfig *StaderConfigCallerSession) GetRewardsThreshold() (*big.Int, error) {
	return _StaderConfig.Contract.GetRewardsThreshold(&_StaderConfig.CallOpts)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_StaderConfig *StaderConfigCaller) GetRoleAdmin(opts *bind.CallOpts, role [32]byte) ([32]byte, error) {
	var out []interface{}
	err := _StaderConfig.contract.Call(opts, &out, "getRoleAdmin", role)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_StaderConfig *StaderConfigSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _StaderConfig.Contract.GetRoleAdmin(&_StaderConfig.CallOpts, role)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_StaderConfig *StaderConfigCallerSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _StaderConfig.Contract.GetRoleAdmin(&_StaderConfig.CallOpts, role)
}

// GetSDCollateral is a free data retrieval call binding the contract method 0xaa953795.
//
// Solidity: function getSDCollateral() view returns(address)
func (_StaderConfig *StaderConfigCaller) GetSDCollateral(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _StaderConfig.contract.Call(opts, &out, "getSDCollateral")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetSDCollateral is a free data retrieval call binding the contract method 0xaa953795.
//
// Solidity: function getSDCollateral() view returns(address)
func (_StaderConfig *StaderConfigSession) GetSDCollateral() (common.Address, error) {
	return _StaderConfig.Contract.GetSDCollateral(&_StaderConfig.CallOpts)
}

// GetSDCollateral is a free data retrieval call binding the contract method 0xaa953795.
//
// Solidity: function getSDCollateral() view returns(address)
func (_StaderConfig *StaderConfigCallerSession) GetSDCollateral() (common.Address, error) {
	return _StaderConfig.Contract.GetSDCollateral(&_StaderConfig.CallOpts)
}

// GetSocializingPoolCycleDuration is a free data retrieval call binding the contract method 0x1ca197a5.
//
// Solidity: function getSocializingPoolCycleDuration() view returns(uint256)
func (_StaderConfig *StaderConfigCaller) GetSocializingPoolCycleDuration(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _StaderConfig.contract.Call(opts, &out, "getSocializingPoolCycleDuration")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetSocializingPoolCycleDuration is a free data retrieval call binding the contract method 0x1ca197a5.
//
// Solidity: function getSocializingPoolCycleDuration() view returns(uint256)
func (_StaderConfig *StaderConfigSession) GetSocializingPoolCycleDuration() (*big.Int, error) {
	return _StaderConfig.Contract.GetSocializingPoolCycleDuration(&_StaderConfig.CallOpts)
}

// GetSocializingPoolCycleDuration is a free data retrieval call binding the contract method 0x1ca197a5.
//
// Solidity: function getSocializingPoolCycleDuration() view returns(uint256)
func (_StaderConfig *StaderConfigCallerSession) GetSocializingPoolCycleDuration() (*big.Int, error) {
	return _StaderConfig.Contract.GetSocializingPoolCycleDuration(&_StaderConfig.CallOpts)
}

// GetSocializingPoolOptInCoolingPeriod is a free data retrieval call binding the contract method 0x6e0fddfc.
//
// Solidity: function getSocializingPoolOptInCoolingPeriod() view returns(uint256)
func (_StaderConfig *StaderConfigCaller) GetSocializingPoolOptInCoolingPeriod(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _StaderConfig.contract.Call(opts, &out, "getSocializingPoolOptInCoolingPeriod")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetSocializingPoolOptInCoolingPeriod is a free data retrieval call binding the contract method 0x6e0fddfc.
//
// Solidity: function getSocializingPoolOptInCoolingPeriod() view returns(uint256)
func (_StaderConfig *StaderConfigSession) GetSocializingPoolOptInCoolingPeriod() (*big.Int, error) {
	return _StaderConfig.Contract.GetSocializingPoolOptInCoolingPeriod(&_StaderConfig.CallOpts)
}

// GetSocializingPoolOptInCoolingPeriod is a free data retrieval call binding the contract method 0x6e0fddfc.
//
// Solidity: function getSocializingPoolOptInCoolingPeriod() view returns(uint256)
func (_StaderConfig *StaderConfigCallerSession) GetSocializingPoolOptInCoolingPeriod() (*big.Int, error) {
	return _StaderConfig.Contract.GetSocializingPoolOptInCoolingPeriod(&_StaderConfig.CallOpts)
}

// GetStaderInsuranceFund is a free data retrieval call binding the contract method 0xb5cfee6c.
//
// Solidity: function getStaderInsuranceFund() view returns(address)
func (_StaderConfig *StaderConfigCaller) GetStaderInsuranceFund(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _StaderConfig.contract.Call(opts, &out, "getStaderInsuranceFund")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetStaderInsuranceFund is a free data retrieval call binding the contract method 0xb5cfee6c.
//
// Solidity: function getStaderInsuranceFund() view returns(address)
func (_StaderConfig *StaderConfigSession) GetStaderInsuranceFund() (common.Address, error) {
	return _StaderConfig.Contract.GetStaderInsuranceFund(&_StaderConfig.CallOpts)
}

// GetStaderInsuranceFund is a free data retrieval call binding the contract method 0xb5cfee6c.
//
// Solidity: function getStaderInsuranceFund() view returns(address)
func (_StaderConfig *StaderConfigCallerSession) GetStaderInsuranceFund() (common.Address, error) {
	return _StaderConfig.Contract.GetStaderInsuranceFund(&_StaderConfig.CallOpts)
}

// GetStaderOracle is a free data retrieval call binding the contract method 0xdefd024d.
//
// Solidity: function getStaderOracle() view returns(address)
func (_StaderConfig *StaderConfigCaller) GetStaderOracle(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _StaderConfig.contract.Call(opts, &out, "getStaderOracle")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetStaderOracle is a free data retrieval call binding the contract method 0xdefd024d.
//
// Solidity: function getStaderOracle() view returns(address)
func (_StaderConfig *StaderConfigSession) GetStaderOracle() (common.Address, error) {
	return _StaderConfig.Contract.GetStaderOracle(&_StaderConfig.CallOpts)
}

// GetStaderOracle is a free data retrieval call binding the contract method 0xdefd024d.
//
// Solidity: function getStaderOracle() view returns(address)
func (_StaderConfig *StaderConfigCallerSession) GetStaderOracle() (common.Address, error) {
	return _StaderConfig.Contract.GetStaderOracle(&_StaderConfig.CallOpts)
}

// GetStaderToken is a free data retrieval call binding the contract method 0xe069f714.
//
// Solidity: function getStaderToken() view returns(address)
func (_StaderConfig *StaderConfigCaller) GetStaderToken(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _StaderConfig.contract.Call(opts, &out, "getStaderToken")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetStaderToken is a free data retrieval call binding the contract method 0xe069f714.
//
// Solidity: function getStaderToken() view returns(address)
func (_StaderConfig *StaderConfigSession) GetStaderToken() (common.Address, error) {
	return _StaderConfig.Contract.GetStaderToken(&_StaderConfig.CallOpts)
}

// GetStaderToken is a free data retrieval call binding the contract method 0xe069f714.
//
// Solidity: function getStaderToken() view returns(address)
func (_StaderConfig *StaderConfigCallerSession) GetStaderToken() (common.Address, error) {
	return _StaderConfig.Contract.GetStaderToken(&_StaderConfig.CallOpts)
}

// GetStaderTreasury is a free data retrieval call binding the contract method 0x72ce78b0.
//
// Solidity: function getStaderTreasury() view returns(address)
func (_StaderConfig *StaderConfigCaller) GetStaderTreasury(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _StaderConfig.contract.Call(opts, &out, "getStaderTreasury")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetStaderTreasury is a free data retrieval call binding the contract method 0x72ce78b0.
//
// Solidity: function getStaderTreasury() view returns(address)
func (_StaderConfig *StaderConfigSession) GetStaderTreasury() (common.Address, error) {
	return _StaderConfig.Contract.GetStaderTreasury(&_StaderConfig.CallOpts)
}

// GetStaderTreasury is a free data retrieval call binding the contract method 0x72ce78b0.
//
// Solidity: function getStaderTreasury() view returns(address)
func (_StaderConfig *StaderConfigCallerSession) GetStaderTreasury() (common.Address, error) {
	return _StaderConfig.Contract.GetStaderTreasury(&_StaderConfig.CallOpts)
}

// GetStakePoolManager is a free data retrieval call binding the contract method 0x2ec5e018.
//
// Solidity: function getStakePoolManager() view returns(address)
func (_StaderConfig *StaderConfigCaller) GetStakePoolManager(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _StaderConfig.contract.Call(opts, &out, "getStakePoolManager")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetStakePoolManager is a free data retrieval call binding the contract method 0x2ec5e018.
//
// Solidity: function getStakePoolManager() view returns(address)
func (_StaderConfig *StaderConfigSession) GetStakePoolManager() (common.Address, error) {
	return _StaderConfig.Contract.GetStakePoolManager(&_StaderConfig.CallOpts)
}

// GetStakePoolManager is a free data retrieval call binding the contract method 0x2ec5e018.
//
// Solidity: function getStakePoolManager() view returns(address)
func (_StaderConfig *StaderConfigCallerSession) GetStakePoolManager() (common.Address, error) {
	return _StaderConfig.Contract.GetStakePoolManager(&_StaderConfig.CallOpts)
}

// GetStakedEthPerNode is a free data retrieval call binding the contract method 0xff387f3a.
//
// Solidity: function getStakedEthPerNode() view returns(uint256)
func (_StaderConfig *StaderConfigCaller) GetStakedEthPerNode(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _StaderConfig.contract.Call(opts, &out, "getStakedEthPerNode")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetStakedEthPerNode is a free data retrieval call binding the contract method 0xff387f3a.
//
// Solidity: function getStakedEthPerNode() view returns(uint256)
func (_StaderConfig *StaderConfigSession) GetStakedEthPerNode() (*big.Int, error) {
	return _StaderConfig.Contract.GetStakedEthPerNode(&_StaderConfig.CallOpts)
}

// GetStakedEthPerNode is a free data retrieval call binding the contract method 0xff387f3a.
//
// Solidity: function getStakedEthPerNode() view returns(uint256)
func (_StaderConfig *StaderConfigCallerSession) GetStakedEthPerNode() (*big.Int, error) {
	return _StaderConfig.Contract.GetStakedEthPerNode(&_StaderConfig.CallOpts)
}

// GetTotalFee is a free data retrieval call binding the contract method 0x7ae316d0.
//
// Solidity: function getTotalFee() view returns(uint256)
func (_StaderConfig *StaderConfigCaller) GetTotalFee(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _StaderConfig.contract.Call(opts, &out, "getTotalFee")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetTotalFee is a free data retrieval call binding the contract method 0x7ae316d0.
//
// Solidity: function getTotalFee() view returns(uint256)
func (_StaderConfig *StaderConfigSession) GetTotalFee() (*big.Int, error) {
	return _StaderConfig.Contract.GetTotalFee(&_StaderConfig.CallOpts)
}

// GetTotalFee is a free data retrieval call binding the contract method 0x7ae316d0.
//
// Solidity: function getTotalFee() view returns(uint256)
func (_StaderConfig *StaderConfigCallerSession) GetTotalFee() (*big.Int, error) {
	return _StaderConfig.Contract.GetTotalFee(&_StaderConfig.CallOpts)
}

// GetUserWithdrawManager is a free data retrieval call binding the contract method 0xecf170a8.
//
// Solidity: function getUserWithdrawManager() view returns(address)
func (_StaderConfig *StaderConfigCaller) GetUserWithdrawManager(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _StaderConfig.contract.Call(opts, &out, "getUserWithdrawManager")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetUserWithdrawManager is a free data retrieval call binding the contract method 0xecf170a8.
//
// Solidity: function getUserWithdrawManager() view returns(address)
func (_StaderConfig *StaderConfigSession) GetUserWithdrawManager() (common.Address, error) {
	return _StaderConfig.Contract.GetUserWithdrawManager(&_StaderConfig.CallOpts)
}

// GetUserWithdrawManager is a free data retrieval call binding the contract method 0xecf170a8.
//
// Solidity: function getUserWithdrawManager() view returns(address)
func (_StaderConfig *StaderConfigCallerSession) GetUserWithdrawManager() (common.Address, error) {
	return _StaderConfig.Contract.GetUserWithdrawManager(&_StaderConfig.CallOpts)
}

// GetValidatorWithdrawalVaultImplementation is a free data retrieval call binding the contract method 0x6d28ad1c.
//
// Solidity: function getValidatorWithdrawalVaultImplementation() view returns(address)
func (_StaderConfig *StaderConfigCaller) GetValidatorWithdrawalVaultImplementation(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _StaderConfig.contract.Call(opts, &out, "getValidatorWithdrawalVaultImplementation")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetValidatorWithdrawalVaultImplementation is a free data retrieval call binding the contract method 0x6d28ad1c.
//
// Solidity: function getValidatorWithdrawalVaultImplementation() view returns(address)
func (_StaderConfig *StaderConfigSession) GetValidatorWithdrawalVaultImplementation() (common.Address, error) {
	return _StaderConfig.Contract.GetValidatorWithdrawalVaultImplementation(&_StaderConfig.CallOpts)
}

// GetValidatorWithdrawalVaultImplementation is a free data retrieval call binding the contract method 0x6d28ad1c.
//
// Solidity: function getValidatorWithdrawalVaultImplementation() view returns(address)
func (_StaderConfig *StaderConfigCallerSession) GetValidatorWithdrawalVaultImplementation() (common.Address, error) {
	return _StaderConfig.Contract.GetValidatorWithdrawalVaultImplementation(&_StaderConfig.CallOpts)
}

// GetVaultFactory is a free data retrieval call binding the contract method 0x18bcb284.
//
// Solidity: function getVaultFactory() view returns(address)
func (_StaderConfig *StaderConfigCaller) GetVaultFactory(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _StaderConfig.contract.Call(opts, &out, "getVaultFactory")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetVaultFactory is a free data retrieval call binding the contract method 0x18bcb284.
//
// Solidity: function getVaultFactory() view returns(address)
func (_StaderConfig *StaderConfigSession) GetVaultFactory() (common.Address, error) {
	return _StaderConfig.Contract.GetVaultFactory(&_StaderConfig.CallOpts)
}

// GetVaultFactory is a free data retrieval call binding the contract method 0x18bcb284.
//
// Solidity: function getVaultFactory() view returns(address)
func (_StaderConfig *StaderConfigCallerSession) GetVaultFactory() (common.Address, error) {
	return _StaderConfig.Contract.GetVaultFactory(&_StaderConfig.CallOpts)
}

// GetWithdrawnKeyBatchSize is a free data retrieval call binding the contract method 0xb479a517.
//
// Solidity: function getWithdrawnKeyBatchSize() view returns(uint256)
func (_StaderConfig *StaderConfigCaller) GetWithdrawnKeyBatchSize(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _StaderConfig.contract.Call(opts, &out, "getWithdrawnKeyBatchSize")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetWithdrawnKeyBatchSize is a free data retrieval call binding the contract method 0xb479a517.
//
// Solidity: function getWithdrawnKeyBatchSize() view returns(uint256)
func (_StaderConfig *StaderConfigSession) GetWithdrawnKeyBatchSize() (*big.Int, error) {
	return _StaderConfig.Contract.GetWithdrawnKeyBatchSize(&_StaderConfig.CallOpts)
}

// GetWithdrawnKeyBatchSize is a free data retrieval call binding the contract method 0xb479a517.
//
// Solidity: function getWithdrawnKeyBatchSize() view returns(uint256)
func (_StaderConfig *StaderConfigCallerSession) GetWithdrawnKeyBatchSize() (*big.Int, error) {
	return _StaderConfig.Contract.GetWithdrawnKeyBatchSize(&_StaderConfig.CallOpts)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_StaderConfig *StaderConfigCaller) HasRole(opts *bind.CallOpts, role [32]byte, account common.Address) (bool, error) {
	var out []interface{}
	err := _StaderConfig.contract.Call(opts, &out, "hasRole", role, account)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_StaderConfig *StaderConfigSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _StaderConfig.Contract.HasRole(&_StaderConfig.CallOpts, role, account)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_StaderConfig *StaderConfigCallerSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _StaderConfig.Contract.HasRole(&_StaderConfig.CallOpts, role, account)
}

// OnlyManagerRole is a free data retrieval call binding the contract method 0x6240fb9c.
//
// Solidity: function onlyManagerRole(address account) view returns(bool)
func (_StaderConfig *StaderConfigCaller) OnlyManagerRole(opts *bind.CallOpts, account common.Address) (bool, error) {
	var out []interface{}
	err := _StaderConfig.contract.Call(opts, &out, "onlyManagerRole", account)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// OnlyManagerRole is a free data retrieval call binding the contract method 0x6240fb9c.
//
// Solidity: function onlyManagerRole(address account) view returns(bool)
func (_StaderConfig *StaderConfigSession) OnlyManagerRole(account common.Address) (bool, error) {
	return _StaderConfig.Contract.OnlyManagerRole(&_StaderConfig.CallOpts, account)
}

// OnlyManagerRole is a free data retrieval call binding the contract method 0x6240fb9c.
//
// Solidity: function onlyManagerRole(address account) view returns(bool)
func (_StaderConfig *StaderConfigCallerSession) OnlyManagerRole(account common.Address) (bool, error) {
	return _StaderConfig.Contract.OnlyManagerRole(&_StaderConfig.CallOpts, account)
}

// OnlyOperatorRole is a free data retrieval call binding the contract method 0x53f5713b.
//
// Solidity: function onlyOperatorRole(address account) view returns(bool)
func (_StaderConfig *StaderConfigCaller) OnlyOperatorRole(opts *bind.CallOpts, account common.Address) (bool, error) {
	var out []interface{}
	err := _StaderConfig.contract.Call(opts, &out, "onlyOperatorRole", account)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// OnlyOperatorRole is a free data retrieval call binding the contract method 0x53f5713b.
//
// Solidity: function onlyOperatorRole(address account) view returns(bool)
func (_StaderConfig *StaderConfigSession) OnlyOperatorRole(account common.Address) (bool, error) {
	return _StaderConfig.Contract.OnlyOperatorRole(&_StaderConfig.CallOpts, account)
}

// OnlyOperatorRole is a free data retrieval call binding the contract method 0x53f5713b.
//
// Solidity: function onlyOperatorRole(address account) view returns(bool)
func (_StaderConfig *StaderConfigCallerSession) OnlyOperatorRole(account common.Address) (bool, error) {
	return _StaderConfig.Contract.OnlyOperatorRole(&_StaderConfig.CallOpts, account)
}

// OnlyStaderContract is a free data retrieval call binding the contract method 0xb3123922.
//
// Solidity: function onlyStaderContract(address _addr, bytes32 _contractName) view returns(bool)
func (_StaderConfig *StaderConfigCaller) OnlyStaderContract(opts *bind.CallOpts, _addr common.Address, _contractName [32]byte) (bool, error) {
	var out []interface{}
	err := _StaderConfig.contract.Call(opts, &out, "onlyStaderContract", _addr, _contractName)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// OnlyStaderContract is a free data retrieval call binding the contract method 0xb3123922.
//
// Solidity: function onlyStaderContract(address _addr, bytes32 _contractName) view returns(bool)
func (_StaderConfig *StaderConfigSession) OnlyStaderContract(_addr common.Address, _contractName [32]byte) (bool, error) {
	return _StaderConfig.Contract.OnlyStaderContract(&_StaderConfig.CallOpts, _addr, _contractName)
}

// OnlyStaderContract is a free data retrieval call binding the contract method 0xb3123922.
//
// Solidity: function onlyStaderContract(address _addr, bytes32 _contractName) view returns(bool)
func (_StaderConfig *StaderConfigCallerSession) OnlyStaderContract(_addr common.Address, _contractName [32]byte) (bool, error) {
	return _StaderConfig.Contract.OnlyStaderContract(&_StaderConfig.CallOpts, _addr, _contractName)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_StaderConfig *StaderConfigCaller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _StaderConfig.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_StaderConfig *StaderConfigSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _StaderConfig.Contract.SupportsInterface(&_StaderConfig.CallOpts, interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_StaderConfig *StaderConfigCallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _StaderConfig.Contract.SupportsInterface(&_StaderConfig.CallOpts, interfaceId)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_StaderConfig *StaderConfigTransactor) GrantRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _StaderConfig.contract.Transact(opts, "grantRole", role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_StaderConfig *StaderConfigSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _StaderConfig.Contract.GrantRole(&_StaderConfig.TransactOpts, role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_StaderConfig *StaderConfigTransactorSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _StaderConfig.Contract.GrantRole(&_StaderConfig.TransactOpts, role, account)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address account) returns()
func (_StaderConfig *StaderConfigTransactor) RenounceRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _StaderConfig.contract.Transact(opts, "renounceRole", role, account)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address account) returns()
func (_StaderConfig *StaderConfigSession) RenounceRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _StaderConfig.Contract.RenounceRole(&_StaderConfig.TransactOpts, role, account)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address account) returns()
func (_StaderConfig *StaderConfigTransactorSession) RenounceRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _StaderConfig.Contract.RenounceRole(&_StaderConfig.TransactOpts, role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_StaderConfig *StaderConfigTransactor) RevokeRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _StaderConfig.contract.Transact(opts, "revokeRole", role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_StaderConfig *StaderConfigSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _StaderConfig.Contract.RevokeRole(&_StaderConfig.TransactOpts, role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_StaderConfig *StaderConfigTransactorSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _StaderConfig.Contract.RevokeRole(&_StaderConfig.TransactOpts, role, account)
}

// UpdateAdmin is a paid mutator transaction binding the contract method 0xe2f273bd.
//
// Solidity: function updateAdmin(address _admin) returns()
func (_StaderConfig *StaderConfigTransactor) UpdateAdmin(opts *bind.TransactOpts, _admin common.Address) (*types.Transaction, error) {
	return _StaderConfig.contract.Transact(opts, "updateAdmin", _admin)
}

// UpdateAdmin is a paid mutator transaction binding the contract method 0xe2f273bd.
//
// Solidity: function updateAdmin(address _admin) returns()
func (_StaderConfig *StaderConfigSession) UpdateAdmin(_admin common.Address) (*types.Transaction, error) {
	return _StaderConfig.Contract.UpdateAdmin(&_StaderConfig.TransactOpts, _admin)
}

// UpdateAdmin is a paid mutator transaction binding the contract method 0xe2f273bd.
//
// Solidity: function updateAdmin(address _admin) returns()
func (_StaderConfig *StaderConfigTransactorSession) UpdateAdmin(_admin common.Address) (*types.Transaction, error) {
	return _StaderConfig.Contract.UpdateAdmin(&_StaderConfig.TransactOpts, _admin)
}

// UpdateAuctionContract is a paid mutator transaction binding the contract method 0x121669f1.
//
// Solidity: function updateAuctionContract(address _auctionContract) returns()
func (_StaderConfig *StaderConfigTransactor) UpdateAuctionContract(opts *bind.TransactOpts, _auctionContract common.Address) (*types.Transaction, error) {
	return _StaderConfig.contract.Transact(opts, "updateAuctionContract", _auctionContract)
}

// UpdateAuctionContract is a paid mutator transaction binding the contract method 0x121669f1.
//
// Solidity: function updateAuctionContract(address _auctionContract) returns()
func (_StaderConfig *StaderConfigSession) UpdateAuctionContract(_auctionContract common.Address) (*types.Transaction, error) {
	return _StaderConfig.Contract.UpdateAuctionContract(&_StaderConfig.TransactOpts, _auctionContract)
}

// UpdateAuctionContract is a paid mutator transaction binding the contract method 0x121669f1.
//
// Solidity: function updateAuctionContract(address _auctionContract) returns()
func (_StaderConfig *StaderConfigTransactorSession) UpdateAuctionContract(_auctionContract common.Address) (*types.Transaction, error) {
	return _StaderConfig.Contract.UpdateAuctionContract(&_StaderConfig.TransactOpts, _auctionContract)
}

// UpdateETHBalancePORFeedProxy is a paid mutator transaction binding the contract method 0x403efe7f.
//
// Solidity: function updateETHBalancePORFeedProxy(address _ethBalanceProxy) returns()
func (_StaderConfig *StaderConfigTransactor) UpdateETHBalancePORFeedProxy(opts *bind.TransactOpts, _ethBalanceProxy common.Address) (*types.Transaction, error) {
	return _StaderConfig.contract.Transact(opts, "updateETHBalancePORFeedProxy", _ethBalanceProxy)
}

// UpdateETHBalancePORFeedProxy is a paid mutator transaction binding the contract method 0x403efe7f.
//
// Solidity: function updateETHBalancePORFeedProxy(address _ethBalanceProxy) returns()
func (_StaderConfig *StaderConfigSession) UpdateETHBalancePORFeedProxy(_ethBalanceProxy common.Address) (*types.Transaction, error) {
	return _StaderConfig.Contract.UpdateETHBalancePORFeedProxy(&_StaderConfig.TransactOpts, _ethBalanceProxy)
}

// UpdateETHBalancePORFeedProxy is a paid mutator transaction binding the contract method 0x403efe7f.
//
// Solidity: function updateETHBalancePORFeedProxy(address _ethBalanceProxy) returns()
func (_StaderConfig *StaderConfigTransactorSession) UpdateETHBalancePORFeedProxy(_ethBalanceProxy common.Address) (*types.Transaction, error) {
	return _StaderConfig.Contract.UpdateETHBalancePORFeedProxy(&_StaderConfig.TransactOpts, _ethBalanceProxy)
}

// UpdateETHXSupplyPORFeedProxy is a paid mutator transaction binding the contract method 0x72195b3e.
//
// Solidity: function updateETHXSupplyPORFeedProxy(address _ethXSupplyProxy) returns()
func (_StaderConfig *StaderConfigTransactor) UpdateETHXSupplyPORFeedProxy(opts *bind.TransactOpts, _ethXSupplyProxy common.Address) (*types.Transaction, error) {
	return _StaderConfig.contract.Transact(opts, "updateETHXSupplyPORFeedProxy", _ethXSupplyProxy)
}

// UpdateETHXSupplyPORFeedProxy is a paid mutator transaction binding the contract method 0x72195b3e.
//
// Solidity: function updateETHXSupplyPORFeedProxy(address _ethXSupplyProxy) returns()
func (_StaderConfig *StaderConfigSession) UpdateETHXSupplyPORFeedProxy(_ethXSupplyProxy common.Address) (*types.Transaction, error) {
	return _StaderConfig.Contract.UpdateETHXSupplyPORFeedProxy(&_StaderConfig.TransactOpts, _ethXSupplyProxy)
}

// UpdateETHXSupplyPORFeedProxy is a paid mutator transaction binding the contract method 0x72195b3e.
//
// Solidity: function updateETHXSupplyPORFeedProxy(address _ethXSupplyProxy) returns()
func (_StaderConfig *StaderConfigTransactorSession) UpdateETHXSupplyPORFeedProxy(_ethXSupplyProxy common.Address) (*types.Transaction, error) {
	return _StaderConfig.Contract.UpdateETHXSupplyPORFeedProxy(&_StaderConfig.TransactOpts, _ethXSupplyProxy)
}

// UpdateETHxToken is a paid mutator transaction binding the contract method 0xb9894a11.
//
// Solidity: function updateETHxToken(address _ethX) returns()
func (_StaderConfig *StaderConfigTransactor) UpdateETHxToken(opts *bind.TransactOpts, _ethX common.Address) (*types.Transaction, error) {
	return _StaderConfig.contract.Transact(opts, "updateETHxToken", _ethX)
}

// UpdateETHxToken is a paid mutator transaction binding the contract method 0xb9894a11.
//
// Solidity: function updateETHxToken(address _ethX) returns()
func (_StaderConfig *StaderConfigSession) UpdateETHxToken(_ethX common.Address) (*types.Transaction, error) {
	return _StaderConfig.Contract.UpdateETHxToken(&_StaderConfig.TransactOpts, _ethX)
}

// UpdateETHxToken is a paid mutator transaction binding the contract method 0xb9894a11.
//
// Solidity: function updateETHxToken(address _ethX) returns()
func (_StaderConfig *StaderConfigTransactorSession) UpdateETHxToken(_ethX common.Address) (*types.Transaction, error) {
	return _StaderConfig.Contract.UpdateETHxToken(&_StaderConfig.TransactOpts, _ethX)
}

// UpdateMaxDepositAmount is a paid mutator transaction binding the contract method 0x0945d42c.
//
// Solidity: function updateMaxDepositAmount(uint256 _maxDepositAmount) returns()
func (_StaderConfig *StaderConfigTransactor) UpdateMaxDepositAmount(opts *bind.TransactOpts, _maxDepositAmount *big.Int) (*types.Transaction, error) {
	return _StaderConfig.contract.Transact(opts, "updateMaxDepositAmount", _maxDepositAmount)
}

// UpdateMaxDepositAmount is a paid mutator transaction binding the contract method 0x0945d42c.
//
// Solidity: function updateMaxDepositAmount(uint256 _maxDepositAmount) returns()
func (_StaderConfig *StaderConfigSession) UpdateMaxDepositAmount(_maxDepositAmount *big.Int) (*types.Transaction, error) {
	return _StaderConfig.Contract.UpdateMaxDepositAmount(&_StaderConfig.TransactOpts, _maxDepositAmount)
}

// UpdateMaxDepositAmount is a paid mutator transaction binding the contract method 0x0945d42c.
//
// Solidity: function updateMaxDepositAmount(uint256 _maxDepositAmount) returns()
func (_StaderConfig *StaderConfigTransactorSession) UpdateMaxDepositAmount(_maxDepositAmount *big.Int) (*types.Transaction, error) {
	return _StaderConfig.Contract.UpdateMaxDepositAmount(&_StaderConfig.TransactOpts, _maxDepositAmount)
}

// UpdateMaxWithdrawAmount is a paid mutator transaction binding the contract method 0x5b9cc8b1.
//
// Solidity: function updateMaxWithdrawAmount(uint256 _maxWithdrawAmount) returns()
func (_StaderConfig *StaderConfigTransactor) UpdateMaxWithdrawAmount(opts *bind.TransactOpts, _maxWithdrawAmount *big.Int) (*types.Transaction, error) {
	return _StaderConfig.contract.Transact(opts, "updateMaxWithdrawAmount", _maxWithdrawAmount)
}

// UpdateMaxWithdrawAmount is a paid mutator transaction binding the contract method 0x5b9cc8b1.
//
// Solidity: function updateMaxWithdrawAmount(uint256 _maxWithdrawAmount) returns()
func (_StaderConfig *StaderConfigSession) UpdateMaxWithdrawAmount(_maxWithdrawAmount *big.Int) (*types.Transaction, error) {
	return _StaderConfig.Contract.UpdateMaxWithdrawAmount(&_StaderConfig.TransactOpts, _maxWithdrawAmount)
}

// UpdateMaxWithdrawAmount is a paid mutator transaction binding the contract method 0x5b9cc8b1.
//
// Solidity: function updateMaxWithdrawAmount(uint256 _maxWithdrawAmount) returns()
func (_StaderConfig *StaderConfigTransactorSession) UpdateMaxWithdrawAmount(_maxWithdrawAmount *big.Int) (*types.Transaction, error) {
	return _StaderConfig.Contract.UpdateMaxWithdrawAmount(&_StaderConfig.TransactOpts, _maxWithdrawAmount)
}

// UpdateMinBlockDelayToFinalizeWithdrawRequest is a paid mutator transaction binding the contract method 0x723b732c.
//
// Solidity: function updateMinBlockDelayToFinalizeWithdrawRequest(uint256 _minBlockDelay) returns()
func (_StaderConfig *StaderConfigTransactor) UpdateMinBlockDelayToFinalizeWithdrawRequest(opts *bind.TransactOpts, _minBlockDelay *big.Int) (*types.Transaction, error) {
	return _StaderConfig.contract.Transact(opts, "updateMinBlockDelayToFinalizeWithdrawRequest", _minBlockDelay)
}

// UpdateMinBlockDelayToFinalizeWithdrawRequest is a paid mutator transaction binding the contract method 0x723b732c.
//
// Solidity: function updateMinBlockDelayToFinalizeWithdrawRequest(uint256 _minBlockDelay) returns()
func (_StaderConfig *StaderConfigSession) UpdateMinBlockDelayToFinalizeWithdrawRequest(_minBlockDelay *big.Int) (*types.Transaction, error) {
	return _StaderConfig.Contract.UpdateMinBlockDelayToFinalizeWithdrawRequest(&_StaderConfig.TransactOpts, _minBlockDelay)
}

// UpdateMinBlockDelayToFinalizeWithdrawRequest is a paid mutator transaction binding the contract method 0x723b732c.
//
// Solidity: function updateMinBlockDelayToFinalizeWithdrawRequest(uint256 _minBlockDelay) returns()
func (_StaderConfig *StaderConfigTransactorSession) UpdateMinBlockDelayToFinalizeWithdrawRequest(_minBlockDelay *big.Int) (*types.Transaction, error) {
	return _StaderConfig.Contract.UpdateMinBlockDelayToFinalizeWithdrawRequest(&_StaderConfig.TransactOpts, _minBlockDelay)
}

// UpdateMinDepositAmount is a paid mutator transaction binding the contract method 0x84780205.
//
// Solidity: function updateMinDepositAmount(uint256 _minDepositAmount) returns()
func (_StaderConfig *StaderConfigTransactor) UpdateMinDepositAmount(opts *bind.TransactOpts, _minDepositAmount *big.Int) (*types.Transaction, error) {
	return _StaderConfig.contract.Transact(opts, "updateMinDepositAmount", _minDepositAmount)
}

// UpdateMinDepositAmount is a paid mutator transaction binding the contract method 0x84780205.
//
// Solidity: function updateMinDepositAmount(uint256 _minDepositAmount) returns()
func (_StaderConfig *StaderConfigSession) UpdateMinDepositAmount(_minDepositAmount *big.Int) (*types.Transaction, error) {
	return _StaderConfig.Contract.UpdateMinDepositAmount(&_StaderConfig.TransactOpts, _minDepositAmount)
}

// UpdateMinDepositAmount is a paid mutator transaction binding the contract method 0x84780205.
//
// Solidity: function updateMinDepositAmount(uint256 _minDepositAmount) returns()
func (_StaderConfig *StaderConfigTransactorSession) UpdateMinDepositAmount(_minDepositAmount *big.Int) (*types.Transaction, error) {
	return _StaderConfig.Contract.UpdateMinDepositAmount(&_StaderConfig.TransactOpts, _minDepositAmount)
}

// UpdateMinWithdrawAmount is a paid mutator transaction binding the contract method 0xff4f3546.
//
// Solidity: function updateMinWithdrawAmount(uint256 _minWithdrawAmount) returns()
func (_StaderConfig *StaderConfigTransactor) UpdateMinWithdrawAmount(opts *bind.TransactOpts, _minWithdrawAmount *big.Int) (*types.Transaction, error) {
	return _StaderConfig.contract.Transact(opts, "updateMinWithdrawAmount", _minWithdrawAmount)
}

// UpdateMinWithdrawAmount is a paid mutator transaction binding the contract method 0xff4f3546.
//
// Solidity: function updateMinWithdrawAmount(uint256 _minWithdrawAmount) returns()
func (_StaderConfig *StaderConfigSession) UpdateMinWithdrawAmount(_minWithdrawAmount *big.Int) (*types.Transaction, error) {
	return _StaderConfig.Contract.UpdateMinWithdrawAmount(&_StaderConfig.TransactOpts, _minWithdrawAmount)
}

// UpdateMinWithdrawAmount is a paid mutator transaction binding the contract method 0xff4f3546.
//
// Solidity: function updateMinWithdrawAmount(uint256 _minWithdrawAmount) returns()
func (_StaderConfig *StaderConfigTransactorSession) UpdateMinWithdrawAmount(_minWithdrawAmount *big.Int) (*types.Transaction, error) {
	return _StaderConfig.Contract.UpdateMinWithdrawAmount(&_StaderConfig.TransactOpts, _minWithdrawAmount)
}

// UpdateNodeELRewardImplementation is a paid mutator transaction binding the contract method 0x5be6ce69.
//
// Solidity: function updateNodeELRewardImplementation(address _nodeELRewardVaultImpl) returns()
func (_StaderConfig *StaderConfigTransactor) UpdateNodeELRewardImplementation(opts *bind.TransactOpts, _nodeELRewardVaultImpl common.Address) (*types.Transaction, error) {
	return _StaderConfig.contract.Transact(opts, "updateNodeELRewardImplementation", _nodeELRewardVaultImpl)
}

// UpdateNodeELRewardImplementation is a paid mutator transaction binding the contract method 0x5be6ce69.
//
// Solidity: function updateNodeELRewardImplementation(address _nodeELRewardVaultImpl) returns()
func (_StaderConfig *StaderConfigSession) UpdateNodeELRewardImplementation(_nodeELRewardVaultImpl common.Address) (*types.Transaction, error) {
	return _StaderConfig.Contract.UpdateNodeELRewardImplementation(&_StaderConfig.TransactOpts, _nodeELRewardVaultImpl)
}

// UpdateNodeELRewardImplementation is a paid mutator transaction binding the contract method 0x5be6ce69.
//
// Solidity: function updateNodeELRewardImplementation(address _nodeELRewardVaultImpl) returns()
func (_StaderConfig *StaderConfigTransactorSession) UpdateNodeELRewardImplementation(_nodeELRewardVaultImpl common.Address) (*types.Transaction, error) {
	return _StaderConfig.Contract.UpdateNodeELRewardImplementation(&_StaderConfig.TransactOpts, _nodeELRewardVaultImpl)
}

// UpdateOperatorRewardsCollector is a paid mutator transaction binding the contract method 0x1de03db8.
//
// Solidity: function updateOperatorRewardsCollector(address _operatorRewardsCollector) returns()
func (_StaderConfig *StaderConfigTransactor) UpdateOperatorRewardsCollector(opts *bind.TransactOpts, _operatorRewardsCollector common.Address) (*types.Transaction, error) {
	return _StaderConfig.contract.Transact(opts, "updateOperatorRewardsCollector", _operatorRewardsCollector)
}

// UpdateOperatorRewardsCollector is a paid mutator transaction binding the contract method 0x1de03db8.
//
// Solidity: function updateOperatorRewardsCollector(address _operatorRewardsCollector) returns()
func (_StaderConfig *StaderConfigSession) UpdateOperatorRewardsCollector(_operatorRewardsCollector common.Address) (*types.Transaction, error) {
	return _StaderConfig.Contract.UpdateOperatorRewardsCollector(&_StaderConfig.TransactOpts, _operatorRewardsCollector)
}

// UpdateOperatorRewardsCollector is a paid mutator transaction binding the contract method 0x1de03db8.
//
// Solidity: function updateOperatorRewardsCollector(address _operatorRewardsCollector) returns()
func (_StaderConfig *StaderConfigTransactorSession) UpdateOperatorRewardsCollector(_operatorRewardsCollector common.Address) (*types.Transaction, error) {
	return _StaderConfig.Contract.UpdateOperatorRewardsCollector(&_StaderConfig.TransactOpts, _operatorRewardsCollector)
}

// UpdatePenaltyContract is a paid mutator transaction binding the contract method 0xf83c7787.
//
// Solidity: function updatePenaltyContract(address _penaltyContract) returns()
func (_StaderConfig *StaderConfigTransactor) UpdatePenaltyContract(opts *bind.TransactOpts, _penaltyContract common.Address) (*types.Transaction, error) {
	return _StaderConfig.contract.Transact(opts, "updatePenaltyContract", _penaltyContract)
}

// UpdatePenaltyContract is a paid mutator transaction binding the contract method 0xf83c7787.
//
// Solidity: function updatePenaltyContract(address _penaltyContract) returns()
func (_StaderConfig *StaderConfigSession) UpdatePenaltyContract(_penaltyContract common.Address) (*types.Transaction, error) {
	return _StaderConfig.Contract.UpdatePenaltyContract(&_StaderConfig.TransactOpts, _penaltyContract)
}

// UpdatePenaltyContract is a paid mutator transaction binding the contract method 0xf83c7787.
//
// Solidity: function updatePenaltyContract(address _penaltyContract) returns()
func (_StaderConfig *StaderConfigTransactorSession) UpdatePenaltyContract(_penaltyContract common.Address) (*types.Transaction, error) {
	return _StaderConfig.Contract.UpdatePenaltyContract(&_StaderConfig.TransactOpts, _penaltyContract)
}

// UpdatePermissionedNodeRegistry is a paid mutator transaction binding the contract method 0x3c128dad.
//
// Solidity: function updatePermissionedNodeRegistry(address _permissionedNodeRegistry) returns()
func (_StaderConfig *StaderConfigTransactor) UpdatePermissionedNodeRegistry(opts *bind.TransactOpts, _permissionedNodeRegistry common.Address) (*types.Transaction, error) {
	return _StaderConfig.contract.Transact(opts, "updatePermissionedNodeRegistry", _permissionedNodeRegistry)
}

// UpdatePermissionedNodeRegistry is a paid mutator transaction binding the contract method 0x3c128dad.
//
// Solidity: function updatePermissionedNodeRegistry(address _permissionedNodeRegistry) returns()
func (_StaderConfig *StaderConfigSession) UpdatePermissionedNodeRegistry(_permissionedNodeRegistry common.Address) (*types.Transaction, error) {
	return _StaderConfig.Contract.UpdatePermissionedNodeRegistry(&_StaderConfig.TransactOpts, _permissionedNodeRegistry)
}

// UpdatePermissionedNodeRegistry is a paid mutator transaction binding the contract method 0x3c128dad.
//
// Solidity: function updatePermissionedNodeRegistry(address _permissionedNodeRegistry) returns()
func (_StaderConfig *StaderConfigTransactorSession) UpdatePermissionedNodeRegistry(_permissionedNodeRegistry common.Address) (*types.Transaction, error) {
	return _StaderConfig.Contract.UpdatePermissionedNodeRegistry(&_StaderConfig.TransactOpts, _permissionedNodeRegistry)
}

// UpdatePermissionedPool is a paid mutator transaction binding the contract method 0xb549dbff.
//
// Solidity: function updatePermissionedPool(address _permissionedPool) returns()
func (_StaderConfig *StaderConfigTransactor) UpdatePermissionedPool(opts *bind.TransactOpts, _permissionedPool common.Address) (*types.Transaction, error) {
	return _StaderConfig.contract.Transact(opts, "updatePermissionedPool", _permissionedPool)
}

// UpdatePermissionedPool is a paid mutator transaction binding the contract method 0xb549dbff.
//
// Solidity: function updatePermissionedPool(address _permissionedPool) returns()
func (_StaderConfig *StaderConfigSession) UpdatePermissionedPool(_permissionedPool common.Address) (*types.Transaction, error) {
	return _StaderConfig.Contract.UpdatePermissionedPool(&_StaderConfig.TransactOpts, _permissionedPool)
}

// UpdatePermissionedPool is a paid mutator transaction binding the contract method 0xb549dbff.
//
// Solidity: function updatePermissionedPool(address _permissionedPool) returns()
func (_StaderConfig *StaderConfigTransactorSession) UpdatePermissionedPool(_permissionedPool common.Address) (*types.Transaction, error) {
	return _StaderConfig.Contract.UpdatePermissionedPool(&_StaderConfig.TransactOpts, _permissionedPool)
}

// UpdatePermissionedSocializingPool is a paid mutator transaction binding the contract method 0xe4f59b6c.
//
// Solidity: function updatePermissionedSocializingPool(address _permissionedSocializePool) returns()
func (_StaderConfig *StaderConfigTransactor) UpdatePermissionedSocializingPool(opts *bind.TransactOpts, _permissionedSocializePool common.Address) (*types.Transaction, error) {
	return _StaderConfig.contract.Transact(opts, "updatePermissionedSocializingPool", _permissionedSocializePool)
}

// UpdatePermissionedSocializingPool is a paid mutator transaction binding the contract method 0xe4f59b6c.
//
// Solidity: function updatePermissionedSocializingPool(address _permissionedSocializePool) returns()
func (_StaderConfig *StaderConfigSession) UpdatePermissionedSocializingPool(_permissionedSocializePool common.Address) (*types.Transaction, error) {
	return _StaderConfig.Contract.UpdatePermissionedSocializingPool(&_StaderConfig.TransactOpts, _permissionedSocializePool)
}

// UpdatePermissionedSocializingPool is a paid mutator transaction binding the contract method 0xe4f59b6c.
//
// Solidity: function updatePermissionedSocializingPool(address _permissionedSocializePool) returns()
func (_StaderConfig *StaderConfigTransactorSession) UpdatePermissionedSocializingPool(_permissionedSocializePool common.Address) (*types.Transaction, error) {
	return _StaderConfig.Contract.UpdatePermissionedSocializingPool(&_StaderConfig.TransactOpts, _permissionedSocializePool)
}

// UpdatePermissionlessNodeRegistry is a paid mutator transaction binding the contract method 0xf63718e7.
//
// Solidity: function updatePermissionlessNodeRegistry(address _permissionlessNodeRegistry) returns()
func (_StaderConfig *StaderConfigTransactor) UpdatePermissionlessNodeRegistry(opts *bind.TransactOpts, _permissionlessNodeRegistry common.Address) (*types.Transaction, error) {
	return _StaderConfig.contract.Transact(opts, "updatePermissionlessNodeRegistry", _permissionlessNodeRegistry)
}

// UpdatePermissionlessNodeRegistry is a paid mutator transaction binding the contract method 0xf63718e7.
//
// Solidity: function updatePermissionlessNodeRegistry(address _permissionlessNodeRegistry) returns()
func (_StaderConfig *StaderConfigSession) UpdatePermissionlessNodeRegistry(_permissionlessNodeRegistry common.Address) (*types.Transaction, error) {
	return _StaderConfig.Contract.UpdatePermissionlessNodeRegistry(&_StaderConfig.TransactOpts, _permissionlessNodeRegistry)
}

// UpdatePermissionlessNodeRegistry is a paid mutator transaction binding the contract method 0xf63718e7.
//
// Solidity: function updatePermissionlessNodeRegistry(address _permissionlessNodeRegistry) returns()
func (_StaderConfig *StaderConfigTransactorSession) UpdatePermissionlessNodeRegistry(_permissionlessNodeRegistry common.Address) (*types.Transaction, error) {
	return _StaderConfig.Contract.UpdatePermissionlessNodeRegistry(&_StaderConfig.TransactOpts, _permissionlessNodeRegistry)
}

// UpdatePermissionlessPool is a paid mutator transaction binding the contract method 0xc20573c1.
//
// Solidity: function updatePermissionlessPool(address _permissionlessPool) returns()
func (_StaderConfig *StaderConfigTransactor) UpdatePermissionlessPool(opts *bind.TransactOpts, _permissionlessPool common.Address) (*types.Transaction, error) {
	return _StaderConfig.contract.Transact(opts, "updatePermissionlessPool", _permissionlessPool)
}

// UpdatePermissionlessPool is a paid mutator transaction binding the contract method 0xc20573c1.
//
// Solidity: function updatePermissionlessPool(address _permissionlessPool) returns()
func (_StaderConfig *StaderConfigSession) UpdatePermissionlessPool(_permissionlessPool common.Address) (*types.Transaction, error) {
	return _StaderConfig.Contract.UpdatePermissionlessPool(&_StaderConfig.TransactOpts, _permissionlessPool)
}

// UpdatePermissionlessPool is a paid mutator transaction binding the contract method 0xc20573c1.
//
// Solidity: function updatePermissionlessPool(address _permissionlessPool) returns()
func (_StaderConfig *StaderConfigTransactorSession) UpdatePermissionlessPool(_permissionlessPool common.Address) (*types.Transaction, error) {
	return _StaderConfig.Contract.UpdatePermissionlessPool(&_StaderConfig.TransactOpts, _permissionlessPool)
}

// UpdatePermissionlessSocializingPool is a paid mutator transaction binding the contract method 0x1049e32e.
//
// Solidity: function updatePermissionlessSocializingPool(address _permissionlessSocializePool) returns()
func (_StaderConfig *StaderConfigTransactor) UpdatePermissionlessSocializingPool(opts *bind.TransactOpts, _permissionlessSocializePool common.Address) (*types.Transaction, error) {
	return _StaderConfig.contract.Transact(opts, "updatePermissionlessSocializingPool", _permissionlessSocializePool)
}

// UpdatePermissionlessSocializingPool is a paid mutator transaction binding the contract method 0x1049e32e.
//
// Solidity: function updatePermissionlessSocializingPool(address _permissionlessSocializePool) returns()
func (_StaderConfig *StaderConfigSession) UpdatePermissionlessSocializingPool(_permissionlessSocializePool common.Address) (*types.Transaction, error) {
	return _StaderConfig.Contract.UpdatePermissionlessSocializingPool(&_StaderConfig.TransactOpts, _permissionlessSocializePool)
}

// UpdatePermissionlessSocializingPool is a paid mutator transaction binding the contract method 0x1049e32e.
//
// Solidity: function updatePermissionlessSocializingPool(address _permissionlessSocializePool) returns()
func (_StaderConfig *StaderConfigTransactorSession) UpdatePermissionlessSocializingPool(_permissionlessSocializePool common.Address) (*types.Transaction, error) {
	return _StaderConfig.Contract.UpdatePermissionlessSocializingPool(&_StaderConfig.TransactOpts, _permissionlessSocializePool)
}

// UpdatePoolSelector is a paid mutator transaction binding the contract method 0x047cb439.
//
// Solidity: function updatePoolSelector(address _poolSelector) returns()
func (_StaderConfig *StaderConfigTransactor) UpdatePoolSelector(opts *bind.TransactOpts, _poolSelector common.Address) (*types.Transaction, error) {
	return _StaderConfig.contract.Transact(opts, "updatePoolSelector", _poolSelector)
}

// UpdatePoolSelector is a paid mutator transaction binding the contract method 0x047cb439.
//
// Solidity: function updatePoolSelector(address _poolSelector) returns()
func (_StaderConfig *StaderConfigSession) UpdatePoolSelector(_poolSelector common.Address) (*types.Transaction, error) {
	return _StaderConfig.Contract.UpdatePoolSelector(&_StaderConfig.TransactOpts, _poolSelector)
}

// UpdatePoolSelector is a paid mutator transaction binding the contract method 0x047cb439.
//
// Solidity: function updatePoolSelector(address _poolSelector) returns()
func (_StaderConfig *StaderConfigTransactorSession) UpdatePoolSelector(_poolSelector common.Address) (*types.Transaction, error) {
	return _StaderConfig.Contract.UpdatePoolSelector(&_StaderConfig.TransactOpts, _poolSelector)
}

// UpdatePoolUtils is a paid mutator transaction binding the contract method 0x2651644c.
//
// Solidity: function updatePoolUtils(address _poolUtils) returns()
func (_StaderConfig *StaderConfigTransactor) UpdatePoolUtils(opts *bind.TransactOpts, _poolUtils common.Address) (*types.Transaction, error) {
	return _StaderConfig.contract.Transact(opts, "updatePoolUtils", _poolUtils)
}

// UpdatePoolUtils is a paid mutator transaction binding the contract method 0x2651644c.
//
// Solidity: function updatePoolUtils(address _poolUtils) returns()
func (_StaderConfig *StaderConfigSession) UpdatePoolUtils(_poolUtils common.Address) (*types.Transaction, error) {
	return _StaderConfig.Contract.UpdatePoolUtils(&_StaderConfig.TransactOpts, _poolUtils)
}

// UpdatePoolUtils is a paid mutator transaction binding the contract method 0x2651644c.
//
// Solidity: function updatePoolUtils(address _poolUtils) returns()
func (_StaderConfig *StaderConfigTransactorSession) UpdatePoolUtils(_poolUtils common.Address) (*types.Transaction, error) {
	return _StaderConfig.Contract.UpdatePoolUtils(&_StaderConfig.TransactOpts, _poolUtils)
}

// UpdateRewardsThreshold is a paid mutator transaction binding the contract method 0x572c686a.
//
// Solidity: function updateRewardsThreshold(uint256 _rewardsThreshold) returns()
func (_StaderConfig *StaderConfigTransactor) UpdateRewardsThreshold(opts *bind.TransactOpts, _rewardsThreshold *big.Int) (*types.Transaction, error) {
	return _StaderConfig.contract.Transact(opts, "updateRewardsThreshold", _rewardsThreshold)
}

// UpdateRewardsThreshold is a paid mutator transaction binding the contract method 0x572c686a.
//
// Solidity: function updateRewardsThreshold(uint256 _rewardsThreshold) returns()
func (_StaderConfig *StaderConfigSession) UpdateRewardsThreshold(_rewardsThreshold *big.Int) (*types.Transaction, error) {
	return _StaderConfig.Contract.UpdateRewardsThreshold(&_StaderConfig.TransactOpts, _rewardsThreshold)
}

// UpdateRewardsThreshold is a paid mutator transaction binding the contract method 0x572c686a.
//
// Solidity: function updateRewardsThreshold(uint256 _rewardsThreshold) returns()
func (_StaderConfig *StaderConfigTransactorSession) UpdateRewardsThreshold(_rewardsThreshold *big.Int) (*types.Transaction, error) {
	return _StaderConfig.Contract.UpdateRewardsThreshold(&_StaderConfig.TransactOpts, _rewardsThreshold)
}

// UpdateSDCollateral is a paid mutator transaction binding the contract method 0x34d17d74.
//
// Solidity: function updateSDCollateral(address _sdCollateral) returns()
func (_StaderConfig *StaderConfigTransactor) UpdateSDCollateral(opts *bind.TransactOpts, _sdCollateral common.Address) (*types.Transaction, error) {
	return _StaderConfig.contract.Transact(opts, "updateSDCollateral", _sdCollateral)
}

// UpdateSDCollateral is a paid mutator transaction binding the contract method 0x34d17d74.
//
// Solidity: function updateSDCollateral(address _sdCollateral) returns()
func (_StaderConfig *StaderConfigSession) UpdateSDCollateral(_sdCollateral common.Address) (*types.Transaction, error) {
	return _StaderConfig.Contract.UpdateSDCollateral(&_StaderConfig.TransactOpts, _sdCollateral)
}

// UpdateSDCollateral is a paid mutator transaction binding the contract method 0x34d17d74.
//
// Solidity: function updateSDCollateral(address _sdCollateral) returns()
func (_StaderConfig *StaderConfigTransactorSession) UpdateSDCollateral(_sdCollateral common.Address) (*types.Transaction, error) {
	return _StaderConfig.Contract.UpdateSDCollateral(&_StaderConfig.TransactOpts, _sdCollateral)
}

// UpdateSocializingPoolCycleDuration is a paid mutator transaction binding the contract method 0x6870bb2b.
//
// Solidity: function updateSocializingPoolCycleDuration(uint256 _socializingPoolCycleDuration) returns()
func (_StaderConfig *StaderConfigTransactor) UpdateSocializingPoolCycleDuration(opts *bind.TransactOpts, _socializingPoolCycleDuration *big.Int) (*types.Transaction, error) {
	return _StaderConfig.contract.Transact(opts, "updateSocializingPoolCycleDuration", _socializingPoolCycleDuration)
}

// UpdateSocializingPoolCycleDuration is a paid mutator transaction binding the contract method 0x6870bb2b.
//
// Solidity: function updateSocializingPoolCycleDuration(uint256 _socializingPoolCycleDuration) returns()
func (_StaderConfig *StaderConfigSession) UpdateSocializingPoolCycleDuration(_socializingPoolCycleDuration *big.Int) (*types.Transaction, error) {
	return _StaderConfig.Contract.UpdateSocializingPoolCycleDuration(&_StaderConfig.TransactOpts, _socializingPoolCycleDuration)
}

// UpdateSocializingPoolCycleDuration is a paid mutator transaction binding the contract method 0x6870bb2b.
//
// Solidity: function updateSocializingPoolCycleDuration(uint256 _socializingPoolCycleDuration) returns()
func (_StaderConfig *StaderConfigTransactorSession) UpdateSocializingPoolCycleDuration(_socializingPoolCycleDuration *big.Int) (*types.Transaction, error) {
	return _StaderConfig.Contract.UpdateSocializingPoolCycleDuration(&_StaderConfig.TransactOpts, _socializingPoolCycleDuration)
}

// UpdateSocializingPoolOptInCoolingPeriod is a paid mutator transaction binding the contract method 0x8a4cfb58.
//
// Solidity: function updateSocializingPoolOptInCoolingPeriod(uint256 _SocializePoolOptInCoolingPeriod) returns()
func (_StaderConfig *StaderConfigTransactor) UpdateSocializingPoolOptInCoolingPeriod(opts *bind.TransactOpts, _SocializePoolOptInCoolingPeriod *big.Int) (*types.Transaction, error) {
	return _StaderConfig.contract.Transact(opts, "updateSocializingPoolOptInCoolingPeriod", _SocializePoolOptInCoolingPeriod)
}

// UpdateSocializingPoolOptInCoolingPeriod is a paid mutator transaction binding the contract method 0x8a4cfb58.
//
// Solidity: function updateSocializingPoolOptInCoolingPeriod(uint256 _SocializePoolOptInCoolingPeriod) returns()
func (_StaderConfig *StaderConfigSession) UpdateSocializingPoolOptInCoolingPeriod(_SocializePoolOptInCoolingPeriod *big.Int) (*types.Transaction, error) {
	return _StaderConfig.Contract.UpdateSocializingPoolOptInCoolingPeriod(&_StaderConfig.TransactOpts, _SocializePoolOptInCoolingPeriod)
}

// UpdateSocializingPoolOptInCoolingPeriod is a paid mutator transaction binding the contract method 0x8a4cfb58.
//
// Solidity: function updateSocializingPoolOptInCoolingPeriod(uint256 _SocializePoolOptInCoolingPeriod) returns()
func (_StaderConfig *StaderConfigTransactorSession) UpdateSocializingPoolOptInCoolingPeriod(_SocializePoolOptInCoolingPeriod *big.Int) (*types.Transaction, error) {
	return _StaderConfig.Contract.UpdateSocializingPoolOptInCoolingPeriod(&_StaderConfig.TransactOpts, _SocializePoolOptInCoolingPeriod)
}

// UpdateStaderInsuranceFund is a paid mutator transaction binding the contract method 0xaa2f56c7.
//
// Solidity: function updateStaderInsuranceFund(address _staderInsuranceFund) returns()
func (_StaderConfig *StaderConfigTransactor) UpdateStaderInsuranceFund(opts *bind.TransactOpts, _staderInsuranceFund common.Address) (*types.Transaction, error) {
	return _StaderConfig.contract.Transact(opts, "updateStaderInsuranceFund", _staderInsuranceFund)
}

// UpdateStaderInsuranceFund is a paid mutator transaction binding the contract method 0xaa2f56c7.
//
// Solidity: function updateStaderInsuranceFund(address _staderInsuranceFund) returns()
func (_StaderConfig *StaderConfigSession) UpdateStaderInsuranceFund(_staderInsuranceFund common.Address) (*types.Transaction, error) {
	return _StaderConfig.Contract.UpdateStaderInsuranceFund(&_StaderConfig.TransactOpts, _staderInsuranceFund)
}

// UpdateStaderInsuranceFund is a paid mutator transaction binding the contract method 0xaa2f56c7.
//
// Solidity: function updateStaderInsuranceFund(address _staderInsuranceFund) returns()
func (_StaderConfig *StaderConfigTransactorSession) UpdateStaderInsuranceFund(_staderInsuranceFund common.Address) (*types.Transaction, error) {
	return _StaderConfig.Contract.UpdateStaderInsuranceFund(&_StaderConfig.TransactOpts, _staderInsuranceFund)
}

// UpdateStaderOracle is a paid mutator transaction binding the contract method 0xca78360c.
//
// Solidity: function updateStaderOracle(address _staderOracle) returns()
func (_StaderConfig *StaderConfigTransactor) UpdateStaderOracle(opts *bind.TransactOpts, _staderOracle common.Address) (*types.Transaction, error) {
	return _StaderConfig.contract.Transact(opts, "updateStaderOracle", _staderOracle)
}

// UpdateStaderOracle is a paid mutator transaction binding the contract method 0xca78360c.
//
// Solidity: function updateStaderOracle(address _staderOracle) returns()
func (_StaderConfig *StaderConfigSession) UpdateStaderOracle(_staderOracle common.Address) (*types.Transaction, error) {
	return _StaderConfig.Contract.UpdateStaderOracle(&_StaderConfig.TransactOpts, _staderOracle)
}

// UpdateStaderOracle is a paid mutator transaction binding the contract method 0xca78360c.
//
// Solidity: function updateStaderOracle(address _staderOracle) returns()
func (_StaderConfig *StaderConfigTransactorSession) UpdateStaderOracle(_staderOracle common.Address) (*types.Transaction, error) {
	return _StaderConfig.Contract.UpdateStaderOracle(&_StaderConfig.TransactOpts, _staderOracle)
}

// UpdateStaderToken is a paid mutator transaction binding the contract method 0x83148593.
//
// Solidity: function updateStaderToken(address _staderToken) returns()
func (_StaderConfig *StaderConfigTransactor) UpdateStaderToken(opts *bind.TransactOpts, _staderToken common.Address) (*types.Transaction, error) {
	return _StaderConfig.contract.Transact(opts, "updateStaderToken", _staderToken)
}

// UpdateStaderToken is a paid mutator transaction binding the contract method 0x83148593.
//
// Solidity: function updateStaderToken(address _staderToken) returns()
func (_StaderConfig *StaderConfigSession) UpdateStaderToken(_staderToken common.Address) (*types.Transaction, error) {
	return _StaderConfig.Contract.UpdateStaderToken(&_StaderConfig.TransactOpts, _staderToken)
}

// UpdateStaderToken is a paid mutator transaction binding the contract method 0x83148593.
//
// Solidity: function updateStaderToken(address _staderToken) returns()
func (_StaderConfig *StaderConfigTransactorSession) UpdateStaderToken(_staderToken common.Address) (*types.Transaction, error) {
	return _StaderConfig.Contract.UpdateStaderToken(&_StaderConfig.TransactOpts, _staderToken)
}

// UpdateStaderTreasury is a paid mutator transaction binding the contract method 0x7b4cd7ec.
//
// Solidity: function updateStaderTreasury(address _staderTreasury) returns()
func (_StaderConfig *StaderConfigTransactor) UpdateStaderTreasury(opts *bind.TransactOpts, _staderTreasury common.Address) (*types.Transaction, error) {
	return _StaderConfig.contract.Transact(opts, "updateStaderTreasury", _staderTreasury)
}

// UpdateStaderTreasury is a paid mutator transaction binding the contract method 0x7b4cd7ec.
//
// Solidity: function updateStaderTreasury(address _staderTreasury) returns()
func (_StaderConfig *StaderConfigSession) UpdateStaderTreasury(_staderTreasury common.Address) (*types.Transaction, error) {
	return _StaderConfig.Contract.UpdateStaderTreasury(&_StaderConfig.TransactOpts, _staderTreasury)
}

// UpdateStaderTreasury is a paid mutator transaction binding the contract method 0x7b4cd7ec.
//
// Solidity: function updateStaderTreasury(address _staderTreasury) returns()
func (_StaderConfig *StaderConfigTransactorSession) UpdateStaderTreasury(_staderTreasury common.Address) (*types.Transaction, error) {
	return _StaderConfig.Contract.UpdateStaderTreasury(&_StaderConfig.TransactOpts, _staderTreasury)
}

// UpdateStakePoolManager is a paid mutator transaction binding the contract method 0x368f9d17.
//
// Solidity: function updateStakePoolManager(address _stakePoolManager) returns()
func (_StaderConfig *StaderConfigTransactor) UpdateStakePoolManager(opts *bind.TransactOpts, _stakePoolManager common.Address) (*types.Transaction, error) {
	return _StaderConfig.contract.Transact(opts, "updateStakePoolManager", _stakePoolManager)
}

// UpdateStakePoolManager is a paid mutator transaction binding the contract method 0x368f9d17.
//
// Solidity: function updateStakePoolManager(address _stakePoolManager) returns()
func (_StaderConfig *StaderConfigSession) UpdateStakePoolManager(_stakePoolManager common.Address) (*types.Transaction, error) {
	return _StaderConfig.Contract.UpdateStakePoolManager(&_StaderConfig.TransactOpts, _stakePoolManager)
}

// UpdateStakePoolManager is a paid mutator transaction binding the contract method 0x368f9d17.
//
// Solidity: function updateStakePoolManager(address _stakePoolManager) returns()
func (_StaderConfig *StaderConfigTransactorSession) UpdateStakePoolManager(_stakePoolManager common.Address) (*types.Transaction, error) {
	return _StaderConfig.Contract.UpdateStakePoolManager(&_StaderConfig.TransactOpts, _stakePoolManager)
}

// UpdateUserWithdrawManager is a paid mutator transaction binding the contract method 0x088ee72d.
//
// Solidity: function updateUserWithdrawManager(address _userWithdrawManager) returns()
func (_StaderConfig *StaderConfigTransactor) UpdateUserWithdrawManager(opts *bind.TransactOpts, _userWithdrawManager common.Address) (*types.Transaction, error) {
	return _StaderConfig.contract.Transact(opts, "updateUserWithdrawManager", _userWithdrawManager)
}

// UpdateUserWithdrawManager is a paid mutator transaction binding the contract method 0x088ee72d.
//
// Solidity: function updateUserWithdrawManager(address _userWithdrawManager) returns()
func (_StaderConfig *StaderConfigSession) UpdateUserWithdrawManager(_userWithdrawManager common.Address) (*types.Transaction, error) {
	return _StaderConfig.Contract.UpdateUserWithdrawManager(&_StaderConfig.TransactOpts, _userWithdrawManager)
}

// UpdateUserWithdrawManager is a paid mutator transaction binding the contract method 0x088ee72d.
//
// Solidity: function updateUserWithdrawManager(address _userWithdrawManager) returns()
func (_StaderConfig *StaderConfigTransactorSession) UpdateUserWithdrawManager(_userWithdrawManager common.Address) (*types.Transaction, error) {
	return _StaderConfig.Contract.UpdateUserWithdrawManager(&_StaderConfig.TransactOpts, _userWithdrawManager)
}

// UpdateValidatorWithdrawalVaultImplementation is a paid mutator transaction binding the contract method 0x5b5961fc.
//
// Solidity: function updateValidatorWithdrawalVaultImplementation(address _validatorWithdrawalVaultImpl) returns()
func (_StaderConfig *StaderConfigTransactor) UpdateValidatorWithdrawalVaultImplementation(opts *bind.TransactOpts, _validatorWithdrawalVaultImpl common.Address) (*types.Transaction, error) {
	return _StaderConfig.contract.Transact(opts, "updateValidatorWithdrawalVaultImplementation", _validatorWithdrawalVaultImpl)
}

// UpdateValidatorWithdrawalVaultImplementation is a paid mutator transaction binding the contract method 0x5b5961fc.
//
// Solidity: function updateValidatorWithdrawalVaultImplementation(address _validatorWithdrawalVaultImpl) returns()
func (_StaderConfig *StaderConfigSession) UpdateValidatorWithdrawalVaultImplementation(_validatorWithdrawalVaultImpl common.Address) (*types.Transaction, error) {
	return _StaderConfig.Contract.UpdateValidatorWithdrawalVaultImplementation(&_StaderConfig.TransactOpts, _validatorWithdrawalVaultImpl)
}

// UpdateValidatorWithdrawalVaultImplementation is a paid mutator transaction binding the contract method 0x5b5961fc.
//
// Solidity: function updateValidatorWithdrawalVaultImplementation(address _validatorWithdrawalVaultImpl) returns()
func (_StaderConfig *StaderConfigTransactorSession) UpdateValidatorWithdrawalVaultImplementation(_validatorWithdrawalVaultImpl common.Address) (*types.Transaction, error) {
	return _StaderConfig.Contract.UpdateValidatorWithdrawalVaultImplementation(&_StaderConfig.TransactOpts, _validatorWithdrawalVaultImpl)
}

// UpdateVaultFactory is a paid mutator transaction binding the contract method 0x98c35927.
//
// Solidity: function updateVaultFactory(address _vaultFactory) returns()
func (_StaderConfig *StaderConfigTransactor) UpdateVaultFactory(opts *bind.TransactOpts, _vaultFactory common.Address) (*types.Transaction, error) {
	return _StaderConfig.contract.Transact(opts, "updateVaultFactory", _vaultFactory)
}

// UpdateVaultFactory is a paid mutator transaction binding the contract method 0x98c35927.
//
// Solidity: function updateVaultFactory(address _vaultFactory) returns()
func (_StaderConfig *StaderConfigSession) UpdateVaultFactory(_vaultFactory common.Address) (*types.Transaction, error) {
	return _StaderConfig.Contract.UpdateVaultFactory(&_StaderConfig.TransactOpts, _vaultFactory)
}

// UpdateVaultFactory is a paid mutator transaction binding the contract method 0x98c35927.
//
// Solidity: function updateVaultFactory(address _vaultFactory) returns()
func (_StaderConfig *StaderConfigTransactorSession) UpdateVaultFactory(_vaultFactory common.Address) (*types.Transaction, error) {
	return _StaderConfig.Contract.UpdateVaultFactory(&_StaderConfig.TransactOpts, _vaultFactory)
}

// UpdateWithdrawnKeysBatchSize is a paid mutator transaction binding the contract method 0xbbb99bb5.
//
// Solidity: function updateWithdrawnKeysBatchSize(uint256 _withdrawnKeysBatchSize) returns()
func (_StaderConfig *StaderConfigTransactor) UpdateWithdrawnKeysBatchSize(opts *bind.TransactOpts, _withdrawnKeysBatchSize *big.Int) (*types.Transaction, error) {
	return _StaderConfig.contract.Transact(opts, "updateWithdrawnKeysBatchSize", _withdrawnKeysBatchSize)
}

// UpdateWithdrawnKeysBatchSize is a paid mutator transaction binding the contract method 0xbbb99bb5.
//
// Solidity: function updateWithdrawnKeysBatchSize(uint256 _withdrawnKeysBatchSize) returns()
func (_StaderConfig *StaderConfigSession) UpdateWithdrawnKeysBatchSize(_withdrawnKeysBatchSize *big.Int) (*types.Transaction, error) {
	return _StaderConfig.Contract.UpdateWithdrawnKeysBatchSize(&_StaderConfig.TransactOpts, _withdrawnKeysBatchSize)
}

// UpdateWithdrawnKeysBatchSize is a paid mutator transaction binding the contract method 0xbbb99bb5.
//
// Solidity: function updateWithdrawnKeysBatchSize(uint256 _withdrawnKeysBatchSize) returns()
func (_StaderConfig *StaderConfigTransactorSession) UpdateWithdrawnKeysBatchSize(_withdrawnKeysBatchSize *big.Int) (*types.Transaction, error) {
	return _StaderConfig.Contract.UpdateWithdrawnKeysBatchSize(&_StaderConfig.TransactOpts, _withdrawnKeysBatchSize)
}

// StaderConfigInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the StaderConfig contract.
type StaderConfigInitializedIterator struct {
	Event *StaderConfigInitialized // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *StaderConfigInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StaderConfigInitialized)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(StaderConfigInitialized)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *StaderConfigInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StaderConfigInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StaderConfigInitialized represents a Initialized event raised by the StaderConfig contract.
type StaderConfigInitialized struct {
	Version uint8
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_StaderConfig *StaderConfigFilterer) FilterInitialized(opts *bind.FilterOpts) (*StaderConfigInitializedIterator, error) {

	logs, sub, err := _StaderConfig.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &StaderConfigInitializedIterator{contract: _StaderConfig.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_StaderConfig *StaderConfigFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *StaderConfigInitialized) (event.Subscription, error) {

	logs, sub, err := _StaderConfig.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StaderConfigInitialized)
				if err := _StaderConfig.contract.UnpackLog(event, "Initialized", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseInitialized is a log parse operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_StaderConfig *StaderConfigFilterer) ParseInitialized(log types.Log) (*StaderConfigInitialized, error) {
	event := new(StaderConfigInitialized)
	if err := _StaderConfig.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StaderConfigRoleAdminChangedIterator is returned from FilterRoleAdminChanged and is used to iterate over the raw logs and unpacked data for RoleAdminChanged events raised by the StaderConfig contract.
type StaderConfigRoleAdminChangedIterator struct {
	Event *StaderConfigRoleAdminChanged // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *StaderConfigRoleAdminChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StaderConfigRoleAdminChanged)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(StaderConfigRoleAdminChanged)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *StaderConfigRoleAdminChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StaderConfigRoleAdminChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StaderConfigRoleAdminChanged represents a RoleAdminChanged event raised by the StaderConfig contract.
type StaderConfigRoleAdminChanged struct {
	Role              [32]byte
	PreviousAdminRole [32]byte
	NewAdminRole      [32]byte
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterRoleAdminChanged is a free log retrieval operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_StaderConfig *StaderConfigFilterer) FilterRoleAdminChanged(opts *bind.FilterOpts, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (*StaderConfigRoleAdminChangedIterator, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var previousAdminRoleRule []interface{}
	for _, previousAdminRoleItem := range previousAdminRole {
		previousAdminRoleRule = append(previousAdminRoleRule, previousAdminRoleItem)
	}
	var newAdminRoleRule []interface{}
	for _, newAdminRoleItem := range newAdminRole {
		newAdminRoleRule = append(newAdminRoleRule, newAdminRoleItem)
	}

	logs, sub, err := _StaderConfig.contract.FilterLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return &StaderConfigRoleAdminChangedIterator{contract: _StaderConfig.contract, event: "RoleAdminChanged", logs: logs, sub: sub}, nil
}

// WatchRoleAdminChanged is a free log subscription operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_StaderConfig *StaderConfigFilterer) WatchRoleAdminChanged(opts *bind.WatchOpts, sink chan<- *StaderConfigRoleAdminChanged, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (event.Subscription, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var previousAdminRoleRule []interface{}
	for _, previousAdminRoleItem := range previousAdminRole {
		previousAdminRoleRule = append(previousAdminRoleRule, previousAdminRoleItem)
	}
	var newAdminRoleRule []interface{}
	for _, newAdminRoleItem := range newAdminRole {
		newAdminRoleRule = append(newAdminRoleRule, newAdminRoleItem)
	}

	logs, sub, err := _StaderConfig.contract.WatchLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StaderConfigRoleAdminChanged)
				if err := _StaderConfig.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseRoleAdminChanged is a log parse operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_StaderConfig *StaderConfigFilterer) ParseRoleAdminChanged(log types.Log) (*StaderConfigRoleAdminChanged, error) {
	event := new(StaderConfigRoleAdminChanged)
	if err := _StaderConfig.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StaderConfigRoleGrantedIterator is returned from FilterRoleGranted and is used to iterate over the raw logs and unpacked data for RoleGranted events raised by the StaderConfig contract.
type StaderConfigRoleGrantedIterator struct {
	Event *StaderConfigRoleGranted // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *StaderConfigRoleGrantedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StaderConfigRoleGranted)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(StaderConfigRoleGranted)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *StaderConfigRoleGrantedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StaderConfigRoleGrantedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StaderConfigRoleGranted represents a RoleGranted event raised by the StaderConfig contract.
type StaderConfigRoleGranted struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleGranted is a free log retrieval operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_StaderConfig *StaderConfigFilterer) FilterRoleGranted(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*StaderConfigRoleGrantedIterator, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _StaderConfig.contract.FilterLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &StaderConfigRoleGrantedIterator{contract: _StaderConfig.contract, event: "RoleGranted", logs: logs, sub: sub}, nil
}

// WatchRoleGranted is a free log subscription operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_StaderConfig *StaderConfigFilterer) WatchRoleGranted(opts *bind.WatchOpts, sink chan<- *StaderConfigRoleGranted, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _StaderConfig.contract.WatchLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StaderConfigRoleGranted)
				if err := _StaderConfig.contract.UnpackLog(event, "RoleGranted", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseRoleGranted is a log parse operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_StaderConfig *StaderConfigFilterer) ParseRoleGranted(log types.Log) (*StaderConfigRoleGranted, error) {
	event := new(StaderConfigRoleGranted)
	if err := _StaderConfig.contract.UnpackLog(event, "RoleGranted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StaderConfigRoleRevokedIterator is returned from FilterRoleRevoked and is used to iterate over the raw logs and unpacked data for RoleRevoked events raised by the StaderConfig contract.
type StaderConfigRoleRevokedIterator struct {
	Event *StaderConfigRoleRevoked // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *StaderConfigRoleRevokedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StaderConfigRoleRevoked)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(StaderConfigRoleRevoked)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *StaderConfigRoleRevokedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StaderConfigRoleRevokedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StaderConfigRoleRevoked represents a RoleRevoked event raised by the StaderConfig contract.
type StaderConfigRoleRevoked struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleRevoked is a free log retrieval operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_StaderConfig *StaderConfigFilterer) FilterRoleRevoked(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*StaderConfigRoleRevokedIterator, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _StaderConfig.contract.FilterLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &StaderConfigRoleRevokedIterator{contract: _StaderConfig.contract, event: "RoleRevoked", logs: logs, sub: sub}, nil
}

// WatchRoleRevoked is a free log subscription operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_StaderConfig *StaderConfigFilterer) WatchRoleRevoked(opts *bind.WatchOpts, sink chan<- *StaderConfigRoleRevoked, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _StaderConfig.contract.WatchLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StaderConfigRoleRevoked)
				if err := _StaderConfig.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseRoleRevoked is a log parse operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_StaderConfig *StaderConfigFilterer) ParseRoleRevoked(log types.Log) (*StaderConfigRoleRevoked, error) {
	event := new(StaderConfigRoleRevoked)
	if err := _StaderConfig.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StaderConfigSetAccountIterator is returned from FilterSetAccount and is used to iterate over the raw logs and unpacked data for SetAccount events raised by the StaderConfig contract.
type StaderConfigSetAccountIterator struct {
	Event *StaderConfigSetAccount // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *StaderConfigSetAccountIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StaderConfigSetAccount)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(StaderConfigSetAccount)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *StaderConfigSetAccountIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StaderConfigSetAccountIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StaderConfigSetAccount represents a SetAccount event raised by the StaderConfig contract.
type StaderConfigSetAccount struct {
	Key        [32]byte
	NewAddress common.Address
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterSetAccount is a free log retrieval operation binding the contract event 0xcbdd341876786c7241ad12a5ce5ea46739a4ce7b1587d0c216dfa655a98e50a6.
//
// Solidity: event SetAccount(bytes32 key, address newAddress)
func (_StaderConfig *StaderConfigFilterer) FilterSetAccount(opts *bind.FilterOpts) (*StaderConfigSetAccountIterator, error) {

	logs, sub, err := _StaderConfig.contract.FilterLogs(opts, "SetAccount")
	if err != nil {
		return nil, err
	}
	return &StaderConfigSetAccountIterator{contract: _StaderConfig.contract, event: "SetAccount", logs: logs, sub: sub}, nil
}

// WatchSetAccount is a free log subscription operation binding the contract event 0xcbdd341876786c7241ad12a5ce5ea46739a4ce7b1587d0c216dfa655a98e50a6.
//
// Solidity: event SetAccount(bytes32 key, address newAddress)
func (_StaderConfig *StaderConfigFilterer) WatchSetAccount(opts *bind.WatchOpts, sink chan<- *StaderConfigSetAccount) (event.Subscription, error) {

	logs, sub, err := _StaderConfig.contract.WatchLogs(opts, "SetAccount")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StaderConfigSetAccount)
				if err := _StaderConfig.contract.UnpackLog(event, "SetAccount", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseSetAccount is a log parse operation binding the contract event 0xcbdd341876786c7241ad12a5ce5ea46739a4ce7b1587d0c216dfa655a98e50a6.
//
// Solidity: event SetAccount(bytes32 key, address newAddress)
func (_StaderConfig *StaderConfigFilterer) ParseSetAccount(log types.Log) (*StaderConfigSetAccount, error) {
	event := new(StaderConfigSetAccount)
	if err := _StaderConfig.contract.UnpackLog(event, "SetAccount", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StaderConfigSetConstantIterator is returned from FilterSetConstant and is used to iterate over the raw logs and unpacked data for SetConstant events raised by the StaderConfig contract.
type StaderConfigSetConstantIterator struct {
	Event *StaderConfigSetConstant // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *StaderConfigSetConstantIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StaderConfigSetConstant)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(StaderConfigSetConstant)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *StaderConfigSetConstantIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StaderConfigSetConstantIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StaderConfigSetConstant represents a SetConstant event raised by the StaderConfig contract.
type StaderConfigSetConstant struct {
	Key    [32]byte
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterSetConstant is a free log retrieval operation binding the contract event 0x9094260c4234c0cb4c44e4a035abb5816b84e5505f9dc571c3ff397c46581630.
//
// Solidity: event SetConstant(bytes32 key, uint256 amount)
func (_StaderConfig *StaderConfigFilterer) FilterSetConstant(opts *bind.FilterOpts) (*StaderConfigSetConstantIterator, error) {

	logs, sub, err := _StaderConfig.contract.FilterLogs(opts, "SetConstant")
	if err != nil {
		return nil, err
	}
	return &StaderConfigSetConstantIterator{contract: _StaderConfig.contract, event: "SetConstant", logs: logs, sub: sub}, nil
}

// WatchSetConstant is a free log subscription operation binding the contract event 0x9094260c4234c0cb4c44e4a035abb5816b84e5505f9dc571c3ff397c46581630.
//
// Solidity: event SetConstant(bytes32 key, uint256 amount)
func (_StaderConfig *StaderConfigFilterer) WatchSetConstant(opts *bind.WatchOpts, sink chan<- *StaderConfigSetConstant) (event.Subscription, error) {

	logs, sub, err := _StaderConfig.contract.WatchLogs(opts, "SetConstant")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StaderConfigSetConstant)
				if err := _StaderConfig.contract.UnpackLog(event, "SetConstant", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseSetConstant is a log parse operation binding the contract event 0x9094260c4234c0cb4c44e4a035abb5816b84e5505f9dc571c3ff397c46581630.
//
// Solidity: event SetConstant(bytes32 key, uint256 amount)
func (_StaderConfig *StaderConfigFilterer) ParseSetConstant(log types.Log) (*StaderConfigSetConstant, error) {
	event := new(StaderConfigSetConstant)
	if err := _StaderConfig.contract.UnpackLog(event, "SetConstant", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StaderConfigSetContractIterator is returned from FilterSetContract and is used to iterate over the raw logs and unpacked data for SetContract events raised by the StaderConfig contract.
type StaderConfigSetContractIterator struct {
	Event *StaderConfigSetContract // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *StaderConfigSetContractIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StaderConfigSetContract)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(StaderConfigSetContract)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *StaderConfigSetContractIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StaderConfigSetContractIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StaderConfigSetContract represents a SetContract event raised by the StaderConfig contract.
type StaderConfigSetContract struct {
	Key        [32]byte
	NewAddress common.Address
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterSetContract is a free log retrieval operation binding the contract event 0x5de40a806536a2029221dac2c8887ac9f11952fcc1ed3d7cfb4476dd5259b740.
//
// Solidity: event SetContract(bytes32 key, address newAddress)
func (_StaderConfig *StaderConfigFilterer) FilterSetContract(opts *bind.FilterOpts) (*StaderConfigSetContractIterator, error) {

	logs, sub, err := _StaderConfig.contract.FilterLogs(opts, "SetContract")
	if err != nil {
		return nil, err
	}
	return &StaderConfigSetContractIterator{contract: _StaderConfig.contract, event: "SetContract", logs: logs, sub: sub}, nil
}

// WatchSetContract is a free log subscription operation binding the contract event 0x5de40a806536a2029221dac2c8887ac9f11952fcc1ed3d7cfb4476dd5259b740.
//
// Solidity: event SetContract(bytes32 key, address newAddress)
func (_StaderConfig *StaderConfigFilterer) WatchSetContract(opts *bind.WatchOpts, sink chan<- *StaderConfigSetContract) (event.Subscription, error) {

	logs, sub, err := _StaderConfig.contract.WatchLogs(opts, "SetContract")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StaderConfigSetContract)
				if err := _StaderConfig.contract.UnpackLog(event, "SetContract", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseSetContract is a log parse operation binding the contract event 0x5de40a806536a2029221dac2c8887ac9f11952fcc1ed3d7cfb4476dd5259b740.
//
// Solidity: event SetContract(bytes32 key, address newAddress)
func (_StaderConfig *StaderConfigFilterer) ParseSetContract(log types.Log) (*StaderConfigSetContract, error) {
	event := new(StaderConfigSetContract)
	if err := _StaderConfig.contract.UnpackLog(event, "SetContract", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StaderConfigSetTokenIterator is returned from FilterSetToken and is used to iterate over the raw logs and unpacked data for SetToken events raised by the StaderConfig contract.
type StaderConfigSetTokenIterator struct {
	Event *StaderConfigSetToken // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *StaderConfigSetTokenIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StaderConfigSetToken)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(StaderConfigSetToken)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *StaderConfigSetTokenIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StaderConfigSetTokenIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StaderConfigSetToken represents a SetToken event raised by the StaderConfig contract.
type StaderConfigSetToken struct {
	Key        [32]byte
	NewAddress common.Address
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterSetToken is a free log retrieval operation binding the contract event 0x19aab10c6a9f5d648eaa15e2d515f8dfda570ee221e7c8cb9dc07694e68005bc.
//
// Solidity: event SetToken(bytes32 key, address newAddress)
func (_StaderConfig *StaderConfigFilterer) FilterSetToken(opts *bind.FilterOpts) (*StaderConfigSetTokenIterator, error) {

	logs, sub, err := _StaderConfig.contract.FilterLogs(opts, "SetToken")
	if err != nil {
		return nil, err
	}
	return &StaderConfigSetTokenIterator{contract: _StaderConfig.contract, event: "SetToken", logs: logs, sub: sub}, nil
}

// WatchSetToken is a free log subscription operation binding the contract event 0x19aab10c6a9f5d648eaa15e2d515f8dfda570ee221e7c8cb9dc07694e68005bc.
//
// Solidity: event SetToken(bytes32 key, address newAddress)
func (_StaderConfig *StaderConfigFilterer) WatchSetToken(opts *bind.WatchOpts, sink chan<- *StaderConfigSetToken) (event.Subscription, error) {

	logs, sub, err := _StaderConfig.contract.WatchLogs(opts, "SetToken")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StaderConfigSetToken)
				if err := _StaderConfig.contract.UnpackLog(event, "SetToken", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseSetToken is a log parse operation binding the contract event 0x19aab10c6a9f5d648eaa15e2d515f8dfda570ee221e7c8cb9dc07694e68005bc.
//
// Solidity: event SetToken(bytes32 key, address newAddress)
func (_StaderConfig *StaderConfigFilterer) ParseSetToken(log types.Log) (*StaderConfigSetToken, error) {
	event := new(StaderConfigSetToken)
	if err := _StaderConfig.contract.UnpackLog(event, "SetToken", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StaderConfigSetVariableIterator is returned from FilterSetVariable and is used to iterate over the raw logs and unpacked data for SetVariable events raised by the StaderConfig contract.
type StaderConfigSetVariableIterator struct {
	Event *StaderConfigSetVariable // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *StaderConfigSetVariableIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StaderConfigSetVariable)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(StaderConfigSetVariable)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *StaderConfigSetVariableIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StaderConfigSetVariableIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StaderConfigSetVariable represents a SetVariable event raised by the StaderConfig contract.
type StaderConfigSetVariable struct {
	Key    [32]byte
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterSetVariable is a free log retrieval operation binding the contract event 0x8091d0a33efc2ef6d3c254604932c813b52281547a6915df7dc1a5554f080c3e.
//
// Solidity: event SetVariable(bytes32 key, uint256 amount)
func (_StaderConfig *StaderConfigFilterer) FilterSetVariable(opts *bind.FilterOpts) (*StaderConfigSetVariableIterator, error) {

	logs, sub, err := _StaderConfig.contract.FilterLogs(opts, "SetVariable")
	if err != nil {
		return nil, err
	}
	return &StaderConfigSetVariableIterator{contract: _StaderConfig.contract, event: "SetVariable", logs: logs, sub: sub}, nil
}

// WatchSetVariable is a free log subscription operation binding the contract event 0x8091d0a33efc2ef6d3c254604932c813b52281547a6915df7dc1a5554f080c3e.
//
// Solidity: event SetVariable(bytes32 key, uint256 amount)
func (_StaderConfig *StaderConfigFilterer) WatchSetVariable(opts *bind.WatchOpts, sink chan<- *StaderConfigSetVariable) (event.Subscription, error) {

	logs, sub, err := _StaderConfig.contract.WatchLogs(opts, "SetVariable")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StaderConfigSetVariable)
				if err := _StaderConfig.contract.UnpackLog(event, "SetVariable", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseSetVariable is a log parse operation binding the contract event 0x8091d0a33efc2ef6d3c254604932c813b52281547a6915df7dc1a5554f080c3e.
//
// Solidity: event SetVariable(bytes32 key, uint256 amount)
func (_StaderConfig *StaderConfigFilterer) ParseSetVariable(log types.Log) (*StaderConfigSetVariable, error) {
	event := new(StaderConfigSetVariable)
	if err := _StaderConfig.contract.UnpackLog(event, "SetVariable", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}