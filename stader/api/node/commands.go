package node

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
		Usage:   "Manage the node",
		Subcommands: []cli.Command{

			{
				Name:      "status",
				Aliases:   []string{"s"},
				Usage:     "Get the node's status",
				UsageText: "stdr-cli api node status",
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
				Name:      "sync",
				Aliases:   []string{"y"},
				Usage:     "Get the sync progress of the eth1 and eth2 clients",
				UsageText: "stdr-cli api node sync",
				Action: func(c *cli.Context) error {

					// Validate args
					if err := cliutils.ValidateArgCount(c, 0); err != nil {
						return err
					}

					// Run
					api.PrintResponse(getSyncProgress(c))
					return nil

				},
			},

			{
				Name:      "can-register",
				Usage:     "Check whether the node can be registered with stdr=",
				UsageText: "stdr-cli api node can-register timezone-location",
				Action: func(c *cli.Context) error {

					// Validate args
					if err := cliutils.ValidateArgCount(c, 3); err != nil {
						return err
					}

					if err := cliutils.ValidateArgCount(c, 3); err != nil {
						return err
					}

					operatorName := c.Args().Get(0)

					operatorRewardAddress, err := cliutils.ValidateAddress("operator-reward-address", c.Args().Get(1))
					if err != nil {
						return err
					}

					socializeMev, err := cliutils.ValidateBool("socialize-mev", c.Args().Get(2))

					// Run
					api.PrintResponse(canRegisterNode(c, operatorName, operatorRewardAddress, socializeMev))
					return nil

				},
			},
			{
				Name:      "register",
				Aliases:   []string{"r"},
				Usage:     "Register the node with Stader",
				UsageText: "stdr-cli api node register operator-name operator-reward-address socialize-mev",
				Action: func(c *cli.Context) error {

					// Validate args
					if err := cliutils.ValidateArgCount(c, 3); err != nil {
						return err
					}

					operatorName := c.Args().Get(0)

					operatorRewardAddress, err := cliutils.ValidateAddress("operator-reward-address", c.Args().Get(1))
					if err != nil {
						return err
					}

					socializeMev, err := cliutils.ValidateBool("socialize-mev", c.Args().Get(2))

					// Run
					api.PrintResponse(registerNode(c, operatorName, operatorRewardAddress, socializeMev))
					return nil

				},
			},

			{
				Name:      "can-node-deposit-sd",
				Usage:     "Check whether the node can stake RPL",
				UsageText: "stdr-cli api node can-node-deposit-sd amount",
				Action: func(c *cli.Context) error {

					// Validate args
					if err := cliutils.ValidateArgCount(c, 1); err != nil {
						return err
					}
					amountWei, err := cliutils.ValidatePositiveWeiAmount("stake amount", c.Args().Get(0))
					if err != nil {
						return err
					}

					// Run
					api.PrintResponse(canNodeDepositSd(c, amountWei))
					return nil

				},
			},
			{
				Name:      "deposit-sd-approve-sd",
				Aliases:   []string{"k1"},
				Usage:     "Approve SD for staking against the node",
				UsageText: "stdr-cli api node deposit-sd-approve-sd amount",
				Action: func(c *cli.Context) error {

					// Validate args
					if err := cliutils.ValidateArgCount(c, 1); err != nil {
						return err
					}
					amountWei, err := cliutils.ValidatePositiveWeiAmount("stake amount", c.Args().Get(0))
					if err != nil {
						return err
					}

					// Run
					api.PrintResponse(approveSd(c, amountWei))
					return nil

				},
			},
			{
				Name:      "wait-and-deposit-sd",
				Aliases:   []string{"k2"},
				Usage:     "Deposit SD against the node, waiting for approval tx-hash to be included in a block first",
				UsageText: "stdr-cli api node wait-and-stake-rpl amount tx-hash",
				Action: func(c *cli.Context) error {

					// Validate args
					if err := cliutils.ValidateArgCount(c, 2); err != nil {
						return err
					}
					amountWei, err := cliutils.ValidatePositiveWeiAmount("stake amount", c.Args().Get(0))
					if err != nil {
						return err
					}
					hash, err := cliutils.ValidateTxHash("tx-hash", c.Args().Get(1))
					if err != nil {
						return err
					}

					// Run
					api.PrintResponse(waitForApprovalAndDepositSd(c, amountWei, hash))
					return nil

				},
			},
			{
				Name:      "get-deposit-sd-approval-gas",
				Usage:     "Estimate the gas cost of new SD interaction approval",
				UsageText: "stdr-cli api node get-deposit-sd-approval-gas",
				Action: func(c *cli.Context) error {

					// Validate args
					if err := cliutils.ValidateArgCount(c, 1); err != nil {
						return err
					}
					amountWei, err := cliutils.ValidatePositiveWeiAmount("approve amount", c.Args().Get(0))
					if err != nil {
						return err
					}

					// Run
					api.PrintResponse(getDepositSdApprovalGas(c, amountWei))
					return nil

				},
			},
			{
				Name:      "deposit-sd-allowance",
				Usage:     "Get the node's SD allowance for the collateral contract",
				UsageText: "stdr-cli api node deposit-sd-allowance",
				Action: func(c *cli.Context) error {

					// Validate args
					if err := cliutils.ValidateArgCount(c, 0); err != nil {
						return err
					}

					// Run
					api.PrintResponse(allowanceSd(c))
					return nil

				},
			},
			{
				Name:      "deposit-sd",
				Aliases:   []string{"k3"},
				Usage:     "Deposit SD against the node",
				UsageText: "stdr-cli api node deposit-sd amount",
				Action: func(c *cli.Context) error {

					// Validate args
					if err := cliutils.ValidateArgCount(c, 1); err != nil {
						return err
					}
					amountWei, err := cliutils.ValidatePositiveWeiAmount("stake amount", c.Args().Get(0))
					if err != nil {
						return err
					}

					// Run
					api.PrintResponse(depositSdAsCollateral(c, amountWei))
					return nil

				},
			},
			{
				Name:      "can-deposit",
				Usage:     "Check whether the node can make a deposit",
				UsageText: "stdr-cli api node can-deposit amount min-fee salt",
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

					submit, err := cliutils.ValidateBool("submit", c.Args().Get(3))
					if err != nil {
						return err
					}

					api.PrintResponse(canNodeDeposit(c, amountWei, salt, numValidators, submit))

					return nil

				},
			},
			{
				Name:      "deposit",
				Aliases:   []string{"d"},
				Usage:     "Make a deposit and create a validator, or just make and sign the transaction (when submit = false)",
				UsageText: "stdr api node deposit amount salt submit",
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

					submit, err := cliutils.ValidateBool("submit", c.Args().Get(3))
					if err != nil {
						return err
					}

					// Run
					response, err := nodeDeposit(c, amountWei, salt, numValidators, submit)
					if submit {
						api.PrintResponse(response, err)
					}
					return nil

				},
			},

			{
				Name:      "can-send",
				Usage:     "Check whether the node can send ETH or tokens to an address",
				UsageText: "stdr-cli api node can-send amount token",
				Action: func(c *cli.Context) error {

					// Validate args
					if err := cliutils.ValidateArgCount(c, 2); err != nil {
						return err
					}
					amountWei, err := cliutils.ValidatePositiveWeiAmount("send amount", c.Args().Get(0))
					if err != nil {
						return err
					}
					token, err := cliutils.ValidateTokenType("token type", c.Args().Get(1))
					if err != nil {
						return err
					}

					// Run
					api.PrintResponse(canNodeSend(c, amountWei, token))
					return nil

				},
			},
			{
				Name:      "send",
				Aliases:   []string{"n"},
				Usage:     "Send ETH or tokens from the node account to an address",
				UsageText: "stdr-cli api node send amount token to",
				Action: func(c *cli.Context) error {

					// Validate args
					if err := cliutils.ValidateArgCount(c, 3); err != nil {
						return err
					}
					amountWei, err := cliutils.ValidatePositiveWeiAmount("send amount", c.Args().Get(0))
					if err != nil {
						return err
					}
					token, err := cliutils.ValidateTokenType("token type", c.Args().Get(1))
					if err != nil {
						return err
					}
					toAddress, err := cliutils.ValidateAddress("to address", c.Args().Get(2))
					if err != nil {
						return err
					}

					// Run
					api.PrintResponse(nodeSend(c, amountWei, token, toAddress))
					return nil

				},
			},

			{
				Name:      "get-contracts-info",
				Usage:     "Get information about the deposit contract and stdr contract on the current network",
				UsageText: "stdr-cli api node get-contracts-info",
				Action: func(c *cli.Context) error {

					// Validate args
					if err := cliutils.ValidateArgCount(c, 0); err != nil {
						return err
					}

					// Run
					api.PrintResponse(getContractsInfo(c))
					return nil

				},
			},

			{
				Name:      "sign",
				Usage:     "Signs a transaction with the node's private key. The TX must be serialized as a hex string.",
				UsageText: "stdr-cli api node sign tx",
				Action: func(c *cli.Context) error {

					// Validate args
					if err := cliutils.ValidateArgCount(c, 1); err != nil {
						return err
					}

					data := c.Args().Get(0)

					// Run
					api.PrintResponse(sign(c, data))
					return nil

				},
			},

			{
				Name:      "sign-message",
				Usage:     "Signs an arbitrary message with the node's private key.",
				UsageText: "stdr-cli api node sign-message 'message'",
				Action: func(c *cli.Context) error {

					// Validate args
					if err := cliutils.ValidateArgCount(c, 1); err != nil {
						return err
					}

					message := c.Args().Get(0)

					// Run
					api.PrintResponse(signMessage(c, message))
					return nil

				},
			},

			{
				Name:      "debug-exit",
				Usage:     "Get exit msg info for a given validator",
				UsageText: "stdr-cli api node debug-exit validator-index",
				Action: func(c *cli.Context) error {

					// Validate args
					if err := cliutils.ValidateArgCount(c, 1); err != nil {
						return err
					}

					valIndex := c.Args().Get(0)

					validatorIndex, err := cliutils.ValidateBigInt("validator-index", valIndex)
					if err != nil {
						return err
					}

					// Run
					api.PrintResponse(DebugExit(c, validatorIndex))
					return nil

				},
			},
		},
	})
}
