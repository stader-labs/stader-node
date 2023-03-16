package network

import (
	"github.com/urfave/cli"
)

// Register subcommands
func RegisterSubcommands(command *cli.Command, name string, aliases []string) {
	command.Subcommands = append(command.Subcommands, cli.Command{
		Name:        name,
		Aliases:     aliases,
		Usage:       "Manage Stader network parameters",
		Subcommands: []cli.Command{},
	})
}
