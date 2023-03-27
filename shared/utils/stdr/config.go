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
package stdr

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/alessio/shellescape"
	"github.com/stader-labs/stader-node/shared/services/config"
	"gopkg.in/yaml.v2"
)

const (
	upgradeFlagFile string = ".firstrun"
)

// Loads a config without updating it if it exists
func LoadConfigFromFile(path string) (*config.StaderConfig, error) {
	_, err := os.Stat(path)
	if os.IsNotExist(err) {
		return nil, nil
	}

	cfg, err := config.LoadFromFile(path)
	if err != nil {
		return nil, err
	}

	return cfg, nil
}

// Saves a config and removes the upgrade flag file
func SaveConfig(cfg *config.StaderConfig, path string) error {

	settings := cfg.Serialize()
	configBytes, err := yaml.Marshal(settings)
	if err != nil {
		return fmt.Errorf("could not serialize settings file: %w", err)
	}

	if err := ioutil.WriteFile(path, configBytes, 0664); err != nil {
		return fmt.Errorf("could not write Stader config to %s: %w", shellescape.Quote(path), err)
	}

	return nil

}

// Checks if this is the first run of the configurator after an install
func IsFirstRun(configDir string) bool {
	upgradeFilePath := filepath.Join(configDir, upgradeFlagFile)

	// Load the config normally if the upgrade flag file isn't there
	_, err := os.Stat(upgradeFilePath)
	if os.IsNotExist(err) {
		return false
	}

	return true
}

// Remove the upgrade flag file
func RemoveUpgradeFlagFile(configDir string) error {

	// Check for the upgrade flag file
	upgradeFilePath := filepath.Join(configDir, upgradeFlagFile)
	_, err := os.Stat(upgradeFilePath)
	if os.IsNotExist(err) {
		return nil
	}

	// Delete the upgrade flag file
	err = os.Remove(upgradeFilePath)
	if err != nil {
		return fmt.Errorf("error removing upgrade flag file: %w", err)
	}

	return nil

}
