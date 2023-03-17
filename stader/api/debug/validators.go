package debug

import (
	"github.com/urfave/cli"
)

// Get all minipool balance details
func ExportValidators(c *cli.Context) error {
	//
	//opts := &bind.CallOpts{}
	//
	//// Get services
	//ec, err := services.GetEthClient(c)
	//if err != nil {
	//	return err
	//}
	//rpl, err := services.GetRocketPool(c)
	//if err != nil {
	//	return err
	//}
	//bc, err := services.GetBeaconClient(c)
	//if err != nil {
	//	return err
	//}
	//
	//// Data
	//var wg1 errgroup.Group
	//var addresses []common.Address
	//var eth2Config beacon.Eth2Config
	//var beaconHead beacon.BeaconHead
	//var blockTime uint64
	//
	//// Get minipool addresses
	//wg1.Go(func() error {
	//	var err error
	//	addresses, err = minipool.GetMinipoolAddresses(rpl, opts)
	//	return err
	//})
	//
	//// Get eth2 config
	//wg1.Go(func() error {
	//	var err error
	//	eth2Config, err = bc.GetEth2Config()
	//	return err
	//})
	//
	//// Get beacon head
	//wg1.Go(func() error {
	//	var err error
	//	beaconHead, err = bc.GetBeaconHead()
	//	return err
	//})
	//
	//// Get block time
	//wg1.Go(func() error {
	//	header, err := ec.HeaderByNumber(context.Background(), opts.BlockNumber)
	//	if err == nil {
	//		blockTime = header.Time
	//	}
	//	return err
	//})
	//
	//// Wait for data
	//if err := wg1.Wait(); err != nil {
	//	return err
	//}
	//
	//// Get & check epoch at block
	//blockEpoch := eth2.EpochAt(eth2Config, blockTime)
	//if blockEpoch > beaconHead.Epoch {
	//	return fmt.Errorf("Epoch %d at block %s is higher than current epoch %d", blockEpoch, opts.BlockNumber.String(), beaconHead.Epoch)
	//}
	//
	//// Get minipool validator statuses
	//validators, err := rp.GetMinipoolValidators(rpl, bc, addresses, opts, &beacon.ValidatorStatusOptions{Epoch: &blockEpoch})
	//if err != nil {
	//	return err
	//}
	//
	//fmt.Printf("%s\t%s\t%s\t%s\t%s\t%s\t%s\t%s\t%s\t%s\n",
	//	"Minipool Address",
	//	"Validator Pub Key",
	//	"Activation Epoch",
	//	"Node Fee",
	//	"Validator Balance",
	//	"Node Balance",
	//	"User Balance",
	//	"Status",
	//	"Finalised",
	//	"Active",
	//	"Pending",
	//)
	//
	//// Load details in batches
	//for bsi := 0; bsi < len(addresses); bsi += MinipoolBalanceDetailsBatchSize {
	//
	//	// Get batch start & end index
	//	msi := bsi
	//	mei := bsi + MinipoolBalanceDetailsBatchSize
	//	if mei > len(addresses) {
	//		mei = len(addresses)
	//	}
	//
	//	// Log
	//	//t.log.Printlnf("Calculating balances for minipools %d - %d of %d...", msi + 1, mei, len(addresses))
	//
	//	// Load details
	//	var wg errgroup.Group
	//	for mi := msi; mi < mei; mi++ {
	//		mi := mi
	//		wg.Go(func() error {
	//			address := addresses[mi]
	//			validator := validators[address]
	//			err := getMinipoolBalanceDetails(rpl, address, opts, validator, eth2Config, blockEpoch)
	//			return err
	//		})
	//	}
	//	if err := wg.Wait(); err != nil {
	//		return err
	//	}
	//
	//}

	// Return
	return nil
}
