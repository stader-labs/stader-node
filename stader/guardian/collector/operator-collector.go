package collector

import (
	"fmt"

	"github.com/stader-labs/stader-node/stader-lib/stader"

	"github.com/ethereum/go-ethereum/common"
	"github.com/stader-labs/stader-node/shared/services/beacon"

	"github.com/prometheus/client_golang/prometheus"
)

// Represents the collector for the stader network metrics
type OperatorCollector struct {
	ActiveValidators                     *prometheus.Desc
	QueuedValidators                     *prometheus.Desc
	SlashedValidators                    *prometheus.Desc
	ExitingValidators                    *prometheus.Desc
	WithdrawnValidators                  *prometheus.Desc
	FrontRunValidators                   *prometheus.Desc
	InvalidSignatureValidators           *prometheus.Desc
	IntializedValidators                 *prometheus.Desc
	UnclaimedClRewards                   *prometheus.Desc
	UnclaimedNonSocializingPoolElRewards *prometheus.Desc
	CumulativePenalty                    *prometheus.Desc
	UnclaimedSocializingPoolELRewards    *prometheus.Desc
	UnclaimedSocializingPoolSdRewards    *prometheus.Desc
	ClaimedSocializingPoolSdRewards      *prometheus.Desc
	ClaimedSocializingPoolElRewards      *prometheus.Desc
	TotalSdCollateral                    *prometheus.Desc
	TotalEthColateral                    *prometheus.Desc

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
func NewOperatorCollector(
	bc beacon.Client,
	ec stader.ExecutionClient,
	nodeAddress common.Address,
	stateLocker *StateCache,
) *OperatorCollector {
	return &OperatorCollector{
		ActiveValidators: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, OperatorSub, ActiveValidators), "", nil, nil,
		),
		QueuedValidators: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, OperatorSub, QueuedValidators), "", nil, nil,
		),
		SlashedValidators: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, OperatorSub, SlashedValidators), "", nil, nil,
		),
		ExitingValidators: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, OperatorSub, ExitingValidators), "", nil, nil,
		),
		WithdrawnValidators: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, OperatorSub, WithdrawnValidators), "", nil, nil,
		),
		InvalidSignatureValidators: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, OperatorSub, InvalidSignatureValidators), "", nil, nil,
		),
		FrontRunValidators: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, OperatorSub, FrontRunValidators), "", nil, nil,
		),
		IntializedValidators: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, OperatorSub, InitializedValidators), "", nil, nil,
		),
		UnclaimedClRewards: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, OperatorSub, UnclaimedCLRewards), "", nil, nil,
		),
		UnclaimedNonSocializingPoolElRewards: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, OperatorSub, UnclaimedNonSocializingPoolELRewards), "", nil, nil,
		),
		CumulativePenalty: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, OperatorSub, CumulativePenalty), "", nil, nil),
		UnclaimedSocializingPoolELRewards: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, OperatorSub, UnclaimedSocializingPoolELRewards), "", nil, nil,
		),
		UnclaimedSocializingPoolSdRewards: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, OperatorSub, UnclaimedSocializingPoolSdRewards), "", nil, nil,
		),
		ClaimedSocializingPoolSdRewards: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, OperatorSub, ClaimedSocializingPoolSDrewards), "", nil, nil),
		ClaimedSocializingPoolElRewards: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, OperatorSub, ClaimedSocializingPoolELRewards), "", nil, nil),
		TotalSdCollateral: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, OperatorSub, SdCollateral), "", nil, nil),
		TotalEthColateral: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, OperatorSub, EthCollateral), "", nil, nil),
		bc:          bc,
		ec:          ec,
		nodeAddress: nodeAddress,
		stateLocker: stateLocker,
		logPrefix:   "Operator Collector",
	}
}

// Write metric descriptions to the Prometheus channel
func (collector *OperatorCollector) Describe(channel chan<- *prometheus.Desc) {
	channel <- collector.ActiveValidators
	channel <- collector.QueuedValidators
	channel <- collector.SlashedValidators
	channel <- collector.ExitingValidators
	channel <- collector.WithdrawnValidators
	channel <- collector.UnclaimedClRewards
	channel <- collector.UnclaimedNonSocializingPoolElRewards
	channel <- collector.CumulativePenalty
	channel <- collector.UnclaimedSocializingPoolELRewards
	channel <- collector.UnclaimedSocializingPoolSdRewards
	channel <- collector.ClaimedSocializingPoolSdRewards
	channel <- collector.ClaimedSocializingPoolElRewards
	channel <- collector.TotalSdCollateral
	channel <- collector.TotalEthColateral
}

// Collect the latest metric values and pass them to Prometheus
func (collector *OperatorCollector) Collect(channel chan<- prometheus.Metric) {
	// Get the latest state
	state := collector.stateLocker.GetState()

	channel <- prometheus.MustNewConstMetric(collector.ActiveValidators, prometheus.GaugeValue, float64(state.StaderNetworkDetails.ActiveValidators.Int64()))
	channel <- prometheus.MustNewConstMetric(collector.QueuedValidators, prometheus.GaugeValue, float64(state.StaderNetworkDetails.QueuedValidators.Int64()))
	channel <- prometheus.MustNewConstMetric(collector.SlashedValidators, prometheus.GaugeValue, float64(state.StaderNetworkDetails.SlashedValidators.Int64()))
	channel <- prometheus.MustNewConstMetric(collector.ExitingValidators, prometheus.GaugeValue, float64(state.StaderNetworkDetails.ExitingValidators.Int64()))
	channel <- prometheus.MustNewConstMetric(collector.WithdrawnValidators, prometheus.GaugeValue, float64(state.StaderNetworkDetails.WithdrawnValidators.Int64()))
	channel <- prometheus.MustNewConstMetric(collector.UnclaimedClRewards, prometheus.GaugeValue, state.StaderNetworkDetails.UnclaimedClRewards)
	channel <- prometheus.MustNewConstMetric(collector.CumulativePenalty, prometheus.GaugeValue, state.StaderNetworkDetails.CumulativePenalty)
	channel <- prometheus.MustNewConstMetric(collector.UnclaimedNonSocializingPoolElRewards, prometheus.GaugeValue, state.StaderNetworkDetails.UnclaimedNonSocializingPoolElRewards)
	channel <- prometheus.MustNewConstMetric(collector.UnclaimedSocializingPoolELRewards, prometheus.GaugeValue, state.StaderNetworkDetails.UnclaimedSocializingPoolElRewards)
	channel <- prometheus.MustNewConstMetric(collector.UnclaimedSocializingPoolSdRewards, prometheus.GaugeValue, state.StaderNetworkDetails.UnclaimedSocializingPoolSDRewards)
	channel <- prometheus.MustNewConstMetric(collector.ClaimedSocializingPoolSdRewards, prometheus.GaugeValue, state.StaderNetworkDetails.ClaimedSocializingPoolSdRewards)
	channel <- prometheus.MustNewConstMetric(collector.ClaimedSocializingPoolElRewards, prometheus.GaugeValue, state.StaderNetworkDetails.ClaimedSocializingPoolElRewards)
	channel <- prometheus.MustNewConstMetric(collector.TotalSdCollateral, prometheus.GaugeValue, state.StaderNetworkDetails.OperatorStakedSd)
	channel <- prometheus.MustNewConstMetric(collector.TotalEthColateral, prometheus.GaugeValue, state.StaderNetworkDetails.OperatorEthCollateral)

}

// Log error messages
func (collector *OperatorCollector) logError(err error) {
	fmt.Printf("[%s] %s\n", collector.logPrefix, err.Error())
}
