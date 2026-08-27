# Builds either cmd/api or cmd/worker from the same module, selected via
# CMD_PATH (see docker-compose.yml). Static binary + distroless runtime:
# no shell, no package manager, minimal attack surface and image size.

FROM golang:1.24-alpine AS build
WORKDIR /src

COPY go.mod go.sum ./
RUN go mod download

COPY . .

ARG CMD_PATH=./cmd/api
RUN CGO_ENABLED=0 GOOS=linux go build -trimpath -ldflags="-s -w" -o /out/service ${CMD_PATH}

FROM gcr.io/distroless/static-debian12:nonroot
COPY --from=build /out/service /service
USER nonroot:nonroot
ENTRYPOINT ["/service"]
