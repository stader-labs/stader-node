package node

import (
	"fmt"
	"github.com/stader-labs/stader-node/shared/services/stader"
	cliutils "github.com/stader-labs/stader-node/shared/utils/cli"
	"github.com/stader-labs/stader-node/shared/utils/log"
	"github.com/urfave/cli"
)

func getContractsInfo(c *cli.Context) error {
	staderClient, err := stader.NewClientFromCtx(c)
	if err != nil {
		return err
	}
	defer staderClient.Close()

	// Check and assign the EC status
	err = cliutils.CheckClientStatus(staderClient)
	if err != nil {
		return err
	}

	// Get node fee
	response, err := staderClient.GetContractsInfo()
	if err != nil {
		return err
	}

	fmt.Printf("%s=== Beacon Network Contract Details ===%s\n", log.ColorGreen, log.ColorReset)
	fmt.Printf("Beacon Network: %d\n\n", response.BeaconNetwork)
	fmt.Printf("Beacon Deposit Contract: %s\n\n", response.BeaconDepositContract)

	fmt.Printf("%s=== Stader Network Contract Details ===%s\n", log.ColorGreen, log.ColorReset)
	fmt.Printf("Network: %d\n\n", response.Network)
	fmt.Printf("Stader Config: %s\n\n", response.StaderConfig)
	fmt.Printf("Permissionless Node Registry: %s\n\n", response.PermissionlessNodeRegistry)
	fmt.Printf("Vault Factory: %s\n\n", response.VaultFactory)
	fmt.Printf("Sd Collateral Lock: %s\n\n", response.SdCollateralContract)
	fmt.Printf("EthX Token: %s\n\n", response.EthxToken)
	fmt.Printf("Sd Token: %s\n\n", response.SdToken)
	fmt.Printf("Socializing Pool: %s\n\n", response.SocializingPoolContract)
	fmt.Printf("Permissionless Pool: %s\n\n", response.PermisionlessPool)
	fmt.Printf("Stader Oracle: %s\n\n", response.StaderOracle)
	fmt.Printf("Pre-sign encryption key is %s\n\n", response.EncryptionKey)

	return nil
}
