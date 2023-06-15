package testing

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	"math/big"
	"testing"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/stader-labs/stader-node/stader-lib/contracts"
	"github.com/stretchr/testify/require"
	"github.com/urfave/cli"
)

var PRE_FUNDED_ACCOUNTS = "ef5177cd0b6b21c87db5a0bf35d4084a8a57a9d6a064f86d51ac85f2b873a4e2"

func deployContracts(t *testing.T, c *cli.Context, eth1URL string) {
	client, err := ethclient.Dial(eth1URL)
	require.Nil(t, err)

	// private key of the deployer
	privateKey, err := crypto.HexToECDSA(PRE_FUNDED_ACCOUNTS)
	require.Nil(t, err)

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
	require.Nil(t, err)

	fmt.Printf("FROM: %+v \n ", fromAddress.Hex())
	// Get Transaction Ops to make a valid Ethereum transaction
	auth, err := GetNextTransaction(client, fromAddress, privateKey, chainID)
	require.Nil(t, err)

	// deploy the contract
	ethXAddr, _, ethxContract, err := contracts.DeployETHX(auth, client)
	require.Nil(t, err)

	auth, err = GetNextTransaction(client, fromAddress, privateKey, chainID)
	require.Nil(t, err)

	staderCfAddress, _, stdCfContract, err := contracts.DeployStaderConfig(auth, client)
	require.Nil(t, err)

	auth, err = GetNextTransaction(client, fromAddress, privateKey, chainID)
	require.Nil(t, err)

	stdCfContract.Initialize(auth, fromAddress, ethXAddr)
	auth, err = GetNextTransaction(client, fromAddress, privateKey, chainID)
	require.Nil(t, err)

	// Init ethx
	_, err = ethxContract.Initialize(auth, fromAddress, staderCfAddress)
	require.Nil(t, err)
	auth, err = GetNextTransaction(client, fromAddress, privateKey, chainID)
	require.Nil(t, err)

	ethxContract.UpdateStaderConfig(auth, staderCfAddress)
	auth, err = GetNextTransaction(client, fromAddress, privateKey, chainID)
	require.Nil(t, err)

	// sdcfg, err := contractsSrv.GetStaderConfigContract(c)
	// require.Nil(t, err)

	// Deploy node permission
	plNodeRegistryAddr, _, _, err := contracts.DeployPermissionlessNodeRegistry(auth, client)
	require.Nil(t, err)
	auth, err = GetNextTransaction(client, fromAddress, privateKey, chainID)
	require.Nil(t, err)

	// Update Node permissionless regis to stader config
	_, err = stdCfContract.UpdatePermissionlessNodeRegistry(auth, plNodeRegistryAddr)
	require.Nil(t, err)
	auth, err = GetNextTransaction(client, fromAddress, privateKey, chainID)
	require.Nil(t, err)

	// Deploy permissionless pool
	permissionlessPool, _, _, err := contracts.DeployPermissionlessPool(auth, client)
	require.Nil(t, err)
	auth, err = GetNextTransaction(client, fromAddress, privateKey, chainID)
	require.Nil(t, err)

	// Update Node permissionless pool to stader config
	tx, err := stdCfContract.UpdatePermissionlessPool(auth, permissionlessPool)
	require.Nil(t, err)
	auth, err = GetNextTransaction(client, fromAddress, privateKey, chainID)
	require.Nil(t, err)

	for {
		_, isPending, err := client.TransactionByHash(context.TODO(), tx.Hash())
		require.Nil(t, err)
		if isPending != true {
			break
		}
		time.Sleep(10 * time.Second)
	}

	auth, err = GetNextTransaction(client, fromAddress, privateKey, chainID)
	require.Nil(t, err)

	auth, err = GetNextTransaction(client, fromAddress, privateKey, chainID)
	require.Nil(t, err)

	// time.Sleep(time.Minute)
	storedPLPool, err := stdCfContract.GetPermissionlessPool(&bind.CallOpts{})
	if err != nil {
		panic(err)
	}

	storedPLRegis, err := stdCfContract.GetPermissionlessNodeRegistry(&bind.CallOpts{
		Pending: true,
	})
	require.Nil(t, err)

	fmt.Printf("Api contract from %s\n", auth.From.Hex())
	fmt.Printf("DeployETHX to %s\n", ethXAddr.Hex())
	fmt.Printf("DeployStaderConfig to %s\n", staderCfAddress.Hex())
	fmt.Printf("DeployPermissionlessPool to %s <> %s \n", permissionlessPool.Hex(), storedPLPool.Hex())
	fmt.Printf("permissionlessNR to %s <> store %s\n", plNodeRegistryAddr.Hex(), storedPLRegis.Hex())

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
	auth.GasLimit = uint64(3000000)        // in units 1000000
	auth.GasPrice = big.NewInt(1000000000) // in wei

	return auth, nil
}
