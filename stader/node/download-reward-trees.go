package node

import (
	"fmt"
	"os"

	"github.com/docker/docker/client"
	"github.com/rocket-pool/rocketpool-go/rewards"
	"github.com/rocket-pool/rocketpool-go/rocketpool"
	"github.com/urfave/cli"

	"github.com/stader-labs/stader-node/shared/services"
	"github.com/stader-labs/stader-node/shared/services/beacon"
	"github.com/stader-labs/stader-node/shared/services/config"
	rprewards "github.com/stader-labs/stader-node/shared/services/rewards"
	"github.com/stader-labs/stader-node/shared/services/wallet"
	cfgtypes "github.com/stader-labs/stader-node/shared/types/config"
	"github.com/stader-labs/stader-node/shared/utils/log"
)

// Manage download rewards trees task
type downloadRewardsTrees struct {
	c   *cli.Context
	log log.ColorLogger
	cfg *config.StaderConfig
	w   *wallet.Wallet
	rp  *rocketpool.RocketPool
	d   *client.Client
	bc  beacon.Client
}

// Create manage fee recipient task
func newDownloadRewardsTrees(c *cli.Context, logger log.ColorLogger) (*downloadRewardsTrees, error) {

	// Get services
	cfg, err := services.GetConfig(c)
	if err != nil {
		return nil, err
	}
	w, err := services.GetWallet(c)
	if err != nil {
		return nil, err
	}
	rp, err := services.GetRocketPool(c)
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
	return &downloadRewardsTrees{
		c:   c,
		log: logger,
		cfg: cfg,
		w:   w,
		rp:  rp,
		d:   d,
		bc:  bc,
	}, nil

}

// Manage fee recipient
func (d *downloadRewardsTrees) run() error {

	// Wait for eth client to sync
	if err := services.WaitEthClientSynced(d.c, true); err != nil {
		return err
	}

	// Check if the user opted into downloading rewards files
	if d.cfg.Smartnode.RewardsTreeMode.Value.(cfgtypes.RewardsMode) != cfgtypes.RewardsMode_Download {
		return nil
	}

	// Get node account
	nodeAccount, err := d.w.GetNodeAccount()
	if err != nil {
		return err
	}

	// Get the current interval
	currentIndexBig, err := rewards.GetRewardIndex(d.rp, nil)
	if err != nil {
		return err
	}
	currentIndex := currentIndexBig.Uint64()

	// Check for missing intervals
	missingIntervals := []uint64{}
	for i := uint64(0); i < currentIndex; i++ {
		// Check if the tree file exists
		treeFilePath := d.cfg.Smartnode.GetRewardsTreePath(i, true)
		_, err = os.Stat(treeFilePath)
		if os.IsNotExist(err) {
			missingIntervals = append(missingIntervals, i)
		} else if err != nil {
			return fmt.Errorf("error checking if rewards interval %d file exists: %w", i, err)
		}
	}

	if len(missingIntervals) == 0 {
		return nil
	}

	// Download missing intervals
	for _, missingInterval := range missingIntervals {
		intervalInfo, err := rprewards.GetIntervalInfo(d.rp, d.cfg, nodeAccount.Address, missingInterval)
		if err != nil {
		}
		err = rprewards.DownloadRewardsFile(d.cfg, missingInterval, intervalInfo.CID, true)
		if err != nil {
			fmt.Println()
			return err
		}
	}

	return nil

}
