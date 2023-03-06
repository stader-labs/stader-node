package protocol

import (
	"fmt"
	"math/big"
	"sync"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"

	protocoldao "github.com/stader-labs/stader-node/stader-lib/dao/protocol"
	"github.com/stader-labs/stader-node/stader-lib/stader"
	"github.com/stader-labs/stader-node/stader-lib/utils/eth"
)

// Config
const AuctionSettingsContractName = "rocketDAOProtocolSettingsAuction"

// Lot creation currently enabled
func GetCreateLotEnabled(rp *stader.PermissionlessNodeRegistryContractManager, opts *bind.CallOpts) (bool, error) {
	auctionSettingsContract, err := getAuctionSettingsContract(rp, opts)
	if err != nil {
		return false, err
	}
	value := new(bool)
	if err := auctionSettingsContract.Call(opts, value, "getCreateLotEnabled"); err != nil {
		return false, fmt.Errorf("Could not get lot creation enabled status: %w", err)
	}
	return *value, nil
}
func BootstrapCreateLotEnabled(rp *stader.PermissionlessNodeRegistryContractManager, value bool, opts *bind.TransactOpts) (common.Hash, error) {
	return protocoldao.BootstrapBool(rp, AuctionSettingsContractName, "auction.lot.create.enabled", value, opts)
}

// Lot bidding currently enabled
func GetBidOnLotEnabled(rp *stader.PermissionlessNodeRegistryContractManager, opts *bind.CallOpts) (bool, error) {
	auctionSettingsContract, err := getAuctionSettingsContract(rp, opts)
	if err != nil {
		return false, err
	}
	value := new(bool)
	if err := auctionSettingsContract.Call(opts, value, "getBidOnLotEnabled"); err != nil {
		return false, fmt.Errorf("Could not get lot bidding enabled status: %w", err)
	}
	return *value, nil
}
func BootstrapBidOnLotEnabled(rp *stader.PermissionlessNodeRegistryContractManager, value bool, opts *bind.TransactOpts) (common.Hash, error) {
	return protocoldao.BootstrapBool(rp, AuctionSettingsContractName, "auction.lot.bidding.enabled", value, opts)
}

// The minimum lot size in ETH value
func GetLotMinimumEthValue(rp *stader.PermissionlessNodeRegistryContractManager, opts *bind.CallOpts) (*big.Int, error) {
	auctionSettingsContract, err := getAuctionSettingsContract(rp, opts)
	if err != nil {
		return nil, err
	}
	value := new(*big.Int)
	if err := auctionSettingsContract.Call(opts, value, "getLotMinimumEthValue"); err != nil {
		return nil, fmt.Errorf("Could not get lot minimum ETH value: %w", err)
	}
	return *value, nil
}
func BootstrapLotMinimumEthValue(rp *stader.PermissionlessNodeRegistryContractManager, value *big.Int, opts *bind.TransactOpts) (common.Hash, error) {
	return protocoldao.BootstrapUint(rp, AuctionSettingsContractName, "auction.lot.value.minimum", value, opts)
}

// The maximum lot size in ETH value
func GetLotMaximumEthValue(rp *stader.PermissionlessNodeRegistryContractManager, opts *bind.CallOpts) (*big.Int, error) {
	auctionSettingsContract, err := getAuctionSettingsContract(rp, opts)
	if err != nil {
		return nil, err
	}
	value := new(*big.Int)
	if err := auctionSettingsContract.Call(opts, value, "getLotMaximumEthValue"); err != nil {
		return nil, fmt.Errorf("Could not get lot maximum ETH value: %w", err)
	}
	return *value, nil
}
func BootstrapLotMaximumEthValue(rp *stader.PermissionlessNodeRegistryContractManager, value *big.Int, opts *bind.TransactOpts) (common.Hash, error) {
	return protocoldao.BootstrapUint(rp, AuctionSettingsContractName, "auction.lot.value.maximum", value, opts)
}

// The lot duration in blocks
func GetLotDuration(rp *stader.PermissionlessNodeRegistryContractManager, opts *bind.CallOpts) (uint64, error) {
	auctionSettingsContract, err := getAuctionSettingsContract(rp, opts)
	if err != nil {
		return 0, err
	}
	value := new(*big.Int)
	if err := auctionSettingsContract.Call(opts, value, "getLotDuration"); err != nil {
		return 0, fmt.Errorf("Could not get lot duration: %w", err)
	}
	return (*value).Uint64(), nil
}
func BootstrapLotDuration(rp *stader.PermissionlessNodeRegistryContractManager, value uint64, opts *bind.TransactOpts) (common.Hash, error) {
	return protocoldao.BootstrapUint(rp, AuctionSettingsContractName, "auction.lot.duration", big.NewInt(int64(value)), opts)
}

// The starting price relative to current ETH price, as a fraction
func GetLotStartingPriceRatio(rp *stader.PermissionlessNodeRegistryContractManager, opts *bind.CallOpts) (float64, error) {
	auctionSettingsContract, err := getAuctionSettingsContract(rp, opts)
	if err != nil {
		return 0, err
	}
	value := new(*big.Int)
	if err := auctionSettingsContract.Call(opts, value, "getStartingPriceRatio"); err != nil {
		return 0, fmt.Errorf("Could not get lot starting price ratio: %w", err)
	}
	return eth.WeiToEth(*value), nil
}
func BootstrapLotStartingPriceRatio(rp *stader.PermissionlessNodeRegistryContractManager, value float64, opts *bind.TransactOpts) (common.Hash, error) {
	return protocoldao.BootstrapUint(rp, AuctionSettingsContractName, "auction.price.start", eth.EthToWei(value), opts)
}

// The reserve price relative to current ETH price, as a fraction
func GetLotReservePriceRatio(rp *stader.PermissionlessNodeRegistryContractManager, opts *bind.CallOpts) (float64, error) {
	auctionSettingsContract, err := getAuctionSettingsContract(rp, opts)
	if err != nil {
		return 0, err
	}
	value := new(*big.Int)
	if err := auctionSettingsContract.Call(opts, value, "getReservePriceRatio"); err != nil {
		return 0, fmt.Errorf("Could not get lot reserve price ratio: %w", err)
	}
	return eth.WeiToEth(*value), nil
}
func BootstrapLotReservePriceRatio(rp *stader.PermissionlessNodeRegistryContractManager, value float64, opts *bind.TransactOpts) (common.Hash, error) {
	return protocoldao.BootstrapUint(rp, AuctionSettingsContractName, "auction.price.reserve", eth.EthToWei(value), opts)
}

// Get contracts
var auctionSettingsContractLock sync.Mutex

func getAuctionSettingsContract(rp *stader.PermissionlessNodeRegistryContractManager, opts *bind.CallOpts) (*stader.Contract, error) {
	auctionSettingsContractLock.Lock()
	defer auctionSettingsContractLock.Unlock()
	return nil, nil
	//return rp.GetContract(AuctionSettingsContractName, opts)
}
