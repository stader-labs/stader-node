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
	"github.com/stader-labs/stader-minipool-go/tokens"
	"github.com/urfave/cli"
	_ "golang.org/x/sync/errgroup"
	"math/big"

	"github.com/stader-labs/stader-node/shared/services"
	"github.com/stader-labs/stader-node/shared/types/api"
	"github.com/stader-labs/stader-node/shared/utils/eth1"
	"github.com/stader-labs/stader-node/shared/utils/validator"
)

func canNodeDeposit(c *cli.Context, amountWei *big.Int, salt *big.Int, numValidators *big.Int, submit bool) (*api.CanNodeDepositResponse, error) {
	canNodeDepositResponse := api.CanNodeDepositResponse{}

	w, err := services.GetWallet(c)
	if err != nil {
		return nil, err
	}
	ec, err := services.GetEthClient(c)
	if err != nil {
		return nil, err
	}
	prn, err := services.GetPermissionlessNodeRegistry(c)
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

	amountToSend := amountWei.Mul(amountWei, numValidators)

	opts, err := w.GetNodeAccountTransactor()
	if err != nil {
		return nil, err
	}

	opts.Value = amountToSend

	userBalance, err := tokens.GetEthBalance(prn.Client, nodeAccount.Address, nil)
	if err != nil {
		return nil, err
	}
	if userBalance.Cmp(amountToSend) < 0 {
		canNodeDepositResponse.CanDeposit = false
		canNodeDepositResponse.InsufficientBalance = true
	}

	operatorRegistryInfo, err := node.GetOperatorRegistry(prn, nodeAccount.Address, nil)
	if err != nil {
		return nil, err
	}
	if operatorRegistryInfo.OperatorName == "" {
		return nil, fmt.Errorf("node is not registered with stader")
	}

	validatorPubKeys := make([][]byte, numValidators.Int64())
	validatorSignatures := make([][]byte, numValidators.Int64())
	depositDataRoots := make([][32]byte, numValidators.Int64())

	nodeAccount, err = w.GetNodeAccount()
	validatorKeyCount, err := node.GetTotalValidatorKeys(prn, nodeAccount.Address, nil)
	if err != nil {
		return nil, err
	}
	//fmt.Printf("Got validator key account %d\n", validatorKeyCount)

	// Adjust the salt
	if salt.Cmp(big.NewInt(0)) == 0 {
		nonce, err := ec.NonceAt(context.Background(), nodeAccount.Address, nil)
		if err != nil {
			return nil, err
		}
		salt.SetUint64(nonce)
	}

	newValidatorKey := validatorKeyCount

	walletIndex := w.GetNextAccount()
	//fmt.Printf("walletIndex is %d\n", walletIndex)

	for i := int64(0); i < numValidators.Int64(); i++ {
		// Create and save a new validator key
		//fmt.Printf("walletIndex is %d\n", walletIndex)
		validatorKey, err := w.GetValidatorKeyAt(walletIndex)
		if err != nil {
			return nil, err
		}
		walletIndex++
		//validatorKey, err := w.CreateValidatorKey()
		//if err != nil {
		//	return nil, err
		//}

		//fmt.Printf("walletIndex post increment is %d\n", walletIndex)
		//fmt.Printf("validator public key is %s\n", validatorKey.PublicKey())

		rewardWithdrawVault, err := node.ComputeWithdrawVaultAddress(srcf, 1, operatorRegistryInfo.OperatorId, newValidatorKey, nil)
		if err != nil {
			return nil, err
		}
		//fmt.Printf("reward withdraw vault is %s\n", rewardWithdrawVault.String())

		withdrawCredentials, err := node.GetValidatorWithdrawalCredential(srcf, rewardWithdrawVault, nil)
		if err != nil {
			return nil, err
		}
		//fmt.Printf("withdraw creds is %s\n", withdrawCredentials.String())

		// Get validator deposit data and associated parameters
		depositData, depositDataRoot, err := validator.GetDepositData(validatorKey, withdrawCredentials, eth2Config)
		if err != nil {
			return nil, err
		}
		pubKey := rptypes.BytesToValidatorPubkey(depositData.PublicKey)
		signature := rptypes.BytesToValidatorSignature(depositData.Signature)

		//fmt.Printf("depositDataRoot is %v\n", depositData)

		// convert deposit data root to 32 bytes
		depositDataRootFixedSize := [32]byte{}
		depositDataRootBytes := depositDataRoot.Bytes()
		for j := 0; j < 32; j++ {
			depositDataRootFixedSize[j] = depositDataRootBytes[j]
		}

		validatorPubKeys[i] = pubKey[:]
		validatorSignatures[i] = signature[:]
		depositDataRoots[i] = depositDataRootFixedSize

		newValidatorKey = validatorKeyCount.Add(validatorKeyCount, big.NewInt(1))

		if err := w.Save(); err != nil {
			return nil, err
		}
	}

	//fmt.Println("Finished generating validator keys")

	// Override the provided pending TX if requested
	err = eth1.CheckForNonceOverride(c, opts)
	if err != nil {
		return nil, fmt.Errorf("error checking for nonce override: %w", err)
	}

	// Do not send transaction unless requested
	opts.NoSend = !submit

	//fmt.Printf("Estimating gas")
	gasInfo, err := node.EstimateAddValidatorKeys(prn, validatorPubKeys, validatorSignatures, depositDataRoots, opts)
	if err != nil {
		return nil, err
	}
	//fmt.Printf("canNodeDepositResponse is %v\n", canNodeDepositResponse)
	//fmt.Printf("Estimated gas!")

	canNodeDepositResponse.CanDeposit = true
	canNodeDepositResponse.GasInfo = gasInfo

	return &canNodeDepositResponse, nil
}

func nodeDeposit(c *cli.Context, amountWei *big.Int, salt *big.Int, numValidators *big.Int, submit bool) (*api.NodeDepositResponse, error) {

	w, err := services.GetWallet(c)
	if err != nil {
		return nil, err
	}
	ec, err := services.GetEthClient(c)
	if err != nil {
		return nil, err
	}
	prn, err := services.GetPermissionlessNodeRegistry(c)
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

	// get the vault address and vault credential
	operatorRegistryInfo, err := node.GetOperatorRegistry(prn, nodeAccount.Address, nil)
	if err != nil {
		return nil, err
	}
	if operatorRegistryInfo.OperatorName == "" {
		return nil, fmt.Errorf("node is not registered with stader")
	}

	validatorPubKeys := make([][]byte, numValidators.Int64())
	validatorSignatures := make([][]byte, numValidators.Int64())
	depositDataRoots := make([][32]byte, numValidators.Int64())

	// Get transactor
	opts, err := w.GetNodeAccountTransactor()
	if err != nil {
		return nil, err
	}

	amountToSend := amountWei.Mul(amountWei, numValidators)
	opts.Value = amountToSend

	nodeAccount, err = w.GetNodeAccount()
	validatorKeyCount, err := node.GetTotalValidatorKeys(prn, nodeAccount.Address, nil)
	if err != nil {
		return nil, err
	}

	// Adjust the salt
	if salt.Cmp(big.NewInt(0)) == 0 {
		nonce, err := ec.NonceAt(context.Background(), nodeAccount.Address, nil)
		if err != nil {
			return nil, err
		}
		salt.SetUint64(nonce)
	}

	newValidatorKey := validatorKeyCount

	for i := int64(0); i < numValidators.Int64(); i++ {
		// Create and save a new validator key
		validatorKey, err := w.CreateValidatorKey()
		if err != nil {
			return nil, err
		}

		rewardWithdrawVault, err := node.ComputeWithdrawVaultAddress(srcf, 1, operatorRegistryInfo.OperatorId, newValidatorKey, nil)
		if err != nil {
			return nil, err
		}

		withdrawCredentials, err := node.GetValidatorWithdrawalCredential(srcf, rewardWithdrawVault, nil)
		if err != nil {
			return nil, err
		}

		// Get validator deposit data and associated parameters
		depositData, depositDataRoot, err := validator.GetDepositData(validatorKey, withdrawCredentials, eth2Config)
		if err != nil {
			return nil, err
		}
		pubKey := rptypes.BytesToValidatorPubkey(depositData.PublicKey)
		signature := rptypes.BytesToValidatorSignature(depositData.Signature)

		// convert deposit data root to 32 bytes
		depositDataRootFixedSize := [32]byte{}
		depositDataRootBytes := depositDataRoot.Bytes()
		for j := 0; j < 32; j++ {
			depositDataRootFixedSize[j] = depositDataRootBytes[j]
		}

		// Make sure a validator with this pubkey doesn't already exist
		status, err := bc.GetValidatorStatus(pubKey, nil)
		if err != nil {
			return nil, fmt.Errorf("Error checking for existing validator status: %w\nYour funds have not been deposited for your own safety.", err)
		}
		if status.Exists {
			return nil, fmt.Errorf("**** ALERT ****\n"+
				"Your validator %s has the following as a validator pubkey:\n\t%s\n"+
				"This key is already in use by validator %d on the Beacon chain!\n"+
				"Stader will not allow you to deposit this validator for your own safety so you do not get slashed.\n"+
				"PLEASE REPORT THIS TO THE STADER DEVELOPERS.\n"+
				"***************\n", operatorRegistryInfo.OperatorName, pubKey.Hex(), status.Index)
		}

		validatorPubKeys[i] = pubKey[:]
		validatorSignatures[i] = signature[:]
		depositDataRoots[i] = depositDataRootFixedSize

		// To save the validator index update
		if err := w.Save(); err != nil {
			return nil, err
		}

		newValidatorKey = validatorKeyCount.Add(validatorKeyCount, big.NewInt(1))
	}

	// Override the provided pending TX if requested
	err = eth1.CheckForNonceOverride(c, opts)
	if err != nil {
		return nil, fmt.Errorf("error checking for nonce override: %w", err)
	}

	// Do not send transaction unless requested
	opts.NoSend = !submit

	tx, err := node.AddValidatorKeys(prn, validatorPubKeys, validatorSignatures, depositDataRoots, opts)
	if err != nil {
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
	response.Status = "success"

	// Return response
	return &response, nil

}
