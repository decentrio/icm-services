# Deploying the Sepolia ⇄ Fuji bridge services

This runbook moves the Sepolia ⇄ Fuji demo from a developer laptop to servers. The contracts are already deployed on both testnets, so only off-chain services and a static UI need to run.

> **Testnet only.** The light client's proving keys come from a development setup, not a ceremony, and the validator-set registry on Sepolia is fixed at one Fuji height (see [Limits](#limits)). Do not reuse this setup for mainnet.

## What runs

| Service | Role | Talks to |
|---|---|---|
| **ICM relayer** | Delivers messages both ways. Fuji → Sepolia: collects Fuji validators' BLS signatures and delivers on Sepolia (≈ 10–30 s). Sepolia → Fuji: waits until the ZK light client covers the message, builds the proof, delivers on Fuji (≈ 15–20 min, mostly Ethereum finality). | Sepolia EL + CL, Fuji C-Chain RPC, local Fuji node, Fuji validators (p2p 9651) |
| **Light-client updater** | Keeps the ZK light client of Sepolia on Fuji current. Proves a Step only when a message waits (and at least hourly), and a Rotate once per sync-committee period (≈ 27 h). | Sepolia CL (light-client API) and EL, Fuji C-Chain RPC |
| **Fuji node** (avalanchego, P-Chain only) | Serves `platform.getValidatorsAt` at past heights, which the public API refuses. | Fuji p2p network |
| **Sepolia node** (Geth + Nimbus) | Already run by the team at `37.27.127.216`. Must serve `eth_getLogs`/receipts for recent blocks, WebSocket log subscriptions, the beacon light-client API, and SSZ beacon states of recent finalized slots (`/eth/v2/debug/beacon/states/{slot}`). | — |
| **UI** (static) | Bridge and monitor pages. Reads public RPCs from the browser; no backend. | Public Sepolia and Fuji RPCs |

Deployed contracts (addresses in `deployment.json` on the laptop):

| Contract | Chain | Address |
|---|---|---|
| ZK light client of Sepolia | Fuji | `0xb72fefdfde361319ebb54bb2f9fef5fa70b42efe` |
| TeleporterMessengerV2 | both | `0xc7e56e3861b19435e4ec5e36035a808a9db51538` |
| LightClientAdapter | both | `0xf76523ba1dc38a930c63a0e64ca5efd853f52413` |
| MerkleValidatorSetRegistry | both | `0x68fc991fd5e846f3c9d969af08eb5fe3e39d16d0` |
| EXMP (token) / ERC20TokenHome | Sepolia | `0xfa476f2a25c82ebcd10568896f886267cf2fa347` / `0x6677Cf31dAf512eA27353B0eFd296f59d426d0F6` |
| EXMP.e (ERC20TokenRemote) | Fuji | `0xF4D2615B9FDE71B848d424D27714c19E8739C707` |

## Machine

One Linux host (x86-64) for the relayer, the updater and the Fuji node:

- **CPU / RAM:** 8+ vCPU, 16 GB RAM. A Step proof takes ≈ 9 s on 10 cores; a Rotate ≈ 25 s and ≈ 7 GB with `GOMEMLIMIT=6GiB`.
- **Disk:** 20 GB SSD (proving keys 3.7 GB, Fuji P-Chain data ≈ 2 GB and growing slowly, relayer database small).
- **Inbound:** TCP 9651 (avalanchego p2p) from anywhere is recommended. Keep 9650 (avalanchego API) on localhost. Expose 18080 (relayer health) and 19090 (Prometheus metrics) only to monitoring.
- **Outbound:** HTTPS/WSS to the Fuji RPCs below, TCP 8545/8546/3500 to the Sepolia node, and TCP 9651 to Fuji validators (signature aggregation).

## Secrets

Never commit them; hand them over through the secret manager.

| Key | Used by | Needs | Notes |
|---|---|---|---|
| **Relayer key** | relayer: `delivery-private-key` (Sepolia) and `account-private-key` (Fuji) | Sepolia ETH and Fuji AVAX (0.1 each lasts long at current gas prices) | Any key. Currently the demo wallet `0x97B1863e137D8B78C7308Dd82503450de2B2142c`; a dedicated key is better. |
| **Updater key** `0x51Fd3FD82397De23210b28A42e2b9c3C335A8613` | updater (`RELAYER_KEY`) | Fuji AVAX (≈ 0.1) | Must be allowlisted on the light client. A different key needs `setRelayer(<address>, true)` from the light client's admin (the demo wallet). |
| **Registry-updater key** | relayer `private-key` | nothing for now | Must differ from the relayer key. Its updates cannot succeed on Fuji today (see Limits), so it is never used to send. |

Run **one** updater per updater key and **one** relayer per relayer key: two instances with the same key collide on nonces.

## Artifacts to copy from the laptop

- **Proving keys:** `avax-light-client/keys/` (`step_512.*`, `rotate_512.*`, 3.7 GB) to `/var/lib/lc-updater/keys`. They match the verifiers deployed on Fuji. **Do not regenerate them**: new keys need new verifiers and a new light client.

## Build

Go 1.26 for both. The light-client module is private (`GOPRIVATE=github.com/decentrio`, GitHub access to `decentrio`).

```sh
# Fuji node: release binary
curl -LO https://github.com/ava-labs/avalanchego/releases/download/v1.15.2/avalanchego-linux-amd64-v1.15.2.tar.gz
tar xzf avalanchego-linux-amd64-v1.15.2.tar.gz && install avalanchego-v1.15.2/avalanchego /opt/bridge/bin/

# Updater: commit dd95dd5 pins the decentrio/gnark fork (a6330801) the proving keys were made with
git clone git@github.com:decentrio/avax-light-client.git && cd avax-light-client
git checkout dd95dd5 && go build -o /opt/bridge/bin/updater ./cmd/updater

# Relayer: icm-services fork, branch hieu/sepolia-fuji-demo (1b56d2356 or later)
git clone git@github.com:decentrio/icm-services.git && cd icm-services
git checkout hieu/sepolia-fuji-demo && go build -o /opt/bridge/bin/icm-relayer ./relayer/main
```

## Configure and run

1. Create a `bridge` user and the directories `/var/lib/avalanchego`, `/var/lib/lc-updater/keys`, `/var/lib/icm-relayer` and `/etc/bridge` (owned by `bridge`, `/etc/bridge` mode 700).
2. Copy the units from [`deploy/`](deploy) to `/etc/systemd/system/`. Replace `<SEPOLIA_EL>` and `<SEPOLIA_CL>` with the Sepolia node (`37.27.127.216`).
3. `/etc/bridge/updater.env`: `RELAYER_KEY=0x<updater key>`, mode 600.
4. `/etc/bridge/relayer.json`: start from [`deploy/relayer.example.json`](deploy/relayer.example.json), fill in `<SEPOLIA_EL>`, `<SEPOLIA_CL>`, `<RELAYER_KEY>`, `<RELAYER_ADDRESS>` and `<REGISTRY_UPDATER_KEY>`, mode 600.
5. Start in order:

```sh
systemctl enable --now avalanchego-fuji
# wait until the P-Chain is bootstrapped (≈ 10 min)
curl -s -X POST -H 'content-type:application/json' \
  --data '{"jsonrpc":"2.0","id":1,"method":"info.isBootstrapped","params":{"chain":"P"}}' \
  http://127.0.0.1:9650/ext/info          # → "isBootstrapped":true
systemctl enable --now lc-updater          # loads the keys (≈ 3 min), then ticks every 12 s
systemctl enable --now icm-relayer
```

6. **Stop the laptop's processes** once these run (same keys).

### UI

The pages live in `decentrio/Dividend-Bridge-UI`, branch `feat/fuji-sepolia-monitor`. Build them with `npm ci && npm run build`, then serve `dist/` from any static host over HTTPS. The Fuji ⇄ Sepolia pages need no backend and no keys: they read public RPCs from the browser, and the Bridge page signs with the user's wallet. The Sepolia node at `37.27.127.216` cannot be used from the browser, because it has no CORS headers and no HTTPS.

### Endpoint choices that matter

- **Fuji C-Chain:** the relayer reads Fuji logs through `avalanche-fuji-c-chain-rpc.publicnode.com`. The official `api.avax-test.network` caps `eth_getLogs` at 100 blocks while the relayer catches up 200 at a time, which crashed it. It sends Fuji transactions through `api.avax-test.network`, because publicnode does not serve pending-block state.
- **Sepolia:** use the team's node. The public Lodestar node sometimes cannot serve the anchor's beacon state (`REGEN_ERROR_NO_SEED_STATE`), which stalls Sepolia → Fuji deliveries.
- **`tx-inclusion-timeout-seconds: 90`** on the Sepolia destination. With 30 s the relayer resent deliveries that had landed one or two blocks late.

## Check it works

| Check | Command / signal |
|---|---|
| Fuji node | `info.isBootstrapped` (P) is `true` |
| Relayer up | `curl -s localhost:18080/health` → `"status":"up"`; log `Initialization complete`, then `Connected to sufficient stake` |
| Updater up | log `stepping on demand for messages from 0x59D6…`, then `tick done … latest finalized slot N` every 12 s |
| Light client | `cast call 0xb72fefdfde361319ebb54bb2f9fef5fa70b42efe 'latestFinalizedSlot()(uint64)' --rpc-url https://api.avax-test.network/ext/bc/C/rpc` advances after each `submitStep` |
| End to end | Send 1 EXMP Sepolia → Fuji and 0.5 EXMP.e back (`go run ./scripts/tools/sepolia-fuji-ictt send …`, or the UI's Bridge page). Relayer logs `Delivered message to destination chain` for each; the UI monitor shows both as Delivered. |

Alert on:

- relayer health not `up`, or no `Delivered message` for a message older than 30 min (Sepolia → Fuji) or 2 min (Fuji → Sepolia);
- updater log `tick:` errors for more than 10 min, or no `submitStep` for more than 2 h;
- relayer-key ETH < 0.01 or AVAX < 0.05; updater-key AVAX < 0.02.

## Limits

- **Validator set on Sepolia is frozen** at Fuji P-Chain height 303781 (59 validators). Avalanche validators running stock avalanchego do not sign the registry's update message, so it cannot follow Fuji's validator set. Fuji → Sepolia keeps working while validators holding ≥ 67% of that set's stake stay online; at current stake end times that holds until about **March 2027**. When deliveries start failing on insufficient stake, ask the dev team to redeploy the registry.
- **Development proving keys**, not from a ceremony: fine for testnet, not for production.
- **Rotate** must land at least once per sync-committee period (≈ 27 h); the updater does it automatically. If the updater is down for longer, it catches up on restart as long as the beacon node still serves the missed periods' light-client updates.
- The demo traffic loop (`sepolia-fuji-run/loop.sh` on the laptop) is not part of the deployment.
