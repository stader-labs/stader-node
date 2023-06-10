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
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"InvalidLimits\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidMaxDepositValue\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidMaxWithdrawValue\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidMinDepositValue\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidMinWithdrawValue\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ZeroAddress\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"version\",\"type\":\"uint8\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"previousAdminRole\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"newAdminRole\",\"type\":\"bytes32\"}],\"name\":\"RoleAdminChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleGranted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleRevoked\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"key\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newAddress\",\"type\":\"address\"}],\"name\":\"SetAccount\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"key\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"SetConstant\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"key\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newAddress\",\"type\":\"address\"}],\"name\":\"SetContract\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"key\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newAddress\",\"type\":\"address\"}],\"name\":\"SetToken\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"key\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"SetVariable\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"ADMIN\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"AUCTION_CONTRACT\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"DECIMALS\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"DEFAULT_ADMIN_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"ETHX_SUPPLY_POR_FEED\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"ETH_BALANCE_POR_FEED\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"ETH_DEPOSIT_CONTRACT\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"ETH_PER_NODE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"ETHx\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"FULL_DEPOSIT_SIZE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"MANAGER\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"MAX_DEPOSIT_AMOUNT\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"MAX_WITHDRAW_AMOUNT\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"MIN_BLOCK_DELAY_TO_FINALIZE_WITHDRAW_REQUEST\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"MIN_DEPOSIT_AMOUNT\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"MIN_WITHDRAW_AMOUNT\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"NODE_EL_REWARD_VAULT_IMPLEMENTATION\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"OPERATOR\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"OPERATOR_MAX_NAME_LENGTH\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"OPERATOR_REWARD_COLLECTOR\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"PENALTY_CONTRACT\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"PERMISSIONED_NODE_REGISTRY\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"PERMISSIONED_POOL\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"PERMISSIONED_SOCIALIZING_POOL\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"PERMISSIONLESS_NODE_REGISTRY\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"PERMISSIONLESS_POOL\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"PERMISSIONLESS_SOCIALIZING_POOL\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"POOL_SELECTOR\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"POOL_UTILS\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"PRE_DEPOSIT_SIZE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"REWARD_THRESHOLD\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"SD\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"SD_COLLATERAL\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"SOCIALIZING_POOL_CYCLE_DURATION\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"SOCIALIZING_POOL_OPT_IN_COOLING_PERIOD\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"STADER_INSURANCE_FUND\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"STADER_ORACLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"STADER_TREASURY\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"STAKE_POOL_MANAGER\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"TOTAL_FEE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"USER_WITHDRAW_MANAGER\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"VALIDATOR_WITHDRAWAL_VAULT_IMPLEMENTATION\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"VAULT_FACTORY\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"WITHDRAWN_KEYS_BATCH_SIZE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getAdmin\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getAuctionContract\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getDecimals\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getETHBalancePORFeedProxy\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getETHDepositContract\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getETHXSupplyPORFeedProxy\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getETHxToken\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getFullDepositSize\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getMaxDepositAmount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getMaxWithdrawAmount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getMinBlockDelayToFinalizeWithdrawRequest\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getMinDepositAmount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getMinWithdrawAmount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getNodeELRewardVaultImplementation\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getOperatorMaxNameLength\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getOperatorRewardsCollector\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getPenaltyContract\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getPermissionedNodeRegistry\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getPermissionedPool\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getPermissionedSocializingPool\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getPermissionlessNodeRegistry\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getPermissionlessPool\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getPermissionlessSocializingPool\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getPoolSelector\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getPoolUtils\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getPreDepositSize\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getRewardsThreshold\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"}],\"name\":\"getRoleAdmin\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getSDCollateral\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getSocializingPoolCycleDuration\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getSocializingPoolOptInCoolingPeriod\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getStaderInsuranceFund\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getStaderOracle\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getStaderToken\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getStaderTreasury\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getStakePoolManager\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getStakedEthPerNode\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getTotalFee\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getUserWithdrawManager\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getValidatorWithdrawalVaultImplementation\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getVaultFactory\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getWithdrawnKeyBatchSize\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"grantRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"hasRole\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_admin\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_ethDepositContract\",\"type\":\"address\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"onlyManagerRole\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"onlyOperatorRole\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_addr\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"_contractName\",\"type\":\"bytes32\"}],\"name\":\"onlyStaderContract\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"renounceRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"revokeRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_admin\",\"type\":\"address\"}],\"name\":\"updateAdmin\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_auctionContract\",\"type\":\"address\"}],\"name\":\"updateAuctionContract\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_ethBalanceProxy\",\"type\":\"address\"}],\"name\":\"updateETHBalancePORFeedProxy\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_ethXSupplyProxy\",\"type\":\"address\"}],\"name\":\"updateETHXSupplyPORFeedProxy\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_ethX\",\"type\":\"address\"}],\"name\":\"updateETHxToken\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_maxDepositAmount\",\"type\":\"uint256\"}],\"name\":\"updateMaxDepositAmount\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_maxWithdrawAmount\",\"type\":\"uint256\"}],\"name\":\"updateMaxWithdrawAmount\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_minBlockDelay\",\"type\":\"uint256\"}],\"name\":\"updateMinBlockDelayToFinalizeWithdrawRequest\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_minDepositAmount\",\"type\":\"uint256\"}],\"name\":\"updateMinDepositAmount\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_minWithdrawAmount\",\"type\":\"uint256\"}],\"name\":\"updateMinWithdrawAmount\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_nodeELRewardVaultImpl\",\"type\":\"address\"}],\"name\":\"updateNodeELRewardImplementation\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_operatorRewardsCollector\",\"type\":\"address\"}],\"name\":\"updateOperatorRewardsCollector\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_penaltyContract\",\"type\":\"address\"}],\"name\":\"updatePenaltyContract\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_permissionedNodeRegistry\",\"type\":\"address\"}],\"name\":\"updatePermissionedNodeRegistry\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_permissionedPool\",\"type\":\"address\"}],\"name\":\"updatePermissionedPool\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_permissionedSocializePool\",\"type\":\"address\"}],\"name\":\"updatePermissionedSocializingPool\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_permissionlessNodeRegistry\",\"type\":\"address\"}],\"name\":\"updatePermissionlessNodeRegistry\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_permissionlessPool\",\"type\":\"address\"}],\"name\":\"updatePermissionlessPool\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_permissionlessSocializePool\",\"type\":\"address\"}],\"name\":\"updatePermissionlessSocializingPool\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_poolSelector\",\"type\":\"address\"}],\"name\":\"updatePoolSelector\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_poolUtils\",\"type\":\"address\"}],\"name\":\"updatePoolUtils\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_rewardsThreshold\",\"type\":\"uint256\"}],\"name\":\"updateRewardsThreshold\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_sdCollateral\",\"type\":\"address\"}],\"name\":\"updateSDCollateral\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_socializingPoolCycleDuration\",\"type\":\"uint256\"}],\"name\":\"updateSocializingPoolCycleDuration\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_SocializePoolOptInCoolingPeriod\",\"type\":\"uint256\"}],\"name\":\"updateSocializingPoolOptInCoolingPeriod\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_staderInsuranceFund\",\"type\":\"address\"}],\"name\":\"updateStaderInsuranceFund\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_staderOracle\",\"type\":\"address\"}],\"name\":\"updateStaderOracle\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_staderToken\",\"type\":\"address\"}],\"name\":\"updateStaderToken\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_staderTreasury\",\"type\":\"address\"}],\"name\":\"updateStaderTreasury\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_stakePoolManager\",\"type\":\"address\"}],\"name\":\"updateStakePoolManager\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_userWithdrawManager\",\"type\":\"address\"}],\"name\":\"updateUserWithdrawManager\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_validatorWithdrawalVaultImpl\",\"type\":\"address\"}],\"name\":\"updateValidatorWithdrawalVaultImplementation\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_vaultFactory\",\"type\":\"address\"}],\"name\":\"updateVaultFactory\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_withdrawnKeysBatchSize\",\"type\":\"uint256\"}],\"name\":\"updateWithdrawnKeysBatchSize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// StaderConfigABI is the input ABI used to generate the binding from.
// Deprecated: Use StaderConfigMetaData.ABI instead.
var StaderConfigABI = StaderConfigMetaData.ABI

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

// Initialize is a paid mutator transaction binding the contract method 0x485cc955.
//
// Solidity: function initialize(address _admin, address _ethDepositContract) returns()
func (_StaderConfig *StaderConfigTransactor) Initialize(opts *bind.TransactOpts, _admin common.Address, _ethDepositContract common.Address) (*types.Transaction, error) {
	return _StaderConfig.contract.Transact(opts, "initialize", _admin, _ethDepositContract)
}

// Initialize is a paid mutator transaction binding the contract method 0x485cc955.
//
// Solidity: function initialize(address _admin, address _ethDepositContract) returns()
func (_StaderConfig *StaderConfigSession) Initialize(_admin common.Address, _ethDepositContract common.Address) (*types.Transaction, error) {
	return _StaderConfig.Contract.Initialize(&_StaderConfig.TransactOpts, _admin, _ethDepositContract)
}

// Initialize is a paid mutator transaction binding the contract method 0x485cc955.
//
// Solidity: function initialize(address _admin, address _ethDepositContract) returns()
func (_StaderConfig *StaderConfigTransactorSession) Initialize(_admin common.Address, _ethDepositContract common.Address) (*types.Transaction, error) {
	return _StaderConfig.Contract.Initialize(&_StaderConfig.TransactOpts, _admin, _ethDepositContract)
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
