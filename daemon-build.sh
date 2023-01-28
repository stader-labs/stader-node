#!/bin/bash

# Get the platform type and run the build script if possible
PLATFORM=$(uname -s)
if [ "$PLATFORM" = "Linux" ]; then
    docker run --rm -v $PWD:/stader-node rocketpool/smartnode-builder:latest /stader-node/stader/build.sh
else
    echo "Platform ${PLATFORM} is not supported by this script, please build the daemon manually."
    exit 1
fi
