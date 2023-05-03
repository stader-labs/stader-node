/*
This work is licensed and released under GNU GPL v3 or any other later versions.
The full text of the license is below/ found at <http://www.gnu.org/licenses/>

(c) 2023 Rocket Pool Pty Ltd. Modified under GNU GPL v3. [0.4.0-beta]

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
package validator

import (
	"context"
	"errors"
	"fmt"
	"os"
	"os/exec"
	"strings"
	"time"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/client"
	"github.com/stader-labs/stader-node/shared/services/beacon"
	"github.com/stader-labs/stader-node/shared/services/config"
	"github.com/stader-labs/stader-node/shared/utils/log"
)

// Settings
const ValidatorContainerSuffix = "_validator"
const BeaconContainerSuffix = "_eth2"

var validatorRestartTimeout, _ = time.ParseDuration("5s")

// Restart validator process
func RestartValidator(cfg *config.StaderConfig, bc beacon.Client, log *log.ColorLogger, d *client.Client) error {

	// Restart validator container
	if !cfg.IsNativeMode {

		// Get validator container name & client type label
		var containerName string
		var clientTypeLabel string
		if cfg.StaderNode.ProjectName.Value == "" {
			return errors.New("Stader docker project name not set")
		}
		clientType, _ := bc.GetClientType()
		switch clientType {
		case beacon.SplitProcess:
			containerName = cfg.StaderNode.ProjectName.Value.(string) + ValidatorContainerSuffix
			clientTypeLabel = "validator"
		case beacon.SingleProcess:
			containerName = cfg.StaderNode.ProjectName.Value.(string) + BeaconContainerSuffix
			clientTypeLabel = "beacon"
		default:
			return fmt.Errorf("Can't restart the validator, unknown client type '%d'", clientType)
		}

		// Log
		if log != nil {
			log.Printlnf("Restarting %s container (%s)...", clientTypeLabel, containerName)
		}

		// Get all containers
		containers, err := d.ContainerList(context.Background(), types.ContainerListOptions{All: true})
		if err != nil {
			return fmt.Errorf("Could not get docker containers: %w", err)
		}

		// Get validator container ID
		var validatorContainerId string
		for _, container := range containers {
			if container.Names[0] == "/"+containerName {
				validatorContainerId = container.ID
				break
			}
		}
		if validatorContainerId == "" {
			return errors.New("Validator container not found")
		}

		// Restart validator container
		if err := d.ContainerRestart(context.Background(), validatorContainerId, &validatorRestartTimeout); err != nil {
			return fmt.Errorf("Could not restart validator container: %w", err)
		}

		// Restart external validator process
	} else {

		// Get validator restart command
		restartCommand := os.ExpandEnv(cfg.Native.ValidatorRestartCommand.Value.(string))

		// Log
		if log != nil {
			log.Printlnf("Restarting validator process with command '%s'...", restartCommand)
		}

		// Run validator restart command bound to os stdout/stderr
		cmd := exec.Command(restartCommand)
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
		if err := cmd.Run(); err != nil {
			return fmt.Errorf("Could not restart validator process: %w", err)
		}

	}

	// Log & return
	if log != nil {
		log.Println("Successfully restarted validator")
	}
	return nil

}

// Stops the validator process
func StopValidator(cfg *config.StaderConfig, bc beacon.Client, log *log.ColorLogger, d *client.Client) error {

	// Stop validator container
	if !cfg.IsNativeMode {

		// Get validator container name & client type label
		var containerName string
		var clientTypeLabel string
		if cfg.StaderNode.ProjectName.Value == "" {
			return errors.New("Stader docker project name not set")
		}
		clientType, _ := bc.GetClientType()
		switch clientType {
		case beacon.SplitProcess:
			containerName = cfg.StaderNode.ProjectName.Value.(string) + ValidatorContainerSuffix
			clientTypeLabel = "validator"
		case beacon.SingleProcess:
			containerName = cfg.StaderNode.ProjectName.Value.(string) + BeaconContainerSuffix
			clientTypeLabel = "beacon"
		default:
			return fmt.Errorf("Can't stop the validator, unknown client type '%d'", clientType)
		}

		// Log
		if log != nil {
			log.Printlnf("Stopping %s container (%s)...", clientTypeLabel, containerName)
		}

		// Get all containers
		containers, err := d.ContainerList(context.Background(), types.ContainerListOptions{All: true})
		if err != nil {
			return fmt.Errorf("Could not get docker containers: %w", err)
		}

		// Get validator container ID
		var validatorContainerId string
		for _, container := range containers {
			if container.Names[0] == "/"+containerName {
				validatorContainerId = container.ID
				break
			}
		}
		if validatorContainerId == "" {
			// TODO: return here if the container doesn't exist? Is erroring out necessary?
			return fmt.Errorf("Validator container %s not found", containerName)
		}

		// Stop validator container
		if err := d.ContainerPause(context.Background(), validatorContainerId); err != nil {
			if strings.Contains(err.Error(), "is not running") {
				// Handle situations where the container is already stopped
				if log != nil {
					log.Printlnf("Validator container %s was not running.", containerName)
				}
				return nil
			}
			return fmt.Errorf("Could not stop validator container %s: %w", containerName, err)
		}

	} else {
		// Stop external validator process

		// Get validator stop command
		stopCommand := os.ExpandEnv(cfg.Native.ValidatorStopCommand.Value.(string))

		// Log
		if log != nil {
			log.Printlnf("Stopping validator process with command '%s'...", stopCommand)
		}

		// Run validator stop command bound to os stdout/stderr
		cmd := exec.Command(stopCommand)
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
		if err := cmd.Run(); err != nil {
			return fmt.Errorf("Could not stop validator process: %w", err)
		}

	}

	// Log & return
	if log != nil {
		log.Println("Successfully stopped validator")
	}
	return nil

}
