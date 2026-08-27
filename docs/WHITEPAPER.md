# ClaimProof — Whitepaper

*Source for the deck/whitepaper PDF required by the submission.*

> **Exported:** [`ClaimProof-Whitepaper.pdf`](ClaimProof-Whitepaper.pdf) — this content plus the architecture excerpt (sequence diagram + components table) from [../ARCHITECTURE.md](../ARCHITECTURE.md). Source HTML at [`whitepaper-export.html`](whitepaper-export.html); regenerate by rendering it to PDF (e.g., open in a browser and print to PDF, or headless Chromium) if this file's content changes.

---

## 1. Summary

ClaimProof is a parametric insurance protocol where payment authorization never depends on a single actor's word — not the insurer, not the adjuster, not even the AI that evaluates the claim. Payment only happens after the trigger event is re-verified on-chain, through the Attestcoin Protocol, against the chain where it actually occurred.

## 2. The problem

Traditional insurance claims — and, in particular, delivery guarantees in e-commerce — depend on slow, opaque dispute processes: the buyer complains, the platform investigates, a human decides. This creates friction, operational cost, and distrust when the parties have conflicting interests (the platform doesn't want to pay; the buyer wants a guarantee they'll be made whole).

## 3. The insight

If the event that should trigger a claim already exists as a verifiable transaction on some chain (a delivery confirmation, a recorded failure, a tracking event), you don't need an adjuster to decide "did this happen?" — you only need to prove it.

## 4. The solution

ClaimProof separates two decisions that traditional insurance systems conflate:

1. **"Did this happen?"** — answered exclusively by cryptographic proof via the Attestcoin Protocol, never by human or AI opinion.
2. **"How much to pay?"** — answered by an AI agent, but always bounded by an on-chain policy cap, never decided freely.

This separation is what makes the system resistant both to buyer fraud (no claim without real proof) and to manipulation of the AI layer (no authorization beyond what policy allows).

## 5. How it works (product view)

1. The buyer purchases delivery protection along with the order.
2. If the delivery fails (an event recorded on-chain on Ethereum Sepolia), the system detects it automatically.
3. The AI evaluates the appropriate refund amount.
4. The contract on Creditcoin re-verifies the event's proof via Attestcoin before releasing any amount.
5. The buyer receives the payout automatically, with no support ticket required.

## 6. Why blockchain, why Creditcoin, why Attestcoin

- **Why blockchain:** eliminates the information asymmetry between the insured and the insurer — the proof is public and verifiable by any party, not just the insurer.
- **Why Creditcoin:** it's the chain where the Attestcoin Protocol's native precompile runs; hosting the business logic there avoids reimplementing cross-chain verification on another L1.
- **Why Attestcoin:** it's the only source of truth accepted to authorize a payment. Remove Attestcoin, and the system loses the ability to reliably know whether the trigger event actually occurred — the AI would have to blindly trust unverified data.

## 7. Competitive differential

No known project in the Creditcoin/Attestcoin ecosystem addresses parametric insurance as of this submission. Most Attestcoin integrations explore credit/settlement (loans, collateral, escrow release) — ClaimProof applies the same verification primitive to a new product domain, with a straightforward business model (fee on the protection premium) and an addressable market far larger than developer-facing DeFi infrastructure.

## 8. Business model

- **Short term:** a fee on the protection premium paid by the buyer (10–20% overhead).
- **Medium term:** a spread on the underwriting pool as more capital providers participate.
- **Long term:** cryptographically verified claims data becomes a proprietary risk-pricing asset — the more volume, the better the pricing, a competitive advantage that's hard to replicate without the same verified history.

## 9. Roadmap

- **Hackathon (MVP):** e-commerce delivery-failure trigger, testnet, AI bounded by policy.
- **Post-hackathon, short term:** replace `DeliveryTrackerMock` with a real logistics provider integration; add a second trigger type (e.g., flight delay).
- **Medium term:** underwriting pool open to multiple capital providers; expansion to additional source chains as the Attestcoin Protocol adds support.
- **CEIP application:** pursue the investment fast-track offered to the hackathon's top three finishers to accelerate this expansion.

## 10. Team

See [TEAM.md](TEAM.md).
