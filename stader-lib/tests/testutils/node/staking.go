package node

import (
	"math/big"

	"github.com/stader-labs/stader-minipool-go/node"
	"github.com/stader-labs/stader-minipool-go/stader"
	"github.com/stader-labs/stader-minipool-go/tokens"

	"github.com/stader-labs/stader-minipool-go/tests/testutils/accounts"
	rplutils "github.com/stader-labs/stader-minipool-go/tests/testutils/tokens/rpl"
)

// Mint & stake an amount of RPL against a node
func StakeRPL(rp *stader.EthxContractManager, ownerAccount, nodeAccount *accounts.Account, amount *big.Int) error {

	// Get RocketNodeStaking contract address
	rocketNodeStakingAddress, err := rp.GetAddress("rocketNodeStaking")
	if err != nil {
		return err
	}

	// Mint, approve & stake RPL
	if err := rplutils.MintRPL(rp, ownerAccount, nodeAccount, amount); err != nil {
		return err
	}
	if _, err := tokens.ApproveRPL(rp, *rocketNodeStakingAddress, amount, nodeAccount.GetTransactor()); err != nil {
		return err
	}
	if _, err := node.StakeRPL(rp, amount, nodeAccount.GetTransactor()); err != nil {
		return err
	}

	// Return
	return nil

}
