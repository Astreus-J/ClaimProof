/**
 * There's no on-chain index of "orders belonging to buyer X" and no
 * database in this MVP (see docs/architecture.md: "Indexer/Database ...
 * never the source of truth for a payout") — so the browser remembers
 * which orders it created. This is purely a local convenience list; every
 * status shown for an order is always read fresh from the chain, never from
 * this cache. Exposed via useSyncExternalStore's subscribe/getSnapshot
 * contract so React components can read it without effects.
 */

export interface TrackedOrder {
  orderId: string;
  buyer: `0x${string}`;
  productName: string;
  protectionAmountWei: string;
  slaSeconds: number;
  createdAt: number;
}

const STORAGE_KEY = "claimproof:orders";
const listeners = new Set<() => void>();
let cache: TrackedOrder[] | null = null;

function readFromStorage(): TrackedOrder[] {
  if (typeof window === "undefined") return [];
  try {
    const raw = window.localStorage.getItem(STORAGE_KEY);
    return raw ? (JSON.parse(raw) as TrackedOrder[]) : [];
  } catch {
    return [];
  }
}

export function getTrackedOrdersSnapshot(): TrackedOrder[] {
  if (cache === null) {
    cache = readFromStorage();
  }
  return cache;
}

const EMPTY_ORDERS: TrackedOrder[] = [];

export function getServerOrdersSnapshot(): TrackedOrder[] {
  return EMPTY_ORDERS;
}

export function subscribeToTrackedOrders(onChange: () => void): () => void {
  listeners.add(onChange);
  return () => listeners.delete(onChange);
}

export function addTrackedOrder(order: TrackedOrder): void {
  if (typeof window === "undefined") return;
  try {
    window.localStorage.setItem(STORAGE_KEY, JSON.stringify([order, ...readFromStorage()]));
    cache = null;
    listeners.forEach((listener) => listener());
  } catch {
    // localStorage can throw (private browsing, quota exceeded) — losing
    // the local list doesn't corrupt any on-chain state, safe to ignore.
  }
}
