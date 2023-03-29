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
	CurrentEpoch    uint64                   `json:"currentEpoch"`
	ValidatorIndex  uint64                   `json:"validatorIndex"`
	SrHash          common.Hash              `json:"srHash"`
	SignedMsg       types.ValidatorSignature `json:"signedMsg"`
	SignatureDomain common.Hash              `json:"signatureDomain"`
}

type CanSendPresignedMsgResponse struct {
	Status                               string `json:"status"`
	Error                                string `json:"error"`
	ValidatorNotRegistered               bool   `json:"validatorNotRegistered"`
	ValidatorPreSignKeyAlreadyRegistered bool   `json:"validatorPreSignKeyAlreadyRegistered"`
}

type SendPresignedMsgResponse struct {
	Status          string                   `json:"status"`
	Error           string                   `json:"error"`
	ValidatorPubKey types.ValidatorPubkey    `json:"validatorPubKey"`
	ExitEpoch       uint64                   `json:"exitEpoch"`
	ValidatorIndex  uint64                   `json:"validatorIndex"`
	SignedMsg       types.ValidatorSignature `json:"signedMsg"`
}

type CanExitValidatorResponse struct {
	Status                 string `json:"status"`
	Error                  string `json:"error"`
	ValidatorNotRegistered bool   `json:"validatorNotRegistered"`
	CanExit                bool   `json:"canExit"`
}

type ExitValidatorResponse struct {
	Status string `json:"status"`
	Error  string `json:"error"`
}

type NodeSignResponse struct {
	Status     string `json:"status"`
	Error      string `json:"error"`
	SignedData string `json:"signedData"`
}
