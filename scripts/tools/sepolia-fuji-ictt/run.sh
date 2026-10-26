#!/usr/bin/env bash
# Copyright (C) 2026, Ava Labs, Inc. All rights reserved.
# See the file LICENSE for licensing terms.

# Transfers a demo ERC-20 from Ethereum Sepolia to Avalanche Fuji C-Chain with ICTT v2. Fuji
# verifies the transfer against Decentrio's ZK light client of Sepolia. Every step is skipped when
# its result already exists in WORKDIR, so the script can be rerun after a failure.
#
# Usage, from the icm-services root:
#   KEY=0x<private key funded on Sepolia and Fuji> ./scripts/tools/sepolia-fuji-ictt/run.sh
#
# Environment:
#   KEY                private key of an account with Sepolia ETH and Fuji AVAX (required)
#   AVAX_LIGHT_CLIENT  avax-light-client checkout (default: ../avax-light-client)
#   WORKDIR            where the deployment and logs are kept (default: ./sepolia-fuji-run)
#   AMOUNT             whole tokens to send (default: 1)
#   SEPOLIA_RPC, FUJI_RPC, BEACON   endpoints (default: public Sepolia and Fuji endpoints)
#   UPDATER_FUNDING    AVAX given to the light client's updater key (default: 0.5)
#   FORKS=1            rehearsal on local forks of both chains: the message is delivered without a
#                      proof, since no beacon root covers a fork's transactions

set -euo pipefail

ICM=$(cd "$(dirname "${BASH_SOURCE[0]}")/../../.." && pwd)
ALC=$(cd "${AVAX_LIGHT_CLIENT:-$ICM/../avax-light-client}" && pwd)
WORKDIR=${WORKDIR:-$ICM/sepolia-fuji-run}
AMOUNT=${AMOUNT:-1}
SEPOLIA_RPC=${SEPOLIA_RPC:-https://ethereum-sepolia-rpc.publicnode.com}
FUJI_RPC=${FUJI_RPC:-https://api.avax-test.network/ext/bc/C/rpc}
BEACON=${BEACON:-https://lodestar-sepolia.chainsafe.io}
UPDATER_FUNDING=${UPDATER_FUNDING:-0.5}
export GOPRIVATE=github.com/decentrio

: "${KEY:?set KEY to the private key of an account funded on Sepolia and Fuji}"
export KEY
mkdir -p "$WORKDIR"
cd "$WORKDIR"
ACCOUNT=$(cast wallet address "$KEY")
echo "account $ACCOUNT"
echo "  Sepolia: $(cast balance "$ACCOUNT" --ether --rpc-url "$SEPOLIA_RPC") ETH"
echo "  Fuji:    $(cast balance "$ACCOUNT" --ether --rpc-url "$FUJI_RPC") AVAX"

TAGS=""
[ "${FORKS:-}" = 1 ] && TAGS="-tags rehearsal"
# shellcheck disable=SC2086
(cd "$ICM" && go build $TAGS -o "$WORKDIR/sepolia-fuji-ictt" ./scripts/tools/sepolia-fuji-ictt)
ictt() { "$WORKDIR/sepolia-fuji-ictt" -sepolia-rpc "$SEPOLIA_RPC" -fuji-rpc "$FUJI_RPC" -beacon "$BEACON" "$@"; }
field() { python3 -c "import json,sys; print(json.load(open('deployment.json')).get(sys.argv[1], ''))" "$1"; }
zero=0x0000000000000000000000000000000000000000

echo
echo "== 1. ZK light client of Sepolia on Fuji"
if [ ! -s light-client.txt ]; then
    if [ ! -s "$ALC/keys/rotate_512.pk" ] || [ ! -s "$ALC/contracts/src/dev/Rotate512Verifier.sol" ]; then
        echo "making the development proving keys (about 15 minutes, 11 GB of memory)"
        (cd "$ALC" && go run ./cmd/devsetup -circuit step -n 512 -out keys \
            -verifier contracts/src/dev/Step512Verifier.sol)
        (cd "$ALC" && go run ./cmd/devsetup -circuit rotate -n 512 -out keys \
            -verifier contracts/src/dev/Rotate512Verifier.sol)
    fi
    mkdir -p "$ALC/contracts/deploy"
    (cd "$ALC" && go run ./cmd/bootstrap -beacon "$BEACON" -out contracts/deploy/sepolia.json)
    (cd "$ALC/contracts" && BOOTSTRAP=deploy/sepolia.json forge script script/Deploy.s.sol \
        --rpc-url "$FUJI_RPC" --private-key "$KEY" --broadcast --slow) | tee deploy-light-client.log
    grep "light client" deploy-light-client.log | awk '{print $3}' > light-client.txt
fi
LIGHT_CLIENT=$(cat light-client.txt)
echo "light client $LIGHT_CLIENT"

# The updater sends from its own key, so that its transactions never race this script's.
if [ ! -s updater.key ]; then
    cast wallet new --json | python3 -c "import json,sys; print(json.load(sys.stdin)[0]['private_key'])" > updater.key
    chmod 600 updater.key
fi
UPDATER=$(cast wallet address "$(cat updater.key)")
if [ "$(cast call "$LIGHT_CLIENT" "relayers(address)(bool)" "$UPDATER" --rpc-url "$FUJI_RPC")" != true ]; then
    cast send "$UPDATER" --value "${UPDATER_FUNDING}ether" --private-key "$KEY" --rpc-url "$FUJI_RPC" > /dev/null
    cast send "$LIGHT_CLIENT" "setRelayer(address,bool)" "$UPDATER" true \
        --private-key "$KEY" --rpc-url "$FUJI_RPC" > /dev/null
fi
if [ ! -s updater.pid ] || ! kill -0 "$(cat updater.pid)" 2>/dev/null; then
    (cd "$ALC" && go build -o "$WORKDIR/updater" ./cmd/updater)
    RELAYER_KEY=$(cat updater.key) nohup ./updater -beacon "$BEACON" -rpc "$FUJI_RPC" \
        -contract "$LIGHT_CLIENT" -keys "$ALC/keys" > updater.log 2>&1 &
    echo $! > updater.pid
    echo "updater started (pid $(cat updater.pid), log $WORKDIR/updater.log); it keeps running"
fi

echo
echo "== 2. Bridge contracts on both chains"
if [ ! -s deployment.json ] || [ "$(field messenger)" = "" ]; then
    ictt deploy-bridge -light-client "$LIGHT_CLIENT"
fi

echo
echo "== 3. Token home on Sepolia, remote on Fuji"
if [ "$(field tokenRemote)" = "" ] || [ "$(field tokenRemote)" = "$zero" ]; then
    ictt deploy-ictt
fi

echo
echo "== 4. Register the remote with the home"
ictt register

echo
echo "== 5. Send $AMOUNT tokens from Sepolia to Fuji"
ictt send -amount "$AMOUNT" | tee send.log
TX=$(grep -o 'relay -tx 0x[0-9a-f]*' send.log | awk '{print $3}')

echo
echo "== 6. Deliver on Fuji"
if [ "${FORKS:-}" = 1 ]; then
    ictt rehearse-delivery -tx "$TX"
else
    echo "waiting for Sepolia to finalize the transfer and the light client to cover it (15-45 min)"
    ictt relay -tx "$TX"
fi

echo
echo "== 7. Balances"
ictt balance
