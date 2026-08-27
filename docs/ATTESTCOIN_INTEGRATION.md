# Attestcoin Protocol Integration Summary

> Technical document required by the BUIDL CTC 2026 Fall submission — explains specifically how and where ClaimProof uses the Attestcoin Protocol. Written to be verifiable by a judge who opens the repository, following the transparency standard set by the competitor `index41`.

## Why Attestcoin is core, not decorative

Test applied (inherited from the competitive analysis, see [decision.md](decision.md)): **"remove Attestcoin — does the project lose its inputs or its ability to act?"**

For ClaimProof, the answer is literally yes: without Attestcoin, `ClaimVault` would have no cryptographically trustworthy way of knowing a delivery failed on another chain. The AI would have to trust unverified text or an API to authorize a payment — exactly the attack vector (fake data, prompt injection, a malicious operator) the product exists to eliminate.

## Where the protocol is used

### 1. Read layer (readability) — the only mode used

ClaimProof uses exclusively the **read** direction of Attestcoin (proof that an event occurred on Ethereum Sepolia, verified on Creditcoin). The **write** direction (outbound messages from Creditcoin to another chain) is **not used**, because the protocol's own documentation describes it as not production-ready ("undergoing 3rd party testing and audits").

### 2. Prover HTTP client — off-chain side (worker, Go)

There is no official Go SDK for Attestcoin. The only official SDK, `@gluwa/usc-sdk`, is TypeScript-only — and it is itself just a thin wrapper around a hosted REST API. ClaimProof's `internal/proofbuilder` package (Go) talks to that same API directly over HTTP, reimplementing only the two calls the worker actually needs:

| Client method (`internal/proofbuilder`) | Equivalent SDK call | Use in ClaimProof |
|---|---|---|
| `WaitUntilHeightAttested(ctx, chainKey, blockNumber)` | `ProofBuilder.waitUntilHeightAttested()` | Polls the Prover API (`https://prover.cc3-testnet.creditcoin.network`) until the Sepolia block containing the `DeliveryFailed` event is attested, before attempting to generate the proof |
| `GetProof(ctx, txHash)` | `ProofBuilder.getProof()` | Obtains the Merkle inclusion proof + continuity proof for the `DeliveryFailed` transaction |

The exact request/response JSON schema was confirmed by reading the official TypeScript SDK's source (`@gluwa/usc-sdk`, via `gluwa/attestcoin-protocol-examples`) — used as a reference only, never as a dependency, since `internal/proofbuilder` reimplements the same three REST calls directly in Go against `https://prover.cc3-testnet.creditcoin.network`:

| Endpoint | Method | Purpose |
|---|---|---|
| `/api/v1/proof-by-tx/{chainKey}/{transactionHash}` | `GET` | Returns the inclusion + continuity proof for one transaction |
| `/api/v1/proof-batch-by-tx/{chainKey}` | `POST`, body: JSON array of transaction hashes | Same, for up to 10 transactions in one call |
| `/api/v1/attested-height/{chainKey}` | `GET` | Returns the latest block height the Prover has attested and cached for that chain, as `{ "attestedHeight": number }` |

`GET /api/v1/proof-by-tx/{chainKey}/{transactionHash}` response shape (this is what `internal/proofbuilder.Proof` models):

```jsonc
{
  "chainKey": 1,
  "headerNumber": 123456,
  "txIndex": 0,
  "txHash": "0x...",
  "txBytes": "0x...",           // raw encoded transaction, decoded on-chain by EvmV1Decoder
  "continuityProof": {
    "lowerEndpointDigest": "0x...",
    "roots": ["0x...", "0x..."]
  },
  "merkleProof": {
    "root": "0x...",
    "siblings": [{ "hash": "0x...", "isLeft": true }]
  },
  "cached": false,
  "generatedAt": "2026-08-26T00:00:00.000Z"
}
```

`waitUntilHeightAttested` is not a separate endpoint — it's a client-side poll loop (15s interval, ~15min timeout in the SDK's defaults) that repeatedly calls `attested-height` until `attestedHeight >= targetHeight`, per the SDK's own doc comment: attestation of a Sepolia-age block typically takes ~8–10 minutes.

### 3. Contract (`ClaimVault.sol`) — on-chain side (Creditcoin)

| Interface / Function | Use in ClaimProof |
|---|---|
| `INativeQueryVerifier.verifyAndEmit(chainKey, blockHeight, txBytes, merkleProof, continuityProof)` on the native `0x0FD2` precompile | Synchronous call, within the `submitClaim` transaction itself, that re-verifies inclusion and continuity of the `DeliveryFailed` transaction before any fund transfer |
| `INativeQueryVerifier.calculateTxIndex(merkleProof)` | Derives the transaction's index within its block, mixed into the anti-replay key (see below) — multiple transactions in the same block share a merkleRoot, so the root alone isn't a unique key |
| `EvmV1Decoder.decodeReceiptFields` / `getLogsByEventSignature` (decoding library) | Extracts the receipt **status** field (`0x1` = success, checked explicitly since the precompile only proves inclusion) and the `DeliveryFailed` log's `orderId`/`buyer` from its topics |
| Emitter-address check (`log.address_ == sourceContract`) | `EvmV1Decoder` filters logs by event signature only, not by which contract emitted them — `ClaimVault` additionally requires the log come from the trusted `DeliveryTrackerMock` address, or any Sepolia contract could forge a same-signature event (see docs/THREAT_MODEL.md, T9) |
| `processedQueries` (mapping, keyed by `keccak256(chainKey, blockHeight, txIndex)`) | Prevents the same proof from being resubmitted to generate multiple payouts (anti-replay protection, T2) |

**Dependency note:** `INativeQueryVerifier` is self-contained boilerplate every Attestcoin dApp defines itself (`contracts/src/interfaces/INativeQueryVerifier.sol`) — Gluwa doesn't publish it as an installable package. `EvmV1Decoder` is a real library from the `@gluwa/usc-contracts` npm package (v0.1.2, MIT), but that package has no public git repository to `forge install` from, so it's vendored verbatim at `contracts/lib/usc-contracts/decoding/EvmV1Decoder.sol` with a provenance comment.

## Full verification flow

1. `DeliveryFailed` is emitted on `DeliveryTrackerMock` (Ethereum Sepolia).
2. The worker waits for the corresponding block's attestation and obtains the proof via the `internal/proofbuilder` HTTP client.
3. The worker calls `ClaimVault.submitClaim(...)`, passing the raw proof.
4. `ClaimVault` calls `verifyAndEmit()` on the `0x0FD2` precompile — if verification fails (invalid proof, unattested block, broken continuity), the whole transaction reverts and no funds move.
5. If verification succeeds, `ClaimVault` decodes the transaction via `EvmV1Decoder`, confirms the `0x1` status, checks `processedQueries`, and only then releases the payout.

## Chains involved

| Role | Chain | Note |
|---|---|---|
| Source (verified event) | Ethereum Sepolia | The only source chain the Attestcoin Protocol supports today, alongside Ethereum Mainnet — the "any chain" marketing language doesn't yet reflect the actual implementation |
| Execution (verification + payout) | Creditcoin CC3 Testnet | Where the `0x0FD2` precompile runs natively |

## Cost and latency (documented, not blindly estimated)

The gas-cost formula published in the protocol's documentation (`attestcoin-protocol/attestcoin-readability/gas-costs.md`) shows that verification cost **grows with proof age** (more continuity hashes to walk). ClaimProof's worker is designed to submit the claim as soon as the attestation is available, minimizing this cost — an architectural decision made directly in response to this documented protocol characteristic, not a generic assumption.

**Update from live Sprint 5 testing:** proof age turned out to matter beyond gas cost. Several claims whose `DeliveryFailed` transactions landed in the same narrow block range — because they were triggered within seconds of each other — shared a continuity proof anchored to the same Attestcoin checkpoint; once that checkpoint rotated before `submitClaim` ran, all of them reverted with `Continuity proof does not match attestation or checkpoint`, reproduced even against a freshly-refetched proof. See [THREAT_MODEL.md](THREAT_MODEL.md) T11 for the full writeup — worth knowing for anyone else integrating against this protocol version, since it isn't obvious from the gas-cost documentation alone.

## Known protocol limitations that shaped this design

- Writability is not used (not production-ready) — see above.
- Only Ethereum/Sepolia as a source — `DeliveryTrackerMock` was deliberately designed for this chain.
- Attestor set in `AuthorizedOnly` mode and a `MinBondRequirement` of 0 CTC on mainnet — protocol-level infrastructure risks, documented in [../SECURITY.md](../SECURITY.md) and [THREAT_MODEL.md](THREAT_MODEL.md), outside ClaimProof's control.
