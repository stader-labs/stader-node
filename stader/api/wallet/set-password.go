/*
This work is licensed and released under GNU GPL v3 or any other later versions.
The full text of the license is below/ found at <http://www.gnu.org/licenses/>

(c) 2023 Rocket Pool Pty Ltd. Modified under GNU GPL v3. [1.2.1]

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
	"errors"

	"github.com/urfave/cli"

	"github.com/stader-labs/stader-node/shared/services"
	"github.com/stader-labs/stader-node/shared/types/api"
)

func setPassword(c *cli.Context, password string) (*api.SetPasswordResponse, error) {

	// Get services
	pm, err := services.GetPasswordManager(c)
	if err != nil {
		return nil, err
	}

	// Response
	response := api.SetPasswordResponse{}

	// Check if password is already set
	if pm.IsPasswordSet() {
		return nil, errors.New("The node password is already set")
	}

	// Set password
	if err := pm.SetPassword(password); err != nil {
		return nil, err
	}

	// Return response
	return &response, nil

}
