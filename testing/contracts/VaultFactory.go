// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package contracts

import (
	"errors"
	"math/big"
	"strings"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/event"
)

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = errors.New
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
	_ = abi.ConvertType
)

// VaultFactoryMetaData contains all meta data concerning the VaultFactory contract.
var VaultFactoryMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_admin\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_staderConfig\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"ZeroAddress\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"version\",\"type\":\"uint8\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"nodeDistributor\",\"type\":\"address\"}],\"name\":\"NodeELRewardVaultCreated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"previousAdminRole\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"newAdminRole\",\"type\":\"bytes32\"}],\"name\":\"RoleAdminChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleGranted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleRevoked\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"staderConfig\",\"type\":\"address\"}],\"name\":\"UpdatedStaderConfig\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"vaultProxyImplementation\",\"type\":\"address\"}],\"name\":\"UpdatedVaultProxyImplementation\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"withdrawVault\",\"type\":\"address\"}],\"name\":\"WithdrawVaultCreated\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"DEFAULT_ADMIN_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"NODE_REGISTRY_CONTRACT\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"_poolId\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"_operatorId\",\"type\":\"uint256\"}],\"name\":\"computeNodeELRewardVaultAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"_poolId\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"_operatorId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_validatorCount\",\"type\":\"uint256\"}],\"name\":\"computeWithdrawVaultAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"_poolId\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"_operatorId\",\"type\":\"uint256\"}],\"name\":\"deployNodeELRewardVault\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"_poolId\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"_operatorId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_validatorCount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_validatorId\",\"type\":\"uint256\"}],\"name\":\"deployWithdrawVault\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"}],\"name\":\"getRoleAdmin\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_withdrawVault\",\"type\":\"address\"}],\"name\":\"getValidatorWithdrawCredential\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"grantRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"hasRole\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"renounceRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"revokeRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"staderConfig\",\"outputs\":[{\"internalType\":\"contractIStaderConfig\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_staderConfig\",\"type\":\"address\"}],\"name\":\"updateStaderConfig\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_vaultProxyImpl\",\"type\":\"address\"}],\"name\":\"updateVaultProxyAddress\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"vaultProxyImplementation\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Bin: "0x608060405234801562000010575f80fd5b5060405162001b0838038062001b0883398101604081905262000033916200033d565b5f54610100900460ff16158080156200005257505f54600160ff909116105b806200006d5750303b1580156200006d57505f5460ff166001145b620000d65760405162461bcd60e51b815260206004820152602e60248201527f496e697469616c697a61626c653a20636f6e747261637420697320616c72656160448201526d191e481a5b9a5d1a585b1a5e995960921b60648201526084015b60405180910390fd5b5f805460ff191660011790558015620000f8575f805461ff0019166101001790555b6200010383620001d7565b6200010e82620001d7565b6200011862000202565b609780546001600160a01b0319166001600160a01b038416179055604051620001419062000313565b604051809103905ff0801580156200015b573d5f803e3d5ffd5b50609880546001600160a01b0319166001600160a01b0392909216919091179055620001885f8462000270565b8015620001ce575f805461ff0019169055604051600181527f7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb38474024989060200160405180910390a15b50505062000373565b6001600160a01b038116620001ff5760405163d92e233d60e01b815260040160405180910390fd5b50565b5f54610100900460ff166200026e5760405162461bcd60e51b815260206004820152602b60248201527f496e697469616c697a61626c653a20636f6e7472616374206973206e6f74206960448201526a6e697469616c697a696e6760a81b6064820152608401620000cd565b565b5f8281526065602090815260408083206001600160a01b038516845290915290205460ff166200030f575f8281526065602090815260408083206001600160a01b03851684529091529020805460ff19166001179055620002ce3390565b6001600160a01b0316816001600160a01b0316837f2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d60405160405180910390a45b5050565b610773806200139583390190565b80516001600160a01b038116811462000338575f80fd5b919050565b5f80604083850312156200034f575f80fd5b6200035a8362000321565b91506200036a6020840162000321565b90509250929050565b61101480620003815f395ff3fe608060405234801561000f575f80fd5b5060043610610106575f3560e01c806391d148541161009e578063ae4e4e451161006e578063ae4e4e451461022e578063ca934f601461028e578063d547741f146102a1578063ee5fca3b146102b4578063f8bc93a5146102c7575f80fd5b806391d14854146101ee5780639ee804cb14610201578063a217fddf14610214578063ab4f36d51461021b575f80fd5b8063490ffa35116100d9578063490ffa351461018a5780636a0b6881146101b557806374903b02146101c85780637f70ce0d146101db575f80fd5b806301ffc9a71461010a578063248a9ca3146101325780632f2ff15d1461016257806336568abe14610177575b5f80fd5b61011d610118366004610d24565b6102ee565b60405190151581526020015b60405180910390f35b610154610140366004610d4b565b5f9081526065602052604090206001015490565b604051908152602001610129565b610175610170366004610d7d565b610324565b005b610175610185366004610d7d565b61034d565b60975461019d906001600160a01b031681565b6040516001600160a01b039091168152602001610129565b61019d6101c3366004610db7565b6103d0565b61019d6101d6366004610ddf565b610541565b61019d6101e9366004610e0f565b6105df565b61011d6101fc366004610d7d565b61075a565b61017561020f366004610e45565b610784565b6101545f81565b610175610229366004610e45565b6107ed565b61028161023c366004610e45565b60408051600160f81b60208201525f6021820152606083811b6bffffffffffffffffffffffff1916602c83015291016040516020818303038152906040529050919050565b6040516101299190610eab565b61019d61029c366004610db7565b61084e565b6101756102af366004610d7d565b6108e4565b60985461019d906001600160a01b031681565b6101547f5f937b69bc198c032799e5628d33386025c498d279cec8ad3c304e5c00bb19f681565b5f6001600160e01b03198216637965db0b60e01b148061031e57506301ffc9a760e01b6001600160e01b03198316145b92915050565b5f8281526065602052604090206001015461033e81610908565b6103488383610915565b505050565b6001600160a01b03811633146103c25760405162461bcd60e51b815260206004820152602f60248201527f416363657373436f6e74726f6c3a2063616e206f6e6c792072656e6f756e636560448201526e103937b632b9903337b91039b2b63360891b60648201526084015b60405180910390fd5b6103cc828261099a565b5050565b5f7f5f937b69bc198c032799e5628d33386025c498d279cec8ad3c304e5c00bb19f66103fb81610908565b6040805160ff861660208201529081018490525f9060029060600160408051601f198184030181529082905261043091610ebd565b602060405180830381855afa15801561044b573d5f803e3d5ffd5b5050506040513d601f19601f8201168201806040525081019061046e9190610ed8565b6098549091505f90610489906001600160a01b031683610a00565b609754604051630483167960e21b81525f600482015260ff89166024820152604481018890526001600160a01b03918216606482015291925082169063120c59e4906084015f604051808303815f87803b1580156104e5575f80fd5b505af11580156104f7573d5f803e3d5ffd5b50506040516001600160a01b03841681527f534e63617d48745c45efd754f53f9292453257107e03d0140db96fc8e3bf7ec99250602001905060405180910390a195945050505050565b6040805160ff85166020820152908101839052606081018290525f90819060029060800160408051601f198184030181529082905261057f91610ebd565b602060405180830381855afa15801561059a573d5f803e3d5ffd5b5050506040513d601f19601f820116820180604052508101906105bd9190610ed8565b6098549091506105d6906001600160a01b031682610a9a565b95945050505050565b5f7f5f937b69bc198c032799e5628d33386025c498d279cec8ad3c304e5c00bb19f661060a81610908565b6040805160ff88166020820152908101869052606081018590525f9060029060800160408051601f198184030181529082905261064691610ebd565b602060405180830381855afa158015610661573d5f803e3d5ffd5b5050506040513d601f19601f820116820180604052508101906106849190610ed8565b6098549091505f9061069f906001600160a01b031683610a00565b609754604051630483167960e21b81526001600482015260ff8b166024820152604481018890526001600160a01b03918216606482015291925082169063120c59e4906084015f604051808303815f87803b1580156106fc575f80fd5b505af115801561070e573d5f803e3d5ffd5b50506040516001600160a01b03841681527f3961159e01aea37eec106905b883c13f18c7c19abfde04fef6835c94d2af51b99250602001905060405180910390a1979650505050505050565b5f9182526065602090815260408084206001600160a01b0393909316845291905290205460ff1690565b5f61078e81610908565b61079782610afc565b609780546001600160a01b0319166001600160a01b0384169081179091556040519081527fdb2219043d7b197cb235f1af0cf6d782d77dee3de19e3f4fb6d39aae633b4485906020015b60405180910390a15050565b5f6107f781610908565b61080082610afc565b609880546001600160a01b0319166001600160a01b0384169081179091556040519081527f77358c5e9bc05f53688720682a283a5de55b771a614836c83cb1428deb2ece2d906020016107e1565b6040805160ff841660208201529081018290525f90819060029060600160408051601f198184030181529082905261088591610ebd565b602060405180830381855afa1580156108a0573d5f803e3d5ffd5b5050506040513d601f19601f820116820180604052508101906108c39190610ed8565b6098549091506108dc906001600160a01b031682610a9a565b949350505050565b5f828152606560205260409020600101546108fe81610908565b610348838361099a565b6109128133610b23565b50565b61091f828261075a565b6103cc575f8281526065602090815260408083206001600160a01b03851684529091529020805460ff191660011790556109563390565b6001600160a01b0316816001600160a01b0316837f2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d60405160405180910390a45050565b6109a4828261075a565b156103cc575f8281526065602090815260408083206001600160a01b0385168085529252808320805460ff1916905551339285917ff6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b9190a45050565b5f763d602d80600a3d3981f3363d3d373d3d3d363d730000008360601b60e81c175f526e5af43d82803e903d91602b57fd5bf38360781b1760205281603760095ff590506001600160a01b03811661031e5760405162461bcd60e51b815260206004820152601760248201527f455243313136373a2063726561746532206661696c656400000000000000000060448201526064016103b9565b6040513060388201526f5af43d82803e903d91602b57fd5bf3ff602482015260148101839052733d602d80600a3d3981f3363d3d373d3d3d363d738152605881018290526037600c820120607882015260556043909101205f905b9392505050565b6001600160a01b0381166109125760405163d92e233d60e01b815260040160405180910390fd5b610b2d828261075a565b6103cc57610b3a81610b7c565b610b45836020610b8e565b604051602001610b56929190610eef565b60408051601f198184030181529082905262461bcd60e51b82526103b991600401610eab565b606061031e6001600160a01b03831660145b60605f610b9c836002610f77565b610ba7906002610f8e565b67ffffffffffffffff811115610bbf57610bbf610fa1565b6040519080825280601f01601f191660200182016040528015610be9576020820181803683370190505b509050600360fc1b815f81518110610c0357610c03610fb5565b60200101906001600160f81b03191690815f1a905350600f60fb1b81600181518110610c3157610c31610fb5565b60200101906001600160f81b03191690815f1a9053505f610c53846002610f77565b610c5e906001610f8e565b90505b6001811115610cd5576f181899199a1a9b1b9c1cb0b131b232b360811b85600f1660108110610c9257610c92610fb5565b1a60f81b828281518110610ca857610ca8610fb5565b60200101906001600160f81b03191690815f1a90535060049490941c93610cce81610fc9565b9050610c61565b508315610af55760405162461bcd60e51b815260206004820181905260248201527f537472696e67733a20686578206c656e67746820696e73756666696369656e7460448201526064016103b9565b5f60208284031215610d34575f80fd5b81356001600160e01b031981168114610af5575f80fd5b5f60208284031215610d5b575f80fd5b5035919050565b80356001600160a01b0381168114610d78575f80fd5b919050565b5f8060408385031215610d8e575f80fd5b82359150610d9e60208401610d62565b90509250929050565b803560ff81168114610d78575f80fd5b5f8060408385031215610dc8575f80fd5b610dd183610da7565b946020939093013593505050565b5f805f60608486031215610df1575f80fd5b610dfa84610da7565b95602085013595506040909401359392505050565b5f805f8060808587031215610e22575f80fd5b610e2b85610da7565b966020860135965060408601359560600135945092505050565b5f60208284031215610e55575f80fd5b610af582610d62565b5f5b83811015610e78578181015183820152602001610e60565b50505f910152565b5f8151808452610e97816020860160208601610e5e565b601f01601f19169290920160200192915050565b602081525f610af56020830184610e80565b5f8251610ece818460208701610e5e565b9190910192915050565b5f60208284031215610ee8575f80fd5b5051919050565b7f416363657373436f6e74726f6c3a206163636f756e742000000000000000000081525f8351610f26816017850160208801610e5e565b7001034b99036b4b9b9b4b733903937b6329607d1b6017918401918201528351610f57816028840160208801610e5e565b01602801949350505050565b634e487b7160e01b5f52601160045260245ffd5b808202811582820484141761031e5761031e610f63565b8082018082111561031e5761031e610f63565b634e487b7160e01b5f52604160045260245ffd5b634e487b7160e01b5f52603260045260245ffd5b5f81610fd757610fd7610f63565b505f19019056fea264697066735822122013081fda5548d27aa7ad80c727a3e4863c1952dd02f6bec4c7424da0b7d865c964736f6c63430008140033608060405234801561000f575f80fd5b506107568061001d5f395ff3fe608060405260043610610090575f3560e01c80637ef4947d116100585780637ef4947d146103045780638da5cb5b1461031c5780639ee804cb1461033b578063af640d0f1461035a578063bc5920ba1461037d57610090565b8063120c59e41461022b578063392e53cd1461024c5780633e0dc34e1461027f578063490ffa35146102b05780637cc0bb90146102e7575b5f3660605f8060019054906101000a900460ff166101215760035f9054906101000a90046001600160a01b03166001600160a01b031663e8fe18736040518163ffffffff1660e01b8152600401602060405180830381865afa1580156100f8573d5f803e3d5ffd5b505050506040513d601f19601f8201168201806040525081019061011c9190610630565b610195565b60035f9054906101000a90046001600160a01b03166001600160a01b0316636d28ad1c6040518163ffffffff1660e01b8152600401602060405180830381865afa158015610171573d5f803e3d5ffd5b505050506040513d601f19601f820116820180604052508101906101959190610630565b90505f80826001600160a01b031686866040516101b3929190610652565b5f60405180830381855af49150503d805f81146101eb576040519150601f19603f3d011682016040523d82523d5f602084013e6101f0565b606091505b50915091508161021d578060405162461bcd60e51b81526004016102149190610661565b60405180910390fd5b805195506020019350505050f35b348015610236575f80fd5b5061024a6102453660046106ac565b610391565b005b348015610257575f80fd5b505f5461026a9062010000900460ff1681565b60405190151581526020015b60405180910390f35b34801561028a575f80fd5b505f5461029e906301000000900460ff1681565b60405160ff9091168152602001610276565b3480156102bb575f80fd5b506003546102cf906001600160a01b031681565b6040516001600160a01b039091168152602001610276565b3480156102f2575f80fd5b505f5461026a90610100900460ff1681565b34801561030f575f80fd5b505f5461026a9060ff1681565b348015610327575f80fd5b506002546102cf906001600160a01b031681565b348015610346575f80fd5b5061024a610355366004610705565b6104a2565b348015610365575f80fd5b5061036f60015481565b604051908152602001610276565b348015610388575f80fd5b5061024a61052a565b5f5462010000900460ff16156103b95760405162dc149f60e41b815260040160405180910390fd5b6103c2816105f2565b5f80546201000062ffff00199091166101008715150262ff00001916171763ff0000001916630100000060ff8616021790556001829055600380546001600160a01b0319166001600160a01b03831690811790915560408051636e9960c360e01b81529051636e9960c3916004808201926020929091908290030181865afa158015610450573d5f803e3d5ffd5b505050506040513d601f19601f820116820180604052508101906104749190610630565b600280546001600160a01b0319166001600160a01b0392909216918217905561049c906105f2565b50505050565b6002546001600160a01b031633146104cd57604051632e6c18c960e11b815260040160405180910390fd5b6104d6816105f2565b600380546001600160a01b0319166001600160a01b0383169081179091556040519081527fdb2219043d7b197cb235f1af0cf6d782d77dee3de19e3f4fb6d39aae633b44859060200160405180910390a150565b60035f9054906101000a90046001600160a01b03166001600160a01b0316636e9960c36040518163ffffffff1660e01b8152600401602060405180830381865afa15801561057a573d5f803e3d5ffd5b505050506040513d601f19601f8201168201806040525081019061059e9190610630565b600280546001600160a01b0319166001600160a01b039290921691821790556040519081527f957090e72c0a1b3ebf83c682eb8c1f88c2a18cd0578b91a819efb28859f0f3a39060200160405180910390a1565b6001600160a01b0381166106195760405163d92e233d60e01b815260040160405180910390fd5b50565b6001600160a01b0381168114610619575f80fd5b5f60208284031215610640575f80fd5b815161064b8161061c565b9392505050565b818382375f9101908152919050565b5f6020808352835180828501525f5b8181101561068c57858101830151858201604001528201610670565b505f604082860101526040601f19601f8301168501019250505092915050565b5f805f80608085870312156106bf575f80fd5b843580151581146106ce575f80fd5b9350602085013560ff811681146106e3575f80fd5b92506040850135915060608501356106fa8161061c565b939692955090935050565b5f60208284031215610715575f80fd5b813561064b8161061c56fea264697066735822122032304ea603f2de08d0b88b402f1d134232dce29b0fb9b0d9bd6b8fdb821fc23b64736f6c63430008140033",
}

// VaultFactoryABI is the input ABI used to generate the binding from.
// Deprecated: Use VaultFactoryMetaData.ABI instead.
var VaultFactoryABI = VaultFactoryMetaData.ABI

// VaultFactoryBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use VaultFactoryMetaData.Bin instead.
var VaultFactoryBin = VaultFactoryMetaData.Bin

// DeployVaultFactory deploys a new Ethereum contract, binding an instance of VaultFactory to it.
func DeployVaultFactory(auth *bind.TransactOpts, backend bind.ContractBackend, _admin common.Address, _staderConfig common.Address) (common.Address, *types.Transaction, *VaultFactory, error) {
	parsed, err := VaultFactoryMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(VaultFactoryBin), backend, _admin, _staderConfig)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &VaultFactory{VaultFactoryCaller: VaultFactoryCaller{contract: contract}, VaultFactoryTransactor: VaultFactoryTransactor{contract: contract}, VaultFactoryFilterer: VaultFactoryFilterer{contract: contract}}, nil
}

// VaultFactory is an auto generated Go binding around an Ethereum contract.
type VaultFactory struct {
	VaultFactoryCaller     // Read-only binding to the contract
	VaultFactoryTransactor // Write-only binding to the contract
	VaultFactoryFilterer   // Log filterer for contract events
}

// VaultFactoryCaller is an auto generated read-only Go binding around an Ethereum contract.
type VaultFactoryCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// VaultFactoryTransactor is an auto generated write-only Go binding around an Ethereum contract.
type VaultFactoryTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// VaultFactoryFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type VaultFactoryFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// VaultFactorySession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type VaultFactorySession struct {
	Contract     *VaultFactory     // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// VaultFactoryCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type VaultFactoryCallerSession struct {
	Contract *VaultFactoryCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts       // Call options to use throughout this session
}

// VaultFactoryTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type VaultFactoryTransactorSession struct {
	Contract     *VaultFactoryTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts       // Transaction auth options to use throughout this session
}

// VaultFactoryRaw is an auto generated low-level Go binding around an Ethereum contract.
type VaultFactoryRaw struct {
	Contract *VaultFactory // Generic contract binding to access the raw methods on
}

// VaultFactoryCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type VaultFactoryCallerRaw struct {
	Contract *VaultFactoryCaller // Generic read-only contract binding to access the raw methods on
}

// VaultFactoryTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type VaultFactoryTransactorRaw struct {
	Contract *VaultFactoryTransactor // Generic write-only contract binding to access the raw methods on
}

// NewVaultFactory creates a new instance of VaultFactory, bound to a specific deployed contract.
func NewVaultFactory(address common.Address, backend bind.ContractBackend) (*VaultFactory, error) {
	contract, err := bindVaultFactory(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &VaultFactory{VaultFactoryCaller: VaultFactoryCaller{contract: contract}, VaultFactoryTransactor: VaultFactoryTransactor{contract: contract}, VaultFactoryFilterer: VaultFactoryFilterer{contract: contract}}, nil
}

// NewVaultFactoryCaller creates a new read-only instance of VaultFactory, bound to a specific deployed contract.
func NewVaultFactoryCaller(address common.Address, caller bind.ContractCaller) (*VaultFactoryCaller, error) {
	contract, err := bindVaultFactory(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &VaultFactoryCaller{contract: contract}, nil
}

// NewVaultFactoryTransactor creates a new write-only instance of VaultFactory, bound to a specific deployed contract.
func NewVaultFactoryTransactor(address common.Address, transactor bind.ContractTransactor) (*VaultFactoryTransactor, error) {
	contract, err := bindVaultFactory(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &VaultFactoryTransactor{contract: contract}, nil
}

// NewVaultFactoryFilterer creates a new log filterer instance of VaultFactory, bound to a specific deployed contract.
func NewVaultFactoryFilterer(address common.Address, filterer bind.ContractFilterer) (*VaultFactoryFilterer, error) {
	contract, err := bindVaultFactory(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &VaultFactoryFilterer{contract: contract}, nil
}

// bindVaultFactory binds a generic wrapper to an already deployed contract.
func bindVaultFactory(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := VaultFactoryMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_VaultFactory *VaultFactoryRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _VaultFactory.Contract.VaultFactoryCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_VaultFactory *VaultFactoryRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _VaultFactory.Contract.VaultFactoryTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_VaultFactory *VaultFactoryRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _VaultFactory.Contract.VaultFactoryTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_VaultFactory *VaultFactoryCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _VaultFactory.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_VaultFactory *VaultFactoryTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _VaultFactory.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_VaultFactory *VaultFactoryTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _VaultFactory.Contract.contract.Transact(opts, method, params...)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_VaultFactory *VaultFactoryCaller) DEFAULTADMINROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _VaultFactory.contract.Call(opts, &out, "DEFAULT_ADMIN_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_VaultFactory *VaultFactorySession) DEFAULTADMINROLE() ([32]byte, error) {
	return _VaultFactory.Contract.DEFAULTADMINROLE(&_VaultFactory.CallOpts)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_VaultFactory *VaultFactoryCallerSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _VaultFactory.Contract.DEFAULTADMINROLE(&_VaultFactory.CallOpts)
}

// NODEREGISTRYCONTRACT is a free data retrieval call binding the contract method 0xf8bc93a5.
//
// Solidity: function NODE_REGISTRY_CONTRACT() view returns(bytes32)
func (_VaultFactory *VaultFactoryCaller) NODEREGISTRYCONTRACT(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _VaultFactory.contract.Call(opts, &out, "NODE_REGISTRY_CONTRACT")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// NODEREGISTRYCONTRACT is a free data retrieval call binding the contract method 0xf8bc93a5.
//
// Solidity: function NODE_REGISTRY_CONTRACT() view returns(bytes32)
func (_VaultFactory *VaultFactorySession) NODEREGISTRYCONTRACT() ([32]byte, error) {
	return _VaultFactory.Contract.NODEREGISTRYCONTRACT(&_VaultFactory.CallOpts)
}

// NODEREGISTRYCONTRACT is a free data retrieval call binding the contract method 0xf8bc93a5.
//
// Solidity: function NODE_REGISTRY_CONTRACT() view returns(bytes32)
func (_VaultFactory *VaultFactoryCallerSession) NODEREGISTRYCONTRACT() ([32]byte, error) {
	return _VaultFactory.Contract.NODEREGISTRYCONTRACT(&_VaultFactory.CallOpts)
}

// ComputeNodeELRewardVaultAddress is a free data retrieval call binding the contract method 0xca934f60.
//
// Solidity: function computeNodeELRewardVaultAddress(uint8 _poolId, uint256 _operatorId) view returns(address)
func (_VaultFactory *VaultFactoryCaller) ComputeNodeELRewardVaultAddress(opts *bind.CallOpts, _poolId uint8, _operatorId *big.Int) (common.Address, error) {
	var out []interface{}
	err := _VaultFactory.contract.Call(opts, &out, "computeNodeELRewardVaultAddress", _poolId, _operatorId)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// ComputeNodeELRewardVaultAddress is a free data retrieval call binding the contract method 0xca934f60.
//
// Solidity: function computeNodeELRewardVaultAddress(uint8 _poolId, uint256 _operatorId) view returns(address)
func (_VaultFactory *VaultFactorySession) ComputeNodeELRewardVaultAddress(_poolId uint8, _operatorId *big.Int) (common.Address, error) {
	return _VaultFactory.Contract.ComputeNodeELRewardVaultAddress(&_VaultFactory.CallOpts, _poolId, _operatorId)
}

// ComputeNodeELRewardVaultAddress is a free data retrieval call binding the contract method 0xca934f60.
//
// Solidity: function computeNodeELRewardVaultAddress(uint8 _poolId, uint256 _operatorId) view returns(address)
func (_VaultFactory *VaultFactoryCallerSession) ComputeNodeELRewardVaultAddress(_poolId uint8, _operatorId *big.Int) (common.Address, error) {
	return _VaultFactory.Contract.ComputeNodeELRewardVaultAddress(&_VaultFactory.CallOpts, _poolId, _operatorId)
}

// ComputeWithdrawVaultAddress is a free data retrieval call binding the contract method 0x74903b02.
//
// Solidity: function computeWithdrawVaultAddress(uint8 _poolId, uint256 _operatorId, uint256 _validatorCount) view returns(address)
func (_VaultFactory *VaultFactoryCaller) ComputeWithdrawVaultAddress(opts *bind.CallOpts, _poolId uint8, _operatorId *big.Int, _validatorCount *big.Int) (common.Address, error) {
	var out []interface{}
	err := _VaultFactory.contract.Call(opts, &out, "computeWithdrawVaultAddress", _poolId, _operatorId, _validatorCount)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// ComputeWithdrawVaultAddress is a free data retrieval call binding the contract method 0x74903b02.
//
// Solidity: function computeWithdrawVaultAddress(uint8 _poolId, uint256 _operatorId, uint256 _validatorCount) view returns(address)
func (_VaultFactory *VaultFactorySession) ComputeWithdrawVaultAddress(_poolId uint8, _operatorId *big.Int, _validatorCount *big.Int) (common.Address, error) {
	return _VaultFactory.Contract.ComputeWithdrawVaultAddress(&_VaultFactory.CallOpts, _poolId, _operatorId, _validatorCount)
}

// ComputeWithdrawVaultAddress is a free data retrieval call binding the contract method 0x74903b02.
//
// Solidity: function computeWithdrawVaultAddress(uint8 _poolId, uint256 _operatorId, uint256 _validatorCount) view returns(address)
func (_VaultFactory *VaultFactoryCallerSession) ComputeWithdrawVaultAddress(_poolId uint8, _operatorId *big.Int, _validatorCount *big.Int) (common.Address, error) {
	return _VaultFactory.Contract.ComputeWithdrawVaultAddress(&_VaultFactory.CallOpts, _poolId, _operatorId, _validatorCount)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_VaultFactory *VaultFactoryCaller) GetRoleAdmin(opts *bind.CallOpts, role [32]byte) ([32]byte, error) {
	var out []interface{}
	err := _VaultFactory.contract.Call(opts, &out, "getRoleAdmin", role)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_VaultFactory *VaultFactorySession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _VaultFactory.Contract.GetRoleAdmin(&_VaultFactory.CallOpts, role)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_VaultFactory *VaultFactoryCallerSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _VaultFactory.Contract.GetRoleAdmin(&_VaultFactory.CallOpts, role)
}

// GetValidatorWithdrawCredential is a free data retrieval call binding the contract method 0xae4e4e45.
//
// Solidity: function getValidatorWithdrawCredential(address _withdrawVault) pure returns(bytes)
func (_VaultFactory *VaultFactoryCaller) GetValidatorWithdrawCredential(opts *bind.CallOpts, _withdrawVault common.Address) ([]byte, error) {
	var out []interface{}
	err := _VaultFactory.contract.Call(opts, &out, "getValidatorWithdrawCredential", _withdrawVault)

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// GetValidatorWithdrawCredential is a free data retrieval call binding the contract method 0xae4e4e45.
//
// Solidity: function getValidatorWithdrawCredential(address _withdrawVault) pure returns(bytes)
func (_VaultFactory *VaultFactorySession) GetValidatorWithdrawCredential(_withdrawVault common.Address) ([]byte, error) {
	return _VaultFactory.Contract.GetValidatorWithdrawCredential(&_VaultFactory.CallOpts, _withdrawVault)
}

// GetValidatorWithdrawCredential is a free data retrieval call binding the contract method 0xae4e4e45.
//
// Solidity: function getValidatorWithdrawCredential(address _withdrawVault) pure returns(bytes)
func (_VaultFactory *VaultFactoryCallerSession) GetValidatorWithdrawCredential(_withdrawVault common.Address) ([]byte, error) {
	return _VaultFactory.Contract.GetValidatorWithdrawCredential(&_VaultFactory.CallOpts, _withdrawVault)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_VaultFactory *VaultFactoryCaller) HasRole(opts *bind.CallOpts, role [32]byte, account common.Address) (bool, error) {
	var out []interface{}
	err := _VaultFactory.contract.Call(opts, &out, "hasRole", role, account)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_VaultFactory *VaultFactorySession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _VaultFactory.Contract.HasRole(&_VaultFactory.CallOpts, role, account)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_VaultFactory *VaultFactoryCallerSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _VaultFactory.Contract.HasRole(&_VaultFactory.CallOpts, role, account)
}

// StaderConfig is a free data retrieval call binding the contract method 0x490ffa35.
//
// Solidity: function staderConfig() view returns(address)
func (_VaultFactory *VaultFactoryCaller) StaderConfig(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _VaultFactory.contract.Call(opts, &out, "staderConfig")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// StaderConfig is a free data retrieval call binding the contract method 0x490ffa35.
//
// Solidity: function staderConfig() view returns(address)
func (_VaultFactory *VaultFactorySession) StaderConfig() (common.Address, error) {
	return _VaultFactory.Contract.StaderConfig(&_VaultFactory.CallOpts)
}

// StaderConfig is a free data retrieval call binding the contract method 0x490ffa35.
//
// Solidity: function staderConfig() view returns(address)
func (_VaultFactory *VaultFactoryCallerSession) StaderConfig() (common.Address, error) {
	return _VaultFactory.Contract.StaderConfig(&_VaultFactory.CallOpts)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_VaultFactory *VaultFactoryCaller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _VaultFactory.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_VaultFactory *VaultFactorySession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _VaultFactory.Contract.SupportsInterface(&_VaultFactory.CallOpts, interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_VaultFactory *VaultFactoryCallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _VaultFactory.Contract.SupportsInterface(&_VaultFactory.CallOpts, interfaceId)
}

// VaultProxyImplementation is a free data retrieval call binding the contract method 0xee5fca3b.
//
// Solidity: function vaultProxyImplementation() view returns(address)
func (_VaultFactory *VaultFactoryCaller) VaultProxyImplementation(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _VaultFactory.contract.Call(opts, &out, "vaultProxyImplementation")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// VaultProxyImplementation is a free data retrieval call binding the contract method 0xee5fca3b.
//
// Solidity: function vaultProxyImplementation() view returns(address)
func (_VaultFactory *VaultFactorySession) VaultProxyImplementation() (common.Address, error) {
	return _VaultFactory.Contract.VaultProxyImplementation(&_VaultFactory.CallOpts)
}

// VaultProxyImplementation is a free data retrieval call binding the contract method 0xee5fca3b.
//
// Solidity: function vaultProxyImplementation() view returns(address)
func (_VaultFactory *VaultFactoryCallerSession) VaultProxyImplementation() (common.Address, error) {
	return _VaultFactory.Contract.VaultProxyImplementation(&_VaultFactory.CallOpts)
}

// DeployNodeELRewardVault is a paid mutator transaction binding the contract method 0x6a0b6881.
//
// Solidity: function deployNodeELRewardVault(uint8 _poolId, uint256 _operatorId) returns(address)
func (_VaultFactory *VaultFactoryTransactor) DeployNodeELRewardVault(opts *bind.TransactOpts, _poolId uint8, _operatorId *big.Int) (*types.Transaction, error) {
	return _VaultFactory.contract.Transact(opts, "deployNodeELRewardVault", _poolId, _operatorId)
}

// DeployNodeELRewardVault is a paid mutator transaction binding the contract method 0x6a0b6881.
//
// Solidity: function deployNodeELRewardVault(uint8 _poolId, uint256 _operatorId) returns(address)
func (_VaultFactory *VaultFactorySession) DeployNodeELRewardVault(_poolId uint8, _operatorId *big.Int) (*types.Transaction, error) {
	return _VaultFactory.Contract.DeployNodeELRewardVault(&_VaultFactory.TransactOpts, _poolId, _operatorId)
}

// DeployNodeELRewardVault is a paid mutator transaction binding the contract method 0x6a0b6881.
//
// Solidity: function deployNodeELRewardVault(uint8 _poolId, uint256 _operatorId) returns(address)
func (_VaultFactory *VaultFactoryTransactorSession) DeployNodeELRewardVault(_poolId uint8, _operatorId *big.Int) (*types.Transaction, error) {
	return _VaultFactory.Contract.DeployNodeELRewardVault(&_VaultFactory.TransactOpts, _poolId, _operatorId)
}

// DeployWithdrawVault is a paid mutator transaction binding the contract method 0x7f70ce0d.
//
// Solidity: function deployWithdrawVault(uint8 _poolId, uint256 _operatorId, uint256 _validatorCount, uint256 _validatorId) returns(address)
func (_VaultFactory *VaultFactoryTransactor) DeployWithdrawVault(opts *bind.TransactOpts, _poolId uint8, _operatorId *big.Int, _validatorCount *big.Int, _validatorId *big.Int) (*types.Transaction, error) {
	return _VaultFactory.contract.Transact(opts, "deployWithdrawVault", _poolId, _operatorId, _validatorCount, _validatorId)
}

// DeployWithdrawVault is a paid mutator transaction binding the contract method 0x7f70ce0d.
//
// Solidity: function deployWithdrawVault(uint8 _poolId, uint256 _operatorId, uint256 _validatorCount, uint256 _validatorId) returns(address)
func (_VaultFactory *VaultFactorySession) DeployWithdrawVault(_poolId uint8, _operatorId *big.Int, _validatorCount *big.Int, _validatorId *big.Int) (*types.Transaction, error) {
	return _VaultFactory.Contract.DeployWithdrawVault(&_VaultFactory.TransactOpts, _poolId, _operatorId, _validatorCount, _validatorId)
}

// DeployWithdrawVault is a paid mutator transaction binding the contract method 0x7f70ce0d.
//
// Solidity: function deployWithdrawVault(uint8 _poolId, uint256 _operatorId, uint256 _validatorCount, uint256 _validatorId) returns(address)
func (_VaultFactory *VaultFactoryTransactorSession) DeployWithdrawVault(_poolId uint8, _operatorId *big.Int, _validatorCount *big.Int, _validatorId *big.Int) (*types.Transaction, error) {
	return _VaultFactory.Contract.DeployWithdrawVault(&_VaultFactory.TransactOpts, _poolId, _operatorId, _validatorCount, _validatorId)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_VaultFactory *VaultFactoryTransactor) GrantRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _VaultFactory.contract.Transact(opts, "grantRole", role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_VaultFactory *VaultFactorySession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _VaultFactory.Contract.GrantRole(&_VaultFactory.TransactOpts, role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_VaultFactory *VaultFactoryTransactorSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _VaultFactory.Contract.GrantRole(&_VaultFactory.TransactOpts, role, account)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address account) returns()
func (_VaultFactory *VaultFactoryTransactor) RenounceRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _VaultFactory.contract.Transact(opts, "renounceRole", role, account)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address account) returns()
func (_VaultFactory *VaultFactorySession) RenounceRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _VaultFactory.Contract.RenounceRole(&_VaultFactory.TransactOpts, role, account)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address account) returns()
func (_VaultFactory *VaultFactoryTransactorSession) RenounceRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _VaultFactory.Contract.RenounceRole(&_VaultFactory.TransactOpts, role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_VaultFactory *VaultFactoryTransactor) RevokeRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _VaultFactory.contract.Transact(opts, "revokeRole", role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_VaultFactory *VaultFactorySession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _VaultFactory.Contract.RevokeRole(&_VaultFactory.TransactOpts, role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_VaultFactory *VaultFactoryTransactorSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _VaultFactory.Contract.RevokeRole(&_VaultFactory.TransactOpts, role, account)
}

// UpdateStaderConfig is a paid mutator transaction binding the contract method 0x9ee804cb.
//
// Solidity: function updateStaderConfig(address _staderConfig) returns()
func (_VaultFactory *VaultFactoryTransactor) UpdateStaderConfig(opts *bind.TransactOpts, _staderConfig common.Address) (*types.Transaction, error) {
	return _VaultFactory.contract.Transact(opts, "updateStaderConfig", _staderConfig)
}

// UpdateStaderConfig is a paid mutator transaction binding the contract method 0x9ee804cb.
//
// Solidity: function updateStaderConfig(address _staderConfig) returns()
func (_VaultFactory *VaultFactorySession) UpdateStaderConfig(_staderConfig common.Address) (*types.Transaction, error) {
	return _VaultFactory.Contract.UpdateStaderConfig(&_VaultFactory.TransactOpts, _staderConfig)
}

// UpdateStaderConfig is a paid mutator transaction binding the contract method 0x9ee804cb.
//
// Solidity: function updateStaderConfig(address _staderConfig) returns()
func (_VaultFactory *VaultFactoryTransactorSession) UpdateStaderConfig(_staderConfig common.Address) (*types.Transaction, error) {
	return _VaultFactory.Contract.UpdateStaderConfig(&_VaultFactory.TransactOpts, _staderConfig)
}

// UpdateVaultProxyAddress is a paid mutator transaction binding the contract method 0xab4f36d5.
//
// Solidity: function updateVaultProxyAddress(address _vaultProxyImpl) returns()
func (_VaultFactory *VaultFactoryTransactor) UpdateVaultProxyAddress(opts *bind.TransactOpts, _vaultProxyImpl common.Address) (*types.Transaction, error) {
	return _VaultFactory.contract.Transact(opts, "updateVaultProxyAddress", _vaultProxyImpl)
}

// UpdateVaultProxyAddress is a paid mutator transaction binding the contract method 0xab4f36d5.
//
// Solidity: function updateVaultProxyAddress(address _vaultProxyImpl) returns()
func (_VaultFactory *VaultFactorySession) UpdateVaultProxyAddress(_vaultProxyImpl common.Address) (*types.Transaction, error) {
	return _VaultFactory.Contract.UpdateVaultProxyAddress(&_VaultFactory.TransactOpts, _vaultProxyImpl)
}

// UpdateVaultProxyAddress is a paid mutator transaction binding the contract method 0xab4f36d5.
//
// Solidity: function updateVaultProxyAddress(address _vaultProxyImpl) returns()
func (_VaultFactory *VaultFactoryTransactorSession) UpdateVaultProxyAddress(_vaultProxyImpl common.Address) (*types.Transaction, error) {
	return _VaultFactory.Contract.UpdateVaultProxyAddress(&_VaultFactory.TransactOpts, _vaultProxyImpl)
}

// VaultFactoryInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the VaultFactory contract.
type VaultFactoryInitializedIterator struct {
	Event *VaultFactoryInitialized // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *VaultFactoryInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VaultFactoryInitialized)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(VaultFactoryInitialized)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *VaultFactoryInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VaultFactoryInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VaultFactoryInitialized represents a Initialized event raised by the VaultFactory contract.
type VaultFactoryInitialized struct {
	Version uint8
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_VaultFactory *VaultFactoryFilterer) FilterInitialized(opts *bind.FilterOpts) (*VaultFactoryInitializedIterator, error) {

	logs, sub, err := _VaultFactory.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &VaultFactoryInitializedIterator{contract: _VaultFactory.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_VaultFactory *VaultFactoryFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *VaultFactoryInitialized) (event.Subscription, error) {

	logs, sub, err := _VaultFactory.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VaultFactoryInitialized)
				if err := _VaultFactory.contract.UnpackLog(event, "Initialized", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseInitialized is a log parse operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_VaultFactory *VaultFactoryFilterer) ParseInitialized(log types.Log) (*VaultFactoryInitialized, error) {
	event := new(VaultFactoryInitialized)
	if err := _VaultFactory.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// VaultFactoryNodeELRewardVaultCreatedIterator is returned from FilterNodeELRewardVaultCreated and is used to iterate over the raw logs and unpacked data for NodeELRewardVaultCreated events raised by the VaultFactory contract.
type VaultFactoryNodeELRewardVaultCreatedIterator struct {
	Event *VaultFactoryNodeELRewardVaultCreated // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *VaultFactoryNodeELRewardVaultCreatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VaultFactoryNodeELRewardVaultCreated)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(VaultFactoryNodeELRewardVaultCreated)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *VaultFactoryNodeELRewardVaultCreatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VaultFactoryNodeELRewardVaultCreatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VaultFactoryNodeELRewardVaultCreated represents a NodeELRewardVaultCreated event raised by the VaultFactory contract.
type VaultFactoryNodeELRewardVaultCreated struct {
	NodeDistributor common.Address
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterNodeELRewardVaultCreated is a free log retrieval operation binding the contract event 0x534e63617d48745c45efd754f53f9292453257107e03d0140db96fc8e3bf7ec9.
//
// Solidity: event NodeELRewardVaultCreated(address nodeDistributor)
func (_VaultFactory *VaultFactoryFilterer) FilterNodeELRewardVaultCreated(opts *bind.FilterOpts) (*VaultFactoryNodeELRewardVaultCreatedIterator, error) {

	logs, sub, err := _VaultFactory.contract.FilterLogs(opts, "NodeELRewardVaultCreated")
	if err != nil {
		return nil, err
	}
	return &VaultFactoryNodeELRewardVaultCreatedIterator{contract: _VaultFactory.contract, event: "NodeELRewardVaultCreated", logs: logs, sub: sub}, nil
}

// WatchNodeELRewardVaultCreated is a free log subscription operation binding the contract event 0x534e63617d48745c45efd754f53f9292453257107e03d0140db96fc8e3bf7ec9.
//
// Solidity: event NodeELRewardVaultCreated(address nodeDistributor)
func (_VaultFactory *VaultFactoryFilterer) WatchNodeELRewardVaultCreated(opts *bind.WatchOpts, sink chan<- *VaultFactoryNodeELRewardVaultCreated) (event.Subscription, error) {

	logs, sub, err := _VaultFactory.contract.WatchLogs(opts, "NodeELRewardVaultCreated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VaultFactoryNodeELRewardVaultCreated)
				if err := _VaultFactory.contract.UnpackLog(event, "NodeELRewardVaultCreated", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseNodeELRewardVaultCreated is a log parse operation binding the contract event 0x534e63617d48745c45efd754f53f9292453257107e03d0140db96fc8e3bf7ec9.
//
// Solidity: event NodeELRewardVaultCreated(address nodeDistributor)
func (_VaultFactory *VaultFactoryFilterer) ParseNodeELRewardVaultCreated(log types.Log) (*VaultFactoryNodeELRewardVaultCreated, error) {
	event := new(VaultFactoryNodeELRewardVaultCreated)
	if err := _VaultFactory.contract.UnpackLog(event, "NodeELRewardVaultCreated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// VaultFactoryRoleAdminChangedIterator is returned from FilterRoleAdminChanged and is used to iterate over the raw logs and unpacked data for RoleAdminChanged events raised by the VaultFactory contract.
type VaultFactoryRoleAdminChangedIterator struct {
	Event *VaultFactoryRoleAdminChanged // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *VaultFactoryRoleAdminChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VaultFactoryRoleAdminChanged)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(VaultFactoryRoleAdminChanged)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *VaultFactoryRoleAdminChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VaultFactoryRoleAdminChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VaultFactoryRoleAdminChanged represents a RoleAdminChanged event raised by the VaultFactory contract.
type VaultFactoryRoleAdminChanged struct {
	Role              [32]byte
	PreviousAdminRole [32]byte
	NewAdminRole      [32]byte
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterRoleAdminChanged is a free log retrieval operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_VaultFactory *VaultFactoryFilterer) FilterRoleAdminChanged(opts *bind.FilterOpts, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (*VaultFactoryRoleAdminChangedIterator, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var previousAdminRoleRule []interface{}
	for _, previousAdminRoleItem := range previousAdminRole {
		previousAdminRoleRule = append(previousAdminRoleRule, previousAdminRoleItem)
	}
	var newAdminRoleRule []interface{}
	for _, newAdminRoleItem := range newAdminRole {
		newAdminRoleRule = append(newAdminRoleRule, newAdminRoleItem)
	}

	logs, sub, err := _VaultFactory.contract.FilterLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return &VaultFactoryRoleAdminChangedIterator{contract: _VaultFactory.contract, event: "RoleAdminChanged", logs: logs, sub: sub}, nil
}

// WatchRoleAdminChanged is a free log subscription operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_VaultFactory *VaultFactoryFilterer) WatchRoleAdminChanged(opts *bind.WatchOpts, sink chan<- *VaultFactoryRoleAdminChanged, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (event.Subscription, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var previousAdminRoleRule []interface{}
	for _, previousAdminRoleItem := range previousAdminRole {
		previousAdminRoleRule = append(previousAdminRoleRule, previousAdminRoleItem)
	}
	var newAdminRoleRule []interface{}
	for _, newAdminRoleItem := range newAdminRole {
		newAdminRoleRule = append(newAdminRoleRule, newAdminRoleItem)
	}

	logs, sub, err := _VaultFactory.contract.WatchLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VaultFactoryRoleAdminChanged)
				if err := _VaultFactory.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseRoleAdminChanged is a log parse operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_VaultFactory *VaultFactoryFilterer) ParseRoleAdminChanged(log types.Log) (*VaultFactoryRoleAdminChanged, error) {
	event := new(VaultFactoryRoleAdminChanged)
	if err := _VaultFactory.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// VaultFactoryRoleGrantedIterator is returned from FilterRoleGranted and is used to iterate over the raw logs and unpacked data for RoleGranted events raised by the VaultFactory contract.
type VaultFactoryRoleGrantedIterator struct {
	Event *VaultFactoryRoleGranted // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *VaultFactoryRoleGrantedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VaultFactoryRoleGranted)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(VaultFactoryRoleGranted)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *VaultFactoryRoleGrantedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VaultFactoryRoleGrantedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VaultFactoryRoleGranted represents a RoleGranted event raised by the VaultFactory contract.
type VaultFactoryRoleGranted struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleGranted is a free log retrieval operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_VaultFactory *VaultFactoryFilterer) FilterRoleGranted(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*VaultFactoryRoleGrantedIterator, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _VaultFactory.contract.FilterLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &VaultFactoryRoleGrantedIterator{contract: _VaultFactory.contract, event: "RoleGranted", logs: logs, sub: sub}, nil
}

// WatchRoleGranted is a free log subscription operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_VaultFactory *VaultFactoryFilterer) WatchRoleGranted(opts *bind.WatchOpts, sink chan<- *VaultFactoryRoleGranted, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _VaultFactory.contract.WatchLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VaultFactoryRoleGranted)
				if err := _VaultFactory.contract.UnpackLog(event, "RoleGranted", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseRoleGranted is a log parse operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_VaultFactory *VaultFactoryFilterer) ParseRoleGranted(log types.Log) (*VaultFactoryRoleGranted, error) {
	event := new(VaultFactoryRoleGranted)
	if err := _VaultFactory.contract.UnpackLog(event, "RoleGranted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// VaultFactoryRoleRevokedIterator is returned from FilterRoleRevoked and is used to iterate over the raw logs and unpacked data for RoleRevoked events raised by the VaultFactory contract.
type VaultFactoryRoleRevokedIterator struct {
	Event *VaultFactoryRoleRevoked // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *VaultFactoryRoleRevokedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VaultFactoryRoleRevoked)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(VaultFactoryRoleRevoked)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *VaultFactoryRoleRevokedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VaultFactoryRoleRevokedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VaultFactoryRoleRevoked represents a RoleRevoked event raised by the VaultFactory contract.
type VaultFactoryRoleRevoked struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleRevoked is a free log retrieval operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_VaultFactory *VaultFactoryFilterer) FilterRoleRevoked(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*VaultFactoryRoleRevokedIterator, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _VaultFactory.contract.FilterLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &VaultFactoryRoleRevokedIterator{contract: _VaultFactory.contract, event: "RoleRevoked", logs: logs, sub: sub}, nil
}

// WatchRoleRevoked is a free log subscription operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_VaultFactory *VaultFactoryFilterer) WatchRoleRevoked(opts *bind.WatchOpts, sink chan<- *VaultFactoryRoleRevoked, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _VaultFactory.contract.WatchLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VaultFactoryRoleRevoked)
				if err := _VaultFactory.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseRoleRevoked is a log parse operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_VaultFactory *VaultFactoryFilterer) ParseRoleRevoked(log types.Log) (*VaultFactoryRoleRevoked, error) {
	event := new(VaultFactoryRoleRevoked)
	if err := _VaultFactory.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// VaultFactoryUpdatedStaderConfigIterator is returned from FilterUpdatedStaderConfig and is used to iterate over the raw logs and unpacked data for UpdatedStaderConfig events raised by the VaultFactory contract.
type VaultFactoryUpdatedStaderConfigIterator struct {
	Event *VaultFactoryUpdatedStaderConfig // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *VaultFactoryUpdatedStaderConfigIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VaultFactoryUpdatedStaderConfig)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(VaultFactoryUpdatedStaderConfig)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *VaultFactoryUpdatedStaderConfigIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VaultFactoryUpdatedStaderConfigIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VaultFactoryUpdatedStaderConfig represents a UpdatedStaderConfig event raised by the VaultFactory contract.
type VaultFactoryUpdatedStaderConfig struct {
	StaderConfig common.Address
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterUpdatedStaderConfig is a free log retrieval operation binding the contract event 0xdb2219043d7b197cb235f1af0cf6d782d77dee3de19e3f4fb6d39aae633b4485.
//
// Solidity: event UpdatedStaderConfig(address staderConfig)
func (_VaultFactory *VaultFactoryFilterer) FilterUpdatedStaderConfig(opts *bind.FilterOpts) (*VaultFactoryUpdatedStaderConfigIterator, error) {

	logs, sub, err := _VaultFactory.contract.FilterLogs(opts, "UpdatedStaderConfig")
	if err != nil {
		return nil, err
	}
	return &VaultFactoryUpdatedStaderConfigIterator{contract: _VaultFactory.contract, event: "UpdatedStaderConfig", logs: logs, sub: sub}, nil
}

// WatchUpdatedStaderConfig is a free log subscription operation binding the contract event 0xdb2219043d7b197cb235f1af0cf6d782d77dee3de19e3f4fb6d39aae633b4485.
//
// Solidity: event UpdatedStaderConfig(address staderConfig)
func (_VaultFactory *VaultFactoryFilterer) WatchUpdatedStaderConfig(opts *bind.WatchOpts, sink chan<- *VaultFactoryUpdatedStaderConfig) (event.Subscription, error) {

	logs, sub, err := _VaultFactory.contract.WatchLogs(opts, "UpdatedStaderConfig")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VaultFactoryUpdatedStaderConfig)
				if err := _VaultFactory.contract.UnpackLog(event, "UpdatedStaderConfig", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseUpdatedStaderConfig is a log parse operation binding the contract event 0xdb2219043d7b197cb235f1af0cf6d782d77dee3de19e3f4fb6d39aae633b4485.
//
// Solidity: event UpdatedStaderConfig(address staderConfig)
func (_VaultFactory *VaultFactoryFilterer) ParseUpdatedStaderConfig(log types.Log) (*VaultFactoryUpdatedStaderConfig, error) {
	event := new(VaultFactoryUpdatedStaderConfig)
	if err := _VaultFactory.contract.UnpackLog(event, "UpdatedStaderConfig", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// VaultFactoryUpdatedVaultProxyImplementationIterator is returned from FilterUpdatedVaultProxyImplementation and is used to iterate over the raw logs and unpacked data for UpdatedVaultProxyImplementation events raised by the VaultFactory contract.
type VaultFactoryUpdatedVaultProxyImplementationIterator struct {
	Event *VaultFactoryUpdatedVaultProxyImplementation // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *VaultFactoryUpdatedVaultProxyImplementationIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VaultFactoryUpdatedVaultProxyImplementation)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(VaultFactoryUpdatedVaultProxyImplementation)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *VaultFactoryUpdatedVaultProxyImplementationIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VaultFactoryUpdatedVaultProxyImplementationIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VaultFactoryUpdatedVaultProxyImplementation represents a UpdatedVaultProxyImplementation event raised by the VaultFactory contract.
type VaultFactoryUpdatedVaultProxyImplementation struct {
	VaultProxyImplementation common.Address
	Raw                      types.Log // Blockchain specific contextual infos
}

// FilterUpdatedVaultProxyImplementation is a free log retrieval operation binding the contract event 0x77358c5e9bc05f53688720682a283a5de55b771a614836c83cb1428deb2ece2d.
//
// Solidity: event UpdatedVaultProxyImplementation(address vaultProxyImplementation)
func (_VaultFactory *VaultFactoryFilterer) FilterUpdatedVaultProxyImplementation(opts *bind.FilterOpts) (*VaultFactoryUpdatedVaultProxyImplementationIterator, error) {

	logs, sub, err := _VaultFactory.contract.FilterLogs(opts, "UpdatedVaultProxyImplementation")
	if err != nil {
		return nil, err
	}
	return &VaultFactoryUpdatedVaultProxyImplementationIterator{contract: _VaultFactory.contract, event: "UpdatedVaultProxyImplementation", logs: logs, sub: sub}, nil
}

// WatchUpdatedVaultProxyImplementation is a free log subscription operation binding the contract event 0x77358c5e9bc05f53688720682a283a5de55b771a614836c83cb1428deb2ece2d.
//
// Solidity: event UpdatedVaultProxyImplementation(address vaultProxyImplementation)
func (_VaultFactory *VaultFactoryFilterer) WatchUpdatedVaultProxyImplementation(opts *bind.WatchOpts, sink chan<- *VaultFactoryUpdatedVaultProxyImplementation) (event.Subscription, error) {

	logs, sub, err := _VaultFactory.contract.WatchLogs(opts, "UpdatedVaultProxyImplementation")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VaultFactoryUpdatedVaultProxyImplementation)
				if err := _VaultFactory.contract.UnpackLog(event, "UpdatedVaultProxyImplementation", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseUpdatedVaultProxyImplementation is a log parse operation binding the contract event 0x77358c5e9bc05f53688720682a283a5de55b771a614836c83cb1428deb2ece2d.
//
// Solidity: event UpdatedVaultProxyImplementation(address vaultProxyImplementation)
func (_VaultFactory *VaultFactoryFilterer) ParseUpdatedVaultProxyImplementation(log types.Log) (*VaultFactoryUpdatedVaultProxyImplementation, error) {
	event := new(VaultFactoryUpdatedVaultProxyImplementation)
	if err := _VaultFactory.contract.UnpackLog(event, "UpdatedVaultProxyImplementation", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// VaultFactoryWithdrawVaultCreatedIterator is returned from FilterWithdrawVaultCreated and is used to iterate over the raw logs and unpacked data for WithdrawVaultCreated events raised by the VaultFactory contract.
type VaultFactoryWithdrawVaultCreatedIterator struct {
	Event *VaultFactoryWithdrawVaultCreated // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *VaultFactoryWithdrawVaultCreatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VaultFactoryWithdrawVaultCreated)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(VaultFactoryWithdrawVaultCreated)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *VaultFactoryWithdrawVaultCreatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VaultFactoryWithdrawVaultCreatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VaultFactoryWithdrawVaultCreated represents a WithdrawVaultCreated event raised by the VaultFactory contract.
type VaultFactoryWithdrawVaultCreated struct {
	WithdrawVault common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterWithdrawVaultCreated is a free log retrieval operation binding the contract event 0x3961159e01aea37eec106905b883c13f18c7c19abfde04fef6835c94d2af51b9.
//
// Solidity: event WithdrawVaultCreated(address withdrawVault)
func (_VaultFactory *VaultFactoryFilterer) FilterWithdrawVaultCreated(opts *bind.FilterOpts) (*VaultFactoryWithdrawVaultCreatedIterator, error) {

	logs, sub, err := _VaultFactory.contract.FilterLogs(opts, "WithdrawVaultCreated")
	if err != nil {
		return nil, err
	}
	return &VaultFactoryWithdrawVaultCreatedIterator{contract: _VaultFactory.contract, event: "WithdrawVaultCreated", logs: logs, sub: sub}, nil
}

// WatchWithdrawVaultCreated is a free log subscription operation binding the contract event 0x3961159e01aea37eec106905b883c13f18c7c19abfde04fef6835c94d2af51b9.
//
// Solidity: event WithdrawVaultCreated(address withdrawVault)
func (_VaultFactory *VaultFactoryFilterer) WatchWithdrawVaultCreated(opts *bind.WatchOpts, sink chan<- *VaultFactoryWithdrawVaultCreated) (event.Subscription, error) {

	logs, sub, err := _VaultFactory.contract.WatchLogs(opts, "WithdrawVaultCreated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VaultFactoryWithdrawVaultCreated)
				if err := _VaultFactory.contract.UnpackLog(event, "WithdrawVaultCreated", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseWithdrawVaultCreated is a log parse operation binding the contract event 0x3961159e01aea37eec106905b883c13f18c7c19abfde04fef6835c94d2af51b9.
//
// Solidity: event WithdrawVaultCreated(address withdrawVault)
func (_VaultFactory *VaultFactoryFilterer) ParseWithdrawVaultCreated(log types.Log) (*VaultFactoryWithdrawVaultCreated, error) {
	event := new(VaultFactoryWithdrawVaultCreated)
	if err := _VaultFactory.contract.UnpackLog(event, "WithdrawVaultCreated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
