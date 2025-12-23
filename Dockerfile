# Build stage
FROM golang:1.25.4-alpine AS builder

WORKDIR /app

# Install build dependencies
RUN apk add --no-cache git

# Copy dependency files
COPY go.mod go.sum ./
RUN go mod download

# Copy source code
COPY . .

# Build the API binary
# CGO_ENABLED=0 ensures a static binary for minimal runtime images
# ldflags -s -w reduces binary size
RUN CGO_ENABLED=0 GOOS=linux go build \
    -trimpath \
    -ldflags "-s -w" \
    -o api \
    ./cmd/api

# Run stage
FROM alpine:latest

WORKDIR /app

# Install ca-certificates for Spotify API (HTTPS)
RUN apk add --no-cache ca-certificates

# Copy the binary from the builder
COPY --from=builder /app/api .

# Expose API port
EXPOSE 8080

# Run the API
ENTRYPOINT ["./api"]
