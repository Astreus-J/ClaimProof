# Submission Copy — ready to paste into the DoraHacks form

## Project name

ClaimProof

## Tagline

A claim that pays itself — because the proof is faster than the adjuster.

## Track

AI

## Short description (1–2 sentences, for the listing card)

A parametric insurance protocol where an AI evaluates the claim, but only authorizes payment after the Attestcoin Protocol re-verifies on-chain that the trigger event actually happened — no human adjuster, no single trusted oracle.

## Long description (for the BUIDL page)

Traditional insurance claims depend on a human adjuster or a single trusted oracle to evaluate and authorize payment. It's slow, contestable, and opaque: whoever is waiting for the refund only has the platform's word.

ClaimProof fixes this by trading trust for proof. When a delivery fails (the trigger chosen for the MVP, generalizable to other verifiable events), the event is recorded as a transaction on Ethereum Sepolia. An AI agent evaluates the order context and suggests a payout amount — but it never has the power to authorize payment alone. The `ClaimVault` contract, running on Creditcoin, only releases the amount after re-verifying on-chain, through the native Attestcoin Protocol precompile (`0x0FD2`), that the event actually occurred — checking inclusion, continuity, the source transaction's success status, and replay protection.

The result: an automatic, auditable payout that doesn't depend on the goodwill of any party — not the store, not the insurer, not even the AI itself.

None of the other projects submitted to this hackathon address parametric insurance; ClaimProof occupies that space by combining the field's most validated cross-chain verification pattern (used by 9 other submissions, in credit/settlement domains) with a new product domain and a real market (insurance is a multi-billion-dollar fintech category).

## Why Attestcoin (for the "Attestcoin Integration Summary" field)

Without Attestcoin, `ClaimVault` would have no cryptographically trustworthy way of knowing a delivery failed on another chain — the AI would have to trust unverified text or an API to authorize money, exactly the attack vector the product exists to eliminate. Full technical breakdown (SDK functions, precompile, decoding, anti-replay) in `docs/ATTESTCOIN_INTEGRATION.md`.

## Why Creditcoin

It's the chain where the Attestcoin Protocol's native precompile runs — there's no reason to reimplement the verification on another L1, and `ClaimVault` benefits directly from Creditcoin's execution and settlement infrastructure.

## Links (fill in at the end)

- GitHub: [fill in]
- Demo video: [fill in]
- Deck / whitepaper: `docs/WHITEPAPER.md` (export to PDF)
- `ClaimVault` contract (Creditcoin CC3 Testnet): [fill in address]
- `DeliveryTrackerMock` contract (Ethereum Sepolia): [fill in address]
