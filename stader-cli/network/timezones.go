package network

import (
	"fmt"
	"sort"

	"github.com/urfave/cli"

	"github.com/stader-labs/stader-node/shared/services/stader"
	cliutils "github.com/stader-labs/stader-node/shared/utils/cli"
)

func getTimezones(c *cli.Context) error {

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

	// Get the timezone map
	response, err := staderClient.TimezoneMap()
	if err != nil {
		return err
	}

	// Sort it by the timezone name
	var maxNameLength int
	timezoneNames := make([]string, 0, len(response.TimezoneCounts))
	for timezoneName := range response.TimezoneCounts {
		if timezoneName != "Other" {
			timezoneNames = append(timezoneNames, timezoneName)
			nameLength := len(timezoneName) + 2
			if nameLength > maxNameLength {
				maxNameLength = nameLength
			}
		}
	}
	sort.Strings(timezoneNames)

	fmt.Printf("There are currently %d nodes across %d timezones.\n\n", response.NodeTotal, response.TimezoneTotal)

	for _, timezoneName := range timezoneNames {
		fmt.Printf("%-*s%d\n", maxNameLength, timezoneName+":", response.TimezoneCounts[timezoneName])
	}
	fmt.Printf("%-*s%d\n", maxNameLength, "Other:", response.TimezoneCounts["Other"])

	return nil

}
