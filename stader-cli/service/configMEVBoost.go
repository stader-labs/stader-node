/*
This work is licensed and released under GNU GPL v3 or any other later versions.
The full text of the license is below/ found at <http://www.gnu.org/licenses/>

(c) 2023 Rocket Pool Pty Ltd. Modified under GNU GPL v3. [1.4.1]

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
	cfgtypes "github.com/stader-labs/stader-node/shared/types/config"
)

func setUIMEVBoost(cfg *stdCf.StaderConfig, newSettings map[string]interface{}) error {
	newSettings[keys.Mev_boost_enabled] = cfg.EnableMevBoost.Value.(bool)
	newSettings[keys.Mev_boost_mode] = makeUIMEVMode(cfg.MevBoost.Mode.Value.(cfgtypes.Mode))
	newSettings[keys.Mev_boost_em_external_url] = cfg.MevBoost.ExternalUrl.Value

	mevSelection := cfg.MevBoost.SelectionMode.Value.(cfgtypes.MevSelectionMode)
	newSettings[keys.Mev_boost_lm_selection_mode] = makeUISelectionMode(mevSelection)

	newSettings[keys.Mev_boost_pm_enable_unregulated] = cfg.MevBoost.EnableUnregulatedAllMev.Value.(bool)
	newSettings[keys.Mev_boost_pm_enable_regulated] = cfg.MevBoost.EnableRegulatedAllMev.Value.(bool)
	newSettings[keys.Mev_boost_pm_port] = format(cfg.MevBoost.Port.Value)
	newSettings[keys.Mev_boost_pm_expose_api_port] = cfg.MevBoost.OpenRpcPort.Value.(bool)
	newSettings[keys.Mev_boost_pm_container_tag] = cfg.MevBoost.ContainerTag.Value
	newSettings[keys.Mev_boost_pm_additional_flags] = cfg.MevBoost.AdditionalFlags.Value

	// Relay Mode
	newSettings[keys.Mev_boost_rm_enable_flashbots] = cfg.MevBoost.FlashbotsRelay.Value.(bool)
	newSettings[keys.Mev_boost_rm_enable_bloXroute_max_profit] = cfg.MevBoost.BloxRouteMaxProfitRelay.Value.(bool)
	newSettings[keys.Mev_boost_rm_enable_eden_network] = cfg.MevBoost.EdenRelay.Value.(bool)
	newSettings[keys.Mev_boost_rm_enable_ultra_sound] = cfg.MevBoost.UltrasoundRelay.Value.(bool)

	// case cfgtypes.MevSelectionMode_Profile:
	newSettings[keys.Mev_boost_pm_port] = format(cfg.MevBoost.Port.Value)
	newSettings[keys.Mev_boost_pm_expose_api_port] = cfg.MevBoost.OpenRpcPort.Value.(bool)
	newSettings[keys.Mev_boost_pm_container_tag] = cfg.MevBoost.ContainerTag.Value
	newSettings[keys.Mev_boost_pm_additional_flags] = cfg.MevBoost.AdditionalFlags.Value
	// case cfgtypes.MevSelectionMode_Relay:
	newSettings[keys.Mev_boost_rm_port] = format(cfg.MevBoost.Port.Value)
	newSettings[keys.Mev_boost_rm_expose_api_port] = cfg.MevBoost.OpenRpcPort.Value.(bool)
	newSettings[keys.Mev_boost_rm_container_tag] = cfg.MevBoost.ContainerTag.Value
	newSettings[keys.Mev_boost_rm_additional_flags] = cfg.MevBoost.AdditionalFlags.Value

	v, _ := cfg.MevBoost.AestusRelay.Value.(bool)
	newSettings[keys.Mev_boost_rm_enable_aestus] = v

	newSettings[keys.Mev_boost_rm_enable_bloXroute_regulated] = cfg.MevBoost.BloxRouteRegulatedRelay.Value.(bool)

	return nil
}

func updateMEVBoost(cfg *stdCf.StaderConfig, newSettings map[string]interface{}) error {
	// Make MEV boot default is true since we remove key from UI
	cfg.EnableMevBoost.Value = true

	cfg.MevBoost.Mode.Value = makeMEVModeFromUi(newSettings[keys.Mev_boost_mode].(string))
	cfg.MevBoost.ExternalUrl.Value = newSettings[keys.Mev_boost_em_external_url]

	mevSelection := makeMEVSelectionModeFromUI(newSettings[keys.Mev_boost_lm_selection_mode].(string))
	cfg.MevBoost.SelectionMode.Value = mevSelection

	// Profile node
	cfg.MevBoost.EnableUnregulatedAllMev.Value = newSettings[keys.Mev_boost_pm_enable_unregulated]
	cfg.MevBoost.EnableRegulatedAllMev.Value = newSettings[keys.Mev_boost_pm_enable_regulated]

	// Relay Mode
	cfg.MevBoost.FlashbotsRelay.Value = newSettings[keys.Mev_boost_rm_enable_flashbots]
	cfg.MevBoost.BloxRouteMaxProfitRelay.Value = newSettings[keys.Mev_boost_rm_enable_bloXroute_max_profit]

	cfg.MevBoost.EdenRelay.Value = newSettings[keys.Mev_boost_rm_enable_eden_network]
	cfg.MevBoost.UltrasoundRelay.Value = newSettings[keys.Mev_boost_rm_enable_ultra_sound]

	cfg.MevBoost.BloxRouteRegulatedRelay.Value = newSettings[keys.Mev_boost_rm_enable_bloXroute_regulated]
	cfg.MevBoost.AestusRelay.Value = newSettings[keys.Mev_boost_rm_enable_aestus]

	switch mevSelection {
	case cfgtypes.MevSelectionMode_Profile:
		cfg.MevBoost.Port.Value = newSettings[keys.Mev_boost_pm_port]
		cfg.MevBoost.OpenRpcPort.Value = newSettings[keys.Mev_boost_pm_expose_api_port]
		cfg.MevBoost.ContainerTag.Value = newSettings[keys.Mev_boost_pm_container_tag]
		cfg.MevBoost.AdditionalFlags.Value = newSettings[keys.Mev_boost_pm_additional_flags]
	case cfgtypes.MevSelectionMode_Relay:
		cfg.MevBoost.Port.Value = newSettings[keys.Mev_boost_rm_port]
		cfg.MevBoost.OpenRpcPort.Value = newSettings[keys.Mev_boost_rm_expose_api_port]
		cfg.MevBoost.ContainerTag.Value = newSettings[keys.Mev_boost_rm_container_tag]
		cfg.MevBoost.AdditionalFlags.Value = newSettings[keys.Mev_boost_rm_additional_flags]
	}

	return nil
}

func makeUIMEVMode(mode cfgtypes.Mode) string {
	switch mode {
	case cfgtypes.Mode_Local:
		return "Locally Managed"
	case cfgtypes.Mode_External:
		return "Externally Managed"
	}

	return ""
}

func makeMEVModeFromUi(mode string) cfgtypes.Mode {
	switch mode {
	case "Locally Managed":
		return cfgtypes.Mode_Local
	case "Externally Managed":
		return cfgtypes.Mode_External
	}

	return cfgtypes.Mode_Unknown
}
func makeUISelectionMode(mode cfgtypes.MevSelectionMode) string {
	switch mode {
	case cfgtypes.MevSelectionMode_Profile:
		return "Standard Mode"
	case cfgtypes.MevSelectionMode_Relay:
		return "Custom Mode"
	}

	return ""
}
func makeMEVSelectionModeFromUI(mode string) cfgtypes.MevSelectionMode {
	switch mode {
	case "Standard Mode":
		return cfgtypes.MevSelectionMode_Profile
	case "Custom Mode":
		return cfgtypes.MevSelectionMode_Relay
	}

	return ""
}
