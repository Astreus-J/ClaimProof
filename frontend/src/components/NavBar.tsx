import Link from "next/link";
import ConnectWalletButton from "./ConnectWalletButton";

export default function NavBar() {
  return (
    <header className="border-b border-border">
      <div className="mx-auto flex max-w-4xl flex-wrap items-center justify-between gap-x-4 gap-y-2 px-6 py-4">
        <Link href="/" className="font-display text-xl italic text-ink">
          ClaimProof
        </Link>
        <nav className="flex flex-wrap items-center gap-x-6 gap-y-2">
          <Link href="/" className="text-sm text-muted transition hover:text-ink">
            Store
          </Link>
          <Link href="/dashboard" className="text-sm text-muted transition hover:text-ink">
            Dashboard
          </Link>
          <ConnectWalletButton />
        </nav>
      </div>
    </header>
  );
}
