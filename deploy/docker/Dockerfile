# Build stage
FROM golang:1.21-alpine AS builder

# Install build dependencies
RUN apk add --no-cache git make

# Set working directory
WORKDIR /app

# Copy go mod files
COPY go.mod go.sum ./

# Download dependencies
RUN go mod download

# Copy source code
COPY . .

# Build the application
RUN make build

# Runtime stage
FROM alpine:latest

# Install runtime dependencies
RUN apk add --no-cache ca-certificates tzdata

# Create app user
RUN addgroup -g 1000 nexuspanel && \
    adduser -D -u 1000 -G nexuspanel nexuspanel

# Set working directory
WORKDIR /app

# Copy binary from builder
COPY --from=builder /app/bin/nexuspanel-server /app/
COPY --from=builder /app/configs/config.example.yaml /app/configs/config.yaml

# Create necessary directories
RUN mkdir -p /app/data /app/logs /app/uploads && \
    chown -R nexuspanel:nexuspanel /app

# Switch to app user
USER nexuspanel

# Expose ports
EXPOSE 8080

# Health check
HEALTHCHECK --interval=30s --timeout=3s --start-period=5s --retries=3 \
  CMD wget --no-verbose --tries=1 --spider http://localhost:8080/health || exit 1

# Run the application
ENTRYPOINT ["/app/nexuspanel-server"]
