/*
This work is licensed and released under GNU GPL v3 or any other later versions.
The full text of the license is below/ found at <http://www.gnu.org/licenses/>

(c) 2023 Rocket Pool Pty Ltd. Modified under GNU GPL v3. [1.4.3]

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
package stader

import (
	"fmt"
	"io/ioutil"

	"github.com/alessio/shellescape"
	"github.com/mitchellh/go-homedir"
	"github.com/stader-labs/stader-node/shared/services/config"
)

// Config
const (
	LegacyGlobalConfigFile    = "config.yml"
	LegacyUserConfigFile      = "settings.yml"
	LegacyComposeFile         = "docker-compose.yml"
	LegacyMetricsComposeFile  = "docker-compose-metrics.yml"
	LegacyFallbackComposeFile = "docker-compose-fallback.yml"
)

// Load the global config
func (c *Client) LoadGlobalConfig_Legacy(globalConfigPath string) (config.LegacyStaderConfig, error) {
	return c.loadConfig_Legacy(globalConfigPath)
}

// Load/save the user config
func (c *Client) LoadUserConfig_Legacy(userConfigPath string) (config.LegacyStaderConfig, error) {
	return c.loadConfig_Legacy(userConfigPath)
}

// Load the merged global & user config
func (c *Client) LoadMergedConfig_Legacy(globalConfigPath string, userConfigPath string) (config.LegacyStaderConfig, error) {
	globalConfig, err := c.LoadGlobalConfig_Legacy(globalConfigPath)
	if err != nil {
		return config.LegacyStaderConfig{}, err
	}
	userConfig, err := c.LoadUserConfig_Legacy(userConfigPath)
	if err != nil {
		return config.LegacyStaderConfig{}, err
	}
	return config.Merge(&globalConfig, &userConfig)
}

// Load a config file
func (c *Client) loadConfig_Legacy(path string) (config.LegacyStaderConfig, error) {
	expandedPath, err := homedir.Expand(path)
	if err != nil {
		return config.LegacyStaderConfig{}, err
	}
	configBytes, err := ioutil.ReadFile(expandedPath)
	if err != nil {
		return config.LegacyStaderConfig{}, fmt.Errorf("Could not read Stader config at %s: %w", shellescape.Quote(path), err)
	}
	return config.Parse(configBytes)
}
