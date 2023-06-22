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

// Validator is an auto generated low-level Go binding around an user-defined struct.
type Validator struct {
	Status               uint8
	Pubkey               []byte
	PreDepositSignature  []byte
	DepositSignature     []byte
	WithdrawVaultAddress common.Address
	OperatorId           *big.Int
	DepositBlock         *big.Int
	WithdrawnBlock       *big.Int
}

// PermissionlessNodeRegistryMetaData contains all meta data concerning the PermissionlessNodeRegistry contract.
var PermissionlessNodeRegistryMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_admin\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_staderConfig\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"CallerNotManager\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"CallerNotOperator\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"CallerNotStaderContract\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"CooldownNotComplete\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"DuplicatePoolIDOrPoolNotAdded\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InSufficientBalance\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidBondEthValue\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidKeyCount\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidStartAndEndIndex\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"MisMatchingInputKeysSize\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NoChangeInState\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotEnoughSDCollateral\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OperatorAlreadyOnBoardedInProtocol\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OperatorIsDeactivate\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OperatorNotOnBoarded\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"PageNumberIsZero\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"PubkeyAlreadyExist\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"TooManyVerifiedKeysReported\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"TooManyWithdrawnKeysReported\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"TransferFailed\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"UNEXPECTED_STATUS\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ZeroAddress\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"maxKeyLimitReached\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"nodeOperator\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"pubkey\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"validatorId\",\"type\":\"uint256\"}],\"name\":\"AddedValidatorKey\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"totalActiveValidatorCount\",\"type\":\"uint256\"}],\"name\":\"DecreasedTotalActiveValidatorCount\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"totalActiveValidatorCount\",\"type\":\"uint256\"}],\"name\":\"IncreasedTotalActiveValidatorCount\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"version\",\"type\":\"uint8\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"nodeOperator\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"nodeRewardAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"operatorId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"optInForSocializingPool\",\"type\":\"bool\"}],\"name\":\"OnboardedOperator\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Paused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"previousAdminRole\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"newAdminRole\",\"type\":\"bytes32\"}],\"name\":\"RoleAdminChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleGranted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleRevoked\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"TransferredCollateralToPool\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Unpaused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"batchKeyDepositLimit\",\"type\":\"uint256\"}],\"name\":\"UpdatedInputKeyCountLimit\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"maxNonTerminalKeyPerOperator\",\"type\":\"uint64\"}],\"name\":\"UpdatedMaxNonTerminalKeyPerOperator\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"nextQueuedValidatorIndex\",\"type\":\"uint256\"}],\"name\":\"UpdatedNextQueuedValidatorIndex\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"nodeOperator\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"operatorName\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"rewardAddress\",\"type\":\"address\"}],\"name\":\"UpdatedOperatorDetails\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"operatorId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"optedForSocializingPool\",\"type\":\"bool\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"block\",\"type\":\"uint256\"}],\"name\":\"UpdatedSocializingPoolState\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"staderConfig\",\"type\":\"address\"}],\"name\":\"UpdatedStaderConfig\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"validatorId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"depositBlock\",\"type\":\"uint256\"}],\"name\":\"UpdatedValidatorDepositBlock\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"verifiedKeysBatchSize\",\"type\":\"uint256\"}],\"name\":\"UpdatedVerifiedKeyBatchSize\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"withdrawnKeysBatchSize\",\"type\":\"uint256\"}],\"name\":\"UpdatedWithdrawnKeyBatchSize\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"pubkey\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"validatorId\",\"type\":\"uint256\"}],\"name\":\"ValidatorMarkedAsFrontRunned\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"pubkey\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"validatorId\",\"type\":\"uint256\"}],\"name\":\"ValidatorMarkedReadyToDeposit\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"pubkey\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"validatorId\",\"type\":\"uint256\"}],\"name\":\"ValidatorStatusMarkedAsInvalidSignature\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"pubkey\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"validatorId\",\"type\":\"uint256\"}],\"name\":\"ValidatorWithdrawn\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"COLLATERAL_ETH\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"DEFAULT_ADMIN_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"FRONT_RUN_PENALTY\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"POOL_ID\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes[]\",\"name\":\"_pubkey\",\"type\":\"bytes[]\"},{\"internalType\":\"bytes[]\",\"name\":\"_preDepositSignature\",\"type\":\"bytes[]\"},{\"internalType\":\"bytes[]\",\"name\":\"_depositSignature\",\"type\":\"bytes[]\"}],\"name\":\"addValidatorKeys\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bool\",\"name\":\"_optInForSocializingPool\",\"type\":\"bool\"}],\"name\":\"changeSocializingPoolState\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"feeRecipientAddress\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_pageNumber\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_pageSize\",\"type\":\"uint256\"}],\"name\":\"getAllActiveValidators\",\"outputs\":[{\"components\":[{\"internalType\":\"enumValidatorStatus\",\"name\":\"status\",\"type\":\"uint8\"},{\"internalType\":\"bytes\",\"name\":\"pubkey\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"preDepositSignature\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"depositSignature\",\"type\":\"bytes\"},{\"internalType\":\"address\",\"name\":\"withdrawVaultAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"operatorId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"depositBlock\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"withdrawnBlock\",\"type\":\"uint256\"}],\"internalType\":\"structValidator[]\",\"name\":\"\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_pageNumber\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_pageSize\",\"type\":\"uint256\"}],\"name\":\"getAllNodeELVaultAddress\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getCollateralETH\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_operatorId\",\"type\":\"uint256\"}],\"name\":\"getOperatorRewardAddress\",\"outputs\":[{\"internalType\":\"addresspayable\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_operatorId\",\"type\":\"uint256\"}],\"name\":\"getOperatorTotalKeys\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"_totalKeys\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_nodeOperator\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_startIndex\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_endIndex\",\"type\":\"uint256\"}],\"name\":\"getOperatorTotalNonTerminalKeys\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"}],\"name\":\"getRoleAdmin\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_operatorId\",\"type\":\"uint256\"}],\"name\":\"getSocializingPoolStateChangeBlock\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getTotalActiveValidatorCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getTotalQueuedValidatorCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_operator\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_pageNumber\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_pageSize\",\"type\":\"uint256\"}],\"name\":\"getValidatorsByOperator\",\"outputs\":[{\"components\":[{\"internalType\":\"enumValidatorStatus\",\"name\":\"status\",\"type\":\"uint8\"},{\"internalType\":\"bytes\",\"name\":\"pubkey\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"preDepositSignature\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"depositSignature\",\"type\":\"bytes\"},{\"internalType\":\"address\",\"name\":\"withdrawVaultAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"operatorId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"depositBlock\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"withdrawnBlock\",\"type\":\"uint256\"}],\"internalType\":\"structValidator[]\",\"name\":\"\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"grantRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"hasRole\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_count\",\"type\":\"uint256\"}],\"name\":\"increaseTotalActiveValidatorCount\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"inputKeyCountLimit\",\"outputs\":[{\"internalType\":\"uint16\",\"name\":\"\",\"type\":\"uint16\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_operAddr\",\"type\":\"address\"}],\"name\":\"isExistingOperator\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"_pubkey\",\"type\":\"bytes\"}],\"name\":\"isExistingPubkey\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes[]\",\"name\":\"_readyToDepositPubkey\",\"type\":\"bytes[]\"},{\"internalType\":\"bytes[]\",\"name\":\"_frontRunPubkey\",\"type\":\"bytes[]\"},{\"internalType\":\"bytes[]\",\"name\":\"_invalidSignaturePubkey\",\"type\":\"bytes[]\"}],\"name\":\"markValidatorReadyToDeposit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"maxNonTerminalKeyPerOperator\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"nextOperatorId\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"nextQueuedValidatorIndex\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"nextValidatorId\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"nodeELRewardVaultByOperatorId\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bool\",\"name\":\"_optInForSocializingPool\",\"type\":\"bool\"},{\"internalType\":\"string\",\"name\":\"_operatorName\",\"type\":\"string\"},{\"internalType\":\"addresspayable\",\"name\":\"_operatorRewardAddress\",\"type\":\"address\"}],\"name\":\"onboardNodeOperator\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"feeRecipientAddress\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"operatorIDByAddress\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"operatorStructById\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"active\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"optedForSocializingPool\",\"type\":\"bool\"},{\"internalType\":\"string\",\"name\":\"operatorName\",\"type\":\"string\"},{\"internalType\":\"addresspayable\",\"name\":\"operatorRewardAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"operatorAddress\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"paused\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"queuedValidators\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"renounceRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"revokeRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"socializingPoolStateChangeBlock\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"staderConfig\",\"outputs\":[{\"internalType\":\"contractIStaderConfig\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalActiveValidatorCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"transferCollateralToPool\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"unpause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_validatorId\",\"type\":\"uint256\"}],\"name\":\"updateDepositStatusAndBlock\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint16\",\"name\":\"_inputKeyCountLimit\",\"type\":\"uint16\"}],\"name\":\"updateInputKeyCountLimit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"_maxNonTerminalKeyPerOperator\",\"type\":\"uint64\"}],\"name\":\"updateMaxNonTerminalKeyPerOperator\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_nextQueuedValidatorIndex\",\"type\":\"uint256\"}],\"name\":\"updateNextQueuedValidatorIndex\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_operatorName\",\"type\":\"string\"},{\"internalType\":\"addresspayable\",\"name\":\"_rewardAddress\",\"type\":\"address\"}],\"name\":\"updateOperatorDetails\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_staderConfig\",\"type\":\"address\"}],\"name\":\"updateStaderConfig\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_verifiedKeysBatchSize\",\"type\":\"uint256\"}],\"name\":\"updateVerifiedKeysBatchSize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"name\":\"validatorIdByPubkey\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"validatorIdsByOperatorId\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"validatorQueueSize\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"validatorRegistry\",\"outputs\":[{\"internalType\":\"enumValidatorStatus\",\"name\":\"status\",\"type\":\"uint8\"},{\"internalType\":\"bytes\",\"name\":\"pubkey\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"preDepositSignature\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"depositSignature\",\"type\":\"bytes\"},{\"internalType\":\"address\",\"name\":\"withdrawVaultAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"operatorId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"depositBlock\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"withdrawnBlock\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"verifiedKeyBatchSize\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes[]\",\"name\":\"_pubkeys\",\"type\":\"bytes[]\"}],\"name\":\"withdrawnValidators\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x608060405234801562000010575f80fd5b50604051620052f3380380620052f38339810160408190526200003391620004c6565b5f54610100900460ff16158080156200005257505f54600160ff909116105b806200006d5750303b1580156200006d57505f5460ff166001145b620000d65760405162461bcd60e51b815260206004820152602e60248201527f496e697469616c697a61626c653a20636f6e747261637420697320616c72656160448201526d191e481a5b9a5d1a585b1a5e995960921b60648201526084015b60405180910390fd5b5f805460ff191660011790558015620000f8575f805461ff0019166101001790555b6200010383620001f1565b6200010e82620001f1565b620001186200021c565b6200012262000278565b6200012c620002dc565b60fb8054600160fd81905560fe55601e7fffff0000000000000000000000000000000000000000ffffffffffffffff00009091166a01000000000000000000006001600160a01b0386160261ffff1916171762010000600160501b03191662320000179055603260fc55620001a25f8462000340565b8015620001e8575f805461ff0019169055604051600181527f7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb38474024989060200160405180910390a15b505050620004fc565b6001600160a01b038116620002195760405163d92e233d60e01b815260040160405180910390fd5b50565b5f54610100900460ff16620002765760405162461bcd60e51b815260206004820152602b60248201525f80516020620052d383398151915260448201526a6e697469616c697a696e6760a81b6064820152608401620000cd565b565b5f54610100900460ff16620002d25760405162461bcd60e51b815260206004820152602b60248201525f80516020620052d383398151915260448201526a6e697469616c697a696e6760a81b6064820152608401620000cd565b62000276620003e3565b5f54610100900460ff16620003365760405162461bcd60e51b815260206004820152602b60248201525f80516020620052d383398151915260448201526a6e697469616c697a696e6760a81b6064820152608401620000cd565b6200027662000449565b5f8281526065602090815260408083206001600160a01b038516845290915290205460ff16620003df575f8281526065602090815260408083206001600160a01b03851684529091529020805460ff191660011790556200039e3390565b6001600160a01b0316816001600160a01b0316837f2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d60405160405180910390a45b5050565b5f54610100900460ff166200043d5760405162461bcd60e51b815260206004820152602b60248201525f80516020620052d383398151915260448201526a6e697469616c697a696e6760a81b6064820152608401620000cd565b6097805460ff19169055565b5f54610100900460ff16620004a35760405162461bcd60e51b815260206004820152602b60248201525f80516020620052d383398151915260448201526a6e697469616c697a696e6760a81b6064820152608401620000cd565b600160c955565b80516001600160a01b0381168114620004c1575f80fd5b919050565b5f8060408385031215620004d8575f80fd5b620004e383620004aa565b9150620004f360208401620004aa565b90509250929050565b614dc9806200050a5f395ff3fe60806040526004361061035b575f3560e01c806383ea2358116101bd578063bb7306bf116100f2578063deacde2b11610092578063ebb5c1741161006d578063ebb5c17414610a64578063f7c0918914610a90578063f90b083814610aa5578063f9c4dda414610ac4575f80fd5b8063deacde2b146109fe578063e0bf8b5314610a11578063e0d7d0e914610a3e575f80fd5b8063c8a00e7a116100cd578063c8a00e7a14610964578063cac8b30614610994578063d547741f146109c0578063d5e1e5ce146109df575f80fd5b8063bb7306bf146108f1578063bc4a3ad51461090c578063c34ade5c14610938575f80fd5b8063998888981161015d578063ab3e71eb11610138578063ab3e71eb14610884578063af533aa814610899578063b01db078146108b8578063b8d2f06c146108d2575f80fd5b806399888898146108335780639ee804cb14610852578063a217fddf14610871575f80fd5b806384b0fa4c1161019857806384b0fa4c146107aa5780638a25bcec146107c057806391d14854146107df5780639344b242146107fe575f80fd5b806383ea23581461073257806384522a6d1461076a5780638456cb5914610796575f80fd5b806349911bfb116102935780635c2c30a511610233578063683547b81161020e578063683547b8146106c757806374338e6d146106f357806377c359e1146107095780637bd977d91461071e575f80fd5b80635c2c30a5146106595780635c975abb1461069157806360c3cf3f146106a8575f80fd5b806358a994ea1161026e57806358a994ea146105c957806359c3c9b7146105e85780635a1239c1146106075780635ae7f25d1461063a575f80fd5b806349911bfb1461055c5780634f59ed801461057157806350d5d7ab1461058c575f80fd5b80632d1dbd74116102fe57806336514d9f116102d957806336514d9f146104e457806336568abe146105035780633f4ba83a14610522578063490ffa3514610536575f80fd5b80632d1dbd74146104845780632d32924f146104995780632f2ff15d146104c5575f80fd5b8063186d954111610339578063186d9541146103eb578063248a9ca31461040a5780632517cfbf14610446578063264f27f314610465575f80fd5b806301ffc9a71461035f578063044d2fe81461039357806313797bff146103ca575b5f80fd5b34801561036a575f80fd5b5061037e610379366004614223565b610afb565b60405190151581526020015b60405180910390f35b34801561039e575f80fd5b506103b26103ad3660046142af565b610b31565b6040516001600160a01b03909116815260200161038a565b3480156103d5575f80fd5b506103e96103e4366004614352565b610f50565b005b3480156103f6575f80fd5b506103e96104053660046143e4565b611397565b348015610415575f80fd5b506104386104243660046143e4565b5f9081526065602052604090206001015490565b60405190815260200161038a565b348015610451575f80fd5b506103e96104603660046143fb565b611447565b348015610470575f80fd5b506103e961047f36600461441c565b6114a9565b34801561048f575f80fd5b5061043860fd5481565b3480156104a4575f80fd5b506104b86104b336600461445a565b6116f4565b60405161038a919061447a565b3480156104d0575f80fd5b506103e96104df3660046144c6565b611828565b3480156104ef575f80fd5b5061037e6104fe3660046144f4565b61184c565b34801561050e575f80fd5b506103e961051d3660046144c6565b611879565b34801561052d575f80fd5b506103e96118fc565b348015610541575f80fd5b5060fb546103b290600160501b90046001600160a01b031681565b348015610567575f80fd5b5061043860ff5481565b34801561057c575f80fd5b50610438673782dace9d90000081565b348015610597575f80fd5b5060fb546105b1906201000090046001600160401b031681565b6040516001600160401b03909116815260200161038a565b3480156105d4575f80fd5b506103e96105e3366004614526565b611924565b3480156105f3575f80fd5b506103e96106023660046143e4565b611aa2565b348015610612575f80fd5b506106266106213660046143e4565b611b42565b60405161038a9897969594939291906145f9565b348015610645575f80fd5b506103e96106543660046143e4565b611d24565b348015610664575f80fd5b50610438610673366004614686565b80516020818301810180516101038252928201919093012091525481565b34801561069c575f80fd5b5060975460ff1661037e565b3480156106b3575f80fd5b506103e96106c2366004614730565b611e8b565b3480156106d2575f80fd5b506106e66106e1366004614756565b611f05565b60405161038a9190614788565b3480156106fe575f80fd5b506104386101005481565b348015610714575f80fd5b5061010154610438565b348015610729575f80fd5b506104386122bc565b34801561073d575f80fd5b506103b261074c3660046143e4565b5f90815261010560205260409020600201546001600160a01b031690565b348015610775575f80fd5b506104386107843660046143e4565b6101086020525f908152604090205481565b3480156107a1575f80fd5b506103e96122d3565b3480156107b5575f80fd5b506104386101015481565b3480156107cb575f80fd5b506105b16107da366004614756565b6122f9565b3480156107ea575f80fd5b5061037e6107f93660046144c6565b6123c4565b348015610809575f80fd5b506103b26108183660046143e4565b6101096020525f90815260409020546001600160a01b031681565b34801561083e575f80fd5b506106e661084d36600461445a565b6123ee565b34801561085d575f80fd5b506103e961086c366004614872565b612739565b34801561087c575f80fd5b506104385f81565b34801561088f575f80fd5b5061043860fc5481565b3480156108a4575f80fd5b506103e96108b33660046143e4565b6127c2565b3480156108c3575f80fd5b50673782dace9d900000610438565b3480156108dd575f80fd5b506103e96108ec3660046143e4565b612815565b3480156108fc575f80fd5b506104386729a2241af62c000081565b348015610917575f80fd5b506104386109263660046143e4565b6101046020525f908152604090205481565b348015610943575f80fd5b506104386109523660046143e4565b5f908152610107602052604090205490565b34801561096f575f80fd5b5061098361097e3660046143e4565b6128a0565b60405161038a95949392919061488d565b34801561099f575f80fd5b506104386109ae366004614872565b6101066020525f908152604090205481565b3480156109cb575f80fd5b506103e96109da3660046144c6565b612969565b3480156109ea575f80fd5b506104386109f936600461445a565b61298d565b6103e9610a0c366004614352565b6129b9565b348015610a1c575f80fd5b5060fb54610a2b9061ffff1681565b60405161ffff909116815260200161038a565b348015610a49575f80fd5b50610a52600181565b60405160ff909116815260200161038a565b348015610a6f575f80fd5b50610438610a7e3660046143e4565b5f908152610108602052604090205490565b348015610a9b575f80fd5b5061043860fe5481565b348015610ab0575f80fd5b506103b2610abf3660046148d1565b612f44565b348015610acf575f80fd5b5061037e610ade366004614872565b6001600160a01b03165f9081526101066020526040902054151590565b5f6001600160e01b03198216637965db0b60e01b1480610b2b57506301ffc9a760e01b6001600160e01b03198316145b92915050565b5f610b3a6131ad565b5f60fb600a9054906101000a90046001600160a01b03166001600160a01b0316636ccb9d706040518163ffffffff1660e01b8152600401602060405180830381865afa158015610b8c573d5f803e3d5ffd5b505050506040513d601f19601f82011682018060405250810190610bb091906148ec565b905060fb600a9054906101000a90046001600160a01b03166001600160a01b0316639ca76b736040518163ffffffff1660e01b8152600401602060405180830381865afa158015610c03573d5f803e3d5ffd5b505050506040513d601f19601f82011682018060405250810190610c2791906148ec565b604051636fc4c27f60e11b8152600160048201526001600160a01b039182169183169063df8984fe90602401602060405180830381865afa158015610c6e573d5f803e3d5ffd5b505050506040513d601f19601f82011682018060405250810190610c9291906148ec565b6001600160a01b031614610cb9576040516303b8ffef60e41b815260040160405180910390fd5b604051639f7053f560e01b81526001600160a01b03821690639f7053f590610ce7908890889060040161492f565b5f604051808303815f87803b158015610cfe575f80fd5b505af1158015610d10573d5f803e3d5ffd5b50505050610d1d836131f3565b604051633e71376960e21b81523360048201526001600160a01b0382169063f9c4dda490602401602060405180830381865afa158015610d5f573d5f803e3d5ffd5b505050506040513d601f19601f82011682018060405250810190610d83919061494a565b15610da15760405163707999fb60e11b815260040160405180910390fd5b5f60fb600a9054906101000a90046001600160a01b03166001600160a01b03166318bcb2846040518163ffffffff1660e01b8152600401602060405180830381865afa158015610df3573d5f803e3d5ffd5b505050506040513d601f19601f82011682018060405250810190610e1791906148ec565b60fd54604051636a0b688160e01b81526001600482015260248101919091526001600160a01b039190911690636a0b6881906044016020604051808303815f875af1158015610e68573d5f803e3d5ffd5b505050506040513d601f19601f82011682018060405250810190610e8c91906148ec565b60fd545f9081526101096020526040902080546001600160a01b0319166001600160a01b038316179055905086610ec35780610f38565b60fb600a9054906101000a90046001600160a01b03166001600160a01b0316630a3fbd9a6040518163ffffffff1660e01b8152600401602060405180830381865afa158015610f14573d5f803e3d5ffd5b505050506040513d601f19601f82011682018060405250810190610f3891906148ec565b9250610f468787878761321a565b5050949350505050565b610f586133a8565b610f606131ad565b60fb5460408051633871d0f160e01b81529051610fde923392600160501b9091046001600160a01b0316918291633871d0f19160048083019260209291908290030181865afa158015610fb5573d5f803e3d5ffd5b505050506040513d601f19601f82011682018060405250810190610fd99190614965565b613401565b60fc5485908490839081610ff28486614990565b610ffc9190614990565b111561101b5760405163525e3de760e01b815260040160405180910390fd5b5f5b838110156110e4575f6101038b8b8481811061103b5761103b6149a3565b905060200281019061104d91906149b7565b60405161105b9291906149f9565b90815260200160405180910390205490506110758161348d565b61107e816134d9565b7f21d79a0b22a7d5a18b9535162fe2f0580e24c042b0541a05afc298a77ddf56938b8b848181106110b1576110b16149a3565b90506020028101906110c391906149b7565b836040516110d393929190614a08565b60405180910390a15060010161101d565b5081156111c15760fb600a9054906101000a90046001600160a01b03166001600160a01b031663b5cfee6c6040518163ffffffff1660e01b8152600401602060405180830381865afa15801561113c573d5f803e3d5ffd5b505050506040513d601f19601f8201168201806040525081019061116091906148ec565b6001600160a01b0316638d0d8cb66111806729a2241af62c000085614a2b565b6040518263ffffffff1660e01b81526004015f604051808303818588803b1580156111a9575f80fd5b505af11580156111bb573d5f803e3d5ffd5b50505050505b5f5b828110156112b7575f6101038989848181106111e1576111e16149a3565b90506020028101906111f391906149b7565b6040516112019291906149f9565b908152602001604051809103902054905061121b8161348d565b5f818152610102602090815260408083208054600260ff199182161782556005909101548452610105909252909120805490911690557f4e93215f00bc729272f0ff71afd3d0f385208cbf6c999fe776ad07c623b83466898984818110611284576112846149a3565b905060200281019061129691906149b7565b836040516112a693929190614a08565b60405180910390a1506001016111c3565b505f5b81811015611381575f6101038787848181106112d8576112d86149a3565b90506020028101906112ea91906149b7565b6040516112f89291906149f9565b90815260200160405180910390205490506113128161348d565b61131b8161351a565b7f596ee835bed6cb827d21ba1785c468f0755ee40d33d87132df5d2ec90b645f9f87878481811061134e5761134e6149a3565b905060200281019061136091906149b7565b8360405161137093929190614a08565b60405180910390a1506001016112ba565b5050505061138f600160c955565b505050505050565b60fb5460408051637a87fa0b60e01b815290516113ec923392600160501b9091046001600160a01b0316918291637a87fa0b9160048083019260209291908290030181865afa158015610fb5573d5f803e3d5ffd5b5f81815261010260205260409020436006820155805460ff19166004179055604080518281524360208201527fce479ab1b7a806fa3704c907b8fae15a191ad8da9a1671659e4f411f516c4c0191015b60405180910390a150565b60fb54611465903390600160501b90046001600160a01b03166136a9565b60fb805461ffff191661ffff83169081179091556040519081527f5fd0fcd821abb4c92d47c4740e5f4a25ef35e99ee092d170faa0e5cb47013c369060200161143c565b60fb5460408051633871d0f160e01b815290516114fe923392600160501b9091046001600160a01b0316918291633871d0f19160048083019260209291908290030181865afa158015610fb5573d5f803e3d5ffd5b60fb546040805163b479a51760e01b815290518392600160501b90046001600160a01b03169163b479a5179160048083019260209291908290030181865afa15801561154c573d5f803e3d5ffd5b505050506040513d601f19601f820116820180604052508101906115709190614965565b81111561159057604051639519af4360e01b815260040160405180910390fd5b5f5b818110156116e5575f6101038585848181106115b0576115b06149a3565b90506020028101906115c291906149b7565b6040516115d09291906149f9565b90815260200160405180910390205490506115ea8161372e565b611607576040516317136fff60e21b815260040160405180910390fd5b5f8181526101026020526040808220805460ff191660051781554360078201556004908101548251630bf8ac4960e41b815292516001600160a01b039091169363bf8ac49093808401939192919082900301818387803b158015611669575f80fd5b505af115801561167b573d5f803e3d5ffd5b505050507f450186694fefe67df6156f60235e4073b623160f28a0b85908ebc864316abf798585848181106116b2576116b26149a3565b90506020028101906116c491906149b7565b836040516116d493929190614a08565b60405180910390a150600101611592565b506116ef81613979565b505050565b6060825f03611716576040516334d6e01560e01b815260040160405180910390fd5b5f82611723600186614a42565b61172d9190614a2b565b611738906001614990565b90505f6117458483614990565b905060fd548111611756578061175a565b60fd545b90505f82821161176a575f611774565b6117748383614a42565b6001600160401b0381111561178b5761178b614672565b6040519080825280602002602001820160405280156117b4578160200160208202803683370190505b509050825b8281101561181e575f81815261010960205260409020546001600160a01b0316826117e48684614a42565b815181106117f4576117f46149a3565b6001600160a01b03909216602092830291909101909101528061181681614a55565b9150506117b9565b5095945050505050565b5f82815260656020526040902060010154611842816139c4565b6116ef83836139ce565b5f61010383836040516118609291906149f9565b9081526040519081900360200190205415159392505050565b6001600160a01b03811633146118ee5760405162461bcd60e51b815260206004820152602f60248201527f416363657373436f6e74726f6c3a2063616e206f6e6c792072656e6f756e636560448201526e103937b632b9903337b91039b2b63360891b60648201526084015b60405180910390fd5b6118f88282613a53565b5050565b60fb5461191a903390600160501b90046001600160a01b0316613ab9565b611922613b3e565b565b60fb600a9054906101000a90046001600160a01b03166001600160a01b0316636ccb9d706040518163ffffffff1660e01b8152600401602060405180830381865afa158015611975573d5f803e3d5ffd5b505050506040513d601f19601f8201168201806040525081019061199991906148ec565b6001600160a01b0316639f7053f584846040518363ffffffff1660e01b81526004016119c692919061492f565b5f604051808303815f87803b1580156119dd575f80fd5b505af11580156119ef573d5f803e3d5ffd5b505050506119fc816131f3565b611a0533613b90565b50335f9081526101066020908152604080832054808452610105909252909120600101611a33848683614aea565b505f81815261010560205260409081902060020180546001600160a01b0319166001600160a01b0385161790555133907fadc8722095edf061d7fdcb583105c05bf9eb15488503b621c39e254d8726977790611a9490879087908790614ba5565b60405180910390a250505050565b60fb5460408051637a87fa0b60e01b81529051611af7923392600160501b9091046001600160a01b0316918291637a87fa0b9160048083019260209291908290030181865afa158015610fb5573d5f803e3d5ffd5b806101015f828254611b099190614990565b9091555050610101546040519081527f5818a627697795ff3c3403f320c7549835866cfb64a0b06a6f7f077bc478e9f29060200161143c565b6101026020525f90815260409020805460018201805460ff9092169291611b6890614a6d565b80601f0160208091040260200160405190810160405280929190818152602001828054611b9490614a6d565b8015611bdf5780601f10611bb657610100808354040283529160200191611bdf565b820191905f5260205f20905b815481529060010190602001808311611bc257829003601f168201915b505050505090806002018054611bf490614a6d565b80601f0160208091040260200160405190810160405280929190818152602001828054611c2090614a6d565b8015611c6b5780601f10611c4257610100808354040283529160200191611c6b565b820191905f5260205f20905b815481529060010190602001808311611c4e57829003601f168201915b505050505090806003018054611c8090614a6d565b80601f0160208091040260200160405190810160405280929190818152602001828054611cac90614a6d565b8015611cf75780601f10611cce57610100808354040283529160200191611cf7565b820191905f5260205f20905b815481529060010190602001808311611cda57829003601f168201915b5050505060048301546005840154600685015460079095015493946001600160a01b039092169390925088565b611d2c6133a8565b60fb5460408051637a87fa0b60e01b81529051611d81923392600160501b9091046001600160a01b0316918291637a87fa0b9160048083019260209291908290030181865afa158015610fb5573d5f803e3d5ffd5b60fb600a9054906101000a90046001600160a01b03166001600160a01b0316639ca76b736040518163ffffffff1660e01b8152600401602060405180830381865afa158015611dd2573d5f803e3d5ffd5b505050506040513d601f19601f82011682018060405250810190611df691906148ec565b6001600160a01b0316631f033ef0826040518263ffffffff1660e01b81526004015f604051808303818588803b158015611e2e575f80fd5b505af1158015611e40573d5f803e3d5ffd5b50505050507f9407b62b10143b3ae08ce1cc7f9b66af41a4431ad59107e53ff54d6401e0730a81604051611e7691815260200190565b60405180910390a1611e88600160c955565b50565b60fb54611ea9903390600160501b90046001600160a01b0316613ab9565b60fb805469ffffffffffffffff00001916620100006001600160401b038481168202929092179283905560405192041681527facda2fe79efeffc359206ddeeb45f26ba1596223e01e1585458603af76e880a29060200161143c565b6060825f03611f27576040516334d6e01560e01b815260040160405180910390fd5b5f82611f34600186614a42565b611f3e9190614a2b565b90505f611f4b8483614990565b6001600160a01b0387165f9081526101066020526040812054919250819003611f875760405163240ebd5960e11b815260040160405180910390fd5b5f8181526101076020526040902054808311611fa35782611fa5565b805b92505f848411611fb5575f611fbf565b611fbf8585614a42565b6001600160401b03811115611fd657611fd6614672565b60405190808252806020026020018201604052801561200f57816020015b611ffc6141d9565b815260200190600190039081611ff45790505b509050845b848110156122af575f8481526101076020526040812080548390811061203c5761203c6149a3565b5f91825260208083209091015480835261010290915260409182902082516101008101909352805491935090829060ff16600581111561207e5761207e614578565b600581111561208f5761208f614578565b81526020016001820180546120a390614a6d565b80601f01602080910402602001604051908101604052809291908181526020018280546120cf90614a6d565b801561211a5780601f106120f15761010080835404028352916020019161211a565b820191905f5260205f20905b8154815290600101906020018083116120fd57829003601f168201915b5050505050815260200160028201805461213390614a6d565b80601f016020809104026020016040519081016040528092919081815260200182805461215f90614a6d565b80156121aa5780601f10612181576101008083540402835291602001916121aa565b820191905f5260205f20905b81548152906001019060200180831161218d57829003601f168201915b505050505081526020016003820180546121c390614a6d565b80601f01602080910402602001604051908101604052809291908181526020018280546121ef90614a6d565b801561223a5780601f106122115761010080835404028352916020019161223a565b820191905f5260205f20905b81548152906001019060200180831161221d57829003601f168201915b505050918352505060048201546001600160a01b031660208201526005820154604082015260068201546060820152600790910154608090910152836122808985614a42565b81518110612290576122906149a3565b60200260200101819052505080806122a790614a55565b915050612014565b5098975050505050505050565b5f6101005460ff546122ce9190614a42565b905090565b60fb546122f1903390600160501b90046001600160a01b0316613ab9565b611922613bfe565b5f8183111561231b5760405163096e13f760e21b815260040160405180910390fd5b6001600160a01b0384165f90815261010660205260408120549061234b825f908152610107602052604090205490565b905080841161235a578361235c565b805b93505f855b858110156123b9575f84815261010760205260408120805483908110612389576123896149a3565b905f5260205f200154905061239d81613c3b565b156123b057826123ac81614bd0565b9350505b50600101612361565b509695505050505050565b5f9182526065602090815260408084206001600160a01b0393909316845291905290205460ff1690565b6060825f03612410576040516334d6e01560e01b815260040160405180910390fd5b5f8261241d600186614a42565b6124279190614a2b565b612432906001614990565b90505f61243f8483614990565b905060fe5481116124505780612454565b60fe545b90505f846001600160401b0381111561246f5761246f614672565b6040519080825280602002602001820160405280156124a857816020015b6124956141d9565b81526020019060019003908161248d5790505b5090505f835b8381101561272d576124bf8161372e565b1561271b575f818152610102602052604090819020815161010081019092528054829060ff1660058111156124f6576124f6614578565b600581111561250757612507614578565b815260200160018201805461251b90614a6d565b80601f016020809104026020016040519081016040528092919081815260200182805461254790614a6d565b80156125925780601f1061256957610100808354040283529160200191612592565b820191905f5260205f20905b81548152906001019060200180831161257557829003601f168201915b505050505081526020016002820180546125ab90614a6d565b80601f01602080910402602001604051908101604052809291908181526020018280546125d790614a6d565b80156126225780601f106125f957610100808354040283529160200191612622565b820191905f5260205f20905b81548152906001019060200180831161260557829003601f168201915b5050505050815260200160038201805461263b90614a6d565b80601f016020809104026020016040519081016040528092919081815260200182805461266790614a6d565b80156126b25780601f10612689576101008083540402835291602001916126b2565b820191905f5260205f20905b81548152906001019060200180831161269557829003601f168201915b505050918352505060048201546001600160a01b0316602082015260058201546040820152600682015460608201526007909101546080909101528351849084908110612701576127016149a3565b6020026020010181905250818061271790614a55565b9250505b8061272581614a55565b9150506124ae565b50815295945050505050565b5f612743816139c4565b61274c826131f3565b60fb80547fffff0000000000000000000000000000000000000000ffffffffffffffffffff16600160501b6001600160a01b038516908102919091179091556040519081527fdb2219043d7b197cb235f1af0cf6d782d77dee3de19e3f4fb6d39aae633b44859060200160405180910390a15050565b60fb546127e0903390600160501b90046001600160a01b03166136a9565b60fc8190556040518181527f5d19c92c6893231b764f3320c712a4d056ff157295c8b620d893dbbed1a869b49060200161143c565b60fb5460408051637a87fa0b60e01b8152905161286a923392600160501b9091046001600160a01b0316918291637a87fa0b9160048083019260209291908290030181865afa158015610fb5573d5f803e3d5ffd5b6101008190556040518181527f711359152f2039f4182a096114b0d199c5f8e9cb268caff34080855c42ff4da99060200161143c565b6101056020525f90815260409020805460018201805460ff80841694610100909404169291906128cf90614a6d565b80601f01602080910402602001604051908101604052809291908181526020018280546128fb90614a6d565b80156129465780601f1061291d57610100808354040283529160200191612946565b820191905f5260205f20905b81548152906001019060200180831161292957829003601f168201915b50505050600283015460039093015491926001600160a01b039081169216905085565b5f82815260656020526040902060010154612983816139c4565b6116ef8383613a53565b610107602052815f5260405f2081815481106129a7575f80fd5b905f5260205f20015f91509150505481565b6129c16133a8565b6129c96131ad565b5f6129d333613b90565b90505f806129e388878686613ec1565b915091505f60fb600a9054906101000a90046001600160a01b03166001600160a01b03166318bcb2846040518163ffffffff1660e01b8152600401602060405180830381865afa158015612a39573d5f803e3d5ffd5b505050506040513d601f19601f82011682018060405250810190612a5d91906148ec565b90505f60fb600a9054906101000a90046001600160a01b03166001600160a01b0316636ccb9d706040518163ffffffff1660e01b8152600401602060405180830381865afa158015612ab1573d5f803e3d5ffd5b505050506040513d601f19601f82011682018060405250810190612ad591906148ec565b90505f5b84811015612f3457816001600160a01b0316639f55941b8d8d84818110612b0257612b026149a3565b9050602002810190612b1491906149b7565b8d8d86818110612b2657612b266149a3565b9050602002810190612b3891906149b7565b8d8d88818110612b4a57612b4a6149a3565b9050602002810190612b5c91906149b7565b6040518763ffffffff1660e01b8152600401612b7d96959493929190614bf5565b5f604051808303815f87803b158015612b94575f80fd5b505af1158015612ba6573d5f803e3d5ffd5b505050505f836001600160a01b0316637f70ce0d6001898589612bc99190614990565b60fe546040516001600160e01b031960e087901b16815260ff90941660048501526024840192909252604483015260648201526084016020604051808303815f875af1158015612c1b573d5f803e3d5ffd5b505050506040513d601f19601f82011682018060405250810190612c3f91906148ec565b604080516101008101909152909150805f81526020018e8e85818110612c6757612c676149a3565b9050602002810190612c7991906149b7565b8080601f0160208091040260200160405190810160405280939291908181526020018383808284375f920191909152505050908252506020018c8c85818110612cc457612cc46149a3565b9050602002810190612cd691906149b7565b8080601f0160208091040260200160405190810160405280939291908181526020018383808284375f920191909152505050908252506020018a8a85818110612d2157612d216149a3565b9050602002810190612d3391906149b7565b8080601f0160208091040260200160405190810160405280939291908181526020018383808284375f9201829052509385525050506001600160a01b03841660208084019190915260408084018c905260608401839052608090930182905260fe54825261010290522081518154829060ff19166001836005811115612dbb57612dbb614578565b021790555060208201516001820190612dd49082614c3d565b5060408201516002820190612de99082614c3d565b5060608201516003820190612dfe9082614c3d565b5060808201516004820180546001600160a01b0319166001600160a01b0390921691909117905560a0820151600582015560c0820151600682015560e09091015160079091015560fe546101038e8e85818110612e5d57612e5d6149a3565b9050602002810190612e6f91906149b7565b604051612e7d9291906149f9565b9081526040805160209281900383019020929092555f898152610107825291822060fe5481546001810183559184529190922090910155337fab5128638b64e6216e80dfafa70d3cb6d54913a536dc41e76eb4a04cfbe979cf8e8e85818110612ee857612ee86149a3565b9050602002810190612efa91906149b7565b60fe54604051612f0c93929190614a08565b60405180910390a260fe8054905f612f2383614a55565b919050555081600101915050612ad9565b50505050505061138f600160c955565b5f80612f4f33613b90565b5f818152610105602052604090205490915083151561010090910460ff16151503612f8d57604051635650952b60e01b815260040160405180910390fd5b60fb600a9054906101000a90046001600160a01b03166001600160a01b0316636e0fddfc6040518163ffffffff1660e01b8152600401602060405180830381865afa158015612fde573d5f803e3d5ffd5b505050506040513d601f19601f820116820180604052508101906130029190614965565b5f828152610108602052604090205461301b9190614990565b43101561303b5760405163111bb2f160e31b815260040160405180910390fd5b5f81815261010960205260409020546001600160a01b031691508215613132576001600160a01b03821631156130ba57816001600160a01b0316633ccfd60b6040518163ffffffff1660e01b81526004015f604051808303815f87803b1580156130a3575f80fd5b505af11580156130b5573d5f803e3d5ffd5b505050505b60fb600a9054906101000a90046001600160a01b03166001600160a01b0316630a3fbd9a6040518163ffffffff1660e01b8152600401602060405180830381865afa15801561310b573d5f803e3d5ffd5b505050506040513d601f19601f8201168201806040525081019061312f91906148ec565b91505b5f81815261010560209081526040808320805461ff0019166101008815159081029190911790915561010883529281902043908190558151858152928301939093528101919091527fc0465abaf1d51829975919c02418d521476b44f330a31d78bb6b4e96465e746b9060600160405180910390a150919050565b60975460ff16156119225760405162461bcd60e51b815260206004820152601060248201526f14185d5cd8589b194e881c185d5cd95960821b60448201526064016118e5565b6001600160a01b038116611e885760405163d92e233d60e01b815260040160405180910390fd5b6040518060a00160405280600115158152602001851515815260200184848080601f0160208091040260200160405190810160405280939291908181526020018383808284375f9201829052509385525050506001600160a01b0384166020808401919091523360409384015260fd548252610105815290829020835181549285015161ffff1990931690151561ff00191617610100921515929092029190911781559082015160018201906132d09082614c3d565b5060608201516002820180546001600160a01b03199081166001600160a01b039384161790915560809093015160039092018054909316911617905560fd8054335f9081526101066020908152604080832084905592825261010890529081204390558154919061334083614a55565b9190505550336001600160a01b03167f55b1a82a03cdb2847b1ec26dcac8ce8b3fc5f310388290b048c0ee9ac1ce8dd482600160fd546133809190614a42565b604080516001600160a01b039093168352602083019190915287151590820152606001611a94565b600260c954036133fa5760405162461bcd60e51b815260206004820152601f60248201527f5265656e7472616e637947756172643a207265656e7472616e742063616c6c0060448201526064016118e5565b600260c955565b6040516359891c9160e11b81526001600160a01b0384811660048301526024820183905283169063b312392290604401602060405180830381865afa15801561344c573d5f803e3d5ffd5b505050506040513d601f19601f82011682018060405250810190613470919061494a565b6116ef5760405163168dfea160e01b815260040160405180910390fd5b8015806134bb57505f818152610102602052604081205460ff1660058111156134b8576134b8614578565b14155b15611e88576040516317136fff60e21b815260040160405180910390fd5b5f81815261010260209081526040808320805460ff1916600317905560ff80548452610104909252822083905580549161351283614a55565b919050555050565b5f81815261010260209081526040808320805460ff19166001178155600501548084526101058352928190206003015460fb54825163278671bb60e01b815292516001600160a01b0392831694600160501b9092049092169263278671bb92600480830193928290030181865afa158015613597573d5f803e3d5ffd5b505050506040513d601f19601f820116820180604052508101906135bb91906148ec565b6001600160a01b031663aa67c91960fb600a9054906101000a90046001600160a01b03166001600160a01b031663082976456040518163ffffffff1660e01b8152600401602060405180830381865afa15801561361a573d5f803e3d5ffd5b505050506040513d601f19601f8201168201806040525081019061363e9190614965565b61365090673782dace9d900000614a42565b6040516001600160e01b031960e084901b1681526001600160a01b03851660048201526024015f604051808303818588803b15801561368d575f80fd5b505af115801561369f573d5f803e3d5ffd5b5050505050505050565b6040516353f5713b60e01b81526001600160a01b0383811660048301528216906353f5713b90602401602060405180830381865afa1580156136ed573d5f803e3d5ffd5b505050506040513d601f19601f82011682018060405250810190613711919061494a565b6118f85760405163a5523ee560e01b815260040160405180910390fd5b5f818152610102602052604080822081516101008101909252805483929190829060ff16600581111561376357613763614578565b600581111561377457613774614578565b815260200160018201805461378890614a6d565b80601f01602080910402602001604051908101604052809291908181526020018280546137b490614a6d565b80156137ff5780601f106137d6576101008083540402835291602001916137ff565b820191905f5260205f20905b8154815290600101906020018083116137e257829003601f168201915b5050505050815260200160028201805461381890614a6d565b80601f016020809104026020016040519081016040528092919081815260200182805461384490614a6d565b801561388f5780601f106138665761010080835404028352916020019161388f565b820191905f5260205f20905b81548152906001019060200180831161387257829003601f168201915b505050505081526020016003820180546138a890614a6d565b80601f01602080910402602001604051908101604052809291908181526020018280546138d490614a6d565b801561391f5780601f106138f65761010080835404028352916020019161391f565b820191905f5260205f20905b81548152906001019060200180831161390257829003601f168201915b50505091835250506004828101546001600160a01b0316602083015260058301546040830152600683015460608301526007909201546080909101529091508151600581111561397157613971614578565b149392505050565b806101015f82825461398b9190614a42565b9091555050610101546040519081527f5040a06a11b7d9b75fc56fbbd207905dbaa4ac86c0dc9cc7fff40cd1d92aece39060200161143c565b611e888133613f89565b6139d882826123c4565b6118f8575f8281526065602090815260408083206001600160a01b03851684529091529020805460ff19166001179055613a0f3390565b6001600160a01b0316816001600160a01b0316837f2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d60405160405180910390a45050565b613a5d82826123c4565b156118f8575f8281526065602090815260408083206001600160a01b0385168085529252808320805460ff1916905551339285917ff6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b9190a45050565b6040516318903ee760e21b81526001600160a01b038381166004830152821690636240fb9c90602401602060405180830381865afa158015613afd573d5f803e3d5ffd5b505050506040513d601f19601f82011682018060405250810190613b21919061494a565b6118f85760405163c4230ae360e01b815260040160405180910390fd5b613b46613fe2565b6097805460ff191690557f5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa335b6040516001600160a01b03909116815260200160405180910390a1565b6001600160a01b0381165f908152610106602052604081205490819003613bca5760405163240ebd5960e11b815260040160405180910390fd5b5f818152610105602052604090205460ff16613bf95760405163c11cb1df60e01b815260040160405180910390fd5b919050565b613c066131ad565b6097805460ff191660011790557f62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258613b733390565b5f818152610102602052604080822081516101008101909252805483929190829060ff166005811115613c7057613c70614578565b6005811115613c8157613c81614578565b8152602001600182018054613c9590614a6d565b80601f0160208091040260200160405190810160405280929190818152602001828054613cc190614a6d565b8015613d0c5780601f10613ce357610100808354040283529160200191613d0c565b820191905f5260205f20905b815481529060010190602001808311613cef57829003601f168201915b50505050508152602001600282018054613d2590614a6d565b80601f0160208091040260200160405190810160405280929190818152602001828054613d5190614a6d565b8015613d9c5780601f10613d7357610100808354040283529160200191613d9c565b820191905f5260205f20905b815481529060010190602001808311613d7f57829003601f168201915b50505050508152602001600382018054613db590614a6d565b80601f0160208091040260200160405190810160405280929190818152602001828054613de190614a6d565b8015613e2c5780601f10613e0357610100808354040283529160200191613e2c565b820191905f5260205f20905b815481529060010190602001808311613e0f57829003601f168201915b505050918352505060048201546001600160a01b0316602082015260058083015460408301526006830154606083015260079092015460809091015290915081516005811115613e7e57613e7e614578565b1480613e9c5750600281516005811115613e9a57613e9a614578565b145b80613eb95750600181516005811115613eb757613eb7614578565b145b159392505050565b5f808486141580613ed25750838614155b15613ef05760405163e5fe884360e01b815260040160405180910390fd5b859150811580613f05575060fb5461ffff1682115b15613f23576040516379b348ff60e11b815260040160405180910390fd5b505f828152610107602052604081205490613f3f3382846122f9565b60fb546001600160401b03918216925062010000900416613f608483614990565b1115613f7f57604051633e10caad60e21b815260040160405180910390fd5b5094509492505050565b613f9382826123c4565b6118f857613fa08161402b565b613fab83602061403d565b604051602001613fbc929190614cf8565b60408051601f198184030181529082905262461bcd60e51b82526118e591600401614d6c565b60975460ff166119225760405162461bcd60e51b815260206004820152601460248201527314185d5cd8589b194e881b9bdd081c185d5cd95960621b60448201526064016118e5565b6060610b2b6001600160a01b03831660145b60605f61404b836002614a2b565b614056906002614990565b6001600160401b0381111561406d5761406d614672565b6040519080825280601f01601f191660200182016040528015614097576020820181803683370190505b509050600360fc1b815f815181106140b1576140b16149a3565b60200101906001600160f81b03191690815f1a905350600f60fb1b816001815181106140df576140df6149a3565b60200101906001600160f81b03191690815f1a9053505f614101846002614a2b565b61410c906001614990565b90505b6001811115614183576f181899199a1a9b1b9c1cb0b131b232b360811b85600f1660108110614140576141406149a3565b1a60f81b828281518110614156576141566149a3565b60200101906001600160f81b03191690815f1a90535060049490941c9361417c81614d7e565b905061410f565b5083156141d25760405162461bcd60e51b815260206004820181905260248201527f537472696e67733a20686578206c656e67746820696e73756666696369656e7460448201526064016118e5565b9392505050565b604080516101008101909152805f81526020016060815260200160608152602001606081526020015f6001600160a01b031681526020015f81526020015f81526020015f81525090565b5f60208284031215614233575f80fd5b81356001600160e01b0319811681146141d2575f80fd5b8015158114611e88575f80fd5b5f8083601f840112614267575f80fd5b5081356001600160401b0381111561427d575f80fd5b602083019150836020828501011115614294575f80fd5b9250929050565b6001600160a01b0381168114611e88575f80fd5b5f805f80606085870312156142c2575f80fd5b84356142cd8161424a565b935060208501356001600160401b038111156142e7575f80fd5b6142f387828801614257565b90945092505060408501356143078161429b565b939692955090935050565b5f8083601f840112614322575f80fd5b5081356001600160401b03811115614338575f80fd5b6020830191508360208260051b8501011115614294575f80fd5b5f805f805f8060608789031215614367575f80fd5b86356001600160401b038082111561437d575f80fd5b6143898a838b01614312565b909850965060208901359150808211156143a1575f80fd5b6143ad8a838b01614312565b909650945060408901359150808211156143c5575f80fd5b506143d289828a01614312565b979a9699509497509295939492505050565b5f602082840312156143f4575f80fd5b5035919050565b5f6020828403121561440b575f80fd5b813561ffff811681146141d2575f80fd5b5f806020838503121561442d575f80fd5b82356001600160401b03811115614442575f80fd5b61444e85828601614312565b90969095509350505050565b5f806040838503121561446b575f80fd5b50508035926020909101359150565b602080825282518282018190525f9190848201906040850190845b818110156144ba5783516001600160a01b031683529284019291840191600101614495565b50909695505050505050565b5f80604083850312156144d7575f80fd5b8235915060208301356144e98161429b565b809150509250929050565b5f8060208385031215614505575f80fd5b82356001600160401b0381111561451a575f80fd5b61444e85828601614257565b5f805f60408486031215614538575f80fd5b83356001600160401b0381111561454d575f80fd5b61455986828701614257565b909450925050602084013561456d8161429b565b809150509250925092565b634e487b7160e01b5f52602160045260245ffd5b600681106145a857634e487b7160e01b5f52602160045260245ffd5b9052565b5f5b838110156145c65781810151838201526020016145ae565b50505f910152565b5f81518084526145e58160208601602086016145ac565b601f01601f19169290920160200192915050565b5f610100614607838c61458c565b8060208401526146198184018b6145ce565b9050828103604084015261462d818a6145ce565b9050828103606084015261464181896145ce565b6001600160a01b03979097166080840152505060a081019390935260c083019190915260e090910152949350505050565b634e487b7160e01b5f52604160045260245ffd5b5f60208284031215614696575f80fd5b81356001600160401b03808211156146ac575f80fd5b818401915084601f8301126146bf575f80fd5b8135818111156146d1576146d1614672565b604051601f8201601f19908116603f011681019083821181831017156146f9576146f9614672565b81604052828152876020848701011115614711575f80fd5b826020860160208301375f928101602001929092525095945050505050565b5f60208284031215614740575f80fd5b81356001600160401b03811681146141d2575f80fd5b5f805f60608486031215614768575f80fd5b83356147738161429b565b95602085013595506040909401359392505050565b5f6020808301818452808551808352604092508286019150828160051b8701018488015f5b8381101561486457603f1989840301855281516101006147ce85835161458c565b88820151818a8701526147e3828701826145ce565b91505087820151858203898701526147fb82826145ce565b9150506060808301518683038288015261481583826145ce565b92505050608080830151614833828801826001600160a01b03169052565b505060a0828101519086015260c0808301519086015260e091820151919094015293860193908601906001016147ad565b509098975050505050505050565b5f60208284031215614882575f80fd5b81356141d28161429b565b8515158152841515602082015260a060408201525f6148af60a08301866145ce565b6001600160a01b03948516606084015292909316608090910152949350505050565b5f602082840312156148e1575f80fd5b81356141d28161424a565b5f602082840312156148fc575f80fd5b81516141d28161429b565b81835281816020850137505f828201602090810191909152601f909101601f19169091010190565b602081525f614942602083018486614907565b949350505050565b5f6020828403121561495a575f80fd5b81516141d28161424a565b5f60208284031215614975575f80fd5b5051919050565b634e487b7160e01b5f52601160045260245ffd5b80820180821115610b2b57610b2b61497c565b634e487b7160e01b5f52603260045260245ffd5b5f808335601e198436030181126149cc575f80fd5b8301803591506001600160401b038211156149e5575f80fd5b602001915036819003821315614294575f80fd5b818382375f9101908152919050565b604081525f614a1b604083018587614907565b9050826020830152949350505050565b8082028115828204841417610b2b57610b2b61497c565b81810381811115610b2b57610b2b61497c565b5f60018201614a6657614a6661497c565b5060010190565b600181811c90821680614a8157607f821691505b602082108103614a9f57634e487b7160e01b5f52602260045260245ffd5b50919050565b601f8211156116ef575f81815260208120601f850160051c81016020861015614acb5750805b601f850160051c820191505b8181101561138f57828155600101614ad7565b6001600160401b03831115614b0157614b01614672565b614b1583614b0f8354614a6d565b83614aa5565b5f601f841160018114614b46575f8515614b2f5750838201355b5f19600387901b1c1916600186901b178355614b9e565b5f83815260209020601f19861690835b82811015614b765786850135825560209485019460019092019101614b56565b5086821015614b92575f1960f88860031b161c19848701351681555b505060018560011b0183555b5050505050565b604081525f614bb8604083018587614907565b905060018060a01b0383166020830152949350505050565b5f6001600160401b03808316818103614beb57614beb61497c565b6001019392505050565b606081525f614c0860608301888a614907565b8281036020840152614c1b818789614907565b90508281036040840152614c30818587614907565b9998505050505050505050565b81516001600160401b03811115614c5657614c56614672565b614c6a81614c648454614a6d565b84614aa5565b602080601f831160018114614c9d575f8415614c865750858301515b5f19600386901b1c1916600185901b17855561138f565b5f85815260208120601f198616915b82811015614ccb57888601518255948401946001909101908401614cac565b5085821015614ce857878501515f19600388901b60f8161c191681555b5050505050600190811b01905550565b7f416363657373436f6e74726f6c3a206163636f756e742000000000000000000081525f8351614d2f8160178501602088016145ac565b7001034b99036b4b9b9b4b733903937b6329607d1b6017918401918201528351614d608160288401602088016145ac565b01602801949350505050565b602081525f6141d260208301846145ce565b5f81614d8c57614d8c61497c565b505f19019056fea26469706673582212207132153208e1744f30ebd23b7fb725c796ba18c64b55f21b0d65f96e536221c664736f6c63430008140033496e697469616c697a61626c653a20636f6e7472616374206973206e6f742069",
}

// PermissionlessNodeRegistryABI is the input ABI used to generate the binding from.
// Deprecated: Use PermissionlessNodeRegistryMetaData.ABI instead.
var PermissionlessNodeRegistryABI = PermissionlessNodeRegistryMetaData.ABI

// PermissionlessNodeRegistryBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use PermissionlessNodeRegistryMetaData.Bin instead.
var PermissionlessNodeRegistryBin = PermissionlessNodeRegistryMetaData.Bin

// DeployPermissionlessNodeRegistry deploys a new Ethereum contract, binding an instance of PermissionlessNodeRegistry to it.
func DeployPermissionlessNodeRegistry(auth *bind.TransactOpts, backend bind.ContractBackend, _admin common.Address, _staderConfig common.Address) (common.Address, *types.Transaction, *PermissionlessNodeRegistry, error) {
	parsed, err := PermissionlessNodeRegistryMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(PermissionlessNodeRegistryBin), backend, _admin, _staderConfig)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &PermissionlessNodeRegistry{PermissionlessNodeRegistryCaller: PermissionlessNodeRegistryCaller{contract: contract}, PermissionlessNodeRegistryTransactor: PermissionlessNodeRegistryTransactor{contract: contract}, PermissionlessNodeRegistryFilterer: PermissionlessNodeRegistryFilterer{contract: contract}}, nil
}

// PermissionlessNodeRegistry is an auto generated Go binding around an Ethereum contract.
type PermissionlessNodeRegistry struct {
	PermissionlessNodeRegistryCaller     // Read-only binding to the contract
	PermissionlessNodeRegistryTransactor // Write-only binding to the contract
	PermissionlessNodeRegistryFilterer   // Log filterer for contract events
}

// PermissionlessNodeRegistryCaller is an auto generated read-only Go binding around an Ethereum contract.
type PermissionlessNodeRegistryCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PermissionlessNodeRegistryTransactor is an auto generated write-only Go binding around an Ethereum contract.
type PermissionlessNodeRegistryTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PermissionlessNodeRegistryFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type PermissionlessNodeRegistryFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PermissionlessNodeRegistrySession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type PermissionlessNodeRegistrySession struct {
	Contract     *PermissionlessNodeRegistry // Generic contract binding to set the session for
	CallOpts     bind.CallOpts               // Call options to use throughout this session
	TransactOpts bind.TransactOpts           // Transaction auth options to use throughout this session
}

// PermissionlessNodeRegistryCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type PermissionlessNodeRegistryCallerSession struct {
	Contract *PermissionlessNodeRegistryCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                     // Call options to use throughout this session
}

// PermissionlessNodeRegistryTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type PermissionlessNodeRegistryTransactorSession struct {
	Contract     *PermissionlessNodeRegistryTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                     // Transaction auth options to use throughout this session
}

// PermissionlessNodeRegistryRaw is an auto generated low-level Go binding around an Ethereum contract.
type PermissionlessNodeRegistryRaw struct {
	Contract *PermissionlessNodeRegistry // Generic contract binding to access the raw methods on
}

// PermissionlessNodeRegistryCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type PermissionlessNodeRegistryCallerRaw struct {
	Contract *PermissionlessNodeRegistryCaller // Generic read-only contract binding to access the raw methods on
}

// PermissionlessNodeRegistryTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type PermissionlessNodeRegistryTransactorRaw struct {
	Contract *PermissionlessNodeRegistryTransactor // Generic write-only contract binding to access the raw methods on
}

// NewPermissionlessNodeRegistry creates a new instance of PermissionlessNodeRegistry, bound to a specific deployed contract.
func NewPermissionlessNodeRegistry(address common.Address, backend bind.ContractBackend) (*PermissionlessNodeRegistry, error) {
	contract, err := bindPermissionlessNodeRegistry(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &PermissionlessNodeRegistry{PermissionlessNodeRegistryCaller: PermissionlessNodeRegistryCaller{contract: contract}, PermissionlessNodeRegistryTransactor: PermissionlessNodeRegistryTransactor{contract: contract}, PermissionlessNodeRegistryFilterer: PermissionlessNodeRegistryFilterer{contract: contract}}, nil
}

// NewPermissionlessNodeRegistryCaller creates a new read-only instance of PermissionlessNodeRegistry, bound to a specific deployed contract.
func NewPermissionlessNodeRegistryCaller(address common.Address, caller bind.ContractCaller) (*PermissionlessNodeRegistryCaller, error) {
	contract, err := bindPermissionlessNodeRegistry(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &PermissionlessNodeRegistryCaller{contract: contract}, nil
}

// NewPermissionlessNodeRegistryTransactor creates a new write-only instance of PermissionlessNodeRegistry, bound to a specific deployed contract.
func NewPermissionlessNodeRegistryTransactor(address common.Address, transactor bind.ContractTransactor) (*PermissionlessNodeRegistryTransactor, error) {
	contract, err := bindPermissionlessNodeRegistry(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &PermissionlessNodeRegistryTransactor{contract: contract}, nil
}

// NewPermissionlessNodeRegistryFilterer creates a new log filterer instance of PermissionlessNodeRegistry, bound to a specific deployed contract.
func NewPermissionlessNodeRegistryFilterer(address common.Address, filterer bind.ContractFilterer) (*PermissionlessNodeRegistryFilterer, error) {
	contract, err := bindPermissionlessNodeRegistry(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &PermissionlessNodeRegistryFilterer{contract: contract}, nil
}

// bindPermissionlessNodeRegistry binds a generic wrapper to an already deployed contract.
func bindPermissionlessNodeRegistry(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := PermissionlessNodeRegistryMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_PermissionlessNodeRegistry *PermissionlessNodeRegistryRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _PermissionlessNodeRegistry.Contract.PermissionlessNodeRegistryCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_PermissionlessNodeRegistry *PermissionlessNodeRegistryRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PermissionlessNodeRegistry.Contract.PermissionlessNodeRegistryTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_PermissionlessNodeRegistry *PermissionlessNodeRegistryRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _PermissionlessNodeRegistry.Contract.PermissionlessNodeRegistryTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_PermissionlessNodeRegistry *PermissionlessNodeRegistryCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _PermissionlessNodeRegistry.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_PermissionlessNodeRegistry *PermissionlessNodeRegistryTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PermissionlessNodeRegistry.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_PermissionlessNodeRegistry *PermissionlessNodeRegistryTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _PermissionlessNodeRegistry.Contract.contract.Transact(opts, method, params...)
}

// COLLATERALETH is a free data retrieval call binding the contract method 0x4f59ed80.
//
// Solidity: function COLLATERAL_ETH() view returns(uint256)
func (_PermissionlessNodeRegistry *PermissionlessNodeRegistryCaller) COLLATERALETH(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _PermissionlessNodeRegistry.contract.Call(opts, &out, "COLLATERAL_ETH")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// COLLATERALETH is a free data retrieval call binding the contract method 0x4f59ed80.
//
// Solidity: function COLLATERAL_ETH() view returns(uint256)
func (_PermissionlessNodeRegistry *PermissionlessNodeRegistrySession) COLLATERALETH() (*big.Int, error) {
	return _PermissionlessNodeRegistry.Contract.COLLATERALETH(&_PermissionlessNodeRegistry.CallOpts)
}

// COLLATERALETH is a free data retrieval call binding the contract method 0x4f59ed80.
//
// Solidity: function COLLATERAL_ETH() view returns(uint256)
func (_PermissionlessNodeRegistry *PermissionlessNodeRegistryCallerSession) COLLATERALETH() (*big.Int, error) {
	return _PermissionlessNodeRegistry.Contract.COLLATERALETH(&_PermissionlessNodeRegistry.CallOpts)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_PermissionlessNodeRegistry *PermissionlessNodeRegistryCaller) DEFAULTADMINROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _PermissionlessNodeRegistry.contract.Call(opts, &out, "DEFAULT_ADMIN_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_PermissionlessNodeRegistry *PermissionlessNodeRegistrySession) DEFAULTADMINROLE() ([32]byte, error) {
	return _PermissionlessNodeRegistry.Contract.DEFAULTADMINROLE(&_PermissionlessNodeRegistry.CallOpts)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_PermissionlessNodeRegistry *PermissionlessNodeRegistryCallerSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _PermissionlessNodeRegistry.Contract.DEFAULTADMINROLE(&_PermissionlessNodeRegistry.CallOpts)
}

// FRONTRUNPENALTY is a free data retrieval call binding the contract method 0xbb7306bf.
//
// Solidity: function FRONT_RUN_PENALTY() view returns(uint256)
func (_PermissionlessNodeRegistry *PermissionlessNodeRegistryCaller) FRONTRUNPENALTY(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _PermissionlessNodeRegistry.contract.Call(opts, &out, "FRONT_RUN_PENALTY")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// FRONTRUNPENALTY is a free data retrieval call binding the contract method 0xbb7306bf.
//
// Solidity: function FRONT_RUN_PENALTY() view returns(uint256)
func (_PermissionlessNodeRegistry *PermissionlessNodeRegistrySession) FRONTRUNPENALTY() (*big.Int, error) {
	return _PermissionlessNodeRegistry.Contract.FRONTRUNPENALTY(&_PermissionlessNodeRegistry.CallOpts)
}

// FRONTRUNPENALTY is a free data retrieval call binding the contract method 0xbb7306bf.
//
// Solidity: function FRONT_RUN_PENALTY() view returns(uint256)
func (_PermissionlessNodeRegistry *PermissionlessNodeRegistryCallerSession) FRONTRUNPENALTY() (*big.Int, error) {
	return _PermissionlessNodeRegistry.Contract.FRONTRUNPENALTY(&_PermissionlessNodeRegistry.CallOpts)
}

// POOLID is a free data retrieval call binding the contract method 0xe0d7d0e9.
//
// Solidity: function POOL_ID() view returns(uint8)
func (_PermissionlessNodeRegistry *PermissionlessNodeRegistryCaller) POOLID(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _PermissionlessNodeRegistry.contract.Call(opts, &out, "POOL_ID")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// POOLID is a free data retrieval call binding the contract method 0xe0d7d0e9.
//
// Solidity: function POOL_ID() view returns(uint8)
func (_PermissionlessNodeRegistry *PermissionlessNodeRegistrySession) POOLID() (uint8, error) {
	return _PermissionlessNodeRegistry.Contract.POOLID(&_PermissionlessNodeRegistry.CallOpts)
}

// POOLID is a free data retrieval call binding the contract method 0xe0d7d0e9.
//
// Solidity: function POOL_ID() view returns(uint8)
func (_PermissionlessNodeRegistry *PermissionlessNodeRegistryCallerSession) POOLID() (uint8, error) {
	return _PermissionlessNodeRegistry.Contract.POOLID(&_PermissionlessNodeRegistry.CallOpts)
}

// GetAllActiveValidators is a free data retrieval call binding the contract method 0x99888898.
//
// Solidity: function getAllActiveValidators(uint256 _pageNumber, uint256 _pageSize) view returns((uint8,bytes,bytes,bytes,address,uint256,uint256,uint256)[])
func (_PermissionlessNodeRegistry *PermissionlessNodeRegistryCaller) GetAllActiveValidators(opts *bind.CallOpts, _pageNumber *big.Int, _pageSize *big.Int) ([]Validator, error) {
	var out []interface{}
	err := _PermissionlessNodeRegistry.contract.Call(opts, &out, "getAllActiveValidators", _pageNumber, _pageSize)

	if err != nil {
		return *new([]Validator), err
	}

	out0 := *abi.ConvertType(out[0], new([]Validator)).(*[]Validator)

	return out0, err

}

// GetAllActiveValidators is a free data retrieval call binding the contract method 0x99888898.
//
// Solidity: function getAllActiveValidators(uint256 _pageNumber, uint256 _pageSize) view returns((uint8,bytes,bytes,bytes,address,uint256,uint256,uint256)[])
func (_PermissionlessNodeRegistry *PermissionlessNodeRegistrySession) GetAllActiveValidators(_pageNumber *big.Int, _pageSize *big.Int) ([]Validator, error) {
	return _PermissionlessNodeRegistry.Contract.GetAllActiveValidators(&_PermissionlessNodeRegistry.CallOpts, _pageNumber, _pageSize)
}

// GetAllActiveValidators is a free data retrieval call binding the contract method 0x99888898.
//
// Solidity: function getAllActiveValidators(uint256 _pageNumber, uint256 _pageSize) view returns((uint8,bytes,bytes,bytes,address,uint256,uint256,uint256)[])
func (_PermissionlessNodeRegistry *PermissionlessNodeRegistryCallerSession) GetAllActiveValidators(_pageNumber *big.Int, _pageSize *big.Int) ([]Validator, error) {
	return _PermissionlessNodeRegistry.Contract.GetAllActiveValidators(&_PermissionlessNodeRegistry.CallOpts, _pageNumber, _pageSize)
}

// GetAllNodeELVaultAddress is a free data retrieval call binding the contract method 0x2d32924f.
//
// Solidity: function getAllNodeELVaultAddress(uint256 _pageNumber, uint256 _pageSize) view returns(address[])
func (_PermissionlessNodeRegistry *PermissionlessNodeRegistryCaller) GetAllNodeELVaultAddress(opts *bind.CallOpts, _pageNumber *big.Int, _pageSize *big.Int) ([]common.Address, error) {
	var out []interface{}
	err := _PermissionlessNodeRegistry.contract.Call(opts, &out, "getAllNodeELVaultAddress", _pageNumber, _pageSize)

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// GetAllNodeELVaultAddress is a free data retrieval call binding the contract method 0x2d32924f.
//
// Solidity: function getAllNodeELVaultAddress(uint256 _pageNumber, uint256 _pageSize) view returns(address[])
func (_PermissionlessNodeRegistry *PermissionlessNodeRegistrySession) GetAllNodeELVaultAddress(_pageNumber *big.Int, _pageSize *big.Int) ([]common.Address, error) {
	return _PermissionlessNodeRegistry.Contract.GetAllNodeELVaultAddress(&_PermissionlessNodeRegistry.CallOpts, _pageNumber, _pageSize)
}

// GetAllNodeELVaultAddress is a free data retrieval call binding the contract method 0x2d32924f.
//
// Solidity: function getAllNodeELVaultAddress(uint256 _pageNumber, uint256 _pageSize) view returns(address[])
func (_PermissionlessNodeRegistry *PermissionlessNodeRegistryCallerSession) GetAllNodeELVaultAddress(_pageNumber *big.Int, _pageSize *big.Int) ([]common.Address, error) {
	return _PermissionlessNodeRegistry.Contract.GetAllNodeELVaultAddress(&_PermissionlessNodeRegistry.CallOpts, _pageNumber, _pageSize)
}

// GetCollateralETH is a free data retrieval call binding the contract method 0xb01db078.
//
// Solidity: function getCollateralETH() pure returns(uint256)
func (_PermissionlessNodeRegistry *PermissionlessNodeRegistryCaller) GetCollateralETH(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _PermissionlessNodeRegistry.contract.Call(opts, &out, "getCollateralETH")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetCollateralETH is a free data retrieval call binding the contract method 0xb01db078.
//
// Solidity: function getCollateralETH() pure returns(uint256)
func (_PermissionlessNodeRegistry *PermissionlessNodeRegistrySession) GetCollateralETH() (*big.Int, error) {
	return _PermissionlessNodeRegistry.Contract.GetCollateralETH(&_PermissionlessNodeRegistry.CallOpts)
}

// GetCollateralETH is a free data retrieval call binding the contract method 0xb01db078.
//
// Solidity: function getCollateralETH() pure returns(uint256)
func (_PermissionlessNodeRegistry *PermissionlessNodeRegistryCallerSession) GetCollateralETH() (*big.Int, error) {
	return _PermissionlessNodeRegistry.Contract.GetCollateralETH(&_PermissionlessNodeRegistry.CallOpts)
}

// GetOperatorRewardAddress is a free data retrieval call binding the contract method 0x83ea2358.
//
// Solidity: function getOperatorRewardAddress(uint256 _operatorId) view returns(address)
func (_PermissionlessNodeRegistry *PermissionlessNodeRegistryCaller) GetOperatorRewardAddress(opts *bind.CallOpts, _operatorId *big.Int) (common.Address, error) {
	var out []interface{}
	err := _PermissionlessNodeRegistry.contract.Call(opts, &out, "getOperatorRewardAddress", _operatorId)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetOperatorRewardAddress is a free data retrieval call binding the contract method 0x83ea2358.
//
// Solidity: function getOperatorRewardAddress(uint256 _operatorId) view returns(address)
func (_PermissionlessNodeRegistry *PermissionlessNodeRegistrySession) GetOperatorRewardAddress(_operatorId *big.Int) (common.Address, error) {
	return _PermissionlessNodeRegistry.Contract.GetOperatorRewardAddress(&_PermissionlessNodeRegistry.CallOpts, _operatorId)
}

// GetOperatorRewardAddress is a free data retrieval call binding the contract method 0x83ea2358.
//
// Solidity: function getOperatorRewardAddress(uint256 _operatorId) view returns(address)
func (_PermissionlessNodeRegistry *PermissionlessNodeRegistryCallerSession) GetOperatorRewardAddress(_operatorId *big.Int) (common.Address, error) {
	return _PermissionlessNodeRegistry.Contract.GetOperatorRewardAddress(&_PermissionlessNodeRegistry.CallOpts, _operatorId)
}

// GetOperatorTotalKeys is a free data retrieval call binding the contract method 0xc34ade5c.
//
// Solidity: function getOperatorTotalKeys(uint256 _operatorId) view returns(uint256 _totalKeys)
func (_PermissionlessNodeRegistry *PermissionlessNodeRegistryCaller) GetOperatorTotalKeys(opts *bind.CallOpts, _operatorId *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _PermissionlessNodeRegistry.contract.Call(opts, &out, "getOperatorTotalKeys", _operatorId)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetOperatorTotalKeys is a free data retrieval call binding the contract method 0xc34ade5c.
//
// Solidity: function getOperatorTotalKeys(uint256 _operatorId) view returns(uint256 _totalKeys)
func (_PermissionlessNodeRegistry *PermissionlessNodeRegistrySession) GetOperatorTotalKeys(_operatorId *big.Int) (*big.Int, error) {
	return _PermissionlessNodeRegistry.Contract.GetOperatorTotalKeys(&_PermissionlessNodeRegistry.CallOpts, _operatorId)
}

// GetOperatorTotalKeys is a free data retrieval call binding the contract method 0xc34ade5c.
//
// Solidity: function getOperatorTotalKeys(uint256 _operatorId) view returns(uint256 _totalKeys)
func (_PermissionlessNodeRegistry *PermissionlessNodeRegistryCallerSession) GetOperatorTotalKeys(_operatorId *big.Int) (*big.Int, error) {
	return _PermissionlessNodeRegistry.Contract.GetOperatorTotalKeys(&_PermissionlessNodeRegistry.CallOpts, _operatorId)
}

// GetOperatorTotalNonTerminalKeys is a free data retrieval call binding the contract method 0x8a25bcec.
//
// Solidity: function getOperatorTotalNonTerminalKeys(address _nodeOperator, uint256 _startIndex, uint256 _endIndex) view returns(uint64)
func (_PermissionlessNodeRegistry *PermissionlessNodeRegistryCaller) GetOperatorTotalNonTerminalKeys(opts *bind.CallOpts, _nodeOperator common.Address, _startIndex *big.Int, _endIndex *big.Int) (uint64, error) {
	var out []interface{}
	err := _PermissionlessNodeRegistry.contract.Call(opts, &out, "getOperatorTotalNonTerminalKeys", _nodeOperator, _startIndex, _endIndex)

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// GetOperatorTotalNonTerminalKeys is a free data retrieval call binding the contract method 0x8a25bcec.
//
// Solidity: function getOperatorTotalNonTerminalKeys(address _nodeOperator, uint256 _startIndex, uint256 _endIndex) view returns(uint64)
func (_PermissionlessNodeRegistry *PermissionlessNodeRegistrySession) GetOperatorTotalNonTerminalKeys(_nodeOperator common.Address, _startIndex *big.Int, _endIndex *big.Int) (uint64, error) {
	return _PermissionlessNodeRegistry.Contract.GetOperatorTotalNonTerminalKeys(&_PermissionlessNodeRegistry.CallOpts, _nodeOperator, _startIndex, _endIndex)
}

// GetOperatorTotalNonTerminalKeys is a free data retrieval call binding the contract method 0x8a25bcec.
//
// Solidity: function getOperatorTotalNonTerminalKeys(address _nodeOperator, uint256 _startIndex, uint256 _endIndex) view returns(uint64)
func (_PermissionlessNodeRegistry *PermissionlessNodeRegistryCallerSession) GetOperatorTotalNonTerminalKeys(_nodeOperator common.Address, _startIndex *big.Int, _endIndex *big.Int) (uint64, error) {
	return _PermissionlessNodeRegistry.Contract.GetOperatorTotalNonTerminalKeys(&_PermissionlessNodeRegistry.CallOpts, _nodeOperator, _startIndex, _endIndex)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_PermissionlessNodeRegistry *PermissionlessNodeRegistryCaller) GetRoleAdmin(opts *bind.CallOpts, role [32]byte) ([32]byte, error) {
	var out []interface{}
	err := _PermissionlessNodeRegistry.contract.Call(opts, &out, "getRoleAdmin", role)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_PermissionlessNodeRegistry *PermissionlessNodeRegistrySession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _PermissionlessNodeRegistry.Contract.GetRoleAdmin(&_PermissionlessNodeRegistry.CallOpts, role)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_PermissionlessNodeRegistry *PermissionlessNodeRegistryCallerSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _PermissionlessNodeRegistry.Contract.GetRoleAdmin(&_PermissionlessNodeRegistry.CallOpts, role)
}

// GetSocializingPoolStateChangeBlock is a free data retrieval call binding the contract method 0xebb5c174.
//
// Solidity: function getSocializingPoolStateChangeBlock(uint256 _operatorId) view returns(uint256)
func (_PermissionlessNodeRegistry *PermissionlessNodeRegistryCaller) GetSocializingPoolStateChangeBlock(opts *bind.CallOpts, _operatorId *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _PermissionlessNodeRegistry.contract.Call(opts, &out, "getSocializingPoolStateChangeBlock", _operatorId)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetSocializingPoolStateChangeBlock is a free data retrieval call binding the contract method 0xebb5c174.
//
// Solidity: function getSocializingPoolStateChangeBlock(uint256 _operatorId) view returns(uint256)
func (_PermissionlessNodeRegistry *PermissionlessNodeRegistrySession) GetSocializingPoolStateChangeBlock(_operatorId *big.Int) (*big.Int, error) {
	return _PermissionlessNodeRegistry.Contract.GetSocializingPoolStateChangeBlock(&_PermissionlessNodeRegistry.CallOpts, _operatorId)
}

// GetSocializingPoolStateChangeBlock is a free data retrieval call binding the contract method 0xebb5c174.
//
// Solidity: function getSocializingPoolStateChangeBlock(uint256 _operatorId) view returns(uint256)
func (_PermissionlessNodeRegistry *PermissionlessNodeRegistryCallerSession) GetSocializingPoolStateChangeBlock(_operatorId *big.Int) (*big.Int, error) {
	return _PermissionlessNodeRegistry.Contract.GetSocializingPoolStateChangeBlock(&_PermissionlessNodeRegistry.CallOpts, _operatorId)
}

// GetTotalActiveValidatorCount is a free data retrieval call binding the contract method 0x77c359e1.
//
// Solidity: function getTotalActiveValidatorCount() view returns(uint256)
func (_PermissionlessNodeRegistry *PermissionlessNodeRegistryCaller) GetTotalActiveValidatorCount(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _PermissionlessNodeRegistry.contract.Call(opts, &out, "getTotalActiveValidatorCount")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetTotalActiveValidatorCount is a free data retrieval call binding the contract method 0x77c359e1.
//
// Solidity: function getTotalActiveValidatorCount() view returns(uint256)
func (_PermissionlessNodeRegistry *PermissionlessNodeRegistrySession) GetTotalActiveValidatorCount() (*big.Int, error) {
	return _PermissionlessNodeRegistry.Contract.GetTotalActiveValidatorCount(&_PermissionlessNodeRegistry.CallOpts)
}

// GetTotalActiveValidatorCount is a free data retrieval call binding the contract method 0x77c359e1.
//
// Solidity: function getTotalActiveValidatorCount() view returns(uint256)
func (_PermissionlessNodeRegistry *PermissionlessNodeRegistryCallerSession) GetTotalActiveValidatorCount() (*big.Int, error) {
	return _PermissionlessNodeRegistry.Contract.GetTotalActiveValidatorCount(&_PermissionlessNodeRegistry.CallOpts)
}

// GetTotalQueuedValidatorCount is a free data retrieval call binding the contract method 0x7bd977d9.
//
// Solidity: function getTotalQueuedValidatorCount() view returns(uint256)
func (_PermissionlessNodeRegistry *PermissionlessNodeRegistryCaller) GetTotalQueuedValidatorCount(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _PermissionlessNodeRegistry.contract.Call(opts, &out, "getTotalQueuedValidatorCount")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetTotalQueuedValidatorCount is a free data retrieval call binding the contract method 0x7bd977d9.
//
// Solidity: function getTotalQueuedValidatorCount() view returns(uint256)
func (_PermissionlessNodeRegistry *PermissionlessNodeRegistrySession) GetTotalQueuedValidatorCount() (*big.Int, error) {
	return _PermissionlessNodeRegistry.Contract.GetTotalQueuedValidatorCount(&_PermissionlessNodeRegistry.CallOpts)
}

// GetTotalQueuedValidatorCount is a free data retrieval call binding the contract method 0x7bd977d9.
//
// Solidity: function getTotalQueuedValidatorCount() view returns(uint256)
func (_PermissionlessNodeRegistry *PermissionlessNodeRegistryCallerSession) GetTotalQueuedValidatorCount() (*big.Int, error) {
	return _PermissionlessNodeRegistry.Contract.GetTotalQueuedValidatorCount(&_PermissionlessNodeRegistry.CallOpts)
}

// GetValidatorsByOperator is a free data retrieval call binding the contract method 0x683547b8.
//
// Solidity: function getValidatorsByOperator(address _operator, uint256 _pageNumber, uint256 _pageSize) view returns((uint8,bytes,bytes,bytes,address,uint256,uint256,uint256)[])
func (_PermissionlessNodeRegistry *PermissionlessNodeRegistryCaller) GetValidatorsByOperator(opts *bind.CallOpts, _operator common.Address, _pageNumber *big.Int, _pageSize *big.Int) ([]Validator, error) {
	var out []interface{}
	err := _PermissionlessNodeRegistry.contract.Call(opts, &out, "getValidatorsByOperator", _operator, _pageNumber, _pageSize)

	if err != nil {
		return *new([]Validator), err
	}

	out0 := *abi.ConvertType(out[0], new([]Validator)).(*[]Validator)

	return out0, err

}

// GetValidatorsByOperator is a free data retrieval call binding the contract method 0x683547b8.
//
// Solidity: function getValidatorsByOperator(address _operator, uint256 _pageNumber, uint256 _pageSize) view returns((uint8,bytes,bytes,bytes,address,uint256,uint256,uint256)[])
func (_PermissionlessNodeRegistry *PermissionlessNodeRegistrySession) GetValidatorsByOperator(_operator common.Address, _pageNumber *big.Int, _pageSize *big.Int) ([]Validator, error) {
	return _PermissionlessNodeRegistry.Contract.GetValidatorsByOperator(&_PermissionlessNodeRegistry.CallOpts, _operator, _pageNumber, _pageSize)
}

// GetValidatorsByOperator is a free data retrieval call binding the contract method 0x683547b8.
//
// Solidity: function getValidatorsByOperator(address _operator, uint256 _pageNumber, uint256 _pageSize) view returns((uint8,bytes,bytes,bytes,address,uint256,uint256,uint256)[])
func (_PermissionlessNodeRegistry *PermissionlessNodeRegistryCallerSession) GetValidatorsByOperator(_operator common.Address, _pageNumber *big.Int, _pageSize *big.Int) ([]Validator, error) {
	return _PermissionlessNodeRegistry.Contract.GetValidatorsByOperator(&_PermissionlessNodeRegistry.CallOpts, _operator, _pageNumber, _pageSize)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_PermissionlessNodeRegistry *PermissionlessNodeRegistryCaller) HasRole(opts *bind.CallOpts, role [32]byte, account common.Address) (bool, error) {
	var out []interface{}
	err := _PermissionlessNodeRegistry.contract.Call(opts, &out, "hasRole", role, account)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_PermissionlessNodeRegistry *PermissionlessNodeRegistrySession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _PermissionlessNodeRegistry.Contract.HasRole(&_PermissionlessNodeRegistry.CallOpts, role, account)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_PermissionlessNodeRegistry *PermissionlessNodeRegistryCallerSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _PermissionlessNodeRegistry.Contract.HasRole(&_PermissionlessNodeRegistry.CallOpts, role, account)
}

// InputKeyCountLimit is a free data retrieval call binding the contract method 0xe0bf8b53.
//
// Solidity: function inputKeyCountLimit() view returns(uint16)
func (_PermissionlessNodeRegistry *PermissionlessNodeRegistryCaller) InputKeyCountLimit(opts *bind.CallOpts) (uint16, error) {
	var out []interface{}
	err := _PermissionlessNodeRegistry.contract.Call(opts, &out, "inputKeyCountLimit")

	if err != nil {
		return *new(uint16), err
	}

	out0 := *abi.ConvertType(out[0], new(uint16)).(*uint16)

	return out0, err

}

// InputKeyCountLimit is a free data retrieval call binding the contract method 0xe0bf8b53.
//
// Solidity: function inputKeyCountLimit() view returns(uint16)
func (_PermissionlessNodeRegistry *PermissionlessNodeRegistrySession) InputKeyCountLimit() (uint16, error) {
	return _PermissionlessNodeRegistry.Contract.InputKeyCountLimit(&_PermissionlessNodeRegistry.CallOpts)
}

// InputKeyCountLimit is a free data retrieval call binding the contract method 0xe0bf8b53.
//
// Solidity: function inputKeyCountLimit() view returns(uint16)
func (_PermissionlessNodeRegistry *PermissionlessNodeRegistryCallerSession) InputKeyCountLimit() (uint16, error) {
	return _PermissionlessNodeRegistry.Contract.InputKeyCountLimit(&_PermissionlessNodeRegistry.CallOpts)
}

// IsExistingOperator is a free data retrieval call binding the contract method 0xf9c4dda4.
//
// Solidity: function isExistingOperator(address _operAddr) view returns(bool)
func (_PermissionlessNodeRegistry *PermissionlessNodeRegistryCaller) IsExistingOperator(opts *bind.CallOpts, _operAddr common.Address) (bool, error) {
	var out []interface{}
	err := _PermissionlessNodeRegistry.contract.Call(opts, &out, "isExistingOperator", _operAddr)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsExistingOperator is a free data retrieval call binding the contract method 0xf9c4dda4.
//
// Solidity: function isExistingOperator(address _operAddr) view returns(bool)
func (_PermissionlessNodeRegistry *PermissionlessNodeRegistrySession) IsExistingOperator(_operAddr common.Address) (bool, error) {
	return _PermissionlessNodeRegistry.Contract.IsExistingOperator(&_PermissionlessNodeRegistry.CallOpts, _operAddr)
}

// IsExistingOperator is a free data retrieval call binding the contract method 0xf9c4dda4.
//
// Solidity: function isExistingOperator(address _operAddr) view returns(bool)
func (_PermissionlessNodeRegistry *PermissionlessNodeRegistryCallerSession) IsExistingOperator(_operAddr common.Address) (bool, error) {
	return _PermissionlessNodeRegistry.Contract.IsExistingOperator(&_PermissionlessNodeRegistry.CallOpts, _operAddr)
}

// IsExistingPubkey is a free data retrieval call binding the contract method 0x36514d9f.
//
// Solidity: function isExistingPubkey(bytes _pubkey) view returns(bool)
func (_PermissionlessNodeRegistry *PermissionlessNodeRegistryCaller) IsExistingPubkey(opts *bind.CallOpts, _pubkey []byte) (bool, error) {
	var out []interface{}
	err := _PermissionlessNodeRegistry.contract.Call(opts, &out, "isExistingPubkey", _pubkey)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsExistingPubkey is a free data retrieval call binding the contract method 0x36514d9f.
//
// Solidity: function isExistingPubkey(bytes _pubkey) view returns(bool)
func (_PermissionlessNodeRegistry *PermissionlessNodeRegistrySession) IsExistingPubkey(_pubkey []byte) (bool, error) {
	return _PermissionlessNodeRegistry.Contract.IsExistingPubkey(&_PermissionlessNodeRegistry.CallOpts, _pubkey)
}

// IsExistingPubkey is a free data retrieval call binding the contract method 0x36514d9f.
//
// Solidity: function isExistingPubkey(bytes _pubkey) view returns(bool)
func (_PermissionlessNodeRegistry *PermissionlessNodeRegistryCallerSession) IsExistingPubkey(_pubkey []byte) (bool, error) {
	return _PermissionlessNodeRegistry.Contract.IsExistingPubkey(&_PermissionlessNodeRegistry.CallOpts, _pubkey)
}

// MaxNonTerminalKeyPerOperator is a free data retrieval call binding the contract method 0x50d5d7ab.
//
// Solidity: function maxNonTerminalKeyPerOperator() view returns(uint64)
func (_PermissionlessNodeRegistry *PermissionlessNodeRegistryCaller) MaxNonTerminalKeyPerOperator(opts *bind.CallOpts) (uint64, error) {
	var out []interface{}
	err := _PermissionlessNodeRegistry.contract.Call(opts, &out, "maxNonTerminalKeyPerOperator")

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// MaxNonTerminalKeyPerOperator is a free data retrieval call binding the contract method 0x50d5d7ab.
//
// Solidity: function maxNonTerminalKeyPerOperator() view returns(uint64)
func (_PermissionlessNodeRegistry *PermissionlessNodeRegistrySession) MaxNonTerminalKeyPerOperator() (uint64, error) {
	return _PermissionlessNodeRegistry.Contract.MaxNonTerminalKeyPerOperator(&_PermissionlessNodeRegistry.CallOpts)
}

// MaxNonTerminalKeyPerOperator is a free data retrieval call binding the contract method 0x50d5d7ab.
//
// Solidity: function maxNonTerminalKeyPerOperator() view returns(uint64)
func (_PermissionlessNodeRegistry *PermissionlessNodeRegistryCallerSession) MaxNonTerminalKeyPerOperator() (uint64, error) {
	return _PermissionlessNodeRegistry.Contract.MaxNonTerminalKeyPerOperator(&_PermissionlessNodeRegistry.CallOpts)
}

// NextOperatorId is a free data retrieval call binding the contract method 0x2d1dbd74.
//
// Solidity: function nextOperatorId() view returns(uint256)
func (_PermissionlessNodeRegistry *PermissionlessNodeRegistryCaller) NextOperatorId(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _PermissionlessNodeRegistry.contract.Call(opts, &out, "nextOperatorId")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// NextOperatorId is a free data retrieval call binding the contract method 0x2d1dbd74.
//
// Solidity: function nextOperatorId() view returns(uint256)
func (_PermissionlessNodeRegistry *PermissionlessNodeRegistrySession) NextOperatorId() (*big.Int, error) {
	return _PermissionlessNodeRegistry.Contract.NextOperatorId(&_PermissionlessNodeRegistry.CallOpts)
}

// NextOperatorId is a free data retrieval call binding the contract method 0x2d1dbd74.
//
// Solidity: function nextOperatorId() view returns(uint256)
func (_PermissionlessNodeRegistry *PermissionlessNodeRegistryCallerSession) NextOperatorId() (*big.Int, error) {
	return _PermissionlessNodeRegistry.Contract.NextOperatorId(&_PermissionlessNodeRegistry.CallOpts)
}

// NextQueuedValidatorIndex is a free data retrieval call binding the contract method 0x74338e6d.
//
// Solidity: function nextQueuedValidatorIndex() view returns(uint256)
func (_PermissionlessNodeRegistry *PermissionlessNodeRegistryCaller) NextQueuedValidatorIndex(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _PermissionlessNodeRegistry.contract.Call(opts, &out, "nextQueuedValidatorIndex")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// NextQueuedValidatorIndex is a free data retrieval call binding the contract method 0x74338e6d.
//
// Solidity: function nextQueuedValidatorIndex() view returns(uint256)
func (_PermissionlessNodeRegistry *PermissionlessNodeRegistrySession) NextQueuedValidatorIndex() (*big.Int, error) {
	return _PermissionlessNodeRegistry.Contract.NextQueuedValidatorIndex(&_PermissionlessNodeRegistry.CallOpts)
}

// NextQueuedValidatorIndex is a free data retrieval call binding the contract method 0x74338e6d.
//
// Solidity: function nextQueuedValidatorIndex() view returns(uint256)
func (_PermissionlessNodeRegistry *PermissionlessNodeRegistryCallerSession) NextQueuedValidatorIndex() (*big.Int, error) {
	return _PermissionlessNodeRegistry.Contract.NextQueuedValidatorIndex(&_PermissionlessNodeRegistry.CallOpts)
}

// NextValidatorId is a free data retrieval call binding the contract method 0xf7c09189.
//
// Solidity: function nextValidatorId() view returns(uint256)
func (_PermissionlessNodeRegistry *PermissionlessNodeRegistryCaller) NextValidatorId(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _PermissionlessNodeRegistry.contract.Call(opts, &out, "nextValidatorId")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// NextValidatorId is a free data retrieval call binding the contract method 0xf7c09189.
//
// Solidity: function nextValidatorId() view returns(uint256)
func (_PermissionlessNodeRegistry *PermissionlessNodeRegistrySession) NextValidatorId() (*big.Int, error) {
	return _PermissionlessNodeRegistry.Contract.NextValidatorId(&_PermissionlessNodeRegistry.CallOpts)
}

// NextValidatorId is a free data retrieval call binding the contract method 0xf7c09189.
//
// Solidity: function nextValidatorId() view returns(uint256)
func (_PermissionlessNodeRegistry *PermissionlessNodeRegistryCallerSession) NextValidatorId() (*big.Int, error) {
	return _PermissionlessNodeRegistry.Contract.NextValidatorId(&_PermissionlessNodeRegistry.CallOpts)
}

// NodeELRewardVaultByOperatorId is a free data retrieval call binding the contract method 0x9344b242.
//
// Solidity: function nodeELRewardVaultByOperatorId(uint256 ) view returns(address)
func (_PermissionlessNodeRegistry *PermissionlessNodeRegistryCaller) NodeELRewardVaultByOperatorId(opts *bind.CallOpts, arg0 *big.Int) (common.Address, error) {
	var out []interface{}
	err := _PermissionlessNodeRegistry.contract.Call(opts, &out, "nodeELRewardVaultByOperatorId", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// NodeELRewardVaultByOperatorId is a free data retrieval call binding the contract method 0x9344b242.
//
// Solidity: function nodeELRewardVaultByOperatorId(uint256 ) view returns(address)
func (_PermissionlessNodeRegistry *PermissionlessNodeRegistrySession) NodeELRewardVaultByOperatorId(arg0 *big.Int) (common.Address, error) {
	return _PermissionlessNodeRegistry.Contract.NodeELRewardVaultByOperatorId(&_PermissionlessNodeRegistry.CallOpts, arg0)
}

// NodeELRewardVaultByOperatorId is a free data retrieval call binding the contract method 0x9344b242.
//
// Solidity: function nodeELRewardVaultByOperatorId(uint256 ) view returns(address)
func (_PermissionlessNodeRegistry *PermissionlessNodeRegistryCallerSession) NodeELRewardVaultByOperatorId(arg0 *big.Int) (common.Address, error) {
	return _PermissionlessNodeRegistry.Contract.NodeELRewardVaultByOperatorId(&_PermissionlessNodeRegistry.CallOpts, arg0)
}

// OperatorIDByAddress is a free data retrieval call binding the contract method 0xcac8b306.
//
// Solidity: function operatorIDByAddress(address ) view returns(uint256)
func (_PermissionlessNodeRegistry *PermissionlessNodeRegistryCaller) OperatorIDByAddress(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _PermissionlessNodeRegistry.contract.Call(opts, &out, "operatorIDByAddress", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// OperatorIDByAddress is a free data retrieval call binding the contract method 0xcac8b306.
//
// Solidity: function operatorIDByAddress(address ) view returns(uint256)
func (_PermissionlessNodeRegistry *PermissionlessNodeRegistrySession) OperatorIDByAddress(arg0 common.Address) (*big.Int, error) {
	return _PermissionlessNodeRegistry.Contract.OperatorIDByAddress(&_PermissionlessNodeRegistry.CallOpts, arg0)
}

// OperatorIDByAddress is a free data retrieval call binding the contract method 0xcac8b306.
//
// Solidity: function operatorIDByAddress(address ) view returns(uint256)
func (_PermissionlessNodeRegistry *PermissionlessNodeRegistryCallerSession) OperatorIDByAddress(arg0 common.Address) (*big.Int, error) {
	return _PermissionlessNodeRegistry.Contract.OperatorIDByAddress(&_PermissionlessNodeRegistry.CallOpts, arg0)
}

// OperatorStructById is a free data retrieval call binding the contract method 0xc8a00e7a.
//
// Solidity: function operatorStructById(uint256 ) view returns(bool active, bool optedForSocializingPool, string operatorName, address operatorRewardAddress, address operatorAddress)
func (_PermissionlessNodeRegistry *PermissionlessNodeRegistryCaller) OperatorStructById(opts *bind.CallOpts, arg0 *big.Int) (struct {
	Active                  bool
	OptedForSocializingPool bool
	OperatorName            string
	OperatorRewardAddress   common.Address
	OperatorAddress         common.Address
}, error) {
	var out []interface{}
	err := _PermissionlessNodeRegistry.contract.Call(opts, &out, "operatorStructById", arg0)

	outstruct := new(struct {
		Active                  bool
		OptedForSocializingPool bool
		OperatorName            string
		OperatorRewardAddress   common.Address
		OperatorAddress         common.Address
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Active = *abi.ConvertType(out[0], new(bool)).(*bool)
	outstruct.OptedForSocializingPool = *abi.ConvertType(out[1], new(bool)).(*bool)
	outstruct.OperatorName = *abi.ConvertType(out[2], new(string)).(*string)
	outstruct.OperatorRewardAddress = *abi.ConvertType(out[3], new(common.Address)).(*common.Address)
	outstruct.OperatorAddress = *abi.ConvertType(out[4], new(common.Address)).(*common.Address)

	return *outstruct, err

}

// OperatorStructById is a free data retrieval call binding the contract method 0xc8a00e7a.
//
// Solidity: function operatorStructById(uint256 ) view returns(bool active, bool optedForSocializingPool, string operatorName, address operatorRewardAddress, address operatorAddress)
func (_PermissionlessNodeRegistry *PermissionlessNodeRegistrySession) OperatorStructById(arg0 *big.Int) (struct {
	Active                  bool
	OptedForSocializingPool bool
	OperatorName            string
	OperatorRewardAddress   common.Address
	OperatorAddress         common.Address
}, error) {
	return _PermissionlessNodeRegistry.Contract.OperatorStructById(&_PermissionlessNodeRegistry.CallOpts, arg0)
}

// OperatorStructById is a free data retrieval call binding the contract method 0xc8a00e7a.
//
// Solidity: function operatorStructById(uint256 ) view returns(bool active, bool optedForSocializingPool, string operatorName, address operatorRewardAddress, address operatorAddress)
func (_PermissionlessNodeRegistry *PermissionlessNodeRegistryCallerSession) OperatorStructById(arg0 *big.Int) (struct {
	Active                  bool
	OptedForSocializingPool bool
	OperatorName            string
	OperatorRewardAddress   common.Address
	OperatorAddress         common.Address
}, error) {
	return _PermissionlessNodeRegistry.Contract.OperatorStructById(&_PermissionlessNodeRegistry.CallOpts, arg0)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_PermissionlessNodeRegistry *PermissionlessNodeRegistryCaller) Paused(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _PermissionlessNodeRegistry.contract.Call(opts, &out, "paused")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_PermissionlessNodeRegistry *PermissionlessNodeRegistrySession) Paused() (bool, error) {
	return _PermissionlessNodeRegistry.Contract.Paused(&_PermissionlessNodeRegistry.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_PermissionlessNodeRegistry *PermissionlessNodeRegistryCallerSession) Paused() (bool, error) {
	return _PermissionlessNodeRegistry.Contract.Paused(&_PermissionlessNodeRegistry.CallOpts)
}

// QueuedValidators is a free data retrieval call binding the contract method 0xbc4a3ad5.
//
// Solidity: function queuedValidators(uint256 ) view returns(uint256)
func (_PermissionlessNodeRegistry *PermissionlessNodeRegistryCaller) QueuedValidators(opts *bind.CallOpts, arg0 *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _PermissionlessNodeRegistry.contract.Call(opts, &out, "queuedValidators", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// QueuedValidators is a free data retrieval call binding the contract method 0xbc4a3ad5.
//
// Solidity: function queuedValidators(uint256 ) view returns(uint256)
func (_PermissionlessNodeRegistry *PermissionlessNodeRegistrySession) QueuedValidators(arg0 *big.Int) (*big.Int, error) {
	return _PermissionlessNodeRegistry.Contract.QueuedValidators(&_PermissionlessNodeRegistry.CallOpts, arg0)
}

// QueuedValidators is a free data retrieval call binding the contract method 0xbc4a3ad5.
//
// Solidity: function queuedValidators(uint256 ) view returns(uint256)
func (_PermissionlessNodeRegistry *PermissionlessNodeRegistryCallerSession) QueuedValidators(arg0 *big.Int) (*big.Int, error) {
	return _PermissionlessNodeRegistry.Contract.QueuedValidators(&_PermissionlessNodeRegistry.CallOpts, arg0)
}

// SocializingPoolStateChangeBlock is a free data retrieval call binding the contract method 0x84522a6d.
//
// Solidity: function socializingPoolStateChangeBlock(uint256 ) view returns(uint256)
func (_PermissionlessNodeRegistry *PermissionlessNodeRegistryCaller) SocializingPoolStateChangeBlock(opts *bind.CallOpts, arg0 *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _PermissionlessNodeRegistry.contract.Call(opts, &out, "socializingPoolStateChangeBlock", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// SocializingPoolStateChangeBlock is a free data retrieval call binding the contract method 0x84522a6d.
//
// Solidity: function socializingPoolStateChangeBlock(uint256 ) view returns(uint256)
func (_PermissionlessNodeRegistry *PermissionlessNodeRegistrySession) SocializingPoolStateChangeBlock(arg0 *big.Int) (*big.Int, error) {
	return _PermissionlessNodeRegistry.Contract.SocializingPoolStateChangeBlock(&_PermissionlessNodeRegistry.CallOpts, arg0)
}

// SocializingPoolStateChangeBlock is a free data retrieval call binding the contract method 0x84522a6d.
//
// Solidity: function socializingPoolStateChangeBlock(uint256 ) view returns(uint256)
func (_PermissionlessNodeRegistry *PermissionlessNodeRegistryCallerSession) SocializingPoolStateChangeBlock(arg0 *big.Int) (*big.Int, error) {
	return _PermissionlessNodeRegistry.Contract.SocializingPoolStateChangeBlock(&_PermissionlessNodeRegistry.CallOpts, arg0)
}

// StaderConfig is a free data retrieval call binding the contract method 0x490ffa35.
//
// Solidity: function staderConfig() view returns(address)
func (_PermissionlessNodeRegistry *PermissionlessNodeRegistryCaller) StaderConfig(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _PermissionlessNodeRegistry.contract.Call(opts, &out, "staderConfig")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// StaderConfig is a free data retrieval call binding the contract method 0x490ffa35.
//
// Solidity: function staderConfig() view returns(address)
func (_PermissionlessNodeRegistry *PermissionlessNodeRegistrySession) StaderConfig() (common.Address, error) {
	return _PermissionlessNodeRegistry.Contract.StaderConfig(&_PermissionlessNodeRegistry.CallOpts)
}

// StaderConfig is a free data retrieval call binding the contract method 0x490ffa35.
//
// Solidity: function staderConfig() view returns(address)
func (_PermissionlessNodeRegistry *PermissionlessNodeRegistryCallerSession) StaderConfig() (common.Address, error) {
	return _PermissionlessNodeRegistry.Contract.StaderConfig(&_PermissionlessNodeRegistry.CallOpts)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_PermissionlessNodeRegistry *PermissionlessNodeRegistryCaller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _PermissionlessNodeRegistry.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_PermissionlessNodeRegistry *PermissionlessNodeRegistrySession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _PermissionlessNodeRegistry.Contract.SupportsInterface(&_PermissionlessNodeRegistry.CallOpts, interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_PermissionlessNodeRegistry *PermissionlessNodeRegistryCallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _PermissionlessNodeRegistry.Contract.SupportsInterface(&_PermissionlessNodeRegistry.CallOpts, interfaceId)
}

// TotalActiveValidatorCount is a free data retrieval call binding the contract method 0x84b0fa4c.
//
// Solidity: function totalActiveValidatorCount() view returns(uint256)
func (_PermissionlessNodeRegistry *PermissionlessNodeRegistryCaller) TotalActiveValidatorCount(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _PermissionlessNodeRegistry.contract.Call(opts, &out, "totalActiveValidatorCount")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalActiveValidatorCount is a free data retrieval call binding the contract method 0x84b0fa4c.
//
// Solidity: function totalActiveValidatorCount() view returns(uint256)
func (_PermissionlessNodeRegistry *PermissionlessNodeRegistrySession) TotalActiveValidatorCount() (*big.Int, error) {
	return _PermissionlessNodeRegistry.Contract.TotalActiveValidatorCount(&_PermissionlessNodeRegistry.CallOpts)
}

// TotalActiveValidatorCount is a free data retrieval call binding the contract method 0x84b0fa4c.
//
// Solidity: function totalActiveValidatorCount() view returns(uint256)
func (_PermissionlessNodeRegistry *PermissionlessNodeRegistryCallerSession) TotalActiveValidatorCount() (*big.Int, error) {
	return _PermissionlessNodeRegistry.Contract.TotalActiveValidatorCount(&_PermissionlessNodeRegistry.CallOpts)
}

// ValidatorIdByPubkey is a free data retrieval call binding the contract method 0x5c2c30a5.
//
// Solidity: function validatorIdByPubkey(bytes ) view returns(uint256)
func (_PermissionlessNodeRegistry *PermissionlessNodeRegistryCaller) ValidatorIdByPubkey(opts *bind.CallOpts, arg0 []byte) (*big.Int, error) {
	var out []interface{}
	err := _PermissionlessNodeRegistry.contract.Call(opts, &out, "validatorIdByPubkey", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ValidatorIdByPubkey is a free data retrieval call binding the contract method 0x5c2c30a5.
//
// Solidity: function validatorIdByPubkey(bytes ) view returns(uint256)
func (_PermissionlessNodeRegistry *PermissionlessNodeRegistrySession) ValidatorIdByPubkey(arg0 []byte) (*big.Int, error) {
	return _PermissionlessNodeRegistry.Contract.ValidatorIdByPubkey(&_PermissionlessNodeRegistry.CallOpts, arg0)
}

// ValidatorIdByPubkey is a free data retrieval call binding the contract method 0x5c2c30a5.
//
// Solidity: function validatorIdByPubkey(bytes ) view returns(uint256)
func (_PermissionlessNodeRegistry *PermissionlessNodeRegistryCallerSession) ValidatorIdByPubkey(arg0 []byte) (*big.Int, error) {
	return _PermissionlessNodeRegistry.Contract.ValidatorIdByPubkey(&_PermissionlessNodeRegistry.CallOpts, arg0)
}

// ValidatorIdsByOperatorId is a free data retrieval call binding the contract method 0xd5e1e5ce.
//
// Solidity: function validatorIdsByOperatorId(uint256 , uint256 ) view returns(uint256)
func (_PermissionlessNodeRegistry *PermissionlessNodeRegistryCaller) ValidatorIdsByOperatorId(opts *bind.CallOpts, arg0 *big.Int, arg1 *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _PermissionlessNodeRegistry.contract.Call(opts, &out, "validatorIdsByOperatorId", arg0, arg1)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ValidatorIdsByOperatorId is a free data retrieval call binding the contract method 0xd5e1e5ce.
//
// Solidity: function validatorIdsByOperatorId(uint256 , uint256 ) view returns(uint256)
func (_PermissionlessNodeRegistry *PermissionlessNodeRegistrySession) ValidatorIdsByOperatorId(arg0 *big.Int, arg1 *big.Int) (*big.Int, error) {
	return _PermissionlessNodeRegistry.Contract.ValidatorIdsByOperatorId(&_PermissionlessNodeRegistry.CallOpts, arg0, arg1)
}

// ValidatorIdsByOperatorId is a free data retrieval call binding the contract method 0xd5e1e5ce.
//
// Solidity: function validatorIdsByOperatorId(uint256 , uint256 ) view returns(uint256)
func (_PermissionlessNodeRegistry *PermissionlessNodeRegistryCallerSession) ValidatorIdsByOperatorId(arg0 *big.Int, arg1 *big.Int) (*big.Int, error) {
	return _PermissionlessNodeRegistry.Contract.ValidatorIdsByOperatorId(&_PermissionlessNodeRegistry.CallOpts, arg0, arg1)
}

// ValidatorQueueSize is a free data retrieval call binding the contract method 0x49911bfb.
//
// Solidity: function validatorQueueSize() view returns(uint256)
func (_PermissionlessNodeRegistry *PermissionlessNodeRegistryCaller) ValidatorQueueSize(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _PermissionlessNodeRegistry.contract.Call(opts, &out, "validatorQueueSize")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ValidatorQueueSize is a free data retrieval call binding the contract method 0x49911bfb.
//
// Solidity: function validatorQueueSize() view returns(uint256)
func (_PermissionlessNodeRegistry *PermissionlessNodeRegistrySession) ValidatorQueueSize() (*big.Int, error) {
	return _PermissionlessNodeRegistry.Contract.ValidatorQueueSize(&_PermissionlessNodeRegistry.CallOpts)
}

// ValidatorQueueSize is a free data retrieval call binding the contract method 0x49911bfb.
//
// Solidity: function validatorQueueSize() view returns(uint256)
func (_PermissionlessNodeRegistry *PermissionlessNodeRegistryCallerSession) ValidatorQueueSize() (*big.Int, error) {
	return _PermissionlessNodeRegistry.Contract.ValidatorQueueSize(&_PermissionlessNodeRegistry.CallOpts)
}

// ValidatorRegistry is a free data retrieval call binding the contract method 0x5a1239c1.
//
// Solidity: function validatorRegistry(uint256 ) view returns(uint8 status, bytes pubkey, bytes preDepositSignature, bytes depositSignature, address withdrawVaultAddress, uint256 operatorId, uint256 depositBlock, uint256 withdrawnBlock)
func (_PermissionlessNodeRegistry *PermissionlessNodeRegistryCaller) ValidatorRegistry(opts *bind.CallOpts, arg0 *big.Int) (struct {
	Status               uint8
	Pubkey               []byte
	PreDepositSignature  []byte
	DepositSignature     []byte
	WithdrawVaultAddress common.Address
	OperatorId           *big.Int
	DepositBlock         *big.Int
	WithdrawnBlock       *big.Int
}, error) {
	var out []interface{}
	err := _PermissionlessNodeRegistry.contract.Call(opts, &out, "validatorRegistry", arg0)

	outstruct := new(struct {
		Status               uint8
		Pubkey               []byte
		PreDepositSignature  []byte
		DepositSignature     []byte
		WithdrawVaultAddress common.Address
		OperatorId           *big.Int
		DepositBlock         *big.Int
		WithdrawnBlock       *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Status = *abi.ConvertType(out[0], new(uint8)).(*uint8)
	outstruct.Pubkey = *abi.ConvertType(out[1], new([]byte)).(*[]byte)
	outstruct.PreDepositSignature = *abi.ConvertType(out[2], new([]byte)).(*[]byte)
	outstruct.DepositSignature = *abi.ConvertType(out[3], new([]byte)).(*[]byte)
	outstruct.WithdrawVaultAddress = *abi.ConvertType(out[4], new(common.Address)).(*common.Address)
	outstruct.OperatorId = *abi.ConvertType(out[5], new(*big.Int)).(**big.Int)
	outstruct.DepositBlock = *abi.ConvertType(out[6], new(*big.Int)).(**big.Int)
	outstruct.WithdrawnBlock = *abi.ConvertType(out[7], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// ValidatorRegistry is a free data retrieval call binding the contract method 0x5a1239c1.
//
// Solidity: function validatorRegistry(uint256 ) view returns(uint8 status, bytes pubkey, bytes preDepositSignature, bytes depositSignature, address withdrawVaultAddress, uint256 operatorId, uint256 depositBlock, uint256 withdrawnBlock)
func (_PermissionlessNodeRegistry *PermissionlessNodeRegistrySession) ValidatorRegistry(arg0 *big.Int) (struct {
	Status               uint8
	Pubkey               []byte
	PreDepositSignature  []byte
	DepositSignature     []byte
	WithdrawVaultAddress common.Address
	OperatorId           *big.Int
	DepositBlock         *big.Int
	WithdrawnBlock       *big.Int
}, error) {
	return _PermissionlessNodeRegistry.Contract.ValidatorRegistry(&_PermissionlessNodeRegistry.CallOpts, arg0)
}

// ValidatorRegistry is a free data retrieval call binding the contract method 0x5a1239c1.
//
// Solidity: function validatorRegistry(uint256 ) view returns(uint8 status, bytes pubkey, bytes preDepositSignature, bytes depositSignature, address withdrawVaultAddress, uint256 operatorId, uint256 depositBlock, uint256 withdrawnBlock)
func (_PermissionlessNodeRegistry *PermissionlessNodeRegistryCallerSession) ValidatorRegistry(arg0 *big.Int) (struct {
	Status               uint8
	Pubkey               []byte
	PreDepositSignature  []byte
	DepositSignature     []byte
	WithdrawVaultAddress common.Address
	OperatorId           *big.Int
	DepositBlock         *big.Int
	WithdrawnBlock       *big.Int
}, error) {
	return _PermissionlessNodeRegistry.Contract.ValidatorRegistry(&_PermissionlessNodeRegistry.CallOpts, arg0)
}

// VerifiedKeyBatchSize is a free data retrieval call binding the contract method 0xab3e71eb.
//
// Solidity: function verifiedKeyBatchSize() view returns(uint256)
func (_PermissionlessNodeRegistry *PermissionlessNodeRegistryCaller) VerifiedKeyBatchSize(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _PermissionlessNodeRegistry.contract.Call(opts, &out, "verifiedKeyBatchSize")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// VerifiedKeyBatchSize is a free data retrieval call binding the contract method 0xab3e71eb.
//
// Solidity: function verifiedKeyBatchSize() view returns(uint256)
func (_PermissionlessNodeRegistry *PermissionlessNodeRegistrySession) VerifiedKeyBatchSize() (*big.Int, error) {
	return _PermissionlessNodeRegistry.Contract.VerifiedKeyBatchSize(&_PermissionlessNodeRegistry.CallOpts)
}

// VerifiedKeyBatchSize is a free data retrieval call binding the contract method 0xab3e71eb.
//
// Solidity: function verifiedKeyBatchSize() view returns(uint256)
func (_PermissionlessNodeRegistry *PermissionlessNodeRegistryCallerSession) VerifiedKeyBatchSize() (*big.Int, error) {
	return _PermissionlessNodeRegistry.Contract.VerifiedKeyBatchSize(&_PermissionlessNodeRegistry.CallOpts)
}

// AddValidatorKeys is a paid mutator transaction binding the contract method 0xdeacde2b.
//
// Solidity: function addValidatorKeys(bytes[] _pubkey, bytes[] _preDepositSignature, bytes[] _depositSignature) payable returns()
func (_PermissionlessNodeRegistry *PermissionlessNodeRegistryTransactor) AddValidatorKeys(opts *bind.TransactOpts, _pubkey [][]byte, _preDepositSignature [][]byte, _depositSignature [][]byte) (*types.Transaction, error) {
	return _PermissionlessNodeRegistry.contract.Transact(opts, "addValidatorKeys", _pubkey, _preDepositSignature, _depositSignature)
}

// AddValidatorKeys is a paid mutator transaction binding the contract method 0xdeacde2b.
//
// Solidity: function addValidatorKeys(bytes[] _pubkey, bytes[] _preDepositSignature, bytes[] _depositSignature) payable returns()
func (_PermissionlessNodeRegistry *PermissionlessNodeRegistrySession) AddValidatorKeys(_pubkey [][]byte, _preDepositSignature [][]byte, _depositSignature [][]byte) (*types.Transaction, error) {
	return _PermissionlessNodeRegistry.Contract.AddValidatorKeys(&_PermissionlessNodeRegistry.TransactOpts, _pubkey, _preDepositSignature, _depositSignature)
}

// AddValidatorKeys is a paid mutator transaction binding the contract method 0xdeacde2b.
//
// Solidity: function addValidatorKeys(bytes[] _pubkey, bytes[] _preDepositSignature, bytes[] _depositSignature) payable returns()
func (_PermissionlessNodeRegistry *PermissionlessNodeRegistryTransactorSession) AddValidatorKeys(_pubkey [][]byte, _preDepositSignature [][]byte, _depositSignature [][]byte) (*types.Transaction, error) {
	return _PermissionlessNodeRegistry.Contract.AddValidatorKeys(&_PermissionlessNodeRegistry.TransactOpts, _pubkey, _preDepositSignature, _depositSignature)
}

// ChangeSocializingPoolState is a paid mutator transaction binding the contract method 0xf90b0838.
//
// Solidity: function changeSocializingPoolState(bool _optInForSocializingPool) returns(address feeRecipientAddress)
func (_PermissionlessNodeRegistry *PermissionlessNodeRegistryTransactor) ChangeSocializingPoolState(opts *bind.TransactOpts, _optInForSocializingPool bool) (*types.Transaction, error) {
	return _PermissionlessNodeRegistry.contract.Transact(opts, "changeSocializingPoolState", _optInForSocializingPool)
}

// ChangeSocializingPoolState is a paid mutator transaction binding the contract method 0xf90b0838.
//
// Solidity: function changeSocializingPoolState(bool _optInForSocializingPool) returns(address feeRecipientAddress)
func (_PermissionlessNodeRegistry *PermissionlessNodeRegistrySession) ChangeSocializingPoolState(_optInForSocializingPool bool) (*types.Transaction, error) {
	return _PermissionlessNodeRegistry.Contract.ChangeSocializingPoolState(&_PermissionlessNodeRegistry.TransactOpts, _optInForSocializingPool)
}

// ChangeSocializingPoolState is a paid mutator transaction binding the contract method 0xf90b0838.
//
// Solidity: function changeSocializingPoolState(bool _optInForSocializingPool) returns(address feeRecipientAddress)
func (_PermissionlessNodeRegistry *PermissionlessNodeRegistryTransactorSession) ChangeSocializingPoolState(_optInForSocializingPool bool) (*types.Transaction, error) {
	return _PermissionlessNodeRegistry.Contract.ChangeSocializingPoolState(&_PermissionlessNodeRegistry.TransactOpts, _optInForSocializingPool)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_PermissionlessNodeRegistry *PermissionlessNodeRegistryTransactor) GrantRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _PermissionlessNodeRegistry.contract.Transact(opts, "grantRole", role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_PermissionlessNodeRegistry *PermissionlessNodeRegistrySession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _PermissionlessNodeRegistry.Contract.GrantRole(&_PermissionlessNodeRegistry.TransactOpts, role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_PermissionlessNodeRegistry *PermissionlessNodeRegistryTransactorSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _PermissionlessNodeRegistry.Contract.GrantRole(&_PermissionlessNodeRegistry.TransactOpts, role, account)
}

// IncreaseTotalActiveValidatorCount is a paid mutator transaction binding the contract method 0x59c3c9b7.
//
// Solidity: function increaseTotalActiveValidatorCount(uint256 _count) returns()
func (_PermissionlessNodeRegistry *PermissionlessNodeRegistryTransactor) IncreaseTotalActiveValidatorCount(opts *bind.TransactOpts, _count *big.Int) (*types.Transaction, error) {
	return _PermissionlessNodeRegistry.contract.Transact(opts, "increaseTotalActiveValidatorCount", _count)
}

// IncreaseTotalActiveValidatorCount is a paid mutator transaction binding the contract method 0x59c3c9b7.
//
// Solidity: function increaseTotalActiveValidatorCount(uint256 _count) returns()
func (_PermissionlessNodeRegistry *PermissionlessNodeRegistrySession) IncreaseTotalActiveValidatorCount(_count *big.Int) (*types.Transaction, error) {
	return _PermissionlessNodeRegistry.Contract.IncreaseTotalActiveValidatorCount(&_PermissionlessNodeRegistry.TransactOpts, _count)
}

// IncreaseTotalActiveValidatorCount is a paid mutator transaction binding the contract method 0x59c3c9b7.
//
// Solidity: function increaseTotalActiveValidatorCount(uint256 _count) returns()
func (_PermissionlessNodeRegistry *PermissionlessNodeRegistryTransactorSession) IncreaseTotalActiveValidatorCount(_count *big.Int) (*types.Transaction, error) {
	return _PermissionlessNodeRegistry.Contract.IncreaseTotalActiveValidatorCount(&_PermissionlessNodeRegistry.TransactOpts, _count)
}

// MarkValidatorReadyToDeposit is a paid mutator transaction binding the contract method 0x13797bff.
//
// Solidity: function markValidatorReadyToDeposit(bytes[] _readyToDepositPubkey, bytes[] _frontRunPubkey, bytes[] _invalidSignaturePubkey) returns()
func (_PermissionlessNodeRegistry *PermissionlessNodeRegistryTransactor) MarkValidatorReadyToDeposit(opts *bind.TransactOpts, _readyToDepositPubkey [][]byte, _frontRunPubkey [][]byte, _invalidSignaturePubkey [][]byte) (*types.Transaction, error) {
	return _PermissionlessNodeRegistry.contract.Transact(opts, "markValidatorReadyToDeposit", _readyToDepositPubkey, _frontRunPubkey, _invalidSignaturePubkey)
}

// MarkValidatorReadyToDeposit is a paid mutator transaction binding the contract method 0x13797bff.
//
// Solidity: function markValidatorReadyToDeposit(bytes[] _readyToDepositPubkey, bytes[] _frontRunPubkey, bytes[] _invalidSignaturePubkey) returns()
func (_PermissionlessNodeRegistry *PermissionlessNodeRegistrySession) MarkValidatorReadyToDeposit(_readyToDepositPubkey [][]byte, _frontRunPubkey [][]byte, _invalidSignaturePubkey [][]byte) (*types.Transaction, error) {
	return _PermissionlessNodeRegistry.Contract.MarkValidatorReadyToDeposit(&_PermissionlessNodeRegistry.TransactOpts, _readyToDepositPubkey, _frontRunPubkey, _invalidSignaturePubkey)
}

// MarkValidatorReadyToDeposit is a paid mutator transaction binding the contract method 0x13797bff.
//
// Solidity: function markValidatorReadyToDeposit(bytes[] _readyToDepositPubkey, bytes[] _frontRunPubkey, bytes[] _invalidSignaturePubkey) returns()
func (_PermissionlessNodeRegistry *PermissionlessNodeRegistryTransactorSession) MarkValidatorReadyToDeposit(_readyToDepositPubkey [][]byte, _frontRunPubkey [][]byte, _invalidSignaturePubkey [][]byte) (*types.Transaction, error) {
	return _PermissionlessNodeRegistry.Contract.MarkValidatorReadyToDeposit(&_PermissionlessNodeRegistry.TransactOpts, _readyToDepositPubkey, _frontRunPubkey, _invalidSignaturePubkey)
}

// OnboardNodeOperator is a paid mutator transaction binding the contract method 0x044d2fe8.
//
// Solidity: function onboardNodeOperator(bool _optInForSocializingPool, string _operatorName, address _operatorRewardAddress) returns(address feeRecipientAddress)
func (_PermissionlessNodeRegistry *PermissionlessNodeRegistryTransactor) OnboardNodeOperator(opts *bind.TransactOpts, _optInForSocializingPool bool, _operatorName string, _operatorRewardAddress common.Address) (*types.Transaction, error) {
	return _PermissionlessNodeRegistry.contract.Transact(opts, "onboardNodeOperator", _optInForSocializingPool, _operatorName, _operatorRewardAddress)
}

// OnboardNodeOperator is a paid mutator transaction binding the contract method 0x044d2fe8.
//
// Solidity: function onboardNodeOperator(bool _optInForSocializingPool, string _operatorName, address _operatorRewardAddress) returns(address feeRecipientAddress)
func (_PermissionlessNodeRegistry *PermissionlessNodeRegistrySession) OnboardNodeOperator(_optInForSocializingPool bool, _operatorName string, _operatorRewardAddress common.Address) (*types.Transaction, error) {
	return _PermissionlessNodeRegistry.Contract.OnboardNodeOperator(&_PermissionlessNodeRegistry.TransactOpts, _optInForSocializingPool, _operatorName, _operatorRewardAddress)
}

// OnboardNodeOperator is a paid mutator transaction binding the contract method 0x044d2fe8.
//
// Solidity: function onboardNodeOperator(bool _optInForSocializingPool, string _operatorName, address _operatorRewardAddress) returns(address feeRecipientAddress)
func (_PermissionlessNodeRegistry *PermissionlessNodeRegistryTransactorSession) OnboardNodeOperator(_optInForSocializingPool bool, _operatorName string, _operatorRewardAddress common.Address) (*types.Transaction, error) {
	return _PermissionlessNodeRegistry.Contract.OnboardNodeOperator(&_PermissionlessNodeRegistry.TransactOpts, _optInForSocializingPool, _operatorName, _operatorRewardAddress)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_PermissionlessNodeRegistry *PermissionlessNodeRegistryTransactor) Pause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PermissionlessNodeRegistry.contract.Transact(opts, "pause")
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_PermissionlessNodeRegistry *PermissionlessNodeRegistrySession) Pause() (*types.Transaction, error) {
	return _PermissionlessNodeRegistry.Contract.Pause(&_PermissionlessNodeRegistry.TransactOpts)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_PermissionlessNodeRegistry *PermissionlessNodeRegistryTransactorSession) Pause() (*types.Transaction, error) {
	return _PermissionlessNodeRegistry.Contract.Pause(&_PermissionlessNodeRegistry.TransactOpts)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address account) returns()
func (_PermissionlessNodeRegistry *PermissionlessNodeRegistryTransactor) RenounceRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _PermissionlessNodeRegistry.contract.Transact(opts, "renounceRole", role, account)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address account) returns()
func (_PermissionlessNodeRegistry *PermissionlessNodeRegistrySession) RenounceRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _PermissionlessNodeRegistry.Contract.RenounceRole(&_PermissionlessNodeRegistry.TransactOpts, role, account)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address account) returns()
func (_PermissionlessNodeRegistry *PermissionlessNodeRegistryTransactorSession) RenounceRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _PermissionlessNodeRegistry.Contract.RenounceRole(&_PermissionlessNodeRegistry.TransactOpts, role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_PermissionlessNodeRegistry *PermissionlessNodeRegistryTransactor) RevokeRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _PermissionlessNodeRegistry.contract.Transact(opts, "revokeRole", role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_PermissionlessNodeRegistry *PermissionlessNodeRegistrySession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _PermissionlessNodeRegistry.Contract.RevokeRole(&_PermissionlessNodeRegistry.TransactOpts, role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_PermissionlessNodeRegistry *PermissionlessNodeRegistryTransactorSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _PermissionlessNodeRegistry.Contract.RevokeRole(&_PermissionlessNodeRegistry.TransactOpts, role, account)
}

// TransferCollateralToPool is a paid mutator transaction binding the contract method 0x5ae7f25d.
//
// Solidity: function transferCollateralToPool(uint256 _amount) returns()
func (_PermissionlessNodeRegistry *PermissionlessNodeRegistryTransactor) TransferCollateralToPool(opts *bind.TransactOpts, _amount *big.Int) (*types.Transaction, error) {
	return _PermissionlessNodeRegistry.contract.Transact(opts, "transferCollateralToPool", _amount)
}

// TransferCollateralToPool is a paid mutator transaction binding the contract method 0x5ae7f25d.
//
// Solidity: function transferCollateralToPool(uint256 _amount) returns()
func (_PermissionlessNodeRegistry *PermissionlessNodeRegistrySession) TransferCollateralToPool(_amount *big.Int) (*types.Transaction, error) {
	return _PermissionlessNodeRegistry.Contract.TransferCollateralToPool(&_PermissionlessNodeRegistry.TransactOpts, _amount)
}

// TransferCollateralToPool is a paid mutator transaction binding the contract method 0x5ae7f25d.
//
// Solidity: function transferCollateralToPool(uint256 _amount) returns()
func (_PermissionlessNodeRegistry *PermissionlessNodeRegistryTransactorSession) TransferCollateralToPool(_amount *big.Int) (*types.Transaction, error) {
	return _PermissionlessNodeRegistry.Contract.TransferCollateralToPool(&_PermissionlessNodeRegistry.TransactOpts, _amount)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_PermissionlessNodeRegistry *PermissionlessNodeRegistryTransactor) Unpause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PermissionlessNodeRegistry.contract.Transact(opts, "unpause")
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_PermissionlessNodeRegistry *PermissionlessNodeRegistrySession) Unpause() (*types.Transaction, error) {
	return _PermissionlessNodeRegistry.Contract.Unpause(&_PermissionlessNodeRegistry.TransactOpts)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_PermissionlessNodeRegistry *PermissionlessNodeRegistryTransactorSession) Unpause() (*types.Transaction, error) {
	return _PermissionlessNodeRegistry.Contract.Unpause(&_PermissionlessNodeRegistry.TransactOpts)
}

// UpdateDepositStatusAndBlock is a paid mutator transaction binding the contract method 0x186d9541.
//
// Solidity: function updateDepositStatusAndBlock(uint256 _validatorId) returns()
func (_PermissionlessNodeRegistry *PermissionlessNodeRegistryTransactor) UpdateDepositStatusAndBlock(opts *bind.TransactOpts, _validatorId *big.Int) (*types.Transaction, error) {
	return _PermissionlessNodeRegistry.contract.Transact(opts, "updateDepositStatusAndBlock", _validatorId)
}

// UpdateDepositStatusAndBlock is a paid mutator transaction binding the contract method 0x186d9541.
//
// Solidity: function updateDepositStatusAndBlock(uint256 _validatorId) returns()
func (_PermissionlessNodeRegistry *PermissionlessNodeRegistrySession) UpdateDepositStatusAndBlock(_validatorId *big.Int) (*types.Transaction, error) {
	return _PermissionlessNodeRegistry.Contract.UpdateDepositStatusAndBlock(&_PermissionlessNodeRegistry.TransactOpts, _validatorId)
}

// UpdateDepositStatusAndBlock is a paid mutator transaction binding the contract method 0x186d9541.
//
// Solidity: function updateDepositStatusAndBlock(uint256 _validatorId) returns()
func (_PermissionlessNodeRegistry *PermissionlessNodeRegistryTransactorSession) UpdateDepositStatusAndBlock(_validatorId *big.Int) (*types.Transaction, error) {
	return _PermissionlessNodeRegistry.Contract.UpdateDepositStatusAndBlock(&_PermissionlessNodeRegistry.TransactOpts, _validatorId)
}

// UpdateInputKeyCountLimit is a paid mutator transaction binding the contract method 0x2517cfbf.
//
// Solidity: function updateInputKeyCountLimit(uint16 _inputKeyCountLimit) returns()
func (_PermissionlessNodeRegistry *PermissionlessNodeRegistryTransactor) UpdateInputKeyCountLimit(opts *bind.TransactOpts, _inputKeyCountLimit uint16) (*types.Transaction, error) {
	return _PermissionlessNodeRegistry.contract.Transact(opts, "updateInputKeyCountLimit", _inputKeyCountLimit)
}

// UpdateInputKeyCountLimit is a paid mutator transaction binding the contract method 0x2517cfbf.
//
// Solidity: function updateInputKeyCountLimit(uint16 _inputKeyCountLimit) returns()
func (_PermissionlessNodeRegistry *PermissionlessNodeRegistrySession) UpdateInputKeyCountLimit(_inputKeyCountLimit uint16) (*types.Transaction, error) {
	return _PermissionlessNodeRegistry.Contract.UpdateInputKeyCountLimit(&_PermissionlessNodeRegistry.TransactOpts, _inputKeyCountLimit)
}

// UpdateInputKeyCountLimit is a paid mutator transaction binding the contract method 0x2517cfbf.
//
// Solidity: function updateInputKeyCountLimit(uint16 _inputKeyCountLimit) returns()
func (_PermissionlessNodeRegistry *PermissionlessNodeRegistryTransactorSession) UpdateInputKeyCountLimit(_inputKeyCountLimit uint16) (*types.Transaction, error) {
	return _PermissionlessNodeRegistry.Contract.UpdateInputKeyCountLimit(&_PermissionlessNodeRegistry.TransactOpts, _inputKeyCountLimit)
}

// UpdateMaxNonTerminalKeyPerOperator is a paid mutator transaction binding the contract method 0x60c3cf3f.
//
// Solidity: function updateMaxNonTerminalKeyPerOperator(uint64 _maxNonTerminalKeyPerOperator) returns()
func (_PermissionlessNodeRegistry *PermissionlessNodeRegistryTransactor) UpdateMaxNonTerminalKeyPerOperator(opts *bind.TransactOpts, _maxNonTerminalKeyPerOperator uint64) (*types.Transaction, error) {
	return _PermissionlessNodeRegistry.contract.Transact(opts, "updateMaxNonTerminalKeyPerOperator", _maxNonTerminalKeyPerOperator)
}

// UpdateMaxNonTerminalKeyPerOperator is a paid mutator transaction binding the contract method 0x60c3cf3f.
//
// Solidity: function updateMaxNonTerminalKeyPerOperator(uint64 _maxNonTerminalKeyPerOperator) returns()
func (_PermissionlessNodeRegistry *PermissionlessNodeRegistrySession) UpdateMaxNonTerminalKeyPerOperator(_maxNonTerminalKeyPerOperator uint64) (*types.Transaction, error) {
	return _PermissionlessNodeRegistry.Contract.UpdateMaxNonTerminalKeyPerOperator(&_PermissionlessNodeRegistry.TransactOpts, _maxNonTerminalKeyPerOperator)
}

// UpdateMaxNonTerminalKeyPerOperator is a paid mutator transaction binding the contract method 0x60c3cf3f.
//
// Solidity: function updateMaxNonTerminalKeyPerOperator(uint64 _maxNonTerminalKeyPerOperator) returns()
func (_PermissionlessNodeRegistry *PermissionlessNodeRegistryTransactorSession) UpdateMaxNonTerminalKeyPerOperator(_maxNonTerminalKeyPerOperator uint64) (*types.Transaction, error) {
	return _PermissionlessNodeRegistry.Contract.UpdateMaxNonTerminalKeyPerOperator(&_PermissionlessNodeRegistry.TransactOpts, _maxNonTerminalKeyPerOperator)
}

// UpdateNextQueuedValidatorIndex is a paid mutator transaction binding the contract method 0xb8d2f06c.
//
// Solidity: function updateNextQueuedValidatorIndex(uint256 _nextQueuedValidatorIndex) returns()
func (_PermissionlessNodeRegistry *PermissionlessNodeRegistryTransactor) UpdateNextQueuedValidatorIndex(opts *bind.TransactOpts, _nextQueuedValidatorIndex *big.Int) (*types.Transaction, error) {
	return _PermissionlessNodeRegistry.contract.Transact(opts, "updateNextQueuedValidatorIndex", _nextQueuedValidatorIndex)
}

// UpdateNextQueuedValidatorIndex is a paid mutator transaction binding the contract method 0xb8d2f06c.
//
// Solidity: function updateNextQueuedValidatorIndex(uint256 _nextQueuedValidatorIndex) returns()
func (_PermissionlessNodeRegistry *PermissionlessNodeRegistrySession) UpdateNextQueuedValidatorIndex(_nextQueuedValidatorIndex *big.Int) (*types.Transaction, error) {
	return _PermissionlessNodeRegistry.Contract.UpdateNextQueuedValidatorIndex(&_PermissionlessNodeRegistry.TransactOpts, _nextQueuedValidatorIndex)
}

// UpdateNextQueuedValidatorIndex is a paid mutator transaction binding the contract method 0xb8d2f06c.
//
// Solidity: function updateNextQueuedValidatorIndex(uint256 _nextQueuedValidatorIndex) returns()
func (_PermissionlessNodeRegistry *PermissionlessNodeRegistryTransactorSession) UpdateNextQueuedValidatorIndex(_nextQueuedValidatorIndex *big.Int) (*types.Transaction, error) {
	return _PermissionlessNodeRegistry.Contract.UpdateNextQueuedValidatorIndex(&_PermissionlessNodeRegistry.TransactOpts, _nextQueuedValidatorIndex)
}

// UpdateOperatorDetails is a paid mutator transaction binding the contract method 0x58a994ea.
//
// Solidity: function updateOperatorDetails(string _operatorName, address _rewardAddress) returns()
func (_PermissionlessNodeRegistry *PermissionlessNodeRegistryTransactor) UpdateOperatorDetails(opts *bind.TransactOpts, _operatorName string, _rewardAddress common.Address) (*types.Transaction, error) {
	return _PermissionlessNodeRegistry.contract.Transact(opts, "updateOperatorDetails", _operatorName, _rewardAddress)
}

// UpdateOperatorDetails is a paid mutator transaction binding the contract method 0x58a994ea.
//
// Solidity: function updateOperatorDetails(string _operatorName, address _rewardAddress) returns()
func (_PermissionlessNodeRegistry *PermissionlessNodeRegistrySession) UpdateOperatorDetails(_operatorName string, _rewardAddress common.Address) (*types.Transaction, error) {
	return _PermissionlessNodeRegistry.Contract.UpdateOperatorDetails(&_PermissionlessNodeRegistry.TransactOpts, _operatorName, _rewardAddress)
}

// UpdateOperatorDetails is a paid mutator transaction binding the contract method 0x58a994ea.
//
// Solidity: function updateOperatorDetails(string _operatorName, address _rewardAddress) returns()
func (_PermissionlessNodeRegistry *PermissionlessNodeRegistryTransactorSession) UpdateOperatorDetails(_operatorName string, _rewardAddress common.Address) (*types.Transaction, error) {
	return _PermissionlessNodeRegistry.Contract.UpdateOperatorDetails(&_PermissionlessNodeRegistry.TransactOpts, _operatorName, _rewardAddress)
}

// UpdateStaderConfig is a paid mutator transaction binding the contract method 0x9ee804cb.
//
// Solidity: function updateStaderConfig(address _staderConfig) returns()
func (_PermissionlessNodeRegistry *PermissionlessNodeRegistryTransactor) UpdateStaderConfig(opts *bind.TransactOpts, _staderConfig common.Address) (*types.Transaction, error) {
	return _PermissionlessNodeRegistry.contract.Transact(opts, "updateStaderConfig", _staderConfig)
}

// UpdateStaderConfig is a paid mutator transaction binding the contract method 0x9ee804cb.
//
// Solidity: function updateStaderConfig(address _staderConfig) returns()
func (_PermissionlessNodeRegistry *PermissionlessNodeRegistrySession) UpdateStaderConfig(_staderConfig common.Address) (*types.Transaction, error) {
	return _PermissionlessNodeRegistry.Contract.UpdateStaderConfig(&_PermissionlessNodeRegistry.TransactOpts, _staderConfig)
}

// UpdateStaderConfig is a paid mutator transaction binding the contract method 0x9ee804cb.
//
// Solidity: function updateStaderConfig(address _staderConfig) returns()
func (_PermissionlessNodeRegistry *PermissionlessNodeRegistryTransactorSession) UpdateStaderConfig(_staderConfig common.Address) (*types.Transaction, error) {
	return _PermissionlessNodeRegistry.Contract.UpdateStaderConfig(&_PermissionlessNodeRegistry.TransactOpts, _staderConfig)
}

// UpdateVerifiedKeysBatchSize is a paid mutator transaction binding the contract method 0xaf533aa8.
//
// Solidity: function updateVerifiedKeysBatchSize(uint256 _verifiedKeysBatchSize) returns()
func (_PermissionlessNodeRegistry *PermissionlessNodeRegistryTransactor) UpdateVerifiedKeysBatchSize(opts *bind.TransactOpts, _verifiedKeysBatchSize *big.Int) (*types.Transaction, error) {
	return _PermissionlessNodeRegistry.contract.Transact(opts, "updateVerifiedKeysBatchSize", _verifiedKeysBatchSize)
}

// UpdateVerifiedKeysBatchSize is a paid mutator transaction binding the contract method 0xaf533aa8.
//
// Solidity: function updateVerifiedKeysBatchSize(uint256 _verifiedKeysBatchSize) returns()
func (_PermissionlessNodeRegistry *PermissionlessNodeRegistrySession) UpdateVerifiedKeysBatchSize(_verifiedKeysBatchSize *big.Int) (*types.Transaction, error) {
	return _PermissionlessNodeRegistry.Contract.UpdateVerifiedKeysBatchSize(&_PermissionlessNodeRegistry.TransactOpts, _verifiedKeysBatchSize)
}

// UpdateVerifiedKeysBatchSize is a paid mutator transaction binding the contract method 0xaf533aa8.
//
// Solidity: function updateVerifiedKeysBatchSize(uint256 _verifiedKeysBatchSize) returns()
func (_PermissionlessNodeRegistry *PermissionlessNodeRegistryTransactorSession) UpdateVerifiedKeysBatchSize(_verifiedKeysBatchSize *big.Int) (*types.Transaction, error) {
	return _PermissionlessNodeRegistry.Contract.UpdateVerifiedKeysBatchSize(&_PermissionlessNodeRegistry.TransactOpts, _verifiedKeysBatchSize)
}

// WithdrawnValidators is a paid mutator transaction binding the contract method 0x264f27f3.
//
// Solidity: function withdrawnValidators(bytes[] _pubkeys) returns()
func (_PermissionlessNodeRegistry *PermissionlessNodeRegistryTransactor) WithdrawnValidators(opts *bind.TransactOpts, _pubkeys [][]byte) (*types.Transaction, error) {
	return _PermissionlessNodeRegistry.contract.Transact(opts, "withdrawnValidators", _pubkeys)
}

// WithdrawnValidators is a paid mutator transaction binding the contract method 0x264f27f3.
//
// Solidity: function withdrawnValidators(bytes[] _pubkeys) returns()
func (_PermissionlessNodeRegistry *PermissionlessNodeRegistrySession) WithdrawnValidators(_pubkeys [][]byte) (*types.Transaction, error) {
	return _PermissionlessNodeRegistry.Contract.WithdrawnValidators(&_PermissionlessNodeRegistry.TransactOpts, _pubkeys)
}

// WithdrawnValidators is a paid mutator transaction binding the contract method 0x264f27f3.
//
// Solidity: function withdrawnValidators(bytes[] _pubkeys) returns()
func (_PermissionlessNodeRegistry *PermissionlessNodeRegistryTransactorSession) WithdrawnValidators(_pubkeys [][]byte) (*types.Transaction, error) {
	return _PermissionlessNodeRegistry.Contract.WithdrawnValidators(&_PermissionlessNodeRegistry.TransactOpts, _pubkeys)
}

// PermissionlessNodeRegistryAddedValidatorKeyIterator is returned from FilterAddedValidatorKey and is used to iterate over the raw logs and unpacked data for AddedValidatorKey events raised by the PermissionlessNodeRegistry contract.
type PermissionlessNodeRegistryAddedValidatorKeyIterator struct {
	Event *PermissionlessNodeRegistryAddedValidatorKey // Event containing the contract specifics and raw log

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
func (it *PermissionlessNodeRegistryAddedValidatorKeyIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PermissionlessNodeRegistryAddedValidatorKey)
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
		it.Event = new(PermissionlessNodeRegistryAddedValidatorKey)
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
func (it *PermissionlessNodeRegistryAddedValidatorKeyIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PermissionlessNodeRegistryAddedValidatorKeyIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PermissionlessNodeRegistryAddedValidatorKey represents a AddedValidatorKey event raised by the PermissionlessNodeRegistry contract.
type PermissionlessNodeRegistryAddedValidatorKey struct {
	NodeOperator common.Address
	Pubkey       []byte
	ValidatorId  *big.Int
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterAddedValidatorKey is a free log retrieval operation binding the contract event 0xab5128638b64e6216e80dfafa70d3cb6d54913a536dc41e76eb4a04cfbe979cf.
//
// Solidity: event AddedValidatorKey(address indexed nodeOperator, bytes pubkey, uint256 validatorId)
func (_PermissionlessNodeRegistry *PermissionlessNodeRegistryFilterer) FilterAddedValidatorKey(opts *bind.FilterOpts, nodeOperator []common.Address) (*PermissionlessNodeRegistryAddedValidatorKeyIterator, error) {

	var nodeOperatorRule []interface{}
	for _, nodeOperatorItem := range nodeOperator {
		nodeOperatorRule = append(nodeOperatorRule, nodeOperatorItem)
	}

	logs, sub, err := _PermissionlessNodeRegistry.contract.FilterLogs(opts, "AddedValidatorKey", nodeOperatorRule)
	if err != nil {
		return nil, err
	}
	return &PermissionlessNodeRegistryAddedValidatorKeyIterator{contract: _PermissionlessNodeRegistry.contract, event: "AddedValidatorKey", logs: logs, sub: sub}, nil
}

// WatchAddedValidatorKey is a free log subscription operation binding the contract event 0xab5128638b64e6216e80dfafa70d3cb6d54913a536dc41e76eb4a04cfbe979cf.
//
// Solidity: event AddedValidatorKey(address indexed nodeOperator, bytes pubkey, uint256 validatorId)
func (_PermissionlessNodeRegistry *PermissionlessNodeRegistryFilterer) WatchAddedValidatorKey(opts *bind.WatchOpts, sink chan<- *PermissionlessNodeRegistryAddedValidatorKey, nodeOperator []common.Address) (event.Subscription, error) {

	var nodeOperatorRule []interface{}
	for _, nodeOperatorItem := range nodeOperator {
		nodeOperatorRule = append(nodeOperatorRule, nodeOperatorItem)
	}

	logs, sub, err := _PermissionlessNodeRegistry.contract.WatchLogs(opts, "AddedValidatorKey", nodeOperatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PermissionlessNodeRegistryAddedValidatorKey)
				if err := _PermissionlessNodeRegistry.contract.UnpackLog(event, "AddedValidatorKey", log); err != nil {
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

// ParseAddedValidatorKey is a log parse operation binding the contract event 0xab5128638b64e6216e80dfafa70d3cb6d54913a536dc41e76eb4a04cfbe979cf.
//
// Solidity: event AddedValidatorKey(address indexed nodeOperator, bytes pubkey, uint256 validatorId)
func (_PermissionlessNodeRegistry *PermissionlessNodeRegistryFilterer) ParseAddedValidatorKey(log types.Log) (*PermissionlessNodeRegistryAddedValidatorKey, error) {
	event := new(PermissionlessNodeRegistryAddedValidatorKey)
	if err := _PermissionlessNodeRegistry.contract.UnpackLog(event, "AddedValidatorKey", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PermissionlessNodeRegistryDecreasedTotalActiveValidatorCountIterator is returned from FilterDecreasedTotalActiveValidatorCount and is used to iterate over the raw logs and unpacked data for DecreasedTotalActiveValidatorCount events raised by the PermissionlessNodeRegistry contract.
type PermissionlessNodeRegistryDecreasedTotalActiveValidatorCountIterator struct {
	Event *PermissionlessNodeRegistryDecreasedTotalActiveValidatorCount // Event containing the contract specifics and raw log

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
func (it *PermissionlessNodeRegistryDecreasedTotalActiveValidatorCountIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PermissionlessNodeRegistryDecreasedTotalActiveValidatorCount)
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
		it.Event = new(PermissionlessNodeRegistryDecreasedTotalActiveValidatorCount)
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
func (it *PermissionlessNodeRegistryDecreasedTotalActiveValidatorCountIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PermissionlessNodeRegistryDecreasedTotalActiveValidatorCountIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PermissionlessNodeRegistryDecreasedTotalActiveValidatorCount represents a DecreasedTotalActiveValidatorCount event raised by the PermissionlessNodeRegistry contract.
type PermissionlessNodeRegistryDecreasedTotalActiveValidatorCount struct {
	TotalActiveValidatorCount *big.Int
	Raw                       types.Log // Blockchain specific contextual infos
}

// FilterDecreasedTotalActiveValidatorCount is a free log retrieval operation binding the contract event 0x5040a06a11b7d9b75fc56fbbd207905dbaa4ac86c0dc9cc7fff40cd1d92aece3.
//
// Solidity: event DecreasedTotalActiveValidatorCount(uint256 totalActiveValidatorCount)
func (_PermissionlessNodeRegistry *PermissionlessNodeRegistryFilterer) FilterDecreasedTotalActiveValidatorCount(opts *bind.FilterOpts) (*PermissionlessNodeRegistryDecreasedTotalActiveValidatorCountIterator, error) {

	logs, sub, err := _PermissionlessNodeRegistry.contract.FilterLogs(opts, "DecreasedTotalActiveValidatorCount")
	if err != nil {
		return nil, err
	}
	return &PermissionlessNodeRegistryDecreasedTotalActiveValidatorCountIterator{contract: _PermissionlessNodeRegistry.contract, event: "DecreasedTotalActiveValidatorCount", logs: logs, sub: sub}, nil
}

// WatchDecreasedTotalActiveValidatorCount is a free log subscription operation binding the contract event 0x5040a06a11b7d9b75fc56fbbd207905dbaa4ac86c0dc9cc7fff40cd1d92aece3.
//
// Solidity: event DecreasedTotalActiveValidatorCount(uint256 totalActiveValidatorCount)
func (_PermissionlessNodeRegistry *PermissionlessNodeRegistryFilterer) WatchDecreasedTotalActiveValidatorCount(opts *bind.WatchOpts, sink chan<- *PermissionlessNodeRegistryDecreasedTotalActiveValidatorCount) (event.Subscription, error) {

	logs, sub, err := _PermissionlessNodeRegistry.contract.WatchLogs(opts, "DecreasedTotalActiveValidatorCount")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PermissionlessNodeRegistryDecreasedTotalActiveValidatorCount)
				if err := _PermissionlessNodeRegistry.contract.UnpackLog(event, "DecreasedTotalActiveValidatorCount", log); err != nil {
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

// ParseDecreasedTotalActiveValidatorCount is a log parse operation binding the contract event 0x5040a06a11b7d9b75fc56fbbd207905dbaa4ac86c0dc9cc7fff40cd1d92aece3.
//
// Solidity: event DecreasedTotalActiveValidatorCount(uint256 totalActiveValidatorCount)
func (_PermissionlessNodeRegistry *PermissionlessNodeRegistryFilterer) ParseDecreasedTotalActiveValidatorCount(log types.Log) (*PermissionlessNodeRegistryDecreasedTotalActiveValidatorCount, error) {
	event := new(PermissionlessNodeRegistryDecreasedTotalActiveValidatorCount)
	if err := _PermissionlessNodeRegistry.contract.UnpackLog(event, "DecreasedTotalActiveValidatorCount", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PermissionlessNodeRegistryIncreasedTotalActiveValidatorCountIterator is returned from FilterIncreasedTotalActiveValidatorCount and is used to iterate over the raw logs and unpacked data for IncreasedTotalActiveValidatorCount events raised by the PermissionlessNodeRegistry contract.
type PermissionlessNodeRegistryIncreasedTotalActiveValidatorCountIterator struct {
	Event *PermissionlessNodeRegistryIncreasedTotalActiveValidatorCount // Event containing the contract specifics and raw log

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
func (it *PermissionlessNodeRegistryIncreasedTotalActiveValidatorCountIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PermissionlessNodeRegistryIncreasedTotalActiveValidatorCount)
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
		it.Event = new(PermissionlessNodeRegistryIncreasedTotalActiveValidatorCount)
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
func (it *PermissionlessNodeRegistryIncreasedTotalActiveValidatorCountIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PermissionlessNodeRegistryIncreasedTotalActiveValidatorCountIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PermissionlessNodeRegistryIncreasedTotalActiveValidatorCount represents a IncreasedTotalActiveValidatorCount event raised by the PermissionlessNodeRegistry contract.
type PermissionlessNodeRegistryIncreasedTotalActiveValidatorCount struct {
	TotalActiveValidatorCount *big.Int
	Raw                       types.Log // Blockchain specific contextual infos
}

// FilterIncreasedTotalActiveValidatorCount is a free log retrieval operation binding the contract event 0x5818a627697795ff3c3403f320c7549835866cfb64a0b06a6f7f077bc478e9f2.
//
// Solidity: event IncreasedTotalActiveValidatorCount(uint256 totalActiveValidatorCount)
func (_PermissionlessNodeRegistry *PermissionlessNodeRegistryFilterer) FilterIncreasedTotalActiveValidatorCount(opts *bind.FilterOpts) (*PermissionlessNodeRegistryIncreasedTotalActiveValidatorCountIterator, error) {

	logs, sub, err := _PermissionlessNodeRegistry.contract.FilterLogs(opts, "IncreasedTotalActiveValidatorCount")
	if err != nil {
		return nil, err
	}
	return &PermissionlessNodeRegistryIncreasedTotalActiveValidatorCountIterator{contract: _PermissionlessNodeRegistry.contract, event: "IncreasedTotalActiveValidatorCount", logs: logs, sub: sub}, nil
}

// WatchIncreasedTotalActiveValidatorCount is a free log subscription operation binding the contract event 0x5818a627697795ff3c3403f320c7549835866cfb64a0b06a6f7f077bc478e9f2.
//
// Solidity: event IncreasedTotalActiveValidatorCount(uint256 totalActiveValidatorCount)
func (_PermissionlessNodeRegistry *PermissionlessNodeRegistryFilterer) WatchIncreasedTotalActiveValidatorCount(opts *bind.WatchOpts, sink chan<- *PermissionlessNodeRegistryIncreasedTotalActiveValidatorCount) (event.Subscription, error) {

	logs, sub, err := _PermissionlessNodeRegistry.contract.WatchLogs(opts, "IncreasedTotalActiveValidatorCount")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PermissionlessNodeRegistryIncreasedTotalActiveValidatorCount)
				if err := _PermissionlessNodeRegistry.contract.UnpackLog(event, "IncreasedTotalActiveValidatorCount", log); err != nil {
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

// ParseIncreasedTotalActiveValidatorCount is a log parse operation binding the contract event 0x5818a627697795ff3c3403f320c7549835866cfb64a0b06a6f7f077bc478e9f2.
//
// Solidity: event IncreasedTotalActiveValidatorCount(uint256 totalActiveValidatorCount)
func (_PermissionlessNodeRegistry *PermissionlessNodeRegistryFilterer) ParseIncreasedTotalActiveValidatorCount(log types.Log) (*PermissionlessNodeRegistryIncreasedTotalActiveValidatorCount, error) {
	event := new(PermissionlessNodeRegistryIncreasedTotalActiveValidatorCount)
	if err := _PermissionlessNodeRegistry.contract.UnpackLog(event, "IncreasedTotalActiveValidatorCount", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PermissionlessNodeRegistryInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the PermissionlessNodeRegistry contract.
type PermissionlessNodeRegistryInitializedIterator struct {
	Event *PermissionlessNodeRegistryInitialized // Event containing the contract specifics and raw log

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
func (it *PermissionlessNodeRegistryInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PermissionlessNodeRegistryInitialized)
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
		it.Event = new(PermissionlessNodeRegistryInitialized)
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
func (it *PermissionlessNodeRegistryInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PermissionlessNodeRegistryInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PermissionlessNodeRegistryInitialized represents a Initialized event raised by the PermissionlessNodeRegistry contract.
type PermissionlessNodeRegistryInitialized struct {
	Version uint8
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_PermissionlessNodeRegistry *PermissionlessNodeRegistryFilterer) FilterInitialized(opts *bind.FilterOpts) (*PermissionlessNodeRegistryInitializedIterator, error) {

	logs, sub, err := _PermissionlessNodeRegistry.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &PermissionlessNodeRegistryInitializedIterator{contract: _PermissionlessNodeRegistry.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_PermissionlessNodeRegistry *PermissionlessNodeRegistryFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *PermissionlessNodeRegistryInitialized) (event.Subscription, error) {

	logs, sub, err := _PermissionlessNodeRegistry.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PermissionlessNodeRegistryInitialized)
				if err := _PermissionlessNodeRegistry.contract.UnpackLog(event, "Initialized", log); err != nil {
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
func (_PermissionlessNodeRegistry *PermissionlessNodeRegistryFilterer) ParseInitialized(log types.Log) (*PermissionlessNodeRegistryInitialized, error) {
	event := new(PermissionlessNodeRegistryInitialized)
	if err := _PermissionlessNodeRegistry.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PermissionlessNodeRegistryOnboardedOperatorIterator is returned from FilterOnboardedOperator and is used to iterate over the raw logs and unpacked data for OnboardedOperator events raised by the PermissionlessNodeRegistry contract.
type PermissionlessNodeRegistryOnboardedOperatorIterator struct {
	Event *PermissionlessNodeRegistryOnboardedOperator // Event containing the contract specifics and raw log

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
func (it *PermissionlessNodeRegistryOnboardedOperatorIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PermissionlessNodeRegistryOnboardedOperator)
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
		it.Event = new(PermissionlessNodeRegistryOnboardedOperator)
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
func (it *PermissionlessNodeRegistryOnboardedOperatorIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PermissionlessNodeRegistryOnboardedOperatorIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PermissionlessNodeRegistryOnboardedOperator represents a OnboardedOperator event raised by the PermissionlessNodeRegistry contract.
type PermissionlessNodeRegistryOnboardedOperator struct {
	NodeOperator            common.Address
	NodeRewardAddress       common.Address
	OperatorId              *big.Int
	OptInForSocializingPool bool
	Raw                     types.Log // Blockchain specific contextual infos
}

// FilterOnboardedOperator is a free log retrieval operation binding the contract event 0x55b1a82a03cdb2847b1ec26dcac8ce8b3fc5f310388290b048c0ee9ac1ce8dd4.
//
// Solidity: event OnboardedOperator(address indexed nodeOperator, address nodeRewardAddress, uint256 operatorId, bool optInForSocializingPool)
func (_PermissionlessNodeRegistry *PermissionlessNodeRegistryFilterer) FilterOnboardedOperator(opts *bind.FilterOpts, nodeOperator []common.Address) (*PermissionlessNodeRegistryOnboardedOperatorIterator, error) {

	var nodeOperatorRule []interface{}
	for _, nodeOperatorItem := range nodeOperator {
		nodeOperatorRule = append(nodeOperatorRule, nodeOperatorItem)
	}

	logs, sub, err := _PermissionlessNodeRegistry.contract.FilterLogs(opts, "OnboardedOperator", nodeOperatorRule)
	if err != nil {
		return nil, err
	}
	return &PermissionlessNodeRegistryOnboardedOperatorIterator{contract: _PermissionlessNodeRegistry.contract, event: "OnboardedOperator", logs: logs, sub: sub}, nil
}

// WatchOnboardedOperator is a free log subscription operation binding the contract event 0x55b1a82a03cdb2847b1ec26dcac8ce8b3fc5f310388290b048c0ee9ac1ce8dd4.
//
// Solidity: event OnboardedOperator(address indexed nodeOperator, address nodeRewardAddress, uint256 operatorId, bool optInForSocializingPool)
func (_PermissionlessNodeRegistry *PermissionlessNodeRegistryFilterer) WatchOnboardedOperator(opts *bind.WatchOpts, sink chan<- *PermissionlessNodeRegistryOnboardedOperator, nodeOperator []common.Address) (event.Subscription, error) {

	var nodeOperatorRule []interface{}
	for _, nodeOperatorItem := range nodeOperator {
		nodeOperatorRule = append(nodeOperatorRule, nodeOperatorItem)
	}

	logs, sub, err := _PermissionlessNodeRegistry.contract.WatchLogs(opts, "OnboardedOperator", nodeOperatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PermissionlessNodeRegistryOnboardedOperator)
				if err := _PermissionlessNodeRegistry.contract.UnpackLog(event, "OnboardedOperator", log); err != nil {
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

// ParseOnboardedOperator is a log parse operation binding the contract event 0x55b1a82a03cdb2847b1ec26dcac8ce8b3fc5f310388290b048c0ee9ac1ce8dd4.
//
// Solidity: event OnboardedOperator(address indexed nodeOperator, address nodeRewardAddress, uint256 operatorId, bool optInForSocializingPool)
func (_PermissionlessNodeRegistry *PermissionlessNodeRegistryFilterer) ParseOnboardedOperator(log types.Log) (*PermissionlessNodeRegistryOnboardedOperator, error) {
	event := new(PermissionlessNodeRegistryOnboardedOperator)
	if err := _PermissionlessNodeRegistry.contract.UnpackLog(event, "OnboardedOperator", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PermissionlessNodeRegistryPausedIterator is returned from FilterPaused and is used to iterate over the raw logs and unpacked data for Paused events raised by the PermissionlessNodeRegistry contract.
type PermissionlessNodeRegistryPausedIterator struct {
	Event *PermissionlessNodeRegistryPaused // Event containing the contract specifics and raw log

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
func (it *PermissionlessNodeRegistryPausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PermissionlessNodeRegistryPaused)
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
		it.Event = new(PermissionlessNodeRegistryPaused)
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
func (it *PermissionlessNodeRegistryPausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PermissionlessNodeRegistryPausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PermissionlessNodeRegistryPaused represents a Paused event raised by the PermissionlessNodeRegistry contract.
type PermissionlessNodeRegistryPaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterPaused is a free log retrieval operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_PermissionlessNodeRegistry *PermissionlessNodeRegistryFilterer) FilterPaused(opts *bind.FilterOpts) (*PermissionlessNodeRegistryPausedIterator, error) {

	logs, sub, err := _PermissionlessNodeRegistry.contract.FilterLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return &PermissionlessNodeRegistryPausedIterator{contract: _PermissionlessNodeRegistry.contract, event: "Paused", logs: logs, sub: sub}, nil
}

// WatchPaused is a free log subscription operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_PermissionlessNodeRegistry *PermissionlessNodeRegistryFilterer) WatchPaused(opts *bind.WatchOpts, sink chan<- *PermissionlessNodeRegistryPaused) (event.Subscription, error) {

	logs, sub, err := _PermissionlessNodeRegistry.contract.WatchLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PermissionlessNodeRegistryPaused)
				if err := _PermissionlessNodeRegistry.contract.UnpackLog(event, "Paused", log); err != nil {
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

// ParsePaused is a log parse operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_PermissionlessNodeRegistry *PermissionlessNodeRegistryFilterer) ParsePaused(log types.Log) (*PermissionlessNodeRegistryPaused, error) {
	event := new(PermissionlessNodeRegistryPaused)
	if err := _PermissionlessNodeRegistry.contract.UnpackLog(event, "Paused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PermissionlessNodeRegistryRoleAdminChangedIterator is returned from FilterRoleAdminChanged and is used to iterate over the raw logs and unpacked data for RoleAdminChanged events raised by the PermissionlessNodeRegistry contract.
type PermissionlessNodeRegistryRoleAdminChangedIterator struct {
	Event *PermissionlessNodeRegistryRoleAdminChanged // Event containing the contract specifics and raw log

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
func (it *PermissionlessNodeRegistryRoleAdminChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PermissionlessNodeRegistryRoleAdminChanged)
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
		it.Event = new(PermissionlessNodeRegistryRoleAdminChanged)
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
func (it *PermissionlessNodeRegistryRoleAdminChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PermissionlessNodeRegistryRoleAdminChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PermissionlessNodeRegistryRoleAdminChanged represents a RoleAdminChanged event raised by the PermissionlessNodeRegistry contract.
type PermissionlessNodeRegistryRoleAdminChanged struct {
	Role              [32]byte
	PreviousAdminRole [32]byte
	NewAdminRole      [32]byte
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterRoleAdminChanged is a free log retrieval operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_PermissionlessNodeRegistry *PermissionlessNodeRegistryFilterer) FilterRoleAdminChanged(opts *bind.FilterOpts, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (*PermissionlessNodeRegistryRoleAdminChangedIterator, error) {

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

	logs, sub, err := _PermissionlessNodeRegistry.contract.FilterLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return &PermissionlessNodeRegistryRoleAdminChangedIterator{contract: _PermissionlessNodeRegistry.contract, event: "RoleAdminChanged", logs: logs, sub: sub}, nil
}

// WatchRoleAdminChanged is a free log subscription operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_PermissionlessNodeRegistry *PermissionlessNodeRegistryFilterer) WatchRoleAdminChanged(opts *bind.WatchOpts, sink chan<- *PermissionlessNodeRegistryRoleAdminChanged, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (event.Subscription, error) {

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

	logs, sub, err := _PermissionlessNodeRegistry.contract.WatchLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PermissionlessNodeRegistryRoleAdminChanged)
				if err := _PermissionlessNodeRegistry.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
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
func (_PermissionlessNodeRegistry *PermissionlessNodeRegistryFilterer) ParseRoleAdminChanged(log types.Log) (*PermissionlessNodeRegistryRoleAdminChanged, error) {
	event := new(PermissionlessNodeRegistryRoleAdminChanged)
	if err := _PermissionlessNodeRegistry.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PermissionlessNodeRegistryRoleGrantedIterator is returned from FilterRoleGranted and is used to iterate over the raw logs and unpacked data for RoleGranted events raised by the PermissionlessNodeRegistry contract.
type PermissionlessNodeRegistryRoleGrantedIterator struct {
	Event *PermissionlessNodeRegistryRoleGranted // Event containing the contract specifics and raw log

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
func (it *PermissionlessNodeRegistryRoleGrantedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PermissionlessNodeRegistryRoleGranted)
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
		it.Event = new(PermissionlessNodeRegistryRoleGranted)
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
func (it *PermissionlessNodeRegistryRoleGrantedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PermissionlessNodeRegistryRoleGrantedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PermissionlessNodeRegistryRoleGranted represents a RoleGranted event raised by the PermissionlessNodeRegistry contract.
type PermissionlessNodeRegistryRoleGranted struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleGranted is a free log retrieval operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_PermissionlessNodeRegistry *PermissionlessNodeRegistryFilterer) FilterRoleGranted(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*PermissionlessNodeRegistryRoleGrantedIterator, error) {

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

	logs, sub, err := _PermissionlessNodeRegistry.contract.FilterLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &PermissionlessNodeRegistryRoleGrantedIterator{contract: _PermissionlessNodeRegistry.contract, event: "RoleGranted", logs: logs, sub: sub}, nil
}

// WatchRoleGranted is a free log subscription operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_PermissionlessNodeRegistry *PermissionlessNodeRegistryFilterer) WatchRoleGranted(opts *bind.WatchOpts, sink chan<- *PermissionlessNodeRegistryRoleGranted, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _PermissionlessNodeRegistry.contract.WatchLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PermissionlessNodeRegistryRoleGranted)
				if err := _PermissionlessNodeRegistry.contract.UnpackLog(event, "RoleGranted", log); err != nil {
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
func (_PermissionlessNodeRegistry *PermissionlessNodeRegistryFilterer) ParseRoleGranted(log types.Log) (*PermissionlessNodeRegistryRoleGranted, error) {
	event := new(PermissionlessNodeRegistryRoleGranted)
	if err := _PermissionlessNodeRegistry.contract.UnpackLog(event, "RoleGranted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PermissionlessNodeRegistryRoleRevokedIterator is returned from FilterRoleRevoked and is used to iterate over the raw logs and unpacked data for RoleRevoked events raised by the PermissionlessNodeRegistry contract.
type PermissionlessNodeRegistryRoleRevokedIterator struct {
	Event *PermissionlessNodeRegistryRoleRevoked // Event containing the contract specifics and raw log

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
func (it *PermissionlessNodeRegistryRoleRevokedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PermissionlessNodeRegistryRoleRevoked)
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
		it.Event = new(PermissionlessNodeRegistryRoleRevoked)
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
func (it *PermissionlessNodeRegistryRoleRevokedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PermissionlessNodeRegistryRoleRevokedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PermissionlessNodeRegistryRoleRevoked represents a RoleRevoked event raised by the PermissionlessNodeRegistry contract.
type PermissionlessNodeRegistryRoleRevoked struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleRevoked is a free log retrieval operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_PermissionlessNodeRegistry *PermissionlessNodeRegistryFilterer) FilterRoleRevoked(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*PermissionlessNodeRegistryRoleRevokedIterator, error) {

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

	logs, sub, err := _PermissionlessNodeRegistry.contract.FilterLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &PermissionlessNodeRegistryRoleRevokedIterator{contract: _PermissionlessNodeRegistry.contract, event: "RoleRevoked", logs: logs, sub: sub}, nil
}

// WatchRoleRevoked is a free log subscription operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_PermissionlessNodeRegistry *PermissionlessNodeRegistryFilterer) WatchRoleRevoked(opts *bind.WatchOpts, sink chan<- *PermissionlessNodeRegistryRoleRevoked, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _PermissionlessNodeRegistry.contract.WatchLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PermissionlessNodeRegistryRoleRevoked)
				if err := _PermissionlessNodeRegistry.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
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
func (_PermissionlessNodeRegistry *PermissionlessNodeRegistryFilterer) ParseRoleRevoked(log types.Log) (*PermissionlessNodeRegistryRoleRevoked, error) {
	event := new(PermissionlessNodeRegistryRoleRevoked)
	if err := _PermissionlessNodeRegistry.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PermissionlessNodeRegistryTransferredCollateralToPoolIterator is returned from FilterTransferredCollateralToPool and is used to iterate over the raw logs and unpacked data for TransferredCollateralToPool events raised by the PermissionlessNodeRegistry contract.
type PermissionlessNodeRegistryTransferredCollateralToPoolIterator struct {
	Event *PermissionlessNodeRegistryTransferredCollateralToPool // Event containing the contract specifics and raw log

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
func (it *PermissionlessNodeRegistryTransferredCollateralToPoolIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PermissionlessNodeRegistryTransferredCollateralToPool)
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
		it.Event = new(PermissionlessNodeRegistryTransferredCollateralToPool)
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
func (it *PermissionlessNodeRegistryTransferredCollateralToPoolIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PermissionlessNodeRegistryTransferredCollateralToPoolIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PermissionlessNodeRegistryTransferredCollateralToPool represents a TransferredCollateralToPool event raised by the PermissionlessNodeRegistry contract.
type PermissionlessNodeRegistryTransferredCollateralToPool struct {
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterTransferredCollateralToPool is a free log retrieval operation binding the contract event 0x9407b62b10143b3ae08ce1cc7f9b66af41a4431ad59107e53ff54d6401e0730a.
//
// Solidity: event TransferredCollateralToPool(uint256 amount)
func (_PermissionlessNodeRegistry *PermissionlessNodeRegistryFilterer) FilterTransferredCollateralToPool(opts *bind.FilterOpts) (*PermissionlessNodeRegistryTransferredCollateralToPoolIterator, error) {

	logs, sub, err := _PermissionlessNodeRegistry.contract.FilterLogs(opts, "TransferredCollateralToPool")
	if err != nil {
		return nil, err
	}
	return &PermissionlessNodeRegistryTransferredCollateralToPoolIterator{contract: _PermissionlessNodeRegistry.contract, event: "TransferredCollateralToPool", logs: logs, sub: sub}, nil
}

// WatchTransferredCollateralToPool is a free log subscription operation binding the contract event 0x9407b62b10143b3ae08ce1cc7f9b66af41a4431ad59107e53ff54d6401e0730a.
//
// Solidity: event TransferredCollateralToPool(uint256 amount)
func (_PermissionlessNodeRegistry *PermissionlessNodeRegistryFilterer) WatchTransferredCollateralToPool(opts *bind.WatchOpts, sink chan<- *PermissionlessNodeRegistryTransferredCollateralToPool) (event.Subscription, error) {

	logs, sub, err := _PermissionlessNodeRegistry.contract.WatchLogs(opts, "TransferredCollateralToPool")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PermissionlessNodeRegistryTransferredCollateralToPool)
				if err := _PermissionlessNodeRegistry.contract.UnpackLog(event, "TransferredCollateralToPool", log); err != nil {
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

// ParseTransferredCollateralToPool is a log parse operation binding the contract event 0x9407b62b10143b3ae08ce1cc7f9b66af41a4431ad59107e53ff54d6401e0730a.
//
// Solidity: event TransferredCollateralToPool(uint256 amount)
func (_PermissionlessNodeRegistry *PermissionlessNodeRegistryFilterer) ParseTransferredCollateralToPool(log types.Log) (*PermissionlessNodeRegistryTransferredCollateralToPool, error) {
	event := new(PermissionlessNodeRegistryTransferredCollateralToPool)
	if err := _PermissionlessNodeRegistry.contract.UnpackLog(event, "TransferredCollateralToPool", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PermissionlessNodeRegistryUnpausedIterator is returned from FilterUnpaused and is used to iterate over the raw logs and unpacked data for Unpaused events raised by the PermissionlessNodeRegistry contract.
type PermissionlessNodeRegistryUnpausedIterator struct {
	Event *PermissionlessNodeRegistryUnpaused // Event containing the contract specifics and raw log

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
func (it *PermissionlessNodeRegistryUnpausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PermissionlessNodeRegistryUnpaused)
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
		it.Event = new(PermissionlessNodeRegistryUnpaused)
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
func (it *PermissionlessNodeRegistryUnpausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PermissionlessNodeRegistryUnpausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PermissionlessNodeRegistryUnpaused represents a Unpaused event raised by the PermissionlessNodeRegistry contract.
type PermissionlessNodeRegistryUnpaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterUnpaused is a free log retrieval operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_PermissionlessNodeRegistry *PermissionlessNodeRegistryFilterer) FilterUnpaused(opts *bind.FilterOpts) (*PermissionlessNodeRegistryUnpausedIterator, error) {

	logs, sub, err := _PermissionlessNodeRegistry.contract.FilterLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return &PermissionlessNodeRegistryUnpausedIterator{contract: _PermissionlessNodeRegistry.contract, event: "Unpaused", logs: logs, sub: sub}, nil
}

// WatchUnpaused is a free log subscription operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_PermissionlessNodeRegistry *PermissionlessNodeRegistryFilterer) WatchUnpaused(opts *bind.WatchOpts, sink chan<- *PermissionlessNodeRegistryUnpaused) (event.Subscription, error) {

	logs, sub, err := _PermissionlessNodeRegistry.contract.WatchLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PermissionlessNodeRegistryUnpaused)
				if err := _PermissionlessNodeRegistry.contract.UnpackLog(event, "Unpaused", log); err != nil {
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

// ParseUnpaused is a log parse operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_PermissionlessNodeRegistry *PermissionlessNodeRegistryFilterer) ParseUnpaused(log types.Log) (*PermissionlessNodeRegistryUnpaused, error) {
	event := new(PermissionlessNodeRegistryUnpaused)
	if err := _PermissionlessNodeRegistry.contract.UnpackLog(event, "Unpaused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PermissionlessNodeRegistryUpdatedInputKeyCountLimitIterator is returned from FilterUpdatedInputKeyCountLimit and is used to iterate over the raw logs and unpacked data for UpdatedInputKeyCountLimit events raised by the PermissionlessNodeRegistry contract.
type PermissionlessNodeRegistryUpdatedInputKeyCountLimitIterator struct {
	Event *PermissionlessNodeRegistryUpdatedInputKeyCountLimit // Event containing the contract specifics and raw log

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
func (it *PermissionlessNodeRegistryUpdatedInputKeyCountLimitIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PermissionlessNodeRegistryUpdatedInputKeyCountLimit)
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
		it.Event = new(PermissionlessNodeRegistryUpdatedInputKeyCountLimit)
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
func (it *PermissionlessNodeRegistryUpdatedInputKeyCountLimitIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PermissionlessNodeRegistryUpdatedInputKeyCountLimitIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PermissionlessNodeRegistryUpdatedInputKeyCountLimit represents a UpdatedInputKeyCountLimit event raised by the PermissionlessNodeRegistry contract.
type PermissionlessNodeRegistryUpdatedInputKeyCountLimit struct {
	BatchKeyDepositLimit *big.Int
	Raw                  types.Log // Blockchain specific contextual infos
}

// FilterUpdatedInputKeyCountLimit is a free log retrieval operation binding the contract event 0x5fd0fcd821abb4c92d47c4740e5f4a25ef35e99ee092d170faa0e5cb47013c36.
//
// Solidity: event UpdatedInputKeyCountLimit(uint256 batchKeyDepositLimit)
func (_PermissionlessNodeRegistry *PermissionlessNodeRegistryFilterer) FilterUpdatedInputKeyCountLimit(opts *bind.FilterOpts) (*PermissionlessNodeRegistryUpdatedInputKeyCountLimitIterator, error) {

	logs, sub, err := _PermissionlessNodeRegistry.contract.FilterLogs(opts, "UpdatedInputKeyCountLimit")
	if err != nil {
		return nil, err
	}
	return &PermissionlessNodeRegistryUpdatedInputKeyCountLimitIterator{contract: _PermissionlessNodeRegistry.contract, event: "UpdatedInputKeyCountLimit", logs: logs, sub: sub}, nil
}

// WatchUpdatedInputKeyCountLimit is a free log subscription operation binding the contract event 0x5fd0fcd821abb4c92d47c4740e5f4a25ef35e99ee092d170faa0e5cb47013c36.
//
// Solidity: event UpdatedInputKeyCountLimit(uint256 batchKeyDepositLimit)
func (_PermissionlessNodeRegistry *PermissionlessNodeRegistryFilterer) WatchUpdatedInputKeyCountLimit(opts *bind.WatchOpts, sink chan<- *PermissionlessNodeRegistryUpdatedInputKeyCountLimit) (event.Subscription, error) {

	logs, sub, err := _PermissionlessNodeRegistry.contract.WatchLogs(opts, "UpdatedInputKeyCountLimit")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PermissionlessNodeRegistryUpdatedInputKeyCountLimit)
				if err := _PermissionlessNodeRegistry.contract.UnpackLog(event, "UpdatedInputKeyCountLimit", log); err != nil {
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

// ParseUpdatedInputKeyCountLimit is a log parse operation binding the contract event 0x5fd0fcd821abb4c92d47c4740e5f4a25ef35e99ee092d170faa0e5cb47013c36.
//
// Solidity: event UpdatedInputKeyCountLimit(uint256 batchKeyDepositLimit)
func (_PermissionlessNodeRegistry *PermissionlessNodeRegistryFilterer) ParseUpdatedInputKeyCountLimit(log types.Log) (*PermissionlessNodeRegistryUpdatedInputKeyCountLimit, error) {
	event := new(PermissionlessNodeRegistryUpdatedInputKeyCountLimit)
	if err := _PermissionlessNodeRegistry.contract.UnpackLog(event, "UpdatedInputKeyCountLimit", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PermissionlessNodeRegistryUpdatedMaxNonTerminalKeyPerOperatorIterator is returned from FilterUpdatedMaxNonTerminalKeyPerOperator and is used to iterate over the raw logs and unpacked data for UpdatedMaxNonTerminalKeyPerOperator events raised by the PermissionlessNodeRegistry contract.
type PermissionlessNodeRegistryUpdatedMaxNonTerminalKeyPerOperatorIterator struct {
	Event *PermissionlessNodeRegistryUpdatedMaxNonTerminalKeyPerOperator // Event containing the contract specifics and raw log

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
func (it *PermissionlessNodeRegistryUpdatedMaxNonTerminalKeyPerOperatorIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PermissionlessNodeRegistryUpdatedMaxNonTerminalKeyPerOperator)
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
		it.Event = new(PermissionlessNodeRegistryUpdatedMaxNonTerminalKeyPerOperator)
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
func (it *PermissionlessNodeRegistryUpdatedMaxNonTerminalKeyPerOperatorIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PermissionlessNodeRegistryUpdatedMaxNonTerminalKeyPerOperatorIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PermissionlessNodeRegistryUpdatedMaxNonTerminalKeyPerOperator represents a UpdatedMaxNonTerminalKeyPerOperator event raised by the PermissionlessNodeRegistry contract.
type PermissionlessNodeRegistryUpdatedMaxNonTerminalKeyPerOperator struct {
	MaxNonTerminalKeyPerOperator uint64
	Raw                          types.Log // Blockchain specific contextual infos
}

// FilterUpdatedMaxNonTerminalKeyPerOperator is a free log retrieval operation binding the contract event 0xacda2fe79efeffc359206ddeeb45f26ba1596223e01e1585458603af76e880a2.
//
// Solidity: event UpdatedMaxNonTerminalKeyPerOperator(uint64 maxNonTerminalKeyPerOperator)
func (_PermissionlessNodeRegistry *PermissionlessNodeRegistryFilterer) FilterUpdatedMaxNonTerminalKeyPerOperator(opts *bind.FilterOpts) (*PermissionlessNodeRegistryUpdatedMaxNonTerminalKeyPerOperatorIterator, error) {

	logs, sub, err := _PermissionlessNodeRegistry.contract.FilterLogs(opts, "UpdatedMaxNonTerminalKeyPerOperator")
	if err != nil {
		return nil, err
	}
	return &PermissionlessNodeRegistryUpdatedMaxNonTerminalKeyPerOperatorIterator{contract: _PermissionlessNodeRegistry.contract, event: "UpdatedMaxNonTerminalKeyPerOperator", logs: logs, sub: sub}, nil
}

// WatchUpdatedMaxNonTerminalKeyPerOperator is a free log subscription operation binding the contract event 0xacda2fe79efeffc359206ddeeb45f26ba1596223e01e1585458603af76e880a2.
//
// Solidity: event UpdatedMaxNonTerminalKeyPerOperator(uint64 maxNonTerminalKeyPerOperator)
func (_PermissionlessNodeRegistry *PermissionlessNodeRegistryFilterer) WatchUpdatedMaxNonTerminalKeyPerOperator(opts *bind.WatchOpts, sink chan<- *PermissionlessNodeRegistryUpdatedMaxNonTerminalKeyPerOperator) (event.Subscription, error) {

	logs, sub, err := _PermissionlessNodeRegistry.contract.WatchLogs(opts, "UpdatedMaxNonTerminalKeyPerOperator")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PermissionlessNodeRegistryUpdatedMaxNonTerminalKeyPerOperator)
				if err := _PermissionlessNodeRegistry.contract.UnpackLog(event, "UpdatedMaxNonTerminalKeyPerOperator", log); err != nil {
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

// ParseUpdatedMaxNonTerminalKeyPerOperator is a log parse operation binding the contract event 0xacda2fe79efeffc359206ddeeb45f26ba1596223e01e1585458603af76e880a2.
//
// Solidity: event UpdatedMaxNonTerminalKeyPerOperator(uint64 maxNonTerminalKeyPerOperator)
func (_PermissionlessNodeRegistry *PermissionlessNodeRegistryFilterer) ParseUpdatedMaxNonTerminalKeyPerOperator(log types.Log) (*PermissionlessNodeRegistryUpdatedMaxNonTerminalKeyPerOperator, error) {
	event := new(PermissionlessNodeRegistryUpdatedMaxNonTerminalKeyPerOperator)
	if err := _PermissionlessNodeRegistry.contract.UnpackLog(event, "UpdatedMaxNonTerminalKeyPerOperator", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PermissionlessNodeRegistryUpdatedNextQueuedValidatorIndexIterator is returned from FilterUpdatedNextQueuedValidatorIndex and is used to iterate over the raw logs and unpacked data for UpdatedNextQueuedValidatorIndex events raised by the PermissionlessNodeRegistry contract.
type PermissionlessNodeRegistryUpdatedNextQueuedValidatorIndexIterator struct {
	Event *PermissionlessNodeRegistryUpdatedNextQueuedValidatorIndex // Event containing the contract specifics and raw log

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
func (it *PermissionlessNodeRegistryUpdatedNextQueuedValidatorIndexIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PermissionlessNodeRegistryUpdatedNextQueuedValidatorIndex)
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
		it.Event = new(PermissionlessNodeRegistryUpdatedNextQueuedValidatorIndex)
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
func (it *PermissionlessNodeRegistryUpdatedNextQueuedValidatorIndexIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PermissionlessNodeRegistryUpdatedNextQueuedValidatorIndexIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PermissionlessNodeRegistryUpdatedNextQueuedValidatorIndex represents a UpdatedNextQueuedValidatorIndex event raised by the PermissionlessNodeRegistry contract.
type PermissionlessNodeRegistryUpdatedNextQueuedValidatorIndex struct {
	NextQueuedValidatorIndex *big.Int
	Raw                      types.Log // Blockchain specific contextual infos
}

// FilterUpdatedNextQueuedValidatorIndex is a free log retrieval operation binding the contract event 0x711359152f2039f4182a096114b0d199c5f8e9cb268caff34080855c42ff4da9.
//
// Solidity: event UpdatedNextQueuedValidatorIndex(uint256 nextQueuedValidatorIndex)
func (_PermissionlessNodeRegistry *PermissionlessNodeRegistryFilterer) FilterUpdatedNextQueuedValidatorIndex(opts *bind.FilterOpts) (*PermissionlessNodeRegistryUpdatedNextQueuedValidatorIndexIterator, error) {

	logs, sub, err := _PermissionlessNodeRegistry.contract.FilterLogs(opts, "UpdatedNextQueuedValidatorIndex")
	if err != nil {
		return nil, err
	}
	return &PermissionlessNodeRegistryUpdatedNextQueuedValidatorIndexIterator{contract: _PermissionlessNodeRegistry.contract, event: "UpdatedNextQueuedValidatorIndex", logs: logs, sub: sub}, nil
}

// WatchUpdatedNextQueuedValidatorIndex is a free log subscription operation binding the contract event 0x711359152f2039f4182a096114b0d199c5f8e9cb268caff34080855c42ff4da9.
//
// Solidity: event UpdatedNextQueuedValidatorIndex(uint256 nextQueuedValidatorIndex)
func (_PermissionlessNodeRegistry *PermissionlessNodeRegistryFilterer) WatchUpdatedNextQueuedValidatorIndex(opts *bind.WatchOpts, sink chan<- *PermissionlessNodeRegistryUpdatedNextQueuedValidatorIndex) (event.Subscription, error) {

	logs, sub, err := _PermissionlessNodeRegistry.contract.WatchLogs(opts, "UpdatedNextQueuedValidatorIndex")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PermissionlessNodeRegistryUpdatedNextQueuedValidatorIndex)
				if err := _PermissionlessNodeRegistry.contract.UnpackLog(event, "UpdatedNextQueuedValidatorIndex", log); err != nil {
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

// ParseUpdatedNextQueuedValidatorIndex is a log parse operation binding the contract event 0x711359152f2039f4182a096114b0d199c5f8e9cb268caff34080855c42ff4da9.
//
// Solidity: event UpdatedNextQueuedValidatorIndex(uint256 nextQueuedValidatorIndex)
func (_PermissionlessNodeRegistry *PermissionlessNodeRegistryFilterer) ParseUpdatedNextQueuedValidatorIndex(log types.Log) (*PermissionlessNodeRegistryUpdatedNextQueuedValidatorIndex, error) {
	event := new(PermissionlessNodeRegistryUpdatedNextQueuedValidatorIndex)
	if err := _PermissionlessNodeRegistry.contract.UnpackLog(event, "UpdatedNextQueuedValidatorIndex", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PermissionlessNodeRegistryUpdatedOperatorDetailsIterator is returned from FilterUpdatedOperatorDetails and is used to iterate over the raw logs and unpacked data for UpdatedOperatorDetails events raised by the PermissionlessNodeRegistry contract.
type PermissionlessNodeRegistryUpdatedOperatorDetailsIterator struct {
	Event *PermissionlessNodeRegistryUpdatedOperatorDetails // Event containing the contract specifics and raw log

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
func (it *PermissionlessNodeRegistryUpdatedOperatorDetailsIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PermissionlessNodeRegistryUpdatedOperatorDetails)
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
		it.Event = new(PermissionlessNodeRegistryUpdatedOperatorDetails)
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
func (it *PermissionlessNodeRegistryUpdatedOperatorDetailsIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PermissionlessNodeRegistryUpdatedOperatorDetailsIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PermissionlessNodeRegistryUpdatedOperatorDetails represents a UpdatedOperatorDetails event raised by the PermissionlessNodeRegistry contract.
type PermissionlessNodeRegistryUpdatedOperatorDetails struct {
	NodeOperator  common.Address
	OperatorName  string
	RewardAddress common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterUpdatedOperatorDetails is a free log retrieval operation binding the contract event 0xadc8722095edf061d7fdcb583105c05bf9eb15488503b621c39e254d87269777.
//
// Solidity: event UpdatedOperatorDetails(address indexed nodeOperator, string operatorName, address rewardAddress)
func (_PermissionlessNodeRegistry *PermissionlessNodeRegistryFilterer) FilterUpdatedOperatorDetails(opts *bind.FilterOpts, nodeOperator []common.Address) (*PermissionlessNodeRegistryUpdatedOperatorDetailsIterator, error) {

	var nodeOperatorRule []interface{}
	for _, nodeOperatorItem := range nodeOperator {
		nodeOperatorRule = append(nodeOperatorRule, nodeOperatorItem)
	}

	logs, sub, err := _PermissionlessNodeRegistry.contract.FilterLogs(opts, "UpdatedOperatorDetails", nodeOperatorRule)
	if err != nil {
		return nil, err
	}
	return &PermissionlessNodeRegistryUpdatedOperatorDetailsIterator{contract: _PermissionlessNodeRegistry.contract, event: "UpdatedOperatorDetails", logs: logs, sub: sub}, nil
}

// WatchUpdatedOperatorDetails is a free log subscription operation binding the contract event 0xadc8722095edf061d7fdcb583105c05bf9eb15488503b621c39e254d87269777.
//
// Solidity: event UpdatedOperatorDetails(address indexed nodeOperator, string operatorName, address rewardAddress)
func (_PermissionlessNodeRegistry *PermissionlessNodeRegistryFilterer) WatchUpdatedOperatorDetails(opts *bind.WatchOpts, sink chan<- *PermissionlessNodeRegistryUpdatedOperatorDetails, nodeOperator []common.Address) (event.Subscription, error) {

	var nodeOperatorRule []interface{}
	for _, nodeOperatorItem := range nodeOperator {
		nodeOperatorRule = append(nodeOperatorRule, nodeOperatorItem)
	}

	logs, sub, err := _PermissionlessNodeRegistry.contract.WatchLogs(opts, "UpdatedOperatorDetails", nodeOperatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PermissionlessNodeRegistryUpdatedOperatorDetails)
				if err := _PermissionlessNodeRegistry.contract.UnpackLog(event, "UpdatedOperatorDetails", log); err != nil {
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

// ParseUpdatedOperatorDetails is a log parse operation binding the contract event 0xadc8722095edf061d7fdcb583105c05bf9eb15488503b621c39e254d87269777.
//
// Solidity: event UpdatedOperatorDetails(address indexed nodeOperator, string operatorName, address rewardAddress)
func (_PermissionlessNodeRegistry *PermissionlessNodeRegistryFilterer) ParseUpdatedOperatorDetails(log types.Log) (*PermissionlessNodeRegistryUpdatedOperatorDetails, error) {
	event := new(PermissionlessNodeRegistryUpdatedOperatorDetails)
	if err := _PermissionlessNodeRegistry.contract.UnpackLog(event, "UpdatedOperatorDetails", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PermissionlessNodeRegistryUpdatedSocializingPoolStateIterator is returned from FilterUpdatedSocializingPoolState and is used to iterate over the raw logs and unpacked data for UpdatedSocializingPoolState events raised by the PermissionlessNodeRegistry contract.
type PermissionlessNodeRegistryUpdatedSocializingPoolStateIterator struct {
	Event *PermissionlessNodeRegistryUpdatedSocializingPoolState // Event containing the contract specifics and raw log

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
func (it *PermissionlessNodeRegistryUpdatedSocializingPoolStateIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PermissionlessNodeRegistryUpdatedSocializingPoolState)
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
		it.Event = new(PermissionlessNodeRegistryUpdatedSocializingPoolState)
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
func (it *PermissionlessNodeRegistryUpdatedSocializingPoolStateIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PermissionlessNodeRegistryUpdatedSocializingPoolStateIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PermissionlessNodeRegistryUpdatedSocializingPoolState represents a UpdatedSocializingPoolState event raised by the PermissionlessNodeRegistry contract.
type PermissionlessNodeRegistryUpdatedSocializingPoolState struct {
	OperatorId              *big.Int
	OptedForSocializingPool bool
	Block                   *big.Int
	Raw                     types.Log // Blockchain specific contextual infos
}

// FilterUpdatedSocializingPoolState is a free log retrieval operation binding the contract event 0xc0465abaf1d51829975919c02418d521476b44f330a31d78bb6b4e96465e746b.
//
// Solidity: event UpdatedSocializingPoolState(uint256 operatorId, bool optedForSocializingPool, uint256 block)
func (_PermissionlessNodeRegistry *PermissionlessNodeRegistryFilterer) FilterUpdatedSocializingPoolState(opts *bind.FilterOpts) (*PermissionlessNodeRegistryUpdatedSocializingPoolStateIterator, error) {

	logs, sub, err := _PermissionlessNodeRegistry.contract.FilterLogs(opts, "UpdatedSocializingPoolState")
	if err != nil {
		return nil, err
	}
	return &PermissionlessNodeRegistryUpdatedSocializingPoolStateIterator{contract: _PermissionlessNodeRegistry.contract, event: "UpdatedSocializingPoolState", logs: logs, sub: sub}, nil
}

// WatchUpdatedSocializingPoolState is a free log subscription operation binding the contract event 0xc0465abaf1d51829975919c02418d521476b44f330a31d78bb6b4e96465e746b.
//
// Solidity: event UpdatedSocializingPoolState(uint256 operatorId, bool optedForSocializingPool, uint256 block)
func (_PermissionlessNodeRegistry *PermissionlessNodeRegistryFilterer) WatchUpdatedSocializingPoolState(opts *bind.WatchOpts, sink chan<- *PermissionlessNodeRegistryUpdatedSocializingPoolState) (event.Subscription, error) {

	logs, sub, err := _PermissionlessNodeRegistry.contract.WatchLogs(opts, "UpdatedSocializingPoolState")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PermissionlessNodeRegistryUpdatedSocializingPoolState)
				if err := _PermissionlessNodeRegistry.contract.UnpackLog(event, "UpdatedSocializingPoolState", log); err != nil {
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

// ParseUpdatedSocializingPoolState is a log parse operation binding the contract event 0xc0465abaf1d51829975919c02418d521476b44f330a31d78bb6b4e96465e746b.
//
// Solidity: event UpdatedSocializingPoolState(uint256 operatorId, bool optedForSocializingPool, uint256 block)
func (_PermissionlessNodeRegistry *PermissionlessNodeRegistryFilterer) ParseUpdatedSocializingPoolState(log types.Log) (*PermissionlessNodeRegistryUpdatedSocializingPoolState, error) {
	event := new(PermissionlessNodeRegistryUpdatedSocializingPoolState)
	if err := _PermissionlessNodeRegistry.contract.UnpackLog(event, "UpdatedSocializingPoolState", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PermissionlessNodeRegistryUpdatedStaderConfigIterator is returned from FilterUpdatedStaderConfig and is used to iterate over the raw logs and unpacked data for UpdatedStaderConfig events raised by the PermissionlessNodeRegistry contract.
type PermissionlessNodeRegistryUpdatedStaderConfigIterator struct {
	Event *PermissionlessNodeRegistryUpdatedStaderConfig // Event containing the contract specifics and raw log

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
func (it *PermissionlessNodeRegistryUpdatedStaderConfigIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PermissionlessNodeRegistryUpdatedStaderConfig)
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
		it.Event = new(PermissionlessNodeRegistryUpdatedStaderConfig)
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
func (it *PermissionlessNodeRegistryUpdatedStaderConfigIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PermissionlessNodeRegistryUpdatedStaderConfigIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PermissionlessNodeRegistryUpdatedStaderConfig represents a UpdatedStaderConfig event raised by the PermissionlessNodeRegistry contract.
type PermissionlessNodeRegistryUpdatedStaderConfig struct {
	StaderConfig common.Address
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterUpdatedStaderConfig is a free log retrieval operation binding the contract event 0xdb2219043d7b197cb235f1af0cf6d782d77dee3de19e3f4fb6d39aae633b4485.
//
// Solidity: event UpdatedStaderConfig(address staderConfig)
func (_PermissionlessNodeRegistry *PermissionlessNodeRegistryFilterer) FilterUpdatedStaderConfig(opts *bind.FilterOpts) (*PermissionlessNodeRegistryUpdatedStaderConfigIterator, error) {

	logs, sub, err := _PermissionlessNodeRegistry.contract.FilterLogs(opts, "UpdatedStaderConfig")
	if err != nil {
		return nil, err
	}
	return &PermissionlessNodeRegistryUpdatedStaderConfigIterator{contract: _PermissionlessNodeRegistry.contract, event: "UpdatedStaderConfig", logs: logs, sub: sub}, nil
}

// WatchUpdatedStaderConfig is a free log subscription operation binding the contract event 0xdb2219043d7b197cb235f1af0cf6d782d77dee3de19e3f4fb6d39aae633b4485.
//
// Solidity: event UpdatedStaderConfig(address staderConfig)
func (_PermissionlessNodeRegistry *PermissionlessNodeRegistryFilterer) WatchUpdatedStaderConfig(opts *bind.WatchOpts, sink chan<- *PermissionlessNodeRegistryUpdatedStaderConfig) (event.Subscription, error) {

	logs, sub, err := _PermissionlessNodeRegistry.contract.WatchLogs(opts, "UpdatedStaderConfig")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PermissionlessNodeRegistryUpdatedStaderConfig)
				if err := _PermissionlessNodeRegistry.contract.UnpackLog(event, "UpdatedStaderConfig", log); err != nil {
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

// ParseUpdatedStaderConfig is a log parse operation binding the contract event 0xdb2219043d7b197cb235f1af0cf6d782d77dee3de19e3f4fb6d39aae633b4485.
//
// Solidity: event UpdatedStaderConfig(address staderConfig)
func (_PermissionlessNodeRegistry *PermissionlessNodeRegistryFilterer) ParseUpdatedStaderConfig(log types.Log) (*PermissionlessNodeRegistryUpdatedStaderConfig, error) {
	event := new(PermissionlessNodeRegistryUpdatedStaderConfig)
	if err := _PermissionlessNodeRegistry.contract.UnpackLog(event, "UpdatedStaderConfig", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PermissionlessNodeRegistryUpdatedValidatorDepositBlockIterator is returned from FilterUpdatedValidatorDepositBlock and is used to iterate over the raw logs and unpacked data for UpdatedValidatorDepositBlock events raised by the PermissionlessNodeRegistry contract.
type PermissionlessNodeRegistryUpdatedValidatorDepositBlockIterator struct {
	Event *PermissionlessNodeRegistryUpdatedValidatorDepositBlock // Event containing the contract specifics and raw log

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
func (it *PermissionlessNodeRegistryUpdatedValidatorDepositBlockIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PermissionlessNodeRegistryUpdatedValidatorDepositBlock)
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
		it.Event = new(PermissionlessNodeRegistryUpdatedValidatorDepositBlock)
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
func (it *PermissionlessNodeRegistryUpdatedValidatorDepositBlockIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PermissionlessNodeRegistryUpdatedValidatorDepositBlockIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PermissionlessNodeRegistryUpdatedValidatorDepositBlock represents a UpdatedValidatorDepositBlock event raised by the PermissionlessNodeRegistry contract.
type PermissionlessNodeRegistryUpdatedValidatorDepositBlock struct {
	ValidatorId  *big.Int
	DepositBlock *big.Int
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterUpdatedValidatorDepositBlock is a free log retrieval operation binding the contract event 0xce479ab1b7a806fa3704c907b8fae15a191ad8da9a1671659e4f411f516c4c01.
//
// Solidity: event UpdatedValidatorDepositBlock(uint256 validatorId, uint256 depositBlock)
func (_PermissionlessNodeRegistry *PermissionlessNodeRegistryFilterer) FilterUpdatedValidatorDepositBlock(opts *bind.FilterOpts) (*PermissionlessNodeRegistryUpdatedValidatorDepositBlockIterator, error) {

	logs, sub, err := _PermissionlessNodeRegistry.contract.FilterLogs(opts, "UpdatedValidatorDepositBlock")
	if err != nil {
		return nil, err
	}
	return &PermissionlessNodeRegistryUpdatedValidatorDepositBlockIterator{contract: _PermissionlessNodeRegistry.contract, event: "UpdatedValidatorDepositBlock", logs: logs, sub: sub}, nil
}

// WatchUpdatedValidatorDepositBlock is a free log subscription operation binding the contract event 0xce479ab1b7a806fa3704c907b8fae15a191ad8da9a1671659e4f411f516c4c01.
//
// Solidity: event UpdatedValidatorDepositBlock(uint256 validatorId, uint256 depositBlock)
func (_PermissionlessNodeRegistry *PermissionlessNodeRegistryFilterer) WatchUpdatedValidatorDepositBlock(opts *bind.WatchOpts, sink chan<- *PermissionlessNodeRegistryUpdatedValidatorDepositBlock) (event.Subscription, error) {

	logs, sub, err := _PermissionlessNodeRegistry.contract.WatchLogs(opts, "UpdatedValidatorDepositBlock")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PermissionlessNodeRegistryUpdatedValidatorDepositBlock)
				if err := _PermissionlessNodeRegistry.contract.UnpackLog(event, "UpdatedValidatorDepositBlock", log); err != nil {
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

// ParseUpdatedValidatorDepositBlock is a log parse operation binding the contract event 0xce479ab1b7a806fa3704c907b8fae15a191ad8da9a1671659e4f411f516c4c01.
//
// Solidity: event UpdatedValidatorDepositBlock(uint256 validatorId, uint256 depositBlock)
func (_PermissionlessNodeRegistry *PermissionlessNodeRegistryFilterer) ParseUpdatedValidatorDepositBlock(log types.Log) (*PermissionlessNodeRegistryUpdatedValidatorDepositBlock, error) {
	event := new(PermissionlessNodeRegistryUpdatedValidatorDepositBlock)
	if err := _PermissionlessNodeRegistry.contract.UnpackLog(event, "UpdatedValidatorDepositBlock", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PermissionlessNodeRegistryUpdatedVerifiedKeyBatchSizeIterator is returned from FilterUpdatedVerifiedKeyBatchSize and is used to iterate over the raw logs and unpacked data for UpdatedVerifiedKeyBatchSize events raised by the PermissionlessNodeRegistry contract.
type PermissionlessNodeRegistryUpdatedVerifiedKeyBatchSizeIterator struct {
	Event *PermissionlessNodeRegistryUpdatedVerifiedKeyBatchSize // Event containing the contract specifics and raw log

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
func (it *PermissionlessNodeRegistryUpdatedVerifiedKeyBatchSizeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PermissionlessNodeRegistryUpdatedVerifiedKeyBatchSize)
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
		it.Event = new(PermissionlessNodeRegistryUpdatedVerifiedKeyBatchSize)
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
func (it *PermissionlessNodeRegistryUpdatedVerifiedKeyBatchSizeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PermissionlessNodeRegistryUpdatedVerifiedKeyBatchSizeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PermissionlessNodeRegistryUpdatedVerifiedKeyBatchSize represents a UpdatedVerifiedKeyBatchSize event raised by the PermissionlessNodeRegistry contract.
type PermissionlessNodeRegistryUpdatedVerifiedKeyBatchSize struct {
	VerifiedKeysBatchSize *big.Int
	Raw                   types.Log // Blockchain specific contextual infos
}

// FilterUpdatedVerifiedKeyBatchSize is a free log retrieval operation binding the contract event 0x5d19c92c6893231b764f3320c712a4d056ff157295c8b620d893dbbed1a869b4.
//
// Solidity: event UpdatedVerifiedKeyBatchSize(uint256 verifiedKeysBatchSize)
func (_PermissionlessNodeRegistry *PermissionlessNodeRegistryFilterer) FilterUpdatedVerifiedKeyBatchSize(opts *bind.FilterOpts) (*PermissionlessNodeRegistryUpdatedVerifiedKeyBatchSizeIterator, error) {

	logs, sub, err := _PermissionlessNodeRegistry.contract.FilterLogs(opts, "UpdatedVerifiedKeyBatchSize")
	if err != nil {
		return nil, err
	}
	return &PermissionlessNodeRegistryUpdatedVerifiedKeyBatchSizeIterator{contract: _PermissionlessNodeRegistry.contract, event: "UpdatedVerifiedKeyBatchSize", logs: logs, sub: sub}, nil
}

// WatchUpdatedVerifiedKeyBatchSize is a free log subscription operation binding the contract event 0x5d19c92c6893231b764f3320c712a4d056ff157295c8b620d893dbbed1a869b4.
//
// Solidity: event UpdatedVerifiedKeyBatchSize(uint256 verifiedKeysBatchSize)
func (_PermissionlessNodeRegistry *PermissionlessNodeRegistryFilterer) WatchUpdatedVerifiedKeyBatchSize(opts *bind.WatchOpts, sink chan<- *PermissionlessNodeRegistryUpdatedVerifiedKeyBatchSize) (event.Subscription, error) {

	logs, sub, err := _PermissionlessNodeRegistry.contract.WatchLogs(opts, "UpdatedVerifiedKeyBatchSize")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PermissionlessNodeRegistryUpdatedVerifiedKeyBatchSize)
				if err := _PermissionlessNodeRegistry.contract.UnpackLog(event, "UpdatedVerifiedKeyBatchSize", log); err != nil {
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

// ParseUpdatedVerifiedKeyBatchSize is a log parse operation binding the contract event 0x5d19c92c6893231b764f3320c712a4d056ff157295c8b620d893dbbed1a869b4.
//
// Solidity: event UpdatedVerifiedKeyBatchSize(uint256 verifiedKeysBatchSize)
func (_PermissionlessNodeRegistry *PermissionlessNodeRegistryFilterer) ParseUpdatedVerifiedKeyBatchSize(log types.Log) (*PermissionlessNodeRegistryUpdatedVerifiedKeyBatchSize, error) {
	event := new(PermissionlessNodeRegistryUpdatedVerifiedKeyBatchSize)
	if err := _PermissionlessNodeRegistry.contract.UnpackLog(event, "UpdatedVerifiedKeyBatchSize", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PermissionlessNodeRegistryUpdatedWithdrawnKeyBatchSizeIterator is returned from FilterUpdatedWithdrawnKeyBatchSize and is used to iterate over the raw logs and unpacked data for UpdatedWithdrawnKeyBatchSize events raised by the PermissionlessNodeRegistry contract.
type PermissionlessNodeRegistryUpdatedWithdrawnKeyBatchSizeIterator struct {
	Event *PermissionlessNodeRegistryUpdatedWithdrawnKeyBatchSize // Event containing the contract specifics and raw log

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
func (it *PermissionlessNodeRegistryUpdatedWithdrawnKeyBatchSizeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PermissionlessNodeRegistryUpdatedWithdrawnKeyBatchSize)
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
		it.Event = new(PermissionlessNodeRegistryUpdatedWithdrawnKeyBatchSize)
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
func (it *PermissionlessNodeRegistryUpdatedWithdrawnKeyBatchSizeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PermissionlessNodeRegistryUpdatedWithdrawnKeyBatchSizeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PermissionlessNodeRegistryUpdatedWithdrawnKeyBatchSize represents a UpdatedWithdrawnKeyBatchSize event raised by the PermissionlessNodeRegistry contract.
type PermissionlessNodeRegistryUpdatedWithdrawnKeyBatchSize struct {
	WithdrawnKeysBatchSize *big.Int
	Raw                    types.Log // Blockchain specific contextual infos
}

// FilterUpdatedWithdrawnKeyBatchSize is a free log retrieval operation binding the contract event 0x5aa519ec64f29fb81c513568f7c6839ee0265b5799bb434dfa467be612125950.
//
// Solidity: event UpdatedWithdrawnKeyBatchSize(uint256 withdrawnKeysBatchSize)
func (_PermissionlessNodeRegistry *PermissionlessNodeRegistryFilterer) FilterUpdatedWithdrawnKeyBatchSize(opts *bind.FilterOpts) (*PermissionlessNodeRegistryUpdatedWithdrawnKeyBatchSizeIterator, error) {

	logs, sub, err := _PermissionlessNodeRegistry.contract.FilterLogs(opts, "UpdatedWithdrawnKeyBatchSize")
	if err != nil {
		return nil, err
	}
	return &PermissionlessNodeRegistryUpdatedWithdrawnKeyBatchSizeIterator{contract: _PermissionlessNodeRegistry.contract, event: "UpdatedWithdrawnKeyBatchSize", logs: logs, sub: sub}, nil
}

// WatchUpdatedWithdrawnKeyBatchSize is a free log subscription operation binding the contract event 0x5aa519ec64f29fb81c513568f7c6839ee0265b5799bb434dfa467be612125950.
//
// Solidity: event UpdatedWithdrawnKeyBatchSize(uint256 withdrawnKeysBatchSize)
func (_PermissionlessNodeRegistry *PermissionlessNodeRegistryFilterer) WatchUpdatedWithdrawnKeyBatchSize(opts *bind.WatchOpts, sink chan<- *PermissionlessNodeRegistryUpdatedWithdrawnKeyBatchSize) (event.Subscription, error) {

	logs, sub, err := _PermissionlessNodeRegistry.contract.WatchLogs(opts, "UpdatedWithdrawnKeyBatchSize")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PermissionlessNodeRegistryUpdatedWithdrawnKeyBatchSize)
				if err := _PermissionlessNodeRegistry.contract.UnpackLog(event, "UpdatedWithdrawnKeyBatchSize", log); err != nil {
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

// ParseUpdatedWithdrawnKeyBatchSize is a log parse operation binding the contract event 0x5aa519ec64f29fb81c513568f7c6839ee0265b5799bb434dfa467be612125950.
//
// Solidity: event UpdatedWithdrawnKeyBatchSize(uint256 withdrawnKeysBatchSize)
func (_PermissionlessNodeRegistry *PermissionlessNodeRegistryFilterer) ParseUpdatedWithdrawnKeyBatchSize(log types.Log) (*PermissionlessNodeRegistryUpdatedWithdrawnKeyBatchSize, error) {
	event := new(PermissionlessNodeRegistryUpdatedWithdrawnKeyBatchSize)
	if err := _PermissionlessNodeRegistry.contract.UnpackLog(event, "UpdatedWithdrawnKeyBatchSize", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PermissionlessNodeRegistryValidatorMarkedAsFrontRunnedIterator is returned from FilterValidatorMarkedAsFrontRunned and is used to iterate over the raw logs and unpacked data for ValidatorMarkedAsFrontRunned events raised by the PermissionlessNodeRegistry contract.
type PermissionlessNodeRegistryValidatorMarkedAsFrontRunnedIterator struct {
	Event *PermissionlessNodeRegistryValidatorMarkedAsFrontRunned // Event containing the contract specifics and raw log

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
func (it *PermissionlessNodeRegistryValidatorMarkedAsFrontRunnedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PermissionlessNodeRegistryValidatorMarkedAsFrontRunned)
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
		it.Event = new(PermissionlessNodeRegistryValidatorMarkedAsFrontRunned)
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
func (it *PermissionlessNodeRegistryValidatorMarkedAsFrontRunnedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PermissionlessNodeRegistryValidatorMarkedAsFrontRunnedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PermissionlessNodeRegistryValidatorMarkedAsFrontRunned represents a ValidatorMarkedAsFrontRunned event raised by the PermissionlessNodeRegistry contract.
type PermissionlessNodeRegistryValidatorMarkedAsFrontRunned struct {
	Pubkey      []byte
	ValidatorId *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterValidatorMarkedAsFrontRunned is a free log retrieval operation binding the contract event 0x4e93215f00bc729272f0ff71afd3d0f385208cbf6c999fe776ad07c623b83466.
//
// Solidity: event ValidatorMarkedAsFrontRunned(bytes pubkey, uint256 validatorId)
func (_PermissionlessNodeRegistry *PermissionlessNodeRegistryFilterer) FilterValidatorMarkedAsFrontRunned(opts *bind.FilterOpts) (*PermissionlessNodeRegistryValidatorMarkedAsFrontRunnedIterator, error) {

	logs, sub, err := _PermissionlessNodeRegistry.contract.FilterLogs(opts, "ValidatorMarkedAsFrontRunned")
	if err != nil {
		return nil, err
	}
	return &PermissionlessNodeRegistryValidatorMarkedAsFrontRunnedIterator{contract: _PermissionlessNodeRegistry.contract, event: "ValidatorMarkedAsFrontRunned", logs: logs, sub: sub}, nil
}

// WatchValidatorMarkedAsFrontRunned is a free log subscription operation binding the contract event 0x4e93215f00bc729272f0ff71afd3d0f385208cbf6c999fe776ad07c623b83466.
//
// Solidity: event ValidatorMarkedAsFrontRunned(bytes pubkey, uint256 validatorId)
func (_PermissionlessNodeRegistry *PermissionlessNodeRegistryFilterer) WatchValidatorMarkedAsFrontRunned(opts *bind.WatchOpts, sink chan<- *PermissionlessNodeRegistryValidatorMarkedAsFrontRunned) (event.Subscription, error) {

	logs, sub, err := _PermissionlessNodeRegistry.contract.WatchLogs(opts, "ValidatorMarkedAsFrontRunned")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PermissionlessNodeRegistryValidatorMarkedAsFrontRunned)
				if err := _PermissionlessNodeRegistry.contract.UnpackLog(event, "ValidatorMarkedAsFrontRunned", log); err != nil {
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

// ParseValidatorMarkedAsFrontRunned is a log parse operation binding the contract event 0x4e93215f00bc729272f0ff71afd3d0f385208cbf6c999fe776ad07c623b83466.
//
// Solidity: event ValidatorMarkedAsFrontRunned(bytes pubkey, uint256 validatorId)
func (_PermissionlessNodeRegistry *PermissionlessNodeRegistryFilterer) ParseValidatorMarkedAsFrontRunned(log types.Log) (*PermissionlessNodeRegistryValidatorMarkedAsFrontRunned, error) {
	event := new(PermissionlessNodeRegistryValidatorMarkedAsFrontRunned)
	if err := _PermissionlessNodeRegistry.contract.UnpackLog(event, "ValidatorMarkedAsFrontRunned", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PermissionlessNodeRegistryValidatorMarkedReadyToDepositIterator is returned from FilterValidatorMarkedReadyToDeposit and is used to iterate over the raw logs and unpacked data for ValidatorMarkedReadyToDeposit events raised by the PermissionlessNodeRegistry contract.
type PermissionlessNodeRegistryValidatorMarkedReadyToDepositIterator struct {
	Event *PermissionlessNodeRegistryValidatorMarkedReadyToDeposit // Event containing the contract specifics and raw log

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
func (it *PermissionlessNodeRegistryValidatorMarkedReadyToDepositIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PermissionlessNodeRegistryValidatorMarkedReadyToDeposit)
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
		it.Event = new(PermissionlessNodeRegistryValidatorMarkedReadyToDeposit)
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
func (it *PermissionlessNodeRegistryValidatorMarkedReadyToDepositIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PermissionlessNodeRegistryValidatorMarkedReadyToDepositIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PermissionlessNodeRegistryValidatorMarkedReadyToDeposit represents a ValidatorMarkedReadyToDeposit event raised by the PermissionlessNodeRegistry contract.
type PermissionlessNodeRegistryValidatorMarkedReadyToDeposit struct {
	Pubkey      []byte
	ValidatorId *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterValidatorMarkedReadyToDeposit is a free log retrieval operation binding the contract event 0x21d79a0b22a7d5a18b9535162fe2f0580e24c042b0541a05afc298a77ddf5693.
//
// Solidity: event ValidatorMarkedReadyToDeposit(bytes pubkey, uint256 validatorId)
func (_PermissionlessNodeRegistry *PermissionlessNodeRegistryFilterer) FilterValidatorMarkedReadyToDeposit(opts *bind.FilterOpts) (*PermissionlessNodeRegistryValidatorMarkedReadyToDepositIterator, error) {

	logs, sub, err := _PermissionlessNodeRegistry.contract.FilterLogs(opts, "ValidatorMarkedReadyToDeposit")
	if err != nil {
		return nil, err
	}
	return &PermissionlessNodeRegistryValidatorMarkedReadyToDepositIterator{contract: _PermissionlessNodeRegistry.contract, event: "ValidatorMarkedReadyToDeposit", logs: logs, sub: sub}, nil
}

// WatchValidatorMarkedReadyToDeposit is a free log subscription operation binding the contract event 0x21d79a0b22a7d5a18b9535162fe2f0580e24c042b0541a05afc298a77ddf5693.
//
// Solidity: event ValidatorMarkedReadyToDeposit(bytes pubkey, uint256 validatorId)
func (_PermissionlessNodeRegistry *PermissionlessNodeRegistryFilterer) WatchValidatorMarkedReadyToDeposit(opts *bind.WatchOpts, sink chan<- *PermissionlessNodeRegistryValidatorMarkedReadyToDeposit) (event.Subscription, error) {

	logs, sub, err := _PermissionlessNodeRegistry.contract.WatchLogs(opts, "ValidatorMarkedReadyToDeposit")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PermissionlessNodeRegistryValidatorMarkedReadyToDeposit)
				if err := _PermissionlessNodeRegistry.contract.UnpackLog(event, "ValidatorMarkedReadyToDeposit", log); err != nil {
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

// ParseValidatorMarkedReadyToDeposit is a log parse operation binding the contract event 0x21d79a0b22a7d5a18b9535162fe2f0580e24c042b0541a05afc298a77ddf5693.
//
// Solidity: event ValidatorMarkedReadyToDeposit(bytes pubkey, uint256 validatorId)
func (_PermissionlessNodeRegistry *PermissionlessNodeRegistryFilterer) ParseValidatorMarkedReadyToDeposit(log types.Log) (*PermissionlessNodeRegistryValidatorMarkedReadyToDeposit, error) {
	event := new(PermissionlessNodeRegistryValidatorMarkedReadyToDeposit)
	if err := _PermissionlessNodeRegistry.contract.UnpackLog(event, "ValidatorMarkedReadyToDeposit", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PermissionlessNodeRegistryValidatorStatusMarkedAsInvalidSignatureIterator is returned from FilterValidatorStatusMarkedAsInvalidSignature and is used to iterate over the raw logs and unpacked data for ValidatorStatusMarkedAsInvalidSignature events raised by the PermissionlessNodeRegistry contract.
type PermissionlessNodeRegistryValidatorStatusMarkedAsInvalidSignatureIterator struct {
	Event *PermissionlessNodeRegistryValidatorStatusMarkedAsInvalidSignature // Event containing the contract specifics and raw log

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
func (it *PermissionlessNodeRegistryValidatorStatusMarkedAsInvalidSignatureIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PermissionlessNodeRegistryValidatorStatusMarkedAsInvalidSignature)
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
		it.Event = new(PermissionlessNodeRegistryValidatorStatusMarkedAsInvalidSignature)
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
func (it *PermissionlessNodeRegistryValidatorStatusMarkedAsInvalidSignatureIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PermissionlessNodeRegistryValidatorStatusMarkedAsInvalidSignatureIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PermissionlessNodeRegistryValidatorStatusMarkedAsInvalidSignature represents a ValidatorStatusMarkedAsInvalidSignature event raised by the PermissionlessNodeRegistry contract.
type PermissionlessNodeRegistryValidatorStatusMarkedAsInvalidSignature struct {
	Pubkey      []byte
	ValidatorId *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterValidatorStatusMarkedAsInvalidSignature is a free log retrieval operation binding the contract event 0x596ee835bed6cb827d21ba1785c468f0755ee40d33d87132df5d2ec90b645f9f.
//
// Solidity: event ValidatorStatusMarkedAsInvalidSignature(bytes pubkey, uint256 validatorId)
func (_PermissionlessNodeRegistry *PermissionlessNodeRegistryFilterer) FilterValidatorStatusMarkedAsInvalidSignature(opts *bind.FilterOpts) (*PermissionlessNodeRegistryValidatorStatusMarkedAsInvalidSignatureIterator, error) {

	logs, sub, err := _PermissionlessNodeRegistry.contract.FilterLogs(opts, "ValidatorStatusMarkedAsInvalidSignature")
	if err != nil {
		return nil, err
	}
	return &PermissionlessNodeRegistryValidatorStatusMarkedAsInvalidSignatureIterator{contract: _PermissionlessNodeRegistry.contract, event: "ValidatorStatusMarkedAsInvalidSignature", logs: logs, sub: sub}, nil
}

// WatchValidatorStatusMarkedAsInvalidSignature is a free log subscription operation binding the contract event 0x596ee835bed6cb827d21ba1785c468f0755ee40d33d87132df5d2ec90b645f9f.
//
// Solidity: event ValidatorStatusMarkedAsInvalidSignature(bytes pubkey, uint256 validatorId)
func (_PermissionlessNodeRegistry *PermissionlessNodeRegistryFilterer) WatchValidatorStatusMarkedAsInvalidSignature(opts *bind.WatchOpts, sink chan<- *PermissionlessNodeRegistryValidatorStatusMarkedAsInvalidSignature) (event.Subscription, error) {

	logs, sub, err := _PermissionlessNodeRegistry.contract.WatchLogs(opts, "ValidatorStatusMarkedAsInvalidSignature")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PermissionlessNodeRegistryValidatorStatusMarkedAsInvalidSignature)
				if err := _PermissionlessNodeRegistry.contract.UnpackLog(event, "ValidatorStatusMarkedAsInvalidSignature", log); err != nil {
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

// ParseValidatorStatusMarkedAsInvalidSignature is a log parse operation binding the contract event 0x596ee835bed6cb827d21ba1785c468f0755ee40d33d87132df5d2ec90b645f9f.
//
// Solidity: event ValidatorStatusMarkedAsInvalidSignature(bytes pubkey, uint256 validatorId)
func (_PermissionlessNodeRegistry *PermissionlessNodeRegistryFilterer) ParseValidatorStatusMarkedAsInvalidSignature(log types.Log) (*PermissionlessNodeRegistryValidatorStatusMarkedAsInvalidSignature, error) {
	event := new(PermissionlessNodeRegistryValidatorStatusMarkedAsInvalidSignature)
	if err := _PermissionlessNodeRegistry.contract.UnpackLog(event, "ValidatorStatusMarkedAsInvalidSignature", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PermissionlessNodeRegistryValidatorWithdrawnIterator is returned from FilterValidatorWithdrawn and is used to iterate over the raw logs and unpacked data for ValidatorWithdrawn events raised by the PermissionlessNodeRegistry contract.
type PermissionlessNodeRegistryValidatorWithdrawnIterator struct {
	Event *PermissionlessNodeRegistryValidatorWithdrawn // Event containing the contract specifics and raw log

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
func (it *PermissionlessNodeRegistryValidatorWithdrawnIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PermissionlessNodeRegistryValidatorWithdrawn)
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
		it.Event = new(PermissionlessNodeRegistryValidatorWithdrawn)
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
func (it *PermissionlessNodeRegistryValidatorWithdrawnIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PermissionlessNodeRegistryValidatorWithdrawnIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PermissionlessNodeRegistryValidatorWithdrawn represents a ValidatorWithdrawn event raised by the PermissionlessNodeRegistry contract.
type PermissionlessNodeRegistryValidatorWithdrawn struct {
	Pubkey      []byte
	ValidatorId *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterValidatorWithdrawn is a free log retrieval operation binding the contract event 0x450186694fefe67df6156f60235e4073b623160f28a0b85908ebc864316abf79.
//
// Solidity: event ValidatorWithdrawn(bytes pubkey, uint256 validatorId)
func (_PermissionlessNodeRegistry *PermissionlessNodeRegistryFilterer) FilterValidatorWithdrawn(opts *bind.FilterOpts) (*PermissionlessNodeRegistryValidatorWithdrawnIterator, error) {

	logs, sub, err := _PermissionlessNodeRegistry.contract.FilterLogs(opts, "ValidatorWithdrawn")
	if err != nil {
		return nil, err
	}
	return &PermissionlessNodeRegistryValidatorWithdrawnIterator{contract: _PermissionlessNodeRegistry.contract, event: "ValidatorWithdrawn", logs: logs, sub: sub}, nil
}

// WatchValidatorWithdrawn is a free log subscription operation binding the contract event 0x450186694fefe67df6156f60235e4073b623160f28a0b85908ebc864316abf79.
//
// Solidity: event ValidatorWithdrawn(bytes pubkey, uint256 validatorId)
func (_PermissionlessNodeRegistry *PermissionlessNodeRegistryFilterer) WatchValidatorWithdrawn(opts *bind.WatchOpts, sink chan<- *PermissionlessNodeRegistryValidatorWithdrawn) (event.Subscription, error) {

	logs, sub, err := _PermissionlessNodeRegistry.contract.WatchLogs(opts, "ValidatorWithdrawn")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PermissionlessNodeRegistryValidatorWithdrawn)
				if err := _PermissionlessNodeRegistry.contract.UnpackLog(event, "ValidatorWithdrawn", log); err != nil {
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

// ParseValidatorWithdrawn is a log parse operation binding the contract event 0x450186694fefe67df6156f60235e4073b623160f28a0b85908ebc864316abf79.
//
// Solidity: event ValidatorWithdrawn(bytes pubkey, uint256 validatorId)
func (_PermissionlessNodeRegistry *PermissionlessNodeRegistryFilterer) ParseValidatorWithdrawn(log types.Log) (*PermissionlessNodeRegistryValidatorWithdrawn, error) {
	event := new(PermissionlessNodeRegistryValidatorWithdrawn)
	if err := _PermissionlessNodeRegistry.contract.UnpackLog(event, "ValidatorWithdrawn", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
