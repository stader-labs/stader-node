/*
This work is licensed and released under GNU GPL v3 or any other later versions.
The full text of the license is below/ found at <http://www.gnu.org/licenses/>

(c) 2023 Rocket Pool Pty Ltd. Modified under GNU GPL v3. [1.5.0]

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
package node

import (
	"fmt"
	"math/big"

	"github.com/stader-labs/stader-node/stader-lib/node"
	"github.com/stader-labs/stader-node/stader-lib/stader"

	"github.com/docker/docker/client"
	"github.com/ethereum/go-ethereum/common"
	"github.com/urfave/cli"

	"github.com/stader-labs/stader-node/shared/services"
	"github.com/stader-labs/stader-node/shared/services/beacon"
	"github.com/stader-labs/stader-node/shared/services/config"
	staderService "github.com/stader-labs/stader-node/shared/services/stader"
	"github.com/stader-labs/stader-node/shared/services/wallet"
	"github.com/stader-labs/stader-node/shared/utils/eth1"
	"github.com/stader-labs/stader-node/shared/utils/log"
	staderUtils "github.com/stader-labs/stader-node/shared/utils/stdr"
	"github.com/stader-labs/stader-node/shared/utils/validator"
)

// Manage fee recipient task
type manageFeeRecipient struct {
	c     *cli.Context
	log   log.ColorLogger
	cfg   *config.StaderConfig
	w     *wallet.Wallet
	prn   *stader.PermissionlessNodeRegistryContractManager
	vf    *stader.VaultFactoryContractManager
	pp    *stader.PermissionlessPoolContractManager
	sdcfg *stader.StaderConfigContractManager
	d     *client.Client
	bc    beacon.Client
}

// Create manage fee recipient task
func newManageFeeRecipient(c *cli.Context, logger log.ColorLogger) (*manageFeeRecipient, error) {

	// Get services
	cfg, err := services.GetConfig(c)
	if err != nil {
		return nil, err
	}
	w, err := services.GetWallet(c)
	if err != nil {
		return nil, err
	}
	prn, err := services.GetPermissionlessNodeRegistry(c)
	if err != nil {
		return nil, err
	}
	vf, err := services.GetVaultFactory(c)
	if err != nil {
		return nil, err
	}
	pp, err := services.GetPermissionlessPoolContract(c)
	if err != nil {
		return nil, err
	}
	d, err := services.GetDocker(c)
	if err != nil {
		return nil, err
	}
	bc, err := services.GetBeaconClient(c)
	if err != nil {
		return nil, err
	}
	sdcfg, err := services.GetStaderConfigContract(c)
	if err != nil {
		return nil, err
	}

	// Return task
	return &manageFeeRecipient{
		c:     c,
		log:   logger,
		cfg:   cfg,
		w:     w,
		prn:   prn,
		vf:    vf,
		pp:    pp,
		d:     d,
		bc:    bc,
		sdcfg: sdcfg,
	}, nil

}

// Manage fee recipient
func (m *manageFeeRecipient) run() error {

	// Wait for eth client to sync
	if err := services.WaitEthClientSynced(m.c, true); err != nil {
		return err
	}

	// Get node account
	nodeAccount, err := m.w.GetNodeAccount()
	if err != nil {
		return err
	}

	currentBlock, err := eth1.GetCurrentBlockNumber(m.c)
	if err != nil {
		return fmt.Errorf("error GetCurrentBlockNumber: %w", err)
	}

	pnr, err := services.GetPermissionlessNodeRegistry(m.c)
	if err != nil {
		return fmt.Errorf("error GetPermissionlessNodeRegistry: %w", err)
	}

	operatorID, err := node.GetOperatorId(pnr, nodeAccount.Address, nil)
	if err != nil {
		return fmt.Errorf("error GetOperatorId: %w", err)
	}

	lastChangeBlock, err := node.GetSocializingPoolStateChangeBlock(pnr, operatorID, nil)
	if err != nil {
		return fmt.Errorf("error GetSocializingPoolStateChangeBlock: %w", err)
	}

	nextUpdatableBlock := lastChangeBlock.Add(lastChangeBlock, big.NewInt(blocksPerThreeEpoch)).Uint64()

	m.log.Printlnf("CurrentBlock: %d, we'll update file at %d block if possible\n", currentBlock, nextUpdatableBlock)

	// Get the fee recipient info for the node
	feeRecipientInfo, err := staderUtils.GetFeeRecipientInfo(m.prn, m.vf, m.sdcfg, nodeAccount.Address, nil)
	if err != nil {
		return fmt.Errorf("error getting fee recipient info: %w", err)
	}

	updatable := currentBlock > nextUpdatableBlock
	// Get the correct fee recipient address
	var correctFeeRecipient common.Address
	if feeRecipientInfo.IsInSocializingPool {
		if updatable {
			correctFeeRecipient = feeRecipientInfo.SocializingPoolAddress
		} else {
			correctFeeRecipient = feeRecipientInfo.FeeDistributorAddress
		}
	} else {
		if updatable {
			correctFeeRecipient = feeRecipientInfo.FeeDistributorAddress
		} else {
			correctFeeRecipient = feeRecipientInfo.SocializingPoolAddress
		}
	}

	// Check if the VC is using the correct fee recipient
	fileExists, correctAddress, err := staderService.CheckFeeRecipientFile(correctFeeRecipient, m.cfg)
	if err != nil {
		return fmt.Errorf("error validating fee recipient files: %w", err)
	}

	if !fileExists {
		m.log.Println("Fee recipient files don't all exist, regenerating...")
	} else if !correctAddress {
		m.log.Printlnf("WARNING: Fee recipient files did not contain the correct fee recipient of %s, regenerating...", correctFeeRecipient.Hex())
	} else {
		// Files are all correct, return.
		m.log.Printlnf("Fee recipient files are all correct, no action required.")
		return nil
	}

	// Regenerate the fee recipient files
	err = staderService.UpdateFeeRecipientFile(correctFeeRecipient, m.cfg)
	if err != nil {
		m.log.Println("***ERROR***")
		m.log.Printlnf("Error updating fee recipient files: %s", err.Error())
		m.log.Println("Shutting down the validator client for safety to prevent you from being penalized...")

		err = validator.StopValidator(m.cfg, m.bc, &m.log, m.d)
		if err != nil {
			return fmt.Errorf("error stopping validator client: %w", err)
		}
		return nil
	}

	// Restart the VC
	m.log.Println("Fee recipient files updated successfully! Restarting validator client...")
	err = validator.RestartValidator(m.cfg, m.bc, &m.log, m.d)
	if err != nil {
		return fmt.Errorf("error restarting validator client: %w", err)
	}

	// Log & return
	m.log.Println("Successfully restarted, you are now validating safely.")
	return nil

}
