package collectors

import (
	"fmt"
	"log"

	"github.com/ethereum/go-ethereum/common"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/rocket-pool/rocketpool-go/rocketpool"
	"github.com/stader-labs/stader-node/shared/services/beacon"
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

	// The Rocket Pool contract manager
	rp *rocketpool.RocketPool

	// The beacon client
	bc beacon.Client

	// The eth1 client
	ec rocketpool.ExecutionClient

	// The node's address
	nodeAddress common.Address
}

// Create a new BeaconCollector instance
func NewBeaconCollector(rp *rocketpool.RocketPool, bc beacon.Client, ec rocketpool.ExecutionClient, nodeAddress common.Address) *BeaconCollector {
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
		rp:          rp,
		bc:          bc,
		ec:          ec,
		nodeAddress: nodeAddress,
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

	// Sync
	var wg errgroup.Group
	var wg2 errgroup.Group

	activeSyncCommittee := float64(0)
	upcomingSyncCommittee := float64(0)
	upcomingProposals := float64(0)

	var validatorIndices []uint64
	var head beacon.BeaconHead

	// TODO - check this - bchain
	//// Get sync committee duties
	//wg.Go(func() error {
	//	var err error
	//	validatorIndices, err = rp.GetNodeValidatorIndices(collector.rp, collector.ec, collector.bc, collector.nodeAddress)
	//	if err != nil {
	//		return fmt.Errorf("Error getting validator indices: %w", err)
	//	}
	//	return nil
	//})

	wg.Go(func() error {
		var err error
		head, err = collector.bc.GetBeaconHead()
		if err != nil {
			return fmt.Errorf("Error getting beaconchain head: %w", err)
		}
		return nil
	})

	// Wait for data
	if err := wg.Wait(); err != nil {
		log.Printf("%s\n", err.Error())
		return
	}

	wg2.Go(func() error {
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

	wg2.Go(func() error {
		// Get epochs per sync committee period config to query next period
		config, err := collector.bc.GetEth2Config()
		if err != nil {
			return fmt.Errorf("Error getting ETH2 config: %w", err)
		}

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

	wg2.Go(func() error {
		// Get proposals in this epoch
		duties, err := collector.bc.GetValidatorProposerDuties(validatorIndices, head.Epoch)
		if err != nil {
			return fmt.Errorf("Error getting proposer duties: %w", err)
		}

		for _, duty := range duties {
			upcomingProposals += float64(duty)
		}

		// TODO: this seems to be illegal according to the official spec:
		// https://eth2book.info/altair/annotated-spec/#compute_proposer_index
		/*
			// Get proposals in the next epoch
			duties, err = collector.bc.GetValidatorProposerDuties(validatorIndices, head.Epoch + 1)
			if err != nil {
				return fmt.Errorf("Error getting proposer duties: %w", err)
			}

			for _, duty := range duties {
				upcomingProposals += float64(duty)
			}
		*/

		return nil
	})

	// Wait for data
	if err := wg2.Wait(); err != nil {
		log.Printf("%s\n", err.Error())
		return
	}

	channel <- prometheus.MustNewConstMetric(
		collector.activeSyncCommittee, prometheus.GaugeValue, activeSyncCommittee)
	channel <- prometheus.MustNewConstMetric(
		collector.upcomingSyncCommittee, prometheus.GaugeValue, upcomingSyncCommittee)
	channel <- prometheus.MustNewConstMetric(
		collector.upcomingProposals, prometheus.GaugeValue, upcomingProposals)
}
