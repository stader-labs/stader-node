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
package node

import (
	"github.com/urfave/cli"

	"github.com/stader-labs/stader-node/shared/utils/api"
	cliutils "github.com/stader-labs/stader-node/shared/utils/cli"
)

// Register subcommands
func RegisterSubcommands(command *cli.Command, name string, aliases []string) {
	command.Subcommands = append(command.Subcommands, cli.Command{
		Name:    name,
		Aliases: aliases,
		Usage:   "Manage the node",
		Subcommands: []cli.Command{

			{
				Name:      "status",
				Aliases:   []string{"s"},
				Usage:     "Get the node's status",
				UsageText: "stader-cli api node status",
				Action: func(c *cli.Context) error {

					// Validate args
					if err := cliutils.ValidateArgCount(c, 0); err != nil {
						return err
					}

					// Run
					api.PrintResponse(getStatus(c))
					return nil

				},
			},

			{
				Name:      "sync",
				Aliases:   []string{"y"},
				Usage:     "Get the sync progress of the eth1 and eth2 clients",
				UsageText: "stader-cli api node sync",
				Action: func(c *cli.Context) error {

					// Validate args
					if err := cliutils.ValidateArgCount(c, 0); err != nil {
						return err
					}

					// Run
					api.PrintResponse(getSyncProgress(c))
					return nil

				},
			},

			{
				Name:      "can-register",
				Usage:     "Check whether the node can be registered with Stader",
				UsageText: "stader-cli api node can-register timezone-location",
				Action: func(c *cli.Context) error {

					// Validate args
					if err := cliutils.ValidateArgCount(c, 3); err != nil {
						return err
					}

					if err := cliutils.ValidateArgCount(c, 3); err != nil {
						return err
					}

					operatorName := c.Args().Get(0)

					operatorRewardAddress, err := cliutils.ValidateAddress("operator-reward-address", c.Args().Get(1))
					if err != nil {
						return err
					}

					socializeMev, err := cliutils.ValidateBool("socialize-mev", c.Args().Get(2))

					// Run
					api.PrintResponse(canRegisterNode(c, operatorName, operatorRewardAddress, socializeMev))
					return nil

				},
			},
			{
				Name:      "register",
				Aliases:   []string{"r"},
				Usage:     "Register the node with Stader",
				UsageText: "stader-cli api node register operator-name operator-reward-address socialize-mev",
				Action: func(c *cli.Context) error {

					// Validate args
					if err := cliutils.ValidateArgCount(c, 3); err != nil {
						return err
					}

					operatorName := c.Args().Get(0)

					operatorRewardAddress, err := cliutils.ValidateAddress("operator-reward-address", c.Args().Get(1))
					if err != nil {
						return err
					}

					socializeMev, err := cliutils.ValidateBool("socialize-mev", c.Args().Get(2))

					// Run
					api.PrintResponse(registerNode(c, operatorName, operatorRewardAddress, socializeMev))
					return nil

				},
			},

			{
				Name:      "can-node-deposit-sd",
				Usage:     "Check whether the node can stake SD",
				UsageText: "stader-cli api node can-node-deposit-sd amount",
				Action: func(c *cli.Context) error {

					// Validate args
					if err := cliutils.ValidateArgCount(c, 1); err != nil {
						return err
					}
					amountWei, err := cliutils.ValidatePositiveWeiAmount("stake amount", c.Args().Get(0))
					if err != nil {
						return err
					}

					// Run
					api.PrintResponse(canNodeDepositSd(c, amountWei))
					return nil

				},
			},
			{
				Name:      "deposit-sd-approve-sd",
				Aliases:   []string{"k1"},
				Usage:     "Approve SD for staking against the node",
				UsageText: "stader-cli api node deposit-sd-approve-sd amount",
				Action: func(c *cli.Context) error {

					// Validate args
					if err := cliutils.ValidateArgCount(c, 1); err != nil {
						return err
					}
					amountWei, err := cliutils.ValidatePositiveWeiAmount("stake amount", c.Args().Get(0))
					if err != nil {
						return err
					}

					// Run
					api.PrintResponse(approveSd(c, amountWei))
					return nil

				},
			},
			{
				Name:      "wait-and-deposit-sd",
				Aliases:   []string{"k2"},
				Usage:     "Deposit SD against the node, waiting for approval tx-hash to be included in a block first",
				UsageText: "stader-cli api node wait-and-deposit-sd amount tx-hash",
				Action: func(c *cli.Context) error {

					// Validate args
					if err := cliutils.ValidateArgCount(c, 2); err != nil {
						return err
					}
					amountWei, err := cliutils.ValidatePositiveWeiAmount("stake amount", c.Args().Get(0))
					if err != nil {
						return err
					}
					hash, err := cliutils.ValidateTxHash("tx-hash", c.Args().Get(1))
					if err != nil {
						return err
					}

					// Run
					api.PrintResponse(waitForApprovalAndDepositSd(c, amountWei, hash))
					return nil

				},
			},
			{
				Name:      "get-deposit-sd-approval-gas",
				Usage:     "Estimate the gas cost of new SD interaction approval",
				UsageText: "stader-cli api node get-deposit-sd-approval-gas",
				Action: func(c *cli.Context) error {

					// Validate args
					if err := cliutils.ValidateArgCount(c, 1); err != nil {
						return err
					}
					amountWei, err := cliutils.ValidatePositiveWeiAmount("approve amount", c.Args().Get(0))
					if err != nil {
						return err
					}

					// Run
					api.PrintResponse(getDepositSdApprovalGas(c, amountWei))
					return nil

				},
			},
			{
				Name:      "deposit-sd-allowance",
				Usage:     "Get the node's SD allowance for the collateral contract",
				UsageText: "stader-cli api node deposit-sd-allowance",
				Action: func(c *cli.Context) error {

					// Validate args
					if err := cliutils.ValidateArgCount(c, 0); err != nil {
						return err
					}

					// Run
					api.PrintResponse(allowanceSd(c))
					return nil

				},
			},
			{
				Name:      "deposit-sd",
				Aliases:   []string{"k3"},
				Usage:     "Deposit SD against the node",
				UsageText: "stader-cli api node deposit-sd amount",
				Action: func(c *cli.Context) error {

					// Validate args
					if err := cliutils.ValidateArgCount(c, 1); err != nil {
						return err
					}
					amountWei, err := cliutils.ValidatePositiveWeiAmount("stake amount", c.Args().Get(0))
					if err != nil {
						return err
					}

					// Run
					api.PrintResponse(depositSdAsCollateral(c, amountWei))
					return nil

				},
			},
			{
				Name:      "can-deposit",
				Usage:     "Check whether the node can make a deposit",
				UsageText: "stader-cli api node can-deposit amount min-fee salt",
				Action: func(c *cli.Context) error {

					//// Validate args
					// Validate args
					if err := cliutils.ValidateArgCount(c, 4); err != nil {
						return err
					}
					amountWei, err := cliutils.ValidateWeiAmount("deposit amount", c.Args().Get(0))
					if err != nil {
						return err
					}

					salt, err := cliutils.ValidateBigInt("salt", c.Args().Get(1))
					if err != nil {
						return err
					}

					numValidators, err := cliutils.ValidateBigInt("num-validators", c.Args().Get(2))
					if err != nil {
						return err
					}

					submit, err := cliutils.ValidateBool("submit", c.Args().Get(3))
					if err != nil {
						return err
					}

					api.PrintResponse(canNodeDeposit(c, amountWei, salt, numValidators, submit))

					return nil

				},
			},
			{
				Name:      "deposit",
				Aliases:   []string{"d"},
				Usage:     "Make a deposit and create a validator, or just make and sign the transaction (when submit = false)",
				UsageText: "stader api node deposit amount salt submit",
				Action: func(c *cli.Context) error {

					// Validate args
					if err := cliutils.ValidateArgCount(c, 4); err != nil {
						return err
					}
					amountWei, err := cliutils.ValidateWeiAmount("deposit amount", c.Args().Get(0))
					if err != nil {
						return err
					}

					salt, err := cliutils.ValidateBigInt("salt", c.Args().Get(1))
					if err != nil {
						return err
					}

					numValidators, err := cliutils.ValidateBigInt("num-validators", c.Args().Get(2))
					if err != nil {
						return err
					}

					submit, err := cliutils.ValidateBool("submit", c.Args().Get(3))
					if err != nil {
						return err
					}

					// Run
					response, err := nodeDeposit(c, amountWei, salt, numValidators, submit)
					if submit {
						api.PrintResponse(response, err)
					}
					return nil

				},
			},

			{
				Name:      "can-send",
				Usage:     "Check whether the node can send ETH or tokens to an address",
				UsageText: "stader-cli api node can-send amount token",
				Action: func(c *cli.Context) error {

					// Validate args
					if err := cliutils.ValidateArgCount(c, 2); err != nil {
						return err
					}
					amountWei, err := cliutils.ValidatePositiveWeiAmount("send amount", c.Args().Get(0))
					if err != nil {
						return err
					}
					token, err := cliutils.ValidateTokenType("token type", c.Args().Get(1))
					if err != nil {
						return err
					}

					// Run
					api.PrintResponse(canNodeSend(c, amountWei, token))
					return nil

				},
			},
			{
				Name:      "send",
				Aliases:   []string{"n"},
				Usage:     "Send ETH or tokens from the node account to an address",
				UsageText: "stader-cli api node send amount token to",
				Action: func(c *cli.Context) error {

					// Validate args
					if err := cliutils.ValidateArgCount(c, 3); err != nil {
						return err
					}
					amountWei, err := cliutils.ValidatePositiveWeiAmount("send amount", c.Args().Get(0))
					if err != nil {
						return err
					}
					token, err := cliutils.ValidateTokenType("token type", c.Args().Get(1))
					if err != nil {
						return err
					}
					toAddress, err := cliutils.ValidateAddress("to address", c.Args().Get(2))
					if err != nil {
						return err
					}

					// Run
					api.PrintResponse(nodeSend(c, amountWei, token, toAddress))
					return nil

				},
			},

			{
				Name:      "get-contracts-info",
				Usage:     "Get information about the deposit contract and stader contract on the current network",
				UsageText: "stader-cli api node get-contracts-info",
				Action: func(c *cli.Context) error {

					// Validate args
					if err := cliutils.ValidateArgCount(c, 0); err != nil {
						return err
					}

					// Run
					api.PrintResponse(getContractsInfo(c))
					return nil

				},
			},

			{
				Name:      "sign",
				Usage:     "Signs a transaction with the node's private key. The TX must be serialized as a hex string.",
				UsageText: "stader-cli api node sign tx",
				Action: func(c *cli.Context) error {

					// Validate args
					if err := cliutils.ValidateArgCount(c, 1); err != nil {
						return err
					}

					data := c.Args().Get(0)

					// Run
					api.PrintResponse(sign(c, data))
					return nil

				},
			},

			{
				Name:      "sign-message",
				Usage:     "Signs an arbitrary message with the node's private key.",
				UsageText: "stader-cli api node sign-message 'message'",
				Action: func(c *cli.Context) error {

					// Validate args
					if err := cliutils.ValidateArgCount(c, 1); err != nil {
						return err
					}

					message := c.Args().Get(0)

					// Run
					api.PrintResponse(signMessage(c, message))
					return nil

				},
			},
			{
				Name:      "can-exit-validator",
				Usage:     "Can validator exit",
				UsageText: "stader-cli api node can-exit-validator validator-pub-key",
				Action: func(c *cli.Context) error {

					// Validate args
					if err := cliutils.ValidateArgCount(c, 1); err != nil {
						return err
					}

					validatorPubKey, err := cliutils.ValidatePubkey("validator-pub-key", c.Args().Get(0))
					if err != nil {
						return err
					}

					api.PrintResponse(canExitValidator(c, validatorPubKey))
					return nil

				},
			},
			{
				Name:      "exit-validator",
				Usage:     "Exit validator",
				UsageText: "stader-cli api node exit-validator validator-pub-key",
				Action: func(c *cli.Context) error {

					// Validate args
					if err := cliutils.ValidateArgCount(c, 1); err != nil {
						return err
					}

					validatorPubKey, err := cliutils.ValidatePubkey("validator-pub-key", c.Args().Get(0))
					if err != nil {
						return err
					}

					api.PrintResponse(exitValidator(c, validatorPubKey))
					return nil

				},
			},
			{
				Name:      "can-update-socialize-el",
				Usage:     "Can opt in or opt out of socializing pool",
				UsageText: "stader-cli api node can-update-socialize-el --socialize-el",
				Action: func(c *cli.Context) error {

					// Validate args
					//if err := cliutils.ValidateArgCount(c, 1); err != nil {
					//	return err
					//}

					socializeEl, err := cliutils.ValidateBool("socialize-el", c.Args().Get(0))
					if err != nil {
						return err
					}

					api.PrintResponse(canUpdateSocializeEl(c, socializeEl))
					return nil

				},
			},
			{
				Name:      "update-socialize-el",
				Usage:     "Opt in or opt out of socializing pool",
				UsageText: "stader-cli api node update-socialize-el --socialize-el",
				Action: func(c *cli.Context) error {

					// Validate args
					if err := cliutils.ValidateArgCount(c, 1); err != nil {
						return err
					}

					socializeEl, err := cliutils.ValidateBool("socialize-el", c.Args().Get(0))
					if err != nil {
						return err
					}

					api.PrintResponse(updateSocializeEl(c, socializeEl))
					return nil

				},
			},
			{
				Name:      "can-claim-el-rewards",
				Usage:     "Can claim el rewards",
				UsageText: "stader-cli api node can-claim-el-rewards",
				Action: func(c *cli.Context) error {

					// Validate args
					if err := cliutils.ValidateArgCount(c, 0); err != nil {
						return err
					}

					api.PrintResponse(CanClaimElRewards(c))
					return nil

				},
			},
			{
				Name:      "claim-el-rewards",
				Usage:     "Claim el rewards",
				UsageText: "stader-cli api node claim-el-rewards",
				Action: func(c *cli.Context) error {

					// Validate args
					if err := cliutils.ValidateArgCount(c, 0); err != nil {
						return err
					}

					api.PrintResponse(ClaimElRewards(c))
					return nil

				},
			},
			{
				Name:      "can-claim-cl-rewards",
				Usage:     "Can claim cl rewards of a validator",
				UsageText: "stader-cli api node can-claim-cl-rewards --validator-pub-key",
				Action: func(c *cli.Context) error {

					validatorPubKey, err := cliutils.ValidatePubkey("validator-pub-key", c.Args().Get(0))
					if err != nil {
						return err
					}

					api.PrintResponse(CanClaimClRewards(c, validatorPubKey))
					return nil

				},
			},
			{
				Name:      "claim-cl-rewards",
				Usage:     "Claim cl rewards of a validator",
				UsageText: "stader-cli api node claim-cl-rewards --validator-pub-key",
				Action: func(c *cli.Context) error {

					validatorPubKey, err := cliutils.ValidatePubkey("validator-pub-key", c.Args().Get(0))
					if err != nil {
						return err
					}

					api.PrintResponse(ClaimClRewards(c, validatorPubKey))
					return nil

				},
			},
			{
				Name:      "can-node-request-sd-withdraw",
				Usage:     "Check whether the node can request to withdraw SD",
				UsageText: "stader-cli api node can-node-request-sd-withdraw amount",
				Action: func(c *cli.Context) error {

					// Validate args
					if err := cliutils.ValidateArgCount(c, 1); err != nil {
						return err
					}
					amountWei, err := cliutils.ValidatePositiveWeiAmount("stake amount", c.Args().Get(0))
					if err != nil {
						return err
					}

					// Run
					api.PrintResponse(canRequestSdWithdraw(c, amountWei))
					return nil

				},
			},
			{
				Name:      "node-request-sd-withdraw",
				Usage:     "Request SD withdraw",
				UsageText: "stader-cli api node node-request-sd-withdraw amount",
				Action: func(c *cli.Context) error {

					// Validate args
					if err := cliutils.ValidateArgCount(c, 1); err != nil {
						return err
					}
					amountWei, err := cliutils.ValidatePositiveWeiAmount("stake amount", c.Args().Get(0))
					if err != nil {
						return err
					}

					// Run
					api.PrintResponse(requestSdWithdraw(c, amountWei))
					return nil

				},
			},
			{
				Name:      "can-node-claim-sd",
				Usage:     "Check whether the node can claim the SD requested to withdraw",
				UsageText: "stader-cli api node can-node-claim-sd",
				Action: func(c *cli.Context) error {

					// Validate args
					if err := cliutils.ValidateArgCount(c, 0); err != nil {
						return err
					}

					// Run
					api.PrintResponse(canClaimSd(c))
					return nil

				},
			},
			{
				Name:      "node-claim-sd",
				Usage:     "Claim the SD requested to withdraw",
				UsageText: "stader-cli api node node-claim-sd",
				Action: func(c *cli.Context) error {

					// Validate args
					if err := cliutils.ValidateArgCount(c, 0); err != nil {
						return err
					}

					// Run
					api.PrintResponse(claimSd(c))
					return nil

				},
			},
			{
				Name:      "can-download-sp-merkle-proofs",
				Usage:     "Can we Download missing socializing merkle proofs",
				UsageText: "stader-cli api node can-download-sp-merkle-proofs",
				Action: func(c *cli.Context) error {

					// Validate args
					if err := cliutils.ValidateArgCount(c, 0); err != nil {
						return err
					}

					// Run
					api.PrintResponse(canDownloadSpMerkleProofs(c))
					return nil

				},
			},
			{
				Name:      "download-sp-merkle-proofs",
				Usage:     "Download missing socializing merkle proofs",
				UsageText: "stader-cli api node download-sp-merkle-proofs",
				Action: func(c *cli.Context) error {

					// Validate args
					if err := cliutils.ValidateArgCount(c, 0); err != nil {
						return err
					}

					// Run
					api.PrintResponse(downloadSpMerkleProofs(c))
					return nil

				},
			},
			{
				Name:      "detailed-cycles-info",
				Usage:     "Get detailed reward cycles info",
				UsageText: "stader-cli api node detailed-cycles-info cycles",
				Action: func(c *cli.Context) error {

					// Validate args
					if err := cliutils.ValidateArgCount(c, 1); err != nil {
						return err
					}

					cycles := c.Args().Get(0)

					// Run
					api.PrintResponse(GetCyclesDetailedInfo(c, cycles))
					return nil

				},
			},
			{
				Name:      "can-claim-sp-rewards",
				Usage:     "Can we claim the SP rewards",
				UsageText: "stader-cli api node can-claim-sp-rewards",
				Action: func(c *cli.Context) error {

					// Validate args
					if err := cliutils.ValidateArgCount(c, 0); err != nil {
						return err
					}

					// Run
					api.PrintResponse(canClaimSpRewards(c))
					return nil

				},
			},
			{
				Name:      "claim-sp-rewards",
				Usage:     "Claim the SP rewards",
				UsageText: "stader-cli api node claim-sp-rewards cycles",
				Action: func(c *cli.Context) error {

					// Validate args
					if err := cliutils.ValidateArgCount(c, 1); err != nil {
						return err
					}

					cycles := c.Args().Get(0)
					//fmt.Printf("cycles is %s\n", cycles)
					// Run
					api.PrintResponse(claimSpRewards(c, cycles))
					return nil

				},
			},
			{
				Name:      "estimate-claim-sp-rewards-gas",
				Usage:     "Estimate the gas required to claim the SP rewards",
				UsageText: "stader-cli api node estimate-claim-sp-rewards-gas cycles",
				Action: func(c *cli.Context) error {

					// Validate args
					if err := cliutils.ValidateArgCount(c, 1); err != nil {
						return err
					}

					cycles := c.Args().Get(0)
					// Run
					api.PrintResponse(estimateSpRewardsGas(c, cycles))
					return nil

				},
			},
			{
				Name:      "can-update-operator-details",
				Usage:     "Can we update the operator details",
				UsageText: "stader-cli api node can-update-operator-details operator-name operator-reward-address",
				Action: func(c *cli.Context) error {

					// Validate args
					if err := cliutils.ValidateArgCount(c, 2); err != nil {
						return err
					}

					operatorName := c.Args().Get(0)

					operatorRewardAddress, err := cliutils.ValidateAddress("operator-reward-address", c.Args().Get(1))
					if err != nil {
						return err
					}

					// Run
					api.PrintResponse(CanUpdateOperatorDetails(c, operatorName, operatorRewardAddress))
					return nil

				},
			},
			{
				Name:      "update-operator-details",
				Usage:     "Update the operator details",
				UsageText: "stader-cli api node update-operator-details operator-name operator-reward-address",
				Action: func(c *cli.Context) error {

					// Validate args
					if err := cliutils.ValidateArgCount(c, 2); err != nil {
						return err
					}

					operatorName := c.Args().Get(0)

					operatorRewardAddress, err := cliutils.ValidateAddress("operator-reward-address", c.Args().Get(1))
					if err != nil {
						return err
					}
					// Run
					api.PrintResponse(UpdateOperatorDetails(c, operatorName, operatorRewardAddress))
					return nil

				},
			},
		},
	})
}
