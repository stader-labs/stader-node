/*
This work is licensed and released under GNU GPL v3 or any other later versions. 
The full text of the license is below/ found at <http://www.gnu.org/licenses/>

(c) 2023 Rocket Pool Pty Ltd. Modified under GNU GPL v3.

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
package stdr

// ROCKETPOOL-OWNED

import (
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/stader-labs/stader-node/shared/services/beacon"
	"github.com/stader-labs/stader-node/stader-lib/stader"
)

type FeeRecipientInfo struct {
	SocializingPoolAddress common.Address `json:"socializingPoolAddress"`
	FeeDistributorAddress  common.Address `json:"feeDistributorAddress"`
	IsInSocializingPool    bool           `json:"isInSocializingPool"`
	IsInOptOutCooldown     bool           `json:"isInOptOutCooldown"`
	OptOutEpoch            uint64         `json:"optOutEpoch"`
}

// TODO - add fee receipient info for socializing pools
func GetFeeRecipientInfo(prn *stader.PermissionlessNodeRegistryContractManager, bc beacon.Client, nodeAddress common.Address, opts *bind.CallOpts) (*FeeRecipientInfo, error) {
	// TODO - Get fee recipient info from sanjay
	return nil, nil
}
