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
package api

import (
	"github.com/stader-labs/stader-node/shared/utils/stdr"
	"github.com/stader-labs/stader-node/stader-lib/stader"
	"github.com/stader-labs/stader-node/stader-lib/types"
	"math/big"

	"github.com/ethereum/go-ethereum/common"

	"github.com/stader-labs/stader-node/stader-lib/tokens"
)

type NodeStatusResponse struct {
	Status                       string               `json:"status"`
	Error                        string               `json:"error"`
	NumberOfValidatorsRegistered string               `json:"numberOfValidatorsRegistered"`
	AccountAddress               common.Address       `json:"accountAddress"`
	AccountAddressFormatted      string               `json:"accountAddressFormatted"`
	OperatorId                   *big.Int             `json:"operatorId"`
	OperatorName                 string               `json:"operatorName"`
	OperatorRewardAddress        common.Address       `json:"operatorRewardAddress"`
	DepositedSdCollateral        *big.Int             `json:"depositedSdCollateral"`
	SdCollateralWorthValidators  *big.Int             `json:"sdCollateralWorthValidators"`
	Registered                   bool                 `json:"registered"`
	AccountBalances              tokens.Balances      `json:"accountBalances"`
	ValidatorInfos               []stdr.ValidatorInfo `json:"validatorInfos"`
}

type CanRegisterNodeResponse struct {
	Status             string         `json:"status"`
	Error              string         `json:"error"`
	CanRegister        bool           `json:"canRegister"`
	AlreadyRegistered  bool           `json:"alreadyRegistered"`
	RegistrationPaused bool           `json:"registrationPaused"`
	GasInfo            stader.GasInfo `json:"gasInfo"`
}
type RegisterNodeResponse struct {
	Status string      `json:"status"`
	Error  string      `json:"error"`
	TxHash common.Hash `json:"txHash"`
}

type CanNodeDepositSdResponse struct {
	Status              string         `json:"status"`
	Error               string         `json:"error"`
	CanDeposit          bool           `json:"CanDeposit"`
	InsufficientBalance bool           `json:"insufficientBalance"`
	InConsensus         bool           `json:"inConsensus"`
	GasInfo             stader.GasInfo `json:"gasInfo"`
}
type NodeDepositSdApproveGasResponse struct {
	Status  string         `json:"status"`
	Error   string         `json:"error"`
	GasInfo stader.GasInfo `json:"gasInfo"`
}
type NodeDepositSdApproveResponse struct {
	Status        string      `json:"status"`
	Error         string      `json:"error"`
	ApproveTxHash common.Hash `json:"approveTxHash"`
}
type NodeDepositSdResponse struct {
	Status        string      `json:"status"`
	Error         string      `json:"error"`
	DepositTxHash common.Hash `json:"stakeTxHash"`
}
type NodeDepositSdAllowanceResponse struct {
	Status    string   `json:"status"`
	Error     string   `json:"error"`
	Allowance *big.Int `json:"allowance"`
}

type CanNodeDepositResponse struct {
	Status                string         `json:"status"`
	Error                 string         `json:"error"`
	CanDeposit            bool           `json:"CanDeposit"`
	InsufficientBalance   bool           `json:"insufficientBalance"`
	InvalidAmount         bool           `json:"invalidAmount"`
	NotRegistered         bool           `json:"notRegistered"`
	DepositPaused         bool           `json:"depositPaused"`
	NotEnoughSdCollateral bool           `json:"notEnoughSdCollateral"`
	GasInfo               stader.GasInfo `json:"gasInfo"`
}

type NodeDepositResponse struct {
	Status string      `json:"status"`
	Error  string      `json:"error"`
	TxHash common.Hash `json:"txHash"`
}

type CanNodeSendResponse struct {
	Status              string         `json:"status"`
	Error               string         `json:"error"`
	CanSend             bool           `json:"canSend"`
	InsufficientBalance bool           `json:"insufficientBalance"`
	GasInfo             stader.GasInfo `json:"gasInfo"`
}
type NodeSendResponse struct {
	Status string      `json:"status"`
	Error  string      `json:"error"`
	TxHash common.Hash `json:"txHash"`
}

type NodeSyncProgressResponse struct {
	Status   string              `json:"status"`
	Error    string              `json:"error"`
	EcStatus ClientManagerStatus `json:"ecStatus"`
	BcStatus ClientManagerStatus `json:"bcStatus"`
}

type ContractsInfoResponse struct {
	Status                     string         `json:"status"`
	Error                      string         `json:"error"`
	Network                    uint64         `json:"network"`
	BeaconDepositContract      common.Address `json:"beaconDepositContract"`
	BeaconNetwork              uint64         `json:"beaconNetwork"`
	PermissionlessNodeRegistry common.Address `json:"permissionlessNodeRegistry"`
	VaultFactory               common.Address `json:"vaultFactory"`
	EthxToken                  common.Address `json:"ethxToken"`
	SdToken                    common.Address `json:"sdToken"`
	SdCollateralContract       common.Address `json:"sdCollateralContract"`
}

type DebugExitResponse struct {
	Status          string                   `json:"status"`
	Error           string                   `json:"error"`
	ValidatorPubKey types.ValidatorPubkey    `json:"validatorPubKey"`
	ExitEpoch       uint64                   `json:"exitEpoch"`
	ValidatorIndex  uint64                   `json:"validatorIndex"`
	SrHash          common.Hash              `json:"srHash"`
	SignedMsg       types.ValidatorSignature `json:"signedMsg"`
	SignatureDomain common.Hash              `json:"signatureDomain"`
}

type NodeSignResponse struct {
	Status     string `json:"status"`
	Error      string `json:"error"`
	SignedData string `json:"signedData"`
}
