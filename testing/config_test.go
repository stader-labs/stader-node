package testing

import (
	"context"
	"fmt"
	"time"

	"github.com/kurtosis-tech/kurtosis/api/golang/engine/lib/kurtosis_context"
	"github.com/sirupsen/logrus"
	"github.com/stader-labs/stader-node/shared/services"
	"github.com/stader-labs/stader-node/shared/services/stader"
	cfgtypes "github.com/stader-labs/stader-node/shared/types/config"
	"github.com/stader-labs/stader-node/stader/api"
	"github.com/stader-labs/stader-node/stader/node"
	"github.com/stretchr/testify/assert"
	"github.com/urfave/cli"

	"github.com/stader-labs/stader-node/shared/services/config"
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

var (
	ConfigPath = "/Users/batphonghan/.stader_testing/user-settings.yml"
)
var cf = []byte(`{
	"participants": [
		{
			"el_client_type": "geth",
			"el_client_image": "ethereum/client-go:v1.11.5",
			"el_client_log_level": "",
			"cl_client_type": "lighthouse",
			"cl_client_image": "sigp/lighthouse:v3.5.0",
			"cl_client_log_level": "",
			"beacon_extra_params": [],
			"el_extra_params": ["--http"],
			"validator_extra_params": [],
			"builder_network_params": null
		}
	],
	"network_params": {
		"preregistered_validator_keys_mnemonic": "giant issue aisle success illegal bike spike question tent bar rely arctic volcano long crawl hungry vocal artwork sniff fantasy very lucky have athlete",
        "num_validator_keys_per_node": 40,
        "network_id": "3151908",
        "deposit_contract_address": "0x4242424242424242424242424242424242424242",
        "seconds_per_slot": 2,
        "genesis_delay": 120,
        "capella_fork_epoch": 5
	}
}`)

func newApp() *cli.App {

	app := cli.NewApp()

	// Set application flags
	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:  "settings",
			Usage: "Stader service user config absolute `path`",
			Value: ConfigPath,
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
	}

	// Register commands
	api.RegisterCommands(app, "api", []string{"a"})
	node.RegisterCommands(app, "node", []string{"n"})
	return app
}
func (s *StaderNodeSuite) setConfig(c *cli.Context, elURL string, clURL string) {
	cfg := config.NewStaderConfig(ConfigPath, false)

	staderClient, err := stader.NewClientFromCtx(c)
	assert.Nil(s.T(), err)
	defer staderClient.Close()

	cfg.ExecutionClientMode.Value = cfgtypes.Mode_External
	cfg.ExternalExecution.HttpUrl.Value = elURL

	cfg.ExternalConsensusClient.Value = cfgtypes.ConsensusClient_Lighthouse
	cfg.ConsensusClientMode.Value = cfgtypes.Mode_External
	cfg.ExternalLighthouse.HttpUrl.Value = clURL

	err = staderClient.SaveConfig(cfg)
	assert.Nil(s.T(), err)
}

func (s *StaderNodeSuite) staderConfig(ctx context.Context, c *cli.Context) {

	t := s.T()

	logrus.Info("------------ CONNECTING TO KURTOSIS ENGINE ---------------")
	kurtosis_context.NewKurtosisContextFromLocalEngine()
	kurtosisCtx, err := kurtosis_context.NewKurtosisContextFromLocalEngine()
	assert.NoError(t, err, "An error occurred connecting to the Kurtosis engine")

	enclaveId := fmt.Sprintf("%s-%d", enclaveIdPrefix, time.Now().Unix())

	enclaveCtx, err := kurtosisCtx.CreateEnclave(ctx, enclaveId, isPartitioningEnabled)

	s.kurtosisCtx = kurtosisCtx
	s.enclaveId = enclaveId
	assert.NoError(t, err, "An error occurred creating the enclave")

	logrus.Info("------------ EXECUTING PACKAGE ---------------")

	starlarkRunResult, err := enclaveCtx.RunStarlarkRemotePackageBlocking(ctx, remotePackage, useDefaultMainFile, useDefaultFunctionName, emptyParams, defaultDryRun, defaultParallelism)

	assert.NoError(t, err, "An error executing loading the package")
	assert.Nil(t, starlarkRunResult.InterpretationError)
	assert.Empty(t, starlarkRunResult.ValidationErrors)
	assert.Nil(t, starlarkRunResult.ExecutionError)

	logrus.Info("------------ EXECUTING TESTS ---------------")
	beaconContext, err := enclaveCtx.GetServiceContext(clClientBeacon)
	assert.Nil(t, err)
	apiServicePublicPorts := beaconContext.GetPublicPorts()
	assert.NotNil(t, apiServicePublicPorts)
	apiServiceHttpPortSpec, found := apiServicePublicPorts["http"]
	assert.True(t, found)
	beaconchainPort := apiServiceHttpPortSpec.GetNumber()

	fmt.Printf("%+v", beaconchainPort)

	validatorContext, err := enclaveCtx.GetServiceContext(clClientValidator)
	assert.Nil(t, err)
	apiServicePublicPorts = validatorContext.GetPublicPorts()
	assert.NotNil(t, apiServicePublicPorts)
	apiServiceHttpPortSpec, found = apiServicePublicPorts["http"]
	assert.True(t, found)
	beaconchainPort = apiServiceHttpPortSpec.GetNumber()

	fmt.Printf("%+v", beaconchainPort)

	elContext, err := enclaveCtx.GetServiceContext(elCient)
	assert.Nil(t, err)
	elPublicPorts := elContext.GetPublicPorts()
	assert.NotNil(t, apiServicePublicPorts)
	apiServiceHttpPortSpec, found = elPublicPorts["rpc"]
	assert.True(t, found)
	elPort := apiServiceHttpPortSpec.GetNumber()

	elUrl := fmt.Sprintf("http://127.0.0.1:%+v", elPort)

	s.setConfig(c, fmt.Sprintf("http://127.0.0.1:%+v", beaconchainPort), elUrl)

	deployContracts(elUrl)

	_, err = services.GetWallet(c)

	assert.Nil(s.T(), err)

}

/*
 Api contract from 0x878705ba3f8Bc32FCf7F4CAa1A35E72AF65CF766
Api contract deployed to 0xAb2A01BC351770D09611Ac80f1DE076D56E0487d
Tx: 0xcc52c30d4d4ea67654b23beda69696485ca60577db8c2482adad82eedd199bf5
*/
