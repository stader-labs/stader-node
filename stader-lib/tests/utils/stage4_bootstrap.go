package utils

import (
	"time"

	"github.com/stader-labs/stader-minipool-go/settings/protocol"
	"github.com/stader-labs/stader-minipool-go/stader"
	"github.com/stader-labs/stader-minipool-go/tests/testutils/accounts"
	"github.com/stader-labs/stader-minipool-go/utils/eth"
)

// Bootstrap all of the parameters to mimic Stage 4 so the unit tests work correctly
func Stage4Bootstrap(rp *stader.EthxContractManager, ownerAccount *accounts.Account) {

	opts := ownerAccount.GetTransactor()

	protocol.BootstrapDepositEnabled(rp, true, opts)
	protocol.BootstrapAssignDepositsEnabled(rp, true, opts)
	protocol.BootstrapMaximumDepositPoolSize(rp, eth.EthToWei(1000), opts)
	protocol.BootstrapNodeRegistrationEnabled(rp, true, opts)
	protocol.BootstrapNodeDepositEnabled(rp, true, opts)
	protocol.BootstrapMinipoolSubmitWithdrawableEnabled(rp, true, opts)
	protocol.BootstrapMinimumNodeFee(rp, 0.05, opts)
	protocol.BootstrapTargetNodeFee(rp, 0.1, opts)
	protocol.BootstrapMaximumNodeFee(rp, 0.2, opts)
	protocol.BootstrapNodeFeeDemandRange(rp, eth.EthToWei(1000), opts)
	protocol.BootstrapInflationStartTime(rp,
		uint64(time.Now().Unix()+(60*60*24*14)), opts)

}
