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
	"github.com/stader-labs/stader-node/shared/types/eth2"
	"github.com/stader-labs/stader-node/stader-lib/types"
	eth2types "github.com/wealdtech/go-eth2-types/v2"
)

// Get a voluntary exit message signature for a given validator key and index
func GetSignedExitMessage(validatorKey *eth2types.BLSPrivateKey, validatorIndex uint64, epoch uint64, signatureDomain []byte) (types.ValidatorSignature, [32]byte, error) {

	// Build voluntary exit message
	exitMessage := eth2.VoluntaryExit{
		Epoch:          epoch,
		ValidatorIndex: validatorIndex,
	}

	// Get object root
	or, err := exitMessage.HashTreeRoot()
	if err != nil {
		return types.ValidatorSignature{}, [32]byte{}, err
	}

	// Get signing root
	sr := eth2.SigningRoot{
		ObjectRoot: or[:],
		Domain:     signatureDomain,
	}

	srHash, err := sr.HashTreeRoot()
	if err != nil {
		return types.ValidatorSignature{}, [32]byte{}, nil
	}

	// Sign message
	signature := validatorKey.Sign(srHash[:]).Marshal()

	// Return
	return types.BytesToValidatorSignature(signature), srHash, nil

}
