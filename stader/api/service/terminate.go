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
package service

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/urfave/cli"

	"github.com/stader-labs/stader-node/shared/types/api"
)

const dataFolder string = "/.stader/data"

// Deletes the contents of the data folder including the wallet file, password file, and all validator keys.
// Don't use this unless you have a very good reason to do it (such as switching from Prater to Mainnet).
func terminateDataFolder(c *cli.Context) (*api.TerminateDataFolderResponse, error) {

	// Response
	response := api.TerminateDataFolderResponse{}

	// Check if it exists
	_, err := os.Stat(dataFolder)
	if os.IsNotExist(err) {
		response.FolderExisted = false
		return &response, nil
	}
	response.FolderExisted = true

	// Traverse it
	files, err := ioutil.ReadDir(dataFolder)
	if err != nil {
		return nil, fmt.Errorf("error enumerating files: %w", err)
	}

	// Delete the children
	for _, file := range files {
		// Skip the validators directory - that get special treatment
		if file.Name() != "validators" && !file.IsDir() {
			fullPath := filepath.Join(dataFolder, file.Name())
			if file.IsDir() {
				err = os.RemoveAll(fullPath)
			} else {
				err = os.Remove(fullPath)
			}
			if err != nil {
				return nil, fmt.Errorf("error removing [%s]: %w", file.Name(), err)
			}
		}
	}

	// Traverse the validators dir
	validatorsDir := filepath.Join(dataFolder, "validators")
	files, err = ioutil.ReadDir(validatorsDir)
	if err != nil {
		return nil, fmt.Errorf("error enumerating validator files: %w", err)
	}

	// Delete the children
	for _, file := range files {
		fullPath := filepath.Join(validatorsDir, file.Name())
		if file.IsDir() {
			err = os.RemoveAll(fullPath)
		} else {
			err = os.Remove(fullPath)
		}
		if err != nil {
			return nil, fmt.Errorf("error removing [%s]: %w", file.Name(), err)
		}
	}

	// Return response
	return &response, nil

}
