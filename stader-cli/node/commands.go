/*
This work is licensed and released under GNU GPL v3 or any other later versions.
The full text of the license is below/ found at <http://www.gnu.org/licenses/>

(c) 2023 Rocket Pool Pty Ltd. Modified under GNU GPL v3. [0.3.0-beta]

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
	"fmt"

	"github.com/urfave/cli"

	cliutils "github.com/stader-labs/stader-node/shared/utils/cli"
)

// Register commands
func RegisterCommands(app *cli.App, name string, aliases []string) {
	app.Commands = append(app.Commands, cli.Command{
		Name:    name,
		Aliases: aliases,
		Usage:   "Manage the node",
		Subcommands: []cli.Command{
			{
				Name:      "status",
				Aliases:   []string{"s"},
				Usage:     "Get the node's status",
				UsageText: "stader-cli node status",
				Action: func(c *cli.Context) error {

					// Validate args
					if err := cliutils.ValidateArgCount(c, 0); err != nil {
						return err
					}

					// Run
					return getStatus(c)

				},
			},
			{
				Name:      "sync",
				Aliases:   []string{"y"},
				Usage:     "Get the sync progress of the eth1 and eth2 clients",
				UsageText: "stader-cli node sync",
				Action: func(c *cli.Context) error {

					// Validate args
					if err := cliutils.ValidateArgCount(c, 0); err != nil {
						return err
					}

					// Run
					return getSyncProgress(c)

				},
			},
			{
				Name:      "update-socialize-el",
				Aliases:   []string{"y"},
				Usage:     "Opt in or Opt out of socializing pool",
				UsageText: "stader-cli node socialize-el [options]",
				Flags: []cli.Flag{
					cli.StringFlag{
						Name:  "socialize-el, sel",
						Usage: "Should EL rewards be socialized (will default to true)",
					},
					cli.BoolFlag{
						Name:  "yes, y",
						Usage: "Automatically confirm socialize-el update",
					},
				},
				Action: func(c *cli.Context) error {

					// Validate args
					socializeEl, err := cliutils.ValidateBool("socialize-el", c.String("socialize-el"))
					if err != nil {
						return err
					}

					// Run
					return UpdateSocializeEl(c, socializeEl)

				},
			},
			{
				Name:      "register",
				Aliases:   []string{"r"},
				Usage:     "Register the node with Stader",
				UsageText: "stader-cli node register [options]",
				Flags: []cli.Flag{
					cli.StringFlag{
						Name:  "operator-name, on",
						Usage: "The name of the operator",
					},
					cli.StringFlag{
						Name:  "operator-reward-address, ora",
						Usage: "The address at which operator will get rewards (will default to the current node address)",
					},
					cli.BoolTFlag{
						Name:  "socialize-el, sel",
						Usage: "Should EL rewards be socialized (will default to true)",
					},
					cli.BoolFlag{
						Name:  "yes, y",
						Usage: "Automatically confirm node registration",
					},
				},
				Action: func(c *cli.Context) error {
					// Validate flags
					if c.String("operator-name") == "" {
						return fmt.Errorf("operator-name is required")
					}

					if c.String("operator-reward-address") != "" {
						if _, err := cliutils.ValidateAddress("operator-reward-address", c.String("operator-reward-address")); err != nil {
							return err
						}
					}

					// Run
					return registerNode(c)

				},
			},
			{
				Name:      "deposit-sd",
				Aliases:   []string{"k"},
				Usage:     "Deposit SD against the node",
				UsageText: "stader-cli node deposit-sd [options]",
				Flags: []cli.Flag{
					cli.StringFlag{
						Name:  "amount, a",
						Usage: "The amount of SD to deposit",
					},
					cli.BoolFlag{
						Name:  "yes, y",
						Usage: "Automatically confirm SD deposit",
					},
				},
				Action: func(c *cli.Context) error {

					if _, err := cliutils.ValidatePositiveEthAmount("sd deposit amount", c.String("amount")); err != nil {
						return err
					}

					// Run
					return nodeDepositSd(c)

				},
			},
			{
				Name:      "deposit",
				Aliases:   []string{"d"},
				Usage:     "Make a deposit and create a validator",
				UsageText: "stader-cli node deposit [options]",
				Flags: []cli.Flag{
					cli.BoolFlag{
						Name:  "yes, y",
						Usage: "Automatically confirm deposit",
					},
					cli.StringFlag{
						Name:  "salt, l",
						Usage: "An optional seed to use when generating the new validator address.",
					},
					cli.Uint64Flag{
						Name:  "num-validators, nv",
						Usage: "Number of validators you want to create (Required)",
					},
				},
				Action: func(c *cli.Context) error {

					// Validate flags
					if c.String("amount") != "" {
						if _, err := cliutils.ValidateDepositEthAmount("deposit amount", c.String("amount")); err != nil {
							return err
						}
					}
					if c.String("salt") != "" {
						if _, err := cliutils.ValidateBigInt("salt", c.String("salt")); err != nil {
							return err
						}
					}
					if c.Uint64("num-validators") == 0 {
						return fmt.Errorf("num-validator needs to be > 0")
					}

					// Run
					return nodeDeposit(c)

				},
			},
			{
				Name:      "send",
				Aliases:   []string{"n"},
				Usage:     "Send ETH or SD, EthX tokens from the node account to an address.",
				UsageText: "stader-cli node send [options] amount token to",
				Flags: []cli.Flag{
					cli.BoolFlag{
						Name:  "yes, y",
						Usage: "Automatically confirm token send",
					},
				},
				Action: func(c *cli.Context) error {

					// Validate args
					if err := cliutils.ValidateArgCount(c, 3); err != nil {
						return err
					}
					amount, err := cliutils.ValidatePositiveEthAmount("send amount", c.Args().Get(0))
					if err != nil {
						return err
					}
					token, err := cliutils.ValidateTokenType("token type", c.Args().Get(1))
					if err != nil {
						return err
					}

					// Run
					return nodeSend(c, amount, token, c.Args().Get(2))

				},
			},
			{
				Name:      "get-contracts-info",
				Aliases:   []string{"c"},
				Usage:     "Get the current network contracts info",
				UsageText: "stader-cli node get-contracts-info",
				Action: func(c *cli.Context) error {

					// Validate args
					if err := cliutils.ValidateArgCount(c, 0); err != nil {
						return err
					}

					// Run
					return getContractsInfo(c)
				},
			},
			{
				Name:      "exit",
				Aliases:   []string{"e"},
				Usage:     "Exit validator",
				UsageText: "stader-cli node exit --validator-pub-key",
				Flags: []cli.Flag{
					cli.StringFlag{
						Name:  "validator-pub-key, vpk",
						Usage: "Public key of validator we want to exit",
					},
					cli.BoolFlag{
						Name:  "yes, y",
						Usage: "Automatically confirm validator exit",
					},
				},
				Action: func(c *cli.Context) error {

					//// Validate args
					validatorPubKey, err := cliutils.ValidatePubkey("validator-pub-key", c.String("validator-pub-key"))
					if err != nil {
						return err
					}

					// Run
					return ExitValidator(c, validatorPubKey)
				},
			},
			{
				Name:      "send-presigned-exit-msg",
				Aliases:   []string{"spem"},
				Usage:     "Send the presigned exit msg to stader",
				UsageText: "stader-cli node send-presigned-exit-msg --validator-pub-key",
				Flags: []cli.Flag{
					cli.StringFlag{
						Name:  "validator-pub-key, vpk",
						Usage: "Validator index for whom we want to generate the debug exit",
					},
					cli.BoolFlag{
						Name:  "yes, y",
						Usage: "Automatically confirm pre-signed message sending",
					},
				},
				Action: func(c *cli.Context) error {
					//// Validate args
					validatorPubKey, err := cliutils.ValidatePubkey("validator-pub-key", c.String("validator-pub-key"))
					if err != nil {
						return err
					}

					// Run
					return SendSignedPresignedMessage(c, validatorPubKey)
				},
			},
			{
				Name:      "withdraw-el-rewards",
				Aliases:   []string{"wer"},
				Usage:     "Withdraw all Execution Layer rewards to the node reward address. This only includes non-socializing pool rewards",
				UsageText: "stader-cli node withdraw-el-rewards",
				Flags: []cli.Flag{cli.BoolFlag{
					Name:  "yes, y",
					Usage: "Automatically confirm EL rewards withdrawal",
				}},
				Action: func(c *cli.Context) error {
					// Run
					return WithdrawElRewards(c)
				},
			},
			{
				Name:      "withdraw-cl-rewards",
				Aliases:   []string{"wcr"},
				Usage:     "Withdraw all Consensus Layer rewards to the node reward address.",
				UsageText: "stader-cli node withdraw-cl-rewards --validator-pub-key",
				Flags: []cli.Flag{
					cli.StringFlag{
						Name:  "validator-pub-key, vpk",
						Usage: "Public key of validator we want to exit",
					},
					cli.BoolFlag{
						Name:  "yes, y",
						Usage: "Automatically confirm CL rewards withdrawal",
					},
				},
				Action: func(c *cli.Context) error {

					validatorPubKey, err := cliutils.ValidatePubkey("validator-pub-key", c.String("validator-pub-key"))
					if err != nil {
						return err
					}
					// Run
					return WithdrawClRewards(c, validatorPubKey)
				},
			},
			{
				Name:      "settle-exit-funds",
				Aliases:   []string{"sef"},
				Usage:     "Settle all funds validator should receive post exit",
				UsageText: "stader-cli node settle-exit-funds --validator-pub-key",
				Flags: []cli.Flag{
					cli.StringFlag{
						Name:  "validator-pub-key, vpk",
						Usage: "Public key of validator",
					},
					cli.BoolFlag{
						Name:  "yes, y",
						Usage: "Automatically confirm exit funds settling",
					},
				},
				Action: func(c *cli.Context) error {

					validatorPubKey, err := cliutils.ValidatePubkey("validator-pub-key", c.String("validator-pub-key"))
					if err != nil {
						return err
					}
					// Run
					return SettleExitFunds(c, validatorPubKey)
				},
			},
			{
				Name:      "request-withdraw-sd-collateral",
				Aliases:   []string{"sef"},
				Usage:     "Request to withdraw SD collateral",
				UsageText: "stader-cli node withdraw-sd-collateral --amount",
				Flags: []cli.Flag{
					cli.StringFlag{
						Name:  "amount, a",
						Usage: "The amount of SD to withdraw",
					},
					cli.BoolFlag{
						Name:  "yes, y",
						Usage: "Automatically confirm withdraw sd collateral",
					},
				},
				Action: func(c *cli.Context) error {

					if _, err := cliutils.ValidatePositiveEthAmount("sd withdraw amount", c.String("amount")); err != nil {
						return err
					}

					// Run
					return WithdrawSd(c)
				},
			},
			{
				Name:      "claim-sd",
				Aliases:   []string{"cs"},
				Usage:     "Claim SD from the stader contract",
				UsageText: "stader-cli node claim-sd",
				Flags: []cli.Flag{
					cli.BoolFlag{
						Name:  "yes, y",
						Usage: "Automatically confirm withdraw sd collateral",
					},
				},
				Action: func(c *cli.Context) error {

					if err := cliutils.ValidateArgCount(c, 0); err != nil {
						return err
					}

					// Run
					return claimSd(c)
				},
			},
			{
				Name:      "download-sp-merkle-proofs",
				Aliases:   []string{"dspmp"},
				Usage:     "Download all the missing Socializing Pool merkle proofs for the operator",
				UsageText: "stader-cli node download-sp-merkle-proofs",
				Flags: []cli.Flag{
					cli.BoolFlag{
						Name:  "yes, y",
						Usage: "Automatically confirm merkle proofs download",
					},
				},
				Action: func(c *cli.Context) error {

					if err := cliutils.ValidateArgCount(c, 0); err != nil {
						return err
					}

					// Run
					return downloadSPMerkleProofs(c)
				},
			},
			{
				Name:      "claim-sp-rewards",
				Aliases:   []string{"cspr"},
				Usage:     "Claim Socializing Pool Rewards for given cycles",
				UsageText: "stader-cli node claim-sp-rewards --download-merkles-proofs --yes",
				Flags: []cli.Flag{
					cli.BoolFlag{
						Name:  "download-merkles-proofs, dmp",
						Usage: "Download merkle proofs for the missing cycles to claim if not present",
					},
					cli.BoolFlag{
						Name:  "yes, y",
						Usage: "Automatically confirm claim of rewards",
					},
				},
				Action: func(c *cli.Context) error {

					if err := cliutils.ValidateArgCount(c, 0); err != nil {
						return err
					}

					downloadMerkleProofs, err := cliutils.ValidateBool("download-merkles-proofs", c.String("download-merkles-proofs"))
					if err != nil {
						return err
					}
					// Run
					return ClaimSpRewards(c, downloadMerkleProofs)
				},
			},
			{
				Name:      "update-operator-details",
				Aliases:   []string{"uod"},
				Usage:     "Update Operator name or Operator reward address",
				UsageText: "stader-cli node update-operator-details --operator-name --operator-reward-address",
				Flags: []cli.Flag{
					cli.StringFlag{
						Name:  "operator-name, on",
						Usage: "The new operator name",
					},
					cli.StringFlag{
						Name:  "operator-reward-address, ora",
						Usage: "New operator reward address",
					},
					cli.BoolFlag{
						Name:  "yes, y",
						Usage: "Automatically confirm claim of rewards",
					},
				},
				Action: func(c *cli.Context) error {

					operatorName := c.String("operator-name")
					if operatorName == "" {
						return fmt.Errorf("operator name can't be empty string")
					}
					operatorRewardAddress, err := cliutils.ValidateAddress("operator-reward-address", c.String("operator-reward-address"))
					if err != nil {
						return err
					}

					// Run
					return updateOperatorDetails(c, operatorName, operatorRewardAddress)
				},
			},
		},
	})
}
