package node

import (
	"github.com/docker/docker/client"
	"github.com/stader-labs/stader-node/shared/services"
	"github.com/stader-labs/stader-node/shared/services/beacon"
	"github.com/stader-labs/stader-node/shared/services/config"
	"github.com/stader-labs/stader-node/shared/services/wallet"
	"github.com/stader-labs/stader-node/shared/utils/log"
	"github.com/stader-labs/stader-node/stader-lib/stader"
	"github.com/urfave/cli"
)

type validatorLauncher struct {
	c   *cli.Context
	log log.ColorLogger
	cfg *config.StaderConfig
	w   *wallet.Wallet
	prn *stader.PermissionlessNodeRegistryContractManager
	d   *client.Client
	bc  beacon.Client
}

func newValidatorLauncher(c *cli.Context, logger log.ColorLogger) (*validatorLauncher, error) {
	// Get services
	cfg, err := services.GetConfig(c)
	if err != nil {
		return nil, err
	}
	w, err := services.GetWallet(c)
	if err != nil {
		return nil, err
	}
	prn, err := services.GetPermissionlessNodeRegistry(c)
	if err != nil {
		return nil, err
	}
	d, err := services.GetDocker(c)
	if err != nil {
		return nil, err
	}
	bc, err := services.GetBeaconClient(c)
	if err != nil {
		return nil, err
	}

	// Return task
	return &validatorLauncher{
		c:   c,
		log: logger,
		cfg: cfg,
		w:   w,
		prn: prn,
		d:   d,
		bc:  bc,
	}, nil
}

func (vl *validatorLauncher) run() error {
	// Wait for eth client to sync
	if err := services.WaitEthClientSynced(vl.c, true); err != nil {
		return err
	}

	// fetch total validator registered

	// start from the latest validator and iterate to the earliest validator

	// if we find a validator we have deposited to in the last 24 hrs based on the

	return nil
}
