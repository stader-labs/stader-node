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
	ActiveValidators                  *prometheus.Desc
	QueuedValidators                  *prometheus.Desc
	SlashedValidators                 *prometheus.Desc
	TotalETHBonded                    *prometheus.Desc
	TotalSDBonded                     *prometheus.Desc
	SdCollateral                      *prometheus.Desc
	BeaconchainReward                 *prometheus.Desc
	ElReward                          *prometheus.Desc
	SDReward                          *prometheus.Desc
	ETHAPR                            *prometheus.Desc
	SDAPR                             *prometheus.Desc
	CumulativePenalty                 *prometheus.Desc
	ClaimedBeaconchainRewards         *prometheus.Desc
	ClaimedELRewards                  *prometheus.Desc
	ClaimedSDrewards                  *prometheus.Desc
	UnclaimedELRewards                *prometheus.Desc
	UnclaimedSDRewards                *prometheus.Desc
	NextSDOrELAndSDRewardsCheckpoint  *prometheus.Desc
	TotalAttestations                 *prometheus.Desc
	AttestationPercent                *prometheus.Desc
	BlocksProduced                    *prometheus.Desc
	BlocksProducedPercent             *prometheus.Desc
	AttestationInclusionEffectiveness *prometheus.Desc
	UptimePercent                     *prometheus.Desc

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
		ActiveValidators: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, RewardSub, ActiveValidators), "", nil, nil,
		),
		QueuedValidators: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, RewardSub, QueuedValidators), "", nil, nil,
		),
		SlashedValidators: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, RewardSub, SlashedValidators), "", nil, nil,
		),
		TotalETHBonded: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, RewardSub, TotalETHBonded), "", nil, nil,
		),
		TotalSDBonded: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, RewardSub, TotalSDBonded), "", nil, nil,
		),
		SdCollateral: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, RewardSub, SdCollateral), "", nil, nil),
		BeaconchainReward: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, RewardSub, BeaconchainReward), "", nil, nil),
		ElReward: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, RewardSub, ElReward), "", nil, nil),
		SDReward: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, RewardSub, SDReward), "", nil, nil),
		ETHAPR: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, RewardSub, ETHAPR), "", nil, nil),
		SDAPR: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, RewardSub, SDAPR), "", nil, nil),
		CumulativePenalty: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, RewardSub, CumulativePenalty), "", nil, nil),
		ClaimedBeaconchainRewards: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, RewardSub, ClaimedBeaconchainRewards), "", nil, nil),
		ClaimedELRewards: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, RewardSub, ClaimedELRewards), "", nil, nil),
		ClaimedSDrewards: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, RewardSub, ClaimedSDrewards), "", nil, nil),
		UnclaimedELRewards: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, RewardSub, UnclaimedELRewards), "", nil, nil),
		UnclaimedSDRewards: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, RewardSub, UnclaimedSDRewards), "", nil, nil),
		NextSDOrELAndSDRewardsCheckpoint: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, RewardSub, NextSDOrELAndSDRewardsCheckpoint), "", nil, nil),
		TotalAttestations: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, RewardSub, TotalAttestations), "", nil, nil),
		AttestationPercent: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, RewardSub, AttestationPercent), "", nil, nil),
		BlocksProduced: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, RewardSub, BlocksProduced), "", nil, nil),
		BlocksProducedPercent: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, RewardSub, BlocksProducedPercent), "", nil, nil),
		AttestationInclusionEffectiveness: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, RewardSub, AttestationInclusionEffectiveness), "", nil, nil),
		UptimePercent: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, RewardSub, UptimePercent), "", nil, nil),

		bc:          bc,
		ec:          ec,
		nodeAddress: nodeAddress,
		stateLocker: stateLocker,
		logPrefix:   "Reward Collector",
	}
}

// Write metric descriptions to the Prometheus channel
func (collector *RewardCollector) Describe(channel chan<- *prometheus.Desc) {
	channel <- collector.ActiveValidators
	channel <- collector.QueuedValidators
	channel <- collector.SlashedValidators
	channel <- collector.TotalETHBonded
	channel <- collector.TotalSDBonded

	channel <- collector.SdCollateral
	channel <- collector.BeaconchainReward
	channel <- collector.ElReward
	channel <- collector.SDReward
	channel <- collector.ETHAPR
	channel <- collector.SDAPR
	channel <- collector.CumulativePenalty
	channel <- collector.ClaimedBeaconchainRewards
	channel <- collector.ClaimedELRewards
	channel <- collector.ClaimedSDrewards
	channel <- collector.UnclaimedELRewards
	channel <- collector.UnclaimedSDRewards
	channel <- collector.NextSDOrELAndSDRewardsCheckpoint
	channel <- collector.TotalAttestations
	channel <- collector.AttestationPercent
	channel <- collector.BlocksProduced
	channel <- collector.BlocksProducedPercent
	channel <- collector.AttestationInclusionEffectiveness
	channel <- collector.UptimePercent
}

// Collect the latest metric values and pass them to Prometheus
func (collector *RewardCollector) Collect(channel chan<- prometheus.Metric) {
	// Get the latest state
	state := collector.stateLocker.GetState()

	fmt.Printf("%+v", state)
	channel <- prometheus.MustNewConstMetric(collector.ActiveValidators, prometheus.GaugeValue, float64(state.StaderNetworkDetails.ActiveValidators.Int64()))
	channel <- prometheus.MustNewConstMetric(collector.QueuedValidators, prometheus.GaugeValue, float64(state.StaderNetworkDetails.QueuedValidators.Int64()))
	channel <- prometheus.MustNewConstMetric(collector.SlashedValidators, prometheus.GaugeValue, float64(state.StaderNetworkDetails.SlashedValidators.Int64()))
	channel <- prometheus.MustNewConstMetric(collector.TotalETHBonded, prometheus.GaugeValue, float64(state.StaderNetworkDetails.TotalETHBonded.Int64()))
	channel <- prometheus.MustNewConstMetric(collector.TotalSDBonded, prometheus.GaugeValue, float64(state.StaderNetworkDetails.TotalSDBonded.Int64()))
	channel <- prometheus.MustNewConstMetric(collector.SdCollateral, prometheus.GaugeValue, float64(state.StaderNetworkDetails.SdCollateral.Int64()))
	channel <- prometheus.MustNewConstMetric(collector.BeaconchainReward, prometheus.GaugeValue, float64(state.StaderNetworkDetails.BeaconchainReward.Int64()))
	channel <- prometheus.MustNewConstMetric(collector.ElReward, prometheus.GaugeValue, float64(state.StaderNetworkDetails.ElReward.Int64()))
	channel <- prometheus.MustNewConstMetric(collector.SDReward, prometheus.GaugeValue, float64(state.StaderNetworkDetails.SDReward.Int64()))
	channel <- prometheus.MustNewConstMetric(collector.ETHAPR, prometheus.GaugeValue, float64(state.StaderNetworkDetails.ETHAPR.Int64()))
	channel <- prometheus.MustNewConstMetric(collector.SDAPR, prometheus.GaugeValue, float64(state.StaderNetworkDetails.SDAPR.Int64()))
	channel <- prometheus.MustNewConstMetric(collector.CumulativePenalty, prometheus.GaugeValue, float64(state.StaderNetworkDetails.CumulativePenalty.Int64()))
	channel <- prometheus.MustNewConstMetric(collector.ClaimedBeaconchainRewards, prometheus.GaugeValue, float64(state.StaderNetworkDetails.ClaimedBeaconchainRewards.Int64()))
	channel <- prometheus.MustNewConstMetric(collector.ClaimedELRewards, prometheus.GaugeValue, float64(state.StaderNetworkDetails.ClaimedELRewards.Int64()))
	channel <- prometheus.MustNewConstMetric(collector.ClaimedSDrewards, prometheus.GaugeValue, float64(state.StaderNetworkDetails.ClaimedSDrewards.Int64()))
	channel <- prometheus.MustNewConstMetric(collector.UnclaimedELRewards, prometheus.GaugeValue, float64(state.StaderNetworkDetails.UnclaimedELRewards.Int64()))
	channel <- prometheus.MustNewConstMetric(collector.UnclaimedSDRewards, prometheus.GaugeValue, float64(state.StaderNetworkDetails.UnclaimedSDRewards.Int64()))
	channel <- prometheus.MustNewConstMetric(collector.NextSDOrELAndSDRewardsCheckpoint, prometheus.GaugeValue, float64(state.StaderNetworkDetails.NextSDOrELAndSDRewardsCheckpoint.Int64()))
	channel <- prometheus.MustNewConstMetric(collector.TotalAttestations, prometheus.GaugeValue, float64(state.StaderNetworkDetails.TotalAttestations.Int64()))
	channel <- prometheus.MustNewConstMetric(collector.AttestationPercent, prometheus.GaugeValue, float64(state.StaderNetworkDetails.AttestationPercent.Int64()))
	channel <- prometheus.MustNewConstMetric(collector.BlocksProduced, prometheus.GaugeValue, float64(state.StaderNetworkDetails.BlocksProduced.Int64()))
	channel <- prometheus.MustNewConstMetric(collector.BlocksProducedPercent, prometheus.GaugeValue, float64(state.StaderNetworkDetails.BlocksProducedPercent.Int64()))
	channel <- prometheus.MustNewConstMetric(collector.AttestationInclusionEffectiveness, prometheus.GaugeValue, float64(state.StaderNetworkDetails.AttestationInclusionEffectiveness.Int64()))
	channel <- prometheus.MustNewConstMetric(collector.UptimePercent, prometheus.GaugeValue, float64(state.StaderNetworkDetails.UptimePercent.Int64()))
}

// Log error messages
func (collector *RewardCollector) logError(err error) {
	fmt.Printf("[%s] %s\n", collector.logPrefix, err.Error())
}
