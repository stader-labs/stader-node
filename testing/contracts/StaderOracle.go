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

// ExchangeRate is an auto generated low-level Go binding around an user-defined struct.
type ExchangeRate struct {
	ReportingBlockNumber *big.Int
	TotalETHBalance      *big.Int
	TotalETHXSupply      *big.Int
}

// MissedAttestationPenaltyData is an auto generated low-level Go binding around an user-defined struct.
type MissedAttestationPenaltyData struct {
	ReportingBlockNumber *big.Int
	Index                *big.Int
	SortedPubkeys        [][]byte
}



// SDPriceData is an auto generated low-level Go binding around an user-defined struct.
type SDPriceData struct {
	ReportingBlockNumber *big.Int
	SdPriceInETH         *big.Int
}

// ValidatorStats is an auto generated low-level Go binding around an user-defined struct.
type ValidatorStats struct {
	ReportingBlockNumber     *big.Int
	ExitingValidatorsBalance *big.Int
	ExitedValidatorsBalance  *big.Int
	SlashedValidatorsBalance *big.Int
	ExitingValidatorsCount   uint32
	ExitedValidatorsCount    uint32
	SlashedValidatorsCount   uint32
}

// ValidatorVerificationDetail is an auto generated low-level Go binding around an user-defined struct.
type ValidatorVerificationDetail struct {
	PoolId                        uint8
	ReportingBlockNumber          *big.Int
	SortedReadyToDepositPubkeys   [][]byte
	SortedFrontRunPubkeys         [][]byte
	SortedInvalidSignaturePubkeys [][]byte
}

// WithdrawnValidators is an auto generated low-level Go binding around an user-defined struct.
type WithdrawnValidators struct {
	PoolId               uint8
	ReportingBlockNumber *big.Int
	SortedPubkeys        [][]byte
}

// StaderOracleMetaData contains all meta data concerning the StaderOracle contract.
var StaderOracleMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_admin\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_staderConfig\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"CallerNotManager\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"CooldownNotComplete\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"DuplicateSubmissionFromNode\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ERChangeLimitCrossed\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ERChangeLimitNotCrossed\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ERPermissibleChangeOutofBounds\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"FrequencyUnchanged\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InspectionModeActive\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InsufficientTrustedNodes\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidERDataSource\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidMAPDIndex\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidMerkleRootIndex\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidPubkeyLength\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidReportingBlock\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidUpdate\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NodeAlreadyTrusted\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NodeNotTrusted\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotATrustedNode\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"PageNumberAlreadyReported\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ReportingFutureBlockData\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ReportingPreviousCycleData\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"UpdateFrequencyNotSet\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ZeroAddress\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ZeroFrequency\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"isPORBasedERData\",\"type\":\"bool\"}],\"name\":\"ERDataSourceToggled\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"erInspectionMode\",\"type\":\"bool\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"time\",\"type\":\"uint256\"}],\"name\":\"ERInspectionModeActivated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"block\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"totalEth\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"ethxSupply\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"time\",\"type\":\"uint256\"}],\"name\":\"ExchangeRateSubmitted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"block\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"totalEth\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"ethxSupply\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"time\",\"type\":\"uint256\"}],\"name\":\"ExchangeRateUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"version\",\"type\":\"uint8\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"node\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"block\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"reportingBlockNumber\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes[]\",\"name\":\"pubkeys\",\"type\":\"bytes[]\"}],\"name\":\"MissedAttestationPenaltySubmitted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"block\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes[]\",\"name\":\"pubkeys\",\"type\":\"bytes[]\"}],\"name\":\"MissedAttestationPenaltyUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Paused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"previousAdminRole\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"newAdminRole\",\"type\":\"bytes32\"}],\"name\":\"RoleAdminChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleGranted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleRevoked\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"node\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"sdPriceInETH\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"reportedBlock\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"block\",\"type\":\"uint256\"}],\"name\":\"SDPriceSubmitted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"sdPriceInETH\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"reportedBlock\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"block\",\"type\":\"uint256\"}],\"name\":\"SDPriceUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"SafeModeDisabled\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"SafeModeEnabled\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"node\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"merkleRoot\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"poolId\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"block\",\"type\":\"uint256\"}],\"name\":\"SocializingRewardsMerkleRootSubmitted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"merkleRoot\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"poolId\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"block\",\"type\":\"uint256\"}],\"name\":\"SocializingRewardsMerkleRootUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"node\",\"type\":\"address\"}],\"name\":\"TrustedNodeAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"trustedNodeChangeCoolingPeriod\",\"type\":\"uint256\"}],\"name\":\"TrustedNodeChangeCoolingPeriodUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"node\",\"type\":\"address\"}],\"name\":\"TrustedNodeRemoved\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Unpaused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"updateFrequency\",\"type\":\"uint256\"}],\"name\":\"UpdateFrequencyUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"erChangeLimit\",\"type\":\"uint256\"}],\"name\":\"UpdatedERChangeLimit\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"staderConfig\",\"type\":\"address\"}],\"name\":\"UpdatedStaderConfig\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"block\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"activeValidatorsBalance\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"exitedValidatorsBalance\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"slashedValidatorsBalance\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"activeValidatorsCount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"exitedValidatorsCount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"slashedValidatorsCount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"time\",\"type\":\"uint256\"}],\"name\":\"ValidatorStatsSubmitted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"block\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"activeValidatorsBalance\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"exitedValidatorsBalance\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"slashedValidatorsBalance\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"activeValidatorsCount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"exitedValidatorsCount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"slashedValidatorsCount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"time\",\"type\":\"uint256\"}],\"name\":\"ValidatorStatsUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"poolId\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"block\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes[]\",\"name\":\"sortedReadyToDepositPubkeys\",\"type\":\"bytes[]\"},{\"indexed\":false,\"internalType\":\"bytes[]\",\"name\":\"sortedFrontRunPubkeys\",\"type\":\"bytes[]\"},{\"indexed\":false,\"internalType\":\"bytes[]\",\"name\":\"sortedInvalidSignaturePubkeys\",\"type\":\"bytes[]\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"time\",\"type\":\"uint256\"}],\"name\":\"ValidatorVerificationDetailSubmitted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"poolId\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"block\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes[]\",\"name\":\"sortedReadyToDepositPubkeys\",\"type\":\"bytes[]\"},{\"indexed\":false,\"internalType\":\"bytes[]\",\"name\":\"sortedFrontRunPubkeys\",\"type\":\"bytes[]\"},{\"indexed\":false,\"internalType\":\"bytes[]\",\"name\":\"sortedInvalidSignaturePubkeys\",\"type\":\"bytes[]\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"time\",\"type\":\"uint256\"}],\"name\":\"ValidatorVerificationDetailUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"poolId\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"block\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes[]\",\"name\":\"pubkeys\",\"type\":\"bytes[]\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"time\",\"type\":\"uint256\"}],\"name\":\"WithdrawnValidatorsSubmitted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"poolId\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"block\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes[]\",\"name\":\"pubkeys\",\"type\":\"bytes[]\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"time\",\"type\":\"uint256\"}],\"name\":\"WithdrawnValidatorsUpdated\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"DEFAULT_ADMIN_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"ER_CHANGE_MAX_BPS\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"ETHX_ER_UF\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"MAX_ER_UPDATE_FREQUENCY\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"MIN_TRUSTED_NODES\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"MISSED_ATTESTATION_PENALTY_UF\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"SD_PRICE_UF\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"VALIDATOR_STATS_UF\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"VALIDATOR_VERIFICATION_DETAIL_UF\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"WITHDRAWN_VALIDATORS_UF\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_nodeAddress\",\"type\":\"address\"}],\"name\":\"addTrustedNode\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"closeERInspectionMode\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"disableERInspectionMode\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"disableSafeMode\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"enableSafeMode\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"erChangeLimit\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"erInspectionMode\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"erInspectionModeStartBlock\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"exchangeRate\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"reportingBlockNumber\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"totalETHBalance\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"totalETHXSupply\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"_poolId\",\"type\":\"uint8\"}],\"name\":\"getCurrentRewardsIndexByPoolId\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getERReportableBlock\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getExchangeRate\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"reportingBlockNumber\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"totalETHBalance\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"totalETHXSupply\",\"type\":\"uint256\"}],\"internalType\":\"structExchangeRate\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"_poolId\",\"type\":\"uint8\"}],\"name\":\"getMerkleRootReportableBlockByPoolId\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getMissedAttestationPenaltyReportableBlock\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"}],\"name\":\"getRoleAdmin\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getSDPriceInETH\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getSDPriceReportableBlock\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getValidatorStats\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"reportingBlockNumber\",\"type\":\"uint256\"},{\"internalType\":\"uint128\",\"name\":\"exitingValidatorsBalance\",\"type\":\"uint128\"},{\"internalType\":\"uint128\",\"name\":\"exitedValidatorsBalance\",\"type\":\"uint128\"},{\"internalType\":\"uint128\",\"name\":\"slashedValidatorsBalance\",\"type\":\"uint128\"},{\"internalType\":\"uint32\",\"name\":\"exitingValidatorsCount\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"exitedValidatorsCount\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"slashedValidatorsCount\",\"type\":\"uint32\"}],\"internalType\":\"structValidatorStats\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getValidatorStatsReportableBlock\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getValidatorVerificationDetailReportableBlock\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getWithdrawnValidatorReportableBlock\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"grantRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"hasRole\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"inspectionModeExchangeRate\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"reportingBlockNumber\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"totalETHBalance\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"totalETHXSupply\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"isPORFeedBasedERData\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"isTrustedNode\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"lastReportedMAPDIndex\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"lastReportedSDPriceData\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"reportingBlockNumber\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"sdPriceInETH\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"name\":\"lastReportingBlockNumberForValidatorVerificationDetailByPoolId\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"name\":\"lastReportingBlockNumberForWithdrawnValidatorsByPoolId\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"lastTrustedNodeCountChangeBlock\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"missedAttestationPenalty\",\"outputs\":[{\"internalType\":\"uint16\",\"name\":\"\",\"type\":\"uint16\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"paused\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_nodeAddress\",\"type\":\"address\"}],\"name\":\"removeTrustedNode\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"renounceRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"revokeRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"safeMode\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_updateFrequency\",\"type\":\"uint256\"}],\"name\":\"setERUpdateFrequency\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_updateFrequency\",\"type\":\"uint256\"}],\"name\":\"setMissedAttestationPenaltyUpdateFrequency\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_updateFrequency\",\"type\":\"uint256\"}],\"name\":\"setSDPriceUpdateFrequency\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_updateFrequency\",\"type\":\"uint256\"}],\"name\":\"setValidatorStatsUpdateFrequency\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_updateFrequency\",\"type\":\"uint256\"}],\"name\":\"setValidatorVerificationDetailUpdateFrequency\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_updateFrequency\",\"type\":\"uint256\"}],\"name\":\"setWithdrawnValidatorsUpdateFrequency\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"staderConfig\",\"outputs\":[{\"internalType\":\"contractIStaderConfig\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"reportingBlockNumber\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"totalETHBalance\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"totalETHXSupply\",\"type\":\"uint256\"}],\"internalType\":\"structExchangeRate\",\"name\":\"_exchangeRate\",\"type\":\"tuple\"}],\"name\":\"submitExchangeRateData\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"reportingBlockNumber\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"},{\"internalType\":\"bytes[]\",\"name\":\"sortedPubkeys\",\"type\":\"bytes[]\"}],\"internalType\":\"structMissedAttestationPenaltyData\",\"name\":\"_mapd\",\"type\":\"tuple\"}],\"name\":\"submitMissedAttestationPenalties\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"reportingBlockNumber\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"sdPriceInETH\",\"type\":\"uint256\"}],\"internalType\":\"structSDPriceData\",\"name\":\"_sdPriceData\",\"type\":\"tuple\"}],\"name\":\"submitSDPrice\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"reportingBlockNumber\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"merkleRoot\",\"type\":\"bytes32\"},{\"internalType\":\"uint8\",\"name\":\"poolId\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"operatorETHRewards\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"userETHRewards\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"protocolETHRewards\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"operatorSDRewards\",\"type\":\"uint256\"}],\"internalType\":\"structRewardsData\",\"name\":\"_rewardsData\",\"type\":\"tuple\"}],\"name\":\"submitSocializingRewardsMerkleRoot\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"reportingBlockNumber\",\"type\":\"uint256\"},{\"internalType\":\"uint128\",\"name\":\"exitingValidatorsBalance\",\"type\":\"uint128\"},{\"internalType\":\"uint128\",\"name\":\"exitedValidatorsBalance\",\"type\":\"uint128\"},{\"internalType\":\"uint128\",\"name\":\"slashedValidatorsBalance\",\"type\":\"uint128\"},{\"internalType\":\"uint32\",\"name\":\"exitingValidatorsCount\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"exitedValidatorsCount\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"slashedValidatorsCount\",\"type\":\"uint32\"}],\"internalType\":\"structValidatorStats\",\"name\":\"_validatorStats\",\"type\":\"tuple\"}],\"name\":\"submitValidatorStats\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint8\",\"name\":\"poolId\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"reportingBlockNumber\",\"type\":\"uint256\"},{\"internalType\":\"bytes[]\",\"name\":\"sortedReadyToDepositPubkeys\",\"type\":\"bytes[]\"},{\"internalType\":\"bytes[]\",\"name\":\"sortedFrontRunPubkeys\",\"type\":\"bytes[]\"},{\"internalType\":\"bytes[]\",\"name\":\"sortedInvalidSignaturePubkeys\",\"type\":\"bytes[]\"}],\"internalType\":\"structValidatorVerificationDetail\",\"name\":\"_validatorVerificationDetail\",\"type\":\"tuple\"}],\"name\":\"submitValidatorVerificationDetail\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint8\",\"name\":\"poolId\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"reportingBlockNumber\",\"type\":\"uint256\"},{\"internalType\":\"bytes[]\",\"name\":\"sortedPubkeys\",\"type\":\"bytes[]\"}],\"internalType\":\"structWithdrawnValidators\",\"name\":\"_withdrawnValidators\",\"type\":\"tuple\"}],\"name\":\"submitWithdrawnValidators\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"togglePORFeedBasedERData\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"trustedNodeChangeCoolingPeriod\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"trustedNodesCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"unpause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_erChangeLimit\",\"type\":\"uint256\"}],\"name\":\"updateERChangeLimit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"updateERFromPORFeed\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"updateFrequencyMap\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_staderConfig\",\"type\":\"address\"}],\"name\":\"updateStaderConfig\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_trustedNodeChangeCoolingPeriod\",\"type\":\"uint256\"}],\"name\":\"updateTrustedNodeChangeCoolingPeriod\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"validatorStats\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"reportingBlockNumber\",\"type\":\"uint256\"},{\"internalType\":\"uint128\",\"name\":\"exitingValidatorsBalance\",\"type\":\"uint128\"},{\"internalType\":\"uint128\",\"name\":\"exitedValidatorsBalance\",\"type\":\"uint128\"},{\"internalType\":\"uint128\",\"name\":\"slashedValidatorsBalance\",\"type\":\"uint128\"},{\"internalType\":\"uint32\",\"name\":\"exitingValidatorsCount\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"exitedValidatorsCount\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"slashedValidatorsCount\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Bin: "0x608060405234801562000010575f80fd5b5060405162004d9338038062004d9383398101604081905262000033916200066e565b5f54610100900460ff16158080156200005257505f54600160ff909116105b806200006d5750303b1580156200006d57505f5460ff166001145b620000d65760405162461bcd60e51b815260206004820152602e60248201527f496e697469616c697a61626c653a20636f6e747261637420697320616c72656160448201526d191e481a5b9a5d1a585b1a5e995960921b60648201526084015b60405180910390fd5b5f805460ff191660011790558015620000f8575f805461ff0019166101001790555b6200010383620002f9565b6200010e82620002f9565b6200011862000324565b6200012262000380565b6200012c620003e4565b6101f461010855620001617f783e3ebf40ee443ac9cbca8dc88c9f47450598583c2168f0ae0021de08ad333b611c2062000448565b6200018f7f8ec4e223bb52129c3d652c6e55a137389860823d9a02acb9d051e591994c6d6f611c2062000448565b620001bd7f7607f5053d01557adb731d4aad009009bba2bf77a5b5f779733919451d426336611c2062000448565b620001eb7f9ae142790790fc18374cd6a6cc28573ecc78841658523afa63055cea42a9e1dd61384062000448565b620002197fedb5588a851185ccd926df348aee898122cd3e827fb7020e3c966fdac62a46af61c4e062000448565b620002477f6a7b51c29672c0ed412394a3e65ab758361d7c963b8657ce8c1f43dc0770b162611c2062000448565b60fe80546001600160a01b0319166001600160a01b0384161790556200026e5f84620004e8565b6040516001600160a01b03831681527fdb2219043d7b197cb235f1af0cf6d782d77dee3de19e3f4fb6d39aae633b44859060200160405180910390a18015620002f0575f805461ff0019169055604051600181527f7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb38474024989060200160405180910390a15b505050620006a4565b6001600160a01b038116620003215760405163d92e233d60e01b815260040160405180910390fd5b50565b5f54610100900460ff166200037e5760405162461bcd60e51b815260206004820152602b60248201525f8051602062004d7383398151915260448201526a6e697469616c697a696e6760a81b6064820152608401620000cd565b565b5f54610100900460ff16620003da5760405162461bcd60e51b815260206004820152602b60248201525f8051602062004d7383398151915260448201526a6e697469616c697a696e6760a81b6064820152608401620000cd565b6200037e6200058b565b5f54610100900460ff166200043e5760405162461bcd60e51b815260206004820152602b60248201525f8051602062004d7383398151915260448201526a6e697469616c697a696e6760a81b6064820152608401620000cd565b6200037e620005f1565b805f036200046957604051637036cfc960e11b815260040160405180910390fd5b5f82815261011660205260409020548103620004985760405163806e577f60e01b815260040160405180910390fd5b5f828152610116602052604090819020829055517f6231a3e049e2072a042ae727208e7650b487871f4080458371e84d6e7d39113890620004dc9083815260200190565b60405180910390a15050565b5f8281526065602090815260408083206001600160a01b038516845290915290205460ff1662000587575f8281526065602090815260408083206001600160a01b03851684529091529020805460ff19166001179055620005463390565b6001600160a01b0316816001600160a01b0316837f2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d60405160405180910390a45b5050565b5f54610100900460ff16620005e55760405162461bcd60e51b815260206004820152602b60248201525f8051602062004d7383398151915260448201526a6e697469616c697a696e6760a81b6064820152608401620000cd565b6097805460ff19169055565b5f54610100900460ff166200064b5760405162461bcd60e51b815260206004820152602b60248201525f8051602062004d7383398151915260448201526a6e697469616c697a696e6760a81b6064820152608401620000cd565b600160c955565b80516001600160a01b038116811462000669575f80fd5b919050565b5f806040838503121562000680575f80fd5b6200068b8362000652565b91506200069b6020840162000652565b90509250929050565b6146c180620006b25f395ff3fe608060405234801561000f575f80fd5b5060043610610422575f3560e01c80639773ee601161022c578063c06a620111610135578063e36c3140116100bf578063f00e022311610084578063f00e022314610a95578063f10b256914610ab5578063f51c0fe714610abe578063f6a3c09014610ad1578063fc8b821c14610ae4575f80fd5b8063e36c314014610a34578063e514fe5514610a3c578063e61befa714610a44578063e6aa216c14610a58578063ea18568b14610a82575f80fd5b8063d6275dd711610105578063d6275dd7146109df578063de271c6d146109f2578063e0bcb378146109fc578063e10025e614610a05578063e2f6339214610a0d575f80fd5b8063c06a620114610993578063d06628ed146109a6578063d0a8f679146109b9578063d547741f146109cc575f80fd5b8063a71b3907116101b6578063ae815a0411610186578063ae815a0414610939578063b17b4d8614610943578063b5c25ba614610963578063b940a0031461096b578063be48e58d1461097f575f80fd5b8063a71b3907146108ed578063a8c3a3a8146108f5578063abe3219c14610918578063ae541d6514610926575f80fd5b8063a0c54387116101fc578063a0c54387146108ae578063a217fddf146108c1578063a220c2d3146108c8578063a3737869146108db578063a6870e5b146108e5575f80fd5b80639773ee601461085257806397a3a10a146108895780639bfdf9a4146108935780639ee804cb1461089b575f80fd5b80634f560abd1161032e5780637150bc5b116102b8578063844007fe11610288578063844007fe146107fe5780638456cb59146108115780638ca8703c1461081957806391d148541461082c578063962c1e051461083f575f80fd5b80637150bc5b146107a5578063735efb96146107c5578063749f7d8a146107d8578063818c8b26146107eb575f80fd5b80635c975abb116102fe5780635c975abb1461076d578063615a02531461077857806361f00c171461078057806367fbf7311461078a578063712033eb1461079d575f80fd5b80634f560abd146107375780635063b5bd1461073f57806352e0fc80146107475780635c7ccd3b1461075a575f80fd5b80632f739b1d116103af5780633ba0b9a91161037f5780633ba0b9a9146105aa5780633e23a827146105da5780633f4ba83a146106fc578063490ffa351461070457806349115a2e1461072f575f80fd5b80632f739b1d14610506578063342280b31461052957806336568abe146105365780633b5eb03a14610549575f80fd5b806312710361116103f5578063127103611461049757806316515428146104ab578063248a9ca3146104bd57806329f96856146104df5780632f2ff15d146104f3575f80fd5b806301ffc9a714610426578063052a68401461044e5780630989001c14610483578063101b6e341461048d575b5f80fd5b610439610434366004613be2565b610aec565b60405190151581526020015b60405180910390f35b6104757fedb5588a851185ccd926df348aee898122cd3e827fb7020e3c966fdac62a46af81565b604051908152602001610445565b61047561010d5481565b610495610b22565b005b6104755f8051602061460c83398151915281565b60fb5461043990610100900460ff1681565b6104756104cb366004613c09565b5f9081526065602052604090206001015490565b6104755f8051602061466c83398151915281565b610495610501366004613c34565b610b6d565b610439610514366004613c62565b61010f6020525f908152604090205460ff1681565b60fb546104399060ff1681565b610495610544366004613c34565b610b96565b61010554610106546101075461059792916001600160801b0380821692600160801b928390048216929181169163ffffffff908204811691600160a01b8104821691600160c01b9091041687565b6040516104459796959493929190613c7d565b6101025461010354610104546105bf92919083565b60408051938452602084019290925290820152606001610445565b6106886040805160e0810182525f80825260208201819052918101829052606081018290526080810182905260a0810182905260c0810191909152506040805160e081018252610105548152610106546001600160801b038082166020840152600160801b9182900481169383019390935261010754928316606083015263ffffffff90830481166080830152600160a01b8304811660a0830152600160c01b90920490911660c082015290565b60405161044591905f60e0820190508251825260208301516001600160801b0380821660208501528060408601511660408501528060608601511660608501525050608083015163ffffffff80821660808501528060a08601511660a08501528060c08601511660c0850152505092915050565b610495610c19565b60fe54610717906001600160a01b031681565b6040516001600160a01b039091168152602001610445565b610475610c2e565b610495610c4a565b610475610c99565b610495610755366004613c62565b610cb0565b610495610768366004613cc1565b610daa565b60975460ff16610439565b610475611191565b6104756101095481565b610495610798366004613ce7565b6111a8565b610495611491565b6104756107b3366004613c09565b6101166020525f908152604090205481565b6104956107d3366004613d21565b61152e565b6104956107e6366004613c09565b6119e9565b6104956107f9366004613d58565b611a2a565b61049561080c366004613c09565b611c75565b610495611ca3565b610495610827366004613c09565b611cc2565b61043961083a366004613c34565b611d42565b61049561084d366004613c09565b611d6c565b610876610860366004613c09565b6101126020525f908152604090205461ffff1681565b60405161ffff9091168152602001610445565b6104756101085481565b610495611db9565b6104956108a9366004613c62565b611e28565b6104756108bc366004613d87565b611e91565b6104755f81565b6104956108d6366004613ce7565b611fda565b61047561010b5481565b60fd54610475565b610475612406565b60fc5460fd54610903919082565b60408051928352602083019190915201610445565b61010e546104399060ff1681565b610495610934366004613da0565b612430565b61047561010a5481565b610475610951366004613d87565b6101146020525f908152604090205481565b610475612870565b60ff5461010054610101546105bf92919083565b6104755f8051602061462c83398151915281565b6104956109a1366004613c09565b612887565b6104756109b4366004613d87565b6128b5565b6104956109c7366004613c09565b6129ea565b6104956109da366004613c34565b612a3b565b6104956109ed366004613c62565b612a5f565b61047561010c5481565b61047561271081565b610495612b5d565b6104757f8ec4e223bb52129c3d652c6e55a137389860823d9a02acb9d051e591994c6d6f81565b610475600581565b610495612c15565b6104755f8051602061464c83398151915281565b610a60612c55565b6040805182518152602080840151908201529181015190820152606001610445565b610495610a90366004613c09565b612c9e565b610475610aa3366004613d87565b6101136020525f908152604090205481565b61047561c4e081565b610495610acc366004613c09565b612ccc565b610495610adf366004613db1565b612d0d565b610475612f21565b5f6001600160e01b03198216637965db0b60e01b1480610b1c57506301ffc9a760e01b6001600160e01b03198316145b92915050565b610b2a612f4b565b60fb5460ff16610b4d5760405163b1df7eb360e01b815260040160405180910390fd5b610b55612b5d565b610100546101015460ff54610b6b929190612f91565b565b5f82815260656020526040902060010154610b8781612fef565b610b918383612ff9565b505050565b6001600160a01b0381163314610c0b5760405162461bcd60e51b815260206004820152602f60248201527f416363657373436f6e74726f6c3a2063616e206f6e6c792072656e6f756e636560448201526e103937b632b9903337b91039b2b63360891b60648201526084015b60405180910390fd5b610c15828261307e565b5050565b5f610c2381612fef565b610c2b6130e4565b50565b5f610c455f8051602061466c833981519152613131565b905090565b60fe54610c619033906001600160a01b031661317c565b61010e805460ff191660011790556040517f3328bda355014adfb66d5d9086ab2c3204d1af7f83a69a3276daeed721f6ce3c905f90a1565b5f610c455f8051602061464c833981519152613131565b60fe54610cc79033906001600160a01b031661317c565b610cd081613201565b6001600160a01b0381165f90815261010f602052604090205460ff16610d095760405163f9644b0760e01b815260040160405180910390fd5b6101095461010d54610d1b9190613dd5565b431015610d3b5760405163111bb2f160e31b815260040160405180910390fd5b4361010d556001600160a01b0381165f90815261010f60205260408120805460ff1916905561010a805491610d6f83613de8565b90915550506040516001600160a01b038216907f6d1de2fb0c63996bae7ba6277c720c0560ba42874ea34c1ebe8e1423b9b47421905f90a250565b335f90815261010f602052604090205460ff16610dda57604051633e2100a160e01b815260040160405180910390fd5b600561010a541015610dff5760405163dfffd22f60e01b815260040160405180910390fd5b610e07612f4b565b43813510610e2857604051633bb0413f60e01b815260040160405180910390fd5b5f8051602061466c8339815191525f9081526101166020527f0683174ee47ba7ded338389c4047f439d06240e01796e9941d2e1ad04002de1b54610e6d908335613e11565b1115610e8c5760405163222ea98560e01b815260040160405180910390fd5b5f338235610ea06040850160208601613e38565b610eb06060860160408701613e38565b610ec06080870160608801613e38565b610ed060a0880160808901613e64565b610ee060c0890160a08a01613e64565b610ef060e08a0160c08b01613e64565b604080516001600160a01b0390991660208a01528801969096526001600160801b0394851660608801529284166080870152921660a085015263ffffffff91821660c0850152811660e084015216610100820152610120016040516020818303038152906040528051906020012090505f825f0135836020016020810190610f789190613e38565b610f886060860160408701613e38565b610f986080870160608801613e38565b610fa860a0880160808901613e64565b610fb860c0890160a08a01613e64565b610fc860e08a0160c08b01613e64565b604051602001610fde9796959493929190613c7d565b6040516020818303038152906040528051906020012090505f6110018383613228565b9050337f72745b0271618e5d84d738ea416e3a3be6eb267e0f639198f63c0ef124c29ffc85356110376040880160208901613e38565b6110476060890160408a01613e38565b61105760808a0160608b01613e38565b61106760a08b0160808c01613e64565b61107760c08c0160a08d01613e64565b61108760e08d0160c08e01613e64565b4260405161109c989796959493929190613e7f565b60405180910390a2600261010a546110b49190613ecc565b6110bf906001613dd5565b8160ff16101580156110d45750610105548435115b1561118b57836101056110e78282613eeb565b507f6988248fd82a7ce842fbdce0c49889ad926bf1548bae4194de0006498d069c949050843561111d6040870160208801613e38565b61112d6060880160408901613e38565b61113d6080890160608a01613e38565b61114d60a08a0160808b01613e64565b61115d60c08b0160a08c01613e64565b61116d60e08c0160c08d01613e64565b42604051611182989796959493929190613e7f565b60405180910390a15b50505050565b5f610c455f8051602061462c833981519152613131565b335f90815261010f602052604090205460ff166111d857604051633e2100a160e01b815260040160405180910390fd5b600561010a5410156111fd5760405163dfffd22f60e01b815260040160405180910390fd5b611205612f4b565b4381351061122657604051633bb0413f60e01b815260040160405180910390fd5b61122e612406565b81351461124e5760405163222ea98560e01b815260040160405180910390fd5b61010b5461125d906001613dd5565b8160200135146112805760405163b59f801760e01b815260040160405180910390fd5b5f61128e6040830183614005565b60405160200161129f929190614107565b60405160208183030381529060405290505f338360200135836040516020016112ca93929190614167565b6040516020818303038152906040528051906020012090505f8360200135836040516020016112fa92919061418d565b6040516020818303038152906040528051906020012090505f61131d8383613228565b9050337f51308cad1da8fe95d4be43112c17d5651d3e3713a675ec61c2214fa16d9f6beb602087013543883561135660408b018b614005565b6040516113679594939291906141a5565b60405180910390a2600261010a5461137f9190613ecc565b61138a906001613dd5565b8160ff161061148a57602085013561010b555f6113aa6040870187614005565b905090505f5b8181101561143b575f6113f06113c960408a018a614005565b848181106113d9576113d96141d5565b90506020028101906113eb91906141e9565b6132c2565b5f81815261011260205260408120805492935061ffff90921691906114148361422c565b91906101000a81548161ffff021916908361ffff16021790555050816001019150506113b0565b507f5454855cf2eeb89296b9e10ba0881425fa305f06ce9ccb1b0ce47bc2f107a19160208701354361147060408a018a614005565b604051611480949392919061424c565b60405180910390a1505b5050505050565b60fb5460ff16156114b557604051632178bc4d60e11b815260040160405180910390fd5b60fe546114cc9033906001600160a01b031661317c565b60fb805460ff610100808304821615810261ff001990931692909217928390556040517fc59a5de02f9d69be770ff0d61bbc894968433bb599f9fd9c2016e149c509c5e5936115249390049091161515815260200190565b60405180910390a1565b611536613354565b335f90815261010f602052604090205460ff1661156657604051633e2100a160e01b815260040160405180910390fd5b600561010a54101561158b5760405163dfffd22f60e01b815260040160405180910390fd5b611593612f4b565b438160200135106115b757604051633bb0413f60e01b815260040160405180910390fd5b5f8051602061462c8339815191525f90815261011660209081527f7a641f2a170436cb9ff0edd342e448ed4a6e9ee295946f1b627f5ee896014a73546115ff91840135613e11565b111561161e5760405163222ea98560e01b815260040160405180910390fd5b5f61162c6040830183614005565b6116396060850185614005565b6116466080870187614005565b60405160200161165b9695949392919061426b565b60408051601f1981840301815291905290505f3361167c6020850185613d87565b84602001358460405160200161169594939291906142b3565b60408051601f19818403018152919052805160209182012091505f906116bd90850185613d87565b8460200135846040516020016116d5939291906142e2565b6040516020818303038152906040528051906020012090505f6116f88383613228565b9050337f9827231af99e07dd2167998d4c6855a2f78e0eead80a6a490393b1c3ead048a16117296020880188613d87565b602088013561173b60408a018a614005565b61174860608c018c614005565b61175560808e018e614005565b4260405161176b99989796959493929190614303565b60405180910390a2600261010a546117839190613ecc565b61178e906001613dd5565b8160ff16101580156117c757506101145f6117ac6020880188613d87565b60ff1660ff1681526020019081526020015f20548560200135115b156119db5760208501803590610114905f906117e39089613d87565b60ff16815260208082019290925260409081015f209290925560fe5482516306ccb9d760e41b815292516001600160a01b0390911692636ccb9d709260048083019391928290030181865afa15801561183e573d5f803e3d5ffd5b505050506040513d601f19601f820116820180604052508101906118629190614364565b6001600160a01b03166399d055c861187d6020880188613d87565b6040516001600160e01b031960e084901b16815260ff9091166004820152602401602060405180830381865afa1580156118b9573d5f803e3d5ffd5b505050506040513d601f19601f820116820180604052508101906118dd9190614364565b6001600160a01b03166313797bff6118f86040880188614005565b61190560608a018a614005565b61191260808c018c614005565b6040518763ffffffff1660e01b81526004016119339695949392919061426b565b5f604051808303815f87803b15801561194a575f80fd5b505af115801561195c573d5f803e3d5ffd5b507f64a4ab6f7a6ca6fe9c5750207ea4a9d3e2d5f5ba81f707146b7b20624f61cf1192506119909150506020870187613d87565b60208701356119a26040890189614005565b6119af60608b018b614005565b6119bc60808d018d614005565b426040516119d299989796959493929190614303565b60405180910390a15b50505050610c2b600160c955565b60fe54611a009033906001600160a01b031661317c565b610c2b7f8ec4e223bb52129c3d652c6e55a137389860823d9a02acb9d051e591994c6d6f826133ad565b335f90815261010f602052604090205460ff16611a5a57604051633e2100a160e01b815260040160405180910390fd5b600561010a541015611a7f5760405163dfffd22f60e01b815260040160405180910390fd5b60fb5460ff1615611aa357604051632178bc4d60e11b815260040160405180910390fd5b611aab612f4b565b60fb54610100900460ff1615611ad35760405162ff240360e21b815260040160405180910390fd5b43813510611af457604051633bb0413f60e01b815260040160405180910390fd5b5f8051602061460c8339815191525f9081526101166020527f78e40c6661d84c085b652d9fa30921a229e88abd691be104ff2436753fe240c454611b39908335613e11565b1115611b585760405163222ea98560e01b815260040160405180910390fd5b604080513360208083019190915283358284015283013560608201529082013560808201525f9060a00160408051808303601f190181528282528051602091820120853582850152908501358383015290840135606083015291505f906080016040516020818303038152906040528051906020012090505f611bdb8383613228565b6040805186358152602080880135908201528682013581830152426060820152905191925033917f73327a5c0fdb3104b4a0f993bc20ce59885ac5bfe5f23e4bfdd19c21fb79cb129181900360800190a2600261010a54611c3c9190613ecc565b611c47906001613dd5565b8160ff1610158015611c5c5750610102548435115b1561118b5761118b60208501356040860135863561343e565b60fe54611c8c9033906001600160a01b031661317c565b610c2b5f8051602061466c833981519152826133ad565b60fe54611cba9033906001600160a01b031661317c565b610b6b613551565b60fe54611cd99033906001600160a01b031661317c565b801580611ce7575061271081115b15611d055760405163b14f197760e01b815260040160405180910390fd5b6101088190556040518181527f94a97bfc9c7a83fe5f75c66931ca7d2d16372fdc244afc5db36044f3ca52a520906020015b60405180910390a150565b5f9182526065602090815260408084206001600160a01b0393909316845291905290205460ff1690565b60fe54611d839033906001600160a01b031661317c565b6101098190556040518181527f4ab6bf3c94e4c92b7b93e89e984ef66d28392f440a58d91d244b6c303e85f04690602001611d37565b60fb5460ff1615611ddd57604051632178bc4d60e11b815260040160405180910390fd5b611de5612f4b565b60fb54610100900460ff16611e0c5760405162ff240360e21b815260040160405180910390fd5b5f805f611e1761358e565b925092509250610b9183838361343e565b5f611e3281612fef565b611e3b82613201565b60fe80546001600160a01b0319166001600160a01b0384169081179091556040519081527fdb2219043d7b197cb235f1af0cf6d782d77dee3de19e3f4fb6d39aae633b4485906020015b60405180910390a15050565b5f8060fe5f9054906101000a90046001600160a01b03166001600160a01b0316636ccb9d706040518163ffffffff1660e01b8152600401602060405180830381865afa158015611ee3573d5f803e3d5ffd5b505050506040513d601f19601f82011682018060405250810190611f079190614364565b604051637526d42960e01b815260ff851660048201526001600160a01b039190911690637526d42990602401602060405180830381865afa158015611f4e573d5f803e3d5ffd5b505050506040513d601f19601f82011682018060405250810190611f729190614364565b6001600160a01b031663d0c402766040518163ffffffff1660e01b8152600401606060405180830381865afa158015611fad573d5f803e3d5ffd5b505050506040513d601f19601f82011682018060405250810190611fd1919061437f565b95945050505050565b611fe2613354565b335f90815261010f602052604090205460ff1661201257604051633e2100a160e01b815260040160405180910390fd5b600561010a5410156120375760405163dfffd22f60e01b815260040160405180910390fd5b61203f612f4b565b4381602001351061206357604051633bb0413f60e01b815260040160405180910390fd5b5f8051602061464c8339815191525f90815261011660209081527fbc6eadc409e9002a7c49dda55566447da97ef8d5367e522b12bf4559b3c929c8546120ab91840135613e11565b11156120ca5760405163222ea98560e01b815260040160405180910390fd5b5f6120d86040830183614005565b6040516020016120e9929190614107565b60408051601f1981840301815291905290505f3361210a6020850185613d87565b84602001358460405160200161212394939291906142b3565b60408051601f19818403018152919052805160209182012091505f9061214b90850185613d87565b846020013584604051602001612163939291906142e2565b6040516020818303038152906040528051906020012090505f6121868383613228565b9050337f3b426b614a89830a3d92d8dead18e2ded3cba56101dbff12dddfc1c77b21fbe66121b76020880188613d87565b60208801356121c960408a018a614005565b426040516121db9594939291906143aa565b60405180910390a2600261010a546121f39190613ecc565b6121fe906001613dd5565b8160ff161015801561223757506101135f61221c6020880188613d87565b60ff1660ff1681526020019081526020015f20548560200135115b156119db5760208501803590610113905f906122539089613d87565b60ff16815260208082019290925260409081015f209290925560fe5482516306ccb9d760e41b815292516001600160a01b0390911692636ccb9d709260048083019391928290030181865afa1580156122ae573d5f803e3d5ffd5b505050506040513d601f19601f820116820180604052508101906122d29190614364565b6001600160a01b03166399d055c86122ed6020880188613d87565b6040516001600160e01b031960e084901b16815260ff9091166004820152602401602060405180830381865afa158015612329573d5f803e3d5ffd5b505050506040513d601f19601f8201168201806040525081019061234d9190614364565b6001600160a01b031663264f27f36123686040880188614005565b6040518363ffffffff1660e01b8152600401612385929190614107565b5f604051808303815f87803b15801561239c575f80fd5b505af11580156123ae573d5f803e3d5ffd5b507f5209293842928a1567d714f34fed8d87769d89d29dfb20f48ea678b02337e84d92506123e29150506020870187613d87565b60208701356123f46040890189614005565b426040516119d29594939291906143aa565b5f610c457fedb5588a851185ccd926df348aee898122cd3e827fb7020e3c966fdac62a46af613131565b612438613354565b335f90815261010f602052604090205460ff1661246857604051633e2100a160e01b815260040160405180910390fd5b600561010a54101561248d5760405163dfffd22f60e01b815260040160405180910390fd5b612495612f4b565b438135106124b657604051633bb0413f60e01b815260040160405180910390fd5b6124c96108bc6080830160608401613d87565b8135146124e95760405163222ea98560e01b815260040160405180910390fd5b6124fc6109b46080830160608401613d87565b81602001351461251f5760405163b4bf916f60e01b815260040160405180910390fd5b5f336020830135604084013561253b6080860160608701613d87565b604080516001600160a01b039095166020860152840192909252606083015260ff1660808281019190915283013560a08281019190915283013560c08281019190915283013560e0828101919091528301356101008201526101200160408051601f19818403018152918152815160209283012092505f91840135908401356125ca6080860160608701613d87565b60408051602081019490945283019190915260ff1660608201526080808501359082015260a0808501359082015260c0808501359082015260e080850135908201526101000160408051601f198184030181529181528151602092830120925033917f97f29f2f9a3ad2e8ebffd3fb4a6dbf5035b92b432c8344609b8368407dd233779190860135908601356126666080880160608901613d87565b60408051938452602084019290925260ff169082015243606082015260800160405180910390a25f6126988383613228565b9050600261010a546126aa9190613ecc565b6126b5906001613dd5565b8160ff16106128635760fe54604080516306ccb9d760e41b815290515f926001600160a01b031691636ccb9d709160048083019260209291908290030181865afa158015612705573d5f803e3d5ffd5b505050506040513d601f19601f820116820180604052508101906127299190614364565b6001600160a01b0316637526d4296127476080880160608901613d87565b6040516001600160e01b031960e084901b16815260ff9091166004820152602401602060405180830381865afa158015612783573d5f803e3d5ffd5b505050506040513d601f19601f820116820180604052508101906127a79190614364565b604051630d83e4ed60e01b81529091506001600160a01b03821690630d83e4ed906127d69088906004016143de565b5f604051808303815f87803b1580156127ed575f80fd5b505af11580156127ff573d5f803e3d5ffd5b507f4394ee7a4ca89453c6900058c69bfaf14014d9fc4d510ead54ae47ac06d1f05e925050506020860135604087013561283f6080890160608a01613d87565b60408051938452602084019290925260ff16908201524360608201526080016119d2565b505050610c2b600160c955565b5f610c455f8051602061460c833981519152613131565b60fe5461289e9033906001600160a01b031661317c565b610c2b5f8051602061464c833981519152826133ad565b60fe54604080516306ccb9d760e41b815290515f926001600160a01b031691636ccb9d709160048083019260209291908290030181865afa1580156128fc573d5f803e3d5ffd5b505050506040513d601f19601f820116820180604052508101906129209190614364565b604051637526d42960e01b815260ff841660048201526001600160a01b039190911690637526d42990602401602060405180830381865afa158015612967573d5f803e3d5ffd5b505050506040513d601f19601f8201168201806040525081019061298b9190614364565b6001600160a01b031663189956a26040518163ffffffff1660e01b8152600401602060405180830381865afa1580156129c6573d5f803e3d5ffd5b505050506040513d601f19601f82011682018060405250810190610b1c919061443d565b60fe54612a019033906001600160a01b031661317c565b61c4e0811115612a2457604051637d5ba07f60e01b815260040160405180910390fd5b610c2b5f8051602061460c833981519152826133ad565b5f82815260656020526040902060010154612a5581612fef565b610b91838361307e565b60fe54612a769033906001600160a01b031661317c565b612a7f81613201565b6001600160a01b0381165f90815261010f602052604090205460ff1615612ab957604051631adb013360e11b815260040160405180910390fd5b6101095461010d54612acb9190613dd5565b431015612aeb5760405163111bb2f160e31b815260040160405180910390fd5b4361010d556001600160a01b0381165f90815261010f60205260408120805460ff1916600117905561010a805491612b2283614454565b90915550506040516001600160a01b038216907fff4a290f0500d113133708d378eb9a151c32d91cb8f5778cfda6328d89cfc4b8905f90a250565b612b65612f4b565b60fe546040516318903ee760e21b81523360048201526001600160a01b0390911690636240fb9c90602401602060405180830381865afa158015612bab573d5f803e3d5ffd5b505050506040513d601f19601f82011682018060405250810190612bcf919061446c565b158015612beb57504361c4e061010c54612be99190613dd5565b115b15612c095760405163111bb2f160e31b815260040160405180910390fd5b60fb805460ff19169055565b5f612c1f81612fef565b61010e805460ff191690556040517ff29e751b3c28b619a215d25fee98a7828471a8e554626186d3f8d122f1650292905f90a150565b612c7660405180606001604052805f81526020015f81526020015f81525090565b5060408051606081018252610102548152610103546020820152610104549181019190915290565b60fe54612cb59033906001600160a01b031661317c565b610c2b5f8051602061462c833981519152826133ad565b60fe54612ce39033906001600160a01b031661317c565b610c2b7fedb5588a851185ccd926df348aee898122cd3e827fb7020e3c966fdac62a46af826133ad565b335f90815261010f602052604090205460ff16612d3d57604051633e2100a160e01b815260040160405180910390fd5b600561010a541015612d625760405163dfffd22f60e01b815260040160405180910390fd5b43813510612d8357604051633bb0413f60e01b815260040160405180910390fd5b612d8b612f21565b813514612dab5760405163222ea98560e01b815260040160405180910390fd5b60fc54813511612dcd5760405162e1442b60e41b815260040160405180910390fd5b604080513360208083019190915283358284018190528351808403850181526060840185528051908301206080808501929092528451808503909201825260a090930190935282519201919091205f612e268383613228565b90508060ff16600103612e3f57612e3f6101155f613bb4565b612e4c8460200135613751565b6040805160208681013582528635908201524381830152905133917f6c291a7ce9906b2982643002c104cb0ba9f2b9fecc8b38e9cc3cf5cfca65b4e8919081900360600190a2600361010a546002612ea4919061448b565b612eae9190613ecc565b612eb9906001613dd5565b8160ff161061118b57833560fc55602084013560fd55612eda61011561385f565b60fd5560408051602086810135825286359082015243918101919091527fbc1a0795e699bbeaa982f6049ae9689f4d0e3e22d554adb7c46626be62f3b8bc90606001611182565b5f610c457f8ec4e223bb52129c3d652c6e55a137389860823d9a02acb9d051e591994c6d6f613131565b60975460ff1615610b6b5760405162461bcd60e51b815260206004820152601060248201526f14185d5cd8589b194e881c185d5cd95960821b6044820152606401610c02565b61010383905561010482905561010281905560408051828152602081018590529081018390524260608201527ff525f19964f35afcb9a475680bb27abecbc5e62b4d6d4f66a81a5bd7e8a107e39060800160405180910390a1505050565b610c2b81336138d0565b6130038282611d42565b610c15575f8281526065602090815260408083206001600160a01b03851684529091529020805460ff1916600117905561303a3390565b6001600160a01b0316816001600160a01b0316837f2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d60405160405180910390a45050565b6130888282611d42565b15610c15575f8281526065602090815260408083206001600160a01b0385168085529252808320805460ff1916905551339285917ff6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b9190a45050565b6130ec613929565b6097805460ff191690557f5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa335b6040516001600160a01b039091168152602001611524565b5f8181526101166020526040812054808203613160576040516379a715fb60e11b815260040160405180910390fd5b8061316b8143613ecc565b613175919061448b565b9392505050565b6040516318903ee760e21b81526001600160a01b038381166004830152821690636240fb9c90602401602060405180830381865afa1580156131c0573d5f803e3d5ffd5b505050506040513d601f19601f820116820180604052508101906131e4919061446c565b610c155760405163c4230ae360e01b815260040160405180910390fd5b6001600160a01b038116610c2b5760405163d92e233d60e01b815260040160405180910390fd5b5f828152610110602052604081205460ff161561325857604051635da1eec160e11b815260040160405180910390fd5b5f83815261011060209081526040808320805460ff191660011790558483526101119091528120805460ff169161328e836144a2565b82546101009290920a60ff8181021990931691831602179091555f9384526101116020526040909320549092169392505050565b5f603082146132e457604051639ca717ed60e01b815260040160405180910390fd5b6040516002906132fc90859085905f906020016144c0565b60408051601f1981840301815290829052613316916144de565b602060405180830381855afa158015613331573d5f803e3d5ffd5b5050506040513d601f19601f82011682018060405250810190613175919061443d565b600260c954036133a65760405162461bcd60e51b815260206004820152601f60248201527f5265656e7472616e637947756172643a207265656e7472616e742063616c6c006044820152606401610c02565b600260c955565b805f036133cd57604051637036cfc960e11b815260040160405180910390fd5b5f828152610116602052604090205481036133fb5760405163806e577f60e01b815260040160405180910390fd5b5f828152610116602052604090819020829055517f6231a3e049e2072a042ae727208e7650b487871f4080458371e84d6e7d39113890611e859083815260200190565b610103546101045460fe545f9261345f9290916001600160a01b0316613972565b60fe549091505f9061347d90869086906001600160a01b0316613972565b90506127106101085461271061349391906144f9565b61349d908461448b565b6134a79190613ecc565b81101580156134dd5750612710610108546127106134c59190613dd5565b6134cf908461448b565b6134d99190613ecc565b8111155b6135465760fb805460ff191660019081179091554361010c5561010086905561010185905560ff849055604080519182524260208301527f9dea6ddefbcfcf9c4f6c1c086e462c2ab380c0be199d0289bf23b09f20814415910160405180910390a15050505050565b61148a858585612f91565b613559612f4b565b6097805460ff191660011790557f62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a2586131193390565b5f805f8060fe5f9054906101000a90046001600160a01b03166001600160a01b031663489ed6516040518163ffffffff1660e01b8152600401602060405180830381865afa1580156135e2573d5f803e3d5ffd5b505050506040513d601f19601f820116820180604052508101906136069190614364565b6001600160a01b031663feaf968c6040518163ffffffff1660e01b815260040160a060405180830381865afa158015613641573d5f803e3d5ffd5b505050506040513d601f19601f820116820180604052508101906136659190614525565b5050509150505f60fe5f9054906101000a90046001600160a01b03166001600160a01b0316632ca03f666040518163ffffffff1660e01b8152600401602060405180830381865afa1580156136bc573d5f803e3d5ffd5b505050506040513d601f19601f820116820180604052508101906136e09190614364565b6001600160a01b031663feaf968c6040518163ffffffff1660e01b815260040160a060405180830381865afa15801561371b573d5f803e3d5ffd5b505050506040513d601f19601f8201168201806040525081019061373f9190614525565b50949891975043965090945050505050565b6101158054600181810183555f8390527f77886ee853a0f02c8259429c7bfca08679ca3ab8bdec2c3bd5cca8557e06493a909101839055905490036137935750565b610115545f906137a5906001906144f9565b90505b600181101580156137de57506101156137c26001836144f9565b815481106137d2576137d26141d5565b905f5260205f20015482105b1561383b576101156137f16001836144f9565b81548110613801576138016141d5565b905f5260205f200154610115828154811061381e5761381e6141d5565b5f918252602090912001558061383381613de8565b9150506137a8565b816101158281548110613850576138506141d5565b5f918252602090912001555050565b80545f906002836138708284613ecc565b81548110613880576138806141d5565b905f5260205f20015484600260018561389991906144f9565b6138a39190613ecc565b815481106138b3576138b36141d5565b905f5260205f2001546138c69190613dd5565b6131759190613ecc565b6138da8282611d42565b610c15576138e781613a0c565b6138f2836020613a1e565b604051602001613903929190614571565b60408051601f198184030181529082905262461bcd60e51b8252610c02916004016145e5565b60975460ff16610b6b5760405162461bcd60e51b815260206004820152601460248201527314185d5cd8589b194e881b9bdd081c185d5cd95960621b6044820152606401610c02565b5f80826001600160a01b031663f0141d846040518163ffffffff1660e01b8152600401602060405180830381865afa1580156139b0573d5f803e3d5ffd5b505050506040513d601f19601f820116820180604052508101906139d4919061443d565b90505f8515806139e2575084155b613a0057846139f1838861448b565b6139fb9190613ecc565b613a02565b815b9695505050505050565b6060610b1c6001600160a01b03831660145b60605f613a2c83600261448b565b613a37906002613dd5565b67ffffffffffffffff811115613a4f57613a4f6145f7565b6040519080825280601f01601f191660200182016040528015613a79576020820181803683370190505b509050600360fc1b815f81518110613a9357613a936141d5565b60200101906001600160f81b03191690815f1a905350600f60fb1b81600181518110613ac157613ac16141d5565b60200101906001600160f81b03191690815f1a9053505f613ae384600261448b565b613aee906001613dd5565b90505b6001811115613b65576f181899199a1a9b1b9c1cb0b131b232b360811b85600f1660108110613b2257613b226141d5565b1a60f81b828281518110613b3857613b386141d5565b60200101906001600160f81b03191690815f1a90535060049490941c93613b5e81613de8565b9050613af1565b5083156131755760405162461bcd60e51b815260206004820181905260248201527f537472696e67733a20686578206c656e67746820696e73756666696369656e746044820152606401610c02565b5080545f8255905f5260205f2090810190610c2b91905b80821115613bde575f8155600101613bcb565b5090565b5f60208284031215613bf2575f80fd5b81356001600160e01b031981168114613175575f80fd5b5f60208284031215613c19575f80fd5b5035919050565b6001600160a01b0381168114610c2b575f80fd5b5f8060408385031215613c45575f80fd5b823591506020830135613c5781613c20565b809150509250929050565b5f60208284031215613c72575f80fd5b813561317581613c20565b9687526001600160801b039586166020880152938516604087015291909316606085015263ffffffff9283166080850152821660a08401521660c082015260e00190565b5f60e08284031215613cd1575f80fd5b50919050565b5f60608284031215613cd1575f80fd5b5f60208284031215613cf7575f80fd5b813567ffffffffffffffff811115613d0d575f80fd5b613d1984828501613cd7565b949350505050565b5f60208284031215613d31575f80fd5b813567ffffffffffffffff811115613d47575f80fd5b820160a08185031215613175575f80fd5b5f60608284031215613d68575f80fd5b6131758383613cd7565b803560ff81168114613d82575f80fd5b919050565b5f60208284031215613d97575f80fd5b61317582613d72565b5f6101008284031215613cd1575f80fd5b5f60408284031215613cd1575f80fd5b634e487b7160e01b5f52601160045260245ffd5b80820180821115610b1c57610b1c613dc1565b5f81613df657613df6613dc1565b505f190190565b634e487b7160e01b5f52601260045260245ffd5b5f82613e1f57613e1f613dfd565b500690565b6001600160801b0381168114610c2b575f80fd5b5f60208284031215613e48575f80fd5b813561317581613e24565b63ffffffff81168114610c2b575f80fd5b5f60208284031215613e74575f80fd5b813561317581613e53565b9788526001600160801b039687166020890152948616604088015292909416606086015263ffffffff908116608086015292831660a085015290911660c083015260e08201526101000190565b5f82613eda57613eda613dfd565b500490565b5f8135610b1c81613e53565b81358155600181016020830135613f0181613e24565b81546001600160801b0319166001600160801b038216178255506040830135613f2981613e24565b81546001600160801b031660809190911b6001600160801b031916179055600281016060830135613f5981613e24565b81546001600160801b0319166001600160801b038216178255506080830135613f8181613e53565b815463ffffffff60801b191660809190911b63ffffffff60801b16178155613fd2613fae60a08501613edf565b82805463ffffffff60a01b191660a09290921b63ffffffff60a01b16919091179055565b610b91613fe160c08501613edf565b82805463ffffffff60c01b191660c09290921b63ffffffff60c01b16919091179055565b5f808335601e1984360301811261401a575f80fd5b83018035915067ffffffffffffffff821115614034575f80fd5b6020019150600581901b360382131561404b575f80fd5b9250929050565b81835281816020850137505f828201602090810191909152601f909101601f19169091010190565b8183525f6020808501808196508560051b81019150845f5b878110156140fa5782840389528135601e198836030181126140b2575f80fd5b8701858101903567ffffffffffffffff8111156140cd575f80fd5b8036038213156140db575f80fd5b6140e6868284614052565b9a87019a9550505090840190600101614092565b5091979650505050505050565b602081525f613d1960208301848661407a565b5f5b8381101561413457818101518382015260200161411c565b50505f910152565b5f815180845261415381602086016020860161411a565b601f01601f19169290920160200192915050565b60018060a01b0384168152826020820152606060408201525f611fd1606083018461413c565b828152604060208201525f613d19604083018461413c565b858152846020820152836040820152608060608201525f6141ca60808301848661407a565b979650505050505050565b634e487b7160e01b5f52603260045260245ffd5b5f808335601e198436030181126141fe575f80fd5b83018035915067ffffffffffffffff821115614218575f80fd5b60200191503681900382131561404b575f80fd5b5f61ffff80831681810361424257614242613dc1565b6001019392505050565b848152836020820152606060408201525f613a0260608301848661407a565b606081525f61427e60608301888a61407a565b828103602084015261429181878961407a565b905082810360408401526142a681858761407a565b9998505050505050505050565b60018060a01b038516815260ff84166020820152826040820152608060608201525f613a02608083018461413c565b60ff84168152826020820152606060408201525f611fd1606083018461413c565b60ff8a16815288602082015260c060408201525f61432560c08301898b61407a565b828103606084015261433881888a61407a565b9050828103608084015261434d81868861407a565b9150508260a08301529a9950505050505050505050565b5f60208284031215614374575f80fd5b815161317581613c20565b5f805f60608486031215614391575f80fd5b8351925060208401519150604084015190509250925092565b60ff86168152846020820152608060408201525f6143cc60808301858761407a565b90508260608301529695505050505050565b813581526020808301359082015260408083013590820152610100810160ff61440960608501613d72565b1660608301526080830135608083015260a083013560a083015260c083013560c083015260e083013560e083015292915050565b5f6020828403121561444d575f80fd5b5051919050565b5f6001820161446557614465613dc1565b5060010190565b5f6020828403121561447c575f80fd5b81518015158114613175575f80fd5b8082028115828204841417610b1c57610b1c613dc1565b5f60ff821660ff81036144b7576144b7613dc1565b60010192915050565b828482376001600160801b0319919091169101908152601001919050565b5f82516144ef81846020870161411a565b9190910192915050565b81810381811115610b1c57610b1c613dc1565b805169ffffffffffffffffffff81168114613d82575f80fd5b5f805f805f60a08688031215614539575f80fd5b6145428661450c565b94506020860151935060408601519250606086015191506145656080870161450c565b90509295509295909350565b7f416363657373436f6e74726f6c3a206163636f756e742000000000000000000081525f83516145a881601785016020880161411a565b7001034b99036b4b9b9b4b733903937b6329607d1b60179184019182015283516145d981602884016020880161411a565b01602801949350505050565b602081525f613175602083018461413c565b634e487b7160e01b5f52604160045260245ffdfe783e3ebf40ee443ac9cbca8dc88c9f47450598583c2168f0ae0021de08ad333b6a7b51c29672c0ed412394a3e65ab758361d7c963b8657ce8c1f43dc0770b1629ae142790790fc18374cd6a6cc28573ecc78841658523afa63055cea42a9e1dd7607f5053d01557adb731d4aad009009bba2bf77a5b5f779733919451d426336a2646970667358221220c4ed90d27fad89c563dc7b8b1e4c2d00c1200316230dbc228abd053ef1a8301d64736f6c63430008140033496e697469616c697a61626c653a20636f6e7472616374206973206e6f742069",
}

// StaderOracleABI is the input ABI used to generate the binding from.
// Deprecated: Use StaderOracleMetaData.ABI instead.
var StaderOracleABI = StaderOracleMetaData.ABI

// StaderOracleBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use StaderOracleMetaData.Bin instead.
var StaderOracleBin = StaderOracleMetaData.Bin

// DeployStaderOracle deploys a new Ethereum contract, binding an instance of StaderOracle to it.
func DeployStaderOracle(auth *bind.TransactOpts, backend bind.ContractBackend, _admin common.Address, _staderConfig common.Address) (common.Address, *types.Transaction, *StaderOracle, error) {
	parsed, err := StaderOracleMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(StaderOracleBin), backend, _admin, _staderConfig)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &StaderOracle{StaderOracleCaller: StaderOracleCaller{contract: contract}, StaderOracleTransactor: StaderOracleTransactor{contract: contract}, StaderOracleFilterer: StaderOracleFilterer{contract: contract}}, nil
}

// StaderOracle is an auto generated Go binding around an Ethereum contract.
type StaderOracle struct {
	StaderOracleCaller     // Read-only binding to the contract
	StaderOracleTransactor // Write-only binding to the contract
	StaderOracleFilterer   // Log filterer for contract events
}

// StaderOracleCaller is an auto generated read-only Go binding around an Ethereum contract.
type StaderOracleCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StaderOracleTransactor is an auto generated write-only Go binding around an Ethereum contract.
type StaderOracleTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StaderOracleFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type StaderOracleFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StaderOracleSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type StaderOracleSession struct {
	Contract     *StaderOracle     // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// StaderOracleCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type StaderOracleCallerSession struct {
	Contract *StaderOracleCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts       // Call options to use throughout this session
}

// StaderOracleTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type StaderOracleTransactorSession struct {
	Contract     *StaderOracleTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts       // Transaction auth options to use throughout this session
}

// StaderOracleRaw is an auto generated low-level Go binding around an Ethereum contract.
type StaderOracleRaw struct {
	Contract *StaderOracle // Generic contract binding to access the raw methods on
}

// StaderOracleCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type StaderOracleCallerRaw struct {
	Contract *StaderOracleCaller // Generic read-only contract binding to access the raw methods on
}

// StaderOracleTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type StaderOracleTransactorRaw struct {
	Contract *StaderOracleTransactor // Generic write-only contract binding to access the raw methods on
}

// NewStaderOracle creates a new instance of StaderOracle, bound to a specific deployed contract.
func NewStaderOracle(address common.Address, backend bind.ContractBackend) (*StaderOracle, error) {
	contract, err := bindStaderOracle(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &StaderOracle{StaderOracleCaller: StaderOracleCaller{contract: contract}, StaderOracleTransactor: StaderOracleTransactor{contract: contract}, StaderOracleFilterer: StaderOracleFilterer{contract: contract}}, nil
}

// NewStaderOracleCaller creates a new read-only instance of StaderOracle, bound to a specific deployed contract.
func NewStaderOracleCaller(address common.Address, caller bind.ContractCaller) (*StaderOracleCaller, error) {
	contract, err := bindStaderOracle(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &StaderOracleCaller{contract: contract}, nil
}

// NewStaderOracleTransactor creates a new write-only instance of StaderOracle, bound to a specific deployed contract.
func NewStaderOracleTransactor(address common.Address, transactor bind.ContractTransactor) (*StaderOracleTransactor, error) {
	contract, err := bindStaderOracle(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &StaderOracleTransactor{contract: contract}, nil
}

// NewStaderOracleFilterer creates a new log filterer instance of StaderOracle, bound to a specific deployed contract.
func NewStaderOracleFilterer(address common.Address, filterer bind.ContractFilterer) (*StaderOracleFilterer, error) {
	contract, err := bindStaderOracle(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &StaderOracleFilterer{contract: contract}, nil
}

// bindStaderOracle binds a generic wrapper to an already deployed contract.
func bindStaderOracle(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := StaderOracleMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_StaderOracle *StaderOracleRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _StaderOracle.Contract.StaderOracleCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_StaderOracle *StaderOracleRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _StaderOracle.Contract.StaderOracleTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_StaderOracle *StaderOracleRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _StaderOracle.Contract.StaderOracleTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_StaderOracle *StaderOracleCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _StaderOracle.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_StaderOracle *StaderOracleTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _StaderOracle.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_StaderOracle *StaderOracleTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _StaderOracle.Contract.contract.Transact(opts, method, params...)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_StaderOracle *StaderOracleCaller) DEFAULTADMINROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _StaderOracle.contract.Call(opts, &out, "DEFAULT_ADMIN_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_StaderOracle *StaderOracleSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _StaderOracle.Contract.DEFAULTADMINROLE(&_StaderOracle.CallOpts)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_StaderOracle *StaderOracleCallerSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _StaderOracle.Contract.DEFAULTADMINROLE(&_StaderOracle.CallOpts)
}

// ERCHANGEMAXBPS is a free data retrieval call binding the contract method 0xe0bcb378.
//
// Solidity: function ER_CHANGE_MAX_BPS() view returns(uint256)
func (_StaderOracle *StaderOracleCaller) ERCHANGEMAXBPS(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _StaderOracle.contract.Call(opts, &out, "ER_CHANGE_MAX_BPS")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ERCHANGEMAXBPS is a free data retrieval call binding the contract method 0xe0bcb378.
//
// Solidity: function ER_CHANGE_MAX_BPS() view returns(uint256)
func (_StaderOracle *StaderOracleSession) ERCHANGEMAXBPS() (*big.Int, error) {
	return _StaderOracle.Contract.ERCHANGEMAXBPS(&_StaderOracle.CallOpts)
}

// ERCHANGEMAXBPS is a free data retrieval call binding the contract method 0xe0bcb378.
//
// Solidity: function ER_CHANGE_MAX_BPS() view returns(uint256)
func (_StaderOracle *StaderOracleCallerSession) ERCHANGEMAXBPS() (*big.Int, error) {
	return _StaderOracle.Contract.ERCHANGEMAXBPS(&_StaderOracle.CallOpts)
}

// ETHXERUF is a free data retrieval call binding the contract method 0x12710361.
//
// Solidity: function ETHX_ER_UF() view returns(bytes32)
func (_StaderOracle *StaderOracleCaller) ETHXERUF(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _StaderOracle.contract.Call(opts, &out, "ETHX_ER_UF")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// ETHXERUF is a free data retrieval call binding the contract method 0x12710361.
//
// Solidity: function ETHX_ER_UF() view returns(bytes32)
func (_StaderOracle *StaderOracleSession) ETHXERUF() ([32]byte, error) {
	return _StaderOracle.Contract.ETHXERUF(&_StaderOracle.CallOpts)
}

// ETHXERUF is a free data retrieval call binding the contract method 0x12710361.
//
// Solidity: function ETHX_ER_UF() view returns(bytes32)
func (_StaderOracle *StaderOracleCallerSession) ETHXERUF() ([32]byte, error) {
	return _StaderOracle.Contract.ETHXERUF(&_StaderOracle.CallOpts)
}

// MAXERUPDATEFREQUENCY is a free data retrieval call binding the contract method 0xf10b2569.
//
// Solidity: function MAX_ER_UPDATE_FREQUENCY() view returns(uint256)
func (_StaderOracle *StaderOracleCaller) MAXERUPDATEFREQUENCY(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _StaderOracle.contract.Call(opts, &out, "MAX_ER_UPDATE_FREQUENCY")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MAXERUPDATEFREQUENCY is a free data retrieval call binding the contract method 0xf10b2569.
//
// Solidity: function MAX_ER_UPDATE_FREQUENCY() view returns(uint256)
func (_StaderOracle *StaderOracleSession) MAXERUPDATEFREQUENCY() (*big.Int, error) {
	return _StaderOracle.Contract.MAXERUPDATEFREQUENCY(&_StaderOracle.CallOpts)
}

// MAXERUPDATEFREQUENCY is a free data retrieval call binding the contract method 0xf10b2569.
//
// Solidity: function MAX_ER_UPDATE_FREQUENCY() view returns(uint256)
func (_StaderOracle *StaderOracleCallerSession) MAXERUPDATEFREQUENCY() (*big.Int, error) {
	return _StaderOracle.Contract.MAXERUPDATEFREQUENCY(&_StaderOracle.CallOpts)
}

// MINTRUSTEDNODES is a free data retrieval call binding the contract method 0xe36c3140.
//
// Solidity: function MIN_TRUSTED_NODES() view returns(uint256)
func (_StaderOracle *StaderOracleCaller) MINTRUSTEDNODES(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _StaderOracle.contract.Call(opts, &out, "MIN_TRUSTED_NODES")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MINTRUSTEDNODES is a free data retrieval call binding the contract method 0xe36c3140.
//
// Solidity: function MIN_TRUSTED_NODES() view returns(uint256)
func (_StaderOracle *StaderOracleSession) MINTRUSTEDNODES() (*big.Int, error) {
	return _StaderOracle.Contract.MINTRUSTEDNODES(&_StaderOracle.CallOpts)
}

// MINTRUSTEDNODES is a free data retrieval call binding the contract method 0xe36c3140.
//
// Solidity: function MIN_TRUSTED_NODES() view returns(uint256)
func (_StaderOracle *StaderOracleCallerSession) MINTRUSTEDNODES() (*big.Int, error) {
	return _StaderOracle.Contract.MINTRUSTEDNODES(&_StaderOracle.CallOpts)
}

// MISSEDATTESTATIONPENALTYUF is a free data retrieval call binding the contract method 0x052a6840.
//
// Solidity: function MISSED_ATTESTATION_PENALTY_UF() view returns(bytes32)
func (_StaderOracle *StaderOracleCaller) MISSEDATTESTATIONPENALTYUF(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _StaderOracle.contract.Call(opts, &out, "MISSED_ATTESTATION_PENALTY_UF")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// MISSEDATTESTATIONPENALTYUF is a free data retrieval call binding the contract method 0x052a6840.
//
// Solidity: function MISSED_ATTESTATION_PENALTY_UF() view returns(bytes32)
func (_StaderOracle *StaderOracleSession) MISSEDATTESTATIONPENALTYUF() ([32]byte, error) {
	return _StaderOracle.Contract.MISSEDATTESTATIONPENALTYUF(&_StaderOracle.CallOpts)
}

// MISSEDATTESTATIONPENALTYUF is a free data retrieval call binding the contract method 0x052a6840.
//
// Solidity: function MISSED_ATTESTATION_PENALTY_UF() view returns(bytes32)
func (_StaderOracle *StaderOracleCallerSession) MISSEDATTESTATIONPENALTYUF() ([32]byte, error) {
	return _StaderOracle.Contract.MISSEDATTESTATIONPENALTYUF(&_StaderOracle.CallOpts)
}

// SDPRICEUF is a free data retrieval call binding the contract method 0xe2f63392.
//
// Solidity: function SD_PRICE_UF() view returns(bytes32)
func (_StaderOracle *StaderOracleCaller) SDPRICEUF(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _StaderOracle.contract.Call(opts, &out, "SD_PRICE_UF")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// SDPRICEUF is a free data retrieval call binding the contract method 0xe2f63392.
//
// Solidity: function SD_PRICE_UF() view returns(bytes32)
func (_StaderOracle *StaderOracleSession) SDPRICEUF() ([32]byte, error) {
	return _StaderOracle.Contract.SDPRICEUF(&_StaderOracle.CallOpts)
}

// SDPRICEUF is a free data retrieval call binding the contract method 0xe2f63392.
//
// Solidity: function SD_PRICE_UF() view returns(bytes32)
func (_StaderOracle *StaderOracleCallerSession) SDPRICEUF() ([32]byte, error) {
	return _StaderOracle.Contract.SDPRICEUF(&_StaderOracle.CallOpts)
}

// VALIDATORSTATSUF is a free data retrieval call binding the contract method 0x29f96856.
//
// Solidity: function VALIDATOR_STATS_UF() view returns(bytes32)
func (_StaderOracle *StaderOracleCaller) VALIDATORSTATSUF(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _StaderOracle.contract.Call(opts, &out, "VALIDATOR_STATS_UF")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// VALIDATORSTATSUF is a free data retrieval call binding the contract method 0x29f96856.
//
// Solidity: function VALIDATOR_STATS_UF() view returns(bytes32)
func (_StaderOracle *StaderOracleSession) VALIDATORSTATSUF() ([32]byte, error) {
	return _StaderOracle.Contract.VALIDATORSTATSUF(&_StaderOracle.CallOpts)
}

// VALIDATORSTATSUF is a free data retrieval call binding the contract method 0x29f96856.
//
// Solidity: function VALIDATOR_STATS_UF() view returns(bytes32)
func (_StaderOracle *StaderOracleCallerSession) VALIDATORSTATSUF() ([32]byte, error) {
	return _StaderOracle.Contract.VALIDATORSTATSUF(&_StaderOracle.CallOpts)
}

// VALIDATORVERIFICATIONDETAILUF is a free data retrieval call binding the contract method 0xbe48e58d.
//
// Solidity: function VALIDATOR_VERIFICATION_DETAIL_UF() view returns(bytes32)
func (_StaderOracle *StaderOracleCaller) VALIDATORVERIFICATIONDETAILUF(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _StaderOracle.contract.Call(opts, &out, "VALIDATOR_VERIFICATION_DETAIL_UF")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// VALIDATORVERIFICATIONDETAILUF is a free data retrieval call binding the contract method 0xbe48e58d.
//
// Solidity: function VALIDATOR_VERIFICATION_DETAIL_UF() view returns(bytes32)
func (_StaderOracle *StaderOracleSession) VALIDATORVERIFICATIONDETAILUF() ([32]byte, error) {
	return _StaderOracle.Contract.VALIDATORVERIFICATIONDETAILUF(&_StaderOracle.CallOpts)
}

// VALIDATORVERIFICATIONDETAILUF is a free data retrieval call binding the contract method 0xbe48e58d.
//
// Solidity: function VALIDATOR_VERIFICATION_DETAIL_UF() view returns(bytes32)
func (_StaderOracle *StaderOracleCallerSession) VALIDATORVERIFICATIONDETAILUF() ([32]byte, error) {
	return _StaderOracle.Contract.VALIDATORVERIFICATIONDETAILUF(&_StaderOracle.CallOpts)
}

// WITHDRAWNVALIDATORSUF is a free data retrieval call binding the contract method 0xe61befa7.
//
// Solidity: function WITHDRAWN_VALIDATORS_UF() view returns(bytes32)
func (_StaderOracle *StaderOracleCaller) WITHDRAWNVALIDATORSUF(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _StaderOracle.contract.Call(opts, &out, "WITHDRAWN_VALIDATORS_UF")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// WITHDRAWNVALIDATORSUF is a free data retrieval call binding the contract method 0xe61befa7.
//
// Solidity: function WITHDRAWN_VALIDATORS_UF() view returns(bytes32)
func (_StaderOracle *StaderOracleSession) WITHDRAWNVALIDATORSUF() ([32]byte, error) {
	return _StaderOracle.Contract.WITHDRAWNVALIDATORSUF(&_StaderOracle.CallOpts)
}

// WITHDRAWNVALIDATORSUF is a free data retrieval call binding the contract method 0xe61befa7.
//
// Solidity: function WITHDRAWN_VALIDATORS_UF() view returns(bytes32)
func (_StaderOracle *StaderOracleCallerSession) WITHDRAWNVALIDATORSUF() ([32]byte, error) {
	return _StaderOracle.Contract.WITHDRAWNVALIDATORSUF(&_StaderOracle.CallOpts)
}

// ErChangeLimit is a free data retrieval call binding the contract method 0x97a3a10a.
//
// Solidity: function erChangeLimit() view returns(uint256)
func (_StaderOracle *StaderOracleCaller) ErChangeLimit(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _StaderOracle.contract.Call(opts, &out, "erChangeLimit")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ErChangeLimit is a free data retrieval call binding the contract method 0x97a3a10a.
//
// Solidity: function erChangeLimit() view returns(uint256)
func (_StaderOracle *StaderOracleSession) ErChangeLimit() (*big.Int, error) {
	return _StaderOracle.Contract.ErChangeLimit(&_StaderOracle.CallOpts)
}

// ErChangeLimit is a free data retrieval call binding the contract method 0x97a3a10a.
//
// Solidity: function erChangeLimit() view returns(uint256)
func (_StaderOracle *StaderOracleCallerSession) ErChangeLimit() (*big.Int, error) {
	return _StaderOracle.Contract.ErChangeLimit(&_StaderOracle.CallOpts)
}

// ErInspectionMode is a free data retrieval call binding the contract method 0x342280b3.
//
// Solidity: function erInspectionMode() view returns(bool)
func (_StaderOracle *StaderOracleCaller) ErInspectionMode(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _StaderOracle.contract.Call(opts, &out, "erInspectionMode")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// ErInspectionMode is a free data retrieval call binding the contract method 0x342280b3.
//
// Solidity: function erInspectionMode() view returns(bool)
func (_StaderOracle *StaderOracleSession) ErInspectionMode() (bool, error) {
	return _StaderOracle.Contract.ErInspectionMode(&_StaderOracle.CallOpts)
}

// ErInspectionMode is a free data retrieval call binding the contract method 0x342280b3.
//
// Solidity: function erInspectionMode() view returns(bool)
func (_StaderOracle *StaderOracleCallerSession) ErInspectionMode() (bool, error) {
	return _StaderOracle.Contract.ErInspectionMode(&_StaderOracle.CallOpts)
}

// ErInspectionModeStartBlock is a free data retrieval call binding the contract method 0xde271c6d.
//
// Solidity: function erInspectionModeStartBlock() view returns(uint256)
func (_StaderOracle *StaderOracleCaller) ErInspectionModeStartBlock(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _StaderOracle.contract.Call(opts, &out, "erInspectionModeStartBlock")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ErInspectionModeStartBlock is a free data retrieval call binding the contract method 0xde271c6d.
//
// Solidity: function erInspectionModeStartBlock() view returns(uint256)
func (_StaderOracle *StaderOracleSession) ErInspectionModeStartBlock() (*big.Int, error) {
	return _StaderOracle.Contract.ErInspectionModeStartBlock(&_StaderOracle.CallOpts)
}

// ErInspectionModeStartBlock is a free data retrieval call binding the contract method 0xde271c6d.
//
// Solidity: function erInspectionModeStartBlock() view returns(uint256)
func (_StaderOracle *StaderOracleCallerSession) ErInspectionModeStartBlock() (*big.Int, error) {
	return _StaderOracle.Contract.ErInspectionModeStartBlock(&_StaderOracle.CallOpts)
}

// ExchangeRate is a free data retrieval call binding the contract method 0x3ba0b9a9.
//
// Solidity: function exchangeRate() view returns(uint256 reportingBlockNumber, uint256 totalETHBalance, uint256 totalETHXSupply)
func (_StaderOracle *StaderOracleCaller) ExchangeRate(opts *bind.CallOpts) (struct {
	ReportingBlockNumber *big.Int
	TotalETHBalance      *big.Int
	TotalETHXSupply      *big.Int
}, error) {
	var out []interface{}
	err := _StaderOracle.contract.Call(opts, &out, "exchangeRate")

	outstruct := new(struct {
		ReportingBlockNumber *big.Int
		TotalETHBalance      *big.Int
		TotalETHXSupply      *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.ReportingBlockNumber = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.TotalETHBalance = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.TotalETHXSupply = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// ExchangeRate is a free data retrieval call binding the contract method 0x3ba0b9a9.
//
// Solidity: function exchangeRate() view returns(uint256 reportingBlockNumber, uint256 totalETHBalance, uint256 totalETHXSupply)
func (_StaderOracle *StaderOracleSession) ExchangeRate() (struct {
	ReportingBlockNumber *big.Int
	TotalETHBalance      *big.Int
	TotalETHXSupply      *big.Int
}, error) {
	return _StaderOracle.Contract.ExchangeRate(&_StaderOracle.CallOpts)
}

// ExchangeRate is a free data retrieval call binding the contract method 0x3ba0b9a9.
//
// Solidity: function exchangeRate() view returns(uint256 reportingBlockNumber, uint256 totalETHBalance, uint256 totalETHXSupply)
func (_StaderOracle *StaderOracleCallerSession) ExchangeRate() (struct {
	ReportingBlockNumber *big.Int
	TotalETHBalance      *big.Int
	TotalETHXSupply      *big.Int
}, error) {
	return _StaderOracle.Contract.ExchangeRate(&_StaderOracle.CallOpts)
}

// GetCurrentRewardsIndexByPoolId is a free data retrieval call binding the contract method 0xd06628ed.
//
// Solidity: function getCurrentRewardsIndexByPoolId(uint8 _poolId) view returns(uint256)
func (_StaderOracle *StaderOracleCaller) GetCurrentRewardsIndexByPoolId(opts *bind.CallOpts, _poolId uint8) (*big.Int, error) {
	var out []interface{}
	err := _StaderOracle.contract.Call(opts, &out, "getCurrentRewardsIndexByPoolId", _poolId)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetCurrentRewardsIndexByPoolId is a free data retrieval call binding the contract method 0xd06628ed.
//
// Solidity: function getCurrentRewardsIndexByPoolId(uint8 _poolId) view returns(uint256)
func (_StaderOracle *StaderOracleSession) GetCurrentRewardsIndexByPoolId(_poolId uint8) (*big.Int, error) {
	return _StaderOracle.Contract.GetCurrentRewardsIndexByPoolId(&_StaderOracle.CallOpts, _poolId)
}

// GetCurrentRewardsIndexByPoolId is a free data retrieval call binding the contract method 0xd06628ed.
//
// Solidity: function getCurrentRewardsIndexByPoolId(uint8 _poolId) view returns(uint256)
func (_StaderOracle *StaderOracleCallerSession) GetCurrentRewardsIndexByPoolId(_poolId uint8) (*big.Int, error) {
	return _StaderOracle.Contract.GetCurrentRewardsIndexByPoolId(&_StaderOracle.CallOpts, _poolId)
}

// GetERReportableBlock is a free data retrieval call binding the contract method 0xb5c25ba6.
//
// Solidity: function getERReportableBlock() view returns(uint256)
func (_StaderOracle *StaderOracleCaller) GetERReportableBlock(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _StaderOracle.contract.Call(opts, &out, "getERReportableBlock")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetERReportableBlock is a free data retrieval call binding the contract method 0xb5c25ba6.
//
// Solidity: function getERReportableBlock() view returns(uint256)
func (_StaderOracle *StaderOracleSession) GetERReportableBlock() (*big.Int, error) {
	return _StaderOracle.Contract.GetERReportableBlock(&_StaderOracle.CallOpts)
}

// GetERReportableBlock is a free data retrieval call binding the contract method 0xb5c25ba6.
//
// Solidity: function getERReportableBlock() view returns(uint256)
func (_StaderOracle *StaderOracleCallerSession) GetERReportableBlock() (*big.Int, error) {
	return _StaderOracle.Contract.GetERReportableBlock(&_StaderOracle.CallOpts)
}

// GetExchangeRate is a free data retrieval call binding the contract method 0xe6aa216c.
//
// Solidity: function getExchangeRate() view returns((uint256,uint256,uint256))
func (_StaderOracle *StaderOracleCaller) GetExchangeRate(opts *bind.CallOpts) (ExchangeRate, error) {
	var out []interface{}
	err := _StaderOracle.contract.Call(opts, &out, "getExchangeRate")

	if err != nil {
		return *new(ExchangeRate), err
	}

	out0 := *abi.ConvertType(out[0], new(ExchangeRate)).(*ExchangeRate)

	return out0, err

}

// GetExchangeRate is a free data retrieval call binding the contract method 0xe6aa216c.
//
// Solidity: function getExchangeRate() view returns((uint256,uint256,uint256))
func (_StaderOracle *StaderOracleSession) GetExchangeRate() (ExchangeRate, error) {
	return _StaderOracle.Contract.GetExchangeRate(&_StaderOracle.CallOpts)
}

// GetExchangeRate is a free data retrieval call binding the contract method 0xe6aa216c.
//
// Solidity: function getExchangeRate() view returns((uint256,uint256,uint256))
func (_StaderOracle *StaderOracleCallerSession) GetExchangeRate() (ExchangeRate, error) {
	return _StaderOracle.Contract.GetExchangeRate(&_StaderOracle.CallOpts)
}

// GetMerkleRootReportableBlockByPoolId is a free data retrieval call binding the contract method 0xa0c54387.
//
// Solidity: function getMerkleRootReportableBlockByPoolId(uint8 _poolId) view returns(uint256)
func (_StaderOracle *StaderOracleCaller) GetMerkleRootReportableBlockByPoolId(opts *bind.CallOpts, _poolId uint8) (*big.Int, error) {
	var out []interface{}
	err := _StaderOracle.contract.Call(opts, &out, "getMerkleRootReportableBlockByPoolId", _poolId)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetMerkleRootReportableBlockByPoolId is a free data retrieval call binding the contract method 0xa0c54387.
//
// Solidity: function getMerkleRootReportableBlockByPoolId(uint8 _poolId) view returns(uint256)
func (_StaderOracle *StaderOracleSession) GetMerkleRootReportableBlockByPoolId(_poolId uint8) (*big.Int, error) {
	return _StaderOracle.Contract.GetMerkleRootReportableBlockByPoolId(&_StaderOracle.CallOpts, _poolId)
}

// GetMerkleRootReportableBlockByPoolId is a free data retrieval call binding the contract method 0xa0c54387.
//
// Solidity: function getMerkleRootReportableBlockByPoolId(uint8 _poolId) view returns(uint256)
func (_StaderOracle *StaderOracleCallerSession) GetMerkleRootReportableBlockByPoolId(_poolId uint8) (*big.Int, error) {
	return _StaderOracle.Contract.GetMerkleRootReportableBlockByPoolId(&_StaderOracle.CallOpts, _poolId)
}

// GetMissedAttestationPenaltyReportableBlock is a free data retrieval call binding the contract method 0xa71b3907.
//
// Solidity: function getMissedAttestationPenaltyReportableBlock() view returns(uint256)
func (_StaderOracle *StaderOracleCaller) GetMissedAttestationPenaltyReportableBlock(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _StaderOracle.contract.Call(opts, &out, "getMissedAttestationPenaltyReportableBlock")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetMissedAttestationPenaltyReportableBlock is a free data retrieval call binding the contract method 0xa71b3907.
//
// Solidity: function getMissedAttestationPenaltyReportableBlock() view returns(uint256)
func (_StaderOracle *StaderOracleSession) GetMissedAttestationPenaltyReportableBlock() (*big.Int, error) {
	return _StaderOracle.Contract.GetMissedAttestationPenaltyReportableBlock(&_StaderOracle.CallOpts)
}

// GetMissedAttestationPenaltyReportableBlock is a free data retrieval call binding the contract method 0xa71b3907.
//
// Solidity: function getMissedAttestationPenaltyReportableBlock() view returns(uint256)
func (_StaderOracle *StaderOracleCallerSession) GetMissedAttestationPenaltyReportableBlock() (*big.Int, error) {
	return _StaderOracle.Contract.GetMissedAttestationPenaltyReportableBlock(&_StaderOracle.CallOpts)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_StaderOracle *StaderOracleCaller) GetRoleAdmin(opts *bind.CallOpts, role [32]byte) ([32]byte, error) {
	var out []interface{}
	err := _StaderOracle.contract.Call(opts, &out, "getRoleAdmin", role)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_StaderOracle *StaderOracleSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _StaderOracle.Contract.GetRoleAdmin(&_StaderOracle.CallOpts, role)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_StaderOracle *StaderOracleCallerSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _StaderOracle.Contract.GetRoleAdmin(&_StaderOracle.CallOpts, role)
}

// GetSDPriceInETH is a free data retrieval call binding the contract method 0xa6870e5b.
//
// Solidity: function getSDPriceInETH() view returns(uint256)
func (_StaderOracle *StaderOracleCaller) GetSDPriceInETH(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _StaderOracle.contract.Call(opts, &out, "getSDPriceInETH")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetSDPriceInETH is a free data retrieval call binding the contract method 0xa6870e5b.
//
// Solidity: function getSDPriceInETH() view returns(uint256)
func (_StaderOracle *StaderOracleSession) GetSDPriceInETH() (*big.Int, error) {
	return _StaderOracle.Contract.GetSDPriceInETH(&_StaderOracle.CallOpts)
}

// GetSDPriceInETH is a free data retrieval call binding the contract method 0xa6870e5b.
//
// Solidity: function getSDPriceInETH() view returns(uint256)
func (_StaderOracle *StaderOracleCallerSession) GetSDPriceInETH() (*big.Int, error) {
	return _StaderOracle.Contract.GetSDPriceInETH(&_StaderOracle.CallOpts)
}

// GetSDPriceReportableBlock is a free data retrieval call binding the contract method 0xfc8b821c.
//
// Solidity: function getSDPriceReportableBlock() view returns(uint256)
func (_StaderOracle *StaderOracleCaller) GetSDPriceReportableBlock(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _StaderOracle.contract.Call(opts, &out, "getSDPriceReportableBlock")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetSDPriceReportableBlock is a free data retrieval call binding the contract method 0xfc8b821c.
//
// Solidity: function getSDPriceReportableBlock() view returns(uint256)
func (_StaderOracle *StaderOracleSession) GetSDPriceReportableBlock() (*big.Int, error) {
	return _StaderOracle.Contract.GetSDPriceReportableBlock(&_StaderOracle.CallOpts)
}

// GetSDPriceReportableBlock is a free data retrieval call binding the contract method 0xfc8b821c.
//
// Solidity: function getSDPriceReportableBlock() view returns(uint256)
func (_StaderOracle *StaderOracleCallerSession) GetSDPriceReportableBlock() (*big.Int, error) {
	return _StaderOracle.Contract.GetSDPriceReportableBlock(&_StaderOracle.CallOpts)
}

// GetValidatorStats is a free data retrieval call binding the contract method 0x3e23a827.
//
// Solidity: function getValidatorStats() view returns((uint256,uint128,uint128,uint128,uint32,uint32,uint32))
func (_StaderOracle *StaderOracleCaller) GetValidatorStats(opts *bind.CallOpts) (ValidatorStats, error) {
	var out []interface{}
	err := _StaderOracle.contract.Call(opts, &out, "getValidatorStats")

	if err != nil {
		return *new(ValidatorStats), err
	}

	out0 := *abi.ConvertType(out[0], new(ValidatorStats)).(*ValidatorStats)

	return out0, err

}

// GetValidatorStats is a free data retrieval call binding the contract method 0x3e23a827.
//
// Solidity: function getValidatorStats() view returns((uint256,uint128,uint128,uint128,uint32,uint32,uint32))
func (_StaderOracle *StaderOracleSession) GetValidatorStats() (ValidatorStats, error) {
	return _StaderOracle.Contract.GetValidatorStats(&_StaderOracle.CallOpts)
}

// GetValidatorStats is a free data retrieval call binding the contract method 0x3e23a827.
//
// Solidity: function getValidatorStats() view returns((uint256,uint128,uint128,uint128,uint32,uint32,uint32))
func (_StaderOracle *StaderOracleCallerSession) GetValidatorStats() (ValidatorStats, error) {
	return _StaderOracle.Contract.GetValidatorStats(&_StaderOracle.CallOpts)
}

// GetValidatorStatsReportableBlock is a free data retrieval call binding the contract method 0x49115a2e.
//
// Solidity: function getValidatorStatsReportableBlock() view returns(uint256)
func (_StaderOracle *StaderOracleCaller) GetValidatorStatsReportableBlock(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _StaderOracle.contract.Call(opts, &out, "getValidatorStatsReportableBlock")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetValidatorStatsReportableBlock is a free data retrieval call binding the contract method 0x49115a2e.
//
// Solidity: function getValidatorStatsReportableBlock() view returns(uint256)
func (_StaderOracle *StaderOracleSession) GetValidatorStatsReportableBlock() (*big.Int, error) {
	return _StaderOracle.Contract.GetValidatorStatsReportableBlock(&_StaderOracle.CallOpts)
}

// GetValidatorStatsReportableBlock is a free data retrieval call binding the contract method 0x49115a2e.
//
// Solidity: function getValidatorStatsReportableBlock() view returns(uint256)
func (_StaderOracle *StaderOracleCallerSession) GetValidatorStatsReportableBlock() (*big.Int, error) {
	return _StaderOracle.Contract.GetValidatorStatsReportableBlock(&_StaderOracle.CallOpts)
}

// GetValidatorVerificationDetailReportableBlock is a free data retrieval call binding the contract method 0x615a0253.
//
// Solidity: function getValidatorVerificationDetailReportableBlock() view returns(uint256)
func (_StaderOracle *StaderOracleCaller) GetValidatorVerificationDetailReportableBlock(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _StaderOracle.contract.Call(opts, &out, "getValidatorVerificationDetailReportableBlock")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetValidatorVerificationDetailReportableBlock is a free data retrieval call binding the contract method 0x615a0253.
//
// Solidity: function getValidatorVerificationDetailReportableBlock() view returns(uint256)
func (_StaderOracle *StaderOracleSession) GetValidatorVerificationDetailReportableBlock() (*big.Int, error) {
	return _StaderOracle.Contract.GetValidatorVerificationDetailReportableBlock(&_StaderOracle.CallOpts)
}

// GetValidatorVerificationDetailReportableBlock is a free data retrieval call binding the contract method 0x615a0253.
//
// Solidity: function getValidatorVerificationDetailReportableBlock() view returns(uint256)
func (_StaderOracle *StaderOracleCallerSession) GetValidatorVerificationDetailReportableBlock() (*big.Int, error) {
	return _StaderOracle.Contract.GetValidatorVerificationDetailReportableBlock(&_StaderOracle.CallOpts)
}

// GetWithdrawnValidatorReportableBlock is a free data retrieval call binding the contract method 0x5063b5bd.
//
// Solidity: function getWithdrawnValidatorReportableBlock() view returns(uint256)
func (_StaderOracle *StaderOracleCaller) GetWithdrawnValidatorReportableBlock(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _StaderOracle.contract.Call(opts, &out, "getWithdrawnValidatorReportableBlock")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetWithdrawnValidatorReportableBlock is a free data retrieval call binding the contract method 0x5063b5bd.
//
// Solidity: function getWithdrawnValidatorReportableBlock() view returns(uint256)
func (_StaderOracle *StaderOracleSession) GetWithdrawnValidatorReportableBlock() (*big.Int, error) {
	return _StaderOracle.Contract.GetWithdrawnValidatorReportableBlock(&_StaderOracle.CallOpts)
}

// GetWithdrawnValidatorReportableBlock is a free data retrieval call binding the contract method 0x5063b5bd.
//
// Solidity: function getWithdrawnValidatorReportableBlock() view returns(uint256)
func (_StaderOracle *StaderOracleCallerSession) GetWithdrawnValidatorReportableBlock() (*big.Int, error) {
	return _StaderOracle.Contract.GetWithdrawnValidatorReportableBlock(&_StaderOracle.CallOpts)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_StaderOracle *StaderOracleCaller) HasRole(opts *bind.CallOpts, role [32]byte, account common.Address) (bool, error) {
	var out []interface{}
	err := _StaderOracle.contract.Call(opts, &out, "hasRole", role, account)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_StaderOracle *StaderOracleSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _StaderOracle.Contract.HasRole(&_StaderOracle.CallOpts, role, account)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_StaderOracle *StaderOracleCallerSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _StaderOracle.Contract.HasRole(&_StaderOracle.CallOpts, role, account)
}

// InspectionModeExchangeRate is a free data retrieval call binding the contract method 0xb940a003.
//
// Solidity: function inspectionModeExchangeRate() view returns(uint256 reportingBlockNumber, uint256 totalETHBalance, uint256 totalETHXSupply)
func (_StaderOracle *StaderOracleCaller) InspectionModeExchangeRate(opts *bind.CallOpts) (struct {
	ReportingBlockNumber *big.Int
	TotalETHBalance      *big.Int
	TotalETHXSupply      *big.Int
}, error) {
	var out []interface{}
	err := _StaderOracle.contract.Call(opts, &out, "inspectionModeExchangeRate")

	outstruct := new(struct {
		ReportingBlockNumber *big.Int
		TotalETHBalance      *big.Int
		TotalETHXSupply      *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.ReportingBlockNumber = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.TotalETHBalance = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.TotalETHXSupply = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// InspectionModeExchangeRate is a free data retrieval call binding the contract method 0xb940a003.
//
// Solidity: function inspectionModeExchangeRate() view returns(uint256 reportingBlockNumber, uint256 totalETHBalance, uint256 totalETHXSupply)
func (_StaderOracle *StaderOracleSession) InspectionModeExchangeRate() (struct {
	ReportingBlockNumber *big.Int
	TotalETHBalance      *big.Int
	TotalETHXSupply      *big.Int
}, error) {
	return _StaderOracle.Contract.InspectionModeExchangeRate(&_StaderOracle.CallOpts)
}

// InspectionModeExchangeRate is a free data retrieval call binding the contract method 0xb940a003.
//
// Solidity: function inspectionModeExchangeRate() view returns(uint256 reportingBlockNumber, uint256 totalETHBalance, uint256 totalETHXSupply)
func (_StaderOracle *StaderOracleCallerSession) InspectionModeExchangeRate() (struct {
	ReportingBlockNumber *big.Int
	TotalETHBalance      *big.Int
	TotalETHXSupply      *big.Int
}, error) {
	return _StaderOracle.Contract.InspectionModeExchangeRate(&_StaderOracle.CallOpts)
}

// IsPORFeedBasedERData is a free data retrieval call binding the contract method 0x16515428.
//
// Solidity: function isPORFeedBasedERData() view returns(bool)
func (_StaderOracle *StaderOracleCaller) IsPORFeedBasedERData(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _StaderOracle.contract.Call(opts, &out, "isPORFeedBasedERData")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsPORFeedBasedERData is a free data retrieval call binding the contract method 0x16515428.
//
// Solidity: function isPORFeedBasedERData() view returns(bool)
func (_StaderOracle *StaderOracleSession) IsPORFeedBasedERData() (bool, error) {
	return _StaderOracle.Contract.IsPORFeedBasedERData(&_StaderOracle.CallOpts)
}

// IsPORFeedBasedERData is a free data retrieval call binding the contract method 0x16515428.
//
// Solidity: function isPORFeedBasedERData() view returns(bool)
func (_StaderOracle *StaderOracleCallerSession) IsPORFeedBasedERData() (bool, error) {
	return _StaderOracle.Contract.IsPORFeedBasedERData(&_StaderOracle.CallOpts)
}

// IsTrustedNode is a free data retrieval call binding the contract method 0x2f739b1d.
//
// Solidity: function isTrustedNode(address ) view returns(bool)
func (_StaderOracle *StaderOracleCaller) IsTrustedNode(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var out []interface{}
	err := _StaderOracle.contract.Call(opts, &out, "isTrustedNode", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsTrustedNode is a free data retrieval call binding the contract method 0x2f739b1d.
//
// Solidity: function isTrustedNode(address ) view returns(bool)
func (_StaderOracle *StaderOracleSession) IsTrustedNode(arg0 common.Address) (bool, error) {
	return _StaderOracle.Contract.IsTrustedNode(&_StaderOracle.CallOpts, arg0)
}

// IsTrustedNode is a free data retrieval call binding the contract method 0x2f739b1d.
//
// Solidity: function isTrustedNode(address ) view returns(bool)
func (_StaderOracle *StaderOracleCallerSession) IsTrustedNode(arg0 common.Address) (bool, error) {
	return _StaderOracle.Contract.IsTrustedNode(&_StaderOracle.CallOpts, arg0)
}

// LastReportedMAPDIndex is a free data retrieval call binding the contract method 0xa3737869.
//
// Solidity: function lastReportedMAPDIndex() view returns(uint256)
func (_StaderOracle *StaderOracleCaller) LastReportedMAPDIndex(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _StaderOracle.contract.Call(opts, &out, "lastReportedMAPDIndex")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// LastReportedMAPDIndex is a free data retrieval call binding the contract method 0xa3737869.
//
// Solidity: function lastReportedMAPDIndex() view returns(uint256)
func (_StaderOracle *StaderOracleSession) LastReportedMAPDIndex() (*big.Int, error) {
	return _StaderOracle.Contract.LastReportedMAPDIndex(&_StaderOracle.CallOpts)
}

// LastReportedMAPDIndex is a free data retrieval call binding the contract method 0xa3737869.
//
// Solidity: function lastReportedMAPDIndex() view returns(uint256)
func (_StaderOracle *StaderOracleCallerSession) LastReportedMAPDIndex() (*big.Int, error) {
	return _StaderOracle.Contract.LastReportedMAPDIndex(&_StaderOracle.CallOpts)
}

// LastReportedSDPriceData is a free data retrieval call binding the contract method 0xa8c3a3a8.
//
// Solidity: function lastReportedSDPriceData() view returns(uint256 reportingBlockNumber, uint256 sdPriceInETH)
func (_StaderOracle *StaderOracleCaller) LastReportedSDPriceData(opts *bind.CallOpts) (struct {
	ReportingBlockNumber *big.Int
	SdPriceInETH         *big.Int
}, error) {
	var out []interface{}
	err := _StaderOracle.contract.Call(opts, &out, "lastReportedSDPriceData")

	outstruct := new(struct {
		ReportingBlockNumber *big.Int
		SdPriceInETH         *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.ReportingBlockNumber = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.SdPriceInETH = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// LastReportedSDPriceData is a free data retrieval call binding the contract method 0xa8c3a3a8.
//
// Solidity: function lastReportedSDPriceData() view returns(uint256 reportingBlockNumber, uint256 sdPriceInETH)
func (_StaderOracle *StaderOracleSession) LastReportedSDPriceData() (struct {
	ReportingBlockNumber *big.Int
	SdPriceInETH         *big.Int
}, error) {
	return _StaderOracle.Contract.LastReportedSDPriceData(&_StaderOracle.CallOpts)
}

// LastReportedSDPriceData is a free data retrieval call binding the contract method 0xa8c3a3a8.
//
// Solidity: function lastReportedSDPriceData() view returns(uint256 reportingBlockNumber, uint256 sdPriceInETH)
func (_StaderOracle *StaderOracleCallerSession) LastReportedSDPriceData() (struct {
	ReportingBlockNumber *big.Int
	SdPriceInETH         *big.Int
}, error) {
	return _StaderOracle.Contract.LastReportedSDPriceData(&_StaderOracle.CallOpts)
}

// LastReportingBlockNumberForValidatorVerificationDetailByPoolId is a free data retrieval call binding the contract method 0xb17b4d86.
//
// Solidity: function lastReportingBlockNumberForValidatorVerificationDetailByPoolId(uint8 ) view returns(uint256)
func (_StaderOracle *StaderOracleCaller) LastReportingBlockNumberForValidatorVerificationDetailByPoolId(opts *bind.CallOpts, arg0 uint8) (*big.Int, error) {
	var out []interface{}
	err := _StaderOracle.contract.Call(opts, &out, "lastReportingBlockNumberForValidatorVerificationDetailByPoolId", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// LastReportingBlockNumberForValidatorVerificationDetailByPoolId is a free data retrieval call binding the contract method 0xb17b4d86.
//
// Solidity: function lastReportingBlockNumberForValidatorVerificationDetailByPoolId(uint8 ) view returns(uint256)
func (_StaderOracle *StaderOracleSession) LastReportingBlockNumberForValidatorVerificationDetailByPoolId(arg0 uint8) (*big.Int, error) {
	return _StaderOracle.Contract.LastReportingBlockNumberForValidatorVerificationDetailByPoolId(&_StaderOracle.CallOpts, arg0)
}

// LastReportingBlockNumberForValidatorVerificationDetailByPoolId is a free data retrieval call binding the contract method 0xb17b4d86.
//
// Solidity: function lastReportingBlockNumberForValidatorVerificationDetailByPoolId(uint8 ) view returns(uint256)
func (_StaderOracle *StaderOracleCallerSession) LastReportingBlockNumberForValidatorVerificationDetailByPoolId(arg0 uint8) (*big.Int, error) {
	return _StaderOracle.Contract.LastReportingBlockNumberForValidatorVerificationDetailByPoolId(&_StaderOracle.CallOpts, arg0)
}

// LastReportingBlockNumberForWithdrawnValidatorsByPoolId is a free data retrieval call binding the contract method 0xf00e0223.
//
// Solidity: function lastReportingBlockNumberForWithdrawnValidatorsByPoolId(uint8 ) view returns(uint256)
func (_StaderOracle *StaderOracleCaller) LastReportingBlockNumberForWithdrawnValidatorsByPoolId(opts *bind.CallOpts, arg0 uint8) (*big.Int, error) {
	var out []interface{}
	err := _StaderOracle.contract.Call(opts, &out, "lastReportingBlockNumberForWithdrawnValidatorsByPoolId", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// LastReportingBlockNumberForWithdrawnValidatorsByPoolId is a free data retrieval call binding the contract method 0xf00e0223.
//
// Solidity: function lastReportingBlockNumberForWithdrawnValidatorsByPoolId(uint8 ) view returns(uint256)
func (_StaderOracle *StaderOracleSession) LastReportingBlockNumberForWithdrawnValidatorsByPoolId(arg0 uint8) (*big.Int, error) {
	return _StaderOracle.Contract.LastReportingBlockNumberForWithdrawnValidatorsByPoolId(&_StaderOracle.CallOpts, arg0)
}

// LastReportingBlockNumberForWithdrawnValidatorsByPoolId is a free data retrieval call binding the contract method 0xf00e0223.
//
// Solidity: function lastReportingBlockNumberForWithdrawnValidatorsByPoolId(uint8 ) view returns(uint256)
func (_StaderOracle *StaderOracleCallerSession) LastReportingBlockNumberForWithdrawnValidatorsByPoolId(arg0 uint8) (*big.Int, error) {
	return _StaderOracle.Contract.LastReportingBlockNumberForWithdrawnValidatorsByPoolId(&_StaderOracle.CallOpts, arg0)
}

// LastTrustedNodeCountChangeBlock is a free data retrieval call binding the contract method 0x0989001c.
//
// Solidity: function lastTrustedNodeCountChangeBlock() view returns(uint256)
func (_StaderOracle *StaderOracleCaller) LastTrustedNodeCountChangeBlock(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _StaderOracle.contract.Call(opts, &out, "lastTrustedNodeCountChangeBlock")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// LastTrustedNodeCountChangeBlock is a free data retrieval call binding the contract method 0x0989001c.
//
// Solidity: function lastTrustedNodeCountChangeBlock() view returns(uint256)
func (_StaderOracle *StaderOracleSession) LastTrustedNodeCountChangeBlock() (*big.Int, error) {
	return _StaderOracle.Contract.LastTrustedNodeCountChangeBlock(&_StaderOracle.CallOpts)
}

// LastTrustedNodeCountChangeBlock is a free data retrieval call binding the contract method 0x0989001c.
//
// Solidity: function lastTrustedNodeCountChangeBlock() view returns(uint256)
func (_StaderOracle *StaderOracleCallerSession) LastTrustedNodeCountChangeBlock() (*big.Int, error) {
	return _StaderOracle.Contract.LastTrustedNodeCountChangeBlock(&_StaderOracle.CallOpts)
}

// MissedAttestationPenalty is a free data retrieval call binding the contract method 0x9773ee60.
//
// Solidity: function missedAttestationPenalty(bytes32 ) view returns(uint16)
func (_StaderOracle *StaderOracleCaller) MissedAttestationPenalty(opts *bind.CallOpts, arg0 [32]byte) (uint16, error) {
	var out []interface{}
	err := _StaderOracle.contract.Call(opts, &out, "missedAttestationPenalty", arg0)

	if err != nil {
		return *new(uint16), err
	}

	out0 := *abi.ConvertType(out[0], new(uint16)).(*uint16)

	return out0, err

}

// MissedAttestationPenalty is a free data retrieval call binding the contract method 0x9773ee60.
//
// Solidity: function missedAttestationPenalty(bytes32 ) view returns(uint16)
func (_StaderOracle *StaderOracleSession) MissedAttestationPenalty(arg0 [32]byte) (uint16, error) {
	return _StaderOracle.Contract.MissedAttestationPenalty(&_StaderOracle.CallOpts, arg0)
}

// MissedAttestationPenalty is a free data retrieval call binding the contract method 0x9773ee60.
//
// Solidity: function missedAttestationPenalty(bytes32 ) view returns(uint16)
func (_StaderOracle *StaderOracleCallerSession) MissedAttestationPenalty(arg0 [32]byte) (uint16, error) {
	return _StaderOracle.Contract.MissedAttestationPenalty(&_StaderOracle.CallOpts, arg0)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_StaderOracle *StaderOracleCaller) Paused(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _StaderOracle.contract.Call(opts, &out, "paused")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_StaderOracle *StaderOracleSession) Paused() (bool, error) {
	return _StaderOracle.Contract.Paused(&_StaderOracle.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_StaderOracle *StaderOracleCallerSession) Paused() (bool, error) {
	return _StaderOracle.Contract.Paused(&_StaderOracle.CallOpts)
}

// SafeMode is a free data retrieval call binding the contract method 0xabe3219c.
//
// Solidity: function safeMode() view returns(bool)
func (_StaderOracle *StaderOracleCaller) SafeMode(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _StaderOracle.contract.Call(opts, &out, "safeMode")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SafeMode is a free data retrieval call binding the contract method 0xabe3219c.
//
// Solidity: function safeMode() view returns(bool)
func (_StaderOracle *StaderOracleSession) SafeMode() (bool, error) {
	return _StaderOracle.Contract.SafeMode(&_StaderOracle.CallOpts)
}

// SafeMode is a free data retrieval call binding the contract method 0xabe3219c.
//
// Solidity: function safeMode() view returns(bool)
func (_StaderOracle *StaderOracleCallerSession) SafeMode() (bool, error) {
	return _StaderOracle.Contract.SafeMode(&_StaderOracle.CallOpts)
}

// StaderConfig is a free data retrieval call binding the contract method 0x490ffa35.
//
// Solidity: function staderConfig() view returns(address)
func (_StaderOracle *StaderOracleCaller) StaderConfig(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _StaderOracle.contract.Call(opts, &out, "staderConfig")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// StaderConfig is a free data retrieval call binding the contract method 0x490ffa35.
//
// Solidity: function staderConfig() view returns(address)
func (_StaderOracle *StaderOracleSession) StaderConfig() (common.Address, error) {
	return _StaderOracle.Contract.StaderConfig(&_StaderOracle.CallOpts)
}

// StaderConfig is a free data retrieval call binding the contract method 0x490ffa35.
//
// Solidity: function staderConfig() view returns(address)
func (_StaderOracle *StaderOracleCallerSession) StaderConfig() (common.Address, error) {
	return _StaderOracle.Contract.StaderConfig(&_StaderOracle.CallOpts)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_StaderOracle *StaderOracleCaller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _StaderOracle.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_StaderOracle *StaderOracleSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _StaderOracle.Contract.SupportsInterface(&_StaderOracle.CallOpts, interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_StaderOracle *StaderOracleCallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _StaderOracle.Contract.SupportsInterface(&_StaderOracle.CallOpts, interfaceId)
}

// TrustedNodeChangeCoolingPeriod is a free data retrieval call binding the contract method 0x61f00c17.
//
// Solidity: function trustedNodeChangeCoolingPeriod() view returns(uint256)
func (_StaderOracle *StaderOracleCaller) TrustedNodeChangeCoolingPeriod(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _StaderOracle.contract.Call(opts, &out, "trustedNodeChangeCoolingPeriod")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TrustedNodeChangeCoolingPeriod is a free data retrieval call binding the contract method 0x61f00c17.
//
// Solidity: function trustedNodeChangeCoolingPeriod() view returns(uint256)
func (_StaderOracle *StaderOracleSession) TrustedNodeChangeCoolingPeriod() (*big.Int, error) {
	return _StaderOracle.Contract.TrustedNodeChangeCoolingPeriod(&_StaderOracle.CallOpts)
}

// TrustedNodeChangeCoolingPeriod is a free data retrieval call binding the contract method 0x61f00c17.
//
// Solidity: function trustedNodeChangeCoolingPeriod() view returns(uint256)
func (_StaderOracle *StaderOracleCallerSession) TrustedNodeChangeCoolingPeriod() (*big.Int, error) {
	return _StaderOracle.Contract.TrustedNodeChangeCoolingPeriod(&_StaderOracle.CallOpts)
}

// TrustedNodesCount is a free data retrieval call binding the contract method 0xae815a04.
//
// Solidity: function trustedNodesCount() view returns(uint256)
func (_StaderOracle *StaderOracleCaller) TrustedNodesCount(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _StaderOracle.contract.Call(opts, &out, "trustedNodesCount")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TrustedNodesCount is a free data retrieval call binding the contract method 0xae815a04.
//
// Solidity: function trustedNodesCount() view returns(uint256)
func (_StaderOracle *StaderOracleSession) TrustedNodesCount() (*big.Int, error) {
	return _StaderOracle.Contract.TrustedNodesCount(&_StaderOracle.CallOpts)
}

// TrustedNodesCount is a free data retrieval call binding the contract method 0xae815a04.
//
// Solidity: function trustedNodesCount() view returns(uint256)
func (_StaderOracle *StaderOracleCallerSession) TrustedNodesCount() (*big.Int, error) {
	return _StaderOracle.Contract.TrustedNodesCount(&_StaderOracle.CallOpts)
}

// UpdateFrequencyMap is a free data retrieval call binding the contract method 0x7150bc5b.
//
// Solidity: function updateFrequencyMap(bytes32 ) view returns(uint256)
func (_StaderOracle *StaderOracleCaller) UpdateFrequencyMap(opts *bind.CallOpts, arg0 [32]byte) (*big.Int, error) {
	var out []interface{}
	err := _StaderOracle.contract.Call(opts, &out, "updateFrequencyMap", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// UpdateFrequencyMap is a free data retrieval call binding the contract method 0x7150bc5b.
//
// Solidity: function updateFrequencyMap(bytes32 ) view returns(uint256)
func (_StaderOracle *StaderOracleSession) UpdateFrequencyMap(arg0 [32]byte) (*big.Int, error) {
	return _StaderOracle.Contract.UpdateFrequencyMap(&_StaderOracle.CallOpts, arg0)
}

// UpdateFrequencyMap is a free data retrieval call binding the contract method 0x7150bc5b.
//
// Solidity: function updateFrequencyMap(bytes32 ) view returns(uint256)
func (_StaderOracle *StaderOracleCallerSession) UpdateFrequencyMap(arg0 [32]byte) (*big.Int, error) {
	return _StaderOracle.Contract.UpdateFrequencyMap(&_StaderOracle.CallOpts, arg0)
}

// ValidatorStats is a free data retrieval call binding the contract method 0x3b5eb03a.
//
// Solidity: function validatorStats() view returns(uint256 reportingBlockNumber, uint128 exitingValidatorsBalance, uint128 exitedValidatorsBalance, uint128 slashedValidatorsBalance, uint32 exitingValidatorsCount, uint32 exitedValidatorsCount, uint32 slashedValidatorsCount)
func (_StaderOracle *StaderOracleCaller) ValidatorStats(opts *bind.CallOpts) (struct {
	ReportingBlockNumber     *big.Int
	ExitingValidatorsBalance *big.Int
	ExitedValidatorsBalance  *big.Int
	SlashedValidatorsBalance *big.Int
	ExitingValidatorsCount   uint32
	ExitedValidatorsCount    uint32
	SlashedValidatorsCount   uint32
}, error) {
	var out []interface{}
	err := _StaderOracle.contract.Call(opts, &out, "validatorStats")

	outstruct := new(struct {
		ReportingBlockNumber     *big.Int
		ExitingValidatorsBalance *big.Int
		ExitedValidatorsBalance  *big.Int
		SlashedValidatorsBalance *big.Int
		ExitingValidatorsCount   uint32
		ExitedValidatorsCount    uint32
		SlashedValidatorsCount   uint32
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.ReportingBlockNumber = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.ExitingValidatorsBalance = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.ExitedValidatorsBalance = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	outstruct.SlashedValidatorsBalance = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)
	outstruct.ExitingValidatorsCount = *abi.ConvertType(out[4], new(uint32)).(*uint32)
	outstruct.ExitedValidatorsCount = *abi.ConvertType(out[5], new(uint32)).(*uint32)
	outstruct.SlashedValidatorsCount = *abi.ConvertType(out[6], new(uint32)).(*uint32)

	return *outstruct, err

}

// ValidatorStats is a free data retrieval call binding the contract method 0x3b5eb03a.
//
// Solidity: function validatorStats() view returns(uint256 reportingBlockNumber, uint128 exitingValidatorsBalance, uint128 exitedValidatorsBalance, uint128 slashedValidatorsBalance, uint32 exitingValidatorsCount, uint32 exitedValidatorsCount, uint32 slashedValidatorsCount)
func (_StaderOracle *StaderOracleSession) ValidatorStats() (struct {
	ReportingBlockNumber     *big.Int
	ExitingValidatorsBalance *big.Int
	ExitedValidatorsBalance  *big.Int
	SlashedValidatorsBalance *big.Int
	ExitingValidatorsCount   uint32
	ExitedValidatorsCount    uint32
	SlashedValidatorsCount   uint32
}, error) {
	return _StaderOracle.Contract.ValidatorStats(&_StaderOracle.CallOpts)
}

// ValidatorStats is a free data retrieval call binding the contract method 0x3b5eb03a.
//
// Solidity: function validatorStats() view returns(uint256 reportingBlockNumber, uint128 exitingValidatorsBalance, uint128 exitedValidatorsBalance, uint128 slashedValidatorsBalance, uint32 exitingValidatorsCount, uint32 exitedValidatorsCount, uint32 slashedValidatorsCount)
func (_StaderOracle *StaderOracleCallerSession) ValidatorStats() (struct {
	ReportingBlockNumber     *big.Int
	ExitingValidatorsBalance *big.Int
	ExitedValidatorsBalance  *big.Int
	SlashedValidatorsBalance *big.Int
	ExitingValidatorsCount   uint32
	ExitedValidatorsCount    uint32
	SlashedValidatorsCount   uint32
}, error) {
	return _StaderOracle.Contract.ValidatorStats(&_StaderOracle.CallOpts)
}

// AddTrustedNode is a paid mutator transaction binding the contract method 0xd6275dd7.
//
// Solidity: function addTrustedNode(address _nodeAddress) returns()
func (_StaderOracle *StaderOracleTransactor) AddTrustedNode(opts *bind.TransactOpts, _nodeAddress common.Address) (*types.Transaction, error) {
	return _StaderOracle.contract.Transact(opts, "addTrustedNode", _nodeAddress)
}

// AddTrustedNode is a paid mutator transaction binding the contract method 0xd6275dd7.
//
// Solidity: function addTrustedNode(address _nodeAddress) returns()
func (_StaderOracle *StaderOracleSession) AddTrustedNode(_nodeAddress common.Address) (*types.Transaction, error) {
	return _StaderOracle.Contract.AddTrustedNode(&_StaderOracle.TransactOpts, _nodeAddress)
}

// AddTrustedNode is a paid mutator transaction binding the contract method 0xd6275dd7.
//
// Solidity: function addTrustedNode(address _nodeAddress) returns()
func (_StaderOracle *StaderOracleTransactorSession) AddTrustedNode(_nodeAddress common.Address) (*types.Transaction, error) {
	return _StaderOracle.Contract.AddTrustedNode(&_StaderOracle.TransactOpts, _nodeAddress)
}

// CloseERInspectionMode is a paid mutator transaction binding the contract method 0x101b6e34.
//
// Solidity: function closeERInspectionMode() returns()
func (_StaderOracle *StaderOracleTransactor) CloseERInspectionMode(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _StaderOracle.contract.Transact(opts, "closeERInspectionMode")
}

// CloseERInspectionMode is a paid mutator transaction binding the contract method 0x101b6e34.
//
// Solidity: function closeERInspectionMode() returns()
func (_StaderOracle *StaderOracleSession) CloseERInspectionMode() (*types.Transaction, error) {
	return _StaderOracle.Contract.CloseERInspectionMode(&_StaderOracle.TransactOpts)
}

// CloseERInspectionMode is a paid mutator transaction binding the contract method 0x101b6e34.
//
// Solidity: function closeERInspectionMode() returns()
func (_StaderOracle *StaderOracleTransactorSession) CloseERInspectionMode() (*types.Transaction, error) {
	return _StaderOracle.Contract.CloseERInspectionMode(&_StaderOracle.TransactOpts)
}

// DisableERInspectionMode is a paid mutator transaction binding the contract method 0xe10025e6.
//
// Solidity: function disableERInspectionMode() returns()
func (_StaderOracle *StaderOracleTransactor) DisableERInspectionMode(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _StaderOracle.contract.Transact(opts, "disableERInspectionMode")
}

// DisableERInspectionMode is a paid mutator transaction binding the contract method 0xe10025e6.
//
// Solidity: function disableERInspectionMode() returns()
func (_StaderOracle *StaderOracleSession) DisableERInspectionMode() (*types.Transaction, error) {
	return _StaderOracle.Contract.DisableERInspectionMode(&_StaderOracle.TransactOpts)
}

// DisableERInspectionMode is a paid mutator transaction binding the contract method 0xe10025e6.
//
// Solidity: function disableERInspectionMode() returns()
func (_StaderOracle *StaderOracleTransactorSession) DisableERInspectionMode() (*types.Transaction, error) {
	return _StaderOracle.Contract.DisableERInspectionMode(&_StaderOracle.TransactOpts)
}

// DisableSafeMode is a paid mutator transaction binding the contract method 0xe514fe55.
//
// Solidity: function disableSafeMode() returns()
func (_StaderOracle *StaderOracleTransactor) DisableSafeMode(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _StaderOracle.contract.Transact(opts, "disableSafeMode")
}

// DisableSafeMode is a paid mutator transaction binding the contract method 0xe514fe55.
//
// Solidity: function disableSafeMode() returns()
func (_StaderOracle *StaderOracleSession) DisableSafeMode() (*types.Transaction, error) {
	return _StaderOracle.Contract.DisableSafeMode(&_StaderOracle.TransactOpts)
}

// DisableSafeMode is a paid mutator transaction binding the contract method 0xe514fe55.
//
// Solidity: function disableSafeMode() returns()
func (_StaderOracle *StaderOracleTransactorSession) DisableSafeMode() (*types.Transaction, error) {
	return _StaderOracle.Contract.DisableSafeMode(&_StaderOracle.TransactOpts)
}

// EnableSafeMode is a paid mutator transaction binding the contract method 0x4f560abd.
//
// Solidity: function enableSafeMode() returns()
func (_StaderOracle *StaderOracleTransactor) EnableSafeMode(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _StaderOracle.contract.Transact(opts, "enableSafeMode")
}

// EnableSafeMode is a paid mutator transaction binding the contract method 0x4f560abd.
//
// Solidity: function enableSafeMode() returns()
func (_StaderOracle *StaderOracleSession) EnableSafeMode() (*types.Transaction, error) {
	return _StaderOracle.Contract.EnableSafeMode(&_StaderOracle.TransactOpts)
}

// EnableSafeMode is a paid mutator transaction binding the contract method 0x4f560abd.
//
// Solidity: function enableSafeMode() returns()
func (_StaderOracle *StaderOracleTransactorSession) EnableSafeMode() (*types.Transaction, error) {
	return _StaderOracle.Contract.EnableSafeMode(&_StaderOracle.TransactOpts)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_StaderOracle *StaderOracleTransactor) GrantRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _StaderOracle.contract.Transact(opts, "grantRole", role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_StaderOracle *StaderOracleSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _StaderOracle.Contract.GrantRole(&_StaderOracle.TransactOpts, role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_StaderOracle *StaderOracleTransactorSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _StaderOracle.Contract.GrantRole(&_StaderOracle.TransactOpts, role, account)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_StaderOracle *StaderOracleTransactor) Pause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _StaderOracle.contract.Transact(opts, "pause")
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_StaderOracle *StaderOracleSession) Pause() (*types.Transaction, error) {
	return _StaderOracle.Contract.Pause(&_StaderOracle.TransactOpts)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_StaderOracle *StaderOracleTransactorSession) Pause() (*types.Transaction, error) {
	return _StaderOracle.Contract.Pause(&_StaderOracle.TransactOpts)
}

// RemoveTrustedNode is a paid mutator transaction binding the contract method 0x52e0fc80.
//
// Solidity: function removeTrustedNode(address _nodeAddress) returns()
func (_StaderOracle *StaderOracleTransactor) RemoveTrustedNode(opts *bind.TransactOpts, _nodeAddress common.Address) (*types.Transaction, error) {
	return _StaderOracle.contract.Transact(opts, "removeTrustedNode", _nodeAddress)
}

// RemoveTrustedNode is a paid mutator transaction binding the contract method 0x52e0fc80.
//
// Solidity: function removeTrustedNode(address _nodeAddress) returns()
func (_StaderOracle *StaderOracleSession) RemoveTrustedNode(_nodeAddress common.Address) (*types.Transaction, error) {
	return _StaderOracle.Contract.RemoveTrustedNode(&_StaderOracle.TransactOpts, _nodeAddress)
}

// RemoveTrustedNode is a paid mutator transaction binding the contract method 0x52e0fc80.
//
// Solidity: function removeTrustedNode(address _nodeAddress) returns()
func (_StaderOracle *StaderOracleTransactorSession) RemoveTrustedNode(_nodeAddress common.Address) (*types.Transaction, error) {
	return _StaderOracle.Contract.RemoveTrustedNode(&_StaderOracle.TransactOpts, _nodeAddress)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address account) returns()
func (_StaderOracle *StaderOracleTransactor) RenounceRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _StaderOracle.contract.Transact(opts, "renounceRole", role, account)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address account) returns()
func (_StaderOracle *StaderOracleSession) RenounceRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _StaderOracle.Contract.RenounceRole(&_StaderOracle.TransactOpts, role, account)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address account) returns()
func (_StaderOracle *StaderOracleTransactorSession) RenounceRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _StaderOracle.Contract.RenounceRole(&_StaderOracle.TransactOpts, role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_StaderOracle *StaderOracleTransactor) RevokeRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _StaderOracle.contract.Transact(opts, "revokeRole", role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_StaderOracle *StaderOracleSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _StaderOracle.Contract.RevokeRole(&_StaderOracle.TransactOpts, role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_StaderOracle *StaderOracleTransactorSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _StaderOracle.Contract.RevokeRole(&_StaderOracle.TransactOpts, role, account)
}

// SetERUpdateFrequency is a paid mutator transaction binding the contract method 0xd0a8f679.
//
// Solidity: function setERUpdateFrequency(uint256 _updateFrequency) returns()
func (_StaderOracle *StaderOracleTransactor) SetERUpdateFrequency(opts *bind.TransactOpts, _updateFrequency *big.Int) (*types.Transaction, error) {
	return _StaderOracle.contract.Transact(opts, "setERUpdateFrequency", _updateFrequency)
}

// SetERUpdateFrequency is a paid mutator transaction binding the contract method 0xd0a8f679.
//
// Solidity: function setERUpdateFrequency(uint256 _updateFrequency) returns()
func (_StaderOracle *StaderOracleSession) SetERUpdateFrequency(_updateFrequency *big.Int) (*types.Transaction, error) {
	return _StaderOracle.Contract.SetERUpdateFrequency(&_StaderOracle.TransactOpts, _updateFrequency)
}

// SetERUpdateFrequency is a paid mutator transaction binding the contract method 0xd0a8f679.
//
// Solidity: function setERUpdateFrequency(uint256 _updateFrequency) returns()
func (_StaderOracle *StaderOracleTransactorSession) SetERUpdateFrequency(_updateFrequency *big.Int) (*types.Transaction, error) {
	return _StaderOracle.Contract.SetERUpdateFrequency(&_StaderOracle.TransactOpts, _updateFrequency)
}

// SetMissedAttestationPenaltyUpdateFrequency is a paid mutator transaction binding the contract method 0xf51c0fe7.
//
// Solidity: function setMissedAttestationPenaltyUpdateFrequency(uint256 _updateFrequency) returns()
func (_StaderOracle *StaderOracleTransactor) SetMissedAttestationPenaltyUpdateFrequency(opts *bind.TransactOpts, _updateFrequency *big.Int) (*types.Transaction, error) {
	return _StaderOracle.contract.Transact(opts, "setMissedAttestationPenaltyUpdateFrequency", _updateFrequency)
}

// SetMissedAttestationPenaltyUpdateFrequency is a paid mutator transaction binding the contract method 0xf51c0fe7.
//
// Solidity: function setMissedAttestationPenaltyUpdateFrequency(uint256 _updateFrequency) returns()
func (_StaderOracle *StaderOracleSession) SetMissedAttestationPenaltyUpdateFrequency(_updateFrequency *big.Int) (*types.Transaction, error) {
	return _StaderOracle.Contract.SetMissedAttestationPenaltyUpdateFrequency(&_StaderOracle.TransactOpts, _updateFrequency)
}

// SetMissedAttestationPenaltyUpdateFrequency is a paid mutator transaction binding the contract method 0xf51c0fe7.
//
// Solidity: function setMissedAttestationPenaltyUpdateFrequency(uint256 _updateFrequency) returns()
func (_StaderOracle *StaderOracleTransactorSession) SetMissedAttestationPenaltyUpdateFrequency(_updateFrequency *big.Int) (*types.Transaction, error) {
	return _StaderOracle.Contract.SetMissedAttestationPenaltyUpdateFrequency(&_StaderOracle.TransactOpts, _updateFrequency)
}

// SetSDPriceUpdateFrequency is a paid mutator transaction binding the contract method 0x749f7d8a.
//
// Solidity: function setSDPriceUpdateFrequency(uint256 _updateFrequency) returns()
func (_StaderOracle *StaderOracleTransactor) SetSDPriceUpdateFrequency(opts *bind.TransactOpts, _updateFrequency *big.Int) (*types.Transaction, error) {
	return _StaderOracle.contract.Transact(opts, "setSDPriceUpdateFrequency", _updateFrequency)
}

// SetSDPriceUpdateFrequency is a paid mutator transaction binding the contract method 0x749f7d8a.
//
// Solidity: function setSDPriceUpdateFrequency(uint256 _updateFrequency) returns()
func (_StaderOracle *StaderOracleSession) SetSDPriceUpdateFrequency(_updateFrequency *big.Int) (*types.Transaction, error) {
	return _StaderOracle.Contract.SetSDPriceUpdateFrequency(&_StaderOracle.TransactOpts, _updateFrequency)
}

// SetSDPriceUpdateFrequency is a paid mutator transaction binding the contract method 0x749f7d8a.
//
// Solidity: function setSDPriceUpdateFrequency(uint256 _updateFrequency) returns()
func (_StaderOracle *StaderOracleTransactorSession) SetSDPriceUpdateFrequency(_updateFrequency *big.Int) (*types.Transaction, error) {
	return _StaderOracle.Contract.SetSDPriceUpdateFrequency(&_StaderOracle.TransactOpts, _updateFrequency)
}

// SetValidatorStatsUpdateFrequency is a paid mutator transaction binding the contract method 0x844007fe.
//
// Solidity: function setValidatorStatsUpdateFrequency(uint256 _updateFrequency) returns()
func (_StaderOracle *StaderOracleTransactor) SetValidatorStatsUpdateFrequency(opts *bind.TransactOpts, _updateFrequency *big.Int) (*types.Transaction, error) {
	return _StaderOracle.contract.Transact(opts, "setValidatorStatsUpdateFrequency", _updateFrequency)
}

// SetValidatorStatsUpdateFrequency is a paid mutator transaction binding the contract method 0x844007fe.
//
// Solidity: function setValidatorStatsUpdateFrequency(uint256 _updateFrequency) returns()
func (_StaderOracle *StaderOracleSession) SetValidatorStatsUpdateFrequency(_updateFrequency *big.Int) (*types.Transaction, error) {
	return _StaderOracle.Contract.SetValidatorStatsUpdateFrequency(&_StaderOracle.TransactOpts, _updateFrequency)
}

// SetValidatorStatsUpdateFrequency is a paid mutator transaction binding the contract method 0x844007fe.
//
// Solidity: function setValidatorStatsUpdateFrequency(uint256 _updateFrequency) returns()
func (_StaderOracle *StaderOracleTransactorSession) SetValidatorStatsUpdateFrequency(_updateFrequency *big.Int) (*types.Transaction, error) {
	return _StaderOracle.Contract.SetValidatorStatsUpdateFrequency(&_StaderOracle.TransactOpts, _updateFrequency)
}

// SetValidatorVerificationDetailUpdateFrequency is a paid mutator transaction binding the contract method 0xea18568b.
//
// Solidity: function setValidatorVerificationDetailUpdateFrequency(uint256 _updateFrequency) returns()
func (_StaderOracle *StaderOracleTransactor) SetValidatorVerificationDetailUpdateFrequency(opts *bind.TransactOpts, _updateFrequency *big.Int) (*types.Transaction, error) {
	return _StaderOracle.contract.Transact(opts, "setValidatorVerificationDetailUpdateFrequency", _updateFrequency)
}

// SetValidatorVerificationDetailUpdateFrequency is a paid mutator transaction binding the contract method 0xea18568b.
//
// Solidity: function setValidatorVerificationDetailUpdateFrequency(uint256 _updateFrequency) returns()
func (_StaderOracle *StaderOracleSession) SetValidatorVerificationDetailUpdateFrequency(_updateFrequency *big.Int) (*types.Transaction, error) {
	return _StaderOracle.Contract.SetValidatorVerificationDetailUpdateFrequency(&_StaderOracle.TransactOpts, _updateFrequency)
}

// SetValidatorVerificationDetailUpdateFrequency is a paid mutator transaction binding the contract method 0xea18568b.
//
// Solidity: function setValidatorVerificationDetailUpdateFrequency(uint256 _updateFrequency) returns()
func (_StaderOracle *StaderOracleTransactorSession) SetValidatorVerificationDetailUpdateFrequency(_updateFrequency *big.Int) (*types.Transaction, error) {
	return _StaderOracle.Contract.SetValidatorVerificationDetailUpdateFrequency(&_StaderOracle.TransactOpts, _updateFrequency)
}

// SetWithdrawnValidatorsUpdateFrequency is a paid mutator transaction binding the contract method 0xc06a6201.
//
// Solidity: function setWithdrawnValidatorsUpdateFrequency(uint256 _updateFrequency) returns()
func (_StaderOracle *StaderOracleTransactor) SetWithdrawnValidatorsUpdateFrequency(opts *bind.TransactOpts, _updateFrequency *big.Int) (*types.Transaction, error) {
	return _StaderOracle.contract.Transact(opts, "setWithdrawnValidatorsUpdateFrequency", _updateFrequency)
}

// SetWithdrawnValidatorsUpdateFrequency is a paid mutator transaction binding the contract method 0xc06a6201.
//
// Solidity: function setWithdrawnValidatorsUpdateFrequency(uint256 _updateFrequency) returns()
func (_StaderOracle *StaderOracleSession) SetWithdrawnValidatorsUpdateFrequency(_updateFrequency *big.Int) (*types.Transaction, error) {
	return _StaderOracle.Contract.SetWithdrawnValidatorsUpdateFrequency(&_StaderOracle.TransactOpts, _updateFrequency)
}

// SetWithdrawnValidatorsUpdateFrequency is a paid mutator transaction binding the contract method 0xc06a6201.
//
// Solidity: function setWithdrawnValidatorsUpdateFrequency(uint256 _updateFrequency) returns()
func (_StaderOracle *StaderOracleTransactorSession) SetWithdrawnValidatorsUpdateFrequency(_updateFrequency *big.Int) (*types.Transaction, error) {
	return _StaderOracle.Contract.SetWithdrawnValidatorsUpdateFrequency(&_StaderOracle.TransactOpts, _updateFrequency)
}

// SubmitExchangeRateData is a paid mutator transaction binding the contract method 0x818c8b26.
//
// Solidity: function submitExchangeRateData((uint256,uint256,uint256) _exchangeRate) returns()
func (_StaderOracle *StaderOracleTransactor) SubmitExchangeRateData(opts *bind.TransactOpts, _exchangeRate ExchangeRate) (*types.Transaction, error) {
	return _StaderOracle.contract.Transact(opts, "submitExchangeRateData", _exchangeRate)
}

// SubmitExchangeRateData is a paid mutator transaction binding the contract method 0x818c8b26.
//
// Solidity: function submitExchangeRateData((uint256,uint256,uint256) _exchangeRate) returns()
func (_StaderOracle *StaderOracleSession) SubmitExchangeRateData(_exchangeRate ExchangeRate) (*types.Transaction, error) {
	return _StaderOracle.Contract.SubmitExchangeRateData(&_StaderOracle.TransactOpts, _exchangeRate)
}

// SubmitExchangeRateData is a paid mutator transaction binding the contract method 0x818c8b26.
//
// Solidity: function submitExchangeRateData((uint256,uint256,uint256) _exchangeRate) returns()
func (_StaderOracle *StaderOracleTransactorSession) SubmitExchangeRateData(_exchangeRate ExchangeRate) (*types.Transaction, error) {
	return _StaderOracle.Contract.SubmitExchangeRateData(&_StaderOracle.TransactOpts, _exchangeRate)
}

// SubmitMissedAttestationPenalties is a paid mutator transaction binding the contract method 0x67fbf731.
//
// Solidity: function submitMissedAttestationPenalties((uint256,uint256,bytes[]) _mapd) returns()
func (_StaderOracle *StaderOracleTransactor) SubmitMissedAttestationPenalties(opts *bind.TransactOpts, _mapd MissedAttestationPenaltyData) (*types.Transaction, error) {
	return _StaderOracle.contract.Transact(opts, "submitMissedAttestationPenalties", _mapd)
}

// SubmitMissedAttestationPenalties is a paid mutator transaction binding the contract method 0x67fbf731.
//
// Solidity: function submitMissedAttestationPenalties((uint256,uint256,bytes[]) _mapd) returns()
func (_StaderOracle *StaderOracleSession) SubmitMissedAttestationPenalties(_mapd MissedAttestationPenaltyData) (*types.Transaction, error) {
	return _StaderOracle.Contract.SubmitMissedAttestationPenalties(&_StaderOracle.TransactOpts, _mapd)
}

// SubmitMissedAttestationPenalties is a paid mutator transaction binding the contract method 0x67fbf731.
//
// Solidity: function submitMissedAttestationPenalties((uint256,uint256,bytes[]) _mapd) returns()
func (_StaderOracle *StaderOracleTransactorSession) SubmitMissedAttestationPenalties(_mapd MissedAttestationPenaltyData) (*types.Transaction, error) {
	return _StaderOracle.Contract.SubmitMissedAttestationPenalties(&_StaderOracle.TransactOpts, _mapd)
}

// SubmitSDPrice is a paid mutator transaction binding the contract method 0xf6a3c090.
//
// Solidity: function submitSDPrice((uint256,uint256) _sdPriceData) returns()
func (_StaderOracle *StaderOracleTransactor) SubmitSDPrice(opts *bind.TransactOpts, _sdPriceData SDPriceData) (*types.Transaction, error) {
	return _StaderOracle.contract.Transact(opts, "submitSDPrice", _sdPriceData)
}

// SubmitSDPrice is a paid mutator transaction binding the contract method 0xf6a3c090.
//
// Solidity: function submitSDPrice((uint256,uint256) _sdPriceData) returns()
func (_StaderOracle *StaderOracleSession) SubmitSDPrice(_sdPriceData SDPriceData) (*types.Transaction, error) {
	return _StaderOracle.Contract.SubmitSDPrice(&_StaderOracle.TransactOpts, _sdPriceData)
}

// SubmitSDPrice is a paid mutator transaction binding the contract method 0xf6a3c090.
//
// Solidity: function submitSDPrice((uint256,uint256) _sdPriceData) returns()
func (_StaderOracle *StaderOracleTransactorSession) SubmitSDPrice(_sdPriceData SDPriceData) (*types.Transaction, error) {
	return _StaderOracle.Contract.SubmitSDPrice(&_StaderOracle.TransactOpts, _sdPriceData)
}

// SubmitSocializingRewardsMerkleRoot is a paid mutator transaction binding the contract method 0xae541d65.
//
// Solidity: function submitSocializingRewardsMerkleRoot((uint256,uint256,bytes32,uint8,uint256,uint256,uint256,uint256) _rewardsData) returns()
func (_StaderOracle *StaderOracleTransactor) SubmitSocializingRewardsMerkleRoot(opts *bind.TransactOpts, _rewardsData RewardsData) (*types.Transaction, error) {
	return _StaderOracle.contract.Transact(opts, "submitSocializingRewardsMerkleRoot", _rewardsData)
}

// SubmitSocializingRewardsMerkleRoot is a paid mutator transaction binding the contract method 0xae541d65.
//
// Solidity: function submitSocializingRewardsMerkleRoot((uint256,uint256,bytes32,uint8,uint256,uint256,uint256,uint256) _rewardsData) returns()
func (_StaderOracle *StaderOracleSession) SubmitSocializingRewardsMerkleRoot(_rewardsData RewardsData) (*types.Transaction, error) {
	return _StaderOracle.Contract.SubmitSocializingRewardsMerkleRoot(&_StaderOracle.TransactOpts, _rewardsData)
}

// SubmitSocializingRewardsMerkleRoot is a paid mutator transaction binding the contract method 0xae541d65.
//
// Solidity: function submitSocializingRewardsMerkleRoot((uint256,uint256,bytes32,uint8,uint256,uint256,uint256,uint256) _rewardsData) returns()
func (_StaderOracle *StaderOracleTransactorSession) SubmitSocializingRewardsMerkleRoot(_rewardsData RewardsData) (*types.Transaction, error) {
	return _StaderOracle.Contract.SubmitSocializingRewardsMerkleRoot(&_StaderOracle.TransactOpts, _rewardsData)
}

// SubmitValidatorStats is a paid mutator transaction binding the contract method 0x5c7ccd3b.
//
// Solidity: function submitValidatorStats((uint256,uint128,uint128,uint128,uint32,uint32,uint32) _validatorStats) returns()
func (_StaderOracle *StaderOracleTransactor) SubmitValidatorStats(opts *bind.TransactOpts, _validatorStats ValidatorStats) (*types.Transaction, error) {
	return _StaderOracle.contract.Transact(opts, "submitValidatorStats", _validatorStats)
}

// SubmitValidatorStats is a paid mutator transaction binding the contract method 0x5c7ccd3b.
//
// Solidity: function submitValidatorStats((uint256,uint128,uint128,uint128,uint32,uint32,uint32) _validatorStats) returns()
func (_StaderOracle *StaderOracleSession) SubmitValidatorStats(_validatorStats ValidatorStats) (*types.Transaction, error) {
	return _StaderOracle.Contract.SubmitValidatorStats(&_StaderOracle.TransactOpts, _validatorStats)
}

// SubmitValidatorStats is a paid mutator transaction binding the contract method 0x5c7ccd3b.
//
// Solidity: function submitValidatorStats((uint256,uint128,uint128,uint128,uint32,uint32,uint32) _validatorStats) returns()
func (_StaderOracle *StaderOracleTransactorSession) SubmitValidatorStats(_validatorStats ValidatorStats) (*types.Transaction, error) {
	return _StaderOracle.Contract.SubmitValidatorStats(&_StaderOracle.TransactOpts, _validatorStats)
}

// SubmitValidatorVerificationDetail is a paid mutator transaction binding the contract method 0x735efb96.
//
// Solidity: function submitValidatorVerificationDetail((uint8,uint256,bytes[],bytes[],bytes[]) _validatorVerificationDetail) returns()
func (_StaderOracle *StaderOracleTransactor) SubmitValidatorVerificationDetail(opts *bind.TransactOpts, _validatorVerificationDetail ValidatorVerificationDetail) (*types.Transaction, error) {
	return _StaderOracle.contract.Transact(opts, "submitValidatorVerificationDetail", _validatorVerificationDetail)
}

// SubmitValidatorVerificationDetail is a paid mutator transaction binding the contract method 0x735efb96.
//
// Solidity: function submitValidatorVerificationDetail((uint8,uint256,bytes[],bytes[],bytes[]) _validatorVerificationDetail) returns()
func (_StaderOracle *StaderOracleSession) SubmitValidatorVerificationDetail(_validatorVerificationDetail ValidatorVerificationDetail) (*types.Transaction, error) {
	return _StaderOracle.Contract.SubmitValidatorVerificationDetail(&_StaderOracle.TransactOpts, _validatorVerificationDetail)
}

// SubmitValidatorVerificationDetail is a paid mutator transaction binding the contract method 0x735efb96.
//
// Solidity: function submitValidatorVerificationDetail((uint8,uint256,bytes[],bytes[],bytes[]) _validatorVerificationDetail) returns()
func (_StaderOracle *StaderOracleTransactorSession) SubmitValidatorVerificationDetail(_validatorVerificationDetail ValidatorVerificationDetail) (*types.Transaction, error) {
	return _StaderOracle.Contract.SubmitValidatorVerificationDetail(&_StaderOracle.TransactOpts, _validatorVerificationDetail)
}

// SubmitWithdrawnValidators is a paid mutator transaction binding the contract method 0xa220c2d3.
//
// Solidity: function submitWithdrawnValidators((uint8,uint256,bytes[]) _withdrawnValidators) returns()
func (_StaderOracle *StaderOracleTransactor) SubmitWithdrawnValidators(opts *bind.TransactOpts, _withdrawnValidators WithdrawnValidators) (*types.Transaction, error) {
	return _StaderOracle.contract.Transact(opts, "submitWithdrawnValidators", _withdrawnValidators)
}

// SubmitWithdrawnValidators is a paid mutator transaction binding the contract method 0xa220c2d3.
//
// Solidity: function submitWithdrawnValidators((uint8,uint256,bytes[]) _withdrawnValidators) returns()
func (_StaderOracle *StaderOracleSession) SubmitWithdrawnValidators(_withdrawnValidators WithdrawnValidators) (*types.Transaction, error) {
	return _StaderOracle.Contract.SubmitWithdrawnValidators(&_StaderOracle.TransactOpts, _withdrawnValidators)
}

// SubmitWithdrawnValidators is a paid mutator transaction binding the contract method 0xa220c2d3.
//
// Solidity: function submitWithdrawnValidators((uint8,uint256,bytes[]) _withdrawnValidators) returns()
func (_StaderOracle *StaderOracleTransactorSession) SubmitWithdrawnValidators(_withdrawnValidators WithdrawnValidators) (*types.Transaction, error) {
	return _StaderOracle.Contract.SubmitWithdrawnValidators(&_StaderOracle.TransactOpts, _withdrawnValidators)
}

// TogglePORFeedBasedERData is a paid mutator transaction binding the contract method 0x712033eb.
//
// Solidity: function togglePORFeedBasedERData() returns()
func (_StaderOracle *StaderOracleTransactor) TogglePORFeedBasedERData(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _StaderOracle.contract.Transact(opts, "togglePORFeedBasedERData")
}

// TogglePORFeedBasedERData is a paid mutator transaction binding the contract method 0x712033eb.
//
// Solidity: function togglePORFeedBasedERData() returns()
func (_StaderOracle *StaderOracleSession) TogglePORFeedBasedERData() (*types.Transaction, error) {
	return _StaderOracle.Contract.TogglePORFeedBasedERData(&_StaderOracle.TransactOpts)
}

// TogglePORFeedBasedERData is a paid mutator transaction binding the contract method 0x712033eb.
//
// Solidity: function togglePORFeedBasedERData() returns()
func (_StaderOracle *StaderOracleTransactorSession) TogglePORFeedBasedERData() (*types.Transaction, error) {
	return _StaderOracle.Contract.TogglePORFeedBasedERData(&_StaderOracle.TransactOpts)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_StaderOracle *StaderOracleTransactor) Unpause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _StaderOracle.contract.Transact(opts, "unpause")
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_StaderOracle *StaderOracleSession) Unpause() (*types.Transaction, error) {
	return _StaderOracle.Contract.Unpause(&_StaderOracle.TransactOpts)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_StaderOracle *StaderOracleTransactorSession) Unpause() (*types.Transaction, error) {
	return _StaderOracle.Contract.Unpause(&_StaderOracle.TransactOpts)
}

// UpdateERChangeLimit is a paid mutator transaction binding the contract method 0x8ca8703c.
//
// Solidity: function updateERChangeLimit(uint256 _erChangeLimit) returns()
func (_StaderOracle *StaderOracleTransactor) UpdateERChangeLimit(opts *bind.TransactOpts, _erChangeLimit *big.Int) (*types.Transaction, error) {
	return _StaderOracle.contract.Transact(opts, "updateERChangeLimit", _erChangeLimit)
}

// UpdateERChangeLimit is a paid mutator transaction binding the contract method 0x8ca8703c.
//
// Solidity: function updateERChangeLimit(uint256 _erChangeLimit) returns()
func (_StaderOracle *StaderOracleSession) UpdateERChangeLimit(_erChangeLimit *big.Int) (*types.Transaction, error) {
	return _StaderOracle.Contract.UpdateERChangeLimit(&_StaderOracle.TransactOpts, _erChangeLimit)
}

// UpdateERChangeLimit is a paid mutator transaction binding the contract method 0x8ca8703c.
//
// Solidity: function updateERChangeLimit(uint256 _erChangeLimit) returns()
func (_StaderOracle *StaderOracleTransactorSession) UpdateERChangeLimit(_erChangeLimit *big.Int) (*types.Transaction, error) {
	return _StaderOracle.Contract.UpdateERChangeLimit(&_StaderOracle.TransactOpts, _erChangeLimit)
}

// UpdateERFromPORFeed is a paid mutator transaction binding the contract method 0x9bfdf9a4.
//
// Solidity: function updateERFromPORFeed() returns()
func (_StaderOracle *StaderOracleTransactor) UpdateERFromPORFeed(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _StaderOracle.contract.Transact(opts, "updateERFromPORFeed")
}

// UpdateERFromPORFeed is a paid mutator transaction binding the contract method 0x9bfdf9a4.
//
// Solidity: function updateERFromPORFeed() returns()
func (_StaderOracle *StaderOracleSession) UpdateERFromPORFeed() (*types.Transaction, error) {
	return _StaderOracle.Contract.UpdateERFromPORFeed(&_StaderOracle.TransactOpts)
}

// UpdateERFromPORFeed is a paid mutator transaction binding the contract method 0x9bfdf9a4.
//
// Solidity: function updateERFromPORFeed() returns()
func (_StaderOracle *StaderOracleTransactorSession) UpdateERFromPORFeed() (*types.Transaction, error) {
	return _StaderOracle.Contract.UpdateERFromPORFeed(&_StaderOracle.TransactOpts)
}

// UpdateStaderConfig is a paid mutator transaction binding the contract method 0x9ee804cb.
//
// Solidity: function updateStaderConfig(address _staderConfig) returns()
func (_StaderOracle *StaderOracleTransactor) UpdateStaderConfig(opts *bind.TransactOpts, _staderConfig common.Address) (*types.Transaction, error) {
	return _StaderOracle.contract.Transact(opts, "updateStaderConfig", _staderConfig)
}

// UpdateStaderConfig is a paid mutator transaction binding the contract method 0x9ee804cb.
//
// Solidity: function updateStaderConfig(address _staderConfig) returns()
func (_StaderOracle *StaderOracleSession) UpdateStaderConfig(_staderConfig common.Address) (*types.Transaction, error) {
	return _StaderOracle.Contract.UpdateStaderConfig(&_StaderOracle.TransactOpts, _staderConfig)
}

// UpdateStaderConfig is a paid mutator transaction binding the contract method 0x9ee804cb.
//
// Solidity: function updateStaderConfig(address _staderConfig) returns()
func (_StaderOracle *StaderOracleTransactorSession) UpdateStaderConfig(_staderConfig common.Address) (*types.Transaction, error) {
	return _StaderOracle.Contract.UpdateStaderConfig(&_StaderOracle.TransactOpts, _staderConfig)
}

// UpdateTrustedNodeChangeCoolingPeriod is a paid mutator transaction binding the contract method 0x962c1e05.
//
// Solidity: function updateTrustedNodeChangeCoolingPeriod(uint256 _trustedNodeChangeCoolingPeriod) returns()
func (_StaderOracle *StaderOracleTransactor) UpdateTrustedNodeChangeCoolingPeriod(opts *bind.TransactOpts, _trustedNodeChangeCoolingPeriod *big.Int) (*types.Transaction, error) {
	return _StaderOracle.contract.Transact(opts, "updateTrustedNodeChangeCoolingPeriod", _trustedNodeChangeCoolingPeriod)
}

// UpdateTrustedNodeChangeCoolingPeriod is a paid mutator transaction binding the contract method 0x962c1e05.
//
// Solidity: function updateTrustedNodeChangeCoolingPeriod(uint256 _trustedNodeChangeCoolingPeriod) returns()
func (_StaderOracle *StaderOracleSession) UpdateTrustedNodeChangeCoolingPeriod(_trustedNodeChangeCoolingPeriod *big.Int) (*types.Transaction, error) {
	return _StaderOracle.Contract.UpdateTrustedNodeChangeCoolingPeriod(&_StaderOracle.TransactOpts, _trustedNodeChangeCoolingPeriod)
}

// UpdateTrustedNodeChangeCoolingPeriod is a paid mutator transaction binding the contract method 0x962c1e05.
//
// Solidity: function updateTrustedNodeChangeCoolingPeriod(uint256 _trustedNodeChangeCoolingPeriod) returns()
func (_StaderOracle *StaderOracleTransactorSession) UpdateTrustedNodeChangeCoolingPeriod(_trustedNodeChangeCoolingPeriod *big.Int) (*types.Transaction, error) {
	return _StaderOracle.Contract.UpdateTrustedNodeChangeCoolingPeriod(&_StaderOracle.TransactOpts, _trustedNodeChangeCoolingPeriod)
}

// StaderOracleERDataSourceToggledIterator is returned from FilterERDataSourceToggled and is used to iterate over the raw logs and unpacked data for ERDataSourceToggled events raised by the StaderOracle contract.
type StaderOracleERDataSourceToggledIterator struct {
	Event *StaderOracleERDataSourceToggled // Event containing the contract specifics and raw log

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
func (it *StaderOracleERDataSourceToggledIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StaderOracleERDataSourceToggled)
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
		it.Event = new(StaderOracleERDataSourceToggled)
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
func (it *StaderOracleERDataSourceToggledIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StaderOracleERDataSourceToggledIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StaderOracleERDataSourceToggled represents a ERDataSourceToggled event raised by the StaderOracle contract.
type StaderOracleERDataSourceToggled struct {
	IsPORBasedERData bool
	Raw              types.Log // Blockchain specific contextual infos
}

// FilterERDataSourceToggled is a free log retrieval operation binding the contract event 0xc59a5de02f9d69be770ff0d61bbc894968433bb599f9fd9c2016e149c509c5e5.
//
// Solidity: event ERDataSourceToggled(bool isPORBasedERData)
func (_StaderOracle *StaderOracleFilterer) FilterERDataSourceToggled(opts *bind.FilterOpts) (*StaderOracleERDataSourceToggledIterator, error) {

	logs, sub, err := _StaderOracle.contract.FilterLogs(opts, "ERDataSourceToggled")
	if err != nil {
		return nil, err
	}
	return &StaderOracleERDataSourceToggledIterator{contract: _StaderOracle.contract, event: "ERDataSourceToggled", logs: logs, sub: sub}, nil
}

// WatchERDataSourceToggled is a free log subscription operation binding the contract event 0xc59a5de02f9d69be770ff0d61bbc894968433bb599f9fd9c2016e149c509c5e5.
//
// Solidity: event ERDataSourceToggled(bool isPORBasedERData)
func (_StaderOracle *StaderOracleFilterer) WatchERDataSourceToggled(opts *bind.WatchOpts, sink chan<- *StaderOracleERDataSourceToggled) (event.Subscription, error) {

	logs, sub, err := _StaderOracle.contract.WatchLogs(opts, "ERDataSourceToggled")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StaderOracleERDataSourceToggled)
				if err := _StaderOracle.contract.UnpackLog(event, "ERDataSourceToggled", log); err != nil {
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

// ParseERDataSourceToggled is a log parse operation binding the contract event 0xc59a5de02f9d69be770ff0d61bbc894968433bb599f9fd9c2016e149c509c5e5.
//
// Solidity: event ERDataSourceToggled(bool isPORBasedERData)
func (_StaderOracle *StaderOracleFilterer) ParseERDataSourceToggled(log types.Log) (*StaderOracleERDataSourceToggled, error) {
	event := new(StaderOracleERDataSourceToggled)
	if err := _StaderOracle.contract.UnpackLog(event, "ERDataSourceToggled", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StaderOracleERInspectionModeActivatedIterator is returned from FilterERInspectionModeActivated and is used to iterate over the raw logs and unpacked data for ERInspectionModeActivated events raised by the StaderOracle contract.
type StaderOracleERInspectionModeActivatedIterator struct {
	Event *StaderOracleERInspectionModeActivated // Event containing the contract specifics and raw log

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
func (it *StaderOracleERInspectionModeActivatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StaderOracleERInspectionModeActivated)
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
		it.Event = new(StaderOracleERInspectionModeActivated)
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
func (it *StaderOracleERInspectionModeActivatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StaderOracleERInspectionModeActivatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StaderOracleERInspectionModeActivated represents a ERInspectionModeActivated event raised by the StaderOracle contract.
type StaderOracleERInspectionModeActivated struct {
	ErInspectionMode bool
	Time             *big.Int
	Raw              types.Log // Blockchain specific contextual infos
}

// FilterERInspectionModeActivated is a free log retrieval operation binding the contract event 0x9dea6ddefbcfcf9c4f6c1c086e462c2ab380c0be199d0289bf23b09f20814415.
//
// Solidity: event ERInspectionModeActivated(bool erInspectionMode, uint256 time)
func (_StaderOracle *StaderOracleFilterer) FilterERInspectionModeActivated(opts *bind.FilterOpts) (*StaderOracleERInspectionModeActivatedIterator, error) {

	logs, sub, err := _StaderOracle.contract.FilterLogs(opts, "ERInspectionModeActivated")
	if err != nil {
		return nil, err
	}
	return &StaderOracleERInspectionModeActivatedIterator{contract: _StaderOracle.contract, event: "ERInspectionModeActivated", logs: logs, sub: sub}, nil
}

// WatchERInspectionModeActivated is a free log subscription operation binding the contract event 0x9dea6ddefbcfcf9c4f6c1c086e462c2ab380c0be199d0289bf23b09f20814415.
//
// Solidity: event ERInspectionModeActivated(bool erInspectionMode, uint256 time)
func (_StaderOracle *StaderOracleFilterer) WatchERInspectionModeActivated(opts *bind.WatchOpts, sink chan<- *StaderOracleERInspectionModeActivated) (event.Subscription, error) {

	logs, sub, err := _StaderOracle.contract.WatchLogs(opts, "ERInspectionModeActivated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StaderOracleERInspectionModeActivated)
				if err := _StaderOracle.contract.UnpackLog(event, "ERInspectionModeActivated", log); err != nil {
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

// ParseERInspectionModeActivated is a log parse operation binding the contract event 0x9dea6ddefbcfcf9c4f6c1c086e462c2ab380c0be199d0289bf23b09f20814415.
//
// Solidity: event ERInspectionModeActivated(bool erInspectionMode, uint256 time)
func (_StaderOracle *StaderOracleFilterer) ParseERInspectionModeActivated(log types.Log) (*StaderOracleERInspectionModeActivated, error) {
	event := new(StaderOracleERInspectionModeActivated)
	if err := _StaderOracle.contract.UnpackLog(event, "ERInspectionModeActivated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StaderOracleExchangeRateSubmittedIterator is returned from FilterExchangeRateSubmitted and is used to iterate over the raw logs and unpacked data for ExchangeRateSubmitted events raised by the StaderOracle contract.
type StaderOracleExchangeRateSubmittedIterator struct {
	Event *StaderOracleExchangeRateSubmitted // Event containing the contract specifics and raw log

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
func (it *StaderOracleExchangeRateSubmittedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StaderOracleExchangeRateSubmitted)
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
		it.Event = new(StaderOracleExchangeRateSubmitted)
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
func (it *StaderOracleExchangeRateSubmittedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StaderOracleExchangeRateSubmittedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StaderOracleExchangeRateSubmitted represents a ExchangeRateSubmitted event raised by the StaderOracle contract.
type StaderOracleExchangeRateSubmitted struct {
	From       common.Address
	Block      *big.Int
	TotalEth   *big.Int
	EthxSupply *big.Int
	Time       *big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterExchangeRateSubmitted is a free log retrieval operation binding the contract event 0x73327a5c0fdb3104b4a0f993bc20ce59885ac5bfe5f23e4bfdd19c21fb79cb12.
//
// Solidity: event ExchangeRateSubmitted(address indexed from, uint256 block, uint256 totalEth, uint256 ethxSupply, uint256 time)
func (_StaderOracle *StaderOracleFilterer) FilterExchangeRateSubmitted(opts *bind.FilterOpts, from []common.Address) (*StaderOracleExchangeRateSubmittedIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}

	logs, sub, err := _StaderOracle.contract.FilterLogs(opts, "ExchangeRateSubmitted", fromRule)
	if err != nil {
		return nil, err
	}
	return &StaderOracleExchangeRateSubmittedIterator{contract: _StaderOracle.contract, event: "ExchangeRateSubmitted", logs: logs, sub: sub}, nil
}

// WatchExchangeRateSubmitted is a free log subscription operation binding the contract event 0x73327a5c0fdb3104b4a0f993bc20ce59885ac5bfe5f23e4bfdd19c21fb79cb12.
//
// Solidity: event ExchangeRateSubmitted(address indexed from, uint256 block, uint256 totalEth, uint256 ethxSupply, uint256 time)
func (_StaderOracle *StaderOracleFilterer) WatchExchangeRateSubmitted(opts *bind.WatchOpts, sink chan<- *StaderOracleExchangeRateSubmitted, from []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}

	logs, sub, err := _StaderOracle.contract.WatchLogs(opts, "ExchangeRateSubmitted", fromRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StaderOracleExchangeRateSubmitted)
				if err := _StaderOracle.contract.UnpackLog(event, "ExchangeRateSubmitted", log); err != nil {
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

// ParseExchangeRateSubmitted is a log parse operation binding the contract event 0x73327a5c0fdb3104b4a0f993bc20ce59885ac5bfe5f23e4bfdd19c21fb79cb12.
//
// Solidity: event ExchangeRateSubmitted(address indexed from, uint256 block, uint256 totalEth, uint256 ethxSupply, uint256 time)
func (_StaderOracle *StaderOracleFilterer) ParseExchangeRateSubmitted(log types.Log) (*StaderOracleExchangeRateSubmitted, error) {
	event := new(StaderOracleExchangeRateSubmitted)
	if err := _StaderOracle.contract.UnpackLog(event, "ExchangeRateSubmitted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StaderOracleExchangeRateUpdatedIterator is returned from FilterExchangeRateUpdated and is used to iterate over the raw logs and unpacked data for ExchangeRateUpdated events raised by the StaderOracle contract.
type StaderOracleExchangeRateUpdatedIterator struct {
	Event *StaderOracleExchangeRateUpdated // Event containing the contract specifics and raw log

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
func (it *StaderOracleExchangeRateUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StaderOracleExchangeRateUpdated)
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
		it.Event = new(StaderOracleExchangeRateUpdated)
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
func (it *StaderOracleExchangeRateUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StaderOracleExchangeRateUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StaderOracleExchangeRateUpdated represents a ExchangeRateUpdated event raised by the StaderOracle contract.
type StaderOracleExchangeRateUpdated struct {
	Block      *big.Int
	TotalEth   *big.Int
	EthxSupply *big.Int
	Time       *big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterExchangeRateUpdated is a free log retrieval operation binding the contract event 0xf525f19964f35afcb9a475680bb27abecbc5e62b4d6d4f66a81a5bd7e8a107e3.
//
// Solidity: event ExchangeRateUpdated(uint256 block, uint256 totalEth, uint256 ethxSupply, uint256 time)
func (_StaderOracle *StaderOracleFilterer) FilterExchangeRateUpdated(opts *bind.FilterOpts) (*StaderOracleExchangeRateUpdatedIterator, error) {

	logs, sub, err := _StaderOracle.contract.FilterLogs(opts, "ExchangeRateUpdated")
	if err != nil {
		return nil, err
	}
	return &StaderOracleExchangeRateUpdatedIterator{contract: _StaderOracle.contract, event: "ExchangeRateUpdated", logs: logs, sub: sub}, nil
}

// WatchExchangeRateUpdated is a free log subscription operation binding the contract event 0xf525f19964f35afcb9a475680bb27abecbc5e62b4d6d4f66a81a5bd7e8a107e3.
//
// Solidity: event ExchangeRateUpdated(uint256 block, uint256 totalEth, uint256 ethxSupply, uint256 time)
func (_StaderOracle *StaderOracleFilterer) WatchExchangeRateUpdated(opts *bind.WatchOpts, sink chan<- *StaderOracleExchangeRateUpdated) (event.Subscription, error) {

	logs, sub, err := _StaderOracle.contract.WatchLogs(opts, "ExchangeRateUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StaderOracleExchangeRateUpdated)
				if err := _StaderOracle.contract.UnpackLog(event, "ExchangeRateUpdated", log); err != nil {
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

// ParseExchangeRateUpdated is a log parse operation binding the contract event 0xf525f19964f35afcb9a475680bb27abecbc5e62b4d6d4f66a81a5bd7e8a107e3.
//
// Solidity: event ExchangeRateUpdated(uint256 block, uint256 totalEth, uint256 ethxSupply, uint256 time)
func (_StaderOracle *StaderOracleFilterer) ParseExchangeRateUpdated(log types.Log) (*StaderOracleExchangeRateUpdated, error) {
	event := new(StaderOracleExchangeRateUpdated)
	if err := _StaderOracle.contract.UnpackLog(event, "ExchangeRateUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StaderOracleInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the StaderOracle contract.
type StaderOracleInitializedIterator struct {
	Event *StaderOracleInitialized // Event containing the contract specifics and raw log

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
func (it *StaderOracleInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StaderOracleInitialized)
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
		it.Event = new(StaderOracleInitialized)
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
func (it *StaderOracleInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StaderOracleInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StaderOracleInitialized represents a Initialized event raised by the StaderOracle contract.
type StaderOracleInitialized struct {
	Version uint8
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_StaderOracle *StaderOracleFilterer) FilterInitialized(opts *bind.FilterOpts) (*StaderOracleInitializedIterator, error) {

	logs, sub, err := _StaderOracle.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &StaderOracleInitializedIterator{contract: _StaderOracle.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_StaderOracle *StaderOracleFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *StaderOracleInitialized) (event.Subscription, error) {

	logs, sub, err := _StaderOracle.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StaderOracleInitialized)
				if err := _StaderOracle.contract.UnpackLog(event, "Initialized", log); err != nil {
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
func (_StaderOracle *StaderOracleFilterer) ParseInitialized(log types.Log) (*StaderOracleInitialized, error) {
	event := new(StaderOracleInitialized)
	if err := _StaderOracle.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StaderOracleMissedAttestationPenaltySubmittedIterator is returned from FilterMissedAttestationPenaltySubmitted and is used to iterate over the raw logs and unpacked data for MissedAttestationPenaltySubmitted events raised by the StaderOracle contract.
type StaderOracleMissedAttestationPenaltySubmittedIterator struct {
	Event *StaderOracleMissedAttestationPenaltySubmitted // Event containing the contract specifics and raw log

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
func (it *StaderOracleMissedAttestationPenaltySubmittedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StaderOracleMissedAttestationPenaltySubmitted)
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
		it.Event = new(StaderOracleMissedAttestationPenaltySubmitted)
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
func (it *StaderOracleMissedAttestationPenaltySubmittedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StaderOracleMissedAttestationPenaltySubmittedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StaderOracleMissedAttestationPenaltySubmitted represents a MissedAttestationPenaltySubmitted event raised by the StaderOracle contract.
type StaderOracleMissedAttestationPenaltySubmitted struct {
	Node                 common.Address
	Index                *big.Int
	Block                *big.Int
	ReportingBlockNumber *big.Int
	Pubkeys              [][]byte
	Raw                  types.Log // Blockchain specific contextual infos
}

// FilterMissedAttestationPenaltySubmitted is a free log retrieval operation binding the contract event 0x51308cad1da8fe95d4be43112c17d5651d3e3713a675ec61c2214fa16d9f6beb.
//
// Solidity: event MissedAttestationPenaltySubmitted(address indexed node, uint256 index, uint256 block, uint256 reportingBlockNumber, bytes[] pubkeys)
func (_StaderOracle *StaderOracleFilterer) FilterMissedAttestationPenaltySubmitted(opts *bind.FilterOpts, node []common.Address) (*StaderOracleMissedAttestationPenaltySubmittedIterator, error) {

	var nodeRule []interface{}
	for _, nodeItem := range node {
		nodeRule = append(nodeRule, nodeItem)
	}

	logs, sub, err := _StaderOracle.contract.FilterLogs(opts, "MissedAttestationPenaltySubmitted", nodeRule)
	if err != nil {
		return nil, err
	}
	return &StaderOracleMissedAttestationPenaltySubmittedIterator{contract: _StaderOracle.contract, event: "MissedAttestationPenaltySubmitted", logs: logs, sub: sub}, nil
}

// WatchMissedAttestationPenaltySubmitted is a free log subscription operation binding the contract event 0x51308cad1da8fe95d4be43112c17d5651d3e3713a675ec61c2214fa16d9f6beb.
//
// Solidity: event MissedAttestationPenaltySubmitted(address indexed node, uint256 index, uint256 block, uint256 reportingBlockNumber, bytes[] pubkeys)
func (_StaderOracle *StaderOracleFilterer) WatchMissedAttestationPenaltySubmitted(opts *bind.WatchOpts, sink chan<- *StaderOracleMissedAttestationPenaltySubmitted, node []common.Address) (event.Subscription, error) {

	var nodeRule []interface{}
	for _, nodeItem := range node {
		nodeRule = append(nodeRule, nodeItem)
	}

	logs, sub, err := _StaderOracle.contract.WatchLogs(opts, "MissedAttestationPenaltySubmitted", nodeRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StaderOracleMissedAttestationPenaltySubmitted)
				if err := _StaderOracle.contract.UnpackLog(event, "MissedAttestationPenaltySubmitted", log); err != nil {
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

// ParseMissedAttestationPenaltySubmitted is a log parse operation binding the contract event 0x51308cad1da8fe95d4be43112c17d5651d3e3713a675ec61c2214fa16d9f6beb.
//
// Solidity: event MissedAttestationPenaltySubmitted(address indexed node, uint256 index, uint256 block, uint256 reportingBlockNumber, bytes[] pubkeys)
func (_StaderOracle *StaderOracleFilterer) ParseMissedAttestationPenaltySubmitted(log types.Log) (*StaderOracleMissedAttestationPenaltySubmitted, error) {
	event := new(StaderOracleMissedAttestationPenaltySubmitted)
	if err := _StaderOracle.contract.UnpackLog(event, "MissedAttestationPenaltySubmitted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StaderOracleMissedAttestationPenaltyUpdatedIterator is returned from FilterMissedAttestationPenaltyUpdated and is used to iterate over the raw logs and unpacked data for MissedAttestationPenaltyUpdated events raised by the StaderOracle contract.
type StaderOracleMissedAttestationPenaltyUpdatedIterator struct {
	Event *StaderOracleMissedAttestationPenaltyUpdated // Event containing the contract specifics and raw log

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
func (it *StaderOracleMissedAttestationPenaltyUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StaderOracleMissedAttestationPenaltyUpdated)
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
		it.Event = new(StaderOracleMissedAttestationPenaltyUpdated)
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
func (it *StaderOracleMissedAttestationPenaltyUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StaderOracleMissedAttestationPenaltyUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StaderOracleMissedAttestationPenaltyUpdated represents a MissedAttestationPenaltyUpdated event raised by the StaderOracle contract.
type StaderOracleMissedAttestationPenaltyUpdated struct {
	Index   *big.Int
	Block   *big.Int
	Pubkeys [][]byte
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterMissedAttestationPenaltyUpdated is a free log retrieval operation binding the contract event 0x5454855cf2eeb89296b9e10ba0881425fa305f06ce9ccb1b0ce47bc2f107a191.
//
// Solidity: event MissedAttestationPenaltyUpdated(uint256 index, uint256 block, bytes[] pubkeys)
func (_StaderOracle *StaderOracleFilterer) FilterMissedAttestationPenaltyUpdated(opts *bind.FilterOpts) (*StaderOracleMissedAttestationPenaltyUpdatedIterator, error) {

	logs, sub, err := _StaderOracle.contract.FilterLogs(opts, "MissedAttestationPenaltyUpdated")
	if err != nil {
		return nil, err
	}
	return &StaderOracleMissedAttestationPenaltyUpdatedIterator{contract: _StaderOracle.contract, event: "MissedAttestationPenaltyUpdated", logs: logs, sub: sub}, nil
}

// WatchMissedAttestationPenaltyUpdated is a free log subscription operation binding the contract event 0x5454855cf2eeb89296b9e10ba0881425fa305f06ce9ccb1b0ce47bc2f107a191.
//
// Solidity: event MissedAttestationPenaltyUpdated(uint256 index, uint256 block, bytes[] pubkeys)
func (_StaderOracle *StaderOracleFilterer) WatchMissedAttestationPenaltyUpdated(opts *bind.WatchOpts, sink chan<- *StaderOracleMissedAttestationPenaltyUpdated) (event.Subscription, error) {

	logs, sub, err := _StaderOracle.contract.WatchLogs(opts, "MissedAttestationPenaltyUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StaderOracleMissedAttestationPenaltyUpdated)
				if err := _StaderOracle.contract.UnpackLog(event, "MissedAttestationPenaltyUpdated", log); err != nil {
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

// ParseMissedAttestationPenaltyUpdated is a log parse operation binding the contract event 0x5454855cf2eeb89296b9e10ba0881425fa305f06ce9ccb1b0ce47bc2f107a191.
//
// Solidity: event MissedAttestationPenaltyUpdated(uint256 index, uint256 block, bytes[] pubkeys)
func (_StaderOracle *StaderOracleFilterer) ParseMissedAttestationPenaltyUpdated(log types.Log) (*StaderOracleMissedAttestationPenaltyUpdated, error) {
	event := new(StaderOracleMissedAttestationPenaltyUpdated)
	if err := _StaderOracle.contract.UnpackLog(event, "MissedAttestationPenaltyUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StaderOraclePausedIterator is returned from FilterPaused and is used to iterate over the raw logs and unpacked data for Paused events raised by the StaderOracle contract.
type StaderOraclePausedIterator struct {
	Event *StaderOraclePaused // Event containing the contract specifics and raw log

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
func (it *StaderOraclePausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StaderOraclePaused)
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
		it.Event = new(StaderOraclePaused)
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
func (it *StaderOraclePausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StaderOraclePausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StaderOraclePaused represents a Paused event raised by the StaderOracle contract.
type StaderOraclePaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterPaused is a free log retrieval operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_StaderOracle *StaderOracleFilterer) FilterPaused(opts *bind.FilterOpts) (*StaderOraclePausedIterator, error) {

	logs, sub, err := _StaderOracle.contract.FilterLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return &StaderOraclePausedIterator{contract: _StaderOracle.contract, event: "Paused", logs: logs, sub: sub}, nil
}

// WatchPaused is a free log subscription operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_StaderOracle *StaderOracleFilterer) WatchPaused(opts *bind.WatchOpts, sink chan<- *StaderOraclePaused) (event.Subscription, error) {

	logs, sub, err := _StaderOracle.contract.WatchLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StaderOraclePaused)
				if err := _StaderOracle.contract.UnpackLog(event, "Paused", log); err != nil {
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
func (_StaderOracle *StaderOracleFilterer) ParsePaused(log types.Log) (*StaderOraclePaused, error) {
	event := new(StaderOraclePaused)
	if err := _StaderOracle.contract.UnpackLog(event, "Paused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StaderOracleRoleAdminChangedIterator is returned from FilterRoleAdminChanged and is used to iterate over the raw logs and unpacked data for RoleAdminChanged events raised by the StaderOracle contract.
type StaderOracleRoleAdminChangedIterator struct {
	Event *StaderOracleRoleAdminChanged // Event containing the contract specifics and raw log

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
func (it *StaderOracleRoleAdminChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StaderOracleRoleAdminChanged)
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
		it.Event = new(StaderOracleRoleAdminChanged)
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
func (it *StaderOracleRoleAdminChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StaderOracleRoleAdminChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StaderOracleRoleAdminChanged represents a RoleAdminChanged event raised by the StaderOracle contract.
type StaderOracleRoleAdminChanged struct {
	Role              [32]byte
	PreviousAdminRole [32]byte
	NewAdminRole      [32]byte
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterRoleAdminChanged is a free log retrieval operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_StaderOracle *StaderOracleFilterer) FilterRoleAdminChanged(opts *bind.FilterOpts, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (*StaderOracleRoleAdminChangedIterator, error) {

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

	logs, sub, err := _StaderOracle.contract.FilterLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return &StaderOracleRoleAdminChangedIterator{contract: _StaderOracle.contract, event: "RoleAdminChanged", logs: logs, sub: sub}, nil
}

// WatchRoleAdminChanged is a free log subscription operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_StaderOracle *StaderOracleFilterer) WatchRoleAdminChanged(opts *bind.WatchOpts, sink chan<- *StaderOracleRoleAdminChanged, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (event.Subscription, error) {

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

	logs, sub, err := _StaderOracle.contract.WatchLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StaderOracleRoleAdminChanged)
				if err := _StaderOracle.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
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
func (_StaderOracle *StaderOracleFilterer) ParseRoleAdminChanged(log types.Log) (*StaderOracleRoleAdminChanged, error) {
	event := new(StaderOracleRoleAdminChanged)
	if err := _StaderOracle.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StaderOracleRoleGrantedIterator is returned from FilterRoleGranted and is used to iterate over the raw logs and unpacked data for RoleGranted events raised by the StaderOracle contract.
type StaderOracleRoleGrantedIterator struct {
	Event *StaderOracleRoleGranted // Event containing the contract specifics and raw log

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
func (it *StaderOracleRoleGrantedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StaderOracleRoleGranted)
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
		it.Event = new(StaderOracleRoleGranted)
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
func (it *StaderOracleRoleGrantedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StaderOracleRoleGrantedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StaderOracleRoleGranted represents a RoleGranted event raised by the StaderOracle contract.
type StaderOracleRoleGranted struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleGranted is a free log retrieval operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_StaderOracle *StaderOracleFilterer) FilterRoleGranted(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*StaderOracleRoleGrantedIterator, error) {

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

	logs, sub, err := _StaderOracle.contract.FilterLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &StaderOracleRoleGrantedIterator{contract: _StaderOracle.contract, event: "RoleGranted", logs: logs, sub: sub}, nil
}

// WatchRoleGranted is a free log subscription operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_StaderOracle *StaderOracleFilterer) WatchRoleGranted(opts *bind.WatchOpts, sink chan<- *StaderOracleRoleGranted, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _StaderOracle.contract.WatchLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StaderOracleRoleGranted)
				if err := _StaderOracle.contract.UnpackLog(event, "RoleGranted", log); err != nil {
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
func (_StaderOracle *StaderOracleFilterer) ParseRoleGranted(log types.Log) (*StaderOracleRoleGranted, error) {
	event := new(StaderOracleRoleGranted)
	if err := _StaderOracle.contract.UnpackLog(event, "RoleGranted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StaderOracleRoleRevokedIterator is returned from FilterRoleRevoked and is used to iterate over the raw logs and unpacked data for RoleRevoked events raised by the StaderOracle contract.
type StaderOracleRoleRevokedIterator struct {
	Event *StaderOracleRoleRevoked // Event containing the contract specifics and raw log

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
func (it *StaderOracleRoleRevokedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StaderOracleRoleRevoked)
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
		it.Event = new(StaderOracleRoleRevoked)
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
func (it *StaderOracleRoleRevokedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StaderOracleRoleRevokedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StaderOracleRoleRevoked represents a RoleRevoked event raised by the StaderOracle contract.
type StaderOracleRoleRevoked struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleRevoked is a free log retrieval operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_StaderOracle *StaderOracleFilterer) FilterRoleRevoked(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*StaderOracleRoleRevokedIterator, error) {

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

	logs, sub, err := _StaderOracle.contract.FilterLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &StaderOracleRoleRevokedIterator{contract: _StaderOracle.contract, event: "RoleRevoked", logs: logs, sub: sub}, nil
}

// WatchRoleRevoked is a free log subscription operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_StaderOracle *StaderOracleFilterer) WatchRoleRevoked(opts *bind.WatchOpts, sink chan<- *StaderOracleRoleRevoked, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _StaderOracle.contract.WatchLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StaderOracleRoleRevoked)
				if err := _StaderOracle.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
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
func (_StaderOracle *StaderOracleFilterer) ParseRoleRevoked(log types.Log) (*StaderOracleRoleRevoked, error) {
	event := new(StaderOracleRoleRevoked)
	if err := _StaderOracle.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StaderOracleSDPriceSubmittedIterator is returned from FilterSDPriceSubmitted and is used to iterate over the raw logs and unpacked data for SDPriceSubmitted events raised by the StaderOracle contract.
type StaderOracleSDPriceSubmittedIterator struct {
	Event *StaderOracleSDPriceSubmitted // Event containing the contract specifics and raw log

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
func (it *StaderOracleSDPriceSubmittedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StaderOracleSDPriceSubmitted)
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
		it.Event = new(StaderOracleSDPriceSubmitted)
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
func (it *StaderOracleSDPriceSubmittedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StaderOracleSDPriceSubmittedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StaderOracleSDPriceSubmitted represents a SDPriceSubmitted event raised by the StaderOracle contract.
type StaderOracleSDPriceSubmitted struct {
	Node          common.Address
	SdPriceInETH  *big.Int
	ReportedBlock *big.Int
	Block         *big.Int
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterSDPriceSubmitted is a free log retrieval operation binding the contract event 0x6c291a7ce9906b2982643002c104cb0ba9f2b9fecc8b38e9cc3cf5cfca65b4e8.
//
// Solidity: event SDPriceSubmitted(address indexed node, uint256 sdPriceInETH, uint256 reportedBlock, uint256 block)
func (_StaderOracle *StaderOracleFilterer) FilterSDPriceSubmitted(opts *bind.FilterOpts, node []common.Address) (*StaderOracleSDPriceSubmittedIterator, error) {

	var nodeRule []interface{}
	for _, nodeItem := range node {
		nodeRule = append(nodeRule, nodeItem)
	}

	logs, sub, err := _StaderOracle.contract.FilterLogs(opts, "SDPriceSubmitted", nodeRule)
	if err != nil {
		return nil, err
	}
	return &StaderOracleSDPriceSubmittedIterator{contract: _StaderOracle.contract, event: "SDPriceSubmitted", logs: logs, sub: sub}, nil
}

// WatchSDPriceSubmitted is a free log subscription operation binding the contract event 0x6c291a7ce9906b2982643002c104cb0ba9f2b9fecc8b38e9cc3cf5cfca65b4e8.
//
// Solidity: event SDPriceSubmitted(address indexed node, uint256 sdPriceInETH, uint256 reportedBlock, uint256 block)
func (_StaderOracle *StaderOracleFilterer) WatchSDPriceSubmitted(opts *bind.WatchOpts, sink chan<- *StaderOracleSDPriceSubmitted, node []common.Address) (event.Subscription, error) {

	var nodeRule []interface{}
	for _, nodeItem := range node {
		nodeRule = append(nodeRule, nodeItem)
	}

	logs, sub, err := _StaderOracle.contract.WatchLogs(opts, "SDPriceSubmitted", nodeRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StaderOracleSDPriceSubmitted)
				if err := _StaderOracle.contract.UnpackLog(event, "SDPriceSubmitted", log); err != nil {
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

// ParseSDPriceSubmitted is a log parse operation binding the contract event 0x6c291a7ce9906b2982643002c104cb0ba9f2b9fecc8b38e9cc3cf5cfca65b4e8.
//
// Solidity: event SDPriceSubmitted(address indexed node, uint256 sdPriceInETH, uint256 reportedBlock, uint256 block)
func (_StaderOracle *StaderOracleFilterer) ParseSDPriceSubmitted(log types.Log) (*StaderOracleSDPriceSubmitted, error) {
	event := new(StaderOracleSDPriceSubmitted)
	if err := _StaderOracle.contract.UnpackLog(event, "SDPriceSubmitted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StaderOracleSDPriceUpdatedIterator is returned from FilterSDPriceUpdated and is used to iterate over the raw logs and unpacked data for SDPriceUpdated events raised by the StaderOracle contract.
type StaderOracleSDPriceUpdatedIterator struct {
	Event *StaderOracleSDPriceUpdated // Event containing the contract specifics and raw log

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
func (it *StaderOracleSDPriceUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StaderOracleSDPriceUpdated)
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
		it.Event = new(StaderOracleSDPriceUpdated)
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
func (it *StaderOracleSDPriceUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StaderOracleSDPriceUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StaderOracleSDPriceUpdated represents a SDPriceUpdated event raised by the StaderOracle contract.
type StaderOracleSDPriceUpdated struct {
	SdPriceInETH  *big.Int
	ReportedBlock *big.Int
	Block         *big.Int
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterSDPriceUpdated is a free log retrieval operation binding the contract event 0xbc1a0795e699bbeaa982f6049ae9689f4d0e3e22d554adb7c46626be62f3b8bc.
//
// Solidity: event SDPriceUpdated(uint256 sdPriceInETH, uint256 reportedBlock, uint256 block)
func (_StaderOracle *StaderOracleFilterer) FilterSDPriceUpdated(opts *bind.FilterOpts) (*StaderOracleSDPriceUpdatedIterator, error) {

	logs, sub, err := _StaderOracle.contract.FilterLogs(opts, "SDPriceUpdated")
	if err != nil {
		return nil, err
	}
	return &StaderOracleSDPriceUpdatedIterator{contract: _StaderOracle.contract, event: "SDPriceUpdated", logs: logs, sub: sub}, nil
}

// WatchSDPriceUpdated is a free log subscription operation binding the contract event 0xbc1a0795e699bbeaa982f6049ae9689f4d0e3e22d554adb7c46626be62f3b8bc.
//
// Solidity: event SDPriceUpdated(uint256 sdPriceInETH, uint256 reportedBlock, uint256 block)
func (_StaderOracle *StaderOracleFilterer) WatchSDPriceUpdated(opts *bind.WatchOpts, sink chan<- *StaderOracleSDPriceUpdated) (event.Subscription, error) {

	logs, sub, err := _StaderOracle.contract.WatchLogs(opts, "SDPriceUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StaderOracleSDPriceUpdated)
				if err := _StaderOracle.contract.UnpackLog(event, "SDPriceUpdated", log); err != nil {
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

// ParseSDPriceUpdated is a log parse operation binding the contract event 0xbc1a0795e699bbeaa982f6049ae9689f4d0e3e22d554adb7c46626be62f3b8bc.
//
// Solidity: event SDPriceUpdated(uint256 sdPriceInETH, uint256 reportedBlock, uint256 block)
func (_StaderOracle *StaderOracleFilterer) ParseSDPriceUpdated(log types.Log) (*StaderOracleSDPriceUpdated, error) {
	event := new(StaderOracleSDPriceUpdated)
	if err := _StaderOracle.contract.UnpackLog(event, "SDPriceUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StaderOracleSafeModeDisabledIterator is returned from FilterSafeModeDisabled and is used to iterate over the raw logs and unpacked data for SafeModeDisabled events raised by the StaderOracle contract.
type StaderOracleSafeModeDisabledIterator struct {
	Event *StaderOracleSafeModeDisabled // Event containing the contract specifics and raw log

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
func (it *StaderOracleSafeModeDisabledIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StaderOracleSafeModeDisabled)
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
		it.Event = new(StaderOracleSafeModeDisabled)
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
func (it *StaderOracleSafeModeDisabledIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StaderOracleSafeModeDisabledIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StaderOracleSafeModeDisabled represents a SafeModeDisabled event raised by the StaderOracle contract.
type StaderOracleSafeModeDisabled struct {
	Raw types.Log // Blockchain specific contextual infos
}

// FilterSafeModeDisabled is a free log retrieval operation binding the contract event 0xf29e751b3c28b619a215d25fee98a7828471a8e554626186d3f8d122f1650292.
//
// Solidity: event SafeModeDisabled()
func (_StaderOracle *StaderOracleFilterer) FilterSafeModeDisabled(opts *bind.FilterOpts) (*StaderOracleSafeModeDisabledIterator, error) {

	logs, sub, err := _StaderOracle.contract.FilterLogs(opts, "SafeModeDisabled")
	if err != nil {
		return nil, err
	}
	return &StaderOracleSafeModeDisabledIterator{contract: _StaderOracle.contract, event: "SafeModeDisabled", logs: logs, sub: sub}, nil
}

// WatchSafeModeDisabled is a free log subscription operation binding the contract event 0xf29e751b3c28b619a215d25fee98a7828471a8e554626186d3f8d122f1650292.
//
// Solidity: event SafeModeDisabled()
func (_StaderOracle *StaderOracleFilterer) WatchSafeModeDisabled(opts *bind.WatchOpts, sink chan<- *StaderOracleSafeModeDisabled) (event.Subscription, error) {

	logs, sub, err := _StaderOracle.contract.WatchLogs(opts, "SafeModeDisabled")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StaderOracleSafeModeDisabled)
				if err := _StaderOracle.contract.UnpackLog(event, "SafeModeDisabled", log); err != nil {
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

// ParseSafeModeDisabled is a log parse operation binding the contract event 0xf29e751b3c28b619a215d25fee98a7828471a8e554626186d3f8d122f1650292.
//
// Solidity: event SafeModeDisabled()
func (_StaderOracle *StaderOracleFilterer) ParseSafeModeDisabled(log types.Log) (*StaderOracleSafeModeDisabled, error) {
	event := new(StaderOracleSafeModeDisabled)
	if err := _StaderOracle.contract.UnpackLog(event, "SafeModeDisabled", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StaderOracleSafeModeEnabledIterator is returned from FilterSafeModeEnabled and is used to iterate over the raw logs and unpacked data for SafeModeEnabled events raised by the StaderOracle contract.
type StaderOracleSafeModeEnabledIterator struct {
	Event *StaderOracleSafeModeEnabled // Event containing the contract specifics and raw log

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
func (it *StaderOracleSafeModeEnabledIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StaderOracleSafeModeEnabled)
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
		it.Event = new(StaderOracleSafeModeEnabled)
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
func (it *StaderOracleSafeModeEnabledIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StaderOracleSafeModeEnabledIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StaderOracleSafeModeEnabled represents a SafeModeEnabled event raised by the StaderOracle contract.
type StaderOracleSafeModeEnabled struct {
	Raw types.Log // Blockchain specific contextual infos
}

// FilterSafeModeEnabled is a free log retrieval operation binding the contract event 0x3328bda355014adfb66d5d9086ab2c3204d1af7f83a69a3276daeed721f6ce3c.
//
// Solidity: event SafeModeEnabled()
func (_StaderOracle *StaderOracleFilterer) FilterSafeModeEnabled(opts *bind.FilterOpts) (*StaderOracleSafeModeEnabledIterator, error) {

	logs, sub, err := _StaderOracle.contract.FilterLogs(opts, "SafeModeEnabled")
	if err != nil {
		return nil, err
	}
	return &StaderOracleSafeModeEnabledIterator{contract: _StaderOracle.contract, event: "SafeModeEnabled", logs: logs, sub: sub}, nil
}

// WatchSafeModeEnabled is a free log subscription operation binding the contract event 0x3328bda355014adfb66d5d9086ab2c3204d1af7f83a69a3276daeed721f6ce3c.
//
// Solidity: event SafeModeEnabled()
func (_StaderOracle *StaderOracleFilterer) WatchSafeModeEnabled(opts *bind.WatchOpts, sink chan<- *StaderOracleSafeModeEnabled) (event.Subscription, error) {

	logs, sub, err := _StaderOracle.contract.WatchLogs(opts, "SafeModeEnabled")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StaderOracleSafeModeEnabled)
				if err := _StaderOracle.contract.UnpackLog(event, "SafeModeEnabled", log); err != nil {
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

// ParseSafeModeEnabled is a log parse operation binding the contract event 0x3328bda355014adfb66d5d9086ab2c3204d1af7f83a69a3276daeed721f6ce3c.
//
// Solidity: event SafeModeEnabled()
func (_StaderOracle *StaderOracleFilterer) ParseSafeModeEnabled(log types.Log) (*StaderOracleSafeModeEnabled, error) {
	event := new(StaderOracleSafeModeEnabled)
	if err := _StaderOracle.contract.UnpackLog(event, "SafeModeEnabled", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StaderOracleSocializingRewardsMerkleRootSubmittedIterator is returned from FilterSocializingRewardsMerkleRootSubmitted and is used to iterate over the raw logs and unpacked data for SocializingRewardsMerkleRootSubmitted events raised by the StaderOracle contract.
type StaderOracleSocializingRewardsMerkleRootSubmittedIterator struct {
	Event *StaderOracleSocializingRewardsMerkleRootSubmitted // Event containing the contract specifics and raw log

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
func (it *StaderOracleSocializingRewardsMerkleRootSubmittedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StaderOracleSocializingRewardsMerkleRootSubmitted)
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
		it.Event = new(StaderOracleSocializingRewardsMerkleRootSubmitted)
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
func (it *StaderOracleSocializingRewardsMerkleRootSubmittedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StaderOracleSocializingRewardsMerkleRootSubmittedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StaderOracleSocializingRewardsMerkleRootSubmitted represents a SocializingRewardsMerkleRootSubmitted event raised by the StaderOracle contract.
type StaderOracleSocializingRewardsMerkleRootSubmitted struct {
	Node       common.Address
	Index      *big.Int
	MerkleRoot [32]byte
	PoolId     uint8
	Block      *big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterSocializingRewardsMerkleRootSubmitted is a free log retrieval operation binding the contract event 0x97f29f2f9a3ad2e8ebffd3fb4a6dbf5035b92b432c8344609b8368407dd23377.
//
// Solidity: event SocializingRewardsMerkleRootSubmitted(address indexed node, uint256 index, bytes32 merkleRoot, uint8 poolId, uint256 block)
func (_StaderOracle *StaderOracleFilterer) FilterSocializingRewardsMerkleRootSubmitted(opts *bind.FilterOpts, node []common.Address) (*StaderOracleSocializingRewardsMerkleRootSubmittedIterator, error) {

	var nodeRule []interface{}
	for _, nodeItem := range node {
		nodeRule = append(nodeRule, nodeItem)
	}

	logs, sub, err := _StaderOracle.contract.FilterLogs(opts, "SocializingRewardsMerkleRootSubmitted", nodeRule)
	if err != nil {
		return nil, err
	}
	return &StaderOracleSocializingRewardsMerkleRootSubmittedIterator{contract: _StaderOracle.contract, event: "SocializingRewardsMerkleRootSubmitted", logs: logs, sub: sub}, nil
}

// WatchSocializingRewardsMerkleRootSubmitted is a free log subscription operation binding the contract event 0x97f29f2f9a3ad2e8ebffd3fb4a6dbf5035b92b432c8344609b8368407dd23377.
//
// Solidity: event SocializingRewardsMerkleRootSubmitted(address indexed node, uint256 index, bytes32 merkleRoot, uint8 poolId, uint256 block)
func (_StaderOracle *StaderOracleFilterer) WatchSocializingRewardsMerkleRootSubmitted(opts *bind.WatchOpts, sink chan<- *StaderOracleSocializingRewardsMerkleRootSubmitted, node []common.Address) (event.Subscription, error) {

	var nodeRule []interface{}
	for _, nodeItem := range node {
		nodeRule = append(nodeRule, nodeItem)
	}

	logs, sub, err := _StaderOracle.contract.WatchLogs(opts, "SocializingRewardsMerkleRootSubmitted", nodeRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StaderOracleSocializingRewardsMerkleRootSubmitted)
				if err := _StaderOracle.contract.UnpackLog(event, "SocializingRewardsMerkleRootSubmitted", log); err != nil {
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

// ParseSocializingRewardsMerkleRootSubmitted is a log parse operation binding the contract event 0x97f29f2f9a3ad2e8ebffd3fb4a6dbf5035b92b432c8344609b8368407dd23377.
//
// Solidity: event SocializingRewardsMerkleRootSubmitted(address indexed node, uint256 index, bytes32 merkleRoot, uint8 poolId, uint256 block)
func (_StaderOracle *StaderOracleFilterer) ParseSocializingRewardsMerkleRootSubmitted(log types.Log) (*StaderOracleSocializingRewardsMerkleRootSubmitted, error) {
	event := new(StaderOracleSocializingRewardsMerkleRootSubmitted)
	if err := _StaderOracle.contract.UnpackLog(event, "SocializingRewardsMerkleRootSubmitted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StaderOracleSocializingRewardsMerkleRootUpdatedIterator is returned from FilterSocializingRewardsMerkleRootUpdated and is used to iterate over the raw logs and unpacked data for SocializingRewardsMerkleRootUpdated events raised by the StaderOracle contract.
type StaderOracleSocializingRewardsMerkleRootUpdatedIterator struct {
	Event *StaderOracleSocializingRewardsMerkleRootUpdated // Event containing the contract specifics and raw log

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
func (it *StaderOracleSocializingRewardsMerkleRootUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StaderOracleSocializingRewardsMerkleRootUpdated)
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
		it.Event = new(StaderOracleSocializingRewardsMerkleRootUpdated)
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
func (it *StaderOracleSocializingRewardsMerkleRootUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StaderOracleSocializingRewardsMerkleRootUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StaderOracleSocializingRewardsMerkleRootUpdated represents a SocializingRewardsMerkleRootUpdated event raised by the StaderOracle contract.
type StaderOracleSocializingRewardsMerkleRootUpdated struct {
	Index      *big.Int
	MerkleRoot [32]byte
	PoolId     uint8
	Block      *big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterSocializingRewardsMerkleRootUpdated is a free log retrieval operation binding the contract event 0x4394ee7a4ca89453c6900058c69bfaf14014d9fc4d510ead54ae47ac06d1f05e.
//
// Solidity: event SocializingRewardsMerkleRootUpdated(uint256 index, bytes32 merkleRoot, uint8 poolId, uint256 block)
func (_StaderOracle *StaderOracleFilterer) FilterSocializingRewardsMerkleRootUpdated(opts *bind.FilterOpts) (*StaderOracleSocializingRewardsMerkleRootUpdatedIterator, error) {

	logs, sub, err := _StaderOracle.contract.FilterLogs(opts, "SocializingRewardsMerkleRootUpdated")
	if err != nil {
		return nil, err
	}
	return &StaderOracleSocializingRewardsMerkleRootUpdatedIterator{contract: _StaderOracle.contract, event: "SocializingRewardsMerkleRootUpdated", logs: logs, sub: sub}, nil
}

// WatchSocializingRewardsMerkleRootUpdated is a free log subscription operation binding the contract event 0x4394ee7a4ca89453c6900058c69bfaf14014d9fc4d510ead54ae47ac06d1f05e.
//
// Solidity: event SocializingRewardsMerkleRootUpdated(uint256 index, bytes32 merkleRoot, uint8 poolId, uint256 block)
func (_StaderOracle *StaderOracleFilterer) WatchSocializingRewardsMerkleRootUpdated(opts *bind.WatchOpts, sink chan<- *StaderOracleSocializingRewardsMerkleRootUpdated) (event.Subscription, error) {

	logs, sub, err := _StaderOracle.contract.WatchLogs(opts, "SocializingRewardsMerkleRootUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StaderOracleSocializingRewardsMerkleRootUpdated)
				if err := _StaderOracle.contract.UnpackLog(event, "SocializingRewardsMerkleRootUpdated", log); err != nil {
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

// ParseSocializingRewardsMerkleRootUpdated is a log parse operation binding the contract event 0x4394ee7a4ca89453c6900058c69bfaf14014d9fc4d510ead54ae47ac06d1f05e.
//
// Solidity: event SocializingRewardsMerkleRootUpdated(uint256 index, bytes32 merkleRoot, uint8 poolId, uint256 block)
func (_StaderOracle *StaderOracleFilterer) ParseSocializingRewardsMerkleRootUpdated(log types.Log) (*StaderOracleSocializingRewardsMerkleRootUpdated, error) {
	event := new(StaderOracleSocializingRewardsMerkleRootUpdated)
	if err := _StaderOracle.contract.UnpackLog(event, "SocializingRewardsMerkleRootUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StaderOracleTrustedNodeAddedIterator is returned from FilterTrustedNodeAdded and is used to iterate over the raw logs and unpacked data for TrustedNodeAdded events raised by the StaderOracle contract.
type StaderOracleTrustedNodeAddedIterator struct {
	Event *StaderOracleTrustedNodeAdded // Event containing the contract specifics and raw log

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
func (it *StaderOracleTrustedNodeAddedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StaderOracleTrustedNodeAdded)
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
		it.Event = new(StaderOracleTrustedNodeAdded)
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
func (it *StaderOracleTrustedNodeAddedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StaderOracleTrustedNodeAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StaderOracleTrustedNodeAdded represents a TrustedNodeAdded event raised by the StaderOracle contract.
type StaderOracleTrustedNodeAdded struct {
	Node common.Address
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterTrustedNodeAdded is a free log retrieval operation binding the contract event 0xff4a290f0500d113133708d378eb9a151c32d91cb8f5778cfda6328d89cfc4b8.
//
// Solidity: event TrustedNodeAdded(address indexed node)
func (_StaderOracle *StaderOracleFilterer) FilterTrustedNodeAdded(opts *bind.FilterOpts, node []common.Address) (*StaderOracleTrustedNodeAddedIterator, error) {

	var nodeRule []interface{}
	for _, nodeItem := range node {
		nodeRule = append(nodeRule, nodeItem)
	}

	logs, sub, err := _StaderOracle.contract.FilterLogs(opts, "TrustedNodeAdded", nodeRule)
	if err != nil {
		return nil, err
	}
	return &StaderOracleTrustedNodeAddedIterator{contract: _StaderOracle.contract, event: "TrustedNodeAdded", logs: logs, sub: sub}, nil
}

// WatchTrustedNodeAdded is a free log subscription operation binding the contract event 0xff4a290f0500d113133708d378eb9a151c32d91cb8f5778cfda6328d89cfc4b8.
//
// Solidity: event TrustedNodeAdded(address indexed node)
func (_StaderOracle *StaderOracleFilterer) WatchTrustedNodeAdded(opts *bind.WatchOpts, sink chan<- *StaderOracleTrustedNodeAdded, node []common.Address) (event.Subscription, error) {

	var nodeRule []interface{}
	for _, nodeItem := range node {
		nodeRule = append(nodeRule, nodeItem)
	}

	logs, sub, err := _StaderOracle.contract.WatchLogs(opts, "TrustedNodeAdded", nodeRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StaderOracleTrustedNodeAdded)
				if err := _StaderOracle.contract.UnpackLog(event, "TrustedNodeAdded", log); err != nil {
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

// ParseTrustedNodeAdded is a log parse operation binding the contract event 0xff4a290f0500d113133708d378eb9a151c32d91cb8f5778cfda6328d89cfc4b8.
//
// Solidity: event TrustedNodeAdded(address indexed node)
func (_StaderOracle *StaderOracleFilterer) ParseTrustedNodeAdded(log types.Log) (*StaderOracleTrustedNodeAdded, error) {
	event := new(StaderOracleTrustedNodeAdded)
	if err := _StaderOracle.contract.UnpackLog(event, "TrustedNodeAdded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StaderOracleTrustedNodeChangeCoolingPeriodUpdatedIterator is returned from FilterTrustedNodeChangeCoolingPeriodUpdated and is used to iterate over the raw logs and unpacked data for TrustedNodeChangeCoolingPeriodUpdated events raised by the StaderOracle contract.
type StaderOracleTrustedNodeChangeCoolingPeriodUpdatedIterator struct {
	Event *StaderOracleTrustedNodeChangeCoolingPeriodUpdated // Event containing the contract specifics and raw log

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
func (it *StaderOracleTrustedNodeChangeCoolingPeriodUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StaderOracleTrustedNodeChangeCoolingPeriodUpdated)
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
		it.Event = new(StaderOracleTrustedNodeChangeCoolingPeriodUpdated)
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
func (it *StaderOracleTrustedNodeChangeCoolingPeriodUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StaderOracleTrustedNodeChangeCoolingPeriodUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StaderOracleTrustedNodeChangeCoolingPeriodUpdated represents a TrustedNodeChangeCoolingPeriodUpdated event raised by the StaderOracle contract.
type StaderOracleTrustedNodeChangeCoolingPeriodUpdated struct {
	TrustedNodeChangeCoolingPeriod *big.Int
	Raw                            types.Log // Blockchain specific contextual infos
}

// FilterTrustedNodeChangeCoolingPeriodUpdated is a free log retrieval operation binding the contract event 0x4ab6bf3c94e4c92b7b93e89e984ef66d28392f440a58d91d244b6c303e85f046.
//
// Solidity: event TrustedNodeChangeCoolingPeriodUpdated(uint256 trustedNodeChangeCoolingPeriod)
func (_StaderOracle *StaderOracleFilterer) FilterTrustedNodeChangeCoolingPeriodUpdated(opts *bind.FilterOpts) (*StaderOracleTrustedNodeChangeCoolingPeriodUpdatedIterator, error) {

	logs, sub, err := _StaderOracle.contract.FilterLogs(opts, "TrustedNodeChangeCoolingPeriodUpdated")
	if err != nil {
		return nil, err
	}
	return &StaderOracleTrustedNodeChangeCoolingPeriodUpdatedIterator{contract: _StaderOracle.contract, event: "TrustedNodeChangeCoolingPeriodUpdated", logs: logs, sub: sub}, nil
}

// WatchTrustedNodeChangeCoolingPeriodUpdated is a free log subscription operation binding the contract event 0x4ab6bf3c94e4c92b7b93e89e984ef66d28392f440a58d91d244b6c303e85f046.
//
// Solidity: event TrustedNodeChangeCoolingPeriodUpdated(uint256 trustedNodeChangeCoolingPeriod)
func (_StaderOracle *StaderOracleFilterer) WatchTrustedNodeChangeCoolingPeriodUpdated(opts *bind.WatchOpts, sink chan<- *StaderOracleTrustedNodeChangeCoolingPeriodUpdated) (event.Subscription, error) {

	logs, sub, err := _StaderOracle.contract.WatchLogs(opts, "TrustedNodeChangeCoolingPeriodUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StaderOracleTrustedNodeChangeCoolingPeriodUpdated)
				if err := _StaderOracle.contract.UnpackLog(event, "TrustedNodeChangeCoolingPeriodUpdated", log); err != nil {
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

// ParseTrustedNodeChangeCoolingPeriodUpdated is a log parse operation binding the contract event 0x4ab6bf3c94e4c92b7b93e89e984ef66d28392f440a58d91d244b6c303e85f046.
//
// Solidity: event TrustedNodeChangeCoolingPeriodUpdated(uint256 trustedNodeChangeCoolingPeriod)
func (_StaderOracle *StaderOracleFilterer) ParseTrustedNodeChangeCoolingPeriodUpdated(log types.Log) (*StaderOracleTrustedNodeChangeCoolingPeriodUpdated, error) {
	event := new(StaderOracleTrustedNodeChangeCoolingPeriodUpdated)
	if err := _StaderOracle.contract.UnpackLog(event, "TrustedNodeChangeCoolingPeriodUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StaderOracleTrustedNodeRemovedIterator is returned from FilterTrustedNodeRemoved and is used to iterate over the raw logs and unpacked data for TrustedNodeRemoved events raised by the StaderOracle contract.
type StaderOracleTrustedNodeRemovedIterator struct {
	Event *StaderOracleTrustedNodeRemoved // Event containing the contract specifics and raw log

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
func (it *StaderOracleTrustedNodeRemovedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StaderOracleTrustedNodeRemoved)
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
		it.Event = new(StaderOracleTrustedNodeRemoved)
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
func (it *StaderOracleTrustedNodeRemovedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StaderOracleTrustedNodeRemovedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StaderOracleTrustedNodeRemoved represents a TrustedNodeRemoved event raised by the StaderOracle contract.
type StaderOracleTrustedNodeRemoved struct {
	Node common.Address
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterTrustedNodeRemoved is a free log retrieval operation binding the contract event 0x6d1de2fb0c63996bae7ba6277c720c0560ba42874ea34c1ebe8e1423b9b47421.
//
// Solidity: event TrustedNodeRemoved(address indexed node)
func (_StaderOracle *StaderOracleFilterer) FilterTrustedNodeRemoved(opts *bind.FilterOpts, node []common.Address) (*StaderOracleTrustedNodeRemovedIterator, error) {

	var nodeRule []interface{}
	for _, nodeItem := range node {
		nodeRule = append(nodeRule, nodeItem)
	}

	logs, sub, err := _StaderOracle.contract.FilterLogs(opts, "TrustedNodeRemoved", nodeRule)
	if err != nil {
		return nil, err
	}
	return &StaderOracleTrustedNodeRemovedIterator{contract: _StaderOracle.contract, event: "TrustedNodeRemoved", logs: logs, sub: sub}, nil
}

// WatchTrustedNodeRemoved is a free log subscription operation binding the contract event 0x6d1de2fb0c63996bae7ba6277c720c0560ba42874ea34c1ebe8e1423b9b47421.
//
// Solidity: event TrustedNodeRemoved(address indexed node)
func (_StaderOracle *StaderOracleFilterer) WatchTrustedNodeRemoved(opts *bind.WatchOpts, sink chan<- *StaderOracleTrustedNodeRemoved, node []common.Address) (event.Subscription, error) {

	var nodeRule []interface{}
	for _, nodeItem := range node {
		nodeRule = append(nodeRule, nodeItem)
	}

	logs, sub, err := _StaderOracle.contract.WatchLogs(opts, "TrustedNodeRemoved", nodeRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StaderOracleTrustedNodeRemoved)
				if err := _StaderOracle.contract.UnpackLog(event, "TrustedNodeRemoved", log); err != nil {
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

// ParseTrustedNodeRemoved is a log parse operation binding the contract event 0x6d1de2fb0c63996bae7ba6277c720c0560ba42874ea34c1ebe8e1423b9b47421.
//
// Solidity: event TrustedNodeRemoved(address indexed node)
func (_StaderOracle *StaderOracleFilterer) ParseTrustedNodeRemoved(log types.Log) (*StaderOracleTrustedNodeRemoved, error) {
	event := new(StaderOracleTrustedNodeRemoved)
	if err := _StaderOracle.contract.UnpackLog(event, "TrustedNodeRemoved", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StaderOracleUnpausedIterator is returned from FilterUnpaused and is used to iterate over the raw logs and unpacked data for Unpaused events raised by the StaderOracle contract.
type StaderOracleUnpausedIterator struct {
	Event *StaderOracleUnpaused // Event containing the contract specifics and raw log

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
func (it *StaderOracleUnpausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StaderOracleUnpaused)
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
		it.Event = new(StaderOracleUnpaused)
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
func (it *StaderOracleUnpausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StaderOracleUnpausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StaderOracleUnpaused represents a Unpaused event raised by the StaderOracle contract.
type StaderOracleUnpaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterUnpaused is a free log retrieval operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_StaderOracle *StaderOracleFilterer) FilterUnpaused(opts *bind.FilterOpts) (*StaderOracleUnpausedIterator, error) {

	logs, sub, err := _StaderOracle.contract.FilterLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return &StaderOracleUnpausedIterator{contract: _StaderOracle.contract, event: "Unpaused", logs: logs, sub: sub}, nil
}

// WatchUnpaused is a free log subscription operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_StaderOracle *StaderOracleFilterer) WatchUnpaused(opts *bind.WatchOpts, sink chan<- *StaderOracleUnpaused) (event.Subscription, error) {

	logs, sub, err := _StaderOracle.contract.WatchLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StaderOracleUnpaused)
				if err := _StaderOracle.contract.UnpackLog(event, "Unpaused", log); err != nil {
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
func (_StaderOracle *StaderOracleFilterer) ParseUnpaused(log types.Log) (*StaderOracleUnpaused, error) {
	event := new(StaderOracleUnpaused)
	if err := _StaderOracle.contract.UnpackLog(event, "Unpaused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StaderOracleUpdateFrequencyUpdatedIterator is returned from FilterUpdateFrequencyUpdated and is used to iterate over the raw logs and unpacked data for UpdateFrequencyUpdated events raised by the StaderOracle contract.
type StaderOracleUpdateFrequencyUpdatedIterator struct {
	Event *StaderOracleUpdateFrequencyUpdated // Event containing the contract specifics and raw log

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
func (it *StaderOracleUpdateFrequencyUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StaderOracleUpdateFrequencyUpdated)
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
		it.Event = new(StaderOracleUpdateFrequencyUpdated)
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
func (it *StaderOracleUpdateFrequencyUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StaderOracleUpdateFrequencyUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StaderOracleUpdateFrequencyUpdated represents a UpdateFrequencyUpdated event raised by the StaderOracle contract.
type StaderOracleUpdateFrequencyUpdated struct {
	UpdateFrequency *big.Int
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterUpdateFrequencyUpdated is a free log retrieval operation binding the contract event 0x6231a3e049e2072a042ae727208e7650b487871f4080458371e84d6e7d391138.
//
// Solidity: event UpdateFrequencyUpdated(uint256 updateFrequency)
func (_StaderOracle *StaderOracleFilterer) FilterUpdateFrequencyUpdated(opts *bind.FilterOpts) (*StaderOracleUpdateFrequencyUpdatedIterator, error) {

	logs, sub, err := _StaderOracle.contract.FilterLogs(opts, "UpdateFrequencyUpdated")
	if err != nil {
		return nil, err
	}
	return &StaderOracleUpdateFrequencyUpdatedIterator{contract: _StaderOracle.contract, event: "UpdateFrequencyUpdated", logs: logs, sub: sub}, nil
}

// WatchUpdateFrequencyUpdated is a free log subscription operation binding the contract event 0x6231a3e049e2072a042ae727208e7650b487871f4080458371e84d6e7d391138.
//
// Solidity: event UpdateFrequencyUpdated(uint256 updateFrequency)
func (_StaderOracle *StaderOracleFilterer) WatchUpdateFrequencyUpdated(opts *bind.WatchOpts, sink chan<- *StaderOracleUpdateFrequencyUpdated) (event.Subscription, error) {

	logs, sub, err := _StaderOracle.contract.WatchLogs(opts, "UpdateFrequencyUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StaderOracleUpdateFrequencyUpdated)
				if err := _StaderOracle.contract.UnpackLog(event, "UpdateFrequencyUpdated", log); err != nil {
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

// ParseUpdateFrequencyUpdated is a log parse operation binding the contract event 0x6231a3e049e2072a042ae727208e7650b487871f4080458371e84d6e7d391138.
//
// Solidity: event UpdateFrequencyUpdated(uint256 updateFrequency)
func (_StaderOracle *StaderOracleFilterer) ParseUpdateFrequencyUpdated(log types.Log) (*StaderOracleUpdateFrequencyUpdated, error) {
	event := new(StaderOracleUpdateFrequencyUpdated)
	if err := _StaderOracle.contract.UnpackLog(event, "UpdateFrequencyUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StaderOracleUpdatedERChangeLimitIterator is returned from FilterUpdatedERChangeLimit and is used to iterate over the raw logs and unpacked data for UpdatedERChangeLimit events raised by the StaderOracle contract.
type StaderOracleUpdatedERChangeLimitIterator struct {
	Event *StaderOracleUpdatedERChangeLimit // Event containing the contract specifics and raw log

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
func (it *StaderOracleUpdatedERChangeLimitIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StaderOracleUpdatedERChangeLimit)
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
		it.Event = new(StaderOracleUpdatedERChangeLimit)
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
func (it *StaderOracleUpdatedERChangeLimitIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StaderOracleUpdatedERChangeLimitIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StaderOracleUpdatedERChangeLimit represents a UpdatedERChangeLimit event raised by the StaderOracle contract.
type StaderOracleUpdatedERChangeLimit struct {
	ErChangeLimit *big.Int
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterUpdatedERChangeLimit is a free log retrieval operation binding the contract event 0x94a97bfc9c7a83fe5f75c66931ca7d2d16372fdc244afc5db36044f3ca52a520.
//
// Solidity: event UpdatedERChangeLimit(uint256 erChangeLimit)
func (_StaderOracle *StaderOracleFilterer) FilterUpdatedERChangeLimit(opts *bind.FilterOpts) (*StaderOracleUpdatedERChangeLimitIterator, error) {

	logs, sub, err := _StaderOracle.contract.FilterLogs(opts, "UpdatedERChangeLimit")
	if err != nil {
		return nil, err
	}
	return &StaderOracleUpdatedERChangeLimitIterator{contract: _StaderOracle.contract, event: "UpdatedERChangeLimit", logs: logs, sub: sub}, nil
}

// WatchUpdatedERChangeLimit is a free log subscription operation binding the contract event 0x94a97bfc9c7a83fe5f75c66931ca7d2d16372fdc244afc5db36044f3ca52a520.
//
// Solidity: event UpdatedERChangeLimit(uint256 erChangeLimit)
func (_StaderOracle *StaderOracleFilterer) WatchUpdatedERChangeLimit(opts *bind.WatchOpts, sink chan<- *StaderOracleUpdatedERChangeLimit) (event.Subscription, error) {

	logs, sub, err := _StaderOracle.contract.WatchLogs(opts, "UpdatedERChangeLimit")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StaderOracleUpdatedERChangeLimit)
				if err := _StaderOracle.contract.UnpackLog(event, "UpdatedERChangeLimit", log); err != nil {
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

// ParseUpdatedERChangeLimit is a log parse operation binding the contract event 0x94a97bfc9c7a83fe5f75c66931ca7d2d16372fdc244afc5db36044f3ca52a520.
//
// Solidity: event UpdatedERChangeLimit(uint256 erChangeLimit)
func (_StaderOracle *StaderOracleFilterer) ParseUpdatedERChangeLimit(log types.Log) (*StaderOracleUpdatedERChangeLimit, error) {
	event := new(StaderOracleUpdatedERChangeLimit)
	if err := _StaderOracle.contract.UnpackLog(event, "UpdatedERChangeLimit", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StaderOracleUpdatedStaderConfigIterator is returned from FilterUpdatedStaderConfig and is used to iterate over the raw logs and unpacked data for UpdatedStaderConfig events raised by the StaderOracle contract.
type StaderOracleUpdatedStaderConfigIterator struct {
	Event *StaderOracleUpdatedStaderConfig // Event containing the contract specifics and raw log

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
func (it *StaderOracleUpdatedStaderConfigIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StaderOracleUpdatedStaderConfig)
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
		it.Event = new(StaderOracleUpdatedStaderConfig)
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
func (it *StaderOracleUpdatedStaderConfigIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StaderOracleUpdatedStaderConfigIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StaderOracleUpdatedStaderConfig represents a UpdatedStaderConfig event raised by the StaderOracle contract.
type StaderOracleUpdatedStaderConfig struct {
	StaderConfig common.Address
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterUpdatedStaderConfig is a free log retrieval operation binding the contract event 0xdb2219043d7b197cb235f1af0cf6d782d77dee3de19e3f4fb6d39aae633b4485.
//
// Solidity: event UpdatedStaderConfig(address staderConfig)
func (_StaderOracle *StaderOracleFilterer) FilterUpdatedStaderConfig(opts *bind.FilterOpts) (*StaderOracleUpdatedStaderConfigIterator, error) {

	logs, sub, err := _StaderOracle.contract.FilterLogs(opts, "UpdatedStaderConfig")
	if err != nil {
		return nil, err
	}
	return &StaderOracleUpdatedStaderConfigIterator{contract: _StaderOracle.contract, event: "UpdatedStaderConfig", logs: logs, sub: sub}, nil
}

// WatchUpdatedStaderConfig is a free log subscription operation binding the contract event 0xdb2219043d7b197cb235f1af0cf6d782d77dee3de19e3f4fb6d39aae633b4485.
//
// Solidity: event UpdatedStaderConfig(address staderConfig)
func (_StaderOracle *StaderOracleFilterer) WatchUpdatedStaderConfig(opts *bind.WatchOpts, sink chan<- *StaderOracleUpdatedStaderConfig) (event.Subscription, error) {

	logs, sub, err := _StaderOracle.contract.WatchLogs(opts, "UpdatedStaderConfig")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StaderOracleUpdatedStaderConfig)
				if err := _StaderOracle.contract.UnpackLog(event, "UpdatedStaderConfig", log); err != nil {
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
func (_StaderOracle *StaderOracleFilterer) ParseUpdatedStaderConfig(log types.Log) (*StaderOracleUpdatedStaderConfig, error) {
	event := new(StaderOracleUpdatedStaderConfig)
	if err := _StaderOracle.contract.UnpackLog(event, "UpdatedStaderConfig", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StaderOracleValidatorStatsSubmittedIterator is returned from FilterValidatorStatsSubmitted and is used to iterate over the raw logs and unpacked data for ValidatorStatsSubmitted events raised by the StaderOracle contract.
type StaderOracleValidatorStatsSubmittedIterator struct {
	Event *StaderOracleValidatorStatsSubmitted // Event containing the contract specifics and raw log

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
func (it *StaderOracleValidatorStatsSubmittedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StaderOracleValidatorStatsSubmitted)
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
		it.Event = new(StaderOracleValidatorStatsSubmitted)
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
func (it *StaderOracleValidatorStatsSubmittedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StaderOracleValidatorStatsSubmittedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StaderOracleValidatorStatsSubmitted represents a ValidatorStatsSubmitted event raised by the StaderOracle contract.
type StaderOracleValidatorStatsSubmitted struct {
	From                     common.Address
	Block                    *big.Int
	ActiveValidatorsBalance  *big.Int
	ExitedValidatorsBalance  *big.Int
	SlashedValidatorsBalance *big.Int
	ActiveValidatorsCount    *big.Int
	ExitedValidatorsCount    *big.Int
	SlashedValidatorsCount   *big.Int
	Time                     *big.Int
	Raw                      types.Log // Blockchain specific contextual infos
}

// FilterValidatorStatsSubmitted is a free log retrieval operation binding the contract event 0x72745b0271618e5d84d738ea416e3a3be6eb267e0f639198f63c0ef124c29ffc.
//
// Solidity: event ValidatorStatsSubmitted(address indexed from, uint256 block, uint256 activeValidatorsBalance, uint256 exitedValidatorsBalance, uint256 slashedValidatorsBalance, uint256 activeValidatorsCount, uint256 exitedValidatorsCount, uint256 slashedValidatorsCount, uint256 time)
func (_StaderOracle *StaderOracleFilterer) FilterValidatorStatsSubmitted(opts *bind.FilterOpts, from []common.Address) (*StaderOracleValidatorStatsSubmittedIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}

	logs, sub, err := _StaderOracle.contract.FilterLogs(opts, "ValidatorStatsSubmitted", fromRule)
	if err != nil {
		return nil, err
	}
	return &StaderOracleValidatorStatsSubmittedIterator{contract: _StaderOracle.contract, event: "ValidatorStatsSubmitted", logs: logs, sub: sub}, nil
}

// WatchValidatorStatsSubmitted is a free log subscription operation binding the contract event 0x72745b0271618e5d84d738ea416e3a3be6eb267e0f639198f63c0ef124c29ffc.
//
// Solidity: event ValidatorStatsSubmitted(address indexed from, uint256 block, uint256 activeValidatorsBalance, uint256 exitedValidatorsBalance, uint256 slashedValidatorsBalance, uint256 activeValidatorsCount, uint256 exitedValidatorsCount, uint256 slashedValidatorsCount, uint256 time)
func (_StaderOracle *StaderOracleFilterer) WatchValidatorStatsSubmitted(opts *bind.WatchOpts, sink chan<- *StaderOracleValidatorStatsSubmitted, from []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}

	logs, sub, err := _StaderOracle.contract.WatchLogs(opts, "ValidatorStatsSubmitted", fromRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StaderOracleValidatorStatsSubmitted)
				if err := _StaderOracle.contract.UnpackLog(event, "ValidatorStatsSubmitted", log); err != nil {
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

// ParseValidatorStatsSubmitted is a log parse operation binding the contract event 0x72745b0271618e5d84d738ea416e3a3be6eb267e0f639198f63c0ef124c29ffc.
//
// Solidity: event ValidatorStatsSubmitted(address indexed from, uint256 block, uint256 activeValidatorsBalance, uint256 exitedValidatorsBalance, uint256 slashedValidatorsBalance, uint256 activeValidatorsCount, uint256 exitedValidatorsCount, uint256 slashedValidatorsCount, uint256 time)
func (_StaderOracle *StaderOracleFilterer) ParseValidatorStatsSubmitted(log types.Log) (*StaderOracleValidatorStatsSubmitted, error) {
	event := new(StaderOracleValidatorStatsSubmitted)
	if err := _StaderOracle.contract.UnpackLog(event, "ValidatorStatsSubmitted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StaderOracleValidatorStatsUpdatedIterator is returned from FilterValidatorStatsUpdated and is used to iterate over the raw logs and unpacked data for ValidatorStatsUpdated events raised by the StaderOracle contract.
type StaderOracleValidatorStatsUpdatedIterator struct {
	Event *StaderOracleValidatorStatsUpdated // Event containing the contract specifics and raw log

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
func (it *StaderOracleValidatorStatsUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StaderOracleValidatorStatsUpdated)
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
		it.Event = new(StaderOracleValidatorStatsUpdated)
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
func (it *StaderOracleValidatorStatsUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StaderOracleValidatorStatsUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StaderOracleValidatorStatsUpdated represents a ValidatorStatsUpdated event raised by the StaderOracle contract.
type StaderOracleValidatorStatsUpdated struct {
	Block                    *big.Int
	ActiveValidatorsBalance  *big.Int
	ExitedValidatorsBalance  *big.Int
	SlashedValidatorsBalance *big.Int
	ActiveValidatorsCount    *big.Int
	ExitedValidatorsCount    *big.Int
	SlashedValidatorsCount   *big.Int
	Time                     *big.Int
	Raw                      types.Log // Blockchain specific contextual infos
}

// FilterValidatorStatsUpdated is a free log retrieval operation binding the contract event 0x6988248fd82a7ce842fbdce0c49889ad926bf1548bae4194de0006498d069c94.
//
// Solidity: event ValidatorStatsUpdated(uint256 block, uint256 activeValidatorsBalance, uint256 exitedValidatorsBalance, uint256 slashedValidatorsBalance, uint256 activeValidatorsCount, uint256 exitedValidatorsCount, uint256 slashedValidatorsCount, uint256 time)
func (_StaderOracle *StaderOracleFilterer) FilterValidatorStatsUpdated(opts *bind.FilterOpts) (*StaderOracleValidatorStatsUpdatedIterator, error) {

	logs, sub, err := _StaderOracle.contract.FilterLogs(opts, "ValidatorStatsUpdated")
	if err != nil {
		return nil, err
	}
	return &StaderOracleValidatorStatsUpdatedIterator{contract: _StaderOracle.contract, event: "ValidatorStatsUpdated", logs: logs, sub: sub}, nil
}

// WatchValidatorStatsUpdated is a free log subscription operation binding the contract event 0x6988248fd82a7ce842fbdce0c49889ad926bf1548bae4194de0006498d069c94.
//
// Solidity: event ValidatorStatsUpdated(uint256 block, uint256 activeValidatorsBalance, uint256 exitedValidatorsBalance, uint256 slashedValidatorsBalance, uint256 activeValidatorsCount, uint256 exitedValidatorsCount, uint256 slashedValidatorsCount, uint256 time)
func (_StaderOracle *StaderOracleFilterer) WatchValidatorStatsUpdated(opts *bind.WatchOpts, sink chan<- *StaderOracleValidatorStatsUpdated) (event.Subscription, error) {

	logs, sub, err := _StaderOracle.contract.WatchLogs(opts, "ValidatorStatsUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StaderOracleValidatorStatsUpdated)
				if err := _StaderOracle.contract.UnpackLog(event, "ValidatorStatsUpdated", log); err != nil {
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

// ParseValidatorStatsUpdated is a log parse operation binding the contract event 0x6988248fd82a7ce842fbdce0c49889ad926bf1548bae4194de0006498d069c94.
//
// Solidity: event ValidatorStatsUpdated(uint256 block, uint256 activeValidatorsBalance, uint256 exitedValidatorsBalance, uint256 slashedValidatorsBalance, uint256 activeValidatorsCount, uint256 exitedValidatorsCount, uint256 slashedValidatorsCount, uint256 time)
func (_StaderOracle *StaderOracleFilterer) ParseValidatorStatsUpdated(log types.Log) (*StaderOracleValidatorStatsUpdated, error) {
	event := new(StaderOracleValidatorStatsUpdated)
	if err := _StaderOracle.contract.UnpackLog(event, "ValidatorStatsUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StaderOracleValidatorVerificationDetailSubmittedIterator is returned from FilterValidatorVerificationDetailSubmitted and is used to iterate over the raw logs and unpacked data for ValidatorVerificationDetailSubmitted events raised by the StaderOracle contract.
type StaderOracleValidatorVerificationDetailSubmittedIterator struct {
	Event *StaderOracleValidatorVerificationDetailSubmitted // Event containing the contract specifics and raw log

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
func (it *StaderOracleValidatorVerificationDetailSubmittedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StaderOracleValidatorVerificationDetailSubmitted)
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
		it.Event = new(StaderOracleValidatorVerificationDetailSubmitted)
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
func (it *StaderOracleValidatorVerificationDetailSubmittedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StaderOracleValidatorVerificationDetailSubmittedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StaderOracleValidatorVerificationDetailSubmitted represents a ValidatorVerificationDetailSubmitted event raised by the StaderOracle contract.
type StaderOracleValidatorVerificationDetailSubmitted struct {
	From                          common.Address
	PoolId                        uint8
	Block                         *big.Int
	SortedReadyToDepositPubkeys   [][]byte
	SortedFrontRunPubkeys         [][]byte
	SortedInvalidSignaturePubkeys [][]byte
	Time                          *big.Int
	Raw                           types.Log // Blockchain specific contextual infos
}

// FilterValidatorVerificationDetailSubmitted is a free log retrieval operation binding the contract event 0x9827231af99e07dd2167998d4c6855a2f78e0eead80a6a490393b1c3ead048a1.
//
// Solidity: event ValidatorVerificationDetailSubmitted(address indexed from, uint8 poolId, uint256 block, bytes[] sortedReadyToDepositPubkeys, bytes[] sortedFrontRunPubkeys, bytes[] sortedInvalidSignaturePubkeys, uint256 time)
func (_StaderOracle *StaderOracleFilterer) FilterValidatorVerificationDetailSubmitted(opts *bind.FilterOpts, from []common.Address) (*StaderOracleValidatorVerificationDetailSubmittedIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}

	logs, sub, err := _StaderOracle.contract.FilterLogs(opts, "ValidatorVerificationDetailSubmitted", fromRule)
	if err != nil {
		return nil, err
	}
	return &StaderOracleValidatorVerificationDetailSubmittedIterator{contract: _StaderOracle.contract, event: "ValidatorVerificationDetailSubmitted", logs: logs, sub: sub}, nil
}

// WatchValidatorVerificationDetailSubmitted is a free log subscription operation binding the contract event 0x9827231af99e07dd2167998d4c6855a2f78e0eead80a6a490393b1c3ead048a1.
//
// Solidity: event ValidatorVerificationDetailSubmitted(address indexed from, uint8 poolId, uint256 block, bytes[] sortedReadyToDepositPubkeys, bytes[] sortedFrontRunPubkeys, bytes[] sortedInvalidSignaturePubkeys, uint256 time)
func (_StaderOracle *StaderOracleFilterer) WatchValidatorVerificationDetailSubmitted(opts *bind.WatchOpts, sink chan<- *StaderOracleValidatorVerificationDetailSubmitted, from []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}

	logs, sub, err := _StaderOracle.contract.WatchLogs(opts, "ValidatorVerificationDetailSubmitted", fromRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StaderOracleValidatorVerificationDetailSubmitted)
				if err := _StaderOracle.contract.UnpackLog(event, "ValidatorVerificationDetailSubmitted", log); err != nil {
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

// ParseValidatorVerificationDetailSubmitted is a log parse operation binding the contract event 0x9827231af99e07dd2167998d4c6855a2f78e0eead80a6a490393b1c3ead048a1.
//
// Solidity: event ValidatorVerificationDetailSubmitted(address indexed from, uint8 poolId, uint256 block, bytes[] sortedReadyToDepositPubkeys, bytes[] sortedFrontRunPubkeys, bytes[] sortedInvalidSignaturePubkeys, uint256 time)
func (_StaderOracle *StaderOracleFilterer) ParseValidatorVerificationDetailSubmitted(log types.Log) (*StaderOracleValidatorVerificationDetailSubmitted, error) {
	event := new(StaderOracleValidatorVerificationDetailSubmitted)
	if err := _StaderOracle.contract.UnpackLog(event, "ValidatorVerificationDetailSubmitted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StaderOracleValidatorVerificationDetailUpdatedIterator is returned from FilterValidatorVerificationDetailUpdated and is used to iterate over the raw logs and unpacked data for ValidatorVerificationDetailUpdated events raised by the StaderOracle contract.
type StaderOracleValidatorVerificationDetailUpdatedIterator struct {
	Event *StaderOracleValidatorVerificationDetailUpdated // Event containing the contract specifics and raw log

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
func (it *StaderOracleValidatorVerificationDetailUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StaderOracleValidatorVerificationDetailUpdated)
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
		it.Event = new(StaderOracleValidatorVerificationDetailUpdated)
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
func (it *StaderOracleValidatorVerificationDetailUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StaderOracleValidatorVerificationDetailUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StaderOracleValidatorVerificationDetailUpdated represents a ValidatorVerificationDetailUpdated event raised by the StaderOracle contract.
type StaderOracleValidatorVerificationDetailUpdated struct {
	PoolId                        uint8
	Block                         *big.Int
	SortedReadyToDepositPubkeys   [][]byte
	SortedFrontRunPubkeys         [][]byte
	SortedInvalidSignaturePubkeys [][]byte
	Time                          *big.Int
	Raw                           types.Log // Blockchain specific contextual infos
}

// FilterValidatorVerificationDetailUpdated is a free log retrieval operation binding the contract event 0x64a4ab6f7a6ca6fe9c5750207ea4a9d3e2d5f5ba81f707146b7b20624f61cf11.
//
// Solidity: event ValidatorVerificationDetailUpdated(uint8 poolId, uint256 block, bytes[] sortedReadyToDepositPubkeys, bytes[] sortedFrontRunPubkeys, bytes[] sortedInvalidSignaturePubkeys, uint256 time)
func (_StaderOracle *StaderOracleFilterer) FilterValidatorVerificationDetailUpdated(opts *bind.FilterOpts) (*StaderOracleValidatorVerificationDetailUpdatedIterator, error) {

	logs, sub, err := _StaderOracle.contract.FilterLogs(opts, "ValidatorVerificationDetailUpdated")
	if err != nil {
		return nil, err
	}
	return &StaderOracleValidatorVerificationDetailUpdatedIterator{contract: _StaderOracle.contract, event: "ValidatorVerificationDetailUpdated", logs: logs, sub: sub}, nil
}

// WatchValidatorVerificationDetailUpdated is a free log subscription operation binding the contract event 0x64a4ab6f7a6ca6fe9c5750207ea4a9d3e2d5f5ba81f707146b7b20624f61cf11.
//
// Solidity: event ValidatorVerificationDetailUpdated(uint8 poolId, uint256 block, bytes[] sortedReadyToDepositPubkeys, bytes[] sortedFrontRunPubkeys, bytes[] sortedInvalidSignaturePubkeys, uint256 time)
func (_StaderOracle *StaderOracleFilterer) WatchValidatorVerificationDetailUpdated(opts *bind.WatchOpts, sink chan<- *StaderOracleValidatorVerificationDetailUpdated) (event.Subscription, error) {

	logs, sub, err := _StaderOracle.contract.WatchLogs(opts, "ValidatorVerificationDetailUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StaderOracleValidatorVerificationDetailUpdated)
				if err := _StaderOracle.contract.UnpackLog(event, "ValidatorVerificationDetailUpdated", log); err != nil {
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

// ParseValidatorVerificationDetailUpdated is a log parse operation binding the contract event 0x64a4ab6f7a6ca6fe9c5750207ea4a9d3e2d5f5ba81f707146b7b20624f61cf11.
//
// Solidity: event ValidatorVerificationDetailUpdated(uint8 poolId, uint256 block, bytes[] sortedReadyToDepositPubkeys, bytes[] sortedFrontRunPubkeys, bytes[] sortedInvalidSignaturePubkeys, uint256 time)
func (_StaderOracle *StaderOracleFilterer) ParseValidatorVerificationDetailUpdated(log types.Log) (*StaderOracleValidatorVerificationDetailUpdated, error) {
	event := new(StaderOracleValidatorVerificationDetailUpdated)
	if err := _StaderOracle.contract.UnpackLog(event, "ValidatorVerificationDetailUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StaderOracleWithdrawnValidatorsSubmittedIterator is returned from FilterWithdrawnValidatorsSubmitted and is used to iterate over the raw logs and unpacked data for WithdrawnValidatorsSubmitted events raised by the StaderOracle contract.
type StaderOracleWithdrawnValidatorsSubmittedIterator struct {
	Event *StaderOracleWithdrawnValidatorsSubmitted // Event containing the contract specifics and raw log

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
func (it *StaderOracleWithdrawnValidatorsSubmittedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StaderOracleWithdrawnValidatorsSubmitted)
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
		it.Event = new(StaderOracleWithdrawnValidatorsSubmitted)
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
func (it *StaderOracleWithdrawnValidatorsSubmittedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StaderOracleWithdrawnValidatorsSubmittedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StaderOracleWithdrawnValidatorsSubmitted represents a WithdrawnValidatorsSubmitted event raised by the StaderOracle contract.
type StaderOracleWithdrawnValidatorsSubmitted struct {
	From    common.Address
	PoolId  uint8
	Block   *big.Int
	Pubkeys [][]byte
	Time    *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterWithdrawnValidatorsSubmitted is a free log retrieval operation binding the contract event 0x3b426b614a89830a3d92d8dead18e2ded3cba56101dbff12dddfc1c77b21fbe6.
//
// Solidity: event WithdrawnValidatorsSubmitted(address indexed from, uint8 poolId, uint256 block, bytes[] pubkeys, uint256 time)
func (_StaderOracle *StaderOracleFilterer) FilterWithdrawnValidatorsSubmitted(opts *bind.FilterOpts, from []common.Address) (*StaderOracleWithdrawnValidatorsSubmittedIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}

	logs, sub, err := _StaderOracle.contract.FilterLogs(opts, "WithdrawnValidatorsSubmitted", fromRule)
	if err != nil {
		return nil, err
	}
	return &StaderOracleWithdrawnValidatorsSubmittedIterator{contract: _StaderOracle.contract, event: "WithdrawnValidatorsSubmitted", logs: logs, sub: sub}, nil
}

// WatchWithdrawnValidatorsSubmitted is a free log subscription operation binding the contract event 0x3b426b614a89830a3d92d8dead18e2ded3cba56101dbff12dddfc1c77b21fbe6.
//
// Solidity: event WithdrawnValidatorsSubmitted(address indexed from, uint8 poolId, uint256 block, bytes[] pubkeys, uint256 time)
func (_StaderOracle *StaderOracleFilterer) WatchWithdrawnValidatorsSubmitted(opts *bind.WatchOpts, sink chan<- *StaderOracleWithdrawnValidatorsSubmitted, from []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}

	logs, sub, err := _StaderOracle.contract.WatchLogs(opts, "WithdrawnValidatorsSubmitted", fromRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StaderOracleWithdrawnValidatorsSubmitted)
				if err := _StaderOracle.contract.UnpackLog(event, "WithdrawnValidatorsSubmitted", log); err != nil {
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

// ParseWithdrawnValidatorsSubmitted is a log parse operation binding the contract event 0x3b426b614a89830a3d92d8dead18e2ded3cba56101dbff12dddfc1c77b21fbe6.
//
// Solidity: event WithdrawnValidatorsSubmitted(address indexed from, uint8 poolId, uint256 block, bytes[] pubkeys, uint256 time)
func (_StaderOracle *StaderOracleFilterer) ParseWithdrawnValidatorsSubmitted(log types.Log) (*StaderOracleWithdrawnValidatorsSubmitted, error) {
	event := new(StaderOracleWithdrawnValidatorsSubmitted)
	if err := _StaderOracle.contract.UnpackLog(event, "WithdrawnValidatorsSubmitted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StaderOracleWithdrawnValidatorsUpdatedIterator is returned from FilterWithdrawnValidatorsUpdated and is used to iterate over the raw logs and unpacked data for WithdrawnValidatorsUpdated events raised by the StaderOracle contract.
type StaderOracleWithdrawnValidatorsUpdatedIterator struct {
	Event *StaderOracleWithdrawnValidatorsUpdated // Event containing the contract specifics and raw log

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
func (it *StaderOracleWithdrawnValidatorsUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StaderOracleWithdrawnValidatorsUpdated)
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
		it.Event = new(StaderOracleWithdrawnValidatorsUpdated)
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
func (it *StaderOracleWithdrawnValidatorsUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StaderOracleWithdrawnValidatorsUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StaderOracleWithdrawnValidatorsUpdated represents a WithdrawnValidatorsUpdated event raised by the StaderOracle contract.
type StaderOracleWithdrawnValidatorsUpdated struct {
	PoolId  uint8
	Block   *big.Int
	Pubkeys [][]byte
	Time    *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterWithdrawnValidatorsUpdated is a free log retrieval operation binding the contract event 0x5209293842928a1567d714f34fed8d87769d89d29dfb20f48ea678b02337e84d.
//
// Solidity: event WithdrawnValidatorsUpdated(uint8 poolId, uint256 block, bytes[] pubkeys, uint256 time)
func (_StaderOracle *StaderOracleFilterer) FilterWithdrawnValidatorsUpdated(opts *bind.FilterOpts) (*StaderOracleWithdrawnValidatorsUpdatedIterator, error) {

	logs, sub, err := _StaderOracle.contract.FilterLogs(opts, "WithdrawnValidatorsUpdated")
	if err != nil {
		return nil, err
	}
	return &StaderOracleWithdrawnValidatorsUpdatedIterator{contract: _StaderOracle.contract, event: "WithdrawnValidatorsUpdated", logs: logs, sub: sub}, nil
}

// WatchWithdrawnValidatorsUpdated is a free log subscription operation binding the contract event 0x5209293842928a1567d714f34fed8d87769d89d29dfb20f48ea678b02337e84d.
//
// Solidity: event WithdrawnValidatorsUpdated(uint8 poolId, uint256 block, bytes[] pubkeys, uint256 time)
func (_StaderOracle *StaderOracleFilterer) WatchWithdrawnValidatorsUpdated(opts *bind.WatchOpts, sink chan<- *StaderOracleWithdrawnValidatorsUpdated) (event.Subscription, error) {

	logs, sub, err := _StaderOracle.contract.WatchLogs(opts, "WithdrawnValidatorsUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StaderOracleWithdrawnValidatorsUpdated)
				if err := _StaderOracle.contract.UnpackLog(event, "WithdrawnValidatorsUpdated", log); err != nil {
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

// ParseWithdrawnValidatorsUpdated is a log parse operation binding the contract event 0x5209293842928a1567d714f34fed8d87769d89d29dfb20f48ea678b02337e84d.
//
// Solidity: event WithdrawnValidatorsUpdated(uint8 poolId, uint256 block, bytes[] pubkeys, uint256 time)
func (_StaderOracle *StaderOracleFilterer) ParseWithdrawnValidatorsUpdated(log types.Log) (*StaderOracleWithdrawnValidatorsUpdated, error) {
	event := new(StaderOracleWithdrawnValidatorsUpdated)
	if err := _StaderOracle.contract.UnpackLog(event, "WithdrawnValidatorsUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
