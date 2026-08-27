# Demo assets

Raw footage (`raw-footage/`, gitignored — not committed, ask the team for the files) for editing into the final 3-5 min demo video per [docs/execution.md](../docs/execution.md)'s script:

- `01-buy-and-simulate-failure.webm` — real screen capture: connect wallet → buy protection → live SLA countdown → click "Simulate Delivery Failure" → card transitions to "Verifying proof". Captured against live testnet (real `createShipment`/`registerOrder`/`reportDeliveryFailure` transactions), not simulated.
- `02-paid-result.webm` — real screen capture of the same order's dashboard card once the worker's pipeline (attestation → proof → AI suggestion → `submitClaim`) completed: "Paid" badge, protection amount, and the AI's reasoning.

No narration or editing — cut/splice per the script in `docs/execution.md`, with a time-lapse or jump cut between the two clips to skip the ~7-9 minute real Attestcoin attestation wait that happens between them (not sped-up or faked — it's a real backend processing delay that isn't part of the story to tell live).

Once the final edited video is ready, this file should link to the published upload, replacing this note.
