"use client";

import { useEffect, useRef, useState } from "react";
import { useWriteContract, useWaitForTransactionReceipt } from "wagmi";
import { formatEther } from "viem";
import StatusBadge, { type BadgeStage } from "./StatusBadge";
import { useOrderStatus } from "@/lib/useOrderStatus";
import { deliveryTrackerMockAbi } from "@/lib/abi";
import { DELIVERY_TRACKER_MOCK_ADDRESS } from "@/lib/contracts";
import { sepolia } from "@/lib/chains";
import type { TrackedOrder } from "@/lib/orders";

// How long the "Failure detected" flash shows before settling into the
// steady "Verifying proof" state — both are real, distinct states (per
// CLAUDE.md §5), but only the transition moment itself is worth announcing.
const FAILURE_FLASH_MS = 4_000;

function useNowSeconds(everyMs: number, enabled: boolean) {
  const [now, setNow] = useState(() => Math.floor(Date.now() / 1000));
  useEffect(() => {
    if (!enabled) return;
    const id = setInterval(() => setNow(Math.floor(Date.now() / 1000)), everyMs);
    return () => clearInterval(id);
  }, [everyMs, enabled]);
  return now;
}

export default function OrderCard({ order }: { order: TrackedOrder }) {
  const orderId = BigInt(order.orderId);
  const status = useOrderStatus(orderId);

  const [justFailed, setJustFailed] = useState(false);
  const prevStage = useRef(status.stage);
  useEffect(() => {
    if (prevStage.current === "active" && status.stage === "failed") {
      setJustFailed(true);
      const t = setTimeout(() => setJustFailed(false), FAILURE_FLASH_MS);
      return () => clearTimeout(t);
    }
    prevStage.current = status.stage;
  }, [status.stage]);

  const badgeStage: BadgeStage =
    status.stage === "failed" ? (justFailed ? "failure-detected" : "verifying-proof") : status.stage;

  const now = useNowSeconds(1_000, status.stage === "active");
  const slaExpired = status.slaDeadline !== undefined && now >= Number(status.slaDeadline);
  const canSimulateFailure = status.stage === "active" && slaExpired;

  const { writeContract, data: txHash, isPending, error: writeError } = useWriteContract();
  const { isLoading: isMining, isSuccess: isMined } = useWaitForTransactionReceipt({
    hash: txHash,
    chainId: sepolia.id,
  });

  return (
    <div className="rounded-2xl border border-border bg-surface p-6">
      <div className="flex items-start justify-between gap-4">
        <div>
          <h3 className="font-display text-lg text-ink">{order.productName}</h3>
          <p className="mt-1 font-mono text-xs tabular-nums text-muted">Order #{order.orderId}</p>
        </div>
        <StatusBadge stage={badgeStage} />
      </div>

      <dl className="mt-4 grid grid-cols-2 gap-3 text-sm">
        <div>
          <dt className="text-muted">Protection amount</dt>
          <dd className="font-mono tabular-nums text-gold">
            {formatEther(BigInt(order.protectionAmountWei))} CTC
          </dd>
        </div>
        <div>
          <dt className="text-muted">SLA deadline</dt>
          <dd className="font-mono tabular-nums text-ink">
            {status.slaDeadline ? new Date(Number(status.slaDeadline) * 1000).toLocaleTimeString() : "—"}
          </dd>
        </div>
      </dl>

      {status.isError && (
        <p className="mt-4 rounded-lg border border-danger/30 bg-danger/10 px-3 py-2 text-sm text-danger">
          Couldn&apos;t read this order&apos;s status just now — retrying automatically.
        </p>
      )}

      {status.stage === "active" && (
        <>
          <button
            onClick={() =>
              writeContract({
                address: DELIVERY_TRACKER_MOCK_ADDRESS,
                abi: deliveryTrackerMockAbi,
                functionName: "reportDeliveryFailure",
                args: [orderId],
                chainId: sepolia.id,
              })
            }
            disabled={!canSimulateFailure || isPending || isMining}
            className="mt-4 min-h-11 w-full rounded-full border border-danger px-4 py-2 text-sm font-medium text-danger transition hover:bg-danger hover:text-danger-foreground disabled:cursor-not-allowed disabled:opacity-40 disabled:hover:bg-transparent"
          >
            {isPending || isMining
              ? "Reporting failure…"
              : canSimulateFailure
                ? "Simulate Delivery Failure"
                : `Available at ${status.slaDeadline ? new Date(Number(status.slaDeadline) * 1000).toLocaleTimeString() : "…"}`}
          </button>
          {writeError && (
            <p className="mt-2 text-sm text-danger">{writeError.message.split("\n")[0]}</p>
          )}
          {isMined && (
            <p className="mt-2 text-sm text-accent">
              Failure reported — the worker is now watching for Attestcoin attestation.
            </p>
          )}
        </>
      )}
    </div>
  );
}
