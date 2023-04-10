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
)

func setUIConsensusClient(cfg *stdCf.StaderConfig, newSettings map[string]interface{}) error {
	newSettings[keys.E2cc_consensus_client] = makeUIExecutionMode(cfg.ConsensusClientMode.Value)
	newSettings[keys.E2cc_em_consensus_client] = strings.Title(format(cfg.ExternalConsensusClient.Value))
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

func updateConsensusClient(newCfg *stdCf.StaderConfig, settings map[string]interface{}) error {
	newCfg.ExternalConsensusClient.Value = settings[keys.E2cc_em_consensus_client]
	newCfg.ConsensusClientMode.Value = makeCfgExecutionMode(settings[keys.E1ec_execution_client_mode])
	return nil
}
