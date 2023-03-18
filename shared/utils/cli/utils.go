package cli

// ROCKETPOOL-OWNED

import (
	"fmt"
	"strings"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/stader-labs/stader-node/shared/services/stader"
	cfgtypes "github.com/stader-labs/stader-node/shared/types/config"
)

const colorReset string = "\033[0m"
const colorRed string = "\033[31m"
const colorGreen string = "\033[32m"
const colorYellow string = "\033[33m"
const colorLightBlue string = "\033[36m"

// Print a TX's details to the console.
func PrintTransactionHash(staderClient *stader.Client, hash common.Hash) {

	finalMessage := "Waiting for the transaction to be included in a block... you may wait here for it, or press CTRL+C to exit and return to the terminal.\n\n"
	printTransactionHashImpl(staderClient, hash, finalMessage)

}

// Print a TX's details to the console, but inform the user NOT to cancel it.
func PrintTransactionHashNoCancel(staderClient *stader.Client, hash common.Hash) {

	finalMessage := "Waiting for the transaction to be included in a block... **DO NOT EXIT!** This transaction is one of several that must be completed.\n\n"
	printTransactionHashImpl(staderClient, hash, finalMessage)

}

// Print a warning to the console if the user set a custom nonce, but this operation involves multiple transactions
func PrintMultiTransactionNonceWarning() {

	fmt.Printf("%sNOTE: You have specified the `nonce` flag to indicate a custom nonce for this transaction.\n"+
		"However, this operation requires multiple transactions.\n"+
		"Stader will use your custom value as a basis, and increment it for each additional transaction.\n"+
		"If you have multiple pending transactions, this MAY OVERRIDE more than the one that you specified.%s\n\n", colorYellow, colorReset)

}

// Implementation of PrintTransactionHash and PrintTransactionHashNoCancel
func printTransactionHashImpl(staderClient *stader.Client, hash common.Hash, finalMessage string) {

	cfg, isNew, err := staderClient.LoadConfig()
	if err != nil {
		fmt.Printf("Warning: couldn't read config file so the transaction URL will be unavailable (%s).\n", err)
		return
	}

	if isNew {
		fmt.Print("Settings file not found. Please run `stader-cli service config` to set up your Smartnode.")
		return
	}

	txWatchUrl := cfg.Smartnode.GetTxWatchUrl()
	hashString := hash.String()

	fmt.Printf("Transaction has been submitted with hash %s.\n", hashString)
	if txWatchUrl != "" {
		fmt.Printf("You may follow its progress by visiting:\n")
		fmt.Printf("%s/%s\n\n", txWatchUrl, hashString)
	}
	fmt.Print(finalMessage)

}

// Convert a Unix datetime to a string, or `---` if it's zero
func GetDateTimeString(dateTime uint64) string {
	timeString := time.Unix(int64(dateTime), 0).Format(time.RFC822)
	if dateTime == 0 {
		timeString = "---"
	}
	return timeString
}

// Gets the hex string of an address, or "none" if it was the 0x0 address
func GetPrettyAddress(address common.Address) string {
	addressString := address.Hex()
	if addressString == "0x0000000000000000000000000000000000000000" {
		return "<none>"
	}
	return addressString
}

// Prints an error in a prettier format, removing the "stack trace" if it represents
// a contract revert message
func PrettyPrintError(err error) {
	errorMessage := err.Error()
	prettyErr := errorMessage
	if strings.Contains(errorMessage, "execution reverted:") {
		elements := strings.Split(errorMessage, ":")
		firstMessage := strings.TrimSpace(elements[0])
		secondMessage := strings.TrimSpace(elements[len(elements)-1])
		prettyErr = fmt.Sprintf("%s: %s", firstMessage, secondMessage)
	}
	fmt.Println(prettyErr)
}

// Prints an error message when the Beacon client is not using the deposit contract address that Stader expects
func PrintDepositMismatchError(rpNetwork, beaconNetwork uint64, rpDepositAddress, beaconDepositAddress common.Address) {
	fmt.Printf("%s***ALERT***\n", colorRed)
	fmt.Println("YOUR ETH2 CLIENT IS NOT CONNECTED TO THE SAME NETWORK THAT Stader IS USING!")
	fmt.Println("This is likely because your ETH2 client is using the wrong configuration.")
	fmt.Println("For the safety of your funds, Stader will not let you deposit your ETH until this is resolved.")
	fmt.Println()
	fmt.Println("To fix it if you are in Docker mode:")
	fmt.Println("\t1. Run 'stader-cli service install -d' to get the latest configuration")
	fmt.Println("\t2. Run 'stader-cli service stop' and 'stader-cli service start' to apply the configuration.")
	fmt.Println("If you are using Hybrid or Native mode, please correct the network flags in your ETH2 launch script.")
	fmt.Println()
	fmt.Println("Details:")
	fmt.Printf("\tStader expects deposit contract %s on chain %d.\n", rpDepositAddress.Hex(), rpNetwork)
	fmt.Printf("\tYour Beacon client is using deposit contract %s on chain %d.%s\n", beaconDepositAddress.Hex(), beaconNetwork, colorReset)
}

// Prints what network you're currently on
func PrintNetwork(staderClient *stader.Client) error {
	cfg, isNew, err := staderClient.LoadConfig()
	if err != nil {
		return fmt.Errorf("error loading global config: %w", err)
	}
	if isNew {
		return fmt.Errorf("settings file not found. Please run `stader-cli service config` to set up your Stader Node")
	}

	currentNetwork := cfg.Smartnode.Network.Value.(cfgtypes.Network)
	switch currentNetwork {
	case cfgtypes.Network_Mainnet:
		fmt.Printf("Your Smartnode is currently using the %sEthereum Mainnet.%s\n\n", colorGreen, colorReset)
	case cfgtypes.Network_Prater:
		fmt.Printf("Your Smartnode is currently using the %sGoerli Test Network.%s\n\n", colorLightBlue, colorReset)
	case cfgtypes.Network_Devnet:
		fmt.Printf("Your Smartnode is currently using the %sPrater Development Network.%s\n\n", colorYellow, colorReset)
	case cfgtypes.Network_Zhejiang:
		fmt.Printf("Your Smartnode is currently using the %sZhejiang Test Network.%s\n\n", colorYellow, colorReset)
	default:
		fmt.Printf("%sYou are on an unexpected network [%v].%s\n\n", colorYellow, currentNetwork, colorReset)
	}

	return nil
}
