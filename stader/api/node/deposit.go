package node

import (
	"context"
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/prysmaticlabs/prysm/v3/beacon-chain/core/signing"
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

	prdeposit "github.com/prysmaticlabs/prysm/v3/contracts/deposit"
	ethpb "github.com/prysmaticlabs/prysm/v3/proto/prysm/v1alpha1"
	"github.com/stader-labs/stader-node/shared/services"
	"github.com/stader-labs/stader-node/shared/services/beacon"
	"github.com/stader-labs/stader-node/shared/types/api"
	"github.com/stader-labs/stader-node/shared/utils/eth1"
	"github.com/stader-labs/stader-node/shared/utils/validator"
	eth2types "github.com/wealdtech/go-eth2-types/v2"
)

type minipoolCreated struct {
	Minipool common.Address
	Node     common.Address
	Time     *big.Int
}

func nodeDeposit(c *cli.Context, amountWei *big.Int, salt *big.Int, submit bool) (*api.NodeDepositResponse, error) {

	// Get services
	//if err := services.RequireNodeRegistered(c); err != nil {
	//	return nil, err
	//}

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

	// Make sure ETH2 is on the correct chain
	//depositContractInfo, err := getDepositContractInfo(c)
	//if err != nil {
	//	return nil, err
	//}
	//if depositContractInfo.RPNetwork != depositContractInfo.BeaconNetwork ||
	//	depositContractInfo.RPDepositContract != depositContractInfo.BeaconDepositContract {
	//	return nil, fmt.Errorf("Beacon network mismatch! Expected %s on chain %d, but beacon is using %s on chain %d.",
	//		depositContractInfo.RPDepositContract.Hex(),
	//		depositContractInfo.RPNetwork,
	//		depositContractInfo.BeaconDepositContract.Hex(),
	//		depositContractInfo.BeaconNetwork)
	//}

	// Get the scrub period
	//scrubPeriodUnix, err := trustednode.GetScrubPeriod(rp, nil)
	//if err != nil {
	//	return nil, err
	//}
	//scrubPeriod := time.Duration(scrubPeriodUnix) * time.Second
	//response.ScrubPeriod = scrubPeriod

	// Get transactor
	opts, err := w.GetNodeAccountTransactor()
	if err != nil {
		return nil, err
	}
	opts.Value = amountWei

	// Get the deposit type
	//depositType, err := node.GetDepositType(rp, amountWei, nil)
	//if err != nil {
	//	return nil, err
	//}

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

	// Do a final sanity check
	//err = validateDepositInfo(eth2Config, uint64(validator.DepositAmount), pubKey, withdrawalCredentials, signature)
	//if err != nil {
	//	return nil, fmt.Errorf("Your deposit failed the validation safety check: %w\n"+
	//		"For your safety, this deposit will not be submitted and your ETH will not be staked.\n"+
	//		"PLEASE REPORT THIS TO THE ROCKET POOL DEVELOPERS and include the following information:\n"+
	//		"\tDomain Type: 0x%s\n"+
	//		"\tGenesis Fork Version: 0x%s\n"+
	//		"\tGenesis Validator Root: 0x%s\n"+
	//		"\tDeposit Amount: %d gwei\n"+
	//		"\tValidator Pubkey: %s\n"+
	//		"\tWithdrawal Credentials: %s\n"+
	//		"\tSignature: %s\n",
	//		err,
	//		hex.EncodeToString(eth2types.DomainDeposit[:]),
	//		hex.EncodeToString(eth2Config.GenesisForkVersion),
	//		hex.EncodeToString(eth2types.ZeroGenesisValidatorsRoot),
	//		uint64(validator.DepositAmount),
	//		pubKey.Hex(),
	//		withdrawalCredentials.Hex(),
	//		signature.Hex(),
	//	)
	//}

	// Override the provided pending TX if requested
	err = eth1.CheckForNonceOverride(c, opts)
	if err != nil {
		return nil, fmt.Errorf("Error checking for nonce override: %w", err)
	}

	// Do not send transaction unless requested
	opts.NoSend = !submit

	// Deposit
	//tx, err := node.Deposit(ethxcm, minNodeFee, types.ValidatorPubkey(pubKey), types.ValidatorSignature(signature), depositDataRoot, salt, minipoolAddress, opts)
	//if err != nil {
	//	return nil, err
	//}

	//tx, err := node.OnboardNodeOperator(sor, true, 0, operatorName, operatorRewardAddress, opts)
	//if err != nil {
	//	return nil, err
	//}

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
	//response.MinipoolAddress = minipoolAddress
	response.ValidatorPubkey = pubKey

	// Return response
	return &response, nil

}

func validateDepositInfo(eth2Config beacon.Eth2Config, depositAmount uint64, pubkey rptypes.ValidatorPubkey, withdrawalCredentials common.Hash, signature rptypes.ValidatorSignature) error {

	// Get the deposit domain based on the eth2 config
	depositDomain, err := signing.ComputeDomain(eth2types.DomainDeposit, eth2Config.GenesisForkVersion, eth2types.ZeroGenesisValidatorsRoot)
	if err != nil {
		return err
	}

	// Create the deposit struct
	depositData := new(ethpb.Deposit_Data)
	depositData.Amount = depositAmount
	depositData.PublicKey = pubkey.Bytes()
	depositData.WithdrawalCredentials = withdrawalCredentials.Bytes()
	depositData.Signature = signature.Bytes()

	// Validate the signature
	err = prdeposit.VerifyDepositSignature(depositData, depositDomain)
	return err

}
