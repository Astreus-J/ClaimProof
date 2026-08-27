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

## USC Integration Summary (for the form field "USC Integration Summary (explain how your project uses USC)" — USC / Attestcoin Protocol)

ClaimProof uses the USC/Attestcoin Protocol exclusively in its **read** direction: proving that an event on one chain (Ethereum Sepolia) really happened, so it can be trusted on another (Creditcoin CC3 Testnet). Without it, our AI claims agent would have no cryptographically trustworthy way to know a delivery failed — it would have to trust unverified text or an API, exactly the failure mode this product exists to eliminate.

**Off-chain (Go worker, `internal/proofbuilder`):** there is no official Go SDK for Attestcoin, so we talk directly to the hosted Prover REST API (`https://prover.cc3-testnet.creditcoin.network`), reimplementing the same calls the official TypeScript SDK (`@gluwa/usc-sdk`) wraps: `GET /api/v1/attested-height/{chainKey}` (polled until the target Sepolia block is attested) and `GET /api/v1/proof-by-tx/{chainKey}/{transactionHash}` (Merkle inclusion proof + continuity proof for the `DeliveryFailed` transaction).

**On-chain (`ClaimVault.sol`, Creditcoin):** `submitClaim` calls `INativeQueryVerifier.verifyAndEmit()` on the native `0x0FD2` precompile inside the same transaction as the payout — if verification fails (invalid proof, unattested block, broken continuity), the whole transaction reverts and no funds move. On success, we decode the transaction via `EvmV1Decoder` (from `@gluwa/usc-contracts`), explicitly check the receipt's success status (the precompile proves inclusion, not success — a common integration mistake we test against), verify the log came from our trusted `DeliveryTrackerMock` address, and check an anti-replay mapping keyed by `(chainKey, blockHeight, transactionIndex)` before releasing payout.

**Result:** the AI claims agent may only ever suggest a payout amount, capped by an on-chain policy — USC/Attestcoin's on-chain re-verification is the only thing that can authorize the transfer. Full technical breakdown, request/response schemas, and known protocol limitations we designed around (e.g., proof-age effects on verification) are documented in `docs/ATTESTCOIN_INTEGRATION.md` in the repository.

## Why Creditcoin

It's the chain where the Attestcoin Protocol's native precompile runs — there's no reason to reimplement the verification on another L1, and `ClaimVault` benefits directly from Creditcoin's execution and settlement infrastructure.

## Links (fill in GitHub + demo video at the end)

- GitHub: [fill in]
- Demo video: [fill in]
- Deck / whitepaper: `docs/ClaimProof-Whitepaper.pdf`
- `ClaimVault` contract (Creditcoin CC3 Testnet): `0xd6f0680F366d2de5849ab00Ff2Ca48aa1D030bCd`
- `DeliveryTrackerMock` contract (Ethereum Sepolia): `0x5c293e0C72E52fAca66befbEd2a65552431Ce46d`
