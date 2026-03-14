package node

import (
	"fmt"

	"github.com/stader-labs/stader-node/shared/services/stader"
	cliutils "github.com/stader-labs/stader-node/shared/utils/cli"
	"github.com/urfave/cli"
)

func downloadSPMerkleProofs(c *cli.Context) error {

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

	canDownloadSpMerkleProofs, err := staderClient.CanDownloadSpMerkleProofs()
	if err != nil {
		return err
	}
	if canDownloadSpMerkleProofs.NoMissingCycles {
		fmt.Println("There are no missing cycles to download! All proofs are up to date!")
		return nil
	}

	fmt.Printf("Following cycles are missing: %v\n", canDownloadSpMerkleProofs.MissingCycles)

	// Prompt for confirmation
	if !(c.Bool("yes") || cliutils.Confirm(fmt.Sprintf(
		"Are you sure you want to download the missing merkle proofs?"))) {
		fmt.Println("Cancelled.")
		return nil
	}

	fmt.Println("Downloading missing merkle proofs.....")

	res, err := staderClient.DownloadSpMerkleProofs()
	if err != nil {
		return err
	}

	fmt.Printf("Successfully downloaded the merkle proofs for cycles: %v\n", res.DownloadedCycles)

	return nil
}
