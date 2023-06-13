package testing

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	api "github.com/stader-labs/stader-node/stader-lib/api"
)

var PRE_FUNDED_ACCOUNTS = "ef5177cd0b6b21c87db5a0bf35d4084a8a57a9d6a064f86d51ac85f2b873a4e2"

func deployContracts(eth1URL string) {
	client, err := ethclient.Dial(eth1URL)
	if err != nil {
		panic(err)
	}

	// private key of the deployer
	privateKey, err := crypto.HexToECDSA(PRE_FUNDED_ACCOUNTS)
	if err != nil {
		panic(err)
	}

	// extract public key of the deployer from private key
	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		panic("cannot assert type: publicKey is not of type *ecdsa.PublicKey")
	}

	// address of the deployer
	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)

	// chain id of the network
	chainID, err := client.NetworkID(context.Background())
	if err != nil {
		panic(err)
	}

	fmt.Printf("FROM: %+v \n ", fromAddress.Hex())
	// Get Transaction Ops to make a valid Ethereum transaction
	auth, err := GetNextTransaction(client, fromAddress, privateKey, chainID)
	if err != nil {
		panic(err)
	}

	// deploy the contract
	address, tx, _, err := api.DeployETHX(auth, client)
	if err != nil {
		panic(err)
	}

	fmt.Printf("Api contract from %s\n", auth.From.Hex())
	fmt.Printf("Api contract deployed to %s\n", address.Hex())
	fmt.Printf("Tx: %s\n", tx.Hash().Hex())

	// // Get Favorite Number
	// // Call SimpleStorage contract Retrieve function to get current favorite number
	// favoriteNumber, err := simpleStorageApi.Retrieve(&bind.CallOpts{})
	// if err != nil {
	// 	panic(err)
	// }
	// fmt.Printf("Favorite Number: %d\n", favoriteNumber)

	// // Set Favorite Number
	// // Get Transaction Ops to make a valid Ethereum transaction
	// auth, err = GetNextTransaction(client, fromAddress, privateKey, chainID)
	// if err != nil {
	// 	panic(err)
	// }

	// // Call SimpleStorate Store function to store favorite number
	// reply, err := simpleStorageApi.Store(auth, big.NewInt(20))
	// if err != nil {
	// 	panic(err)
	// }
	// fmt.Printf("Reply: %s\n", reply.Hash().Hex())

	// // Get Favorite Number
	// // P.S. Retrieve is a Gas Free function hence no need to get Transaction Ops
	// newfavoriteNumber, err := simpleStorageApi.Retrieve(&bind.CallOpts{})
	// if err != nil {
	// 	panic(err)
	// }
	// fmt.Printf("Favorite Number: %d\n", newfavoriteNumber)
}

// GetNextTransaction returns the next transaction in the pending transaction queue
// NOTE: this is not an optimized way
func GetNextTransaction(client *ethclient.Client, fromAddress common.Address, privateKey *ecdsa.PrivateKey, chainID *big.Int) (*bind.TransactOpts, error) {
	// nonce
	nonce, err := client.PendingNonceAt(context.Background(), fromAddress)
	if err != nil {
		return nil, err
	}

	// sign the transaction
	auth, err := bind.NewKeyedTransactorWithChainID(privateKey, chainID)
	if err != nil {
		return nil, err
	}
	auth.Nonce = big.NewInt(int64(nonce))
	auth.Value = big.NewInt(0)             // in wei
	auth.GasLimit = uint64(3000000)        // in units
	auth.GasPrice = big.NewInt(1000000000) // in wei

	return auth, nil
}
