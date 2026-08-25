# ClaimProof — Product Vision

> Phase 3. Depends on the decision in [decision.md](decision.md).

**Tagline:** "A claim that pays itself — because the proof is faster than the adjuster."

## Problem

Insurance claims — especially in e-commerce and logistics — depend on a human adjuster or a single trusted oracle to evaluate and authorize payment. It's slow, contestable, and opaque: the buyer only has the platform's word.

## Solution

A parametric insurance protocol where an AI evaluates the claim and suggests the payout amount, but **never authorizes payment alone** — the contract only releases the amount after re-verifying on-chain, via Attestcoin, that the trigger event (delivery failure) actually happened on the source chain (Ethereum Sepolia).

## Audience & persona

Decentralized marketplaces and e-commerce stores that want to offer "delivery protection" without operating an insurance company.

**Persona — Marina:** owner of a mid-size online store, tired of chargeback disputes and slow manual refunds.

## User journey

1. The buyer completes checkout and pays an optional protection premium.
2. The order is registered in a tracking contract with a delivery SLA (Ethereum Sepolia).
3. If the SLA expires without a delivery confirmation, the `DeliveryFailed` event is emitted.
4. The off-chain worker detects the event, obtains the inclusion proof via Attestcoin, and asks the AI Claims Agent to suggest the payout amount.
5. The `ClaimVault` contract on Creditcoin re-verifies the proof through the Attestcoin precompile before releasing payment.
6. Automatic payout to the buyer's wallet, in seconds, with no human intervention.

## Core feature

Automatic payout locked behind cryptographic proof — payment authorization is a verifiable event, not a trust decision.

## Secondary features

- Real-time claims dashboard
- History of verified events
- Underwriting pool with multiple capital providers
- Per-store claim-rate panel

## Differential & moat

None of the hackathon's 14 competitors touch insurance (see [decision.md](decision.md) for the full comparison). The history of cryptographically verified claims becomes a proprietary data asset for pricing risk better over time — and "proof, not promise" is hard to replicate in a traditional insurance product.

## Business model

A fee on the protection premium paid by the buyer (10–20% initial overhead) and, on the post-hackathon roadmap, a spread on the underwriting pool as the product matures.

## Scalability

The same core architecture (verifiable event → AI suggests value → Attestcoin authorizes) extends to other triggers without redesigning the core: flight delay, DePIN sensor failure, RWA custody event.

## Future roadmap (post-hackathon)

- Integrate real logistics oracles (replacing `DeliveryTrackerMock`)
- Expand to additional source chains as Attestcoin adds support
- Pursue partnerships with real marketplaces
- Apply to the CEIP fast-track available to the hackathon's top 3 finishers
