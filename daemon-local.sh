#!/bin/bash

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
