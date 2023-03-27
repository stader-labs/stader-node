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
package cli

import (
	"encoding/hex"
	"fmt"
	"math/big"
	"regexp"
	"strconv"
	"strings"

	"github.com/ethereum/go-ethereum/common"
	"github.com/stader-labs/stader-node/stader-lib/types"
	"github.com/tyler-smith/go-bip39"
	"github.com/urfave/cli"

	"github.com/stader-labs/stader-node/shared/services/passwords"
	hexutils "github.com/stader-labs/stader-node/shared/utils/hex"
)

//
// General types
//

// Validate command argument count
func ValidateArgCount(c *cli.Context, count int) error {
	if len(c.Args()) != count {
		return fmt.Errorf("incorrect argument count; usage: %s", c.Command.UsageText)
	}
	return nil
}

// Validate a big int
func ValidateBigInt(name, value string) (*big.Int, error) {
	val, success := big.NewInt(0).SetString(value, 0)
	if !success {
		return nil, fmt.Errorf("invalid %s '%s'", name, value)
	}
	return val, nil
}

// Validate a boolean value
func ValidateBool(name, value string) (bool, error) {
	val := strings.ToLower(value)
	if !(val == "true" || val == "yes" || val == "false" || val == "no") {
		return false, fmt.Errorf("invalid %s '%s' - valid values are 'true', 'yes', 'false' and 'no'", name, value)
	}
	if val == "true" || val == "yes" {
		return true, nil
	}
	return false, nil
}

// Validate an unsigned integer value
func ValidateUint(name, value string) (uint64, error) {
	val, err := strconv.ParseUint(value, 10, 64)
	if err != nil {
		return 0, fmt.Errorf("invalid %s '%s'", name, value)
	}
	return val, nil
}

// Validate an address
func ValidateAddress(name, value string) (common.Address, error) {
	if !common.IsHexAddress(value) {
		return common.Address{}, fmt.Errorf("invalid %s '%s'", name, value)
	}
	return common.HexToAddress(value), nil
}

// Validate a wei amount
func ValidateWeiAmount(name, value string) (*big.Int, error) {
	val := new(big.Int)
	if _, ok := val.SetString(value, 10); !ok {
		return nil, fmt.Errorf("invalid %s '%s'", name, value)
	}
	return val, nil
}

// Validate an ether amount
func ValidateEthAmount(name, value string) (float64, error) {
	val, err := strconv.ParseFloat(value, 64)
	if err != nil {
		return 0, fmt.Errorf("invalid %s '%s'", name, value)
	}
	return val, nil
}

// Validate a fraction
func ValidateFraction(name, value string) (float64, error) {
	val, err := strconv.ParseFloat(value, 64)
	if err != nil || val < 0 || val > 1 {
		return 0, fmt.Errorf("invalid %s '%s' - must be a number between 0 and 1", name, value)
	}
	return val, nil
}

// Validate a percentage
func ValidatePercentage(name, value string) (float64, error) {
	val, err := strconv.ParseFloat(value, 64)
	if err != nil || val < 0 || val > 100 {
		return 0, fmt.Errorf("invalid %s '%s' - must be a number between 0 and 100", name, value)
	}
	return val, nil
}

// Validate a token type
func ValidateTokenType(name, value string) (string, error) {
	val := strings.ToLower(value)
	if !(val == "eth" || val == "sd" || val == "ethx") {
		return "", fmt.Errorf("invalid %s '%s' - valid types are 'ETH', 'SD', and 'EthX'", name, value)
	}
	return val, nil
}

//
// Command specific types
//

// Validate a positive unsigned integer value
func ValidatePositiveUint(name, value string) (uint64, error) {
	val, err := ValidateUint(name, value)
	if err != nil {
		return 0, err
	}
	if val == 0 {
		return 0, fmt.Errorf("invalid %s '%s' - must be greater than 0", name, value)
	}
	return val, nil
}

// Validate a positive wei amount
func ValidatePositiveWeiAmount(name, value string) (*big.Int, error) {
	val, err := ValidateWeiAmount(name, value)
	if err != nil {
		return nil, err
	}
	if val.Cmp(big.NewInt(0)) < 1 {
		return nil, fmt.Errorf("invalid %s '%s' - must be greater than 0", name, value)
	}
	return val, nil
}

// Validate a positive or zero wei amount
func ValidatePositiveOrZeroWeiAmount(name, value string) (*big.Int, error) {
	val, err := ValidateWeiAmount(name, value)
	if err != nil {
		return nil, err
	}
	if val.Cmp(big.NewInt(0)) < 0 {
		return nil, fmt.Errorf("invalid %s '%s' - must be greater or equal to 0", name, value)
	}
	return val, nil
}

// Validate a deposit amount in wei
func ValidateDepositWeiAmount(name, value string) (*big.Int, error) {
	val, err := ValidateWeiAmount(name, value)
	if err != nil {
		return nil, err
	}
	if ether := strings.Repeat("0", 18); !(val.String() == "0" || val.String() == "4"+ether || val.String() == "32"+ether) {
		return nil, fmt.Errorf("invalid %s '%s' - valid values are 0, 4 and 32 ether", name, value)
	}
	return val, nil
}

// Validate a positive ether amount
func ValidatePositiveEthAmount(name, value string) (float64, error) {
	val, err := ValidateEthAmount(name, value)
	if err != nil {
		return 0, err
	}
	if val <= 0 {
		return 0, fmt.Errorf("invalid %s '%s' - must be greater than 0", name, value)
	}
	return val, nil
}

// Validate a deposit amount in ether
func ValidateDepositEthAmount(name, value string) (float64, error) {
	val, err := ValidateEthAmount(name, value)
	if err != nil {
		return 0, err
	}
	if !(val == 0 || val == 4 || val == 32) {
		return 0, fmt.Errorf("invalid %s '%s' - valid values are 0, 16 and 32 ether", name, value)
	}
	return val, nil
}

// Validate a node password
func ValidateNodePassword(name, value string) (string, error) {
	if len(value) < passwords.MinPasswordLength {
		return "", fmt.Errorf("invalid %s '%s' - must be at least %d characters long", name, value, passwords.MinPasswordLength)
	}
	return value, nil
}

// Validate a wallet mnemonic phrase
func ValidateWalletMnemonic(name, value string) (string, error) {
	if !bip39.IsMnemonicValid(value) {
		return "", fmt.Errorf("invalid %s '%s'", name, value)
	}
	return value, nil
}

// Validate a timezone location
func ValidateTimezoneLocation(name, value string) (string, error) {
	if !regexp.MustCompile("^([a-zA-Z_]{2,}\\/)+[a-zA-Z_]{2,}$").MatchString(value) {
		return "", fmt.Errorf("invalid %s '%s' - must be in the format 'Country/City'", name, value)
	}
	return value, nil
}

// Validate a transaction hash
func ValidateTxHash(name, value string) (common.Hash, error) {

	// Remove a 0x prefix if present
	if strings.HasPrefix(value, "0x") {
		value = value[2:]
	}

	// Hash should be 64 characters long
	if len(value) != hex.EncodedLen(common.HashLength) {
		return common.Hash{}, fmt.Errorf("invalid %s '%s': it must have 64 characters", name, value)
	}

	// Try to parse the string (removing the prefix)
	bytes, err := hex.DecodeString(value)
	if err != nil {
		return common.Hash{}, fmt.Errorf("invalid %s '%s': %w", name, value, err)
	}
	hash := common.BytesToHash(bytes)

	return hash, nil

}

// Validate a validator pubkey
func ValidatePubkey(name, value string) (types.ValidatorPubkey, error) {
	pubkey, err := types.HexToValidatorPubkey(hexutils.RemovePrefix(value))
	if err != nil {
		return types.ValidatorPubkey{}, fmt.Errorf("invalid %s '%s': %w", name, value, err)
	}
	return pubkey, nil
}
