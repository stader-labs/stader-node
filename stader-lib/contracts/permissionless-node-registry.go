/*
This work is licensed and released under GNU GPL v3 or any other later versions. 
The full text of the license is below/ found at <http://www.gnu.org/licenses/>

(c) 2023 Rocket Pool Pty Ltd. Modified under GNU GPL v3.

This program is free software: you can redistribute it and/or modify
it under the terms of the GNU General Public License as published by
the Free Software Foundation, either version 3 of the License, or
(at your option) any later version.

This program is distributed in the hope that it will be useful,
but WITHOUT ANY WARRANTY; without even the implied warranty of
MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
GNU General Public License for more details.

You should have received a copy of the GNU General Public License
along with this program.  If not, see <http://www.gnu.org/licenses/>.
*/
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

// Operator is an auto generated low-level Go binding around an user-defined struct.
type Operator struct {
	Active                  bool
	OptedForSocializingPool bool
	OperatorName            string
	OperatorRewardAddress   common.Address
	OperatorAddress         common.Address
}

// Validator is an auto generated low-level Go binding around an user-defined struct.
type Validator struct {
	Status               uint8
	Pubkey               []byte
	PreDepositSignature  []byte
	DepositSignature     []byte
	WithdrawVaultAddress common.Address
	OperatorId           *big.Int
	InitialBondEth       *big.Int
	DepositTime          *big.Int
	WithdrawnTime        *big.Int
}

// PermissionlessNodeRegistryMetaData contains all meta data concerning the PermissionlessNodeRegistry contract.
var PermissionlessNodeRegistryMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"name\":\"CooldownNotComplete\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"EmptyNameString\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InSufficientBalance\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidBondEthValue\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidKeyCount\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidLengthOfPubkey\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidLengthOfSignature\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidStartAndEndIndex\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"MisMatchingInputKeysSize\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NameCrossedMaxLength\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NoChangeInState\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OperatorAlreadyOnBoarded\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OperatorIsDeactivate\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OperatorNotOnBoarded\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"PubkeyAlreadyExist\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"PubkeyDoesNotExist\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"PubkeyNotFoundOrDuplicateInput\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"TransferFailed\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"UNEXPECTED_STATUS\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ZeroAddress\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"maxKeyLimitReached\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_nodeOperator\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"_pubkey\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_validatorId\",\"type\":\"uint256\"}],\"name\":\"AddedKeys\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"version\",\"type\":\"uint8\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_nodeOperator\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_operatorId\",\"type\":\"uint256\"}],\"name\":\"OnboardedOperator\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Paused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"previousAdminRole\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"newAdminRole\",\"type\":\"bytes32\"}],\"name\":\"RoleAdminChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleGranted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleRevoked\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Unpaused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_elRewardSocializePool\",\"type\":\"address\"}],\"name\":\"UpdatedELRewardSocializePool\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint16\",\"name\":\"_inputKeyCountLimit\",\"type\":\"uint16\"}],\"name\":\"UpdatedInputKeyCountLimit\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"_keyDepositLimit\",\"type\":\"uint64\"}],\"name\":\"UpdatedMaxKeyPerOperator\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_nextQueuedValidatorIndex\",\"type\":\"uint256\"}],\"name\":\"UpdatedNextQueuedValidatorIndex\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_nodeOperator\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"_operatorName\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_rewardAddress\",\"type\":\"address\"}],\"name\":\"UpdatedOperatorDetails\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_permissionlessPool\",\"type\":\"address\"}],\"name\":\"UpdatedPermissionlessPoolAddress\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_poolFactoryAddress\",\"type\":\"address\"}],\"name\":\"UpdatedPoolFactoryAddress\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_sdCollateral\",\"type\":\"address\"}],\"name\":\"UpdatedSDCollateralAddress\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_operatorId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"_optedForSocializingPool\",\"type\":\"bool\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"}],\"name\":\"UpdatedSocializingPoolState\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_staderPenaltyFund\",\"type\":\"address\"}],\"name\":\"UpdatedStaderPenaltyFund\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_vaultFactoryAddress\",\"type\":\"address\"}],\"name\":\"UpdatedVaultFactoryAddress\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_validatorId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_depositTime\",\"type\":\"uint256\"}],\"name\":\"ValidatorDepositTimeSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes\",\"name\":\"_frontRunnedPubkey\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_validatorId\",\"type\":\"uint256\"}],\"name\":\"ValidatorMarkedAsFrontRunned\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes\",\"name\":\"_pubkey\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_validatorId\",\"type\":\"uint256\"}],\"name\":\"ValidatorMarkedReadyToDeposit\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes\",\"name\":\"invalidSignaturePubkey\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_validatorId\",\"type\":\"uint256\"}],\"name\":\"ValidatorStatusMarkedAsInvalidSignature\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes\",\"name\":\"_pubkey\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_validatorId\",\"type\":\"uint256\"}],\"name\":\"ValidatorWithdrawn\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"DEFAULT_ADMIN_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"FRONT_RUN_PENALTY\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"OPERATOR_MAX_NAME_LENGTH\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"PERMISSIONLESS_NODE_REGISTRY_OWNER\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"PERMISSIONLESS_POOL\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"PRE_DEPOSIT\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"STADER_ORACLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"VALIDATOR_STATUS_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes[]\",\"name\":\"_pubkey\",\"type\":\"bytes[]\"},{\"internalType\":\"bytes[]\",\"name\":\"_preDepositSignature\",\"type\":\"bytes[]\"},{\"internalType\":\"bytes[]\",\"name\":\"_depositSignature\",\"type\":\"bytes[]\"}],\"name\":\"addValidatorKeys\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bool\",\"name\":\"_optInForSocializingPool\",\"type\":\"bool\"}],\"name\":\"changeSocializingPoolState\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"feeRecipientAddress\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"collateralETH\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"elRewardSocializePool\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getAllActiveValidators\",\"outputs\":[{\"components\":[{\"internalType\":\"enumValidatorStatus\",\"name\":\"status\",\"type\":\"uint8\"},{\"internalType\":\"bytes\",\"name\":\"pubkey\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"preDepositSignature\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"depositSignature\",\"type\":\"bytes\"},{\"internalType\":\"address\",\"name\":\"withdrawVaultAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"operatorId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"initialBondEth\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"depositTime\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"withdrawnTime\",\"type\":\"uint256\"}],\"internalType\":\"structValidator[]\",\"name\":\"\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getCollateralETH\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"_pubkey\",\"type\":\"bytes\"}],\"name\":\"getOperator\",\"outputs\":[{\"components\":[{\"internalType\":\"bool\",\"name\":\"active\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"optedForSocializingPool\",\"type\":\"bool\"},{\"internalType\":\"string\",\"name\":\"operatorName\",\"type\":\"string\"},{\"internalType\":\"addresspayable\",\"name\":\"operatorRewardAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"operatorAddress\",\"type\":\"address\"}],\"internalType\":\"structOperator\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_operatorId\",\"type\":\"uint256\"}],\"name\":\"getOperatorRewardAddress\",\"outputs\":[{\"internalType\":\"addresspayable\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_operatorId\",\"type\":\"uint256\"}],\"name\":\"getOperatorTotalKeys\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"_totalKeys\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_nodeOperator\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"startIndex\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"endIndex\",\"type\":\"uint256\"}],\"name\":\"getOperatorTotalNonTerminalKeys\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"}],\"name\":\"getRoleAdmin\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_operatorId\",\"type\":\"uint256\"}],\"name\":\"getSocializingPoolStateChangeTimestamp\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getTotalActiveValidatorCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getTotalQueuedValidatorCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"_pubkey\",\"type\":\"bytes\"}],\"name\":\"getValidator\",\"outputs\":[{\"components\":[{\"internalType\":\"enumValidatorStatus\",\"name\":\"status\",\"type\":\"uint8\"},{\"internalType\":\"bytes\",\"name\":\"pubkey\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"preDepositSignature\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"depositSignature\",\"type\":\"bytes\"},{\"internalType\":\"address\",\"name\":\"withdrawVaultAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"operatorId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"initialBondEth\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"depositTime\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"withdrawnTime\",\"type\":\"uint256\"}],\"internalType\":\"structValidator\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_validatorId\",\"type\":\"uint256\"}],\"name\":\"getValidator\",\"outputs\":[{\"components\":[{\"internalType\":\"enumValidatorStatus\",\"name\":\"status\",\"type\":\"uint8\"},{\"internalType\":\"bytes\",\"name\":\"pubkey\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"preDepositSignature\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"depositSignature\",\"type\":\"bytes\"},{\"internalType\":\"address\",\"name\":\"withdrawVaultAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"operatorId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"initialBondEth\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"depositTime\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"withdrawnTime\",\"type\":\"uint256\"}],\"internalType\":\"structValidator\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"grantRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"hasRole\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_count\",\"type\":\"uint256\"}],\"name\":\"increaseTotalActiveValidatorCount\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_adminOwner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_sdCollateral\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_staderPenaltyFund\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_vaultFactoryAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_elRewardSocializePool\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_poolFactoryAddress\",\"type\":\"address\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"inputKeyCountLimit\",\"outputs\":[{\"internalType\":\"uint16\",\"name\":\"\",\"type\":\"uint16\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"_pubkey\",\"type\":\"bytes\"}],\"name\":\"isExistingPubkey\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes[]\",\"name\":\"_readyToDepositPubkey\",\"type\":\"bytes[]\"},{\"internalType\":\"bytes[]\",\"name\":\"_frontRunnedPubkey\",\"type\":\"bytes[]\"},{\"internalType\":\"bytes[]\",\"name\":\"_invalidSignaturePubkey\",\"type\":\"bytes[]\"}],\"name\":\"markValidatorReadyToDeposit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"maxKeyPerOperator\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"nextOperatorId\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"nextQueuedValidatorIndex\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"nextValidatorId\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bool\",\"name\":\"_optInForSocializingPool\",\"type\":\"bool\"},{\"internalType\":\"string\",\"name\":\"_operatorName\",\"type\":\"string\"},{\"internalType\":\"addresspayable\",\"name\":\"_operatorRewardAddress\",\"type\":\"address\"}],\"name\":\"onboardNodeOperator\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"feeRecipientAddress\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"operatorIDByAddress\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"operatorStructById\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"active\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"optedForSocializingPool\",\"type\":\"bool\"},{\"internalType\":\"string\",\"name\":\"operatorName\",\"type\":\"string\"},{\"internalType\":\"addresspayable\",\"name\":\"operatorRewardAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"operatorAddress\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"paused\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"permissionlessPool\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"poolFactoryAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"poolId\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"queuedValidators\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"renounceRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"revokeRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"sdCollateral\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"socializePoolRewardDistributionCycle\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"socializingPoolStateChangeTimestamp\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"staderPenaltyFund\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalActiveValidatorCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"transferCollateralToPool\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"unpause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_validatorId\",\"type\":\"uint256\"}],\"name\":\"updateDepositStatusAndTime\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_elRewardSocializePool\",\"type\":\"address\"}],\"name\":\"updateELRewardSocializePool\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint16\",\"name\":\"_inputKeyCountLimit\",\"type\":\"uint16\"}],\"name\":\"updateInputKeyCountLimit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"_maxKeyPerOperator\",\"type\":\"uint64\"}],\"name\":\"updateMaxKeyPerOperator\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_nextQueuedValidatorIndex\",\"type\":\"uint256\"}],\"name\":\"updateNextQueuedValidatorIndex\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_operatorName\",\"type\":\"string\"},{\"internalType\":\"addresspayable\",\"name\":\"_rewardAddress\",\"type\":\"address\"}],\"name\":\"updateOperatorDetails\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_permissionlessPool\",\"type\":\"address\"}],\"name\":\"updatePermissionlessPoolAddress\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_poolFactoryAddress\",\"type\":\"address\"}],\"name\":\"updatePoolFactoryAddress\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_sdCollateral\",\"type\":\"address\"}],\"name\":\"updateSDCollateralAddress\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_staderPenaltyFund\",\"type\":\"address\"}],\"name\":\"updateStaderPenaltyFundAddress\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"_pubkey\",\"type\":\"bytes\"},{\"internalType\":\"enumValidatorStatus\",\"name\":\"_status\",\"type\":\"uint8\"}],\"name\":\"updateValidatorStatus\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_vaultFactoryAddress\",\"type\":\"address\"}],\"name\":\"updateVaultFactoryAddress\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"name\":\"validatorIdByPubkey\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"validatorIdsByOperatorId\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"validatorQueueSize\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"validatorRegistry\",\"outputs\":[{\"internalType\":\"enumValidatorStatus\",\"name\":\"status\",\"type\":\"uint8\"},{\"internalType\":\"bytes\",\"name\":\"pubkey\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"preDepositSignature\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"depositSignature\",\"type\":\"bytes\"},{\"internalType\":\"address\",\"name\":\"withdrawVaultAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"operatorId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"initialBondEth\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"depositTime\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"withdrawnTime\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"vaultFactoryAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes[]\",\"name\":\"_pubkeys\",\"type\":\"bytes[]\"}],\"name\":\"withdrawnValidators\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// PermissionlessNodeRegistryABI is the input ABI used to generate the binding from.
// Deprecated: Use PermissionlessNodeRegistryMetaData.ABI instead.
var PermissionlessNodeRegistryABI = PermissionlessNodeRegistryMetaData.ABI

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

// OPERATORMAXNAMELENGTH is a free data retrieval call binding the contract method 0x5455e472.
//
// Solidity: function OPERATOR_MAX_NAME_LENGTH() view returns(uint256)
func (_PermissionlessNodeRegistry *PermissionlessNodeRegistryCaller) OPERATORMAXNAMELENGTH(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _PermissionlessNodeRegistry.contract.Call(opts, &out, "OPERATOR_MAX_NAME_LENGTH")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// OPERATORMAXNAMELENGTH is a free data retrieval call binding the contract method 0x5455e472.
//
// Solidity: function OPERATOR_MAX_NAME_LENGTH() view returns(uint256)
func (_PermissionlessNodeRegistry *PermissionlessNodeRegistrySession) OPERATORMAXNAMELENGTH() (*big.Int, error) {
	return _PermissionlessNodeRegistry.Contract.OPERATORMAXNAMELENGTH(&_PermissionlessNodeRegistry.CallOpts)
}

// OPERATORMAXNAMELENGTH is a free data retrieval call binding the contract method 0x5455e472.
//
// Solidity: function OPERATOR_MAX_NAME_LENGTH() view returns(uint256)
func (_PermissionlessNodeRegistry *PermissionlessNodeRegistryCallerSession) OPERATORMAXNAMELENGTH() (*big.Int, error) {
	return _PermissionlessNodeRegistry.Contract.OPERATORMAXNAMELENGTH(&_PermissionlessNodeRegistry.CallOpts)
}

// PERMISSIONLESSNODEREGISTRYOWNER is a free data retrieval call binding the contract method 0x0d138435.
//
// Solidity: function PERMISSIONLESS_NODE_REGISTRY_OWNER() view returns(bytes32)
func (_PermissionlessNodeRegistry *PermissionlessNodeRegistryCaller) PERMISSIONLESSNODEREGISTRYOWNER(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _PermissionlessNodeRegistry.contract.Call(opts, &out, "PERMISSIONLESS_NODE_REGISTRY_OWNER")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// PERMISSIONLESSNODEREGISTRYOWNER is a free data retrieval call binding the contract method 0x0d138435.
//
// Solidity: function PERMISSIONLESS_NODE_REGISTRY_OWNER() view returns(bytes32)
func (_PermissionlessNodeRegistry *PermissionlessNodeRegistrySession) PERMISSIONLESSNODEREGISTRYOWNER() ([32]byte, error) {
	return _PermissionlessNodeRegistry.Contract.PERMISSIONLESSNODEREGISTRYOWNER(&_PermissionlessNodeRegistry.CallOpts)
}

// PERMISSIONLESSNODEREGISTRYOWNER is a free data retrieval call binding the contract method 0x0d138435.
//
// Solidity: function PERMISSIONLESS_NODE_REGISTRY_OWNER() view returns(bytes32)
func (_PermissionlessNodeRegistry *PermissionlessNodeRegistryCallerSession) PERMISSIONLESSNODEREGISTRYOWNER() ([32]byte, error) {
	return _PermissionlessNodeRegistry.Contract.PERMISSIONLESSNODEREGISTRYOWNER(&_PermissionlessNodeRegistry.CallOpts)
}

// PERMISSIONLESSPOOL is a free data retrieval call binding the contract method 0x7a87fa0b.
//
// Solidity: function PERMISSIONLESS_POOL() view returns(bytes32)
func (_PermissionlessNodeRegistry *PermissionlessNodeRegistryCaller) PERMISSIONLESSPOOL(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _PermissionlessNodeRegistry.contract.Call(opts, &out, "PERMISSIONLESS_POOL")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// PERMISSIONLESSPOOL is a free data retrieval call binding the contract method 0x7a87fa0b.
//
// Solidity: function PERMISSIONLESS_POOL() view returns(bytes32)
func (_PermissionlessNodeRegistry *PermissionlessNodeRegistrySession) PERMISSIONLESSPOOL() ([32]byte, error) {
	return _PermissionlessNodeRegistry.Contract.PERMISSIONLESSPOOL(&_PermissionlessNodeRegistry.CallOpts)
}

// PERMISSIONLESSPOOL is a free data retrieval call binding the contract method 0x7a87fa0b.
//
// Solidity: function PERMISSIONLESS_POOL() view returns(bytes32)
func (_PermissionlessNodeRegistry *PermissionlessNodeRegistryCallerSession) PERMISSIONLESSPOOL() ([32]byte, error) {
	return _PermissionlessNodeRegistry.Contract.PERMISSIONLESSPOOL(&_PermissionlessNodeRegistry.CallOpts)
}

// PREDEPOSIT is a free data retrieval call binding the contract method 0x3703e09b.
//
// Solidity: function PRE_DEPOSIT() view returns(uint256)
func (_PermissionlessNodeRegistry *PermissionlessNodeRegistryCaller) PREDEPOSIT(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _PermissionlessNodeRegistry.contract.Call(opts, &out, "PRE_DEPOSIT")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PREDEPOSIT is a free data retrieval call binding the contract method 0x3703e09b.
//
// Solidity: function PRE_DEPOSIT() view returns(uint256)
func (_PermissionlessNodeRegistry *PermissionlessNodeRegistrySession) PREDEPOSIT() (*big.Int, error) {
	return _PermissionlessNodeRegistry.Contract.PREDEPOSIT(&_PermissionlessNodeRegistry.CallOpts)
}

// PREDEPOSIT is a free data retrieval call binding the contract method 0x3703e09b.
//
// Solidity: function PRE_DEPOSIT() view returns(uint256)
func (_PermissionlessNodeRegistry *PermissionlessNodeRegistryCallerSession) PREDEPOSIT() (*big.Int, error) {
	return _PermissionlessNodeRegistry.Contract.PREDEPOSIT(&_PermissionlessNodeRegistry.CallOpts)
}

// STADERORACLE is a free data retrieval call binding the contract method 0x3871d0f1.
//
// Solidity: function STADER_ORACLE() view returns(bytes32)
func (_PermissionlessNodeRegistry *PermissionlessNodeRegistryCaller) STADERORACLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _PermissionlessNodeRegistry.contract.Call(opts, &out, "STADER_ORACLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// STADERORACLE is a free data retrieval call binding the contract method 0x3871d0f1.
//
// Solidity: function STADER_ORACLE() view returns(bytes32)
func (_PermissionlessNodeRegistry *PermissionlessNodeRegistrySession) STADERORACLE() ([32]byte, error) {
	return _PermissionlessNodeRegistry.Contract.STADERORACLE(&_PermissionlessNodeRegistry.CallOpts)
}

// STADERORACLE is a free data retrieval call binding the contract method 0x3871d0f1.
//
// Solidity: function STADER_ORACLE() view returns(bytes32)
func (_PermissionlessNodeRegistry *PermissionlessNodeRegistryCallerSession) STADERORACLE() ([32]byte, error) {
	return _PermissionlessNodeRegistry.Contract.STADERORACLE(&_PermissionlessNodeRegistry.CallOpts)
}

// VALIDATORSTATUSROLE is a free data retrieval call binding the contract method 0x4da1145a.
//
// Solidity: function VALIDATOR_STATUS_ROLE() view returns(bytes32)
func (_PermissionlessNodeRegistry *PermissionlessNodeRegistryCaller) VALIDATORSTATUSROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _PermissionlessNodeRegistry.contract.Call(opts, &out, "VALIDATOR_STATUS_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// VALIDATORSTATUSROLE is a free data retrieval call binding the contract method 0x4da1145a.
//
// Solidity: function VALIDATOR_STATUS_ROLE() view returns(bytes32)
func (_PermissionlessNodeRegistry *PermissionlessNodeRegistrySession) VALIDATORSTATUSROLE() ([32]byte, error) {
	return _PermissionlessNodeRegistry.Contract.VALIDATORSTATUSROLE(&_PermissionlessNodeRegistry.CallOpts)
}

// VALIDATORSTATUSROLE is a free data retrieval call binding the contract method 0x4da1145a.
//
// Solidity: function VALIDATOR_STATUS_ROLE() view returns(bytes32)
func (_PermissionlessNodeRegistry *PermissionlessNodeRegistryCallerSession) VALIDATORSTATUSROLE() ([32]byte, error) {
	return _PermissionlessNodeRegistry.Contract.VALIDATORSTATUSROLE(&_PermissionlessNodeRegistry.CallOpts)
}

// CollateralETH is a free data retrieval call binding the contract method 0xf75ea242.
//
// Solidity: function collateralETH() view returns(uint256)
func (_PermissionlessNodeRegistry *PermissionlessNodeRegistryCaller) CollateralETH(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _PermissionlessNodeRegistry.contract.Call(opts, &out, "collateralETH")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CollateralETH is a free data retrieval call binding the contract method 0xf75ea242.
//
// Solidity: function collateralETH() view returns(uint256)
func (_PermissionlessNodeRegistry *PermissionlessNodeRegistrySession) CollateralETH() (*big.Int, error) {
	return _PermissionlessNodeRegistry.Contract.CollateralETH(&_PermissionlessNodeRegistry.CallOpts)
}

// CollateralETH is a free data retrieval call binding the contract method 0xf75ea242.
//
// Solidity: function collateralETH() view returns(uint256)
func (_PermissionlessNodeRegistry *PermissionlessNodeRegistryCallerSession) CollateralETH() (*big.Int, error) {
	return _PermissionlessNodeRegistry.Contract.CollateralETH(&_PermissionlessNodeRegistry.CallOpts)
}

// ElRewardSocializePool is a free data retrieval call binding the contract method 0xad62c227.
//
// Solidity: function elRewardSocializePool() view returns(address)
func (_PermissionlessNodeRegistry *PermissionlessNodeRegistryCaller) ElRewardSocializePool(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _PermissionlessNodeRegistry.contract.Call(opts, &out, "elRewardSocializePool")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// ElRewardSocializePool is a free data retrieval call binding the contract method 0xad62c227.
//
// Solidity: function elRewardSocializePool() view returns(address)
func (_PermissionlessNodeRegistry *PermissionlessNodeRegistrySession) ElRewardSocializePool() (common.Address, error) {
	return _PermissionlessNodeRegistry.Contract.ElRewardSocializePool(&_PermissionlessNodeRegistry.CallOpts)
}

// ElRewardSocializePool is a free data retrieval call binding the contract method 0xad62c227.
//
// Solidity: function elRewardSocializePool() view returns(address)
func (_PermissionlessNodeRegistry *PermissionlessNodeRegistryCallerSession) ElRewardSocializePool() (common.Address, error) {
	return _PermissionlessNodeRegistry.Contract.ElRewardSocializePool(&_PermissionlessNodeRegistry.CallOpts)
}

// GetAllActiveValidators is a free data retrieval call binding the contract method 0xe1fe9d15.
//
// Solidity: function getAllActiveValidators() view returns((uint8,bytes,bytes,bytes,address,uint256,uint256,uint256,uint256)[])
func (_PermissionlessNodeRegistry *PermissionlessNodeRegistryCaller) GetAllActiveValidators(opts *bind.CallOpts) ([]Validator, error) {
	var out []interface{}
	err := _PermissionlessNodeRegistry.contract.Call(opts, &out, "getAllActiveValidators")

	if err != nil {
		return *new([]Validator), err
	}

	out0 := *abi.ConvertType(out[0], new([]Validator)).(*[]Validator)

	return out0, err

}

// GetAllActiveValidators is a free data retrieval call binding the contract method 0xe1fe9d15.
//
// Solidity: function getAllActiveValidators() view returns((uint8,bytes,bytes,bytes,address,uint256,uint256,uint256,uint256)[])
func (_PermissionlessNodeRegistry *PermissionlessNodeRegistrySession) GetAllActiveValidators() ([]Validator, error) {
	return _PermissionlessNodeRegistry.Contract.GetAllActiveValidators(&_PermissionlessNodeRegistry.CallOpts)
}

// GetAllActiveValidators is a free data retrieval call binding the contract method 0xe1fe9d15.
//
// Solidity: function getAllActiveValidators() view returns((uint8,bytes,bytes,bytes,address,uint256,uint256,uint256,uint256)[])
func (_PermissionlessNodeRegistry *PermissionlessNodeRegistryCallerSession) GetAllActiveValidators() ([]Validator, error) {
	return _PermissionlessNodeRegistry.Contract.GetAllActiveValidators(&_PermissionlessNodeRegistry.CallOpts)
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

// GetOperator is a free data retrieval call binding the contract method 0x9eaffa96.
//
// Solidity: function getOperator(bytes _pubkey) view returns((bool,bool,string,address,address))
func (_PermissionlessNodeRegistry *PermissionlessNodeRegistryCaller) GetOperator(opts *bind.CallOpts, _pubkey []byte) (Operator, error) {
	var out []interface{}
	err := _PermissionlessNodeRegistry.contract.Call(opts, &out, "getOperator", _pubkey)

	if err != nil {
		return *new(Operator), err
	}

	out0 := *abi.ConvertType(out[0], new(Operator)).(*Operator)

	return out0, err

}

// GetOperator is a free data retrieval call binding the contract method 0x9eaffa96.
//
// Solidity: function getOperator(bytes _pubkey) view returns((bool,bool,string,address,address))
func (_PermissionlessNodeRegistry *PermissionlessNodeRegistrySession) GetOperator(_pubkey []byte) (Operator, error) {
	return _PermissionlessNodeRegistry.Contract.GetOperator(&_PermissionlessNodeRegistry.CallOpts, _pubkey)
}

// GetOperator is a free data retrieval call binding the contract method 0x9eaffa96.
//
// Solidity: function getOperator(bytes _pubkey) view returns((bool,bool,string,address,address))
func (_PermissionlessNodeRegistry *PermissionlessNodeRegistryCallerSession) GetOperator(_pubkey []byte) (Operator, error) {
	return _PermissionlessNodeRegistry.Contract.GetOperator(&_PermissionlessNodeRegistry.CallOpts, _pubkey)
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
// Solidity: function getOperatorTotalNonTerminalKeys(address _nodeOperator, uint256 startIndex, uint256 endIndex) view returns(uint64)
func (_PermissionlessNodeRegistry *PermissionlessNodeRegistryCaller) GetOperatorTotalNonTerminalKeys(opts *bind.CallOpts, _nodeOperator common.Address, startIndex *big.Int, endIndex *big.Int) (uint64, error) {
	var out []interface{}
	err := _PermissionlessNodeRegistry.contract.Call(opts, &out, "getOperatorTotalNonTerminalKeys", _nodeOperator, startIndex, endIndex)

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// GetOperatorTotalNonTerminalKeys is a free data retrieval call binding the contract method 0x8a25bcec.
//
// Solidity: function getOperatorTotalNonTerminalKeys(address _nodeOperator, uint256 startIndex, uint256 endIndex) view returns(uint64)
func (_PermissionlessNodeRegistry *PermissionlessNodeRegistrySession) GetOperatorTotalNonTerminalKeys(_nodeOperator common.Address, startIndex *big.Int, endIndex *big.Int) (uint64, error) {
	return _PermissionlessNodeRegistry.Contract.GetOperatorTotalNonTerminalKeys(&_PermissionlessNodeRegistry.CallOpts, _nodeOperator, startIndex, endIndex)
}

// GetOperatorTotalNonTerminalKeys is a free data retrieval call binding the contract method 0x8a25bcec.
//
// Solidity: function getOperatorTotalNonTerminalKeys(address _nodeOperator, uint256 startIndex, uint256 endIndex) view returns(uint64)
func (_PermissionlessNodeRegistry *PermissionlessNodeRegistryCallerSession) GetOperatorTotalNonTerminalKeys(_nodeOperator common.Address, startIndex *big.Int, endIndex *big.Int) (uint64, error) {
	return _PermissionlessNodeRegistry.Contract.GetOperatorTotalNonTerminalKeys(&_PermissionlessNodeRegistry.CallOpts, _nodeOperator, startIndex, endIndex)
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

// GetSocializingPoolStateChangeTimestamp is a free data retrieval call binding the contract method 0x3520862a.
//
// Solidity: function getSocializingPoolStateChangeTimestamp(uint256 _operatorId) view returns(uint256)
func (_PermissionlessNodeRegistry *PermissionlessNodeRegistryCaller) GetSocializingPoolStateChangeTimestamp(opts *bind.CallOpts, _operatorId *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _PermissionlessNodeRegistry.contract.Call(opts, &out, "getSocializingPoolStateChangeTimestamp", _operatorId)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetSocializingPoolStateChangeTimestamp is a free data retrieval call binding the contract method 0x3520862a.
//
// Solidity: function getSocializingPoolStateChangeTimestamp(uint256 _operatorId) view returns(uint256)
func (_PermissionlessNodeRegistry *PermissionlessNodeRegistrySession) GetSocializingPoolStateChangeTimestamp(_operatorId *big.Int) (*big.Int, error) {
	return _PermissionlessNodeRegistry.Contract.GetSocializingPoolStateChangeTimestamp(&_PermissionlessNodeRegistry.CallOpts, _operatorId)
}

// GetSocializingPoolStateChangeTimestamp is a free data retrieval call binding the contract method 0x3520862a.
//
// Solidity: function getSocializingPoolStateChangeTimestamp(uint256 _operatorId) view returns(uint256)
func (_PermissionlessNodeRegistry *PermissionlessNodeRegistryCallerSession) GetSocializingPoolStateChangeTimestamp(_operatorId *big.Int) (*big.Int, error) {
	return _PermissionlessNodeRegistry.Contract.GetSocializingPoolStateChangeTimestamp(&_PermissionlessNodeRegistry.CallOpts, _operatorId)
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

// GetValidator is a free data retrieval call binding the contract method 0xb4891bfd.
//
// Solidity: function getValidator(bytes _pubkey) view returns((uint8,bytes,bytes,bytes,address,uint256,uint256,uint256,uint256))
func (_PermissionlessNodeRegistry *PermissionlessNodeRegistryCaller) GetValidator(opts *bind.CallOpts, _pubkey []byte) (Validator, error) {
	var out []interface{}
	err := _PermissionlessNodeRegistry.contract.Call(opts, &out, "getValidator", _pubkey)

	if err != nil {
		return *new(Validator), err
	}

	out0 := *abi.ConvertType(out[0], new(Validator)).(*Validator)

	return out0, err

}

// GetValidator is a free data retrieval call binding the contract method 0xb4891bfd.
//
// Solidity: function getValidator(bytes _pubkey) view returns((uint8,bytes,bytes,bytes,address,uint256,uint256,uint256,uint256))
func (_PermissionlessNodeRegistry *PermissionlessNodeRegistrySession) GetValidator(_pubkey []byte) (Validator, error) {
	return _PermissionlessNodeRegistry.Contract.GetValidator(&_PermissionlessNodeRegistry.CallOpts, _pubkey)
}

// GetValidator is a free data retrieval call binding the contract method 0xb4891bfd.
//
// Solidity: function getValidator(bytes _pubkey) view returns((uint8,bytes,bytes,bytes,address,uint256,uint256,uint256,uint256))
func (_PermissionlessNodeRegistry *PermissionlessNodeRegistryCallerSession) GetValidator(_pubkey []byte) (Validator, error) {
	return _PermissionlessNodeRegistry.Contract.GetValidator(&_PermissionlessNodeRegistry.CallOpts, _pubkey)
}

// GetValidator0 is a free data retrieval call binding the contract method 0xb5d89627.
//
// Solidity: function getValidator(uint256 _validatorId) view returns((uint8,bytes,bytes,bytes,address,uint256,uint256,uint256,uint256))
func (_PermissionlessNodeRegistry *PermissionlessNodeRegistryCaller) GetValidator0(opts *bind.CallOpts, _validatorId *big.Int) (Validator, error) {
	var out []interface{}
	err := _PermissionlessNodeRegistry.contract.Call(opts, &out, "getValidator0", _validatorId)

	if err != nil {
		return *new(Validator), err
	}

	out0 := *abi.ConvertType(out[0], new(Validator)).(*Validator)

	return out0, err

}

// GetValidator0 is a free data retrieval call binding the contract method 0xb5d89627.
//
// Solidity: function getValidator(uint256 _validatorId) view returns((uint8,bytes,bytes,bytes,address,uint256,uint256,uint256,uint256))
func (_PermissionlessNodeRegistry *PermissionlessNodeRegistrySession) GetValidator0(_validatorId *big.Int) (Validator, error) {
	return _PermissionlessNodeRegistry.Contract.GetValidator0(&_PermissionlessNodeRegistry.CallOpts, _validatorId)
}

// GetValidator0 is a free data retrieval call binding the contract method 0xb5d89627.
//
// Solidity: function getValidator(uint256 _validatorId) view returns((uint8,bytes,bytes,bytes,address,uint256,uint256,uint256,uint256))
func (_PermissionlessNodeRegistry *PermissionlessNodeRegistryCallerSession) GetValidator0(_validatorId *big.Int) (Validator, error) {
	return _PermissionlessNodeRegistry.Contract.GetValidator0(&_PermissionlessNodeRegistry.CallOpts, _validatorId)
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

// MaxKeyPerOperator is a free data retrieval call binding the contract method 0xdf606985.
//
// Solidity: function maxKeyPerOperator() view returns(uint64)
func (_PermissionlessNodeRegistry *PermissionlessNodeRegistryCaller) MaxKeyPerOperator(opts *bind.CallOpts) (uint64, error) {
	var out []interface{}
	err := _PermissionlessNodeRegistry.contract.Call(opts, &out, "maxKeyPerOperator")

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// MaxKeyPerOperator is a free data retrieval call binding the contract method 0xdf606985.
//
// Solidity: function maxKeyPerOperator() view returns(uint64)
func (_PermissionlessNodeRegistry *PermissionlessNodeRegistrySession) MaxKeyPerOperator() (uint64, error) {
	return _PermissionlessNodeRegistry.Contract.MaxKeyPerOperator(&_PermissionlessNodeRegistry.CallOpts)
}

// MaxKeyPerOperator is a free data retrieval call binding the contract method 0xdf606985.
//
// Solidity: function maxKeyPerOperator() view returns(uint64)
func (_PermissionlessNodeRegistry *PermissionlessNodeRegistryCallerSession) MaxKeyPerOperator() (uint64, error) {
	return _PermissionlessNodeRegistry.Contract.MaxKeyPerOperator(&_PermissionlessNodeRegistry.CallOpts)
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

// PermissionlessPool is a free data retrieval call binding the contract method 0xccf4da22.
//
// Solidity: function permissionlessPool() view returns(address)
func (_PermissionlessNodeRegistry *PermissionlessNodeRegistryCaller) PermissionlessPool(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _PermissionlessNodeRegistry.contract.Call(opts, &out, "permissionlessPool")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// PermissionlessPool is a free data retrieval call binding the contract method 0xccf4da22.
//
// Solidity: function permissionlessPool() view returns(address)
func (_PermissionlessNodeRegistry *PermissionlessNodeRegistrySession) PermissionlessPool() (common.Address, error) {
	return _PermissionlessNodeRegistry.Contract.PermissionlessPool(&_PermissionlessNodeRegistry.CallOpts)
}

// PermissionlessPool is a free data retrieval call binding the contract method 0xccf4da22.
//
// Solidity: function permissionlessPool() view returns(address)
func (_PermissionlessNodeRegistry *PermissionlessNodeRegistryCallerSession) PermissionlessPool() (common.Address, error) {
	return _PermissionlessNodeRegistry.Contract.PermissionlessPool(&_PermissionlessNodeRegistry.CallOpts)
}

// PoolFactoryAddress is a free data retrieval call binding the contract method 0xa32cf536.
//
// Solidity: function poolFactoryAddress() view returns(address)
func (_PermissionlessNodeRegistry *PermissionlessNodeRegistryCaller) PoolFactoryAddress(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _PermissionlessNodeRegistry.contract.Call(opts, &out, "poolFactoryAddress")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// PoolFactoryAddress is a free data retrieval call binding the contract method 0xa32cf536.
//
// Solidity: function poolFactoryAddress() view returns(address)
func (_PermissionlessNodeRegistry *PermissionlessNodeRegistrySession) PoolFactoryAddress() (common.Address, error) {
	return _PermissionlessNodeRegistry.Contract.PoolFactoryAddress(&_PermissionlessNodeRegistry.CallOpts)
}

// PoolFactoryAddress is a free data retrieval call binding the contract method 0xa32cf536.
//
// Solidity: function poolFactoryAddress() view returns(address)
func (_PermissionlessNodeRegistry *PermissionlessNodeRegistryCallerSession) PoolFactoryAddress() (common.Address, error) {
	return _PermissionlessNodeRegistry.Contract.PoolFactoryAddress(&_PermissionlessNodeRegistry.CallOpts)
}

// PoolId is a free data retrieval call binding the contract method 0x3e0dc34e.
//
// Solidity: function poolId() view returns(uint8)
func (_PermissionlessNodeRegistry *PermissionlessNodeRegistryCaller) PoolId(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _PermissionlessNodeRegistry.contract.Call(opts, &out, "poolId")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// PoolId is a free data retrieval call binding the contract method 0x3e0dc34e.
//
// Solidity: function poolId() view returns(uint8)
func (_PermissionlessNodeRegistry *PermissionlessNodeRegistrySession) PoolId() (uint8, error) {
	return _PermissionlessNodeRegistry.Contract.PoolId(&_PermissionlessNodeRegistry.CallOpts)
}

// PoolId is a free data retrieval call binding the contract method 0x3e0dc34e.
//
// Solidity: function poolId() view returns(uint8)
func (_PermissionlessNodeRegistry *PermissionlessNodeRegistryCallerSession) PoolId() (uint8, error) {
	return _PermissionlessNodeRegistry.Contract.PoolId(&_PermissionlessNodeRegistry.CallOpts)
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

// SdCollateral is a free data retrieval call binding the contract method 0xe048b2da.
//
// Solidity: function sdCollateral() view returns(address)
func (_PermissionlessNodeRegistry *PermissionlessNodeRegistryCaller) SdCollateral(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _PermissionlessNodeRegistry.contract.Call(opts, &out, "sdCollateral")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// SdCollateral is a free data retrieval call binding the contract method 0xe048b2da.
//
// Solidity: function sdCollateral() view returns(address)
func (_PermissionlessNodeRegistry *PermissionlessNodeRegistrySession) SdCollateral() (common.Address, error) {
	return _PermissionlessNodeRegistry.Contract.SdCollateral(&_PermissionlessNodeRegistry.CallOpts)
}

// SdCollateral is a free data retrieval call binding the contract method 0xe048b2da.
//
// Solidity: function sdCollateral() view returns(address)
func (_PermissionlessNodeRegistry *PermissionlessNodeRegistryCallerSession) SdCollateral() (common.Address, error) {
	return _PermissionlessNodeRegistry.Contract.SdCollateral(&_PermissionlessNodeRegistry.CallOpts)
}

// SocializePoolRewardDistributionCycle is a free data retrieval call binding the contract method 0xd2986f39.
//
// Solidity: function socializePoolRewardDistributionCycle() view returns(uint256)
func (_PermissionlessNodeRegistry *PermissionlessNodeRegistryCaller) SocializePoolRewardDistributionCycle(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _PermissionlessNodeRegistry.contract.Call(opts, &out, "socializePoolRewardDistributionCycle")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// SocializePoolRewardDistributionCycle is a free data retrieval call binding the contract method 0xd2986f39.
//
// Solidity: function socializePoolRewardDistributionCycle() view returns(uint256)
func (_PermissionlessNodeRegistry *PermissionlessNodeRegistrySession) SocializePoolRewardDistributionCycle() (*big.Int, error) {
	return _PermissionlessNodeRegistry.Contract.SocializePoolRewardDistributionCycle(&_PermissionlessNodeRegistry.CallOpts)
}

// SocializePoolRewardDistributionCycle is a free data retrieval call binding the contract method 0xd2986f39.
//
// Solidity: function socializePoolRewardDistributionCycle() view returns(uint256)
func (_PermissionlessNodeRegistry *PermissionlessNodeRegistryCallerSession) SocializePoolRewardDistributionCycle() (*big.Int, error) {
	return _PermissionlessNodeRegistry.Contract.SocializePoolRewardDistributionCycle(&_PermissionlessNodeRegistry.CallOpts)
}

// SocializingPoolStateChangeTimestamp is a free data retrieval call binding the contract method 0x33f807de.
//
// Solidity: function socializingPoolStateChangeTimestamp(uint256 ) view returns(uint256)
func (_PermissionlessNodeRegistry *PermissionlessNodeRegistryCaller) SocializingPoolStateChangeTimestamp(opts *bind.CallOpts, arg0 *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _PermissionlessNodeRegistry.contract.Call(opts, &out, "socializingPoolStateChangeTimestamp", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// SocializingPoolStateChangeTimestamp is a free data retrieval call binding the contract method 0x33f807de.
//
// Solidity: function socializingPoolStateChangeTimestamp(uint256 ) view returns(uint256)
func (_PermissionlessNodeRegistry *PermissionlessNodeRegistrySession) SocializingPoolStateChangeTimestamp(arg0 *big.Int) (*big.Int, error) {
	return _PermissionlessNodeRegistry.Contract.SocializingPoolStateChangeTimestamp(&_PermissionlessNodeRegistry.CallOpts, arg0)
}

// SocializingPoolStateChangeTimestamp is a free data retrieval call binding the contract method 0x33f807de.
//
// Solidity: function socializingPoolStateChangeTimestamp(uint256 ) view returns(uint256)
func (_PermissionlessNodeRegistry *PermissionlessNodeRegistryCallerSession) SocializingPoolStateChangeTimestamp(arg0 *big.Int) (*big.Int, error) {
	return _PermissionlessNodeRegistry.Contract.SocializingPoolStateChangeTimestamp(&_PermissionlessNodeRegistry.CallOpts, arg0)
}

// StaderPenaltyFund is a free data retrieval call binding the contract method 0x8d69b872.
//
// Solidity: function staderPenaltyFund() view returns(address)
func (_PermissionlessNodeRegistry *PermissionlessNodeRegistryCaller) StaderPenaltyFund(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _PermissionlessNodeRegistry.contract.Call(opts, &out, "staderPenaltyFund")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// StaderPenaltyFund is a free data retrieval call binding the contract method 0x8d69b872.
//
// Solidity: function staderPenaltyFund() view returns(address)
func (_PermissionlessNodeRegistry *PermissionlessNodeRegistrySession) StaderPenaltyFund() (common.Address, error) {
	return _PermissionlessNodeRegistry.Contract.StaderPenaltyFund(&_PermissionlessNodeRegistry.CallOpts)
}

// StaderPenaltyFund is a free data retrieval call binding the contract method 0x8d69b872.
//
// Solidity: function staderPenaltyFund() view returns(address)
func (_PermissionlessNodeRegistry *PermissionlessNodeRegistryCallerSession) StaderPenaltyFund() (common.Address, error) {
	return _PermissionlessNodeRegistry.Contract.StaderPenaltyFund(&_PermissionlessNodeRegistry.CallOpts)
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
// Solidity: function validatorRegistry(uint256 ) view returns(uint8 status, bytes pubkey, bytes preDepositSignature, bytes depositSignature, address withdrawVaultAddress, uint256 operatorId, uint256 initialBondEth, uint256 depositTime, uint256 withdrawnTime)
func (_PermissionlessNodeRegistry *PermissionlessNodeRegistryCaller) ValidatorRegistry(opts *bind.CallOpts, arg0 *big.Int) (struct {
	Status               uint8
	Pubkey               []byte
	PreDepositSignature  []byte
	DepositSignature     []byte
	WithdrawVaultAddress common.Address
	OperatorId           *big.Int
	InitialBondEth       *big.Int
	DepositTime          *big.Int
	WithdrawnTime        *big.Int
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
		InitialBondEth       *big.Int
		DepositTime          *big.Int
		WithdrawnTime        *big.Int
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
	outstruct.InitialBondEth = *abi.ConvertType(out[6], new(*big.Int)).(**big.Int)
	outstruct.DepositTime = *abi.ConvertType(out[7], new(*big.Int)).(**big.Int)
	outstruct.WithdrawnTime = *abi.ConvertType(out[8], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// ValidatorRegistry is a free data retrieval call binding the contract method 0x5a1239c1.
//
// Solidity: function validatorRegistry(uint256 ) view returns(uint8 status, bytes pubkey, bytes preDepositSignature, bytes depositSignature, address withdrawVaultAddress, uint256 operatorId, uint256 initialBondEth, uint256 depositTime, uint256 withdrawnTime)
func (_PermissionlessNodeRegistry *PermissionlessNodeRegistrySession) ValidatorRegistry(arg0 *big.Int) (struct {
	Status               uint8
	Pubkey               []byte
	PreDepositSignature  []byte
	DepositSignature     []byte
	WithdrawVaultAddress common.Address
	OperatorId           *big.Int
	InitialBondEth       *big.Int
	DepositTime          *big.Int
	WithdrawnTime        *big.Int
}, error) {
	return _PermissionlessNodeRegistry.Contract.ValidatorRegistry(&_PermissionlessNodeRegistry.CallOpts, arg0)
}

// ValidatorRegistry is a free data retrieval call binding the contract method 0x5a1239c1.
//
// Solidity: function validatorRegistry(uint256 ) view returns(uint8 status, bytes pubkey, bytes preDepositSignature, bytes depositSignature, address withdrawVaultAddress, uint256 operatorId, uint256 initialBondEth, uint256 depositTime, uint256 withdrawnTime)
func (_PermissionlessNodeRegistry *PermissionlessNodeRegistryCallerSession) ValidatorRegistry(arg0 *big.Int) (struct {
	Status               uint8
	Pubkey               []byte
	PreDepositSignature  []byte
	DepositSignature     []byte
	WithdrawVaultAddress common.Address
	OperatorId           *big.Int
	InitialBondEth       *big.Int
	DepositTime          *big.Int
	WithdrawnTime        *big.Int
}, error) {
	return _PermissionlessNodeRegistry.Contract.ValidatorRegistry(&_PermissionlessNodeRegistry.CallOpts, arg0)
}

// VaultFactoryAddress is a free data retrieval call binding the contract method 0x42089791.
//
// Solidity: function vaultFactoryAddress() view returns(address)
func (_PermissionlessNodeRegistry *PermissionlessNodeRegistryCaller) VaultFactoryAddress(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _PermissionlessNodeRegistry.contract.Call(opts, &out, "vaultFactoryAddress")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// VaultFactoryAddress is a free data retrieval call binding the contract method 0x42089791.
//
// Solidity: function vaultFactoryAddress() view returns(address)
func (_PermissionlessNodeRegistry *PermissionlessNodeRegistrySession) VaultFactoryAddress() (common.Address, error) {
	return _PermissionlessNodeRegistry.Contract.VaultFactoryAddress(&_PermissionlessNodeRegistry.CallOpts)
}

// VaultFactoryAddress is a free data retrieval call binding the contract method 0x42089791.
//
// Solidity: function vaultFactoryAddress() view returns(address)
func (_PermissionlessNodeRegistry *PermissionlessNodeRegistryCallerSession) VaultFactoryAddress() (common.Address, error) {
	return _PermissionlessNodeRegistry.Contract.VaultFactoryAddress(&_PermissionlessNodeRegistry.CallOpts)
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

// Initialize is a paid mutator transaction binding the contract method 0xcc2a9a5b.
//
// Solidity: function initialize(address _adminOwner, address _sdCollateral, address _staderPenaltyFund, address _vaultFactoryAddress, address _elRewardSocializePool, address _poolFactoryAddress) returns()
func (_PermissionlessNodeRegistry *PermissionlessNodeRegistryTransactor) Initialize(opts *bind.TransactOpts, _adminOwner common.Address, _sdCollateral common.Address, _staderPenaltyFund common.Address, _vaultFactoryAddress common.Address, _elRewardSocializePool common.Address, _poolFactoryAddress common.Address) (*types.Transaction, error) {
	return _PermissionlessNodeRegistry.contract.Transact(opts, "initialize", _adminOwner, _sdCollateral, _staderPenaltyFund, _vaultFactoryAddress, _elRewardSocializePool, _poolFactoryAddress)
}

// Initialize is a paid mutator transaction binding the contract method 0xcc2a9a5b.
//
// Solidity: function initialize(address _adminOwner, address _sdCollateral, address _staderPenaltyFund, address _vaultFactoryAddress, address _elRewardSocializePool, address _poolFactoryAddress) returns()
func (_PermissionlessNodeRegistry *PermissionlessNodeRegistrySession) Initialize(_adminOwner common.Address, _sdCollateral common.Address, _staderPenaltyFund common.Address, _vaultFactoryAddress common.Address, _elRewardSocializePool common.Address, _poolFactoryAddress common.Address) (*types.Transaction, error) {
	return _PermissionlessNodeRegistry.Contract.Initialize(&_PermissionlessNodeRegistry.TransactOpts, _adminOwner, _sdCollateral, _staderPenaltyFund, _vaultFactoryAddress, _elRewardSocializePool, _poolFactoryAddress)
}

// Initialize is a paid mutator transaction binding the contract method 0xcc2a9a5b.
//
// Solidity: function initialize(address _adminOwner, address _sdCollateral, address _staderPenaltyFund, address _vaultFactoryAddress, address _elRewardSocializePool, address _poolFactoryAddress) returns()
func (_PermissionlessNodeRegistry *PermissionlessNodeRegistryTransactorSession) Initialize(_adminOwner common.Address, _sdCollateral common.Address, _staderPenaltyFund common.Address, _vaultFactoryAddress common.Address, _elRewardSocializePool common.Address, _poolFactoryAddress common.Address) (*types.Transaction, error) {
	return _PermissionlessNodeRegistry.Contract.Initialize(&_PermissionlessNodeRegistry.TransactOpts, _adminOwner, _sdCollateral, _staderPenaltyFund, _vaultFactoryAddress, _elRewardSocializePool, _poolFactoryAddress)
}

// MarkValidatorReadyToDeposit is a paid mutator transaction binding the contract method 0x13797bff.
//
// Solidity: function markValidatorReadyToDeposit(bytes[] _readyToDepositPubkey, bytes[] _frontRunnedPubkey, bytes[] _invalidSignaturePubkey) returns()
func (_PermissionlessNodeRegistry *PermissionlessNodeRegistryTransactor) MarkValidatorReadyToDeposit(opts *bind.TransactOpts, _readyToDepositPubkey [][]byte, _frontRunnedPubkey [][]byte, _invalidSignaturePubkey [][]byte) (*types.Transaction, error) {
	return _PermissionlessNodeRegistry.contract.Transact(opts, "markValidatorReadyToDeposit", _readyToDepositPubkey, _frontRunnedPubkey, _invalidSignaturePubkey)
}

// MarkValidatorReadyToDeposit is a paid mutator transaction binding the contract method 0x13797bff.
//
// Solidity: function markValidatorReadyToDeposit(bytes[] _readyToDepositPubkey, bytes[] _frontRunnedPubkey, bytes[] _invalidSignaturePubkey) returns()
func (_PermissionlessNodeRegistry *PermissionlessNodeRegistrySession) MarkValidatorReadyToDeposit(_readyToDepositPubkey [][]byte, _frontRunnedPubkey [][]byte, _invalidSignaturePubkey [][]byte) (*types.Transaction, error) {
	return _PermissionlessNodeRegistry.Contract.MarkValidatorReadyToDeposit(&_PermissionlessNodeRegistry.TransactOpts, _readyToDepositPubkey, _frontRunnedPubkey, _invalidSignaturePubkey)
}

// MarkValidatorReadyToDeposit is a paid mutator transaction binding the contract method 0x13797bff.
//
// Solidity: function markValidatorReadyToDeposit(bytes[] _readyToDepositPubkey, bytes[] _frontRunnedPubkey, bytes[] _invalidSignaturePubkey) returns()
func (_PermissionlessNodeRegistry *PermissionlessNodeRegistryTransactorSession) MarkValidatorReadyToDeposit(_readyToDepositPubkey [][]byte, _frontRunnedPubkey [][]byte, _invalidSignaturePubkey [][]byte) (*types.Transaction, error) {
	return _PermissionlessNodeRegistry.Contract.MarkValidatorReadyToDeposit(&_PermissionlessNodeRegistry.TransactOpts, _readyToDepositPubkey, _frontRunnedPubkey, _invalidSignaturePubkey)
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

// UpdateDepositStatusAndTime is a paid mutator transaction binding the contract method 0x26d70e98.
//
// Solidity: function updateDepositStatusAndTime(uint256 _validatorId) returns()
func (_PermissionlessNodeRegistry *PermissionlessNodeRegistryTransactor) UpdateDepositStatusAndTime(opts *bind.TransactOpts, _validatorId *big.Int) (*types.Transaction, error) {
	return _PermissionlessNodeRegistry.contract.Transact(opts, "updateDepositStatusAndTime", _validatorId)
}

// UpdateDepositStatusAndTime is a paid mutator transaction binding the contract method 0x26d70e98.
//
// Solidity: function updateDepositStatusAndTime(uint256 _validatorId) returns()
func (_PermissionlessNodeRegistry *PermissionlessNodeRegistrySession) UpdateDepositStatusAndTime(_validatorId *big.Int) (*types.Transaction, error) {
	return _PermissionlessNodeRegistry.Contract.UpdateDepositStatusAndTime(&_PermissionlessNodeRegistry.TransactOpts, _validatorId)
}

// UpdateDepositStatusAndTime is a paid mutator transaction binding the contract method 0x26d70e98.
//
// Solidity: function updateDepositStatusAndTime(uint256 _validatorId) returns()
func (_PermissionlessNodeRegistry *PermissionlessNodeRegistryTransactorSession) UpdateDepositStatusAndTime(_validatorId *big.Int) (*types.Transaction, error) {
	return _PermissionlessNodeRegistry.Contract.UpdateDepositStatusAndTime(&_PermissionlessNodeRegistry.TransactOpts, _validatorId)
}

// UpdateELRewardSocializePool is a paid mutator transaction binding the contract method 0x14382841.
//
// Solidity: function updateELRewardSocializePool(address _elRewardSocializePool) returns()
func (_PermissionlessNodeRegistry *PermissionlessNodeRegistryTransactor) UpdateELRewardSocializePool(opts *bind.TransactOpts, _elRewardSocializePool common.Address) (*types.Transaction, error) {
	return _PermissionlessNodeRegistry.contract.Transact(opts, "updateELRewardSocializePool", _elRewardSocializePool)
}

// UpdateELRewardSocializePool is a paid mutator transaction binding the contract method 0x14382841.
//
// Solidity: function updateELRewardSocializePool(address _elRewardSocializePool) returns()
func (_PermissionlessNodeRegistry *PermissionlessNodeRegistrySession) UpdateELRewardSocializePool(_elRewardSocializePool common.Address) (*types.Transaction, error) {
	return _PermissionlessNodeRegistry.Contract.UpdateELRewardSocializePool(&_PermissionlessNodeRegistry.TransactOpts, _elRewardSocializePool)
}

// UpdateELRewardSocializePool is a paid mutator transaction binding the contract method 0x14382841.
//
// Solidity: function updateELRewardSocializePool(address _elRewardSocializePool) returns()
func (_PermissionlessNodeRegistry *PermissionlessNodeRegistryTransactorSession) UpdateELRewardSocializePool(_elRewardSocializePool common.Address) (*types.Transaction, error) {
	return _PermissionlessNodeRegistry.Contract.UpdateELRewardSocializePool(&_PermissionlessNodeRegistry.TransactOpts, _elRewardSocializePool)
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

// UpdateMaxKeyPerOperator is a paid mutator transaction binding the contract method 0x25cfe633.
//
// Solidity: function updateMaxKeyPerOperator(uint64 _maxKeyPerOperator) returns()
func (_PermissionlessNodeRegistry *PermissionlessNodeRegistryTransactor) UpdateMaxKeyPerOperator(opts *bind.TransactOpts, _maxKeyPerOperator uint64) (*types.Transaction, error) {
	return _PermissionlessNodeRegistry.contract.Transact(opts, "updateMaxKeyPerOperator", _maxKeyPerOperator)
}

// UpdateMaxKeyPerOperator is a paid mutator transaction binding the contract method 0x25cfe633.
//
// Solidity: function updateMaxKeyPerOperator(uint64 _maxKeyPerOperator) returns()
func (_PermissionlessNodeRegistry *PermissionlessNodeRegistrySession) UpdateMaxKeyPerOperator(_maxKeyPerOperator uint64) (*types.Transaction, error) {
	return _PermissionlessNodeRegistry.Contract.UpdateMaxKeyPerOperator(&_PermissionlessNodeRegistry.TransactOpts, _maxKeyPerOperator)
}

// UpdateMaxKeyPerOperator is a paid mutator transaction binding the contract method 0x25cfe633.
//
// Solidity: function updateMaxKeyPerOperator(uint64 _maxKeyPerOperator) returns()
func (_PermissionlessNodeRegistry *PermissionlessNodeRegistryTransactorSession) UpdateMaxKeyPerOperator(_maxKeyPerOperator uint64) (*types.Transaction, error) {
	return _PermissionlessNodeRegistry.Contract.UpdateMaxKeyPerOperator(&_PermissionlessNodeRegistry.TransactOpts, _maxKeyPerOperator)
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

// UpdatePermissionlessPoolAddress is a paid mutator transaction binding the contract method 0x4ca6d6da.
//
// Solidity: function updatePermissionlessPoolAddress(address _permissionlessPool) returns()
func (_PermissionlessNodeRegistry *PermissionlessNodeRegistryTransactor) UpdatePermissionlessPoolAddress(opts *bind.TransactOpts, _permissionlessPool common.Address) (*types.Transaction, error) {
	return _PermissionlessNodeRegistry.contract.Transact(opts, "updatePermissionlessPoolAddress", _permissionlessPool)
}

// UpdatePermissionlessPoolAddress is a paid mutator transaction binding the contract method 0x4ca6d6da.
//
// Solidity: function updatePermissionlessPoolAddress(address _permissionlessPool) returns()
func (_PermissionlessNodeRegistry *PermissionlessNodeRegistrySession) UpdatePermissionlessPoolAddress(_permissionlessPool common.Address) (*types.Transaction, error) {
	return _PermissionlessNodeRegistry.Contract.UpdatePermissionlessPoolAddress(&_PermissionlessNodeRegistry.TransactOpts, _permissionlessPool)
}

// UpdatePermissionlessPoolAddress is a paid mutator transaction binding the contract method 0x4ca6d6da.
//
// Solidity: function updatePermissionlessPoolAddress(address _permissionlessPool) returns()
func (_PermissionlessNodeRegistry *PermissionlessNodeRegistryTransactorSession) UpdatePermissionlessPoolAddress(_permissionlessPool common.Address) (*types.Transaction, error) {
	return _PermissionlessNodeRegistry.Contract.UpdatePermissionlessPoolAddress(&_PermissionlessNodeRegistry.TransactOpts, _permissionlessPool)
}

// UpdatePoolFactoryAddress is a paid mutator transaction binding the contract method 0x116cffef.
//
// Solidity: function updatePoolFactoryAddress(address _poolFactoryAddress) returns()
func (_PermissionlessNodeRegistry *PermissionlessNodeRegistryTransactor) UpdatePoolFactoryAddress(opts *bind.TransactOpts, _poolFactoryAddress common.Address) (*types.Transaction, error) {
	return _PermissionlessNodeRegistry.contract.Transact(opts, "updatePoolFactoryAddress", _poolFactoryAddress)
}

// UpdatePoolFactoryAddress is a paid mutator transaction binding the contract method 0x116cffef.
//
// Solidity: function updatePoolFactoryAddress(address _poolFactoryAddress) returns()
func (_PermissionlessNodeRegistry *PermissionlessNodeRegistrySession) UpdatePoolFactoryAddress(_poolFactoryAddress common.Address) (*types.Transaction, error) {
	return _PermissionlessNodeRegistry.Contract.UpdatePoolFactoryAddress(&_PermissionlessNodeRegistry.TransactOpts, _poolFactoryAddress)
}

// UpdatePoolFactoryAddress is a paid mutator transaction binding the contract method 0x116cffef.
//
// Solidity: function updatePoolFactoryAddress(address _poolFactoryAddress) returns()
func (_PermissionlessNodeRegistry *PermissionlessNodeRegistryTransactorSession) UpdatePoolFactoryAddress(_poolFactoryAddress common.Address) (*types.Transaction, error) {
	return _PermissionlessNodeRegistry.Contract.UpdatePoolFactoryAddress(&_PermissionlessNodeRegistry.TransactOpts, _poolFactoryAddress)
}

// UpdateSDCollateralAddress is a paid mutator transaction binding the contract method 0x94c23cc7.
//
// Solidity: function updateSDCollateralAddress(address _sdCollateral) returns()
func (_PermissionlessNodeRegistry *PermissionlessNodeRegistryTransactor) UpdateSDCollateralAddress(opts *bind.TransactOpts, _sdCollateral common.Address) (*types.Transaction, error) {
	return _PermissionlessNodeRegistry.contract.Transact(opts, "updateSDCollateralAddress", _sdCollateral)
}

// UpdateSDCollateralAddress is a paid mutator transaction binding the contract method 0x94c23cc7.
//
// Solidity: function updateSDCollateralAddress(address _sdCollateral) returns()
func (_PermissionlessNodeRegistry *PermissionlessNodeRegistrySession) UpdateSDCollateralAddress(_sdCollateral common.Address) (*types.Transaction, error) {
	return _PermissionlessNodeRegistry.Contract.UpdateSDCollateralAddress(&_PermissionlessNodeRegistry.TransactOpts, _sdCollateral)
}

// UpdateSDCollateralAddress is a paid mutator transaction binding the contract method 0x94c23cc7.
//
// Solidity: function updateSDCollateralAddress(address _sdCollateral) returns()
func (_PermissionlessNodeRegistry *PermissionlessNodeRegistryTransactorSession) UpdateSDCollateralAddress(_sdCollateral common.Address) (*types.Transaction, error) {
	return _PermissionlessNodeRegistry.Contract.UpdateSDCollateralAddress(&_PermissionlessNodeRegistry.TransactOpts, _sdCollateral)
}

// UpdateStaderPenaltyFundAddress is a paid mutator transaction binding the contract method 0xc7c7b304.
//
// Solidity: function updateStaderPenaltyFundAddress(address _staderPenaltyFund) returns()
func (_PermissionlessNodeRegistry *PermissionlessNodeRegistryTransactor) UpdateStaderPenaltyFundAddress(opts *bind.TransactOpts, _staderPenaltyFund common.Address) (*types.Transaction, error) {
	return _PermissionlessNodeRegistry.contract.Transact(opts, "updateStaderPenaltyFundAddress", _staderPenaltyFund)
}

// UpdateStaderPenaltyFundAddress is a paid mutator transaction binding the contract method 0xc7c7b304.
//
// Solidity: function updateStaderPenaltyFundAddress(address _staderPenaltyFund) returns()
func (_PermissionlessNodeRegistry *PermissionlessNodeRegistrySession) UpdateStaderPenaltyFundAddress(_staderPenaltyFund common.Address) (*types.Transaction, error) {
	return _PermissionlessNodeRegistry.Contract.UpdateStaderPenaltyFundAddress(&_PermissionlessNodeRegistry.TransactOpts, _staderPenaltyFund)
}

// UpdateStaderPenaltyFundAddress is a paid mutator transaction binding the contract method 0xc7c7b304.
//
// Solidity: function updateStaderPenaltyFundAddress(address _staderPenaltyFund) returns()
func (_PermissionlessNodeRegistry *PermissionlessNodeRegistryTransactorSession) UpdateStaderPenaltyFundAddress(_staderPenaltyFund common.Address) (*types.Transaction, error) {
	return _PermissionlessNodeRegistry.Contract.UpdateStaderPenaltyFundAddress(&_PermissionlessNodeRegistry.TransactOpts, _staderPenaltyFund)
}

// UpdateValidatorStatus is a paid mutator transaction binding the contract method 0x1c6197e8.
//
// Solidity: function updateValidatorStatus(bytes _pubkey, uint8 _status) returns()
func (_PermissionlessNodeRegistry *PermissionlessNodeRegistryTransactor) UpdateValidatorStatus(opts *bind.TransactOpts, _pubkey []byte, _status uint8) (*types.Transaction, error) {
	return _PermissionlessNodeRegistry.contract.Transact(opts, "updateValidatorStatus", _pubkey, _status)
}

// UpdateValidatorStatus is a paid mutator transaction binding the contract method 0x1c6197e8.
//
// Solidity: function updateValidatorStatus(bytes _pubkey, uint8 _status) returns()
func (_PermissionlessNodeRegistry *PermissionlessNodeRegistrySession) UpdateValidatorStatus(_pubkey []byte, _status uint8) (*types.Transaction, error) {
	return _PermissionlessNodeRegistry.Contract.UpdateValidatorStatus(&_PermissionlessNodeRegistry.TransactOpts, _pubkey, _status)
}

// UpdateValidatorStatus is a paid mutator transaction binding the contract method 0x1c6197e8.
//
// Solidity: function updateValidatorStatus(bytes _pubkey, uint8 _status) returns()
func (_PermissionlessNodeRegistry *PermissionlessNodeRegistryTransactorSession) UpdateValidatorStatus(_pubkey []byte, _status uint8) (*types.Transaction, error) {
	return _PermissionlessNodeRegistry.Contract.UpdateValidatorStatus(&_PermissionlessNodeRegistry.TransactOpts, _pubkey, _status)
}

// UpdateVaultFactoryAddress is a paid mutator transaction binding the contract method 0xf65694e3.
//
// Solidity: function updateVaultFactoryAddress(address _vaultFactoryAddress) returns()
func (_PermissionlessNodeRegistry *PermissionlessNodeRegistryTransactor) UpdateVaultFactoryAddress(opts *bind.TransactOpts, _vaultFactoryAddress common.Address) (*types.Transaction, error) {
	return _PermissionlessNodeRegistry.contract.Transact(opts, "updateVaultFactoryAddress", _vaultFactoryAddress)
}

// UpdateVaultFactoryAddress is a paid mutator transaction binding the contract method 0xf65694e3.
//
// Solidity: function updateVaultFactoryAddress(address _vaultFactoryAddress) returns()
func (_PermissionlessNodeRegistry *PermissionlessNodeRegistrySession) UpdateVaultFactoryAddress(_vaultFactoryAddress common.Address) (*types.Transaction, error) {
	return _PermissionlessNodeRegistry.Contract.UpdateVaultFactoryAddress(&_PermissionlessNodeRegistry.TransactOpts, _vaultFactoryAddress)
}

// UpdateVaultFactoryAddress is a paid mutator transaction binding the contract method 0xf65694e3.
//
// Solidity: function updateVaultFactoryAddress(address _vaultFactoryAddress) returns()
func (_PermissionlessNodeRegistry *PermissionlessNodeRegistryTransactorSession) UpdateVaultFactoryAddress(_vaultFactoryAddress common.Address) (*types.Transaction, error) {
	return _PermissionlessNodeRegistry.Contract.UpdateVaultFactoryAddress(&_PermissionlessNodeRegistry.TransactOpts, _vaultFactoryAddress)
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

// PermissionlessNodeRegistryAddedKeysIterator is returned from FilterAddedKeys and is used to iterate over the raw logs and unpacked data for AddedKeys events raised by the PermissionlessNodeRegistry contract.
type PermissionlessNodeRegistryAddedKeysIterator struct {
	Event *PermissionlessNodeRegistryAddedKeys // Event containing the contract specifics and raw log

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
func (it *PermissionlessNodeRegistryAddedKeysIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PermissionlessNodeRegistryAddedKeys)
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
		it.Event = new(PermissionlessNodeRegistryAddedKeys)
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
func (it *PermissionlessNodeRegistryAddedKeysIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PermissionlessNodeRegistryAddedKeysIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PermissionlessNodeRegistryAddedKeys represents a AddedKeys event raised by the PermissionlessNodeRegistry contract.
type PermissionlessNodeRegistryAddedKeys struct {
	NodeOperator common.Address
	Pubkey       []byte
	ValidatorId  *big.Int
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterAddedKeys is a free log retrieval operation binding the contract event 0x05128c29faed4dff2adba6177ad093bdf0e2f18e49fa884c794b73027b6b213c.
//
// Solidity: event AddedKeys(address indexed _nodeOperator, bytes _pubkey, uint256 _validatorId)
func (_PermissionlessNodeRegistry *PermissionlessNodeRegistryFilterer) FilterAddedKeys(opts *bind.FilterOpts, _nodeOperator []common.Address) (*PermissionlessNodeRegistryAddedKeysIterator, error) {

	var _nodeOperatorRule []interface{}
	for _, _nodeOperatorItem := range _nodeOperator {
		_nodeOperatorRule = append(_nodeOperatorRule, _nodeOperatorItem)
	}

	logs, sub, err := _PermissionlessNodeRegistry.contract.FilterLogs(opts, "AddedKeys", _nodeOperatorRule)
	if err != nil {
		return nil, err
	}
	return &PermissionlessNodeRegistryAddedKeysIterator{contract: _PermissionlessNodeRegistry.contract, event: "AddedKeys", logs: logs, sub: sub}, nil
}

// WatchAddedKeys is a free log subscription operation binding the contract event 0x05128c29faed4dff2adba6177ad093bdf0e2f18e49fa884c794b73027b6b213c.
//
// Solidity: event AddedKeys(address indexed _nodeOperator, bytes _pubkey, uint256 _validatorId)
func (_PermissionlessNodeRegistry *PermissionlessNodeRegistryFilterer) WatchAddedKeys(opts *bind.WatchOpts, sink chan<- *PermissionlessNodeRegistryAddedKeys, _nodeOperator []common.Address) (event.Subscription, error) {

	var _nodeOperatorRule []interface{}
	for _, _nodeOperatorItem := range _nodeOperator {
		_nodeOperatorRule = append(_nodeOperatorRule, _nodeOperatorItem)
	}

	logs, sub, err := _PermissionlessNodeRegistry.contract.WatchLogs(opts, "AddedKeys", _nodeOperatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PermissionlessNodeRegistryAddedKeys)
				if err := _PermissionlessNodeRegistry.contract.UnpackLog(event, "AddedKeys", log); err != nil {
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

// ParseAddedKeys is a log parse operation binding the contract event 0x05128c29faed4dff2adba6177ad093bdf0e2f18e49fa884c794b73027b6b213c.
//
// Solidity: event AddedKeys(address indexed _nodeOperator, bytes _pubkey, uint256 _validatorId)
func (_PermissionlessNodeRegistry *PermissionlessNodeRegistryFilterer) ParseAddedKeys(log types.Log) (*PermissionlessNodeRegistryAddedKeys, error) {
	event := new(PermissionlessNodeRegistryAddedKeys)
	if err := _PermissionlessNodeRegistry.contract.UnpackLog(event, "AddedKeys", log); err != nil {
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
	NodeOperator common.Address
	OperatorId   *big.Int
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterOnboardedOperator is a free log retrieval operation binding the contract event 0x0f4b23f0e8f5ece9f9346a59e6754a462ccd648010097e6cd69d0537b4ca3c08.
//
// Solidity: event OnboardedOperator(address indexed _nodeOperator, uint256 _operatorId)
func (_PermissionlessNodeRegistry *PermissionlessNodeRegistryFilterer) FilterOnboardedOperator(opts *bind.FilterOpts, _nodeOperator []common.Address) (*PermissionlessNodeRegistryOnboardedOperatorIterator, error) {

	var _nodeOperatorRule []interface{}
	for _, _nodeOperatorItem := range _nodeOperator {
		_nodeOperatorRule = append(_nodeOperatorRule, _nodeOperatorItem)
	}

	logs, sub, err := _PermissionlessNodeRegistry.contract.FilterLogs(opts, "OnboardedOperator", _nodeOperatorRule)
	if err != nil {
		return nil, err
	}
	return &PermissionlessNodeRegistryOnboardedOperatorIterator{contract: _PermissionlessNodeRegistry.contract, event: "OnboardedOperator", logs: logs, sub: sub}, nil
}

// WatchOnboardedOperator is a free log subscription operation binding the contract event 0x0f4b23f0e8f5ece9f9346a59e6754a462ccd648010097e6cd69d0537b4ca3c08.
//
// Solidity: event OnboardedOperator(address indexed _nodeOperator, uint256 _operatorId)
func (_PermissionlessNodeRegistry *PermissionlessNodeRegistryFilterer) WatchOnboardedOperator(opts *bind.WatchOpts, sink chan<- *PermissionlessNodeRegistryOnboardedOperator, _nodeOperator []common.Address) (event.Subscription, error) {

	var _nodeOperatorRule []interface{}
	for _, _nodeOperatorItem := range _nodeOperator {
		_nodeOperatorRule = append(_nodeOperatorRule, _nodeOperatorItem)
	}

	logs, sub, err := _PermissionlessNodeRegistry.contract.WatchLogs(opts, "OnboardedOperator", _nodeOperatorRule)
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

// ParseOnboardedOperator is a log parse operation binding the contract event 0x0f4b23f0e8f5ece9f9346a59e6754a462ccd648010097e6cd69d0537b4ca3c08.
//
// Solidity: event OnboardedOperator(address indexed _nodeOperator, uint256 _operatorId)
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

// PermissionlessNodeRegistryUpdatedELRewardSocializePoolIterator is returned from FilterUpdatedELRewardSocializePool and is used to iterate over the raw logs and unpacked data for UpdatedELRewardSocializePool events raised by the PermissionlessNodeRegistry contract.
type PermissionlessNodeRegistryUpdatedELRewardSocializePoolIterator struct {
	Event *PermissionlessNodeRegistryUpdatedELRewardSocializePool // Event containing the contract specifics and raw log

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
func (it *PermissionlessNodeRegistryUpdatedELRewardSocializePoolIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PermissionlessNodeRegistryUpdatedELRewardSocializePool)
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
		it.Event = new(PermissionlessNodeRegistryUpdatedELRewardSocializePool)
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
func (it *PermissionlessNodeRegistryUpdatedELRewardSocializePoolIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PermissionlessNodeRegistryUpdatedELRewardSocializePoolIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PermissionlessNodeRegistryUpdatedELRewardSocializePool represents a UpdatedELRewardSocializePool event raised by the PermissionlessNodeRegistry contract.
type PermissionlessNodeRegistryUpdatedELRewardSocializePool struct {
	ElRewardSocializePool common.Address
	Raw                   types.Log // Blockchain specific contextual infos
}

// FilterUpdatedELRewardSocializePool is a free log retrieval operation binding the contract event 0xd8b67edadb4681f8e02711834e23a92c852e3861c503be003227c01b66635fee.
//
// Solidity: event UpdatedELRewardSocializePool(address _elRewardSocializePool)
func (_PermissionlessNodeRegistry *PermissionlessNodeRegistryFilterer) FilterUpdatedELRewardSocializePool(opts *bind.FilterOpts) (*PermissionlessNodeRegistryUpdatedELRewardSocializePoolIterator, error) {

	logs, sub, err := _PermissionlessNodeRegistry.contract.FilterLogs(opts, "UpdatedELRewardSocializePool")
	if err != nil {
		return nil, err
	}
	return &PermissionlessNodeRegistryUpdatedELRewardSocializePoolIterator{contract: _PermissionlessNodeRegistry.contract, event: "UpdatedELRewardSocializePool", logs: logs, sub: sub}, nil
}

// WatchUpdatedELRewardSocializePool is a free log subscription operation binding the contract event 0xd8b67edadb4681f8e02711834e23a92c852e3861c503be003227c01b66635fee.
//
// Solidity: event UpdatedELRewardSocializePool(address _elRewardSocializePool)
func (_PermissionlessNodeRegistry *PermissionlessNodeRegistryFilterer) WatchUpdatedELRewardSocializePool(opts *bind.WatchOpts, sink chan<- *PermissionlessNodeRegistryUpdatedELRewardSocializePool) (event.Subscription, error) {

	logs, sub, err := _PermissionlessNodeRegistry.contract.WatchLogs(opts, "UpdatedELRewardSocializePool")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PermissionlessNodeRegistryUpdatedELRewardSocializePool)
				if err := _PermissionlessNodeRegistry.contract.UnpackLog(event, "UpdatedELRewardSocializePool", log); err != nil {
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

// ParseUpdatedELRewardSocializePool is a log parse operation binding the contract event 0xd8b67edadb4681f8e02711834e23a92c852e3861c503be003227c01b66635fee.
//
// Solidity: event UpdatedELRewardSocializePool(address _elRewardSocializePool)
func (_PermissionlessNodeRegistry *PermissionlessNodeRegistryFilterer) ParseUpdatedELRewardSocializePool(log types.Log) (*PermissionlessNodeRegistryUpdatedELRewardSocializePool, error) {
	event := new(PermissionlessNodeRegistryUpdatedELRewardSocializePool)
	if err := _PermissionlessNodeRegistry.contract.UnpackLog(event, "UpdatedELRewardSocializePool", log); err != nil {
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
	InputKeyCountLimit uint16
	Raw                types.Log // Blockchain specific contextual infos
}

// FilterUpdatedInputKeyCountLimit is a free log retrieval operation binding the contract event 0x16da2836197a2cac50b2d1b5e0fae9874b59cf21a0b06d92f34bab1cd4e893a1.
//
// Solidity: event UpdatedInputKeyCountLimit(uint16 _inputKeyCountLimit)
func (_PermissionlessNodeRegistry *PermissionlessNodeRegistryFilterer) FilterUpdatedInputKeyCountLimit(opts *bind.FilterOpts) (*PermissionlessNodeRegistryUpdatedInputKeyCountLimitIterator, error) {

	logs, sub, err := _PermissionlessNodeRegistry.contract.FilterLogs(opts, "UpdatedInputKeyCountLimit")
	if err != nil {
		return nil, err
	}
	return &PermissionlessNodeRegistryUpdatedInputKeyCountLimitIterator{contract: _PermissionlessNodeRegistry.contract, event: "UpdatedInputKeyCountLimit", logs: logs, sub: sub}, nil
}

// WatchUpdatedInputKeyCountLimit is a free log subscription operation binding the contract event 0x16da2836197a2cac50b2d1b5e0fae9874b59cf21a0b06d92f34bab1cd4e893a1.
//
// Solidity: event UpdatedInputKeyCountLimit(uint16 _inputKeyCountLimit)
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

// ParseUpdatedInputKeyCountLimit is a log parse operation binding the contract event 0x16da2836197a2cac50b2d1b5e0fae9874b59cf21a0b06d92f34bab1cd4e893a1.
//
// Solidity: event UpdatedInputKeyCountLimit(uint16 _inputKeyCountLimit)
func (_PermissionlessNodeRegistry *PermissionlessNodeRegistryFilterer) ParseUpdatedInputKeyCountLimit(log types.Log) (*PermissionlessNodeRegistryUpdatedInputKeyCountLimit, error) {
	event := new(PermissionlessNodeRegistryUpdatedInputKeyCountLimit)
	if err := _PermissionlessNodeRegistry.contract.UnpackLog(event, "UpdatedInputKeyCountLimit", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PermissionlessNodeRegistryUpdatedMaxKeyPerOperatorIterator is returned from FilterUpdatedMaxKeyPerOperator and is used to iterate over the raw logs and unpacked data for UpdatedMaxKeyPerOperator events raised by the PermissionlessNodeRegistry contract.
type PermissionlessNodeRegistryUpdatedMaxKeyPerOperatorIterator struct {
	Event *PermissionlessNodeRegistryUpdatedMaxKeyPerOperator // Event containing the contract specifics and raw log

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
func (it *PermissionlessNodeRegistryUpdatedMaxKeyPerOperatorIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PermissionlessNodeRegistryUpdatedMaxKeyPerOperator)
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
		it.Event = new(PermissionlessNodeRegistryUpdatedMaxKeyPerOperator)
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
func (it *PermissionlessNodeRegistryUpdatedMaxKeyPerOperatorIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PermissionlessNodeRegistryUpdatedMaxKeyPerOperatorIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PermissionlessNodeRegistryUpdatedMaxKeyPerOperator represents a UpdatedMaxKeyPerOperator event raised by the PermissionlessNodeRegistry contract.
type PermissionlessNodeRegistryUpdatedMaxKeyPerOperator struct {
	KeyDepositLimit uint64
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterUpdatedMaxKeyPerOperator is a free log retrieval operation binding the contract event 0xae32a0fc0865ac1ff229d2eb9e38d025ba79c51aa94e691493815ec2b327ef42.
//
// Solidity: event UpdatedMaxKeyPerOperator(uint64 _keyDepositLimit)
func (_PermissionlessNodeRegistry *PermissionlessNodeRegistryFilterer) FilterUpdatedMaxKeyPerOperator(opts *bind.FilterOpts) (*PermissionlessNodeRegistryUpdatedMaxKeyPerOperatorIterator, error) {

	logs, sub, err := _PermissionlessNodeRegistry.contract.FilterLogs(opts, "UpdatedMaxKeyPerOperator")
	if err != nil {
		return nil, err
	}
	return &PermissionlessNodeRegistryUpdatedMaxKeyPerOperatorIterator{contract: _PermissionlessNodeRegistry.contract, event: "UpdatedMaxKeyPerOperator", logs: logs, sub: sub}, nil
}

// WatchUpdatedMaxKeyPerOperator is a free log subscription operation binding the contract event 0xae32a0fc0865ac1ff229d2eb9e38d025ba79c51aa94e691493815ec2b327ef42.
//
// Solidity: event UpdatedMaxKeyPerOperator(uint64 _keyDepositLimit)
func (_PermissionlessNodeRegistry *PermissionlessNodeRegistryFilterer) WatchUpdatedMaxKeyPerOperator(opts *bind.WatchOpts, sink chan<- *PermissionlessNodeRegistryUpdatedMaxKeyPerOperator) (event.Subscription, error) {

	logs, sub, err := _PermissionlessNodeRegistry.contract.WatchLogs(opts, "UpdatedMaxKeyPerOperator")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PermissionlessNodeRegistryUpdatedMaxKeyPerOperator)
				if err := _PermissionlessNodeRegistry.contract.UnpackLog(event, "UpdatedMaxKeyPerOperator", log); err != nil {
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

// ParseUpdatedMaxKeyPerOperator is a log parse operation binding the contract event 0xae32a0fc0865ac1ff229d2eb9e38d025ba79c51aa94e691493815ec2b327ef42.
//
// Solidity: event UpdatedMaxKeyPerOperator(uint64 _keyDepositLimit)
func (_PermissionlessNodeRegistry *PermissionlessNodeRegistryFilterer) ParseUpdatedMaxKeyPerOperator(log types.Log) (*PermissionlessNodeRegistryUpdatedMaxKeyPerOperator, error) {
	event := new(PermissionlessNodeRegistryUpdatedMaxKeyPerOperator)
	if err := _PermissionlessNodeRegistry.contract.UnpackLog(event, "UpdatedMaxKeyPerOperator", log); err != nil {
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
// Solidity: event UpdatedNextQueuedValidatorIndex(uint256 _nextQueuedValidatorIndex)
func (_PermissionlessNodeRegistry *PermissionlessNodeRegistryFilterer) FilterUpdatedNextQueuedValidatorIndex(opts *bind.FilterOpts) (*PermissionlessNodeRegistryUpdatedNextQueuedValidatorIndexIterator, error) {

	logs, sub, err := _PermissionlessNodeRegistry.contract.FilterLogs(opts, "UpdatedNextQueuedValidatorIndex")
	if err != nil {
		return nil, err
	}
	return &PermissionlessNodeRegistryUpdatedNextQueuedValidatorIndexIterator{contract: _PermissionlessNodeRegistry.contract, event: "UpdatedNextQueuedValidatorIndex", logs: logs, sub: sub}, nil
}

// WatchUpdatedNextQueuedValidatorIndex is a free log subscription operation binding the contract event 0x711359152f2039f4182a096114b0d199c5f8e9cb268caff34080855c42ff4da9.
//
// Solidity: event UpdatedNextQueuedValidatorIndex(uint256 _nextQueuedValidatorIndex)
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
// Solidity: event UpdatedNextQueuedValidatorIndex(uint256 _nextQueuedValidatorIndex)
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
// Solidity: event UpdatedOperatorDetails(address indexed _nodeOperator, string _operatorName, address _rewardAddress)
func (_PermissionlessNodeRegistry *PermissionlessNodeRegistryFilterer) FilterUpdatedOperatorDetails(opts *bind.FilterOpts, _nodeOperator []common.Address) (*PermissionlessNodeRegistryUpdatedOperatorDetailsIterator, error) {

	var _nodeOperatorRule []interface{}
	for _, _nodeOperatorItem := range _nodeOperator {
		_nodeOperatorRule = append(_nodeOperatorRule, _nodeOperatorItem)
	}

	logs, sub, err := _PermissionlessNodeRegistry.contract.FilterLogs(opts, "UpdatedOperatorDetails", _nodeOperatorRule)
	if err != nil {
		return nil, err
	}
	return &PermissionlessNodeRegistryUpdatedOperatorDetailsIterator{contract: _PermissionlessNodeRegistry.contract, event: "UpdatedOperatorDetails", logs: logs, sub: sub}, nil
}

// WatchUpdatedOperatorDetails is a free log subscription operation binding the contract event 0xadc8722095edf061d7fdcb583105c05bf9eb15488503b621c39e254d87269777.
//
// Solidity: event UpdatedOperatorDetails(address indexed _nodeOperator, string _operatorName, address _rewardAddress)
func (_PermissionlessNodeRegistry *PermissionlessNodeRegistryFilterer) WatchUpdatedOperatorDetails(opts *bind.WatchOpts, sink chan<- *PermissionlessNodeRegistryUpdatedOperatorDetails, _nodeOperator []common.Address) (event.Subscription, error) {

	var _nodeOperatorRule []interface{}
	for _, _nodeOperatorItem := range _nodeOperator {
		_nodeOperatorRule = append(_nodeOperatorRule, _nodeOperatorItem)
	}

	logs, sub, err := _PermissionlessNodeRegistry.contract.WatchLogs(opts, "UpdatedOperatorDetails", _nodeOperatorRule)
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
// Solidity: event UpdatedOperatorDetails(address indexed _nodeOperator, string _operatorName, address _rewardAddress)
func (_PermissionlessNodeRegistry *PermissionlessNodeRegistryFilterer) ParseUpdatedOperatorDetails(log types.Log) (*PermissionlessNodeRegistryUpdatedOperatorDetails, error) {
	event := new(PermissionlessNodeRegistryUpdatedOperatorDetails)
	if err := _PermissionlessNodeRegistry.contract.UnpackLog(event, "UpdatedOperatorDetails", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PermissionlessNodeRegistryUpdatedPermissionlessPoolAddressIterator is returned from FilterUpdatedPermissionlessPoolAddress and is used to iterate over the raw logs and unpacked data for UpdatedPermissionlessPoolAddress events raised by the PermissionlessNodeRegistry contract.
type PermissionlessNodeRegistryUpdatedPermissionlessPoolAddressIterator struct {
	Event *PermissionlessNodeRegistryUpdatedPermissionlessPoolAddress // Event containing the contract specifics and raw log

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
func (it *PermissionlessNodeRegistryUpdatedPermissionlessPoolAddressIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PermissionlessNodeRegistryUpdatedPermissionlessPoolAddress)
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
		it.Event = new(PermissionlessNodeRegistryUpdatedPermissionlessPoolAddress)
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
func (it *PermissionlessNodeRegistryUpdatedPermissionlessPoolAddressIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PermissionlessNodeRegistryUpdatedPermissionlessPoolAddressIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PermissionlessNodeRegistryUpdatedPermissionlessPoolAddress represents a UpdatedPermissionlessPoolAddress event raised by the PermissionlessNodeRegistry contract.
type PermissionlessNodeRegistryUpdatedPermissionlessPoolAddress struct {
	PermissionlessPool common.Address
	Raw                types.Log // Blockchain specific contextual infos
}

// FilterUpdatedPermissionlessPoolAddress is a free log retrieval operation binding the contract event 0x575e1a5f20a9e3dba21bf0c6f1b31e2a7b348242016f34c2d3a51488e3e0c8c4.
//
// Solidity: event UpdatedPermissionlessPoolAddress(address _permissionlessPool)
func (_PermissionlessNodeRegistry *PermissionlessNodeRegistryFilterer) FilterUpdatedPermissionlessPoolAddress(opts *bind.FilterOpts) (*PermissionlessNodeRegistryUpdatedPermissionlessPoolAddressIterator, error) {

	logs, sub, err := _PermissionlessNodeRegistry.contract.FilterLogs(opts, "UpdatedPermissionlessPoolAddress")
	if err != nil {
		return nil, err
	}
	return &PermissionlessNodeRegistryUpdatedPermissionlessPoolAddressIterator{contract: _PermissionlessNodeRegistry.contract, event: "UpdatedPermissionlessPoolAddress", logs: logs, sub: sub}, nil
}

// WatchUpdatedPermissionlessPoolAddress is a free log subscription operation binding the contract event 0x575e1a5f20a9e3dba21bf0c6f1b31e2a7b348242016f34c2d3a51488e3e0c8c4.
//
// Solidity: event UpdatedPermissionlessPoolAddress(address _permissionlessPool)
func (_PermissionlessNodeRegistry *PermissionlessNodeRegistryFilterer) WatchUpdatedPermissionlessPoolAddress(opts *bind.WatchOpts, sink chan<- *PermissionlessNodeRegistryUpdatedPermissionlessPoolAddress) (event.Subscription, error) {

	logs, sub, err := _PermissionlessNodeRegistry.contract.WatchLogs(opts, "UpdatedPermissionlessPoolAddress")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PermissionlessNodeRegistryUpdatedPermissionlessPoolAddress)
				if err := _PermissionlessNodeRegistry.contract.UnpackLog(event, "UpdatedPermissionlessPoolAddress", log); err != nil {
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

// ParseUpdatedPermissionlessPoolAddress is a log parse operation binding the contract event 0x575e1a5f20a9e3dba21bf0c6f1b31e2a7b348242016f34c2d3a51488e3e0c8c4.
//
// Solidity: event UpdatedPermissionlessPoolAddress(address _permissionlessPool)
func (_PermissionlessNodeRegistry *PermissionlessNodeRegistryFilterer) ParseUpdatedPermissionlessPoolAddress(log types.Log) (*PermissionlessNodeRegistryUpdatedPermissionlessPoolAddress, error) {
	event := new(PermissionlessNodeRegistryUpdatedPermissionlessPoolAddress)
	if err := _PermissionlessNodeRegistry.contract.UnpackLog(event, "UpdatedPermissionlessPoolAddress", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PermissionlessNodeRegistryUpdatedPoolFactoryAddressIterator is returned from FilterUpdatedPoolFactoryAddress and is used to iterate over the raw logs and unpacked data for UpdatedPoolFactoryAddress events raised by the PermissionlessNodeRegistry contract.
type PermissionlessNodeRegistryUpdatedPoolFactoryAddressIterator struct {
	Event *PermissionlessNodeRegistryUpdatedPoolFactoryAddress // Event containing the contract specifics and raw log

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
func (it *PermissionlessNodeRegistryUpdatedPoolFactoryAddressIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PermissionlessNodeRegistryUpdatedPoolFactoryAddress)
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
		it.Event = new(PermissionlessNodeRegistryUpdatedPoolFactoryAddress)
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
func (it *PermissionlessNodeRegistryUpdatedPoolFactoryAddressIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PermissionlessNodeRegistryUpdatedPoolFactoryAddressIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PermissionlessNodeRegistryUpdatedPoolFactoryAddress represents a UpdatedPoolFactoryAddress event raised by the PermissionlessNodeRegistry contract.
type PermissionlessNodeRegistryUpdatedPoolFactoryAddress struct {
	PoolFactoryAddress common.Address
	Raw                types.Log // Blockchain specific contextual infos
}

// FilterUpdatedPoolFactoryAddress is a free log retrieval operation binding the contract event 0xfcc03bf9f8fce7313625292145aee2957468243bb7c022ddf4809d6fe639ab38.
//
// Solidity: event UpdatedPoolFactoryAddress(address _poolFactoryAddress)
func (_PermissionlessNodeRegistry *PermissionlessNodeRegistryFilterer) FilterUpdatedPoolFactoryAddress(opts *bind.FilterOpts) (*PermissionlessNodeRegistryUpdatedPoolFactoryAddressIterator, error) {

	logs, sub, err := _PermissionlessNodeRegistry.contract.FilterLogs(opts, "UpdatedPoolFactoryAddress")
	if err != nil {
		return nil, err
	}
	return &PermissionlessNodeRegistryUpdatedPoolFactoryAddressIterator{contract: _PermissionlessNodeRegistry.contract, event: "UpdatedPoolFactoryAddress", logs: logs, sub: sub}, nil
}

// WatchUpdatedPoolFactoryAddress is a free log subscription operation binding the contract event 0xfcc03bf9f8fce7313625292145aee2957468243bb7c022ddf4809d6fe639ab38.
//
// Solidity: event UpdatedPoolFactoryAddress(address _poolFactoryAddress)
func (_PermissionlessNodeRegistry *PermissionlessNodeRegistryFilterer) WatchUpdatedPoolFactoryAddress(opts *bind.WatchOpts, sink chan<- *PermissionlessNodeRegistryUpdatedPoolFactoryAddress) (event.Subscription, error) {

	logs, sub, err := _PermissionlessNodeRegistry.contract.WatchLogs(opts, "UpdatedPoolFactoryAddress")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PermissionlessNodeRegistryUpdatedPoolFactoryAddress)
				if err := _PermissionlessNodeRegistry.contract.UnpackLog(event, "UpdatedPoolFactoryAddress", log); err != nil {
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

// ParseUpdatedPoolFactoryAddress is a log parse operation binding the contract event 0xfcc03bf9f8fce7313625292145aee2957468243bb7c022ddf4809d6fe639ab38.
//
// Solidity: event UpdatedPoolFactoryAddress(address _poolFactoryAddress)
func (_PermissionlessNodeRegistry *PermissionlessNodeRegistryFilterer) ParseUpdatedPoolFactoryAddress(log types.Log) (*PermissionlessNodeRegistryUpdatedPoolFactoryAddress, error) {
	event := new(PermissionlessNodeRegistryUpdatedPoolFactoryAddress)
	if err := _PermissionlessNodeRegistry.contract.UnpackLog(event, "UpdatedPoolFactoryAddress", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PermissionlessNodeRegistryUpdatedSDCollateralAddressIterator is returned from FilterUpdatedSDCollateralAddress and is used to iterate over the raw logs and unpacked data for UpdatedSDCollateralAddress events raised by the PermissionlessNodeRegistry contract.
type PermissionlessNodeRegistryUpdatedSDCollateralAddressIterator struct {
	Event *PermissionlessNodeRegistryUpdatedSDCollateralAddress // Event containing the contract specifics and raw log

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
func (it *PermissionlessNodeRegistryUpdatedSDCollateralAddressIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PermissionlessNodeRegistryUpdatedSDCollateralAddress)
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
		it.Event = new(PermissionlessNodeRegistryUpdatedSDCollateralAddress)
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
func (it *PermissionlessNodeRegistryUpdatedSDCollateralAddressIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PermissionlessNodeRegistryUpdatedSDCollateralAddressIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PermissionlessNodeRegistryUpdatedSDCollateralAddress represents a UpdatedSDCollateralAddress event raised by the PermissionlessNodeRegistry contract.
type PermissionlessNodeRegistryUpdatedSDCollateralAddress struct {
	SdCollateral common.Address
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterUpdatedSDCollateralAddress is a free log retrieval operation binding the contract event 0xe1304098c1073db7fb2fe96b01be93d6d980015e6edf1ec0d06f2b2d539c0751.
//
// Solidity: event UpdatedSDCollateralAddress(address _sdCollateral)
func (_PermissionlessNodeRegistry *PermissionlessNodeRegistryFilterer) FilterUpdatedSDCollateralAddress(opts *bind.FilterOpts) (*PermissionlessNodeRegistryUpdatedSDCollateralAddressIterator, error) {

	logs, sub, err := _PermissionlessNodeRegistry.contract.FilterLogs(opts, "UpdatedSDCollateralAddress")
	if err != nil {
		return nil, err
	}
	return &PermissionlessNodeRegistryUpdatedSDCollateralAddressIterator{contract: _PermissionlessNodeRegistry.contract, event: "UpdatedSDCollateralAddress", logs: logs, sub: sub}, nil
}

// WatchUpdatedSDCollateralAddress is a free log subscription operation binding the contract event 0xe1304098c1073db7fb2fe96b01be93d6d980015e6edf1ec0d06f2b2d539c0751.
//
// Solidity: event UpdatedSDCollateralAddress(address _sdCollateral)
func (_PermissionlessNodeRegistry *PermissionlessNodeRegistryFilterer) WatchUpdatedSDCollateralAddress(opts *bind.WatchOpts, sink chan<- *PermissionlessNodeRegistryUpdatedSDCollateralAddress) (event.Subscription, error) {

	logs, sub, err := _PermissionlessNodeRegistry.contract.WatchLogs(opts, "UpdatedSDCollateralAddress")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PermissionlessNodeRegistryUpdatedSDCollateralAddress)
				if err := _PermissionlessNodeRegistry.contract.UnpackLog(event, "UpdatedSDCollateralAddress", log); err != nil {
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

// ParseUpdatedSDCollateralAddress is a log parse operation binding the contract event 0xe1304098c1073db7fb2fe96b01be93d6d980015e6edf1ec0d06f2b2d539c0751.
//
// Solidity: event UpdatedSDCollateralAddress(address _sdCollateral)
func (_PermissionlessNodeRegistry *PermissionlessNodeRegistryFilterer) ParseUpdatedSDCollateralAddress(log types.Log) (*PermissionlessNodeRegistryUpdatedSDCollateralAddress, error) {
	event := new(PermissionlessNodeRegistryUpdatedSDCollateralAddress)
	if err := _PermissionlessNodeRegistry.contract.UnpackLog(event, "UpdatedSDCollateralAddress", log); err != nil {
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
	Timestamp               *big.Int
	Raw                     types.Log // Blockchain specific contextual infos
}

// FilterUpdatedSocializingPoolState is a free log retrieval operation binding the contract event 0xc0465abaf1d51829975919c02418d521476b44f330a31d78bb6b4e96465e746b.
//
// Solidity: event UpdatedSocializingPoolState(uint256 _operatorId, bool _optedForSocializingPool, uint256 timestamp)
func (_PermissionlessNodeRegistry *PermissionlessNodeRegistryFilterer) FilterUpdatedSocializingPoolState(opts *bind.FilterOpts) (*PermissionlessNodeRegistryUpdatedSocializingPoolStateIterator, error) {

	logs, sub, err := _PermissionlessNodeRegistry.contract.FilterLogs(opts, "UpdatedSocializingPoolState")
	if err != nil {
		return nil, err
	}
	return &PermissionlessNodeRegistryUpdatedSocializingPoolStateIterator{contract: _PermissionlessNodeRegistry.contract, event: "UpdatedSocializingPoolState", logs: logs, sub: sub}, nil
}

// WatchUpdatedSocializingPoolState is a free log subscription operation binding the contract event 0xc0465abaf1d51829975919c02418d521476b44f330a31d78bb6b4e96465e746b.
//
// Solidity: event UpdatedSocializingPoolState(uint256 _operatorId, bool _optedForSocializingPool, uint256 timestamp)
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
// Solidity: event UpdatedSocializingPoolState(uint256 _operatorId, bool _optedForSocializingPool, uint256 timestamp)
func (_PermissionlessNodeRegistry *PermissionlessNodeRegistryFilterer) ParseUpdatedSocializingPoolState(log types.Log) (*PermissionlessNodeRegistryUpdatedSocializingPoolState, error) {
	event := new(PermissionlessNodeRegistryUpdatedSocializingPoolState)
	if err := _PermissionlessNodeRegistry.contract.UnpackLog(event, "UpdatedSocializingPoolState", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PermissionlessNodeRegistryUpdatedStaderPenaltyFundIterator is returned from FilterUpdatedStaderPenaltyFund and is used to iterate over the raw logs and unpacked data for UpdatedStaderPenaltyFund events raised by the PermissionlessNodeRegistry contract.
type PermissionlessNodeRegistryUpdatedStaderPenaltyFundIterator struct {
	Event *PermissionlessNodeRegistryUpdatedStaderPenaltyFund // Event containing the contract specifics and raw log

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
func (it *PermissionlessNodeRegistryUpdatedStaderPenaltyFundIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PermissionlessNodeRegistryUpdatedStaderPenaltyFund)
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
		it.Event = new(PermissionlessNodeRegistryUpdatedStaderPenaltyFund)
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
func (it *PermissionlessNodeRegistryUpdatedStaderPenaltyFundIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PermissionlessNodeRegistryUpdatedStaderPenaltyFundIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PermissionlessNodeRegistryUpdatedStaderPenaltyFund represents a UpdatedStaderPenaltyFund event raised by the PermissionlessNodeRegistry contract.
type PermissionlessNodeRegistryUpdatedStaderPenaltyFund struct {
	StaderPenaltyFund common.Address
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterUpdatedStaderPenaltyFund is a free log retrieval operation binding the contract event 0x186340150b943649316d1a10d14c47e9906bd95b16fa6fdb763b235fcc9693b0.
//
// Solidity: event UpdatedStaderPenaltyFund(address _staderPenaltyFund)
func (_PermissionlessNodeRegistry *PermissionlessNodeRegistryFilterer) FilterUpdatedStaderPenaltyFund(opts *bind.FilterOpts) (*PermissionlessNodeRegistryUpdatedStaderPenaltyFundIterator, error) {

	logs, sub, err := _PermissionlessNodeRegistry.contract.FilterLogs(opts, "UpdatedStaderPenaltyFund")
	if err != nil {
		return nil, err
	}
	return &PermissionlessNodeRegistryUpdatedStaderPenaltyFundIterator{contract: _PermissionlessNodeRegistry.contract, event: "UpdatedStaderPenaltyFund", logs: logs, sub: sub}, nil
}

// WatchUpdatedStaderPenaltyFund is a free log subscription operation binding the contract event 0x186340150b943649316d1a10d14c47e9906bd95b16fa6fdb763b235fcc9693b0.
//
// Solidity: event UpdatedStaderPenaltyFund(address _staderPenaltyFund)
func (_PermissionlessNodeRegistry *PermissionlessNodeRegistryFilterer) WatchUpdatedStaderPenaltyFund(opts *bind.WatchOpts, sink chan<- *PermissionlessNodeRegistryUpdatedStaderPenaltyFund) (event.Subscription, error) {

	logs, sub, err := _PermissionlessNodeRegistry.contract.WatchLogs(opts, "UpdatedStaderPenaltyFund")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PermissionlessNodeRegistryUpdatedStaderPenaltyFund)
				if err := _PermissionlessNodeRegistry.contract.UnpackLog(event, "UpdatedStaderPenaltyFund", log); err != nil {
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

// ParseUpdatedStaderPenaltyFund is a log parse operation binding the contract event 0x186340150b943649316d1a10d14c47e9906bd95b16fa6fdb763b235fcc9693b0.
//
// Solidity: event UpdatedStaderPenaltyFund(address _staderPenaltyFund)
func (_PermissionlessNodeRegistry *PermissionlessNodeRegistryFilterer) ParseUpdatedStaderPenaltyFund(log types.Log) (*PermissionlessNodeRegistryUpdatedStaderPenaltyFund, error) {
	event := new(PermissionlessNodeRegistryUpdatedStaderPenaltyFund)
	if err := _PermissionlessNodeRegistry.contract.UnpackLog(event, "UpdatedStaderPenaltyFund", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PermissionlessNodeRegistryUpdatedVaultFactoryAddressIterator is returned from FilterUpdatedVaultFactoryAddress and is used to iterate over the raw logs and unpacked data for UpdatedVaultFactoryAddress events raised by the PermissionlessNodeRegistry contract.
type PermissionlessNodeRegistryUpdatedVaultFactoryAddressIterator struct {
	Event *PermissionlessNodeRegistryUpdatedVaultFactoryAddress // Event containing the contract specifics and raw log

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
func (it *PermissionlessNodeRegistryUpdatedVaultFactoryAddressIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PermissionlessNodeRegistryUpdatedVaultFactoryAddress)
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
		it.Event = new(PermissionlessNodeRegistryUpdatedVaultFactoryAddress)
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
func (it *PermissionlessNodeRegistryUpdatedVaultFactoryAddressIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PermissionlessNodeRegistryUpdatedVaultFactoryAddressIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PermissionlessNodeRegistryUpdatedVaultFactoryAddress represents a UpdatedVaultFactoryAddress event raised by the PermissionlessNodeRegistry contract.
type PermissionlessNodeRegistryUpdatedVaultFactoryAddress struct {
	VaultFactoryAddress common.Address
	Raw                 types.Log // Blockchain specific contextual infos
}

// FilterUpdatedVaultFactoryAddress is a free log retrieval operation binding the contract event 0x4918fdae000037899c8f3e83a59ced7bf92783cad4d6b3025ec319a4a01a1669.
//
// Solidity: event UpdatedVaultFactoryAddress(address _vaultFactoryAddress)
func (_PermissionlessNodeRegistry *PermissionlessNodeRegistryFilterer) FilterUpdatedVaultFactoryAddress(opts *bind.FilterOpts) (*PermissionlessNodeRegistryUpdatedVaultFactoryAddressIterator, error) {

	logs, sub, err := _PermissionlessNodeRegistry.contract.FilterLogs(opts, "UpdatedVaultFactoryAddress")
	if err != nil {
		return nil, err
	}
	return &PermissionlessNodeRegistryUpdatedVaultFactoryAddressIterator{contract: _PermissionlessNodeRegistry.contract, event: "UpdatedVaultFactoryAddress", logs: logs, sub: sub}, nil
}

// WatchUpdatedVaultFactoryAddress is a free log subscription operation binding the contract event 0x4918fdae000037899c8f3e83a59ced7bf92783cad4d6b3025ec319a4a01a1669.
//
// Solidity: event UpdatedVaultFactoryAddress(address _vaultFactoryAddress)
func (_PermissionlessNodeRegistry *PermissionlessNodeRegistryFilterer) WatchUpdatedVaultFactoryAddress(opts *bind.WatchOpts, sink chan<- *PermissionlessNodeRegistryUpdatedVaultFactoryAddress) (event.Subscription, error) {

	logs, sub, err := _PermissionlessNodeRegistry.contract.WatchLogs(opts, "UpdatedVaultFactoryAddress")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PermissionlessNodeRegistryUpdatedVaultFactoryAddress)
				if err := _PermissionlessNodeRegistry.contract.UnpackLog(event, "UpdatedVaultFactoryAddress", log); err != nil {
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

// ParseUpdatedVaultFactoryAddress is a log parse operation binding the contract event 0x4918fdae000037899c8f3e83a59ced7bf92783cad4d6b3025ec319a4a01a1669.
//
// Solidity: event UpdatedVaultFactoryAddress(address _vaultFactoryAddress)
func (_PermissionlessNodeRegistry *PermissionlessNodeRegistryFilterer) ParseUpdatedVaultFactoryAddress(log types.Log) (*PermissionlessNodeRegistryUpdatedVaultFactoryAddress, error) {
	event := new(PermissionlessNodeRegistryUpdatedVaultFactoryAddress)
	if err := _PermissionlessNodeRegistry.contract.UnpackLog(event, "UpdatedVaultFactoryAddress", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PermissionlessNodeRegistryValidatorDepositTimeSetIterator is returned from FilterValidatorDepositTimeSet and is used to iterate over the raw logs and unpacked data for ValidatorDepositTimeSet events raised by the PermissionlessNodeRegistry contract.
type PermissionlessNodeRegistryValidatorDepositTimeSetIterator struct {
	Event *PermissionlessNodeRegistryValidatorDepositTimeSet // Event containing the contract specifics and raw log

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
func (it *PermissionlessNodeRegistryValidatorDepositTimeSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PermissionlessNodeRegistryValidatorDepositTimeSet)
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
		it.Event = new(PermissionlessNodeRegistryValidatorDepositTimeSet)
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
func (it *PermissionlessNodeRegistryValidatorDepositTimeSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PermissionlessNodeRegistryValidatorDepositTimeSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PermissionlessNodeRegistryValidatorDepositTimeSet represents a ValidatorDepositTimeSet event raised by the PermissionlessNodeRegistry contract.
type PermissionlessNodeRegistryValidatorDepositTimeSet struct {
	ValidatorId *big.Int
	DepositTime *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterValidatorDepositTimeSet is a free log retrieval operation binding the contract event 0xd7d27755850483d7c171ef91a1319df4293f437d249db3d03183d9fec8c2b313.
//
// Solidity: event ValidatorDepositTimeSet(uint256 _validatorId, uint256 _depositTime)
func (_PermissionlessNodeRegistry *PermissionlessNodeRegistryFilterer) FilterValidatorDepositTimeSet(opts *bind.FilterOpts) (*PermissionlessNodeRegistryValidatorDepositTimeSetIterator, error) {

	logs, sub, err := _PermissionlessNodeRegistry.contract.FilterLogs(opts, "ValidatorDepositTimeSet")
	if err != nil {
		return nil, err
	}
	return &PermissionlessNodeRegistryValidatorDepositTimeSetIterator{contract: _PermissionlessNodeRegistry.contract, event: "ValidatorDepositTimeSet", logs: logs, sub: sub}, nil
}

// WatchValidatorDepositTimeSet is a free log subscription operation binding the contract event 0xd7d27755850483d7c171ef91a1319df4293f437d249db3d03183d9fec8c2b313.
//
// Solidity: event ValidatorDepositTimeSet(uint256 _validatorId, uint256 _depositTime)
func (_PermissionlessNodeRegistry *PermissionlessNodeRegistryFilterer) WatchValidatorDepositTimeSet(opts *bind.WatchOpts, sink chan<- *PermissionlessNodeRegistryValidatorDepositTimeSet) (event.Subscription, error) {

	logs, sub, err := _PermissionlessNodeRegistry.contract.WatchLogs(opts, "ValidatorDepositTimeSet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PermissionlessNodeRegistryValidatorDepositTimeSet)
				if err := _PermissionlessNodeRegistry.contract.UnpackLog(event, "ValidatorDepositTimeSet", log); err != nil {
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

// ParseValidatorDepositTimeSet is a log parse operation binding the contract event 0xd7d27755850483d7c171ef91a1319df4293f437d249db3d03183d9fec8c2b313.
//
// Solidity: event ValidatorDepositTimeSet(uint256 _validatorId, uint256 _depositTime)
func (_PermissionlessNodeRegistry *PermissionlessNodeRegistryFilterer) ParseValidatorDepositTimeSet(log types.Log) (*PermissionlessNodeRegistryValidatorDepositTimeSet, error) {
	event := new(PermissionlessNodeRegistryValidatorDepositTimeSet)
	if err := _PermissionlessNodeRegistry.contract.UnpackLog(event, "ValidatorDepositTimeSet", log); err != nil {
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
	FrontRunnedPubkey common.Hash
	ValidatorId       *big.Int
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterValidatorMarkedAsFrontRunned is a free log retrieval operation binding the contract event 0x4e93215f00bc729272f0ff71afd3d0f385208cbf6c999fe776ad07c623b83466.
//
// Solidity: event ValidatorMarkedAsFrontRunned(bytes indexed _frontRunnedPubkey, uint256 _validatorId)
func (_PermissionlessNodeRegistry *PermissionlessNodeRegistryFilterer) FilterValidatorMarkedAsFrontRunned(opts *bind.FilterOpts, _frontRunnedPubkey [][]byte) (*PermissionlessNodeRegistryValidatorMarkedAsFrontRunnedIterator, error) {

	var _frontRunnedPubkeyRule []interface{}
	for _, _frontRunnedPubkeyItem := range _frontRunnedPubkey {
		_frontRunnedPubkeyRule = append(_frontRunnedPubkeyRule, _frontRunnedPubkeyItem)
	}

	logs, sub, err := _PermissionlessNodeRegistry.contract.FilterLogs(opts, "ValidatorMarkedAsFrontRunned", _frontRunnedPubkeyRule)
	if err != nil {
		return nil, err
	}
	return &PermissionlessNodeRegistryValidatorMarkedAsFrontRunnedIterator{contract: _PermissionlessNodeRegistry.contract, event: "ValidatorMarkedAsFrontRunned", logs: logs, sub: sub}, nil
}

// WatchValidatorMarkedAsFrontRunned is a free log subscription operation binding the contract event 0x4e93215f00bc729272f0ff71afd3d0f385208cbf6c999fe776ad07c623b83466.
//
// Solidity: event ValidatorMarkedAsFrontRunned(bytes indexed _frontRunnedPubkey, uint256 _validatorId)
func (_PermissionlessNodeRegistry *PermissionlessNodeRegistryFilterer) WatchValidatorMarkedAsFrontRunned(opts *bind.WatchOpts, sink chan<- *PermissionlessNodeRegistryValidatorMarkedAsFrontRunned, _frontRunnedPubkey [][]byte) (event.Subscription, error) {

	var _frontRunnedPubkeyRule []interface{}
	for _, _frontRunnedPubkeyItem := range _frontRunnedPubkey {
		_frontRunnedPubkeyRule = append(_frontRunnedPubkeyRule, _frontRunnedPubkeyItem)
	}

	logs, sub, err := _PermissionlessNodeRegistry.contract.WatchLogs(opts, "ValidatorMarkedAsFrontRunned", _frontRunnedPubkeyRule)
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
// Solidity: event ValidatorMarkedAsFrontRunned(bytes indexed _frontRunnedPubkey, uint256 _validatorId)
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
	Pubkey      common.Hash
	ValidatorId *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterValidatorMarkedReadyToDeposit is a free log retrieval operation binding the contract event 0x21d79a0b22a7d5a18b9535162fe2f0580e24c042b0541a05afc298a77ddf5693.
//
// Solidity: event ValidatorMarkedReadyToDeposit(bytes indexed _pubkey, uint256 _validatorId)
func (_PermissionlessNodeRegistry *PermissionlessNodeRegistryFilterer) FilterValidatorMarkedReadyToDeposit(opts *bind.FilterOpts, _pubkey [][]byte) (*PermissionlessNodeRegistryValidatorMarkedReadyToDepositIterator, error) {

	var _pubkeyRule []interface{}
	for _, _pubkeyItem := range _pubkey {
		_pubkeyRule = append(_pubkeyRule, _pubkeyItem)
	}

	logs, sub, err := _PermissionlessNodeRegistry.contract.FilterLogs(opts, "ValidatorMarkedReadyToDeposit", _pubkeyRule)
	if err != nil {
		return nil, err
	}
	return &PermissionlessNodeRegistryValidatorMarkedReadyToDepositIterator{contract: _PermissionlessNodeRegistry.contract, event: "ValidatorMarkedReadyToDeposit", logs: logs, sub: sub}, nil
}

// WatchValidatorMarkedReadyToDeposit is a free log subscription operation binding the contract event 0x21d79a0b22a7d5a18b9535162fe2f0580e24c042b0541a05afc298a77ddf5693.
//
// Solidity: event ValidatorMarkedReadyToDeposit(bytes indexed _pubkey, uint256 _validatorId)
func (_PermissionlessNodeRegistry *PermissionlessNodeRegistryFilterer) WatchValidatorMarkedReadyToDeposit(opts *bind.WatchOpts, sink chan<- *PermissionlessNodeRegistryValidatorMarkedReadyToDeposit, _pubkey [][]byte) (event.Subscription, error) {

	var _pubkeyRule []interface{}
	for _, _pubkeyItem := range _pubkey {
		_pubkeyRule = append(_pubkeyRule, _pubkeyItem)
	}

	logs, sub, err := _PermissionlessNodeRegistry.contract.WatchLogs(opts, "ValidatorMarkedReadyToDeposit", _pubkeyRule)
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
// Solidity: event ValidatorMarkedReadyToDeposit(bytes indexed _pubkey, uint256 _validatorId)
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
	InvalidSignaturePubkey common.Hash
	ValidatorId            *big.Int
	Raw                    types.Log // Blockchain specific contextual infos
}

// FilterValidatorStatusMarkedAsInvalidSignature is a free log retrieval operation binding the contract event 0x596ee835bed6cb827d21ba1785c468f0755ee40d33d87132df5d2ec90b645f9f.
//
// Solidity: event ValidatorStatusMarkedAsInvalidSignature(bytes indexed invalidSignaturePubkey, uint256 _validatorId)
func (_PermissionlessNodeRegistry *PermissionlessNodeRegistryFilterer) FilterValidatorStatusMarkedAsInvalidSignature(opts *bind.FilterOpts, invalidSignaturePubkey [][]byte) (*PermissionlessNodeRegistryValidatorStatusMarkedAsInvalidSignatureIterator, error) {

	var invalidSignaturePubkeyRule []interface{}
	for _, invalidSignaturePubkeyItem := range invalidSignaturePubkey {
		invalidSignaturePubkeyRule = append(invalidSignaturePubkeyRule, invalidSignaturePubkeyItem)
	}

	logs, sub, err := _PermissionlessNodeRegistry.contract.FilterLogs(opts, "ValidatorStatusMarkedAsInvalidSignature", invalidSignaturePubkeyRule)
	if err != nil {
		return nil, err
	}
	return &PermissionlessNodeRegistryValidatorStatusMarkedAsInvalidSignatureIterator{contract: _PermissionlessNodeRegistry.contract, event: "ValidatorStatusMarkedAsInvalidSignature", logs: logs, sub: sub}, nil
}

// WatchValidatorStatusMarkedAsInvalidSignature is a free log subscription operation binding the contract event 0x596ee835bed6cb827d21ba1785c468f0755ee40d33d87132df5d2ec90b645f9f.
//
// Solidity: event ValidatorStatusMarkedAsInvalidSignature(bytes indexed invalidSignaturePubkey, uint256 _validatorId)
func (_PermissionlessNodeRegistry *PermissionlessNodeRegistryFilterer) WatchValidatorStatusMarkedAsInvalidSignature(opts *bind.WatchOpts, sink chan<- *PermissionlessNodeRegistryValidatorStatusMarkedAsInvalidSignature, invalidSignaturePubkey [][]byte) (event.Subscription, error) {

	var invalidSignaturePubkeyRule []interface{}
	for _, invalidSignaturePubkeyItem := range invalidSignaturePubkey {
		invalidSignaturePubkeyRule = append(invalidSignaturePubkeyRule, invalidSignaturePubkeyItem)
	}

	logs, sub, err := _PermissionlessNodeRegistry.contract.WatchLogs(opts, "ValidatorStatusMarkedAsInvalidSignature", invalidSignaturePubkeyRule)
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
// Solidity: event ValidatorStatusMarkedAsInvalidSignature(bytes indexed invalidSignaturePubkey, uint256 _validatorId)
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
	Pubkey      common.Hash
	ValidatorId *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterValidatorWithdrawn is a free log retrieval operation binding the contract event 0x450186694fefe67df6156f60235e4073b623160f28a0b85908ebc864316abf79.
//
// Solidity: event ValidatorWithdrawn(bytes indexed _pubkey, uint256 _validatorId)
func (_PermissionlessNodeRegistry *PermissionlessNodeRegistryFilterer) FilterValidatorWithdrawn(opts *bind.FilterOpts, _pubkey [][]byte) (*PermissionlessNodeRegistryValidatorWithdrawnIterator, error) {

	var _pubkeyRule []interface{}
	for _, _pubkeyItem := range _pubkey {
		_pubkeyRule = append(_pubkeyRule, _pubkeyItem)
	}

	logs, sub, err := _PermissionlessNodeRegistry.contract.FilterLogs(opts, "ValidatorWithdrawn", _pubkeyRule)
	if err != nil {
		return nil, err
	}
	return &PermissionlessNodeRegistryValidatorWithdrawnIterator{contract: _PermissionlessNodeRegistry.contract, event: "ValidatorWithdrawn", logs: logs, sub: sub}, nil
}

// WatchValidatorWithdrawn is a free log subscription operation binding the contract event 0x450186694fefe67df6156f60235e4073b623160f28a0b85908ebc864316abf79.
//
// Solidity: event ValidatorWithdrawn(bytes indexed _pubkey, uint256 _validatorId)
func (_PermissionlessNodeRegistry *PermissionlessNodeRegistryFilterer) WatchValidatorWithdrawn(opts *bind.WatchOpts, sink chan<- *PermissionlessNodeRegistryValidatorWithdrawn, _pubkey [][]byte) (event.Subscription, error) {

	var _pubkeyRule []interface{}
	for _, _pubkeyItem := range _pubkey {
		_pubkeyRule = append(_pubkeyRule, _pubkeyItem)
	}

	logs, sub, err := _PermissionlessNodeRegistry.contract.WatchLogs(opts, "ValidatorWithdrawn", _pubkeyRule)
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
// Solidity: event ValidatorWithdrawn(bytes indexed _pubkey, uint256 _validatorId)
func (_PermissionlessNodeRegistry *PermissionlessNodeRegistryFilterer) ParseValidatorWithdrawn(log types.Log) (*PermissionlessNodeRegistryValidatorWithdrawn, error) {
	event := new(PermissionlessNodeRegistryValidatorWithdrawn)
	if err := _PermissionlessNodeRegistry.contract.UnpackLog(event, "ValidatorWithdrawn", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
