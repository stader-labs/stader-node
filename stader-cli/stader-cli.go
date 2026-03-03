/*
This work is licensed and released under GNU GPL v3 or any other later versions.
The full text of the license is below/ found at <http://www.gnu.org/licenses/>

(c) 2023 Rocket Pool Pty Ltd. Modified under GNU GPL v3. [1.5.0]

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
package main

import (
	"fmt"
	"os"

	"github.com/mitchellh/go-homedir"
	"github.com/stader-labs/stader-node/shared"
	cliutils "github.com/stader-labs/stader-node/shared/utils/cli"
	"github.com/stader-labs/stader-node/stader-cli/node"
	"github.com/stader-labs/stader-node/stader-cli/service"
	"github.com/stader-labs/stader-node/stader-cli/validator"
	"github.com/stader-labs/stader-node/stader-cli/wallet"
	"github.com/urfave/cli"
)

// Run
func main() {

	// Add logo to application help template
	cli.AppHelpTemplate = fmt.Sprintf(`%s
   _____ _            _             _           _         
  / ____| |          | |           | |         | |        
 | (___ | |_ __ _  __| | ___ _ __  | |     __ _| |__  ___ 
  \___ \| __/ _' |/ _' |/ _ \ '__| | |    / _' | '_ \/ __|
  ____) | || (_| | (_| |  __/ |    | |___| (_| | |_) \__ \
 |_____/ \__\__,_|\__,_|\___|_|    |______\__,_|_.__/|___/
                                                          
%s`, "\033[34m", "\033[0m"+cli.AppHelpTemplate)

	// Initialise application
	app := cli.NewApp()

	// Set application info
	app.Name = "Stader-node"
	app.Usage = "Stader Node CLI"
	app.Version = shared.StaderVersion
	// app.Description = ``

	// Set application flags
	app.Flags = []cli.Flag{
		cli.BoolFlag{
			Name:  "allow-root, r",
			Usage: "Allow Stader node to be run as the root user",
		},
		cli.StringFlag{
			Name:  "config-path, c",
			Usage: "Stader node config asset `path`",
			Value: "~/.stader",
		},
		cli.StringFlag{
			Name:  "daemon-path, d",
			Usage: "Interact with a Stader node service daemon at a `path` on the host OS, running outside of docker",
		},
		cli.Float64Flag{
			Name:  "maxFee, f",
			Usage: "The max fee (including the priority fee) you want a transaction to cost, in gwei",
		},
		cli.Float64Flag{
			Name:  "maxPrioFee, i",
			Usage: "The max priority fee you want a transaction to use, in gwei",
		},
		cli.Uint64Flag{
			Name:  "gasLimit",
			Usage: "Custom gas limit, using gas estimate if not specified",
		},
		cli.StringFlag{
			Name:  "nonce",
			Usage: "Use this flag to explicitly specify the nonce that this transaction should use, so it can override an existing 'stuck' transaction",
		},
		cli.BoolFlag{
			Name:  "debug",
			Usage: "Enable debug printing of API commands",
		},
		cli.BoolFlag{
			Name: "secure-session, s",
			Usage: "Some commands may print sensitive information to your terminal. " +
				"Use this flag when nobody can see your screen to allow sensitive data to be printed without prompting",
		},
	}
	app.Authors = []cli.Author{
		{
			Name: "David Rugendyke",
			// Email: "david@rocketpool.net",
		},
		{
			Name: "Jake Pospischil",
			// Email: "jake@rocketpool.net",
		},
		{
			Name: "Joe Clapis",
			// Email: "joe@rocketpool.net",
		},
		{
			Name: "Kane Wallmann",
			// Email: "kane@rocketpool.net",
		},
	}
	// Register commands

	// Get the config path from the arguments (or use the default)
	configPath := "~/.stader"
	for index, arg := range os.Args {
		if arg == "-c" || arg == "--config-path" {
			if len(os.Args)-1 == index {
				fmt.Fprintf(os.Stderr, "Expected config path after %s but none was given.\n", arg)
				os.Exit(1)
			}
			configPath = os.Args[index+1]
		}
	}

	// Get and parse the config file
	configFile := fmt.Sprintf("%s/%s", configPath, "user-settings.yml")
	_, err := homedir.Expand(configFile)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to get the global config file path: %s\n", err.Error())
		os.Exit(1)
	}

	node.RegisterCommands(app, "node", []string{"n"})
	service.RegisterCommands(app, "service", []string{"s"})
	wallet.RegisterCommands(app, "wallet", []string{"w"})
	validator.RegisterCommands(app, "validator", []string{"v"})
	app.Commands = append(app.Commands, cli.Command{
		Name:    "license",
		Aliases: []string{"l"},
		Usage:   "Show the license",
		Action: func(c *cli.Context) error {
			// fmt.Println(GPLv3)
			return nil
		},
	})

	app.Commands = append(app.Commands, cli.Command{
		Name:    "copyright",
		Aliases: []string{"c"},
		Usage:   "Show the copyright",
		Action: func(c *cli.Context) error {
			fmt.Println("(c) 2023 Rocket Pool Pty Ltd. Modified under GNU GPL v3. [1.5.0].\n(c) 2023 Stakeinfra Technologies Inc.")
			return nil
		},
	})

	app.Before = func(c *cli.Context) error {
		// Check user ID
		if os.Getuid() == 0 && !c.GlobalBool("allow-root") {
			fmt.Fprintln(os.Stderr, "Stader node should not be run as root. Please try again without 'sudo'.")
			fmt.Fprintln(os.Stderr, "If you want to run Stader node as root anyway, use the '--allow-root' option to override this warning.")
			os.Exit(1)
		}

		return nil
	}

	// Run application
	fmt.Println("")
	if err := app.Run(os.Args); err != nil {
		cliutils.PrettyPrintError(err)
	}
	fmt.Println("")

}
