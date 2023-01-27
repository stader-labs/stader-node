package main

import (
	"fmt"
	"os"

	"github.com/urfave/cli"

	"github.com/stader-labs/stader-node/shared"
	apiutils "github.com/stader-labs/stader-node/shared/utils/api"
	"github.com/stader-labs/stader-node/stader/api"
	"github.com/stader-labs/stader-node/stader/node"
	"github.com/stader-labs/stader-node/stader/watchtower"
)

// Run
func main() {

	// Initialise application
	app := cli.NewApp()

	// Set application info
	app.Name = "Stader"
	app.Usage = "Stader service"
	app.Version = shared.RocketPoolVersion
	app.Authors = []cli.Author{
		{
			Name:  "Rahul Jaguste",
			Email: "rahul@staderlabs.com",
		},
	}
	app.Copyright = "(c) 2023 Stader labs"

	// Set application flags
	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:  "settings, s",
			Usage: "Stader service user config absolute `path`",
			Value: "/.stader/user-settings.yml",
		},
		cli.StringFlag{
			Name:  "storageAddress, a",
			Usage: "Stader storage contract `address`",
		},
		cli.StringFlag{
			Name:  "oneInchOracleAddress, o",
			Usage: "1inch exchange oracle contract `address`",
		},
		cli.StringFlag{
			Name:  "rplTokenAddress, t",
			Usage: "RPL token contract `address`",
		},
		cli.StringFlag{
			Name:  "rplFaucetAddress, f",
			Usage: "Stader RPL token faucet `address`",
		},
		cli.StringFlag{
			Name:  "password, p",
			Usage: "Stader wallet password file absolute `path`",
		},
		cli.StringFlag{
			Name:  "wallet, w",
			Usage: "Stader wallet file absolute `path`",
		},
		cli.StringFlag{
			Name:  "validatorKeychain, k",
			Usage: "Stader validator keychain absolute `path`",
		},
		cli.StringFlag{
			Name:  "eth1Provider, e",
			Usage: "Eth 1.0 provider `address`",
		},
		cli.StringFlag{
			Name:  "eth2Provider, b",
			Usage: "Eth 2.0 provider `address`",
		},
		cli.Float64Flag{
			Name:  "maxFee",
			Usage: "Desired max fee in gwei",
		},
		cli.Float64Flag{
			Name:  "maxPrioFee",
			Usage: "Desired max priority fee in gwei",
		},
		cli.Uint64Flag{
			Name:  "gasLimit, l",
			Usage: "Desired gas limit",
		},
		cli.StringFlag{
			Name:  "nonce",
			Usage: "Use this flag to explicitly specify the nonce that this transaction should use, so it can override an existing 'stuck' transaction",
		},
		cli.StringFlag{
			Name:  "metricsAddress, m",
			Usage: "Address to serve metrics on if enabled",
			Value: "0.0.0.0",
		},
		cli.UintFlag{
			Name:  "metricsPort, r",
			Usage: "Port to serve metrics on if enabled",
			Value: 9102,
		},
		cli.BoolFlag{
			Name:  "ignore-sync-check",
			Usage: "Set this to true if you already checked the sync status of the execution client(s) and don't need to re-check it for this command",
		},
		cli.BoolFlag{
			Name:  "force-fallbacks",
			Usage: "Set this to true if you know the primary EC or CC is offline and want to bypass its health checks, and just use the fallback EC and CC instead",
		},
	}

	// Register commands
	api.RegisterCommands(app, "api", []string{"a"})
	node.RegisterCommands(app, "node", []string{"n"})
	watchtower.RegisterCommands(app, "watchtower", []string{"w"})

	// Get command being run
	var commandName string
	app.Before = func(c *cli.Context) error {
		commandName = c.Args().First()
		return nil
	}

	// Run application
	if err := app.Run(os.Args); err != nil {
		if commandName == "api" {
			apiutils.PrintErrorResponse(err)
		} else {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}
	}

}
