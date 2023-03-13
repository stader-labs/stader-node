package node

import (
	"context"
	"fmt"
	"math/big"
	"time"

	"github.com/docker/docker/client"
	"github.com/ethereum/go-ethereum/common"
	"github.com/rocket-pool/rocketpool-go/minipool"
	"github.com/rocket-pool/rocketpool-go/rocketpool"
	"github.com/rocket-pool/rocketpool-go/settings/trustednode"
	rptypes "github.com/rocket-pool/rocketpool-go/types"
	"github.com/rocket-pool/rocketpool-go/utils/eth"
	"github.com/urfave/cli"
	"golang.org/x/sync/errgroup"

	"github.com/stader-labs/stader-node/shared/services"
	"github.com/stader-labs/stader-node/shared/services/beacon"
	"github.com/stader-labs/stader-node/shared/services/config"
	rpgas "github.com/stader-labs/stader-node/shared/services/gas"
	"github.com/stader-labs/stader-node/shared/services/wallet"
	"github.com/stader-labs/stader-node/shared/utils/api"
	"github.com/stader-labs/stader-node/shared/utils/log"
	"github.com/stader-labs/stader-node/shared/utils/validator"
)

// Stake prelaunch minipools task
type stakePrelaunchMinipools struct {
	c              *cli.Context
	log            log.ColorLogger
	cfg            *config.StaderConfig
	w              *wallet.Wallet
	rp             *rocketpool.RocketPool
	bc             beacon.Client
	d              *client.Client
	gasThreshold   float64
	maxFee         *big.Int
	maxPriorityFee *big.Int
	gasLimit       uint64
}

// Create stake prelaunch minipools task
func newStakePrelaunchMinipools(c *cli.Context, logger log.ColorLogger) (*stakePrelaunchMinipools, error) {

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
	bc, err := services.GetBeaconClient(c)
	if err != nil {
		return nil, err
	}
	d, err := services.GetDocker(c)
	if err != nil {
		return nil, err
	}

	// Check if auto-staking is disabled
	gasThreshold := cfg.Smartnode.MinipoolStakeGasThreshold.Value.(float64)

	// Get the user-requested max fee
	maxFeeGwei := cfg.Smartnode.ManualMaxFee.Value.(float64)
	var maxFee *big.Int
	if maxFeeGwei == 0 {
		maxFee = nil
	} else {
		maxFee = eth.GweiToWei(maxFeeGwei)
	}

	// Get the user-requested max fee
	priorityFeeGwei := cfg.Smartnode.PriorityFee.Value.(float64)
	var priorityFee *big.Int
	if priorityFeeGwei == 0 {
		logger.Println("WARNING: priority fee was missing or 0, setting a default of 2.")
		priorityFee = eth.GweiToWei(2)
	} else {
		priorityFee = eth.GweiToWei(priorityFeeGwei)
	}

	// Return task
	return &stakePrelaunchMinipools{
		c:              c,
		log:            logger,
		cfg:            cfg,
		w:              w,
		rp:             rp,
		bc:             bc,
		d:              d,
		gasThreshold:   gasThreshold,
		maxFee:         maxFee,
		maxPriorityFee: priorityFee,
		gasLimit:       0,
	}, nil

}

// Stake prelaunch minipools
func (t *stakePrelaunchMinipools) run() error {

	// Reload the wallet (in case a call to `node deposit` changed it)
	if err := t.w.Reload(); err != nil {
		return err
	}

	// Wait for eth client to sync
	if err := services.WaitEthClientSynced(t.c, true); err != nil {
		return err
	}

	// Get node account
	nodeAccount, err := t.w.GetNodeAccount()
	if err != nil {
		return err
	}

	// Get prelaunch minipools
	minipools, err := t.getPrelaunchMinipools(nodeAccount.Address)
	if err != nil {
		return err
	}
	if len(minipools) == 0 {
		return nil
	}

	// Get eth2 config
	eth2Config, err := t.bc.GetEth2Config()
	if err != nil {
		return err
	}
	
	// Stake minipools
	successCount := 0
	for _, mp := range minipools {
		success, err := t.stakeMinipool(mp, eth2Config)
		if err != nil {
			t.log.Println(fmt.Errorf("Could not stake minipool %s: %w", mp.Address.Hex(), err))
			return err
		}
		if success {
			successCount++
		}
	}

	// Restart validator process if any minipools were staked successfully
	if successCount > 0 {
		if err := validator.RestartValidator(t.cfg, t.bc, &t.log, t.d); err != nil {
			return err
		}
	}

	// Return
	return nil

}

// Get prelaunch minipools
func (t *stakePrelaunchMinipools) getPrelaunchMinipools(nodeAddress common.Address) ([]*minipool.Minipool, error) {

	// Get node minipool addresses
	addresses, err := minipool.GetNodeMinipoolAddresses(t.rp, nodeAddress, nil)
	if err != nil {
		return []*minipool.Minipool{}, err
	}

	// Create minipool contracts
	minipools := make([]*minipool.Minipool, len(addresses))
	for mi, address := range addresses {
		mp, err := minipool.NewMinipool(t.rp, address, nil)
		if err != nil {
			return []*minipool.Minipool{}, err
		}
		minipools[mi] = mp
	}

	// Data
	var wg errgroup.Group
	statuses := make([]minipool.StatusDetails, len(minipools))

	// Load minipool statuses
	for mi, mp := range minipools {
		mi, mp := mi, mp
		wg.Go(func() error {
			status, err := mp.GetStatusDetails(nil)
			if err == nil {
				statuses[mi] = status
			}
			return err
		})
	}

	// Wait for data
	if err := wg.Wait(); err != nil {
		return []*minipool.Minipool{}, err
	}

	// Get the scrub period
	scrubPeriodSeconds, err := trustednode.GetScrubPeriod(t.rp, nil)
	if err != nil {
		return []*minipool.Minipool{}, err
	}
	scrubPeriod := time.Duration(scrubPeriodSeconds) * time.Second

	// Get the time of the latest block
	latestEth1Block, err := t.rp.Client.HeaderByNumber(context.Background(), nil)
	if err != nil {
		return []*minipool.Minipool{}, fmt.Errorf("Can't get the latest block time: %w", err)
	}
	latestBlockTime := time.Unix(int64(latestEth1Block.Time), 0)

	// Filter minipools by status
	prelaunchMinipools := []*minipool.Minipool{}
	for mi, mp := range minipools {
		if statuses[mi].Status == rptypes.Prelaunch {
			creationTime := statuses[mi].StatusTime
			remainingTime := creationTime.Add(scrubPeriod).Sub(latestBlockTime)
			if remainingTime < 0 {
				prelaunchMinipools = append(prelaunchMinipools, mp)
			} else {
				t.log.Printlnf("Minipool %s has %s left until it can be staked.", mp.Address.Hex(), remainingTime)
			}
		}
	}

	// Return
	return prelaunchMinipools, nil

}

// Stake a minipool
func (t *stakePrelaunchMinipools) stakeMinipool(mp *minipool.Minipool, eth2Config beacon.Eth2Config) (bool, error) {

	// Log
	t.log.Printlnf("Staking minipool %s...", mp.Address.Hex())

	// Get minipool withdrawal credentials
	withdrawalCredentials, err := minipool.GetMinipoolWithdrawalCredentials(t.rp, mp.Address, nil)
	if err != nil {
		return false, err
	}

	// Get the validator key for the minipool
	validatorPubkey, err := minipool.GetMinipoolPubkey(t.rp, mp.Address, nil)
	if err != nil {
		return false, err
	}
	validatorKey, err := t.w.GetValidatorKeyByPubkey(validatorPubkey)
	if err != nil {
		return false, err
	}

	// Get validator deposit data
	depositData, depositDataRoot, err := validator.GetDepositData(validatorKey, withdrawalCredentials, eth2Config, 1)
	if err != nil {
		return false, err
	}

	// Get transactor
	opts, err := t.w.GetNodeAccountTransactor()
	if err != nil {
		return false, err
	}

	// Get the gas limit
	signature := rptypes.BytesToValidatorSignature(depositData.Signature)
	gasInfo, err := mp.EstimateStakeGas(signature, depositDataRoot, opts)
	if err != nil {
		return false, fmt.Errorf("Could not estimate the gas required to stake the minipool: %w", err)
	}
	var gas *big.Int
	if t.gasLimit != 0 {
		gas = new(big.Int).SetUint64(t.gasLimit)
	} else {
		gas = new(big.Int).SetUint64(gasInfo.SafeGasLimit)
	}

	// Get the max fee
	maxFee := t.maxFee
	if maxFee == nil || maxFee.Uint64() == 0 {
		maxFee, err = rpgas.GetHeadlessMaxFeeWei()
		if err != nil {
			return false, err
		}
	}

	// Print the gas info
	if !api.PrintAndCheckGasInfo(gasInfo, true, t.gasThreshold, t.log, maxFee, t.gasLimit) {
		// Check for the timeout buffer
		prelaunchTime, err := mp.GetStatusTime(nil)
		if err != nil {
			t.log.Printlnf("Error checking minipool launch time: %s\nStaking now for safety...", err.Error())
		}
		isDue, timeUntilDue, err := api.IsTransactionDue(t.rp, prelaunchTime)
		if err != nil {
			t.log.Printlnf("Error checking if minipool is due: %s\nStaking now for safety...", err.Error())
		}
		if !isDue {
			t.log.Printlnf("Time until staking will be forced for safety: %s", timeUntilDue)
			return false, nil
		}

		t.log.Println("NOTICE: The minipool has exceeded half of the timeout period, so it will be force-staked at the current gas price.")
	}

	opts.GasFeeCap = maxFee
	opts.GasTipCap = t.maxPriorityFee
	opts.GasLimit = gas.Uint64()

	// Stake minipool
	hash, err := mp.Stake(
		signature,
		depositDataRoot,
		opts,
	)
	if err != nil {
		return false, err
	}

	// Print TX info and wait for it to be included in a block
	err = api.PrintAndWaitForTransaction(t.cfg, hash, t.rp.Client, t.log)
	if err != nil {
		return false, err
	}

	// Log
	t.log.Printlnf("Successfully staked minipool %s.", mp.Address.Hex())

	// Return
	return true, nil

}
