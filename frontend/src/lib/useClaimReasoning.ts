"use client";

import { useQuery } from "@tanstack/react-query";
import { API_BASE_URL } from "./contracts";

export interface ClaimReasoning {
  reasoning: string;
  suggestedPayoutWei: string;
}

/**
 * Fetches the AI claims agent's payout reasoning for an order from
 * cmd/api's GET /api/claims/{orderId} — purely informational (see
 * backend/internal/reasoningreporter's package doc), so a 404 (nothing
 * reported yet, or never will be) is a normal, silent "no data" case, not
 * an error state worth surfacing to the user.
 */
export function useClaimReasoning(orderId: string, enabled: boolean) {
  return useQuery<ClaimReasoning | null>({
    queryKey: ["claim-reasoning", orderId],
    enabled,
    refetchInterval: (query) => (query.state.data ? false : 5_000),
    queryFn: async () => {
      const res = await fetch(`${API_BASE_URL}/api/claims/${orderId}`);
      if (res.status === 404) return null;
      if (!res.ok) throw new Error(`Failed to fetch AI reasoning (HTTP ${res.status})`);
      return (await res.json()) as ClaimReasoning;
    },
  });
}
