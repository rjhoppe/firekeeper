# ---- Build Stage ----
FROM golang:1.23-alpine AS builder

# Install git (required for go get) and ca-certificates
RUN apk add --no-cache git ca-certificates

# Set the working directory inside the container
WORKDIR /app

# Copy go.mod and go.sum first for better caching
COPY go.mod go.sum ./
RUN go mod download

# Copy the rest of the source code
COPY . .

# Build the Go binary
RUN go build -o firekeeper main.go

# ---- Run Stage ----
FROM alpine:latest

RUN apk --no-cache add ca-certificates

WORKDIR /root/

# Copy the binary from the builder stage
COPY --from=builder /app/firekeeper .

# For health check endpoint
EXPOSE 8081

CMD ["./firekeeper"]
