#!/bin/bash

# This work is licensed and released under GNU GPL v3 or any other later versions. 
# The full text of the license is below/ found at <http://www.gnu.org/licenses/>

# (c) 2023 Rocket Pool Pty Ltd. Modified under GNU GPL v3. [1.3.0]

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

# This script will build all of the artifacts involved in a new Stader Stader node release
# (except for the macOS daemons, which need to be built manually on a macOS system) and put
# them into a convenient folder for ease of uploading.


# =================
# === Functions ===
# =================

DOCKER_ACCOUNT=staderdev
S3_BUCKET=stader-cli-permissionless/eth/releases/stader-node-build/permissionless

# Print a failure message to stderr and exit
fail() {
    MESSAGE=$1
    RED='\033[0;31m'
    RESET='\033[;0m'
    >&2 echo -e "\n${RED}**ERROR**\n$MESSAGE${RESET}\n"
    exit 1
}


# Builds all of the CLI binaries
build_cli() {
    echo -n "Building CLI binaries... "
    docker run --rm -v $PWD:/stader-node staderdev/stader-node-builder:latest /stader-node/stader-cli/build.sh || fail "Error building CLI binaries."
    mv stader-cli/stader-cli-* build/$VERSION
    echo "done!"
}


# Builds the .tar.xz file packages with the Stader configuration files
build_install_packages() {
    rm -f stader-node-install.tar.xz
    echo -n "Building Stader node installer packages... "
    tar cfJ stader-node-install.tar.xz install || fail "Error building installer package."
    mv stader-node-install.tar.xz build/$VERSION
    cp install.sh build/$VERSION
    aws s3 cp build/$VERSION s3://$S3_BUCKET/$VERSION --recursive
    echo "done!"

}


# Builds the daemon binaries and Docker images, and pushes them to Docker Hub
build_daemon() {
    echo -n "Building Daemon binary... "
    ./daemon-build.sh || fail "Error building daemon binary."
    cp stader/stader-daemon-* build/$VERSION
    echo "done!"
        # ensure support for arm64 is installed by  sudo apt install -y qemu-user-static binfmt-support
    echo "Building Docker Stader Daemon image..."
    docker buildx build --platform=linux/amd64 -t $DOCKER_ACCOUNT/stader-permissionless:$VERSION-amd64 -f docker/stader-dockerfile --load . || fail "Error building amd64 Docker Stader Daemon image."
    docker buildx build --platform=linux/arm64 -t $DOCKER_ACCOUNT/stader-permissionless:$VERSION-arm64 -f docker/stader-dockerfile --load . || fail "Error building arm64 Docker Stader Daemon image."
    echo "done!"

    echo -n "Pushing to Docker Hub... "
    docker push $DOCKER_ACCOUNT/stader-permissionless:$VERSION-amd64 || fail "Error pushing amd64 Docker Stader Daemon image to Docker Hub."
    docker push $DOCKER_ACCOUNT/stader-permissionless:$VERSION-arm64 || fail "Error pushing arm Docker Stader Daemon image to Docker Hub."
    rm -f stader/stader-daemon-*
    echo "done!"
}


# Builds the Docker prune provisioner image and pushes it to Docker Hub
build_docker_prune_provision() {
    echo "Building Docker Prune Provisioner image..."
    docker buildx build --platform=linux/amd64 -t $DOCKER_ACCOUNT/stader-permissionless:$VERSION-amd64 -f docker/stader-prune-provision --load . || fail "Error building amd64 Docker Prune Provision  image."
    docker buildx build --platform=linux/arm64 -t $DOCKER_ACCOUNT/stader-permissionless:$VERSION-arm64 -f docker/stader-prune-provision --load . || fail "Error building arm64 Docker Prune Provision  image."
    echo "done!"

    echo -n "Pushing to Docker Hub... "
    docker push $DOCKER_ACCOUNT/eth1-prune-provision:$VERSION-amd64 || fail "Error pushing amd64 Docker Prune Provision image to Docker Hub."
    docker push $DOCKER_ACCOUNT/eth1-prune-provision:$VERSION-arm64 || fail "Error pushing arm Docker Prune Provision image to Docker Hub."
    echo "done!"
}


# Builds the Docker Manifests and pushes them to Docker Hub
build_docker_manifest() {
    echo -n "Building Docker manifest... "
    rm -f ~/.docker/manifests/docker.io_$DOCKER_ACCOUNT_stader-permissionless-$VERSION
    docker manifest create $DOCKER_ACCOUNT/stader-permissionless:$VERSION --amend $DOCKER_ACCOUNT/stader-permissionless:$VERSION-amd64 --amend $DOCKER_ACCOUNT/stader-permissionless:$VERSION-arm64
    echo "done!"

    echo -n "Pushing to Docker Hub... "
    docker manifest push --purge $DOCKER_ACCOUNT/stader-permissionless:$VERSION
    echo "done!"
}


# Builds the 'latest' Docker Manifests and pushes them to Docker Hub
build_latest_docker_manifest() {
    echo -n "Building 'latest' Docker manifest... "
    rm -f ~/.docker/manifests/docker.io_$DOCKER_ACCOUNT_stader-permissionless-latest
    docker manifest create $DOCKER_ACCOUNT/stader-permissionless:latest --amend $DOCKER_ACCOUNT/stader-permissionless:$VERSION-amd64 --amend $DOCKER_ACCOUNT/stader-permissionless:$VERSION-arm64
    echo "done!"

    echo -n "Pushing to Docker Hub... "
    docker manifest push --purge $DOCKER_ACCOUNT/stader-permissionless:latest
    echo "done!"
}


# Builds the Docker Manifest for the prune provisioner and pushes it to Docker Hub
build_docker_prune_provision_manifest() {
    echo -n "Building Docker Prune Provision manifest... "
    rm -f ~/.docker/manifests/docker.io_$DOCKER_ACCOUNT_eth1-prune-provision-$VERSION
    docker manifest create $DOCKER_ACCOUNT/eth1-prune-provision:$VERSION --amend $DOCKER_ACCOUNT/eth1-prune-provision:$VERSION-amd64 --amend $DOCKER_ACCOUNT/eth1-prune-provision:$VERSION-arm64
    echo "done!"

    echo -n "Pushing to Docker Hub... "
    docker manifest push --purge $DOCKER_ACCOUNT/eth1-prune-provision:$VERSION
    echo "done!"
}

# Print usage
usage() {
    echo "Usage: build-release.sh [options] -v <version number>"
    echo "This script assumes it is in a directory that contains subdirectories for all of the Stader repositories."
    echo "Options:"
    echo $'\t-a\tBuild all of the artifacts, except for the prune provisioner'
    echo $'\t-c\tBuild the CLI binaries for all platforms'
    echo $'\t-p\tBuild the Stader installer packages'
    echo $'\t-d\tBuild the Daemon binaries and Docker images, and push them to Docker Hub'
    echo $'\t-x\tBuild the Docker POW Proxy image and push it to Docker Hub'
    echo $'\t-n\tBuild the Docker manifests (Stader node and POW Proxy), and push them to Docker Hub'
    echo $'\t-r\tBuild the Docker Prune Provisioner image and push it to Docker Hub'
    echo $'\t-f\tBuild the Docker manifest for the Prune Provisioner and push it to Docker Hub'
    exit 0
}


# =================
# === Main Body ===
# =================

# Get CPU architecture
UNAME_VAL=$(uname -m)
ARCH=""
case $UNAME_VAL in
    x86_64)  ARCH="amd64" ;;
    aarch64) ARCH="arm64" ;;
    arm64)   ARCH="arm64" ;;
    *)       fail "CPU architecture not supported: $UNAME_VAL" ;;
esac

# Parse arguments
while getopts "acpdnlrfv:" FLAG; do
    case "$FLAG" in
        a) CLI=true PACKAGES=true DAEMON=true MANIFEST=true LATEST_MANIFEST=true ;;
        c) CLI=true ;;
        p) PACKAGES=true ;;
        d) DAEMON=true ;;
        n) MANIFEST=true ;;
        l) LATEST_MANIFEST=true ;;
        r) PRUNE=true ;;
        f) PRUNE_MANIFEST=true ;;
        v) VERSION="$OPTARG" ;;
        *) usage ;;
    esac
done
if [ -z "$VERSION" ]; then
    usage
fi

# Cleanup old artifacts
rm -f build/$VERSION/*
mkdir -p build/$VERSION

# Build the artifacts
if [ "$CLI" = true ]; then
    build_cli
fi
if [ "$PACKAGES" = true ]; then
    build_install_packages
fi
if [ "$DAEMON" = true ]; then
    build_daemon
fi
if [ "$MANIFEST" = true ]; then
    build_docker_manifest
fi
if [ "$LATEST_MANIFEST" = true ]; then
    build_latest_docker_manifest
fi
if [ "$PRUNE" = true ]; then
    build_docker_prune_provision
fi
if [ "$PRUNE_MANIFEST" = true ]; then
    build_docker_prune_provision_manifest
fi