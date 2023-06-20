package testing

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	"log"
	"math/big"
	"testing"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/rpc"
	"github.com/stader-labs/stader-node/shared/services"
	"github.com/stader-labs/stader-node/stader-lib/node"
	"github.com/stader-labs/stader-node/testing/contracts"
	"github.com/stretchr/testify/require"
	"github.com/urfave/cli"
)

const (
	staderConfigAddr  = "0x700b6A60ce7EaaEA56F065753d8dcB9653dbAD35"
	ethXPredefineAddr = "0xA15BB66138824a1c7167f5E85b957d04Dd34E468"
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

	// Get Transaction Ops to make a valid Ethereum transaction
	auth, err := GetNextTransaction(client, fromAddress, privateKey, chainID)
	require.Nil(t, err)

	// deploy the config contract
	staderCfAddress, _, stdCfContract, err := contracts.DeployStaderConfig(auth, client, fromAddress, common.HexToAddress(ethXPredefineAddr))
	require.Nil(t, err)

	fmt.Printf("Stader config address: %+v \n", staderCfAddress.Hex())
	auth, _ = GetNextTransaction(client, fromAddress, privateKey, chainID)

	// Update Node permissionless regis to stader config
	mn, _ := stdCfContract.MANAGER(&bind.CallOpts{})
	_, err = stdCfContract.GrantRole(auth, mn, fromAddress)
	require.Nil(t, err)
	auth, _ = GetNextTransaction(client, fromAddress, privateKey, chainID)

	// am, _ := stdCfContract.ADMIN(&bind.CallOpts{})
	_, err = stdCfContract.UpdateAdmin(auth, fromAddress)
	require.Nil(t, err)
	auth, _ = GetNextTransaction(client, fromAddress, privateKey, chainID)

	// deploy the ethx contract
	ethXAddr, _, ethxContract, err := contracts.DeployETHX(auth, client, fromAddress, staderCfAddress)
	require.Nil(t, err)
	fmt.Printf("EthXAddr %+v", ethXAddr.Hex())
	auth, _ = GetNextTransaction(client, fromAddress, privateKey, chainID)

	ethxContract.UpdateStaderConfig(auth, staderCfAddress)
	auth, _ = GetNextTransaction(client, fromAddress, privateKey, chainID)

	// Deploy node permission regis
	plNodeRegistryAddr, _, nrContact, err := contracts.DeployPermissionlessNodeRegistry(auth, client, fromAddress, staderCfAddress)
	require.Nil(t, err)
	auth, _ = GetNextTransaction(client, fromAddress, privateKey, chainID)

	_, err = stdCfContract.UpdatePermissionlessNodeRegistry(auth, plNodeRegistryAddr)
	require.Nil(t, err)
	auth, _ = GetNextTransaction(client, fromAddress, privateKey, chainID)

	// Permissionless pool
	permissionlessPoolAddr, _, _, err := contracts.DeployPermissionlessPool(auth, client, fromAddress, staderCfAddress)
	require.Nil(t, err)
	auth, _ = GetNextTransaction(client, fromAddress, privateKey, chainID)

	_, err = stdCfContract.UpdatePermissionlessPool(auth, permissionlessPoolAddr)
	require.Nil(t, err)
	auth, _ = GetNextTransaction(client, fromAddress, privateKey, chainID)

	// PoolUtils
	poolUtils, _, poolUtilsContract, err := contracts.DeployPoolUtils(auth, client, fromAddress, staderCfAddress)
	require.Nil(t, err)
	auth, _ = GetNextTransaction(client, fromAddress, privateKey, chainID)

	poolUtilsContract.AddNewPool(auth, uint8(1), permissionlessPoolAddr)
	auth, _ = GetNextTransaction(client, fromAddress, privateKey, chainID)

	x, err := poolUtilsContract.GetPoolIdArray(&bind.CallOpts{})
	fmt.Printf("X %+v", x)

	require.Nil(t, err)
	_, err = stdCfContract.UpdatePoolUtils(auth, poolUtils)
	require.Nil(t, err)
	auth, _ = GetNextTransaction(client, fromAddress, privateKey, chainID)

	// Vault factory
	vaultFactoryAddr, _, vaultFactoryContract, err := contracts.DeployVaultFactory(
		auth,
		client,
		fromAddress,
		staderCfAddress,
	)
	require.Nil(t, err)
	auth, _ = GetNextTransaction(client, fromAddress, privateKey, chainID)

	nrcRole, err := vaultFactoryContract.NODEREGISTRYCONTRACT(&bind.CallOpts{})
	require.Nil(t, err)
	_, err = vaultFactoryContract.GrantRole(auth, nrcRole, plNodeRegistryAddr)
	require.Nil(t, err)
	auth, _ = GetNextTransaction(client, fromAddress, privateKey, chainID)

	_, err = stdCfContract.UpdateVaultFactory(auth, vaultFactoryAddr)
	require.Nil(t, err)
	auth, _ = GetNextTransaction(client, fromAddress, privateKey, chainID)

	vaultStore, err := stdCfContract.GetVaultFactory(&bind.CallOpts{})
	require.Nil(t, err)
	require.Equal(t, vaultStore, vaultFactoryAddr)

	// SocializingPool
	spAddr, _, _, err := contracts.DeploySocializingPool(auth, client, fromAddress, staderCfAddress)
	require.Nil(t, err)

	auth, _ = GetNextTransaction(client, fromAddress, privateKey, chainID)
	_, err = stdCfContract.UpdatePermissionedSocializingPool(auth, spAddr)
	require.Nil(t, err)
	auth, _ = GetNextTransaction(client, fromAddress, privateKey, chainID)

	fmt.Printf("Api contract from %s\n", auth.From.Hex())
	fmt.Printf("DeployETHX to %s\n", ethXAddr.Hex())
	fmt.Printf("DeployStaderConfig to %s\n", staderCfAddress.Hex())
	fmt.Printf("DeployPoolUtils to %s\n", poolUtils.Hex())
	fmt.Printf("PermissionlessPoolAddr %s\n", permissionlessPoolAddr.Hex())
	fmt.Printf("PLNodeRegistryAddr %s\n", plNodeRegistryAddr.Hex())

	prn, err := services.GetPermissionlessNodeRegistry(c)
	require.Nil(t, err)

	w, err := services.GetWallet(c)
	require.Nil(t, err)

	nodePrivateKey, err := w.GetNodePrivateKey()
	require.Nil(t, err)

	acc, err := w.GetNodeAccount()
	require.Nil(t, err)

	send1EthTransaction(client, fromAddress, acc.Address, privateKey, chainID)

	auth, err = GetNextTransaction(client, acc.Address, nodePrivateKey, chainID)
	require.Nil(t, err)

	_, err = node.OnboardNodeOperator(prn, true, "nodetesting", acc.Address, auth)
	require.Nil(t, err)

	exist, err := nrContact.IsExistingOperator(&bind.CallOpts{}, acc.Address)
	require.Nil(t, err)
	require.True(t, exist)
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
	auth.Value = big.NewInt(0)              // in wei
	auth.GasLimit = uint64(10000000)        // in units
	auth.GasPrice = big.NewInt(10000000000) // in wei

	return auth, nil
}

func send1EthTransaction(client *ethclient.Client,
	fromAddress common.Address,
	toAddress common.Address,
	privateKey *ecdsa.PrivateKey,
	chainID *big.Int,
) {
	// nonce
	nonce, err := client.PendingNonceAt(context.Background(), fromAddress)
	if err != nil {

		log.Fatal(err)
	}

	gasLimit := uint64(10000000)        // in units
	gasPrice := big.NewInt(10000000000) // in wei

	var data []byte
	tx := types.NewTransaction(
		nonce,
		toAddress,
		big.NewInt(9000000000000000000),
		gasLimit,
		gasPrice,
		data)

	signedTx, err := types.SignTx(tx, types.NewEIP155Signer(chainID), privateKey)
	if err != nil {
		log.Fatal(err)
	}

	err = client.SendTransaction(context.Background(), signedTx)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("tx sent: %s", signedTx.Hash().Hex()) // tx sent: 0x77006fcb3938f648e2cc65bafd27dec30b9bfbe9df41f78498b9c8b7322a249e
}

func reset(t *testing.T, rawurl string) {
	ctx := context.Background()
	c, err := rpc.DialContext(ctx, rawurl)
	require.Nil(t, err)

	err = c.CallContext(ctx, nil, "anvil_reset")
	require.Nil(t, err)
}
