package network

import (
	"github.com/urfave/cli"

	cliutils "github.com/stader-labs/stader-node/shared/utils/cli"
)

// Register commands
func RegisterCommands(app *cli.App, name string, aliases []string) {
	app.Commands = append(app.Commands, cli.Command{
		Name:    name,
		Aliases: aliases,
		Usage:   "Manage Stader network parameters",
		Subcommands: []cli.Command{
			{
				Name:      "stats",
				Aliases:   []string{"s"},
				Usage:     "Get stats about the Stader network and its tokens",
				UsageText: "Stader network stats",
				Action: func(c *cli.Context) error {

					// Validate args
					if err := cliutils.ValidateArgCount(c, 0); err != nil {
						return err
					}

					// Run
					return getStats(c)

				},
			},
			{
				Name:      "node-fee",
				Aliases:   []string{"f"},
				Usage:     "Get the current network node commission rate",
				UsageText: "rocketpool network node-fee",
				Action: func(c *cli.Context) error {

					// Validate args
					if err := cliutils.ValidateArgCount(c, 0); err != nil {
						return err
					}

					// Run
					return getNodeFee(c)

				},
			},
			{
				Name:      "get-contracts-info",
				Aliases:   []string{"c"},
				Usage:     "Get the current network contracts info",
				UsageText: "stader-cli network get-contracts-info",
				Action: func(c *cli.Context) error {

					// Validate args
					if err := cliutils.ValidateArgCount(c, 0); err != nil {
						return err
					}

					// Run
					return getNodeFee(c)

				},
			},
		},
	})
}
