# syntax=docker/dockerfile:1
FROM golang:latest AS builder
WORKDIR /build
COPY . .
RUN --mount=type=cache,target=/go/pkg/mod \
    go mod download -x
RUN --mount=type=cache,target=/go/pkg/mod \
    --mount=type=cache,target=/root/.cache/go-build \
    go build -ldflags="-s -w" -trimpath ./cmd/app

FROM debian:stable-slim AS app
WORKDIR /app
COPY --from=builder /build/app .
CMD [ "./app" ]
