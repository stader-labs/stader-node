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
	BeaconChainQueuedValidators          *prometheus.Desc
	StaderQueuedValidators               *prometheus.Desc
	SlashedValidators                    *prometheus.Desc
	ExitingValidators                    *prometheus.Desc
	WithdrawnValidators                  *prometheus.Desc
	FrontRunValidators                   *prometheus.Desc
	InvalidSignatureValidators           *prometheus.Desc
	IntializedValidators                 *prometheus.Desc
	FundsSettledValidators               *prometheus.Desc
	UnclaimedClRewards                   *prometheus.Desc
	UnclaimedNonSocializingPoolElRewards *prometheus.Desc
	CumulativePenalty                    *prometheus.Desc
	UnclaimedSocializingPoolELRewards    *prometheus.Desc
	UnclaimedSocializingPoolSdRewards    *prometheus.Desc
	ClaimedSocializingPoolSdRewards      *prometheus.Desc
	ClaimedSocializingPoolElRewards      *prometheus.Desc
	TotalSdCollateral                    *prometheus.Desc
	TotalSdCollateralInEth               *prometheus.Desc
	TotalEthColateral                    *prometheus.Desc
	TotalSDUtilizationPosition           *prometheus.Desc
	TotalSDSelfBond                      *prometheus.Desc
	TotalSDUtilized                      *prometheus.Desc
	TotalSDUtilizedInterest              *prometheus.Desc
	SdCollateralPct                      *prometheus.Desc
	LockedEth                            *prometheus.Desc
	HealthFactor                         *prometheus.Desc
	LiquidationStatus                    *prometheus.Desc
	ClaimVaultBalance                    *prometheus.Desc

	SDSelfBond *prometheus.Desc

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
func NewOperatorCollector(
	bc beacon.Client,
	ec stader.ExecutionClient,
	nodeAddress common.Address,
	stateLocker *MetricsCacheContainer,
) *OperatorCollector {
	return &OperatorCollector{
		ActiveValidators: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, OperatorSub, ActiveValidators), "", nil, nil,
		),
		BeaconChainQueuedValidators: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, OperatorSub, BeaconChainQueuedValidators), "", nil, nil,
		),
		StaderQueuedValidators: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, OperatorSub, StaderQueuedValidators), "", nil, nil,
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
		FundsSettledValidators: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, OperatorSub, FundsSettledValidators), "", nil, nil,
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
		TotalSdCollateralInEth: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, OperatorSub, SdCollateralInEth), "", nil, nil),
		TotalEthColateral: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, OperatorSub, EthCollateral), "", nil, nil),
		TotalSDUtilized: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, OperatorSub, SDUtilized), "", nil, nil),
		TotalSDUtilizedInterest: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, OperatorSub, SDUtilizedInterest), "", nil, nil),
		SdCollateralPct: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, OperatorSub, SdCollateralPct), "", nil, nil),
		LockedEth: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, OperatorSub, LockedEth), "", nil, nil),
		HealthFactor: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, OperatorSub, HeathFactor), "", nil, nil),
		TotalSDUtilizationPosition: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, OperatorSub, TotalSDUtilizationPosition), "", nil, nil),
		TotalSDSelfBond: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, OperatorSub, TotalSDSelfBonded), "", nil, nil),
		LiquidationStatus: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, OperatorSub, LiquidationStatus), "", nil, nil),
		ClaimVaultBalance: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, OperatorSub, ClaimVaultBalance), "", nil, nil),
		SDSelfBond: prometheus.NewDesc(prometheus.BuildFQName(namespace, OperatorSub, "sd_self_bond"),
			"The current balance of the SD utility pool",
			nil, nil,
		),
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
	channel <- collector.BeaconChainQueuedValidators
	channel <- collector.StaderQueuedValidators
	channel <- collector.SlashedValidators
	channel <- collector.ExitingValidators
	channel <- collector.FrontRunValidators
	channel <- collector.IntializedValidators
	channel <- collector.InvalidSignatureValidators
	channel <- collector.WithdrawnValidators
	channel <- collector.FundsSettledValidators
	channel <- collector.UnclaimedClRewards
	channel <- collector.UnclaimedNonSocializingPoolElRewards
	channel <- collector.CumulativePenalty
	channel <- collector.UnclaimedSocializingPoolELRewards
	channel <- collector.UnclaimedSocializingPoolSdRewards
	channel <- collector.ClaimedSocializingPoolSdRewards
	channel <- collector.ClaimedSocializingPoolElRewards
	channel <- collector.TotalSdCollateral
	channel <- collector.TotalSdCollateralInEth
	channel <- collector.TotalEthColateral
	channel <- collector.TotalSDUtilized
	channel <- collector.TotalSDUtilizedInterest
	channel <- collector.SdCollateralPct
	channel <- collector.LockedEth
	channel <- collector.HealthFactor
	channel <- collector.TotalSDUtilizationPosition
	channel <- collector.TotalSDSelfBond
	channel <- collector.LiquidationStatus
	channel <- collector.ClaimVaultBalance
	channel <- collector.SDSelfBond
}

// Collect the latest metric values and pass them to Prometheus
func (collector *OperatorCollector) Collect(channel chan<- prometheus.Metric) {
	// Get the latest state
	state := collector.stateLocker.GetMetricsContainer()

	channel <- prometheus.MustNewConstMetric(collector.ActiveValidators, prometheus.GaugeValue, float64(state.StaderNetworkDetails.ActiveValidators.Int64()))
	channel <- prometheus.MustNewConstMetric(collector.BeaconChainQueuedValidators, prometheus.GaugeValue, float64(state.StaderNetworkDetails.BeaconChainQueuedValidators.Int64()))
	channel <- prometheus.MustNewConstMetric(collector.StaderQueuedValidators, prometheus.GaugeValue, float64(state.StaderNetworkDetails.StaderQueuedValidators.Int64()))
	channel <- prometheus.MustNewConstMetric(collector.SlashedValidators, prometheus.GaugeValue, float64(state.StaderNetworkDetails.SlashedValidators.Int64()))
	channel <- prometheus.MustNewConstMetric(collector.ExitingValidators, prometheus.GaugeValue, float64(state.StaderNetworkDetails.ExitingValidators.Int64()))
	channel <- prometheus.MustNewConstMetric(collector.WithdrawnValidators, prometheus.GaugeValue, float64(state.StaderNetworkDetails.WithdrawnValidators.Int64()))
	channel <- prometheus.MustNewConstMetric(collector.InvalidSignatureValidators, prometheus.GaugeValue, float64(state.StaderNetworkDetails.InvalidSignatureValidators.Int64()))
	channel <- prometheus.MustNewConstMetric(collector.IntializedValidators, prometheus.GaugeValue, float64(state.StaderNetworkDetails.InitializedValidators.Int64()))
	channel <- prometheus.MustNewConstMetric(collector.FrontRunValidators, prometheus.GaugeValue, float64(state.StaderNetworkDetails.FrontRunValidators.Int64()))
	channel <- prometheus.MustNewConstMetric(collector.FundsSettledValidators, prometheus.GaugeValue, float64(state.StaderNetworkDetails.FundsSettledValidators.Int64()))
	channel <- prometheus.MustNewConstMetric(collector.UnclaimedClRewards, prometheus.GaugeValue, state.StaderNetworkDetails.UnclaimedClRewards)
	channel <- prometheus.MustNewConstMetric(collector.CumulativePenalty, prometheus.GaugeValue, state.StaderNetworkDetails.CumulativePenalty)
	channel <- prometheus.MustNewConstMetric(collector.UnclaimedNonSocializingPoolElRewards, prometheus.GaugeValue, state.StaderNetworkDetails.UnclaimedNonSocializingPoolElRewards)
	channel <- prometheus.MustNewConstMetric(collector.UnclaimedSocializingPoolELRewards, prometheus.GaugeValue, state.StaderNetworkDetails.UnclaimedSocializingPoolElRewards)
	channel <- prometheus.MustNewConstMetric(collector.UnclaimedSocializingPoolSdRewards, prometheus.GaugeValue, state.StaderNetworkDetails.UnclaimedSocializingPoolSDRewards)
	channel <- prometheus.MustNewConstMetric(collector.ClaimedSocializingPoolSdRewards, prometheus.GaugeValue, state.StaderNetworkDetails.ClaimedSocializingPoolSdRewards)
	channel <- prometheus.MustNewConstMetric(collector.ClaimedSocializingPoolElRewards, prometheus.GaugeValue, state.StaderNetworkDetails.ClaimedSocializingPoolElRewards)
	channel <- prometheus.MustNewConstMetric(collector.TotalSdCollateral, prometheus.GaugeValue, state.StaderNetworkDetails.OperatorStakedSd)
	channel <- prometheus.MustNewConstMetric(collector.TotalSdCollateralInEth, prometheus.GaugeValue, state.StaderNetworkDetails.OperatorStakedSdInEth)
	channel <- prometheus.MustNewConstMetric(collector.TotalEthColateral, prometheus.GaugeValue, state.StaderNetworkDetails.OperatorEthCollateral)
	channel <- prometheus.MustNewConstMetric(collector.TotalSDUtilized, prometheus.GaugeValue, state.StaderNetworkDetails.OperatorSDUtilized)
	channel <- prometheus.MustNewConstMetric(collector.TotalSDUtilizedInterest, prometheus.GaugeValue, state.StaderNetworkDetails.OperatorSDInterest)
	channel <- prometheus.MustNewConstMetric(collector.SdCollateralPct, prometheus.GaugeValue, state.StaderNetworkDetails.SdCollateralPct)
	channel <- prometheus.MustNewConstMetric(collector.LockedEth, prometheus.GaugeValue, state.StaderNetworkDetails.LockedEth)
	channel <- prometheus.MustNewConstMetric(collector.HealthFactor, prometheus.GaugeValue, state.StaderNetworkDetails.HealthFactor)
	channel <- prometheus.MustNewConstMetric(collector.TotalSDUtilizationPosition, prometheus.GaugeValue, state.StaderNetworkDetails.OperatorSDUtilizationPosition)
	channel <- prometheus.MustNewConstMetric(collector.TotalSDSelfBond, prometheus.GaugeValue, state.StaderNetworkDetails.OperatorSDSelfBond)
	channel <- prometheus.MustNewConstMetric(collector.LiquidationStatus, prometheus.GaugeValue, state.StaderNetworkDetails.LiquidationStatus)
	channel <- prometheus.MustNewConstMetric(collector.ClaimVaultBalance, prometheus.GaugeValue, state.StaderNetworkDetails.ClaimVaultBalance)
	channel <- prometheus.MustNewConstMetric(collector.SDSelfBond, prometheus.GaugeValue, state.StaderNetworkDetails.OperatorSDSelfBond)
}

// Log error messages
func (collector *OperatorCollector) logError(err error) {
	fmt.Printf("[%s] %s\n", collector.logPrefix, err.Error())
}
