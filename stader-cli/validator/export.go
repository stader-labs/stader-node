package validator

import (
	"encoding/csv"
	"fmt"
	"github.com/stader-labs/stader-node/shared/services/stader"
	cliutils "github.com/stader-labs/stader-node/shared/utils/cli"
	"github.com/stader-labs/stader-node/shared/utils/log"
	"github.com/stader-labs/stader-node/shared/utils/math"
	"github.com/stader-labs/stader-node/stader-lib/types"
	"github.com/stader-labs/stader-node/stader-lib/utils/eth"
	"github.com/urfave/cli"
	"os"
)

func exportValidatorStatus(c *cli.Context) error {

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

	// Print what network we're on
	err = cliutils.PrintNetwork(staderClient)
	if err != nil {
		return err
	}

	// Get node status
	status, err := staderClient.NodeStatus()
	if err != nil {
		return err
	}

	if !status.Registered {
		fmt.Printf("The node is not registered with Stader. Please use the %sstader-cli node register%s to register with Stader", log.ColorGreen, log.ColorReset)
		return nil
	}

	if len(status.ValidatorInfos) == 0 {
		fmt.Printf("The node has no registered validators. Please use the %sstader-cli validator deposit%s command to register a validator with Stader\n\n", log.ColorGreen, log.ColorReset)
		return nil
	}

	fmt.Printf("Exporting validator status for %d validators\n\n", len(status.ValidatorInfos))

	file, err := os.Create("validator_info.csv")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	csvWriter := csv.NewWriter(file)

	header := []string{"PubKey", "Status", "WithdrawVaultAddress", "WithdrawVaultRewardBalance", "WithdrawVaultWithdrawableBalance", "DepositBlock", "WithdrawBlock"}
	err = csvWriter.Write(header)
	if err != nil {
		return err
	}

	for i := 0; i < len(status.ValidatorInfos); i++ {
		validatorRewards := 0.0
		validatorInfo := status.ValidatorInfos[i]
		if validatorInfo.WithdrawVaultRewardBalance.Int64() > 0 && !validatorInfo.CrossedRewardsThreshold {
			validatorRewards = math.RoundDown(eth.WeiToEth(validatorInfo.WithdrawVaultRewardBalance), 18)
		}

		row := []string{types.BytesToValidatorPubkey(validatorInfo.Pubkey).String(), validatorInfo.StatusToDisplay, validatorInfo.WithdrawVaultAddress.String(), fmt.Sprintf("%f", validatorRewards), fmt.Sprintf("%f", math.RoundDown(eth.WeiToEth(validatorInfo.WithdrawVaultWithdrawableBalance), 18)), validatorInfo.DepositBlock.String(), validatorInfo.WithdrawnBlock.String()}
		err = csvWriter.Write(row)
		if err != nil {
			return err
		}
	}

	csvWriter.Flush()
	if err := csvWriter.Error(); err != nil {
		panic(err)
	}

	return nil
}
