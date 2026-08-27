import type { Address } from "viem";

// Next.js only inlines `process.env.NEXT_PUBLIC_*` into the client bundle
// when the access is a static, literal property lookup — a dynamic
// `process.env[key]` helper can't be detected at build time and silently
// resolves to undefined in the browser. Each variable is therefore read
// with its own literal statement below.

function required(name: string, value: string | undefined): Address {
  if (!value) {
    throw new Error(`${name} is not set — check frontend/.env.local`);
  }
  return value as Address;
}

export const DELIVERY_TRACKER_MOCK_ADDRESS = required(
  "NEXT_PUBLIC_DELIVERY_TRACKER_MOCK_ADDRESS",
  process.env.NEXT_PUBLIC_DELIVERY_TRACKER_MOCK_ADDRESS,
);
export const CLAIM_VAULT_ADDRESS = required(
  "NEXT_PUBLIC_CLAIM_VAULT_ADDRESS",
  process.env.NEXT_PUBLIC_CLAIM_VAULT_ADDRESS,
);

export const API_BASE_URL = process.env.NEXT_PUBLIC_API_BASE_URL ?? "http://localhost:8080";

/** Attestcoin's chain key identifying Ethereum Sepolia as a source chain. */
export const SEPOLIA_CHAIN_KEY = 1;
