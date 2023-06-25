/*
This work is licensed and released under GNU GPL v3 or any other later versions.
The full text of the license is below/ found at <http://www.gnu.org/licenses/>

(c) 2023 Rocket Pool Pty Ltd. Modified under GNU GPL v3. [1.0.0]

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
package types

import (
	"fmt"

	"encoding/hex"

	"github.com/stader-labs/stader-node/stader-lib/utils/json"
)

// Validator pubkey
const ValidatorPubkeyLength = 48 // bytes
type ValidatorPubkey [ValidatorPubkeyLength]byte

// Bytes conversion
func (v ValidatorPubkey) Bytes() []byte {
	return v[:]
}
func BytesToValidatorPubkey(value []byte) ValidatorPubkey {
	var pubkey ValidatorPubkey
	copy(pubkey[:], value)
	return pubkey
}

// String conversion
func (v ValidatorPubkey) Hex() string {
	return hex.EncodeToString(v.Bytes())
}
func (v ValidatorPubkey) String() string {
	return v.Hex()
}
func HexToValidatorPubkey(value string) (ValidatorPubkey, error) {
	pubkey := make([]byte, ValidatorPubkeyLength)
	if len(value) != hex.EncodedLen(ValidatorPubkeyLength) {
		return ValidatorPubkey{}, fmt.Errorf("Invalid validator public key hex string %s: invalid length %d", value, len(value))
	}
	if _, err := hex.Decode(pubkey, []byte(value)); err != nil {
		return ValidatorPubkey{}, err
	}
	return BytesToValidatorPubkey(pubkey), nil
}

// JSON encoding
func (v ValidatorPubkey) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.Hex())
}
func (v *ValidatorPubkey) UnmarshalJSON(data []byte) error {
	var dataStr string
	if err := json.Unmarshal(data, &dataStr); err != nil {
		return err
	}
	pubkey, err := HexToValidatorPubkey(dataStr)
	if err == nil {
		*v = pubkey
	}
	return err
}

// Validator signature
const ValidatorSignatureLength = 96 // bytes
type ValidatorSignature [ValidatorSignatureLength]byte

// Bytes conversion
func (v ValidatorSignature) Bytes() []byte {
	return v[:]
}
func BytesToValidatorSignature(value []byte) ValidatorSignature {
	var signature ValidatorSignature
	copy(signature[:], value)
	return signature
}

// String conversion
func (v ValidatorSignature) Hex() string {
	return hex.EncodeToString(v.Bytes())
}
func (v ValidatorSignature) String() string {
	return v.Hex()
}
func HexToValidatorSignature(value string) (ValidatorSignature, error) {
	signature := make([]byte, ValidatorSignatureLength)
	if len(value) != hex.EncodedLen(ValidatorSignatureLength) {
		return ValidatorSignature{}, fmt.Errorf("Invalid validator signature hex string %s: invalid length %d", value, len(value))
	}
	if _, err := hex.Decode(signature, []byte(value)); err != nil {
		return ValidatorSignature{}, err
	}
	return BytesToValidatorSignature(signature), nil
}

// JSON encoding
func (v ValidatorSignature) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.Hex())
}
func (v *ValidatorSignature) UnmarshalJSON(data []byte) error {
	var dataStr string
	if err := json.Unmarshal(data, &dataStr); err != nil {
		return err
	}
	signature, err := HexToValidatorSignature(dataStr)
	if err == nil {
		*v = signature
	}
	return err
}
