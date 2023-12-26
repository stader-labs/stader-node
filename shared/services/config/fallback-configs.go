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

// Configuration for fallback Lighthouse
type FallbackNormalConfig struct {
	Title string `yaml:"-"`

	// The URL of the Execution Client HTTP endpoint
	EcHttpUrl config.Parameter `yaml:"ecHttpUrl,omitempty"`

	// The URL of the Beacon Node HTTP endpoint
	CcHttpUrl config.Parameter `yaml:"ccHttpUrl,omitempty"`
}

// Configuration for fallback Prysm
type FallbackPrysmConfig struct {
	Title string `yaml:"-"`

	// The URL of the Execution Client HTTP endpoint
	EcHttpUrl config.Parameter `yaml:"ecHttpUrl,omitempty"`

	// The URL of the Beacon Node HTTP endpoint
	CcHttpUrl config.Parameter `yaml:"ccHttpUrl,omitempty"`

	// The URL of the JSON-RPC endpoint for the Validator client
	JsonRpcUrl config.Parameter `yaml:"jsonRpcUrl,omitempty"`
}

// Generates a new FallbackNormalConfig configuration
func NewFallbackNormalConfig(cfg *StaderConfig) *FallbackNormalConfig {
	return &FallbackNormalConfig{
		Title: "Fallback Client Settings",

		EcHttpUrl: config.Parameter{
			ID:                   "ecHttpUrl",
			Name:                 "Execution Client URL",
			Description:          "The URL of the HTTP API endpoint for your fallback Execution client.\n\nNOTE: If you are running it on the same machine as the Stadernode, addresses like `localhost` and `127.0.0.1` will not work due to Docker limitations. Enter your machine's LAN IP address instead.",
			Type:                 config.ParameterType_String,
			Default:              map[config.Network]interface{}{config.Network_All: ""},
			AffectsContainers:    []config.ContainerID{config.ContainerID_Api, config.ContainerID_Node, config.ContainerID_Guardian},
			EnvironmentVariables: []string{"FALLBACK_EC_API_ENDPOINT"},
			CanBeBlank:           false,
			OverwriteOnUpgrade:   false,
		},

		CcHttpUrl: config.Parameter{
			ID:                   "ccHttpUrl",
			Name:                 "Beacon Node URL",
			Description:          "The URL of the HTTP Beacon API endpoint for your fallback Consensus client.\n\nNOTE: If you are running it on the same machine as the Stadernode, addresses like `localhost` and `127.0.0.1` will not work due to Docker limitations. Enter your machine's LAN IP address instead.",
			Type:                 config.ParameterType_String,
			Default:              map[config.Network]interface{}{config.Network_All: ""},
			AffectsContainers:    []config.ContainerID{config.ContainerID_Api, config.ContainerID_Node, config.ContainerID_Validator, config.ContainerID_Guardian},
			EnvironmentVariables: []string{"FALLBACK_CC_API_ENDPOINT"},
			CanBeBlank:           false,
			OverwriteOnUpgrade:   false,
		},
	}
}

// Generates a new FallbackPrysmConfig configuration
func NewFallbackPrysmConfig(cfg *StaderConfig) *FallbackPrysmConfig {
	return &FallbackPrysmConfig{
		Title: "Fallback Prysm Settings",

		EcHttpUrl: config.Parameter{
			ID:                   "ecHttpUrl",
			Name:                 "Execution Client URL",
			Description:          "The URL of the HTTP API endpoint for your fallback Execution client.\n\nNOTE: If you are running it on the same machine as the Stadernode, addresses like `localhost` and `127.0.0.1` will not work due to Docker limitations. Enter your machine's LAN IP address instead.",
			Type:                 config.ParameterType_String,
			Default:              map[config.Network]interface{}{config.Network_All: ""},
			AffectsContainers:    []config.ContainerID{config.ContainerID_Api, config.ContainerID_Node, config.ContainerID_Guardian},
			EnvironmentVariables: []string{"FALLBACK_EC_API_ENDPOINT"},
			CanBeBlank:           false,
			OverwriteOnUpgrade:   false,
		},

		CcHttpUrl: config.Parameter{
			ID:                   "ccHttpUrl",
			Name:                 "Beacon Node HTTP URL",
			Description:          "The URL of the HTTP Beacon API endpoint for your fallback Prysm client.\n\nNOTE: If you are running it on the same machine as the Stadernode, addresses like `localhost` and `127.0.0.1` will not work due to Docker limitations. Enter your machine's LAN IP address instead.",
			Type:                 config.ParameterType_String,
			Default:              map[config.Network]interface{}{config.Network_All: ""},
			AffectsContainers:    []config.ContainerID{config.ContainerID_Api, config.ContainerID_Node, config.ContainerID_Validator, config.ContainerID_Guardian},
			EnvironmentVariables: []string{"FALLBACK_CC_API_ENDPOINT"},
			CanBeBlank:           false,
			OverwriteOnUpgrade:   false,
		},

		JsonRpcUrl: config.Parameter{
			ID:                   "jsonRpcUrl",
			Name:                 "Beacon Node JSON-RPC URL",
			Description:          "The URL of the JSON-RPC API endpoint for your fallback client. Prysm's validator client will need this in order to connect to it.\n\nNOTE: If you are running it on the same machine as the Stadernode, addresses like `localhost` and `127.0.0.1` will not work due to Docker limitations. Enter your machine's LAN IP address instead.",
			Type:                 config.ParameterType_String,
			Default:              map[config.Network]interface{}{config.Network_All: ""},
			AffectsContainers:    []config.ContainerID{config.ContainerID_Eth1},
			EnvironmentVariables: []string{"FALLBACK_CC_RPC_ENDPOINT"},
			CanBeBlank:           false,
			OverwriteOnUpgrade:   false,
		},
	}
}

// Get the config.Parameters for this config
func (cfg *FallbackNormalConfig) GetParameters() []*config.Parameter {
	return []*config.Parameter{
		&cfg.EcHttpUrl,
		&cfg.CcHttpUrl,
	}
}

// Get the config.Parameters for this config
func (cfg *FallbackPrysmConfig) GetParameters() []*config.Parameter {
	return []*config.Parameter{
		&cfg.EcHttpUrl,
		&cfg.CcHttpUrl,
		&cfg.JsonRpcUrl,
	}
}

// The the title for the config
func (config *FallbackNormalConfig) GetConfigTitle() string {
	return config.Title
}

// The the title for the config
func (config *FallbackPrysmConfig) GetConfigTitle() string {
	return config.Title
}
