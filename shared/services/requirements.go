/*
This work is licensed and released under GNU GPL v3 or any other later versions.
The full text of the license is below/ found at <http://www.gnu.org/licenses/>

(c) 2023 Rocket Pool Pty Ltd. Modified under GNU GPL v3. [1.2.1]

This program is free software: you can redistribute it and/or modify
it under the terms of the GNU General Public License as published by
the Free Software Foundation, either version 3 of the License, or
(at your option) any later version.

This program is distributed in the hope that it will be useful,
but WITHOUT ANY WARRANTY; without even the implied warranty of
MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
GNU General Public License for more details.

You should have received a copy of the GNU General Public License
along with this program.  If not, see <http://www.gnu.org/licenses/>.
*/
package services

import (
	"context"
	"errors"
	"fmt"
	"log"
	"math/big"
	"sync"
	"time"

	"github.com/ethereum/go-ethereum/common"

	"github.com/stader-labs/stader-node/shared/services/config"
	"github.com/stader-labs/stader-node/stader-lib/node"
	"github.com/stader-labs/stader-node/stader-lib/stader"
	"github.com/urfave/cli"
)

// Settings
const EthClientSyncTimeout = 16    // 16 seconds
const BeaconClientSyncTimeout = 16 // 16 seconds
var checkNodePasswordInterval, _ = time.ParseDuration("15s")
var checkNodeWalletInterval, _ = time.ParseDuration("15s")
var checkNodeRegisteredInterval, _ = time.ParseDuration("10s")
var ethClientSyncPollInterval, _ = time.ParseDuration("5s")
var beaconClientSyncPollInterval, _ = time.ParseDuration("5s")
var ethClientRecentBlockThreshold, _ = time.ParseDuration("5m")
var ethClientStatusRefreshInterval, _ = time.ParseDuration("60s")

//
// Service requirements
//

func RequireNodePassword(c *cli.Context) error {
	nodePasswordSet, err := getNodePasswordSet(c)
	if err != nil {
		return err
	}
	if !nodePasswordSet {
		return errors.New("The node password has not been set. Please run './stader-cli wallet init' and try again.")
	}
	return nil
}

func RequireNodeWallet(c *cli.Context) error {
	if err := RequireNodePassword(c); err != nil {
		return err
	}
	nodeWalletInitialized, err := getNodeWalletInitialized(c)
	if err != nil {
		return err
	}
	if !nodeWalletInitialized {
		return errors.New("The node wallet has not been initialized. Please run './stader-cli wallet init' and try again.")
	}
	return nil
}

func RequireEthClientSynced(c *cli.Context) error {
	ethClientSynced, err := waitEthClientSynced(c, false, EthClientSyncTimeout)
	if err != nil {
		return err
	}
	if !ethClientSynced {
		return errors.New("The Eth 1.0 node is currently syncing. Please try again later.")
	}
	return nil
}

func RequireBeaconClientSynced(c *cli.Context) error {
	beaconClientSynced, err := waitBeaconClientSynced(c, false, BeaconClientSyncTimeout)
	if err != nil {
		return err
	}
	if !beaconClientSynced {
		return errors.New("The Eth 2.0 node is currently syncing. Please try again later.")
	}
	return nil
}

func RequireNodeRegistered(c *cli.Context) error {
	if err := RequireNodeWallet(c); err != nil {
		return err
	}
	nodeRegistered, err := isNodeRegistered(c)
	if err != nil {
		return err
	}
	if !nodeRegistered {
		return errors.New("The node is not registered with Stader. Please run 'stader-cli node register' and try again.")
	}
	return nil
}

func RequireNodeActive(c *cli.Context) error {
	if err := RequireNodeWallet(c); err != nil {
		return err
	}
	nodeActive, err := isNodeActive(c)
	if err != nil {
		return err
	}
	if !nodeActive {
		return errors.New("The node has been deactivated with Stader. You might have front run one of your keys, Please reach out to the Stader team on discord for more information.")
	}
	return nil
}

//
// Service synchronization
//

func WaitNodePassword(c *cli.Context, verbose bool) error {
	for {
		nodePasswordSet, err := getNodePasswordSet(c)
		if err != nil {
			return err
		}
		if nodePasswordSet {
			return nil
		}
		if verbose {
			log.Printf("The node password has not been set, retrying in %s...\n", checkNodePasswordInterval.String())
		}
		time.Sleep(checkNodePasswordInterval)
	}
}

func WaitNodeWallet(c *cli.Context, verbose bool) error {
	if err := WaitNodePassword(c, verbose); err != nil {
		return err
	}
	for {
		nodeWalletInitialized, err := getNodeWalletInitialized(c)
		if err != nil {
			return err
		}
		if nodeWalletInitialized {
			return nil
		}
		if verbose {
			log.Printf("The node wallet has not been initialized, retrying in %s...\n", checkNodeWalletInterval.String())
		}
		time.Sleep(checkNodeWalletInterval)
	}
}

func WaitNodeRegistered(c *cli.Context, operatorAddress common.Address, verbose bool) error {
	if err := WaitNodeWallet(c, verbose); err != nil {
		return err
	}
	for {
		pnr, err := GetPermissionlessNodeRegistry(c)
		if err != nil {
			return err
		}
		operatorId, err := node.GetOperatorId(pnr, operatorAddress, nil)
		if err != nil {
			return err
		}
		if operatorId.Cmp(big.NewInt(0)) != 0 {
			return nil
		}
		if verbose {
			log.Printf("The node is not registered with Stader, retrying in %s...\n", checkNodeRegisteredInterval.String())
		}
		time.Sleep(checkNodeRegisteredInterval)
	}
}

func WaitEthClientSynced(c *cli.Context, verbose bool) error {
	_, err := waitEthClientSynced(c, verbose, 0)
	return err
}

func WaitBeaconClientSynced(c *cli.Context, verbose bool) error {
	_, err := waitBeaconClientSynced(c, verbose, 0)
	return err
}

//
// Helpers
//

// Check if the node password is set
func getNodePasswordSet(c *cli.Context) (bool, error) {
	pm, err := GetPasswordManager(c)
	if err != nil {
		return false, err
	}
	return pm.IsPasswordSet(), nil
}

// Check if the node wallet is initialized
func getNodeWalletInitialized(c *cli.Context) (bool, error) {
	w, err := GetWallet(c)
	if err != nil {
		return false, err
	}
	return w.GetInitialized()
}

// Check if the node is registered
func isNodeRegistered(c *cli.Context) (bool, error) {
	w, err := GetWallet(c)
	if err != nil {
		return false, err
	}
	pnr, err := GetPermissionlessNodeRegistry(c)
	if err != nil {
		return false, err
	}
	nodeAccount, err := w.GetNodeAccount()
	if err != nil {
		return false, err
	}
	operatorId, err := node.GetOperatorId(pnr, nodeAccount.Address, nil)
	if err != nil {
		return false, err
	}

	return operatorId.Int64() != 0, nil
}

func isNodeActive(c *cli.Context) (bool, error) {
	w, err := GetWallet(c)
	if err != nil {
		return false, err
	}
	pnr, err := GetPermissionlessNodeRegistry(c)
	if err != nil {
		return false, err
	}
	nodeAccount, err := w.GetNodeAccount()
	if err != nil {
		return false, err
	}
	operatorId, err := node.GetOperatorId(pnr, nodeAccount.Address, nil)
	if err != nil {
		return false, err
	}
	operatorInfo, err := node.GetOperatorInfo(pnr, operatorId, nil)
	if err != nil {
		return false, err
	}

	return operatorInfo.Active, nil
}

// Wait for the eth client to sync
// timeout of 0 indicates no timeout
var ethClientSyncLock sync.Mutex

func checkExecutionClientStatus(ecMgr *ExecutionClientManager, cfg *config.StaderConfig) (bool, stader.ExecutionClient, error) {

	// Check the EC status
	mgrStatus := ecMgr.CheckStatus(cfg)
	if ecMgr.primaryReady {
		return true, nil, nil
	}

	// If the primary isn't synced but there's a fallback and it is, return true
	if ecMgr.fallbackReady {
		if mgrStatus.PrimaryClientStatus.Error != "" {
			log.Printf("Primary execution client is unavailable (%s), using fallback execution client...\n", mgrStatus.PrimaryClientStatus.Error)
		} else {
			log.Printf("Primary execution client is still syncing (%.2f%%), using fallback execution client...\n", mgrStatus.PrimaryClientStatus.SyncProgress*100)
		}
		return true, nil, nil
	}

	// If neither is synced, go through the status to figure out what to do

	// Is the primary working and syncing? If so, wait for it
	if mgrStatus.PrimaryClientStatus.IsWorking && mgrStatus.PrimaryClientStatus.Error == "" {
		log.Printf("Fallback execution client is not configured or unavailable, waiting for primary execution client to finish syncing (%.2f%%)\n", mgrStatus.PrimaryClientStatus.SyncProgress*100)
		return false, ecMgr.primaryEc, nil
	}

	// Is the fallback working and syncing? If so, wait for it
	if mgrStatus.FallbackEnabled && mgrStatus.FallbackClientStatus.IsWorking && mgrStatus.FallbackClientStatus.Error == "" {
		log.Printf("Primary execution client is unavailable (%s), waiting for the fallback execution client to finish syncing (%.2f%%)\n", mgrStatus.PrimaryClientStatus.Error, mgrStatus.FallbackClientStatus.SyncProgress*100)
		return false, ecMgr.fallbackEc, nil
	}

	// If neither client is working, report the errors
	if mgrStatus.FallbackEnabled {
		return false, nil, fmt.Errorf("Primary execution client is unavailable (%s) and fallback execution client is unavailable (%s), no execution clients are ready.", mgrStatus.PrimaryClientStatus.Error, mgrStatus.FallbackClientStatus.Error)
	}

	return false, nil, fmt.Errorf("Primary execution client is unavailable (%s) and no fallback execution client is configured.", mgrStatus.PrimaryClientStatus.Error)
}

func checkBeaconClientStatus(bcMgr *BeaconClientManager) (bool, error) {

	// Check the BC status
	mgrStatus := bcMgr.CheckStatus()
	if bcMgr.primaryReady {
		return true, nil
	}

	// If the primary isn't synced but there's a fallback and it is, return true
	if bcMgr.fallbackReady {
		if mgrStatus.PrimaryClientStatus.Error != "" {
			log.Printf("Primary consensus client is unavailable (%s), using fallback consensus client...\n", mgrStatus.PrimaryClientStatus.Error)
		} else {
			log.Printf("Primary consensus client is still syncing (%.2f%%), using fallback consensus client...\n", mgrStatus.PrimaryClientStatus.SyncProgress*100)
		}
		return true, nil
	}

	// If neither is synced, go through the status to figure out what to do

	// Is the primary working and syncing? If so, wait for it
	if mgrStatus.PrimaryClientStatus.IsWorking && mgrStatus.PrimaryClientStatus.Error == "" {
		log.Printf("Fallback consensus client is not configured or unavailable, waiting for primary consensus client to finish syncing (%.2f%%)\n", mgrStatus.PrimaryClientStatus.SyncProgress*100)
		return false, nil
	}

	// Is the fallback working and syncing? If so, wait for it
	if mgrStatus.FallbackEnabled && mgrStatus.FallbackClientStatus.IsWorking && mgrStatus.FallbackClientStatus.Error == "" {
		log.Printf("Primary cosnensus client is unavailable (%s), waiting for the fallback consensus client to finish syncing (%.2f%%)\n", mgrStatus.PrimaryClientStatus.Error, mgrStatus.FallbackClientStatus.SyncProgress*100)
		return false, nil
	}

	// If neither client is working, report the errors
	if mgrStatus.FallbackEnabled {
		return false, fmt.Errorf("Primary consensus client is unavailable (%s) and fallback consensus client is unavailable (%s), no consensus clients are ready.", mgrStatus.PrimaryClientStatus.Error, mgrStatus.FallbackClientStatus.Error)
	}

	return false, fmt.Errorf("Primary consensus client is unavailable (%s) and no fallback consensus client is configured.", mgrStatus.PrimaryClientStatus.Error)
}

func waitEthClientSynced(c *cli.Context, verbose bool, timeout int64) (bool, error) {

	// Prevent multiple waiting goroutines from requesting sync progress
	ethClientSyncLock.Lock()
	defer ethClientSyncLock.Unlock()

	// Get eth client
	ecMgr, err := GetEthClient(c)
	if err != nil {
		return false, err
	}

	cfg, err := GetConfig(c)
	if err != nil {
		return false, err
	}

	synced, clientToCheck, err := checkExecutionClientStatus(ecMgr, cfg)
	if err != nil {
		return false, err
	}
	if synced {
		return true, nil
	}

	// Get wait start time
	startTime := time.Now()

	// Get EC status refresh time
	ecRefreshTime := startTime

	// Wait for sync
	for {

		// Check timeout
		if (timeout > 0) && (time.Since(startTime).Seconds() > float64(timeout)) {
			return false, nil
		}

		// Check if the EC status needs to be refreshed
		if time.Since(ecRefreshTime) > ethClientStatusRefreshInterval {
			log.Println("Refreshing primary / fallback execution client status...")
			ecRefreshTime = time.Now()
			synced, clientToCheck, err = checkExecutionClientStatus(ecMgr, cfg)
			if err != nil {
				return false, err
			}
			if synced {
				return true, nil
			}
		}

		// Get sync progress
		progress, err := clientToCheck.SyncProgress(context.Background())
		if err != nil {
			return false, err
		}

		// Check sync progress
		if progress != nil {
			if verbose {
				p := float64(progress.CurrentBlock-progress.StartingBlock) / float64(progress.HighestBlock-progress.StartingBlock)
				if p > 1 {
					log.Println("Eth 1.0 node syncing...")
				} else {
					log.Printf("Eth 1.0 node syncing: %.2f%%\n", p*100)
				}
			}
		} else {
			// Eth 1 client is not in "syncing" state but may be behind head
			// Get the latest block it knows about and make sure it's recent compared to system clock time
			isUpToDate, _, err := IsSyncWithinThreshold(clientToCheck)
			if err != nil {
				return false, err
			}
			// Only return true if the last reportedly known block is within our defined threshold
			if isUpToDate {
				return true, nil
			}
		}

		// Pause before next poll
		time.Sleep(ethClientSyncPollInterval)

	}

}

// Wait for the beacon client to sync
// timeout of 0 indicates no timeout
var beaconClientSyncLock sync.Mutex

func waitBeaconClientSynced(c *cli.Context, verbose bool, timeout int64) (bool, error) {

	// Prevent multiple waiting goroutines from requesting sync progress
	beaconClientSyncLock.Lock()
	defer beaconClientSyncLock.Unlock()

	// Get beacon client
	bcMgr, err := GetBeaconClient(c)
	if err != nil {
		return false, err
	}

	synced, err := checkBeaconClientStatus(bcMgr)
	if err != nil {
		return false, err
	}
	if synced {
		return true, nil
	}

	// Get wait start time
	startTime := time.Now()

	// Get BC status refresh time
	bcRefreshTime := startTime

	// Wait for sync
	for {

		// Check timeout
		if (timeout > 0) && (time.Since(startTime).Seconds() > float64(timeout)) {
			return false, nil
		}

		// Check if the BC status needs to be refreshed
		if time.Since(bcRefreshTime) > ethClientStatusRefreshInterval {
			log.Println("Refreshing primary / fallback consensus client status...")
			bcRefreshTime = time.Now()
			synced, err = checkBeaconClientStatus(bcMgr)
			if err != nil {
				return false, err
			}
			if synced {
				return true, nil
			}
		}

		// Get sync status
		syncStatus, err := bcMgr.GetSyncStatus()
		if err != nil {
			return false, err
		}

		// Check sync status
		if syncStatus.Syncing {
			if verbose {
				log.Println("Eth 2.0 node syncing: %.2f%%\n", syncStatus.Progress*100)
			}
		} else {
			return true, nil
		}

		// Pause before next poll
		time.Sleep(beaconClientSyncPollInterval)

	}

}

// Confirm the EC's latest block is within the threshold of the current system clock
func IsSyncWithinThreshold(ec stader.ExecutionClient) (bool, time.Time, error) {
	timestamp, err := GetEthClientLatestBlockTimestamp(ec)
	if err != nil {
		return false, time.Time{}, err
	}

	// Return true if the latest block is under the threshold
	blockTime := time.Unix(int64(timestamp), 0)
	if time.Since(blockTime) < ethClientRecentBlockThreshold {
		return true, blockTime, nil
	}

	return false, blockTime, nil
}
