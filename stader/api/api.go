package api

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/stader-labs/stader-node/stader/api/debug"
	"github.com/urfave/cli"

	"github.com/rocket-pool/rocketpool-go/utils"
	"github.com/stader-labs/stader-node/shared/services"
	apitypes "github.com/stader-labs/stader-node/shared/types/api"
	"github.com/stader-labs/stader-node/shared/utils/api"
	cliutils "github.com/stader-labs/stader-node/shared/utils/cli"
	"github.com/stader-labs/stader-node/stader/api/network"
	"github.com/stader-labs/stader-node/stader/api/node"
	apiservice "github.com/stader-labs/stader-node/stader/api/service"
	apivalidator "github.com/stader-labs/stader-node/stader/api/validator"
	"github.com/stader-labs/stader-node/stader/api/wallet"
)

// Waits for an auction transaction
func waitForTransaction(c *cli.Context, hash common.Hash) (*apitypes.APIResponse, error) {

	rp, err := services.GetRocketPool(c)
	if err != nil {
		return nil, err
	}

	// Response
	response := apitypes.APIResponse{}
	_, err = utils.WaitForTransaction(rp.Client, hash)
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
	network.RegisterSubcommands(&command, "network", []string{"e"})
	node.RegisterSubcommands(&command, "node", []string{"n"})
	wallet.RegisterSubcommands(&command, "wallet", []string{"w"})
	apiservice.RegisterSubcommands(&command, "service", []string{"s"})
	debug.RegisterSubcommands(&command, "debug", []string{"d"})
	apivalidator.RegisterSubcommands(&command, "validator", []string{"v"})

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
