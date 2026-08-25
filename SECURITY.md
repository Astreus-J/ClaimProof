# Security — ClaimProof

## General posture

ClaimProof handles automatic insurance-claim payments. The core design principle is: **no single actor — not the AI, not the off-chain worker, not a single operator — can, on its own, authorize funds to move.** The only valid authorization is an on-chain proof re-verified by the Attestcoin Protocol.

## Attack surface

| Component | Risk | Mitigation |
|---|---|---|
| `DeliveryTrackerMock.sol` (Sepolia) | For the demo, it's team-controlled — in production it would need to be replaced by a real logistics oracle with its own set of guarantees | Explicitly documented as a mock; the post-hackathon roadmap lists replacing it with a real oracle |
| AI Claims Agent | Could, in theory, suggest a manipulated or out-of-policy payout value | The contract enforces an **on-chain policy cap** independent of the AI's suggestion — the AI never has signing power or final authority |
| Off-chain worker | A compromised worker private key could submit fake claims | The worker can only submit a claim that passes Attestcoin precompile re-verification — a claim without valid proof of a real event is rejected by the contract, regardless of who submitted it |
| Claim replay | Resending the same proof to claim multiple payouts | `processedQueries` mapping (chainKey, blockHeight, transactionIndex) in `ClaimVault`, inherited from the Attestcoin Protocol's reference pattern |
| Inclusion proof without checking source tx success | The precompile only proves **inclusion**, not **success** — a common integration mistake | `ClaimVault` explicitly checks the `0x1` status of the decoded transaction before any action |
| Permissioned attestor set | The Attestcoin Protocol operates today in `AuthorizedOnly` mode — the attestor set isn't fully decentralized yet | Risk inherited from the protocol, outside the project's control; documented for transparency with judges |
| `MinBondRequirement` of 0 CTC on mainnet | Economic security against attestor collusion isn't active yet | Same note above — an infrastructure risk, not ClaimProof's |

## Testnet scope

The entire system is deployed and demonstrated **exclusively on testnet** (Ethereum Sepolia + Creditcoin CC3 Testnet), as required by the hackathon submission. No real funds are at risk.

## What's not covered in this MVP

- Third-party audit of the contracts (out of scope for the hackathon timeline; the CertiK audit credits offered by the hackathon are a natural next step post-submission)
- Front-running resistance on claim submission (mitigable, but out of MVP scope)
- A compliance/KYC model for high-value payouts

## Responsible disclosure

This is a hackathon project on testnet. For security questions about this repository, open a GitHub issue tagged `security`.

## Full threat model

See [docs/THREAT_MODEL.md](docs/THREAT_MODEL.md) for the detailed analysis of attack vectors and mitigations.
