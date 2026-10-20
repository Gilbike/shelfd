# Build frontend
FROM node:22-alpine AS frontend
WORKDIR /app

ENV PNPM_HOME="/pnpm"
ENV PATH="$PNPM_HOME:$PATH"

RUN npm install -g pnpm

COPY ./web/package.json ./web/pnpm-lock.yaml* ./
RUN --mount=type=cache,id=pnpm-store,target=/pnpm/store \
    pnpm install --frozen-lockfile

COPY ./web .
RUN pnpm build

# Build backend
FROM golang:1.27.0-alpine AS backend
WORKDIR /build

COPY go.mod go.sum ./
RUN --mount=type=cache,target=/go/pkg/mod \
    go mod download

COPY ./cmd ./cmd
COPY ./internal ./internal
COPY ./web/embed.go ./web/embed.go
COPY --from=frontend /app/build ./web/build

RUN --mount=type=cache,target=/go/pkg/mod \
    --mount=type=cache,target=/root/.cache/go-build \
    CGO_ENABLED=0 GOOS=linux go build \
    -trimpath \
    -ldflags="-s -w" \
    -o /build/server \
    ./cmd/server

# Run the app
FROM alpine:3.21 AS runner
WORKDIR /app

RUN apk --no-cache add ca-certificates tzdata && \
    addgroup -S appgroup && adduser -S appuser -G appgroup && \
    mkdir -p /app/data && chown -R appuser:appgroup /app

COPY --from=backend --chown=appuser:appgroup /build/server .

USER appuser

EXPOSE 8080

CMD ["./server"]