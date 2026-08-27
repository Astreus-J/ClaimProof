"use client";

import { useSyncExternalStore } from "react";

const subscribe = () => () => {};

/**
 * True only after the client has hydrated. wagmi can auto-reconnect a
 * previously-connected wallet before hydration finishes, so `isConnected`
 * can legitimately differ between the server render (always false) and the
 * first client render (possibly true) — gate any UI branch that depends on
 * connection state behind this to avoid a hydration mismatch.
 */
export function useMounted(): boolean {
  return useSyncExternalStore(
    subscribe,
    () => true,
    () => false,
  );
}
