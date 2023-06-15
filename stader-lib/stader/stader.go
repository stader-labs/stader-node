package stader

import (
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/stader-labs/stader-node/stader-lib/contracts"
)

type Erc20TokenContractManager struct {
	Client             ExecutionClient
	Erc20Token         *contracts.Erc20
	Erc20TokenContract *Contract
}

func NewErc20TokenContract(client ExecutionClient, erc20TokenAddress common.Address) (*Erc20TokenContractManager, error) {
	erc20TokenFactory, err := contracts.NewErc20(erc20TokenAddress, client)
	if err != nil {
		return nil, err
	}

	erc20Abi, err := abi.JSON(strings.NewReader(contracts.Erc20MetaData.ABI))
	if err != nil {
		return nil, err
	}
	erc20Contract := &Contract{
		Contract: bind.NewBoundContract(erc20TokenAddress, erc20Abi, client, client, client),
		Address:  &erc20TokenAddress,
		ABI:      &erc20Abi,
		Client:   client,
	}

	return &Erc20TokenContractManager{
		Client:             client,
		Erc20Token:         erc20TokenFactory,
		Erc20TokenContract: erc20Contract,
	}, nil

}

type EthxContractManager struct {
	Client       ExecutionClient
	EthX         *contracts.ETHX
	EthXContract *Contract
}

func NewEthxContractManager(client ExecutionClient, ethxAddress common.Address) (*EthxContractManager, error) {
	ethxFactory, err := contracts.NewETHX(ethxAddress, client)
	if err != nil {
		return nil, err
	}

	ethxAbi, err := abi.JSON(strings.NewReader(contracts.ETHXABI))
	if err != nil {
		return nil, err
	}
	ethXContract := &Contract{
		Contract: bind.NewBoundContract(ethxAddress, ethxAbi, client, client, client),
		Address:  &ethxAddress,
		ABI:      &ethxAbi,
		Client:   client,
	}

	return &EthxContractManager{
		Client:       client,
		EthX:         ethxFactory,
		EthXContract: ethXContract,
	}, nil

}

type SdCollateralContractManager struct {
	Client               ExecutionClient
	SdCollateral         *contracts.SdCollateral
	SdCollateralContract *Contract
}

func NewSdCollateralContract(client ExecutionClient, sdCollateralAddress common.Address) (*SdCollateralContractManager, error) {
	sdCollateralFactory, err := contracts.NewSdCollateral(sdCollateralAddress, client)
	if err != nil {
		return nil, err
	}

	sdCollateralAbi, err := abi.JSON(strings.NewReader(contracts.SdCollateralMetaData.ABI))
	if err != nil {
		return nil, err
	}
	sdCollateralContract := &Contract{
		Contract: bind.NewBoundContract(sdCollateralAddress, sdCollateralAbi, client, client, client),
		Address:  &sdCollateralAddress,
		ABI:      &sdCollateralAbi,
		Client:   client,
	}

	return &SdCollateralContractManager{
		Client:               client,
		SdCollateral:         sdCollateralFactory,
		SdCollateralContract: sdCollateralContract,
	}, nil

}

type PermissionlessNodeRegistryContractManager struct {
	Client                             ExecutionClient
	PermissionlessNodeRegistry         *contracts.PermissionlessNodeRegistry
	PermissionlessNodeRegistryContract *Contract
}

func NewPermissionlessNodeRegistry(client ExecutionClient, permissionlessNodeRegistryAddress common.Address) (*PermissionlessNodeRegistryContractManager, error) {
	permissionlessNodeRegistryFactory, err := contracts.NewPermissionlessNodeRegistry(permissionlessNodeRegistryAddress, client)
	if err != nil {
		return nil, err
	}

	permissionlessNodeRegistyAbi, err := abi.JSON(strings.NewReader(contracts.PermissionlessNodeRegistryMetaData.ABI))
	if err != nil {
		return nil, err
	}
	permissionlessNodeRegistryContract := &Contract{
		Contract: bind.NewBoundContract(permissionlessNodeRegistryAddress, permissionlessNodeRegistyAbi, client, client, client),
		Address:  &permissionlessNodeRegistryAddress,
		ABI:      &permissionlessNodeRegistyAbi,
		Client:   client,
	}

	return &PermissionlessNodeRegistryContractManager{
		Client:                             client,
		PermissionlessNodeRegistry:         permissionlessNodeRegistryFactory,
		PermissionlessNodeRegistryContract: permissionlessNodeRegistryContract,
	}, nil

}

type VaultFactoryContractManager struct {
	Client               ExecutionClient
	VaultFactory         *contracts.VaultFactory
	VaultFactoryContract *Contract
}

func NewVaultFactory(client ExecutionClient, vaultFactoryAddress common.Address) (*VaultFactoryContractManager, error) {
	vaultFactory, err := contracts.NewVaultFactory(vaultFactoryAddress, client)
	if err != nil {
		return nil, err
	}

	vaultFactoryContractAbi, err := abi.JSON(strings.NewReader(contracts.VaultFactoryMetaData.ABI))
	if err != nil {
		return nil, err
	}
	vaultFactoryContract := &Contract{
		Contract: bind.NewBoundContract(vaultFactoryAddress, vaultFactoryContractAbi, client, client, client),
		Address:  &vaultFactoryAddress,
		ABI:      &vaultFactoryContractAbi,
		Client:   client,
	}

	return &VaultFactoryContractManager{
		Client:               client,
		VaultFactory:         vaultFactory,
		VaultFactoryContract: vaultFactoryContract,
	}, nil

}

type PermissionlessPoolContractManager struct {
	Client                     ExecutionClient
	PermissionlessPool         *contracts.PermissionlessPool
	PermissionlessPoolContract *Contract
}

func NewPermissionlessPoolFactory(client ExecutionClient, permissionPoolAddress common.Address) (*PermissionlessPoolContractManager, error) {
	permissionlessPool, err := contracts.NewPermissionlessPool(permissionPoolAddress, client)
	if err != nil {
		return nil, err
	}

	permissionlessPoolContractAbi, err := abi.JSON(strings.NewReader(contracts.PermissionlessPoolMetaData.ABI))
	if err != nil {
		return nil, err
	}
	permissionlessPoolContract := &Contract{
		Contract: bind.NewBoundContract(permissionPoolAddress, permissionlessPoolContractAbi, client, client, client),
		Address:  &permissionPoolAddress,
		ABI:      &permissionlessPoolContractAbi,
		Client:   client,
	}

	return &PermissionlessPoolContractManager{
		Client:                     client,
		PermissionlessPool:         permissionlessPool,
		PermissionlessPoolContract: permissionlessPoolContract,
	}, nil

}

type NodeElRewardVaultContractManager struct {
	Client                    ExecutionClient
	NodeElRewardVault         *contracts.NodeElRewardVault
	NodeElRewardVaultContract *Contract
}

func NewNodeElRewardVaultFactory(client ExecutionClient, nodeElRewardVaultAddress common.Address) (*NodeElRewardVaultContractManager, error) {
	nodeElRewardVault, err := contracts.NewNodeElRewardVault(nodeElRewardVaultAddress, client)
	if err != nil {
		return nil, err
	}

	nodeElRewardVaultContractAbi, err := abi.JSON(strings.NewReader(contracts.NodeElRewardVaultMetaData.ABI))
	if err != nil {
		return nil, err
	}
	nodeElRewardContract := &Contract{
		Contract: bind.NewBoundContract(nodeElRewardVaultAddress, nodeElRewardVaultContractAbi, client, client, client),
		Address:  &nodeElRewardVaultAddress,
		ABI:      &nodeElRewardVaultContractAbi,
		Client:   client,
	}

	return &NodeElRewardVaultContractManager{
		Client:                    client,
		NodeElRewardVault:         nodeElRewardVault,
		NodeElRewardVaultContract: nodeElRewardContract,
	}, nil

}

type ValidatorWithdrawVaultContractManager struct {
	Client                         ExecutionClient
	ValidatorWithdrawVault         *contracts.ValidatorWithdrawVault
	ValidatorWithdrawVaultContract *Contract
}

func NewValidatorWithdrawVaultFactory(client ExecutionClient, validatorWithdrawVaultAddress common.Address) (*ValidatorWithdrawVaultContractManager, error) {
	validatorWithdrawVault, err := contracts.NewValidatorWithdrawVault(validatorWithdrawVaultAddress, client)
	if err != nil {
		return nil, err
	}

	validatorWithdrawVaultContractAbi, err := abi.JSON(strings.NewReader(contracts.ValidatorWithdrawVaultMetaData.ABI))
	if err != nil {
		return nil, err
	}
	validatorWithdrawContract := &Contract{
		Contract: bind.NewBoundContract(validatorWithdrawVaultAddress, validatorWithdrawVaultContractAbi, client, client, client),
		Address:  &validatorWithdrawVaultAddress,
		ABI:      &validatorWithdrawVaultContractAbi,
		Client:   client,
	}

	return &ValidatorWithdrawVaultContractManager{
		Client:                         client,
		ValidatorWithdrawVault:         validatorWithdrawVault,
		ValidatorWithdrawVaultContract: validatorWithdrawContract,
	}, nil

}

type StaderConfigContractManager struct {
	Client               ExecutionClient
	StaderConfig         *contracts.StaderConfig
	StaderConfigContract *Contract
}

func NewStaderConfig(client ExecutionClient, staderConfigAddress common.Address) (*StaderConfigContractManager, error) {
	staderConfig, err := contracts.NewStaderConfig(staderConfigAddress, client)
	if err != nil {
		return nil, err
	}

	staderConfigContractAbi, err := abi.JSON(strings.NewReader(contracts.StaderConfigMetaData.ABI))
	if err != nil {
		return nil, err
	}
	staderConfigContract := &Contract{
		Contract: bind.NewBoundContract(staderConfigAddress, staderConfigContractAbi, client, client, client),
		Address:  &staderConfigAddress,
		ABI:      &staderConfigContractAbi,
		Client:   client,
	}

	return &StaderConfigContractManager{
		Client:               client,
		StaderConfig:         staderConfig,
		StaderConfigContract: staderConfigContract,
	}, nil

}

type SocializingPoolContractManager struct {
	Client                  ExecutionClient
	SocializingPool         *contracts.SocializingPool
	SocializingPoolContract *Contract
}

func NewSocializingPool(client ExecutionClient, socializingPoolAddress common.Address) (*SocializingPoolContractManager, error) {
	socializingPool, err := contracts.NewSocializingPool(socializingPoolAddress, client)
	if err != nil {
		return nil, err
	}

	socializingPoolContractAbi, err := abi.JSON(strings.NewReader(contracts.SocializingPoolMetaData.ABI))
	if err != nil {
		return nil, err
	}
	socializingPoolContract := &Contract{
		Contract: bind.NewBoundContract(socializingPoolAddress, socializingPoolContractAbi, client, client, client),
		Address:  &socializingPoolAddress,
		ABI:      &socializingPoolContractAbi,
		Client:   client,
	}

	return &SocializingPoolContractManager{
		Client:                  client,
		SocializingPool:         socializingPool,
		SocializingPoolContract: socializingPoolContract,
	}, nil

}

// Write bindings similar to the above for PoolUtils

type PoolUtilsContractManager struct {
	Client            ExecutionClient
	PoolUtils         *contracts.PoolUtils
	PoolUtilsContract *Contract
}

func NewPoolUtils(client ExecutionClient, poolUtilsAddress common.Address) (*PoolUtilsContractManager, error) {
	poolUtils, err := contracts.NewPoolUtils(poolUtilsAddress, client)
	if err != nil {
		return nil, err
	}

	poolUtilsContractAbi, err := abi.JSON(strings.NewReader(contracts.PoolUtilsMetaData.ABI))
	if err != nil {
		return nil, err
	}
	poolUtilsContract := &Contract{
		Contract: bind.NewBoundContract(poolUtilsAddress, poolUtilsContractAbi, client, client, client),
		Address:  &poolUtilsAddress,
		ABI:      &poolUtilsContractAbi,
		Client:   client,
	}

	return &PoolUtilsContractManager{
		Client:            client,
		PoolUtils:         poolUtils,
		PoolUtilsContract: poolUtilsContract,
	}, nil

}

// implement similar bindings as above for Penalty

type PenaltyTrackerContractManager struct {
	Client          ExecutionClient
	Penalty         *contracts.PenaltyTracker
	PenaltyContract *Contract
}

func NewPenaltyTracker(client ExecutionClient, penaltyAddress common.Address) (*PenaltyTrackerContractManager, error) {
	penalty, err := contracts.NewPenaltyTracker(penaltyAddress, client)
	if err != nil {
		return nil, err
	}

	penaltyContractAbi, err := abi.JSON(strings.NewReader(contracts.PenaltyTrackerMetaData.ABI))
	if err != nil {
		return nil, err
	}
	penaltyContract := &Contract{
		Contract: bind.NewBoundContract(penaltyAddress, penaltyContractAbi, client, client, client),
		Address:  &penaltyAddress,
		ABI:      &penaltyContractAbi,
		Client:   client,
	}

	return &PenaltyTrackerContractManager{
		Client:          client,
		Penalty:         penalty,
		PenaltyContract: penaltyContract,
	}, nil

}

// implement simliar bindings as above for stake pool manager

type StakePoolManagerContractManager struct {
	Client                   ExecutionClient
	StakePoolManager         *contracts.StakePoolManager
	StakePoolManagerContract *Contract
}

func NewStakePoolManager(client ExecutionClient, stakePoolManagerAddress common.Address) (*StakePoolManagerContractManager, error) {
	stakePoolManager, err := contracts.NewStakePoolManager(stakePoolManagerAddress, client)
	if err != nil {
		return nil, err
	}

	stakePoolManagerContractAbi, err := abi.JSON(strings.NewReader(contracts.StakePoolManagerMetaData.ABI))
	if err != nil {
		return nil, err
	}
	stakePoolManagerContract := &Contract{
		Contract: bind.NewBoundContract(stakePoolManagerAddress, stakePoolManagerContractAbi, client, client, client),
		Address:  &stakePoolManagerAddress,
		ABI:      &stakePoolManagerContractAbi,
		Client:   client,
	}

	return &StakePoolManagerContractManager{
		Client:                   client,
		StakePoolManager:         stakePoolManager,
		StakePoolManagerContract: stakePoolManagerContract,
	}, nil

}

// Write above bindings for VaultProxy
type VaultProxyContractManager struct {
	Client             ExecutionClient
	VaultProxy         *contracts.VaultProxy
	VaultProxyContract *Contract
}

func NewVaultProxy(client ExecutionClient, vaultProxyAddress common.Address) (*VaultProxyContractManager, error) {
	vaultProxy, err := contracts.NewVaultProxy(vaultProxyAddress, client)
	if err != nil {
		return nil, err
	}

	vaultProxyContractAbi, err := abi.JSON(strings.NewReader(contracts.VaultProxyMetaData.ABI))
	if err != nil {
		return nil, err
	}
	vaultProxyContract := &Contract{
		Contract: bind.NewBoundContract(vaultProxyAddress, vaultProxyContractAbi, client, client, client),
		Address:  &vaultProxyAddress,
		ABI:      &vaultProxyContractAbi,
		Client:   client,
	}

	return &VaultProxyContractManager{
		Client:             client,
		VaultProxy:         vaultProxy,
		VaultProxyContract: vaultProxyContract,
	}, nil

}

// Add bindings similar to the above for OperatorRewardsCollector

type OperatorRewardsCollectorContractManager struct {
	Client                           ExecutionClient
	OperatorRewardsCollector         *contracts.OperatorRewardsCollector
	OperatorRewardsCollectorContract *Contract
}

func NewOperatorRewardsCollector(client ExecutionClient, operatorRewardsCollectorAddress common.Address) (*OperatorRewardsCollectorContractManager, error) {
	operatorRewardsCollector, err := contracts.NewOperatorRewardsCollector(operatorRewardsCollectorAddress, client)
	if err != nil {
		return nil, err
	}

	operatorRewardsCollectorContractAbi, err := abi.JSON(strings.NewReader(contracts.OperatorRewardsCollectorMetaData.ABI))
	if err != nil {
		return nil, err
	}
	operatorRewardsCollectorContract := &Contract{
		Contract: bind.NewBoundContract(operatorRewardsCollectorAddress, operatorRewardsCollectorContractAbi, client, client, client),
		Address:  &operatorRewardsCollectorAddress,
		ABI:      &operatorRewardsCollectorContractAbi,
		Client:   client,
	}

	return &OperatorRewardsCollectorContractManager{
		Client:                           client,
		OperatorRewardsCollector:         operatorRewardsCollector,
		OperatorRewardsCollectorContract: operatorRewardsCollectorContract,
	}, nil

}
