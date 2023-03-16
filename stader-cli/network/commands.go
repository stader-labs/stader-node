package network

import (
	"github.com/urfave/cli"
)

// Register commands
func RegisterCommands(app *cli.App, name string, aliases []string) {
	app.Commands = append(app.Commands, cli.Command{
		Name:        name,
		Aliases:     aliases,
		Usage:       "Manage Stader network parameters",
		Subcommands: []cli.Command{},
	})
}
