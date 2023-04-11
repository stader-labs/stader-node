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

func makeConfigFromUISetting(oldCfg *stdCf.StaderConfig, settings map[string]interface{}) stdCf.StaderConfig {
	newCfg := *oldCfg
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
	updateConsensusClient(&newCfg, settings)
	updateFallbackClient(&newCfg, settings)
	updateMEVBoost(&newCfg, settings)
	updateMonitoring(&newCfg, settings)

	return newCfg
}

func makeUISettingFromConfig(cfg *stdCf.StaderConfig) map[string]interface{} {
	settings := make(map[string]interface{})

	setUIStaderNode(cfg, settings)
	setUIExecutionClient(cfg, settings)
	setUIConsensusClient(cfg, settings)
	setUIFallbackClient(cfg, settings)
	setUIMEVBoost(cfg, settings)
	setUIMonitoring(cfg, settings)
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

func updateFeeAndReward(newCfg *stdCf.StaderConfig, settings map[string]interface{}) {
	staderNode := newCfg.StaderNode
	staderNode.ManualMaxFee.Value = settings[keys.Fr_max_fee_limit]
	staderNode.PriorityFee.Value = settings[keys.Fr_priority_fee]
	staderNode.ArchiveECUrl.Value = settings[keys.Fr_archive_mode_ec_url]
}

func openConfigurationSetting(c *cli.Context) (bool, error) {
	staderClient, err := stader.NewClientFromCtx(c)
	if err != nil {
		return false, err
	}
	defer staderClient.Close()

	cfg, _, err := staderClient.LoadConfig()
	if err != nil {
		return false, fmt.Errorf("error loading user settings: %w", err)
	}

	oldSetting := makeUISettingFromConfig(cfg)
	saved, openConfig, m := ethCfUI.Run(oldSetting)

	if saved {
		fmt.Printf("Your settings have not changed.\n")
		newConfig := makeConfigFromUISetting(cfg, m)
		err = staderClient.SaveConfig(&newConfig)
		if err != nil {
			return false, err
		}
	}

	if err != nil {
		return false, err
	}

	return openConfig, nil
}
