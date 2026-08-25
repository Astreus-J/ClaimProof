# ClaimProof — Roadmap, Demo, Pitch & Submission

> Phase 3. Complements [product.md](product.md) and [architecture.md](architecture.md).

## Roadmap through 2026-09-13

19 days starting 2026-08-25. Phases 1 and 2 (research and idea selection) are already complete.

| Date | Phase |
|---|---|
| Aug 26–28 | Project setup, base contracts (`DeliveryTrackerMock` + `ClaimVault` skeleton), Foundry environment, initial integration with the Attestcoin SDK on testnet |
| Aug 29–Sep 1 | Full verification flow (precompile, `EvmV1Decoder`, `processedQueries`), unit and negative-path tests |
| Sep 2–4 | Off-chain worker + integration with the AI Claims Agent |
| Sep 5–7 | Frontend: purchase flow + claims dashboard |
| Sep 8–9 | Full E2E integration on testnet, final deploy, bug fixing |
| Sep 10–11 | Final tests, documented threat model, README, Attestcoin technical doc |
| Sep 12 | Demo video recording, pitch deck |
| Sep 13 | Buffer for surprises + final submission before 23:59 ET |

## Demo script (3–5 min)

| Time | Content |
|---|---|
| 0:00–0:20 | **Hook** — "What if your money came back on its own the second a delivery fails — no ticket, no adjuster, no waiting days?" |
| 0:20–0:50 | **Problem** — claims depend on a trusted third party's word: slow, contestable, opaque |
| 0:50–1:20 | **Solution** — ClaimProof: payout locked by proof via Attestcoin, AI only decides the amount |
| 1:20–3:30 | **Live demo** — buy protection → simulate a delivery failure on Sepolia → proof generated and re-verified on-chain on Creditcoin → instant payout in the wallet |
| 3:30–4:00 | **Attestcoin / Creditcoin** — show the call to the `0x0FD2` precompile and each chain's role |
| 4:00–4:30 | **Impact** — marketplaces offer delivery guarantees without becoming an insurer; buyers get verifiable trust |
| 4:30–5:00 | **Future / business / closing** — expansion to other triggers, revenue model, CEIP fast-track ask |

## Pitch (60 seconds)

- **First line:** "ClaimProof is insurance that pays for itself, because the proof arrives before the adjuster does."
- **Problem:** traditional claims depend on trusting a third party — slow, contestable, opaque.
- **Insight:** if the event that triggers a claim already exists as a verifiable transaction on another chain, you don't need an adjuster — you need proof.
- **Solution:** an AI evaluates and suggests the amount; the contract only pays after re-verifying on-chain, via Attestcoin, that the event happened.
- **Differential:** none of the 14 already-submitted projects touch insurance, and none combine AI-driven evaluation with payout authorization locked by cryptographic proof.
- **Attestcoin:** the only source of truth the contract accepts — without it, the AI would have to trust unverified text to authorize money.
- **Creditcoin:** the execution and payment chain, where the Attestcoin precompile runs natively.
- **Impact & business:** any marketplace can offer a guarantee without becoming an insurer; revenue comes from the fee on the protection premium.
- **Future:** the same architecture scales to flights, DePIN, and RWA — anywhere a verifiable event can serve as a trigger.

## Submission checklist

| Item | How to execute it better than the competition |
|---|---|
| Name, logo, short/long description | "ClaimProof" + tagline; short description focused on the verifiable trigger, long description detailing the E2E flow |
| Track | AI — maps directly onto the track's official description |
| Attestcoin Integration Summary | Spell out exactly which SDK/precompile functions are called, mirroring the transparency standard set by the competitor `index41` |
| GitHub + README | Public repository, README that lets someone run the full flow in minutes |
| Demo video | Follow the script above, showing the real flow, not slides |
| Deck / whitepaper PDF | Condensed Product Vision + Architecture, with the sequence diagram |
| Documented architecture | Sequence diagram + end-to-end transaction flow explanation |
| Team data | Name, email, country, role, bio for each member |
| Testnet deployment | `DeliveryTrackerMock` on Sepolia + `ClaimVault` on Creditcoin CC3 testnet, addresses published in the README |
| Contracts + tests | Full source code with the Foundry suite visible in the repository |

## GitHub structure

```text
claimproof/
├── contracts/
│   ├── src/
│   │   ├── ClaimVault.sol
│   │   ├── DeliveryTrackerMock.sol
│   │   └── lib/EvmV1Decoder.sol
│   ├── script/
│   │   └── Deploy.s.sol                # Foundry deploy script
│   └── test/
│       ├── ClaimVault.t.sol
│       └── ReplayProtection.t.sol      # adversarial cases
├── backend/                            # off-chain worker (Go)
│   ├── cmd/worker/main.go
│   └── internal/
│       ├── listener/
│       ├── proofbuilder/               # Attestcoin Prover HTTP client
│       ├── claimsagent/
│       └── chain/                      # go-ethereum client + abigen bindings
├── frontend/
│   ├── app/store/                      # purchase flow
│   └── app/dashboard/                  # live claim status
├── scripts/
│   └── deployments.json                # deployed contract addresses (data only)
├── docs/
│   ├── ATTESTCOIN_INTEGRATION.md       # required by the submission
│   └── THREAT_MODEL.md
├── demo/
│   └── demo-video-link.md
├── README.md
├── ARCHITECTURE.md
└── SECURITY.md
```
