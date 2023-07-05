/*
This work is licensed and released under GNU GPL v3 or any other later versions.
The full text of the license is below/ found at <http://www.gnu.org/licenses/>

(c) 2023 Rocket Pool Pty Ltd. Modified under GNU GPL v3. [1.2.0]

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

	"github.com/urfave/cli"

	"github.com/stader-labs/stader-node/shared/services/stader"
	cliutils "github.com/stader-labs/stader-node/shared/utils/cli"
	"github.com/stader-labs/stader-node/shared/utils/term"
)

func initWallet(c *cli.Context) error {

	staderClient, err := stader.NewClientFromCtx(c)
	if err != nil {
		return err
	}
	defer staderClient.Close()

	// Get & check wallet status
	status, err := staderClient.WalletStatus()
	if err != nil {
		return err
	}
	if status.WalletInitialized {
		fmt.Println("The node wallet is already initialized.")
		return nil
	}

	// Prompt for user confirmation before printing sensitive information
	if !(c.GlobalBool("secure-session") ||
		cliutils.ConfirmSecureSession("Creating a wallet will print sensitive information to your screen.")) {
		return nil
	}

	// Set password if not set
	if !status.PasswordSet {
		var password string
		if c.String("password") != "" {
			password = c.String("password")
		} else {
			password = promptPassword()
		}
		if _, err := staderClient.SetPassword(password); err != nil {
			return err
		}
	}

	// Get the derivation path
	derivationPath := c.String("derivation-path")
	if derivationPath != "" {
		fmt.Printf("Using a custom derivation path (%s).\n\n", derivationPath)
	}

	// Initialize wallet
	response, err := staderClient.InitWallet(derivationPath)
	if err != nil {
		return err
	}

	// Print mnemonic
	fmt.Println("Your mnemonic phrase to recover your wallet is printed below. It can be used to recover your node account and validator keys if they are lost.")
	fmt.Println("Record this phrase somewhere secure and private. Do not share it with anyone as it will give them control of your node account and validators.")
	fmt.Println("==============================================================================================================================================")
	fmt.Println("")
	fmt.Println(response.Mnemonic)
	fmt.Println("")
	fmt.Println("==============================================================================================================================================")
	fmt.Println("")

	// Confirm mnemonic
	if !c.Bool("confirm-mnemonic") {
		confirmMnemonic(response.Mnemonic)
	}

	// Do a recover to save the wallet
	recoverResponse, err := staderClient.RecoverWallet(response.Mnemonic, true, derivationPath, 0)
	if err != nil {
		return fmt.Errorf("error saving wallet: %w", err)
	}

	// Sanity check the addresses
	if recoverResponse.AccountAddress != response.AccountAddress {
		return fmt.Errorf("Expected %s, but generated %s upon saving", response.AccountAddress, recoverResponse.AccountAddress)
	}

	// Clear terminal output
	_ = term.Clear()

	// Log & return
	fmt.Println("The node wallet was successfully initialized.")
	fmt.Printf("Node account: %s\n", response.AccountAddress.Hex())
	return nil

}
