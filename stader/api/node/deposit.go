package node

import (
	"context"
	"fmt"
	"github.com/rocket-pool/rocketpool-go/types"
	_ "github.com/stader-labs/stader-node/stader-lib/network"
	"github.com/stader-labs/stader-node/stader-lib/node"
	"github.com/stader-labs/stader-node/stader-lib/tokens"
	stadertypes "github.com/stader-labs/stader-node/stader-lib/types"
	"github.com/urfave/cli"
	_ "golang.org/x/sync/errgroup"
	"math/big"

	"github.com/stader-labs/stader-node/shared/services"
	"github.com/stader-labs/stader-node/shared/types/api"
	"github.com/stader-labs/stader-node/shared/utils/eth1"
	"github.com/stader-labs/stader-node/shared/utils/validator"
)

// TODO: refactor canNodeDeposit and nodeDeposit bchain

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
	vfc, err := services.GetVaultFactory(c)
	if err != nil {
		return nil, err
	}
	sdc, err := services.GetSdCollateralContract(c)
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
		canNodeDepositResponse.InsufficientBalance = true
	}

	operatorId, err := node.GetOperatorId(prn, nodeAccount.Address, nil)
	if err != nil {
		return nil, err
	}
	operatorRegistryInfo, err := node.GetOperatorInfo(prn, operatorId, nil)
	if err != nil {
		return nil, err
	}
	if operatorRegistryInfo.OperatorName == "" {
		canNodeDepositResponse.NotRegistered = true
		return &canNodeDepositResponse, nil
	}

	isPermissionlessNodeRegistryPaused, err := node.IsPermissionlessNodeRegistryPaused(prn, nil)
	if err != nil {
		return nil, err
	}
	if isPermissionlessNodeRegistryPaused {
		canNodeDepositResponse.DepositPaused = true
		return &canNodeDepositResponse, nil
	}

	hasEnoughSdCollateral, err := sd_collateral.HasEnoughSdCollateral(sdc, nodeAccount.Address, 1, uint32(numValidators.Int64()), nil)
	if err != nil {
		return nil, err
	}
	if !hasEnoughSdCollateral {
		canNodeDepositResponse.NotEnoughSdCollateral = true
		return &canNodeDepositResponse, nil
	}

	pubKeys := make([][]byte, numValidators.Int64())
	preDepositSignatures := make([][]byte, numValidators.Int64())
	depositSignatures := make([][]byte, numValidators.Int64())

	operatorKeyCount, err := node.GetTotalValidatorKeys(prn, operatorId, nil)
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

	newValidatorKey := operatorKeyCount

	walletIndex := w.GetNextAccount()

	for i := int64(0); i < numValidators.Int64(); i++ {
		// Create and save a new validator key
		validatorKey, err := w.GetValidatorKeyAt(walletIndex)
		if err != nil {
			return nil, err
		}
		walletIndex++

		rewardWithdrawVault, err := node.ComputeWithdrawVaultAddress(vfc, 1, operatorId, newValidatorKey, nil)
		if err != nil {
			return nil, err
		}

		withdrawCredentials, err := node.GetValidatorWithdrawalCredential(vfc, rewardWithdrawVault, nil)
		if err != nil {
			return nil, err
		}

		// Get validator deposit data for 1 eth and 31eth
		preDepositData, _, err := validator.GetDepositData(validatorKey, withdrawCredentials, eth2Config, 1000000000)
		if err != nil {
			return nil, err
		}
		pubKey := stadertypes.BytesToValidatorPubkey(preDepositData.PublicKey)
		preDepositSignature := stadertypes.BytesToValidatorSignature(preDepositData.Signature)

		depositData, _, err := validator.GetDepositData(validatorKey, withdrawCredentials, eth2Config, 31000000000)
		if err != nil {
			return nil, err
		}
		depositSignature := stadertypes.BytesToValidatorSignature(depositData.Signature)

		pubKeys[i] = pubKey[:]
		preDepositSignatures[i] = preDepositSignature[:]
		depositSignatures[i] = depositSignature[:]

		newValidatorKey = operatorKeyCount.Add(operatorKeyCount, big.NewInt(1))
	}

	// Override the provided pending TX if requested
	err = eth1.CheckForNonceOverride(c, opts)
	if err != nil {
		return nil, fmt.Errorf("error checking for nonce override: %w", err)
	}

	// Do not send transaction unless requested
	opts.NoSend = !submit

	gasInfo, err := node.EstimateAddValidatorKeys(prn, pubKeys, preDepositSignatures, depositSignatures, opts)
	if err != nil {
		return nil, err
	}

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
	srcf, err := services.GetVaultFactory(c)
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
	operatorId, err := node.GetOperatorId(prn, nodeAccount.Address, nil)
	if err != nil {
		return nil, err
	}

	operatorRegistryInfo, err := node.GetOperatorInfo(prn, operatorId, nil)
	if err != nil {
		return nil, err
	}

	// Get transactor
	opts, err := w.GetNodeAccountTransactor()
	if err != nil {
		return nil, err
	}

	pubKeys := make([][]byte, numValidators.Int64())
	preDepositSignatures := make([][]byte, numValidators.Int64())
	depositSignatures := make([][]byte, numValidators.Int64())

	amountToSend := amountWei.Mul(amountWei, numValidators)
	opts.Value = amountToSend

	nodeAccount, err = w.GetNodeAccount()
	validatorKeyCount, err := node.GetTotalValidatorKeys(prn, operatorId, nil)
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

		rewardWithdrawVault, err := node.ComputeWithdrawVaultAddress(srcf, 1, operatorId, newValidatorKey, nil)
		if err != nil {
			return nil, err
		}

		withdrawCredentials, err := node.GetValidatorWithdrawalCredential(srcf, rewardWithdrawVault, nil)
		if err != nil {
			return nil, err
		}

		// Get validator deposit data for 1 eth and 31eth
		preDepositData, _, err := validator.GetDepositData(validatorKey, withdrawCredentials, eth2Config, 1000000000)
		if err != nil {
			return nil, err
		}
		pubKey := stadertypes.BytesToValidatorPubkey(preDepositData.PublicKey)
		preDepositSignature := stadertypes.BytesToValidatorSignature(preDepositData.Signature)

		depositData, _, err := validator.GetDepositData(validatorKey, withdrawCredentials, eth2Config, 31000000000)
		if err != nil {
			return nil, err
		}
		depositSignature := stadertypes.BytesToValidatorSignature(depositData.Signature)

		pubKeys[i] = pubKey[:]
		preDepositSignatures[i] = preDepositSignature[:]
		depositSignatures[i] = depositSignature[:]

		// Make sure a validator with this pubkey doesn't already exist
		status, err := bc.GetValidatorStatus(types.ValidatorPubkey(pubKey), nil)
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

	tx, err := node.AddValidatorKeys(prn, pubKeys, preDepositSignatures, depositSignatures, opts)
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

	// Return response
	return &response, nil

}
