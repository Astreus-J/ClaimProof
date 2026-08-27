import { createConfig, http } from "wagmi";
import { injected, walletConnect } from "wagmi/connectors";
import { sepolia, creditcoinTestnet } from "./chains";

const walletConnectProjectId = process.env.NEXT_PUBLIC_WALLETCONNECT_PROJECT_ID;

// WalletConnect requires a Reown Cloud project ID (cloud.reown.com). It's
// optional here: without one, the injected connector (MetaMask, etc.) still
// works standalone — the safer default for a live demo, where a QR-code
// flow is one more thing that can go wrong on stage.
const connectors = walletConnectProjectId
  ? [injected(), walletConnect({ projectId: walletConnectProjectId })]
  : [injected()];

export const wagmiConfig = createConfig({
  chains: [sepolia, creditcoinTestnet],
  connectors,
  transports: {
    [sepolia.id]: http(process.env.NEXT_PUBLIC_SEPOLIA_RPC_URL),
    [creditcoinTestnet.id]: http(process.env.NEXT_PUBLIC_CREDITCOIN_RPC_URL),
  },
});

declare module "wagmi" {
  interface Register {
    config: typeof wagmiConfig;
  }
}
