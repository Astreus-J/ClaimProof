"use client";

import { useSyncExternalStore } from "react";
import { useConnection } from "wagmi";
import ConnectWalletButton from "@/components/ConnectWalletButton";
import OrderCard from "@/components/OrderCard";
import { getTrackedOrdersSnapshot, getServerOrdersSnapshot, subscribeToTrackedOrders } from "@/lib/orders";
import { useMounted } from "@/lib/useMounted";

export default function DashboardPage() {
  const mounted = useMounted();
  const { address, isConnected } = useConnection();
  const showConnected = mounted && isConnected;
  const orders = useSyncExternalStore(subscribeToTrackedOrders, getTrackedOrdersSnapshot, getServerOrdersSnapshot);

  const myOrders = orders.filter((o) => address && o.buyer.toLowerCase() === address.toLowerCase());

  return (
    <div className="mx-auto max-w-3xl px-6 py-16">
      <h1 className="font-display text-3xl italic text-ink">Your Claims</h1>
      <p className="mt-2 text-muted">Live status, straight from the chain — no refresh needed.</p>

      {!showConnected ? (
        <div className="mt-10 flex flex-col items-center gap-3 rounded-xl border border-dashed border-border p-10 text-center">
          <p className="text-sm text-muted">Connect the wallet you bought with to see your orders.</p>
          <ConnectWalletButton />
        </div>
      ) : myOrders.length === 0 ? (
        <div className="mt-10 rounded-xl border border-dashed border-border p-10 text-center">
          <p className="text-sm text-muted">No orders yet for this wallet — buy some protection first.</p>
        </div>
      ) : (
        <div className="mt-8 flex flex-col gap-4">
          {myOrders.map((order) => (
            <OrderCard key={order.orderId} order={order} />
          ))}
        </div>
      )}
    </div>
  );
}
