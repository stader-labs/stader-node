/*
This work is licensed and released under GNU GPL v3 or any other later versions.
The full text of the license is below/ found at <http://www.gnu.org/licenses/>

(c) 2023 Rocket Pool Pty Ltd. Modified under GNU GPL v3. [0.4.0-beta]

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

// Constants
const prometheusTag string = "prom/prometheus:v2.41.0"

// Defaults
const defaultPrometheusPort uint16 = 9091
const defaultPrometheusOpenPort bool = false

// Configuration for Prometheus
type PrometheusConfig struct {
	Title string `yaml:"-"`

	// The port to serve metrics on
	Port config.Parameter `yaml:"port,omitempty"`

	// Toggle for forwarding the API port outside of Docker
	OpenPort config.Parameter `yaml:"openPort,omitempty"`

	// The Docker Hub tag for Prometheus
	ContainerTag config.Parameter `yaml:"containerTag,omitempty"`

	// Custom command line flags
	AdditionalFlags config.Parameter `yaml:"additionalFlags,omitempty"`
}

// Generates a new Prometheus config
func NewPrometheusConfig(cfg *StaderConfig) *PrometheusConfig {
	return &PrometheusConfig{
		Title: "Prometheus Settings",

		Port: config.Parameter{
			ID:                   "port",
			Name:                 "Prometheus Port",
			Description:          "The port Prometheus should make its statistics available on.",
			Type:                 config.ParameterType_Uint16,
			Default:              map[config.Network]interface{}{config.Network_All: defaultPrometheusPort},
			AffectsContainers:    []config.ContainerID{config.ContainerID_Prometheus},
			EnvironmentVariables: []string{"PROMETHEUS_PORT"},
			CanBeBlank:           true,
			OverwriteOnUpgrade:   false,
		},

		OpenPort: config.Parameter{
			ID:                   "openPort",
			Name:                 "Expose Prometheus Port",
			Description:          "Enable this to expose Prometheus's port to your local network, so other machines can access it too.",
			Type:                 config.ParameterType_Bool,
			Default:              map[config.Network]interface{}{config.Network_All: defaultPrometheusOpenPort},
			AffectsContainers:    []config.ContainerID{config.ContainerID_Prometheus},
			EnvironmentVariables: []string{"PROMETHEUS_OPEN_PORT"},
			CanBeBlank:           false,
			OverwriteOnUpgrade:   false,
		},

		ContainerTag: config.Parameter{
			ID:                   "containerTag",
			Name:                 "Prometheus Container Tag",
			Description:          "The tag name of the Prometheus container you want to use on Docker Hub.",
			Type:                 config.ParameterType_String,
			Default:              map[config.Network]interface{}{config.Network_All: prometheusTag},
			AffectsContainers:    []config.ContainerID{config.ContainerID_Prometheus},
			EnvironmentVariables: []string{"PROMETHEUS_CONTAINER_TAG"},
			CanBeBlank:           false,
			OverwriteOnUpgrade:   true,
		},

		AdditionalFlags: config.Parameter{
			ID:                   "additionalFlags",
			Name:                 "Additional Prometheus Flags",
			Description:          "Additional custom command line flags you want to pass to Prometheus, to take advantage of other settings that the Stadernode's configuration doesn't cover.",
			Type:                 config.ParameterType_String,
			Default:              map[config.Network]interface{}{config.Network_All: ""},
			AffectsContainers:    []config.ContainerID{config.ContainerID_Grafana},
			EnvironmentVariables: []string{},
			CanBeBlank:           true,
			OverwriteOnUpgrade:   false,
		},
	}
}

// Get the parameters for this config
func (cfg *PrometheusConfig) GetParameters() []*config.Parameter {
	return []*config.Parameter{
		&cfg.Port,
		&cfg.OpenPort,
		&cfg.ContainerTag,
		&cfg.AdditionalFlags,
	}
}

// The the title for the config
func (cfg *PrometheusConfig) GetConfigTitle() string {
	return cfg.Title
}
