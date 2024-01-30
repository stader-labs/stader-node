#!/bin/bash

# This work is licensed and released under GNU GPL v3 or any other later versions. 
# The full text of the license is below/ found at <http://www.gnu.org/licenses/>

# (c) 2023 Rocket Pool Pty Ltd. Modified under GNU GPL v3. [1.4.7]

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

# Print usage
usage() {
    echo "Usage: build-ec-migrate.sh -v <version number>"
    echo "This script builds the EC Migrator image."
    exit 0
}

# =================
# === Main Body ===
# =================

# Get the version
while getopts "admv:" FLAG; do
    case "$FLAG" in
        v) VERSION="$OPTARG" ;;
        *) usage ;;
    esac
done
if [ -z "$VERSION" ]; then
    usage
fi

echo -n "Building Docker image... "
echo "Building Docker EC Migrator image..."
docker buildx build --platform=linux/amd64 -t staderdev/ec-migrator:$VERSION-amd64 -f docker/stader-ec-migrator --load . || fail "Error building amd64 Docker ec-migrator image."
docker buildx build --platform=linux/arm64 -t staderdev/ec-migrator:$VERSION-arm64 -f docker/stader-ec-migrator --load . || fail "Error building arm64 Docker ec-migrator image."
echo "done!"

echo -n "Pushing to Docker Hub... "
docker push staderdev/ec-migrator:$VERSION-amd64 || fail "Error pushing amd64 Docker EC Migrator image to Docker Hub."
docker push staderdev/ec-migrator:$VERSION-arm64 || fail "Error pushing arm Docker EC Migrator image to Docker Hub."
echo "done!"

 echo -n "Building Docker manifest... "
rm -f ~/.docker/manifests/docker.io_staderdev_ec-migrator-$VERSION
docker manifest create staderdev/ec-migrator:$VERSION --amend staderdev/ec-migrator:$VERSION-amd64 --amend staderdev/ec-migrator:$VERSION-arm64
echo "done!"

echo -n "Pushing to Docker Hub... "
docker manifest push --purge staderdev/ec-migrator:$VERSION
echo "done!"