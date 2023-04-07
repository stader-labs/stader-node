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

	"github.com/davecgh/go-spew/spew"
	ethcliconfigurationui "github.com/stader-labs/ethcli-configuration-ui"
	"github.com/stader-labs/ethcli-configuration-ui/config"
	"github.com/stader-labs/ethcli-configuration-ui/logger"
	stdCf "github.com/stader-labs/stader-node/shared/services/config"
	"github.com/stader-labs/stader-node/shared/services/stader"
	"github.com/urfave/cli"
)

var (
	log = logger.Log
)

func format(v interface{}) string {
	return fmt.Sprintf("%v", v)
}

func makeSetting(cfg *stdCf.StaderConfig) map[string]interface{} {
	v := make(map[string]interface{})
	f := config.GetFieldKey()

	// The StaderNode configuration
	staderNode := cfg.StaderNode
	fmt.Printf("LOAD %+v", staderNode.Network.Value)
	v[f.Sn_node_network] = "Goerli Testnet"
	v[f.Sn_project_title] = staderNode.Title
	v[f.Sn_storage_location] = staderNode.DataPath.Value

	return v
}

func makeConfig(cfg *stdCf.StaderConfig, setting map[string]interface{}) stdCf.StaderConfig {
	newCfg := *cfg
	f := config.GetFieldKey()

	spew.Dump(setting)
	staderNode := newCfg.StaderNode
	staderNode.Network.Value = setting[f.Sn_node_network]
	staderNode.ProjectName.Value = setting[f.Sn_project_title].(string)
	// staderNode.DataPath.Value = setting[f.Sn_storage_location].(string)

	return newCfg
}

func configureNode(c *cli.Context) error {

	staderClient, err := stader.NewClientFromCtx(c)
	if err != nil {
		return err
	}
	defer staderClient.Close()

	// Load the config, checking to see if it's new (hasn't been installed before)
	// var oldCfg *config.StaderConfig
	cfg, _, err := staderClient.LoadConfig()
	if err != nil {
		return fmt.Errorf("error loading user settings: %w", err)
	}

	_, _, m := ethcliconfigurationui.Run(makeSetting(cfg))

	fmt.Printf("APP DONE \n")
	newConfig := makeConfig(cfg, m)
	err = staderClient.SaveConfig(&newConfig)
	if err != nil {
		return err
	}

	return nil
}
