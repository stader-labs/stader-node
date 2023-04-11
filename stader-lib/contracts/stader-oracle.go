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
	ReportingBlockNumber   *big.Int
	TotalETHBalance        *big.Int
	TotalStakingETHBalance *big.Int
	TotalETHXSupply        *big.Int
}

// MissedAttestationPenaltyData is an auto generated low-level Go binding around an user-defined struct.
type MissedAttestationPenaltyData struct {
	KeyCount             uint16
	ReportingBlockNumber *big.Int
	Index                *big.Int
	PageNumber           *big.Int
	Pubkeys              []byte
}

// RewardsData is an auto generated low-level Go binding around an user-defined struct.
type RewardsData struct {
	ReportingBlockNumber *big.Int
	Index                *big.Int
	MerkleRoot           [32]byte
	PoolId               uint8
	OperatorETHRewards   *big.Int
	UserETHRewards       *big.Int
	ProtocolETHRewards   *big.Int
	OperatorSDRewards    *big.Int
}

// SDPriceData is an auto generated low-level Go binding around an user-defined struct.
type SDPriceData struct {
	ReportingBlockNumber *big.Int
	SdPriceInETH         *big.Int
}

// ValidatorStats is an auto generated low-level Go binding around an user-defined struct.
type ValidatorStats struct {
	ReportingBlockNumber     *big.Int
	ActiveValidatorsBalance  *big.Int
	ExitedValidatorsBalance  *big.Int
	SlashedValidatorsBalance *big.Int
	ActiveValidatorsCount    uint32
	ExitedValidatorsCount    uint32
	SlashedValidatorsCount   uint32
}

// WithdrawnValidators is an auto generated low-level Go binding around an user-defined struct.
type WithdrawnValidators struct {
	ReportingBlockNumber *big.Int
	NodeRegistry         common.Address
	SortedPubkeys        [][]byte
}

// StaderOracleMetaData contains all meta data concerning the StaderOracle contract.
var StaderOracleMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"DuplicateSubmissionFromNode\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"FrequencyUnchanged\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidData\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidMerkleRootIndex\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidNetworkBalances\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidPubkeyLength\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidReportingBlock\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NodeAlreadyTrusted\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NodeNotTrusted\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotATrustedNode\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"PageNumberAlreadyReported\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"PubkeysNotSorted\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ReportingFutureBlockData\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ReportingPreviousCycleData\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"StaleData\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"UpdateFrequencyNotSet\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ZeroAddress\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ZeroFrequency\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"block\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"totalEth\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"stakingEth\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"ethxSupply\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"time\",\"type\":\"uint256\"}],\"name\":\"BalancesSubmitted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"block\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"totalEth\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"stakingEth\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"ethxSupply\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"time\",\"type\":\"uint256\"}],\"name\":\"BalancesUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"version\",\"type\":\"uint8\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"node\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"pageNumber\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"block\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"reportingBlockNumber\",\"type\":\"uint256\"}],\"name\":\"MissedAttestationPenaltySubmitted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"block\",\"type\":\"uint256\"}],\"name\":\"MissedAttestationPenaltyUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"previousAdminRole\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"newAdminRole\",\"type\":\"bytes32\"}],\"name\":\"RoleAdminChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleGranted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleRevoked\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"node\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"sdPriceInETH\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"reportedBlock\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"block\",\"type\":\"uint256\"}],\"name\":\"SDPriceSubmitted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"sdPriceInETH\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"reportedBlock\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"block\",\"type\":\"uint256\"}],\"name\":\"SDPriceUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"node\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"merkleRoot\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"block\",\"type\":\"uint256\"}],\"name\":\"SocializingRewardsMerkleRootSubmitted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"merkleRoot\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"block\",\"type\":\"uint256\"}],\"name\":\"SocializingRewardsMerkleRootUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"node\",\"type\":\"address\"}],\"name\":\"TrustedNodeAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"node\",\"type\":\"address\"}],\"name\":\"TrustedNodeRemoved\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"updateFrequency\",\"type\":\"uint256\"}],\"name\":\"UpdateFrequencyUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"safeMode\",\"type\":\"bool\"}],\"name\":\"UpdatedSafeMode\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"staderConfig\",\"type\":\"address\"}],\"name\":\"UpdatedStaderConfig\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"block\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"activeValidatorsBalance\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"exitedValidatorsBalance\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"slashedValidatorsBalance\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"activeValidatorsCount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"exitedValidatorsCount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"slashedValidatorsCount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"time\",\"type\":\"uint256\"}],\"name\":\"ValidatorStatsSubmitted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"block\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"activeValidatorsBalance\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"exitedValidatorsBalance\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"slashedValidatorsBalance\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"activeValidatorsCount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"exitedValidatorsCount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"slashedValidatorsCount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"time\",\"type\":\"uint256\"}],\"name\":\"ValidatorStatsUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"block\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"nodeRegistry\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes[]\",\"name\":\"pubkeys\",\"type\":\"bytes[]\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"time\",\"type\":\"uint256\"}],\"name\":\"WithdrawnValidatorsSubmitted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"block\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"nodeRegistry\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes[]\",\"name\":\"pubkeys\",\"type\":\"bytes[]\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"time\",\"type\":\"uint256\"}],\"name\":\"WithdrawnValidatorsUpdated\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"DEFAULT_ADMIN_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"ETHX_ER_UF\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"MERKLE_UF\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"MISSED_ATTESTATION_PENALTY_UF\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"SD_PRICE_UF\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"STADER_MANAGER\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"VALIDATOR_STATS_UF\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"WITHDRAWN_VALIDATORS_UF\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_nodeAddress\",\"type\":\"address\"}],\"name\":\"addTrustedNode\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"exchangeRate\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"reportingBlockNumber\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"totalETHBalance\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"totalStakingETHBalance\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"totalETHXSupply\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getCurrentRewardsIndex\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getERReportableBlock\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getExchangeRate\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"reportingBlockNumber\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"totalETHBalance\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"totalStakingETHBalance\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"totalETHXSupply\",\"type\":\"uint256\"}],\"internalType\":\"structExchangeRate\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getMerkleRootReportableBlock\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getMissedAttestationPenaltyReportableBlock\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"_pubkey\",\"type\":\"bytes\"}],\"name\":\"getPubkeyRoot\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"}],\"name\":\"getRoleAdmin\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getSDPriceInETH\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getSDPriceReportableBlock\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getValidatorStats\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"reportingBlockNumber\",\"type\":\"uint256\"},{\"internalType\":\"uint128\",\"name\":\"activeValidatorsBalance\",\"type\":\"uint128\"},{\"internalType\":\"uint128\",\"name\":\"exitedValidatorsBalance\",\"type\":\"uint128\"},{\"internalType\":\"uint128\",\"name\":\"slashedValidatorsBalance\",\"type\":\"uint128\"},{\"internalType\":\"uint32\",\"name\":\"activeValidatorsCount\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"exitedValidatorsCount\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"slashedValidatorsCount\",\"type\":\"uint32\"}],\"internalType\":\"structValidatorStats\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getValidatorStatsReportableBlock\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getWithdrawnValidatorReportableBlock\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"grantRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"hasRole\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_staderConfig\",\"type\":\"address\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"isTrustedNode\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"lastReportedSDPriceData\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"reportingBlockNumber\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"sdPriceInETH\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"latestMissedAttestationConsensusIndex\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"missedAttestationDataByTrustedNode\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"pageNumber\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"missedAttestationPenalty\",\"outputs\":[{\"internalType\":\"uint16\",\"name\":\"\",\"type\":\"uint16\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_nodeAddress\",\"type\":\"address\"}],\"name\":\"removeTrustedNode\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"renounceRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"reportingBlockNumberForWithdrawnValidators\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"revokeRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"rewardsData\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"reportingBlockNumber\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"merkleRoot\",\"type\":\"bytes32\"},{\"internalType\":\"uint8\",\"name\":\"poolId\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"operatorETHRewards\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"userETHRewards\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"protocolETHRewards\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"operatorSDRewards\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"safeMode\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_updateFrequency\",\"type\":\"uint256\"}],\"name\":\"setERUpdateFrequency\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_updateFrequency\",\"type\":\"uint256\"}],\"name\":\"setMerkleRootUpdateFrequency\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_updateFrequency\",\"type\":\"uint256\"}],\"name\":\"setMissedAttestationPenaltyUpdateFrequency\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_updateFrequency\",\"type\":\"uint256\"}],\"name\":\"setSDPriceUpdateFrequency\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bool\",\"name\":\"_safeMode\",\"type\":\"bool\"}],\"name\":\"setSafeMode\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_updateFrequency\",\"type\":\"uint256\"}],\"name\":\"setValidatorStatsUpdateFrequency\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_updateFrequency\",\"type\":\"uint256\"}],\"name\":\"setWithdrawnValidatorsUpdateFrequency\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"socializingRewardsMerkleRoot\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"staderConfig\",\"outputs\":[{\"internalType\":\"contractIStaderConfig\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"reportingBlockNumber\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"totalETHBalance\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"totalStakingETHBalance\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"totalETHXSupply\",\"type\":\"uint256\"}],\"internalType\":\"structExchangeRate\",\"name\":\"_exchangeRate\",\"type\":\"tuple\"}],\"name\":\"submitBalances\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint16\",\"name\":\"keyCount\",\"type\":\"uint16\"},{\"internalType\":\"uint256\",\"name\":\"reportingBlockNumber\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"pageNumber\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"pubkeys\",\"type\":\"bytes\"}],\"internalType\":\"structMissedAttestationPenaltyData\",\"name\":\"_mapd\",\"type\":\"tuple\"}],\"name\":\"submitMissedAttestationPenalties\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"reportingBlockNumber\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"sdPriceInETH\",\"type\":\"uint256\"}],\"internalType\":\"structSDPriceData\",\"name\":\"_sdPriceData\",\"type\":\"tuple\"}],\"name\":\"submitSDPrice\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"reportingBlockNumber\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"merkleRoot\",\"type\":\"bytes32\"},{\"internalType\":\"uint8\",\"name\":\"poolId\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"operatorETHRewards\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"userETHRewards\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"protocolETHRewards\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"operatorSDRewards\",\"type\":\"uint256\"}],\"internalType\":\"structRewardsData\",\"name\":\"_rewardsData\",\"type\":\"tuple\"}],\"name\":\"submitSocializingRewardsMerkleRoot\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"reportingBlockNumber\",\"type\":\"uint256\"},{\"internalType\":\"uint128\",\"name\":\"activeValidatorsBalance\",\"type\":\"uint128\"},{\"internalType\":\"uint128\",\"name\":\"exitedValidatorsBalance\",\"type\":\"uint128\"},{\"internalType\":\"uint128\",\"name\":\"slashedValidatorsBalance\",\"type\":\"uint128\"},{\"internalType\":\"uint32\",\"name\":\"activeValidatorsCount\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"exitedValidatorsCount\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"slashedValidatorsCount\",\"type\":\"uint32\"}],\"internalType\":\"structValidatorStats\",\"name\":\"_validatorStats\",\"type\":\"tuple\"}],\"name\":\"submitValidatorStats\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"reportingBlockNumber\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"nodeRegistry\",\"type\":\"address\"},{\"internalType\":\"bytes[]\",\"name\":\"sortedPubkeys\",\"type\":\"bytes[]\"}],\"internalType\":\"structWithdrawnValidators\",\"name\":\"_withdrawnValidators\",\"type\":\"tuple\"}],\"name\":\"submitWithdrawnValidators\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"trustedNodesCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"updateFrequencyMap\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_staderConfig\",\"type\":\"address\"}],\"name\":\"updateStaderConfig\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"validatorStats\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"reportingBlockNumber\",\"type\":\"uint256\"},{\"internalType\":\"uint128\",\"name\":\"activeValidatorsBalance\",\"type\":\"uint128\"},{\"internalType\":\"uint128\",\"name\":\"exitedValidatorsBalance\",\"type\":\"uint128\"},{\"internalType\":\"uint128\",\"name\":\"slashedValidatorsBalance\",\"type\":\"uint128\"},{\"internalType\":\"uint32\",\"name\":\"activeValidatorsCount\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"exitedValidatorsCount\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"slashedValidatorsCount\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// StaderOracleABI is the input ABI used to generate the binding from.
// Deprecated: Use StaderOracleMetaData.ABI instead.
var StaderOracleABI = StaderOracleMetaData.ABI

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

// MERKLEUF is a free data retrieval call binding the contract method 0x091f3b19.
//
// Solidity: function MERKLE_UF() view returns(bytes32)
func (_StaderOracle *StaderOracleCaller) MERKLEUF(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _StaderOracle.contract.Call(opts, &out, "MERKLE_UF")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// MERKLEUF is a free data retrieval call binding the contract method 0x091f3b19.
//
// Solidity: function MERKLE_UF() view returns(bytes32)
func (_StaderOracle *StaderOracleSession) MERKLEUF() ([32]byte, error) {
	return _StaderOracle.Contract.MERKLEUF(&_StaderOracle.CallOpts)
}

// MERKLEUF is a free data retrieval call binding the contract method 0x091f3b19.
//
// Solidity: function MERKLE_UF() view returns(bytes32)
func (_StaderOracle *StaderOracleCallerSession) MERKLEUF() ([32]byte, error) {
	return _StaderOracle.Contract.MERKLEUF(&_StaderOracle.CallOpts)
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

// STADERMANAGER is a free data retrieval call binding the contract method 0xb8fd4c76.
//
// Solidity: function STADER_MANAGER() view returns(bytes32)
func (_StaderOracle *StaderOracleCaller) STADERMANAGER(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _StaderOracle.contract.Call(opts, &out, "STADER_MANAGER")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// STADERMANAGER is a free data retrieval call binding the contract method 0xb8fd4c76.
//
// Solidity: function STADER_MANAGER() view returns(bytes32)
func (_StaderOracle *StaderOracleSession) STADERMANAGER() ([32]byte, error) {
	return _StaderOracle.Contract.STADERMANAGER(&_StaderOracle.CallOpts)
}

// STADERMANAGER is a free data retrieval call binding the contract method 0xb8fd4c76.
//
// Solidity: function STADER_MANAGER() view returns(bytes32)
func (_StaderOracle *StaderOracleCallerSession) STADERMANAGER() ([32]byte, error) {
	return _StaderOracle.Contract.STADERMANAGER(&_StaderOracle.CallOpts)
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

// ExchangeRate is a free data retrieval call binding the contract method 0x3ba0b9a9.
//
// Solidity: function exchangeRate() view returns(uint256 reportingBlockNumber, uint256 totalETHBalance, uint256 totalStakingETHBalance, uint256 totalETHXSupply)
func (_StaderOracle *StaderOracleCaller) ExchangeRate(opts *bind.CallOpts) (struct {
	ReportingBlockNumber   *big.Int
	TotalETHBalance        *big.Int
	TotalStakingETHBalance *big.Int
	TotalETHXSupply        *big.Int
}, error) {
	var out []interface{}
	err := _StaderOracle.contract.Call(opts, &out, "exchangeRate")

	outstruct := new(struct {
		ReportingBlockNumber   *big.Int
		TotalETHBalance        *big.Int
		TotalStakingETHBalance *big.Int
		TotalETHXSupply        *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.ReportingBlockNumber = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.TotalETHBalance = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.TotalStakingETHBalance = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	outstruct.TotalETHXSupply = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// ExchangeRate is a free data retrieval call binding the contract method 0x3ba0b9a9.
//
// Solidity: function exchangeRate() view returns(uint256 reportingBlockNumber, uint256 totalETHBalance, uint256 totalStakingETHBalance, uint256 totalETHXSupply)
func (_StaderOracle *StaderOracleSession) ExchangeRate() (struct {
	ReportingBlockNumber   *big.Int
	TotalETHBalance        *big.Int
	TotalStakingETHBalance *big.Int
	TotalETHXSupply        *big.Int
}, error) {
	return _StaderOracle.Contract.ExchangeRate(&_StaderOracle.CallOpts)
}

// ExchangeRate is a free data retrieval call binding the contract method 0x3ba0b9a9.
//
// Solidity: function exchangeRate() view returns(uint256 reportingBlockNumber, uint256 totalETHBalance, uint256 totalStakingETHBalance, uint256 totalETHXSupply)
func (_StaderOracle *StaderOracleCallerSession) ExchangeRate() (struct {
	ReportingBlockNumber   *big.Int
	TotalETHBalance        *big.Int
	TotalStakingETHBalance *big.Int
	TotalETHXSupply        *big.Int
}, error) {
	return _StaderOracle.Contract.ExchangeRate(&_StaderOracle.CallOpts)
}

// GetCurrentRewardsIndex is a free data retrieval call binding the contract method 0x189956a2.
//
// Solidity: function getCurrentRewardsIndex() view returns(uint256)
func (_StaderOracle *StaderOracleCaller) GetCurrentRewardsIndex(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _StaderOracle.contract.Call(opts, &out, "getCurrentRewardsIndex")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetCurrentRewardsIndex is a free data retrieval call binding the contract method 0x189956a2.
//
// Solidity: function getCurrentRewardsIndex() view returns(uint256)
func (_StaderOracle *StaderOracleSession) GetCurrentRewardsIndex() (*big.Int, error) {
	return _StaderOracle.Contract.GetCurrentRewardsIndex(&_StaderOracle.CallOpts)
}

// GetCurrentRewardsIndex is a free data retrieval call binding the contract method 0x189956a2.
//
// Solidity: function getCurrentRewardsIndex() view returns(uint256)
func (_StaderOracle *StaderOracleCallerSession) GetCurrentRewardsIndex() (*big.Int, error) {
	return _StaderOracle.Contract.GetCurrentRewardsIndex(&_StaderOracle.CallOpts)
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
// Solidity: function getExchangeRate() view returns((uint256,uint256,uint256,uint256))
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
// Solidity: function getExchangeRate() view returns((uint256,uint256,uint256,uint256))
func (_StaderOracle *StaderOracleSession) GetExchangeRate() (ExchangeRate, error) {
	return _StaderOracle.Contract.GetExchangeRate(&_StaderOracle.CallOpts)
}

// GetExchangeRate is a free data retrieval call binding the contract method 0xe6aa216c.
//
// Solidity: function getExchangeRate() view returns((uint256,uint256,uint256,uint256))
func (_StaderOracle *StaderOracleCallerSession) GetExchangeRate() (ExchangeRate, error) {
	return _StaderOracle.Contract.GetExchangeRate(&_StaderOracle.CallOpts)
}

// GetMerkleRootReportableBlock is a free data retrieval call binding the contract method 0x68db0aa7.
//
// Solidity: function getMerkleRootReportableBlock() view returns(uint256)
func (_StaderOracle *StaderOracleCaller) GetMerkleRootReportableBlock(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _StaderOracle.contract.Call(opts, &out, "getMerkleRootReportableBlock")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetMerkleRootReportableBlock is a free data retrieval call binding the contract method 0x68db0aa7.
//
// Solidity: function getMerkleRootReportableBlock() view returns(uint256)
func (_StaderOracle *StaderOracleSession) GetMerkleRootReportableBlock() (*big.Int, error) {
	return _StaderOracle.Contract.GetMerkleRootReportableBlock(&_StaderOracle.CallOpts)
}

// GetMerkleRootReportableBlock is a free data retrieval call binding the contract method 0x68db0aa7.
//
// Solidity: function getMerkleRootReportableBlock() view returns(uint256)
func (_StaderOracle *StaderOracleCallerSession) GetMerkleRootReportableBlock() (*big.Int, error) {
	return _StaderOracle.Contract.GetMerkleRootReportableBlock(&_StaderOracle.CallOpts)
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

// GetPubkeyRoot is a free data retrieval call binding the contract method 0xf0c418c8.
//
// Solidity: function getPubkeyRoot(bytes _pubkey) pure returns(bytes32)
func (_StaderOracle *StaderOracleCaller) GetPubkeyRoot(opts *bind.CallOpts, _pubkey []byte) ([32]byte, error) {
	var out []interface{}
	err := _StaderOracle.contract.Call(opts, &out, "getPubkeyRoot", _pubkey)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetPubkeyRoot is a free data retrieval call binding the contract method 0xf0c418c8.
//
// Solidity: function getPubkeyRoot(bytes _pubkey) pure returns(bytes32)
func (_StaderOracle *StaderOracleSession) GetPubkeyRoot(_pubkey []byte) ([32]byte, error) {
	return _StaderOracle.Contract.GetPubkeyRoot(&_StaderOracle.CallOpts, _pubkey)
}

// GetPubkeyRoot is a free data retrieval call binding the contract method 0xf0c418c8.
//
// Solidity: function getPubkeyRoot(bytes _pubkey) pure returns(bytes32)
func (_StaderOracle *StaderOracleCallerSession) GetPubkeyRoot(_pubkey []byte) ([32]byte, error) {
	return _StaderOracle.Contract.GetPubkeyRoot(&_StaderOracle.CallOpts, _pubkey)
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

// LatestMissedAttestationConsensusIndex is a free data retrieval call binding the contract method 0x808b9f7d.
//
// Solidity: function latestMissedAttestationConsensusIndex() view returns(uint256)
func (_StaderOracle *StaderOracleCaller) LatestMissedAttestationConsensusIndex(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _StaderOracle.contract.Call(opts, &out, "latestMissedAttestationConsensusIndex")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// LatestMissedAttestationConsensusIndex is a free data retrieval call binding the contract method 0x808b9f7d.
//
// Solidity: function latestMissedAttestationConsensusIndex() view returns(uint256)
func (_StaderOracle *StaderOracleSession) LatestMissedAttestationConsensusIndex() (*big.Int, error) {
	return _StaderOracle.Contract.LatestMissedAttestationConsensusIndex(&_StaderOracle.CallOpts)
}

// LatestMissedAttestationConsensusIndex is a free data retrieval call binding the contract method 0x808b9f7d.
//
// Solidity: function latestMissedAttestationConsensusIndex() view returns(uint256)
func (_StaderOracle *StaderOracleCallerSession) LatestMissedAttestationConsensusIndex() (*big.Int, error) {
	return _StaderOracle.Contract.LatestMissedAttestationConsensusIndex(&_StaderOracle.CallOpts)
}

// MissedAttestationDataByTrustedNode is a free data retrieval call binding the contract method 0x295e6566.
//
// Solidity: function missedAttestationDataByTrustedNode(address ) view returns(uint256 index, uint256 pageNumber)
func (_StaderOracle *StaderOracleCaller) MissedAttestationDataByTrustedNode(opts *bind.CallOpts, arg0 common.Address) (struct {
	Index      *big.Int
	PageNumber *big.Int
}, error) {
	var out []interface{}
	err := _StaderOracle.contract.Call(opts, &out, "missedAttestationDataByTrustedNode", arg0)

	outstruct := new(struct {
		Index      *big.Int
		PageNumber *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Index = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.PageNumber = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// MissedAttestationDataByTrustedNode is a free data retrieval call binding the contract method 0x295e6566.
//
// Solidity: function missedAttestationDataByTrustedNode(address ) view returns(uint256 index, uint256 pageNumber)
func (_StaderOracle *StaderOracleSession) MissedAttestationDataByTrustedNode(arg0 common.Address) (struct {
	Index      *big.Int
	PageNumber *big.Int
}, error) {
	return _StaderOracle.Contract.MissedAttestationDataByTrustedNode(&_StaderOracle.CallOpts, arg0)
}

// MissedAttestationDataByTrustedNode is a free data retrieval call binding the contract method 0x295e6566.
//
// Solidity: function missedAttestationDataByTrustedNode(address ) view returns(uint256 index, uint256 pageNumber)
func (_StaderOracle *StaderOracleCallerSession) MissedAttestationDataByTrustedNode(arg0 common.Address) (struct {
	Index      *big.Int
	PageNumber *big.Int
}, error) {
	return _StaderOracle.Contract.MissedAttestationDataByTrustedNode(&_StaderOracle.CallOpts, arg0)
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

// ReportingBlockNumberForWithdrawnValidators is a free data retrieval call binding the contract method 0x1445f7a0.
//
// Solidity: function reportingBlockNumberForWithdrawnValidators() view returns(uint256)
func (_StaderOracle *StaderOracleCaller) ReportingBlockNumberForWithdrawnValidators(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _StaderOracle.contract.Call(opts, &out, "reportingBlockNumberForWithdrawnValidators")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ReportingBlockNumberForWithdrawnValidators is a free data retrieval call binding the contract method 0x1445f7a0.
//
// Solidity: function reportingBlockNumberForWithdrawnValidators() view returns(uint256)
func (_StaderOracle *StaderOracleSession) ReportingBlockNumberForWithdrawnValidators() (*big.Int, error) {
	return _StaderOracle.Contract.ReportingBlockNumberForWithdrawnValidators(&_StaderOracle.CallOpts)
}

// ReportingBlockNumberForWithdrawnValidators is a free data retrieval call binding the contract method 0x1445f7a0.
//
// Solidity: function reportingBlockNumberForWithdrawnValidators() view returns(uint256)
func (_StaderOracle *StaderOracleCallerSession) ReportingBlockNumberForWithdrawnValidators() (*big.Int, error) {
	return _StaderOracle.Contract.ReportingBlockNumberForWithdrawnValidators(&_StaderOracle.CallOpts)
}

// RewardsData is a free data retrieval call binding the contract method 0x03fb4295.
//
// Solidity: function rewardsData() view returns(uint256 reportingBlockNumber, uint256 index, bytes32 merkleRoot, uint8 poolId, uint256 operatorETHRewards, uint256 userETHRewards, uint256 protocolETHRewards, uint256 operatorSDRewards)
func (_StaderOracle *StaderOracleCaller) RewardsData(opts *bind.CallOpts) (struct {
	ReportingBlockNumber *big.Int
	Index                *big.Int
	MerkleRoot           [32]byte
	PoolId               uint8
	OperatorETHRewards   *big.Int
	UserETHRewards       *big.Int
	ProtocolETHRewards   *big.Int
	OperatorSDRewards    *big.Int
}, error) {
	var out []interface{}
	err := _StaderOracle.contract.Call(opts, &out, "rewardsData")

	outstruct := new(struct {
		ReportingBlockNumber *big.Int
		Index                *big.Int
		MerkleRoot           [32]byte
		PoolId               uint8
		OperatorETHRewards   *big.Int
		UserETHRewards       *big.Int
		ProtocolETHRewards   *big.Int
		OperatorSDRewards    *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.ReportingBlockNumber = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.Index = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.MerkleRoot = *abi.ConvertType(out[2], new([32]byte)).(*[32]byte)
	outstruct.PoolId = *abi.ConvertType(out[3], new(uint8)).(*uint8)
	outstruct.OperatorETHRewards = *abi.ConvertType(out[4], new(*big.Int)).(**big.Int)
	outstruct.UserETHRewards = *abi.ConvertType(out[5], new(*big.Int)).(**big.Int)
	outstruct.ProtocolETHRewards = *abi.ConvertType(out[6], new(*big.Int)).(**big.Int)
	outstruct.OperatorSDRewards = *abi.ConvertType(out[7], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// RewardsData is a free data retrieval call binding the contract method 0x03fb4295.
//
// Solidity: function rewardsData() view returns(uint256 reportingBlockNumber, uint256 index, bytes32 merkleRoot, uint8 poolId, uint256 operatorETHRewards, uint256 userETHRewards, uint256 protocolETHRewards, uint256 operatorSDRewards)
func (_StaderOracle *StaderOracleSession) RewardsData() (struct {
	ReportingBlockNumber *big.Int
	Index                *big.Int
	MerkleRoot           [32]byte
	PoolId               uint8
	OperatorETHRewards   *big.Int
	UserETHRewards       *big.Int
	ProtocolETHRewards   *big.Int
	OperatorSDRewards    *big.Int
}, error) {
	return _StaderOracle.Contract.RewardsData(&_StaderOracle.CallOpts)
}

// RewardsData is a free data retrieval call binding the contract method 0x03fb4295.
//
// Solidity: function rewardsData() view returns(uint256 reportingBlockNumber, uint256 index, bytes32 merkleRoot, uint8 poolId, uint256 operatorETHRewards, uint256 userETHRewards, uint256 protocolETHRewards, uint256 operatorSDRewards)
func (_StaderOracle *StaderOracleCallerSession) RewardsData() (struct {
	ReportingBlockNumber *big.Int
	Index                *big.Int
	MerkleRoot           [32]byte
	PoolId               uint8
	OperatorETHRewards   *big.Int
	UserETHRewards       *big.Int
	ProtocolETHRewards   *big.Int
	OperatorSDRewards    *big.Int
}, error) {
	return _StaderOracle.Contract.RewardsData(&_StaderOracle.CallOpts)
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

// SocializingRewardsMerkleRoot is a free data retrieval call binding the contract method 0xe2485740.
//
// Solidity: function socializingRewardsMerkleRoot(uint256 ) view returns(bytes32)
func (_StaderOracle *StaderOracleCaller) SocializingRewardsMerkleRoot(opts *bind.CallOpts, arg0 *big.Int) ([32]byte, error) {
	var out []interface{}
	err := _StaderOracle.contract.Call(opts, &out, "socializingRewardsMerkleRoot", arg0)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// SocializingRewardsMerkleRoot is a free data retrieval call binding the contract method 0xe2485740.
//
// Solidity: function socializingRewardsMerkleRoot(uint256 ) view returns(bytes32)
func (_StaderOracle *StaderOracleSession) SocializingRewardsMerkleRoot(arg0 *big.Int) ([32]byte, error) {
	return _StaderOracle.Contract.SocializingRewardsMerkleRoot(&_StaderOracle.CallOpts, arg0)
}

// SocializingRewardsMerkleRoot is a free data retrieval call binding the contract method 0xe2485740.
//
// Solidity: function socializingRewardsMerkleRoot(uint256 ) view returns(bytes32)
func (_StaderOracle *StaderOracleCallerSession) SocializingRewardsMerkleRoot(arg0 *big.Int) ([32]byte, error) {
	return _StaderOracle.Contract.SocializingRewardsMerkleRoot(&_StaderOracle.CallOpts, arg0)
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
// Solidity: function validatorStats() view returns(uint256 reportingBlockNumber, uint128 activeValidatorsBalance, uint128 exitedValidatorsBalance, uint128 slashedValidatorsBalance, uint32 activeValidatorsCount, uint32 exitedValidatorsCount, uint32 slashedValidatorsCount)
func (_StaderOracle *StaderOracleCaller) ValidatorStats(opts *bind.CallOpts) (struct {
	ReportingBlockNumber     *big.Int
	ActiveValidatorsBalance  *big.Int
	ExitedValidatorsBalance  *big.Int
	SlashedValidatorsBalance *big.Int
	ActiveValidatorsCount    uint32
	ExitedValidatorsCount    uint32
	SlashedValidatorsCount   uint32
}, error) {
	var out []interface{}
	err := _StaderOracle.contract.Call(opts, &out, "validatorStats")

	outstruct := new(struct {
		ReportingBlockNumber     *big.Int
		ActiveValidatorsBalance  *big.Int
		ExitedValidatorsBalance  *big.Int
		SlashedValidatorsBalance *big.Int
		ActiveValidatorsCount    uint32
		ExitedValidatorsCount    uint32
		SlashedValidatorsCount   uint32
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.ReportingBlockNumber = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.ActiveValidatorsBalance = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.ExitedValidatorsBalance = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	outstruct.SlashedValidatorsBalance = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)
	outstruct.ActiveValidatorsCount = *abi.ConvertType(out[4], new(uint32)).(*uint32)
	outstruct.ExitedValidatorsCount = *abi.ConvertType(out[5], new(uint32)).(*uint32)
	outstruct.SlashedValidatorsCount = *abi.ConvertType(out[6], new(uint32)).(*uint32)

	return *outstruct, err

}

// ValidatorStats is a free data retrieval call binding the contract method 0x3b5eb03a.
//
// Solidity: function validatorStats() view returns(uint256 reportingBlockNumber, uint128 activeValidatorsBalance, uint128 exitedValidatorsBalance, uint128 slashedValidatorsBalance, uint32 activeValidatorsCount, uint32 exitedValidatorsCount, uint32 slashedValidatorsCount)
func (_StaderOracle *StaderOracleSession) ValidatorStats() (struct {
	ReportingBlockNumber     *big.Int
	ActiveValidatorsBalance  *big.Int
	ExitedValidatorsBalance  *big.Int
	SlashedValidatorsBalance *big.Int
	ActiveValidatorsCount    uint32
	ExitedValidatorsCount    uint32
	SlashedValidatorsCount   uint32
}, error) {
	return _StaderOracle.Contract.ValidatorStats(&_StaderOracle.CallOpts)
}

// ValidatorStats is a free data retrieval call binding the contract method 0x3b5eb03a.
//
// Solidity: function validatorStats() view returns(uint256 reportingBlockNumber, uint128 activeValidatorsBalance, uint128 exitedValidatorsBalance, uint128 slashedValidatorsBalance, uint32 activeValidatorsCount, uint32 exitedValidatorsCount, uint32 slashedValidatorsCount)
func (_StaderOracle *StaderOracleCallerSession) ValidatorStats() (struct {
	ReportingBlockNumber     *big.Int
	ActiveValidatorsBalance  *big.Int
	ExitedValidatorsBalance  *big.Int
	SlashedValidatorsBalance *big.Int
	ActiveValidatorsCount    uint32
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

// Initialize is a paid mutator transaction binding the contract method 0xc4d66de8.
//
// Solidity: function initialize(address _staderConfig) returns()
func (_StaderOracle *StaderOracleTransactor) Initialize(opts *bind.TransactOpts, _staderConfig common.Address) (*types.Transaction, error) {
	return _StaderOracle.contract.Transact(opts, "initialize", _staderConfig)
}

// Initialize is a paid mutator transaction binding the contract method 0xc4d66de8.
//
// Solidity: function initialize(address _staderConfig) returns()
func (_StaderOracle *StaderOracleSession) Initialize(_staderConfig common.Address) (*types.Transaction, error) {
	return _StaderOracle.Contract.Initialize(&_StaderOracle.TransactOpts, _staderConfig)
}

// Initialize is a paid mutator transaction binding the contract method 0xc4d66de8.
//
// Solidity: function initialize(address _staderConfig) returns()
func (_StaderOracle *StaderOracleTransactorSession) Initialize(_staderConfig common.Address) (*types.Transaction, error) {
	return _StaderOracle.Contract.Initialize(&_StaderOracle.TransactOpts, _staderConfig)
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

// SetMerkleRootUpdateFrequency is a paid mutator transaction binding the contract method 0x5a583321.
//
// Solidity: function setMerkleRootUpdateFrequency(uint256 _updateFrequency) returns()
func (_StaderOracle *StaderOracleTransactor) SetMerkleRootUpdateFrequency(opts *bind.TransactOpts, _updateFrequency *big.Int) (*types.Transaction, error) {
	return _StaderOracle.contract.Transact(opts, "setMerkleRootUpdateFrequency", _updateFrequency)
}

// SetMerkleRootUpdateFrequency is a paid mutator transaction binding the contract method 0x5a583321.
//
// Solidity: function setMerkleRootUpdateFrequency(uint256 _updateFrequency) returns()
func (_StaderOracle *StaderOracleSession) SetMerkleRootUpdateFrequency(_updateFrequency *big.Int) (*types.Transaction, error) {
	return _StaderOracle.Contract.SetMerkleRootUpdateFrequency(&_StaderOracle.TransactOpts, _updateFrequency)
}

// SetMerkleRootUpdateFrequency is a paid mutator transaction binding the contract method 0x5a583321.
//
// Solidity: function setMerkleRootUpdateFrequency(uint256 _updateFrequency) returns()
func (_StaderOracle *StaderOracleTransactorSession) SetMerkleRootUpdateFrequency(_updateFrequency *big.Int) (*types.Transaction, error) {
	return _StaderOracle.Contract.SetMerkleRootUpdateFrequency(&_StaderOracle.TransactOpts, _updateFrequency)
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

// SetSafeMode is a paid mutator transaction binding the contract method 0x44b7e5f2.
//
// Solidity: function setSafeMode(bool _safeMode) returns()
func (_StaderOracle *StaderOracleTransactor) SetSafeMode(opts *bind.TransactOpts, _safeMode bool) (*types.Transaction, error) {
	return _StaderOracle.contract.Transact(opts, "setSafeMode", _safeMode)
}

// SetSafeMode is a paid mutator transaction binding the contract method 0x44b7e5f2.
//
// Solidity: function setSafeMode(bool _safeMode) returns()
func (_StaderOracle *StaderOracleSession) SetSafeMode(_safeMode bool) (*types.Transaction, error) {
	return _StaderOracle.Contract.SetSafeMode(&_StaderOracle.TransactOpts, _safeMode)
}

// SetSafeMode is a paid mutator transaction binding the contract method 0x44b7e5f2.
//
// Solidity: function setSafeMode(bool _safeMode) returns()
func (_StaderOracle *StaderOracleTransactorSession) SetSafeMode(_safeMode bool) (*types.Transaction, error) {
	return _StaderOracle.Contract.SetSafeMode(&_StaderOracle.TransactOpts, _safeMode)
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

// SubmitBalances is a paid mutator transaction binding the contract method 0xcf06fb72.
//
// Solidity: function submitBalances((uint256,uint256,uint256,uint256) _exchangeRate) returns()
func (_StaderOracle *StaderOracleTransactor) SubmitBalances(opts *bind.TransactOpts, _exchangeRate ExchangeRate) (*types.Transaction, error) {
	return _StaderOracle.contract.Transact(opts, "submitBalances", _exchangeRate)
}

// SubmitBalances is a paid mutator transaction binding the contract method 0xcf06fb72.
//
// Solidity: function submitBalances((uint256,uint256,uint256,uint256) _exchangeRate) returns()
func (_StaderOracle *StaderOracleSession) SubmitBalances(_exchangeRate ExchangeRate) (*types.Transaction, error) {
	return _StaderOracle.Contract.SubmitBalances(&_StaderOracle.TransactOpts, _exchangeRate)
}

// SubmitBalances is a paid mutator transaction binding the contract method 0xcf06fb72.
//
// Solidity: function submitBalances((uint256,uint256,uint256,uint256) _exchangeRate) returns()
func (_StaderOracle *StaderOracleTransactorSession) SubmitBalances(_exchangeRate ExchangeRate) (*types.Transaction, error) {
	return _StaderOracle.Contract.SubmitBalances(&_StaderOracle.TransactOpts, _exchangeRate)
}

// SubmitMissedAttestationPenalties is a paid mutator transaction binding the contract method 0x6b895903.
//
// Solidity: function submitMissedAttestationPenalties((uint16,uint256,uint256,uint256,bytes) _mapd) returns()
func (_StaderOracle *StaderOracleTransactor) SubmitMissedAttestationPenalties(opts *bind.TransactOpts, _mapd MissedAttestationPenaltyData) (*types.Transaction, error) {
	return _StaderOracle.contract.Transact(opts, "submitMissedAttestationPenalties", _mapd)
}

// SubmitMissedAttestationPenalties is a paid mutator transaction binding the contract method 0x6b895903.
//
// Solidity: function submitMissedAttestationPenalties((uint16,uint256,uint256,uint256,bytes) _mapd) returns()
func (_StaderOracle *StaderOracleSession) SubmitMissedAttestationPenalties(_mapd MissedAttestationPenaltyData) (*types.Transaction, error) {
	return _StaderOracle.Contract.SubmitMissedAttestationPenalties(&_StaderOracle.TransactOpts, _mapd)
}

// SubmitMissedAttestationPenalties is a paid mutator transaction binding the contract method 0x6b895903.
//
// Solidity: function submitMissedAttestationPenalties((uint16,uint256,uint256,uint256,bytes) _mapd) returns()
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

// SubmitWithdrawnValidators is a paid mutator transaction binding the contract method 0xf0626d72.
//
// Solidity: function submitWithdrawnValidators((uint256,address,bytes[]) _withdrawnValidators) returns()
func (_StaderOracle *StaderOracleTransactor) SubmitWithdrawnValidators(opts *bind.TransactOpts, _withdrawnValidators WithdrawnValidators) (*types.Transaction, error) {
	return _StaderOracle.contract.Transact(opts, "submitWithdrawnValidators", _withdrawnValidators)
}

// SubmitWithdrawnValidators is a paid mutator transaction binding the contract method 0xf0626d72.
//
// Solidity: function submitWithdrawnValidators((uint256,address,bytes[]) _withdrawnValidators) returns()
func (_StaderOracle *StaderOracleSession) SubmitWithdrawnValidators(_withdrawnValidators WithdrawnValidators) (*types.Transaction, error) {
	return _StaderOracle.Contract.SubmitWithdrawnValidators(&_StaderOracle.TransactOpts, _withdrawnValidators)
}

// SubmitWithdrawnValidators is a paid mutator transaction binding the contract method 0xf0626d72.
//
// Solidity: function submitWithdrawnValidators((uint256,address,bytes[]) _withdrawnValidators) returns()
func (_StaderOracle *StaderOracleTransactorSession) SubmitWithdrawnValidators(_withdrawnValidators WithdrawnValidators) (*types.Transaction, error) {
	return _StaderOracle.Contract.SubmitWithdrawnValidators(&_StaderOracle.TransactOpts, _withdrawnValidators)
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

// StaderOracleBalancesSubmittedIterator is returned from FilterBalancesSubmitted and is used to iterate over the raw logs and unpacked data for BalancesSubmitted events raised by the StaderOracle contract.
type StaderOracleBalancesSubmittedIterator struct {
	Event *StaderOracleBalancesSubmitted // Event containing the contract specifics and raw log

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
func (it *StaderOracleBalancesSubmittedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StaderOracleBalancesSubmitted)
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
		it.Event = new(StaderOracleBalancesSubmitted)
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
func (it *StaderOracleBalancesSubmittedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StaderOracleBalancesSubmittedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StaderOracleBalancesSubmitted represents a BalancesSubmitted event raised by the StaderOracle contract.
type StaderOracleBalancesSubmitted struct {
	From       common.Address
	Block      *big.Int
	TotalEth   *big.Int
	StakingEth *big.Int
	EthxSupply *big.Int
	Time       *big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterBalancesSubmitted is a free log retrieval operation binding the contract event 0xe657a6d6957f4fabb37b86d4d6571e82df061bd2d8a3ede5d197b0b98a5a1bdf.
//
// Solidity: event BalancesSubmitted(address indexed from, uint256 block, uint256 totalEth, uint256 stakingEth, uint256 ethxSupply, uint256 time)
func (_StaderOracle *StaderOracleFilterer) FilterBalancesSubmitted(opts *bind.FilterOpts, from []common.Address) (*StaderOracleBalancesSubmittedIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}

	logs, sub, err := _StaderOracle.contract.FilterLogs(opts, "BalancesSubmitted", fromRule)
	if err != nil {
		return nil, err
	}
	return &StaderOracleBalancesSubmittedIterator{contract: _StaderOracle.contract, event: "BalancesSubmitted", logs: logs, sub: sub}, nil
}

// WatchBalancesSubmitted is a free log subscription operation binding the contract event 0xe657a6d6957f4fabb37b86d4d6571e82df061bd2d8a3ede5d197b0b98a5a1bdf.
//
// Solidity: event BalancesSubmitted(address indexed from, uint256 block, uint256 totalEth, uint256 stakingEth, uint256 ethxSupply, uint256 time)
func (_StaderOracle *StaderOracleFilterer) WatchBalancesSubmitted(opts *bind.WatchOpts, sink chan<- *StaderOracleBalancesSubmitted, from []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}

	logs, sub, err := _StaderOracle.contract.WatchLogs(opts, "BalancesSubmitted", fromRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StaderOracleBalancesSubmitted)
				if err := _StaderOracle.contract.UnpackLog(event, "BalancesSubmitted", log); err != nil {
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

// ParseBalancesSubmitted is a log parse operation binding the contract event 0xe657a6d6957f4fabb37b86d4d6571e82df061bd2d8a3ede5d197b0b98a5a1bdf.
//
// Solidity: event BalancesSubmitted(address indexed from, uint256 block, uint256 totalEth, uint256 stakingEth, uint256 ethxSupply, uint256 time)
func (_StaderOracle *StaderOracleFilterer) ParseBalancesSubmitted(log types.Log) (*StaderOracleBalancesSubmitted, error) {
	event := new(StaderOracleBalancesSubmitted)
	if err := _StaderOracle.contract.UnpackLog(event, "BalancesSubmitted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StaderOracleBalancesUpdatedIterator is returned from FilterBalancesUpdated and is used to iterate over the raw logs and unpacked data for BalancesUpdated events raised by the StaderOracle contract.
type StaderOracleBalancesUpdatedIterator struct {
	Event *StaderOracleBalancesUpdated // Event containing the contract specifics and raw log

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
func (it *StaderOracleBalancesUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StaderOracleBalancesUpdated)
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
		it.Event = new(StaderOracleBalancesUpdated)
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
func (it *StaderOracleBalancesUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StaderOracleBalancesUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StaderOracleBalancesUpdated represents a BalancesUpdated event raised by the StaderOracle contract.
type StaderOracleBalancesUpdated struct {
	Block      *big.Int
	TotalEth   *big.Int
	StakingEth *big.Int
	EthxSupply *big.Int
	Time       *big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterBalancesUpdated is a free log retrieval operation binding the contract event 0x7bbbb137fdad433d6168b1c75c714c72b8abe8d07460f0c0b433063e7bf1f394.
//
// Solidity: event BalancesUpdated(uint256 block, uint256 totalEth, uint256 stakingEth, uint256 ethxSupply, uint256 time)
func (_StaderOracle *StaderOracleFilterer) FilterBalancesUpdated(opts *bind.FilterOpts) (*StaderOracleBalancesUpdatedIterator, error) {

	logs, sub, err := _StaderOracle.contract.FilterLogs(opts, "BalancesUpdated")
	if err != nil {
		return nil, err
	}
	return &StaderOracleBalancesUpdatedIterator{contract: _StaderOracle.contract, event: "BalancesUpdated", logs: logs, sub: sub}, nil
}

// WatchBalancesUpdated is a free log subscription operation binding the contract event 0x7bbbb137fdad433d6168b1c75c714c72b8abe8d07460f0c0b433063e7bf1f394.
//
// Solidity: event BalancesUpdated(uint256 block, uint256 totalEth, uint256 stakingEth, uint256 ethxSupply, uint256 time)
func (_StaderOracle *StaderOracleFilterer) WatchBalancesUpdated(opts *bind.WatchOpts, sink chan<- *StaderOracleBalancesUpdated) (event.Subscription, error) {

	logs, sub, err := _StaderOracle.contract.WatchLogs(opts, "BalancesUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StaderOracleBalancesUpdated)
				if err := _StaderOracle.contract.UnpackLog(event, "BalancesUpdated", log); err != nil {
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

// ParseBalancesUpdated is a log parse operation binding the contract event 0x7bbbb137fdad433d6168b1c75c714c72b8abe8d07460f0c0b433063e7bf1f394.
//
// Solidity: event BalancesUpdated(uint256 block, uint256 totalEth, uint256 stakingEth, uint256 ethxSupply, uint256 time)
func (_StaderOracle *StaderOracleFilterer) ParseBalancesUpdated(log types.Log) (*StaderOracleBalancesUpdated, error) {
	event := new(StaderOracleBalancesUpdated)
	if err := _StaderOracle.contract.UnpackLog(event, "BalancesUpdated", log); err != nil {
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
	PageNumber           *big.Int
	Block                *big.Int
	ReportingBlockNumber *big.Int
	Raw                  types.Log // Blockchain specific contextual infos
}

// FilterMissedAttestationPenaltySubmitted is a free log retrieval operation binding the contract event 0xea8e9b227ba9887fe5c583b382794c05eb4fa60a89a2ee024a8ab86894afc8d0.
//
// Solidity: event MissedAttestationPenaltySubmitted(address indexed node, uint256 index, uint256 pageNumber, uint256 block, uint256 reportingBlockNumber)
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

// WatchMissedAttestationPenaltySubmitted is a free log subscription operation binding the contract event 0xea8e9b227ba9887fe5c583b382794c05eb4fa60a89a2ee024a8ab86894afc8d0.
//
// Solidity: event MissedAttestationPenaltySubmitted(address indexed node, uint256 index, uint256 pageNumber, uint256 block, uint256 reportingBlockNumber)
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

// ParseMissedAttestationPenaltySubmitted is a log parse operation binding the contract event 0xea8e9b227ba9887fe5c583b382794c05eb4fa60a89a2ee024a8ab86894afc8d0.
//
// Solidity: event MissedAttestationPenaltySubmitted(address indexed node, uint256 index, uint256 pageNumber, uint256 block, uint256 reportingBlockNumber)
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
	Index *big.Int
	Block *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterMissedAttestationPenaltyUpdated is a free log retrieval operation binding the contract event 0x73c2611cfd9a60d77e3e57cf6fd72e13a689232db5deb08784b0de3861401ed6.
//
// Solidity: event MissedAttestationPenaltyUpdated(uint256 index, uint256 block)
func (_StaderOracle *StaderOracleFilterer) FilterMissedAttestationPenaltyUpdated(opts *bind.FilterOpts) (*StaderOracleMissedAttestationPenaltyUpdatedIterator, error) {

	logs, sub, err := _StaderOracle.contract.FilterLogs(opts, "MissedAttestationPenaltyUpdated")
	if err != nil {
		return nil, err
	}
	return &StaderOracleMissedAttestationPenaltyUpdatedIterator{contract: _StaderOracle.contract, event: "MissedAttestationPenaltyUpdated", logs: logs, sub: sub}, nil
}

// WatchMissedAttestationPenaltyUpdated is a free log subscription operation binding the contract event 0x73c2611cfd9a60d77e3e57cf6fd72e13a689232db5deb08784b0de3861401ed6.
//
// Solidity: event MissedAttestationPenaltyUpdated(uint256 index, uint256 block)
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

// ParseMissedAttestationPenaltyUpdated is a log parse operation binding the contract event 0x73c2611cfd9a60d77e3e57cf6fd72e13a689232db5deb08784b0de3861401ed6.
//
// Solidity: event MissedAttestationPenaltyUpdated(uint256 index, uint256 block)
func (_StaderOracle *StaderOracleFilterer) ParseMissedAttestationPenaltyUpdated(log types.Log) (*StaderOracleMissedAttestationPenaltyUpdated, error) {
	event := new(StaderOracleMissedAttestationPenaltyUpdated)
	if err := _StaderOracle.contract.UnpackLog(event, "MissedAttestationPenaltyUpdated", log); err != nil {
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
	Block      *big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterSocializingRewardsMerkleRootSubmitted is a free log retrieval operation binding the contract event 0x910be9ecf7e303516c34fc1a37ed32381e228262fd49d11837b4d72575aaa412.
//
// Solidity: event SocializingRewardsMerkleRootSubmitted(address indexed node, uint256 index, bytes32 merkleRoot, uint256 block)
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

// WatchSocializingRewardsMerkleRootSubmitted is a free log subscription operation binding the contract event 0x910be9ecf7e303516c34fc1a37ed32381e228262fd49d11837b4d72575aaa412.
//
// Solidity: event SocializingRewardsMerkleRootSubmitted(address indexed node, uint256 index, bytes32 merkleRoot, uint256 block)
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

// ParseSocializingRewardsMerkleRootSubmitted is a log parse operation binding the contract event 0x910be9ecf7e303516c34fc1a37ed32381e228262fd49d11837b4d72575aaa412.
//
// Solidity: event SocializingRewardsMerkleRootSubmitted(address indexed node, uint256 index, bytes32 merkleRoot, uint256 block)
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
	Block      *big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterSocializingRewardsMerkleRootUpdated is a free log retrieval operation binding the contract event 0xe48c5fc63a8909e91fbfa2a2aa57b864359869d20d9187a723dedc4d5ea7cb1b.
//
// Solidity: event SocializingRewardsMerkleRootUpdated(uint256 index, bytes32 merkleRoot, uint256 block)
func (_StaderOracle *StaderOracleFilterer) FilterSocializingRewardsMerkleRootUpdated(opts *bind.FilterOpts) (*StaderOracleSocializingRewardsMerkleRootUpdatedIterator, error) {

	logs, sub, err := _StaderOracle.contract.FilterLogs(opts, "SocializingRewardsMerkleRootUpdated")
	if err != nil {
		return nil, err
	}
	return &StaderOracleSocializingRewardsMerkleRootUpdatedIterator{contract: _StaderOracle.contract, event: "SocializingRewardsMerkleRootUpdated", logs: logs, sub: sub}, nil
}

// WatchSocializingRewardsMerkleRootUpdated is a free log subscription operation binding the contract event 0xe48c5fc63a8909e91fbfa2a2aa57b864359869d20d9187a723dedc4d5ea7cb1b.
//
// Solidity: event SocializingRewardsMerkleRootUpdated(uint256 index, bytes32 merkleRoot, uint256 block)
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

// ParseSocializingRewardsMerkleRootUpdated is a log parse operation binding the contract event 0xe48c5fc63a8909e91fbfa2a2aa57b864359869d20d9187a723dedc4d5ea7cb1b.
//
// Solidity: event SocializingRewardsMerkleRootUpdated(uint256 index, bytes32 merkleRoot, uint256 block)
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

// StaderOracleUpdatedSafeModeIterator is returned from FilterUpdatedSafeMode and is used to iterate over the raw logs and unpacked data for UpdatedSafeMode events raised by the StaderOracle contract.
type StaderOracleUpdatedSafeModeIterator struct {
	Event *StaderOracleUpdatedSafeMode // Event containing the contract specifics and raw log

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
func (it *StaderOracleUpdatedSafeModeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StaderOracleUpdatedSafeMode)
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
		it.Event = new(StaderOracleUpdatedSafeMode)
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
func (it *StaderOracleUpdatedSafeModeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StaderOracleUpdatedSafeModeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StaderOracleUpdatedSafeMode represents a UpdatedSafeMode event raised by the StaderOracle contract.
type StaderOracleUpdatedSafeMode struct {
	SafeMode bool
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterUpdatedSafeMode is a free log retrieval operation binding the contract event 0xde98c74f54ffa25bf71add6ec149cd942ed4ec368c2841f8be61a2a2896bc306.
//
// Solidity: event UpdatedSafeMode(bool safeMode)
func (_StaderOracle *StaderOracleFilterer) FilterUpdatedSafeMode(opts *bind.FilterOpts) (*StaderOracleUpdatedSafeModeIterator, error) {

	logs, sub, err := _StaderOracle.contract.FilterLogs(opts, "UpdatedSafeMode")
	if err != nil {
		return nil, err
	}
	return &StaderOracleUpdatedSafeModeIterator{contract: _StaderOracle.contract, event: "UpdatedSafeMode", logs: logs, sub: sub}, nil
}

// WatchUpdatedSafeMode is a free log subscription operation binding the contract event 0xde98c74f54ffa25bf71add6ec149cd942ed4ec368c2841f8be61a2a2896bc306.
//
// Solidity: event UpdatedSafeMode(bool safeMode)
func (_StaderOracle *StaderOracleFilterer) WatchUpdatedSafeMode(opts *bind.WatchOpts, sink chan<- *StaderOracleUpdatedSafeMode) (event.Subscription, error) {

	logs, sub, err := _StaderOracle.contract.WatchLogs(opts, "UpdatedSafeMode")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StaderOracleUpdatedSafeMode)
				if err := _StaderOracle.contract.UnpackLog(event, "UpdatedSafeMode", log); err != nil {
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

// ParseUpdatedSafeMode is a log parse operation binding the contract event 0xde98c74f54ffa25bf71add6ec149cd942ed4ec368c2841f8be61a2a2896bc306.
//
// Solidity: event UpdatedSafeMode(bool safeMode)
func (_StaderOracle *StaderOracleFilterer) ParseUpdatedSafeMode(log types.Log) (*StaderOracleUpdatedSafeMode, error) {
	event := new(StaderOracleUpdatedSafeMode)
	if err := _StaderOracle.contract.UnpackLog(event, "UpdatedSafeMode", log); err != nil {
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
	From         common.Address
	Block        *big.Int
	NodeRegistry common.Address
	Pubkeys      [][]byte
	Time         *big.Int
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterWithdrawnValidatorsSubmitted is a free log retrieval operation binding the contract event 0x1e004d900c9787bf2cfed87d7704cf2a6b1956a7141dab327aea41eb6aa143f2.
//
// Solidity: event WithdrawnValidatorsSubmitted(address indexed from, uint256 block, address nodeRegistry, bytes[] pubkeys, uint256 time)
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

// WatchWithdrawnValidatorsSubmitted is a free log subscription operation binding the contract event 0x1e004d900c9787bf2cfed87d7704cf2a6b1956a7141dab327aea41eb6aa143f2.
//
// Solidity: event WithdrawnValidatorsSubmitted(address indexed from, uint256 block, address nodeRegistry, bytes[] pubkeys, uint256 time)
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

// ParseWithdrawnValidatorsSubmitted is a log parse operation binding the contract event 0x1e004d900c9787bf2cfed87d7704cf2a6b1956a7141dab327aea41eb6aa143f2.
//
// Solidity: event WithdrawnValidatorsSubmitted(address indexed from, uint256 block, address nodeRegistry, bytes[] pubkeys, uint256 time)
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
	Block        *big.Int
	NodeRegistry common.Address
	Pubkeys      [][]byte
	Time         *big.Int
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterWithdrawnValidatorsUpdated is a free log retrieval operation binding the contract event 0x70a26ae3b547e42cc7a3819a181db1a108440c2dedd23eb34c2357680cfd056b.
//
// Solidity: event WithdrawnValidatorsUpdated(uint256 block, address nodeRegistry, bytes[] pubkeys, uint256 time)
func (_StaderOracle *StaderOracleFilterer) FilterWithdrawnValidatorsUpdated(opts *bind.FilterOpts) (*StaderOracleWithdrawnValidatorsUpdatedIterator, error) {

	logs, sub, err := _StaderOracle.contract.FilterLogs(opts, "WithdrawnValidatorsUpdated")
	if err != nil {
		return nil, err
	}
	return &StaderOracleWithdrawnValidatorsUpdatedIterator{contract: _StaderOracle.contract, event: "WithdrawnValidatorsUpdated", logs: logs, sub: sub}, nil
}

// WatchWithdrawnValidatorsUpdated is a free log subscription operation binding the contract event 0x70a26ae3b547e42cc7a3819a181db1a108440c2dedd23eb34c2357680cfd056b.
//
// Solidity: event WithdrawnValidatorsUpdated(uint256 block, address nodeRegistry, bytes[] pubkeys, uint256 time)
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

// ParseWithdrawnValidatorsUpdated is a log parse operation binding the contract event 0x70a26ae3b547e42cc7a3819a181db1a108440c2dedd23eb34c2357680cfd056b.
//
// Solidity: event WithdrawnValidatorsUpdated(uint256 block, address nodeRegistry, bytes[] pubkeys, uint256 time)
func (_StaderOracle *StaderOracleFilterer) ParseWithdrawnValidatorsUpdated(log types.Log) (*StaderOracleWithdrawnValidatorsUpdated, error) {
	event := new(StaderOracleWithdrawnValidatorsUpdated)
	if err := _StaderOracle.contract.UnpackLog(event, "WithdrawnValidatorsUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
