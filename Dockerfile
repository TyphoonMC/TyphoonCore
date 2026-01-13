# Build stage
FROM golang:1.23-alpine AS builder

WORKDIR /app

# Install build dependencies
RUN apk add --no-cache git ca-certificates

# Copy go mod files first for caching
COPY go.mod go.sum ./
RUN go mod download

# Copy source code
COPY . .

# Build with optimizations
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build \
    -ldflags="-w -s" \
    -o /typhoon ./cmd/typhoon

# Final stage
FROM alpine:3.21

# Add non-root user
RUN addgroup -S typhoon && adduser -S typhoon -G typhoon

# Install netcat for health check
RUN apk add --no-cache netcat-openbsd

WORKDIR /app

# Copy binary from builder
COPY --from=builder /typhoon /app/typhoon
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/

# Copy default config
COPY config.json /app/config.json
COPY favicon.png /app/favicon.png

# Create data directories
RUN mkdir -p /app/modules && chown -R typhoon:typhoon /app

USER typhoon

# Expose Minecraft port
EXPOSE 25565

# Health check
HEALTHCHECK --interval=30s --timeout=10s --start-period=5s --retries=3 \
    CMD nc -z localhost 25565 || exit 1

ENTRYPOINT ["/app/typhoon"]
