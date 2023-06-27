package testing

import (
	"context"
	"fmt"
	"os"
	"path/filepath"
	"time"

	"github.com/kurtosis-tech/kurtosis/api/golang/engine/lib/kurtosis_context"
	"github.com/mitchellh/go-homedir"
	"github.com/stader-labs/stader-node/shared/services"
	"github.com/stader-labs/stader-node/shared/services/passwords"
	"github.com/stader-labs/stader-node/shared/services/stader"
	"github.com/stader-labs/stader-node/shared/services/wallet"
	cfgtypes "github.com/stader-labs/stader-node/shared/types/config"
	"github.com/stader-labs/stader-node/stader/api"
	"github.com/stader-labs/stader-node/stader/node"
	"github.com/stretchr/testify/require"
	"github.com/urfave/cli"

	"github.com/stader-labs/stader-node/shared/services/config"
)

var ( //f02daebbf456faf787c5cd61a33ce780857c1ca10b00972aa451f0e9688e4ead
	preFundedKeyKur   = "ef5177cd0b6b21c87db5a0bf35d4084a8a57a9d6a064f86d51ac85f2b873a4e2"
	preFundedKeyAnvil = "2a871d0798f97d79848a013d4936a73bf4cc922c825d33c1cf7073dff6d409c6"
	UserSettingPath   = filepath.Join(ConfigPath, "user-settings.yml")
	ConfigPath, _     = homedir.Expand("~/.stader_testing")
	PasswordPath      = filepath.Join(ConfigPath, "password")
	cf                = []byte(`{
		"participants": [
			{
				"el_client_type": "geth",
				"el_client_image": "",
				"el_client_log_level": "",
				"cl_client_type": "lighthouse",
				"cl_client_image": "",
				"cl_client_log_level": "",
				"beacon_extra_params": [],
				"el_extra_params": ["--http"],
				"validator_extra_params": [],
				"builder_network_params": null
			}
		],
		"network_params": {
			"preregistered_validator_keys_mnemonic": "giant issue aisle success illegal bike spike question tent bar rely arctic volcano long crawl hungry vocal artwork sniff fantasy very lucky have athlete",
			"num_validator_keys_per_node": 64,
			"network_id": "3151908",
			"deposit_contract_address": "0x4242424242424242424242424242424242424242",
			"seconds_per_slot": 12,
			"genesis_delay": 120,
			"capella_fork_epoch": 5
		},
		"global_client_log_level": "info"
	}
	
	`)
)

const (
	enclaveIdPrefix = "stader"

	remotePackage = "github.com/kurtosis-tech/eth-network-package"

	noDryRun = false

	emptyPackageParams = "{}"
	clClientBeacon     = "cl-client-0-beacon"
	clClientValidator  = "cl-client-0-validator"
	elCient            = "el-client-0"

	contentType = "application/json"

	useDefaultMainFile     = ""
	useDefaultFunctionName = ""

	emptyParams           = "{}"
	isPartitioningEnabled = false
	emptyRunParams        = "{}"
	defaultDryRun         = false
	defaultParallelism    = 4
)

func newApp() *cli.App {

	app := cli.NewApp()

	// Set application flags
	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:  "settings",
			Usage: "Stader service user config absolute `path`",
			Value: UserSettingPath,
		},
		cli.StringFlag{
			Name:  "password, p",
			Usage: "Stader wallet password file absolute `path`",
		},
		cli.StringFlag{
			Name:  "wallet, w",
			Usage: "Stader wallet file absolute `path`",
		},
		cli.StringFlag{
			Name:  "validatorKeychain, k",
			Usage: "Stader validator keychain absolute `path`",
		},
		cli.StringFlag{
			Name:  "eth1Provider, e",
			Usage: "Eth 1.0 provider `address`",
		},
		cli.StringFlag{
			Name:  "eth2Provider, b",
			Usage: "Eth 2.0 provider `address`",
		},
		cli.Float64Flag{
			Name:  "maxFee",
			Usage: "Desired max fee in gwei",
		},
		cli.Float64Flag{
			Name:  "maxPrioFee",
			Usage: "Desired max priority fee in gwei",
		},
		cli.Uint64Flag{
			Name:  "gasLimit, l",
			Usage: "Desired gas limit",
		},
		cli.StringFlag{
			Name:  "nonce",
			Usage: "Use this flag to explicitly specify the nonce that this transaction should use, so it can override an existing 'stuck' transaction",
		},
		cli.StringFlag{
			Name:  "metricsAddress, m",
			Usage: "Address to serve metrics on if enabled",
			Value: "0.0.0.0",
		},
		cli.UintFlag{
			Name:  "metricsPort, r",
			Usage: "Port to serve metrics on if enabled",
			Value: 9102,
		},
		cli.BoolFlag{
			Name:  "ignore-sync-check",
			Usage: "Set this to true if you already checked the sync status of the execution client(s) and don't need to re-check it for this command",
		},
		cli.BoolFlag{
			Name:  "force-fallbacks",
			Usage: "Set this to true if you know the primary EC or CC is offline and want to bypass its health checks, and just use the fallback EC and CC instead",
		},
		cli.BoolFlag{
			Name: "local-testnet",
		},
		cli.StringFlag{
			Name: "presign-cooldown",
		},
	}

	// Register commands
	api.RegisterCommands(app, "api", []string{"a"})
	node.RegisterCommands(app, "node", []string{"n"})
	return app
}
func (s *StaderNodeSuite) setConfig(c *cli.Context, elURL string, clURL string) {
	cfg := config.NewStaderConfig(UserSettingPath, false)

	staderClient, err := stader.NewClientFromCtx(c)
	staderClient.GetContractsInfo()
	require.Nil(s.T(), err)
	defer staderClient.Close()

	cfg.ExecutionClientMode.Value = cfgtypes.Mode_External
	cfg.ExternalExecution.HttpUrl.Value = elURL

	cfg.ExternalConsensusClient.Value = cfgtypes.ConsensusClient_Lighthouse
	cfg.ConsensusClientMode.Value = cfgtypes.Mode_External
	cfg.ExternalLighthouse.HttpUrl.Value = clURL

	cfg.StaderNode.DataPath.Value = ConfigPath

	cfg.Native.EcHttpUrl.Value = elURL
	cfg.Native.ConsensusClient.Value = clURL
	cfg.Native.CcHttpUrl.Value = clURL

	cfg.ChangeNetwork(cfgtypes.Network_Local)
	path, err := homedir.Expand(ConfigPath)
	require.Nil(s.T(), err)

	_ = os.Mkdir(path, os.ModePerm)

	err = staderClient.SaveConfig(cfg)
	require.Nil(s.T(), err)

	bc, err := services.GetBeaconClient(c)
	require.Nil(s.T(), err)
	s.bc = bc
}

func (s *StaderNodeSuite) staderConfig(
	ctx context.Context,
	c *cli.Context,
	clUrl *string,
	elUrl string,
) {

	t := s.T()
	if clUrl == nil {
		fmt.Println("------------ CONNECTING TO KURTOSIS ENGINE ---------------")
		kurtosis_context.NewKurtosisContextFromLocalEngine()
		kurtosisCtx, err := kurtosis_context.NewKurtosisContextFromLocalEngine()
		require.NoError(t, err, "An error occurred connecting to the Kurtosis engine")

		enclaveId := fmt.Sprintf("%s-%d", enclaveIdPrefix, time.Now().Unix())

		enclaveCtx, err := kurtosisCtx.CreateEnclave(ctx, enclaveId, isPartitioningEnabled)

		s.kurtosisCtx = kurtosisCtx
		s.enclaveId = enclaveId
		require.NoError(t, err, "An error occurred creating the enclave")

		fmt.Println("------------ EXECUTING PACKAGE ---------------")

		starlarkRunResult, err := enclaveCtx.RunStarlarkRemotePackageBlocking(ctx, remotePackage, useDefaultMainFile, useDefaultFunctionName, emptyParams, defaultDryRun, defaultParallelism)

		require.NoError(t, err, "An error executing loading the package")
		require.Nil(t, starlarkRunResult.InterpretationError)
		require.Empty(t, starlarkRunResult.ValidationErrors)
		require.Nil(t, starlarkRunResult.ExecutionError)

		fmt.Println("------------ EXECUTING TESTS ---------------")
		beaconContext, err := enclaveCtx.GetServiceContext(clClientBeacon)
		require.Nil(t, err)
		apiServicePublicPorts := beaconContext.GetPublicPorts()
		require.NotNil(t, apiServicePublicPorts)
		apiServiceHttpPortSpec, found := apiServicePublicPorts["http"]
		require.True(t, found)
		clPort := apiServiceHttpPortSpec.GetNumber()

		// elContext, err := enclaveCtx.GetServiceContext(elCient)
		// require.Nil(t, err)
		// elPublicPorts := elContext.GetPublicPorts()
		// require.NotNil(t, apiServicePublicPorts)
		// apiServiceHttpPortSpec, found = elPublicPorts["rpc"]
		// require.True(t, found)
		// elPort := apiServiceHttpPortSpec.GetNumber()

		_clUrl := fmt.Sprintf("http://127.0.0.1:%d", clPort)
		clUrl = &_clUrl
	}

	s.setConfig(c, elUrl, *clUrl)
	s.setupWallet(ctx, c)

	fmt.Println("------------ DEPLOYING CONTRACT ---------------")
	fmt.Println("CURL: ", *clUrl)
	fmt.Println("EURL: ", elUrl)

	deployContracts(s.T(), c, elUrl)

}

func (s *StaderNodeSuite) setupWallet(ctx context.Context, c *cli.Context) {

	// Get services
	pm := passwords.NewPasswordManager(PasswordPath)

	// Set password
	err := pm.SetPassword("stader_testing_pass")
	require.Nil(s.T(), err)

	w, err := services.GetWallet(c)
	require.Nil(s.T(), err)

	_, err = w.Initialize(wallet.DefaultNodeKeyPath, 0)

	require.Nil(s.T(), err)

	err = w.Save()
	require.Nil(s.T(), err)
}
