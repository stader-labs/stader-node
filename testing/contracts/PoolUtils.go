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

// PoolUtilsMetaData contains all meta data concerning the PoolUtils contract.
var PoolUtilsMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_admin\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_staderConfig\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"CallerNotManager\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"CallerNotOperator\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"EmptyNameString\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ExistingOrMismatchingPoolId\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidLengthOfPubkey\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidLengthOfSignature\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NameCrossedMaxLength\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OperatorIsNotOnboarded\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"PoolIdNotPresent\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"PubkeyAlreadyExist\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"PubkeyDoesNotExit\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ZeroAddress\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint8\",\"name\":\"poolId\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"poolAddress\",\"type\":\"address\"}],\"name\":\"DeactivatedPool\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"pubkey\",\"type\":\"bytes\"}],\"name\":\"ExitValidator\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"version\",\"type\":\"uint8\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint8\",\"name\":\"poolId\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"poolAddress\",\"type\":\"address\"}],\"name\":\"PoolAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint8\",\"name\":\"poolId\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"poolAddress\",\"type\":\"address\"}],\"name\":\"PoolAddressUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"previousAdminRole\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"newAdminRole\",\"type\":\"bytes32\"}],\"name\":\"RoleAdminChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleGranted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleRevoked\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"staderConfig\",\"type\":\"address\"}],\"name\":\"UpdatedStaderConfig\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"DEFAULT_ADMIN_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"_poolId\",\"type\":\"uint8\"},{\"internalType\":\"address\",\"name\":\"_poolAddress\",\"type\":\"address\"}],\"name\":\"addNewPool\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"_poolId\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"_totalRewards\",\"type\":\"uint256\"}],\"name\":\"calculateRewardShare\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"userShare\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"operatorShare\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"protocolShare\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"_poolId\",\"type\":\"uint8\"}],\"name\":\"getActiveValidatorCountByPool\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"_poolId\",\"type\":\"uint8\"}],\"name\":\"getCollateralETH\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"_poolId\",\"type\":\"uint8\"}],\"name\":\"getNodeRegistry\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"_poolId\",\"type\":\"uint8\"}],\"name\":\"getOperatorFee\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_operAddr\",\"type\":\"address\"}],\"name\":\"getOperatorPoolId\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"_poolId\",\"type\":\"uint8\"},{\"internalType\":\"address\",\"name\":\"_nodeOperator\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_startIndex\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_endIndex\",\"type\":\"uint256\"}],\"name\":\"getOperatorTotalNonTerminalKeys\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getPoolIdArray\",\"outputs\":[{\"internalType\":\"uint8[]\",\"name\":\"\",\"type\":\"uint8[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"_poolId\",\"type\":\"uint8\"}],\"name\":\"getProtocolFee\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"_poolId\",\"type\":\"uint8\"}],\"name\":\"getQueuedValidatorCountByPool\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"}],\"name\":\"getRoleAdmin\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"_poolId\",\"type\":\"uint8\"}],\"name\":\"getSocializingPoolAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getTotalActiveValidatorCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"_pubkey\",\"type\":\"bytes\"}],\"name\":\"getValidatorPoolId\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"grantRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"hasRole\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_operAddr\",\"type\":\"address\"}],\"name\":\"isExistingOperator\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"_poolId\",\"type\":\"uint8\"}],\"name\":\"isExistingPoolId\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"_pubkey\",\"type\":\"bytes\"}],\"name\":\"isExistingPubkey\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"_pubkey\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"_preDepositSignature\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"_depositSignature\",\"type\":\"bytes\"}],\"name\":\"onlyValidKeys\",\"outputs\":[],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_name\",\"type\":\"string\"}],\"name\":\"onlyValidName\",\"outputs\":[],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"name\":\"poolAddressById\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"poolIdArray\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes[]\",\"name\":\"_pubkeys\",\"type\":\"bytes[]\"}],\"name\":\"processValidatorExitList\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"renounceRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"revokeRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"staderConfig\",\"outputs\":[{\"internalType\":\"contractIStaderConfig\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"_poolId\",\"type\":\"uint8\"},{\"internalType\":\"address\",\"name\":\"_newPoolAddress\",\"type\":\"address\"}],\"name\":\"updatePoolAddress\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_staderConfig\",\"type\":\"address\"}],\"name\":\"updateStaderConfig\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x608060405234801562000010575f80fd5b50604051620023d9380380620023d98339810160408190526200003391620002e6565b5f54610100900460ff16158080156200005257505f54600160ff909116105b806200006d5750303b1580156200006d57505f5460ff166001145b620000d65760405162461bcd60e51b815260206004820152602e60248201527f496e697469616c697a61626c653a20636f6e747261637420697320616c72656160448201526d191e481a5b9a5d1a585b1a5e995960921b60648201526084015b60405180910390fd5b5f805460ff191660011790558015620000f8575f805461ff0019166101001790555b62000103836200018e565b6200010e826200018e565b62000118620001b9565b609780546001600160a01b0319166001600160a01b0384161790556200013f5f8462000227565b801562000185575f805461ff0019169055604051600181527f7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb38474024989060200160405180910390a15b5050506200031c565b6001600160a01b038116620001b65760405163d92e233d60e01b815260040160405180910390fd5b50565b5f54610100900460ff16620002255760405162461bcd60e51b815260206004820152602b60248201527f496e697469616c697a61626c653a20636f6e7472616374206973206e6f74206960448201526a6e697469616c697a696e6760a81b6064820152608401620000cd565b565b5f8281526065602090815260408083206001600160a01b038516845290915290205460ff16620002c6575f8281526065602090815260408083206001600160a01b03851684529091529020805460ff19166001179055620002853390565b6001600160a01b0316816001600160a01b0316837f2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d60405160405180910390a45b5050565b80516001600160a01b0381168114620002e1575f80fd5b919050565b5f8060408385031215620002f8575f80fd5b6200030383620002ca565b91506200031360208401620002ca565b90509250929050565b6120af806200032a5f395ff3fe608060405234801561000f575f80fd5b50600436106101e7575f3560e01c806399d055c811610109578063b7b32d4b1161009e578063df8984fe1161006e578063df8984fe14610463578063ef7bba861461048b578063f9c4dda41461049e578063ff6bceec146104b1575f80fd5b8063b7b32d4b14610417578063bda0bc891461042a578063d547741f1461043d578063d97ac84714610450575f80fd5b8063a217fddf116100d9578063a217fddf146103ba578063a92e1faf146103c1578063afc2afba146103d6578063b0ef1e1814610404575f80fd5b806399d055c81461036e5780639ee804cb146103815780639f55941b146103945780639f7053f5146103a7575f80fd5b80635d713ec31161017f57806377c359e11161014f57806377c359e11461031b5780638465bef5146103235780638e43c53a1461034857806391d148541461035b575f80fd5b80635d713ec3146102cf57806363d0d5c0146102e25780636cdf1252146102f55780637526d42914610308575f80fd5b80632f2ff15d116101ba5780632f2ff15d1461026957806336514d9f1461027e57806336568abe14610291578063490ffa35146102a4575f80fd5b806301ffc9a7146101eb5780631ec2db3c14610213578063248a9ca314610234578063261d41f514610256575b5f80fd5b6101fe6101f9366004611abf565b6104c4565b60405190151581526020015b60405180910390f35b610226610221366004611af4565b6104fa565b60405190815260200161020a565b610226610242366004611b0f565b5f9081526065602052604090206001015490565b610226610264366004611af4565b610596565b61027c610277366004611b3a565b61063e565b005b6101fe61028c366004611bad565b610667565b61027c61029f366004611b3a565b610756565b6097546102b7906001600160a01b031681565b6040516001600160a01b03909116815260200161020a565b6102266102dd366004611bec565b6107d9565b61027c6102f0366004611c2f565b61089a565b6101fe610303366004611af4565b610971565b6102b7610316366004611af4565b6109e8565b610226610a89565b610336610331366004611b0f565b610b00565b60405160ff909116815260200161020a565b610336610356366004611c5b565b610b31565b6101fe610369366004611b3a565b610c38565b6102b761037c366004611af4565b610c62565b61027c61038f366004611c5b565b610cdf565b61027c6103a2366004611c76565b610d47565b61027c6103b5366004611bad565b610dda565b6102265f81565b6103c9610e8f565b60405161020a9190611d09565b6103e96103e4366004611d4f565b610f02565b6040805193845260208401929092529082015260600161020a565b610226610412366004611af4565b61111e565b610226610425366004611af4565b61119b565b610336610438366004611bad565b61120b565b61027c61044b366004611b3a565b611319565b61027c61045e366004611c2f565b61133d565b6102b7610471366004611af4565b60986020525f90815260409020546001600160a01b031681565b610226610499366004611af4565b6113e8565b6101fe6104ac366004611c5b565b611458565b61027c6104bf366004611d79565b611515565b5f6001600160e01b03198216637965db0b60e01b14806104f457506301ffc9a760e01b6001600160e01b03198316145b92915050565b5f8161050581610971565b61052257604051636c7cd82760e01b815260040160405180910390fd5b5f61052c84610c62565b9050806001600160a01b03166377c359e16040518163ffffffff1660e01b8152600401602060405180830381865afa15801561056a573d5f803e3d5ffd5b505050506040513d601f19601f8201168201806040525081019061058e9190611de8565b949350505050565b5f816105a181610971565b6105be57604051636c7cd82760e01b815260040160405180910390fd5b60ff83165f908152609860209081526040918290205482516358710f4560e11b815292516001600160a01b039091169263b0e21e8a9260048083019391928290030181865afa158015610613573d5f803e3d5ffd5b505050506040513d601f19601f820116820180604052508101906106379190611de8565b9392505050565b5f82815260656020526040902060010154610658816115a0565b61066283836115ad565b505050565b5f8061067260995490565b90505f5b8181101561074c575f6106b86099838154811061069557610695611dff565b905f5260205f2090602091828204019190069054906101000a900460ff16610c62565b6040516336514d9f60e01b81529091506001600160a01b038216906336514d9f906106e99089908990600401611e13565b602060405180830381865afa158015610704573d5f803e3d5ffd5b505050506040513d601f19601f820116820180604052508101906107289190611e41565b1561073957600193505050506104f4565b508061074481611e74565b915050610676565b505f949350505050565b6001600160a01b03811633146107cb5760405162461bcd60e51b815260206004820152602f60248201527f416363657373436f6e74726f6c3a2063616e206f6e6c792072656e6f756e636560448201526e103937b632b9903337b91039b2b63360891b60648201526084015b60405180910390fd5b6107d58282611632565b5050565b5f846107e481610971565b61080157604051636c7cd82760e01b815260040160405180910390fd5b5f61080b87610c62565b6040516322896f3b60e21b81526001600160a01b038881166004830152602482018890526044820187905291925090821690638a25bcec90606401602060405180830381865afa158015610861573d5f803e3d5ffd5b505050506040513d601f19601f820116820180604052508101906108859190611e8c565b67ffffffffffffffff16979650505050505050565b6097546108b19033906001600160a01b0316611698565b6108ba8161171d565b6108c48282611744565b609980546001810190915560208082047f72a152ddfb8e864297c917af52ea6c1c68aead0fee1a62673fcc7e0c94979d0001805460ff868116601f9095166101000a8581029102199091161790555f8281526098825260409081902080546001600160a01b0386166001600160a01b0319909116811790915590519081527f697362d5a2939aff718fb2db4145cb1b4ffc68872c82b2e64d805d8e94845af1910160405180910390a25050565b5f8061097c60995490565b90505f5b818110156109df578360ff166099828154811061099f5761099f611dff565b5f9182526020918290209181049091015460ff601f9092166101000a900416036109cd575060019392505050565b806109d781611e74565b915050610980565b505f9392505050565b5f816109f381610971565b610a1057604051636c7cd82760e01b815260040160405180910390fd5b60ff83165f9081526098602090815260409182902054825163f74b4cd160e01b815292516001600160a01b039091169263f74b4cd19260048083019391928290030181865afa158015610a65573d5f803e3d5ffd5b505050506040513d601f19601f820116820180604052508101906106379190611eb3565b5f805f610a9560995490565b90505f5b81811015610af857610ada60998281548110610ab757610ab7611dff565b905f5260205f2090602091828204019190069054906101000a900460ff166104fa565b610ae49084611ece565b925080610af081611e74565b915050610a99565b509092915050565b60998181548110610b0f575f80fd5b905f5260205f209060209182820401919006915054906101000a900460ff1681565b5f80610b3c60995490565b90505f5b81811015610c1e575f610b5f6099838154811061069557610695611dff565b604051633e71376960e21b81526001600160a01b0387811660048301529192509082169063f9c4dda490602401602060405180830381865afa158015610ba7573d5f803e3d5ffd5b505050506040513d601f19601f82011682018060405250810190610bcb9190611e41565b15610c0b5760998281548110610be357610be3611dff565b905f5260205f2090602091828204019190069054906101000a900460ff169350505050919050565b5080610c1681611e74565b915050610b40565b50604051633e945e6d60e21b815260040160405180910390fd5b5f9182526065602090815260408084206001600160a01b0393909316845291905290205460ff1690565b5f81610c6d81610971565b610c8a57604051636c7cd82760e01b815260040160405180910390fd5b60ff83165f90815260986020908152604091829020548251632dbecfeb60e21b815292516001600160a01b039091169263b6fb3fac9260048083019391928290030181865afa158015610a65573d5f803e3d5ffd5b5f610ce9816115a0565b610cf28261171d565b609780546001600160a01b0319166001600160a01b0384169081179091556040519081527fdb2219043d7b197cb235f1af0cf6d782d77dee3de19e3f4fb6d39aae633b44859060200160405180910390a15050565b60308514610d6857604051636eeda58f60e11b815260040160405180910390fd5b60608314610d895760405163a52467c160e01b815260040160405180910390fd5b60608114610daa5760405163a52467c160e01b815260040160405180910390fd5b610db48686610667565b15610dd25760405163c7de6d2b60e01b815260040160405180910390fd5b505050505050565b5f819003610dfb57604051630f35201760e41b815260040160405180910390fd5b60975f9054906101000a90046001600160a01b03166001600160a01b03166310deba2b6040518163ffffffff1660e01b8152600401602060405180830381865afa158015610e4b573d5f803e3d5ffd5b505050506040513d601f19601f82011682018060405250810190610e6f9190611de8565b8111156107d557604051630ea0087f60e11b815260040160405180910390fd5b60606099805480602002602001604051908101604052809291908181526020018280548015610ef857602002820191905f5260205f20905f905b825461010083900a900460ff16815260206001928301818104948501949093039092029101808411610ec95790505b5050505050905090565b5f805f8060975f9054906101000a90046001600160a01b03166001600160a01b031663ff387f3a6040518163ffffffff1660e01b8152600401602060405180830381865afa158015610f56573d5f803e3d5ffd5b505050506040513d601f19601f82011682018060405250810190610f7a9190611de8565b90505f610f86876113e8565b90505f610f938284611ee1565b90505f610f9f89610596565b90505f610fab8a61111e565b90505f85610fb9858c611ef4565b610fc39190611f0b565b905060975f9054906101000a90046001600160a01b03166001600160a01b0316637ae316d06040518163ffffffff1660e01b8152600401602060405180830381865afa158015611015573d5f803e3d5ffd5b505050506040513d601f19601f820116820180604052508101906110399190611de8565b6110438285611ef4565b61104d9190611f0b565b96508561105a868c611ef4565b6110649190611f0b565b975060975f9054906101000a90046001600160a01b03166001600160a01b0316637ae316d06040518163ffffffff1660e01b8152600401602060405180830381865afa1580156110b6573d5f803e3d5ffd5b505050506040513d601f19601f820116820180604052508101906110da9190611de8565b6110e48284611ef4565b6110ee9190611f0b565b6110f89089611ece565b975087611105888c611ee1565b61110f9190611ee1565b98505050505050509250925092565b5f8161112981610971565b61114657604051636c7cd82760e01b815260040160405180910390fd5b60ff83165f908152609860209081526040918290205482516389afc0f160e01b815292516001600160a01b03909116926389afc0f19260048083019391928290030181865afa158015610613573d5f803e3d5ffd5b5f816111a681610971565b6111c357604051636c7cd82760e01b815260040160405180910390fd5b5f6111cd84610c62565b9050806001600160a01b0316637bd977d96040518163ffffffff1660e01b8152600401602060405180830381865afa15801561056a573d5f803e3d5ffd5b5f8061121660995490565b90505f5b818110156112fc575f6112396099838154811061069557610695611dff565b6040516336514d9f60e01b81529091506001600160a01b038216906336514d9f9061126a9089908990600401611e13565b602060405180830381865afa158015611285573d5f803e3d5ffd5b505050506040513d601f19601f820116820180604052508101906112a99190611e41565b156112e957609982815481106112c1576112c1611dff565b905f5260205f2090602091828204019190069054906101000a900460ff1693505050506104f4565b50806112f481611e74565b91505061121a565b506040516001623899b360e11b0319815260040160405180910390fd5b5f82815260656020526040902060010154611333816115a0565b6106628383611632565b8161134781610971565b61136457604051636c7cd82760e01b815260040160405180910390fd5b5f61136e816115a0565b6113778361171d565b6113818484611744565b60ff84165f8181526098602090815260409182902080546001600160a01b0319166001600160a01b03881690811790915591519182527ff732deab68331ad20834cfc15d686fed4bac945cf3af5d7f729205c9bf846199910160405180910390a250505050565b5f816113f381610971565b61141057604051636c7cd82760e01b815260040160405180910390fd5b5f61141a84610c62565b9050806001600160a01b031663b01db0786040518163ffffffff1660e01b8152600401602060405180830381865afa15801561056a573d5f803e3d5ffd5b5f8061146360995490565b90505f5b818110156109df575f6114866099838154811061069557610695611dff565b604051633e71376960e21b81526001600160a01b0387811660048301529192509082169063f9c4dda490602401602060405180830381865afa1580156114ce573d5f803e3d5ffd5b505050506040513d601f19601f820116820180604052508101906114f29190611e41565b1561150257506001949350505050565b508061150d81611e74565b915050611467565b60975461152c9033906001600160a01b0316611839565b805f5b8181101561159a577fce4931e23262c5aa14c3b95b5e67c07bb38447fda706a4e5e4019e4f7014281284848381811061156a5761156a611dff565b905060200281019061157c9190611f2a565b60405161158a929190611e13565b60405180910390a160010161152f565b50505050565b6115aa81336118be565b50565b6115b78282610c38565b6107d5575f8281526065602090815260408083206001600160a01b03851684529091529020805460ff191660011790556115ee3390565b6001600160a01b0316816001600160a01b0316837f2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d60405160405180910390a45050565b61163c8282610c38565b156107d5575f8281526065602090815260408083206001600160a01b0385168085529252808320805460ff1916905551339285917ff6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b9190a45050565b6040516318903ee760e21b81526001600160a01b038381166004830152821690636240fb9c90602401602060405180830381865afa1580156116dc573d5f803e3d5ffd5b505050506040513d601f19601f820116820180604052508101906117009190611e41565b6107d55760405163c4230ae360e01b815260040160405180910390fd5b6001600160a01b0381166115aa5760405163d92e233d60e01b815260040160405180910390fd5b8160ff16816001600160a01b031663b6fb3fac6040518163ffffffff1660e01b8152600401602060405180830381865afa158015611784573d5f803e3d5ffd5b505050506040513d601f19601f820116820180604052508101906117a89190611eb3565b6001600160a01b031663e0d7d0e96040518163ffffffff1660e01b8152600401602060405180830381865afa1580156117e3573d5f803e3d5ffd5b505050506040513d601f19601f820116820180604052508101906118079190611f6d565b60ff1614158061181b575061181b82610971565b156107d55760405163870a828560e01b815260040160405180910390fd5b6040516353f5713b60e01b81526001600160a01b0383811660048301528216906353f5713b90602401602060405180830381865afa15801561187d573d5f803e3d5ffd5b505050506040513d601f19601f820116820180604052508101906118a19190611e41565b6107d55760405163a5523ee560e01b815260040160405180910390fd5b6118c88282610c38565b6107d5576118d581611917565b6118e0836020611929565b6040516020016118f1929190611faa565b60408051601f198184030181529082905262461bcd60e51b82526107c29160040161201e565b60606104f46001600160a01b03831660145b60605f611937836002611ef4565b611942906002611ece565b67ffffffffffffffff81111561195a5761195a612050565b6040519080825280601f01601f191660200182016040528015611984576020820181803683370190505b509050600360fc1b815f8151811061199e5761199e611dff565b60200101906001600160f81b03191690815f1a905350600f60fb1b816001815181106119cc576119cc611dff565b60200101906001600160f81b03191690815f1a9053505f6119ee846002611ef4565b6119f9906001611ece565b90505b6001811115611a70576f181899199a1a9b1b9c1cb0b131b232b360811b85600f1660108110611a2d57611a2d611dff565b1a60f81b828281518110611a4357611a43611dff565b60200101906001600160f81b03191690815f1a90535060049490941c93611a6981612064565b90506119fc565b5083156106375760405162461bcd60e51b815260206004820181905260248201527f537472696e67733a20686578206c656e67746820696e73756666696369656e7460448201526064016107c2565b5f60208284031215611acf575f80fd5b81356001600160e01b031981168114610637575f80fd5b60ff811681146115aa575f80fd5b5f60208284031215611b04575f80fd5b813561063781611ae6565b5f60208284031215611b1f575f80fd5b5035919050565b6001600160a01b03811681146115aa575f80fd5b5f8060408385031215611b4b575f80fd5b823591506020830135611b5d81611b26565b809150509250929050565b5f8083601f840112611b78575f80fd5b50813567ffffffffffffffff811115611b8f575f80fd5b602083019150836020828501011115611ba6575f80fd5b9250929050565b5f8060208385031215611bbe575f80fd5b823567ffffffffffffffff811115611bd4575f80fd5b611be085828601611b68565b90969095509350505050565b5f805f8060808587031215611bff575f80fd5b8435611c0a81611ae6565b93506020850135611c1a81611b26565b93969395505050506040820135916060013590565b5f8060408385031215611c40575f80fd5b8235611c4b81611ae6565b91506020830135611b5d81611b26565b5f60208284031215611c6b575f80fd5b813561063781611b26565b5f805f805f8060608789031215611c8b575f80fd5b863567ffffffffffffffff80821115611ca2575f80fd5b611cae8a838b01611b68565b90985096506020890135915080821115611cc6575f80fd5b611cd28a838b01611b68565b90965094506040890135915080821115611cea575f80fd5b50611cf789828a01611b68565b979a9699509497509295939492505050565b602080825282518282018190525f9190848201906040850190845b81811015611d4357835160ff1683529284019291840191600101611d24565b50909695505050505050565b5f8060408385031215611d60575f80fd5b8235611d6b81611ae6565b946020939093013593505050565b5f8060208385031215611d8a575f80fd5b823567ffffffffffffffff80821115611da1575f80fd5b818501915085601f830112611db4575f80fd5b813581811115611dc2575f80fd5b8660208260051b8501011115611dd6575f80fd5b60209290920196919550909350505050565b5f60208284031215611df8575f80fd5b5051919050565b634e487b7160e01b5f52603260045260245ffd5b60208152816020820152818360408301375f818301604090810191909152601f909201601f19160101919050565b5f60208284031215611e51575f80fd5b81518015158114610637575f80fd5b634e487b7160e01b5f52601160045260245ffd5b5f60018201611e8557611e85611e60565b5060010190565b5f60208284031215611e9c575f80fd5b815167ffffffffffffffff81168114610637575f80fd5b5f60208284031215611ec3575f80fd5b815161063781611b26565b808201808211156104f4576104f4611e60565b818103818111156104f4576104f4611e60565b80820281158282048414176104f4576104f4611e60565b5f82611f2557634e487b7160e01b5f52601260045260245ffd5b500490565b5f808335601e19843603018112611f3f575f80fd5b83018035915067ffffffffffffffff821115611f59575f80fd5b602001915036819003821315611ba6575f80fd5b5f60208284031215611f7d575f80fd5b815161063781611ae6565b5f5b83811015611fa2578181015183820152602001611f8a565b50505f910152565b7f416363657373436f6e74726f6c3a206163636f756e742000000000000000000081525f8351611fe1816017850160208801611f88565b7001034b99036b4b9b9b4b733903937b6329607d1b6017918401918201528351612012816028840160208801611f88565b01602801949350505050565b602081525f825180602084015261203c816040850160208701611f88565b601f01601f19169190910160400192915050565b634e487b7160e01b5f52604160045260245ffd5b5f8161207257612072611e60565b505f19019056fea2646970667358221220ca3f28d91dac665c1dc4418c1eb9875829632fc125e9aecaa4590b6ebee23f3264736f6c63430008140033",
}

// PoolUtilsABI is the input ABI used to generate the binding from.
// Deprecated: Use PoolUtilsMetaData.ABI instead.
var PoolUtilsABI = PoolUtilsMetaData.ABI

// PoolUtilsBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use PoolUtilsMetaData.Bin instead.
var PoolUtilsBin = PoolUtilsMetaData.Bin

// DeployPoolUtils deploys a new Ethereum contract, binding an instance of PoolUtils to it.
func DeployPoolUtils(auth *bind.TransactOpts, backend bind.ContractBackend, _admin common.Address, _staderConfig common.Address) (common.Address, *types.Transaction, *PoolUtils, error) {
	parsed, err := PoolUtilsMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(PoolUtilsBin), backend, _admin, _staderConfig)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &PoolUtils{PoolUtilsCaller: PoolUtilsCaller{contract: contract}, PoolUtilsTransactor: PoolUtilsTransactor{contract: contract}, PoolUtilsFilterer: PoolUtilsFilterer{contract: contract}}, nil
}

// PoolUtils is an auto generated Go binding around an Ethereum contract.
type PoolUtils struct {
	PoolUtilsCaller     // Read-only binding to the contract
	PoolUtilsTransactor // Write-only binding to the contract
	PoolUtilsFilterer   // Log filterer for contract events
}

// PoolUtilsCaller is an auto generated read-only Go binding around an Ethereum contract.
type PoolUtilsCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PoolUtilsTransactor is an auto generated write-only Go binding around an Ethereum contract.
type PoolUtilsTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PoolUtilsFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type PoolUtilsFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PoolUtilsSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type PoolUtilsSession struct {
	Contract     *PoolUtils        // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// PoolUtilsCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type PoolUtilsCallerSession struct {
	Contract *PoolUtilsCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts    // Call options to use throughout this session
}

// PoolUtilsTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type PoolUtilsTransactorSession struct {
	Contract     *PoolUtilsTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// PoolUtilsRaw is an auto generated low-level Go binding around an Ethereum contract.
type PoolUtilsRaw struct {
	Contract *PoolUtils // Generic contract binding to access the raw methods on
}

// PoolUtilsCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type PoolUtilsCallerRaw struct {
	Contract *PoolUtilsCaller // Generic read-only contract binding to access the raw methods on
}

// PoolUtilsTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type PoolUtilsTransactorRaw struct {
	Contract *PoolUtilsTransactor // Generic write-only contract binding to access the raw methods on
}

// NewPoolUtils creates a new instance of PoolUtils, bound to a specific deployed contract.
func NewPoolUtils(address common.Address, backend bind.ContractBackend) (*PoolUtils, error) {
	contract, err := bindPoolUtils(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &PoolUtils{PoolUtilsCaller: PoolUtilsCaller{contract: contract}, PoolUtilsTransactor: PoolUtilsTransactor{contract: contract}, PoolUtilsFilterer: PoolUtilsFilterer{contract: contract}}, nil
}

// NewPoolUtilsCaller creates a new read-only instance of PoolUtils, bound to a specific deployed contract.
func NewPoolUtilsCaller(address common.Address, caller bind.ContractCaller) (*PoolUtilsCaller, error) {
	contract, err := bindPoolUtils(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &PoolUtilsCaller{contract: contract}, nil
}

// NewPoolUtilsTransactor creates a new write-only instance of PoolUtils, bound to a specific deployed contract.
func NewPoolUtilsTransactor(address common.Address, transactor bind.ContractTransactor) (*PoolUtilsTransactor, error) {
	contract, err := bindPoolUtils(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &PoolUtilsTransactor{contract: contract}, nil
}

// NewPoolUtilsFilterer creates a new log filterer instance of PoolUtils, bound to a specific deployed contract.
func NewPoolUtilsFilterer(address common.Address, filterer bind.ContractFilterer) (*PoolUtilsFilterer, error) {
	contract, err := bindPoolUtils(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &PoolUtilsFilterer{contract: contract}, nil
}

// bindPoolUtils binds a generic wrapper to an already deployed contract.
func bindPoolUtils(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := PoolUtilsMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_PoolUtils *PoolUtilsRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _PoolUtils.Contract.PoolUtilsCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_PoolUtils *PoolUtilsRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PoolUtils.Contract.PoolUtilsTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_PoolUtils *PoolUtilsRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _PoolUtils.Contract.PoolUtilsTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_PoolUtils *PoolUtilsCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _PoolUtils.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_PoolUtils *PoolUtilsTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PoolUtils.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_PoolUtils *PoolUtilsTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _PoolUtils.Contract.contract.Transact(opts, method, params...)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_PoolUtils *PoolUtilsCaller) DEFAULTADMINROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _PoolUtils.contract.Call(opts, &out, "DEFAULT_ADMIN_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_PoolUtils *PoolUtilsSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _PoolUtils.Contract.DEFAULTADMINROLE(&_PoolUtils.CallOpts)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_PoolUtils *PoolUtilsCallerSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _PoolUtils.Contract.DEFAULTADMINROLE(&_PoolUtils.CallOpts)
}

// CalculateRewardShare is a free data retrieval call binding the contract method 0xafc2afba.
//
// Solidity: function calculateRewardShare(uint8 _poolId, uint256 _totalRewards) view returns(uint256 userShare, uint256 operatorShare, uint256 protocolShare)
func (_PoolUtils *PoolUtilsCaller) CalculateRewardShare(opts *bind.CallOpts, _poolId uint8, _totalRewards *big.Int) (struct {
	UserShare     *big.Int
	OperatorShare *big.Int
	ProtocolShare *big.Int
}, error) {
	var out []interface{}
	err := _PoolUtils.contract.Call(opts, &out, "calculateRewardShare", _poolId, _totalRewards)

	outstruct := new(struct {
		UserShare     *big.Int
		OperatorShare *big.Int
		ProtocolShare *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.UserShare = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.OperatorShare = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.ProtocolShare = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// CalculateRewardShare is a free data retrieval call binding the contract method 0xafc2afba.
//
// Solidity: function calculateRewardShare(uint8 _poolId, uint256 _totalRewards) view returns(uint256 userShare, uint256 operatorShare, uint256 protocolShare)
func (_PoolUtils *PoolUtilsSession) CalculateRewardShare(_poolId uint8, _totalRewards *big.Int) (struct {
	UserShare     *big.Int
	OperatorShare *big.Int
	ProtocolShare *big.Int
}, error) {
	return _PoolUtils.Contract.CalculateRewardShare(&_PoolUtils.CallOpts, _poolId, _totalRewards)
}

// CalculateRewardShare is a free data retrieval call binding the contract method 0xafc2afba.
//
// Solidity: function calculateRewardShare(uint8 _poolId, uint256 _totalRewards) view returns(uint256 userShare, uint256 operatorShare, uint256 protocolShare)
func (_PoolUtils *PoolUtilsCallerSession) CalculateRewardShare(_poolId uint8, _totalRewards *big.Int) (struct {
	UserShare     *big.Int
	OperatorShare *big.Int
	ProtocolShare *big.Int
}, error) {
	return _PoolUtils.Contract.CalculateRewardShare(&_PoolUtils.CallOpts, _poolId, _totalRewards)
}

// GetActiveValidatorCountByPool is a free data retrieval call binding the contract method 0x1ec2db3c.
//
// Solidity: function getActiveValidatorCountByPool(uint8 _poolId) view returns(uint256)
func (_PoolUtils *PoolUtilsCaller) GetActiveValidatorCountByPool(opts *bind.CallOpts, _poolId uint8) (*big.Int, error) {
	var out []interface{}
	err := _PoolUtils.contract.Call(opts, &out, "getActiveValidatorCountByPool", _poolId)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetActiveValidatorCountByPool is a free data retrieval call binding the contract method 0x1ec2db3c.
//
// Solidity: function getActiveValidatorCountByPool(uint8 _poolId) view returns(uint256)
func (_PoolUtils *PoolUtilsSession) GetActiveValidatorCountByPool(_poolId uint8) (*big.Int, error) {
	return _PoolUtils.Contract.GetActiveValidatorCountByPool(&_PoolUtils.CallOpts, _poolId)
}

// GetActiveValidatorCountByPool is a free data retrieval call binding the contract method 0x1ec2db3c.
//
// Solidity: function getActiveValidatorCountByPool(uint8 _poolId) view returns(uint256)
func (_PoolUtils *PoolUtilsCallerSession) GetActiveValidatorCountByPool(_poolId uint8) (*big.Int, error) {
	return _PoolUtils.Contract.GetActiveValidatorCountByPool(&_PoolUtils.CallOpts, _poolId)
}

// GetCollateralETH is a free data retrieval call binding the contract method 0xef7bba86.
//
// Solidity: function getCollateralETH(uint8 _poolId) view returns(uint256)
func (_PoolUtils *PoolUtilsCaller) GetCollateralETH(opts *bind.CallOpts, _poolId uint8) (*big.Int, error) {
	var out []interface{}
	err := _PoolUtils.contract.Call(opts, &out, "getCollateralETH", _poolId)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetCollateralETH is a free data retrieval call binding the contract method 0xef7bba86.
//
// Solidity: function getCollateralETH(uint8 _poolId) view returns(uint256)
func (_PoolUtils *PoolUtilsSession) GetCollateralETH(_poolId uint8) (*big.Int, error) {
	return _PoolUtils.Contract.GetCollateralETH(&_PoolUtils.CallOpts, _poolId)
}

// GetCollateralETH is a free data retrieval call binding the contract method 0xef7bba86.
//
// Solidity: function getCollateralETH(uint8 _poolId) view returns(uint256)
func (_PoolUtils *PoolUtilsCallerSession) GetCollateralETH(_poolId uint8) (*big.Int, error) {
	return _PoolUtils.Contract.GetCollateralETH(&_PoolUtils.CallOpts, _poolId)
}

// GetNodeRegistry is a free data retrieval call binding the contract method 0x99d055c8.
//
// Solidity: function getNodeRegistry(uint8 _poolId) view returns(address)
func (_PoolUtils *PoolUtilsCaller) GetNodeRegistry(opts *bind.CallOpts, _poolId uint8) (common.Address, error) {
	var out []interface{}
	err := _PoolUtils.contract.Call(opts, &out, "getNodeRegistry", _poolId)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetNodeRegistry is a free data retrieval call binding the contract method 0x99d055c8.
//
// Solidity: function getNodeRegistry(uint8 _poolId) view returns(address)
func (_PoolUtils *PoolUtilsSession) GetNodeRegistry(_poolId uint8) (common.Address, error) {
	return _PoolUtils.Contract.GetNodeRegistry(&_PoolUtils.CallOpts, _poolId)
}

// GetNodeRegistry is a free data retrieval call binding the contract method 0x99d055c8.
//
// Solidity: function getNodeRegistry(uint8 _poolId) view returns(address)
func (_PoolUtils *PoolUtilsCallerSession) GetNodeRegistry(_poolId uint8) (common.Address, error) {
	return _PoolUtils.Contract.GetNodeRegistry(&_PoolUtils.CallOpts, _poolId)
}

// GetOperatorFee is a free data retrieval call binding the contract method 0xb0ef1e18.
//
// Solidity: function getOperatorFee(uint8 _poolId) view returns(uint256)
func (_PoolUtils *PoolUtilsCaller) GetOperatorFee(opts *bind.CallOpts, _poolId uint8) (*big.Int, error) {
	var out []interface{}
	err := _PoolUtils.contract.Call(opts, &out, "getOperatorFee", _poolId)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetOperatorFee is a free data retrieval call binding the contract method 0xb0ef1e18.
//
// Solidity: function getOperatorFee(uint8 _poolId) view returns(uint256)
func (_PoolUtils *PoolUtilsSession) GetOperatorFee(_poolId uint8) (*big.Int, error) {
	return _PoolUtils.Contract.GetOperatorFee(&_PoolUtils.CallOpts, _poolId)
}

// GetOperatorFee is a free data retrieval call binding the contract method 0xb0ef1e18.
//
// Solidity: function getOperatorFee(uint8 _poolId) view returns(uint256)
func (_PoolUtils *PoolUtilsCallerSession) GetOperatorFee(_poolId uint8) (*big.Int, error) {
	return _PoolUtils.Contract.GetOperatorFee(&_PoolUtils.CallOpts, _poolId)
}

// GetOperatorPoolId is a free data retrieval call binding the contract method 0x8e43c53a.
//
// Solidity: function getOperatorPoolId(address _operAddr) view returns(uint8)
func (_PoolUtils *PoolUtilsCaller) GetOperatorPoolId(opts *bind.CallOpts, _operAddr common.Address) (uint8, error) {
	var out []interface{}
	err := _PoolUtils.contract.Call(opts, &out, "getOperatorPoolId", _operAddr)

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// GetOperatorPoolId is a free data retrieval call binding the contract method 0x8e43c53a.
//
// Solidity: function getOperatorPoolId(address _operAddr) view returns(uint8)
func (_PoolUtils *PoolUtilsSession) GetOperatorPoolId(_operAddr common.Address) (uint8, error) {
	return _PoolUtils.Contract.GetOperatorPoolId(&_PoolUtils.CallOpts, _operAddr)
}

// GetOperatorPoolId is a free data retrieval call binding the contract method 0x8e43c53a.
//
// Solidity: function getOperatorPoolId(address _operAddr) view returns(uint8)
func (_PoolUtils *PoolUtilsCallerSession) GetOperatorPoolId(_operAddr common.Address) (uint8, error) {
	return _PoolUtils.Contract.GetOperatorPoolId(&_PoolUtils.CallOpts, _operAddr)
}

// GetOperatorTotalNonTerminalKeys is a free data retrieval call binding the contract method 0x5d713ec3.
//
// Solidity: function getOperatorTotalNonTerminalKeys(uint8 _poolId, address _nodeOperator, uint256 _startIndex, uint256 _endIndex) view returns(uint256)
func (_PoolUtils *PoolUtilsCaller) GetOperatorTotalNonTerminalKeys(opts *bind.CallOpts, _poolId uint8, _nodeOperator common.Address, _startIndex *big.Int, _endIndex *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _PoolUtils.contract.Call(opts, &out, "getOperatorTotalNonTerminalKeys", _poolId, _nodeOperator, _startIndex, _endIndex)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetOperatorTotalNonTerminalKeys is a free data retrieval call binding the contract method 0x5d713ec3.
//
// Solidity: function getOperatorTotalNonTerminalKeys(uint8 _poolId, address _nodeOperator, uint256 _startIndex, uint256 _endIndex) view returns(uint256)
func (_PoolUtils *PoolUtilsSession) GetOperatorTotalNonTerminalKeys(_poolId uint8, _nodeOperator common.Address, _startIndex *big.Int, _endIndex *big.Int) (*big.Int, error) {
	return _PoolUtils.Contract.GetOperatorTotalNonTerminalKeys(&_PoolUtils.CallOpts, _poolId, _nodeOperator, _startIndex, _endIndex)
}

// GetOperatorTotalNonTerminalKeys is a free data retrieval call binding the contract method 0x5d713ec3.
//
// Solidity: function getOperatorTotalNonTerminalKeys(uint8 _poolId, address _nodeOperator, uint256 _startIndex, uint256 _endIndex) view returns(uint256)
func (_PoolUtils *PoolUtilsCallerSession) GetOperatorTotalNonTerminalKeys(_poolId uint8, _nodeOperator common.Address, _startIndex *big.Int, _endIndex *big.Int) (*big.Int, error) {
	return _PoolUtils.Contract.GetOperatorTotalNonTerminalKeys(&_PoolUtils.CallOpts, _poolId, _nodeOperator, _startIndex, _endIndex)
}

// GetPoolIdArray is a free data retrieval call binding the contract method 0xa92e1faf.
//
// Solidity: function getPoolIdArray() view returns(uint8[])
func (_PoolUtils *PoolUtilsCaller) GetPoolIdArray(opts *bind.CallOpts) ([]uint8, error) {
	var out []interface{}
	err := _PoolUtils.contract.Call(opts, &out, "getPoolIdArray")

	if err != nil {
		return *new([]uint8), err
	}

	out0 := *abi.ConvertType(out[0], new([]uint8)).(*[]uint8)

	return out0, err

}

// GetPoolIdArray is a free data retrieval call binding the contract method 0xa92e1faf.
//
// Solidity: function getPoolIdArray() view returns(uint8[])
func (_PoolUtils *PoolUtilsSession) GetPoolIdArray() ([]uint8, error) {
	return _PoolUtils.Contract.GetPoolIdArray(&_PoolUtils.CallOpts)
}

// GetPoolIdArray is a free data retrieval call binding the contract method 0xa92e1faf.
//
// Solidity: function getPoolIdArray() view returns(uint8[])
func (_PoolUtils *PoolUtilsCallerSession) GetPoolIdArray() ([]uint8, error) {
	return _PoolUtils.Contract.GetPoolIdArray(&_PoolUtils.CallOpts)
}

// GetProtocolFee is a free data retrieval call binding the contract method 0x261d41f5.
//
// Solidity: function getProtocolFee(uint8 _poolId) view returns(uint256)
func (_PoolUtils *PoolUtilsCaller) GetProtocolFee(opts *bind.CallOpts, _poolId uint8) (*big.Int, error) {
	var out []interface{}
	err := _PoolUtils.contract.Call(opts, &out, "getProtocolFee", _poolId)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetProtocolFee is a free data retrieval call binding the contract method 0x261d41f5.
//
// Solidity: function getProtocolFee(uint8 _poolId) view returns(uint256)
func (_PoolUtils *PoolUtilsSession) GetProtocolFee(_poolId uint8) (*big.Int, error) {
	return _PoolUtils.Contract.GetProtocolFee(&_PoolUtils.CallOpts, _poolId)
}

// GetProtocolFee is a free data retrieval call binding the contract method 0x261d41f5.
//
// Solidity: function getProtocolFee(uint8 _poolId) view returns(uint256)
func (_PoolUtils *PoolUtilsCallerSession) GetProtocolFee(_poolId uint8) (*big.Int, error) {
	return _PoolUtils.Contract.GetProtocolFee(&_PoolUtils.CallOpts, _poolId)
}

// GetQueuedValidatorCountByPool is a free data retrieval call binding the contract method 0xb7b32d4b.
//
// Solidity: function getQueuedValidatorCountByPool(uint8 _poolId) view returns(uint256)
func (_PoolUtils *PoolUtilsCaller) GetQueuedValidatorCountByPool(opts *bind.CallOpts, _poolId uint8) (*big.Int, error) {
	var out []interface{}
	err := _PoolUtils.contract.Call(opts, &out, "getQueuedValidatorCountByPool", _poolId)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetQueuedValidatorCountByPool is a free data retrieval call binding the contract method 0xb7b32d4b.
//
// Solidity: function getQueuedValidatorCountByPool(uint8 _poolId) view returns(uint256)
func (_PoolUtils *PoolUtilsSession) GetQueuedValidatorCountByPool(_poolId uint8) (*big.Int, error) {
	return _PoolUtils.Contract.GetQueuedValidatorCountByPool(&_PoolUtils.CallOpts, _poolId)
}

// GetQueuedValidatorCountByPool is a free data retrieval call binding the contract method 0xb7b32d4b.
//
// Solidity: function getQueuedValidatorCountByPool(uint8 _poolId) view returns(uint256)
func (_PoolUtils *PoolUtilsCallerSession) GetQueuedValidatorCountByPool(_poolId uint8) (*big.Int, error) {
	return _PoolUtils.Contract.GetQueuedValidatorCountByPool(&_PoolUtils.CallOpts, _poolId)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_PoolUtils *PoolUtilsCaller) GetRoleAdmin(opts *bind.CallOpts, role [32]byte) ([32]byte, error) {
	var out []interface{}
	err := _PoolUtils.contract.Call(opts, &out, "getRoleAdmin", role)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_PoolUtils *PoolUtilsSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _PoolUtils.Contract.GetRoleAdmin(&_PoolUtils.CallOpts, role)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_PoolUtils *PoolUtilsCallerSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _PoolUtils.Contract.GetRoleAdmin(&_PoolUtils.CallOpts, role)
}

// GetSocializingPoolAddress is a free data retrieval call binding the contract method 0x7526d429.
//
// Solidity: function getSocializingPoolAddress(uint8 _poolId) view returns(address)
func (_PoolUtils *PoolUtilsCaller) GetSocializingPoolAddress(opts *bind.CallOpts, _poolId uint8) (common.Address, error) {
	var out []interface{}
	err := _PoolUtils.contract.Call(opts, &out, "getSocializingPoolAddress", _poolId)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetSocializingPoolAddress is a free data retrieval call binding the contract method 0x7526d429.
//
// Solidity: function getSocializingPoolAddress(uint8 _poolId) view returns(address)
func (_PoolUtils *PoolUtilsSession) GetSocializingPoolAddress(_poolId uint8) (common.Address, error) {
	return _PoolUtils.Contract.GetSocializingPoolAddress(&_PoolUtils.CallOpts, _poolId)
}

// GetSocializingPoolAddress is a free data retrieval call binding the contract method 0x7526d429.
//
// Solidity: function getSocializingPoolAddress(uint8 _poolId) view returns(address)
func (_PoolUtils *PoolUtilsCallerSession) GetSocializingPoolAddress(_poolId uint8) (common.Address, error) {
	return _PoolUtils.Contract.GetSocializingPoolAddress(&_PoolUtils.CallOpts, _poolId)
}

// GetTotalActiveValidatorCount is a free data retrieval call binding the contract method 0x77c359e1.
//
// Solidity: function getTotalActiveValidatorCount() view returns(uint256)
func (_PoolUtils *PoolUtilsCaller) GetTotalActiveValidatorCount(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _PoolUtils.contract.Call(opts, &out, "getTotalActiveValidatorCount")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetTotalActiveValidatorCount is a free data retrieval call binding the contract method 0x77c359e1.
//
// Solidity: function getTotalActiveValidatorCount() view returns(uint256)
func (_PoolUtils *PoolUtilsSession) GetTotalActiveValidatorCount() (*big.Int, error) {
	return _PoolUtils.Contract.GetTotalActiveValidatorCount(&_PoolUtils.CallOpts)
}

// GetTotalActiveValidatorCount is a free data retrieval call binding the contract method 0x77c359e1.
//
// Solidity: function getTotalActiveValidatorCount() view returns(uint256)
func (_PoolUtils *PoolUtilsCallerSession) GetTotalActiveValidatorCount() (*big.Int, error) {
	return _PoolUtils.Contract.GetTotalActiveValidatorCount(&_PoolUtils.CallOpts)
}

// GetValidatorPoolId is a free data retrieval call binding the contract method 0xbda0bc89.
//
// Solidity: function getValidatorPoolId(bytes _pubkey) view returns(uint8)
func (_PoolUtils *PoolUtilsCaller) GetValidatorPoolId(opts *bind.CallOpts, _pubkey []byte) (uint8, error) {
	var out []interface{}
	err := _PoolUtils.contract.Call(opts, &out, "getValidatorPoolId", _pubkey)

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// GetValidatorPoolId is a free data retrieval call binding the contract method 0xbda0bc89.
//
// Solidity: function getValidatorPoolId(bytes _pubkey) view returns(uint8)
func (_PoolUtils *PoolUtilsSession) GetValidatorPoolId(_pubkey []byte) (uint8, error) {
	return _PoolUtils.Contract.GetValidatorPoolId(&_PoolUtils.CallOpts, _pubkey)
}

// GetValidatorPoolId is a free data retrieval call binding the contract method 0xbda0bc89.
//
// Solidity: function getValidatorPoolId(bytes _pubkey) view returns(uint8)
func (_PoolUtils *PoolUtilsCallerSession) GetValidatorPoolId(_pubkey []byte) (uint8, error) {
	return _PoolUtils.Contract.GetValidatorPoolId(&_PoolUtils.CallOpts, _pubkey)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_PoolUtils *PoolUtilsCaller) HasRole(opts *bind.CallOpts, role [32]byte, account common.Address) (bool, error) {
	var out []interface{}
	err := _PoolUtils.contract.Call(opts, &out, "hasRole", role, account)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_PoolUtils *PoolUtilsSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _PoolUtils.Contract.HasRole(&_PoolUtils.CallOpts, role, account)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_PoolUtils *PoolUtilsCallerSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _PoolUtils.Contract.HasRole(&_PoolUtils.CallOpts, role, account)
}

// IsExistingOperator is a free data retrieval call binding the contract method 0xf9c4dda4.
//
// Solidity: function isExistingOperator(address _operAddr) view returns(bool)
func (_PoolUtils *PoolUtilsCaller) IsExistingOperator(opts *bind.CallOpts, _operAddr common.Address) (bool, error) {
	var out []interface{}
	err := _PoolUtils.contract.Call(opts, &out, "isExistingOperator", _operAddr)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsExistingOperator is a free data retrieval call binding the contract method 0xf9c4dda4.
//
// Solidity: function isExistingOperator(address _operAddr) view returns(bool)
func (_PoolUtils *PoolUtilsSession) IsExistingOperator(_operAddr common.Address) (bool, error) {
	return _PoolUtils.Contract.IsExistingOperator(&_PoolUtils.CallOpts, _operAddr)
}

// IsExistingOperator is a free data retrieval call binding the contract method 0xf9c4dda4.
//
// Solidity: function isExistingOperator(address _operAddr) view returns(bool)
func (_PoolUtils *PoolUtilsCallerSession) IsExistingOperator(_operAddr common.Address) (bool, error) {
	return _PoolUtils.Contract.IsExistingOperator(&_PoolUtils.CallOpts, _operAddr)
}

// IsExistingPoolId is a free data retrieval call binding the contract method 0x6cdf1252.
//
// Solidity: function isExistingPoolId(uint8 _poolId) view returns(bool)
func (_PoolUtils *PoolUtilsCaller) IsExistingPoolId(opts *bind.CallOpts, _poolId uint8) (bool, error) {
	var out []interface{}
	err := _PoolUtils.contract.Call(opts, &out, "isExistingPoolId", _poolId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsExistingPoolId is a free data retrieval call binding the contract method 0x6cdf1252.
//
// Solidity: function isExistingPoolId(uint8 _poolId) view returns(bool)
func (_PoolUtils *PoolUtilsSession) IsExistingPoolId(_poolId uint8) (bool, error) {
	return _PoolUtils.Contract.IsExistingPoolId(&_PoolUtils.CallOpts, _poolId)
}

// IsExistingPoolId is a free data retrieval call binding the contract method 0x6cdf1252.
//
// Solidity: function isExistingPoolId(uint8 _poolId) view returns(bool)
func (_PoolUtils *PoolUtilsCallerSession) IsExistingPoolId(_poolId uint8) (bool, error) {
	return _PoolUtils.Contract.IsExistingPoolId(&_PoolUtils.CallOpts, _poolId)
}

// IsExistingPubkey is a free data retrieval call binding the contract method 0x36514d9f.
//
// Solidity: function isExistingPubkey(bytes _pubkey) view returns(bool)
func (_PoolUtils *PoolUtilsCaller) IsExistingPubkey(opts *bind.CallOpts, _pubkey []byte) (bool, error) {
	var out []interface{}
	err := _PoolUtils.contract.Call(opts, &out, "isExistingPubkey", _pubkey)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsExistingPubkey is a free data retrieval call binding the contract method 0x36514d9f.
//
// Solidity: function isExistingPubkey(bytes _pubkey) view returns(bool)
func (_PoolUtils *PoolUtilsSession) IsExistingPubkey(_pubkey []byte) (bool, error) {
	return _PoolUtils.Contract.IsExistingPubkey(&_PoolUtils.CallOpts, _pubkey)
}

// IsExistingPubkey is a free data retrieval call binding the contract method 0x36514d9f.
//
// Solidity: function isExistingPubkey(bytes _pubkey) view returns(bool)
func (_PoolUtils *PoolUtilsCallerSession) IsExistingPubkey(_pubkey []byte) (bool, error) {
	return _PoolUtils.Contract.IsExistingPubkey(&_PoolUtils.CallOpts, _pubkey)
}

// OnlyValidKeys is a free data retrieval call binding the contract method 0x9f55941b.
//
// Solidity: function onlyValidKeys(bytes _pubkey, bytes _preDepositSignature, bytes _depositSignature) view returns()
func (_PoolUtils *PoolUtilsCaller) OnlyValidKeys(opts *bind.CallOpts, _pubkey []byte, _preDepositSignature []byte, _depositSignature []byte) error {
	var out []interface{}
	err := _PoolUtils.contract.Call(opts, &out, "onlyValidKeys", _pubkey, _preDepositSignature, _depositSignature)

	if err != nil {
		return err
	}

	return err

}

// OnlyValidKeys is a free data retrieval call binding the contract method 0x9f55941b.
//
// Solidity: function onlyValidKeys(bytes _pubkey, bytes _preDepositSignature, bytes _depositSignature) view returns()
func (_PoolUtils *PoolUtilsSession) OnlyValidKeys(_pubkey []byte, _preDepositSignature []byte, _depositSignature []byte) error {
	return _PoolUtils.Contract.OnlyValidKeys(&_PoolUtils.CallOpts, _pubkey, _preDepositSignature, _depositSignature)
}

// OnlyValidKeys is a free data retrieval call binding the contract method 0x9f55941b.
//
// Solidity: function onlyValidKeys(bytes _pubkey, bytes _preDepositSignature, bytes _depositSignature) view returns()
func (_PoolUtils *PoolUtilsCallerSession) OnlyValidKeys(_pubkey []byte, _preDepositSignature []byte, _depositSignature []byte) error {
	return _PoolUtils.Contract.OnlyValidKeys(&_PoolUtils.CallOpts, _pubkey, _preDepositSignature, _depositSignature)
}

// OnlyValidName is a free data retrieval call binding the contract method 0x9f7053f5.
//
// Solidity: function onlyValidName(string _name) view returns()
func (_PoolUtils *PoolUtilsCaller) OnlyValidName(opts *bind.CallOpts, _name string) error {
	var out []interface{}
	err := _PoolUtils.contract.Call(opts, &out, "onlyValidName", _name)

	if err != nil {
		return err
	}

	return err

}

// OnlyValidName is a free data retrieval call binding the contract method 0x9f7053f5.
//
// Solidity: function onlyValidName(string _name) view returns()
func (_PoolUtils *PoolUtilsSession) OnlyValidName(_name string) error {
	return _PoolUtils.Contract.OnlyValidName(&_PoolUtils.CallOpts, _name)
}

// OnlyValidName is a free data retrieval call binding the contract method 0x9f7053f5.
//
// Solidity: function onlyValidName(string _name) view returns()
func (_PoolUtils *PoolUtilsCallerSession) OnlyValidName(_name string) error {
	return _PoolUtils.Contract.OnlyValidName(&_PoolUtils.CallOpts, _name)
}

// PoolAddressById is a free data retrieval call binding the contract method 0xdf8984fe.
//
// Solidity: function poolAddressById(uint8 ) view returns(address)
func (_PoolUtils *PoolUtilsCaller) PoolAddressById(opts *bind.CallOpts, arg0 uint8) (common.Address, error) {
	var out []interface{}
	err := _PoolUtils.contract.Call(opts, &out, "poolAddressById", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// PoolAddressById is a free data retrieval call binding the contract method 0xdf8984fe.
//
// Solidity: function poolAddressById(uint8 ) view returns(address)
func (_PoolUtils *PoolUtilsSession) PoolAddressById(arg0 uint8) (common.Address, error) {
	return _PoolUtils.Contract.PoolAddressById(&_PoolUtils.CallOpts, arg0)
}

// PoolAddressById is a free data retrieval call binding the contract method 0xdf8984fe.
//
// Solidity: function poolAddressById(uint8 ) view returns(address)
func (_PoolUtils *PoolUtilsCallerSession) PoolAddressById(arg0 uint8) (common.Address, error) {
	return _PoolUtils.Contract.PoolAddressById(&_PoolUtils.CallOpts, arg0)
}

// PoolIdArray is a free data retrieval call binding the contract method 0x8465bef5.
//
// Solidity: function poolIdArray(uint256 ) view returns(uint8)
func (_PoolUtils *PoolUtilsCaller) PoolIdArray(opts *bind.CallOpts, arg0 *big.Int) (uint8, error) {
	var out []interface{}
	err := _PoolUtils.contract.Call(opts, &out, "poolIdArray", arg0)

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// PoolIdArray is a free data retrieval call binding the contract method 0x8465bef5.
//
// Solidity: function poolIdArray(uint256 ) view returns(uint8)
func (_PoolUtils *PoolUtilsSession) PoolIdArray(arg0 *big.Int) (uint8, error) {
	return _PoolUtils.Contract.PoolIdArray(&_PoolUtils.CallOpts, arg0)
}

// PoolIdArray is a free data retrieval call binding the contract method 0x8465bef5.
//
// Solidity: function poolIdArray(uint256 ) view returns(uint8)
func (_PoolUtils *PoolUtilsCallerSession) PoolIdArray(arg0 *big.Int) (uint8, error) {
	return _PoolUtils.Contract.PoolIdArray(&_PoolUtils.CallOpts, arg0)
}

// StaderConfig is a free data retrieval call binding the contract method 0x490ffa35.
//
// Solidity: function staderConfig() view returns(address)
func (_PoolUtils *PoolUtilsCaller) StaderConfig(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _PoolUtils.contract.Call(opts, &out, "staderConfig")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// StaderConfig is a free data retrieval call binding the contract method 0x490ffa35.
//
// Solidity: function staderConfig() view returns(address)
func (_PoolUtils *PoolUtilsSession) StaderConfig() (common.Address, error) {
	return _PoolUtils.Contract.StaderConfig(&_PoolUtils.CallOpts)
}

// StaderConfig is a free data retrieval call binding the contract method 0x490ffa35.
//
// Solidity: function staderConfig() view returns(address)
func (_PoolUtils *PoolUtilsCallerSession) StaderConfig() (common.Address, error) {
	return _PoolUtils.Contract.StaderConfig(&_PoolUtils.CallOpts)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_PoolUtils *PoolUtilsCaller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _PoolUtils.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_PoolUtils *PoolUtilsSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _PoolUtils.Contract.SupportsInterface(&_PoolUtils.CallOpts, interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_PoolUtils *PoolUtilsCallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _PoolUtils.Contract.SupportsInterface(&_PoolUtils.CallOpts, interfaceId)
}

// AddNewPool is a paid mutator transaction binding the contract method 0x63d0d5c0.
//
// Solidity: function addNewPool(uint8 _poolId, address _poolAddress) returns()
func (_PoolUtils *PoolUtilsTransactor) AddNewPool(opts *bind.TransactOpts, _poolId uint8, _poolAddress common.Address) (*types.Transaction, error) {
	return _PoolUtils.contract.Transact(opts, "addNewPool", _poolId, _poolAddress)
}

// AddNewPool is a paid mutator transaction binding the contract method 0x63d0d5c0.
//
// Solidity: function addNewPool(uint8 _poolId, address _poolAddress) returns()
func (_PoolUtils *PoolUtilsSession) AddNewPool(_poolId uint8, _poolAddress common.Address) (*types.Transaction, error) {
	return _PoolUtils.Contract.AddNewPool(&_PoolUtils.TransactOpts, _poolId, _poolAddress)
}

// AddNewPool is a paid mutator transaction binding the contract method 0x63d0d5c0.
//
// Solidity: function addNewPool(uint8 _poolId, address _poolAddress) returns()
func (_PoolUtils *PoolUtilsTransactorSession) AddNewPool(_poolId uint8, _poolAddress common.Address) (*types.Transaction, error) {
	return _PoolUtils.Contract.AddNewPool(&_PoolUtils.TransactOpts, _poolId, _poolAddress)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_PoolUtils *PoolUtilsTransactor) GrantRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _PoolUtils.contract.Transact(opts, "grantRole", role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_PoolUtils *PoolUtilsSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _PoolUtils.Contract.GrantRole(&_PoolUtils.TransactOpts, role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_PoolUtils *PoolUtilsTransactorSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _PoolUtils.Contract.GrantRole(&_PoolUtils.TransactOpts, role, account)
}

// ProcessValidatorExitList is a paid mutator transaction binding the contract method 0xff6bceec.
//
// Solidity: function processValidatorExitList(bytes[] _pubkeys) returns()
func (_PoolUtils *PoolUtilsTransactor) ProcessValidatorExitList(opts *bind.TransactOpts, _pubkeys [][]byte) (*types.Transaction, error) {
	return _PoolUtils.contract.Transact(opts, "processValidatorExitList", _pubkeys)
}

// ProcessValidatorExitList is a paid mutator transaction binding the contract method 0xff6bceec.
//
// Solidity: function processValidatorExitList(bytes[] _pubkeys) returns()
func (_PoolUtils *PoolUtilsSession) ProcessValidatorExitList(_pubkeys [][]byte) (*types.Transaction, error) {
	return _PoolUtils.Contract.ProcessValidatorExitList(&_PoolUtils.TransactOpts, _pubkeys)
}

// ProcessValidatorExitList is a paid mutator transaction binding the contract method 0xff6bceec.
//
// Solidity: function processValidatorExitList(bytes[] _pubkeys) returns()
func (_PoolUtils *PoolUtilsTransactorSession) ProcessValidatorExitList(_pubkeys [][]byte) (*types.Transaction, error) {
	return _PoolUtils.Contract.ProcessValidatorExitList(&_PoolUtils.TransactOpts, _pubkeys)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address account) returns()
func (_PoolUtils *PoolUtilsTransactor) RenounceRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _PoolUtils.contract.Transact(opts, "renounceRole", role, account)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address account) returns()
func (_PoolUtils *PoolUtilsSession) RenounceRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _PoolUtils.Contract.RenounceRole(&_PoolUtils.TransactOpts, role, account)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address account) returns()
func (_PoolUtils *PoolUtilsTransactorSession) RenounceRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _PoolUtils.Contract.RenounceRole(&_PoolUtils.TransactOpts, role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_PoolUtils *PoolUtilsTransactor) RevokeRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _PoolUtils.contract.Transact(opts, "revokeRole", role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_PoolUtils *PoolUtilsSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _PoolUtils.Contract.RevokeRole(&_PoolUtils.TransactOpts, role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_PoolUtils *PoolUtilsTransactorSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _PoolUtils.Contract.RevokeRole(&_PoolUtils.TransactOpts, role, account)
}

// UpdatePoolAddress is a paid mutator transaction binding the contract method 0xd97ac847.
//
// Solidity: function updatePoolAddress(uint8 _poolId, address _newPoolAddress) returns()
func (_PoolUtils *PoolUtilsTransactor) UpdatePoolAddress(opts *bind.TransactOpts, _poolId uint8, _newPoolAddress common.Address) (*types.Transaction, error) {
	return _PoolUtils.contract.Transact(opts, "updatePoolAddress", _poolId, _newPoolAddress)
}

// UpdatePoolAddress is a paid mutator transaction binding the contract method 0xd97ac847.
//
// Solidity: function updatePoolAddress(uint8 _poolId, address _newPoolAddress) returns()
func (_PoolUtils *PoolUtilsSession) UpdatePoolAddress(_poolId uint8, _newPoolAddress common.Address) (*types.Transaction, error) {
	return _PoolUtils.Contract.UpdatePoolAddress(&_PoolUtils.TransactOpts, _poolId, _newPoolAddress)
}

// UpdatePoolAddress is a paid mutator transaction binding the contract method 0xd97ac847.
//
// Solidity: function updatePoolAddress(uint8 _poolId, address _newPoolAddress) returns()
func (_PoolUtils *PoolUtilsTransactorSession) UpdatePoolAddress(_poolId uint8, _newPoolAddress common.Address) (*types.Transaction, error) {
	return _PoolUtils.Contract.UpdatePoolAddress(&_PoolUtils.TransactOpts, _poolId, _newPoolAddress)
}

// UpdateStaderConfig is a paid mutator transaction binding the contract method 0x9ee804cb.
//
// Solidity: function updateStaderConfig(address _staderConfig) returns()
func (_PoolUtils *PoolUtilsTransactor) UpdateStaderConfig(opts *bind.TransactOpts, _staderConfig common.Address) (*types.Transaction, error) {
	return _PoolUtils.contract.Transact(opts, "updateStaderConfig", _staderConfig)
}

// UpdateStaderConfig is a paid mutator transaction binding the contract method 0x9ee804cb.
//
// Solidity: function updateStaderConfig(address _staderConfig) returns()
func (_PoolUtils *PoolUtilsSession) UpdateStaderConfig(_staderConfig common.Address) (*types.Transaction, error) {
	return _PoolUtils.Contract.UpdateStaderConfig(&_PoolUtils.TransactOpts, _staderConfig)
}

// UpdateStaderConfig is a paid mutator transaction binding the contract method 0x9ee804cb.
//
// Solidity: function updateStaderConfig(address _staderConfig) returns()
func (_PoolUtils *PoolUtilsTransactorSession) UpdateStaderConfig(_staderConfig common.Address) (*types.Transaction, error) {
	return _PoolUtils.Contract.UpdateStaderConfig(&_PoolUtils.TransactOpts, _staderConfig)
}

// PoolUtilsDeactivatedPoolIterator is returned from FilterDeactivatedPool and is used to iterate over the raw logs and unpacked data for DeactivatedPool events raised by the PoolUtils contract.
type PoolUtilsDeactivatedPoolIterator struct {
	Event *PoolUtilsDeactivatedPool // Event containing the contract specifics and raw log

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
func (it *PoolUtilsDeactivatedPoolIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PoolUtilsDeactivatedPool)
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
		it.Event = new(PoolUtilsDeactivatedPool)
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
func (it *PoolUtilsDeactivatedPoolIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PoolUtilsDeactivatedPoolIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PoolUtilsDeactivatedPool represents a DeactivatedPool event raised by the PoolUtils contract.
type PoolUtilsDeactivatedPool struct {
	PoolId      uint8
	PoolAddress common.Address
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterDeactivatedPool is a free log retrieval operation binding the contract event 0xf711845001a9a7fade2a40e12b1fb02c31952a41f7c999ae1f84a283d32671f6.
//
// Solidity: event DeactivatedPool(uint8 indexed poolId, address poolAddress)
func (_PoolUtils *PoolUtilsFilterer) FilterDeactivatedPool(opts *bind.FilterOpts, poolId []uint8) (*PoolUtilsDeactivatedPoolIterator, error) {

	var poolIdRule []interface{}
	for _, poolIdItem := range poolId {
		poolIdRule = append(poolIdRule, poolIdItem)
	}

	logs, sub, err := _PoolUtils.contract.FilterLogs(opts, "DeactivatedPool", poolIdRule)
	if err != nil {
		return nil, err
	}
	return &PoolUtilsDeactivatedPoolIterator{contract: _PoolUtils.contract, event: "DeactivatedPool", logs: logs, sub: sub}, nil
}

// WatchDeactivatedPool is a free log subscription operation binding the contract event 0xf711845001a9a7fade2a40e12b1fb02c31952a41f7c999ae1f84a283d32671f6.
//
// Solidity: event DeactivatedPool(uint8 indexed poolId, address poolAddress)
func (_PoolUtils *PoolUtilsFilterer) WatchDeactivatedPool(opts *bind.WatchOpts, sink chan<- *PoolUtilsDeactivatedPool, poolId []uint8) (event.Subscription, error) {

	var poolIdRule []interface{}
	for _, poolIdItem := range poolId {
		poolIdRule = append(poolIdRule, poolIdItem)
	}

	logs, sub, err := _PoolUtils.contract.WatchLogs(opts, "DeactivatedPool", poolIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PoolUtilsDeactivatedPool)
				if err := _PoolUtils.contract.UnpackLog(event, "DeactivatedPool", log); err != nil {
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

// ParseDeactivatedPool is a log parse operation binding the contract event 0xf711845001a9a7fade2a40e12b1fb02c31952a41f7c999ae1f84a283d32671f6.
//
// Solidity: event DeactivatedPool(uint8 indexed poolId, address poolAddress)
func (_PoolUtils *PoolUtilsFilterer) ParseDeactivatedPool(log types.Log) (*PoolUtilsDeactivatedPool, error) {
	event := new(PoolUtilsDeactivatedPool)
	if err := _PoolUtils.contract.UnpackLog(event, "DeactivatedPool", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PoolUtilsExitValidatorIterator is returned from FilterExitValidator and is used to iterate over the raw logs and unpacked data for ExitValidator events raised by the PoolUtils contract.
type PoolUtilsExitValidatorIterator struct {
	Event *PoolUtilsExitValidator // Event containing the contract specifics and raw log

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
func (it *PoolUtilsExitValidatorIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PoolUtilsExitValidator)
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
		it.Event = new(PoolUtilsExitValidator)
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
func (it *PoolUtilsExitValidatorIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PoolUtilsExitValidatorIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PoolUtilsExitValidator represents a ExitValidator event raised by the PoolUtils contract.
type PoolUtilsExitValidator struct {
	Pubkey []byte
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterExitValidator is a free log retrieval operation binding the contract event 0xce4931e23262c5aa14c3b95b5e67c07bb38447fda706a4e5e4019e4f70142812.
//
// Solidity: event ExitValidator(bytes pubkey)
func (_PoolUtils *PoolUtilsFilterer) FilterExitValidator(opts *bind.FilterOpts) (*PoolUtilsExitValidatorIterator, error) {

	logs, sub, err := _PoolUtils.contract.FilterLogs(opts, "ExitValidator")
	if err != nil {
		return nil, err
	}
	return &PoolUtilsExitValidatorIterator{contract: _PoolUtils.contract, event: "ExitValidator", logs: logs, sub: sub}, nil
}

// WatchExitValidator is a free log subscription operation binding the contract event 0xce4931e23262c5aa14c3b95b5e67c07bb38447fda706a4e5e4019e4f70142812.
//
// Solidity: event ExitValidator(bytes pubkey)
func (_PoolUtils *PoolUtilsFilterer) WatchExitValidator(opts *bind.WatchOpts, sink chan<- *PoolUtilsExitValidator) (event.Subscription, error) {

	logs, sub, err := _PoolUtils.contract.WatchLogs(opts, "ExitValidator")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PoolUtilsExitValidator)
				if err := _PoolUtils.contract.UnpackLog(event, "ExitValidator", log); err != nil {
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

// ParseExitValidator is a log parse operation binding the contract event 0xce4931e23262c5aa14c3b95b5e67c07bb38447fda706a4e5e4019e4f70142812.
//
// Solidity: event ExitValidator(bytes pubkey)
func (_PoolUtils *PoolUtilsFilterer) ParseExitValidator(log types.Log) (*PoolUtilsExitValidator, error) {
	event := new(PoolUtilsExitValidator)
	if err := _PoolUtils.contract.UnpackLog(event, "ExitValidator", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PoolUtilsInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the PoolUtils contract.
type PoolUtilsInitializedIterator struct {
	Event *PoolUtilsInitialized // Event containing the contract specifics and raw log

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
func (it *PoolUtilsInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PoolUtilsInitialized)
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
		it.Event = new(PoolUtilsInitialized)
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
func (it *PoolUtilsInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PoolUtilsInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PoolUtilsInitialized represents a Initialized event raised by the PoolUtils contract.
type PoolUtilsInitialized struct {
	Version uint8
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_PoolUtils *PoolUtilsFilterer) FilterInitialized(opts *bind.FilterOpts) (*PoolUtilsInitializedIterator, error) {

	logs, sub, err := _PoolUtils.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &PoolUtilsInitializedIterator{contract: _PoolUtils.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_PoolUtils *PoolUtilsFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *PoolUtilsInitialized) (event.Subscription, error) {

	logs, sub, err := _PoolUtils.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PoolUtilsInitialized)
				if err := _PoolUtils.contract.UnpackLog(event, "Initialized", log); err != nil {
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
func (_PoolUtils *PoolUtilsFilterer) ParseInitialized(log types.Log) (*PoolUtilsInitialized, error) {
	event := new(PoolUtilsInitialized)
	if err := _PoolUtils.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PoolUtilsPoolAddedIterator is returned from FilterPoolAdded and is used to iterate over the raw logs and unpacked data for PoolAdded events raised by the PoolUtils contract.
type PoolUtilsPoolAddedIterator struct {
	Event *PoolUtilsPoolAdded // Event containing the contract specifics and raw log

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
func (it *PoolUtilsPoolAddedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PoolUtilsPoolAdded)
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
		it.Event = new(PoolUtilsPoolAdded)
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
func (it *PoolUtilsPoolAddedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PoolUtilsPoolAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PoolUtilsPoolAdded represents a PoolAdded event raised by the PoolUtils contract.
type PoolUtilsPoolAdded struct {
	PoolId      uint8
	PoolAddress common.Address
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterPoolAdded is a free log retrieval operation binding the contract event 0x697362d5a2939aff718fb2db4145cb1b4ffc68872c82b2e64d805d8e94845af1.
//
// Solidity: event PoolAdded(uint8 indexed poolId, address poolAddress)
func (_PoolUtils *PoolUtilsFilterer) FilterPoolAdded(opts *bind.FilterOpts, poolId []uint8) (*PoolUtilsPoolAddedIterator, error) {

	var poolIdRule []interface{}
	for _, poolIdItem := range poolId {
		poolIdRule = append(poolIdRule, poolIdItem)
	}

	logs, sub, err := _PoolUtils.contract.FilterLogs(opts, "PoolAdded", poolIdRule)
	if err != nil {
		return nil, err
	}
	return &PoolUtilsPoolAddedIterator{contract: _PoolUtils.contract, event: "PoolAdded", logs: logs, sub: sub}, nil
}

// WatchPoolAdded is a free log subscription operation binding the contract event 0x697362d5a2939aff718fb2db4145cb1b4ffc68872c82b2e64d805d8e94845af1.
//
// Solidity: event PoolAdded(uint8 indexed poolId, address poolAddress)
func (_PoolUtils *PoolUtilsFilterer) WatchPoolAdded(opts *bind.WatchOpts, sink chan<- *PoolUtilsPoolAdded, poolId []uint8) (event.Subscription, error) {

	var poolIdRule []interface{}
	for _, poolIdItem := range poolId {
		poolIdRule = append(poolIdRule, poolIdItem)
	}

	logs, sub, err := _PoolUtils.contract.WatchLogs(opts, "PoolAdded", poolIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PoolUtilsPoolAdded)
				if err := _PoolUtils.contract.UnpackLog(event, "PoolAdded", log); err != nil {
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

// ParsePoolAdded is a log parse operation binding the contract event 0x697362d5a2939aff718fb2db4145cb1b4ffc68872c82b2e64d805d8e94845af1.
//
// Solidity: event PoolAdded(uint8 indexed poolId, address poolAddress)
func (_PoolUtils *PoolUtilsFilterer) ParsePoolAdded(log types.Log) (*PoolUtilsPoolAdded, error) {
	event := new(PoolUtilsPoolAdded)
	if err := _PoolUtils.contract.UnpackLog(event, "PoolAdded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PoolUtilsPoolAddressUpdatedIterator is returned from FilterPoolAddressUpdated and is used to iterate over the raw logs and unpacked data for PoolAddressUpdated events raised by the PoolUtils contract.
type PoolUtilsPoolAddressUpdatedIterator struct {
	Event *PoolUtilsPoolAddressUpdated // Event containing the contract specifics and raw log

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
func (it *PoolUtilsPoolAddressUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PoolUtilsPoolAddressUpdated)
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
		it.Event = new(PoolUtilsPoolAddressUpdated)
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
func (it *PoolUtilsPoolAddressUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PoolUtilsPoolAddressUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PoolUtilsPoolAddressUpdated represents a PoolAddressUpdated event raised by the PoolUtils contract.
type PoolUtilsPoolAddressUpdated struct {
	PoolId      uint8
	PoolAddress common.Address
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterPoolAddressUpdated is a free log retrieval operation binding the contract event 0xf732deab68331ad20834cfc15d686fed4bac945cf3af5d7f729205c9bf846199.
//
// Solidity: event PoolAddressUpdated(uint8 indexed poolId, address poolAddress)
func (_PoolUtils *PoolUtilsFilterer) FilterPoolAddressUpdated(opts *bind.FilterOpts, poolId []uint8) (*PoolUtilsPoolAddressUpdatedIterator, error) {

	var poolIdRule []interface{}
	for _, poolIdItem := range poolId {
		poolIdRule = append(poolIdRule, poolIdItem)
	}

	logs, sub, err := _PoolUtils.contract.FilterLogs(opts, "PoolAddressUpdated", poolIdRule)
	if err != nil {
		return nil, err
	}
	return &PoolUtilsPoolAddressUpdatedIterator{contract: _PoolUtils.contract, event: "PoolAddressUpdated", logs: logs, sub: sub}, nil
}

// WatchPoolAddressUpdated is a free log subscription operation binding the contract event 0xf732deab68331ad20834cfc15d686fed4bac945cf3af5d7f729205c9bf846199.
//
// Solidity: event PoolAddressUpdated(uint8 indexed poolId, address poolAddress)
func (_PoolUtils *PoolUtilsFilterer) WatchPoolAddressUpdated(opts *bind.WatchOpts, sink chan<- *PoolUtilsPoolAddressUpdated, poolId []uint8) (event.Subscription, error) {

	var poolIdRule []interface{}
	for _, poolIdItem := range poolId {
		poolIdRule = append(poolIdRule, poolIdItem)
	}

	logs, sub, err := _PoolUtils.contract.WatchLogs(opts, "PoolAddressUpdated", poolIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PoolUtilsPoolAddressUpdated)
				if err := _PoolUtils.contract.UnpackLog(event, "PoolAddressUpdated", log); err != nil {
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

// ParsePoolAddressUpdated is a log parse operation binding the contract event 0xf732deab68331ad20834cfc15d686fed4bac945cf3af5d7f729205c9bf846199.
//
// Solidity: event PoolAddressUpdated(uint8 indexed poolId, address poolAddress)
func (_PoolUtils *PoolUtilsFilterer) ParsePoolAddressUpdated(log types.Log) (*PoolUtilsPoolAddressUpdated, error) {
	event := new(PoolUtilsPoolAddressUpdated)
	if err := _PoolUtils.contract.UnpackLog(event, "PoolAddressUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PoolUtilsRoleAdminChangedIterator is returned from FilterRoleAdminChanged and is used to iterate over the raw logs and unpacked data for RoleAdminChanged events raised by the PoolUtils contract.
type PoolUtilsRoleAdminChangedIterator struct {
	Event *PoolUtilsRoleAdminChanged // Event containing the contract specifics and raw log

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
func (it *PoolUtilsRoleAdminChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PoolUtilsRoleAdminChanged)
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
		it.Event = new(PoolUtilsRoleAdminChanged)
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
func (it *PoolUtilsRoleAdminChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PoolUtilsRoleAdminChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PoolUtilsRoleAdminChanged represents a RoleAdminChanged event raised by the PoolUtils contract.
type PoolUtilsRoleAdminChanged struct {
	Role              [32]byte
	PreviousAdminRole [32]byte
	NewAdminRole      [32]byte
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterRoleAdminChanged is a free log retrieval operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_PoolUtils *PoolUtilsFilterer) FilterRoleAdminChanged(opts *bind.FilterOpts, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (*PoolUtilsRoleAdminChangedIterator, error) {

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

	logs, sub, err := _PoolUtils.contract.FilterLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return &PoolUtilsRoleAdminChangedIterator{contract: _PoolUtils.contract, event: "RoleAdminChanged", logs: logs, sub: sub}, nil
}

// WatchRoleAdminChanged is a free log subscription operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_PoolUtils *PoolUtilsFilterer) WatchRoleAdminChanged(opts *bind.WatchOpts, sink chan<- *PoolUtilsRoleAdminChanged, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (event.Subscription, error) {

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

	logs, sub, err := _PoolUtils.contract.WatchLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PoolUtilsRoleAdminChanged)
				if err := _PoolUtils.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
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
func (_PoolUtils *PoolUtilsFilterer) ParseRoleAdminChanged(log types.Log) (*PoolUtilsRoleAdminChanged, error) {
	event := new(PoolUtilsRoleAdminChanged)
	if err := _PoolUtils.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PoolUtilsRoleGrantedIterator is returned from FilterRoleGranted and is used to iterate over the raw logs and unpacked data for RoleGranted events raised by the PoolUtils contract.
type PoolUtilsRoleGrantedIterator struct {
	Event *PoolUtilsRoleGranted // Event containing the contract specifics and raw log

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
func (it *PoolUtilsRoleGrantedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PoolUtilsRoleGranted)
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
		it.Event = new(PoolUtilsRoleGranted)
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
func (it *PoolUtilsRoleGrantedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PoolUtilsRoleGrantedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PoolUtilsRoleGranted represents a RoleGranted event raised by the PoolUtils contract.
type PoolUtilsRoleGranted struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleGranted is a free log retrieval operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_PoolUtils *PoolUtilsFilterer) FilterRoleGranted(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*PoolUtilsRoleGrantedIterator, error) {

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

	logs, sub, err := _PoolUtils.contract.FilterLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &PoolUtilsRoleGrantedIterator{contract: _PoolUtils.contract, event: "RoleGranted", logs: logs, sub: sub}, nil
}

// WatchRoleGranted is a free log subscription operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_PoolUtils *PoolUtilsFilterer) WatchRoleGranted(opts *bind.WatchOpts, sink chan<- *PoolUtilsRoleGranted, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _PoolUtils.contract.WatchLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PoolUtilsRoleGranted)
				if err := _PoolUtils.contract.UnpackLog(event, "RoleGranted", log); err != nil {
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
func (_PoolUtils *PoolUtilsFilterer) ParseRoleGranted(log types.Log) (*PoolUtilsRoleGranted, error) {
	event := new(PoolUtilsRoleGranted)
	if err := _PoolUtils.contract.UnpackLog(event, "RoleGranted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PoolUtilsRoleRevokedIterator is returned from FilterRoleRevoked and is used to iterate over the raw logs and unpacked data for RoleRevoked events raised by the PoolUtils contract.
type PoolUtilsRoleRevokedIterator struct {
	Event *PoolUtilsRoleRevoked // Event containing the contract specifics and raw log

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
func (it *PoolUtilsRoleRevokedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PoolUtilsRoleRevoked)
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
		it.Event = new(PoolUtilsRoleRevoked)
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
func (it *PoolUtilsRoleRevokedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PoolUtilsRoleRevokedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PoolUtilsRoleRevoked represents a RoleRevoked event raised by the PoolUtils contract.
type PoolUtilsRoleRevoked struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleRevoked is a free log retrieval operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_PoolUtils *PoolUtilsFilterer) FilterRoleRevoked(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*PoolUtilsRoleRevokedIterator, error) {

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

	logs, sub, err := _PoolUtils.contract.FilterLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &PoolUtilsRoleRevokedIterator{contract: _PoolUtils.contract, event: "RoleRevoked", logs: logs, sub: sub}, nil
}

// WatchRoleRevoked is a free log subscription operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_PoolUtils *PoolUtilsFilterer) WatchRoleRevoked(opts *bind.WatchOpts, sink chan<- *PoolUtilsRoleRevoked, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _PoolUtils.contract.WatchLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PoolUtilsRoleRevoked)
				if err := _PoolUtils.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
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
func (_PoolUtils *PoolUtilsFilterer) ParseRoleRevoked(log types.Log) (*PoolUtilsRoleRevoked, error) {
	event := new(PoolUtilsRoleRevoked)
	if err := _PoolUtils.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PoolUtilsUpdatedStaderConfigIterator is returned from FilterUpdatedStaderConfig and is used to iterate over the raw logs and unpacked data for UpdatedStaderConfig events raised by the PoolUtils contract.
type PoolUtilsUpdatedStaderConfigIterator struct {
	Event *PoolUtilsUpdatedStaderConfig // Event containing the contract specifics and raw log

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
func (it *PoolUtilsUpdatedStaderConfigIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PoolUtilsUpdatedStaderConfig)
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
		it.Event = new(PoolUtilsUpdatedStaderConfig)
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
func (it *PoolUtilsUpdatedStaderConfigIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PoolUtilsUpdatedStaderConfigIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PoolUtilsUpdatedStaderConfig represents a UpdatedStaderConfig event raised by the PoolUtils contract.
type PoolUtilsUpdatedStaderConfig struct {
	StaderConfig common.Address
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterUpdatedStaderConfig is a free log retrieval operation binding the contract event 0xdb2219043d7b197cb235f1af0cf6d782d77dee3de19e3f4fb6d39aae633b4485.
//
// Solidity: event UpdatedStaderConfig(address staderConfig)
func (_PoolUtils *PoolUtilsFilterer) FilterUpdatedStaderConfig(opts *bind.FilterOpts) (*PoolUtilsUpdatedStaderConfigIterator, error) {

	logs, sub, err := _PoolUtils.contract.FilterLogs(opts, "UpdatedStaderConfig")
	if err != nil {
		return nil, err
	}
	return &PoolUtilsUpdatedStaderConfigIterator{contract: _PoolUtils.contract, event: "UpdatedStaderConfig", logs: logs, sub: sub}, nil
}

// WatchUpdatedStaderConfig is a free log subscription operation binding the contract event 0xdb2219043d7b197cb235f1af0cf6d782d77dee3de19e3f4fb6d39aae633b4485.
//
// Solidity: event UpdatedStaderConfig(address staderConfig)
func (_PoolUtils *PoolUtilsFilterer) WatchUpdatedStaderConfig(opts *bind.WatchOpts, sink chan<- *PoolUtilsUpdatedStaderConfig) (event.Subscription, error) {

	logs, sub, err := _PoolUtils.contract.WatchLogs(opts, "UpdatedStaderConfig")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PoolUtilsUpdatedStaderConfig)
				if err := _PoolUtils.contract.UnpackLog(event, "UpdatedStaderConfig", log); err != nil {
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
func (_PoolUtils *PoolUtilsFilterer) ParseUpdatedStaderConfig(log types.Log) (*PoolUtilsUpdatedStaderConfig, error) {
	event := new(PoolUtilsUpdatedStaderConfig)
	if err := _PoolUtils.contract.UnpackLog(event, "UpdatedStaderConfig", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
