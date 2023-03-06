package tokens

import (
	"math/big"

	"github.com/stader-labs/stader-minipool-go/stader"
	"github.com/stader-labs/stader-minipool-go/tests/testutils/accounts"
)

// Mint an amount of rETH to an account
func MintRETH(rp *stader.PermissionlessNodeRegistryContractManager, toAccount *accounts.Account, amount *big.Int) error {

	//// Get ETH value of amount
	//ethValue, err := tokens.GetETHValueOfRETH(rp, amount, nil)
	//if err != nil {
	//	return err
	//}
	//
	//// Deposit from account to mint rETH
	//opts := toAccount.GetTransactor()
	//opts.Value = ethValue
	//_, err = deposit.Deposit(rp, opts)
	return nil

}
