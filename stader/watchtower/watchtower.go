package watchtower

import (
	"math/rand"
	"net/http"
	"sync"
	"time"

	"github.com/fatih/color"
	"github.com/urfave/cli"

	"github.com/stader-labs/stader-node/shared/services"
	"github.com/stader-labs/stader-node/shared/utils/log"
	"github.com/stader-labs/stader-node/stader/watchtower/collectors"
)

// Config
var minTasksInterval, _ = time.ParseDuration("4m")
var maxTasksInterval, _ = time.ParseDuration("6m")

const (
	MaxConcurrentEth1Requests = 200

	ErrorColor   = color.FgRed
	MetricsColor = color.FgHiYellow
)

// Register watchtower command
func RegisterCommands(app *cli.App, name string, aliases []string) {
	app.Commands = append(app.Commands, cli.Command{
		Name:    name,
		Aliases: aliases,
		Usage:   "Run Rocket Pool watchtower activity daemon",
		Action: func(c *cli.Context) error {
			return run(c)
		},
	})
}

// Run daemon
func run(c *cli.Context) error {

	// Configure
	configureHTTP()

	// Wait until node is registered
	if err := services.WaitNodeRegistered(c, true); err != nil {
		return err
	}

	// Initialize the scrub metrics reporter
	scrubCollector := collectors.NewScrubCollector()

	// Initialize error logger
	errorLog := log.NewColorLogger(ErrorColor)

	/*processPenalties, err := newProcessPenalties(c, log.NewColorLogger(ProcessPenaltiesColor), errorLog)
	if err != nil {
		return fmt.Errorf("error during penalties check: %w", err)
	}*/

	intervalDelta := maxTasksInterval - minTasksInterval
	secondsDelta := intervalDelta.Seconds()

	// Wait group to handle the various threads
	wg := new(sync.WaitGroup)
	wg.Add(2)

	// Run task loop
	go func() {
		for {
			// Randomize the next interval
			randomSeconds := rand.Intn(int(secondsDelta))
			interval := time.Duration(randomSeconds)*time.Second + minTasksInterval

			// Check the EC status
			err := services.WaitEthClientSynced(c, false) // Force refresh the primary / fallback EC status
			if err != nil {
				errorLog.Println(err)
			} else {
				// Check the BC status
				err := services.WaitBeaconClientSynced(c, false) // Force refresh the primary / fallback BC status
				if err != nil {
					errorLog.Println(err)
				} else {
				}
			}
			time.Sleep(interval)
		}
		wg.Done()
	}()

	// Run metrics loop
	go func() {
		err := runMetricsServer(c, log.NewColorLogger(MetricsColor), scrubCollector)
		if err != nil {
			errorLog.Println(err)
		}
		wg.Done()
	}()

	// Wait for both threads to stop
	wg.Wait()
	return nil
}

// Configure HTTP transport settings
func configureHTTP() {

	// The watchtower daemon makes a large number of concurrent RPC requests to the Eth1 client
	// The HTTP transport is set to cache connections for future re-use equal to the maximum expected number of concurrent requests
	// This prevents issues related to memory consumption and address allowance from repeatedly opening and closing connections
	http.DefaultTransport.(*http.Transport).MaxIdleConnsPerHost = MaxConcurrentEth1Requests

}
