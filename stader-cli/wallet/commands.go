package wallet

import (
	"github.com/urfave/cli"

	cliutils "github.com/stader-labs/stader-node/shared/utils/cli"
)

// Register commands
func RegisterCommands(app *cli.App, name string, aliases []string) {
	app.Commands = append(app.Commands, cli.Command{
		Name:    name,
		Aliases: aliases,
		Usage:   "Manage the node wallet",
		Subcommands: []cli.Command{

			{
				Name:      "status",
				Aliases:   []string{"s"},
				Usage:     "Get the node wallet status",
				UsageText: "stdr-cli wallet status",
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
				Name:      "init",
				Aliases:   []string{"i"},
				Usage:     "Initialize the node wallet",
				UsageText: "stdr-cli wallet init [options]",
				Flags: []cli.Flag{
					cli.StringFlag{
						Name:  "password, p",
						Usage: "The password to secure the wallet with (if not already set)",
					},
					cli.BoolFlag{
						Name:  "confirm-mnemonic, c",
						Usage: "Automatically confirm the mnemonic phrase",
					},
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

					// Validate flags
					if c.String("password") != "" {
						if _, err := cliutils.ValidateNodePassword("password", c.String("password")); err != nil {
							return err
						}
					}

					// Run
					return initWallet(c)

				},
			},

			{
				Name:      "export",
				Aliases:   []string{"e"},
				Usage:     "Export the node wallet in JSON format",
				UsageText: "stdr-cli wallet export",
				Action: func(c *cli.Context) error {

					// Validate args
					if err := cliutils.ValidateArgCount(c, 0); err != nil {
						return err
					}

					// Run
					return exportWallet(c)

				},
			},
		},
	})
}
