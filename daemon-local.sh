#!/bin/bash

# This work is licensed and released under GNU GPL v3 or any other later versions. 
# The full text of the license is below/ found at <http://www.gnu.org/licenses/>

# (c) 2023 Rocket Pool Pty Ltd. Modified under GNU GPL v3. [1.4.1]

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

# Get the platform type and run the build script if possible

# Print usage
usage() {
    echo "Usage: daemon-local.sh -v <version number>"
    exit 0
}

DOCKER_ACCOUNT=staderlabs
# Parse arguments
while getopts "acpdnlrfv:" FLAG; do
    case "$FLAG" in
        v) VERSION="$OPTARG" ;;
        *) usage ;;
    esac
done
if [ -z "$VERSION" ]; then
    usage
fi

# Get CPU architecture
UNAME_VAL=$(uname -m)
ARCH=""
case $UNAME_VAL in
    x86_64)  ARCH="amd64" ;;
    aarch64) ARCH="arm64" ;;
    arm64)   ARCH="arm64" ;;
    *)       fail "CPU architecture not supported: $UNAME_VAL" ;;
esac


echo "Start buiding: " $VERSION
 
docker run --rm -v $PWD:/stader-node staderdev/stader-node-builder:latest /stader-node/stader/build.sh

cp stader/stader-daemon-* build/$VERSION

echo "done!"

docker buildx build --platform=linux/$ARCH -t $DOCKER_ACCOUNT/stader-permissionless:$VERSION -f docker/stader-dockerfile --load . || fail "Error building $ARCH Docker Stader Daemon image."

echo "done!"
