/*
This work is licensed and released under GNU GPL v3 or any other later versions.
The full text of the license is below/ found at <http://www.gnu.org/licenses/>

(c) 2023 Rocket Pool Pty Ltd. Modified under GNU GPL v3. [1.4.9]

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
					return getNodeStatus(c)

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
				UsageText: "stader-cli node update-socialize-el [options]",
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

					operatorName := c.String("operator-name")

					// Run
					return registerNode(c, operatorName)
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
				Name:      "send-el-rewards",
				Aliases:   []string{"wer"},
				Usage:     "Send all Execution Layer rewards to the operator claim vault. This only includes non-socializing pool rewards",
				UsageText: "stader-cli node send-el-rewards",
				Flags: []cli.Flag{cli.BoolFlag{
					Name:  "yes, y",
					Usage: "Automatically confirm EL rewards send to operator claim vault",
				}},
				Action: func(c *cli.Context) error {
					// Run
					return SendElRewards(c)
				},
			},
			{
				Name:      "claim-rewards",
				Aliases:   []string{"wer"},
				Usage:     "Claim rewards from claim vault to the operator reward address",
				UsageText: "stader-cli node claim-rewards",
				Flags: []cli.Flag{cli.BoolFlag{
					Name:  "yes, y",
					Usage: "Automatically confirm rewards transfer to operator reward address",
				}},
				Action: func(c *cli.Context) error {
					// Run
					return ClaimRewards(c)
				},
			},
			{
				Name:      "withdraw-sd-collateral",
				Aliases:   []string{"sef"},
				Usage:     "Withdraw Sd collateral",
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
				UsageText: "stader-cli node claim-sp-rewards",
				Flags: []cli.Flag{
					cli.BoolFlag{
						Name:  "yes, y",
						Usage: "Automatically confirm claim of rewards",
					},
				},
				Action: func(c *cli.Context) error {

					if err := cliutils.ValidateArgCount(c, 0); err != nil {
						return err
					}

					// Run
					return ClaimSpRewards(c)
				},
			},
			{
				Name:      "update-operator-name",
				Aliases:   []string{"uon"},
				Usage:     "Update Operator name",
				UsageText: "stader-cli node update-operator-name --operator-name",
				Flags: []cli.Flag{
					cli.StringFlag{
						Name:  "operator-name, on",
						Usage: "The new operator name",
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

					// Run
					return updateOperatorName(c, operatorName)
				},
			},
			{
				Name:      "set-reward-address",
				Aliases:   []string{"sra"},
				Usage:     "Set operator reward address",
				UsageText: "stader-cli node set-reward-address --operator-reward-address",
				Flags: []cli.Flag{
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

					operatorRewardAddress, err := cliutils.ValidateAddress("operator-reward-address", c.String("operator-reward-address"))
					if err != nil {
						return err
					}

					// Run
					return SetRewardAddress(c, operatorRewardAddress)
				},
			},
			{
				Name:      "approve-sd",
				Aliases:   []string{"k"},
				Usage:     "Approve SD against the node",
				UsageText: "stader-cli node approve-sd [options]",
				Flags: []cli.Flag{
					cli.StringFlag{
						Name:  "amount, a",
						Usage: "The amount of SD to approve",
					},
					cli.BoolFlag{
						Name:  "yes, y",
						Usage: "Automatically confirm SD approve",
					},
				},
				Action: func(c *cli.Context) error {

					if _, err := cliutils.ValidatePositiveEthAmount("sd deposit amount", c.String("amount")); err != nil {
						return err
					}

					// Run
					return nodeApproveSd(c)
				},
			},
		},
	})
}
