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

func deployContracts(t *testing.T, c *cli.Context, eth1URL string) {
	client, err := ethclient.Dial(eth1URL)
	require.Nil(t, err)

	// private key of the deployer
	privateKey, err := crypto.HexToECDSA(preFundedKeyAnvil)
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

	// deploy the ethx contract
	ethXAddr, _, ethxContract, err := contracts.DeployETHX(auth, client, fromAddress, fromAddress)
	require.Nil(t, err)
	auth, err = GetNextTransaction(client, fromAddress, privateKey, chainID)
	require.Nil(t, err)

	StaderConfigInETHX, _ := ethxContract.StaderConfig(&bind.CallOpts{Pending: true})
	fmt.Printf(" StaderConfig %+v", StaderConfigInETHX.Hex())

	// deploy the config contract
	staderCfAddress, _, stdCfContract, err := contracts.DeployStaderConfig(auth, client, fromAddress)
	require.Nil(t, err)
	auth, _ = GetNextTransaction(client, fromAddress, privateKey, chainID)

	// // Init ethx
	_, err = ethxContract.Initialize(auth, fromAddress, staderCfAddress)
	require.Nil(t, err)
	auth, err = GetNextTransaction(client, fromAddress, privateKey, chainID)
	require.Nil(t, err)

	admin, _ := stdCfContract.GetAdmin(&bind.CallOpts{Pending: true})
	fmt.Printf("ADMIN %+v", admin.Hex())
	// auth, err = GetNextTransaction(client, fromAddress, privateKey, chainID)
	// require.Nil(t, err)

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

	// client.Commit()
	require.Nil(t, err)
	auth, err = GetNextTransaction(client, fromAddress, privateKey, chainID)
	require.Nil(t, err)

	// Update Node permissionless pool to stader config
	_, err = stdCfContract.UpdatePermissionlessPool(auth, permissionlessPool)
	require.Nil(t, err)
	auth, err = GetNextTransaction(client, fromAddress, privateKey, chainID)
	require.Nil(t, err)

	time.Sleep(time.Second)
	opt := bind.CallOpts{
		Pending: true,
	}
	ETHxToken, err := ethxContract.Decimals(&opt)
	if err != nil {
		panic(err)
	}

	fmt.Printf("ETHxToken from %+v\n", ETHxToken)
	// time.Sleep(time.Minute)
	storedPLPool, err := stdCfContract.GetPermissionlessPool(&opt)
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

	fmt.Println(" >>>>>>>>>>>>>> Nouce ", nonce)
	// sign the transaction
	auth, err := bind.NewKeyedTransactorWithChainID(privateKey, chainID)
	if err != nil {
		return nil, err
	}
	auth.Nonce = big.NewInt(int64(nonce))
	auth.Value = big.NewInt(0)              // in wei
	auth.GasLimit = uint64(10000000)        // in units
	auth.GasPrice = big.NewInt(10000000000) // in wei

	return auth, nil
}
