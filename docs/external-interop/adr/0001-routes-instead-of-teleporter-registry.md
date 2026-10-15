# ICTT v2 trusts a route per counterparty chain, not TeleporterRegistry

ICTT v1 finds its messenger in `TeleporterRegistry`, which trusts every registered messenger for every chain and is managed through Warp, so it cannot exist on Ethereum. Once a Warp messenger and a bridge messenger run on the same chain, a registry would let each deliver the other's traffic. ICTT v2 instead keeps its own route for each counterparty chain (one outbound messenger and a set of inbound messengers), managed by the contract's owner.

## Considered Options

- **A registry on Ethereum that does not use Warp.** It would still trust every messenger for every chain.
- **One fixed messenger per ICTT contract.** A Home could not serve Avalanche L1s and Ethereum at the same time, and moving to a new bridge messenger would mean redeploying the token contracts.

## Consequences

- The owner of the routes controls minting and releasing, so it must be a multisig behind a timelock.
- A chain must be routed before a Remote on it can register with a Home. Any Avalanche chain could register in v1.
- The storage of upgradeable ICTT contracts now holds routes, so changing this later is a storage migration.
