package node

import (
	"encoding/json"
	"fmt"
	"github.com/mitchellh/go-homedir"
	"github.com/stader-labs/stader-node/shared/services"
	"github.com/stader-labs/stader-node/shared/services/config"
	"github.com/stader-labs/stader-node/shared/services/wallet"
	"github.com/stader-labs/stader-node/shared/utils/log"
	"github.com/stader-labs/stader-node/shared/utils/stader"
	"github.com/urfave/cli"
	"os"
)

type MerkleProofsDownloader struct {
	c   *cli.Context
	log log.ColorLogger
	cfg *config.StaderConfig
	w   *wallet.Wallet
}

func NewMerkleProofsDownloader(c *cli.Context, logger log.ColorLogger) (*MerkleProofsDownloader, error) {
	cfg, err := services.GetConfig(c)
	if err != nil {
		return nil, err
	}
	w, err := services.GetWallet(c)
	if err != nil {
		return nil, err
	}

	return &MerkleProofsDownloader{
		c:   c,
		log: logger,
		cfg: cfg,
		w:   w,
	}, nil
}

func (m *MerkleProofsDownloader) run() error {
	// Wait for eth client to sync
	if err := services.WaitEthClientSynced(m.c, true); err != nil {
		return err
	}

	nodeAccount, err := m.w.GetNodeAccount()
	if err != nil {
		return err
	}

	allMerkleProofs, err := stader.GetAllMerkleProofsForOperator(m.c, nodeAccount.Address)
	if err != nil {
		return err
	}

	downloadedCycles := []int64{}

	for _, cycleMerkleProof := range allMerkleProofs {
		cycleMerkleProofFile := m.cfg.StaderNode.GetSpRewardCyclePath(cycleMerkleProof.Cycle, true)
		absolutePathOfProofFile, err := homedir.Expand(cycleMerkleProofFile)
		if err != nil {
			return err
		}

		_, err = os.Stat(cycleMerkleProofFile)
		if !os.IsNotExist(err) && err != nil {
			return err
		}
		if !os.IsNotExist(err) {
			m.log.Printlnf("Merkle proof for cycle %d already exists, skipping", cycleMerkleProof.Cycle)
			continue
		}

		m.log.Printlnf("Downloading merkle proof for cycle %d", cycleMerkleProof.Cycle)
		file, err := os.Create(absolutePathOfProofFile)
		if err != nil {
			return err
		}

		encoder := json.NewEncoder(file)
		err = encoder.Encode(cycleMerkleProof)
		if err != nil {
			return fmt.Errorf("error encoding JSON: %v", err)
		}

		downloadedCycles = append(downloadedCycles, cycleMerkleProof.Cycle)
	}

	if len(downloadedCycles) == 0 {
		m.log.Printlnf("No merkle proofs to download")
		return nil
	} else {
		m.log.Printlnf("Downloaded merkle proofs for cycles: %v", downloadedCycles)
	}

	return nil
}
