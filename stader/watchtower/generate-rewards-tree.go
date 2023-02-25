package watchtower

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/rocket-pool/rocketpool-go/rewards"
	"github.com/rocket-pool/rocketpool-go/rocketpool"
	"github.com/stader-labs/stader-node/shared/services"
	"github.com/stader-labs/stader-node/shared/services/beacon"
	"github.com/stader-labs/stader-node/shared/services/config"
	rprewards "github.com/stader-labs/stader-node/shared/services/rewards"
	"github.com/stader-labs/stader-node/shared/utils/log"
	"github.com/urfave/cli"
)

// Generate rewards Merkle Tree task
type generateRewardsTree struct {
	c         *cli.Context
	log       log.ColorLogger
	errLog    log.ColorLogger
	cfg       *config.StaderConfig
	rp        *rocketpool.RocketPool
	ec        rocketpool.ExecutionClient
	bc        beacon.Client
	lock      *sync.Mutex
	isRunning bool
}

// Create generate rewards Merkle Tree task
func newGenerateRewardsTree(c *cli.Context, logger log.ColorLogger, errorLogger log.ColorLogger) (*generateRewardsTree, error) {

	// Get services
	cfg, err := services.GetConfig(c)
	if err != nil {
		return nil, err
	}
	ec, err := services.GetEthClient(c)
	if err != nil {
		return nil, err
	}
	bc, err := services.GetBeaconClient(c)
	if err != nil {
		return nil, err
	}
	rp, err := services.GetRocketPool(c)
	if err != nil {
		return nil, err
	}

	lock := &sync.Mutex{}
	generator := &generateRewardsTree{
		c:         c,
		log:       logger,
		errLog:    errorLogger,
		cfg:       cfg,
		ec:        ec,
		bc:        bc,
		rp:        rp,
		lock:      lock,
		isRunning: false,
	}

	return generator, nil
}

// Check for generation requests
func (t *generateRewardsTree) run() error {
	t.log.Println("Checking for manual rewards tree generation requests...")

	// Check if rewards generation is already running
	t.lock.Lock()
	if t.isRunning {
		t.log.Println("Tree generation is already running.")
		t.lock.Unlock()
		return nil
	}
	t.lock.Unlock()

	// Check for requests
	requestDir := t.cfg.Smartnode.GetWatchtowerFolder(true)
	files, err := ioutil.ReadDir(requestDir)
	if os.IsNotExist(err) {
		t.log.Println("Watchtower storage directory doesn't exist, creating...")
		err = os.Mkdir(requestDir, 0755)
		if err != nil {
			return fmt.Errorf("Error creating watchtower storage directory: %w", err)
		}
	} else if err != nil {
		return fmt.Errorf("Error enumerating files in watchtower storage directory: %w", err)
	}

	for _, file := range files {
		filename := file.Name()
		if strings.HasSuffix(filename, config.RegenerateRewardsTreeRequestSuffix) && !file.IsDir() {
			// Get the index
			indexString := strings.TrimSuffix(filename, config.RegenerateRewardsTreeRequestSuffix)
			index, err := strconv.ParseUint(indexString, 0, 64)
			if err != nil {
				return fmt.Errorf("Error parsing index from [%s]: %w", filename, err)
			}

			// Delete the file
			path := filepath.Join(requestDir, filename)
			err = os.Remove(path)
			if err != nil {
				return fmt.Errorf("Error removing request file [%s]: %w", path, err)
			}

			// Generate the rewards tree
			t.lock.Lock()
			t.isRunning = true
			t.lock.Unlock()
			go t.generateRewardsTree(index)

			// Return after the first request, do others at other intervals
			return nil
		}
	}

	return nil
}

func (t *generateRewardsTree) generateRewardsTree(index uint64) {
	// Begin generation of the tree
	generationPrefix := fmt.Sprintf("[Interval %d Tree]", index)
	t.log.Printlnf("%s Starting generation of Merkle rewards tree for interval %d.", generationPrefix, index)

	// Find the event for this interval
	rewardsEvent, err := rprewards.GetRewardSnapshotEvent(t.rp, t.cfg, index)
	if err != nil {
		t.handleError(fmt.Errorf("%s Error getting event for interval %d: %w", generationPrefix, index, err))
		return
	}
	t.log.Printlnf("%s Found snapshot event: Beacon block %s, execution block %s", generationPrefix, rewardsEvent.ConsensusBlock.String(), rewardsEvent.ExecutionBlock.String())

	// Get the EL block
	elBlockHeader, err := t.ec.HeaderByNumber(context.Background(), rewardsEvent.ExecutionBlock)
	if err != nil {
		t.handleError(fmt.Errorf("%s Error getting execution block: %w", generationPrefix, err))
		return
	}

	// Try getting the rETH address as a canary to see if the block is available
	client := t.rp
	opts := &bind.CallOpts{
		BlockNumber: elBlockHeader.Number,
	}
	address, err := client.RocketStorage.GetAddress(opts, crypto.Keccak256Hash([]byte("contract.addressrocketTokenRETH")))
	if err != nil {
		errMessage := err.Error()
		t.log.Printlnf("%s Error getting state for block %d: %s", generationPrefix, elBlockHeader.Number.Uint64(), errMessage)
		if strings.Contains(errMessage, "missing trie node") || // Geth
			strings.Contains(errMessage, "No state available for block") || // Nethermind
			strings.Contains(errMessage, "Internal error") { // Besu

			// The state was missing so fall back to the archive node
			archiveEcUrl := t.cfg.Smartnode.ArchiveECUrl.Value.(string)
			if archiveEcUrl != "" {
				t.log.Printlnf("%s Primary EC cannot retrieve state for historical block %d, using archive EC [%s]", generationPrefix, elBlockHeader.Number.Uint64(), archiveEcUrl)
				ec, err := ethclient.Dial(archiveEcUrl)
				if err != nil {
					t.handleError(fmt.Errorf("Error connecting to archive EC: %w", err))
					return
				}
				client, err = rocketpool.NewRocketPool(ec, common.HexToAddress(t.cfg.Smartnode.GetStorageAddress()))
				if err != nil {
					t.handleError(fmt.Errorf("%s Error creating Rocket Pool client connected to archive EC: %w", err))
					return
				}

				// Get the rETH address from the archive EC
				address, err = client.RocketStorage.GetAddress(opts, crypto.Keccak256Hash([]byte("contract.addressrocketTokenRETH")))
				if err != nil {
					//t.handleError(fmt.Errorf("%s Error verifying rETH address with Archive EC: %w", err))
					return
				}
			} else {
				// No archive node specified
				t.handleError(fmt.Errorf("***ERROR*** Primary EC cannot retrieve state for historical block %d and the Archive EC is not specified.", elBlockHeader.Number.Uint64()))
				return
			}

		}
	}

	// Sanity check the rETH address to make sure the client is working right
	if address != t.cfg.Smartnode.GetRethAddress() {
		//t.handleError(fmt.Errorf("***ERROR*** Your Primary EC provided %s as the rETH address, but it should have been %s!", address.Hex(), t.cfg.Smartnode.GetRethAddress().Hex()))
		return
	}

	// Generate the tree
	t.generateRewardsTreeImpl(client, index, generationPrefix, rewardsEvent, elBlockHeader)
}

// Implementation for rewards tree generation using a viable EC
func (t *generateRewardsTree) generateRewardsTreeImpl(rp *rocketpool.RocketPool, index uint64, generationPrefix string, rewardsEvent rewards.RewardsEvent, elBlockHeader *types.Header) {

	// Generate the rewards file
	start := time.Now()
	treegen, err := rprewards.NewTreeGenerator(t.log, generationPrefix, rp, t.cfg, t.bc, index, rewardsEvent.IntervalStartTime, rewardsEvent.IntervalEndTime, rewardsEvent.ConsensusBlock.Uint64(), elBlockHeader, rewardsEvent.IntervalsPassed.Uint64())
	if err != nil {
		t.handleError(fmt.Errorf("%s Error creating Merkle tree generator: %w", generationPrefix, err))
		return
	}
	rewardsFile, err := treegen.GenerateTree()
	if err != nil {
		t.handleError(fmt.Errorf("%s Error generating Merkle tree: %w", generationPrefix, err))
		return
	}
	for address, network := range rewardsFile.InvalidNetworkNodes {
		t.log.Printlnf("%s WARNING: Node %s has invalid network %d assigned! Using 0 (mainnet) instead.", generationPrefix, address.Hex(), network)
	}
	t.log.Printlnf("%s Finished in %s", generationPrefix, time.Since(start).String())

	// Validate the Merkle root
	root := common.BytesToHash(rewardsFile.MerkleTree.Root())
	if root != rewardsEvent.MerkleRoot {
		t.log.Printlnf("%s WARNING: your Merkle tree had a root of %s, but the canonical Merkle tree's root was %s. This file will not be usable for claiming rewards.", generationPrefix, root.Hex(), rewardsEvent.MerkleRoot.Hex())
	} else {
		t.log.Printlnf("%s Your Merkle tree's root of %s matches the canonical root! You will be able to use this file for claiming rewards.", generationPrefix, rewardsFile.MerkleRoot)
	}

	// Create the JSON files
	rewardsFile.MinipoolPerformanceFileCID = "---"
	t.log.Printlnf("%s Saving JSON files...", generationPrefix)
	minipoolPerformanceBytes, err := json.Marshal(rewardsFile.MinipoolPerformanceFile)
	if err != nil {
		t.handleError(fmt.Errorf("%s Error serializing minipool performance file into JSON: %w", generationPrefix, err))
		return
	}
	wrapperBytes, err := json.Marshal(rewardsFile)
	if err != nil {
		t.handleError(fmt.Errorf("%s Error serializing proof wrapper into JSON: %w", generationPrefix, err))
		return
	}

	// Write the files
	path := t.cfg.Smartnode.GetRewardsTreePath(index, true)
	minipoolPerformancePath := t.cfg.Smartnode.GetMinipoolPerformancePath(index, true)
	err = ioutil.WriteFile(minipoolPerformancePath, minipoolPerformanceBytes, 0644)
	if err != nil {
		t.handleError(fmt.Errorf("%s Error saving minipool performance file to %s: %w", generationPrefix, minipoolPerformancePath, err))
		return
	}
	err = ioutil.WriteFile(path, wrapperBytes, 0644)
	if err != nil {
		t.handleError(fmt.Errorf("%s Error saving rewards file to %s: %w", generationPrefix, path, err))
		return
	}

	t.log.Printlnf("%s Merkle tree generation complete!", generationPrefix)
	t.lock.Lock()
	t.isRunning = false
	t.lock.Unlock()

}

func (t *generateRewardsTree) handleError(err error) {
	t.errLog.Println(err)
	t.errLog.Println("*** Rewards tree generation failed. ***")
	t.lock.Lock()
	t.isRunning = false
	t.lock.Unlock()
}
