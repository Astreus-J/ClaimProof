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
**Mitigation:** two independent layers. Off-chain, `internal/claimsagent.Agent.SuggestPayout` clamps the LLM's suggested percentage to [0, 100] and the resulting amount to a configured policy cap before ever building a transaction. On-chain, `ClaimVault.submitClaim`'s `suggestedPayout` parameter is bounded by `min(suggestedPayout, protectionAmount, payoutCap)` regardless of what either layer above computed — the AI never has signing authority, and a compromised worker relaying an inflated `suggestedPayout` still cannot exceed the registered order's own protection amount or the on-chain cap.
**Status:** mitigated by design — validated by automated tests at both layers (`claimsagent` tests for the off-chain cap; `ClaimVault.t.sol`'s `test_SubmitClaim_SuggestedValueAbovePolicyCap_IsCappedNotHonored` for the on-chain cap) and by a real end-to-end run where the Gemini-backed claims agent suggested a genuine payout, submitted automatically by the worker with no manual step.

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

### T8 — Self-registering an order to drain the pool with a genuine, self-triggered event
**Attack:** `DeliveryTrackerMock.createShipment` and `reportDeliveryFailure` are permissionless by design (see T6). If `ClaimVault.registerOrder` were also unrestricted, an attacker could register an order with themselves as `buyer` and an arbitrary `protectionAmount`, then create and fail their own shipment on Sepolia — a real, honestly-attestable `DeliveryFailed` event — and call `submitClaim` to drain the pool for the cost of gas, repeatable per `orderId`.
**Mitigation:** `registerOrder` is restricted to the authorized `worker` address (`onlyWorker`) — it is the only gate on which `(orderId, buyer, protectionAmount)` triples `submitClaim` will ever honor a payout for, independent of how easy it is to trigger the on-chain event itself.
**Status:** mitigated — covered by an automated test (`ClaimVault.t.sol`, `test_RegisterOrder_RevertsIfCallerIsNotWorker`). Found during Sprint 2's `entry-point-analyzer` pass, not present in the original Sprint 1 skeleton's design intent (the NatSpec comment predated the actual restriction).

### T9 — Fake `DeliveryFailed` event from an untrusted Sepolia contract
**Attack:** the Attestcoin precompile proves that a decoded log's *content* (topics/data) was genuinely included and attested — it does not know or care which contract emitted it. Without an extra check, any attacker-deployed Sepolia contract could emit an event with the exact same signature and topics as `DeliveryTrackerMock.DeliveryFailed`, get it attested (it's a real, honest transaction), and pass it off as a legitimate trigger for any registered `orderId`.
**Mitigation:** `ClaimVault` stores `DeliveryTrackerMock`'s address as an immutable `sourceContract` set at deployment, and `submitClaim` only accepts a `DeliveryFailed` log whose emitter address (`log.address_`) matches it exactly — logs from any other contract are ignored even if the signature and topics match.
**Status:** mitigated by design — covered by an automated test (`test_SubmitClaim_RevertsOnEventFromUntrustedContract`).

### T10 — AI claims agent hallucinating or inventing data
**Attack:** an LLM can fabricate plausible-sounding details, answer about the wrong claim (cross-talk), or invent facts not present in its input — independent of whether its numeric suggestion happens to be reasonable. `internal/claimsagent.Agent.SuggestPayout` mitigates this with two independent layers, neither of which relies on trusting the model's own text:
1. **Grounding at generation time** — a system prompt (Gemini's `systemInstruction`, kept separate from the per-claim user prompt) explicitly forbids inventing or assuming facts not given, and the LLM is never told the claim's real monetary value in the first place — it only ever judges severity as a 0-100 percentage, narrowing what it could hallucinate about the payout amount to begin with. `generationConfig.responseMimeType: "application/json"` additionally constrains the output format itself.
2. **Verification after generation** — the LLM must echo back the exact order ID it was given as `orderIdConfirmation`. If that echo doesn't match the order actually being processed, the response is treated as a hallucination/cross-talk signal and the suggested percentage is discarded outright (not merely distrusted-but-used) in favor of the same safe full-refund fallback used for unparsable responses.
Whatever percentage survives both layers is still bounded by the off-chain `policyCap` and, independently, by `ClaimVault`'s on-chain `payoutCap` (see T4) — a hallucinated suggestion can be wrong, but it cannot authorize an out-of-bounds payout on its own.
**Status:** mitigated by design — covered by automated tests (`TestSuggestPayout_FallsBackToFullAmountWhenOrderIdConfirmationMismatches`, `TestSuggestPayout_FallsBackToFullAmountWhenOrderIdConfirmationMissing`, `TestSuggestPayout_NeverExposesTheMonetaryAmountToTheLLM`) and verified against the real Gemini API, which correctly echoed the given order ID and produced valid, unfenced JSON.

### T11 — Continuity-proof staleness when `submitClaim` is delayed after attestation

**Discovered:** Sprint 5's full-integration stress test (five purchases made within ~90 seconds of each other, all failing near-simultaneously). Reproduced again during Sprint 7's demo recording, this time with a single, isolated claim delayed roughly 15 minutes by an unrelated debugging detour.
**Symptom:** `submitClaim` reverts inside the Attestcoin precompile with `Continuity proof does not match attestation or checkpoint` — reproduced against a brand-new, freshly-refetched proof each time, ruling out a nonce/RPC race or a cached/stale proof object as the cause. Every claim submitted promptly (within a couple of minutes of attestation) across many runs this session succeeded; every claim submitted after a longer delay, whether alone or in a same-checkpoint batch, failed the same way.
**Analysis:** `docs/ATTESTCOIN_INTEGRATION.md` already notes that a continuity proof's verification cost grows with age as Attestcoin's checkpoint advances — this session's evidence across two separate incidents shows age doesn't just raise cost past some point, it invalidates the proof outright once the checkpoint it's anchored to rotates. The original Sprint 5 batch (five claims failing together) is one instance of this: they shared a checkpoint window and it rotated out from under all of them at once. The Sprint 7 incident shows the same failure needs no concurrency at all — a single claim, delayed on its own, hits the identical revert. The real trigger is elapsed time since attestation, not how many claims are in flight.
**Impact:** no funds moved and no over-payout occurred in either incident — the precompile correctly refused every stale claim rather than verifying incorrectly. The impact is availability, not fund safety: a buyer whose claim's proof goes stale before submission needs a fresh trigger (a new `DeliveryFailed`-equivalent event) to be paid, since `submitClaim`'s anti-replay key is tied to the specific proof, and a stale proof cannot be "refreshed" back to validity by refetching it — the underlying checkpoint has already moved on.
**Mitigation:** not yet implemented. `cmd/worker` already submits "as soon as attestation is available" per the architecture's stated design, which is the right default, but nothing currently bounds *how* soon is required, and any delay in the pipeline — a slow LLM call, a retry backoff, an operator debugging something else, several claims clustering — can push a claim past whatever that window turns out to be. Candidate fixes for post-hackathon hardening: measure the actual checkpoint-rotation interval empirically, add an explicit deadline/alert if a claim is at risk of missing it, or detect this specific revert and retry with a freshly-fetched proof against the *new* checkpoint rather than exhausting retries against the same stale one.
**Status:** disclosed, unresolved — documented rather than silently retried, since retrying against a proof that cannot become valid again just burns gas and delays the operator noticing.

## Negative test cases covered (Foundry)

- Submitting a claim with an invalid proof → revert
- Submitting a claim referencing an already-processed transaction (replay) → revert
- Submitting a claim for a source transaction with a failure status (`!= 0x1`) → revert
- Submitting a suggested value above the policy cap → the value is capped, not honored as-is
- Submitting a claim for a nonexistent `orderId` → revert
- Registering an order as anyone other than the authorized worker → revert
- Submitting a claim whose event's buyer doesn't match the registered order's buyer → revert
- Submitting a claim whose `DeliveryFailed` log was emitted by a contract other than the trusted `sourceContract` → revert
- Submitting a claim for an order that was already paid out → revert
- Submitting a claim the pool's balance can't cover → revert, order left unclaimed and resubmittable (fuzzed across the balance/protection-amount range)
- The payout formula `min(suggestedPayout, protectionAmount, payoutCap)` holds across the full `uint256` range of all three inputs, including zero on either bound (fuzzed)
