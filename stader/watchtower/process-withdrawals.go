package watchtower

import (
	"github.com/rocket-pool/rocketpool-go/rocketpool"
	"github.com/urfave/cli"

	"github.com/stader-labs/stader-node/shared/services"
	"github.com/stader-labs/stader-node/shared/services/wallet"
	"github.com/stader-labs/stader-node/shared/utils/log"
)

// Process withdrawals task
type processWithdrawals struct {
	c   *cli.Context
	log log.ColorLogger
	w   *wallet.Wallet
	rp  *rocketpool.RocketPool
}

// Create process withdrawals task
func newProcessWithdrawals(c *cli.Context, logger log.ColorLogger) (*processWithdrawals, error) {

	// Get services
	w, err := services.GetWallet(c)
	if err != nil {
		return nil, err
	}
	rp, err := services.GetRocketPool(c)
	if err != nil {
		return nil, err
	}

	// Return task
	return &processWithdrawals{
		c:   c,
		log: logger,
		w:   w,
		rp:  rp,
	}, nil

}

// Process withdrawals
func (t *processWithdrawals) run() error {

	// Process withdrawals
	// TODO: implement

	// Return
	return nil

}
