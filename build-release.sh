#!/bin/bash

# This script will build all of the artifacts involved in a new Stader Stader node release
# (except for the macOS daemons, which need to be built manually on a macOS system) and put
# them into a convenient folder for ease of uploading.

# NOTE: You MUST put this in a directory that has the `stader-node` and `stader-node-install`
# repositories cloned as subdirectories.


# =================
# === Functions ===
# =================

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
    cd stader-node || fail "Directory ${PWD}/stader-node/stader-cli does not exist or you don't have permissions to access it."

    echo -n "Building CLI binaries... "
    docker run --rm -v $PWD:/stader-node staderdev/stader-node-builder:latest /stader-node/stader-cli/build.sh || fail "Error building CLI binaries."
    mv stader-cli/stader-cli-* ../$VERSION
    cd ..
    # aws s3 cp $VERSION s3://stadernode/$VERSION --recursive
    echo "done!"


}


# Builds the .tar.xz file packages with the Stader configuration files
build_install_packages() {
    cd stader-node || fail "Directory ${PWD}/stader-node-install does not exist or you don't have permissions to access it."
    rm -f stader-node-install.tar.xz

    echo -n "Building Stader node installer packages... "
    tar cfJ stader-node-install.tar.xz install || fail "Error building installer package."
    mv stader-node-install.tar.xz ../$VERSION
    cp install.sh ../$VERSION
    cp install-update-tracker.sh ../$VERSION
    cd ..
    echo "done!"

    cd stader-node-install || fail "Directory ${PWD}/stader-node-install does not exist or you don't have permissions to access it."
    echo -n "Building update tracker package... "
    tar cfJ stader-update-tracker.tar.xz stader-update-tracker || fail "Error building update tracker package."
    mv stader-update-tracker.tar.xz ../$VERSION
    cd ..
    aws s3 cp $VERSION s3://temps3node/$VERSION --recursive
    echo "done!"

    #cd ..
}


# Builds the daemon binaries and Docker images, and pushes them to Docker Hub
build_daemon() {
    cd stader-node || fail "Directory ${PWD}/stader-node does not exist or you don't have permissions to access it."

    echo -n "Building Daemon binary... "
    ./daemon-build.sh || fail "Error building daemon binary."
    cp stader/stader-daemon-* ../$VERSION
    echo "done!"
        # ensure support for arm64 is installed by  sudo apt install -y qemu-user-static binfmt-support
    echo "Building Docker Stader Daemon image..."
    docker buildx build --platform=linux/amd64 -t staderdev/stader-node:$VERSION-amd64 -f docker/stader-dockerfile --load . || fail "Error building amd64 Docker Stader Daemon image."
    docker buildx build --platform=linux/arm64 -t staderdev/stader-node:$VERSION-arm64 -f docker/stader-dockerfile --load . || fail "Error building arm64 Docker Stader Daemon image."
    echo "done!"

    echo -n "Pushing to Docker Hub... "
    docker push staderdev/stader-node:$VERSION-amd64 || fail "Error pushing amd64 Docker Stader Daemon image to Docker Hub."
    docker push staderdev/stader-node:$VERSION-arm64 || fail "Error pushing arm Docker Stader Daemon image to Docker Hub."
    rm -f stader/stader-daemon-*
    echo "done!"

    cd ..
}


# Builds the Docker prune provisioner image and pushes it to Docker Hub
build_docker_prune_provision() {
    cd stader-node || fail "Directory ${PWD}/stader-node does not exist or you don't have permissions to access it."

    echo "Building Docker Prune Provisioner image..."
    docker buildx build --platform=linux/amd64 -t staderdev/stader-node:$VERSION-amd64 -f docker/stader-prune-provision --load . || fail "Error building amd64 Docker Prune Provision  image."
    docker buildx build --platform=linux/arm64 -t staderdev/stader-node:$VERSION-arm64 -f docker/stader-prune-provision --load . || fail "Error building arm64 Docker Prune Provision  image."
    echo "done!"

    echo -n "Pushing to Docker Hub... "
    docker push staderdev/eth1-prune-provision:$VERSION-amd64 || fail "Error pushing amd64 Docker Prune Provision image to Docker Hub."
    docker push staderdev/eth1-prune-provision:$VERSION-arm64 || fail "Error pushing arm Docker Prune Provision image to Docker Hub."
    echo "done!"

    cd ..
}


# Builds the Docker Manifests and pushes them to Docker Hub
build_docker_manifest() {
    echo -n "Building Docker manifest... "
    rm -f ~/.docker/manifests/docker.io_staderdev_stader-node-$VERSION
    docker manifest create staderdev/stader-node:$VERSION --amend staderdev/stader-node:$VERSION-amd64 --amend staderdev/stader-node:$VERSION-arm64
    echo "done!"

    echo -n "Pushing to Docker Hub... "
    docker manifest push --purge staderdev/stader-node:$VERSION
    echo "done!"
}


# Builds the 'latest' Docker Manifests and pushes them to Docker Hub
build_latest_docker_manifest() {
    echo -n "Building 'latest' Docker manifest... "
    rm -f ~/.docker/manifests/docker.io_staderdev_stader-node-latest
    docker manifest create staderdev/stader-node:latest --amend staderdev/stader-node:$VERSION-amd64 --amend staderdev/stader-node:$VERSION-arm64
    echo "done!"

    echo -n "Pushing to Docker Hub... "
    docker manifest push --purge staderdev/stader-node:latest
    echo "done!"
}


# Builds the Docker Manifest for the prune provisioner and pushes it to Docker Hub
build_docker_prune_provision_manifest() {
    echo -n "Building Docker Prune Provision manifest... "
    rm -f ~/.docker/manifests/docker.io_staderdev_eth1-prune-provision-$VERSION
    docker manifest create staderdev/eth1-prune-provision:$VERSION --amend staderdev/eth1-prune-provision:$VERSION-amd64 --amend staderdev/eth1-prune-provision:$VERSION-arm64
    echo "done!"

    echo -n "Pushing to Docker Hub... "
    docker manifest push --purge staderdev/eth1-prune-provision:$VERSION
    echo "done!"
}

send_slack() {
    # curl -X POST --data-urlencode 'payload={"channel": "'${SLACK_CHAN}'", "username": "'${SLACK_USERNAME}'", "text": "'"${*}"'", "icon_emoji": "'${SLACK_ICON}'"}' ${URL} > /dev/null 2>&1
    curl -X POST --data-urlencode "payload={\"channel\": \"#stader-node-build-alerts\", \"username\": \"webhookbot\", \"text\": \"${*}\", \"icon_emoji\": \":ghost:\"}" https://hooks.slack.com/services/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 > /dev/null 2>&1
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
rm -f ./$VERSION/*
mkdir -p ./$VERSION

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
# if all successful, send slack message
if [ $? -eq 0 ]; then
    send_slack "Stader Node $VERSION has been built and pushed to Docker Hub."
# else send failure message
else
    send_slack "Stader Node $VERSION has failed to build and push to Docker Hub."
fi