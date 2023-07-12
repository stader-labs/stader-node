package collector

const namespace = "stader"

// Validator rewards & performance => stader_validator_rewards_performance + key
const OperatorSub = "operator"

const ActiveValidators = "active_validators"                         //GetAllActiveValidators _PermissionlessNodeRegistry
const BeaconChainQueuedValidators = "beacon_chain_queued_validators" //BeaconChainQueuedValidators _PermissionlessNodeRegistry
const StaderQueuedValidators = "stader_queued_validators"
const SlashedValidators = "slashed_validators"                    //GetValidatorStatus
const ExitingValidators = "exiting_validators"                    //GetValidatorStatus
const WithdrawnValidators = "withdrawn_validators"                //GetValidatorStatus
const InitializedValidators = "initialized_validators"            //GetValidatorStatus
const InvalidSignatureValidators = "invalid_signature_validators" //GetValidatorStatus
const FrontRunValidators = "front_run_validators"                 //GetValidatorStatus
const FundsSettledValidators = "funds_settled_validators"         //GetValidatorStatus

const TotalETHBonded = "total_eth_bonded"
const TotalSDBonded = "total_sd_bonded"
const SdCollateral = "sd_collateral"
const SdCollateralInEth = "sd_collateral_in_eth"
const EthCollateral = "eth_collateral"
const CumulativePenalty = "cumulative_penalty"
const ClaimedSocializingPoolELRewards = "claimed_socializing_pool_el_rewards"
const ClaimedSocializingPoolSDrewards = "claimed_socializing_pool_sd_rewards"
const UnclaimedNonSocializingPoolELRewards = "unclaimed_non_socializing_pool_el_rewards"
const UnclaimedSocializingPoolELRewards = "unclaimed_socializing_pool_el_rewards"
const UnclaimedSocializingPoolSdRewards = "unclaimed_socializing_pool_sd_rewards"
const UnclaimedCLRewards = "unclaimed_cl_rewards"
const NextRewardCycleTime = "next_reward_cycle_time"

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
