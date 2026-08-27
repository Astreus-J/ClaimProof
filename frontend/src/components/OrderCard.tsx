"use client";

import { useEffect, useRef, useState } from "react";
import { useWriteContract, useWaitForTransactionReceipt } from "wagmi";
import { formatEther } from "viem";
import StatusBadge, { type BadgeStage } from "./StatusBadge";
import Spinner from "./Spinner";
import { useOrderStatus } from "@/lib/useOrderStatus";
import { deliveryTrackerMockAbi } from "@/lib/abi";
import { DELIVERY_TRACKER_MOCK_ADDRESS } from "@/lib/contracts";
import { formatContractError } from "@/lib/errors";
import { sepolia } from "@/lib/chains";
import type { TrackedOrder } from "@/lib/orders";

// How long the "Failure detected" flash shows before settling into the
// steady "Verifying proof" state — both are real, distinct states (per
// CLAUDE.md §5), but only the transition moment itself is worth announcing.
const FAILURE_FLASH_MS = 4_000;

// A few seconds of margin absorbs clock skew between the browser's local
// clock (what `now` is built from) and the chain's block.timestamp, which
// is what the contract actually checks — without it, a fast local clock
// can enable the simulate button slightly before the chain agrees the SLA
// has expired, and the transaction reverts with SlaNotExpired.
const CLOCK_SKEW_MARGIN_SECONDS = 5;

function useNowSeconds(everyMs: number, enabled: boolean) {
  const [now, setNow] = useState(() => Math.floor(Date.now() / 1000));
  useEffect(() => {
    if (!enabled) return;
    const id = setInterval(() => setNow(Math.floor(Date.now() / 1000)), everyMs);
    return () => clearInterval(id);
  }, [everyMs, enabled]);
  return now;
}

function formatCountdown(totalSeconds: number): string {
  const clamped = Math.max(0, totalSeconds);
  const minutes = Math.floor(clamped / 60);
  const seconds = clamped % 60;
  return `${minutes}:${String(seconds).padStart(2, "0")}`;
}

function OrderCardSkeleton() {
  return (
    <div className="rounded-2xl border border-border bg-surface p-6" aria-hidden>
      <div className="flex items-start justify-between gap-4">
        <div className="flex flex-col gap-2">
          <div className="skeleton h-5 w-40" />
          <div className="skeleton h-3 w-24" />
        </div>
        <div className="skeleton h-6 w-20 rounded-full" />
      </div>
      <div className="mt-4 grid grid-cols-2 gap-3">
        <div className="flex flex-col gap-2">
          <div className="skeleton h-3 w-28" />
          <div className="skeleton h-4 w-16" />
        </div>
        <div className="flex flex-col gap-2">
          <div className="skeleton h-3 w-20" />
          <div className="skeleton h-4 w-16" />
        </div>
      </div>
      <div className="skeleton mt-4 h-11 w-full" />
    </div>
  );
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
  const effectiveDeadline =
    status.slaDeadline !== undefined ? Number(status.slaDeadline) + CLOCK_SKEW_MARGIN_SECONDS : undefined;
  const secondsRemaining = effectiveDeadline !== undefined ? effectiveDeadline - now : undefined;
  const canSimulateFailure = status.stage === "active" && secondsRemaining !== undefined && secondsRemaining <= 0;

  const { writeContract, data: txHash, isPending, error: writeError } = useWriteContract();
  const { isLoading: isMining, isSuccess: isMined } = useWaitForTransactionReceipt({
    hash: txHash,
    chainId: sepolia.id,
  });

  if (status.isLoading) return <OrderCardSkeleton />;

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
            className="mt-4 flex min-h-11 w-full items-center justify-center gap-2 rounded-full border border-danger px-4 py-2 text-sm font-medium text-danger transition hover:bg-danger hover:text-danger-foreground disabled:cursor-not-allowed disabled:opacity-40 disabled:hover:bg-transparent"
          >
            {(isPending || isMining) && <Spinner className="h-4 w-4" />}
            {isPending || isMining
              ? "Reporting failure…"
              : canSimulateFailure
                ? "Simulate Delivery Failure"
                : `Available in ${secondsRemaining !== undefined ? formatCountdown(secondsRemaining) : "…"}`}
          </button>
          {writeError && (
            <p className="mt-2 text-sm text-danger">{formatContractError(writeError)}</p>
          )}
          {isMined && (
            <p className="badge-transition mt-2 text-sm text-accent">
              Failure reported — the worker is now watching for Attestcoin attestation.
            </p>
          )}
        </>
      )}
    </div>
  );
}
