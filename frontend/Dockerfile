# NEXT_PUBLIC_* variables are inlined into the client bundle at build time
# (see src/lib/contracts.ts), so they must be passed as build args, not
# runtime env vars — docker-compose.yml wires them from the repo-root .env.

FROM node:22-alpine AS deps
WORKDIR /app
COPY package.json package-lock.json ./
RUN npm ci

FROM node:22-alpine AS builder
WORKDIR /app
COPY --from=deps /app/node_modules ./node_modules
COPY . .

ARG NEXT_PUBLIC_DELIVERY_TRACKER_MOCK_ADDRESS
ARG NEXT_PUBLIC_CLAIM_VAULT_ADDRESS
ARG NEXT_PUBLIC_API_BASE_URL
ARG NEXT_PUBLIC_SEPOLIA_RPC_URL
ARG NEXT_PUBLIC_CREDITCOIN_RPC_URL
ARG NEXT_PUBLIC_WALLETCONNECT_PROJECT_ID
ENV NEXT_PUBLIC_DELIVERY_TRACKER_MOCK_ADDRESS=$NEXT_PUBLIC_DELIVERY_TRACKER_MOCK_ADDRESS \
    NEXT_PUBLIC_CLAIM_VAULT_ADDRESS=$NEXT_PUBLIC_CLAIM_VAULT_ADDRESS \
    NEXT_PUBLIC_API_BASE_URL=$NEXT_PUBLIC_API_BASE_URL \
    NEXT_PUBLIC_SEPOLIA_RPC_URL=$NEXT_PUBLIC_SEPOLIA_RPC_URL \
    NEXT_PUBLIC_CREDITCOIN_RPC_URL=$NEXT_PUBLIC_CREDITCOIN_RPC_URL \
    NEXT_PUBLIC_WALLETCONNECT_PROJECT_ID=$NEXT_PUBLIC_WALLETCONNECT_PROJECT_ID

RUN npm run build

FROM node:22-alpine AS runner
WORKDIR /app
ENV NODE_ENV=production PORT=3000 HOSTNAME=0.0.0.0
RUN addgroup -S nextjs && adduser -S nextjs -G nextjs

COPY --from=builder /app/public ./public
COPY --from=builder --chown=nextjs:nextjs /app/.next/standalone ./
COPY --from=builder --chown=nextjs:nextjs /app/.next/static ./.next/static

USER nextjs
EXPOSE 3000
CMD ["node", "server.js"]
