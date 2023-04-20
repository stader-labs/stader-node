package collector

const namespace = "stader"

// Validator rewards & performance => stader_validator_rewards_performance + key
const RewardSub = "rewards"

const ActiveValidators = "active_validators"   //GetAllActiveValidators _PermissionlessNodeRegistry
const QueuedValidators = "queued_validators"   //QueuedValidators _PermissionlessNodeRegistry
const SlashedValidators = "slashed_validators" //GetValidatorStatus
const TotalETHBonded = "total_eth_bonded"
const TotalSDBonded = "total_sd_bonded"
const SdCollateral = "sd_collateral"
const BeaconchainReward = "beaconchain_reward"
const ElReward = "el_reward"
const SDReward = "sd_reward"
const ETHAPR = "eth_apr"
const SDAPR = "sd_apr"
const CumulativePenalty = "cumulative_penalty"
const ClaimedBeaconchainRewards = "claimed_beaconchain_rewards"
const ClaimedELRewards = "claimed_el_rewards"
const ClaimedSDrewards = "claimed_sd_rewards"
const UnclaimedELRewards = "unclaimed_el_rewards"
const UnclaimedSDRewards = "unclaimed_sd_rewards"
const NextSDOrELAndSDRewardsCheckpoint = "next_sd_or_el_and_sd_rewards_checkpoint"
const TotalAttestations = "total_attestations"
const AttestationPercent = "attestation_percent"
const BlocksProduced = "blocks_produced"
const BlocksProducedPercent = "blocks_produced_percent"
const AttestationInclusionEffectiveness = "attestation_inclusion_effectiveness"
const UptimePercent = "uptime_percent"

// Node Health => stader_node_health+ key
const NodeSub = "node_health"
const CPUUsage = "cpu_usage"
const CPUUsageTimeSeries = "cpu_usage_time_series"
const RAMUsage = "ram_usage"
const RAMUsageTimeSeries = "ram_usage_time_series"
const DiskSpaceUsed = "disk_space_used"
const SSDLatency = "ssd_latency"
const TotalIO = "total_io"
const IOWaiTTime = "io_wait_time"
const NetworkUsage = "network_usage"
const NetworkLatency = "network_latency"
const ECPeers = "ec_peers"
const NBCPeers = "nbc_peers"
