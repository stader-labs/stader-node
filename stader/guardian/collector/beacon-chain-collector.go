/*
This work is licensed and released under GNU GPL v3 or any other later versions.
The full text of the license is below/ found at <http://www.gnu.org/licenses/>

(c) 2023 Rocket Pool Pty Ltd. Modified under GNU GPL v3. [0.4.0-beta]

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

package collector

import (
	"fmt"
	"github.com/stader-labs/stader-node/stader-lib/stader"

	"github.com/ethereum/go-ethereum/common"
	"github.com/stader-labs/stader-node/shared/services/beacon"

	"github.com/prometheus/client_golang/prometheus"
	"golang.org/x/sync/errgroup"
)

// Represents the collector for the beaconchain metrics
type BeaconCollector struct {
	// The number of this node's validators is currently in a sync committee
	activeSyncCommittee *prometheus.Desc

	// The number of this node's validators on the next sync committee
	upcomingSyncCommittee *prometheus.Desc

	// The number of upcoming proposals for this node's validators
	upcomingProposals *prometheus.Desc

	// The beacon client
	bc beacon.Client

	// The eth1 client
	ec stader.ExecutionClient

	// The node's address
	nodeAddress common.Address

	// The thread-safe locker for the network state
	stateLocker *MetricsCacheContainer

	// Prefix for logging
	logPrefix string
}

// Create a new NetworkCollector instance
func NewBeaconCollector(bc beacon.Client, ec stader.ExecutionClient, nodeAddress common.Address, stateLocker *MetricsCacheContainer) *BeaconCollector {
	subsystem := "beacon"
	return &BeaconCollector{
		activeSyncCommittee: prometheus.NewDesc(prometheus.BuildFQName(namespace, subsystem, "active_sync_committee"),
			"The number of validators on a current sync committee",
			nil, nil,
		),
		upcomingSyncCommittee: prometheus.NewDesc(prometheus.BuildFQName(namespace, subsystem, "upcoming_sync_committee"),
			"The number of validators on the next sync committee",
			nil, nil,
		),
		upcomingProposals: prometheus.NewDesc(prometheus.BuildFQName(namespace, subsystem, "upcoming_proposals"),
			"The number of proposals assigned to validators in this epoch and the next",
			nil, nil,
		),
		bc:          bc,
		ec:          ec,
		nodeAddress: nodeAddress,
		stateLocker: stateLocker,
		logPrefix:   "Beacon Collector",
	}
}

// Write metric descriptions to the Prometheus channel
func (collector *BeaconCollector) Describe(channel chan<- *prometheus.Desc) {
	channel <- collector.activeSyncCommittee
	channel <- collector.upcomingSyncCommittee
	channel <- collector.upcomingProposals
}

// Collect the latest metric values and pass them to Prometheus
func (collector *BeaconCollector) Collect(channel chan<- prometheus.Metric) {
	// Get the latest state
	state := collector.stateLocker.GetMetricsContainer()

	var wg errgroup.Group
	activeSyncCommittee := float64(0)
	upcomingSyncCommittee := float64(0)
	upcomingProposals := float64(0)

	var validatorIndices []uint64
	var head beacon.BeaconHead

	// Get sync committee duties
	for _, validator := range state.ValidatorDetails {
		if validator.Exists {
			validatorIndices = append(validatorIndices, validator.Index)
		}
	}

	head, err := collector.bc.GetBeaconHead()
	if err != nil {
		collector.logError(fmt.Errorf("error getting Beacon chain head: %w", err))
		return
	}

	wg.Go(func() error {
		// Get current duties
		duties, err := collector.bc.GetValidatorSyncDuties(validatorIndices, head.Epoch)
		if err != nil {
			return fmt.Errorf("Error getting sync duties: %w", err)
		}

		for _, duty := range duties {
			if duty {
				activeSyncCommittee++
			}
		}

		return nil
	})

	wg.Go(func() error {
		// Get epochs per sync committee period config to query next period
		config := state.BeaconConfig

		// Get upcoming duties
		duties, err := collector.bc.GetValidatorSyncDuties(validatorIndices, head.Epoch+config.EpochsPerSyncCommitteePeriod)
		if err != nil {
			return fmt.Errorf("Error getting sync duties: %w", err)
		}

		for _, duty := range duties {
			if duty {
				upcomingSyncCommittee++
			}
		}

		return nil
	})

	wg.Go(func() error {
		// Get proposals in this epoch
		duties, err := collector.bc.GetValidatorProposerDuties(validatorIndices, head.Epoch)
		if err != nil {
			return fmt.Errorf("Error getting proposer duties: %w", err)
		}

		for _, duty := range duties {
			upcomingProposals += float64(duty)
		}

		return nil
	})

	// Wait for data
	if err := wg.Wait(); err != nil {
		collector.logError(err)
		return
	}

	channel <- prometheus.MustNewConstMetric(
		collector.activeSyncCommittee, prometheus.GaugeValue, activeSyncCommittee)
	channel <- prometheus.MustNewConstMetric(
		collector.upcomingSyncCommittee, prometheus.GaugeValue, upcomingSyncCommittee)
	channel <- prometheus.MustNewConstMetric(
		collector.upcomingProposals, prometheus.GaugeValue, upcomingProposals)

}

// Log error messages
func (collector *BeaconCollector) logError(err error) {
	fmt.Printf("[%s] %s\n", collector.logPrefix, err.Error())
}
