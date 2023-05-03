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
	"os"
	"path/filepath"
	"reflect"
	"regexp"
	"strings"
	"time"

	"github.com/mitchellh/go-homedir"
	"github.com/urfave/cli"
	"gopkg.in/yaml.v2"

	"github.com/dustin/go-humanize"
	"github.com/shirou/gopsutil/v3/disk"
	"github.com/stader-labs/stader-node/shared"
	"github.com/stader-labs/stader-node/shared/services/config"

	"github.com/stader-labs/ethcli-ui/wizard/pages"
	"github.com/stader-labs/stader-node/shared/services/stader"
	cfgtypes "github.com/stader-labs/stader-node/shared/types/config"
	cliutils "github.com/stader-labs/stader-node/shared/utils/cli"
	"github.com/stader-labs/stader-node/shared/utils/sys"
)

// Settings
const (
	ExporterContainerSuffix         string = "_exporter"
	ValidatorContainerSuffix        string = "_validator"
	BeaconContainerSuffix           string = "_eth2"
	ExecutionContainerSuffix        string = "_eth1"
	NodeContainerSuffix             string = "_node"
	ApiContainerSuffix              string = "_api"
	GuardianContainerSuffix         string = "guardian"
	PruneProvisionerContainerSuffix string = "_prune_provisioner"
	EcMigratorContainerSuffix       string = "_ec_migrator"
	clientDataVolumeName            string = "/ethclient"
	dataFolderVolumeName            string = "/.stader/data"

	PruneFreeSpaceRequired uint64 = 50 * 1024 * 1024 * 1024
	dockerImageRegex       string = ".*/(?P<image>.*):.*"
	colorReset             string = "\033[0m"
	colorBold              string = "\033[1m"
	colorRed               string = "\033[31m"
	colorYellow            string = "\033[33m"
	colorGreen             string = "\033[32m"
	colorLightBlue         string = "\033[36m"
	clearLine              string = "\033[2K"
)

// Install the Stader service
func installService(c *cli.Context) error {
	dataPath := ""

	if c.String("network") != "" {
		fmt.Printf("%sNOTE: The --network flag is deprecated. You no longer need to specify it.%s\n\n", colorLightBlue, colorReset)
	}

	// Prompt for confirmation
	if !(c.Bool("yes") || cliutils.Confirm(fmt.Sprintf(
		"The Stader service will be installed --Version: %s\n\n%sIf you're upgrading, your existing configuration will be backed up and preserved.\nAll of your previous settings will be migrated automatically.%s\nAre you sure you want to continue?",
		c.String("version"), colorGreen, colorReset,
	))) {
		fmt.Println("Cancelled.")
		return nil
	}

	staderClient, err := stader.NewClientFromCtx(c)
	if err != nil {
		return err
	}
	defer staderClient.Close()

	// Attempt to load the config to see if any settings need to be passed along to the install script
	cfg, isNew, err := staderClient.LoadConfig()
	if err != nil {
		return fmt.Errorf("error loading old configuration: %w", err)
	}
	if !isNew {
		dataPath = cfg.StaderNode.DataPath.Value.(string)
		dataPath, err = homedir.Expand(dataPath)
		if err != nil {
			return fmt.Errorf("error getting data path from old configuration: %w", err)
		}
	}
	fmt.Println("Installing Stader service...")

	// Install service
	err = staderClient.InstallService(c.Bool("verbose"), c.Bool("no-deps"), c.String("network"), c.String("version"), c.String("path"), dataPath)
	if err != nil {
		return err
	}

	// Print success message & return
	fmt.Println("")
	fmt.Println("The Stader service was successfully installed!")

	// printPatchNotes(c)

	// Reload the config after installation
	_, isNew, err = staderClient.LoadConfig()
	if err != nil {
		return fmt.Errorf("error loading new configuration: %w", err)
	}

	// Check if this is a migration
	isMigration := false
	if isNew {
		// Look for a legacy config to migrate
		migratedConfig, err := staderClient.LoadLegacyConfigFromBackup()
		if err != nil {
			return err
		}
		if migratedConfig != nil {
			isMigration = true
		}
	}

	// Report next steps
	fmt.Printf("%s\n=== Next Steps ===\n", colorLightBlue)
	fmt.Printf("Run 'stader-cli service config' to review the settings changes for this update, or to continue setting up your node.%s\n", colorReset)

	// Print the docker permissions notice
	if isNew && !isMigration {
		fmt.Printf("\n%sNOTE:\nSince this is your first time installing Stader, please start a new shell session by logging out and back in or restarting the machine.\n", colorYellow)
		fmt.Printf("This is necessary for your user account to have permissions to use Docker.%s", colorReset)
	}

	return nil

}

// Print the latest patch notes for this release
func printPatchNotes(c *cli.Context) {

	fmt.Printf(`%s
   _____ _            _             _           _         
  / ____| |          | |           | |         | |        
 | (___ | |_ __ _  __| | ___ _ __  | |     __ _| |__  ___ 
  \___ \| __/ _' |/ _' |/ _ \ '__| | |    / _' | '_ \/ __|
  ____) | || (_| | (_| |  __/ |    | |___| (_| | |_) \__ \
 |_____/ \__\__,_|\__,_|\___|_|    |______\__,_|_.__/|___/

`, "\033[34m")
	fmt.Printf("%s=== Stadernode v%s ===%s\n\n", colorGreen, shared.StaderVersion, colorReset)
	fmt.Printf("Changes you should be aware of before starting:\n\n")

	fmt.Printf("%s=== MEV-Boost Updates ===%s\n", colorGreen, colorReset)
	fmt.Println("MEV-Boost is now opt-out instead of opt-in. Furthermore, there is a new way to select relays: you can now select \"profiles\" instead of individual relays. As new relays are added to the Stadernode, any that belong to the profiles you've selected will automatically be enabled for you.\nNOTE: everyone will have to configure either profile-mode or individual-relay mode when first upgrading from v1.6, even if you had previously configured MEV-Boost.")

	fmt.Printf("%s=== ENS Support ===%s\n", colorGreen, colorReset)
	fmt.Println("`stader-cli node set-withdrawal-address`, `stader-cli node send`, and `stader-cli node set-voting-delegate` can now use ENS names instead of addresses! This requires your Execution Client to be online and synced.\nAlso, use the `stader-cli wallet set-ens-name` command to confirm an ENS domain or subdomain name that you assign to your node wallet. Once you do this, you can refer to your node's address by its ENS name on explorers like Etherscan.")

	fmt.Printf("%s=== Modern vs. Portable ===%s\n", colorGreen, colorReset)
	fmt.Println("The Stadernode now automatically checks your node's CPU features and defaults to either the \"modern\" optimized version of certain clients, or the more generic \"portable\" version based on what your machine supports. This only applies to MEV-Boost and Lighthouse.")

}

// View the Stader service status
func serviceStatus(c *cli.Context) error {

	staderClient, err := stader.NewClientFromCtx(c)
	if err != nil {
		return err
	}
	defer staderClient.Close()

	// Print what network we're on
	err = cliutils.PrintNetwork(staderClient)
	if err != nil {
		return err
	}

	// Print service status
	return staderClient.PrintServiceStatus(getComposeFiles(c))

}

func ConvertBoolToString(boolValue bool) string {
	if boolValue {
		return "Yes"
	}

	return "No"
}

func ConvertStringToBool(stringValue string) bool {
	if stringValue == "Yes" {
		return true
	} else {
		return false
	}
}

func HasBeenUpdated(settings1 pages.SettingsType, settings2 pages.SettingsType) bool {
	if reflect.DeepEqual(settings1, settings2) {
		return false
	}

	return true
}

func UpdateConfig(_cfg *config.StaderConfig, newSettings *pages.SettingsType) (config.StaderConfig, error) {
	cfg := *_cfg
	// update the network
	cfg.ChangeNetwork(cfgtypes.Network(newSettings.Network))

	// update the consensus and execution client
	cfg.ConsensusClientMode.Value = cfgtypes.Mode(newSettings.EthClient)
	cfg.ExecutionClientMode.Value = cfgtypes.Mode(newSettings.EthClient)

	if newSettings.EthClient == "local" {
		cfg.ExecutionClient.Value = cfgtypes.ExecutionClient(strings.ToLower(newSettings.ExecutionClient.SelectionOption))
		cfg.ConsensusClient.Value = cfgtypes.ConsensusClient(newSettings.ConsensusClient.Selection)
	} else if newSettings.EthClient == "external" {
		cfg.ExternalConsensusClient.Value = newSettings.ConsensusClient.ExternalSelection
		cfg.ExternalExecution.WsUrl.Value = newSettings.ExecutionClient.External.WebsocketBasedRpcApi
		cfg.ExternalExecution.HttpUrl.Value = newSettings.ExecutionClient.External.HTTPBasedRpcApi
		cfg.ExternalPrysm.DoppelgangerDetection.Value = ConvertStringToBool(newSettings.ConsensusClient.DoppelgangerProtection)
		cfg.ExternalPrysm.HttpUrl.Value = newSettings.ConsensusClient.External.Prysm.HTTPUrl
		cfg.ExternalPrysm.JsonRpcUrl.Value = newSettings.ConsensusClient.External.Prysm.JSONRpcUrl
		cfg.ExternalLighthouse.DoppelgangerDetection.Value = ConvertStringToBool(newSettings.ConsensusClient.DoppelgangerProtection)
		cfg.ExternalLighthouse.HttpUrl.Value = newSettings.ConsensusClient.External.Lighthouse.HTTPUrl
		cfg.ExternalTeku.Graffiti.Value = newSettings.ConsensusClient.Graffit
		cfg.ExternalTeku.HttpUrl.Value = newSettings.ConsensusClient.External.Teku.HTTPUrl
	}
	cfg.ConsensusCommon.DoppelgangerDetection.Value = ConvertStringToBool(newSettings.ConsensusClient.DoppelgangerProtection)
	cfg.ConsensusCommon.Graffiti.Value = newSettings.ConsensusClient.Graffit
	cfg.ConsensusCommon.CheckpointSyncProvider.Value = newSettings.ConsensusClient.CheckpointUrl

	cfg.UseFallbackClients.Value = ConvertStringToBool(newSettings.FallbackClients.SelectionOption)
	// get the consensus client we are using for fallback
	fallBackConsensusClient := newSettings.ConsensusClient.Selection
	if newSettings.EthClient == "external" {
		fallBackConsensusClient = newSettings.ConsensusClient.ExternalSelection
	}

	switch fallBackConsensusClient {
	case "prysm":
		cfg.FallbackPrysm.EcHttpUrl.Value = newSettings.FallbackClients.Prysm.ExecutionClientUrl
		cfg.FallbackPrysm.CcHttpUrl.Value = newSettings.FallbackClients.Prysm.BeaconNodeHttpUrl
		cfg.FallbackPrysm.JsonRpcUrl.Value = newSettings.FallbackClients.Prysm.BeaconNodeJsonRpcpUrl
	case "lighthouse":
		cfg.FallbackNormal.EcHttpUrl.Value = newSettings.FallbackClients.Lighthouse.ExecutionClientUrl
		cfg.FallbackNormal.CcHttpUrl.Value = newSettings.FallbackClients.Lighthouse.BeaconNodeHttpUrl
	case "teku":
		cfg.FallbackNormal.EcHttpUrl.Value = newSettings.FallbackClients.Teku.ExecutionClientUrl
		cfg.FallbackNormal.CcHttpUrl.Value = newSettings.FallbackClients.Teku.BeaconNodeHttpUrl
	}

	// update monitoring
	cfg.EnableMetrics.Value = ConvertStringToBool(newSettings.Monitoring)

	// update mev boost
	cfg.EnableMevBoost.Value = true
	if newSettings.MEVBoost == "external" {
		cfg.MevBoost.ExternalUrl.Value = newSettings.MEVBoostExternalMevUrl
		cfg.MevBoost.Mode.Value = cfgtypes.Mode_External
	} else if newSettings.MEVBoost == "local" {
		cfg.MevBoost.Mode.Value = cfgtypes.Mode_Local
		cfg.MevBoost.EnableUnregulatedAllMev.Value = newSettings.MEVBoostLocalUnregulated
		cfg.MevBoost.EnableRegulatedAllMev.Value = newSettings.MEVBoostLocalRegulated
	}

	// unset mev boost mode value if mev boost is disabled
	if newSettings.MEVBoost == "local" && cfg.EnableMevBoost.Value.(bool) == false {
		cfg.MevBoost.Mode.Value = cfgtypes.Mode_Local
		cfg.MevBoost.SelectionMode.Value = cfgtypes.MevSelectionMode_Unknow
	}

	return cfg, nil
}

func NewSettingsType(cfg *config.StaderConfig) pages.SettingsType {
	currentSettings := pages.SettingsType{
		Network:   "prater",
		EthClient: string(cfg.ConsensusClientMode.Value.(cfgtypes.Mode)),
		ExecutionClient: pages.ExecutionClientSettingsType{
			SelectionOption: string(cfg.ExecutionClient.Value.(cfgtypes.ExecutionClient)),
			External: pages.ExecutionClientExternalType{
				HTTPBasedRpcApi:      cfg.ExternalExecution.HttpUrl.Value.(string),
				WebsocketBasedRpcApi: cfg.ExternalExecution.WsUrl.Value.(string),
			},
		},
		ConsensusClient: pages.ConsensusClientSettingsType{
			Selection:              string(cfg.ConsensusClient.Value.(cfgtypes.ConsensusClient)),
			ExternalSelection:      string(cfg.ExternalConsensusClient.Value.(cfgtypes.ConsensusClient)),
			Graffit:                cfg.ConsensusCommon.Graffiti.Value.(string),
			CheckpointUrl:          cfg.ConsensusCommon.CheckpointSyncProvider.Value.(string),
			DoppelgangerProtection: ConvertBoolToString(cfg.ConsensusCommon.DoppelgangerDetection.Value.(bool)),
			External: pages.ConsensusClientExternalType{
				Lighthouse: pages.ConsensusClientExternalSelectedLighthouseType{
					HTTPUrl: cfg.ExternalLighthouse.HttpUrl.Value.(string),
				},
				Prysm: pages.ConsensusClientExternalSelectedPrysmType{
					HTTPUrl:    cfg.ExternalPrysm.HttpUrl.Value.(string),
					JSONRpcUrl: cfg.ExternalPrysm.JsonRpcUrl.Value.(string),
				},
				Teku: pages.ConsensusClientExternalSelectedTekuType{
					HTTPUrl: cfg.ExternalTeku.HttpUrl.Value.(string),
				},
			},
		},
		Monitoring:               ConvertBoolToString(cfg.EnableMetrics.Value.(bool)),
		MEVBoost:                 string(cfg.MevBoost.Mode.Value.(cfgtypes.Mode)),
		MEVBoostExternalMevUrl:   cfg.MevBoost.ExternalUrl.Value.(string),
		MEVBoostLocalRegulated:   cfg.MevBoost.EnableRegulatedAllMev.Value.(bool),
		MEVBoostLocalUnregulated: cfg.MevBoost.EnableUnregulatedAllMev.Value.(bool),
		FallbackClients: pages.FallbackClientsSettingsType{
			SelectionOption: ConvertBoolToString(cfg.UseFallbackClients.Value.(bool)),
			Lighthouse: pages.FallbackClientsLighthouseType{
				ExecutionClientUrl: cfg.FallbackNormal.EcHttpUrl.Value.(string),
				BeaconNodeHttpUrl:  cfg.FallbackNormal.CcHttpUrl.Value.(string),
			},
			Prysm: pages.FallbackClientsPrysmType{
				ExecutionClientUrl:    cfg.FallbackPrysm.EcHttpUrl.Value.(string),
				BeaconNodeHttpUrl:     cfg.FallbackPrysm.CcHttpUrl.Value.(string),
				BeaconNodeJsonRpcpUrl: cfg.FallbackPrysm.JsonRpcUrl.Value.(string),
			},
			Teku: pages.FallbackClientsTekuType{
				ExecutionClientUrl: cfg.FallbackNormal.EcHttpUrl.Value.(string),
				BeaconNodeHttpUrl:  cfg.FallbackNormal.CcHttpUrl.Value.(string),
			},
		},
	}

	return currentSettings
}

// Configure the service
func loadConfig(c *cli.Context) (*config.StaderConfig, error) {
	// Make sure the config directory exists first
	configPath := c.GlobalString("config-path")
	fmt.Println("configPath", configPath)
	path, err := homedir.Expand(configPath)
	if err != nil {
		return nil, fmt.Errorf("error expanding config path [%s]: %w", configPath, err)
	}
	staderClient, err := stader.NewClientFromCtx(c)
	if err != nil {
		return nil, err
	}
	defer staderClient.Close()

	_, err = os.Stat(path)
	if os.IsNotExist(err) {
		fmt.Printf("%sYour configured Stader config directory of [%s] does not exist.\n%s\n", colorYellow, path, colorReset)
		if err != nil {
			return nil, fmt.Errorf("error installService: %w", err)
		}
	}

	// Load the config, checking to see if it's new (hasn't been installed before)
	// var oldCfg *config.StaderConfig
	cfg, isNew, err := staderClient.LoadConfig()
	if err != nil {
		return nil, fmt.Errorf("error loading user settings: %w", err)
	}

	// Check to see if this is a migration from a legacy config
	isMigration := false
	if isNew {
		// Look for a legacy config to migrate
		migratedConfig, err := staderClient.LoadLegacyConfigFromBackup()
		if err != nil {
			return nil, err
		}
		if migratedConfig != nil {
			cfg = migratedConfig
			isMigration = true
		}
	}

	// Check if this is a new install
	isUpdate, err := staderClient.IsFirstRun()
	if err != nil {
		return nil, fmt.Errorf("error checking for first-run status: %w", err)
	}

	// For migrations and upgrades, move the config to the old one and create a new upgraded copy
	if isMigration || isUpdate {
		// oldCfg = cfg
		cfg = cfg.CreateCopy()
		err = cfg.UpdateDefaults()
		if err != nil {
			return nil, fmt.Errorf("error upgrading configuration with the latest parameters: %w", err)
		}
	}

	// Save the config and exit in headless mode
	if c.NumFlags() > 0 {
		err := configureHeadless(c, cfg)
		if err != nil {
			return nil, fmt.Errorf("error updating config from provided arguments: %w", err)
		}
		return nil, staderClient.SaveConfig(cfg)
	}
	return cfg, nil
}

// Updates a configuration from the provided CLI arguments headlessly
func configureHeadless(c *cli.Context, cfg *config.StaderConfig) error {

	// Root params
	for _, param := range cfg.GetParameters() {
		err := updateConfigParamFromCliArg(c, "", param, cfg)
		if err != nil {
			return err
		}
	}

	// Subconfigs
	for sectionName, subconfig := range cfg.GetSubconfigs() {
		for _, param := range subconfig.GetParameters() {
			err := updateConfigParamFromCliArg(c, sectionName, param, cfg)
			if err != nil {
				return err
			}
		}
	}

	return nil

}

// Updates a config parameter from a CLI flag
func updateConfigParamFromCliArg(c *cli.Context, sectionName string, param *cfgtypes.Parameter, cfg *config.StaderConfig) error {

	var paramName string
	if sectionName == "" {
		paramName = param.ID
	} else {
		paramName = fmt.Sprintf("%s-%s", sectionName, param.ID)
	}

	if c.IsSet(paramName) {
		switch param.Type {
		case cfgtypes.ParameterType_Bool:
			param.Value = c.Bool(paramName)
		case cfgtypes.ParameterType_Int:
			param.Value = c.Int(paramName)
		case cfgtypes.ParameterType_Float:
			param.Value = c.Float64(paramName)
		case cfgtypes.ParameterType_String:
			setting := c.String(paramName)
			if param.MaxLength > 0 && len(setting) > param.MaxLength {
				return fmt.Errorf("error setting value for %s: [%s] is too long (max length %d)", paramName, setting, param.MaxLength)
			}
			param.Value = c.String(paramName)
		case cfgtypes.ParameterType_Uint:
			param.Value = c.Uint(paramName)
		case cfgtypes.ParameterType_Uint16:
			param.Value = uint16(c.Uint(paramName))
		case cfgtypes.ParameterType_Choice:
			selection := c.String(paramName)
			found := false
			for _, option := range param.Options {
				if fmt.Sprint(option.Value) == selection {
					param.Value = option.Value
					found = true
					break
				}
			}
			if !found {
				return fmt.Errorf("error setting value for %s: [%s] is not one of the valid options", paramName, selection)
			}
		}
	}

	return nil

}

// Handle a network change by terminating the service, deleting everything, and starting over
func changeNetworks(c *cli.Context, stader *stader.Client, apiContainerName string) error {

	// Stop all of the containers
	fmt.Print("Stopping containers... ")
	err := stader.PauseService(getComposeFiles(c))
	if err != nil {
		return fmt.Errorf("error stopping service: %w", err)
	}
	fmt.Println("done")

	// Restart the API container
	fmt.Print("Starting API container... ")
	output, err := stader.StartContainer(apiContainerName)
	if err != nil {
		return fmt.Errorf("error starting API container: %w", err)
	}
	if output != apiContainerName {
		return fmt.Errorf("starting API container had unexpected output: %s", output)
	}
	fmt.Println("done")

	// Get the path of the user's data folder
	fmt.Print("Retrieving data folder path... ")
	volumePath, err := stader.GetClientVolumeSource(apiContainerName, dataFolderVolumeName)
	if err != nil {
		return fmt.Errorf("error getting data folder path: %w", err)
	}
	fmt.Printf("done, data folder = %s\n", volumePath)

	// Delete the data folder
	fmt.Print("Removing data folder... ")
	_, err = stader.TerminateDataFolder()
	if err != nil {
		return err
	}
	fmt.Println("done")

	// Terminate the current setup
	fmt.Print("Removing old installation... ")
	err = stader.StopService(getComposeFiles(c))
	if err != nil {
		return fmt.Errorf("error terminating old installation: %w", err)
	}
	fmt.Println("done")

	// Create new validator folder
	fmt.Print("Recreating data folder... ")
	err = os.MkdirAll(filepath.Join(volumePath, "validators"), 0775)
	if err != nil {
		return fmt.Errorf("error recreating data folder: %w", err)
	}

	// Start the service
	fmt.Print("Starting Stader... ")
	err = stader.StartService(getComposeFiles(c))
	if err != nil {
		return fmt.Errorf("error starting service: %w", err)
	}
	fmt.Println("done")

	return nil

}

// Start the Stader service
func startService(c *cli.Context, ignoreConfigSuggestion bool) error {

	staderClient, err := stader.NewClientFromCtx(c)
	if err != nil {
		return err
	}
	defer staderClient.Close()

	// Update the Prometheus template with the assigned ports
	cfg, isNew, err := staderClient.LoadConfig()
	if err != nil {
		return fmt.Errorf("Error loading user settings: %w", err)
	}

	// Check for unsupported clients
	if cfg.ExecutionClientMode.Value.(cfgtypes.Mode) == cfgtypes.Mode_Local {
		selectedEc := cfg.ExecutionClient.Value.(cfgtypes.ExecutionClient)
		switch selectedEc {
		case cfgtypes.ExecutionClient_Obs_Infura:
			fmt.Printf("%sYou currently have Infura configured as your primary Execution client, but it is no longer supported because it is not compatible with the upcoming Ethereum Merge.\nPlease run `stader-cli service config` and select a full Execution client.%s\n", colorRed, colorReset)
			return nil
		case cfgtypes.ExecutionClient_Obs_Pocket:
			fmt.Printf("%sYou currently have Pocket configured as your primary Execution client, but it is no longer supported because it is not compatible with the upcoming Ethereum Merge.\nPlease run `stader-cli service config` and select a full Execution client.%s\n", colorRed, colorReset)
			return nil
		}
	}

	// Force all Docker or all Hybrid
	if cfg.ExecutionClientMode.Value.(cfgtypes.Mode) == cfgtypes.Mode_Local && cfg.ConsensusClientMode.Value.(cfgtypes.Mode) == cfgtypes.Mode_External {
		fmt.Printf("%sYou are using a locally-managed Execution client and an externally-managed Consensus client.\nThis configuration is not compatible with The Merge; please select either locally-managed or externally-managed for both the EC and CC.%s\n", colorRed, colorReset)
	} else if cfg.ExecutionClientMode.Value.(cfgtypes.Mode) == cfgtypes.Mode_External && cfg.ConsensusClientMode.Value.(cfgtypes.Mode) == cfgtypes.Mode_Local {
		fmt.Printf("%sYou are using an externally-managed Execution client and a locally-managed Consensus client.\nThis configuration is not compatible with The Merge; please select either locally-managed or externally-managed for both the EC and CC.%s\n", colorRed, colorReset)
	}

	isMigration := false
	if isNew {
		// Look for a legacy config to migrate
		migratedConfig, err := staderClient.LoadLegacyConfigFromBackup()
		if err != nil {
			return err
		}
		if migratedConfig != nil {
			cfg = migratedConfig
			isMigration = true
		}
	}

	if isMigration {
		return fmt.Errorf("You must upgrade your configuration before starting the Stadernode.\nPlease run `stader-cli service config` to confirm your settings were migrated correctly, and enjoy the new configuration UI!")
	} else if isNew {
		return fmt.Errorf("No configuration detected. Please run `stader-cli service config` to set up your Stadernode before running it.")
	}

	if !ignoreConfigSuggestion {
		if c.Bool("yes") || cliutils.Confirm("Stadernode upgrade detected - starting will overwrite certain settings with the latest defaults (such as container versions).\nYou may want to run `service config` first to see what's changed.\n\nWould you like to continue starting the service?") {
			err = cfg.UpdateDefaults()
			if err != nil {
				return fmt.Errorf("error upgrading configuration with the latest parameters: %w", err)
			}
			staderClient.SaveConfig(cfg)
			fmt.Printf("%sUpdated settings successfully.%s\n", colorGreen, colorReset)
		} else {
			fmt.Println("Cancelled.")
			return nil
		}
	}

	// Update the Prometheus template with the assigned ports
	metricsEnabled := cfg.EnableMetrics.Value.(bool)
	if metricsEnabled {
		err := staderClient.UpdatePrometheusConfiguration(cfg.GenerateEnvironmentVariables())
		if err != nil {
			return err
		}
	}

	// Validate the config
	errors := cfg.Validate()
	if len(errors) > 0 {
		fmt.Printf("%sYour configuration encountered errors. You must correct the following in order to start Stader:\n\n", colorRed)
		for _, err := range errors {
			fmt.Printf("%s\n\n", err)
		}
		fmt.Println(colorReset)
		return nil
	}

	if !c.Bool("ignore-slash-timer") {
		// Do the client swap check
		err := checkForValidatorChange(staderClient, cfg)
		if err != nil {
			fmt.Printf("%sWarning: couldn't verify that the validator container can be safely restarted:\n\t%s\n", colorYellow, err.Error())
			fmt.Println("If you are changing to a different ETH2 client, it may resubmit an attestation you have already submitted.")
			fmt.Println("This will slash your validator!")
			fmt.Println("To prevent slashing, you must wait 15 minutes from the time you stopped the clients before starting them again.\n")
			fmt.Println("**If you did NOT change clients, you can safely ignore this warning.**\n")
			if !cliutils.Confirm(fmt.Sprintf("Press y when you understand the above warning, have waited, and are ready to start Stader:%s", colorReset)) {
				fmt.Println("Cancelled.")
				return nil
			}
		}
	} else {
		fmt.Printf("%sIgnoring anti-slashing safety delay.%s\n", colorYellow, colorReset)
	}

	// Write a note on doppelganger protection
	doppelgangerEnabled, err := cfg.IsDoppelgangerEnabled()
	if err != nil {
		fmt.Printf("%sCouldn't check if you have Doppelganger Protection enabled: %s\nIf you do, your validator will miss up to 3 attestations when it starts.\nThis is *intentional* and does not indicate a problem with your node.%s\n\n", colorYellow, err.Error(), colorReset)
	} else if doppelgangerEnabled {
		fmt.Printf("%sNOTE: You currently have Doppelganger Protection enabled.\nYour validator will miss up to 3 attestations when it starts.\nThis is *intentional* and does not indicate a problem with your node.%s\n\n", colorYellow, colorReset)
	}

	println("Starting Stader Service")

	// Start service
	err = staderClient.StartService(getComposeFiles(c))
	if err != nil {
		return err
	}

	// Remove the upgrade flag if it's there
	return staderClient.RemoveUpgradeFlagFile()

}

func checkForValidatorChange(stader *stader.Client, cfg *config.StaderConfig) error {

	// Get the container prefix
	prefix, err := getContainerPrefix(stader)
	if err != nil {
		return fmt.Errorf("Error getting validator container prefix: %w", err)
	}

	// Get the current validator client
	currentValidatorImageString, err := stader.GetDockerImage(prefix + ValidatorContainerSuffix)
	if err != nil {
		return fmt.Errorf("Error getting current validator image: %w", err)
	}

	currentValidatorName, err := getDockerImageName(currentValidatorImageString)
	if err != nil {
		return fmt.Errorf("Error getting current validator image name: %w", err)
	}

	// Get the new validator client according to the settings file
	selectedConsensusClientConfig, err := cfg.GetSelectedConsensusClientConfig()
	if err != nil {
		return fmt.Errorf("Error getting selected consensus client config: %w", err)
	}
	pendingValidatorName, err := getDockerImageName(selectedConsensusClientConfig.GetValidatorImage())
	if err != nil {
		return fmt.Errorf("Error getting pending validator image name: %w", err)
	}

	// Compare the clients and warn if necessary
	if currentValidatorName == pendingValidatorName {
		fmt.Printf("Validator client [%s] was previously used - no slashing prevention delay necessary.\n", currentValidatorName)
	} else if currentValidatorName == "" {
		fmt.Println("This is the first time starting Stader - no slashing prevention delay necessary.")
	} else {

		// Get the time that the container responsible for validator duties exited
		validatorDutyContainerName, err := getContainerNameForValidatorDuties(currentValidatorName, stader)
		if err != nil {
			return fmt.Errorf("Error getting validator container name: %w", err)
		}
		validatorFinishTime, err := stader.GetDockerContainerShutdownTime(validatorDutyContainerName)
		if err != nil {
			return fmt.Errorf("Error getting validator shutdown time: %w", err)
		}

		// If it hasn't exited yet, shut it down
		zeroTime := time.Time{}
		status, err := stader.GetDockerStatus(validatorDutyContainerName)
		if err != nil {
			return fmt.Errorf("Error getting container [%s] status: %w", validatorDutyContainerName, err)
		}
		if validatorFinishTime == zeroTime || status == "running" {
			fmt.Printf("%sValidator is currently running, stopping it...%s\n", colorYellow, colorReset)
			response, err := stader.StopContainer(validatorDutyContainerName)
			validatorFinishTime = time.Now()
			if err != nil {
				return fmt.Errorf("Error stopping container [%s]: %w", validatorDutyContainerName, err)
			}
			if response != validatorDutyContainerName {
				return fmt.Errorf("Unexpected response when stopping container [%s]: %s", validatorDutyContainerName, response)
			}
		}

		// Print the warning and start the time lockout
		safeStartTime := validatorFinishTime.Add(15 * time.Minute)
		remainingTime := time.Until(safeStartTime)
		if remainingTime <= 0 {
			fmt.Printf("The validator has been offline for %s, which is long enough to prevent slashing.\n", time.Since(validatorFinishTime))
			fmt.Println("The new client can be safely started.")
		} else {
			fmt.Printf("%s=== WARNING ===\n", colorRed)
			fmt.Printf("You have changed your validator client from %s to %s.\n", currentValidatorName, pendingValidatorName)
			fmt.Println("If you have active validators, starting the new client immediately will cause them to be slashed due to duplicate attestations!")
			fmt.Println("To prevent slashing, Stader will delay activating the new client for 15 minutes.")
			fmt.Printf("If you want to bypass this cooldown and understand the risks, run `stader-cli service start --ignore-slash-timer`.%s\n\n", colorReset)

			// Wait for 15 minutes
			for remainingTime > 0 {
				fmt.Printf("Remaining time: %s", remainingTime)
				time.Sleep(1 * time.Second)
				remainingTime = time.Until(safeStartTime)
				fmt.Printf("%s\r", clearLine)
			}

			fmt.Println(colorReset)
			fmt.Println("You may now safely start the validator without fear of being slashed.")
		}
	}

	return nil
}

// Get the name of the container responsible for validator duties based on the client name
// TODO: this is temporary and can change, clean it up when Nimbus supports split mode
func getContainerNameForValidatorDuties(CurrentValidatorClientName string, sd *stader.Client) (string, error) {

	prefix, err := getContainerPrefix(sd)
	if err != nil {
		return "", err
	}

	if CurrentValidatorClientName == "nimbus" {
		return prefix + BeaconContainerSuffix, nil
	}

	return prefix + ValidatorContainerSuffix, nil

}

// Get the time that the container responsible for validator duties exited
func getValidatorFinishTime(CurrentValidatorClientName string, staderClient *stader.Client) (time.Time, error) {

	prefix, err := getContainerPrefix(staderClient)
	if err != nil {
		return time.Time{}, err
	}

	var validatorFinishTime time.Time
	if CurrentValidatorClientName == "nimbus" {
		validatorFinishTime, err = staderClient.GetDockerContainerShutdownTime(prefix + BeaconContainerSuffix)
	} else {
		validatorFinishTime, err = staderClient.GetDockerContainerShutdownTime(prefix + ValidatorContainerSuffix)
	}

	return validatorFinishTime, err

}

// Extract the image name from a Docker image string
func getDockerImageName(imageString string) (string, error) {

	// Return the empty string if the validator didn't exist (probably because this is the first time starting it up)
	if imageString == "" {
		return "", nil
	}

	reg := regexp.MustCompile(dockerImageRegex)
	matches := reg.FindStringSubmatch(imageString)
	if matches == nil {
		return "", fmt.Errorf("Couldn't parse the Docker image string [%s]", imageString)
	}
	imageIndex := reg.SubexpIndex("image")
	if imageIndex == -1 {
		return "", fmt.Errorf("Image name not found in Docker image [%s]", imageString)
	}

	imageName := matches[imageIndex]
	return imageName, nil
}

// Gets the prefix specified for Stader's Docker containers
func getContainerPrefix(staderClient *stader.Client) (string, error) {

	cfg, isNew, err := staderClient.LoadConfig()
	if err != nil {
		return "", err
	}
	if isNew {
		return "", fmt.Errorf("Settings file not found. Please run `stader-cli service config` to set up your Stadernode.")
	}

	return cfg.StaderNode.ProjectName.Value.(string), nil
}

// Prepares the execution client for pruning
func pruneExecutionClient(c *cli.Context) error {

	staderClient, err := stader.NewClientFromCtx(c)
	if err != nil {
		return err
	}
	defer staderClient.Close()

	// Get the config
	cfg, isNew, err := staderClient.LoadConfig()
	if err != nil {
		return err
	}
	if isNew {
		return fmt.Errorf("Settings file not found. Please run `stader-cli service config` to set up your Stadernode.")
	}

	// Sanity checks
	if cfg.ExecutionClientMode.Value.(cfgtypes.Mode) == cfgtypes.Mode_External {
		fmt.Println("You are using an externally managed Execution client.\nThe Stadernode cannot prune it for you.")
		return nil
	}
	if cfg.IsNativeMode {
		fmt.Println("You are using Native Mode.\nThe Stadernode cannot prune your Execution client for you, you'll have to do it manually.")
	}
	selectedEc := cfg.ExecutionClient.Value.(cfgtypes.ExecutionClient)
	switch selectedEc {
	case cfgtypes.ExecutionClient_Besu:
		fmt.Println("You are using Besu as your Execution client.\nBesu does not need pruning.")
		return nil
	}

	fmt.Println("This will shut down your main execution client and prune its database, freeing up disk space.")
	fmt.Println("Once pruning is complete, your execution client will restart automatically.\n")

	if selectedEc == cfgtypes.ExecutionClient_Geth {
		if cfg.UseFallbackClients.Value == false {
			fmt.Printf("%sYou do not have a fallback execution client configured.\nYour node will no longer be able to perform any validation duties (attesting or proposing blocks) until Geth is done pruning and has synced again.\nPlease configure a fallback client with `stader-cli service config` before running this.%s\n", colorRed, colorReset)
		} else {
			fmt.Println("You have fallback clients enabled. Stader (and your consensus client) will use that while the main client is pruning.")
		}
	}

	// Get the container prefix
	prefix, err := getContainerPrefix(staderClient)
	if err != nil {
		return fmt.Errorf("Error getting container prefix: %w", err)
	}

	// Prompt for confirmation
	if !(c.Bool("yes") || cliutils.Confirm("Are you sure you want to prune your main execution client?")) {
		fmt.Println("Cancelled.")
		return nil
	}

	// Get the prune provisioner image
	pruneProvisioner := cfg.StaderNode.GetPruneProvisionerContainerTag()

	// Check for enough free space
	executionContainerName := prefix + ExecutionContainerSuffix
	volumePath, err := staderClient.GetClientVolumeSource(executionContainerName, clientDataVolumeName)
	if err != nil {
		return fmt.Errorf("Error getting execution volume source path: %w", err)
	}
	partitions, err := disk.Partitions(true)
	if err != nil {
		return fmt.Errorf("Error getting partition list: %w", err)
	}

	longestPath := 0
	bestPartition := disk.PartitionStat{}
	for _, partition := range partitions {
		if strings.HasPrefix(volumePath, partition.Mountpoint) && len(partition.Mountpoint) > longestPath {
			bestPartition = partition
			longestPath = len(partition.Mountpoint)
		}
	}

	diskUsage, err := disk.Usage(bestPartition.Mountpoint)
	if err != nil {
		return fmt.Errorf("Error getting free disk space available: %w", err)
	}
	freeSpaceHuman := humanize.IBytes(diskUsage.Free)
	if diskUsage.Free < PruneFreeSpaceRequired {
		return fmt.Errorf("%sYour disk must have 50 GiB free to prune, but it only has %s free. Please free some space before pruning.%s", colorRed, freeSpaceHuman, colorReset)
	}

	fmt.Printf("Your disk has %s free, which is enough to prune.\n", freeSpaceHuman)

	fmt.Printf("Stopping %s...\n", executionContainerName)
	result, err := staderClient.StopContainer(executionContainerName)
	if err != nil {
		return fmt.Errorf("Error stopping main execution container: %w", err)
	}
	if result != executionContainerName {
		return fmt.Errorf("Unexpected output while stopping main execution container: %s", result)
	}

	// Get the ETH1 volume name
	volume, err := staderClient.GetClientVolumeName(executionContainerName, clientDataVolumeName)
	if err != nil {
		return fmt.Errorf("Error getting execution client volume name: %w", err)
	}

	// Run the prune provisioner
	fmt.Printf("Provisioning pruning on volume %s...\n", volume)
	err = staderClient.RunPruneProvisioner(prefix+PruneProvisionerContainerSuffix, volume, pruneProvisioner)
	if err != nil {
		return fmt.Errorf("Error running prune provisioner: %w", err)
	}

	// Restart ETH1
	fmt.Printf("Restarting %s...\n", executionContainerName)
	result, err = staderClient.StartContainer(executionContainerName)
	if err != nil {
		return fmt.Errorf("Error starting main execution client: %w", err)
	}
	if result != executionContainerName {
		return fmt.Errorf("Unexpected output while starting main execution client: %s", result)
	}

	if selectedEc == cfgtypes.ExecutionClient_Nethermind {
		err = staderClient.RunNethermindPruneStarter(executionContainerName)
		if err != nil {
			return fmt.Errorf("Error starting Nethermind prune starter: %w", err)
		}
	}

	fmt.Printf("\nDone! Your main execution client is now pruning. You can follow its progress with `stader-cli service logs eth1`.\n")
	fmt.Println("Once it's done, it will restart automatically and resume normal operation.")

	fmt.Printf("%sNOTE: While pruning, you **cannot** interrupt the client (e.g. by restarting) or you risk corrupting the database!\nYou must let it run to completion!%s\n", colorYellow, colorReset)

	return nil

}

// Pause the Stader service
func pauseService(c *cli.Context) error {

	staderClient, err := stader.NewClientFromCtx(c)
	if err != nil {
		return err
	}
	defer staderClient.Close()

	// Get the config
	cfg, _, err := staderClient.LoadConfig()
	if err != nil {
		return err
	}

	// Write a note on doppelganger protection
	doppelgangerEnabled, err := cfg.IsDoppelgangerEnabled()
	if err != nil {
		fmt.Printf("%sCouldn't check if you have Doppelganger Protection enabled: %s\nIf you do, stopping your validator will cause it to miss up to 3 attestations when it next starts.\nThis is *intentional* and does not indicate a problem with your node.%s\n\n", colorYellow, err.Error(), colorReset)
	} else if doppelgangerEnabled {
		fmt.Printf("%sNOTE: You currently have Doppelganger Protection enabled.\nIf you stop your validator, it will miss up to 3 attestations when it next starts.\nThis is *intentional* and does not indicate a problem with your node.%s\n\n", colorYellow, colorReset)
	}

	// Prompt for confirmation
	if !(c.Bool("yes") || cliutils.Confirm("Are you sure you want to pause the Stader service? Any staking validators will be penalized!")) {
		fmt.Println("Cancelled.")
		return nil
	}

	// Pause service
	return staderClient.PauseService(getComposeFiles(c))

}

// Stop the Stader service
func stopService(c *cli.Context) error {

	// Prompt for confirmation
	if !(c.Bool("yes") || cliutils.Confirm(fmt.Sprintf("%sWARNING: Are you sure you want to terminate the Stader service? Any validators will be penalized, your ETH1 and ETH2 chain databases will be deleted, you will lose ALL of your sync progress, and you will lose your Prometheus metrics database!%s", colorRed, colorReset))) {
		fmt.Println("Cancelled.")
		return nil
	}

	staderClient, err := stader.NewClientFromCtx(c)
	if err != nil {
		return err
	}
	defer staderClient.Close()

	// Stop service
	return staderClient.StopService(getComposeFiles(c))

}

// View the Stader service logs
func serviceLogs(c *cli.Context, serviceNames ...string) error {

	staderClient, err := stader.NewClientFromCtx(c)
	if err != nil {
		return err
	}
	defer staderClient.Close()

	// Print service logs
	return staderClient.PrintServiceLogs(getComposeFiles(c), c.String("tail"), serviceNames...)

}

// View the Stader service stats
func serviceStats(c *cli.Context) error {

	staderClient, err := stader.NewClientFromCtx(c)
	if err != nil {
		return err
	}
	defer staderClient.Close()

	// Print service stats
	return staderClient.PrintServiceStats(getComposeFiles(c))

}

// View the Stader service compose config
func serviceCompose(c *cli.Context) error {

	staderClient, err := stader.NewClientFromCtx(c)
	if err != nil {
		return err
	}
	defer staderClient.Close()

	// Print service compose config
	return staderClient.PrintServiceCompose(getComposeFiles(c))

}

// View the Stader service version information
func serviceVersion(c *cli.Context) error {

	stader, err := stader.NewClientFromCtx(c)
	if err != nil {
		return err
	}
	defer stader.Close()

	// Print what network we're on
	err = cliutils.PrintNetwork(stader)
	if err != nil {
		return err
	}

	serviceVersion, err := stader.GetServiceVersion()
	if err != nil {
		return err
	}

	// Get config
	cfg, isNew, err := stader.LoadConfig()
	if err != nil {
		return err
	}
	if isNew {
		return fmt.Errorf("settings file not found. Please run `stader-cli service config` to set up your Stadernode")
	}

	// Handle native mode
	if cfg.IsNativeMode {
		fmt.Printf("Stader client version: %s\n", c.App.Version)
		fmt.Printf("Stader service version: %s\n", serviceVersion)
		fmt.Println("Configured for Native Mode")
		return nil
	}

	// Get the execution client string
	var eth1ClientString string
	eth1ClientMode := cfg.ExecutionClientMode.Value.(cfgtypes.Mode)
	switch eth1ClientMode {
	case cfgtypes.Mode_Local:
		eth1Client := cfg.ExecutionClient.Value.(cfgtypes.ExecutionClient)
		format := "%s (Locally managed)\n\tImage: %s"
		switch eth1Client {
		case cfgtypes.ExecutionClient_Geth:
			eth1ClientString = fmt.Sprintf(format, "Geth", cfg.Geth.ContainerTag.Value.(string))
		case cfgtypes.ExecutionClient_Nethermind:
			eth1ClientString = fmt.Sprintf(format, "Nethermind", cfg.Nethermind.ContainerTag.Value.(string))
		case cfgtypes.ExecutionClient_Besu:
			eth1ClientString = fmt.Sprintf(format, "Besu", cfg.Besu.ContainerTag.Value.(string))
		default:
			return fmt.Errorf("unknown local execution client [%v]", eth1Client)
		}

	case cfgtypes.Mode_External:
		eth1ClientString = "Externally managed"

	default:
		return fmt.Errorf("unknown execution client mode [%v]", eth1ClientMode)
	}

	// Get the consensus client string
	var eth2ClientString string
	eth2ClientMode := cfg.ConsensusClientMode.Value.(cfgtypes.Mode)
	switch eth2ClientMode {
	case cfgtypes.Mode_Local:
		eth2Client := cfg.ConsensusClient.Value.(cfgtypes.ConsensusClient)
		format := "%s (Locally managed)\n\tImage: %s"
		switch eth2Client {
		case cfgtypes.ConsensusClient_Lighthouse:
			eth2ClientString = fmt.Sprintf(format, "Lighthouse", cfg.Lighthouse.ContainerTag.Value.(string))
		case cfgtypes.ConsensusClient_Nimbus:
			eth2ClientString = fmt.Sprintf(format, "Nimbus", cfg.Nimbus.BnContainerTag.Value.(string))
		case cfgtypes.ConsensusClient_Prysm:
			// Prysm is a special case, as the BN and VC image versions may differ
			eth2ClientString = fmt.Sprintf(format+"\n\tVC image: %s", "Prysm", cfg.Prysm.BnContainerTag.Value.(string), cfg.Prysm.VcContainerTag.Value.(string))
		case cfgtypes.ConsensusClient_Teku:
			eth2ClientString = fmt.Sprintf(format, "Teku", cfg.Teku.ContainerTag.Value.(string))
		default:
			return fmt.Errorf("unknown local consensus client [%v]", eth2Client)
		}

	case cfgtypes.Mode_External:
		eth2Client := cfg.ExternalConsensusClient.Value.(cfgtypes.ConsensusClient)
		format := "%s (Externally managed)\n\tVC Image: %s"
		switch eth2Client {
		case cfgtypes.ConsensusClient_Lighthouse:
			eth2ClientString = fmt.Sprintf(format, "Lighthouse", cfg.ExternalLighthouse.ContainerTag.Value.(string))
		case cfgtypes.ConsensusClient_Prysm:
			eth2ClientString = fmt.Sprintf(format, "Prysm", cfg.ExternalPrysm.ContainerTag.Value.(string))
		case cfgtypes.ConsensusClient_Teku:
			eth2ClientString = fmt.Sprintf(format, "Teku", cfg.ExternalTeku.ContainerTag.Value.(string))
		default:
			return fmt.Errorf("unknown external consensus client [%v]", eth2Client)
		}

	default:
		return fmt.Errorf("unknown consensus client mode [%v]", eth2ClientMode)
	}

	// Print version info
	fmt.Printf("Stader client version: %s\n", c.App.Version)
	fmt.Printf("Stader service version: %s\n", serviceVersion)
	fmt.Printf("Selected Eth 1.0 client: %s\n", eth1ClientString)
	fmt.Printf("Selected Eth 2.0 client: %s\n", eth2ClientString)
	return nil

}

// Get the compose file paths for a CLI context
func getComposeFiles(c *cli.Context) []string {
	return c.Parent().StringSlice("compose-file")
}

// Destroy and resync the eth1 client from scratch
func resyncEth1(c *cli.Context) error {

	staderClient, err := stader.NewClientFromCtx(c)
	if err != nil {
		return err
	}
	defer staderClient.Close()

	// Get the config
	_, isNew, err := staderClient.LoadConfig()
	if err != nil {
		return err
	}
	if isNew {
		return fmt.Errorf("Settings file not found. Please run `stader-cli service config` to set up your Stadernode.")
	}

	fmt.Println("This will delete the chain data of your primary ETH1 client and resync it from scratch.")
	fmt.Printf("%sYou should only do this if your ETH1 client has failed and can no longer start or sync properly.\nThis is meant to be a last resort.%s\n", colorYellow, colorReset)

	// Get the container prefix
	prefix, err := getContainerPrefix(staderClient)
	if err != nil {
		return fmt.Errorf("Error getting container prefix: %w", err)
	}

	// Prompt for confirmation
	if !(c.Bool("yes") || cliutils.Confirm(fmt.Sprintf("%sAre you SURE you want to delete and resync your main ETH1 client from scratch? This cannot be undone!%s", colorRed, colorReset))) {
		fmt.Println("Cancelled.")
		return nil
	}

	// Stop ETH1
	executionContainerName := prefix + ExecutionContainerSuffix
	fmt.Printf("Stopping %s...\n", executionContainerName)
	result, err := staderClient.StopContainer(executionContainerName)
	if err != nil {
		fmt.Printf("%sWARNING: Stopping main ETH1 container failed: %s%s\n", colorYellow, err.Error(), colorReset)
	}
	if result != executionContainerName {
		fmt.Printf("%sWARNING: Unexpected output while stopping main ETH1 container: %s%s\n", colorYellow, result, colorReset)
	}

	// Get ETH1 volume name
	volume, err := staderClient.GetClientVolumeName(executionContainerName, clientDataVolumeName)
	if err != nil {
		return fmt.Errorf("Error getting ETH1 volume name: %w", err)
	}

	// Remove ETH1
	fmt.Printf("Deleting %s...\n", executionContainerName)
	result, err = staderClient.RemoveContainer(executionContainerName)
	if err != nil {
		return fmt.Errorf("Error deleting main ETH1 container: %w", err)
	}
	if result != executionContainerName {
		return fmt.Errorf("Unexpected output while deleting main ETH1 container: %s", result)
	}

	// Delete the ETH1 volume
	fmt.Printf("Deleting volume %s...\n", volume)
	result, err = staderClient.DeleteVolume(volume)
	if err != nil {
		return fmt.Errorf("Error deleting volume: %w", err)
	}
	if result != volume {
		return fmt.Errorf("Unexpected output while deleting volume: %s", result)
	}

	// Restart Stader
	fmt.Printf("Rebuilding %s and restarting Stader...\n", executionContainerName)
	err = startService(c, true)
	if err != nil {
		return fmt.Errorf("Error starting Stader: %s", err)
	}

	fmt.Printf("\nDone! Your main ETH1 client is now resyncing. You can follow its progress with `stader-cli service logs eth1`.\n")

	return nil

}

// Destroy and resync the eth2 client from scratch
func resyncEth2(c *cli.Context) error {

	staderClient, err := stader.NewClientFromCtx(c)
	if err != nil {
		return err
	}
	defer staderClient.Close()

	// Get the merged config
	cfg, isNew, err := staderClient.LoadConfig()
	if err != nil {
		return err
	}
	if isNew {
		return fmt.Errorf("Settings file not found. Please run `stader-cli service config` to set up your Stadernode.")
	}

	fmt.Println("This will delete the chain data of your ETH2 client and resync it from scratch.")
	fmt.Printf("%sYou should only do this if your ETH2 client has failed and can no longer start or sync properly.\nThis is meant to be a last resort.%s\n\n", colorYellow, colorReset)

	// Get the parameters that the selected client doesn't support
	var unsupportedParams []string
	var clientName string
	eth2ClientMode := cfg.ConsensusClientMode.Value.(cfgtypes.Mode)
	switch eth2ClientMode {
	case cfgtypes.Mode_Local:
		selectedClientConfig, err := cfg.GetSelectedConsensusClientConfig()
		if err != nil {
			return fmt.Errorf("error getting selected consensus client config: %w", err)
		}
		unsupportedParams = selectedClientConfig.(cfgtypes.LocalConsensusConfig).GetUnsupportedCommonParams()
		clientName = selectedClientConfig.GetName()

	case cfgtypes.Mode_External:
		fmt.Println("You use an externally-managed Consensus client. Stader cannot resync it for you.")
		return nil

	default:
		return fmt.Errorf("unknown consensus client mode [%v]", eth2ClientMode)
	}

	// Check if the selected client supports checkpoint sync
	supportsCheckpointSync := true
	for _, param := range unsupportedParams {
		if param == config.CheckpointSyncUrlID {
			supportsCheckpointSync = false
		}
	}
	if !supportsCheckpointSync {
		fmt.Printf("%sYour ETH2 client (%s) does not support checkpoint sync.\nIf you have active validators, they %swill be considered offline and will leak ETH%s%s while the client is syncing.%s\n\n", colorRed, clientName, colorBold, colorReset, colorRed, colorReset)
	} else {
		// Get the current checkpoint sync URL
		checkpointSyncUrl := cfg.ConsensusCommon.CheckpointSyncProvider.Value.(string)
		if checkpointSyncUrl == "" {
			fmt.Printf("%sYou do not have a checkpoint sync provider configured.\nIf you have active validators, they %swill be considered offline and will lose ETH%s%s until your ETH2 client finishes syncing.\nWe strongly recommend you configure a checkpoint sync provider with `stader-cli service config` so it syncs instantly before running this.%s\n\n", colorRed, colorBold, colorReset, colorRed, colorReset)
		} else {
			fmt.Printf("You have a checkpoint sync provider configured (%s).\nYour ETH2 client will use it to sync to the head of the Beacon Chain instantly after being rebuilt.\n\n", checkpointSyncUrl)
		}
	}

	// Get the container prefix
	prefix, err := getContainerPrefix(staderClient)
	if err != nil {
		return fmt.Errorf("Error getting container prefix: %w", err)
	}

	// Prompt for confirmation
	if !(c.Bool("yes") || cliutils.Confirm(fmt.Sprintf("%sAre you SURE you want to delete and resync your main ETH2 client from scratch? This cannot be undone!%s", colorRed, colorReset))) {
		fmt.Println("Cancelled.")
		return nil
	}

	// Stop ETH2
	beaconContainerName := prefix + BeaconContainerSuffix
	fmt.Printf("Stopping %s...\n", beaconContainerName)
	result, err := staderClient.StopContainer(beaconContainerName)
	if err != nil {
		fmt.Printf("%sWARNING: Stopping ETH2 container failed: %s%s\n", colorYellow, err.Error(), colorReset)
	}
	if result != beaconContainerName {
		fmt.Printf("%sWARNING: Unexpected output while stopping ETH2 container: %s%s\n", colorYellow, result, colorReset)
	}

	// Get ETH2 volume name
	volume, err := staderClient.GetClientVolumeName(beaconContainerName, clientDataVolumeName)
	if err != nil {
		return fmt.Errorf("Error getting ETH2 volume name: %w", err)
	}

	// Remove ETH2
	fmt.Printf("Deleting %s...\n", beaconContainerName)
	result, err = staderClient.RemoveContainer(beaconContainerName)
	if err != nil {
		return fmt.Errorf("Error deleting ETH2 container: %w", err)
	}
	if result != beaconContainerName {
		return fmt.Errorf("Unexpected output while deleting ETH2 container: %s", result)
	}

	// Delete the ETH2 volume
	fmt.Printf("Deleting volume %s...\n", volume)
	result, err = staderClient.DeleteVolume(volume)
	if err != nil {
		return fmt.Errorf("Error deleting volume: %w", err)
	}
	if result != volume {
		return fmt.Errorf("Unexpected output while deleting volume: %s", result)
	}

	// Restart Stader
	fmt.Printf("Rebuilding %s and restarting Stader...\n", beaconContainerName)
	err = startService(c, true)
	if err != nil {
		return fmt.Errorf("Error starting Stader: %s", err)
	}

	fmt.Printf("\nDone! Your ETH2 client is now resyncing. You can follow its progress with `stader-cli service logs eth2`.\n")

	return nil

}

// Generate a YAML file that shows the current configuration schema, including all of the parameters and their descriptions
func getConfigYaml(c *cli.Context) error {
	cfg := config.NewStaderConfig("", false)
	bytes, err := yaml.Marshal(cfg)
	if err != nil {
		return fmt.Errorf("error serializing configuration schema: %w", err)
	}

	fmt.Println(string(bytes))
	return nil
}

// Export the EC volume to an external folder
func exportEcData(c *cli.Context, targetDir string) error {

	staderClient, err := stader.NewClientFromCtx(c)
	if err != nil {
		return err
	}
	defer staderClient.Close()

	// Get the config
	cfg, isNew, err := staderClient.LoadConfig()
	if err != nil {
		return err
	}
	if isNew {
		return fmt.Errorf("Settings file not found. Please run `stader-cli service config` to set up your Stadernode.")
	}

	// Make the path absolute
	targetDir, err = filepath.Abs(targetDir)
	if err != nil {
		return fmt.Errorf("Error converting to absolute path: %w", err)
	}

	// Make sure the target dir exists and is accessible
	targetDirInfo, err := os.Stat(targetDir)
	if os.IsNotExist(err) {
		return fmt.Errorf("Target directory [%s] does not exist.", targetDir)
	} else if err != nil {
		return fmt.Errorf("Error reading target dir: %w", err)
	}
	if !targetDirInfo.IsDir() {
		return fmt.Errorf("Target directory [%s] is not a directory.", targetDir)
	}

	fmt.Println("This will export your execution client's chain data to an external directory, such as a portable hard drive.")
	fmt.Println("If your execution client is running, it will be shut down.")
	fmt.Println("Once the export is complete, your execution client will restart automatically.\n")

	// Get the container prefix
	prefix, err := getContainerPrefix(staderClient)
	if err != nil {
		return fmt.Errorf("Error getting container prefix: %w", err)
	}

	// Get the EC volume name
	executionContainerName := prefix + ExecutionContainerSuffix
	volume, err := staderClient.GetClientVolumeName(executionContainerName, clientDataVolumeName)
	if err != nil {
		return fmt.Errorf("Error getting execution client volume name: %w", err)
	}

	if !c.Bool("force") {
		// Make sure the target dir has enough space
		volumeBytes, err := getVolumeSpaceUsed(staderClient, volume)
		if err != nil {
			fmt.Printf("%sWARNING: Couldn't check the disk space used by the Execution client volume: %s\nPlease verify you have enough free space to store the chain data in the target folder before proceeding!%s\n\n", colorRed, err.Error(), colorReset)
		} else {
			volumeBytesHuman := humanize.IBytes(volumeBytes)
			targetFree, err := getPartitionFreeSpace(staderClient, targetDir)
			if err != nil {
				fmt.Printf("%sWARNING: Couldn't get the free space available on the target folder: %s\nPlease verify you have enough free space to store the chain data in the target folder before proceeding!%s\n\n", colorRed, err.Error(), colorReset)
			} else {
				freeSpaceHuman := humanize.IBytes(targetFree)
				fmt.Printf("%sChain data size:       %s%s\n", colorLightBlue, volumeBytesHuman, colorReset)
				fmt.Printf("%sTarget dir free space: %s%s\n", colorLightBlue, freeSpaceHuman, colorReset)
				if targetFree < volumeBytes {
					return fmt.Errorf("%sYour target directory does not have enough space to hold the chain data. Please free up more space and try again or use the --force flag to ignore this check.%s", colorRed, colorReset)
				}

				fmt.Printf("%sYour target directory has enough space to store the chain data.%s\n\n", colorGreen, colorReset)
			}
		}
	}

	// Prompt for confirmation
	fmt.Printf("%sNOTE: Once started, this process *will not stop* until the export is complete - even if you exit the command with Ctrl+C.\nPlease do not exit until it finishes so you can watch its progress.%s\n\n", colorYellow, colorReset)
	if !(c.Bool("yes") || cliutils.Confirm("Are you sure you want to export your execution layer chain data?")) {
		fmt.Println("Cancelled.")
		return nil
	}

	fmt.Printf("Stopping %s...\n", executionContainerName)
	result, err := staderClient.StopContainer(executionContainerName)
	if err != nil {
		return fmt.Errorf("Error stopping main execution container: %w", err)
	}
	if result != executionContainerName {
		return fmt.Errorf("Unexpected output while stopping main execution container: %s", result)
	}

	// Run the migrator
	ecMigrator := cfg.StaderNode.GetEcMigratorContainerTag()
	fmt.Printf("Exporting data from volume %s to %s...\n", volume, targetDir)
	err = staderClient.RunEcMigrator(prefix+EcMigratorContainerSuffix, volume, targetDir, "export", ecMigrator)
	if err != nil {
		return fmt.Errorf("Error running EC migrator: %w", err)
	}

	// Restart ETH1
	fmt.Printf("Restarting %s...\n", executionContainerName)
	result, err = staderClient.StartContainer(executionContainerName)
	if err != nil {
		return fmt.Errorf("Error starting main execution client: %w", err)
	}
	if result != executionContainerName {
		return fmt.Errorf("Unexpected output while starting main execution client: %s", result)
	}

	fmt.Println("\nDone! Your chain data has been exported.")

	return nil

}

// Import the EC volume from an external folder
func importEcData(c *cli.Context, sourceDir string) error {

	staderClient, err := stader.NewClientFromCtx(c)
	if err != nil {
		return err
	}
	defer staderClient.Close()

	// Get the config
	cfg, isNew, err := staderClient.LoadConfig()
	if err != nil {
		return err
	}
	if isNew {
		return fmt.Errorf("Settings file not found. Please run `stader-cli service config` to set up your Stadernode.")
	}

	// Make the path absolute
	sourceDir, err = filepath.Abs(sourceDir)
	if err != nil {
		return fmt.Errorf("Error converting to absolute path: %w", err)
	}

	// Get the container prefix
	prefix, err := getContainerPrefix(staderClient)
	if err != nil {
		return fmt.Errorf("Error getting container prefix: %w", err)
	}

	// Check the source dir
	fmt.Println("Checking source directory...")
	ecMigrator := cfg.StaderNode.GetEcMigratorContainerTag()
	sourceBytes, err := staderClient.GetDirSizeViaEcMigrator(prefix+EcMigratorContainerSuffix, sourceDir, ecMigrator)
	if err != nil {
		return err
	}

	fmt.Println("This will import execution layer chain data that you previously exported into your execution client.")
	fmt.Println("If your execution client is running, it will be shut down.")
	fmt.Println("Once the import is complete, your execution client will restart automatically.\n")

	// Get the volume to import into
	executionContainerName := prefix + ExecutionContainerSuffix
	volume, err := staderClient.GetClientVolumeName(executionContainerName, clientDataVolumeName)
	if err != nil {
		return fmt.Errorf("Error getting execution client volume name: %w", err)
	}

	// Make sure the target volume has enough space
	if err != nil {
		fmt.Printf("%sWARNING: Couldn't check the disk space used by the source folder: %s\nPlease verify you have enough free space to import the chain data before proceeding!%s\n\n", colorRed, err.Error(), colorReset)
	} else {
		sourceBytesHuman := humanize.IBytes(sourceBytes)
		volumePath, err := staderClient.GetClientVolumeSource(executionContainerName, clientDataVolumeName)
		if err != nil {
			err = fmt.Errorf("error getting execution volume source path: %w", err)
			fmt.Printf("%sWARNING: Couldn't check the disk space free on the Docker volume partition: %s\nPlease verify you have enough free space to import the chain data before proceeding!%s\n\n", colorRed, err.Error(), colorReset)
		} else {
			targetFree, err := getPartitionFreeSpace(staderClient, volumePath)
			if err != nil {
				fmt.Printf("%sWARNING: Couldn't check the disk space free on the Docker volume partition: %s\nPlease verify you have enough free space to import the chain data before proceeding!%s\n\n", colorRed, err.Error(), colorReset)
			} else {
				freeSpaceHuman := humanize.IBytes(targetFree)

				fmt.Printf("%sChain data size:         %s%s\n", colorLightBlue, sourceBytesHuman, colorReset)
				fmt.Printf("%sDocker drive free space: %s%s\n", colorLightBlue, freeSpaceHuman, colorReset)
				if targetFree < sourceBytes {
					return fmt.Errorf("%sYour Docker drive does not have enough space to hold the chain data. Please free up more space and try again.%s", colorRed, colorReset)
				}

				fmt.Printf("%sYour Docker drive has enough space to store the chain data.%s\n\n", colorGreen, colorReset)
			}
		}
	}

	// Prompt for confirmation
	fmt.Printf("%sNOTE: Importing will *delete* your existing chain data!%s\n\n", colorYellow, colorReset)
	fmt.Printf("%sOnce started, this process *will not stop* until the import is complete - even if you exit the command with Ctrl+C.\nPlease do not exit until it finishes so you can watch its progress.%s\n\n", colorYellow, colorReset)
	if !(c.Bool("yes") || cliutils.Confirm("Are you sure you want to delete your existing execution layer chain data and import other data from a backup?")) {
		fmt.Println("Cancelled.")
		return nil
	}

	fmt.Printf("Stopping %s...\n", executionContainerName)
	result, err := staderClient.StopContainer(executionContainerName)
	if err != nil {
		return fmt.Errorf("Error stopping main execution container: %w", err)
	}
	if result != executionContainerName {
		return fmt.Errorf("Unexpected output while stopping main execution container: %s", result)
	}

	// Run the migrator
	fmt.Printf("Importing data from %s to volume %s...\n", sourceDir, volume)
	err = staderClient.RunEcMigrator(prefix+EcMigratorContainerSuffix, volume, sourceDir, "import", ecMigrator)
	if err != nil {
		return fmt.Errorf("Error running EC migrator: %w", err)
	}

	// Restart ETH1
	fmt.Printf("Restarting %s...\n", executionContainerName)
	result, err = staderClient.StartContainer(executionContainerName)
	if err != nil {
		return fmt.Errorf("Error starting main execution client: %w", err)
	}
	if result != executionContainerName {
		return fmt.Errorf("Unexpected output while starting main execution client: %s", result)
	}

	fmt.Println("\nDone! Your chain data has been imported.")

	return nil
}

// Get the amount of space used by a Docker volume
func getVolumeSpaceUsed(staderClient *stader.Client, volume string) (uint64, error) {
	size, err := staderClient.GetVolumeSize(volume)
	if err != nil {
		return 0, fmt.Errorf("error getting execution client volume name: %w", err)
	}
	volumeBytes, err := humanize.ParseBytes(size)
	if err != nil {
		return 0, fmt.Errorf("couldn't parse size of EC volume (%s): %w", size, err)
	}
	return volumeBytes, nil
}

// Get the amount of free space available in the target dir
func getPartitionFreeSpace(staderClient *stader.Client, targetDir string) (uint64, error) {
	partitions, err := disk.Partitions(true)
	if err != nil {
		return 0, fmt.Errorf("error getting partition list: %w", err)
	}
	longestPath := 0
	bestPartition := disk.PartitionStat{}
	for _, partition := range partitions {
		if strings.HasPrefix(targetDir, partition.Mountpoint) && len(partition.Mountpoint) > longestPath {
			bestPartition = partition
			longestPath = len(partition.Mountpoint)
		}
	}
	diskUsage, err := disk.Usage(bestPartition.Mountpoint)
	if err != nil {
		return 0, fmt.Errorf("error getting free disk space available: %w", err)
	}
	return diskUsage.Free, nil
}

// Get the list of features required for modern client containers but not supported by the CPU
func checkCpuFeatures() error {
	unsupportedFeatures := sys.GetMissingModernCpuFeatures()
	if len(unsupportedFeatures) > 0 {
		fmt.Println("Your CPU is missing support for the following features:")
		for _, name := range unsupportedFeatures {
			fmt.Printf("  - %s\n", name)
		}

		fmt.Println("\nYou must use the 'portable' image.")
		return nil
	}

	fmt.Println("Your CPU supports all required features for 'modern' images.")
	return nil
}
