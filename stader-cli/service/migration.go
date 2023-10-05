package service

import (
	_ "embed"
	"fmt"

	"github.com/hashicorp/go-version"
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
	v131, err := parseVersion("1.3.1")
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
			version:     v131,
			upgradeFunc: upgradeFuncV31,
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

func upgradeFuncV31(c *cli.Context) error {
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
