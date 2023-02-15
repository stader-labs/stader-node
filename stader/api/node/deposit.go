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
	"github.com/urfave/cli"
	_ "golang.org/x/sync/errgroup"
	"math/big"

	"github.com/stader-labs/stader-node/shared/services"
	"github.com/stader-labs/stader-node/shared/types/api"
	"github.com/stader-labs/stader-node/shared/utils/eth1"
	"github.com/stader-labs/stader-node/shared/utils/validator"
)

func nodeDeposit(c *cli.Context, amountWei *big.Int, salt *big.Int, numValidators uint64, submit bool) (*api.NodeDepositResponse, error) {

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

	validatorPubKeys := make([][]byte, numValidators)
	validatorSignatures := make([][]byte, numValidators)
	depositDataRoots := make([][32]byte, numValidators)

	// Get transactor
	opts, err := w.GetNodeAccountTransactor()
	if err != nil {
		return nil, err
	}
	opts.Value = amountWei.Mul(amountWei, big.NewInt(int64(numValidators)))

	nodeAccount, err = w.GetNodeAccount()
	validatorKeyCount, err := node.GetTotalValidatorKeys(prn, nodeAccount.Address, nil)
	if err != nil {
		return nil, err
	}

	// bchain - TODO - convert numValidators to int64
	for i := uint64(0); i < numValidators; i++ {
		// Adjust the salt
		if salt.Cmp(big.NewInt(0)) == 0 {
			nonce, err := ec.NonceAt(context.Background(), nodeAccount.Address, nil)
			if err != nil {
				return nil, err
			}
			salt.SetUint64(nonce)
		}

		nextValidatorKeyCount := validatorKeyCount.Abs(big.NewInt(int64(i)))

		// Create and save a new validator key
		validatorKey, err := w.CreateValidatorKey()
		if err != nil {
			return nil, err
		}

		rewardWithdrawVault, err := node.ComputeWithdrawVaultAddress(srcf, 1, operatorRegistryInfo.OperatorId, nextValidatorKeyCount, nil)
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
		for i := 0; i < 32; i++ {
			depositDataRootFixedSize[i] = depositDataRootBytes[i]
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

	// Return response
	return &response, nil

}
