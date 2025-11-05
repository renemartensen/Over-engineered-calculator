# Stage 1: Build Go binary
FROM golang:1.22-alpine AS builder
WORKDIR /app

# Copy go.mod and go.sum first for caching
COPY go.mod go.sum ./
RUN go mod download

# Copy the rest of the source
COPY . .

# Build the Go binary
RUN go build -o server ./cmd/server

# Stage 2: Create minimal runtime image
FROM alpine:latest
WORKDIR /root/

# Copy Go binary
COPY --from=builder /app/server .

# Copy static frontend
COPY --from=builder /app/web ./web

# Expose port
EXPOSE 8080

# Run the binary
CMD ["./server"]
