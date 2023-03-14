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
				Name:      "register",
				Aliases:   []string{"r"},
				Usage:     "Register the node with stader",
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
					cli.StringFlag{
						Name:  "socialize-mev, sm",
						Usage: "Should Mev be socialized (will default to true, can be be only true or false)",
					},
				},
				Action: func(c *cli.Context) error {

					// Validate args
					fmt.Printf("Operator name is %s\n", c.String("operator-name"))
					fmt.Printf("Operator reward address is %s\n", c.String("operator-reward-address"))
					fmt.Printf("socialize mev is %s\n", c.String("socialize-mev"))

					fmt.Printf("c is %v\n", c.Args())

					// Validate flags
					if c.String("operator-name") == "" {
						return fmt.Errorf("operator-name is required")
					}

					if c.String("operator-reward-address") != "" {
						if _, err := cliutils.ValidateAddress("operator-reward-address", c.String("operator-reward-address")); err != nil {
							return err
						}
					}

					if c.String("socialize-mev") != "" && c.String("socialize-mev") != "true" && c.String("socialize-mev") != "false" {
						return fmt.Errorf("invalid value for socialize mev, it should be exactly true or false")
					}

					// Run
					return registerNode(c)

				},
			},
			{
				Name:      "deposit-sd",
				Aliases:   []string{"k"},
				Usage:     "Deposit SD against the node",
				UsageText: "stader-cli node stake-rpl [options]",
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

					// Validate args
					if err := cliutils.ValidateArgCount(c, 0); err != nil {
						return err
					}

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

					// Validate args
					if err := cliutils.ValidateArgCount(c, 0); err != nil {
						return err
					}

					fmt.Printf("num-validator is %d\n", c.Uint64("num-validators"))
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
				Name:      "debug-exit",
				Aliases:   []string{"c"},
				Usage:     "get the debug exit info",
				UsageText: "stader-cli node debug-exit index",
				Flags: []cli.Flag{
					cli.Uint64Flag{
						Name:  "validator-index, vi",
						Usage: "Validator index for whom we want to generate the debug exit",
					},
					cli.Uint64Flag{
						Name:  "epoch-delta, ed",
						Usage: "Delta to add to the epoch",
					},
				},
				Action: func(c *cli.Context) error {

					//// Validate args
					//if err := cliutils.ValidateArgCount(c, 1); err != nil {
					//	return err
					//}
					index := c.Uint64("validator-index")
					fmt.Printf("index is %d\n", index)
					epochDelta := c.Uint64("epoch-delta")
					fmt.Printf("epoch delta is %d\n", epochDelta)

					// Run
					return debugExitMsg(c, index, epochDelta)
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
				},
				Action: func(c *cli.Context) error {

					//// Validate args
					//if err := cliutils.ValidateArgCount(c, 1); err != nil {
					//	return err
					//}
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
					cli.Uint64Flag{
						Name:  "validator-pub-key, vpk",
						Usage: "Validator index for whom we want to generate the debug exit",
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
		},
	})
}
