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
	"strings"

	"github.com/stader-labs/stader-node/shared/types/config"
)

// Constants
const (
	mevBoostTagProd             string = "flashbots/mev-boost:1.9"
	mevBoostTagTest             string = "flashbots/mev-boost:1.9"
	mevBoostUrlEnvVar           string = "MEV_BOOST_URL"
	mevBoostRelaysEnvVar        string = "MEV_BOOST_RELAYS"
	mevDocsUrl                  string = "#"
	RegulatedRelayDescription   string = "Select this to enable the relays that comply with government regulations (e.g. OFAC sanctions), "
	UnregulatedRelayDescription string = "Select this to enable the relays that do not follow any sanctions lists (do not censor transactions), "
	NoSandwichRelayDescription  string = "and do not allow front-running or sandwich attacks."
	AllMevRelayDescription      string = "and allow for all types of MEV (including sandwich attacks)."
)

// Configuration for MEV-Boost
type MevBoostConfig struct {
	Title string `yaml:"-"`

	// Ownership mode
	Mode config.Parameter `yaml:"mode,omitempty"`

	// The mode for relay selection
	SelectionMode config.Parameter `yaml:"selectionMode,omitempty"`

	// Regulated, all types
	EnableRegulatedAllMev config.Parameter `yaml:"enableRegulatedAllMev,omitempty"`

	// Regulated, no sandwiching
	EnableRegulatedNoSandwich config.Parameter `yaml:"enableRegulatedNoSandwich,omitempty"`

	// Unregulated, all types
	EnableUnregulatedAllMev config.Parameter `yaml:"enableUnregulatedAllMev,omitempty"`

	// Unregulated, no sandwiching
	EnableUnregulatedNoSandwich config.Parameter `yaml:"enableUnregulatedNoSandwich,omitempty"`

	// Flashbots relay
	FlashbotsRelay config.Parameter `yaml:"flashbotsEnabled,omitempty"`

	// bloXroute max profit relay
	BloxRouteMaxProfitRelay config.Parameter `yaml:"bloxRouteMaxProfitEnabled,omitempty"`

	// bloXroute regulated relay
	BloxRouteRegulatedRelay config.Parameter `yaml:"bloxRouteRegulatedEnabled,omitempty"`

	// Eden relay
	EdenRelay config.Parameter `yaml:"edenEnabled,omitempty"`

	// Ultra sound relay
	UltrasoundRelay config.Parameter `yaml:"ultrasoundEnabled,omitempty"`

	// Aestus relay
	AestusRelay config.Parameter `yaml:"aestusEnabled,omitempty"`

	// Agnostic relay
	AgnosticRelay config.Parameter `yaml:"AgnoticEnabled,omitempty"`

	// Titan regulated relay
	TitanRegulatedRelay config.Parameter `yaml:"TitanRegulatedEnable,omitempty"`

	// Titan unregulated relay
	TitanUnRegulatedRelay config.Parameter `yaml:"TitanUnRegulatedEnable,omitempty"`

	// The RPC port
	Port config.Parameter `yaml:"port,omitempty"`

	// Toggle for forwarding the HTTP port outside of Docker
	OpenRpcPort config.Parameter `yaml:"openRpcPort,omitempty"`

	// The Docker Hub tag for MEV-Boost
	ContainerTag config.Parameter `yaml:"containerTag,omitempty"`

	// Custom command line flags
	AdditionalFlags config.Parameter `yaml:"additionalFlags,omitempty"`

	// The URL of an external MEV-Boost client
	ExternalUrl config.Parameter `yaml:"externalUrl"`

	///////////////////////////
	// Non-editable settings //
	///////////////////////////

	parentConfig *StaderConfig                         `yaml:"-"`
	relays       []config.MevRelay                     `yaml:"-"`
	relayMap     map[config.MevRelayID]config.MevRelay `yaml:"-"`
}

// Generates a new MEV-Boost configuration
func NewMevBoostConfig(cfg *StaderConfig) *MevBoostConfig {
	// Generate the relays
	relays := createDefaultRelays()
	relayMap := map[config.MevRelayID]config.MevRelay{}
	for _, relay := range relays {
		relayMap[relay.ID] = relay
	}

	return &MevBoostConfig{
		Title: "MEV-Boost Settings",

		parentConfig: cfg,

		Mode: config.Parameter{
			ID:                   "mode",
			Name:                 "MEV-Boost Mode",
			Description:          "Choose whether to let the Stadernode manage your MEV-Boost instance (Locally Managed), or if you manage your own outside of the Stadernode stack (Externally Managed).",
			Type:                 config.ParameterType_Choice,
			Default:              map[config.Network]interface{}{config.Network_All: config.Mode_Local},
			AffectsContainers:    []config.ContainerID{config.ContainerID_Eth2, config.ContainerID_MevBoost},
			EnvironmentVariables: []string{},
			CanBeBlank:           false,
			OverwriteOnUpgrade:   false,
			Options: []config.ParameterOption{{
				Name:        "Locally Managed",
				Description: "Allow the Stadernode to manage the MEV-Boost client for you",
				Value:       config.Mode_Local,
			}, {
				Name:        "Externally Managed",
				Description: "Use an existing MEV-Boost client that you manage on your own",
				Value:       config.Mode_External,
			}},
		},

		SelectionMode: config.Parameter{
			ID:                   "selectionMode",
			Name:                 "Selection Mode",
			Description:          "Select how the TUI shows you the options for which MEV relays to enable.",
			Type:                 config.ParameterType_Choice,
			Default:              map[config.Network]interface{}{config.Network_All: config.MevSelectionMode_Profile},
			AffectsContainers:    []config.ContainerID{config.ContainerID_MevBoost},
			EnvironmentVariables: []string{},
			CanBeBlank:           false,
			OverwriteOnUpgrade:   false,
			Options: []config.ParameterOption{{
				Name:        "Profile Mode",
				Description: "Relays will be bundled up based on whether or not they're regulated, and whether or not they allow sandwich attacks.\nUse this if you simply want to specify which type of relay you want to use without needing to read about each individual relay.",
				Value:       config.MevSelectionMode_Profile,
			}, {
				Name:        "Relay Mode",
				Description: "Each relay will be shown, and you can enable each one individually as you see fit.\nUse this if you already know about the relays and want to customize the ones you will use.",
				Value:       config.MevSelectionMode_Relay,
			}},
		},

		EnableRegulatedAllMev:       generateProfileParameter("enableRegulatedAllMev", relays, true, false),
		EnableRegulatedNoSandwich:   generateProfileParameter("enableRegulatedNoSandwich", relays, true, true),
		EnableUnregulatedAllMev:     generateProfileParameter("enableUnregulatedAllMev", relays, false, false),
		EnableUnregulatedNoSandwich: generateProfileParameter("enableUnregulatedNoSandwich", relays, false, true),

		// Explicit relay params
		FlashbotsRelay:          generateRelayParameter("flashbotsEnabled", relayMap[config.MevRelayID_Flashbots]),
		BloxRouteMaxProfitRelay: generateRelayParameter("bloxRouteMaxProfitEnabled", relayMap[config.MevRelayID_BloxrouteMaxProfit]),
		BloxRouteRegulatedRelay: generateRelayParameter("bloxRouteRegulatedEnabled", relayMap[config.MevRelayID_BloxrouteRegulated]),
		EdenRelay:               generateRelayParameter("edenEnabled", relayMap[config.MevRelayID_Eden]),
		UltrasoundRelay:         generateRelayParameter("ultrasoundEnabled", relayMap[config.MevRelayID_Ultrasound]),
		AestusRelay:             generateRelayParameter("aestusEnabled", relayMap[config.MevRelayID_Aestus]),
		AgnosticRelay:           generateRelayParameter("agnosticEnabled", relayMap[config.MevRelayID_Agnostic]),
		TitanRegulatedRelay:     generateRelayParameter("titanRegulatedEnabled", relayMap[config.MevRelayID_TitanRegulated]),
		TitanUnRegulatedRelay:   generateRelayParameter("titanUnRegulatedEnabled", relayMap[config.MevRelayID_TitanUnRegulated]),

		Port: config.Parameter{
			ID:                   "port",
			Name:                 "Port",
			Description:          "The port that MEV-Boost should serve its API on.",
			Type:                 config.ParameterType_Uint16,
			Default:              map[config.Network]interface{}{config.Network_All: uint16(18550)},
			AffectsContainers:    []config.ContainerID{config.ContainerID_Eth2, config.ContainerID_MevBoost},
			EnvironmentVariables: []string{"MEV_BOOST_PORT"},
			CanBeBlank:           false,
			OverwriteOnUpgrade:   false,
		},

		OpenRpcPort: config.Parameter{
			ID:                   "openRpcPort",
			Name:                 "Expose API Port",
			Description:          "Expose the API port to your local network, so other local machines can access MEV-Boost's API.",
			Type:                 config.ParameterType_Bool,
			Default:              map[config.Network]interface{}{config.Network_All: false},
			AffectsContainers:    []config.ContainerID{config.ContainerID_MevBoost},
			EnvironmentVariables: []string{},
			CanBeBlank:           false,
			OverwriteOnUpgrade:   false,
		},

		ContainerTag: config.Parameter{
			ID:          "containerTag",
			Name:        "Container Tag",
			Description: "The tag name of the MEV-Boost container you want to use on Docker Hub.",
			Type:        config.ParameterType_String,
			Default: map[config.Network]interface{}{
				config.Network_Mainnet: mevBoostTagProd,
				config.Network_Holesky: mevBoostTagTest,
			},
			AffectsContainers:    []config.ContainerID{config.ContainerID_MevBoost},
			EnvironmentVariables: []string{"MEV_BOOST_CONTAINER_TAG"},
			CanBeBlank:           false,
			OverwriteOnUpgrade:   true,
		},

		AdditionalFlags: config.Parameter{
			ID:                   "additionalFlags",
			Name:                 "Additional Flags",
			Description:          "Additional custom command line flags you want to pass to MEV-Boost, to take advantage of other settings that the Stadernode's configuration doesn't cover.",
			Type:                 config.ParameterType_String,
			Default:              map[config.Network]interface{}{config.Network_All: ""},
			AffectsContainers:    []config.ContainerID{config.ContainerID_MevBoost},
			EnvironmentVariables: []string{"MEV_BOOST_ADDITIONAL_FLAGS"},
			CanBeBlank:           true,
			OverwriteOnUpgrade:   false,
		},

		ExternalUrl: config.Parameter{
			ID:                   "externalUrl",
			Name:                 "External URL",
			Description:          "The URL of the external MEV-Boost client or provider",
			Type:                 config.ParameterType_String,
			Default:              map[config.Network]interface{}{config.Network_All: ""},
			AffectsContainers:    []config.ContainerID{config.ContainerID_Eth2},
			EnvironmentVariables: []string{mevBoostUrlEnvVar},
			CanBeBlank:           true,
			OverwriteOnUpgrade:   false,
		},

		relays:   relays,
		relayMap: relayMap,
	}
}

// Get the config.Parameters for this config
func (cfg *MevBoostConfig) GetParameters() []*config.Parameter {
	return []*config.Parameter{
		&cfg.Mode,
		&cfg.SelectionMode,
		&cfg.EnableRegulatedAllMev,
		&cfg.EnableRegulatedNoSandwich,
		&cfg.EnableUnregulatedAllMev,
		&cfg.EnableUnregulatedNoSandwich,
		&cfg.FlashbotsRelay,
		&cfg.BloxRouteMaxProfitRelay,
		&cfg.BloxRouteRegulatedRelay,
		&cfg.EdenRelay,
		&cfg.UltrasoundRelay,
		&cfg.AestusRelay,
		&cfg.AgnosticRelay,
		&cfg.TitanRegulatedRelay,
		&cfg.TitanUnRegulatedRelay,
		&cfg.Port,
		&cfg.OpenRpcPort,
		&cfg.ContainerTag,
		&cfg.AdditionalFlags,
		&cfg.ExternalUrl,
	}
}

// The title for the config
func (cfg *MevBoostConfig) GetConfigTitle() string {
	return cfg.Title
}

// Get the profiles that are available for the current network
func (cfg *MevBoostConfig) GetAvailableProfiles() (bool, bool, bool, bool) {
	regulatedAllMev := false
	regulatedNoSandwich := false
	unregulatedAllMev := false
	unregulatedNoSandwich := false

	currentNetwork := cfg.parentConfig.StaderNode.Network.Value.(config.Network)
	for _, relay := range cfg.relays {
		_, exists := relay.Urls[currentNetwork]
		if !exists {
			continue
		}
		regulatedAllMev = regulatedAllMev || (relay.Regulated && !relay.NoSandwiching)
		regulatedNoSandwich = regulatedNoSandwich || (relay.Regulated && relay.NoSandwiching)
		unregulatedAllMev = unregulatedAllMev || (!relay.Regulated && !relay.NoSandwiching)
		unregulatedNoSandwich = unregulatedNoSandwich || (!relay.Regulated && relay.NoSandwiching)
	}

	return regulatedAllMev, regulatedNoSandwich, unregulatedAllMev, unregulatedNoSandwich
}

// Get the relays that are available for the current network
func (cfg *MevBoostConfig) GetAvailableRelays() []config.MevRelay {
	relays := []config.MevRelay{}
	currentNetwork := cfg.parentConfig.StaderNode.Network.Value.(config.Network)
	for _, relay := range cfg.relays {
		_, exists := relay.Urls[currentNetwork]
		if !exists {
			continue
		}
		relays = append(relays, relay)
	}

	return relays
}

// Get which MEV-boost relays are enabled
func (cfg *MevBoostConfig) GetEnabledMevRelays() []config.MevRelay {
	relays := []config.MevRelay{}

	currentNetwork := cfg.parentConfig.StaderNode.Network.Value.(config.Network)
	switch cfg.SelectionMode.Value.(config.MevSelectionMode) {
	case config.MevSelectionMode_Profile:
		for _, relay := range cfg.relays {
			_, exists := relay.Urls[currentNetwork]
			if !exists {
				// Skip relays that don't exist on the current network
				continue
			}
			if relay.Regulated {
				if relay.NoSandwiching {
					if cfg.EnableRegulatedNoSandwich.Value == true {
						relays = append(relays, relay)
					}
				} else {
					if cfg.EnableRegulatedAllMev.Value == true {
						relays = append(relays, relay)
					}
				}
			} else {
				if relay.NoSandwiching {
					if cfg.EnableUnregulatedNoSandwich.Value == true {
						relays = append(relays, relay)
					}
				} else {
					if cfg.EnableUnregulatedAllMev.Value == true {
						relays = append(relays, relay)
					}
				}
			}
		}

	case config.MevSelectionMode_Relay:
		if cfg.FlashbotsRelay.Value == true {
			_, exists := cfg.relayMap[config.MevRelayID_Flashbots].Urls[currentNetwork]
			if exists {
				relays = append(relays, cfg.relayMap[config.MevRelayID_Flashbots])
			}
		}
		if cfg.BloxRouteMaxProfitRelay.Value == true {
			_, exists := cfg.relayMap[config.MevRelayID_BloxrouteMaxProfit].Urls[currentNetwork]
			if exists {
				relays = append(relays, cfg.relayMap[config.MevRelayID_BloxrouteMaxProfit])
			}
		}
		if cfg.BloxRouteRegulatedRelay.Value == true {
			_, exists := cfg.relayMap[config.MevRelayID_BloxrouteRegulated].Urls[currentNetwork]
			if exists {
				relays = append(relays, cfg.relayMap[config.MevRelayID_BloxrouteRegulated])
			}
		}

		if cfg.EdenRelay.Value == true {
			_, exists := cfg.relayMap[config.MevRelayID_Eden].Urls[currentNetwork]
			if exists {
				relays = append(relays, cfg.relayMap[config.MevRelayID_Eden])
			}
		}
		if cfg.UltrasoundRelay.Value == true {
			_, exists := cfg.relayMap[config.MevRelayID_Ultrasound].Urls[currentNetwork]
			if exists {
				relays = append(relays, cfg.relayMap[config.MevRelayID_Ultrasound])
			}
		}

		if cfg.AestusRelay.Value == true {
			_, exists := cfg.relayMap[config.MevRelayID_Aestus].Urls[currentNetwork]
			if exists {
				relays = append(relays, cfg.relayMap[config.MevRelayID_Aestus])
			}
		}

		if cfg.AgnosticRelay.Value == true {
			_, exists := cfg.relayMap[config.MevRelayID_Agnostic].Urls[currentNetwork]
			if exists {
				relays = append(relays, cfg.relayMap[config.MevRelayID_Agnostic])
			}
		}

		if cfg.TitanRegulatedRelay.Value == true {
			_, exists := cfg.relayMap[config.MevRelayID_TitanRegulated].Urls[currentNetwork]
			if exists {
				relays = append(relays, cfg.relayMap[config.MevRelayID_TitanRegulated])
			}
		}

		if cfg.TitanUnRegulatedRelay.Value == true {
			_, exists := cfg.relayMap[config.MevRelayID_TitanUnRegulated].Urls[currentNetwork]
			if exists {
				relays = append(relays, cfg.relayMap[config.MevRelayID_TitanUnRegulated])
			}
		}
	}

	return relays
}

func (cfg *MevBoostConfig) GetRelayString() string {
	relayUrls := []string{}
	currentNetwork := cfg.parentConfig.StaderNode.Network.Value.(config.Network)

	relays := cfg.GetEnabledMevRelays()
	for _, relay := range relays {
		relayUrls = append(relayUrls, relay.Urls[currentNetwork])
	}

	relayString := strings.Join(relayUrls, ",")
	return relayString
}

// Create the default MEV relays
func createDefaultRelays() []config.MevRelay {
	relays := []config.MevRelay{
		// Flashbots
		{
			ID:          config.MevRelayID_Flashbots,
			Name:        "Flashbots",
			Description: "Flashbots is the developer of MEV-Boost, and one of the best-known and most trusted relays in the space.",
			Urls: map[config.Network]string{
				config.Network_Mainnet: "https://0xac6e77dfe25ecd6110b8e780608cce0dab71fdd5ebea22a16c0205200f2f8e2e3ad3b71d3499c54ad14d6c21b41a37ae@boost-relay.flashbots.net?id=staderlabs",
				config.Network_Holesky: "https://0xafa4c6985aa049fb79dd37010438cfebeb0f2bd42b115b89dd678dab0670c1de38da0c4e9138c9290a398ecd9a0b3110@boost-relay-holesky.flashbots.net?id=staderlabs",
			},
			Regulated:     true,
			NoSandwiching: false,
		},

		// bloXroute Max Profit
		{
			ID:          config.MevRelayID_BloxrouteMaxProfit,
			Name:        "bloXroute Max Profit",
			Description: "Select this to enable the \"max profit\" relay from bloXroute.",
			Urls: map[config.Network]string{
				config.Network_Mainnet: "https://0x8b5d2e73e2a3a55c6c87b8b6eb92e0149a125c852751db1422fa951e42a09b82c142c3ea98d0d9930b056a3bc9896b8f@bloxroute.max-profit.blxrbdn.com?id=staderlabs",
			},
			Regulated:     true,
			NoSandwiching: false,
		},
		// bloXroute Regulated
		{
			ID:          config.MevRelayID_BloxrouteRegulated,
			Name:        "bloXroute Regulated",
			Description: "Select this to enable the \"regulated\" relay from bloXroute.",
			Urls: map[config.Network]string{
				config.Network_Mainnet: "https://0xb0b07cd0abef743db4260b0ed50619cf6ad4d82064cb4fbec9d3ec530f7c5e6793d9f286c4e082c0244ffb9f2658fe88@bloxroute.regulated.blxrbdn.com?id=staderlabs",
				config.Network_Holesky: "https://0x821f2a65afb70e7f2e820a925a9b4c80a159620582c1766b1b09729fec178b11ea22abb3a51f07b288be815a1a2ff516@bloxroute.holesky.blxrbdn.com?id=staderlabs",
			},
			Regulated:     true,
			NoSandwiching: false,
		},

		// Eden
		{
			ID:          config.MevRelayID_Eden,
			Name:        "Eden Network",
			Description: "Eden Network is the home of Eden Relay, a block building hub focused on optimising block rewards for validators.",
			Urls: map[config.Network]string{
				config.Network_Mainnet: "https://0xb3ee7afcf27f1f1259ac1787876318c6584ee353097a50ed84f51a1f21a323b3736f271a895c7ce918c038e4265918be@relay.edennetwork.io?id=staderlabs",
				config.Network_Holesky: "https://0xb1d229d9c21298a87846c7022ebeef277dfc321fe674fa45312e20b5b6c400bfde9383f801848d7837ed5fc449083a12@relay-holesky.edennetwork.io?id=staderlabs",
			},
			Regulated:     true,
			NoSandwiching: false,
		},

		// Ultrasound
		{
			ID:          config.MevRelayID_Ultrasound,
			Name:        "Ultra Sound",
			Description: "The ultra sound relay is a credibly-neutral and permissionless relay — a public good from the ultrasound.money team.",
			Urls: map[config.Network]string{
				config.Network_Mainnet: "https://0xa1559ace749633b997cb3fdacffb890aeebdb0f5a3b6aaa7eeeaf1a38af0a8fe88b9e4b1f61f236d2e64d95733327a62@relay.ultrasound.money?id=staderlabs",
				config.Network_Holesky: "https://0xb1559beef7b5ba3127485bbbb090362d9f497ba64e177ee2c8e7db74746306efad687f2cf8574e38d70067d40ef136dc@relay-stag.ultrasound.money?id=staderlabs",
			},
			Regulated:     false,
			NoSandwiching: false,
		},
		// Aestus
		{
			ID:          config.MevRelayID_Aestus,
			Name:        "Aestus",
			Description: "The Aestus MEV-Boost Relay is an independent and non-censoring relay. It is committed to neutrality and the development of a healthy MEV-Boost ecosystem.",
			Urls: map[config.Network]string{
				config.Network_Mainnet: "https://0xa15b52576bcbf1072f4a011c0f99f9fb6c66f3e1ff321f11f461d15e31b1cb359caa092c71bbded0bae5b5ea401aab7e@aestus.live?id=staderlabs",
				config.Network_Holesky: "https://0xab78bf8c781c58078c3beb5710c57940874dd96aef2835e7742c866b4c7c0406754376c2c8285a36c630346aa5c5f833@holesky.aestus.live?id=staderlabs",
			},
			Regulated:     false,
			NoSandwiching: false,
		},
		// Agnostic
		{
			ID:          config.MevRelayID_Agnostic,
			Name:        "Agnostic",
			Description: "Agnostic Relay is an open-source MEV Boost relay available to anyone, anywhere in the world, without prejudice or privilege. It is an ideal relay for block producers and block builders trying to provide neutral features.",
			Urls: map[config.Network]string{
				config.Network_Mainnet: "https://0xa7ab7a996c8584251c8f925da3170bdfd6ebc75d50f5ddc4050a6fdc77f2a3b5fce2cc750d0865e05d7228af97d69561@agnostic-relay.net?id=staderlabs",
			},
			Regulated:     true,
			NoSandwiching: false,
		},

		// Titan Regulated
		{
			ID:          config.MevRelayID_TitanRegulated,
			Name:        "Titan Regulated (filtering)",
			Description: "Titan Relay is a neutral, Rust-based MEV-Boost Relay optimized for low latency throughput, geographical distribution, and robustness. Select this to enable the \"filtering\" relay from Titan.",
			Urls: map[config.Network]string{
				config.Network_Mainnet: "https://0x8c4ed5e24fe5c6ae21018437bde147693f68cda427cd1122cf20819c30eda7ed74f72dece09bb313f2a1855595ab677d@regional.titanrelay.xyz",
				config.Network_Holesky: "https://0xaa58208899c6105603b74396734a6263cc7d947f444f396a90f7b7d3e65d102aec7e5e5291b27e08d02c50a050825c2f@holesky.titanrelay.xyz",
			},
			Regulated:     true,
			NoSandwiching: false,
		},

		// Titan UnRegulated
		{
			ID:          config.MevRelayID_TitanUnRegulated,
			Name:        "Titan UnRegulated (non-filtering)",
			Description: "Titan Relay is a neutral, Rust-based MEV-Boost Relay optimized for low latency throughput, geographical distribution, and robustness. Select this to enable the \"non-filtering\" relay from Titan.",
			Urls: map[config.Network]string{
				config.Network_Mainnet: "https://0x8c4ed5e24fe5c6ae21018437bde147693f68cda427cd1122cf20819c30eda7ed74f72dece09bb313f2a1855595ab677d@global.titanrelay.xyz",
			},
			Regulated:     false,
			NoSandwiching: false,
		},
	}

	return relays
}

// Generate one of the profile parameters
func generateProfileParameter(id string, relays []config.MevRelay, regulated bool, noSandwiching bool) config.Parameter {
	name := "Enable "
	description := fmt.Sprintf("[lime]NOTE: You can enable multiple options.")

	if regulated {
		name += "Regulated "
		description += RegulatedRelayDescription
	} else {
		name += "Unregulated "
		description += UnregulatedRelayDescription
	}

	if noSandwiching {
		name += "(No Sandwiching)"
		description += NoSandwichRelayDescription
	} else {
		name += "(All MEV Types)"
		description += AllMevRelayDescription
	}

	// Generate the Mainnet description
	mainnetRelays := []string{}
	mainnetDescription := description + "\n\nRelays: "
	for _, relay := range relays {
		_, exists := relay.Urls[config.Network_Mainnet]
		if !exists {
			continue
		}
		if relay.Regulated == regulated && relay.NoSandwiching == noSandwiching {
			mainnetRelays = append(mainnetRelays, relay.Name)
		}
	}
	mainnetDescription += strings.Join(mainnetRelays, ", ")

	// Generate the holesky description
	holeskyRelays := []string{}
	holeskyDescription := description + "\n\nRelays:\n"
	for _, relay := range relays {
		_, exists := relay.Urls[config.Network_Holesky]
		if !exists {
			continue
		}
		if relay.Regulated == regulated && relay.NoSandwiching == noSandwiching {
			holeskyRelays = append(holeskyRelays, relay.Name)
		}
	}

	holeskyDescription += strings.Join(holeskyRelays, ", ")

	return config.Parameter{
		ID:                   id,
		Name:                 name,
		Description:          mainnetDescription,
		Type:                 config.ParameterType_Bool,
		Default:              map[config.Network]interface{}{config.Network_All: false},
		AffectsContainers:    []config.ContainerID{config.ContainerID_MevBoost},
		EnvironmentVariables: []string{},
		CanBeBlank:           false,
		OverwriteOnUpgrade:   false,
		DescriptionsByNetwork: map[config.Network]string{
			config.Network_Mainnet: mainnetDescription,
			config.Network_Holesky: holeskyDescription,
		},
	}
}

// Generate one of the relay parameters
func generateRelayParameter(id string, relay config.MevRelay) config.Parameter {
	description := fmt.Sprintf("[lime]NOTE: You can enable multiple options.")

	if relay.Regulated {
		description += "Complies with Regulations: YES\n"
	} else {
		description += "Complies with Regulations: NO\n"
	}

	if relay.NoSandwiching {
		description += "Allows Sandwich Attacks: NO"
	} else {
		description += "Allows Sandwich Attacks: YES"
	}

	return config.Parameter{
		ID:                   id,
		Name:                 fmt.Sprintf("Enable %s", relay.Name),
		Description:          description,
		Type:                 config.ParameterType_Bool,
		Default:              map[config.Network]interface{}{config.Network_All: false},
		AffectsContainers:    []config.ContainerID{config.ContainerID_MevBoost},
		EnvironmentVariables: []string{},
		CanBeBlank:           false,
		OverwriteOnUpgrade:   false,
	}
}
