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
	"github.com/stader-labs/stader-node/shared/types/config"
)

// Param IDs
const GraffitiID string = "graffiti"
const CheckpointSyncUrlID string = "checkpointSyncUrl"
const P2pPortID string = "p2pPort"
const ApiPortID string = "apiPort"
const OpenApiPortID string = "openApiPort"
const DoppelgangerDetectionID string = "doppelgangerDetection"

// Defaults
const defaultGraffiti string = "StaderLabs"
const defaultCheckpointSyncProvider string = ""
const defaultP2pPort uint16 = 9001
const defaultBnApiPort uint16 = 5052
const defaultOpenBnApiPort bool = false
const defaultDoppelgangerDetection bool = true

// Env var names
const CustomGraffitiEnvVar string = "CUSTOM_GRAFFITI"

// Common parameters shared by all of the Beacon Clients
type ConsensusCommonConfig struct {
	Title string `yaml:"-"`

	// Custom proposal graffiti
	Graffiti config.Parameter `yaml:"graffiti,omitempty"`

	// The checkpoint sync URL if used
	CheckpointSyncProvider config.Parameter `yaml:"checkpointSyncProvider,omitempty"`

	// The port to use for gossip traffic
	P2pPort config.Parameter `yaml:"p2pPort,omitempty"`

	// The port to expose the HTTP API on
	ApiPort config.Parameter `yaml:"apiPort,omitempty"`

	// Toggle for forwarding the HTTP API port outside of Docker
	OpenApiPort config.Parameter `yaml:"openApiPort,omitempty"`

	// Toggle for enabling doppelganger detection
	DoppelgangerDetection config.Parameter `yaml:"doppelgangerDetection,omitempty"`
}

// Create a new ConsensusCommonParams struct
func NewConsensusCommonConfig(cfg *StaderConfig) *ConsensusCommonConfig {
	return &ConsensusCommonConfig{
		Title: "Common Consensus Client Settings",

		Graffiti: config.Parameter{
			ID:                   GraffitiID,
			Name:                 "Custom Graffiti",
			Description:          "Add a short message to any blocks you propose, so the world can see what you have to say!\nIt has a 16 character limit.",
			Type:                 config.ParameterType_String,
			Default:              map[config.Network]interface{}{config.Network_All: defaultGraffiti},
			MaxLength:            16,
			AffectsContainers:    []config.ContainerID{config.ContainerID_Validator},
			EnvironmentVariables: []string{CustomGraffitiEnvVar},
			CanBeBlank:           true,
			OverwriteOnUpgrade:   false,
		},

		CheckpointSyncProvider: config.Parameter{
			ID:   CheckpointSyncUrlID,
			Name: "Checkpoint Sync URL",
			Description: "If you would like to instantly sync using an existing Beacon node, enter its URL.\n" +
				"Example: https://<project ID>:<secret>@eth2-beacon-prater.infura.io\n" +
				"Leave this blank if you want to sync normally from the start of the chain.",
			Type:                 config.ParameterType_String,
			Default:              map[config.Network]interface{}{config.Network_All: defaultCheckpointSyncProvider},
			AffectsContainers:    []config.ContainerID{config.ContainerID_Eth2},
			EnvironmentVariables: []string{"CHECKPOINT_SYNC_URL"},
			CanBeBlank:           true,
			OverwriteOnUpgrade:   false,
		},

		P2pPort: config.Parameter{
			ID:                   P2pPortID,
			Name:                 "P2P Port",
			Description:          "The port to use for P2P (blockchain) traffic.",
			Type:                 config.ParameterType_Uint16,
			Default:              map[config.Network]interface{}{config.Network_All: defaultP2pPort},
			AffectsContainers:    []config.ContainerID{config.ContainerID_Eth2},
			EnvironmentVariables: []string{"BN_P2P_PORT"},
			CanBeBlank:           false,
			OverwriteOnUpgrade:   false,
		},

		ApiPort: config.Parameter{
			ID:                   ApiPortID,
			Name:                 "HTTP API Port",
			Description:          "The port your Consensus client should run its HTTP API on.",
			Type:                 config.ParameterType_Uint16,
			Default:              map[config.Network]interface{}{config.Network_All: defaultBnApiPort},
			AffectsContainers:    []config.ContainerID{config.ContainerID_Api, config.ContainerID_Node, config.ContainerID_Guardian, config.ContainerID_Eth2, config.ContainerID_Validator, config.ContainerID_Prometheus},
			EnvironmentVariables: []string{"BN_API_PORT"},
			CanBeBlank:           false,
			OverwriteOnUpgrade:   false,
		},

		OpenApiPort: config.Parameter{
			ID:                   OpenApiPortID,
			Name:                 "Expose API Port",
			Description:          "Enable this to expose your Consensus client's API port to your local network, so other machines can access it too.",
			Type:                 config.ParameterType_Bool,
			Default:              map[config.Network]interface{}{config.Network_All: defaultOpenBnApiPort},
			AffectsContainers:    []config.ContainerID{config.ContainerID_Eth2},
			EnvironmentVariables: []string{"BN_OPEN_API_PORT"},
			CanBeBlank:           false,
			OverwriteOnUpgrade:   false,
		},

		DoppelgangerDetection: config.Parameter{
			ID:                   DoppelgangerDetectionID,
			Name:                 "Enable Doppelgänger Detection",
			Description:          "If enabled, your client will *intentionally* miss 1 or 2 attestations on startup to check if validator keys are already running elsewhere. If they are, it will disable validation duties for them to prevent you from being slashed.",
			Type:                 config.ParameterType_Bool,
			Default:              map[config.Network]interface{}{config.Network_All: defaultDoppelgangerDetection},
			AffectsContainers:    []config.ContainerID{config.ContainerID_Validator},
			EnvironmentVariables: []string{"DOPPELGANGER_DETECTION"},
			CanBeBlank:           false,
			OverwriteOnUpgrade:   false,
		},
	}
}

// Get the parameters for this config
func (cfg *ConsensusCommonConfig) GetParameters() []*config.Parameter {
	return []*config.Parameter{
		&cfg.Graffiti,
		&cfg.CheckpointSyncProvider,
		&cfg.P2pPort,
		&cfg.ApiPort,
		&cfg.OpenApiPort,
		&cfg.DoppelgangerDetection,
	}
}

// The the title for the config
func (cfg *ConsensusCommonConfig) GetConfigTitle() string {
	return cfg.Title
}
