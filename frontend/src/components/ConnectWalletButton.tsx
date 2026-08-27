"use client";

import { useConnection, useConnect, useConnectors, useDisconnect } from "wagmi";
import { useMounted } from "@/lib/useMounted";

function truncate(address: string) {
  return `${address.slice(0, 6)}…${address.slice(-4)}`;
}

export default function ConnectWalletButton() {
  const mounted = useMounted();
  const { address, isConnected } = useConnection();
  const { mutate: connect, isPending } = useConnect();
  const { mutate: disconnect } = useDisconnect();
  const connectors = useConnectors();

  if (mounted && isConnected && address) {
    return (
      <button
        onClick={() => disconnect()}
        className="min-h-11 rounded-full border border-border bg-surface px-4 py-2.5 font-mono text-sm tabular-nums text-ink transition hover:border-accent"
        title="Disconnect wallet"
      >
        {truncate(address)}
      </button>
    );
  }

  const primaryConnector = connectors[0];

  return (
    <button
      onClick={() => primaryConnector && connect({ connector: primaryConnector })}
      disabled={isPending || !primaryConnector}
      className="min-h-11 rounded-full bg-accent px-4 py-2 text-sm font-medium text-accent-foreground transition hover:opacity-90 disabled:opacity-50"
    >
      {isPending ? "Connecting…" : "Connect Wallet"}
    </button>
  );
}
