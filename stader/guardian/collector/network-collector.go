package collector

import (
	"fmt"
	"github.com/stader-labs/stader-node/stader-lib/stader"
	"github.com/stader-labs/stader-node/stader-lib/utils/eth"

	"github.com/ethereum/go-ethereum/common"
	"github.com/stader-labs/stader-node/shared/services/beacon"

	"github.com/prometheus/client_golang/prometheus"
)

// Represents the collector for the stader network metrics
type NetworkCollector struct {
	// The current SD price
	SdPrice *prometheus.Desc

	// The total number of validators created
	TotalValidatorsCreated *prometheus.Desc

	// The total number of registered operators
	TotalOperators *prometheus.Desc

	// Total SD staked as collateral
	TotalStakedSd *prometheus.Desc

	// Total Eth staked by Users
	TotalStakedEthByUsers *prometheus.Desc

	// Total Eth staked by NOs
	TotalStakedEthByNos *prometheus.Desc

	// Total EthX supply
	TotalEthxSupply *prometheus.Desc

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
func NewNetworkCollector(bc beacon.Client, ec stader.ExecutionClient, nodeAddress common.Address, stateLocker *StateCache) *NetworkCollector {
	subsystem := "network"
	return &NetworkCollector{
		SdPrice: prometheus.NewDesc(prometheus.BuildFQName(namespace, subsystem, "sd_price"),
			"The current SD price",
			nil, nil,
		),
		TotalValidatorsCreated: prometheus.NewDesc(prometheus.BuildFQName(namespace, subsystem, "total_validators_created"),
			"The total number of validators created in the Stader network",
			nil, nil,
		),
		TotalOperators: prometheus.NewDesc(prometheus.BuildFQName(namespace, subsystem, "total_operator_registered"),
			"The total number of operator registered in the Stader network",
			nil, nil,
		),
		TotalStakedSd: prometheus.NewDesc(prometheus.BuildFQName(namespace, subsystem, "total_staked_sd_collateral"),
			"The total amount of SD staked as collateral",
			nil, nil,
		),
		TotalStakedEthByUsers: prometheus.NewDesc(prometheus.BuildFQName(namespace, subsystem, "total_staked_user_eth"),
			"The total amount of SD staked as collateral",
			nil, nil,
		),
		TotalStakedEthByNos: prometheus.NewDesc(prometheus.BuildFQName(namespace, subsystem, "total_staked_nos_eth"),
			"The total amount of SD staked as collateral",
			nil, nil,
		),
		TotalEthxSupply: prometheus.NewDesc(prometheus.BuildFQName(namespace, subsystem, "total_ethx_supply"),
			"The total amount of SD staked as collateral",
			nil, nil,
		),
		bc:          bc,
		ec:          ec,
		nodeAddress: nodeAddress,
		stateLocker: stateLocker,
		logPrefix:   "Network Collector",
	}
}

// Write metric descriptions to the Prometheus channel
func (collector *NetworkCollector) Describe(channel chan<- *prometheus.Desc) {
	channel <- collector.SdPrice
	channel <- collector.TotalValidatorsCreated
	channel <- collector.TotalOperators
	channel <- collector.TotalEthxSupply
	channel <- collector.TotalStakedEthByUsers
	channel <- collector.TotalStakedEthByNos
	channel <- collector.TotalStakedSd
}

// Collect the latest metric values and pass them to Prometheus
func (collector *NetworkCollector) Collect(channel chan<- prometheus.Metric) {
	// Get the latest state
	state := collector.stateLocker.GetState()

	channel <- prometheus.MustNewConstMetric(
		collector.SdPrice, prometheus.GaugeValue, eth.WeiToEth(state.StaderNetworkDetails.SdPrice))
	channel <- prometheus.MustNewConstMetric(
		collector.TotalValidatorsCreated, prometheus.GaugeValue, float64(state.StaderNetworkDetails.TotalValidators.Int64()))
	channel <- prometheus.MustNewConstMetric(
		collector.TotalOperators, prometheus.GaugeValue, float64(state.StaderNetworkDetails.TotalOperators.Int64()))
	channel <- prometheus.MustNewConstMetric(
		collector.TotalEthxSupply, prometheus.GaugeValue, float64(state.StaderNetworkDetails.TotalEthxSupply.Int64()))
	channel <- prometheus.MustNewConstMetric(
		collector.TotalStakedEthByUsers, prometheus.GaugeValue, float64(state.StaderNetworkDetails.TotalStakedEthByUsers.Int64()))
	channel <- prometheus.MustNewConstMetric(
		collector.TotalStakedEthByNos, prometheus.GaugeValue, float64(state.StaderNetworkDetails.TotalStakedEthByNos.Int64()))
	channel <- prometheus.MustNewConstMetric(
		collector.TotalStakedSd, prometheus.GaugeValue, float64(state.StaderNetworkDetails.TotalStakedSd.Int64()))
}

// Log error messages
func (collector *NetworkCollector) logError(err error) {
	fmt.Printf("[%s] %s\n", collector.logPrefix, err.Error())
}
