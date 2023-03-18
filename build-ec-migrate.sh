#!/bin/bash

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