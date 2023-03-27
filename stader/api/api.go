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
package api

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/urfave/cli"

	"github.com/stader-labs/stader-node/shared/services"
	apitypes "github.com/stader-labs/stader-node/shared/types/api"
	"github.com/stader-labs/stader-node/shared/utils/api"
	cliutils "github.com/stader-labs/stader-node/shared/utils/cli"
	"github.com/stader-labs/stader-node/stader-lib/utils"
	"github.com/stader-labs/stader-node/stader/api/node"
	apiservice "github.com/stader-labs/stader-node/stader/api/service"
	"github.com/stader-labs/stader-node/stader/api/wallet"
)

// Waits for an auction transaction
func waitForTransaction(c *cli.Context, hash common.Hash) (*apitypes.APIResponse, error) {

	prn, err := services.GetPermissionlessNodeRegistry(c)
	if err != nil {
		return nil, err
	}

	// Response
	response := apitypes.APIResponse{}
	_, err = utils.WaitForTransaction(prn.Client, hash)
	if err != nil {
		return nil, err
	}

	// Return response
	return &response, nil

}

// Register commands
func RegisterCommands(app *cli.App, name string, aliases []string) {

	// CLI command
	command := cli.Command{
		Name:        name,
		Aliases:     aliases,
		Usage:       "Run Stader API commands",
		Subcommands: []cli.Command{},
	}

	// Don't show help message for api errors because of JSON serialisation
	command.OnUsageError = func(context *cli.Context, err error, isSubcommand bool) error {
		return err
	}

	// Register subcommands
	node.RegisterSubcommands(&command, "node", []string{"n"})
	wallet.RegisterSubcommands(&command, "wallet", []string{"w"})
	apiservice.RegisterSubcommands(&command, "service", []string{"s"})

	// Append a general wait-for-transaction command to support async operations
	command.Subcommands = append(command.Subcommands, cli.Command{
		Name:      "wait",
		Aliases:   []string{"t"},
		Usage:     "Wait for a transaction to complete",
		UsageText: "stader api wait tx-hash",
		Action: func(c *cli.Context) error {
			// Validate args
			if err := cliutils.ValidateArgCount(c, 1); err != nil {
				return err
			}
			hash, err := cliutils.ValidateTxHash("tx-hash", c.Args().Get(0))
			if err != nil {
				return err
			}

			// Run
			api.PrintResponse(waitForTransaction(c, hash))
			return nil
		},
	})

	// Register CLI command
	app.Commands = append(app.Commands, command)

}
