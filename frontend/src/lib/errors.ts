import { BaseError, ContractFunctionRevertedError } from "viem";

// Mirrors contracts/src/DeliveryTrackerMock.sol's custom errors.
const FRIENDLY_REVERT_REASONS: Record<string, string> = {
  SlaNotExpired: "The delivery SLA hasn't expired yet — try again after the deadline shown above.",
  ShipmentNotPending: "This order's delivery status has already been reported.",
  ShipmentNotFound: "This order isn't registered on Sepolia yet.",
};

/**
 * Extracts a clean, user-facing message from a wagmi/viem write error: the
 * decoded contract revert reason when the ABI has it, otherwise viem's own
 * one-line summary. Never surfaces the raw multi-line message (docs links,
 * contract call params, stack context) directly to the UI.
 */
export function formatContractError(error: Error): string {
  if (error instanceof BaseError) {
    const revertError = error.walk((e) => e instanceof ContractFunctionRevertedError);
    if (revertError instanceof ContractFunctionRevertedError) {
      const errorName = revertError.data?.errorName;
      if (errorName && errorName in FRIENDLY_REVERT_REASONS) {
        return FRIENDLY_REVERT_REASONS[errorName];
      }
      return revertError.shortMessage;
    }
    return error.shortMessage;
  }
  return error.message;
}

/**
 * Same idea for a plain fetch()-based request (the Store's buy flow): a
 * network-level failure throws a generic, browser-specific TypeError
 * ("Failed to fetch", "NetworkError when attempting to fetch resource.")
 * that means nothing to a buyer — replace it with one clear sentence.
 */
export function formatRequestError(error: unknown): string {
  if (error instanceof TypeError) {
    return "Couldn't reach the ClaimProof service — check your connection and try again.";
  }
  if (error instanceof Error) return error.message;
  return "Something went wrong — please try again.";
}
