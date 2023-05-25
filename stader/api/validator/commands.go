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
	"github.com/urfave/cli"

	"github.com/stader-labs/stader-node/shared/utils/api"
	cliutils "github.com/stader-labs/stader-node/shared/utils/cli"
)

// Register subcommands
func RegisterSubcommands(command *cli.Command, name string, aliases []string) {
	command.Subcommands = append(command.Subcommands, cli.Command{
		Name:    name,
		Aliases: aliases,
		Usage:   "validator related commands",
		Subcommands: []cli.Command{

			{
				Name:      "can-deposit",
				Usage:     "Check whether the node can make a deposit to create a validator",
				UsageText: "stader-cli api node can-deposit amount salt num-validators reload-keys",
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

					reloadKeys, err := cliutils.ValidateBool("reload-keys", c.Args().Get(3))
					if err != nil {
						return err
					}

					api.PrintResponse(canNodeDeposit(c, amountWei, salt, numValidators, reloadKeys))

					return nil

				},
			},
			{
				Name:      "deposit",
				Aliases:   []string{"d"},
				Usage:     "Make a deposit and create a validator",
				UsageText: "stader-cli api node deposit amount salt num-validators reload-keys",
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

					reloadKeys, err := cliutils.ValidateBool("reload-keys", c.Args().Get(3))
					if err != nil {
						return err
					}

					// Run
					response, err := nodeDeposit(c, amountWei, salt, numValidators, reloadKeys)
					api.PrintResponse(response, err)

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
				Name:      "can-send-cl-rewards",
				Usage:     "Can send cl rewards of a validator to the operator claim vault",
				UsageText: "stader-cli api node can-send-cl-rewards --validator-pub-key",
				Action: func(c *cli.Context) error {

					validatorPubKey, err := cliutils.ValidatePubkey("validator-pub-key", c.Args().Get(0))
					if err != nil {
						return err
					}

					api.PrintResponse(CanSendClRewards(c, validatorPubKey))
					return nil

				},
			},
			{
				Name:      "send-cl-rewards",
				Usage:     "Send cl rewards of a validator to the operator claim vault",
				UsageText: "stader-cli api node send-cl-rewards --validator-pub-key",
				Action: func(c *cli.Context) error {

					validatorPubKey, err := cliutils.ValidatePubkey("validator-pub-key", c.Args().Get(0))
					if err != nil {
						return err
					}

					api.PrintResponse(SendClRewards(c, validatorPubKey))
					return nil

				},
			},
		},
	})
}
