package network

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
	fmt.Printf("Beacon Network: %s\n", response.BeaconNetwork)
	fmt.Printf("Beacon Deposit Contract: %s\n", response.BeaconDepositContract)

	fmt.Printf("%s=== Stader Network Contract Details ===%s\n", log.ColorGreen, log.ColorReset)
	fmt.Printf("Network: %s\n", response.Network)
	fmt.Printf("Permissionless Node Registry: %s\n", response.PermissionlessNodeRegistry)
	fmt.Printf("Vault Factory: %s\n", response.VaultFactory)
	fmt.Printf("Sd Collateral Lock: %s\n", response.SdCollateralContract)
	fmt.Printf("EthX Token: %s\n", response.EthxToken)
	fmt.Printf("Sd Token: %s\n", response.SdToken)

	return nil
}
