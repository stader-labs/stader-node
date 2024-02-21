/*
This work is licensed and released under GNU GPL v3 or any other later versions.
The full text of the license is below/ found at <http://www.gnu.org/licenses/>

(c) 2023 Rocket Pool Pty Ltd. Modified under GNU GPL v3. [1.4.9]

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
package node

import (
	"fmt"

	"encoding/json"

	"github.com/ethereum/go-ethereum/common"
	"github.com/urfave/cli"

	"github.com/stader-labs/stader-node/shared/services/stader"
	cliutils "github.com/stader-labs/stader-node/shared/utils/cli"
)

const signatureVersion = 1

type PersonalSignature struct {
	Address   common.Address `json:"address"`
	Message   string         `json:"msg"`
	Signature string         `json:"sig"`
	Version   string         `json:"version"` // beaconcha.in expects a string
}

func signMessage(c *cli.Context) error {

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

	if !status.WalletInitialized {
		fmt.Println("The node wallet is not initialized.")
		return nil
	}

	message := c.String("message")
	for message == "" {
		message = cliutils.Prompt("Please enter the message you want to sign: (EIP-191 personal_sign)", "^.+$", "Please enter the message you want to sign: (EIP-191 personal_sign)")
	}

	response, err := staderClient.SignMessage(message)
	if err != nil {
		return err
	}

	// Print the signature
	formattedSignature := PersonalSignature{
		Address:   status.AccountAddress,
		Message:   message,
		Signature: response.SignedData,
		Version:   fmt.Sprint(signatureVersion),
	}
	bytes, err := json.MarshalIndent(formattedSignature, "", "    ")
	if err != nil {
		return err
	}

	fmt.Printf("Signed Error:\n\n%s\n", string(bytes))

	return nil

}
