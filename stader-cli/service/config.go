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
package service

import (
	"fmt"
	"strings"

	"github.com/davecgh/go-spew/spew"
	ethCfUI "github.com/stader-labs/ethcli-configuration-ui"
	"github.com/stader-labs/ethcli-configuration-ui/config"
	"github.com/stader-labs/ethcli-configuration-ui/logger"
	stdCf "github.com/stader-labs/stader-node/shared/services/config"
	"github.com/stader-labs/stader-node/shared/services/stader"
	cfgtypes "github.com/stader-labs/stader-node/shared/types/config"
	"github.com/urfave/cli"
)

var (
	log = logger.Log
)

func format(v interface{}) string {
	return fmt.Sprintf("%v", v)
}

var keys = config.GetFieldKey()

func makeConfigFromUISetting(cfg *stdCf.StaderConfig, settings map[string]interface{}) stdCf.StaderConfig {
	newCfg := *cfg

	spew.Dump(settings)

	// update the network
	network := settings[keys.Sn_node_network].(string)
	if network == "Goerli Testnet" {
		newCfg.ChangeNetwork(cfgtypes.Network("prater"))
	} else {
		newCfg.ChangeNetwork(cfgtypes.Network("mainnet"))
	}

	// Stader node config
	staderNode := newCfg.StaderNode
	staderNode.ProjectName.Value = settings[keys.Sn_project_title]
	staderNode.DataPath.Value = settings[keys.Sn_storage_location]

	updateFeeAndReward(&newCfg, settings)
	updateExecutionClient(&newCfg, settings)

	// update the consensus client
	newCfg.ConsensusClientMode.Value = makeCfgExecutionMode(settings[keys.E1ec_execution_client_mode])
	return newCfg
}

func makeUISettingFromConfig(cfg *stdCf.StaderConfig) map[string]interface{} {
	settings := make(map[string]interface{})

	setUIStaderNode(cfg, settings)

	settings[keys.E1ec_execution_client_mode] = makeUIExecutionMode(cfg.ExecutionClientMode.Value)
	settings[keys.E2cc_consensus_client] = makeUIExecutionMode(cfg.ConsensusClientMode.Value)

	settings[keys.E1ec_em_websocket_url] = cfg.ExternalExecution.WsUrl.Value
	settings[keys.E1ec_em_http_url] = cfg.ExternalExecution.HttpUrl.Value

	settings[keys.E2cc_em_consensus_client] = strings.Title(format(cfg.ExternalConsensusClient.Value))

	setUIConsensusClient(cfg, settings)
	return settings
}

func setUIStaderNode(cfg *stdCf.StaderConfig, settings map[string]interface{}) error {
	staderNode := cfg.StaderNode

	switch staderNode.Network.Value.(cfgtypes.Network) {
	case cfgtypes.Network_Prater:
		settings[keys.Sn_node_network] = "Goerli Testnet"
	case cfgtypes.Network_Mainnet:
		settings[keys.Sn_node_network] = "Ethereum Mainnet"
	}

	settings[keys.Sn_project_title] = staderNode.ProjectName.Value
	settings[keys.Sn_storage_location] = staderNode.DataPath.Value

	settings[keys.Fr_max_fee_limit] = format(staderNode.ManualMaxFee.Value)
	settings[keys.Fr_priority_fee] =
		format(staderNode.PriorityFee.Value)
	settings[keys.Fr_archive_mode_ec_url] = staderNode.ArchiveECUrl.Value
	return nil
}

func setUIConsensusClient(cfg *stdCf.StaderConfig, newSettings map[string]interface{}) error {

	// case cfgtypes.ConsensusClient_Teku:
	newSettings[keys.E2cc_em_custom_graffiti_teku] = cfg.ExternalTeku.Graffiti.Value
	newSettings[keys.E2cc_em_http_url_teku] = cfg.ExternalTeku.HttpUrl.Value
	newSettings[keys.E2cc_em_container_tag_teku] = cfg.ExternalTeku.ContainerTag.Value
	newSettings[keys.E2cc_em_additional_client_flags_teku] = cfg.ExternalTeku.AdditionalVcFlags.Value
	// case cfgtypes.ConsensusClient_Lighthouse:
	newSettings[keys.E2cc_em_custom_graffiti_lighthouse] = cfg.ExternalLighthouse.Graffiti.Value
	newSettings[keys.E2cc_em_http_url_lighthouse] = cfg.ExternalLighthouse.HttpUrl.Value
	newSettings[keys.E2cc_em_container_tag_lighthouse] = cfg.ExternalLighthouse.ContainerTag.Value
	newSettings[keys.E2cc_em_additional_client_flags_lighthouse] = cfg.ExternalLighthouse.AdditionalVcFlags.Value
	newSettings[keys.E2cc_em_doppelganger_detection_lighthouse] = cfg.ExternalLighthouse.DoppelgangerDetection.Value.(bool)

	// case cfgtypes.ConsensusClient_Prysm:
	newSettings[keys.E2cc_em_custom_graffiti_prysm] = cfg.ExternalPrysm.Graffiti.Value
	newSettings[keys.E2cc_em_http_url_prysm] = cfg.ExternalPrysm.HttpUrl.Value
	newSettings[keys.E2cc_em_container_tag_prysm] = cfg.ExternalPrysm.ContainerTag.Value
	newSettings[keys.E2cc_em_additional_client_flags_prysm] = cfg.ExternalPrysm.AdditionalVcFlags.Value
	newSettings[keys.E2cc_em_doppelganger_detection_prysm] = cfg.ExternalPrysm.DoppelgangerDetection.Value.(bool)

	return nil
}

func updateFeeAndReward(newCfg *stdCf.StaderConfig, settings map[string]interface{}) {
	staderNode := newCfg.StaderNode
	staderNode.ManualMaxFee.Value = settings[keys.Fr_max_fee_limit]
	staderNode.PriorityFee.Value = settings[keys.Fr_priority_fee]
	staderNode.ArchiveECUrl.Value = settings[keys.Fr_archive_mode_ec_url]
}

func configureNode(c *cli.Context) error {
	staderClient, err := stader.NewClientFromCtx(c)
	if err != nil {
		return err
	}
	defer staderClient.Close()

	cfg, _, err := staderClient.LoadConfig()
	if err != nil {
		return fmt.Errorf("error loading user settings: %w", err)
	}

	oldSetting := makeUISettingFromConfig(cfg)
	_, _, m := ethCfUI.Run(oldSetting)

	fmt.Printf("APP DONE \n")
	newConfig := makeConfigFromUISetting(cfg, m)
	err = staderClient.SaveConfig(&newConfig)
	if err != nil {
		return err
	}

	return nil
}
