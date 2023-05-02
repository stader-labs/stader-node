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
	"strings"

	stdCf "github.com/stader-labs/stader-node/shared/services/config"
	cfgtypes "github.com/stader-labs/stader-node/shared/types/config"
)

func updateExecutionClient(cfg *stdCf.StaderConfig, newSettings map[string]interface{}) error {
	// update the execution client
	executionMod := makeCfgExecutionMode(newSettings[keys.E1ec_execution_and_consensus_mode])
	cfg.ExecutionClientMode.Value = executionMod

	if err := updateExternalExecutionClient(cfg, newSettings); err != nil {
		return err
	}
	if err := updateLocalExecutionClient(cfg, newSettings); err != nil {
		return err
	}

	return nil
}

func updateExternalExecutionClient(cfg *stdCf.StaderConfig, newSettings map[string]interface{}) error {
	cfg.ExternalExecution.WsUrl.Value = newSettings[keys.E1ec_em_websocket_url]
	cfg.ExternalExecution.HttpUrl.Value = newSettings[keys.E1ec_em_http_url]
	return nil
}

func updateLocalExecutionClient(cfg *stdCf.StaderConfig, newSettings map[string]interface{}) error {
	cfg.ExecutionClient.Value = cfgtypes.ExecutionClient(strings.ToLower(newSettings[keys.E1ec_lm_execution_client].(string)))

	cfg.ExecutionCommon.HttpPort.Value = newSettings[keys.E1ec_lm_common_http_port]
	cfg.ExecutionCommon.WsPort.Value = newSettings[keys.E1ec_lm_common_websocket_port]
	cfg.ExecutionCommon.P2pPort.Value = newSettings[keys.E1ec_lm_common_p2p_port]
	cfg.ExecutionCommon.OpenRpcPorts.Value = newSettings[keys.E1ec_lm_common_expose_rpc_port]
	cfg.ExecutionCommon.EnginePort.Value = newSettings[keys.E1ec_lm_common_engine_api_port]
	cfg.ExecutionCommon.EthstatsLabel.Value = newSettings[keys.E1ec_lm_common_ethstats_label]
	cfg.ExecutionCommon.EthstatsLogin.Value = newSettings[keys.E1ec_lm_common_ethstats_login]

	cfg.Geth.CacheSize.Value = newSettings[keys.E1ec_lm_geth_cache_size]
	cfg.Geth.MaxPeers.Value = newSettings[keys.E1ec_lm_geth_max_peers]
	cfg.Geth.ContainerTag.Value = newSettings[keys.E1ec_lm_geth_container_tag]
	cfg.Geth.AdditionalFlags.Value = newSettings[keys.E1ec_lm_geth_additional_flags]

	// Nethermind
	cfg.Nethermind.CacheSize.Value = newSettings[keys.E1ec_lm_nethermind_cache_size]
	cfg.Nethermind.MaxPeers.Value = newSettings[keys.E1ec_lm_nethermind_max_peers]
	cfg.Nethermind.ContainerTag.Value = newSettings[keys.E1ec_lm_nethermind_container_tag]
	cfg.Nethermind.AdditionalFlags.Value = newSettings[keys.E1ec_lm_nethermind_additional_flags]
	cfg.Nethermind.PruneMemSize.Value = newSettings[keys.E1ec_lm_nethermind_pruning_cache_size]
	cfg.Nethermind.AdditionalModules.Value = newSettings[keys.E1ec_lm_nethermind_additional_modules]
	cfg.Nethermind.AdditionalUrls.Value = newSettings[keys.E1ec_lm_nethermind_additional_urls]

	// Besu
	cfg.Besu.JvmHeapSize.Value = newSettings[keys.E1ec_lm_besu_jvm_heap_size]
	cfg.Besu.MaxPeers.Value = newSettings[keys.E1ec_lm_besu_max_peers]
	cfg.Besu.ContainerTag.Value = newSettings[keys.E1ec_lm_besu_container_tag]
	cfg.Besu.AdditionalFlags.Value = newSettings[keys.E1ec_lm_besu_additional_flags]
	cfg.Besu.MaxBackLayers.Value = newSettings[keys.E1ec_lm_besu_historical_block_replay_limit]

	return nil
}

func setUIExecutionClient(cfg *stdCf.StaderConfig, newSettings map[string]interface{}) error {
	newSettings[keys.E1ec_execution_and_consensus_mode] = makeUIExecutionMode(cfg.ExecutionClientMode.Value)

	newSettings[keys.E1ec_em_websocket_url] = cfg.ExternalExecution.WsUrl.Value
	newSettings[keys.E1ec_em_http_url] = cfg.ExternalExecution.HttpUrl.Value

	newSettings[keys.E1ec_lm_execution_client] = strings.Title(string((cfg.ExecutionClient.Value.(cfgtypes.ExecutionClient))))

	// Common
	newSettings[keys.E1ec_lm_common_http_port] = format(cfg.ExecutionCommon.HttpPort.Value)
	newSettings[keys.E1ec_lm_common_websocket_port] = format(cfg.ExecutionCommon.WsPort.Value)
	newSettings[keys.E1ec_lm_common_expose_rpc_port] = cfg.ExecutionCommon.OpenRpcPorts.Value
	newSettings[keys.E1ec_lm_common_p2p_port] = format(cfg.ExecutionCommon.P2pPort.Value)
	newSettings[keys.E1ec_lm_common_engine_api_port] = format(cfg.ExecutionCommon.EnginePort.Value)
	newSettings[keys.E1ec_lm_common_ethstats_label] = cfg.ExecutionCommon.EthstatsLabel.Value
	newSettings[keys.E1ec_lm_common_ethstats_login] = cfg.ExecutionCommon.EthstatsLogin.Value

	// Geth
	newSettings[keys.E1ec_lm_geth_cache_size] = format(cfg.Geth.CacheSize.Value)
	newSettings[keys.E1ec_lm_geth_max_peers] = format(cfg.Geth.MaxPeers.Value)
	newSettings[keys.E1ec_lm_geth_container_tag] = cfg.Geth.ContainerTag.Value
	newSettings[keys.E1ec_lm_geth_additional_flags] = cfg.Geth.AdditionalFlags.Value

	// Nethermind
	newSettings[keys.E1ec_lm_nethermind_cache_size] = format(cfg.Nethermind.CacheSize.Value)
	newSettings[keys.E1ec_lm_nethermind_max_peers] = format(cfg.Nethermind.MaxPeers.Value)
	newSettings[keys.E1ec_lm_nethermind_container_tag] = cfg.Nethermind.ContainerTag.Value
	newSettings[keys.E1ec_lm_nethermind_additional_flags] = cfg.Nethermind.AdditionalFlags.Value
	newSettings[keys.E1ec_lm_nethermind_pruning_cache_size] = format(cfg.Nethermind.PruneMemSize.Value)
	newSettings[keys.E1ec_lm_nethermind_additional_modules] = cfg.Nethermind.AdditionalModules.Value
	newSettings[keys.E1ec_lm_nethermind_additional_urls] = cfg.Nethermind.AdditionalUrls.Value

	// Besu
	newSettings[keys.E1ec_lm_besu_jvm_heap_size] = format(cfg.Besu.JvmHeapSize.Value)
	newSettings[keys.E1ec_lm_besu_max_peers] = format(cfg.Besu.MaxPeers.Value)
	newSettings[keys.E1ec_lm_besu_container_tag] = cfg.Besu.ContainerTag.Value
	newSettings[keys.E1ec_lm_besu_additional_flags] = cfg.Besu.AdditionalFlags.Value
	newSettings[keys.E1ec_lm_besu_historical_block_replay_limit] = cfg.Besu.MaxBackLayers.Value

	return nil
}

func makeCfgExecutionMode(i interface{}) cfgtypes.Mode {
	mod := format(i)
	switch mod {
	case "Locally Managed":
		return cfgtypes.Mode_Local
	case "Externally Managed":
		return cfgtypes.Mode_External
	default:
		return cfgtypes.Mode_Unknown
	}
}

func makeUIExecutionMode(i interface{}) string {
	mod := i.(cfgtypes.Mode)
	switch mod {
	case cfgtypes.Mode_Local:
		return "Locally Managed"
	case cfgtypes.Mode_External:
		return "Externally Managed"
	default:
		return ""
	}
}
