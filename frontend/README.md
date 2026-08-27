# ClaimProof — Frontend

Next.js (App Router) storefront demo and live claims dashboard. See [docs/architecture.md](../docs/architecture.md) for the full system design and [docs/SPRINTS.md](../docs/SPRINTS.md) for this app's build plan (Sprint 4).

## Screens

- `/` — Store: a mock product with delivery protection included; "Buy" calls the `backend/cmd/api` service, which signs `createShipment` (Sepolia) and `registerOrder` (Creditcoin CC3) on the buyer's behalf.
- `/dashboard` — the connected wallet's orders with live status (`Active → Failure detected → Verifying proof → Paid`), read directly from both chains via wagmi, no manual refresh.

## Getting started

```bash
npm install
cp .env.example .env.local   # fill in contract addresses and RPC URLs
npm run dev
```

The Store's "Buy" flow requires `backend/cmd/api` running (default `http://localhost:8080`) — see `backend/.env.example`.

## Stack

Next.js 16, TypeScript (`strict`), Tailwind CSS v4, wagmi + viem, @tanstack/react-query.

## Design tokens

Defined in `src/app/globals.css` — see CLAUDE.md §5 for the full visual quality profile (typography, color semantics, light/dark themes).
