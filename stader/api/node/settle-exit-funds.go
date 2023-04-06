package node

import (
	"github.com/stader-labs/stader-node/shared/types/api"
	"github.com/stader-labs/stader-node/stader-lib/types"
	"github.com/urfave/cli"
)

func CanSettleExitFunds(c *cli.Context, validatorPubKey types.ValidatorPubkey) (*api.CanWithdrawElRewardsResponse, error) {
	return nil, nil
}

func SettleExitFunds(c *cli.Context, validatorPubKey types.ValidatorPubkey) (*api.CanWithdrawElRewardsResponse, error) {
	return nil, nil
}
