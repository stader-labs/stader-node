package service

import (
	_ "embed"
	"fmt"
	"os"

	"github.com/hashicorp/go-version"
	"github.com/mitchellh/go-homedir"
	"github.com/stader-labs/stader-node/shared/services/stader"
	"github.com/urfave/cli"
)

type ConfigUpgrader struct {
	version     *version.Version
	upgradeFunc func(c *cli.Context) error
}

//go:embed guardian.tmpl
var guardian []byte

func migrate(c *cli.Context) ([]ConfigUpgrader, error) {
	// Create versions
	v130, err := parseVersion("1.3.0")
	if err != nil {
		return nil, err
	}

	// Create versions
	v140, err := parseVersion("1.4.0")
	if err != nil {
		return nil, err
	}

	v142, err := parseVersion("1.4.2")
	if err != nil {
		return nil, err
	}

	// Create the collection of upgraders
	upgraders := []ConfigUpgrader{
		{
			version:     v130,
			upgradeFunc: upgradeFuncV30,
		},
		{
			version:     v140,
			upgradeFunc: upgradeFuncV140,
		},
		{
			version:     v142,
			upgradeFunc: upgradeFuncV142,
		},
	}

	cfg, err := loadConfig(c)
	if err != nil {
		return nil, fmt.Errorf("error loading user settings: %w", err)
	}

	// cfg nill or version empty in case fresh install
	if cfg == nil || len(cfg.Version) == 0 {
		return nil, nil
	}

	cfgVer, err := parseVersion(cfg.Version)
	if err != nil {
		return nil, fmt.Errorf("error checking for config version: %w", err)
	}

	// Find the index of the provided config's version
	var targetUpgrades []ConfigUpgrader
	for _, upgrader := range upgraders {
		cfgVerCore := cfgVer.Core()
		targetVerCore := upgrader.version.Core()

		if cfgVerCore.LessThan(targetVerCore) {
			targetUpgrades = append(targetUpgrades, upgrader)
		}
	}

	// If there are no upgrades to apply, return
	if len(targetUpgrades) == 0 {
		return nil, nil
	}

	return targetUpgrades, nil
}

func upgradeFuncV30(c *cli.Context) error {
	staderClient, err := stader.NewClientFromCtx(c)
	if err != nil {
		return fmt.Errorf("error NewClientFromCtx: %w", err)
	}
	_, err = staderClient.RebuildWallet()

	if err != nil {
		return fmt.Errorf("error NewClientFromCtx: %w", err)
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
		return fmt.Errorf("error NewClientFromCtx: %w", err)
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

	cycleMerkleRewardFile := cfg.StaderNode.GetSpRewardCyclePath(5, true)
	expandedCycleMerkleRewardFile, err := homedir.Expand(cycleMerkleRewardFile)

	// Remove old cycle 5 proof
	_, err = os.Stat(expandedCycleMerkleRewardFile)
	if err == nil {
		if err = os.Remove(expandedCycleMerkleRewardFile); err != nil {
			return fmt.Errorf("error Remove old cycle 5: %w", err)
		}
	}

	// Download new one
	_, err = staderClient.DownloadSpMerkleProofs()
	if err != nil {
		return fmt.Errorf("error DownloadSpMerkleProofs: %w", err)
	}

	return nil
}
