# Stage 1: Build Stage
FROM golang:1.22 AS builder

# Set the working directory inside the container
WORKDIR /app

# Copy only the go.mod and go.sum files to download dependencies
COPY go.mod go.sum ./

# Download all dependencies
RUN go mod download

# Copy all application code to the container
COPY . .

# Build the application
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o app ./cmd/main.go

# Stage 2: Runtime Stage
FROM debian:bullseye-slim

# Set the working directory
WORKDIR /app

# Copy the built binary from the builder stage
COPY --from=builder /app/app .
COPY ./.env .

RUN chmod +x ./app
# Expose port if required (Not mandatory, no explicit service port is needed for this app)
# EXPOSE 8080

# Command to run the application, allowing runtime arguments
ENTRYPOINT ["./app"]
