package collector

import (
	"fmt"

	"github.com/stader-labs/stader-node/stader-lib/stader"

	"github.com/ethereum/go-ethereum/common"
	"github.com/stader-labs/stader-node/shared/services/beacon"

	"github.com/prometheus/client_golang/prometheus"
)

// Represents the collector for the stader network metrics
type RewardCollector struct {
	ActiveValidators  *prometheus.Desc
	QueuedValidators  *prometheus.Desc
	SlashedValidators *prometheus.Desc
	TotalETHBonded    *prometheus.Desc
	TotalSDBonded     *prometheus.Desc

	// The beacon client
	bc beacon.Client

	// The eth1 client
	ec stader.ExecutionClient

	// The node's address
	nodeAddress common.Address

	// The thread-safe locker for the network state
	stateLocker *StateCache

	// Prefix for logging
	logPrefix string
}

// Create a new NetworkCollector instance
func NewRewardCollector(
	bc beacon.Client,
	ec stader.ExecutionClient,
	nodeAddress common.Address,
	stateLocker *StateCache,
) *RewardCollector {
	return &RewardCollector{
		ActiveValidators:  prometheus.NewDesc(prometheus.BuildFQName(namespace, ValidatorSub, ActiveValidators), "", nil, nil),
		QueuedValidators:  prometheus.NewDesc(prometheus.BuildFQName(namespace, ValidatorSub, QueuedValidators), "", nil, nil),
		SlashedValidators: prometheus.NewDesc(prometheus.BuildFQName(namespace, ValidatorSub, SlashedValidators), "", nil, nil),
		TotalETHBonded:    prometheus.NewDesc(prometheus.BuildFQName(namespace, ValidatorSub, TotalETHBonded), "", nil, nil),
		TotalSDBonded:     prometheus.NewDesc(prometheus.BuildFQName(namespace, ValidatorSub, TotalSDBonded), "", nil, nil),
		bc:                bc,
		ec:                ec,
		nodeAddress:       nodeAddress,
		stateLocker:       stateLocker,
		logPrefix:         "Reward Collector",
	}
}

// Write metric descriptions to the Prometheus channel
func (collector *RewardCollector) Describe(channel chan<- *prometheus.Desc) {
	channel <- collector.ActiveValidators
	channel <- collector.QueuedValidators
	channel <- collector.SlashedValidators
	channel <- collector.TotalETHBonded
	channel <- collector.TotalSDBonded
}

// Collect the latest metric values and pass them to Prometheus
func (collector *RewardCollector) Collect(channel chan<- prometheus.Metric) {
	// Get the latest state
	state := collector.stateLocker.GetState()

	fmt.Printf("%+v", state)
	channel <- prometheus.MustNewConstMetric(collector.ActiveValidators, prometheus.GaugeValue, float64(9696969))
	channel <- prometheus.MustNewConstMetric(collector.QueuedValidators, prometheus.GaugeValue, float64(9696969))
	channel <- prometheus.MustNewConstMetric(collector.SlashedValidators, prometheus.GaugeValue, float64(9696969))
	channel <- prometheus.MustNewConstMetric(collector.TotalETHBonded, prometheus.GaugeValue, float64(9696969))
	channel <- prometheus.MustNewConstMetric(collector.TotalSDBonded, prometheus.GaugeValue, float64(9696969))
}

// Log error messages
func (collector *RewardCollector) logError(err error) {
	fmt.Printf("[%s] %s\n", collector.logPrefix, err.Error())
}
