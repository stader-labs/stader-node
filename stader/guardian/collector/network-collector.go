package collector

import (
	"fmt"
	"github.com/stader-labs/stader-node/shared/utils/math"

	"github.com/ethereum/go-ethereum/common"
	"github.com/stader-labs/stader-node/shared/services/beacon"
	"github.com/stader-labs/stader-node/stader-lib/stader"

	"github.com/prometheus/client_golang/prometheus"
)

// Represents the collector for the stader network metrics
type NetworkCollector struct {
	// The current SD price in Eth
	SdPrice *prometheus.Desc

	// The current Eth price in SD
	EthPrice *prometheus.Desc

	// The total number of validators created
	TotalValidatorsCreated *prometheus.Desc

	// The total number of validators active on beacon chain
	TotalActiveValidators *prometheus.Desc

	// The total number of validators waiting to receive the 28eth
	TotalQueuedValidators *prometheus.Desc

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

	// The next block at which Sd and socializing el rewards wil be given
	NextRewardBlock *prometheus.Desc

	// The operator collateral ratio in ETH
	CollateralRatio *prometheus.Desc

	// The operator collateral ratio in SD
	CollateralRatioInSd *prometheus.Desc

	// The max amount of sd value that can be staked to get rewards
	MaxEthThreshold *prometheus.Desc

	// The min amount of sd value that can be staked to get rewards
	MinEthThreshold *prometheus.Desc

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
		EthPrice: prometheus.NewDesc(prometheus.BuildFQName(namespace, subsystem, "eth_price"),
			"The current Eth price",
			nil, nil,
		),
		TotalValidatorsCreated: prometheus.NewDesc(prometheus.BuildFQName(namespace, subsystem, "total_validators_created"),
			"The total number of validators created in the Stader network",
			nil, nil,
		),
		TotalQueuedValidators: prometheus.NewDesc(prometheus.BuildFQName(namespace, subsystem, "total_validators_queued"),
			"The total number of validators waiting to receive the 28eth",
			nil, nil,
		),
		TotalActiveValidators: prometheus.NewDesc(prometheus.BuildFQName(namespace, subsystem, "total_validators_active"),
			"The total number of validators active on beacon chain",
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
			"The total amount of Eth staked by users",
			nil, nil,
		),
		TotalStakedEthByNos: prometheus.NewDesc(prometheus.BuildFQName(namespace, subsystem, "total_staked_nos_eth"),
			"The total amount of Eth staked by NOs",
			nil, nil,
		),
		TotalEthxSupply: prometheus.NewDesc(prometheus.BuildFQName(namespace, subsystem, "total_ethx_supply"),
			"The total Ethx Supply",
			nil, nil,
		),
		NextRewardBlock: prometheus.NewDesc(prometheus.BuildFQName(namespace, subsystem, "next_reward_block"),
			"The next block at which SD and socializing el rewards wil be given",
			nil, nil,
		),
		CollateralRatio: prometheus.NewDesc(prometheus.BuildFQName(namespace, subsystem, "collateral_ratio"),
			"The collateral ratio for adding a new validator in Eth",
			nil, nil,
		),
		CollateralRatioInSd: prometheus.NewDesc(prometheus.BuildFQName(namespace, subsystem, "collateral_ratio_sd"),
			"The collateral ratio for adding a new validator in SD",
			nil, nil,
		),
		MinEthThreshold: prometheus.NewDesc(prometheus.BuildFQName(namespace, subsystem, "min_eth_threshold"),
			"The minimum amount of sd value that can be staked to get rewards",
			nil, nil,
		),
		MaxEthThreshold: prometheus.NewDesc(prometheus.BuildFQName(namespace, subsystem, "max_eth_threshold"),
			"The maximum amount of sd value that can be staked to get rewards",
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
	channel <- collector.EthPrice
	channel <- collector.TotalValidatorsCreated
	channel <- collector.TotalQueuedValidators
	channel <- collector.TotalActiveValidators
	channel <- collector.TotalOperators
	channel <- collector.TotalEthxSupply
	channel <- collector.TotalStakedEthByUsers
	channel <- collector.TotalStakedEthByNos
	channel <- collector.TotalStakedSd
	channel <- collector.NextRewardBlock
	channel <- collector.CollateralRatio
	channel <- collector.CollateralRatioInSd
	channel <- collector.MinEthThreshold
	channel <- collector.MaxEthThreshold
}

// Collect the latest metric values and pass them to Prometheus
func (collector *NetworkCollector) Collect(channel chan<- prometheus.Metric) {
	// Get the latest state
	state := collector.stateLocker.GetState()

	currentStartBlock := math.RoundDown(float64(state.StaderNetworkDetails.NextSocializingPoolRewardCycle.CurrentStartBlock.Int64()), 0)

	channel <- prometheus.MustNewConstMetric(
		collector.SdPrice, prometheus.GaugeValue, state.StaderNetworkDetails.SdPrice)
	channel <- prometheus.MustNewConstMetric(
		collector.EthPrice, prometheus.GaugeValue, state.StaderNetworkDetails.EthPrice)
	channel <- prometheus.MustNewConstMetric(
		collector.TotalValidatorsCreated, prometheus.GaugeValue, float64(state.StaderNetworkDetails.TotalValidators.Int64()))
	channel <- prometheus.MustNewConstMetric(
		collector.TotalActiveValidators, prometheus.GaugeValue, float64(state.StaderNetworkDetails.TotalActiveValidators.Int64()))
	channel <- prometheus.MustNewConstMetric(
		collector.TotalQueuedValidators, prometheus.GaugeValue, float64(state.StaderNetworkDetails.TotalQueuedValidators.Int64()))
	channel <- prometheus.MustNewConstMetric(
		collector.TotalOperators, prometheus.GaugeValue, float64(state.StaderNetworkDetails.TotalOperators.Int64()))
	channel <- prometheus.MustNewConstMetric(
		collector.TotalEthxSupply, prometheus.GaugeValue, state.StaderNetworkDetails.TotalEthxSupply)
	channel <- prometheus.MustNewConstMetric(
		collector.TotalStakedEthByUsers, prometheus.GaugeValue, float64(state.StaderNetworkDetails.TotalStakedEthByUsers.Int64()))
	channel <- prometheus.MustNewConstMetric(
		collector.TotalStakedEthByNos, prometheus.GaugeValue, float64(state.StaderNetworkDetails.TotalStakedEthByNos.Int64()))
	channel <- prometheus.MustNewConstMetric(
		collector.TotalStakedSd, prometheus.GaugeValue, state.StaderNetworkDetails.TotalStakedSd)
	channel <- prometheus.MustNewConstMetric(
		collector.NextRewardBlock, prometheus.GaugeValue, currentStartBlock)
	channel <- prometheus.MustNewConstMetric(
		collector.CollateralRatio, prometheus.GaugeValue, state.StaderNetworkDetails.CollateralRatio)
	channel <- prometheus.MustNewConstMetric(
		collector.CollateralRatioInSd, prometheus.GaugeValue, state.StaderNetworkDetails.CollateralRatioInSd)
	channel <- prometheus.MustNewConstMetric(
		collector.MinEthThreshold, prometheus.GaugeValue, state.StaderNetworkDetails.MinEthThreshold)
	channel <- prometheus.MustNewConstMetric(
		collector.MaxEthThreshold, prometheus.GaugeValue, state.StaderNetworkDetails.MaxEthThreshold)
}

// Log error messages
func (collector *NetworkCollector) logError(err error) {
	fmt.Printf("[%s] %s\n", collector.logPrefix, err.Error())
}
