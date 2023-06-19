/*
This work is licensed and released under GNU GPL v3 or any other later versions.
The full text of the license is below/ found at <http://www.gnu.org/licenses/>

(c) 2023 Rocket Pool Pty Ltd. Modified under GNU GPL v3. [1.0.0]

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
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strconv"
	"strings"

	"github.com/stader-labs/stader-node/shared/utils/log"

	"github.com/mitchellh/go-homedir"
	"github.com/stader-labs/stader-node/shared/services/config"
	"github.com/stader-labs/stader-node/shared/services/passwords"
	"github.com/stader-labs/stader-node/shared/services/stader"
	"github.com/stader-labs/stader-node/shared/types/api"
	cliutils "github.com/stader-labs/stader-node/shared/utils/cli"
	hexutils "github.com/stader-labs/stader-node/shared/utils/hex"
	"github.com/stader-labs/stader-node/stader-cli/wallet/bip39"
	"github.com/stader-labs/stader-node/stader-lib/types"
	"gopkg.in/yaml.v2"
)

const bold string = "\033[1m"
const unbold string = "\033[0m"

// Prompt for a wallet password
func promptPassword() string {
	for {
		password := cliutils.PromptPassword(
			"Please enter a password to secure your wallet with:",
			fmt.Sprintf("^.{%d,}$", passwords.MinPasswordLength),
			fmt.Sprintf("Your password must be at least %d characters long. Please try again:", passwords.MinPasswordLength),
		)
		confirmation := cliutils.PromptPassword("Please confirm your password:", "^.*$", "")
		if password == confirmation {
			return password
		}
		fmt.Println("Password confirmation does not match.")
		fmt.Println("")
	}
}

// Prompt for a recovery mnemonic phrase
func promptMnemonic() string {
	for {
		lengthInput := cliutils.Prompt(
			"Please enter the "+bold+"number"+unbold+" of words in your mnemonic phrase (24 by default):",
			"^[1-9][0-9]*$",
			"Please enter a valid number.")

		length, err := strconv.Atoi(lengthInput)
		if err != nil {
			fmt.Println("Please enter a valid number.")
			continue
		}

		mv := bip39.Create(length)
		if mv == nil {
			fmt.Println("Please enter a valid mnemonic length.")
			continue
		}

		i := 0
		for mv.Filled() == false {
			prompt := fmt.Sprintf("Enter %sWord Number %d%s of your mnemonic:", bold, i+1, unbold)
			word := cliutils.PromptPassword(prompt, "^[a-zA-Z]+$", "Please enter a single word only.")

			if err := mv.AddWord(strings.ToLower(word)); err != nil {
				fmt.Println("Inputted word not valid, please retry.")
				continue
			}

			i++
		}

		mnemonic, err := mv.Finalize()
		if err != nil {
			fmt.Printf("Error validating mnemonic: %s\n", err)
			fmt.Println("Please try again.")
			fmt.Println("")
			continue
		}

		return mnemonic
	}
}

// Confirm a recovery mnemonic phrase
func confirmMnemonic(mnemonic string) {
	for {
		fmt.Println("Please enter your mnemonic phrase to confirm.")
		confirmation := promptMnemonic()
		if mnemonic == confirmation {
			return
		}
		fmt.Println("The mnemonic phrase you entered does not match your recovery phrase. Please try again.")
		fmt.Println("")
	}
}

// Check for custom keys, prompt for their passwords, and store them in the custom keys file
func promptForCustomKeyPasswords(sd *stader.Client, cfg *config.StaderConfig, testOnly bool) (string, error) {

	// Check for the custom key directory
	datapath, err := homedir.Expand(cfg.StaderNode.DataPath.Value.(string))
	if err != nil {
		return "", fmt.Errorf("error expanding data directory: %w", err)
	}
	customKeyDir := filepath.Join(datapath, "custom-keys")
	info, err := os.Stat(customKeyDir)
	if os.IsNotExist(err) || !info.IsDir() {
		return "", nil
	}

	// Get the custom keystore files
	files, err := ioutil.ReadDir(customKeyDir)
	if err != nil {
		return "", fmt.Errorf("error enumerating custom keystores: %w", err)
	}
	if len(files) == 0 {
		return "", nil
	}

	// Prompt the user with a warning message
	if !testOnly {
		fmt.Printf("%sWARNING:\nThe Stadernode has detected that you have custom (externally-derived) validator keys for your validators.\nIf these keys were actively used for validation by a service such as Allnodes, you MUST CONFIRM WITH THAT SERVICE that they have stopped validating and disabled those keys, and will NEVER validate with them again.\nOtherwise, you may both run the same keys at the same time which WILL RESULT IN YOUR VALIDATORS BEING SLASHED.%s\n\n", log.ColorRed, log.ColorReset)

		if !cliutils.Confirm("Please confirm that you have coordinated with the service that was running your validators previously to ensure they have STOPPED validation for your validators, will NEVER start them again, and you have manually confirmed on a Blockchain explorer such as https://beaconcha.in that your validators are no longer attesting.") {
			fmt.Println("Cancelled.")
			os.Exit(0)
		}
	}

	// Get the pubkeys for the custom keystores
	customPubkeys := []types.ValidatorPubkey{}
	for _, file := range files {
		// Read the file
		bytes, err := ioutil.ReadFile(filepath.Join(customKeyDir, file.Name()))
		if err != nil {
			return "", fmt.Errorf("error reading custom keystore %s: %w", file.Name(), err)
		}

		// Deserialize it
		keystore := api.ValidatorKeystore{}
		err = json.Unmarshal(bytes, &keystore)
		if err != nil {
			return "", fmt.Errorf("error deserializing custom keystore %s: %w", file.Name(), err)
		}

		customPubkeys = append(customPubkeys, keystore.Pubkey)
	}

	// Notify the user
	fmt.Println("It looks like you have some custom keystores for your stader validators.\nYou will be prompted for the passwords each one was encrypted with, so they can be loaded into the Validator Client that stader manages for you.")

	// Get the passwords for each one
	pubkeyPasswords := map[string]string{}
	for _, pubkey := range customPubkeys {
		password := cliutils.PromptPassword(
			fmt.Sprintf("Please enter the password that the keystore for %s was encrypted with:", pubkey.Hex()), "^.*$", "",
		)

		formattedPubkey := strings.ToUpper(hexutils.RemovePrefix(pubkey.Hex()))
		pubkeyPasswords[formattedPubkey] = password

		fmt.Println()
	}

	// Store them in the file
	fileBytes, err := yaml.Marshal(pubkeyPasswords)
	if err != nil {
		return "", fmt.Errorf("error serializing keystore passwords file: %w", err)
	}
	passwordFile := filepath.Join(datapath, "custom-key-passwords")
	err = ioutil.WriteFile(passwordFile, fileBytes, 0600)
	if err != nil {
		return "", fmt.Errorf("error writing keystore passwords file: %w", err)
	}

	return passwordFile, nil

}

// Deletes the custom key password file
func deleteCustomKeyPasswordFile(passwordFile string) error {
	_, err := os.Stat(passwordFile)
	if os.IsNotExist(err) {
		return nil
	}

	err = os.Remove(passwordFile)
	return err
}
