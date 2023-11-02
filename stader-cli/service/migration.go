package service

import (
	_ "embed"
	"fmt"
	"os"

	"github.com/hashicorp/go-version"
	"github.com/mitchellh/go-homedir"
	"github.com/stader-labs/stader-node/shared"
	"github.com/stader-labs/stader-node/shared/services/stader"
	"github.com/urfave/cli"
)

const RefreshingCycle = 5

type ConfigUpgrader struct {
	version                 *version.Version
	upgradeFunc             func(c *cli.Context) error
	needInstall             bool
	runAfterStaderNodeStart bool
}

//go:embed guardian.tmpl
var guardian []byte

func migrate(c *cli.Context) (runBeforeUpgrades, rundAfterUpgrades []ConfigUpgrader, err error) {
	v0, _ := parseVersion("1.0.0")

	// Create versions
	v130, err := parseVersion("1.3.0")
	if err != nil {
		return nil, nil, err
	}

	v140, err := parseVersion("1.4.0")
	if err != nil {
		return nil, nil, err
	}

	v142, err := parseVersion("1.4.2")
	if err != nil {
		return nil, nil, err
	}

	v143, err := parseVersion("1.4.3")
	if err != nil {
		return nil, nil, err
	}

	// Create the collection of upgraders
	upgraders := []ConfigUpgrader{
		{
			version:                 v130,
			upgradeFunc:             upgradeFuncV30,
			runAfterStaderNodeStart: true,
		},
		{
			version:     v140,
			upgradeFunc: upgradeFuncV140,
		},
		{
			version:     v142,
			upgradeFunc: upgradeFuncV142,
		},
		{
			version:     v143,
			upgradeFunc: func(c *cli.Context) error { return nil },
			needInstall: true,
		},
	}

	staderClient, err := stader.NewClientFromCtx(c)
	if err != nil {
		return nil, nil, fmt.Errorf("error NewClientFromCtx: %w", err)
	}

	cfg, _, err := staderClient.LoadConfig()
	if err != nil {
		return nil, nil, fmt.Errorf("error LoadConfig: %w", err)
	}

	defer staderClient.Close()

	// cfg nill or version empty in case fresh install
	if cfg == nil || len(cfg.Version) == 0 {
		fmt.Printf("Can not found config %+v\n", cfg)
		return nil, nil, nil
	}

	cfgVer, err := parseVersion(cfg.Version)
	if err != nil {
		return nil, nil, fmt.Errorf("error checking for config version: %w", err)
	}

	// Find the index of the provided config's version
	var needInstall bool
	for _, upgrader := range upgraders {
		cfgVerCore := cfgVer.Core()
		targetVerCore := upgrader.version.Core()

		if cfgVerCore.LessThan(targetVerCore) {
			if upgrader.runAfterStaderNodeStart {
				rundAfterUpgrades = append(rundAfterUpgrades, upgrader)
			} else {
				runBeforeUpgrades = append(runBeforeUpgrades, upgrader)
			}

			needInstall = needInstall || upgrader.needInstall
		}
	}

	if needInstall {
		runBeforeUpgrades = append([]ConfigUpgrader{
			{
				upgradeFunc: updateStaderPackage,
				version:     v0,
			},
		}, runBeforeUpgrades...)
	}

	// If there are no upgrades to apply, return
	if len(runBeforeUpgrades) == 0 && len(rundAfterUpgrades) == 0 {
		return nil, nil, nil
	}

	return runBeforeUpgrades, rundAfterUpgrades, nil
}

func upgradeFuncV30(c *cli.Context) error {
	staderClient, err := stader.NewClientFromCtx(c)
	if err != nil {
		return fmt.Errorf("error NewClientFromCtx: %w", err)
	}

	_, err = staderClient.RebuildWallet()
	if err != nil {
		return fmt.Errorf("error RebuildWallet: %w", err)
	}

	return nil
}

func upgradeFuncV140(c *cli.Context) error {
	staderClient, err := stader.NewClientFromCtx(c)
	if err != nil {
		return fmt.Errorf("error NewClientFromCtx: %w", err)
	}

	err = staderClient.UpdateGuardianConfiguration(guardian)
	if err != nil {
		return fmt.Errorf("error UpdateGuardianConfiguration: %w", err)
	}

	return nil
}

func updateStaderPackage(c *cli.Context) error {
	staderClient, err := stader.NewClientFromCtx(c)
	if err != nil {
		return fmt.Errorf("error NewClientFromCtx: %w", err)
	}

	defer staderClient.Close()

	cfg, _, err := staderClient.LoadConfig()
	if err != nil {
		return fmt.Errorf("error LoadConfig: %w", err)
	}

	dataPath, ok := cfg.StaderNode.DataPath.Value.(string)
	if !ok {
		return fmt.Errorf("error path: %s", cfg.StaderNode.DataPath.Value)
	}

	dataPath, err = homedir.Expand(dataPath)
	if err != nil {
		return fmt.Errorf("error getting data path from old configuration: %w", err)
	}

	// Install service
	path, err := homedir.Expand(cfg.StaderDirectory)
	if err != nil {
		return fmt.Errorf("error getting data path from old configuration: %w", err)
	}

	err = staderClient.UpdateStaderPackage(
		false,
		true,
		fmt.Sprintf("%s", cfg.StaderNode.Network.Value),
		fmt.Sprintf("v%s", shared.StaderVersion),
		path,
		dataPath,
	)

	if err != nil {
		return fmt.Errorf("error UpdateStaderPackage: %w", err)
	}

	return nil
}

func upgradeFuncV142(c *cli.Context) error {
	staderClient, err := stader.NewClientFromCtx(c)
	if err != nil {
		return fmt.Errorf("error NewClientFromCtx: %w", err)
	}

	cfg, _, err := staderClient.LoadConfig()
	if err != nil {
		return fmt.Errorf("error loading user settings: %w", err)
	}

	cycleMerkleRewardFile := cfg.StaderNode.GetSpRewardCyclePath(RefreshingCycle, false)

	expandedCycleMerkleRewardFile, err := homedir.Expand(cycleMerkleRewardFile)
	if err != nil {
		return fmt.Errorf("error expand cycleMerkleRewardFile: %w", err)
	}

	// Remove old cycle 5 proof
	_, err = os.Stat(expandedCycleMerkleRewardFile)
	if err == nil {
		if err = os.Remove(expandedCycleMerkleRewardFile); err != nil {
			return fmt.Errorf("error Remove old cycle 5: %w", err)
		}

		fmt.Printf("Success remove %+v \n", cycleMerkleRewardFile)
	}

	// Download new one
	_, err = staderClient.DownloadSpMerkleProofs()
	if err != nil {
		return fmt.Errorf("error DownloadSpMerkleProofs: %w", err)
	}

	return nil
}
