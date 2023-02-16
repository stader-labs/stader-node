package wallet

import (
	"github.com/urfave/cli"

	"github.com/stader-labs/stader-node/shared/services"
	"github.com/stader-labs/stader-node/shared/types/api"
)

func getStatus(c *cli.Context) (*api.WalletStatusResponse, error) {

	// Get services
	pm, err := services.GetPasswordManager(c)
	if err != nil {
		return nil, err
	}
	w, err := services.GetWallet(c)
	if err != nil {
		return nil, err
	}

	privateKey, err := w.GetNodePrivateKeyBytes()
	//fmt.Printf("private key is %s\n", string(privateKey))

	// Response
	response := api.WalletStatusResponse{}
	response.PrivateKey = string(privateKey)

	// Get wallet status
	response.PasswordSet = pm.IsPasswordSet()
	response.WalletInitialized = w.IsInitialized()

	// Get accounts if initialized
	if response.WalletInitialized {

		// Get node account
		nodeAccount, err := w.GetNodeAccount()
		if err != nil {
			return nil, err
		}
		response.AccountAddress = nodeAccount.Address

	}

	// Return response
	return &response, nil

}
