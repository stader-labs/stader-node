#!/bin/bash

# Print usage
usage() {
    echo "Usage: build-release.sh -v <version number>"
    echo "This script builds the Smartnode builder image used to build the daemon binaries."
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
docker build -t staderdev/stader-node:$VERSION -f docker/smartnode-builder .
docker tag staderdev/stader-node:$VERSION staderdev/stader-node:latest
docker push staderdev/stader-node:$VERSION
docker push staderdev/stader-node:latestx  x64
echo "done!"