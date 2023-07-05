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
package config

import (
	"github.com/stader-labs/stader-node/shared/types/config"
)

// Constants
const exporterTag string = "prom/node-exporter:v1.6.0"

// Defaults
const defaultExporterRootFs bool = false

// Configuration for Exporter
type ExporterConfig struct {
	Title string `yaml:"-"`

	// Toggle for enabling access to the root filesystem (for multiple disk usage metrics)
	RootFs config.Parameter `yaml:"rootFs,omitempty"`

	// The Docker Hub tag for Prometheus
	ContainerTag config.Parameter `yaml:"containerTag,omitempty"`

	// Custom command line flags
	AdditionalFlags config.Parameter `yaml:"additionalFlags,omitempty"`
}

type ExternalExporterConfig struct {
	Title string `yaml:"-"`

	// The URL of the HTTP endpoint
	HttpUrl config.Parameter `yaml:"httpUrl,omitempty"`
}

// Generates a new Exporter config
func NewExporterConfig(cfg *StaderConfig) *ExporterConfig {
	return &ExporterConfig{
		Title: "Node Exporter Settings",

		RootFs: config.Parameter{
			ID:                   "enableRootFs",
			Name:                 "Allow Root Filesystem Access",
			Description:          "Give Prometheus's Node Exporter permission to view your root filesystem instead of being limited to its own Docker container.\nThis is needed if you want the Grafana dashboard to report the used disk space of a second SSD.",
			Type:                 config.ParameterType_Bool,
			Default:              map[config.Network]interface{}{config.Network_All: defaultExporterRootFs},
			AffectsContainers:    []config.ContainerID{config.ContainerID_Exporter},
			EnvironmentVariables: []string{"EXPORTER_ROOT_FS"},
			CanBeBlank:           false,
			OverwriteOnUpgrade:   false,
		},

		ContainerTag: config.Parameter{
			ID:                   "containerTag",
			Name:                 "Exporter Container Tag",
			Description:          "The tag name of the Prometheus Node Exporter container you want to use on Docker Hub.",
			Type:                 config.ParameterType_String,
			Default:              map[config.Network]interface{}{config.Network_All: exporterTag},
			AffectsContainers:    []config.ContainerID{config.ContainerID_Exporter},
			EnvironmentVariables: []string{"EXPORTER_CONTAINER_TAG"},
			CanBeBlank:           false,
			OverwriteOnUpgrade:   true,
		},

		AdditionalFlags: config.Parameter{
			ID:                   "additionalFlags",
			Name:                 "Additional Exporter Flags",
			Description:          "Additional custom command line flags you want to pass to the Node Exporter, to take advantage of other settings that the Stadernode's configuration doesn't cover.",
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
func (cfg *ExporterConfig) GetParameters() []*config.Parameter {
	return []*config.Parameter{
		&cfg.RootFs,
		&cfg.ContainerTag,
		&cfg.AdditionalFlags,
	}
}

// The the title for the config
func (cfg *ExporterConfig) GetConfigTitle() string {
	return cfg.Title
}
