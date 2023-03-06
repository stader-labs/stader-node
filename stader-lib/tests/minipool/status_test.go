package minipool

import (
	"fmt"
	"testing"

	"github.com/stader-labs/stader-minipool-go/settings/trustednode"
	"github.com/stader-labs/stader-minipool-go/types"

	"github.com/stader-labs/stader-minipool-go/minipool"
	"github.com/stader-labs/stader-minipool-go/node"
	"github.com/stader-labs/stader-minipool-go/utils/eth"

	"github.com/stader-labs/stader-minipool-go/tests/testutils/evm"
	minipoolutils "github.com/stader-labs/stader-minipool-go/tests/testutils/minipool"
	nodeutils "github.com/stader-labs/stader-minipool-go/tests/testutils/node"
)

func TestSubmitMinipoolWithdrawable(t *testing.T) {

	// State snapshotting
	if err := evm.TakeSnapshot(); err != nil {
		t.Fatal(err)
	}
	t.Cleanup(func() {
		if err := evm.RevertSnapshot(); err != nil {
			t.Fatal(err)
		}
	})

	// Register nodes
	if _, err := node.RegisterNode(rp, "Australia/Brisbane", nodeAccount.GetTransactor()); err != nil {
		t.Fatal(err)
	}
	if err := nodeutils.RegisterTrustedNode(rp, ownerAccount, trustedNodeAccount); err != nil {
		t.Fatal(err)
	}

	// Create & stake minipool
	mp, err := minipoolutils.CreateMinipool(t, rp, ownerAccount, nodeAccount, eth.EthToWei(32), 1)
	if err != nil {
		t.Fatal(err)
	}

	// Delay for the time between depositing and staking
	scrubPeriod, err := trustednode.GetScrubPeriod(rp, nil)
	if err != nil {
		t.Fatal(err)
	}
	err = evm.IncreaseTime(int(scrubPeriod + 1))
	if err != nil {
		t.Fatal(fmt.Errorf("Could not increase time: %w", err))
	}

	if err := minipoolutils.StakeMinipool(rp, mp, nodeAccount); err != nil {
		t.Fatal(err)
	}

	// Get & check initial minipool withdrawable status
	if status, err := mp.GetStatus(nil); err != nil {
		t.Error(err)
	} else if status == types.Withdrawable {
		t.Error("Incorrect initial minipool withdrawable status")
	}

	// Submit minipool withdrawable status
	if _, err := minipool.SubmitMinipoolWithdrawable(rp, mp.Address, trustedNodeAccount.GetTransactor()); err != nil {
		t.Fatal(err)
	}

	// Get & check updated minipool withdrawable status
	if status, err := mp.GetStatus(nil); err != nil {
		t.Error(err)
	} else if status != types.Withdrawable {
		t.Error("Incorrect updated minipool withdrawable status")
	}

}
