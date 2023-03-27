/*
This work is licensed and released under GNU GPL v3 or any other later versions.
The full text of the license is below/ found at <http://www.gnu.org/licenses/>

(c) 2023 Rocket Pool Pty Ltd. Modified under GNU GPL v3. [0.3.0-beta]

This program is free software: you can redistribute it and/or modify
it under the terms of the GNU General Public License as published by
the Free Software Foundation, either version 3 of the License, or
(at your option) any later version.

This program is distributed in the hope that it will be useful,
but WITHOUT ANY WARRANTY; without even the implied warranty of
MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
GNU General Public License for more details.

You should have received a copy of the GNU General Public License
along with this program.  If not, see <http://www.gnu.org/licenses/>.
*/
package wallet

import (
	"fmt"
	"os"

	"github.com/urfave/cli"

	"github.com/stader-labs/stader-node/shared/services/stader"
	cliutils "github.com/stader-labs/stader-node/shared/utils/cli"
)

func exportWallet(c *cli.Context) error {

	staderCLient, err := stader.NewClientFromCtx(c)
	if err != nil {
		return err
	}
	defer staderCLient.Close()

	// Get & check wallet status
	status, err := staderCLient.WalletStatus()
	if err != nil {
		return err
	}
	if !status.WalletInitialized {
		fmt.Println("The node wallet is not initialized.")
		return nil
	}

	if !c.GlobalBool("secure-session") {
		// Check if stdout is interactive
		stat, err := os.Stdout.Stat()
		if err != nil {
			fmt.Fprintf(os.Stderr, "An error occured while determining whether or not the output is a tty: %w\n"+
				"Use \"stader-cli --secure-session wallet export\" to bypass.\n", err)
			os.Exit(1)
		}

		if (stat.Mode()&os.ModeCharDevice) == os.ModeCharDevice &&
			!cliutils.ConfirmSecureSession("Exporting a wallet will print sensitive information to your screen.") {
			return nil
		}
	}

	// Export wallet
	export, err := staderCLient.ExportWallet()
	if err != nil {
		return err
	}

	// Print wallet & return
	fmt.Println("Node account private key:")
	fmt.Println("")
	fmt.Println(export.AccountPrivateKey)
	fmt.Println("")
	fmt.Println("Wallet password:")
	fmt.Println("")
	fmt.Println(export.Password)
	fmt.Println("")
	fmt.Println("Wallet file:")
	fmt.Println("============")
	fmt.Println("")
	fmt.Println(export.Wallet)
	fmt.Println("")
	fmt.Println("============")
	return nil

}
