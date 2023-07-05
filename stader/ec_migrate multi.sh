#!/bin/sh

# This work is licensed and released under GNU GPL v3 or any other later versions. 
# The full text of the license is below/ found at <http://www.gnu.org/licenses/>

# (c) 2023 Rocket Pool Pty Ltd. Modified under GNU GPL v3. [1.2.0]

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


EC_CHAINDATA_DIR=/ethclient
EXTERNAL_DIR=/mnt/external

# Get the core count
CORE_COUNT=$(grep -c ^processor /proc/cpuinfo)
if [ -z "$CORE_COUNT" ]; then
    CORE_COUNT=1
elif [ "$CORE_COUNT" -lt 1 ]; then
    CORE_COUNT=1
fi

RSYNC_CMD="xargs -n1 -P${CORE_COUNT} -I% rsync -a --progress %"

if [ "$EC_MIGRATE_MODE" = "export" ]; then
    ls $EC_CHAINDATA_DIR/* | $RSYNC_CMD $EXTERNAL_DIR
elif [ "$EC_MIGRATE_MODE" = "import" ]; then
    rm -rf $EC_CHAINDATA_DIR/*
    ls $EXTERNAL_DIR/* | $RSYNC_CMD $EC_CHAINDATA_DIR
else
    echo "Unknown migrate mode \"$EC_MIGRATE_MODE\""
fi