package node

import (
	"fmt"
	"github.com/stader-labs/stader-node/shared/services/stader"
	cliutils "github.com/stader-labs/stader-node/shared/utils/cli"
	"github.com/stader-labs/stader-node/shared/utils/log"
	"github.com/urfave/cli"
	"math/big"
)

func debugExitMsg(c *cli.Context, validatorIndex uint64, epochDelta uint64) error {
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
	response, err := staderClient.DebugExit(big.NewInt(int64(validatorIndex)), big.NewInt(int64(epochDelta)))
	if err != nil {
		return err
	}

	fmt.Printf("%s=== Exit details ===%s\n", log.ColorGreen, log.ColorReset)
	fmt.Printf("Validator pub key is %s\n", response.ValidatorPubKey)
	fmt.Printf("Exit epoch is %d\n", response.ExitEpoch)
	fmt.Printf("Current epoch is %d\n", response.CurrentEpoch)
	fmt.Printf("Validator index is %d\n", response.ValidatorIndex)
	fmt.Printf("Signed msg is %s\n", response.SignedMsg.Hex())
	fmt.Printf("Signature domain is %s\n", response.SignatureDomain.Hex())
	fmt.Printf("SrHash is %s\n", response.SrHash.Hex())

	return nil
}
