package main

import (
	"github.com/stader-labs/stader-node/stader/guardian"
	"github.com/stader-labs/stader-node/stader/node"
	"github.com/urfave/cli"
)

func main() {
	// Initialise application
	app := cli.NewApp()
	guardian.RegisterCommands(app, "guardian", []string{"w"})
	node.RegisterCommands(app, "node", []string{"w"})
}
