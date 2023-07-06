# Use official Go image as builder
FROM golang:1.18 as builder

# Set working directory
WORKDIR /app

# Copy go.mod and go.sum
COPY go.mod go.sum ./

# Download dependencies
RUN go mod download

# Copy source files
COPY . .

# Copy .env file
COPY .env .env

# Build the Go app
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main .

# Start a new stage with alpine
FROM alpine:latest

# Add ca-certificates for https calls
RUN apk --no-cache add ca-certificates

# Create appuser
RUN adduser -D -g '' appuser

# Set working directory
WORKDIR /app

# Copy binary from the builder stage
COPY --from=builder /app/main .
COPY --from=builder /app/.env .env

# Set ownership and permissions
RUN chown -R appuser:appuser /app
RUN chmod -R 755 /app

# Switch to appuser
USER appuser

# Expose the port
EXPOSE 3001

# Run the binary
CMD ["./main"]
