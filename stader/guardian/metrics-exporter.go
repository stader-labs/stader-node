/*
This work is licensed and released under GNU GPL v3 or any other later versions.
The full text of the license is below/ found at <http://www.gnu.org/licenses/>

(c) 2023 Rocket Pool Pty Ltd. Modified under GNU GPL v3. [0.3.0-beta]

This program is free software: you can redistribute it and/or modify
it under the terms of the GNU General Public License as published by
the Free Software Foundation, either version 3 of the License, or
(at your option) any later version.

This program is distributed in the hope that it will be useful,
but WITHOUT ANY WARRANTY; without even the implied warranty of
MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
GNU General Public License for more details.

You should have received a copy of the GNU General Public License
along with this program.  If not, see <http://www.gnu.org/licenses/>.
*/
package guardian

import (
	"fmt"
	"github.com/stader-labs/stader-node/stader/guardian/collector"
	"net/http"
	"os"
	"strings"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/stader-labs/stader-node/shared/services"
	"github.com/stader-labs/stader-node/shared/utils/log"
	"github.com/urfave/cli"
)

func runMetricsServer(c *cli.Context, logger log.ColorLogger, stateLocker *collector.StateLocker) error {

	// Get services
	cfg, err := services.GetConfig(c)
	if err != nil {
		return err
	}
	w, err := services.GetWallet(c)
	if err != nil {
		return err
	}
	bc, err := services.GetBeaconClient(c)
	if err != nil {
		return err
	}
	ec, err := services.GetEthClient(c)
	if err != nil {
		return err
	}

	nodeAccount, err := w.GetNodeAccount()
	if err != nil {
		return err
	}

	// Return if metrics are disabled
	if cfg.EnableMetrics.Value == false {
		if strings.ToLower(os.Getenv("ENABLE_METRICS")) == "true" {
			logger.Printlnf("ENABLE_METRICS override set to true, will start Metrics exporter anyway!")
		} else {
			return nil
		}
	}

	beaconCollector := collector.NewBeaconCollector(bc, ec, nodeAccount.Address, stateLocker)
	networkCollector := collector.NewNetworkCollector(bc, ec, nodeAccount.Address, stateLocker)

	// Set up Prometheus
	registry := prometheus.NewRegistry()
	registry.MustRegister(beaconCollector)
	registry.MustRegister(networkCollector)

	handler := promhttp.HandlerFor(registry, promhttp.HandlerOpts{})

	// Start the HTTP server
	metricsAddress := c.GlobalString("metricsAddress")
	metricsPort := c.GlobalUint("metricsPort")
	logger.Printlnf("Starting metrics exporter on %s:%d.", metricsAddress, metricsPort)
	metricsPath := "/metrics"
	http.Handle(metricsPath, handler)
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(`<html>
            <head><title>Stader Guardian Metrics Exporter</title></head>
            <body>
            <h1>Stader Guardian Metrics Exporter</h1>
            <p><a href='` + metricsPath + `'>Metrics</a></p>
            </body>
            </html>`,
		))
	})
	err = http.ListenAndServe(fmt.Sprintf("%s:%d", metricsAddress, metricsPort), nil)
	if err != nil {
		return fmt.Errorf("Error running HTTP server: %w", err)
	}

	return nil

}
