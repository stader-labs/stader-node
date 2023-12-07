package validator

import (
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/stader-labs/stader-node/stader-lib/node"
	sd_collateral "github.com/stader-labs/stader-node/stader-lib/sd-collateral"
	"github.com/stader-labs/stader-node/stader-lib/sdutility"
	"github.com/stader-labs/stader-node/stader-lib/stader"
	"github.com/stader-labs/stader-node/stader-lib/tokens"
	stadertypes "github.com/stader-labs/stader-node/stader-lib/types"
	"github.com/urfave/cli"
	_ "golang.org/x/sync/errgroup"

	"github.com/stader-labs/stader-node/shared/services"
	"github.com/stader-labs/stader-node/shared/types/api"
	"github.com/stader-labs/stader-node/shared/utils/eth1"
	"github.com/stader-labs/stader-node/shared/utils/validator"
)

const (
	MinToUtilizeRatioNum   = 5
	MinToUtilizeRatioDenom = 2
)

func GetSDStatus(
	sdc *stader.SdCollateralContractManager,
	sdu *stader.SDUtilityPoolContractManager,
	sdt *stader.Erc20TokenContractManager,
	operatorAddress common.Address,
	totalValidatorsPostAddition *big.Int,
) (*api.SdStatusResponse, error) {
	sdUtilityLatestBalance, err := sdutility.GetUtilizerLatestBalance(sdu, operatorAddress, nil)
	if err != nil {
		return nil, err
	}

	poolAvailableSDBalance, err := sdutility.GetPoolAvailableSDBalance(sdu, nil)
	if err != nil {
		return nil, err
	}

	minimumSDToBond, err := sd_collateral.MinimumSDToBond(sdc, 1, totalValidatorsPostAddition, nil)
	if err != nil {
		return nil, err
	}

	sdCollateralCurrentAmount, err := sd_collateral.GetOperatorSdBalance(sdc, operatorAddress, nil)
	if err != nil {
		return nil, err
	}

	sdUtilizedBalance, err := sd_collateral.GetOperatorUtilizedSDBalance(sdc, operatorAddress, nil)
	if err != nil {
		return nil, err
	}

	rewardEligibleSD, err := sd_collateral.RewardEligibleSD(sdc, 1, totalValidatorsPostAddition, nil)
	if err != nil {
		return nil, err
	}

	// Check Sd balance
	sdBalance, err := sdt.Erc20Token.BalanceOf(nil, operatorAddress)
	if err != nil {
		return nil, err
	}

	sdMaxUtilizableAmount := new(big.Int).Mul(minimumSDToBond, big.NewInt(MinToUtilizeRatioNum))
	sdMaxUtilizableAmount = new(big.Int).Div(sdMaxUtilizableAmount, big.NewInt(MinToUtilizeRatioDenom))

	hasEnoughSdCollateral, err := sd_collateral.HasEnoughSdCollateral(sdc, operatorAddress, 1, totalValidatorsPostAddition, nil)
	if err != nil {
		return nil, err
	}

	return &api.SdStatusResponse{
		NotEnoughSdCollateral:     !hasEnoughSdCollateral,
		SdUtilizerLatestBalance:   sdUtilityLatestBalance,
		SdBalance:                 sdBalance,
		SdCollateralCurrentAmount: sdCollateralCurrentAmount,
		SdCollateralRequireAmount: minimumSDToBond,
		SdMaxUtilizableAmount:     sdMaxUtilizableAmount,
		SdUtilizedBalance:         sdUtilizedBalance,
		PoolAvailableSDBalance:    poolAvailableSDBalance,
		SdRewardEligible:          rewardEligibleSD,
	}, nil
}

func canNodeDeposit(c *cli.Context, baseAmountWei, utilityAmountWei, numValidators *big.Int, reloadKeys bool) (*api.CanNodeDepositResponse, error) {
	if err := services.RequireNodeWallet(c); err != nil {
		return nil, err
	}
	if err := services.RequireNodeRegistered(c); err != nil {
		return nil, err
	}
	if err := services.RequireNodeActive(c); err != nil {
		return nil, err
	}
	w, err := services.GetWallet(c)
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

	amountToSend := baseAmountWei.Mul(baseAmountWei, numValidators)

	opts, err := w.GetNodeAccountTransactor()
	if err != nil {
		return nil, err
	}

	opts.Value = amountToSend

	canNodeDepositResponse := api.CanNodeDepositResponse{}

	userBalance, err := tokens.GetEthBalance(prn.Client, nodeAccount.Address, nil)
	if err != nil {
		return nil, err
	}
	if userBalance.Cmp(amountToSend) < 0 {
		canNodeDepositResponse.InsufficientBalance = true
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

	operatorId, err := node.GetOperatorId(prn, nodeAccount.Address, nil)
	if err != nil {
		return nil, err
	}

	totalValidatorKeys, err := node.GetTotalValidatorKeys(prn, operatorId, nil)
	if err != nil {
		return nil, err
	}
	totalValidatorNonTerminalKeys, err := node.GetTotalNonTerminalValidatorKeys(prn, nodeAccount.Address, totalValidatorKeys, nil)
	if err != nil {
		return nil, err
	}
	maxKeysPerOperator, err := node.GetMaxValidatorKeysPerOperator(prn, nil)
	if err != nil {
		return nil, err
	}

	totalValidatorsPostAddition := totalValidatorNonTerminalKeys + numValidators.Uint64()

	inputKeyLimitCount, err := node.GetInputKeyLimitCount(prn, nil)
	if err != nil {
		return nil, err
	}
	if numValidators.Cmp(big.NewInt(int64(inputKeyLimitCount))) > 0 {
		canNodeDepositResponse.InputKeyLimitReached = true
		canNodeDepositResponse.InputKeyLimit = inputKeyLimitCount
		return &canNodeDepositResponse, nil
	}

	if totalValidatorsPostAddition > maxKeysPerOperator {
		canNodeDepositResponse.MaxValidatorLimitReached = true
		return &canNodeDepositResponse, nil
	}

	pubKeys := make([][]byte, numValidators.Int64())
	preDepositSignatures := make([][]byte, numValidators.Int64())
	depositSignatures := make([][]byte, numValidators.Int64())

	operatorKeyCount, err := node.GetTotalValidatorKeys(prn, operatorId, nil)
	if err != nil {
		return nil, err
	}

	newValidatorKey := operatorKeyCount

	walletIndex, err := w.GetValidatorKeyCount()
	if err != nil {
		return nil, err
	}

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

		// Get validator deposit data for 1 eth
		preDepositData, _, err := validator.GetDepositData(validatorKey, withdrawCredentials, eth2Config, 1000000000)
		if err != nil {
			return nil, err
		}
		preDepositSignature := stadertypes.BytesToValidatorSignature(preDepositData.Signature)

		depositData, _, err := validator.GetDepositData(validatorKey, withdrawCredentials, eth2Config, 31000000000)
		if err != nil {
			return nil, err
		}
		depositSignature := stadertypes.BytesToValidatorSignature(depositData.Signature)

		pubKey := stadertypes.BytesToValidatorPubkey(preDepositData.PublicKey)

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

	gasInfo, err := node.EstimateAddValidatorKeys(
		prn,
		utilityAmountWei,
		pubKeys,
		preDepositSignatures,
		depositSignatures,
		opts,
	)
	if err != nil {
		return nil, err
	}

	canNodeDepositResponse.CanDeposit = true
	canNodeDepositResponse.GasInfo = gasInfo

	return &canNodeDepositResponse, nil
}

func nodeDeposit(c *cli.Context, baseAmountWei, utilityAmountWei, numValidators *big.Int, reloadKeys bool) (*api.NodeDepositResponse, error) {
	cfg, err := services.GetConfig(c)
	if err != nil {
		return nil, err
	}
	w, err := services.GetWallet(c)
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

	amountToSend := baseAmountWei.Mul(baseAmountWei, numValidators)
	opts.Value = amountToSend

	validatorKeyCount, err := node.GetTotalValidatorKeys(prn, operatorId, nil)
	if err != nil {
		return nil, err
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

		// Get validator deposit data for 1 eth
		preDepositData, _, err := validator.GetDepositData(validatorKey, withdrawCredentials, eth2Config, 1000000000)
		if err != nil {
			return nil, err
		}
		preDepositSignature := stadertypes.BytesToValidatorSignature(preDepositData.Signature)

		pubKey := stadertypes.BytesToValidatorPubkey(preDepositData.PublicKey)

		depositData, _, err := validator.GetDepositData(validatorKey, withdrawCredentials, eth2Config, 31000000000)
		if err != nil {
			return nil, err
		}
		depositSignature := stadertypes.BytesToValidatorSignature(depositData.Signature)

		pubKeys[i] = pubKey[:]
		preDepositSignatures[i] = preDepositSignature[:]
		depositSignatures[i] = depositSignature[:]

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

		newValidatorKey = validatorKeyCount.Add(validatorKeyCount, big.NewInt(1))
	}

	if reloadKeys {
		d, err := services.GetDocker(c)
		if err != nil {
			return nil, err
		}

		// Restart the validator container when a new key have been saved
		err = validator.RestartValidator(cfg, bc, nil, d)
		if err != nil {
			return nil, err
		}
	}

	// Override the provided pending TX if requested
	err = eth1.CheckForNonceOverride(c, opts)
	if err != nil {
		return nil, fmt.Errorf("error checking for nonce override: %w", err)
	}

	tx, err := node.AddValidatorKeysWithAmount(prn,
		pubKeys,
		preDepositSignatures,
		depositSignatures,
		utilityAmountWei,
		opts)
	if err != nil {
		return nil, err
	}

	// To save the validator index update
	if err := w.Save(); err != nil {
		return nil, err
	}

	response.TxHash = tx.Hash()

	// Return response
	return &response, nil

}
