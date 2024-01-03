#!/bin/sh

# This work is licensed and released under GNU GPL v3 or any other later versions. 
# The full text of the license is below/ found at <http://www.gnu.org/licenses/>

# (c) 2023 Rocket Pool Pty Ltd. Modified under GNU GPL v3. [1.4.4]

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

##
# Stader service installation script
# Prints progress messages to stdout
# All command output is redirected to stderr
##

COLOR_RED='\033[0;31m'
COLOR_YELLOW='\033[33m'
COLOR_RESET='\033[0m'

# Print a failure message to stderr and exit
fail() {
    MESSAGE=$1
    >&2 echo -e "\n${COLOR_RED}**ERROR**\n$MESSAGE${COLOR_RESET}"
    exit 1
}


# Get CPU architecture
UNAME_VAL=$(uname -m)
ARCH=""
case $UNAME_VAL in
    x86_64)  ARCH="amd64" ;;
    aarch64) ARCH="arm64" ;;
    arm64)   ARCH="arm64" ;;
    *)       fail "CPU architecture not supported: $UNAME_VAL" ;;
esac


# Get the platform type
PLATFORM=$(uname -s)
if [ "$PLATFORM" = "Linux" ]; then
    if command -v lsb_release &>/dev/null ; then
        PLATFORM=$(lsb_release -si)
    elif [ -f "/etc/centos-release" ]; then
        PLATFORM="CentOS"
    elif [ -f "/etc/fedora-release" ]; then
        PLATFORM="Fedora"
    fi
fi

##
# Config
##


# The total number of steps in the installation process
TOTAL_STEPS="2"
# The Stader user data path
STADER_PATH="$HOME/.stader"
# The default stader node package version to download
PACKAGE_VERSION="latest"
# The default network to run Stader on
NETWORK="mainnet"

##
# Utils
##


# Print progress
progress() {
    STEP_NUMBER=$1
    MESSAGE=$2
    echo "Step $STEP_NUMBER of $TOTAL_STEPS: $MESSAGE"
}


# Docker installation steps
add_user_docker() {
    $SUDO_CMD usermod -aG docker $USER || fail "Could not add user to docker group."
}


# Detect installed privilege escalation programs
get_escalation_cmd() {
    if type sudo > /dev/null 2>&1; then
        SUDO_CMD="sudo"
    elif type doas > /dev/null 2>&1; then
        echo "NOTE: sudo not found, using doas instead"
        SUDO_CMD="doas"
    else
        fail "Please make sure a privilege escalation command such as \"sudo\" is installed and available before installing Stader."
    fi
}

# Install
install() {


##
# Initialization
##


# Parse arguments
while getopts "dp:u:n:v:" FLAG; do
    case "$FLAG" in
        d) NO_DEPS=true ;;
        p) STADER_PATH="$OPTARG" ;;
        u) DATA_PATH="$OPTARG" ;;
        n) NETWORK="$OPTARG" ;;
        v) PACKAGE_VERSION="$OPTARG" ;;
        *) fail "Incorrect usage." ;;
    esac
done

if [ -z "$DATA_PATH" ]; then
    DATA_PATH="$STADER_PATH/data"
fi

if [ "$PACKAGE_VERSION" = "latest" ]; then
    PACKAGE_URL="https://staderlabs.com/eth/releases/stader-node-build/permissionless/latest/stader-node-install.tar.xz"
else
    PACKAGE_URL="https://staderlabs.com/eth/releases/stader-node-build/permissionless/$PACKAGE_VERSION/stader-node-install.tar.xz"
fi

# Create temporary data folder; clean up on exit
TEMPDIR=$(mktemp -d 2>/dev/null) || fail "Could not create temporary data directory."
trap 'rm -rf "$TEMPDIR"' EXIT


# Get temporary data paths
PACKAGE_FILES_PATH="$TEMPDIR/install"


##
# Installation
##

# Check for existing installation
progress 1 "Checking for existing installation..."
if [ -d $STADER_PATH ]; then
    # Check for legacy files - key on the old config.yml
    if [ -f "$STADER_PATH/config.yml" ]; then
        progress 1 "Old installation detected, backing it up and migrating to new config system..."
        OLD_CONFIG_BACKUP_PATH="$STADER_PATH/old_config_backup"
        { mkdir -p $OLD_CONFIG_BACKUP_PATH || fail "Could not create old config backup folder."; } >&2

        if [ -f "$STADER_PATH/config.yml" ]; then
            { mv "$STADER_PATH/config.yml" "$OLD_CONFIG_BACKUP_PATH" || fail "Could not move config.yml to backup folder."; } >&2
        fi
        if [ -f "$STADER_PATH/settings.yml" ]; then
            { mv "$STADER_PATH/settings.yml" "$OLD_CONFIG_BACKUP_PATH" || fail "Could not move settings.yml to backup folder."; } >&2
        fi
        if [ -f "$STADER_PATH/docker-compose.yml" ]; then
            { mv "$STADER_PATH/docker-compose.yml" "$OLD_CONFIG_BACKUP_PATH" || fail "Could not move docker-compose.yml to backup folder."; } >&2
        fi
        if [ -f "$STADER_PATH/docker-compose-metrics.yml" ]; then
            { mv "$STADER_PATH/docker-compose-metrics.yml" "$OLD_CONFIG_BACKUP_PATH" || fail "Could not move docker-compose-metrics.yml to backup folder."; } >&2
        fi
        if [ -f "$STADER_PATH/docker-compose-fallback.yml" ]; then
            { mv "$STADER_PATH/docker-compose-fallback.yml" "$OLD_CONFIG_BACKUP_PATH" || fail "Could not move docker-compose-fallback.yml to backup folder."; }>&2
        fi
        if [ -f "$STADER_PATH/prometheus.tmpl" ]; then
            { mv "$STADER_PATH/prometheus.tmpl" "$OLD_CONFIG_BACKUP_PATH" || fail "Could not move prometheus.tmpl to backup folder."; } >&2
        fi
        if [ -f "$STADER_PATH/grafana-prometheus-datasource.yml" ]; then
            { mv "$STADER_PATH/grafana-prometheus-datasource.yml" "$OLD_CONFIG_BACKUP_PATH" || fail "Could not move grafana-prometheus-datasource.yml to backupfolder."; } >&2
        fi
        if [ -d "$STADER_PATH/chains" ]; then
            { mv "$STADER_PATH/chains" "$OLD_CONFIG_BACKUP_PATH" || fail "Could not move chains directory to backup folder."; } >&2
        fi
    fi

    # Back up existing config file
    if [ -f "$STADER_PATH/user-settings.yml" ]; then
        progress 1 "Backing up configuration settings to user-settings-backup.yml..."
        { cp "$STADER_PATH/user-settings.yml" "$STADER_PATH/user-settings-backup.yml" || fail "Could not backup configuration settings."; } >&2
    fi
fi


# Download and extract package files
progress 2 "Downloading Stader package files..."
{ curl -L "$PACKAGE_URL" | tar -xJ -C "$TEMPDIR" || fail "Could not download and extract the Stader package files."; } >&2
{ test -d "$PACKAGE_FILES_PATH" || fail "Could not extract the Stader package files."; } >&2


# Copy package files
progress 2 "Copying package files to Stader user data directory..."
{ cp -r -n "$PACKAGE_FILES_PATH/override" "$STADER_PATH" || rsync -r --ignore-existing "$PACKAGE_FILES_PATH/override" "$STADER_PATH" || fail "Could not copy new override files to the Stader user data directory."; } >&2
{ cp -r "$PACKAGE_FILES_PATH/scripts" "$STADER_PATH" || fail "Could not copy scripts folder to the Stader user data directory."; } >&2
{ cp -r "$PACKAGE_FILES_PATH/templates" "$STADER_PATH" || fail "Could not copy templates folder to the Stader user data directory."; } >&2
{ cp "$PACKAGE_FILES_PATH/grafana-prometheus-datasource.yml" "$PACKAGE_FILES_PATH/prometheus.tmpl" "$STADER_PATH" || fail "Could not copy base files to the Stader user data directory."; } >&2
{ find "$STADER_PATH/scripts" -name "*.sh" -exec chmod +x {} \; 2>/dev/null || fail "Could not set executable permissions on package files."; } >&2

# Clean up unnecessary files from old installations
progress 2 "Cleaning up obsolete files from previous installs..."

}

install "$@"
