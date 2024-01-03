/*
This work is licensed and released under GNU GPL v3 or any other later versions.
The full text of the license is below/ found at <http://www.gnu.org/licenses/>

(c) 2023 Rocket Pool Pty Ltd. Modified under GNU GPL v3. [1.4.4]

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
package config

import (
	"github.com/stader-labs/stader-node/shared/types/config"
)

// Defaults
const (
	defaultBitflyNodeMetricsSecret      string = ""
	defaultBitflyNodeMetricsEndpoint    string = "https://beaconcha.in/api/v1/client/metrics"
	defaultBitflyNodeMetricsMachineName string = "Stadernode"
)

// Configuration for Bitfly Node Metrics
type BitflyNodeMetricsConfig struct {
	Title string `yaml:"-"`

	Secret config.Parameter `yaml:"secret,omitempty"`

	Endpoint config.Parameter `yaml:"endpoint,omitempty"`

	MachineName config.Parameter `yaml:"machineName,omitempty"`
}

// Generates a new Bitfly Node Metrics config
func NewBitflyNodeMetricsConfig(cfg *StaderConfig) *BitflyNodeMetricsConfig {
	return &BitflyNodeMetricsConfig{
		Title: "Bitfly Node Metrics Settings",

		Secret: config.Parameter{
			ID:                   "bitflySecret",
			Name:                 "Beaconcha.in API Key",
			Description:          "The API key used to authenticate your Beaconcha.in node metrics integration. Can be found in your Beaconcha.in account settings.\n\nPlease visit https://beaconcha.in/user/settings#api to access your account information.",
			Type:                 config.ParameterType_String,
			Default:              map[config.Network]interface{}{config.Network_All: defaultBitflyNodeMetricsSecret},
			AffectsContainers:    []config.ContainerID{config.ContainerID_Validator, config.ContainerID_Eth2},
			EnvironmentVariables: []string{"BITFLY_NODE_METRICS_SECRET"},
			// ensures the string is 28 characters of Base64
			Regex:              "^[A-Za-z0-9+/]{28}$",
			CanBeBlank:         false,
			OverwriteOnUpgrade: false,
		},

		Endpoint: config.Parameter{
			ID:                   "bitflyEndpoint",
			Name:                 "Node Metrics Endpoint",
			Description:          "The endpoint to send your Beaconcha.in Node Metrics data to. Should be left as the default.",
			Type:                 config.ParameterType_String,
			Default:              map[config.Network]interface{}{config.Network_All: defaultBitflyNodeMetricsEndpoint},
			AffectsContainers:    []config.ContainerID{config.ContainerID_Validator, config.ContainerID_Eth2},
			EnvironmentVariables: []string{"BITFLY_NODE_METRICS_ENDPOINT"},
			CanBeBlank:           false,
			OverwriteOnUpgrade:   false,
		},

		MachineName: config.Parameter{
			ID:                   "bitflyMachineName",
			Name:                 "Node Metrics Machine Name",
			Description:          "The name of the machine you are running on. This is used to identify your machine in the mobile app.\nChange this if you are running multiple Stader nodes with the same Secret.",
			Type:                 config.ParameterType_String,
			Default:              map[config.Network]interface{}{config.Network_All: defaultBitflyNodeMetricsMachineName},
			AffectsContainers:    []config.ContainerID{config.ContainerID_Validator, config.ContainerID_Eth2},
			EnvironmentVariables: []string{"BITFLY_NODE_METRICS_MACHINE_NAME"},
			CanBeBlank:           false,
			OverwriteOnUpgrade:   false,
		},
	}
}

// Get the parameters for this config
func (cfg *BitflyNodeMetricsConfig) GetParameters() []*config.Parameter {
	return []*config.Parameter{
		&cfg.Secret,
		&cfg.Endpoint,
		&cfg.MachineName,
	}
}

// The the title for the config
func (cfg *BitflyNodeMetricsConfig) GetConfigTitle() string {
	return cfg.Title
}
