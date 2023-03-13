package network

import (
	"fmt"
	"github.com/stader-labs/stader-node/shared/utils/log"

	"github.com/urfave/cli"

	"github.com/stader-labs/stader-node/shared/services/stader"
	cliutils "github.com/stader-labs/stader-node/shared/utils/cli"
)

func getStats(c *cli.Context) error {

	// Get RP client
	staderClient, err := stader.NewClientFromCtx(c)
	if err != nil {
		return err
	}
	defer staderClient.Close()

	// Check and assign the EC status
	err = cliutils.CheckClientStatus(staderClient)

	if err != nil {
		return err
	}

	// Get network stats
	response, err := staderClient.NetworkStats()
	if err != nil {
		return err
	}
	activeMinipools := response.InitializedMinipoolCount +
		response.PrelaunchMinipoolCount +
		response.StakingMinipoolCount +
		response.WithdrawableMinipoolCount +
		response.DissolvedMinipoolCount

	// Print & return
	fmt.Printf("%s========== General Stats ==========%s\n", log.ColorGreen, log.ColorReset)
	fmt.Printf("Total Value Locked:      %f ETH\n", response.TotalValueLocked)
	fmt.Printf("Staking Pool Balance:    %f ETH\n", response.DepositPoolBalance)
	fmt.Printf("Minipool Queue Demand:   %f ETH\n", response.MinipoolCapacity)
	fmt.Printf("Staking Pool ETH Used:   %f%%\n\n", response.StakerUtilization*100)

	fmt.Printf("%s============== Nodes ==============%s\n", log.ColorGreen, log.ColorReset)
	fmt.Printf("Current Commission Rate: %f%%\n", response.NodeFee*100)
	fmt.Printf("Node Count:              %d\n", response.NodeCount)
	fmt.Printf("Active Minipools:        %d\n", activeMinipools)
	fmt.Printf("    Initialized:         %d\n", response.InitializedMinipoolCount)
	fmt.Printf("    Prelaunch:           %d\n", response.PrelaunchMinipoolCount)
	fmt.Printf("    Staking:             %d\n", response.StakingMinipoolCount)
	fmt.Printf("    Withdrawable:        %d\n", response.WithdrawableMinipoolCount)
	fmt.Printf("    Dissolved:           %d\n", response.DissolvedMinipoolCount)
	fmt.Printf("Inactive Minipools:      %d\n\n", response.FinalizedMinipoolCount)

	fmt.Printf("%s========== Smoothing Pool =========%s\n", log.ColorGreen, log.ColorReset)
	fmt.Printf("Contract Address:        %s%s%s\n", log.ColorBlue, response.SmoothingPoolAddress.Hex(), log.ColorReset)
	fmt.Printf("Nodes Opted in:          %d\n", response.SmoothingPoolNodes)
	fmt.Printf("Pending Balance:         %f\n\n", response.SmoothingPoolBalance)

	fmt.Printf("%s============== Tokens =============%s\n", log.ColorGreen, log.ColorReset)
	fmt.Printf("rETH Price (ETH / rETH): %f ETH\n", response.RethPrice)
	fmt.Printf("RPL Price (ETH / RPL):   %f ETH\n", response.RplPrice)
	fmt.Printf("Total RPL staked:        %f RPL\n", response.TotalRplStaked)
	fmt.Printf("Effective RPL staked:    %f RPL\n", response.EffectiveRplStaked)

	return nil

}
