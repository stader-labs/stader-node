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
package services

import (
	"fmt"
	"math/big"
	"os"
	"sync"

	"github.com/docker/docker/client"
	"github.com/stader-labs/stader-node/stader-lib/stader"
	"github.com/stader-labs/stader-node/stader-lib/utils/eth"
	"github.com/urfave/cli"

	"github.com/stader-labs/stader-node/shared/services/config"
	"github.com/stader-labs/stader-node/shared/services/passwords"
	"github.com/stader-labs/stader-node/shared/services/wallet"
	lhkeystore "github.com/stader-labs/stader-node/shared/services/wallet/keystore/lighthouse"
	nmkeystore "github.com/stader-labs/stader-node/shared/services/wallet/keystore/nimbus"
	prkeystore "github.com/stader-labs/stader-node/shared/services/wallet/keystore/prysm"
	tkkeystore "github.com/stader-labs/stader-node/shared/services/wallet/keystore/teku"
	staderUtils "github.com/stader-labs/stader-node/shared/utils/stdr"
)

// Config
const (
	DockerAPIVersion string = "1.40"
	BnContainerName  string = "eth2"
)

// Service instances & initializers
var (
	cfg             *config.StaderConfig
	passwordManager *passwords.PasswordManager
	nodeWallet      *wallet.Wallet
	ecManager       *ExecutionClientManager
	bcManager       *BeaconClientManager
	docker          *client.Client

	initCfg             sync.Once
	initPasswordManager sync.Once
	initNodeWallet      sync.Once
	initECManager       sync.Once
	initBCManager       sync.Once
	initDocker          sync.Once
)

//
// Service providers
//

func GetConfig(c *cli.Context) (*config.StaderConfig, error) {
	return getConfig(c)
}

func GetPasswordManager(c *cli.Context) (*passwords.PasswordManager, error) {
	cfg, err := getConfig(c)
	if err != nil {
		return nil, err
	}
	return getPasswordManager(cfg), nil
}

func GetWallet(c *cli.Context) (*wallet.Wallet, error) {
	cfg, err := getConfig(c)
	if err != nil {
		return nil, err
	}
	pm := getPasswordManager(cfg)
	return getWallet(c, cfg, pm)
}

func GetEthClient(c *cli.Context) (*ExecutionClientManager, error) {
	cfg, err := getConfig(c)
	if err != nil {
		return nil, err
	}
	ec, err := getEthClient(c, cfg)
	if err != nil {
		return nil, err
	}
	return ec, nil
}

func GetPermissionlessNodeRegistry(c *cli.Context) (*stader.PermissionlessNodeRegistryContractManager, error) {
	cfg, err := getConfig(c)
	if err != nil {
		return nil, err
	}
	ec, err := getEthClient(c, cfg)
	if err != nil {
		return nil, err
	}

	return stader.NewPermissionlessNodeRegistry(ec, cfg.StaderNode.GetPermissionlessNodeRegistryAddress())
}

func GetVaultFactory(c *cli.Context) (*stader.VaultFactoryContractManager, error) {
	cfg, err := getConfig(c)
	if err != nil {
		return nil, err
	}
	ec, err := getEthClient(c, cfg)
	if err != nil {
		return nil, err
	}

	return stader.NewVaultFactory(ec, cfg.StaderNode.GetVaultFactoryAddress())
}

func GetPermissionlessPoolContract(c *cli.Context) (*stader.PermissionlessPoolContractManager, error) {
	cfg, err := getConfig(c)
	if err != nil {
		return nil, err
	}
	ec, err := getEthClient(c, cfg)
	if err != nil {
		return nil, err
	}

	return stader.NewPermissionlessPoolFactory(ec, cfg.StaderNode.GetPermissionlessPoolAddress())
}

func GetSdCollateralContract(c *cli.Context) (*stader.SdCollateralContractManager, error) {
	cfg, err := getConfig(c)
	if err != nil {
		return nil, err
	}
	ec, err := getEthClient(c, cfg)
	if err != nil {
		return nil, err
	}

	return stader.NewSdCollateralContract(ec, cfg.StaderNode.GetSdCollateralContractAddress())
}

func GetStaderConfigContract(c *cli.Context) (*stader.StaderConfigContractManager, error) {
	cfg, err := getConfig(c)
	if err != nil {
		return nil, err
	}
	ec, err := getEthClient(c, cfg)
	if err != nil {
		return nil, err
	}

	return stader.NewStaderConfig(ec, cfg.StaderNode.GetStaderConfigAddress())
}

func GetSocializingPoolContract(c *cli.Context) (*stader.SocializingPoolContractManager, error) {
	cfg, err := getConfig(c)
	if err != nil {
		return nil, err
	}
	ec, err := getEthClient(c, cfg)
	if err != nil {
		return nil, err
	}

	return stader.NewSocializingPool(ec, cfg.StaderNode.GetSocializingPoolAddress())
}

func GetSdTokenContract(c *cli.Context) (*stader.Erc20TokenContractManager, error) {
	cfg, err := getConfig(c)
	if err != nil {
		return nil, err
	}
	ec, err := getEthClient(c, cfg)
	if err != nil {
		return nil, err
	}

	return stader.NewErc20TokenContract(ec, cfg.StaderNode.GetSdTokenAddress())
}

func GetEthxTokenContract(c *cli.Context) (*stader.Erc20TokenContractManager, error) {
	cfg, err := getConfig(c)
	if err != nil {
		return nil, err
	}
	ec, err := getEthClient(c, cfg)
	if err != nil {
		return nil, err
	}

	return stader.NewErc20TokenContract(ec, cfg.StaderNode.GetEthxTokenAddress())
}

func GetPoolUtilsContract(c *cli.Context) (*stader.PoolUtilsContractManager, error) {
	cfg, err := getConfig(c)
	if err != nil {
		return nil, err
	}
	ec, err := getEthClient(c, cfg)
	if err != nil {
		return nil, err
	}

	return stader.NewPoolUtils(ec, cfg.StaderNode.GetPoolUtilsAddress())
}

func GetPenaltyTrackerContract(c *cli.Context) (*stader.PenaltyTrackerContractManager, error) {
	cfg, err := getConfig(c)
	if err != nil {
		return nil, err
	}
	ec, err := getEthClient(c, cfg)
	if err != nil {
		return nil, err
	}

	return stader.NewPenaltyTracker(ec, cfg.StaderNode.GetPenaltyTrackerAddress())
}

func GetStakePoolManager(c *cli.Context) (*stader.StakePoolManagerContractManager, error) {
	cfg, err := getConfig(c)
	if err != nil {
		return nil, err
	}
	ec, err := getEthClient(c, cfg)
	if err != nil {
		return nil, err
	}
	return stader.NewStakePoolManager(ec, cfg.StaderNode.GetStakePoolManagerAddress())
}

func GetBeaconClient(c *cli.Context) (*BeaconClientManager, error) {
	cfg, err := getConfig(c)
	if err != nil {
		return nil, err
	}
	return getBeaconClient(c, cfg)
}

func GetDocker(c *cli.Context) (*client.Client, error) {
	return getDocker()
}

//
// Service instance getters
//

func getConfig(c *cli.Context) (*config.StaderConfig, error) {
	var err error
	initCfg.Do(func() {
		settingsFile := os.ExpandEnv(c.GlobalString("settings"))
		cfg, err = staderUtils.LoadConfigFromFile(settingsFile)
		if cfg == nil && err == nil {
			err = fmt.Errorf("Settings file [%s] not found.", settingsFile)
		}
	})
	return cfg, err
}

func getPasswordManager(cfg *config.StaderConfig) *passwords.PasswordManager {
	initPasswordManager.Do(func() {
		passwordManager = passwords.NewPasswordManager(os.ExpandEnv(cfg.StaderNode.GetPasswordPath()))
	})
	return passwordManager
}

func getWallet(c *cli.Context, cfg *config.StaderConfig, pm *passwords.PasswordManager) (*wallet.Wallet, error) {
	var err error
	initNodeWallet.Do(func() {
		var maxFee *big.Int
		maxFeeFloat := c.GlobalFloat64("maxFee")
		if maxFeeFloat == 0 {
			maxFeeFloat = cfg.StaderNode.ManualMaxFee.Value.(float64)
		}
		if maxFeeFloat != 0 {
			maxFee = eth.GweiToWei(maxFeeFloat)
		}

		var maxPriorityFee *big.Int
		maxPriorityFeeFloat := c.GlobalFloat64("maxPrioFee")
		if maxPriorityFeeFloat == 0 {
			maxPriorityFeeFloat = cfg.StaderNode.PriorityFee.Value.(float64)
		}
		if maxPriorityFeeFloat != 0 {
			maxPriorityFee = eth.GweiToWei(maxPriorityFeeFloat)
		}

		chainId := cfg.StaderNode.GetChainID()

		nodeWallet, err = wallet.NewWallet(os.ExpandEnv(cfg.StaderNode.GetWalletPath()), chainId, maxFee, maxPriorityFee, 0, pm)
		if err != nil {
			return
		}

		// Keystores
		lighthouseKeystore := lhkeystore.NewKeystore(os.ExpandEnv(cfg.StaderNode.GetValidatorKeychainPath()), pm)
		nimbusKeystore := nmkeystore.NewKeystore(os.ExpandEnv(cfg.StaderNode.GetValidatorKeychainPath()), pm)
		prysmKeystore := prkeystore.NewKeystore(os.ExpandEnv(cfg.StaderNode.GetValidatorKeychainPath()), pm)
		tekuKeystore := tkkeystore.NewKeystore(os.ExpandEnv(cfg.StaderNode.GetValidatorKeychainPath()), pm)
		nodeWallet.AddKeystore("lighthouse", lighthouseKeystore)
		nodeWallet.AddKeystore("nimbus", nimbusKeystore)
		nodeWallet.AddKeystore("prysm", prysmKeystore)
		nodeWallet.AddKeystore("teku", tekuKeystore)
	})
	return nodeWallet, err
}

func getEthClient(c *cli.Context, cfg *config.StaderConfig) (*ExecutionClientManager, error) {
	var err error
	initECManager.Do(func() {
		// Create a new client manager
		ecManager, err = NewExecutionClientManager(cfg)
		if err == nil {
			// Check if the manager should ignore sync checks and/or default to using the fallback (used by the API container when driven by the CLI)
			if c.GlobalBool("ignore-sync-check") {
				ecManager.ignoreSyncCheck = true
			}
			if c.GlobalBool("force-fallbacks") {
				ecManager.primaryReady = false
			}
		}
	})
	return ecManager, err
}

func getBeaconClient(c *cli.Context, cfg *config.StaderConfig) (*BeaconClientManager, error) {
	var err error
	initBCManager.Do(func() {
		// Create a new client manager
		bcManager, err = NewBeaconClientManager(cfg)
		if err == nil {
			// Check if the manager should ignore sync checks and/or default to using the fallback (used by the API container when driven by the CLI)
			if c.GlobalBool("ignore-sync-check") {
				bcManager.ignoreSyncCheck = true
			}
			if c.GlobalBool("force-fallbacks") {
				bcManager.primaryReady = false
			}
		}
	})
	return bcManager, err
}

func getDocker() (*client.Client, error) {
	var err error
	initDocker.Do(func() {
		docker, err = client.NewClientWithOpts(client.WithVersion(DockerAPIVersion))
	})
	return docker, err
}
