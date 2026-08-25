# BUIDL CTC 2026 Fall — Ideation and Project Selection

> Reference document for Phase 2 (Brainstorm → Decision Matrix → Selection → Red Team).
> Full factual background (rules, deadlines, Attestcoin deep dive, landscape of the 14 competing submissions) lives in Phase 1 — see the "BUIDL CTC Intelligence" artifact.
> Direction set by the user: **prioritize the niche with the highest real probability of winning**, not the "most creative" idea in isolation.

---

## 1. Context that shaped the brainstorm

From the Phase 1 research, three concrete facts determined how the 30 ideas were allocated across tracks:

- **Gaming: 0 real competitors** among the 14 mapped submissions.
- **DePIN: 1 competitor**, and a weak one (recycled pitch from another hackathon, Attestcoin mentioned as an afterthought).
- **DeFi (5 submissions) and RWA (4 submissions) hold 9 of the 14 submissions**, including the three technically strongest projects in the entire field (`index41`, `crosscredit`, `Cr3dX`) — all following the same pattern: "prove an event on Ethereum/Sepolia → release credit/escrow on Creditcoin" via the BlockProver precompile.
- **AI: 2 competitors** (`Oracle-Free Council`, `BountyOps Verified Execution`), each with a different product thesis — the track still has room.

Allocation of the 30 ideas: **Gaming 8 · DePIN 7 · AI 7 · DeFi 4 · RWA 4** — heavy weight on open tracks, light coverage on saturated ones.

"Core vs. decorative" test applied to every idea (drawn from the competitor `Oracle-Free Council` itself): *"remove Attestcoin — does the project lose its inputs or its ability to act?"* — if the answer is no, the idea is decorative and was discarded or downgraded on the Attestcoin Fit score.

---

## 2. The 30 ideas

Score 0–10 on: **Orig** (originality) · **Imp** (impact) · **Att** (Attestcoin usage) · **Viab** (feasibility in 19 days) · **Biz** (business potential) · **Demo** (demo quality) · **Judge** (judge-impressing potential). **Win** = synthesized average of these seven scores.

### Gaming (8) — track with no real competition

| # | Idea | Problem / Solution in one line | Verified event | Orig | Imp | Att | Viab | Biz | Demo | Judge | **Win** |
|---|---|---|---|---|---|---|---|---|---|---|---|
| G-01 | **CrossMint Foundry** | Duplicated items across chains break game economies → mint on Creditcoin only after a verified burn on Ethereum (burn-and-mint, no custodian) | `Transfer/Burn` of the source NFT | 6 | 6 | 9 | 8 | 6 | 8 | 7 | **7.5** |
| G-02 | Guild Ledger | Player reputation locked inside a single game/chain → cross-game reputation passport | Tournament achievement/prize tx | 7 | 5 | 6 | 6 | 5 | 6 | 6 | 6.0 |
| G-03 | Provenance Arena | PvP wagering with an NFT stake suffers double-spend across chains → the match only starts after verifying custody of the staked NFT | Transfer-to-escrow tx | 5 | 5 | 7 | 5 | 5 | 7 | 6 | 5.5 |
| G-04 | Loot Bridge | Cross-chain loot boxes rely on a custodial bridge today → burn unlocks a verifiable opening + mint of the contents | Loot box burn tx | 6 | 5 | 9 | 8 | 5 | 8 | 6 | 7.1 |
| G-05 | **Tournament Escrow** | Esports prize payouts depend on a trusted organizer → the escrowed prize only releases after proof of a result recorded in a scoreboard contract | Final result tx from the scoreboard | 7 | 6 | 8 | 7 | 6 | 8 | 8 | **7.5** |
| G-06 | Season Pass Relay | Season passes don't reward progress made in a partner game/chain → cross-game milestones unlock tiers | Milestone tx in the partner game | 5 | 4 | 6 | 6 | 5 | 5 | 5 | 5.0 |
| G-07 | Anti-Cheat Ledger | Bots exploit game economies → high-value economic actions require a prior "verified human" attestation | Humanity-verification tx | 7 | 6 | 5 | 4 | 5 | 5 | 5 | 5.0 |
| G-08 | Cross-Game Currency Vault | Currency earned on Ethereum has no trustless exit → verified burn unlocks 1:1 minting on Creditcoin | Currency token burn tx | 6 | 5 | 8 | 7 | 6 | 7 | 6 | 6.8 |

### DePIN (7) — 1 weak competitor

| # | Idea | Problem / Solution in one line | Verified event | Orig | Imp | Att | Viab | Biz | Demo | Judge | **Win** |
|---|---|---|---|---|---|---|---|---|---|---|---|
| D-01 | SensorProof Grid | Sensor networks rely on a central aggregator → a reading posted on Ethereum unlocks the incentive | Sensor-reading tx | 6 | 7 | 7 | 6 | 7 | 6 | 6 | 6.5 |
| D-02 | Compute Attest | GPU/compute marketplaces need a trusted central operator → proof of job completion unlocks payment | "Job completed" tx | 6 | 7 | 8 | 6 | 7 | 6 | 7 | 7.1 |
| D-03 | Bandwidth Relay Pay | Shared-bandwidth networks (Helium-style) need trustless settlement | Node usage-report tx | 5 | 5 | 6 | 5 | 5 | 5 | 5 | 5.0 |
| D-04 | **EV Charge Verify** | EV charging networks lack automatic, auditable settlement → completing a session unlocks a dual payment | "Session completed" tx w/ kWh delivered | 7 | 7 | 7 | 6 | 7 | 7 | 7 | **7.25** |
| D-05 | Solar Yield Ledger | Renewable generation needs proof of production for carbon credits — *risk: name/niche overlaps a weak existing competitor ("Solar DePin")* | Energy-generation tx | 4 | 6 | 6 | 6 | 6 | 6 | 5 | 5.5 |
| D-06 | Storage Proof Bridge | Decentralized storage needs settlement via storage proofs | Storage-proof tx | 5 | 6 | 7 | 5 | 6 | 5 | 5 | 5.7 |
| D-07 | **Fleet Telemetry Trust** | Logistics depends on the carrier's word → delivery confirmation unlocks payment **and** triggers parametric insurance | Delivery confirmation/failure tx | 7 | 7 | 8 | 6 | 7 | 7 | 7 | **7.25** |

### AI (7) — 2 competitors

| # | Idea | Problem / Solution in one line | Verified event | Orig | Imp | Att | Viab | Biz | Demo | Judge | **Win** |
|---|---|---|---|---|---|---|---|---|---|---|---|
| A-01 | Verified-Fact Trading Agent | AI agents are vulnerable to fake data/prompt injection → they only act on re-verified events | Relevant on-chain transfer tx | 7 | 7 | 8 | 6 | 7 | 6 | 7 | 7.0 |
| A-02 | Compliance Copilot | Regulated DeFi reimplements compliance per chain → an agent grants access after verifying a compliance tx on another chain | Compliance-check tx | 7 | 8 | 7 | 5 | 8 | 5 | 6 | 6.5 |
| A-03 | Audit Agent | An AI analyzing contracts can hallucinate facts → the report only cites state backed by an Attestcoin proof | Audited cross-chain state tx | 6 | 6 | 7 | 6 | 6 | 6 | 6 | 6.0 |
| A-04 | **★ Insurance Claims Adjuster AI** | Parametric claims depend on a human adjuster/single oracle → an AI evaluates, but only pays after on-chain re-verification of the trigger event | Trigger-event tx (delay, delivery failure, sensor reading) | 7 | 8 | 8 | 6 | 8 | 8 | 8 | **7.6** |
| A-05 | Reputation Scoring Agent | Undercollateralized lending needs a cross-chain score — *risk: close to the saturated credit pattern* | Repayment tx across multiple protocols | 6 | 7 | 7 | 6 | 7 | 6 | 6 | 6.2 |
| A-06 | AI Underwriter | Risk pricing uses loose data, not verifiable history | Portfolio of verified events | 6 | 6 | 6 | 5 | 6 | 5 | 5 | 5.6 |
| A-07 | Grant Quality Evaluator Agent | DAOs pay bounties on a binary basis only — *risk: overlaps the "BountyOps" competitor* | Contribution merge tx | 6 | 6 | 6 | 6 | 6 | 6 | 6 | 6.0 |

### DeFi (4) — light coverage, track saturated by the field's 3 strongest technical entries

| # | Idea | Problem / Solution in one line | Orig | Imp | Att | Viab | Biz | Demo | Judge | **Win** |
|---|---|---|---|---|---|---|---|---|---|---|
| F-01 | Cross-Chain Collateral Aggregator | Lowers collateral by aggregating verified history across multiple Ethereum protocols — *risk: integration complexity too high for 19 days* | 6 | 7 | 7 | 5 | 7 | 5 | 6 | 5.8 |
| F-02 | Governance-Gated Treasury | A DAO treasury only spends after verifying a `ProposalExecuted` event from a vote on another chain | 7 | 6 | 8 | 6 | 6 | 6 | 6 | 6.4 |
| F-03 | Airdrop Anti-Sybil Vault | Claims use Attestcoin's native anti-replay protection against sybil/double claims | 6 | 5 | 7 | 7 | 5 | 6 | 5 | 5.9 |
| F-04 | Perpetuals Funding Bridge | Perp funding-rate settlement depends on a verified price-execution tx from a reference DEX on Ethereum | 5 | 6 | 6 | 5 | 6 | 5 | 5 | 5.2 |

### RWA (4) — light coverage, track with 4 competitors

| # | Idea | Problem / Solution in one line | Orig | Imp | Att | Viab | Biz | Demo | Judge | **Win** |
|---|---|---|---|---|---|---|---|---|---|---|
| R-01 | **Reserve-Proof Stablecoin** | An asset-backed token only mints/burns after a proof of reserve/redemption from the custodian on Ethereum | 7 | 8 | **9** | 6 | 8 | 7 | 7 | **7.25** |
| R-02 | Real Estate Escrow Verify | *Risk: property registries are rarely on Ethereum — the event is hard to justify* | 6 | 7 | 6 | 4 | 6 | 5 | 5 | 5.4 |
| R-03 | Commodity Custody Ledger | Tokenized commodities only move after a verified custody attestation | 6 | 7 | 8 | 6 | 7 | 6 | 6 | 6.4 |
| R-04 | Trade Finance Verifier | Trade finance releases invoice payment after proof of customs clearance | 6 | 7 | 7 | 5 | 7 | 6 | 6 | 6.1 |

**Discarded for score < 6.0 or explicit feasibility/overlap risk:** Guild Ledger, Provenance Arena, Season Pass Relay, Anti-Cheat Ledger, Bandwidth Relay Pay, Solar Yield Ledger, Storage Proof Bridge, Audit Agent, Reputation Scoring Agent, AI Underwriter, Grant Quality Evaluator Agent, Cross-Chain Collateral Aggregator, Airdrop Anti-Sybil Vault, Perpetuals Funding Bridge, Real Estate Escrow Verify.

---

## 3. Top 10 → Top 5 → Top 3

**Top 10** (by Win score): Insurance Claims Adjuster AI (7.6) · CrossMint Foundry (7.5) · Tournament Escrow (7.5) · EV Charge Verify (7.25) · Fleet Telemetry Trust (7.25) · Reserve-Proof Stablecoin (7.25) · Compute Attest (7.1) · Loot Bridge (7.1) · Verified-Fact Trading Agent (7.0) · Cross-Game Currency Vault (6.8).

**Top 5** (removing conceptual overlap — Loot Bridge duplicates CrossMint Foundry, Cross-Game Currency Vault duplicates the same burn-and-mint pattern):
1. Insurance Claims Adjuster AI
2. Fleet Telemetry Trust
3. Tournament Escrow
4. Reserve-Proof Stablecoin
5. CrossMint Foundry

**Top 3:**
1. Insurance Claims Adjuster AI
2. Tournament Escrow
3. Fleet Telemetry Trust

---

The final decision, the full reasoning behind it, the differential against each of the other 29 ideas, and the differential against the 14 already-submitted competitors are documented separately in [decision.md](decision.md).
