package guardian

import (
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math/big"
	"net/http"
	"sync"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/stader-labs/stader-node/shared/services/beacon"
	"github.com/stader-labs/stader-node/shared/services/config"
	"github.com/urfave/cli"
	"golang.org/x/sync/errgroup"

	"github.com/stader-labs/stader-node/shared/services"
	"github.com/stader-labs/stader-node/shared/services/wallet"
	"github.com/stader-labs/stader-node/shared/utils/log"
)

const (
	beaconChainAPI = "https://beaconcha.in/api/v1"
)

type submitValidatorStats struct {
	c   *cli.Context
	log log.ColorLogger
	cfg *config.StaderConfig
	w   *wallet.Wallet
	bc  beacon.Client
}

func getValidatorStatus(pubKey string, blockNumber uint64) (string, error) {
	decodedPubKey, err := hex.DecodeString(pubKey)
	if err != nil {
		return "", fmt.Errorf("invalid public key: %w", err)
	}

	encodedPubKey := hex.EncodeToString(decodedPubKey)
	url := fmt.Sprintf("%s/validator/%s?epoch=%d", beaconChainAPI, encodedPubKey, blockNumber)

	resp, err := http.Get(url)
	if err != nil {
		return "", fmt.Errorf("failed to fetch validator data: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", fmt.Errorf("failed to read response body: %w", err)
	}

	var jsonResponse map[string]interface{}
	err = json.Unmarshal(body, &jsonResponse)
	if err != nil {
		return "", fmt.Errorf("failed to parse JSON response: %w", err)
	}

	status, ok := jsonResponse["status"].(string)
	if !ok {
		return "", fmt.Errorf("status field not found in JSON response")
	}

	return status, nil
}

func countValidatorsByStatus(pubKeys []string, blockNumber uint64, status string) (int, error) {
	var count int
	var wg sync.WaitGroup
	statusChan := make(chan string)

	wg.Add(len(pubKeys))
	for _, pubKey := range pubKeys {
		go func(pubKey string) {
			defer wg.Done()
			status, err := getValidatorStatus(pubKey, blockNumber)
			if err != nil {
				fmt.Printf("Error fetching validator status for %s: %v\n", pubKey, err)
				return
			}
			statusChan <- status
		}(pubKey)
	}

	go func() {
		wg.Wait()
		close(statusChan)
	}()

	for status := range statusChan {
		if status == status {
			count++
		}
	}

	return count, nil
}

func submitValidatorStats() {
	blockNumber := GetLatestReportableBlock()
	validatorPubKeys, err := getValidatorPubKeys()
	activeValidatorsCount, err := countValidatorsByStatus(validatorPubKeys, blockNumber, "active_exiting")
	exitedValidatorsCount, err := countValidatorsByStatus(validatorPubKeys, blockNumber, "exited_unslashed")
	exitedValidatorsCount, err += countValidatorsByStatus(validatorPubKeys, blockNumber, "withdrawal_possible")
	slashedValidatorsCount, err := countValidatorsByStatus(validatorPubKeys, blockNumber, "exited_slashed")
	slashedValidatorsCount, err += countValidatorsByStatus(validatorPubKeys, blockNumber, "active_slashed")
	
	// Call the submitValidatorCounts function
	tx, err := contractInstance.SubmitValidatorCounts(auth, blockNumber, activeValidatorsCount, exitedValidatorsCount, slashedValidatorsCount)
	if err != nil {
		fmt.Errorf("Failed to submit validator counts: %v", err)
	}

	fmt.Printf("Transaction sent: %s\n", tx.Hash().Hex())
}

// Get the latest block number to report RPL price for
func (t *submitValidatorStats) getLatestReportableBlock() (uint64, error) {

	// Require eth client synced
	if err := services.RequireEthClientSynced(t.c); err != nil {
		return 0, err
	}

	latestBlock, err := GetLatestReportableBlock(t.rp, nil)
	if err != nil {
		return 0, fmt.Errorf("Error getting latest reportable block: %w", err)
	}
	return latestBlock.Uint64(), nil

}

// Returns the latest block number that oracles should be reporting prices for
func GetLatestReportableBlock() (*big.Int, error) {
	staderOracle, err := GetContract("StaderOracle")
	if err != nil {
		return nil, err
	}
	latestReportableBlock := new(*big.Int)
	if err := staderOracle.Call(opts, latestReportableBlock, "getLatestReportableBlock"); err != nil {
		return nil, fmt.Errorf("Could not get latest reportable block: %w", err)
	}
	return *latestReportableBlock, nil
}

func GetContract(contractName string) (*Contract, error) {
	// Data
	var wg errgroup.Group
	var address *common.Address
	var abi *abi.ABI

	// Create contract
	contract := &Contract{
		Contract: bind.NewBoundContract(*address, *abi),
		Address:  address,
		ABI:      abi,
		Client:   rp.Client,
	}

	// Return
	return contract, nil
}