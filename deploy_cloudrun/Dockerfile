# Stage 1: Build the application
FROM golang:1.22 AS builder

# Set environment variables
WORKDIR /app

# Copy app files to working directory
COPY . .

# Build the application binary
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o main ./cmd/api/main.go

# Stage 2: Create a minimal image to run the application
FROM gcr.io/distroless/base-debian10

# Set working directory in container
WORKDIR /app

# Copy the compiled binary from the builder stage
COPY --from=builder /app/main .

# Expose the application port
EXPOSE 8080

ENV WEATHER_API_KEY=1234
ENV SERVICE_PORT=:8080

# Command to run the application
CMD ["/app/main"]