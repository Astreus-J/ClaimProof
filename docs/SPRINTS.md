# Sprint-by-Sprint Development Plan — ClaimProof

> 19 calendar days, from 2026-08-26 through the submission deadline on 2026-09-13, 23:59 ET. Short sprints (2–4 days) given the hackathon timeline. Every sprint has a goal, tasks, and an exit criterion — don't move to the next sprint without meeting the exit criterion, unless there's an explicit scope-cut decision (see [architecture.md](architecture.md) for what can be cut without compromising the core: the "DO NOT DO" section).

References: [product.md](product.md) (vision) · [architecture.md](architecture.md) (architecture + MVP) · [../ARCHITECTURE.md](../ARCHITECTURE.md) (submission-facing technical doc) · [execution.md](execution.md) (demo/pitch/submission).

---

## Sprint 1 — Setup & Base Contracts

**Aug 26–28, 2026 (3 days)**

**Goal:** environment ready, skeleton contracts compiling and deployed on testnet, first successful interaction with the Attestcoin SDK.

### Tasks

- [x] Create a public GitHub repository with the folder structure defined in the README
- [x] Configure Foundry (`forge init`, `foundry.toml` pointing to Creditcoin CC3 Testnet and Ethereum Sepolia via RPC)
- [x] Get testnet CTC from the faucet and configure a deploy wallet (separate from any personal wallet)
- [x] Run the official `@gluwa/usc-sdk` "Hello Bridge" tutorial (`gluwa/attestcoin-protocol-examples`) once, in a scratch directory outside the repo, purely as a reference to confirm the Prover API's request/response shape and testnet setup — it is **not** a project dependency, since the backend is Go. Ran end-to-end against live testnet (burn on Sepolia → attested → proof generated → minted on Creditcoin), confirming the full pipeline works; the confirmed Prover API schema is recorded in [ATTESTCOIN_INTEGRATION.md](ATTESTCOIN_INTEGRATION.md).
- [x] Initialize the Go module (`go mod init github.com/<org>/claimproof/backend` or equivalent), create `cmd/worker/main.go` and the `internal/{listener,proofbuilder,claimsagent,chain}/` package skeletons, add `go-ethereum` as a dependency
- [x] Write the `DeliveryTrackerMock.sol` skeleton (Sepolia): shipment struct, `createShipment` function, `DeliveryFailed` event
- [x] Write the `ClaimVault.sol` skeleton (Creditcoin): `registerOrder`, pool structure, basic access guard
- [x] Deploy both skeletons to testnet via `contracts/script/` (`DeployDeliveryTrackerMock.s.sol` on Sepolia via `forge script`; `DeployClaimVault.s.sol` on Creditcoin CC3 via `forge create`, since Creditcoin's RPC omits `mixHash` and breaks `forge script`'s local simulation — documented in the script file)
- [x] Create `scripts/deployments.json` with the deployed addresses
- [x] Generate Go contract bindings for both skeletons with `abigen` into `backend/internal/chain` — required installing a separate stable Go 1.24.6 toolchain alongside the system's non-standard 1.27 build to get `abigen`/`golangci-lint` working; bindings generated from the compiled ABI/bytecode in `contracts/out/`

### Exit criterion

Both skeleton contracts are deployed on testnet, the official SDK tutorial runs without error locally, and the repository is published with the correct structure.

---

## Sprint 2 — Full Verification Flow (technical core)

**Aug 29–Sep 1, 2026 (4 days)**

**Goal:** `ClaimVault` verifies a real `DeliveryFailed` event end-to-end via Attestcoin, with tests covering the negative cases. This is the most critical sprint — it's what separates the project from a mock.

### Tasks

- [x] Implement the call to `INativeQueryVerifier.verifyAndEmit()` on the `0x0FD2` precompile inside `ClaimVault.submitClaim(...)`
- [x] Integrate `EvmV1Decoder` to extract `orderId`, `buyer`, `timestamp`, and the status field from the decoded transaction
- [x] Implement an explicit check of the `0x1` status (the precompile only proves inclusion, not success — a documented protocol gotcha)
- [x] Implement the `processedQueries` mapping for anti-replay protection
- [x] Write `ClaimVault.t.sol`: successful verification, invalid proof rejected, nonexistent orderId rejected
- [x] Write `ReplayProtection.t.sol`: resubmitting the same proof is rejected
- [x] Manually test end-to-end on testnet: emit a real `DeliveryFailed` on Sepolia → generate the proof via `ProofBuilder` → submit it to `ClaimVault` → confirm the payout

### Exit criterion

A real event emitted on Sepolia is verified and results in a correct payout on `ClaimVault` on Creditcoin, with no mocked part in the verification. Negative tests (T1–T3 from [THREAT_MODEL.md](THREAT_MODEL.md)) pass.

**Verified 2026-08-26:** order 42 registered on `ClaimVault` (`0xE05C7771921368e3d433accCa93e9E185acb12D3`), a real `DeliveryFailed` emitted on `DeliveryTrackerMock` (Sepolia tx `0x7033d1940b59e974c4c73900cf4503642069c2d4157ec1b681215d2a2a817908`), proof fetched from the live Prover API after attestation, submitted via `submitClaim` — payout landed for the exact registered `protectionAmount`, `orders[42].claimed == true`, pool balance decreased accordingly. All 32 Foundry tests pass, including T1/T2/T3 and two additional threats found during this sprint's `entry-point-analyzer`/`guidelines-advisor` pass (T8: `registerOrder` was missing its intended `onlyWorker` gate; T9: a `DeliveryFailed` log's emitter address wasn't checked against the trusted `DeliveryTrackerMock`) — both fixed and documented in [THREAT_MODEL.md](THREAT_MODEL.md).

---

## Sprint 3 — Off-chain Worker + AI Claims Agent

**Sep 2–4, 2026 (3 days)**

**Goal:** the process runs automatically, from the event to the claim submission, with no manual intervention.

### Tasks

- [ ] Implement `internal/listener` — listens for `DeliveryFailed` via Sepolia's WSS RPC using `go-ethereum`'s `ethclient`, publishes events on a Go channel
- [ ] Implement `internal/proofbuilder` — hand-written HTTP client wrapping `WaitUntilHeightAttested()` + `GetProof()` against the Attestcoin Prover REST API
- [ ] Implement `internal/claimsagent` — LLM call with order context, returns a suggested value within a configurable policy (LLM client behind an interface, so it can be mocked in tests)
- [ ] Implement `internal/chain` — automatic submission to `ClaimVault` (`submitClaim`), signed by the worker's wallet, using the `abigen`-generated binding
- [ ] Wire everything in `cmd/worker/main.go`: read config from environment variables, construct each component via its constructor, run the listener loop, handle SIGINT/SIGTERM for graceful shutdown
- [ ] Add structured logging (`log/slog`) and retry-with-backoff on RPC/HTTP failure
- [ ] Run the worker end-to-end locally against testnet, with no manual step between the failure and the payout

### Exit criterion

When a delivery failure is simulated, the worker detects, proves, consults the AI, and submits the claim automatically — with no intermediate manual command.

---

## Sprint 4 — Frontend

**Sep 5–7, 2026 (3 days)**

**Goal:** a navigable interface covering protection purchase and live claim tracking — the part judges actually see.

### Tasks

- [ ] Set up Next.js + wagmi/ethers + WalletConnect
- [ ] "Store" screen: mock product + protection checkbox + buy button (calls `registerOrder` + `createShipment`)
- [ ] "Dashboard" screen: the connected user's orders with live status (`Active → Failure detected → Verifying proof → Paid`)
- [ ] Delivery-failure simulation button (calls the test function on `DeliveryTrackerMock`) — used when recording the demo
- [ ] Subscribe to on-chain events so status updates without a manual reload
- [ ] Handle loading and error states explicitly (no broken screen during the live demo)

### Exit criterion

Without touching code or the console, it's possible to buy protection, simulate the failure, and watch the status change to "Paid" live in the interface.

---

## Sprint 5 — E2E Integration & Final Deploy

**Sep 8–9, 2026 (2 days)**

**Goal:** the complete system, with no mocked parts outside `DeliveryTrackerMock` (documented as a deliberate simplification), ready for recording.

### Tasks

- [ ] Run the full flow (purchase → failure → proof → AI → payout) 5 times in a row without failure
- [ ] Fix bugs found during full integration
- [ ] Final "frozen" deploy of the contracts (the version used for the demo and the submission)
- [ ] Update `scripts/deployments.json` with the final addresses
- [ ] Create a GitHub release tag for the frozen version

### Exit criterion

Five consecutive runs of the full flow with no manual intervention beyond the expected UI clicks. Final addresses published.

---

## Sprint 6 — Technical Rigor & Final Documentation

**Sep 10–11, 2026 (2 days)**

**Goal:** a rigor level comparable to the field's technically strongest competitors (`index41`, `crosscredit`, `Cr3dX`).

### Tasks

- [ ] Expand the Foundry test suite (fuzzing on payout values and the policy cap)
- [ ] Review and finalize [THREAT_MODEL.md](THREAT_MODEL.md) with any findings from previous sprints
- [ ] Review [ATTESTCOIN_INTEGRATION.md](ATTESTCOIN_INTEGRATION.md) with the final real addresses and behavior
- [ ] Review [../SECURITY.md](../SECURITY.md)
- [ ] Finalize [../README.md](../README.md) with setup instructions tested from scratch (ideally by someone outside the team)
- [ ] Review [../ARCHITECTURE.md](../ARCHITECTURE.md) against the final implemented flow, adjusting the diagram if anything changed in practice

### Exit criterion

Someone outside the project can clone the repository and run the full flow following only the README, with no help.

---

## Sprint 7 — Demo & Pitch

**Sep 12, 2026 (1 day)**

**Goal:** the video and pitch materials are ready.

### Tasks

- [ ] Record the demo script (see [execution.md](execution.md)) in a real testnet environment, not simulated
- [ ] Edit the video (3–5 min) and publish it (upload/link)
- [ ] Export [WHITEPAPER.md](WHITEPAPER.md) + the architecture excerpt to PDF (the submission's deck/whitepaper)
- [ ] Rehearse the 60-second pitch (see [execution.md](execution.md))

### Exit criterion

The video is published and watchable end-to-end; the whitepaper/deck PDF is ready to attach.

---

## Sprint 8 — Buffer & Submission

**Sep 13, 2026 (1 day, deadline 23:59 ET)**

**Goal:** the submission is sent with margin before the deadline.

### Tasks

- [ ] Run the final submission checklist (see [execution.md](execution.md))
- [ ] Fill in [TEAM.md](TEAM.md) with the real team data
- [ ] Fill in the remaining fields in [SUBMISSION_COPY.md](SUBMISSION_COPY.md) (GitHub links, video, contract addresses)
- [ ] Fill out the DoraHacks form using the content from `SUBMISSION_COPY.md`
- [ ] Final cross-review (or a self-review after a few hours' break, to catch mistakes with fresh eyes)
- [ ] Submit officially with at least a 2-hour margin before the deadline (23:59 ET)

### Exit criterion

Submission confirmed on DoraHacks, with every required field filled in, at least 2 hours before the final deadline.

---

## Consolidated view

| Sprint | Period | Focus | Days |
|---|---|---|---|
| 1 | Aug 26–28 | Setup & base contracts | 3 |
| 2 | Aug 29–Sep 1 | Attestcoin verification flow (core) | 4 |
| 3 | Sep 2–4 | Off-chain worker + AI | 3 |
| 4 | Sep 5–7 | Frontend | 3 |
| 5 | Sep 8–9 | E2E integration + final deploy | 2 |
| 6 | Sep 10–11 | Technical rigor + documentation | 2 |
| 7 | Sep 12 | Demo + pitch | 1 |
| 8 | Sep 13 | Buffer + submission | 1 |

**Scope-cut rule if something slips:** cut the SHOULD/NICE-HAVE items listed in [architecture.md](architecture.md) first (multi-provider pool, second trigger scenario, independent verification) — never cut time from Sprint 2 (real verification) or Sprint 8 (submission buffer).
