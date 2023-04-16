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
	stdCf "github.com/stader-labs/stader-node/shared/services/config"
)

func setUIFallbackClient(cfg *stdCf.StaderConfig, newSettings map[string]interface{}) error {
	newSettings[keys.Fc_use_fallback_clients] = cfg.UseFallbackClients.Value.(bool)
	newSettings[keys.Fc_true_reconnect_delay] = cfg.ReconnectDelay.Value
	newSettings[keys.Fc_true_execution_client_url] = cfg.FallbackNormal.EcHttpUrl.Value
	newSettings[keys.Fc_true_beacon_node_url] = cfg.FallbackNormal.CcHttpUrl.Value
	newSettings[keys.Fc_true_beacon_node_json_rpc_url] = cfg.FallbackPrysm.JsonRpcUrl.Value

	return nil
}

func updateFallbackClient(cfg *stdCf.StaderConfig, newSettings map[string]interface{}) error {
	cfg.UseFallbackClients.Value = newSettings[keys.Fc_use_fallback_clients]

	if cfg.UseFallbackClients.Value.(bool) {
		cfg.ReconnectDelay.Value = newSettings[keys.Fc_true_reconnect_delay]
		cfg.FallbackNormal.EcHttpUrl.Value = newSettings[keys.Fc_true_execution_client_url]
		cfg.FallbackNormal.CcHttpUrl.Value = newSettings[keys.Fc_true_beacon_node_url]

		cfg.FallbackPrysm.EcHttpUrl.Value = newSettings[keys.Fc_true_execution_client_url]
		cfg.FallbackPrysm.CcHttpUrl.Value = newSettings[keys.Fc_true_beacon_node_url]
		cfg.FallbackPrysm.JsonRpcUrl.Value = newSettings[keys.Fc_true_beacon_node_json_rpc_url]
	}
	return nil
}
