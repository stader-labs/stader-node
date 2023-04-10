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
	stdCf "github.com/stader-labs/stader-node/shared/services/config"
	cfgtypes "github.com/stader-labs/stader-node/shared/types/config"
)

func updateExecutionClient(cfg *stdCf.StaderConfig, newSettings map[string]interface{}) error {
	// update the execution client
	cfg.ExecutionClientMode.Value = makeCfgExecutionMode(newSettings[keys.E1ec_execution_client_mode])
	switch newSettings[keys.E1ec_execution_client_mode] {
	case "Externally Managed":
		updateExternalExecutionClient(cfg, newSettings)

	case "Locally Managed":
	}

	return nil
}

func updateExternalExecutionClient(cfg *stdCf.StaderConfig, newSettings map[string]interface{}) error {
	cfg.ExternalExecution.WsUrl.Value = newSettings[keys.E1ec_em_websocket_url]
	cfg.ExternalExecution.HttpUrl.Value = newSettings[keys.E1ec_em_http_url]

	cfg.ExternalConsensusClient.Value = newSettings[keys.E2cc_em_consensus_client]

	spew.Dump(" newSettings[keys.E2cc_consensus_client] ")
	clientStr, ok := newSettings[keys.E2cc_em_consensus_client].(string)
	if !ok {
		return fmt.Errorf("Invalid External client %+v", newSettings[keys.E2cc_em_consensus_client])
	}

	// case cfgtypes.ConsensusClient_Teku:
	cfg.ExternalTeku.Graffiti.Value = newSettings[keys.E2cc_em_custom_graffiti_teku]
	cfg.ExternalTeku.HttpUrl.Value = newSettings[keys.E2cc_em_http_url_teku]
	cfg.ExternalTeku.ContainerTag.Value = newSettings[keys.E2cc_em_container_tag_teku]
	cfg.ExternalTeku.AdditionalVcFlags.Value = newSettings[keys.E2cc_em_additional_client_flags_teku]

	// case cfgtypes.ConsensusClient_Lighthouse:
	cfg.ExternalLighthouse.Graffiti.Value = newSettings[keys.E2cc_em_custom_graffiti_lighthouse]
	cfg.ExternalLighthouse.HttpUrl.Value = newSettings[keys.E2cc_em_http_url_lighthouse]
	cfg.ExternalLighthouse.ContainerTag.Value = newSettings[keys.E2cc_em_container_tag_lighthouse]
	cfg.ExternalLighthouse.AdditionalVcFlags.Value = newSettings[keys.E2cc_em_additional_client_flags_lighthouse]
	cfg.ExternalLighthouse.DoppelgangerDetection.Value = newSettings[keys.E2cc_em_doppelganger_detection_lighthouse]

	// case cfgtypes.ConsensusClient_Prysm:
	cfg.ExternalPrysm.Graffiti.Value = newSettings[keys.E2cc_em_custom_graffiti_prysm]
	cfg.ExternalPrysm.HttpUrl.Value = newSettings[keys.E2cc_em_http_url_prysm]
	cfg.ExternalPrysm.ContainerTag.Value = newSettings[keys.E2cc_em_container_tag_prysm]
	cfg.ExternalPrysm.AdditionalVcFlags.Value = newSettings[keys.E2cc_em_additional_client_flags_prysm]
	cfg.ExternalPrysm.DoppelgangerDetection.Value = newSettings[keys.E2cc_em_doppelganger_detection_prysm]

	cfg.ConsensusClient.Value = strings.ToLower(clientStr)
	cfg.ExternalExecution.WsUrl.Value = newSettings[keys.E1ec_em_websocket_url]
	cfg.ExternalExecution.HttpUrl.Value = newSettings[keys.E1ec_em_http_url]

	return nil
}
func makeCfgExecutionMode(i interface{}) string {
	mod, ok := i.(string)
	if !ok {
		return ""
	}
	switch mod {
	case "Locally Managed":
		return "local"
	case "Externally Managed":
		return "external"
	default:
		return ""
	}
}

func makeUIExecutionMode(i interface{}) string {
	mod, ok := i.(cfgtypes.Mode)
	if !ok {
		return ""
	}
	switch mod {
	case cfgtypes.Mode_Local:
		return "Locally Managed"
	case cfgtypes.Mode_External:
		return "Externally Managed"
	default:
		return ""
	}
}
