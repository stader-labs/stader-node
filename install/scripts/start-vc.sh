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

# This script launches ETH2 validator clients for Stader docker stack; only edit if you know what you're doing ;)

GWW_GRAFFITI_FILE="/addons/gww/graffiti.txt"

# Set up the network-based flags
if [ "$NETWORK" = "mainnet" ]; then
    LH_NETWORK="mainnet"
    LODESTAR_NETWORK="mainnet"
    PRYSM_NETWORK="--mainnet"
    TEKU_NETWORK="mainnet"
elif [ "$NETWORK" = "prater" ]; then
    LH_NETWORK="prater"
    LODESTAR_NETWORK="goerli"
    PRYSM_NETWORK="--prater"
    TEKU_NETWORK="prater"
elif [ "$NETWORK" = "devnet" ]; then
    LH_NETWORK="prater"
    LODESTAR_NETWORK="goerli"
    PRYSM_NETWORK="--prater"
    TEKU_NETWORK="prater"
elif [ "$NETWORK" = "zhejiang" ]; then
    LH_NETWORK=""
    LODESTAR_NETWORK=""
    PRYSM_NETWORK="--chain-config-file=/zhejiang/config.yaml"
    TEKU_NETWORK="/zhejiang/config.yaml"
else
    echo "Unknown network [$NETWORK]"
    exit 1
fi

# Report a missing fee recipient file
if [ ! -f "/validators/$FEE_RECIPIENT_FILE" ]; then
    echo "Fee recipient file not found, please wait for the Stader node process to create one."
    exit 1
fi


# Lighthouse startup
if [ "$CC_CLIENT" = "lighthouse" ]; then

    if [ "$NETWORK" = "zhejiang" ]; then
        LH_NETWORK_ARG="--testnet-dir=/zhejiang"
    else
        LH_NETWORK_ARG="--network $LH_NETWORK" 
    fi

    # Set up the CC + fallback string
    CC_URL_STRING=$CC_API_ENDPOINT
    if [ ! -z "$FALLBACK_CC_API_ENDPOINT" ]; then
        CC_URL_STRING="$CC_API_ENDPOINT,$FALLBACK_CC_API_ENDPOINT"
    fi

    CMD="/usr/local/bin/lighthouse validator \
        $LH_NETWORK_ARG \
        --datadir /validators/lighthouse \
        --init-slashing-protection \
        --logfile-max-number 0 \
        --beacon-nodes $CC_URL_STRING \
        --suggested-fee-recipient $(cat /validators/$FEE_RECIPIENT_FILE) \
        $VC_ADDITIONAL_FLAGS"

    if [ "$DOPPELGANGER_DETECTION" = "true" ]; then
        CMD="$CMD --enable-doppelganger-protection"
    fi

    if [ "$ENABLE_MEV_BOOST" = "true" ]; then
        CMD="$CMD --builder-proposals"
    fi

    if [ "$ENABLE_METRICS" = "true" ]; then
        CMD="$CMD --metrics --metrics-address 0.0.0.0 --metrics-port $VC_METRICS_PORT"
    fi

    if [ "$ENABLE_BITFLY_NODE_METRICS" = "true" ]; then
        CMD="$CMD --monitoring-endpoint $BITFLY_NODE_METRICS_ENDPOINT?apikey=$BITFLY_NODE_METRICS_SECRET&machine=$BITFLY_NODE_METRICS_MACHINE_NAME"
    fi

    exec ${CMD} --graffiti "$GRAFFITI"

fi

# Lodestar startup
if [ "$CC_CLIENT" = "lodestar" ]; then

    if [ "$NETWORK" = "zhejiang" ]; then
        LODESTAR_NETWORK_ARG="--paramsFile=/zhejiang/config.yaml"
    else
        LODESTAR_NETWORK_ARG="--network $LODESTAR_NETWORK" 
    fi

    # Remove any lock files that were left over accidentally after an unclean shutdown
    find /validators/lodestar/validators -name voting-keystore.json.lock -delete

    # Set up the CC + fallback string
    CC_URL_STRING=$CC_API_ENDPOINT
    if [ ! -z "$FALLBACK_CC_API_ENDPOINT" ]; then
        CC_URL_STRING="$CC_API_ENDPOINT,$FALLBACK_CC_API_ENDPOINT"
    fi

    CMD="/usr/app/node_modules/.bin/lodestar validator \
        $LODESTAR_NETWORK_ARG \
        --dataDir /validators/lodestar \
        --beacon-nodes $CC_URL_STRING \
        $FALLBACK_CC_STRING \
        --keystoresDir /validators/lodestar/validators \
        --secretsDir /validators/lodestar/secrets \
        --suggestedFeeRecipient $(cat /validators/$FEE_RECIPIENT_FILE) \
        $VC_ADDITIONAL_FLAGS"

    if [ "$DOPPELGANGER_DETECTION" = "true" ]; then
        CMD="$CMD --doppelgangerProtectionEnabled"
    fi

    if [ ! -z "$MEV_BOOST_URL" ]; then
        CMD="$CMD --builder"
    fi

    if [ "$ENABLE_METRICS" = "true" ]; then
        CMD="$CMD --metrics --metrics.address 0.0.0.0 --metrics.port $VC_METRICS_PORT"
    fi

    exec ${CMD} --graffiti "$GRAFFITI"

fi


# Nimbus startup
if [ "$CC_CLIENT" = "nimbus" ]; then

    # Nimbus won't start unless the validator directories already exist
    mkdir -p /validators/nimbus/validators
    mkdir -p /validators/nimbus/secrets

    # Set up the fallback arg
    if [ ! -z "$FALLBACK_CC_API_ENDPOINT" ]; then
        FALLBACK_CC_ARG="--beacon-node=$FALLBACK_CC_API_ENDPOINT"
    fi

    CMD="/home/user/nimbus_validator_client \
        --non-interactive \
        --beacon-node=$CC_API_ENDPOINT $FALLBACK_CC_ARG \
        --data-dir=/ethclient/nimbus_vc \
        --validators-dir=/validators/nimbus/validators \
        --secrets-dir=/validators/nimbus/secrets \
        --doppelganger-detection=$DOPPELGANGER_DETECTION \
        --suggested-fee-recipient=$(cat /validators/$FEE_RECIPIENT_FILE) \
        $VC_ADDITIONAL_FLAGS"

    if [ "$ENABLE_MEV_BOOST" = "true" ]; then
        CMD="$CMD --payload-builder"
    fi

    if [ "$ENABLE_METRICS" = "true" ]; then
        CMD="$CMD --metrics --metrics-address=0.0.0.0 --metrics-port=$VC_METRICS_PORT"
    fi

    # Graffiti breaks if it's in the CMD string instead of here because of spaces
    exec ${CMD} --graffiti="$GRAFFITI"

fi


# Prysm startup
if [ "$CC_CLIENT" = "prysm" ]; then

    # Make the Prysm dir
    mkdir -p /validators/prysm-non-hd/

    # Get rid of the protocol prefix
    CC_RPC_ENDPOINT=$(echo $CC_RPC_ENDPOINT | sed -E 's/.*\:\/\/(.*)/\1/')
    if [ ! -z "$FALLBACK_CC_RPC_ENDPOINT" ]; then
        FALLBACK_CC_RPC_ENDPOINT=$(echo $FALLBACK_CC_RPC_ENDPOINT | sed -E 's/.*\:\/\/(.*)/\1/')
    fi

    # Set up the CC + fallback string
    CC_URL_STRING=$CC_RPC_ENDPOINT
    if [ ! -z "$FALLBACK_CC_RPC_ENDPOINT" ]; then
        CC_URL_STRING="$CC_RPC_ENDPOINT,$FALLBACK_CC_RPC_ENDPOINT"
    fi

    CMD="/app/cmd/validator/validator \
        --accept-terms-of-use \
        $PRYSM_NETWORK \
        --wallet-dir /validators/prysm-non-hd \
        --wallet-password-file /validators/prysm-non-hd/direct/accounts/secret \
        --beacon-rpc-provider $CC_URL_STRING \
        --suggested-fee-recipient $(cat /validators/$FEE_RECIPIENT_FILE) \
        $VC_ADDITIONAL_FLAGS"

    if [ "$ENABLE_MEV_BOOST" = "true" ]; then
        CMD="$CMD --enable-builder"
    fi

    if [ "$DOPPELGANGER_DETECTION" = "true" ]; then
        CMD="$CMD --enable-doppelganger"
    fi

    if [ "$ENABLE_METRICS" = "true" ]; then
        CMD="$CMD --monitoring-host 0.0.0.0 --monitoring-port $VC_METRICS_PORT"
    else
        CMD="$CMD --disable-account-metrics"
    fi


    exec ${CMD} --graffiti "$GRAFFITI"


fi


# Teku startup
if [ "$CC_CLIENT" = "teku" ]; then

    # Teku won't start unless the validator directories already exist
    mkdir -p /validators/teku/keys
    mkdir -p /validators/teku/passwords

    # Remove any lock files that were left over accidentally after an unclean shutdown
    rm -f /validators/teku/keys/*.lock

    # Set up the CC + fallback string
    CC_URL_STRING=$CC_API_ENDPOINT
    if [ ! -z "$FALLBACK_CC_API_ENDPOINT" ]; then
        CC_URL_STRING="$CC_API_ENDPOINT,$FALLBACK_CC_API_ENDPOINT"
    fi

    CMD="/opt/teku/bin/teku validator-client \
        --network=$TEKU_NETWORK \
        --data-path=/validators/teku \
        --validator-keys=/validators/teku/keys:/validators/teku/passwords \
        --beacon-node-api-endpoints=$CC_URL_STRING \
        --validators-keystore-locking-enabled=false \
        --log-destination=CONSOLE \
        --validators-proposer-default-fee-recipient=$(cat /validators/$FEE_RECIPIENT_FILE) \
        $VC_ADDITIONAL_FLAGS"

    if [ "$ENABLE_MEV_BOOST" = "true" ]; then
        CMD="$CMD --validators-builder-registration-default-enabled=true"
    fi

    if [ "$ENABLE_METRICS" = "true" ]; then
        CMD="$CMD --metrics-enabled=true --metrics-interface=0.0.0.0 --metrics-port=$VC_METRICS_PORT --metrics-host-allowlist=*"
    fi

    if [ "$ENABLE_BITFLY_NODE_METRICS" = "true" ]; then
        CMD="$CMD --metrics-publish-endpoint=$BITFLY_NODE_METRICS_ENDPOINT?apikey=$BITFLY_NODE_METRICS_SECRET&machine=$BITFLY_NODE_METRICS_MACHINE_NAME"
    fi

    exec ${CMD} --validators-graffiti="$GRAFFITI"

fi

