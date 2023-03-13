package stader

import (
	"encoding/json"
	"fmt"
	"github.com/ethereum/go-ethereum/common"
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

// Check whether the node can stake RPL
func (c *Client) CanNodeDepositSd(amountWei *big.Int) (api.CanNodeStakeRplResponse, error) {
	responseBytes, err := c.callAPI(fmt.Sprintf("node can-node-deposit-sd %s", amountWei.String()))
	if err != nil {
		return api.CanNodeStakeRplResponse{}, fmt.Errorf("could not get can node stake RPL status: %w", err)
	}
	var response api.CanNodeStakeRplResponse
	if err := json.Unmarshal(responseBytes, &response); err != nil {
		return api.CanNodeStakeRplResponse{}, fmt.Errorf("could not decode can node stake RPL response: %w", err)
	}
	if response.Error != "" {
		return api.CanNodeStakeRplResponse{}, fmt.Errorf("could not get can node stake RPL status: %s", response.Error)
	}
	return response, nil
}

// Get the gas estimate for approving new RPL interaction
func (c *Client) NodeDepositSdApprovalGas(amountWei *big.Int) (api.NodeStakeRplApproveGasResponse, error) {
	responseBytes, err := c.callAPI(fmt.Sprintf("node get-deposit-sd-approval-gas %s", amountWei.String()))
	if err != nil {
		return api.NodeStakeRplApproveGasResponse{}, fmt.Errorf("could not get new RPL approval gas: %w", err)
	}
	var response api.NodeStakeRplApproveGasResponse
	if err := json.Unmarshal(responseBytes, &response); err != nil {
		return api.NodeStakeRplApproveGasResponse{}, fmt.Errorf("could not decode node stake RPL approve gas response: %w", err)
	}
	if response.Error != "" {
		return api.NodeStakeRplApproveGasResponse{}, fmt.Errorf("could not get new RPL approval gas: %s", response.Error)
	}
	return response, nil
}

// Approve RPL for staking against the node
func (c *Client) NodeDepositSdApprove(amountWei *big.Int) (api.NodeStakeRplApproveResponse, error) {
	responseBytes, err := c.callAPI(fmt.Sprintf("node deposit-sd-approve-sd %s", amountWei.String()))
	if err != nil {
		return api.NodeStakeRplApproveResponse{}, fmt.Errorf("could not approve RPL for staking: %w", err)
	}
	var response api.NodeStakeRplApproveResponse
	if err := json.Unmarshal(responseBytes, &response); err != nil {
		return api.NodeStakeRplApproveResponse{}, fmt.Errorf("could not decode stake node RPL approve response: %w", err)
	}
	if response.Error != "" {
		return api.NodeStakeRplApproveResponse{}, fmt.Errorf("could not approve RPL for staking: %s", response.Error)
	}
	return response, nil
}

// Stake RPL against the node waiting for approvalTxHash to be included in a block first
func (c *Client) NodeWaitAndStakeRpl(amountWei *big.Int, approvalTxHash common.Hash) (api.NodeStakeRplStakeResponse, error) {
	responseBytes, err := c.callAPI(fmt.Sprintf("node wait-and-stake-rpl %s %s", amountWei.String(), approvalTxHash.String()))
	if err != nil {
		return api.NodeStakeRplStakeResponse{}, fmt.Errorf("could not stake node RPL: %w", err)
	}
	var response api.NodeStakeRplStakeResponse
	if err := json.Unmarshal(responseBytes, &response); err != nil {
		return api.NodeStakeRplStakeResponse{}, fmt.Errorf("could not decode stake node RPL response: %w", err)
	}
	if response.Error != "" {
		return api.NodeStakeRplStakeResponse{}, fmt.Errorf("could not stake node RPL: %s", response.Error)
	}
	return response, nil
}

// Stake RPL against the node
func (c *Client) NodeDepositSd(amountWei *big.Int) (api.NodeStakeRplStakeResponse, error) {
	responseBytes, err := c.callAPI(fmt.Sprintf("node deposit-sd %s", amountWei.String()))
	if err != nil {
		return api.NodeStakeRplStakeResponse{}, fmt.Errorf("could not stake node RPL: %w", err)
	}
	var response api.NodeStakeRplStakeResponse
	if err := json.Unmarshal(responseBytes, &response); err != nil {
		return api.NodeStakeRplStakeResponse{}, fmt.Errorf("could not decode stake node RPL response: %w", err)
	}
	if response.Error != "" {
		return api.NodeStakeRplStakeResponse{}, fmt.Errorf("could not stake node RPL: %s", response.Error)
	}
	return response, nil
}

// Get a node's RPL allowance for the staking contract
func (c *Client) GetNodeDepositSdAllowance() (api.NodeStakeRplAllowanceResponse, error) {
	responseBytes, err := c.callAPI(fmt.Sprintf("node deposit-sd-allowance"))
	if err != nil {
		return api.NodeStakeRplAllowanceResponse{}, fmt.Errorf("could not get node stake RPL allowance: %w", err)
	}
	var response api.NodeStakeRplAllowanceResponse
	if err := json.Unmarshal(responseBytes, &response); err != nil {
		return api.NodeStakeRplAllowanceResponse{}, fmt.Errorf("could not decode node stake RPL allowance response: %w", err)
	}
	if response.Error != "" {
		return api.NodeStakeRplAllowanceResponse{}, fmt.Errorf("could not get node stake RPL allowance: %s", response.Error)
	}
	return response, nil
}

// Check whether the node can make a deposit
func (c *Client) CanNodeDeposit(amountWei *big.Int, salt *big.Int, numValidators *big.Int, submit bool) (api.CanNodeDepositResponse, error) {
	responseBytes, err := c.callAPI(fmt.Sprintf("node can-deposit %s %s %d %t", amountWei.String(), salt.String(), numValidators, submit))
	if err != nil {
		return api.CanNodeDepositResponse{}, fmt.Errorf("could not get can node deposit status: %w", err)
	}
	var response api.CanNodeDepositResponse
	if err := json.Unmarshal(responseBytes, &response); err != nil {
		return api.CanNodeDepositResponse{}, fmt.Errorf("could not decode can node deposit response: %w", err)
	}
	fmt.Printf("response in CanNodeDeposit is %v\n", response)
	if response.Error != "" {
		return api.CanNodeDepositResponse{}, fmt.Errorf("could not get can node deposit status: %s", response.Error)
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

func (c *Client) DebugExit(validatorIndex *big.Int) (api.DebugExitResponse, error) {
	fmt.Printf("cli: validatorIndex is %d\n", validatorIndex)
	responseBytes, err := c.callAPI(fmt.Sprintf("node debug-exit %d", validatorIndex))
	if err != nil {
		return api.DebugExitResponse{}, fmt.Errorf("could not get debug exit info: %w", err)
	}
	var response api.DebugExitResponse
	if err := json.Unmarshal(responseBytes, &response); err != nil {
		return api.DebugExitResponse{}, fmt.Errorf("could not decode debug exit info: %w", err)
	}
	if response.Error != "" {
		return api.DebugExitResponse{}, fmt.Errorf("could not get debug exit info: %s", response.Error)
	}
	return response, nil
}

// Make a node deposit
func (c *Client) NodeDeposit(amountWei *big.Int, salt *big.Int, numValidators *big.Int, submit bool) (api.NodeDepositResponse, error) {
	responseBytes, err := c.callAPI(fmt.Sprintf("node deposit %s %s %d %t", amountWei.String(), salt.String(), numValidators, submit))
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

// Get the deposit contract info for Stader and the Beacon Client
func (c *Client) DepositContractInfo() (api.ContractsInfoResponse, error) {
	responseBytes, err := c.callAPI("node deposit-contract-info")
	if err != nil {
		return api.ContractsInfoResponse{}, fmt.Errorf("could not get deposit contract info: %w", err)
	}
	var response api.ContractsInfoResponse
	if err := json.Unmarshal(responseBytes, &response); err != nil {
		return api.ContractsInfoResponse{}, fmt.Errorf("could not decode deposit contract info response: %w", err)
	}
	if response.Error != "" {
		return api.ContractsInfoResponse{}, fmt.Errorf("could not get deposit contract info: %s", response.Error)
	}
	return response, nil
}

func (c *Client) ResolveEnsName(name string) (api.ResolveEnsNameResponse, error) {
	responseBytes, err := c.callAPI(fmt.Sprintf("node resolve-ens-name %s", name))
	if err != nil {
		return api.ResolveEnsNameResponse{}, fmt.Errorf("could not resolve ENS name: %w", err)
	}
	var response api.ResolveEnsNameResponse
	if err := json.Unmarshal(responseBytes, &response); err != nil {
		return api.ResolveEnsNameResponse{}, fmt.Errorf("could not decode resolve-ens-name: %w", err)
	}
	if response.Error != "" {
		return api.ResolveEnsNameResponse{}, fmt.Errorf("could not resolve ENS name: %s", response.Error)
	}
	return response, nil
}
func (c *Client) ReverseResolveEnsName(name string) (api.ResolveEnsNameResponse, error) {
	responseBytes, err := c.callAPI(fmt.Sprintf("node reverse-resolve-ens-name %s", name))
	if err != nil {
		return api.ResolveEnsNameResponse{}, fmt.Errorf("could not reverse resolve ENS name: %w", err)
	}
	var response api.ResolveEnsNameResponse
	if err := json.Unmarshal(responseBytes, &response); err != nil {
		return api.ResolveEnsNameResponse{}, fmt.Errorf("could not decode reverse-resolve-ens-name: %w", err)
	}
	if response.Error != "" {
		return api.ResolveEnsNameResponse{}, fmt.Errorf("could not reverse resolve ENS name: %s", response.Error)
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
