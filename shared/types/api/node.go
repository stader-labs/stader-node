package api

import (
	stader_backend "github.com/stader-labs/stader-node/shared/types/stader-backend"
	"math/big"
	"time"

	"github.com/stader-labs/stader-node/shared/utils/stdr"
	"github.com/stader-labs/stader-node/stader-lib/stader"
	"github.com/stader-labs/stader-node/stader-lib/types"

	"github.com/ethereum/go-ethereum/common"

	"github.com/stader-labs/stader-node/stader-lib/tokens"
)

type NodeStatusResponse struct {
	Status                            string                             `json:"status"`
	Error                             string                             `json:"error"`
	NumberOfValidatorsRegistered      string                             `json:"numberOfValidatorsRegistered"`
	AccountAddress                    common.Address                     `json:"accountAddress"`
	AccountAddressFormatted           string                             `json:"accountAddressFormatted"`
	OperatorId                        *big.Int                           `json:"operatorId"`
	OperatorName                      string                             `json:"operatorName"`
	OperatorActive                    bool                               `json:"operatorActive"`
	OperatorAddress                   common.Address                     `json:"operatorAddress"`
	OperatorRewardAddress             common.Address                     `json:"operatorRewardAddress"`
	OperatorRewardInETH               *big.Int                           `json:"operatorRewardInETH"`
	AccumulatedClRewards              *big.Int                           `json:"accumulatedClRewards"`
	OptedInForSocializingPool         bool                               `json:"optedInForSocializingPool"`
	SocializingPoolRewardCycleDetails types.RewardCycleDetails           `json:"socializingPoolRewardCycleDetails"`
	SocializingPoolEndTime            time.Time                          `json:"socializingPoolEndTime"`
	SocializingPoolAddress            common.Address                     `json:"socializingPoolAddress"`
	OperatorELRewardsAddress          common.Address                     `json:"operatorELRewardsAddress"`
	OperatorELRewardsAddressBalance   *big.Int                           `json:"operatorELRewardsAddressBalance"`
	DepositedSdCollateral             *big.Int                           `json:"depositedSdCollateral"`
	SdCollateralRequestedToWithdraw   *big.Int                           `json:"sdCollateralRequestedToWithdraw"`
	SdCollateralWithdrawTime          *big.Int                           `json:"withdrawTime"`
	SdCollateralWorthValidators       *big.Int                           `json:"sdCollateralWorthValidators"`
	Registered                        bool                               `json:"registered"`
	AccountBalances                   tokens.Balances                    `json:"accountBalances"`
	TotalNonTerminalValidators        *big.Int                           `json:"nonTerminalValidators"`
	ValidatorInfos                    []stdr.ValidatorInfo               `json:"validatorInfos"`
	ClaimedSocializingPoolMerkles     []stader_backend.CycleMerkleProofs `json:"claimedSocializingPoolMerkles"`
	UnclaimedSocializingPoolMerkles   []stader_backend.CycleMerkleProofs `json:"unclaimedSocializingPoolMerkles"`
}

type CanRegisterNodeResponse struct {
	Status                    string         `json:"status"`
	Error                     string         `json:"error"`
	AlreadyRegistered         bool           `json:"alreadyRegistered"`
	RegistrationPaused        bool           `json:"registrationPaused"`
	OperatorNameTooLong       bool           `json:"operatorNameTooLong"`
	OperatorRewardAddressZero bool           `json:"operatorRewardAddressZero"`
	GasInfo                   stader.GasInfo `json:"gasInfo"`
}

type RegisterNodeResponse struct {
	Status string      `json:"status"`
	Error  string      `json:"error"`
	TxHash common.Hash `json:"txHash"`
}

type CanNodeDepositSdResponse struct {
	Status                   string         `json:"status"`
	Error                    string         `json:"error"`
	CollateralContractPaused bool           `json:"collateralContractPaused"`
	InsufficientBalance      bool           `json:"insufficientBalance"`
	GasInfo                  stader.GasInfo `json:"gasInfo"`
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
	Status                   string         `json:"status"`
	Error                    string         `json:"error"`
	CanDeposit               bool           `json:"CanDeposit"`
	InsufficientBalance      bool           `json:"insufficientBalance"`
	InvalidAmount            bool           `json:"invalidAmount"`
	DepositPaused            bool           `json:"depositPaused"`
	NotEnoughSdCollateral    bool           `json:"notEnoughSdCollateral"`
	MaxValidatorLimitReached bool           `json:"maxValidatorLimitReached"`
	GasInfo                  stader.GasInfo `json:"gasInfo"`
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
	SocializingPoolContract    common.Address `json:"socializingPoolContract"`
	PermisionlessPool          common.Address `json:"permisionlessPool"`
	StaderOracle               common.Address `json:"staderOracle"`
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
	ValidatorNotRegisteredWithStader     bool   `json:"validatorNotRegisteredWithStader"`
	ValidatorNotRegisteredWithOperator   bool   `json:"validatorNotRegisteredWithOperator"`
	ValidatorNotRegistered               bool   `json:"validatorNotRegistered"`
	ValidatorPreSignKeyAlreadyRegistered bool   `json:"validatorPreSignKeyAlreadyRegistered"`
	ValidatorIsNotActive                 bool   `json:"validatorIsNotActive"`
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
	ValidatorTooYoung      bool   `json:"validatorTooYoung"`
	ValidatorExiting       bool   `json:"validatorExiting"`
}

type ExitValidatorResponse struct {
	Status string `json:"status"`
	Error  string `json:"error"`
}

type CanUpdateSocializeElResponse struct {
	Status                        string         `json:"status"`
	Error                         string         `json:"error"`
	SocializingPoolContractPaused bool           `json:"socializingPoolContractPaused"`
	AlreadyOptedIn                bool           `json:"alreadyOptedIn"`
	AlreadyOptedOut               bool           `json:"alreadyOptedOut"`
	InCooldown                    bool           `json:"inCooldown"`
	GasInfo                       stader.GasInfo `json:"gasInfo"`
}

type UpdateSocializeElResponse struct {
	Status string      `json:"status"`
	Error  string      `json:"error"`
	TxHash common.Hash `json:"txHash"`
}

type CanClaimClRewardsResponse struct {
	Status              string         `json:"status"`
	Error               string         `json:"error"`
	ValidatorWithdrawn  bool           `json:"validatorWithdrawn"`
	VaultAlreadySettled bool           `json:"vaultAlreadySettled"`
	NoClRewards         bool           `json:"noClRewards"`
	TooManyClRewards    bool           `json:"tooManyClRewards"`
	ValidatorNotFound   bool           `json:"validatorNotFound"`
	GasInfo             stader.GasInfo `json:"gasInfo"`
}

type ClaimClRewardsResponse struct {
	Status                string         `json:"status"`
	Error                 string         `json:"error"`
	ClRewardsAmount       *big.Int       `json:"clRewardsAmount"`
	OperatorRewardAddress common.Address `json:"operatorRewardAddress"`
	TxHash                common.Hash    `json:"txHash"`
}

type CanSettleExitFunds struct {
	Status                 string         `json:"status"`
	Error                  string         `json:"error"`
	ValidatorNotWithdrawn  bool           `json:"validatorNotWithdrawn"`
	ValidatorNotRegistered bool           `json:"validatorNotRegistered"`
	NoEthToWithdraw        bool           `json:"notEthToWithdraw"`
	VaultAlreadySettled    bool           `json:"vaultAlreadySettled"`
	GasInfo                stader.GasInfo `json:"gasInfo"`
}

type SettleExitFunds struct {
	Status                string         `json:"status"`
	Error                 string         `json:"error"`
	ExitAmount            *big.Int       `json:"exitShare"`
	OperatorRewardAddress common.Address `json:"operatorRewardAddress"`
	TxHash                common.Hash    `json:"txHash"`
}

type CanClaimElRewardsResponse struct {
	Status      string         `json:"status"`
	Error       string         `json:"error"`
	NoElRewards bool           `json:"noElRewards"`
	GasInfo     stader.GasInfo `json:"gasInfo"`
}

type ClaimElRewardsResponse struct {
	Status                string         `json:"status"`
	Error                 string         `json:"error"`
	ElRewardsAmount       *big.Int       `json:"elRewardsAmount"`
	OperatorRewardAddress common.Address `json:"operatorRewardAddress"`
	TxHash                common.Hash    `json:"txHash"`
}

type CanRequestWithdrawSdResponse struct {
	Status                     string         `json:"status"`
	Error                      string         `json:"error"`
	InsufficientSdCollateral   bool           `json:"insufficientSdCollateral"`
	InsufficientWithdrawableSd bool           `json:"insufficientWithdrawableSd"`
	GasInfo                    stader.GasInfo `json:"gasInfo"`
}

type RequestWithdrawSdResponse struct {
	Status string      `json:"status"`
	Error  string      `json:"error"`
	TxHash common.Hash `json:"txHash"`
}

type CanClaimSdResponse struct {
	Status                   string         `json:"status"`
	Error                    string         `json:"error"`
	NoExistingClaim          bool           `json:"noExistingClaim"`
	ClaimIsInUnbondingPeriod bool           `json:"claimIsInUnbondingPeriod"`
	GasInfo                  stader.GasInfo `json:"gasInfo"`
}

type ClaimSdResponse struct {
	Status string      `json:"status"`
	Error  string      `json:"error"`
	TxHash common.Hash `json:"txHash"`
}

type CanDownloadSpMerkleProofsResponse struct {
	Status          string  `json:"status"`
	Error           string  `json:"error"`
	NoMissingCycles bool    `json:"noMissingCycles"`
	MissingCycles   []int64 `json:"missingCycles"`
	CurrentCycle    int64   `json:"currentCycle"`
}

type DownloadSpMerkleProofsResponse struct {
	Status           string  `json:"status"`
	Error            string  `json:"error"`
	DownloadedCycles []int64 `json:"downloadedCycles"`
}

type CanClaimSpRewardsResponse struct {
	Status                        string     `json:"status"`
	Error                         string     `json:"error"`
	SocializingPoolContractPaused bool       `json:"socializingPoolContractPaused"`
	IneligibleCycles              []*big.Int `json:"ineligibleCycles"`
	ClaimedCycles                 []*big.Int `json:"claimedCycles"`
	UnclaimedCycles               []*big.Int `json:"unclaimedCycles"`
	CyclesToDownload              []*big.Int `json:"cyclesToDownload"`
}

type EstimateClaimSpRewardsGasResponse struct {
	Status  string         `json:"status"`
	Error   string         `json:"error"`
	GasInfo stader.GasInfo `json:"gasInfo"`
}

type ClaimSpRewardsResponse struct {
	Status string      `json:"status"`
	Error  string      `json:"error"`
	TxHash common.Hash `json:"txHash"`
}

type CanUpdateOperatorDetails struct {
	Status                    string         `json:"status"`
	Error                     string         `json:"error"`
	OperatorNameTooLong       bool           `json:"operatorNameTooLong"`
	OperatorRewardAddressZero bool           `json:"operatorRewardAddressZero"`
	NothingToUpdate           bool           `json:"nothingToUpdate"`
	GasInfo                   stader.GasInfo `json:"gasInfo"`
}

type UpdateOperatorDetails struct {
	Status string      `json:"status"`
	Error  string      `json:"error"`
	TxHash common.Hash `json:"txHash"`
}

type NodeSignResponse struct {
	Status     string `json:"status"`
	Error      string `json:"error"`
	SignedData string `json:"signedData"`
}
