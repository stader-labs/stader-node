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

// SDCollateralMetaData contains all meta data concerning the SDCollateral contract.
var SDCollateralMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_admin\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_staderConfig\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"CallerNotManager\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"CallerNotWithdrawVault\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"operatorSDCollateral\",\"type\":\"uint256\"}],\"name\":\"InsufficientSDToWithdraw\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidPoolId\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidPoolLimit\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NoStateChange\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"SDTransferFailed\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ZeroAddress\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"version\",\"type\":\"uint8\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"previousAdminRole\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"newAdminRole\",\"type\":\"bytes32\"}],\"name\":\"RoleAdminChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleGranted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleRevoked\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"sdAmount\",\"type\":\"uint256\"}],\"name\":\"SDDeposited\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"auction\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"sdSlashed\",\"type\":\"uint256\"}],\"name\":\"SDSlashed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"sdAmount\",\"type\":\"uint256\"}],\"name\":\"SDWithdrawn\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"poolId\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"}],\"name\":\"UpdatedPoolIdForOperator\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"poolId\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"minThreshold\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"withdrawThreshold\",\"type\":\"uint256\"}],\"name\":\"UpdatedPoolThreshold\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"staderConfig\",\"type\":\"address\"}],\"name\":\"UpdatedStaderConfig\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"DEFAULT_ADMIN_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_ethAmount\",\"type\":\"uint256\"}],\"name\":\"convertETHToSD\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_sdAmount\",\"type\":\"uint256\"}],\"name\":\"convertSDToETH\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_sdAmount\",\"type\":\"uint256\"}],\"name\":\"depositSDAsCollateral\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"_poolId\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"_numValidator\",\"type\":\"uint256\"}],\"name\":\"getMinimumSDToBond\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"_minSDToBond\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_operator\",\"type\":\"address\"}],\"name\":\"getOperatorWithdrawThreshold\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"operatorWithdrawThreshold\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_operator\",\"type\":\"address\"},{\"internalType\":\"uint8\",\"name\":\"_poolId\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"_numValidator\",\"type\":\"uint256\"}],\"name\":\"getRemainingSDToBond\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_operator\",\"type\":\"address\"}],\"name\":\"getRewardEligibleSD\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"_rewardEligibleSD\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"}],\"name\":\"getRoleAdmin\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"grantRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_operator\",\"type\":\"address\"},{\"internalType\":\"uint8\",\"name\":\"_poolId\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"_numValidator\",\"type\":\"uint256\"}],\"name\":\"hasEnoughSDCollateral\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"hasRole\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"maxApproveSD\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"operatorSDBalance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"name\":\"poolThresholdbyPoolId\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"minThreshold\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxThreshold\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"withdrawThreshold\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"units\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"renounceRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"revokeRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_validatorId\",\"type\":\"uint256\"},{\"internalType\":\"uint8\",\"name\":\"_poolId\",\"type\":\"uint8\"}],\"name\":\"slashValidatorSD\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"staderConfig\",\"outputs\":[{\"internalType\":\"contractIStaderConfig\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"_poolId\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"_minThreshold\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_maxThreshold\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_withdrawThreshold\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"_units\",\"type\":\"string\"}],\"name\":\"updatePoolThreshold\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_staderConfig\",\"type\":\"address\"}],\"name\":\"updateStaderConfig\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_requestedSD\",\"type\":\"uint256\"}],\"name\":\"withdraw\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x608060405234801562000010575f80fd5b50604051620026ad380380620026ad8339810160408190526200003391620003d6565b5f54610100900460ff16158080156200005257505f54600160ff909116105b806200006d5750303b1580156200006d57505f5460ff166001145b620000d65760405162461bcd60e51b815260206004820152602e60248201527f496e697469616c697a61626c653a20636f6e747261637420697320616c72656160448201526d191e481a5b9a5d1a585b1a5e995960921b60648201526084015b60405180910390fd5b5f805460ff191660011790558015620000f8575f805461ff0019166101001790555b6200010383620001cb565b6200010e82620001cb565b62000118620001f6565b6200012262000252565b60c980546001600160a01b0319166001600160a01b038416179055620001495f84620002b6565b6040516001600160a01b038316907fdb2219043d7b197cb235f1af0cf6d782d77dee3de19e3f4fb6d39aae633b4485905f90a28015620001c2575f805461ff0019169055604051600181527f7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb38474024989060200160405180910390a15b5050506200040c565b6001600160a01b038116620001f35760405163d92e233d60e01b815260040160405180910390fd5b50565b5f54610100900460ff16620002505760405162461bcd60e51b815260206004820152602b60248201525f805160206200268d83398151915260448201526a6e697469616c697a696e6760a81b6064820152608401620000cd565b565b5f54610100900460ff16620002ac5760405162461bcd60e51b815260206004820152602b60248201525f805160206200268d83398151915260448201526a6e697469616c697a696e6760a81b6064820152608401620000cd565b6200025062000359565b5f8281526065602090815260408083206001600160a01b038516845290915290205460ff1662000355575f8281526065602090815260408083206001600160a01b03851684529091529020805460ff19166001179055620003143390565b6001600160a01b0316816001600160a01b0316837f2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d60405160405180910390a45b5050565b5f54610100900460ff16620003b35760405162461bcd60e51b815260206004820152602b60248201525f805160206200268d83398151915260448201526a6e697469616c697a696e6760a81b6064820152608401620000cd565b6001609755565b80516001600160a01b0381168114620003d1575f80fd5b919050565b5f8060408385031215620003e8575f80fd5b620003f383620003ba565b91506200040360208401620003ba565b90509250929050565b612273806200041a5f395ff3fe608060405234801561000f575f80fd5b5060043610610148575f3560e01c80638a9b3738116100bf578063d547741f11610079578063d547741f146102d4578063dfdafccb146102e7578063e0412f0e146102fa578063e614e17c1461030d578063f9af40b814610320578063fcb7e0321461033f575f80fd5b80638a9b37381461025e57806391d14854146102815780639871a30a146102945780639ee804cb146102a7578063a217fddf146102ba578063b178e38e146102c1575f80fd5b806336568abe1161011057806336568abe146101df578063379b727e146101f25780633909afd3146102055780633e04cd3514610218578063490ffa35146102205780634c538f581461024b575f80fd5b806301ffc9a71461014c578063248a9ca3146101745780632e1a7d4d146101a45780632f2ff15d146101b9578063351691ab146101cc575b5f80fd5b61015f61015a366004611acc565b610352565b60405190151581526020015b60405180910390f35b610196610182366004611af3565b5f9081526065602052604090206001015490565b60405190815260200161016b565b6101b76101b2366004611af3565b610388565b005b6101b76101c7366004611b1e565b610542565b6101966101da366004611b5a565b61056b565b6101b76101ed366004611b1e565b6105b5565b610196610200366004611b98565b610633565b610196610213366004611bc2565b61066d565b6101b7610709565b60c954610233906001600160a01b031681565b6040516001600160a01b03909116815260200161016b565b6101b7610259366004611bdd565b61087c565b61027161026c366004611c00565b6108e4565b60405161016b9493929190611c68565b61015f61028f366004611b1e565b610994565b6101966102a2366004611bc2565b6109be565b6101b76102b5366004611bc2565b610a06565b6101965f81565b61015f6102cf366004611b5a565b610a92565b6101b76102e2366004611b1e565b610aa7565b6101966102f5366004611af3565b610acb565b6101b7610308366004611d02565b610c2a565b61019661031b366004611af3565b610d13565b61019661032e366004611bc2565b60cb6020525f908152604090205481565b6101b761034d366004611af3565b610e69565b5f6001600160e01b03198216637965db0b60e01b148061038257506301ffc9a760e01b6001600160e01b03198316145b92915050565b335f81815260cb6020526040902054826103a1836109be565b6103ab9190611dbe565b8110156103d35760405163f669f60360e01b8152600481018290526024015b60405180910390fd5b6001600160a01b0382165f90815260cb6020526040812080548592906103fa908490611dd1565b909155505060c9546040805163381a7dc560e21b815290516001600160a01b039092169163e069f714916004808201926020929091908290030181865afa158015610447573d5f803e3d5ffd5b505050506040513d601f19601f8201168201806040525081019061046b9190611df4565b60405163a9059cbb60e01b81526001600160a01b03848116600483015260248201869052919091169063a9059cbb906044016020604051808303815f875af11580156104b9573d5f803e3d5ffd5b505050506040513d601f19601f820116820180604052508101906104dd9190611e1e565b6104fa5760405163d3544e3f60e01b815260040160405180910390fd5b816001600160a01b03167f48c1f846fa4bc05385324ee60316f9c6778ed2b5f205a6319678a609b87676078460405161053591815260200190565b60405180910390a2505050565b5f8281526065602052604090206001015461055c81610fd4565b6105668383610fe1565b505050565b6001600160a01b0383165f90815260cb60205260408120548161058e8585610633565b9050808210156105a7576105a28282611dd1565b6105a9565b5f5b925050505b9392505050565b6001600160a01b03811633146106255760405162461bcd60e51b815260206004820152602f60248201527f416363657373436f6e74726f6c3a2063616e206f6e6c792072656e6f756e636560448201526e103937b632b9903337b91039b2b63360891b60648201526084016103ca565b61062f8282611066565b5050565b5f61063d836110cc565b60ff83165f90815260ca60205260409020805461065990610d13565b91506106658383611e37565b949350505050565b5f805f6106798461110c565b9250509150610687826110cc565b60ff82165f90815260ca6020526040812080549091906106a690610d13565b6106b09084611e37565b90505f6106c08360010154610d13565b6106ca9085611e37565b6001600160a01b0388165f90815260cb60205260409020549091508281106106fb576106f681836113cf565b6106fd565b5f5b98975050505050505050565b60c9546107209033906001600160a01b03166113e4565b60c95460408051630db055fd60e21b815290515f926001600160a01b0316916336c157f49160048083019260209291908290030181865afa158015610767573d5f803e3d5ffd5b505050506040513d601f19601f8201168201806040525081019061078b9190611df4565b905061079681611469565b60c95f9054906101000a90046001600160a01b03166001600160a01b031663e069f7146040518163ffffffff1660e01b8152600401602060405180830381865afa1580156107e6573d5f803e3d5ffd5b505050506040513d601f19601f8201168201806040525081019061080a9190611df4565b60405163095ea7b360e01b81526001600160a01b0383811660048301525f196024830152919091169063095ea7b3906044016020604051808303815f875af1158015610858573d5f803e3d5ffd5b505050506040513d601f19601f8201168201806040525081019061062f9190611e1e565b610884611490565b60c9545f906108a1908390859033906001600160a01b03166114e9565b90506108ac826110cc565b60ff82165f90815260ca6020526040812080549091906108cb90610d13565b90506108d783826116e1565b50505061062f6001609755565b60ca6020525f908152604090208054600182015460028301546003840180549394929391929161091390611e4e565b80601f016020809104026020016040519081016040528092919081815260200182805461093f90611e4e565b801561098a5780601f106109615761010080835404028352916020019161098a565b820191905f5260205f20905b81548152906001019060200180831161096d57829003601f168201915b5050505050905084565b5f9182526065602090815260408084206001600160a01b0393909316845291905290205460ff1690565b5f805f6109ca8461110c565b92505091506109d8826110cc565b60ff82165f90815260ca6020526040902060028101546109fd9061031b908490611e37565b95945050505050565b5f610a1081610fd4565b610a1982611469565b60c9546001600160a01b0390811690831603610a485760405163a28a88c160e01b815260040160405180910390fd5b60c980546001600160a01b0319166001600160a01b0384169081179091556040517fdb2219043d7b197cb235f1af0cf6d782d77dee3de19e3f4fb6d39aae633b4485905f90a25050565b5f610a9e84848461056b565b15949350505050565b5f82815260656020526040902060010154610ac181610fd4565b6105668383611066565b5f8060c95f9054906101000a90046001600160a01b03166001600160a01b031663defd024d6040518163ffffffff1660e01b8152600401602060405180830381865afa158015610b1d573d5f803e3d5ffd5b505050506040513d601f19601f82011682018060405250810190610b419190611df4565b6001600160a01b031663a6870e5b6040518163ffffffff1660e01b8152600401602060405180830381865afa158015610b7c573d5f803e3d5ffd5b505050506040513d601f19601f82011682018060405250810190610ba09190611e86565b905060c95f9054906101000a90046001600160a01b03166001600160a01b031663f0141d846040518163ffffffff1660e01b8152600401602060405180830381865afa158015610bf2573d5f803e3d5ffd5b505050506040513d601f19601f82011682018060405250810190610c169190611e86565b610c208285611e37565b6105ae9190611e9d565b60c954610c419033906001600160a01b03166113e4565b81841180610c4e57508284115b15610c6c576040516375075b2960e11b815260040160405180910390fd5b6040805160808101825285815260208082018681528284018681526060840186815260ff8b165f90815260ca9094529490922083518155905160018201559051600282015591519091906003820190610cc59082611f09565b50506040805160ff88168152602081018790529081018490527f18757dd1fbfe2ad823e1bd4de3f8a2ee76b49f92f6aa34cc7cbf717cdf4d1758915060600160405180910390a15050505050565b5f8060c95f9054906101000a90046001600160a01b03166001600160a01b031663defd024d6040518163ffffffff1660e01b8152600401602060405180830381865afa158015610d65573d5f803e3d5ffd5b505050506040513d601f19601f82011682018060405250810190610d899190611df4565b6001600160a01b031663a6870e5b6040518163ffffffff1660e01b8152600401602060405180830381865afa158015610dc4573d5f803e3d5ffd5b505050506040513d601f19601f82011682018060405250810190610de89190611e86565b90508060c95f9054906101000a90046001600160a01b03166001600160a01b031663f0141d846040518163ffffffff1660e01b8152600401602060405180830381865afa158015610e3b573d5f803e3d5ffd5b505050506040513d601f19601f82011682018060405250810190610e5f9190611e86565b610c209085611e37565b335f81815260cb602052604081208054849290610e87908490611dbe565b909155505060c9546040805163381a7dc560e21b815290516001600160a01b039092169163e069f714916004808201926020929091908290030181865afa158015610ed4573d5f803e3d5ffd5b505050506040513d601f19601f82011682018060405250810190610ef89190611df4565b6040516323b872dd60e01b81526001600160a01b0383811660048301523060248301526001604483015291909116906323b872dd906064016020604051808303815f875af1158015610f4c573d5f803e3d5ffd5b505050506040513d601f19601f82011682018060405250810190610f709190611e1e565b610f8d5760405163d3544e3f60e01b815260040160405180910390fd5b806001600160a01b03167f112973aa2e3b182a34447572d830c1e97afc414558a08f33554be8e54522425983604051610fc891815260200190565b60405180910390a25050565b610fde81336118cb565b50565b610feb8282610994565b61062f575f8281526065602090815260408083206001600160a01b03851684529091529020805460ff191660011790556110223390565b6001600160a01b0316816001600160a01b0316837f2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d60405160405180910390a45050565b6110708282610994565b1561062f575f8281526065602090815260408083206001600160a01b0385168085529252808320805460ff1916905551339285917ff6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b9190a45050565b60ff81165f90815260ca6020526040902060030180546110eb90611e4e565b90505f03610fde5760405163015f4fdd60e31b815260040160405180910390fd5b5f805f8060c95f9054906101000a90046001600160a01b03166001600160a01b0316636ccb9d706040518163ffffffff1660e01b8152600401602060405180830381865afa158015611160573d5f803e3d5ffd5b505050506040513d601f19601f820116820180604052508101906111849190611df4565b604051634721e29d60e11b81526001600160a01b03878116600483015291925090821690638e43c53a90602401602060405180830381865afa1580156111cc573d5f803e3d5ffd5b505050506040513d601f19601f820116820180604052508101906111f09190611fc5565b60405163133a0ab960e31b815260ff821660048201529094505f906001600160a01b038316906399d055c890602401602060405180830381865afa15801561123a573d5f803e3d5ffd5b505050506040513d601f19601f8201168201806040525081019061125e9190611df4565b604051636564598360e11b81526001600160a01b0388811660048301529192509082169063cac8b30690602401602060405180830381865afa1580156112a6573d5f803e3d5ffd5b505050506040513d601f19601f820116820180604052508101906112ca9190611e86565b9350816001600160a01b0316635d713ec386885f856001600160a01b031663c34ade5c8a6040518263ffffffff1660e01b815260040161130c91815260200190565b602060405180830381865afa158015611327573d5f803e3d5ffd5b505050506040513d601f19601f8201168201806040525081019061134b9190611e86565b6040516001600160e01b031960e087901b16815260ff90941660048501526001600160a01b03909216602484015260448301526064820152608401602060405180830381865afa1580156113a1573d5f803e3d5ffd5b505050506040513d601f19601f820116820180604052508101906113c59190611e86565b9496939550505050565b5f8183106113dd57816105ae565b5090919050565b6040516318903ee760e21b81526001600160a01b038381166004830152821690636240fb9c90602401602060405180830381865afa158015611428573d5f803e3d5ffd5b505050506040513d601f19601f8201168201806040525081019061144c9190611e1e565b61062f5760405163c4230ae360e01b815260040160405180910390fd5b6001600160a01b038116610fde5760405163d92e233d60e01b815260040160405180910390fd5b6002609754036114e25760405162461bcd60e51b815260206004820152601f60248201527f5265656e7472616e637947756172643a207265656e7472616e742063616c6c0060448201526064016103ca565b6002609755565b5f80826001600160a01b0316636ccb9d706040518163ffffffff1660e01b8152600401602060405180830381865afa158015611527573d5f803e3d5ffd5b505050506040513d601f19601f8201168201806040525081019061154b9190611df4565b60405163133a0ab960e31b815260ff881660048201526001600160a01b0391909116906399d055c890602401602060405180830381865afa158015611592573d5f803e3d5ffd5b505050506040513d601f19601f820116820180604052508101906115b69190611df4565b90505f80826001600160a01b0316635a1239c1886040518263ffffffff1660e01b81526004016115e891815260200190565b5f60405180830381865afa158015611602573d5f803e3d5ffd5b505050506040513d5f823e601f3d908101601f19168201604052611629919081019061203a565b50509550955050505050816001600160a01b0316866001600160a01b03161461166557604051636aa52cb560e01b815260040160405180910390fd5b604051636450073d60e11b8152600481018290525f906001600160a01b0385169063c8a00e7a906024015f60405180830381865afa1580156116a9573d5f803e3d5ffd5b505050506040513d5f823e601f3d908101601f191682016040526116d091908101906120f9565b9d9c50505050505050505050505050565b6001600160a01b0382165f90815260cb60205260408120549061170483836113cf565b9050805f036117135750505050565b6001600160a01b0384165f90815260cb60205260408120805483929061173a908490611dd1565b909155505060c95460408051630db055fd60e21b815290516001600160a01b03909216916336c157f4916004808201926020929091908290030181865afa158015611787573d5f803e3d5ffd5b505050506040513d601f19601f820116820180604052508101906117ab9190611df4565b6001600160a01b0316637a78e12d826040518263ffffffff1660e01b81526004016117d891815260200190565b5f604051808303815f87803b1580156117ef575f80fd5b505af1158015611801573d5f803e3d5ffd5b5050505060c95f9054906101000a90046001600160a01b03166001600160a01b03166336c157f46040518163ffffffff1660e01b8152600401602060405180830381865afa158015611855573d5f803e3d5ffd5b505050506040513d601f19601f820116820180604052508101906118799190611df4565b6001600160a01b0316846001600160a01b03167fe4a6f5b1a676a94b2af7a506093c172e23d46a5bea6f2d783d4d32c9047800f4836040516118bd91815260200190565b60405180910390a350505050565b6118d58282610994565b61062f576118e281611924565b6118ed836020611936565b6040516020016118fe92919061218e565b60408051601f198184030181529082905262461bcd60e51b82526103ca91600401612202565b60606103826001600160a01b03831660145b60605f611944836002611e37565b61194f906002611dbe565b67ffffffffffffffff81111561196757611967611c96565b6040519080825280601f01601f191660200182016040528015611991576020820181803683370190505b509050600360fc1b815f815181106119ab576119ab612214565b60200101906001600160f81b03191690815f1a905350600f60fb1b816001815181106119d9576119d9612214565b60200101906001600160f81b03191690815f1a9053505f6119fb846002611e37565b611a06906001611dbe565b90505b6001811115611a7d576f181899199a1a9b1b9c1cb0b131b232b360811b85600f1660108110611a3a57611a3a612214565b1a60f81b828281518110611a5057611a50612214565b60200101906001600160f81b03191690815f1a90535060049490941c93611a7681612228565b9050611a09565b5083156105ae5760405162461bcd60e51b815260206004820181905260248201527f537472696e67733a20686578206c656e67746820696e73756666696369656e7460448201526064016103ca565b5f60208284031215611adc575f80fd5b81356001600160e01b0319811681146105ae575f80fd5b5f60208284031215611b03575f80fd5b5035919050565b6001600160a01b0381168114610fde575f80fd5b5f8060408385031215611b2f575f80fd5b823591506020830135611b4181611b0a565b809150509250929050565b60ff81168114610fde575f80fd5b5f805f60608486031215611b6c575f80fd5b8335611b7781611b0a565b92506020840135611b8781611b4c565b929592945050506040919091013590565b5f8060408385031215611ba9575f80fd5b8235611bb481611b4c565b946020939093013593505050565b5f60208284031215611bd2575f80fd5b81356105ae81611b0a565b5f8060408385031215611bee575f80fd5b823591506020830135611b4181611b4c565b5f60208284031215611c10575f80fd5b81356105ae81611b4c565b5f5b83811015611c35578181015183820152602001611c1d565b50505f910152565b5f8151808452611c54816020860160208601611c1b565b601f01601f19169290920160200192915050565b848152836020820152826040820152608060608201525f611c8c6080830184611c3d565b9695505050505050565b634e487b7160e01b5f52604160045260245ffd5b604051601f8201601f1916810167ffffffffffffffff81118282101715611cd357611cd3611c96565b604052919050565b5f67ffffffffffffffff821115611cf457611cf4611c96565b50601f01601f191660200190565b5f805f805f60a08688031215611d16575f80fd5b8535611d2181611b4c565b9450602086013593506040860135925060608601359150608086013567ffffffffffffffff811115611d51575f80fd5b8601601f81018813611d61575f80fd5b8035611d74611d6f82611cdb565b611caa565b818152896020838501011115611d88575f80fd5b816020840160208301375f602083830101528093505050509295509295909350565b634e487b7160e01b5f52601160045260245ffd5b8082018082111561038257610382611daa565b8181038181111561038257610382611daa565b8051611def81611b0a565b919050565b5f60208284031215611e04575f80fd5b81516105ae81611b0a565b80518015158114611def575f80fd5b5f60208284031215611e2e575f80fd5b6105ae82611e0f565b808202811582820484141761038257610382611daa565b600181811c90821680611e6257607f821691505b602082108103611e8057634e487b7160e01b5f52602260045260245ffd5b50919050565b5f60208284031215611e96575f80fd5b5051919050565b5f82611eb757634e487b7160e01b5f52601260045260245ffd5b500490565b601f821115610566575f81815260208120601f850160051c81016020861015611ee25750805b601f850160051c820191505b81811015611f0157828155600101611eee565b505050505050565b815167ffffffffffffffff811115611f2357611f23611c96565b611f3781611f318454611e4e565b84611ebc565b602080601f831160018114611f6a575f8415611f535750858301515b5f19600386901b1c1916600185901b178555611f01565b5f85815260208120601f198616915b82811015611f9857888601518255948401946001909101908401611f79565b5085821015611fb557878501515f19600388901b60f8161c191681555b5050505050600190811b01905550565b5f60208284031215611fd5575f80fd5b81516105ae81611b4c565b805160068110611def575f80fd5b5f611ffb611d6f84611cdb565b905082815283838301111561200e575f80fd5b6105ae836020830184611c1b565b5f82601f83011261202b575f80fd5b6105ae83835160208501611fee565b5f805f805f805f80610100898b031215612052575f80fd5b61205b89611fe0565b9750602089015167ffffffffffffffff80821115612077575f80fd5b6120838c838d0161201c565b985060408b0151915080821115612098575f80fd5b6120a48c838d0161201c565b975060608b01519150808211156120b9575f80fd5b506120c68b828c0161201c565b9550506120d560808a01611de4565b935060a0890151925060c0890151915060e089015190509295985092959890939650565b5f805f805f60a0868803121561210d575f80fd5b61211686611e0f565b945061212460208701611e0f565b9350604086015167ffffffffffffffff81111561213f575f80fd5b8601601f8101881361214f575f80fd5b61215e88825160208401611fee565b935050606086015161216f81611b0a565b608087015190925061218081611b0a565b809150509295509295909350565b7f416363657373436f6e74726f6c3a206163636f756e742000000000000000000081525f83516121c5816017850160208801611c1b565b7001034b99036b4b9b9b4b733903937b6329607d1b60179184019182015283516121f6816028840160208801611c1b565b01602801949350505050565b602081525f6105ae6020830184611c3d565b634e487b7160e01b5f52603260045260245ffd5b5f8161223657612236611daa565b505f19019056fea26469706673582212205882f62f0d1caf55b0022380d57ddb99d840014f537ebbdc49f5b7a0e1b9fc7164736f6c63430008140033496e697469616c697a61626c653a20636f6e7472616374206973206e6f742069",
}

// SDCollateralABI is the input ABI used to generate the binding from.
// Deprecated: Use SDCollateralMetaData.ABI instead.
var SDCollateralABI = SDCollateralMetaData.ABI

// SDCollateralBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use SDCollateralMetaData.Bin instead.
var SDCollateralBin = SDCollateralMetaData.Bin

// DeploySDCollateral deploys a new Ethereum contract, binding an instance of SDCollateral to it.
func DeploySDCollateral(auth *bind.TransactOpts, backend bind.ContractBackend, _admin common.Address, _staderConfig common.Address) (common.Address, *types.Transaction, *SDCollateral, error) {
	parsed, err := SDCollateralMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(SDCollateralBin), backend, _admin, _staderConfig)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &SDCollateral{SDCollateralCaller: SDCollateralCaller{contract: contract}, SDCollateralTransactor: SDCollateralTransactor{contract: contract}, SDCollateralFilterer: SDCollateralFilterer{contract: contract}}, nil
}

// SDCollateral is an auto generated Go binding around an Ethereum contract.
type SDCollateral struct {
	SDCollateralCaller     // Read-only binding to the contract
	SDCollateralTransactor // Write-only binding to the contract
	SDCollateralFilterer   // Log filterer for contract events
}

// SDCollateralCaller is an auto generated read-only Go binding around an Ethereum contract.
type SDCollateralCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SDCollateralTransactor is an auto generated write-only Go binding around an Ethereum contract.
type SDCollateralTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SDCollateralFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type SDCollateralFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SDCollateralSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type SDCollateralSession struct {
	Contract     *SDCollateral     // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// SDCollateralCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type SDCollateralCallerSession struct {
	Contract *SDCollateralCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts       // Call options to use throughout this session
}

// SDCollateralTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type SDCollateralTransactorSession struct {
	Contract     *SDCollateralTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts       // Transaction auth options to use throughout this session
}

// SDCollateralRaw is an auto generated low-level Go binding around an Ethereum contract.
type SDCollateralRaw struct {
	Contract *SDCollateral // Generic contract binding to access the raw methods on
}

// SDCollateralCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type SDCollateralCallerRaw struct {
	Contract *SDCollateralCaller // Generic read-only contract binding to access the raw methods on
}

// SDCollateralTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type SDCollateralTransactorRaw struct {
	Contract *SDCollateralTransactor // Generic write-only contract binding to access the raw methods on
}

// NewSDCollateral creates a new instance of SDCollateral, bound to a specific deployed contract.
func NewSDCollateral(address common.Address, backend bind.ContractBackend) (*SDCollateral, error) {
	contract, err := bindSDCollateral(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &SDCollateral{SDCollateralCaller: SDCollateralCaller{contract: contract}, SDCollateralTransactor: SDCollateralTransactor{contract: contract}, SDCollateralFilterer: SDCollateralFilterer{contract: contract}}, nil
}

// NewSDCollateralCaller creates a new read-only instance of SDCollateral, bound to a specific deployed contract.
func NewSDCollateralCaller(address common.Address, caller bind.ContractCaller) (*SDCollateralCaller, error) {
	contract, err := bindSDCollateral(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &SDCollateralCaller{contract: contract}, nil
}

// NewSDCollateralTransactor creates a new write-only instance of SDCollateral, bound to a specific deployed contract.
func NewSDCollateralTransactor(address common.Address, transactor bind.ContractTransactor) (*SDCollateralTransactor, error) {
	contract, err := bindSDCollateral(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &SDCollateralTransactor{contract: contract}, nil
}

// NewSDCollateralFilterer creates a new log filterer instance of SDCollateral, bound to a specific deployed contract.
func NewSDCollateralFilterer(address common.Address, filterer bind.ContractFilterer) (*SDCollateralFilterer, error) {
	contract, err := bindSDCollateral(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &SDCollateralFilterer{contract: contract}, nil
}

// bindSDCollateral binds a generic wrapper to an already deployed contract.
func bindSDCollateral(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := SDCollateralMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SDCollateral *SDCollateralRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _SDCollateral.Contract.SDCollateralCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SDCollateral *SDCollateralRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SDCollateral.Contract.SDCollateralTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SDCollateral *SDCollateralRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SDCollateral.Contract.SDCollateralTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SDCollateral *SDCollateralCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _SDCollateral.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SDCollateral *SDCollateralTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SDCollateral.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SDCollateral *SDCollateralTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SDCollateral.Contract.contract.Transact(opts, method, params...)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_SDCollateral *SDCollateralCaller) DEFAULTADMINROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _SDCollateral.contract.Call(opts, &out, "DEFAULT_ADMIN_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_SDCollateral *SDCollateralSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _SDCollateral.Contract.DEFAULTADMINROLE(&_SDCollateral.CallOpts)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_SDCollateral *SDCollateralCallerSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _SDCollateral.Contract.DEFAULTADMINROLE(&_SDCollateral.CallOpts)
}

// ConvertETHToSD is a free data retrieval call binding the contract method 0xe614e17c.
//
// Solidity: function convertETHToSD(uint256 _ethAmount) view returns(uint256)
func (_SDCollateral *SDCollateralCaller) ConvertETHToSD(opts *bind.CallOpts, _ethAmount *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _SDCollateral.contract.Call(opts, &out, "convertETHToSD", _ethAmount)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ConvertETHToSD is a free data retrieval call binding the contract method 0xe614e17c.
//
// Solidity: function convertETHToSD(uint256 _ethAmount) view returns(uint256)
func (_SDCollateral *SDCollateralSession) ConvertETHToSD(_ethAmount *big.Int) (*big.Int, error) {
	return _SDCollateral.Contract.ConvertETHToSD(&_SDCollateral.CallOpts, _ethAmount)
}

// ConvertETHToSD is a free data retrieval call binding the contract method 0xe614e17c.
//
// Solidity: function convertETHToSD(uint256 _ethAmount) view returns(uint256)
func (_SDCollateral *SDCollateralCallerSession) ConvertETHToSD(_ethAmount *big.Int) (*big.Int, error) {
	return _SDCollateral.Contract.ConvertETHToSD(&_SDCollateral.CallOpts, _ethAmount)
}

// ConvertSDToETH is a free data retrieval call binding the contract method 0xdfdafccb.
//
// Solidity: function convertSDToETH(uint256 _sdAmount) view returns(uint256)
func (_SDCollateral *SDCollateralCaller) ConvertSDToETH(opts *bind.CallOpts, _sdAmount *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _SDCollateral.contract.Call(opts, &out, "convertSDToETH", _sdAmount)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ConvertSDToETH is a free data retrieval call binding the contract method 0xdfdafccb.
//
// Solidity: function convertSDToETH(uint256 _sdAmount) view returns(uint256)
func (_SDCollateral *SDCollateralSession) ConvertSDToETH(_sdAmount *big.Int) (*big.Int, error) {
	return _SDCollateral.Contract.ConvertSDToETH(&_SDCollateral.CallOpts, _sdAmount)
}

// ConvertSDToETH is a free data retrieval call binding the contract method 0xdfdafccb.
//
// Solidity: function convertSDToETH(uint256 _sdAmount) view returns(uint256)
func (_SDCollateral *SDCollateralCallerSession) ConvertSDToETH(_sdAmount *big.Int) (*big.Int, error) {
	return _SDCollateral.Contract.ConvertSDToETH(&_SDCollateral.CallOpts, _sdAmount)
}

// GetMinimumSDToBond is a free data retrieval call binding the contract method 0x379b727e.
//
// Solidity: function getMinimumSDToBond(uint8 _poolId, uint256 _numValidator) view returns(uint256 _minSDToBond)
func (_SDCollateral *SDCollateralCaller) GetMinimumSDToBond(opts *bind.CallOpts, _poolId uint8, _numValidator *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _SDCollateral.contract.Call(opts, &out, "getMinimumSDToBond", _poolId, _numValidator)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetMinimumSDToBond is a free data retrieval call binding the contract method 0x379b727e.
//
// Solidity: function getMinimumSDToBond(uint8 _poolId, uint256 _numValidator) view returns(uint256 _minSDToBond)
func (_SDCollateral *SDCollateralSession) GetMinimumSDToBond(_poolId uint8, _numValidator *big.Int) (*big.Int, error) {
	return _SDCollateral.Contract.GetMinimumSDToBond(&_SDCollateral.CallOpts, _poolId, _numValidator)
}

// GetMinimumSDToBond is a free data retrieval call binding the contract method 0x379b727e.
//
// Solidity: function getMinimumSDToBond(uint8 _poolId, uint256 _numValidator) view returns(uint256 _minSDToBond)
func (_SDCollateral *SDCollateralCallerSession) GetMinimumSDToBond(_poolId uint8, _numValidator *big.Int) (*big.Int, error) {
	return _SDCollateral.Contract.GetMinimumSDToBond(&_SDCollateral.CallOpts, _poolId, _numValidator)
}

// GetOperatorWithdrawThreshold is a free data retrieval call binding the contract method 0x9871a30a.
//
// Solidity: function getOperatorWithdrawThreshold(address _operator) view returns(uint256 operatorWithdrawThreshold)
func (_SDCollateral *SDCollateralCaller) GetOperatorWithdrawThreshold(opts *bind.CallOpts, _operator common.Address) (*big.Int, error) {
	var out []interface{}
	err := _SDCollateral.contract.Call(opts, &out, "getOperatorWithdrawThreshold", _operator)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetOperatorWithdrawThreshold is a free data retrieval call binding the contract method 0x9871a30a.
//
// Solidity: function getOperatorWithdrawThreshold(address _operator) view returns(uint256 operatorWithdrawThreshold)
func (_SDCollateral *SDCollateralSession) GetOperatorWithdrawThreshold(_operator common.Address) (*big.Int, error) {
	return _SDCollateral.Contract.GetOperatorWithdrawThreshold(&_SDCollateral.CallOpts, _operator)
}

// GetOperatorWithdrawThreshold is a free data retrieval call binding the contract method 0x9871a30a.
//
// Solidity: function getOperatorWithdrawThreshold(address _operator) view returns(uint256 operatorWithdrawThreshold)
func (_SDCollateral *SDCollateralCallerSession) GetOperatorWithdrawThreshold(_operator common.Address) (*big.Int, error) {
	return _SDCollateral.Contract.GetOperatorWithdrawThreshold(&_SDCollateral.CallOpts, _operator)
}

// GetRemainingSDToBond is a free data retrieval call binding the contract method 0x351691ab.
//
// Solidity: function getRemainingSDToBond(address _operator, uint8 _poolId, uint256 _numValidator) view returns(uint256)
func (_SDCollateral *SDCollateralCaller) GetRemainingSDToBond(opts *bind.CallOpts, _operator common.Address, _poolId uint8, _numValidator *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _SDCollateral.contract.Call(opts, &out, "getRemainingSDToBond", _operator, _poolId, _numValidator)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetRemainingSDToBond is a free data retrieval call binding the contract method 0x351691ab.
//
// Solidity: function getRemainingSDToBond(address _operator, uint8 _poolId, uint256 _numValidator) view returns(uint256)
func (_SDCollateral *SDCollateralSession) GetRemainingSDToBond(_operator common.Address, _poolId uint8, _numValidator *big.Int) (*big.Int, error) {
	return _SDCollateral.Contract.GetRemainingSDToBond(&_SDCollateral.CallOpts, _operator, _poolId, _numValidator)
}

// GetRemainingSDToBond is a free data retrieval call binding the contract method 0x351691ab.
//
// Solidity: function getRemainingSDToBond(address _operator, uint8 _poolId, uint256 _numValidator) view returns(uint256)
func (_SDCollateral *SDCollateralCallerSession) GetRemainingSDToBond(_operator common.Address, _poolId uint8, _numValidator *big.Int) (*big.Int, error) {
	return _SDCollateral.Contract.GetRemainingSDToBond(&_SDCollateral.CallOpts, _operator, _poolId, _numValidator)
}

// GetRewardEligibleSD is a free data retrieval call binding the contract method 0x3909afd3.
//
// Solidity: function getRewardEligibleSD(address _operator) view returns(uint256 _rewardEligibleSD)
func (_SDCollateral *SDCollateralCaller) GetRewardEligibleSD(opts *bind.CallOpts, _operator common.Address) (*big.Int, error) {
	var out []interface{}
	err := _SDCollateral.contract.Call(opts, &out, "getRewardEligibleSD", _operator)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetRewardEligibleSD is a free data retrieval call binding the contract method 0x3909afd3.
//
// Solidity: function getRewardEligibleSD(address _operator) view returns(uint256 _rewardEligibleSD)
func (_SDCollateral *SDCollateralSession) GetRewardEligibleSD(_operator common.Address) (*big.Int, error) {
	return _SDCollateral.Contract.GetRewardEligibleSD(&_SDCollateral.CallOpts, _operator)
}

// GetRewardEligibleSD is a free data retrieval call binding the contract method 0x3909afd3.
//
// Solidity: function getRewardEligibleSD(address _operator) view returns(uint256 _rewardEligibleSD)
func (_SDCollateral *SDCollateralCallerSession) GetRewardEligibleSD(_operator common.Address) (*big.Int, error) {
	return _SDCollateral.Contract.GetRewardEligibleSD(&_SDCollateral.CallOpts, _operator)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_SDCollateral *SDCollateralCaller) GetRoleAdmin(opts *bind.CallOpts, role [32]byte) ([32]byte, error) {
	var out []interface{}
	err := _SDCollateral.contract.Call(opts, &out, "getRoleAdmin", role)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_SDCollateral *SDCollateralSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _SDCollateral.Contract.GetRoleAdmin(&_SDCollateral.CallOpts, role)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_SDCollateral *SDCollateralCallerSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _SDCollateral.Contract.GetRoleAdmin(&_SDCollateral.CallOpts, role)
}

// HasEnoughSDCollateral is a free data retrieval call binding the contract method 0xb178e38e.
//
// Solidity: function hasEnoughSDCollateral(address _operator, uint8 _poolId, uint256 _numValidator) view returns(bool)
func (_SDCollateral *SDCollateralCaller) HasEnoughSDCollateral(opts *bind.CallOpts, _operator common.Address, _poolId uint8, _numValidator *big.Int) (bool, error) {
	var out []interface{}
	err := _SDCollateral.contract.Call(opts, &out, "hasEnoughSDCollateral", _operator, _poolId, _numValidator)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// HasEnoughSDCollateral is a free data retrieval call binding the contract method 0xb178e38e.
//
// Solidity: function hasEnoughSDCollateral(address _operator, uint8 _poolId, uint256 _numValidator) view returns(bool)
func (_SDCollateral *SDCollateralSession) HasEnoughSDCollateral(_operator common.Address, _poolId uint8, _numValidator *big.Int) (bool, error) {
	return _SDCollateral.Contract.HasEnoughSDCollateral(&_SDCollateral.CallOpts, _operator, _poolId, _numValidator)
}

// HasEnoughSDCollateral is a free data retrieval call binding the contract method 0xb178e38e.
//
// Solidity: function hasEnoughSDCollateral(address _operator, uint8 _poolId, uint256 _numValidator) view returns(bool)
func (_SDCollateral *SDCollateralCallerSession) HasEnoughSDCollateral(_operator common.Address, _poolId uint8, _numValidator *big.Int) (bool, error) {
	return _SDCollateral.Contract.HasEnoughSDCollateral(&_SDCollateral.CallOpts, _operator, _poolId, _numValidator)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_SDCollateral *SDCollateralCaller) HasRole(opts *bind.CallOpts, role [32]byte, account common.Address) (bool, error) {
	var out []interface{}
	err := _SDCollateral.contract.Call(opts, &out, "hasRole", role, account)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_SDCollateral *SDCollateralSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _SDCollateral.Contract.HasRole(&_SDCollateral.CallOpts, role, account)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_SDCollateral *SDCollateralCallerSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _SDCollateral.Contract.HasRole(&_SDCollateral.CallOpts, role, account)
}

// OperatorSDBalance is a free data retrieval call binding the contract method 0xf9af40b8.
//
// Solidity: function operatorSDBalance(address ) view returns(uint256)
func (_SDCollateral *SDCollateralCaller) OperatorSDBalance(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _SDCollateral.contract.Call(opts, &out, "operatorSDBalance", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// OperatorSDBalance is a free data retrieval call binding the contract method 0xf9af40b8.
//
// Solidity: function operatorSDBalance(address ) view returns(uint256)
func (_SDCollateral *SDCollateralSession) OperatorSDBalance(arg0 common.Address) (*big.Int, error) {
	return _SDCollateral.Contract.OperatorSDBalance(&_SDCollateral.CallOpts, arg0)
}

// OperatorSDBalance is a free data retrieval call binding the contract method 0xf9af40b8.
//
// Solidity: function operatorSDBalance(address ) view returns(uint256)
func (_SDCollateral *SDCollateralCallerSession) OperatorSDBalance(arg0 common.Address) (*big.Int, error) {
	return _SDCollateral.Contract.OperatorSDBalance(&_SDCollateral.CallOpts, arg0)
}

// PoolThresholdbyPoolId is a free data retrieval call binding the contract method 0x8a9b3738.
//
// Solidity: function poolThresholdbyPoolId(uint8 ) view returns(uint256 minThreshold, uint256 maxThreshold, uint256 withdrawThreshold, string units)
func (_SDCollateral *SDCollateralCaller) PoolThresholdbyPoolId(opts *bind.CallOpts, arg0 uint8) (struct {
	MinThreshold      *big.Int
	MaxThreshold      *big.Int
	WithdrawThreshold *big.Int
	Units             string
}, error) {
	var out []interface{}
	err := _SDCollateral.contract.Call(opts, &out, "poolThresholdbyPoolId", arg0)

	outstruct := new(struct {
		MinThreshold      *big.Int
		MaxThreshold      *big.Int
		WithdrawThreshold *big.Int
		Units             string
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.MinThreshold = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.MaxThreshold = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.WithdrawThreshold = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	outstruct.Units = *abi.ConvertType(out[3], new(string)).(*string)

	return *outstruct, err

}

// PoolThresholdbyPoolId is a free data retrieval call binding the contract method 0x8a9b3738.
//
// Solidity: function poolThresholdbyPoolId(uint8 ) view returns(uint256 minThreshold, uint256 maxThreshold, uint256 withdrawThreshold, string units)
func (_SDCollateral *SDCollateralSession) PoolThresholdbyPoolId(arg0 uint8) (struct {
	MinThreshold      *big.Int
	MaxThreshold      *big.Int
	WithdrawThreshold *big.Int
	Units             string
}, error) {
	return _SDCollateral.Contract.PoolThresholdbyPoolId(&_SDCollateral.CallOpts, arg0)
}

// PoolThresholdbyPoolId is a free data retrieval call binding the contract method 0x8a9b3738.
//
// Solidity: function poolThresholdbyPoolId(uint8 ) view returns(uint256 minThreshold, uint256 maxThreshold, uint256 withdrawThreshold, string units)
func (_SDCollateral *SDCollateralCallerSession) PoolThresholdbyPoolId(arg0 uint8) (struct {
	MinThreshold      *big.Int
	MaxThreshold      *big.Int
	WithdrawThreshold *big.Int
	Units             string
}, error) {
	return _SDCollateral.Contract.PoolThresholdbyPoolId(&_SDCollateral.CallOpts, arg0)
}

// StaderConfig is a free data retrieval call binding the contract method 0x490ffa35.
//
// Solidity: function staderConfig() view returns(address)
func (_SDCollateral *SDCollateralCaller) StaderConfig(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _SDCollateral.contract.Call(opts, &out, "staderConfig")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// StaderConfig is a free data retrieval call binding the contract method 0x490ffa35.
//
// Solidity: function staderConfig() view returns(address)
func (_SDCollateral *SDCollateralSession) StaderConfig() (common.Address, error) {
	return _SDCollateral.Contract.StaderConfig(&_SDCollateral.CallOpts)
}

// StaderConfig is a free data retrieval call binding the contract method 0x490ffa35.
//
// Solidity: function staderConfig() view returns(address)
func (_SDCollateral *SDCollateralCallerSession) StaderConfig() (common.Address, error) {
	return _SDCollateral.Contract.StaderConfig(&_SDCollateral.CallOpts)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_SDCollateral *SDCollateralCaller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _SDCollateral.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_SDCollateral *SDCollateralSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _SDCollateral.Contract.SupportsInterface(&_SDCollateral.CallOpts, interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_SDCollateral *SDCollateralCallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _SDCollateral.Contract.SupportsInterface(&_SDCollateral.CallOpts, interfaceId)
}

// DepositSDAsCollateral is a paid mutator transaction binding the contract method 0xfcb7e032.
//
// Solidity: function depositSDAsCollateral(uint256 _sdAmount) returns()
func (_SDCollateral *SDCollateralTransactor) DepositSDAsCollateral(opts *bind.TransactOpts, _sdAmount *big.Int) (*types.Transaction, error) {
	return _SDCollateral.contract.Transact(opts, "depositSDAsCollateral", _sdAmount)
}

// DepositSDAsCollateral is a paid mutator transaction binding the contract method 0xfcb7e032.
//
// Solidity: function depositSDAsCollateral(uint256 _sdAmount) returns()
func (_SDCollateral *SDCollateralSession) DepositSDAsCollateral(_sdAmount *big.Int) (*types.Transaction, error) {
	return _SDCollateral.Contract.DepositSDAsCollateral(&_SDCollateral.TransactOpts, _sdAmount)
}

// DepositSDAsCollateral is a paid mutator transaction binding the contract method 0xfcb7e032.
//
// Solidity: function depositSDAsCollateral(uint256 _sdAmount) returns()
func (_SDCollateral *SDCollateralTransactorSession) DepositSDAsCollateral(_sdAmount *big.Int) (*types.Transaction, error) {
	return _SDCollateral.Contract.DepositSDAsCollateral(&_SDCollateral.TransactOpts, _sdAmount)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_SDCollateral *SDCollateralTransactor) GrantRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _SDCollateral.contract.Transact(opts, "grantRole", role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_SDCollateral *SDCollateralSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _SDCollateral.Contract.GrantRole(&_SDCollateral.TransactOpts, role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_SDCollateral *SDCollateralTransactorSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _SDCollateral.Contract.GrantRole(&_SDCollateral.TransactOpts, role, account)
}

// MaxApproveSD is a paid mutator transaction binding the contract method 0x3e04cd35.
//
// Solidity: function maxApproveSD() returns()
func (_SDCollateral *SDCollateralTransactor) MaxApproveSD(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SDCollateral.contract.Transact(opts, "maxApproveSD")
}

// MaxApproveSD is a paid mutator transaction binding the contract method 0x3e04cd35.
//
// Solidity: function maxApproveSD() returns()
func (_SDCollateral *SDCollateralSession) MaxApproveSD() (*types.Transaction, error) {
	return _SDCollateral.Contract.MaxApproveSD(&_SDCollateral.TransactOpts)
}

// MaxApproveSD is a paid mutator transaction binding the contract method 0x3e04cd35.
//
// Solidity: function maxApproveSD() returns()
func (_SDCollateral *SDCollateralTransactorSession) MaxApproveSD() (*types.Transaction, error) {
	return _SDCollateral.Contract.MaxApproveSD(&_SDCollateral.TransactOpts)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address account) returns()
func (_SDCollateral *SDCollateralTransactor) RenounceRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _SDCollateral.contract.Transact(opts, "renounceRole", role, account)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address account) returns()
func (_SDCollateral *SDCollateralSession) RenounceRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _SDCollateral.Contract.RenounceRole(&_SDCollateral.TransactOpts, role, account)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address account) returns()
func (_SDCollateral *SDCollateralTransactorSession) RenounceRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _SDCollateral.Contract.RenounceRole(&_SDCollateral.TransactOpts, role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_SDCollateral *SDCollateralTransactor) RevokeRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _SDCollateral.contract.Transact(opts, "revokeRole", role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_SDCollateral *SDCollateralSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _SDCollateral.Contract.RevokeRole(&_SDCollateral.TransactOpts, role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_SDCollateral *SDCollateralTransactorSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _SDCollateral.Contract.RevokeRole(&_SDCollateral.TransactOpts, role, account)
}

// SlashValidatorSD is a paid mutator transaction binding the contract method 0x4c538f58.
//
// Solidity: function slashValidatorSD(uint256 _validatorId, uint8 _poolId) returns()
func (_SDCollateral *SDCollateralTransactor) SlashValidatorSD(opts *bind.TransactOpts, _validatorId *big.Int, _poolId uint8) (*types.Transaction, error) {
	return _SDCollateral.contract.Transact(opts, "slashValidatorSD", _validatorId, _poolId)
}

// SlashValidatorSD is a paid mutator transaction binding the contract method 0x4c538f58.
//
// Solidity: function slashValidatorSD(uint256 _validatorId, uint8 _poolId) returns()
func (_SDCollateral *SDCollateralSession) SlashValidatorSD(_validatorId *big.Int, _poolId uint8) (*types.Transaction, error) {
	return _SDCollateral.Contract.SlashValidatorSD(&_SDCollateral.TransactOpts, _validatorId, _poolId)
}

// SlashValidatorSD is a paid mutator transaction binding the contract method 0x4c538f58.
//
// Solidity: function slashValidatorSD(uint256 _validatorId, uint8 _poolId) returns()
func (_SDCollateral *SDCollateralTransactorSession) SlashValidatorSD(_validatorId *big.Int, _poolId uint8) (*types.Transaction, error) {
	return _SDCollateral.Contract.SlashValidatorSD(&_SDCollateral.TransactOpts, _validatorId, _poolId)
}

// UpdatePoolThreshold is a paid mutator transaction binding the contract method 0xe0412f0e.
//
// Solidity: function updatePoolThreshold(uint8 _poolId, uint256 _minThreshold, uint256 _maxThreshold, uint256 _withdrawThreshold, string _units) returns()
func (_SDCollateral *SDCollateralTransactor) UpdatePoolThreshold(opts *bind.TransactOpts, _poolId uint8, _minThreshold *big.Int, _maxThreshold *big.Int, _withdrawThreshold *big.Int, _units string) (*types.Transaction, error) {
	return _SDCollateral.contract.Transact(opts, "updatePoolThreshold", _poolId, _minThreshold, _maxThreshold, _withdrawThreshold, _units)
}

// UpdatePoolThreshold is a paid mutator transaction binding the contract method 0xe0412f0e.
//
// Solidity: function updatePoolThreshold(uint8 _poolId, uint256 _minThreshold, uint256 _maxThreshold, uint256 _withdrawThreshold, string _units) returns()
func (_SDCollateral *SDCollateralSession) UpdatePoolThreshold(_poolId uint8, _minThreshold *big.Int, _maxThreshold *big.Int, _withdrawThreshold *big.Int, _units string) (*types.Transaction, error) {
	return _SDCollateral.Contract.UpdatePoolThreshold(&_SDCollateral.TransactOpts, _poolId, _minThreshold, _maxThreshold, _withdrawThreshold, _units)
}

// UpdatePoolThreshold is a paid mutator transaction binding the contract method 0xe0412f0e.
//
// Solidity: function updatePoolThreshold(uint8 _poolId, uint256 _minThreshold, uint256 _maxThreshold, uint256 _withdrawThreshold, string _units) returns()
func (_SDCollateral *SDCollateralTransactorSession) UpdatePoolThreshold(_poolId uint8, _minThreshold *big.Int, _maxThreshold *big.Int, _withdrawThreshold *big.Int, _units string) (*types.Transaction, error) {
	return _SDCollateral.Contract.UpdatePoolThreshold(&_SDCollateral.TransactOpts, _poolId, _minThreshold, _maxThreshold, _withdrawThreshold, _units)
}

// UpdateStaderConfig is a paid mutator transaction binding the contract method 0x9ee804cb.
//
// Solidity: function updateStaderConfig(address _staderConfig) returns()
func (_SDCollateral *SDCollateralTransactor) UpdateStaderConfig(opts *bind.TransactOpts, _staderConfig common.Address) (*types.Transaction, error) {
	return _SDCollateral.contract.Transact(opts, "updateStaderConfig", _staderConfig)
}

// UpdateStaderConfig is a paid mutator transaction binding the contract method 0x9ee804cb.
//
// Solidity: function updateStaderConfig(address _staderConfig) returns()
func (_SDCollateral *SDCollateralSession) UpdateStaderConfig(_staderConfig common.Address) (*types.Transaction, error) {
	return _SDCollateral.Contract.UpdateStaderConfig(&_SDCollateral.TransactOpts, _staderConfig)
}

// UpdateStaderConfig is a paid mutator transaction binding the contract method 0x9ee804cb.
//
// Solidity: function updateStaderConfig(address _staderConfig) returns()
func (_SDCollateral *SDCollateralTransactorSession) UpdateStaderConfig(_staderConfig common.Address) (*types.Transaction, error) {
	return _SDCollateral.Contract.UpdateStaderConfig(&_SDCollateral.TransactOpts, _staderConfig)
}

// Withdraw is a paid mutator transaction binding the contract method 0x2e1a7d4d.
//
// Solidity: function withdraw(uint256 _requestedSD) returns()
func (_SDCollateral *SDCollateralTransactor) Withdraw(opts *bind.TransactOpts, _requestedSD *big.Int) (*types.Transaction, error) {
	return _SDCollateral.contract.Transact(opts, "withdraw", _requestedSD)
}

// Withdraw is a paid mutator transaction binding the contract method 0x2e1a7d4d.
//
// Solidity: function withdraw(uint256 _requestedSD) returns()
func (_SDCollateral *SDCollateralSession) Withdraw(_requestedSD *big.Int) (*types.Transaction, error) {
	return _SDCollateral.Contract.Withdraw(&_SDCollateral.TransactOpts, _requestedSD)
}

// Withdraw is a paid mutator transaction binding the contract method 0x2e1a7d4d.
//
// Solidity: function withdraw(uint256 _requestedSD) returns()
func (_SDCollateral *SDCollateralTransactorSession) Withdraw(_requestedSD *big.Int) (*types.Transaction, error) {
	return _SDCollateral.Contract.Withdraw(&_SDCollateral.TransactOpts, _requestedSD)
}

// SDCollateralInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the SDCollateral contract.
type SDCollateralInitializedIterator struct {
	Event *SDCollateralInitialized // Event containing the contract specifics and raw log

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
func (it *SDCollateralInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SDCollateralInitialized)
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
		it.Event = new(SDCollateralInitialized)
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
func (it *SDCollateralInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SDCollateralInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SDCollateralInitialized represents a Initialized event raised by the SDCollateral contract.
type SDCollateralInitialized struct {
	Version uint8
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_SDCollateral *SDCollateralFilterer) FilterInitialized(opts *bind.FilterOpts) (*SDCollateralInitializedIterator, error) {

	logs, sub, err := _SDCollateral.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &SDCollateralInitializedIterator{contract: _SDCollateral.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_SDCollateral *SDCollateralFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *SDCollateralInitialized) (event.Subscription, error) {

	logs, sub, err := _SDCollateral.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SDCollateralInitialized)
				if err := _SDCollateral.contract.UnpackLog(event, "Initialized", log); err != nil {
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
func (_SDCollateral *SDCollateralFilterer) ParseInitialized(log types.Log) (*SDCollateralInitialized, error) {
	event := new(SDCollateralInitialized)
	if err := _SDCollateral.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SDCollateralRoleAdminChangedIterator is returned from FilterRoleAdminChanged and is used to iterate over the raw logs and unpacked data for RoleAdminChanged events raised by the SDCollateral contract.
type SDCollateralRoleAdminChangedIterator struct {
	Event *SDCollateralRoleAdminChanged // Event containing the contract specifics and raw log

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
func (it *SDCollateralRoleAdminChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SDCollateralRoleAdminChanged)
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
		it.Event = new(SDCollateralRoleAdminChanged)
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
func (it *SDCollateralRoleAdminChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SDCollateralRoleAdminChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SDCollateralRoleAdminChanged represents a RoleAdminChanged event raised by the SDCollateral contract.
type SDCollateralRoleAdminChanged struct {
	Role              [32]byte
	PreviousAdminRole [32]byte
	NewAdminRole      [32]byte
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterRoleAdminChanged is a free log retrieval operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_SDCollateral *SDCollateralFilterer) FilterRoleAdminChanged(opts *bind.FilterOpts, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (*SDCollateralRoleAdminChangedIterator, error) {

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

	logs, sub, err := _SDCollateral.contract.FilterLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return &SDCollateralRoleAdminChangedIterator{contract: _SDCollateral.contract, event: "RoleAdminChanged", logs: logs, sub: sub}, nil
}

// WatchRoleAdminChanged is a free log subscription operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_SDCollateral *SDCollateralFilterer) WatchRoleAdminChanged(opts *bind.WatchOpts, sink chan<- *SDCollateralRoleAdminChanged, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (event.Subscription, error) {

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

	logs, sub, err := _SDCollateral.contract.WatchLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SDCollateralRoleAdminChanged)
				if err := _SDCollateral.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
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
func (_SDCollateral *SDCollateralFilterer) ParseRoleAdminChanged(log types.Log) (*SDCollateralRoleAdminChanged, error) {
	event := new(SDCollateralRoleAdminChanged)
	if err := _SDCollateral.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SDCollateralRoleGrantedIterator is returned from FilterRoleGranted and is used to iterate over the raw logs and unpacked data for RoleGranted events raised by the SDCollateral contract.
type SDCollateralRoleGrantedIterator struct {
	Event *SDCollateralRoleGranted // Event containing the contract specifics and raw log

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
func (it *SDCollateralRoleGrantedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SDCollateralRoleGranted)
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
		it.Event = new(SDCollateralRoleGranted)
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
func (it *SDCollateralRoleGrantedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SDCollateralRoleGrantedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SDCollateralRoleGranted represents a RoleGranted event raised by the SDCollateral contract.
type SDCollateralRoleGranted struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleGranted is a free log retrieval operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_SDCollateral *SDCollateralFilterer) FilterRoleGranted(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*SDCollateralRoleGrantedIterator, error) {

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

	logs, sub, err := _SDCollateral.contract.FilterLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &SDCollateralRoleGrantedIterator{contract: _SDCollateral.contract, event: "RoleGranted", logs: logs, sub: sub}, nil
}

// WatchRoleGranted is a free log subscription operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_SDCollateral *SDCollateralFilterer) WatchRoleGranted(opts *bind.WatchOpts, sink chan<- *SDCollateralRoleGranted, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _SDCollateral.contract.WatchLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SDCollateralRoleGranted)
				if err := _SDCollateral.contract.UnpackLog(event, "RoleGranted", log); err != nil {
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
func (_SDCollateral *SDCollateralFilterer) ParseRoleGranted(log types.Log) (*SDCollateralRoleGranted, error) {
	event := new(SDCollateralRoleGranted)
	if err := _SDCollateral.contract.UnpackLog(event, "RoleGranted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SDCollateralRoleRevokedIterator is returned from FilterRoleRevoked and is used to iterate over the raw logs and unpacked data for RoleRevoked events raised by the SDCollateral contract.
type SDCollateralRoleRevokedIterator struct {
	Event *SDCollateralRoleRevoked // Event containing the contract specifics and raw log

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
func (it *SDCollateralRoleRevokedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SDCollateralRoleRevoked)
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
		it.Event = new(SDCollateralRoleRevoked)
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
func (it *SDCollateralRoleRevokedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SDCollateralRoleRevokedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SDCollateralRoleRevoked represents a RoleRevoked event raised by the SDCollateral contract.
type SDCollateralRoleRevoked struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleRevoked is a free log retrieval operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_SDCollateral *SDCollateralFilterer) FilterRoleRevoked(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*SDCollateralRoleRevokedIterator, error) {

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

	logs, sub, err := _SDCollateral.contract.FilterLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &SDCollateralRoleRevokedIterator{contract: _SDCollateral.contract, event: "RoleRevoked", logs: logs, sub: sub}, nil
}

// WatchRoleRevoked is a free log subscription operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_SDCollateral *SDCollateralFilterer) WatchRoleRevoked(opts *bind.WatchOpts, sink chan<- *SDCollateralRoleRevoked, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _SDCollateral.contract.WatchLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SDCollateralRoleRevoked)
				if err := _SDCollateral.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
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
func (_SDCollateral *SDCollateralFilterer) ParseRoleRevoked(log types.Log) (*SDCollateralRoleRevoked, error) {
	event := new(SDCollateralRoleRevoked)
	if err := _SDCollateral.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SDCollateralSDDepositedIterator is returned from FilterSDDeposited and is used to iterate over the raw logs and unpacked data for SDDeposited events raised by the SDCollateral contract.
type SDCollateralSDDepositedIterator struct {
	Event *SDCollateralSDDeposited // Event containing the contract specifics and raw log

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
func (it *SDCollateralSDDepositedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SDCollateralSDDeposited)
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
		it.Event = new(SDCollateralSDDeposited)
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
func (it *SDCollateralSDDepositedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SDCollateralSDDepositedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SDCollateralSDDeposited represents a SDDeposited event raised by the SDCollateral contract.
type SDCollateralSDDeposited struct {
	Operator common.Address
	SdAmount *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterSDDeposited is a free log retrieval operation binding the contract event 0x112973aa2e3b182a34447572d830c1e97afc414558a08f33554be8e545224259.
//
// Solidity: event SDDeposited(address indexed operator, uint256 sdAmount)
func (_SDCollateral *SDCollateralFilterer) FilterSDDeposited(opts *bind.FilterOpts, operator []common.Address) (*SDCollateralSDDepositedIterator, error) {

	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _SDCollateral.contract.FilterLogs(opts, "SDDeposited", operatorRule)
	if err != nil {
		return nil, err
	}
	return &SDCollateralSDDepositedIterator{contract: _SDCollateral.contract, event: "SDDeposited", logs: logs, sub: sub}, nil
}

// WatchSDDeposited is a free log subscription operation binding the contract event 0x112973aa2e3b182a34447572d830c1e97afc414558a08f33554be8e545224259.
//
// Solidity: event SDDeposited(address indexed operator, uint256 sdAmount)
func (_SDCollateral *SDCollateralFilterer) WatchSDDeposited(opts *bind.WatchOpts, sink chan<- *SDCollateralSDDeposited, operator []common.Address) (event.Subscription, error) {

	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _SDCollateral.contract.WatchLogs(opts, "SDDeposited", operatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SDCollateralSDDeposited)
				if err := _SDCollateral.contract.UnpackLog(event, "SDDeposited", log); err != nil {
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

// ParseSDDeposited is a log parse operation binding the contract event 0x112973aa2e3b182a34447572d830c1e97afc414558a08f33554be8e545224259.
//
// Solidity: event SDDeposited(address indexed operator, uint256 sdAmount)
func (_SDCollateral *SDCollateralFilterer) ParseSDDeposited(log types.Log) (*SDCollateralSDDeposited, error) {
	event := new(SDCollateralSDDeposited)
	if err := _SDCollateral.contract.UnpackLog(event, "SDDeposited", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SDCollateralSDSlashedIterator is returned from FilterSDSlashed and is used to iterate over the raw logs and unpacked data for SDSlashed events raised by the SDCollateral contract.
type SDCollateralSDSlashedIterator struct {
	Event *SDCollateralSDSlashed // Event containing the contract specifics and raw log

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
func (it *SDCollateralSDSlashedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SDCollateralSDSlashed)
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
		it.Event = new(SDCollateralSDSlashed)
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
func (it *SDCollateralSDSlashedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SDCollateralSDSlashedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SDCollateralSDSlashed represents a SDSlashed event raised by the SDCollateral contract.
type SDCollateralSDSlashed struct {
	Operator  common.Address
	Auction   common.Address
	SdSlashed *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterSDSlashed is a free log retrieval operation binding the contract event 0xe4a6f5b1a676a94b2af7a506093c172e23d46a5bea6f2d783d4d32c9047800f4.
//
// Solidity: event SDSlashed(address indexed operator, address indexed auction, uint256 sdSlashed)
func (_SDCollateral *SDCollateralFilterer) FilterSDSlashed(opts *bind.FilterOpts, operator []common.Address, auction []common.Address) (*SDCollateralSDSlashedIterator, error) {

	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}
	var auctionRule []interface{}
	for _, auctionItem := range auction {
		auctionRule = append(auctionRule, auctionItem)
	}

	logs, sub, err := _SDCollateral.contract.FilterLogs(opts, "SDSlashed", operatorRule, auctionRule)
	if err != nil {
		return nil, err
	}
	return &SDCollateralSDSlashedIterator{contract: _SDCollateral.contract, event: "SDSlashed", logs: logs, sub: sub}, nil
}

// WatchSDSlashed is a free log subscription operation binding the contract event 0xe4a6f5b1a676a94b2af7a506093c172e23d46a5bea6f2d783d4d32c9047800f4.
//
// Solidity: event SDSlashed(address indexed operator, address indexed auction, uint256 sdSlashed)
func (_SDCollateral *SDCollateralFilterer) WatchSDSlashed(opts *bind.WatchOpts, sink chan<- *SDCollateralSDSlashed, operator []common.Address, auction []common.Address) (event.Subscription, error) {

	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}
	var auctionRule []interface{}
	for _, auctionItem := range auction {
		auctionRule = append(auctionRule, auctionItem)
	}

	logs, sub, err := _SDCollateral.contract.WatchLogs(opts, "SDSlashed", operatorRule, auctionRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SDCollateralSDSlashed)
				if err := _SDCollateral.contract.UnpackLog(event, "SDSlashed", log); err != nil {
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

// ParseSDSlashed is a log parse operation binding the contract event 0xe4a6f5b1a676a94b2af7a506093c172e23d46a5bea6f2d783d4d32c9047800f4.
//
// Solidity: event SDSlashed(address indexed operator, address indexed auction, uint256 sdSlashed)
func (_SDCollateral *SDCollateralFilterer) ParseSDSlashed(log types.Log) (*SDCollateralSDSlashed, error) {
	event := new(SDCollateralSDSlashed)
	if err := _SDCollateral.contract.UnpackLog(event, "SDSlashed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SDCollateralSDWithdrawnIterator is returned from FilterSDWithdrawn and is used to iterate over the raw logs and unpacked data for SDWithdrawn events raised by the SDCollateral contract.
type SDCollateralSDWithdrawnIterator struct {
	Event *SDCollateralSDWithdrawn // Event containing the contract specifics and raw log

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
func (it *SDCollateralSDWithdrawnIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SDCollateralSDWithdrawn)
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
		it.Event = new(SDCollateralSDWithdrawn)
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
func (it *SDCollateralSDWithdrawnIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SDCollateralSDWithdrawnIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SDCollateralSDWithdrawn represents a SDWithdrawn event raised by the SDCollateral contract.
type SDCollateralSDWithdrawn struct {
	Operator common.Address
	SdAmount *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterSDWithdrawn is a free log retrieval operation binding the contract event 0x48c1f846fa4bc05385324ee60316f9c6778ed2b5f205a6319678a609b8767607.
//
// Solidity: event SDWithdrawn(address indexed operator, uint256 sdAmount)
func (_SDCollateral *SDCollateralFilterer) FilterSDWithdrawn(opts *bind.FilterOpts, operator []common.Address) (*SDCollateralSDWithdrawnIterator, error) {

	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _SDCollateral.contract.FilterLogs(opts, "SDWithdrawn", operatorRule)
	if err != nil {
		return nil, err
	}
	return &SDCollateralSDWithdrawnIterator{contract: _SDCollateral.contract, event: "SDWithdrawn", logs: logs, sub: sub}, nil
}

// WatchSDWithdrawn is a free log subscription operation binding the contract event 0x48c1f846fa4bc05385324ee60316f9c6778ed2b5f205a6319678a609b8767607.
//
// Solidity: event SDWithdrawn(address indexed operator, uint256 sdAmount)
func (_SDCollateral *SDCollateralFilterer) WatchSDWithdrawn(opts *bind.WatchOpts, sink chan<- *SDCollateralSDWithdrawn, operator []common.Address) (event.Subscription, error) {

	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _SDCollateral.contract.WatchLogs(opts, "SDWithdrawn", operatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SDCollateralSDWithdrawn)
				if err := _SDCollateral.contract.UnpackLog(event, "SDWithdrawn", log); err != nil {
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

// ParseSDWithdrawn is a log parse operation binding the contract event 0x48c1f846fa4bc05385324ee60316f9c6778ed2b5f205a6319678a609b8767607.
//
// Solidity: event SDWithdrawn(address indexed operator, uint256 sdAmount)
func (_SDCollateral *SDCollateralFilterer) ParseSDWithdrawn(log types.Log) (*SDCollateralSDWithdrawn, error) {
	event := new(SDCollateralSDWithdrawn)
	if err := _SDCollateral.contract.UnpackLog(event, "SDWithdrawn", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SDCollateralUpdatedPoolIdForOperatorIterator is returned from FilterUpdatedPoolIdForOperator and is used to iterate over the raw logs and unpacked data for UpdatedPoolIdForOperator events raised by the SDCollateral contract.
type SDCollateralUpdatedPoolIdForOperatorIterator struct {
	Event *SDCollateralUpdatedPoolIdForOperator // Event containing the contract specifics and raw log

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
func (it *SDCollateralUpdatedPoolIdForOperatorIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SDCollateralUpdatedPoolIdForOperator)
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
		it.Event = new(SDCollateralUpdatedPoolIdForOperator)
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
func (it *SDCollateralUpdatedPoolIdForOperatorIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SDCollateralUpdatedPoolIdForOperatorIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SDCollateralUpdatedPoolIdForOperator represents a UpdatedPoolIdForOperator event raised by the SDCollateral contract.
type SDCollateralUpdatedPoolIdForOperator struct {
	PoolId   uint8
	Operator common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterUpdatedPoolIdForOperator is a free log retrieval operation binding the contract event 0x834f00ba6adeb9f7123fa03b8252cdda3f81509cc96c3c2239420138fa1b895e.
//
// Solidity: event UpdatedPoolIdForOperator(uint8 poolId, address operator)
func (_SDCollateral *SDCollateralFilterer) FilterUpdatedPoolIdForOperator(opts *bind.FilterOpts) (*SDCollateralUpdatedPoolIdForOperatorIterator, error) {

	logs, sub, err := _SDCollateral.contract.FilterLogs(opts, "UpdatedPoolIdForOperator")
	if err != nil {
		return nil, err
	}
	return &SDCollateralUpdatedPoolIdForOperatorIterator{contract: _SDCollateral.contract, event: "UpdatedPoolIdForOperator", logs: logs, sub: sub}, nil
}

// WatchUpdatedPoolIdForOperator is a free log subscription operation binding the contract event 0x834f00ba6adeb9f7123fa03b8252cdda3f81509cc96c3c2239420138fa1b895e.
//
// Solidity: event UpdatedPoolIdForOperator(uint8 poolId, address operator)
func (_SDCollateral *SDCollateralFilterer) WatchUpdatedPoolIdForOperator(opts *bind.WatchOpts, sink chan<- *SDCollateralUpdatedPoolIdForOperator) (event.Subscription, error) {

	logs, sub, err := _SDCollateral.contract.WatchLogs(opts, "UpdatedPoolIdForOperator")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SDCollateralUpdatedPoolIdForOperator)
				if err := _SDCollateral.contract.UnpackLog(event, "UpdatedPoolIdForOperator", log); err != nil {
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

// ParseUpdatedPoolIdForOperator is a log parse operation binding the contract event 0x834f00ba6adeb9f7123fa03b8252cdda3f81509cc96c3c2239420138fa1b895e.
//
// Solidity: event UpdatedPoolIdForOperator(uint8 poolId, address operator)
func (_SDCollateral *SDCollateralFilterer) ParseUpdatedPoolIdForOperator(log types.Log) (*SDCollateralUpdatedPoolIdForOperator, error) {
	event := new(SDCollateralUpdatedPoolIdForOperator)
	if err := _SDCollateral.contract.UnpackLog(event, "UpdatedPoolIdForOperator", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SDCollateralUpdatedPoolThresholdIterator is returned from FilterUpdatedPoolThreshold and is used to iterate over the raw logs and unpacked data for UpdatedPoolThreshold events raised by the SDCollateral contract.
type SDCollateralUpdatedPoolThresholdIterator struct {
	Event *SDCollateralUpdatedPoolThreshold // Event containing the contract specifics and raw log

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
func (it *SDCollateralUpdatedPoolThresholdIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SDCollateralUpdatedPoolThreshold)
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
		it.Event = new(SDCollateralUpdatedPoolThreshold)
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
func (it *SDCollateralUpdatedPoolThresholdIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SDCollateralUpdatedPoolThresholdIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SDCollateralUpdatedPoolThreshold represents a UpdatedPoolThreshold event raised by the SDCollateral contract.
type SDCollateralUpdatedPoolThreshold struct {
	PoolId            uint8
	MinThreshold      *big.Int
	WithdrawThreshold *big.Int
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterUpdatedPoolThreshold is a free log retrieval operation binding the contract event 0x18757dd1fbfe2ad823e1bd4de3f8a2ee76b49f92f6aa34cc7cbf717cdf4d1758.
//
// Solidity: event UpdatedPoolThreshold(uint8 poolId, uint256 minThreshold, uint256 withdrawThreshold)
func (_SDCollateral *SDCollateralFilterer) FilterUpdatedPoolThreshold(opts *bind.FilterOpts) (*SDCollateralUpdatedPoolThresholdIterator, error) {

	logs, sub, err := _SDCollateral.contract.FilterLogs(opts, "UpdatedPoolThreshold")
	if err != nil {
		return nil, err
	}
	return &SDCollateralUpdatedPoolThresholdIterator{contract: _SDCollateral.contract, event: "UpdatedPoolThreshold", logs: logs, sub: sub}, nil
}

// WatchUpdatedPoolThreshold is a free log subscription operation binding the contract event 0x18757dd1fbfe2ad823e1bd4de3f8a2ee76b49f92f6aa34cc7cbf717cdf4d1758.
//
// Solidity: event UpdatedPoolThreshold(uint8 poolId, uint256 minThreshold, uint256 withdrawThreshold)
func (_SDCollateral *SDCollateralFilterer) WatchUpdatedPoolThreshold(opts *bind.WatchOpts, sink chan<- *SDCollateralUpdatedPoolThreshold) (event.Subscription, error) {

	logs, sub, err := _SDCollateral.contract.WatchLogs(opts, "UpdatedPoolThreshold")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SDCollateralUpdatedPoolThreshold)
				if err := _SDCollateral.contract.UnpackLog(event, "UpdatedPoolThreshold", log); err != nil {
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

// ParseUpdatedPoolThreshold is a log parse operation binding the contract event 0x18757dd1fbfe2ad823e1bd4de3f8a2ee76b49f92f6aa34cc7cbf717cdf4d1758.
//
// Solidity: event UpdatedPoolThreshold(uint8 poolId, uint256 minThreshold, uint256 withdrawThreshold)
func (_SDCollateral *SDCollateralFilterer) ParseUpdatedPoolThreshold(log types.Log) (*SDCollateralUpdatedPoolThreshold, error) {
	event := new(SDCollateralUpdatedPoolThreshold)
	if err := _SDCollateral.contract.UnpackLog(event, "UpdatedPoolThreshold", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SDCollateralUpdatedStaderConfigIterator is returned from FilterUpdatedStaderConfig and is used to iterate over the raw logs and unpacked data for UpdatedStaderConfig events raised by the SDCollateral contract.
type SDCollateralUpdatedStaderConfigIterator struct {
	Event *SDCollateralUpdatedStaderConfig // Event containing the contract specifics and raw log

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
func (it *SDCollateralUpdatedStaderConfigIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SDCollateralUpdatedStaderConfig)
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
		it.Event = new(SDCollateralUpdatedStaderConfig)
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
func (it *SDCollateralUpdatedStaderConfigIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SDCollateralUpdatedStaderConfigIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SDCollateralUpdatedStaderConfig represents a UpdatedStaderConfig event raised by the SDCollateral contract.
type SDCollateralUpdatedStaderConfig struct {
	StaderConfig common.Address
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterUpdatedStaderConfig is a free log retrieval operation binding the contract event 0xdb2219043d7b197cb235f1af0cf6d782d77dee3de19e3f4fb6d39aae633b4485.
//
// Solidity: event UpdatedStaderConfig(address indexed staderConfig)
func (_SDCollateral *SDCollateralFilterer) FilterUpdatedStaderConfig(opts *bind.FilterOpts, staderConfig []common.Address) (*SDCollateralUpdatedStaderConfigIterator, error) {

	var staderConfigRule []interface{}
	for _, staderConfigItem := range staderConfig {
		staderConfigRule = append(staderConfigRule, staderConfigItem)
	}

	logs, sub, err := _SDCollateral.contract.FilterLogs(opts, "UpdatedStaderConfig", staderConfigRule)
	if err != nil {
		return nil, err
	}
	return &SDCollateralUpdatedStaderConfigIterator{contract: _SDCollateral.contract, event: "UpdatedStaderConfig", logs: logs, sub: sub}, nil
}

// WatchUpdatedStaderConfig is a free log subscription operation binding the contract event 0xdb2219043d7b197cb235f1af0cf6d782d77dee3de19e3f4fb6d39aae633b4485.
//
// Solidity: event UpdatedStaderConfig(address indexed staderConfig)
func (_SDCollateral *SDCollateralFilterer) WatchUpdatedStaderConfig(opts *bind.WatchOpts, sink chan<- *SDCollateralUpdatedStaderConfig, staderConfig []common.Address) (event.Subscription, error) {

	var staderConfigRule []interface{}
	for _, staderConfigItem := range staderConfig {
		staderConfigRule = append(staderConfigRule, staderConfigItem)
	}

	logs, sub, err := _SDCollateral.contract.WatchLogs(opts, "UpdatedStaderConfig", staderConfigRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SDCollateralUpdatedStaderConfig)
				if err := _SDCollateral.contract.UnpackLog(event, "UpdatedStaderConfig", log); err != nil {
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
// Solidity: event UpdatedStaderConfig(address indexed staderConfig)
func (_SDCollateral *SDCollateralFilterer) ParseUpdatedStaderConfig(log types.Log) (*SDCollateralUpdatedStaderConfig, error) {
	event := new(SDCollateralUpdatedStaderConfig)
	if err := _SDCollateral.contract.UnpackLog(event, "UpdatedStaderConfig", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
