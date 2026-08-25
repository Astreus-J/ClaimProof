# Threat Model — ClaimProof

> Target rigor bar: comparable to the hackathon's technically strongest competitors (`crosscredit`, which publishes a threat model + negative-path tests; `Cr3dX`, with independent cross-verification).

## Assets to protect

1. **Funds in the `ClaimVault` pool** (protection premiums deposited by buyers).
2. **Payout integrity** — the right amount, to the right buyer, exactly once per real event.
3. **Trust in the verification** — the guarantee that a payout only happens if the trigger event actually occurred.

## Actors

| Actor | Assumed trust level |
|---|---|
| Buyer | Untrusted — may attempt to claim a payout without a real failure |
| Store / `DeliveryTrackerMock` operator | Trusted in the demo (it's the team itself) — in production would be replaced by an independent logistics oracle |
| Off-chain worker | Semi-trusted — could be compromised, but cannot forge a valid proof |
| AI Claims Agent | Untrusted for authorization — only suggests a value, capped by an on-chain policy |
| Attestcoin Protocol attestors | Trust inherited from the protocol — today operate in permissioned (`AuthorizedOnly`) mode |

## Attack vectors and mitigations

### T1 — Claim without a real event (forging a fake "DeliveryFailed")
**Attack:** a buyer or the worker tries to submit a claim without the event having actually occurred on Sepolia.
**Mitigation:** `ClaimVault` only accepts the claim if `verifyAndEmit()` on the Attestcoin precompile returns success, which requires a valid Merkle inclusion + continuity proof against the real attested state of Sepolia. This cannot be forged without compromising the protocol's attestor set.
**Status:** mitigated by design (inherited from the Attestcoin Protocol).

### T2 — Replaying the same proof for multiple payouts
**Attack:** resubmitting the proof of an already-processed `DeliveryFailed` to drain the pool again.
**Mitigation:** the `processedQueries` mapping, keyed by `(chainKey, blockHeight, transactionIndex)`, rejects any already-processed proof.
**Status:** mitigated — covered by an automated test (`ReplayProtection.t.sol`).

### T3 — Inclusion proof without checking transaction success
**Attack:** the precompile proves the transaction was *included*, not that it *succeeded*. A naive contract could pay out even if the source transaction reverted.
**Mitigation:** `ClaimVault` explicitly checks the decoded status field (`0x1`) via `EvmV1Decoder` before any transfer.
**Status:** mitigated by design — an explicitly documented gotcha, since it's a common integration mistake with the protocol.

### T4 — AI suggests a manipulated or excessive value
**Attack:** the AI Claims Agent (compromised, prompt-injected, or simply buggy) suggests a payout far above what's reasonable.
**Mitigation:** the contract enforces an on-chain policy cap (a maximum percentage of the original order value, set during `registerOrder`), independent of the AI's suggestion. The AI never has signing authority.
**Status:** mitigated by design — validated by an automated test covering out-of-cap values.

### T5 — Compromise of the worker's private key
**Attack:** an attacker obtains the private key the worker uses to submit claims.
**Mitigation:** even with the key compromised, the attacker can only submit claims that pass Attestcoin re-verification — they cannot generate payouts for events that never happened. The maximum damage is submitting legitimate claims out of order or with delay, not value fraud.
**Status:** residual risk accepted for hackathon scope; production mitigation (HSM/multisig for the worker) is out of MVP scope.

### T6 — `DeliveryTrackerMock` being, by design, team-controlled
**Attack/limitation:** in the demo, the source-of-truth on the origin side (Sepolia) is controlled by the developers themselves — not an independent logistics oracle.
**Mitigation:** explicitly documented as a deliberate MVP simplification (see [../ARCHITECTURE.md](../ARCHITECTURE.md) and [../README.md](../README.md)); the post-hackathon roadmap lists replacing it with a real oracle as a next step. This doesn't invalidate Attestcoin's value in the design — the protocol remains the only way for Creditcoin to trust any event from any source chain, including a future independent data source.
**Status:** known, disclosed limitation, not a hidden vulnerability.

### T7 — Risks inherited from the Attestcoin Protocol's infrastructure
**Description:** the protocol currently operates with a permissioned attestor set (`AuthorizedOnly`) and a `MinBondRequirement` of 0 CTC on mainnet — economic security against attestor collusion isn't fully active yet.
**Mitigation:** outside ClaimProof's control; documented for transparency with judges. ClaimProof inherits exactly the security level the protocol offers today, without adding optimistic assumptions about guarantees that don't exist yet.
**Status:** disclosed platform risk.

## Negative test cases covered (Foundry)

- Submitting a claim with an invalid proof → revert
- Submitting a claim referencing an already-processed transaction (replay) → revert
- Submitting a claim for a source transaction with a failure status (`!= 0x1`) → revert
- Submitting a suggested value above the policy cap → the value is capped, not honored as-is
- Submitting a claim for a nonexistent `orderId` → revert
