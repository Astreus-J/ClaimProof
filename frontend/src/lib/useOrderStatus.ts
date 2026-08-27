"use client";

import { useMemo } from "react";
import { useReadContract, useWatchContractEvent } from "wagmi";
import { deliveryTrackerMockAbi, claimVaultAbi, ShipmentStatus } from "./abi";
import { DELIVERY_TRACKER_MOCK_ADDRESS, CLAIM_VAULT_ADDRESS } from "./contracts";
import { sepolia, creditcoinTestnet } from "./chains";

export type OrderStage = "active" | "failed" | "paid";

export interface OrderStatus {
  isLoading: boolean;
  isError: boolean;
  stage: OrderStage;
  protectionAmountWei: bigint | undefined;
  slaDeadline: bigint | undefined;
}

/**
 * Reads an order's live status by combining a Sepolia read (shipment
 * status) and a Creditcoin read (claim status), and re-fetches both
 * whenever their respective on-chain events fire — so the dashboard
 * updates without a manual reload.
 */
export function useOrderStatus(orderId: bigint): OrderStatus {
  const shipment = useReadContract({
    address: DELIVERY_TRACKER_MOCK_ADDRESS,
    abi: deliveryTrackerMockAbi,
    functionName: "shipments",
    args: [orderId],
    chainId: sepolia.id,
    query: { refetchInterval: 8_000 },
  });

  const order = useReadContract({
    address: CLAIM_VAULT_ADDRESS,
    abi: claimVaultAbi,
    functionName: "orders",
    args: [orderId],
    chainId: creditcoinTestnet.id,
    query: { refetchInterval: 8_000 },
  });

  useWatchContractEvent({
    address: DELIVERY_TRACKER_MOCK_ADDRESS,
    abi: deliveryTrackerMockAbi,
    eventName: "DeliveryFailed",
    chainId: sepolia.id,
    args: { orderId },
    onLogs: () => shipment.refetch(),
  });

  useWatchContractEvent({
    address: CLAIM_VAULT_ADDRESS,
    abi: claimVaultAbi,
    eventName: "ClaimPaid",
    chainId: creditcoinTestnet.id,
    args: { orderId },
    onLogs: () => order.refetch(),
  });

  const [, , slaDeadline, shipmentStatus] = shipment.data ?? [];
  const [, protectionAmountWei, claimed] = order.data ?? [];

  const stage = useMemo<OrderStage>(() => {
    if (claimed) return "paid";
    if (shipmentStatus === ShipmentStatus.Failed) return "failed";
    return "active";
  }, [claimed, shipmentStatus]);

  return {
    isLoading: shipment.isLoading || order.isLoading,
    isError: shipment.isError || order.isError,
    stage,
    protectionAmountWei,
    slaDeadline,
  };
}
