export type BadgeStage = "loading" | "active" | "failure-detected" | "verifying-proof" | "paid";

const STYLES: Record<BadgeStage, { label: string; dot: string; classes: string }> = {
  loading: {
    label: "Loading…",
    dot: "bg-muted",
    classes: "border-border bg-surface text-muted",
  },
  active: {
    label: "Active",
    dot: "bg-muted",
    classes: "border-border bg-surface text-ink",
  },
  "failure-detected": {
    label: "Failure detected",
    dot: "bg-danger-foreground",
    classes: "border-danger bg-danger text-danger-foreground",
  },
  "verifying-proof": {
    label: "Verifying proof",
    dot: "bg-accent",
    classes: "border-accent bg-surface text-ink",
  },
  paid: {
    label: "Paid",
    dot: "bg-gold-foreground",
    classes: "border-gold bg-gold text-gold-foreground",
  },
};

export default function StatusBadge({ stage }: { stage: BadgeStage }) {
  const style = STYLES[stage];
  return (
    <span
      key={stage}
      role="status"
      aria-live="polite"
      className={`badge-transition inline-flex items-center gap-2 whitespace-nowrap rounded-full border px-3 py-1 text-sm font-medium ${style.classes}`}
    >
      <span
        className={`h-2 w-2 rounded-full ${style.dot} ${stage === "verifying-proof" || stage === "loading" ? "animate-pulse" : ""}`}
        aria-hidden
      />
      {style.label}
    </span>
  );
}
