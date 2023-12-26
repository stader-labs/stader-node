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
package config

import (
	"github.com/stader-labs/stader-node/shared/types/config"
)

// Constants
const grafanaTag string = "grafana/grafana:9.4.15"

// Defaults
const defaultGrafanaPort uint16 = 3100

// Configuration for Grafana
type GrafanaConfig struct {
	Title string `yaml:"-"`

	// The HTTP port to serve on
	Port config.Parameter `yaml:"port,omitempty"`

	// The Docker Hub tag for Grafana
	ContainerTag config.Parameter `yaml:"containerTag,omitempty"`
}
type ExternalGrafanaConfig struct {
	Title string `yaml:"-"`

	// The URL of the HTTP endpoint
	HttpUrl config.Parameter `yaml:"httpUrl,omitempty"`
}

// Generates a new Grafana config
func NewGrafanaConfig(cfg *StaderConfig) *GrafanaConfig {
	return &GrafanaConfig{
		Title: "Grafana Settings",

		Port: config.Parameter{
			ID:                   "port",
			Name:                 "Grafana Port",
			Description:          "The port Grafana should run its HTTP server on - this is the port you will connect to in your browser.",
			Type:                 config.ParameterType_Uint16,
			Default:              map[config.Network]interface{}{config.Network_All: defaultGrafanaPort},
			AffectsContainers:    []config.ContainerID{config.ContainerID_Grafana},
			EnvironmentVariables: []string{"GRAFANA_PORT"},
			CanBeBlank:           false,
			OverwriteOnUpgrade:   false,
		},

		ContainerTag: config.Parameter{
			ID:                   "containerTag",
			Name:                 "Grafana Container Tag",
			Description:          "The tag name of the Grafana container you want to use on Docker Hub.",
			Type:                 config.ParameterType_String,
			Default:              map[config.Network]interface{}{config.Network_All: grafanaTag},
			AffectsContainers:    []config.ContainerID{config.ContainerID_Grafana},
			EnvironmentVariables: []string{"GRAFANA_CONTAINER_TAG"},
			CanBeBlank:           false,
			OverwriteOnUpgrade:   true,
		},
	}
}

// Get the parameters for this config
func (cfg *GrafanaConfig) GetParameters() []*config.Parameter {
	return []*config.Parameter{
		&cfg.Port,
		&cfg.ContainerTag,
	}
}

// The the title for the config
func (cfg *GrafanaConfig) GetConfigTitle() string {
	return cfg.Title
}
