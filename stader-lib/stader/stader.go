package stader

import (
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/stader-labs/stader-node/stader-lib/contracts"
	"strings"
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
