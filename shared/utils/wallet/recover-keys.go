package wallet

import (
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/stader-labs/stader-node/shared/services/wallet"
	"github.com/stader-labs/stader-node/shared/utils/stdr"
	"github.com/stader-labs/stader-node/stader-lib/node"
	"github.com/stader-labs/stader-node/stader-lib/stader"
	"github.com/stader-labs/stader-node/stader-lib/types"
	"math/big"
)

const (
	pageSize  uint = 20
	pageLimit uint = 2000
)

func RecoverStaderKeys(pnr *stader.PermissionlessNodeRegistryContractManager, address common.Address, w *wallet.Wallet, testOnly bool) ([]types.ValidatorPubkey, error) {
	recoveredKeys := []types.ValidatorPubkey{}
	operatorId, err := node.GetOperatorId(pnr, address, nil)
	if err != nil {
		return nil, err
	}
	if operatorId.Cmp(big.NewInt(0)) == 0 {
		return recoveredKeys, nil
	}
	// Get node's validating pubkeys
	allOperatorValidators, _, err := stdr.GetAllValidatorsRegisteredWithOperator(pnr, operatorId, address, nil)
	if err != nil {
		return nil, err
	}

	// Recover conventionally generated keys
	pageStart := uint(0)
	for {
		if pageStart >= pageLimit {
			return nil, fmt.Errorf("attempt limit exceeded (%d keys)", pageLimit)
		}
		pageEnd := pageStart + pageSize
		if pageEnd > pageLimit {
			pageEnd = pageLimit
		}

		// Get the keys for this bucket
		keys, err := w.GetValidatorKeys(pageStart, pageEnd-pageStart)
		if err != nil {
			return nil, err
		}
		for _, validatorKey := range keys {
			_, exists := allOperatorValidators[validatorKey.PublicKey]
			if exists {
				// Found one!
				delete(allOperatorValidators, validatorKey.PublicKey)
				if !testOnly {
					err := w.SaveValidatorKey(validatorKey)
					if err != nil {
						return nil, fmt.Errorf("error recovering validator keys: %w", err)
					}
					recoveredKeys = append(recoveredKeys, validatorKey.PublicKey)
				}
			}
		}

		if len(allOperatorValidators) == 0 {
			// All keys recovered!
			break
		}

		// Run another iteration with the next bucket
		pageStart = pageEnd
	}

	return recoveredKeys, nil

}
