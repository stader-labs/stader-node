package node

import (
	"context"
	"fmt"
	_ "github.com/rocket-pool/rocketpool-go/dao/trustednode"
	_ "github.com/rocket-pool/rocketpool-go/network"
	_ "github.com/rocket-pool/rocketpool-go/settings/protocol"
	_ "github.com/rocket-pool/rocketpool-go/settings/trustednode"
	rptypes "github.com/rocket-pool/rocketpool-go/types"
	"github.com/stader-labs/stader-minipool-go/node"
	"github.com/stader-labs/stader-minipool-go/types"
	"github.com/urfave/cli"
	_ "golang.org/x/sync/errgroup"
	"math/big"

	"github.com/stader-labs/stader-node/shared/services"
	"github.com/stader-labs/stader-node/shared/types/api"
	"github.com/stader-labs/stader-node/shared/utils/eth1"
	"github.com/stader-labs/stader-node/shared/utils/validator"
)

func nodeDeposit(c *cli.Context, amountWei *big.Int, salt *big.Int, submit bool) (*api.NodeDepositResponse, error) {

	w, err := services.GetWallet(c)
	if err != nil {
		return nil, err
	}
	ec, err := services.GetEthClient(c)
	if err != nil {
		return nil, err
	}
	sor, err := services.GetStaderOperatorRegistry(c)
	if err != nil {
		return nil, err
	}
	svr, err := services.GetStaderValidatorRegistry(c)
	if err != nil {
		return nil, err
	}
	srcf, err := services.GetStaderRewardContractFactory(c)
	if err != nil {
		return nil, err
	}
	bc, err := services.GetBeaconClient(c)
	if err != nil {
		return nil, err
	}

	// Get eth2 config
	eth2Config, err := bc.GetEth2Config()
	if err != nil {
		return nil, err
	}

	// Get node account
	nodeAccount, err := w.GetNodeAccount()
	if err != nil {
		return nil, err
	}

	// Response
	response := api.NodeDepositResponse{}

	// Adjust the salt
	if salt.Cmp(big.NewInt(0)) == 0 {
		nonce, err := ec.NonceAt(context.Background(), nodeAccount.Address, nil)
		if err != nil {
			return nil, err
		}
		salt.SetUint64(nonce)
	}
	// Get transactor
	opts, err := w.GetNodeAccountTransactor()
	if err != nil {
		return nil, err
	}
	opts.Value = amountWei

	// Create and save a new validator key
	validatorKey, err := w.CreateValidatorKey()
	if err != nil {
		return nil, err
	}

	nodeAccount, err = w.GetNodeAccount()
	// get the vault address and vault credential
	fmt.Printf("node account is %s\n", nodeAccount.Address)
	operatorRegistryInfo, err := node.GetOperatorRegistry(sor, nodeAccount.Address, nil)
	if err != nil {
		return nil, err
	}
	if operatorRegistryInfo.OperatorName == "" {
		return nil, fmt.Errorf("node is not registered")
	}

	fmt.Printf("operator registry info is %v\n", operatorRegistryInfo)
	validatorKeyCount, err := node.GetTotalValidatorKeys(sor, nodeAccount.Address, nil)
	if err != nil {
		return nil, err
	}
	fmt.Printf("validator key count is %v\n", validatorKeyCount)
	rewardWithdrawVault, err := node.ComputeWithdrawVaultAddress(srcf, operatorRegistryInfo.OperatorId, validatorKeyCount, nil)
	if err != nil {
		return nil, err
	}
	fmt.Printf("reward withdraw vault is %v\n", rewardWithdrawVault)
	withdrawCredentials, err := node.GetValidatorWithdrawalCredential(srcf, rewardWithdrawVault, nil)
	if err != nil {
		return nil, err
	}

	fmt.Printf("bharath-build: withdraw creds is %v\n", withdrawCredentials)

	// Get validator deposit data and associated parameters
	depositData, depositDataRoot, err := validator.GetDepositData(validatorKey, withdrawCredentials, eth2Config)
	if err != nil {
		return nil, err
	}
	pubKey := rptypes.BytesToValidatorPubkey(depositData.PublicKey)
	signature := rptypes.BytesToValidatorSignature(depositData.Signature)

	// Make sure a validator with this pubkey doesn't already exist
	status, err := bc.GetValidatorStatus(pubKey, nil)
	if err != nil {
		return nil, fmt.Errorf("Error checking for existing validator status: %w\nYour funds have not been deposited for your own safety.", err)
	}
	if status.Exists {
		return nil, fmt.Errorf("**** ALERT ****\n"+
			"Your minipool %s has the following as a validator pubkey:\n\t%s\n"+
			"This key is already in use by validator %d on the Beacon chain!\n"+
			"Rocket Pool will not allow you to deposit this validator for your own safety so you do not get slashed.\n"+
			"PLEASE REPORT THIS TO THE ROCKET POOL DEVELOPERS.\n"+
			"***************\n", operatorRegistryInfo.OperatorRewardAddress.Hex(), pubKey.Hex(), status.Index)
	}

	// Override the provided pending TX if requested
	err = eth1.CheckForNonceOverride(c, opts)
	if err != nil {
		return nil, fmt.Errorf("Error checking for nonce override: %w", err)
	}

	// Do not send transaction unless requested
	opts.NoSend = !submit

	tx, err := node.AddValidatorKeys(svr, types.ValidatorPubkey(pubKey), types.ValidatorSignature(signature), depositDataRoot, opts)
	if err != nil {
		return nil, err
	}

	// Save wallet
	if err := w.Save(); err != nil {
		return nil, err
	}

	// Print transaction if requested
	if !submit {
		b, err := tx.MarshalBinary()
		if err != nil {
			return nil, err
		}
		fmt.Printf("%x\n", b)
	}

	response.TxHash = tx.Hash()
	response.ValidatorPubkey = pubKey

	// Return response
	return &response, nil

}
