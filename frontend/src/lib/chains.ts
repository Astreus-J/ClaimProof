import { defineChain } from "viem";
import { sepolia } from "viem/chains";

export { sepolia };

/**
 * Creditcoin CC3 Testnet — execution and payout chain, hosts ClaimVault and
 * the Attestcoin precompile at 0x0FD2. Not in viem/chains, so defined here.
 */
export const creditcoinTestnet = defineChain({
  id: 102031,
  name: "Creditcoin CC3 Testnet",
  nativeCurrency: { name: "Creditcoin", symbol: "CTC", decimals: 18 },
  rpcUrls: {
    default: { http: [process.env.NEXT_PUBLIC_CREDITCOIN_RPC_URL ?? "https://rpc.cc3-testnet.creditcoin.network"] },
  },
  blockExplorers: {
    default: { name: "Blockscout", url: "https://creditcoin-testnet.blockscout.com" },
  },
  testnet: true,
});
