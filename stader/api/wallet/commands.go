package wallet

// ROCKETPOOL-OWNED

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
		Usage:   "Manage the node wallet",
		Subcommands: []cli.Command{

			{
				Name:      "status",
				Aliases:   []string{"s"},
				Usage:     "Get the node wallet status",
				UsageText: "stader-cli api wallet status",
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
				Name:      "set-password",
				Aliases:   []string{"p"},
				Usage:     "Set the node wallet password",
				UsageText: "stader-cli api wallet set-password password",
				Action: func(c *cli.Context) error {

					// Validate args
					if err := cliutils.ValidateArgCount(c, 1); err != nil {
						return err
					}
					password, err := cliutils.ValidateNodePassword("wallet password", c.Args().Get(0))
					if err != nil {
						return err
					}

					// Run
					api.PrintResponse(setPassword(c, password))
					return nil

				},
			},

			{
				Name:      "init",
				Aliases:   []string{"i"},
				Usage:     "Initialize the node wallet",
				UsageText: "stader-cli api wallet init",
				Flags: []cli.Flag{
					cli.StringFlag{
						Name:  "derivation-path, d",
						Usage: "Specify the derivation path for the wallet.\nOmit this flag (or leave it blank) for the default of \"m/44'/60'/0'/0/%d\" (where %d is the index).\nSet this to \"ledgerLive\" to use Ledger Live's path of \"m/44'/60'/%d/0/0\".\nSet this to \"mew\" to use MyEtherWallet's path of \"m/44'/60'/0'/%d\".\nFor custom paths, simply enter them here.",
					},
				},
				Action: func(c *cli.Context) error {

					// Validate args
					if err := cliutils.ValidateArgCount(c, 0); err != nil {
						return err
					}

					// Run
					api.PrintResponse(initWallet(c))
					return nil

				},
			},

			{
				Name:      "export",
				Aliases:   []string{"e"},
				Usage:     "Export the node wallet in JSON format",
				UsageText: "stader-cli api wallet export",
				Action: func(c *cli.Context) error {

					// Validate args
					if err := cliutils.ValidateArgCount(c, 0); err != nil {
						return err
					}

					// Run
					api.PrintResponse(exportWallet(c))
					return nil

				},
			},

			{
				Name:      "purge",
				Usage:     "Deletes your node wallet, your validator keys, and restarts your Validator Client while preserving your chain data. WARNING: Only use this if you want to stop validating with this machine!",
				UsageText: "stader-cli api wallet purge",
				Action: func(c *cli.Context) error {

					// Validate args
					if err := cliutils.ValidateArgCount(c, 0); err != nil {
						return err
					}

					// Run
					api.PrintResponse(purge(c))
					return nil

				},
			},
		},
	})
}
