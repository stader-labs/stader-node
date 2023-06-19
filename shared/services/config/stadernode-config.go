/*
This work is licensed and released under GNU GPL v3 or any other later versions.
The full text of the license is below/ found at <http://www.gnu.org/licenses/>

(c) 2023 Rocket Pool Pty Ltd. Modified under GNU GPL v3. [1.0.0]

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
package config

import (
	"encoding/hex"
	"encoding/json"
	"fmt"
	"math/big"
	"os"
	"path/filepath"

	"github.com/ethereum/go-ethereum/common"
	"github.com/mitchellh/go-homedir"
	"github.com/stader-labs/stader-node/shared"
	"github.com/stader-labs/stader-node/shared/types/config"
	stader_backend "github.com/stader-labs/stader-node/shared/types/stader-backend"
)

// Constants
const (
	stadernodeTag                      = shared.DockerAccount + "/stader-permissionless:v" + shared.StaderVersion
	pruneProvisionerTag                = shared.DockerAccount + "/eth1-prune-provision:v1.0.0"
	ecMigratorTag                      = shared.DockerAccount + "/ec-migrator:v1.2.0"
	NetworkID                   string = "network"
	ProjectNameID               string = "projectName"
	DaemonDataPath              string = "/.stader/data"
	GuardianFolder              string = "guardian"
	SpRewardsMerkleProofsFolder string = "sp-rewards-merkle-proofs"
	MerkleProofsFormat          string = "cycle-%s-%d.json"
	FeeRecipientFilename        string = "stader-fee-recipient.txt"
	NativeFeeRecipientFilename  string = "stader-fee-recipient-env.txt"
)

// --ignore-sync-check
// Defaults
const defaultProjectName string = "stader"

// Configuration for the Stader node
type StaderNodeConfig struct {
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

	// The path of the guardians's persistent state storage
	GuardianStatePath config.Parameter `yaml:"guardianStatePath"`

	// Which network we're on
	Network config.Parameter `yaml:"network,omitempty"`

	// Manual max fee override
	ManualMaxFee config.Parameter `yaml:"manualMaxFee,omitempty"`

	// Manual priority fee override
	PriorityFee config.Parameter `yaml:"priorityFee,omitempty"`

	// Max tx fee for a single tx override
	TxFeeCap config.Parameter `yaml:"txFeeCap,omitempty"`

	// URL for an EC with archive mode, for manual rewards tree generation
	ArchiveECUrl config.Parameter `yaml:"archiveEcUrl,omitempty"`

	///////////////////////////
	// Non-editable settings //
	///////////////////////////

	// The URL to provide the user so they can follow pending transactions
	txWatchUrl map[config.Network]string `yaml:"-"`

	// The URL to use for staking EthX
	stakeUrl map[config.Network]string `yaml:"-"`

	// The URL of the beacon chain explorer
	beaconChainUrl map[config.Network]string `yaml:"-"`

	// The map of networks to execution chain IDs
	chainID map[config.Network]uint `yaml:"-"`

	// The contract address of EthX ERC20
	ethxTokenAddress map[config.Network]string `yaml:"-"`

	// The contract address of stader config
	staderConfigAddress map[config.Network]string `yaml:"-"`

	// The base url of stader backend
	baseStaderBackendUrl map[config.Network]string `yaml:"-"`
}

// Generates a new Stadernode configuration
func NewStadernodeConfig(cfg *StaderConfig) *StaderNodeConfig {

	return &StaderNodeConfig{
		Title:  "Stadernode Settings",
		parent: cfg,

		ProjectName: config.Parameter{
			ID:                   ProjectNameID,
			Name:                 "Project Name",
			Description:          "This is the prefix that will be attached to all of the Docker containers managed by the Stadernode.",
			Type:                 config.ParameterType_String,
			Default:              map[config.Network]interface{}{config.Network_All: defaultProjectName},
			AffectsContainers:    []config.ContainerID{config.ContainerID_Api, config.ContainerID_Node, config.ContainerID_Guardian, config.ContainerID_Eth1, config.ContainerID_Eth2, config.ContainerID_Validator, config.ContainerID_Grafana, config.ContainerID_Prometheus, config.ContainerID_Exporter},
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
			AffectsContainers:    []config.ContainerID{config.ContainerID_Api, config.ContainerID_Node, config.ContainerID_Guardian, config.ContainerID_Validator},
			EnvironmentVariables: []string{"STADER_DATA_FOLDER"},
			CanBeBlank:           false,
			OverwriteOnUpgrade:   false,
		},

		GuardianStatePath: config.Parameter{
			ID:                   "guardianPath",
			Name:                 "Guardian Path",
			Description:          "The absolute path of the guardian state folder that contains persistent state that is used by the guardian process on trusted nodes. **Only relevant for trusted nodes.**",
			Type:                 config.ParameterType_String,
			Default:              map[config.Network]interface{}{config.Network_All: "$HOME/.stader/guardian"},
			AffectsContainers:    []config.ContainerID{config.ContainerID_Guardian},
			EnvironmentVariables: []string{"STADER_GUARDIAN_FOLDER"},
			CanBeBlank:           false,
			OverwriteOnUpgrade:   false,
		},

		Network: config.Parameter{
			ID:                   NetworkID,
			Name:                 "Network",
			Description:          "The Ethereum network you want to use - select Goerli Testnet to practice with Goerli ETH, or Mainnet to stake on the real network using real ETH.",
			Type:                 config.ParameterType_Choice,
			Default:              map[config.Network]interface{}{config.Network_All: config.Network_Mainnet},
			AffectsContainers:    []config.ContainerID{config.ContainerID_Api, config.ContainerID_Node, config.ContainerID_Guardian, config.ContainerID_Eth1, config.ContainerID_Eth2, config.ContainerID_Validator},
			EnvironmentVariables: []string{"NETWORK"},
			CanBeBlank:           false,
			OverwriteOnUpgrade:   false,
			Options:              getNetworkOptions(),
		},

		ManualMaxFee: config.Parameter{
			ID:                   "manualMaxFee",
			Name:                 "Manual Max Fee",
			Description:          "Set this if you want all of the Stadernode's transactions to use this specific max fee value (in gwei), which is the most you'd be willing to pay (*including the priority fee*).\n\nA value of 0 will show you the current suggested max fee based on the current network conditions and let you specify it each time you do a transaction.\n\nAny other value will ignore the recommended max fee and explicitly use this value instead.\n\nThis applies to automated transactions as well.",
			Type:                 config.ParameterType_Float,
			Default:              map[config.Network]interface{}{config.Network_All: float64(0)},
			AffectsContainers:    []config.ContainerID{config.ContainerID_Node, config.ContainerID_Guardian},
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
			AffectsContainers:    []config.ContainerID{config.ContainerID_Node, config.ContainerID_Guardian},
			EnvironmentVariables: []string{},
			CanBeBlank:           false,
			OverwriteOnUpgrade:   false,
		},

		TxFeeCap: config.Parameter{
			ID:                   "txFeeCap",
			Name:                 "Tx Fee Cap",
			Description:          "The max fee which can be paid for a transaction (in eth) including priority fee",
			Type:                 config.ParameterType_Float,
			Default:              map[config.Network]interface{}{config.Network_All: float64(1)},
			AffectsContainers:    []config.ContainerID{config.ContainerID_Node, config.ContainerID_Guardian},
			EnvironmentVariables: []string{"TX_FEE_CAP"},
			CanBeBlank:           false,
			OverwriteOnUpgrade:   false,
		},

		ArchiveECUrl: config.Parameter{
			ID:                   "archiveECUrl",
			Name:                 "Archive-Mode EC URL",
			Description:          "[orange]**For manual Merkle rewards tree generation only.**[white]\n\nGenerating the Merkle rewards tree files for past rewards intervals typically requires an Execution client with Archive mode enabled, which is usually disabled on your primary and fallback Execution clients to save disk space.\nIf you want to generate your own rewards tree files for intervals from a long time ago, you may enter the URL of an Execution client with Archive access here.\n\nFor a free light client with Archive access, you may use https://www.alchemy.com/supernode.",
			Type:                 config.ParameterType_String,
			Default:              map[config.Network]interface{}{config.Network_All: ""},
			AffectsContainers:    []config.ContainerID{config.ContainerID_Guardian},
			EnvironmentVariables: []string{},
			CanBeBlank:           true,
			OverwriteOnUpgrade:   false,
		},

		beaconChainUrl: map[config.Network]string{
			config.Network_Mainnet: "https://beaconcha.in",
			config.Network_Prater:  "https://prater.beaconcha.in",
			config.Network_Devnet:  "https://prater.beaconcha.in",
		},

		txWatchUrl: map[config.Network]string{
			config.Network_Mainnet: "https://etherscan.io/tx",
			config.Network_Prater:  "https://goerli.etherscan.io/tx",
			config.Network_Devnet:  "https://goerli.etherscan.io/tx",
		},

		chainID: map[config.Network]uint{
			config.Network_Mainnet:  1,       // Mainnet
			config.Network_Prater:   5,       // Goerli
			config.Network_Devnet:   5,       // Also goerli
			config.Network_Zhejiang: 1337803, // Zhejiang
		},

		ethxTokenAddress: map[config.Network]string{
			config.Network_Prater:   "0x38DE8Df722B4032Cc6987F00bCA0d9B37d9F9438",
			config.Network_Devnet:   "0x38DE8Df722B4032Cc6987F00bCA0d9B37d9F9438",
			config.Network_Mainnet:  "0xA35b1B31Ce002FBF2058D22F30f95D405200A15b",
			config.Network_Zhejiang: "0x90Da3CA75532A17ca38440a32595F036ecE46E85",
		},

		staderConfigAddress: map[config.Network]string{
			config.Network_Prater:   "0x749Ed651c4F41E0D705960e815A58815ffFd3afe",
			config.Network_Devnet:   "0x749Ed651c4F41E0D705960e815A58815ffFd3afe",
			config.Network_Mainnet:  "0x4ABEF2263d5A5ED582FC9A9789a41D85b68d69DB",
			config.Network_Zhejiang: "0x90Da3CA75532A17ca38440a32595F036ecE46E85",
		},

		baseStaderBackendUrl: map[config.Network]string{
			config.Network_Prater:   "https://stage-ethx-offchain.staderlabs.click",
			config.Network_Devnet:   "https://1r6l0g1nkd.execute-api.us-east-1.amazonaws.com/prod",
			config.Network_Mainnet:  "https://ethx-offchain.staderlabs.com",
			config.Network_Zhejiang: "0x90Da3CA75532A17ca38440a32595F036ecE46E85",
		},
	}
}

// Get the parameters for this config
func (cfg *StaderNodeConfig) GetParameters() []*config.Parameter {
	return []*config.Parameter{
		&cfg.Network,
		&cfg.ProjectName,
		&cfg.DataPath,
		&cfg.ManualMaxFee,
		&cfg.PriorityFee,
		&cfg.TxFeeCap,
		&cfg.ArchiveECUrl,
	}
}

// Getters for the non-editable parameters

func (cfg *StaderNodeConfig) GetBeaconChainUrl() string {
	return cfg.beaconChainUrl[cfg.Network.Value.(config.Network)]
}

// Getters for the non-editable parameters
func (cfg *StaderNodeConfig) GetPresignSendApi() string {
	return cfg.baseStaderBackendUrl[cfg.Network.Value.(config.Network)] + "/presign"
}

func (cfg *StaderNodeConfig) GetBulkPresignSendApi() string {
	return cfg.baseStaderBackendUrl[cfg.Network.Value.(config.Network)] + "/presigns"
}

func (cfg *StaderNodeConfig) GetPresignCheckApi() string {
	return cfg.baseStaderBackendUrl[cfg.Network.Value.(config.Network)] + "/msgSubmitted"
}

func (cfg *StaderNodeConfig) GetBulkPresignCheckApi() string {
	return cfg.baseStaderBackendUrl[cfg.Network.Value.(config.Network)] + "/presignsSubmitted"
}

func (cfg *StaderNodeConfig) GetPresignPublicKeyApi() string {
	return cfg.baseStaderBackendUrl[cfg.Network.Value.(config.Network)] + "/publicKey"
}

func (cfg *StaderNodeConfig) GetMerkleProofApi() string {
	return cfg.baseStaderBackendUrl[cfg.Network.Value.(config.Network)] + "/merklesForElRewards/proofs/%s"
}

func (cfg *StaderNodeConfig) GetTxWatchUrl() string {
	return cfg.txWatchUrl[cfg.Network.Value.(config.Network)]
}

func (cfg *StaderNodeConfig) GetStakeUrl() string {
	return cfg.stakeUrl[cfg.Network.Value.(config.Network)]
}

func (cfg *StaderNodeConfig) GetChainID() uint {
	return cfg.chainID[cfg.Network.Value.(config.Network)]
}

func (cfg *StaderNodeConfig) GetWalletPath() string {
	if cfg.parent.IsNativeMode {
		return filepath.Join(cfg.DataPath.Value.(string), "wallet")
	}

	return filepath.Join(DaemonDataPath, "wallet")
}

func (cfg *StaderNodeConfig) GetPasswordPath() string {
	if cfg.parent.IsNativeMode {
		return filepath.Join(cfg.DataPath.Value.(string), "password")
	}

	return filepath.Join(DaemonDataPath, "password")
}

func (cfg *StaderNodeConfig) GetValidatorKeychainPath() string {
	if cfg.parent.IsNativeMode {
		return filepath.Join(cfg.DataPath.Value.(string), "validators")
	}

	return filepath.Join(DaemonDataPath, "validators")
}

func (cfg *StaderNodeConfig) GetWalletPathInCLI() string {
	return filepath.Join(cfg.DataPath.Value.(string), "wallet")
}

func (cfg *StaderNodeConfig) GetPasswordPathInCLI() string {
	return filepath.Join(cfg.DataPath.Value.(string), "password")
}

func (cfg *StaderNodeConfig) GetValidatorKeychainPathInCLI() string {
	return filepath.Join(cfg.DataPath.Value.(string), "validators")
}

func (config *StaderNodeConfig) GetGuardianStatePath() string {
	if config.parent.IsNativeMode {
		return filepath.Join(config.DataPath.Value.(string), GuardianFolder, "state.yml")
	}

	return filepath.Join(DaemonDataPath, GuardianFolder, "state.yml")
}

func (cfg *StaderNodeConfig) GetCustomKeyPath() string {
	if cfg.parent.IsNativeMode {
		return filepath.Join(cfg.DataPath.Value.(string), "custom-keys")
	}

	return filepath.Join(DaemonDataPath, "custom-keys")
}

func (cfg *StaderNodeConfig) GetCustomKeyPasswordFilePath() string {
	if cfg.parent.IsNativeMode {
		return filepath.Join(cfg.DataPath.Value.(string), "custom-key-passwords")
	}

	return filepath.Join(DaemonDataPath, "custom-key-passwords")
}

func (cfg *StaderNodeConfig) GetStadernodeContainerTag() string {
	return stadernodeTag
}

func (config *StaderNodeConfig) GetPruneProvisionerContainerTag() string {
	return pruneProvisionerTag
}

func (cfg *StaderNodeConfig) GetEcMigratorContainerTag() string {
	return ecMigratorTag
}

// The the title for the config
func (cfg *StaderNodeConfig) GetConfigTitle() string {
	return cfg.Title
}

func (cfg *StaderNodeConfig) GetEthxTokenAddress() common.Address {
	return common.HexToAddress(cfg.ethxTokenAddress[cfg.Network.Value.(config.Network)])
}

func (cfg *StaderNodeConfig) GetStaderConfigAddress() common.Address {
	return common.HexToAddress(cfg.staderConfigAddress[cfg.Network.Value.(config.Network)])
}

func getDefaultDataDir(config *StaderConfig) string {
	return filepath.Join(config.StaderDirectory, "data")
}

func (cfg *StaderNodeConfig) GetGuardianFolder(daemon bool) string {
	if daemon && !cfg.parent.IsNativeMode {
		return filepath.Join(DaemonDataPath, GuardianFolder)
	}

	return filepath.Join(cfg.DataPath.Value.(string), GuardianFolder)
}

func (cfg *StaderNodeConfig) GetSpRewardsMerkleProofFolder(daemon bool) string {
	if daemon && !cfg.parent.IsNativeMode {
		return filepath.Join(DaemonDataPath, SpRewardsMerkleProofsFolder)
	}

	return filepath.Join(cfg.DataPath.Value.(string), GuardianFolder)
}

func (cfg *StaderNodeConfig) GetSpRewardCyclePath(cycle int64, daemon bool) string {
	if daemon && !cfg.parent.IsNativeMode {
		return filepath.Join(DaemonDataPath, SpRewardsMerkleProofsFolder, fmt.Sprintf(MerkleProofsFormat, string(cfg.Network.Value.(config.Network)), cycle))
	}

	return filepath.Join(cfg.DataPath.Value.(string), SpRewardsMerkleProofsFolder, fmt.Sprintf(MerkleProofsFormat, string(cfg.Network.Value.(config.Network)), cycle))
}

func (cfg *StaderNodeConfig) GetFeeRecipientFilePath() string {
	if !cfg.parent.IsNativeMode {
		return filepath.Join(DaemonDataPath, "validators", FeeRecipientFilename)
	}

	return filepath.Join(cfg.DataPath.Value.(string), "validators", NativeFeeRecipientFilename)
}

func (cfg *StaderNodeConfig) GetClaimData(cycles []*big.Int) ([]*big.Int, []*big.Int, [][][32]byte, error) {
	// data to pass to socializing pool contract
	amountSd := []*big.Int{}
	amountEth := []*big.Int{}
	merkleProofs := [][][32]byte{}

	// get the proofs for each cycle and claim them. throw error if cycle proofs is not downloaded
	for _, cycle := range cycles {
		cycleMerkleRewardFile := cfg.GetSpRewardCyclePath(cycle.Int64(), true)
		expandedCycleMerkleRewardFile, err := homedir.Expand(cycleMerkleRewardFile)
		if err != nil {
			return nil, nil, nil, err
		}
		data, err := os.ReadFile(expandedCycleMerkleRewardFile)
		if err != nil {
			return nil, nil, nil, err
		}
		// need to make sure we have downloaded the data
		if os.IsNotExist(err) {
			return nil, nil, nil, err
		}
		merkleData := stader_backend.CycleMerkleProofs{}
		err = json.Unmarshal(data, &merkleData)
		if err != nil {
			return nil, nil, nil, err
		}

		amountSdBigInt := big.NewInt(0)
		amountSdBigInt, ok := amountSdBigInt.SetString(merkleData.Sd, 10)
		if !ok {
			return nil, nil, nil, fmt.Errorf("could not parse sd amount %s", merkleData.Sd)
		}

		// same thing above for eth
		amountEthBigInt := big.NewInt(0)
		amountEthBigInt, ok = amountEthBigInt.SetString(merkleData.Eth, 10)
		if !ok {
			return nil, nil, nil, fmt.Errorf("could not parse eth amount %s", merkleData.Eth)
		}

		amountSd = append(amountSd, amountSdBigInt)
		amountEth = append(amountEth, amountEthBigInt)

		// convert merkle proofs to [32]byte
		cycleMerkleProofs := [][32]byte{}
		for _, proof := range merkleData.Proof {
			merkleProofBytes, err := hex.DecodeString(proof[2:])
			if err != nil {
				return nil, nil, nil, err
			}
			var proofBytes [32]byte
			copy(proofBytes[:], merkleProofBytes[:32])
			cycleMerkleProofs = append(cycleMerkleProofs, proofBytes)
		}

		merkleProofs = append(merkleProofs, cycleMerkleProofs)
	}

	return amountSd, amountEth, merkleProofs, nil
}

func (cfg *StaderNodeConfig) ReadCycleCache(cycle int64) (stader_backend.CycleMerkleProofs, bool, error) {
	//fmt.Printf("Reading cycle cache for cycle %d\n", cycle)
	cycleMerkleProofFile := cfg.GetSpRewardCyclePath(cycle, true)
	absolutePathOfProofFile, err := homedir.Expand(cycleMerkleProofFile)
	if err != nil {
		return stader_backend.CycleMerkleProofs{}, false, err
	}

	_, err = os.Stat(cycleMerkleProofFile)
	if !os.IsNotExist(err) && err != nil {
		return stader_backend.CycleMerkleProofs{}, false, err
	}
	if os.IsNotExist(err) {
		return stader_backend.CycleMerkleProofs{}, false, nil
	}

	// Open the JSON file
	file, err := os.Open(absolutePathOfProofFile)
	if err != nil {
		return stader_backend.CycleMerkleProofs{}, false, err
	}

	var cycleMerkleProof stader_backend.CycleMerkleProofs
	decoder := json.NewDecoder(file)
	err = decoder.Decode(&cycleMerkleProof)
	if err != nil {
		return stader_backend.CycleMerkleProofs{}, false, err
	}

	return cycleMerkleProof, true, nil
}

func getNetworkOptions() []config.ParameterOption {
	options := []config.ParameterOption{
		{
			Name:        "Ethereum Mainnet",
			Description: "This is the real Ethereum main network, using real ETH to make real validators.",
			Value:       config.Network_Mainnet,
		}, {
			Name:        "Goerli Testnet",
			Description: "This is the Goerli test network, using Goerli ETH to make demo validators.\nUse this if you want to practice running the Stadernode in a free, safe environment before moving to Mainnet.",
			Value:       config.Network_Prater,
		},
		{
			Name:        "Zhejiang Testnet",
			Description: "This is the Zhejiang test network, using free fake ETH and free fake SD to make fake validators.\nUse this if you want to test the ZHejiang network, along with the Shanghai and Capella upgrades to Ethereum that enable validator withdrawals.",
			Value:       config.Network_Zhejiang,
		},
	}

	// if string-utils.HasSuffix(shared.StaderVersion, "-dev") {
	// 	options = append(options, config.ParameterOption{
	// 		Name:        "Devnet",
	// 		Description: "This is a development network used by Stader engineers to test new features and contract upgrades before they are promoted to Prater for staging. You should not use this network unless invited to do so by the developers.",
	// 		Value:       config.Network_Devnet,
	// 	})
	// }

	return options
}
