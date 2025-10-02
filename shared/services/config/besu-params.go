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

// Constants
const (
	besuTagTest          string = "hyperledger/besu:25.9.0"
	besuTagProd          string = "hyperledger/besu:25.9.0"
	besuEventLogInterval int    = 1000
	besuMaxPeers         uint16 = 25
	besuStopSignal       string = "SIGTERM"
)

// Configuration for Besu
type BesuConfig struct {
	Title string `yaml:"-"`

	// Common parameters that Besu doesn't support and should be hidden
	UnsupportedCommonParams []string `yaml:"-"`

	// Compatible consensus clients
	CompatibleConsensusClients []config.ConsensusClient `yaml:"-"`

	// Max number of P2P peers to connect to
	JvmHeapSize config.Parameter `yaml:"jvmHeapSize,omitempty"`

	// The max number of events to query in a single event log query
	EventLogInterval int `yaml:"-"`

	// Max number of P2P peers to connect to
	MaxPeers config.Parameter `yaml:"maxPeers,omitempty"`

	// Historical state block regeneration limit
	MaxBackLayers config.Parameter `yaml:"maxBackLayers,omitempty"`

	// The Docker Hub tag for Besu
	ContainerTag config.Parameter `yaml:"containerTag,omitempty"`

	// Custom command line flags
	AdditionalFlags config.Parameter `yaml:"additionalFlags,omitempty"`
}

// Generates a new Besu configuration
func NewBesuConfig(cfg *StaderConfig) *BesuConfig {
	return &BesuConfig{
		Title: "Besu Settings",

		UnsupportedCommonParams: []string{},

		CompatibleConsensusClients: []config.ConsensusClient{
			config.ConsensusClient_Lighthouse,
			config.ConsensusClient_Nimbus,
			config.ConsensusClient_Prysm,
			config.ConsensusClient_Teku,
		},

		EventLogInterval: besuEventLogInterval,

		JvmHeapSize: config.Parameter{
			ID:                   "jvmHeapSize",
			Name:                 "JVM Heap Size",
			Description:          "The max amount of RAM, in MB, that Besu's JVM should limit itself to. Setting this lower will cause Besu to use less RAM, though it will always use more than this limit.\n\nUse 0 for automatic allocation.",
			Type:                 config.ParameterType_Uint,
			Default:              map[config.Network]interface{}{config.Network_All: uint64(0)},
			AffectsContainers:    []config.ContainerID{config.ContainerID_Eth1},
			EnvironmentVariables: []string{"BESU_JVM_HEAP_SIZE"},
			CanBeBlank:           false,
			OverwriteOnUpgrade:   false,
		},

		MaxPeers: config.Parameter{
			ID:                   "maxPeers",
			Name:                 "Max Peers",
			Description:          "The maximum number of peers Besu should connect to. This can be lowered to improve performance on low-power systems or constrained networks. We recommend keeping it at 12 or higher.",
			Type:                 config.ParameterType_Uint16,
			Default:              map[config.Network]interface{}{config.Network_All: besuMaxPeers},
			AffectsContainers:    []config.ContainerID{config.ContainerID_Eth1},
			EnvironmentVariables: []string{"EC_MAX_PEERS"},
			CanBeBlank:           false,
			OverwriteOnUpgrade:   false,
		},

		MaxBackLayers: config.Parameter{
			ID:                   "maxBackLayers",
			Name:                 "Historical Block Replay Limit",
			Description:          "Besu has the ability to revisit the state of any historical block on the chain by \"replaying\" all of the previous blocks to get back to the target. This limit controls how many blocks you can replay - in other words, how far back Besu can go in time. Normal Execution client processing will be paused while a replay is in progress.\n\n[orange]NOTE: If you try to replay a state from a long time ago, it may take Besu several minutes to rebuild the state!",
			Type:                 config.ParameterType_Uint,
			Default:              map[config.Network]interface{}{config.Network_All: uint64(512)},
			AffectsContainers:    []config.ContainerID{config.ContainerID_Eth1},
			EnvironmentVariables: []string{"BESU_MAX_BACK_LAYERS"},
			CanBeBlank:           false,
			OverwriteOnUpgrade:   false,
		},

		ContainerTag: config.Parameter{
			ID:          "containerTag",
			Name:        "Container Tag",
			Description: "The tag name of the Besu container you want to use on Docker Hub.",
			Type:        config.ParameterType_String,
			Default: map[config.Network]interface{}{
				config.Network_Mainnet: besuTagProd,
				config.Network_Holesky: besuTagTest,
			},
			AffectsContainers:    []config.ContainerID{config.ContainerID_Eth1},
			EnvironmentVariables: []string{"EC_CONTAINER_TAG"},
			CanBeBlank:           false,
			OverwriteOnUpgrade:   true,
		},

		AdditionalFlags: config.Parameter{
			ID:                   "additionalFlags",
			Name:                 "Additional Flags",
			Description:          "Additional custom command line flags you want to pass to Besu, to take advantage of other settings that the Stadernode's configuration doesn't cover.",
			Type:                 config.ParameterType_String,
			Default:              map[config.Network]interface{}{config.Network_All: ""},
			AffectsContainers:    []config.ContainerID{config.ContainerID_Eth1},
			EnvironmentVariables: []string{"EC_ADDITIONAL_FLAGS"},
			CanBeBlank:           true,
			OverwriteOnUpgrade:   false,
		},
	}
}

// Get the parameters for this config
func (cfg *BesuConfig) GetParameters() []*config.Parameter {
	return []*config.Parameter{
		&cfg.JvmHeapSize,
		&cfg.MaxPeers,
		&cfg.MaxBackLayers,
		&cfg.ContainerTag,
		&cfg.AdditionalFlags,
	}
}

// The the title for the config
func (cfg *BesuConfig) GetConfigTitle() string {
	return cfg.Title
}
