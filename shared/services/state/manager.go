package state

import (
	"context"
	"fmt"
	"github.com/stader-labs/stader-node/shared/services/beacon"
	"github.com/stader-labs/stader-node/shared/services/config"
	cfgtypes "github.com/stader-labs/stader-node/shared/types/config"
	"github.com/stader-labs/stader-node/shared/utils/log"
	"github.com/stader-labs/stader-node/stader-lib/stader"
	"time"

	"github.com/ethereum/go-ethereum/common"
)

type MetricsCacheManager struct {
	cfg          *config.StaderConfig
	ec           stader.ExecutionClient
	bc           beacon.Client
	log          *log.ColorLogger
	Config       *config.StaderConfig
	Network      cfgtypes.Network
	ChainID      uint
	BeaconConfig beacon.Eth2Config
}

// Create a new manager for the network state
func NewMetricsCache(cfg *config.StaderConfig, ec stader.ExecutionClient, bc beacon.Client, log *log.ColorLogger) (*MetricsCacheManager, error) {

	// Create the manager
	m := &MetricsCacheManager{
		cfg:     cfg,
		ec:      ec,
		bc:      bc,
		log:     log,
		Config:  cfg,
		Network: cfg.StaderNode.Network.Value.(cfgtypes.Network),
		ChainID: cfg.StaderNode.GetChainID(),
	}

	// Get the Beacon config info
	var err error
	m.BeaconConfig, err = m.bc.GetEth2Config()
	if err != nil {
		return nil, err
	}

	return m, nil

}

func (m *MetricsCacheManager) GetHeadState(nodeAddress common.Address) (*MetricsCache, error) {
	targetSlot, err := m.GetHeadSlot()
	if err != nil {
		return nil, fmt.Errorf("error getting latest Beacon slot: %w", err)
	}
	return m.getNodeMetrics(nodeAddress, targetSlot)
}

func (m *MetricsCacheManager) GetHeadStateForNode(nodeAddress common.Address) (*MetricsCache, error) {
	targetSlot, err := m.GetHeadSlot()
	if err != nil {
		return nil, fmt.Errorf("error getting latest Beacon slot: %w", err)
	}
	return m.getNodeMetrics(nodeAddress, targetSlot)
}

func (m *MetricsCacheManager) GetHeadSlot() (uint64, error) {
	// Get the latest EL block
	latestBlockHeader, err := m.ec.HeaderByNumber(context.Background(), nil)
	if err != nil {
		return 0, fmt.Errorf("error getting latest EL block: %w", err)
	}

	// Get the corresponding Beacon slot based on the timestamp
	latestBlockTime := time.Unix(int64(latestBlockHeader.Time), 0)
	genesisTime := time.Unix(int64(m.BeaconConfig.GenesisTime), 0)
	secondsSinceGenesis := uint64(latestBlockTime.Sub(genesisTime).Seconds())
	targetSlot := secondsSinceGenesis / m.BeaconConfig.SecondsPerSlot
	return targetSlot, nil
}

func (m *MetricsCacheManager) getNodeMetrics(nodeAddress common.Address, slotNumber uint64) (*MetricsCache, error) {
	state, err := CreateMetricsCache(m.cfg.StaderNode, m.ec, m.bc, m.log, slotNumber, m.BeaconConfig, nodeAddress)
	if err != nil {
		return nil, err
	}
	return state, nil
}

// Logs a line if the logger is specified
func (m *MetricsCacheManager) logLine(format string, v ...interface{}) {
	if m.log != nil {
		m.log.Printlnf(format, v)
	}
}
