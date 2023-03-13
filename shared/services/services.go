package services

import (
	"fmt"
	"math/big"
	"os"
	"sync"

	"github.com/docker/docker/client"
	"github.com/ethereum/go-ethereum/common"
	"github.com/rocket-pool/rocketpool-go/rocketpool"
	"github.com/rocket-pool/rocketpool-go/utils/eth"
	"github.com/stader-labs/stader-node/stader-lib/stader"
	"github.com/urfave/cli"

	"github.com/stader-labs/stader-node/shared/services/beacon"
	"github.com/stader-labs/stader-node/shared/services/config"
	"github.com/stader-labs/stader-node/shared/services/contracts"
	"github.com/stader-labs/stader-node/shared/services/passwords"
	"github.com/stader-labs/stader-node/shared/services/wallet"
	lhkeystore "github.com/stader-labs/stader-node/shared/services/wallet/keystore/lighthouse"
	nmkeystore "github.com/stader-labs/stader-node/shared/services/wallet/keystore/nimbus"
	prkeystore "github.com/stader-labs/stader-node/shared/services/wallet/keystore/prysm"
	tkkeystore "github.com/stader-labs/stader-node/shared/services/wallet/keystore/teku"
	"github.com/stader-labs/stader-node/shared/utils/rp"
)

// Config
const (
	DockerAPIVersion        string = "1.40"
	EcContainerName         string = "eth1"
	FallbackEcContainerName string = "eth1-fallback"
	BnContainerName         string = "eth2"
)

// Service instances & initializers
var (
	cfg                *config.StaderConfig
	passwordManager    *passwords.PasswordManager
	nodeWallet         *wallet.Wallet
	ecManager          *ExecutionClientManager
	bcManager          *BeaconClientManager
	rocketPool         *rocketpool.RocketPool
	oneInchOracle      *contracts.OneInchOracle
	snapshotDelegation *contracts.SnapshotDelegation
	beaconClient       beacon.Client
	docker             *client.Client

	initCfg                sync.Once
	initPasswordManager    sync.Once
	initNodeWallet         sync.Once
	initECManager          sync.Once
	initBCManager          sync.Once
	initRocketPool         sync.Once
	initOneInchOracle      sync.Once
	initRplFaucet          sync.Once
	initSnapshotDelegation sync.Once
	initBeaconClient       sync.Once
	initDocker             sync.Once
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

	return stader.NewPermissionlessNodeRegistry(ec, cfg.Smartnode.GetPermissionlessNodeRegistryAddress())
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

	return stader.NewVaultFactory(ec, cfg.Smartnode.GetVaultFactoryAddress())
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

	return stader.NewSdCollateralContract(ec, cfg.Smartnode.GetSdCollateralContractAddress())
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

	return stader.NewErc20TokenContract(ec, cfg.Smartnode.GetSdTokenAddress())
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

	return stader.NewErc20TokenContract(ec, cfg.Smartnode.GetEthxTokenAddress())
}

func GetRocketPool(c *cli.Context) (*rocketpool.RocketPool, error) {
	cfg, err := getConfig(c)
	if err != nil {
		return nil, err
	}
	ec, err := getEthClient(c, cfg)
	if err != nil {
		return nil, err
	}

	return getRocketPool(cfg, ec)
}

func GetOneInchOracle(c *cli.Context) (*contracts.OneInchOracle, error) {
	cfg, err := getConfig(c)
	if err != nil {
		return nil, err
	}
	ec, err := getEthClient(c, cfg)
	if err != nil {
		return nil, err
	}
	return getOneInchOracle(cfg, ec)
}

func GetSnapshotDelegation(c *cli.Context) (*contracts.SnapshotDelegation, error) {
	cfg, err := getConfig(c)
	if err != nil {
		return nil, err
	}
	ec, err := getEthClient(c, cfg)
	if err != nil {
		return nil, err
	}
	return getSnapshotDelegation(cfg, ec)
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
		cfg, err = rp.LoadConfigFromFile(settingsFile)
		if cfg == nil && err == nil {
			err = fmt.Errorf("Settings file [%s] not found.", settingsFile)
		}
	})
	return cfg, err
}

func getPasswordManager(cfg *config.StaderConfig) *passwords.PasswordManager {
	initPasswordManager.Do(func() {
		passwordManager = passwords.NewPasswordManager(os.ExpandEnv(cfg.Smartnode.GetPasswordPath()))
	})
	return passwordManager
}

func getWallet(c *cli.Context, cfg *config.StaderConfig, pm *passwords.PasswordManager) (*wallet.Wallet, error) {
	var err error
	initNodeWallet.Do(func() {
		var maxFee *big.Int
		maxFeeFloat := c.GlobalFloat64("maxFee")
		if maxFeeFloat == 0 {
			maxFeeFloat = cfg.Smartnode.ManualMaxFee.Value.(float64)
		}
		if maxFeeFloat != 0 {
			maxFee = eth.GweiToWei(maxFeeFloat)
		}

		var maxPriorityFee *big.Int
		maxPriorityFeeFloat := c.GlobalFloat64("maxPrioFee")
		if maxPriorityFeeFloat == 0 {
			maxPriorityFeeFloat = cfg.Smartnode.PriorityFee.Value.(float64)
		}
		if maxPriorityFeeFloat != 0 {
			maxPriorityFee = eth.GweiToWei(maxPriorityFeeFloat)
		}

		chainId := cfg.Smartnode.GetChainID()

		nodeWallet, err = wallet.NewWallet(os.ExpandEnv(cfg.Smartnode.GetWalletPath()), chainId, maxFee, maxPriorityFee, 0, pm)
		if err != nil {
			return
		}

		// Keystores
		lighthouseKeystore := lhkeystore.NewKeystore(os.ExpandEnv(cfg.Smartnode.GetValidatorKeychainPath()), pm)
		nimbusKeystore := nmkeystore.NewKeystore(os.ExpandEnv(cfg.Smartnode.GetValidatorKeychainPath()), pm)
		prysmKeystore := prkeystore.NewKeystore(os.ExpandEnv(cfg.Smartnode.GetValidatorKeychainPath()), pm)
		tekuKeystore := tkkeystore.NewKeystore(os.ExpandEnv(cfg.Smartnode.GetValidatorKeychainPath()), pm)
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

func getRocketPool(cfg *config.StaderConfig, client rocketpool.ExecutionClient) (*rocketpool.RocketPool, error) {
	var err error
	initRocketPool.Do(func() {
		rocketPool, err = rocketpool.NewRocketPool(client, common.HexToAddress(cfg.Smartnode.GetStorageAddress()))
	})
	return rocketPool, err
}

func getOneInchOracle(cfg *config.StaderConfig, client rocketpool.ExecutionClient) (*contracts.OneInchOracle, error) {
	var err error
	initOneInchOracle.Do(func() {
		oneInchOracle, err = contracts.NewOneInchOracle(common.HexToAddress(cfg.Smartnode.GetOneInchOracleAddress()), client)
	})
	return oneInchOracle, err
}

func getSnapshotDelegation(cfg *config.StaderConfig, client rocketpool.ExecutionClient) (*contracts.SnapshotDelegation, error) {
	var err error
	initSnapshotDelegation.Do(func() {
		snapshotDelegation, err = contracts.NewSnapshotDelegation(common.HexToAddress(cfg.Smartnode.GetSnapshotDelegationAddress()), client)
	})
	return snapshotDelegation, err
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
