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
	"github.com/pbnjay/memory"
	"github.com/stader-labs/stader-node/shared/types/config"
)

const (
	tekuTagTest         string = "consensys/teku:23.4.0"
	tekuTagProd         string = "consensys/teku:23.4.0"
	defaultTekuMaxPeers uint16 = 100
)

// Configuration for Teku
type TekuConfig struct {
	Title string `yaml:"-"`

	// Common parameters that Teku doesn't support and should be hidden
	UnsupportedCommonParams []string `yaml:"-"`

	// Max number of P2P peers to connect to
	JvmHeapSize config.Parameter `yaml:"jvmHeapSize,omitempty"`

	// The max number of P2P peers to connect to
	MaxPeers config.Parameter `yaml:"maxPeers,omitempty"`

	// The archive mode flag
	ArchiveMode config.Parameter `yaml:"archiveMode,omitempty"`

	// The Docker Hub tag for Lighthouse
	ContainerTag config.Parameter `yaml:"containerTag,omitempty"`

	// Custom command line flags for the BN
	AdditionalBnFlags config.Parameter `yaml:"additionalBnFlags,omitempty"`

	// Custom command line flags for the VC
	AdditionalVcFlags config.Parameter `yaml:"additionalVcFlags,omitempty"`
}

// Generates a new Teku configuration
func NewTekuConfig(cfg *StaderConfig) *TekuConfig {
	return &TekuConfig{
		Title: "Teku Settings",

		UnsupportedCommonParams: []string{
			DoppelgangerDetectionID,
		},

		JvmHeapSize: config.Parameter{
			ID:                   "jvmHeapSize",
			Name:                 "JVM Heap Size",
			Description:          "The max amount of RAM, in MB, that Teku's JVM should limit itself to. Setting this lower will cause Teku to use less RAM, though it will always use more than this limit.\n\nUse 0 for automatic allocation.",
			Type:                 config.ParameterType_Uint,
			Default:              map[config.Network]interface{}{config.Network_All: getTekuHeapSize()},
			AffectsContainers:    []config.ContainerID{config.ContainerID_Eth1},
			EnvironmentVariables: []string{"TEKU_JVM_HEAP_SIZE"},
			CanBeBlank:           false,
			OverwriteOnUpgrade:   false,
		},

		MaxPeers: config.Parameter{
			ID:                   "maxPeers",
			Name:                 "Max Peers",
			Description:          "The maximum number of peers your client should try to maintain. You can try lowering this if you have a low-resource system or a constrained network.",
			Type:                 config.ParameterType_Uint16,
			Default:              map[config.Network]interface{}{config.Network_All: defaultTekuMaxPeers},
			AffectsContainers:    []config.ContainerID{config.ContainerID_Eth2},
			EnvironmentVariables: []string{"BN_MAX_PEERS"},
			CanBeBlank:           false,
			OverwriteOnUpgrade:   false,
		},

		ArchiveMode: config.Parameter{
			ID:                   "archiveMode",
			Name:                 "Enable Archive Mode",
			Description:          "When enabled, Teku will run in \"archive\" mode which means it can recreate the state of the Beacon chain for a previous block. This is required for manually generating the Merkle rewards tree.\n\nIf you are sure you will never be manually generating a tree, you can disable archive mode.",
			Type:                 config.ParameterType_Bool,
			Default:              map[config.Network]interface{}{config.Network_All: false},
			AffectsContainers:    []config.ContainerID{config.ContainerID_Eth2},
			EnvironmentVariables: []string{"TEKU_ARCHIVE_MODE"},
			CanBeBlank:           false,
			OverwriteOnUpgrade:   false,
		},

		ContainerTag: config.Parameter{
			ID:          "containerTag",
			Name:        "Container Tag",
			Description: "The tag name of the Teku container you want to use on Docker Hub.",
			Type:        config.ParameterType_String,
			Default: map[config.Network]interface{}{
				config.Network_Mainnet:  tekuTagProd,
				config.Network_Prater:   tekuTagTest,
				config.Network_Devnet:   tekuTagTest,
				config.Network_Zhejiang: tekuTagTest,
			},
			AffectsContainers:    []config.ContainerID{config.ContainerID_Eth2, config.ContainerID_Validator},
			EnvironmentVariables: []string{"BN_CONTAINER_TAG", "VC_CONTAINER_TAG"},
			CanBeBlank:           false,
			OverwriteOnUpgrade:   true,
		},

		AdditionalBnFlags: config.Parameter{
			ID:                   "additionalBnFlags",
			Name:                 "Additional Beacon Node Flags",
			Description:          "Additional custom command line flags you want to pass Teku's Beacon Node, to take advantage of other settings that the Stadernode's configuration doesn't cover.",
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
			Description:          "Additional custom command line flags you want to pass Teku's Validator Client, to take advantage of other settings that the Stadernode's configuration doesn't cover.",
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
func (cfg *TekuConfig) GetParameters() []*config.Parameter {
	return []*config.Parameter{
		&cfg.JvmHeapSize,
		&cfg.MaxPeers,
		&cfg.ArchiveMode,
		&cfg.ContainerTag,
		&cfg.AdditionalBnFlags,
		&cfg.AdditionalVcFlags,
	}
}

// Get the recommended heap size for Teku
func getTekuHeapSize() uint64 {
	totalMemoryGB := memory.TotalMemory() / 1024 / 1024 / 1024
	if totalMemoryGB < 9 {
		return 2048
	}
	return 0
}

// Get the common params that this client doesn't support
func (cfg *TekuConfig) GetUnsupportedCommonParams() []string {
	return cfg.UnsupportedCommonParams
}

// Get the Docker container name of the validator client
func (cfg *TekuConfig) GetValidatorImage() string {
	return cfg.ContainerTag.Value.(string)
}

// Get the name of the client
func (cfg *TekuConfig) GetName() string {
	return "Teku"
}

// The the title for the config
func (cfg *TekuConfig) GetConfigTitle() string {
	return cfg.Title
}
