/*
This work is licensed and released under GNU GPL v3 or any other later versions.
The full text of the license is below/ found at <http://www.gnu.org/licenses/>

(c) 2023 Rocket Pool Pty Ltd. Modified under GNU GPL v3. [1.4.3]

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

package state

import (
	"context"
	"fmt"
	"time"

	"github.com/stader-labs/stader-node/shared/services/beacon"
	"github.com/stader-labs/stader-node/shared/services/config"
	cfgtypes "github.com/stader-labs/stader-node/shared/types/config"
	"github.com/stader-labs/stader-node/shared/utils/log"
	"github.com/stader-labs/stader-node/stader-lib/stader"
	"github.com/urfave/cli"

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
	c            *cli.Context
}

// Create a new manager for the network state
func NewMetricsCache(c *cli.Context, cfg *config.StaderConfig, ec stader.ExecutionClient, bc beacon.Client, log *log.ColorLogger) (*MetricsCacheManager, error) {

	// Create the manager
	m := &MetricsCacheManager{
		cfg:     cfg,
		ec:      ec,
		bc:      bc,
		log:     log,
		Config:  cfg,
		Network: cfg.StaderNode.Network.Value.(cfgtypes.Network),
		ChainID: cfg.StaderNode.GetChainID(),
		c:       c,
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
	state, err := CreateMetricsCache(m.c, m.cfg.StaderNode, m.ec, m.bc, m.log, slotNumber, m.BeaconConfig, nodeAddress)
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
