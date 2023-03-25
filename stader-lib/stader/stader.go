/*
This work is licensed and released under GNU GPL v3 or any other later versions. 
The full text of the license is below/ found at <http://www.gnu.org/licenses/>

(c) 2023 Rocket Pool Pty Ltd. Modified under GNU GPL v3.

This program is free software: you can redistribute it and/or modify
it under the terms of the GNU General Public License as published by
the Free Software Foundation, either version 3 of the License, or
(at your option) any later version.

This program is distributed in the hope that it will be useful,
but WITHOUT ANY WARRANTY; without even the implied warranty of
MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
GNU General Public License for more details.

You should have received a copy of the GNU General Public License
along with this program.  If not, see <http://www.gnu.org/licenses/>.
*/
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
