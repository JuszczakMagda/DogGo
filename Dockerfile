# --- Stage 1: Build the Go binary ---
FROM golang:1.24-alpine AS builder

# Set working directory inside the container
WORKDIR /app

# Copy go.mod and go.sum files
COPY go.mod go.sum ./

# Download dependencies
RUN go mod download

# Copy the rest of the application code
COPY . .

# Build the binary
RUN go build -o /bin/doggo

# --- Stage 2: Create a minimal runtime image ---
FROM alpine:latest

# Copy the binary from the builder stage
COPY --from=builder /bin/doggo /doggo

# Expose port 8080
EXPOSE 8080

# Run the binary
ENTRYPOINT ["/doggo"]
