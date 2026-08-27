/**
 * Minimal ABI subsets — only the functions/events the frontend actually
 * calls or watches. The full ABIs live in contracts/out/ after `forge build`.
 */

export const deliveryTrackerMockAbi = [
  {
    type: "function",
    name: "shipments",
    stateMutability: "view",
    inputs: [{ name: "", type: "uint256" }],
    outputs: [
      { name: "buyer", type: "address" },
      { name: "orderId", type: "uint256" },
      { name: "slaDeadline", type: "uint256" },
      { name: "status", type: "uint8" },
    ],
  },
  {
    type: "function",
    name: "reportDeliveryFailure",
    stateMutability: "nonpayable",
    inputs: [{ name: "orderId", type: "uint256" }],
    outputs: [],
  },
  {
    type: "event",
    name: "DeliveryFailed",
    inputs: [
      { name: "orderId", type: "uint256", indexed: true },
      { name: "buyer", type: "address", indexed: true },
      { name: "timestamp", type: "uint256", indexed: false },
    ],
    anonymous: false,
  },
] as const;

export const claimVaultAbi = [
  {
    type: "function",
    name: "orders",
    stateMutability: "view",
    inputs: [{ name: "", type: "uint256" }],
    outputs: [
      { name: "buyer", type: "address" },
      { name: "protectionAmount", type: "uint256" },
      { name: "claimed", type: "bool" },
    ],
  },
  {
    type: "event",
    name: "ClaimPaid",
    inputs: [
      { name: "orderId", type: "uint256", indexed: true },
      { name: "buyer", type: "address", indexed: true },
      { name: "amount", type: "uint256", indexed: false },
      { name: "queryId", type: "bytes32", indexed: true },
    ],
    anonymous: false,
  },
] as const;

/** Mirrors DeliveryTrackerMock.ShipmentStatus. */
export enum ShipmentStatus {
  Pending = 0,
  Delivered = 1,
  Failed = 2,
}
