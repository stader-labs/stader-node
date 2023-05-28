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
package validator

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
		Usage:   "Validator related commands",
		Subcommands: []cli.Command{
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
				Name:      "exit-validator",
				Aliases:   []string{"e"},
				Usage:     "Exit validator",
				UsageText: "stader-cli node exit-validator --validator-pub-key",
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
				Name:      "send-cl-rewards",
				Aliases:   []string{"wcr"},
				Usage:     "Send all Consensus Layer rewards to the operator claim vault",
				UsageText: "stader-cli node send-cl-rewards --validator-pub-key",
				Flags: []cli.Flag{
					cli.StringFlag{
						Name:  "validator-pub-key, vpk",
						Usage: "Public key of the validator whose CL rewards we want to send to operator claim vault",
					},
					cli.BoolFlag{
						Name:  "yes, y",
						Usage: "Automatically confirm CL rewards send",
					},
				},
				Action: func(c *cli.Context) error {

					validatorPubKey, err := cliutils.ValidatePubkey("validator-pub-key", c.String("validator-pub-key"))
					if err != nil {
						return err
					}
					// Run
					return SendClRewards(c, validatorPubKey)
				},
			},
			{
				Name:      "status",
				Aliases:   []string{"s"},
				Usage:     "Validator Status",
				UsageText: "stader-cli validator status",
				Flags:     []cli.Flag{},
				Action: func(c *cli.Context) error {

					// Run
					return getValidatorStatus(c)
				},
			},
			{
				Name:      "export",
				Aliases:   []string{"e"},
				Usage:     "Export validator info to a csv file",
				UsageText: "stader-cli validator export",
				Flags:     []cli.Flag{},
				Action: func(c *cli.Context) error {

					// Run
					return exportValidatorStatus(c)
				},
			},
		},
	})
}
