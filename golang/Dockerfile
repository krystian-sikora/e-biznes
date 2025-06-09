# Build stage
FROM golang:1.24-alpine AS builder

WORKDIR /app

# Install build dependencies for CGO
RUN apk --no-cache add gcc musl-dev

# Copy go mod and sum files
COPY go.mod ./
COPY go.sum* ./

# Download dependencies
RUN go mod download

# Copy source code
COPY . .

# Build the application
RUN CGO_ENABLED=1 GOOS=linux go build -o /app/main ./src

# Final stage
FROM alpine:latest

WORKDIR /app

# Install required packages including SQLite
RUN apk --no-cache add ca-certificates sqlite

# Create data directory with proper permissions
RUN mkdir -p /app/data && chmod 777 /app/data

# Copy binary from builder
COPY --from=builder /app/main .
COPY test.db /app/data/test.db
RUN chmod 777 /app/data/test.db

# Create a non-root user to run the application
RUN adduser -D appuser
USER appuser

# Expose the port the app runs on
EXPOSE 8080

# Run the application
CMD ["/app/main"]