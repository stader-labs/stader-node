package odao

import (
	"fmt"
	"time"

	"github.com/rocket-pool/rocketpool-go/utils/eth"
	"github.com/urfave/cli"

	"github.com/stader-labs/stader-node/shared/services/gas"
	"github.com/stader-labs/stader-node/shared/services/stader"
	cliutils "github.com/stader-labs/stader-node/shared/utils/cli"
)

func proposeSettingMembersQuorum(c *cli.Context, quorumPercent float64) error {

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

	// Check if proposal can be made
	canPropose, err := staderClient.CanProposeTNDAOSettingMembersQuorum(quorumPercent / 100)
	if err != nil {
		return err
	}
	if !canPropose.CanPropose {
		fmt.Println("Cannot propose setting update:")
		if canPropose.ProposalCooldownActive {
			fmt.Println("The node must wait for the proposal cooldown period to pass before making another proposal.")
		}
		return nil
	}

	// Assign max fees
	err = gas.AssignMaxFeeAndLimit(canPropose.GasInfo, staderClient, c.Bool("yes"))
	if err != nil {
		return err
	}

	// Prompt for confirmation
	if !(c.Bool("yes") || cliutils.Confirm("Are you sure you want to submit this proposal?")) {
		fmt.Println("Cancelled.")
		return nil
	}

	// Submit proposal
	response, err := staderClient.ProposeTNDAOSettingMembersQuorum(quorumPercent / 100)
	if err != nil {
		return err
	}

	fmt.Printf("Submitting proposal...\n")
	cliutils.PrintTransactionHash(staderClient, response.TxHash)
	if _, err = staderClient.WaitForTransaction(response.TxHash); err != nil {
		return err
	}

	// Log & return
	fmt.Printf("Successfully submitted a members.quorum setting update proposal with ID %d.\n", response.ProposalId)
	return nil

}

func proposeSettingMembersRplBond(c *cli.Context, bondAmountEth float64) error {

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

	// Check if proposal can be made
	canPropose, err := staderClient.CanProposeTNDAOSettingMembersRplBond(eth.EthToWei(bondAmountEth))
	if err != nil {
		return err
	}
	if !canPropose.CanPropose {
		fmt.Println("Cannot propose setting update:")
		if canPropose.ProposalCooldownActive {
			fmt.Println("The node must wait for the proposal cooldown period to pass before making another proposal.")
		}
		return nil
	}

	// Assign max fees
	err = gas.AssignMaxFeeAndLimit(canPropose.GasInfo, staderClient, c.Bool("yes"))
	if err != nil {
		return err
	}

	// Prompt for confirmation
	if !(c.Bool("yes") || cliutils.Confirm("Are you sure you want to submit this proposal?")) {
		fmt.Println("Cancelled.")
		return nil
	}

	// Submit proposal
	response, err := staderClient.ProposeTNDAOSettingMembersRplBond(eth.EthToWei(bondAmountEth))
	if err != nil {
		return err
	}

	fmt.Printf("Submitting proposal...\n")
	cliutils.PrintTransactionHash(staderClient, response.TxHash)
	if _, err = staderClient.WaitForTransaction(response.TxHash); err != nil {
		return err
	}

	// Log & return
	fmt.Printf("Successfully submitted a members.rplbond setting update proposal with ID %d.\n", response.ProposalId)
	return nil

}

func proposeSettingMinipoolUnbondedMax(c *cli.Context, unbondedMinipoolMax uint64) error {

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

	// Check if proposal can be made
	canPropose, err := staderClient.CanProposeTNDAOSettingMinipoolUnbondedMax(unbondedMinipoolMax)
	if err != nil {
		return err
	}
	if !canPropose.CanPropose {
		fmt.Println("Cannot propose setting update:")
		if canPropose.ProposalCooldownActive {
			fmt.Println("The node must wait for the proposal cooldown period to pass before making another proposal.")
		}
		return nil
	}

	// Assign max fees
	err = gas.AssignMaxFeeAndLimit(canPropose.GasInfo, staderClient, c.Bool("yes"))
	if err != nil {
		return err
	}

	// Prompt for confirmation
	if !(c.Bool("yes") || cliutils.Confirm("Are you sure you want to submit this proposal?")) {
		fmt.Println("Cancelled.")
		return nil
	}

	// Submit proposal
	response, err := staderClient.ProposeTNDAOSettingMinipoolUnbondedMax(unbondedMinipoolMax)
	if err != nil {
		return err
	}

	fmt.Printf("Submitting proposal...\n")
	cliutils.PrintTransactionHash(staderClient, response.TxHash)
	if _, err = staderClient.WaitForTransaction(response.TxHash); err != nil {
		return err
	}

	// Log & return
	fmt.Printf("Successfully submitted a members.minipool.unbonded.max setting update proposal with ID %d.\n", response.ProposalId)
	return nil

}

func proposeSettingProposalCooldown(c *cli.Context, proposalCooldownTimespan string) error {

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

	// Parse the timespan
	timespan, err := time.ParseDuration(proposalCooldownTimespan)
	if err != nil {
		return fmt.Errorf("Error parsing time: %w\n", err)
	}
	seconds := uint64(timespan.Seconds())

	// Check if proposal can be made
	canPropose, err := staderClient.CanProposeTNDAOSettingProposalCooldown(seconds)
	if err != nil {
		return err
	}
	if !canPropose.CanPropose {
		fmt.Println("Cannot propose setting update:")
		if canPropose.ProposalCooldownActive {
			fmt.Println("The node must wait for the proposal cooldown period to pass before making another proposal.")
		}
		return nil
	}

	// Assign max fees
	err = gas.AssignMaxFeeAndLimit(canPropose.GasInfo, staderClient, c.Bool("yes"))
	if err != nil {
		return err
	}

	// Prompt for confirmation
	if !(c.Bool("yes") || cliutils.Confirm("Are you sure you want to submit this proposal?")) {
		fmt.Println("Cancelled.")
		return nil
	}

	// Submit proposal
	response, err := staderClient.ProposeTNDAOSettingProposalCooldown(seconds)
	if err != nil {
		return err
	}

	fmt.Printf("Submitting proposal...\n")
	cliutils.PrintTransactionHash(staderClient, response.TxHash)
	if _, err = staderClient.WaitForTransaction(response.TxHash); err != nil {
		return err
	}

	// Log & return
	fmt.Printf("Successfully submitted a proposal.cooldown setting update proposal with ID %d.\n", response.ProposalId)
	return nil

}

func proposeSettingProposalVoteTimespan(c *cli.Context, proposalVoteTimespan string) error {

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

	// Parse the timespan
	timespan, err := time.ParseDuration(proposalVoteTimespan)
	if err != nil {
		return fmt.Errorf("Error parsing time: %w\n", err)
	}
	seconds := uint64(timespan.Seconds())

	// Check if proposal can be made
	canPropose, err := staderClient.CanProposeTNDAOSettingProposalVoteTimespan(seconds)
	if err != nil {
		return err
	}
	if !canPropose.CanPropose {
		fmt.Println("Cannot propose setting update:")
		if canPropose.ProposalCooldownActive {
			fmt.Println("The node must wait for the proposal cooldown period to pass before making another proposal.")
		}
		return nil
	}

	// Assign max fees
	err = gas.AssignMaxFeeAndLimit(canPropose.GasInfo, staderClient, c.Bool("yes"))
	if err != nil {
		return err
	}

	// Prompt for confirmation
	if !(c.Bool("yes") || cliutils.Confirm("Are you sure you want to submit this proposal?")) {
		fmt.Println("Cancelled.")
		return nil
	}

	// Submit proposal
	response, err := staderClient.ProposeTNDAOSettingProposalVoteTimespan(seconds)
	if err != nil {
		return err
	}

	fmt.Printf("Submitting proposal...\n")
	cliutils.PrintTransactionHash(staderClient, response.TxHash)
	if _, err = staderClient.WaitForTransaction(response.TxHash); err != nil {
		return err
	}

	// Log & return
	fmt.Printf("Successfully submitted a proposal.vote.time setting update proposal with ID %d.\n", response.ProposalId)
	return nil

}

func proposeSettingProposalVoteDelayTimespan(c *cli.Context, proposalDelayTimespan string) error {

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

	// Parse the timespan
	timespan, err := time.ParseDuration(proposalDelayTimespan)
	if err != nil {
		return fmt.Errorf("Error parsing time: %w\n", err)
	}
	seconds := uint64(timespan.Seconds())

	// Check if proposal can be made
	canPropose, err := staderClient.CanProposeTNDAOSettingProposalVoteDelayTimespan(seconds)
	if err != nil {
		return err
	}
	if !canPropose.CanPropose {
		fmt.Println("Cannot propose setting update:")
		if canPropose.ProposalCooldownActive {
			fmt.Println("The node must wait for the proposal cooldown period to pass before making another proposal.")
		}
		return nil
	}

	// Assign max fees
	err = gas.AssignMaxFeeAndLimit(canPropose.GasInfo, staderClient, c.Bool("yes"))
	if err != nil {
		return err
	}

	// Prompt for confirmation
	if !(c.Bool("yes") || cliutils.Confirm("Are you sure you want to submit this proposal?")) {
		fmt.Println("Cancelled.")
		return nil
	}

	// Submit proposal
	response, err := staderClient.ProposeTNDAOSettingProposalVoteDelayTimespan(seconds)
	if err != nil {
		return err
	}

	fmt.Printf("Submitting proposal...\n")
	cliutils.PrintTransactionHash(staderClient, response.TxHash)
	if _, err = staderClient.WaitForTransaction(response.TxHash); err != nil {
		return err
	}

	// Log & return
	fmt.Printf("Successfully submitted a proposal.vote.delay.time setting update proposal with ID %d.\n", response.ProposalId)
	return nil

}

func proposeSettingProposalExecuteTimespan(c *cli.Context, proposalExecuteTimespan string) error {

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

	// Parse the timespan
	timespan, err := time.ParseDuration(proposalExecuteTimespan)
	if err != nil {
		return fmt.Errorf("Error parsing time: %w\n", err)
	}
	seconds := uint64(timespan.Seconds())

	// Check if proposal can be made
	canPropose, err := staderClient.CanProposeTNDAOSettingProposalExecuteTimespan(seconds)
	if err != nil {
		return err
	}
	if !canPropose.CanPropose {
		fmt.Println("Cannot propose setting update:")
		if canPropose.ProposalCooldownActive {
			fmt.Println("The node must wait for the proposal cooldown period to pass before making another proposal.")
		}
		return nil
	}

	// Assign max fees
	err = gas.AssignMaxFeeAndLimit(canPropose.GasInfo, staderClient, c.Bool("yes"))
	if err != nil {
		return err
	}

	// Prompt for confirmation
	if !(c.Bool("yes") || cliutils.Confirm("Are you sure you want to submit this proposal?")) {
		fmt.Println("Cancelled.")
		return nil
	}

	// Submit proposal
	response, err := staderClient.ProposeTNDAOSettingProposalExecuteTimespan(seconds)
	if err != nil {
		return err
	}

	fmt.Printf("Submitting proposal...\n")
	cliutils.PrintTransactionHash(staderClient, response.TxHash)
	if _, err = staderClient.WaitForTransaction(response.TxHash); err != nil {
		return err
	}

	// Log & return
	fmt.Printf("Successfully submitted a proposal.execute.time setting update proposal with ID %d.\n", response.ProposalId)
	return nil

}

func proposeSettingProposalActionTimespan(c *cli.Context, proposalActionTimespan string) error {

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

	// Parse the timespan
	timespan, err := time.ParseDuration(proposalActionTimespan)
	if err != nil {
		return fmt.Errorf("Error parsing time: %w\n", err)
	}
	seconds := uint64(timespan.Seconds())

	// Check if proposal can be made
	canPropose, err := staderClient.CanProposeTNDAOSettingProposalActionTimespan(seconds)
	if err != nil {
		return err
	}
	if !canPropose.CanPropose {
		fmt.Println("Cannot propose setting update:")
		if canPropose.ProposalCooldownActive {
			fmt.Println("The node must wait for the proposal cooldown period to pass before making another proposal.")
		}
		return nil
	}

	// Assign max fees
	err = gas.AssignMaxFeeAndLimit(canPropose.GasInfo, staderClient, c.Bool("yes"))
	if err != nil {
		return err
	}

	// Prompt for confirmation
	if !(c.Bool("yes") || cliutils.Confirm("Are you sure you want to submit this proposal?")) {
		fmt.Println("Cancelled.")
		return nil
	}

	// Submit proposal
	response, err := staderClient.ProposeTNDAOSettingProposalActionTimespan(seconds)
	if err != nil {
		return err
	}

	fmt.Printf("Submitting proposal...\n")
	cliutils.PrintTransactionHash(staderClient, response.TxHash)
	if _, err = staderClient.WaitForTransaction(response.TxHash); err != nil {
		return err
	}

	// Log & return
	fmt.Printf("Successfully submitted a proposal.action.time setting update proposal with ID %d.\n", response.ProposalId)
	return nil

}

func proposeSettingScrubPeriod(c *cli.Context, scrubPeriod string) error {

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

	// Parse the timespan
	timespan, err := time.ParseDuration(scrubPeriod)
	if err != nil {
		return fmt.Errorf("Error parsing time: %w\n", err)
	}
	seconds := uint64(timespan.Seconds())

	// Check if proposal can be made
	canPropose, err := staderClient.CanProposeTNDAOSettingScrubPeriod(seconds)
	if err != nil {
		return err
	}
	if !canPropose.CanPropose {
		fmt.Println("Cannot propose setting update:")
		if canPropose.ProposalCooldownActive {
			fmt.Println("The node must wait for the proposal cooldown period to pass before making another proposal.")
		}
		return nil
	}

	// Assign max fees
	err = gas.AssignMaxFeeAndLimit(canPropose.GasInfo, staderClient, c.Bool("yes"))
	if err != nil {
		return err
	}

	// Prompt for confirmation
	if !(c.Bool("yes") || cliutils.Confirm("Are you sure you want to submit this proposal?")) {
		fmt.Println("Cancelled.")
		return nil
	}

	// Submit proposal
	response, err := staderClient.ProposeTNDAOSettingScrubPeriod(seconds)
	if err != nil {
		return err
	}

	fmt.Printf("Submitting proposal...\n")
	cliutils.PrintTransactionHash(staderClient, response.TxHash)
	if _, err = staderClient.WaitForTransaction(response.TxHash); err != nil {
		return err
	}

	// Log & return
	fmt.Printf("Successfully submitted a minipool.scrub.period setting update proposal with ID %d.\n", response.ProposalId)
	return nil

}
