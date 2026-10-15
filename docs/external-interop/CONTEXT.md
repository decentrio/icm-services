# External Interop

Moving ICM messages and ICTT tokens between Avalanche chains and Ethereum over the bridge built on `TeleporterMessengerV2`.

## Language

### Bridge

**Connection**:
The bridge between one Avalanche chain and Ethereum. It carries messages in both directions.
_Avoid_: two-way messenger, pair, link

**Direction**:
One way across a connection: Avalanche → Ethereum, or Ethereum → Avalanche. Each direction has its own verifier.
_Avoid_: path, leg, lane

**Verifier**:
The bridge part that proves, on the destination chain, that a message was sent on its source chain. Validator signatures checked by the Merkle validator-set registry for Avalanche → Ethereum; a ZK proof of Ethereum consensus for Ethereum → Avalanche.
_Avoid_: authenticator, prover (the prover only produces the proof)

**Bridge messenger**:
The `TeleporterMessengerV2` deployment that carries one connection's messages. It sits at the same address on both chains.
_Avoid_: V2 messenger, two-way messenger

**Warp messenger**:
The Teleporter V1 messenger, which carries messages between Avalanche chains using Warp.
_Avoid_: V1 messenger, registry messenger

**Relayer**:
The off-chain service that delivers messages across a connection and keeps each verifier's view of the other chain up to date.

**Blockchain ID**:
The 32-byte identifier of a chain. Avalanche chains use their Avalanche blockchain ID; Ethereum uses its EVM chain ID as a 32-byte value.
_Avoid_: chain ID (unless meaning the EVM chain ID)

### Routing

**Counterparty chain**:
The chain at the other end of a message, seen from one ICTT contract.
_Avoid_: remote chain (a Remote is a token contract), destination chain (it may be the source)

**Route**:
What an ICTT contract trusts for one counterparty chain: one outbound messenger and a set of inbound messengers.
_Avoid_: messenger table, messenger binding, messenger-per-chain

**Outbound messenger**:
The messenger a route sends through.

**Inbound messenger**:
A messenger a route accepts deliveries from. A route has two only while one messenger replaces another.
_Avoid_: allowed messenger, trusted messenger

**Paused messenger**:
A messenger that no route may send or receive through, on any counterparty chain.

### Token transfer

**Home**:
The ICTT contract on the chain where the token originally lives. It holds the tokens locked for its Remotes.
_Avoid_: hub, source, vault

**Remote**:
The ICTT contract that represents the token on another chain. It mints when tokens arrive and burns when they leave.
_Avoid_: spoke, wrapped token, destination

**Backing rule**:
Every token a Remote has minted is backed one-for-one by a token the Home has locked for that Remote.

**Multi-hop**:
A transfer from one Remote to another that passes through their Home.
_Avoid_: forwarding, routing (a route is something else)
