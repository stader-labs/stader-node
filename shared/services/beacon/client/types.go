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
package client

import (
	"encoding/hex"
	"encoding/json"
	"strconv"

	"github.com/ethereum/go-ethereum/common"
	hexutil "github.com/stader-labs/stader-node/shared/utils/hex"
)

// Request types
type VoluntaryExitRequest struct {
	Message   VoluntaryExitMessage `json:"message"`
	Signature byteArray            `json:"signature"`
}
type VoluntaryExitMessage struct {
	Epoch          uinteger `json:"epoch"`
	ValidatorIndex uinteger `json:"validator_index"`
}

// Response types
type SyncStatusResponse struct {
	Data struct {
		IsSyncing    bool     `json:"is_syncing"`
		HeadSlot     uinteger `json:"head_slot"`
		SyncDistance uinteger `json:"sync_distance"`
	} `json:"data"`
}
type Eth2ConfigResponse struct {
	Data struct {
		SecondsPerSlot               uinteger `json:"SECONDS_PER_SLOT"`
		SlotsPerEpoch                uinteger `json:"SLOTS_PER_EPOCH"`
		EpochsPerSyncCommitteePeriod uinteger `json:"EPOCHS_PER_SYNC_COMMITTEE_PERIOD"`
	} `json:"data"`
}
type Eth2DepositContractResponse struct {
	Data struct {
		ChainID uinteger       `json:"chain_id"`
		Address common.Address `json:"address"`
	} `json:"data"`
}
type GenesisResponse struct {
	Data struct {
		GenesisTime           uinteger  `json:"genesis_time"`
		GenesisForkVersion    byteArray `json:"genesis_fork_version"`
		GenesisValidatorsRoot byteArray `json:"genesis_validators_root"`
	} `json:"data"`
}
type FinalityCheckpointsResponse struct {
	Data struct {
		PreviousJustified struct {
			Epoch uinteger `json:"epoch"`
		} `json:"previous_justified"`
		CurrentJustified struct {
			Epoch uinteger `json:"epoch"`
		} `json:"current_justified"`
		Finalized struct {
			Epoch uinteger `json:"epoch"`
		} `json:"finalized"`
	} `json:"data"`
}
type ForkResponse struct {
	Data struct {
		PreviousVersion byteArray `json:"previous_version"`
		CurrentVersion  byteArray `json:"current_version"`
		Epoch           uinteger  `json:"epoch"`
	} `json:"data"`
}
type AttestationsResponse struct {
	Data []Attestation `json:"data"`
}
type BeaconBlockResponse struct {
	Data struct {
		Message struct {
			Slot          uinteger `json:"slot"`
			ProposerIndex uinteger `json:"proposer_index"`
			Body          struct {
				Eth1Data struct {
					DepositRoot  byteArray `json:"deposit_root"`
					DepositCount uinteger  `json:"deposit_count"`
					BlockHash    byteArray `json:"block_hash"`
				} `json:"eth1_data"`
				Attestations     []Attestation `json:"attestations"`
				ExecutionPayload *struct {
					FeeRecipient byteArray `json:"fee_recipient"`
					BlockNumber  uinteger  `json:"block_number"`
				} `json:"execution_payload"`
			} `json:"body"`
		} `json:"message"`
	} `json:"data"`
}
type ValidatorsResponse struct {
	Data []Validator `json:"data"`
}
type Validator struct {
	Index     uinteger `json:"index"`
	Balance   uinteger `json:"balance"`
	Status    string   `json:"status"`
	Validator struct {
		Pubkey                     byteArray `json:"pubkey"`
		WithdrawalCredentials      byteArray `json:"withdrawal_credentials"`
		EffectiveBalance           uinteger  `json:"effective_balance"`
		Slashed                    bool      `json:"slashed"`
		ActivationEligibilityEpoch uinteger  `json:"activation_eligibility_epoch"`
		ActivationEpoch            uinteger  `json:"activation_epoch"`
		ExitEpoch                  uinteger  `json:"exit_epoch"`
		WithdrawableEpoch          uinteger  `json:"withdrawable_epoch"`
	} `json:"validator"`
}
type SyncDutiesResponse struct {
	Data []SyncDuty `json:"data"`
}
type SyncDuty struct {
	Pubkey               byteArray  `json:"pubkey"`
	ValidatorIndex       uinteger   `json:"validator_index"`
	SyncCommitteeIndices []uinteger `json:"validator_sync_committee_indices"`
}
type ProposerDutiesResponse struct {
	Data []ProposerDuty `json:"data"`
}
type ProposerDuty struct {
	ValidatorIndex uinteger `json:"validator_index"`
}

type CommitteesResponse struct {
	Data []Committee `json:"data"`
}

type Committee struct {
	Index      uinteger   `json:"index"`
	Slot       uinteger   `json:"slot"`
	Validators []uinteger `json:"validators"`
}

type Attestation struct {
	AggregationBits string `json:"aggregation_bits"`
	Data            struct {
		Slot  uinteger `json:"slot"`
		Index uinteger `json:"index"`
	} `json:"data"`
}

// Unsigned integer type
type uinteger uint64

func (i uinteger) MarshalJSON() ([]byte, error) {
	return json.Marshal(strconv.Itoa(int(i)))
}
func (i *uinteger) UnmarshalJSON(data []byte) error {

	// Unmarshal string
	var dataStr string
	if err := json.Unmarshal(data, &dataStr); err != nil {
		return err
	}

	// Parse integer value
	value, err := strconv.ParseUint(dataStr, 10, 64)
	if err != nil {
		return err
	}

	// Set value and return
	*i = uinteger(value)
	return nil

}

// Byte array type
type byteArray []byte

func (b byteArray) MarshalJSON() ([]byte, error) {
	return json.Marshal(hexutil.AddPrefix(hex.EncodeToString(b)))
}
func (b *byteArray) UnmarshalJSON(data []byte) error {

	// Unmarshal string
	var dataStr string
	if err := json.Unmarshal(data, &dataStr); err != nil {
		return err
	}

	// Decode hex
	value, err := hex.DecodeString(hexutil.RemovePrefix(dataStr))
	if err != nil {
		return err
	}

	// Set value and return
	*b = value
	return nil

}
