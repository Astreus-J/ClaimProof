"use client";

import { useState } from "react";
import { useRouter } from "next/navigation";
import { useConnection } from "wagmi";
import { formatEther } from "viem";
import ConnectWalletButton from "@/components/ConnectWalletButton";
import { addTrackedOrder } from "@/lib/orders";
import { API_BASE_URL } from "@/lib/contracts";
import { useMounted } from "@/lib/useMounted";

const PRODUCT = {
  name: "Aurora Wireless Headphones",
  priceLabel: "$129.00",
  protectionAmountWei: "30000000000000000",
  slaSeconds: 90,
};

type BuyState = "idle" | "submitting" | "error";

export default function StorePage() {
  const mounted = useMounted();
  const { address, isConnected } = useConnection();
  const showConnected = mounted && isConnected;
  const [state, setState] = useState<BuyState>("idle");
  const [errorMessage, setErrorMessage] = useState<string | null>(null);
  const router = useRouter();

  async function handleBuy() {
    if (!address) return;
    setState("submitting");
    setErrorMessage(null);

    const orderId = Date.now();
    try {
      const res = await fetch(`${API_BASE_URL}/api/orders`, {
        method: "POST",
        headers: { "Content-Type": "application/json" },
        body: JSON.stringify({
          orderId,
          buyer: address,
          protectionAmountWei: PRODUCT.protectionAmountWei,
          slaSeconds: PRODUCT.slaSeconds,
        }),
      });

      if (!res.ok) {
        const body = (await res.json().catch(() => null)) as { error?: string } | null;
        throw new Error(body?.error ?? `Purchase failed (HTTP ${res.status})`);
      }

      addTrackedOrder({
        orderId: String(orderId),
        buyer: address,
        productName: PRODUCT.name,
        protectionAmountWei: PRODUCT.protectionAmountWei,
        slaSeconds: PRODUCT.slaSeconds,
        createdAt: Date.now(),
      });

      router.push("/dashboard");
    } catch (err) {
      setState("error");
      setErrorMessage(err instanceof Error ? err.message : "Something went wrong — please try again.");
    }
  }

  return (
    <div className="mx-auto max-w-2xl px-6 py-16">
      <p className="font-mono text-xs uppercase tracking-wide text-accent">ClaimProof Storefront</p>
      <h1 className="mt-2 font-display text-4xl italic text-ink">{PRODUCT.name}</h1>
      <p className="mt-4 text-muted">
        A claim that pays itself — because the proof is faster than the adjuster. Add delivery
        protection at checkout: if this order isn&apos;t confirmed delivered in time, you&apos;re
        refunded automatically the moment the failure is verified on-chain.
      </p>

      <div className="mt-8 rounded-2xl border border-border bg-surface p-6">
        <div className="flex items-baseline justify-between">
          <span className="text-lg text-ink">{PRODUCT.name}</span>
          <span className="font-mono text-lg tabular-nums text-ink">{PRODUCT.priceLabel}</span>
        </div>
        <div className="mt-3 flex items-baseline justify-between border-t border-border pt-3">
          <span className="text-sm text-muted">Delivery protection (included)</span>
          <span className="font-mono text-sm tabular-nums text-gold">
            up to {formatEther(BigInt(PRODUCT.protectionAmountWei))} CTC refund
          </span>
        </div>

        {!showConnected ? (
          <div className="mt-6 flex flex-col items-center gap-3 rounded-xl border border-dashed border-border p-6 text-center">
            <p className="text-sm text-muted">Connect a wallet to buy — it&apos;s where your refund would go.</p>
            <ConnectWalletButton />
          </div>
        ) : (
          <button
            onClick={handleBuy}
            disabled={state === "submitting"}
            className="mt-6 w-full rounded-full bg-accent px-4 py-3 font-medium text-accent-foreground transition hover:opacity-90 disabled:opacity-50"
          >
            {state === "submitting" ? "Processing purchase…" : "Buy with Delivery Protection"}
          </button>
        )}

        {state === "error" && errorMessage && (
          <p className="mt-3 rounded-lg border border-danger/30 bg-danger/10 px-3 py-2 text-sm text-danger">
            {errorMessage}
          </p>
        )}
      </div>
    </div>
  );
}
