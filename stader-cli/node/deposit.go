package node

import (
	"crypto/rand"
	"fmt"
	"math/big"

	"github.com/rocket-pool/rocketpool-go/rocketpool"
	"github.com/stader-labs/stader-minipool-go/utils/eth"
	"github.com/stader-labs/stader-node/shared/services/gas"
	"github.com/urfave/cli"

	"github.com/stader-labs/stader-node/shared/services/stader"
	cliutils "github.com/stader-labs/stader-node/shared/utils/cli"
	"github.com/stader-labs/stader-node/shared/utils/math"
)

// Config
const DefaultMaxNodeFeeSlippage = 0.01 // 1% below current network fee

func nodeDeposit(c *cli.Context) error {

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

	fmt.Println("Your eth2 client is on the correct network.")

	// Post a warning about fee distribution
	if !(c.Bool("yes") || cliutils.Confirm(fmt.Sprintf("%sNOTE: by creating a new validator, your node will automatically claim and distribute any balance you have in your fee distributor contract. If you don't want to claim your balance at this time, you should not create a new validator.%s\nWould you like to continue?", colorYellow, colorReset))) {
		fmt.Println("Cancelled.")
		return nil
	}

	// Force 4 ETH minipools as the only option after much community discussion
	amountWei := eth.EthToWei(4.0)

	// Get minipool salt
	var salt *big.Int
	if c.String("salt") != "" {
		var success bool
		salt, success = big.NewInt(0).SetString(c.String("salt"), 0)
		if !success {
			return fmt.Errorf("Invalid validator salt: %s", c.String("salt"))
		}
	} else {
		buffer := make([]byte, 32)
		_, err = rand.Read(buffer)
		if err != nil {
			return fmt.Errorf("Error generating random salt: %w", err)
		}
		salt = big.NewInt(0).SetBytes(buffer)
	}

	// Check to see if eth2 is synced
	colorReset := "\033[0m"
	colorRed := "\033[31m"
	colorYellow := "\033[33m"
	syncResponse, err := staderClient.NodeSync()
	if err != nil {
		fmt.Printf("%s**WARNING**: Can't verify the sync status of your consensus client.\nYOU WILL LOSE ETH if your validator is activated before it is fully synced.\n"+
			"Reason: %s\n%s", colorRed, err, colorReset)
	} else {
		if syncResponse.BcStatus.PrimaryClientStatus.IsSynced {
			fmt.Printf("Your consensus client is synced, you may safely create a validator.\n")
		} else if syncResponse.BcStatus.FallbackEnabled {
			if syncResponse.BcStatus.FallbackClientStatus.IsSynced {
				fmt.Printf("Your fallback consensus client is synced, you may safely create a validator.\n")
			} else {
				fmt.Printf("%s**WARNING**: neither your primary nor fallback consensus clients are fully synced.\nYOU WILL LOSE ETH if your validator is activated before they are fully synced.\n%s", colorRed, colorReset)
			}
		} else {
			fmt.Printf("%s**WARNING**: your primary consensus client is either not fully synced or offline and you do not have a fallback client configured.\nYOU WILL LOSE ETH if your validator is activated before it is fully synced.\n%s", colorRed, colorReset)
		}
	}

	// Assign max fees
	err = gas.AssignMaxFeeAndLimit(rocketpool.GasInfo{
		EstGasLimit:  25000000,
		SafeGasLimit: 30000000,
	}, staderClient, c.Bool("yes"))
	if err != nil {
		return err
	}

	// Prompt for confirmation
	if !(c.Bool("yes") || cliutils.Confirm(fmt.Sprintf(
		"You are about to deposit %.6f ETH to create a validator with a minimum possible commission rate of %f%%.\n"+
			"%sARE YOU SURE YOU WANT TO DO THIS? Running a validator is a long-term commitment, and this action cannot be undone!%s",
		math.RoundDown(eth.WeiToEth(amountWei), 6),
		5.0,
		colorYellow,
		colorReset))) {
		fmt.Println("Cancelled.")
		return nil
	}

	// Make deposit
	response, err := staderClient.NodeDeposit(amountWei, salt, true)
	if err != nil {
		return err
	}

	// Log and wait for the minipool address
	fmt.Printf("Creating validator...\n")
	cliutils.PrintTransactionHash(staderClient, response.TxHash)
	_, err = staderClient.WaitForTransaction(response.TxHash)
	if err != nil {
		return err
	}

	// Log & return
	fmt.Printf("The node deposit of %.6f ETH was made successfully!\n", math.RoundDown(eth.WeiToEth(amountWei), 6))
	// fmt.Printf("Your new validator's address is: %s\n", response.MinipoolAddress)
	fmt.Printf("The validator pubkey is: %s\n\n", response.ValidatorPubkey.Hex())

	fmt.Println("Your validator is now in Initialized status.")
	fmt.Println("Once the 4 ETH deposit has been matched by the staking pool, it will move to Prelaunch status.")
	// fmt.Printf("After that, it will move to Staking status once %s have passed.\n", response.ScrubPeriod)
	fmt.Println("You can watch its progress using `stader-cli service logs node`.")

	return nil

}
