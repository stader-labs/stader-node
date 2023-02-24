package node

import (
	"fmt"

	trustednodedao "github.com/stader-labs/stader-minipool-go/dao/trustednode"
	"github.com/stader-labs/stader-minipool-go/node"
	trustednodesettings "github.com/stader-labs/stader-minipool-go/settings/trustednode"
	"github.com/stader-labs/stader-minipool-go/stader"
	"github.com/stader-labs/stader-minipool-go/tokens"

	"github.com/stader-labs/stader-minipool-go/tests/testutils/accounts"
	rplutils "github.com/stader-labs/stader-minipool-go/tests/testutils/tokens/rpl"
)

// Trusted node counter
var trustedNodeIndex = 0

// Register a trusted node
func RegisterTrustedNode(rp *stader.EthxContractManager, ownerAccount *accounts.Account, trustedNodeAccount *accounts.Account) error {

	// Register node
	if _, err := node.RegisterNode(rp, "Australia/Brisbane", trustedNodeAccount.GetTransactor()); err != nil {
		return err
	}

	// Bootstrap trusted node DAO member
	if _, err := trustednodedao.BootstrapMember(rp, fmt.Sprintf("tn%d", trustedNodeIndex), fmt.Sprintf("tn%d@stader.net", trustedNodeIndex), trustedNodeAccount.Address, ownerAccount.GetTransactor()); err != nil {
		return err
	}

	// Mint trusted node RPL bond
	if err := MintTrustedNodeBond(rp, ownerAccount, trustedNodeAccount); err != nil {
		return err
	}

	// Join trusted node DAO
	if _, err := trustednodedao.Join(rp, trustedNodeAccount.GetTransactor()); err != nil {
		return err
	}

	// Increment trusted node counter & return
	trustedNodeIndex++
	return nil

}

// Mint trusted node DAO RPL bond to a node account and approve it for spending
func MintTrustedNodeBond(rp *stader.EthxContractManager, ownerAccount *accounts.Account, trustedNodeAccount *accounts.Account) error {

	// Get RPL bond amount
	rplBondAmount, err := trustednodesettings.GetRPLBond(rp, nil)
	if err != nil {
		return err
	}

	// Get RocketDAONodeTrustedActions contract address
	rocketDAONodeTrustedActionsAddress, err := rp.GetAddress("rocketDAONodeTrustedActions")
	if err != nil {
		return err
	}

	// Mint RPL to node & allow trusted node DAO contract to spend it
	if err := rplutils.MintRPL(rp, ownerAccount, trustedNodeAccount, rplBondAmount); err != nil {
		return err
	}
	if _, err := tokens.ApproveRPL(rp, *rocketDAONodeTrustedActionsAddress, rplBondAmount, trustedNodeAccount.GetTransactor()); err != nil {
		return err
	}

	// Return
	return nil

}
