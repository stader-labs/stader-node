#!/bin/sh
# This work is licensed and released under GNU GPL v3 or any other later versions.
# The full text of the license is below/ found at <http://www.gnu.org/licenses/>

# (c) 2023 Rocket Pool Pty Ltd. Modified under GNU GPL v3. [0.4.0-beta]

# This program is free software: you can redistribute it and/or modify
# it under the terms of the GNU General Public License as published by
# the Free Software Foundation, either version 3 of the License, or
# (at your option) any later version.

# This program is distributed in the hope that it will be useful,
# but WITHOUT ANY WARRANTY; without even the implied warranty of
# MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
# GNU General Public License for more details.

# You should have received a copy of the GNU General Public License
# along with this program.  If not, see <http://www.gnu.org/licenses/>.


# Set up the network-based flag
if [ "$NETWORK" = "mainnet" ]; then
    MEV_NETWORK="mainnet"
elif [ "$NETWORK" = "prater" ]; then
    MEV_NETWORK="goerli"
elif [ "$NETWORK" = "devnet" ]; then
    MEV_NETWORK="goerli"
else
    echo "Unknown network [$NETWORK]"
    exit 1
fi

# Run MEV-boost
exec /app/mev-boost -${MEV_NETWORK} -addr 0.0.0.0:${MEV_BOOST_PORT} -relay-check -relays ${MEV_BOOST_RELAYS}