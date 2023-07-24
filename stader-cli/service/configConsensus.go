/*
This work is licensed and released under GNU GPL v3 or any other later versions.
The full text of the license is below/ found at <http://www.gnu.org/licenses/>

(c) 2023 Rocket Pool Pty Ltd. Modified under GNU GPL v3. [1.2.1]

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

	stdCf "github.com/stader-labs/stader-node/shared/services/config"
	cfgtypes "github.com/stader-labs/stader-node/shared/types/config"
)

func setUIConsensusClient(cfg *stdCf.StaderConfig, newSettings map[string]interface{}) error {
	// newSettings[keys.E1ec_execution_and_consensus_mode] = makeUIExecutionMode(cfg.ConsensusClientMode.Value)

	newSettings[keys.E2cc_lc_consensus_client] = strings.Title(format(cfg.ConsensusClient.Value))
	newSettings[keys.E2cc_em_consensus_client] = strings.Title(format(cfg.ExternalConsensusClient.Value))

	// External
	// Teku:
	newSettings[keys.E2cc_em_custom_graffiti_teku] = cfg.ExternalTeku.Graffiti.Value
	newSettings[keys.E2cc_em_http_teku] = cfg.ExternalTeku.HttpUrl.Value
	newSettings[keys.E2cc_em_container_tag_teku] = cfg.ExternalTeku.ContainerTag.Value
	newSettings[keys.E2cc_em_additional_client_flags_teku] = cfg.ExternalTeku.AdditionalVcFlags.Value
	newSettings[keys.E2cc_lc_enable_archive_mode_teku] = cfg.Teku.ArchiveMode.Value.(bool)

	// Lighthouse:
	newSettings[keys.E2cc_em_custom_graffiti_lighthouse] = cfg.ExternalLighthouse.Graffiti.Value
	newSettings[keys.E2cc_em_http_lighthouse] = cfg.ExternalLighthouse.HttpUrl.Value
	newSettings[keys.E2cc_em_container_tag_lighthouse] = cfg.ExternalLighthouse.ContainerTag.Value
	newSettings[keys.E2cc_em_additional_client_flags_lighthouse] = cfg.ExternalLighthouse.AdditionalVcFlags.Value
	newSettings[keys.E2cc_em_doppelganger_detection_lighthouse] = cfg.ExternalLighthouse.DoppelgangerDetection.Value.(bool)

	// Nimbus:
	newSettings[keys.E2cc_em_custom_graffiti_nimbus] = cfg.ExternalNimbus.Graffiti.Value
	newSettings[keys.E2cc_em_http_nimbus] = cfg.ExternalNimbus.HttpUrl.Value
	newSettings[keys.E2cc_em_container_tag_nimbus] = cfg.ExternalNimbus.ContainerTag.Value
	newSettings[keys.E2cc_em_additional_client_flags_nimbus] = cfg.ExternalNimbus.AdditionalVcFlags.Value
	newSettings[keys.E2cc_em_doppelganger_detection_nimbus] = cfg.ExternalNimbus.DoppelgangerDetection.Value.(bool)

	// Prysm:
	newSettings[keys.E2cc_em_custom_graffiti_prysm] = cfg.ExternalPrysm.Graffiti.Value
	newSettings[keys.E2cc_em_http_prysm] = cfg.ExternalPrysm.HttpUrl.Value
	newSettings[keys.E2cc_em_container_tag_prysm] = cfg.ExternalPrysm.ContainerTag.Value
	newSettings[keys.E2cc_em_additional_client_flags_prysm] = cfg.ExternalPrysm.AdditionalVcFlags.Value
	newSettings[keys.E2cc_em_doppelganger_detection_prysm] = cfg.ExternalPrysm.DoppelgangerDetection.Value.(bool)
	newSettings[keys.E2cc_em_json_rpc_prysm] = cfg.ExternalPrysm.JsonRpcUrl.Value

	// Local
	newSettings[keys.E2cc_lc_common_graffiti] = cfg.ConsensusCommon.Graffiti.Value
	newSettings[keys.E2cc_lc_common_checkpoint_sync_url] = cfg.ConsensusCommon.CheckpointSyncProvider.Value
	newSettings[keys.E2cc_lc_common_p2p_port] = format(cfg.ConsensusCommon.P2pPort.Value)
	newSettings[keys.E2cc_lc_common_http_api_port] = format(cfg.ConsensusCommon.ApiPort.Value)
	newSettings[keys.E2cc_lc_common_expose_api_port] = cfg.ConsensusCommon.OpenApiPort.Value.(bool)
	newSettings[keys.E2cc_lc_common_doppelganger_detection] = cfg.ConsensusCommon.DoppelgangerDetection.Value.(bool)

	// Max peers
	newSettings[keys.E2cc_lc_max_peer_lighthouse] = format(cfg.Lighthouse.MaxPeers.Value)
	newSettings[keys.E2cc_lc_max_peer_nimbus] = format(cfg.Nimbus.MaxPeers.Value)
	newSettings[keys.E2cc_lc_max_peer_prysm] = format(cfg.Prysm.MaxPeers.Value)
	newSettings[keys.E2cc_lc_max_peer_teku] = format(cfg.Teku.MaxPeers.Value)

	// Lighthouse
	newSettings[keys.E2cc_lc_container_tag_lighthouse] = cfg.Lighthouse.ContainerTag.Value
	newSettings[keys.E2cc_lc_additional_beacon_node_flags_lighthouse] = cfg.Lighthouse.AdditionalBnFlags.Value
	newSettings[keys.E2cc_lc_additional_client_flags_lighthouse] = cfg.Lighthouse.AdditionalVcFlags.Value

	// Teku
	newSettings[keys.E2cc_lc_jvm_heap_size] = format(cfg.Teku.JvmHeapSize.Value)
	newSettings[keys.E2cc_lc_container_tag_teku] = cfg.Teku.ContainerTag.Value
	newSettings[keys.E2cc_lc_additional_beacon_node_flags_teku] = cfg.Teku.AdditionalBnFlags.Value
	newSettings[keys.E2cc_lc_additional_client_flags_teku] = cfg.Teku.AdditionalVcFlags.Value

	// Nimbus
	newSettings[keys.E2cc_lc_container_tag_nimbus] = cfg.Nimbus.BnContainerTag.Value
	newSettings[keys.E2cc_lc_additional_flags_nimbus] = cfg.Nimbus.AdditionalBnFlags.Value

	// Prysm
	newSettings[keys.E2cc_lc_rpc_port_prysm] = format(cfg.Prysm.RpcPort.Value)
	newSettings[keys.E2cc_lc_expose_rpc_port_prysm] = cfg.Prysm.OpenRpcPort.Value

	newSettings[keys.E2cc_lc_beacon_container_tag_prysm] = cfg.Prysm.BnContainerTag.Value
	newSettings[keys.E2cc_lc_validator_client_container_tag_prysm] = cfg.Prysm.VcContainerTag.Value
	newSettings[keys.E2cc_lc_additional_beacon_node_flags_prysm] = cfg.Prysm.AdditionalBnFlags.Value
	newSettings[keys.E2cc_lc_additional_client_flags_prysm] = cfg.Prysm.AdditionalVcFlags.Value

	return nil
}

func updateConsensusClient(newCfg *stdCf.StaderConfig, settings map[string]interface{}) error {

	if err := updateExternalConsensusClient(newCfg, settings); err != nil {
		return err
	}

	if err := updateLocalConsensusClient(newCfg, settings); err != nil {
		return err
	}

	return nil
}

func updateExternalConsensusClient(cfg *stdCf.StaderConfig, newSettings map[string]interface{}) error {
	cfg.ExternalConsensusClient.Value = cfgtypes.ConsensusClient(strings.ToLower(newSettings[keys.E2cc_em_consensus_client].(string)))
	cfg.ConsensusClientMode.Value = makeCfgExecutionMode(newSettings[keys.E1ec_execution_and_consensus_mode])

	// case cfgtypes.ConsensusClient_Teku:
	cfg.ExternalTeku.Graffiti.Value = newSettings[keys.E2cc_em_custom_graffiti_teku]
	cfg.ExternalTeku.HttpUrl.Value = newSettings[keys.E2cc_em_http_teku]
	cfg.ExternalTeku.ContainerTag.Value = newSettings[keys.E2cc_em_container_tag_teku]
	cfg.ExternalTeku.AdditionalVcFlags.Value = newSettings[keys.E2cc_em_additional_client_flags_teku]

	// case cfgtypes.ConsensusClient_Lighthouse:
	cfg.ExternalLighthouse.Graffiti.Value = newSettings[keys.E2cc_em_custom_graffiti_lighthouse]
	cfg.ExternalLighthouse.HttpUrl.Value = newSettings[keys.E2cc_em_http_lighthouse]
	cfg.ExternalLighthouse.ContainerTag.Value = newSettings[keys.E2cc_em_container_tag_lighthouse]
	cfg.ExternalLighthouse.AdditionalVcFlags.Value = newSettings[keys.E2cc_em_additional_client_flags_lighthouse]
	cfg.ExternalLighthouse.DoppelgangerDetection.Value = newSettings[keys.E2cc_em_doppelganger_detection_lighthouse]

	// case cfgtypes.ConsensusClient_Nimbus:
	cfg.ExternalNimbus.Graffiti.Value = newSettings[keys.E2cc_em_custom_graffiti_nimbus]
	cfg.ExternalNimbus.HttpUrl.Value = newSettings[keys.E2cc_em_http_nimbus]
	cfg.ExternalNimbus.ContainerTag.Value = newSettings[keys.E2cc_em_container_tag_nimbus]
	cfg.ExternalNimbus.AdditionalVcFlags.Value = newSettings[keys.E2cc_em_additional_client_flags_nimbus]
	cfg.ExternalNimbus.DoppelgangerDetection.Value = newSettings[keys.E2cc_em_doppelganger_detection_nimbus]

	// case cfgtypes.ConsensusClient_Prysm:
	cfg.ExternalPrysm.Graffiti.Value = newSettings[keys.E2cc_em_custom_graffiti_prysm]
	cfg.ExternalPrysm.HttpUrl.Value = newSettings[keys.E2cc_em_http_prysm]
	cfg.ExternalPrysm.ContainerTag.Value = newSettings[keys.E2cc_em_container_tag_prysm]
	cfg.ExternalPrysm.AdditionalVcFlags.Value = newSettings[keys.E2cc_em_additional_client_flags_prysm]
	cfg.ExternalPrysm.DoppelgangerDetection.Value = newSettings[keys.E2cc_em_doppelganger_detection_prysm]
	cfg.ExternalPrysm.JsonRpcUrl.Value = newSettings[keys.E2cc_em_json_rpc_prysm]

	return nil
}

func updateLocalConsensusClient(newCfg *stdCf.StaderConfig, settings map[string]interface{}) error {
	// Common
	newCfg.ConsensusCommon.Graffiti.Value = settings[keys.E2cc_lc_common_graffiti]
	newCfg.ConsensusCommon.CheckpointSyncProvider.Value = settings[keys.E2cc_lc_common_checkpoint_sync_url]
	newCfg.ConsensusCommon.P2pPort.Value = settings[keys.E2cc_lc_common_p2p_port]
	newCfg.ConsensusCommon.ApiPort.Value = settings[keys.E2cc_lc_common_http_api_port]
	newCfg.ConsensusCommon.OpenApiPort.Value = settings[keys.E2cc_lc_common_expose_api_port]
	newCfg.ConsensusCommon.DoppelgangerDetection.Value = settings[keys.E2cc_lc_common_doppelganger_detection]

	clientStr, ok := settings[keys.E2cc_lc_consensus_client].(string)
	if !ok {
		return fmt.Errorf("Invalid External client %+v", settings[keys.E1ec_lm_execution_client])
	}

	newCfg.ConsensusClient.Value = cfgtypes.ConsensusClient(strings.ToLower(clientStr))

	newCfg.Lighthouse.MaxPeers.Value = settings[keys.E2cc_lc_max_peer_lighthouse]
	newCfg.Lighthouse.ContainerTag.Value = settings[keys.E2cc_lc_container_tag_lighthouse]
	newCfg.Lighthouse.AdditionalBnFlags.Value = settings[keys.E2cc_lc_additional_beacon_node_flags_lighthouse]
	newCfg.Lighthouse.AdditionalVcFlags.Value = settings[keys.E2cc_lc_additional_client_flags_lighthouse]

	// Common Max Peer
	newCfg.Lighthouse.MaxPeers.Value = settings[keys.E2cc_lc_max_peer_lighthouse]
	newCfg.Nimbus.MaxPeers.Value = settings[keys.E2cc_lc_max_peer_nimbus]
	newCfg.Prysm.MaxPeers.Value = settings[keys.E2cc_lc_max_peer_prysm]
	newCfg.Teku.MaxPeers.Value = settings[keys.E2cc_lc_max_peer_teku]

	// Lighthouse
	newCfg.Lighthouse.ContainerTag.Value = settings[keys.E2cc_lc_container_tag_lighthouse]
	newCfg.Lighthouse.AdditionalBnFlags.Value = settings[keys.E2cc_lc_additional_beacon_node_flags_lighthouse]
	newCfg.Lighthouse.AdditionalVcFlags.Value = settings[keys.E2cc_lc_additional_client_flags_lighthouse]

	// Teku
	newCfg.Teku.JvmHeapSize.Value = settings[keys.E2cc_lc_jvm_heap_size]
	newCfg.Teku.ArchiveMode.Value = settings[keys.E2cc_lc_enable_archive_mode_teku]
	newCfg.Teku.ContainerTag.Value = settings[keys.E2cc_lc_container_tag_teku]
	newCfg.Teku.AdditionalBnFlags.Value = settings[keys.E2cc_lc_additional_beacon_node_flags_teku]
	newCfg.Teku.AdditionalVcFlags.Value = settings[keys.E2cc_lc_additional_client_flags_teku]

	newCfg.Nimbus.BnContainerTag.Value = settings[keys.E2cc_lc_container_tag_nimbus]
	newCfg.Nimbus.AdditionalBnFlags.Value = settings[keys.E2cc_lc_additional_flags_nimbus]

	// Prysm
	newCfg.Prysm.RpcPort.Value = settings[keys.E2cc_lc_rpc_port_prysm]
	newCfg.Prysm.OpenRpcPort.Value = settings[keys.E2cc_lc_expose_rpc_port_prysm]
	newCfg.Prysm.BnContainerTag.Value = settings[keys.E2cc_lc_beacon_container_tag_prysm]
	newCfg.Prysm.VcContainerTag.Value = settings[keys.E2cc_lc_validator_client_container_tag_prysm]
	newCfg.Prysm.AdditionalBnFlags.Value = settings[keys.E2cc_lc_additional_beacon_node_flags_prysm]
	newCfg.Prysm.AdditionalVcFlags.Value = settings[keys.E2cc_lc_additional_client_flags_prysm]

	return nil
}
