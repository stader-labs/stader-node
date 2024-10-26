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
	"fmt"
	"runtime"

	"github.com/stader-labs/stader-node/shared/types/config"
)

const (
	prysmBnTagTest string = "staderlabs/prysm:v5.1.2"
	prysmVcTagTest string = "staderlabs/prysm:v5.1.2"

	prysmBnTagProd string = "staderlabs/prysm:v5.1.2"
	prysmVcTagProd string = "staderlabs/prysm:v5.1.2"

	defaultPrysmRpcPort     uint16 = 5053
	defaultPrysmOpenRpcPort bool   = false
	defaultPrysmMaxPeers    uint16 = 70
)

// Configuration for Prysm
type PrysmConfig struct {
	Title string `yaml:"title,omitempty"`

	// Common parameters that Prysm doesn't support and should be hidden
	UnsupportedCommonParams []string `yaml:"unsupportedCommonParams,omitempty"`

	// The max number of P2P peers to connect to
	MaxPeers config.Parameter `yaml:"maxPeers,omitempty"`

	// The RPC port for BN / VC connections
	RpcPort config.Parameter `yaml:"rpcPort,omitempty"`

	// Toggle for forwarding the RPC API outside of Docker
	OpenRpcPort config.Parameter `yaml:"openRpcPort,omitempty"`

	// The Docker Hub tag for the Prysm BN
	BnContainerTag config.Parameter `yaml:"bnContainerTag,omitempty"`

	// The Docker Hub tag for the Prysm VC
	VcContainerTag config.Parameter `yaml:"vcContainerTag,omitempty"`

	// Custom command line flags for the BN
	AdditionalBnFlags config.Parameter `yaml:"additionalBnFlags,omitempty"`

	// Custom command line flags for the VC
	AdditionalVcFlags config.Parameter `yaml:"additionalVcFlags,omitempty"`
}

// Generates a new Prysm configuration
func NewPrysmConfig(cfg *StaderConfig) *PrysmConfig {
	return &PrysmConfig{
		Title: "Prysm Settings",

		UnsupportedCommonParams: []string{},

		MaxPeers: config.Parameter{
			ID:                   "maxPeers",
			Name:                 "Max Peers",
			Description:          "The maximum number of peers your client should try to maintain. You can try lowering this if you have a low-resource system or a constrained network.",
			Type:                 config.ParameterType_Uint16,
			Default:              map[config.Network]interface{}{config.Network_All: defaultPrysmMaxPeers},
			AffectsContainers:    []config.ContainerID{config.ContainerID_Eth2},
			EnvironmentVariables: []string{"BN_MAX_PEERS"},
			CanBeBlank:           false,
			OverwriteOnUpgrade:   false,
		},

		RpcPort: config.Parameter{
			ID:                   "rpcPort",
			Name:                 "RPC Port",
			Description:          "The port Prysm should run its JSON-RPC API on.",
			Type:                 config.ParameterType_Uint16,
			Default:              map[config.Network]interface{}{config.Network_All: defaultPrysmRpcPort},
			AffectsContainers:    []config.ContainerID{config.ContainerID_Eth2, config.ContainerID_Validator},
			EnvironmentVariables: []string{"BN_RPC_PORT"},
			CanBeBlank:           false,
			OverwriteOnUpgrade:   false,
		},

		OpenRpcPort: config.Parameter{
			ID:                   "openRpcPort",
			Name:                 "Expose RPC Port",
			Description:          "Enable this to expose Prysm's JSON-RPC port to your local network, so other machines can access it too.",
			Type:                 config.ParameterType_Bool,
			Default:              map[config.Network]interface{}{config.Network_All: defaultPrysmOpenRpcPort},
			AffectsContainers:    []config.ContainerID{config.ContainerID_Eth2},
			EnvironmentVariables: []string{"BN_OPEN_RPC_PORT"},
			CanBeBlank:           false,
			OverwriteOnUpgrade:   false,
		},

		BnContainerTag: config.Parameter{
			ID:          "bnContainerTag",
			Name:        "Beacon Node Container Tag",
			Description: "The tag name of the Prysm Beacon Node container you want to use on Docker Hub.",
			Type:        config.ParameterType_String,
			Default: map[config.Network]interface{}{
				config.Network_Mainnet: getPrysmBnProdTag(),
				config.Network_Holesky: getPrysmBnTestTag(),
			},
			AffectsContainers:    []config.ContainerID{config.ContainerID_Eth2},
			EnvironmentVariables: []string{"BN_CONTAINER_TAG"},
			CanBeBlank:           false,
			OverwriteOnUpgrade:   true,
		},

		VcContainerTag: config.Parameter{
			ID:          "vcContainerTag",
			Name:        "Validator Client Container Tag",
			Description: "The tag name of the Prysm Validator Client container you want to use on Docker Hub.",
			Type:        config.ParameterType_String,
			Default: map[config.Network]interface{}{
				config.Network_Mainnet: getPrysmVcProdTag(),
				config.Network_Holesky: getPrysmVcTestTag(),
			},
			AffectsContainers:    []config.ContainerID{config.ContainerID_Validator},
			EnvironmentVariables: []string{"VC_CONTAINER_TAG"},
			CanBeBlank:           false,
			OverwriteOnUpgrade:   true,
		},

		AdditionalBnFlags: config.Parameter{
			ID:                   "additionalBnFlags",
			Name:                 "Additional Beacon Node Flags",
			Description:          "Additional custom command line flags you want to pass Prysm's Beacon Node, to take advantage of other settings that the Stadernode's configuration doesn't cover.",
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
			Description:          "Additional custom command line flags you want to pass Prysm's Validator Client, to take advantage of other settings that the Stadernode's configuration doesn't cover.",
			Type:                 config.ParameterType_String,
			Default:              map[config.Network]interface{}{config.Network_All: ""},
			AffectsContainers:    []config.ContainerID{config.ContainerID_Validator},
			EnvironmentVariables: []string{"VC_ADDITIONAL_FLAGS"},
			CanBeBlank:           true,
			OverwriteOnUpgrade:   false,
		},
	}
}

// Get the container tag for the Prysm BN based on the current architecture
func getPrysmBnProdTag() string {
	if runtime.GOARCH == "arm64" || runtime.GOARCH == "amd64" {
		return prysmBnTagProd
	}
	panic(fmt.Sprintf("Prysm doesn't support architecture %s", runtime.GOARCH))
}

// Get the container tag for the Prysm BN based on the current architecture
func getPrysmBnTestTag() string {
	if runtime.GOARCH == "arm64" || runtime.GOARCH == "amd64" {
		return prysmBnTagTest
	}
	panic(fmt.Sprintf("Prysm doesn't support architecture %s", runtime.GOARCH))
}

// Get the container tag for the Prysm VC based on the current architecture
func getPrysmVcProdTag() string {
	if runtime.GOARCH == "arm64" || runtime.GOARCH == "amd64" {
		return prysmVcTagProd
	}
	panic(fmt.Sprintf("Prysm doesn't support architecture %s", runtime.GOARCH))
}

// Get the container tag for the Prysm VC based on the current architecture
func getPrysmVcTestTag() string {
	if runtime.GOARCH == "arm64" || runtime.GOARCH == "amd64" {
		return prysmVcTagTest
	}
	panic(fmt.Sprintf("Prysm doesn't support architecture %s", runtime.GOARCH))
}

// Get the parameters for this config
func (cfg *PrysmConfig) GetParameters() []*config.Parameter {
	return []*config.Parameter{
		&cfg.MaxPeers,
		&cfg.RpcPort,
		&cfg.OpenRpcPort,
		&cfg.BnContainerTag,
		&cfg.VcContainerTag,
		&cfg.AdditionalBnFlags,
		&cfg.AdditionalVcFlags,
	}
}

// Get the common params that this client doesn't support
func (cfg *PrysmConfig) GetUnsupportedCommonParams() []string {
	return cfg.UnsupportedCommonParams
}

// Get the Docker container name of the validator client
func (cfg *PrysmConfig) GetValidatorImage() string {
	return cfg.VcContainerTag.Value.(string)
}

// Get the name of the client
func (cfg *PrysmConfig) GetName() string {
	return "Prysm"
}

// The the title for the config
func (cfg *PrysmConfig) GetConfigTitle() string {
	return cfg.Title
}
