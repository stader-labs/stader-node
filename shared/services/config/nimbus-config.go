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
package config

import (
	"runtime"

	"github.com/stader-labs/stader-node/shared/types/config"
)

const (
	// Prater
	nimbusBnTagTest string = "statusim/nimbus-eth2:multiarch-v23.3.2"
	nimbusVcTagTest string = "statusim/nimbus-validator-client:multiarch-v23.3.2"

	// Mainnet
	nimbusBnTagProd string = "statusim/nimbus-eth2:multiarch-v23.3.2"
	nimbusVcTagProd string = "statusim/nimbus-validator-client:multiarch-v23.3.2"

	defaultNimbusMaxPeersArm uint16 = 100
	defaultNimbusMaxPeersAmd uint16 = 160
)

// Configuration for Nimbus
type NimbusConfig struct {
	Title string `yaml:"-"`

	// The max number of P2P peers to connect to
	MaxPeers config.Parameter `yaml:"maxPeers,omitempty"`

	// Common parameters that Nimbus doesn't support and should be hidden
	UnsupportedCommonParams []string `yaml:"-"`

	// The Docker Hub tag for the BN
	BnContainerTag config.Parameter `yaml:"bnContainerTag,omitempty"`

	// The Docker Hub tag for the VC
	VcContainerTag config.Parameter `yaml:"vcContainerTag,omitempty"`

	// The pruning mode to use in the BN
	PruningMode config.Parameter `yaml:"pruningMode,omitempty"`

	// Custom command line flags for the BN
	AdditionalBnFlags config.Parameter `yaml:"additionalBnFlags,omitempty"`

	// Custom command line flags for the VC
	AdditionalVcFlags config.Parameter `yaml:"additionalVcFlags,omitempty"`
}

// Generates a new Nimbus configuration
func NewNimbusConfig(cfg *StaderConfig) *NimbusConfig {
	return &NimbusConfig{
		Title: "Nimbus Settings",

		MaxPeers: config.Parameter{
			ID:                   "maxPeers",
			Name:                 "Max Peers",
			Description:          "The maximum number of peers your client should try to maintain. You can try lowering this if you have a low-resource system or a constrained network.",
			Type:                 config.ParameterType_Uint16,
			Default:              map[config.Network]interface{}{config.Network_All: getNimbusDefaultPeers()},
			AffectsContainers:    []config.ContainerID{config.ContainerID_Eth2},
			EnvironmentVariables: []string{"BN_MAX_PEERS"},
			CanBeBlank:           false,
			OverwriteOnUpgrade:   false,
		},

		BnContainerTag: config.Parameter{
			ID:          "bnContainerTag",
			Name:        "Beacon Node Container Tag",
			Description: "The tag name of the Nimbus Beacon Node container you want to use on Docker Hub.",
			Type:        config.ParameterType_String,
			Default: map[config.Network]interface{}{
				config.Network_Mainnet:  nimbusBnTagProd,
				config.Network_Prater:   nimbusBnTagTest,
				config.Network_Devnet:   nimbusBnTagTest,
				config.Network_Zhejiang: nimbusBnTagTest,
			},
			AffectsContainers:    []config.ContainerID{config.ContainerID_Eth2},
			EnvironmentVariables: []string{"BN_CONTAINER_TAG"},
			CanBeBlank:           false,
			OverwriteOnUpgrade:   true,
		},

		VcContainerTag: config.Parameter{
			ID:          "containerTag",
			Name:        "Validator Client Container Tag",
			Description: "The tag name of the Nimbus Validator Client container you want to use on Docker Hub.",
			Type:        config.ParameterType_String,
			Default: map[config.Network]interface{}{
				config.Network_Mainnet:  nimbusVcTagProd,
				config.Network_Prater:   nimbusVcTagTest,
				config.Network_Devnet:   nimbusVcTagTest,
				config.Network_Zhejiang: nimbusVcTagTest,
			},
			AffectsContainers:    []config.ContainerID{config.ContainerID_Validator},
			EnvironmentVariables: []string{"VC_CONTAINER_TAG"},
			CanBeBlank:           false,
			OverwriteOnUpgrade:   true,
		},

		PruningMode: config.Parameter{
			ID:                   "pruningMode",
			Name:                 "Pruning Mode",
			Description:          "Choose how Nimbus will prune its database. Highlight each option to learn more about it.",
			Type:                 config.ParameterType_Choice,
			Default:              map[config.Network]interface{}{config.Network_All: config.NimbusPruningMode_Archive},
			AffectsContainers:    []config.ContainerID{config.ContainerID_Eth2},
			EnvironmentVariables: []string{"NIMBUS_PRUNING_MODE"},
			CanBeBlank:           false,
			OverwriteOnUpgrade:   false,
			Options: []config.ParameterOption{{
				Name:        "Archive",
				Description: "Nimbus will download the entire Beacon Chain history and store it forever. This is healthier for the overall network, since people will be able to sync the entire chain from scratch using your node.",
				Value:       config.NimbusPruningMode_Archive,
			}, {
				Name:        "Pruned",
				Description: "Nimbus will only keep the last 5 months of data available, and will delete everything older than that. This will make Nimbus use less disk space overall, but you won't be able to access state older than 5 months (such as regenerating old rewards trees).\n\n[orange]WARNING: Pruning an *existing* database will take a VERY long time when Nimbus first starts. If you change from Archive to Pruned, you should delete your old chain data and do a checkpoint sync using `stader-cli service resync-eth2`. Make sure you have a checkpoint sync provider specified first!",
				Value:       config.NimbusPruningMode_Prune,
			}},
		},

		AdditionalBnFlags: config.Parameter{
			ID:                   "additionalBnFlags",
			Name:                 "Additional Beacon Client Flags",
			Description:          "Additional custom command line flags you want to pass Nimbus's Beacon Client, to take advantage of other settings that the Stadernode's configuration doesn't cover.",
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
			Description:          "Additional custom command line flags you want to pass Nimbus's Validator Client, to take advantage of other settings that the Stadernode's configuration doesn't cover.",
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
func (cfg *NimbusConfig) GetParameters() []*config.Parameter {
	return []*config.Parameter{
		&cfg.MaxPeers,
		&cfg.PruningMode,
		&cfg.BnContainerTag,
		&cfg.VcContainerTag,
		&cfg.AdditionalBnFlags,
		&cfg.AdditionalVcFlags,
	}
}

// Get the common params that this client doesn't support
func (cfg *NimbusConfig) GetUnsupportedCommonParams() []string {
	return cfg.UnsupportedCommonParams
}

// Get the Docker container name of the validator client
func (cfg *NimbusConfig) GetValidatorImage() string {
	return cfg.VcContainerTag.Value.(string)
}

// Get the name of the client
func (cfg *NimbusConfig) GetName() string {
	return "Nimbus"
}

// The the title for the config
func (cfg *NimbusConfig) GetConfigTitle() string {
	return cfg.Title
}

func getNimbusDefaultPeers() uint16 {
	if runtime.GOARCH == "arm64" {
		return defaultNimbusMaxPeersArm
	}

	return defaultNimbusMaxPeersAmd
}
