# ICTT v2 from Sepolia to Fuji, verified by the ZK light client

This tool moves an ERC-20 from Ethereum Sepolia (the token's home) to Avalanche Fuji C-Chain (a remote) with ICTT v2.

- **Verification.** On Fuji, a `LightClientAdapter` verifies each message against the beacon roots of Decentrio's ZK light client of Sepolia. There is no trusted relayer and no third-party prover.
- **Relaying.** The tool proves the message's log against the light client's latest finalized root (the Gloas receipt path) and delivers it to `TeleporterMessengerV2` on Fuji. The remote then mints the tokens.

## Run it

You need:
- an account with **Sepolia ETH and Fuji AVAX**, about 0.05 ETH and 1 AVAX (gas is currently a fraction of that);
- Go, Foundry and Python 3;
- `icm-services` and `avax-light-client` checked out side by side, with access to the private `github.com/decentrio` Go modules;
- about 16 GB of free memory for the light client's prover.

From the `icm-services` root:

```sh
KEY=0x<private key> ./scripts/tools/sepolia-fuji-ictt/run.sh
```

`run.sh` keeps its state in `./sepolia-fuji-run` (`WORKDIR`) and skips every step that is already done. If a step fails, run it again. To send more tokens later, run it again with `AMOUNT=5`.

| Step | What happens | Time |
|---|---|---|
| 1. Light client | Takes the root of trust from Sepolia's latest finalized checkpoint and deploys the verifiers and the light client on Fuji. Then it starts the light client's updater, which keeps running: a Step each epoch, a Rotate each period. The updater sends from its own key, which the script creates, funds with `UPDATER_FUNDING` AVAX (default 0.5) and allowlists. The first time, it also makes the development proving keys. | 1-2 min; plus about 15 min once for the keys |
| 2. Bridge | Deploys `MerkleValidatorSetRegistry`, `LightClientAdapter`, `Adapter` and `TeleporterMessengerV2` with CREATE2, at the same address on both chains, and initializes the messengers. | 1-2 min |
| 3. Token | Deploys a demo ERC-20 (minted to the account) and its `ERC20TokenHome` on Sepolia, routed to Fuji through the messenger. Then it deploys the `ERC20TokenRemote` on Fuji. | 1 min |
| 4. Register | Registers the remote with the home through a `RemoteRegistrar` (see [Registration](#registration)). | 1 min |
| 5. Send | Approves the home and sends `AMOUNT` tokens to the account on Fuji. | 30 s |
| 6. Deliver | Waits until Sepolia finalizes the transfer and the light client covers it. Then it proves the message and delivers it on Fuji, where the remote mints. | 15-45 min |
| 7. Balances | Shows the remote balance on Fuji, and what the home holds for Fuji on Sepolia. | |

The commands can also be run one by one:

```sh
go run ./scripts/tools/sepolia-fuji-ictt <command> -h
```

The commands are `deploy-bridge`, `deploy-ictt`, `register`, `send`, `relay` and `balance`. Addresses are kept in `deployment.json`.

### Rehearse on forks first

The whole flow runs against local forks of both chains:

```sh
anvil --fork-url https://ethereum-sepolia-rpc.publicnode.com --port 8548
anvil --fork-url https://api.avax-test.network/ext/bc/C/rpc --port 8547 --hardfork cancun
FORKS=1 SEPOLIA_RPC=http://127.0.0.1:8548 FUJI_RPC=http://127.0.0.1:8547 \
    KEY=0x47e179ec197488593b187f80a00eb0da91f1b9d0b13f8733639f19c30a34926a \
    ./scripts/tools/sepolia-fuji-ictt/run.sh
```

- **What differs on forks.** A fork's transactions are not in Sepolia's beacon chain, so no light-client root covers them. Step 6 therefore runs `rehearse-delivery` instead of `relay`. That command is built only with `-tags rehearsal`.
- **How the rehearsal delivers.** It swaps the `LightClientAdapter`'s code on the Fuji fork for one that accepts any message, delivers, and restores the code.
- **What it covers.** Everything except the proof. The proof itself is exercised on real Sepolia data by `../sepolia-fuji-demo`.

## Registration

A `TokenRemote` must be registered with its home before the home sends to it. ICTT does this with a message from the remote, here from Fuji to Sepolia. That direction cannot run on public Fuji without our own Fuji node (see [Fuji to Sepolia](#fuji-to-sepolia)).

The home's owner therefore registers the remote with a `RemoteRegistrar` (`icm-contracts/common/ictt-v2/RemoteRegistrar.sol`):
1. The owner adds the registrar as the home's messenger for Fuji.
2. The registrar delivers the registration message the remote would have sent.
3. The owner removes the registrar.

The home trusts it as it trusts any messenger its owner adds, and still applies its own registration rules. The ICTT contracts are unchanged. Transfers themselves are always verified by the light client.

## Fuji to Sepolia

The bridge deployment already includes the Sepolia side of the return direction. The `MerkleValidatorSetRegistry` holds Fuji's primary network validator set, at the P-Chain height recorded in `deployment.json`, under the C-Chain's ID.

Running the ICM relayer for that direction (`relayer/`) needs three things:

- **A Fuji node of our own as its P-Chain API.** The public P-Chain API refuses the historical validator queries the relayer makes for each message. A node with `--partial-sync-primary-network` is enough.
- **Outbound access to the Fuji validators' staking port (9651)**, so the signature aggregator can reach them.
- **A configuration with one `external-evm-destinations` entry**:
  - `contract-address` is the registry;
  - `blockchain-id` is Fuji C-Chain (`yH8D7ThNJkxmtkuv2jgBa4P1Rn3Qpr4pPr7QYNfcdoS6k6HWp`);
  - `subnet-id` is the primary network;
  - `teleporter-address` is the messenger;
  - the source's `message-contracts` entry is the registry, with the `teleporter-v2` format.

The registry's set is frozen at that height. Public Fuji validators will not sign the validator-set updates the registry expects, so their weight must keep at least 67% of the frozen set's weight online. At deployment that means the 13 heaviest of 65 validators.

## Limits

- **Development keys.** The light client's verifiers come from a development setup, not a ceremony.
- **Public nodes.** The default endpoints have no SLA. The Sepolia beacon node serves recent states quickly and older ones slowly, so deliver soon after a transfer is finalized.
- **The registrar.** Its registration trusts the home's owner, as described above.
