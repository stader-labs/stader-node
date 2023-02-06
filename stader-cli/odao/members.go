package odao

import (
	"fmt"

	"github.com/rocket-pool/rocketpool-go/utils/eth"
	"github.com/urfave/cli"

	"github.com/stader-labs/stader-node/shared/services/stader"
	cliutils "github.com/stader-labs/stader-node/shared/utils/cli"
	"github.com/stader-labs/stader-node/shared/utils/math"
)

func getMembers(c *cli.Context) error {

	// Get RP client
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

	// Get oracle DAO members
	members, err := staderClient.TNDAOMembers()
	if err != nil {
		return err
	}

	// Print & return
	if len(members.Members) > 0 {
		fmt.Printf("The oracle DAO has %d members:\n", len(members.Members))
		fmt.Println("")
	} else {
		fmt.Println("The oracle DAO does not have any members yet.")
	}
	for _, member := range members.Members {
		fmt.Printf("--------------------\n")
		fmt.Printf("\n")
		fmt.Printf("Member ID:            %s\n", member.ID)
		fmt.Printf("URL:                  %s\n", member.Url)
		fmt.Printf("Node address:         %s\n", member.Address.Hex())
		fmt.Printf("Joined at:            %s\n", cliutils.GetDateTimeString(member.JoinedTime))
		fmt.Printf("Last proposal:        %s\n", cliutils.GetDateTimeString(member.LastProposalTime))
		fmt.Printf("RPL bond amount:      %.6f\n", math.RoundDown(eth.WeiToEth(member.RPLBondAmount), 6))
		fmt.Printf("Unbonded minipools:   %d\n", member.UnbondedValidatorCount)
		fmt.Printf("\n")
	}
	return nil

}
