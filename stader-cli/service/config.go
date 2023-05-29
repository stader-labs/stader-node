/*
This work is licensed and released under GNU GPL v3 or any other later versions.
The full text of the license is below/ found at <http://www.gnu.org/licenses/>

(c) 2023 Rocket Pool Pty Ltd. Modified under GNU GPL v3. [0.4.0-beta]

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

	"github.com/stader-labs/ethcli-ui/configuration"
	"github.com/stader-labs/ethcli-ui/configuration/config"
	"github.com/stader-labs/ethcli-ui/configuration/logger"
	"github.com/stader-labs/ethcli-ui/wizard"
	stdCf "github.com/stader-labs/stader-node/shared/services/config"
	"github.com/stader-labs/stader-node/shared/services/stader"
	cfgtypes "github.com/stader-labs/stader-node/shared/types/config"
	"github.com/urfave/cli"
)

var (
	log = logger.Log
)

type ConfigContext struct {
	staderClient    *stader.Client
	openWizardPage  bool
	newConfigFromUI stdCf.StaderConfig
	cliCtx          *cli.Context
}

func format(v interface{}) string {
	return fmt.Sprintf("%v", v)
}

var keys = config.GetFieldKey()

func updateConfigFromUISetting(oldCfg *stdCf.StaderConfig, settings map[string]interface{}) (*stdCf.StaderConfig, error) {
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
	if err := updateExecutionClient(&newCfg, settings); err != nil {
		return nil, fmt.Errorf("Error updateExecutionClient %+v", err)
	}
	if err := updateConsensusClient(&newCfg, settings); err != nil {
		return nil, fmt.Errorf("Error updateConsensusClient %+v", err)
	}
	if err := updateFallbackClient(&newCfg, settings); err != nil {
		return nil, fmt.Errorf("Error updateFallbackClient %+v", err)
	}
	if err := updateMEVBoost(&newCfg, settings); err != nil {
		return nil, fmt.Errorf("Error updateMEVBoost %+v", err)
	}
	if err := updateMonitoring(&newCfg, settings); err != nil {
		return nil, fmt.Errorf("Error updateMonitoring %+v", err)
	}

	return &newCfg, nil
}

func makeUISettingFromConfig(cfg *stdCf.StaderConfig) (map[string]interface{}, error) {
	settings := make(map[string]interface{})

	if err := setUIStaderNode(cfg, settings); err != nil {
		return nil, fmt.Errorf("Error setUIStaderNode %+v ", err)
	}
	if err := setUIExecutionClient(cfg, settings); err != nil {
		return nil, fmt.Errorf("Error setUIExecutionClient %+v ", err)
	}
	if err := setUIConsensusClient(cfg, settings); err != nil {
		return nil, fmt.Errorf("Error setUIConsensusClient %+v ", err)
	}
	if err := setUIFallbackClient(cfg, settings); err != nil {
		return nil, fmt.Errorf("Error setUIFallbackClient %+v ", err)
	}
	if err := setUIMEVBoost(cfg, settings); err != nil {
		return nil, fmt.Errorf("Error setUIMonitoring %+v ", err)
	}
	if err := setUIMonitoring(cfg, settings); err != nil {
		return nil, fmt.Errorf("Error setUIMEVBoost %+v ", err)
	}
	return settings, nil
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

func configureService(c *cli.Context) error {
	staderClient, err := stader.NewClientFromCtx(c)
	if err != nil {
		return err
	}
	defer staderClient.Close()

	landingUI := "wizard"
	cfg, err := loadConfig(c)
	if err != nil {
		return fmt.Errorf("error loading user settings: %w", err)
	}

	saved, newCg, err := handleUI(landingUI, cfg)

	if err != nil {
		fmt.Printf("Error from parsing UI to config model %+v", err)
		return err
	}

	if !saved {
		fmt.Printf("Your settings have not changed.\n")
		return nil
	}

	err = staderClient.SaveConfig(newCg)
	if err != nil {
		return fmt.Errorf("error SaveConfig: %w", err)
	}

	// Restart the services
	err = startService(c, true)
	if err != nil {
		return fmt.Errorf("error startService: %w", err)
	}

	// Remove the upgrade flag if it's there
	return staderClient.RemoveUpgradeFlagFile()
}

func handleUI(
	initialUI string,
	cfg *stdCf.StaderConfig,
) (bool, *stdCf.StaderConfig, error) {

	if initialUI == "wizard" {
		oldWSettings := NewSettingsType(cfg)
		confirmed, openConfiguring, newWSettings := wizard.Run(&oldWSettings)

		cfg, err := UpdateConfig(cfg, newWSettings)

		if err != nil {
			return false, nil, err
		}
		if openConfiguring {
			return handleUI("configuration", &cfg)
		}

		return confirmed, &cfg, nil
	}

	if initialUI == "configuration" {
		oldCSetting, err := makeUISettingFromConfig(cfg)

		if err != nil {
			return false, nil, err
		}
		cSaved, cOpenWizard, newCSettings := configuration.Run(&oldCSetting)
		cfg, err := updateConfigFromUISetting(cfg, *newCSettings)

		if err != nil {
			return false, nil, err
		}

		if cOpenWizard {
			return handleUI("wizard", cfg)
		}

		return cSaved, cfg, nil
	}

	return false, cfg, nil
}
