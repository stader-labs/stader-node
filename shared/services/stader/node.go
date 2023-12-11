/*
This work is licensed and released under GNU GPL v3 or any other later versions.
The full text of the license is below/ found at <http://www.gnu.org/licenses/>

(c) 2023 Rocket Pool Pty Ltd. Modified under GNU GPL v3. [1.4.3]

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
	"math/big"
	"strconv"

	"github.com/ethereum/go-ethereum/common"
	string_utils "github.com/stader-labs/stader-node/shared/utils/string-utils"
	"github.com/stader-labs/stader-node/stader-lib/types"

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
func (c *Client) NodeSdApprovalGas(amountWei *big.Int, address common.Address) (api.SdApproveGasResponse, error) {
	responseBytes, err := c.callAPI(fmt.Sprintf("node get-sd-approval-gas %s %s", amountWei.String(), address))
	if err != nil {
		return api.SdApproveGasResponse{}, fmt.Errorf("could not get new SD approval gas: %w", err)
	}
	var response api.SdApproveGasResponse
	if err := json.Unmarshal(responseBytes, &response); err != nil {
		return api.SdApproveGasResponse{}, fmt.Errorf("could not decode node deposit S approve gas response: %w", err)
	}
	if response.Error != "" {
		return api.SdApproveGasResponse{}, fmt.Errorf("could not get new SD approval gas: %s", response.Error)
	}
	return response, nil
}

// Approve SD for depositing as collateral
func (c *Client) NodeSdApprove(amountWei *big.Int, address common.Address) (api.SdApproveResponse, error) {
	responseBytes, err := c.callAPI(fmt.Sprintf("node sd-approve-sd %s %s", amountWei.String(), address))
	if err != nil {
		return api.SdApproveResponse{}, fmt.Errorf("could not approve SD for staking: %w", err)
	}
	var response api.SdApproveResponse
	if err := json.Unmarshal(responseBytes, &response); err != nil {
		return api.SdApproveResponse{}, fmt.Errorf("could not decode deposit node SD approve response: %w", err)
	}
	if response.Error != "" {
		return api.SdApproveResponse{}, fmt.Errorf("could not approve SD for staking: %s", response.Error)
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
func (c *Client) GetNodeSdAllowance(contractAddress common.Address) (api.SdAllowanceResponse, error) {
	responseBytes, err := c.callAPI(fmt.Sprintf("node sd-allowance %s", contractAddress.String()))
	if err != nil {
		return api.SdAllowanceResponse{}, fmt.Errorf("could not get node deposit SD allowance: %w", err)
	}
	var response api.SdAllowanceResponse
	if err := json.Unmarshal(responseBytes, &response); err != nil {
		return api.SdAllowanceResponse{}, fmt.Errorf("could not decode node deposit SD allowance response: %w", err)
	}
	if response.Error != "" {
		return api.SdAllowanceResponse{}, fmt.Errorf("could not get node deposit SD allowance: %s", response.Error)
	}
	return response, nil
}

// Check whether the node can make a deposit
func (c *Client) CanNodeDeposit(amountBasedWei, amountUtilityWei, numValidators *big.Int, reloadKeys bool) (api.CanNodeDepositResponse, error) {
	responseBytes, err := c.callAPI(fmt.Sprintf("validator can-deposit %s %s %s %t", amountBasedWei.String(), amountUtilityWei.String(), numValidators, reloadKeys))
	if err != nil {
		return api.CanNodeDepositResponse{}, fmt.Errorf("could not get can validator deposit status: %w", err)
	}
	var response api.CanNodeDepositResponse
	if err := json.Unmarshal(responseBytes, &response); err != nil {
		return api.CanNodeDepositResponse{}, fmt.Errorf("could not decode can validator deposit response: %w", err)
	}
	if response.Error != "" {
		return api.CanNodeDepositResponse{}, fmt.Errorf("could not get can validator deposit status: %s", response.Error)
	}
	return response, nil
}

func (c *Client) CanExitValidator(validatorPubKey types.ValidatorPubkey) (api.CanExitValidatorResponse, error) {
	responseBytes, err := c.callAPI(fmt.Sprintf("validator can-exit-validator %s", validatorPubKey))
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
	responseBytes, err := c.callAPI(fmt.Sprintf("validator exit-validator %s", validatorPubKey))
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
func (c *Client) NodeDeposit(amountWei, numValidators, utilitySDAmount *big.Int, reloadKeys bool) (api.NodeDepositResponse, error) {
	responseBytes, err := c.callAPI(fmt.Sprintf("validator deposit %s %s %s %t", amountWei.String(), utilitySDAmount.String(), numValidators, reloadKeys))
	if err != nil {
		return api.NodeDepositResponse{}, fmt.Errorf("could not make validator deposit as er: %w", err)
	}
	var response api.NodeDepositResponse
	if err := json.Unmarshal(responseBytes, &response); err != nil {
		return api.NodeDepositResponse{}, fmt.Errorf("could not decode validator deposit response: %w", err)
	}
	if response.Error != "" {
		return api.NodeDepositResponse{}, fmt.Errorf("could not make validator deposit: %s", response.Error)
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

func (c *Client) CanSendElRewards() (api.CanSendElRewardsResponse, error) {
	responseBytes, err := c.callAPI("node can-send-el-rewards")
	if err != nil {
		return api.CanSendElRewardsResponse{}, fmt.Errorf("could not get node can-send-el-rewards response: %w", err)
	}
	var response api.CanSendElRewardsResponse
	if err := json.Unmarshal(responseBytes, &response); err != nil {
		return api.CanSendElRewardsResponse{}, fmt.Errorf("could not decode node can-send-el-rewards response: %w", err)
	}
	if response.Error != "" {
		return api.CanSendElRewardsResponse{}, fmt.Errorf("could not get node can-send-el-rewards response: %s", response.Error)
	}

	return response, nil
}

func (c *Client) SendElRewards() (api.SendElRewardsResponse, error) {
	responseBytes, err := c.callAPI("node send-el-rewards")
	if err != nil {
		return api.SendElRewardsResponse{}, fmt.Errorf("could not get node send-el-rewards response: %w", err)
	}
	var response api.SendElRewardsResponse
	if err := json.Unmarshal(responseBytes, &response); err != nil {
		return api.SendElRewardsResponse{}, fmt.Errorf("could not decode node send-el-rewards response: %w", err)
	}
	if response.Error != "" {
		return api.SendElRewardsResponse{}, fmt.Errorf("could not get node send-el-rewards response: %s", response.Error)
	}

	return response, nil
}

func (c *Client) CanClaimRewards() (api.CanClaimRewards, error) {
	responseBytes, err := c.callAPI("node can-claim-rewards")
	if err != nil {
		return api.CanClaimRewards{}, fmt.Errorf("could not get node can-claim-rewards response: %w", err)
	}
	var response api.CanClaimRewards
	if err := json.Unmarshal(responseBytes, &response); err != nil {
		return api.CanClaimRewards{}, fmt.Errorf("could not decode node can-claim-rewards response: %w", err)
	}
	if response.Error != "" {
		return api.CanClaimRewards{}, fmt.Errorf("could not get node can-claim-rewards response: %s", response.Error)
	}

	return response, nil
}

func (c *Client) ClaimRewards() (api.ClaimRewards, error) {
	responseBytes, err := c.callAPI("node claim-rewards")
	if err != nil {
		return api.ClaimRewards{}, fmt.Errorf("could not get node claim-rewards response: %w", err)
	}
	var response api.ClaimRewards
	if err := json.Unmarshal(responseBytes, &response); err != nil {
		return api.ClaimRewards{}, fmt.Errorf("could not decode node claim-rewards response: %w", err)
	}
	if response.Error != "" {
		return api.ClaimRewards{}, fmt.Errorf("could not get node claim-rewards response: %s", response.Error)
	}

	return response, nil
}

func (c *Client) CanSendClRewards(validatorPubKey types.ValidatorPubkey) (api.CanSendClRewardsResponse, error) {
	responseBytes, err := c.callAPI(fmt.Sprintf("validator can-send-cl-rewards %s", validatorPubKey))
	if err != nil {
		return api.CanSendClRewardsResponse{}, fmt.Errorf("could not get validator can-send-cl-rewards response: %w", err)
	}
	var response api.CanSendClRewardsResponse
	if err := json.Unmarshal(responseBytes, &response); err != nil {
		return api.CanSendClRewardsResponse{}, fmt.Errorf("could not decode validator can-send-cl-rewards response: %w", err)
	}
	if response.Error != "" {
		return api.CanSendClRewardsResponse{}, fmt.Errorf("could not get validator can-send-cl-rewards response: %s", response.Error)
	}

	return response, nil
}

func (c *Client) SendClRewards(validatorPubKey types.ValidatorPubkey) (api.SendClRewardsResponse, error) {
	responseBytes, err := c.callAPI(fmt.Sprintf("validator send-cl-rewards %s", validatorPubKey))
	if err != nil {
		return api.SendClRewardsResponse{}, fmt.Errorf("could not get validator send-cl-rewards response: %w", err)
	}
	var response api.SendClRewardsResponse
	if err := json.Unmarshal(responseBytes, &response); err != nil {
		return api.SendClRewardsResponse{}, fmt.Errorf("could not decode validator send-cl-rewards response: %w", err)
	}
	if response.Error != "" {
		return api.SendClRewardsResponse{}, fmt.Errorf("could not get validator send-cl-rewards response: %s", response.Error)
	}

	return response, nil
}

func (c *Client) CanWithdrawSd(amount *big.Int) (api.CanWithdrawSdResponse, error) {
	responseBytes, err := c.callAPI(fmt.Sprintf("node can-withdraw-sd %s", amount.String()))
	if err != nil {
		return api.CanWithdrawSdResponse{}, fmt.Errorf("could not get node can-withdraw-sd response: %w", err)
	}
	var response api.CanWithdrawSdResponse
	if err := json.Unmarshal(responseBytes, &response); err != nil {
		return api.CanWithdrawSdResponse{}, fmt.Errorf("could not decode node can-withdraw-sd response: %w", err)
	}
	if response.Error != "" {
		return api.CanWithdrawSdResponse{}, fmt.Errorf("could not get node can-withdraw-sd response: %s", response.Error)
	}

	return response, nil
}

func (c *Client) WithdrawSd(amount *big.Int) (api.WithdrawSdResponse, error) {
	responseBytes, err := c.callAPI(fmt.Sprintf("node withdraw-sd %s", amount.String()))
	if err != nil {
		return api.WithdrawSdResponse{}, fmt.Errorf("could not get node withdraw-sd response: %w", err)
	}
	var response api.WithdrawSdResponse
	if err := json.Unmarshal(responseBytes, &response); err != nil {
		return api.WithdrawSdResponse{}, fmt.Errorf("could not decode node withdraw-sd response: %w", err)
	}
	if response.Error != "" {
		return api.WithdrawSdResponse{}, fmt.Errorf("could not get node withdraw-sd response: %s", response.Error)
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

func (c *Client) EstimateClaimSpRewardsGas(cycles []*big.Int, depositSd bool) (api.EstimateClaimSpRewardsGasResponse, error) {
	stringifiedCycleList := string_utils.StringifyArray(cycles)
	responseBytes, err := c.callAPI(fmt.Sprintf("node estimate-claim-sp-rewards-gas %s %s", stringifiedCycleList, strconv.FormatBool(depositSd)))
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

func (c *Client) ClaimSpRewards(cycles []*big.Int, depositSd bool) (api.ClaimSpRewardsResponse, error) {
	stringifiedCycleList := string_utils.StringifyArray(cycles)
	responseBytes, err := c.callAPI(fmt.Sprintf("node claim-sp-rewards %s %s", stringifiedCycleList, strconv.FormatBool(depositSd)))
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

func (c *Client) CanUpdateOperatorName(operatorName string) (api.CanUpdateOperatorName, error) {
	responseBytes, err := c.callAPI("node can-update-operator-name", operatorName)
	if err != nil {
		return api.CanUpdateOperatorName{}, fmt.Errorf("could not get can-update-operator-name response: %w", err)
	}
	var response api.CanUpdateOperatorName
	if err := json.Unmarshal(responseBytes, &response); err != nil {
		return api.CanUpdateOperatorName{}, fmt.Errorf("could not decode can-update-operator-name response: %w", err)
	}
	if response.Error != "" {
		return api.CanUpdateOperatorName{}, fmt.Errorf("could not get can-update-operator-name response: %s", response.Error)
	}
	return response, nil
}

func (c *Client) UpdateOperatorName(operatorName string) (api.UpdateOperatorName, error) {
	responseBytes, err := c.callAPI("node update-operator-name", operatorName)
	if err != nil {
		return api.UpdateOperatorName{}, fmt.Errorf("could not get update-operator-name response: %w", err)
	}
	var response api.UpdateOperatorName
	if err := json.Unmarshal(responseBytes, &response); err != nil {
		return api.UpdateOperatorName{}, fmt.Errorf("could not decode update-operator-name response: %w", err)
	}
	if response.Error != "" {
		return api.UpdateOperatorName{}, fmt.Errorf("could not get update-operator-name response: %s", response.Error)
	}
	return response, nil
}

func (c *Client) CanUpdateOperatorRewardAddress(operatorRewardAddress common.Address) (api.CanUpdateOperatorRewardAddress, error) {
	responseBytes, err := c.callAPI("node can-update-operator-reward-address", operatorRewardAddress.Hex())
	if err != nil {
		return api.CanUpdateOperatorRewardAddress{}, fmt.Errorf("could not get can-update-operator-reward-address response: %w", err)
	}
	var response api.CanUpdateOperatorRewardAddress
	if err := json.Unmarshal(responseBytes, &response); err != nil {
		return api.CanUpdateOperatorRewardAddress{}, fmt.Errorf("could not decode can-update-operator-reward-address response: %w", err)
	}
	if response.Error != "" {
		return api.CanUpdateOperatorRewardAddress{}, fmt.Errorf("could not get can-update-operator-reward-address response: %s", response.Error)
	}
	return response, nil
}

func (c *Client) SetRewardAddress(operatorRewardAddress common.Address) (api.SetRewardAddress, error) {
	responseBytes, err := c.callAPI("node set-reward-address", operatorRewardAddress.Hex())
	if err != nil {
		return api.SetRewardAddress{}, fmt.Errorf("could not get set-reward-address response: %w", err)
	}

	var response api.SetRewardAddress
	if err := json.Unmarshal(responseBytes, &response); err != nil {
		return api.SetRewardAddress{}, fmt.Errorf("could not decode set-reward-address response: %w", err)
	}

	if response.Error != "" {
		return api.SetRewardAddress{}, fmt.Errorf("could not get set-reward-address response: %s", response.Error)
	}

	return response, nil
}

func (c *Client) RepaySd(amountWei *big.Int) (api.NodeRepaySDResponse, error) {
	responseBytes, err := c.callAPI(fmt.Sprintf("node repay-sd %s", amountWei.String()))
	if err != nil {
		return api.NodeRepaySDResponse{}, fmt.Errorf("could not repay SD: %w", err)
	}

	var response api.NodeRepaySDResponse
	if err := json.Unmarshal(responseBytes, &response); err != nil {
		return api.NodeRepaySDResponse{}, fmt.Errorf("could not decode repay node SD response: %w", err)
	}

	if response.Error != "" {
		return api.NodeRepaySDResponse{}, fmt.Errorf("could not repay SD: %s", response.Error)
	}

	return response, nil
}

func (c *Client) CanRepaySd(amountWei *big.Int) (api.CanRepaySDResponse, error) {
	responseBytes, err := c.callAPI(fmt.Sprintf("node can-repay-sd %s", amountWei.String()))
	if err != nil {
		return api.CanRepaySDResponse{}, fmt.Errorf("could not get CanNodeRepaySd SD: %w", err)
	}

	var response api.CanRepaySDResponse
	if err := json.Unmarshal(responseBytes, &response); err != nil {
		return api.CanRepaySDResponse{}, fmt.Errorf("could not decode  CanNodeRepaySd response: %w", err)
	}

	if response.Error != "" {
		return api.CanRepaySDResponse{}, fmt.Errorf("could not can-repay SD: %s", response.Error)
	}

	return response, nil
}

func (c *Client) NodeUtilizeSd(amountWei *big.Int) (api.NodeUtilitySDResponse, error) {
	responseBytes, err := c.callAPI(fmt.Sprintf("node utilize-sd %s", amountWei.String()))
	if err != nil {
		return api.NodeUtilitySDResponse{}, fmt.Errorf("could not utilize SD: %w", err)
	}

	var response api.NodeUtilitySDResponse
	if err := json.Unmarshal(responseBytes, &response); err != nil {
		return api.NodeUtilitySDResponse{}, fmt.Errorf("could not decode utilize response: %w", err)
	}

	if response.Error != "" {
		return api.NodeUtilitySDResponse{}, fmt.Errorf("could not utilize SD: %s", response.Error)
	}

	return response, nil
}

func (c *Client) CanNodeUtilizeSd(amountWei *big.Int) (api.CanUtilitySDResponse, error) {
	responseBytes, err := c.callAPI(fmt.Sprintf("node can-utilize-sd %s", amountWei.String()))
	if err != nil {
		return api.CanUtilitySDResponse{}, fmt.Errorf("could not utilize SD: %w", err)
	}

	var response api.CanUtilitySDResponse
	if err := json.Unmarshal(responseBytes, &response); err != nil {
		return api.CanUtilitySDResponse{}, fmt.Errorf("could not decode node SD utilize response: %w", err)
	}

	if response.Error != "" {
		return api.CanUtilitySDResponse{}, fmt.Errorf("could not utilize SD: %s", response.Error)
	}

	return response, nil
}

func (c *Client) GetSDStatus(numValidators *big.Int) (api.GetSdStatusResponse, error) {
	responseBytes, err := c.callAPI(fmt.Sprintf("node get-sd-status %s", numValidators))
	if err != nil {
		return api.GetSdStatusResponse{}, fmt.Errorf("could not get-sd-status: %w", err)
	}

	var response api.GetSdStatusResponse
	if err := json.Unmarshal(responseBytes, &response); err != nil {
		return api.GetSdStatusResponse{}, fmt.Errorf("could not decode node get SD status response: %w", err)
	}

	if response.Error != "" {
		return api.GetSdStatusResponse{}, fmt.Errorf("could not get SD status: %s", response.Error)
	}

	return response, nil
}
