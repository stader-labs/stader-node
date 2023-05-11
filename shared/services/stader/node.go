/*
This work is licensed and released under GNU GPL v3 or any other later versions.
The full text of the license is below/ found at <http://www.gnu.org/licenses/>

(c) 2023 Rocket Pool Pty Ltd. Modified under GNU GPL v3. [0.4.0-beta]

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
package stader

import (
	"encoding/json"
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	string_utils "github.com/stader-labs/stader-node/shared/utils/string-utils"
	"github.com/stader-labs/stader-node/stader-lib/types"
	"math/big"
	"strconv"

	"github.com/stader-labs/stader-node/shared/types/api"
)

// Get node status
func (c *Client) NodeStatus() (api.NodeStatusResponse, error) {
	responseBytes, err := c.callAPI("node status")
	if err != nil {
		return api.NodeStatusResponse{}, fmt.Errorf("could not get node status: %w", err)
	}
	var response api.NodeStatusResponse
	if err := json.Unmarshal(responseBytes, &response); err != nil {
		return api.NodeStatusResponse{}, fmt.Errorf("could not decode node status response: %w", err)
	}
	if response.Error != "" {
		return api.NodeStatusResponse{}, fmt.Errorf("could not get node status: %s", response.Error)
	}
	if response.AccountBalances.ETH == nil {
		response.AccountBalances.ETH = big.NewInt(0)
	}
	if response.AccountBalances.Sd == nil {
		response.AccountBalances.Sd = big.NewInt(0)
	}
	if response.OperatorELRewardsAddressBalance == nil {
		response.OperatorELRewardsAddressBalance = big.NewInt(0)
	}
	if response.OperatorRewardInETH == nil {
		response.OperatorRewardInETH = big.NewInt(0)
	}
	if response.SdCollateralWithdrawTime == nil {
		response.SdCollateralWithdrawTime = big.NewInt(0)
	}
	if response.SdCollateralRequestedToWithdraw == nil {
		response.SdCollateralRequestedToWithdraw = big.NewInt(0)
	}
	if response.TotalNonTerminalValidators == nil {
		response.TotalNonTerminalValidators = big.NewInt(0)
	}

	return response, nil
}

// Check whether the node can be registered
func (c *Client) CanRegisterNode(operatorName string, operatorRewardAddress common.Address, socializeMev bool) (api.CanRegisterNodeResponse, error) {
	responseBytes, err := c.callAPI("node can-register", operatorName, operatorRewardAddress.Hex(), strconv.FormatBool(socializeMev))
	if err != nil {
		return api.CanRegisterNodeResponse{}, fmt.Errorf("could not get can register node status: %w", err)
	}
	var response api.CanRegisterNodeResponse
	if err := json.Unmarshal(responseBytes, &response); err != nil {
		return api.CanRegisterNodeResponse{}, fmt.Errorf("could not decode can register node response: %w", err)
	}
	if response.Error != "" {
		return api.CanRegisterNodeResponse{}, fmt.Errorf("could not get can register node status: %s", response.Error)
	}
	return response, nil
}

// Register the node
func (c *Client) RegisterNode(operatorName string, operatorRewardAddress common.Address, socializeMev bool) (api.RegisterNodeResponse, error) {
	responseBytes, err := c.callAPI("node register", operatorName, operatorRewardAddress.Hex(), strconv.FormatBool(socializeMev))
	if err != nil {
		return api.RegisterNodeResponse{}, fmt.Errorf("could not register node: %w", err)
	}
	var response api.RegisterNodeResponse
	if err := json.Unmarshal(responseBytes, &response); err != nil {
		return api.RegisterNodeResponse{}, fmt.Errorf("could not decode register node response: %w", err)
	}
	if response.Error != "" {
		return api.RegisterNodeResponse{}, fmt.Errorf("could not register node: %s", response.Error)
	}
	return response, nil
}

// Check whether the node can deposit SD
func (c *Client) CanNodeDepositSd(amountWei *big.Int) (api.CanNodeDepositSdResponse, error) {
	responseBytes, err := c.callAPI(fmt.Sprintf("node can-node-deposit-sd %s", amountWei.String()))
	if err != nil {
		return api.CanNodeDepositSdResponse{}, fmt.Errorf("could not get can node deposit SD status: %w", err)
	}
	var response api.CanNodeDepositSdResponse
	if err := json.Unmarshal(responseBytes, &response); err != nil {
		return api.CanNodeDepositSdResponse{}, fmt.Errorf("could not decode can node deposit SD response: %w", err)
	}
	if response.Error != "" {
		return api.CanNodeDepositSdResponse{}, fmt.Errorf("could not get can node deposit SD status: %s", response.Error)
	}
	return response, nil
}

// Get the gas estimate for approving new SD interaction
func (c *Client) NodeDepositSdApprovalGas(amountWei *big.Int) (api.NodeDepositSdApproveGasResponse, error) {
	responseBytes, err := c.callAPI(fmt.Sprintf("node get-deposit-sd-approval-gas %s", amountWei.String()))
	if err != nil {
		return api.NodeDepositSdApproveGasResponse{}, fmt.Errorf("could not get new SD approval gas: %w", err)
	}
	var response api.NodeDepositSdApproveGasResponse
	if err := json.Unmarshal(responseBytes, &response); err != nil {
		return api.NodeDepositSdApproveGasResponse{}, fmt.Errorf("could not decode node deposit S approve gas response: %w", err)
	}
	if response.Error != "" {
		return api.NodeDepositSdApproveGasResponse{}, fmt.Errorf("could not get new SD approval gas: %s", response.Error)
	}
	return response, nil
}

// Approve SD for depositing as collateral
func (c *Client) NodeDepositSdApprove(amountWei *big.Int) (api.NodeDepositSdApproveResponse, error) {
	responseBytes, err := c.callAPI(fmt.Sprintf("node deposit-sd-approve-sd %s", amountWei.String()))
	if err != nil {
		return api.NodeDepositSdApproveResponse{}, fmt.Errorf("could not approve SD for staking: %w", err)
	}
	var response api.NodeDepositSdApproveResponse
	if err := json.Unmarshal(responseBytes, &response); err != nil {
		return api.NodeDepositSdApproveResponse{}, fmt.Errorf("could not decode deposit node SD approve response: %w", err)
	}
	if response.Error != "" {
		return api.NodeDepositSdApproveResponse{}, fmt.Errorf("could not approve SD for staking: %s", response.Error)
	}
	return response, nil
}

// Deposit SD as collateral
func (c *Client) NodeDepositSd(amountWei *big.Int) (api.NodeDepositSdResponse, error) {
	responseBytes, err := c.callAPI(fmt.Sprintf("node deposit-sd %s", amountWei.String()))
	if err != nil {
		return api.NodeDepositSdResponse{}, fmt.Errorf("could not deposit node SD: %w", err)
	}
	var response api.NodeDepositSdResponse
	if err := json.Unmarshal(responseBytes, &response); err != nil {
		return api.NodeDepositSdResponse{}, fmt.Errorf("could not decode deposit node SD response: %w", err)
	}
	if response.Error != "" {
		return api.NodeDepositSdResponse{}, fmt.Errorf("could not deposit node SD: %s", response.Error)
	}
	return response, nil
}

func (c *Client) CanUpdateSocializeEl(socializeEl bool) (api.CanUpdateSocializeElResponse, error) {
	responseBytes, err := c.callAPI(fmt.Sprintf("node can-update-socialize-el %s", strconv.FormatBool(socializeEl)))
	if err != nil {
		return api.CanUpdateSocializeElResponse{}, fmt.Errorf("could not opt-in or opt-out of socializing pool: %w", err)
	}
	var response api.CanUpdateSocializeElResponse
	if err := json.Unmarshal(responseBytes, &response); err != nil {
		return api.CanUpdateSocializeElResponse{}, fmt.Errorf("could not decode opt-in or opt-out of socializing pool response: %w", err)
	}
	if response.Error != "" {
		return api.CanUpdateSocializeElResponse{}, fmt.Errorf("could not opt-in ot opt-out of socializing pool: %s", response.Error)
	}
	return response, nil
}

func (c *Client) UpdateSocializeEl(socializeEl bool) (api.UpdateSocializeElResponse, error) {
	responseBytes, err := c.callAPI(fmt.Sprintf("node update-socialize-el %s", strconv.FormatBool(socializeEl)))
	if err != nil {
		return api.UpdateSocializeElResponse{}, fmt.Errorf("could not opt-in or opt-out of socializing pool: %w", err)
	}
	var response api.UpdateSocializeElResponse
	if err := json.Unmarshal(responseBytes, &response); err != nil {
		return api.UpdateSocializeElResponse{}, fmt.Errorf("could not decode opt-in or opt-out of socializing pool response: %w", err)
	}
	if response.Error != "" {
		return api.UpdateSocializeElResponse{}, fmt.Errorf("could not opt-in ot opt-out of socializing pool: %s", response.Error)
	}
	return response, nil
}

// Get the node's SD allowance
func (c *Client) GetNodeDepositSdAllowance() (api.NodeDepositSdAllowanceResponse, error) {
	responseBytes, err := c.callAPI(fmt.Sprintf("node deposit-sd-allowance"))
	if err != nil {
		return api.NodeDepositSdAllowanceResponse{}, fmt.Errorf("could not get node deposit SD allowance: %w", err)
	}
	var response api.NodeDepositSdAllowanceResponse
	if err := json.Unmarshal(responseBytes, &response); err != nil {
		return api.NodeDepositSdAllowanceResponse{}, fmt.Errorf("could not decode node deposit SD allowance response: %w", err)
	}
	if response.Error != "" {
		return api.NodeDepositSdAllowanceResponse{}, fmt.Errorf("could not get node deposit SD allowance: %s", response.Error)
	}
	return response, nil
}

// Check whether the node can make a deposit
func (c *Client) CanNodeDeposit(amountWei *big.Int, salt *big.Int, numValidators *big.Int, reloadKeys bool) (api.CanNodeDepositResponse, error) {
	responseBytes, err := c.callAPI(fmt.Sprintf("node can-deposit %s %s %d %t", amountWei.String(), salt.String(), numValidators, reloadKeys))
	if err != nil {
		return api.CanNodeDepositResponse{}, fmt.Errorf("could not get can node deposit status: %w", err)
	}
	var response api.CanNodeDepositResponse
	if err := json.Unmarshal(responseBytes, &response); err != nil {
		return api.CanNodeDepositResponse{}, fmt.Errorf("could not decode can node deposit response: %w", err)
	}
	if response.Error != "" {
		return api.CanNodeDepositResponse{}, fmt.Errorf("could not get can node deposit status: %s", response.Error)
	}
	return response, nil
}

func (c *Client) CanExitValidator(validatorPubKey types.ValidatorPubkey) (api.CanExitValidatorResponse, error) {
	responseBytes, err := c.callAPI(fmt.Sprintf("node can-exit-validator %s", validatorPubKey))
	if err != nil {
		return api.CanExitValidatorResponse{}, fmt.Errorf("could not get can-exit-validator status: %w", err)
	}
	var response api.CanExitValidatorResponse
	if err := json.Unmarshal(responseBytes, &response); err != nil {
		return api.CanExitValidatorResponse{}, fmt.Errorf("could not decode can-exit-validator response: %w", err)
	}
	if response.Error != "" {
		return api.CanExitValidatorResponse{}, fmt.Errorf("could not get can-exit-validator status: %s", response.Error)
	}
	return response, nil
}

func (c *Client) ExitValidator(validatorPubKey types.ValidatorPubkey) (api.ExitValidatorResponse, error) {
	responseBytes, err := c.callAPI(fmt.Sprintf("node exit-validator %s", validatorPubKey))
	if err != nil {
		return api.ExitValidatorResponse{}, fmt.Errorf("could not get exit-validator status: %w", err)
	}
	var response api.ExitValidatorResponse
	if err := json.Unmarshal(responseBytes, &response); err != nil {
		return api.ExitValidatorResponse{}, fmt.Errorf("could not decode exit-validator response: %w", err)
	}
	if response.Error != "" {
		return api.ExitValidatorResponse{}, fmt.Errorf("could not get exit-validator status: %s", response.Error)
	}
	return response, nil
}

func (c *Client) GetContractsInfo() (api.ContractsInfoResponse, error) {
	responseBytes, err := c.callAPI("node get-contracts-info")
	if err != nil {
		return api.ContractsInfoResponse{}, fmt.Errorf("could not get networks contracts info: %w", err)
	}
	var response api.ContractsInfoResponse
	if err := json.Unmarshal(responseBytes, &response); err != nil {
		return api.ContractsInfoResponse{}, fmt.Errorf("could not decode networks contracts info: %w", err)
	}
	if response.Error != "" {
		return api.ContractsInfoResponse{}, fmt.Errorf("could not get networks contract info: %s", response.Error)
	}
	return response, nil
}

// Make a node deposit
func (c *Client) NodeDeposit(amountWei *big.Int, salt *big.Int, numValidators *big.Int, reloadKeys bool) (api.NodeDepositResponse, error) {
	responseBytes, err := c.callAPI(fmt.Sprintf("node deposit %s %s %d %t", amountWei.String(), salt.String(), numValidators, reloadKeys))
	if err != nil {
		return api.NodeDepositResponse{}, fmt.Errorf("could not make node deposit as er: %w", err)
	}
	var response api.NodeDepositResponse
	if err := json.Unmarshal(responseBytes, &response); err != nil {
		return api.NodeDepositResponse{}, fmt.Errorf("could not decode node deposit response: %w", err)
	}
	if response.Error != "" {
		return api.NodeDepositResponse{}, fmt.Errorf("could not make node deposit: %s", response.Error)
	}
	return response, nil
}

// Check whether the node can send tokens
func (c *Client) CanNodeSend(amountWei *big.Int, token string) (api.CanNodeSendResponse, error) {
	responseBytes, err := c.callAPI(fmt.Sprintf("node can-send %s %s", amountWei.String(), token))
	if err != nil {
		return api.CanNodeSendResponse{}, fmt.Errorf("could not get can node send status: %w", err)
	}
	var response api.CanNodeSendResponse
	if err := json.Unmarshal(responseBytes, &response); err != nil {
		return api.CanNodeSendResponse{}, fmt.Errorf("could not decode can node send response: %w", err)
	}
	if response.Error != "" {
		return api.CanNodeSendResponse{}, fmt.Errorf("could not get can node send status: %s", response.Error)
	}
	return response, nil
}

// Send tokens from the node to an address
func (c *Client) NodeSend(amountWei *big.Int, token string, toAddress common.Address) (api.NodeSendResponse, error) {
	responseBytes, err := c.callAPI(fmt.Sprintf("node send %s %s %s", amountWei.String(), token, toAddress.Hex()))
	if err != nil {
		return api.NodeSendResponse{}, fmt.Errorf("could not send tokens from node: %w", err)
	}
	var response api.NodeSendResponse
	if err := json.Unmarshal(responseBytes, &response); err != nil {
		return api.NodeSendResponse{}, fmt.Errorf("could not decode node send response: %w", err)
	}
	if response.Error != "" {
		return api.NodeSendResponse{}, fmt.Errorf("could not send tokens from node: %s", response.Error)
	}
	return response, nil
}

// Get node sync progress
func (c *Client) NodeSync() (api.NodeSyncProgressResponse, error) {
	responseBytes, err := c.callAPI("node sync")
	if err != nil {
		return api.NodeSyncProgressResponse{}, fmt.Errorf("could not get node sync: %w", err)
	}
	var response api.NodeSyncProgressResponse
	if err := json.Unmarshal(responseBytes, &response); err != nil {
		return api.NodeSyncProgressResponse{}, fmt.Errorf("could not decode node sync response: %w", err)
	}
	if response.Error != "" {
		return api.NodeSyncProgressResponse{}, fmt.Errorf("could not get node sync: %s", response.Error)
	}
	return response, nil
}

// Use the node private key to sign an arbitrary message
func (c *Client) SignMessage(message string) (api.NodeSignResponse, error) {
	responseBytes, err := c.callAPI("node sign-message", message)
	if err != nil {
		return api.NodeSignResponse{}, fmt.Errorf("could not sign message: %w", err)
	}

	var response api.NodeSignResponse
	if err := json.Unmarshal(responseBytes, &response); err != nil {
		return api.NodeSignResponse{}, fmt.Errorf("could not decode node sign response: %w", err)
	}
	if response.Error != "" {
		return api.NodeSignResponse{}, fmt.Errorf("could not sign message: %s", response.Error)
	}
	return response, nil
}

func (c *Client) CanClaimElRewards() (api.CanClaimElRewardsResponse, error) {
	responseBytes, err := c.callAPI("node can-claim-el-rewards")
	if err != nil {
		return api.CanClaimElRewardsResponse{}, fmt.Errorf("could not get node can-claim-el-rewards response: %w", err)
	}
	var response api.CanClaimElRewardsResponse
	if err := json.Unmarshal(responseBytes, &response); err != nil {
		return api.CanClaimElRewardsResponse{}, fmt.Errorf("could not decode node can-claim-el-rewards response: %w", err)
	}
	if response.Error != "" {
		return api.CanClaimElRewardsResponse{}, fmt.Errorf("could not get node can-claim-el-rewards response: %s", response.Error)
	}

	return response, nil
}

func (c *Client) ClaimElRewards() (api.ClaimElRewardsResponse, error) {
	responseBytes, err := c.callAPI("node claim-el-rewards")
	if err != nil {
		return api.ClaimElRewardsResponse{}, fmt.Errorf("could not get node claim-el-rewards response: %w", err)
	}
	var response api.ClaimElRewardsResponse
	if err := json.Unmarshal(responseBytes, &response); err != nil {
		return api.ClaimElRewardsResponse{}, fmt.Errorf("could not decode node claim-el-rewards response: %w", err)
	}
	if response.Error != "" {
		return api.ClaimElRewardsResponse{}, fmt.Errorf("could not get node claim-el-rewards response: %s", response.Error)
	}

	return response, nil
}

func (c *Client) CanClaimClRewards(validatorPubKey types.ValidatorPubkey) (api.CanClaimClRewardsResponse, error) {
	responseBytes, err := c.callAPI(fmt.Sprintf("node can-claim-cl-rewards %s", validatorPubKey))
	if err != nil {
		return api.CanClaimClRewardsResponse{}, fmt.Errorf("could not get node can-claim-cl-rewards response: %w", err)
	}
	var response api.CanClaimClRewardsResponse
	if err := json.Unmarshal(responseBytes, &response); err != nil {
		return api.CanClaimClRewardsResponse{}, fmt.Errorf("could not decode node can-claim-cl-rewards response: %w", err)
	}
	if response.Error != "" {
		return api.CanClaimClRewardsResponse{}, fmt.Errorf("could not get node can-claim-cl-rewards response: %s", response.Error)
	}

	return response, nil
}

func (c *Client) ClaimClRewards(validatorPubKey types.ValidatorPubkey) (api.ClaimClRewardsResponse, error) {
	responseBytes, err := c.callAPI(fmt.Sprintf("node claim-cl-rewards %s", validatorPubKey))
	if err != nil {
		return api.ClaimClRewardsResponse{}, fmt.Errorf("could not get node claim-cl-rewards response: %w", err)
	}
	var response api.ClaimClRewardsResponse
	if err := json.Unmarshal(responseBytes, &response); err != nil {
		return api.ClaimClRewardsResponse{}, fmt.Errorf("could not decode node claim-cl-rewards response: %w", err)
	}
	if response.Error != "" {
		return api.ClaimClRewardsResponse{}, fmt.Errorf("could not get node claim-cl-rewards response: %s", response.Error)
	}

	return response, nil
}

func (c *Client) CanRequestSdCollateralWithdraw(amount *big.Int) (api.CanRequestWithdrawSdResponse, error) {
	responseBytes, err := c.callAPI(fmt.Sprintf("node can-node-request-sd-withdraw %s", amount.String()))
	if err != nil {
		return api.CanRequestWithdrawSdResponse{}, fmt.Errorf("could not get node can-node-request-sd-withdraw response: %w", err)
	}
	var response api.CanRequestWithdrawSdResponse
	if err := json.Unmarshal(responseBytes, &response); err != nil {
		return api.CanRequestWithdrawSdResponse{}, fmt.Errorf("could not decode node can-node-request-sd-withdraw response: %w", err)
	}
	if response.Error != "" {
		return api.CanRequestWithdrawSdResponse{}, fmt.Errorf("could not get node can-node-request-sd-withdraw response: %s", response.Error)
	}

	return response, nil
}

func (c *Client) RequestSdCollateralWithdraw(amount *big.Int) (api.RequestWithdrawSdResponse, error) {
	responseBytes, err := c.callAPI(fmt.Sprintf("node node-request-sd-withdraw %s", amount.String()))
	if err != nil {
		return api.RequestWithdrawSdResponse{}, fmt.Errorf("could not get node node-request-sd-withdraw response: %w", err)
	}
	var response api.RequestWithdrawSdResponse
	if err := json.Unmarshal(responseBytes, &response); err != nil {
		return api.RequestWithdrawSdResponse{}, fmt.Errorf("could not decode node node-request-sd-withdraw response: %w", err)
	}
	if response.Error != "" {
		return api.RequestWithdrawSdResponse{}, fmt.Errorf("could not get node node-request-sd-withdraw response: %s", response.Error)
	}

	return response, nil
}

func (c *Client) CanClaimSd() (api.CanClaimSdResponse, error) {
	// TODO - bchain - normalize these response messages
	responseBytes, err := c.callAPI(fmt.Sprintf("node can-node-claim-sd"))
	if err != nil {
		return api.CanClaimSdResponse{}, fmt.Errorf("could not get node can-node-claim-sd response: %w", err)
	}
	var response api.CanClaimSdResponse
	if err := json.Unmarshal(responseBytes, &response); err != nil {
		return api.CanClaimSdResponse{}, fmt.Errorf("could not decode node can-node-claim-sd response: %w", err)
	}
	if response.Error != "" {
		return api.CanClaimSdResponse{}, fmt.Errorf("could not get node can-node-claim-sd response: %s", response.Error)
	}

	return response, nil
}

func (c *Client) ClaimSd() (api.ClaimSdResponse, error) {
	// TODO - bchain - normalize these response messages
	responseBytes, err := c.callAPI(fmt.Sprintf("node node-claim-sd"))
	if err != nil {
		return api.ClaimSdResponse{}, fmt.Errorf("could not get node node-claim-sd response: %w", err)
	}
	var response api.ClaimSdResponse
	if err := json.Unmarshal(responseBytes, &response); err != nil {
		return api.ClaimSdResponse{}, fmt.Errorf("could not decode node node-claim-sd response: %w", err)
	}
	if response.Error != "" {
		return api.ClaimSdResponse{}, fmt.Errorf("could not get node node-claim-sd response: %s", response.Error)
	}

	return response, nil
}

func (c *Client) CanDownloadSpMerkleProofs() (api.CanDownloadSpMerkleProofsResponse, error) {
	responseBytes, err := c.callAPI(fmt.Sprintf("node can-download-sp-merkle-proofs"))
	if err != nil {
		return api.CanDownloadSpMerkleProofsResponse{}, fmt.Errorf("could not get node can-download-sp-merkle-proofs response: %w", err)
	}
	var response api.CanDownloadSpMerkleProofsResponse
	if err := json.Unmarshal(responseBytes, &response); err != nil {
		return api.CanDownloadSpMerkleProofsResponse{}, fmt.Errorf("could not decode node can-download-sp-merkle-proofs response: %w", err)
	}
	if response.Error != "" {
		return api.CanDownloadSpMerkleProofsResponse{}, fmt.Errorf("could not get node can-download-sp-merkle-proofs response: %s", response.Error)
	}

	return response, nil
}

func (c *Client) DownloadSpMerkleProofs() (api.DownloadSpMerkleProofsResponse, error) {
	responseBytes, err := c.callAPI(fmt.Sprintf("node download-sp-merkle-proofs"))
	if err != nil {
		return api.DownloadSpMerkleProofsResponse{}, fmt.Errorf("could not get node download-sp-merkle-proofs response: %w", err)
	}
	var response api.DownloadSpMerkleProofsResponse
	if err := json.Unmarshal(responseBytes, &response); err != nil {
		return api.DownloadSpMerkleProofsResponse{}, fmt.Errorf("could not decode node download-sp-merkle-proofs response: %w", err)
	}
	if response.Error != "" {
		return api.DownloadSpMerkleProofsResponse{}, fmt.Errorf("could not get node download-sp-merkle-proofs response: %s", response.Error)
	}

	return response, nil
}

func (c *Client) GetDetailedCyclesInfo(cycles []*big.Int) (api.CyclesDetailedInfo, error) {
	stringifiedCycleList := string_utils.StringifyArray(cycles)
	responseBytes, err := c.callAPI(fmt.Sprintf("node detailed-cycles-info %s", stringifiedCycleList))
	if err != nil {
		return api.CyclesDetailedInfo{}, fmt.Errorf("could not get node detailed-cycles-info response: %w", err)
	}
	var response api.CyclesDetailedInfo
	if err := json.Unmarshal(responseBytes, &response); err != nil {
		return api.CyclesDetailedInfo{}, fmt.Errorf("could not decode node detailed-cycles-info response: %w", err)
	}
	if response.Error != "" {
		return api.CyclesDetailedInfo{}, fmt.Errorf("could not get node detailed-cycles-info response: %s", response.Error)
	}

	return response, nil
}

func (c *Client) CanClaimSpRewards() (api.CanClaimSpRewardsResponse, error) {
	responseBytes, err := c.callAPI(fmt.Sprintf("node can-claim-sp-rewards"))
	if err != nil {
		return api.CanClaimSpRewardsResponse{}, fmt.Errorf("could not get node can-claim-sp-rewards response: %w", err)
	}
	var response api.CanClaimSpRewardsResponse
	if err := json.Unmarshal(responseBytes, &response); err != nil {
		return api.CanClaimSpRewardsResponse{}, fmt.Errorf("could not decode node can-claim-sp-rewards response: %w", err)
	}
	if response.Error != "" {
		return api.CanClaimSpRewardsResponse{}, fmt.Errorf("could not get node can-claim-sp-rewards response: %s", response.Error)
	}

	return response, nil
}

func (c *Client) EstimateClaimSpRewardsGas(cycles []*big.Int) (api.EstimateClaimSpRewardsGasResponse, error) {
	stringifiedCycleList := string_utils.StringifyArray(cycles)
	responseBytes, err := c.callAPI(fmt.Sprintf("node estimate-claim-sp-rewards-gas %s", stringifiedCycleList))
	if err != nil {
		return api.EstimateClaimSpRewardsGasResponse{}, fmt.Errorf("could not get node estimate-claim-sp-rewards-gas response: %w", err)
	}
	var response api.EstimateClaimSpRewardsGasResponse
	if err := json.Unmarshal(responseBytes, &response); err != nil {
		return api.EstimateClaimSpRewardsGasResponse{}, fmt.Errorf("could not decode node estimate-claim-sp-rewards-gas response: %w", err)
	}
	if response.Error != "" {
		return api.EstimateClaimSpRewardsGasResponse{}, fmt.Errorf("could not get node estimate-claim-sp-rewards-gas response: %s", response.Error)
	}

	return response, nil
}

func (c *Client) ClaimSpRewards(cycles []*big.Int) (api.ClaimSpRewardsResponse, error) {
	stringifiedCycleList := string_utils.StringifyArray(cycles)
	responseBytes, err := c.callAPI(fmt.Sprintf("node claim-sp-rewards %s", stringifiedCycleList))
	if err != nil {
		return api.ClaimSpRewardsResponse{}, fmt.Errorf("could not get node claim-sp-rewards response: %w", err)
	}
	var response api.ClaimSpRewardsResponse
	if err := json.Unmarshal(responseBytes, &response); err != nil {
		return api.ClaimSpRewardsResponse{}, fmt.Errorf("could not decode node claim-sp-rewards response: %w", err)
	}
	if response.Error != "" {
		return api.ClaimSpRewardsResponse{}, fmt.Errorf("could not get node claim-sp-rewards response: %s", response.Error)
	}

	return response, nil
}

func (c *Client) CanUpdateOperatorDetails(operatorName string, operatorRewardAddress common.Address) (api.CanUpdateOperatorDetails, error) {
	responseBytes, err := c.callAPI("node can-update-operator-details", operatorName, operatorRewardAddress.Hex())
	if err != nil {
		return api.CanUpdateOperatorDetails{}, fmt.Errorf("could not get can-update-operator-details response: %w", err)
	}
	var response api.CanUpdateOperatorDetails
	if err := json.Unmarshal(responseBytes, &response); err != nil {
		return api.CanUpdateOperatorDetails{}, fmt.Errorf("could not decode can-update-operator-details response: %w", err)
	}
	if response.Error != "" {
		return api.CanUpdateOperatorDetails{}, fmt.Errorf("could not get can-update-operator-details response: %s", response.Error)
	}
	return response, nil
}

func (c *Client) UpdateOperatorDetails(operatorName string, operatorRewardAddress common.Address) (api.UpdateOperatorDetails, error) {
	responseBytes, err := c.callAPI("node update-operator-details", operatorName, operatorRewardAddress.Hex())
	if err != nil {
		return api.UpdateOperatorDetails{}, fmt.Errorf("could not get update-operator-details response: %w", err)
	}
	var response api.UpdateOperatorDetails
	if err := json.Unmarshal(responseBytes, &response); err != nil {
		return api.UpdateOperatorDetails{}, fmt.Errorf("could not decode update-operator-details response: %w", err)
	}
	if response.Error != "" {
		return api.UpdateOperatorDetails{}, fmt.Errorf("could not get update-operator-details response: %s", response.Error)
	}
	return response, nil
}
