package config

import (
	"path/filepath"

	"github.com/ethereum/go-ethereum/common"
	"github.com/stader-labs/stader-node/shared"
	"github.com/stader-labs/stader-node/shared/types/config"
)

// Constants
const (
	smartnodeTag = "staderdev/stader-node:v" + shared.StaderVersion
	// TODO - arjay remove this and add them to stader docker namespace
	pruneProvisionerTag        string = "rocketpool/eth1-prune-provision:v0.0.1"
	ecMigratorTag              string = "rocketpool/ec-migrator:v1.0.0"
	NetworkID                  string = "network"
	ProjectNameID              string = "projectName"
	DaemonDataPath             string = "/.stader/data"
	WatchtowerFolder           string = "watchtower"
	FeeRecipientFilename       string = "stader-fee-recipient.txt"
	NativeFeeRecipientFilename string = "stader-fee-recipient-env.txt"
)

// --ignore-sync-check
// Defaults
const defaultProjectName string = "stader"

// Configuration for the Smartnode
type SmartnodeConfig struct {
	Title string `yaml:"-"`

	// The parent config
	parent *StaderConfig

	////////////////////////////
	// User-editable settings //
	////////////////////////////

	// Docker container prefix
	ProjectName config.Parameter `yaml:"projectName,omitempty"`

	// The path of the data folder where everything is stored
	DataPath config.Parameter `yaml:"dataPath,omitempty"`

	// The path of the watchtower's persistent state storage
	WatchtowerStatePath config.Parameter `yaml:"watchtowerStatePath"`

	// Which network we're on
	Network config.Parameter `yaml:"network,omitempty"`

	// Manual max fee override
	ManualMaxFee config.Parameter `yaml:"manualMaxFee,omitempty"`

	// Manual priority fee override
	PriorityFee config.Parameter `yaml:"priorityFee,omitempty"`

	// URL for an EC with archive mode, for manual rewards tree generation
	ArchiveECUrl config.Parameter `yaml:"archiveEcUrl,omitempty"`

	///////////////////////////
	// Non-editable settings //
	///////////////////////////

	// The URL to provide the user so they can follow pending transactions
	txWatchUrl map[config.Network]string `yaml:"-"`

	// The URL to use for staking EthX
	stakeUrl map[config.Network]string `yaml:"-"`

	// The map of networks to execution chain IDs
	chainID map[config.Network]uint `yaml:"-"`

	// The contract address of the permissionless node operator registry
	permissionlessNodeRegistryAddress map[config.Network]string `yaml:"-"`

	// The contract address of the withdraw vault factory
	vaultFactoryAddress map[config.Network]string `yaml:"-"`

	// The contract address of the SD collateral holder
	sdCollateralAddress map[config.Network]string `yaml:"-"`

	// The contract address of SD token ERC20
	sdTokenAddress map[config.Network]string `yaml:"-"`

	// The contract address of EthX ERC20
	ethxTokenAddress map[config.Network]string `yaml:"-"`
}

// Generates a new Smartnode configuration
func NewSmartnodeConfig(cfg *StaderConfig) *SmartnodeConfig {

	return &SmartnodeConfig{
		Title:  "Smartnode Settings",
		parent: cfg,

		ProjectName: config.Parameter{
			ID:                   ProjectNameID,
			Name:                 "Project Name",
			Description:          "This is the prefix that will be attached to all of the Docker containers managed by the Smartnode.",
			Type:                 config.ParameterType_String,
			Default:              map[config.Network]interface{}{config.Network_All: defaultProjectName},
			AffectsContainers:    []config.ContainerID{config.ContainerID_Api, config.ContainerID_Node, config.ContainerID_Watchtower, config.ContainerID_Eth1, config.ContainerID_Eth2, config.ContainerID_Validator, config.ContainerID_Grafana, config.ContainerID_Prometheus, config.ContainerID_Exporter},
			EnvironmentVariables: []string{"COMPOSE_PROJECT_NAME"},
			CanBeBlank:           false,
			OverwriteOnUpgrade:   false,
		},

		DataPath: config.Parameter{
			ID:                   "dataPath",
			Name:                 "Data Path",
			Description:          "The absolute path of the `data` folder that contains your node wallet's encrypted file, the password for your node wallet, and all of the validator keys for your validators. You may use environment variables in this string.",
			Type:                 config.ParameterType_String,
			Default:              map[config.Network]interface{}{config.Network_All: getDefaultDataDir(cfg)},
			AffectsContainers:    []config.ContainerID{config.ContainerID_Api, config.ContainerID_Node, config.ContainerID_Watchtower, config.ContainerID_Validator},
			EnvironmentVariables: []string{"ROCKETPOOL_DATA_FOLDER"},
			CanBeBlank:           false,
			OverwriteOnUpgrade:   false,
		},

		WatchtowerStatePath: config.Parameter{
			ID:                   "watchtowerPath",
			Name:                 "Watchtower Path",
			Description:          "The absolute path of the watchtower state folder that contains persistent state that is used by the watchtower process on trusted nodes. **Only relevant for trusted nodes.**",
			Type:                 config.ParameterType_String,
			Default:              map[config.Network]interface{}{config.Network_All: "$HOME/.stader/watchtower"},
			AffectsContainers:    []config.ContainerID{config.ContainerID_Watchtower},
			EnvironmentVariables: []string{"ROCKETPOOL_WATCHTOWER_FOLDER"},
			CanBeBlank:           false,
			OverwriteOnUpgrade:   false,
		},

		Network: config.Parameter{
			ID:                   NetworkID,
			Name:                 "Network",
			Description:          "The Ethereum network you want to use - select Goerli Testnet to practice with Goerli ETH, or Mainnet to stake on the real network using real ETH.",
			Type:                 config.ParameterType_Choice,
			Default:              map[config.Network]interface{}{config.Network_All: config.Network_Mainnet},
			AffectsContainers:    []config.ContainerID{config.ContainerID_Api, config.ContainerID_Node, config.ContainerID_Watchtower, config.ContainerID_Eth1, config.ContainerID_Eth2, config.ContainerID_Validator},
			EnvironmentVariables: []string{"NETWORK"},
			CanBeBlank:           false,
			OverwriteOnUpgrade:   false,
			Options:              getNetworkOptions(),
		},

		ManualMaxFee: config.Parameter{
			ID:                   "manualMaxFee",
			Name:                 "Manual Max Fee",
			Description:          "Set this if you want all of the Smartnode's transactions to use this specific max fee value (in gwei), which is the most you'd be willing to pay (*including the priority fee*).\n\nA value of 0 will show you the current suggested max fee based on the current network conditions and let you specify it each time you do a transaction.\n\nAny other value will ignore the recommended max fee and explicitly use this value instead.\n\nThis applies to automated transactions as well.",
			Type:                 config.ParameterType_Float,
			Default:              map[config.Network]interface{}{config.Network_All: float64(0)},
			AffectsContainers:    []config.ContainerID{config.ContainerID_Node, config.ContainerID_Watchtower},
			EnvironmentVariables: []string{},
			CanBeBlank:           false,
			OverwriteOnUpgrade:   false,
		},

		PriorityFee: config.Parameter{
			ID:                   "priorityFee",
			Name:                 "Priority Fee",
			Description:          "The default value for the priority fee (in gwei) for all of your transactions. This describes how much you're willing to pay *above the network's current base fee* - the higher this is, the more ETH you give to the validators for including your transaction, which generally means it will be included in a block faster (as long as your max fee is sufficiently high to cover the current network conditions).\n\nMust be larger than 0.",
			Type:                 config.ParameterType_Float,
			Default:              map[config.Network]interface{}{config.Network_All: float64(2)},
			AffectsContainers:    []config.ContainerID{config.ContainerID_Node, config.ContainerID_Watchtower},
			EnvironmentVariables: []string{},
			CanBeBlank:           false,
			OverwriteOnUpgrade:   false,
		},

		ArchiveECUrl: config.Parameter{
			ID:                   "archiveECUrl",
			Name:                 "Archive-Mode EC URL",
			Description:          "[orange]**For manual Merkle rewards tree generation only.**[white]\n\nGenerating the Merkle rewards tree files for past rewards intervals typically requires an Execution client with Archive mode enabled, which is usually disabled on your primary and fallback Execution clients to save disk space.\nIf you want to generate your own rewards tree files for intervals from a long time ago, you may enter the URL of an Execution client with Archive access here.\n\nFor a free light client with Archive access, you may use https://www.alchemy.com/supernode.",
			Type:                 config.ParameterType_String,
			Default:              map[config.Network]interface{}{config.Network_All: ""},
			AffectsContainers:    []config.ContainerID{config.ContainerID_Watchtower},
			EnvironmentVariables: []string{},
			CanBeBlank:           true,
			OverwriteOnUpgrade:   false,
		},

		txWatchUrl: map[config.Network]string{
			config.Network_Mainnet: "https://etherscan.io/tx",
			config.Network_Prater:  "https://goerli.etherscan.io/tx",
			config.Network_Devnet:  "https://goerli.etherscan.io/tx",
		},

		stakeUrl: map[config.Network]string{
			config.Network_Mainnet: "https://stake.rocketpool.net",
			config.Network_Prater:  "https://testnet.rocketpool.net",
			config.Network_Devnet:  "TBD",
		},

		chainID: map[config.Network]uint{
			config.Network_Mainnet:  1,       // Mainnet
			config.Network_Prater:   5,       // Goerli
			config.Network_Devnet:   5,       // Also goerli
			config.Network_Zhejiang: 1337803, // Zhejiang
		},

		permissionlessNodeRegistryAddress: map[config.Network]string{
			config.Network_Devnet:   "0xeb9e32b2aae13fF98311Bf818C094a4d14FFE694",
			config.Network_Prater:   "0xeb9e32b2aae13fF98311Bf818C094a4d14FFE694",
			config.Network_Mainnet:  "0xeb9e32b2aae13fF98311Bf818C094a4d14FFE694",
			config.Network_Zhejiang: "0x55c1D10b097dAf0E565B6C6D44f9E04ea3EEe2c7",
		},

		vaultFactoryAddress: map[config.Network]string{
			config.Network_Prater:   "0xC60D9b8688F7B12a1ecB5D1349103fBEa7b02906",
			config.Network_Devnet:   "0xC60D9b8688F7B12a1ecB5D1349103fBEa7b02906",
			config.Network_Mainnet:  "0xC60D9b8688F7B12a1ecB5D1349103fBEa7b02906",
			config.Network_Zhejiang: "0xacC1766b4a6dacbB67063a639F588EaB8b6b5A2d",
		},

		sdTokenAddress: map[config.Network]string{
			config.Network_Prater:   "0xD311878a010a94e4500eb5B056DfeaEcAc349AD2",
			config.Network_Devnet:   "0xD311878a010a94e4500eb5B056DfeaEcAc349AD2",
			config.Network_Mainnet:  "0xD311878a010a94e4500eb5B056DfeaEcAc349AD2",
			config.Network_Zhejiang: "0x117f8c20b91e34049798ca66f9234dea4a700f19",
		},

		sdCollateralAddress: map[config.Network]string{
			config.Network_Prater:   "0xb0999B893937DCA77d26bc5923184D0bdD17Fbaf",
			config.Network_Devnet:   "0xb0999B893937DCA77d26bc5923184D0bdD17Fbaf",
			config.Network_Mainnet:  "0xb0999B893937DCA77d26bc5923184D0bdD17Fbaf",
			config.Network_Zhejiang: "0x206fdA2D637C05F69E9d5F0C91a6949F7d0555Fc",
		},

		ethxTokenAddress: map[config.Network]string{
			config.Network_Prater:   "0x9A3f0A872bf39E0cd28a4C25FE66F4B6586953F6",
			config.Network_Devnet:   "0x9A3f0A872bf39E0cd28a4C25FE66F4B6586953F6",
			config.Network_Mainnet:  "0x9A3f0A872bf39E0cd28a4C25FE66F4B6586953F6",
			config.Network_Zhejiang: "0x90Da3CA75532A17ca38440a32595F036ecE46E85",
		},
	}

}

// Get the parameters for this config
func (cfg *SmartnodeConfig) GetParameters() []*config.Parameter {
	return []*config.Parameter{
		&cfg.Network,
		&cfg.ProjectName,
		&cfg.DataPath,
		&cfg.ManualMaxFee,
		&cfg.PriorityFee,
		&cfg.ArchiveECUrl,
	}
}

// Getters for the non-editable parameters

func (cfg *SmartnodeConfig) GetTxWatchUrl() string {
	return cfg.txWatchUrl[cfg.Network.Value.(config.Network)]
}

func (cfg *SmartnodeConfig) GetStakeUrl() string {
	return cfg.stakeUrl[cfg.Network.Value.(config.Network)]
}

func (cfg *SmartnodeConfig) GetChainID() uint {
	return cfg.chainID[cfg.Network.Value.(config.Network)]
}

func (cfg *SmartnodeConfig) GetWalletPath() string {
	if cfg.parent.IsNativeMode {
		return filepath.Join(cfg.DataPath.Value.(string), "wallet")
	}

	return filepath.Join(DaemonDataPath, "wallet")
}

func (cfg *SmartnodeConfig) GetPasswordPath() string {
	if cfg.parent.IsNativeMode {
		return filepath.Join(cfg.DataPath.Value.(string), "password")
	}

	return filepath.Join(DaemonDataPath, "password")
}

func (cfg *SmartnodeConfig) GetValidatorKeychainPath() string {
	if cfg.parent.IsNativeMode {
		return filepath.Join(cfg.DataPath.Value.(string), "validators")
	}

	return filepath.Join(DaemonDataPath, "validators")
}

func (cfg *SmartnodeConfig) GetWalletPathInCLI() string {
	return filepath.Join(cfg.DataPath.Value.(string), "wallet")
}

func (cfg *SmartnodeConfig) GetPasswordPathInCLI() string {
	return filepath.Join(cfg.DataPath.Value.(string), "password")
}

func (cfg *SmartnodeConfig) GetValidatorKeychainPathInCLI() string {
	return filepath.Join(cfg.DataPath.Value.(string), "validators")
}

func (config *SmartnodeConfig) GetWatchtowerStatePath() string {
	if config.parent.IsNativeMode {
		return filepath.Join(config.DataPath.Value.(string), WatchtowerFolder, "state.yml")
	}

	return filepath.Join(DaemonDataPath, WatchtowerFolder, "state.yml")
}

func (cfg *SmartnodeConfig) GetCustomKeyPath() string {
	if cfg.parent.IsNativeMode {
		return filepath.Join(cfg.DataPath.Value.(string), "custom-keys")
	}

	return filepath.Join(DaemonDataPath, "custom-keys")
}

func (cfg *SmartnodeConfig) GetCustomKeyPasswordFilePath() string {
	if cfg.parent.IsNativeMode {
		return filepath.Join(cfg.DataPath.Value.(string), "custom-key-passwords")
	}

	return filepath.Join(DaemonDataPath, "custom-key-passwords")
}

func (cfg *SmartnodeConfig) GetStadernodeContainerTag() string {
	return smartnodeTag
}

func (config *SmartnodeConfig) GetPruneProvisionerContainerTag() string {
	return pruneProvisionerTag
}

func (cfg *SmartnodeConfig) GetEcMigratorContainerTag() string {
	return ecMigratorTag
}

// The the title for the config
func (cfg *SmartnodeConfig) GetConfigTitle() string {
	return cfg.Title
}

func (cfg *SmartnodeConfig) GetPermissionlessNodeRegistryAddress() common.Address {
	return common.HexToAddress(cfg.permissionlessNodeRegistryAddress[cfg.Network.Value.(config.Network)])
}

func (cfg *SmartnodeConfig) GetVaultFactoryAddress() common.Address {
	return common.HexToAddress(cfg.vaultFactoryAddress[cfg.Network.Value.(config.Network)])
}

func (cfg *SmartnodeConfig) GetSdCollateralContractAddress() common.Address {
	return common.HexToAddress(cfg.sdCollateralAddress[cfg.Network.Value.(config.Network)])
}

func (cfg *SmartnodeConfig) GetSdTokenAddress() common.Address {
	return common.HexToAddress(cfg.sdTokenAddress[cfg.Network.Value.(config.Network)])
}

func (cfg *SmartnodeConfig) GetEthxTokenAddress() common.Address {
	return common.HexToAddress(cfg.ethxTokenAddress[cfg.Network.Value.(config.Network)])
}

func getDefaultDataDir(config *StaderConfig) string {
	return filepath.Join(config.StaderDirectory, "data")
}

func (cfg *SmartnodeConfig) GetWatchtowerFolder(daemon bool) string {
	if daemon && !cfg.parent.IsNativeMode {
		return filepath.Join(DaemonDataPath, WatchtowerFolder)
	}

	return filepath.Join(cfg.DataPath.Value.(string), WatchtowerFolder)
}

func (cfg *SmartnodeConfig) GetFeeRecipientFilePath() string {
	if !cfg.parent.IsNativeMode {
		return filepath.Join(DaemonDataPath, "validators", FeeRecipientFilename)
	}

	return filepath.Join(cfg.DataPath.Value.(string), "validators", NativeFeeRecipientFilename)
}

func getNetworkOptions() []config.ParameterOption {
	options := []config.ParameterOption{
		{
			Name:        "Ethereum Mainnet",
			Description: "This is the real Ethereum main network, using real ETH to make real validators.",
			Value:       config.Network_Mainnet,
		}, {
			Name:        "Goerli Testnet",
			Description: "This is the Goerli test network, using Goerli ETH to make demo validators.\nUse this if you want to practice running the Smartnode in a free, safe environment before moving to Mainnet.",
			Value:       config.Network_Prater,
		},
		{
			Name:        "Zhejiang Testnet",
			Description: "This is the Zhejiang test network, using free fake ETH and free fake RPL to make fake validators.\nUse this if you want to test the upcoming Atlas upgrade to Rocket Pool, along with the Shanghai and Capella upgrades to Ethereum that enable validator withdrawals.",
			Value:       config.Network_Zhejiang,
		},
	}

	// if strings.HasSuffix(shared.StaderVersion, "-dev") {
	// 	options = append(options, config.ParameterOption{
	// 		Name:        "Devnet",
	// 		Description: "This is a development network used by Stader engineers to test new features and contract upgrades before they are promoted to Prater for staging. You should not use this network unless invited to do so by the developers.",
	// 		Value:       config.Network_Devnet,
	// 	})
	// }

	return options
}
