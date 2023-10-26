/*
This work is licensed and released under GNU GPL v3 or any other later versions.
The full text of the license is below/ found at <http://www.gnu.org/licenses/>

(c) 2023 Rocket Pool Pty Ltd. Modified under GNU GPL v3. [1.4.1]

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
	"strings"

	"github.com/stader-labs/stader-node/shared/utils/log"

	"github.com/ethereum/go-ethereum/common"
	"github.com/urfave/cli"

	"github.com/stader-labs/stader-node/shared/services/stader"
	cliutils "github.com/stader-labs/stader-node/shared/utils/cli"
)

func recoverWallet(c *cli.Context) error {

	staderOwner, err := stader.NewClientFromCtx(c)
	if err != nil {
		return err
	}
	defer staderOwner.Close()

	// Get & check wallet status
	status, err := staderOwner.WalletStatus()
	if err != nil {
		return err
	}
	if status.WalletInitialized {
		fmt.Println("The node wallet is already initialized.")
		return nil
	}

	// Prompt a notice about test recovery
	fmt.Printf("%sNOTE:\nThis command will fully regenerate your node wallet's private key and (unless explicitly disabled) the validator keys for your validators.%s\n\n", log.ColorYellow, log.ColorReset)

	// Set password if not set
	if !status.PasswordSet {
		var password string
		if c.String("password") != "" {
			password = c.String("password")
		} else {
			password = promptPassword()
		}
		if _, err := staderOwner.SetPassword(password); err != nil {
			return err
		}
	}

	// Prompt for mnemonic
	var mnemonic string
	if c.String("mnemonic") != "" {
		mnemonic = c.String("mnemonic")
	} else {
		mnemonic = promptMnemonic()
	}
	mnemonic = strings.TrimSpace(mnemonic)

	// Handle validator key recovery skipping
	skipValidatorKeyRecovery := c.Bool("skip-validator-key-recovery")

	// Check for a search-by-address operation
	addressString := c.String("address")
	if addressString != "" {

		// Get the address to search for
		address := common.HexToAddress(addressString)
		fmt.Printf("Searching for the derivation path and index for wallet %s...\nNOTE: this may take several minutes depending on how large your wallet's index is.\n", address.Hex())

		// Log
		if skipValidatorKeyRecovery {
			fmt.Println("Ignoring validator keys, searching for wallet only...")
		} else {
			// Check and assign the EC status
			err = cliutils.CheckClientStatus(staderOwner)
			if err != nil {
				return err
			}
		}

		// Recover wallet
		response, err := staderOwner.SearchAndRecoverWallet(mnemonic, address, skipValidatorKeyRecovery)
		if err != nil {
			return err
		}

		// Log & return
		fmt.Println("The node wallet was successfully recovered.")
		fmt.Printf("Derivation path: %s\n", response.DerivationPath)
		fmt.Printf("Wallet index:    %d\n", response.Index)
		fmt.Printf("Node account:    %s\n", response.AccountAddress.Hex())

		if !skipValidatorKeyRecovery {
			if len(response.ValidatorKeys) > 0 {
				fmt.Println("Validator keys:")
				for _, key := range response.ValidatorKeys {
					fmt.Println(key.Hex())
				}
			} else {
				fmt.Println("No validator keys were found.")
			}
		}

	} else {

		// Get the derivation path
		derivationPath := c.String("derivation-path")
		if derivationPath != "" {
			fmt.Printf("Using a custom derivation path (%s).\n", derivationPath)
		}

		// Get the wallet index
		walletIndex := c.Uint("wallet-index")
		if walletIndex != 0 {
			fmt.Printf("Using a custom wallet index (%d).\n", walletIndex)
		}

		fmt.Println()

		// Log
		if skipValidatorKeyRecovery {
			fmt.Println("Recovering node wallet only (ignoring validator keys)...")
		} else {
			// Check and assign the EC status
			err = cliutils.CheckClientStatus(staderOwner)
			if err != nil {
				return err
			}
			fmt.Println("Recovering node wallet and validator keys...")
		}

		// Recover wallet
		response, err := staderOwner.RecoverWallet(mnemonic, skipValidatorKeyRecovery, derivationPath, walletIndex)
		if err != nil {
			return err
		}

		// Log & return
		fmt.Println("The node wallet was successfully recovered.")
		fmt.Printf("Node account: %s\n", response.AccountAddress.Hex())
		if !skipValidatorKeyRecovery {

			if !response.OperatorExists {
				fmt.Println("Operator not registered with Stader, no validator keys to recover")
				return nil
			}

			if len(response.ValidatorKeys) > 0 {
				fmt.Println("Validator keys:")
				for _, key := range response.ValidatorKeys {
					fmt.Println(key.Hex())
				}
			} else {
				fmt.Println("No validator keys were found.")
			}
		}
	}

	return nil
}
