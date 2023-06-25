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
	"fmt"

	"github.com/stader-labs/stader-node/shared/services/stader"
	"github.com/urfave/cli"

	cliutils "github.com/stader-labs/stader-node/shared/utils/cli"
)

func getStatus(c *cli.Context) error {

	staderClient, err := stader.NewClientFromCtx(c)
	if err != nil {
		return err
	}
	defer staderClient.Close()

	// Print what network we're on
	err = cliutils.PrintNetwork(staderClient)
	if err != nil {
		return err
	}

	// Get wallet status
	status, err := staderClient.WalletStatus()
	if err != nil {
		return err
	}

	// Print status & return
	if status.WalletInitialized {
		fmt.Println("The node wallet is initialized.")
		fmt.Printf("Node account: %s\n", status.AccountAddress.Hex())
		fmt.Printf("Current Nonce: %d\n", status.CurrentNonce)
		fmt.Printf("Pending Nonce: %d\n", status.PendingNonce)
	} else {
		fmt.Println("The node wallet has not been initialized.")
	}
	return nil

}
