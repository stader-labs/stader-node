/*
This work is licensed and released under GNU GPL v3 or any other later versions.
The full text of the license is below/ found at <http://www.gnu.org/licenses/>

(c) 2023 Rocket Pool Pty Ltd. Modified under GNU GPL v3. [1.5.0]

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

const (
	lighthouseTagPortableTest string = "sigp/lighthouse:v7.0.0-beta.0"
	lighthouseTagPortableProd string = "sigp/lighthouse:v6.0.1"
	defaultLhMaxPeers         uint16 = 100
)

// Configuration for Lighthouse
type LighthouseConfig struct {
	Title string `yaml:"-"`

	// The max number of P2P peers to connect to
	MaxPeers config.Parameter `yaml:"maxPeers,omitempty"`

	// Common parameters that Lighthouse doesn't support and should be hidden
	UnsupportedCommonParams []string `yaml:"-"`

	// The Docker Hub tag for Lighthouse
	ContainerTag config.Parameter `yaml:"containerTag,omitempty"`

	// Custom command line flags for the BN
	AdditionalBnFlags config.Parameter `yaml:"additionalBnFlags,omitempty"`

	// Custom command line flags for the VC
	AdditionalVcFlags config.Parameter `yaml:"additionalVcFlags,omitempty"`
}

// Generates a new Lighthouse configuration
func NewLighthouseConfig(cfg *StaderConfig) *LighthouseConfig {
	return &LighthouseConfig{
		Title: "Lighthouse Settings",

		MaxPeers: config.Parameter{
			ID:                   "maxPeers",
			Name:                 "Max Peers",
			Description:          "The maximum number of peers your client should try to maintain. You can try lowering this if you have a low-resource system or a constrained network.",
			Type:                 config.ParameterType_Uint16,
			Default:              map[config.Network]interface{}{config.Network_All: defaultLhMaxPeers},
			AffectsContainers:    []config.ContainerID{config.ContainerID_Eth2},
			EnvironmentVariables: []string{"BN_MAX_PEERS"},
			CanBeBlank:           false,
			OverwriteOnUpgrade:   false,
		},

		ContainerTag: config.Parameter{
			ID:          "containerTag",
			Name:        "Container Tag",
			Description: "The tag name of the Lighthouse container you want to use from Docker Hub.",
			Type:        config.ParameterType_String,
			Default: map[config.Network]interface{}{
				config.Network_Mainnet: lighthouseTagPortableProd,
				config.Network_Holesky: lighthouseTagPortableTest,
			},
			AffectsContainers:    []config.ContainerID{config.ContainerID_Eth2, config.ContainerID_Validator},
			EnvironmentVariables: []string{"BN_CONTAINER_TAG", "VC_CONTAINER_TAG"},
			CanBeBlank:           false,
			OverwriteOnUpgrade:   true,
		},

		AdditionalBnFlags: config.Parameter{
			ID:                   "additionalBnFlags",
			Name:                 "Additional Beacon Client Flags",
			Description:          "Additional custom command line flags you want to pass Lighthouse's Beacon Client, to take advantage of other settings that the Stadernode's configuration doesn't cover.",
			Type:                 config.ParameterType_String,
			Default:              map[config.Network]interface{}{config.Network_All: ""},
			AffectsContainers:    []config.ContainerID{config.ContainerID_Eth2},
			EnvironmentVariables: []string{"BN_ADDITIONAL_FLAGS"},
			CanBeBlank:           true,
			OverwriteOnUpgrade:   false,
		},

		AdditionalVcFlags: config.Parameter{
			ID:                   "additionalVcFlags",
			Name:                 "Additional Validator Client Flags",
			Description:          "Additional custom command line flags you want to pass Lighthouse's Validator Client, to take advantage of other settings that the Stadernode's configuration doesn't cover.",
			Type:                 config.ParameterType_String,
			Default:              map[config.Network]interface{}{config.Network_All: ""},
			AffectsContainers:    []config.ContainerID{config.ContainerID_Validator},
			EnvironmentVariables: []string{"VC_ADDITIONAL_FLAGS"},
			CanBeBlank:           true,
			OverwriteOnUpgrade:   false,
		},
	}
}

// Get the parameters for this config
func (cfg *LighthouseConfig) GetParameters() []*config.Parameter {
	return []*config.Parameter{
		&cfg.MaxPeers,
		&cfg.ContainerTag,
		&cfg.AdditionalBnFlags,
		&cfg.AdditionalVcFlags,
	}
}

// Get the common params that this client doesn't support
func (cfg *LighthouseConfig) GetUnsupportedCommonParams() []string {
	return cfg.UnsupportedCommonParams
}

// Get the Docker container name of the validator client
func (cfg *LighthouseConfig) GetValidatorImage() string {
	return cfg.ContainerTag.Value.(string)
}

// Get the name of the client
func (cfg *LighthouseConfig) GetName() string {
	return "Lighthouse"
}

// The the title for the config
func (cfg *LighthouseConfig) GetConfigTitle() string {
	return cfg.Title
}
