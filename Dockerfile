# syntax=docker/dockerfile:1
FROM golang:latest AS builder
WORKDIR /build
COPY . .
RUN --mount=type=cache,target=/go/pkg/mod \
    go mod download -x
RUN --mount=type=cache,target=/go/pkg/mod \
    --mount=type=cache,target=/root/.cache/go-build \
    go build -ldflags="-s -w" -trimpath -o server ./cmd/webapiserver/main.go

FROM debian:stable-slim AS deploy
WORKDIR /app
COPY --from=builder /build/server .
EXPOSE 8080
CMD [ "./server" ]
