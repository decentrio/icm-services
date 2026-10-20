# Sepolia → Fuji demo: an Ethereum token transfer proven on Avalanche

This demo proves a real ERC-20 transfer on Sepolia on an Avalanche chain. The proof uses beacon roots from Decentrio's ZK light client, which follows Sepolia's finality with Groth16 proofs of the sync committee.

The demo uses the parts of the ICM path from Ethereum to Avalanche that do not depend on the messaging layer:

1. The ZK light client (`avax-light-client`) runs on the Avalanche chain. It stores Sepolia's finalized beacon roots by slot:
   - a Step proof every epoch;
   - a Rotate proof once per sync period.
2. `LightClientAdapter` reads those roots through `IBeaconRootSource`. It then verifies a log with the Gloas execution proof:
   - anchor block → `state_roots` → state of the first block that built on the payload;
   - that state's `latest_block_hash` → the execution header → its receipts root;
   - the receipt → the log.
3. This tool picks the transfer, builds the proofs from Sepolia nodes, and submits them.

ICTT messages take the same verification path. To run ICTT transfers end to end, `TeleporterMessengerV2` must be on Sepolia (see [Limits](#limits)).

## Run it

You need Go, Foundry and both repositories. The commands below use a local fork of Fuji C-Chain. To use Fuji itself, point `--rpc-url`/`-rpc` at `https://api.avax-test.network/ext/bc/C/rpc` and use a funded key.

### 1. Start a Fuji fork

```sh
anvil --fork-url https://api.avax-test.network/ext/bc/C/rpc --port 8547 --hardfork cancun
```

Use anvil's first two accounts:
- `KEY0` for the light client's admin and relayer;
- `KEY1` for the demo.

### 2. Deploy the light client and keep it following Sepolia

Run these in `avax-light-client`, with `GOPRIVATE=github.com/decentrio`.

1. Make the development keys and verifiers, once. This takes about 15 minutes and peaks at 11 GB of memory. It is not a ceremony: use it for testnets only.

   ```sh
   go run ./cmd/devsetup -circuit step -n 512 -out keys -verifier contracts/src/dev/Step512Verifier.sol
   go run ./cmd/devsetup -circuit rotate -n 512 -out keys -verifier contracts/src/dev/Rotate512Verifier.sol
   ```

2. Take the root of trust from Sepolia's latest finalized checkpoint. The tool checks it against the network's genesis and fork schedule.

   ```sh
   go run ./cmd/bootstrap -out contracts/deploy/sepolia.json
   ```

3. Deploy the verifiers and the light client:

   ```sh
   (cd contracts && BOOTSTRAP=deploy/sepolia.json forge script script/Deploy.s.sol \
       --rpc-url http://127.0.0.1:8547 --private-key $KEY0 --broadcast)
   ```

4. Run the updater, which submits a Step each epoch and a Rotate each period:

   ```sh
   RELAYER_KEY=$KEY0 go run ./cmd/updater -rpc http://127.0.0.1:8547 -contract <light client> -keys keys
   ```

### 3. Prove a transfer

Run this in `icm-services`:

```sh
DEMO_KEY=$KEY1 go run ./scripts/tools/sepolia-fuji-demo -rpc http://127.0.0.1:8547 -light-client <light client>
```

It works in this order:
1. It deploys a `LightClientAdapter` for Sepolia, or uses `-adapter`.
2. It waits until the light client's latest finalized root covers a recent USDC transfer. Use `-token` for another token, or `-tx` for a given transaction.
3. It builds the proofs and has the adapter verify the transfer's log.

## A run (2026-10-21, Fuji fork at block 59,743,481)

```text
chain 43113: light client 0xb3e0586F61a2dC5C1055528f0939E2fbDcE38Ad0, adapter 0x7074569F5c6871b2D36cafd9d16120806fC95308
light client anchor: slot 11398624, root 0x9f0ab2fbb3f072393dd4b492a7ff663deeda5b84c89747a99d8d5d73f173cf8a
transfer 0x33beb4b64699820092b13e077296b19703c746b466d30d958e912818d75d36b5 in execution block 11956765 (payload slot 11398622)
execution proof in 57s: target slot 11398623, 3 + 21 + 11 SSZ siblings, 669-byte header
receipt proof: 3 trie nodes

verified on chain 43113 in tx 0x3417088fbe200df2ec2a9174434cb62f40efecff3abdad85af3f3371048f8968 (661828 gas)
  ZKEventImported: source chain 11155111, beacon slot 11398623, emitter 0x1c7D4B196Cb0C7B01d743Fbc6116a902379C7238, execution block 0x760cda23cb2584e5cb5787af806b6a43c9b83f6937085f425c0fde9f2b643390, log 8
  Sepolia transfer 0x33beb4b64699820092b13e077296b19703c746b466d30d958e912818d75d36b5: 0x6588c789c6B6dF8f8b6B92e6433F813cc2D1D4d0 -> 0xB62eFc43C16b4E2Bf1aAAD33002368e55708538D, amount 3000000 (raw units)
```

**Independent checks of this run** (all match):
- The execution block hash and `slotNumber` come from the Sepolia RPC.
- The transfer is receipt log 8 of that transaction.
- The anchor root comes from a second beacon node, `ethereum-sepolia-beacon-api.publicnode.com`.

**Measurements from the same run:**
- **Updater:** a Rotate costs 1,558,265 gas, a Step 462,309 gas, and one tick with both took 65 s.
- **Artifacts:** loading the 512-member proving artifacts took 3 to 3.5 minutes.

## Limits

- **Public nodes.**
  - The defaults are public Sepolia nodes without an SLA.
  - The beacon node serves recent states quickly and older ones slowly. The proof needs the anchor state and the target state, so run the demo soon after the transfer is finalized.
  - A testnet deployment needs a beacon node that keeps historical states.
- **Development keys.** The verifiers come from a development setup, not a ceremony.
- **ICTT end to end.** An ICTT transfer from Sepolia to Fuji also needs `TeleporterMessengerV2`, an Ethereum adapter and an ICTT `TokenHome` on Sepolia. On Fuji it needs `TeleporterMessengerV2`, a `LightClientAdapter` and a `TokenRemote`. It also needs funded keys on both chains.
  - A `TokenRemote` registers with its home through a message from Avalanche to Ethereum. That direction needs a verifier on Sepolia (for example the validator set registry), which is outside this design.
