# syntax=docker/dockerfile:1
FROM golang:latest AS base
WORKDIR /base
COPY . .
RUN --mount=type=cache,target=/go/pkg/mod \
    go mod download -x

FROM base AS migrate-builder
RUN --mount=type=cache,target=/go/pkg/mod \
    --mount=type=cache,target=/root/.cache/go-build \
    go build -ldflags="-s -w" -trimpath ./cmd/migrate
FROM debian:stable-slim AS migrate
WORKDIR /migrate
COPY --from=migrate-builder /base/migrate .

FROM base AS app-builder
RUN --mount=type=cache,target=/go/pkg/mod \
    --mount=type=cache,target=/root/.cache/go-build \
    go build -ldflags="-s -w" -trimpath ./cmd/app
FROM debian:stable-slim AS app
WORKDIR /app
COPY --from=app-builder /base/app .
CMD [ "./app" ]
