# enter in stader-lib/contracts and throw and error if it fails
echo "Updating ABIs..."

cd ./stader-lib/contracts;

abigen --abi ./../../abis/VaultFactory.abi.json --pkg contracts --type VaultFactory --out vault-factory.go;
abigen --abi ./../../abis/ValidatorWithdrawVault.abi.json --pkg contracts --type ValidatorWithdrawVault --out validator-withdraw-vault.go;
abigen --abi ./../../abis/StakePoolManager.abi.json --pkg contracts --type StakePoolManager --out stake-pool-manager.go;
abigen --abi ./../../abis/StaderConfig.abi.json --pkg contracts --type StaderConfig --out stader-config.go;
abigen --abi ./../../abis/SocializingPool.abi.json --pkg contracts --type SocializingPool --out socializing-pool.go;
abigen --abi ./../../abis/PoolUtils.abi.json --pkg contracts --type PoolUtils --out pool-utils.go;
abigen --abi ./../../abis/PermissionlessPool.abi.json --pkg contracts --type PermissionlessPool --out permissionless-pool.go;
abigen --abi ./../../abis/PermissionlessNodeRegistry.abi.json --pkg contracts --type PermissionlessNodeRegistry --out permissionless-node-registry.go;
abigen --abi ./../../abis/Penalty.abi.json --pkg contracts --type PenaltyTracker --out penalty.go;
abigen --abi ./../../abis/NodeElRewardVault.abi.json --pkg contracts --type NodeElRewardVault --out node-el-reward-vault.go;
abigen --abi ./../../abis/OperatorRewardsCollector.abi.json --pkg contracts --type OperatorRewardsCollector --out operator-rewards-collector.go;
abigen --abi ./../../abis/SdCollateral.abi.json --pkg contracts --type SdCollateral --out sd-collateral.go;

cd ../..;

echo "Done updating ABIs."