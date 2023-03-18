//go:build !windows
// +build !windows

package term

// ROCKETPOOL-OWNED

import (
	"os"
	"os/exec"
)

// Clear terminal output
func Clear() error {
	cmd := exec.Command("clear")
	cmd.Stdout = os.Stdout
	return cmd.Run()
}
